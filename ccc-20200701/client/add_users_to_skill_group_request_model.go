// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddUsersToSkillGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *AddUsersToSkillGroupRequest
	GetInstanceId() *string
	SetSkillGroupId(v string) *AddUsersToSkillGroupRequest
	GetSkillGroupId() *string
	SetUserSkillLevelList(v string) *AddUsersToSkillGroupRequest
	GetUserSkillLevelList() *string
}

type AddUsersToSkillGroupRequest struct {
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
	// test1@ccc-test
	SkillGroupId *string `json:"SkillGroupId,omitempty" xml:"SkillGroupId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// [{"userId":"agent1@ccc-test","skillLevel":10},{"userId":"agent2@ccc-test","skillLevel":10}]
	UserSkillLevelList *string `json:"UserSkillLevelList,omitempty" xml:"UserSkillLevelList,omitempty"`
}

func (s AddUsersToSkillGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s AddUsersToSkillGroupRequest) GoString() string {
	return s.String()
}

func (s *AddUsersToSkillGroupRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *AddUsersToSkillGroupRequest) GetSkillGroupId() *string {
	return s.SkillGroupId
}

func (s *AddUsersToSkillGroupRequest) GetUserSkillLevelList() *string {
	return s.UserSkillLevelList
}

func (s *AddUsersToSkillGroupRequest) SetInstanceId(v string) *AddUsersToSkillGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *AddUsersToSkillGroupRequest) SetSkillGroupId(v string) *AddUsersToSkillGroupRequest {
	s.SkillGroupId = &v
	return s
}

func (s *AddUsersToSkillGroupRequest) SetUserSkillLevelList(v string) *AddUsersToSkillGroupRequest {
	s.UserSkillLevelList = &v
	return s
}

func (s *AddUsersToSkillGroupRequest) Validate() error {
	return dara.Validate(s)
}
