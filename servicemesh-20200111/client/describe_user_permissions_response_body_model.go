// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserPermissionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPermissions(v []*DescribeUserPermissionsResponseBodyPermissions) *DescribeUserPermissionsResponseBody
	GetPermissions() []*DescribeUserPermissionsResponseBodyPermissions
	SetRequestId(v string) *DescribeUserPermissionsResponseBody
	GetRequestId() *string
}

type DescribeUserPermissionsResponseBody struct {
	// The permissions that are granted to an entity.
	Permissions []*DescribeUserPermissionsResponseBodyPermissions `json:"Permissions,omitempty" xml:"Permissions,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 5A7C9E37-C171-584F-9A99-869B48C4196D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeUserPermissionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserPermissionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUserPermissionsResponseBody) GetPermissions() []*DescribeUserPermissionsResponseBodyPermissions {
	return s.Permissions
}

func (s *DescribeUserPermissionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeUserPermissionsResponseBody) SetPermissions(v []*DescribeUserPermissionsResponseBodyPermissions) *DescribeUserPermissionsResponseBody {
	s.Permissions = v
	return s
}

func (s *DescribeUserPermissionsResponseBody) SetRequestId(v string) *DescribeUserPermissionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeUserPermissionsResponseBody) Validate() error {
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

type DescribeUserPermissionsResponseBodyPermissions struct {
	// The entity to which the permissions are granted. A value of `true` indicates that the permissions are granted to a RAM user. A value of `false` indicates that the permissions are granted to a RAM role.
	//
	// example:
	//
	// false
	IsRamRole *string `json:"IsRamRole,omitempty" xml:"IsRamRole,omitempty"`
	// The value is fixed as `0`.
	//
	// example:
	//
	// 0
	ParentId *string `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
	// The ID of the ASM instance.
	//
	// example:
	//
	// c57b848115458460583a4260cb713****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The value is fixed as `cluster`.
	//
	// example:
	//
	// cluster
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The name of the permissions. Valid values:
	//
	// 	- `istio-admin`: the permissions of ASM administrators.
	//
	// 	- `istio-ops`: the permissions of ASM restricted users.
	//
	// 	- `istio-readonly`: the read-only permissions.
	//
	// example:
	//
	// istio-admin
	RoleName *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
	// The value is fixed as `custom`.
	//
	// example:
	//
	// custom
	RoleType *string `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
}

func (s DescribeUserPermissionsResponseBodyPermissions) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserPermissionsResponseBodyPermissions) GoString() string {
	return s.String()
}

func (s *DescribeUserPermissionsResponseBodyPermissions) GetIsRamRole() *string {
	return s.IsRamRole
}

func (s *DescribeUserPermissionsResponseBodyPermissions) GetParentId() *string {
	return s.ParentId
}

func (s *DescribeUserPermissionsResponseBodyPermissions) GetResourceId() *string {
	return s.ResourceId
}

func (s *DescribeUserPermissionsResponseBodyPermissions) GetResourceType() *string {
	return s.ResourceType
}

func (s *DescribeUserPermissionsResponseBodyPermissions) GetRoleName() *string {
	return s.RoleName
}

func (s *DescribeUserPermissionsResponseBodyPermissions) GetRoleType() *string {
	return s.RoleType
}

func (s *DescribeUserPermissionsResponseBodyPermissions) SetIsRamRole(v string) *DescribeUserPermissionsResponseBodyPermissions {
	s.IsRamRole = &v
	return s
}

func (s *DescribeUserPermissionsResponseBodyPermissions) SetParentId(v string) *DescribeUserPermissionsResponseBodyPermissions {
	s.ParentId = &v
	return s
}

func (s *DescribeUserPermissionsResponseBodyPermissions) SetResourceId(v string) *DescribeUserPermissionsResponseBodyPermissions {
	s.ResourceId = &v
	return s
}

func (s *DescribeUserPermissionsResponseBodyPermissions) SetResourceType(v string) *DescribeUserPermissionsResponseBodyPermissions {
	s.ResourceType = &v
	return s
}

func (s *DescribeUserPermissionsResponseBodyPermissions) SetRoleName(v string) *DescribeUserPermissionsResponseBodyPermissions {
	s.RoleName = &v
	return s
}

func (s *DescribeUserPermissionsResponseBodyPermissions) SetRoleType(v string) *DescribeUserPermissionsResponseBodyPermissions {
	s.RoleType = &v
	return s
}

func (s *DescribeUserPermissionsResponseBodyPermissions) Validate() error {
	return dara.Validate(s)
}
