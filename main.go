package main

import "fmt"

func main() {
	data := []int{5, 2, 4, 8, 6, 1, 3, 7, 0, 14}
	fmt.Println("source:", data)
	fmt.Println("sorting by insert:", sortingByInsert(data))
	fmt.Println("sorting by merge:", sortingByMerge(data, 0, len(data) - 1))
}

// Сортировка вставкой
func sortingByInsert(a []int) []int {
	for j := 1; j < len(a); j++ {
		key := a[j]
		i := j - 1
		for {
			if i >= 0 && a[i] > key {
				a[i+1], a[i] = a[i], key
				i--
			} else {
				break
			}
		}
	}
	return a
}

func merge(a []int, p int, q int, r int) []int {
	n1 := q - p + 1
	n2 := r - q
	var tl  = []int{}
	var tr  = []int{}
	for i := 0; i < n1; i++ {
		tl = append(tl, a[p+i])
	}
	for j := 0; j < n2; j++ {
		tr = append(tr, a[q+j+1])
	}
	//tl = append(tl, ^int(0))
	tl = append(tl, 999)
	tr = append(tr, 999)

	i := 0
	j := 0
	for k := p; k < r; k++ {
		if tl[i] <= tr[j] {
			a[k] = tl[i]
			i++
		} else {
			a[k] = tr[j]
			j++
		}
	}
	return a
}

// Сортировка вставкой
func sortingByMerge(a []int, p int, r int) []int {
	if p < r {
		q := (p + r) / 2
		a = sortingByMerge(a, p, q)
		a = sortingByMerge(a, q+1, r)
		a = merge(a, p, q, r)
	}

	return a
}
