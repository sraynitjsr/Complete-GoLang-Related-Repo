package main

import (
	"fmt"
)

func Variables() {
	//Variable Declaration With Initial Value
	var employee1 = "Subhradeep"
	var employee2 = "Ray"
	employee3 := "Subhradeep Ray"
	fmt.Println("Employee1 = " + employee1 + " Employee2 = " + employee2 + " Employee3 = " + employee3)

	//Variable Declaration Without Initial Value
	var a string
	var b int
	var c bool
	fmt.Println("A = ", a)
	fmt.Println("B = ", b)
	fmt.Println("C = ", c)

	//Value Assignment After Declaration
	var temp = 10
	temp = 15 / 3
	fmt.Println("Modified Temp = ", temp)

	/*
		Difference Between var and :=
		1. var Can be used inside and outside of functions  BUT  := Can only be used inside functions
		2. var -> Variable declaration and value assignment can be done separately BUT := -> Variable declaration and value assignment cannot be done separately (must be done in the same line)
	*/

	//Go Multiple Variable Declaration
	var x, y, z int = 10, 20, 30
	fmt.Printf("\nX = %d Y = %d Z = %d", x, y, z)

	var p, q = 100, "My String"
	u, v := "UU", true
	fmt.Printf("\nP = %d Q = %s U = %s V = %v", p, q, u, v)

	//Go Variable Declaration in a Block
	var (
		varA float64 = 10.50
		varB string  = "Simple String"
	)
	fmt.Println("\n", "varA = ", varA)
	fmt.Println("\n", "varA = ", varB)
}
