// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMemoriesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentId(v string) *ListMemoriesRequest
	GetAgentId() *string
	SetMaxResults(v int32) *ListMemoriesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListMemoriesRequest
	GetNextToken() *string
	SetSessionId(v string) *ListMemoriesRequest
	GetSessionId() *string
	SetUserId(v string) *ListMemoriesRequest
	GetUserId() *string
}

type ListMemoriesRequest struct {
	// The agent ID. If you do not specify this parameter, all agents are matched. The value must be 1 to 128 bytes in length and encoded in UTF-8. The value cannot contain a forward slash (/) or an asterisk (*).
	//
	// example:
	//
	// agent-001
	AgentId *string `json:"agentId,omitempty" xml:"agentId,omitempty"`
	// The maximum number of records per page. Default value: 50. Valid values: 1 to 100.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The pagination token for the next page.
	//
	// example:
	//
	// eyJvZmZzZXQiOjEwfQ==
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The session ID. If you do not specify this parameter, all sessions are matched. The value must be 1 to 128 bytes in length and encoded in UTF-8. The value cannot contain a forward slash (/) or an asterisk (*).
	//
	// example:
	//
	// session-001
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
	// The user ID. If you do not specify this parameter, all users are matched. The value must be 1 to 128 bytes in length and encoded in UTF-8. The value cannot contain a forward slash (/) or an asterisk (*).
	//
	// example:
	//
	// user-001
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s ListMemoriesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListMemoriesRequest) GoString() string {
	return s.String()
}

func (s *ListMemoriesRequest) GetAgentId() *string {
	return s.AgentId
}

func (s *ListMemoriesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListMemoriesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListMemoriesRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *ListMemoriesRequest) GetUserId() *string {
	return s.UserId
}

func (s *ListMemoriesRequest) SetAgentId(v string) *ListMemoriesRequest {
	s.AgentId = &v
	return s
}

func (s *ListMemoriesRequest) SetMaxResults(v int32) *ListMemoriesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListMemoriesRequest) SetNextToken(v string) *ListMemoriesRequest {
	s.NextToken = &v
	return s
}

func (s *ListMemoriesRequest) SetSessionId(v string) *ListMemoriesRequest {
	s.SessionId = &v
	return s
}

func (s *ListMemoriesRequest) SetUserId(v string) *ListMemoriesRequest {
	s.UserId = &v
	return s
}

func (s *ListMemoriesRequest) Validate() error {
	return dara.Validate(s)
}
