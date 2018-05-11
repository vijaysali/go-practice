package main

import (
	"fmt"
)

func main() {
	var array [10]int

	fmt.Println("Size:", len(array))

	array[3] = 3
	// index and value , each_with_index
	for index, value := range array {
		fmt.Println(index, value)
	}

	for _, value := range array {
		if value != 0 {
			fmt.Println(value)
		}
	}

	var q [3]int = [3]int{1, 2, 3}
	fmt.Println(q[1])

	// ellipsis(...), Size is determined by initializers

	array1 := [...]int{10, 20, 30, 40, 50}
	fmt.Println(array1)

	array2 := [...]int{10: 0}
	fmt.Println("1.upto(10).to_a: ", array2)
}
