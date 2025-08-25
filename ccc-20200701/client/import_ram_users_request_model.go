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
