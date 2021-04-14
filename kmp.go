package kmp

type KMP struct {
	p string
	m []int
}

func (s *KMP) Matching(src string) []int {
	ret := make([]int, 0)
	i := 0
	j := 0
	ti := i
	for i < len(src) {
		for ti < len(src) && j < len(s.p) && src[ti] == s.p[j] {
			j++
			ti++
			if j == len(s.p) {
				ret = append(ret, i)
			}
		}
		if j == 0 {
			i++
			ti = i
		} else {
			i += ti - i - s.m[j-1]
			ti = i + s.m[j-1]
			j = s.m[j-1]
		}
	}
	return ret
}

func BuildNext(p string) *KMP {
	m := make([]int, len(p), len(p))

	for i1 := 1; i1 < len(p); i1++ {
		i2 := 0
		c := 0
		for i1+c < len(p) && p[i1+c] == p[i2] {
			if m[i1+c] < c+1 {
				m[i1+c] = c + 1
			}
			c++
			i2++
		}
	}
	return &KMP{p: p, m: m}
}
