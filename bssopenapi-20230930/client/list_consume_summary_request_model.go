// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListConsumeSummaryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCommonFilters(v []*ListConsumeSummaryRequestCommonFilters) *ListConsumeSummaryRequest
	GetCommonFilters() []*ListConsumeSummaryRequestCommonFilters
	SetDimensionType(v string) *ListConsumeSummaryRequest
	GetDimensionType() *string
	SetEcIdAccountIds(v []*ListConsumeSummaryRequestEcIdAccountIds) *ListConsumeSummaryRequest
	GetEcIdAccountIds() []*ListConsumeSummaryRequestEcIdAccountIds
	SetMaxResults(v int32) *ListConsumeSummaryRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListConsumeSummaryRequest
	GetNextToken() *string
	SetPeriod(v *ListConsumeSummaryRequestPeriod) *ListConsumeSummaryRequest
	GetPeriod() *ListConsumeSummaryRequestPeriod
}

type ListConsumeSummaryRequest struct {
	// The common filter conditions. Valid conditions for different fields take effect simultaneously. For supported fields and operations, see "Common filtering" in the supplementary description of request parameters.
	//
	// example:
	//
	// [{"FilterField":"ProductCode","Operation":"IN","Values":["ecs"]}]
	CommonFilters []*ListConsumeSummaryRequestCommonFilters `json:"CommonFilters,omitempty" xml:"CommonFilters,omitempty" type:"Repeated"`
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
	EcIdAccountIds []*ListConsumeSummaryRequestEcIdAccountIds `json:"EcIdAccountIds,omitempty" xml:"EcIdAccountIds,omitempty" type:"Repeated"`
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
	Period *ListConsumeSummaryRequestPeriod `json:"Period,omitempty" xml:"Period,omitempty" type:"Struct"`
}

func (s ListConsumeSummaryRequest) String() string {
	return dara.Prettify(s)
}

func (s ListConsumeSummaryRequest) GoString() string {
	return s.String()
}

func (s *ListConsumeSummaryRequest) GetCommonFilters() []*ListConsumeSummaryRequestCommonFilters {
	return s.CommonFilters
}

func (s *ListConsumeSummaryRequest) GetDimensionType() *string {
	return s.DimensionType
}

func (s *ListConsumeSummaryRequest) GetEcIdAccountIds() []*ListConsumeSummaryRequestEcIdAccountIds {
	return s.EcIdAccountIds
}

func (s *ListConsumeSummaryRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListConsumeSummaryRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListConsumeSummaryRequest) GetPeriod() *ListConsumeSummaryRequestPeriod {
	return s.Period
}

func (s *ListConsumeSummaryRequest) SetCommonFilters(v []*ListConsumeSummaryRequestCommonFilters) *ListConsumeSummaryRequest {
	s.CommonFilters = v
	return s
}

func (s *ListConsumeSummaryRequest) SetDimensionType(v string) *ListConsumeSummaryRequest {
	s.DimensionType = &v
	return s
}

func (s *ListConsumeSummaryRequest) SetEcIdAccountIds(v []*ListConsumeSummaryRequestEcIdAccountIds) *ListConsumeSummaryRequest {
	s.EcIdAccountIds = v
	return s
}

func (s *ListConsumeSummaryRequest) SetMaxResults(v int32) *ListConsumeSummaryRequest {
	s.MaxResults = &v
	return s
}

func (s *ListConsumeSummaryRequest) SetNextToken(v string) *ListConsumeSummaryRequest {
	s.NextToken = &v
	return s
}

func (s *ListConsumeSummaryRequest) SetPeriod(v *ListConsumeSummaryRequestPeriod) *ListConsumeSummaryRequest {
	s.Period = v
	return s
}

