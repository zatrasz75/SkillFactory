package main

import (
	fmt "fmt"
	"helloapp/calc"
	"helloapp/maths"
)

func main() {
	//от прошлого задания
	fmt.Println(maths.E)
	fmt.Println(maths.B)

	//------------
	x := calc.ReadIntager1()
	s := calc.ReadStr()
	y := calc.ReadIntager2()

	switch s {
	case "+":
		fmt.Println(x + y)
	case "-":
		fmt.Println(x - y)
	case "*":
		fmt.Println(x * y)
	case "/":
		if y != 0 {
			fmt.Println(x / y)
		} else {
			fmt.Printf("Нельзя делить число на %g", y)
		}

	}
}
