// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClearAgentInstanceMemoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentInstanceId(v string) *ClearAgentInstanceMemoryRequest
	GetAgentInstanceId() *string
}

type ClearAgentInstanceMemoryRequest struct {
	// The ID of the agent instance whose memory you want to clear. This parameter is required and can be obtained from the response of ListAgentInstanceMemory.
	//
	// example:
	//
	// agi-08vm****
	AgentInstanceId *string `json:"AgentInstanceId,omitempty" xml:"AgentInstanceId,omitempty"`
}

func (s ClearAgentInstanceMemoryRequest) String() string {
	return dara.Prettify(s)
}

func (s ClearAgentInstanceMemoryRequest) GoString() string {
	return s.String()
}

func (s *ClearAgentInstanceMemoryRequest) GetAgentInstanceId() *string {
	return s.AgentInstanceId
}

func (s *ClearAgentInstanceMemoryRequest) SetAgentInstanceId(v string) *ClearAgentInstanceMemoryRequest {
	s.AgentInstanceId = &v
	return s
}

func (s *ClearAgentInstanceMemoryRequest) Validate() error {
	return dara.Validate(s)
}
