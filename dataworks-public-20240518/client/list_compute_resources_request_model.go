// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListComputeResourcesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnvType(v string) *ListComputeResourcesRequest
	GetEnvType() *string
	SetName(v string) *ListComputeResourcesRequest
	GetName() *string
	SetOrder(v string) *ListComputeResourcesRequest
	GetOrder() *string
	SetPageNumber(v int32) *ListComputeResourcesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListComputeResourcesRequest
	GetPageSize() *int32
	SetProjectId(v int64) *ListComputeResourcesRequest
	GetProjectId() *int64
	SetSortBy(v string) *ListComputeResourcesRequest
	GetSortBy() *string
	SetTypes(v []*string) *ListComputeResourcesRequest
	GetTypes() []*string
}

type ListComputeResourcesRequest struct {
	// The environment type of the computing resource. Valid values:
	//
	// 	- Dev
	//
	// 	- Prod
	//
	// example:
	//
	// Dev
	EnvType *string `json:"EnvType,omitempty" xml:"EnvType,omitempty"`
	// The name of the computing resource.
	//
	// example:
	//
	// category name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The sort direction of the computing resource list. Valid values:
	//
	// 	- Desc: descending order.
	//
	// 	- Asc: ascending order.
	//
	// Default value: Desc
	//
	// example:
	//
	// Asc
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// The page number to query. The default value is 1, which indicates the first page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. The default value is 10, and the maximum value is 100.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the DataWorks workspace.
	//
	// This parameter is required.
	//
	// example:
	//
	// 21229
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The field to sort the computing resource list by. Supported fields include name, creation time, and computing resource ID.
	//
	// 	- CreateTime: Sorts by creation time
	//
	// 	- Id: Sorts by computing resource ID
	//
	// 	- Name: Sorts by computing resource name.
	//
	// 	- CreateTimeWithDefaultFirst: Sorts based on whether it is the default resource and by creation time, with the default computing resource listed first.
	//
	// Default value: CreateTime
	//
	// example:
	//
	// CreateTimeWithDefaultFirst
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// The filter for computing resource types. You can configure multiple types for filtering.
	Types []*string `json:"Types,omitempty" xml:"Types,omitempty" type:"Repeated"`
}

func (s ListComputeResourcesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListComputeResourcesRequest) GoString() string {
	return s.String()
}

func (s *ListComputeResourcesRequest) GetEnvType() *string {
	return s.EnvType
}

func (s *ListComputeResourcesRequest) GetName() *string {
	return s.Name
}

func (s *ListComputeResourcesRequest) GetOrder() *string {
	return s.Order
}

func (s *ListComputeResourcesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListComputeResourcesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListComputeResourcesRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListComputeResourcesRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListComputeResourcesRequest) GetTypes() []*string {
	return s.Types
}

func (s *ListComputeResourcesRequest) SetEnvType(v string) *ListComputeResourcesRequest {
	s.EnvType = &v
	return s
}

func (s *ListComputeResourcesRequest) SetName(v string) *ListComputeResourcesRequest {
	s.Name = &v
	return s
}

func (s *ListComputeResourcesRequest) SetOrder(v string) *ListComputeResourcesRequest {
	s.Order = &v
	return s
}

func (s *ListComputeResourcesRequest) SetPageNumber(v int32) *ListComputeResourcesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListComputeResourcesRequest) SetPageSize(v int32) *ListComputeResourcesRequest {
	s.PageSize = &v
	return s
}

func (s *ListComputeResourcesRequest) SetProjectId(v int64) *ListComputeResourcesRequest {
	s.ProjectId = &v
	return s
}

func (s *ListComputeResourcesRequest) SetSortBy(v string) *ListComputeResourcesRequest {
	s.SortBy = &v
	return s
}

func (s *ListComputeResourcesRequest) SetTypes(v []*string) *ListComputeResourcesRequest {
	s.Types = v
	return s
}

func (s *ListComputeResourcesRequest) Validate() error {
	return dara.Validate(s)
}
