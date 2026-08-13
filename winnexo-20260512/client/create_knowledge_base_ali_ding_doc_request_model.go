// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateKnowledgeBaseAliDingDocRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateKnowledgeBaseAliDingDocRequest
	GetDescription() *string
	SetDirectoryId(v string) *CreateKnowledgeBaseAliDingDocRequest
	GetDirectoryId() *string
	SetFilePublicUrl(v string) *CreateKnowledgeBaseAliDingDocRequest
	GetFilePublicUrl() *string
	SetKnowledgeId(v string) *CreateKnowledgeBaseAliDingDocRequest
	GetKnowledgeId() *string
	SetName(v string) *CreateKnowledgeBaseAliDingDocRequest
	GetName() *string
	SetOperatingObjectName(v string) *CreateKnowledgeBaseAliDingDocRequest
	GetOperatingObjectName() *string
	SetSourceTags(v string) *CreateKnowledgeBaseAliDingDocRequest
	GetSourceTags() *string
	SetTenantId(v string) *CreateKnowledgeBaseAliDingDocRequest
	GetTenantId() *string
}

type CreateKnowledgeBaseAliDingDocRequest struct {
	// 资源描述（可选）
	//
	// example:
	//
	// 示例描述
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 目标企业知识库目录 ID；不传时自动绑定到当前数字员工默认根目录，传入时必须是当前租户下已有的企业知识库目录
	//
	// example:
	//
	// exampleDirectoryId
	DirectoryId *string `json:"directoryId,omitempty" xml:"directoryId,omitempty"`
	// 阿里钉在线文档的可公开访问 URL
	//
	// This parameter is required.
	//
	// example:
	//
	// https://example.com/winnexo/resource
	FilePublicUrl *string `json:"filePublicUrl,omitempty" xml:"filePublicUrl,omitempty"`
	// 知识库 ID（可选，透传给 document_agent）
	//
	// example:
	//
	// exampleKnowledgeId
	KnowledgeId *string `json:"knowledgeId,omitempty" xml:"knowledgeId,omitempty"`
	// 资源显示名称（建议传入钉钉文档标题）
	//
	// This parameter is required.
	//
	// example:
	//
	// 示例名称.pdf
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 数字员工名称（运营对象 name，可选）
	//
	// example:
	//
	// string_value
	OperatingObjectName *string `json:"operatingObjectName,omitempty" xml:"operatingObjectName,omitempty"`
	// 资源标签（可选，JSON 字符串列表，如 ["tagA","tagB"]）
	//
	// example:
	//
	// string_value
	SourceTags *string `json:"sourceTags,omitempty" xml:"sourceTags,omitempty"`
	// 租户ID，公共参数，缺省时使用调用方默认租户
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s CreateKnowledgeBaseAliDingDocRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateKnowledgeBaseAliDingDocRequest) GoString() string {
	return s.String()
}

func (s *CreateKnowledgeBaseAliDingDocRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateKnowledgeBaseAliDingDocRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *CreateKnowledgeBaseAliDingDocRequest) GetFilePublicUrl() *string {
	return s.FilePublicUrl
}

func (s *CreateKnowledgeBaseAliDingDocRequest) GetKnowledgeId() *string {
	return s.KnowledgeId
}

func (s *CreateKnowledgeBaseAliDingDocRequest) GetName() *string {
	return s.Name
}

func (s *CreateKnowledgeBaseAliDingDocRequest) GetOperatingObjectName() *string {
	return s.OperatingObjectName
}

func (s *CreateKnowledgeBaseAliDingDocRequest) GetSourceTags() *string {
	return s.SourceTags
}

func (s *CreateKnowledgeBaseAliDingDocRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *CreateKnowledgeBaseAliDingDocRequest) SetDescription(v string) *CreateKnowledgeBaseAliDingDocRequest {
	s.Description = &v
	return s
}

func (s *CreateKnowledgeBaseAliDingDocRequest) SetDirectoryId(v string) *CreateKnowledgeBaseAliDingDocRequest {
	s.DirectoryId = &v
	return s
}

func (s *CreateKnowledgeBaseAliDingDocRequest) SetFilePublicUrl(v string) *CreateKnowledgeBaseAliDingDocRequest {
	s.FilePublicUrl = &v
	return s
}

func (s *CreateKnowledgeBaseAliDingDocRequest) SetKnowledgeId(v string) *CreateKnowledgeBaseAliDingDocRequest {
	s.KnowledgeId = &v
	return s
}

func (s *CreateKnowledgeBaseAliDingDocRequest) SetName(v string) *CreateKnowledgeBaseAliDingDocRequest {
	s.Name = &v
	return s
}

func (s *CreateKnowledgeBaseAliDingDocRequest) SetOperatingObjectName(v string) *CreateKnowledgeBaseAliDingDocRequest {
	s.OperatingObjectName = &v
	return s
}

func (s *CreateKnowledgeBaseAliDingDocRequest) SetSourceTags(v string) *CreateKnowledgeBaseAliDingDocRequest {
	s.SourceTags = &v
	return s
}

func (s *CreateKnowledgeBaseAliDingDocRequest) SetTenantId(v string) *CreateKnowledgeBaseAliDingDocRequest {
	s.TenantId = &v
	return s
}

func (s *CreateKnowledgeBaseAliDingDocRequest) Validate() error {
	return dara.Validate(s)
}
