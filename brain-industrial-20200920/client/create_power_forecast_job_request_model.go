// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePowerForecastJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBusinessKey(v string) *CreatePowerForecastJobRequest
	GetBusinessKey() *string
	SetDataMode(v string) *CreatePowerForecastJobRequest
	GetDataMode() *string
	SetDeviceType(v string) *CreatePowerForecastJobRequest
	GetDeviceType() *string
	SetDuration(v int32) *CreatePowerForecastJobRequest
	GetDuration() *int32
	SetForecastHorizon(v string) *CreatePowerForecastJobRequest
	GetForecastHorizon() *string
	SetFreq(v string) *CreatePowerForecastJobRequest
	GetFreq() *string
	SetHistoryData(v []*CreatePowerForecastJobRequestHistoryData) *CreatePowerForecastJobRequest
	GetHistoryData() []*CreatePowerForecastJobRequestHistoryData
	SetLocation(v *CreatePowerForecastJobRequestLocation) *CreatePowerForecastJobRequest
	GetLocation() *CreatePowerForecastJobRequestLocation
	SetModelVersion(v string) *CreatePowerForecastJobRequest
	GetModelVersion() *string
	SetRunDate(v string) *CreatePowerForecastJobRequest
	GetRunDate() *string
	SetSystemType(v string) *CreatePowerForecastJobRequest
	GetSystemType() *string
	SetTimeZone(v string) *CreatePowerForecastJobRequest
	GetTimeZone() *string
}

