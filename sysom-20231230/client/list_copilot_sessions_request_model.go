// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCopilotSessionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentId(v string) *ListCopilotSessionsRequest
	GetAgentId() *string
	SetMaxResults(v int32) *ListCopilotSessionsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListCopilotSessionsRequest
	GetNextToken() *string
	SetScheduledTaskId(v int64) *ListCopilotSessionsRequest
	GetScheduledTaskId() *int64
	SetSessionType(v string) *ListCopilotSessionsRequest
	GetSessionType() *string
}

type ListCopilotSessionsRequest struct {
	// The Agent ID. If this parameter is not specified, all Agent sessions of the current user are queried.
	//
	// example:
	//
	// sysagent
	AgentId *string `json:"agentId,omitempty" xml:"agentId,omitempty"`
	// The maximum number of entries per page. Default value: 50. Valid values: 1 to 1000.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The pagination token for the next page. Do not specify this parameter for the first query.
	//
	// example:
	//
	// 50
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The scheduled task ID, which must be a positive integer. If this parameter is specified, sessionType must be omitted or set to cron.
	//
	// example:
	//
	// 42
	ScheduledTaskId *int64 `json:"scheduledTaskId,omitempty" xml:"scheduledTaskId,omitempty"`
	// The session source. If this parameter is not specified, sessions are not filtered by source.
	//
	// example:
	//
	// openapi
	SessionType *string `json:"sessionType,omitempty" xml:"sessionType,omitempty"`
}

func (s ListCopilotSessionsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCopilotSessionsRequest) GoString() string {
	return s.String()
}

func (s *ListCopilotSessionsRequest) GetAgentId() *string {
	return s.AgentId
}

func (s *ListCopilotSessionsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListCopilotSessionsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListCopilotSessionsRequest) GetScheduledTaskId() *int64 {
	return s.ScheduledTaskId
}

func (s *ListCopilotSessionsRequest) GetSessionType() *string {
	return s.SessionType
}

func (s *ListCopilotSessionsRequest) SetAgentId(v string) *ListCopilotSessionsRequest {
	s.AgentId = &v
	return s
}

func (s *ListCopilotSessionsRequest) SetMaxResults(v int32) *ListCopilotSessionsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListCopilotSessionsRequest) SetNextToken(v string) *ListCopilotSessionsRequest {
	s.NextToken = &v
	return s
}

func (s *ListCopilotSessionsRequest) SetScheduledTaskId(v int64) *ListCopilotSessionsRequest {
	s.ScheduledTaskId = &v
	return s
}

func (s *ListCopilotSessionsRequest) SetSessionType(v string) *ListCopilotSessionsRequest {
	s.SessionType = &v
	return s
}

func (s *ListCopilotSessionsRequest) Validate() error {
	return dara.Validate(s)
}
