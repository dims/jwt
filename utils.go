// +build !go1.14

package jwt

import "math/big"

func FillBytes(x *big.Int, buf []byte) []byte {
	return x.FillBytes(buf)
}
