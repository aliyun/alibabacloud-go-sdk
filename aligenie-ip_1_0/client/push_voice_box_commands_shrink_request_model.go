// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPushVoiceBoxCommandsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCommandsShrink(v string) *PushVoiceBoxCommandsShrinkRequest
	GetCommandsShrink() *string
	SetHotelId(v string) *PushVoiceBoxCommandsShrinkRequest
	GetHotelId() *string
	SetRoomNo(v string) *PushVoiceBoxCommandsShrinkRequest
	GetRoomNo() *string
}

type PushVoiceBoxCommandsShrinkRequest struct {
	// This parameter is required.
	CommandsShrink *string `json:"Commands,omitempty" xml:"Commands,omitempty"`
	// This parameter is required.
	HotelId *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
	// This parameter is required.
	RoomNo *string `json:"RoomNo,omitempty" xml:"RoomNo,omitempty"`
}

func (s PushVoiceBoxCommandsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s PushVoiceBoxCommandsShrinkRequest) GoString() string {
	return s.String()
}

func (s *PushVoiceBoxCommandsShrinkRequest) GetCommandsShrink() *string {
	return s.CommandsShrink
}

func (s *PushVoiceBoxCommandsShrinkRequest) GetHotelId() *string {
	return s.HotelId
}

func (s *PushVoiceBoxCommandsShrinkRequest) GetRoomNo() *string {
	return s.RoomNo
}

func (s *PushVoiceBoxCommandsShrinkRequest) SetCommandsShrink(v string) *PushVoiceBoxCommandsShrinkRequest {
	s.CommandsShrink = &v
	return s
}

func (s *PushVoiceBoxCommandsShrinkRequest) SetHotelId(v string) *PushVoiceBoxCommandsShrinkRequest {
	s.HotelId = &v
	return s
}

func (s *PushVoiceBoxCommandsShrinkRequest) SetRoomNo(v string) *PushVoiceBoxCommandsShrinkRequest {
	s.RoomNo = &v
	return s
}

func (s *PushVoiceBoxCommandsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
