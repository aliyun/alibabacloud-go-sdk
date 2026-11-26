// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetKnowledgeBaseConfigurationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *GetKnowledgeBaseConfigurationRequest
	GetAppToken() *string
	SetDirectoryId(v string) *GetKnowledgeBaseConfigurationRequest
	GetDirectoryId() *string
	SetOpsUserCode(v string) *GetKnowledgeBaseConfigurationRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *GetKnowledgeBaseConfigurationRequest
	GetOpsUserId() *string
	SetTenantId(v string) *GetKnowledgeBaseConfigurationRequest
	GetTenantId() *string
}

type GetKnowledgeBaseConfigurationRequest struct {
	// 平台应用 Token（控制台签发）。传入后将以 opsUserId 或 opsUserCode 指定的用户身份执行
	//
	// example:
	//
	// wn_app_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
	AppToken *string `json:"appToken,omitempty" xml:"appToken,omitempty"`
	// 普通企业知识库根目录 ID
	//
	// This parameter is required.
	//
	// example:
	//
	// kb-example
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
}

func (s GetKnowledgeBaseConfigurationRequest) String() string {
	return dara.Prettify(s)
}

func (s GetKnowledgeBaseConfigurationRequest) GoString() string {
	return s.String()
}

func (s *GetKnowledgeBaseConfigurationRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *GetKnowledgeBaseConfigurationRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *GetKnowledgeBaseConfigurationRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *GetKnowledgeBaseConfigurationRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *GetKnowledgeBaseConfigurationRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *GetKnowledgeBaseConfigurationRequest) SetAppToken(v string) *GetKnowledgeBaseConfigurationRequest {
	s.AppToken = &v
	return s
}

func (s *GetKnowledgeBaseConfigurationRequest) SetDirectoryId(v string) *GetKnowledgeBaseConfigurationRequest {
	s.DirectoryId = &v
	return s
}

func (s *GetKnowledgeBaseConfigurationRequest) SetOpsUserCode(v string) *GetKnowledgeBaseConfigurationRequest {
	s.OpsUserCode = &v
	return s
}

func (s *GetKnowledgeBaseConfigurationRequest) SetOpsUserId(v string) *GetKnowledgeBaseConfigurationRequest {
	s.OpsUserId = &v
	return s
}

func (s *GetKnowledgeBaseConfigurationRequest) SetTenantId(v string) *GetKnowledgeBaseConfigurationRequest {
	s.TenantId = &v
	return s
}

func (s *GetKnowledgeBaseConfigurationRequest) Validate() error {
	return dara.Validate(s)
}
