package node

// Node interface
type Node interface {
	isNode()
}

// Root struct
type Root struct {
	Val string
}

func (Root) isNode() {}

// Leaf struct
type Leaf struct {
	Val string
}

func (Leaf) isNode() {}
