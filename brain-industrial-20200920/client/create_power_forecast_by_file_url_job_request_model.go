// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePowerForecastByFileUrlJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBusinessKey(v string) *CreatePowerForecastByFileUrlJobRequest
	GetBusinessKey() *string
	SetDataMode(v string) *CreatePowerForecastByFileUrlJobRequest
	GetDataMode() *string
	SetDeviceType(v string) *CreatePowerForecastByFileUrlJobRequest
	GetDeviceType() *string
	SetDuration(v int32) *CreatePowerForecastByFileUrlJobRequest
	GetDuration() *int32
	SetForecastHorizon(v string) *CreatePowerForecastByFileUrlJobRequest
	GetForecastHorizon() *string
	SetFreq(v string) *CreatePowerForecastByFileUrlJobRequest
	GetFreq() *string
	SetHistoryUrl(v string) *CreatePowerForecastByFileUrlJobRequest
	GetHistoryUrl() *string
	SetLocation(v *CreatePowerForecastByFileUrlJobRequestLocation) *CreatePowerForecastByFileUrlJobRequest
	GetLocation() *CreatePowerForecastByFileUrlJobRequestLocation
	SetModelVersion(v string) *CreatePowerForecastByFileUrlJobRequest
	GetModelVersion() *string
	SetRunDate(v string) *CreatePowerForecastByFileUrlJobRequest
	GetRunDate() *string
	SetSystemType(v string) *CreatePowerForecastByFileUrlJobRequest
	GetSystemType() *string
	SetTimeColumn(v string) *CreatePowerForecastByFileUrlJobRequest
	GetTimeColumn() *string
	SetTimeZone(v string) *CreatePowerForecastByFileUrlJobRequest
	GetTimeZone() *string
	SetValueColumn(v string) *CreatePowerForecastByFileUrlJobRequest
	GetValueColumn() *string
}

