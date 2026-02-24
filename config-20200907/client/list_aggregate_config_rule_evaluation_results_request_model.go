// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAggregateConfigRuleEvaluationResultsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAggregatorId(v string) *ListAggregateConfigRuleEvaluationResultsRequest
	GetAggregatorId() *string
	SetCompliancePackId(v string) *ListAggregateConfigRuleEvaluationResultsRequest
	GetCompliancePackId() *string
	SetComplianceType(v string) *ListAggregateConfigRuleEvaluationResultsRequest
	GetComplianceType() *string
	SetConfigRuleId(v string) *ListAggregateConfigRuleEvaluationResultsRequest
	GetConfigRuleId() *string
	SetMaxResults(v int32) *ListAggregateConfigRuleEvaluationResultsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListAggregateConfigRuleEvaluationResultsRequest
	GetNextToken() *string
	SetRegions(v string) *ListAggregateConfigRuleEvaluationResultsRequest
	GetRegions() *string
	SetResourceAccountId(v int64) *ListAggregateConfigRuleEvaluationResultsRequest
	GetResourceAccountId() *int64
	SetResourceGroupIds(v string) *ListAggregateConfigRuleEvaluationResultsRequest
	GetResourceGroupIds() *string
	SetResourceOwnerId(v int64) *ListAggregateConfigRuleEvaluationResultsRequest
	GetResourceOwnerId() *int64
	SetResourceTypes(v string) *ListAggregateConfigRuleEvaluationResultsRequest
	GetResourceTypes() *string
}

