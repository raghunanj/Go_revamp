// In go, vairables are explicitly declared and used by the compiler 
// for type correctness of function calls. 
package main 

import ("fmt"
		"reflect"
)
		// The Go reflection package has methods for inspecting the type of variables. 


func main(){
	var a = "initial"
	fmt.Println(a)

	var b,c int = 1,2 
	fmt.Println(b,c)

	// var can declare 1 or more variables and multiple variables at once.
	//like python, go will infer the type of initialuzed variables. should mention var at the beginning thats must. 
	// variables declared without a corresponding initialization are zero-valued. 
	// For ex: var e int stores 0. var e  

	/*
	basic types in golang.

	bool

	string

	int  int8  int16  int32  int64
	uint uint8 uint16 uint32 uint64 uintptr

	byte // alias for uint8

	rune // alias for int32
	     // represents a Unicode code point

	float32 float64

	complex64 complex128
	
	*/
	var d = true
	fmt.Println(d)

	test1 := "string"
    test2 := 10
    test3 := 1.2
    test4 := false

    fmt.Println(reflect.TypeOf(test1))
    fmt.Println(reflect.TypeOf(test2))
    fmt.Println(reflect.TypeOf(test3))
    fmt.Println(reflect.TypeOf(test4))

    // Simple nice thread about typeof vars in go:-   https://stackoverflow.com/questions/20170275/how-to-find-a-type-of-an-object-in-go
    // Some stuff about packages and set data structure. https://code.tutsplus.com/tutorials/12-indispensable-go-packages-and-libraries--cms-29008



}