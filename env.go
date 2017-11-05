// Package env contains a parser for reading and parsing environment variables to go structs
package env

import (
	"errors"
	"fmt"
	"os"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const (
	// DefaultSeparator is used as the split character by default.
	// If your env variables contain this value, you'll need to override separator to something else
	DefaultSeparator = ","

	constAliasTypeByte = "byte"
	constAliasTypeRune = "rune"
)

// slices of supported types
var (
	sliceInt      = reflect.TypeOf([]int(nil))
	sliceInt8     = reflect.TypeOf([]int8(nil))
	sliceInt16    = reflect.TypeOf([]int16(nil))
	sliceInt32    = reflect.TypeOf([]int32(nil))
	sliceInt64    = reflect.TypeOf([]int64(nil))
	sliceUint     = reflect.TypeOf([]uint(nil))
	sliceUint8    = reflect.TypeOf([]uint8(nil))
	sliceUint16   = reflect.TypeOf([]uint16(nil))
	sliceUint32   = reflect.TypeOf([]uint32(nil))
	sliceUint64   = reflect.TypeOf([]uint64(nil))
	sliceString   = reflect.TypeOf([]string(nil))
	sliceBool     = reflect.TypeOf([]bool(nil))
	sliceFloat32  = reflect.TypeOf([]float32(nil))
	sliceFloat64  = reflect.TypeOf([]float64(nil))
	sliceDuration = reflect.TypeOf([]time.Duration(nil))
	sliceTime     = reflect.TypeOf([]time.Time(nil))
	sliceRegexp   = reflect.TypeOf([]*regexp.Regexp(nil))

	// internal aliases, so not strictly necessary. Useful for documentational purposes.
	sliceByte = reflect.TypeOf([]byte(nil))
	sliceRune = reflect.TypeOf([]rune(nil))
)

// Parse parses the environment values to the specified struct based on the struct tags
//
// Possible tag options are:
// 	required       - the field must have a non-zero value
// 	default=Y      - the default value to use if variable is unset in environment
// 	separator=X    - separator for multivalue environment values
// 	type=byte|rune - type of value for values which reflect cannot distinguish between itself
//
// Example usage:
//  type Config struct {
//  	Host 	 string   `env:"HOST,required,default=localhost"`
//  	Secret 	 []byte   `env:"SECRET,required,type=byte"`
//  	Versions []string `env:"VALUES,default=v1"`
//  	Names 	 []string `env:"VALUES,default=n1:n2:n3,separator=:"`
//  }
//
//  config := &Config{}
//  env.Parse(&config)
//
// See env_test.go for complete examples.
func Parse(v interface{}) error {
	ptr := reflect.ValueOf(v)
	if ptr.Kind() != reflect.Ptr {
		return errors.New("Expected a pointer value")
	}

	elem := ptr.Elem()
	if elem.Kind() != reflect.Struct {
		return errors.New("Expected a struct pointer")
	}

	if err := parseEnv(elem); err != nil {
		return err
	}

	return nil
}

func parseEnv(s reflect.Value) error {
	sType := s.Type()
	fieldCount := sType.NumField()
	for i := 0; i < fieldCount; i++ {
		field := s.Field(i)
		fieldType := sType.Field(i)
		tagValue := fieldType.Tag.Get("env")

		if tagValue == "" {
			// if the field is a nested struct, parse it and continue to next field
			if field.Kind() == reflect.Struct {
				if err := parseEnv(field); err != nil {
					return err
				}
			}
			continue
		}

		tags := strings.Split(tagValue, ",")
		if len(tags) < 1 || len(tags[0]) == 0 {
			return errors.New("env variable name cannot be empty")
		}

		envVariableName := tags[0]
		value := os.Getenv(envVariableName)

		// parse environment based on tags
		separator := DefaultSeparator
		aliasType := ""
		for _, tagValue := range tags[1:] {
			if tagValue == "required" {
				if value == "" {
					return asParseError(envVariableName, "value is required but was empty")
				}
			} else if value == "" && strings.HasPrefix(tagValue, "default") {
				value = namedOptionValue(tagValue)
			} else if strings.HasPrefix(tagValue, "type") { // type allows override for go native aliases (byte,rune)
				aliasType = namedOptionValue(tagValue)

				if aliasType != constAliasTypeRune && aliasType != constAliasTypeByte {
					return asParseError(envVariableName, fmt.Sprintf("invalid type \"%s\", valid options are: \"%s\", \"%s\"", tagValue, constAliasTypeByte, constAliasTypeRune))
				}
			} else if strings.HasPrefix(tagValue, "separator") {
				if tmp := namedOptionValue(tagValue); tmp != "" {
					separator = tmp
				}
			} else {
				return asParseError(envVariableName, fmt.Sprintf("unknown option %s", tagValue))
			}
		}

		// parse value to correct type and set it to field
		switch field.Kind() {
		default:
			if err := parseSingle(fieldType, field, envVariableName, value, aliasType); err != nil {
				return err
			}

		case reflect.Slice:
			if err := parseSlice(fieldType, field, envVariableName, value, separator, aliasType); err != nil {
				return err
			}
		}
	}

	return nil
}

func parseSlice(fieldType reflect.StructField, field reflect.Value, envVariableName, value, separator, aliasType string) error {
	data := strings.Split(value, separator)

	switch field.Type() {
	case sliceUint:
		parsed := make([]uint, len(data))
		for idx, d := range data {
			p, err := strconv.ParseUint(d, 10, 32)
			if err != nil {
				return asParseError(envVariableName, err.Error())
			}
			parsed[idx] = uint(p)
		}
		field.Set(reflect.ValueOf(parsed))
	case sliceUint8, sliceByte:
		if aliasType == constAliasTypeByte {
			if len(data) > 1 {
				return asParseError(envVariableName, "byte slice cannot have multiple values")
			}
			field.Set(reflect.ValueOf([]byte(data[0])))
			return nil
		}
		parsed := make([]uint8, len(data))
		for idx, d := range data {
			p, err := strconv.ParseUint(d, 10, 8)
			if err != nil {
				return asParseError(envVariableName, err.Error())
			}
			parsed[idx] = uint8(p)
		}
		field.Set(reflect.ValueOf(parsed))
	case sliceUint16:
		parsed := make([]uint16, len(data))
		for idx, d := range data {
			p, err := strconv.ParseUint(d, 10, 16)
			if err != nil {
				return asParseError(envVariableName, err.Error())
			}
			parsed[idx] = uint16(p)
		}
		field.Set(reflect.ValueOf(parsed))
	case sliceUint32:
		parsed := make([]uint32, len(data))
		for idx, d := range data {
			p, err := strconv.ParseUint(d, 10, 32)
			if err != nil {
				return asParseError(envVariableName, err.Error())
			}
			parsed[idx] = uint32(p)
		}
		field.Set(reflect.ValueOf(parsed))
	case sliceUint64:
		parsed := make([]uint64, len(data))
		for idx, d := range data {
			p, err := strconv.ParseUint(d, 10, 64)
			if err != nil {
				return asParseError(envVariableName, err.Error())
			}
			parsed[idx] = uint64(p)
		}
		field.Set(reflect.ValueOf(parsed))
	case sliceInt:
		parsed := make([]int, len(data))
		for idx, d := range data {
			p, err := strconv.ParseInt(d, 10, 32)
			if err != nil {
				return asParseError(envVariableName, err.Error())
			}
			parsed[idx] = int(p)
		}
		field.Set(reflect.ValueOf(parsed))
	case sliceInt8:
		parsed := make([]int8, len(data))
		for idx, d := range data {
			p, err := strconv.ParseInt(d, 10, 8)
			if err != nil {
				return asParseError(envVariableName, err.Error())
			}
			parsed[idx] = int8(p)
		}
		field.Set(reflect.ValueOf(parsed))
	case sliceInt16:
		parsed := make([]int16, len(data))
		for idx, d := range data {
			p, err := strconv.ParseInt(d, 10, 16)
			if err != nil {
				return asParseError(envVariableName, err.Error())
			}
			parsed[idx] = int16(p)
		}
		field.Set(reflect.ValueOf(parsed))
	case sliceInt32, sliceRune:
		if aliasType == constAliasTypeRune {
			if len(data) > 1 {
				return asParseError(envVariableName, "rune slice cannot have multiple values")
			}
			field.Set(reflect.ValueOf([]rune(data[0])))
			return nil
		}
		parsed := make([]int32, len(data))
		for idx, d := range data {
			p, err := strconv.ParseInt(d, 10, 32)
			if err != nil {
				return asParseError(envVariableName, err.Error())
			}
			parsed[idx] = int32(p)
		}
		field.Set(reflect.ValueOf(parsed))
	case sliceDuration:
		parsed := make([]time.Duration, len(data))
		for idx, d := range data {
			v, err := time.ParseDuration(d)
			if err != nil {
				return asParseError(envVariableName, err.Error())
			}
			parsed[idx] = v
		}
		field.Set(reflect.ValueOf(parsed))
		return nil
	case sliceTime:
		parsed := make([]time.Time, len(data))
		for idx, d := range data {
			v, err := time.Parse(time.RFC3339, d)
			if err != nil {
				return asParseError(envVariableName, err.Error())
			}
			parsed[idx] = v
		}
		field.Set(reflect.ValueOf(parsed))
		return nil
	case sliceRegexp:
		parsed := make([]*regexp.Regexp, len(data))
		for idx, d := range data {
			v, err := regexp.Compile(d)
			if err != nil {
				return asParseError(envVariableName, err.Error())
			}
			parsed[idx] = v
		}
		field.Set(reflect.ValueOf(parsed))
		return nil
	case sliceInt64:
		parsed := make([]int64, len(data))
		for idx, d := range data {
			p, err := strconv.ParseInt(d, 10, 64)
			if err != nil {
				return asParseError(envVariableName, err.Error())
			}
			parsed[idx] = int64(p)
		}
		field.Set(reflect.ValueOf(parsed))
	case sliceString:
		field.Set(reflect.ValueOf(data))
	case sliceBool:
		parsed := make([]bool, len(data))
		for idx, d := range data {
			p, err := strconv.ParseBool(d)
			if err != nil {
				return asParseError(envVariableName, err.Error())
			}
			parsed[idx] = bool(p)
		}
		field.Set(reflect.ValueOf(parsed))
	case sliceFloat32:
		parsed := make([]float32, len(data))
		for idx, d := range data {
			p, err := strconv.ParseFloat(d, 32)
			if err != nil {
				return asParseError(envVariableName, err.Error())
			}
			parsed[idx] = float32(p)
		}
		field.Set(reflect.ValueOf(parsed))
	case sliceFloat64:
		parsed := make([]float64, len(data))
		for idx, d := range data {
			p, err := strconv.ParseFloat(d, 64)
			if err != nil {
				return asParseError(envVariableName, err.Error())
			}
			parsed[idx] = float64(p)
		}
		field.Set(reflect.ValueOf(parsed))

	default:
		return errors.New("Unrecognized slice type")
	}
	return nil
}

func parseSingle(fieldType reflect.StructField, field reflect.Value, envVariableName, value, aliasType string) error {
	if fieldType.Type.String() == "time.Duration" {
		v, err := time.ParseDuration(value)
		if err != nil {
			return asParseError(envVariableName, err.Error())
		}
		field.Set(reflect.ValueOf(v))
		return nil
	}

	if fieldType.Type.String() == "time.Time" {
		v, err := time.Parse(time.RFC3339, value)
		if err != nil {
			return asParseError(envVariableName, err.Error())
		}
		field.Set(reflect.ValueOf(v))
		return nil
	}

	if fieldType.Type.String() == "*regexp.Regexp" {
		v, err := regexp.Compile(value)
		if err != nil {
			return asParseError(envVariableName, err.Error())
		}
		field.Set(reflect.ValueOf(v))
		return nil
	}

	// handle byte
	if field.Kind() == reflect.Uint8 && aliasType == constAliasTypeByte {
		if len(value) != 1 {
			return asParseError(envVariableName, "byte must be a single character value")
		}
		field.Set(reflect.ValueOf(byte(value[0])))
		return nil
	}

	// handle rune
	if field.Kind() == reflect.Int32 && aliasType == constAliasTypeRune {
		if len(value) != 1 {
			return asParseError(envVariableName, "rune must be a single character value")
		}
		field.Set(reflect.ValueOf(rune(value[0])))
		return nil
	}

	switch field.Kind() {
	case reflect.String:
		field.SetString(value)

	case reflect.Bool:
		bvalue, err := strconv.ParseBool(value)
		if err != nil {
			return asParseError(envVariableName, err.Error())
		}
		field.SetBool(bvalue)

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		var bitSize int
		switch field.Kind() {
		case reflect.Uint8:
			bitSize = 8
		case reflect.Uint16:
			bitSize = 16
		case reflect.Uint, reflect.Uint32:
			bitSize = 32
		case reflect.Uint64:
			bitSize = 64
		}
		uintValue, err := strconv.ParseUint(value, 10, bitSize)
		if err != nil {
			return asParseError(envVariableName, err.Error())
		}
		field.SetUint(uintValue)

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		var bitSize int
		switch field.Kind() {
		case reflect.Int8:
			bitSize = 8
		case reflect.Int16:
			bitSize = 16
		case reflect.Int, reflect.Int32:
			bitSize = 32
		case reflect.Int64:
			bitSize = 64
		}
		intValue, err := strconv.ParseInt(value, 10, bitSize)
		if err != nil {
			return asParseError(envVariableName, err.Error())
		}
		field.SetInt(intValue)

	case reflect.Float32, reflect.Float64:
		var bitSize int
		switch field.Kind() {
		case reflect.Float32:
			bitSize = 32
		case reflect.Float64:
			bitSize = 64
		}
		v, err := strconv.ParseFloat(value, bitSize)
		if err != nil {
			return asParseError(envVariableName, err.Error())
		}
		field.SetFloat(v)

	default:
		return errors.New("Unrecognized type")
	}
	return nil
}

func asParseError(envVariableName, err string) error {
	return fmt.Errorf("%s: %s", envVariableName, err)
}

func namedOptionValue(val string) string {
	split := strings.SplitN(val, "=", 2)
	if len(split) != 2 {
		return ""
	}
	return split[1]
}
