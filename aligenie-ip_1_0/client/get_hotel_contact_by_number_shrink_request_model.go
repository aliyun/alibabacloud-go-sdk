// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHotelContactByNumberShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNumber(v string) *GetHotelContactByNumberShrinkRequest
	GetNumber() *string
	SetUserInfoShrink(v string) *GetHotelContactByNumberShrinkRequest
	GetUserInfoShrink() *string
}

type GetHotelContactByNumberShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 101
	Number         *string `json:"Number,omitempty" xml:"Number,omitempty"`
	UserInfoShrink *string `json:"UserInfo,omitempty" xml:"UserInfo,omitempty"`
}

func (s GetHotelContactByNumberShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetHotelContactByNumberShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetHotelContactByNumberShrinkRequest) GetNumber() *string {
	return s.Number
}

func (s *GetHotelContactByNumberShrinkRequest) GetUserInfoShrink() *string {
	return s.UserInfoShrink
}

func (s *GetHotelContactByNumberShrinkRequest) SetNumber(v string) *GetHotelContactByNumberShrinkRequest {
	s.Number = &v
	return s
}

func (s *GetHotelContactByNumberShrinkRequest) SetUserInfoShrink(v string) *GetHotelContactByNumberShrinkRequest {
	s.UserInfoShrink = &v
	return s
}

func (s *GetHotelContactByNumberShrinkRequest) Validate() error {
	return dara.Validate(s)
}
