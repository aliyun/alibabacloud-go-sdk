// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateKnowledgeBaseTextRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *CreateKnowledgeBaseTextRequest
	GetAppToken() *string
	SetDescription(v string) *CreateKnowledgeBaseTextRequest
	GetDescription() *string
	SetDirectoryId(v string) *CreateKnowledgeBaseTextRequest
	GetDirectoryId() *string
	SetKnowledgeId(v string) *CreateKnowledgeBaseTextRequest
	GetKnowledgeId() *string
	SetName(v string) *CreateKnowledgeBaseTextRequest
	GetName() *string
	SetOperatingObjectName(v string) *CreateKnowledgeBaseTextRequest
	GetOperatingObjectName() *string
	SetOpsUserCode(v string) *CreateKnowledgeBaseTextRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *CreateKnowledgeBaseTextRequest
	GetOpsUserId() *string
	SetSourceTags(v string) *CreateKnowledgeBaseTextRequest
	GetSourceTags() *string
	SetTenantId(v string) *CreateKnowledgeBaseTextRequest
	GetTenantId() *string
	SetTextContent(v string) *CreateKnowledgeBaseTextRequest
	GetTextContent() *string
}

type CreateKnowledgeBaseTextRequest struct {
	// 平台应用 Token（控制台签发）。传入后将以 opsUserId 或 opsUserCode 指定的用户身份执行
	//
	// example:
	//
	// wn_app_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
	AppToken *string `json:"appToken,omitempty" xml:"appToken,omitempty"`
	// The description of the alias.
	//
	// example:
	//
	// InterviewMaster operations and health check service
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The folder ID.
	//
	// example:
	//
	// exampleDirectoryId
	DirectoryId *string `json:"directoryId,omitempty" xml:"directoryId,omitempty"`
	// Not supported. Ignore this parameter.
	//
	// example:
	//
	// exampleKnowledgeId
	KnowledgeId *string `json:"knowledgeId,omitempty" xml:"knowledgeId,omitempty"`
	// The image name.
	//
	// This parameter is required.
	//
	// example:
	//
	// KL_tongyici
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The name of the operating object.
	//
	// example:
	//
	// string_value
	OperatingObjectName *string `json:"operatingObjectName,omitempty" xml:"operatingObjectName,omitempty"`
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
	// The source tags.
	//
	// example:
	//
	// string_value
	SourceTags *string `json:"sourceTags,omitempty" xml:"sourceTags,omitempty"`
	// The tenant ID.
	//
	// example:
	//
	// 10001
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// The message content for text messages.
	//
	// This parameter is required.
	//
	// example:
	//
	// string_value
	TextContent *string `json:"textContent,omitempty" xml:"textContent,omitempty"`
}

func (s CreateKnowledgeBaseTextRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateKnowledgeBaseTextRequest) GoString() string {
	return s.String()
}

func (s *CreateKnowledgeBaseTextRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *CreateKnowledgeBaseTextRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateKnowledgeBaseTextRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *CreateKnowledgeBaseTextRequest) GetKnowledgeId() *string {
	return s.KnowledgeId
}

func (s *CreateKnowledgeBaseTextRequest) GetName() *string {
	return s.Name
}

func (s *CreateKnowledgeBaseTextRequest) GetOperatingObjectName() *string {
	return s.OperatingObjectName
}

func (s *CreateKnowledgeBaseTextRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *CreateKnowledgeBaseTextRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *CreateKnowledgeBaseTextRequest) GetSourceTags() *string {
	return s.SourceTags
}

func (s *CreateKnowledgeBaseTextRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *CreateKnowledgeBaseTextRequest) GetTextContent() *string {
	return s.TextContent
}

func (s *CreateKnowledgeBaseTextRequest) SetAppToken(v string) *CreateKnowledgeBaseTextRequest {
	s.AppToken = &v
	return s
}

func (s *CreateKnowledgeBaseTextRequest) SetDescription(v string) *CreateKnowledgeBaseTextRequest {
	s.Description = &v
	return s
}

func (s *CreateKnowledgeBaseTextRequest) SetDirectoryId(v string) *CreateKnowledgeBaseTextRequest {
	s.DirectoryId = &v
	return s
}

func (s *CreateKnowledgeBaseTextRequest) SetKnowledgeId(v string) *CreateKnowledgeBaseTextRequest {
	s.KnowledgeId = &v
	return s
}

func (s *CreateKnowledgeBaseTextRequest) SetName(v string) *CreateKnowledgeBaseTextRequest {
	s.Name = &v
	return s
}

func (s *CreateKnowledgeBaseTextRequest) SetOperatingObjectName(v string) *CreateKnowledgeBaseTextRequest {
	s.OperatingObjectName = &v
	return s
}

func (s *CreateKnowledgeBaseTextRequest) SetOpsUserCode(v string) *CreateKnowledgeBaseTextRequest {
	s.OpsUserCode = &v
	return s
}

func (s *CreateKnowledgeBaseTextRequest) SetOpsUserId(v string) *CreateKnowledgeBaseTextRequest {
	s.OpsUserId = &v
	return s
}

func (s *CreateKnowledgeBaseTextRequest) SetSourceTags(v string) *CreateKnowledgeBaseTextRequest {
	s.SourceTags = &v
	return s
}

func (s *CreateKnowledgeBaseTextRequest) SetTenantId(v string) *CreateKnowledgeBaseTextRequest {
	s.TenantId = &v
	return s
}

func (s *CreateKnowledgeBaseTextRequest) SetTextContent(v string) *CreateKnowledgeBaseTextRequest {
	s.TextContent = &v
	return s
}

func (s *CreateKnowledgeBaseTextRequest) Validate() error {
	return dara.Validate(s)
}
