package main

import (
	"errors"
	"fmt"
	"log"
)

func minus(x int, y int) (int, error) {
	result := x - y

	if result > 0 {
		return result, nil
	} else {
		return result, errors.New("the result is less than zero")
	}
}

func main() {
	sum, err := minus(5, 10)

	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("Sum: ", sum)
}
