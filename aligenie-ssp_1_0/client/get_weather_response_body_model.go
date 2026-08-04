// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWeatherResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *GetWeatherResponseBody
	GetCode() *int32
	SetMessage(v string) *GetWeatherResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetWeatherResponseBody
	GetRequestId() *string
	SetResult(v *GetWeatherResponseBodyResult) *GetWeatherResponseBody
	GetResult() *GetWeatherResponseBodyResult
}

type GetWeatherResponseBody struct {
	// HttpCode
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// error message
	//
	// example:
	//
	// 调用成功
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID
	//
	// example:
	//
	// F7E2****B7C94
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// model data
	Result *GetWeatherResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s GetWeatherResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetWeatherResponseBody) GoString() string {
	return s.String()
}

func (s *GetWeatherResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetWeatherResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetWeatherResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetWeatherResponseBody) GetResult() *GetWeatherResponseBodyResult {
	return s.Result
}

func (s *GetWeatherResponseBody) SetCode(v int32) *GetWeatherResponseBody {
	s.Code = &v
	return s
}

func (s *GetWeatherResponseBody) SetMessage(v string) *GetWeatherResponseBody {
	s.Message = &v
	return s
}

func (s *GetWeatherResponseBody) SetRequestId(v string) *GetWeatherResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetWeatherResponseBody) SetResult(v *GetWeatherResponseBodyResult) *GetWeatherResponseBody {
	s.Result = v
	return s
}

