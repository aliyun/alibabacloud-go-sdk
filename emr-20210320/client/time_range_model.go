// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTimeRange interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *TimeRange
	GetEndTime() *string
	SetStartTime(v string) *TimeRange
	GetStartTime() *string
}

type TimeRange struct {
	// 结束时间。
	//
	// example:
	//
	// 1676441972000
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// 起始时间。
	//
	// example:
	//
	// 1676441971000
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s TimeRange) String() string {
	return dara.Prettify(s)
}

func (s TimeRange) GoString() string {
	return s.String()
}

func (s *TimeRange) GetEndTime() *string {
	return s.EndTime
}

func (s *TimeRange) GetStartTime() *string {
	return s.StartTime
}

func (s *TimeRange) SetEndTime(v string) *TimeRange {
	s.EndTime = &v
	return s
}

func (s *TimeRange) SetStartTime(v string) *TimeRange {
	s.StartTime = &v
	return s
}

func (s *TimeRange) Validate() error {
	return dara.Validate(s)
}
