package main

import "fmt"

type Vertex struct {
    label string
	edges []*Edge
}

type Edge struct {
	from  *Vertex
	to    *Vertex
	value float32
}

type Graph struct {
	vertices []Vertex
}

func main() {

	edges := [5]*Edge{}
	s := Vertex{"s",edges[0:2]}
	t := Vertex{"t",edges[5:5]} // empty
	u := Vertex{"u",edges[2:4]}
	v := Vertex{"v",edges[4:5]}

    // s := Vertex{label:"s"}
	// t := Vertex{label:"t"}
	// u := Vertex{label:"u"}
	// v := Vertex{label:"v"}

	su := Edge{from: &s, to: &u, value: 10}

	sv := Edge{from: &s, to: &v, value: 5}
	ut := Edge{from: &u, to: &t, value: 5}
	uv := Edge{from: &u, to: &v, value: 15}
	vt := Edge{from: &v, to: &t, value: 10}
	edges = [5]*Edge{&su, &sv, &ut, &uv, &vt}

    vertices := [4]Vertex{s, t, u, v}
	mygraph := Graph{vertices: vertices[0:4]}
	for i, v := range mygraph.vertices {
		fmt.Println(i, v.label, ": ")
		for j, e := range v.edges {
			fmt.Println(j, (*e).to.label,(*e).value)
		}
        fmt.Println()
	}
}
