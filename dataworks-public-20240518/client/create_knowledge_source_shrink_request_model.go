// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateKnowledgeSourceShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateKnowledgeSourceShrinkRequest
	GetClientToken() *string
	SetDataSourceShrink(v string) *CreateKnowledgeSourceShrinkRequest
	GetDataSourceShrink() *string
	SetDescription(v string) *CreateKnowledgeSourceShrinkRequest
	GetDescription() *string
	SetDocumentShrink(v string) *CreateKnowledgeSourceShrinkRequest
	GetDocumentShrink() *string
	SetKnowledgeBaseId(v string) *CreateKnowledgeSourceShrinkRequest
	GetKnowledgeBaseId() *string
	SetName(v string) *CreateKnowledgeSourceShrinkRequest
	GetName() *string
	SetProjectShrink(v string) *CreateKnowledgeSourceShrinkRequest
	GetProjectShrink() *string
	SetType(v string) *CreateKnowledgeSourceShrinkRequest
	GetType() *string
}

type CreateKnowledgeSourceShrinkRequest struct {
	// The client token used to ensure idempotence.
	//
	// example:
	//
	// source-create-001
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The data source configuration. This parameter is required when Type is set to DATA_SOURCE.
	DataSourceShrink *string `json:"DataSource,omitempty" xml:"DataSource,omitempty"`
	// The description.
	//
	// example:
	//
	// Sales metrics and policy documents
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The document configuration. This parameter is required when Type is set to DOCUMENT.
	DocumentShrink *string `json:"Document,omitempty" xml:"Document,omitempty"`
	// The ID of the knowledge base.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12345678-1234-4234-8234-123456789abc
	KnowledgeBaseId *string `json:"KnowledgeBaseId,omitempty" xml:"KnowledgeBaseId,omitempty"`
	// The name of the source.
	//
	// This parameter is required.
	//
	// example:
	//
	// Sales topic project
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The project configuration. This parameter is required when Type is set to PROJECT.
	ProjectShrink *string `json:"Project,omitempty" xml:"Project,omitempty"`
	// The type of the knowledge source. Valid values: PROJECT, DATA_SOURCE, and DOCUMENT.
	//
	// This parameter is required.
	//
	// example:
	//
	// PROJECT
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateKnowledgeSourceShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateKnowledgeSourceShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateKnowledgeSourceShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateKnowledgeSourceShrinkRequest) GetDataSourceShrink() *string {
	return s.DataSourceShrink
}

func (s *CreateKnowledgeSourceShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateKnowledgeSourceShrinkRequest) GetDocumentShrink() *string {
	return s.DocumentShrink
}

func (s *CreateKnowledgeSourceShrinkRequest) GetKnowledgeBaseId() *string {
	return s.KnowledgeBaseId
}

func (s *CreateKnowledgeSourceShrinkRequest) GetName() *string {
	return s.Name
}

func (s *CreateKnowledgeSourceShrinkRequest) GetProjectShrink() *string {
	return s.ProjectShrink
}

func (s *CreateKnowledgeSourceShrinkRequest) GetType() *string {
	return s.Type
}

func (s *CreateKnowledgeSourceShrinkRequest) SetClientToken(v string) *CreateKnowledgeSourceShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateKnowledgeSourceShrinkRequest) SetDataSourceShrink(v string) *CreateKnowledgeSourceShrinkRequest {
	s.DataSourceShrink = &v
	return s
}

func (s *CreateKnowledgeSourceShrinkRequest) SetDescription(v string) *CreateKnowledgeSourceShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreateKnowledgeSourceShrinkRequest) SetDocumentShrink(v string) *CreateKnowledgeSourceShrinkRequest {
	s.DocumentShrink = &v
	return s
}

func (s *CreateKnowledgeSourceShrinkRequest) SetKnowledgeBaseId(v string) *CreateKnowledgeSourceShrinkRequest {
	s.KnowledgeBaseId = &v
	return s
}

func (s *CreateKnowledgeSourceShrinkRequest) SetName(v string) *CreateKnowledgeSourceShrinkRequest {
	s.Name = &v
	return s
}

func (s *CreateKnowledgeSourceShrinkRequest) SetProjectShrink(v string) *CreateKnowledgeSourceShrinkRequest {
	s.ProjectShrink = &v
	return s
}

func (s *CreateKnowledgeSourceShrinkRequest) SetType(v string) *CreateKnowledgeSourceShrinkRequest {
	s.Type = &v
	return s
}

func (s *CreateKnowledgeSourceShrinkRequest) Validate() error {
	return dara.Validate(s)
}
