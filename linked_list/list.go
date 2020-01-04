package linked_list

var head *Node

type Node struct {
	data string
	next *Node
}

func Add(data string) {
	node := Node{data, nil}
	head = &node
}
