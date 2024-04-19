package jsonx

import (
	"bytes"
	"encoding/json"
	"io"
)

// RawMessage is a raw encoded JSON value.
// It implements Marshaler and Unmarshaler and can
// be used to delay JSON decoding or precompute a JSON encoding.
type RawMessage = json.RawMessage

// Marshal adapts to json/encoding Marshal API.
//
// Marshal returns the JSON encoding of v, adapts to json/encoding Marshal API
// Refer to https://godoc.org/encoding/json#Marshal for more information.
func Marshal(v interface{}) (marshaledBytes []byte, err error) {
	return json.Marshal(v)
}

// MarshalIndent same as json.MarshalIndent.
func MarshalIndent(v interface{}, prefix, indent string) (marshaledBytes []byte, err error) {
	return json.MarshalIndent(v, prefix, indent)
}

// Unmarshal adapts to json/encoding Unmarshal API
//
// Unmarshal parses the JSON-encoded data and stores the result in the value pointed to by v.
// Refer to https://godoc.org/encoding/json#Unmarshal for more information.
func Unmarshal(data []byte, v interface{}) (err error) {
	return json.Unmarshal(data, v)
}

// UnmarshalUseNumber decodes the json data bytes to target interface using number option.
func UnmarshalUseNumber(data []byte, v interface{}) (err error) {
	decoder := NewDecoder(bytes.NewReader(data))
	decoder.UseNumber()
	err = decoder.Decode(v)
	return
}

// NewEncoder same as json.NewEncoder
func NewEncoder(writer io.Writer) *json.Encoder {
	return json.NewEncoder(writer)
}

// NewDecoder adapts to json/stream NewDecoder API.
//
// NewDecoder returns a new decoder that reads from r.
//
// Instead of a json/encoding Decoder, a Decoder is returned
// Refer to https://godoc.org/encoding/json#NewDecoder for more information.
func NewDecoder(reader io.Reader) *json.Decoder {
	return json.NewDecoder(reader)
}

// Valid reports whether data is a valid JSON encoding.
func Valid(data []byte) bool {
	return json.Valid(data)
}
