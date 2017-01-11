// M receivers, one sender, the sender says "no more sends" by closing the data channel

package main

import (
	"log"
	"math/rand"
	"sync"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	log.SetFlags(0)

	const MaxRandomNumber = 100000
	const NumReceivers = 100

	wgReceivers := sync.WaitGroup{}
	wgReceivers.Add(NumReceivers)

	dataCh := make(chan int, 100)

	// the sender
	go func() {
		for {
			value := rand.Intn(MaxRandomNumber)
			if value == 0 {
				close(dataCh)
				return
			}
			dataCh <- value
		}
	}()

	// receivers
	for i := 0; i < NumReceivers; i++ {
		go func() {
			defer wgReceivers.Done()

			for value := range dataCh {
				log.Println(value)
			}
		}()
	}

	wgReceivers.Wait()
}
