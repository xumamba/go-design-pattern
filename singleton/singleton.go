package singleton

/**
 * @Time       : 2021/3/28
 * @Author     : xumamba
 * @Description: 单例模式，采用懒加载方式，线程安全。
 **/

import (
	"sync"
)

var obj *singleton
var once sync.Once

// singleton 考虑开闭原则，封闭结构体本身对外的可见性。
type singleton struct {
	name string // 声明包含属性的结构体（Go种所有的空结构体对象都指向相同的内存空间）
}

// GetSingletonObj 获取全局唯一的实例化对象
func GetSingletonObj() *singleton {
	once.Do(func() {
		obj = &singleton{}
	})
	return obj
}
