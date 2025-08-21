// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetScheduleTaskShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceInfoShrink(v string) *GetScheduleTaskShrinkRequest
	GetDeviceInfoShrink() *string
	SetPayloadShrink(v string) *GetScheduleTaskShrinkRequest
	GetPayloadShrink() *string
	SetUserInfoShrink(v string) *GetScheduleTaskShrinkRequest
	GetUserInfoShrink() *string
}

type GetScheduleTaskShrinkRequest struct {
	// This parameter is required.
	DeviceInfoShrink *string `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty"`
	// This parameter is required.
	PayloadShrink *string `json:"Payload,omitempty" xml:"Payload,omitempty"`
	// This parameter is required.
	UserInfoShrink *string `json:"UserInfo,omitempty" xml:"UserInfo,omitempty"`
}

func (s GetScheduleTaskShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetScheduleTaskShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetScheduleTaskShrinkRequest) GetDeviceInfoShrink() *string {
	return s.DeviceInfoShrink
}

func (s *GetScheduleTaskShrinkRequest) GetPayloadShrink() *string {
	return s.PayloadShrink
}

func (s *GetScheduleTaskShrinkRequest) GetUserInfoShrink() *string {
	return s.UserInfoShrink
}

func (s *GetScheduleTaskShrinkRequest) SetDeviceInfoShrink(v string) *GetScheduleTaskShrinkRequest {
	s.DeviceInfoShrink = &v
	return s
}

func (s *GetScheduleTaskShrinkRequest) SetPayloadShrink(v string) *GetScheduleTaskShrinkRequest {
	s.PayloadShrink = &v
	return s
}

func (s *GetScheduleTaskShrinkRequest) SetUserInfoShrink(v string) *GetScheduleTaskShrinkRequest {
	s.UserInfoShrink = &v
	return s
}

func (s *GetScheduleTaskShrinkRequest) Validate() error {
	return dara.Validate(s)
}
