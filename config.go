package main

import (
	"encoding/json"
	"io/ioutil"
)

type conf struct {
	Listen  string
	Backend string
}

func loadConfig(fn string) (*conf, error) {
	data, err := ioutil.ReadFile(fn)
	if err != nil {
		return nil, err
	}

	c := new(conf)
	err = json.Unmarshal(data, c)
	if err != nil {
		return nil, err
	}

	return c, nil
}
