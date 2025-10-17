// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePowerForecastByFileUrlJobShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBusinessKey(v string) *CreatePowerForecastByFileUrlJobShrinkRequest
	GetBusinessKey() *string
	SetDataMode(v string) *CreatePowerForecastByFileUrlJobShrinkRequest
	GetDataMode() *string
	SetDeviceType(v string) *CreatePowerForecastByFileUrlJobShrinkRequest
	GetDeviceType() *string
	SetDuration(v int32) *CreatePowerForecastByFileUrlJobShrinkRequest
	GetDuration() *int32
	SetForecastHorizon(v string) *CreatePowerForecastByFileUrlJobShrinkRequest
	GetForecastHorizon() *string
	SetFreq(v string) *CreatePowerForecastByFileUrlJobShrinkRequest
	GetFreq() *string
	SetHistoryUrl(v string) *CreatePowerForecastByFileUrlJobShrinkRequest
	GetHistoryUrl() *string
	SetLocationShrink(v string) *CreatePowerForecastByFileUrlJobShrinkRequest
	GetLocationShrink() *string
	SetModelVersion(v string) *CreatePowerForecastByFileUrlJobShrinkRequest
	GetModelVersion() *string
	SetRunDate(v string) *CreatePowerForecastByFileUrlJobShrinkRequest
	GetRunDate() *string
	SetSystemType(v string) *CreatePowerForecastByFileUrlJobShrinkRequest
	GetSystemType() *string
	SetTimeColumn(v string) *CreatePowerForecastByFileUrlJobShrinkRequest
	GetTimeColumn() *string
	SetTimeZone(v string) *CreatePowerForecastByFileUrlJobShrinkRequest
	GetTimeZone() *string
	SetValueColumn(v string) *CreatePowerForecastByFileUrlJobShrinkRequest
	GetValueColumn() *string
}

type CreatePowerForecastByFileUrlJobShrinkRequest struct {
	// example:
	//
	// stationA
	BusinessKey *string `json:"BusinessKey,omitempty" xml:"BusinessKey,omitempty"`
	// example:
	//
	// FULL
	DataMode *string `json:"DataMode,omitempty" xml:"DataMode,omitempty"`
	// example:
	//
	// solarInverter
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
	Freq *string `json:"Freq,omitempty" xml:"Freq,omitempty"`
	// example:
	//
	// https://bucket.oss-cn-hangzhou.aliyuncs.com/dir/target_file.csv
	HistoryUrl     *string `json:"HistoryUrl,omitempty" xml:"HistoryUrl,omitempty"`
	LocationShrink *string `json:"Location,omitempty" xml:"Location,omitempty"`
	// example:
	//
	// latest
	ModelVersion *string `json:"ModelVersion,omitempty" xml:"ModelVersion,omitempty"`
	// example:
	//
	// 2025-01-01
	RunDate *string `json:"RunDate,omitempty" xml:"RunDate,omitempty"`
	// example:
	//
	// solar
	SystemType *string `json:"SystemType,omitempty" xml:"SystemType,omitempty"`
	// example:
	//
	// runTime
	TimeColumn *string `json:"TimeColumn,omitempty" xml:"TimeColumn,omitempty"`
	// example:
	//
	// Asia/Shanghai
	TimeZone *string `json:"TimeZone,omitempty" xml:"TimeZone,omitempty"`
	// example:
	//
	// value
	ValueColumn *string `json:"ValueColumn,omitempty" xml:"ValueColumn,omitempty"`
}

func (s CreatePowerForecastByFileUrlJobShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePowerForecastByFileUrlJobShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreatePowerForecastByFileUrlJobShrinkRequest) GetBusinessKey() *string {
	return s.BusinessKey
}

func (s *CreatePowerForecastByFileUrlJobShrinkRequest) GetDataMode() *string {
	return s.DataMode
}

func (s *CreatePowerForecastByFileUrlJobShrinkRequest) GetDeviceType() *string {
	return s.DeviceType
}

func (s *CreatePowerForecastByFileUrlJobShrinkRequest) GetDuration() *int32 {
	return s.Duration
}

func (s *CreatePowerForecastByFileUrlJobShrinkRequest) GetForecastHorizon() *string {
	return s.ForecastHorizon
}

