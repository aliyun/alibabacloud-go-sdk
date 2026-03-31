// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCompliancePackTemplatesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCompliancePackTemplatesResult(v *ListCompliancePackTemplatesResponseBodyCompliancePackTemplatesResult) *ListCompliancePackTemplatesResponseBody
	GetCompliancePackTemplatesResult() *ListCompliancePackTemplatesResponseBodyCompliancePackTemplatesResult
	SetRequestId(v string) *ListCompliancePackTemplatesResponseBody
	GetRequestId() *string
}

type ListCompliancePackTemplatesResponseBody struct {
	// The information about the compliance package templates returned.
	CompliancePackTemplatesResult *ListCompliancePackTemplatesResponseBodyCompliancePackTemplatesResult `json:"CompliancePackTemplatesResult,omitempty" xml:"CompliancePackTemplatesResult,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// D67FC82F-25AE-4268-A94C-3348340748F9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListCompliancePackTemplatesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListCompliancePackTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *ListCompliancePackTemplatesResponseBody) GetCompliancePackTemplatesResult() *ListCompliancePackTemplatesResponseBodyCompliancePackTemplatesResult {
	return s.CompliancePackTemplatesResult
}

func (s *ListCompliancePackTemplatesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListCompliancePackTemplatesResponseBody) SetCompliancePackTemplatesResult(v *ListCompliancePackTemplatesResponseBodyCompliancePackTemplatesResult) *ListCompliancePackTemplatesResponseBody {
	s.CompliancePackTemplatesResult = v
	return s
}

func (s *ListCompliancePackTemplatesResponseBody) SetRequestId(v string) *ListCompliancePackTemplatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCompliancePackTemplatesResponseBody) Validate() error {
	if s.CompliancePackTemplatesResult != nil {
		if err := s.CompliancePackTemplatesResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListCompliancePackTemplatesResponseBodyCompliancePackTemplatesResult struct {
	// The compliance package templates.
	CompliancePackTemplates []*ListCompliancePackTemplatesResponseBodyCompliancePackTemplatesResultCompliancePackTemplates `json:"CompliancePackTemplates,omitempty" xml:"CompliancePackTemplates,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of the compliance package templates returned.
	//
	// example:
	//
	// 1
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListCompliancePackTemplatesResponseBodyCompliancePackTemplatesResult) String() string {
	return dara.Prettify(s)
}

func (s ListCompliancePackTemplatesResponseBodyCompliancePackTemplatesResult) GoString() string {
	return s.String()
}

func (s *ListCompliancePackTemplatesResponseBodyCompliancePackTemplatesResult) GetCompliancePackTemplates() []*ListCompliancePackTemplatesResponseBodyCompliancePackTemplatesResultCompliancePackTemplates {
	return s.CompliancePackTemplates
}

func (s *ListCompliancePackTemplatesResponseBodyCompliancePackTemplatesResult) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListCompliancePackTemplatesResponseBodyCompliancePackTemplatesResult) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListCompliancePackTemplatesResponseBodyCompliancePackTemplatesResult) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListCompliancePackTemplatesResponseBodyCompliancePackTemplatesResult) SetCompliancePackTemplates(v []*ListCompliancePackTemplatesResponseBodyCompliancePackTemplatesResultCompliancePackTemplates) *ListCompliancePackTemplatesResponseBodyCompliancePackTemplatesResult {
	s.CompliancePackTemplates = v
	return s
}

func (s *ListCompliancePackTemplatesResponseBodyCompliancePackTemplatesResult) SetPageNumber(v int32) *ListCompliancePackTemplatesResponseBodyCompliancePackTemplatesResult {
	s.PageNumber = &v
	return s
}

func (s *ListCompliancePackTemplatesResponseBodyCompliancePackTemplatesResult) SetPageSize(v int32) *ListCompliancePackTemplatesResponseBodyCompliancePackTemplatesResult {
	s.PageSize = &v
	return s
}

func (s *ListCompliancePackTemplatesResponseBodyCompliancePackTemplatesResult) SetTotalCount(v int64) *ListCompliancePackTemplatesResponseBodyCompliancePackTemplatesResult {
	s.TotalCount = &v
	return s
}

