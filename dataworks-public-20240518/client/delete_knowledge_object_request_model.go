// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteKnowledgeObjectRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKnowledgeBaseId(v string) *DeleteKnowledgeObjectRequest
	GetKnowledgeBaseId() *string
	SetKnowledgeObjectId(v string) *DeleteKnowledgeObjectRequest
	GetKnowledgeObjectId() *string
}

type DeleteKnowledgeObjectRequest struct {
	// The ID of the knowledge base to which the object belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12345678-1234-4234-8234-123456789abc
	KnowledgeBaseId *string `json:"KnowledgeBaseId,omitempty" xml:"KnowledgeBaseId,omitempty"`
	// The stable ID of the object to delete.
	//
	// This parameter is required.
	//
	// example:
	//
	// obj-document-001
	KnowledgeObjectId *string `json:"KnowledgeObjectId,omitempty" xml:"KnowledgeObjectId,omitempty"`
}

func (s DeleteKnowledgeObjectRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteKnowledgeObjectRequest) GoString() string {
	return s.String()
}

func (s *DeleteKnowledgeObjectRequest) GetKnowledgeBaseId() *string {
	return s.KnowledgeBaseId
}

func (s *DeleteKnowledgeObjectRequest) GetKnowledgeObjectId() *string {
	return s.KnowledgeObjectId
}

func (s *DeleteKnowledgeObjectRequest) SetKnowledgeBaseId(v string) *DeleteKnowledgeObjectRequest {
	s.KnowledgeBaseId = &v
	return s
}

func (s *DeleteKnowledgeObjectRequest) SetKnowledgeObjectId(v string) *DeleteKnowledgeObjectRequest {
	s.KnowledgeObjectId = &v
	return s
}

func (s *DeleteKnowledgeObjectRequest) Validate() error {
	return dara.Validate(s)
}
