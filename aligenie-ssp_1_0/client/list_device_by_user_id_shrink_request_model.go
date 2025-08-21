// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDeviceByUserIdShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetUserInfoShrink(v string) *ListDeviceByUserIdShrinkRequest
	GetUserInfoShrink() *string
}

type ListDeviceByUserIdShrinkRequest struct {
	// This parameter is required.
	UserInfoShrink *string `json:"UserInfo,omitempty" xml:"UserInfo,omitempty"`
}

func (s ListDeviceByUserIdShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDeviceByUserIdShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListDeviceByUserIdShrinkRequest) GetUserInfoShrink() *string {
	return s.UserInfoShrink
}

func (s *ListDeviceByUserIdShrinkRequest) SetUserInfoShrink(v string) *ListDeviceByUserIdShrinkRequest {
	s.UserInfoShrink = &v
	return s
}

func (s *ListDeviceByUserIdShrinkRequest) Validate() error {
	return dara.Validate(s)
}
