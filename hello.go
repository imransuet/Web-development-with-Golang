package main

import "fmt"

func main() {

	var i int = 10
	var j string
	i = 20
	j = "hi"
	fmt.Println("%T %T", i, j)
	fmt.Println("%v %v", i, j)
	fmt.Println(j)
}
