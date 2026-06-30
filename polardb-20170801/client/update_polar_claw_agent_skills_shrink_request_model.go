// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePolarClawAgentSkillsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentId(v string) *UpdatePolarClawAgentSkillsShrinkRequest
	GetAgentId() *string
	SetApplicationId(v string) *UpdatePolarClawAgentSkillsShrinkRequest
	GetApplicationId() *string
	SetSkillsShrink(v string) *UpdatePolarClawAgentSkillsShrinkRequest
	GetSkillsShrink() *string
}

type UpdatePolarClawAgentSkillsShrinkRequest struct {
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
	SkillsShrink *string `json:"Skills,omitempty" xml:"Skills,omitempty"`
}

func (s UpdatePolarClawAgentSkillsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdatePolarClawAgentSkillsShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdatePolarClawAgentSkillsShrinkRequest) GetAgentId() *string {
	return s.AgentId
}

func (s *UpdatePolarClawAgentSkillsShrinkRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *UpdatePolarClawAgentSkillsShrinkRequest) GetSkillsShrink() *string {
	return s.SkillsShrink
}

func (s *UpdatePolarClawAgentSkillsShrinkRequest) SetAgentId(v string) *UpdatePolarClawAgentSkillsShrinkRequest {
	s.AgentId = &v
	return s
}

func (s *UpdatePolarClawAgentSkillsShrinkRequest) SetApplicationId(v string) *UpdatePolarClawAgentSkillsShrinkRequest {
	s.ApplicationId = &v
	return s
}

func (s *UpdatePolarClawAgentSkillsShrinkRequest) SetSkillsShrink(v string) *UpdatePolarClawAgentSkillsShrinkRequest {
	s.SkillsShrink = &v
	return s
}

func (s *UpdatePolarClawAgentSkillsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
