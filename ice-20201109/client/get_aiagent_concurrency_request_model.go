// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAIAgentConcurrencyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAIAgentId(v string) *GetAIAgentConcurrencyRequest
	GetAIAgentId() *string
}

type GetAIAgentConcurrencyRequest struct {
	// The ID of the AI agent.
	//
	// This parameter is required.
	//
	// example:
	//
	// *******3b3d94abda22******
	AIAgentId *string `json:"AIAgentId,omitempty" xml:"AIAgentId,omitempty"`
}

func (s GetAIAgentConcurrencyRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAIAgentConcurrencyRequest) GoString() string {
	return s.String()
}

func (s *GetAIAgentConcurrencyRequest) GetAIAgentId() *string {
	return s.AIAgentId
}

func (s *GetAIAgentConcurrencyRequest) SetAIAgentId(v string) *GetAIAgentConcurrencyRequest {
	s.AIAgentId = &v
	return s
}

func (s *GetAIAgentConcurrencyRequest) Validate() error {
	return dara.Validate(s)
}
