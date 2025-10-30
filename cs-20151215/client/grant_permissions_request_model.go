// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGrantPermissionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v []*GrantPermissionsRequestBody) *GrantPermissionsRequest
	GetBody() []*GrantPermissionsRequestBody
}

type GrantPermissionsRequest struct {
	// The request body.
	Body []*GrantPermissionsRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Repeated"`
}

func (s GrantPermissionsRequest) String() string {
	return dara.Prettify(s)
}

func (s GrantPermissionsRequest) GoString() string {
	return s.String()
}

func (s *GrantPermissionsRequest) GetBody() []*GrantPermissionsRequestBody {
	return s.Body
}

func (s *GrantPermissionsRequest) SetBody(v []*GrantPermissionsRequestBody) *GrantPermissionsRequest {
	s.Body = v
	return s
}

func (s *GrantPermissionsRequest) Validate() error {
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

type GrantPermissionsRequestBody struct {
	// The ID of the cluster on which you want to grant permissions to the RAM role or RAM role.
	//
	// 	- Set this parameter to an empty string if `role_type` is set to `all-clusters`.
	//
	// This parameter is required.
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
	// 	- You cannot grant namespace-level permissions to the `admin`, `admin-view`, and `ops` roles.
	//
	// 	- You cannot grant all cluster-level permissions to the `admin-view` role.
	//
	// This parameter is required.
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
	// This parameter is required.
	//
	// example:
	//
	// cluster
	RoleType *string `json:"role_type,omitempty" xml:"role_type,omitempty"`
}

func (s GrantPermissionsRequestBody) String() string {
	return dara.Prettify(s)
}

func (s GrantPermissionsRequestBody) GoString() string {
	return s.String()
}

func (s *GrantPermissionsRequestBody) GetCluster() *string {
	return s.Cluster
}

func (s *GrantPermissionsRequestBody) GetIsCustom() *bool {
	return s.IsCustom
}

func (s *GrantPermissionsRequestBody) GetIsRamRole() *bool {
	return s.IsRamRole
}

func (s *GrantPermissionsRequestBody) GetNamespace() *string {
	return s.Namespace
}

func (s *GrantPermissionsRequestBody) GetRoleName() *string {
	return s.RoleName
}

func (s *GrantPermissionsRequestBody) GetRoleType() *string {
	return s.RoleType
}

func (s *GrantPermissionsRequestBody) SetCluster(v string) *GrantPermissionsRequestBody {
	s.Cluster = &v
	return s
}

func (s *GrantPermissionsRequestBody) SetIsCustom(v bool) *GrantPermissionsRequestBody {
	s.IsCustom = &v
	return s
}

func (s *GrantPermissionsRequestBody) SetIsRamRole(v bool) *GrantPermissionsRequestBody {
	s.IsRamRole = &v
	return s
}

func (s *GrantPermissionsRequestBody) SetNamespace(v string) *GrantPermissionsRequestBody {
	s.Namespace = &v
	return s
}

func (s *GrantPermissionsRequestBody) SetRoleName(v string) *GrantPermissionsRequestBody {
	s.RoleName = &v
	return s
}

func (s *GrantPermissionsRequestBody) SetRoleType(v string) *GrantPermissionsRequestBody {
	s.RoleType = &v
	return s
}

func (s *GrantPermissionsRequestBody) Validate() error {
	return dara.Validate(s)
}
