package main

import (
	"reflect"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	a := ListNode{Val: 1}
	b := ListNode{Val: 2, Next: &a}
	c := ListNode{Val: 5, Next: &b}
	ans := addTwoNumbers(&c, &c)

	actual := []int{}
	for ans != nil {
		actual = append(actual, ans.Val)
		ans = ans.Next
	}
	expect := []int{0, 5, 2}

	if !reflect.DeepEqual(actual, expect) {
		t.Errorf("result error")
	}
}