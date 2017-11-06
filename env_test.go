package env_test

import (
	"errors"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
	"testing"
	"time"

	env "github.com/stenhagglund/go-env"
	"github.com/stretchr/testify/require"
)

type testEnvStruct struct {
	IntValue           int              `env:"GO_ENV_TEST_INT_VALUE"`
	Int8Value          int8             `env:"GO_ENV_TEST_INT8_VALUE"`
	Int16Value         int16            `env:"GO_ENV_TEST_INT16_VALUE"`
	Int32Value         int32            `env:"GO_ENV_TEST_INT32_VALUE"`
	Int64Value         int64            `env:"GO_ENV_TEST_INT64_VALUE"`
	UintValue          uint             `env:"GO_ENV_TEST_UINT_VALUE"`
	Uint8Value         uint8            `env:"GO_ENV_TEST_UINT8_VALUE"`
	Uint16Value        uint16           `env:"GO_ENV_TEST_UINT16_VALUE"`
	Uint32Value        uint32           `env:"GO_ENV_TEST_UINT32_VALUE"`
	Uint64Value        uint64           `env:"GO_ENV_TEST_UINT64_VALUE"`
	Float32Value       float32          `env:"GO_ENV_TEST_FLOAT32_VALUE"`
	Float64Value       float64          `env:"GO_ENV_TEST_FLOAT64_VALUE"`
	IntSliceValue      []int            `env:"GO_ENV_TEST_INT_SLICE_VALUE"`
	Int8SliceValue     []int8           `env:"GO_ENV_TEST_INT8_SLICE_VALUE"`
	Int16SliceValue    []int16          `env:"GO_ENV_TEST_INT16_SLICE_VALUE"`
	Int32SliceValue    []int32          `env:"GO_ENV_TEST_INT32_SLICE_VALUE"`
	Int64SliceValue    []int64          `env:"GO_ENV_TEST_INT64_SLICE_VALUE"`
	UintSliceValue     []uint           `env:"GO_ENV_TEST_UINT_SLICE_VALUE"`
	Uint8SliceValue    []uint8          `env:"GO_ENV_TEST_UINT8_SLICE_VALUE"`
	Uint16SliceValue   []uint16         `env:"GO_ENV_TEST_UINT16_SLICE_VALUE"`
	Uint32SliceValue   []uint32         `env:"GO_ENV_TEST_UINT32_SLICE_VALUE"`
	Uint64SliceValue   []uint64         `env:"GO_ENV_TEST_UINT64_SLICE_VALUE"`
	Float32SliceValue  []float32        `env:"GO_ENV_TEST_FLOAT32_SLICE_VALUE"`
	Float64SliceValue  []float64        `env:"GO_ENV_TEST_FLOAT64_SLICE_VALUE"`
	BoolValue          bool             `env:"GO_ENV_TEST_BOOL_VALUE"`
	BoolSliceValue     []bool           `env:"GO_ENV_TEST_BOOL_SLICE_VALUE"`
	StringValue        string           `env:"GO_ENV_TEST_STRING_VALUE"`
	StringSliceValue   []string         `env:"GO_ENV_TEST_STRING_SLICE_VALUE"`
	ByteValue          byte             `env:"GO_ENV_TEST_BYTE_VALUE,type=byte"`
	ByteSliceValue     []byte           `env:"GO_ENV_TEST_BYTE_SLICE_VALUE,type=byte"`
	RuneValue          rune             `env:"GO_ENV_TEST_RUNE_VALUE,type=rune"`
	RuneSliceValue     []rune           `env:"GO_ENV_TEST_RUNE_SLICE_VALUE,type=rune"`
	DurationValue      time.Duration    `env:"GO_ENV_TEST_TIME_DURATION_VALUE"`
	DurationSliceValue []time.Duration  `env:"GO_ENV_TEST_TIME_DURATION_SLICE_VALUE"`
	TimeValue          time.Time        `env:"GO_ENV_TEST_TIME_VALUE"`
	TimeSliceValue     []time.Time      `env:"GO_ENV_TEST_TIME_SLICE_VALUE"`
	RegexpValue        *regexp.Regexp   `env:"GO_ENV_TEST_REGEXP_POINTER_VALUE"`
	RegexpSliceValue   []*regexp.Regexp `env:"GO_ENV_TEST_REGEXP_POINTER_SLICE_VALUE"`
}

type nestedTestEnvStruct struct {
	testEnvStruct
	Nested testEnvStruct
}

