package main

import (
	"fmt"
	"math"
)

const precision = 0.000001

// a2 = 0.5 * (a1 + x/a1)
// a2 = (a1 + x/a1) / 2

// Sqrt - Getting the square root using Newton's method
func Sqrt(x float64) float64 {
	var (
		a1, a2 float64
	)

	a1 = x / 2
	for {
		a2 = (a1 + x/a1) / 2
		if (a1 - a2) < precision {
			break
		}
		a1 = a2
	}
	return a2
}

func main() {
	fmt.Println(Sqrt(12345678))
	fmt.Println(math.Sqrt(12345678))
}
