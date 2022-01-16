package main

import "fmt"

type Number struct {
	n int
}

func (n Number) Add(i int) int {
	return n.n + i
}

func main() {
	// メソッド式の第一引数はレシーバになる
	add := Number.Add
	fmt.Printf("%T\n", add)
	fmt.Println(add(Number{1}, 2))

	increment := Number{1}.Add
	fmt.Printf("%T\n", increment)
	fmt.Println(increment(2))
}
