// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetKnowledgeSourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKnowledgeBaseId(v string) *GetKnowledgeSourceRequest
	GetKnowledgeBaseId() *string
	SetSourceId(v string) *GetKnowledgeSourceRequest
	GetSourceId() *string
}

type GetKnowledgeSourceRequest struct {
	// The ID of the knowledge base to which the source belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12345678-1234-4234-8234-123456789abc
	KnowledgeBaseId *string `json:"KnowledgeBaseId,omitempty" xml:"KnowledgeBaseId,omitempty"`
	// The source ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// src-001
	SourceId *string `json:"SourceId,omitempty" xml:"SourceId,omitempty"`
}

func (s GetKnowledgeSourceRequest) String() string {
	return dara.Prettify(s)
}

func (s GetKnowledgeSourceRequest) GoString() string {
	return s.String()
}

func (s *GetKnowledgeSourceRequest) GetKnowledgeBaseId() *string {
	return s.KnowledgeBaseId
}

func (s *GetKnowledgeSourceRequest) GetSourceId() *string {
	return s.SourceId
}

func (s *GetKnowledgeSourceRequest) SetKnowledgeBaseId(v string) *GetKnowledgeSourceRequest {
	s.KnowledgeBaseId = &v
	return s
}

func (s *GetKnowledgeSourceRequest) SetSourceId(v string) *GetKnowledgeSourceRequest {
	s.SourceId = &v
	return s
}

func (s *GetKnowledgeSourceRequest) Validate() error {
	return dara.Validate(s)
}
