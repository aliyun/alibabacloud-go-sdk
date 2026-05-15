// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSkillGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChannelType(v int32) *CreateSkillGroupRequest
	GetChannelType() *int32
	SetClientToken(v string) *CreateSkillGroupRequest
	GetClientToken() *string
	SetDepartmentId(v int64) *CreateSkillGroupRequest
	GetDepartmentId() *int64
	SetDescription(v string) *CreateSkillGroupRequest
	GetDescription() *string
	SetDisplayName(v string) *CreateSkillGroupRequest
	GetDisplayName() *string
	SetInstanceId(v string) *CreateSkillGroupRequest
	GetInstanceId() *string
	SetSkillGroupName(v string) *CreateSkillGroupRequest
	GetSkillGroupName() *string
}

type CreateSkillGroupRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 2
	ChannelType *int32 `json:"ChannelType,omitempty" xml:"ChannelType,omitempty"`
	// example:
	//
	// 46c1341e-2648-447a-****-70b6a298d94d
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// example:
	//
	// 123
	DepartmentId *int64  `json:"DepartmentId,omitempty" xml:"DepartmentId,omitempty"`
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DisplayName  *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ccc_xp_pre-cn-***
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	SkillGroupName *string `json:"SkillGroupName,omitempty" xml:"SkillGroupName,omitempty"`
}

func (s CreateSkillGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSkillGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateSkillGroupRequest) GetChannelType() *int32 {
	return s.ChannelType
}

func (s *CreateSkillGroupRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateSkillGroupRequest) GetDepartmentId() *int64 {
	return s.DepartmentId
}

func (s *CreateSkillGroupRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateSkillGroupRequest) GetDisplayName() *string {
	return s.DisplayName
}

func (s *CreateSkillGroupRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateSkillGroupRequest) GetSkillGroupName() *string {
	return s.SkillGroupName
}

func (s *CreateSkillGroupRequest) SetChannelType(v int32) *CreateSkillGroupRequest {
	s.ChannelType = &v
	return s
}

func (s *CreateSkillGroupRequest) SetClientToken(v string) *CreateSkillGroupRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateSkillGroupRequest) SetDepartmentId(v int64) *CreateSkillGroupRequest {
	s.DepartmentId = &v
	return s
}

func (s *CreateSkillGroupRequest) SetDescription(v string) *CreateSkillGroupRequest {
	s.Description = &v
	return s
}

func (s *CreateSkillGroupRequest) SetDisplayName(v string) *CreateSkillGroupRequest {
	s.DisplayName = &v
	return s
}

func (s *CreateSkillGroupRequest) SetInstanceId(v string) *CreateSkillGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateSkillGroupRequest) SetSkillGroupName(v string) *CreateSkillGroupRequest {
	s.SkillGroupName = &v
	return s
}

func (s *CreateSkillGroupRequest) Validate() error {
	return dara.Validate(s)
}
