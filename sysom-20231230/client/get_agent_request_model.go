// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAgentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentId(v string) *GetAgentRequest
	GetAgentId() *string
}

type GetAgentRequest struct {
	AgentId *string `json:"agent_id,omitempty" xml:"agent_id,omitempty"`
}

func (s GetAgentRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAgentRequest) GoString() string {
	return s.String()
}

func (s *GetAgentRequest) GetAgentId() *string {
	return s.AgentId
}

func (s *GetAgentRequest) SetAgentId(v string) *GetAgentRequest {
	s.AgentId = &v
	return s
}

func (s *GetAgentRequest) Validate() error {
	return dara.Validate(s)
}
