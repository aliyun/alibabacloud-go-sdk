// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHotelHomeBackImageAndModesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetUserInfoShrink(v string) *GetHotelHomeBackImageAndModesShrinkRequest
	GetUserInfoShrink() *string
}

type GetHotelHomeBackImageAndModesShrinkRequest struct {
	// This parameter is required.
	UserInfoShrink *string `json:"UserInfo,omitempty" xml:"UserInfo,omitempty"`
}

func (s GetHotelHomeBackImageAndModesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetHotelHomeBackImageAndModesShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetHotelHomeBackImageAndModesShrinkRequest) GetUserInfoShrink() *string {
	return s.UserInfoShrink
}

func (s *GetHotelHomeBackImageAndModesShrinkRequest) SetUserInfoShrink(v string) *GetHotelHomeBackImageAndModesShrinkRequest {
	s.UserInfoShrink = &v
	return s
}

func (s *GetHotelHomeBackImageAndModesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
