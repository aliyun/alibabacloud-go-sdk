// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListResourceTagsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListResourceTagsResponseBody
	GetRequestId() *string
	SetTags(v *ListResourceTagsResponseBodyTags) *ListResourceTagsResponseBody
	GetTags() *ListResourceTagsResponseBodyTags
}

type ListResourceTagsResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 4162a6af-bc99-40b3-a552-89dcc8aaf7c8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The tags of the CMK.
	Tags *ListResourceTagsResponseBodyTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
}

func (s ListResourceTagsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListResourceTagsResponseBody) GoString() string {
	return s.String()
}

func (s *ListResourceTagsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListResourceTagsResponseBody) GetTags() *ListResourceTagsResponseBodyTags {
	return s.Tags
}

func (s *ListResourceTagsResponseBody) SetRequestId(v string) *ListResourceTagsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListResourceTagsResponseBody) SetTags(v *ListResourceTagsResponseBodyTags) *ListResourceTagsResponseBody {
	s.Tags = v
	return s
}

func (s *ListResourceTagsResponseBody) Validate() error {
	if s.Tags != nil {
		if err := s.Tags.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListResourceTagsResponseBodyTags struct {
	Tag []*ListResourceTagsResponseBodyTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListResourceTagsResponseBodyTags) String() string {
	return dara.Prettify(s)
}

func (s ListResourceTagsResponseBodyTags) GoString() string {
	return s.String()
}

func (s *ListResourceTagsResponseBodyTags) GetTag() []*ListResourceTagsResponseBodyTagsTag {
	return s.Tag
}

func (s *ListResourceTagsResponseBodyTags) SetTag(v []*ListResourceTagsResponseBodyTagsTag) *ListResourceTagsResponseBodyTags {
	s.Tag = v
	return s
}

func (s *ListResourceTagsResponseBodyTags) Validate() error {
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

type ListResourceTagsResponseBodyTagsTag struct {
	// The globally unique ID of the CMK.
	//
	// example:
	//
	// 33caea95-c3e5-4b3e-a9c6-cec76e4e****
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	// The tag key.
	//
	// example:
	//
	// Project
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The tag value.
	//
	// example:
	//
	// Test
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s ListResourceTagsResponseBodyTagsTag) String() string {
	return dara.Prettify(s)
}

func (s ListResourceTagsResponseBodyTagsTag) GoString() string {
	return s.String()
}

func (s *ListResourceTagsResponseBodyTagsTag) GetKeyId() *string {
	return s.KeyId
}

func (s *ListResourceTagsResponseBodyTagsTag) GetTagKey() *string {
	return s.TagKey
}

func (s *ListResourceTagsResponseBodyTagsTag) GetTagValue() *string {
	return s.TagValue
}

func (s *ListResourceTagsResponseBodyTagsTag) SetKeyId(v string) *ListResourceTagsResponseBodyTagsTag {
	s.KeyId = &v
	return s
}

func (s *ListResourceTagsResponseBodyTagsTag) SetTagKey(v string) *ListResourceTagsResponseBodyTagsTag {
	s.TagKey = &v
	return s
}

func (s *ListResourceTagsResponseBodyTagsTag) SetTagValue(v string) *ListResourceTagsResponseBodyTagsTag {
	s.TagValue = &v
	return s
}

func (s *ListResourceTagsResponseBodyTagsTag) Validate() error {
	return dara.Validate(s)
}
