package uuid

import (
	"encoding/hex"
	//"bufio"
	"bytes"
	"fmt"
	"math/rand"
	"strconv"
	"testing"
	"time"
	"unsafe"

	"github.com/missionMeteora/binny.v2"
	puuid "github.com/pborman/uuid"
)

var (
	pubCurr UUID
	pubStr  string
)

func TestMain(t *testing.T) {
	c := New()
	fmt.Println(c.String())
	fmt.Println(c.Time())

	buf := bytes.NewBuffer(nil)
	binny.NewEncoder(buf).Encode(c)

	var n UUID
	binny.NewDecoder(buf).Decode(&n)
	fmt.Println(n.String())
	fmt.Println(n.Time())
}

func BenchmarkCurrent(b *testing.B) {
	for i := 0; i < b.N; i++ {
		pubCurr = New()
	}
	b.ReportAllocs()
}

func BenchmarkCurrentString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		pubStr = New().String()
	}
	b.ReportAllocs()
}

func BenchmarkCurrentParallel(b *testing.B) {
	b.SetParallelism(4)
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			pubCurr = New()
		}
	})

	b.ReportAllocs()
}

func BenchmarkCurrentParallelGen(b *testing.B) {
	b.SetParallelism(4)
	b.RunParallel(func(p *testing.PB) {
		g := NewGen()
		for p.Next() {
			pubCurr = g.New()
		}
	})

	b.ReportAllocs()
}

func BenchmarkBasic(b *testing.B) {
	for i := 0; i < b.N; i++ {
		pubStr = BasicUUID()
	}

	b.ReportAllocs()
}

func BenchmarkBasicParallel(b *testing.B) {
	b.SetParallelism(4)
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			pubStr = BasicUUID()
		}
	})

	b.ReportAllocs()
}

func BenchmarkPbor(b *testing.B) {
	for i := 0; i < b.N; i++ {
		pubStr = puuid.New()
	}

	b.ReportAllocs()
}

func BenchmarkPborParallel(b *testing.B) {
	b.SetParallelism(4)
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			pubStr = puuid.New()
		}
	})

	b.ReportAllocs()
}

func BasicUUID() string {
	now := time.Now().UnixNano()
	randPart := make([]byte, 7)
	if _, err := rand.Read(randPart); err != nil {
		copy(randPart, (*(*[8]byte)(unsafe.Pointer(&now)))[:7])
	}
	return strconv.FormatInt(now, 10)[1:] + hex.EncodeToString(randPart)
}
