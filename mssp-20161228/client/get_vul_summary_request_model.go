// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVulSummaryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDateType(v string) *GetVulSummaryRequest
	GetDateType() *string
	SetEndDate(v int64) *GetVulSummaryRequest
	GetEndDate() *int64
	SetStartDate(v int64) *GetVulSummaryRequest
	GetStartDate() *int64
	SetSuspEventSource(v string) *GetVulSummaryRequest
	GetSuspEventSource() *string
}

type GetVulSummaryRequest struct {
	// Filter time type. Supports filtering by the last 7 days, the last 30 days, the last half year, or a custom time range.
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

func (s GetVulSummaryRequest) String() string {
	return dara.Prettify(s)
}

func (s GetVulSummaryRequest) GoString() string {
	return s.String()
}

func (s *GetVulSummaryRequest) GetDateType() *string {
	return s.DateType
}

func (s *GetVulSummaryRequest) GetEndDate() *int64 {
	return s.EndDate
}

func (s *GetVulSummaryRequest) GetStartDate() *int64 {
	return s.StartDate
}

func (s *GetVulSummaryRequest) GetSuspEventSource() *string {
	return s.SuspEventSource
}

func (s *GetVulSummaryRequest) SetDateType(v string) *GetVulSummaryRequest {
	s.DateType = &v
	return s
}

func (s *GetVulSummaryRequest) SetEndDate(v int64) *GetVulSummaryRequest {
	s.EndDate = &v
	return s
}

func (s *GetVulSummaryRequest) SetStartDate(v int64) *GetVulSummaryRequest {
	s.StartDate = &v
	return s
}

func (s *GetVulSummaryRequest) SetSuspEventSource(v string) *GetVulSummaryRequest {
	s.SuspEventSource = &v
	return s
}

func (s *GetVulSummaryRequest) Validate() error {
	return dara.Validate(s)
}
