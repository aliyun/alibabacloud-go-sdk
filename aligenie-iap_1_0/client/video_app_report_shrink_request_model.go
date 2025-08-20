// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVideoAppReportShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceInfoShrink(v string) *VideoAppReportShrinkRequest
	GetDeviceInfoShrink() *string
	SetPayloadShrink(v string) *VideoAppReportShrinkRequest
	GetPayloadShrink() *string
	SetUserInfoShrink(v string) *VideoAppReportShrinkRequest
	GetUserInfoShrink() *string
}

type VideoAppReportShrinkRequest struct {
	DeviceInfoShrink *string `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty"`
	PayloadShrink    *string `json:"Payload,omitempty" xml:"Payload,omitempty"`
	// This parameter is required.
	UserInfoShrink *string `json:"UserInfo,omitempty" xml:"UserInfo,omitempty"`
}

func (s VideoAppReportShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s VideoAppReportShrinkRequest) GoString() string {
	return s.String()
}

func (s *VideoAppReportShrinkRequest) GetDeviceInfoShrink() *string {
	return s.DeviceInfoShrink
}

func (s *VideoAppReportShrinkRequest) GetPayloadShrink() *string {
	return s.PayloadShrink
}

func (s *VideoAppReportShrinkRequest) GetUserInfoShrink() *string {
	return s.UserInfoShrink
}

func (s *VideoAppReportShrinkRequest) SetDeviceInfoShrink(v string) *VideoAppReportShrinkRequest {
	s.DeviceInfoShrink = &v
	return s
}

func (s *VideoAppReportShrinkRequest) SetPayloadShrink(v string) *VideoAppReportShrinkRequest {
	s.PayloadShrink = &v
	return s
}

func (s *VideoAppReportShrinkRequest) SetUserInfoShrink(v string) *VideoAppReportShrinkRequest {
	s.UserInfoShrink = &v
	return s
}

func (s *VideoAppReportShrinkRequest) Validate() error {
	return dara.Validate(s)
}
