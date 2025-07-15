// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddSkillGroupsToUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *AddSkillGroupsToUserRequest
	GetInstanceId() *string
	SetSkillLevelList(v string) *AddSkillGroupsToUserRequest
	GetSkillLevelList() *string
	SetUserId(v string) *AddSkillGroupsToUserRequest
	GetUserId() *string
}

type AddSkillGroupsToUserRequest struct {
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
	// [{"skillGroupId":"test1@ccc-test","skillLevel":5},{"skillGroupId":"test2@ccc-test","skillLevel":5}]
	SkillLevelList *string `json:"SkillLevelList,omitempty" xml:"SkillLevelList,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// agent@ccc-test
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s AddSkillGroupsToUserRequest) String() string {
	return dara.Prettify(s)
}

func (s AddSkillGroupsToUserRequest) GoString() string {
	return s.String()
}

func (s *AddSkillGroupsToUserRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *AddSkillGroupsToUserRequest) GetSkillLevelList() *string {
	return s.SkillLevelList
}

func (s *AddSkillGroupsToUserRequest) GetUserId() *string {
	return s.UserId
}

func (s *AddSkillGroupsToUserRequest) SetInstanceId(v string) *AddSkillGroupsToUserRequest {
	s.InstanceId = &v
	return s
}

func (s *AddSkillGroupsToUserRequest) SetSkillLevelList(v string) *AddSkillGroupsToUserRequest {
	s.SkillLevelList = &v
	return s
}

func (s *AddSkillGroupsToUserRequest) SetUserId(v string) *AddSkillGroupsToUserRequest {
	s.UserId = &v
	return s
}

func (s *AddSkillGroupsToUserRequest) Validate() error {
	return dara.Validate(s)
}
