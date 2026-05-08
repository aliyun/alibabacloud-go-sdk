// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInteractTextRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentId(v string) *InteractTextRequest
	GetAgentId() *string
	SetContent(v string) *InteractTextRequest
	GetContent() *string
	SetSessionId(v string) *InteractTextRequest
	GetSessionId() *string
}

type InteractTextRequest struct {
	// example:
	//
	// 1000222
	AgentId *string `json:"agentId,omitempty" xml:"agentId,omitempty"`
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// example:
	//
	// 144285195534941
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
}

func (s InteractTextRequest) String() string {
	return dara.Prettify(s)
}

func (s InteractTextRequest) GoString() string {
	return s.String()
}

func (s *InteractTextRequest) GetAgentId() *string {
	return s.AgentId
}

func (s *InteractTextRequest) GetContent() *string {
	return s.Content
}

func (s *InteractTextRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *InteractTextRequest) SetAgentId(v string) *InteractTextRequest {
	s.AgentId = &v
	return s
}

func (s *InteractTextRequest) SetContent(v string) *InteractTextRequest {
	s.Content = &v
	return s
}

func (s *InteractTextRequest) SetSessionId(v string) *InteractTextRequest {
	s.SessionId = &v
	return s
}

func (s *InteractTextRequest) Validate() error {
	return dara.Validate(s)
}