type testEnvStructRequiredValues struct {
	IntValue           int              `env:"GO_ENV_TEST_INT_VALUE,required"`
	Int8Value          int8             `env:"GO_ENV_TEST_INT8_VALUE,required"`
	Int16Value         int16            `env:"GO_ENV_TEST_INT16_VALUE,required"`
	Int32Value         int32            `env:"GO_ENV_TEST_INT32_VALUE,required"`
	Int64Value         int64            `env:"GO_ENV_TEST_INT64_VALUE,required"`
	UintValue          uint             `env:"GO_ENV_TEST_UINT_VALUE,required"`
	Uint8Value         uint8            `env:"GO_ENV_TEST_UINT8_VALUE,required"`
	Uint16Value        uint16           `env:"GO_ENV_TEST_UINT16_VALUE,required"`
	Uint32Value        uint32           `env:"GO_ENV_TEST_UINT32_VALUE,required"`
	Uint64Value        uint64           `env:"GO_ENV_TEST_UINT64_VALUE,required"`
	Float32Value       float32          `env:"GO_ENV_TEST_FLOAT32_VALUE,required"`
	Float64Value       float64          `env:"GO_ENV_TEST_FLOAT64_VALUE,required"`
	IntSliceValue      []int            `env:"GO_ENV_TEST_INT_SLICE_VALUE,required"`
	Int8SliceValue     []int8           `env:"GO_ENV_TEST_INT8_SLICE_VALUE,required"`
	Int16SliceValue    []int16          `env:"GO_ENV_TEST_INT16_SLICE_VALUE,required"`
	Int32SliceValue    []int32          `env:"GO_ENV_TEST_INT32_SLICE_VALUE,required"`
	Int64SliceValue    []int64          `env:"GO_ENV_TEST_INT64_SLICE_VALUE,required"`
	UintSliceValue     []uint           `env:"GO_ENV_TEST_UINT_SLICE_VALUE,required"`
	Uint8SliceValue    []uint8          `env:"GO_ENV_TEST_UINT8_SLICE_VALUE,required"`
	Uint16SliceValue   []uint16         `env:"GO_ENV_TEST_UINT16_SLICE_VALUE,required"`
	Uint32SliceValue   []uint32         `env:"GO_ENV_TEST_UINT32_SLICE_VALUE,required"`
	Uint64SliceValue   []uint64         `env:"GO_ENV_TEST_UINT64_SLICE_VALUE,required"`
	Float32SliceValue  []float32        `env:"GO_ENV_TEST_FLOAT32_SLICE_VALUE,required"`
	Float64SliceValue  []float64        `env:"GO_ENV_TEST_FLOAT64_SLICE_VALUE,required"`
	BoolValue          bool             `env:"GO_ENV_TEST_BOOL_VALUE,required"`
	BoolSliceValue     []bool           `env:"GO_ENV_TEST_BOOL_SLICE_VALUE,required"`
	StringValue        string           `env:"GO_ENV_TEST_STRING_VALUE,required"`
	StringSliceValue   []string         `env:"GO_ENV_TEST_STRING_SLICE_VALUE,required"`
	ByteValue          byte             `env:"GO_ENV_TEST_BYTE_VALUE,type=byte,required"`
	ByteSliceValue     []byte           `env:"GO_ENV_TEST_BYTE_SLICE_VALUE,type=byte,required"`
	RuneValue          rune             `env:"GO_ENV_TEST_RUNE_VALUE,type=rune,required"`
	RuneSliceValue     []rune           `env:"GO_ENV_TEST_RUNE_SLICE_VALUE,type=rune,required"`
	DurationValue      time.Duration    `env:"GO_ENV_TEST_TIME_DURATION_VALUE,required"`
	DurationSliceValue []time.Duration  `env:"GO_ENV_TEST_TIME_DURATION_SLICE_VALUE,required"`
	TimeValue          time.Time        `env:"GO_ENV_TEST_TIME_VALUE,required"`
	TimeSliceValue     []time.Time      `env:"GO_ENV_TEST_TIME_SLICE_VALUE,required"`
	RegexpValue        *regexp.Regexp   `env:"GO_ENV_TEST_REGEXP_POINTER_VALUE,required"`
	RegexpSliceValue   []*regexp.Regexp `env:"GO_ENV_TEST_REGEXP_POINTER_SLICE_VALUE,required"`
}

// assert both struct are compatible to ensure exactly same keys
var _ = testEnvStruct(testEnvStructRequiredValues{})

type testEnvStructDefaultValues struct {
	IntValue           int              `env:"GO_ENV_TEST_INT_VALUE,default=12"`
	Int8Value          int8             `env:"GO_ENV_TEST_INT8_VALUE,default=22"`
	Int16Value         int16            `env:"GO_ENV_TEST_INT16_VALUE,default=32"`
	Int32Value         int32            `env:"GO_ENV_TEST_INT32_VALUE,default=42"`
	Int64Value         int64            `env:"GO_ENV_TEST_INT64_VALUE,default=52"`
	UintValue          uint             `env:"GO_ENV_TEST_UINT_VALUE,default=16"`
	Uint8Value         uint8            `env:"GO_ENV_TEST_UINT8_VALUE,default=26"`
	Uint16Value        uint16           `env:"GO_ENV_TEST_UINT16_VALUE,default=36"`
	Uint32Value        uint32           `env:"GO_ENV_TEST_UINT32_VALUE,default=46"`
	Uint64Value        uint64           `env:"GO_ENV_TEST_UINT64_VALUE,default=56"`
	Float32Value       float32          `env:"GO_ENV_TEST_FLOAT32_VALUE,default=12.0"`
	Float64Value       float64          `env:"GO_ENV_TEST_FLOAT64_VALUE,default=22.0"`
	IntSliceValue      []int            `env:"GO_ENV_TEST_INT_SLICE_VALUE,default=112"`
	Int8SliceValue     []int8           `env:"GO_ENV_TEST_INT8_SLICE_VALUE,default=122"`
	Int16SliceValue    []int16          `env:"GO_ENV_TEST_INT16_SLICE_VALUE,default=132"`
	Int32SliceValue    []int32          `env:"GO_ENV_TEST_INT32_SLICE_VALUE,default=142"`
	Int64SliceValue    []int64          `env:"GO_ENV_TEST_INT64_SLICE_VALUE,default=152"`
	UintSliceValue     []uint           `env:"GO_ENV_TEST_UINT_SLICE_VALUE,default=117"`
	Uint8SliceValue    []uint8          `env:"GO_ENV_TEST_UINT8_SLICE_VALUE,default=127"`
	Uint16SliceValue   []uint16         `env:"GO_ENV_TEST_UINT16_SLICE_VALUE,default=137"`
	Uint32SliceValue   []uint32         `env:"GO_ENV_TEST_UINT32_SLICE_VALUE,default=147"`
	Uint64SliceValue   []uint64         `env:"GO_ENV_TEST_UINT64_SLICE_VALUE,default=157"`
	Float32SliceValue  []float32        `env:"GO_ENV_TEST_FLOAT32_SLICE_VALUE,default=112.0"`
	Float64SliceValue  []float64        `env:"GO_ENV_TEST_FLOAT64_SLICE_VALUE,default=122.0"`
	BoolValue          bool             `env:"GO_ENV_TEST_BOOL_VALUE,default=1"`
	BoolSliceValue     []bool           `env:"GO_ENV_TEST_BOOL_SLICE_VALUE,default=1"`
	StringValue        string           `env:"GO_ENV_TEST_STRING_VALUE,default=exp/def/string"`
	StringSliceValue   []string         `env:"GO_ENV_TEST_STRING_SLICE_VALUE,default=exp/def/stringslice"`
	ByteValue          byte             `env:"GO_ENV_TEST_BYTE_VALUE,type=byte,default=b"`
	ByteSliceValue     []byte           `env:"GO_ENV_TEST_BYTE_SLICE_VALUE,type=byte,default=exp/def/byteslice"`
	RuneValue          rune             `env:"GO_ENV_TEST_RUNE_VALUE,type=rune,default=r"`
	RuneSliceValue     []rune           `env:"GO_ENV_TEST_RUNE_SLICE_VALUE,type=rune,default=exp/def/runeslice"`
	DurationValue      time.Duration    `env:"GO_ENV_TEST_TIME_DURATION_VALUE,default=10s"`
	DurationSliceValue []time.Duration  `env:"GO_ENV_TEST_TIME_DURATION_SLICE_VALUE,default=15s"`
	TimeValue          time.Time        `env:"GO_ENV_TEST_TIME_VALUE,default=2006-01-02T15:04:05Z"`
	TimeSliceValue     []time.Time      `env:"GO_ENV_TEST_TIME_SLICE_VALUE,default=2006-02-02T15:04:05Z"`
	RegexpValue        *regexp.Regexp   `env:"GO_ENV_TEST_REGEXP_POINTER_VALUE,default=def"`        // only simple regexes may be in defaults. TODO: research tags
	RegexpSliceValue   []*regexp.Regexp `env:"GO_ENV_TEST_REGEXP_POINTER_SLICE_VALUE,default=def2"` // only simple regexes may be in defaults. TODO: research tags
}

