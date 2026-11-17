// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListKnowledgeBaseAdministratorsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKnowledgeBaseId(v string) *ListKnowledgeBaseAdministratorsRequest
	GetKnowledgeBaseId() *string
	SetMaxResults(v int32) *ListKnowledgeBaseAdministratorsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListKnowledgeBaseAdministratorsRequest
	GetNextToken() *string
}

type ListKnowledgeBaseAdministratorsRequest struct {
	// The ID of the knowledge base.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12345678-1234-4234-8234-123456789abc
	KnowledgeBaseId *string `json:"KnowledgeBaseId,omitempty" xml:"KnowledgeBaseId,omitempty"`
	// The number of entries per page. Default value: 20. Valid values: 1 to 100.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The cursor returned from the previous page. Do not specify this parameter for the first request.
	//
	// example:
	//
	// eyJvZmZzZXQiOjIwfQ
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListKnowledgeBaseAdministratorsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListKnowledgeBaseAdministratorsRequest) GoString() string {
	return s.String()
}

func (s *ListKnowledgeBaseAdministratorsRequest) GetKnowledgeBaseId() *string {
	return s.KnowledgeBaseId
}

func (s *ListKnowledgeBaseAdministratorsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListKnowledgeBaseAdministratorsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListKnowledgeBaseAdministratorsRequest) SetKnowledgeBaseId(v string) *ListKnowledgeBaseAdministratorsRequest {
	s.KnowledgeBaseId = &v
	return s
}

func (s *ListKnowledgeBaseAdministratorsRequest) SetMaxResults(v int32) *ListKnowledgeBaseAdministratorsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListKnowledgeBaseAdministratorsRequest) SetNextToken(v string) *ListKnowledgeBaseAdministratorsRequest {
	s.NextToken = &v
	return s
}

func (s *ListKnowledgeBaseAdministratorsRequest) Validate() error {
	return dara.Validate(s)
}
