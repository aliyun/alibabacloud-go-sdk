// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAiEngineEvidenceRequestsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEventId(v string) *DescribeAiEngineEvidenceRequestsRequest
	GetEventId() *string
	SetFilter(v *DescribeAiEngineEvidenceRequestsRequestFilter) *DescribeAiEngineEvidenceRequestsRequest
	GetFilter() *DescribeAiEngineEvidenceRequestsRequestFilter
	SetInstanceId(v string) *DescribeAiEngineEvidenceRequestsRequest
	GetInstanceId() *string
	SetLanguage(v string) *DescribeAiEngineEvidenceRequestsRequest
	GetLanguage() *string
	SetMaxResults(v int32) *DescribeAiEngineEvidenceRequestsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeAiEngineEvidenceRequestsRequest
	GetNextToken() *string
	SetRegionId(v string) *DescribeAiEngineEvidenceRequestsRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *DescribeAiEngineEvidenceRequestsRequest
	GetResourceManagerResourceGroupId() *string
	SetSeachValue(v string) *DescribeAiEngineEvidenceRequestsRequest
	GetSeachValue() *string
	SetSearchKey(v string) *DescribeAiEngineEvidenceRequestsRequest
	GetSearchKey() *string
	SetSortBy(v string) *DescribeAiEngineEvidenceRequestsRequest
	GetSortBy() *string
	SetSortOrder(v string) *DescribeAiEngineEvidenceRequestsRequest
	GetSortOrder() *string
}

type DescribeAiEngineEvidenceRequestsRequest struct {
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
	Filter *DescribeAiEngineEvidenceRequestsRequestFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Struct"`
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

func (s DescribeAiEngineEvidenceRequestsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAiEngineEvidenceRequestsRequest) GoString() string {
	return s.String()
}

func (s *DescribeAiEngineEvidenceRequestsRequest) GetEventId() *string {
	return s.EventId
}

func (s *DescribeAiEngineEvidenceRequestsRequest) GetFilter() *DescribeAiEngineEvidenceRequestsRequestFilter {
	return s.Filter
}

func (s *DescribeAiEngineEvidenceRequestsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeAiEngineEvidenceRequestsRequest) GetLanguage() *string {
	return s.Language
}

func (s *DescribeAiEngineEvidenceRequestsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeAiEngineEvidenceRequestsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeAiEngineEvidenceRequestsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeAiEngineEvidenceRequestsRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *DescribeAiEngineEvidenceRequestsRequest) GetSeachValue() *string {
	return s.SeachValue
}

func (s *DescribeAiEngineEvidenceRequestsRequest) GetSearchKey() *string {
	return s.SearchKey
}

func (s *DescribeAiEngineEvidenceRequestsRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *DescribeAiEngineEvidenceRequestsRequest) GetSortOrder() *string {
	return s.SortOrder
}

func (s *DescribeAiEngineEvidenceRequestsRequest) SetEventId(v string) *DescribeAiEngineEvidenceRequestsRequest {
	s.EventId = &v
	return s
}

func (s *DescribeAiEngineEvidenceRequestsRequest) SetFilter(v *DescribeAiEngineEvidenceRequestsRequestFilter) *DescribeAiEngineEvidenceRequestsRequest {
	s.Filter = v
	return s
}

func (s *DescribeAiEngineEvidenceRequestsRequest) SetInstanceId(v string) *DescribeAiEngineEvidenceRequestsRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeAiEngineEvidenceRequestsRequest) SetLanguage(v string) *DescribeAiEngineEvidenceRequestsRequest {
	s.Language = &v
	return s
}

func (s *DescribeAiEngineEvidenceRequestsRequest) SetMaxResults(v int32) *DescribeAiEngineEvidenceRequestsRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeAiEngineEvidenceRequestsRequest) SetNextToken(v string) *DescribeAiEngineEvidenceRequestsRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeAiEngineEvidenceRequestsRequest) SetRegionId(v string) *DescribeAiEngineEvidenceRequestsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAiEngineEvidenceRequestsRequest) SetResourceManagerResourceGroupId(v string) *DescribeAiEngineEvidenceRequestsRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeAiEngineEvidenceRequestsRequest) SetSeachValue(v string) *DescribeAiEngineEvidenceRequestsRequest {
	s.SeachValue = &v
	return s
}

