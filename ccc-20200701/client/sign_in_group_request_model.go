// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSignInGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAdditivity(v bool) *SignInGroupRequest
	GetAdditivity() *bool
	SetChatDeviceId(v string) *SignInGroupRequest
	GetChatDeviceId() *string
	SetDeviceId(v string) *SignInGroupRequest
	GetDeviceId() *string
	SetInstanceId(v string) *SignInGroupRequest
	GetInstanceId() *string
	SetSignedSkillGroupIdList(v string) *SignInGroupRequest
	GetSignedSkillGroupIdList() *string
	SetUserId(v string) *SignInGroupRequest
	GetUserId() *string
}

type SignInGroupRequest struct {
	Additivity   *bool   `json:"Additivity,omitempty" xml:"Additivity,omitempty"`
	ChatDeviceId *string `json:"ChatDeviceId,omitempty" xml:"ChatDeviceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// device
	DeviceId *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
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
	SignedSkillGroupIdList *string `json:"SignedSkillGroupIdList,omitempty" xml:"SignedSkillGroupIdList,omitempty"`
	// example:
	//
	// agent@ccc-test
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s SignInGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s SignInGroupRequest) GoString() string {
	return s.String()
}

func (s *SignInGroupRequest) GetAdditivity() *bool {
	return s.Additivity
}

func (s *SignInGroupRequest) GetChatDeviceId() *string {
	return s.ChatDeviceId
}

func (s *SignInGroupRequest) GetDeviceId() *string {
	return s.DeviceId
}

func (s *SignInGroupRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *SignInGroupRequest) GetSignedSkillGroupIdList() *string {
	return s.SignedSkillGroupIdList
}

func (s *SignInGroupRequest) GetUserId() *string {
	return s.UserId
}

func (s *SignInGroupRequest) SetAdditivity(v bool) *SignInGroupRequest {
	s.Additivity = &v
	return s
}

func (s *SignInGroupRequest) SetChatDeviceId(v string) *SignInGroupRequest {
	s.ChatDeviceId = &v
	return s
}

func (s *SignInGroupRequest) SetDeviceId(v string) *SignInGroupRequest {
	s.DeviceId = &v
	return s
}

func (s *SignInGroupRequest) SetInstanceId(v string) *SignInGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *SignInGroupRequest) SetSignedSkillGroupIdList(v string) *SignInGroupRequest {
	s.SignedSkillGroupIdList = &v
	return s
}

func (s *SignInGroupRequest) SetUserId(v string) *SignInGroupRequest {
	s.UserId = &v
	return s
}

func (s *SignInGroupRequest) Validate() error {
	return dara.Validate(s)
}
