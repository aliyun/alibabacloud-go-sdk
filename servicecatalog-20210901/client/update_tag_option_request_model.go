// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTagOptionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetActive(v bool) *UpdateTagOptionRequest
	GetActive() *bool
	SetTagOptionId(v string) *UpdateTagOptionRequest
	GetTagOptionId() *string
	SetValue(v string) *UpdateTagOptionRequest
	GetValue() *string
}

type UpdateTagOptionRequest struct {
	// Specifies whether to enable the tag option. Valid values:
	//
	// 	- true (default)
	//
	// 	- false
	//
	// example:
	//
	// true
	Active *bool `json:"Active,omitempty" xml:"Active,omitempty"`
	// The ID of the tag option.
	//
	// This parameter is required.
	//
	// example:
	//
	// tag-bp1u6mdf3d****
	TagOptionId *string `json:"TagOptionId,omitempty" xml:"TagOptionId,omitempty"`
	// The value of the tag option.
	//
	// The value can be up to 128 characters in length. It cannot start with `acs:` and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// v1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateTagOptionRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateTagOptionRequest) GoString() string {
	return s.String()
}

func (s *UpdateTagOptionRequest) GetActive() *bool {
	return s.Active
}

func (s *UpdateTagOptionRequest) GetTagOptionId() *string {
	return s.TagOptionId
}

func (s *UpdateTagOptionRequest) GetValue() *string {
	return s.Value
}

func (s *UpdateTagOptionRequest) SetActive(v bool) *UpdateTagOptionRequest {
	s.Active = &v
	return s
}

func (s *UpdateTagOptionRequest) SetTagOptionId(v string) *UpdateTagOptionRequest {
	s.TagOptionId = &v
	return s
}

func (s *UpdateTagOptionRequest) SetValue(v string) *UpdateTagOptionRequest {
	s.Value = &v
	return s
}

func (s *UpdateTagOptionRequest) Validate() error {
	return dara.Validate(s)
}
