package main

import "fmt"

func main() {
	var i int
	i = 45
	fmt.Println(i)

	var f float32 = 3.14
	fmt.Println(f)

	firstName := "Bob"
	fmt.Println(firstName)

	b := true
	fmt.Println(b)

	c := complex(3, 4)
	fmt.Println(c)

	r, im := real(c), imag(c)
	fmt.Println(r, im)
}
