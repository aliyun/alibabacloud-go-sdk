// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMusicTypeShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceInfoShrink(v string) *QueryMusicTypeShrinkRequest
	GetDeviceInfoShrink() *string
	SetPayloadShrink(v string) *QueryMusicTypeShrinkRequest
	GetPayloadShrink() *string
	SetUserInfoShrink(v string) *QueryMusicTypeShrinkRequest
	GetUserInfoShrink() *string
}

type QueryMusicTypeShrinkRequest struct {
	// Device identity information
	//
	// This parameter is required.
	DeviceInfoShrink *string `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty"`
	// Input parameters for the service request
	PayloadShrink *string `json:"Payload,omitempty" xml:"Payload,omitempty"`
	// User identifier information
	//
	// This parameter is required.
	UserInfoShrink *string `json:"UserInfo,omitempty" xml:"UserInfo,omitempty"`
}

func (s QueryMusicTypeShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryMusicTypeShrinkRequest) GoString() string {
	return s.String()
}

func (s *QueryMusicTypeShrinkRequest) GetDeviceInfoShrink() *string {
	return s.DeviceInfoShrink
}

func (s *QueryMusicTypeShrinkRequest) GetPayloadShrink() *string {
	return s.PayloadShrink
}

func (s *QueryMusicTypeShrinkRequest) GetUserInfoShrink() *string {
	return s.UserInfoShrink
}

func (s *QueryMusicTypeShrinkRequest) SetDeviceInfoShrink(v string) *QueryMusicTypeShrinkRequest {
	s.DeviceInfoShrink = &v
	return s
}

func (s *QueryMusicTypeShrinkRequest) SetPayloadShrink(v string) *QueryMusicTypeShrinkRequest {
	s.PayloadShrink = &v
	return s
}

func (s *QueryMusicTypeShrinkRequest) SetUserInfoShrink(v string) *QueryMusicTypeShrinkRequest {
	s.UserInfoShrink = &v
	return s
}

func (s *QueryMusicTypeShrinkRequest) Validate() error {
	return dara.Validate(s)
}
