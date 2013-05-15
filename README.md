PseudoCrypt
===========

A small library for creating reversible obfuscated identifiers for integer values.

```go
ps := Create()
for i := 0; i < 10; i++ {
    hash := ps.Hash(int64(i), 5)
    fmt.Println(i, " - ", hash, " - ", ps.Unhash(hash))
}
```

```
0  -  00000  -  0
1  -  CjIO3  -  1
2  -  eDrC6  -  2
3  -  QXaq9  -  3
4  -  tgTec  -  4
5  -  5AC2f  -  5
6  -  HUkQi  -  6
7  -  ke3El  -  7
8  -  WxMso  -  8
9  -  yRvgr  -  9
```

Classic use case: URL shorteners where:
* The long url is stored in a database with an integer primary key
* The short url must be unique
* Valid short urls must be difficult to deduce

Adapted from: http://blog.kevburnsjr.com/php-unique-hash