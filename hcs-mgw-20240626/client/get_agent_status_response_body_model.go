// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAgentStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetImportAgentStatus(v *GetAgentStatusResp) *GetAgentStatusResponseBody
	GetImportAgentStatus() *GetAgentStatusResp
}

type GetAgentStatusResponseBody struct {
	// The details for obtaining the status of the agent.
	ImportAgentStatus *GetAgentStatusResp `json:"ImportAgentStatus,omitempty" xml:"ImportAgentStatus,omitempty"`
}

func (s GetAgentStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAgentStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetAgentStatusResponseBody) GetImportAgentStatus() *GetAgentStatusResp {
	return s.ImportAgentStatus
}

func (s *GetAgentStatusResponseBody) SetImportAgentStatus(v *GetAgentStatusResp) *GetAgentStatusResponseBody {
	s.ImportAgentStatus = v
	return s
}

func (s *GetAgentStatusResponseBody) Validate() error {
	if s.ImportAgentStatus != nil {
		if err := s.ImportAgentStatus.Validate(); err != nil {
			return err
		}
	}
	return nil
}
