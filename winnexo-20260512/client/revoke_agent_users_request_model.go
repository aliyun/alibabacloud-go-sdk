// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokeAgentUsersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *RevokeAgentUsersRequest
	GetAppToken() *string
	SetOperatingObjectName(v string) *RevokeAgentUsersRequest
	GetOperatingObjectName() *string
	SetOpsUserCode(v string) *RevokeAgentUsersRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *RevokeAgentUsersRequest
	GetOpsUserId() *string
	SetTenantId(v string) *RevokeAgentUsersRequest
	GetTenantId() *string
	SetUserGroupIds(v []*string) *RevokeAgentUsersRequest
	GetUserGroupIds() []*string
	SetUserIds(v []*string) *RevokeAgentUsersRequest
	GetUserIds() []*string
}

type RevokeAgentUsersRequest struct {
	// 平台应用 Token（控制台签发）。传入后将以 opsUserId 或 opsUserCode 指定的用户身份执行
	//
	// example:
	//
	// wn_app_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
	AppToken *string `json:"appToken,omitempty" xml:"appToken,omitempty"`
	// The name of the digital human.
	//
	// This parameter is required.
	//
	// example:
	//
	// string_value
	OperatingObjectName *string `json:"operatingObjectName,omitempty" xml:"operatingObjectName,omitempty"`
	// 代理用户编码（租户内唯一），仅当应用凭证配置为 source_uid 时使用
	//
	// example:
	//
	// alice
	OpsUserCode *string `json:"opsUserCode,omitempty" xml:"opsUserCode,omitempty"`
	// 代理用户ID（WINNEXO 平台用户ID），仅当应用凭证配置为 winnexo_uid 时使用
	//
	// example:
	//
	// 10001
	OpsUserId *string `json:"opsUserId,omitempty" xml:"opsUserId,omitempty"`
	// The tenant ID. This is a common parameter. If not specified, the default tenant of the caller is used.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// The list of user group IDs to be revoked (16-character hex strings).
	//
	// example:
	//
	// string_value
	UserGroupIds []*string `json:"userGroupIds,omitempty" xml:"userGroupIds,omitempty" type:"Repeated"`
	// The list of user IDs to be revoked.
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

func (s *RevokeAgentUsersRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *RevokeAgentUsersRequest) GetOperatingObjectName() *string {
	return s.OperatingObjectName
}

func (s *RevokeAgentUsersRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *RevokeAgentUsersRequest) GetOpsUserId() *string {
	return s.OpsUserId
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

func (s *RevokeAgentUsersRequest) SetAppToken(v string) *RevokeAgentUsersRequest {
	s.AppToken = &v
	return s
}

func (s *RevokeAgentUsersRequest) SetOperatingObjectName(v string) *RevokeAgentUsersRequest {
	s.OperatingObjectName = &v
	return s
}

func (s *RevokeAgentUsersRequest) SetOpsUserCode(v string) *RevokeAgentUsersRequest {
	s.OpsUserCode = &v
	return s
}

func (s *RevokeAgentUsersRequest) SetOpsUserId(v string) *RevokeAgentUsersRequest {
	s.OpsUserId = &v
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
