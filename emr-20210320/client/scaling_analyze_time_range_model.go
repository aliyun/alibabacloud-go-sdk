// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iScalingAnalyzeTimeRange interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v int64) *ScalingAnalyzeTimeRange
	GetEndTime() *int64
	SetStartTime(v int64) *ScalingAnalyzeTimeRange
	GetStartTime() *int64
	SetType(v string) *ScalingAnalyzeTimeRange
	GetType() *string
}

type ScalingAnalyzeTimeRange struct {
	// 结束时间。
	//
	// example:
	//
	// 1676441972000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// 起始时间。
	//
	// example:
	//
	// 1676441971000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// 峰谷类型。 peak/valley
	//
	// example:
	//
	// peak
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ScalingAnalyzeTimeRange) String() string {
	return dara.Prettify(s)
}

func (s ScalingAnalyzeTimeRange) GoString() string {
	return s.String()
}

func (s *ScalingAnalyzeTimeRange) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ScalingAnalyzeTimeRange) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ScalingAnalyzeTimeRange) GetType() *string {
	return s.Type
}

func (s *ScalingAnalyzeTimeRange) SetEndTime(v int64) *ScalingAnalyzeTimeRange {
	s.EndTime = &v
	return s
}

func (s *ScalingAnalyzeTimeRange) SetStartTime(v int64) *ScalingAnalyzeTimeRange {
	s.StartTime = &v
	return s
}

func (s *ScalingAnalyzeTimeRange) SetType(v string) *ScalingAnalyzeTimeRange {
	s.Type = &v
	return s
}

func (s *ScalingAnalyzeTimeRange) Validate() error {
	return dara.Validate(s)
}
