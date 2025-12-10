// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAgentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImportAgent(v *CreateAgentInfo) *CreateAgentRequest
	GetImportAgent() *CreateAgentInfo
}

type CreateAgentRequest struct {
	// The details for creating the agent.
	ImportAgent *CreateAgentInfo `json:"ImportAgent,omitempty" xml:"ImportAgent,omitempty"`
}

func (s CreateAgentRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAgentRequest) GoString() string {
	return s.String()
}

func (s *CreateAgentRequest) GetImportAgent() *CreateAgentInfo {
	return s.ImportAgent
}

func (s *CreateAgentRequest) SetImportAgent(v *CreateAgentInfo) *CreateAgentRequest {
	s.ImportAgent = v
	return s
}

func (s *CreateAgentRequest) Validate() error {
	if s.ImportAgent != nil {
		if err := s.ImportAgent.Validate(); err != nil {
			return err
		}
	}
	return nil
}
