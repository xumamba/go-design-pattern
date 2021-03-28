package singleton

/**
 * @Time       : 2021/3/28
 * @Author     : xumamba
 * @Description: singleton_test.go
 */

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSingleton(t *testing.T) {
	singletonObjOne := GetSingletonObj()
	singletonObjTwo := GetSingletonObj()
	assert.Same(t, singletonObjOne, singletonObjTwo) // Same asserts that two pointers reference the same object.
	singletonObjThree := &singleton{}
	assert.NotSame(t, singletonObjTwo, singletonObjThree) // NotSame asserts that two pointers do not reference the same object.
}

func TestParallelSafe(t *testing.T) {
	signal := make(chan struct{})
	wg := &sync.WaitGroup{}
	maxG := 100

	wg.Add(maxG)
	container := make([]*singleton, 100)
	for i := 0; i < maxG; i++ {
		go func(index int) {
			<-signal // maxG个goroutine等待'号施令' 一起执行
			container[index] = GetSingletonObj()
			wg.Done()
		}(i)
	}
	close(signal)
	wg.Wait()
	for i := 1; i < maxG; i++ {
		assert.Same(t, container[i-1], container[i])
	}

}
