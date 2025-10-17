// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLoadForecastByFileUrlJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBusinessKey(v string) *CreateLoadForecastByFileUrlJobRequest
	GetBusinessKey() *string
	SetDataMode(v string) *CreateLoadForecastByFileUrlJobRequest
	GetDataMode() *string
	SetDeviceType(v string) *CreateLoadForecastByFileUrlJobRequest
	GetDeviceType() *string
	SetDuration(v int32) *CreateLoadForecastByFileUrlJobRequest
	GetDuration() *int32
	SetForecastHorizon(v string) *CreateLoadForecastByFileUrlJobRequest
	GetForecastHorizon() *string
	SetFreq(v string) *CreateLoadForecastByFileUrlJobRequest
	GetFreq() *string
	SetHistoryUrl(v string) *CreateLoadForecastByFileUrlJobRequest
	GetHistoryUrl() *string
	SetModelVersion(v string) *CreateLoadForecastByFileUrlJobRequest
	GetModelVersion() *string
	SetRunDate(v string) *CreateLoadForecastByFileUrlJobRequest
	GetRunDate() *string
	SetSystemType(v string) *CreateLoadForecastByFileUrlJobRequest
	GetSystemType() *string
	SetTimeColumn(v string) *CreateLoadForecastByFileUrlJobRequest
	GetTimeColumn() *string
	SetTimeZone(v string) *CreateLoadForecastByFileUrlJobRequest
	GetTimeZone() *string
	SetValueColumn(v string) *CreateLoadForecastByFileUrlJobRequest
	GetValueColumn() *string
}

type CreateLoadForecastByFileUrlJobRequest struct {
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
	Freq *string `json:"Freq,omitempty" xml:"Freq,omitempty"`
	// example:
	//
	// https://bucket.oss-cn-hangzhou.aliyuncs.com/dir/target_file.csv
	HistoryUrl *string `json:"HistoryUrl,omitempty" xml:"HistoryUrl,omitempty"`
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
	// load
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

func (s CreateLoadForecastByFileUrlJobRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateLoadForecastByFileUrlJobRequest) GoString() string {
	return s.String()
}

func (s *CreateLoadForecastByFileUrlJobRequest) GetBusinessKey() *string {
	return s.BusinessKey
}

func (s *CreateLoadForecastByFileUrlJobRequest) GetDataMode() *string {
	return s.DataMode
}

func (s *CreateLoadForecastByFileUrlJobRequest) GetDeviceType() *string {
	return s.DeviceType
}

func (s *CreateLoadForecastByFileUrlJobRequest) GetDuration() *int32 {
	return s.Duration
}

func (s *CreateLoadForecastByFileUrlJobRequest) GetForecastHorizon() *string {
	return s.ForecastHorizon
}

func (s *CreateLoadForecastByFileUrlJobRequest) GetFreq() *string {
	return s.Freq
}

func (s *CreateLoadForecastByFileUrlJobRequest) GetHistoryUrl() *string {
	return s.HistoryUrl
}

func (s *CreateLoadForecastByFileUrlJobRequest) GetModelVersion() *string {
	return s.ModelVersion
}

func (s *CreateLoadForecastByFileUrlJobRequest) GetRunDate() *string {
	return s.RunDate
}

func (s *CreateLoadForecastByFileUrlJobRequest) GetSystemType() *string {
	return s.SystemType
}

func (s *CreateLoadForecastByFileUrlJobRequest) GetTimeColumn() *string {
	return s.TimeColumn
}

func (s *CreateLoadForecastByFileUrlJobRequest) GetTimeZone() *string {
	return s.TimeZone
}

func (s *CreateLoadForecastByFileUrlJobRequest) GetValueColumn() *string {
	return s.ValueColumn
}

func (s *CreateLoadForecastByFileUrlJobRequest) SetBusinessKey(v string) *CreateLoadForecastByFileUrlJobRequest {
	s.BusinessKey = &v
	return s
}

func (s *CreateLoadForecastByFileUrlJobRequest) SetDataMode(v string) *CreateLoadForecastByFileUrlJobRequest {
	s.DataMode = &v
	return s
}

func (s *CreateLoadForecastByFileUrlJobRequest) SetDeviceType(v string) *CreateLoadForecastByFileUrlJobRequest {
	s.DeviceType = &v
	return s
}

func (s *CreateLoadForecastByFileUrlJobRequest) SetDuration(v int32) *CreateLoadForecastByFileUrlJobRequest {
	s.Duration = &v
	return s
}

func (s *CreateLoadForecastByFileUrlJobRequest) SetForecastHorizon(v string) *CreateLoadForecastByFileUrlJobRequest {
	s.ForecastHorizon = &v
	return s
}

func (s *CreateLoadForecastByFileUrlJobRequest) SetFreq(v string) *CreateLoadForecastByFileUrlJobRequest {
	s.Freq = &v
	return s
}

func (s *CreateLoadForecastByFileUrlJobRequest) SetHistoryUrl(v string) *CreateLoadForecastByFileUrlJobRequest {
	s.HistoryUrl = &v
	return s
}

func (s *CreateLoadForecastByFileUrlJobRequest) SetModelVersion(v string) *CreateLoadForecastByFileUrlJobRequest {
	s.ModelVersion = &v
	return s
}

func (s *CreateLoadForecastByFileUrlJobRequest) SetRunDate(v string) *CreateLoadForecastByFileUrlJobRequest {
	s.RunDate = &v
	return s
}

func (s *CreateLoadForecastByFileUrlJobRequest) SetSystemType(v string) *CreateLoadForecastByFileUrlJobRequest {
	s.SystemType = &v
	return s
}

func (s *CreateLoadForecastByFileUrlJobRequest) SetTimeColumn(v string) *CreateLoadForecastByFileUrlJobRequest {
	s.TimeColumn = &v
	return s
}

func (s *CreateLoadForecastByFileUrlJobRequest) SetTimeZone(v string) *CreateLoadForecastByFileUrlJobRequest {
	s.TimeZone = &v
	return s
}

func (s *CreateLoadForecastByFileUrlJobRequest) SetValueColumn(v string) *CreateLoadForecastByFileUrlJobRequest {
	s.ValueColumn = &v
	return s
}

func (s *CreateLoadForecastByFileUrlJobRequest) Validate() error {
	return dara.Validate(s)
}
