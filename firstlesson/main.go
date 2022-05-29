package main

import "fmt"

func main() {
	i := int(3)
	f := float64(2.0)
	b := bool(true)
	s := string("zalupa")
	fmt.Println("%v %v %v %q\n", i, f, b, s)
	loop()
}

func loop() {
	x := float64(-1.09)

	if x < 9 {
		println("hello")
	}
}