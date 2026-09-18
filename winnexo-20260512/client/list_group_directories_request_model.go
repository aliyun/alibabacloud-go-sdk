// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGroupDirectoriesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDirectoryId(v string) *ListGroupDirectoriesRequest
	GetDirectoryId() *string
	SetGroupId(v string) *ListGroupDirectoriesRequest
	GetGroupId() *string
	SetSortField(v string) *ListGroupDirectoriesRequest
	GetSortField() *string
	SetSortOrder(v string) *ListGroupDirectoriesRequest
	GetSortOrder() *string
	SetTenantId(v string) *ListGroupDirectoriesRequest
	GetTenantId() *string
}

type ListGroupDirectoriesRequest struct {
	// The ID of a visible directory in the current space. If this parameter is omitted or set to root, the space root is queried. The first query reuses the existing service-initialized internal root directory.
	//
	// example:
	//
	// dir_example
	DirectoryId *string `json:"directoryId,omitempty" xml:"directoryId,omitempty"`
	// The ID of the collaborative share.
	//
	// This parameter is required.
	//
	// example:
	//
	// group_example
	GroupId *string `json:"groupId,omitempty" xml:"groupId,omitempty"`
	// The sort field. Valid values: name, gmt_create, and gmt_modified.
	//
	// example:
	//
	// gmt_create
	SortField *string `json:"sortField,omitempty" xml:"sortField,omitempty"`
	// The sort order. Valid values: asc and desc.
	//
	// example:
	//
	// desc
	SortOrder *string `json:"sortOrder,omitempty" xml:"sortOrder,omitempty"`
	// The tenant ID. This is a common parameter. If not specified, the default tenant of the caller is used.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s ListGroupDirectoriesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListGroupDirectoriesRequest) GoString() string {
	return s.String()
}

func (s *ListGroupDirectoriesRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *ListGroupDirectoriesRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *ListGroupDirectoriesRequest) GetSortField() *string {
	return s.SortField
}

func (s *ListGroupDirectoriesRequest) GetSortOrder() *string {
	return s.SortOrder
}

func (s *ListGroupDirectoriesRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ListGroupDirectoriesRequest) SetDirectoryId(v string) *ListGroupDirectoriesRequest {
	s.DirectoryId = &v
	return s
}

func (s *ListGroupDirectoriesRequest) SetGroupId(v string) *ListGroupDirectoriesRequest {
	s.GroupId = &v
	return s
}

func (s *ListGroupDirectoriesRequest) SetSortField(v string) *ListGroupDirectoriesRequest {
	s.SortField = &v
	return s
}

func (s *ListGroupDirectoriesRequest) SetSortOrder(v string) *ListGroupDirectoriesRequest {
	s.SortOrder = &v
	return s
}

func (s *ListGroupDirectoriesRequest) SetTenantId(v string) *ListGroupDirectoriesRequest {
	s.TenantId = &v
	return s
}

func (s *ListGroupDirectoriesRequest) Validate() error {
	return dara.Validate(s)
}
