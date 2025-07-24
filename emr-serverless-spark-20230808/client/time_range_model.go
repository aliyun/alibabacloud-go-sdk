// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTimeRange interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v int64) *TimeRange
	GetEndTime() *int64
	SetStartTime(v int64) *TimeRange
	GetStartTime() *int64
}

type TimeRange struct {
	// 时间范围结束时间。
	//
	// example:
	//
	// 1688370894339
	EndTime *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// 时间范围开始时间。
	//
	// example:
	//
	// 1688370894339
	StartTime *int64 `json:"startTime,omitempty" xml:"startTime,omitempty"`
}

func (s TimeRange) String() string {
	return dara.Prettify(s)
}

func (s TimeRange) GoString() string {
	return s.String()
}

func (s *TimeRange) GetEndTime() *int64 {
	return s.EndTime
}

func (s *TimeRange) GetStartTime() *int64 {
	return s.StartTime
}

func (s *TimeRange) SetEndTime(v int64) *TimeRange {
	s.EndTime = &v
	return s
}

func (s *TimeRange) SetStartTime(v int64) *TimeRange {
	s.StartTime = &v
	return s
}

func (s *TimeRange) Validate() error {
	return dara.Validate(s)
}
