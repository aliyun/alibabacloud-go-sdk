// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSavingsPlansCoverageTotalRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBillOwnerId(v int64) *DescribeSavingsPlansCoverageTotalRequest
	GetBillOwnerId() *int64
	SetEndPeriod(v string) *DescribeSavingsPlansCoverageTotalRequest
	GetEndPeriod() *string
	SetPeriodType(v string) *DescribeSavingsPlansCoverageTotalRequest
	GetPeriodType() *string
	SetStartPeriod(v string) *DescribeSavingsPlansCoverageTotalRequest
	GetStartPeriod() *string
}

type DescribeSavingsPlansCoverageTotalRequest struct {
	// The ID of the account for which you want to query coverage summary. If you do not set this parameter, the data of the current Alibaba Cloud account and its RAM users is queried. To query the data of a RAM user, specify the ID of the RAM user.
	//
	// example:
	//
	// 123745698925000
	BillOwnerId *int64 `json:"BillOwnerId,omitempty" xml:"BillOwnerId,omitempty"`
	// The end of the time range to query. The end is excluded from the time range. If you do not set this parameter, the end time is the current time. Specify the time in the format of yyyy-MM-dd HH:mm:ss.
	//
	// example:
	//
	// 2021-07-20 00:00:00
	EndPeriod *string `json:"EndPeriod,omitempty" xml:"EndPeriod,omitempty"`
	// The time granularity at which coverage summary are queried. Valid values: MONTH, DAY, and HOUR.
	//
	// This parameter is required.
	//
	// example:
	//
	// HOUR
	PeriodType *string `json:"PeriodType,omitempty" xml:"PeriodType,omitempty"`
	// The beginning of the time range to query. The beginning is included in the time range. Specify the time in the format of yyyy-MM-dd HH:mm:ss.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2021-07-15 00:00:00
	StartPeriod *string `json:"StartPeriod,omitempty" xml:"StartPeriod,omitempty"`
}

func (s DescribeSavingsPlansCoverageTotalRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSavingsPlansCoverageTotalRequest) GoString() string {
	return s.String()
}

func (s *DescribeSavingsPlansCoverageTotalRequest) GetBillOwnerId() *int64 {
	return s.BillOwnerId
}

func (s *DescribeSavingsPlansCoverageTotalRequest) GetEndPeriod() *string {
	return s.EndPeriod
}

func (s *DescribeSavingsPlansCoverageTotalRequest) GetPeriodType() *string {
	return s.PeriodType
}

func (s *DescribeSavingsPlansCoverageTotalRequest) GetStartPeriod() *string {
	return s.StartPeriod
}

func (s *DescribeSavingsPlansCoverageTotalRequest) SetBillOwnerId(v int64) *DescribeSavingsPlansCoverageTotalRequest {
	s.BillOwnerId = &v
	return s
}

func (s *DescribeSavingsPlansCoverageTotalRequest) SetEndPeriod(v string) *DescribeSavingsPlansCoverageTotalRequest {
	s.EndPeriod = &v
	return s
}

func (s *DescribeSavingsPlansCoverageTotalRequest) SetPeriodType(v string) *DescribeSavingsPlansCoverageTotalRequest {
	s.PeriodType = &v
	return s
}

func (s *DescribeSavingsPlansCoverageTotalRequest) SetStartPeriod(v string) *DescribeSavingsPlansCoverageTotalRequest {
	s.StartPeriod = &v
	return s
}

func (s *DescribeSavingsPlansCoverageTotalRequest) Validate() error {
	return dara.Validate(s)
}
