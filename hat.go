package gohat

import (
    "strconv"
    "math"
    "math/rand"
    "time"
)

func Hat(b ...int) string {
    base := 16
    bits := 128

    if len(b) > 0 {
       base = b[0]
    }
 
    if len(b) > 1 {
       bits = b[1]
    }

    rand.Seed(time.Now().UTC().UnixNano())
    
    var digits float64 = math.Log(math.Pow(2, float64(bits))) / math.Log(float64(base))
    for i := 2; math.IsInf(digits, 0); i*=2 {
        digits = math.Log(math.Pow(2, float64(bits) / float64(i))) / math.Log(float64(base)) * float64(i);
    }

    rem := digits - math.Floor(digits);
    res := ""

    for i := 0; float64(i) < math.Floor(digits); i++ {
        x := math.Floor(rand.Float64() * float64(base));
        res = strconv.FormatInt(int64(x), base) + res;
    }

    if int64(rem) != 0 {
        b := math.Pow(float64(base), rem)
        x := math.Floor(rand.Float64() * float64(b));
        res = strconv.FormatInt(int64(x), base) + res;
    }    

    parsed,_ := strconv.ParseInt(res, base,64)

    if float64(parsed) >= math.Pow(2, float64(bits)) {
        return Hat(bits, base)
    } else {
        return res
    }
}

func Rack(b ...int) (func() string) {
    base := 16
    bits := 128
    expandBy := 16
    if len(b) > 0 {
       base = b[0]
    }

    if len(b) > 1 {
       bits = b[1]
    }

    if len(b) > 2 {
        expandBy = b[2]
    }

    hats := make(map[string]bool)
    return func() string {
        iter := 0
        id := Hat(base, bits)
        for hats[id] {
            if iter >= 10 {
                bits += expandBy
                iter = 0
            }
            id = Hat(base, bits)
            iter++
        }
        hats[id] = true
        return id
    }
}

