// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOnCallSchedulesDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *GetOnCallSchedulesDetailRequest
	GetEndTime() *string
	SetId(v int64) *GetOnCallSchedulesDetailRequest
	GetId() *int64
	SetStartTime(v string) *GetOnCallSchedulesDetailRequest
	GetStartTime() *string
}

type GetOnCallSchedulesDetailRequest struct {
	// The date on which the shift ends. Format: `yyyy-MM-dd`.
	//
	// example:
	//
	// 2022-10-30
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the scheduling policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The date from which the shift starts. Format: `yyyy-MM-dd`.
	//
	// example:
	//
	// 2022-10-01
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s GetOnCallSchedulesDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s GetOnCallSchedulesDetailRequest) GoString() string {
	return s.String()
}

func (s *GetOnCallSchedulesDetailRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *GetOnCallSchedulesDetailRequest) GetId() *int64 {
	return s.Id
}

func (s *GetOnCallSchedulesDetailRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *GetOnCallSchedulesDetailRequest) SetEndTime(v string) *GetOnCallSchedulesDetailRequest {
	s.EndTime = &v
	return s
}

func (s *GetOnCallSchedulesDetailRequest) SetId(v int64) *GetOnCallSchedulesDetailRequest {
	s.Id = &v
	return s
}

func (s *GetOnCallSchedulesDetailRequest) SetStartTime(v string) *GetOnCallSchedulesDetailRequest {
	s.StartTime = &v
	return s
}

func (s *GetOnCallSchedulesDetailRequest) Validate() error {
	return dara.Validate(s)
}
