// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVariableListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeVariableListResponseBody
	GetRequestId() *string
	SetCurrentPage(v int32) *DescribeVariableListResponseBody
	GetCurrentPage() *int32
	SetPageSize(v int32) *DescribeVariableListResponseBody
	GetPageSize() *int32
	SetResultObject(v []*DescribeVariableListResponseBodyResultObject) *DescribeVariableListResponseBody
	GetResultObject() []*DescribeVariableListResponseBodyResultObject
	SetTotalItem(v int32) *DescribeVariableListResponseBody
	GetTotalItem() *int32
	SetTotalPage(v int32) *DescribeVariableListResponseBody
	GetTotalPage() *int32
}

type DescribeVariableListResponseBody struct {
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
	// Returned object.
	ResultObject []*DescribeVariableListResponseBodyResultObject `json:"resultObject,omitempty" xml:"resultObject,omitempty" type:"Repeated"`
	// Total number of items.
	//
	// example:
	//
	// 27
	TotalItem *int32 `json:"totalItem,omitempty" xml:"totalItem,omitempty"`
	// Total number of pages.
	//
	// example:
	//
	// 3
	TotalPage *int32 `json:"totalPage,omitempty" xml:"totalPage,omitempty"`
}

func (s DescribeVariableListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVariableListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVariableListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVariableListResponseBody) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeVariableListResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeVariableListResponseBody) GetResultObject() []*DescribeVariableListResponseBodyResultObject {
	return s.ResultObject
}

func (s *DescribeVariableListResponseBody) GetTotalItem() *int32 {
	return s.TotalItem
}

func (s *DescribeVariableListResponseBody) GetTotalPage() *int32 {
	return s.TotalPage
}

func (s *DescribeVariableListResponseBody) SetRequestId(v string) *DescribeVariableListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVariableListResponseBody) SetCurrentPage(v int32) *DescribeVariableListResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *DescribeVariableListResponseBody) SetPageSize(v int32) *DescribeVariableListResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeVariableListResponseBody) SetResultObject(v []*DescribeVariableListResponseBodyResultObject) *DescribeVariableListResponseBody {
	s.ResultObject = v
	return s
}

func (s *DescribeVariableListResponseBody) SetTotalItem(v int32) *DescribeVariableListResponseBody {
	s.TotalItem = &v
	return s
}

func (s *DescribeVariableListResponseBody) SetTotalPage(v int32) *DescribeVariableListResponseBody {
	s.TotalPage = &v
	return s
}

func (s *DescribeVariableListResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeVariableListResponseBodyResultObject struct {
	// Capacity.
	//
	// example:
	//
	// 100000
	Capacity *int64 `json:"capacity,omitempty" xml:"capacity,omitempty"`
	// Variable definition ID.
	//
	// example:
	//
	// 10
	DefineId *string `json:"defineId,omitempty" xml:"defineId,omitempty"`
	// Description information.
	//
	// example:
	//
	// 变量描述
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// Extended information.
	//
	// example:
	//
	// 暂无
	ExtendInfo map[string]interface{} `json:"extendInfo,omitempty" xml:"extendInfo,omitempty"`
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
	// Variable ID.
	//
	// example:
	//
	// 497
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// Remaining capacity.
	//
	// example:
	//
	// 100000
	LeftCapacity *int64 `json:"leftCapacity,omitempty" xml:"leftCapacity,omitempty"`
	// Variable name
	//
	// example:
	//
	// __ipLocationCityCode__
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// Variable output type
	//
	// example:
	//
	// STRING
	OutputsType *string `json:"outputsType,omitempty" xml:"outputsType,omitempty"`
	// Associated event ID.
	//
	// example:
	//
	// de_agdxgz0246
	RefObjId *string `json:"refObjId,omitempty" xml:"refObjId,omitempty"`
	// Associated event name.
	//
	// example:
	//
	// 注册事件
	RefObjName *string `json:"refObjName,omitempty" xml:"refObjName,omitempty"`
	// Associated object type of the variable
	//
	// example:
	//
	// EVENT
	RefObjType *string `json:"refObjType,omitempty" xml:"refObjType,omitempty"`
	// Source type.
	//
	// example:
	//
	// SAF
	SourceType *string `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
	// Title.
	//
	// example:
	//
	// 设备风险识别_标签
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// Variable type.
	//
	// example:
	//
	// NATIVE
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// User ID to which the data belongs.
	//
	// example:
	//
	// 100000
	UserId *int64 `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s DescribeVariableListResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s DescribeVariableListResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *DescribeVariableListResponseBodyResultObject) GetCapacity() *int64 {
	return s.Capacity
}

func (s *DescribeVariableListResponseBodyResultObject) GetDefineId() *string {
	return s.DefineId
}

func (s *DescribeVariableListResponseBodyResultObject) GetDescription() *string {
	return s.Description
}

func (s *DescribeVariableListResponseBodyResultObject) GetExtendInfo() map[string]interface{} {
	return s.ExtendInfo
}

func (s *DescribeVariableListResponseBodyResultObject) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *DescribeVariableListResponseBodyResultObject) GetGmtModified() *int64 {
	return s.GmtModified
}

