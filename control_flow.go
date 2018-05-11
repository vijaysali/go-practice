package main

import (
	"fmt"
)

func main() {
	answer := 10

	// If block
	if answer == 10 {
		fmt.Println("Answer is correct")
	}

	// if else block
	if answer == 9 {
		fmt.Println("If block")
	} else {
		fmt.Println("Else block")
	}

	// switch statement
	switch answer {
	case 10:
		fmt.Println("Block 10 executed")
	case 9:
		fmt.Println("Block 9 executed")
	default:
		fmt.Println("Default block")
	}

	// loop structure -- do, while, for all one :)

	j := 10
	// for loop
	fmt.Println("Printing j")
	for j >= 0 {
		fmt.Println(j)
		j--
	}

	fmt.Println("Printing i")
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}
