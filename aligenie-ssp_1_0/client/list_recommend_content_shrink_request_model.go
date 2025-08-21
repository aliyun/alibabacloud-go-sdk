// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRecommendContentShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceInfoShrink(v string) *ListRecommendContentShrinkRequest
	GetDeviceInfoShrink() *string
	SetRequestShrink(v string) *ListRecommendContentShrinkRequest
	GetRequestShrink() *string
	SetUserInfoShrink(v string) *ListRecommendContentShrinkRequest
	GetUserInfoShrink() *string
}

type ListRecommendContentShrinkRequest struct {
	// This parameter is required.
	DeviceInfoShrink *string `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty"`
	// This parameter is required.
	RequestShrink *string `json:"Request,omitempty" xml:"Request,omitempty"`
	// This parameter is required.
	UserInfoShrink *string `json:"UserInfo,omitempty" xml:"UserInfo,omitempty"`
}

func (s ListRecommendContentShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListRecommendContentShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListRecommendContentShrinkRequest) GetDeviceInfoShrink() *string {
	return s.DeviceInfoShrink
}

func (s *ListRecommendContentShrinkRequest) GetRequestShrink() *string {
	return s.RequestShrink
}

func (s *ListRecommendContentShrinkRequest) GetUserInfoShrink() *string {
	return s.UserInfoShrink
}

func (s *ListRecommendContentShrinkRequest) SetDeviceInfoShrink(v string) *ListRecommendContentShrinkRequest {
	s.DeviceInfoShrink = &v
	return s
}

func (s *ListRecommendContentShrinkRequest) SetRequestShrink(v string) *ListRecommendContentShrinkRequest {
	s.RequestShrink = &v
	return s
}

func (s *ListRecommendContentShrinkRequest) SetUserInfoShrink(v string) *ListRecommendContentShrinkRequest {
	s.UserInfoShrink = &v
	return s
}

func (s *ListRecommendContentShrinkRequest) Validate() error {
	return dara.Validate(s)
}
