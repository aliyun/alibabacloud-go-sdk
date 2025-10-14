// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTagsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int64) *DescribeTagsResponseBody
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribeTagsResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *DescribeTagsResponseBody
	GetRequestId() *string
	SetTags(v []*DescribeTagsResponseBodyTags) *DescribeTagsResponseBody
	GetTags() []*DescribeTagsResponseBodyTags
	SetTotalCount(v int64) *DescribeTagsResponseBody
	GetTotalCount() *int64
}

type DescribeTagsResponseBody struct {
	// The page number. Pages start from page **1**. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 200.
	//
	// example:
	//
	// 200
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 89184F33-48A1-4401-9C0F-40E45DB091AB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The tags added to the resource.
	Tags []*DescribeTagsResponseBodyTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 5
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeTagsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeTagsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTagsResponseBody) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeTagsResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeTagsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeTagsResponseBody) GetTags() []*DescribeTagsResponseBodyTags {
	return s.Tags
}

func (s *DescribeTagsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeTagsResponseBody) SetPageNumber(v int64) *DescribeTagsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeTagsResponseBody) SetPageSize(v int64) *DescribeTagsResponseBody {
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

func (s *DescribeTagsResponseBody) SetTotalCount(v int64) *DescribeTagsResponseBody {
	s.TotalCount = &v
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
	// The key of tag N added to the resource.
	//
	// example:
	//
	// abc
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The values of tags added to the resource.
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
