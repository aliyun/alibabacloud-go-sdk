// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWorkspaceRoleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetModulePermissions(v []*CreateWorkspaceRoleRequestModulePermissions) *CreateWorkspaceRoleRequest
	GetModulePermissions() []*CreateWorkspaceRoleRequestModulePermissions
	SetRoleName(v string) *CreateWorkspaceRoleRequest
	GetRoleName() *string
	SetRoleType(v string) *CreateWorkspaceRoleRequest
	GetRoleType() *string
	SetUserId(v string) *CreateWorkspaceRoleRequest
	GetUserId() *string
}

type CreateWorkspaceRoleRequest struct {
	// The permission settings for the role.
	ModulePermissions []*CreateWorkspaceRoleRequestModulePermissions `json:"ModulePermissions,omitempty" xml:"ModulePermissions,omitempty" type:"Repeated"`
	// The unique name for the custom role within the workspace. It can contain letters, digits, underscores (_), and hyphens (-), and be up to 64 characters long.
	//
	// example:
	//
	// dev-test
	RoleName *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
	// The role type. This operation only creates custom roles. Valid value: custom.
	//
	// example:
	//
	// custom
	RoleType *string `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
	UserId   *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s CreateWorkspaceRoleRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateWorkspaceRoleRequest) GoString() string {
	return s.String()
}

func (s *CreateWorkspaceRoleRequest) GetModulePermissions() []*CreateWorkspaceRoleRequestModulePermissions {
	return s.ModulePermissions
}

func (s *CreateWorkspaceRoleRequest) GetRoleName() *string {
	return s.RoleName
}

func (s *CreateWorkspaceRoleRequest) GetRoleType() *string {
	return s.RoleType
}

func (s *CreateWorkspaceRoleRequest) GetUserId() *string {
	return s.UserId
}

func (s *CreateWorkspaceRoleRequest) SetModulePermissions(v []*CreateWorkspaceRoleRequestModulePermissions) *CreateWorkspaceRoleRequest {
	s.ModulePermissions = v
	return s
}

func (s *CreateWorkspaceRoleRequest) SetRoleName(v string) *CreateWorkspaceRoleRequest {
	s.RoleName = &v
	return s
}

func (s *CreateWorkspaceRoleRequest) SetRoleType(v string) *CreateWorkspaceRoleRequest {
	s.RoleType = &v
	return s
}

func (s *CreateWorkspaceRoleRequest) SetUserId(v string) *CreateWorkspaceRoleRequest {
	s.UserId = &v
	return s
}

func (s *CreateWorkspaceRoleRequest) Validate() error {
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

type CreateWorkspaceRoleRequestModulePermissions struct {
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
	Permissions []*CreateWorkspaceRoleRequestModulePermissionsPermissions `json:"Permissions,omitempty" xml:"Permissions,omitempty" type:"Repeated"`
}

func (s CreateWorkspaceRoleRequestModulePermissions) String() string {
	return dara.Prettify(s)
}

func (s CreateWorkspaceRoleRequestModulePermissions) GoString() string {
	return s.String()
}

func (s *CreateWorkspaceRoleRequestModulePermissions) GetModuleName() *string {
	return s.ModuleName
}

func (s *CreateWorkspaceRoleRequestModulePermissions) GetPermissionType() *string {
	return s.PermissionType
}

func (s *CreateWorkspaceRoleRequestModulePermissions) GetPermissions() []*CreateWorkspaceRoleRequestModulePermissionsPermissions {
	return s.Permissions
}

func (s *CreateWorkspaceRoleRequestModulePermissions) SetModuleName(v string) *CreateWorkspaceRoleRequestModulePermissions {
	s.ModuleName = &v
	return s
}

func (s *CreateWorkspaceRoleRequestModulePermissions) SetPermissionType(v string) *CreateWorkspaceRoleRequestModulePermissions {
	s.PermissionType = &v
	return s
}

func (s *CreateWorkspaceRoleRequestModulePermissions) SetPermissions(v []*CreateWorkspaceRoleRequestModulePermissionsPermissions) *CreateWorkspaceRoleRequestModulePermissions {
	s.Permissions = v
	return s
}

func (s *CreateWorkspaceRoleRequestModulePermissions) Validate() error {
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

type CreateWorkspaceRoleRequestModulePermissionsPermissions struct {
	// A list of permissions.
	PermissionCodes []*string `json:"PermissionCodes,omitempty" xml:"PermissionCodes,omitempty" type:"Repeated"`
	// A list of permission rules.
	PermissionRules []*CreateWorkspaceRoleRequestModulePermissionsPermissionsPermissionRules `json:"PermissionRules,omitempty" xml:"PermissionRules,omitempty" type:"Repeated"`
}

func (s CreateWorkspaceRoleRequestModulePermissionsPermissions) String() string {
	return dara.Prettify(s)
}

func (s CreateWorkspaceRoleRequestModulePermissionsPermissions) GoString() string {
	return s.String()
}

func (s *CreateWorkspaceRoleRequestModulePermissionsPermissions) GetPermissionCodes() []*string {
	return s.PermissionCodes
}

func (s *CreateWorkspaceRoleRequestModulePermissionsPermissions) GetPermissionRules() []*CreateWorkspaceRoleRequestModulePermissionsPermissionsPermissionRules {
	return s.PermissionRules
}

func (s *CreateWorkspaceRoleRequestModulePermissionsPermissions) SetPermissionCodes(v []*string) *CreateWorkspaceRoleRequestModulePermissionsPermissions {
	s.PermissionCodes = v
	return s
}

func (s *CreateWorkspaceRoleRequestModulePermissionsPermissions) SetPermissionRules(v []*CreateWorkspaceRoleRequestModulePermissionsPermissionsPermissionRules) *CreateWorkspaceRoleRequestModulePermissionsPermissions {
	s.PermissionRules = v
	return s
}

func (s *CreateWorkspaceRoleRequestModulePermissionsPermissions) Validate() error {
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

type CreateWorkspaceRoleRequestModulePermissionsPermissionsPermissionRules struct {
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

func (s CreateWorkspaceRoleRequestModulePermissionsPermissionsPermissionRules) String() string {
	return dara.Prettify(s)
}

func (s CreateWorkspaceRoleRequestModulePermissionsPermissionsPermissionRules) GoString() string {
	return s.String()
}

func (s *CreateWorkspaceRoleRequestModulePermissionsPermissionsPermissionRules) GetAccessibility() *string {
	return s.Accessibility
}

func (s *CreateWorkspaceRoleRequestModulePermissionsPermissionsPermissionRules) GetEntityAccessType() *string {
	return s.EntityAccessType
}

func (s *CreateWorkspaceRoleRequestModulePermissionsPermissionsPermissionRules) SetAccessibility(v string) *CreateWorkspaceRoleRequestModulePermissionsPermissionsPermissionRules {
	s.Accessibility = &v
	return s
}

func (s *CreateWorkspaceRoleRequestModulePermissionsPermissionsPermissionRules) SetEntityAccessType(v string) *CreateWorkspaceRoleRequestModulePermissionsPermissionsPermissionRules {
	s.EntityAccessType = &v
	return s
}

func (s *CreateWorkspaceRoleRequestModulePermissionsPermissionsPermissionRules) Validate() error {
	return dara.Validate(s)
}
