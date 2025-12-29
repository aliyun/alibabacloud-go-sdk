// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitHotelOrderShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPayloadShrink(v string) *SubmitHotelOrderShrinkRequest
	GetPayloadShrink() *string
	SetUserInfoShrink(v string) *SubmitHotelOrderShrinkRequest
	GetUserInfoShrink() *string
}

type SubmitHotelOrderShrinkRequest struct {
	// This parameter is required.
	PayloadShrink *string `json:"Payload,omitempty" xml:"Payload,omitempty"`
	// This parameter is required.
	UserInfoShrink *string `json:"UserInfo,omitempty" xml:"UserInfo,omitempty"`
}

func (s SubmitHotelOrderShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitHotelOrderShrinkRequest) GoString() string {
	return s.String()
}

func (s *SubmitHotelOrderShrinkRequest) GetPayloadShrink() *string {
	return s.PayloadShrink
}

func (s *SubmitHotelOrderShrinkRequest) GetUserInfoShrink() *string {
	return s.UserInfoShrink
}

func (s *SubmitHotelOrderShrinkRequest) SetPayloadShrink(v string) *SubmitHotelOrderShrinkRequest {
	s.PayloadShrink = &v
	return s
}

func (s *SubmitHotelOrderShrinkRequest) SetUserInfoShrink(v string) *SubmitHotelOrderShrinkRequest {
	s.UserInfoShrink = &v
	return s
}

func (s *SubmitHotelOrderShrinkRequest) Validate() error {
	return dara.Validate(s)
}