type CreatePowerForecastJobRequest struct {
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
	Freq        *string                                     `json:"Freq,omitempty" xml:"Freq,omitempty"`
	HistoryData []*CreatePowerForecastJobRequestHistoryData `json:"HistoryData,omitempty" xml:"HistoryData,omitempty" type:"Repeated"`
	Location    *CreatePowerForecastJobRequestLocation      `json:"Location,omitempty" xml:"Location,omitempty" type:"Struct"`
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

func (s CreatePowerForecastJobRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePowerForecastJobRequest) GoString() string {
	return s.String()
}

func (s *CreatePowerForecastJobRequest) GetBusinessKey() *string {
	return s.BusinessKey
}

func (s *CreatePowerForecastJobRequest) GetDataMode() *string {
	return s.DataMode
}

func (s *CreatePowerForecastJobRequest) GetDeviceType() *string {
	return s.DeviceType
}

func (s *CreatePowerForecastJobRequest) GetDuration() *int32 {
	return s.Duration
}

func (s *CreatePowerForecastJobRequest) GetForecastHorizon() *string {
	return s.ForecastHorizon
}

func (s *CreatePowerForecastJobRequest) GetFreq() *string {
	return s.Freq
}

func (s *CreatePowerForecastJobRequest) GetHistoryData() []*CreatePowerForecastJobRequestHistoryData {
	return s.HistoryData
}

func (s *CreatePowerForecastJobRequest) GetLocation() *CreatePowerForecastJobRequestLocation {
	return s.Location
}

func (s *CreatePowerForecastJobRequest) GetModelVersion() *string {
	return s.ModelVersion
}

func (s *CreatePowerForecastJobRequest) GetRunDate() *string {
	return s.RunDate
}

func (s *CreatePowerForecastJobRequest) GetSystemType() *string {
	return s.SystemType
}

func (s *CreatePowerForecastJobRequest) GetTimeZone() *string {
	return s.TimeZone
}

func (s *CreatePowerForecastJobRequest) SetBusinessKey(v string) *CreatePowerForecastJobRequest {
	s.BusinessKey = &v
	return s
}

func (s *CreatePowerForecastJobRequest) SetDataMode(v string) *CreatePowerForecastJobRequest {
	s.DataMode = &v
	return s
}

func (s *CreatePowerForecastJobRequest) SetDeviceType(v string) *CreatePowerForecastJobRequest {
	s.DeviceType = &v
	return s
}

func (s *CreatePowerForecastJobRequest) SetDuration(v int32) *CreatePowerForecastJobRequest {
	s.Duration = &v
	return s
}

func (s *CreatePowerForecastJobRequest) SetForecastHorizon(v string) *CreatePowerForecastJobRequest {
	s.ForecastHorizon = &v
	return s
}

func (s *CreatePowerForecastJobRequest) SetFreq(v string) *CreatePowerForecastJobRequest {
	s.Freq = &v
	return s
}

func (s *CreatePowerForecastJobRequest) SetHistoryData(v []*CreatePowerForecastJobRequestHistoryData) *CreatePowerForecastJobRequest {
	s.HistoryData = v
	return s
}

func (s *CreatePowerForecastJobRequest) SetLocation(v *CreatePowerForecastJobRequestLocation) *CreatePowerForecastJobRequest {
	s.Location = v
	return s
}

func (s *CreatePowerForecastJobRequest) SetModelVersion(v string) *CreatePowerForecastJobRequest {
	s.ModelVersion = &v
	return s
}

func (s *CreatePowerForecastJobRequest) SetRunDate(v string) *CreatePowerForecastJobRequest {
	s.RunDate = &v
	return s
}

func (s *CreatePowerForecastJobRequest) SetSystemType(v string) *CreatePowerForecastJobRequest {
	s.SystemType = &v
	return s
}

func (s *CreatePowerForecastJobRequest) SetTimeZone(v string) *CreatePowerForecastJobRequest {
	s.TimeZone = &v
	return s
}

func (s *CreatePowerForecastJobRequest) Validate() error {
	if s.HistoryData != nil {
		for _, item := range s.HistoryData {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Location != nil {
		if err := s.Location.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreatePowerForecastJobRequestHistoryData struct {
	// example:
	//
	// 2025-02-12 00:00:00
	RunTime *string `json:"RunTime,omitempty" xml:"RunTime,omitempty"`
	// example:
	//
	// 1.0
	Value *float64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreatePowerForecastJobRequestHistoryData) String() string {
	return dara.Prettify(s)
}

func (s CreatePowerForecastJobRequestHistoryData) GoString() string {
	return s.String()
}

func (s *CreatePowerForecastJobRequestHistoryData) GetRunTime() *string {
	return s.RunTime
}

func (s *CreatePowerForecastJobRequestHistoryData) GetValue() *float64 {
	return s.Value
}

func (s *CreatePowerForecastJobRequestHistoryData) SetRunTime(v string) *CreatePowerForecastJobRequestHistoryData {
	s.RunTime = &v
	return s
}

func (s *CreatePowerForecastJobRequestHistoryData) SetValue(v float64) *CreatePowerForecastJobRequestHistoryData {
	s.Value = &v
	return s
}

func (s *CreatePowerForecastJobRequestHistoryData) Validate() error {
	return dara.Validate(s)
}

type CreatePowerForecastJobRequestLocation struct {
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

func (s CreatePowerForecastJobRequestLocation) String() string {
	return dara.Prettify(s)
}

func (s CreatePowerForecastJobRequestLocation) GoString() string {
	return s.String()
}

func (s *CreatePowerForecastJobRequestLocation) GetAltitude() *float64 {
	return s.Altitude
}

func (s *CreatePowerForecastJobRequestLocation) GetLatitude() *float64 {
	return s.Latitude
}

func (s *CreatePowerForecastJobRequestLocation) GetLongitude() *float64 {
	return s.Longitude
}

func (s *CreatePowerForecastJobRequestLocation) SetAltitude(v float64) *CreatePowerForecastJobRequestLocation {
	s.Altitude = &v
	return s
}

func (s *CreatePowerForecastJobRequestLocation) SetLatitude(v float64) *CreatePowerForecastJobRequestLocation {
	s.Latitude = &v
	return s
}

func (s *CreatePowerForecastJobRequestLocation) SetLongitude(v float64) *CreatePowerForecastJobRequestLocation {
	s.Longitude = &v
	return s
}

func (s *CreatePowerForecastJobRequestLocation) Validate() error {
	return dara.Validate(s)
}
