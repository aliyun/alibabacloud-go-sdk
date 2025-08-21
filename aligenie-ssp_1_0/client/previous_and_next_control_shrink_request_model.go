// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPreviousAndNextControlShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceInfoShrink(v string) *PreviousAndNextControlShrinkRequest
	GetDeviceInfoShrink() *string
	SetOpenControlPlayingListRequestShrink(v string) *PreviousAndNextControlShrinkRequest
	GetOpenControlPlayingListRequestShrink() *string
	SetUserInfoShrink(v string) *PreviousAndNextControlShrinkRequest
	GetUserInfoShrink() *string
}

type PreviousAndNextControlShrinkRequest struct {
	// This parameter is required.
	DeviceInfoShrink *string `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty"`
	// This parameter is required.
	OpenControlPlayingListRequestShrink *string `json:"OpenControlPlayingListRequest,omitempty" xml:"OpenControlPlayingListRequest,omitempty"`
	// This parameter is required.
	UserInfoShrink *string `json:"UserInfo,omitempty" xml:"UserInfo,omitempty"`
}

func (s PreviousAndNextControlShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s PreviousAndNextControlShrinkRequest) GoString() string {
	return s.String()
}

func (s *PreviousAndNextControlShrinkRequest) GetDeviceInfoShrink() *string {
	return s.DeviceInfoShrink
}

func (s *PreviousAndNextControlShrinkRequest) GetOpenControlPlayingListRequestShrink() *string {
	return s.OpenControlPlayingListRequestShrink
}

func (s *PreviousAndNextControlShrinkRequest) GetUserInfoShrink() *string {
	return s.UserInfoShrink
}

func (s *PreviousAndNextControlShrinkRequest) SetDeviceInfoShrink(v string) *PreviousAndNextControlShrinkRequest {
	s.DeviceInfoShrink = &v
	return s
}

func (s *PreviousAndNextControlShrinkRequest) SetOpenControlPlayingListRequestShrink(v string) *PreviousAndNextControlShrinkRequest {
	s.OpenControlPlayingListRequestShrink = &v
	return s
}

func (s *PreviousAndNextControlShrinkRequest) SetUserInfoShrink(v string) *PreviousAndNextControlShrinkRequest {
	s.UserInfoShrink = &v
	return s
}

func (s *PreviousAndNextControlShrinkRequest) Validate() error {
	return dara.Validate(s)
}
