package modem

import (
	"log"
	"io"
	"time"

	"github.com/messagebird/sachet"
	"github.com/warthog618/modem/at"
	"github.com/warthog618/modem/gsm"
	"github.com/warthog618/modem/serial"
	"github.com/warthog618/modem/trace"
)

type Config struct {
	Device string `yaml:device`
	Mode string `yaml:mode`
	BaudRate int `yaml:baud_rate`
	Timeout time.Duration `yaml:timeout`
	Verbose bool `yaml:verbose`
	Hex bool `yaml:hex`
	Attempts int `yaml:attempts`
	Cooldown float32 `yaml:cooldown`
}

var _ (sachet.Provider) = (*Modem)(nil)

type Modem struct {
	Gsm *gsm.GSM
	Pdu bool
	Attempts int
	Cooldown float32
	channel chan sachet.Message
}

var modemInstance *Modem = nil

func SetupModem(config Config) error {
	modem := Modem{Attempts: 5, Cooldown: 1, channel: make(chan sachet.Message, 32)}
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
	go modem.runDispatcher()
	modemInstance = &modem
	return nil
}

func NewModem(config Config) *Modem {
	return modemInstance
}

func (m *Modem) Send(message sachet.Message) error {
	go m.sendInBackground(message)
	return nil
}

func (m *Modem) runDispatcher() {
	for {
		message := <-m.channel
		m.sendSmses(message)
	}
}

func (m *Modem) sendInBackground(message sachet.Message) {
	m.channel <- message
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
