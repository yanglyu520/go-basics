package main

import "fmt"

func main() {
	a := 42
	b := &a
	fmt.Println(&a)
	fmt.Printf("%T", &a)
	fmt.Println(*b)

	//change the pointed value to something else
	*b = 55
	fmt.Println(a)
}
