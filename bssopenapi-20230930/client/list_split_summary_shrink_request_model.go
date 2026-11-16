// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSplitSummaryShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCommonFiltersShrink(v string) *ListSplitSummaryShrinkRequest
	GetCommonFiltersShrink() *string
	SetCostCategoriesFiltersShrink(v string) *ListSplitSummaryShrinkRequest
	GetCostCategoriesFiltersShrink() *string
	SetEcIdAccountIdsShrink(v string) *ListSplitSummaryShrinkRequest
	GetEcIdAccountIdsShrink() *string
	SetMaxResults(v int32) *ListSplitSummaryShrinkRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListSplitSummaryShrinkRequest
	GetNextToken() *string
	SetPeriodShrink(v string) *ListSplitSummaryShrinkRequest
	GetPeriodShrink() *string
	SetTagFiltersShrink(v string) *ListSplitSummaryShrinkRequest
	GetTagFiltersShrink() *string
}

type ListSplitSummaryShrinkRequest struct {
	// The common filter conditions. Supports filtering by product, commodity, region, billing item, fee type, resource, account, financial unit, resource group, and split item. Different fields are combined by AND. Only one object can be passed for the same field. For more information, see the supplementary description of request parameters.
	//
	// example:
	//
	// [{"FilterField":"SplitItemId","Operation":"IN","Values":["example-bucket"]}]
	CommonFiltersShrink *string `json:"CommonFilters,omitempty" xml:"CommonFilters,omitempty"`
	// The cost category filter conditions. Different objects in the list and multiple values of the same cost category are matched by OR. The overall cost category filter is combined with the tag filter and CommonFilters by AND.
	//
	// example:
	//
	// [{"CostCategoriesKey":"business","CostCategoriesValues":["commerce"]}]
	CostCategoriesFiltersShrink *string `json:"CostCategoriesFilters,omitempty" xml:"CostCategoriesFilters,omitempty"`
	// The query scope for enterprises and accounts. If not passed, the bill scope of the current account is queried. Only one enterprise object can be specified per request. The actual accessible scope is limited by permissions.
	//
	// example:
	//
	// [{"EcId":"123456","AccountIds":[1234567890123456]}]
	EcIdAccountIdsShrink *string `json:"EcIdAccountIds,omitempty" xml:"EcIdAccountIds,omitempty"`
	// The maximum number of records expected to be returned per page. It is used in conjunction with `NextToken`. Pass a positive integer. We recommend that you explicitly set this parameter and keep it unchanged during pagination. The actual number of returned records may be less than this value.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token for the next page of query results. Do not pass this parameter for the first query. For subsequent queries, pass the `NextToken` returned in the previous response as is. The query ends when the token is empty or no longer returned. Do not construct or parse the token yourself. Keep `Period`, account scope, all filter conditions, and `MaxResults` consistent during pagination.
	//
	// example:
	//
	// opaque-token-example
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The query time range. It is required for business logic and includes Granularity, StartTime, and EndTime. Only one billing month can be queried per request. If passed, it overrides the legacy flat time parameters.
	//
	// example:
	//
	// {"Granularity":"DAILY","StartTime":"2026-09-01","EndTime":"2026-09-07"}
	PeriodShrink *string `json:"Period,omitempty" xml:"Period,omitempty"`
	// The resource tag filter conditions. Different objects in the list and multiple values of the same tag key are matched by OR. The overall tag filter is combined with the cost category filter and CommonFilters by AND.
	//
	// example:
	//
	// [{"TagKey":"env","TagValues":["prod"]}]
	TagFiltersShrink *string `json:"TagFilters,omitempty" xml:"TagFilters,omitempty"`
}

func (s ListSplitSummaryShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSplitSummaryShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListSplitSummaryShrinkRequest) GetCommonFiltersShrink() *string {
	return s.CommonFiltersShrink
}

func (s *ListSplitSummaryShrinkRequest) GetCostCategoriesFiltersShrink() *string {
	return s.CostCategoriesFiltersShrink
}

func (s *ListSplitSummaryShrinkRequest) GetEcIdAccountIdsShrink() *string {
	return s.EcIdAccountIdsShrink
}

func (s *ListSplitSummaryShrinkRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListSplitSummaryShrinkRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListSplitSummaryShrinkRequest) GetPeriodShrink() *string {
	return s.PeriodShrink
}

func (s *ListSplitSummaryShrinkRequest) GetTagFiltersShrink() *string {
	return s.TagFiltersShrink
}

func (s *ListSplitSummaryShrinkRequest) SetCommonFiltersShrink(v string) *ListSplitSummaryShrinkRequest {
	s.CommonFiltersShrink = &v
	return s
}

func (s *ListSplitSummaryShrinkRequest) SetCostCategoriesFiltersShrink(v string) *ListSplitSummaryShrinkRequest {
	s.CostCategoriesFiltersShrink = &v
	return s
}

func (s *ListSplitSummaryShrinkRequest) SetEcIdAccountIdsShrink(v string) *ListSplitSummaryShrinkRequest {
	s.EcIdAccountIdsShrink = &v
	return s
}

func (s *ListSplitSummaryShrinkRequest) SetMaxResults(v int32) *ListSplitSummaryShrinkRequest {
	s.MaxResults = &v
	return s
}

func (s *ListSplitSummaryShrinkRequest) SetNextToken(v string) *ListSplitSummaryShrinkRequest {
	s.NextToken = &v
	return s
}

func (s *ListSplitSummaryShrinkRequest) SetPeriodShrink(v string) *ListSplitSummaryShrinkRequest {
	s.PeriodShrink = &v
	return s
}

func (s *ListSplitSummaryShrinkRequest) SetTagFiltersShrink(v string) *ListSplitSummaryShrinkRequest {
	s.TagFiltersShrink = &v
	return s
}

func (s *ListSplitSummaryShrinkRequest) Validate() error {
	return dara.Validate(s)
}
