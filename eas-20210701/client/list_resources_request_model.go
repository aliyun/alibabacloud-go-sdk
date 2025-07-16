// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListResourcesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOrder(v string) *ListResourcesRequest
	GetOrder() *string
	SetPageNumber(v int32) *ListResourcesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListResourcesRequest
	GetPageSize() *int32
	SetResourceId(v string) *ListResourcesRequest
	GetResourceId() *string
	SetResourceName(v string) *ListResourcesRequest
	GetResourceName() *string
	SetResourceStatus(v string) *ListResourcesRequest
	GetResourceStatus() *string
	SetResourceType(v string) *ListResourcesRequest
	GetResourceType() *string
	SetSort(v string) *ListResourcesRequest
	GetSort() *string
}

type ListResourcesRequest struct {
	// The sorting order. Valid values:
	//
	// 	- Desc
	//
	// 	- Asc
	//
	// example:
	//
	// Desc
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// The page number. Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 100.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the resource group. You can call the [CreateResource](https://help.aliyun.com/document_detail/412111.html) operation to query the ID of the resource group.
	//
	// example:
	//
	// eas-r-h7lcw24dyqztwxxxxxx
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The name of the resource group. You can call the [CreateResource](https://help.aliyun.com/document_detail/412111.html) operation to query the name of the resource group.
	//
	// example:
	//
	// MyResource
	ResourceName *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	// The resource group status.
	//
	// example:
	//
	// ResourceReady
	ResourceStatus *string `json:"ResourceStatus,omitempty" xml:"ResourceStatus,omitempty"`
	// The type of the resource group. Valid values:
	//
	// 	- Dedicated: the dedicated resource group.
	//
	// 	- SelfManaged: the self-managed resource group.
	//
	// example:
	//
	// Dedicated
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The condition by which the results are sorted. By default, the query results are sorted by the timestamp type in descending order.
	//
	// Valid values:
	//
	// 	- PrePaidInstanceCount
	//
	// 	- CpuCount
	//
	// 	- Memory
	//
	// 	- CreateTime
	//
	// 	- PostPaidInstanceCount
	//
	// 	- MemoryUsed
	//
	// 	- GpuCount
	//
	// 	- GpuUsed
	//
	// 	- CpuUsed
	//
	// 	- ServiceCount
	//
	// example:
	//
	// CreateTime
	Sort *string `json:"Sort,omitempty" xml:"Sort,omitempty"`
}

func (s ListResourcesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListResourcesRequest) GoString() string {
	return s.String()
}

func (s *ListResourcesRequest) GetOrder() *string {
	return s.Order
}

func (s *ListResourcesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListResourcesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListResourcesRequest) GetResourceId() *string {
	return s.ResourceId
}

func (s *ListResourcesRequest) GetResourceName() *string {
	return s.ResourceName
}

func (s *ListResourcesRequest) GetResourceStatus() *string {
	return s.ResourceStatus
}

func (s *ListResourcesRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListResourcesRequest) GetSort() *string {
	return s.Sort
}

func (s *ListResourcesRequest) SetOrder(v string) *ListResourcesRequest {
	s.Order = &v
	return s
}

func (s *ListResourcesRequest) SetPageNumber(v int32) *ListResourcesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListResourcesRequest) SetPageSize(v int32) *ListResourcesRequest {
	s.PageSize = &v
	return s
}

func (s *ListResourcesRequest) SetResourceId(v string) *ListResourcesRequest {
	s.ResourceId = &v
	return s
}

func (s *ListResourcesRequest) SetResourceName(v string) *ListResourcesRequest {
	s.ResourceName = &v
	return s
}

func (s *ListResourcesRequest) SetResourceStatus(v string) *ListResourcesRequest {
	s.ResourceStatus = &v
	return s
}

func (s *ListResourcesRequest) SetResourceType(v string) *ListResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *ListResourcesRequest) SetSort(v string) *ListResourcesRequest {
	s.Sort = &v
	return s
}

func (s *ListResourcesRequest) Validate() error {
	return dara.Validate(s)
}
