package main 

import ("fmt" 
		"math"
		"reflect"
)
const s string = "Constant string"

// Go supports constants of char, string, bool and numeric values

// A const statement can appear anywhere a var statement can.

func main(){
	fmt.Println(s)

	const n = 9232382
	//fmt.Println(reflect.TypeOf(n))

	const d = 33193103138103109 / n
	fmt.Println(d)

	fmt.Println(reflect.TypeOf(d))
	fmt.Println(int64(d)) // A numeric constant has no type untils its given one by casting(similar to python)
	fmt.Println(math.Sin(n)) // the other way is using it in a context that requires one, 
	// Like a variable assignment or function call. For Ex: math.Sin takes a float64
	fmt.Println(reflect.TypeOf(n)) // gives int --  types wont change in this case
}