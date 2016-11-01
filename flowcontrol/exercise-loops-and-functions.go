package main

import "fmt"

func Sqrt(x float64) float64 {
	z := 1.0

	for step := 1; step <= 10; step++ {
		z = z - (z*z-x)/(2*z)
		fmt.Printf("step:%d z=%g\n", step, z)
	}

	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
