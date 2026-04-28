// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUserTag interface {
	dara.Model
	String() string
	GoString() string
	SetKey(v string) *UserTag
	GetKey() *string
	SetValue(v string) *UserTag
	GetValue() *string
}

type UserTag struct {
	// The tag key. This parameter must be specified and the tag key cannot contain number signs (#). The tag key and tag value cannot exceed 2,000 bytes in length in total.
	//
	// This parameter is required.
	//
	// example:
	//
	// k1
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// This parameter is required.
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s UserTag) String() string {
	return dara.Prettify(s)
}

func (s UserTag) GoString() string {
	return s.String()
}

func (s *UserTag) GetKey() *string {
	return s.Key
}

func (s *UserTag) GetValue() *string {
	return s.Value
}

func (s *UserTag) SetKey(v string) *UserTag {
	s.Key = &v
	return s
}

func (s *UserTag) SetValue(v string) *UserTag {
	s.Value = &v
	return s
}

func (s *UserTag) Validate() error {
	return dara.Validate(s)
}
