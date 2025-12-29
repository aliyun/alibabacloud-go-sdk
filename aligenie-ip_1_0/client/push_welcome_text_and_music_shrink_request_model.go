// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPushWelcomeTextAndMusicShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHotelId(v string) *PushWelcomeTextAndMusicShrinkRequest
	GetHotelId() *string
	SetRoomName(v string) *PushWelcomeTextAndMusicShrinkRequest
	GetRoomName() *string
	SetRoomNo(v string) *PushWelcomeTextAndMusicShrinkRequest
	GetRoomNo() *string
	SetTemplateVariableShrink(v string) *PushWelcomeTextAndMusicShrinkRequest
	GetTemplateVariableShrink() *string
}

type PushWelcomeTextAndMusicShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// af7***536
	HotelId  *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
	RoomName *string `json:"RoomName,omitempty" xml:"RoomName,omitempty"`
	// example:
	//
	// 1211
	RoomNo                 *string `json:"RoomNo,omitempty" xml:"RoomNo,omitempty"`
	TemplateVariableShrink *string `json:"TemplateVariable,omitempty" xml:"TemplateVariable,omitempty"`
}

func (s PushWelcomeTextAndMusicShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s PushWelcomeTextAndMusicShrinkRequest) GoString() string {
	return s.String()
}

func (s *PushWelcomeTextAndMusicShrinkRequest) GetHotelId() *string {
	return s.HotelId
}

func (s *PushWelcomeTextAndMusicShrinkRequest) GetRoomName() *string {
	return s.RoomName
}

func (s *PushWelcomeTextAndMusicShrinkRequest) GetRoomNo() *string {
	return s.RoomNo
}

func (s *PushWelcomeTextAndMusicShrinkRequest) GetTemplateVariableShrink() *string {
	return s.TemplateVariableShrink
}

func (s *PushWelcomeTextAndMusicShrinkRequest) SetHotelId(v string) *PushWelcomeTextAndMusicShrinkRequest {
	s.HotelId = &v
	return s
}

func (s *PushWelcomeTextAndMusicShrinkRequest) SetRoomName(v string) *PushWelcomeTextAndMusicShrinkRequest {
	s.RoomName = &v
	return s
}

func (s *PushWelcomeTextAndMusicShrinkRequest) SetRoomNo(v string) *PushWelcomeTextAndMusicShrinkRequest {
	s.RoomNo = &v
	return s
}

func (s *PushWelcomeTextAndMusicShrinkRequest) SetTemplateVariableShrink(v string) *PushWelcomeTextAndMusicShrinkRequest {
	s.TemplateVariableShrink = &v
	return s
}

func (s *PushWelcomeTextAndMusicShrinkRequest) Validate() error {
	return dara.Validate(s)
}
