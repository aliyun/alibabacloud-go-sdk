// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHotelSettingResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *GetHotelSettingResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetHotelSettingResponseBody
	GetRequestId() *string
	SetResult(v *GetHotelSettingResponseBodyResult) *GetHotelSettingResponseBody
	GetResult() *GetHotelSettingResponseBodyResult
	SetStatusCode(v int32) *GetHotelSettingResponseBody
	GetStatusCode() *int32
}

type GetHotelSettingResponseBody struct {
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// F7E2****B7C94
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *GetHotelSettingResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// example:
	//
	// 200
	StatusCode *int32 `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
}

func (s GetHotelSettingResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetHotelSettingResponseBody) GoString() string {
	return s.String()
}

func (s *GetHotelSettingResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetHotelSettingResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetHotelSettingResponseBody) GetResult() *GetHotelSettingResponseBodyResult {
	return s.Result
}

func (s *GetHotelSettingResponseBody) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetHotelSettingResponseBody) SetMessage(v string) *GetHotelSettingResponseBody {
	s.Message = &v
	return s
}

func (s *GetHotelSettingResponseBody) SetRequestId(v string) *GetHotelSettingResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetHotelSettingResponseBody) SetResult(v *GetHotelSettingResponseBodyResult) *GetHotelSettingResponseBody {
	s.Result = v
	return s
}

func (s *GetHotelSettingResponseBody) SetStatusCode(v int32) *GetHotelSettingResponseBody {
	s.StatusCode = &v
	return s
}

func (s *GetHotelSettingResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetHotelSettingResponseBodyResult struct {
	// example:
	//
	// 0
	DeleteToken *int64 `json:"DeleteToken,omitempty" xml:"DeleteToken,omitempty"`
	// example:
	//
	// {}
	ExtInfo             *string   `json:"ExtInfo,omitempty" xml:"ExtInfo,omitempty"`
	HotelDeviceModeList []*string `json:"HotelDeviceModeList,omitempty" xml:"HotelDeviceModeList,omitempty" type:"Repeated"`
	// example:
	//
	// af7***536
	HotelId          *string                                            `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
	HotelScreenSaver *GetHotelSettingResponseBodyResultHotelScreenSaver `json:"HotelScreenSaver,omitempty" xml:"HotelScreenSaver,omitempty" type:"Struct"`
	NightMode        *GetHotelSettingResponseBodyResultNightMode        `json:"NightMode,omitempty" xml:"NightMode,omitempty" type:"Struct"`
	// example:
	//
	// SCREENSAVER
	SettingType *string `json:"SettingType,omitempty" xml:"SettingType,omitempty"`
	Value       *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetHotelSettingResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s GetHotelSettingResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetHotelSettingResponseBodyResult) GetDeleteToken() *int64 {
	return s.DeleteToken
}

func (s *GetHotelSettingResponseBodyResult) GetExtInfo() *string {
	return s.ExtInfo
}

func (s *GetHotelSettingResponseBodyResult) GetHotelDeviceModeList() []*string {
	return s.HotelDeviceModeList
}

func (s *GetHotelSettingResponseBodyResult) GetHotelId() *string {
	return s.HotelId
}

func (s *GetHotelSettingResponseBodyResult) GetHotelScreenSaver() *GetHotelSettingResponseBodyResultHotelScreenSaver {
	return s.HotelScreenSaver
}

func (s *GetHotelSettingResponseBodyResult) GetNightMode() *GetHotelSettingResponseBodyResultNightMode {
	return s.NightMode
}

func (s *GetHotelSettingResponseBodyResult) GetSettingType() *string {
	return s.SettingType
}

func (s *GetHotelSettingResponseBodyResult) GetValue() *string {
	return s.Value
}

func (s *GetHotelSettingResponseBodyResult) SetDeleteToken(v int64) *GetHotelSettingResponseBodyResult {
	s.DeleteToken = &v
	return s
}

func (s *GetHotelSettingResponseBodyResult) SetExtInfo(v string) *GetHotelSettingResponseBodyResult {
	s.ExtInfo = &v
	return s
}

func (s *GetHotelSettingResponseBodyResult) SetHotelDeviceModeList(v []*string) *GetHotelSettingResponseBodyResult {
	s.HotelDeviceModeList = v
	return s
}

func (s *GetHotelSettingResponseBodyResult) SetHotelId(v string) *GetHotelSettingResponseBodyResult {
	s.HotelId = &v
	return s
}

func (s *GetHotelSettingResponseBodyResult) SetHotelScreenSaver(v *GetHotelSettingResponseBodyResultHotelScreenSaver) *GetHotelSettingResponseBodyResult {
	s.HotelScreenSaver = v
	return s
}

