package algorithm

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShellSort(t *testing.T) {
	arr := []int{3, 2, 0, 5, 10, 1, 13}
	arrSort := ShellSort(arr)
	fmt.Printf("arr after sored: %v\n", arrSort)
	assert.Equal(t, 10, arrSort[5])
	assert.Equal(t, 5, arrSort[4])
}
