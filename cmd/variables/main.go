package main

import "fmt"

var a int = 10
var b int = 10

func main() {
	sum := a + b
	text := "The sum of a and b is: "

	fmt.Println(text, sum)
	fmt.Printf("Type of a %T \n", a)
	fmt.Printf("Type of b %T \n", b)
	fmt.Printf("Type of text %T \n", text)

	biggerText := `
		this is a bigger text,
		with multiple lines
	`

	fmt.Println(biggerText)
	fmt.Printf("Type of biggerText %T \n", biggerText)

	boolVar := true
	fmt.Printf("Type of boolVar %T \n", boolVar)
}
