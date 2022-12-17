package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scores := map[string]struct{ f, s int }{
		"A X": {1 + 3, 3 + 0}, "A Y": {2 + 6, 1 + 3}, "A Z": {3 + 0, 2 + 6},
		"B X": {1 + 0, 1 + 0}, "B Y": {2 + 3, 2 + 3}, "B Z": {3 + 6, 3 + 6},
		"C X": {1 + 6, 2 + 0}, "C Y": {2 + 0, 3 + 3}, "C Z": {3 + 3, 1 + 6},
	}
	score, score2 := 0, 0
	for scanner.Scan() {
		score += scores[scanner.Text()].f
		score2 += scores[scanner.Text()].s
	}
	fmt.Println("Part 1: %d , Part2: %d", score2)
}

