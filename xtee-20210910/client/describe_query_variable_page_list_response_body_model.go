// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeQueryVariablePageListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeQueryVariablePageListResponseBody
	GetRequestId() *string
	SetCurrentPage(v int64) *DescribeQueryVariablePageListResponseBody
	GetCurrentPage() *int64
	SetPageSize(v int64) *DescribeQueryVariablePageListResponseBody
	GetPageSize() *int64
	SetResultObject(v []*DescribeQueryVariablePageListResponseBodyResultObject) *DescribeQueryVariablePageListResponseBody
	GetResultObject() []*DescribeQueryVariablePageListResponseBodyResultObject
	SetTotalItem(v int64) *DescribeQueryVariablePageListResponseBody
	GetTotalItem() *int64
	SetTotalPage(v int64) *DescribeQueryVariablePageListResponseBody
	GetTotalPage() *int64
}

type DescribeQueryVariablePageListResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// A32FE941-35F2-5378-B37C-4B8FDB16F094
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Pagination parameter, current page.
	//
	// example:
	//
	// 1
	CurrentPage *int64 `json:"currentPage,omitempty" xml:"currentPage,omitempty"`
	// Page size, default value is 10
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// Return object
	ResultObject []*DescribeQueryVariablePageListResponseBodyResultObject `json:"resultObject,omitempty" xml:"resultObject,omitempty" type:"Repeated"`
	// Total items
	//
	// example:
	//
	// 6
	TotalItem *int64 `json:"totalItem,omitempty" xml:"totalItem,omitempty"`
	// Total pages
	//
	// example:
	//
	// 1
	TotalPage *int64 `json:"totalPage,omitempty" xml:"totalPage,omitempty"`
}

func (s DescribeQueryVariablePageListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeQueryVariablePageListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeQueryVariablePageListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeQueryVariablePageListResponseBody) GetCurrentPage() *int64 {
	return s.CurrentPage
}

func (s *DescribeQueryVariablePageListResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeQueryVariablePageListResponseBody) GetResultObject() []*DescribeQueryVariablePageListResponseBodyResultObject {
	return s.ResultObject
}

func (s *DescribeQueryVariablePageListResponseBody) GetTotalItem() *int64 {
	return s.TotalItem
}

func (s *DescribeQueryVariablePageListResponseBody) GetTotalPage() *int64 {
	return s.TotalPage
}

func (s *DescribeQueryVariablePageListResponseBody) SetRequestId(v string) *DescribeQueryVariablePageListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeQueryVariablePageListResponseBody) SetCurrentPage(v int64) *DescribeQueryVariablePageListResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *DescribeQueryVariablePageListResponseBody) SetPageSize(v int64) *DescribeQueryVariablePageListResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeQueryVariablePageListResponseBody) SetResultObject(v []*DescribeQueryVariablePageListResponseBodyResultObject) *DescribeQueryVariablePageListResponseBody {
	s.ResultObject = v
	return s
}

func (s *DescribeQueryVariablePageListResponseBody) SetTotalItem(v int64) *DescribeQueryVariablePageListResponseBody {
	s.TotalItem = &v
	return s
}

func (s *DescribeQueryVariablePageListResponseBody) SetTotalPage(v int64) *DescribeQueryVariablePageListResponseBody {
	s.TotalPage = &v
	return s
}

