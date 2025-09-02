// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUserEntityTag interface {
	dara.Model
	String() string
	GoString() string
	SetTagKey(v string) *UserEntityTag
	GetTagKey() *string
	SetTagValue(v string) *UserEntityTag
	GetTagValue() *string
}

type UserEntityTag struct {
	// example:
	//
	// priority
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// example:
	//
	// p1
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s UserEntityTag) String() string {
	return dara.Prettify(s)
}

func (s UserEntityTag) GoString() string {
	return s.String()
}

func (s *UserEntityTag) GetTagKey() *string {
	return s.TagKey
}

func (s *UserEntityTag) GetTagValue() *string {
	return s.TagValue
}

func (s *UserEntityTag) SetTagKey(v string) *UserEntityTag {
	s.TagKey = &v
	return s
}

func (s *UserEntityTag) SetTagValue(v string) *UserEntityTag {
	s.TagValue = &v
	return s
}

func (s *UserEntityTag) Validate() error {
	return dara.Validate(s)
}
