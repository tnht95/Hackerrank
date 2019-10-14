//https://www.hackerrank.com/challenges/journey-to-the-moon/problem
package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	matrix  [][]int
	visited []bool
	remain  int
)

func dfs(x int) int {
	stack := []int{x}
	visited[x] = true
	//count the number of people from the same country
	count := 0

	for len(stack) > 0 {
		cur := pop(&stack)
		count++
		//people still available to pair
		remain--
		for i := range matrix[cur] {
			if !visited[matrix[cur][i]] {
				stack = append(stack, matrix[cur][i])
				visited[matrix[cur][i]] = true
			}
		}

	}

	//each person from the same country can be pair with same amount of people that are still available
	return count * remain
}

func pop(s *[]int) int {
	n := (*s)[len(*s)-1]
	(*s) = (*s)[:len(*s)-1]
	return n
}

func main() {
	r := bufio.NewReader(os.Stdin)
	var n, q int
	fmt.Fscanf(r, "%d %d\n", &n, &q)

	matrix = make([][]int, n)

	for i := 0; i < q; i++ {
		var p1, p2 int
		fmt.Fscanf(r, "%d %d\n", &p1, &p2)
		matrix[p1] = append(matrix[p1], p2)
		matrix[p2] = append(matrix[p2], p1)
	}

	visited = make([]bool, n)
	remain = n
	result := 0
	for i := range matrix {
		if !visited[i] {
			result += dfs(i)
		}
	}

	fmt.Fprintf(os.Stdout, "%d\n", result)
}
