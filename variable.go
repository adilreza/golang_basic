package main

import "fmt"

func main() {
	//fmt.Printf("Hello printf")
	fmt.Println("Helloo println")
	//fmt.Print("Hello print")
	var i, j, sum int
	var a, b, c = 3, 4.99, "foo"
	sum = 0
	i = 10
	j = 20
	sum = i + j
	fmt.Println(sum)
	fmt.Printf("a is of type %T\n", a)
	fmt.Printf("b is of type %T\n", b)
	fmt.Printf("c is of type %T\n", c)

}
