package main

import "fmt"

func main() {
	const a float64 = 5.1 //typed constant

	const b = 6.7 // untyped constant

	const c float64 = a * b

	const str = "Hello " + "GO!"
	const d = 5 > 10
	fmt.Println(d)

	// const x int = 5
	// const y float64 = 2.2 * x
	// Error occurs
	const x = 5
	const y = 2.2 * x

	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", y)

	var i int = x     // x changes to int
	var j float64 = x // float64(x)
	var p byte = x

	fmt.Printf("%v,%v,%v\n", i, j, p)
}
