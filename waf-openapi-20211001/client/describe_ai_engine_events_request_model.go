// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAiEngineEventsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v int64) *DescribeAiEngineEventsRequest
	GetEndTime() *int64
	SetFilter(v *DescribeAiEngineEventsRequestFilter) *DescribeAiEngineEventsRequest
	GetFilter() *DescribeAiEngineEventsRequestFilter
	SetInstanceId(v string) *DescribeAiEngineEventsRequest
	GetInstanceId() *string
	SetLanguage(v string) *DescribeAiEngineEventsRequest
	GetLanguage() *string
	SetMaxResults(v int32) *DescribeAiEngineEventsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeAiEngineEventsRequest
	GetNextToken() *string
	SetRegionId(v string) *DescribeAiEngineEventsRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *DescribeAiEngineEventsRequest
	GetResourceManagerResourceGroupId() *string
	SetSeachValue(v string) *DescribeAiEngineEventsRequest
	GetSeachValue() *string
	SetSearchKey(v string) *DescribeAiEngineEventsRequest
	GetSearchKey() *string
	SetSortBy(v string) *DescribeAiEngineEventsRequest
	GetSortBy() *string
	SetSortOrder(v string) *DescribeAiEngineEventsRequest
	GetSortOrder() *string
	SetStartTime(v int64) *DescribeAiEngineEventsRequest
	GetStartTime() *int64
}

type DescribeAiEngineEventsRequest struct {
	// The end time. The time is a UNIX timestamp in UTC. Unit: seconds. The data is filtered based on the time when the event was last detected. The end time must be later than the start time.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1793059200
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The filter conditions. Exact match is used, and multiple conditions are connected by the AND logical operator.
	//
	// example:
	//
	// {"riskLevel":"high","resource":"resource-api-example"}
	Filter *DescribeAiEngineEventsRequestFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Struct"`
	// The ID of the WAF instance.
	//
	// > You can call the [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) operation to query the ID of the current WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_dddsdf3****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The language of the response.
	//
	// example:
	//
	// en
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// The number of entries to return on each page in a paging query. Valid values: 1 to 100. Default value: 20.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token. The value is the **NextToken*	- parameter value returned by the previous API call. You do not need to specify this parameter for the first query.
	//
	// example:
	//
	// AAAAAGBgV9tolsLfijC4wam2htS*****D/46H3X2wIS
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The region ID of the WAF instance.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the Alibaba Cloud resource group.
	//
	// example:
	//
	// rg-****mo
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
	// The search value. This parameter is used together with SearchKey. Maximum length: 256 characters.
	//
	// example:
	//
	// testKey
	SeachValue *string `json:"SeachValue,omitempty" xml:"SeachValue,omitempty"`
	// The search field. This parameter is used together with SeachValue.
	//
	// example:
	//
	// eventName
	SearchKey *string `json:"SearchKey,omitempty" xml:"SearchKey,omitempty"`
	// The field by which to sort the results.
	//
	// example:
	//
	// last_seen
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// The sort order.
	//
	// example:
	//
	// desc
	SortOrder *string `json:"SortOrder,omitempty" xml:"SortOrder,omitempty"`
	// The start time. The time is a UNIX timestamp in UTC. Unit: seconds. The data is filtered based on the time when the event was last detected. You can query data within the last 30 days.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1792972800
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeAiEngineEventsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAiEngineEventsRequest) GoString() string {
	return s.String()
}

func (s *DescribeAiEngineEventsRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeAiEngineEventsRequest) GetFilter() *DescribeAiEngineEventsRequestFilter {
	return s.Filter
}

func (s *DescribeAiEngineEventsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeAiEngineEventsRequest) GetLanguage() *string {
	return s.Language
}

func (s *DescribeAiEngineEventsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeAiEngineEventsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeAiEngineEventsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeAiEngineEventsRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *DescribeAiEngineEventsRequest) GetSeachValue() *string {
	return s.SeachValue
}

func (s *DescribeAiEngineEventsRequest) GetSearchKey() *string {
	return s.SearchKey
}

func (s *DescribeAiEngineEventsRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *DescribeAiEngineEventsRequest) GetSortOrder() *string {
	return s.SortOrder
}

func (s *DescribeAiEngineEventsRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeAiEngineEventsRequest) SetEndTime(v int64) *DescribeAiEngineEventsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeAiEngineEventsRequest) SetFilter(v *DescribeAiEngineEventsRequestFilter) *DescribeAiEngineEventsRequest {
	s.Filter = v
	return s
}

func (s *DescribeAiEngineEventsRequest) SetInstanceId(v string) *DescribeAiEngineEventsRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeAiEngineEventsRequest) SetLanguage(v string) *DescribeAiEngineEventsRequest {
	s.Language = &v
	return s
}

func (s *DescribeAiEngineEventsRequest) SetMaxResults(v int32) *DescribeAiEngineEventsRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeAiEngineEventsRequest) SetNextToken(v string) *DescribeAiEngineEventsRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeAiEngineEventsRequest) SetRegionId(v string) *DescribeAiEngineEventsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAiEngineEventsRequest) SetResourceManagerResourceGroupId(v string) *DescribeAiEngineEventsRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeAiEngineEventsRequest) SetSeachValue(v string) *DescribeAiEngineEventsRequest {
	s.SeachValue = &v
	return s
}

