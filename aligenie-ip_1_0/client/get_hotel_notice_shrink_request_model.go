// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHotelNoticeShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetUserInfoShrink(v string) *GetHotelNoticeShrinkRequest
	GetUserInfoShrink() *string
}

type GetHotelNoticeShrinkRequest struct {
	// This parameter is required.
	UserInfoShrink *string `json:"UserInfo,omitempty" xml:"UserInfo,omitempty"`
}

func (s GetHotelNoticeShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetHotelNoticeShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetHotelNoticeShrinkRequest) GetUserInfoShrink() *string {
	return s.UserInfoShrink
}

func (s *GetHotelNoticeShrinkRequest) SetUserInfoShrink(v string) *GetHotelNoticeShrinkRequest {
	s.UserInfoShrink = &v
	return s
}

func (s *GetHotelNoticeShrinkRequest) Validate() error {
	return dara.Validate(s)
}
