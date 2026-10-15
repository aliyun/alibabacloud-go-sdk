// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSavingsPlansCoverageTotalShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBillOwnerId(v int64) *DescribeSavingsPlansCoverageTotalShrinkRequest
	GetBillOwnerId() *int64
	SetEndPeriod(v string) *DescribeSavingsPlansCoverageTotalShrinkRequest
	GetEndPeriod() *string
	SetFilterParamShrink(v string) *DescribeSavingsPlansCoverageTotalShrinkRequest
	GetFilterParamShrink() *string
	SetPeriodType(v string) *DescribeSavingsPlansCoverageTotalShrinkRequest
	GetPeriodType() *string
	SetStartPeriod(v string) *DescribeSavingsPlansCoverageTotalShrinkRequest
	GetStartPeriod() *string
}

type DescribeSavingsPlansCoverageTotalShrinkRequest struct {
	// The ID of the resource ownership account. The resource ownership account is the account that actually uses the resource.
	//
	// example:
	//
	// 123745698925000
	BillOwnerId *int64 `json:"BillOwnerId,omitempty" xml:"BillOwnerId,omitempty"`
	// Format: yyyy-MM-dd HH:mm:ss. The query excludes this time. If this parameter is left empty, the query covers data up to the current time.
	//
	// example:
	//
	// 2021-07-20 00:00:00
	EndPeriod         *string `json:"EndPeriod,omitempty" xml:"EndPeriod,omitempty"`
	FilterParamShrink *string `json:"FilterParam,omitempty" xml:"FilterParam,omitempty"`
	// Valid values:
	//
	// - MONTH: month granularity.
	//
	// - DAY: day granularity.
	//
	// - HOUR: hour granularity.
	//
	// This parameter is required.
	//
	// example:
	//
	// HOUR
	PeriodType *string `json:"PeriodType,omitempty" xml:"PeriodType,omitempty"`
	// Format: yyyy-MM-dd HH:mm:ss. The query includes this time.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2021-07-15 00:00:00
	StartPeriod *string `json:"StartPeriod,omitempty" xml:"StartPeriod,omitempty"`
}

func (s DescribeSavingsPlansCoverageTotalShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSavingsPlansCoverageTotalShrinkRequest) GoString() string {
	return s.String()
}

func (s *DescribeSavingsPlansCoverageTotalShrinkRequest) GetBillOwnerId() *int64 {
	return s.BillOwnerId
}

func (s *DescribeSavingsPlansCoverageTotalShrinkRequest) GetEndPeriod() *string {
	return s.EndPeriod
}

func (s *DescribeSavingsPlansCoverageTotalShrinkRequest) GetFilterParamShrink() *string {
	return s.FilterParamShrink
}

func (s *DescribeSavingsPlansCoverageTotalShrinkRequest) GetPeriodType() *string {
	return s.PeriodType
}

func (s *DescribeSavingsPlansCoverageTotalShrinkRequest) GetStartPeriod() *string {
	return s.StartPeriod
}

func (s *DescribeSavingsPlansCoverageTotalShrinkRequest) SetBillOwnerId(v int64) *DescribeSavingsPlansCoverageTotalShrinkRequest {
	s.BillOwnerId = &v
	return s
}

func (s *DescribeSavingsPlansCoverageTotalShrinkRequest) SetEndPeriod(v string) *DescribeSavingsPlansCoverageTotalShrinkRequest {
	s.EndPeriod = &v
	return s
}

func (s *DescribeSavingsPlansCoverageTotalShrinkRequest) SetFilterParamShrink(v string) *DescribeSavingsPlansCoverageTotalShrinkRequest {
	s.FilterParamShrink = &v
	return s
}

func (s *DescribeSavingsPlansCoverageTotalShrinkRequest) SetPeriodType(v string) *DescribeSavingsPlansCoverageTotalShrinkRequest {
	s.PeriodType = &v
	return s
}

func (s *DescribeSavingsPlansCoverageTotalShrinkRequest) SetStartPeriod(v string) *DescribeSavingsPlansCoverageTotalShrinkRequest {
	s.StartPeriod = &v
	return s
}

func (s *DescribeSavingsPlansCoverageTotalShrinkRequest) Validate() error {
	return dara.Validate(s)
}
