package one

import (
	"errors"
	"math"
)

func P3CBit(input1, input2 uint8) (uint8, error) {
	// Find index of first difference
	var index int16 = -1
	for i := 0; i < 256; i++ {
		mask := uint8(math.Pow(2, float64(i)))
		different := (input1 & mask) & (input2 & mask)
		if different > 0 {
			index = int16(different)
			break
		}
	}
	if index < 0 {
		return 0, errors.New("The numbers are not unique")
	}
	var value int16 = int16(input1) & int16(math.Pow(2, float64(index)))
	value = value >> index // So we get single bit
	index += 1
	hash := (2 * index) + value
	if hash < 0 {
		return 0, errors.New("P3CBit value was less than zero")
	}
	return uint8(hash), nil
}
