// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAgentDataSemanticsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentName(v string) *GetAgentDataSemanticsRequest
	GetAgentName() *string
}

type GetAgentDataSemanticsRequest struct {
	// The name of the agent.
	//
	// This parameter is required.
	//
	// example:
	//
	// bakehouse_agent
	AgentName *string `json:"AgentName,omitempty" xml:"AgentName,omitempty"`
}

func (s GetAgentDataSemanticsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAgentDataSemanticsRequest) GoString() string {
	return s.String()
}

func (s *GetAgentDataSemanticsRequest) GetAgentName() *string {
	return s.AgentName
}

func (s *GetAgentDataSemanticsRequest) SetAgentName(v string) *GetAgentDataSemanticsRequest {
	s.AgentName = &v
	return s
}

func (s *GetAgentDataSemanticsRequest) Validate() error {
	return dara.Validate(s)
}
