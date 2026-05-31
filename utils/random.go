package utils

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

func radomInit(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func randomString(n int) string {
	var sb strings.Builder

	k := len(alphabet)

	for i := 0; i < k; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

// random owner generates random owner name
func RandomOwner() string {
	return randomString(6)
}

// random money
func RandomMoney() int64 {
	return radomInit(0, 1000)
}

// random currency
func RandomCurrency() string {
	currency := []string{"USD", "EUR", "CAD"}

	k := len(currency)

	return currency[rand.Intn(k)]
}
