package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"time"
)

const (
	WHITE = 0
	RED   = iota
	BLUE  = iota
)

func main() {
	totalIterations, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	rand.Seed(time.Now().UnixNano() + int64(os.Getpid()))

	fmt.Printf("# %d iterations\n", totalIterations)
	wins, losses := 0, 0

	distribution := make(map[int]int)

	bag := make([]int, 256)
	bag[0] = RED
	bag[1] = BLUE
	bag[2] = WHITE

	l := 3

	for i := 0; i < totalIterations; i++ {
		playing := true
		for playing {
			ball := bag[rand.Intn(l)]
			switch ball {
			case RED:
				playing = false
				losses++
				distribution[l-2]++
			case BLUE:
				playing = false
				wins++
				distribution[l-2]++
			case WHITE:
				l++
				if l > len(bag) {
					bag = append(bag, WHITE)
				}
			}
		}
		bag = bag[0:3]
		l = 3
	}

	iterations := make([]int, len(distribution))
	i := 0
	for n, _ := range distribution {
		iterations[i] = n
		i++
	}
	sort.Sort(sort.IntSlice(iterations))

	total := 0
	moment := 0
	for j := range iterations {
		k := iterations[j]
		n := distribution[k]
		total += n
		moment += k * n
		if n > 0 {
			fmt.Printf("%d\t%d\t%.5f\n", k, n, float64(n)/float64(totalIterations))
		}
	}
	fmt.Printf("# %d wins, %d losses\n", wins, losses)
	fmt.Printf("# %.4f wins, %.4f losses\n", float64(wins)/float64(total), float64(losses)/float64(total))
	fmt.Printf("# Total %d\n", total)
	fmt.Printf("# Ave number of choices in a game: %.4f\n", float64(moment)/float64(total))
}
