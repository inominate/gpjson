package gpjson

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
)

var ErrNotMap = fmt.Errorf("value is not a map")
var ErrNotArray = fmt.Errorf("value is not an array")

var ErrNameNotFound = fmt.Errorf("name not found in map")
var ErrIdxRange = fmt.Errorf("array index out of range")

func (j *Json) Map() (map[string]*Json, error) {
	if j.internalError != nil {
		return nil, j.internalError
	}

	m, ok := j.Value.(map[string]*Json)
	if !ok {
		return nil, ErrNotMap
	}

	return m, nil
}

func (j *Json) Slice() ([]*Json, error) {
	if j.internalError != nil {
		return nil, j.internalError
	}

	arr, ok := j.Value.([]*Json)
	if !ok {
		return nil, ErrNotArray
	}

	return arr, nil
}

func (j *Json) Get(name string) *Json {
	m, err := j.Map()
	if err != nil {
		return &Json{internalError: err}
	}

	v, exists := m[name]
	if !exists {
		return &Json{internalError: ErrNameNotFound}
	}

	return v
}

func (j *Json) Idx(idx int) *Json {
	arr, err := j.Slice()
	if err != nil {
		return &Json{internalError: err}
	}

	if idx < 0 || idx+1 > len(arr) {
		return &Json{internalError: ErrIdxRange}
	}

	return arr[idx]
}

/*
func (j *Json) () (, error) {
	if j.internalError != nil {
		return nil, j.internalError
	}

}
*/
func (j *Json) String() (string, error) {
	if j.internalError != nil {
		return "", j.internalError
	}

	if v, ok := j.Value.(string); ok {
		return v, nil
	}

	if v, ok := j.Value.(json.Number); ok {
		return v.String(), nil
	}

	if v, ok := j.Value.(bool); ok {
		return fmt.Sprint(v), nil
	}

	return "", errors.New("cannot stringify this element")
}

func (j *Json) Bool() (bool, error) {
	if j.internalError != nil {
		return false, j.internalError
	}

	v, ok := j.Value.(bool)
	if !ok {
		return false, errors.New("not a bool")
	}
	return v, nil
}

func (j *Json) Int64() (int64, error) {
	if j.internalError != nil {
		return 0, j.internalError
	}

	istr, err := j.String()
	if err != nil {
		return 0, err
	}

	return strconv.ParseInt(istr, 10, 64)
}

func (j *Json) Float64() (float64, error) {
	if j.internalError != nil {
		return 0, j.internalError
	}

	fstr, err := j.String()
	if err != nil {
		return 0, err
	}

	return strconv.ParseFloat(fstr, 64)
}
