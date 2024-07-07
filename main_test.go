package main

import (
	"testing"
)

func TestPutAndGet(t *testing.T) {
	m := NewMap()

	m.Put("key1", "value1")
	m.Put(2, 42)
	m.Put(3.14, "pi")
	m.Put(true, false)

	if val, ok := m.Get("key1"); !ok || val != "value1" {
		t.Errorf("expected value1, got %v", val)
	}

	if val, ok := m.Get(2); !ok || val != 42 {
		t.Errorf("expected 42, got %v", val)
	}

	if val, ok := m.Get(3.14); !ok || val != "pi" {
		t.Errorf("expected pi, got %v", val)
	}

	if val, ok := m.Get(true); !ok || val != false {
		t.Errorf("expected false, got %v", val)
	}

	if _, ok := m.Get("key3"); ok {
		t.Errorf("expected key3 to not be found")
	}
}

func TestDelete(t *testing.T) {
	m := NewMap()

	m.Put("key1", "value1")
	m.Put(2, 42)
	m.Put(3.14, "pi")

	m.Delete("key1")
	m.Delete(2)

	if _, ok := m.Get("key1"); ok {
		t.Errorf("expected key1 to be deleted")
	}

	if _, ok := m.Get(2); ok {
		t.Errorf("expected key 2 to be deleted")
	}

	if val, ok := m.Get(3.14); !ok || val != "pi" {
		t.Errorf("expected pi to be present, got %v", val)
	}
}

func TestContains(t *testing.T) {
	m := NewMap()

	m.Put("key1", "value1")
	m.Put(2, 42)

	if !m.Contains("key1") {
		t.Errorf("expected hashmap to contain key1")
	}

	if !m.Contains(2) {
		t.Errorf("expected hashmap to contain key 2")
	}

	if m.Contains("key2") {
		t.Errorf("expected hashmap to not contain key2")
	}
}

func TestSizeAndIsEmpty(t *testing.T) {
	m := NewMap()

	if !m.IsEmpty() {
		t.Errorf("expected hashmap to be empty")
	}

	m.Put("key1", "value1")

	if m.Size() != 1 {
		t.Errorf("expected size to be 1, got %v", m.Size())
	}

	m.Put("key2", "value2")

	if m.Size() != 2 {
		t.Errorf("expected size to be 2, got %v", m.Size())
	}

	m.Delete("key1")

	if m.Size() != 1 {
		t.Errorf("expected size to be 1, got %v", m.Size())
	}

	if m.IsEmpty() {
		t.Errorf("expected hashmap to not be empty")
	}
}

func TestKeysAndValues(t *testing.T) {
	m := NewMap()

	m.Put("key1", "value1")
	m.Put(2, 42)
	m.Put(3.14, "pi")
	m.Put(true, false)

	keys := m.Keys()
	values := m.Values()

	expectedKeys := map[interface{}]bool{"key1": true, 2: true, 3.14: true, true: true}
	for _, key := range keys {
		if !expectedKeys[key] {
			t.Errorf("unexpected key %v", key)
		}
		delete(expectedKeys, key)
	}
	if len(expectedKeys) != 0 {
		t.Errorf("missing keys %v", expectedKeys)
	}

	expectedValues := map[interface{}]bool{"value1": true, 42: true, "pi": true, false: true}
	for _, value := range values {
		if !expectedValues[value] {
			t.Errorf("unexpected value %v", value)
		}
		delete(expectedValues, value)
	}
	if len(expectedValues) != 0 {
		t.Errorf("missing values %v", expectedValues)
	}
}

func TestClear(t *testing.T) {
	m := NewMap()

	m.Put("key1", "value1")
	m.Put(2, 42)

	m.Clear()

	if m.Size() != 0 {
		t.Errorf("expected size to be 0 after clear, got %v", m.Size())
	}

	if !m.IsEmpty() {
		t.Errorf("expected hashmap to be empty after clear")
	}
}
