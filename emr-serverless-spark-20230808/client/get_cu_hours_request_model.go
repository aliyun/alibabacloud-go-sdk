// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCuHoursRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *GetCuHoursRequest
	GetEndTime() *string
	SetStartTime(v string) *GetCuHoursRequest
	GetStartTime() *string
}

type GetCuHoursRequest struct {
	// The end time of the query time range.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2024-01-08 00:00:00
	EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// The start time of the query time range.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2024-01-01 00:00:00
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
}

func (s GetCuHoursRequest) String() string {
	return dara.Prettify(s)
}

func (s GetCuHoursRequest) GoString() string {
	return s.String()
}

func (s *GetCuHoursRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *GetCuHoursRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *GetCuHoursRequest) SetEndTime(v string) *GetCuHoursRequest {
	s.EndTime = &v
	return s
}

func (s *GetCuHoursRequest) SetStartTime(v string) *GetCuHoursRequest {
	s.StartTime = &v
	return s
}

func (s *GetCuHoursRequest) Validate() error {
	return dara.Validate(s)
}
