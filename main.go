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

func sloane(angka int) []int {
	array := []int{1}
	for i := 1; i <= angka; i++ {
		array = append(array, i+array[i-1])
	}
	return array
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
	var angka int
	hasil1 := sloane(20)
	fmt.Printf("%d \nMasukan angka : ", hasil1[len(hasil1)-1])

	fmt.Scan(&angka)
	deret := sloane(angka)
	deret = deret[:len(deret)-1]
	fmt.Println(deret)

	root := Node{"A", nil, nil, nil}
	root.left = &Node{"B", nil, nil, nil}
	root.right = &Node{"C", nil, nil, nil}
	root.left.left = &Node{"D", nil, nil, nil}
	root.left.right = &Node{"E", nil, nil, nil}
	root.right.left = &Node{"F", nil, nil, nil}
	root.right.mid = &Node{"G", nil, nil, nil}
	root.right.right = &Node{"H", nil, nil, nil}

	root.recursive(0)
	var s, child string

	fmt.Print("program mencari path pada tree\n")
	fmt.Print("masukan alphabet : ")
	fmt.Scan(&s)
	fmt.Print(root.search(strings.ToUpper(s)))

	fmt.Print("\nprogram mencari child pada tree\n")
	fmt.Print("masukan alphabet : ")
	fmt.Scan(&child)
	fmt.Print(root.searchChild(strings.ToUpper(child)))

}

func (root *Node) searchChild(c string) string {
	temp = nil
	root.recursiveSearch(c, 0, &temp)
	if temp[len(temp)-1].left == nil && temp[len(temp)-1].mid == nil && temp[len(temp)-1].right == nil {
		return "Tidak ada child"
	}
	return temp[len(temp)-1].left.data + "-" + temp[len(temp)-1].right.data
}

func (root *Node) search(s string) bool {
	if string(s[0]) != "A" {
		return false
	}
	root.recursiveSearch(s, 0, &temp)
	return len(temp) >= len(s)-1

}

var temp []Node

func (root *Node) recursiveSearch(s string, level int, temp *[]Node) {
	if root != nil {
		if len(s) > level+1 {
			if root.left != nil {
				if root.left.data == string(s[level+1]) {
					*temp = append(*temp, Node{
						data:  root.left.data,
						left:  root.left.left,
						mid:   root.left.mid,
						right: root.left.right,
					})
					root.left.recursiveSearch(s, level+1, temp)
				}
			}
			if root.mid != nil {
				if root.mid.data == string(s[level+1]) {
					*temp = append(*temp, Node{
						data:  root.mid.data,
						left:  root.mid.left,
						mid:   root.mid.mid,
						right: root.mid.right,
					})
					root.mid.recursiveSearch(s, level+1, temp)
				}
			}
			if root.right != nil {
				if root.right.data == string(s[level+1]) {
					*temp = append(*temp, Node{
						data:  root.right.data,
						left:  root.right.left,
						mid:   root.right.mid,
						right: root.right.right,
					})
					root.right.recursiveSearch(s, level+1, temp)
				}
			}
		}
	}
}
