// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRecallManagementServicesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListRecallManagementServicesRequest
	GetInstanceId() *string
	SetMaxResults(v int32) *ListRecallManagementServicesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListRecallManagementServicesRequest
	GetNextToken() *string
	SetOrder(v string) *ListRecallManagementServicesRequest
	GetOrder() *string
	SetPageNumber(v int64) *ListRecallManagementServicesRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *ListRecallManagementServicesRequest
	GetPageSize() *int64
	SetSortBy(v string) *ListRecallManagementServicesRequest
	GetSortBy() *string
}

type ListRecallManagementServicesRequest struct {
	// example:
	//
	// pairec-cn-test123
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// ""
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// 0
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// ASC
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 50
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// GmtCreateTime
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
}

func (s ListRecallManagementServicesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListRecallManagementServicesRequest) GoString() string {
	return s.String()
}

func (s *ListRecallManagementServicesRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListRecallManagementServicesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListRecallManagementServicesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListRecallManagementServicesRequest) GetOrder() *string {
	return s.Order
}

func (s *ListRecallManagementServicesRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListRecallManagementServicesRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListRecallManagementServicesRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListRecallManagementServicesRequest) SetInstanceId(v string) *ListRecallManagementServicesRequest {
	s.InstanceId = &v
	return s
}

func (s *ListRecallManagementServicesRequest) SetMaxResults(v int32) *ListRecallManagementServicesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListRecallManagementServicesRequest) SetNextToken(v string) *ListRecallManagementServicesRequest {
	s.NextToken = &v
	return s
}

func (s *ListRecallManagementServicesRequest) SetOrder(v string) *ListRecallManagementServicesRequest {
	s.Order = &v
	return s
}

func (s *ListRecallManagementServicesRequest) SetPageNumber(v int64) *ListRecallManagementServicesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListRecallManagementServicesRequest) SetPageSize(v int64) *ListRecallManagementServicesRequest {
	s.PageSize = &v
	return s
}

func (s *ListRecallManagementServicesRequest) SetSortBy(v string) *ListRecallManagementServicesRequest {
	s.SortBy = &v
	return s
}

func (s *ListRecallManagementServicesRequest) Validate() error {
	return dara.Validate(s)
}
