// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListConfigRuleEvaluationResultsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCompliancePackId(v string) *ListConfigRuleEvaluationResultsRequest
	GetCompliancePackId() *string
	SetComplianceType(v string) *ListConfigRuleEvaluationResultsRequest
	GetComplianceType() *string
	SetConfigRuleId(v string) *ListConfigRuleEvaluationResultsRequest
	GetConfigRuleId() *string
	SetMaxResults(v int32) *ListConfigRuleEvaluationResultsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListConfigRuleEvaluationResultsRequest
	GetNextToken() *string
	SetRegions(v string) *ListConfigRuleEvaluationResultsRequest
	GetRegions() *string
	SetResourceGroupIds(v string) *ListConfigRuleEvaluationResultsRequest
	GetResourceGroupIds() *string
	SetResourceTypes(v string) *ListConfigRuleEvaluationResultsRequest
	GetResourceTypes() *string
}

type ListConfigRuleEvaluationResultsRequest struct {
	// The compliance package ID.
	//
	// For more information about how to obtain a compliance package ID, see [ListCompliancePacks](https://help.aliyun.com/document_detail/263332.html).
	//
	// example:
	//
	// cp-f1e3326622af00cb****
	CompliancePackId *string `json:"CompliancePackId,omitempty" xml:"CompliancePackId,omitempty"`
	// The compliance evaluation result. Valid values:
	//
	// - COMPLIANT: The resource is compliant.
	//
	// - NON_COMPLIANT: The resource is non-compliant.
	//
	// - NOT_APPLICABLE: The rule does not apply to the resource.
	//
	// - INSUFFICIENT_DATA: No data is available.
	//
	// - IGNORED: The evaluation result is ignored.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// NON_COMPLIANT
	ComplianceType *string `json:"ComplianceType,omitempty" xml:"ComplianceType,omitempty"`
	// The rule ID.
	//
	// For more information about how to obtain a rule ID, see [ListConfigRules](https://help.aliyun.com/document_detail/169607.html).
	//
	// example:
	//
	// cr-cac56457e0d900d3****
	ConfigRuleId *string `json:"ConfigRuleId,omitempty" xml:"ConfigRuleId,omitempty"`
	// The maximum number of entries to return on each page. Valid values: 1 to 100.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// If the response is truncated, use the `NextToken` to retrieve the next page of results.
	//
	// example:
	//
	// IWBjqMYSy0is7zSMGu16****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The region where the evaluated resource resides. Separate multiple regions with commas (,).
	//
	// example:
	//
	// cn-shanghai
	Regions *string `json:"Regions,omitempty" xml:"Regions,omitempty"`
	// The ID of the resource group to which the evaluated resource belongs. Separate multiple resource group IDs with commas (,).
	//
	// example:
	//
	// rg-aek2indxn3g****
	ResourceGroupIds *string `json:"ResourceGroupIds,omitempty" xml:"ResourceGroupIds,omitempty"`
	// The type of the evaluated resource. Separate multiple resource types with commas (,).
	//
	// example:
	//
	// ACS::ECS::Instance
	ResourceTypes *string `json:"ResourceTypes,omitempty" xml:"ResourceTypes,omitempty"`
}

func (s ListConfigRuleEvaluationResultsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListConfigRuleEvaluationResultsRequest) GoString() string {
	return s.String()
}

func (s *ListConfigRuleEvaluationResultsRequest) GetCompliancePackId() *string {
	return s.CompliancePackId
}

func (s *ListConfigRuleEvaluationResultsRequest) GetComplianceType() *string {
	return s.ComplianceType
}

func (s *ListConfigRuleEvaluationResultsRequest) GetConfigRuleId() *string {
	return s.ConfigRuleId
}

func (s *ListConfigRuleEvaluationResultsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListConfigRuleEvaluationResultsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListConfigRuleEvaluationResultsRequest) GetRegions() *string {
	return s.Regions
}

func (s *ListConfigRuleEvaluationResultsRequest) GetResourceGroupIds() *string {
	return s.ResourceGroupIds
}

func (s *ListConfigRuleEvaluationResultsRequest) GetResourceTypes() *string {
	return s.ResourceTypes
}

func (s *ListConfigRuleEvaluationResultsRequest) SetCompliancePackId(v string) *ListConfigRuleEvaluationResultsRequest {
	s.CompliancePackId = &v
	return s
}

func (s *ListConfigRuleEvaluationResultsRequest) SetComplianceType(v string) *ListConfigRuleEvaluationResultsRequest {
	s.ComplianceType = &v
	return s
}

func (s *ListConfigRuleEvaluationResultsRequest) SetConfigRuleId(v string) *ListConfigRuleEvaluationResultsRequest {
	s.ConfigRuleId = &v
	return s
}

func (s *ListConfigRuleEvaluationResultsRequest) SetMaxResults(v int32) *ListConfigRuleEvaluationResultsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListConfigRuleEvaluationResultsRequest) SetNextToken(v string) *ListConfigRuleEvaluationResultsRequest {
	s.NextToken = &v
	return s
}

func (s *ListConfigRuleEvaluationResultsRequest) SetRegions(v string) *ListConfigRuleEvaluationResultsRequest {
	s.Regions = &v
	return s
}

func (s *ListConfigRuleEvaluationResultsRequest) SetResourceGroupIds(v string) *ListConfigRuleEvaluationResultsRequest {
	s.ResourceGroupIds = &v
	return s
}

func (s *ListConfigRuleEvaluationResultsRequest) SetResourceTypes(v string) *ListConfigRuleEvaluationResultsRequest {
	s.ResourceTypes = &v
	return s
}

func (s *ListConfigRuleEvaluationResultsRequest) Validate() error {
	return dara.Validate(s)
}
