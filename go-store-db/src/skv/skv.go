package skv

import (
	"encoding/json"
	"io/ioutil"
)

func objToString(obj interface{}) string {
	bolB, _ := json.Marshal(obj)

	return string(bolB)
}

func storeFromSource(path string) (map[string]interface{}, error) {
	rawData, ioErr := ioutil.ReadFile(path)

	if ioErr != nil {
		return nil, ioErr
	}

	var obj map[string]interface{} = nil
	var jsonErr = json.Unmarshal(rawData, &obj)

	if jsonErr != nil {
		return nil, jsonErr
	}

	return obj, nil
}

type skv struct {
	path  string
	store map[string]interface{}
}

func (s skv) saveToSource() {
	bolB, err := json.Marshal(s.store)

	if err == nil {
		ioutil.WriteFile(s.path, bolB, 0644)
	}
}

func (s skv) Put(key string, value interface{}) {
	s.store[key] = objToString(value)
	s.saveToSource()
}

func (s skv) Delete(key string) {
	delete(s.store, key)
	s.saveToSource()
}

func (s skv) Get(key string) interface{} {
	return s.store[key]
}

func Open(path string) (s *skv, err error) {
	store, err := storeFromSource(path)

	if err != nil {
		return nil, err
	}

	return &skv{path, store}, nil
}
