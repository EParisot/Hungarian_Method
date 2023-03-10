package main

import (
	"testing"
)

func TestUnsquared(t *testing.T) {
	costs := [][]int{
		{0, 98, 95, 85},
		{0, 2, 4, 2},
		{97, 0, 0, 2},
		{0, 0, 2, 0},
		{0, 0, 2, 0},
	}
	_, err := hungarian_method(costs)
	if err == nil {
		t.Fatalf("want %s\n", err)
	}
}

func TestUnsquared2(t *testing.T) {
	costs := [][]int{
		{0, 98, 95, 85},
		{0, 2, 4, 2},
		{97, 0, 0, 2},
		{0, 0, 2, 0, 0},
	}
	_, err := hungarian_method(costs)
	if err == nil {
		t.Fatalf("want %s\n", err)
	}
}

func TestEasy(t *testing.T) {
	costs := [][]int{
		{0, 98, 95, 85},
		{0, 2, 4, 2},
		{97, 0, 0, 2},
		{0, 0, 2, 0},
	}
	cost := 2
	res, err := hungarian_method(costs)
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
	res, err := hungarian_method(costs)
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
	res, err := hungarian_method(costs)
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
	res, err := hungarian_method(costs)
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
	res, err := hungarian_method(costs)
	if cost != res || err != nil {
		t.Fatalf("%s, want %d, got %d\n", err, cost, res)
	}
}
