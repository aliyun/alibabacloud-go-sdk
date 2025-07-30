// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTagKeysResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCategory(v string) *DescribeTagKeysResponseBody
	GetCategory() *string
	SetPageNumber(v int32) *DescribeTagKeysResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeTagKeysResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeTagKeysResponseBody
	GetRequestId() *string
	SetTagKeys(v []*string) *DescribeTagKeysResponseBody
	GetTagKeys() []*string
	SetTotalCount(v int32) *DescribeTagKeysResponseBody
	GetTotalCount() *int32
}

type DescribeTagKeysResponseBody struct {
	// The type of the tag.
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
	// The number of tags returned per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// AD110813-9AD6-5F07-BFC8-4C841309****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The key of the tag.
	TagKeys []*string `json:"TagKeys,omitempty" xml:"TagKeys,omitempty" type:"Repeated"`
	// The total number of returned tag keys.
	//
	// example:
	//
	// 10
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeTagKeysResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeTagKeysResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTagKeysResponseBody) GetCategory() *string {
	return s.Category
}

func (s *DescribeTagKeysResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeTagKeysResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeTagKeysResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeTagKeysResponseBody) GetTagKeys() []*string {
	return s.TagKeys
}

func (s *DescribeTagKeysResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeTagKeysResponseBody) SetCategory(v string) *DescribeTagKeysResponseBody {
	s.Category = &v
	return s
}

func (s *DescribeTagKeysResponseBody) SetPageNumber(v int32) *DescribeTagKeysResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeTagKeysResponseBody) SetPageSize(v int32) *DescribeTagKeysResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeTagKeysResponseBody) SetRequestId(v string) *DescribeTagKeysResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTagKeysResponseBody) SetTagKeys(v []*string) *DescribeTagKeysResponseBody {
	s.TagKeys = v
	return s
}

func (s *DescribeTagKeysResponseBody) SetTotalCount(v int32) *DescribeTagKeysResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeTagKeysResponseBody) Validate() error {
	return dara.Validate(s)
}
