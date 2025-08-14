// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLoanTaskListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeLoanTaskListResponseBody
	GetRequestId() *string
	SetCurrentPage(v int32) *DescribeLoanTaskListResponseBody
	GetCurrentPage() *int32
	SetPageSize(v int32) *DescribeLoanTaskListResponseBody
	GetPageSize() *int32
	SetResultObject(v bool) *DescribeLoanTaskListResponseBody
	GetResultObject() *bool
	SetTotalItem(v int32) *DescribeLoanTaskListResponseBody
	GetTotalItem() *int32
	SetTotalPage(v int32) *DescribeLoanTaskListResponseBody
	GetTotalPage() *int32
}

type DescribeLoanTaskListResponseBody struct {
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
	// Page size, default value is 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// Return object.
	//
	// example:
	//
	// true
	ResultObject *bool `json:"resultObject,omitempty" xml:"resultObject,omitempty"`
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

func (s DescribeLoanTaskListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLoanTaskListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLoanTaskListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLoanTaskListResponseBody) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeLoanTaskListResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeLoanTaskListResponseBody) GetResultObject() *bool {
	return s.ResultObject
}

func (s *DescribeLoanTaskListResponseBody) GetTotalItem() *int32 {
	return s.TotalItem
}

func (s *DescribeLoanTaskListResponseBody) GetTotalPage() *int32 {
	return s.TotalPage
}

func (s *DescribeLoanTaskListResponseBody) SetRequestId(v string) *DescribeLoanTaskListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLoanTaskListResponseBody) SetCurrentPage(v int32) *DescribeLoanTaskListResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *DescribeLoanTaskListResponseBody) SetPageSize(v int32) *DescribeLoanTaskListResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeLoanTaskListResponseBody) SetResultObject(v bool) *DescribeLoanTaskListResponseBody {
	s.ResultObject = &v
	return s
}

func (s *DescribeLoanTaskListResponseBody) SetTotalItem(v int32) *DescribeLoanTaskListResponseBody {
	s.TotalItem = &v
	return s
}

func (s *DescribeLoanTaskListResponseBody) SetTotalPage(v int32) *DescribeLoanTaskListResponseBody {
	s.TotalPage = &v
	return s
}

func (s *DescribeLoanTaskListResponseBody) Validate() error {
	return dara.Validate(s)
}
