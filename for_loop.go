// there are no while loop in go language because, for loop can be used in different purpose
package main

import "fmt"

func main() {

	k := 1 //this := is used for initialize the short size of variable
	for ; k <= 10; k++ {
		fmt.Println(k)
	}

	//this like a while loop
	k = 1
	for k <= 10 {
		fmt.Println(k)
		k++
	}

	for k := 1; ; k++ {
		fmt.Println(k)
		if k == 10 {
			break
		}
	}
}
