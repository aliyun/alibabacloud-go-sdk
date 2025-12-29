// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPushWelcomeTextAndMusicRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHotelId(v string) *PushWelcomeTextAndMusicRequest
	GetHotelId() *string
	SetRoomName(v string) *PushWelcomeTextAndMusicRequest
	GetRoomName() *string
	SetRoomNo(v string) *PushWelcomeTextAndMusicRequest
	GetRoomNo() *string
	SetTemplateVariable(v map[string]*string) *PushWelcomeTextAndMusicRequest
	GetTemplateVariable() map[string]*string
}

type PushWelcomeTextAndMusicRequest struct {
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
	RoomNo           *string            `json:"RoomNo,omitempty" xml:"RoomNo,omitempty"`
	TemplateVariable map[string]*string `json:"TemplateVariable,omitempty" xml:"TemplateVariable,omitempty"`
}

func (s PushWelcomeTextAndMusicRequest) String() string {
	return dara.Prettify(s)
}

func (s PushWelcomeTextAndMusicRequest) GoString() string {
	return s.String()
}

func (s *PushWelcomeTextAndMusicRequest) GetHotelId() *string {
	return s.HotelId
}

func (s *PushWelcomeTextAndMusicRequest) GetRoomName() *string {
	return s.RoomName
}

func (s *PushWelcomeTextAndMusicRequest) GetRoomNo() *string {
	return s.RoomNo
}

func (s *PushWelcomeTextAndMusicRequest) GetTemplateVariable() map[string]*string {
	return s.TemplateVariable
}

func (s *PushWelcomeTextAndMusicRequest) SetHotelId(v string) *PushWelcomeTextAndMusicRequest {
	s.HotelId = &v
	return s
}

func (s *PushWelcomeTextAndMusicRequest) SetRoomName(v string) *PushWelcomeTextAndMusicRequest {
	s.RoomName = &v
	return s
}

func (s *PushWelcomeTextAndMusicRequest) SetRoomNo(v string) *PushWelcomeTextAndMusicRequest {
	s.RoomNo = &v
	return s
}

func (s *PushWelcomeTextAndMusicRequest) SetTemplateVariable(v map[string]*string) *PushWelcomeTextAndMusicRequest {
	s.TemplateVariable = v
	return s
}

func (s *PushWelcomeTextAndMusicRequest) Validate() error {
	return dara.Validate(s)
}
