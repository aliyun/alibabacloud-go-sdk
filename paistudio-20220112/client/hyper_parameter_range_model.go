// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHyperParameterRange interface {
	dara.Model
	String() string
	GoString() string
	SetEnum(v []*string) *HyperParameterRange
	GetEnum() []*string
	SetExclusiveMaximum(v bool) *HyperParameterRange
	GetExclusiveMaximum() *bool
	SetExclusiveMinimum(v bool) *HyperParameterRange
	GetExclusiveMinimum() *bool
	SetMaxLength(v int64) *HyperParameterRange
	GetMaxLength() *int64
	SetMaximum(v string) *HyperParameterRange
	GetMaximum() *string
	SetMinLength(v int64) *HyperParameterRange
	GetMinLength() *int64
	SetMinimum(v string) *HyperParameterRange
	GetMinimum() *string
	SetPattern(v string) *HyperParameterRange
	GetPattern() *string
}

type HyperParameterRange struct {
	// Hyperparameter enumeration list.
	Enum []*string `json:"Enum,omitempty" xml:"Enum,omitempty" type:"Repeated"`
	// Whether the maximum value is exclusive.
	//
	// example:
	//
	// true
	ExclusiveMaximum *bool `json:"ExclusiveMaximum,omitempty" xml:"ExclusiveMaximum,omitempty"`
	// Whether the minimum value is exclusive.
	//
	// example:
	//
	// true
	ExclusiveMinimum *bool `json:"ExclusiveMinimum,omitempty" xml:"ExclusiveMinimum,omitempty"`
	// Maximum length.
	//
	// example:
	//
	// 30
	MaxLength *int64 `json:"MaxLength,omitempty" xml:"MaxLength,omitempty"`
	// Maximum value.
	//
	// example:
	//
	// 10
	Maximum *string `json:"Maximum,omitempty" xml:"Maximum,omitempty"`
	// Minimum length.
	//
	// example:
	//
	// 1
	MinLength *int64 `json:"MinLength,omitempty" xml:"MinLength,omitempty"`
	// Minimum value.
	//
	// example:
	//
	// 0
	Minimum *string `json:"Minimum,omitempty" xml:"Minimum,omitempty"`
	// Regular expression.
	//
	// example:
	//
	// ^\\+?[1-9][0-9]*$
	Pattern *string `json:"Pattern,omitempty" xml:"Pattern,omitempty"`
}

func (s HyperParameterRange) String() string {
	return dara.Prettify(s)
}

func (s HyperParameterRange) GoString() string {
	return s.String()
}

func (s *HyperParameterRange) GetEnum() []*string {
	return s.Enum
}

func (s *HyperParameterRange) GetExclusiveMaximum() *bool {
	return s.ExclusiveMaximum
}

func (s *HyperParameterRange) GetExclusiveMinimum() *bool {
	return s.ExclusiveMinimum
}

func (s *HyperParameterRange) GetMaxLength() *int64 {
	return s.MaxLength
}

func (s *HyperParameterRange) GetMaximum() *string {
	return s.Maximum
}

func (s *HyperParameterRange) GetMinLength() *int64 {
	return s.MinLength
}

func (s *HyperParameterRange) GetMinimum() *string {
	return s.Minimum
}

func (s *HyperParameterRange) GetPattern() *string {
	return s.Pattern
}

func (s *HyperParameterRange) SetEnum(v []*string) *HyperParameterRange {
	s.Enum = v
	return s
}

func (s *HyperParameterRange) SetExclusiveMaximum(v bool) *HyperParameterRange {
	s.ExclusiveMaximum = &v
	return s
}

func (s *HyperParameterRange) SetExclusiveMinimum(v bool) *HyperParameterRange {
	s.ExclusiveMinimum = &v
	return s
}

func (s *HyperParameterRange) SetMaxLength(v int64) *HyperParameterRange {
	s.MaxLength = &v
	return s
}

func (s *HyperParameterRange) SetMaximum(v string) *HyperParameterRange {
	s.Maximum = &v
	return s
}

func (s *HyperParameterRange) SetMinLength(v int64) *HyperParameterRange {
	s.MinLength = &v
	return s
}

func (s *HyperParameterRange) SetMinimum(v string) *HyperParameterRange {
	s.Minimum = &v
	return s
}

func (s *HyperParameterRange) SetPattern(v string) *HyperParameterRange {
	s.Pattern = &v
	return s
}

func (s *HyperParameterRange) Validate() error {
	return dara.Validate(s)
}
