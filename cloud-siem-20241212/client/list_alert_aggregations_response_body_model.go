// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAlertAggregationsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAlertAggregations(v []*ListAlertAggregationsResponseBodyAlertAggregations) *ListAlertAggregationsResponseBody
	GetAlertAggregations() []*ListAlertAggregationsResponseBodyAlertAggregations
	SetMaxResults(v int32) *ListAlertAggregationsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListAlertAggregationsResponseBody
	GetNextToken() *string
	SetPageNumber(v int32) *ListAlertAggregationsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListAlertAggregationsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListAlertAggregationsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListAlertAggregationsResponseBody
	GetTotalCount() *int32
}

type ListAlertAggregationsResponseBody struct {
	AlertAggregations []*ListAlertAggregationsResponseBodyAlertAggregations `json:"AlertAggregations,omitempty" xml:"AlertAggregations,omitempty" type:"Repeated"`
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// AAAAAUqcj6VO4E3ECWIrFczs****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 9AAA9ED9-78F4-5021-86DC-D51C7511****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 57
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListAlertAggregationsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAlertAggregationsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAlertAggregationsResponseBody) GetAlertAggregations() []*ListAlertAggregationsResponseBodyAlertAggregations {
	return s.AlertAggregations
}

func (s *ListAlertAggregationsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListAlertAggregationsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAlertAggregationsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListAlertAggregationsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAlertAggregationsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAlertAggregationsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListAlertAggregationsResponseBody) SetAlertAggregations(v []*ListAlertAggregationsResponseBodyAlertAggregations) *ListAlertAggregationsResponseBody {
	s.AlertAggregations = v
	return s
}

