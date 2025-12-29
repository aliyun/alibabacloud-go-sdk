// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHotelSceneItemShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPayloadShrink(v string) *ListHotelSceneItemShrinkRequest
	GetPayloadShrink() *string
	SetUserInfoShrink(v string) *ListHotelSceneItemShrinkRequest
	GetUserInfoShrink() *string
}

type ListHotelSceneItemShrinkRequest struct {
	// This parameter is required.
	PayloadShrink *string `json:"Payload,omitempty" xml:"Payload,omitempty"`
	// This parameter is required.
	UserInfoShrink *string `json:"UserInfo,omitempty" xml:"UserInfo,omitempty"`
}

func (s ListHotelSceneItemShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListHotelSceneItemShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListHotelSceneItemShrinkRequest) GetPayloadShrink() *string {
	return s.PayloadShrink
}

func (s *ListHotelSceneItemShrinkRequest) GetUserInfoShrink() *string {
	return s.UserInfoShrink
}

func (s *ListHotelSceneItemShrinkRequest) SetPayloadShrink(v string) *ListHotelSceneItemShrinkRequest {
	s.PayloadShrink = &v
	return s
}

func (s *ListHotelSceneItemShrinkRequest) SetUserInfoShrink(v string) *ListHotelSceneItemShrinkRequest {
	s.UserInfoShrink = &v
	return s
}

func (s *ListHotelSceneItemShrinkRequest) Validate() error {
	return dara.Validate(s)
}
