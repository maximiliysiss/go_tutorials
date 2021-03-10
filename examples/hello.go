package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

type CommonInterface interface {
	init()
	second(one int)
}

func main() {
	random := rand.New(rand.NewSource(int64(time.Now().Nanosecond())))

	scores := make([]int, 100)
	for i := 0; i < 100; i++ {
		scores[i] = int(random.Int31n(1000))
	}
	sort.Ints(scores)

	worst := make([]int, 5)
	copy(worst, scores[:5])
	fmt.Println(worst)
}
