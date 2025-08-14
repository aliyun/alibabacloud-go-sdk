// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNameListPageListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v string) *DescribeNameListPageListResponseBody
	GetCurrentPage() *string
	SetPageSize(v string) *DescribeNameListPageListResponseBody
	GetPageSize() *string
	SetRequestId(v string) *DescribeNameListPageListResponseBody
	GetRequestId() *string
	SetResultObject(v []*DescribeNameListPageListResponseBodyResultObject) *DescribeNameListPageListResponseBody
	GetResultObject() []*DescribeNameListPageListResponseBodyResultObject
	SetTotalItem(v string) *DescribeNameListPageListResponseBody
	GetTotalItem() *string
	SetTotalPage(v string) *DescribeNameListPageListResponseBody
	GetTotalPage() *string
}

type DescribeNameListPageListResponseBody struct {
	// Current page number.
	//
	// example:
	//
	// 1
	CurrentPage *string `json:"currentPage,omitempty" xml:"currentPage,omitempty"`
	// Page size, default value is 10
	//
	// example:
	//
	// 10
	PageSize *string `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// Request ID.
	//
	// example:
	//
	// AE7E6105-7DEB-5125-9B24-DCBC139F6CD2
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Returned object
	ResultObject []*DescribeNameListPageListResponseBodyResultObject `json:"resultObject,omitempty" xml:"resultObject,omitempty" type:"Repeated"`
	// Total number of items
	//
	// example:
	//
	// 101
	TotalItem *string `json:"totalItem,omitempty" xml:"totalItem,omitempty"`
	// Total number of pages
	//
	// example:
	//
	// 9
	TotalPage *string `json:"totalPage,omitempty" xml:"totalPage,omitempty"`
}

func (s DescribeNameListPageListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeNameListPageListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeNameListPageListResponseBody) GetCurrentPage() *string {
	return s.CurrentPage
}

func (s *DescribeNameListPageListResponseBody) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribeNameListPageListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeNameListPageListResponseBody) GetResultObject() []*DescribeNameListPageListResponseBodyResultObject {
	return s.ResultObject
}

func (s *DescribeNameListPageListResponseBody) GetTotalItem() *string {
	return s.TotalItem
}

func (s *DescribeNameListPageListResponseBody) GetTotalPage() *string {
	return s.TotalPage
}

func (s *DescribeNameListPageListResponseBody) SetCurrentPage(v string) *DescribeNameListPageListResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *DescribeNameListPageListResponseBody) SetPageSize(v string) *DescribeNameListPageListResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeNameListPageListResponseBody) SetRequestId(v string) *DescribeNameListPageListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeNameListPageListResponseBody) SetResultObject(v []*DescribeNameListPageListResponseBodyResultObject) *DescribeNameListPageListResponseBody {
	s.ResultObject = v
	return s
}

func (s *DescribeNameListPageListResponseBody) SetTotalItem(v string) *DescribeNameListPageListResponseBody {
	s.TotalItem = &v
	return s
}

func (s *DescribeNameListPageListResponseBody) SetTotalPage(v string) *DescribeNameListPageListResponseBody {
	s.TotalPage = &v
	return s
}

func (s *DescribeNameListPageListResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeNameListPageListResponseBodyResultObject struct {
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
	// ID of the list variable content data
	//
	// example:
	//
	// 497
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// NameList Content memo
	//
	// example:
	//
	// 名单内容描述
	Memo *string `json:"memo,omitempty" xml:"memo,omitempty"`
	// Variable name
	//
	// example:
	//
	// nl_UN8otElLb490
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// Variable type
	//
	// example:
	//
	// accountId
	NameListType *string `json:"nameListType,omitempty" xml:"nameListType,omitempty"`
	// Title.
	//
	// example:
	//
	// 白名单
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// User UID
	//
	// example:
	//
	// 130433202307074287
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// Variable value
	//
	// example:
	//
	// 130433202307074287
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
	// Variable ID.
	//
	// example:
	//
	// 762
	VariableId *int64 `json:"variableId,omitempty" xml:"variableId,omitempty"`
}

func (s DescribeNameListPageListResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s DescribeNameListPageListResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *DescribeNameListPageListResponseBodyResultObject) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *DescribeNameListPageListResponseBodyResultObject) GetGmtModified() *int64 {
	return s.GmtModified
}

func (s *DescribeNameListPageListResponseBodyResultObject) GetId() *int64 {
	return s.Id
}

func (s *DescribeNameListPageListResponseBodyResultObject) GetMemo() *string {
	return s.Memo
}

func (s *DescribeNameListPageListResponseBodyResultObject) GetName() *string {
	return s.Name
}

func (s *DescribeNameListPageListResponseBodyResultObject) GetNameListType() *string {
	return s.NameListType
}

func (s *DescribeNameListPageListResponseBodyResultObject) GetTitle() *string {
	return s.Title
}

func (s *DescribeNameListPageListResponseBodyResultObject) GetUserId() *string {
	return s.UserId
}

func (s *DescribeNameListPageListResponseBodyResultObject) GetValue() *string {
	return s.Value
}

func (s *DescribeNameListPageListResponseBodyResultObject) GetVariableId() *int64 {
	return s.VariableId
}

func (s *DescribeNameListPageListResponseBodyResultObject) SetGmtCreate(v int64) *DescribeNameListPageListResponseBodyResultObject {
	s.GmtCreate = &v
	return s
}

func (s *DescribeNameListPageListResponseBodyResultObject) SetGmtModified(v int64) *DescribeNameListPageListResponseBodyResultObject {
	s.GmtModified = &v
	return s
}

func (s *DescribeNameListPageListResponseBodyResultObject) SetId(v int64) *DescribeNameListPageListResponseBodyResultObject {
	s.Id = &v
	return s
}

func (s *DescribeNameListPageListResponseBodyResultObject) SetMemo(v string) *DescribeNameListPageListResponseBodyResultObject {
	s.Memo = &v
	return s
}

func (s *DescribeNameListPageListResponseBodyResultObject) SetName(v string) *DescribeNameListPageListResponseBodyResultObject {
	s.Name = &v
	return s
}

func (s *DescribeNameListPageListResponseBodyResultObject) SetNameListType(v string) *DescribeNameListPageListResponseBodyResultObject {
	s.NameListType = &v
	return s
}

func (s *DescribeNameListPageListResponseBodyResultObject) SetTitle(v string) *DescribeNameListPageListResponseBodyResultObject {
	s.Title = &v
	return s
}

func (s *DescribeNameListPageListResponseBodyResultObject) SetUserId(v string) *DescribeNameListPageListResponseBodyResultObject {
	s.UserId = &v
	return s
}

func (s *DescribeNameListPageListResponseBodyResultObject) SetValue(v string) *DescribeNameListPageListResponseBodyResultObject {
	s.Value = &v
	return s
}

func (s *DescribeNameListPageListResponseBodyResultObject) SetVariableId(v int64) *DescribeNameListPageListResponseBodyResultObject {
	s.VariableId = &v
	return s
}

func (s *DescribeNameListPageListResponseBodyResultObject) Validate() error {
	return dara.Validate(s)
}
