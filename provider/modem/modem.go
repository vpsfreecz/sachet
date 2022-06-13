package modem

import (
	"io"
	"log"
	"time"

	"github.com/messagebird/sachet"
	"github.com/warthog618/modem/at"
	"github.com/warthog618/modem/gsm"
	"github.com/warthog618/modem/serial"
	"github.com/warthog618/modem/trace"
)

type Config struct {
	Device    string        `yaml:device`
	Mode      string        `yaml:mode`
	BaudRate  int           `yaml:baud_rate`
	Timeout   time.Duration `yaml:timeout`
	Verbose   bool          `yaml:verbose`
	Hex       bool          `yaml:hex`
	QueueSize int           `yaml:queue_size`
	Attempts  int           `yaml:attempts`
	Cooldown  float32       `yaml:cooldown`
}

var _ (sachet.Provider) = (*Modem)(nil)

type Modem struct {
	Gsm       *gsm.GSM
	Pdu       bool
	Attempts  int
	Cooldown  float32
	inQueue   chan queueCommand
	outQueue  chan queueResponse
	queueSize int
}

const (
	put uint = iota
	fetch
	check
)

type queueCommand struct {
	command uint
	message sachet.Message
}

type queueResponse struct {
	status  bool
	message sachet.Message
}

var modemInstance *Modem = nil

func SetupModem(config Config) error {
	modem := Modem{
		Attempts:  5,
		Cooldown:  1,
		queueSize: 5,
		inQueue:   make(chan queueCommand, 1),
		outQueue:  make(chan queueResponse, 1),
	}
	mode := "pdu"
	baudRate := 115200
	timeout := 5 * time.Second
	verbose := false
	hex := false

	if config.Mode != "" {
		mode = config.Mode
	}

	if mode == "pdu" {
		modem.Pdu = true
	}

	if config.BaudRate != 0 {
		baudRate = config.BaudRate
	}

	if config.Timeout != 0 {
		timeout = config.Timeout * time.Second
	}

	if config.QueueSize != 0 {
		modem.queueSize = config.QueueSize
	}

	if config.Attempts != 0 {
		modem.Attempts = config.Attempts
	}

	if config.Cooldown != 0 {
		modem.Cooldown = config.Cooldown
	}

	if config.Verbose {
		verbose = true
	}

	if config.Hex {
		hex = true
	}

	m, err := serial.New(serial.WithPort(config.Device), serial.WithBaud(baudRate))
	if err != nil {
		return err
	}

	var mio io.ReadWriter = m

	if hex {
		mio = trace.New(m, trace.WithReadFormat("r: %v"))
	} else if verbose {
		mio = trace.New(m)
	}

	gopts := []gsm.Option{}

	if !modem.Pdu {
		gopts = append(gopts, gsm.WithTextMode)
	}

	g := gsm.New(at.New(mio, at.WithTimeout(timeout)), gopts...)
	if err = g.Init(); err != nil {
		return err
	}

	modem.Gsm = g
	go modem.runFilter()
	go modem.runDispatcher()
	modemInstance = &modem
	return nil
}

func NewModem(config Config) *Modem {
	return modemInstance
}

func (m *Modem) Send(message sachet.Message) error {
	m.inQueue <- queueCommand{command: put, message: message}
	return nil
}

func (m *Modem) runFilter() {
	queue := make([]sachet.Message, 0)
	waiting := false

	for {
		cmd := <-m.inQueue

		if cmd.command == fetch {
			if len(queue) == 0 {
				waiting = true
				continue
			}

			m.outQueue <- queueResponse{status: true, message: queue[0]}
			queue = queue[1:]
		} else if cmd.command == check {
			m.outQueue <- queueResponse{status: len(queue) > 0}
		} else if cmd.command == put {
			if waiting {
				m.outQueue <- queueResponse{status: true, message: cmd.message}
				waiting = false
			} else {
				queue = append(queue, cmd.message)

				if len(queue) > m.queueSize {
					log.Printf("modem: dropping message from send queue")
					queue = queue[1:]
				}
			}
		}
	}
}

func (m *Modem) runDispatcher() {
	for {
		m.inQueue <- queueCommand{command: fetch}
		cmd := <-m.outQueue
		m.sendSmses(cmd.message)
	}
}

func (m *Modem) sendSmses(message sachet.Message) {
	smses := make(map[string]bool)

	for _, number := range message.To {
		smses[number] = false
	}

	for i := 1; i <= m.Attempts; i++ {
		for number := range smses {
			log.Printf("Sending SMS to %s (attempt %d/%d)", number, i, m.Attempts)

			err := m.sendOneSms(message, number)
			if err == nil {
				log.Printf("Sent SMS to %s (attempt %d/%d)", number, i, m.Attempts)
				delete(smses, number)
				continue
			}

			if i == m.Attempts {
				log.Printf("Unable to send SMS to: %s", number)
			} else {
				log.Printf("Failed to send SMS to %s: %w (attempt %d/%d)", number, err, i, m.Attempts)
			}
		}

		if len(smses) == 0 {
			return
		}

		time.Sleep(time.Duration(m.Cooldown) * time.Second)

		// If there is another queued message, give it precedence over
		// further attempts to send the current message, as we're falling
		// behind.
		m.inQueue <- queueCommand{command: check}
		if resp := <-m.outQueue; resp.status {
			log.Printf("Skipping further attempts, continue with another message")
			return
		}
	}

}

func (m *Modem) sendOneSms(message sachet.Message, number string) (err error) {
	if m.Pdu {
		_, err = m.Gsm.SendLongMessage(number, message.Text)
	} else {
		_, err = m.Gsm.SendShortMessage(number, message.Text)
	}

	return
}
