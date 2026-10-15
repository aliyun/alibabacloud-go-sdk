// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSavingsPlansCoverageDetailShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBillOwnerId(v int64) *DescribeSavingsPlansCoverageDetailShrinkRequest
	GetBillOwnerId() *int64
	SetEndPeriod(v string) *DescribeSavingsPlansCoverageDetailShrinkRequest
	GetEndPeriod() *string
	SetFilterParamShrink(v string) *DescribeSavingsPlansCoverageDetailShrinkRequest
	GetFilterParamShrink() *string
	SetMaxResults(v int32) *DescribeSavingsPlansCoverageDetailShrinkRequest
	GetMaxResults() *int32
	SetPeriodType(v string) *DescribeSavingsPlansCoverageDetailShrinkRequest
	GetPeriodType() *string
	SetStartPeriod(v string) *DescribeSavingsPlansCoverageDetailShrinkRequest
	GetStartPeriod() *string
	SetToken(v string) *DescribeSavingsPlansCoverageDetailShrinkRequest
	GetToken() *string
}

type DescribeSavingsPlansCoverageDetailShrinkRequest struct {
	// The ID of the resource ownership account. The resource ownership account is the account that actually uses the resource.
	//
	// example:
	//
	// 123745698925000
	BillOwnerId *int64 `json:"BillOwnerId,omitempty" xml:"BillOwnerId,omitempty"`
	// The end of the time range to query. The end time is excluded from the time range. If you do not set this parameter, the end time is the current time. Format: yyyy-MM-dd HH:mm:ss.
	//
	// example:
	//
	// 2021-01-01 00:00:00
	EndPeriod         *string `json:"EndPeriod,omitempty" xml:"EndPeriod,omitempty"`
	FilterParamShrink *string `json:"FilterParam,omitempty" xml:"FilterParam,omitempty"`
	// The maximum number of entries to return.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The time granularity for querying the usage summary. Valid values: MONTH, DAY, HOUR.
	//
	// This parameter is required.
	//
	// example:
	//
	// HOUR
	PeriodType *string `json:"PeriodType,omitempty" xml:"PeriodType,omitempty"`
	// The beginning of the time range to query. The start time is included in the time range. Format: yyyy-MM-dd HH:mm:ss.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2021-01-01 00:00:00
	StartPeriod *string `json:"StartPeriod,omitempty" xml:"StartPeriod,omitempty"`
	// The pagination token. Leave this parameter empty for the first page. The response returns the token for the next page. An empty token in the response indicates that all data has been returned.
	//
	// example:
	//
	// eyJwYWdlTnVtIjoyLCJwYWdlU2l6ZSI6MTB9
	Token *string `json:"Token,omitempty" xml:"Token,omitempty"`
}

func (s DescribeSavingsPlansCoverageDetailShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSavingsPlansCoverageDetailShrinkRequest) GoString() string {
	return s.String()
}

func (s *DescribeSavingsPlansCoverageDetailShrinkRequest) GetBillOwnerId() *int64 {
	return s.BillOwnerId
}

func (s *DescribeSavingsPlansCoverageDetailShrinkRequest) GetEndPeriod() *string {
	return s.EndPeriod
}

func (s *DescribeSavingsPlansCoverageDetailShrinkRequest) GetFilterParamShrink() *string {
	return s.FilterParamShrink
}

func (s *DescribeSavingsPlansCoverageDetailShrinkRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeSavingsPlansCoverageDetailShrinkRequest) GetPeriodType() *string {
	return s.PeriodType
}

func (s *DescribeSavingsPlansCoverageDetailShrinkRequest) GetStartPeriod() *string {
	return s.StartPeriod
}

func (s *DescribeSavingsPlansCoverageDetailShrinkRequest) GetToken() *string {
	return s.Token
}

func (s *DescribeSavingsPlansCoverageDetailShrinkRequest) SetBillOwnerId(v int64) *DescribeSavingsPlansCoverageDetailShrinkRequest {
	s.BillOwnerId = &v
	return s
}

func (s *DescribeSavingsPlansCoverageDetailShrinkRequest) SetEndPeriod(v string) *DescribeSavingsPlansCoverageDetailShrinkRequest {
	s.EndPeriod = &v
	return s
}

func (s *DescribeSavingsPlansCoverageDetailShrinkRequest) SetFilterParamShrink(v string) *DescribeSavingsPlansCoverageDetailShrinkRequest {
	s.FilterParamShrink = &v
	return s
}

func (s *DescribeSavingsPlansCoverageDetailShrinkRequest) SetMaxResults(v int32) *DescribeSavingsPlansCoverageDetailShrinkRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeSavingsPlansCoverageDetailShrinkRequest) SetPeriodType(v string) *DescribeSavingsPlansCoverageDetailShrinkRequest {
	s.PeriodType = &v
	return s
}

func (s *DescribeSavingsPlansCoverageDetailShrinkRequest) SetStartPeriod(v string) *DescribeSavingsPlansCoverageDetailShrinkRequest {
	s.StartPeriod = &v
	return s
}

func (s *DescribeSavingsPlansCoverageDetailShrinkRequest) SetToken(v string) *DescribeSavingsPlansCoverageDetailShrinkRequest {
	s.Token = &v
	return s
}

func (s *DescribeSavingsPlansCoverageDetailShrinkRequest) Validate() error {
	return dara.Validate(s)
}
