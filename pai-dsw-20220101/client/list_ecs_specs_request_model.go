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
	SetOrder(v string) *ListEcsSpecsRequest
	GetOrder() *string
	SetPageNumber(v int64) *ListEcsSpecsRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *ListEcsSpecsRequest
	GetPageSize() *int64
	SetResourceType(v string) *ListEcsSpecsRequest
	GetResourceType() *string
	SetSortBy(v string) *ListEcsSpecsRequest
	GetSortBy() *string
}

type ListEcsSpecsRequest struct {
	// The accelerator type.
	//
	// 	- CPU: Only CPU computing is used.
	//
	// 	- GPU: GPUs are used to accelerate computing.
	//
	// This parameter is required.
	//
	// example:
	//
	// CPU
	AcceleratorType *string `json:"AcceleratorType,omitempty" xml:"AcceleratorType,omitempty"`
	// The sorting order. Valid values:
	//
	// 	- ASC
	//
	// 	- DESC
	//
	// example:
	//
	// DESC
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// The page number. Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize     *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The field by which the query results are sorted. Set the value to gmtCreate.
	//
	// example:
	//
	// gmtCreate
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

func (s *ListEcsSpecsRequest) GetOrder() *string {
	return s.Order
}

func (s *ListEcsSpecsRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListEcsSpecsRequest) GetPageSize() *int64 {
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

func (s *ListEcsSpecsRequest) SetOrder(v string) *ListEcsSpecsRequest {
	s.Order = &v
	return s
}

func (s *ListEcsSpecsRequest) SetPageNumber(v int64) *ListEcsSpecsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListEcsSpecsRequest) SetPageSize(v int64) *ListEcsSpecsRequest {
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
