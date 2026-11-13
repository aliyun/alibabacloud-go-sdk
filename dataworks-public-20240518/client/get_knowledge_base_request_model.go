// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetKnowledgeBaseRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKnowledgeBaseId(v string) *GetKnowledgeBaseRequest
	GetKnowledgeBaseId() *string
}

type GetKnowledgeBaseRequest struct {
	// The stable ID of the target knowledge base, which can be obtained by calling ListKnowledgeBases.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12345678-1234-4234-8234-123456789abc
	KnowledgeBaseId *string `json:"KnowledgeBaseId,omitempty" xml:"KnowledgeBaseId,omitempty"`
}

func (s GetKnowledgeBaseRequest) String() string {
	return dara.Prettify(s)
}

func (s GetKnowledgeBaseRequest) GoString() string {
	return s.String()
}

func (s *GetKnowledgeBaseRequest) GetKnowledgeBaseId() *string {
	return s.KnowledgeBaseId
}

func (s *GetKnowledgeBaseRequest) SetKnowledgeBaseId(v string) *GetKnowledgeBaseRequest {
	s.KnowledgeBaseId = &v
	return s
}

func (s *GetKnowledgeBaseRequest) Validate() error {
	return dara.Validate(s)
}
