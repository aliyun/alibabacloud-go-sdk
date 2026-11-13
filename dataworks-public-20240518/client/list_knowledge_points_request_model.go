// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListKnowledgePointsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKnowledgeBaseId(v string) *ListKnowledgePointsRequest
	GetKnowledgeBaseId() *string
	SetMaxResults(v int32) *ListKnowledgePointsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListKnowledgePointsRequest
	GetNextToken() *string
	SetType(v string) *ListKnowledgePointsRequest
	GetType() *string
}

type ListKnowledgePointsRequest struct {
	// The stable ID of the target knowledge base. You can obtain the ID by calling the ListKnowledgeBases operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12345678-1234-4234-8234-123456789abc
	KnowledgeBaseId *string `json:"KnowledgeBaseId,omitempty" xml:"KnowledgeBaseId,omitempty"`
	// The maximum number of knowledge points to return per page. Valid values: 1 to 100. Default value: 20.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination cursor. Do not specify this parameter for the first page. For subsequent pages, pass the NextToken value returned in the previous response as is, and keep the knowledge base, type, and page size unchanged. An empty cursor indicates that no more data is available. If the current page is empty but the cursor is not empty, continue querying.
	//
	// example:
	//
	// eyJvZmZzZXQiOjIwfQ
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The type of the knowledge point. Valid values:
	//
	// - GRAPH: local graph entry.
	//
	// - WIKI: Wiki page.
	//
	// - KNOWLEDGE_CONTENT: knowledge content.
	//
	// If you do not specify this parameter, knowledge points of all types are listed.
	//
	// example:
	//
	// GRAPH
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListKnowledgePointsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListKnowledgePointsRequest) GoString() string {
	return s.String()
}

func (s *ListKnowledgePointsRequest) GetKnowledgeBaseId() *string {
	return s.KnowledgeBaseId
}

func (s *ListKnowledgePointsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListKnowledgePointsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListKnowledgePointsRequest) GetType() *string {
	return s.Type
}

func (s *ListKnowledgePointsRequest) SetKnowledgeBaseId(v string) *ListKnowledgePointsRequest {
	s.KnowledgeBaseId = &v
	return s
}

func (s *ListKnowledgePointsRequest) SetMaxResults(v int32) *ListKnowledgePointsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListKnowledgePointsRequest) SetNextToken(v string) *ListKnowledgePointsRequest {
	s.NextToken = &v
	return s
}

func (s *ListKnowledgePointsRequest) SetType(v string) *ListKnowledgePointsRequest {
	s.Type = &v
	return s
}

func (s *ListKnowledgePointsRequest) Validate() error {
	return dara.Validate(s)
}
