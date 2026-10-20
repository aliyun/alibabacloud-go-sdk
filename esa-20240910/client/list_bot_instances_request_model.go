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
	// The instance ID. You can obtain the ID by calling the [ListSites](https://help.aliyun.com/document_detail/2850189.html) operation.
	//
	// example:
	//
	// sp-xcdn-96wblslz****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The page number Settings for paging in a paged query. Default value: 1. Valid values: 1 to 100000.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page for paging in a paged query. Valid values: 1 to 500.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The sort field. By default, instances are sorted by creation time. Valid values:
	//
	// - **CreateTime**: creation time.
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
	// - **online**: normal.
	//
	// - **offline**: expired but not overdue, and unavailable.
	//
	// - **disable**: released.
	//
	// - **overdue**: overdue payment.
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
