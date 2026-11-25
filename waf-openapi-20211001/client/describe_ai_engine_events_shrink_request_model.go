// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAiEngineEventsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v int64) *DescribeAiEngineEventsShrinkRequest
	GetEndTime() *int64
	SetFilterShrink(v string) *DescribeAiEngineEventsShrinkRequest
	GetFilterShrink() *string
	SetInstanceId(v string) *DescribeAiEngineEventsShrinkRequest
	GetInstanceId() *string
	SetLanguage(v string) *DescribeAiEngineEventsShrinkRequest
	GetLanguage() *string
	SetMaxResults(v int32) *DescribeAiEngineEventsShrinkRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeAiEngineEventsShrinkRequest
	GetNextToken() *string
	SetRegionId(v string) *DescribeAiEngineEventsShrinkRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *DescribeAiEngineEventsShrinkRequest
	GetResourceManagerResourceGroupId() *string
	SetSeachValue(v string) *DescribeAiEngineEventsShrinkRequest
	GetSeachValue() *string
	SetSearchKey(v string) *DescribeAiEngineEventsShrinkRequest
	GetSearchKey() *string
	SetSortBy(v string) *DescribeAiEngineEventsShrinkRequest
	GetSortBy() *string
	SetSortOrder(v string) *DescribeAiEngineEventsShrinkRequest
	GetSortOrder() *string
	SetStartTime(v int64) *DescribeAiEngineEventsShrinkRequest
	GetStartTime() *int64
}

type DescribeAiEngineEventsShrinkRequest struct {
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
	FilterShrink *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
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

func (s DescribeAiEngineEventsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAiEngineEventsShrinkRequest) GoString() string {
	return s.String()
}

func (s *DescribeAiEngineEventsShrinkRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeAiEngineEventsShrinkRequest) GetFilterShrink() *string {
	return s.FilterShrink
}

func (s *DescribeAiEngineEventsShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeAiEngineEventsShrinkRequest) GetLanguage() *string {
	return s.Language
}

func (s *DescribeAiEngineEventsShrinkRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeAiEngineEventsShrinkRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeAiEngineEventsShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeAiEngineEventsShrinkRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *DescribeAiEngineEventsShrinkRequest) GetSeachValue() *string {
	return s.SeachValue
}

func (s *DescribeAiEngineEventsShrinkRequest) GetSearchKey() *string {
	return s.SearchKey
}

func (s *DescribeAiEngineEventsShrinkRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *DescribeAiEngineEventsShrinkRequest) GetSortOrder() *string {
	return s.SortOrder
}

func (s *DescribeAiEngineEventsShrinkRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeAiEngineEventsShrinkRequest) SetEndTime(v int64) *DescribeAiEngineEventsShrinkRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeAiEngineEventsShrinkRequest) SetFilterShrink(v string) *DescribeAiEngineEventsShrinkRequest {
	s.FilterShrink = &v
	return s
}

func (s *DescribeAiEngineEventsShrinkRequest) SetInstanceId(v string) *DescribeAiEngineEventsShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeAiEngineEventsShrinkRequest) SetLanguage(v string) *DescribeAiEngineEventsShrinkRequest {
	s.Language = &v
	return s
}

func (s *DescribeAiEngineEventsShrinkRequest) SetMaxResults(v int32) *DescribeAiEngineEventsShrinkRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeAiEngineEventsShrinkRequest) SetNextToken(v string) *DescribeAiEngineEventsShrinkRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeAiEngineEventsShrinkRequest) SetRegionId(v string) *DescribeAiEngineEventsShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAiEngineEventsShrinkRequest) SetResourceManagerResourceGroupId(v string) *DescribeAiEngineEventsShrinkRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeAiEngineEventsShrinkRequest) SetSeachValue(v string) *DescribeAiEngineEventsShrinkRequest {
	s.SeachValue = &v
	return s
}

func (s *DescribeAiEngineEventsShrinkRequest) SetSearchKey(v string) *DescribeAiEngineEventsShrinkRequest {
	s.SearchKey = &v
	return s
}

func (s *DescribeAiEngineEventsShrinkRequest) SetSortBy(v string) *DescribeAiEngineEventsShrinkRequest {
	s.SortBy = &v
	return s
}

func (s *DescribeAiEngineEventsShrinkRequest) SetSortOrder(v string) *DescribeAiEngineEventsShrinkRequest {
	s.SortOrder = &v
	return s
}

func (s *DescribeAiEngineEventsShrinkRequest) SetStartTime(v int64) *DescribeAiEngineEventsShrinkRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeAiEngineEventsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
