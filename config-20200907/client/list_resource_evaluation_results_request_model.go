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
	// The compliance evaluation result of the resource. Valid values:
	//
	// 	- COMPLIANT: The resource is evaluated as compliant.
	//
	// 	- NON_COMPLIANT: The resource is evaluated as non-compliant.
	//
	// 	- NOT_APPLICABLE: The rule does not apply to the resources.
	//
	// 	- INSUFFICIENT_DATA: No data is available.
	//
	// 	- IGNORED: The resource is ignored during compliance evaluation.
	//
	// example:
	//
	// NON_COMPLIANT
	ComplianceType *string `json:"ComplianceType,omitempty" xml:"ComplianceType,omitempty"`
	// The maximum number of entries to return in a request. Valid values: 1 to 100.
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
	// The ID of the region where one or more resources you want to query reside. For example, the value `global` indicates global regions and the value `cn-hangzhou` indicates the China (Hangzhou) region.
	//
	// For more information about how to obtain the ID of the region where a resource resides, see [ListDiscoveredResources](https://help.aliyun.com/document_detail/169620.html).
	//
	// example:
	//
	// global
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The ID of the resource.
	//
	// For more information about how to obtain the ID of a resource, see [ListDiscoveredResources](https://help.aliyun.com/document_detail/169620.html).
	//
	// example:
	//
	// 23642660635396****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The type of the resource.
	//
	// For more information about how to query the type of a resource, see [ListDiscoveredResources](https://help.aliyun.com/document_detail/169620.html).
	//
	// example:
	//
	// ACS::RAM::User
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	RiskLevel    *int32  `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	SortBy       *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
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
