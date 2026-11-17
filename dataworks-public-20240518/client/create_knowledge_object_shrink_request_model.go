// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateKnowledgeObjectShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateKnowledgeObjectShrinkRequest
	GetClientToken() *string
	SetContentShrink(v string) *CreateKnowledgeObjectShrinkRequest
	GetContentShrink() *string
	SetDescription(v string) *CreateKnowledgeObjectShrinkRequest
	GetDescription() *string
	SetKnowledgeBaseId(v string) *CreateKnowledgeObjectShrinkRequest
	GetKnowledgeBaseId() *string
	SetName(v string) *CreateKnowledgeObjectShrinkRequest
	GetName() *string
	SetType(v string) *CreateKnowledgeObjectShrinkRequest
	GetType() *string
	SetUploadId(v string) *CreateKnowledgeObjectShrinkRequest
	GetUploadId() *string
}

type CreateKnowledgeObjectShrinkRequest struct {
	// The idempotency token for the creation request. Keep the original request unchanged when retrying.
	//
	// example:
	//
	// create-document-001
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Only Content.Document is supported. It includes Text and ContentFormat (MARKDOWN or TEXT).
	ContentShrink *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The document description.
	//
	// example:
	//
	// Sales metrics and policy document
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the target knowledge base.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12345678-1234-4234-8234-123456789abc
	KnowledgeBaseId *string `json:"KnowledgeBaseId,omitempty" xml:"KnowledgeBaseId,omitempty"`
	// The document name. Existing objects are not overwritten based on the name.
	//
	// This parameter is required.
	//
	// example:
	//
	// Business description
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The type of the object. For this operation, the value is fixed to DOCUMENT.
	//
	// This parameter is required.
	//
	// example:
	//
	// DOCUMENT
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The ID of the completed PUT upload session for this knowledge base. This parameter is mutually exclusive with Content.
	//
	// example:
	//
	// 12345678-1234-4234-8234-123456789abd
	UploadId *string `json:"UploadId,omitempty" xml:"UploadId,omitempty"`
}

func (s CreateKnowledgeObjectShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateKnowledgeObjectShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateKnowledgeObjectShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateKnowledgeObjectShrinkRequest) GetContentShrink() *string {
	return s.ContentShrink
}

func (s *CreateKnowledgeObjectShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateKnowledgeObjectShrinkRequest) GetKnowledgeBaseId() *string {
	return s.KnowledgeBaseId
}

func (s *CreateKnowledgeObjectShrinkRequest) GetName() *string {
	return s.Name
}

func (s *CreateKnowledgeObjectShrinkRequest) GetType() *string {
	return s.Type
}

func (s *CreateKnowledgeObjectShrinkRequest) GetUploadId() *string {
	return s.UploadId
}

func (s *CreateKnowledgeObjectShrinkRequest) SetClientToken(v string) *CreateKnowledgeObjectShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateKnowledgeObjectShrinkRequest) SetContentShrink(v string) *CreateKnowledgeObjectShrinkRequest {
	s.ContentShrink = &v
	return s
}

func (s *CreateKnowledgeObjectShrinkRequest) SetDescription(v string) *CreateKnowledgeObjectShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreateKnowledgeObjectShrinkRequest) SetKnowledgeBaseId(v string) *CreateKnowledgeObjectShrinkRequest {
	s.KnowledgeBaseId = &v
	return s
}

func (s *CreateKnowledgeObjectShrinkRequest) SetName(v string) *CreateKnowledgeObjectShrinkRequest {
	s.Name = &v
	return s
}

func (s *CreateKnowledgeObjectShrinkRequest) SetType(v string) *CreateKnowledgeObjectShrinkRequest {
	s.Type = &v
	return s
}

func (s *CreateKnowledgeObjectShrinkRequest) SetUploadId(v string) *CreateKnowledgeObjectShrinkRequest {
	s.UploadId = &v
	return s
}

func (s *CreateKnowledgeObjectShrinkRequest) Validate() error {
	return dara.Validate(s)
}
