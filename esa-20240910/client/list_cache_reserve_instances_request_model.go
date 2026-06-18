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
	// The instance ID.
	//
	// example:
	//
	// sp-xcdn-96wblslz****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values: **1 to 500**. The default value is **500**.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The field to sort the results by. Valid values:
	//
	// - **ExpireTime**: Expiration time.
	//
	// - **CreateTime**: Creation time.
	//
	// example:
	//
	// ExpireTime
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// The sort order. Valid values:
	//
	// - **asc**: Ascending order.
	//
	// - **desc**: Descending order.
	//
	// example:
	//
	// desc
	SortOrder *string `json:"SortOrder,omitempty" xml:"SortOrder,omitempty"`
	// The status of the cache reserve instance. Valid values:
	//
	// - **online**: The instance is running normally.
	//
	// - **offline**: The instance has expired and is unavailable but remains within the grace period.
	//
	// - **disable**: The instance is disabled.
	//
	// - **overdue**: The instance is suspended due to an overdue payment.
	//
	// if can be null:
	// false
	//
	// example:
	//
	// online
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
