// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSessionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLimit(v int32) *ListSessionsRequest
	GetLimit() *int32
	SetNextToken(v string) *ListSessionsRequest
	GetNextToken() *string
	SetQualifier(v string) *ListSessionsRequest
	GetQualifier() *string
	SetSessionId(v string) *ListSessionsRequest
	GetSessionId() *string
	SetSessionStatus(v string) *ListSessionsRequest
	GetSessionStatus() *string
}

type ListSessionsRequest struct {
	// example:
	//
	// 10
	Limit *int32 `json:"limit,omitempty" xml:"limit,omitempty"`
	// example:
	//
	// MTIzNCNhYmM=
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// example:
	//
	// aliasName1
	Qualifier *string `json:"qualifier,omitempty" xml:"qualifier,omitempty"`
	// example:
	//
	// test-session-id-1
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
	// example:
	//
	// Active
	SessionStatus *string `json:"sessionStatus,omitempty" xml:"sessionStatus,omitempty"`
}

func (s ListSessionsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSessionsRequest) GoString() string {
	return s.String()
}

func (s *ListSessionsRequest) GetLimit() *int32 {
	return s.Limit
}

func (s *ListSessionsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListSessionsRequest) GetQualifier() *string {
	return s.Qualifier
}

func (s *ListSessionsRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *ListSessionsRequest) GetSessionStatus() *string {
	return s.SessionStatus
}

func (s *ListSessionsRequest) SetLimit(v int32) *ListSessionsRequest {
	s.Limit = &v
	return s
}

func (s *ListSessionsRequest) SetNextToken(v string) *ListSessionsRequest {
	s.NextToken = &v
	return s
}

func (s *ListSessionsRequest) SetQualifier(v string) *ListSessionsRequest {
	s.Qualifier = &v
	return s
}

func (s *ListSessionsRequest) SetSessionId(v string) *ListSessionsRequest {
	s.SessionId = &v
	return s
}

func (s *ListSessionsRequest) SetSessionStatus(v string) *ListSessionsRequest {
	s.SessionStatus = &v
	return s
}

func (s *ListSessionsRequest) Validate() error {
	return dara.Validate(s)
}
