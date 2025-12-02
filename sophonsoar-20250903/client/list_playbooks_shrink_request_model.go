// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPlaybooksShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *ListPlaybooksShrinkRequest
	GetLang() *string
	SetMaxResults(v int32) *ListPlaybooksShrinkRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListPlaybooksShrinkRequest
	GetNextToken() *string
	SetOrder(v string) *ListPlaybooksShrinkRequest
	GetOrder() *string
	SetOrderField(v string) *ListPlaybooksShrinkRequest
	GetOrderField() *string
	SetPageNumber(v int32) *ListPlaybooksShrinkRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListPlaybooksShrinkRequest
	GetPageSize() *int32
	SetPlaybookExecutionEndTime(v int64) *ListPlaybooksShrinkRequest
	GetPlaybookExecutionEndTime() *int64
	SetPlaybookExecutionStartTime(v int64) *ListPlaybooksShrinkRequest
	GetPlaybookExecutionStartTime() *int64
	SetPlaybookName(v string) *ListPlaybooksShrinkRequest
	GetPlaybookName() *string
	SetPlaybookParamTypesShrink(v string) *ListPlaybooksShrinkRequest
	GetPlaybookParamTypesShrink() *string
	SetPlaybookStatus(v int32) *ListPlaybooksShrinkRequest
	GetPlaybookStatus() *int32
	SetPlaybookType(v string) *ListPlaybooksShrinkRequest
	GetPlaybookType() *string
	SetPlaybookUuidsShrink(v string) *ListPlaybooksShrinkRequest
	GetPlaybookUuidsShrink() *string
	SetRoleFor(v int64) *ListPlaybooksShrinkRequest
	GetRoleFor() *int64
}

type ListPlaybooksShrinkRequest struct {
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
	PlaybookParamTypesShrink *string `json:"PlaybookParamTypes,omitempty" xml:"PlaybookParamTypes,omitempty"`
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
	PlaybookUuidsShrink *string `json:"PlaybookUuids,omitempty" xml:"PlaybookUuids,omitempty"`
	// User ID for the administrator to switch to another member\\"s perspective.
	//
	// example:
	//
	// 13760*****718726
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
}

func (s ListPlaybooksShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPlaybooksShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListPlaybooksShrinkRequest) GetLang() *string {
	return s.Lang
}

func (s *ListPlaybooksShrinkRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListPlaybooksShrinkRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListPlaybooksShrinkRequest) GetOrder() *string {
	return s.Order
}

func (s *ListPlaybooksShrinkRequest) GetOrderField() *string {
	return s.OrderField
}

func (s *ListPlaybooksShrinkRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListPlaybooksShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListPlaybooksShrinkRequest) GetPlaybookExecutionEndTime() *int64 {
	return s.PlaybookExecutionEndTime
}

func (s *ListPlaybooksShrinkRequest) GetPlaybookExecutionStartTime() *int64 {
	return s.PlaybookExecutionStartTime
}

func (s *ListPlaybooksShrinkRequest) GetPlaybookName() *string {
	return s.PlaybookName
}

func (s *ListPlaybooksShrinkRequest) GetPlaybookParamTypesShrink() *string {
	return s.PlaybookParamTypesShrink
}

func (s *ListPlaybooksShrinkRequest) GetPlaybookStatus() *int32 {
	return s.PlaybookStatus
}

func (s *ListPlaybooksShrinkRequest) GetPlaybookType() *string {
	return s.PlaybookType
}

func (s *ListPlaybooksShrinkRequest) GetPlaybookUuidsShrink() *string {
	return s.PlaybookUuidsShrink
}

func (s *ListPlaybooksShrinkRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *ListPlaybooksShrinkRequest) SetLang(v string) *ListPlaybooksShrinkRequest {
	s.Lang = &v
	return s
}

func (s *ListPlaybooksShrinkRequest) SetMaxResults(v int32) *ListPlaybooksShrinkRequest {
	s.MaxResults = &v
	return s
}

func (s *ListPlaybooksShrinkRequest) SetNextToken(v string) *ListPlaybooksShrinkRequest {
	s.NextToken = &v
	return s
}

func (s *ListPlaybooksShrinkRequest) SetOrder(v string) *ListPlaybooksShrinkRequest {
	s.Order = &v
	return s
}

func (s *ListPlaybooksShrinkRequest) SetOrderField(v string) *ListPlaybooksShrinkRequest {
	s.OrderField = &v
	return s
}

func (s *ListPlaybooksShrinkRequest) SetPageNumber(v int32) *ListPlaybooksShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListPlaybooksShrinkRequest) SetPageSize(v int32) *ListPlaybooksShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListPlaybooksShrinkRequest) SetPlaybookExecutionEndTime(v int64) *ListPlaybooksShrinkRequest {
	s.PlaybookExecutionEndTime = &v
	return s
}

func (s *ListPlaybooksShrinkRequest) SetPlaybookExecutionStartTime(v int64) *ListPlaybooksShrinkRequest {
	s.PlaybookExecutionStartTime = &v
	return s
}

func (s *ListPlaybooksShrinkRequest) SetPlaybookName(v string) *ListPlaybooksShrinkRequest {
	s.PlaybookName = &v
	return s
}

func (s *ListPlaybooksShrinkRequest) SetPlaybookParamTypesShrink(v string) *ListPlaybooksShrinkRequest {
	s.PlaybookParamTypesShrink = &v
	return s
}

func (s *ListPlaybooksShrinkRequest) SetPlaybookStatus(v int32) *ListPlaybooksShrinkRequest {
	s.PlaybookStatus = &v
	return s
}

func (s *ListPlaybooksShrinkRequest) SetPlaybookType(v string) *ListPlaybooksShrinkRequest {
	s.PlaybookType = &v
	return s
}

func (s *ListPlaybooksShrinkRequest) SetPlaybookUuidsShrink(v string) *ListPlaybooksShrinkRequest {
	s.PlaybookUuidsShrink = &v
	return s
}

func (s *ListPlaybooksShrinkRequest) SetRoleFor(v int64) *ListPlaybooksShrinkRequest {
	s.RoleFor = &v
	return s
}

func (s *ListPlaybooksShrinkRequest) Validate() error {
	return dara.Validate(s)
}
