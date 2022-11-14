package main

import "fmt"

type Person struct {
	Age  int
	Name string
}

func (p *Person) ChangeMe() {
	p.Name = "yang"
}

func main() {
	p := Person{
		Age:  32,
		Name: "x"}
	fmt.Println(p)
	(&p).ChangeMe()
	fmt.Println(p)
}
