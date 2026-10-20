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

type DescribeUserPermissionResponseBody struct {
	// The cluster access configuration. The format varies based on the authorization scope:
	//
	// - Cluster-level authorization: `{cluster_id}`.
	//
	// - Namespace-level authorization: `{cluster_id}/{namespace}`.
	//
	// - All-cluster authorization: `all-clusters`.
	//
	// example:
	//
	// c1b542****
	ResourceId *string `json:"resource_id,omitempty" xml:"resource_id,omitempty"`
	// The authorization type. Valid values:
	//
	// - `cluster`: cluster-level.
	//
	// - `namespace`: namespace-level.
	//
	// - `console`: all-cluster-level authorization.
	//
	// example:
	//
	// cluster
	ResourceType *string `json:"resource_type,omitempty" xml:"resource_type,omitempty"`
	// The custom role name. When a custom role is authorized, this field specifies the name of the custom cluster management role.
	//
	// example:
	//
	// terway-pod-reader
	RoleName *string `json:"role_name,omitempty" xml:"role_name,omitempty"`
	// The preset role type. Valid values:
	//
	// - `admin`: Administrator.
	//
	// - `ops`: O&M engineer.
	//
	// - `dev`: Developer.
	//
	// - `restricted`: Restricted user.
	//
	// - `custom`: Custom cluster management role.
	//
	// example:
	//
	// admin
	RoleType *string `json:"role_type,omitempty" xml:"role_type,omitempty"`
	// Indicates whether the authorization belongs to the cluster creator. Valid values:
	//
	// - `0`: The authorization record does not belong to the cluster creator.
	//
	// - `1`: The authorization record represents the administrator permissions of the cluster creator.
	//
	// example:
	//
	// 1
	IsOwner *int64 `json:"is_owner,omitempty" xml:"is_owner,omitempty"`
	// Indicates whether the authorization is for a RAM role. Valid values:
	//
	// - `0`: The record is not a RAM role authorization.
	//
	// - `1`: The record is a RAM role authorization.
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
