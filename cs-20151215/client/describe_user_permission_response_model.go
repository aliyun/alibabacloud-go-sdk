// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserPermissionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeUserPermissionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeUserPermissionResponse
	GetStatusCode() *int32
	SetBody(v []*DescribeUserPermissionResponseBody) *DescribeUserPermissionResponse
	GetBody() []*DescribeUserPermissionResponseBody
}

type DescribeUserPermissionResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       []*DescribeUserPermissionResponseBody `json:"body,omitempty" xml:"body,omitempty" type:"Repeated"`
}

func (s DescribeUserPermissionResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserPermissionResponse) GoString() string {
	return s.String()
}

func (s *DescribeUserPermissionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeUserPermissionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeUserPermissionResponse) GetBody() []*DescribeUserPermissionResponseBody {
	return s.Body
}

func (s *DescribeUserPermissionResponse) SetHeaders(v map[string]*string) *DescribeUserPermissionResponse {
	s.Headers = v
	return s
}

func (s *DescribeUserPermissionResponse) SetStatusCode(v int32) *DescribeUserPermissionResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeUserPermissionResponse) SetBody(v []*DescribeUserPermissionResponseBody) *DescribeUserPermissionResponse {
	s.Body = v
	return s
}

func (s *DescribeUserPermissionResponse) Validate() error {
	return dara.Validate(s)
}

type DescribeUserPermissionResponseBody struct {
	// The authorization setting. Valid values:
	//
	// 	- `{cluster_id}` is returned if the permissions are scoped to a cluster.
	//
	// 	- `{cluster_id}/{namespace}` is returned if the permissions are scoped to a namespace of a cluster.
	//
	// 	- `all-clusters` is returned if the permissions are scoped to all clusters.
	//
	// example:
	//
	// c1b542****
	ResourceId *string `json:"resource_id,omitempty" xml:"resource_id,omitempty"`
	// The authorization type. Valid values:
	//
	// 	- `cluster`: indicates that the permissions are scoped to a cluster.
	//
	// 	- `namespace`: indicates that the permissions are scoped to a namespace of a cluster.
	//
	// 	- `console`: indicates that the permissions are scoped to all clusters.
	//
	// example:
	//
	// cluster
	ResourceType *string `json:"resource_type,omitempty" xml:"resource_type,omitempty"`
	// The name of the custom role. If a custom role is assigned, the value is the name of the assigned custom role.
	//
	// example:
	//
	// view
	RoleName *string `json:"role_name,omitempty" xml:"role_name,omitempty"`
	// The type of predefined role. Valid values:
	//
	// 	- `admin`: administrator
	//
	// 	- `ops`: O\\&M engineer
	//
	// 	- `dev`: developer
	//
	// 	- `restricted`: restricted user
	//
	// 	- `custom`: custom role
	//
	// example:
	//
	// admin
	RoleType *string `json:"role_type,omitempty" xml:"role_type,omitempty"`
	// Indicates whether the permissions are granted to the cluster owner.
	//
	// 	- `0`: indicates that the permissions are not granted to the cluster owner.
	//
	// 	- `1`: indicates that the permissions are granted to the cluster owner. The cluster owner is the administrator.
	//
	// example:
	//
	// 1
	IsOwner *int64 `json:"is_owner,omitempty" xml:"is_owner,omitempty"`
	// Indicates whether the permissions are granted to the RAM role. Valid values:
	//
	// 	- `0`: indicates that the permissions are not granted to the RAM role.
	//
	// 	- `1`: indicates that the permissions are granted to the RAM role.
	//
	// example:
	//
	// 1
	IsRamRole *int64 `json:"is_ram_role,omitempty" xml:"is_ram_role,omitempty"`
}

func (s DescribeUserPermissionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserPermissionResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUserPermissionResponseBody) GetResourceId() *string {
	return s.ResourceId
}

func (s *DescribeUserPermissionResponseBody) GetResourceType() *string {
	return s.ResourceType
}

func (s *DescribeUserPermissionResponseBody) GetRoleName() *string {
	return s.RoleName
}

func (s *DescribeUserPermissionResponseBody) GetRoleType() *string {
	return s.RoleType
}

func (s *DescribeUserPermissionResponseBody) GetIsOwner() *int64 {
	return s.IsOwner
}

func (s *DescribeUserPermissionResponseBody) GetIsRamRole() *int64 {
	return s.IsRamRole
}

func (s *DescribeUserPermissionResponseBody) SetResourceId(v string) *DescribeUserPermissionResponseBody {
	s.ResourceId = &v
	return s
}

func (s *DescribeUserPermissionResponseBody) SetResourceType(v string) *DescribeUserPermissionResponseBody {
	s.ResourceType = &v
	return s
}

func (s *DescribeUserPermissionResponseBody) SetRoleName(v string) *DescribeUserPermissionResponseBody {
	s.RoleName = &v
	return s
}

func (s *DescribeUserPermissionResponseBody) SetRoleType(v string) *DescribeUserPermissionResponseBody {
	s.RoleType = &v
	return s
}

func (s *DescribeUserPermissionResponseBody) SetIsOwner(v int64) *DescribeUserPermissionResponseBody {
	s.IsOwner = &v
	return s
}

func (s *DescribeUserPermissionResponseBody) SetIsRamRole(v int64) *DescribeUserPermissionResponseBody {
	s.IsRamRole = &v
	return s
}

func (s *DescribeUserPermissionResponseBody) Validate() error {
	return dara.Validate(s)
}
