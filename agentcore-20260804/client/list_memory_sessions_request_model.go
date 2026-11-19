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
	SetTeamId(v string) *ListMemorySessionsRequest
	GetTeamId() *string
	SetUserId(v string) *ListMemorySessionsRequest
	GetUserId() *string
}

type ListMemorySessionsRequest struct {
	// The agent ID. If you do not specify this parameter, all agents are matched. The value must be 1 to 128 UTF-8 bytes in length and cannot contain a forward slash (/) or an asterisk (*).
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
	// The business team ID defined by the caller.
	//
	// example:
	//
	// team-001
	TeamId *string `json:"teamId,omitempty" xml:"teamId,omitempty"`
	// The user ID. If you do not specify this parameter, all users are matched. The value must be 1 to 128 UTF-8 bytes in length and cannot contain a forward slash (/) or an asterisk (*).
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

func (s *ListMemorySessionsRequest) GetTeamId() *string {
	return s.TeamId
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

func (s *ListMemorySessionsRequest) SetTeamId(v string) *ListMemorySessionsRequest {
	s.TeamId = &v
	return s
}

func (s *ListMemorySessionsRequest) SetUserId(v string) *ListMemorySessionsRequest {
	s.UserId = &v
	return s
}

func (s *ListMemorySessionsRequest) Validate() error {
	return dara.Validate(s)
}
