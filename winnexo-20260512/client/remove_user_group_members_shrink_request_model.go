// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveUserGroupMembersShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *RemoveUserGroupMembersShrinkRequest
	GetAppToken() *string
	SetOpsUserCode(v string) *RemoveUserGroupMembersShrinkRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *RemoveUserGroupMembersShrinkRequest
	GetOpsUserId() *string
	SetTenantId(v string) *RemoveUserGroupMembersShrinkRequest
	GetTenantId() *string
	SetUserGroupId(v string) *RemoveUserGroupMembersShrinkRequest
	GetUserGroupId() *string
	SetUserIdsShrink(v string) *RemoveUserGroupMembersShrinkRequest
	GetUserIdsShrink() *string
}

type RemoveUserGroupMembersShrinkRequest struct {
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
	// The tenant ID. This is a common parameter. You can explicitly pass this parameter in winnexo-cli by using `--tenant-id`.
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
	// The list of platform user IDs to be removed. You can specify one or more IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	UserIdsShrink *string `json:"userIds,omitempty" xml:"userIds,omitempty"`
}

func (s RemoveUserGroupMembersShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveUserGroupMembersShrinkRequest) GoString() string {
	return s.String()
}

func (s *RemoveUserGroupMembersShrinkRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *RemoveUserGroupMembersShrinkRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *RemoveUserGroupMembersShrinkRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *RemoveUserGroupMembersShrinkRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *RemoveUserGroupMembersShrinkRequest) GetUserGroupId() *string {
	return s.UserGroupId
}

func (s *RemoveUserGroupMembersShrinkRequest) GetUserIdsShrink() *string {
	return s.UserIdsShrink
}

func (s *RemoveUserGroupMembersShrinkRequest) SetAppToken(v string) *RemoveUserGroupMembersShrinkRequest {
	s.AppToken = &v
	return s
}

func (s *RemoveUserGroupMembersShrinkRequest) SetOpsUserCode(v string) *RemoveUserGroupMembersShrinkRequest {
	s.OpsUserCode = &v
	return s
}

func (s *RemoveUserGroupMembersShrinkRequest) SetOpsUserId(v string) *RemoveUserGroupMembersShrinkRequest {
	s.OpsUserId = &v
	return s
}

func (s *RemoveUserGroupMembersShrinkRequest) SetTenantId(v string) *RemoveUserGroupMembersShrinkRequest {
	s.TenantId = &v
	return s
}

func (s *RemoveUserGroupMembersShrinkRequest) SetUserGroupId(v string) *RemoveUserGroupMembersShrinkRequest {
	s.UserGroupId = &v
	return s
}

func (s *RemoveUserGroupMembersShrinkRequest) SetUserIdsShrink(v string) *RemoveUserGroupMembersShrinkRequest {
	s.UserIdsShrink = &v
	return s
}

func (s *RemoveUserGroupMembersShrinkRequest) Validate() error {
	return dara.Validate(s)
}
