// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCompliancePackTemplatesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCompliancePackTemplateId(v string) *ListCompliancePackTemplatesRequest
	GetCompliancePackTemplateId() *string
	SetFilterType(v string) *ListCompliancePackTemplatesRequest
	GetFilterType() *string
	SetPageNumber(v int32) *ListCompliancePackTemplatesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListCompliancePackTemplatesRequest
	GetPageSize() *int32
	SetResourceTypes(v string) *ListCompliancePackTemplatesRequest
	GetResourceTypes() *string
	SetRuleRiskLevel(v int32) *ListCompliancePackTemplatesRequest
	GetRuleRiskLevel() *int32
}

type ListCompliancePackTemplatesRequest struct {
	// The ID of the compliance pack template.
	//
	// example:
	//
	// ct-5f26ff4e06a300c4****
	CompliancePackTemplateId *string `json:"CompliancePackTemplateId,omitempty" xml:"CompliancePackTemplateId,omitempty"`
	// The field used to filter the query results.
	//
	// example:
	//
	// LAW
	FilterType *string `json:"FilterType,omitempty" xml:"FilterType,omitempty"`
	// The page number.
	//
	// Minimum value: 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page.
	//
	// Valid values: 1 to 100. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The resource type that is evaluated by the rules. If you specify this parameter, only the compliance pack templates that contain rules for the specified resource type are returned.
	//
	// example:
	//
	// ACS::ECS::Instance
	ResourceTypes *string `json:"ResourceTypes,omitempty" xml:"ResourceTypes,omitempty"`
	// The risk level of the rules in the compliance pack. Valid values:
	//
	// - 1: high
	//
	// - 2: medium
	//
	// - 3: low
	//
	// example:
	//
	// 2
	RuleRiskLevel *int32 `json:"RuleRiskLevel,omitempty" xml:"RuleRiskLevel,omitempty"`
}

func (s ListCompliancePackTemplatesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCompliancePackTemplatesRequest) GoString() string {
	return s.String()
}

func (s *ListCompliancePackTemplatesRequest) GetCompliancePackTemplateId() *string {
	return s.CompliancePackTemplateId
}

func (s *ListCompliancePackTemplatesRequest) GetFilterType() *string {
	return s.FilterType
}

func (s *ListCompliancePackTemplatesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListCompliancePackTemplatesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListCompliancePackTemplatesRequest) GetResourceTypes() *string {
	return s.ResourceTypes
}

func (s *ListCompliancePackTemplatesRequest) GetRuleRiskLevel() *int32 {
	return s.RuleRiskLevel
}

func (s *ListCompliancePackTemplatesRequest) SetCompliancePackTemplateId(v string) *ListCompliancePackTemplatesRequest {
	s.CompliancePackTemplateId = &v
	return s
}

func (s *ListCompliancePackTemplatesRequest) SetFilterType(v string) *ListCompliancePackTemplatesRequest {
	s.FilterType = &v
	return s
}

func (s *ListCompliancePackTemplatesRequest) SetPageNumber(v int32) *ListCompliancePackTemplatesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListCompliancePackTemplatesRequest) SetPageSize(v int32) *ListCompliancePackTemplatesRequest {
	s.PageSize = &v
	return s
}

func (s *ListCompliancePackTemplatesRequest) SetResourceTypes(v string) *ListCompliancePackTemplatesRequest {
	s.ResourceTypes = &v
	return s
}

func (s *ListCompliancePackTemplatesRequest) SetRuleRiskLevel(v int32) *ListCompliancePackTemplatesRequest {
	s.RuleRiskLevel = &v
	return s
}

func (s *ListCompliancePackTemplatesRequest) Validate() error {
	return dara.Validate(s)
}
