// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProjectsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *ListProjectsRequest
	GetName() *string
	SetOrder(v string) *ListProjectsRequest
	GetOrder() *string
	SetOwner(v string) *ListProjectsRequest
	GetOwner() *string
	SetPageNumber(v int32) *ListProjectsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListProjectsRequest
	GetPageSize() *int32
	SetProjectIds(v []*string) *ListProjectsRequest
	GetProjectIds() []*string
	SetSortBy(v string) *ListProjectsRequest
	GetSortBy() *string
	SetWorkspaceId(v string) *ListProjectsRequest
	GetWorkspaceId() *string
}

type ListProjectsRequest struct {
	// The project name to filter by.
	//
	// example:
	//
	// fs1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The sort order. Valid values:
	//
	// - `Asc`: ascending order.
	//
	// - `Desc`: descending order.
	//
	// example:
	//
	// Asc
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// The Alibaba Cloud account ID of the project owner.
	//
	// example:
	//
	// 134324352****
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The page number. Starts from 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// An array of project IDs to filter by.
	ProjectIds []*string `json:"ProjectIds,omitempty" xml:"ProjectIds,omitempty" type:"Repeated"`
	// The field to sort the results by.
	//
	// example:
	//
	// GmtModifiedTime
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// The ID of the workspace. Call the [ListWorkspaces](https://help.aliyun.com/document_detail/449124.html) operation to get this ID.
	//
	// example:
	//
	// 234
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListProjectsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListProjectsRequest) GoString() string {
	return s.String()
}

func (s *ListProjectsRequest) GetName() *string {
	return s.Name
}

func (s *ListProjectsRequest) GetOrder() *string {
	return s.Order
}

func (s *ListProjectsRequest) GetOwner() *string {
	return s.Owner
}

func (s *ListProjectsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListProjectsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListProjectsRequest) GetProjectIds() []*string {
	return s.ProjectIds
}

func (s *ListProjectsRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListProjectsRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListProjectsRequest) SetName(v string) *ListProjectsRequest {
	s.Name = &v
	return s
}

func (s *ListProjectsRequest) SetOrder(v string) *ListProjectsRequest {
	s.Order = &v
	return s
}

func (s *ListProjectsRequest) SetOwner(v string) *ListProjectsRequest {
	s.Owner = &v
	return s
}

func (s *ListProjectsRequest) SetPageNumber(v int32) *ListProjectsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListProjectsRequest) SetPageSize(v int32) *ListProjectsRequest {
	s.PageSize = &v
	return s
}

func (s *ListProjectsRequest) SetProjectIds(v []*string) *ListProjectsRequest {
	s.ProjectIds = v
	return s
}

func (s *ListProjectsRequest) SetSortBy(v string) *ListProjectsRequest {
	s.SortBy = &v
	return s
}

func (s *ListProjectsRequest) SetWorkspaceId(v string) *ListProjectsRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ListProjectsRequest) Validate() error {
	return dara.Validate(s)
}
