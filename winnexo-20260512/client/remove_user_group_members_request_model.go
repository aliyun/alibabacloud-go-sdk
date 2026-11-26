// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveUserGroupMembersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *RemoveUserGroupMembersRequest
	GetAppToken() *string
	SetOpsUserCode(v string) *RemoveUserGroupMembersRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *RemoveUserGroupMembersRequest
	GetOpsUserId() *string
	SetTenantId(v string) *RemoveUserGroupMembersRequest
	GetTenantId() *string
	SetUserGroupId(v string) *RemoveUserGroupMembersRequest
	GetUserGroupId() *string
	SetUserIds(v []*int64) *RemoveUserGroupMembersRequest
	GetUserIds() []*int64
}

type RemoveUserGroupMembersRequest struct {
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
	UserIds []*int64 `json:"userIds,omitempty" xml:"userIds,omitempty" type:"Repeated"`
}

func (s RemoveUserGroupMembersRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveUserGroupMembersRequest) GoString() string {
	return s.String()
}

func (s *RemoveUserGroupMembersRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *RemoveUserGroupMembersRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *RemoveUserGroupMembersRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *RemoveUserGroupMembersRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *RemoveUserGroupMembersRequest) GetUserGroupId() *string {
	return s.UserGroupId
}

func (s *RemoveUserGroupMembersRequest) GetUserIds() []*int64 {
	return s.UserIds
}

func (s *RemoveUserGroupMembersRequest) SetAppToken(v string) *RemoveUserGroupMembersRequest {
	s.AppToken = &v
	return s
}

func (s *RemoveUserGroupMembersRequest) SetOpsUserCode(v string) *RemoveUserGroupMembersRequest {
	s.OpsUserCode = &v
	return s
}

func (s *RemoveUserGroupMembersRequest) SetOpsUserId(v string) *RemoveUserGroupMembersRequest {
	s.OpsUserId = &v
	return s
}

func (s *RemoveUserGroupMembersRequest) SetTenantId(v string) *RemoveUserGroupMembersRequest {
	s.TenantId = &v
	return s
}

func (s *RemoveUserGroupMembersRequest) SetUserGroupId(v string) *RemoveUserGroupMembersRequest {
	s.UserGroupId = &v
	return s
}

func (s *RemoveUserGroupMembersRequest) SetUserIds(v []*int64) *RemoveUserGroupMembersRequest {
	s.UserIds = v
	return s
}

func (s *RemoveUserGroupMembersRequest) Validate() error {
	return dara.Validate(s)
}
