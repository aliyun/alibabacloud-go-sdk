// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLoadForecastJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBusinessKey(v string) *CreateLoadForecastJobRequest
	GetBusinessKey() *string
	SetDataMode(v string) *CreateLoadForecastJobRequest
	GetDataMode() *string
	SetDeviceType(v string) *CreateLoadForecastJobRequest
	GetDeviceType() *string
	SetDuration(v int32) *CreateLoadForecastJobRequest
	GetDuration() *int32
	SetForecastHorizon(v string) *CreateLoadForecastJobRequest
	GetForecastHorizon() *string
	SetFreq(v string) *CreateLoadForecastJobRequest
	GetFreq() *string
	SetHistoryData(v []*CreateLoadForecastJobRequestHistoryData) *CreateLoadForecastJobRequest
	GetHistoryData() []*CreateLoadForecastJobRequestHistoryData
	SetModelVersion(v string) *CreateLoadForecastJobRequest
	GetModelVersion() *string
	SetRunDate(v string) *CreateLoadForecastJobRequest
	GetRunDate() *string
	SetSystemType(v string) *CreateLoadForecastJobRequest
	GetSystemType() *string
	SetTimeZone(v string) *CreateLoadForecastJobRequest
	GetTimeZone() *string
}

type CreateLoadForecastJobRequest struct {
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
	Freq        *string                                    `json:"Freq,omitempty" xml:"Freq,omitempty"`
	HistoryData []*CreateLoadForecastJobRequestHistoryData `json:"HistoryData,omitempty" xml:"HistoryData,omitempty" type:"Repeated"`
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

func (s CreateLoadForecastJobRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateLoadForecastJobRequest) GoString() string {
	return s.String()
}

func (s *CreateLoadForecastJobRequest) GetBusinessKey() *string {
	return s.BusinessKey
}

func (s *CreateLoadForecastJobRequest) GetDataMode() *string {
	return s.DataMode
}

func (s *CreateLoadForecastJobRequest) GetDeviceType() *string {
	return s.DeviceType
}

func (s *CreateLoadForecastJobRequest) GetDuration() *int32 {
	return s.Duration
}

func (s *CreateLoadForecastJobRequest) GetForecastHorizon() *string {
	return s.ForecastHorizon
}

func (s *CreateLoadForecastJobRequest) GetFreq() *string {
	return s.Freq
}

func (s *CreateLoadForecastJobRequest) GetHistoryData() []*CreateLoadForecastJobRequestHistoryData {
	return s.HistoryData
}

func (s *CreateLoadForecastJobRequest) GetModelVersion() *string {
	return s.ModelVersion
}

func (s *CreateLoadForecastJobRequest) GetRunDate() *string {
	return s.RunDate
}

func (s *CreateLoadForecastJobRequest) GetSystemType() *string {
	return s.SystemType
}

func (s *CreateLoadForecastJobRequest) GetTimeZone() *string {
	return s.TimeZone
}

func (s *CreateLoadForecastJobRequest) SetBusinessKey(v string) *CreateLoadForecastJobRequest {
	s.BusinessKey = &v
	return s
}

func (s *CreateLoadForecastJobRequest) SetDataMode(v string) *CreateLoadForecastJobRequest {
	s.DataMode = &v
	return s
}

func (s *CreateLoadForecastJobRequest) SetDeviceType(v string) *CreateLoadForecastJobRequest {
	s.DeviceType = &v
	return s
}

func (s *CreateLoadForecastJobRequest) SetDuration(v int32) *CreateLoadForecastJobRequest {
	s.Duration = &v
	return s
}

func (s *CreateLoadForecastJobRequest) SetForecastHorizon(v string) *CreateLoadForecastJobRequest {
	s.ForecastHorizon = &v
	return s
}

func (s *CreateLoadForecastJobRequest) SetFreq(v string) *CreateLoadForecastJobRequest {
	s.Freq = &v
	return s
}

func (s *CreateLoadForecastJobRequest) SetHistoryData(v []*CreateLoadForecastJobRequestHistoryData) *CreateLoadForecastJobRequest {
	s.HistoryData = v
	return s
}

func (s *CreateLoadForecastJobRequest) SetModelVersion(v string) *CreateLoadForecastJobRequest {
	s.ModelVersion = &v
	return s
}

func (s *CreateLoadForecastJobRequest) SetRunDate(v string) *CreateLoadForecastJobRequest {
	s.RunDate = &v
	return s
}

func (s *CreateLoadForecastJobRequest) SetSystemType(v string) *CreateLoadForecastJobRequest {
	s.SystemType = &v
	return s
}

func (s *CreateLoadForecastJobRequest) SetTimeZone(v string) *CreateLoadForecastJobRequest {
	s.TimeZone = &v
	return s
}

func (s *CreateLoadForecastJobRequest) Validate() error {
	if s.HistoryData != nil {
		for _, item := range s.HistoryData {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateLoadForecastJobRequestHistoryData struct {
	// example:
	//
	// 2025-12-12 00:00:00
	RunTime *string `json:"RunTime,omitempty" xml:"RunTime,omitempty"`
	// example:
	//
	// 1.0
	Value *float64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateLoadForecastJobRequestHistoryData) String() string {
	return dara.Prettify(s)
}

func (s CreateLoadForecastJobRequestHistoryData) GoString() string {
	return s.String()
}

func (s *CreateLoadForecastJobRequestHistoryData) GetRunTime() *string {
	return s.RunTime
}

func (s *CreateLoadForecastJobRequestHistoryData) GetValue() *float64 {
	return s.Value
}

func (s *CreateLoadForecastJobRequestHistoryData) SetRunTime(v string) *CreateLoadForecastJobRequestHistoryData {
	s.RunTime = &v
	return s
}

func (s *CreateLoadForecastJobRequestHistoryData) SetValue(v float64) *CreateLoadForecastJobRequestHistoryData {
	s.Value = &v
	return s
}

func (s *CreateLoadForecastJobRequestHistoryData) Validate() error {
	return dara.Validate(s)
}
