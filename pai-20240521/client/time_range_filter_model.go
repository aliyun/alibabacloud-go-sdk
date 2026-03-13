// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTimeRangeFilter interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *TimeRangeFilter
	GetEndTime() *string
	SetStartTime(v string) *TimeRangeFilter
	GetStartTime() *string
}

type TimeRangeFilter struct {
	// example:
	//
	// 2023-06-22T00:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 2023-06-22T00:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s TimeRangeFilter) String() string {
	return dara.Prettify(s)
}

func (s TimeRangeFilter) GoString() string {
	return s.String()
}

func (s *TimeRangeFilter) GetEndTime() *string {
	return s.EndTime
}

func (s *TimeRangeFilter) GetStartTime() *string {
	return s.StartTime
}

func (s *TimeRangeFilter) SetEndTime(v string) *TimeRangeFilter {
	s.EndTime = &v
	return s
}

func (s *TimeRangeFilter) SetStartTime(v string) *TimeRangeFilter {
	s.StartTime = &v
	return s
}

func (s *TimeRangeFilter) Validate() error {
	return dara.Validate(s)
}
