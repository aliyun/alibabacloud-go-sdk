// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTagValuesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCategory(v string) *DescribeTagValuesResponseBody
	GetCategory() *string
	SetPageNumber(v int32) *DescribeTagValuesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeTagValuesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeTagValuesResponseBody
	GetRequestId() *string
	SetTagValues(v []*string) *DescribeTagValuesResponseBody
	GetTagValues() []*string
	SetTotalCount(v int32) *DescribeTagValuesResponseBody
	GetTotalCount() *int32
}

type DescribeTagValuesResponseBody struct {
	// The type of the tag key.
	//
	// example:
	//
	// Custom
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The start page of the returned pages.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of tag values returned on each page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// AD110813-9AD6-5F07-BFC8-4C841309****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The tag values that are associated with the tag key.
	TagValues []*string `json:"TagValues,omitempty" xml:"TagValues,omitempty" type:"Repeated"`
	// The total number of tag values that are associated with the tag key.
	//
	// example:
	//
	// 13
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeTagValuesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeTagValuesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTagValuesResponseBody) GetCategory() *string {
	return s.Category
}

func (s *DescribeTagValuesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeTagValuesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeTagValuesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeTagValuesResponseBody) GetTagValues() []*string {
	return s.TagValues
}

func (s *DescribeTagValuesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeTagValuesResponseBody) SetCategory(v string) *DescribeTagValuesResponseBody {
	s.Category = &v
	return s
}

func (s *DescribeTagValuesResponseBody) SetPageNumber(v int32) *DescribeTagValuesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeTagValuesResponseBody) SetPageSize(v int32) *DescribeTagValuesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeTagValuesResponseBody) SetRequestId(v string) *DescribeTagValuesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTagValuesResponseBody) SetTagValues(v []*string) *DescribeTagValuesResponseBody {
	s.TagValues = v
	return s
}

func (s *DescribeTagValuesResponseBody) SetTotalCount(v int32) *DescribeTagValuesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeTagValuesResponseBody) Validate() error {
	return dara.Validate(s)
}
