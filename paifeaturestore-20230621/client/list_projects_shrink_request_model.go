// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProjectsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *ListProjectsShrinkRequest
	GetName() *string
	SetOrder(v string) *ListProjectsShrinkRequest
	GetOrder() *string
	SetOwner(v string) *ListProjectsShrinkRequest
	GetOwner() *string
	SetPageNumber(v int32) *ListProjectsShrinkRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListProjectsShrinkRequest
	GetPageSize() *int32
	SetProjectIdsShrink(v string) *ListProjectsShrinkRequest
	GetProjectIdsShrink() *string
	SetSortBy(v string) *ListProjectsShrinkRequest
	GetSortBy() *string
	SetWorkspaceId(v string) *ListProjectsShrinkRequest
	GetWorkspaceId() *string
}

type ListProjectsShrinkRequest struct {
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
	ProjectIdsShrink *string `json:"ProjectIds,omitempty" xml:"ProjectIds,omitempty"`
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

func (s ListProjectsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListProjectsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListProjectsShrinkRequest) GetName() *string {
	return s.Name
}

func (s *ListProjectsShrinkRequest) GetOrder() *string {
	return s.Order
}

func (s *ListProjectsShrinkRequest) GetOwner() *string {
	return s.Owner
}

func (s *ListProjectsShrinkRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListProjectsShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListProjectsShrinkRequest) GetProjectIdsShrink() *string {
	return s.ProjectIdsShrink
}

func (s *ListProjectsShrinkRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListProjectsShrinkRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListProjectsShrinkRequest) SetName(v string) *ListProjectsShrinkRequest {
	s.Name = &v
	return s
}

func (s *ListProjectsShrinkRequest) SetOrder(v string) *ListProjectsShrinkRequest {
	s.Order = &v
	return s
}

func (s *ListProjectsShrinkRequest) SetOwner(v string) *ListProjectsShrinkRequest {
	s.Owner = &v
	return s
}

func (s *ListProjectsShrinkRequest) SetPageNumber(v int32) *ListProjectsShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListProjectsShrinkRequest) SetPageSize(v int32) *ListProjectsShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListProjectsShrinkRequest) SetProjectIdsShrink(v string) *ListProjectsShrinkRequest {
	s.ProjectIdsShrink = &v
	return s
}

func (s *ListProjectsShrinkRequest) SetSortBy(v string) *ListProjectsShrinkRequest {
	s.SortBy = &v
	return s
}

func (s *ListProjectsShrinkRequest) SetWorkspaceId(v string) *ListProjectsShrinkRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ListProjectsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
