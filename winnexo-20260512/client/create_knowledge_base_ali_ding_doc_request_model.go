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
	// The description of the alias.
	//
	// example:
	//
	// created by eventbridge
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The directory ID.
	//
	// example:
	//
	// exampleDirectoryId
	DirectoryId *string `json:"directoryId,omitempty" xml:"directoryId,omitempty"`
	// The publicly accessible URL of the AliDing online document.
	//
	// This parameter is required.
	//
	// example:
	//
	// https://example.com/winnexo/resource
	FilePublicUrl *string `json:"filePublicUrl,omitempty" xml:"filePublicUrl,omitempty"`
	// Not supported. This parameter is ignored.
	//
	// example:
	//
	// exampleKnowledgeId
	KnowledgeId *string `json:"knowledgeId,omitempty" xml:"knowledgeId,omitempty"`
	// The name.
	//
	// This parameter is required.
	//
	// example:
	//
	// p-toolset-4dd55d81-f98f-4806-8596-43f1c95f1ff0
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The digital employee name (operating object name, optional).
	//
	// example:
	//
	// string_value
	OperatingObjectName *string `json:"operatingObjectName,omitempty" xml:"operatingObjectName,omitempty"`
	// The resource tags (optional, a JSON string list, such as ["tagA","tagB"]).
	//
	// example:
	//
	// string_value
	SourceTags *string `json:"sourceTags,omitempty" xml:"sourceTags,omitempty"`
	// The tenant ID.
	//
	// example:
	//
	// 520539530998273
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
