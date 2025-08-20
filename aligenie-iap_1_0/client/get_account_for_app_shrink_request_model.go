// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAccountForAppShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceInfoShrink(v string) *GetAccountForAppShrinkRequest
	GetDeviceInfoShrink() *string
	SetPayloadShrink(v string) *GetAccountForAppShrinkRequest
	GetPayloadShrink() *string
	SetUserInfoShrink(v string) *GetAccountForAppShrinkRequest
	GetUserInfoShrink() *string
}

type GetAccountForAppShrinkRequest struct {
	DeviceInfoShrink *string `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty"`
	PayloadShrink    *string `json:"Payload,omitempty" xml:"Payload,omitempty"`
	// This parameter is required.
	UserInfoShrink *string `json:"UserInfo,omitempty" xml:"UserInfo,omitempty"`
}

func (s GetAccountForAppShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAccountForAppShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetAccountForAppShrinkRequest) GetDeviceInfoShrink() *string {
	return s.DeviceInfoShrink
}

func (s *GetAccountForAppShrinkRequest) GetPayloadShrink() *string {
	return s.PayloadShrink
}

func (s *GetAccountForAppShrinkRequest) GetUserInfoShrink() *string {
	return s.UserInfoShrink
}

func (s *GetAccountForAppShrinkRequest) SetDeviceInfoShrink(v string) *GetAccountForAppShrinkRequest {
	s.DeviceInfoShrink = &v
	return s
}

func (s *GetAccountForAppShrinkRequest) SetPayloadShrink(v string) *GetAccountForAppShrinkRequest {
	s.PayloadShrink = &v
	return s
}

func (s *GetAccountForAppShrinkRequest) SetUserInfoShrink(v string) *GetAccountForAppShrinkRequest {
	s.UserInfoShrink = &v
	return s
}

func (s *GetAccountForAppShrinkRequest) Validate() error {
	return dara.Validate(s)
}
