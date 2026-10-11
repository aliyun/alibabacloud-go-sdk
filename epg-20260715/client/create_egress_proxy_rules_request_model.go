// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEgressProxyRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateEgressProxyRulesRequest
	GetClientToken() *string
	SetDryRun(v bool) *CreateEgressProxyRulesRequest
	GetDryRun() *bool
	SetEgressProxyRuleGroupId(v string) *CreateEgressProxyRulesRequest
	GetEgressProxyRuleGroupId() *string
	SetEgressProxyRules(v *CreateEgressProxyRulesRequestEgressProxyRules) *CreateEgressProxyRulesRequest
	GetEgressProxyRules() *CreateEgressProxyRulesRequestEgressProxyRules
}

type CreateEgressProxyRulesRequest struct {
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
	// Specifies whether to perform only a dry run on the request. Valid values:
	//
	// - **true**: performs a dry run. The system checks the request for potential issues, including missing required parameters, invalid request formats, and service limits, but does not create the rules. If the request fails the dry run, an error message is returned. If the request passes the dry run, the DryRunOperation error code is returned.
	//
	// - **false*	- (default): performs a dry run and sends the request. If the request passes the dry run, an HTTP 2xx status code is returned and the operation is performed.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ID of the egress proxy rule group.
	//
	// This parameter is required.
	//
	// example:
	//
	// eprg-uf60966a9633667c****
	EgressProxyRuleGroupId *string `json:"EgressProxyRuleGroupId,omitempty" xml:"EgressProxyRuleGroupId,omitempty"`
	// The list of egress rules grouped by effective phase (PreDns, PreRequest, or PostResponse).
	//
	// This parameter is required.
	EgressProxyRules *CreateEgressProxyRulesRequestEgressProxyRules `json:"EgressProxyRules,omitempty" xml:"EgressProxyRules,omitempty" type:"Struct"`
}

func (s CreateEgressProxyRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateEgressProxyRulesRequest) GoString() string {
	return s.String()
}

func (s *CreateEgressProxyRulesRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateEgressProxyRulesRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *CreateEgressProxyRulesRequest) GetEgressProxyRuleGroupId() *string {
	return s.EgressProxyRuleGroupId
}

func (s *CreateEgressProxyRulesRequest) GetEgressProxyRules() *CreateEgressProxyRulesRequestEgressProxyRules {
	return s.EgressProxyRules
}

func (s *CreateEgressProxyRulesRequest) SetClientToken(v string) *CreateEgressProxyRulesRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateEgressProxyRulesRequest) SetDryRun(v bool) *CreateEgressProxyRulesRequest {
	s.DryRun = &v
	return s
}

func (s *CreateEgressProxyRulesRequest) SetEgressProxyRuleGroupId(v string) *CreateEgressProxyRulesRequest {
	s.EgressProxyRuleGroupId = &v
	return s
}

func (s *CreateEgressProxyRulesRequest) SetEgressProxyRules(v *CreateEgressProxyRulesRequestEgressProxyRules) *CreateEgressProxyRulesRequest {
	s.EgressProxyRules = v
	return s
}