func (s *ListConsumeSummaryRequest) Validate() error {
	if s.CommonFilters != nil {
		for _, item := range s.CommonFilters {
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
	return nil
}

type ListConsumeSummaryRequestCommonFilters struct {
	// The name of the filter field. For supported fields and operations, see "Common filtering" in the supplementary description of request parameters.
	//
	// example:
	//
	// ProductCode
	FilterField *string `json:"FilterField,omitempty" xml:"FilterField,omitempty"`
	// The operator. Default value: `IN`. `IN` indicates inclusion or exact match, `NOT` indicates exclusion, and `PREFIX` indicates prefix match. For the specific support scope, see "Common filtering" in the supplementary description of request parameters.
	//
	// example:
	//
	// IN
	Operation *string `json:"Operation,omitempty" xml:"Operation,omitempty"`
	// The filter values always use a string array, including single values, Boolean values, and account IDs.
	//
	// example:
	//
	// ["ecs"]
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s ListConsumeSummaryRequestCommonFilters) String() string {
	return dara.Prettify(s)
}

func (s ListConsumeSummaryRequestCommonFilters) GoString() string {
	return s.String()
}

func (s *ListConsumeSummaryRequestCommonFilters) GetFilterField() *string {
	return s.FilterField
}

func (s *ListConsumeSummaryRequestCommonFilters) GetOperation() *string {
	return s.Operation
}

func (s *ListConsumeSummaryRequestCommonFilters) GetValues() []*string {
	return s.Values
}

func (s *ListConsumeSummaryRequestCommonFilters) SetFilterField(v string) *ListConsumeSummaryRequestCommonFilters {
	s.FilterField = &v
	return s
}

func (s *ListConsumeSummaryRequestCommonFilters) SetOperation(v string) *ListConsumeSummaryRequestCommonFilters {
	s.Operation = &v
	return s
}

func (s *ListConsumeSummaryRequestCommonFilters) SetValues(v []*string) *ListConsumeSummaryRequestCommonFilters {
	s.Values = v
	return s
}

func (s *ListConsumeSummaryRequestCommonFilters) Validate() error {
	return dara.Validate(s)
}

type ListConsumeSummaryRequestEcIdAccountIds struct {
	// The list of account IDs to query. This parameter is required when an enterprise object is passed. If the entire EcIdAccountIds parameter is omitted, the data of the current calling account is queried by default.
	//
	// example:
	//
	// [1234567890123456]
	AccountIds []*int64 `json:"AccountIds,omitempty" xml:"AccountIds,omitempty" type:"Repeated"`
	// The ID of the enterprise entity.
	//
	// example:
	//
	// 123456
	EcId *string `json:"EcId,omitempty" xml:"EcId,omitempty"`
}

func (s ListConsumeSummaryRequestEcIdAccountIds) String() string {
	return dara.Prettify(s)
}

func (s ListConsumeSummaryRequestEcIdAccountIds) GoString() string {
	return s.String()
}

func (s *ListConsumeSummaryRequestEcIdAccountIds) GetAccountIds() []*int64 {
	return s.AccountIds
}

func (s *ListConsumeSummaryRequestEcIdAccountIds) GetEcId() *string {
	return s.EcId
}

func (s *ListConsumeSummaryRequestEcIdAccountIds) SetAccountIds(v []*int64) *ListConsumeSummaryRequestEcIdAccountIds {
	s.AccountIds = v
	return s
}

func (s *ListConsumeSummaryRequestEcIdAccountIds) SetEcId(v string) *ListConsumeSummaryRequestEcIdAccountIds {
	s.EcId = &v
	return s
}

func (s *ListConsumeSummaryRequestEcIdAccountIds) Validate() error {
	return dara.Validate(s)
}

type ListConsumeSummaryRequestPeriod struct {
	// For monthly queries, only one month is supported at a time. The value must be the same as the start month, and the format is `yyyy-MM`.
	//
	// For daily queries, this parameter is required. The format is `yyyy-MM-dd`. The end date cannot be earlier than the start date, and the start and end dates must be in the same month.
	//
	// example:
	//
	// 2026-10-07
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// `MONTHLY`: aggregates by month. `DAILY`: aggregates by day.
	//
	// example:
	//
	// DAILY
	Granularity *string `json:"Granularity,omitempty" xml:"Granularity,omitempty"`
	// For monthly queries, the format is `yyyy-MM`.
	//
	// For daily queries, the format is `yyyy-MM-dd`.
	//
	// example:
	//
	// 2026-10-01
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ListConsumeSummaryRequestPeriod) String() string {
	return dara.Prettify(s)
}

func (s ListConsumeSummaryRequestPeriod) GoString() string {
	return s.String()
}

func (s *ListConsumeSummaryRequestPeriod) GetEndTime() *string {
	return s.EndTime
}

func (s *ListConsumeSummaryRequestPeriod) GetGranularity() *string {
	return s.Granularity
}

func (s *ListConsumeSummaryRequestPeriod) GetStartTime() *string {
	return s.StartTime
}

func (s *ListConsumeSummaryRequestPeriod) SetEndTime(v string) *ListConsumeSummaryRequestPeriod {
	s.EndTime = &v
	return s
}

func (s *ListConsumeSummaryRequestPeriod) SetGranularity(v string) *ListConsumeSummaryRequestPeriod {
	s.Granularity = &v
	return s
}

func (s *ListConsumeSummaryRequestPeriod) SetStartTime(v string) *ListConsumeSummaryRequestPeriod {
	s.StartTime = &v
	return s
}

func (s *ListConsumeSummaryRequestPeriod) Validate() error {
	return dara.Validate(s)
}
