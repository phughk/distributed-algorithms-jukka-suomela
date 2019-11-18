package one

import (
	"fmt"
	"math"
)

func P3CBit(input1, input2 uint) int16 {
	// Find index of first difference
	var index int16 = -1
	for i := 0; i < 256; i++ {
		mask := uint(math.Pow(2, float64(i)))
		different := (input1 & mask) & (input2 & mask)
		if different > 0 {
			index = int16(different)
			break
		}
	}
	return index
}

func hashCollisionTest() {
	fmt.Println("Hash Collision Test on P3CBit")
}

func Main() {
	fmt.Println("One")
	hashCollisionTest()
}
