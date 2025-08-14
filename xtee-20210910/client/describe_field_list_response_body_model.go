// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFieldListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeFieldListResponseBody
	GetRequestId() *string
	SetCurrentPage(v int32) *DescribeFieldListResponseBody
	GetCurrentPage() *int32
	SetPageSize(v int32) *DescribeFieldListResponseBody
	GetPageSize() *int32
	SetResultObject(v []*DescribeFieldListResponseBodyResultObject) *DescribeFieldListResponseBody
	GetResultObject() []*DescribeFieldListResponseBodyResultObject
	SetTotalItem(v int32) *DescribeFieldListResponseBody
	GetTotalItem() *int32
	SetTotalPage(v int32) *DescribeFieldListResponseBody
	GetTotalPage() *int32
}

type DescribeFieldListResponseBody struct {
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
	// Page size, default value is 10
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// Returned object
	ResultObject []*DescribeFieldListResponseBodyResultObject `json:"resultObject,omitempty" xml:"resultObject,omitempty" type:"Repeated"`
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
	// 9
	TotalPage *int32 `json:"totalPage,omitempty" xml:"totalPage,omitempty"`
}

func (s DescribeFieldListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeFieldListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFieldListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeFieldListResponseBody) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeFieldListResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeFieldListResponseBody) GetResultObject() []*DescribeFieldListResponseBodyResultObject {
	return s.ResultObject
}

func (s *DescribeFieldListResponseBody) GetTotalItem() *int32 {
	return s.TotalItem
}

func (s *DescribeFieldListResponseBody) GetTotalPage() *int32 {
	return s.TotalPage
}

func (s *DescribeFieldListResponseBody) SetRequestId(v string) *DescribeFieldListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFieldListResponseBody) SetCurrentPage(v int32) *DescribeFieldListResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *DescribeFieldListResponseBody) SetPageSize(v int32) *DescribeFieldListResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeFieldListResponseBody) SetResultObject(v []*DescribeFieldListResponseBodyResultObject) *DescribeFieldListResponseBody {
	s.ResultObject = v
	return s
}

func (s *DescribeFieldListResponseBody) SetTotalItem(v int32) *DescribeFieldListResponseBody {
	s.TotalItem = &v
	return s
}

func (s *DescribeFieldListResponseBody) SetTotalPage(v int32) *DescribeFieldListResponseBody {
	s.TotalPage = &v
	return s
}

func (s *DescribeFieldListResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeFieldListResponseBodyResultObject struct {
	// Field classification
	//
	// example:
	//
	// REQUEST_PARAM
	Classify *string `json:"classify,omitempty" xml:"classify,omitempty"`
	// Description information.
	//
	// example:
	//
	// 描述
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// Enum data
	//
	// example:
	//
	// STATUS
	EnumData *string `json:"enumData,omitempty" xml:"enumData,omitempty"`
	// Field ID.
	//
	// example:
	//
	// 497
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// Field name
	//
	// example:
	//
	// age
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// Field source
	//
	// example:
	//
	// DEFINE
	Source *string `json:"source,omitempty" xml:"source,omitempty"`
	// Status.
	//
	// example:
	//
	// ENABLE
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// Title.
	//
	// example:
	//
	// 年龄
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// Field type
	//
	// example:
	//
	// STRING
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s DescribeFieldListResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s DescribeFieldListResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *DescribeFieldListResponseBodyResultObject) GetClassify() *string {
	return s.Classify
}

func (s *DescribeFieldListResponseBodyResultObject) GetDescription() *string {
	return s.Description
}

func (s *DescribeFieldListResponseBodyResultObject) GetEnumData() *string {
	return s.EnumData
}

func (s *DescribeFieldListResponseBodyResultObject) GetId() *int64 {
	return s.Id
}

func (s *DescribeFieldListResponseBodyResultObject) GetName() *string {
	return s.Name
}

func (s *DescribeFieldListResponseBodyResultObject) GetSource() *string {
	return s.Source
}

func (s *DescribeFieldListResponseBodyResultObject) GetStatus() *string {
	return s.Status
}

func (s *DescribeFieldListResponseBodyResultObject) GetTitle() *string {
	return s.Title
}

func (s *DescribeFieldListResponseBodyResultObject) GetType() *string {
	return s.Type
}

func (s *DescribeFieldListResponseBodyResultObject) SetClassify(v string) *DescribeFieldListResponseBodyResultObject {
	s.Classify = &v
	return s
}

func (s *DescribeFieldListResponseBodyResultObject) SetDescription(v string) *DescribeFieldListResponseBodyResultObject {
	s.Description = &v
	return s
}

func (s *DescribeFieldListResponseBodyResultObject) SetEnumData(v string) *DescribeFieldListResponseBodyResultObject {
	s.EnumData = &v
	return s
}

func (s *DescribeFieldListResponseBodyResultObject) SetId(v int64) *DescribeFieldListResponseBodyResultObject {
	s.Id = &v
	return s
}

func (s *DescribeFieldListResponseBodyResultObject) SetName(v string) *DescribeFieldListResponseBodyResultObject {
	s.Name = &v
	return s
}

func (s *DescribeFieldListResponseBodyResultObject) SetSource(v string) *DescribeFieldListResponseBodyResultObject {
	s.Source = &v
	return s
}

func (s *DescribeFieldListResponseBodyResultObject) SetStatus(v string) *DescribeFieldListResponseBodyResultObject {
	s.Status = &v
	return s
}

func (s *DescribeFieldListResponseBodyResultObject) SetTitle(v string) *DescribeFieldListResponseBodyResultObject {
	s.Title = &v
	return s
}

func (s *DescribeFieldListResponseBodyResultObject) SetType(v string) *DescribeFieldListResponseBodyResultObject {
	s.Type = &v
	return s
}

func (s *DescribeFieldListResponseBodyResultObject) Validate() error {
	return dara.Validate(s)
}
