package linear

import (
	"container/heap"
	"fmt"
)

type IntegerHeap []int

func (h IntegerHeap) Len() int           { return len(h) }
func (h IntegerHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntegerHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntegerHeap) Push(item interface{}) {
	*h = append(*h, item.(int))
}

func (h *IntegerHeap) Pop() interface{} {
	var n, x1 int
	var prev = *h
	n = len(prev)
	x1 = prev[n-1]
	*h = prev[0 : n-1]

	return x1
}
