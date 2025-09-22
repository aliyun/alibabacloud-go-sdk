// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChatRequest interface {
	dara.Model
	String() string
	GoString() string
	SetQuestion(v string) *ChatRequest
	GetQuestion() *string
	SetSessionId(v string) *ChatRequest
	GetSessionId() *string
}

type ChatRequest struct {
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

func (s ChatRequest) String() string {
	return dara.Prettify(s)
}

func (s ChatRequest) GoString() string {
	return s.String()
}

func (s *ChatRequest) GetQuestion() *string {
	return s.Question
}

func (s *ChatRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *ChatRequest) SetQuestion(v string) *ChatRequest {
	s.Question = &v
	return s
}

func (s *ChatRequest) SetSessionId(v string) *ChatRequest {
	s.SessionId = &v
	return s
}

func (s *ChatRequest) Validate() error {
	return dara.Validate(s)
}
