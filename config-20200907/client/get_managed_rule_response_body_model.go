// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetManagedRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetManagedRule(v *GetManagedRuleResponseBodyManagedRule) *GetManagedRuleResponseBody
	GetManagedRule() *GetManagedRuleResponseBodyManagedRule
	SetRequestId(v string) *GetManagedRuleResponseBody
	GetRequestId() *string
}

type GetManagedRuleResponseBody struct {
	// The details of the rule template.
	ManagedRule *GetManagedRuleResponseBodyManagedRule `json:"ManagedRule,omitempty" xml:"ManagedRule,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 7E6DDC09-87C1-5310-A924-3491EAAE6F90
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetManagedRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetManagedRuleResponseBody) GoString() string {
	return s.String()
}

func (s *GetManagedRuleResponseBody) GetManagedRule() *GetManagedRuleResponseBodyManagedRule {
	return s.ManagedRule
}

func (s *GetManagedRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetManagedRuleResponseBody) SetManagedRule(v *GetManagedRuleResponseBodyManagedRule) *GetManagedRuleResponseBody {
	s.ManagedRule = v
	return s
}

func (s *GetManagedRuleResponseBody) SetRequestId(v string) *GetManagedRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetManagedRuleResponseBody) Validate() error {
	if s.ManagedRule != nil {
		if err := s.ManagedRule.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetManagedRuleResponseBodyManagedRule struct {
	// The information about the required input parameters for the rule template.
	//
	// example:
	//
	// {}
	CompulsoryInputParameterDetails map[string]interface{} `json:"CompulsoryInputParameterDetails,omitempty" xml:"CompulsoryInputParameterDetails,omitempty"`
	// The name of the rule template.
	//
	// example:
	//
	// CDN域名开启HTTPS加密
	ConfigRuleName *string `json:"ConfigRuleName,omitempty" xml:"ConfigRuleName,omitempty"`
	// The description of the rule template.
	//
	// example:
	//
	// CDN域名开启HTTPS协议加密，视为“合规”。
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The URL of the document that provides guidance on how to fix the issue.
	//
	// example:
	//
	// https://example.aliyundoc.com
	HelpUrls *string `json:"HelpUrls,omitempty" xml:"HelpUrls,omitempty"`
	// The identifier of the rule template.
	//
	// example:
	//
	// cdn-domain-https-enabled
	Identifier *string `json:"Identifier,omitempty" xml:"Identifier,omitempty"`
	// The labels of the rule template.
	Labels []*string `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	// The information about the optional input parameters for the rule template.
	//
	// example:
	//
	// {}
	OptionalInputParameterDetails map[string]interface{} `json:"OptionalInputParameterDetails,omitempty" xml:"OptionalInputParameterDetails,omitempty"`
	// The risk level of the rule template. Valid values:
	//
	// - 1: high risk.
	//
	// - 2: medium risk.
	//
	// - 3: low risk.
	//
	// example:
	//
	// 1
	RiskLevel *int32 `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	// The effective scope of the rule template.
	Scope *GetManagedRuleResponseBodyManagedRuleScope `json:"Scope,omitempty" xml:"Scope,omitempty" type:"Struct"`
	// The trigger methods for the rule.
	SourceDetails []*GetManagedRuleResponseBodyManagedRuleSourceDetails `json:"SourceDetails,omitempty" xml:"SourceDetails,omitempty" type:"Repeated"`
}

func (s GetManagedRuleResponseBodyManagedRule) String() string {
	return dara.Prettify(s)
}

func (s GetManagedRuleResponseBodyManagedRule) GoString() string {
	return s.String()
}

func (s *GetManagedRuleResponseBodyManagedRule) GetCompulsoryInputParameterDetails() map[string]interface{} {
	return s.CompulsoryInputParameterDetails
}

func (s *GetManagedRuleResponseBodyManagedRule) GetConfigRuleName() *string {
	return s.ConfigRuleName
}

func (s *GetManagedRuleResponseBodyManagedRule) GetDescription() *string {
	return s.Description
}

func (s *GetManagedRuleResponseBodyManagedRule) GetHelpUrls() *string {
	return s.HelpUrls
}

func (s *GetManagedRuleResponseBodyManagedRule) GetIdentifier() *string {
	return s.Identifier
}

func (s *GetManagedRuleResponseBodyManagedRule) GetLabels() []*string {
	return s.Labels
}

func (s *GetManagedRuleResponseBodyManagedRule) GetOptionalInputParameterDetails() map[string]interface{} {
	return s.OptionalInputParameterDetails
}

func (s *GetManagedRuleResponseBodyManagedRule) GetRiskLevel() *int32 {
	return s.RiskLevel
}

func (s *GetManagedRuleResponseBodyManagedRule) GetScope() *GetManagedRuleResponseBodyManagedRuleScope {
	return s.Scope
}

func (s *GetManagedRuleResponseBodyManagedRule) GetSourceDetails() []*GetManagedRuleResponseBodyManagedRuleSourceDetails {
	return s.SourceDetails
}

func (s *GetManagedRuleResponseBodyManagedRule) SetCompulsoryInputParameterDetails(v map[string]interface{}) *GetManagedRuleResponseBodyManagedRule {
	s.CompulsoryInputParameterDetails = v
	return s
}

func (s *GetManagedRuleResponseBodyManagedRule) SetConfigRuleName(v string) *GetManagedRuleResponseBodyManagedRule {
	s.ConfigRuleName = &v
	return s
}

func (s *GetManagedRuleResponseBodyManagedRule) SetDescription(v string) *GetManagedRuleResponseBodyManagedRule {
	s.Description = &v
	return s
}

func (s *GetManagedRuleResponseBodyManagedRule) SetHelpUrls(v string) *GetManagedRuleResponseBodyManagedRule {
	s.HelpUrls = &v
	return s
}

func (s *GetManagedRuleResponseBodyManagedRule) SetIdentifier(v string) *GetManagedRuleResponseBodyManagedRule {
	s.Identifier = &v
	return s
}

func (s *GetManagedRuleResponseBodyManagedRule) SetLabels(v []*string) *GetManagedRuleResponseBodyManagedRule {
	s.Labels = v
	return s
}

func (s *GetManagedRuleResponseBodyManagedRule) SetOptionalInputParameterDetails(v map[string]interface{}) *GetManagedRuleResponseBodyManagedRule {
	s.OptionalInputParameterDetails = v
	return s
}

func (s *GetManagedRuleResponseBodyManagedRule) SetRiskLevel(v int32) *GetManagedRuleResponseBodyManagedRule {
	s.RiskLevel = &v
	return s
}

func (s *GetManagedRuleResponseBodyManagedRule) SetScope(v *GetManagedRuleResponseBodyManagedRuleScope) *GetManagedRuleResponseBodyManagedRule {
	s.Scope = v
	return s
}

func (s *GetManagedRuleResponseBodyManagedRule) SetSourceDetails(v []*GetManagedRuleResponseBodyManagedRuleSourceDetails) *GetManagedRuleResponseBodyManagedRule {
	s.SourceDetails = v
	return s
}

func (s *GetManagedRuleResponseBodyManagedRule) Validate() error {
	if s.Scope != nil {
		if err := s.Scope.Validate(); err != nil {
			return err
		}
	}
	if s.SourceDetails != nil {
		for _, item := range s.SourceDetails {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetManagedRuleResponseBodyManagedRuleScope struct {
	// The resource types for which the rule template is effective.
	ComplianceResourceTypes []*string `json:"ComplianceResourceTypes,omitempty" xml:"ComplianceResourceTypes,omitempty" type:"Repeated"`
}

func (s GetManagedRuleResponseBodyManagedRuleScope) String() string {
	return dara.Prettify(s)
}

func (s GetManagedRuleResponseBodyManagedRuleScope) GoString() string {
	return s.String()
}

func (s *GetManagedRuleResponseBodyManagedRuleScope) GetComplianceResourceTypes() []*string {
	return s.ComplianceResourceTypes
}

func (s *GetManagedRuleResponseBodyManagedRuleScope) SetComplianceResourceTypes(v []*string) *GetManagedRuleResponseBodyManagedRuleScope {
	s.ComplianceResourceTypes = v
	return s
}

func (s *GetManagedRuleResponseBodyManagedRuleScope) Validate() error {
	return dara.Validate(s)
}

type GetManagedRuleResponseBodyManagedRuleSourceDetails struct {
	// The execution period of the rule. Valid values:
	//
	// - One_Hour: 1 hour.
	//
	// - Three_Hours: 3 hours.
	//
	// - Six_Hours: 6 hours.
	//
	// - Twelve_Hours: 12 hours.
	//
	// - TwentyFour_Hours: 24 hours.
	//
	// example:
	//
	// TwentyFour_Hours
	MaximumExecutionFrequency *string `json:"MaximumExecutionFrequency,omitempty" xml:"MaximumExecutionFrequency,omitempty"`
	// The trigger type of the rule. Valid values:
	//
	// - ConfigurationItemChangeNotification: The rule is triggered by a configuration change.
	//
	// - ScheduledNotification: The rule is triggered periodically.
	//
	// example:
	//
	// ConfigurationItemChangeNotification
	MessageType *string `json:"MessageType,omitempty" xml:"MessageType,omitempty"`
}

func (s GetManagedRuleResponseBodyManagedRuleSourceDetails) String() string {
	return dara.Prettify(s)
}

func (s GetManagedRuleResponseBodyManagedRuleSourceDetails) GoString() string {
	return s.String()
}

func (s *GetManagedRuleResponseBodyManagedRuleSourceDetails) GetMaximumExecutionFrequency() *string {
	return s.MaximumExecutionFrequency
}

func (s *GetManagedRuleResponseBodyManagedRuleSourceDetails) GetMessageType() *string {
	return s.MessageType
}

func (s *GetManagedRuleResponseBodyManagedRuleSourceDetails) SetMaximumExecutionFrequency(v string) *GetManagedRuleResponseBodyManagedRuleSourceDetails {
	s.MaximumExecutionFrequency = &v
	return s
}

func (s *GetManagedRuleResponseBodyManagedRuleSourceDetails) SetMessageType(v string) *GetManagedRuleResponseBodyManagedRuleSourceDetails {
	s.MessageType = &v
	return s
}

func (s *GetManagedRuleResponseBodyManagedRuleSourceDetails) Validate() error {
	return dara.Validate(s)
}
