// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListResourceGroupsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetComputingResourceProvider(v string) *ListResourceGroupsRequest
	GetComputingResourceProvider() *string
	SetHasResource(v bool) *ListResourceGroupsRequest
	GetHasResource() *bool
	SetName(v string) *ListResourceGroupsRequest
	GetName() *string
	SetOrder(v string) *ListResourceGroupsRequest
	GetOrder() *string
	SetPageNumber(v int64) *ListResourceGroupsRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *ListResourceGroupsRequest
	GetPageSize() *int64
	SetResourceGroupIDs(v string) *ListResourceGroupsRequest
	GetResourceGroupIDs() *string
	SetResourceType(v string) *ListResourceGroupsRequest
	GetResourceType() *string
	SetShowAll(v bool) *ListResourceGroupsRequest
	GetShowAll() *bool
	SetSortBy(v string) *ListResourceGroupsRequest
	GetSortBy() *string
	SetStatus(v string) *ListResourceGroupsRequest
	GetStatus() *string
	SetVersions(v string) *ListResourceGroupsRequest
	GetVersions() *string
}

type ListResourceGroupsRequest struct {
	// example:
	//
	// Ecs
	ComputingResourceProvider *string `json:"ComputingResourceProvider,omitempty" xml:"ComputingResourceProvider,omitempty"`
	HasResource               *bool   `json:"HasResource,omitempty" xml:"HasResource,omitempty"`
	// example:
	//
	// rgf0zhfqn1d4ity2
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// desc
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// example:
	//
	// 2
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize         *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceGroupIDs *string `json:"ResourceGroupIDs,omitempty" xml:"ResourceGroupIDs,omitempty"`
	// example:
	//
	// Lingjun
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// example:
	//
	// false
	ShowAll *bool `json:"ShowAll,omitempty" xml:"ShowAll,omitempty"`
	// example:
	//
	// DisplayName
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// example:
	//
	// Creating
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 1.0
	Versions *string `json:"Versions,omitempty" xml:"Versions,omitempty"`
}

func (s ListResourceGroupsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListResourceGroupsRequest) GoString() string {
	return s.String()
}

func (s *ListResourceGroupsRequest) GetComputingResourceProvider() *string {
	return s.ComputingResourceProvider
}

func (s *ListResourceGroupsRequest) GetHasResource() *bool {
	return s.HasResource
}

func (s *ListResourceGroupsRequest) GetName() *string {
	return s.Name
}

func (s *ListResourceGroupsRequest) GetOrder() *string {
	return s.Order
}

func (s *ListResourceGroupsRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListResourceGroupsRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListResourceGroupsRequest) GetResourceGroupIDs() *string {
	return s.ResourceGroupIDs
}

func (s *ListResourceGroupsRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListResourceGroupsRequest) GetShowAll() *bool {
	return s.ShowAll
}

func (s *ListResourceGroupsRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListResourceGroupsRequest) GetStatus() *string {
	return s.Status
}

func (s *ListResourceGroupsRequest) GetVersions() *string {
	return s.Versions
}

func (s *ListResourceGroupsRequest) SetComputingResourceProvider(v string) *ListResourceGroupsRequest {
	s.ComputingResourceProvider = &v
	return s
}

func (s *ListResourceGroupsRequest) SetHasResource(v bool) *ListResourceGroupsRequest {
	s.HasResource = &v
	return s
}

func (s *ListResourceGroupsRequest) SetName(v string) *ListResourceGroupsRequest {
	s.Name = &v
	return s
}

func (s *ListResourceGroupsRequest) SetOrder(v string) *ListResourceGroupsRequest {
	s.Order = &v
	return s
}

func (s *ListResourceGroupsRequest) SetPageNumber(v int64) *ListResourceGroupsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListResourceGroupsRequest) SetPageSize(v int64) *ListResourceGroupsRequest {
	s.PageSize = &v
	return s
}

func (s *ListResourceGroupsRequest) SetResourceGroupIDs(v string) *ListResourceGroupsRequest {
	s.ResourceGroupIDs = &v
	return s
}

func (s *ListResourceGroupsRequest) SetResourceType(v string) *ListResourceGroupsRequest {
	s.ResourceType = &v
	return s
}

func (s *ListResourceGroupsRequest) SetShowAll(v bool) *ListResourceGroupsRequest {
	s.ShowAll = &v
	return s
}

func (s *ListResourceGroupsRequest) SetSortBy(v string) *ListResourceGroupsRequest {
	s.SortBy = &v
	return s
}

func (s *ListResourceGroupsRequest) SetStatus(v string) *ListResourceGroupsRequest {
	s.Status = &v
	return s
}

func (s *ListResourceGroupsRequest) SetVersions(v string) *ListResourceGroupsRequest {
	s.Versions = &v
	return s
}

func (s *ListResourceGroupsRequest) Validate() error {
	return dara.Validate(s)
}
