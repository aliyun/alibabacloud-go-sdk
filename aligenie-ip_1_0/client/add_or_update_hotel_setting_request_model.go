// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddOrUpdateHotelSettingRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHotelDeviceModeList(v []*string) *AddOrUpdateHotelSettingRequest
	GetHotelDeviceModeList() []*string
	SetHotelId(v string) *AddOrUpdateHotelSettingRequest
	GetHotelId() *string
	SetHotelScreenSaver(v *AddOrUpdateHotelSettingRequestHotelScreenSaver) *AddOrUpdateHotelSettingRequest
	GetHotelScreenSaver() *AddOrUpdateHotelSettingRequestHotelScreenSaver
	SetNightMode(v *AddOrUpdateHotelSettingRequestNightMode) *AddOrUpdateHotelSettingRequest
	GetNightMode() *AddOrUpdateHotelSettingRequestNightMode
	SetSettingType(v string) *AddOrUpdateHotelSettingRequest
	GetSettingType() *string
	SetValue(v string) *AddOrUpdateHotelSettingRequest
	GetValue() *string
}

type AddOrUpdateHotelSettingRequest struct {
	HotelDeviceModeList []*string `json:"HotelDeviceModeList,omitempty" xml:"HotelDeviceModeList,omitempty" type:"Repeated"`
	// example:
	//
	// a7a3***013
	HotelId          *string                                         `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
	HotelScreenSaver *AddOrUpdateHotelSettingRequestHotelScreenSaver `json:"HotelScreenSaver,omitempty" xml:"HotelScreenSaver,omitempty" type:"Struct"`
	NightMode        *AddOrUpdateHotelSettingRequestNightMode        `json:"NightMode,omitempty" xml:"NightMode,omitempty" type:"Struct"`
	// example:
	//
	// SCREENSAVER
	SettingType *string `json:"SettingType,omitempty" xml:"SettingType,omitempty"`
	Value       *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s AddOrUpdateHotelSettingRequest) String() string {
	return dara.Prettify(s)
}

func (s AddOrUpdateHotelSettingRequest) GoString() string {
	return s.String()
}

func (s *AddOrUpdateHotelSettingRequest) GetHotelDeviceModeList() []*string {
	return s.HotelDeviceModeList
}

func (s *AddOrUpdateHotelSettingRequest) GetHotelId() *string {
	return s.HotelId
}

func (s *AddOrUpdateHotelSettingRequest) GetHotelScreenSaver() *AddOrUpdateHotelSettingRequestHotelScreenSaver {
	return s.HotelScreenSaver
}

func (s *AddOrUpdateHotelSettingRequest) GetNightMode() *AddOrUpdateHotelSettingRequestNightMode {
	return s.NightMode
}

func (s *AddOrUpdateHotelSettingRequest) GetSettingType() *string {
	return s.SettingType
}

func (s *AddOrUpdateHotelSettingRequest) GetValue() *string {
	return s.Value
}

func (s *AddOrUpdateHotelSettingRequest) SetHotelDeviceModeList(v []*string) *AddOrUpdateHotelSettingRequest {
	s.HotelDeviceModeList = v
	return s
}

func (s *AddOrUpdateHotelSettingRequest) SetHotelId(v string) *AddOrUpdateHotelSettingRequest {
	s.HotelId = &v
	return s
}

func (s *AddOrUpdateHotelSettingRequest) SetHotelScreenSaver(v *AddOrUpdateHotelSettingRequestHotelScreenSaver) *AddOrUpdateHotelSettingRequest {
	s.HotelScreenSaver = v
	return s
}

func (s *AddOrUpdateHotelSettingRequest) SetNightMode(v *AddOrUpdateHotelSettingRequestNightMode) *AddOrUpdateHotelSettingRequest {
	s.NightMode = v
	return s
}

func (s *AddOrUpdateHotelSettingRequest) SetSettingType(v string) *AddOrUpdateHotelSettingRequest {
	s.SettingType = &v
	return s
}

func (s *AddOrUpdateHotelSettingRequest) SetValue(v string) *AddOrUpdateHotelSettingRequest {
	s.Value = &v
	return s
}

func (s *AddOrUpdateHotelSettingRequest) Validate() error {
	if s.HotelScreenSaver != nil {
		if err := s.HotelScreenSaver.Validate(); err != nil {
			return err
		}
	}
	if s.NightMode != nil {
		if err := s.NightMode.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AddOrUpdateHotelSettingRequestHotelScreenSaver struct {
	// example:
	//
	// https://a****jpg
	ScreenSaverPicUrl *string `json:"ScreenSaverPicUrl,omitempty" xml:"ScreenSaverPicUrl,omitempty"`
	// example:
	//
	// common-weather
	ScreenSaverStyle *string `json:"ScreenSaverStyle,omitempty" xml:"ScreenSaverStyle,omitempty"`
}

func (s AddOrUpdateHotelSettingRequestHotelScreenSaver) String() string {
	return dara.Prettify(s)
}

func (s AddOrUpdateHotelSettingRequestHotelScreenSaver) GoString() string {
	return s.String()
}

func (s *AddOrUpdateHotelSettingRequestHotelScreenSaver) GetScreenSaverPicUrl() *string {
	return s.ScreenSaverPicUrl
}

func (s *AddOrUpdateHotelSettingRequestHotelScreenSaver) GetScreenSaverStyle() *string {
	return s.ScreenSaverStyle
}

func (s *AddOrUpdateHotelSettingRequestHotelScreenSaver) SetScreenSaverPicUrl(v string) *AddOrUpdateHotelSettingRequestHotelScreenSaver {
	s.ScreenSaverPicUrl = &v
	return s
}

func (s *AddOrUpdateHotelSettingRequestHotelScreenSaver) SetScreenSaverStyle(v string) *AddOrUpdateHotelSettingRequestHotelScreenSaver {
	s.ScreenSaverStyle = &v
	return s
}

func (s *AddOrUpdateHotelSettingRequestHotelScreenSaver) Validate() error {
	return dara.Validate(s)
}

type AddOrUpdateHotelSettingRequestNightMode struct {
	DefaultBright *string `json:"DefaultBright,omitempty" xml:"DefaultBright,omitempty"`
	DefaultVolume *string `json:"DefaultVolume,omitempty" xml:"DefaultVolume,omitempty"`
	// example:
	//
	// true
	Enable *bool `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// example:
	//
	// 22:00
	End *string `json:"End,omitempty" xml:"End,omitempty"`
	// example:
	//
	// screenoff
	StandbyAction *string `json:"StandbyAction,omitempty" xml:"StandbyAction,omitempty"`
	// example:
	//
	// 7:00
	Start *string `json:"Start,omitempty" xml:"Start,omitempty"`
}

