// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetRunningPlanRequest interface {
	dara.Model
	String() string
	GoString() string
	SetControlType(v int32) *SetRunningPlanRequest
	GetControlType() *int32
	SetDateType(v int32) *SetRunningPlanRequest
	GetDateType() *int32
	SetEarliestStartupTime(v string) *SetRunningPlanRequest
	GetEarliestStartupTime() *string
	SetEndTime(v string) *SetRunningPlanRequest
	GetEndTime() *string
	SetFactoryId(v string) *SetRunningPlanRequest
	GetFactoryId() *string
	SetLatestShutdownTime(v string) *SetRunningPlanRequest
	GetLatestShutdownTime() *string
	SetMaxCarbonDioxide(v float64) *SetRunningPlanRequest
	GetMaxCarbonDioxide() *float64
	SetMaxTem(v float64) *SetRunningPlanRequest
	GetMaxTem() *float64
	SetMinTem(v float64) *SetRunningPlanRequest
	GetMinTem() *float64
	SetPKey(v string) *SetRunningPlanRequest
	GetPKey() *string
	SetSeasonMode(v int32) *SetRunningPlanRequest
	GetSeasonMode() *int32
	SetStartTime(v string) *SetRunningPlanRequest
	GetStartTime() *string
	SetStatisticsTime(v string) *SetRunningPlanRequest
	GetStatisticsTime() *string
	SetSystemId(v string) *SetRunningPlanRequest
	GetSystemId() *string
	SetWorkingEndTime(v string) *SetRunningPlanRequest
	GetWorkingEndTime() *string
	SetWorkingStartTime(v string) *SetRunningPlanRequest
	GetWorkingStartTime() *string
}

