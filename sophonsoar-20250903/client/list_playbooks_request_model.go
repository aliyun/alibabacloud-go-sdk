// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPlaybooksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *ListPlaybooksRequest
	GetLang() *string
	SetMaxResults(v int32) *ListPlaybooksRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListPlaybooksRequest
	GetNextToken() *string
	SetOrder(v string) *ListPlaybooksRequest
	GetOrder() *string
	SetOrderField(v string) *ListPlaybooksRequest
	GetOrderField() *string
	SetPageNumber(v int32) *ListPlaybooksRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListPlaybooksRequest
	GetPageSize() *int32
	SetPlaybookExecutionEndTime(v int64) *ListPlaybooksRequest
	GetPlaybookExecutionEndTime() *int64
	SetPlaybookExecutionStartTime(v int64) *ListPlaybooksRequest
	GetPlaybookExecutionStartTime() *int64
	SetPlaybookName(v string) *ListPlaybooksRequest
	GetPlaybookName() *string
	SetPlaybookParamTypes(v []*string) *ListPlaybooksRequest
	GetPlaybookParamTypes() []*string
	SetPlaybookStatus(v int32) *ListPlaybooksRequest
	GetPlaybookStatus() *int32
	SetPlaybookType(v string) *ListPlaybooksRequest
	GetPlaybookType() *string
	SetPlaybookUuids(v []*string) *ListPlaybooksRequest
	GetPlaybookUuids() []*string
	SetRoleFor(v int64) *ListPlaybooksRequest
	GetRoleFor() *int64
}

type ListPlaybooksRequest struct {
	// Language type for request and response messages.
	//
	// - **zh*	- (default): Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Maximum number of results returned in a single request. Range: 1~100.
	//
	// example:
	//
	// 100
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// Token for the next query start.
	//
	// example:
	//
	// 7fbb1c****07c1f79a740f039a8dcfda
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Sorting logic, default is **desc**, with values as follows:
	//
	// - **desc**: Descending order.
	//
	// - **asc**: Ascending order.
	//
	// example:
	//
	// desc
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// Sorting field. Values:
	//
	// - **UpdateTime**: Sort by update time.
	//
	// - **ExecutionTime**: Sort by the latest execution time.
	//
	// example:
	//
	// UpdateTime
	OrderField *string `json:"OrderField,omitempty" xml:"OrderField,omitempty"`
	// Page number for pagination, default value is 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// Number of items per page for pagination. Range: 1~100.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// End time of the latest execution of the playbook.
	//
	// example:
	//
	// 1731379251000
	PlaybookExecutionEndTime *int64 `json:"PlaybookExecutionEndTime,omitempty" xml:"PlaybookExecutionEndTime,omitempty"`
	// Start time of the latest execution of the playbook.
	//
	// example:
	//
	// 1731378251000
	PlaybookExecutionStartTime *int64 `json:"PlaybookExecutionStartTime,omitempty" xml:"PlaybookExecutionStartTime,omitempty"`
	// Name of the playbook, supports fuzzy search.
	//
	// example:
	//
	// demo_test
	PlaybookName *string `json:"PlaybookName,omitempty" xml:"PlaybookName,omitempty"`
	// Collection of input parameter types for the playbook.
	PlaybookParamTypes []*string `json:"PlaybookParamTypes,omitempty" xml:"PlaybookParamTypes,omitempty" type:"Repeated"`
	// Publication status of the playbook, with values as follows:
	//
	// - **0**: Unpublished.
	//
	// - **1**: Published.
	//
	// example:
	//
	// 1
	PlaybookStatus *int32 `json:"PlaybookStatus,omitempty" xml:"PlaybookStatus,omitempty"`
	// Type of the playbook, with values as follows:
	//
	// - **preset**: Predefined playbook.
	//
	// - **user**: Custom playbook.
	//
	// - **component**: Security response component.
	//
	// example:
	//
	// preset
	PlaybookType *string `json:"PlaybookType,omitempty" xml:"PlaybookType,omitempty"`
	// Collection of UUIDs of playbooks.
	PlaybookUuids []*string `json:"PlaybookUuids,omitempty" xml:"PlaybookUuids,omitempty" type:"Repeated"`
	// User ID for the administrator to switch to another member\\"s perspective.
	//
	// example:
	//
	// 13760*****718726
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
}

