package main

import "fmt"

// 单向通道

func counter(out chan<- int) {
	for x := 0; x <= 10; x++ {
		out <- x
	}
	close(out)
}

func squarer(out chan<- int, in <-chan int) {
	for x := range in {
		out <- x * x
	}
	close(out)
}

func printer(in <-chan int) {
	for x := range in {
		fmt.Println(x)
	}
}

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	// counter
	go counter(naturals)

	// squarer
	go squarer(squares, naturals)

	// printer, in main goroutine
	printer(squares)

}
