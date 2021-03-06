// generated by jsonenums -to_string -type=toString; DO NOT EDIT

package main

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

var (
	_toStringNameToValue = map[string]toString{
		"elemA": elemA,
		"elemB": elemB,
		"elemC": elemC,
	}

	_toStringValueToName = map[toString]string{
		elemA: "elemA",
		elemB: "elemB",
		elemC: "elemC",
	}
)

func init() {
	var v toString
	if _, ok := interface{}(v).(fmt.Stringer); ok {
		_toStringNameToValue = map[string]toString{
			interface{}(elemA).(fmt.Stringer).String(): elemA,
			interface{}(elemB).(fmt.Stringer).String(): elemB,
			interface{}(elemC).(fmt.Stringer).String(): elemC,
		}
	}
}

func (r toString) toString() (string, error) {
	s, ok := _toStringValueToName[r]
	if !ok {
		return "", fmt.Errorf("invalid toString: %d", r)
	}
	return s, nil
}

func (r toString) ToString() (string, error) {
	return r.toString()
}

func (r toString) getString() (string, error) {
	if s, ok := interface{}(r).(fmt.Stringer); ok {
		return s.String(), nil
	}
	return r.toString()
}

func (r *toString) setValue(str string) error {
	v, ok := _toStringNameToValue[str]
	if !ok {
		return fmt.Errorf("invalid toString %q", str)
	}
	*r = v
	return nil
}

// MarshalJSON is generated so toString satisfies json.Marshaler.
func (r toString) MarshalJSON() ([]byte, error) {
	s, err := r.getString()
	if err != nil {
		return nil, err
	}
	return json.Marshal(s)
}

// UnmarshalJSON is generated so toString satisfies json.Unmarshaler.
func (r *toString) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("toString should be a string, got %s", data)
	}
	return r.setValue(s)
}

//Scan an input string into this structure for use with GORP
func (r *toString) Scan(i interface{}) error {
	switch t := i.(type) {
	case []byte:
		return r.setValue(string(t))
	case string:
		return r.setValue(t)
	default:
		return fmt.Errorf("Can't scan %T into type %T", i, r)
	}
	return nil
}

func (r toString) Value() (driver.Value, error) {
	return r.getString()
}
