package Onyx

import "testing"

func TestPickRandomVertext(T *testing.T) {
	graph, _ := NewGraph("/tmp/onyxsdlkjf", false)
	defer graph.Close()
	_ = graph.AddEdge("a", "b", nil)
	_ = graph.AddEdge("b", "c", nil)
	_ = graph.AddEdge("c", "d", nil)
	_ = graph.AddEdge("d", "e", nil)

	v, _ := graph.PickRandomVertex(nil)
	T.Log("==", string(v), "==")
}

func TestInsertAndRead(T *testing.T) {
	graph, _ := NewGraph("/tmp/onyxsdlkjf", false)
	defer graph.Close()
	err := graph.AddEdge("a", "b", nil)
	if err != nil {
		T.Fatal(err)
	}
	err = graph.AddEdge("a", "c", nil)
	if err != nil {
		T.Fatal(err)
	}
	err = graph.AddEdge("a", "d", nil)
	if err != nil {
		T.Fatal(err)
	}

	dstNodes, err := graph.GetEdges("a", nil)
	if err != nil {
		T.Fatal(err)
	}

	for _, node := range []string{"b", "c", "d"} {
		if _, ok := dstNodes[node]; !ok {
			T.Fatalf("%s not in edgelist", node)
		}
	}
}
