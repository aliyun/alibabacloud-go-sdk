// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCopilotSessionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentId(v string) *DeleteCopilotSessionRequest
	GetAgentId() *string
	SetSessionId(v string) *DeleteCopilotSessionRequest
	GetSessionId() *string
}

type DeleteCopilotSessionRequest struct {
	// example:
	//
	// sre_agent
	AgentId *string `json:"agentId,omitempty" xml:"agentId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// c356bc310e4345c3a602a2912322ad21
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
}

func (s DeleteCopilotSessionRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteCopilotSessionRequest) GoString() string {
	return s.String()
}

func (s *DeleteCopilotSessionRequest) GetAgentId() *string {
	return s.AgentId
}

func (s *DeleteCopilotSessionRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *DeleteCopilotSessionRequest) SetAgentId(v string) *DeleteCopilotSessionRequest {
	s.AgentId = &v
	return s
}

func (s *DeleteCopilotSessionRequest) SetSessionId(v string) *DeleteCopilotSessionRequest {
	s.SessionId = &v
	return s
}

func (s *DeleteCopilotSessionRequest) Validate() error {
	return dara.Validate(s)
}
