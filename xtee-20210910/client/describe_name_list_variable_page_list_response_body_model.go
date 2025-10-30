// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNameListVariablePageListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *DescribeNameListVariablePageListResponseBody
	GetCurrentPage() *int32
	SetPageSize(v int32) *DescribeNameListVariablePageListResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeNameListVariablePageListResponseBody
	GetRequestId() *string
	SetResultObject(v []*DescribeNameListVariablePageListResponseBodyResultObject) *DescribeNameListVariablePageListResponseBody
	GetResultObject() []*DescribeNameListVariablePageListResponseBodyResultObject
	SetTotalItem(v int32) *DescribeNameListVariablePageListResponseBody
	GetTotalItem() *int32
	SetTotalPage(v int32) *DescribeNameListVariablePageListResponseBody
	GetTotalPage() *int32
}

type DescribeNameListVariablePageListResponseBody struct {
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
	// Request ID
	//
	// example:
	//
	// AE7E6105-7DEB-5125-9B24-DCBC139F6CD2
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Return object
	ResultObject []*DescribeNameListVariablePageListResponseBodyResultObject `json:"resultObject,omitempty" xml:"resultObject,omitempty" type:"Repeated"`
	// Total number of items
	//
	// example:
	//
	// 101
	TotalItem *int32 `json:"totalItem,omitempty" xml:"totalItem,omitempty"`
	// Total number of pages
	//
	// example:
	//
	// 1
	TotalPage *int32 `json:"totalPage,omitempty" xml:"totalPage,omitempty"`
}

func (s DescribeNameListVariablePageListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeNameListVariablePageListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeNameListVariablePageListResponseBody) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeNameListVariablePageListResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeNameListVariablePageListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeNameListVariablePageListResponseBody) GetResultObject() []*DescribeNameListVariablePageListResponseBodyResultObject {
	return s.ResultObject
}

func (s *DescribeNameListVariablePageListResponseBody) GetTotalItem() *int32 {
	return s.TotalItem
}

func (s *DescribeNameListVariablePageListResponseBody) GetTotalPage() *int32 {
	return s.TotalPage
}

func (s *DescribeNameListVariablePageListResponseBody) SetCurrentPage(v int32) *DescribeNameListVariablePageListResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *DescribeNameListVariablePageListResponseBody) SetPageSize(v int32) *DescribeNameListVariablePageListResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeNameListVariablePageListResponseBody) SetRequestId(v string) *DescribeNameListVariablePageListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeNameListVariablePageListResponseBody) SetResultObject(v []*DescribeNameListVariablePageListResponseBodyResultObject) *DescribeNameListVariablePageListResponseBody {
	s.ResultObject = v
	return s
}

func (s *DescribeNameListVariablePageListResponseBody) SetTotalItem(v int32) *DescribeNameListVariablePageListResponseBody {
	s.TotalItem = &v
	return s
}

func (s *DescribeNameListVariablePageListResponseBody) SetTotalPage(v int32) *DescribeNameListVariablePageListResponseBody {
	s.TotalPage = &v
	return s
}

