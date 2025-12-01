// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetConsoleScoreRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDateType(v string) *GetConsoleScoreRequest
	GetDateType() *string
	SetEndDate(v int64) *GetConsoleScoreRequest
	GetEndDate() *int64
	SetStartDate(v int64) *GetConsoleScoreRequest
	GetStartDate() *int64
	SetSuspEventSource(v string) *GetConsoleScoreRequest
	GetSuspEventSource() *string
}

type GetConsoleScoreRequest struct {
	// Filter time type, supports filtering by the last 7 days, last 30 days, last half year, or custom. If empty, it represents the last 7 days.
	//
	// This parameter is required.
	//
	// example:
	//
	// month
	DateType *string `json:"DateType,omitempty" xml:"DateType,omitempty"`
	// End date.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1732156885986
	EndDate *int64 `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// Start date.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1729478485000
	StartDate *int64 `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	// Source of alert events.
	//
	// example:
	//
	// 该字段暂未使用，有问题请联系管理员
	SuspEventSource *string `json:"SuspEventSource,omitempty" xml:"SuspEventSource,omitempty"`
}

func (s GetConsoleScoreRequest) String() string {
	return dara.Prettify(s)
}

func (s GetConsoleScoreRequest) GoString() string {
	return s.String()
}

func (s *GetConsoleScoreRequest) GetDateType() *string {
	return s.DateType
}

func (s *GetConsoleScoreRequest) GetEndDate() *int64 {
	return s.EndDate
}

func (s *GetConsoleScoreRequest) GetStartDate() *int64 {
	return s.StartDate
}

func (s *GetConsoleScoreRequest) GetSuspEventSource() *string {
	return s.SuspEventSource
}

func (s *GetConsoleScoreRequest) SetDateType(v string) *GetConsoleScoreRequest {
	s.DateType = &v
	return s
}

func (s *GetConsoleScoreRequest) SetEndDate(v int64) *GetConsoleScoreRequest {
	s.EndDate = &v
	return s
}

func (s *GetConsoleScoreRequest) SetStartDate(v int64) *GetConsoleScoreRequest {
	s.StartDate = &v
	return s
}

func (s *GetConsoleScoreRequest) SetSuspEventSource(v string) *GetConsoleScoreRequest {
	s.SuspEventSource = &v
	return s
}

func (s *GetConsoleScoreRequest) Validate() error {
	return dara.Validate(s)
}
