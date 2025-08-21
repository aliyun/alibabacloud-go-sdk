// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteScheduleTaskShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceInfoShrink(v string) *DeleteScheduleTaskShrinkRequest
	GetDeviceInfoShrink() *string
	SetPayloadShrink(v string) *DeleteScheduleTaskShrinkRequest
	GetPayloadShrink() *string
	SetUserInfoShrink(v string) *DeleteScheduleTaskShrinkRequest
	GetUserInfoShrink() *string
}

type DeleteScheduleTaskShrinkRequest struct {
	// This parameter is required.
	DeviceInfoShrink *string `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty"`
	// This parameter is required.
	PayloadShrink *string `json:"Payload,omitempty" xml:"Payload,omitempty"`
	// This parameter is required.
	UserInfoShrink *string `json:"UserInfo,omitempty" xml:"UserInfo,omitempty"`
}

func (s DeleteScheduleTaskShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteScheduleTaskShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteScheduleTaskShrinkRequest) GetDeviceInfoShrink() *string {
	return s.DeviceInfoShrink
}

func (s *DeleteScheduleTaskShrinkRequest) GetPayloadShrink() *string {
	return s.PayloadShrink
}

func (s *DeleteScheduleTaskShrinkRequest) GetUserInfoShrink() *string {
	return s.UserInfoShrink
}

func (s *DeleteScheduleTaskShrinkRequest) SetDeviceInfoShrink(v string) *DeleteScheduleTaskShrinkRequest {
	s.DeviceInfoShrink = &v
	return s
}

func (s *DeleteScheduleTaskShrinkRequest) SetPayloadShrink(v string) *DeleteScheduleTaskShrinkRequest {
	s.PayloadShrink = &v
	return s
}

func (s *DeleteScheduleTaskShrinkRequest) SetUserInfoShrink(v string) *DeleteScheduleTaskShrinkRequest {
	s.UserInfoShrink = &v
	return s
}

func (s *DeleteScheduleTaskShrinkRequest) Validate() error {
	return dara.Validate(s)
}
