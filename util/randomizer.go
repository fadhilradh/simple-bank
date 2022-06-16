package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

func GenerateRandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min +1)
}

func GenerateRandomStr(n int) string {
	var sb strings.Builder

	strLength := len(alphabet)

	// for i := 0; i< n; i++
}