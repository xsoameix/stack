package stack

type Stack struct{
	Quenu []string
}

func (s *Stack) Push(value string) {
	s.Quenu = append(s.Quenu, value)
}

func (s *Stack) Pop() string {
	pop := s.Quenu[len(s.Quenu)-1]
	s.Quenu = s.Quenu[:len(s.Quenu)-1]
	return pop
}

func (s *Stack) Top() string {
	return s.Quenu[len(s.Quenu)-1]
}

func (s *Stack) String() string {
	info := "["
	for i, value := range s.Quenu {
		info += value+", "
		if i == len(s.Quenu) - 1 {
			info = info[:len(info)-len(", ")]
		}
	}
	return info + "]"
}