// assert both struct are compatible to ensure exactly same keys
var _ = testEnvStruct(testEnvStructDefaultValues{})

type testEnvStructSeparator struct {
	IntValue           int              `env:"GO_ENV_TEST_INT_VALUE,separator=;"`
	Int8Value          int8             `env:"GO_ENV_TEST_INT8_VALUE,separator=;"`
	Int16Value         int16            `env:"GO_ENV_TEST_INT16_VALUE,separator=;"`
	Int32Value         int32            `env:"GO_ENV_TEST_INT32_VALUE,separator=;"`
	Int64Value         int64            `env:"GO_ENV_TEST_INT64_VALUE,separator=;"`
	UintValue          uint             `env:"GO_ENV_TEST_UINT_VALUE,separator=;"`
	Uint8Value         uint8            `env:"GO_ENV_TEST_UINT8_VALUE,separator=;"`
	Uint16Value        uint16           `env:"GO_ENV_TEST_UINT16_VALUE,separator=;"`
	Uint32Value        uint32           `env:"GO_ENV_TEST_UINT32_VALUE,separator=;"`
	Uint64Value        uint64           `env:"GO_ENV_TEST_UINT64_VALUE,separator=;"`
	Float32Value       float32          `env:"GO_ENV_TEST_FLOAT32_VALUE,separator=;"`
	Float64Value       float64          `env:"GO_ENV_TEST_FLOAT64_VALUE,separator=;"`
	IntSliceValue      []int            `env:"GO_ENV_TEST_INT_SLICE_VALUE,separator=;"`
	Int8SliceValue     []int8           `env:"GO_ENV_TEST_INT8_SLICE_VALUE,separator=;"`
	Int16SliceValue    []int16          `env:"GO_ENV_TEST_INT16_SLICE_VALUE,separator=;"`
	Int32SliceValue    []int32          `env:"GO_ENV_TEST_INT32_SLICE_VALUE,separator=;"`
	Int64SliceValue    []int64          `env:"GO_ENV_TEST_INT64_SLICE_VALUE,separator=;"`
	UintSliceValue     []uint           `env:"GO_ENV_TEST_UINT_SLICE_VALUE,separator=;"`
	Uint8SliceValue    []uint8          `env:"GO_ENV_TEST_UINT8_SLICE_VALUE,separator=;"`
	Uint16SliceValue   []uint16         `env:"GO_ENV_TEST_UINT16_SLICE_VALUE,separator=;"`
	Uint32SliceValue   []uint32         `env:"GO_ENV_TEST_UINT32_SLICE_VALUE,separator=;"`
	Uint64SliceValue   []uint64         `env:"GO_ENV_TEST_UINT64_SLICE_VALUE,separator=;"`
	Float32SliceValue  []float32        `env:"GO_ENV_TEST_FLOAT32_SLICE_VALUE,separator=;"`
	Float64SliceValue  []float64        `env:"GO_ENV_TEST_FLOAT64_SLICE_VALUE,separator=;"`
	BoolValue          bool             `env:"GO_ENV_TEST_BOOL_VALUE,separator=;"`
	BoolSliceValue     []bool           `env:"GO_ENV_TEST_BOOL_SLICE_VALUE,separator=;"`
	StringValue        string           `env:"GO_ENV_TEST_STRING_VALUE,separator=;"`
	StringSliceValue   []string         `env:"GO_ENV_TEST_STRING_SLICE_VALUE,separator=;"`
	ByteValue          byte             `env:"GO_ENV_TEST_BYTE_VALUE,type=byte,separator=;"`
	ByteSliceValue     []byte           `env:"GO_ENV_TEST_BYTE_SLICE_VALUE,type=byte,separator=;"`
	RuneValue          rune             `env:"GO_ENV_TEST_RUNE_VALUE,type=rune,separator=;"`
	RuneSliceValue     []rune           `env:"GO_ENV_TEST_RUNE_SLICE_VALUE,type=rune,separator=;"`
	DurationValue      time.Duration    `env:"GO_ENV_TEST_TIME_DURATION_VALUE,separator=;"`
	DurationSliceValue []time.Duration  `env:"GO_ENV_TEST_TIME_DURATION_SLICE_VALUE,separator=;"`
	TimeValue          time.Time        `env:"GO_ENV_TEST_TIME_VALUE,separator=;"`
	TimeSliceValue     []time.Time      `env:"GO_ENV_TEST_TIME_SLICE_VALUE,separator=;"`
	RegexpValue        *regexp.Regexp   `env:"GO_ENV_TEST_REGEXP_POINTER_VALUE,separator=;"`
	RegexpSliceValue   []*regexp.Regexp `env:"GO_ENV_TEST_REGEXP_POINTER_SLICE_VALUE,separator=;"`
}

// assert both struct are compatible to ensure exactly same keys
var _ = testEnvStruct(testEnvStructSeparator{})

type nestedTestEnvStructSeparator struct {
	testEnvStructSeparator
	Nested testEnvStructSeparator
}

type nestedTestEnvStructDefaultValues struct {
	testEnvStructDefaultValues
	Nested testEnvStructDefaultValues
}

var (
	// HACK(SH): time equality comparison is broken in testify, see: https://github.com/stretchr/testify/issues/502
	// so ensure these would match the unmarshalled ones by marshaling + unmarshaling them
	testTime, _       = time.ParseInLocation(time.RFC3339, time.Now().Round(time.Second).Format(time.RFC3339), time.Local)
	testTimeSlice1, _ = time.ParseInLocation(time.RFC3339, time.Now().Add(time.Hour).Round(time.Second).Format(time.RFC3339), time.Local)
	testTimeSlice2, _ = time.ParseInLocation(time.RFC3339, time.Now().Add(-time.Hour).Round(time.Second).Format(time.RFC3339), time.Local)

	testRegexp       = regexp.MustCompile("\\Adef\\z")
	testRegexpSlice1 = regexp.MustCompile(".*")
	testRegexpSlice2 = regexp.MustCompile("\\A[A-z]*\\z")
)