func (s *CreatePowerForecastByFileUrlJobShrinkRequest) GetFreq() *string {
	return s.Freq
}

func (s *CreatePowerForecastByFileUrlJobShrinkRequest) GetHistoryUrl() *string {
	return s.HistoryUrl
}

func (s *CreatePowerForecastByFileUrlJobShrinkRequest) GetLocationShrink() *string {
	return s.LocationShrink
}

func (s *CreatePowerForecastByFileUrlJobShrinkRequest) GetModelVersion() *string {
	return s.ModelVersion
}

func (s *CreatePowerForecastByFileUrlJobShrinkRequest) GetRunDate() *string {
	return s.RunDate
}

func (s *CreatePowerForecastByFileUrlJobShrinkRequest) GetSystemType() *string {
	return s.SystemType
}

func (s *CreatePowerForecastByFileUrlJobShrinkRequest) GetTimeColumn() *string {
	return s.TimeColumn
}

func (s *CreatePowerForecastByFileUrlJobShrinkRequest) GetTimeZone() *string {
	return s.TimeZone
}

func (s *CreatePowerForecastByFileUrlJobShrinkRequest) GetValueColumn() *string {
	return s.ValueColumn
}

func (s *CreatePowerForecastByFileUrlJobShrinkRequest) SetBusinessKey(v string) *CreatePowerForecastByFileUrlJobShrinkRequest {
	s.BusinessKey = &v
	return s
}

func (s *CreatePowerForecastByFileUrlJobShrinkRequest) SetDataMode(v string) *CreatePowerForecastByFileUrlJobShrinkRequest {
	s.DataMode = &v
	return s
}

func (s *CreatePowerForecastByFileUrlJobShrinkRequest) SetDeviceType(v string) *CreatePowerForecastByFileUrlJobShrinkRequest {
	s.DeviceType = &v
	return s
}

func (s *CreatePowerForecastByFileUrlJobShrinkRequest) SetDuration(v int32) *CreatePowerForecastByFileUrlJobShrinkRequest {
	s.Duration = &v
	return s
}

func (s *CreatePowerForecastByFileUrlJobShrinkRequest) SetForecastHorizon(v string) *CreatePowerForecastByFileUrlJobShrinkRequest {
	s.ForecastHorizon = &v
	return s
}

func (s *CreatePowerForecastByFileUrlJobShrinkRequest) SetFreq(v string) *CreatePowerForecastByFileUrlJobShrinkRequest {
	s.Freq = &v
	return s
}

func (s *CreatePowerForecastByFileUrlJobShrinkRequest) SetHistoryUrl(v string) *CreatePowerForecastByFileUrlJobShrinkRequest {
	s.HistoryUrl = &v
	return s
}

func (s *CreatePowerForecastByFileUrlJobShrinkRequest) SetLocationShrink(v string) *CreatePowerForecastByFileUrlJobShrinkRequest {
	s.LocationShrink = &v
	return s
}

func (s *CreatePowerForecastByFileUrlJobShrinkRequest) SetModelVersion(v string) *CreatePowerForecastByFileUrlJobShrinkRequest {
	s.ModelVersion = &v
	return s
}

func (s *CreatePowerForecastByFileUrlJobShrinkRequest) SetRunDate(v string) *CreatePowerForecastByFileUrlJobShrinkRequest {
	s.RunDate = &v
	return s
}

func (s *CreatePowerForecastByFileUrlJobShrinkRequest) SetSystemType(v string) *CreatePowerForecastByFileUrlJobShrinkRequest {
	s.SystemType = &v
	return s
}

func (s *CreatePowerForecastByFileUrlJobShrinkRequest) SetTimeColumn(v string) *CreatePowerForecastByFileUrlJobShrinkRequest {
	s.TimeColumn = &v
	return s
}

func (s *CreatePowerForecastByFileUrlJobShrinkRequest) SetTimeZone(v string) *CreatePowerForecastByFileUrlJobShrinkRequest {
	s.TimeZone = &v
	return s
}

func (s *CreatePowerForecastByFileUrlJobShrinkRequest) SetValueColumn(v string) *CreatePowerForecastByFileUrlJobShrinkRequest {
	s.ValueColumn = &v
	return s
}

func (s *CreatePowerForecastByFileUrlJobShrinkRequest) Validate() error {
	return dara.Validate(s)
}
