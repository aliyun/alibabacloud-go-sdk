// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteKnowledgeSourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKnowledgeBaseId(v string) *DeleteKnowledgeSourceRequest
	GetKnowledgeBaseId() *string
	SetSourceId(v string) *DeleteKnowledgeSourceRequest
	GetSourceId() *string
}

type DeleteKnowledgeSourceRequest struct {
	// The ID of the knowledge base to which the source belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12345678-1234-4234-8234-123456789abc
	KnowledgeBaseId *string `json:"KnowledgeBaseId,omitempty" xml:"KnowledgeBaseId,omitempty"`
	// The ID of the source.
	//
	// This parameter is required.
	//
	// example:
	//
	// src-001
	SourceId *string `json:"SourceId,omitempty" xml:"SourceId,omitempty"`
}

func (s DeleteKnowledgeSourceRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteKnowledgeSourceRequest) GoString() string {
	return s.String()
}

func (s *DeleteKnowledgeSourceRequest) GetKnowledgeBaseId() *string {
	return s.KnowledgeBaseId
}

func (s *DeleteKnowledgeSourceRequest) GetSourceId() *string {
	return s.SourceId
}

func (s *DeleteKnowledgeSourceRequest) SetKnowledgeBaseId(v string) *DeleteKnowledgeSourceRequest {
	s.KnowledgeBaseId = &v
	return s
}

func (s *DeleteKnowledgeSourceRequest) SetSourceId(v string) *DeleteKnowledgeSourceRequest {
	s.SourceId = &v
	return s
}

func (s *DeleteKnowledgeSourceRequest) Validate() error {
	return dara.Validate(s)
}
