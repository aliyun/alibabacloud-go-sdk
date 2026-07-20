// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBotInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListBotInstancesRequest
	GetInstanceId() *string
	SetPageNumber(v int32) *ListBotInstancesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListBotInstancesRequest
	GetPageSize() *int32
	SetSortBy(v string) *ListBotInstancesRequest
	GetSortBy() *string
	SetSortOrder(v string) *ListBotInstancesRequest
	GetSortOrder() *string
	SetStatus(v string) *ListBotInstancesRequest
	GetStatus() *string
}

type ListBotInstancesRequest struct {
	// The instance ID. You can call the [ListSites](https://help.aliyun.com/document_detail/2850189.html) operation to obtain the ID.
	//
	// example:
	//
	// sp-xcdn-96wblslz****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The page number for the paged query. Settings this parameter for paging. Default value: 1. Valid values: 1 to 100000.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page for the paged query. This parameter is used for paging. Valid values: 1 to 500.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The field by which to sort the results. By default, the results are sorted by purchase time. Valid values:
	//
	// - **CreateTime**: purchase time.
	//
	// - **ExpireTime**: expiration time.
	//
	// example:
	//
	// CreateTime
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// The sort order. Default value: desc. Valid values:
	//
	// - **asc**: ascending order.
	//
	// - **desc**: descending order.
	//
	// example:
	//
	// desc
	SortOrder *string `json:"SortOrder,omitempty" xml:"SortOrder,omitempty"`
	// The instance status. Valid values:
	//
	// - **online**: The instance is running normally.
	//
	// - **offline**: The instance has expired but has not exceeded the retention period and is unavailable.
	//
	// - **disable**: The instance has been released.
	//
	// - **overdue**: The instance has an overdue payment.
	//
	// example:
	//
	// online
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListBotInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListBotInstancesRequest) GoString() string {
	return s.String()
}

func (s *ListBotInstancesRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListBotInstancesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListBotInstancesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListBotInstancesRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListBotInstancesRequest) GetSortOrder() *string {
	return s.SortOrder
}

func (s *ListBotInstancesRequest) GetStatus() *string {
	return s.Status
}

func (s *ListBotInstancesRequest) SetInstanceId(v string) *ListBotInstancesRequest {
	s.InstanceId = &v
	return s
}

func (s *ListBotInstancesRequest) SetPageNumber(v int32) *ListBotInstancesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListBotInstancesRequest) SetPageSize(v int32) *ListBotInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *ListBotInstancesRequest) SetSortBy(v string) *ListBotInstancesRequest {
	s.SortBy = &v
	return s
}

func (s *ListBotInstancesRequest) SetSortOrder(v string) *ListBotInstancesRequest {
	s.SortOrder = &v
	return s
}

func (s *ListBotInstancesRequest) SetStatus(v string) *ListBotInstancesRequest {
	s.Status = &v
	return s
}

func (s *ListBotInstancesRequest) Validate() error {
	return dara.Validate(s)
}
