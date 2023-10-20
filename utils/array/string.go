/*
 * Copyright (c) 2023 by EricWinn<eng.eric.winn@gmail.com>, All Rights Reserved.
 * @Author: Eric Winn
 * @Email: eng.eric.winn@gmail.com
 * @Date: 2023-10-20 19:09:55
 * @FilePath: /chinese-holiday/utils/array/string.go
 * @Software: VS Code
 */
package array

import (
	"sync"
)

type StringSet struct {
	m map[string]bool
	sync.RWMutex
}

func NewStringSet() *StringSet {
	return &StringSet{
		m: map[string]bool{},
	}
}

func (s *StringSet) Add(items ...string) {
	s.Lock()
	defer s.Unlock()
	if len(items) == 0 {
		return
	}
	for _, item := range items {
		s.m[item] = true
	}
}

func (s *StringSet) List() []string {
	s.RLock()
	defer s.RUnlock()
	list := []string{}
	for item := range s.m {
		list = append(list, item)
	}
	return list
}
