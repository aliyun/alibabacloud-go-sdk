// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDataSourcePageListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeDataSourcePageListResponseBody
	GetRequestId() *string
	SetCurrentPage(v int32) *DescribeDataSourcePageListResponseBody
	GetCurrentPage() *int32
	SetPageSize(v int32) *DescribeDataSourcePageListResponseBody
	GetPageSize() *int32
	SetResultObject(v []*DescribeDataSourcePageListResponseBodyResultObject) *DescribeDataSourcePageListResponseBody
	GetResultObject() []*DescribeDataSourcePageListResponseBodyResultObject
	SetTotalItem(v int32) *DescribeDataSourcePageListResponseBody
	GetTotalItem() *int32
	SetTotalPage(v int32) *DescribeDataSourcePageListResponseBody
	GetTotalPage() *int32
}

type DescribeDataSourcePageListResponseBody struct {
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
	// Return object
	ResultObject []*DescribeDataSourcePageListResponseBodyResultObject `json:"resultObject,omitempty" xml:"resultObject,omitempty" type:"Repeated"`
	// Total items
	//
	// example:
	//
	// 3
	TotalItem *int32 `json:"totalItem,omitempty" xml:"totalItem,omitempty"`
	// Total pages
	//
	// example:
	//
	// 9
	TotalPage *int32 `json:"totalPage,omitempty" xml:"totalPage,omitempty"`
}

func (s DescribeDataSourcePageListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataSourcePageListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDataSourcePageListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDataSourcePageListResponseBody) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeDataSourcePageListResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDataSourcePageListResponseBody) GetResultObject() []*DescribeDataSourcePageListResponseBodyResultObject {
	return s.ResultObject
}

func (s *DescribeDataSourcePageListResponseBody) GetTotalItem() *int32 {
	return s.TotalItem
}

func (s *DescribeDataSourcePageListResponseBody) GetTotalPage() *int32 {
	return s.TotalPage
}

func (s *DescribeDataSourcePageListResponseBody) SetRequestId(v string) *DescribeDataSourcePageListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDataSourcePageListResponseBody) SetCurrentPage(v int32) *DescribeDataSourcePageListResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *DescribeDataSourcePageListResponseBody) SetPageSize(v int32) *DescribeDataSourcePageListResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeDataSourcePageListResponseBody) SetResultObject(v []*DescribeDataSourcePageListResponseBodyResultObject) *DescribeDataSourcePageListResponseBody {
	s.ResultObject = v
	return s
}

func (s *DescribeDataSourcePageListResponseBody) SetTotalItem(v int32) *DescribeDataSourcePageListResponseBody {
	s.TotalItem = &v
	return s
}

func (s *DescribeDataSourcePageListResponseBody) SetTotalPage(v int32) *DescribeDataSourcePageListResponseBody {
	s.TotalPage = &v
	return s
}

func (s *DescribeDataSourcePageListResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDataSourcePageListResponseBodyResultObject struct {
	// Creator of the data source.
	//
	// example:
	//
	// xxxx
	Creator *string `json:"creator,omitempty" xml:"creator,omitempty"`
	// Data source description.
	//
	// example:
	//
	// 描述
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// Time when the data source was created.
	//
	// example:
	//
	// 1621578648000
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// Time when the data source was last modified.
	//
	// example:
	//
	// 1565701886000
	GmtModified *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// Data source ID.
	//
	// example:
	//
	// 497
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// Last modifier.
	//
	// example:
	//
	// xxxxx
	Modifier *string `json:"modifier,omitempty" xml:"modifier,omitempty"`
	// Data source name.
	//
	// example:
	//
	// data_source
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// Total pages.
	//
	// example:
	//
	// 4
	Total *int64 `json:"total,omitempty" xml:"total,omitempty"`
	// Data source type.
	//
	// example:
	//
	// FILE
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s DescribeDataSourcePageListResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataSourcePageListResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *DescribeDataSourcePageListResponseBodyResultObject) GetCreator() *string {
	return s.Creator
}

func (s *DescribeDataSourcePageListResponseBodyResultObject) GetDescription() *string {
	return s.Description
}

func (s *DescribeDataSourcePageListResponseBodyResultObject) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *DescribeDataSourcePageListResponseBodyResultObject) GetGmtModified() *string {
	return s.GmtModified
}

func (s *DescribeDataSourcePageListResponseBodyResultObject) GetId() *int64 {
	return s.Id
}

func (s *DescribeDataSourcePageListResponseBodyResultObject) GetModifier() *string {
	return s.Modifier
}

func (s *DescribeDataSourcePageListResponseBodyResultObject) GetName() *string {
	return s.Name
}

func (s *DescribeDataSourcePageListResponseBodyResultObject) GetTotal() *int64 {
	return s.Total
}

func (s *DescribeDataSourcePageListResponseBodyResultObject) GetType() *string {
	return s.Type
}

func (s *DescribeDataSourcePageListResponseBodyResultObject) SetCreator(v string) *DescribeDataSourcePageListResponseBodyResultObject {
	s.Creator = &v
	return s
}

func (s *DescribeDataSourcePageListResponseBodyResultObject) SetDescription(v string) *DescribeDataSourcePageListResponseBodyResultObject {
	s.Description = &v
	return s
}

func (s *DescribeDataSourcePageListResponseBodyResultObject) SetGmtCreate(v string) *DescribeDataSourcePageListResponseBodyResultObject {
	s.GmtCreate = &v
	return s
}

func (s *DescribeDataSourcePageListResponseBodyResultObject) SetGmtModified(v string) *DescribeDataSourcePageListResponseBodyResultObject {
	s.GmtModified = &v
	return s
}

func (s *DescribeDataSourcePageListResponseBodyResultObject) SetId(v int64) *DescribeDataSourcePageListResponseBodyResultObject {
	s.Id = &v
	return s
}

func (s *DescribeDataSourcePageListResponseBodyResultObject) SetModifier(v string) *DescribeDataSourcePageListResponseBodyResultObject {
	s.Modifier = &v
	return s
}

func (s *DescribeDataSourcePageListResponseBodyResultObject) SetName(v string) *DescribeDataSourcePageListResponseBodyResultObject {
	s.Name = &v
	return s
}

func (s *DescribeDataSourcePageListResponseBodyResultObject) SetTotal(v int64) *DescribeDataSourcePageListResponseBodyResultObject {
	s.Total = &v
	return s
}

func (s *DescribeDataSourcePageListResponseBodyResultObject) SetType(v string) *DescribeDataSourcePageListResponseBodyResultObject {
	s.Type = &v
	return s
}

func (s *DescribeDataSourcePageListResponseBodyResultObject) Validate() error {
	return dara.Validate(s)
}
