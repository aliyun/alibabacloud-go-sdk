// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMemorySessionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentId(v string) *ListMemorySessionsRequest
	GetAgentId() *string
	SetMaxResults(v int32) *ListMemorySessionsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListMemorySessionsRequest
	GetNextToken() *string
	SetUserId(v string) *ListMemorySessionsRequest
	GetUserId() *string
}

type ListMemorySessionsRequest struct {
	// The agent ID. If this parameter is not specified, all agents are matched. The length is limited to 1 to 128 UTF-8 bytes, and the value cannot contain / or *.
	//
	// example:
	//
	// agent-001
	AgentId *string `json:"agentId,omitempty" xml:"agentId,omitempty"`
	// The maximum number of records per page. Default value: 100. Valid values: 1 to 100.
	//
	// example:
	//
	// 100
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The pagination token for the next page.
	//
	// example:
	//
	// eyJvZmZzZXQiOjEwMH0=
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The user ID. If this parameter is not specified, all users are matched. The length is limited to 1 to 128 UTF-8 bytes, and the value cannot contain / or *.
	//
	// example:
	//
	// user-001
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s ListMemorySessionsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListMemorySessionsRequest) GoString() string {
	return s.String()
}

func (s *ListMemorySessionsRequest) GetAgentId() *string {
	return s.AgentId
}

func (s *ListMemorySessionsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListMemorySessionsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListMemorySessionsRequest) GetUserId() *string {
	return s.UserId
}

func (s *ListMemorySessionsRequest) SetAgentId(v string) *ListMemorySessionsRequest {
	s.AgentId = &v
	return s
}

func (s *ListMemorySessionsRequest) SetMaxResults(v int32) *ListMemorySessionsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListMemorySessionsRequest) SetNextToken(v string) *ListMemorySessionsRequest {
	s.NextToken = &v
	return s
}

func (s *ListMemorySessionsRequest) SetUserId(v string) *ListMemorySessionsRequest {
	s.UserId = &v
	return s
}

func (s *ListMemorySessionsRequest) Validate() error {
	return dara.Validate(s)
}