func (s *GetHotelSettingResponseBodyResult) SetNightMode(v *GetHotelSettingResponseBodyResultNightMode) *GetHotelSettingResponseBodyResult {
	s.NightMode = v
	return s
}

func (s *GetHotelSettingResponseBodyResult) SetSettingType(v string) *GetHotelSettingResponseBodyResult {
	s.SettingType = &v
	return s
}

func (s *GetHotelSettingResponseBodyResult) SetValue(v string) *GetHotelSettingResponseBodyResult {
	s.Value = &v
	return s
}

func (s *GetHotelSettingResponseBodyResult) Validate() error {
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

type GetHotelSettingResponseBodyResultHotelScreenSaver struct {
	// example:
	//
	// https://a***png
	ScreenSaverPicUrl *string `json:"ScreenSaverPicUrl,omitempty" xml:"ScreenSaverPicUrl,omitempty"`
	// example:
	//
	// common-weather
	ScreenSaverStyle *string `json:"ScreenSaverStyle,omitempty" xml:"ScreenSaverStyle,omitempty"`
}

func (s GetHotelSettingResponseBodyResultHotelScreenSaver) String() string {
	return dara.Prettify(s)
}

func (s GetHotelSettingResponseBodyResultHotelScreenSaver) GoString() string {
	return s.String()
}

func (s *GetHotelSettingResponseBodyResultHotelScreenSaver) GetScreenSaverPicUrl() *string {
	return s.ScreenSaverPicUrl
}

func (s *GetHotelSettingResponseBodyResultHotelScreenSaver) GetScreenSaverStyle() *string {
	return s.ScreenSaverStyle
}

func (s *GetHotelSettingResponseBodyResultHotelScreenSaver) SetScreenSaverPicUrl(v string) *GetHotelSettingResponseBodyResultHotelScreenSaver {
	s.ScreenSaverPicUrl = &v
	return s
}

func (s *GetHotelSettingResponseBodyResultHotelScreenSaver) SetScreenSaverStyle(v string) *GetHotelSettingResponseBodyResultHotelScreenSaver {
	s.ScreenSaverStyle = &v
	return s
}

func (s *GetHotelSettingResponseBodyResultHotelScreenSaver) Validate() error {
	return dara.Validate(s)
}

type GetHotelSettingResponseBodyResultNightMode struct {
	// 夜间模式下的默认亮度
	DefaultBright *string `json:"DefaultBright,omitempty" xml:"DefaultBright,omitempty"`
	// 夜间模式下的默认音量
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
	// 07:00
	Start *string `json:"Start,omitempty" xml:"Start,omitempty"`
}

func (s GetHotelSettingResponseBodyResultNightMode) String() string {
	return dara.Prettify(s)
}

func (s GetHotelSettingResponseBodyResultNightMode) GoString() string {
	return s.String()
}

func (s *GetHotelSettingResponseBodyResultNightMode) GetDefaultBright() *string {
	return s.DefaultBright
}

func (s *GetHotelSettingResponseBodyResultNightMode) GetDefaultVolume() *string {
	return s.DefaultVolume
}

func (s *GetHotelSettingResponseBodyResultNightMode) GetEnable() *bool {
	return s.Enable
}

func (s *GetHotelSettingResponseBodyResultNightMode) GetEnd() *string {
	return s.End
}

func (s *GetHotelSettingResponseBodyResultNightMode) GetStandbyAction() *string {
	return s.StandbyAction
}

func (s *GetHotelSettingResponseBodyResultNightMode) GetStart() *string {
	return s.Start
}

func (s *GetHotelSettingResponseBodyResultNightMode) SetDefaultBright(v string) *GetHotelSettingResponseBodyResultNightMode {
	s.DefaultBright = &v
	return s
}

func (s *GetHotelSettingResponseBodyResultNightMode) SetDefaultVolume(v string) *GetHotelSettingResponseBodyResultNightMode {
	s.DefaultVolume = &v
	return s
}

func (s *GetHotelSettingResponseBodyResultNightMode) SetEnable(v bool) *GetHotelSettingResponseBodyResultNightMode {
	s.Enable = &v
	return s
}

func (s *GetHotelSettingResponseBodyResultNightMode) SetEnd(v string) *GetHotelSettingResponseBodyResultNightMode {
	s.End = &v
	return s
}

func (s *GetHotelSettingResponseBodyResultNightMode) SetStandbyAction(v string) *GetHotelSettingResponseBodyResultNightMode {
	s.StandbyAction = &v
	return s
}

func (s *GetHotelSettingResponseBodyResultNightMode) SetStart(v string) *GetHotelSettingResponseBodyResultNightMode {
	s.Start = &v
	return s
}

func (s *GetHotelSettingResponseBodyResultNightMode) Validate() error {
	return dara.Validate(s)
}
