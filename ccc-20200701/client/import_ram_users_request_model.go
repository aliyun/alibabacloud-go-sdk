// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportRamUsersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ImportRamUsersRequest
	GetInstanceId() *string
	SetRamIdList(v string) *ImportRamUsersRequest
	GetRamIdList() *string
	SetRoleId(v string) *ImportRamUsersRequest
	GetRoleId() *string
	SetSkillLevelList(v string) *ImportRamUsersRequest
	GetSkillLevelList() *string
	SetWorkMode(v string) *ImportRamUsersRequest
	GetWorkMode() *string
}

type ImportRamUsersRequest struct {
	// Instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// List of RAM user IDs to add.
	//
	// This parameter is required.
	//
	// example:
	//
	// ["28036411123456****","29234301123456****"]
	RamIdList *string `json:"RamIdList,omitempty" xml:"RamIdList,omitempty"`
	// Role ID. This is the agent\\"s role in the instance after successful import. Roles include administrator, skill group leader, and agent.
	//
	// This parameter is required.
	//
	// example:
	//
	// Agent@ccc-test
	RoleId *string `json:"RoleId,omitempty" xml:"RoleId,omitempty"`
	// Skill level list for the skill group. This is a JSON array string. Each array element is an object with two fields: skillGroupId and skillLevel. skillGroupId is the ID of the skill group the agent wants to associate with. skillLevel is the agent\\"s skill level in that skill group. The range is 1-10. A smaller value indicates stronger business capability and the ability to handle more calls per unit of time.
	//
	// example:
	//
	// [{"skillGroupId":"skillgroup@ccc-test","skillLevel":5}]
	SkillLevelList *string `json:"SkillLevelList,omitempty" xml:"SkillLevelList,omitempty"`
	// Work mode.
	//
	// This parameter is required.
	//
	// example:
	//
	// ON_SITE
	WorkMode *string `json:"WorkMode,omitempty" xml:"WorkMode,omitempty"`
}

func (s ImportRamUsersRequest) String() string {
	return dara.Prettify(s)
}

func (s ImportRamUsersRequest) GoString() string {
	return s.String()
}

func (s *ImportRamUsersRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ImportRamUsersRequest) GetRamIdList() *string {
	return s.RamIdList
}

func (s *ImportRamUsersRequest) GetRoleId() *string {
	return s.RoleId
}

func (s *ImportRamUsersRequest) GetSkillLevelList() *string {
	return s.SkillLevelList
}

func (s *ImportRamUsersRequest) GetWorkMode() *string {
	return s.WorkMode
}

func (s *ImportRamUsersRequest) SetInstanceId(v string) *ImportRamUsersRequest {
	s.InstanceId = &v
	return s
}

func (s *ImportRamUsersRequest) SetRamIdList(v string) *ImportRamUsersRequest {
	s.RamIdList = &v
	return s
}

func (s *ImportRamUsersRequest) SetRoleId(v string) *ImportRamUsersRequest {
	s.RoleId = &v
	return s
}

func (s *ImportRamUsersRequest) SetSkillLevelList(v string) *ImportRamUsersRequest {
	s.SkillLevelList = &v
	return s
}

func (s *ImportRamUsersRequest) SetWorkMode(v string) *ImportRamUsersRequest {
	s.WorkMode = &v
	return s
}

func (s *ImportRamUsersRequest) Validate() error {
	return dara.Validate(s)
}
