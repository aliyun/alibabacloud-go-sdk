// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFieldPageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeFieldPageResponseBody
	GetRequestId() *string
	SetCurrentPage(v int32) *DescribeFieldPageResponseBody
	GetCurrentPage() *int32
	SetPageSize(v int32) *DescribeFieldPageResponseBody
	GetPageSize() *int32
	SetResultObject(v []*DescribeFieldPageResponseBodyResultObject) *DescribeFieldPageResponseBody
	GetResultObject() []*DescribeFieldPageResponseBodyResultObject
	SetTotalItem(v int32) *DescribeFieldPageResponseBody
	GetTotalItem() *int32
	SetTotalPage(v int32) *DescribeFieldPageResponseBody
	GetTotalPage() *int32
}

type DescribeFieldPageResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// A32FE941-35F2-5378-B37C-4B8FDB16F094
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Current page number
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"currentPage,omitempty" xml:"currentPage,omitempty"`
	// Number of items per page, default value is 10
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// Returned object
	ResultObject []*DescribeFieldPageResponseBodyResultObject `json:"resultObject,omitempty" xml:"resultObject,omitempty" type:"Repeated"`
	// Total number of items
	//
	// example:
	//
	// 40
	TotalItem *int32 `json:"totalItem,omitempty" xml:"totalItem,omitempty"`
	// Total number of pages.
	//
	// example:
	//
	// 4
	TotalPage *int32 `json:"totalPage,omitempty" xml:"totalPage,omitempty"`
}

func (s DescribeFieldPageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeFieldPageResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFieldPageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeFieldPageResponseBody) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeFieldPageResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeFieldPageResponseBody) GetResultObject() []*DescribeFieldPageResponseBodyResultObject {
	return s.ResultObject
}

func (s *DescribeFieldPageResponseBody) GetTotalItem() *int32 {
	return s.TotalItem
}

func (s *DescribeFieldPageResponseBody) GetTotalPage() *int32 {
	return s.TotalPage
}

func (s *DescribeFieldPageResponseBody) SetRequestId(v string) *DescribeFieldPageResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFieldPageResponseBody) SetCurrentPage(v int32) *DescribeFieldPageResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *DescribeFieldPageResponseBody) SetPageSize(v int32) *DescribeFieldPageResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeFieldPageResponseBody) SetResultObject(v []*DescribeFieldPageResponseBodyResultObject) *DescribeFieldPageResponseBody {
	s.ResultObject = v
	return s
}

func (s *DescribeFieldPageResponseBody) SetTotalItem(v int32) *DescribeFieldPageResponseBody {
	s.TotalItem = &v
	return s
}

func (s *DescribeFieldPageResponseBody) SetTotalPage(v int32) *DescribeFieldPageResponseBody {
	s.TotalPage = &v
	return s
}

func (s *DescribeFieldPageResponseBody) Validate() error {
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

type DescribeFieldPageResponseBodyResultObject struct {
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
	// Unique table ID.
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
	// File source.
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

func (s DescribeFieldPageResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s DescribeFieldPageResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *DescribeFieldPageResponseBodyResultObject) GetClassify() *string {
	return s.Classify
}

func (s *DescribeFieldPageResponseBodyResultObject) GetDescription() *string {
	return s.Description
}

func (s *DescribeFieldPageResponseBodyResultObject) GetEnumData() *string {
	return s.EnumData
}

func (s *DescribeFieldPageResponseBodyResultObject) GetId() *int64 {
	return s.Id
}

func (s *DescribeFieldPageResponseBodyResultObject) GetName() *string {
	return s.Name
}

func (s *DescribeFieldPageResponseBodyResultObject) GetSource() *string {
	return s.Source
}

func (s *DescribeFieldPageResponseBodyResultObject) GetStatus() *string {
	return s.Status
}

func (s *DescribeFieldPageResponseBodyResultObject) GetTitle() *string {
	return s.Title
}

func (s *DescribeFieldPageResponseBodyResultObject) GetType() *string {
	return s.Type
}

func (s *DescribeFieldPageResponseBodyResultObject) SetClassify(v string) *DescribeFieldPageResponseBodyResultObject {
	s.Classify = &v
	return s
}

func (s *DescribeFieldPageResponseBodyResultObject) SetDescription(v string) *DescribeFieldPageResponseBodyResultObject {
	s.Description = &v
	return s
}

func (s *DescribeFieldPageResponseBodyResultObject) SetEnumData(v string) *DescribeFieldPageResponseBodyResultObject {
	s.EnumData = &v
	return s
}

func (s *DescribeFieldPageResponseBodyResultObject) SetId(v int64) *DescribeFieldPageResponseBodyResultObject {
	s.Id = &v
	return s
}

func (s *DescribeFieldPageResponseBodyResultObject) SetName(v string) *DescribeFieldPageResponseBodyResultObject {
	s.Name = &v
	return s
}

func (s *DescribeFieldPageResponseBodyResultObject) SetSource(v string) *DescribeFieldPageResponseBodyResultObject {
	s.Source = &v
	return s
}

func (s *DescribeFieldPageResponseBodyResultObject) SetStatus(v string) *DescribeFieldPageResponseBodyResultObject {
	s.Status = &v
	return s
}

func (s *DescribeFieldPageResponseBodyResultObject) SetTitle(v string) *DescribeFieldPageResponseBodyResultObject {
	s.Title = &v
	return s
}

func (s *DescribeFieldPageResponseBodyResultObject) SetType(v string) *DescribeFieldPageResponseBodyResultObject {
	s.Type = &v
	return s
}

func (s *DescribeFieldPageResponseBodyResultObject) Validate() error {
	return dara.Validate(s)
}
