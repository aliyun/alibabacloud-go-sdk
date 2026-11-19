// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEventsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentId(v string) *ListEventsRequest
	GetAgentId() *string
	SetMaxResults(v int32) *ListEventsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListEventsRequest
	GetNextToken() *string
	SetSessionId(v string) *ListEventsRequest
	GetSessionId() *string
	SetTeamId(v string) *ListEventsRequest
	GetTeamId() *string
	SetUserId(v string) *ListEventsRequest
	GetUserId() *string
}

type ListEventsRequest struct {
	// The exact filter for the business agent.
	//
	// example:
	//
	// agent-001
	AgentId *string `json:"agentId,omitempty" xml:"agentId,omitempty"`
	// The maximum number of entries per page. Default value: 200. Valid values: 1 to 500.
	//
	// example:
	//
	// 100
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The pagination token used to continue reading the same query.
	//
	// example:
	//
	// eyJvZmZzZXQiOjIwMH0=
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The ID of the session.
	//
	// This parameter is required.
	//
	// example:
	//
	// session-001
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
	// The exact filter for the business team.
	//
	// example:
	//
	// team-001
	TeamId *string `json:"teamId,omitempty" xml:"teamId,omitempty"`
	// The exact filter for the business user. This is not the platform tenant identity.
	//
	// example:
	//
	// user-001
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s ListEventsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListEventsRequest) GoString() string {
	return s.String()
}

func (s *ListEventsRequest) GetAgentId() *string {
	return s.AgentId
}

func (s *ListEventsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListEventsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListEventsRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *ListEventsRequest) GetTeamId() *string {
	return s.TeamId
}

func (s *ListEventsRequest) GetUserId() *string {
	return s.UserId
}

func (s *ListEventsRequest) SetAgentId(v string) *ListEventsRequest {
	s.AgentId = &v
	return s
}

func (s *ListEventsRequest) SetMaxResults(v int32) *ListEventsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListEventsRequest) SetNextToken(v string) *ListEventsRequest {
	s.NextToken = &v
	return s
}

func (s *ListEventsRequest) SetSessionId(v string) *ListEventsRequest {
	s.SessionId = &v
	return s
}

func (s *ListEventsRequest) SetTeamId(v string) *ListEventsRequest {
	s.TeamId = &v
	return s
}

func (s *ListEventsRequest) SetUserId(v string) *ListEventsRequest {
	s.UserId = &v
	return s
}

func (s *ListEventsRequest) Validate() error {
	return dara.Validate(s)
}
