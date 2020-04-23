package skiplist

import (
	"testing"
)

func TestNew(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		s := New()
		s.Insert(1, "tester")
		s.Insert(2, "tester")
		s.Insert(3, "tester")
		s.Insert(4, "tester")
		s.Insert(5, "tester")
		s.Insert(6, "tester")
		s.Insert(7, "tester")
		s.Insert(8, "tester")
		s.Insert(9, "tester")
		s.Insert(10, "tester")
		s.Insert(11, "tester")
		s.Insert(12, "tester")

		//result,ok := s.Search(1)
		//if ok {
		//	t.Log(result)
		//}
		//fmt.Println(s.Head.Forward[0])
		//fmt.Println(s.Head.Forward[0].Forward[0])
		//fmt.Println(s.Head.Forward[0].Forward[0].Forward[0])
		//fmt.Println(s.Head.Forward[0].Forward[0].Forward[0].Forward[0])
		//fmt.Println("---------------------")
		//fmt.Println(s.Head.Forward[4])
		//fmt.Println(s.Head.Forward[0].Forward[0])

		s.DumpSkipList()
		//x:=s.Head
		//for x.Forward[0] != nil  {
		//	fmt.Println(x.Key)
		//	x = x.Forward[0]
		//}

	})
	t.Run("2", func(t *testing.T) {
		
	})
}
