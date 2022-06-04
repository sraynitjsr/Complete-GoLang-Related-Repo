package main

import (
	"reflect"
	"strconv"
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Enter Integer Input")
	reader1  := bufio.NewReader(os.Stdin)
	input1,_ := reader1.ReadString('\n')
	myInt,_ := strconv.Atoi(input1)
	fmt.Println("Entered Integrer is ",input1+" of type",reflect.TypeOf(myInt))

	fmt.Println("Enter Float64 Input")
	reader2  := bufio.NewReader(os.Stdin)
	input2,_ := reader2.ReadString('\n')
	myFloat64,_ := strconv.ParseFloat(input2, 64)
	fmt.Println("Entered Floating Number is "+input2," of type",reflect.TypeOf(myFloat64))

	fmt.Println("Enter Third Input")
	reader3  := bufio.NewReader(os.Stdin)
	input3,_ := reader3.ReadString('\n')
	fmt.Println("Entered String is ",input3+" of type",reflect.TypeOf(input3))
}