func (s *DescribeAiEngineEvidenceRequestsRequest) SetSearchKey(v string) *DescribeAiEngineEvidenceRequestsRequest {
	s.SearchKey = &v
	return s
}

func (s *DescribeAiEngineEvidenceRequestsRequest) SetSortBy(v string) *DescribeAiEngineEvidenceRequestsRequest {
	s.SortBy = &v
	return s
}

func (s *DescribeAiEngineEvidenceRequestsRequest) SetSortOrder(v string) *DescribeAiEngineEvidenceRequestsRequest {
	s.SortOrder = &v
	return s
}

func (s *DescribeAiEngineEvidenceRequestsRequest) Validate() error {
	if s.Filter != nil {
		if err := s.Filter.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeAiEngineEvidenceRequestsRequestFilter struct {
	// The HTTP status code.
	//
	// example:
	//
	// 500
	HttpStatus *string `json:"httpStatus,omitempty" xml:"httpStatus,omitempty"`
	// The evidence impact tag.
	//
	// example:
	//
	// bypass
	ImpactType *string `json:"impactType,omitempty" xml:"impactType,omitempty"`
	// The HTTP request method. Exact match is used.
	//
	// example:
	//
	// POST
	Method *string `json:"method,omitempty" xml:"method,omitempty"`
	// The node ID. Exact match is used. The value can be up to 64 ASCII characters in length.
	//
	// example:
	//
	// n109-04
	NodeId *string `json:"nodeId,omitempty" xml:"nodeId,omitempty"`
	// The attack source type.
	//
	// example:
	//
	// ip
	SourceType *string `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
	// Specifies whether the request is a suspected bypass.
	//
	// example:
	//
	// 1
	SuspectedBypass *string `json:"suspectedBypass,omitempty" xml:"suspectedBypass,omitempty"`
}

func (s DescribeAiEngineEvidenceRequestsRequestFilter) String() string {
	return dara.Prettify(s)
}

func (s DescribeAiEngineEvidenceRequestsRequestFilter) GoString() string {
	return s.String()
}

func (s *DescribeAiEngineEvidenceRequestsRequestFilter) GetHttpStatus() *string {
	return s.HttpStatus
}

func (s *DescribeAiEngineEvidenceRequestsRequestFilter) GetImpactType() *string {
	return s.ImpactType
}

func (s *DescribeAiEngineEvidenceRequestsRequestFilter) GetMethod() *string {
	return s.Method
}

func (s *DescribeAiEngineEvidenceRequestsRequestFilter) GetNodeId() *string {
	return s.NodeId
}

func (s *DescribeAiEngineEvidenceRequestsRequestFilter) GetSourceType() *string {
	return s.SourceType
}

func (s *DescribeAiEngineEvidenceRequestsRequestFilter) GetSuspectedBypass() *string {
	return s.SuspectedBypass
}

func (s *DescribeAiEngineEvidenceRequestsRequestFilter) SetHttpStatus(v string) *DescribeAiEngineEvidenceRequestsRequestFilter {
	s.HttpStatus = &v
	return s
}

func (s *DescribeAiEngineEvidenceRequestsRequestFilter) SetImpactType(v string) *DescribeAiEngineEvidenceRequestsRequestFilter {
	s.ImpactType = &v
	return s
}

func (s *DescribeAiEngineEvidenceRequestsRequestFilter) SetMethod(v string) *DescribeAiEngineEvidenceRequestsRequestFilter {
	s.Method = &v
	return s
}

func (s *DescribeAiEngineEvidenceRequestsRequestFilter) SetNodeId(v string) *DescribeAiEngineEvidenceRequestsRequestFilter {
	s.NodeId = &v
	return s
}

func (s *DescribeAiEngineEvidenceRequestsRequestFilter) SetSourceType(v string) *DescribeAiEngineEvidenceRequestsRequestFilter {
	s.SourceType = &v
	return s
}

func (s *DescribeAiEngineEvidenceRequestsRequestFilter) SetSuspectedBypass(v string) *DescribeAiEngineEvidenceRequestsRequestFilter {
	s.SuspectedBypass = &v
	return s
}

func (s *DescribeAiEngineEvidenceRequestsRequestFilter) Validate() error {
	return dara.Validate(s)
}
