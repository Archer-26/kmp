package kmp

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	str := "fecaba"
	m := BuildNext(str)
	fmt.Println(m.Matching("ababa"))
	fmt.Println(m.Matching("hvew9rh834n ff4feba39rjfv nfhufecabaht43q0yffecab87hv"))

}
