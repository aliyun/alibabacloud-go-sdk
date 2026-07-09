// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEvaluatorSkillRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentSpace(v string) *GetEvaluatorSkillRequest
	GetAgentSpace() *string
	SetVersion(v string) *GetEvaluatorSkillRequest
	GetVersion() *string
}

type GetEvaluatorSkillRequest struct {
	// The AgentSpace name.
	//
	// This parameter is required.
	//
	// example:
	//
	// prod-agentspace
	AgentSpace *string `json:"agentSpace,omitempty" xml:"agentSpace,omitempty"`
	// The skill version.
	//
	// example:
	//
	// 1782816000000
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s GetEvaluatorSkillRequest) String() string {
	return dara.Prettify(s)
}

func (s GetEvaluatorSkillRequest) GoString() string {
	return s.String()
}

func (s *GetEvaluatorSkillRequest) GetAgentSpace() *string {
	return s.AgentSpace
}

func (s *GetEvaluatorSkillRequest) GetVersion() *string {
	return s.Version
}

func (s *GetEvaluatorSkillRequest) SetAgentSpace(v string) *GetEvaluatorSkillRequest {
	s.AgentSpace = &v
	return s
}

func (s *GetEvaluatorSkillRequest) SetVersion(v string) *GetEvaluatorSkillRequest {
	s.Version = &v
	return s
}

func (s *GetEvaluatorSkillRequest) Validate() error {
	return dara.Validate(s)
}
