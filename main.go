package main

import (
	"fmt"
	"math"
)

func main() {
	var x = 1
	var x32 int32 = 1
	var x64 int64 = 1

	fmt.Println("x = ", x)
	fmt.Println("x32 = ", x32)
	fmt.Println("x64 = ", x64)

	fmt.Printf("FORMAT x = %T, x32= %T, x64= %T", x, x32, x64)
	fmt.Println()

	x = int(x32)
	fmt.Printf("%T , ", x)
	fmt.Println(x)

	y := 2
	fmt.Println(x + y)  //3
	fmt.Println(x % y)  //1
	fmt.Println(x == y) //false
	fmt.Println(x < y)  //true

	// *** FLOADING POINT NUMBERS
	pi := 3.141
	fmt.Println(pi)
	fmt.Printf("%T", pi)
	fmt.Println()

	var pi32 float32 = 3.14
	fmt.Println(pi32)
	fmt.Printf("%T", pi32)
	println()

	pi = float64(3.14)
	fmt.Println(pi)

	example := 1e10
	fmt.Println(example)

	// Operation +
	a := 1
	b := 3.231
	// It's invallid fmt.Println(a + b)
	//Valid :
	fmt.Println(float64(a) + b)

	//MATH LIBRARY
	fmt.Println("MATH LIBRARY")
	fmt.Println(math.Ceil(math.Pi))  //4
	fmt.Println(math.Floor(math.Pi)) //3
	fmt.Println(math.Min(1, 0))      //0
	fmt.Println(math.Max(1, 0))      //1
	fmt.Println(math.Abs(-1))        //1
	fmt.Println(math.Sqrt(100))      //10
	fmt.Println(math.Pow(2, 3))      //8

	// *** COMPLEX NUMBERS ***
	fmt.Println(complex(1, 2))
}
