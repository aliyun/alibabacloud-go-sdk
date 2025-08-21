// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCateContentShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceInfoShrink(v string) *ListCateContentShrinkRequest
	GetDeviceInfoShrink() *string
	SetRequestShrink(v string) *ListCateContentShrinkRequest
	GetRequestShrink() *string
	SetUserInfoShrink(v string) *ListCateContentShrinkRequest
	GetUserInfoShrink() *string
}

type ListCateContentShrinkRequest struct {
	// This parameter is required.
	DeviceInfoShrink *string `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty"`
	// This parameter is required.
	RequestShrink *string `json:"Request,omitempty" xml:"Request,omitempty"`
	// This parameter is required.
	UserInfoShrink *string `json:"UserInfo,omitempty" xml:"UserInfo,omitempty"`
}

func (s ListCateContentShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCateContentShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListCateContentShrinkRequest) GetDeviceInfoShrink() *string {
	return s.DeviceInfoShrink
}

func (s *ListCateContentShrinkRequest) GetRequestShrink() *string {
	return s.RequestShrink
}

func (s *ListCateContentShrinkRequest) GetUserInfoShrink() *string {
	return s.UserInfoShrink
}

func (s *ListCateContentShrinkRequest) SetDeviceInfoShrink(v string) *ListCateContentShrinkRequest {
	s.DeviceInfoShrink = &v
	return s
}

func (s *ListCateContentShrinkRequest) SetRequestShrink(v string) *ListCateContentShrinkRequest {
	s.RequestShrink = &v
	return s
}

func (s *ListCateContentShrinkRequest) SetUserInfoShrink(v string) *ListCateContentShrinkRequest {
	s.UserInfoShrink = &v
	return s
}

func (s *ListCateContentShrinkRequest) Validate() error {
	return dara.Validate(s)
}
