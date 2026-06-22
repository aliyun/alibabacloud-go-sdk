// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTag interface {
	dara.Model
	String() string
	GoString() string
	SetKey(v string) *Tag
	GetKey() *string
	SetValue(v string) *Tag
	GetValue() *string
}

type Tag struct {
	// The key of the tag. This parameter is required and cannot be an empty string. The key can be up to 128 characters in length. It cannot start with `aliyun` or `acs:` and cannot contain `http://` or `https://`.
	//
	// This parameter is required.
	//
	// example:
	//
	// department
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag This parameter is optional and can be an empty string. The value can be up to 128 characters in length. It cannot start with `acs:` and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// IT
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s Tag) String() string {
	return dara.Prettify(s)
}

func (s Tag) GoString() string {
	return s.String()
}

func (s *Tag) GetKey() *string {
	return s.Key
}

func (s *Tag) GetValue() *string {
	return s.Value
}

func (s *Tag) SetKey(v string) *Tag {
	s.Key = &v
	return s
}

func (s *Tag) SetValue(v string) *Tag {
	s.Value = &v
	return s
}

func (s *Tag) Validate() error {
	return dara.Validate(s)
}
