// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build go1.23

package top

import (
	"cmp"
	"math/rand/v2"
	"slices"
	"testing"
)

func Test(t *testing.T) {
	const N = 20
	perm := rand.Perm(N)

	order := make([]int, N)
	for i := range order {
		order[i] = N - 1 - i
	}
	for i := 0; i <= 20; i++ {
		top := New[int](i, cmp.Compare)
		for _, x := range perm {
			top.Add(x)
		}
		all := top.Take()
		if !slices.Equal(all, order[:i]) {
			t.Errorf("Top%d = %v, want %v", i, all, order[:i])
		}
	}
}
