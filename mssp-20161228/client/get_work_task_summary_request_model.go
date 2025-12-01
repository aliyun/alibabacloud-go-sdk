// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWorkTaskSummaryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDateType(v string) *GetWorkTaskSummaryRequest
	GetDateType() *string
	SetEndDate(v int64) *GetWorkTaskSummaryRequest
	GetEndDate() *int64
	SetStartDate(v int64) *GetWorkTaskSummaryRequest
	GetStartDate() *int64
	SetSuspEventSource(v string) *GetWorkTaskSummaryRequest
	GetSuspEventSource() *string
}

type GetWorkTaskSummaryRequest struct {
	// Filter time type, supports filtering by the last 7 days, the last 30 days, the last half year, or custom time periods.
	//
	// This parameter is required.
	//
	// example:
	//
	// month
	DateType *string `json:"DateType,omitempty" xml:"DateType,omitempty"`
	// End time.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1732156885986
	EndDate *int64 `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// Start time.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1729478485000
	StartDate *int64 `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	// Alert event source.
	//
	// example:
	//
	// 该字段暂时未用，有问题请联系管理员
	SuspEventSource *string `json:"SuspEventSource,omitempty" xml:"SuspEventSource,omitempty"`
}

func (s GetWorkTaskSummaryRequest) String() string {
	return dara.Prettify(s)
}

func (s GetWorkTaskSummaryRequest) GoString() string {
	return s.String()
}

func (s *GetWorkTaskSummaryRequest) GetDateType() *string {
	return s.DateType
}

func (s *GetWorkTaskSummaryRequest) GetEndDate() *int64 {
	return s.EndDate
}

func (s *GetWorkTaskSummaryRequest) GetStartDate() *int64 {
	return s.StartDate
}

func (s *GetWorkTaskSummaryRequest) GetSuspEventSource() *string {
	return s.SuspEventSource
}

func (s *GetWorkTaskSummaryRequest) SetDateType(v string) *GetWorkTaskSummaryRequest {
	s.DateType = &v
	return s
}

func (s *GetWorkTaskSummaryRequest) SetEndDate(v int64) *GetWorkTaskSummaryRequest {
	s.EndDate = &v
	return s
}

func (s *GetWorkTaskSummaryRequest) SetStartDate(v int64) *GetWorkTaskSummaryRequest {
	s.StartDate = &v
	return s
}

func (s *GetWorkTaskSummaryRequest) SetSuspEventSource(v string) *GetWorkTaskSummaryRequest {
	s.SuspEventSource = &v
	return s
}

func (s *GetWorkTaskSummaryRequest) Validate() error {
	return dara.Validate(s)
}
