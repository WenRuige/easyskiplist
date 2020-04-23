package skiplist

import (
	"fmt"
	"math/rand"
	"sync"
)

const (
	MaxLevel    int     = 16 //可以容纳的大小为 2 ^ 32
	Probability float32 = 0.5
)

// 创建一个新的skiplist
func New() *SkipList {
	return &SkipList{
		Head:  &Element{Forward: make([]*Element, MaxLevel)},
		Level: 0,
	}
}

// 随机的层
func randomLevel() int {
	level := 1
	// 根据概率生成一个Probability 并且要求 level < MaxLevel
	for rand.Float32() < Probability && level < MaxLevel {
		level++
	}
	return level
}

// 跳表
type SkipList struct {
	Level int      // 跳跃表的层数
	Len   int      // 当前元素的个数
	Head  *Element // 每一层的数据
	Mutex sync.RWMutex
}

// 元素
type Element struct {
	Forward []*Element  // 存储向前的指针列表
	Key     int         // 元素的key
	Value   interface{} //元素的value
}

// 搜索  17
func (s *SkipList) Search(key int) (element *Element, ok bool) {
	x := s.Head
	// 1.从最高层开始搜索
	for i := s.Level - 1; i >= 0; i-- {
		// 2.看这一层的元素是否有小于这个key的
		for x.Forward[i] != nil && x.Forward[i].Key < key {
			x = x.Forward[i]
		}
		// 3.如果没有就遍历下一层
	}
	x = x.Forward[0]
	if x != nil && x.Key == key {
		return x, true
	}
	return nil, false
}

// 创建一个新的元素
func newElement(key int, value interface{}, level int) *Element {
	return &Element{
		Key:     key,
		Value:   value,
		Forward: make([]*Element, level),
	}
}

// 插入一个新的元素
func (s *SkipList) Insert(key int, value interface{}) *Element {
	// 1.加读写锁
	s.Mutex.RLock()
	defer s.Mutex.RUnlock()
	// 创建一个新的element
	update := make([]*Element, MaxLevel)
	x := s.Head
	for i := s.Level - 1; i >= 0; i-- {
		for x.Forward[i] != nil && x.Forward[i].Key < key {
			x = x.Forward[i]
		}
		update[i] = x
	}
	x = x.Forward[0]
	// 如果当前x!=nil && x.key = key
	if x != nil && x.Key == key {
		x.Value = value
		return x
	}
	// 随机生成一个level
	level := randomLevel()
	// 如果生成的level大于当前level
	if level > s.Level {
		level = s.Level + 1
		update[s.Level] = s.Head
		s.Level = level
	}
	// 创建一个元素，指定level的层数
	e := newElement(key, value, level)
	for i := 0; i < level; i++ {
		e.Forward[i] = update[i].Forward[i]
		update[i].Forward[i] = e
	}
	s.Len++
	return e
}

// 删除节点
func (s *SkipList) Delete(key int) *Element {
	// 1.加读写锁
	s.Mutex.RLock()
	defer s.Mutex.RUnlock()
	// 2.找到key对应
	update := make([]*Element, MaxLevel)
	x := s.Head
	for i := s.Level - 1; i >= 0; i-- {
		for x.Forward[i] != nil && x.Forward[i].Key < key {
			x = x.Forward[i]
		}
		update[i] = x
	}
	x = x.Forward[0]
	// 3.如果找到这个key
	if x != nil && x.Key == key {
		// 4.从最上层开始移除
		for i := 0; i < s.Level; i++ {
			if update[i].Forward[i] != x {
				return nil
			}
			// 当前层的右侧节点
			update[i].Forward[i] = x.Forward[i]
		}
		s.Len--
	}
	return x
}

// 返回skiplist的长度
func (s *SkipList) Length() int {
	return s.Len
}

func (s *SkipList) DumpSkipList() {
	x := s.Head
	// 1.从最高层开始搜索
	for i := s.Level - 1; i >= 0; i-- {
		x = s.Head
		// 2.看这一层的元素是否有小于这个key的
		for x.Forward[i] != nil {
			if x.Forward[i].Forward[i] == nil {
				fmt.Print(x.Forward[i].Key)
			} else {
				fmt.Print(x.Forward[i].Key, "-->")
			}
			x = x.Forward[i]
		}
		fmt.Println("")
	}

}
