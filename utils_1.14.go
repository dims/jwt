// +build go1.14

package jwt

import (
	"math/big"
	"math/bits"
)

const (
	_S = _W / 8        // word size in bytes
	_W = bits.UintSize // word size in bits
)

func FillBytes(x *big.Int, buf []byte) []byte {
	// Clear whole buffer. (This gets optimized into a memclr.)
	for i := range buf {
		buf[i] = 0
	}
	toBytes(x.Bits(), buf)
	return buf
}

// bytes writes the value of z into buf using big-endian encoding.
// The value of z is encoded in the slice buf[i:]. If the value of z
// cannot be represented in buf, bytes panics. The number i of unused
// bytes at the beginning of buf is returned as result.
func toBytes(z []big.Word, buf []byte) (i int) {
	i = len(buf)
	for _, d := range z {
		for j := 0; j < _S; j++ {
			i--
			if i >= 0 {
				buf[i] = byte(d)
			} else if byte(d) != 0 {
				panic("math/big: buffer too small to fit value")
			}
			d >>= 8
		}
	}

	if i < 0 {
		i = 0
	}
	for i < len(buf) && buf[i] == 0 {
		i++
	}
	return
}
