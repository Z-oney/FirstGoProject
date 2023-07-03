package src

import (
	"testing"
	"zoney.cn/src/apple"
)

func BenchmarkTrieB(c *testing.B) {
	c.Log(c.N)
}
func BenchmarkTrie_Remove_and_Compact(b *testing.B) {
	apple.Print()
}
