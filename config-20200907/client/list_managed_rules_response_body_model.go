// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListManagedRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetManagedRules(v *ListManagedRulesResponseBodyManagedRules) *ListManagedRulesResponseBody
	GetManagedRules() *ListManagedRulesResponseBodyManagedRules
	SetRequestId(v string) *ListManagedRulesResponseBody
	GetRequestId() *string
}

type ListManagedRulesResponseBody struct {
	// The managed rules.
	ManagedRules *ListManagedRulesResponseBodyManagedRules `json:"ManagedRules,omitempty" xml:"ManagedRules,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// B3E605AB-63D5-1EE0-BFA6-0BAC247B0461
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListManagedRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListManagedRulesResponseBody) GoString() string {
	return s.String()
}

func (s *ListManagedRulesResponseBody) GetManagedRules() *ListManagedRulesResponseBodyManagedRules {
	return s.ManagedRules
}

func (s *ListManagedRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListManagedRulesResponseBody) SetManagedRules(v *ListManagedRulesResponseBodyManagedRules) *ListManagedRulesResponseBody {
	s.ManagedRules = v
	return s
}

func (s *ListManagedRulesResponseBody) SetRequestId(v string) *ListManagedRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListManagedRulesResponseBody) Validate() error {
	if s.ManagedRules != nil {
		if err := s.ManagedRules.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListManagedRulesResponseBodyManagedRules struct {
	// The details of the managed rule.
	ManagedRuleList []*ListManagedRulesResponseBodyManagedRulesManagedRuleList `json:"ManagedRuleList,omitempty" xml:"ManagedRuleList,omitempty" type:"Repeated"`
	// The page number.
	//
	// Page start from page 1.
	//
	// example:
	//
	// 10
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page. Valid values: 1 to 500.
	//
	// example:
	//
	// 1
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListManagedRulesResponseBodyManagedRules) String() string {
	return dara.Prettify(s)
}

func (s ListManagedRulesResponseBodyManagedRules) GoString() string {
	return s.String()
}

func (s *ListManagedRulesResponseBodyManagedRules) GetManagedRuleList() []*ListManagedRulesResponseBodyManagedRulesManagedRuleList {
	return s.ManagedRuleList
}

func (s *ListManagedRulesResponseBodyManagedRules) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListManagedRulesResponseBodyManagedRules) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListManagedRulesResponseBodyManagedRules) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListManagedRulesResponseBodyManagedRules) SetManagedRuleList(v []*ListManagedRulesResponseBodyManagedRulesManagedRuleList) *ListManagedRulesResponseBodyManagedRules {
	s.ManagedRuleList = v
	return s
}

func (s *ListManagedRulesResponseBodyManagedRules) SetPageNumber(v int32) *ListManagedRulesResponseBodyManagedRules {
	s.PageNumber = &v
	return s
}

func (s *ListManagedRulesResponseBodyManagedRules) SetPageSize(v int32) *ListManagedRulesResponseBodyManagedRules {
	s.PageSize = &v
	return s
}

func (s *ListManagedRulesResponseBodyManagedRules) SetTotalCount(v int64) *ListManagedRulesResponseBodyManagedRules {
	s.TotalCount = &v
	return s
}

func (s *ListManagedRulesResponseBodyManagedRules) Validate() error {
	if s.ManagedRuleList != nil {
		for _, item := range s.ManagedRuleList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListManagedRulesResponseBodyManagedRulesManagedRuleList struct {
	// The name of the managed rule.
	//
	// example:
	//
	// test-rule-name
	ConfigRuleName *string `json:"ConfigRuleName,omitempty" xml:"ConfigRuleName,omitempty"`
	// The description of the managed rule.
	//
	// example:
	//
	// The description of the test rule.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The URL of the topic that describes how the managed rule remediates the incompliant configurations.
	//
	// example:
	//
	// https://example.aliyundoc.com
	HelpUrls *string `json:"HelpUrls,omitempty" xml:"HelpUrls,omitempty"`
	// The unique identifier of the managed rule.
	//
	// example:
	//
	// cdn-domain-https-enabled
	Identifier *string `json:"Identifier,omitempty" xml:"Identifier,omitempty"`
	// The classification description of the managed rule.
	Labels []*string `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	// The ID of the remediation template.
	//
	// example:
	//
	// ACS-CDN-SetDomainServerCertificate
	RemediationTemplateIdentifier *string `json:"RemediationTemplateIdentifier,omitempty" xml:"RemediationTemplateIdentifier,omitempty"`
	// The name of the remediation template.
	//
	// example:
	//
	// Configure encryption rules for OSS buckets
	RemediationTemplateName *string `json:"RemediationTemplateName,omitempty" xml:"RemediationTemplateName,omitempty"`
	// The risk level of the resources that do not comply with the rule. Valid values:
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
	// The effective scope of the managed rule.
	Scope *ListManagedRulesResponseBodyManagedRulesManagedRuleListScope `json:"Scope,omitempty" xml:"Scope,omitempty" type:"Struct"`
	// Indicates whether precheck is supported. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	SupportPreviewManagedRule *bool `json:"SupportPreviewManagedRule,omitempty" xml:"SupportPreviewManagedRule,omitempty"`
}

func (s ListManagedRulesResponseBodyManagedRulesManagedRuleList) String() string {
	return dara.Prettify(s)
}

