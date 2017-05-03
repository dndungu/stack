package stack

import "testing"

func TestStack(t *testing.T) {
	s := NewStack()
	if s.head != nil {
		t.Errorf("Expected a nil head, got a non nil Stack head")
	}
}

func TestPush(t *testing.T) {
	tests := []struct {
		data     []interface{}
		expected interface{}
	}{
		{[]interface{}{10, 11, 12, 13}, 13},
		{[]interface{}{14, 15, 16, 17}, 17},
		{[]interface{}{18, 19, 10, 11}, 11},
		{[]interface{}{12, 13, 14, 15}, 15},
		{[]interface{}{16, 17, 18, 19}, 19},
	}
	for i, test := range tests {
		s := NewStack()
		for _, v := range test.data {
			s.Push(v)
		}
		actual := s.head.value
		if actual != test.expected {
			t.Errorf("%d: Error, expected, %d, got %d", i, test.expected, actual)
		}
	}
}

func TestPop(t *testing.T) {
	tests := []struct {
		data     []interface{}
		expected []interface{}
	}{
		{[]interface{}{10, 11, 12, 13}, []interface{}{13}},
		{[]interface{}{14, 15, 16, 17}, []interface{}{17, 16}},
		{[]interface{}{18, 19, 10, 11}, []interface{}{11, 10, 19}},
		{[]interface{}{12, 13, 14, 15}, []interface{}{15, 14, 13, 12}},
		{[]interface{}{16, 17, 18, 19}, []interface{}{19, 18, 17, 16, nil}},
	}
	for i, test := range tests {
		s := NewStack()
		for _, v := range test.data {
			s.Push(v)
		}
		for _, expected := range test.expected {
			actual, _ := s.Pop()
			if expected == nil {
				if actual != nil {
					t.Errorf("%d: Error, expected nil got %d", i, actual)
				}
			} else {
				if actual != expected {
					t.Errorf("%d: Error, expected, %d, got %d", i, expected, actual)
				}
			}
		}
	}
}

func TestPeek(t *testing.T) {
	tests := []struct {
		data     []interface{}
		expected interface{}
	}{
		{[]interface{}{}, nil},
		{[]interface{}{13}, 13},
		{[]interface{}{16, 17}, 17},
		{[]interface{}{19, 10, 11}, 11},
		{[]interface{}{12, 13, 14, 15}, 15},
		{[]interface{}{15, 16, 17, 18, 19}, 19},
	}
	for i, test := range tests {
		s := NewStack()
		for _, v := range test.data {
			s.Push(v)
		}
		actual, _ := s.Peek()
		if actual != test.expected {
			t.Errorf("%d: Error, expected, %d, got %d", i, test.expected, actual)
		}
	}
}
