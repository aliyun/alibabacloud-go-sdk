// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResourceTag interface {
	dara.Model
	String() string
	GoString() string
	SetKey(v string) *ResourceTag
	GetKey() *string
	SetValue(v string) *ResourceTag
	GetValue() *string
}

type ResourceTag struct {
	// The tag key, up to 128 characters in length.
	//
	// example:
	//
	// team
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// The tag value, up to 256 characters in length.
	//
	// example:
	//
	// recommendation
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s ResourceTag) String() string {
	return dara.Prettify(s)
}

func (s ResourceTag) GoString() string {
	return s.String()
}

func (s *ResourceTag) GetKey() *string {
	return s.Key
}

func (s *ResourceTag) GetValue() *string {
	return s.Value
}

func (s *ResourceTag) SetKey(v string) *ResourceTag {
	s.Key = &v
	return s
}

func (s *ResourceTag) SetValue(v string) *ResourceTag {
	s.Value = &v
	return s
}

func (s *ResourceTag) Validate() error {
	return dara.Validate(s)
}
