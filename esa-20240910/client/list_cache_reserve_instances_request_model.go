// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCacheReserveInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListCacheReserveInstancesRequest
	GetInstanceId() *string
	SetPageNumber(v int32) *ListCacheReserveInstancesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListCacheReserveInstancesRequest
	GetPageSize() *int32
	SetSortBy(v string) *ListCacheReserveInstancesRequest
	GetSortBy() *string
	SetSortOrder(v string) *ListCacheReserveInstancesRequest
	GetSortOrder() *string
	SetStatus(v string) *ListCacheReserveInstancesRequest
	GetStatus() *string
}

type ListCacheReserveInstancesRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SortBy     *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	SortOrder  *string `json:"SortOrder,omitempty" xml:"SortOrder,omitempty"`
	// if can be null:
	// false
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListCacheReserveInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCacheReserveInstancesRequest) GoString() string {
	return s.String()
}

func (s *ListCacheReserveInstancesRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListCacheReserveInstancesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListCacheReserveInstancesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListCacheReserveInstancesRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListCacheReserveInstancesRequest) GetSortOrder() *string {
	return s.SortOrder
}

func (s *ListCacheReserveInstancesRequest) GetStatus() *string {
	return s.Status
}

func (s *ListCacheReserveInstancesRequest) SetInstanceId(v string) *ListCacheReserveInstancesRequest {
	s.InstanceId = &v
	return s
}

func (s *ListCacheReserveInstancesRequest) SetPageNumber(v int32) *ListCacheReserveInstancesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListCacheReserveInstancesRequest) SetPageSize(v int32) *ListCacheReserveInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *ListCacheReserveInstancesRequest) SetSortBy(v string) *ListCacheReserveInstancesRequest {
	s.SortBy = &v
	return s
}

func (s *ListCacheReserveInstancesRequest) SetSortOrder(v string) *ListCacheReserveInstancesRequest {
	s.SortOrder = &v
	return s
}

func (s *ListCacheReserveInstancesRequest) SetStatus(v string) *ListCacheReserveInstancesRequest {
	s.Status = &v
	return s
}

func (s *ListCacheReserveInstancesRequest) Validate() error {
	return dara.Validate(s)
}
