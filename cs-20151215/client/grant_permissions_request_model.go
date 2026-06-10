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
	// The request body parameters.
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
	// The ID of the target cluster.
	//
	// - If you set the `role_type` parameter to `all-clusters`, set this parameter to an empty string.
	//
	// This parameter is required.
	//
	// example:
	//
	// c796c60***
	Cluster *string `json:"cluster,omitempty" xml:"cluster,omitempty"`
	// Set to true if `role_name` specifies a custom ClusterRole.
	//
	// example:
	//
	// false
	IsCustom *bool `json:"is_custom,omitempty" xml:"is_custom,omitempty"`
	// Set to true if you are granting permissions to a RAM role.
	//
	// example:
	//
	// false
	IsRamRole *bool `json:"is_ram_role,omitempty" xml:"is_ram_role,omitempty"`
	// The name of the namespace. This parameter is required only when `role_type` is set to `namespace`.
	//
	// example:
	//
	// test
	Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
	// The name of the role to grant. Valid values:
	//
	// - `admin`: The administrator role.
	//
	// - `admin-view`: The read-only administrator role.
	//
	// - `ops`: The operations role.
	//
	// - `dev`: The developer role.
	//
	// - `restricted`: The restricted role.
	//
	// - The name of a custom ClusterRole.
	//
	// 	Notice:
	//
	// - The `admin`, `admin-view`, and `ops` roles cannot be granted at the namespace scope.
	//
	// - The `admin-view` role is not currently supported for the all-clusters scope.
	//
	// This parameter is required.
	//
	// example:
	//
	// ops
	RoleName *string `json:"role_name,omitempty" xml:"role_name,omitempty"`
	// The authorization scope. Valid values:
	//
	// - `cluster`: Grants permissions at the cluster scope.
	//
	// - `namespace`: Grants permissions at the namespace scope.
	//
	// - `all-clusters`: Grants permissions at the all-clusters scope.
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
