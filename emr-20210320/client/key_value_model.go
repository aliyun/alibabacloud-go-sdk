// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iKeyValue interface {
	dara.Model
	String() string
	GoString() string
	SetKey(v string) *KeyValue
	GetKey() *string
	SetValue(v string) *KeyValue
	GetValue() *string
}

type KeyValue struct {
	// 键。
	//
	// example:
	//
	// key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// 值。
	//
	// example:
	//
	// value
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s KeyValue) String() string {
	return dara.Prettify(s)
}

func (s KeyValue) GoString() string {
	return s.String()
}

func (s *KeyValue) GetKey() *string {
	return s.Key
}

func (s *KeyValue) GetValue() *string {
	return s.Value
}

func (s *KeyValue) SetKey(v string) *KeyValue {
	s.Key = &v
	return s
}

func (s *KeyValue) SetValue(v string) *KeyValue {
	s.Value = &v
	return s
}

func (s *KeyValue) Validate() error {
	return dara.Validate(s)
}
