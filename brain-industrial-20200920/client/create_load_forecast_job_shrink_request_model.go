// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLoadForecastJobShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBusinessKey(v string) *CreateLoadForecastJobShrinkRequest
	GetBusinessKey() *string
	SetDataMode(v string) *CreateLoadForecastJobShrinkRequest
	GetDataMode() *string
	SetDeviceType(v string) *CreateLoadForecastJobShrinkRequest
	GetDeviceType() *string
	SetDuration(v int32) *CreateLoadForecastJobShrinkRequest
	GetDuration() *int32
	SetForecastHorizon(v string) *CreateLoadForecastJobShrinkRequest
	GetForecastHorizon() *string
	SetFreq(v string) *CreateLoadForecastJobShrinkRequest
	GetFreq() *string
	SetHistoryDataShrink(v string) *CreateLoadForecastJobShrinkRequest
	GetHistoryDataShrink() *string
	SetModelVersion(v string) *CreateLoadForecastJobShrinkRequest
	GetModelVersion() *string
	SetRunDate(v string) *CreateLoadForecastJobShrinkRequest
	GetRunDate() *string
	SetSystemType(v string) *CreateLoadForecastJobShrinkRequest
	GetSystemType() *string
	SetTimeZone(v string) *CreateLoadForecastJobShrinkRequest
	GetTimeZone() *string
}

type CreateLoadForecastJobShrinkRequest struct {
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
	// example:
	//
	// latest
	ModelVersion *string `json:"ModelVersion,omitempty" xml:"ModelVersion,omitempty"`
	// example:
	//
	// 2025-12-12
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

func (s CreateLoadForecastJobShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateLoadForecastJobShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateLoadForecastJobShrinkRequest) GetBusinessKey() *string {
	return s.BusinessKey
}

func (s *CreateLoadForecastJobShrinkRequest) GetDataMode() *string {
	return s.DataMode
}

func (s *CreateLoadForecastJobShrinkRequest) GetDeviceType() *string {
	return s.DeviceType
}

func (s *CreateLoadForecastJobShrinkRequest) GetDuration() *int32 {
	return s.Duration
}

func (s *CreateLoadForecastJobShrinkRequest) GetForecastHorizon() *string {
	return s.ForecastHorizon
}

func (s *CreateLoadForecastJobShrinkRequest) GetFreq() *string {
	return s.Freq
}

func (s *CreateLoadForecastJobShrinkRequest) GetHistoryDataShrink() *string {
	return s.HistoryDataShrink
}

func (s *CreateLoadForecastJobShrinkRequest) GetModelVersion() *string {
	return s.ModelVersion
}

func (s *CreateLoadForecastJobShrinkRequest) GetRunDate() *string {
	return s.RunDate
}

func (s *CreateLoadForecastJobShrinkRequest) GetSystemType() *string {
	return s.SystemType
}

func (s *CreateLoadForecastJobShrinkRequest) GetTimeZone() *string {
	return s.TimeZone
}

func (s *CreateLoadForecastJobShrinkRequest) SetBusinessKey(v string) *CreateLoadForecastJobShrinkRequest {
	s.BusinessKey = &v
	return s
}

func (s *CreateLoadForecastJobShrinkRequest) SetDataMode(v string) *CreateLoadForecastJobShrinkRequest {
	s.DataMode = &v
	return s
}

func (s *CreateLoadForecastJobShrinkRequest) SetDeviceType(v string) *CreateLoadForecastJobShrinkRequest {
	s.DeviceType = &v
	return s
}

func (s *CreateLoadForecastJobShrinkRequest) SetDuration(v int32) *CreateLoadForecastJobShrinkRequest {
	s.Duration = &v
	return s
}

func (s *CreateLoadForecastJobShrinkRequest) SetForecastHorizon(v string) *CreateLoadForecastJobShrinkRequest {
	s.ForecastHorizon = &v
	return s
}

func (s *CreateLoadForecastJobShrinkRequest) SetFreq(v string) *CreateLoadForecastJobShrinkRequest {
	s.Freq = &v
	return s
}

func (s *CreateLoadForecastJobShrinkRequest) SetHistoryDataShrink(v string) *CreateLoadForecastJobShrinkRequest {
	s.HistoryDataShrink = &v
	return s
}

func (s *CreateLoadForecastJobShrinkRequest) SetModelVersion(v string) *CreateLoadForecastJobShrinkRequest {
	s.ModelVersion = &v
	return s
}

func (s *CreateLoadForecastJobShrinkRequest) SetRunDate(v string) *CreateLoadForecastJobShrinkRequest {
	s.RunDate = &v
	return s
}

func (s *CreateLoadForecastJobShrinkRequest) SetSystemType(v string) *CreateLoadForecastJobShrinkRequest {
	s.SystemType = &v
	return s
}

func (s *CreateLoadForecastJobShrinkRequest) SetTimeZone(v string) *CreateLoadForecastJobShrinkRequest {
	s.TimeZone = &v
	return s
}

func (s *CreateLoadForecastJobShrinkRequest) Validate() error {
	return dara.Validate(s)
}
