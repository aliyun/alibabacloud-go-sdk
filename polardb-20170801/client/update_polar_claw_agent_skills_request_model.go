// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePolarClawAgentSkillsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentId(v string) *UpdatePolarClawAgentSkillsRequest
	GetAgentId() *string
	SetApplicationId(v string) *UpdatePolarClawAgentSkillsRequest
	GetApplicationId() *string
	SetSkills(v []*string) *UpdatePolarClawAgentSkillsRequest
	GetSkills() []*string
}

type UpdatePolarClawAgentSkillsRequest struct {
	// Agent ID
	//
	// This parameter is required.
	//
	// example:
	//
	// work
	AgentId *string `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	// The application ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pa-xxx
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The skill allowlist. A value of null indicates that all skills are allowed.
	//
	// example:
	//
	// ["alibacloud-rds-copilot"]
	Skills []*string `json:"Skills,omitempty" xml:"Skills,omitempty" type:"Repeated"`
}

func (s UpdatePolarClawAgentSkillsRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdatePolarClawAgentSkillsRequest) GoString() string {
	return s.String()
}

func (s *UpdatePolarClawAgentSkillsRequest) GetAgentId() *string {
	return s.AgentId
}

func (s *UpdatePolarClawAgentSkillsRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *UpdatePolarClawAgentSkillsRequest) GetSkills() []*string {
	return s.Skills
}

func (s *UpdatePolarClawAgentSkillsRequest) SetAgentId(v string) *UpdatePolarClawAgentSkillsRequest {
	s.AgentId = &v
	return s
}

func (s *UpdatePolarClawAgentSkillsRequest) SetApplicationId(v string) *UpdatePolarClawAgentSkillsRequest {
	s.ApplicationId = &v
	return s
}

func (s *UpdatePolarClawAgentSkillsRequest) SetSkills(v []*string) *UpdatePolarClawAgentSkillsRequest {
	s.Skills = v
	return s
}

func (s *UpdatePolarClawAgentSkillsRequest) Validate() error {
	return dara.Validate(s)
}
