// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListKnowledgeObjectsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKnowledgeBaseId(v string) *ListKnowledgeObjectsRequest
	GetKnowledgeBaseId() *string
	SetMaxResults(v int32) *ListKnowledgeObjectsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListKnowledgeObjectsRequest
	GetNextToken() *string
	SetType(v string) *ListKnowledgeObjectsRequest
	GetType() *string
}

type ListKnowledgeObjectsRequest struct {
	// The target knowledge base.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12345678-1234-4234-8234-123456789abc
	KnowledgeBaseId *string `json:"KnowledgeBaseId,omitempty" xml:"KnowledgeBaseId,omitempty"`
	// The maximum number of entries per page.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination cursor.
	//
	// example:
	//
	// eyJvZmZzZXQiOjIwfQ
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The type of the knowledge object. Valid values: SEMANTIC_MODEL and DOCUMENT.
	//
	// example:
	//
	// DOCUMENT
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListKnowledgeObjectsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListKnowledgeObjectsRequest) GoString() string {
	return s.String()
}

func (s *ListKnowledgeObjectsRequest) GetKnowledgeBaseId() *string {
	return s.KnowledgeBaseId
}

func (s *ListKnowledgeObjectsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListKnowledgeObjectsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListKnowledgeObjectsRequest) GetType() *string {
	return s.Type
}

func (s *ListKnowledgeObjectsRequest) SetKnowledgeBaseId(v string) *ListKnowledgeObjectsRequest {
	s.KnowledgeBaseId = &v
	return s
}

func (s *ListKnowledgeObjectsRequest) SetMaxResults(v int32) *ListKnowledgeObjectsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListKnowledgeObjectsRequest) SetNextToken(v string) *ListKnowledgeObjectsRequest {
	s.NextToken = &v
	return s
}

func (s *ListKnowledgeObjectsRequest) SetType(v string) *ListKnowledgeObjectsRequest {
	s.Type = &v
	return s
}

func (s *ListKnowledgeObjectsRequest) Validate() error {
	return dara.Validate(s)
}
