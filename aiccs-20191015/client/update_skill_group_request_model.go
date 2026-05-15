// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSkillGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *UpdateSkillGroupRequest
	GetClientToken() *string
	SetDescription(v string) *UpdateSkillGroupRequest
	GetDescription() *string
	SetDisplayName(v string) *UpdateSkillGroupRequest
	GetDisplayName() *string
	SetInstanceId(v string) *UpdateSkillGroupRequest
	GetInstanceId() *string
	SetSkillGroupId(v int64) *UpdateSkillGroupRequest
	GetSkillGroupId() *int64
	SetSkillGroupName(v string) *UpdateSkillGroupRequest
	GetSkillGroupName() *string
}

type UpdateSkillGroupRequest struct {
	// example:
	//
	// 46c1341e-2648-447a-9b11-70b6a298d94d
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ccc_xp_pre-cn-***
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 123456
	SkillGroupId   *int64  `json:"SkillGroupId,omitempty" xml:"SkillGroupId,omitempty"`
	SkillGroupName *string `json:"SkillGroupName,omitempty" xml:"SkillGroupName,omitempty"`
}

func (s UpdateSkillGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateSkillGroupRequest) GoString() string {
	return s.String()
}

func (s *UpdateSkillGroupRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateSkillGroupRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateSkillGroupRequest) GetDisplayName() *string {
	return s.DisplayName
}

func (s *UpdateSkillGroupRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateSkillGroupRequest) GetSkillGroupId() *int64 {
	return s.SkillGroupId
}

func (s *UpdateSkillGroupRequest) GetSkillGroupName() *string {
	return s.SkillGroupName
}

func (s *UpdateSkillGroupRequest) SetClientToken(v string) *UpdateSkillGroupRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateSkillGroupRequest) SetDescription(v string) *UpdateSkillGroupRequest {
	s.Description = &v
	return s
}

func (s *UpdateSkillGroupRequest) SetDisplayName(v string) *UpdateSkillGroupRequest {
	s.DisplayName = &v
	return s
}

func (s *UpdateSkillGroupRequest) SetInstanceId(v string) *UpdateSkillGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateSkillGroupRequest) SetSkillGroupId(v int64) *UpdateSkillGroupRequest {
	s.SkillGroupId = &v
	return s
}

func (s *UpdateSkillGroupRequest) SetSkillGroupName(v string) *UpdateSkillGroupRequest {
	s.SkillGroupName = &v
	return s
}

func (s *UpdateSkillGroupRequest) Validate() error {
	return dara.Validate(s)
}
