package main

import "fmt"

type nameWithOption struct {
	Name string
	*options
}

func newNameWithOption(name string, opts ...option) *nameWithOption {
	o := &options{"default string", 0}
	for _, opt := range opts {
		opt(o)
	}
	return &nameWithOption{name, o}
}

type options struct {
	Option1 string
	Option2 int
}

type option func(*options)

func option1(s string) option {
	return func(options *options) {
		options.Option1 = s
	}
}

func option2(i int) option {
	return func(options *options) {
		options.Option2 = i
	}
}

func main() {
	n := newNameWithOption("test", option1("hoge"), option2(10))
	fmt.Println(n.Name)
	fmt.Println(n.Option1)
	fmt.Println(n.Option2)
}