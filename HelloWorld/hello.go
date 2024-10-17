package main

import "fmt"

func Hello(receipent string) string {
	return fmt.Sprintf("Hello, %s", receipent)
}

func main() {
	fmt.Println(Hello("world"))
}
