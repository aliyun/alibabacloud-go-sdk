// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListComponentsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetComponentIds(v string) *ListComponentsRequest
	GetComponentIds() *string
	SetName(v string) *ListComponentsRequest
	GetName() *string
	SetOrder(v string) *ListComponentsRequest
	GetOrder() *string
	SetPageNumber(v int32) *ListComponentsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListComponentsRequest
	GetPageSize() *int32
	SetSortBy(v string) *ListComponentsRequest
	GetSortBy() *string
	SetVersion(v string) *ListComponentsRequest
	GetVersion() *string
}

type ListComponentsRequest struct {
	// example:
	//
	// cmpt-my1tk3jggooi5uwks5,cmpt-n69468yvjz8d12as7d,cmpt-tga4omjrepklkg1onn
	ComponentIds *string `json:"ComponentIds,omitempty" xml:"ComponentIds,omitempty"`
	// example:
	//
	// dataset-accelerator
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// Desc
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
	// v1.0.0
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s ListComponentsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListComponentsRequest) GoString() string {
	return s.String()
}

func (s *ListComponentsRequest) GetComponentIds() *string {
	return s.ComponentIds
}

func (s *ListComponentsRequest) GetName() *string {
	return s.Name
}

func (s *ListComponentsRequest) GetOrder() *string {
	return s.Order
}

func (s *ListComponentsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListComponentsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListComponentsRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListComponentsRequest) GetVersion() *string {
	return s.Version
}

func (s *ListComponentsRequest) SetComponentIds(v string) *ListComponentsRequest {
	s.ComponentIds = &v
	return s
}

func (s *ListComponentsRequest) SetName(v string) *ListComponentsRequest {
	s.Name = &v
	return s
}

func (s *ListComponentsRequest) SetOrder(v string) *ListComponentsRequest {
	s.Order = &v
	return s
}

func (s *ListComponentsRequest) SetPageNumber(v int32) *ListComponentsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListComponentsRequest) SetPageSize(v int32) *ListComponentsRequest {
	s.PageSize = &v
	return s
}

func (s *ListComponentsRequest) SetSortBy(v string) *ListComponentsRequest {
	s.SortBy = &v
	return s
}

func (s *ListComponentsRequest) SetVersion(v string) *ListComponentsRequest {
	s.Version = &v
	return s
}

func (s *ListComponentsRequest) Validate() error {
	return dara.Validate(s)
}
