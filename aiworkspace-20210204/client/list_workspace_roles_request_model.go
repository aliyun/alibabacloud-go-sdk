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
	// The sort order for the specified sort field. Valid values:
	//
	// - `ASC` (default): Ascending order.
	//
	// - `DESC`: Descending order.
	//
	// example:
	//
	// DESC
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// The page number. Pages start from 1. The default value is 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. The default value is 20.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// A comma-separated list of role IDs.
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
	// The role type.
	//
	// - `custom`: A custom role.
	//
	// example:
	//
	// custom
	RoleType *string `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
	// The sort field for paginated queries. Valid values:
	//
	// - `GmtCreateTime` (default): Sorts by creation time.
	//
	// - `GmtModifiedTime`: Sorts by modification time.
	//
	// example:
	//
	// GmtCreateTime
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// The status. Valid values:
	//
	// - `Creating`
	//
	// - `Updating`
	//
	// - `Deleting`
	//
	// - `Succeeded`: A terminal state.
	//
	// - `Failed`: A terminal state.
	//
	// example:
	//
	// Succeeded
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// A comma-separated list of fields to return. Currently, only `ModulePermissions` is supported, which returns the permission configuration of the role.
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
