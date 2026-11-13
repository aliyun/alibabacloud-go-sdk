// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListKnowledgeBasesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListKnowledgeBasesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListKnowledgeBasesRequest
	GetNextToken() *string
}

type ListKnowledgeBasesRequest struct {
	// The maximum number of knowledge bases to return per page. Valid values: 1 to 100.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination cursor. Do not specify this parameter when you request the first page. For subsequent pages, use the NextToken value returned in the previous response as is. An empty response indicates that no more pages are available. If the current page is empty but the cursor is not empty, continue pagination. Keep the same MaxResults value during pagination.
	//
	// example:
	//
	// eyJvZmZzZXQiOjIwfQ
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListKnowledgeBasesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListKnowledgeBasesRequest) GoString() string {
	return s.String()
}

func (s *ListKnowledgeBasesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListKnowledgeBasesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListKnowledgeBasesRequest) SetMaxResults(v int32) *ListKnowledgeBasesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListKnowledgeBasesRequest) SetNextToken(v string) *ListKnowledgeBasesRequest {
	s.NextToken = &v
	return s
}

func (s *ListKnowledgeBasesRequest) Validate() error {
	return dara.Validate(s)
}
