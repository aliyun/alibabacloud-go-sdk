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
	// The details of the compliance pack templates.
	CompliancePackTemplatesResult *ListCompliancePackTemplatesResponseBodyCompliancePackTemplatesResult `json:"CompliancePackTemplatesResult,omitempty" xml:"CompliancePackTemplatesResult,omitempty" type:"Struct"`
	// The ID of the request.
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
	// A list of compliance pack templates.
	CompliancePackTemplates []*ListCompliancePackTemplatesResponseBodyCompliancePackTemplatesResultCompliancePackTemplates `json:"CompliancePackTemplates,omitempty" xml:"CompliancePackTemplates,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of compliance pack templates.
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
	// The ID of the compliance pack template.
	//
	// example:
	//
	// ct-5f26ff4e06a300c4****
	CompliancePackTemplateId *string `json:"CompliancePackTemplateId,omitempty" xml:"CompliancePackTemplateId,omitempty"`
	// The name of the compliance pack template.
	//
	// example:
	//
	// ClassifiedProtectionPreCheck
	CompliancePackTemplateName *string `json:"CompliancePackTemplateName,omitempty" xml:"CompliancePackTemplateName,omitempty"`
	// A list of default rules in the compliance pack.
	ConfigRules []*ListCompliancePackTemplatesResponseBodyCompliancePackTemplatesResultCompliancePackTemplatesConfigRules `json:"ConfigRules,omitempty" xml:"ConfigRules,omitempty" type:"Repeated"`
	// The description of the compliance pack.
	//
	// example:
	//
	// Checks the compliance of Alibaba Cloud resources based on the specific requirements of MLPS 2.0 Level 3.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The classification labels of the compliance pack.
	//
	// example:
	//
	// Regulation
	Labels *string `json:"Labels,omitempty" xml:"Labels,omitempty"`
	// The UNIX timestamp when the compliance pack was last updated.
	//
	// example:
	//
	// 1747983081
	LastUpdate *int32 `json:"LastUpdate,omitempty" xml:"LastUpdate,omitempty"`
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
	// The parameters of the managed rule.
	ConfigRuleParameters []*ListCompliancePackTemplatesResponseBodyCompliancePackTemplatesResultCompliancePackTemplatesConfigRulesConfigRuleParameters `json:"ConfigRuleParameters,omitempty" xml:"ConfigRuleParameters,omitempty" type:"Repeated"`
	// The description of the control. This parameter is returned only for compliance packs that are created based on a regulation.
	//
	// example:
	//
	// e) Possible known vulnerabilities should be found and remedied in a timely manner after adequate testing and evaluation.\\nf) It shall be able to detect intrusion into important nodes and provide an alarm in case of serious intrusion events.
	ControlDescription *string `json:"ControlDescription,omitempty" xml:"ControlDescription,omitempty"`
	// The ID of the control.
	//
	// > This parameter appears only for regulatory compliance packages.
	//
	// example:
	//
	// 8.1.4.4
	ControlId *string `json:"ControlId,omitempty" xml:"ControlId,omitempty"`
	// Indicates whether the rule can be quickly enabled. Valid values:
	//
	// - true: The rule can be quickly enabled.
	//
	// - false: The rule cannot be quickly enabled.
	//
	// example:
	//
	// false
	DefaultEnable *bool `json:"DefaultEnable,omitempty" xml:"DefaultEnable,omitempty"`
	// The description of the rule.
	//
	// example:
	//
	// If no vulnerabilities that are of the specified type and severity level and to be fixed exist on the running ECS instances that are protected by Security Center, the configuration is considered compliant. This rule does not apply to ECS instances that are not in the running state.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The identifier of the managed rule.
	//
	// example:
	//
	// ecs-instance-updated-security-vul
	ManagedRuleIdentifier *string `json:"ManagedRuleIdentifier,omitempty" xml:"ManagedRuleIdentifier,omitempty"`
	// The name of the managed rule.
	//
	// example:
	//
	// ecs-instance-updated-security-vul
	ManagedRuleName *string `json:"ManagedRuleName,omitempty" xml:"ManagedRuleName,omitempty"`
	// The resource types that are evaluated by the rule.
	//
	// example:
	//
	// ACS::ECS::Instance
	ResourceTypesScope *string `json:"ResourceTypesScope,omitempty" xml:"ResourceTypesScope,omitempty"`
	// The risk level of the managed rule. Valid values:
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
	// The name of the parameter for the managed rule.
	//
	// example:
	//
	// necessity
	ParameterName *string `json:"ParameterName,omitempty" xml:"ParameterName,omitempty"`
	// The value of the parameter for the managed rule.
	//
	// example:
	//
	// asap
	ParameterValue *string `json:"ParameterValue,omitempty" xml:"ParameterValue,omitempty"`
	// Indicates whether the parameter is required for the managed rule. Valid values:
	//
	// - true: The parameter is required.
	//
	// - false: The parameter is not required.
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
