// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokeAgentUsersShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOperatingObjectName(v string) *RevokeAgentUsersShrinkRequest
	GetOperatingObjectName() *string
	SetTenantId(v string) *RevokeAgentUsersShrinkRequest
	GetTenantId() *string
	SetUserGroupIdsShrink(v string) *RevokeAgentUsersShrinkRequest
	GetUserGroupIdsShrink() *string
	SetUserIdsShrink(v string) *RevokeAgentUsersShrinkRequest
	GetUserIdsShrink() *string
}

type RevokeAgentUsersShrinkRequest struct {
	// 数字员工名称
	//
	// This parameter is required.
	//
	// example:
	//
	// string_value
	OperatingObjectName *string `json:"operatingObjectName,omitempty" xml:"operatingObjectName,omitempty"`
	// 租户ID，公共参数，缺省时使用调用方默认租户
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// 被撤销的用户组 ID 列表（16位 hex 字符串）
	//
	// example:
	//
	// string_value
	UserGroupIdsShrink *string `json:"userGroupIds,omitempty" xml:"userGroupIds,omitempty"`
	// 被撤销的用户 ID 列表
	//
	// example:
	//
	// 1
	UserIdsShrink *string `json:"userIds,omitempty" xml:"userIds,omitempty"`
}

func (s RevokeAgentUsersShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s RevokeAgentUsersShrinkRequest) GoString() string {
	return s.String()
}

func (s *RevokeAgentUsersShrinkRequest) GetOperatingObjectName() *string {
	return s.OperatingObjectName
}

func (s *RevokeAgentUsersShrinkRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *RevokeAgentUsersShrinkRequest) GetUserGroupIdsShrink() *string {
	return s.UserGroupIdsShrink
}

func (s *RevokeAgentUsersShrinkRequest) GetUserIdsShrink() *string {
	return s.UserIdsShrink
}

func (s *RevokeAgentUsersShrinkRequest) SetOperatingObjectName(v string) *RevokeAgentUsersShrinkRequest {
	s.OperatingObjectName = &v
	return s
}

func (s *RevokeAgentUsersShrinkRequest) SetTenantId(v string) *RevokeAgentUsersShrinkRequest {
	s.TenantId = &v
	return s
}

func (s *RevokeAgentUsersShrinkRequest) SetUserGroupIdsShrink(v string) *RevokeAgentUsersShrinkRequest {
	s.UserGroupIdsShrink = &v
	return s
}

func (s *RevokeAgentUsersShrinkRequest) SetUserIdsShrink(v string) *RevokeAgentUsersShrinkRequest {
	s.UserIdsShrink = &v
	return s
}

func (s *RevokeAgentUsersShrinkRequest) Validate() error {
	return dara.Validate(s)
}
