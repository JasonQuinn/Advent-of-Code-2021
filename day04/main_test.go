package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetBingo(t *testing.T) {
	input := []string{"7,4,9,5,11,17,23,2,0,14,21,24,10,16,13,6,15,25,12,22,18,20,8,19,3,26,1", "", "22 13 17 11  0", "8  2 23  4 24", "21  9 14 16  7", "6 10  3 18  5", "1 12 20 15 19", "", "3 15  0  2 22", "9 18 13 17  5", "19  8  7 25 23", "20 11 10 24  4", "14 21 16 12  6", "", "14 21 17 24  4", "10 16 15  9 19", "18  8 23 26 20", "22 11 13  6  5", "2  0 12  3  7"}
	bingo, err := getNumbersAndCards(input)
	assert.Nil(t, err)
	assert.Equal(t, 27, len(bingo.Numbers))
	assert.Equal(t, 3, len(bingo.Cards))
	assert.Equal(t, 5, len(bingo.Cards[0]))
	assert.ElementsMatch(t, []int{22, 13, 17, 11, 0}, bingo.Cards[0][0])
	assert.ElementsMatch(t, []int{8, 2, 23, 4, 24}, bingo.Cards[0][1])
	assert.ElementsMatch(t, []int{21, 9, 14, 16, 7}, bingo.Cards[0][2])
	assert.ElementsMatch(t, []int{6, 10, 3, 18, 5}, bingo.Cards[0][3])
	assert.ElementsMatch(t, []int{1, 12, 20, 15, 19}, bingo.Cards[0][4])
	assert.Equal(t, 5, len(bingo.Cards[1]))
	assert.ElementsMatch(t, []int{3, 15, 0, 2, 22}, bingo.Cards[1][0])
	assert.Equal(t, 5, len(bingo.Cards[2]))
}

func TestPart1(t *testing.T) {
	input := []string{"7,4,9,5,11,17,23,2,0,14,21,24,10,16,13,6,15,25,12,22,18,20,8,19,3,26,1", "", "22 13 17 11  0", "8  2 23  4 24", "21  9 14 16  7", "6 10  3 18  5", "1 12 20 15 19", "", "3 15  0  2 22", "9 18 13 17  5", "19  8  7 25 23", "20 11 10 24  4", "14 21 16 12  6", "", "14 21 17 24  4", "10 16 15  9 19", "18  8 23 26 20", "22 11 13  6  5", "2  0 12  3  7"}
	bingo, err := getNumbersAndCards(input)
	assert.Nil(t, err)

	result := part1(bingo)
	assert.Equal(t, 4512, result)
}

func TestPart2(t *testing.T) {
	input := []string{"7,4,9,5,11,17,23,2,0,14,21,24,10,16,13,6,15,25,12,22,18,20,8,19,3,26,1", "", "22 13 17 11  0", "8  2 23  4 24", "21  9 14 16  7", "6 10  3 18  5", "1 12 20 15 19", "", "3 15  0  2 22", "9 18 13 17  5", "19  8  7 25 23", "20 11 10 24  4", "14 21 16 12  6", "", "14 21 17 24  4", "10 16 15  9 19", "18  8 23 26 20", "22 11 13  6  5", "2  0 12  3  7"}
	bingo, err := getNumbersAndCards(input)
	assert.Nil(t, err)

	result := part2(bingo)
	assert.Equal(t, 1924, result)
}
