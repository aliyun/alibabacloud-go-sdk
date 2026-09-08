// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSkillGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateSkillGroupRequest
	GetDescription() *string
	SetDisplayName(v string) *CreateSkillGroupRequest
	GetDisplayName() *string
	SetInstanceId(v string) *CreateSkillGroupRequest
	GetInstanceId() *string
	SetMediaType(v string) *CreateSkillGroupRequest
	GetMediaType() *string
	SetName(v string) *CreateSkillGroupRequest
	GetName() *string
}

type CreateSkillGroupRequest struct {
	// Description of the skill group. This parameter is optional. If not specified, it defaults to empty.
	//
	// example:
	//
	// 云联络中心的测试技能组。
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Display name of the skill group. It must be 1 to 64 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// 测试技能组
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// Instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Media type. The default value is AUDIO. Other valid values include CHAT and VIDEO.
	//
	// example:
	//
	// CHAT
	MediaType *string `json:"MediaType,omitempty" xml:"MediaType,omitempty"`
	// The skill group name. It must be 1 to 64 characters in length, can contain digits, letters, underscores (_), and hyphens (-), and must start with a letter.
	//
	// This parameter is required.
	//
	// example:
	//
	// skillgroup
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s CreateSkillGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSkillGroupRequest) GoString() string {
	return s.String()
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

func (s *CreateSkillGroupRequest) GetMediaType() *string {
	return s.MediaType
}

func (s *CreateSkillGroupRequest) GetName() *string {
	return s.Name
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

func (s *CreateSkillGroupRequest) SetMediaType(v string) *CreateSkillGroupRequest {
	s.MediaType = &v
	return s
}

func (s *CreateSkillGroupRequest) SetName(v string) *CreateSkillGroupRequest {
	s.Name = &v
	return s
}

func (s *CreateSkillGroupRequest) Validate() error {
	return dara.Validate(s)
}
