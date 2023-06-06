package main

import "fmt"

func main() {
	graph := Graph{}

	graph.AddNode(Node{}) // 0
	graph.AddNode(Node{}) // 1
	graph.AddNode(Node{}) // 2
	graph.AddNode(Node{}) // 3

	graph.AddEdge(Edge{From: 0, To: 1})
	graph.AddEdge(Edge{From: 0, To: 2})
	graph.AddEdge(Edge{From: 0, To: 3})
	graph.AddEdge(Edge{From: 1, To: 2})
	graph.AddEdge(Edge{From: 1, To: 3})
	graph.AddEdge(Edge{From: 2, To: 3})

	fmt.Println("Graph 1")
	graph.PrintGraph()
	fmt.Println("Eulerian Graph:", graph.IsEulerian())
	fmt.Println("Complete Graph:", graph.IsComplete())

	graph2 := Graph{}

	graph2.AddNode(Node{}) // 0
	graph2.AddNode(Node{}) // 1
	graph2.AddNode(Node{}) // 2
	graph2.AddNode(Node{}) // 3
	graph2.AddNode(Node{}) // 4

	graph2.AddEdge(Edge{From: 0, To: 1})
	graph2.AddEdge(Edge{From: 0, To: 2})
	graph2.AddEdge(Edge{From: 0, To: 3})
	graph2.AddEdge(Edge{From: 0, To: 4})
	graph2.AddEdge(Edge{From: 1, To: 2})
	graph2.AddEdge(Edge{From: 3, To: 4})

	fmt.Printf("\nGraph 2\n")
	graph2.PrintGraph()
	fmt.Println("Eulerian Graph:", graph2.IsEulerian())
	fmt.Println("Complete Graph:", graph2.IsComplete())

}
