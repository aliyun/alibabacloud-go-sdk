// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAgentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentId(v string) *DeleteAgentRequest
	GetAgentId() *string
}

type DeleteAgentRequest struct {
	// example:
	//
	// 840016700254633984
	AgentId *string `json:"agentId,omitempty" xml:"agentId,omitempty"`
}

func (s DeleteAgentRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAgentRequest) GoString() string {
	return s.String()
}

func (s *DeleteAgentRequest) GetAgentId() *string {
	return s.AgentId
}

func (s *DeleteAgentRequest) SetAgentId(v string) *DeleteAgentRequest {
	s.AgentId = &v
	return s
}

func (s *DeleteAgentRequest) Validate() error {
	return dara.Validate(s)
}