type CreatePowerForecastByFileUrlJobRequest struct {
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
	HistoryUrl *string                                         `json:"HistoryUrl,omitempty" xml:"HistoryUrl,omitempty"`
	Location   *CreatePowerForecastByFileUrlJobRequestLocation `json:"Location,omitempty" xml:"Location,omitempty" type:"Struct"`
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

func (s CreatePowerForecastByFileUrlJobRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePowerForecastByFileUrlJobRequest) GoString() string {
	return s.String()
}

func (s *CreatePowerForecastByFileUrlJobRequest) GetBusinessKey() *string {
	return s.BusinessKey
}

func (s *CreatePowerForecastByFileUrlJobRequest) GetDataMode() *string {
	return s.DataMode
}

func (s *CreatePowerForecastByFileUrlJobRequest) GetDeviceType() *string {
	return s.DeviceType
}

func (s *CreatePowerForecastByFileUrlJobRequest) GetDuration() *int32 {
	return s.Duration
}

func (s *CreatePowerForecastByFileUrlJobRequest) GetForecastHorizon() *string {
	return s.ForecastHorizon
}

func (s *CreatePowerForecastByFileUrlJobRequest) GetFreq() *string {
	return s.Freq
}

func (s *CreatePowerForecastByFileUrlJobRequest) GetHistoryUrl() *string {
	return s.HistoryUrl
}

func (s *CreatePowerForecastByFileUrlJobRequest) GetLocation() *CreatePowerForecastByFileUrlJobRequestLocation {
	return s.Location
}

func (s *CreatePowerForecastByFileUrlJobRequest) GetModelVersion() *string {
	return s.ModelVersion
}

func (s *CreatePowerForecastByFileUrlJobRequest) GetRunDate() *string {
	return s.RunDate
}

func (s *CreatePowerForecastByFileUrlJobRequest) GetSystemType() *string {
	return s.SystemType
}

func (s *CreatePowerForecastByFileUrlJobRequest) GetTimeColumn() *string {
	return s.TimeColumn
}

func (s *CreatePowerForecastByFileUrlJobRequest) GetTimeZone() *string {
	return s.TimeZone
}

func (s *CreatePowerForecastByFileUrlJobRequest) GetValueColumn() *string {
	return s.ValueColumn
}

func (s *CreatePowerForecastByFileUrlJobRequest) SetBusinessKey(v string) *CreatePowerForecastByFileUrlJobRequest {
	s.BusinessKey = &v
	return s
}

func (s *CreatePowerForecastByFileUrlJobRequest) SetDataMode(v string) *CreatePowerForecastByFileUrlJobRequest {
	s.DataMode = &v
	return s
}

func (s *CreatePowerForecastByFileUrlJobRequest) SetDeviceType(v string) *CreatePowerForecastByFileUrlJobRequest {
	s.DeviceType = &v
	return s
}

func (s *CreatePowerForecastByFileUrlJobRequest) SetDuration(v int32) *CreatePowerForecastByFileUrlJobRequest {
	s.Duration = &v
	return s
}

func (s *CreatePowerForecastByFileUrlJobRequest) SetForecastHorizon(v string) *CreatePowerForecastByFileUrlJobRequest {
	s.ForecastHorizon = &v
	return s
}

func (s *CreatePowerForecastByFileUrlJobRequest) SetFreq(v string) *CreatePowerForecastByFileUrlJobRequest {
	s.Freq = &v
	return s
}

func (s *CreatePowerForecastByFileUrlJobRequest) SetHistoryUrl(v string) *CreatePowerForecastByFileUrlJobRequest {
	s.HistoryUrl = &v
	return s
}

func (s *CreatePowerForecastByFileUrlJobRequest) SetLocation(v *CreatePowerForecastByFileUrlJobRequestLocation) *CreatePowerForecastByFileUrlJobRequest {
	s.Location = v
	return s
}

func (s *CreatePowerForecastByFileUrlJobRequest) SetModelVersion(v string) *CreatePowerForecastByFileUrlJobRequest {
	s.ModelVersion = &v
	return s
}

func (s *CreatePowerForecastByFileUrlJobRequest) SetRunDate(v string) *CreatePowerForecastByFileUrlJobRequest {
	s.RunDate = &v
	return s
}

func (s *CreatePowerForecastByFileUrlJobRequest) SetSystemType(v string) *CreatePowerForecastByFileUrlJobRequest {
	s.SystemType = &v
	return s
}

func (s *CreatePowerForecastByFileUrlJobRequest) SetTimeColumn(v string) *CreatePowerForecastByFileUrlJobRequest {
	s.TimeColumn = &v
	return s
}

func (s *CreatePowerForecastByFileUrlJobRequest) SetTimeZone(v string) *CreatePowerForecastByFileUrlJobRequest {
	s.TimeZone = &v
	return s
}

func (s *CreatePowerForecastByFileUrlJobRequest) SetValueColumn(v string) *CreatePowerForecastByFileUrlJobRequest {
	s.ValueColumn = &v
	return s
}

func (s *CreatePowerForecastByFileUrlJobRequest) Validate() error {
	if s.Location != nil {
		if err := s.Location.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreatePowerForecastByFileUrlJobRequestLocation struct {
	// example:
	//
	// 10.123
	Altitude *float64 `json:"Altitude,omitempty" xml:"Altitude,omitempty"`
	// example:
	//
	// 40.027
	Latitude *float64 `json:"Latitude,omitempty" xml:"Latitude,omitempty"`
	// example:
	//
	// 120.042
	Longitude *float64 `json:"Longitude,omitempty" xml:"Longitude,omitempty"`
}

func (s CreatePowerForecastByFileUrlJobRequestLocation) String() string {
	return dara.Prettify(s)
}

func (s CreatePowerForecastByFileUrlJobRequestLocation) GoString() string {
	return s.String()
}

func (s *CreatePowerForecastByFileUrlJobRequestLocation) GetAltitude() *float64 {
	return s.Altitude
}

func (s *CreatePowerForecastByFileUrlJobRequestLocation) GetLatitude() *float64 {
	return s.Latitude
}

func (s *CreatePowerForecastByFileUrlJobRequestLocation) GetLongitude() *float64 {
	return s.Longitude
}

func (s *CreatePowerForecastByFileUrlJobRequestLocation) SetAltitude(v float64) *CreatePowerForecastByFileUrlJobRequestLocation {
	s.Altitude = &v
	return s
}

func (s *CreatePowerForecastByFileUrlJobRequestLocation) SetLatitude(v float64) *CreatePowerForecastByFileUrlJobRequestLocation {
	s.Latitude = &v
	return s
}

func (s *CreatePowerForecastByFileUrlJobRequestLocation) SetLongitude(v float64) *CreatePowerForecastByFileUrlJobRequestLocation {
	s.Longitude = &v
	return s
}

func (s *CreatePowerForecastByFileUrlJobRequestLocation) Validate() error {
	return dara.Validate(s)
}
