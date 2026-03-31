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
	// The ID of the compliance package template.
	//
	// example:
	//
	// ct-d254ff4e06a300cf****
	CompliancePackTemplateId *string `json:"CompliancePackTemplateId,omitempty" xml:"CompliancePackTemplateId,omitempty"`
	// example:
	//
	// LAW
	FilterType *string `json:"FilterType,omitempty" xml:"FilterType,omitempty"`
	// The page number.
	//
	// Pages start from page 1. Default value: 1
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// Valid values: 1 to 100. Minimum value: 1. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The types of the resources evaluated based on the rule. If you configure this parameter, only the rules that include the resource types in the compliance package template are returned.
	//
	// example:
	//
	// ACS::ECS::Instance
	ResourceTypes *string `json:"ResourceTypes,omitempty" xml:"ResourceTypes,omitempty"`
	RuleRiskLevel *int32  `json:"RuleRiskLevel,omitempty" xml:"RuleRiskLevel,omitempty"`
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
