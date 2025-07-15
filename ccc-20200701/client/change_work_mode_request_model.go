// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangeWorkModeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceId(v string) *ChangeWorkModeRequest
	GetDeviceId() *string
	SetInstanceId(v string) *ChangeWorkModeRequest
	GetInstanceId() *string
	SetMobile(v string) *ChangeWorkModeRequest
	GetMobile() *string
	SetSignedSkillGroupIdList(v string) *ChangeWorkModeRequest
	GetSignedSkillGroupIdList() *string
	SetUserId(v string) *ChangeWorkModeRequest
	GetUserId() *string
	SetWorkMode(v string) *ChangeWorkModeRequest
	GetWorkMode() *string
}

type ChangeWorkModeRequest struct {
	// example:
	//
	// ACC-YUNBS-1.0.10-****
	DeviceId *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 1382114****
	Mobile *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	// example:
	//
	// ["online-1@ccc-test","online-2@ccc-test","online-3@ccc-test","skg-default@ccc-test"]
	SignedSkillGroupIdList *string `json:"SignedSkillGroupIdList,omitempty" xml:"SignedSkillGroupIdList,omitempty"`
	// example:
	//
	// agent@ccc-test
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ON_SITE
	WorkMode *string `json:"WorkMode,omitempty" xml:"WorkMode,omitempty"`
}

func (s ChangeWorkModeRequest) String() string {
	return dara.Prettify(s)
}

func (s ChangeWorkModeRequest) GoString() string {
	return s.String()
}

func (s *ChangeWorkModeRequest) GetDeviceId() *string {
	return s.DeviceId
}

func (s *ChangeWorkModeRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ChangeWorkModeRequest) GetMobile() *string {
	return s.Mobile
}

func (s *ChangeWorkModeRequest) GetSignedSkillGroupIdList() *string {
	return s.SignedSkillGroupIdList
}

func (s *ChangeWorkModeRequest) GetUserId() *string {
	return s.UserId
}

func (s *ChangeWorkModeRequest) GetWorkMode() *string {
	return s.WorkMode
}

func (s *ChangeWorkModeRequest) SetDeviceId(v string) *ChangeWorkModeRequest {
	s.DeviceId = &v
	return s
}

func (s *ChangeWorkModeRequest) SetInstanceId(v string) *ChangeWorkModeRequest {
	s.InstanceId = &v
	return s
}

func (s *ChangeWorkModeRequest) SetMobile(v string) *ChangeWorkModeRequest {
	s.Mobile = &v
	return s
}

func (s *ChangeWorkModeRequest) SetSignedSkillGroupIdList(v string) *ChangeWorkModeRequest {
	s.SignedSkillGroupIdList = &v
	return s
}

func (s *ChangeWorkModeRequest) SetUserId(v string) *ChangeWorkModeRequest {
	s.UserId = &v
	return s
}

func (s *ChangeWorkModeRequest) SetWorkMode(v string) *ChangeWorkModeRequest {
	s.WorkMode = &v
	return s
}

func (s *ChangeWorkModeRequest) Validate() error {
	return dara.Validate(s)
}
