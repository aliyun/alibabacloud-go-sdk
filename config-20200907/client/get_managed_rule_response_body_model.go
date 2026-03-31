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
	// The details of the managed rule.
	ManagedRule *GetManagedRuleResponseBodyManagedRule `json:"ManagedRule,omitempty" xml:"ManagedRule,omitempty" type:"Struct"`
	// The ID of the request.
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
	// The details of the required input parameters for the managed rule.
	//
	// example:
	//
	// {}
	CompulsoryInputParameterDetails map[string]interface{} `json:"CompulsoryInputParameterDetails,omitempty" xml:"CompulsoryInputParameterDetails,omitempty"`
	// The name of the managed rule.
	//
	// example:
	//
	// cdn-domain-https-enabled
	ConfigRuleName *string `json:"ConfigRuleName,omitempty" xml:"ConfigRuleName,omitempty"`
	// The description of the managed rule.
	//
	// example:
	//
	// If HTTPS encryption is enabled for the CDN domain name, the configuration is considered compliant.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The URL of the topic that provides guidance on remediation for the managed rule.
	//
	// example:
	//
	// https://example.aliyundoc.com
	HelpUrls *string `json:"HelpUrls,omitempty" xml:"HelpUrls,omitempty"`
	// The identifier of the managed rule.
	//
	// example:
	//
	// cdn-domain-https-enabled
	Identifier *string `json:"Identifier,omitempty" xml:"Identifier,omitempty"`
	// The tags of the managed rule.
	Labels []*string `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	// The details of the optional input parameters for the managed rule.
	//
	// example:
	//
	// {}
	OptionalInputParameterDetails map[string]interface{} `json:"OptionalInputParameterDetails,omitempty" xml:"OptionalInputParameterDetails,omitempty"`
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
	// The effective scope of the managed rule.
	Scope *GetManagedRuleResponseBodyManagedRuleScope `json:"Scope,omitempty" xml:"Scope,omitempty" type:"Struct"`
	// The information about the trigger type of the managed rule.
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
	// The types of resources to which the managed rule applies.
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
	// The interval at which the rule is triggered. Valid values: Valid values:
	//
	// 	- One_Hour
	//
	// 	- Three_Hours
	//
	// 	- Six_Hours
	//
	// 	- Twelve_Hours
	//
	// 	- TwentyFour_Hours
	//
	// example:
	//
	// TwentyFour_Hours
	MaximumExecutionFrequency *string `json:"MaximumExecutionFrequency,omitempty" xml:"MaximumExecutionFrequency,omitempty"`
	// The trigger type of the rule. Valid values:
	//
	// 	- ConfigurationItemChangeNotification: The rule is triggered by configuration changes.
	//
	// 	- ScheduledNotification: The rule is periodically triggered.
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
