// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSplitSummaryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCommonFilters(v []*ListSplitSummaryRequestCommonFilters) *ListSplitSummaryRequest
	GetCommonFilters() []*ListSplitSummaryRequestCommonFilters
	SetCostCategoriesFilters(v []*ListSplitSummaryRequestCostCategoriesFilters) *ListSplitSummaryRequest
	GetCostCategoriesFilters() []*ListSplitSummaryRequestCostCategoriesFilters
	SetEcIdAccountIds(v []*ListSplitSummaryRequestEcIdAccountIds) *ListSplitSummaryRequest
	GetEcIdAccountIds() []*ListSplitSummaryRequestEcIdAccountIds
	SetMaxResults(v int32) *ListSplitSummaryRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListSplitSummaryRequest
	GetNextToken() *string
	SetPeriod(v *ListSplitSummaryRequestPeriod) *ListSplitSummaryRequest
	GetPeriod() *ListSplitSummaryRequestPeriod
	SetTagFilters(v []*ListSplitSummaryRequestTagFilters) *ListSplitSummaryRequest
	GetTagFilters() []*ListSplitSummaryRequestTagFilters
}

type ListSplitSummaryRequest struct {
	// The common filter conditions. Supports filtering by product, commodity, region, billing item, fee type, resource, account, financial unit, resource group, and split item. Different fields are combined by AND. Only one object can be passed for the same field. For more information, see the supplementary description of request parameters.
	//
	// example:
	//
	// [{"FilterField":"SplitItemId","Operation":"IN","Values":["example-bucket"]}]
	CommonFilters []*ListSplitSummaryRequestCommonFilters `json:"CommonFilters,omitempty" xml:"CommonFilters,omitempty" type:"Repeated"`
	// The cost category filter conditions. Different objects in the list and multiple values of the same cost category are matched by OR. The overall cost category filter is combined with the tag filter and CommonFilters by AND.
	//
	// example:
	//
	// [{"CostCategoriesKey":"business","CostCategoriesValues":["commerce"]}]
	CostCategoriesFilters []*ListSplitSummaryRequestCostCategoriesFilters `json:"CostCategoriesFilters,omitempty" xml:"CostCategoriesFilters,omitempty" type:"Repeated"`
	// The query scope for enterprises and accounts. If not passed, the bill scope of the current account is queried. Only one enterprise object can be specified per request. The actual accessible scope is limited by permissions.
	//
	// example:
	//
	// [{"EcId":"123456","AccountIds":[1234567890123456]}]
	EcIdAccountIds []*ListSplitSummaryRequestEcIdAccountIds `json:"EcIdAccountIds,omitempty" xml:"EcIdAccountIds,omitempty" type:"Repeated"`
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
	Period *ListSplitSummaryRequestPeriod `json:"Period,omitempty" xml:"Period,omitempty" type:"Struct"`
	// The resource tag filter conditions. Different objects in the list and multiple values of the same tag key are matched by OR. The overall tag filter is combined with the cost category filter and CommonFilters by AND.
	//
	// example:
	//
	// [{"TagKey":"env","TagValues":["prod"]}]
	TagFilters []*ListSplitSummaryRequestTagFilters `json:"TagFilters,omitempty" xml:"TagFilters,omitempty" type:"Repeated"`
}

func (s ListSplitSummaryRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSplitSummaryRequest) GoString() string {
	return s.String()
}

func (s *ListSplitSummaryRequest) GetCommonFilters() []*ListSplitSummaryRequestCommonFilters {
	return s.CommonFilters
}

func (s *ListSplitSummaryRequest) GetCostCategoriesFilters() []*ListSplitSummaryRequestCostCategoriesFilters {
	return s.CostCategoriesFilters
}

func (s *ListSplitSummaryRequest) GetEcIdAccountIds() []*ListSplitSummaryRequestEcIdAccountIds {
	return s.EcIdAccountIds
}

func (s *ListSplitSummaryRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListSplitSummaryRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListSplitSummaryRequest) GetPeriod() *ListSplitSummaryRequestPeriod {
	return s.Period
}

func (s *ListSplitSummaryRequest) GetTagFilters() []*ListSplitSummaryRequestTagFilters {
	return s.TagFilters
}

func (s *ListSplitSummaryRequest) SetCommonFilters(v []*ListSplitSummaryRequestCommonFilters) *ListSplitSummaryRequest {
	s.CommonFilters = v
	return s
}

func (s *ListSplitSummaryRequest) SetCostCategoriesFilters(v []*ListSplitSummaryRequestCostCategoriesFilters) *ListSplitSummaryRequest {
	s.CostCategoriesFilters = v
	return s
}

func (s *ListSplitSummaryRequest) SetEcIdAccountIds(v []*ListSplitSummaryRequestEcIdAccountIds) *ListSplitSummaryRequest {
	s.EcIdAccountIds = v
	return s
}

