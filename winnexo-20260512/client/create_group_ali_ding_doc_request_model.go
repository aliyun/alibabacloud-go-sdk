// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGroupAliDingDocRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *CreateGroupAliDingDocRequest
	GetAppToken() *string
	SetDescription(v string) *CreateGroupAliDingDocRequest
	GetDescription() *string
	SetDirectoryId(v string) *CreateGroupAliDingDocRequest
	GetDirectoryId() *string
	SetFilePublicUrl(v string) *CreateGroupAliDingDocRequest
	GetFilePublicUrl() *string
	SetGroupId(v string) *CreateGroupAliDingDocRequest
	GetGroupId() *string
	SetName(v string) *CreateGroupAliDingDocRequest
	GetName() *string
	SetOpsUserCode(v string) *CreateGroupAliDingDocRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *CreateGroupAliDingDocRequest
	GetOpsUserId() *string
	SetSourceTags(v string) *CreateGroupAliDingDocRequest
	GetSourceTags() *string
	SetTenantId(v string) *CreateGroupAliDingDocRequest
	GetTenantId() *string
}

type CreateGroupAliDingDocRequest struct {
	// 平台应用 Token（控制台签发）。传入后将以 opsUserId 或 opsUserCode 指定的用户身份执行
	//
	// example:
	//
	// wn_app_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
	AppToken *string `json:"appToken,omitempty" xml:"appToken,omitempty"`
	// The description of the AI assistant.
	//
	// example:
	//
	// Sample description
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The folder ID.
	//
	// example:
	//
	// exampleDirectoryId
	DirectoryId *string `json:"directoryId,omitempty" xml:"directoryId,omitempty"`
	// The publicly accessible URL of the Alibaba DingTalk online document.
	//
	// This parameter is required.
	//
	// example:
	//
	// https://example.com/winnexo/resource
	FilePublicUrl *string `json:"filePublicUrl,omitempty" xml:"filePublicUrl,omitempty"`
	// The project group ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// exampleGroupId
	GroupId *string `json:"groupId,omitempty" xml:"groupId,omitempty"`
	// The image name.
	//
	// This parameter is required.
	//
	// example:
	//
	// user_paswd_104
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
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
	// The resource tags. This parameter is optional. Specify a JSON string list, such as ["tagA","tagB"].
	//
	// example:
	//
	// ["Customer","GroupChat"]
	SourceTags *string `json:"sourceTags,omitempty" xml:"sourceTags,omitempty"`
	// The tenant ID.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s CreateGroupAliDingDocRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateGroupAliDingDocRequest) GoString() string {
	return s.String()
}

func (s *CreateGroupAliDingDocRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *CreateGroupAliDingDocRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateGroupAliDingDocRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *CreateGroupAliDingDocRequest) GetFilePublicUrl() *string {
	return s.FilePublicUrl
}

func (s *CreateGroupAliDingDocRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *CreateGroupAliDingDocRequest) GetName() *string {
	return s.Name
}

func (s *CreateGroupAliDingDocRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *CreateGroupAliDingDocRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *CreateGroupAliDingDocRequest) GetSourceTags() *string {
	return s.SourceTags
}

func (s *CreateGroupAliDingDocRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *CreateGroupAliDingDocRequest) SetAppToken(v string) *CreateGroupAliDingDocRequest {
	s.AppToken = &v
	return s
}

func (s *CreateGroupAliDingDocRequest) SetDescription(v string) *CreateGroupAliDingDocRequest {
	s.Description = &v
	return s
}

func (s *CreateGroupAliDingDocRequest) SetDirectoryId(v string) *CreateGroupAliDingDocRequest {
	s.DirectoryId = &v
	return s
}

func (s *CreateGroupAliDingDocRequest) SetFilePublicUrl(v string) *CreateGroupAliDingDocRequest {
	s.FilePublicUrl = &v
	return s
}

func (s *CreateGroupAliDingDocRequest) SetGroupId(v string) *CreateGroupAliDingDocRequest {
	s.GroupId = &v
	return s
}

func (s *CreateGroupAliDingDocRequest) SetName(v string) *CreateGroupAliDingDocRequest {
	s.Name = &v
	return s
}

func (s *CreateGroupAliDingDocRequest) SetOpsUserCode(v string) *CreateGroupAliDingDocRequest {
	s.OpsUserCode = &v
	return s
}

func (s *CreateGroupAliDingDocRequest) SetOpsUserId(v string) *CreateGroupAliDingDocRequest {
	s.OpsUserId = &v
	return s
}

func (s *CreateGroupAliDingDocRequest) SetSourceTags(v string) *CreateGroupAliDingDocRequest {
	s.SourceTags = &v
	return s
}

func (s *CreateGroupAliDingDocRequest) SetTenantId(v string) *CreateGroupAliDingDocRequest {
	s.TenantId = &v
	return s
}

func (s *CreateGroupAliDingDocRequest) Validate() error {
	return dara.Validate(s)
}
