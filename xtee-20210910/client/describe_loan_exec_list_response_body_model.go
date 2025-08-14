// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLoanExecListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeLoanExecListResponseBody
	GetRequestId() *string
	SetCurrentPage(v int32) *DescribeLoanExecListResponseBody
	GetCurrentPage() *int32
	SetPageSize(v int32) *DescribeLoanExecListResponseBody
	GetPageSize() *int32
	SetResultObject(v bool) *DescribeLoanExecListResponseBody
	GetResultObject() *bool
	SetTotalItem(v int32) *DescribeLoanExecListResponseBody
	GetTotalItem() *int32
	SetTotalPage(v int32) *DescribeLoanExecListResponseBody
	GetTotalPage() *int32
}

type DescribeLoanExecListResponseBody struct {
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

func (s DescribeLoanExecListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLoanExecListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLoanExecListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLoanExecListResponseBody) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeLoanExecListResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeLoanExecListResponseBody) GetResultObject() *bool {
	return s.ResultObject
}

func (s *DescribeLoanExecListResponseBody) GetTotalItem() *int32 {
	return s.TotalItem
}

func (s *DescribeLoanExecListResponseBody) GetTotalPage() *int32 {
	return s.TotalPage
}

func (s *DescribeLoanExecListResponseBody) SetRequestId(v string) *DescribeLoanExecListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLoanExecListResponseBody) SetCurrentPage(v int32) *DescribeLoanExecListResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *DescribeLoanExecListResponseBody) SetPageSize(v int32) *DescribeLoanExecListResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeLoanExecListResponseBody) SetResultObject(v bool) *DescribeLoanExecListResponseBody {
	s.ResultObject = &v
	return s
}

func (s *DescribeLoanExecListResponseBody) SetTotalItem(v int32) *DescribeLoanExecListResponseBody {
	s.TotalItem = &v
	return s
}

func (s *DescribeLoanExecListResponseBody) SetTotalPage(v int32) *DescribeLoanExecListResponseBody {
	s.TotalPage = &v
	return s
}

func (s *DescribeLoanExecListResponseBody) Validate() error {
	return dara.Validate(s)
}
