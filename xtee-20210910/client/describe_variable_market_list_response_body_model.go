// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVariableMarketListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeVariableMarketListResponseBody
	GetRequestId() *string
	SetCurrentPage(v int32) *DescribeVariableMarketListResponseBody
	GetCurrentPage() *int32
	SetPageSize(v int32) *DescribeVariableMarketListResponseBody
	GetPageSize() *int32
	SetResultObject(v bool) *DescribeVariableMarketListResponseBody
	GetResultObject() *bool
	SetTotalItem(v int32) *DescribeVariableMarketListResponseBody
	GetTotalItem() *int32
	SetTotalPage(v int32) *DescribeVariableMarketListResponseBody
	GetTotalPage() *int32
}

type DescribeVariableMarketListResponseBody struct {
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
	// Page size, with a default value of 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// Return object
	//
	// example:
	//
	// true
	ResultObject *bool `json:"resultObject,omitempty" xml:"resultObject,omitempty"`
	// Total number of items
	//
	// example:
	//
	// 6
	TotalItem *int32 `json:"totalItem,omitempty" xml:"totalItem,omitempty"`
	// Total number of pages
	//
	// example:
	//
	// 9
	TotalPage *int32 `json:"totalPage,omitempty" xml:"totalPage,omitempty"`
}

func (s DescribeVariableMarketListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVariableMarketListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVariableMarketListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVariableMarketListResponseBody) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeVariableMarketListResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeVariableMarketListResponseBody) GetResultObject() *bool {
	return s.ResultObject
}

func (s *DescribeVariableMarketListResponseBody) GetTotalItem() *int32 {
	return s.TotalItem
}

func (s *DescribeVariableMarketListResponseBody) GetTotalPage() *int32 {
	return s.TotalPage
}

func (s *DescribeVariableMarketListResponseBody) SetRequestId(v string) *DescribeVariableMarketListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVariableMarketListResponseBody) SetCurrentPage(v int32) *DescribeVariableMarketListResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *DescribeVariableMarketListResponseBody) SetPageSize(v int32) *DescribeVariableMarketListResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeVariableMarketListResponseBody) SetResultObject(v bool) *DescribeVariableMarketListResponseBody {
	s.ResultObject = &v
	return s
}

func (s *DescribeVariableMarketListResponseBody) SetTotalItem(v int32) *DescribeVariableMarketListResponseBody {
	s.TotalItem = &v
	return s
}

func (s *DescribeVariableMarketListResponseBody) SetTotalPage(v int32) *DescribeVariableMarketListResponseBody {
	s.TotalPage = &v
	return s
}

func (s *DescribeVariableMarketListResponseBody) Validate() error {
	return dara.Validate(s)
}
