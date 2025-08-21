// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSubShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceInfoShrink(v string) *ListSubShrinkRequest
	GetDeviceInfoShrink() *string
	SetPageShrink(v string) *ListSubShrinkRequest
	GetPageShrink() *string
	SetUserInfoShrink(v string) *ListSubShrinkRequest
	GetUserInfoShrink() *string
}

type ListSubShrinkRequest struct {
	// This parameter is required.
	DeviceInfoShrink *string `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty"`
	// This parameter is required.
	PageShrink *string `json:"Page,omitempty" xml:"Page,omitempty"`
	// This parameter is required.
	UserInfoShrink *string `json:"UserInfo,omitempty" xml:"UserInfo,omitempty"`
}

func (s ListSubShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSubShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListSubShrinkRequest) GetDeviceInfoShrink() *string {
	return s.DeviceInfoShrink
}

func (s *ListSubShrinkRequest) GetPageShrink() *string {
	return s.PageShrink
}

func (s *ListSubShrinkRequest) GetUserInfoShrink() *string {
	return s.UserInfoShrink
}

func (s *ListSubShrinkRequest) SetDeviceInfoShrink(v string) *ListSubShrinkRequest {
	s.DeviceInfoShrink = &v
	return s
}

func (s *ListSubShrinkRequest) SetPageShrink(v string) *ListSubShrinkRequest {
	s.PageShrink = &v
	return s
}

func (s *ListSubShrinkRequest) SetUserInfoShrink(v string) *ListSubShrinkRequest {
	s.UserInfoShrink = &v
	return s
}

func (s *ListSubShrinkRequest) Validate() error {
	return dara.Validate(s)
}