func (s *DescribeQueryVariablePageListResponseBody) Validate() error {
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

type DescribeQueryVariablePageListResponseBodyResultObject struct {
	// Data source code.
	//
	// example:
	//
	// ds_vcaoii1697
	DataSourceCode *int64 `json:"dataSourceCode,omitempty" xml:"dataSourceCode,omitempty"`
	// Data source name.
	//
	// example:
	//
	// 姓名数据源
	DataSourceName *string `json:"dataSourceName,omitempty" xml:"dataSourceName,omitempty"`
	// Description information.
	//
	// example:
	//
	// 变量描述
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// Event name.
	//
	// example:
	//
	// 注册风险
	EventName *string `json:"eventName,omitempty" xml:"eventName,omitempty"`
	// Return value type
	//
	// example:
	//
	// STRING
	FieldType *string `json:"fieldType,omitempty" xml:"fieldType,omitempty"`
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
	// Query variable primary key ID
	//
	// example:
	//
	// 497
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// Query variable name
	//
	// example:
	//
	// 查询变量名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// Status.
	//
	// example:
	//
	// ENABLE
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// Total count
	//
	// example:
	//
	// 24
	Total *int64 `json:"total,omitempty" xml:"total,omitempty"`
	// Version number.
	//
	// example:
	//
	// 1
	Version *int64 `json:"version,omitempty" xml:"version,omitempty"`
}

func (s DescribeQueryVariablePageListResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s DescribeQueryVariablePageListResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *DescribeQueryVariablePageListResponseBodyResultObject) GetDataSourceCode() *int64 {
	return s.DataSourceCode
}

func (s *DescribeQueryVariablePageListResponseBodyResultObject) GetDataSourceName() *string {
	return s.DataSourceName
}

func (s *DescribeQueryVariablePageListResponseBodyResultObject) GetDescription() *string {
	return s.Description
}

func (s *DescribeQueryVariablePageListResponseBodyResultObject) GetEventName() *string {
	return s.EventName
}

func (s *DescribeQueryVariablePageListResponseBodyResultObject) GetFieldType() *string {
	return s.FieldType
}

func (s *DescribeQueryVariablePageListResponseBodyResultObject) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *DescribeQueryVariablePageListResponseBodyResultObject) GetGmtModified() *int64 {
	return s.GmtModified
}

func (s *DescribeQueryVariablePageListResponseBodyResultObject) GetId() *int64 {
	return s.Id
}

func (s *DescribeQueryVariablePageListResponseBodyResultObject) GetName() *string {
	return s.Name
}

func (s *DescribeQueryVariablePageListResponseBodyResultObject) GetStatus() *string {
	return s.Status
}

func (s *DescribeQueryVariablePageListResponseBodyResultObject) GetTotal() *int64 {
	return s.Total
}

func (s *DescribeQueryVariablePageListResponseBodyResultObject) GetVersion() *int64 {
	return s.Version
}

func (s *DescribeQueryVariablePageListResponseBodyResultObject) SetDataSourceCode(v int64) *DescribeQueryVariablePageListResponseBodyResultObject {
	s.DataSourceCode = &v
	return s
}

func (s *DescribeQueryVariablePageListResponseBodyResultObject) SetDataSourceName(v string) *DescribeQueryVariablePageListResponseBodyResultObject {
	s.DataSourceName = &v
	return s
}

func (s *DescribeQueryVariablePageListResponseBodyResultObject) SetDescription(v string) *DescribeQueryVariablePageListResponseBodyResultObject {
	s.Description = &v
	return s
}

func (s *DescribeQueryVariablePageListResponseBodyResultObject) SetEventName(v string) *DescribeQueryVariablePageListResponseBodyResultObject {
	s.EventName = &v
	return s
}

func (s *DescribeQueryVariablePageListResponseBodyResultObject) SetFieldType(v string) *DescribeQueryVariablePageListResponseBodyResultObject {
	s.FieldType = &v
	return s
}

func (s *DescribeQueryVariablePageListResponseBodyResultObject) SetGmtCreate(v int64) *DescribeQueryVariablePageListResponseBodyResultObject {
	s.GmtCreate = &v
	return s
}

func (s *DescribeQueryVariablePageListResponseBodyResultObject) SetGmtModified(v int64) *DescribeQueryVariablePageListResponseBodyResultObject {
	s.GmtModified = &v
	return s
}

func (s *DescribeQueryVariablePageListResponseBodyResultObject) SetId(v int64) *DescribeQueryVariablePageListResponseBodyResultObject {
	s.Id = &v
	return s
}

func (s *DescribeQueryVariablePageListResponseBodyResultObject) SetName(v string) *DescribeQueryVariablePageListResponseBodyResultObject {
	s.Name = &v
	return s
}

func (s *DescribeQueryVariablePageListResponseBodyResultObject) SetStatus(v string) *DescribeQueryVariablePageListResponseBodyResultObject {
	s.Status = &v
	return s
}

func (s *DescribeQueryVariablePageListResponseBodyResultObject) SetTotal(v int64) *DescribeQueryVariablePageListResponseBodyResultObject {
	s.Total = &v
	return s
}

func (s *DescribeQueryVariablePageListResponseBodyResultObject) SetVersion(v int64) *DescribeQueryVariablePageListResponseBodyResultObject {
	s.Version = &v
	return s
}

func (s *DescribeQueryVariablePageListResponseBodyResultObject) Validate() error {
	return dara.Validate(s)
}
