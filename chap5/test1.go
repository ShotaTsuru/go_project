package main

import (
	"fmt"
)

type MyfuncOpts struct {
	FirstName string
	LastName  string
	Age       int
}

func Myfunc(opts MyfuncOpts) {
	fmt.Printf("Hello %s %s, your age is %d\n", opts.FirstName, opts.LastName, opts.Age)
}

func main() {
	opts := MyfuncOpts{
		FirstName: "John",
		LastName:  "Doe",
		Age:       30,
	}
	Myfunc(opts)

	Myfunc(MyfuncOpts {
		FirstName: "Jane",
		LastName:  "Doe",
		Age:       25,
	})
}
