package pseudocrypt

import (
    "math"
    "math/big"
    "strings"
)

type PseudoCrypt struct {
    Chars string
    Primes []int64
    ModMulInv []int64
}

func Create () *PseudoCrypt {
    return &PseudoCrypt {
        /* Base 62 (ASCII ONLY) */
        Chars: "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ",
        /* Next prime greater than 62 ^ n / 1.618033988749894848 */
        Primes: []int64{
            1,
            41,
            2377,
            147299,
            9132313,
            566201239,
            35104476161,
            2176477521929,
            134941606358731,
            8366379594239857,
            518715534842869223,
        },
        /* Modular multiplicative inverse */
        ModMulInv: []int64{
            1,
            59,
            1677,
            187507,
            5952585,
            643566407,
            22071637057,
            294289236153,
            88879354792675,
            7275288500431249,
            280042546585394647,
        },
    }
}

func (ps *PseudoCrypt) ToBase(n int64) (key string) {
    base := int64(len(ps.Chars))
    for n > 0 {
        mod := n % base
        key = ps.Chars[mod:mod+1] + key
        n = n / base
    }
    return key
}

func (ps *PseudoCrypt) FromBase(key string) (n int64) {
    base := int64(len(ps.Chars))
    for i := int64(0) ; len(key) > 0 ; i++ {
        c := key[len(key)-1:]
        key = key[:len(key)-1]
        dec := int64(strings.Index(ps.Chars, c))
        n += dec * int64(math.Pow(float64(base), float64(i)))
    }
    return n
}

func (ps *PseudoCrypt) Hash(n int64, length int) (hash string) {
    base := int64(len(ps.Chars))
    ceil := int64(math.Pow(float64(base), float64(length)))
    prime := ps.Primes[length]
    dec := big.NewInt(0).Mod(big.NewInt(0).Mul(big.NewInt(n), big.NewInt(prime)), big.NewInt(ceil))
    hash = ps.ToBase(dec.Int64());
    if len(hash) < length {
        hash = strings.Repeat("0", length-len(hash)) + hash;
    }
    return hash
}

func (ps *PseudoCrypt) Unhash(key string) (n int64) {
    length := len(key)
    base := int64(len(ps.Chars))
    ceil := int64(math.Pow(float64(base), float64(length)))
    mmi := ps.ModMulInv[length]
    n = ps.FromBase(key)
    dec := big.NewInt(0).Mod(big.NewInt(0).Mul(big.NewInt(n), big.NewInt(mmi)), big.NewInt(ceil))
    return dec.Int64()
}