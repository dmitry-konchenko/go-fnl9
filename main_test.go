package main

import (
	"fmt"
	"math/rand"
	"slices"
	"testing"
)

// Функция проверяет размер слайса и положительные значения
func TestGenerateRandomElements(t *testing.T) {
	requests := []struct {
		size int
		want int
	}{
		{-5, 0},
		{0, 0},
		{1, 1},
		{1000, 1000},
	}

	for _, v := range requests {
		t.Run(fmt.Sprintf("size=%d", v.size), func(t *testing.T) {
			data := generateRandomElements(v.size)
			if len(data) != v.want {
				t.Fatalf("Ожидалось %d элементов, получено %d", v.want, len(data))
			}

			for i, value := range data {
				if value <= 0 {
					t.Errorf("Элемент с индексом %d должен быть положительным, получено %d", i, value)
				}
			}
		})
	}
}

// Функция проверяет оба способа поиска максимума
func TestMaximum(t *testing.T) {
	requests := []struct {
		name string
		data []int
		want int
	}{
		{"nil", nil, 0},
		{"empty", []int{}, 0},
		{"single", []int{7}, 7},
		{"short", []int{2, 9, 3}, 9},
		{"first", []int{20, 2, 3, 4, 5, 6, 7, 8}, 20},
		{"middle", []int{1, 2, 3, 20, 5, 6, 7, 8}, 20},
		{"last", []int{1, 2, 3, 4, 5, 6, 7, 20}, 20},
		{"remainder", []int{1, 2, 3, 4, 5, 6, 7, 8, 50, 9}, 50},
		{"equal", []int{5, 5, 5, 5, 5, 5, 5, 5}, 5},
		{"zero", []int{0, 0, 0}, 0},
		{"negative", []int{-9, -8, -7, -6, -5, -4, -3, -2, -1}, -1},
		{"mixed", []int{-10, 0, 6, -4, 2, 5, -3, 1}, 6},
		{"min_int", []int{-int(^uint(0)>>1) - 1}, -int(^uint(0)>>1) - 1},
		{"max_int", []int{1, int(^uint(0) >> 1)}, int(^uint(0) >> 1)},
	}

	functions := []struct {
		name string
		find func([]int) int
	}{
		{"maximum", maximum},
		{"maxChunks", maxChunks},
	}

	for _, function := range functions {
		t.Run(function.name, func(t *testing.T) {
			for _, v := range requests {
				t.Run(v.name, func(t *testing.T) {
					original := slices.Clone(v.data)
					if got := function.find(v.data); got != v.want {
						t.Errorf("Ожидался максимум %d, получено %d", v.want, got)
					}
					if !slices.Equal(original, v.data) {
						t.Error("Исходный слайс изменился")
					}
				})
			}
		})
	}
}

// Функция проверяет границы частей при разной длине слайса
func TestMaxChunksBoundaries(t *testing.T) {
	for _, size := range []int{8, 9, 15, 16, 17, 23, 24, 25} {
		t.Run(fmt.Sprintf("size=%d", size), func(t *testing.T) {
			data := make([]int, size)
			for i := range data {
				data[i] = 100
				if got := maxChunks(data); got != 100 {
					t.Fatalf("Максимум в позиции %d: ожидалось 100, получено %d", i, got)
				}
				data[i] = 0
			}
		})
	}
}

// Функция сравнивает поиск по частям со стандартным поиском
func TestMaxChunksRandom(t *testing.T) {
	rnd := rand.New(rand.NewSource(1))
	for _, size := range []int{7, 8, 9, 31, 1000, 100003} {
		t.Run(fmt.Sprintf("size=%d", size), func(t *testing.T) {
			data := make([]int, size)
			for i := range data {
				data[i] = rnd.Intn(2001) - 1000
			}
			want := slices.Max(data)
			if got := maxChunks(data); got != want {
				t.Errorf("Ожидался максимум %d, получено %d", want, got)
			}
		})
	}
}
