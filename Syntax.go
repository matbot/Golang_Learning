/* Language syntax basics for Golang
Learning directly from gobyexample.com exercises
*/

package main

import (
	"fmt"
	"time"
)

func main() {
	// Slices

	// Arrays
	var emptyArray [5]int
	fmt.Println("emptyLength:", len(emptyArray))
	fmt.Println("emptyContents:", emptyArray)
	emptyArray[4] = 100
	fmt.Println("updatedSet:", emptyArray)

	declarativeArray := [5]int{1, 2, 3, 4, 5}
	fmt.Println("declarativeArray:", declarativeArray)

	var twoDimensionalArray [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoDimensionalArray[i][j] = i + j
		}
	}
	fmt.Println("2dArray:", twoDimensionalArray)

	//Switch Statements
	switchInteger := 2
	fmt.Print("Write ", switchInteger, " as ")
	switch switchInteger {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend!")
	default:
		fmt.Println("It's a weekday :(")
	}

	nowTime := time.Now()
	switch {
	case nowTime.Hour() < 12:
		fmt.Println("It's before noon.")
	default:
		fmt.Println("It's afternoon.")
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("Booly")
		case int:
			fmt.Println("Intish")
		default:
			fmt.Printf("Unknown type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")

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
