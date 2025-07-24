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
	// 标签键。必填参数，不允许为空字符串。最多支持128个字符，不能以aliyun和acs:开头，不能包含http://或https://。
	//
	// This parameter is required.
	//
	// example:
	//
	// department
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// 标签值。非必填，可以为空字符串。最多支持128个字符，不能以acs:开头，不能包含http://或者https://。
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
