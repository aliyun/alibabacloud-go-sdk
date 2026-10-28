// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iViewOncallScheduleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndDate(v string) *ViewOncallScheduleRequest
	GetEndDate() *string
	SetStartDate(v string) *ViewOncallScheduleRequest
	GetStartDate() *string
}

type ViewOncallScheduleRequest struct {
	// example:
	//
	// 2024-07-05
	EndDate *string `json:"endDate,omitempty" xml:"endDate,omitempty"`
	// example:
	//
	// 2024-04-04
	StartDate *string `json:"startDate,omitempty" xml:"startDate,omitempty"`
}

func (s ViewOncallScheduleRequest) String() string {
	return dara.Prettify(s)
}

func (s ViewOncallScheduleRequest) GoString() string {
	return s.String()
}

func (s *ViewOncallScheduleRequest) GetEndDate() *string {
	return s.EndDate
}

func (s *ViewOncallScheduleRequest) GetStartDate() *string {
	return s.StartDate
}

func (s *ViewOncallScheduleRequest) SetEndDate(v string) *ViewOncallScheduleRequest {
	s.EndDate = &v
	return s
}

func (s *ViewOncallScheduleRequest) SetStartDate(v string) *ViewOncallScheduleRequest {
	s.StartDate = &v
	return s
}

func (s *ViewOncallScheduleRequest) Validate() error {
	return dara.Validate(s)
}
