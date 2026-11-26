// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateKnowledgeBaseSourceContentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *UpdateKnowledgeBaseSourceContentRequest
	GetAppToken() *string
	SetContent(v string) *UpdateKnowledgeBaseSourceContentRequest
	GetContent() *string
	SetForceSync(v bool) *UpdateKnowledgeBaseSourceContentRequest
	GetForceSync() *bool
	SetOpsUserCode(v string) *UpdateKnowledgeBaseSourceContentRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *UpdateKnowledgeBaseSourceContentRequest
	GetOpsUserId() *string
	SetSourceId(v string) *UpdateKnowledgeBaseSourceContentRequest
	GetSourceId() *string
	SetTenantId(v string) *UpdateKnowledgeBaseSourceContentRequest
	GetTenantId() *string
}

type UpdateKnowledgeBaseSourceContentRequest struct {
	// 平台应用 Token（控制台签发）。传入后将以 opsUserId 或 opsUserCode 指定的用户身份执行
	//
	// example:
	//
	// wn_app_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
	AppToken *string `json:"appToken,omitempty" xml:"appToken,omitempty"`
	// Specifies whether the operation is successful. A value of true indicates success.
	//
	// This parameter is required.
	//
	// example:
	//
	// Sample content
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// Specifies whether to force synchronous processing.
	//
	// example:
	//
	// false
	ForceSync *bool `json:"forceSync,omitempty" xml:"forceSync,omitempty"`
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
	// The source ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// exampleSourceId
	SourceId *string `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	// The tenant ID that takes effect.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s UpdateKnowledgeBaseSourceContentRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateKnowledgeBaseSourceContentRequest) GoString() string {
	return s.String()
}

func (s *UpdateKnowledgeBaseSourceContentRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *UpdateKnowledgeBaseSourceContentRequest) GetContent() *string {
	return s.Content
}

func (s *UpdateKnowledgeBaseSourceContentRequest) GetForceSync() *bool {
	return s.ForceSync
}

func (s *UpdateKnowledgeBaseSourceContentRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *UpdateKnowledgeBaseSourceContentRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *UpdateKnowledgeBaseSourceContentRequest) GetSourceId() *string {
	return s.SourceId
}

func (s *UpdateKnowledgeBaseSourceContentRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *UpdateKnowledgeBaseSourceContentRequest) SetAppToken(v string) *UpdateKnowledgeBaseSourceContentRequest {
	s.AppToken = &v
	return s
}

func (s *UpdateKnowledgeBaseSourceContentRequest) SetContent(v string) *UpdateKnowledgeBaseSourceContentRequest {
	s.Content = &v
	return s
}

func (s *UpdateKnowledgeBaseSourceContentRequest) SetForceSync(v bool) *UpdateKnowledgeBaseSourceContentRequest {
	s.ForceSync = &v
	return s
}

func (s *UpdateKnowledgeBaseSourceContentRequest) SetOpsUserCode(v string) *UpdateKnowledgeBaseSourceContentRequest {
	s.OpsUserCode = &v
	return s
}

func (s *UpdateKnowledgeBaseSourceContentRequest) SetOpsUserId(v string) *UpdateKnowledgeBaseSourceContentRequest {
	s.OpsUserId = &v
	return s
}

func (s *UpdateKnowledgeBaseSourceContentRequest) SetSourceId(v string) *UpdateKnowledgeBaseSourceContentRequest {
	s.SourceId = &v
	return s
}

func (s *UpdateKnowledgeBaseSourceContentRequest) SetTenantId(v string) *UpdateKnowledgeBaseSourceContentRequest {
	s.TenantId = &v
	return s
}

func (s *UpdateKnowledgeBaseSourceContentRequest) Validate() error {
	return dara.Validate(s)
}
