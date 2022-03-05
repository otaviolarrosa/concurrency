package main

import "fmt"

func main() {
	c := make(chan string)
	fmt.Println("Sending to the channel")
	go func(input chan string) {
		c <- "hello"
	}(c)
	fmt.Println("receiving from the channel")
	greeting := <-c
	fmt.Println("greeting received")

	fmt.Println(greeting)
}

func HelloWorld() {
	fmt.Println("Hello World")
}

//a única maneira de uma goroutine se comunicar com outra, é através de channels
