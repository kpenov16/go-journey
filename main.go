package main

import "fmt"

//expresion families
//const family 1
const (
	loFirst  = iota + 6
	loSecond = 2 << iota
	loThird
)

//const family 2
const (
	loFlamily2 = iota
	loFlamily2Second
)

func main() {
	//variables
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
	//end variables

	//pointers
	var lastName *string = new(string)
	*lastName = "BobLast"
	fmt.Println(*lastName)

	name := "BobName"
	fmt.Println(name)

	ptrToName := &name
	fmt.Println(ptrToName, *ptrToName)

	name = "ChangedBobName"
	fmt.Println(ptrToName, *ptrToName)
	//end pointers

	//constants
	const pi = 2.1415
	fmt.Println(pi)

	const n int = 2
	fmt.Println(n + 2)
	//code
	fmt.Println(float32(n) + 2.2)

	//constants iota
	fmt.Println(loFirst, loSecond, loThird, loFlamily2, loFlamily2Second)

}
