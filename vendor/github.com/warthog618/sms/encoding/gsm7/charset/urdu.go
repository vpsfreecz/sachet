// SPDX-License-Identifier: MIT
//
// Copyright © 2018 Kent Gibson <warthog618@gmail.com>.

package charset

var (
	urduDecoder = Decoder{
		// awkward layout as most editors have problems with these RTL characters.
		0x00: 'ا',
		0x01: 'آ',
		0x02: 'ب',
		0x03: 'ٻ',
		0x04: 'ڀ',
		0x05: 'پ',
		0x06: 'ڦ',
		0x07: 'ت',
		0x08: 'ۂ',
		0x09: 'ٿ',
		0x0a: '\n',
		0x0b: 'ٹ',
		0x0c: 'ٽ',
		0x0d: '\r',
		0x0e: 'ٺ',
		0x0f: 'ټ',
		0x10: 'ث',
		0x11: 'ج',
		0x12: 'ځ',
		0x13: 'ڄ',
		0x14: 'ڃ',
		0x15: 'څ',
		0x16: 'چ',
		0x17: 'ڇ',
		0x18: 'ح',
		0x19: 'خ',
		0x1a: 'د',
		0x1b: 0x1b,
		0x1c: 'ڌ',
		0x1d: 'ڈ',
		0x1e: 'ډ',
		0x1f: 'ڊ',
		0x20: 0x20,
		0x21: '!',
		0x22: 'ڏ',
		0x23: 'ڍ',
		0x24: 'ذ',
		0x25: 'ر',
		0x26: 'ڑ',
		0x27: 'ړ',
		0x28: ')',
		0x29: '(',
		0x2a: 'ڙ',
		0x2b: 'ز',
		0x2c: ',',
		0x2d: 'ږ',
		0x2e: '.',
		0x2f: 'ژ',
		0x30: '0',
		0x31: '1',
		0x32: '2',
		0x33: '3',
		0x34: '4',
		0x35: '5',
		0x36: '6',
		0x37: '7',
		0x38: '8',
		0x39: '9',
		0x3a: ':',
		0x3b: ';',
		0x3c: 'ښ',
		0x3d: 'س',
		0x3e: 'ش',
		0x3f: '?',
		0x40: 'ص',
		0x41: 'ض',
		0x42: 'ط',
		0x43: 'ظ',
		0x44: 'ع',
		0x45: 'ف',
		0x46: 'ق',
		0x47: 'ک',
		0x48: 'ڪ',
		0x49: 'ګ',
		0x4a: 'گ',
		0x4b: 'ڳ',
		0x4c: 'ڱ',
		0x4d: 'ل',
		0x4e: 'م',
		0x4f: 'ن',
		0x50: 'ں',
		0x51: 'ڻ',
		0x52: 'ڼ',
		0x53: 'و',
		0x54: 'ۄ',
		0x55: 'ە',
		0x56: 'ہ',
		0x57: 'ھ',
		0x58: 'ء',
		0x59: 'ی',
		0x5a: 'ې',
		0x5b: 'ے',
		0x5c: '\u064D',
		0x5d: '\u0650',
		0x5e: '\u064f',
		0x5f: '\u0657',
		0x60: '\u0654',
		0x61: 'a',
		0x62: 'b',
		0x63: 'c',
		0x64: 'd',
		0x65: 'e',
		0x66: 'f',
		0x67: 'g',
		0x68: 'h',
		0x69: 'i',
		0x6a: 'j',
		0x6b: 'k',
		0x6c: 'l',
		0x6d: 'm',
		0x6e: 'n',
		0x6f: 'o',
		0x70: 'p',
		0x71: 'q',
		0x72: 'r',
		0x73: 's',
		0x74: 't',
		0x75: 'u',
		0x76: 'v',
		0x77: 'w',
		0x78: 'x',
		0x79: 'y',
		0x7a: 'z',
		0x7b: '\u0655',
		0x7c: '\u0651',
		0x7d: '\u0653',
		0x7e: '\u0656',
		0x7f: '\u0670',
	}
	urduExtDecoder = Decoder{
		0x00: '@',
		0x01: '£',
		0x02: '$',
		0x03: '¥',
		0x04: '¿',
		0x05: '"',
		0x06: '¤',
		0x07: '%',
		0x08: '&',
		0x09: '\'',
		0x0a: '\f',
		0x0b: '*',
		0x0c: '+',
		0x0d: '\r',
		0x0e: '-',
		0x0f: '/',
		0x10: '<',
		0x11: '=',
		0x12: '>',
		0x13: '¡',
		0x14: '^',
		0x15: '¡',
		0x16: '_',
		0x17: '#',
		0x18: '*',
		0x19: '؀',
		0x1a: '؁',
		0x1b: 0x1b,
		0x1c: '۰',
		0x1d: '۱',
		0x1e: '۲',
		0x1f: '۳',
		0x20: '۴',
		0x21: '۵',
		0x22: '۶',
		0x23: '۷',
		0x24: '۸',
		0x25: '۹',
		0x26: '،',
		0x27: '؍',
		0x28: '{',
		0x29: '}',
		0x2a: '؎',
		0x2b: '؏',
		0x2c: '\u0610',
		0x2d: '\u0611',
		0x2e: '\u0612',
		0x2f: '\\',
		0x30: '\u0613',
		0x31: '\u0614',
		0x32: '\u061b',
		0x33: '\u061f',
		0x34: '\u0640',
		0x35: '\u0652',
		0x36: '\u0658',
		0x37: '٫',
		0x38: '٬',
		0x39: 'ٲ',
		0x3a: 'ٳ',
		0x3b: 'ۍ',
		0x3c: '[',
		0x3d: '~',
		0x3e: ']',
		0x3f: '۔',
		0x40: '|',
		0x41: 'A',
		0x42: 'B',
		0x43: 'C',
		0x44: 'D',
		0x45: 'E',
		0x46: 'F',
		0x47: 'G',
		0x48: 'H',
		0x49: 'I',
		0x4a: 'J',
		0x4b: 'K',
		0x4c: 'L',
		0x4d: 'M',
		0x4e: 'N',
		0x4f: 'O',
		0x50: 'P',
		0x51: 'Q',
		0x52: 'R',
		0x53: 'S',
		0x54: 'T',
		0x55: 'U',
		0x56: 'V',
		0x57: 'W',
		0x58: 'X',
		0x59: 'Y',
		0x5a: 'Z',
		0x65: '€',
	}
	urduEncoder    Encoder
	urduExtEncoder Encoder
)

func generateUrduEncoder() Encoder {
	return generateEncoder(urduDecoder)
}

func generateUrduExtEncoder() Encoder {
	return generateEncoder(urduExtDecoder)
}
