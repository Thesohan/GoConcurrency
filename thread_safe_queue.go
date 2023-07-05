package main

import (
	"fmt"
	"math/rand"
	"sync"
)

type ConcurrentQueue struct {
	queue []int32
	muTax sync.Mutex
}

func (q *ConcurrentQueue) Enqueue(value int32) {
	q.muTax.Lock()
	defer q.muTax.Unlock()
	q.queue = append(q.queue, value)
}

func (q *ConcurrentQueue) Deueue() int32 {
	q.muTax.Lock()
	defer q.muTax.Unlock()
	if len(q.queue) == 0 {
		panic("can't dequeu from an empty queue")
	}
	value := q.queue[0]
	q.queue = q.queue[1:]
	return value
}

func (q *ConcurrentQueue) Size() int {
	q.muTax.Lock()
	defer q.muTax.Unlock()
	return len(q.queue)
}

func main() {
	q1 := ConcurrentQueue{queue: make([]int32, 0)}

	wg := sync.WaitGroup{}
	for i := 0; i < 1000000; i++ {
		wg.Add(1)
		go func() {
			q1.Enqueue(rand.Int31())
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(q1.Size())

	for i := 0; i < 1000000; i++ {
		wg.Add(1)
		go func() {
			q1.Deueue()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(q1.Size())

}
