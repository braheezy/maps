package main

import (
	"fmt"
	"hash/fnv"
)

const initialCapacity = 16

type Map struct {
	buckets []*entry
	size    int
}

type entry struct {
	key   interface{}
	value interface{}
	next  *entry
}

func NewMap() Map {
	return Map{
		buckets: make([]*entry, initialCapacity),
	}
}

func (m *Map) hash(key interface{}) int {
	return int(fnv32(key)) % len(m.buckets)
}

func fnv32(key interface{}) uint32 {
	h := fnv.New32a()
	h.Write([]byte(fmt.Sprintf("%v", key)))
	return h.Sum32()
}

func (m *Map) Put(key, value interface{}) {
	index := m.hash(key)
	head := m.buckets[index]

	// Collision resolution
	for e := head; e != nil; e = e.next {
		if e.key == key {
			e.value = value
			return
		}
	}

	newEntry := &entry{key: key, value: value, next: head}
	m.buckets[index] = newEntry
	m.size++
}

func (m *Map) Get(key interface{}) (interface{}, bool) {
	index := m.hash(key)
	head := m.buckets[index]

	for e := head; e != nil; e = e.next {
		if e.key == key {
			return e.value, true
		}
	}

	return nil, false
}
