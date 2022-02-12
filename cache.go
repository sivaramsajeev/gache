package main

import (
	"container/list"
	"errors"
	"log"
	"sync"
)

type Cache struct {
	store Store
	list  *list.List
	size  int
}

type Store interface {
	Get(param interface{}) (interface{}, error)
	Set(key, val interface{}) error
	Del(key interface{}) error
}

func New(size int) *Cache {
	if size < 1 {
		log.Fatal("Cache size should be greater than 0")
	}
	return &Cache{
		size:  size,
		list:  list.New(),
		store: NewStore(),
	}
}

func (c *Cache) Update(key interface{}) {
	for e := c.list.Front(); e != nil; e = e.Next() {
		if e.Value == key {
			c.list.Remove(e)
		}
	}
	c.list.PushBack(key)
	for c.list.Len() > c.size {
		front := c.list.Front()
		c.list.Remove(front)
		c.Delete(front.Value)
	}
}

type InMemory struct {
	sync.RWMutex
	dict map[interface{}]interface{}
}

func NewStore() *InMemory {
	return &InMemory{
		dict: make(map[interface{}]interface{}),
	}
}

func (m *InMemory) Get(param interface{}) (interface{}, error) {
	m.RLock()
	defer m.RUnlock()
	val, ok := m.dict[param]
	if !ok {
		return nil, errors.New("value not found")
	}
	return val, nil
}

func (m *InMemory) Set(key, val interface{}) error {
	m.Lock()
	defer m.Unlock()
	m.dict[key] = val
	return nil
}

func (m *InMemory) Del(key interface{}) error {
	m.Lock()
	defer m.Unlock()
	delete(m.dict, key)
	return nil
}

func (c *Cache) Get(key interface{}) (interface{}, error) {
	val, err := c.store.Get(key)
	if err != nil {
		return nil, err
	}
	c.Update(key)
	return val, nil
}

func (c *Cache) Set(key, val interface{}) error {
	err := c.store.Set(key, val)
	if err != nil {
		return err
	}
	c.Update(key)
	return nil
}

func (c *Cache) Delete(key interface{}) error {
	return c.store.Del(key)
}
