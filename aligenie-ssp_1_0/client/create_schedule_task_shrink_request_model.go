// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateScheduleTaskShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceInfoShrink(v string) *CreateScheduleTaskShrinkRequest
	GetDeviceInfoShrink() *string
	SetPayloadShrink(v string) *CreateScheduleTaskShrinkRequest
	GetPayloadShrink() *string
	SetUserInfoShrink(v string) *CreateScheduleTaskShrinkRequest
	GetUserInfoShrink() *string
}

type CreateScheduleTaskShrinkRequest struct {
	// This parameter is required.
	DeviceInfoShrink *string `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty"`
	// This parameter is required.
	PayloadShrink *string `json:"Payload,omitempty" xml:"Payload,omitempty"`
	// This parameter is required.
	UserInfoShrink *string `json:"UserInfo,omitempty" xml:"UserInfo,omitempty"`
}

func (s CreateScheduleTaskShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateScheduleTaskShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateScheduleTaskShrinkRequest) GetDeviceInfoShrink() *string {
	return s.DeviceInfoShrink
}

func (s *CreateScheduleTaskShrinkRequest) GetPayloadShrink() *string {
	return s.PayloadShrink
}

func (s *CreateScheduleTaskShrinkRequest) GetUserInfoShrink() *string {
	return s.UserInfoShrink
}

func (s *CreateScheduleTaskShrinkRequest) SetDeviceInfoShrink(v string) *CreateScheduleTaskShrinkRequest {
	s.DeviceInfoShrink = &v
	return s
}

func (s *CreateScheduleTaskShrinkRequest) SetPayloadShrink(v string) *CreateScheduleTaskShrinkRequest {
	s.PayloadShrink = &v
	return s
}

func (s *CreateScheduleTaskShrinkRequest) SetUserInfoShrink(v string) *CreateScheduleTaskShrinkRequest {
	s.UserInfoShrink = &v
	return s
}

func (s *CreateScheduleTaskShrinkRequest) Validate() error {
	return dara.Validate(s)
}
