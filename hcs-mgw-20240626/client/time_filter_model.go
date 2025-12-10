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
	EndTime   *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
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
