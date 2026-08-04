// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNodeMetricsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *GetNodeMetricsRequest
	GetEndTime() *string
	SetGPUType(v string) *GetNodeMetricsRequest
	GetGPUType() *string
	SetStartTime(v string) *GetNodeMetricsRequest
	GetStartTime() *string
	SetTimeStep(v string) *GetNodeMetricsRequest
	GetTimeStep() *string
	SetVerbose(v bool) *GetNodeMetricsRequest
	GetVerbose() *bool
}

type GetNodeMetricsRequest struct {
	// The end time.
	//
	// example:
	//
	// 2024-07-10T10:17:06
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The GPU type. Valid values:
	//
	// - V100
	//
	// - A100
	//
	// - A10
	//
	// - T4
	//
	// - P100
	//
	// example:
	//
	// V100
	GPUType *string `json:"GPUType,omitempty" xml:"GPUType,omitempty"`
	// The start time.
	//
	// example:
	//
	// 2024-07-08T02:23:30.292Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The time step. Default value: 5m. The supported time units are:
	//
	// - h: hours.
	//
	// - m: minutes.
	//
	// - s: seconds.
	//
	// If the value does not contain a unit, the default unit is s (seconds).
	//
	// example:
	//
	// 1h
	TimeStep *string `json:"TimeStep,omitempty" xml:"TimeStep,omitempty"`
	// Specifies whether to display non-essential information. Non-essential information currently includes Labels.
	//
	// Valid values:
	//
	// false: does not display non-essential information (default value)
	//
	// true: displays non-essential information
	//
	// example:
	//
	// true
	Verbose *bool `json:"Verbose,omitempty" xml:"Verbose,omitempty"`
}

func (s GetNodeMetricsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetNodeMetricsRequest) GoString() string {
	return s.String()
}

func (s *GetNodeMetricsRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *GetNodeMetricsRequest) GetGPUType() *string {
	return s.GPUType
}

func (s *GetNodeMetricsRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *GetNodeMetricsRequest) GetTimeStep() *string {
	return s.TimeStep
}

func (s *GetNodeMetricsRequest) GetVerbose() *bool {
	return s.Verbose
}

func (s *GetNodeMetricsRequest) SetEndTime(v string) *GetNodeMetricsRequest {
	s.EndTime = &v
	return s
}

func (s *GetNodeMetricsRequest) SetGPUType(v string) *GetNodeMetricsRequest {
	s.GPUType = &v
	return s
}

func (s *GetNodeMetricsRequest) SetStartTime(v string) *GetNodeMetricsRequest {
	s.StartTime = &v
	return s
}

func (s *GetNodeMetricsRequest) SetTimeStep(v string) *GetNodeMetricsRequest {
	s.TimeStep = &v
	return s
}

func (s *GetNodeMetricsRequest) SetVerbose(v bool) *GetNodeMetricsRequest {
	s.Verbose = &v
	return s
}

func (s *GetNodeMetricsRequest) Validate() error {
	return dara.Validate(s)
}