func TestParseInvalidType(t *testing.T) {
	assert := require.New(t)

	withResetEnv(func() {
		os.Setenv("GO_ENV_TEST_INVALID_TYPE", "abc")
		for _, tc := range []interface{}{
			struct {
				Invalid uintptr `env:"GO_ENV_TEST_INVALID_TYPE"`
			}{},
			struct {
				Invalid []uintptr `env:"GO_ENV_TEST_INVALID_TYPE"`
			}{},
		} {
			assert.Error(env.Parse(&tc))
		}
	})
}

func TestParseInvalidArgument(t *testing.T) {
	assert := require.New(t)

	withResetEnv(func() {
		testInvalidParseTarget := []string{"abc"}
		assert.Equal(errors.New("Expected a struct pointer"), env.Parse(&testInvalidParseTarget))
		assert.Equal(errors.New("Expected a pointer value"), env.Parse(struct{}{}))
	})
}

func TestParseStructWithNonEnvValues(t *testing.T) {
	assert := require.New(t)

	withResetEnv(func() {
		type textNonEnvStruct struct {
			Skipped     int
			StringValue string `env:"GO_ENV_TEST_STRING_VALUE,required"`
		}

		os.Setenv("GO_ENV_TEST_STRING_VALUE", "abc")

		testStruct := textNonEnvStruct{}
		testStructExp := textNonEnvStruct{StringValue: "abc"}

		assert.Nil(env.Parse(&testStruct))
		assert.Equal(testStructExp, testStruct)
	})
}

func TestParseStructWithMissingEnvVariableName(t *testing.T) {
	assert := require.New(t)

	testStruct := struct {
		Skipped     int
		StringValue string `env:",required"`
	}{}
	testNestedStruct := struct {
		Skipped     int
		StructValue struct {
			StringValue string `env:",required"`
		}
	}{}

	withResetEnv(func() {
		assert.Equal(errors.New("env variable name cannot be empty"), env.Parse(&testStruct))
		assert.Equal(errors.New("env variable name cannot be empty"), env.Parse(&testNestedStruct))
	})
}

func TestParseStructUnknownOption(t *testing.T) {
	assert := require.New(t)

	testStruct := struct {
		Skipped     int
		StringValue string `env:"S1,abc"`
	}{}

	withResetEnv(func() {
		assert.Equal(errors.New("S1: unknown option abc"), env.Parse(&testStruct))
	})
}

func TestParseAliasType(t *testing.T) {
	assert := require.New(t)

	withResetEnv(func() {
		os.Setenv("S1", "a")
		testRuneStruct := struct {
			Value rune `env:"S1,type=rune"`
		}{}
		assert.Nil(env.Parse(&testRuneStruct))

		testByteStruct := struct {
			Value byte `env:"S1,type=byte"`
		}{}
		assert.Nil(env.Parse(&testByteStruct))

		testInvalidStruct := struct {
			Value interface{} `env:"S1,type=invalid"`
		}{}
		assert.Equal(errors.New("S1: invalid type \"type=invalid\", valid options are: \"byte\", \"rune\""), env.Parse(&testInvalidStruct))

		testInvalidValueStruct := struct {
			Value interface{} `env:"S1,type"`
		}{}
		assert.Equal(errors.New("S1: invalid type \"type\", valid options are: \"byte\", \"rune\""), env.Parse(&testInvalidValueStruct))
	})
}

func TestParse(t *testing.T) {
	assert := require.New(t)
	expectedEnv := genValidEnvStruct()

	withResetEnv(func() {
		testStruct := &testEnvStruct{}
		setupAllEnvVariables(expectedEnv, env.DefaultSeparator)
		assert.Nil(env.Parse(testStruct))
		assert.Equal(expectedEnv, testStruct)
	})
}

func TestParseRequired(t *testing.T) {

	assert := require.New(t)
	withResetEnv(func() {
		setupAllEnvVariables(genValidEnvStruct(), env.DefaultSeparator)
		for _, envKey := range []string{
			"GO_ENV_TEST_INT_VALUE",
			"GO_ENV_TEST_INT8_VALUE",
			"GO_ENV_TEST_INT16_VALUE",
			"GO_ENV_TEST_INT32_VALUE",
			"GO_ENV_TEST_INT64_VALUE",
			"GO_ENV_TEST_UINT_VALUE",
			"GO_ENV_TEST_UINT8_VALUE",
			"GO_ENV_TEST_UINT16_VALUE",
			"GO_ENV_TEST_UINT32_VALUE",
			"GO_ENV_TEST_UINT64_VALUE",
			"GO_ENV_TEST_FLOAT32_VALUE",
			"GO_ENV_TEST_FLOAT64_VALUE",
			"GO_ENV_TEST_INT_SLICE_VALUE",
			"GO_ENV_TEST_INT8_SLICE_VALUE",
			"GO_ENV_TEST_INT16_SLICE_VALUE",
			"GO_ENV_TEST_INT32_SLICE_VALUE",
			"GO_ENV_TEST_INT64_SLICE_VALUE",
			"GO_ENV_TEST_FLOAT32_SLICE_VALUE",
			"GO_ENV_TEST_FLOAT64_SLICE_VALUE",
			"GO_ENV_TEST_UINT_SLICE_VALUE",
			"GO_ENV_TEST_UINT8_SLICE_VALUE",
			"GO_ENV_TEST_UINT16_SLICE_VALUE",
			"GO_ENV_TEST_UINT32_SLICE_VALUE",
			"GO_ENV_TEST_UINT64_SLICE_VALUE",
			"GO_ENV_TEST_BOOL_VALUE",
			"GO_ENV_TEST_BOOL_SLICE_VALUE",
			"GO_ENV_TEST_STRING_VALUE",
			"GO_ENV_TEST_STRING_SLICE_VALUE",
			"GO_ENV_TEST_BYTE_SLICE_VALUE",
			"GO_ENV_TEST_BYTE_VALUE",
			"GO_ENV_TEST_RUNE_SLICE_VALUE",
			"GO_ENV_TEST_RUNE_VALUE",
			"GO_ENV_TEST_TIME_DURATION_VALUE",
			"GO_ENV_TEST_TIME_DURATION_SLICE_VALUE",
			"GO_ENV_TEST_TIME_VALUE",
			"GO_ENV_TEST_TIME_SLICE_VALUE",
		} {
			testStruct := &testEnvStructRequiredValues{}
			prev := os.Getenv(envKey)
			os.Unsetenv(envKey)

			err := env.Parse(testStruct)
			assert.Error(err)
			assert.Equal(errors.New(envKey+": value is required but was empty"), err)

			os.Setenv(envKey, prev)
		}
	})
}

