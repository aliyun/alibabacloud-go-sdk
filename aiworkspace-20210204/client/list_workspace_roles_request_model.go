// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWorkspaceRolesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOrder(v string) *ListWorkspaceRolesRequest
	GetOrder() *string
	SetPageNumber(v int32) *ListWorkspaceRolesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListWorkspaceRolesRequest
	GetPageSize() *int32
	SetRoleIds(v string) *ListWorkspaceRolesRequest
	GetRoleIds() *string
	SetRoleName(v string) *ListWorkspaceRolesRequest
	GetRoleName() *string
	SetRoleType(v string) *ListWorkspaceRolesRequest
	GetRoleType() *string
	SetSortBy(v string) *ListWorkspaceRolesRequest
	GetSortBy() *string
	SetStatus(v string) *ListWorkspaceRolesRequest
	GetStatus() *string
	SetVerboseFields(v string) *ListWorkspaceRolesRequest
	GetVerboseFields() *string
}

type ListWorkspaceRolesRequest struct {
	// The sorting order for the specified sorting field in paging. Valid values:
	//
	// - ASC: Ascending order. Default value.
	//
	// - DESC: Descending order.
	//
	// example:
	//
	// DESC
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// The page number. Pages start from 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 20.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The list of role IDs. Separate multiple IDs with commas (,).
	//
	// example:
	//
	// role-dhg*******
	RoleIds *string `json:"RoleIds,omitempty" xml:"RoleIds,omitempty"`
	// The role name.
	//
	// example:
	//
	// dev-test
	RoleName *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
	// The role type. Valid values:
	//
	// - custom: custom role.
	//
	// example:
	//
	// custom
	RoleType *string `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
	// The sorting field for paging. Valid values:
	//
	// - GmtCreateTime: Sort by creation time. Default value.
	//
	// - GmtModifiedTime: Sort by modification time.
	//
	// example:
	//
	// GmtCreateTime
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// **[Deprecated]*	- The status.
	//
	// example:
	//
	// Succeeded
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The comma-separated list of fields to return. Currently, only ModulePermissions is supported, which indicates whether to return the permission configurations of the role.
	//
	// example:
	//
	// ModulePermissions
	VerboseFields *string `json:"VerboseFields,omitempty" xml:"VerboseFields,omitempty"`
}

func (s ListWorkspaceRolesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListWorkspaceRolesRequest) GoString() string {
	return s.String()
}

func (s *ListWorkspaceRolesRequest) GetOrder() *string {
	return s.Order
}

func (s *ListWorkspaceRolesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListWorkspaceRolesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListWorkspaceRolesRequest) GetRoleIds() *string {
	return s.RoleIds
}

func (s *ListWorkspaceRolesRequest) GetRoleName() *string {
	return s.RoleName
}

func (s *ListWorkspaceRolesRequest) GetRoleType() *string {
	return s.RoleType
}

func (s *ListWorkspaceRolesRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListWorkspaceRolesRequest) GetStatus() *string {
	return s.Status
}

func (s *ListWorkspaceRolesRequest) GetVerboseFields() *string {
	return s.VerboseFields
}

func (s *ListWorkspaceRolesRequest) SetOrder(v string) *ListWorkspaceRolesRequest {
	s.Order = &v
	return s
}

func (s *ListWorkspaceRolesRequest) SetPageNumber(v int32) *ListWorkspaceRolesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListWorkspaceRolesRequest) SetPageSize(v int32) *ListWorkspaceRolesRequest {
	s.PageSize = &v
	return s
}

func (s *ListWorkspaceRolesRequest) SetRoleIds(v string) *ListWorkspaceRolesRequest {
	s.RoleIds = &v
	return s
}

func (s *ListWorkspaceRolesRequest) SetRoleName(v string) *ListWorkspaceRolesRequest {
	s.RoleName = &v
	return s
}

func (s *ListWorkspaceRolesRequest) SetRoleType(v string) *ListWorkspaceRolesRequest {
	s.RoleType = &v
	return s
}

func (s *ListWorkspaceRolesRequest) SetSortBy(v string) *ListWorkspaceRolesRequest {
	s.SortBy = &v
	return s
}

func (s *ListWorkspaceRolesRequest) SetStatus(v string) *ListWorkspaceRolesRequest {
	s.Status = &v
	return s
}

func (s *ListWorkspaceRolesRequest) SetVerboseFields(v string) *ListWorkspaceRolesRequest {
	s.VerboseFields = &v
	return s
}

func (s *ListWorkspaceRolesRequest) Validate() error {
	return dara.Validate(s)
}
