// You can edit this code!
// Click here and start typing.
package main

import "fmt"

type vehicle struct {
	doors string
	color string
}

type truct struct {
	fourwheel bool
	vehicle
}
type sedan struct {
	luxury bool
	vehicle
}

func main() {
	v1 := vehicle{"black", "yellow"}
	s1 := sedan{
		true,
		v1,
	}
	t1 := truct{
		true,
		v1}

	test := struct {
		name string
		ID   int
	}{
		name: "test1",
		ID:   0001}

	fmt.Println(t1, s1, test)

}
