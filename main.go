package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	index := 20
	cost := 10
	chanel_cost := make(chan int, 1)
	chanel_index := make(chan int, 1)
	chanel_cost <- cost
	chanel_index <- index
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go broker(chanel_cost, chanel_index, &wg)
	}
	wg.Wait()
}
func sell(chanel_index, chanel_cost chan int) {
	cost := <-chanel_cost
	index := <-chanel_index
	if cost <= 0 {
		cost = 20
	}
	index -= 1
	cost -= 1
	chanel_index <- index
	chanel_cost <- cost
}
func buy(chanel_index, chanel_cost chan int) {
	cost := <-chanel_cost
	index := <-chanel_index
	if cost <= 0 {
		cost = 20
	}
	index += 1
	cost += 1
	chanel_index <- index
	chanel_cost <- cost
}
func broker(chanel_index, chanel_cost chan int, wg *sync.WaitGroup) {
	for {
		time.Sleep(3 * time.Second)
		index := <-chanel_index
		chanel_index <- index
		if index <= 0 {
			fmt.Println("Index = 0")
			break
		}
		rand.Seed(time.Now().UnixNano())
		count_shares := 50 + rand.Intn(50)
		rand.Seed(time.Now().UnixNano())
		random := rand.Intn(2)
		count := 1 + rand.Intn(10)
		if random == 0 {
			fmt.Println("Sold ", "count ", count)
			sell(chanel_index, chanel_cost)
			count_shares -= count
		} else {
			fmt.Println("Bought ", "count ", count)
			buy(chanel_index, chanel_cost)
			count_shares += count
		}
	}
	wg.Done()
}
