// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEgressProxyRuleGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateEgressProxyRuleGroupRequest
	GetClientToken() *string
	SetDescription(v string) *CreateEgressProxyRuleGroupRequest
	GetDescription() *string
	SetDryRun(v bool) *CreateEgressProxyRuleGroupRequest
	GetDryRun() *bool
	SetEgressProxyIpPoolId(v string) *CreateEgressProxyRuleGroupRequest
	GetEgressProxyIpPoolId() *string
	SetEgressProxyRuleGroupName(v string) *CreateEgressProxyRuleGroupRequest
	GetEgressProxyRuleGroupName() *string
	SetEgressProxyRules(v *CreateEgressProxyRuleGroupRequestEgressProxyRules) *CreateEgressProxyRuleGroupRequest
	GetEgressProxyRules() *CreateEgressProxyRuleGroupRequestEgressProxyRules
	SetResourceGroupId(v string) *CreateEgressProxyRuleGroupRequest
	GetResourceGroupId() *string
	SetTag(v []*CreateEgressProxyRuleGroupRequestTag) *CreateEgressProxyRuleGroupRequest
	GetTag() []*CreateEgressProxyRuleGroupRequestTag
}

type CreateEgressProxyRuleGroupRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the value, but you must ensure that it is unique among different requests. The client token can contain only ASCII characters.
	//
	// > Note If you do not specify this parameter, the system automatically uses the RequestId of the API request as the client token. The RequestId of each API request is different.
	//
	// example:
	//
	// 5A2CFF0E-5718-45B5-9D4D-70B******
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the egress rule group. The description must be 2 to 256 characters in length and can contain Chinese characters, letters, digits, hyphens (-), forward slashes (/), periods (.), and underscores (_).
	//
	// example:
	//
	// testEpg
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether to perform only a dry run. Valid values:
	//
	// - **true**: sends a dry run request and does not create the rule group. The system checks whether required parameters are specified, the request format is valid, and service limits are met. If the request fails the dry run, an error message is returned. If the request passes the dry run, the DryRunOperation error code is returned.
	//
	// - **false*	- (default): sends a normal request. If the request passes the dry run, an HTTP 2xx status code is returned and the operation is performed.
	//
	// example:
	//
	// true
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ID of the egress IP address pool associated with the egress rule group.
	//
	// example:
	//
	// eppool-gw80d96a96ca0b00****
	EgressProxyIpPoolId *string `json:"EgressProxyIpPoolId,omitempty" xml:"EgressProxyIpPoolId,omitempty"`
	// The name of the egress rule group. The name must be 2 to 128 characters in length and must start with a letter or a Chinese character. It can contain digits, periods (.), underscores (_), and hyphens (-).
	//
	// This parameter is required.
	//
	// example:
	//
	// testEpg
	EgressProxyRuleGroupName *string `json:"EgressProxyRuleGroupName,omitempty" xml:"EgressProxyRuleGroupName,omitempty"`
	// The list of egress rules grouped by phase (PreDns, PreRequest, and PostResponse).
	EgressProxyRules *CreateEgressProxyRuleGroupRequestEgressProxyRules `json:"EgressProxyRules,omitempty" xml:"EgressProxyRules,omitempty" type:"Struct"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-gw80bb6a94f835db****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The list of tags.
	Tag []*CreateEgressProxyRuleGroupRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s CreateEgressProxyRuleGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateEgressProxyRuleGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateEgressProxyRuleGroupRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateEgressProxyRuleGroupRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateEgressProxyRuleGroupRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *CreateEgressProxyRuleGroupRequest) GetEgressProxyIpPoolId() *string {
	return s.EgressProxyIpPoolId
}

func (s *CreateEgressProxyRuleGroupRequest) GetEgressProxyRuleGroupName() *string {
	return s.EgressProxyRuleGroupName
}

func (s *CreateEgressProxyRuleGroupRequest) GetEgressProxyRules() *CreateEgressProxyRuleGroupRequestEgressProxyRules {
	return s.EgressProxyRules
}

func (s *CreateEgressProxyRuleGroupRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateEgressProxyRuleGroupRequest) GetTag() []*CreateEgressProxyRuleGroupRequestTag {
	return s.Tag
}

func (s *CreateEgressProxyRuleGroupRequest) SetClientToken(v string) *CreateEgressProxyRuleGroupRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateEgressProxyRuleGroupRequest) SetDescription(v string) *CreateEgressProxyRuleGroupRequest {
	s.Description = &v
	return s
}

func (s *CreateEgressProxyRuleGroupRequest) SetDryRun(v bool) *CreateEgressProxyRuleGroupRequest {
	s.DryRun = &v
	return s
}

func (s *CreateEgressProxyRuleGroupRequest) SetEgressProxyIpPoolId(v string) *CreateEgressProxyRuleGroupRequest {
	s.EgressProxyIpPoolId = &v
	return s
}

func (s *CreateEgressProxyRuleGroupRequest) SetEgressProxyRuleGroupName(v string) *CreateEgressProxyRuleGroupRequest {
	s.EgressProxyRuleGroupName = &v
	return s
}

func (s *CreateEgressProxyRuleGroupRequest) SetEgressProxyRules(v *CreateEgressProxyRuleGroupRequestEgressProxyRules) *CreateEgressProxyRuleGroupRequest {
	s.EgressProxyRules = v
	return s
}

func (s *CreateEgressProxyRuleGroupRequest) SetResourceGroupId(v string) *CreateEgressProxyRuleGroupRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateEgressProxyRuleGroupRequest) SetTag(v []*CreateEgressProxyRuleGroupRequestTag) *CreateEgressProxyRuleGroupRequest {
	s.Tag = v
	return s
}

func (s *CreateEgressProxyRuleGroupRequest) Validate() error {
	if s.EgressProxyRules != nil {
		if err := s.EgressProxyRules.Validate(); err != nil {
			return err
		}
	}
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateEgressProxyRuleGroupRequestEgressProxyRules struct {
	// The list of rules in the PostResponse phase (after the response is returned).
	PostResponse []*CreateEgressProxyRuleGroupRequestEgressProxyRulesPostResponse `json:"PostResponse,omitempty" xml:"PostResponse,omitempty" type:"Repeated"`
	// The list of rules for the Pre-DNS stage (before DNS resolution).
	PreDns []*CreateEgressProxyRuleGroupRequestEgressProxyRulesPreDns `json:"PreDns,omitempty" xml:"PreDns,omitempty" type:"Repeated"`
	// The list of rules in the PreRequest stage, which is executed before the request is forwarded.
	PreRequest []*CreateEgressProxyRuleGroupRequestEgressProxyRulesPreRequest `json:"PreRequest,omitempty" xml:"PreRequest,omitempty" type:"Repeated"`
}

func (s CreateEgressProxyRuleGroupRequestEgressProxyRules) String() string {
	return dara.Prettify(s)
}

func (s CreateEgressProxyRuleGroupRequestEgressProxyRules) GoString() string {
	return s.String()
}

func (s *CreateEgressProxyRuleGroupRequestEgressProxyRules) GetPostResponse() []*CreateEgressProxyRuleGroupRequestEgressProxyRulesPostResponse {
	return s.PostResponse
}

func (s *CreateEgressProxyRuleGroupRequestEgressProxyRules) GetPreDns() []*CreateEgressProxyRuleGroupRequestEgressProxyRulesPreDns {
	return s.PreDns
}

func (s *CreateEgressProxyRuleGroupRequestEgressProxyRules) GetPreRequest() []*CreateEgressProxyRuleGroupRequestEgressProxyRulesPreRequest {
	return s.PreRequest
}

func (s *CreateEgressProxyRuleGroupRequestEgressProxyRules) SetPostResponse(v []*CreateEgressProxyRuleGroupRequestEgressProxyRulesPostResponse) *CreateEgressProxyRuleGroupRequestEgressProxyRules {
	s.PostResponse = v
	return s
}

func (s *CreateEgressProxyRuleGroupRequestEgressProxyRules) SetPreDns(v []*CreateEgressProxyRuleGroupRequestEgressProxyRulesPreDns) *CreateEgressProxyRuleGroupRequestEgressProxyRules {
	s.PreDns = v
	return s
}

func (s *CreateEgressProxyRuleGroupRequestEgressProxyRules) SetPreRequest(v []*CreateEgressProxyRuleGroupRequestEgressProxyRulesPreRequest) *CreateEgressProxyRuleGroupRequestEgressProxyRules {
	s.PreRequest = v
	return s
}

func (s *CreateEgressProxyRuleGroupRequestEgressProxyRules) Validate() error {
	if s.PostResponse != nil {
		for _, item := range s.PostResponse {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.PreDns != nil {
		for _, item := range s.PreDns {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.PreRequest != nil {
		for _, item := range s.PreRequest {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateEgressProxyRuleGroupRequestEgressProxyRulesPostResponse struct {
	// The action to perform on the matched traffic after the rule is hit. Valid values:
	//
	// - **Allow**: allows the traffic that matches the rule.
	//
	// - **Deny**: denies the traffic that matches the rule.
	//
	// This parameter is required.
	//
	// example:
	//
	// Allow
	Action *string `json:"Action,omitempty" xml:"Action,omitempty"`
	// The list of matching conditions for the rule. Each condition consists of a condition key (ConditionKey), an operator (ConditionOperator), and condition values (ConditionValues). The combination of a condition key and an operator must be unique within the same rule.
	//
	// This parameter is required.
	Conditions []*CreateEgressProxyRuleGroupRequestEgressProxyRulesPostResponseConditions `json:"Conditions,omitempty" xml:"Conditions,omitempty" type:"Repeated"`
	// The description of the egress rule. The description must be 2 to 256 characters in length and can contain Chinese characters, letters, digits, hyphens (-), forward slashes (/), periods (.), and underscores (_).
	//
	// example:
	//
	// testEpg
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the egress rule. The name must be 2 to 128 characters in length and must start with a letter or a Chinese character. It can contain digits, periods (.), underscores (_), and hyphens (-).
	//
	// This parameter is required.
	//
	// example:
	//
	// testEpg
	EgressProxyRuleName *string `json:"EgressProxyRuleName,omitempty" xml:"EgressProxyRuleName,omitempty"`
	// The priority of the egress rule within its phase. Valid values: 1 to 1000. A smaller value indicates a higher priority. The priority must be unique within the same phase of a rule group.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
}

func (s CreateEgressProxyRuleGroupRequestEgressProxyRulesPostResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateEgressProxyRuleGroupRequestEgressProxyRulesPostResponse) GoString() string {
	return s.String()
}

func (s *CreateEgressProxyRuleGroupRequestEgressProxyRulesPostResponse) GetAction() *string {
	return s.Action
}

func (s *CreateEgressProxyRuleGroupRequestEgressProxyRulesPostResponse) GetConditions() []*CreateEgressProxyRuleGroupRequestEgressProxyRulesPostResponseConditions {
	return s.Conditions
}

func (s *CreateEgressProxyRuleGroupRequestEgressProxyRulesPostResponse) GetDescription() *string {
	return s.Description
}

func (s *CreateEgressProxyRuleGroupRequestEgressProxyRulesPostResponse) GetEgressProxyRuleName() *string {
	return s.EgressProxyRuleName
}

func (s *CreateEgressProxyRuleGroupRequestEgressProxyRulesPostResponse) GetPriority() *int32 {
	return s.Priority
}

func (s *CreateEgressProxyRuleGroupRequestEgressProxyRulesPostResponse) SetAction(v string) *CreateEgressProxyRuleGroupRequestEgressProxyRulesPostResponse {
	s.Action = &v
	return s
}

func (s *CreateEgressProxyRuleGroupRequestEgressProxyRulesPostResponse) SetConditions(v []*CreateEgressProxyRuleGroupRequestEgressProxyRulesPostResponseConditions) *CreateEgressProxyRuleGroupRequestEgressProxyRulesPostResponse {
	s.Conditions = v
	return s
}

func (s *CreateEgressProxyRuleGroupRequestEgressProxyRulesPostResponse) SetDescription(v string) *CreateEgressProxyRuleGroupRequestEgressProxyRulesPostResponse {
	s.Description = &v
	return s
}

func (s *CreateEgressProxyRuleGroupRequestEgressProxyRulesPostResponse) SetEgressProxyRuleName(v string) *CreateEgressProxyRuleGroupRequestEgressProxyRulesPostResponse {
	s.EgressProxyRuleName = &v
	return s
}

func (s *CreateEgressProxyRuleGroupRequestEgressProxyRulesPostResponse) SetPriority(v int32) *CreateEgressProxyRuleGroupRequestEgressProxyRulesPostResponse {
	s.Priority = &v
	return s
}

func (s *CreateEgressProxyRuleGroupRequestEgressProxyRulesPostResponse) Validate() error {
	if s.Conditions != nil {
		for _, item := range s.Conditions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateEgressProxyRuleGroupRequestEgressProxyRulesPostResponseConditions struct {
	// The key of the matching condition. The value must match the phase of the rule and be used with an operator (ConditionOperator) of the corresponding category. Valid values:
	//
	// - **request:SourceIp**: the source IP address of the request. Used with IP address operators. Supported in the PreDns and PreRequest phases.
	//
	// - **request:DestinationDomain**: the destination domain name of the request. Used with string operators. Supported in the PreDns and PreRequest phases.
	//
	// - **request:Http:Uri**: the URI of the HTTP request. Used with string operators. Supported in the PreRequest phase.
	//
	// - **request:Http:Method**: the HTTP request method. Used with string operators. Supported in the PreRequest phase.
	//
	// - **request:Http:UserAgent**: the User-Agent of the HTTP request. Used with string operators. Supported in the PreRequest phase.
	//
	// - **request:Http:ContentType**: the Content-Type of the HTTP request. Used with string operators. Supported in the PreRequest phase.
	//
	// - **request:Http:Header**: the HTTP request header. Used with string operators. Supported in the PreRequest phase. When you use this key, append the header name to the key in the format of request:Http:Header/<HeaderName>.
	//
	// - **request:Protocol**: the request protocol. Used with string operators. Supported in the PreRequest phase.
	//
	// This parameter is required.
	//
	// example:
	//
	// request:SourceIp
	ConditionKey *string `json:"ConditionKey,omitempty" xml:"ConditionKey,omitempty"`
	// The operator of the matching condition. It must be used with a condition key (ConditionKey). Valid values:
	//
	// - **StringEquals**: exact string match.
	//
	// - **StringNotEquals**: string inequality.
	//
	// - **StringEqualsIgnoreCase**: case-insensitive exact string match.
	//
	// - **StringNotEqualsIgnoreCase**: case-insensitive string inequality.
	//
	// - **StringPrefixMatch**: string prefix match.
	//
	// - **StringPrefixMatchIgnoreCase**: case-insensitive string prefix match.
	//
	// - **StringPrefixNotMatch**: string prefix mismatch.
	//
	// - **StringPrefixNotMatchIgnoreCase**: case-insensitive string prefix mismatch.
	//
	// - **StringSuffixMatch**: string suffix match.
	//
	// - **StringSuffixMatchIgnoreCase**: case-insensitive string suffix match.
	//
	// - **StringSuffixNotMatch**: string suffix mismatch.
	//
	// - **StringSuffixNotMatchIgnoreCase**: case-insensitive string suffix mismatch.
	//
	// - **StringRegexMatch**: string regular expression match.
	//
	// - **StringRegexNotMatch**: string regular expression mismatch.
	//
	// - **StringWildcardMatch**: string wildcard match.
	//
	// - **StringWildcardMatchIgnoreCase**: case-insensitive string wildcard match.
	//
	// - **StringWildcardNotMatch**: string wildcard mismatch.
	//
	// - **StringWildcardNotMatchIgnoreCase**: case-insensitive string wildcard mismatch.
	//
	// - **IpAddress**: The source IP address belongs to the specified IP address or CIDR block.
	//
	// - **NotIpAddress**: The source IP address does not belong to the specified IP address or CIDR block.
	//
	// This parameter is required.
	//
	// example:
	//
	// IpAddress
	ConditionOperator *string `json:"ConditionOperator,omitempty" xml:"ConditionOperator,omitempty"`
	// The list of condition values.
	//
	// This parameter is required.
	ConditionValues []*string `json:"ConditionValues,omitempty" xml:"ConditionValues,omitempty" type:"Repeated"`
}

func (s CreateEgressProxyRuleGroupRequestEgressProxyRulesPostResponseConditions) String() string {
	return dara.Prettify(s)
}

func (s CreateEgressProxyRuleGroupRequestEgressProxyRulesPostResponseConditions) GoString() string {
	return s.String()
}

func (s *CreateEgressProxyRuleGroupRequestEgressProxyRulesPostResponseConditions) GetConditionKey() *string {
	return s.ConditionKey
}

func (s *CreateEgressProxyRuleGroupRequestEgressProxyRulesPostResponseConditions) GetConditionOperator() *string {
	return s.ConditionOperator
}

func (s *CreateEgressProxyRuleGroupRequestEgressProxyRulesPostResponseConditions) GetConditionValues() []*string {
	return s.ConditionValues
}

func (s *CreateEgressProxyRuleGroupRequestEgressProxyRulesPostResponseConditions) SetConditionKey(v string) *CreateEgressProxyRuleGroupRequestEgressProxyRulesPostResponseConditions {
	s.ConditionKey = &v
	return s
}

func (s *CreateEgressProxyRuleGroupRequestEgressProxyRulesPostResponseConditions) SetConditionOperator(v string) *CreateEgressProxyRuleGroupRequestEgressProxyRulesPostResponseConditions {
	s.ConditionOperator = &v
	return s
}

func (s *CreateEgressProxyRuleGroupRequestEgressProxyRulesPostResponseConditions) SetConditionValues(v []*string) *CreateEgressProxyRuleGroupRequestEgressProxyRulesPostResponseConditions {
	s.ConditionValues = v
	return s
}

func (s *CreateEgressProxyRuleGroupRequestEgressProxyRulesPostResponseConditions) Validate() error {
	return dara.Validate(s)
}

type CreateEgressProxyRuleGroupRequestEgressProxyRulesPreDns struct {
	// The action to perform on the matched traffic after the rule is hit. Valid values:
	//
	// - **Allow**: allows the traffic that matches the rule.
	//
	// - **Deny**: denies the traffic that matches the rule.
	//
	// This parameter is required.
	//
	// example:
	//
	// Allow
	Action *string `json:"Action,omitempty" xml:"Action,omitempty"`
	// The list of matching conditions for the rule. Each condition consists of a condition key (ConditionKey), an operator (ConditionOperator), and condition values (ConditionValues). The combination of a condition key and an operator must be unique within the same rule.
	//
	// This parameter is required.
	Conditions []*CreateEgressProxyRuleGroupRequestEgressProxyRulesPreDnsConditions `json:"Conditions,omitempty" xml:"Conditions,omitempty" type:"Repeated"`
	// The description of the egress rule. The description must be 2 to 256 characters in length and can contain Chinese characters, letters, digits, hyphens (-), forward slashes (/), periods (.), and underscores (_).
	//
	// example:
	//
	// testEpg
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the egress rule. The name must be 2 to 128 characters in length and must start with a letter or a Chinese character. It can contain digits, periods (.), underscores (_), and hyphens (-).
	//
	// This parameter is required.
	//
	// example:
	//
	// testEpg
	EgressProxyRuleName *string `json:"EgressProxyRuleName,omitempty" xml:"EgressProxyRuleName,omitempty"`
	// The priority of the egress rule within its phase. Valid values: 1 to 1000. A smaller value indicates a higher priority. The priority must be unique within the same phase of a rule group.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
}

func (s CreateEgressProxyRuleGroupRequestEgressProxyRulesPreDns) String() string {
	return dara.Prettify(s)
}

func (s CreateEgressProxyRuleGroupRequestEgressProxyRulesPreDns) GoString() string {
	return s.String()
}

func (s *CreateEgressProxyRuleGroupRequestEgressProxyRulesPreDns) GetAction() *string {
	return s.Action
}

func (s *CreateEgressProxyRuleGroupRequestEgressProxyRulesPreDns) GetConditions() []*CreateEgressProxyRuleGroupRequestEgressProxyRulesPreDnsConditions {
	return s.Conditions
}

func (s *CreateEgressProxyRuleGroupRequestEgressProxyRulesPreDns) GetDescription() *string {
	return s.Description
}

func (s *CreateEgressProxyRuleGroupRequestEgressProxyRulesPreDns) GetEgressProxyRuleName() *string {
	return s.EgressProxyRuleName
}

func (s *CreateEgressProxyRuleGroupRequestEgressProxyRulesPreDns) GetPriority() *int32 {
	return s.Priority
}

func (s *CreateEgressProxyRuleGroupRequestEgressProxyRulesPreDns) SetAction(v string) *CreateEgressProxyRuleGroupRequestEgressProxyRulesPreDns {
	s.Action = &v
	return s
}

func (s *CreateEgressProxyRuleGroupRequestEgressProxyRulesPreDns) SetConditions(v []*CreateEgressProxyRuleGroupRequestEgressProxyRulesPreDnsConditions) *CreateEgressProxyRuleGroupRequestEgressProxyRulesPreDns {
	s.Conditions = v
	return s
}

func (s *CreateEgressProxyRuleGroupRequestEgressProxyRulesPreDns) SetDescription(v string) *CreateEgressProxyRuleGroupRequestEgressProxyRulesPreDns {
	s.Description = &v
	return s
}

func (s *CreateEgressProxyRuleGroupRequestEgressProxyRulesPreDns) SetEgressProxyRuleName(v string) *CreateEgressProxyRuleGroupRequestEgressProxyRulesPreDns {
	s.EgressProxyRuleName = &v
	return s
}

func (s *CreateEgressProxyRuleGroupRequestEgressProxyRulesPreDns) SetPriority(v int32) *CreateEgressProxyRuleGroupRequestEgressProxyRulesPreDns {
	s.Priority = &v
	return s
}

func (s *CreateEgressProxyRuleGroupRequestEgressProxyRulesPreDns) Validate() error {
	if s.Conditions != nil {
		for _, item := range s.Conditions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateEgressProxyRuleGroupRequestEgressProxyRulesPreDnsConditions struct {
	// The key of the matching condition. The value must match the phase of the rule and be used with an operator (ConditionOperator) of the corresponding category. Valid values:
	//
	// - **request:SourceIp**: the source IP address of the request. Used with IP address operators. Supported in the PreDns and PreRequest phases.
	//
	// - **request:DestinationDomain**: the destination domain name of the request. Used with string operators. Supported in the PreDns and PreRequest phases.
	//
	// - **request:Http:Uri**: the URI of the HTTP request. Used with string operators. Supported in the PreRequest phase.
	//
	// - **request:Http:Method**: the HTTP request method. Used with string operators. Supported in the PreRequest phase.
	//
	// - **request:Http:UserAgent**: the User-Agent of the HTTP request. Used with string operators. Supported in the PreRequest phase.
	//
	// - **request:Http:ContentType**: the Content-Type of the HTTP request. Used with string operators. Supported in the PreRequest phase.
	//
	// - **request:Http:Header**: the HTTP request header. Used with string operators. Supported in the PreRequest phase. When you use this key, append the header name to the key in the format of request:Http:Header/<HeaderName>.
	//
	// - **request:Protocol**: the request protocol. Used with string operators. Supported in the PreRequest phase.
	//
	// This parameter is required.
	//
	// example:
	//
	// request:SourceIp
	ConditionKey *string `json:"ConditionKey,omitempty" xml:"ConditionKey,omitempty"`
	// The operator of the matching condition. It must be used with a condition key (ConditionKey). Valid values:
	//
	// - **StringEquals**: exact string match.
	//
	// - **StringNotEquals**: string inequality.
	//
	// - **StringEqualsIgnoreCase**: case-insensitive exact string match.
	//
	// - **StringNotEqualsIgnoreCase**: case-insensitive string inequality.
	//
	// - **StringPrefixMatch**: string prefix match.
	//
	// - **StringPrefixMatchIgnoreCase**: case-insensitive string prefix match.
	//
	// - **StringPrefixNotMatch**: string prefix mismatch.
	//
	// - **StringPrefixNotMatchIgnoreCase**: case-insensitive string prefix mismatch.
	//
	// - **StringSuffixMatch**: string suffix match.
	//
	// - **StringSuffixMatchIgnoreCase**: case-insensitive string suffix match.
	//
	// - **StringSuffixNotMatch**: string suffix mismatch.
	//
	// - **StringSuffixNotMatchIgnoreCase**: case-insensitive string suffix mismatch.
	//
	// - **StringRegexMatch**: string regular expression match.
	//
	// - **StringRegexNotMatch**: string regular expression mismatch.
	//
	// - **StringWildcardMatch**: string wildcard match.
	//
	// - **StringWildcardMatchIgnoreCase**: case-insensitive string wildcard match.
	//
	// - **StringWildcardNotMatch**: string wildcard mismatch.
	//
	// - **StringWildcardNotMatchIgnoreCase**: case-insensitive string wildcard mismatch.
	//
	// - **IpAddress**: The source IP address belongs to the specified IP address or CIDR block.
	//
	// - **NotIpAddress**: The source IP address does not belong to the specified IP address or CIDR block.
	//
	// This parameter is required.
	//
	// example:
	//
	// IpAddress
	ConditionOperator *string `json:"ConditionOperator,omitempty" xml:"ConditionOperator,omitempty"`
	// The list of condition values.
	//
	// This parameter is required.
	ConditionValues []*string `json:"ConditionValues,omitempty" xml:"ConditionValues,omitempty" type:"Repeated"`
}

func (s CreateEgressProxyRuleGroupRequestEgressProxyRulesPreDnsConditions) String() string {
	return dara.Prettify(s)
}

func (s CreateEgressProxyRuleGroupRequestEgressProxyRulesPreDnsConditions) GoString() string {
	return s.String()
}

func (s *CreateEgressProxyRuleGroupRequestEgressProxyRulesPreDnsConditions) GetConditionKey() *string {
	return s.ConditionKey
}

func (s *CreateEgressProxyRuleGroupRequestEgressProxyRulesPreDnsConditions) GetConditionOperator() *string {
	return s.ConditionOperator
}

func (s *CreateEgressProxyRuleGroupRequestEgressProxyRulesPreDnsConditions) GetConditionValues() []*string {
	return s.ConditionValues
}

func (s *CreateEgressProxyRuleGroupRequestEgressProxyRulesPreDnsConditions) SetConditionKey(v string) *CreateEgressProxyRuleGroupRequestEgressProxyRulesPreDnsConditions {
	s.ConditionKey = &v
	return s
}

func (s *CreateEgressProxyRuleGroupRequestEgressProxyRulesPreDnsConditions) SetConditionOperator(v string) *CreateEgressProxyRuleGroupRequestEgressProxyRulesPreDnsConditions {
	s.ConditionOperator = &v
	return s
}

func (s *CreateEgressProxyRuleGroupRequestEgressProxyRulesPreDnsConditions) SetConditionValues(v []*string) *CreateEgressProxyRuleGroupRequestEgressProxyRulesPreDnsConditions {
	s.ConditionValues = v
	return s
}

func (s *CreateEgressProxyRuleGroupRequestEgressProxyRulesPreDnsConditions) Validate() error {
	return dara.Validate(s)
}

type CreateEgressProxyRuleGroupRequestEgressProxyRulesPreRequest struct {
	// The action to perform on the matched traffic after the rule is hit. Valid values:
	//
	// - **Allow**: allows the traffic that matches the rule.
	//
	// - **Deny**: denies the traffic that matches the rule.
	//
	// This parameter is required.
	//
	// example:
	//
	// Allow
	Action *string `json:"Action,omitempty" xml:"Action,omitempty"`
	// The list of matching conditions for the rule. Each condition consists of a condition key (ConditionKey), an operator (ConditionOperator), and condition values (ConditionValues). The combination of a condition key and an operator must be unique within the same rule.
	//
	// This parameter is required.
	Conditions []*CreateEgressProxyRuleGroupRequestEgressProxyRulesPreRequestConditions `json:"Conditions,omitempty" xml:"Conditions,omitempty" type:"Repeated"`
	// The description of the egress rule. The description must be 2 to 256 characters in length and can contain Chinese characters, letters, digits, hyphens (-), forward slashes (/), periods (.), and underscores (_).
	//
	// example:
	//
	// testEpg
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the egress rule. The name must be 2 to 128 characters in length and must start with a letter or a Chinese character. It can contain digits, periods (.), underscores (_), and hyphens (-).
	//
	// This parameter is required.
	//
	// example:
	//
	// testEpg
	EgressProxyRuleName *string `json:"EgressProxyRuleName,omitempty" xml:"EgressProxyRuleName,omitempty"`
	// The priority of the egress rule within its phase. Valid values: 1 to 1000. A smaller value indicates a higher priority. The priority must be unique within the same phase of a rule group.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
}

func (s CreateEgressProxyRuleGroupRequestEgressProxyRulesPreRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateEgressProxyRuleGroupRequestEgressProxyRulesPreRequest) GoString() string {
	return s.String()
}

func (s *CreateEgressProxyRuleGroupRequestEgressProxyRulesPreRequest) GetAction() *string {
	return s.Action
}

func (s *CreateEgressProxyRuleGroupRequestEgressProxyRulesPreRequest) GetConditions() []*CreateEgressProxyRuleGroupRequestEgressProxyRulesPreRequestConditions {
	return s.Conditions
}

func (s *CreateEgressProxyRuleGroupRequestEgressProxyRulesPreRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateEgressProxyRuleGroupRequestEgressProxyRulesPreRequest) GetEgressProxyRuleName() *string {
	return s.EgressProxyRuleName
}

func (s *CreateEgressProxyRuleGroupRequestEgressProxyRulesPreRequest) GetPriority() *int32 {
	return s.Priority
}

func (s *CreateEgressProxyRuleGroupRequestEgressProxyRulesPreRequest) SetAction(v string) *CreateEgressProxyRuleGroupRequestEgressProxyRulesPreRequest {
	s.Action = &v
	return s
}

func (s *CreateEgressProxyRuleGroupRequestEgressProxyRulesPreRequest) SetConditions(v []*CreateEgressProxyRuleGroupRequestEgressProxyRulesPreRequestConditions) *CreateEgressProxyRuleGroupRequestEgressProxyRulesPreRequest {
	s.Conditions = v
	return s
}

func (s *CreateEgressProxyRuleGroupRequestEgressProxyRulesPreRequest) SetDescription(v string) *CreateEgressProxyRuleGroupRequestEgressProxyRulesPreRequest {
	s.Description = &v
	return s
}

func (s *CreateEgressProxyRuleGroupRequestEgressProxyRulesPreRequest) SetEgressProxyRuleName(v string) *CreateEgressProxyRuleGroupRequestEgressProxyRulesPreRequest {
	s.EgressProxyRuleName = &v
	return s
}

func (s *CreateEgressProxyRuleGroupRequestEgressProxyRulesPreRequest) SetPriority(v int32) *CreateEgressProxyRuleGroupRequestEgressProxyRulesPreRequest {
	s.Priority = &v
	return s
}

func (s *CreateEgressProxyRuleGroupRequestEgressProxyRulesPreRequest) Validate() error {
	if s.Conditions != nil {
		for _, item := range s.Conditions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateEgressProxyRuleGroupRequestEgressProxyRulesPreRequestConditions struct {
	// The key of the matching condition. The value must match the phase of the rule and be used with an operator (ConditionOperator) of the corresponding category. Valid values:
	//
	// - **request:SourceIp**: the source IP address of the request. Used with IP address operators. Supported in the PreDns and PreRequest phases.
	//
	// - **request:DestinationDomain**: the destination domain name of the request. Used with string operators. Supported in the PreDns and PreRequest phases.
	//
	// - **request:Http:Uri**: the URI of the HTTP request. Used with string operators. Supported in the PreRequest phase.
	//
	// - **request:Http:Method**: the HTTP request method. Used with string operators. Supported in the PreRequest phase.
	//
	// - **request:Http:UserAgent**: the User-Agent of the HTTP request. Used with string operators. Supported in the PreRequest phase.
	//
	// - **request:Http:ContentType**: the Content-Type of the HTTP request. Used with string operators. Supported in the PreRequest phase.
	//
	// - **request:Http:Header**: the HTTP request header. Used with string operators. Supported in the PreRequest phase. When you use this key, append the header name to the key in the format of request:Http:Header/<HeaderName>.
	//
	// - **request:Protocol**: the request protocol. Used with string operators. Supported in the PreRequest phase.
	//
	// This parameter is required.
	//
	// example:
	//
	// request:DestinationDomain
	ConditionKey *string `json:"ConditionKey,omitempty" xml:"ConditionKey,omitempty"`
	// The operator of the matching condition. It must be used with a condition key (ConditionKey). Valid values:
	//
	// - **StringEquals**: exact string match.
	//
	// - **StringNotEquals**: string inequality.
	//
	// - **StringEqualsIgnoreCase**: case-insensitive exact string match.
	//
	// - **StringNotEqualsIgnoreCase**: case-insensitive string inequality.
	//
	// - **StringPrefixMatch**: string prefix match.
	//
	// - **StringPrefixMatchIgnoreCase**: case-insensitive string prefix match.
	//
	// - **StringPrefixNotMatch**: string prefix mismatch.
	//
	// - **StringPrefixNotMatchIgnoreCase**: case-insensitive string prefix mismatch.
	//
	// - **StringSuffixMatch**: string suffix match.
	//
	// - **StringSuffixMatchIgnoreCase**: case-insensitive string suffix match.
	//
	// - **StringSuffixNotMatch**: string suffix mismatch.
	//
	// - **StringSuffixNotMatchIgnoreCase**: case-insensitive string suffix mismatch.
	//
	// - **StringRegexMatch**: string regular expression match.
	//
	// - **StringRegexNotMatch**: string regular expression mismatch.
	//
	// - **StringWildcardMatch**: string wildcard match.
	//
	// - **StringWildcardMatchIgnoreCase**: case-insensitive string wildcard match.
	//
	// - **StringWildcardNotMatch**: string wildcard mismatch.
	//
	// - **StringWildcardNotMatchIgnoreCase**: case-insensitive string wildcard mismatch.
	//
	// - **IpAddress**: The source IP address belongs to the specified IP address or CIDR block.
	//
	// - **NotIpAddress**: The source IP address does not belong to the specified IP address or CIDR block.
	//
	// This parameter is required.
	//
	// example:
	//
	// StringEquals
	ConditionOperator *string `json:"ConditionOperator,omitempty" xml:"ConditionOperator,omitempty"`
	// The list of condition values.
	//
	// This parameter is required.
	ConditionValues []*string `json:"ConditionValues,omitempty" xml:"ConditionValues,omitempty" type:"Repeated"`
}

func (s CreateEgressProxyRuleGroupRequestEgressProxyRulesPreRequestConditions) String() string {
	return dara.Prettify(s)
}

func (s CreateEgressProxyRuleGroupRequestEgressProxyRulesPreRequestConditions) GoString() string {
	return s.String()
}

func (s *CreateEgressProxyRuleGroupRequestEgressProxyRulesPreRequestConditions) GetConditionKey() *string {
	return s.ConditionKey
}

func (s *CreateEgressProxyRuleGroupRequestEgressProxyRulesPreRequestConditions) GetConditionOperator() *string {
	return s.ConditionOperator
}

func (s *CreateEgressProxyRuleGroupRequestEgressProxyRulesPreRequestConditions) GetConditionValues() []*string {
	return s.ConditionValues
}

func (s *CreateEgressProxyRuleGroupRequestEgressProxyRulesPreRequestConditions) SetConditionKey(v string) *CreateEgressProxyRuleGroupRequestEgressProxyRulesPreRequestConditions {
	s.ConditionKey = &v
	return s
}

func (s *CreateEgressProxyRuleGroupRequestEgressProxyRulesPreRequestConditions) SetConditionOperator(v string) *CreateEgressProxyRuleGroupRequestEgressProxyRulesPreRequestConditions {
	s.ConditionOperator = &v
	return s
}

func (s *CreateEgressProxyRuleGroupRequestEgressProxyRulesPreRequestConditions) SetConditionValues(v []*string) *CreateEgressProxyRuleGroupRequestEgressProxyRulesPreRequestConditions {
	s.ConditionValues = v
	return s
}

func (s *CreateEgressProxyRuleGroupRequestEgressProxyRulesPreRequestConditions) Validate() error {
	return dara.Validate(s)
}

type CreateEgressProxyRuleGroupRequestTag struct {
	// The tag key. It can be up to 128 characters in length, cannot start with `aliyun` or `acs:`, and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// testKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value. It can be up to 256 characters in length and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// testValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEgressProxyRuleGroupRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateEgressProxyRuleGroupRequestTag) GoString() string {
	return s.String()
}

func (s *CreateEgressProxyRuleGroupRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateEgressProxyRuleGroupRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateEgressProxyRuleGroupRequestTag) SetKey(v string) *CreateEgressProxyRuleGroupRequestTag {
	s.Key = &v
	return s
}

func (s *CreateEgressProxyRuleGroupRequestTag) SetValue(v string) *CreateEgressProxyRuleGroupRequestTag {
	s.Value = &v
	return s
}

func (s *CreateEgressProxyRuleGroupRequestTag) Validate() error {
	return dara.Validate(s)
}
