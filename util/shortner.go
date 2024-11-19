package util

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"sync"
)

var mu sync.Mutex

func hashSha256(data string) string {
	h := sha256.New()
	h.Write([]byte(data))
	return string(h.Sum(nil))
}

func ShortUrl(url string) string {
	mu.Lock()
	defer mu.Unlock()
	enc := hashSha256(url)
	enc = enc[:8]
	result := base64.URLEncoding.EncodeToString([]byte(enc))
	fmt.Println("result: ", result)
	return result
}
