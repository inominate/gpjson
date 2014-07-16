package gpjson

import (
	"bytes"
	"encoding/json"
	"io"
)

const (
	JSON_STRING = iota
	JSON_NUMBER
	JSON_BOOL
	JSON_ARRAY
	JSON_MAP
)

type Json struct {
	Name  string
	Value interface{}

	internalError error
}

func NewJson(buf []byte) (*Json, error) {
	return NewJsonFromReader(bytes.NewReader(buf))
}

func NewJsonFromReader(r io.Reader) (*Json, error) {
	raw := map[string]interface{}{}

	dec := json.NewDecoder(r)
	dec.UseNumber()
	dec.Decode(&raw)

	newJ := Json{}
	err := newJ.marshal(raw)

	return &newJ, err
}

func (j *Json) marshal(i interface{}) error {
	switch t := i.(type) {
	case []interface{}:
		arr := []*Json{}

		for k := range t {
			nj := Json{}

			err := nj.marshal(t[k])
			if err != nil {
				return err
			}

			arr = append(arr, &nj)
		}

		j.Value = arr

	case map[string]interface{}:
		m := make(map[string]*Json)

		for k := range t {
			nj := Json{}
			nj.Name = k

			err := nj.marshal(t[k])
			if err != nil {
				return err
			}

			m[k] = &nj
		}

		j.Value = m

	default:
		j.Value = i
	}

	return nil
}
