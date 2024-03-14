package main

import (
	"fmt"

	"github.com/arrudadev/hello-go/internal/math"
)

func main() {
	sum := math.Sum(10, 5)
	minus := math.Minus(10, 5)
	multiply := math.Multiply(10, 5)
	divide := math.Divide(10, 5)

	fmt.Println("Sum: ", sum)
	fmt.Println("Minus: ", minus)
	fmt.Println("Multiply: ", multiply)
	fmt.Println("Divide: ", divide)
}
