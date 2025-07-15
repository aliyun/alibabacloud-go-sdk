// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveUsersFromSkillGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *RemoveUsersFromSkillGroupRequest
	GetInstanceId() *string
	SetSkillGroupId(v string) *RemoveUsersFromSkillGroupRequest
	GetSkillGroupId() *string
	SetUserIdList(v string) *RemoveUsersFromSkillGroupRequest
	GetUserIdList() *string
}

type RemoveUsersFromSkillGroupRequest struct {
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
	// ["agent1@ccc-test","agent2@ccc-test"]
	UserIdList *string `json:"UserIdList,omitempty" xml:"UserIdList,omitempty"`
}

func (s RemoveUsersFromSkillGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveUsersFromSkillGroupRequest) GoString() string {
	return s.String()
}

func (s *RemoveUsersFromSkillGroupRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *RemoveUsersFromSkillGroupRequest) GetSkillGroupId() *string {
	return s.SkillGroupId
}

func (s *RemoveUsersFromSkillGroupRequest) GetUserIdList() *string {
	return s.UserIdList
}

func (s *RemoveUsersFromSkillGroupRequest) SetInstanceId(v string) *RemoveUsersFromSkillGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *RemoveUsersFromSkillGroupRequest) SetSkillGroupId(v string) *RemoveUsersFromSkillGroupRequest {
	s.SkillGroupId = &v
	return s
}

func (s *RemoveUsersFromSkillGroupRequest) SetUserIdList(v string) *RemoveUsersFromSkillGroupRequest {
	s.UserIdList = &v
	return s
}

func (s *RemoveUsersFromSkillGroupRequest) Validate() error {
	return dara.Validate(s)
}
