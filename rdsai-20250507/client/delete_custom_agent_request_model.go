// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCustomAgentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustomAgentId(v string) *DeleteCustomAgentRequest
	GetCustomAgentId() *string
}

type DeleteCustomAgentRequest struct {
	// AgentId。
	//
	// example:
	//
	// ebe44453-3b41-4c74-94d1-01d088d7****
	CustomAgentId *string `json:"CustomAgentId,omitempty" xml:"CustomAgentId,omitempty"`
}

func (s DeleteCustomAgentRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteCustomAgentRequest) GoString() string {
	return s.String()
}

func (s *DeleteCustomAgentRequest) GetCustomAgentId() *string {
	return s.CustomAgentId
}

func (s *DeleteCustomAgentRequest) SetCustomAgentId(v string) *DeleteCustomAgentRequest {
	s.CustomAgentId = &v
	return s
}

func (s *DeleteCustomAgentRequest) Validate() error {
	return dara.Validate(s)
}
