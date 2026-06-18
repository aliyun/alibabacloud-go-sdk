// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveAgentFromSkillGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentIds(v []*int64) *RemoveAgentFromSkillGroupRequest
	GetAgentIds() []*int64
	SetInstanceId(v string) *RemoveAgentFromSkillGroupRequest
	GetInstanceId() *string
	SetSkillGroupId(v int64) *RemoveAgentFromSkillGroupRequest
	GetSkillGroupId() *int64
}

type RemoveAgentFromSkillGroupRequest struct {
	// A list of agent IDs.
	//
	// This parameter is required.
	AgentIds []*int64 `json:"AgentIds,omitempty" xml:"AgentIds,omitempty" type:"Repeated"`
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

func (s RemoveAgentFromSkillGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveAgentFromSkillGroupRequest) GoString() string {
	return s.String()
}

func (s *RemoveAgentFromSkillGroupRequest) GetAgentIds() []*int64 {
	return s.AgentIds
}

func (s *RemoveAgentFromSkillGroupRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *RemoveAgentFromSkillGroupRequest) GetSkillGroupId() *int64 {
	return s.SkillGroupId
}

func (s *RemoveAgentFromSkillGroupRequest) SetAgentIds(v []*int64) *RemoveAgentFromSkillGroupRequest {
	s.AgentIds = v
	return s
}

func (s *RemoveAgentFromSkillGroupRequest) SetInstanceId(v string) *RemoveAgentFromSkillGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *RemoveAgentFromSkillGroupRequest) SetSkillGroupId(v int64) *RemoveAgentFromSkillGroupRequest {
	s.SkillGroupId = &v
	return s
}

func (s *RemoveAgentFromSkillGroupRequest) Validate() error {
	return dara.Validate(s)
}
