// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAiEngineEventNodesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEventId(v string) *DescribeAiEngineEventNodesRequest
	GetEventId() *string
	SetInstanceId(v string) *DescribeAiEngineEventNodesRequest
	GetInstanceId() *string
	SetLanguage(v string) *DescribeAiEngineEventNodesRequest
	GetLanguage() *string
	SetMaxResults(v int32) *DescribeAiEngineEventNodesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeAiEngineEventNodesRequest
	GetNextToken() *string
	SetRegionId(v string) *DescribeAiEngineEventNodesRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *DescribeAiEngineEventNodesRequest
	GetResourceManagerResourceGroupId() *string
	SetSortBy(v string) *DescribeAiEngineEventNodesRequest
	GetSortBy() *string
	SetSortOrder(v string) *DescribeAiEngineEventNodesRequest
	GetSortOrder() *string
}

type DescribeAiEngineEventNodesRequest struct {
	// The event ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// EVT-20241025109
	EventId *string `json:"EventId,omitempty" xml:"EventId,omitempty"`
	// The ID of the WAF instance.
	//
	// > You can call the [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) operation to query the ID of the current WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf-dem****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The language of the response.
	//
	// example:
	//
	// en
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// The number of entries to return on each page in a paging query. Valid values: 1 to 20. Default value: 20.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used to query the next page. Set the value to the **NextToken*	- value returned by the previous API call. You do not need to specify this parameter for the first query.
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
	// rg-deccc****
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
	// The sort field.
	//
	// example:
	//
	// occurred_at
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// The sort order.
	//
	// example:
	//
	// asc
	SortOrder *string `json:"SortOrder,omitempty" xml:"SortOrder,omitempty"`
}

func (s DescribeAiEngineEventNodesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAiEngineEventNodesRequest) GoString() string {
	return s.String()
}

func (s *DescribeAiEngineEventNodesRequest) GetEventId() *string {
	return s.EventId
}

func (s *DescribeAiEngineEventNodesRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeAiEngineEventNodesRequest) GetLanguage() *string {
	return s.Language
}

func (s *DescribeAiEngineEventNodesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeAiEngineEventNodesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeAiEngineEventNodesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeAiEngineEventNodesRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *DescribeAiEngineEventNodesRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *DescribeAiEngineEventNodesRequest) GetSortOrder() *string {
	return s.SortOrder
}

func (s *DescribeAiEngineEventNodesRequest) SetEventId(v string) *DescribeAiEngineEventNodesRequest {
	s.EventId = &v
	return s
}

func (s *DescribeAiEngineEventNodesRequest) SetInstanceId(v string) *DescribeAiEngineEventNodesRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeAiEngineEventNodesRequest) SetLanguage(v string) *DescribeAiEngineEventNodesRequest {
	s.Language = &v
	return s
}

func (s *DescribeAiEngineEventNodesRequest) SetMaxResults(v int32) *DescribeAiEngineEventNodesRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeAiEngineEventNodesRequest) SetNextToken(v string) *DescribeAiEngineEventNodesRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeAiEngineEventNodesRequest) SetRegionId(v string) *DescribeAiEngineEventNodesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAiEngineEventNodesRequest) SetResourceManagerResourceGroupId(v string) *DescribeAiEngineEventNodesRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeAiEngineEventNodesRequest) SetSortBy(v string) *DescribeAiEngineEventNodesRequest {
	s.SortBy = &v
	return s
}

func (s *DescribeAiEngineEventNodesRequest) SetSortOrder(v string) *DescribeAiEngineEventNodesRequest {
	s.SortOrder = &v
	return s
}

func (s *DescribeAiEngineEventNodesRequest) Validate() error {
	return dara.Validate(s)
}
