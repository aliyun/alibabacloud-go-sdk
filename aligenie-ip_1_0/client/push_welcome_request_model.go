// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPushWelcomeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHotelId(v string) *PushWelcomeRequest
	GetHotelId() *string
	SetRoomName(v string) *PushWelcomeRequest
	GetRoomName() *string
	SetRoomNo(v string) *PushWelcomeRequest
	GetRoomNo() *string
	SetWelcomeMusicUrl(v string) *PushWelcomeRequest
	GetWelcomeMusicUrl() *string
	SetWelcomeText(v string) *PushWelcomeRequest
	GetWelcomeText() *string
}

type PushWelcomeRequest struct {
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
	RoomNo *string `json:"RoomNo,omitempty" xml:"RoomNo,omitempty"`
	// example:
	//
	// http://ailabsaicloudservice.alicdn.com/tmp/a.wav
	WelcomeMusicUrl *string `json:"WelcomeMusicUrl,omitempty" xml:"WelcomeMusicUrl,omitempty"`
	// This parameter is required.
	WelcomeText *string `json:"WelcomeText,omitempty" xml:"WelcomeText,omitempty"`
}

func (s PushWelcomeRequest) String() string {
	return dara.Prettify(s)
}

func (s PushWelcomeRequest) GoString() string {
	return s.String()
}

func (s *PushWelcomeRequest) GetHotelId() *string {
	return s.HotelId
}

func (s *PushWelcomeRequest) GetRoomName() *string {
	return s.RoomName
}

func (s *PushWelcomeRequest) GetRoomNo() *string {
	return s.RoomNo
}

func (s *PushWelcomeRequest) GetWelcomeMusicUrl() *string {
	return s.WelcomeMusicUrl
}

func (s *PushWelcomeRequest) GetWelcomeText() *string {
	return s.WelcomeText
}

func (s *PushWelcomeRequest) SetHotelId(v string) *PushWelcomeRequest {
	s.HotelId = &v
	return s
}

func (s *PushWelcomeRequest) SetRoomName(v string) *PushWelcomeRequest {
	s.RoomName = &v
	return s
}

func (s *PushWelcomeRequest) SetRoomNo(v string) *PushWelcomeRequest {
	s.RoomNo = &v
	return s
}

func (s *PushWelcomeRequest) SetWelcomeMusicUrl(v string) *PushWelcomeRequest {
	s.WelcomeMusicUrl = &v
	return s
}

func (s *PushWelcomeRequest) SetWelcomeText(v string) *PushWelcomeRequest {
	s.WelcomeText = &v
	return s
}

func (s *PushWelcomeRequest) Validate() error {
	return dara.Validate(s)
}
