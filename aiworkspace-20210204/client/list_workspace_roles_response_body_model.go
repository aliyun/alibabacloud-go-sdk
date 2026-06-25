// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWorkspaceRolesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListWorkspaceRolesResponseBody
	GetRequestId() *string
	SetRoles(v []*ListWorkspaceRolesResponseBodyRoles) *ListWorkspaceRolesResponseBody
	GetRoles() []*ListWorkspaceRolesResponseBodyRoles
	SetTotalCount(v int64) *ListWorkspaceRolesResponseBody
	GetTotalCount() *int64
}

type ListWorkspaceRolesResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// A519F77D-28A0-52F5-AB82-5********8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of custom roles.
	Roles []*ListWorkspaceRolesResponseBodyRoles `json:"Roles,omitempty" xml:"Roles,omitempty" type:"Repeated"`
	// The total count of matching entries.
	//
	// example:
	//
	// 15
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListWorkspaceRolesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListWorkspaceRolesResponseBody) GoString() string {
	return s.String()
}

func (s *ListWorkspaceRolesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListWorkspaceRolesResponseBody) GetRoles() []*ListWorkspaceRolesResponseBodyRoles {
	return s.Roles
}

func (s *ListWorkspaceRolesResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListWorkspaceRolesResponseBody) SetRequestId(v string) *ListWorkspaceRolesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListWorkspaceRolesResponseBody) SetRoles(v []*ListWorkspaceRolesResponseBodyRoles) *ListWorkspaceRolesResponseBody {
	s.Roles = v
	return s
}

