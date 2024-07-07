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

func (m *Map) Delete(key interface{}) {
	index := m.hash(key)
	head := m.buckets[index]

	var prev *entry
	for e := head; e != nil; e = e.next {
		if e.key == key {
			if prev == nil {
				// Entry to delete is the head
				m.buckets[index] = e.next
			} else {
				// Entry to delete is in the middle or end of list
				prev.next = e.next
			}
			m.size--
			return
		}
		prev = e
	}
}

func (m *Map) Contains(key interface{}) bool {
	index := m.hash(key)
	head := m.buckets[index]

	for e := head; e != nil; e = e.next {
		if e.key == key {
			return true
		}
	}

	return false
}

func (m *Map) Size() int {
	return m.size
}

func (m *Map) IsEmpty() bool {
	return m.size == 0
}

func (m *Map) Keys() []interface{} {
	var keys []interface{}
	for _, bucket := range m.buckets {
		for e := bucket; e != nil; e = e.next {
			keys = append(keys, e.key)
		}

	}
	return keys
}

func (m *Map) Values() []interface{} {
	var values []interface{}
	for _, bucket := range m.buckets {
		for e := bucket; e != nil; e = e.next {
			values = append(values, e.value)
		}

	}
	return values
}

func (m *Map) Clear() {
	for i := range m.buckets {
		m.buckets[i] = nil
	}
	m.size = 0
}
