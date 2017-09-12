 /*
 Can we randomly choose 1 item out of 3 by flipping 2 coins to generate 2 bits
 and convert the bits to base 10 to get our item number. If we flipped a 4 then
 we will just discard and try again.  Question is, is this uniform?
 */

package main

import ("fmt";"math/rand"; "time")

func flip() int {
    return rand.Intn(2)
}

/*
 The chance that any 1 of the 3 items getting chosen on first try is 1/4, on
 second try is 1/4*1/4, on third try is 1/4*1/4*1/4. So we will just sum this up
 Or we can probably use the geometric progression formula
*/
func calculateUniform() {
   sum := 0.25
   num_terms := 100
   for i, term := 0, sum; i < num_terms; i++ {
        term *= 0.25 
        sum += term
   }
   fmt.Printf("Probability of choosing 1 item out of 3: %f\n", sum)
}

/*
 Alternatively we can just simulate the flipping for a large number of trials
 and check that we do get an even distribution in the end
*/
func trialUniform() {
    num_trials := 10000
    num_flips := 2
    const set_size = 3
    var hits [set_size]int
    for i := 0; i < num_trials; i++ {
        currNum := 0
        currMultiplier := 1
        for k :=0; k < num_flips; k++{
            currNum += flip() * currMultiplier
            currMultiplier *= 2
        }
        if currNum == 3 {
            i--
            continue
        }
        hits[currNum]++
    }
    for i := 0; i < set_size; i++ {
        fmt.Printf("Number of hits for %d : %d\n", i, hits[i]);
    }
}

func main() {
    rand.Seed(time.Now().UTC().UnixNano())
    calculateUniform()
    trialUniform()
}
