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
	// 集群访问配置，格式为：
	//
	// - 当是集群维度授权时，格式为：`{cluster_id}`。
	//
	// - 当是命名空间维度授权时，格式为：`{cluster_id}/{namespace}`。
	//
	// - 当是所有集群授权时，值固定为：`all-clusters`。
	//
	// example:
	//
	// c1b542****
	ResourceId *string `json:"resource_id,omitempty" xml:"resource_id,omitempty"`
	// 授权类型，取值：
	//
	// - `cluster`：集群维度。
	//
	// - `namespace`：命名空间维度。
	//
	// - `console`：所有集群维度的授权。
	//
	// example:
	//
	// cluster
	ResourceType *string `json:"resource_type,omitempty" xml:"resource_type,omitempty"`
	// 自定义角色名称，当授权自定义角色时，该字段为指定的自定义集群管理角色名称。
	//
	// example:
	//
	// terway-pod-reader
	RoleName *string `json:"role_name,omitempty" xml:"role_name,omitempty"`
	// 预置的角色类型，取值：
	//
	// - `admin`：管理员。
	//
	// - `ops`：运维人员。
	//
	// - `dev`：开发人员。
	//
	// - `restricted`：受限用户。
	//
	// - `custom`：使用自定义的集群管理角色。
	//
	// example:
	//
	// admin
	RoleType *string `json:"role_type,omitempty" xml:"role_type,omitempty"`
	// 是否为集群创建者的授权，取值：
	//
	// - `0`：代表不是集群创建者的授权记录。
	//
	// - `1`：代表该授权记录为集群创建者的管理员权限。
	//
	// example:
	//
	// 1
	IsOwner *int64 `json:"is_owner,omitempty" xml:"is_owner,omitempty"`
	// 是否为RAM角色授权，取值：
	//
	// - `0`：代表当前记录不是RAM角色授权。
	//
	// - `1`：代表当前记录是RAM角色授权。
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
