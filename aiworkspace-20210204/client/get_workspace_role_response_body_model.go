// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWorkspaceRoleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCreator(v string) *GetWorkspaceRoleResponseBody
	GetCreator() *string
	SetGmtCreateTime(v string) *GetWorkspaceRoleResponseBody
	GetGmtCreateTime() *string
	SetGmtModifiedTime(v string) *GetWorkspaceRoleResponseBody
	GetGmtModifiedTime() *string
	SetModulePermissions(v []*GetWorkspaceRoleResponseBodyModulePermissions) *GetWorkspaceRoleResponseBody
	GetModulePermissions() []*GetWorkspaceRoleResponseBodyModulePermissions
	SetRequestId(v string) *GetWorkspaceRoleResponseBody
	GetRequestId() *string
	SetRoleId(v string) *GetWorkspaceRoleResponseBody
	GetRoleId() *string
	SetRoleName(v string) *GetWorkspaceRoleResponseBody
	GetRoleName() *string
	SetStatus(v string) *GetWorkspaceRoleResponseBody
	GetStatus() *string
}

type GetWorkspaceRoleResponseBody struct {
	// The Alibaba Cloud account UID of the creator.
	//
	// example:
	//
	// 2680******4103
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// The time when the role was created. It is specified in UTC and formatted in ISO 8601.
	//
	// example:
	//
	// 2025-06-11T08:58:35.438Z
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// The time when the role was last modified. It is specified in UTC and formatted in ISO 8601.
	//
	// example:
	//
	// 2026-03-27T02:26:33.872Z
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// The permission configuration of the role.
	ModulePermissions []*GetWorkspaceRoleResponseBodyModulePermissions `json:"ModulePermissions,omitempty" xml:"ModulePermissions,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// A519F77D-28A0-52F5-AB82-5********8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The role ID.
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
	// The task status. Valid values:
	//
	// - `Creating`: The role is being created.
	//
	// - `Updating`: The role is being updated.
	//
	// - `Deleting`: The role is being deleted.
	//
	// - `Succeeded`: The task succeeded.
	//
	// - `Failed`: The task failed.
	//
	// example:
	//
	// Successful
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetWorkspaceRoleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetWorkspaceRoleResponseBody) GoString() string {
	return s.String()
}

func (s *GetWorkspaceRoleResponseBody) GetCreator() *string {
	return s.Creator
}

func (s *GetWorkspaceRoleResponseBody) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *GetWorkspaceRoleResponseBody) GetGmtModifiedTime() *string {
	return s.GmtModifiedTime
}

func (s *GetWorkspaceRoleResponseBody) GetModulePermissions() []*GetWorkspaceRoleResponseBodyModulePermissions {
	return s.ModulePermissions
}

func (s *GetWorkspaceRoleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetWorkspaceRoleResponseBody) GetRoleId() *string {
	return s.RoleId
}

func (s *GetWorkspaceRoleResponseBody) GetRoleName() *string {
	return s.RoleName
}

func (s *GetWorkspaceRoleResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetWorkspaceRoleResponseBody) SetCreator(v string) *GetWorkspaceRoleResponseBody {
	s.Creator = &v
	return s
}

func (s *GetWorkspaceRoleResponseBody) SetGmtCreateTime(v string) *GetWorkspaceRoleResponseBody {
	s.GmtCreateTime = &v
	return s
}

func (s *GetWorkspaceRoleResponseBody) SetGmtModifiedTime(v string) *GetWorkspaceRoleResponseBody {
	s.GmtModifiedTime = &v
	return s
}

func (s *GetWorkspaceRoleResponseBody) SetModulePermissions(v []*GetWorkspaceRoleResponseBodyModulePermissions) *GetWorkspaceRoleResponseBody {
	s.ModulePermissions = v
	return s
}

func (s *GetWorkspaceRoleResponseBody) SetRequestId(v string) *GetWorkspaceRoleResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetWorkspaceRoleResponseBody) SetRoleId(v string) *GetWorkspaceRoleResponseBody {
	s.RoleId = &v
	return s
}

func (s *GetWorkspaceRoleResponseBody) SetRoleName(v string) *GetWorkspaceRoleResponseBody {
	s.RoleName = &v
	return s
}

func (s *GetWorkspaceRoleResponseBody) SetStatus(v string) *GetWorkspaceRoleResponseBody {
	s.Status = &v
	return s
}

func (s *GetWorkspaceRoleResponseBody) Validate() error {
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

type GetWorkspaceRoleResponseBodyModulePermissions struct {
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
	Permissions []*GetWorkspaceRoleResponseBodyModulePermissionsPermissions `json:"Permissions,omitempty" xml:"Permissions,omitempty" type:"Repeated"`
}

func (s GetWorkspaceRoleResponseBodyModulePermissions) String() string {
	return dara.Prettify(s)
}

func (s GetWorkspaceRoleResponseBodyModulePermissions) GoString() string {
	return s.String()
}

func (s *GetWorkspaceRoleResponseBodyModulePermissions) GetModuleName() *string {
	return s.ModuleName
}

func (s *GetWorkspaceRoleResponseBodyModulePermissions) GetPermissionType() *string {
	return s.PermissionType
}

func (s *GetWorkspaceRoleResponseBodyModulePermissions) GetPermissions() []*GetWorkspaceRoleResponseBodyModulePermissionsPermissions {
	return s.Permissions
}

func (s *GetWorkspaceRoleResponseBodyModulePermissions) SetModuleName(v string) *GetWorkspaceRoleResponseBodyModulePermissions {
	s.ModuleName = &v
	return s
}

func (s *GetWorkspaceRoleResponseBodyModulePermissions) SetPermissionType(v string) *GetWorkspaceRoleResponseBodyModulePermissions {
	s.PermissionType = &v
	return s
}

func (s *GetWorkspaceRoleResponseBodyModulePermissions) SetPermissions(v []*GetWorkspaceRoleResponseBodyModulePermissionsPermissions) *GetWorkspaceRoleResponseBodyModulePermissions {
	s.Permissions = v
	return s
}

func (s *GetWorkspaceRoleResponseBodyModulePermissions) Validate() error {
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

type GetWorkspaceRoleResponseBodyModulePermissionsPermissions struct {
	// A list of permissions.
	PermissionCodes []*string `json:"PermissionCodes,omitempty" xml:"PermissionCodes,omitempty" type:"Repeated"`
	// A list of permission rules.
	PermissionRules []*GetWorkspaceRoleResponseBodyModulePermissionsPermissionsPermissionRules `json:"PermissionRules,omitempty" xml:"PermissionRules,omitempty" type:"Repeated"`
}

func (s GetWorkspaceRoleResponseBodyModulePermissionsPermissions) String() string {
	return dara.Prettify(s)
}

func (s GetWorkspaceRoleResponseBodyModulePermissionsPermissions) GoString() string {
	return s.String()
}

func (s *GetWorkspaceRoleResponseBodyModulePermissionsPermissions) GetPermissionCodes() []*string {
	return s.PermissionCodes
}

func (s *GetWorkspaceRoleResponseBodyModulePermissionsPermissions) GetPermissionRules() []*GetWorkspaceRoleResponseBodyModulePermissionsPermissionsPermissionRules {
	return s.PermissionRules
}

func (s *GetWorkspaceRoleResponseBodyModulePermissionsPermissions) SetPermissionCodes(v []*string) *GetWorkspaceRoleResponseBodyModulePermissionsPermissions {
	s.PermissionCodes = v
	return s
}

func (s *GetWorkspaceRoleResponseBodyModulePermissionsPermissions) SetPermissionRules(v []*GetWorkspaceRoleResponseBodyModulePermissionsPermissionsPermissionRules) *GetWorkspaceRoleResponseBodyModulePermissionsPermissions {
	s.PermissionRules = v
	return s
}

func (s *GetWorkspaceRoleResponseBodyModulePermissionsPermissions) Validate() error {
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

type GetWorkspaceRoleResponseBodyModulePermissionsPermissionsPermissionRules struct {
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

func (s GetWorkspaceRoleResponseBodyModulePermissionsPermissionsPermissionRules) String() string {
	return dara.Prettify(s)
}

func (s GetWorkspaceRoleResponseBodyModulePermissionsPermissionsPermissionRules) GoString() string {
	return s.String()
}

func (s *GetWorkspaceRoleResponseBodyModulePermissionsPermissionsPermissionRules) GetAccessibility() *string {
	return s.Accessibility
}

func (s *GetWorkspaceRoleResponseBodyModulePermissionsPermissionsPermissionRules) GetEntityAccessType() *string {
	return s.EntityAccessType
}

func (s *GetWorkspaceRoleResponseBodyModulePermissionsPermissionsPermissionRules) SetAccessibility(v string) *GetWorkspaceRoleResponseBodyModulePermissionsPermissionsPermissionRules {
	s.Accessibility = &v
	return s
}

func (s *GetWorkspaceRoleResponseBodyModulePermissionsPermissionsPermissionRules) SetEntityAccessType(v string) *GetWorkspaceRoleResponseBodyModulePermissionsPermissionsPermissionRules {
	s.EntityAccessType = &v
	return s
}

func (s *GetWorkspaceRoleResponseBodyModulePermissionsPermissionsPermissionRules) Validate() error {
	return dara.Validate(s)
}
