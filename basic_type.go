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

	f := 3.14
	fmt.Println("PI value is ", f)

	// dealing with const
	const telno = 9742506107
	fmt.Println("Constant telephone:", telno)

	// Complex numbers
	var complex1 complex128 = complex(1, 2)
	var complex2 complex128 = complex(2, 4)
	fmt.Println("Complex number:", complex1*complex2)
	fmt.Println("Real of complex number:", real(complex1*complex2))
	fmt.Println("Imaginary of Complex number:", imag(complex1*complex2))

	//Boolean
	var bool bool = true
	fmt.Println("Boolean True", bool)
	bool = false
	fmt.Println("Boolean False", bool)
	fmt.Println("Compare a and b", a == b)

	//Strings
	var x string = "Hello"
	fmt.Println("X is ", x)

	string1 := "World"
	fmt.Println("string1 is ", string1)
	fmt.Println("Length of string1", len(string1))

	var (
		name, address string
		age           int
	)
	name = "vijay"
	address = "Bangalore"
	age = 29
	fmt.Println(name, address, age)
}
