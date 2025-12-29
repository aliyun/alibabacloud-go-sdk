// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddOrUpdateScreenSaverRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHotelId(v string) *AddOrUpdateScreenSaverRequest
	GetHotelId() *string
	SetHotelScreenSaver(v *AddOrUpdateScreenSaverRequestHotelScreenSaver) *AddOrUpdateScreenSaverRequest
	GetHotelScreenSaver() *AddOrUpdateScreenSaverRequestHotelScreenSaver
}

type AddOrUpdateScreenSaverRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// a7a3***013
	HotelId *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
	// This parameter is required.
	HotelScreenSaver *AddOrUpdateScreenSaverRequestHotelScreenSaver `json:"HotelScreenSaver,omitempty" xml:"HotelScreenSaver,omitempty" type:"Struct"`
}

func (s AddOrUpdateScreenSaverRequest) String() string {
	return dara.Prettify(s)
}

func (s AddOrUpdateScreenSaverRequest) GoString() string {
	return s.String()
}

func (s *AddOrUpdateScreenSaverRequest) GetHotelId() *string {
	return s.HotelId
}

func (s *AddOrUpdateScreenSaverRequest) GetHotelScreenSaver() *AddOrUpdateScreenSaverRequestHotelScreenSaver {
	return s.HotelScreenSaver
}

func (s *AddOrUpdateScreenSaverRequest) SetHotelId(v string) *AddOrUpdateScreenSaverRequest {
	s.HotelId = &v
	return s
}

func (s *AddOrUpdateScreenSaverRequest) SetHotelScreenSaver(v *AddOrUpdateScreenSaverRequestHotelScreenSaver) *AddOrUpdateScreenSaverRequest {
	s.HotelScreenSaver = v
	return s
}

func (s *AddOrUpdateScreenSaverRequest) Validate() error {
	if s.HotelScreenSaver != nil {
		if err := s.HotelScreenSaver.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AddOrUpdateScreenSaverRequestHotelScreenSaver struct {
	// example:
	//
	// xxx.png
	ScreenSaverPicUrl *string `json:"ScreenSaverPicUrl,omitempty" xml:"ScreenSaverPicUrl,omitempty"`
	// example:
	//
	// common-weather
	ScreenSaverStyle *string `json:"ScreenSaverStyle,omitempty" xml:"ScreenSaverStyle,omitempty"`
}

func (s AddOrUpdateScreenSaverRequestHotelScreenSaver) String() string {
	return dara.Prettify(s)
}

func (s AddOrUpdateScreenSaverRequestHotelScreenSaver) GoString() string {
	return s.String()
}

func (s *AddOrUpdateScreenSaverRequestHotelScreenSaver) GetScreenSaverPicUrl() *string {
	return s.ScreenSaverPicUrl
}

func (s *AddOrUpdateScreenSaverRequestHotelScreenSaver) GetScreenSaverStyle() *string {
	return s.ScreenSaverStyle
}

func (s *AddOrUpdateScreenSaverRequestHotelScreenSaver) SetScreenSaverPicUrl(v string) *AddOrUpdateScreenSaverRequestHotelScreenSaver {
	s.ScreenSaverPicUrl = &v
	return s
}

func (s *AddOrUpdateScreenSaverRequestHotelScreenSaver) SetScreenSaverStyle(v string) *AddOrUpdateScreenSaverRequestHotelScreenSaver {
	s.ScreenSaverStyle = &v
	return s
}

func (s *AddOrUpdateScreenSaverRequestHotelScreenSaver) Validate() error {
	return dara.Validate(s)
}
