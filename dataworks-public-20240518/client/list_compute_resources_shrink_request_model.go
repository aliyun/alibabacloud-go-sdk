// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListComputeResourcesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnvType(v string) *ListComputeResourcesShrinkRequest
	GetEnvType() *string
	SetName(v string) *ListComputeResourcesShrinkRequest
	GetName() *string
	SetOrder(v string) *ListComputeResourcesShrinkRequest
	GetOrder() *string
	SetPageNumber(v int32) *ListComputeResourcesShrinkRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListComputeResourcesShrinkRequest
	GetPageSize() *int32
	SetProjectId(v int64) *ListComputeResourcesShrinkRequest
	GetProjectId() *int64
	SetSortBy(v string) *ListComputeResourcesShrinkRequest
	GetSortBy() *string
	SetTypesShrink(v string) *ListComputeResourcesShrinkRequest
	GetTypesShrink() *string
}

type ListComputeResourcesShrinkRequest struct {
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
	TypesShrink *string `json:"Types,omitempty" xml:"Types,omitempty"`
}

func (s ListComputeResourcesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListComputeResourcesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListComputeResourcesShrinkRequest) GetEnvType() *string {
	return s.EnvType
}

func (s *ListComputeResourcesShrinkRequest) GetName() *string {
	return s.Name
}

func (s *ListComputeResourcesShrinkRequest) GetOrder() *string {
	return s.Order
}

func (s *ListComputeResourcesShrinkRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListComputeResourcesShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListComputeResourcesShrinkRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListComputeResourcesShrinkRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListComputeResourcesShrinkRequest) GetTypesShrink() *string {
	return s.TypesShrink
}

func (s *ListComputeResourcesShrinkRequest) SetEnvType(v string) *ListComputeResourcesShrinkRequest {
	s.EnvType = &v
	return s
}

func (s *ListComputeResourcesShrinkRequest) SetName(v string) *ListComputeResourcesShrinkRequest {
	s.Name = &v
	return s
}

func (s *ListComputeResourcesShrinkRequest) SetOrder(v string) *ListComputeResourcesShrinkRequest {
	s.Order = &v
	return s
}

func (s *ListComputeResourcesShrinkRequest) SetPageNumber(v int32) *ListComputeResourcesShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListComputeResourcesShrinkRequest) SetPageSize(v int32) *ListComputeResourcesShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListComputeResourcesShrinkRequest) SetProjectId(v int64) *ListComputeResourcesShrinkRequest {
	s.ProjectId = &v
	return s
}

func (s *ListComputeResourcesShrinkRequest) SetSortBy(v string) *ListComputeResourcesShrinkRequest {
	s.SortBy = &v
	return s
}

func (s *ListComputeResourcesShrinkRequest) SetTypesShrink(v string) *ListComputeResourcesShrinkRequest {
	s.TypesShrink = &v
	return s
}

func (s *ListComputeResourcesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
