// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTimeFilter interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *TimeFilter
	GetEndTime() *string
	SetStartTime(v string) *TimeFilter
	GetStartTime() *string
}

type TimeFilter struct {
	// The time when the migration task ended.
	//
	// example:
	//
	// 2006-12-31T59:59:59Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The time when the migration task started.
	//
	// example:
	//
	// 2006-01-01T00:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s TimeFilter) String() string {
	return dara.Prettify(s)
}

func (s TimeFilter) GoString() string {
	return s.String()
}

func (s *TimeFilter) GetEndTime() *string {
	return s.EndTime
}

func (s *TimeFilter) GetStartTime() *string {
	return s.StartTime
}

func (s *TimeFilter) SetEndTime(v string) *TimeFilter {
	s.EndTime = &v
	return s
}

func (s *TimeFilter) SetStartTime(v string) *TimeFilter {
	s.StartTime = &v
	return s
}

func (s *TimeFilter) Validate() error {
	return dara.Validate(s)
}
