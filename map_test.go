package belajar_golang_goroutines

import (
	"sync"
)

func AddToMap(data *sync.Map, value int, group *sync.WaitGroup) {
	defer group.Done()

	group.Add(1)
	data.Store(value, value)
}

//func TestMap() {
//
//	data := &sync.Map{}
//	group := &sync.WaitGroup{}
//
//	for i := 0; i < 100; i++ {
//		go AddToMap(data, i, group)
//	}
//
//	group.Wait()
//
//	data.Range(func(key, value interface{}) bool {
//		fmt.Println(key, ":", value)
//		return true
//	})
//}
