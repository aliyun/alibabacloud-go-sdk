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
	// A list of agent IDs.
	//
	// This parameter is required.
	AgentIdsShrink *string `json:"AgentIds,omitempty" xml:"AgentIds,omitempty"`
	// The Artificial Intelligence Cloud Call Service (AICCS) instance ID.
	//
	// You can obtain it from **Instance Management*	- in the left-side navigation pane of the [Artificial Intelligence Cloud Call Service console](https://aiccs.console.aliyun.com/overview).
	//
	// This parameter is required.
	//
	// example:
	//
	// ccc_xp_pre-cn-78v1gnp97002
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Skill group ID.
	//
	// You can invoke the [QuerySkillGroups](https://help.aliyun.com/zh/aiccs/developer-reference/api-aiccs-2019-10-15-queryskillgroups) API and view the **SkillGroupId*	- in the response parameters to obtain the skill group ID.
	//
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
