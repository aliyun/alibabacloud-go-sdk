// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTag interface {
	dara.Model
	String() string
	GoString() string
	SetTagKey(v string) *Tag
	GetTagKey() *string
	SetTagValue(v string) *Tag
	GetTagValue() *string
}

type Tag struct {
	TagKey   *string `json:"tagKey,omitempty" xml:"tagKey,omitempty"`
	TagValue *string `json:"tagValue,omitempty" xml:"tagValue,omitempty"`
}

func (s Tag) String() string {
	return dara.Prettify(s)
}

func (s Tag) GoString() string {
	return s.String()
}

func (s *Tag) GetTagKey() *string {
	return s.TagKey
}

func (s *Tag) GetTagValue() *string {
	return s.TagValue
}

func (s *Tag) SetTagKey(v string) *Tag {
	s.TagKey = &v
	return s
}

func (s *Tag) SetTagValue(v string) *Tag {
	s.TagValue = &v
	return s
}

func (s *Tag) Validate() error {
	return dara.Validate(s)
}
