package main

import "fmt"

//can declare global
var name = "Nate"

func main() {
	//using vars
	//type is inferred
	var name = "nate"
	fmt.Println(name)

	//again, type is inferred
	var age = 18
	fmt.Println(name, age)

	//notation can only be used within function body
	//Notice, can only declare initially, not redeclare
	//Needs to be a new variable
	newName := "Nathan"

	//to find the type, use Printf and a formatter
	fmt.Printf("%T\n%T\n", newName, age)

	size := 3.87
	fmt.Printf("%T\n", size)

	//can declare multiple variables on same line
	customerName, email := "nathan", "nate@email.com"
	fmt.Println(customerName, email)

}
