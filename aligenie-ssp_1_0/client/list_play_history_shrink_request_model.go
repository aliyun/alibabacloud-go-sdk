// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPlayHistoryShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceInfoShrink(v string) *ListPlayHistoryShrinkRequest
	GetDeviceInfoShrink() *string
	SetRequestShrink(v string) *ListPlayHistoryShrinkRequest
	GetRequestShrink() *string
	SetUserInfoShrink(v string) *ListPlayHistoryShrinkRequest
	GetUserInfoShrink() *string
}

type ListPlayHistoryShrinkRequest struct {
	// This parameter is required.
	DeviceInfoShrink *string `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty"`
	// This parameter is required.
	RequestShrink *string `json:"Request,omitempty" xml:"Request,omitempty"`
	// This parameter is required.
	UserInfoShrink *string `json:"UserInfo,omitempty" xml:"UserInfo,omitempty"`
}

func (s ListPlayHistoryShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPlayHistoryShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListPlayHistoryShrinkRequest) GetDeviceInfoShrink() *string {
	return s.DeviceInfoShrink
}

func (s *ListPlayHistoryShrinkRequest) GetRequestShrink() *string {
	return s.RequestShrink
}

func (s *ListPlayHistoryShrinkRequest) GetUserInfoShrink() *string {
	return s.UserInfoShrink
}

func (s *ListPlayHistoryShrinkRequest) SetDeviceInfoShrink(v string) *ListPlayHistoryShrinkRequest {
	s.DeviceInfoShrink = &v
	return s
}

func (s *ListPlayHistoryShrinkRequest) SetRequestShrink(v string) *ListPlayHistoryShrinkRequest {
	s.RequestShrink = &v
	return s
}

func (s *ListPlayHistoryShrinkRequest) SetUserInfoShrink(v string) *ListPlayHistoryShrinkRequest {
	s.UserInfoShrink = &v
	return s
}

func (s *ListPlayHistoryShrinkRequest) Validate() error {
	return dara.Validate(s)
}