func TestParseDefaults(t *testing.T) {
	expectedEnv := testEnvStruct{
		IntValue:           12,
		Int8Value:          22,
		Int16Value:         32,
		Int32Value:         42,
		Int64Value:         52,
		UintValue:          16,
		Uint8Value:         26,
		Uint16Value:        36,
		Uint32Value:        46,
		Uint64Value:        56,
		Float32Value:       12.0,
		Float64Value:       22.0,
		IntSliceValue:      []int{112},
		Int8SliceValue:     []int8{122},
		Int16SliceValue:    []int16{132},
		Int32SliceValue:    []int32{142},
		Int64SliceValue:    []int64{152},
		UintSliceValue:     []uint{117},
		Uint8SliceValue:    []uint8{127},
		Uint16SliceValue:   []uint16{137},
		Uint32SliceValue:   []uint32{147},
		Uint64SliceValue:   []uint64{157},
		Float32SliceValue:  []float32{112.0},
		Float64SliceValue:  []float64{122.0},
		BoolValue:          true,
		BoolSliceValue:     []bool{true},
		StringValue:        "exp/def/string",
		StringSliceValue:   []string{"exp/def/stringslice"},
		ByteValue:          byte('b'),
		ByteSliceValue:     []byte("exp/def/byteslice"),
		RuneValue:          'r',
		RuneSliceValue:     []rune("exp/def/runeslice"),
		DurationValue:      10 * time.Second,
		DurationSliceValue: []time.Duration{15 * time.Second},
		TimeValue:          time.Date(2006, 01, 02, 15, 04, 05, 0, time.UTC),
		TimeSliceValue:     []time.Time{time.Date(2006, 02, 02, 15, 04, 05, 0, time.UTC)},
		RegexpValue:        regexp.MustCompile("def"),
		RegexpSliceValue:   []*regexp.Regexp{regexp.MustCompile("def2")},
	}

	expectedNestedEnv := nestedTestEnvStructDefaultValues{
		testEnvStructDefaultValues: testEnvStructDefaultValues(expectedEnv),
		Nested: testEnvStructDefaultValues(expectedEnv),
	}

	assert := require.New(t)

	withResetEnv(func() {
		os.Clearenv()

		// test unnested
		testStruct := testEnvStructDefaultValues{}
		assert.Nil(env.Parse(&testStruct))
		assert.Equal(expectedEnv, testEnvStruct(testStruct))

		// test nested
		testNestedStruct := nestedTestEnvStructDefaultValues{}
		assert.Nil(env.Parse(&testNestedStruct))
		assert.Equal(expectedNestedEnv, testNestedStruct)
	})
}
func TestParseMultivalueSlices(t *testing.T) {
	assert := require.New(t)

	expectedEnv := testEnvStruct{
		IntValue:           13,
		Int8Value:          23,
		Int16Value:         33,
		Int32Value:         43,
		Int64Value:         53,
		UintValue:          14,
		Uint8Value:         24,
		Uint16Value:        34,
		Uint32Value:        44,
		Uint64Value:        54,
		Float32Value:       13.0,
		Float64Value:       23.0,
		IntSliceValue:      []int{113, 131},
		Int8SliceValue:     []int8{123, 121},
		Int16SliceValue:    []int16{133, 133},
		Int32SliceValue:    []int32{143, 143},
		Int64SliceValue:    []int64{153, 153},
		UintSliceValue:     []uint{213, 231},
		Uint8SliceValue:    []uint8{223, 221},
		Uint16SliceValue:   []uint16{233, 233},
		Uint32SliceValue:   []uint32{243, 243},
		Uint64SliceValue:   []uint64{253, 253},
		Float32SliceValue:  []float32{113.0, 131.1},
		Float64SliceValue:  []float64{123.0, 132.1},
		BoolValue:          true,
		BoolSliceValue:     []bool{true, true, false},
		StringValue:        "exp/multi/string",
		StringSliceValue:   []string{"exp/multi/stringslice1", "exp/multi/stringslice2"},
		ByteValue:          byte('l'),
		ByteSliceValue:     []byte("exp/multi/byteslice"),
		RuneValue:          'l',
		RuneSliceValue:     []rune("exp/multi/runeslice"),
		DurationValue:      time.Minute,
		DurationSliceValue: []time.Duration{time.Second, time.Minute},
		TimeValue:          testTime,
		TimeSliceValue:     []time.Time{testTimeSlice1, testTimeSlice2},
		RegexpValue:        testRegexp,
		RegexpSliceValue:   []*regexp.Regexp{testRegexpSlice1, testRegexpSlice2},
	}

	expectedNestedEnv := nestedTestEnvStructSeparator{
		testEnvStructSeparator: testEnvStructSeparator(expectedEnv),
		Nested:                 testEnvStructSeparator(expectedEnv),
	}

	withResetEnv(func() {
		setupAllEnvVariables(&expectedEnv, ";")

		// test unnested
		testStruct := testEnvStructSeparator{}
		assert.Nil(env.Parse(&testStruct))
		assert.Equal(expectedEnv, testEnvStruct(testStruct))

		// test nested
		testNestedStruct := nestedTestEnvStructSeparator{}
		assert.Nil(env.Parse(&testNestedStruct))
		assert.Equal(expectedNestedEnv, testNestedStruct)
	})
}

