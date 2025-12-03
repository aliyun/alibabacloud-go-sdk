// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTagsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *DescribeTagsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeTagsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeTagsResponseBody
	GetRequestId() *string
	SetTagSets(v *DescribeTagsResponseBodyTagSets) *DescribeTagsResponseBody
	GetTagSets() *DescribeTagsResponseBodyTagSets
	SetTotalCount(v int32) *DescribeTagsResponseBody
	GetTotalCount() *int32
}

type DescribeTagsResponseBody struct {
	// The number of the returned page. Minimum value: 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page. Default value: 50. Maximum value: 100.
	//
	// example:
	//
	// 50
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 365F4154-92F6-4AE4-92F8-7FF34B540710
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The tags that are queried.
	TagSets *DescribeTagsResponseBodyTagSets `json:"TagSets,omitempty" xml:"TagSets,omitempty" type:"Struct"`
	// The number of instances returned.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeTagsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeTagsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTagsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeTagsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeTagsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeTagsResponseBody) GetTagSets() *DescribeTagsResponseBodyTagSets {
	return s.TagSets
}

func (s *DescribeTagsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeTagsResponseBody) SetPageNumber(v int32) *DescribeTagsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeTagsResponseBody) SetPageSize(v int32) *DescribeTagsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeTagsResponseBody) SetRequestId(v string) *DescribeTagsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTagsResponseBody) SetTagSets(v *DescribeTagsResponseBodyTagSets) *DescribeTagsResponseBody {
	s.TagSets = v
	return s
}

func (s *DescribeTagsResponseBody) SetTotalCount(v int32) *DescribeTagsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeTagsResponseBody) Validate() error {
	if s.TagSets != nil {
		if err := s.TagSets.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeTagsResponseBodyTagSets struct {
	TagSet []*DescribeTagsResponseBodyTagSetsTagSet `json:"TagSet,omitempty" xml:"TagSet,omitempty" type:"Repeated"`
}

func (s DescribeTagsResponseBodyTagSets) String() string {
	return dara.Prettify(s)
}

func (s DescribeTagsResponseBodyTagSets) GoString() string {
	return s.String()
}

func (s *DescribeTagsResponseBodyTagSets) GetTagSet() []*DescribeTagsResponseBodyTagSetsTagSet {
	return s.TagSet
}

func (s *DescribeTagsResponseBodyTagSets) SetTagSet(v []*DescribeTagsResponseBodyTagSetsTagSet) *DescribeTagsResponseBodyTagSets {
	s.TagSet = v
	return s
}

func (s *DescribeTagsResponseBodyTagSets) Validate() error {
	if s.TagSet != nil {
		for _, item := range s.TagSet {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeTagsResponseBodyTagSetsTagSet struct {
	// The number of instances to which the tag is added.
	//
	// example:
	//
	// 10
	InstanceCount *int32 `json:"InstanceCount,omitempty" xml:"InstanceCount,omitempty"`
	// The tag key.
	//
	// example:
	//
	// test
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The tag value.
	//
	// example:
	//
	// api
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s DescribeTagsResponseBodyTagSetsTagSet) String() string {
	return dara.Prettify(s)
}

func (s DescribeTagsResponseBodyTagSetsTagSet) GoString() string {
	return s.String()
}

func (s *DescribeTagsResponseBodyTagSetsTagSet) GetInstanceCount() *int32 {
	return s.InstanceCount
}

func (s *DescribeTagsResponseBodyTagSetsTagSet) GetTagKey() *string {
	return s.TagKey
}

func (s *DescribeTagsResponseBodyTagSetsTagSet) GetTagValue() *string {
	return s.TagValue
}

func (s *DescribeTagsResponseBodyTagSetsTagSet) SetInstanceCount(v int32) *DescribeTagsResponseBodyTagSetsTagSet {
	s.InstanceCount = &v
	return s
}

func (s *DescribeTagsResponseBodyTagSetsTagSet) SetTagKey(v string) *DescribeTagsResponseBodyTagSetsTagSet {
	s.TagKey = &v
	return s
}

func (s *DescribeTagsResponseBodyTagSetsTagSet) SetTagValue(v string) *DescribeTagsResponseBodyTagSetsTagSet {
	s.TagValue = &v
	return s
}

func (s *DescribeTagsResponseBodyTagSetsTagSet) Validate() error {
	return dara.Validate(s)
}
