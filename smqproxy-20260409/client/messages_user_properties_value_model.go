// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMessagesUserPropertiesValue interface {
	dara.Model
	String() string
	GoString() string
	SetDataType(v string) *MessagesUserPropertiesValue
	GetDataType() *string
	SetStringValue(v string) *MessagesUserPropertiesValue
	GetStringValue() *string
	SetBinaryValue(v string) *MessagesUserPropertiesValue
	GetBinaryValue() *string
}

type MessagesUserPropertiesValue struct {
	DataType    *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	StringValue *string `json:"StringValue,omitempty" xml:"StringValue,omitempty"`
	BinaryValue *string `json:"BinaryValue,omitempty" xml:"BinaryValue,omitempty"`
}

func (s MessagesUserPropertiesValue) String() string {
	return dara.Prettify(s)
}

func (s MessagesUserPropertiesValue) GoString() string {
	return s.String()
}

func (s *MessagesUserPropertiesValue) GetDataType() *string {
	return s.DataType
}

func (s *MessagesUserPropertiesValue) GetStringValue() *string {
	return s.StringValue
}

func (s *MessagesUserPropertiesValue) GetBinaryValue() *string {
	return s.BinaryValue
}

func (s *MessagesUserPropertiesValue) SetDataType(v string) *MessagesUserPropertiesValue {
	s.DataType = &v
	return s
}

func (s *MessagesUserPropertiesValue) SetStringValue(v string) *MessagesUserPropertiesValue {
	s.StringValue = &v
	return s
}

func (s *MessagesUserPropertiesValue) SetBinaryValue(v string) *MessagesUserPropertiesValue {
	s.BinaryValue = &v
	return s
}

func (s *MessagesUserPropertiesValue) Validate() error {
	return dara.Validate(s)
}
