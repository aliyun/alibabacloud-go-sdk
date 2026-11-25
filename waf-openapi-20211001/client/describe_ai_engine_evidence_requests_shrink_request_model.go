// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAiEngineEvidenceRequestsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEventId(v string) *DescribeAiEngineEvidenceRequestsShrinkRequest
	GetEventId() *string
	SetFilterShrink(v string) *DescribeAiEngineEvidenceRequestsShrinkRequest
	GetFilterShrink() *string
	SetInstanceId(v string) *DescribeAiEngineEvidenceRequestsShrinkRequest
	GetInstanceId() *string
	SetLanguage(v string) *DescribeAiEngineEvidenceRequestsShrinkRequest
	GetLanguage() *string
	SetMaxResults(v int32) *DescribeAiEngineEvidenceRequestsShrinkRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeAiEngineEvidenceRequestsShrinkRequest
	GetNextToken() *string
	SetRegionId(v string) *DescribeAiEngineEvidenceRequestsShrinkRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *DescribeAiEngineEvidenceRequestsShrinkRequest
	GetResourceManagerResourceGroupId() *string
	SetSeachValue(v string) *DescribeAiEngineEvidenceRequestsShrinkRequest
	GetSeachValue() *string
	SetSearchKey(v string) *DescribeAiEngineEvidenceRequestsShrinkRequest
	GetSearchKey() *string
	SetSortBy(v string) *DescribeAiEngineEvidenceRequestsShrinkRequest
	GetSortBy() *string
	SetSortOrder(v string) *DescribeAiEngineEvidenceRequestsShrinkRequest
	GetSortOrder() *string
}

type DescribeAiEngineEvidenceRequestsShrinkRequest struct {
	// The event ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// EVT-20241025-109
	EventId *string `json:"EventId,omitempty" xml:"EventId,omitempty"`
	// The filter conditions. Exact match is used, and multiple conditions are connected by the AND logical operator.
	//
	// example:
	//
	// {"suspectedBypass":"1","httpStatus":"500"}
	FilterShrink *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	// The ID of the WAF instance.
	//
	// > You can call the [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) operation to query the ID of the current WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf-de****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The response language.
	//
	// example:
	//
	// en
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// The number of entries to return per page. Valid values: 1 to 100. Default value: 20.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token for the next page. The value is the **NextToken*	- parameter value returned by the previous API call. You do not need to specify this parameter for the first query.
	//
	// example:
	//
	// AAAAAGBgV9tolsLfijC4wam2htS*****D/46H3X2wIS
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The region where the WAF instance resides.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the Alibaba Cloud resource group.
	//
	// example:
	//
	// rg-dsdo*****
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
	// The search value. This parameter is used together with SearchKey. The value can be up to 256 characters in length.
	//
	// example:
	//
	// trace-20241***ssdas
	SeachValue *string `json:"SeachValue,omitempty" xml:"SeachValue,omitempty"`
	// The search key. This parameter is used together with SeachValue.
	//
	// example:
	//
	// traceId
	SearchKey *string `json:"SearchKey,omitempty" xml:"SearchKey,omitempty"`
	// The sort field.
	//
	// example:
	//
	// request_time
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// The sort order.
	//
	// example:
	//
	// asc
	SortOrder *string `json:"SortOrder,omitempty" xml:"SortOrder,omitempty"`
}

func (s DescribeAiEngineEvidenceRequestsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAiEngineEvidenceRequestsShrinkRequest) GoString() string {
	return s.String()
}

func (s *DescribeAiEngineEvidenceRequestsShrinkRequest) GetEventId() *string {
	return s.EventId
}

func (s *DescribeAiEngineEvidenceRequestsShrinkRequest) GetFilterShrink() *string {
	return s.FilterShrink
}

func (s *DescribeAiEngineEvidenceRequestsShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeAiEngineEvidenceRequestsShrinkRequest) GetLanguage() *string {
	return s.Language
}

func (s *DescribeAiEngineEvidenceRequestsShrinkRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeAiEngineEvidenceRequestsShrinkRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeAiEngineEvidenceRequestsShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeAiEngineEvidenceRequestsShrinkRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *DescribeAiEngineEvidenceRequestsShrinkRequest) GetSeachValue() *string {
	return s.SeachValue
}

func (s *DescribeAiEngineEvidenceRequestsShrinkRequest) GetSearchKey() *string {
	return s.SearchKey
}

func (s *DescribeAiEngineEvidenceRequestsShrinkRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *DescribeAiEngineEvidenceRequestsShrinkRequest) GetSortOrder() *string {
	return s.SortOrder
}

func (s *DescribeAiEngineEvidenceRequestsShrinkRequest) SetEventId(v string) *DescribeAiEngineEvidenceRequestsShrinkRequest {
	s.EventId = &v
	return s
}

func (s *DescribeAiEngineEvidenceRequestsShrinkRequest) SetFilterShrink(v string) *DescribeAiEngineEvidenceRequestsShrinkRequest {
	s.FilterShrink = &v
	return s
}

func (s *DescribeAiEngineEvidenceRequestsShrinkRequest) SetInstanceId(v string) *DescribeAiEngineEvidenceRequestsShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeAiEngineEvidenceRequestsShrinkRequest) SetLanguage(v string) *DescribeAiEngineEvidenceRequestsShrinkRequest {
	s.Language = &v
	return s
}

func (s *DescribeAiEngineEvidenceRequestsShrinkRequest) SetMaxResults(v int32) *DescribeAiEngineEvidenceRequestsShrinkRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeAiEngineEvidenceRequestsShrinkRequest) SetNextToken(v string) *DescribeAiEngineEvidenceRequestsShrinkRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeAiEngineEvidenceRequestsShrinkRequest) SetRegionId(v string) *DescribeAiEngineEvidenceRequestsShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAiEngineEvidenceRequestsShrinkRequest) SetResourceManagerResourceGroupId(v string) *DescribeAiEngineEvidenceRequestsShrinkRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeAiEngineEvidenceRequestsShrinkRequest) SetSeachValue(v string) *DescribeAiEngineEvidenceRequestsShrinkRequest {
	s.SeachValue = &v
	return s
}

func (s *DescribeAiEngineEvidenceRequestsShrinkRequest) SetSearchKey(v string) *DescribeAiEngineEvidenceRequestsShrinkRequest {
	s.SearchKey = &v
	return s
}

func (s *DescribeAiEngineEvidenceRequestsShrinkRequest) SetSortBy(v string) *DescribeAiEngineEvidenceRequestsShrinkRequest {
	s.SortBy = &v
	return s
}

func (s *DescribeAiEngineEvidenceRequestsShrinkRequest) SetSortOrder(v string) *DescribeAiEngineEvidenceRequestsShrinkRequest {
	s.SortOrder = &v
	return s
}

func (s *DescribeAiEngineEvidenceRequestsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
