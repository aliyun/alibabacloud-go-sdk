// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGenerateAgentDataSemanticsProgressRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentName(v string) *GetGenerateAgentDataSemanticsProgressRequest
	GetAgentName() *string
}

type GetGenerateAgentDataSemanticsProgressRequest struct {
	// The name of the agent.
	//
	// This parameter is required.
	//
	// example:
	//
	// bakehouse_agent
	AgentName *string `json:"AgentName,omitempty" xml:"AgentName,omitempty"`
}

func (s GetGenerateAgentDataSemanticsProgressRequest) String() string {
	return dara.Prettify(s)
}

func (s GetGenerateAgentDataSemanticsProgressRequest) GoString() string {
	return s.String()
}

func (s *GetGenerateAgentDataSemanticsProgressRequest) GetAgentName() *string {
	return s.AgentName
}

func (s *GetGenerateAgentDataSemanticsProgressRequest) SetAgentName(v string) *GetGenerateAgentDataSemanticsProgressRequest {
	s.AgentName = &v
	return s
}

func (s *GetGenerateAgentDataSemanticsProgressRequest) Validate() error {
	return dara.Validate(s)
}
