// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReleaseAgentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentId(v string) *ReleaseAgentRequest
	GetAgentId() *string
}

type ReleaseAgentRequest struct {
	// example:
	//
	// 1
	AgentId *string `json:"agentId,omitempty" xml:"agentId,omitempty"`
}

func (s ReleaseAgentRequest) String() string {
	return dara.Prettify(s)
}

func (s ReleaseAgentRequest) GoString() string {
	return s.String()
}

func (s *ReleaseAgentRequest) GetAgentId() *string {
	return s.AgentId
}

func (s *ReleaseAgentRequest) SetAgentId(v string) *ReleaseAgentRequest {
	s.AgentId = &v
	return s
}

func (s *ReleaseAgentRequest) Validate() error {
	return dara.Validate(s)
}
