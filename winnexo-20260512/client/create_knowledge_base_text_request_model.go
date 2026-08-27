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
