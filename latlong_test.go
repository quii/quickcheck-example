package main

import (
"reflect"
"math/rand"
"testing"
	"testing/quick"
)

type LatLong struct{
	Lat float64
	Long float64
}

func (l LatLong) Generate(rand *rand.Rand, size int) reflect.Value{
	randomLatLong := LatLong{
		Lat: rand.Float64(),
		Long: rand.Float64(),
	}
	return reflect.ValueOf(randomLatLong)
}

func TestLatLongIsAlwaysTrue(t *testing.T){
	assertion := func(x LatLong) bool {
		t.Log("Random LatLong:", x)
		//todo: Do some interesting assertions!
		return true
	}

	if err := quick.Check(assertion, nil); err != nil {
		t.Error(err)
	}
}