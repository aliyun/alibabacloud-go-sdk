// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSavingsPlansUsageTotalShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBillOwnerId(v int64) *DescribeSavingsPlansUsageTotalShrinkRequest
	GetBillOwnerId() *int64
	SetEndPeriod(v string) *DescribeSavingsPlansUsageTotalShrinkRequest
	GetEndPeriod() *string
	SetFilterParamShrink(v string) *DescribeSavingsPlansUsageTotalShrinkRequest
	GetFilterParamShrink() *string
	SetPeriodType(v string) *DescribeSavingsPlansUsageTotalShrinkRequest
	GetPeriodType() *string
	SetStartPeriod(v string) *DescribeSavingsPlansUsageTotalShrinkRequest
	GetStartPeriod() *string
}

type DescribeSavingsPlansUsageTotalShrinkRequest struct {
	// The ID of the account with resource ownership, which is the account that actually uses the resources.
	//
	// example:
	//
	// 123745698925000
	BillOwnerId *int64 `json:"BillOwnerId,omitempty" xml:"BillOwnerId,omitempty"`
	// The end time of the time range to query. The end time is not included in the time range. If you do not specify this parameter, the end time is the current time. Specify the time in the yyyy-MM-dd HH:mm:ss format.
	//
	// example:
	//
	// 2021-01-02 00:00:00
	EndPeriod *string `json:"EndPeriod,omitempty" xml:"EndPeriod,omitempty"`
	// The filter parameters.
	FilterParamShrink *string `json:"FilterParam,omitempty" xml:"FilterParam,omitempty"`
	// The time granularity for querying the usage summary. Valid values: MONTH, DAY, and HOUR.
	//
	// This parameter is required.
	//
	// example:
	//
	// HOUR
	PeriodType *string `json:"PeriodType,omitempty" xml:"PeriodType,omitempty"`
	// The start time of the time range to query. The start time is included in the time range. Specify the time in the yyyy-MM-dd HH:mm:ss format.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2021-01-01 00:00:00
	StartPeriod *string `json:"StartPeriod,omitempty" xml:"StartPeriod,omitempty"`
}

func (s DescribeSavingsPlansUsageTotalShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSavingsPlansUsageTotalShrinkRequest) GoString() string {
	return s.String()
}

func (s *DescribeSavingsPlansUsageTotalShrinkRequest) GetBillOwnerId() *int64 {
	return s.BillOwnerId
}

func (s *DescribeSavingsPlansUsageTotalShrinkRequest) GetEndPeriod() *string {
	return s.EndPeriod
}

func (s *DescribeSavingsPlansUsageTotalShrinkRequest) GetFilterParamShrink() *string {
	return s.FilterParamShrink
}

func (s *DescribeSavingsPlansUsageTotalShrinkRequest) GetPeriodType() *string {
	return s.PeriodType
}

func (s *DescribeSavingsPlansUsageTotalShrinkRequest) GetStartPeriod() *string {
	return s.StartPeriod
}

func (s *DescribeSavingsPlansUsageTotalShrinkRequest) SetBillOwnerId(v int64) *DescribeSavingsPlansUsageTotalShrinkRequest {
	s.BillOwnerId = &v
	return s
}

func (s *DescribeSavingsPlansUsageTotalShrinkRequest) SetEndPeriod(v string) *DescribeSavingsPlansUsageTotalShrinkRequest {
	s.EndPeriod = &v
	return s
}

func (s *DescribeSavingsPlansUsageTotalShrinkRequest) SetFilterParamShrink(v string) *DescribeSavingsPlansUsageTotalShrinkRequest {
	s.FilterParamShrink = &v
	return s
}

func (s *DescribeSavingsPlansUsageTotalShrinkRequest) SetPeriodType(v string) *DescribeSavingsPlansUsageTotalShrinkRequest {
	s.PeriodType = &v
	return s
}

func (s *DescribeSavingsPlansUsageTotalShrinkRequest) SetStartPeriod(v string) *DescribeSavingsPlansUsageTotalShrinkRequest {
	s.StartPeriod = &v
	return s
}

func (s *DescribeSavingsPlansUsageTotalShrinkRequest) Validate() error {
	return dara.Validate(s)
}
