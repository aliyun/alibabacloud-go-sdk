// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHotelSampleUtterancesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetUserInfoShrink(v string) *GetHotelSampleUtterancesShrinkRequest
	GetUserInfoShrink() *string
}

type GetHotelSampleUtterancesShrinkRequest struct {
	UserInfoShrink *string `json:"UserInfo,omitempty" xml:"UserInfo,omitempty"`
}

func (s GetHotelSampleUtterancesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetHotelSampleUtterancesShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetHotelSampleUtterancesShrinkRequest) GetUserInfoShrink() *string {
	return s.UserInfoShrink
}

func (s *GetHotelSampleUtterancesShrinkRequest) SetUserInfoShrink(v string) *GetHotelSampleUtterancesShrinkRequest {
	s.UserInfoShrink = &v
	return s
}

func (s *GetHotelSampleUtterancesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
