// GO-DOM a go package that implements the document object model in golang
// the go-dom is a full implementation of the model that browsers implement minus xml namespaces and legacy depreciated features. 
package dom

import "fmt"

type Node struct {
	parent *Node
	desc string
	name string
	attributes map[string]map[string]string
	children []*Node
}
type DOM interface {
	sortTree() *Node
	makeTree() *Node
}

// family node
// var Cain, Abel, Seth, Awan, Azura = "Cain", "Abel", "Seth", "Awan", "Azura"

// var firstChildren = []string{Cain, Abel, Seth, Awan, Azura}

// Adam := {
//     parent: "",
//     children: make([&Cain, &Abel, &Seth, &Awan, &Azura]string)
// }

func newNode (name string, parent *Node, children []*Node) *Node {
    node := Node{
        parent: parent,
				name: name,
				children: children,
    }
    node.children = children
		node.name = name
    fmt.Println(node)
    return &node
	}

// add error handling for functions at some point
func strhas (slicePointer *[]string) bool {
	slice := *slicePointer
	return len(slice) < 1
}

func strpop (slicePointer *[]string) string {
	slice := *slicePointer
	result := slice[len(slice) - 1]
	slice = slice[:len(slice) - 1]
	return result
}

func strshift (slicePointer *[]string) string {
	slice := *slicePointer 
	result := slice[0]
	slice = slice[1:]
	return result
}

func (slice nodeSlice) empty() bool {
	return len(slice) < 1
}

func (slice nodeSlice) pop()  nodeSlice {
	return slice[:len(slice)-1]
}

type nodeSlice []Node

type treeNode interface {
	pop() nodeSlice
	empty() bool
}

// returns a node and an interface on a node
func bfs (nodePointer *Node, treePointer *map[string]Node, find string) *Node {
	tree := *treePointer
	node := *nodePointer
	var found Node
	visited := make(map[string]bool)
	queue := []string{}
	visited[node.name] = true
	queue = append(queue, node.name)
	for !strhas(&queue) {
		m := strshift(&queue)
		fmt.Sprintf("%s",m)
		for _, nodeP := range tree[m].children {
			_, ok := 
			if nodeP 
		}
	}

	return &found
}

func (nodePointer *Node) makeTree() *Node {
		node := *nodePointer
		// traverse the nodeSlice and create a new node for each element
			// empty node pointer list
			// depth or breadth first search needs to be done here 
			bfs()

			for _, child := range node.children {
				newNode(child.name, nodePointer, []*Node{})
			}

		return nodePointer
}

func (nodePointer *Node) sortTree() *Node {
	return nodePointer
}

type Tree struct {
	tree map[string]Node	
}