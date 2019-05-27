package main

import "fmt"

func main() {
	//using vars
	//type is inferred
	var name = "nate"
	fmt.Println(name)

	//again, type is inferred
	var age = 18
	fmt.Println(name, age)

	//to find the type, use a formatter
	fmt.Printf("%T\n", name)

}
