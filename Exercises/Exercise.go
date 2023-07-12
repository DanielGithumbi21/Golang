// As a way to play with functions and loops, let's implement a square root function:
//given a number x, we want to find the number z for which z² is most nearly x.
// Computers typically compute the square root of x using a loop.
//Starting with some guess z, we can adjust z based on how close z² is to x, producing a better guess:
// z -= (z*z - x) / (2*z)
// Repeating this adjustment makes the guess better and better until we reach an answer
//that is as close to the actual square root as can be.

// Implement this in the func Sqrt provided. A decent starting guess for z is 1, no matter what the input.
//To begin with, repeat the calculation 10 times and print each z along the way.
//See how close you get to the answer for various values of x (1, 2, 3, ...) and how quickly the guess improves.

// SOLUTION
package main

import (
	"fmt"

	"golang.org/x/tour/pic"
)

func sqrt(x float64) float64 {
	z := 1.0
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
		fmt.Println(z)
	}
	return z
}

func main() {
	fmt.Println(sqrt(2))
	pic.Show(Pic)

}

// SLICES EXERCISE

//Implement Pic.
//It should return a slice of length dy, each element of which is a slice of dx 8-bit unsigned integers.
//When you run the program, it will display your picture, interpreting the integers as grayscale (well, bluescale) values.
//The choice of image is up to you. Interesting functions include (x+y)/2, x*y, and x^y.
//(You need to use a loop to allocate each []uint8 inside the [][]uint8.)
//(Use uint8(intValue) to convert between types.)

func Pic(dx, dy int) [][]uint8 {
	picture := make([][]uint8, dy)

	for y := 0; y < dy; y++ {
		picture[y] = make([]uint8, dx)
		for x := 0; x < dx; x++ {
			// Modify this line to choose the desired image function
			picture[y][x] = uint8((x + y) / 2)
			// picture[y][x] = uint8(x * y)
			// picture[y][x] = uint8(x ^ y)
		}
	}

	return picture
}