func (s *ListSplitSummaryRequest) SetMaxResults(v int32) *ListSplitSummaryRequest {
	s.MaxResults = &v
	return s
}

func (s *ListSplitSummaryRequest) SetNextToken(v string) *ListSplitSummaryRequest {
	s.NextToken = &v
	return s
}

func (s *ListSplitSummaryRequest) SetPeriod(v *ListSplitSummaryRequestPeriod) *ListSplitSummaryRequest {
	s.Period = v
	return s
}

func (s *ListSplitSummaryRequest) SetTagFilters(v []*ListSplitSummaryRequestTagFilters) *ListSplitSummaryRequest {
	s.TagFilters = v
	return s
}

func (s *ListSplitSummaryRequest) Validate() error {
	if s.CommonFilters != nil {
		for _, item := range s.CommonFilters {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.CostCategoriesFilters != nil {
		for _, item := range s.CostCategoriesFilters {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.EcIdAccountIds != nil {
		for _, item := range s.EcIdAccountIds {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Period != nil {
		if err := s.Period.Validate(); err != nil {
			return err
		}
	}
	if s.TagFilters != nil {
		for _, item := range s.TagFilters {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListSplitSummaryRequestCommonFilters struct {
	// The name of the filter field. It is required when setting conditions. For supported fields and matching rules, see the supplementary description of request parameters.
	//
	// example:
	//
	// SplitItemId
	FilterField *string `json:"FilterField,omitempty" xml:"FilterField,omitempty"`
	// The matching method. This parameter is case-insensitive and defaults to IN if it is not specified or is empty. Valid values:
	//
	// - IN: includes
	//
	// - NOT: excludes
	//
	// - PREFIX: literal prefix match
	//
	// For the supported scope of each field, refer to the supplementary description of the request parameters.
	//
	// example:
	//
	// IN
	Operation *string `json:"Operation,omitempty" xml:"Operation,omitempty"`
	// The filter values. It is always an array of strings. Arrays are also used for single values, Boolean values, and account IDs.
	//
	// example:
	//
	// ["example-bucket"]
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s ListSplitSummaryRequestCommonFilters) String() string {
	return dara.Prettify(s)
}

func (s ListSplitSummaryRequestCommonFilters) GoString() string {
	return s.String()
}

func (s *ListSplitSummaryRequestCommonFilters) GetFilterField() *string {
	return s.FilterField
}

func (s *ListSplitSummaryRequestCommonFilters) GetOperation() *string {
	return s.Operation
}

func (s *ListSplitSummaryRequestCommonFilters) GetValues() []*string {
	return s.Values
}

func (s *ListSplitSummaryRequestCommonFilters) SetFilterField(v string) *ListSplitSummaryRequestCommonFilters {
	s.FilterField = &v
	return s
}

func (s *ListSplitSummaryRequestCommonFilters) SetOperation(v string) *ListSplitSummaryRequestCommonFilters {
	s.Operation = &v
	return s
}

func (s *ListSplitSummaryRequestCommonFilters) SetValues(v []*string) *ListSplitSummaryRequestCommonFilters {
	s.Values = v
	return s
}

func (s *ListSplitSummaryRequestCommonFilters) Validate() error {
	return dara.Validate(s)
}

type ListSplitSummaryRequestCostCategoriesFilters struct {
	// The cost category name. It cannot be empty or a string of only whitespace characters.
	//
	// This parameter is required.
	//
	// example:
	//
	// business
	CostCategoriesKey *string `json:"CostCategoriesKey,omitempty" xml:"CostCategoriesKey,omitempty"`
	// The list of cost category values. Matches any given value under the specified cost category. If this parameter is not passed or an empty array is passed, it only requires the existence of the cost category.
	//
	// example:
	//
	// ["commerce"]
	CostCategoriesValues []*string `json:"CostCategoriesValues,omitempty" xml:"CostCategoriesValues,omitempty" type:"Repeated"`
}

func (s ListSplitSummaryRequestCostCategoriesFilters) String() string {
	return dara.Prettify(s)
}

func (s ListSplitSummaryRequestCostCategoriesFilters) GoString() string {
	return s.String()
}

func (s *ListSplitSummaryRequestCostCategoriesFilters) GetCostCategoriesKey() *string {
	return s.CostCategoriesKey
}

func (s *ListSplitSummaryRequestCostCategoriesFilters) GetCostCategoriesValues() []*string {
	return s.CostCategoriesValues
}

func (s *ListSplitSummaryRequestCostCategoriesFilters) SetCostCategoriesKey(v string) *ListSplitSummaryRequestCostCategoriesFilters {
	s.CostCategoriesKey = &v
	return s
}

func (s *ListSplitSummaryRequestCostCategoriesFilters) SetCostCategoriesValues(v []*string) *ListSplitSummaryRequestCostCategoriesFilters {
	s.CostCategoriesValues = v
	return s
}

func (s *ListSplitSummaryRequestCostCategoriesFilters) Validate() error {
	return dara.Validate(s)
}

type ListSplitSummaryRequestEcIdAccountIds struct {
	// The list of account IDs to access. If it is empty, all accounts under the current enterprise entity are selected. The actual scope is limited by permissions.
	//
	// example:
	//
	// [1234567890123456]
	AccountIds []*int64 `json:"AccountIds,omitempty" xml:"AccountIds,omitempty" type:"Repeated"`
	// The enterprise entity ID.
	//
	// example:
	//
	// 123456
	EcId *string `json:"EcId,omitempty" xml:"EcId,omitempty"`
}

func (s ListSplitSummaryRequestEcIdAccountIds) String() string {
	return dara.Prettify(s)
}

func (s ListSplitSummaryRequestEcIdAccountIds) GoString() string {
	return s.String()
}

func (s *ListSplitSummaryRequestEcIdAccountIds) GetAccountIds() []*int64 {
	return s.AccountIds
}

func (s *ListSplitSummaryRequestEcIdAccountIds) GetEcId() *string {
	return s.EcId
}

func (s *ListSplitSummaryRequestEcIdAccountIds) SetAccountIds(v []*int64) *ListSplitSummaryRequestEcIdAccountIds {
	s.AccountIds = v
	return s
}

func (s *ListSplitSummaryRequestEcIdAccountIds) SetEcId(v string) *ListSplitSummaryRequestEcIdAccountIds {
	s.EcId = &v
	return s
}

func (s *ListSplitSummaryRequestEcIdAccountIds) Validate() error {
	return dara.Validate(s)
}

type ListSplitSummaryRequestPeriod struct {
	// The end time. It is inclusive. It is required for `DAILY` and the format is `yyyy-MM-dd`. It cannot be earlier than `StartTime` and must be in the same month as `StartTime`. It is optional for `MONTHLY`, but if passed, it must be the same as `StartTime`.
	//
	// example:
	//
	// 2026-09-07
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The time granularity. This parameter is required for business logic and is case-insensitive. Valid values:
	//
	// - DAILY: aggregates by day
	//
	// - MONTHLY: aggregates by month
	//
	// example:
	//
	// DAILY
	Granularity *string `json:"Granularity,omitempty" xml:"Granularity,omitempty"`
	// The start time. It is required for business logic and is inclusive. The format is yyyy-MM for MONTHLY and yyyy-MM-dd for DAILY.
	//
	// example:
	//
	// 2026-09-01
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ListSplitSummaryRequestPeriod) String() string {
	return dara.Prettify(s)
}

func (s ListSplitSummaryRequestPeriod) GoString() string {
	return s.String()
}

func (s *ListSplitSummaryRequestPeriod) GetEndTime() *string {
	return s.EndTime
}

func (s *ListSplitSummaryRequestPeriod) GetGranularity() *string {
	return s.Granularity
}

func (s *ListSplitSummaryRequestPeriod) GetStartTime() *string {
	return s.StartTime
}

func (s *ListSplitSummaryRequestPeriod) SetEndTime(v string) *ListSplitSummaryRequestPeriod {
	s.EndTime = &v
	return s
}

func (s *ListSplitSummaryRequestPeriod) SetGranularity(v string) *ListSplitSummaryRequestPeriod {
	s.Granularity = &v
	return s
}

func (s *ListSplitSummaryRequestPeriod) SetStartTime(v string) *ListSplitSummaryRequestPeriod {
	s.StartTime = &v
	return s
}

func (s *ListSplitSummaryRequestPeriod) Validate() error {
	return dara.Validate(s)
}

type ListSplitSummaryRequestTagFilters struct {
	// The tag key. It cannot be empty or a string of only whitespace characters.
	//
	// This parameter is required.
	//
	// example:
	//
	// env
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The list of tag values. Matches any given value under the specified tag key. If this parameter is not passed or an empty array is passed, it only requires the existence of the tag key.
	//
	// example:
	//
	// ["prod"]
	TagValues []*string `json:"TagValues,omitempty" xml:"TagValues,omitempty" type:"Repeated"`
}

func (s ListSplitSummaryRequestTagFilters) String() string {
	return dara.Prettify(s)
}

func (s ListSplitSummaryRequestTagFilters) GoString() string {
	return s.String()
}

func (s *ListSplitSummaryRequestTagFilters) GetTagKey() *string {
	return s.TagKey
}

func (s *ListSplitSummaryRequestTagFilters) GetTagValues() []*string {
	return s.TagValues
}

func (s *ListSplitSummaryRequestTagFilters) SetTagKey(v string) *ListSplitSummaryRequestTagFilters {
	s.TagKey = &v
	return s
}

func (s *ListSplitSummaryRequestTagFilters) SetTagValues(v []*string) *ListSplitSummaryRequestTagFilters {
	s.TagValues = v
	return s
}

func (s *ListSplitSummaryRequestTagFilters) Validate() error {
	return dara.Validate(s)
}
