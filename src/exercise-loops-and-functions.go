package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	prev := z
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
		if z == prev {
			break
		}
		prev = z
		fmt.Println(i, ":", z)
	}
	return z
}

func main() {
	value := 2.0
	fmt.Println(Sqrt(value))
	fmt.Println(math.Sqrt(value))
}
