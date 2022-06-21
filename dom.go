// GO-DOM a go package that implements the document object model in golang
// the go-dom is a full implementation of the model that browsers implement minus xml namespaces and legacy depreciated features. 
package tree

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
// returns a node and an interface on a node
func bfs (node *Node, find string) *Node {
	var found Node
	visited := make(map[string]bool)
	queue := []Node{}
	visited[`%s`, node] = true
	queue = append(queue, node)
	for !queue.hasElements() {

	}

	return &found
}


func makeTree (nodePointer *Node) *Node {
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

func sortTree (nodePointer *Node) *Node {
	return nodePointer
}
