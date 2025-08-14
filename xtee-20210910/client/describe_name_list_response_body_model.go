// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNameListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeNameListResponseBody
	GetRequestId() *string
	SetCurrentPage(v int32) *DescribeNameListResponseBody
	GetCurrentPage() *int32
	SetPageSize(v int32) *DescribeNameListResponseBody
	GetPageSize() *int32
	SetResultObject(v []*DescribeNameListResponseBodyResultObject) *DescribeNameListResponseBody
	GetResultObject() []*DescribeNameListResponseBodyResultObject
	SetTotalItem(v int32) *DescribeNameListResponseBody
	GetTotalItem() *int32
	SetTotalPage(v int32) *DescribeNameListResponseBody
	GetTotalPage() *int32
}

type DescribeNameListResponseBody struct {
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
	// Returned object.
	ResultObject []*DescribeNameListResponseBodyResultObject `json:"resultObject,omitempty" xml:"resultObject,omitempty" type:"Repeated"`
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

func (s DescribeNameListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeNameListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeNameListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeNameListResponseBody) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeNameListResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeNameListResponseBody) GetResultObject() []*DescribeNameListResponseBodyResultObject {
	return s.ResultObject
}

func (s *DescribeNameListResponseBody) GetTotalItem() *int32 {
	return s.TotalItem
}

func (s *DescribeNameListResponseBody) GetTotalPage() *int32 {
	return s.TotalPage
}

func (s *DescribeNameListResponseBody) SetRequestId(v string) *DescribeNameListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeNameListResponseBody) SetCurrentPage(v int32) *DescribeNameListResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *DescribeNameListResponseBody) SetPageSize(v int32) *DescribeNameListResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeNameListResponseBody) SetResultObject(v []*DescribeNameListResponseBodyResultObject) *DescribeNameListResponseBody {
	s.ResultObject = v
	return s
}

func (s *DescribeNameListResponseBody) SetTotalItem(v int32) *DescribeNameListResponseBody {
	s.TotalItem = &v
	return s
}

func (s *DescribeNameListResponseBody) SetTotalPage(v int32) *DescribeNameListResponseBody {
	s.TotalPage = &v
	return s
}

func (s *DescribeNameListResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeNameListResponseBodyResultObject struct {
	// Creation time.
	//
	// example:
	//
	// 1621578648000
	GmtCreate *int64 `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// Modification time.
	//
	// example:
	//
	// 1565701886000
	GmtModified *int64 `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// Name list content ID.
	//
	// example:
	//
	// 30
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// Variable identifier.
	//
	// example:
	//
	// NAME_LIST
	Identifier *string `json:"identifier,omitempty" xml:"identifier,omitempty"`
	// Title.
	//
	// example:
	//
	// 变量的title
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// Value.
	//
	// example:
	//
	// 321311193502064288
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
	// Variable ID.
	//
	// example:
	//
	// 393314
	VariableId *int64 `json:"variableId,omitempty" xml:"variableId,omitempty"`
}

func (s DescribeNameListResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s DescribeNameListResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *DescribeNameListResponseBodyResultObject) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *DescribeNameListResponseBodyResultObject) GetGmtModified() *int64 {
	return s.GmtModified
}

func (s *DescribeNameListResponseBodyResultObject) GetId() *int64 {
	return s.Id
}

func (s *DescribeNameListResponseBodyResultObject) GetIdentifier() *string {
	return s.Identifier
}

func (s *DescribeNameListResponseBodyResultObject) GetTitle() *string {
	return s.Title
}

func (s *DescribeNameListResponseBodyResultObject) GetValue() *string {
	return s.Value
}

func (s *DescribeNameListResponseBodyResultObject) GetVariableId() *int64 {
	return s.VariableId
}

func (s *DescribeNameListResponseBodyResultObject) SetGmtCreate(v int64) *DescribeNameListResponseBodyResultObject {
	s.GmtCreate = &v
	return s
}

func (s *DescribeNameListResponseBodyResultObject) SetGmtModified(v int64) *DescribeNameListResponseBodyResultObject {
	s.GmtModified = &v
	return s
}

func (s *DescribeNameListResponseBodyResultObject) SetId(v int64) *DescribeNameListResponseBodyResultObject {
	s.Id = &v
	return s
}

func (s *DescribeNameListResponseBodyResultObject) SetIdentifier(v string) *DescribeNameListResponseBodyResultObject {
	s.Identifier = &v
	return s
}

func (s *DescribeNameListResponseBodyResultObject) SetTitle(v string) *DescribeNameListResponseBodyResultObject {
	s.Title = &v
	return s
}

func (s *DescribeNameListResponseBodyResultObject) SetValue(v string) *DescribeNameListResponseBodyResultObject {
	s.Value = &v
	return s
}

func (s *DescribeNameListResponseBodyResultObject) SetVariableId(v int64) *DescribeNameListResponseBodyResultObject {
	s.VariableId = &v
	return s
}

func (s *DescribeNameListResponseBodyResultObject) Validate() error {
	return dara.Validate(s)
}
