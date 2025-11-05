// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserTagsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeUserTagsResponseBody
	GetRequestId() *string
	SetTags(v []*DescribeUserTagsResponseBodyTags) *DescribeUserTagsResponseBody
	GetTags() []*DescribeUserTagsResponseBodyTags
}

type DescribeUserTagsResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 34AB41F1-04A5-496F-8C8D-634BDBE6A9FB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The tags.
	Tags []*DescribeUserTagsResponseBodyTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s DescribeUserTagsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserTagsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUserTagsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeUserTagsResponseBody) GetTags() []*DescribeUserTagsResponseBodyTags {
	return s.Tags
}

func (s *DescribeUserTagsResponseBody) SetRequestId(v string) *DescribeUserTagsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeUserTagsResponseBody) SetTags(v []*DescribeUserTagsResponseBodyTags) *DescribeUserTagsResponseBody {
	s.Tags = v
	return s
}

func (s *DescribeUserTagsResponseBody) Validate() error {
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

type DescribeUserTagsResponseBodyTags struct {
	// The key of a tag.
	//
	// example:
	//
	// env
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag values returned.
	Value []*string `json:"Value,omitempty" xml:"Value,omitempty" type:"Repeated"`
}

func (s DescribeUserTagsResponseBodyTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserTagsResponseBodyTags) GoString() string {
	return s.String()
}

func (s *DescribeUserTagsResponseBodyTags) GetKey() *string {
	return s.Key
}

func (s *DescribeUserTagsResponseBodyTags) GetValue() []*string {
	return s.Value
}

func (s *DescribeUserTagsResponseBodyTags) SetKey(v string) *DescribeUserTagsResponseBodyTags {
	s.Key = &v
	return s
}

func (s *DescribeUserTagsResponseBodyTags) SetValue(v []*string) *DescribeUserTagsResponseBodyTags {
	s.Value = v
	return s
}

func (s *DescribeUserTagsResponseBodyTags) Validate() error {
	return dara.Validate(s)
}
