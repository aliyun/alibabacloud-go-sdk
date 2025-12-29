// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHotelOrderShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPayloadShrink(v string) *ListHotelOrderShrinkRequest
	GetPayloadShrink() *string
	SetUserInfoShrink(v string) *ListHotelOrderShrinkRequest
	GetUserInfoShrink() *string
}

type ListHotelOrderShrinkRequest struct {
	// This parameter is required.
	PayloadShrink *string `json:"Payload,omitempty" xml:"Payload,omitempty"`
	// This parameter is required.
	UserInfoShrink *string `json:"UserInfo,omitempty" xml:"UserInfo,omitempty"`
}

func (s ListHotelOrderShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListHotelOrderShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListHotelOrderShrinkRequest) GetPayloadShrink() *string {
	return s.PayloadShrink
}

func (s *ListHotelOrderShrinkRequest) GetUserInfoShrink() *string {
	return s.UserInfoShrink
}

func (s *ListHotelOrderShrinkRequest) SetPayloadShrink(v string) *ListHotelOrderShrinkRequest {
	s.PayloadShrink = &v
	return s
}

func (s *ListHotelOrderShrinkRequest) SetUserInfoShrink(v string) *ListHotelOrderShrinkRequest {
	s.UserInfoShrink = &v
	return s
}

func (s *ListHotelOrderShrinkRequest) Validate() error {
	return dara.Validate(s)
}
