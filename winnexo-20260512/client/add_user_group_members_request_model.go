// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddUserGroupMembersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *AddUserGroupMembersRequest
	GetAppToken() *string
	SetOpsUserCode(v string) *AddUserGroupMembersRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *AddUserGroupMembersRequest
	GetOpsUserId() *string
	SetTenantId(v string) *AddUserGroupMembersRequest
	GetTenantId() *string
	SetUserGroupId(v string) *AddUserGroupMembersRequest
	GetUserGroupId() *string
	SetUserIds(v []*int64) *AddUserGroupMembersRequest
	GetUserIds() []*int64
}

type AddUserGroupMembersRequest struct {
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
	UserIds []*int64 `json:"userIds,omitempty" xml:"userIds,omitempty" type:"Repeated"`
}

func (s AddUserGroupMembersRequest) String() string {
	return dara.Prettify(s)
}

func (s AddUserGroupMembersRequest) GoString() string {
	return s.String()
}

func (s *AddUserGroupMembersRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *AddUserGroupMembersRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *AddUserGroupMembersRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *AddUserGroupMembersRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *AddUserGroupMembersRequest) GetUserGroupId() *string {
	return s.UserGroupId
}

func (s *AddUserGroupMembersRequest) GetUserIds() []*int64 {
	return s.UserIds
}

func (s *AddUserGroupMembersRequest) SetAppToken(v string) *AddUserGroupMembersRequest {
	s.AppToken = &v
	return s
}

func (s *AddUserGroupMembersRequest) SetOpsUserCode(v string) *AddUserGroupMembersRequest {
	s.OpsUserCode = &v
	return s
}

func (s *AddUserGroupMembersRequest) SetOpsUserId(v string) *AddUserGroupMembersRequest {
	s.OpsUserId = &v
	return s
}

func (s *AddUserGroupMembersRequest) SetTenantId(v string) *AddUserGroupMembersRequest {
	s.TenantId = &v
	return s
}

func (s *AddUserGroupMembersRequest) SetUserGroupId(v string) *AddUserGroupMembersRequest {
	s.UserGroupId = &v
	return s
}

func (s *AddUserGroupMembersRequest) SetUserIds(v []*int64) *AddUserGroupMembersRequest {
	s.UserIds = v
	return s
}

func (s *AddUserGroupMembersRequest) Validate() error {
	return dara.Validate(s)
}