func (s ListPlaybooksRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPlaybooksRequest) GoString() string {
	return s.String()
}

func (s *ListPlaybooksRequest) GetLang() *string {
	return s.Lang
}

func (s *ListPlaybooksRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListPlaybooksRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListPlaybooksRequest) GetOrder() *string {
	return s.Order
}

func (s *ListPlaybooksRequest) GetOrderField() *string {
	return s.OrderField
}

func (s *ListPlaybooksRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListPlaybooksRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListPlaybooksRequest) GetPlaybookExecutionEndTime() *int64 {
	return s.PlaybookExecutionEndTime
}

func (s *ListPlaybooksRequest) GetPlaybookExecutionStartTime() *int64 {
	return s.PlaybookExecutionStartTime
}

func (s *ListPlaybooksRequest) GetPlaybookName() *string {
	return s.PlaybookName
}

func (s *ListPlaybooksRequest) GetPlaybookParamTypes() []*string {
	return s.PlaybookParamTypes
}

func (s *ListPlaybooksRequest) GetPlaybookStatus() *int32 {
	return s.PlaybookStatus
}

func (s *ListPlaybooksRequest) GetPlaybookType() *string {
	return s.PlaybookType
}

func (s *ListPlaybooksRequest) GetPlaybookUuids() []*string {
	return s.PlaybookUuids
}

func (s *ListPlaybooksRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *ListPlaybooksRequest) SetLang(v string) *ListPlaybooksRequest {
	s.Lang = &v
	return s
}

func (s *ListPlaybooksRequest) SetMaxResults(v int32) *ListPlaybooksRequest {
	s.MaxResults = &v
	return s
}

func (s *ListPlaybooksRequest) SetNextToken(v string) *ListPlaybooksRequest {
	s.NextToken = &v
	return s
}

func (s *ListPlaybooksRequest) SetOrder(v string) *ListPlaybooksRequest {
	s.Order = &v
	return s
}

func (s *ListPlaybooksRequest) SetOrderField(v string) *ListPlaybooksRequest {
	s.OrderField = &v
	return s
}

func (s *ListPlaybooksRequest) SetPageNumber(v int32) *ListPlaybooksRequest {
	s.PageNumber = &v
	return s
}

func (s *ListPlaybooksRequest) SetPageSize(v int32) *ListPlaybooksRequest {
	s.PageSize = &v
	return s
}

func (s *ListPlaybooksRequest) SetPlaybookExecutionEndTime(v int64) *ListPlaybooksRequest {
	s.PlaybookExecutionEndTime = &v
	return s
}

func (s *ListPlaybooksRequest) SetPlaybookExecutionStartTime(v int64) *ListPlaybooksRequest {
	s.PlaybookExecutionStartTime = &v
	return s
}

func (s *ListPlaybooksRequest) SetPlaybookName(v string) *ListPlaybooksRequest {
	s.PlaybookName = &v
	return s
}

func (s *ListPlaybooksRequest) SetPlaybookParamTypes(v []*string) *ListPlaybooksRequest {
	s.PlaybookParamTypes = v
	return s
}

func (s *ListPlaybooksRequest) SetPlaybookStatus(v int32) *ListPlaybooksRequest {
	s.PlaybookStatus = &v
	return s
}

func (s *ListPlaybooksRequest) SetPlaybookType(v string) *ListPlaybooksRequest {
	s.PlaybookType = &v
	return s
}

func (s *ListPlaybooksRequest) SetPlaybookUuids(v []*string) *ListPlaybooksRequest {
	s.PlaybookUuids = v
	return s
}

func (s *ListPlaybooksRequest) SetRoleFor(v int64) *ListPlaybooksRequest {
	s.RoleFor = &v
	return s
}

func (s *ListPlaybooksRequest) Validate() error {
	return dara.Validate(s)
}
