// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateForwardingRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceleratorId(v string) *UpdateForwardingRulesRequest
	GetAcceleratorId() *string
	SetClientToken(v string) *UpdateForwardingRulesRequest
	GetClientToken() *string
	SetForwardingRules(v []*UpdateForwardingRulesRequestForwardingRules) *UpdateForwardingRulesRequest
	GetForwardingRules() []*UpdateForwardingRulesRequestForwardingRules
	SetListenerId(v string) *UpdateForwardingRulesRequest
	GetListenerId() *string
	SetRegionId(v string) *UpdateForwardingRulesRequest
	GetRegionId() *string
}

type UpdateForwardingRulesRequest struct {
	// The GA instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ga-bp17frjjh0udz4q****
	AcceleratorId *string `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The client token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 02fb3da4****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The configurations of the forwarding rules.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	ForwardingRules []*UpdateForwardingRulesRequestForwardingRules `json:"ForwardingRules,omitempty" xml:"ForwardingRules,omitempty" type:"Repeated"`
	// The listener ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// lsr-bp1s0vzbi5bxlx5****
	ListenerId *string `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	// The region ID of the GA instance. Set the value to **cn-hangzhou**.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s UpdateForwardingRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateForwardingRulesRequest) GoString() string {
	return s.String()
}

func (s *UpdateForwardingRulesRequest) GetAcceleratorId() *string {
	return s.AcceleratorId
}

func (s *UpdateForwardingRulesRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateForwardingRulesRequest) GetForwardingRules() []*UpdateForwardingRulesRequestForwardingRules {
	return s.ForwardingRules
}

func (s *UpdateForwardingRulesRequest) GetListenerId() *string {
	return s.ListenerId
}

func (s *UpdateForwardingRulesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateForwardingRulesRequest) SetAcceleratorId(v string) *UpdateForwardingRulesRequest {
	s.AcceleratorId = &v
	return s
}

func (s *UpdateForwardingRulesRequest) SetClientToken(v string) *UpdateForwardingRulesRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateForwardingRulesRequest) SetForwardingRules(v []*UpdateForwardingRulesRequestForwardingRules) *UpdateForwardingRulesRequest {
	s.ForwardingRules = v
	return s
}

func (s *UpdateForwardingRulesRequest) SetListenerId(v string) *UpdateForwardingRulesRequest {
	s.ListenerId = &v
	return s
}

