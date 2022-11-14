package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Icecream  []string
}

func main() {
	p1 := Person{
		FirstName: "y",
		LastName:  "z",
		Icecream:  []string{"1", "2"},
	}

	p2 := Person{
		FirstName: "p",
		LastName:  "t",
		Icecream:  []string{"3", "4"},
	}

	fmt.Println(p1.Icecream, p2.Icecream)
	for _, v := range p1.Icecream {
		fmt.Println(v)
	}

	m := map[string]Person{
		p1.LastName: p1,
		p2.LastName: p2}

	fmt.Println(m)
}
