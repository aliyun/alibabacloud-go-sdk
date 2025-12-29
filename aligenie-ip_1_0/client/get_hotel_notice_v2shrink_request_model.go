// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHotelNoticeV2ShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetUserInfoShrink(v string) *GetHotelNoticeV2ShrinkRequest
	GetUserInfoShrink() *string
}

type GetHotelNoticeV2ShrinkRequest struct {
	// This parameter is required.
	UserInfoShrink *string `json:"UserInfo,omitempty" xml:"UserInfo,omitempty"`
}

func (s GetHotelNoticeV2ShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetHotelNoticeV2ShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetHotelNoticeV2ShrinkRequest) GetUserInfoShrink() *string {
	return s.UserInfoShrink
}

func (s *GetHotelNoticeV2ShrinkRequest) SetUserInfoShrink(v string) *GetHotelNoticeV2ShrinkRequest {
	s.UserInfoShrink = &v
	return s
}

func (s *GetHotelNoticeV2ShrinkRequest) Validate() error {
	return dara.Validate(s)
}
