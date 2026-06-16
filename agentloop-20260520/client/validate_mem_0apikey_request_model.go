// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iValidateMem0APIKeyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentSpace(v string) *ValidateMem0APIKeyRequest
	GetAgentSpace() *string
}

type ValidateMem0APIKeyRequest struct {
	// example:
	//
	// my-agent-space
	AgentSpace *string `json:"agentSpace,omitempty" xml:"agentSpace,omitempty"`
}

func (s ValidateMem0APIKeyRequest) String() string {
	return dara.Prettify(s)
}

func (s ValidateMem0APIKeyRequest) GoString() string {
	return s.String()
}

func (s *ValidateMem0APIKeyRequest) GetAgentSpace() *string {
	return s.AgentSpace
}

func (s *ValidateMem0APIKeyRequest) SetAgentSpace(v string) *ValidateMem0APIKeyRequest {
	s.AgentSpace = &v
	return s
}

func (s *ValidateMem0APIKeyRequest) Validate() error {
	return dara.Validate(s)
}
