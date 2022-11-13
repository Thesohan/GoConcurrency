package main

import "fmt"

func bufferedChannel(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
		fmt.Println("successfully wrote ", i, "to ch")
	}
	close(ch)
}

func unbufferedChannel(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
		fmt.Println("successfully wrote ", i, "to ch")
	}

	close(ch)

}
func main() {

	buffered := make(chan int, 2)
	buffered <- 1
	buffered <- 2

	go bufferedChannel(buffered)
	for i := range buffered {
		fmt.Println(i)
	}

	unbuffered := make(chan int)
	go unbufferedChannel(unbuffered)

	for i := range unbuffered {
		fmt.Println(i, "from unbuffered")
	}

}
