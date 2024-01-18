package main

import (
	"testing"
)

func TestEasy(t *testing.T) {
	costs := [][]int{
		{0, 98, 95, 85},
		{0, 2, 4, 2},
		{97, 0, 0, 2},
		{0, 0, 2, 0},
	}
	cost := 2
	res, err := hungarian_method(costs, "minimise")
	if cost != res || err != nil {
		t.Fatalf("%s, want %d, got %d\n", err, cost, res)
	}
}

func TestEasy2(t *testing.T) {
	costs := [][]int{
		{10, 12, 19, 11},
		{5, 10, 7, 8},
		{12, 14, 13, 11},
		{8, 15, 11, 9},
	}
	cost := 38
	res, err := hungarian_method(costs, "minimise")
	if cost != res || err != nil {
		t.Fatalf("%s, want %d, got %d\n", err, cost, res)
	}
}

func TestMedium(t *testing.T) {
	costs := [][]int{
		{1, 1, 1, 5, 3, 7, 5, 8},
		{5, 5, 5, 1, 7, 3, 9, 10},
		{2, 2, 2, 4, 4, 6, 6, 7},
		{2, 2, 2, 4, 4, 6, 6, 7},
		{4, 4, 4, 2, 6, 4, 8, 11},
		{5, 5, 5, 3, 7, 5, 9, 12},
		{6, 6, 6, 8, 8, 10, 10, 13},
		{7, 7, 7, 9, 7, 11, 9, 12},
	}
	cost := 37
	res, err := hungarian_method(costs, "minimise")
	if cost != res || err != nil {
		t.Fatalf("%s, want %d, got %d\n", err, cost, res)
	}
}

func TestHard(t *testing.T) {
	costs := [][]int{
		{1, 3, 5, 4, 2, 2, 8, 9, 11, 11, 14},
		{1, 3, 5, 4, 2, 2, 8, 9, 11, 11, 14},
		{5, 1, 1, 2, 6, 6, 4, 5, 7, 7, 10},
		{8, 6, 2, 1, 9, 9, 3, 4, 6, 10, 11},
		{3, 7, 7, 6, 2, 2, 10, 11, 13, 13, 16},
		{3, 3, 3, 2, 4, 4, 6, 7, 9, 9, 12},
		{9, 7, 3, 2, 10, 10, 4, 3, 5, 9, 10},
		{4, 6, 6, 5, 5, 5, 9, 10, 12, 12, 15},
		{4, 6, 6, 5, 5, 5, 9, 10, 12, 12, 15},
		{7, 7, 5, 4, 8, 8, 6, 7, 9, 9, 12},
		{7, 7, 5, 4, 8, 8, 6, 7, 9, 9, 12},
	}
	cost := 53
	res, err := hungarian_method(costs, "minimise")
	if cost != res || err != nil {
		t.Fatalf("%s, want %d, got %d\n", err, cost, res)
	}
}

func TestHard2(t *testing.T) {
	costs := [][]int{
		{1, 3, 3, 6, 4, 99, 5, 9, 7},
		{2, 4, 4, 5, 7, 5, 6, 6, 8},
		{2, 4, 4, 5, 7, 5, 6, 6, 8},
		{3, 99, 7, 4, 99, 99, 99, 99, 99},
		{3, 99, 5, 10, 99, 99, 99, 99, 99},
		{4, 6, 6, 9, 9, 7, 8, 10, 10},
		{4, 6, 6, 9, 9, 7, 8, 10, 10},
		{5, 99, 7, 8, 99, 99, 99, 99, 99},
		{6, 99, 8, 7, 99, 99, 99, 99, 99},
	}
	cost := 142
	res, err := hungarian_method(costs, "minimise")
	if cost != res || err != nil {
		t.Fatalf("%s, want %d, got %d\n", err, cost, res)
	}
}

func TestUnsquared(t *testing.T) {
	costs := [][]int{
		{7, 5, 11},
		{5, 4, 1},
	}
	cost := 6
	res, err := hungarian_method(costs, "minimise")
	if cost != res || err != nil {
		t.Fatalf("%s, want %d, got %d\n", err, cost, res)
	}
}

func TestUnsquared2(t *testing.T) {
	costs := [][]int{
		{94, 98, 95, 85},
		{8, 2, 4, 2},
		{97, 12, 54, 2},
		{1, 5, 2, 8},
		{16, 20, 2, 30},
	}
	cost := 7
	res, err := hungarian_method(costs, "minimise")
	if cost != res || err != nil {
		t.Fatalf("%s, want %d, got %d\n", err, cost, res)
	}
}

func TestUnsquared3(t *testing.T) {
	costs := [][]int{
		{18, 11, 16, 20},
		{14, 19, 26, 18},
		{21, 23, 35, 29},
		{32, 27, 21, 17},
		{16, 15, 28, 25},
	}
	cost := 62
	res, err := hungarian_method(costs, "minimise")
	if cost != res || err != nil {
		t.Fatalf("%s, want %d, got %d\n", err, cost, res)
	}
}

func TestUnsquared4(t *testing.T) {
	costs := [][]int{
		{12, 4, 10, 4, 12},
		{4, 10, 16, 14, 14},
		{14, 16, 12, 18, 16},
		{12, 4, 6, 8, 10},
		{18, 6, 16, 18, 14},
		{8, 14, 8, 12, 16},
	}
	cost := 32
	res, err := hungarian_method(costs, "minimise")
	if cost != res || err != nil {
		t.Fatalf("%s, want %d, got %d\n", err, cost, res)
	}
}

func TestUnsquared5(t *testing.T) {
	costs := [][]int{
		{12, 4, 10, 4, 12},
		{4, 10, 16, 14, 14},
		{14, 16, 12, 18, 16},
		{12, 4, 6, 8, 10},
	}
	cost := 24
	res, err := hungarian_method(costs, "minimise")
	if cost != res || err != nil {
		t.Fatalf("%s, want %d, got %d\n", err, cost, res)
	}
}

func TestUnsquared6(t *testing.T) {
	costs := [][]int{
		{12, 4, 14, 12},
		{4, 10, 16, 4},
		{10, 16, 12, 6},
		{4, 14, 18, 8},
		{12, 14, 16, 10},
	}
	cost := 24
	res, err := hungarian_method(costs, "minimise")
	if cost != res || err != nil {
		t.Fatalf("%s, want %d, got %d\n", err, cost, res)
	}
}

func TestUnsquared7(t *testing.T) {
	costs := [][]int{
		{12, 22, 6, 10, 10},
		{22, 24, 16, 18, 10},
		{18, 38, 26, 30, 10},
		{18, 38, 26, 30, 10},
	}
	cost := 52
	res, err := hungarian_method(costs, "minimise")
	if cost != res || err != nil {
		t.Fatalf("%s, want %d, got %d\n", err, cost, res)
	}
}
