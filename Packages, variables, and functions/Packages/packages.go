package main

import (
	"fmt"
	"math/rand" // math package has a sub-package rand.
	/*
		A common good practice is that you have the same package name
		as that of directory of that package.
	*/)

func main() {
	fmt.Println("My favorite number is", rand.Intn(10))
}