func (s *GetWeatherResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetWeatherResponseBodyResult struct {
	// Current weather
	CurrentMeteorology *GetWeatherResponseBodyResultCurrentMeteorology `json:"CurrentMeteorology,omitempty" xml:"CurrentMeteorology,omitempty" type:"Struct"`
}

func (s GetWeatherResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s GetWeatherResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetWeatherResponseBodyResult) GetCurrentMeteorology() *GetWeatherResponseBodyResultCurrentMeteorology {
	return s.CurrentMeteorology
}

func (s *GetWeatherResponseBodyResult) SetCurrentMeteorology(v *GetWeatherResponseBodyResultCurrentMeteorology) *GetWeatherResponseBodyResult {
	s.CurrentMeteorology = v
	return s
}

func (s *GetWeatherResponseBodyResult) Validate() error {
	if s.CurrentMeteorology != nil {
		if err := s.CurrentMeteorology.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetWeatherResponseBodyResultCurrentMeteorology struct {
	// Temperature
	Temperature *GetWeatherResponseBodyResultCurrentMeteorologyTemperature `json:"Temperature,omitempty" xml:"Temperature,omitempty" type:"Struct"`
	// Daytime weather
	Weather *GetWeatherResponseBodyResultCurrentMeteorologyWeather `json:"Weather,omitempty" xml:"Weather,omitempty" type:"Struct"`
}

func (s GetWeatherResponseBodyResultCurrentMeteorology) String() string {
	return dara.Prettify(s)
}

func (s GetWeatherResponseBodyResultCurrentMeteorology) GoString() string {
	return s.String()
}

func (s *GetWeatherResponseBodyResultCurrentMeteorology) GetTemperature() *GetWeatherResponseBodyResultCurrentMeteorologyTemperature {
	return s.Temperature
}

func (s *GetWeatherResponseBodyResultCurrentMeteorology) GetWeather() *GetWeatherResponseBodyResultCurrentMeteorologyWeather {
	return s.Weather
}

func (s *GetWeatherResponseBodyResultCurrentMeteorology) SetTemperature(v *GetWeatherResponseBodyResultCurrentMeteorologyTemperature) *GetWeatherResponseBodyResultCurrentMeteorology {
	s.Temperature = v
	return s
}

func (s *GetWeatherResponseBodyResultCurrentMeteorology) SetWeather(v *GetWeatherResponseBodyResultCurrentMeteorologyWeather) *GetWeatherResponseBodyResultCurrentMeteorology {
	s.Weather = v
	return s
}

func (s *GetWeatherResponseBodyResultCurrentMeteorology) Validate() error {
	if s.Temperature != nil {
		if err := s.Temperature.Validate(); err != nil {
			return err
		}
	}
	if s.Weather != nil {
		if err := s.Weather.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetWeatherResponseBodyResultCurrentMeteorologyTemperature struct {
	// Current temperature value
	//
	// example:
	//
	// 36
	Current *string `json:"Current,omitempty" xml:"Current,omitempty"`
	// Description of the current temperature value
	//
	// example:
	//
	// 36度
	CurrentDesc *string `json:"CurrentDesc,omitempty" xml:"CurrentDesc,omitempty"`
	// Maximum temperature value
	//
	// example:
	//
	// 37
	High *string `json:"High,omitempty" xml:"High,omitempty"`
	// Description of the maximum temperature value
	//
	// example:
	//
	// 37度
	HighDesc *string `json:"HighDesc,omitempty" xml:"HighDesc,omitempty"`
	// Temperature with logic, as follows:
	//
	// example:
	//
	// 41
	Logical *string `json:"Logical,omitempty" xml:"Logical,omitempty"`
	// Lowest temperature
	//
	// example:
	//
	// 28
	Low *string `json:"Low,omitempty" xml:"Low,omitempty"`
	// Description of the lowest temperature
	//
	// example:
	//
	// 28度
	LowDesc *string `json:"LowDesc,omitempty" xml:"LowDesc,omitempty"`
}

func (s GetWeatherResponseBodyResultCurrentMeteorologyTemperature) String() string {
	return dara.Prettify(s)
}

func (s GetWeatherResponseBodyResultCurrentMeteorologyTemperature) GoString() string {
	return s.String()
}

func (s *GetWeatherResponseBodyResultCurrentMeteorologyTemperature) GetCurrent() *string {
	return s.Current
}

func (s *GetWeatherResponseBodyResultCurrentMeteorologyTemperature) GetCurrentDesc() *string {
	return s.CurrentDesc
}

func (s *GetWeatherResponseBodyResultCurrentMeteorologyTemperature) GetHigh() *string {
	return s.High
}

func (s *GetWeatherResponseBodyResultCurrentMeteorologyTemperature) GetHighDesc() *string {
	return s.HighDesc
}

func (s *GetWeatherResponseBodyResultCurrentMeteorologyTemperature) GetLogical() *string {
	return s.Logical
}

func (s *GetWeatherResponseBodyResultCurrentMeteorologyTemperature) GetLow() *string {
	return s.Low
}

func (s *GetWeatherResponseBodyResultCurrentMeteorologyTemperature) GetLowDesc() *string {
	return s.LowDesc
}

func (s *GetWeatherResponseBodyResultCurrentMeteorologyTemperature) SetCurrent(v string) *GetWeatherResponseBodyResultCurrentMeteorologyTemperature {
	s.Current = &v
	return s
}

func (s *GetWeatherResponseBodyResultCurrentMeteorologyTemperature) SetCurrentDesc(v string) *GetWeatherResponseBodyResultCurrentMeteorologyTemperature {
	s.CurrentDesc = &v
	return s
}

func (s *GetWeatherResponseBodyResultCurrentMeteorologyTemperature) SetHigh(v string) *GetWeatherResponseBodyResultCurrentMeteorologyTemperature {
	s.High = &v
	return s
}

func (s *GetWeatherResponseBodyResultCurrentMeteorologyTemperature) SetHighDesc(v string) *GetWeatherResponseBodyResultCurrentMeteorologyTemperature {
	s.HighDesc = &v
	return s
}

func (s *GetWeatherResponseBodyResultCurrentMeteorologyTemperature) SetLogical(v string) *GetWeatherResponseBodyResultCurrentMeteorologyTemperature {
	s.Logical = &v
	return s
}

func (s *GetWeatherResponseBodyResultCurrentMeteorologyTemperature) SetLow(v string) *GetWeatherResponseBodyResultCurrentMeteorologyTemperature {
	s.Low = &v
	return s
}

func (s *GetWeatherResponseBodyResultCurrentMeteorologyTemperature) SetLowDesc(v string) *GetWeatherResponseBodyResultCurrentMeteorologyTemperature {
	s.LowDesc = &v
	return s
}

func (s *GetWeatherResponseBodyResultCurrentMeteorologyTemperature) Validate() error {
	return dara.Validate(s)
}

type GetWeatherResponseBodyResultCurrentMeteorologyWeather struct {
	// Weather code: for example, "000,100"
	//
	// example:
	//
	// 000
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Weather name description: "Sunny (000), Multicloud (100), Overcast (101), Rain (200), Light rain (201), Light to moderate rain (202), Moderate rain (203), Moderate to heavy rain (204), Heavy rain (205), Heavy to storm rain (206), Storm rain (207), Heavy storm rain (209), Severe storm rain (211), Showers (212), Thunderstorms (213), Freezing rain (214), Snow (300), Sleet (301), Snow showers (302), Light snow (303), Light to moderate snow (304), Moderate snow (305), Heavy snow (307), Blizzard (309), Fog (400), Dust (501), Sand blowing (502), Sandstorm (503), Severe sandstorm (504), Mostly sunny (000), Partly cloudy (100), Light showers (212), Lightning (213), Ice pellets (214), Thunderstorms with hail (215), Light snow showers (302), Freezing fog (400), Haze (500), Dust whirls (502), Localized showers (212), Thunderstorm (213), Ice needles (214), Hail (215), Intense showers (212)"
	//
	// example:
	//
	// 晴
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetWeatherResponseBodyResultCurrentMeteorologyWeather) String() string {
	return dara.Prettify(s)
}

func (s GetWeatherResponseBodyResultCurrentMeteorologyWeather) GoString() string {
	return s.String()
}

func (s *GetWeatherResponseBodyResultCurrentMeteorologyWeather) GetCode() *string {
	return s.Code
}

func (s *GetWeatherResponseBodyResultCurrentMeteorologyWeather) GetName() *string {
	return s.Name
}

func (s *GetWeatherResponseBodyResultCurrentMeteorologyWeather) SetCode(v string) *GetWeatherResponseBodyResultCurrentMeteorologyWeather {
	s.Code = &v
	return s
}

func (s *GetWeatherResponseBodyResultCurrentMeteorologyWeather) SetName(v string) *GetWeatherResponseBodyResultCurrentMeteorologyWeather {
	s.Name = &v
	return s
}

func (s *GetWeatherResponseBodyResultCurrentMeteorologyWeather) Validate() error {
	return dara.Validate(s)
}
