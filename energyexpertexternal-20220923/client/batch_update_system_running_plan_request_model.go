// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchUpdateSystemRunningPlanRequest interface {
	dara.Model
	String() string
	GoString() string
	SetControlType(v int32) *BatchUpdateSystemRunningPlanRequest
	GetControlType() *int32
	SetDateType(v int32) *BatchUpdateSystemRunningPlanRequest
	GetDateType() *int32
	SetEarliestStartupTime(v string) *BatchUpdateSystemRunningPlanRequest
	GetEarliestStartupTime() *string
	SetEndTime(v string) *BatchUpdateSystemRunningPlanRequest
	GetEndTime() *string
	SetFactoryId(v string) *BatchUpdateSystemRunningPlanRequest
	GetFactoryId() *string
	SetLatestShutdownTime(v string) *BatchUpdateSystemRunningPlanRequest
	GetLatestShutdownTime() *string
	SetMaxCarbonDioxide(v float64) *BatchUpdateSystemRunningPlanRequest
	GetMaxCarbonDioxide() *float64
	SetMaxTem(v float64) *BatchUpdateSystemRunningPlanRequest
	GetMaxTem() *float64
	SetMinTem(v float64) *BatchUpdateSystemRunningPlanRequest
	GetMinTem() *float64
	SetSeasonMode(v int32) *BatchUpdateSystemRunningPlanRequest
	GetSeasonMode() *int32
	SetStartTime(v string) *BatchUpdateSystemRunningPlanRequest
	GetStartTime() *string
	SetSystemId(v string) *BatchUpdateSystemRunningPlanRequest
	GetSystemId() *string
	SetWorkingEndTime(v string) *BatchUpdateSystemRunningPlanRequest
	GetWorkingEndTime() *string
	SetWorkingStartTime(v string) *BatchUpdateSystemRunningPlanRequest
	GetWorkingStartTime() *string
}

type BatchUpdateSystemRunningPlanRequest struct {
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
	// This parameter is required.
	//
	// example:
	//
	// 2024-08-30
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
	// 37.1
	MaxCarbonDioxide *float64 `json:"maxCarbonDioxide,omitempty" xml:"maxCarbonDioxide,omitempty"`
	// example:
	//
	// 25.3
	MaxTem *float64 `json:"maxTem,omitempty" xml:"maxTem,omitempty"`
	// example:
	//
	// 20.1
	MinTem *float64 `json:"minTem,omitempty" xml:"minTem,omitempty"`
	// example:
	//
	// 0
	SeasonMode *int32 `json:"seasonMode,omitempty" xml:"seasonMode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2024-08-21
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
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

func (s BatchUpdateSystemRunningPlanRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchUpdateSystemRunningPlanRequest) GoString() string {
	return s.String()
}

func (s *BatchUpdateSystemRunningPlanRequest) GetControlType() *int32 {
	return s.ControlType
}

func (s *BatchUpdateSystemRunningPlanRequest) GetDateType() *int32 {
	return s.DateType
}

func (s *BatchUpdateSystemRunningPlanRequest) GetEarliestStartupTime() *string {
	return s.EarliestStartupTime
}

func (s *BatchUpdateSystemRunningPlanRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *BatchUpdateSystemRunningPlanRequest) GetFactoryId() *string {
	return s.FactoryId
}

func (s *BatchUpdateSystemRunningPlanRequest) GetLatestShutdownTime() *string {
	return s.LatestShutdownTime
}

func (s *BatchUpdateSystemRunningPlanRequest) GetMaxCarbonDioxide() *float64 {
	return s.MaxCarbonDioxide
}

func (s *BatchUpdateSystemRunningPlanRequest) GetMaxTem() *float64 {
	return s.MaxTem
}

func (s *BatchUpdateSystemRunningPlanRequest) GetMinTem() *float64 {
	return s.MinTem
}

func (s *BatchUpdateSystemRunningPlanRequest) GetSeasonMode() *int32 {
	return s.SeasonMode
}

func (s *BatchUpdateSystemRunningPlanRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *BatchUpdateSystemRunningPlanRequest) GetSystemId() *string {
	return s.SystemId
}

func (s *BatchUpdateSystemRunningPlanRequest) GetWorkingEndTime() *string {
	return s.WorkingEndTime
}

func (s *BatchUpdateSystemRunningPlanRequest) GetWorkingStartTime() *string {
	return s.WorkingStartTime
}

func (s *BatchUpdateSystemRunningPlanRequest) SetControlType(v int32) *BatchUpdateSystemRunningPlanRequest {
	s.ControlType = &v
	return s
}

func (s *BatchUpdateSystemRunningPlanRequest) SetDateType(v int32) *BatchUpdateSystemRunningPlanRequest {
	s.DateType = &v
	return s
}

func (s *BatchUpdateSystemRunningPlanRequest) SetEarliestStartupTime(v string) *BatchUpdateSystemRunningPlanRequest {
	s.EarliestStartupTime = &v
	return s
}

func (s *BatchUpdateSystemRunningPlanRequest) SetEndTime(v string) *BatchUpdateSystemRunningPlanRequest {
	s.EndTime = &v
	return s
}

func (s *BatchUpdateSystemRunningPlanRequest) SetFactoryId(v string) *BatchUpdateSystemRunningPlanRequest {
	s.FactoryId = &v
	return s
}

func (s *BatchUpdateSystemRunningPlanRequest) SetLatestShutdownTime(v string) *BatchUpdateSystemRunningPlanRequest {
	s.LatestShutdownTime = &v
	return s
}

func (s *BatchUpdateSystemRunningPlanRequest) SetMaxCarbonDioxide(v float64) *BatchUpdateSystemRunningPlanRequest {
	s.MaxCarbonDioxide = &v
	return s
}

func (s *BatchUpdateSystemRunningPlanRequest) SetMaxTem(v float64) *BatchUpdateSystemRunningPlanRequest {
	s.MaxTem = &v
	return s
}

func (s *BatchUpdateSystemRunningPlanRequest) SetMinTem(v float64) *BatchUpdateSystemRunningPlanRequest {
	s.MinTem = &v
	return s
}

func (s *BatchUpdateSystemRunningPlanRequest) SetSeasonMode(v int32) *BatchUpdateSystemRunningPlanRequest {
	s.SeasonMode = &v
	return s
}

func (s *BatchUpdateSystemRunningPlanRequest) SetStartTime(v string) *BatchUpdateSystemRunningPlanRequest {
	s.StartTime = &v
	return s
}

func (s *BatchUpdateSystemRunningPlanRequest) SetSystemId(v string) *BatchUpdateSystemRunningPlanRequest {
	s.SystemId = &v
	return s
}

func (s *BatchUpdateSystemRunningPlanRequest) SetWorkingEndTime(v string) *BatchUpdateSystemRunningPlanRequest {
	s.WorkingEndTime = &v
	return s
}

func (s *BatchUpdateSystemRunningPlanRequest) SetWorkingStartTime(v string) *BatchUpdateSystemRunningPlanRequest {
	s.WorkingStartTime = &v
	return s
}

func (s *BatchUpdateSystemRunningPlanRequest) Validate() error {
	return dara.Validate(s)
}