func (s *DescribeVariableListResponseBodyResultObject) GetId() *int64 {
	return s.Id
}

func (s *DescribeVariableListResponseBodyResultObject) GetLeftCapacity() *int64 {
	return s.LeftCapacity
}

func (s *DescribeVariableListResponseBodyResultObject) GetName() *string {
	return s.Name
}

func (s *DescribeVariableListResponseBodyResultObject) GetOutputsType() *string {
	return s.OutputsType
}

func (s *DescribeVariableListResponseBodyResultObject) GetRefObjId() *string {
	return s.RefObjId
}

func (s *DescribeVariableListResponseBodyResultObject) GetRefObjName() *string {
	return s.RefObjName
}

func (s *DescribeVariableListResponseBodyResultObject) GetRefObjType() *string {
	return s.RefObjType
}

func (s *DescribeVariableListResponseBodyResultObject) GetSourceType() *string {
	return s.SourceType
}

func (s *DescribeVariableListResponseBodyResultObject) GetTitle() *string {
	return s.Title
}

func (s *DescribeVariableListResponseBodyResultObject) GetType() *string {
	return s.Type
}

func (s *DescribeVariableListResponseBodyResultObject) GetUserId() *int64 {
	return s.UserId
}

func (s *DescribeVariableListResponseBodyResultObject) SetCapacity(v int64) *DescribeVariableListResponseBodyResultObject {
	s.Capacity = &v
	return s
}

func (s *DescribeVariableListResponseBodyResultObject) SetDefineId(v string) *DescribeVariableListResponseBodyResultObject {
	s.DefineId = &v
	return s
}

func (s *DescribeVariableListResponseBodyResultObject) SetDescription(v string) *DescribeVariableListResponseBodyResultObject {
	s.Description = &v
	return s
}

func (s *DescribeVariableListResponseBodyResultObject) SetExtendInfo(v map[string]interface{}) *DescribeVariableListResponseBodyResultObject {
	s.ExtendInfo = v
	return s
}

func (s *DescribeVariableListResponseBodyResultObject) SetGmtCreate(v int64) *DescribeVariableListResponseBodyResultObject {
	s.GmtCreate = &v
	return s
}

func (s *DescribeVariableListResponseBodyResultObject) SetGmtModified(v int64) *DescribeVariableListResponseBodyResultObject {
	s.GmtModified = &v
	return s
}

func (s *DescribeVariableListResponseBodyResultObject) SetId(v int64) *DescribeVariableListResponseBodyResultObject {
	s.Id = &v
	return s
}

func (s *DescribeVariableListResponseBodyResultObject) SetLeftCapacity(v int64) *DescribeVariableListResponseBodyResultObject {
	s.LeftCapacity = &v
	return s
}

func (s *DescribeVariableListResponseBodyResultObject) SetName(v string) *DescribeVariableListResponseBodyResultObject {
	s.Name = &v
	return s
}

func (s *DescribeVariableListResponseBodyResultObject) SetOutputsType(v string) *DescribeVariableListResponseBodyResultObject {
	s.OutputsType = &v
	return s
}

func (s *DescribeVariableListResponseBodyResultObject) SetRefObjId(v string) *DescribeVariableListResponseBodyResultObject {
	s.RefObjId = &v
	return s
}

func (s *DescribeVariableListResponseBodyResultObject) SetRefObjName(v string) *DescribeVariableListResponseBodyResultObject {
	s.RefObjName = &v
	return s
}

func (s *DescribeVariableListResponseBodyResultObject) SetRefObjType(v string) *DescribeVariableListResponseBodyResultObject {
	s.RefObjType = &v
	return s
}

func (s *DescribeVariableListResponseBodyResultObject) SetSourceType(v string) *DescribeVariableListResponseBodyResultObject {
	s.SourceType = &v
	return s
}

func (s *DescribeVariableListResponseBodyResultObject) SetTitle(v string) *DescribeVariableListResponseBodyResultObject {
	s.Title = &v
	return s
}

func (s *DescribeVariableListResponseBodyResultObject) SetType(v string) *DescribeVariableListResponseBodyResultObject {
	s.Type = &v
	return s
}

func (s *DescribeVariableListResponseBodyResultObject) SetUserId(v int64) *DescribeVariableListResponseBodyResultObject {
	s.UserId = &v
	return s
}

func (s *DescribeVariableListResponseBodyResultObject) Validate() error {
	return dara.Validate(s)
}
