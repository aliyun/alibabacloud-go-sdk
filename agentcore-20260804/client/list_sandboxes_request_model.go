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
	// The pagination token for the next page. When paginating, keep workspaceId, agentId, searchText, sessionId, and maxResults unchanged.
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// When explicitly specified, this parameter cannot be empty or consist entirely of whitespace. The value can contain up to 36 characters and allows only hexadecimal characters and hyphens. This parameter can be used together with sessionId, combined with AND logic.
	//
	// example:
	//
	// 2f360
	SearchText *string `json:"searchText,omitempty" xml:"searchText,omitempty"`
	// When explicitly specified, this parameter cannot be empty or consist entirely of whitespace. The value must be valid UTF-8 of 1 to 128 bytes and cannot contain forward slashes (/), vertical bars (|), control characters, or format characters. The original case and leading/trailing spaces are preserved. This parameter can be used together with searchText, combined with AND logic.
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