func (s ListManagedRulesResponseBodyManagedRulesManagedRuleList) GoString() string {
	return s.String()
}

func (s *ListManagedRulesResponseBodyManagedRulesManagedRuleList) GetConfigRuleName() *string {
	return s.ConfigRuleName
}

func (s *ListManagedRulesResponseBodyManagedRulesManagedRuleList) GetDescription() *string {
	return s.Description
}

func (s *ListManagedRulesResponseBodyManagedRulesManagedRuleList) GetHelpUrls() *string {
	return s.HelpUrls
}

func (s *ListManagedRulesResponseBodyManagedRulesManagedRuleList) GetIdentifier() *string {
	return s.Identifier
}

func (s *ListManagedRulesResponseBodyManagedRulesManagedRuleList) GetLabels() []*string {
	return s.Labels
}

func (s *ListManagedRulesResponseBodyManagedRulesManagedRuleList) GetRemediationTemplateIdentifier() *string {
	return s.RemediationTemplateIdentifier
}

func (s *ListManagedRulesResponseBodyManagedRulesManagedRuleList) GetRemediationTemplateName() *string {
	return s.RemediationTemplateName
}

func (s *ListManagedRulesResponseBodyManagedRulesManagedRuleList) GetRiskLevel() *int32 {
	return s.RiskLevel
}

func (s *ListManagedRulesResponseBodyManagedRulesManagedRuleList) GetScope() *ListManagedRulesResponseBodyManagedRulesManagedRuleListScope {
	return s.Scope
}

func (s *ListManagedRulesResponseBodyManagedRulesManagedRuleList) GetSupportPreviewManagedRule() *bool {
	return s.SupportPreviewManagedRule
}

func (s *ListManagedRulesResponseBodyManagedRulesManagedRuleList) SetConfigRuleName(v string) *ListManagedRulesResponseBodyManagedRulesManagedRuleList {
	s.ConfigRuleName = &v
	return s
}

func (s *ListManagedRulesResponseBodyManagedRulesManagedRuleList) SetDescription(v string) *ListManagedRulesResponseBodyManagedRulesManagedRuleList {
	s.Description = &v
	return s
}

func (s *ListManagedRulesResponseBodyManagedRulesManagedRuleList) SetHelpUrls(v string) *ListManagedRulesResponseBodyManagedRulesManagedRuleList {
	s.HelpUrls = &v
	return s
}

func (s *ListManagedRulesResponseBodyManagedRulesManagedRuleList) SetIdentifier(v string) *ListManagedRulesResponseBodyManagedRulesManagedRuleList {
	s.Identifier = &v
	return s
}

func (s *ListManagedRulesResponseBodyManagedRulesManagedRuleList) SetLabels(v []*string) *ListManagedRulesResponseBodyManagedRulesManagedRuleList {
	s.Labels = v
	return s
}

func (s *ListManagedRulesResponseBodyManagedRulesManagedRuleList) SetRemediationTemplateIdentifier(v string) *ListManagedRulesResponseBodyManagedRulesManagedRuleList {
	s.RemediationTemplateIdentifier = &v
	return s
}

func (s *ListManagedRulesResponseBodyManagedRulesManagedRuleList) SetRemediationTemplateName(v string) *ListManagedRulesResponseBodyManagedRulesManagedRuleList {
	s.RemediationTemplateName = &v
	return s
}

func (s *ListManagedRulesResponseBodyManagedRulesManagedRuleList) SetRiskLevel(v int32) *ListManagedRulesResponseBodyManagedRulesManagedRuleList {
	s.RiskLevel = &v
	return s
}

func (s *ListManagedRulesResponseBodyManagedRulesManagedRuleList) SetScope(v *ListManagedRulesResponseBodyManagedRulesManagedRuleListScope) *ListManagedRulesResponseBodyManagedRulesManagedRuleList {
	s.Scope = v
	return s
}

func (s *ListManagedRulesResponseBodyManagedRulesManagedRuleList) SetSupportPreviewManagedRule(v bool) *ListManagedRulesResponseBodyManagedRulesManagedRuleList {
	s.SupportPreviewManagedRule = &v
	return s
}

func (s *ListManagedRulesResponseBodyManagedRulesManagedRuleList) Validate() error {
	if s.Scope != nil {
		if err := s.Scope.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListManagedRulesResponseBodyManagedRulesManagedRuleListScope struct {
	// The types of resources to which the managed rule applies.
	ComplianceResourceTypes []*string `json:"ComplianceResourceTypes,omitempty" xml:"ComplianceResourceTypes,omitempty" type:"Repeated"`
}

func (s ListManagedRulesResponseBodyManagedRulesManagedRuleListScope) String() string {
	return dara.Prettify(s)
}

func (s ListManagedRulesResponseBodyManagedRulesManagedRuleListScope) GoString() string {
	return s.String()
}

func (s *ListManagedRulesResponseBodyManagedRulesManagedRuleListScope) GetComplianceResourceTypes() []*string {
	return s.ComplianceResourceTypes
}

func (s *ListManagedRulesResponseBodyManagedRulesManagedRuleListScope) SetComplianceResourceTypes(v []*string) *ListManagedRulesResponseBodyManagedRulesManagedRuleListScope {
	s.ComplianceResourceTypes = v
	return s
}

func (s *ListManagedRulesResponseBodyManagedRulesManagedRuleListScope) Validate() error {
	return dara.Validate(s)
}
