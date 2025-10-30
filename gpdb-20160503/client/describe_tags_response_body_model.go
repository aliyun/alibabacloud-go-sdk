// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTagsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeTagsResponseBody
	GetRequestId() *string
	SetTags(v []*DescribeTagsResponseBodyTags) *DescribeTagsResponseBody
	GetTags() []*DescribeTagsResponseBodyTags
}

type DescribeTagsResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// A29EC547-B392-4340-AA4F-7C0A7B626E74
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The queried tags.
	Tags []*DescribeTagsResponseBodyTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s DescribeTagsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeTagsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTagsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeTagsResponseBody) GetTags() []*DescribeTagsResponseBodyTags {
	return s.Tags
}

func (s *DescribeTagsResponseBody) SetRequestId(v string) *DescribeTagsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTagsResponseBody) SetTags(v []*DescribeTagsResponseBodyTags) *DescribeTagsResponseBody {
	s.Tags = v
	return s
}

func (s *DescribeTagsResponseBody) Validate() error {
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeTagsResponseBodyTags struct {
	// The tag key.
	//
	// example:
	//
	// user
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The tag value.
	//
	// example:
	//
	// test
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s DescribeTagsResponseBodyTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeTagsResponseBodyTags) GoString() string {
	return s.String()
}

func (s *DescribeTagsResponseBodyTags) GetTagKey() *string {
	return s.TagKey
}

func (s *DescribeTagsResponseBodyTags) GetTagValue() *string {
	return s.TagValue
}

func (s *DescribeTagsResponseBodyTags) SetTagKey(v string) *DescribeTagsResponseBodyTags {
	s.TagKey = &v
	return s
}

func (s *DescribeTagsResponseBodyTags) SetTagValue(v string) *DescribeTagsResponseBodyTags {
	s.TagValue = &v
	return s
}

func (s *DescribeTagsResponseBodyTags) Validate() error {
	return dara.Validate(s)
}
