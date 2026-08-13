// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateKnowledgeBaseTextRequest interface {
	dara.Model
	String() string
	GoString() string
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
	SetSourceTags(v string) *CreateKnowledgeBaseTextRequest
	GetSourceTags() *string
	SetTenantId(v string) *CreateKnowledgeBaseTextRequest
	GetTenantId() *string
	SetTextContent(v string) *CreateKnowledgeBaseTextRequest
	GetTextContent() *string
}

type CreateKnowledgeBaseTextRequest struct {
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
	// 知识库 ID（可选，透传给 document_agent）
	//
	// example:
	//
	// exampleKnowledgeId
	KnowledgeId *string `json:"knowledgeId,omitempty" xml:"knowledgeId,omitempty"`
	// 资源显示名称
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
	// 纯文本正文（必填）
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

func (s *CreateKnowledgeBaseTextRequest) GetSourceTags() *string {
	return s.SourceTags
}

func (s *CreateKnowledgeBaseTextRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *CreateKnowledgeBaseTextRequest) GetTextContent() *string {
	return s.TextContent
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
