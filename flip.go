 /*
 Can we randomly choose 1 item out of n by flipping m = ceil(log_2(n)) coins to
 generate m bits and convert the bits to base 10 to get our item number l. If l
 > n  then we will just discard and try again.  Question is, is this uniform?
 */
package main

import ("fmt";"math/rand"; "math"; "time"; "flag")

func flip() int {
    return rand.Intn(2)
}

/*
 Given an integer n, returns the least number of flips needed to generate a
 bitstring that covers at least until n. E.g., for n = 7, we have to flip 3
 times for n = 30, we have to flip 5 times
*/
func numflips(set_size int) int {
    log := math.Log2(float64(set_size))
    return int(math.Pow(2, math.Ceil(log)))
}

/*
 Define m to be the number of flips needed to cover n. The chance that any
 1 of the n items getting chosen on first try is 1/m, on second try is
 (m-n)/m*1/m, on third try is (m-n)/m*(m-n)/m*1/m. So we will just sum this up
 Or we can probably use the geometric progression formula
*/
func calculateUniform(set_size int, num_terms int) {
   totalItems := numflips(set_size)
   multiplier := float64(totalItems - set_size)/float64(totalItems)
   sum := 1.0/float64(totalItems)
   for i, term := 0, sum; i < num_terms; i++ {
        term *= multiplier
        sum += term
   }
   fmt.Printf("Probability of choosing 1 item out of %d: %f\n", set_size, sum)
}

/*
 Alternatively we can just simulate the flipping for a large number of trials
 and check that each item do get selected more or less with equal frequency in
 the end
*/
func trialUniform(set_size int, num_trials int) {
    num_flips := numflips(set_size)
    hits := make([]int, set_size)
    for i := 0; i < num_trials; i++ {
        currNum := 0
        currMultiplier := 1
        for k :=0; k < num_flips; k++{
            currNum += flip() * currMultiplier
            currMultiplier *= 2
        }
        if currNum >= set_size {
            //discard this trial
            i--
            continue
        }
        hits[currNum]++
    }
    for i := 0; i < set_size; i++ {
        fmt.Printf("Number of hits for %d : %d\n", i, hits[i]);
    }
}

func testPowerOfTwo(){
    for i := 0; i < 100; i++ {
        fmt.Printf("the smallestpoweroftwo of %d is %d\n", i, numflips(i))
    }
}
func testCalculateUniform(){
    for i := 0; i < 100; i++ {
        calculateUniform(i, 100)
    }
}
func getOpts() (int, int, int) {
    setsize := flag.Int("n", 3, "number of items in a set")
    numtrials := flag.Int("t", 1000, "number of trials to simulate")
    numterms := flag.Int("m", 100, "number of terms to sum")
    flag.Parse()
    return *setsize, *numtrials, *numterms
}

func main() {
    rand.Seed(time.Now().UTC().UnixNano())
    setsize, numtrials, numterms := getOpts()

    calculateUniform(setsize, numterms)
    trialUniform(setsize, numtrials)
}
