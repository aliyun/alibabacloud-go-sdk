// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateAgentDataSemanticsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentName(v string) *GenerateAgentDataSemanticsRequest
	GetAgentName() *string
}

type GenerateAgentDataSemanticsRequest struct {
	// The agent name. If no current official version exists or all four knowledge categories are empty, a first-time generation is performed. If at least one knowledge category is non-empty, only Text is regenerated while existing Metrics, Joins, and Examples are retained. The caller cannot specify the generation mode.
	//
	// This parameter is required.
	//
	// example:
	//
	// bakehouse_agent
	AgentName *string `json:"AgentName,omitempty" xml:"AgentName,omitempty"`
}

func (s GenerateAgentDataSemanticsRequest) String() string {
	return dara.Prettify(s)
}

func (s GenerateAgentDataSemanticsRequest) GoString() string {
	return s.String()
}

func (s *GenerateAgentDataSemanticsRequest) GetAgentName() *string {
	return s.AgentName
}

func (s *GenerateAgentDataSemanticsRequest) SetAgentName(v string) *GenerateAgentDataSemanticsRequest {
	s.AgentName = &v
	return s
}

func (s *GenerateAgentDataSemanticsRequest) Validate() error {
	return dara.Validate(s)
}