type SetRunningPlanRequest struct {
	// example:
	//
	// 0
	ControlType *int32 `json:"controlType,omitempty" xml:"controlType,omitempty"`
	// example:
	//
	// 0
	DateType *int32 `json:"dateType,omitempty" xml:"dateType,omitempty"`
	// example:
	//
	// 05:00:00
	EarliestStartupTime *string `json:"earliestStartupTime,omitempty" xml:"earliestStartupTime,omitempty"`
	// example:
	//
	// 2024-07-21
	EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ***
	FactoryId *string `json:"factoryId,omitempty" xml:"factoryId,omitempty"`
	// example:
	//
	// 05:30:00
	LatestShutdownTime *string `json:"latestShutdownTime,omitempty" xml:"latestShutdownTime,omitempty"`
	// example:
	//
	// 2.1
	MaxCarbonDioxide *float64 `json:"maxCarbonDioxide,omitempty" xml:"maxCarbonDioxide,omitempty"`
	// example:
	//
	// 3.1
	MaxTem *float64 `json:"maxTem,omitempty" xml:"maxTem,omitempty"`
	// example:
	//
	// 2.1
	MinTem *float64 `json:"minTem,omitempty" xml:"minTem,omitempty"`
	// example:
	//
	// ib
	PKey *string `json:"pKey,omitempty" xml:"pKey,omitempty"`
	// example:
	//
	// 0
	SeasonMode *int32 `json:"seasonMode,omitempty" xml:"seasonMode,omitempty"`
	// example:
	//
	// 2024-07-20
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// example:
	//
	// 2024-07-31
	StatisticsTime *string `json:"statisticsTime,omitempty" xml:"statisticsTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// system1
	SystemId *string `json:"systemId,omitempty" xml:"systemId,omitempty"`
	// example:
	//
	// 05:30:00
	WorkingEndTime *string `json:"workingEndTime,omitempty" xml:"workingEndTime,omitempty"`
	// example:
	//
	// 05:00:00
	WorkingStartTime *string `json:"workingStartTime,omitempty" xml:"workingStartTime,omitempty"`
}

func (s SetRunningPlanRequest) String() string {
	return dara.Prettify(s)
}

func (s SetRunningPlanRequest) GoString() string {
	return s.String()
}

func (s *SetRunningPlanRequest) GetControlType() *int32 {
	return s.ControlType
}

func (s *SetRunningPlanRequest) GetDateType() *int32 {
	return s.DateType
}

func (s *SetRunningPlanRequest) GetEarliestStartupTime() *string {
	return s.EarliestStartupTime
}

func (s *SetRunningPlanRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *SetRunningPlanRequest) GetFactoryId() *string {
	return s.FactoryId
}

func (s *SetRunningPlanRequest) GetLatestShutdownTime() *string {
	return s.LatestShutdownTime
}

func (s *SetRunningPlanRequest) GetMaxCarbonDioxide() *float64 {
	return s.MaxCarbonDioxide
}

func (s *SetRunningPlanRequest) GetMaxTem() *float64 {
	return s.MaxTem
}

func (s *SetRunningPlanRequest) GetMinTem() *float64 {
	return s.MinTem
}

func (s *SetRunningPlanRequest) GetPKey() *string {
	return s.PKey
}

func (s *SetRunningPlanRequest) GetSeasonMode() *int32 {
	return s.SeasonMode
}

func (s *SetRunningPlanRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *SetRunningPlanRequest) GetStatisticsTime() *string {
	return s.StatisticsTime
}

func (s *SetRunningPlanRequest) GetSystemId() *string {
	return s.SystemId
}

func (s *SetRunningPlanRequest) GetWorkingEndTime() *string {
	return s.WorkingEndTime
}

func (s *SetRunningPlanRequest) GetWorkingStartTime() *string {
	return s.WorkingStartTime
}

func (s *SetRunningPlanRequest) SetControlType(v int32) *SetRunningPlanRequest {
	s.ControlType = &v
	return s
}

func (s *SetRunningPlanRequest) SetDateType(v int32) *SetRunningPlanRequest {
	s.DateType = &v
	return s
}

func (s *SetRunningPlanRequest) SetEarliestStartupTime(v string) *SetRunningPlanRequest {
	s.EarliestStartupTime = &v
	return s
}

func (s *SetRunningPlanRequest) SetEndTime(v string) *SetRunningPlanRequest {
	s.EndTime = &v
	return s
}

func (s *SetRunningPlanRequest) SetFactoryId(v string) *SetRunningPlanRequest {
	s.FactoryId = &v
	return s
}

func (s *SetRunningPlanRequest) SetLatestShutdownTime(v string) *SetRunningPlanRequest {
	s.LatestShutdownTime = &v
	return s
}

func (s *SetRunningPlanRequest) SetMaxCarbonDioxide(v float64) *SetRunningPlanRequest {
	s.MaxCarbonDioxide = &v
	return s
}

func (s *SetRunningPlanRequest) SetMaxTem(v float64) *SetRunningPlanRequest {
	s.MaxTem = &v
	return s
}

func (s *SetRunningPlanRequest) SetMinTem(v float64) *SetRunningPlanRequest {
	s.MinTem = &v
	return s
}

func (s *SetRunningPlanRequest) SetPKey(v string) *SetRunningPlanRequest {
	s.PKey = &v
	return s
}

func (s *SetRunningPlanRequest) SetSeasonMode(v int32) *SetRunningPlanRequest {
	s.SeasonMode = &v
	return s
}

func (s *SetRunningPlanRequest) SetStartTime(v string) *SetRunningPlanRequest {
	s.StartTime = &v
	return s
}

func (s *SetRunningPlanRequest) SetStatisticsTime(v string) *SetRunningPlanRequest {
	s.StatisticsTime = &v
	return s
}

func (s *SetRunningPlanRequest) SetSystemId(v string) *SetRunningPlanRequest {
	s.SystemId = &v
	return s
}

func (s *SetRunningPlanRequest) SetWorkingEndTime(v string) *SetRunningPlanRequest {
	s.WorkingEndTime = &v
	return s
}

func (s *SetRunningPlanRequest) SetWorkingStartTime(v string) *SetRunningPlanRequest {
	s.WorkingStartTime = &v
	return s
}

func (s *SetRunningPlanRequest) Validate() error {
	return dara.Validate(s)
}
