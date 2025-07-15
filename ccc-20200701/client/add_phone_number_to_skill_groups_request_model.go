// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddPhoneNumberToSkillGroupsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *AddPhoneNumberToSkillGroupsRequest
	GetInstanceId() *string
	SetNumber(v string) *AddPhoneNumberToSkillGroupsRequest
	GetNumber() *string
	SetSkillGroupIdList(v string) *AddPhoneNumberToSkillGroupsRequest
	GetSkillGroupIdList() *string
}

type AddPhoneNumberToSkillGroupsRequest struct {
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

func (s AddPhoneNumberToSkillGroupsRequest) String() string {
	return dara.Prettify(s)
}

func (s AddPhoneNumberToSkillGroupsRequest) GoString() string {
	return s.String()
}

func (s *AddPhoneNumberToSkillGroupsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *AddPhoneNumberToSkillGroupsRequest) GetNumber() *string {
	return s.Number
}

func (s *AddPhoneNumberToSkillGroupsRequest) GetSkillGroupIdList() *string {
	return s.SkillGroupIdList
}

func (s *AddPhoneNumberToSkillGroupsRequest) SetInstanceId(v string) *AddPhoneNumberToSkillGroupsRequest {
	s.InstanceId = &v
	return s
}

func (s *AddPhoneNumberToSkillGroupsRequest) SetNumber(v string) *AddPhoneNumberToSkillGroupsRequest {
	s.Number = &v
	return s
}

func (s *AddPhoneNumberToSkillGroupsRequest) SetSkillGroupIdList(v string) *AddPhoneNumberToSkillGroupsRequest {
	s.SkillGroupIdList = &v
	return s
}

func (s *AddPhoneNumberToSkillGroupsRequest) Validate() error {
	return dara.Validate(s)
}
