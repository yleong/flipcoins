package main

import "testing"
import "math"

var THRESHOLD = 0.0001

func TestCoinflip(t *testing.T){
    heads, tails := 0, 0
    for i := 0; i < 1000; i++ {
        if flip() == 0 {
            heads ++;
        } else {
            tails ++;
        }
    }
    percent := (float64(heads)/float64(tails))
    if math.Abs(percent - 1.0) > 0.1 {
        t.Fatalf("Percentage is not 50%%, got heads %d tails %d", heads, tails)
    }
}

type testpair struct {
    setsize int
    numflips int
}

var tests = []testpair{
	{1, 0},
	{2, 1},
	{3, 2},
	{4, 2},
	{5, 3},
	{7, 3},
	{8, 3},
	{9, 4},
	{1024, 10},
	{127, 7},
	{128, 7},
	{129, 8},
}


func TestNumflips(t *testing.T){
	for _, pair := range tests {
		actual := numflips(pair.setsize)
		if actual  != pair.numflips {
			t.Fatalf("expected numflips(%d) = %d, got %d", pair.setsize, pair.numflips, actual)
		}
	}
}

func TestCalculateUniform(t *testing.T){
	numTerms := 100
	numTests := 1000
	for i := 1; i < numTests; i++ {
		actual := calculateUniform(i, numTerms)
		expected := (1.0)/float64(i)
		if math.Abs(actual - expected) > THRESHOLD {
			t.Fatalf("expected calculateUniform(%d, %d) = %f, got %f", i, numTerms, expected, actual)
		}
	}
}


func TestTrialUniform(t *testing.T){
	numTrials := 1000000
	setSize := 1000
	ERROR_THRESH := 0.1
	hits := trialUniform(setSize, numTrials)
	for i := 0; i < setSize; i++ {
		actual := float64(hits[i]) / float64(numTrials)
		expected := 1.0/float64(setSize)
		if math.Abs(actual - expected)/expected > ERROR_THRESH {
			t.Fatalf("expected hits[%d] = %f, got %f", i, expected, actual)
		}
	}
}

