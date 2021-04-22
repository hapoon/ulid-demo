package main

import (
	"fmt"
	"io"
	"math/rand"
	"time"

	"github.com/oklog/ulid"
)

func main() {
	{
		gen1 := func() string {
			now := time.Now()
			t := now.UnixNano()
			entropy := ulid.Monotonic(rand.New(rand.NewSource(t)), 0)
			id := ulid.MustNew(ulid.Timestamp(now), entropy).String()
			return id
		}

		gen2 := func() string {
			t := time.Unix(time.Now().Unix(), 0)
			entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
			id := ulid.MustNew(ulid.Timestamp(t), entropy).String()
			return id
		}

		fmt.Println("Unixtime with nanosec")
		fmt.Println(gen1())
		fmt.Println(gen1())
		fmt.Println(gen1())
		fmt.Println(gen1())

		fmt.Println("Unixtime with sec")
		fmt.Println(gen2())
		fmt.Println(gen2())
		fmt.Println(gen2())
		fmt.Println(gen2())

		e := ulid.Monotonic(rand.New(rand.NewSource(time.Now().UnixNano())), 0)
		fmt.Println("genUlid")
		fmt.Println(genUlid(e))
		fmt.Println(genUlid(e))
		fmt.Println(genUlid(e))
		fmt.Println(genUlid(e))
	}
}

func genUlid(entropy io.Reader) string {
	id := ulid.MustNew(ulid.Timestamp(time.Now()), entropy)
	return id.String()
}
