package pseudocrypt

import (
    "testing"
    "fmt"
)

func TestToBase(t *testing.T) {
    const in, out = 10, "a"
    ps := Create()
	if x := ps.ToBase(in); x != out {
		t.Errorf("ToBase(%v) = %v, want %v", in, x, out)
	}
}

func TestToBaseLarge(t *testing.T) {
    const in, out = 99999999999999999, "7o044QyIzB"
    ps := Create()
	if x := ps.ToBase(in); x != out {
		t.Errorf("ToBase(%v) = %v, want %v", in, x, out)
	}
}

func TestFromBase(t *testing.T) {
    const in, out = "a", 10
    ps := Create()
	if x := ps.FromBase(in); x != out {
		t.Errorf("FromBase(%v) = %v, want %v", in, x, out)
	}
}

func TestFromBaseLarge(t *testing.T) {
    const in, out = "7o044QyIzB", 99999999999999999
    ps := Create()
	if x := ps.FromBase(in); x != out {
		t.Errorf("FromBase(%v) = %v, want %v", in, x, out)
	}
}

func TestHashLeadingZeros(t *testing.T) {
    const in, out = 0, "00000"
    ps := Create()
	if x := ps.Hash(in, 5); x != out {
		t.Errorf("Hash(%v) = %v, want %v", in, x, out)
	}
}

func TestHash(t *testing.T) {
    const in, out = 50, "TU8mq"
    ps := Create()
	if x := ps.Hash(in, 5); x != out {
		t.Errorf("Hash(%v) = %v, want %v", in, x, out)
	}
}

func TestUnhash(t *testing.T) {
    const in, out = "TU8mq", 50
    ps := Create()
	if x := ps.Unhash(in); x != out {
		t.Errorf("Unhash(%v) = %v, want %v", in, x, out)
	}
}

func TestHashLarge(t *testing.T) {
    const in, out = 99999999999999999, "aL6iBLdAxz"
    ps := Create()
	if x := ps.Hash(in, 10); x != out {
		t.Errorf("Hash(%v) = %v, want %v", in, x, out)
	}
}

func TestUnashLarge(t *testing.T) {
    const in, out = "aL6iBLdAxz", 99999999999999999
    ps := Create()
	if x := ps.Unhash(in); x != out {
		t.Errorf("Unhash(%v) = %v, want %v", in, x, out)
	}
}

func TestRange10(t *testing.T) {
    ps := Create()
    for i := 0; i < 10; i++ {
        hash := ps.Hash(int64(i), 5)
        fmt.Println(i, " - ", hash, " - ", ps.Unhash(hash))
    }
}