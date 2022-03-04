package main

import (
	"fmt"

	"github.com/vrnvu/go-sum-types/internal/node"
)

type foo struct{}

func (foo) isNode() {}

func howToUseIt() {
	r := node.Root{Val: "i'm a root"}
	l := node.Leaf{Val: "i'm a leaf"}
	printNode(r)
	printNode(l)
}

func compileErr() {
	f := foo{}
	printNode(f)
}

func printNode(n node.Node) {
	switch t := n.(type) {
	case node.Root:
		fmt.Printf("is root: %+v\n", t)
	case node.Leaf:
		fmt.Printf("is leaf: %+v\n", t)
	default:
		panic("unexpetec type")
	}
}
