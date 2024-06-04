// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package top implements a “top N” accumulator.
package top

import (
	"container/heap"
	"sort"
)

// A TopN keeps the top (greatest) N elements from
// a set of elements added incrementally.
type TopN[E any] struct {
	top  []E
	n    int
	cmp  func(E, E) int
	heap bool
}

type topN[E any] TopN[E]

func (t *topN[E]) Len() int           { return len(t.top) }
func (t *topN[E]) Swap(i, j int)      { t.top[i], t.top[j] = t.top[j], t.top[i] }
func (t *topN[E]) Less(i, j int) bool { return t.cmp(t.top[i], t.top[j]) < 0 }
func (t *topN[E]) Push(x any)         { t.top = append(t.top, x.(E)) }
func (t *topN[E]) Pop() any           { t.top = t.top[:len(t.top)-1]; return nil }

// New creates a new TopN keeping the top (greatest) N elements
// according to the ordering function cmp.
func New[E any](N int, cmp func(E, E) int) *TopN[E] {
	return &TopN[E]{n: N, cmp: cmp}
}

// Add adds a new element to the set.
func (t *TopN[E]) Add(x E) {
	if t.n == 0 {
		return
	}
	if len(t.top) < t.n {
		t.top = append(t.top, x)
		return
	}
	if !t.heap {
		heap.Init((*topN[E])(t))
		t.heap = true
	}
	if t.cmp(x, t.top[0]) < 0 {
		return
	}
	heap.Push((*topN[E])(t), x)
	heap.Pop((*topN[E])(t))
}

// Take returns the top N elements among those added with [TopN.Add].
// The order of elements that compare equal is unspecified.
// Take takes the elements out of t, resetting it for another collection.
func (t *TopN[E]) Take() []E {
	sort.Sort(sort.Reverse((*topN[E])(t)))
	t.heap = false
	x := t.top
	t.top = nil
	return x
}
