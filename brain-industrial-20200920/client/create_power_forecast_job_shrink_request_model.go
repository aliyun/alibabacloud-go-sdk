// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePowerForecastJobShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBusinessKey(v string) *CreatePowerForecastJobShrinkRequest
	GetBusinessKey() *string
	SetDataMode(v string) *CreatePowerForecastJobShrinkRequest
	GetDataMode() *string
	SetDeviceType(v string) *CreatePowerForecastJobShrinkRequest
	GetDeviceType() *string
	SetDuration(v int32) *CreatePowerForecastJobShrinkRequest
	GetDuration() *int32
	SetForecastHorizon(v string) *CreatePowerForecastJobShrinkRequest
	GetForecastHorizon() *string
	SetFreq(v string) *CreatePowerForecastJobShrinkRequest
	GetFreq() *string
	SetHistoryDataShrink(v string) *CreatePowerForecastJobShrinkRequest
	GetHistoryDataShrink() *string
	SetLocationShrink(v string) *CreatePowerForecastJobShrinkRequest
	GetLocationShrink() *string
	SetModelVersion(v string) *CreatePowerForecastJobShrinkRequest
	GetModelVersion() *string
	SetRunDate(v string) *CreatePowerForecastJobShrinkRequest
	GetRunDate() *string
	SetSystemType(v string) *CreatePowerForecastJobShrinkRequest
	GetSystemType() *string
	SetTimeZone(v string) *CreatePowerForecastJobShrinkRequest
	GetTimeZone() *string
}

type CreatePowerForecastJobShrinkRequest struct {
	BusinessKey *string `json:"BusinessKey,omitempty" xml:"BusinessKey,omitempty"`
	// example:
	//
	// FULL
	DataMode *string `json:"DataMode,omitempty" xml:"DataMode,omitempty"`
	// example:
	//
	// electricityMeter
	DeviceType *string `json:"DeviceType,omitempty" xml:"DeviceType,omitempty"`
	// example:
	//
	// 1
	Duration *int32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// example:
	//
	// DAY_AHEAD
	ForecastHorizon *string `json:"ForecastHorizon,omitempty" xml:"ForecastHorizon,omitempty"`
	// example:
	//
	// FIFTEEN_MIN
	Freq              *string `json:"Freq,omitempty" xml:"Freq,omitempty"`
	HistoryDataShrink *string `json:"HistoryData,omitempty" xml:"HistoryData,omitempty"`
	LocationShrink    *string `json:"Location,omitempty" xml:"Location,omitempty"`
	// example:
	//
	// latest
	ModelVersion *string `json:"ModelVersion,omitempty" xml:"ModelVersion,omitempty"`
	// example:
	//
	// 2025-02-12
	RunDate *string `json:"RunDate,omitempty" xml:"RunDate,omitempty"`
	// example:
	//
	// load
	SystemType *string `json:"SystemType,omitempty" xml:"SystemType,omitempty"`
	// example:
	//
	// Asia/Shanghai
	TimeZone *string `json:"TimeZone,omitempty" xml:"TimeZone,omitempty"`
}

func (s CreatePowerForecastJobShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePowerForecastJobShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreatePowerForecastJobShrinkRequest) GetBusinessKey() *string {
	return s.BusinessKey
}

func (s *CreatePowerForecastJobShrinkRequest) GetDataMode() *string {
	return s.DataMode
}

func (s *CreatePowerForecastJobShrinkRequest) GetDeviceType() *string {
	return s.DeviceType
}

func (s *CreatePowerForecastJobShrinkRequest) GetDuration() *int32 {
	return s.Duration
}

func (s *CreatePowerForecastJobShrinkRequest) GetForecastHorizon() *string {
	return s.ForecastHorizon
}

func (s *CreatePowerForecastJobShrinkRequest) GetFreq() *string {
	return s.Freq
}

func (s *CreatePowerForecastJobShrinkRequest) GetHistoryDataShrink() *string {
	return s.HistoryDataShrink
}

func (s *CreatePowerForecastJobShrinkRequest) GetLocationShrink() *string {
	return s.LocationShrink
}

func (s *CreatePowerForecastJobShrinkRequest) GetModelVersion() *string {
	return s.ModelVersion
}

func (s *CreatePowerForecastJobShrinkRequest) GetRunDate() *string {
	return s.RunDate
}

func (s *CreatePowerForecastJobShrinkRequest) GetSystemType() *string {
	return s.SystemType
}

func (s *CreatePowerForecastJobShrinkRequest) GetTimeZone() *string {
	return s.TimeZone
}

func (s *CreatePowerForecastJobShrinkRequest) SetBusinessKey(v string) *CreatePowerForecastJobShrinkRequest {
	s.BusinessKey = &v
	return s
}

func (s *CreatePowerForecastJobShrinkRequest) SetDataMode(v string) *CreatePowerForecastJobShrinkRequest {
	s.DataMode = &v
	return s
}

func (s *CreatePowerForecastJobShrinkRequest) SetDeviceType(v string) *CreatePowerForecastJobShrinkRequest {
	s.DeviceType = &v
	return s
}

func (s *CreatePowerForecastJobShrinkRequest) SetDuration(v int32) *CreatePowerForecastJobShrinkRequest {
	s.Duration = &v
	return s
}

func (s *CreatePowerForecastJobShrinkRequest) SetForecastHorizon(v string) *CreatePowerForecastJobShrinkRequest {
	s.ForecastHorizon = &v
	return s
}

func (s *CreatePowerForecastJobShrinkRequest) SetFreq(v string) *CreatePowerForecastJobShrinkRequest {
	s.Freq = &v
	return s
}

func (s *CreatePowerForecastJobShrinkRequest) SetHistoryDataShrink(v string) *CreatePowerForecastJobShrinkRequest {
	s.HistoryDataShrink = &v
	return s
}

func (s *CreatePowerForecastJobShrinkRequest) SetLocationShrink(v string) *CreatePowerForecastJobShrinkRequest {
	s.LocationShrink = &v
	return s
}

func (s *CreatePowerForecastJobShrinkRequest) SetModelVersion(v string) *CreatePowerForecastJobShrinkRequest {
	s.ModelVersion = &v
	return s
}

func (s *CreatePowerForecastJobShrinkRequest) SetRunDate(v string) *CreatePowerForecastJobShrinkRequest {
	s.RunDate = &v
	return s
}

func (s *CreatePowerForecastJobShrinkRequest) SetSystemType(v string) *CreatePowerForecastJobShrinkRequest {
	s.SystemType = &v
	return s
}

func (s *CreatePowerForecastJobShrinkRequest) SetTimeZone(v string) *CreatePowerForecastJobShrinkRequest {
	s.TimeZone = &v
	return s
}

func (s *CreatePowerForecastJobShrinkRequest) Validate() error {
	return dara.Validate(s)
}
