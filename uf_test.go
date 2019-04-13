package uf

import (
	"testing"
)

func TestInterface(t *testing.T) {
	t.Parallel()
	u := New(5)
	data := [][]int{
		{0, 1},
		{1, 2},
		{2, 3},
		{3, 4},
	}
	for _, pair := range data {
		p, q := pair[0], pair[1]
		if !u.Connected(p, q) {
			u.Union(p, q)
		}
	}
	if c := u.Count(); c != 1 {
		t.Errorf("Expected 1 got %d", c)
	}
	if u.Find(0) != u.Find(1) ||
		u.Find(1) != u.Find(2) ||
		u.Find(2) != u.Find(3) ||
		u.Find(3) != u.Find(4) {
		t.Errorf("all of vertices should be joined now")
	}
	u.Union(0, 1)
	u.Union(1, 2)
	u.Union(2, 3)
	u.Union(3, 4)
	if u.Find(0) != u.Find(1) ||
		u.Find(1) != u.Find(2) ||
		u.Find(2) != u.Find(3) ||
		u.Find(3) != u.Find(4) {
		t.Errorf("all of vertices should still be joined now")
	}
}

func TestInterfaceReversed(t *testing.T) {
	t.Parallel()
	u := New(5)
	data := [][]int{
		{4, 3},
		{3, 2},
		{1, 2},
		{1, 0},
	}
	for _, pair := range data {
		p, q := pair[0], pair[1]
		if !u.Connected(p, q) {
			u.Union(p, q)
		}
	}
	if c := u.Count(); c != 1 {
		t.Errorf("Expected 1 got %d", c)
	}
}

func TestInterfacePanic(t *testing.T) {
	t.Parallel()
	u := New(5)
	defer func() {
		if err := recover(); err == nil {
			t.Errorf("Expected a panic")
		}
	}()
	u.Find(5)
}

func BenchmarkInterface(b *testing.B) {
	u := New(5)
	for i := 0; i < b.N; i++ {
		p, q := i%5, (i+1)%5
		if !u.Connected(p, q) {
			u.Union(p, q)
		}
	}
}
