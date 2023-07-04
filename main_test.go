package FirstGoProject

import (
	"github.com/Z-oney/FirstGoProject/src/apple"
	"testing"
)

func BenchmarkTrieB(c *testing.B) {
	c.Log(c.N)
}
func BenchmarkTrie_Remove_and_Compact(b *testing.B) {
	apple.Print()
}
