// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListKnowledgeSourcesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKnowledgeBaseId(v string) *ListKnowledgeSourcesRequest
	GetKnowledgeBaseId() *string
	SetMaxResults(v int32) *ListKnowledgeSourcesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListKnowledgeSourcesRequest
	GetNextToken() *string
	SetType(v string) *ListKnowledgeSourcesRequest
	GetType() *string
}

type ListKnowledgeSourcesRequest struct {
	// The ID of the target knowledge base.
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
	// The pagination token.
	//
	// example:
	//
	// eyJvZmZzZXQiOjIwfQ
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Filters the results by source type.
	//
	// example:
	//
	// PROJECT
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListKnowledgeSourcesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListKnowledgeSourcesRequest) GoString() string {
	return s.String()
}

func (s *ListKnowledgeSourcesRequest) GetKnowledgeBaseId() *string {
	return s.KnowledgeBaseId
}

func (s *ListKnowledgeSourcesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListKnowledgeSourcesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListKnowledgeSourcesRequest) GetType() *string {
	return s.Type
}

func (s *ListKnowledgeSourcesRequest) SetKnowledgeBaseId(v string) *ListKnowledgeSourcesRequest {
	s.KnowledgeBaseId = &v
	return s
}

func (s *ListKnowledgeSourcesRequest) SetMaxResults(v int32) *ListKnowledgeSourcesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListKnowledgeSourcesRequest) SetNextToken(v string) *ListKnowledgeSourcesRequest {
	s.NextToken = &v
	return s
}

func (s *ListKnowledgeSourcesRequest) SetType(v string) *ListKnowledgeSourcesRequest {
	s.Type = &v
	return s
}

func (s *ListKnowledgeSourcesRequest) Validate() error {
	return dara.Validate(s)
}
