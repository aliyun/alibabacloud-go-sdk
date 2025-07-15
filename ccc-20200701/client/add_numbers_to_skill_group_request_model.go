// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddNumbersToSkillGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstNumberGroupIdList(v string) *AddNumbersToSkillGroupRequest
	GetInstNumberGroupIdList() *string
	SetInstanceId(v string) *AddNumbersToSkillGroupRequest
	GetInstanceId() *string
	SetNumberList(v string) *AddNumbersToSkillGroupRequest
	GetNumberList() *string
	SetSkillGroupId(v string) *AddNumbersToSkillGroupRequest
	GetSkillGroupId() *string
}

type AddNumbersToSkillGroupRequest struct {
	InstNumberGroupIdList *string `json:"InstNumberGroupIdList,omitempty" xml:"InstNumberGroupIdList,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// ["0103182****","0102387****"]
	NumberList *string `json:"NumberList,omitempty" xml:"NumberList,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// skillgroup@ccc-test
	SkillGroupId *string `json:"SkillGroupId,omitempty" xml:"SkillGroupId,omitempty"`
}

func (s AddNumbersToSkillGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s AddNumbersToSkillGroupRequest) GoString() string {
	return s.String()
}

func (s *AddNumbersToSkillGroupRequest) GetInstNumberGroupIdList() *string {
	return s.InstNumberGroupIdList
}

func (s *AddNumbersToSkillGroupRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *AddNumbersToSkillGroupRequest) GetNumberList() *string {
	return s.NumberList
}

func (s *AddNumbersToSkillGroupRequest) GetSkillGroupId() *string {
	return s.SkillGroupId
}

func (s *AddNumbersToSkillGroupRequest) SetInstNumberGroupIdList(v string) *AddNumbersToSkillGroupRequest {
	s.InstNumberGroupIdList = &v
	return s
}

func (s *AddNumbersToSkillGroupRequest) SetInstanceId(v string) *AddNumbersToSkillGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *AddNumbersToSkillGroupRequest) SetNumberList(v string) *AddNumbersToSkillGroupRequest {
	s.NumberList = &v
	return s
}

func (s *AddNumbersToSkillGroupRequest) SetSkillGroupId(v string) *AddNumbersToSkillGroupRequest {
	s.SkillGroupId = &v
	return s
}

func (s *AddNumbersToSkillGroupRequest) Validate() error {
	return dara.Validate(s)
}
