// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iField interface {
	dara.Model
	String() string
	GoString() string
	SetBlobValue(v string) *Field
	GetBlobValue() *string
	SetBooleanValue(v bool) *Field
	GetBooleanValue() *bool
	SetDoubleValue(v float64) *Field
	GetDoubleValue() *float64
	SetIsNull(v bool) *Field
	GetIsNull() *bool
	SetLongValue(v int64) *Field
	GetLongValue() *int64
	SetStringValue(v string) *Field
	GetStringValue() *string
}

type Field struct {
	BlobValue    *string  `json:"BlobValue,omitempty" xml:"BlobValue,omitempty"`
	BooleanValue *bool    `json:"BooleanValue,omitempty" xml:"BooleanValue,omitempty"`
	DoubleValue  *float64 `json:"DoubleValue,omitempty" xml:"DoubleValue,omitempty"`
	IsNull       *bool    `json:"IsNull,omitempty" xml:"IsNull,omitempty"`
	LongValue    *int64   `json:"LongValue,omitempty" xml:"LongValue,omitempty"`
	StringValue  *string  `json:"StringValue,omitempty" xml:"StringValue,omitempty"`
}

func (s Field) String() string {
	return dara.Prettify(s)
}

func (s Field) GoString() string {
	return s.String()
}

func (s *Field) GetBlobValue() *string {
	return s.BlobValue
}

func (s *Field) GetBooleanValue() *bool {
	return s.BooleanValue
}

func (s *Field) GetDoubleValue() *float64 {
	return s.DoubleValue
}

func (s *Field) GetIsNull() *bool {
	return s.IsNull
}

func (s *Field) GetLongValue() *int64 {
	return s.LongValue
}

func (s *Field) GetStringValue() *string {
	return s.StringValue
}

func (s *Field) SetBlobValue(v string) *Field {
	s.BlobValue = &v
	return s
}

func (s *Field) SetBooleanValue(v bool) *Field {
	s.BooleanValue = &v
	return s
}

func (s *Field) SetDoubleValue(v float64) *Field {
	s.DoubleValue = &v
	return s
}

func (s *Field) SetIsNull(v bool) *Field {
	s.IsNull = &v
	return s
}

func (s *Field) SetLongValue(v int64) *Field {
	s.LongValue = &v
	return s
}

func (s *Field) SetStringValue(v string) *Field {
	s.StringValue = &v
	return s
}

func (s *Field) Validate() error {
	return dara.Validate(s)
}
