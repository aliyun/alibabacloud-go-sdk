// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRecallManagementTablesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListRecallManagementTablesRequest
	GetInstanceId() *string
	SetMaxResults(v int32) *ListRecallManagementTablesRequest
	GetMaxResults() *int32
	SetName(v string) *ListRecallManagementTablesRequest
	GetName() *string
	SetNextToken(v string) *ListRecallManagementTablesRequest
	GetNextToken() *string
	SetOrder(v string) *ListRecallManagementTablesRequest
	GetOrder() *string
	SetPageNumber(v int32) *ListRecallManagementTablesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListRecallManagementTablesRequest
	GetPageSize() *int32
	SetSortBy(v string) *ListRecallManagementTablesRequest
	GetSortBy() *string
	SetType(v string) *ListRecallManagementTablesRequest
	GetType() *string
}

type ListRecallManagementTablesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pairec-cn-test123
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 0
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// table-1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
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
	// X2I
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListRecallManagementTablesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListRecallManagementTablesRequest) GoString() string {
	return s.String()
}

func (s *ListRecallManagementTablesRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListRecallManagementTablesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListRecallManagementTablesRequest) GetName() *string {
	return s.Name
}

func (s *ListRecallManagementTablesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListRecallManagementTablesRequest) GetOrder() *string {
	return s.Order
}

func (s *ListRecallManagementTablesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListRecallManagementTablesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListRecallManagementTablesRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListRecallManagementTablesRequest) GetType() *string {
	return s.Type
}

func (s *ListRecallManagementTablesRequest) SetInstanceId(v string) *ListRecallManagementTablesRequest {
	s.InstanceId = &v
	return s
}

func (s *ListRecallManagementTablesRequest) SetMaxResults(v int32) *ListRecallManagementTablesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListRecallManagementTablesRequest) SetName(v string) *ListRecallManagementTablesRequest {
	s.Name = &v
	return s
}

func (s *ListRecallManagementTablesRequest) SetNextToken(v string) *ListRecallManagementTablesRequest {
	s.NextToken = &v
	return s
}

func (s *ListRecallManagementTablesRequest) SetOrder(v string) *ListRecallManagementTablesRequest {
	s.Order = &v
	return s
}

func (s *ListRecallManagementTablesRequest) SetPageNumber(v int32) *ListRecallManagementTablesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListRecallManagementTablesRequest) SetPageSize(v int32) *ListRecallManagementTablesRequest {
	s.PageSize = &v
	return s
}

func (s *ListRecallManagementTablesRequest) SetSortBy(v string) *ListRecallManagementTablesRequest {
	s.SortBy = &v
	return s
}

func (s *ListRecallManagementTablesRequest) SetType(v string) *ListRecallManagementTablesRequest {
	s.Type = &v
	return s
}

func (s *ListRecallManagementTablesRequest) Validate() error {
	return dara.Validate(s)
}
