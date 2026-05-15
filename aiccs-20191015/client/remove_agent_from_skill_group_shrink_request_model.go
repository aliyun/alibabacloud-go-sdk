// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveAgentFromSkillGroupShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentIdsShrink(v string) *RemoveAgentFromSkillGroupShrinkRequest
	GetAgentIdsShrink() *string
	SetInstanceId(v string) *RemoveAgentFromSkillGroupShrinkRequest
	GetInstanceId() *string
	SetSkillGroupId(v int64) *RemoveAgentFromSkillGroupShrinkRequest
	GetSkillGroupId() *int64
}

type RemoveAgentFromSkillGroupShrinkRequest struct {
	// This parameter is required.
	AgentIdsShrink *string `json:"AgentIds,omitempty" xml:"AgentIds,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// agent_***
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1146****
	SkillGroupId *int64 `json:"SkillGroupId,omitempty" xml:"SkillGroupId,omitempty"`
}

func (s RemoveAgentFromSkillGroupShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveAgentFromSkillGroupShrinkRequest) GoString() string {
	return s.String()
}

func (s *RemoveAgentFromSkillGroupShrinkRequest) GetAgentIdsShrink() *string {
	return s.AgentIdsShrink
}

func (s *RemoveAgentFromSkillGroupShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *RemoveAgentFromSkillGroupShrinkRequest) GetSkillGroupId() *int64 {
	return s.SkillGroupId
}

func (s *RemoveAgentFromSkillGroupShrinkRequest) SetAgentIdsShrink(v string) *RemoveAgentFromSkillGroupShrinkRequest {
	s.AgentIdsShrink = &v
	return s
}

func (s *RemoveAgentFromSkillGroupShrinkRequest) SetInstanceId(v string) *RemoveAgentFromSkillGroupShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *RemoveAgentFromSkillGroupShrinkRequest) SetSkillGroupId(v int64) *RemoveAgentFromSkillGroupShrinkRequest {
	s.SkillGroupId = &v
	return s
}

func (s *RemoveAgentFromSkillGroupShrinkRequest) Validate() error {
	return dara.Validate(s)
}