func TestParseInvalidEnvValues(t *testing.T) {
	assert := require.New(t)
	withResetEnv(func() {
		setupAllEnvVariables(genValidEnvStruct(), env.DefaultSeparator)

		for _, tc := range []struct {
			EnvKey        string
			EnvValue      string
			ExpectedError error
		}{
			{
				EnvKey:        "GO_ENV_TEST_INT_VALUE",
				EnvValue:      "abc",
				ExpectedError: errors.New("GO_ENV_TEST_INT_VALUE: strconv.ParseInt: parsing \"abc\": invalid syntax"),
			},
			{
				EnvKey:        "GO_ENV_TEST_INT8_VALUE",
				EnvValue:      "abc",
				ExpectedError: errors.New("GO_ENV_TEST_INT8_VALUE: strconv.ParseInt: parsing \"abc\": invalid syntax"),
			},
			{
				EnvKey:        "GO_ENV_TEST_INT16_VALUE",
				EnvValue:      "abc",
				ExpectedError: errors.New("GO_ENV_TEST_INT16_VALUE: strconv.ParseInt: parsing \"abc\": invalid syntax"),
			},
			{
				EnvKey:        "GO_ENV_TEST_INT32_VALUE",
				EnvValue:      "abc",
				ExpectedError: errors.New("GO_ENV_TEST_INT32_VALUE: strconv.ParseInt: parsing \"abc\": invalid syntax"),
			},
			{
				EnvKey:        "GO_ENV_TEST_INT64_VALUE",
				EnvValue:      "abc",
				ExpectedError: errors.New("GO_ENV_TEST_INT64_VALUE: strconv.ParseInt: parsing \"abc\": invalid syntax"),
			},
			{
				EnvKey:        "GO_ENV_TEST_UINT_VALUE",
				EnvValue:      "abc",
				ExpectedError: errors.New("GO_ENV_TEST_UINT_VALUE: strconv.ParseUint: parsing \"abc\": invalid syntax"),
			},
			{
				EnvKey:        "GO_ENV_TEST_UINT8_VALUE",
				EnvValue:      "abc",
				ExpectedError: errors.New("GO_ENV_TEST_UINT8_VALUE: strconv.ParseUint: parsing \"abc\": invalid syntax"),
			},
			{
				EnvKey:        "GO_ENV_TEST_UINT16_VALUE",
				EnvValue:      "abc",
				ExpectedError: errors.New("GO_ENV_TEST_UINT16_VALUE: strconv.ParseUint: parsing \"abc\": invalid syntax"),
			},
			{
				EnvKey:        "GO_ENV_TEST_UINT32_VALUE",
				EnvValue:      "abc",
				ExpectedError: errors.New("GO_ENV_TEST_UINT32_VALUE: strconv.ParseUint: parsing \"abc\": invalid syntax"),
			},
			{
				EnvKey:        "GO_ENV_TEST_UINT64_VALUE",
				EnvValue:      "abc",
				ExpectedError: errors.New("GO_ENV_TEST_UINT64_VALUE: strconv.ParseUint: parsing \"abc\": invalid syntax"),
			},
			{
				EnvKey:        "GO_ENV_TEST_FLOAT32_VALUE",
				EnvValue:      "abc",
				ExpectedError: errors.New("GO_ENV_TEST_FLOAT32_VALUE: strconv.ParseFloat: parsing \"abc\": invalid syntax"),
			},
			{
				EnvKey:        "GO_ENV_TEST_FLOAT64_VALUE",
				EnvValue:      "abc",
				ExpectedError: errors.New("GO_ENV_TEST_FLOAT64_VALUE: strconv.ParseFloat: parsing \"abc\": invalid syntax"),
			},
			{
				EnvKey:        "GO_ENV_TEST_INT_SLICE_VALUE",
				EnvValue:      "abc",
				ExpectedError: errors.New("GO_ENV_TEST_INT_SLICE_VALUE: strconv.ParseInt: parsing \"abc\": invalid syntax"),
			},
			{
				EnvKey:        "GO_ENV_TEST_INT8_SLICE_VALUE",
				EnvValue:      "abc",
				ExpectedError: errors.New("GO_ENV_TEST_INT8_SLICE_VALUE: strconv.ParseInt: parsing \"abc\": invalid syntax"),
			},
			{
				EnvKey:        "GO_ENV_TEST_INT16_SLICE_VALUE",
				EnvValue:      "abc",
				ExpectedError: errors.New("GO_ENV_TEST_INT16_SLICE_VALUE: strconv.ParseInt: parsing \"abc\": invalid syntax"),
			},
			{
				EnvKey:        "GO_ENV_TEST_INT32_SLICE_VALUE",
				EnvValue:      "abc",
				ExpectedError: errors.New("GO_ENV_TEST_INT32_SLICE_VALUE: strconv.ParseInt: parsing \"abc\": invalid syntax"),
			},
			{
				EnvKey:        "GO_ENV_TEST_INT64_SLICE_VALUE",
				EnvValue:      "abc",
				ExpectedError: errors.New("GO_ENV_TEST_INT64_SLICE_VALUE: strconv.ParseInt: parsing \"abc\": invalid syntax"),
			},
			{
				EnvKey:        "GO_ENV_TEST_UINT_SLICE_VALUE",
				EnvValue:      "abc",
				ExpectedError: errors.New("GO_ENV_TEST_UINT_SLICE_VALUE: strconv.ParseUint: parsing \"abc\": invalid syntax"),
			},
			{
				EnvKey:        "GO_ENV_TEST_UINT8_SLICE_VALUE",
				EnvValue:      "abc",
				ExpectedError: errors.New("GO_ENV_TEST_UINT8_SLICE_VALUE: strconv.ParseUint: parsing \"abc\": invalid syntax"),
			},
			{
				EnvKey:        "GO_ENV_TEST_UINT16_SLICE_VALUE",
				EnvValue:      "abc",
				ExpectedError: errors.New("GO_ENV_TEST_UINT16_SLICE_VALUE: strconv.ParseUint: parsing \"abc\": invalid syntax"),
			},
			{
				EnvKey:        "GO_ENV_TEST_UINT32_SLICE_VALUE",
				EnvValue:      "abc",
				ExpectedError: errors.New("GO_ENV_TEST_UINT32_SLICE_VALUE: strconv.ParseUint: parsing \"abc\": invalid syntax"),
			},
			{
				EnvKey:        "GO_ENV_TEST_UINT64_SLICE_VALUE",
				EnvValue:      "abc",
				ExpectedError: errors.New("GO_ENV_TEST_UINT64_SLICE_VALUE: strconv.ParseUint: parsing \"abc\": invalid syntax"),
			},
			{
				EnvKey:        "GO_ENV_TEST_FLOAT32_SLICE_VALUE",
				EnvValue:      "abc",
				ExpectedError: errors.New("GO_ENV_TEST_FLOAT32_SLICE_VALUE: strconv.ParseFloat: parsing \"abc\": invalid syntax"),
			},
			{
				EnvKey:        "GO_ENV_TEST_FLOAT64_SLICE_VALUE",
				EnvValue:      "abc",
				ExpectedError: errors.New("GO_ENV_TEST_FLOAT64_SLICE_VALUE: strconv.ParseFloat: parsing \"abc\": invalid syntax"),
			},
			{
				EnvKey:        "GO_ENV_TEST_BOOL_VALUE",
				EnvValue:      "abc",
				ExpectedError: errors.New("GO_ENV_TEST_BOOL_VALUE: strconv.ParseBool: parsing \"abc\": invalid syntax"),
			},
			{
				EnvKey:        "GO_ENV_TEST_BOOL_SLICE_VALUE",
				EnvValue:      "abc",
				ExpectedError: errors.New("GO_ENV_TEST_BOOL_SLICE_VALUE: strconv.ParseBool: parsing \"abc\": invalid syntax"),
			},
			{
				EnvKey:        "GO_ENV_TEST_RUNE_VALUE",
				EnvValue:      "abc",
				ExpectedError: errors.New("GO_ENV_TEST_RUNE_VALUE: rune must be a single character value"),
			},
			{
				EnvKey:        "GO_ENV_TEST_RUNE_SLICE_VALUE",
				EnvValue:      "abc,def",
				ExpectedError: errors.New("GO_ENV_TEST_RUNE_SLICE_VALUE: rune slice cannot have multiple values"),
			},
			{
				EnvKey:        "GO_ENV_TEST_BYTE_VALUE",
				EnvValue:      "abc",
				ExpectedError: errors.New("GO_ENV_TEST_BYTE_VALUE: byte must be a single character value"),
			},
			{
				EnvKey:        "GO_ENV_TEST_BYTE_SLICE_VALUE",
				EnvValue:      "abc,def",
				ExpectedError: errors.New("GO_ENV_TEST_BYTE_SLICE_VALUE: byte slice cannot have multiple values"),
			},
			{
				EnvKey:        "GO_ENV_TEST_TIME_DURATION_VALUE",
				EnvValue:      "abc",
				ExpectedError: errors.New("GO_ENV_TEST_TIME_DURATION_VALUE: time: invalid duration abc"),
			},
			{
				EnvKey:        "GO_ENV_TEST_TIME_DURATION_SLICE_VALUE",
				EnvValue:      "abc",
				ExpectedError: errors.New("GO_ENV_TEST_TIME_DURATION_SLICE_VALUE: time: invalid duration abc"),
			},
			{
				EnvKey:        "GO_ENV_TEST_TIME_VALUE",
				EnvValue:      "abc",
				ExpectedError: errors.New("GO_ENV_TEST_TIME_VALUE: parsing time \"abc\" as \"2006-01-02T15:04:05Z07:00\": cannot parse \"abc\" as \"2006\""),
			},
			{
				EnvKey:        "GO_ENV_TEST_TIME_SLICE_VALUE",
				EnvValue:      "abc",
				ExpectedError: errors.New("GO_ENV_TEST_TIME_SLICE_VALUE: parsing time \"abc\" as \"2006-01-02T15:04:05Z07:00\": cannot parse \"abc\" as \"2006\""),
			},
			{
				EnvKey:        "GO_ENV_TEST_REGEXP_POINTER_VALUE",
				EnvValue:      "*",
				ExpectedError: errors.New("GO_ENV_TEST_REGEXP_POINTER_VALUE: error parsing regexp: missing argument to repetition operator: `*`"),
			},
			{
				EnvKey:        "GO_ENV_TEST_REGEXP_POINTER_SLICE_VALUE",
				EnvValue:      "*",
				ExpectedError: errors.New("GO_ENV_TEST_REGEXP_POINTER_SLICE_VALUE: error parsing regexp: missing argument to repetition operator: `*`"),
			},
		} {
			testStruct := &testEnvStruct{}
			prev := os.Getenv(tc.EnvKey)
			os.Setenv(tc.EnvKey, tc.EnvValue)

			err := env.Parse(testStruct)
			assert.Equal(tc.ExpectedError, err)

			os.Setenv(tc.EnvKey, prev)
		}
	})
}

