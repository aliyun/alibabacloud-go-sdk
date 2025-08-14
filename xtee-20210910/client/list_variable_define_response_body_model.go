// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVariableDefineResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListVariableDefineResponseBody
	GetRequestId() *string
	SetCurrentPage(v int32) *ListVariableDefineResponseBody
	GetCurrentPage() *int32
	SetPageSize(v int32) *ListVariableDefineResponseBody
	GetPageSize() *int32
	SetResultObject(v bool) *ListVariableDefineResponseBody
	GetResultObject() *bool
	SetTotalItem(v int32) *ListVariableDefineResponseBody
	GetTotalItem() *int32
	SetTotalPage(v int32) *ListVariableDefineResponseBody
	GetTotalPage() *int32
}

type ListVariableDefineResponseBody struct {
	// Request ID
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
	// Page size, default value is 10
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// Returned object
	//
	// example:
	//
	// true
	ResultObject *bool `json:"resultObject,omitempty" xml:"resultObject,omitempty"`
	// Total items
	//
	// example:
	//
	// 27
	TotalItem *int32 `json:"totalItem,omitempty" xml:"totalItem,omitempty"`
	// Total pages.
	//
	// example:
	//
	// 4
	TotalPage *int32 `json:"totalPage,omitempty" xml:"totalPage,omitempty"`
}

func (s ListVariableDefineResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListVariableDefineResponseBody) GoString() string {
	return s.String()
}

func (s *ListVariableDefineResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListVariableDefineResponseBody) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListVariableDefineResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListVariableDefineResponseBody) GetResultObject() *bool {
	return s.ResultObject
}

func (s *ListVariableDefineResponseBody) GetTotalItem() *int32 {
	return s.TotalItem
}

func (s *ListVariableDefineResponseBody) GetTotalPage() *int32 {
	return s.TotalPage
}

func (s *ListVariableDefineResponseBody) SetRequestId(v string) *ListVariableDefineResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListVariableDefineResponseBody) SetCurrentPage(v int32) *ListVariableDefineResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *ListVariableDefineResponseBody) SetPageSize(v int32) *ListVariableDefineResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListVariableDefineResponseBody) SetResultObject(v bool) *ListVariableDefineResponseBody {
	s.ResultObject = &v
	return s
}

func (s *ListVariableDefineResponseBody) SetTotalItem(v int32) *ListVariableDefineResponseBody {
	s.TotalItem = &v
	return s
}

func (s *ListVariableDefineResponseBody) SetTotalPage(v int32) *ListVariableDefineResponseBody {
	s.TotalPage = &v
	return s
}

func (s *ListVariableDefineResponseBody) Validate() error {
	return dara.Validate(s)
}
