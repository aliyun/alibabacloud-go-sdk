// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAuthorizedAgentsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPermission(v string) *ListAuthorizedAgentsRequest
	GetPermission() *string
	SetTargetUserId(v int64) *ListAuthorizedAgentsRequest
	GetTargetUserId() *int64
	SetTenantId(v string) *ListAuthorizedAgentsRequest
	GetTenantId() *string
}

type ListAuthorizedAgentsRequest struct {
	// 权限类型：USE=使用权限, MANAGE=管理权限，默认 USE
	//
	// example:
	//
	// USE
	Permission *string `json:"permission,omitempty" xml:"permission,omitempty"`
	// 目标用户 ID，管理员代查指定用户可用的数字员工时传入（需 APPLICATION_AGENT_VIEW 权限）；不传则查询调用方自身
	//
	// example:
	//
	// 1
	TargetUserId *int64 `json:"targetUserId,omitempty" xml:"targetUserId,omitempty"`
	// 租户ID，公共参数，缺省时使用调用方默认租户
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s ListAuthorizedAgentsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAuthorizedAgentsRequest) GoString() string {
	return s.String()
}

func (s *ListAuthorizedAgentsRequest) GetPermission() *string {
	return s.Permission
}

func (s *ListAuthorizedAgentsRequest) GetTargetUserId() *int64 {
	return s.TargetUserId
}

func (s *ListAuthorizedAgentsRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ListAuthorizedAgentsRequest) SetPermission(v string) *ListAuthorizedAgentsRequest {
	s.Permission = &v
	return s
}

func (s *ListAuthorizedAgentsRequest) SetTargetUserId(v int64) *ListAuthorizedAgentsRequest {
	s.TargetUserId = &v
	return s
}

func (s *ListAuthorizedAgentsRequest) SetTenantId(v string) *ListAuthorizedAgentsRequest {
	s.TenantId = &v
	return s
}

func (s *ListAuthorizedAgentsRequest) Validate() error {
	return dara.Validate(s)
}
