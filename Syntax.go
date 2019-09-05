/* Language syntax basics for Golang */

package main

import "fmt"

func main() {
	//For Loops
	//Classic var; condition; increment Loop
	for i := 1; i <= 3; i++ {
		fmt.Println(i)
	}
	//While True Endless Loop
	for {
		fmt.Println("...")
		break
	}
	//Loop with Continue
	for i := 0; i <= 5; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}

	//If, else if, else
	num := 9
	if num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has one digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}

	//Constants
	const s string = "constant"
	fmt.Println(s)
	const n = 5000
	fmt.Println(n)

	//Simple Variables
	var a = "initial"
	fmt.Println(a)
	var b, c int = 1, 2
	fmt.Println(b, c)
	var d = true
	fmt.Println(d)
	var e int
	fmt.Println(e)
	//Combined Assignment
	f := "apple"
	fmt.Println(f)

	//Boolean Logic
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)

	//String Printing
	fmt.Println("Hello World")
	fmt.Println("go" + "lang")
}
