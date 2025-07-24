// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTimeConstraint interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *TimeConstraint
	GetEndTime() *string
	SetStartTime(v string) *TimeConstraint
	GetStartTime() *string
}

type TimeConstraint struct {
	// 结束时间。取值范围：00:00:00至23:59:59
	//
	// example:
	//
	// 23:59:59
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// 开始时间。取值范围：00:00:00至23:59:59
	//
	// example:
	//
	// 06:00:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s TimeConstraint) String() string {
	return dara.Prettify(s)
}

func (s TimeConstraint) GoString() string {
	return s.String()
}

func (s *TimeConstraint) GetEndTime() *string {
	return s.EndTime
}

func (s *TimeConstraint) GetStartTime() *string {
	return s.StartTime
}

func (s *TimeConstraint) SetEndTime(v string) *TimeConstraint {
	s.EndTime = &v
	return s
}

func (s *TimeConstraint) SetStartTime(v string) *TimeConstraint {
	s.StartTime = &v
	return s
}

func (s *TimeConstraint) Validate() error {
	return dara.Validate(s)
}
