// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAgentCardRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnv(v string) *AgentCardRequest
	GetEnv() *string
}

type AgentCardRequest struct {
	// example:
	//
	// a2a
	Env *string `json:"Env,omitempty" xml:"Env,omitempty"`
}

func (s AgentCardRequest) String() string {
	return dara.Prettify(s)
}

func (s AgentCardRequest) GoString() string {
	return s.String()
}

func (s *AgentCardRequest) GetEnv() *string {
	return s.Env
}

func (s *AgentCardRequest) SetEnv(v string) *AgentCardRequest {
	s.Env = &v
	return s
}

func (s *AgentCardRequest) Validate() error {
	return dara.Validate(s)
}
