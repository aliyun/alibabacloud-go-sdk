// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAPIKeyAuthParameter interface {
	dara.Model
	String() string
	GoString() string
	SetEncrypted(v bool) *APIKeyAuthParameter
	GetEncrypted() *bool
	SetIn(v string) *APIKeyAuthParameter
	GetIn() *string
	SetKey(v string) *APIKeyAuthParameter
	GetKey() *string
	SetValue(v string) *APIKeyAuthParameter
	GetValue() *string
}

type APIKeyAuthParameter struct {
	Encrypted *bool `json:"encrypted,omitempty" xml:"encrypted,omitempty"`
	// example:
	//
	// header
	In *string `json:"in,omitempty" xml:"in,omitempty"`
	// example:
	//
	// Authorization
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// example:
	//
	// mock_value
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s APIKeyAuthParameter) String() string {
	return dara.Prettify(s)
}

func (s APIKeyAuthParameter) GoString() string {
	return s.String()
}

func (s *APIKeyAuthParameter) GetEncrypted() *bool {
	return s.Encrypted
}

func (s *APIKeyAuthParameter) GetIn() *string {
	return s.In
}

func (s *APIKeyAuthParameter) GetKey() *string {
	return s.Key
}

func (s *APIKeyAuthParameter) GetValue() *string {
	return s.Value
}

func (s *APIKeyAuthParameter) SetEncrypted(v bool) *APIKeyAuthParameter {
	s.Encrypted = &v
	return s
}

func (s *APIKeyAuthParameter) SetIn(v string) *APIKeyAuthParameter {
	s.In = &v
	return s
}

func (s *APIKeyAuthParameter) SetKey(v string) *APIKeyAuthParameter {
	s.Key = &v
	return s
}

func (s *APIKeyAuthParameter) SetValue(v string) *APIKeyAuthParameter {
	s.Value = &v
	return s
}

func (s *APIKeyAuthParameter) Validate() error {
	return dara.Validate(s)
}
