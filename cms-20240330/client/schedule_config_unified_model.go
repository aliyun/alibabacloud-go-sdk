// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iScheduleConfigUnified interface {
	dara.Model
	String() string
	GoString() string
	SetIntervalSecs(v int32) *ScheduleConfigUnified
	GetIntervalSecs() *int32
	SetType(v string) *ScheduleConfigUnified
	GetType() *string
}

type ScheduleConfigUnified struct {
	// The scheduling interval in seconds. This parameter is used when type is set to FIXED.
	//
	// example:
	//
	// 30
	IntervalSecs *int32 `json:"intervalSecs,omitempty" xml:"intervalSecs,omitempty"`
	// The type of the scheduling configuration. FIXED indicates fixed-interval scheduling, which executes periodically based on the interval specified by intervalSecs.
	//
	// This parameter is required.
	//
	// example:
	//
	// FIXED
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ScheduleConfigUnified) String() string {
	return dara.Prettify(s)
}

func (s ScheduleConfigUnified) GoString() string {
	return s.String()
}

func (s *ScheduleConfigUnified) GetIntervalSecs() *int32 {
	return s.IntervalSecs
}

func (s *ScheduleConfigUnified) GetType() *string {
	return s.Type
}

func (s *ScheduleConfigUnified) SetIntervalSecs(v int32) *ScheduleConfigUnified {
	s.IntervalSecs = &v
	return s
}

func (s *ScheduleConfigUnified) SetType(v string) *ScheduleConfigUnified {
	s.Type = &v
	return s
}

func (s *ScheduleConfigUnified) Validate() error {
	return dara.Validate(s)
}
