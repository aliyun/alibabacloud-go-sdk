// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChatStreamRequest interface {
	dara.Model
	String() string
	GoString() string
	SetQuestion(v string) *ChatStreamRequest
	GetQuestion() *string
	SetSessionId(v string) *ChatStreamRequest
	GetSessionId() *string
}

type ChatStreamRequest struct {
	// Q&A content.
	//
	// This parameter is required.
	//
	// example:
	//
	// How to access knowledge base Q&A documents.
	Question *string `json:"question,omitempty" xml:"question,omitempty"`
	// - Q&A session ID.
	//
	// - Historical sessions can be retrieved through GetSessionList.
	//
	// - A new session can also be created via CreateChatSession.
	//
	// This parameter is required.
	//
	// example:
	//
	// bfce2248-1546-4298-8bcf-70ac26e69646
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
}

func (s ChatStreamRequest) String() string {
	return dara.Prettify(s)
}

func (s ChatStreamRequest) GoString() string {
	return s.String()
}

func (s *ChatStreamRequest) GetQuestion() *string {
	return s.Question
}

func (s *ChatStreamRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *ChatStreamRequest) SetQuestion(v string) *ChatStreamRequest {
	s.Question = &v
	return s
}

func (s *ChatStreamRequest) SetSessionId(v string) *ChatStreamRequest {
	s.SessionId = &v
	return s
}

func (s *ChatStreamRequest) Validate() error {
	return dara.Validate(s)
}
