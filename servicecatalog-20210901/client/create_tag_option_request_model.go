// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTagOptionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKey(v string) *CreateTagOptionRequest
	GetKey() *string
	SetValue(v string) *CreateTagOptionRequest
	GetValue() *string
}

type CreateTagOptionRequest struct {
	// The key of the tag option.
	//
	// The key can be up to 128 characters in length. It cannot start with `aliyun` or `acs:` and cannot contain `http://` or `https://`.
	//
	// This parameter is required.
	//
	// example:
	//
	// k1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag option.
	//
	// The value can be up to 128 characters in length. It cannot start with `acs:`and cannot contain `http://` or `https://`.
	//
	// This parameter is required.
	//
	// example:
	//
	// v1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateTagOptionRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateTagOptionRequest) GoString() string {
	return s.String()
}

func (s *CreateTagOptionRequest) GetKey() *string {
	return s.Key
}

func (s *CreateTagOptionRequest) GetValue() *string {
	return s.Value
}

func (s *CreateTagOptionRequest) SetKey(v string) *CreateTagOptionRequest {
	s.Key = &v
	return s
}

func (s *CreateTagOptionRequest) SetValue(v string) *CreateTagOptionRequest {
	s.Value = &v
	return s
}

func (s *CreateTagOptionRequest) Validate() error {
	return dara.Validate(s)
}
