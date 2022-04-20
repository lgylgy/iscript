package selection

import (
	"fmt"
	"log"
	"math/rand"
	"sort"
)

func Select(values []string, number, seed int) ([]string, error) {
	if len(values) < number {
		return nil, fmt.Errorf("not enough values")
	}
	sort.Strings(values)
	rand.Seed(int64(seed))
	rand.Shuffle(len(values), func(i, j int) {
		values[i], values[j] = values[j], values[i]
	})
	log.Println(values)
	return values[:number], nil
}