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
	SetServiceEndpointId(v string) *ListSandboxesRequest
	GetServiceEndpointId() *string
	SetSessionId(v string) *ListSandboxesRequest
	GetSessionId() *string
}

type ListSandboxesRequest struct {
	// The maximum number of entries to return per page. Default value: 20. Valid values: 1 to 100.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The cursor for querying the next page. During pagination, you must keep workspaceId, agentId, searchText, sessionId, and maxResults unchanged.
	//
	// example:
	//
	// next-token-1
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// Performs a case-insensitive fuzzy search by Sandbox ID fragment.
	//
	// example:
	//
	// 2f360
	SearchText *string `json:"searchText,omitempty" xml:"searchText,omitempty"`
	// If omitted or left blank, queries the latest running unit of the agent. If specified, the service endpoint must belong to the agent and the target type must be AGENT_VERSION. Otherwise, Resource.NotFound is returned. An independently running service endpoint returns only the sandboxes of its own running unit. If the DEFAULT service endpoint has not taken over independent running, it falls back to the latest running unit. If a non-DEFAULT service endpoint is not running independently, ServiceEndpointTargetNotReady is returned. If the running unit has not been created, an empty list is returned and totalCount is 0, without reporting an error.
	//
	// example:
	//
	// se-7f3a92c1
	ServiceEndpointId *string `json:"serviceEndpointId,omitempty" xml:"serviceEndpointId,omitempty"`
	// Performs a case-insensitive fuzzy search by current valid session ID fragment.
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

func (s *ListSandboxesRequest) GetServiceEndpointId() *string {
	return s.ServiceEndpointId
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

func (s *ListSandboxesRequest) SetServiceEndpointId(v string) *ListSandboxesRequest {
	s.ServiceEndpointId = &v
	return s
}

func (s *ListSandboxesRequest) SetSessionId(v string) *ListSandboxesRequest {
	s.SessionId = &v
	return s
}

func (s *ListSandboxesRequest) Validate() error {
	return dara.Validate(s)
}
