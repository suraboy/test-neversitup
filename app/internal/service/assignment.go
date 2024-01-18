package service

import (
	"errors"
	"fmt"
	"regexp"
	"sort"
)

type AssignTestService interface {
	GeneratePermutations(input string) []string
	FindOdd(array []int) (int, error)
	CheckCountSmileys(array []string) int
}

type service struct{}

func NewService() AssignTestService {
	return &service{}
}

func (s service) GeneratePermutations(input string) []string {
	result := make([]string, 0)
	used := make(map[int]bool)
	generatePermutationsHelper(input, "", used, &result)
	sort.Strings(result)
	return result
}

func generatePermutationsHelper(input, current string, used map[int]bool, result *[]string) {
	if len(current) == len(input) {
		*result = append(*result, current)
		return
	}

	for i := 0; i < len(input); i++ {
		if !used[i] {
			used[i] = true
			generatePermutationsHelper(input, current+string(input[i]), used, result)
			used[i] = false
		}
	}
}

func (s service) FindOdd(array []int) (int, error) {
	occurrences := make(map[int]int)

	for _, num := range array {
		occurrences[num]++
	}

	for num, count := range occurrences {
		if count%2 != 0 {
			return num, nil
		}
	}

	return 0, errors.New(fmt.Sprintf("Number not found in Odd : %d\n", array))
}

func (s service) CheckCountSmileys(arr []string) int {
	regx := `[:;][-~]?[)D]`
	re := regexp.MustCompile(regx)

	count := 0
	for _, face := range arr {
		if re.MatchString(face) {
			count++
		}
	}

	return count
}
