package main 

import "fmt"


// Go has various value types including strings, integers, floats, booleans, etc. Here are a few basic examples.

func main(){
	// Strings, which can be added together with +. 
	fmt.Println("go" + "lang")
	fmt.Println("1+1.5 = ", 1+1.5)
	fmt.Println("7.0/3 =", 7.0/3)
	fmt.Println(true && false)
	fmt.Println(!true)
}