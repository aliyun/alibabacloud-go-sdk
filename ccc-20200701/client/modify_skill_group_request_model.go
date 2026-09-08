// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySkillGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *ModifySkillGroupRequest
	GetDescription() *string
	SetDisplayName(v string) *ModifySkillGroupRequest
	GetDisplayName() *string
	SetInstanceId(v string) *ModifySkillGroupRequest
	GetInstanceId() *string
	SetSkillGroupId(v string) *ModifySkillGroupRequest
	GetSkillGroupId() *string
}

type ModifySkillGroupRequest struct {
	// Description of the skill group. This parameter is optional and defaults to empty.
	//
	// example:
	//
	// 云联络中心的新测试技能组。
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Display name of the skill group, which must be 1 to 64 characters in length.
	//
	// example:
	//
	// 新测试技能组
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// Instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Skill group ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// skillgroup@ccc-test
	SkillGroupId *string `json:"SkillGroupId,omitempty" xml:"SkillGroupId,omitempty"`
}

func (s ModifySkillGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifySkillGroupRequest) GoString() string {
	return s.String()
}

func (s *ModifySkillGroupRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifySkillGroupRequest) GetDisplayName() *string {
	return s.DisplayName
}

func (s *ModifySkillGroupRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifySkillGroupRequest) GetSkillGroupId() *string {
	return s.SkillGroupId
}

func (s *ModifySkillGroupRequest) SetDescription(v string) *ModifySkillGroupRequest {
	s.Description = &v
	return s
}

func (s *ModifySkillGroupRequest) SetDisplayName(v string) *ModifySkillGroupRequest {
	s.DisplayName = &v
	return s
}

func (s *ModifySkillGroupRequest) SetInstanceId(v string) *ModifySkillGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifySkillGroupRequest) SetSkillGroupId(v string) *ModifySkillGroupRequest {
	s.SkillGroupId = &v
	return s
}

func (s *ModifySkillGroupRequest) Validate() error {
	return dara.Validate(s)
}
