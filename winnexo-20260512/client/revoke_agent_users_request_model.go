// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokeAgentUsersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOperatingObjectName(v string) *RevokeAgentUsersRequest
	GetOperatingObjectName() *string
	SetTenantId(v string) *RevokeAgentUsersRequest
	GetTenantId() *string
	SetUserGroupIds(v []*string) *RevokeAgentUsersRequest
	GetUserGroupIds() []*string
	SetUserIds(v []*string) *RevokeAgentUsersRequest
	GetUserIds() []*string
}

type RevokeAgentUsersRequest struct {
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
	UserGroupIds []*string `json:"userGroupIds,omitempty" xml:"userGroupIds,omitempty" type:"Repeated"`
	// 被撤销的用户 ID 列表
	//
	// example:
	//
	// 1
	UserIds []*string `json:"userIds,omitempty" xml:"userIds,omitempty" type:"Repeated"`
}

func (s RevokeAgentUsersRequest) String() string {
	return dara.Prettify(s)
}

func (s RevokeAgentUsersRequest) GoString() string {
	return s.String()
}

func (s *RevokeAgentUsersRequest) GetOperatingObjectName() *string {
	return s.OperatingObjectName
}

func (s *RevokeAgentUsersRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *RevokeAgentUsersRequest) GetUserGroupIds() []*string {
	return s.UserGroupIds
}

func (s *RevokeAgentUsersRequest) GetUserIds() []*string {
	return s.UserIds
}

func (s *RevokeAgentUsersRequest) SetOperatingObjectName(v string) *RevokeAgentUsersRequest {
	s.OperatingObjectName = &v
	return s
}

func (s *RevokeAgentUsersRequest) SetTenantId(v string) *RevokeAgentUsersRequest {
	s.TenantId = &v
	return s
}

func (s *RevokeAgentUsersRequest) SetUserGroupIds(v []*string) *RevokeAgentUsersRequest {
	s.UserGroupIds = v
	return s
}

func (s *RevokeAgentUsersRequest) SetUserIds(v []*string) *RevokeAgentUsersRequest {
	s.UserIds = v
	return s
}

func (s *RevokeAgentUsersRequest) Validate() error {
	return dara.Validate(s)
}
