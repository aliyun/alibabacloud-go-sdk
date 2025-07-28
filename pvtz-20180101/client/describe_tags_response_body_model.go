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
	SetTags(v []*DescribeTagsResponseBodyTags) *DescribeTagsResponseBody
	GetTags() []*DescribeTagsResponseBodyTags
	SetTotalCount(v int32) *DescribeTagsResponseBody
	GetTotalCount() *int32
}

type DescribeTagsResponseBody struct {
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9EA7F720-B7C0-45C1-9CF4-B6A5A1179B68
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The tags added to the resources.
	Tags []*DescribeTagsResponseBodyTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 200
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

func (s *DescribeTagsResponseBody) GetTags() []*DescribeTagsResponseBodyTags {
	return s.Tags
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

func (s *DescribeTagsResponseBody) SetTags(v []*DescribeTagsResponseBodyTags) *DescribeTagsResponseBody {
	s.Tags = v
	return s
}

func (s *DescribeTagsResponseBody) SetTotalCount(v int32) *DescribeTagsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeTagsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeTagsResponseBodyTags struct {
	// The key of tag N added to the resource.
	//
	// example:
	//
	// env
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The values of tags added to the resources.
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s DescribeTagsResponseBodyTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeTagsResponseBodyTags) GoString() string {
	return s.String()
}

func (s *DescribeTagsResponseBodyTags) GetKey() *string {
	return s.Key
}

func (s *DescribeTagsResponseBodyTags) GetValues() []*string {
	return s.Values
}

func (s *DescribeTagsResponseBodyTags) SetKey(v string) *DescribeTagsResponseBodyTags {
	s.Key = &v
	return s
}

func (s *DescribeTagsResponseBodyTags) SetValues(v []*string) *DescribeTagsResponseBodyTags {
	s.Values = v
	return s
}

func (s *DescribeTagsResponseBodyTags) Validate() error {
	return dara.Validate(s)
}
