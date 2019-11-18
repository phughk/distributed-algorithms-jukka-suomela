package one

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestP3CBit(t *testing.T) {
	a := assert.New(t)

	// Taken from Table 1.2
	value, err := P3CBit(0b01111011, 0b00101111)
	a.NoError(err)
	a.Equal(value, uint8(4))

	value, err = P3CBit(0b00101111, 0b01101011)
	a.NoError(err)
	a.Equal(value, uint8(5))

	value, err = P3CBit(0b01111011, 0b00101111)
	a.NoError(err)
	a.Equal(value, uint8(4))

	value, err = P3CBit(0b00101111, 0b01101111)
	a.NoError(err)
	a.Equal(value, uint8(12))
}
