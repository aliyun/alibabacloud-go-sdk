// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHotelContactsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetUserInfoShrink(v string) *GetHotelContactsShrinkRequest
	GetUserInfoShrink() *string
}

type GetHotelContactsShrinkRequest struct {
	UserInfoShrink *string `json:"UserInfo,omitempty" xml:"UserInfo,omitempty"`
}

func (s GetHotelContactsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetHotelContactsShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetHotelContactsShrinkRequest) GetUserInfoShrink() *string {
	return s.UserInfoShrink
}

func (s *GetHotelContactsShrinkRequest) SetUserInfoShrink(v string) *GetHotelContactsShrinkRequest {
	s.UserInfoShrink = &v
	return s
}

func (s *GetHotelContactsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
