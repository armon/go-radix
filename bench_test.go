package radix

import (
	"strconv"
	"testing"
)

var suffixes = []string{
	"a",
	"ab",
	"abc",
	"abx",
	"abya",
	"abyb",
	"aby",
	"abz",
	"abcd",
	"abcdef",
	"abcde",
	"x",
	"xyz",
	"xy",
}

func createTree() (*Tree[string], []string) {
	r := New[string]()
	paths := make([]string, 0)
	for c := 'a'; c < 'i'; c++ {
		for _, path := range suffixes {
			paths = append(paths, string(c)+path)
		}
	}

	for _, path := range paths {
		r.Insert(path, path)
	}
	return r, paths
}

func BenchmarkInsert(b *testing.B) {
	r, paths := createTree()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		for _, path := range paths {
			r.Insert(path+strconv.Itoa(n), path)
		}
	}
}

func BenchmarkGet(b *testing.B) {
	r, paths := createTree()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		for _, path := range paths {
			actual, ok := r.Get(path)
			if !ok {
				b.Fatalf("Expected %s, got nothing", path)
			}
			if actual != path {
				b.Fatalf("Expected %s, got %s", path, actual)
			}
		}
	}
}

func BenchmarkLongestPrefix(b *testing.B) {
	r, paths := createTree()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		for _, path := range paths {
			actual, _, ok := r.LongestPrefix(path)
			if !ok {
				b.Fatalf("Expected %s, got nothing", path)
			}
			if actual != path {
				b.Fatalf("Expected %s, got %s", path, actual)
			}
		}
	}
}
