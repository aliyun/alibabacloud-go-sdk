// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iKeyValuePair interface {
	dara.Model
	String() string
	GoString() string
	SetKey(v string) *KeyValuePair
	GetKey() *string
	SetValue(v string) *KeyValuePair
	GetValue() *string
}

type KeyValuePair struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s KeyValuePair) String() string {
	return dara.Prettify(s)
}

func (s KeyValuePair) GoString() string {
	return s.String()
}

func (s *KeyValuePair) GetKey() *string {
	return s.Key
}

func (s *KeyValuePair) GetValue() *string {
	return s.Value
}

func (s *KeyValuePair) SetKey(v string) *KeyValuePair {
	s.Key = &v
	return s
}

func (s *KeyValuePair) SetValue(v string) *KeyValuePair {
	s.Value = &v
	return s
}

func (s *KeyValuePair) Validate() error {
	return dara.Validate(s)
}
