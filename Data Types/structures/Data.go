package main

import (
	"fmt"
)

// POINTERS

var i = 43

// The & operator generates a pointer to its operand.
var p = &i

// STRUCTS
// A collection of fields

type Vertex struct {
	X int
	Y int
	Z string
}

// accessing struct using dot notation
var v = Vertex{3, 9, "poor"}

// Struct Literals
var (
	v1 = Vertex{1, 2, "one"}
	v2 = Vertex{}
	v3 = Vertex{X: 1}
	v4 = &v1
)

// ARRAYS
//The type [n]T is an array of n values of type T.

// SLICES
// An array has a fixed size. A slice, on the other hand, is a dynamically-sized, flexible view into the elements of an array.
//In practice, slices are much more common than arrays.

var x = [4]int{33, 78, 23, 90} //[Array]

var s []int = x[0:2] //[Slice]

// A slice does not store any data, it just describes a section of an underlying array.

var names = [5]string{"Daniel", "Kevin", "Francis", "Davis", "Alexis"}
var a = names[0:2]
var b = names[2:3]

//slice literals - the same as an array literal

var q = []int{1, 2, 3, 4, 5, 6, 7, 8}

var y = []bool{true, false, false}

var w = []struct {
	i int
	t string
}{{22, "middle"}, {18, "Young"}}

//When slicing, you may omit the high or low bounds to use their defaults instead.

//The default is zero for the low bound and the length of the slice for the high bound.

//Creating a slice with make

var o = make([]int, 0, 5)

// slices in slices
// Create a tic-tac-toe board.
var board = [][]string{
	[]string{"_", "_", "_"},
	[]string{"_", "_", "_"},
	[]string{"_", "_", "_"},
}

//Appending to a slice

var h []int

// RANGE
// The range form of the for loop iterates over a slice or map.
// It returns the first element as index and the copy of the element (just as map in ruby and javascript)
var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	fmt.Println("Everything works well")
	fmt.Println(p)
	//The * operator denotes the pointer's underlying value.
	fmt.Println(*p)

	fmt.Println(Vertex{1, 2, "hello"})
	v.X = 4
	v.Y = 10
	v.Z = "life"
	fmt.Println(v)

	// Pointers to structs
	p := &v
	p.X = 25

	fmt.Println(v)
	fmt.Println(v1, v2, v3, v4)

	var j [2]string
	j[0] = "Hello"
	j[1] = "World"

	var k = [6]int{1, 2, 3, 4, 5, 6}

	fmt.Println(j)
	fmt.Println(k)
	fmt.Println(s)
	fmt.Println(names)
	fmt.Println(a)
	fmt.Println(b)

	a[0] = "Mary"
	fmt.Println(names)
	fmt.Println(a)
	fmt.Println(q, y, w)
	fmt.Println(o)
	//The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	fmt.Println(board)
	// Appending to our empty slice
	h = append(h, 10, 20, 79, 34, 54, 34, 65)

	fmt.Println(h)

	// using range
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}
