// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListResourceEvaluationResultsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetComplianceType(v string) *ListResourceEvaluationResultsRequest
	GetComplianceType() *string
	SetMaxResults(v int32) *ListResourceEvaluationResultsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListResourceEvaluationResultsRequest
	GetNextToken() *string
	SetRegion(v string) *ListResourceEvaluationResultsRequest
	GetRegion() *string
	SetResourceId(v string) *ListResourceEvaluationResultsRequest
	GetResourceId() *string
	SetResourceType(v string) *ListResourceEvaluationResultsRequest
	GetResourceType() *string
	SetRiskLevel(v int32) *ListResourceEvaluationResultsRequest
	GetRiskLevel() *int32
	SetSortBy(v string) *ListResourceEvaluationResultsRequest
	GetSortBy() *string
}

type ListResourceEvaluationResultsRequest struct {
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
	// The maximum number of entries to return for a single request. Valid values: 1 to 100.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// If the response is truncated, pass the `NextToken` value in a subsequent request to retrieve the remaining results.
	//
	// example:
	//
	// IWBjqMYSy0is7zSMGu16****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the region where the resource resides. For example, `global` indicates Global, and `cn-hangzhou` indicates China (Hangzhou).
	//
	// For more information, see [ListDiscoveredResources](https://help.aliyun.com/document_detail/169620.html).
	//
	// example:
	//
	// global
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The resource ID.
	//
	// For more information, see [ListDiscoveredResources](https://help.aliyun.com/document_detail/169620.html).
	//
	// example:
	//
	// 23642660635396****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The resource type.
	//
	// For more information, see [ListDiscoveredResources](https://help.aliyun.com/document_detail/169620.html).
	//
	// example:
	//
	// ACS::RAM::User
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The risk level of the rule. Valid values:
	//
	// - 1: High risk.
	//
	// - 2: Medium risk.
	//
	// - 3: Low risk.
	//
	// example:
	//
	// 1
	RiskLevel *int32 `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	// This parameter is optional. The only supported value is `LastNonCompliantRecordTimestamp-Asc`. This value sorts resources based on when they first became non-compliant, with the earliest appearing first. Use this parameter only when you set the `ComplianceType` parameter to `NON_COMPLIANT`.
	//
	// example:
	//
	// LastNonCompliantRecordTimestamp-Asc
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
}

func (s ListResourceEvaluationResultsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListResourceEvaluationResultsRequest) GoString() string {
	return s.String()
}

func (s *ListResourceEvaluationResultsRequest) GetComplianceType() *string {
	return s.ComplianceType
}

func (s *ListResourceEvaluationResultsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListResourceEvaluationResultsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListResourceEvaluationResultsRequest) GetRegion() *string {
	return s.Region
}

func (s *ListResourceEvaluationResultsRequest) GetResourceId() *string {
	return s.ResourceId
}

func (s *ListResourceEvaluationResultsRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListResourceEvaluationResultsRequest) GetRiskLevel() *int32 {
	return s.RiskLevel
}

func (s *ListResourceEvaluationResultsRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListResourceEvaluationResultsRequest) SetComplianceType(v string) *ListResourceEvaluationResultsRequest {
	s.ComplianceType = &v
	return s
}

func (s *ListResourceEvaluationResultsRequest) SetMaxResults(v int32) *ListResourceEvaluationResultsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListResourceEvaluationResultsRequest) SetNextToken(v string) *ListResourceEvaluationResultsRequest {
	s.NextToken = &v
	return s
}

func (s *ListResourceEvaluationResultsRequest) SetRegion(v string) *ListResourceEvaluationResultsRequest {
	s.Region = &v
	return s
}

func (s *ListResourceEvaluationResultsRequest) SetResourceId(v string) *ListResourceEvaluationResultsRequest {
	s.ResourceId = &v
	return s
}

func (s *ListResourceEvaluationResultsRequest) SetResourceType(v string) *ListResourceEvaluationResultsRequest {
	s.ResourceType = &v
	return s
}

func (s *ListResourceEvaluationResultsRequest) SetRiskLevel(v int32) *ListResourceEvaluationResultsRequest {
	s.RiskLevel = &v
	return s
}

func (s *ListResourceEvaluationResultsRequest) SetSortBy(v string) *ListResourceEvaluationResultsRequest {
	s.SortBy = &v
	return s
}

func (s *ListResourceEvaluationResultsRequest) Validate() error {
	return dara.Validate(s)
}
