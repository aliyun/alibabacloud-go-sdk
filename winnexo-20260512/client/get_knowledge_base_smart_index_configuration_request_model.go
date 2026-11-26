// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetKnowledgeBaseSmartIndexConfigurationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *GetKnowledgeBaseSmartIndexConfigurationRequest
	GetAppToken() *string
	SetDirectoryId(v string) *GetKnowledgeBaseSmartIndexConfigurationRequest
	GetDirectoryId() *string
	SetOpsUserCode(v string) *GetKnowledgeBaseSmartIndexConfigurationRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *GetKnowledgeBaseSmartIndexConfigurationRequest
	GetOpsUserId() *string
	SetTenantId(v string) *GetKnowledgeBaseSmartIndexConfigurationRequest
	GetTenantId() *string
}

type GetKnowledgeBaseSmartIndexConfigurationRequest struct {
	// 平台应用 Token（控制台签发）。传入后将以 opsUserId 或 opsUserCode 指定的用户身份执行
	//
	// example:
	//
	// wn_app_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
	AppToken *string `json:"appToken,omitempty" xml:"appToken,omitempty"`
	// 企业知识库根目录业务 ID
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

func (s GetKnowledgeBaseSmartIndexConfigurationRequest) String() string {
	return dara.Prettify(s)
}

func (s GetKnowledgeBaseSmartIndexConfigurationRequest) GoString() string {
	return s.String()
}

func (s *GetKnowledgeBaseSmartIndexConfigurationRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *GetKnowledgeBaseSmartIndexConfigurationRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *GetKnowledgeBaseSmartIndexConfigurationRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *GetKnowledgeBaseSmartIndexConfigurationRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *GetKnowledgeBaseSmartIndexConfigurationRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *GetKnowledgeBaseSmartIndexConfigurationRequest) SetAppToken(v string) *GetKnowledgeBaseSmartIndexConfigurationRequest {
	s.AppToken = &v
	return s
}

func (s *GetKnowledgeBaseSmartIndexConfigurationRequest) SetDirectoryId(v string) *GetKnowledgeBaseSmartIndexConfigurationRequest {
	s.DirectoryId = &v
	return s
}

func (s *GetKnowledgeBaseSmartIndexConfigurationRequest) SetOpsUserCode(v string) *GetKnowledgeBaseSmartIndexConfigurationRequest {
	s.OpsUserCode = &v
	return s
}

func (s *GetKnowledgeBaseSmartIndexConfigurationRequest) SetOpsUserId(v string) *GetKnowledgeBaseSmartIndexConfigurationRequest {
	s.OpsUserId = &v
	return s
}

func (s *GetKnowledgeBaseSmartIndexConfigurationRequest) SetTenantId(v string) *GetKnowledgeBaseSmartIndexConfigurationRequest {
	s.TenantId = &v
	return s
}

func (s *GetKnowledgeBaseSmartIndexConfigurationRequest) Validate() error {
	return dara.Validate(s)
}
