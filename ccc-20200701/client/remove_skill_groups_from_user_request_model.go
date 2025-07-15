// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveSkillGroupsFromUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *RemoveSkillGroupsFromUserRequest
	GetInstanceId() *string
	SetSkillGroupIdList(v string) *RemoveSkillGroupsFromUserRequest
	GetSkillGroupIdList() *string
	SetUserId(v string) *RemoveSkillGroupsFromUserRequest
	GetUserId() *string
}

type RemoveSkillGroupsFromUserRequest struct {
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
	// ["skillgroup1@ccc-test","skillgroup2@ccc-test"]
	SkillGroupIdList *string `json:"SkillGroupIdList,omitempty" xml:"SkillGroupIdList,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// agent@ccc-test
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s RemoveSkillGroupsFromUserRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveSkillGroupsFromUserRequest) GoString() string {
	return s.String()
}

func (s *RemoveSkillGroupsFromUserRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *RemoveSkillGroupsFromUserRequest) GetSkillGroupIdList() *string {
	return s.SkillGroupIdList
}

func (s *RemoveSkillGroupsFromUserRequest) GetUserId() *string {
	return s.UserId
}

func (s *RemoveSkillGroupsFromUserRequest) SetInstanceId(v string) *RemoveSkillGroupsFromUserRequest {
	s.InstanceId = &v
	return s
}

func (s *RemoveSkillGroupsFromUserRequest) SetSkillGroupIdList(v string) *RemoveSkillGroupsFromUserRequest {
	s.SkillGroupIdList = &v
	return s
}

func (s *RemoveSkillGroupsFromUserRequest) SetUserId(v string) *RemoveSkillGroupsFromUserRequest {
	s.UserId = &v
	return s
}

func (s *RemoveSkillGroupsFromUserRequest) Validate() error {
	return dara.Validate(s)
}
