// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAgentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetImportAgent(v *GetAgentResp) *GetAgentResponseBody
	GetImportAgent() *GetAgentResp
}

type GetAgentResponseBody struct {
	// The details for obtaining the details of the agent.
	ImportAgent *GetAgentResp `json:"ImportAgent,omitempty" xml:"ImportAgent,omitempty"`
}

func (s GetAgentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAgentResponseBody) GoString() string {
	return s.String()
}

func (s *GetAgentResponseBody) GetImportAgent() *GetAgentResp {
	return s.ImportAgent
}

func (s *GetAgentResponseBody) SetImportAgent(v *GetAgentResp) *GetAgentResponseBody {
	s.ImportAgent = v
	return s
}

func (s *GetAgentResponseBody) Validate() error {
	if s.ImportAgent != nil {
		if err := s.ImportAgent.Validate(); err != nil {
			return err
		}
	}
	return nil
}
