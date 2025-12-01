// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBaselineSummaryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDateType(v string) *GetBaselineSummaryRequest
	GetDateType() *string
	SetEndDate(v int64) *GetBaselineSummaryRequest
	GetEndDate() *int64
	SetStartDate(v int64) *GetBaselineSummaryRequest
	GetStartDate() *int64
	SetSuspEventSource(v string) *GetBaselineSummaryRequest
	GetSuspEventSource() *string
}

type GetBaselineSummaryRequest struct {
	// Time filter type, supports filtering by the last 7 days, the last 30 days, the last half year, or custom time periods.
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
	// 该字段暂未使用，有问题请联系管理员
	SuspEventSource *string `json:"SuspEventSource,omitempty" xml:"SuspEventSource,omitempty"`
}

func (s GetBaselineSummaryRequest) String() string {
	return dara.Prettify(s)
}

func (s GetBaselineSummaryRequest) GoString() string {
	return s.String()
}

func (s *GetBaselineSummaryRequest) GetDateType() *string {
	return s.DateType
}

func (s *GetBaselineSummaryRequest) GetEndDate() *int64 {
	return s.EndDate
}

func (s *GetBaselineSummaryRequest) GetStartDate() *int64 {
	return s.StartDate
}

func (s *GetBaselineSummaryRequest) GetSuspEventSource() *string {
	return s.SuspEventSource
}

func (s *GetBaselineSummaryRequest) SetDateType(v string) *GetBaselineSummaryRequest {
	s.DateType = &v
	return s
}

func (s *GetBaselineSummaryRequest) SetEndDate(v int64) *GetBaselineSummaryRequest {
	s.EndDate = &v
	return s
}

func (s *GetBaselineSummaryRequest) SetStartDate(v int64) *GetBaselineSummaryRequest {
	s.StartDate = &v
	return s
}

func (s *GetBaselineSummaryRequest) SetSuspEventSource(v string) *GetBaselineSummaryRequest {
	s.SuspEventSource = &v
	return s
}

func (s *GetBaselineSummaryRequest) Validate() error {
	return dara.Validate(s)
}
