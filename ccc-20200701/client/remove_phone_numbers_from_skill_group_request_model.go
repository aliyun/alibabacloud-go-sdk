// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemovePhoneNumbersFromSkillGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *RemovePhoneNumbersFromSkillGroupRequest
	GetInstanceId() *string
	SetNumberList(v string) *RemovePhoneNumbersFromSkillGroupRequest
	GetNumberList() *string
	SetSkillGroupId(v string) *RemovePhoneNumbersFromSkillGroupRequest
	GetSkillGroupId() *string
}

type RemovePhoneNumbersFromSkillGroupRequest struct {
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
	// ["0101234****","0105678****"]
	NumberList *string `json:"NumberList,omitempty" xml:"NumberList,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// skillgroup@ccc-test
	SkillGroupId *string `json:"SkillGroupId,omitempty" xml:"SkillGroupId,omitempty"`
}

func (s RemovePhoneNumbersFromSkillGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s RemovePhoneNumbersFromSkillGroupRequest) GoString() string {
	return s.String()
}

func (s *RemovePhoneNumbersFromSkillGroupRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *RemovePhoneNumbersFromSkillGroupRequest) GetNumberList() *string {
	return s.NumberList
}

func (s *RemovePhoneNumbersFromSkillGroupRequest) GetSkillGroupId() *string {
	return s.SkillGroupId
}

func (s *RemovePhoneNumbersFromSkillGroupRequest) SetInstanceId(v string) *RemovePhoneNumbersFromSkillGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *RemovePhoneNumbersFromSkillGroupRequest) SetNumberList(v string) *RemovePhoneNumbersFromSkillGroupRequest {
	s.NumberList = &v
	return s
}

func (s *RemovePhoneNumbersFromSkillGroupRequest) SetSkillGroupId(v string) *RemovePhoneNumbersFromSkillGroupRequest {
	s.SkillGroupId = &v
	return s
}

func (s *RemovePhoneNumbersFromSkillGroupRequest) Validate() error {
	return dara.Validate(s)
}
