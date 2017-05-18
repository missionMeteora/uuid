package uuid

import (
	crand "crypto/rand"
	"encoding/binary"
	"encoding/hex"
	"math/rand"
	"time"
	"unsafe"

	"github.com/missionMeteora/binny.v2"
	"github.com/missionMeteora/lockie"
	"github.com/missionMeteora/toolkit/errors"
)

const (
	// ErrInvalidLength is returned when a UUID is not a proper length
	ErrInvalidLength = errors.Error("invalid uuid length")
)

var gen *Gen

func init() {
	gen = NewGen()
}

// New returns a new 16 byte array with a randomized uuid
func New() (u UUID) {
	return gen.New()
}

// NewGen returns a new Gen (generator)
func NewGen() *Gen {
	var seed [8]byte
	crand.Read(seed[:])
	nseed := binary.LittleEndian.Uint64(seed[:])

	return &Gen{
		rnd: rand.New(rand.NewSource(int64(nseed))),
		mux: lockie.NewLockie(),
	}
}

// Parse will parse a byteslice as a UUID
func Parse(b []byte) (u UUID, err error) {
	if len(b) > 16 {
		err = ErrInvalidLength
		return
	}

	copy(u[:], b)
	return
}

// ParseStr will parse a string
func ParseStr(str string) (u UUID, err error) {
	var decoded []byte
	if decoded, err = hex.DecodeString(str); err != nil {
		return
	}

	return Parse(decoded)
}

// Gen is a generator of uuid's which uses it's own locking mechanism (instead of a global lock)
// Intended to be used for highly concurrent systems.
type Gen struct {
	rnd *rand.Rand
	mux lockie.Lockie
}

// New returns a new UUID
func (g *Gen) New() (u UUID) {
	ts := time.Now().UnixNano()
	copy(u[:8], ((*[8]byte)(unsafe.Pointer(&ts))[:]))
	g.mux.Lock()
	g.rnd.Read(u[8:])
	g.mux.Unlock()
	return
}

// UUID is returned when a UUID is requested. It is a 16 byte array with helper functions
type UUID [16]byte

// String returns a string representation of UUID
func (u UUID) String() string {
	return hex.EncodeToString(u[:])
}

// Short returns a string representation of 24-byte UUID
func (u UUID) Short() string {
	return hex.EncodeToString(u[:])[:24]
}

// Time returns a time representation of UUID
func (u UUID) Time() time.Time {
	return time.Unix(0, *(*int64)(unsafe.Pointer(&u[0])))
}

// MarshalBinny is used to export as binny format
func (u UUID) MarshalBinny(enc *binny.Encoder) (err error) {
	if err = enc.WriteBytes(u[:]); err != nil {
		return
	}

	enc.Flush()
	return
}

// UnmarshalBinny is used to import as binny format
func (u *UUID) UnmarshalBinny(dec *binny.Decoder) (err error) {
	var b []byte
	if b, err = dec.ReadBytes(); err == nil {
		copy(u[:], b)
	}

	return
}
