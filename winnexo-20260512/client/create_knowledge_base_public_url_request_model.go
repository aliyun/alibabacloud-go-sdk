// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateKnowledgeBasePublicUrlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateKnowledgeBasePublicUrlRequest
	GetDescription() *string
	SetDirectoryId(v string) *CreateKnowledgeBasePublicUrlRequest
	GetDirectoryId() *string
	SetName(v string) *CreateKnowledgeBasePublicUrlRequest
	GetName() *string
	SetNotes(v string) *CreateKnowledgeBasePublicUrlRequest
	GetNotes() *string
	SetOperatingObjectName(v string) *CreateKnowledgeBasePublicUrlRequest
	GetOperatingObjectName() *string
	SetOriginalUrl(v string) *CreateKnowledgeBasePublicUrlRequest
	GetOriginalUrl() *string
	SetSourceTags(v string) *CreateKnowledgeBasePublicUrlRequest
	GetSourceTags() *string
	SetTenantId(v string) *CreateKnowledgeBasePublicUrlRequest
	GetTenantId() *string
}

type CreateKnowledgeBasePublicUrlRequest struct {
	// The resource description.
	//
	// example:
	//
	// Project design document
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The ID of the destination folder in the enterprise knowledge base. This parameter is required. You must have knowledge base management permissions on the knowledge base.
	//
	// This parameter is required.
	//
	// example:
	//
	// dir_tenant_child
	DirectoryId *string `json:"directoryId,omitempty" xml:"directoryId,omitempty"`
	// The resource name. If not specified, the URL is used.
	//
	// example:
	//
	// Project Plan
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The analysis instruction.
	//
	// example:
	//
	// Extract decisions and to-do items
	Notes *string `json:"notes,omitempty" xml:"notes,omitempty"`
	// The name of the operating object.
	//
	// example:
	//
	// R&D Assistant
	OperatingObjectName *string `json:"operatingObjectName,omitempty" xml:"operatingObjectName,omitempty"`
	// The URL of the public HTTP/HTTPS web page.
	//
	// This parameter is required.
	//
	// example:
	//
	// https://example.com
	OriginalUrl *string `json:"originalUrl,omitempty" xml:"originalUrl,omitempty"`
	// The list of resource tags as JSON strings.
	//
	// example:
	//
	// ["R&D"]
	SourceTags *string `json:"sourceTags,omitempty" xml:"sourceTags,omitempty"`
	// The tenant ID. This is a common parameter. If not specified, the default tenant of the caller is used.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s CreateKnowledgeBasePublicUrlRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateKnowledgeBasePublicUrlRequest) GoString() string {
	return s.String()
}

func (s *CreateKnowledgeBasePublicUrlRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateKnowledgeBasePublicUrlRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *CreateKnowledgeBasePublicUrlRequest) GetName() *string {
	return s.Name
}

func (s *CreateKnowledgeBasePublicUrlRequest) GetNotes() *string {
	return s.Notes
}

func (s *CreateKnowledgeBasePublicUrlRequest) GetOperatingObjectName() *string {
	return s.OperatingObjectName
}

func (s *CreateKnowledgeBasePublicUrlRequest) GetOriginalUrl() *string {
	return s.OriginalUrl
}

func (s *CreateKnowledgeBasePublicUrlRequest) GetSourceTags() *string {
	return s.SourceTags
}

func (s *CreateKnowledgeBasePublicUrlRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *CreateKnowledgeBasePublicUrlRequest) SetDescription(v string) *CreateKnowledgeBasePublicUrlRequest {
	s.Description = &v
	return s
}

func (s *CreateKnowledgeBasePublicUrlRequest) SetDirectoryId(v string) *CreateKnowledgeBasePublicUrlRequest {
	s.DirectoryId = &v
	return s
}

func (s *CreateKnowledgeBasePublicUrlRequest) SetName(v string) *CreateKnowledgeBasePublicUrlRequest {
	s.Name = &v
	return s
}

func (s *CreateKnowledgeBasePublicUrlRequest) SetNotes(v string) *CreateKnowledgeBasePublicUrlRequest {
	s.Notes = &v
	return s
}

func (s *CreateKnowledgeBasePublicUrlRequest) SetOperatingObjectName(v string) *CreateKnowledgeBasePublicUrlRequest {
	s.OperatingObjectName = &v
	return s
}

func (s *CreateKnowledgeBasePublicUrlRequest) SetOriginalUrl(v string) *CreateKnowledgeBasePublicUrlRequest {
	s.OriginalUrl = &v
	return s
}

func (s *CreateKnowledgeBasePublicUrlRequest) SetSourceTags(v string) *CreateKnowledgeBasePublicUrlRequest {
	s.SourceTags = &v
	return s
}

func (s *CreateKnowledgeBasePublicUrlRequest) SetTenantId(v string) *CreateKnowledgeBasePublicUrlRequest {
	s.TenantId = &v
	return s
}

func (s *CreateKnowledgeBasePublicUrlRequest) Validate() error {
	return dara.Validate(s)
}
