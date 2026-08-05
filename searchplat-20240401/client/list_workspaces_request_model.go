// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWorkspacesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListWorkspacesRequest
	GetInstanceId() *string
	SetName(v string) *ListWorkspacesRequest
	GetName() *string
	SetPageNumber(v int32) *ListWorkspacesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListWorkspacesRequest
	GetPageSize() *int32
	SetResourceGroupId(v string) *ListWorkspacesRequest
	GetResourceGroupId() *string
	SetSortBy(v int32) *ListWorkspacesRequest
	GetSortBy() *int32
	SetType(v string) *ListWorkspacesRequest
	GetType() *string
}

type ListWorkspacesRequest struct {
	// The instance ID.
	//
	// example:
	//
	// ops-cn-jte49bevd04
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// The workspace name.
	//
	// example:
	//
	// default
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-aekzivjfrlpyn3y
	ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	// The sort order. Valid values:
	//
	// - 0: sorted by creation time in descending order.
	//
	// - 1: sorted by modification time in descending order.
	//
	// Default value: 0.
	//
	// example:
	//
	// 1
	SortBy *int32 `json:"sortBy,omitempty" xml:"sortBy,omitempty"`
	// The type.
	//
	// example:
	//
	// standard
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListWorkspacesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListWorkspacesRequest) GoString() string {
	return s.String()
}

func (s *ListWorkspacesRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListWorkspacesRequest) GetName() *string {
	return s.Name
}

func (s *ListWorkspacesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListWorkspacesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListWorkspacesRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListWorkspacesRequest) GetSortBy() *int32 {
	return s.SortBy
}

func (s *ListWorkspacesRequest) GetType() *string {
	return s.Type
}

func (s *ListWorkspacesRequest) SetInstanceId(v string) *ListWorkspacesRequest {
	s.InstanceId = &v
	return s
}

func (s *ListWorkspacesRequest) SetName(v string) *ListWorkspacesRequest {
	s.Name = &v
	return s
}

func (s *ListWorkspacesRequest) SetPageNumber(v int32) *ListWorkspacesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListWorkspacesRequest) SetPageSize(v int32) *ListWorkspacesRequest {
	s.PageSize = &v
	return s
}

func (s *ListWorkspacesRequest) SetResourceGroupId(v string) *ListWorkspacesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListWorkspacesRequest) SetSortBy(v int32) *ListWorkspacesRequest {
	s.SortBy = &v
	return s
}

func (s *ListWorkspacesRequest) SetType(v string) *ListWorkspacesRequest {
	s.Type = &v
	return s
}

func (s *ListWorkspacesRequest) Validate() error {
	return dara.Validate(s)
}