func (s *UpdateForwardingRulesRequest) SetRegionId(v string) *UpdateForwardingRulesRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateForwardingRulesRequest) Validate() error {
	if s.ForwardingRules != nil {
		for _, item := range s.ForwardingRules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateForwardingRulesRequestForwardingRules struct {
	// The forwarding rule ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// frule-bp1dii16gu9qdvb34****
	ForwardingRuleId *string `json:"ForwardingRuleId,omitempty" xml:"ForwardingRuleId,omitempty"`
	// The forwarding rule name.
	//
	// The name must be 2 to 128 characters in length, and can contain letters, digits, periods (.), underscores (_), and hyphens (-). The name must start with a letter.
	//
	// example:
	//
	// test
	ForwardingRuleName *string `json:"ForwardingRuleName,omitempty" xml:"ForwardingRuleName,omitempty"`
	// The priority of the forwarding rule. Valid values: **1*	- to **10000**. A smaller value specifies a higher priority.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1000
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The configurations of the forwarding actions.
	//
	// This parameter is required.
	RuleActions []*UpdateForwardingRulesRequestForwardingRulesRuleActions `json:"RuleActions,omitempty" xml:"RuleActions,omitempty" type:"Repeated"`
	// The conditions that trigger the forwarding rule.
	//
	// This parameter is required.
	RuleConditions []*UpdateForwardingRulesRequestForwardingRulesRuleConditions `json:"RuleConditions,omitempty" xml:"RuleConditions,omitempty" type:"Repeated"`
	// The direction in which the rule takes effect. You do not need to specify this parameter.
	//
	// By default, this parameter is set to **request**, which specifies that the rule takes effect on requests.
	//
	// example:
	//
	// request
	RuleDirection *string `json:"RuleDirection,omitempty" xml:"RuleDirection,omitempty"`
}

func (s UpdateForwardingRulesRequestForwardingRules) String() string {
	return dara.Prettify(s)
}

func (s UpdateForwardingRulesRequestForwardingRules) GoString() string {
	return s.String()
}

func (s *UpdateForwardingRulesRequestForwardingRules) GetForwardingRuleId() *string {
	return s.ForwardingRuleId
}

func (s *UpdateForwardingRulesRequestForwardingRules) GetForwardingRuleName() *string {
	return s.ForwardingRuleName
}

func (s *UpdateForwardingRulesRequestForwardingRules) GetPriority() *int32 {
	return s.Priority
}

func (s *UpdateForwardingRulesRequestForwardingRules) GetRuleActions() []*UpdateForwardingRulesRequestForwardingRulesRuleActions {
	return s.RuleActions
}

func (s *UpdateForwardingRulesRequestForwardingRules) GetRuleConditions() []*UpdateForwardingRulesRequestForwardingRulesRuleConditions {
	return s.RuleConditions
}

func (s *UpdateForwardingRulesRequestForwardingRules) GetRuleDirection() *string {
	return s.RuleDirection
}

func (s *UpdateForwardingRulesRequestForwardingRules) SetForwardingRuleId(v string) *UpdateForwardingRulesRequestForwardingRules {
	s.ForwardingRuleId = &v
	return s
}

func (s *UpdateForwardingRulesRequestForwardingRules) SetForwardingRuleName(v string) *UpdateForwardingRulesRequestForwardingRules {
	s.ForwardingRuleName = &v
	return s
}

func (s *UpdateForwardingRulesRequestForwardingRules) SetPriority(v int32) *UpdateForwardingRulesRequestForwardingRules {
	s.Priority = &v
	return s
}

func (s *UpdateForwardingRulesRequestForwardingRules) SetRuleActions(v []*UpdateForwardingRulesRequestForwardingRulesRuleActions) *UpdateForwardingRulesRequestForwardingRules {
	s.RuleActions = v
	return s
}

func (s *UpdateForwardingRulesRequestForwardingRules) SetRuleConditions(v []*UpdateForwardingRulesRequestForwardingRulesRuleConditions) *UpdateForwardingRulesRequestForwardingRules {
	s.RuleConditions = v
	return s
}

func (s *UpdateForwardingRulesRequestForwardingRules) SetRuleDirection(v string) *UpdateForwardingRulesRequestForwardingRules {
	s.RuleDirection = &v
	return s
}

func (s *UpdateForwardingRulesRequestForwardingRules) Validate() error {
	if s.RuleActions != nil {
		for _, item := range s.RuleActions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.RuleConditions != nil {
		for _, item := range s.RuleConditions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateForwardingRulesRequestForwardingRulesRuleActions struct {
	// The forwarding configuration.
	//
	// >  We recommend that you use **RuleActionType*	- and **RuleActionValue*	- rather than this parameter to configure forwarding actions.
	ForwardGroupConfig *UpdateForwardingRulesRequestForwardingRulesRuleActionsForwardGroupConfig `json:"ForwardGroupConfig,omitempty" xml:"ForwardGroupConfig,omitempty" type:"Struct"`
	// The forwarding priority.
	//
	// >  This parameter does not take effect. Ignore this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20
	Order *int32 `json:"Order,omitempty" xml:"Order,omitempty"`
	// The type of the forwarding action. Valid values:
	//
	// 	- **ForwardGroup**: forwards a request.
	//
	// 	- **Redirect**: redirects a request.
	//
	// 	- **FixResponse**: returns a fixed response.
	//
	// 	- **Rewrite**: rewrites a request.
	//
	// 	- **AddHeader**: adds a header to a request.
	//
	// 	- **RemoveHeaderConfig**: deletes the header of a request.
	//
	// 	- **Drop**: drops a request.
	//
	// This parameter is required.
	//
	// example:
	//
	// ForwardGroup
	RuleActionType *string `json:"RuleActionType,omitempty" xml:"RuleActionType,omitempty"`
	// The value of the forwarding action.
	//
	// You must specify different JSON strings based on **RuleActionType**.
	//
	// A forwarding rule can contain only one forwarding action whose type is **ForwardGroup**, **Redirect**, or **FixResponse**. You must specify a forwarding action whose type is **Rewrite**, **AddHeader**, or **RemoveHeader*	- before a forwarding action whose type is **ForwardGroup**.
	//
	// 	- If **RuleActionType*	- is set to **ForwardGroup**, this parameter specifies the information of a virtual endpoint group. You can forward requests to only one virtual endpoint group. Example: `{"type":"endpointgroup", "value":"epg-bp1enpdcrqhl78g6r****"}`.
	//
	//     	- `type`: Set this parameter to `endpointgroup`.
	//
	//     	- `value`: Set this parameter to the ID of a virtual endpoint group.
	//
	// 	- If **RuleActionType*	- is set to **Redirect**, this parameter specifies redirecting configurations. You cannot leave the following parameters empty or use the default values for the following parameters for a forwarding action whose type is **Redirect**: `protocol`, `domain`, `port`, `path`, and `query`. Example: `{"protocol":"HTTP", "domain":"www.example.com", "port":"80", "path":"/a","query":"value1", "code":"301" }`.
	//
	//     	- `protocol`: the protocol of requests after the requests are redirected. Valid values: `${protocol}` (default), `HTTP`, and `HTTPS`.
	//
	//     	- `domain`: the domain name to which requests are redirected. Default value: `${host}`. You can also enter a domain name. The domain name must be 3 to 128 characters in length, and can contain only letters, digits, and the following special characters: `. - ? = ~ _ - + / ^ 	- ! $ & | ( ) [ ]`.
	//
	//     	- `port`: the port to which requests are redirected. Default value: `${port}`. You can enter a port number from 1 to 63335.
	//
	//     	- `path`: the path to which requests are redirected. Default value: `${path}`. The path must be 1 to 128 characters in length. To use a regular expression, the path can contain letters, digits, and the following special characters: `. - _ / = ? ~ ^ 	- $ : ( ) [ ] + |`. The path must start with a tilde (~). If you do not want to use a regular expression, the path can contain letters, digits, and the following special characters: `. - _ / = ? :`. The path must start with a forward slash (/).
	//
	//     	- `query`: the query string of the requests that are redirected. Default value: `${query}`. You can also specify a query string. The query string must be 1 to 128 characters in length, and can contain printable characters whose ASCII values are `greater than or equal to 32 and smaller than 127`. The query string cannot contain uppercase letters, space characters, or the following special characters: `[ ] { } < > # | &`.
	//
	//     	- `code`: the redirect code. Valid values: `301`, `302`, `303`, `307`, and `308`.
	//
	// 	- If **RuleActionType*	- is set to **FixResponse**, this parameter specifies a fixed response. Example: `{"code":"200", "type":"text/plain", "content":"dssacav" }`.
	//
	//     	- `code`: the HTTP response status code. The response status code must be one of the following numeric strings: `2xx`, `4xx`, and `5xx`. The letter `x` indicates a number from 0 to 9.
	//
	//     	- `type`: the type of the response content. Valid values: **text/plain**, **text/css**, **text/html**, **application/javascript**, and **application/json**.
	//
	//     	- `content`: the response content. The response content cannot exceed 1,000 characters in length, and does not support Chinese characters.
	//
	// 	- If **RuleActionType*	- is set to **AddHeader**, this parameter specifies an HTTP header to be added. If a forwarding rule contains a forwarding action whose type is **AddHeader**, you must specify another forwarding action whose type is **ForwardGroup**. Example: `[{"name":"header1","type":"userdefined", "value":"value"}]`.
	//
	//     	- `name`: the name of the HTTP header. The name must be 1 to 40 characters in length, and can contain letters, digits, hyphens (-), and underscores (_). The name of the HTTP header specified by **AddHeader*	- must be unique and cannot be the same as the name of the HTTP header specified by **RemoveHeader**.
	//
	//     	- `type`: the content type of the HTTP header. Valid values: `user-defined`, `ref`, and `system-defined`.
	//
	//     	- `value`: the content of the HTTP header. You cannot leave this parameter empty. If you set `type` to `user-defined`, the content must be 1 to 128 characters in length, and can contain printable characters whose ASCII values are `greater than or equal to 32 and smaller than 127`. The content can contain letters, digits, hyphens (-), and underscores (_*). The content cannot start or end with a space character. If you set `type` to `ref`, the content must be 1 to 128 characters in length, and can contain letters, digits, hyphens (-), and underscores (_*). The content cannot start or end with a space character. If you set `type` to `system-defined`, only `ClientSrcIp` is supported.
	//
	// 	- If **RuleActionType*	- is set to **RemoveHeader**, this parameter specifies an HTTP header to be removed. If a forwarding rule contains a forwarding action whose type is **RemoveHeader**, you must specify another forwarding action whose type is **ForwardGroup**. The header must be 1 to 40 characters in length, and can contain letters, digits, hyphens (-), and underscores (_). Example: `["header1"]`.
	//
	// 	- If **RuleActionType*	- is set to **Rewrite**, this parameter specifies the rewriting configuration. If a forwarding rule contains a forwarding action whose type is **Rewrite**, you must specify another forwarding action whose type is **ForwardGroup**. Example: `{"domain":"value1", "path":"value2", "query":"value3"}`.
	//
	//     	- `domain`: the domain name to which requests are redirected. Default value: `${host}`. You can also enter a domain name. The domain name must be 3 to 128 characters in length, and can contain only lowercase letters, digits, and the following special characters: `. - ? = ~ _ - + / ^ 	- ! $ & | ( ) [ ]`.
	//
	//     	- `path`: the path to which requests are redirected. Default value: `${path}`. The path must be 1 to 128 characters in length. To use a regular expression, the path can contain letters, digits, and the following special characters: `. - _ / = ? ~ ^ 	- $ : ( ) [ ] + |`. The path must start with a tilde (~). If you do not want to use a regular expression, the path can contain letters, digits, and the following special characters: `. - _ / = ? :`. The path must start with a forward slash (/).
	//
	//     	- `query`: the query string of the requests that are redirected. Default value: `${query}`. You can also specify a query string. The query string must be 1 to 128 characters in length, and can contain printable characters whose ASCII values are `greater than or equal to 32 and smaller than 127`. The query string cannot contain uppercase letters, space characters, or the following special characters: `[ ] { } < > # | &`.
	//
	// 	- If **RuleActionType*	- is set to **Drop**, you do not need to specify this parameter.
	//
	// example:
	//
	// [{"type":"endpointgroup", "value":"epg-bp1enpdcrqhl78g6r****"}]
	RuleActionValue *string `json:"RuleActionValue,omitempty" xml:"RuleActionValue,omitempty"`
}

func (s UpdateForwardingRulesRequestForwardingRulesRuleActions) String() string {
	return dara.Prettify(s)
}

func (s UpdateForwardingRulesRequestForwardingRulesRuleActions) GoString() string {
	return s.String()
}

func (s *UpdateForwardingRulesRequestForwardingRulesRuleActions) GetForwardGroupConfig() *UpdateForwardingRulesRequestForwardingRulesRuleActionsForwardGroupConfig {
	return s.ForwardGroupConfig
}

func (s *UpdateForwardingRulesRequestForwardingRulesRuleActions) GetOrder() *int32 {
	return s.Order
}

func (s *UpdateForwardingRulesRequestForwardingRulesRuleActions) GetRuleActionType() *string {
	return s.RuleActionType
}

func (s *UpdateForwardingRulesRequestForwardingRulesRuleActions) GetRuleActionValue() *string {
	return s.RuleActionValue
}

func (s *UpdateForwardingRulesRequestForwardingRulesRuleActions) SetForwardGroupConfig(v *UpdateForwardingRulesRequestForwardingRulesRuleActionsForwardGroupConfig) *UpdateForwardingRulesRequestForwardingRulesRuleActions {
	s.ForwardGroupConfig = v
	return s
}

func (s *UpdateForwardingRulesRequestForwardingRulesRuleActions) SetOrder(v int32) *UpdateForwardingRulesRequestForwardingRulesRuleActions {
	s.Order = &v
	return s
}

func (s *UpdateForwardingRulesRequestForwardingRulesRuleActions) SetRuleActionType(v string) *UpdateForwardingRulesRequestForwardingRulesRuleActions {
	s.RuleActionType = &v
	return s
}

func (s *UpdateForwardingRulesRequestForwardingRulesRuleActions) SetRuleActionValue(v string) *UpdateForwardingRulesRequestForwardingRulesRuleActions {
	s.RuleActionValue = &v
	return s
}

func (s *UpdateForwardingRulesRequestForwardingRulesRuleActions) Validate() error {
	if s.ForwardGroupConfig != nil {
		if err := s.ForwardGroupConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateForwardingRulesRequestForwardingRulesRuleActionsForwardGroupConfig struct {
	// The configuration of an endpoint group.
	//
	// >  We recommend that you use **RuleActionType*	- and **RuleActionValue*	- rather than this parameter to configure forwarding actions.
	//
	// This parameter is required.
	ServerGroupTuples []*UpdateForwardingRulesRequestForwardingRulesRuleActionsForwardGroupConfigServerGroupTuples `json:"ServerGroupTuples,omitempty" xml:"ServerGroupTuples,omitempty" type:"Repeated"`
}

func (s UpdateForwardingRulesRequestForwardingRulesRuleActionsForwardGroupConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateForwardingRulesRequestForwardingRulesRuleActionsForwardGroupConfig) GoString() string {
	return s.String()
}

func (s *UpdateForwardingRulesRequestForwardingRulesRuleActionsForwardGroupConfig) GetServerGroupTuples() []*UpdateForwardingRulesRequestForwardingRulesRuleActionsForwardGroupConfigServerGroupTuples {
	return s.ServerGroupTuples
}

func (s *UpdateForwardingRulesRequestForwardingRulesRuleActionsForwardGroupConfig) SetServerGroupTuples(v []*UpdateForwardingRulesRequestForwardingRulesRuleActionsForwardGroupConfigServerGroupTuples) *UpdateForwardingRulesRequestForwardingRulesRuleActionsForwardGroupConfig {
	s.ServerGroupTuples = v
	return s
}

func (s *UpdateForwardingRulesRequestForwardingRulesRuleActionsForwardGroupConfig) Validate() error {
	if s.ServerGroupTuples != nil {
		for _, item := range s.ServerGroupTuples {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateForwardingRulesRequestForwardingRulesRuleActionsForwardGroupConfigServerGroupTuples struct {
	// The ID of the endpoint group.
	//
	// >  We recommend that you use **RuleActionType*	- and **RuleActionValue*	- rather than this parameter to configure forwarding actions.
	//
	// This parameter is required.
	//
	// example:
	//
	// epg-bp1ieei9664r5nv****
	EndpointGroupId *string `json:"EndpointGroupId,omitempty" xml:"EndpointGroupId,omitempty"`
}

func (s UpdateForwardingRulesRequestForwardingRulesRuleActionsForwardGroupConfigServerGroupTuples) String() string {
	return dara.Prettify(s)
}

func (s UpdateForwardingRulesRequestForwardingRulesRuleActionsForwardGroupConfigServerGroupTuples) GoString() string {
	return s.String()
}

func (s *UpdateForwardingRulesRequestForwardingRulesRuleActionsForwardGroupConfigServerGroupTuples) GetEndpointGroupId() *string {
	return s.EndpointGroupId
}

func (s *UpdateForwardingRulesRequestForwardingRulesRuleActionsForwardGroupConfigServerGroupTuples) SetEndpointGroupId(v string) *UpdateForwardingRulesRequestForwardingRulesRuleActionsForwardGroupConfigServerGroupTuples {
	s.EndpointGroupId = &v
	return s
}

func (s *UpdateForwardingRulesRequestForwardingRulesRuleActionsForwardGroupConfigServerGroupTuples) Validate() error {
	return dara.Validate(s)
}

type UpdateForwardingRulesRequestForwardingRulesRuleConditions struct {
	// The domain name configuration.
	//
	// >  We recommend that you use **RuleConditionType*	- and **RuleConditionValue*	- rather than this parameter to configure forwarding conditions.
	HostConfig *UpdateForwardingRulesRequestForwardingRulesRuleConditionsHostConfig `json:"HostConfig,omitempty" xml:"HostConfig,omitempty" type:"Struct"`
	// The path configuration.
	//
	// >  We recommend that you use **RuleConditionType*	- and **RuleConditionValue*	- rather than this parameter to configure forwarding conditions.
	PathConfig *UpdateForwardingRulesRequestForwardingRulesRuleConditionsPathConfig `json:"PathConfig,omitempty" xml:"PathConfig,omitempty" type:"Struct"`
	// The type of the forwarding condition. Valid values:
	//
	// 	- **Host**: Requests are forwarded based on domain names.
	//
	// 	- **Path**: Requests are forwarded based on paths.
	//
	// 	- **RequestHeader**: Requests are forwarded based on HTTP headers.
	//
	// 	- **Query**: Requests are forwarded based on query strings.
	//
	// 	- **Method**: Requests are forwarded based on HTTP request methods.
	//
	// 	- **Cookie**: Requests are forwarded based on cookies.
	//
	// 	- **SourceIp**: Requests are forwarded based on source IP addresses.
	//
	// This parameter is required.
	//
	// example:
	//
	// Host
	RuleConditionType *string `json:"RuleConditionType,omitempty" xml:"RuleConditionType,omitempty"`
	// The value of the forwarding condition. You must specify different JSON strings based on **RuleConditionType**.
	//
	// 	- If **RuleConditionType*	- is set to **Host**, RuleConditionValue specifies a domain name condition. A forwarding rule can contain only one forwarding condition of the host type. You can specify multiple domain names in a forwarding condition. The relationship between multiple domain names is OR. The domain name must be 3 to 128 characters in length, and can contain letters, digits, hyphens (-), and periods (.). You can use asterisks (\\*) and question marks (?) as wildcard characters. Example: `["www.example.com", "www.aliyun.com"]`.
	//
	// 	- If **RuleConditionType*	- is set to **Path**, RuleConditionValue specifies a path condition. A forwarding rule can contain multiple forwarding conditions of the path type. The relationship between multiple path conditions is OR. You can specify multiple paths in a forwarding condition. The relationship between multiple paths is OR. The path must be 1 to 128 characters in length, and must start with a forward slash (/). The path can contain letters, digits, and the following special characters: $ - _ . + / & ~ @ : \\". Supported wildcard characters are asterisks (\\*) and question marks (?). Example: `["/a", "/b/"]`.
	//
	// 	- If **RuleConditionType*	- is set to **RequestHeader**, RuleConditionValue specifies an HTTP header condition. An HTTP header consists of a key and a value. The header values in a forwarding condition must be unique. Example: `[{"header1":["value1","value2"]}]`.
	//
	//     	- Key: The key of an HTTP header must be 1 to 40 characters in length and can contain letters, digits, hyphens (-), and underscores (_).
	//
	//     	- Value: The value of an HTTP header must be 1 to 128 characters in length and can contain printable characters whose ASCII values `are larger than or equal to 32 and smaller than 127`. The value cannot start or end with a space.
	//
	// 	- If **RuleConditionType*	- is set to **Query**, RuleConditionValue specifies a query string condition. A query string consists of a key and a value. Example: `[{"query1":["value1"]}, {"query2":["value2"]}]`.
	//
	//     	- Key: The key must be 1 to 100 characters in length and can contain printable characters whose ASCII values `are larger than or equal to 32 and smaller than 127`, excluding uppercase letters, spaces, and the following special characters: `[ ] { } < > \\ ; / ? : @ & = + , $ % " ^ ~`.
	//
	//     	- Value: The value must be 1 to 128 characters in length and can contain printable characters whose ASCII values `are larger than or equal to 32 and smaller than 127`, excluding uppercase letters, spaces, and the following special characters: `[ ] { } < > \\ ; / ? : @ & = + , $ % " ^ ~`.
	//
	// 	- If **RuleConditionType*	- is set to **Method**, RuleConditionValue specifies an HTTP method condition. Valid values: **HEAD**, **GET**, **POST**, **OPTIONS**, **PUT**, **PATCH**, and **DELETE**. Example: `["GET", "OPTIONS", "POST"]`.
	//
	// 	- If **RuleConditionType*	- is set to **Cookie**, RuleConditionValue specifies a cookie condition. A cookie consists of a key and a value. Example: `[{"cookie1":["value1"]}, {"cookie2":["value2"]}]`.
	//
	//     	- Key: The key of a cookie must be 1 to 100 characters in length and can contain printable characters whose ASCII values `are larger than or equal to 32 and smaller than 127`, excluding uppercase letters, spaces, and the following special characters: `# [ ] { } \\ < > &`.
	//
	//     	- Value: The value of a cookie must be 1 to 128 characters in length and can contain printable characters whose ASCII values `are larger than or equal to 32 and smaller than 127`, excluding uppercase letters, spaces, and the following special characters: `# [ ] { } \\ < > &`.
	//
	// 	- If **RuleConditionType*	- is set to **SourceIP**, RuleConditionValue specifies a source IP address condition. IP addresses, such as 1.1.XX.XX/32, and CIDR blocks, such as 2.2.XX.XX/24, are supported. A forwarding rule can contain only one forwarding condition of the SourceIP type. You can specify multiple source IP addresses or CIDR blocks in a forwarding condition. The relationship between multiple IP addresses or CIDR blocks is OR. Example: `["1.1.XX.XX/32", "2.2.XX.XX/24"]`.
	//
	// example:
	//
	// ["www.example.com", "www.aliyun.com"]
	RuleConditionValue *string `json:"RuleConditionValue,omitempty" xml:"RuleConditionValue,omitempty"`
}

func (s UpdateForwardingRulesRequestForwardingRulesRuleConditions) String() string {
	return dara.Prettify(s)
}

func (s UpdateForwardingRulesRequestForwardingRulesRuleConditions) GoString() string {
	return s.String()
}

func (s *UpdateForwardingRulesRequestForwardingRulesRuleConditions) GetHostConfig() *UpdateForwardingRulesRequestForwardingRulesRuleConditionsHostConfig {
	return s.HostConfig
}

func (s *UpdateForwardingRulesRequestForwardingRulesRuleConditions) GetPathConfig() *UpdateForwardingRulesRequestForwardingRulesRuleConditionsPathConfig {
	return s.PathConfig
}

func (s *UpdateForwardingRulesRequestForwardingRulesRuleConditions) GetRuleConditionType() *string {
	return s.RuleConditionType
}

func (s *UpdateForwardingRulesRequestForwardingRulesRuleConditions) GetRuleConditionValue() *string {
	return s.RuleConditionValue
}

func (s *UpdateForwardingRulesRequestForwardingRulesRuleConditions) SetHostConfig(v *UpdateForwardingRulesRequestForwardingRulesRuleConditionsHostConfig) *UpdateForwardingRulesRequestForwardingRulesRuleConditions {
	s.HostConfig = v
	return s
}

func (s *UpdateForwardingRulesRequestForwardingRulesRuleConditions) SetPathConfig(v *UpdateForwardingRulesRequestForwardingRulesRuleConditionsPathConfig) *UpdateForwardingRulesRequestForwardingRulesRuleConditions {
	s.PathConfig = v
	return s
}

func (s *UpdateForwardingRulesRequestForwardingRulesRuleConditions) SetRuleConditionType(v string) *UpdateForwardingRulesRequestForwardingRulesRuleConditions {
	s.RuleConditionType = &v
	return s
}

func (s *UpdateForwardingRulesRequestForwardingRulesRuleConditions) SetRuleConditionValue(v string) *UpdateForwardingRulesRequestForwardingRulesRuleConditions {
	s.RuleConditionValue = &v
	return s
}

func (s *UpdateForwardingRulesRequestForwardingRulesRuleConditions) Validate() error {
	if s.HostConfig != nil {
		if err := s.HostConfig.Validate(); err != nil {
			return err
		}
	}
	if s.PathConfig != nil {
		if err := s.PathConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateForwardingRulesRequestForwardingRulesRuleConditionsHostConfig struct {
	// The domain name configuration.
	//
	// >  We recommend that you use **RuleConditionType*	- and **RuleConditionValue*	- rather than this parameter to configure forwarding conditions.
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s UpdateForwardingRulesRequestForwardingRulesRuleConditionsHostConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateForwardingRulesRequestForwardingRulesRuleConditionsHostConfig) GoString() string {
	return s.String()
}

func (s *UpdateForwardingRulesRequestForwardingRulesRuleConditionsHostConfig) GetValues() []*string {
	return s.Values
}

func (s *UpdateForwardingRulesRequestForwardingRulesRuleConditionsHostConfig) SetValues(v []*string) *UpdateForwardingRulesRequestForwardingRulesRuleConditionsHostConfig {
	s.Values = v
	return s
}

func (s *UpdateForwardingRulesRequestForwardingRulesRuleConditionsHostConfig) Validate() error {
	return dara.Validate(s)
}

type UpdateForwardingRulesRequestForwardingRulesRuleConditionsPathConfig struct {
	// The path configuration.
	//
	// >  We recommend that you use **RuleConditionType*	- and **RuleConditionValue*	- rather than this parameter to configure forwarding conditions.
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s UpdateForwardingRulesRequestForwardingRulesRuleConditionsPathConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateForwardingRulesRequestForwardingRulesRuleConditionsPathConfig) GoString() string {
	return s.String()
}

func (s *UpdateForwardingRulesRequestForwardingRulesRuleConditionsPathConfig) GetValues() []*string {
	return s.Values
}

func (s *UpdateForwardingRulesRequestForwardingRulesRuleConditionsPathConfig) SetValues(v []*string) *UpdateForwardingRulesRequestForwardingRulesRuleConditionsPathConfig {
	s.Values = v
	return s
}

func (s *UpdateForwardingRulesRequestForwardingRulesRuleConditionsPathConfig) Validate() error {
	return dara.Validate(s)
}
