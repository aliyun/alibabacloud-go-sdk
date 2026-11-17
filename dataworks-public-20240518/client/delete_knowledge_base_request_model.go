// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteKnowledgeBaseRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKnowledgeBaseId(v string) *DeleteKnowledgeBaseRequest
	GetKnowledgeBaseId() *string
}

type DeleteKnowledgeBaseRequest struct {
	// The ID of the target knowledge base.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12345678-1234-4234-8234-123456789abc
	KnowledgeBaseId *string `json:"KnowledgeBaseId,omitempty" xml:"KnowledgeBaseId,omitempty"`
}

func (s DeleteKnowledgeBaseRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteKnowledgeBaseRequest) GoString() string {
	return s.String()
}

func (s *DeleteKnowledgeBaseRequest) GetKnowledgeBaseId() *string {
	return s.KnowledgeBaseId
}

func (s *DeleteKnowledgeBaseRequest) SetKnowledgeBaseId(v string) *DeleteKnowledgeBaseRequest {
	s.KnowledgeBaseId = &v
	return s
}

func (s *DeleteKnowledgeBaseRequest) Validate() error {
	return dara.Validate(s)
}
