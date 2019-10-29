package main

import "fmt"

func main() {
	var a, b, c = 20, 3, 40
	fmt.Println("--------------------testing for if else---------------")
	fmt.Println("--------------------this code only for learning ---------------")
	if a >= b {
		if a >= c {
			fmt.Printf("BIg is a = %d", a)
		} else {
			fmt.Printf("Big is c = %d", c)
		}

	} else if b >= a {

		if b >= c {
			fmt.Println("BIg is b = %d", b)
		}
	} else {
		fmt.Println("Big is c = %d", c)
	}
}
