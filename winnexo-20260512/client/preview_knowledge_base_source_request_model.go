// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPreviewKnowledgeBaseSourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *PreviewKnowledgeBaseSourceRequest
	GetAppToken() *string
	SetOpsUserCode(v string) *PreviewKnowledgeBaseSourceRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *PreviewKnowledgeBaseSourceRequest
	GetOpsUserId() *string
	SetSourceId(v string) *PreviewKnowledgeBaseSourceRequest
	GetSourceId() *string
	SetTenantId(v string) *PreviewKnowledgeBaseSourceRequest
	GetTenantId() *string
}

type PreviewKnowledgeBaseSourceRequest struct {
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
	// The data source ID, which is unique within the tenant.
	//
	// This parameter is required.
	//
	// example:
	//
	// exampleSourceId
	SourceId *string `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	// The tenant ID. This is a common parameter. In winnexo-cli, pass it explicitly with --tenant-id.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s PreviewKnowledgeBaseSourceRequest) String() string {
	return dara.Prettify(s)
}

func (s PreviewKnowledgeBaseSourceRequest) GoString() string {
	return s.String()
}

func (s *PreviewKnowledgeBaseSourceRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *PreviewKnowledgeBaseSourceRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *PreviewKnowledgeBaseSourceRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *PreviewKnowledgeBaseSourceRequest) GetSourceId() *string {
	return s.SourceId
}

func (s *PreviewKnowledgeBaseSourceRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *PreviewKnowledgeBaseSourceRequest) SetAppToken(v string) *PreviewKnowledgeBaseSourceRequest {
	s.AppToken = &v
	return s
}

func (s *PreviewKnowledgeBaseSourceRequest) SetOpsUserCode(v string) *PreviewKnowledgeBaseSourceRequest {
	s.OpsUserCode = &v
	return s
}

func (s *PreviewKnowledgeBaseSourceRequest) SetOpsUserId(v string) *PreviewKnowledgeBaseSourceRequest {
	s.OpsUserId = &v
	return s
}

func (s *PreviewKnowledgeBaseSourceRequest) SetSourceId(v string) *PreviewKnowledgeBaseSourceRequest {
	s.SourceId = &v
	return s
}

func (s *PreviewKnowledgeBaseSourceRequest) SetTenantId(v string) *PreviewKnowledgeBaseSourceRequest {
	s.TenantId = &v
	return s
}

func (s *PreviewKnowledgeBaseSourceRequest) Validate() error {
	return dara.Validate(s)
}
