package main

import "fmt"

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

}
