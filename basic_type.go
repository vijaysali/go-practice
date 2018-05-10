package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Hello Welcome to go-world")

	// Variable declarations
	var a int = 430
	fmt.Println("Value for a is:", a, "And the type is: ", reflect.TypeOf(a))

	b := 30
	fmt.Println("Value for b is:", b)

	// other options are int8, int16, int32 and int64
	// similary for unsignedint
	var c int32
	c = 19999
	fmt.Println("Value for c is ", c)

	var x string = "Hello"
	fmt.Println("X is ", x)

	f := 3.14
	fmt.Println("PI value is ", f)
}
