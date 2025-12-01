// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSuspEventSummaryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDateType(v string) *GetSuspEventSummaryRequest
	GetDateType() *string
	SetEndDate(v int64) *GetSuspEventSummaryRequest
	GetEndDate() *int64
	SetStartDate(v int64) *GetSuspEventSummaryRequest
	GetStartDate() *int64
	SetSuspEventSource(v string) *GetSuspEventSummaryRequest
	GetSuspEventSource() *string
}

type GetSuspEventSummaryRequest struct {
	// Filter time type. Supports filtering by the last 7 days, the last 30 days, the last half year, or custom time ranges.
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
	// SUSP_EVENT
	SuspEventSource *string `json:"SuspEventSource,omitempty" xml:"SuspEventSource,omitempty"`
}

func (s GetSuspEventSummaryRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSuspEventSummaryRequest) GoString() string {
	return s.String()
}

func (s *GetSuspEventSummaryRequest) GetDateType() *string {
	return s.DateType
}

func (s *GetSuspEventSummaryRequest) GetEndDate() *int64 {
	return s.EndDate
}

func (s *GetSuspEventSummaryRequest) GetStartDate() *int64 {
	return s.StartDate
}

func (s *GetSuspEventSummaryRequest) GetSuspEventSource() *string {
	return s.SuspEventSource
}

func (s *GetSuspEventSummaryRequest) SetDateType(v string) *GetSuspEventSummaryRequest {
	s.DateType = &v
	return s
}

func (s *GetSuspEventSummaryRequest) SetEndDate(v int64) *GetSuspEventSummaryRequest {
	s.EndDate = &v
	return s
}

func (s *GetSuspEventSummaryRequest) SetStartDate(v int64) *GetSuspEventSummaryRequest {
	s.StartDate = &v
	return s
}

func (s *GetSuspEventSummaryRequest) SetSuspEventSource(v string) *GetSuspEventSummaryRequest {
	s.SuspEventSource = &v
	return s
}

func (s *GetSuspEventSummaryRequest) Validate() error {
	return dara.Validate(s)
}
