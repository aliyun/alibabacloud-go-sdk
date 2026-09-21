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
	SetResume(v string) *ChatRequest
	GetResume() *string
	SetSessionId(v string) *ChatRequest
	GetSessionId() *string
	SetSummary(v string) *ChatRequest
	GetSummary() *string
}

type ChatRequest struct {
	// The agent ID used for the service. This parameter is optional. You can specify an agent generated after DAS Agent is enabled or an agent that you manually created. If this parameter is not specified, the default agent is used.
	//
	// example:
	//
	// ag-472T0DxtmjIxxxxx
	AgentId *string `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	// The message.
	//
	// example:
	//
	// {"id":"68fe0321-37fe-4c75-a118-b61b33156f6a","role":"user","content":[{"type":"text","text":"hello"}]}
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The user interaction feedback.
	//
	// example:
	//
	// [{"metadata":{"a2uiClientDataModel":{"decided":true,"interaction":{"request":{"allowFreeText":true,"question":"question?","options":[{"label":"162","value":"162"},{"label":"242","value":"242"},{"label":"243","value":"243"},{"label":"486","value":"486"}]},"kind":"choice","interruptId":"5857955e-4856-4ab3-969e-f8b4c484bda4","status":"pending"},"free_text":""},"a2uiAction":{"name":"das_interaction_response","context":{"value":"162"},"surfaceId":"render-choice-7c03eb4f927f42ff8c30b30a2cc0797d","sourceComponentId":"opts-0","timestamp":"2026-09-15T06:54:15.556Z"}},"payload":{"value":"162"},"interruptId":"5857955e-4856-4ab3-969e-f8b4c484bda4","status":"resolved"}]
	Resume *string `json:"Resume,omitempty" xml:"Resume,omitempty"`
	// The session ID in UUID string format. This parameter is optional. If this parameter is not specified, a new session is created by default. To maintain context across conversations, use the same session ID.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-xxxxxxxxxxxx
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// Specifies whether to output summary information.
	//
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

func (s *ChatRequest) GetResume() *string {
	return s.Resume
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

func (s *ChatRequest) SetResume(v string) *ChatRequest {
	s.Resume = &v
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
