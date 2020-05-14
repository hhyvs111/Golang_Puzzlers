package main

import "fmt"

func main() {

ch2 := make(chan int); // zero channel
ch3 := make(chan int, 10); // buffered channel

fmt.Printf("ch2,len: %v, cap:%v \n",len(ch2),cap(ch2)) 
fmt.Printf("ch2,len: %v, cap:%v \n",len(ch2),cap(ch2)) 
fmt.Printf("ch3,len: %v, cap:%v \n",len(ch3),cap(ch3)) 



	ch1 := make(chan int, 2)
	// 发送方。
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Printf("Sender: sending element %v...\n", i)
			ch1 <- i
		}
		fmt.Println("Sender: close the channel...")
		close(ch1)
	}()

	// 接收方。
	for {
		elem, ok := <-ch1
		if !ok {
			fmt.Println("Receiver: closed channel")
			break
		}
		fmt.Printf("Receiver: received an element: %v\n", elem)
	}

	fmt.Println("End.")
}
