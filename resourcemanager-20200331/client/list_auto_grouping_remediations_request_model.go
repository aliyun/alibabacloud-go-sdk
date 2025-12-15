// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAutoGroupingRemediationsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEarliestRemediationTime(v string) *ListAutoGroupingRemediationsRequest
	GetEarliestRemediationTime() *string
	SetLatestRemediationTime(v string) *ListAutoGroupingRemediationsRequest
	GetLatestRemediationTime() *string
	SetMaxResults(v int32) *ListAutoGroupingRemediationsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListAutoGroupingRemediationsRequest
	GetNextToken() *string
	SetResourceId(v string) *ListAutoGroupingRemediationsRequest
	GetResourceId() *string
	SetResourceType(v string) *ListAutoGroupingRemediationsRequest
	GetResourceType() *string
	SetRuleId(v string) *ListAutoGroupingRemediationsRequest
	GetRuleId() *string
	SetService(v string) *ListAutoGroupingRemediationsRequest
	GetService() *string
	SetTargetResourceGroupId(v string) *ListAutoGroupingRemediationsRequest
	GetTargetResourceGroupId() *string
}

type ListAutoGroupingRemediationsRequest struct {
	// The earliest remediation time. This parameter is empty by default.
	//
	// example:
	//
	// 2022-01-01 00:00:00
	EarliestRemediationTime *string `json:"EarliestRemediationTime,omitempty" xml:"EarliestRemediationTime,omitempty"`
	// The latest remediation time. This parameter is empty by default.
	//
	// example:
	//
	// 2022-02-01 00:00:00
	LatestRemediationTime *string `json:"LatestRemediationTime,omitempty" xml:"LatestRemediationTime,omitempty"`
	// The maximum number of data entries to return.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. If you leave this parameter empty, the query starts from the beginning.
	//
	// example:
	//
	// TGlzdFJlc291cm****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The resource ID,
	//
	// example:
	//
	// i-23v38****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The resource type,
	//
	// You can obtain the resource type from the **Resource type*	- column in [Services that work with Resource Group](https://help.aliyun.com/document_detail/94479.html).
	//
	// example:
	//
	// instance
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The rule ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// gr-acfnugygwms32yy
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The ID of the Alibaba Cloud service.
	//
	// You can obtain the ID from the **Service code*	- column in [Services that work with Resource Group](https://help.aliyun.com/document_detail/94479.html).
	//
	// example:
	//
	// ecs
	Service *string `json:"Service,omitempty" xml:"Service,omitempty"`
	// The ID of the new resource group.
	//
	// example:
	//
	// rg-aekz26emqhc****
	TargetResourceGroupId *string `json:"TargetResourceGroupId,omitempty" xml:"TargetResourceGroupId,omitempty"`
}

func (s ListAutoGroupingRemediationsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAutoGroupingRemediationsRequest) GoString() string {
	return s.String()
}

func (s *ListAutoGroupingRemediationsRequest) GetEarliestRemediationTime() *string {
	return s.EarliestRemediationTime
}

func (s *ListAutoGroupingRemediationsRequest) GetLatestRemediationTime() *string {
	return s.LatestRemediationTime
}

func (s *ListAutoGroupingRemediationsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListAutoGroupingRemediationsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAutoGroupingRemediationsRequest) GetResourceId() *string {
	return s.ResourceId
}

func (s *ListAutoGroupingRemediationsRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListAutoGroupingRemediationsRequest) GetRuleId() *string {
	return s.RuleId
}

func (s *ListAutoGroupingRemediationsRequest) GetService() *string {
	return s.Service
}

func (s *ListAutoGroupingRemediationsRequest) GetTargetResourceGroupId() *string {
	return s.TargetResourceGroupId
}

func (s *ListAutoGroupingRemediationsRequest) SetEarliestRemediationTime(v string) *ListAutoGroupingRemediationsRequest {
	s.EarliestRemediationTime = &v
	return s
}

func (s *ListAutoGroupingRemediationsRequest) SetLatestRemediationTime(v string) *ListAutoGroupingRemediationsRequest {
	s.LatestRemediationTime = &v
	return s
}

func (s *ListAutoGroupingRemediationsRequest) SetMaxResults(v int32) *ListAutoGroupingRemediationsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListAutoGroupingRemediationsRequest) SetNextToken(v string) *ListAutoGroupingRemediationsRequest {
	s.NextToken = &v
	return s
}

func (s *ListAutoGroupingRemediationsRequest) SetResourceId(v string) *ListAutoGroupingRemediationsRequest {
	s.ResourceId = &v
	return s
}

func (s *ListAutoGroupingRemediationsRequest) SetResourceType(v string) *ListAutoGroupingRemediationsRequest {
	s.ResourceType = &v
	return s
}

func (s *ListAutoGroupingRemediationsRequest) SetRuleId(v string) *ListAutoGroupingRemediationsRequest {
	s.RuleId = &v
	return s
}

func (s *ListAutoGroupingRemediationsRequest) SetService(v string) *ListAutoGroupingRemediationsRequest {
	s.Service = &v
	return s
}

func (s *ListAutoGroupingRemediationsRequest) SetTargetResourceGroupId(v string) *ListAutoGroupingRemediationsRequest {
	s.TargetResourceGroupId = &v
	return s
}

func (s *ListAutoGroupingRemediationsRequest) Validate() error {
	return dara.Validate(s)
}
