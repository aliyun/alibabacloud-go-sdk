// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySkillLevelsOfUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ModifySkillLevelsOfUserRequest
	GetInstanceId() *string
	SetSkillLevelList(v string) *ModifySkillLevelsOfUserRequest
	GetSkillLevelList() *string
	SetUserId(v string) *ModifySkillLevelsOfUserRequest
	GetUserId() *string
}

type ModifySkillLevelsOfUserRequest struct {
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
	// [{"skillGroupId":"skillgroup1@ccc-test","skillLevel":1},{"skillGroupId":"skillgroup2@ccc-test","skillLevel":10}]
	SkillLevelList *string `json:"SkillLevelList,omitempty" xml:"SkillLevelList,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// agent@ccc-test
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ModifySkillLevelsOfUserRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifySkillLevelsOfUserRequest) GoString() string {
	return s.String()
}

func (s *ModifySkillLevelsOfUserRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifySkillLevelsOfUserRequest) GetSkillLevelList() *string {
	return s.SkillLevelList
}

func (s *ModifySkillLevelsOfUserRequest) GetUserId() *string {
	return s.UserId
}

func (s *ModifySkillLevelsOfUserRequest) SetInstanceId(v string) *ModifySkillLevelsOfUserRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifySkillLevelsOfUserRequest) SetSkillLevelList(v string) *ModifySkillLevelsOfUserRequest {
	s.SkillLevelList = &v
	return s
}

func (s *ModifySkillLevelsOfUserRequest) SetUserId(v string) *ModifySkillLevelsOfUserRequest {
	s.UserId = &v
	return s
}

func (s *ModifySkillLevelsOfUserRequest) Validate() error {
	return dara.Validate(s)
}
