// PACKAGES
package main

// IMPORTS
import (
	"fmt"
	"math"
	"math/cmplx"
)

// BASIC FUNCTIONAL CALLS
func add(x int, y int) int {
	return x + y
}
func swap(x string, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

// VARIABLES

var python, java, c bool

// Variables with intializers
var age, height, weight = false, true, "No!"
var j, k = 2, 3

// Declaring constants
const PI = 3.142

// TYPES
// boolean
var toBe bool = false

// string
var hello string = "Hello"

// int
var maxInt uint64 = 1<<64 - 1

// Complex numbers
var z complex128 = cmplx.Sqrt(5 + 12i)

// Zero values (Variables declared with no value)
var i int            //Gives zero for numeric value
var isLoggedIn bool  // false for boolean values
var greetings string // "" empty quotes for strings

func main() {
	//Type conversion
	var n, p int = 3, 4
	var h float64 = math.Sqrt(float64(n*n + p*p))
	var z uint = uint(h)
	//Type inference
	var m = "42"
	var o = 42

	fmt.Printf("Type: %T Value: %v\n", m, m)
	fmt.Printf("Type: %T Value: %v\n", o, o)

	var i string
	// Short variable declarations (only used inside functions)
	length := 11
	fmt.Println(z)
	fmt.Println(split(17))
	fmt.Println(i, python, java, c)
	fmt.Println(age, height, weight, j, k)
	fmt.Println(length)
	fmt.Printf("Type: %T Value: %v\n", toBe, toBe)
	fmt.Printf("Type: %T Value: %v\n", hello, hello)
	fmt.Printf("Type: %T Value: %v\n", maxInt, maxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
}
