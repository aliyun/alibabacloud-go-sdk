// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateUserPermissionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v []*UpdateUserPermissionsRequestBody) *UpdateUserPermissionsRequest
	GetBody() []*UpdateUserPermissionsRequestBody
	SetMode(v string) *UpdateUserPermissionsRequest
	GetMode() *string
}

type UpdateUserPermissionsRequest struct {
	// The request body.
	Body []*UpdateUserPermissionsRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Repeated"`
	// The authorization method. Valid values:
	//
	// 	- `apply`: The global update mode. Overwrites all existing permissions of the RAM user or RAM role on the cluster. You must specify all the permissions you want to grant to the RAM user or RAM role in the request parameters when you call this operation.
	//
	// 	- `delete`: The deletion mode. Revokes only the cluster permissions specified in the request, preserving other existing permissions of the RAM user or RAM role.
	//
	// 	- `patch`: The incremental mode. Adds only the cluster permissions specified in the request, preserving other existing permissions of the RAM user or RAM role.
	//
	// Default value: `apply`.
	//
	// example:
	//
	// apply
	Mode *string `json:"mode,omitempty" xml:"mode,omitempty"`
}

func (s UpdateUserPermissionsRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateUserPermissionsRequest) GoString() string {
	return s.String()
}

func (s *UpdateUserPermissionsRequest) GetBody() []*UpdateUserPermissionsRequestBody {
	return s.Body
}

func (s *UpdateUserPermissionsRequest) GetMode() *string {
	return s.Mode
}

func (s *UpdateUserPermissionsRequest) SetBody(v []*UpdateUserPermissionsRequestBody) *UpdateUserPermissionsRequest {
	s.Body = v
	return s
}

func (s *UpdateUserPermissionsRequest) SetMode(v string) *UpdateUserPermissionsRequest {
	s.Mode = &v
	return s
}

func (s *UpdateUserPermissionsRequest) Validate() error {
	return dara.Validate(s)
}

type UpdateUserPermissionsRequestBody struct {
	// The ID of the cluster on which you want to grant permissions to the RAM role or RAM role.
	//
	// 	- Set this parameter to an empty string if `role_type` is set to `all-clusters`.
	//
	// example:
	//
	// c796c60***
	Cluster *string `json:"cluster,omitempty" xml:"cluster,omitempty"`
	// Specifies whether to assign a custom role to the RAM user or RAM role. If you want to assign a custom role to the RAM user or RAM role, set `role_name` to the name of the custom role.
	//
	// example:
	//
	// false
	IsCustom *bool `json:"is_custom,omitempty" xml:"is_custom,omitempty"`
	// Specifies whether to use a RAM role to grant permissions.
	//
	// example:
	//
	// false
	IsRamRole *bool `json:"is_ram_role,omitempty" xml:"is_ram_role,omitempty"`
	// The namespace that you want to authorize the RAM user or RAM role to manage. This parameter is required only if you set role_type to namespace.
	//
	// example:
	//
	// test
	Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
	// The predefined role name. Valid values:
	//
	// 	- `admin`: administrator
	//
	// 	- `admin-view`: read-only administrator
	//
	// 	- `ops`: O\\&M engineer
	//
	// 	- `dev`: developer
	//
	// 	- `restricted`: restricted user
	//
	// 	- Custom role
	//
	// Note:
	//
	// 	- You cannot grant **namespace-level*	- permissions to the `admin`, `admin-view`, and `ops` roles.
	//
	// 	- You cannot grant **all cluster-level*	- permissions to the `admin-view` role.
	//
	// example:
	//
	// ops
	RoleName *string `json:"role_name,omitempty" xml:"role_name,omitempty"`
	// The authorization type. Valid values:
	//
	// 	- `cluster`: authorizes the RAM user or RAM role to manage the specified clusters.
	//
	// 	- `namespace`: authorizes the RAM user or RAM role to manage the specified namespaces.
	//
	// 	- `all-clusters`: authorizes the RAM user or RAM role to manage all clusters.
	//
	// example:
	//
	// cluster
	RoleType *string `json:"role_type,omitempty" xml:"role_type,omitempty"`
}

func (s UpdateUserPermissionsRequestBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateUserPermissionsRequestBody) GoString() string {
	return s.String()
}

func (s *UpdateUserPermissionsRequestBody) GetCluster() *string {
	return s.Cluster
}

func (s *UpdateUserPermissionsRequestBody) GetIsCustom() *bool {
	return s.IsCustom
}

func (s *UpdateUserPermissionsRequestBody) GetIsRamRole() *bool {
	return s.IsRamRole
}

func (s *UpdateUserPermissionsRequestBody) GetNamespace() *string {
	return s.Namespace
}

func (s *UpdateUserPermissionsRequestBody) GetRoleName() *string {
	return s.RoleName
}

func (s *UpdateUserPermissionsRequestBody) GetRoleType() *string {
	return s.RoleType
}

func (s *UpdateUserPermissionsRequestBody) SetCluster(v string) *UpdateUserPermissionsRequestBody {
	s.Cluster = &v
	return s
}

func (s *UpdateUserPermissionsRequestBody) SetIsCustom(v bool) *UpdateUserPermissionsRequestBody {
	s.IsCustom = &v
	return s
}

func (s *UpdateUserPermissionsRequestBody) SetIsRamRole(v bool) *UpdateUserPermissionsRequestBody {
	s.IsRamRole = &v
	return s
}

func (s *UpdateUserPermissionsRequestBody) SetNamespace(v string) *UpdateUserPermissionsRequestBody {
	s.Namespace = &v
	return s
}

func (s *UpdateUserPermissionsRequestBody) SetRoleName(v string) *UpdateUserPermissionsRequestBody {
	s.RoleName = &v
	return s
}

func (s *UpdateUserPermissionsRequestBody) SetRoleType(v string) *UpdateUserPermissionsRequestBody {
	s.RoleType = &v
	return s
}

func (s *UpdateUserPermissionsRequestBody) Validate() error {
	return dara.Validate(s)
}
