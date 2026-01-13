// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRecallManagementServiceVersionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListRecallManagementServiceVersionsRequest
	GetInstanceId() *string
	SetOrder(v string) *ListRecallManagementServiceVersionsRequest
	GetOrder() *string
	SetPageNumber(v int64) *ListRecallManagementServiceVersionsRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *ListRecallManagementServiceVersionsRequest
	GetPageSize() *int64
	SetSortBy(v string) *ListRecallManagementServiceVersionsRequest
	GetSortBy() *string
}

type ListRecallManagementServiceVersionsRequest struct {
	// example:
	//
	// pairec-cn-test123
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
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

func (s ListRecallManagementServiceVersionsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListRecallManagementServiceVersionsRequest) GoString() string {
	return s.String()
}

func (s *ListRecallManagementServiceVersionsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListRecallManagementServiceVersionsRequest) GetOrder() *string {
	return s.Order
}

func (s *ListRecallManagementServiceVersionsRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListRecallManagementServiceVersionsRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListRecallManagementServiceVersionsRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListRecallManagementServiceVersionsRequest) SetInstanceId(v string) *ListRecallManagementServiceVersionsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListRecallManagementServiceVersionsRequest) SetOrder(v string) *ListRecallManagementServiceVersionsRequest {
	s.Order = &v
	return s
}

func (s *ListRecallManagementServiceVersionsRequest) SetPageNumber(v int64) *ListRecallManagementServiceVersionsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListRecallManagementServiceVersionsRequest) SetPageSize(v int64) *ListRecallManagementServiceVersionsRequest {
	s.PageSize = &v
	return s
}

func (s *ListRecallManagementServiceVersionsRequest) SetSortBy(v string) *ListRecallManagementServiceVersionsRequest {
	s.SortBy = &v
	return s
}

func (s *ListRecallManagementServiceVersionsRequest) Validate() error {
	return dara.Validate(s)
}
