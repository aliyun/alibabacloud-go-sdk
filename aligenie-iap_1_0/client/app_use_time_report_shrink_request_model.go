// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAppUseTimeReportShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceInfoShrink(v string) *AppUseTimeReportShrinkRequest
	GetDeviceInfoShrink() *string
	SetPayloadShrink(v string) *AppUseTimeReportShrinkRequest
	GetPayloadShrink() *string
	SetUserInfoShrink(v string) *AppUseTimeReportShrinkRequest
	GetUserInfoShrink() *string
}

type AppUseTimeReportShrinkRequest struct {
	// Device identity information
	DeviceInfoShrink *string `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty"`
	// Business parameters
	PayloadShrink *string `json:"Payload,omitempty" xml:"Payload,omitempty"`
	// User identifier information
	//
	// This parameter is required.
	UserInfoShrink *string `json:"UserInfo,omitempty" xml:"UserInfo,omitempty"`
}

func (s AppUseTimeReportShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s AppUseTimeReportShrinkRequest) GoString() string {
	return s.String()
}

func (s *AppUseTimeReportShrinkRequest) GetDeviceInfoShrink() *string {
	return s.DeviceInfoShrink
}

func (s *AppUseTimeReportShrinkRequest) GetPayloadShrink() *string {
	return s.PayloadShrink
}

func (s *AppUseTimeReportShrinkRequest) GetUserInfoShrink() *string {
	return s.UserInfoShrink
}

func (s *AppUseTimeReportShrinkRequest) SetDeviceInfoShrink(v string) *AppUseTimeReportShrinkRequest {
	s.DeviceInfoShrink = &v
	return s
}

func (s *AppUseTimeReportShrinkRequest) SetPayloadShrink(v string) *AppUseTimeReportShrinkRequest {
	s.PayloadShrink = &v
	return s
}

func (s *AppUseTimeReportShrinkRequest) SetUserInfoShrink(v string) *AppUseTimeReportShrinkRequest {
	s.UserInfoShrink = &v
	return s
}

func (s *AppUseTimeReportShrinkRequest) Validate() error {
	return dara.Validate(s)
}
