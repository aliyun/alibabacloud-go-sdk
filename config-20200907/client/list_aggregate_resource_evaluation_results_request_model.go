// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAggregateResourceEvaluationResultsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAggregatorId(v string) *ListAggregateResourceEvaluationResultsRequest
	GetAggregatorId() *string
	SetComplianceType(v string) *ListAggregateResourceEvaluationResultsRequest
	GetComplianceType() *string
	SetMaxResults(v int32) *ListAggregateResourceEvaluationResultsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListAggregateResourceEvaluationResultsRequest
	GetNextToken() *string
	SetRegion(v string) *ListAggregateResourceEvaluationResultsRequest
	GetRegion() *string
	SetResourceId(v string) *ListAggregateResourceEvaluationResultsRequest
	GetResourceId() *string
	SetResourceType(v string) *ListAggregateResourceEvaluationResultsRequest
	GetResourceType() *string
	SetRiskLevel(v int32) *ListAggregateResourceEvaluationResultsRequest
	GetRiskLevel() *int32
	SetSortBy(v string) *ListAggregateResourceEvaluationResultsRequest
	GetSortBy() *string
}

type ListAggregateResourceEvaluationResultsRequest struct {
	// The ID of the account group.
	//
	// For more information about how to obtain the ID of an account group, see [ListAggregators](https://help.aliyun.com/document_detail/255797.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// ca-7f00626622af0041****
	AggregatorId *string `json:"AggregatorId,omitempty" xml:"AggregatorId,omitempty"`
	// The compliance evaluation result of the resource. Valid values:
	//
	// 	- COMPLIANT: The resource is evaluated as compliant.
	//
	// 	- NON_COMPLIANT: The resource is evaluated as non-compliant.
	//
	// 	- NOT_APPLICABLE: The rule does not apply to the resource.
	//
	// 	- INSUFFICIENT_DATA: No data is available.
	//
	// 	- IGNORED: The resource is ignored during compliance evaluation.
	//
	// example:
	//
	// NON_COMPLIANT
	ComplianceType *string `json:"ComplianceType,omitempty" xml:"ComplianceType,omitempty"`
	// The maximum number of entries to return for a single request. Valid values: 1 to 100.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that you want to use to initiate the current request. If the response of the previous request is truncated, you can use this token to initiate another request and obtain the remaining entries.``
	//
	// example:
	//
	// IWBjqMYSy0is7zSMGu16****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the region where one or more resources that you want to query reside. For example, the value `global` indicates global regions and the value `cn-hangzhou` indicates the China (Hangzhou) region.
	//
	// For more information about how to obtain the ID of a region, see [ListAggregateDiscoveredResources](https://help.aliyun.com/document_detail/265983.html).
	//
	// example:
	//
	// global
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The ID of the resource.
	//
	// For more information about how to obtain the ID of a resource, see [ListAggregateDiscoveredResources](https://help.aliyun.com/document_detail/265983.html).
	//
	// example:
	//
	// 23642660635396****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The type of the resource.
	//
	// For more information about how to query the type of a resource, see [ListAggregateDiscoveredResources](https://help.aliyun.com/document_detail/265983.html).
	//
	// example:
	//
	// ACS::RAM::User
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	RiskLevel    *int32  `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	SortBy       *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
}

func (s ListAggregateResourceEvaluationResultsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAggregateResourceEvaluationResultsRequest) GoString() string {
	return s.String()
}

func (s *ListAggregateResourceEvaluationResultsRequest) GetAggregatorId() *string {
	return s.AggregatorId
}

func (s *ListAggregateResourceEvaluationResultsRequest) GetComplianceType() *string {
	return s.ComplianceType
}

func (s *ListAggregateResourceEvaluationResultsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListAggregateResourceEvaluationResultsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAggregateResourceEvaluationResultsRequest) GetRegion() *string {
	return s.Region
}

func (s *ListAggregateResourceEvaluationResultsRequest) GetResourceId() *string {
	return s.ResourceId
}

func (s *ListAggregateResourceEvaluationResultsRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListAggregateResourceEvaluationResultsRequest) GetRiskLevel() *int32 {
	return s.RiskLevel
}

func (s *ListAggregateResourceEvaluationResultsRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListAggregateResourceEvaluationResultsRequest) SetAggregatorId(v string) *ListAggregateResourceEvaluationResultsRequest {
	s.AggregatorId = &v
	return s
}

func (s *ListAggregateResourceEvaluationResultsRequest) SetComplianceType(v string) *ListAggregateResourceEvaluationResultsRequest {
	s.ComplianceType = &v
	return s
}

func (s *ListAggregateResourceEvaluationResultsRequest) SetMaxResults(v int32) *ListAggregateResourceEvaluationResultsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListAggregateResourceEvaluationResultsRequest) SetNextToken(v string) *ListAggregateResourceEvaluationResultsRequest {
	s.NextToken = &v
	return s
}

func (s *ListAggregateResourceEvaluationResultsRequest) SetRegion(v string) *ListAggregateResourceEvaluationResultsRequest {
	s.Region = &v
	return s
}

func (s *ListAggregateResourceEvaluationResultsRequest) SetResourceId(v string) *ListAggregateResourceEvaluationResultsRequest {
	s.ResourceId = &v
	return s
}

func (s *ListAggregateResourceEvaluationResultsRequest) SetResourceType(v string) *ListAggregateResourceEvaluationResultsRequest {
	s.ResourceType = &v
	return s
}

func (s *ListAggregateResourceEvaluationResultsRequest) SetRiskLevel(v int32) *ListAggregateResourceEvaluationResultsRequest {
	s.RiskLevel = &v
	return s
}

func (s *ListAggregateResourceEvaluationResultsRequest) SetSortBy(v string) *ListAggregateResourceEvaluationResultsRequest {
	s.SortBy = &v
	return s
}

func (s *ListAggregateResourceEvaluationResultsRequest) Validate() error {
	return dara.Validate(s)
}
