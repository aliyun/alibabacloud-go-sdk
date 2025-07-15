// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssignUsersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *AssignUsersRequest
	GetInstanceId() *string
	SetRamIdList(v string) *AssignUsersRequest
	GetRamIdList() *string
	SetRoleId(v string) *AssignUsersRequest
	GetRoleId() *string
	SetSkillLevelList(v string) *AssignUsersRequest
	GetSkillLevelList() *string
	SetWorkMode(v string) *AssignUsersRequest
	GetWorkMode() *string
}

type AssignUsersRequest struct {
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
	// ["28036411123456****","29234301123456****"]
	RamIdList *string `json:"RamIdList,omitempty" xml:"RamIdList,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Agent@ccc-test
	RoleId *string `json:"RoleId,omitempty" xml:"RoleId,omitempty"`
	// example:
	//
	// [{"skillGroupId":"skillgroup@ccc-test","skillLevel":5}]
	SkillLevelList *string `json:"SkillLevelList,omitempty" xml:"SkillLevelList,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ON_SITE
	WorkMode *string `json:"WorkMode,omitempty" xml:"WorkMode,omitempty"`
}

func (s AssignUsersRequest) String() string {
	return dara.Prettify(s)
}

func (s AssignUsersRequest) GoString() string {
	return s.String()
}

func (s *AssignUsersRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *AssignUsersRequest) GetRamIdList() *string {
	return s.RamIdList
}

func (s *AssignUsersRequest) GetRoleId() *string {
	return s.RoleId
}

func (s *AssignUsersRequest) GetSkillLevelList() *string {
	return s.SkillLevelList
}

func (s *AssignUsersRequest) GetWorkMode() *string {
	return s.WorkMode
}

func (s *AssignUsersRequest) SetInstanceId(v string) *AssignUsersRequest {
	s.InstanceId = &v
	return s
}

func (s *AssignUsersRequest) SetRamIdList(v string) *AssignUsersRequest {
	s.RamIdList = &v
	return s
}

func (s *AssignUsersRequest) SetRoleId(v string) *AssignUsersRequest {
	s.RoleId = &v
	return s
}

func (s *AssignUsersRequest) SetSkillLevelList(v string) *AssignUsersRequest {
	s.SkillLevelList = &v
	return s
}

func (s *AssignUsersRequest) SetWorkMode(v string) *AssignUsersRequest {
	s.WorkMode = &v
	return s
}

func (s *AssignUsersRequest) Validate() error {
	return dara.Validate(s)
}