func (s *DescribeNameListVariablePageListResponseBody) Validate() error {
	if s.ResultObject != nil {
		for _, item := range s.ResultObject {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeNameListVariablePageListResponseBodyResultObject struct {
	// Used capacity
	//
	// example:
	//
	// 3
	Capacity *int64 `json:"capacity,omitempty" xml:"capacity,omitempty"`
	// Description information.
	//
	// example:
	//
	// 变量描述
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// Creation time.
	//
	// example:
	//
	// 1621578648000
	GmtCreate *int64 `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// Modification time
	//
	// example:
	//
	// 1565701886000
	GmtModified *int64 `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// Name list ID
	//
	// example:
	//
	// 497
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// Remaining capacity
	//
	// example:
	//
	// 99997
	LeftCapacity *int64 `json:"leftCapacity,omitempty" xml:"leftCapacity,omitempty"`
	// Parameter name.
	//
	// example:
	//
	// nl_UN8otElLb490
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// Name list type
	//
	// example:
	//
	// 手机号
	NameListType *string `json:"nameListType,omitempty" xml:"nameListType,omitempty"`
	// Associated event eventCode
	//
	// example:
	//
	// -1
	RefObjId *string `json:"refObjId,omitempty" xml:"refObjId,omitempty"`
	// Association type
	//
	// example:
	//
	// EVENT
	RefObjType *int64 `json:"refObjType,omitempty" xml:"refObjType,omitempty"`
	// Data source
	//
	// example:
	//
	// SAF
	SourceType *string `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
	// Title.
	//
	// example:
	//
	// 白名单
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// Variable type
	//
	// example:
	//
	// NAME_LIST
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// User UID
	//
	// example:
	//
	// 180650758xxxxxxx
	UserId *int64 `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s DescribeNameListVariablePageListResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s DescribeNameListVariablePageListResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *DescribeNameListVariablePageListResponseBodyResultObject) GetCapacity() *int64 {
	return s.Capacity
}

func (s *DescribeNameListVariablePageListResponseBodyResultObject) GetDescription() *string {
	return s.Description
}

func (s *DescribeNameListVariablePageListResponseBodyResultObject) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *DescribeNameListVariablePageListResponseBodyResultObject) GetGmtModified() *int64 {
	return s.GmtModified
}

func (s *DescribeNameListVariablePageListResponseBodyResultObject) GetId() *int64 {
	return s.Id
}

func (s *DescribeNameListVariablePageListResponseBodyResultObject) GetLeftCapacity() *int64 {
	return s.LeftCapacity
}

func (s *DescribeNameListVariablePageListResponseBodyResultObject) GetName() *string {
	return s.Name
}

func (s *DescribeNameListVariablePageListResponseBodyResultObject) GetNameListType() *string {
	return s.NameListType
}

func (s *DescribeNameListVariablePageListResponseBodyResultObject) GetRefObjId() *string {
	return s.RefObjId
}

func (s *DescribeNameListVariablePageListResponseBodyResultObject) GetRefObjType() *int64 {
	return s.RefObjType
}

func (s *DescribeNameListVariablePageListResponseBodyResultObject) GetSourceType() *string {
	return s.SourceType
}

func (s *DescribeNameListVariablePageListResponseBodyResultObject) GetTitle() *string {
	return s.Title
}

func (s *DescribeNameListVariablePageListResponseBodyResultObject) GetType() *string {
	return s.Type
}

func (s *DescribeNameListVariablePageListResponseBodyResultObject) GetUserId() *int64 {
	return s.UserId
}

func (s *DescribeNameListVariablePageListResponseBodyResultObject) SetCapacity(v int64) *DescribeNameListVariablePageListResponseBodyResultObject {
	s.Capacity = &v
	return s
}

func (s *DescribeNameListVariablePageListResponseBodyResultObject) SetDescription(v string) *DescribeNameListVariablePageListResponseBodyResultObject {
	s.Description = &v
	return s
}

func (s *DescribeNameListVariablePageListResponseBodyResultObject) SetGmtCreate(v int64) *DescribeNameListVariablePageListResponseBodyResultObject {
	s.GmtCreate = &v
	return s
}

func (s *DescribeNameListVariablePageListResponseBodyResultObject) SetGmtModified(v int64) *DescribeNameListVariablePageListResponseBodyResultObject {
	s.GmtModified = &v
	return s
}

func (s *DescribeNameListVariablePageListResponseBodyResultObject) SetId(v int64) *DescribeNameListVariablePageListResponseBodyResultObject {
	s.Id = &v
	return s
}

func (s *DescribeNameListVariablePageListResponseBodyResultObject) SetLeftCapacity(v int64) *DescribeNameListVariablePageListResponseBodyResultObject {
	s.LeftCapacity = &v
	return s
}

func (s *DescribeNameListVariablePageListResponseBodyResultObject) SetName(v string) *DescribeNameListVariablePageListResponseBodyResultObject {
	s.Name = &v
	return s
}

func (s *DescribeNameListVariablePageListResponseBodyResultObject) SetNameListType(v string) *DescribeNameListVariablePageListResponseBodyResultObject {
	s.NameListType = &v
	return s
}

func (s *DescribeNameListVariablePageListResponseBodyResultObject) SetRefObjId(v string) *DescribeNameListVariablePageListResponseBodyResultObject {
	s.RefObjId = &v
	return s
}

func (s *DescribeNameListVariablePageListResponseBodyResultObject) SetRefObjType(v int64) *DescribeNameListVariablePageListResponseBodyResultObject {
	s.RefObjType = &v
	return s
}

func (s *DescribeNameListVariablePageListResponseBodyResultObject) SetSourceType(v string) *DescribeNameListVariablePageListResponseBodyResultObject {
	s.SourceType = &v
	return s
}

func (s *DescribeNameListVariablePageListResponseBodyResultObject) SetTitle(v string) *DescribeNameListVariablePageListResponseBodyResultObject {
	s.Title = &v
	return s
}

func (s *DescribeNameListVariablePageListResponseBodyResultObject) SetType(v string) *DescribeNameListVariablePageListResponseBodyResultObject {
	s.Type = &v
	return s
}

func (s *DescribeNameListVariablePageListResponseBodyResultObject) SetUserId(v int64) *DescribeNameListVariablePageListResponseBodyResultObject {
	s.UserId = &v
	return s
}

func (s *DescribeNameListVariablePageListResponseBodyResultObject) Validate() error {
	return dara.Validate(s)
}
