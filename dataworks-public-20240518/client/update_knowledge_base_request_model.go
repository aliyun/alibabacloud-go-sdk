// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateKnowledgeBaseRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *UpdateKnowledgeBaseRequest
	GetDescription() *string
	SetKnowledgeBaseId(v string) *UpdateKnowledgeBaseRequest
	GetKnowledgeBaseId() *string
	SetName(v string) *UpdateKnowledgeBaseRequest
	GetName() *string
}

type UpdateKnowledgeBaseRequest struct {
	// The new description.
	//
	// example:
	//
	// Sales metrics and policy documents
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the target knowledge base.
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
	// Sales Knowledge
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s UpdateKnowledgeBaseRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateKnowledgeBaseRequest) GoString() string {
	return s.String()
}

func (s *UpdateKnowledgeBaseRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateKnowledgeBaseRequest) GetKnowledgeBaseId() *string {
	return s.KnowledgeBaseId
}

func (s *UpdateKnowledgeBaseRequest) GetName() *string {
	return s.Name
}

func (s *UpdateKnowledgeBaseRequest) SetDescription(v string) *UpdateKnowledgeBaseRequest {
	s.Description = &v
	return s
}

func (s *UpdateKnowledgeBaseRequest) SetKnowledgeBaseId(v string) *UpdateKnowledgeBaseRequest {
	s.KnowledgeBaseId = &v
	return s
}

func (s *UpdateKnowledgeBaseRequest) SetName(v string) *UpdateKnowledgeBaseRequest {
	s.Name = &v
	return s
}

func (s *UpdateKnowledgeBaseRequest) Validate() error {
	return dara.Validate(s)
}
