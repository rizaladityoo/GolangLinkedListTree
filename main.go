package main

import (
	"fmt"
	"strings"
)

type Node struct {
	data  string
	left  *Node
	mid   *Node
	right *Node
}

func (root *Node) recursive(level int) {
	if root != nil {
		for x := 1; x <= level; x++ {
			fmt.Print("-")
		}
		fmt.Printf("%s\n", root.data)
		root.left.recursive(level + 1)
		root.mid.recursive(level + 1)
		root.right.recursive(level + 1)
	}
}

func main() {
	root := Node{"A", nil, nil, nil}
	root.left = &Node{"B", nil, nil, nil}
	root.right = &Node{"C", nil, nil, nil}
	root.left.left = &Node{"D", nil, nil, nil}
	root.left.right = &Node{"E", nil, nil, nil}
	root.right.left = &Node{"F", nil, nil, nil}
	root.right.mid = &Node{"G", nil, nil, nil}
	root.right.right = &Node{"H", nil, nil, nil}

	root.recursive(0)
	var s string

	fmt.Print("program mencari alphabet pada tree\n")
	fmt.Print("masukan alphabet : ")
	fmt.Scan(&s)
	fmt.Print(root.search(strings.ToUpper(s)))

}

func (root *Node) search(s string) bool {
	if string(s[0]) != "A" {
		return false
	}

	return len(root.recursiveSearch(s, 0)) >= len(s)

}

var temp []string

func (root *Node) recursiveSearch(s string, level int) []string {
	if root != nil {
		temp = append(temp, root.data)
		if root.left != nil {
			if root.left.data == string(s[level+1]) {
				root.left.recursiveSearch(s, level+1)
			}
		}
		if root.mid != nil {
			if root.mid.data == string(s[level+1]) {
				root.mid.recursiveSearch(s, level+1)
			}
		}
		if root.right != nil {
			if root.right.data == string(s[level+1]) {
				root.right.recursiveSearch(s, level+1)
			}
		}
	}
	return temp
}
