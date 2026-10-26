// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAgentConversationsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListAgentConversationsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListAgentConversationsRequest
	GetNextToken() *string
	SetPurpose(v string) *ListAgentConversationsRequest
	GetPurpose() *string
	SetQ(v string) *ListAgentConversationsRequest
	GetQ() *string
}

type ListAgentConversationsRequest struct {
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// eyJpZCI6IjAxOThjOGE5In0=
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// example:
	//
	// chat
	Purpose *string `json:"purpose,omitempty" xml:"purpose,omitempty"`
	// example:
	//
	// 消息堆积
	Q *string `json:"q,omitempty" xml:"q,omitempty"`
}

func (s ListAgentConversationsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAgentConversationsRequest) GoString() string {
	return s.String()
}

func (s *ListAgentConversationsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListAgentConversationsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAgentConversationsRequest) GetPurpose() *string {
	return s.Purpose
}

func (s *ListAgentConversationsRequest) GetQ() *string {
	return s.Q
}

func (s *ListAgentConversationsRequest) SetMaxResults(v int32) *ListAgentConversationsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListAgentConversationsRequest) SetNextToken(v string) *ListAgentConversationsRequest {
	s.NextToken = &v
	return s
}

func (s *ListAgentConversationsRequest) SetPurpose(v string) *ListAgentConversationsRequest {
	s.Purpose = &v
	return s
}

func (s *ListAgentConversationsRequest) SetQ(v string) *ListAgentConversationsRequest {
	s.Q = &v
	return s
}

func (s *ListAgentConversationsRequest) Validate() error {
	return dara.Validate(s)
}