func withResetEnv(cb func()) {
	existing := os.Environ()
	defer func() {
		os.Clearenv()
		for _, val := range existing {
			envVar := strings.SplitN(val, "=", 2)
			os.Setenv(envVar[0], envVar[1])
		}
	}()
	cb()
}

// helper to setup all expected env variables
func setupAllEnvVariables(te *testEnvStruct, sep string) {
	os.Setenv("GO_ENV_TEST_INT_VALUE", strconv.Itoa(te.IntValue))
	os.Setenv("GO_ENV_TEST_INT8_VALUE", strconv.FormatInt(int64(te.Int8Value), 10))
	os.Setenv("GO_ENV_TEST_INT16_VALUE", strconv.FormatInt(int64(te.Int16Value), 10))
	os.Setenv("GO_ENV_TEST_INT32_VALUE", strconv.FormatInt(int64(te.Int32Value), 10))
	os.Setenv("GO_ENV_TEST_INT64_VALUE", strconv.FormatInt(te.Int64Value, 10))
	os.Setenv("GO_ENV_TEST_UINT_VALUE", strconv.FormatUint(uint64(te.UintValue), 10))
	os.Setenv("GO_ENV_TEST_UINT8_VALUE", strconv.FormatUint(uint64(te.Uint8Value), 10))
	os.Setenv("GO_ENV_TEST_UINT16_VALUE", strconv.FormatUint(uint64(te.Uint16Value), 10))
	os.Setenv("GO_ENV_TEST_UINT32_VALUE", strconv.FormatUint(uint64(te.Uint32Value), 10))
	os.Setenv("GO_ENV_TEST_UINT64_VALUE", strconv.FormatUint(te.Uint64Value, 10))
	os.Setenv("GO_ENV_TEST_FLOAT32_VALUE", fmt.Sprintf("%g", te.Float32Value))
	os.Setenv("GO_ENV_TEST_FLOAT64_VALUE", fmt.Sprintf("%g", te.Float64Value))
	os.Setenv("GO_ENV_TEST_INT_SLICE_VALUE", strings.Trim(strings.Join(strings.Fields(fmt.Sprint(te.IntSliceValue)), sep), "[]"))
	os.Setenv("GO_ENV_TEST_INT8_SLICE_VALUE", strings.Trim(strings.Join(strings.Fields(fmt.Sprint(te.Int8SliceValue)), sep), "[]"))
	os.Setenv("GO_ENV_TEST_INT16_SLICE_VALUE", strings.Trim(strings.Join(strings.Fields(fmt.Sprint(te.Int16SliceValue)), sep), "[]"))
	os.Setenv("GO_ENV_TEST_INT32_SLICE_VALUE", strings.Trim(strings.Join(strings.Fields(fmt.Sprint(te.Int32SliceValue)), sep), "[]"))
	os.Setenv("GO_ENV_TEST_INT64_SLICE_VALUE", strings.Trim(strings.Join(strings.Fields(fmt.Sprint(te.Int64SliceValue)), sep), "[]"))
	os.Setenv("GO_ENV_TEST_UINT_SLICE_VALUE", strings.Trim(strings.Join(strings.Fields(fmt.Sprint(te.UintSliceValue)), sep), "[]"))
	os.Setenv("GO_ENV_TEST_UINT8_SLICE_VALUE", strings.Trim(strings.Join(strings.Fields(fmt.Sprint(te.Uint8SliceValue)), sep), "[]"))
	os.Setenv("GO_ENV_TEST_UINT16_SLICE_VALUE", strings.Trim(strings.Join(strings.Fields(fmt.Sprint(te.Uint16SliceValue)), sep), "[]"))
	os.Setenv("GO_ENV_TEST_UINT32_SLICE_VALUE", strings.Trim(strings.Join(strings.Fields(fmt.Sprint(te.Uint32SliceValue)), sep), "[]"))
	os.Setenv("GO_ENV_TEST_UINT64_SLICE_VALUE", strings.Trim(strings.Join(strings.Fields(fmt.Sprint(te.Uint64SliceValue)), sep), "[]"))
	os.Setenv("GO_ENV_TEST_FLOAT32_SLICE_VALUE", strings.Trim(strings.Join(strings.Fields(fmt.Sprint(te.Float32SliceValue)), sep), "[]"))
	os.Setenv("GO_ENV_TEST_FLOAT64_SLICE_VALUE", strings.Trim(strings.Join(strings.Fields(fmt.Sprint(te.Float64SliceValue)), sep), "[]"))
	os.Setenv("GO_ENV_TEST_BOOL_VALUE", fmt.Sprintf("%t", te.BoolValue))
	os.Setenv("GO_ENV_TEST_BOOL_SLICE_VALUE", strings.Trim(strings.Join(strings.Fields(fmt.Sprint(te.BoolSliceValue)), sep), "[]"))
	os.Setenv("GO_ENV_TEST_STRING_VALUE", te.StringValue)
	os.Setenv("GO_ENV_TEST_STRING_SLICE_VALUE", strings.Join(te.StringSliceValue, sep))
	os.Setenv("GO_ENV_TEST_BYTE_SLICE_VALUE", string(te.ByteSliceValue))
	os.Setenv("GO_ENV_TEST_BYTE_VALUE", string(te.ByteValue))
	os.Setenv("GO_ENV_TEST_RUNE_SLICE_VALUE", string(te.RuneSliceValue))
	os.Setenv("GO_ENV_TEST_RUNE_VALUE", string(te.RuneValue))
	os.Setenv("GO_ENV_TEST_TIME_DURATION_VALUE", te.DurationValue.String())
	os.Setenv("GO_ENV_TEST_TIME_DURATION_SLICE_VALUE", strings.Trim(strings.Join(strings.Fields(fmt.Sprint(te.DurationSliceValue)), sep), "[]"))
	os.Setenv("GO_ENV_TEST_TIME_VALUE", te.TimeValue.Format(time.RFC3339))
	timeStrings := make([]string, len(te.TimeSliceValue))
	for idx, t := range te.TimeSliceValue {
		timeStrings[idx] = t.Format(time.RFC3339)
	}
	os.Setenv("GO_ENV_TEST_TIME_SLICE_VALUE", strings.Trim(strings.Join(timeStrings, sep), "[]"))
	os.Setenv("GO_ENV_TEST_REGEXP_POINTER_VALUE", te.RegexpValue.String())
	os.Setenv("GO_ENV_TEST_REGEXP_POINTER_SLICE_VALUE", strings.Trim(strings.Join(strings.Fields(fmt.Sprint(te.RegexpSliceValue)), sep), "[]"))
}

