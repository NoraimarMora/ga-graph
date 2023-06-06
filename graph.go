package main

import "fmt"

type Graph struct {
	Nodes         []Node
	OutGoingEdges [][]Edge
}

type Edge struct {
	From, To int
	Weight   float64
}

type Node struct {
	ID int
}

func (g *Graph) AddNode(n Node) {
	n.ID = len(g.Nodes)
	g.Nodes = append(g.Nodes, n)
	g.OutGoingEdges = append(g.OutGoingEdges, []Edge{})
}

func (g *Graph) AddEdge(e Edge) {
	if (e.From < 0 || e.From >= len(g.Nodes)) || (e.To < 0 || e.To >= len(g.Nodes)) {
		return
	}
	g.OutGoingEdges[e.From] = append(g.OutGoingEdges[e.From], e)
	g.OutGoingEdges[e.To] = append(g.OutGoingEdges[e.To], Edge{From: e.To, To: e.From, Weight: e.Weight})
}

func (g *Graph) IsEulerian() bool {
	oddVertices := 0

	for _, e := range g.OutGoingEdges {
		if len(e) == 0 {
			return false
		}

		if len(e)%2 != 0 {
			oddVertices++
		}

		if !(oddVertices == 0 || oddVertices == 2) {
			return false
		}
	}

	return true
}

func (g *Graph) IsComplete() bool {
	for _, e := range g.OutGoingEdges {
		if len(e) != len(g.Nodes)-1 {
			return false
		}
	}

	return true
}

func (g *Graph) PrintGraph() {
	for _, v := range g.Nodes {
		fmt.Printf("Vertex %d -> Adjacent nodes: [ ", v.ID)
		for _, e := range g.OutGoingEdges[v.ID] {
			fmt.Printf("%d ", e.To)
		}
		fmt.Println("]")
	}
}