func (s AddOrUpdateHotelSettingRequestNightMode) String() string {
	return dara.Prettify(s)
}

func (s AddOrUpdateHotelSettingRequestNightMode) GoString() string {
	return s.String()
}

func (s *AddOrUpdateHotelSettingRequestNightMode) GetDefaultBright() *string {
	return s.DefaultBright
}

func (s *AddOrUpdateHotelSettingRequestNightMode) GetDefaultVolume() *string {
	return s.DefaultVolume
}

func (s *AddOrUpdateHotelSettingRequestNightMode) GetEnable() *bool {
	return s.Enable
}

func (s *AddOrUpdateHotelSettingRequestNightMode) GetEnd() *string {
	return s.End
}

func (s *AddOrUpdateHotelSettingRequestNightMode) GetStandbyAction() *string {
	return s.StandbyAction
}

func (s *AddOrUpdateHotelSettingRequestNightMode) GetStart() *string {
	return s.Start
}

func (s *AddOrUpdateHotelSettingRequestNightMode) SetDefaultBright(v string) *AddOrUpdateHotelSettingRequestNightMode {
	s.DefaultBright = &v
	return s
}

func (s *AddOrUpdateHotelSettingRequestNightMode) SetDefaultVolume(v string) *AddOrUpdateHotelSettingRequestNightMode {
	s.DefaultVolume = &v
	return s
}

func (s *AddOrUpdateHotelSettingRequestNightMode) SetEnable(v bool) *AddOrUpdateHotelSettingRequestNightMode {
	s.Enable = &v
	return s
}

func (s *AddOrUpdateHotelSettingRequestNightMode) SetEnd(v string) *AddOrUpdateHotelSettingRequestNightMode {
	s.End = &v
	return s
}

func (s *AddOrUpdateHotelSettingRequestNightMode) SetStandbyAction(v string) *AddOrUpdateHotelSettingRequestNightMode {
	s.StandbyAction = &v
	return s
}

func (s *AddOrUpdateHotelSettingRequestNightMode) SetStart(v string) *AddOrUpdateHotelSettingRequestNightMode {
	s.Start = &v
	return s
}

func (s *AddOrUpdateHotelSettingRequestNightMode) Validate() error {
	return dara.Validate(s)
}