func (s *ListAlertAggregationsResponseBody) SetMaxResults(v int32) *ListAlertAggregationsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListAlertAggregationsResponseBody) SetNextToken(v string) *ListAlertAggregationsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListAlertAggregationsResponseBody) SetPageNumber(v int32) *ListAlertAggregationsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListAlertAggregationsResponseBody) SetPageSize(v int32) *ListAlertAggregationsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListAlertAggregationsResponseBody) SetRequestId(v string) *ListAlertAggregationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAlertAggregationsResponseBody) SetTotalCount(v int32) *ListAlertAggregationsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListAlertAggregationsResponseBody) Validate() error {
	if s.AlertAggregations != nil {
		for _, item := range s.AlertAggregations {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAlertAggregationsResponseBodyAlertAggregations struct {
	AccountDisplayValues []*string `json:"AccountDisplayValues,omitempty" xml:"AccountDisplayValues,omitempty" type:"Repeated"`
	AccountIds           []*string `json:"AccountIds,omitempty" xml:"AccountIds,omitempty" type:"Repeated"`
	// example:
	//
	// vendor_product:alibaba_cloud:sas
	AggregationKey *string `json:"AggregationKey,omitempty" xml:"AggregationKey,omitempty"`
	// example:
	//
	// 5m-avg
	AggregationType *string `json:"AggregationType,omitempty" xml:"AggregationType,omitempty"`
	// example:
	//
	// 25
	AlertCount          *int64                                                                   `json:"AlertCount,omitempty" xml:"AlertCount,omitempty"`
	AttackStages        []*string                                                                `json:"AttackStages,omitempty" xml:"AttackStages,omitempty" type:"Repeated"`
	DefenseActionCounts []*ListAlertAggregationsResponseBodyAlertAggregationsDefenseActionCounts `json:"DefenseActionCounts,omitempty" xml:"DefenseActionCounts,omitempty" type:"Repeated"`
	// example:
	//
	// 1785290308
	FirstTime *int64 `json:"FirstTime,omitempty" xml:"FirstTime,omitempty"`
	// example:
	//
	// 1785293908000
	LatestTime *int64 `json:"LatestTime,omitempty" xml:"LatestTime,omitempty"`
	// example:
	//
	// Test incident
	PrimaryDisplayValue *string `json:"PrimaryDisplayValue,omitempty" xml:"PrimaryDisplayValue,omitempty"`
	// example:
	//
	// alibaba_cloud
	PrimaryValue *string `json:"PrimaryValue,omitempty" xml:"PrimaryValue,omitempty"`
	// example:
	//
	// sas
	SecondaryValue *string                                                           `json:"SecondaryValue,omitempty" xml:"SecondaryValue,omitempty"`
	SourceCodes    []*string                                                         `json:"SourceCodes,omitempty" xml:"SourceCodes,omitempty" type:"Repeated"`
	StatusCounts   []*ListAlertAggregationsResponseBodyAlertAggregationsStatusCounts `json:"StatusCounts,omitempty" xml:"StatusCounts,omitempty" type:"Repeated"`
}

func (s ListAlertAggregationsResponseBodyAlertAggregations) String() string {
	return dara.Prettify(s)
}

func (s ListAlertAggregationsResponseBodyAlertAggregations) GoString() string {
	return s.String()
}

func (s *ListAlertAggregationsResponseBodyAlertAggregations) GetAccountDisplayValues() []*string {
	return s.AccountDisplayValues
}

func (s *ListAlertAggregationsResponseBodyAlertAggregations) GetAccountIds() []*string {
	return s.AccountIds
}

func (s *ListAlertAggregationsResponseBodyAlertAggregations) GetAggregationKey() *string {
	return s.AggregationKey
}

func (s *ListAlertAggregationsResponseBodyAlertAggregations) GetAggregationType() *string {
	return s.AggregationType
}

func (s *ListAlertAggregationsResponseBodyAlertAggregations) GetAlertCount() *int64 {
	return s.AlertCount
}

func (s *ListAlertAggregationsResponseBodyAlertAggregations) GetAttackStages() []*string {
	return s.AttackStages
}

func (s *ListAlertAggregationsResponseBodyAlertAggregations) GetDefenseActionCounts() []*ListAlertAggregationsResponseBodyAlertAggregationsDefenseActionCounts {
	return s.DefenseActionCounts
}

func (s *ListAlertAggregationsResponseBodyAlertAggregations) GetFirstTime() *int64 {
	return s.FirstTime
}

func (s *ListAlertAggregationsResponseBodyAlertAggregations) GetLatestTime() *int64 {
	return s.LatestTime
}

func (s *ListAlertAggregationsResponseBodyAlertAggregations) GetPrimaryDisplayValue() *string {
	return s.PrimaryDisplayValue
}

func (s *ListAlertAggregationsResponseBodyAlertAggregations) GetPrimaryValue() *string {
	return s.PrimaryValue
}

func (s *ListAlertAggregationsResponseBodyAlertAggregations) GetSecondaryValue() *string {
	return s.SecondaryValue
}

func (s *ListAlertAggregationsResponseBodyAlertAggregations) GetSourceCodes() []*string {
	return s.SourceCodes
}

func (s *ListAlertAggregationsResponseBodyAlertAggregations) GetStatusCounts() []*ListAlertAggregationsResponseBodyAlertAggregationsStatusCounts {
	return s.StatusCounts
}

func (s *ListAlertAggregationsResponseBodyAlertAggregations) SetAccountDisplayValues(v []*string) *ListAlertAggregationsResponseBodyAlertAggregations {
	s.AccountDisplayValues = v
	return s
}

func (s *ListAlertAggregationsResponseBodyAlertAggregations) SetAccountIds(v []*string) *ListAlertAggregationsResponseBodyAlertAggregations {
	s.AccountIds = v
	return s
}

func (s *ListAlertAggregationsResponseBodyAlertAggregations) SetAggregationKey(v string) *ListAlertAggregationsResponseBodyAlertAggregations {
	s.AggregationKey = &v
	return s
}

func (s *ListAlertAggregationsResponseBodyAlertAggregations) SetAggregationType(v string) *ListAlertAggregationsResponseBodyAlertAggregations {
	s.AggregationType = &v
	return s
}

func (s *ListAlertAggregationsResponseBodyAlertAggregations) SetAlertCount(v int64) *ListAlertAggregationsResponseBodyAlertAggregations {
	s.AlertCount = &v
	return s
}

func (s *ListAlertAggregationsResponseBodyAlertAggregations) SetAttackStages(v []*string) *ListAlertAggregationsResponseBodyAlertAggregations {
	s.AttackStages = v
	return s
}

func (s *ListAlertAggregationsResponseBodyAlertAggregations) SetDefenseActionCounts(v []*ListAlertAggregationsResponseBodyAlertAggregationsDefenseActionCounts) *ListAlertAggregationsResponseBodyAlertAggregations {
	s.DefenseActionCounts = v
	return s
}

func (s *ListAlertAggregationsResponseBodyAlertAggregations) SetFirstTime(v int64) *ListAlertAggregationsResponseBodyAlertAggregations {
	s.FirstTime = &v
	return s
}

func (s *ListAlertAggregationsResponseBodyAlertAggregations) SetLatestTime(v int64) *ListAlertAggregationsResponseBodyAlertAggregations {
	s.LatestTime = &v
	return s
}

func (s *ListAlertAggregationsResponseBodyAlertAggregations) SetPrimaryDisplayValue(v string) *ListAlertAggregationsResponseBodyAlertAggregations {
	s.PrimaryDisplayValue = &v
	return s
}

func (s *ListAlertAggregationsResponseBodyAlertAggregations) SetPrimaryValue(v string) *ListAlertAggregationsResponseBodyAlertAggregations {
	s.PrimaryValue = &v
	return s
}

func (s *ListAlertAggregationsResponseBodyAlertAggregations) SetSecondaryValue(v string) *ListAlertAggregationsResponseBodyAlertAggregations {
	s.SecondaryValue = &v
	return s
}

func (s *ListAlertAggregationsResponseBodyAlertAggregations) SetSourceCodes(v []*string) *ListAlertAggregationsResponseBodyAlertAggregations {
	s.SourceCodes = v
	return s
}

func (s *ListAlertAggregationsResponseBodyAlertAggregations) SetStatusCounts(v []*ListAlertAggregationsResponseBodyAlertAggregationsStatusCounts) *ListAlertAggregationsResponseBodyAlertAggregations {
	s.StatusCounts = v
	return s
}

func (s *ListAlertAggregationsResponseBodyAlertAggregations) Validate() error {
	if s.DefenseActionCounts != nil {
		for _, item := range s.DefenseActionCounts {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.StatusCounts != nil {
		for _, item := range s.StatusCounts {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAlertAggregationsResponseBodyAlertAggregationsDefenseActionCounts struct {
	// example:
	//
	// accesstoken
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 测试-前途系统AI
	Value *int64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListAlertAggregationsResponseBodyAlertAggregationsDefenseActionCounts) String() string {
	return dara.Prettify(s)
}

func (s ListAlertAggregationsResponseBodyAlertAggregationsDefenseActionCounts) GoString() string {
	return s.String()
}

func (s *ListAlertAggregationsResponseBodyAlertAggregationsDefenseActionCounts) GetName() *string {
	return s.Name
}

func (s *ListAlertAggregationsResponseBodyAlertAggregationsDefenseActionCounts) GetValue() *int64 {
	return s.Value
}

func (s *ListAlertAggregationsResponseBodyAlertAggregationsDefenseActionCounts) SetName(v string) *ListAlertAggregationsResponseBodyAlertAggregationsDefenseActionCounts {
	s.Name = &v
	return s
}

func (s *ListAlertAggregationsResponseBodyAlertAggregationsDefenseActionCounts) SetValue(v int64) *ListAlertAggregationsResponseBodyAlertAggregationsDefenseActionCounts {
	s.Value = &v
	return s
}

func (s *ListAlertAggregationsResponseBodyAlertAggregationsDefenseActionCounts) Validate() error {
	return dara.Validate(s)
}

type ListAlertAggregationsResponseBodyAlertAggregationsStatusCounts struct {
	// example:
	//
	// http://domain.com/
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 123.123.123.123
	Value *int64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListAlertAggregationsResponseBodyAlertAggregationsStatusCounts) String() string {
	return dara.Prettify(s)
}

func (s ListAlertAggregationsResponseBodyAlertAggregationsStatusCounts) GoString() string {
	return s.String()
}

func (s *ListAlertAggregationsResponseBodyAlertAggregationsStatusCounts) GetName() *string {
	return s.Name
}

func (s *ListAlertAggregationsResponseBodyAlertAggregationsStatusCounts) GetValue() *int64 {
	return s.Value
}

func (s *ListAlertAggregationsResponseBodyAlertAggregationsStatusCounts) SetName(v string) *ListAlertAggregationsResponseBodyAlertAggregationsStatusCounts {
	s.Name = &v
	return s
}

func (s *ListAlertAggregationsResponseBodyAlertAggregationsStatusCounts) SetValue(v int64) *ListAlertAggregationsResponseBodyAlertAggregationsStatusCounts {
	s.Value = &v
	return s
}

func (s *ListAlertAggregationsResponseBodyAlertAggregationsStatusCounts) Validate() error {
	return dara.Validate(s)
}
