// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRLProgressStage interface {
	dara.Model
	String() string
	GoString() string
	SetDuration(v float64) *RLProgressStage
	GetDuration() *float64
	SetEndTime(v int64) *RLProgressStage
	GetEndTime() *int64
	SetKey(v string) *RLProgressStage
	GetKey() *string
	SetLabel(v string) *RLProgressStage
	GetLabel() *string
	SetMarker(v string) *RLProgressStage
	GetMarker() *string
	SetOptional(v bool) *RLProgressStage
	GetOptional() *bool
	SetStartTime(v int64) *RLProgressStage
	GetStartTime() *int64
	SetStatus(v string) *RLProgressStage
	GetStatus() *string
}

type RLProgressStage struct {
	// The duration of the stage in seconds, with three decimal places. A step often falls entirely within the same second, so the value is not rounded to an integer.
	//
	// example:
	//
	// 0.483
	Duration *float64 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The end time of the stage in UNIX seconds.
	//
	// example:
	//
	// 1787474487
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The stage identifier.
	//
	// example:
	//
	// traj
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The Chinese name of the stage.
	//
	// example:
	//
	// Generate trace
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The log marker text used to match the stage.
	//
	// example:
	//
	// start/end generation
	Marker *string `json:"Marker,omitempty" xml:"Marker,omitempty"`
	// Specifies whether the stage is optional. If an optional stage does not occur, its status is recorded as skipped.
	//
	// example:
	//
	// false
	Optional *bool `json:"Optional,omitempty" xml:"Optional,omitempty"`
	// The start time of the stage in UNIX seconds.
	//
	// example:
	//
	// 1787474487
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// done / running / waiting / pending / skipped
	//
	// example:
	//
	// done
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s RLProgressStage) String() string {
	return dara.Prettify(s)
}

func (s RLProgressStage) GoString() string {
	return s.String()
}

func (s *RLProgressStage) GetDuration() *float64 {
	return s.Duration
}

func (s *RLProgressStage) GetEndTime() *int64 {
	return s.EndTime
}

func (s *RLProgressStage) GetKey() *string {
	return s.Key
}

func (s *RLProgressStage) GetLabel() *string {
	return s.Label
}

func (s *RLProgressStage) GetMarker() *string {
	return s.Marker
}

func (s *RLProgressStage) GetOptional() *bool {
	return s.Optional
}

func (s *RLProgressStage) GetStartTime() *int64 {
	return s.StartTime
}

func (s *RLProgressStage) GetStatus() *string {
	return s.Status
}

func (s *RLProgressStage) SetDuration(v float64) *RLProgressStage {
	s.Duration = &v
	return s
}

func (s *RLProgressStage) SetEndTime(v int64) *RLProgressStage {
	s.EndTime = &v
	return s
}

func (s *RLProgressStage) SetKey(v string) *RLProgressStage {
	s.Key = &v
	return s
}

func (s *RLProgressStage) SetLabel(v string) *RLProgressStage {
	s.Label = &v
	return s
}

func (s *RLProgressStage) SetMarker(v string) *RLProgressStage {
	s.Marker = &v
	return s
}

func (s *RLProgressStage) SetOptional(v bool) *RLProgressStage {
	s.Optional = &v
	return s
}

func (s *RLProgressStage) SetStartTime(v int64) *RLProgressStage {
	s.StartTime = &v
	return s
}

func (s *RLProgressStage) SetStatus(v string) *RLProgressStage {
	s.Status = &v
	return s
}

func (s *RLProgressStage) Validate() error {
	return dara.Validate(s)
}
