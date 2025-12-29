// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHotelScreenSaverShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetUserInfoShrink(v string) *GetHotelScreenSaverShrinkRequest
	GetUserInfoShrink() *string
}

type GetHotelScreenSaverShrinkRequest struct {
	// This parameter is required.
	UserInfoShrink *string `json:"UserInfo,omitempty" xml:"UserInfo,omitempty"`
}

func (s GetHotelScreenSaverShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetHotelScreenSaverShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetHotelScreenSaverShrinkRequest) GetUserInfoShrink() *string {
	return s.UserInfoShrink
}

func (s *GetHotelScreenSaverShrinkRequest) SetUserInfoShrink(v string) *GetHotelScreenSaverShrinkRequest {
	s.UserInfoShrink = &v
	return s
}

func (s *GetHotelScreenSaverShrinkRequest) Validate() error {
	return dara.Validate(s)
}
