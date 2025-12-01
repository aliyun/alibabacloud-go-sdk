// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSafetyCoverRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDateType(v string) *GetSafetyCoverRequest
	GetDateType() *string
	SetEndDate(v int64) *GetSafetyCoverRequest
	GetEndDate() *int64
	SetStartDate(v int64) *GetSafetyCoverRequest
	GetStartDate() *int64
	SetSuspEventSource(v string) *GetSafetyCoverRequest
	GetSuspEventSource() *string
}

type GetSafetyCoverRequest struct {
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
	// 1732268720000
	EndDate *int64 `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// Start time.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1732255620000
	StartDate *int64 `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	// Alert event source.
	//
	// example:
	//
	// 该接口不用传
	SuspEventSource *string `json:"SuspEventSource,omitempty" xml:"SuspEventSource,omitempty"`
}

func (s GetSafetyCoverRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSafetyCoverRequest) GoString() string {
	return s.String()
}

func (s *GetSafetyCoverRequest) GetDateType() *string {
	return s.DateType
}

func (s *GetSafetyCoverRequest) GetEndDate() *int64 {
	return s.EndDate
}

func (s *GetSafetyCoverRequest) GetStartDate() *int64 {
	return s.StartDate
}

func (s *GetSafetyCoverRequest) GetSuspEventSource() *string {
	return s.SuspEventSource
}

func (s *GetSafetyCoverRequest) SetDateType(v string) *GetSafetyCoverRequest {
	s.DateType = &v
	return s
}

func (s *GetSafetyCoverRequest) SetEndDate(v int64) *GetSafetyCoverRequest {
	s.EndDate = &v
	return s
}

func (s *GetSafetyCoverRequest) SetStartDate(v int64) *GetSafetyCoverRequest {
	s.StartDate = &v
	return s
}

func (s *GetSafetyCoverRequest) SetSuspEventSource(v string) *GetSafetyCoverRequest {
	s.SuspEventSource = &v
	return s
}

func (s *GetSafetyCoverRequest) Validate() error {
	return dara.Validate(s)
}
