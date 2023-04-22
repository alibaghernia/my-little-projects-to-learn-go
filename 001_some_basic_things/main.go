package main

import "fmt"

func main() {
	// Comment:
	//	this is a single line comment

	/*
		multi line comment
		second line of multi line comment
	*/

	// ---------------------------------------------
	// Variable
	var name string = "Ali"
	name2 := "Kourosh"
	fmt.Println("name is:", name, "name2 is:", name2)

	var number int
	number = 5 + 1
	fmt.Println("number is:", number)

	// Multiple variable declaration
	var a, b, c, d int = 1, 2, 3, 4
	fmt.Println("a, b, c, d contains:", a, b, c, d)

	var (
		v1 float32
		v2 bool   = false
		v3 string = "a text..."
	)
	v2 = true
	v1 = 5.5
	fmt.Printf("v1 is: %g,"+
		" v2 is %t,"+
		" v3 is %s \n",
		v1, v2, v3)

}
