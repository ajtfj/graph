package graph

import (
	"container/heap"
	"fmt"
	"math"
)

type Node string

type Edge struct {
	node   Node
	weight int
}

type Graph struct {
	adj map[Node][]Edge
}

func NewGraph() *Graph {
	return &Graph{
		adj: make(map[Node][]Edge),
	}
}

func (g *Graph) AddEdge(ori Node, dest Node, weight int) {
	g.adj[ori] = append(g.adj[ori], Edge{
		node:   dest,
		weight: weight,
	})
	g.adj[dest] = append(g.adj[dest], Edge{
		node:   ori,
		weight: weight,
	})
}

func (g *Graph) ShortestPath(ori Node, dest Node) ([]Node, error) {
	trace := make(map[Node]*Trace)
	for v := range g.adj {
		trace[v] = &Trace{math.MaxInt, nil}
	}
	trace[ori] = &Trace{0, nil}

	h := distanceHeap{
		distance{0, ori},
	}
	heap.Init(&h)

	for len(h) > 0 {
		top := heap.Pop(&h)
		u := top.(distance).node
		if u == dest {
			path := getPath(trace, &dest)
			return path, nil
		}

		for _, e := range g.adj[u] {
			v := e.node
			weight := e.weight
			if trace[v].dist > trace[u].dist+weight {
				t := trace[v]
				t.dist = trace[u].dist + weight
				t.prev = &u
				heap.Push(&h, distance{t.dist, v})
			}
		}
	}

	return nil, fmt.Errorf("no path found")
}

type Trace struct {
	dist int
	prev *Node
}

func getPath(trace map[Node]*Trace, dest *Node) []Node {
	var path []Node

	if dest == nil {
		return path
	}

	path = getPath(trace, trace[*dest].prev)
	path = append(path, *dest)
	return path
}
