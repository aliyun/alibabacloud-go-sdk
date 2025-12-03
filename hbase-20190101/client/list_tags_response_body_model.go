// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTagsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListTagsResponseBody
	GetRequestId() *string
	SetTags(v *ListTagsResponseBodyTags) *ListTagsResponseBody
	GetTags() *ListTagsResponseBodyTags
}

type ListTagsResponseBody struct {
	// example:
	//
	// 36D1BE9B-3C4A-425B-947A-69E3D77999C4
	RequestId *string                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Tags      *ListTagsResponseBodyTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
}

func (s ListTagsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTagsResponseBody) GoString() string {
	return s.String()
}

func (s *ListTagsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTagsResponseBody) GetTags() *ListTagsResponseBodyTags {
	return s.Tags
}

func (s *ListTagsResponseBody) SetRequestId(v string) *ListTagsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTagsResponseBody) SetTags(v *ListTagsResponseBodyTags) *ListTagsResponseBody {
	s.Tags = v
	return s
}

func (s *ListTagsResponseBody) Validate() error {
	if s.Tags != nil {
		if err := s.Tags.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListTagsResponseBodyTags struct {
	Tag []*ListTagsResponseBodyTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListTagsResponseBodyTags) String() string {
	return dara.Prettify(s)
}

func (s ListTagsResponseBodyTags) GoString() string {
	return s.String()
}

func (s *ListTagsResponseBodyTags) GetTag() []*ListTagsResponseBodyTagsTag {
	return s.Tag
}

func (s *ListTagsResponseBodyTags) SetTag(v []*ListTagsResponseBodyTagsTag) *ListTagsResponseBodyTags {
	s.Tag = v
	return s
}

func (s *ListTagsResponseBodyTags) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListTagsResponseBodyTagsTag struct {
	// example:
	//
	// k1
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// example:
	//
	// v2
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s ListTagsResponseBodyTagsTag) String() string {
	return dara.Prettify(s)
}

func (s ListTagsResponseBodyTagsTag) GoString() string {
	return s.String()
}

func (s *ListTagsResponseBodyTagsTag) GetTagKey() *string {
	return s.TagKey
}

func (s *ListTagsResponseBodyTagsTag) GetTagValue() *string {
	return s.TagValue
}

func (s *ListTagsResponseBodyTagsTag) SetTagKey(v string) *ListTagsResponseBodyTagsTag {
	s.TagKey = &v
	return s
}

func (s *ListTagsResponseBodyTagsTag) SetTagValue(v string) *ListTagsResponseBodyTagsTag {
	s.TagValue = &v
	return s
}

func (s *ListTagsResponseBodyTagsTag) Validate() error {
	return dara.Validate(s)
}
