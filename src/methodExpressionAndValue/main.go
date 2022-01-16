package main

import "fmt"

type Number struct {
	n int
}

func (n *Number) Add(i int) int {
	n.n = n.n + i
	return n.n
}

func main() {
	// メソッド式の第一引数はレシーバになる
	add := (*Number).Add
	n := &Number{1}
	fmt.Printf("%T\n", add)
	fmt.Println(add(n, 2))
	fmt.Println(add(n, 2))

	increment := (&Number{1}).Add
	fmt.Printf("%T\n", increment)
	fmt.Println(increment(2))
	fmt.Println(increment(2))
}
