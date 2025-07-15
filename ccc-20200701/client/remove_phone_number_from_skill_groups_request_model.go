// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemovePhoneNumberFromSkillGroupsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *RemovePhoneNumberFromSkillGroupsRequest
	GetInstanceId() *string
	SetNumber(v string) *RemovePhoneNumberFromSkillGroupsRequest
	GetNumber() *string
	SetSkillGroupIdList(v string) *RemovePhoneNumberFromSkillGroupsRequest
	GetSkillGroupIdList() *string
}

type RemovePhoneNumberFromSkillGroupsRequest struct {
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
	// 0101234****
	Number *string `json:"Number,omitempty" xml:"Number,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ["skillgroup1@ccc-test","skillgroup2@ccc-test"]
	SkillGroupIdList *string `json:"SkillGroupIdList,omitempty" xml:"SkillGroupIdList,omitempty"`
}

func (s RemovePhoneNumberFromSkillGroupsRequest) String() string {
	return dara.Prettify(s)
}

func (s RemovePhoneNumberFromSkillGroupsRequest) GoString() string {
	return s.String()
}

func (s *RemovePhoneNumberFromSkillGroupsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *RemovePhoneNumberFromSkillGroupsRequest) GetNumber() *string {
	return s.Number
}

func (s *RemovePhoneNumberFromSkillGroupsRequest) GetSkillGroupIdList() *string {
	return s.SkillGroupIdList
}

func (s *RemovePhoneNumberFromSkillGroupsRequest) SetInstanceId(v string) *RemovePhoneNumberFromSkillGroupsRequest {
	s.InstanceId = &v
	return s
}

func (s *RemovePhoneNumberFromSkillGroupsRequest) SetNumber(v string) *RemovePhoneNumberFromSkillGroupsRequest {
	s.Number = &v
	return s
}

func (s *RemovePhoneNumberFromSkillGroupsRequest) SetSkillGroupIdList(v string) *RemovePhoneNumberFromSkillGroupsRequest {
	s.SkillGroupIdList = &v
	return s
}

func (s *RemovePhoneNumberFromSkillGroupsRequest) Validate() error {
	return dara.Validate(s)
}