func (s *ListWorkspaceRolesResponseBody) SetTotalCount(v int64) *ListWorkspaceRolesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListWorkspaceRolesResponseBody) Validate() error {
	if s.Roles != nil {
		for _, item := range s.Roles {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListWorkspaceRolesResponseBodyRoles struct {
	// The Alibaba Cloud account UID of the creator.
	//
	// example:
	//
	// 2680******4103
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// The creation time, in UTC and ISO 8601 format.
	//
	// example:
	//
	// 2025-09-14T07:40:01.000Z
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// The modification time, in UTC and ISO 8601 format.
	//
	// example:
	//
	// 2026-04-15T02:29:52Z
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// The permission configuration of the role.
	ModulePermissions []*ListWorkspaceRolesResponseBodyRolesModulePermissions `json:"ModulePermissions,omitempty" xml:"ModulePermissions,omitempty" type:"Repeated"`
	// The ID of the custom role.
	//
	// example:
	//
	// role-dhg*******
	RoleId *string `json:"RoleId,omitempty" xml:"RoleId,omitempty"`
	// The role name.
	//
	// example:
	//
	// dev-test
	RoleName *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
}

func (s ListWorkspaceRolesResponseBodyRoles) String() string {
	return dara.Prettify(s)
}

func (s ListWorkspaceRolesResponseBodyRoles) GoString() string {
	return s.String()
}

func (s *ListWorkspaceRolesResponseBodyRoles) GetCreator() *string {
	return s.Creator
}

func (s *ListWorkspaceRolesResponseBodyRoles) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *ListWorkspaceRolesResponseBodyRoles) GetGmtModifiedTime() *string {
	return s.GmtModifiedTime
}

func (s *ListWorkspaceRolesResponseBodyRoles) GetModulePermissions() []*ListWorkspaceRolesResponseBodyRolesModulePermissions {
	return s.ModulePermissions
}

func (s *ListWorkspaceRolesResponseBodyRoles) GetRoleId() *string {
	return s.RoleId
}

func (s *ListWorkspaceRolesResponseBodyRoles) GetRoleName() *string {
	return s.RoleName
}

func (s *ListWorkspaceRolesResponseBodyRoles) SetCreator(v string) *ListWorkspaceRolesResponseBodyRoles {
	s.Creator = &v
	return s
}

func (s *ListWorkspaceRolesResponseBodyRoles) SetGmtCreateTime(v string) *ListWorkspaceRolesResponseBodyRoles {
	s.GmtCreateTime = &v
	return s
}

func (s *ListWorkspaceRolesResponseBodyRoles) SetGmtModifiedTime(v string) *ListWorkspaceRolesResponseBodyRoles {
	s.GmtModifiedTime = &v
	return s
}

func (s *ListWorkspaceRolesResponseBodyRoles) SetModulePermissions(v []*ListWorkspaceRolesResponseBodyRolesModulePermissions) *ListWorkspaceRolesResponseBodyRoles {
	s.ModulePermissions = v
	return s
}

func (s *ListWorkspaceRolesResponseBodyRoles) SetRoleId(v string) *ListWorkspaceRolesResponseBodyRoles {
	s.RoleId = &v
	return s
}

func (s *ListWorkspaceRolesResponseBodyRoles) SetRoleName(v string) *ListWorkspaceRolesResponseBodyRoles {
	s.RoleName = &v
	return s
}

func (s *ListWorkspaceRolesResponseBodyRoles) Validate() error {
	if s.ModulePermissions != nil {
		for _, item := range s.ModulePermissions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListWorkspaceRolesResponseBodyRolesModulePermissions struct {
	// The module name. Valid values are:
	//
	// - PaiDesigner: PAI-Designer
	//
	// - Paiflow: workflow
	//
	// - DSW: PAI-DSW
	//
	// - DLC: PAI-DLC
	//
	// - Dataset: dataset
	//
	// - Model: model
	//
	// - Image: image
	//
	// - CodeSource: code source
	//
	// - PaiWorkspace@@Workspace: Basic workspace information
	//
	// - PaiWorkspace@@MemberRole: workspace member management
	//
	// - PaiWorkspace@@Resource: workspace computing resource management
	//
	// - PaiWorkspace@@Config: workspace configuration center
	//
	// - ArtLab: ArtLab
	//
	// example:
	//
	// PaiDesigner
	ModuleName *string `json:"ModuleName,omitempty" xml:"ModuleName,omitempty"`
	// The permission type. Valid values are:
	//
	// - ReadOnly: read-only access.
	//
	// - ReadWrite: Allows users to edit and run.
	//
	// - FullAccess: full control.
	//
	// - NoPrivilege: no permissions.
	//
	// - CustomPermissions: custom permissions.
	//
	// example:
	//
	// ReadOnly
	PermissionType *string `json:"PermissionType,omitempty" xml:"PermissionType,omitempty"`
	// The permissions. This parameter is required and applies only when `PermissionType` is set to `CustomPermissions`.
	Permissions []*ListWorkspaceRolesResponseBodyRolesModulePermissionsPermissions `json:"Permissions,omitempty" xml:"Permissions,omitempty" type:"Repeated"`
}

func (s ListWorkspaceRolesResponseBodyRolesModulePermissions) String() string {
	return dara.Prettify(s)
}

func (s ListWorkspaceRolesResponseBodyRolesModulePermissions) GoString() string {
	return s.String()
}

func (s *ListWorkspaceRolesResponseBodyRolesModulePermissions) GetModuleName() *string {
	return s.ModuleName
}

func (s *ListWorkspaceRolesResponseBodyRolesModulePermissions) GetPermissionType() *string {
	return s.PermissionType
}

func (s *ListWorkspaceRolesResponseBodyRolesModulePermissions) GetPermissions() []*ListWorkspaceRolesResponseBodyRolesModulePermissionsPermissions {
	return s.Permissions
}

func (s *ListWorkspaceRolesResponseBodyRolesModulePermissions) SetModuleName(v string) *ListWorkspaceRolesResponseBodyRolesModulePermissions {
	s.ModuleName = &v
	return s
}

func (s *ListWorkspaceRolesResponseBodyRolesModulePermissions) SetPermissionType(v string) *ListWorkspaceRolesResponseBodyRolesModulePermissions {
	s.PermissionType = &v
	return s
}

func (s *ListWorkspaceRolesResponseBodyRolesModulePermissions) SetPermissions(v []*ListWorkspaceRolesResponseBodyRolesModulePermissionsPermissions) *ListWorkspaceRolesResponseBodyRolesModulePermissions {
	s.Permissions = v
	return s
}

func (s *ListWorkspaceRolesResponseBodyRolesModulePermissions) Validate() error {
	if s.Permissions != nil {
		for _, item := range s.Permissions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListWorkspaceRolesResponseBodyRolesModulePermissionsPermissions struct {
	// A list of permissions.
	PermissionCodes []*string `json:"PermissionCodes,omitempty" xml:"PermissionCodes,omitempty" type:"Repeated"`
	// A list of permission rules.
	PermissionRules []*ListWorkspaceRolesResponseBodyRolesModulePermissionsPermissionsPermissionRules `json:"PermissionRules,omitempty" xml:"PermissionRules,omitempty" type:"Repeated"`
}

func (s ListWorkspaceRolesResponseBodyRolesModulePermissionsPermissions) String() string {
	return dara.Prettify(s)
}

func (s ListWorkspaceRolesResponseBodyRolesModulePermissionsPermissions) GoString() string {
	return s.String()
}

func (s *ListWorkspaceRolesResponseBodyRolesModulePermissionsPermissions) GetPermissionCodes() []*string {
	return s.PermissionCodes
}

func (s *ListWorkspaceRolesResponseBodyRolesModulePermissionsPermissions) GetPermissionRules() []*ListWorkspaceRolesResponseBodyRolesModulePermissionsPermissionsPermissionRules {
	return s.PermissionRules
}

func (s *ListWorkspaceRolesResponseBodyRolesModulePermissionsPermissions) SetPermissionCodes(v []*string) *ListWorkspaceRolesResponseBodyRolesModulePermissionsPermissions {
	s.PermissionCodes = v
	return s
}

func (s *ListWorkspaceRolesResponseBodyRolesModulePermissionsPermissions) SetPermissionRules(v []*ListWorkspaceRolesResponseBodyRolesModulePermissionsPermissionsPermissionRules) *ListWorkspaceRolesResponseBodyRolesModulePermissionsPermissions {
	s.PermissionRules = v
	return s
}

func (s *ListWorkspaceRolesResponseBodyRolesModulePermissionsPermissions) Validate() error {
	if s.PermissionRules != nil {
		for _, item := range s.PermissionRules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListWorkspaceRolesResponseBodyRolesModulePermissionsPermissionsPermissionRules struct {
	// The access type. Valid values are:
	//
	// - PUBLIC: All members in the current workspace can perform this action.
	//
	// - PRIVATE: Only the creator can perform this action.
	//
	// - ANY: Both the creator and non-creators can perform this action.
	//
	// example:
	//
	// PRIVATE
	Accessibility *string `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	// The entity access type. This parameter is ignored if `Accessibility` is set to `PUBLIC`. If `Accessibility` is set to `PRIVATE`, the value of this parameter determines the permissions. Valid values are:
	//
	// - CREATOR: Only the creator can perform this action.
	//
	// - ANY: Both the creator and non-creators can perform this action.
	//
	// example:
	//
	// CREATOR
	EntityAccessType *string `json:"EntityAccessType,omitempty" xml:"EntityAccessType,omitempty"`
}

func (s ListWorkspaceRolesResponseBodyRolesModulePermissionsPermissionsPermissionRules) String() string {
	return dara.Prettify(s)
}

func (s ListWorkspaceRolesResponseBodyRolesModulePermissionsPermissionsPermissionRules) GoString() string {
	return s.String()
}

func (s *ListWorkspaceRolesResponseBodyRolesModulePermissionsPermissionsPermissionRules) GetAccessibility() *string {
	return s.Accessibility
}

func (s *ListWorkspaceRolesResponseBodyRolesModulePermissionsPermissionsPermissionRules) GetEntityAccessType() *string {
	return s.EntityAccessType
}

func (s *ListWorkspaceRolesResponseBodyRolesModulePermissionsPermissionsPermissionRules) SetAccessibility(v string) *ListWorkspaceRolesResponseBodyRolesModulePermissionsPermissionsPermissionRules {
	s.Accessibility = &v
	return s
}

func (s *ListWorkspaceRolesResponseBodyRolesModulePermissionsPermissionsPermissionRules) SetEntityAccessType(v string) *ListWorkspaceRolesResponseBodyRolesModulePermissionsPermissionsPermissionRules {
	s.EntityAccessType = &v
	return s
}

func (s *ListWorkspaceRolesResponseBodyRolesModulePermissionsPermissionsPermissionRules) Validate() error {
	return dara.Validate(s)
}
