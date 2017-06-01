package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/JoaquinTrinanes/semaphore"
)

var freeSlots, fullSlots, mutex *semaphore.Semaphore
var wg sync.WaitGroup //to continue execution when all goroutines are finished

const bufferSize uint32 = 8

//number of elements in buffer, 0 by default
var count uint

var buffer [bufferSize]int

func productor() {
	defer wg.Done()

	//create 20 elements
	for i := 0; i < 20; i++ {
		freeSlots.Down()
		mutex.Down()
		buffer[count] = i
		count++
		fmt.Printf("Produced %d at %d\n", i, count)

		mutex.Up()
		fullSlots.Up()
		//sleep between 0 and 4 seconds
		time.Sleep(time.Duration(rand.Intn(4)) * time.Second)
	}
}

func consumer() {
	defer wg.Done()

	for i := 0; i < 20; i++ {
		fullSlots.Down()
		mutex.Down()

		count--
		fmt.Printf("Consumed %d at %d\n", buffer[count], count)

		mutex.Up()
		freeSlots.Up()
		time.Sleep(time.Duration(rand.Intn(4)) * time.Second)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	wg.Add(2) //we'll be waiting for two goroutines to finish

	freeSlots = semaphore.SemInit(bufferSize)
	fullSlots = semaphore.SemInit(0) //not necessary, 0, by default
	mutex = semaphore.SemInit(1)

	go productor()
	go consumer()
	wg.Wait() //wait for both goroutines to finish
}
