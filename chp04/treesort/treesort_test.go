package treesort_test

// using standard test style from gopl.io

import (
	"math/rand"
	"sort"
	"testing"

	"github.com/aliwatters/go-course-02/chp04/treesort"
)

func TestSort(t *testing.T) {
	data := make([]int, 50)
	for i := range data {
		data[i] = rand.Int() % 50
	}
	treesort.Sort(data)

	if !sort.IntsAreSorted(data) {
		t.Errorf("not sorted: %v", data)
	}
}
