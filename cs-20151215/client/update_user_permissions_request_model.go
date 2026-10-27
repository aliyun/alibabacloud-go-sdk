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
	// The request body parameters.
	Body []*UpdateUserPermissionsRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Repeated"`
	// The authorization method. Valid values:
	//
	// - `apply`: Fully updates permissions. This operation overwrites the existing cluster permissions of the target RAM user or RAM role. The request must include all permission configurations to be granted to the target RAM user or RAM role.
	//
	// - `delete`: Deletes permissions. This operation deletes only the cluster authorization information specified in the request and does not affect other cluster permissions of the RAM user or RAM role.
	//
	// - `patch`: Adds permissions. This operation adds only the cluster authorization information specified in the request and does not affect other cluster permissions of the RAM user or RAM role.
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
	if s.Body != nil {
		for _, item := range s.Body {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateUserPermissionsRequestBody struct {
	// The ID of the target cluster.
	//
	// This parameter is not required when the role_type parameter is set to all-clusters.
	//
	// example:
	//
	// c796c60***
	Cluster *string `json:"cluster,omitempty" xml:"cluster,omitempty"`
	// Specifies whether the authorization is custom. A custom authorization uses a custom ClusterRole name for the role_name parameter.
	//
	// Valid values:
	//
	// - true: The authorized role is a custom cluster role.
	//
	// - false: The authorized role is a preset cluster role.
	//
	// example:
	//
	// false
	IsCustom *bool `json:"is_custom,omitempty" xml:"is_custom,omitempty"`
	// Specifies whether the authorized entity is a RAM role.
	//
	// Valid values:
	//
	// - true: The authorization is for a RAM role.
	//
	// - false: The authorization is for a Resource Access Management (RAM) user.
	//
	// example:
	//
	// false
	IsRamRole *bool `json:"is_ram_role,omitempty" xml:"is_ram_role,omitempty"`
	// The namespace name. This parameter is empty by default for cluster-level authorization.
	//
	// example:
	//
	// test
	Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
	// The preset role name. Valid values:
	//
	// - admin: administrator.
	//
	// - admin-view: read-only administrator.
	//
	// - ops: O&M engineer.
	//
	// - dev: developer.
	//
	// - restricted: restricted user.
	//
	// - A custom ClusterRole name.
	//
	// > - The admin, admin-view, and ops roles do not support authorization at the **namespace level**.
	//
	// > - The admin-view role does not support authorization at the **all-clusters level**.
	//
	// example:
	//
	// ops
	RoleName *string `json:"role_name,omitempty" xml:"role_name,omitempty"`
	// The authorization type. Valid values:
	//
	// - cluster: cluster level.
	//
	// - namespace: namespace level.
	//
	// - all-clusters: all-clusters level.
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
