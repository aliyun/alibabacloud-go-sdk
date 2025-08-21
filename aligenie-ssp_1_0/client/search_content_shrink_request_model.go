// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchContentShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceInfoShrink(v string) *SearchContentShrinkRequest
	GetDeviceInfoShrink() *string
	SetRequestShrink(v string) *SearchContentShrinkRequest
	GetRequestShrink() *string
	SetUserInfoShrink(v string) *SearchContentShrinkRequest
	GetUserInfoShrink() *string
}

type SearchContentShrinkRequest struct {
	// This parameter is required.
	DeviceInfoShrink *string `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty"`
	// This parameter is required.
	RequestShrink *string `json:"Request,omitempty" xml:"Request,omitempty"`
	// This parameter is required.
	UserInfoShrink *string `json:"UserInfo,omitempty" xml:"UserInfo,omitempty"`
}

func (s SearchContentShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s SearchContentShrinkRequest) GoString() string {
	return s.String()
}

func (s *SearchContentShrinkRequest) GetDeviceInfoShrink() *string {
	return s.DeviceInfoShrink
}

func (s *SearchContentShrinkRequest) GetRequestShrink() *string {
	return s.RequestShrink
}

func (s *SearchContentShrinkRequest) GetUserInfoShrink() *string {
	return s.UserInfoShrink
}

func (s *SearchContentShrinkRequest) SetDeviceInfoShrink(v string) *SearchContentShrinkRequest {
	s.DeviceInfoShrink = &v
	return s
}

func (s *SearchContentShrinkRequest) SetRequestShrink(v string) *SearchContentShrinkRequest {
	s.RequestShrink = &v
	return s
}

func (s *SearchContentShrinkRequest) SetUserInfoShrink(v string) *SearchContentShrinkRequest {
	s.UserInfoShrink = &v
	return s
}

func (s *SearchContentShrinkRequest) Validate() error {
	return dara.Validate(s)
}
