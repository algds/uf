package uf

import "fmt"

// Interface is how you use the union-find data structure.
type Interface interface {
	Find(p int) int
	Count() int
	Connected(p, q int) bool
	Union(p, q int)
	validate(p int)
}

type instance struct {
	parent []int
	rank   []byte
	count  int
}

func (i *instance) Find(p int) int {
	i.validate(p)
	for p != i.parent[p] {
		i.parent[p] = i.parent[i.parent[p]]
		p = i.parent[p]
	}
	return p
}

func (i *instance) Count() int {
	return i.count
}

func (i *instance) Connected(p, q int) bool {
	return i.Find(p) == i.Find(q)
}

func (i *instance) Union(p, q int) {
	rootP := i.Find(p)
	rootQ := i.Find(q)
	if rootP == rootQ {
		return
	}
	switch {
	case i.rank[rootP] < i.rank[rootQ]:
		i.parent[rootP] = rootQ
	case i.rank[rootP] > i.rank[rootQ]:
		i.parent[rootQ] = rootP
	default:
		i.parent[rootQ] = rootP
		i.rank[rootP]++
	}
	i.count--
}

func (i *instance) validate(p int) {
	if n := len(i.parent); p < 0 || p >= n {
		panic(fmt.Sprintf("index %d is not between 0 and %d", p, n-1))
	}
}

// New creates a new instance of a union-find interface.
func New(n int) Interface {
	created := &instance{
		parent: make([]int, n),
		rank:   make([]byte, n),
		count:  n,
	}
	for i := 0; i < n; i++ {
		created.parent[i] = i
	}
	return created
}
