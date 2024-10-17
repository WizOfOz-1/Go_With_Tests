package main

import "fmt"

const englishHelloPrefix = "Hello, "

func Hello(receipent string) string {
	if receipent == "" {
        receipent = "World "
	}
	return englishHelloPrefix + receipent
}

func main() {
	fmt.Println(Hello("world"))
}
