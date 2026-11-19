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
	SetSource(v string) *ListMemoriesRequest
	GetSource() *string
	SetTeamId(v string) *ListMemoriesRequest
	GetTeamId() *string
	SetTypes(v string) *ListMemoriesRequest
	GetTypes() *string
	SetUpdatedAtFrom(v string) *ListMemoriesRequest
	GetUpdatedAtFrom() *string
	SetUpdatedAtTo(v string) *ListMemoriesRequest
	GetUpdatedAtTo() *string
	SetUserId(v string) *ListMemoriesRequest
	GetUserId() *string
	SetValid(v bool) *ListMemoriesRequest
	GetValid() *bool
}

type ListMemoriesRequest struct {
	// The agent ID. If not specified, all agents are matched. The length is limited to 1 to 128 UTF-8 bytes. The value cannot contain a forward slash (/) or an asterisk (*).
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
	// The session ID. If not specified, all sessions are matched. The length is limited to 1 to 128 UTF-8 bytes. The value cannot contain a forward slash (/) or an asterisk (*).
	//
	// example:
	//
	// session-001
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
	// The source of the memory. Valid values:
	//
	// - session_extraction: Extracted from session events.
	//
	// - content_extraction: Extracted from text or messages.
	//
	// - direct: Written directly.
	//
	// example:
	//
	// session_extraction
	Source *string `json:"source,omitempty" xml:"source,omitempty"`
	// The business team ID defined by the caller.
	//
	// example:
	//
	// team-001
	TeamId *string `json:"teamId,omitempty" xml:"teamId,omitempty"`
	// The memory content types. Specify the value as a JSON array string, such as ["skill"]. Valid values: semantic, skill, user_profile, and session_summary. If this parameter is omitted or set to [], all types are queried. The value auto is not supported. Keep the filter conditions consistent when you retrieve subsequent pages.
	//
	// example:
	//
	// ["semantic"]
	Types *string `json:"types,omitempty" xml:"types,omitempty"`
	// The start time for filtering by update time. The time must be in RFC 3339 format with a time zone.
	//
	// example:
	//
	// 2026-09-23T10:00:00Z
	UpdatedAtFrom *string `json:"updatedAtFrom,omitempty" xml:"updatedAtFrom,omitempty"`
	// The end time for filtering by update time. The time must be in RFC 3339 format with a time zone.
	//
	// example:
	//
	// 2026-09-24T10:00:00Z
	UpdatedAtTo *string `json:"updatedAtTo,omitempty" xml:"updatedAtTo,omitempty"`
	// The user ID. If not specified, all users are matched. The length is limited to 1 to 128 UTF-8 bytes. The value cannot contain a forward slash (/) or an asterisk (*).
	//
	// example:
	//
	// user-001
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// Specifies whether the memory is valid. Default value: true. If this parameter is omitted, only valid records are returned. If set to false, revoked records are returned.
	Valid *bool `json:"valid,omitempty" xml:"valid,omitempty"`
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

func (s *ListMemoriesRequest) GetSource() *string {
	return s.Source
}

func (s *ListMemoriesRequest) GetTeamId() *string {
	return s.TeamId
}

func (s *ListMemoriesRequest) GetTypes() *string {
	return s.Types
}

func (s *ListMemoriesRequest) GetUpdatedAtFrom() *string {
	return s.UpdatedAtFrom
}

func (s *ListMemoriesRequest) GetUpdatedAtTo() *string {
	return s.UpdatedAtTo
}

func (s *ListMemoriesRequest) GetUserId() *string {
	return s.UserId
}

func (s *ListMemoriesRequest) GetValid() *bool {
	return s.Valid
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

func (s *ListMemoriesRequest) SetSource(v string) *ListMemoriesRequest {
	s.Source = &v
	return s
}

func (s *ListMemoriesRequest) SetTeamId(v string) *ListMemoriesRequest {
	s.TeamId = &v
	return s
}

func (s *ListMemoriesRequest) SetTypes(v string) *ListMemoriesRequest {
	s.Types = &v
	return s
}

func (s *ListMemoriesRequest) SetUpdatedAtFrom(v string) *ListMemoriesRequest {
	s.UpdatedAtFrom = &v
	return s
}

func (s *ListMemoriesRequest) SetUpdatedAtTo(v string) *ListMemoriesRequest {
	s.UpdatedAtTo = &v
	return s
}

func (s *ListMemoriesRequest) SetUserId(v string) *ListMemoriesRequest {
	s.UserId = &v
	return s
}

func (s *ListMemoriesRequest) SetValid(v bool) *ListMemoriesRequest {
	s.Valid = &v
	return s
}

func (s *ListMemoriesRequest) Validate() error {
	return dara.Validate(s)
}
