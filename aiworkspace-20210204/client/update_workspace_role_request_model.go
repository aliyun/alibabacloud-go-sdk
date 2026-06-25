// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWorkspaceRoleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetModulePermissions(v []*UpdateWorkspaceRoleRequestModulePermissions) *UpdateWorkspaceRoleRequest
	GetModulePermissions() []*UpdateWorkspaceRoleRequestModulePermissions
	SetRoleName(v string) *UpdateWorkspaceRoleRequest
	GetRoleName() *string
}

type UpdateWorkspaceRoleRequest struct {
	// The permission settings for the role.
	ModulePermissions []*UpdateWorkspaceRoleRequestModulePermissions `json:"ModulePermissions,omitempty" xml:"ModulePermissions,omitempty" type:"Repeated"`
	// The name of the custom role. The name must be unique within the workspace. It can be up to 64 characters long and can contain letters, digits, underscores (_), and hyphens (-).
	//
	// example:
	//
	// dev-test
	RoleName *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
}

func (s UpdateWorkspaceRoleRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateWorkspaceRoleRequest) GoString() string {
	return s.String()
}

func (s *UpdateWorkspaceRoleRequest) GetModulePermissions() []*UpdateWorkspaceRoleRequestModulePermissions {
	return s.ModulePermissions
}

func (s *UpdateWorkspaceRoleRequest) GetRoleName() *string {
	return s.RoleName
}

func (s *UpdateWorkspaceRoleRequest) SetModulePermissions(v []*UpdateWorkspaceRoleRequestModulePermissions) *UpdateWorkspaceRoleRequest {
	s.ModulePermissions = v
	return s
}

func (s *UpdateWorkspaceRoleRequest) SetRoleName(v string) *UpdateWorkspaceRoleRequest {
	s.RoleName = &v
	return s
}

func (s *UpdateWorkspaceRoleRequest) Validate() error {
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

type UpdateWorkspaceRoleRequestModulePermissions struct {
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
	Permissions []*UpdateWorkspaceRoleRequestModulePermissionsPermissions `json:"Permissions,omitempty" xml:"Permissions,omitempty" type:"Repeated"`
}

func (s UpdateWorkspaceRoleRequestModulePermissions) String() string {
	return dara.Prettify(s)
}

func (s UpdateWorkspaceRoleRequestModulePermissions) GoString() string {
	return s.String()
}

func (s *UpdateWorkspaceRoleRequestModulePermissions) GetModuleName() *string {
	return s.ModuleName
}

func (s *UpdateWorkspaceRoleRequestModulePermissions) GetPermissionType() *string {
	return s.PermissionType
}

func (s *UpdateWorkspaceRoleRequestModulePermissions) GetPermissions() []*UpdateWorkspaceRoleRequestModulePermissionsPermissions {
	return s.Permissions
}

func (s *UpdateWorkspaceRoleRequestModulePermissions) SetModuleName(v string) *UpdateWorkspaceRoleRequestModulePermissions {
	s.ModuleName = &v
	return s
}

func (s *UpdateWorkspaceRoleRequestModulePermissions) SetPermissionType(v string) *UpdateWorkspaceRoleRequestModulePermissions {
	s.PermissionType = &v
	return s
}

func (s *UpdateWorkspaceRoleRequestModulePermissions) SetPermissions(v []*UpdateWorkspaceRoleRequestModulePermissionsPermissions) *UpdateWorkspaceRoleRequestModulePermissions {
	s.Permissions = v
	return s
}

func (s *UpdateWorkspaceRoleRequestModulePermissions) Validate() error {
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

type UpdateWorkspaceRoleRequestModulePermissionsPermissions struct {
	// A list of permissions.
	PermissionCodes []*string `json:"PermissionCodes,omitempty" xml:"PermissionCodes,omitempty" type:"Repeated"`
	// A list of permission rules.
	PermissionRules []*UpdateWorkspaceRoleRequestModulePermissionsPermissionsPermissionRules `json:"PermissionRules,omitempty" xml:"PermissionRules,omitempty" type:"Repeated"`
}

func (s UpdateWorkspaceRoleRequestModulePermissionsPermissions) String() string {
	return dara.Prettify(s)
}

func (s UpdateWorkspaceRoleRequestModulePermissionsPermissions) GoString() string {
	return s.String()
}

func (s *UpdateWorkspaceRoleRequestModulePermissionsPermissions) GetPermissionCodes() []*string {
	return s.PermissionCodes
}

func (s *UpdateWorkspaceRoleRequestModulePermissionsPermissions) GetPermissionRules() []*UpdateWorkspaceRoleRequestModulePermissionsPermissionsPermissionRules {
	return s.PermissionRules
}

func (s *UpdateWorkspaceRoleRequestModulePermissionsPermissions) SetPermissionCodes(v []*string) *UpdateWorkspaceRoleRequestModulePermissionsPermissions {
	s.PermissionCodes = v
	return s
}

func (s *UpdateWorkspaceRoleRequestModulePermissionsPermissions) SetPermissionRules(v []*UpdateWorkspaceRoleRequestModulePermissionsPermissionsPermissionRules) *UpdateWorkspaceRoleRequestModulePermissionsPermissions {
	s.PermissionRules = v
	return s
}

func (s *UpdateWorkspaceRoleRequestModulePermissionsPermissions) Validate() error {
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

type UpdateWorkspaceRoleRequestModulePermissionsPermissionsPermissionRules struct {
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

func (s UpdateWorkspaceRoleRequestModulePermissionsPermissionsPermissionRules) String() string {
	return dara.Prettify(s)
}

func (s UpdateWorkspaceRoleRequestModulePermissionsPermissionsPermissionRules) GoString() string {
	return s.String()
}

func (s *UpdateWorkspaceRoleRequestModulePermissionsPermissionsPermissionRules) GetAccessibility() *string {
	return s.Accessibility
}

func (s *UpdateWorkspaceRoleRequestModulePermissionsPermissionsPermissionRules) GetEntityAccessType() *string {
	return s.EntityAccessType
}

func (s *UpdateWorkspaceRoleRequestModulePermissionsPermissionsPermissionRules) SetAccessibility(v string) *UpdateWorkspaceRoleRequestModulePermissionsPermissionsPermissionRules {
	s.Accessibility = &v
	return s
}

func (s *UpdateWorkspaceRoleRequestModulePermissionsPermissionsPermissionRules) SetEntityAccessType(v string) *UpdateWorkspaceRoleRequestModulePermissionsPermissionsPermissionRules {
	s.EntityAccessType = &v
	return s
}

func (s *UpdateWorkspaceRoleRequestModulePermissionsPermissionsPermissionRules) Validate() error {
	return dara.Validate(s)
}