func (s *CreateEgressProxyRulesRequest) Validate() error {
	if s.EgressProxyRules != nil {
		if err := s.EgressProxyRules.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateEgressProxyRulesRequestEgressProxyRules struct {
	// The list of rules in the PostResponse phase (after the response is returned).
	PostResponse []*CreateEgressProxyRulesRequestEgressProxyRulesPostResponse `json:"PostResponse,omitempty" xml:"PostResponse,omitempty" type:"Repeated"`
	// The list of rules for the pre-DNS stage (before DNS resolution).
	PreDns []*CreateEgressProxyRulesRequestEgressProxyRulesPreDns `json:"PreDns,omitempty" xml:"PreDns,omitempty" type:"Repeated"`
	// The list of rules in the PreRequest stage (before request forwarding).
	PreRequest []*CreateEgressProxyRulesRequestEgressProxyRulesPreRequest `json:"PreRequest,omitempty" xml:"PreRequest,omitempty" type:"Repeated"`
}

func (s CreateEgressProxyRulesRequestEgressProxyRules) String() string {
	return dara.Prettify(s)
}

func (s CreateEgressProxyRulesRequestEgressProxyRules) GoString() string {
	return s.String()
}

func (s *CreateEgressProxyRulesRequestEgressProxyRules) GetPostResponse() []*CreateEgressProxyRulesRequestEgressProxyRulesPostResponse {
	return s.PostResponse
}

func (s *CreateEgressProxyRulesRequestEgressProxyRules) GetPreDns() []*CreateEgressProxyRulesRequestEgressProxyRulesPreDns {
	return s.PreDns
}

func (s *CreateEgressProxyRulesRequestEgressProxyRules) GetPreRequest() []*CreateEgressProxyRulesRequestEgressProxyRulesPreRequest {
	return s.PreRequest
}

func (s *CreateEgressProxyRulesRequestEgressProxyRules) SetPostResponse(v []*CreateEgressProxyRulesRequestEgressProxyRulesPostResponse) *CreateEgressProxyRulesRequestEgressProxyRules {
	s.PostResponse = v
	return s
}

func (s *CreateEgressProxyRulesRequestEgressProxyRules) SetPreDns(v []*CreateEgressProxyRulesRequestEgressProxyRulesPreDns) *CreateEgressProxyRulesRequestEgressProxyRules {
	s.PreDns = v
	return s
}

func (s *CreateEgressProxyRulesRequestEgressProxyRules) SetPreRequest(v []*CreateEgressProxyRulesRequestEgressProxyRulesPreRequest) *CreateEgressProxyRulesRequestEgressProxyRules {
	s.PreRequest = v
	return s
}

func (s *CreateEgressProxyRulesRequestEgressProxyRules) Validate() error {
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

type CreateEgressProxyRulesRequestEgressProxyRulesPostResponse struct {
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
	Conditions []*CreateEgressProxyRulesRequestEgressProxyRulesPostResponseConditions `json:"Conditions,omitempty" xml:"Conditions,omitempty" type:"Repeated"`
	// The description of the egress rule. The description must be 2 to 256 characters in length, and can contain Chinese characters, letters, digits, hyphens (-), forward slashes (/), periods (.), and underscores (_).
	//
	// example:
	//
	// testEpg
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the egress rule. The name must be 2 to 128 characters in length, and must start with a letter or a Chinese character. It can contain digits, periods (.), underscores (_), and hyphens (-).
	//
	// This parameter is required.
	//
	// example:
	//
	// testEpg
	EgressProxyRuleName *string `json:"EgressProxyRuleName,omitempty" xml:"EgressProxyRuleName,omitempty"`
	// The priority of the egress rule within the phase to which it belongs. Valid values: 1 to 1000. A smaller value indicates a higher priority. Within the same phase of a rule group, the priority of rules must be unique.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
}

func (s CreateEgressProxyRulesRequestEgressProxyRulesPostResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateEgressProxyRulesRequestEgressProxyRulesPostResponse) GoString() string {
	return s.String()
}

func (s *CreateEgressProxyRulesRequestEgressProxyRulesPostResponse) GetAction() *string {
	return s.Action
}

func (s *CreateEgressProxyRulesRequestEgressProxyRulesPostResponse) GetConditions() []*CreateEgressProxyRulesRequestEgressProxyRulesPostResponseConditions {
	return s.Conditions
}

func (s *CreateEgressProxyRulesRequestEgressProxyRulesPostResponse) GetDescription() *string {
	return s.Description
}

func (s *CreateEgressProxyRulesRequestEgressProxyRulesPostResponse) GetEgressProxyRuleName() *string {
	return s.EgressProxyRuleName
}

func (s *CreateEgressProxyRulesRequestEgressProxyRulesPostResponse) GetPriority() *int32 {
	return s.Priority
}

func (s *CreateEgressProxyRulesRequestEgressProxyRulesPostResponse) SetAction(v string) *CreateEgressProxyRulesRequestEgressProxyRulesPostResponse {
	s.Action = &v
	return s
}

func (s *CreateEgressProxyRulesRequestEgressProxyRulesPostResponse) SetConditions(v []*CreateEgressProxyRulesRequestEgressProxyRulesPostResponseConditions) *CreateEgressProxyRulesRequestEgressProxyRulesPostResponse {
	s.Conditions = v
	return s
}

func (s *CreateEgressProxyRulesRequestEgressProxyRulesPostResponse) SetDescription(v string) *CreateEgressProxyRulesRequestEgressProxyRulesPostResponse {
	s.Description = &v
	return s
}

func (s *CreateEgressProxyRulesRequestEgressProxyRulesPostResponse) SetEgressProxyRuleName(v string) *CreateEgressProxyRulesRequestEgressProxyRulesPostResponse {
	s.EgressProxyRuleName = &v
	return s
}

func (s *CreateEgressProxyRulesRequestEgressProxyRulesPostResponse) SetPriority(v int32) *CreateEgressProxyRulesRequestEgressProxyRulesPostResponse {
	s.Priority = &v
	return s
}

func (s *CreateEgressProxyRulesRequestEgressProxyRulesPostResponse) Validate() error {
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

type CreateEgressProxyRulesRequestEgressProxyRulesPostResponseConditions struct {
	// The key of the matching condition. The value must match the effective phase (Phase) of the rule and be used with an operator (ConditionOperator) of the corresponding category. Valid values:
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
	// - **request:Http:Header**: the HTTP request header. Used with string operators. Supported in the PreRequest phase. When you use this key, append the request header name to the key in the format of request:Http:Header/<HeaderName>.
	//
	// - **request:Protocol**: the request protocol. Used with string operators. Supported in the PreRequest phase.
	//
	// This parameter is required.
	//
	// example:
	//
	// request:SourceIp
	ConditionKey *string `json:"ConditionKey,omitempty" xml:"ConditionKey,omitempty"`
	// The operator of the matching condition. Must be used with a condition key (ConditionKey). Valid values:
	//
	// - **StringEquals**: exact string match.
	//
	// - **StringNotEquals**: string not equal.
	//
	// - **StringEqualsIgnoreCase**: case-insensitive exact string match.
	//
	// - **StringNotEqualsIgnoreCase**: case-insensitive string not equal.
	//
	// - **StringPrefixMatch**: string prefix match.
	//
	// - **StringPrefixMatchIgnoreCase**: case-insensitive string prefix match.
	//
	// - **StringPrefixNotMatch**: string prefix not match.
	//
	// - **StringPrefixNotMatchIgnoreCase**: case-insensitive string prefix not match.
	//
	// - **StringSuffixMatch**: string suffix match.
	//
	// - **StringSuffixMatchIgnoreCase**: case-insensitive string suffix match.
	//
	// - **StringSuffixNotMatch**: string suffix not match.
	//
	// - **StringSuffixNotMatchIgnoreCase**: case-insensitive string suffix not match.
	//
	// - **StringRegexMatch**: string regular expression match.
	//
	// - **StringRegexNotMatch**: string regular expression not match.
	//
	// - **StringWildcardMatch**: string wildcard match.
	//
	// - **StringWildcardMatchIgnoreCase**: case-insensitive string wildcard match.
	//
	// - **StringWildcardNotMatch**: string wildcard not match.
	//
	// - **StringWildcardNotMatchIgnoreCase**: case-insensitive string wildcard not match.
	//
	// - **IpAddress**: the source IP address belongs to the specified IP address or CIDR block.
	//
	// - **NotIpAddress**: the source IP address does not belong to the specified IP address or CIDR block.
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

func (s CreateEgressProxyRulesRequestEgressProxyRulesPostResponseConditions) String() string {
	return dara.Prettify(s)
}

func (s CreateEgressProxyRulesRequestEgressProxyRulesPostResponseConditions) GoString() string {
	return s.String()
}

func (s *CreateEgressProxyRulesRequestEgressProxyRulesPostResponseConditions) GetConditionKey() *string {
	return s.ConditionKey
}

func (s *CreateEgressProxyRulesRequestEgressProxyRulesPostResponseConditions) GetConditionOperator() *string {
	return s.ConditionOperator
}

func (s *CreateEgressProxyRulesRequestEgressProxyRulesPostResponseConditions) GetConditionValues() []*string {
	return s.ConditionValues
}

func (s *CreateEgressProxyRulesRequestEgressProxyRulesPostResponseConditions) SetConditionKey(v string) *CreateEgressProxyRulesRequestEgressProxyRulesPostResponseConditions {
	s.ConditionKey = &v
	return s
}

func (s *CreateEgressProxyRulesRequestEgressProxyRulesPostResponseConditions) SetConditionOperator(v string) *CreateEgressProxyRulesRequestEgressProxyRulesPostResponseConditions {
	s.ConditionOperator = &v
	return s
}

func (s *CreateEgressProxyRulesRequestEgressProxyRulesPostResponseConditions) SetConditionValues(v []*string) *CreateEgressProxyRulesRequestEgressProxyRulesPostResponseConditions {
	s.ConditionValues = v
	return s
}

func (s *CreateEgressProxyRulesRequestEgressProxyRulesPostResponseConditions) Validate() error {
	return dara.Validate(s)
}

type CreateEgressProxyRulesRequestEgressProxyRulesPreDns struct {
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
	Conditions []*CreateEgressProxyRulesRequestEgressProxyRulesPreDnsConditions `json:"Conditions,omitempty" xml:"Conditions,omitempty" type:"Repeated"`
	// The description of the egress rule. The description must be 2 to 256 characters in length and can contain Chinese characters, letters, digits, hyphens (-), forward slashes (/), periods (.), and underscores (_).
	//
	// example:
	//
	// testEpg
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the egress rule. The name must be 2 to 128 characters in length, and must start with a letter or a Chinese character. It can contain digits, periods (.), underscores (_), and hyphens (-).
	//
	// This parameter is required.
	//
	// example:
	//
	// testEpg
	EgressProxyRuleName *string `json:"EgressProxyRuleName,omitempty" xml:"EgressProxyRuleName,omitempty"`
	// The priority of the egress rule within the phase to which it belongs. Valid values: 1 to 1000. A smaller value indicates a higher priority. Within the same phase of a rule group, the priority of rules must be unique.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
}

func (s CreateEgressProxyRulesRequestEgressProxyRulesPreDns) String() string {
	return dara.Prettify(s)
}

func (s CreateEgressProxyRulesRequestEgressProxyRulesPreDns) GoString() string {
	return s.String()
}

func (s *CreateEgressProxyRulesRequestEgressProxyRulesPreDns) GetAction() *string {
	return s.Action
}

func (s *CreateEgressProxyRulesRequestEgressProxyRulesPreDns) GetConditions() []*CreateEgressProxyRulesRequestEgressProxyRulesPreDnsConditions {
	return s.Conditions
}

func (s *CreateEgressProxyRulesRequestEgressProxyRulesPreDns) GetDescription() *string {
	return s.Description
}

func (s *CreateEgressProxyRulesRequestEgressProxyRulesPreDns) GetEgressProxyRuleName() *string {
	return s.EgressProxyRuleName
}

func (s *CreateEgressProxyRulesRequestEgressProxyRulesPreDns) GetPriority() *int32 {
	return s.Priority
}

func (s *CreateEgressProxyRulesRequestEgressProxyRulesPreDns) SetAction(v string) *CreateEgressProxyRulesRequestEgressProxyRulesPreDns {
	s.Action = &v
	return s
}

func (s *CreateEgressProxyRulesRequestEgressProxyRulesPreDns) SetConditions(v []*CreateEgressProxyRulesRequestEgressProxyRulesPreDnsConditions) *CreateEgressProxyRulesRequestEgressProxyRulesPreDns {
	s.Conditions = v
	return s
}

func (s *CreateEgressProxyRulesRequestEgressProxyRulesPreDns) SetDescription(v string) *CreateEgressProxyRulesRequestEgressProxyRulesPreDns {
	s.Description = &v
	return s
}

func (s *CreateEgressProxyRulesRequestEgressProxyRulesPreDns) SetEgressProxyRuleName(v string) *CreateEgressProxyRulesRequestEgressProxyRulesPreDns {
	s.EgressProxyRuleName = &v
	return s
}

func (s *CreateEgressProxyRulesRequestEgressProxyRulesPreDns) SetPriority(v int32) *CreateEgressProxyRulesRequestEgressProxyRulesPreDns {
	s.Priority = &v
	return s
}

func (s *CreateEgressProxyRulesRequestEgressProxyRulesPreDns) Validate() error {
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

type CreateEgressProxyRulesRequestEgressProxyRulesPreDnsConditions struct {
	// The key of the matching condition. The value must match the effective phase (Phase) of the rule and be used with an operator (ConditionOperator) of the corresponding category. Valid values:
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
	// - **request:Http:Header**: the HTTP request header. Used with string operators. Supported in the PreRequest phase. When you use this key, append the request header name to the key in the format of request:Http:Header/<HeaderName>.
	//
	// - **request:Protocol**: the request protocol. Used with string operators. Supported in the PreRequest phase.
	//
	// This parameter is required.
	//
	// example:
	//
	// request:DestinationDomain
	ConditionKey *string `json:"ConditionKey,omitempty" xml:"ConditionKey,omitempty"`
	// The operator of the matching condition. Must be used with a condition key (ConditionKey). Valid values:
	//
	// - **StringEquals**: exact string match.
	//
	// - **StringNotEquals**: string not equal.
	//
	// - **StringEqualsIgnoreCase**: case-insensitive exact string match.
	//
	// - **StringNotEqualsIgnoreCase**: case-insensitive string not equal.
	//
	// - **StringPrefixMatch**: string prefix match.
	//
	// - **StringPrefixMatchIgnoreCase**: case-insensitive string prefix match.
	//
	// - **StringPrefixNotMatch**: string prefix not match.
	//
	// - **StringPrefixNotMatchIgnoreCase**: case-insensitive string prefix not match.
	//
	// - **StringSuffixMatch**: string suffix match.
	//
	// - **StringSuffixMatchIgnoreCase**: case-insensitive string suffix match.
	//
	// - **StringSuffixNotMatch**: string suffix not match.
	//
	// - **StringSuffixNotMatchIgnoreCase**: case-insensitive string suffix not match.
	//
	// - **StringRegexMatch**: string regular expression match.
	//
	// - **StringRegexNotMatch**: string regular expression not match.
	//
	// - **StringWildcardMatch**: string wildcard match.
	//
	// - **StringWildcardMatchIgnoreCase**: case-insensitive string wildcard match.
	//
	// - **StringWildcardNotMatch**: string wildcard not match.
	//
	// - **StringWildcardNotMatchIgnoreCase**: case-insensitive string wildcard not match.
	//
	// - **IpAddress**: the source IP address belongs to the specified IP address or CIDR block.
	//
	// - **NotIpAddress**: the source IP address does not belong to the specified IP address or CIDR block.
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

func (s CreateEgressProxyRulesRequestEgressProxyRulesPreDnsConditions) String() string {
	return dara.Prettify(s)
}

func (s CreateEgressProxyRulesRequestEgressProxyRulesPreDnsConditions) GoString() string {
	return s.String()
}

func (s *CreateEgressProxyRulesRequestEgressProxyRulesPreDnsConditions) GetConditionKey() *string {
	return s.ConditionKey
}

func (s *CreateEgressProxyRulesRequestEgressProxyRulesPreDnsConditions) GetConditionOperator() *string {
	return s.ConditionOperator
}

func (s *CreateEgressProxyRulesRequestEgressProxyRulesPreDnsConditions) GetConditionValues() []*string {
	return s.ConditionValues
}

func (s *CreateEgressProxyRulesRequestEgressProxyRulesPreDnsConditions) SetConditionKey(v string) *CreateEgressProxyRulesRequestEgressProxyRulesPreDnsConditions {
	s.ConditionKey = &v
	return s
}

func (s *CreateEgressProxyRulesRequestEgressProxyRulesPreDnsConditions) SetConditionOperator(v string) *CreateEgressProxyRulesRequestEgressProxyRulesPreDnsConditions {
	s.ConditionOperator = &v
	return s
}

func (s *CreateEgressProxyRulesRequestEgressProxyRulesPreDnsConditions) SetConditionValues(v []*string) *CreateEgressProxyRulesRequestEgressProxyRulesPreDnsConditions {
	s.ConditionValues = v
	return s
}

func (s *CreateEgressProxyRulesRequestEgressProxyRulesPreDnsConditions) Validate() error {
	return dara.Validate(s)
}

type CreateEgressProxyRulesRequestEgressProxyRulesPreRequest struct {
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
	Conditions []*CreateEgressProxyRulesRequestEgressProxyRulesPreRequestConditions `json:"Conditions,omitempty" xml:"Conditions,omitempty" type:"Repeated"`
	// The description of the egress rule. The description must be 2 to 256 characters in length, and can contain Chinese characters, letters, digits, hyphens (-), forward slashes (/), periods (.), and underscores (_).
	//
	// example:
	//
	// testEpg
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the egress rule. The name must be 2 to 128 characters in length, and must start with a letter or a Chinese character. It can contain digits, periods (.), underscores (_), and hyphens (-).
	//
	// This parameter is required.
	//
	// example:
	//
	// testEpg
	EgressProxyRuleName *string `json:"EgressProxyRuleName,omitempty" xml:"EgressProxyRuleName,omitempty"`
	// The priority of the egress rule within the phase to which it belongs. Valid values: 1 to 1000. A smaller value indicates a higher priority. Within the same phase of a rule group, the priority of rules must be unique.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
}

func (s CreateEgressProxyRulesRequestEgressProxyRulesPreRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateEgressProxyRulesRequestEgressProxyRulesPreRequest) GoString() string {
	return s.String()
}

func (s *CreateEgressProxyRulesRequestEgressProxyRulesPreRequest) GetAction() *string {
	return s.Action
}

func (s *CreateEgressProxyRulesRequestEgressProxyRulesPreRequest) GetConditions() []*CreateEgressProxyRulesRequestEgressProxyRulesPreRequestConditions {
	return s.Conditions
}

func (s *CreateEgressProxyRulesRequestEgressProxyRulesPreRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateEgressProxyRulesRequestEgressProxyRulesPreRequest) GetEgressProxyRuleName() *string {
	return s.EgressProxyRuleName
}

func (s *CreateEgressProxyRulesRequestEgressProxyRulesPreRequest) GetPriority() *int32 {
	return s.Priority
}

func (s *CreateEgressProxyRulesRequestEgressProxyRulesPreRequest) SetAction(v string) *CreateEgressProxyRulesRequestEgressProxyRulesPreRequest {
	s.Action = &v
	return s
}

func (s *CreateEgressProxyRulesRequestEgressProxyRulesPreRequest) SetConditions(v []*CreateEgressProxyRulesRequestEgressProxyRulesPreRequestConditions) *CreateEgressProxyRulesRequestEgressProxyRulesPreRequest {
	s.Conditions = v
	return s
}

func (s *CreateEgressProxyRulesRequestEgressProxyRulesPreRequest) SetDescription(v string) *CreateEgressProxyRulesRequestEgressProxyRulesPreRequest {
	s.Description = &v
	return s
}

func (s *CreateEgressProxyRulesRequestEgressProxyRulesPreRequest) SetEgressProxyRuleName(v string) *CreateEgressProxyRulesRequestEgressProxyRulesPreRequest {
	s.EgressProxyRuleName = &v
	return s
}

func (s *CreateEgressProxyRulesRequestEgressProxyRulesPreRequest) SetPriority(v int32) *CreateEgressProxyRulesRequestEgressProxyRulesPreRequest {
	s.Priority = &v
	return s
}

func (s *CreateEgressProxyRulesRequestEgressProxyRulesPreRequest) Validate() error {
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

type CreateEgressProxyRulesRequestEgressProxyRulesPreRequestConditions struct {
	// The key of the matching condition. The value must match the effective phase (Phase) of the rule and be used with an operator (ConditionOperator) of the corresponding category. Valid values:
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
	// - **request:Http:Header**: the HTTP request header. Used with string operators. Supported in the PreRequest phase. When you use this key, append the request header name to the key in the format of request:Http:Header/<HeaderName>.
	//
	// - **request:Protocol**: the request protocol. Used with string operators. Supported in the PreRequest phase.
	//
	// This parameter is required.
	//
	// example:
	//
	// request:DestinationDomain
	ConditionKey *string `json:"ConditionKey,omitempty" xml:"ConditionKey,omitempty"`
	// The operator of the matching condition. Must be used with a condition key (ConditionKey). Valid values:
	//
	// - **StringEquals**: exact string match.
	//
	// - **StringNotEquals**: string not equal.
	//
	// - **StringEqualsIgnoreCase**: case-insensitive exact string match.
	//
	// - **StringNotEqualsIgnoreCase**: case-insensitive string not equal.
	//
	// - **StringPrefixMatch**: string prefix match.
	//
	// - **StringPrefixMatchIgnoreCase**: case-insensitive string prefix match.
	//
	// - **StringPrefixNotMatch**: string prefix not match.
	//
	// - **StringPrefixNotMatchIgnoreCase**: case-insensitive string prefix not match.
	//
	// - **StringSuffixMatch**: string suffix match.
	//
	// - **StringSuffixMatchIgnoreCase**: case-insensitive string suffix match.
	//
	// - **StringSuffixNotMatch**: string suffix not match.
	//
	// - **StringSuffixNotMatchIgnoreCase**: case-insensitive string suffix not match.
	//
	// - **StringRegexMatch**: string regular expression match.
	//
	// - **StringRegexNotMatch**: string regular expression not match.
	//
	// - **StringWildcardMatch**: string wildcard match.
	//
	// - **StringWildcardMatchIgnoreCase**: case-insensitive string wildcard match.
	//
	// - **StringWildcardNotMatch**: string wildcard not match.
	//
	// - **StringWildcardNotMatchIgnoreCase**: case-insensitive string wildcard not match.
	//
	// - **IpAddress**: the source IP address belongs to the specified IP address or CIDR block.
	//
	// - **NotIpAddress**: the source IP address does not belong to the specified IP address or CIDR block.
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

func (s CreateEgressProxyRulesRequestEgressProxyRulesPreRequestConditions) String() string {
	return dara.Prettify(s)
}

func (s CreateEgressProxyRulesRequestEgressProxyRulesPreRequestConditions) GoString() string {
	return s.String()
}

func (s *CreateEgressProxyRulesRequestEgressProxyRulesPreRequestConditions) GetConditionKey() *string {
	return s.ConditionKey
}

func (s *CreateEgressProxyRulesRequestEgressProxyRulesPreRequestConditions) GetConditionOperator() *string {
	return s.ConditionOperator
}

func (s *CreateEgressProxyRulesRequestEgressProxyRulesPreRequestConditions) GetConditionValues() []*string {
	return s.ConditionValues
}

func (s *CreateEgressProxyRulesRequestEgressProxyRulesPreRequestConditions) SetConditionKey(v string) *CreateEgressProxyRulesRequestEgressProxyRulesPreRequestConditions {
	s.ConditionKey = &v
	return s
}

func (s *CreateEgressProxyRulesRequestEgressProxyRulesPreRequestConditions) SetConditionOperator(v string) *CreateEgressProxyRulesRequestEgressProxyRulesPreRequestConditions {
	s.ConditionOperator = &v
	return s
}

func (s *CreateEgressProxyRulesRequestEgressProxyRulesPreRequestConditions) SetConditionValues(v []*string) *CreateEgressProxyRulesRequestEgressProxyRulesPreRequestConditions {
	s.ConditionValues = v
	return s
}

func (s *CreateEgressProxyRulesRequestEgressProxyRulesPreRequestConditions) Validate() error {
	return dara.Validate(s)
}