func genValidEnvStruct() *testEnvStruct {
	return &testEnvStruct{
		IntValue:           1,
		Int8Value:          2,
		Int16Value:         3,
		Int32Value:         4,
		Int64Value:         5,
		UintValue:          7,
		Uint8Value:         8,
		Uint16Value:        9,
		Uint32Value:        10,
		Uint64Value:        11,
		Float32Value:       1.0,
		Float64Value:       2.0,
		IntSliceValue:      []int{12},
		Int8SliceValue:     []int8{math.MaxInt8},
		Int16SliceValue:    []int16{math.MaxInt16},
		Int32SliceValue:    []int32{math.MaxInt32},
		Int64SliceValue:    []int64{math.MaxInt64},
		Float32SliceValue:  []float32{math.MaxFloat32},
		Float64SliceValue:  []float64{math.MaxFloat64},
		UintSliceValue:     []uint{13},
		Uint8SliceValue:    []uint8{math.MaxUint8},
		Uint16SliceValue:   []uint16{math.MaxUint16},
		Uint32SliceValue:   []uint32{math.MaxUint32},
		Uint64SliceValue:   []uint64{math.MaxUint64},
		BoolValue:          true,
		BoolSliceValue:     []bool{true},
		StringValue:        "abc",
		StringSliceValue:   []string{"def"},
		ByteValue:          byte('a'),
		ByteSliceValue:     []byte("ced"),
		RuneValue:          's',
		RuneSliceValue:     []rune{'a', 'c'},
		DurationValue:      time.Second,
		DurationSliceValue: []time.Duration{time.Second, 2 * time.Hour},
		TimeValue:          testTime,
		TimeSliceValue:     []time.Time{testTimeSlice1, testTimeSlice2},
		RegexpValue:        testRegexp,
		RegexpSliceValue:   []*regexp.Regexp{testRegexpSlice1, testRegexpSlice2},
	}
}

func sliceToString(sep, format string, v []interface{}) string {
	slice := make([]string, len(v))
	for idx, val := range v {
		slice[idx] = fmt.Sprintf(format, val)
	}
	return strings.Join(slice, sep)
}
