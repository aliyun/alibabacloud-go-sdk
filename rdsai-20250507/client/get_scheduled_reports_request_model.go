// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetScheduledReportsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *GetScheduledReportsRequest
	GetEndTime() *string
	SetPageNumber(v int64) *GetScheduledReportsRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *GetScheduledReportsRequest
	GetPageSize() *int64
	SetScheduledId(v string) *GetScheduledReportsRequest
	GetScheduledId() *string
	SetStartTime(v string) *GetScheduledReportsRequest
	GetStartTime() *string
}

type GetScheduledReportsRequest struct {
	// example:
	//
	// 2026-01-25T02:02:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 847268a4-196f-416b-aa12-bfe0c115****
	ScheduledId *string `json:"ScheduledId,omitempty" xml:"ScheduledId,omitempty"`
	// example:
	//
	// 2026-01-25T01:02:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s GetScheduledReportsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetScheduledReportsRequest) GoString() string {
	return s.String()
}

func (s *GetScheduledReportsRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *GetScheduledReportsRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *GetScheduledReportsRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *GetScheduledReportsRequest) GetScheduledId() *string {
	return s.ScheduledId
}

func (s *GetScheduledReportsRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *GetScheduledReportsRequest) SetEndTime(v string) *GetScheduledReportsRequest {
	s.EndTime = &v
	return s
}

func (s *GetScheduledReportsRequest) SetPageNumber(v int64) *GetScheduledReportsRequest {
	s.PageNumber = &v
	return s
}

func (s *GetScheduledReportsRequest) SetPageSize(v int64) *GetScheduledReportsRequest {
	s.PageSize = &v
	return s
}

func (s *GetScheduledReportsRequest) SetScheduledId(v string) *GetScheduledReportsRequest {
	s.ScheduledId = &v
	return s
}

func (s *GetScheduledReportsRequest) SetStartTime(v string) *GetScheduledReportsRequest {
	s.StartTime = &v
	return s
}

func (s *GetScheduledReportsRequest) Validate() error {
	return dara.Validate(s)
}
