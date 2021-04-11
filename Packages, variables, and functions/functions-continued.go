package main

import "fmt"

/*
characteristics - has a name
	- Input - Arguments
	- Processing - Function logic/purpose
	- Output - return values
*/
//  name  arg1, arg2    return-type
/*

int add(int x, int y) {
	return x + y
}

*/
func add(x, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}
