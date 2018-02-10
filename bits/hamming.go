package astibits

import "fmt"

var hamming84tab = [256]uint8{
	0x01, 0xff, 0xff, 0x08, 0xff, 0x0c, 0x04, 0xff, 0xff, 0x08, 0x08, 0x08, 0x06, 0xff, 0xff, 0x08,
	0xff, 0x0a, 0x02, 0xff, 0x06, 0xff, 0xff, 0x0f, 0x06, 0xff, 0xff, 0x08, 0x06, 0x06, 0x06, 0xff,
	0xff, 0x0a, 0x04, 0xff, 0x04, 0xff, 0x04, 0x04, 0x00, 0xff, 0xff, 0x08, 0xff, 0x0d, 0x04, 0xff,
	0x0a, 0x0a, 0xff, 0x0a, 0xff, 0x0a, 0x04, 0xff, 0xff, 0x0a, 0x03, 0xff, 0x06, 0xff, 0xff, 0x0e,
	0x01, 0x01, 0x01, 0xff, 0x01, 0xff, 0xff, 0x0f, 0x01, 0xff, 0xff, 0x08, 0xff, 0x0d, 0x05, 0xff,
	0x01, 0xff, 0xff, 0x0f, 0xff, 0x0f, 0x0f, 0x0f, 0xff, 0x0b, 0x03, 0xff, 0x06, 0xff, 0xff, 0x0f,
	0x01, 0xff, 0xff, 0x09, 0xff, 0x0d, 0x04, 0xff, 0xff, 0x0d, 0x03, 0xff, 0x0d, 0x0d, 0xff, 0x0d,
	0xff, 0x0a, 0x03, 0xff, 0x07, 0xff, 0xff, 0x0f, 0x03, 0xff, 0x03, 0x03, 0xff, 0x0d, 0x03, 0xff,
	0xff, 0x0c, 0x02, 0xff, 0x0c, 0x0c, 0xff, 0x0c, 0x00, 0xff, 0xff, 0x08, 0xff, 0x0c, 0x05, 0xff,
	0x02, 0xff, 0x02, 0x02, 0xff, 0x0c, 0x02, 0xff, 0xff, 0x0b, 0x02, 0xff, 0x06, 0xff, 0xff, 0x0e,
	0x00, 0xff, 0xff, 0x09, 0xff, 0x0c, 0x04, 0xff, 0x00, 0x00, 0x00, 0xff, 0x00, 0xff, 0xff, 0x0e,
	0xff, 0x0a, 0x02, 0xff, 0x07, 0xff, 0xff, 0x0e, 0x00, 0xff, 0xff, 0x0e, 0xff, 0x0e, 0x0e, 0x0e,
	0x01, 0xff, 0xff, 0x09, 0xff, 0x0c, 0x05, 0xff, 0xff, 0x0b, 0x05, 0xff, 0x05, 0xff, 0x05, 0x05,
	0xff, 0x0b, 0x02, 0xff, 0x07, 0xff, 0xff, 0x0f, 0x0b, 0x0b, 0xff, 0x0b, 0xff, 0x0b, 0x05, 0xff,
	0xff, 0x09, 0x09, 0x09, 0x07, 0xff, 0xff, 0x09, 0x00, 0xff, 0xff, 0x09, 0xff, 0x0d, 0x05, 0xff,
	0x07, 0xff, 0xff, 0x09, 0x07, 0x07, 0x07, 0xff, 0xff, 0x0b, 0x03, 0xff, 0x07, 0xff, 0xff, 0x0e,
}

func Hamming84Decode(i uint8) (o uint8, err error) {
	o = hamming84tab[i]
	if o == 0xff {
		err = fmt.Errorf("astibits: hamming 8/4 decode of %.8b failed", i)
		return
	}
	return
}