// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEvaluatorSkillRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentSpace(v string) *DeleteEvaluatorSkillRequest
	GetAgentSpace() *string
}

type DeleteEvaluatorSkillRequest struct {
	// The AgentSpace name.
	//
	// This parameter is required.
	//
	// example:
	//
	// prod-agentspace
	AgentSpace *string `json:"agentSpace,omitempty" xml:"agentSpace,omitempty"`
}

func (s DeleteEvaluatorSkillRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteEvaluatorSkillRequest) GoString() string {
	return s.String()
}

func (s *DeleteEvaluatorSkillRequest) GetAgentSpace() *string {
	return s.AgentSpace
}

func (s *DeleteEvaluatorSkillRequest) SetAgentSpace(v string) *DeleteEvaluatorSkillRequest {
	s.AgentSpace = &v
	return s
}

func (s *DeleteEvaluatorSkillRequest) Validate() error {
	return dara.Validate(s)
}
