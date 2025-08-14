// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeOperationLogPageListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeOperationLogPageListResponseBody
	GetRequestId() *string
	SetCurrentPage(v int32) *DescribeOperationLogPageListResponseBody
	GetCurrentPage() *int32
	SetPageSize(v int32) *DescribeOperationLogPageListResponseBody
	GetPageSize() *int32
	SetResultObject(v []*DescribeOperationLogPageListResponseBodyResultObject) *DescribeOperationLogPageListResponseBody
	GetResultObject() []*DescribeOperationLogPageListResponseBodyResultObject
	SetTotalItem(v int32) *DescribeOperationLogPageListResponseBody
	GetTotalItem() *int32
	SetTotalPage(v int32) *DescribeOperationLogPageListResponseBody
	GetTotalPage() *int32
}

type DescribeOperationLogPageListResponseBody struct {
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
	ResultObject []*DescribeOperationLogPageListResponseBodyResultObject `json:"resultObject,omitempty" xml:"resultObject,omitempty" type:"Repeated"`
	// Total number of items
	//
	// example:
	//
	// 3
	TotalItem *int32 `json:"totalItem,omitempty" xml:"totalItem,omitempty"`
	// Total number of pages
	//
	// example:
	//
	// 3
	TotalPage *int32 `json:"totalPage,omitempty" xml:"totalPage,omitempty"`
}

func (s DescribeOperationLogPageListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeOperationLogPageListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeOperationLogPageListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeOperationLogPageListResponseBody) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeOperationLogPageListResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeOperationLogPageListResponseBody) GetResultObject() []*DescribeOperationLogPageListResponseBodyResultObject {
	return s.ResultObject
}

func (s *DescribeOperationLogPageListResponseBody) GetTotalItem() *int32 {
	return s.TotalItem
}

func (s *DescribeOperationLogPageListResponseBody) GetTotalPage() *int32 {
	return s.TotalPage
}

func (s *DescribeOperationLogPageListResponseBody) SetRequestId(v string) *DescribeOperationLogPageListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeOperationLogPageListResponseBody) SetCurrentPage(v int32) *DescribeOperationLogPageListResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *DescribeOperationLogPageListResponseBody) SetPageSize(v int32) *DescribeOperationLogPageListResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeOperationLogPageListResponseBody) SetResultObject(v []*DescribeOperationLogPageListResponseBodyResultObject) *DescribeOperationLogPageListResponseBody {
	s.ResultObject = v
	return s
}

func (s *DescribeOperationLogPageListResponseBody) SetTotalItem(v int32) *DescribeOperationLogPageListResponseBody {
	s.TotalItem = &v
	return s
}

func (s *DescribeOperationLogPageListResponseBody) SetTotalPage(v int32) *DescribeOperationLogPageListResponseBody {
	s.TotalPage = &v
	return s
}

func (s *DescribeOperationLogPageListResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeOperationLogPageListResponseBodyResultObject struct {
	// Client IP.
	//
	// example:
	//
	// 100.68.***.166
	ClientIp *string `json:"clientIp,omitempty" xml:"clientIp,omitempty"`
	// Creation time.
	//
	// example:
	//
	// 1621578648000
	GmtCreate *int64 `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// Content after operation
	//
	// example:
	//
	// @selfvariable_02
	NewContent *string `json:"newContent,omitempty" xml:"newContent,omitempty"`
	// Content before operation
	//
	// example:
	//
	// @selfvariable_02 + 1001
	OldContent *string `json:"oldContent,omitempty" xml:"oldContent,omitempty"`
	// Operation summary
	//
	// example:
	//
	// 更新事件:决策引擎可观测性持续建设_事件A(de_afghcf6411)
	OperationSummary *string `json:"operationSummary,omitempty" xml:"operationSummary,omitempty"`
	// Operation type.
	//
	// example:
	//
	// CREATE_EVENT
	OperationType *string `json:"operationType,omitempty" xml:"operationType,omitempty"`
	// Operator
	//
	// example:
	//
	// root
	UserName *string `json:"userName,omitempty" xml:"userName,omitempty"`
}

func (s DescribeOperationLogPageListResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s DescribeOperationLogPageListResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *DescribeOperationLogPageListResponseBodyResultObject) GetClientIp() *string {
	return s.ClientIp
}

func (s *DescribeOperationLogPageListResponseBodyResultObject) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *DescribeOperationLogPageListResponseBodyResultObject) GetNewContent() *string {
	return s.NewContent
}

func (s *DescribeOperationLogPageListResponseBodyResultObject) GetOldContent() *string {
	return s.OldContent
}

func (s *DescribeOperationLogPageListResponseBodyResultObject) GetOperationSummary() *string {
	return s.OperationSummary
}

func (s *DescribeOperationLogPageListResponseBodyResultObject) GetOperationType() *string {
	return s.OperationType
}

func (s *DescribeOperationLogPageListResponseBodyResultObject) GetUserName() *string {
	return s.UserName
}

func (s *DescribeOperationLogPageListResponseBodyResultObject) SetClientIp(v string) *DescribeOperationLogPageListResponseBodyResultObject {
	s.ClientIp = &v
	return s
}

func (s *DescribeOperationLogPageListResponseBodyResultObject) SetGmtCreate(v int64) *DescribeOperationLogPageListResponseBodyResultObject {
	s.GmtCreate = &v
	return s
}

func (s *DescribeOperationLogPageListResponseBodyResultObject) SetNewContent(v string) *DescribeOperationLogPageListResponseBodyResultObject {
	s.NewContent = &v
	return s
}

func (s *DescribeOperationLogPageListResponseBodyResultObject) SetOldContent(v string) *DescribeOperationLogPageListResponseBodyResultObject {
	s.OldContent = &v
	return s
}

func (s *DescribeOperationLogPageListResponseBodyResultObject) SetOperationSummary(v string) *DescribeOperationLogPageListResponseBodyResultObject {
	s.OperationSummary = &v
	return s
}

func (s *DescribeOperationLogPageListResponseBodyResultObject) SetOperationType(v string) *DescribeOperationLogPageListResponseBodyResultObject {
	s.OperationType = &v
	return s
}

func (s *DescribeOperationLogPageListResponseBodyResultObject) SetUserName(v string) *DescribeOperationLogPageListResponseBodyResultObject {
	s.UserName = &v
	return s
}

func (s *DescribeOperationLogPageListResponseBodyResultObject) Validate() error {
	return dara.Validate(s)
}
