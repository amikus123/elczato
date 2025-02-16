package learning

import "fmt"

func toChan(arr []int) <-chan int {
	chan1 := make(chan int)
	go func() {
		for _, i := range arr {
			chan1 <- i
		}
		close(chan1)
	}()
	return chan1
}

func double(chan1 <-chan int) <-chan int {
	chan2 := make(chan int)
	go func() {
		for c := range chan1 {
			chan2 <- c * 2
		}
		close(chan2)
	}()
	return chan2
}

func square(chan1 <-chan int) <-chan int {
	chan2 := make(chan int)
	go func() {
		for c := range chan1 {
			chan2 <- c * c
		}
		close(chan2)
	}()
	return chan2
}

func Main() {
	ints := []int{1, 2, 3, 4, 5}

	chan1 := toChan(ints)
	chan2 := double(chan1)
	chan3 := square(chan2)

	for c := range chan3 {
		fmt.Println(c)
	}
}
