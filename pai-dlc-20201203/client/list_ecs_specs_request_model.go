// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEcsSpecsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceleratorType(v string) *ListEcsSpecsRequest
	GetAcceleratorType() *string
	SetInstanceTypes(v string) *ListEcsSpecsRequest
	GetInstanceTypes() *string
	SetOrder(v string) *ListEcsSpecsRequest
	GetOrder() *string
	SetPageNumber(v int32) *ListEcsSpecsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListEcsSpecsRequest
	GetPageSize() *int32
	SetResourceType(v string) *ListEcsSpecsRequest
	GetResourceType() *string
	SetSortBy(v string) *ListEcsSpecsRequest
	GetSortBy() *string
}

type ListEcsSpecsRequest struct {
	// Filter by accelerator type. Valid values:
	//
	// 	- CPU
	//
	// 	- GPU
	//
	// example:
	//
	// GPU
	AcceleratorType *string `json:"AcceleratorType,omitempty" xml:"AcceleratorType,omitempty"`
	// The instance types to query. Separate the types with commas (,).
	//
	// example:
	//
	// ecs.g6.large,ecs.g6.xlarge
	InstanceTypes *string `json:"InstanceTypes,omitempty" xml:"InstanceTypes,omitempty"`
	// The sorting order. Valid values:
	//
	// 	- desc: descending order.
	//
	// 	- asc: ascending order.
	//
	// example:
	//
	// desc
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// The number of the page to query. The start value is 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The type of the resource. Valid values:
	//
	// 	- ECS
	//
	// 	- Lingjun
	//
	// example:
	//
	// ECS
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The field based on which the results are sorted. Valid values:
	//
	// 	- CPU
	//
	// 	- GPU
	//
	// 	- Memory
	//
	// 	- GmtCreateTime
	//
	// example:
	//
	// Gpu
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
}

func (s ListEcsSpecsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListEcsSpecsRequest) GoString() string {
	return s.String()
}

func (s *ListEcsSpecsRequest) GetAcceleratorType() *string {
	return s.AcceleratorType
}

func (s *ListEcsSpecsRequest) GetInstanceTypes() *string {
	return s.InstanceTypes
}

func (s *ListEcsSpecsRequest) GetOrder() *string {
	return s.Order
}

func (s *ListEcsSpecsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListEcsSpecsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListEcsSpecsRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListEcsSpecsRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListEcsSpecsRequest) SetAcceleratorType(v string) *ListEcsSpecsRequest {
	s.AcceleratorType = &v
	return s
}

func (s *ListEcsSpecsRequest) SetInstanceTypes(v string) *ListEcsSpecsRequest {
	s.InstanceTypes = &v
	return s
}

func (s *ListEcsSpecsRequest) SetOrder(v string) *ListEcsSpecsRequest {
	s.Order = &v
	return s
}

func (s *ListEcsSpecsRequest) SetPageNumber(v int32) *ListEcsSpecsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListEcsSpecsRequest) SetPageSize(v int32) *ListEcsSpecsRequest {
	s.PageSize = &v
	return s
}

func (s *ListEcsSpecsRequest) SetResourceType(v string) *ListEcsSpecsRequest {
	s.ResourceType = &v
	return s
}

func (s *ListEcsSpecsRequest) SetSortBy(v string) *ListEcsSpecsRequest {
	s.SortBy = &v
	return s
}

func (s *ListEcsSpecsRequest) Validate() error {
	return dara.Validate(s)
}
