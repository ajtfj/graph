package graph

type distance struct {
	dist int
	node Node
}

type distanceHeap []distance

func (h distanceHeap) Len() int           { return len(h) }
func (h distanceHeap) Less(i, j int) bool { return h[i].dist < h[j].dist }
func (h distanceHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *distanceHeap) Push(dist interface{}) {
	*h = append(*h, dist.(distance))
}

func (h *distanceHeap) Pop() interface{} {
	old := *h
	n := len(old)
	dist := old[n-1]
	*h = old[0 : n-1]
	return dist
}
