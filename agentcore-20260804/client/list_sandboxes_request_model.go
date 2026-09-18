// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSandboxesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListSandboxesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListSandboxesRequest
	GetNextToken() *string
	SetSearchText(v string) *ListSandboxesRequest
	GetSearchText() *string
	SetSessionId(v string) *ListSandboxesRequest
	GetSessionId() *string
}

type ListSandboxesRequest struct {
	// The maximum number of records per page. Default value: 20.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The pagination token for querying the next page. When paginating, keep workspaceId, agentId, searchText, sessionId, and maxResults unchanged.
	//
	// example:
	//
	// next-token-1
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// Performs a case-insensitive fuzzy search by sandbox ID fragment.
	//
	// example:
	//
	// 2f360
	SearchText *string `json:"searchText,omitempty" xml:"searchText,omitempty"`
	// Performs a case-insensitive fuzzy search by active session ID fragment.
	//
	// example:
	//
	// session-a
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
}

func (s ListSandboxesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSandboxesRequest) GoString() string {
	return s.String()
}

func (s *ListSandboxesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListSandboxesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListSandboxesRequest) GetSearchText() *string {
	return s.SearchText
}

func (s *ListSandboxesRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *ListSandboxesRequest) SetMaxResults(v int32) *ListSandboxesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListSandboxesRequest) SetNextToken(v string) *ListSandboxesRequest {
	s.NextToken = &v
	return s
}

func (s *ListSandboxesRequest) SetSearchText(v string) *ListSandboxesRequest {
	s.SearchText = &v
	return s
}

func (s *ListSandboxesRequest) SetSessionId(v string) *ListSandboxesRequest {
	s.SessionId = &v
	return s
}

func (s *ListSandboxesRequest) Validate() error {
	return dara.Validate(s)
}
