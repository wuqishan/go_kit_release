package conv

import (
	"encoding/json"
	"github.com/wuqishan/go_kit/base/jsonx"
	"github.com/wuqishan/go_kit/base/reflectionx"
	"reflect"
)

// SliceFloat ...
func SliceFloat(any interface{}) []float64 {
	return Floats(any)
}

// SliceFloat32 ...
func SliceFloat32(any interface{}) []float32 {
	return Float32s(any)
}

// SliceFloat64 ...
func SliceFloat64(any interface{}) []float64 {
	return Floats(any)
}

// Floats ...
func Floats(any interface{}) []float64 {
	return Float64s(any)
}

// Float32s ...
func Float32s(any interface{}) []float32 {
	if any == nil {
		return nil
	}
	var (
		array []float32 = nil
	)
	switch value := any.(type) {
	case string:
		if value == "" {
			return []float32{}
		}
		return []float32{Float32(value)}
	case []string:
		array = make([]float32, len(value))
		for k, v := range value {
			array[k] = Float32(v)
		}
	case []int:
		array = make([]float32, len(value))
		for k, v := range value {
			array[k] = Float32(v)
		}
	case []int8:
		array = make([]float32, len(value))
		for k, v := range value {
			array[k] = Float32(v)
		}
	case []int16:
		array = make([]float32, len(value))
		for k, v := range value {
			array[k] = Float32(v)
		}
	case []int32:
		array = make([]float32, len(value))
		for k, v := range value {
			array[k] = Float32(v)
		}
	case []int64:
		array = make([]float32, len(value))
		for k, v := range value {
			array[k] = Float32(v)
		}
	case []uint:
		for _, v := range value {
			array = append(array, Float32(v))
		}
	case []uint8:
		if json.Valid(value) {
			_ = jsonx.UnmarshalUseNumber(value, &array)
		} else {
			array = make([]float32, len(value))
			for k, v := range value {
				array[k] = Float32(v)
			}
		}
	case []uint16:
		array = make([]float32, len(value))
		for k, v := range value {
			array[k] = Float32(v)
		}
	case []uint32:
		array = make([]float32, len(value))
		for k, v := range value {
			array[k] = Float32(v)
		}
	case []uint64:
		array = make([]float32, len(value))
		for k, v := range value {
			array[k] = Float32(v)
		}
	case []bool:
		array = make([]float32, len(value))
		for k, v := range value {
			array[k] = Float32(v)
		}
	case []float32:
		array = value
	case []float64:
		array = make([]float32, len(value))
		for k, v := range value {
			array[k] = Float32(v)
		}
	case []interface{}:
		array = make([]float32, len(value))
		for k, v := range value {
			array[k] = Float32(v)
		}
	}
	if array != nil {
		return array
	}
	if v, ok := any.(iFloats); ok {
		return Float32s(v.Floats())
	}
	if v, ok := any.(iInterfaces); ok {
		return Float32s(v.Interfaces())
	}
	// JSON format string value converting.
	if checkJsonAndUnmarshalUseNumber(any, &array) {
		return array
	}
	// Not a common type, it then uses reflection for conversion.
	originValueAndKind := reflectionx.OriginValueAndKind(any)
	switch originValueAndKind.OriginKind {
	case reflect.Slice, reflect.Array:
		var (
			length = originValueAndKind.OriginValue.Len()
			slice  = make([]float32, length)
		)
		for i := 0; i < length; i++ {
			slice[i] = Float32(originValueAndKind.OriginValue.Index(i).Interface())
		}
		return slice

	default:
		if originValueAndKind.OriginValue.IsZero() {
			return []float32{}
		}
		return []float32{Float32(any)}
	}
}

// Float64s ...
func Float64s(any interface{}) []float64 {
	if any == nil {
		return nil
	}
	var (
		array []float64 = nil
	)
	switch value := any.(type) {
	case string:
		if value == "" {
			return []float64{}
		}
		return []float64{Float64(value)}
	case []string:
		array = make([]float64, len(value))
		for k, v := range value {
			array[k] = Float64(v)
		}
	case []int:
		array = make([]float64, len(value))
		for k, v := range value {
			array[k] = Float64(v)
		}
	case []int8:
		array = make([]float64, len(value))
		for k, v := range value {
			array[k] = Float64(v)
		}
	case []int16:
		array = make([]float64, len(value))
		for k, v := range value {
			array[k] = Float64(v)
		}
	case []int32:
		array = make([]float64, len(value))
		for k, v := range value {
			array[k] = Float64(v)
		}
	case []int64:
		array = make([]float64, len(value))
		for k, v := range value {
			array[k] = Float64(v)
		}
	case []uint:
		for _, v := range value {
			array = append(array, Float64(v))
		}
	case []uint8:
		if json.Valid(value) {
			_ = jsonx.UnmarshalUseNumber(value, &array)
		} else {
			array = make([]float64, len(value))
			for k, v := range value {
				array[k] = Float64(v)
			}
		}
	case []uint16:
		array = make([]float64, len(value))
		for k, v := range value {
			array[k] = Float64(v)
		}
	case []uint32:
		array = make([]float64, len(value))
		for k, v := range value {
			array[k] = Float64(v)
		}
	case []uint64:
		array = make([]float64, len(value))
		for k, v := range value {
			array[k] = Float64(v)
		}
	case []bool:
		array = make([]float64, len(value))
		for k, v := range value {
			array[k] = Float64(v)
		}
	case []float32:
		array = make([]float64, len(value))
		for k, v := range value {
			array[k] = Float64(v)
		}
	case []float64:
		array = value
	case []interface{}:
		array = make([]float64, len(value))
		for k, v := range value {
			array[k] = Float64(v)
		}
	}
	if array != nil {
		return array
	}
	if v, ok := any.(iFloats); ok {
		return v.Floats()
	}
	if v, ok := any.(iInterfaces); ok {
		return Floats(v.Interfaces())
	}
	// JSON format string value converting.
	if checkJsonAndUnmarshalUseNumber(any, &array) {
		return array
	}
	// Not a common type, it then uses reflection for conversion.
	originValueAndKind := reflectionx.OriginValueAndKind(any)
	switch originValueAndKind.OriginKind {
	case reflect.Slice, reflect.Array:
		var (
			length = originValueAndKind.OriginValue.Len()
			slice  = make([]float64, length)
		)
		for i := 0; i < length; i++ {
			slice[i] = Float64(originValueAndKind.OriginValue.Index(i).Interface())
		}
		return slice

	default:
		if originValueAndKind.OriginValue.IsZero() {
			return []float64{}
		}
		return []float64{Float64(any)}
	}
}