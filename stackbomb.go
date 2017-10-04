package main;

import "fmt"

func stackbomb(n int) {
	fmt.Println(n)
	stackbomb(n+1)
}

func main() {
	stackbomb(1)
}
