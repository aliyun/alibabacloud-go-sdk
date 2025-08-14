// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGroupAccountPageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeGroupAccountPageResponseBody
	GetRequestId() *string
	SetCurrentPage(v int32) *DescribeGroupAccountPageResponseBody
	GetCurrentPage() *int32
	SetData(v bool) *DescribeGroupAccountPageResponseBody
	GetData() *bool
	SetPageSize(v int32) *DescribeGroupAccountPageResponseBody
	GetPageSize() *int32
	SetTotalItem(v int32) *DescribeGroupAccountPageResponseBody
	GetTotalItem() *int32
	SetTotalPage(v int32) *DescribeGroupAccountPageResponseBody
	GetTotalPage() *int32
}

type DescribeGroupAccountPageResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// A32FE941-35F2-5378-B37C-4B8FDB16F094
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Current page number.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"currentPage,omitempty" xml:"currentPage,omitempty"`
	// Returned data object.
	//
	// example:
	//
	// true
	Data *bool `json:"data,omitempty" xml:"data,omitempty"`
	// Page size, default value is 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// Total number of items.
	//
	// example:
	//
	// 6
	TotalItem *int32 `json:"totalItem,omitempty" xml:"totalItem,omitempty"`
	// Total number of pages.
	//
	// example:
	//
	// 1
	TotalPage *int32 `json:"totalPage,omitempty" xml:"totalPage,omitempty"`
}

func (s DescribeGroupAccountPageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeGroupAccountPageResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeGroupAccountPageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeGroupAccountPageResponseBody) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeGroupAccountPageResponseBody) GetData() *bool {
	return s.Data
}

func (s *DescribeGroupAccountPageResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeGroupAccountPageResponseBody) GetTotalItem() *int32 {
	return s.TotalItem
}

func (s *DescribeGroupAccountPageResponseBody) GetTotalPage() *int32 {
	return s.TotalPage
}

func (s *DescribeGroupAccountPageResponseBody) SetRequestId(v string) *DescribeGroupAccountPageResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeGroupAccountPageResponseBody) SetCurrentPage(v int32) *DescribeGroupAccountPageResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *DescribeGroupAccountPageResponseBody) SetData(v bool) *DescribeGroupAccountPageResponseBody {
	s.Data = &v
	return s
}

func (s *DescribeGroupAccountPageResponseBody) SetPageSize(v int32) *DescribeGroupAccountPageResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeGroupAccountPageResponseBody) SetTotalItem(v int32) *DescribeGroupAccountPageResponseBody {
	s.TotalItem = &v
	return s
}

func (s *DescribeGroupAccountPageResponseBody) SetTotalPage(v int32) *DescribeGroupAccountPageResponseBody {
	s.TotalPage = &v
	return s
}

func (s *DescribeGroupAccountPageResponseBody) Validate() error {
	return dara.Validate(s)
}
