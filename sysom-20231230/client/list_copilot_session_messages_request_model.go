// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCopilotSessionMessagesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentId(v string) *ListCopilotSessionMessagesRequest
	GetAgentId() *string
	SetLimit(v int64) *ListCopilotSessionMessagesRequest
	GetLimit() *int64
	SetMaxResults(v int32) *ListCopilotSessionMessagesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListCopilotSessionMessagesRequest
	GetNextToken() *string
	SetOffset(v int64) *ListCopilotSessionMessagesRequest
	GetOffset() *int64
	SetSessionId(v string) *ListCopilotSessionMessagesRequest
	GetSessionId() *string
}

type ListCopilotSessionMessagesRequest struct {
	// The agent ID retained for compatibility with historical calls. This parameter can be omitted. Currently, the entire session is queried by sessionId, and this parameter does not change the query scope of historical events.
	//
	// example:
	//
	// sysagent
	AgentId *string `json:"agentId,omitempty" xml:"agentId,omitempty"`
	// The page size retained for compatibility with historical calls. Default value: 50. Valid values: 1 to 200. For new integrations, we recommend using maxResults. If both parameters are specified, maxResults takes precedence.
	//
	// example:
	//
	// 50
	Limit *int64 `json:"limit,omitempty" xml:"limit,omitempty"`
	// The page size. Default value: 50. Valid values: 1 to 1000.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The token for the next page. Do not specify this parameter for the first request. For subsequent requests, use the nextToken value from the previous response. Clear this value when you change the filter conditions.
	//
	// example:
	//
	// 50
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The pagination offset retained for compatibility with historical calls. Default value: 0. For new integrations, we recommend using nextToken. Do not specify this parameter and nextToken at the same time.
	//
	// example:
	//
	// 0
	Offset *int64 `json:"offset,omitempty" xml:"offset,omitempty"`
	// The session ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// s001
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
}

func (s ListCopilotSessionMessagesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCopilotSessionMessagesRequest) GoString() string {
	return s.String()
}

func (s *ListCopilotSessionMessagesRequest) GetAgentId() *string {
	return s.AgentId
}

func (s *ListCopilotSessionMessagesRequest) GetLimit() *int64 {
	return s.Limit
}

func (s *ListCopilotSessionMessagesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListCopilotSessionMessagesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListCopilotSessionMessagesRequest) GetOffset() *int64 {
	return s.Offset
}

func (s *ListCopilotSessionMessagesRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *ListCopilotSessionMessagesRequest) SetAgentId(v string) *ListCopilotSessionMessagesRequest {
	s.AgentId = &v
	return s
}

func (s *ListCopilotSessionMessagesRequest) SetLimit(v int64) *ListCopilotSessionMessagesRequest {
	s.Limit = &v
	return s
}

func (s *ListCopilotSessionMessagesRequest) SetMaxResults(v int32) *ListCopilotSessionMessagesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListCopilotSessionMessagesRequest) SetNextToken(v string) *ListCopilotSessionMessagesRequest {
	s.NextToken = &v
	return s
}

func (s *ListCopilotSessionMessagesRequest) SetOffset(v int64) *ListCopilotSessionMessagesRequest {
	s.Offset = &v
	return s
}

func (s *ListCopilotSessionMessagesRequest) SetSessionId(v string) *ListCopilotSessionMessagesRequest {
	s.SessionId = &v
	return s
}

func (s *ListCopilotSessionMessagesRequest) Validate() error {
	return dara.Validate(s)
}
