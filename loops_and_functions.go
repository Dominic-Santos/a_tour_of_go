package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	z := x/2
	for i := 1; i < 10; i++ {
	    fmt.Println(x, i, z)
		z -= (z*z - x) / (2*z)
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(4))
	fmt.Println(Sqrt(8))
	fmt.Println(Sqrt(128))
}
