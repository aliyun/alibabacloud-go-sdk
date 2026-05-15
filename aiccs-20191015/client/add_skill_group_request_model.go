// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddSkillGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOuterDepartmentId(v string) *AddSkillGroupRequest
	GetOuterDepartmentId() *string
	SetOuterDepartmentType(v string) *AddSkillGroupRequest
	GetOuterDepartmentType() *string
	SetOuterGroupId(v string) *AddSkillGroupRequest
	GetOuterGroupId() *string
	SetOuterGroupName(v string) *AddSkillGroupRequest
	GetOuterGroupName() *string
	SetOuterGroupType(v string) *AddSkillGroupRequest
	GetOuterGroupType() *string
}

type AddSkillGroupRequest struct {
	// example:
	//
	// 123456
	OuterDepartmentId *string `json:"OuterDepartmentId,omitempty" xml:"OuterDepartmentId,omitempty"`
	// example:
	//
	// type_invalid
	OuterDepartmentType *string `json:"OuterDepartmentType,omitempty" xml:"OuterDepartmentType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 123456
	OuterGroupId *string `json:"OuterGroupId,omitempty" xml:"OuterGroupId,omitempty"`
	// This parameter is required.
	OuterGroupName *string `json:"OuterGroupName,omitempty" xml:"OuterGroupName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// mybank
	OuterGroupType *string `json:"OuterGroupType,omitempty" xml:"OuterGroupType,omitempty"`
}

func (s AddSkillGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s AddSkillGroupRequest) GoString() string {
	return s.String()
}

func (s *AddSkillGroupRequest) GetOuterDepartmentId() *string {
	return s.OuterDepartmentId
}

func (s *AddSkillGroupRequest) GetOuterDepartmentType() *string {
	return s.OuterDepartmentType
}

func (s *AddSkillGroupRequest) GetOuterGroupId() *string {
	return s.OuterGroupId
}

func (s *AddSkillGroupRequest) GetOuterGroupName() *string {
	return s.OuterGroupName
}

func (s *AddSkillGroupRequest) GetOuterGroupType() *string {
	return s.OuterGroupType
}

func (s *AddSkillGroupRequest) SetOuterDepartmentId(v string) *AddSkillGroupRequest {
	s.OuterDepartmentId = &v
	return s
}

func (s *AddSkillGroupRequest) SetOuterDepartmentType(v string) *AddSkillGroupRequest {
	s.OuterDepartmentType = &v
	return s
}

func (s *AddSkillGroupRequest) SetOuterGroupId(v string) *AddSkillGroupRequest {
	s.OuterGroupId = &v
	return s
}

func (s *AddSkillGroupRequest) SetOuterGroupName(v string) *AddSkillGroupRequest {
	s.OuterGroupName = &v
	return s
}

func (s *AddSkillGroupRequest) SetOuterGroupType(v string) *AddSkillGroupRequest {
	s.OuterGroupType = &v
	return s
}

func (s *AddSkillGroupRequest) Validate() error {
	return dara.Validate(s)
}
