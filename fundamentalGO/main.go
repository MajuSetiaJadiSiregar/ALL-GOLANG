package main

import (
	"AG/fundamentalGO/calculation"
	"fmt"
)

func main() {
	fmt.Println("Hallo, Belajar Golang")

	sentence := testPackage()

	result := calculation.Add(8, 9)

	fmt.Print(sentence)

	fmt.Println(result)
}
