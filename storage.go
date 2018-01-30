package transaction

// Core
// Storage
// Copyright © 2017-2018 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

import (
	"sync"
)

const storageDegree uint64 = 16
const storageNumber uint64 = 1 << storageDegree
const storageShift uint64 = 64 - storageDegree

/*
storage - provides access to sections with units.
*/
type storage struct {
	data [storageNumber]*section
}

/*
newStorage - create new storage
*/
func newStorage() *storage {
	s := &storage{}
	for i := uint64(0); i < storageNumber; i++ {
		s.data[i] = newSection()
	}
	return s
}

func (s *storage) addUnit(id int64) bool {
	section := s.data[(uint64(id)<<storageShift)>>storageShift]
	return section.addUnit(id)
}

func (s *storage) getUnit(id int64) *unit {
	return s.data[(uint64(id)<<storageShift)>>storageShift].getUnit(id)
}

func (s *storage) delUnit(id int64) (*unit, bool) {
	return s.data[(uint64(id)<<storageShift)>>storageShift].delUnit(id)
}

func (s *storage) id(id int64) uint64 {
	return (uint64(id) << storageShift) >> storageShift
}

/*
section - provides access to units.
*/
type section struct {
	sync.Mutex
	data map[int64]*unit
}

/*
 newSection - create new section
*/
func newSection() *section {
	s := &section{
		data: make(map[int64]*unit),
	}
	return s
}

func (s *section) addUnit(id int64) bool {
	s.Lock()
	defer s.Unlock()
	if _, ok := s.data[id]; !ok {
		s.data[id] = newUnit()
		return true
	}
	return false
}

func (s *section) getUnit(id int64) *unit {
	if u, ok := s.data[id]; ok {
		return u
	}
	return nil
}

func (s *section) delUnit(id int64) (*unit, bool) {
	s.Lock()
	defer s.Unlock()

	if u, ok := s.data[id]; ok {
		delete(s.data, id)
		return u, true
	}
	return nil, false
}