type ListAggregateConfigRuleEvaluationResultsRequest struct {
	// The ID of the account group.
	//
	// For more information about how to obtain the ID of an account group, see [ListAggregators](https://help.aliyun.com/document_detail/255797.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// ca-b1e6626622af00cb****
	AggregatorId *string `json:"AggregatorId,omitempty" xml:"AggregatorId,omitempty"`
	// The ID of the compliance package.
	//
	// For more information about how to obtain the ID of a compliance package, see [ListAggregateCompliancePacks](https://help.aliyun.com/document_detail/262059.html).
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
	// example:
	//
	// NON_COMPLIANT
	ComplianceType *string `json:"ComplianceType,omitempty" xml:"ComplianceType,omitempty"`
	// The rule ID.
	//
	// For more information about how to obtain a rule ID, see [ListAggregateConfigRules](https://help.aliyun.com/document_detail/264148.html).
	//
	// example:
	//
	// cr-888f626622af00ae****
	ConfigRuleId *string `json:"ConfigRuleId,omitempty" xml:"ConfigRuleId,omitempty"`
	// The maximum number of entries to return on a single page. Valid values: 1 to 100.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is returned when the response is truncated. Use this token in the next request to retrieve the next page of results.
	//
	// example:
	//
	// IWBjqMYSy0is7zSMGu16****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the region where the resource resides. Separate multiple region IDs with commas (,).
	//
	// example:
	//
	// cn-shanghai
	Regions *string `json:"Regions,omitempty" xml:"Regions,omitempty"`
	// The ID of the Alibaba Cloud account to which the resource in the account group belongs.
	//
	// > Specify this parameter or \\`ResourceOwnerId\\`. We recommend that you specify this parameter.
	//
	// example:
	//
	// 100931896542****
	ResourceAccountId *int64 `json:"ResourceAccountId,omitempty" xml:"ResourceAccountId,omitempty"`
	// The ID of the resource group to which the resource belongs. Separate multiple resource group IDs with commas (,).
	//
	// example:
	//
	// rg-aek2cqyzvuj****
	ResourceGroupIds *string `json:"ResourceGroupIds,omitempty" xml:"ResourceGroupIds,omitempty"`
	// Deprecated
	ResourceOwnerId *int64 `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The resource type. Separate multiple resource types with commas (,).
	//
	// example:
	//
	// ACS::ECS::Instance
	ResourceTypes *string `json:"ResourceTypes,omitempty" xml:"ResourceTypes,omitempty"`
}

func (s ListAggregateConfigRuleEvaluationResultsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAggregateConfigRuleEvaluationResultsRequest) GoString() string {
	return s.String()
}

func (s *ListAggregateConfigRuleEvaluationResultsRequest) GetAggregatorId() *string {
	return s.AggregatorId
}

func (s *ListAggregateConfigRuleEvaluationResultsRequest) GetCompliancePackId() *string {
	return s.CompliancePackId
}

func (s *ListAggregateConfigRuleEvaluationResultsRequest) GetComplianceType() *string {
	return s.ComplianceType
}

func (s *ListAggregateConfigRuleEvaluationResultsRequest) GetConfigRuleId() *string {
	return s.ConfigRuleId
}

func (s *ListAggregateConfigRuleEvaluationResultsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListAggregateConfigRuleEvaluationResultsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAggregateConfigRuleEvaluationResultsRequest) GetRegions() *string {
	return s.Regions
}

func (s *ListAggregateConfigRuleEvaluationResultsRequest) GetResourceAccountId() *int64 {
	return s.ResourceAccountId
}

func (s *ListAggregateConfigRuleEvaluationResultsRequest) GetResourceGroupIds() *string {
	return s.ResourceGroupIds
}

func (s *ListAggregateConfigRuleEvaluationResultsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ListAggregateConfigRuleEvaluationResultsRequest) GetResourceTypes() *string {
	return s.ResourceTypes
}

func (s *ListAggregateConfigRuleEvaluationResultsRequest) SetAggregatorId(v string) *ListAggregateConfigRuleEvaluationResultsRequest {
	s.AggregatorId = &v
	return s
}

func (s *ListAggregateConfigRuleEvaluationResultsRequest) SetCompliancePackId(v string) *ListAggregateConfigRuleEvaluationResultsRequest {
	s.CompliancePackId = &v
	return s
}

func (s *ListAggregateConfigRuleEvaluationResultsRequest) SetComplianceType(v string) *ListAggregateConfigRuleEvaluationResultsRequest {
	s.ComplianceType = &v
	return s
}

func (s *ListAggregateConfigRuleEvaluationResultsRequest) SetConfigRuleId(v string) *ListAggregateConfigRuleEvaluationResultsRequest {
	s.ConfigRuleId = &v
	return s
}

func (s *ListAggregateConfigRuleEvaluationResultsRequest) SetMaxResults(v int32) *ListAggregateConfigRuleEvaluationResultsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListAggregateConfigRuleEvaluationResultsRequest) SetNextToken(v string) *ListAggregateConfigRuleEvaluationResultsRequest {
	s.NextToken = &v
	return s
}

func (s *ListAggregateConfigRuleEvaluationResultsRequest) SetRegions(v string) *ListAggregateConfigRuleEvaluationResultsRequest {
	s.Regions = &v
	return s
}

func (s *ListAggregateConfigRuleEvaluationResultsRequest) SetResourceAccountId(v int64) *ListAggregateConfigRuleEvaluationResultsRequest {
	s.ResourceAccountId = &v
	return s
}

func (s *ListAggregateConfigRuleEvaluationResultsRequest) SetResourceGroupIds(v string) *ListAggregateConfigRuleEvaluationResultsRequest {
	s.ResourceGroupIds = &v
	return s
}

func (s *ListAggregateConfigRuleEvaluationResultsRequest) SetResourceOwnerId(v int64) *ListAggregateConfigRuleEvaluationResultsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListAggregateConfigRuleEvaluationResultsRequest) SetResourceTypes(v string) *ListAggregateConfigRuleEvaluationResultsRequest {
	s.ResourceTypes = &v
	return s
}

func (s *ListAggregateConfigRuleEvaluationResultsRequest) Validate() error {
	return dara.Validate(s)
}
