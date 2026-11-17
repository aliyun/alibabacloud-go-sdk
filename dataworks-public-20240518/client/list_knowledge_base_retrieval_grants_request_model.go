// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListKnowledgeBaseRetrievalGrantsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKnowledgeBaseId(v string) *ListKnowledgeBaseRetrievalGrantsRequest
	GetKnowledgeBaseId() *string
	SetMaxResults(v int32) *ListKnowledgeBaseRetrievalGrantsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListKnowledgeBaseRetrievalGrantsRequest
	GetNextToken() *string
	SetPrincipalType(v string) *ListKnowledgeBaseRetrievalGrantsRequest
	GetPrincipalType() *string
}

type ListKnowledgeBaseRetrievalGrantsRequest struct {
	// The knowledge base ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12345678-1234-4234-8234-123456789abc
	KnowledgeBaseId *string `json:"KnowledgeBaseId,omitempty" xml:"KnowledgeBaseId,omitempty"`
	// The number of entries per page. This parameter is optional. Default value: 20. Valid values: 1 to 100.
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
	// The type of the principal that receives the share. This parameter is optional. Default value: USER. Only USER is supported in the initial release.
	//
	// example:
	//
	// USER
	PrincipalType *string `json:"PrincipalType,omitempty" xml:"PrincipalType,omitempty"`
}

func (s ListKnowledgeBaseRetrievalGrantsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListKnowledgeBaseRetrievalGrantsRequest) GoString() string {
	return s.String()
}

func (s *ListKnowledgeBaseRetrievalGrantsRequest) GetKnowledgeBaseId() *string {
	return s.KnowledgeBaseId
}

func (s *ListKnowledgeBaseRetrievalGrantsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListKnowledgeBaseRetrievalGrantsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListKnowledgeBaseRetrievalGrantsRequest) GetPrincipalType() *string {
	return s.PrincipalType
}

func (s *ListKnowledgeBaseRetrievalGrantsRequest) SetKnowledgeBaseId(v string) *ListKnowledgeBaseRetrievalGrantsRequest {
	s.KnowledgeBaseId = &v
	return s
}

func (s *ListKnowledgeBaseRetrievalGrantsRequest) SetMaxResults(v int32) *ListKnowledgeBaseRetrievalGrantsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListKnowledgeBaseRetrievalGrantsRequest) SetNextToken(v string) *ListKnowledgeBaseRetrievalGrantsRequest {
	s.NextToken = &v
	return s
}

func (s *ListKnowledgeBaseRetrievalGrantsRequest) SetPrincipalType(v string) *ListKnowledgeBaseRetrievalGrantsRequest {
	s.PrincipalType = &v
	return s
}

func (s *ListKnowledgeBaseRetrievalGrantsRequest) Validate() error {
	return dara.Validate(s)
}
