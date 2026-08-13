// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGrantAgentUsersShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExpireDate(v int64) *GrantAgentUsersShrinkRequest
	GetExpireDate() *int64
	SetOperatingObjectName(v string) *GrantAgentUsersShrinkRequest
	GetOperatingObjectName() *string
	SetPermissionsShrink(v string) *GrantAgentUsersShrinkRequest
	GetPermissionsShrink() *string
	SetTenantId(v string) *GrantAgentUsersShrinkRequest
	GetTenantId() *string
	SetUserGroupIdsShrink(v string) *GrantAgentUsersShrinkRequest
	GetUserGroupIdsShrink() *string
	SetUserIdsShrink(v string) *GrantAgentUsersShrinkRequest
	GetUserIdsShrink() *string
}

type GrantAgentUsersShrinkRequest struct {
	// 授权截止时间戳（毫秒），不传表示永不过期
	//
	// example:
	//
	// 1
	ExpireDate *int64 `json:"expireDate,omitempty" xml:"expireDate,omitempty"`
	// 数字员工名称
	//
	// This parameter is required.
	//
	// example:
	//
	// string_value
	OperatingObjectName *string `json:"operatingObjectName,omitempty" xml:"operatingObjectName,omitempty"`
	// 权限列表：USE（使用权限）和/或 MANAGE（管理权限），不传时默认仅 USE；不得为空列表
	//
	// example:
	//
	// string_value
	PermissionsShrink *string `json:"permissions,omitempty" xml:"permissions,omitempty"`
	// 租户ID，公共参数，缺省时使用调用方默认租户
	//
	// example:
	//
	// 676577544219585
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// 被授权的用户组 ID 列表（16位 hex 字符串）
	//
	// example:
	//
	// string_value
	UserGroupIdsShrink *string `json:"userGroupIds,omitempty" xml:"userGroupIds,omitempty"`
	// 被授权的用户 ID 列表
	//
	// example:
	//
	// 1
	UserIdsShrink *string `json:"userIds,omitempty" xml:"userIds,omitempty"`
}

func (s GrantAgentUsersShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GrantAgentUsersShrinkRequest) GoString() string {
	return s.String()
}

func (s *GrantAgentUsersShrinkRequest) GetExpireDate() *int64 {
	return s.ExpireDate
}

func (s *GrantAgentUsersShrinkRequest) GetOperatingObjectName() *string {
	return s.OperatingObjectName
}

func (s *GrantAgentUsersShrinkRequest) GetPermissionsShrink() *string {
	return s.PermissionsShrink
}

func (s *GrantAgentUsersShrinkRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *GrantAgentUsersShrinkRequest) GetUserGroupIdsShrink() *string {
	return s.UserGroupIdsShrink
}

func (s *GrantAgentUsersShrinkRequest) GetUserIdsShrink() *string {
	return s.UserIdsShrink
}

func (s *GrantAgentUsersShrinkRequest) SetExpireDate(v int64) *GrantAgentUsersShrinkRequest {
	s.ExpireDate = &v
	return s
}

func (s *GrantAgentUsersShrinkRequest) SetOperatingObjectName(v string) *GrantAgentUsersShrinkRequest {
	s.OperatingObjectName = &v
	return s
}

func (s *GrantAgentUsersShrinkRequest) SetPermissionsShrink(v string) *GrantAgentUsersShrinkRequest {
	s.PermissionsShrink = &v
	return s
}

func (s *GrantAgentUsersShrinkRequest) SetTenantId(v string) *GrantAgentUsersShrinkRequest {
	s.TenantId = &v
	return s
}

func (s *GrantAgentUsersShrinkRequest) SetUserGroupIdsShrink(v string) *GrantAgentUsersShrinkRequest {
	s.UserGroupIdsShrink = &v
	return s
}

func (s *GrantAgentUsersShrinkRequest) SetUserIdsShrink(v string) *GrantAgentUsersShrinkRequest {
	s.UserIdsShrink = &v
	return s
}

func (s *GrantAgentUsersShrinkRequest) Validate() error {
	return dara.Validate(s)
}