func (s *ListCompliancePackTemplatesResponseBodyCompliancePackTemplatesResult) Validate() error {
	if s.CompliancePackTemplates != nil {
		for _, item := range s.CompliancePackTemplates {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListCompliancePackTemplatesResponseBodyCompliancePackTemplatesResultCompliancePackTemplates struct {
	// The ID of the compliance package template.
	//
	// example:
	//
	// ct-d254ff4e06a300cf****
	CompliancePackTemplateId *string `json:"CompliancePackTemplateId,omitempty" xml:"CompliancePackTemplateId,omitempty"`
	// The name of the compliance package template.
	//
	// example:
	//
	// BestPracticesForResourceStability
	CompliancePackTemplateName *string `json:"CompliancePackTemplateName,omitempty" xml:"CompliancePackTemplateName,omitempty"`
	// The default rules in the compliance package.
	ConfigRules []*ListCompliancePackTemplatesResponseBodyCompliancePackTemplatesResultCompliancePackTemplatesConfigRules `json:"ConfigRules,omitempty" xml:"ConfigRules,omitempty" type:"Repeated"`
	// The description of the compliance package.
	//
	// example:
	//
	// example-description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The tag of the compliance package.
	//
	// example:
	//
	// tagKey-1
	Labels *string `json:"Labels,omitempty" xml:"Labels,omitempty"`
	// The time when the compliance package was last updated.
	//
	// example:
	//
	// 1663408308
	LastUpdate *int32 `json:"LastUpdate,omitempty" xml:"LastUpdate,omitempty"`
	// The risk level of the managed rule in the compliance package. Valid values:
	//
	// 	- 1: high
	//
	// 	- 2: medium
	//
	// 	- 3: low
	//
	// example:
	//
	// 1
	RiskLevel *int32 `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
}

func (s ListCompliancePackTemplatesResponseBodyCompliancePackTemplatesResultCompliancePackTemplates) String() string {
	return dara.Prettify(s)
}

func (s ListCompliancePackTemplatesResponseBodyCompliancePackTemplatesResultCompliancePackTemplates) GoString() string {
	return s.String()
}

func (s *ListCompliancePackTemplatesResponseBodyCompliancePackTemplatesResultCompliancePackTemplates) GetCompliancePackTemplateId() *string {
	return s.CompliancePackTemplateId
}

func (s *ListCompliancePackTemplatesResponseBodyCompliancePackTemplatesResultCompliancePackTemplates) GetCompliancePackTemplateName() *string {
	return s.CompliancePackTemplateName
}

func (s *ListCompliancePackTemplatesResponseBodyCompliancePackTemplatesResultCompliancePackTemplates) GetConfigRules() []*ListCompliancePackTemplatesResponseBodyCompliancePackTemplatesResultCompliancePackTemplatesConfigRules {
	return s.ConfigRules
}

func (s *ListCompliancePackTemplatesResponseBodyCompliancePackTemplatesResultCompliancePackTemplates) GetDescription() *string {
	return s.Description
}

func (s *ListCompliancePackTemplatesResponseBodyCompliancePackTemplatesResultCompliancePackTemplates) GetLabels() *string {
	return s.Labels
}

func (s *ListCompliancePackTemplatesResponseBodyCompliancePackTemplatesResultCompliancePackTemplates) GetLastUpdate() *int32 {
	return s.LastUpdate
}

func (s *ListCompliancePackTemplatesResponseBodyCompliancePackTemplatesResultCompliancePackTemplates) GetRiskLevel() *int32 {
	return s.RiskLevel
}

func (s *ListCompliancePackTemplatesResponseBodyCompliancePackTemplatesResultCompliancePackTemplates) SetCompliancePackTemplateId(v string) *ListCompliancePackTemplatesResponseBodyCompliancePackTemplatesResultCompliancePackTemplates {
	s.CompliancePackTemplateId = &v
	return s
}

func (s *ListCompliancePackTemplatesResponseBodyCompliancePackTemplatesResultCompliancePackTemplates) SetCompliancePackTemplateName(v string) *ListCompliancePackTemplatesResponseBodyCompliancePackTemplatesResultCompliancePackTemplates {
	s.CompliancePackTemplateName = &v
	return s
}

func (s *ListCompliancePackTemplatesResponseBodyCompliancePackTemplatesResultCompliancePackTemplates) SetConfigRules(v []*ListCompliancePackTemplatesResponseBodyCompliancePackTemplatesResultCompliancePackTemplatesConfigRules) *ListCompliancePackTemplatesResponseBodyCompliancePackTemplatesResultCompliancePackTemplates {
	s.ConfigRules = v
	return s
}

func (s *ListCompliancePackTemplatesResponseBodyCompliancePackTemplatesResultCompliancePackTemplates) SetDescription(v string) *ListCompliancePackTemplatesResponseBodyCompliancePackTemplatesResultCompliancePackTemplates {
	s.Description = &v
	return s
}

func (s *ListCompliancePackTemplatesResponseBodyCompliancePackTemplatesResultCompliancePackTemplates) SetLabels(v string) *ListCompliancePackTemplatesResponseBodyCompliancePackTemplatesResultCompliancePackTemplates {
	s.Labels = &v
	return s
}

func (s *ListCompliancePackTemplatesResponseBodyCompliancePackTemplatesResultCompliancePackTemplates) SetLastUpdate(v int32) *ListCompliancePackTemplatesResponseBodyCompliancePackTemplatesResultCompliancePackTemplates {
	s.LastUpdate = &v
	return s
}

func (s *ListCompliancePackTemplatesResponseBodyCompliancePackTemplatesResultCompliancePackTemplates) SetRiskLevel(v int32) *ListCompliancePackTemplatesResponseBodyCompliancePackTemplatesResultCompliancePackTemplates {
	s.RiskLevel = &v
	return s
}

func (s *ListCompliancePackTemplatesResponseBodyCompliancePackTemplatesResultCompliancePackTemplates) Validate() error {
	if s.ConfigRules != nil {
		for _, item := range s.ConfigRules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListCompliancePackTemplatesResponseBodyCompliancePackTemplatesResultCompliancePackTemplatesConfigRules struct {
	// The input parameter of the managed rule.
	ConfigRuleParameters []*ListCompliancePackTemplatesResponseBodyCompliancePackTemplatesResultCompliancePackTemplatesConfigRulesConfigRuleParameters `json:"ConfigRuleParameters,omitempty" xml:"ConfigRuleParameters,omitempty" type:"Repeated"`
	// The description of the regulation. This parameter is available only for regulation compliance packages.
	//
	// example:
	//
	// No classic networks exist.
	ControlDescription *string `json:"ControlDescription,omitempty" xml:"ControlDescription,omitempty"`
	// The regulation ID.
	//
	// >  This parameter is available only for regulation compliance packages.
	//
	// example:
	//
	// 3.1
	ControlId *string `json:"ControlId,omitempty" xml:"ControlId,omitempty"`
	// Indicates whether the rules are enabled together with the compliance package. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// false
	DefaultEnable *bool `json:"DefaultEnable,omitempty" xml:"DefaultEnable,omitempty"`
	// The description of the rule.
	//
	// example:
	//
	// If the expiration time of the SLB certificate is later than the specified number of days after the check time, the configuration is considered compliant. Default value: 90 days.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The identifier of the managed rule.
	//
	// example:
	//
	// slb-servercertificate-expired-check
	ManagedRuleIdentifier *string `json:"ManagedRuleIdentifier,omitempty" xml:"ManagedRuleIdentifier,omitempty"`
	// The name of the managed rule.
	//
	// example:
	//
	// slb-servercertificate-expired-check
	ManagedRuleName *string `json:"ManagedRuleName,omitempty" xml:"ManagedRuleName,omitempty"`
	// The types of the resources evaluated based on the rule.
	//
	// example:
	//
	// ACS::SLB::ServerCertificate
	ResourceTypesScope *string `json:"ResourceTypesScope,omitempty" xml:"ResourceTypesScope,omitempty"`
	// The risk level of the managed rule. Valid values:
	//
	// 	- 1: high
	//
	// 	- 2: medium
	//
	// 	- 3: low
	//
	// example:
	//
	// 1
	RiskLevel *int32 `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
}

func (s ListCompliancePackTemplatesResponseBodyCompliancePackTemplatesResultCompliancePackTemplatesConfigRules) String() string {
	return dara.Prettify(s)
}

func (s ListCompliancePackTemplatesResponseBodyCompliancePackTemplatesResultCompliancePackTemplatesConfigRules) GoString() string {
	return s.String()
}

func (s *ListCompliancePackTemplatesResponseBodyCompliancePackTemplatesResultCompliancePackTemplatesConfigRules) GetConfigRuleParameters() []*ListCompliancePackTemplatesResponseBodyCompliancePackTemplatesResultCompliancePackTemplatesConfigRulesConfigRuleParameters {
	return s.ConfigRuleParameters
}

func (s *ListCompliancePackTemplatesResponseBodyCompliancePackTemplatesResultCompliancePackTemplatesConfigRules) GetControlDescription() *string {
	return s.ControlDescription
}

func (s *ListCompliancePackTemplatesResponseBodyCompliancePackTemplatesResultCompliancePackTemplatesConfigRules) GetControlId() *string {
	return s.ControlId
}

func (s *ListCompliancePackTemplatesResponseBodyCompliancePackTemplatesResultCompliancePackTemplatesConfigRules) GetDefaultEnable() *bool {
	return s.DefaultEnable
}

func (s *ListCompliancePackTemplatesResponseBodyCompliancePackTemplatesResultCompliancePackTemplatesConfigRules) GetDescription() *string {
	return s.Description
}

func (s *ListCompliancePackTemplatesResponseBodyCompliancePackTemplatesResultCompliancePackTemplatesConfigRules) GetManagedRuleIdentifier() *string {
	return s.ManagedRuleIdentifier
}

func (s *ListCompliancePackTemplatesResponseBodyCompliancePackTemplatesResultCompliancePackTemplatesConfigRules) GetManagedRuleName() *string {
	return s.ManagedRuleName
}

func (s *ListCompliancePackTemplatesResponseBodyCompliancePackTemplatesResultCompliancePackTemplatesConfigRules) GetResourceTypesScope() *string {
	return s.ResourceTypesScope
}

func (s *ListCompliancePackTemplatesResponseBodyCompliancePackTemplatesResultCompliancePackTemplatesConfigRules) GetRiskLevel() *int32 {
	return s.RiskLevel
}

func (s *ListCompliancePackTemplatesResponseBodyCompliancePackTemplatesResultCompliancePackTemplatesConfigRules) SetConfigRuleParameters(v []*ListCompliancePackTemplatesResponseBodyCompliancePackTemplatesResultCompliancePackTemplatesConfigRulesConfigRuleParameters) *ListCompliancePackTemplatesResponseBodyCompliancePackTemplatesResultCompliancePackTemplatesConfigRules {
	s.ConfigRuleParameters = v
	return s
}

func (s *ListCompliancePackTemplatesResponseBodyCompliancePackTemplatesResultCompliancePackTemplatesConfigRules) SetControlDescription(v string) *ListCompliancePackTemplatesResponseBodyCompliancePackTemplatesResultCompliancePackTemplatesConfigRules {
	s.ControlDescription = &v
	return s
}

func (s *ListCompliancePackTemplatesResponseBodyCompliancePackTemplatesResultCompliancePackTemplatesConfigRules) SetControlId(v string) *ListCompliancePackTemplatesResponseBodyCompliancePackTemplatesResultCompliancePackTemplatesConfigRules {
	s.ControlId = &v
	return s
}

func (s *ListCompliancePackTemplatesResponseBodyCompliancePackTemplatesResultCompliancePackTemplatesConfigRules) SetDefaultEnable(v bool) *ListCompliancePackTemplatesResponseBodyCompliancePackTemplatesResultCompliancePackTemplatesConfigRules {
	s.DefaultEnable = &v
	return s
}

func (s *ListCompliancePackTemplatesResponseBodyCompliancePackTemplatesResultCompliancePackTemplatesConfigRules) SetDescription(v string) *ListCompliancePackTemplatesResponseBodyCompliancePackTemplatesResultCompliancePackTemplatesConfigRules {
	s.Description = &v
	return s
}

func (s *ListCompliancePackTemplatesResponseBodyCompliancePackTemplatesResultCompliancePackTemplatesConfigRules) SetManagedRuleIdentifier(v string) *ListCompliancePackTemplatesResponseBodyCompliancePackTemplatesResultCompliancePackTemplatesConfigRules {
	s.ManagedRuleIdentifier = &v
	return s
}

func (s *ListCompliancePackTemplatesResponseBodyCompliancePackTemplatesResultCompliancePackTemplatesConfigRules) SetManagedRuleName(v string) *ListCompliancePackTemplatesResponseBodyCompliancePackTemplatesResultCompliancePackTemplatesConfigRules {
	s.ManagedRuleName = &v
	return s
}

func (s *ListCompliancePackTemplatesResponseBodyCompliancePackTemplatesResultCompliancePackTemplatesConfigRules) SetResourceTypesScope(v string) *ListCompliancePackTemplatesResponseBodyCompliancePackTemplatesResultCompliancePackTemplatesConfigRules {
	s.ResourceTypesScope = &v
	return s
}

func (s *ListCompliancePackTemplatesResponseBodyCompliancePackTemplatesResultCompliancePackTemplatesConfigRules) SetRiskLevel(v int32) *ListCompliancePackTemplatesResponseBodyCompliancePackTemplatesResultCompliancePackTemplatesConfigRules {
	s.RiskLevel = &v
	return s
}

func (s *ListCompliancePackTemplatesResponseBodyCompliancePackTemplatesResultCompliancePackTemplatesConfigRules) Validate() error {
	if s.ConfigRuleParameters != nil {
		for _, item := range s.ConfigRuleParameters {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListCompliancePackTemplatesResponseBodyCompliancePackTemplatesResultCompliancePackTemplatesConfigRulesConfigRuleParameters struct {
	// The name of the input parameter of the managed rule.
	//
	// example:
	//
	// days
	ParameterName *string `json:"ParameterName,omitempty" xml:"ParameterName,omitempty"`
	// The value of the input parameter of the managed rule.
	//
	// example:
	//
	// 90
	ParameterValue *string `json:"ParameterValue,omitempty" xml:"ParameterValue,omitempty"`
	// Indicates whether the parameter is required in the managed rule. Valid values:
	//
	// 	- true: required
	//
	// 	- false: optional
	//
	// example:
	//
	// true
	Required *bool `json:"Required,omitempty" xml:"Required,omitempty"`
}

func (s ListCompliancePackTemplatesResponseBodyCompliancePackTemplatesResultCompliancePackTemplatesConfigRulesConfigRuleParameters) String() string {
	return dara.Prettify(s)
}

func (s ListCompliancePackTemplatesResponseBodyCompliancePackTemplatesResultCompliancePackTemplatesConfigRulesConfigRuleParameters) GoString() string {
	return s.String()
}

func (s *ListCompliancePackTemplatesResponseBodyCompliancePackTemplatesResultCompliancePackTemplatesConfigRulesConfigRuleParameters) GetParameterName() *string {
	return s.ParameterName
}

func (s *ListCompliancePackTemplatesResponseBodyCompliancePackTemplatesResultCompliancePackTemplatesConfigRulesConfigRuleParameters) GetParameterValue() *string {
	return s.ParameterValue
}

func (s *ListCompliancePackTemplatesResponseBodyCompliancePackTemplatesResultCompliancePackTemplatesConfigRulesConfigRuleParameters) GetRequired() *bool {
	return s.Required
}

func (s *ListCompliancePackTemplatesResponseBodyCompliancePackTemplatesResultCompliancePackTemplatesConfigRulesConfigRuleParameters) SetParameterName(v string) *ListCompliancePackTemplatesResponseBodyCompliancePackTemplatesResultCompliancePackTemplatesConfigRulesConfigRuleParameters {
	s.ParameterName = &v
	return s
}

func (s *ListCompliancePackTemplatesResponseBodyCompliancePackTemplatesResultCompliancePackTemplatesConfigRulesConfigRuleParameters) SetParameterValue(v string) *ListCompliancePackTemplatesResponseBodyCompliancePackTemplatesResultCompliancePackTemplatesConfigRulesConfigRuleParameters {
	s.ParameterValue = &v
	return s
}

func (s *ListCompliancePackTemplatesResponseBodyCompliancePackTemplatesResultCompliancePackTemplatesConfigRulesConfigRuleParameters) SetRequired(v bool) *ListCompliancePackTemplatesResponseBodyCompliancePackTemplatesResultCompliancePackTemplatesConfigRulesConfigRuleParameters {
	s.Required = &v
	return s
}

func (s *ListCompliancePackTemplatesResponseBodyCompliancePackTemplatesResultCompliancePackTemplatesConfigRulesConfigRuleParameters) Validate() error {
	return dara.Validate(s)
}
