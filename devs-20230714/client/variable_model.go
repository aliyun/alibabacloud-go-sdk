// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVariable interface {
	dara.Model
	String() string
	GoString() string
	SetEncrypted(v bool) *Variable
	GetEncrypted() *bool
	SetSensitive(v bool) *Variable
	GetSensitive() *bool
	SetValue(v interface{}) *Variable
	GetValue() interface{}
}

type Variable struct {
	// example:
	//
	// false
	Encrypted *bool `json:"encrypted,omitempty" xml:"encrypted,omitempty"`
	// example:
	//
	// false
	Sensitive *bool `json:"sensitive,omitempty" xml:"sensitive,omitempty"`
	// example:
	//
	// object_value
	Value interface{} `json:"value,omitempty" xml:"value,omitempty"`
}

func (s Variable) String() string {
	return dara.Prettify(s)
}

func (s Variable) GoString() string {
	return s.String()
}

func (s *Variable) GetEncrypted() *bool {
	return s.Encrypted
}

func (s *Variable) GetSensitive() *bool {
	return s.Sensitive
}

func (s *Variable) GetValue() interface{} {
	return s.Value
}

func (s *Variable) SetEncrypted(v bool) *Variable {
	s.Encrypted = &v
	return s
}

func (s *Variable) SetSensitive(v bool) *Variable {
	s.Sensitive = &v
	return s
}

func (s *Variable) SetValue(v interface{}) *Variable {
	s.Value = v
	return s
}

func (s *Variable) Validate() error {
	return dara.Validate(s)
}
