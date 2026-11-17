// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateKnowledgeSourceShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDataSourceShrink(v string) *UpdateKnowledgeSourceShrinkRequest
	GetDataSourceShrink() *string
	SetDescription(v string) *UpdateKnowledgeSourceShrinkRequest
	GetDescription() *string
	SetDocumentShrink(v string) *UpdateKnowledgeSourceShrinkRequest
	GetDocumentShrink() *string
	SetKnowledgeBaseId(v string) *UpdateKnowledgeSourceShrinkRequest
	GetKnowledgeBaseId() *string
	SetName(v string) *UpdateKnowledgeSourceShrinkRequest
	GetName() *string
	SetProjectShrink(v string) *UpdateKnowledgeSourceShrinkRequest
	GetProjectShrink() *string
	SetSourceId(v string) *UpdateKnowledgeSourceShrinkRequest
	GetSourceId() *string
}

type UpdateKnowledgeSourceShrinkRequest struct {
	// The complete configuration for Type=DATA_SOURCE.
	DataSourceShrink *string `json:"DataSource,omitempty" xml:"DataSource,omitempty"`
	// The new description.
	//
	// example:
	//
	// Sales metrics and policy documents
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The complete configuration for Type=DOCUMENT. Credentials are not returned after they are written.
	DocumentShrink *string `json:"Document,omitempty" xml:"Document,omitempty"`
	// The ID of the knowledge base to which the source belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12345678-1234-4234-8234-123456789abc
	KnowledgeBaseId *string `json:"KnowledgeBaseId,omitempty" xml:"KnowledgeBaseId,omitempty"`
	// The new name.
	//
	// example:
	//
	// Sales Project
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The complete configuration for Type=PROJECT.
	ProjectShrink *string `json:"Project,omitempty" xml:"Project,omitempty"`
	// The source ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// src-001
	SourceId *string `json:"SourceId,omitempty" xml:"SourceId,omitempty"`
}

func (s UpdateKnowledgeSourceShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateKnowledgeSourceShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateKnowledgeSourceShrinkRequest) GetDataSourceShrink() *string {
	return s.DataSourceShrink
}

func (s *UpdateKnowledgeSourceShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateKnowledgeSourceShrinkRequest) GetDocumentShrink() *string {
	return s.DocumentShrink
}

func (s *UpdateKnowledgeSourceShrinkRequest) GetKnowledgeBaseId() *string {
	return s.KnowledgeBaseId
}

func (s *UpdateKnowledgeSourceShrinkRequest) GetName() *string {
	return s.Name
}

func (s *UpdateKnowledgeSourceShrinkRequest) GetProjectShrink() *string {
	return s.ProjectShrink
}

func (s *UpdateKnowledgeSourceShrinkRequest) GetSourceId() *string {
	return s.SourceId
}

func (s *UpdateKnowledgeSourceShrinkRequest) SetDataSourceShrink(v string) *UpdateKnowledgeSourceShrinkRequest {
	s.DataSourceShrink = &v
	return s
}

func (s *UpdateKnowledgeSourceShrinkRequest) SetDescription(v string) *UpdateKnowledgeSourceShrinkRequest {
	s.Description = &v
	return s
}

func (s *UpdateKnowledgeSourceShrinkRequest) SetDocumentShrink(v string) *UpdateKnowledgeSourceShrinkRequest {
	s.DocumentShrink = &v
	return s
}

func (s *UpdateKnowledgeSourceShrinkRequest) SetKnowledgeBaseId(v string) *UpdateKnowledgeSourceShrinkRequest {
	s.KnowledgeBaseId = &v
	return s
}

func (s *UpdateKnowledgeSourceShrinkRequest) SetName(v string) *UpdateKnowledgeSourceShrinkRequest {
	s.Name = &v
	return s
}

func (s *UpdateKnowledgeSourceShrinkRequest) SetProjectShrink(v string) *UpdateKnowledgeSourceShrinkRequest {
	s.ProjectShrink = &v
	return s
}

func (s *UpdateKnowledgeSourceShrinkRequest) SetSourceId(v string) *UpdateKnowledgeSourceShrinkRequest {
	s.SourceId = &v
	return s
}

func (s *UpdateKnowledgeSourceShrinkRequest) Validate() error {
	return dara.Validate(s)
}
