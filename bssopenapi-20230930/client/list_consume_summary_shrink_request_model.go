// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListConsumeSummaryShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCommonFiltersShrink(v string) *ListConsumeSummaryShrinkRequest
	GetCommonFiltersShrink() *string
	SetDimensionType(v string) *ListConsumeSummaryShrinkRequest
	GetDimensionType() *string
	SetEcIdAccountIdsShrink(v string) *ListConsumeSummaryShrinkRequest
	GetEcIdAccountIdsShrink() *string
	SetMaxResults(v int32) *ListConsumeSummaryShrinkRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListConsumeSummaryShrinkRequest
	GetNextToken() *string
	SetPeriodShrink(v string) *ListConsumeSummaryShrinkRequest
	GetPeriodShrink() *string
}

type ListConsumeSummaryShrinkRequest struct {
	// The common filter conditions. Valid conditions for different fields take effect simultaneously. For supported fields and operations, see "Common filtering" in the supplementary description of request parameters.
	//
	// example:
	//
	// [{"FilterField":"ProductCode","Operation":"IN","Values":["ecs"]}]
	CommonFiltersShrink *string `json:"CommonFilters,omitempty" xml:"CommonFilters,omitempty"`
	// The aggregation dimension.
	//
	// `BILLING_ITEM`: aggregates by billing item. This is the default value.
	//
	// `RESOURCE`: aggregates by resource.
	//
	// `PRODUCT`: aggregates by product.
	//
	// `ACCOUNT`: aggregates by account.
	//
	// example:
	//
	// BILLING_ITEM
	DimensionType *string `json:"DimensionType,omitempty" xml:"DimensionType,omitempty"`
	// The query scope for enterprises and accounts. If this parameter is not specified, the default billing scope of the current calling account is used. Only one enterprise object can be specified per request, and the actual scope is limited by permissions.
	//
	// example:
	//
	// [{"EcId":"123456","AccountIds":[1234567890123456]}]
	EcIdAccountIdsShrink *string `json:"EcIdAccountIds,omitempty" xml:"EcIdAccountIds,omitempty"`
	// The maximum number of records expected to be returned for a single query. Specify a positive integer. Keep this value unchanged during pagination. The actual number of returned records may be less than this value.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The query token for the next page. Do not specify this parameter for the first request. For subsequent requests, pass the value returned in the previous response.
	//
	// example:
	//
	// example-next-token
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The query time range, which specifies the time granularity, start time, and end time.
	//
	// example:
	//
	// {"Granularity":"DAILY","StartTime":"2026-10-01","EndTime":"2026-10-07"}
	PeriodShrink *string `json:"Period,omitempty" xml:"Period,omitempty"`
}

func (s ListConsumeSummaryShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListConsumeSummaryShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListConsumeSummaryShrinkRequest) GetCommonFiltersShrink() *string {
	return s.CommonFiltersShrink
}

func (s *ListConsumeSummaryShrinkRequest) GetDimensionType() *string {
	return s.DimensionType
}

func (s *ListConsumeSummaryShrinkRequest) GetEcIdAccountIdsShrink() *string {
	return s.EcIdAccountIdsShrink
}

func (s *ListConsumeSummaryShrinkRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListConsumeSummaryShrinkRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListConsumeSummaryShrinkRequest) GetPeriodShrink() *string {
	return s.PeriodShrink
}

func (s *ListConsumeSummaryShrinkRequest) SetCommonFiltersShrink(v string) *ListConsumeSummaryShrinkRequest {
	s.CommonFiltersShrink = &v
	return s
}

func (s *ListConsumeSummaryShrinkRequest) SetDimensionType(v string) *ListConsumeSummaryShrinkRequest {
	s.DimensionType = &v
	return s
}

func (s *ListConsumeSummaryShrinkRequest) SetEcIdAccountIdsShrink(v string) *ListConsumeSummaryShrinkRequest {
	s.EcIdAccountIdsShrink = &v
	return s
}

func (s *ListConsumeSummaryShrinkRequest) SetMaxResults(v int32) *ListConsumeSummaryShrinkRequest {
	s.MaxResults = &v
	return s
}

func (s *ListConsumeSummaryShrinkRequest) SetNextToken(v string) *ListConsumeSummaryShrinkRequest {
	s.NextToken = &v
	return s
}

func (s *ListConsumeSummaryShrinkRequest) SetPeriodShrink(v string) *ListConsumeSummaryShrinkRequest {
	s.PeriodShrink = &v
	return s
}

func (s *ListConsumeSummaryShrinkRequest) Validate() error {
	return dara.Validate(s)
}
