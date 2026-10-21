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
	// The list of alert groups returned based on the specified aggregation method, sorted by alert count in descending order.
	AlertAggregations []*ListAlertAggregationsResponseBodyAlertAggregations `json:"AlertAggregations,omitempty" xml:"AlertAggregations,omitempty" type:"Repeated"`
	// The maximum number of entries returned in this request.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token for the next query.
	//
	// example:
	//
	// AAAAAUqcj6VO4E3ECWIrFczs****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9AAA9ED9-78F4-5021-86DC-D51C7511****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of records.
	//
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
	// The list of account display values associated with the aggregation group, in one-to-one correspondence with AccountIds. When the username is valid and differs from the account ID, the format is "username/account ID". Otherwise, only the account ID is returned.
	AccountDisplayValues []*string `json:"AccountDisplayValues,omitempty" xml:"AccountDisplayValues,omitempty" type:"Repeated"`
	// The list of member accounts IDs for which the report is generated. You need to specify this parameter only when you generate a multi-account report and want to limit the report to specific member accounts.
	AccountIds []*string `json:"AccountIds,omitempty" xml:"AccountIds,omitempty" type:"Repeated"`
	// The unique identifier of the aggregation group, which is a concatenation of AggregationType, PrimaryValue, and SecondaryValue joined by colons (:). The last segment is empty if no secondary aggregation value exists.
	//
	// example:
	//
	// vendor_product:alibaba_cloud:sas
	AggregationKey *string `json:"AggregationKey,omitempty" xml:"AggregationKey,omitempty"`
	// The aggregation method.
	//
	// example:
	//
	// 5m-avg
	AggregationType *string `json:"AggregationType,omitempty" xml:"AggregationType,omitempty"`
	// The number of consecutive failures. This parameter is optional.
	//
	// example:
	//
	// 25
	AlertCount *int64 `json:"AlertCount,omitempty" xml:"AlertCount,omitempty"`
	// The collection of MITRE ATT&CK attack techniques associated with alerts in the aggregation group. Each element is a JSON array string.
	AttackStages []*string `json:"AttackStages,omitempty" xml:"AttackStages,omitempty" type:"Repeated"`
	// The quantity distribution by defense action within the aggregation group. Name is the action field value, and Value is the corresponding alert count.
	DefenseActionCounts []*ListAlertAggregationsResponseBodyAlertAggregationsDefenseActionCounts `json:"DefenseActionCounts,omitempty" xml:"DefenseActionCounts,omitempty" type:"Repeated"`
	// The time of the first access, represented as a UNIX timestamp in seconds.
	//
	// example:
	//
	// 1785290308
	FirstTime *int64 `json:"FirstTime,omitempty" xml:"FirstTime,omitempty"`
	// The most recent time when the attack event occurred.
	//
	// example:
	//
	// 1785293908000
	LatestTime *int64 `json:"LatestTime,omitempty" xml:"LatestTime,omitempty"`
	// The display text of the primary aggregation value. For account aggregation, the format is "username/account ID". If no valid username exists, only the account ID is returned. For event summaries, the event name is returned based on the Lang parameter. If the name is missing, the event UUID is returned.
	//
	// example:
	//
	// Test incident
	PrimaryDisplayValue *string `json:"PrimaryDisplayValue,omitempty" xml:"PrimaryDisplayValue,omitempty"`
	// The value corresponding to the primary query condition.
	//
	// example:
	//
	// alibaba_cloud
	PrimaryValue *string `json:"PrimaryValue,omitempty" xml:"PrimaryValue,omitempty"`
	// The value corresponding to the secondary query condition.
	//
	// example:
	//
	// sas
	SecondaryValue *string `json:"SecondaryValue,omitempty" xml:"SecondaryValue,omitempty"`
	// The collection of event source codes.
	//
	// > Example: [\\\\\\"MRM\\\\\\"]
	SourceCodes []*string `json:"SourceCodes,omitempty" xml:"SourceCodes,omitempty" type:"Repeated"`
	// The quantity distribution by alert status within the aggregation group. Name is the alert status value, and Value is the corresponding alert count.
	StatusCounts []*ListAlertAggregationsResponseBodyAlertAggregationsStatusCounts `json:"StatusCounts,omitempty" xml:"StatusCounts,omitempty" type:"Repeated"`
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
	// The display name of the entity attribute field that can be added to the whitelist.
	//
	// example:
	//
	// accesstoken
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The default right-hand value displayed for the whitelist rule.
	//
	// example:
	//
	// Test-FutureSystemAI
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
	// The display name of the entity attribute field that can be added to the whitelist.
	//
	// example:
	//
	// http://domain.com/
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The right-hand value.
	//
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
