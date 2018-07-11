package yukpiz

import (
	"math"
	"reflect"
	"testing"
)

func TestUniqInts1(t *testing.T) {
	i := []int{1, 1, 2, 2}
	ii := UniqInts(i)
	if !reflect.DeepEqual(ii, []int{1, 2}) {
		t.Fatalf("Not match!")
	}
}

func TestUniqInts2(t *testing.T) {
	i := []int{1, 2, 3, 4, 5}
	ii := UniqInts(i)
	if !reflect.DeepEqual(ii, []int{1, 2, 3, 4, 5}) {
		t.Fatalf("Not match!")
	}
}

func TestUniqInts3(t *testing.T) {
	i := []int{}
	ii := UniqInts(i)
	if !reflect.DeepEqual(ii, []int{}) {
		t.Fatalf("Not match!")
	}
}

func TestUniqInts4(t *testing.T) {
	i := []int{-1, -1, 0, 0, 1}
	ii := UniqInts(i)
	if !reflect.DeepEqual(ii, []int{-1, 0, 1}) {
		t.Fatalf("Not match!")
	}
}

func TestUniqInts5(t *testing.T) {
	i := []int{math.MinInt64, math.MinInt64, math.MaxInt64, math.MaxInt64}
	ii := UniqInts(i)
	if !reflect.DeepEqual(ii, []int{math.MinInt64, math.MaxInt64}) {
		t.Fatalf("Not match!")
	}
}
