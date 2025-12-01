// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSowListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDateType(v string) *GetSowListRequest
	GetDateType() *string
	SetEndDate(v int64) *GetSowListRequest
	GetEndDate() *int64
	SetStartDate(v int64) *GetSowListRequest
	GetStartDate() *int64
	SetSuspEventSource(v string) *GetSowListRequest
	GetSuspEventSource() *string
}

type GetSowListRequest struct {
	// Filter time type, supports filtering by the last 7 days, the last 30 days, the last half year, or custom time ranges.
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

func (s GetSowListRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSowListRequest) GoString() string {
	return s.String()
}

func (s *GetSowListRequest) GetDateType() *string {
	return s.DateType
}

func (s *GetSowListRequest) GetEndDate() *int64 {
	return s.EndDate
}

func (s *GetSowListRequest) GetStartDate() *int64 {
	return s.StartDate
}

func (s *GetSowListRequest) GetSuspEventSource() *string {
	return s.SuspEventSource
}

func (s *GetSowListRequest) SetDateType(v string) *GetSowListRequest {
	s.DateType = &v
	return s
}

func (s *GetSowListRequest) SetEndDate(v int64) *GetSowListRequest {
	s.EndDate = &v
	return s
}

func (s *GetSowListRequest) SetStartDate(v int64) *GetSowListRequest {
	s.StartDate = &v
	return s
}

func (s *GetSowListRequest) SetSuspEventSource(v string) *GetSowListRequest {
	s.SuspEventSource = &v
	return s
}

func (s *GetSowListRequest) Validate() error {
	return dara.Validate(s)
}
