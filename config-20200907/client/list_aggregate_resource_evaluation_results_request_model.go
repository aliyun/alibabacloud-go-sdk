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
	// The compliance evaluation result. Valid values:
	//
	// - COMPLIANT: The resource is compliant.
	//
	// - NON_COMPLIANT: The resource is non-compliant.
	//
	// - NOT_APPLICABLE: The rule does not apply to the resource.
	//
	// - INSUFFICIENT_DATA: No data is available for the resource.
	//
	// - IGNORED: The evaluation result is ignored.
	//
	// example:
	//
	// NON_COMPLIANT
	ComplianceType *string `json:"ComplianceType,omitempty" xml:"ComplianceType,omitempty"`
	// The maximum number of entries to return on each page. Valid values: 1 to 100.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. If the response is truncated, you can use this token to initiate another request to retrieve the remaining records.
	//
	// example:
	//
	// IWBjqMYSy0is7zSMGu16****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the region where the resource resides. For example, `global` indicates Global and `cn-hangzhou` indicates China (Hangzhou).
	//
	// For more information about how to obtain the region ID of a resource, see [ListAggregateDiscoveredResources](https://help.aliyun.com/document_detail/265983.html).
	//
	// example:
	//
	// global
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The resource ID.
	//
	// For more information about how to obtain the resource ID, see [ListAggregateDiscoveredResources](https://help.aliyun.com/document_detail/265983.html).
	//
	// example:
	//
	// 23642660635396****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The resource type.
	//
	// For more information about how to obtain the resource type, see [ListAggregateDiscoveredResources](https://help.aliyun.com/document_detail/265983.html).
	//
	// example:
	//
	// ACS::RAM::User
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The risk level of the compliance package. Valid values:
	//
	// - 1: high
	//
	// - 2: medium
	//
	// - 3: low
	//
	// example:
	//
	// 1
	RiskLevel *int32 `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	// The sorting method. By default, this parameter is not specified. The supported value is `LastNonCompliantRecordTimestamp-Asc`. This value sorts resources based on the time when they first became non-compliant, in ascending order. You must set the `ComplianceType` parameter to `NON_COMPLIANT`.
	//
	// example:
	//
	// LastNonCompliantRecordTimestamp-Asc
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
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
