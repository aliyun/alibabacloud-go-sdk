// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSavingsPlansUsageDetailShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBillOwnerId(v int64) *DescribeSavingsPlansUsageDetailShrinkRequest
	GetBillOwnerId() *int64
	SetEndPeriod(v string) *DescribeSavingsPlansUsageDetailShrinkRequest
	GetEndPeriod() *string
	SetFilterParamShrink(v string) *DescribeSavingsPlansUsageDetailShrinkRequest
	GetFilterParamShrink() *string
	SetMaxResults(v int32) *DescribeSavingsPlansUsageDetailShrinkRequest
	GetMaxResults() *int32
	SetPeriodType(v string) *DescribeSavingsPlansUsageDetailShrinkRequest
	GetPeriodType() *string
	SetStartPeriod(v string) *DescribeSavingsPlansUsageDetailShrinkRequest
	GetStartPeriod() *string
	SetToken(v string) *DescribeSavingsPlansUsageDetailShrinkRequest
	GetToken() *string
}

type DescribeSavingsPlansUsageDetailShrinkRequest struct {
	// The ID of the account whose data you want to query. If this parameter is left empty, the data of the current account and its trusteeship linked accounts is queried. You can specify the ID of a trusteeship linked account to query only the data of that account.
	//
	// example:
	//
	// 123745698925000
	BillOwnerId *int64 `json:"BillOwnerId,omitempty" xml:"BillOwnerId,omitempty"`
	// The end time of the query. Format: yyyy-MM-dd HH:mm:ss. The specified time is excluded from the query. If this parameter is left empty, the current time is used.
	//
	// example:
	//
	// 2021-01-02 00:00:00
	EndPeriod *string `json:"EndPeriod,omitempty" xml:"EndPeriod,omitempty"`
	// The filter parameters.
	FilterParamShrink *string `json:"FilterParam,omitempty" xml:"FilterParam,omitempty"`
	// The maximum number of data records to return in the current request. Default value: 20. Maximum value: 300.
	//
	// example:
	//
	// 300
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The time granularity for querying the usage summary. Valid values:
	//
	// - MONTH: monthly
	//
	// - DAY: daily
	//
	// - HOUR: hourly
	//
	// This parameter is required.
	//
	// example:
	//
	// HOUR
	PeriodType *string `json:"PeriodType,omitempty" xml:"PeriodType,omitempty"`
	// The start time of the query. Format: yyyy-MM-dd HH:mm:ss. The specified time is included in the query.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2021-01-01 00:00:00
	StartPeriod *string `json:"StartPeriod,omitempty" xml:"StartPeriod,omitempty"`
	// The pagination token. Leave this parameter empty when you query the first page. The token for the next page is returned in the response. If the returned token is empty, all results are returned.
	//
	// example:
	//
	// eyJwYWdlTnVtIjoyLCJwYWdlU2l6ZSI6MTB9
	Token *string `json:"Token,omitempty" xml:"Token,omitempty"`
}

func (s DescribeSavingsPlansUsageDetailShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSavingsPlansUsageDetailShrinkRequest) GoString() string {
	return s.String()
}

func (s *DescribeSavingsPlansUsageDetailShrinkRequest) GetBillOwnerId() *int64 {
	return s.BillOwnerId
}

func (s *DescribeSavingsPlansUsageDetailShrinkRequest) GetEndPeriod() *string {
	return s.EndPeriod
}

func (s *DescribeSavingsPlansUsageDetailShrinkRequest) GetFilterParamShrink() *string {
	return s.FilterParamShrink
}

func (s *DescribeSavingsPlansUsageDetailShrinkRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeSavingsPlansUsageDetailShrinkRequest) GetPeriodType() *string {
	return s.PeriodType
}

func (s *DescribeSavingsPlansUsageDetailShrinkRequest) GetStartPeriod() *string {
	return s.StartPeriod
}

func (s *DescribeSavingsPlansUsageDetailShrinkRequest) GetToken() *string {
	return s.Token
}

func (s *DescribeSavingsPlansUsageDetailShrinkRequest) SetBillOwnerId(v int64) *DescribeSavingsPlansUsageDetailShrinkRequest {
	s.BillOwnerId = &v
	return s
}

func (s *DescribeSavingsPlansUsageDetailShrinkRequest) SetEndPeriod(v string) *DescribeSavingsPlansUsageDetailShrinkRequest {
	s.EndPeriod = &v
	return s
}

func (s *DescribeSavingsPlansUsageDetailShrinkRequest) SetFilterParamShrink(v string) *DescribeSavingsPlansUsageDetailShrinkRequest {
	s.FilterParamShrink = &v
	return s
}

func (s *DescribeSavingsPlansUsageDetailShrinkRequest) SetMaxResults(v int32) *DescribeSavingsPlansUsageDetailShrinkRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeSavingsPlansUsageDetailShrinkRequest) SetPeriodType(v string) *DescribeSavingsPlansUsageDetailShrinkRequest {
	s.PeriodType = &v
	return s
}

func (s *DescribeSavingsPlansUsageDetailShrinkRequest) SetStartPeriod(v string) *DescribeSavingsPlansUsageDetailShrinkRequest {
	s.StartPeriod = &v
	return s
}

func (s *DescribeSavingsPlansUsageDetailShrinkRequest) SetToken(v string) *DescribeSavingsPlansUsageDetailShrinkRequest {
	s.Token = &v
	return s
}

func (s *DescribeSavingsPlansUsageDetailShrinkRequest) Validate() error {
	return dara.Validate(s)
}
