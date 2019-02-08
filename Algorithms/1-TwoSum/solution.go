func twoSum(nums []int, target int) []int {
	ss := &SliceStruct{nums}
	last_index := len(ss.Body) - 1
	find_index := ss.Find(target - ss.Pop())
	if find_index != -1 {
		return []int{find_index, last_index}
	} else {
		if len(ss.Body) > 0 {
			return twoSum(ss.Body, target)
		} else {
			return []int{}
		}
	}
}

type SliceStruct struct {
	Body []int
}

func (s SliceStruct) Find(val int) int {
	for i, e := range s.Body {
		if e == val {
			return i
		}
	}
	return -1
}

func (s *SliceStruct) Pop() int {
	pop := s.Body[len(s.Body)-1]
	s.Body = s.Body[:len(s.Body)-1]
	return pop
}

func Find(arr []int, val int) int {
	for i, e := range arr {
		if e == val {
			return i
		}
	}
	return -1
}

func Pop(arr []int) ([]int, int) {
	pop := arr[len(arr)-1]
	arr = arr[:len(arr)-1]
	return arr, pop
}