// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddUserGroupMembersShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *AddUserGroupMembersShrinkRequest
	GetAppToken() *string
	SetOpsUserCode(v string) *AddUserGroupMembersShrinkRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *AddUserGroupMembersShrinkRequest
	GetOpsUserId() *string
	SetTenantId(v string) *AddUserGroupMembersShrinkRequest
	GetTenantId() *string
	SetUserGroupId(v string) *AddUserGroupMembersShrinkRequest
	GetUserGroupId() *string
	SetUserIdsShrink(v string) *AddUserGroupMembersShrinkRequest
	GetUserIdsShrink() *string
}

type AddUserGroupMembersShrinkRequest struct {
	// 平台应用 Token（控制台签发）。传入后将以 opsUserId 或 opsUserCode 指定的用户身份执行
	//
	// example:
	//
	// wn_app_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
	AppToken *string `json:"appToken,omitempty" xml:"appToken,omitempty"`
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
	// The tenant ID. This is a common parameter. In winnexo-cli, pass this parameter explicitly by using `--tenant-id`.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// The ID of the target user group.
	//
	// This parameter is required.
	//
	// example:
	//
	// 7ea8973f-7a5c-4e8a-956b-4fe0e7e2eb11
	UserGroupId *string `json:"userGroupId,omitempty" xml:"userGroupId,omitempty"`
	// The list of platform user IDs to add. Supports single or batch input. Duplicate relationships are idempotent.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	UserIdsShrink *string `json:"userIds,omitempty" xml:"userIds,omitempty"`
}

func (s AddUserGroupMembersShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s AddUserGroupMembersShrinkRequest) GoString() string {
	return s.String()
}

func (s *AddUserGroupMembersShrinkRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *AddUserGroupMembersShrinkRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *AddUserGroupMembersShrinkRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *AddUserGroupMembersShrinkRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *AddUserGroupMembersShrinkRequest) GetUserGroupId() *string {
	return s.UserGroupId
}

func (s *AddUserGroupMembersShrinkRequest) GetUserIdsShrink() *string {
	return s.UserIdsShrink
}

func (s *AddUserGroupMembersShrinkRequest) SetAppToken(v string) *AddUserGroupMembersShrinkRequest {
	s.AppToken = &v
	return s
}

func (s *AddUserGroupMembersShrinkRequest) SetOpsUserCode(v string) *AddUserGroupMembersShrinkRequest {
	s.OpsUserCode = &v
	return s
}

func (s *AddUserGroupMembersShrinkRequest) SetOpsUserId(v string) *AddUserGroupMembersShrinkRequest {
	s.OpsUserId = &v
	return s
}

func (s *AddUserGroupMembersShrinkRequest) SetTenantId(v string) *AddUserGroupMembersShrinkRequest {
	s.TenantId = &v
	return s
}

func (s *AddUserGroupMembersShrinkRequest) SetUserGroupId(v string) *AddUserGroupMembersShrinkRequest {
	s.UserGroupId = &v
	return s
}

func (s *AddUserGroupMembersShrinkRequest) SetUserIdsShrink(v string) *AddUserGroupMembersShrinkRequest {
	s.UserIdsShrink = &v
	return s
}

func (s *AddUserGroupMembersShrinkRequest) Validate() error {
	return dara.Validate(s)
}
