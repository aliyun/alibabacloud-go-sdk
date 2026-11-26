// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReplaceKnowledgeBaseAdminsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *ReplaceKnowledgeBaseAdminsShrinkRequest
	GetAppToken() *string
	SetDirectoryId(v string) *ReplaceKnowledgeBaseAdminsShrinkRequest
	GetDirectoryId() *string
	SetOpsUserCode(v string) *ReplaceKnowledgeBaseAdminsShrinkRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *ReplaceKnowledgeBaseAdminsShrinkRequest
	GetOpsUserId() *string
	SetTenantId(v string) *ReplaceKnowledgeBaseAdminsShrinkRequest
	GetTenantId() *string
	SetUserGroupIdsShrink(v string) *ReplaceKnowledgeBaseAdminsShrinkRequest
	GetUserGroupIdsShrink() *string
	SetUserIdsShrink(v string) *ReplaceKnowledgeBaseAdminsShrinkRequest
	GetUserIdsShrink() *string
}

type ReplaceKnowledgeBaseAdminsShrinkRequest struct {
	// 平台应用 Token（控制台签发）。传入后将以 opsUserId 或 opsUserCode 指定的用户身份执行
	//
	// example:
	//
	// wn_app_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
	AppToken *string `json:"appToken,omitempty" xml:"appToken,omitempty"`
	// 企业知识库或内置帮助知识库根目录业务 ID
	//
	// This parameter is required.
	//
	// example:
	//
	// kb_001
	DirectoryId *string `json:"directoryId,omitempty" xml:"directoryId,omitempty"`
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
	// 租户ID，公共参数；winnexo-cli 通过 --tenant-id 显式传入
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// 管理员用户组 ID 列表；默认空列表表示清空用户组管理员
	//
	// example:
	//
	// ["group-1"]
	UserGroupIdsShrink *string `json:"userGroupIds,omitempty" xml:"userGroupIds,omitempty"`
	// 直接管理员的平台用户 ID 列表；默认空列表表示清空可变用户管理员，普通库创建者保留
	//
	// example:
	//
	// [10001]
	UserIdsShrink *string `json:"userIds,omitempty" xml:"userIds,omitempty"`
}

func (s ReplaceKnowledgeBaseAdminsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ReplaceKnowledgeBaseAdminsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ReplaceKnowledgeBaseAdminsShrinkRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *ReplaceKnowledgeBaseAdminsShrinkRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *ReplaceKnowledgeBaseAdminsShrinkRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *ReplaceKnowledgeBaseAdminsShrinkRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *ReplaceKnowledgeBaseAdminsShrinkRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ReplaceKnowledgeBaseAdminsShrinkRequest) GetUserGroupIdsShrink() *string {
	return s.UserGroupIdsShrink
}

func (s *ReplaceKnowledgeBaseAdminsShrinkRequest) GetUserIdsShrink() *string {
	return s.UserIdsShrink
}

func (s *ReplaceKnowledgeBaseAdminsShrinkRequest) SetAppToken(v string) *ReplaceKnowledgeBaseAdminsShrinkRequest {
	s.AppToken = &v
	return s
}

func (s *ReplaceKnowledgeBaseAdminsShrinkRequest) SetDirectoryId(v string) *ReplaceKnowledgeBaseAdminsShrinkRequest {
	s.DirectoryId = &v
	return s
}

func (s *ReplaceKnowledgeBaseAdminsShrinkRequest) SetOpsUserCode(v string) *ReplaceKnowledgeBaseAdminsShrinkRequest {
	s.OpsUserCode = &v
	return s
}

func (s *ReplaceKnowledgeBaseAdminsShrinkRequest) SetOpsUserId(v string) *ReplaceKnowledgeBaseAdminsShrinkRequest {
	s.OpsUserId = &v
	return s
}

func (s *ReplaceKnowledgeBaseAdminsShrinkRequest) SetTenantId(v string) *ReplaceKnowledgeBaseAdminsShrinkRequest {
	s.TenantId = &v
	return s
}

func (s *ReplaceKnowledgeBaseAdminsShrinkRequest) SetUserGroupIdsShrink(v string) *ReplaceKnowledgeBaseAdminsShrinkRequest {
	s.UserGroupIdsShrink = &v
	return s
}

func (s *ReplaceKnowledgeBaseAdminsShrinkRequest) SetUserIdsShrink(v string) *ReplaceKnowledgeBaseAdminsShrinkRequest {
	s.UserIdsShrink = &v
	return s
}

func (s *ReplaceKnowledgeBaseAdminsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
