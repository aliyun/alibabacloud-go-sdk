// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUserPropertiesValue interface {
	dara.Model
	String() string
	GoString() string
	SetDataType(v string) *UserPropertiesValue
	GetDataType() *string
	SetStringValue(v string) *UserPropertiesValue
	GetStringValue() *string
	SetBinaryValue(v string) *UserPropertiesValue
	GetBinaryValue() *string
}

type UserPropertiesValue struct {
	DataType    *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	StringValue *string `json:"StringValue,omitempty" xml:"StringValue,omitempty"`
	BinaryValue *string `json:"BinaryValue,omitempty" xml:"BinaryValue,omitempty"`
}

func (s UserPropertiesValue) String() string {
	return dara.Prettify(s)
}

func (s UserPropertiesValue) GoString() string {
	return s.String()
}

func (s *UserPropertiesValue) GetDataType() *string {
	return s.DataType
}

func (s *UserPropertiesValue) GetStringValue() *string {
	return s.StringValue
}

func (s *UserPropertiesValue) GetBinaryValue() *string {
	return s.BinaryValue
}

func (s *UserPropertiesValue) SetDataType(v string) *UserPropertiesValue {
	s.DataType = &v
	return s
}

func (s *UserPropertiesValue) SetStringValue(v string) *UserPropertiesValue {
	s.StringValue = &v
	return s
}

func (s *UserPropertiesValue) SetBinaryValue(v string) *UserPropertiesValue {
	s.BinaryValue = &v
	return s
}

func (s *UserPropertiesValue) Validate() error {
	return dara.Validate(s)
}