func (s *DescribeAiEngineEventsRequest) SetSearchKey(v string) *DescribeAiEngineEventsRequest {
	s.SearchKey = &v
	return s
}

func (s *DescribeAiEngineEventsRequest) SetSortBy(v string) *DescribeAiEngineEventsRequest {
	s.SortBy = &v
	return s
}

func (s *DescribeAiEngineEventsRequest) SetSortOrder(v string) *DescribeAiEngineEventsRequest {
	s.SortOrder = &v
	return s
}

func (s *DescribeAiEngineEventsRequest) SetStartTime(v int64) *DescribeAiEngineEventsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeAiEngineEventsRequest) Validate() error {
	if s.Filter != nil {
		if err := s.Filter.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeAiEngineEventsRequestFilter struct {
	// The domain name.
	//
	// example:
	//
	// api.example.com
	Domain *string `json:"domain,omitempty" xml:"domain,omitempty"`
	// The event ID.
	//
	// example:
	//
	// EVT-20261025109
	EventId *string `json:"eventId,omitempty" xml:"eventId,omitempty"`
	// The number of nodes included in the event.
	//
	// example:
	//
	// 10
	EventNodeCount *int32 `json:"eventNodeCount,omitempty" xml:"eventNodeCount,omitempty"`
	// The event handling status.
	//
	// example:
	//
	// unhandled
	HandleStatus *string `json:"handleStatus,omitempty" xml:"handleStatus,omitempty"`
	// The protected object.
	//
	// example:
	//
	// resource-api-example
	Resource *string `json:"resource,omitempty" xml:"resource,omitempty"`
	// The risk level.
	//
	// example:
	//
	// high
	RiskLevel *string `json:"riskLevel,omitempty" xml:"riskLevel,omitempty"`
	// The source account of the attack.
	//
	// example:
	//
	// account-demo-***
	SourceAccount *string `json:"sourceAccount,omitempty" xml:"sourceAccount,omitempty"`
	// The source fingerprint of the attack.
	//
	// example:
	//
	// fingerprint-demo-***
	SourceFingerprint *string `json:"sourceFingerprint,omitempty" xml:"sourceFingerprint,omitempty"`
	// The source IP address of the attack.
	//
	// example:
	//
	// 203.0.113.109
	SourceIp *string `json:"sourceIp,omitempty" xml:"sourceIp,omitempty"`
}

func (s DescribeAiEngineEventsRequestFilter) String() string {
	return dara.Prettify(s)
}

func (s DescribeAiEngineEventsRequestFilter) GoString() string {
	return s.String()
}

func (s *DescribeAiEngineEventsRequestFilter) GetDomain() *string {
	return s.Domain
}

func (s *DescribeAiEngineEventsRequestFilter) GetEventId() *string {
	return s.EventId
}

func (s *DescribeAiEngineEventsRequestFilter) GetEventNodeCount() *int32 {
	return s.EventNodeCount
}

func (s *DescribeAiEngineEventsRequestFilter) GetHandleStatus() *string {
	return s.HandleStatus
}

func (s *DescribeAiEngineEventsRequestFilter) GetResource() *string {
	return s.Resource
}

func (s *DescribeAiEngineEventsRequestFilter) GetRiskLevel() *string {
	return s.RiskLevel
}

func (s *DescribeAiEngineEventsRequestFilter) GetSourceAccount() *string {
	return s.SourceAccount
}

func (s *DescribeAiEngineEventsRequestFilter) GetSourceFingerprint() *string {
	return s.SourceFingerprint
}

func (s *DescribeAiEngineEventsRequestFilter) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeAiEngineEventsRequestFilter) SetDomain(v string) *DescribeAiEngineEventsRequestFilter {
	s.Domain = &v
	return s
}

func (s *DescribeAiEngineEventsRequestFilter) SetEventId(v string) *DescribeAiEngineEventsRequestFilter {
	s.EventId = &v
	return s
}

func (s *DescribeAiEngineEventsRequestFilter) SetEventNodeCount(v int32) *DescribeAiEngineEventsRequestFilter {
	s.EventNodeCount = &v
	return s
}

func (s *DescribeAiEngineEventsRequestFilter) SetHandleStatus(v string) *DescribeAiEngineEventsRequestFilter {
	s.HandleStatus = &v
	return s
}

func (s *DescribeAiEngineEventsRequestFilter) SetResource(v string) *DescribeAiEngineEventsRequestFilter {
	s.Resource = &v
	return s
}

func (s *DescribeAiEngineEventsRequestFilter) SetRiskLevel(v string) *DescribeAiEngineEventsRequestFilter {
	s.RiskLevel = &v
	return s
}

func (s *DescribeAiEngineEventsRequestFilter) SetSourceAccount(v string) *DescribeAiEngineEventsRequestFilter {
	s.SourceAccount = &v
	return s
}

func (s *DescribeAiEngineEventsRequestFilter) SetSourceFingerprint(v string) *DescribeAiEngineEventsRequestFilter {
	s.SourceFingerprint = &v
	return s
}

func (s *DescribeAiEngineEventsRequestFilter) SetSourceIp(v string) *DescribeAiEngineEventsRequestFilter {
	s.SourceIp = &v
	return s
}

func (s *DescribeAiEngineEventsRequestFilter) Validate() error {
	return dara.Validate(s)
}
