// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRecallManagementJobsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCondition(v map[string]interface{}) *ListRecallManagementJobsRequest
	GetCondition() map[string]interface{}
	SetInstanceId(v string) *ListRecallManagementJobsRequest
	GetInstanceId() *string
	SetMaxResults(v int32) *ListRecallManagementJobsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListRecallManagementJobsRequest
	GetNextToken() *string
	SetOrder(v string) *ListRecallManagementJobsRequest
	GetOrder() *string
	SetPageNumber(v int32) *ListRecallManagementJobsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListRecallManagementJobsRequest
	GetPageSize() *int32
	SetSortBy(v string) *ListRecallManagementJobsRequest
	GetSortBy() *string
	SetType(v string) *ListRecallManagementJobsRequest
	GetType() *string
}

type ListRecallManagementJobsRequest struct {
	// example:
	//
	// {"RecallManagementTableId":"1"}
	Condition map[string]interface{} `json:"Condition,omitempty" xml:"Condition,omitempty"`
	// example:
	//
	// pairec-test1
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 0
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// ""
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// ASC
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// GmtCreateTime
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// example:
	//
	// Table
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListRecallManagementJobsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListRecallManagementJobsRequest) GoString() string {
	return s.String()
}

func (s *ListRecallManagementJobsRequest) GetCondition() map[string]interface{} {
	return s.Condition
}

func (s *ListRecallManagementJobsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListRecallManagementJobsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListRecallManagementJobsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListRecallManagementJobsRequest) GetOrder() *string {
	return s.Order
}

func (s *ListRecallManagementJobsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListRecallManagementJobsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListRecallManagementJobsRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListRecallManagementJobsRequest) GetType() *string {
	return s.Type
}

func (s *ListRecallManagementJobsRequest) SetCondition(v map[string]interface{}) *ListRecallManagementJobsRequest {
	s.Condition = v
	return s
}

func (s *ListRecallManagementJobsRequest) SetInstanceId(v string) *ListRecallManagementJobsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListRecallManagementJobsRequest) SetMaxResults(v int32) *ListRecallManagementJobsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListRecallManagementJobsRequest) SetNextToken(v string) *ListRecallManagementJobsRequest {
	s.NextToken = &v
	return s
}

func (s *ListRecallManagementJobsRequest) SetOrder(v string) *ListRecallManagementJobsRequest {
	s.Order = &v
	return s
}

func (s *ListRecallManagementJobsRequest) SetPageNumber(v int32) *ListRecallManagementJobsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListRecallManagementJobsRequest) SetPageSize(v int32) *ListRecallManagementJobsRequest {
	s.PageSize = &v
	return s
}

func (s *ListRecallManagementJobsRequest) SetSortBy(v string) *ListRecallManagementJobsRequest {
	s.SortBy = &v
	return s
}

func (s *ListRecallManagementJobsRequest) SetType(v string) *ListRecallManagementJobsRequest {
	s.Type = &v
	return s
}

func (s *ListRecallManagementJobsRequest) Validate() error {
	return dara.Validate(s)
}
