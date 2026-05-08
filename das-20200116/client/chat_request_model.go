// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChatRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentId(v string) *ChatRequest
	GetAgentId() *string
	SetMessage(v string) *ChatRequest
	GetMessage() *string
	SetSessionId(v string) *ChatRequest
	GetSessionId() *string
	SetSummary(v string) *ChatRequest
	GetSummary() *string
}

type ChatRequest struct {
	// example:
	//
	// ag-472T0DxtmjIxxxxx
	AgentId *string `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// {"id":"68fe0321-37fe-4c75-a118-b61b33156f6a","role":"user","content":[{"type":"text","text":"hello"}]}
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 123e4567-e89b-12d3-a456-xxxxxxxxxxxx
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// example:
	//
	// false
	Summary *string `json:"Summary,omitempty" xml:"Summary,omitempty"`
}

func (s ChatRequest) String() string {
	return dara.Prettify(s)
}

func (s ChatRequest) GoString() string {
	return s.String()
}

func (s *ChatRequest) GetAgentId() *string {
	return s.AgentId
}

func (s *ChatRequest) GetMessage() *string {
	return s.Message
}

func (s *ChatRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *ChatRequest) GetSummary() *string {
	return s.Summary
}

func (s *ChatRequest) SetAgentId(v string) *ChatRequest {
	s.AgentId = &v
	return s
}

func (s *ChatRequest) SetMessage(v string) *ChatRequest {
	s.Message = &v
	return s
}

func (s *ChatRequest) SetSessionId(v string) *ChatRequest {
	s.SessionId = &v
	return s
}

func (s *ChatRequest) SetSummary(v string) *ChatRequest {
	s.Summary = &v
	return s
}

func (s *ChatRequest) Validate() error {
	return dara.Validate(s)
}
