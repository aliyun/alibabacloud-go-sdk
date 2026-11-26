// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReplaceKnowledgeBaseAdminsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *ReplaceKnowledgeBaseAdminsRequest
	GetAppToken() *string
	SetDirectoryId(v string) *ReplaceKnowledgeBaseAdminsRequest
	GetDirectoryId() *string
	SetOpsUserCode(v string) *ReplaceKnowledgeBaseAdminsRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *ReplaceKnowledgeBaseAdminsRequest
	GetOpsUserId() *string
	SetTenantId(v string) *ReplaceKnowledgeBaseAdminsRequest
	GetTenantId() *string
	SetUserGroupIds(v []*string) *ReplaceKnowledgeBaseAdminsRequest
	GetUserGroupIds() []*string
	SetUserIds(v []*int64) *ReplaceKnowledgeBaseAdminsRequest
	GetUserIds() []*int64
}

type ReplaceKnowledgeBaseAdminsRequest struct {
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
	UserGroupIds []*string `json:"userGroupIds,omitempty" xml:"userGroupIds,omitempty" type:"Repeated"`
	// 直接管理员的平台用户 ID 列表；默认空列表表示清空可变用户管理员，普通库创建者保留
	//
	// example:
	//
	// [10001]
	UserIds []*int64 `json:"userIds,omitempty" xml:"userIds,omitempty" type:"Repeated"`
}

func (s ReplaceKnowledgeBaseAdminsRequest) String() string {
	return dara.Prettify(s)
}

func (s ReplaceKnowledgeBaseAdminsRequest) GoString() string {
	return s.String()
}

func (s *ReplaceKnowledgeBaseAdminsRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *ReplaceKnowledgeBaseAdminsRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *ReplaceKnowledgeBaseAdminsRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *ReplaceKnowledgeBaseAdminsRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *ReplaceKnowledgeBaseAdminsRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ReplaceKnowledgeBaseAdminsRequest) GetUserGroupIds() []*string {
	return s.UserGroupIds
}

func (s *ReplaceKnowledgeBaseAdminsRequest) GetUserIds() []*int64 {
	return s.UserIds
}

func (s *ReplaceKnowledgeBaseAdminsRequest) SetAppToken(v string) *ReplaceKnowledgeBaseAdminsRequest {
	s.AppToken = &v
	return s
}

func (s *ReplaceKnowledgeBaseAdminsRequest) SetDirectoryId(v string) *ReplaceKnowledgeBaseAdminsRequest {
	s.DirectoryId = &v
	return s
}

func (s *ReplaceKnowledgeBaseAdminsRequest) SetOpsUserCode(v string) *ReplaceKnowledgeBaseAdminsRequest {
	s.OpsUserCode = &v
	return s
}

func (s *ReplaceKnowledgeBaseAdminsRequest) SetOpsUserId(v string) *ReplaceKnowledgeBaseAdminsRequest {
	s.OpsUserId = &v
	return s
}

func (s *ReplaceKnowledgeBaseAdminsRequest) SetTenantId(v string) *ReplaceKnowledgeBaseAdminsRequest {
	s.TenantId = &v
	return s
}

func (s *ReplaceKnowledgeBaseAdminsRequest) SetUserGroupIds(v []*string) *ReplaceKnowledgeBaseAdminsRequest {
	s.UserGroupIds = v
	return s
}

func (s *ReplaceKnowledgeBaseAdminsRequest) SetUserIds(v []*int64) *ReplaceKnowledgeBaseAdminsRequest {
	s.UserIds = v
	return s
}

func (s *ReplaceKnowledgeBaseAdminsRequest) Validate() error {
	return dara.Validate(s)
}
