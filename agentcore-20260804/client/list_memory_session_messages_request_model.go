// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMemorySessionMessagesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentId(v string) *ListMemorySessionMessagesRequest
	GetAgentId() *string
	SetMaxResults(v int32) *ListMemorySessionMessagesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListMemorySessionMessagesRequest
	GetNextToken() *string
	SetSessionId(v string) *ListMemorySessionMessagesRequest
	GetSessionId() *string
	SetUserId(v string) *ListMemorySessionMessagesRequest
	GetUserId() *string
}

type ListMemorySessionMessagesRequest struct {
	// The agent ID. At least one of UserId and AgentId must be provided. The length must be 1 to 128 UTF-8 bytes. The value cannot contain forward slashes (/) or asterisks (*).
	//
	// example:
	//
	// agent-001
	AgentId *string `json:"agentId,omitempty" xml:"agentId,omitempty"`
	// The maximum number of records per page. Default value: 200. Valid values: 1 to 500.
	//
	// example:
	//
	// 200
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The pagination token for the next page.
	//
	// example:
	//
	// eyJvZmZzZXQiOjIwMH0=
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The session ID. The length must be 1 to 128 UTF-8 bytes. The value cannot contain forward slashes (/) or asterisks (*).
	//
	// This parameter is required.
	//
	// example:
	//
	// session-001
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
	// The user ID. At least one of UserId and AgentId must be provided. The length must be 1 to 128 UTF-8 bytes. The value cannot contain forward slashes (/) or asterisks (*).
	//
	// example:
	//
	// user-001
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s ListMemorySessionMessagesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListMemorySessionMessagesRequest) GoString() string {
	return s.String()
}

func (s *ListMemorySessionMessagesRequest) GetAgentId() *string {
	return s.AgentId
}

func (s *ListMemorySessionMessagesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListMemorySessionMessagesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListMemorySessionMessagesRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *ListMemorySessionMessagesRequest) GetUserId() *string {
	return s.UserId
}

func (s *ListMemorySessionMessagesRequest) SetAgentId(v string) *ListMemorySessionMessagesRequest {
	s.AgentId = &v
	return s
}

func (s *ListMemorySessionMessagesRequest) SetMaxResults(v int32) *ListMemorySessionMessagesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListMemorySessionMessagesRequest) SetNextToken(v string) *ListMemorySessionMessagesRequest {
	s.NextToken = &v
	return s
}

func (s *ListMemorySessionMessagesRequest) SetSessionId(v string) *ListMemorySessionMessagesRequest {
	s.SessionId = &v
	return s
}

func (s *ListMemorySessionMessagesRequest) SetUserId(v string) *ListMemorySessionMessagesRequest {
	s.UserId = &v
	return s
}

func (s *ListMemorySessionMessagesRequest) Validate() error {
	return dara.Validate(s)
}
