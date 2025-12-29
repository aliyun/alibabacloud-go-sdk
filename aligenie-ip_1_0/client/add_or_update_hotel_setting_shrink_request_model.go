// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddOrUpdateHotelSettingShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHotelDeviceModeListShrink(v string) *AddOrUpdateHotelSettingShrinkRequest
	GetHotelDeviceModeListShrink() *string
	SetHotelId(v string) *AddOrUpdateHotelSettingShrinkRequest
	GetHotelId() *string
	SetHotelScreenSaverShrink(v string) *AddOrUpdateHotelSettingShrinkRequest
	GetHotelScreenSaverShrink() *string
	SetNightModeShrink(v string) *AddOrUpdateHotelSettingShrinkRequest
	GetNightModeShrink() *string
	SetSettingType(v string) *AddOrUpdateHotelSettingShrinkRequest
	GetSettingType() *string
	SetValue(v string) *AddOrUpdateHotelSettingShrinkRequest
	GetValue() *string
}

type AddOrUpdateHotelSettingShrinkRequest struct {
	HotelDeviceModeListShrink *string `json:"HotelDeviceModeList,omitempty" xml:"HotelDeviceModeList,omitempty"`
	// example:
	//
	// a7a3***013
	HotelId                *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
	HotelScreenSaverShrink *string `json:"HotelScreenSaver,omitempty" xml:"HotelScreenSaver,omitempty"`
	NightModeShrink        *string `json:"NightMode,omitempty" xml:"NightMode,omitempty"`
	// example:
	//
	// SCREENSAVER
	SettingType *string `json:"SettingType,omitempty" xml:"SettingType,omitempty"`
	Value       *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s AddOrUpdateHotelSettingShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s AddOrUpdateHotelSettingShrinkRequest) GoString() string {
	return s.String()
}

func (s *AddOrUpdateHotelSettingShrinkRequest) GetHotelDeviceModeListShrink() *string {
	return s.HotelDeviceModeListShrink
}

func (s *AddOrUpdateHotelSettingShrinkRequest) GetHotelId() *string {
	return s.HotelId
}

func (s *AddOrUpdateHotelSettingShrinkRequest) GetHotelScreenSaverShrink() *string {
	return s.HotelScreenSaverShrink
}

func (s *AddOrUpdateHotelSettingShrinkRequest) GetNightModeShrink() *string {
	return s.NightModeShrink
}

func (s *AddOrUpdateHotelSettingShrinkRequest) GetSettingType() *string {
	return s.SettingType
}

func (s *AddOrUpdateHotelSettingShrinkRequest) GetValue() *string {
	return s.Value
}

func (s *AddOrUpdateHotelSettingShrinkRequest) SetHotelDeviceModeListShrink(v string) *AddOrUpdateHotelSettingShrinkRequest {
	s.HotelDeviceModeListShrink = &v
	return s
}

func (s *AddOrUpdateHotelSettingShrinkRequest) SetHotelId(v string) *AddOrUpdateHotelSettingShrinkRequest {
	s.HotelId = &v
	return s
}

func (s *AddOrUpdateHotelSettingShrinkRequest) SetHotelScreenSaverShrink(v string) *AddOrUpdateHotelSettingShrinkRequest {
	s.HotelScreenSaverShrink = &v
	return s
}

func (s *AddOrUpdateHotelSettingShrinkRequest) SetNightModeShrink(v string) *AddOrUpdateHotelSettingShrinkRequest {
	s.NightModeShrink = &v
	return s
}

func (s *AddOrUpdateHotelSettingShrinkRequest) SetSettingType(v string) *AddOrUpdateHotelSettingShrinkRequest {
	s.SettingType = &v
	return s
}

func (s *AddOrUpdateHotelSettingShrinkRequest) SetValue(v string) *AddOrUpdateHotelSettingShrinkRequest {
	s.Value = &v
	return s
}

func (s *AddOrUpdateHotelSettingShrinkRequest) Validate() error {
	return dara.Validate(s)
}
