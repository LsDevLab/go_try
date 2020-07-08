package main

import (
	"fmt"
	"math"
)

var i, j int
var s = 3
var d = "ciao"

const c = 3.14

func addAndSub(x, y int) (a, b int) {
	a = x + y
	b = x - y
	return
}

func swap(x, y string) (string, string, string) {
	valore := "VAL" //only in a func
	return y, x, valore
}

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func main() {
	fmt.Printf("This is pi %f\n", math.Pi)
	fmt.Println(addAndSub(4, 6))
	fmt.Println(swap("prima", "seconda"))

	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	sum = 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}
