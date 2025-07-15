// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyUserLevelsOfSkillGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ModifyUserLevelsOfSkillGroupRequest
	GetInstanceId() *string
	SetSkillGroupId(v string) *ModifyUserLevelsOfSkillGroupRequest
	GetSkillGroupId() *string
	SetUserLevelList(v string) *ModifyUserLevelsOfSkillGroupRequest
	GetUserLevelList() *string
}

type ModifyUserLevelsOfSkillGroupRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// skillgroup@ccc-test
	SkillGroupId *string `json:"SkillGroupId,omitempty" xml:"SkillGroupId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// [{"userId":"agent1@ccc-test","skillLevel":1},{"userId":"agent2@ccc-test","skillLevel":10}]
	UserLevelList *string `json:"UserLevelList,omitempty" xml:"UserLevelList,omitempty"`
}

func (s ModifyUserLevelsOfSkillGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyUserLevelsOfSkillGroupRequest) GoString() string {
	return s.String()
}

func (s *ModifyUserLevelsOfSkillGroupRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyUserLevelsOfSkillGroupRequest) GetSkillGroupId() *string {
	return s.SkillGroupId
}

func (s *ModifyUserLevelsOfSkillGroupRequest) GetUserLevelList() *string {
	return s.UserLevelList
}

func (s *ModifyUserLevelsOfSkillGroupRequest) SetInstanceId(v string) *ModifyUserLevelsOfSkillGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyUserLevelsOfSkillGroupRequest) SetSkillGroupId(v string) *ModifyUserLevelsOfSkillGroupRequest {
	s.SkillGroupId = &v
	return s
}

func (s *ModifyUserLevelsOfSkillGroupRequest) SetUserLevelList(v string) *ModifyUserLevelsOfSkillGroupRequest {
	s.UserLevelList = &v
	return s
}

func (s *ModifyUserLevelsOfSkillGroupRequest) Validate() error {
	return dara.Validate(s)
}
