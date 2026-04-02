// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateForwardingRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceleratorId(v string) *CreateForwardingRulesRequest
	GetAcceleratorId() *string
	SetClientToken(v string) *CreateForwardingRulesRequest
	GetClientToken() *string
	SetForwardingRules(v []*CreateForwardingRulesRequestForwardingRules) *CreateForwardingRulesRequest
	GetForwardingRules() []*CreateForwardingRulesRequestForwardingRules
	SetListenerId(v string) *CreateForwardingRulesRequest
	GetListenerId() *string
	SetRegionId(v string) *CreateForwardingRulesRequest
	GetRegionId() *string
}

type CreateForwardingRulesRequest struct {
	// The ID of the GA instance.
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
	ForwardingRules []*CreateForwardingRulesRequestForwardingRules `json:"ForwardingRules,omitempty" xml:"ForwardingRules,omitempty" type:"Repeated"`
	// The ID of the listener.
	//
	// This parameter is required.
	//
	// example:
	//
	// lsr-bp1s0vzbi5bxlx5****
	ListenerId *string `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	// The ID of the region where the GA instance is deployed. Set the value to **cn-hangzhou**.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateForwardingRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateForwardingRulesRequest) GoString() string {
	return s.String()
}

func (s *CreateForwardingRulesRequest) GetAcceleratorId() *string {
	return s.AcceleratorId
}

func (s *CreateForwardingRulesRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateForwardingRulesRequest) GetForwardingRules() []*CreateForwardingRulesRequestForwardingRules {
	return s.ForwardingRules
}

func (s *CreateForwardingRulesRequest) GetListenerId() *string {
	return s.ListenerId
}

func (s *CreateForwardingRulesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateForwardingRulesRequest) SetAcceleratorId(v string) *CreateForwardingRulesRequest {
	s.AcceleratorId = &v
	return s
}

func (s *CreateForwardingRulesRequest) SetClientToken(v string) *CreateForwardingRulesRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateForwardingRulesRequest) SetForwardingRules(v []*CreateForwardingRulesRequestForwardingRules) *CreateForwardingRulesRequest {
	s.ForwardingRules = v
	return s
}

func (s *CreateForwardingRulesRequest) SetListenerId(v string) *CreateForwardingRulesRequest {
	s.ListenerId = &v
	return s
}

func (s *CreateForwardingRulesRequest) SetRegionId(v string) *CreateForwardingRulesRequest {
	s.RegionId = &v
	return s
}

func (s *CreateForwardingRulesRequest) Validate() error {
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

type CreateForwardingRulesRequestForwardingRules struct {
	// The name of the forwarding rule. The name must be 2 to 128 characters in length, and can contain letters, digits, periods (.), underscores (_), and hyphens (-). The name must start with a letter.
	//
	// example:
	//
	// test
	ForwardingRuleName *string `json:"ForwardingRuleName,omitempty" xml:"ForwardingRuleName,omitempty"`
	// The priority of the forwarding rule. Valid values: **1*	- to **10000**. A lower value specifies a higher priority.
	//
	// example:
	//
	// 1000
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The configurations of the forwarding action.
	//
	// This parameter is required.
	RuleActions []*CreateForwardingRulesRequestForwardingRulesRuleActions `json:"RuleActions,omitempty" xml:"RuleActions,omitempty" type:"Repeated"`
	// The forwarding conditions.
	//
	// This parameter is required.
	RuleConditions []*CreateForwardingRulesRequestForwardingRulesRuleConditions `json:"RuleConditions,omitempty" xml:"RuleConditions,omitempty" type:"Repeated"`
	// The direction in which the rule takes effect. You do not need to specify this parameter.
	//
	// By default, this parameter is set to **request**, which specifies that the rule takes effect on requests.
	//
	// example:
	//
	// request
	RuleDirection *string `json:"RuleDirection,omitempty" xml:"RuleDirection,omitempty"`
}

func (s CreateForwardingRulesRequestForwardingRules) String() string {
	return dara.Prettify(s)
}

func (s CreateForwardingRulesRequestForwardingRules) GoString() string {
	return s.String()
}

func (s *CreateForwardingRulesRequestForwardingRules) GetForwardingRuleName() *string {
	return s.ForwardingRuleName
}

func (s *CreateForwardingRulesRequestForwardingRules) GetPriority() *int32 {
	return s.Priority
}

func (s *CreateForwardingRulesRequestForwardingRules) GetRuleActions() []*CreateForwardingRulesRequestForwardingRulesRuleActions {
	return s.RuleActions
}

func (s *CreateForwardingRulesRequestForwardingRules) GetRuleConditions() []*CreateForwardingRulesRequestForwardingRulesRuleConditions {
	return s.RuleConditions
}

func (s *CreateForwardingRulesRequestForwardingRules) GetRuleDirection() *string {
	return s.RuleDirection
}

func (s *CreateForwardingRulesRequestForwardingRules) SetForwardingRuleName(v string) *CreateForwardingRulesRequestForwardingRules {
	s.ForwardingRuleName = &v
	return s
}

func (s *CreateForwardingRulesRequestForwardingRules) SetPriority(v int32) *CreateForwardingRulesRequestForwardingRules {
	s.Priority = &v
	return s
}

func (s *CreateForwardingRulesRequestForwardingRules) SetRuleActions(v []*CreateForwardingRulesRequestForwardingRulesRuleActions) *CreateForwardingRulesRequestForwardingRules {
	s.RuleActions = v
	return s
}

func (s *CreateForwardingRulesRequestForwardingRules) SetRuleConditions(v []*CreateForwardingRulesRequestForwardingRulesRuleConditions) *CreateForwardingRulesRequestForwardingRules {
	s.RuleConditions = v
	return s
}

func (s *CreateForwardingRulesRequestForwardingRules) SetRuleDirection(v string) *CreateForwardingRulesRequestForwardingRules {
	s.RuleDirection = &v
	return s
}

func (s *CreateForwardingRulesRequestForwardingRules) Validate() error {
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

type CreateForwardingRulesRequestForwardingRulesRuleActions struct {
	// The forwarding configurations.
	//
	// >  We recommend that you use **RuleActionType*	- and **RuleActionValue*	- rather than this parameter to configure forwarding conditions.
	ForwardGroupConfig *CreateForwardingRulesRequestForwardingRulesRuleActionsForwardGroupConfig `json:"ForwardGroupConfig,omitempty" xml:"ForwardGroupConfig,omitempty" type:"Struct"`
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
	// 	- **RemoveHeaderConfig**: deletes the header from a request.
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
	// You must specify different JSON strings based on the **RuleActionType*	- parameter.
	//
	// A forwarding rule can contain only one forwarding action whose type is **ForwardGroup**, **Redirect**, or **FixResponse**. You must specify a forwarding action whose type is **Rewrite**, **AddHeader**, or **RemoveHeader*	- before a forwarding action whose type is **ForwardGroup**.
	//
	// 	- If **RuleActionType*	- is set to **ForwardGroup**, this parameter specifies the information of a virtual endpoint group. You can forward requests to only one virtual endpoint group. Example: `{"type":"endpointgroup", "value":"epg-bp1enpdcrqhl78g6r****"}`.
	//
	//     	- `type`: Set this parameter to `endpointgroup`.
	//
	//     	- `value`: Set this parameter to the ID of a virtual endpoint group.
	//
	// 	- If **RuleActionType*	- is set to **Redirect**, this parameter specifies redirecting configurations. You cannot leave all the following parameters empty or use the default values for all the following parameters for a forwarding action whose type is **Redirect**: `protocol`, `domain`, `port`, `path`, and `query`. Example: `{"protocol":"HTTP", "domain":"www.example.com", "port":"80", "path":"/a","query":"value1", "code":"301" }`.
	//
	//     	- `protocol`: the protocol of requests after the requests are redirected. Valid values: `${protocol}` (default), `HTTP`, and `HTTPS`.
	//
	//     	- `domain`: the domain name to which requests are redirected. Default value: `${host}`. You can also enter a domain name. The domain name must be 3 to 128 characters in length, and can contain only letters, digits, and the following special characters: `. - ? = ~ _ - + / ^ 	- ! $ & | ( ) [ ]`.
	//
	//     	- `port`: the port to which requests are redirected. Default value: `${port}`. You can enter a port number that ranges from 1 to 63335.
	//
	//     	- `path`: the path to which requests are redirected. Default value: `${path}`. The path must be 1 to 128 characters in length. To use a regular expression, the path can contain letters, digits, and the following special characters: `. - _ / = ? ~ ^ 	- $ : ( ) [ ] +`. The path must start with a tilde (~). If you do not want to use a regular expression, the path can contain letters, digits, and the following special characters: `. - _ / = ? :`. The path must start with a forward slash (/).
	//
	//     	- `query`: the query string to which requests are redirected. Default value: `${query}`. You can also specify a query string. The query string must be 1 to 128 characters in length, and can contain printable characters whose ASCII values are `greater than or equal to 32 and smaller than 127`. The query string cannot contain uppercase letters, space characters, or the following special characters: `[ ] { } < > # &`.
	//
	//     	- `code`: the redirect code. Valid values: `301`, `302`, `303`, `307`, and `308`.
	//
	// 	- If **RuleActionType*	- is set to **FixResponse**, this parameter specifies a fixed response. Example: `{"code":"200", "type":"text/plain", "content":"dssacav" }`.
	//
	//     	- `code`: the HTTP response status code. The response status code must be one of the following numeric strings: `2xx`, `4xx`, and `5xx`. The letter `x` is a digit.
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
	//     	- `path`: the path to which requests are redirected. Default value: `${path}`. The path must be 1 to 128 characters in length. To use a regular expression, the path can contain letters, digits, and the following special characters: `. - _ / = ? ~ ^ 	- $ : ( ) [ ] +`. The path must start with a tilde (~). If you do not want to use a regular expression, the path can contain letters, digits, and the following special characters: `. - _ / = ? :`. The path must start with a forward slash (/).
	//
	//     	- `query`: the query string to which requests are redirected. Default value: `${query}`. You can also specify a query string. The query string must be 1 to 128 characters in length, and can contain printable characters whose ASCII values are `greater than or equal to 32 and smaller than 127`. The query string cannot contain uppercase letters, space characters, or the following special characters: `[ ] { } < > # &`.
	//
	// 	- If **RuleActionType*	- is set to **Drop**, you do not need to specify this parameter.
	//
	// example:
	//
	// [{"type":"endpointgroup", "value":"epg-bp1enpdcrqhl78g6r****"}]
	RuleActionValue *string `json:"RuleActionValue,omitempty" xml:"RuleActionValue,omitempty"`
}

func (s CreateForwardingRulesRequestForwardingRulesRuleActions) String() string {
	return dara.Prettify(s)
}

func (s CreateForwardingRulesRequestForwardingRulesRuleActions) GoString() string {
	return s.String()
}

func (s *CreateForwardingRulesRequestForwardingRulesRuleActions) GetForwardGroupConfig() *CreateForwardingRulesRequestForwardingRulesRuleActionsForwardGroupConfig {
	return s.ForwardGroupConfig
}

func (s *CreateForwardingRulesRequestForwardingRulesRuleActions) GetOrder() *int32 {
	return s.Order
}

func (s *CreateForwardingRulesRequestForwardingRulesRuleActions) GetRuleActionType() *string {
	return s.RuleActionType
}

func (s *CreateForwardingRulesRequestForwardingRulesRuleActions) GetRuleActionValue() *string {
	return s.RuleActionValue
}

func (s *CreateForwardingRulesRequestForwardingRulesRuleActions) SetForwardGroupConfig(v *CreateForwardingRulesRequestForwardingRulesRuleActionsForwardGroupConfig) *CreateForwardingRulesRequestForwardingRulesRuleActions {
	s.ForwardGroupConfig = v
	return s
}

func (s *CreateForwardingRulesRequestForwardingRulesRuleActions) SetOrder(v int32) *CreateForwardingRulesRequestForwardingRulesRuleActions {
	s.Order = &v
	return s
}

func (s *CreateForwardingRulesRequestForwardingRulesRuleActions) SetRuleActionType(v string) *CreateForwardingRulesRequestForwardingRulesRuleActions {
	s.RuleActionType = &v
	return s
}

func (s *CreateForwardingRulesRequestForwardingRulesRuleActions) SetRuleActionValue(v string) *CreateForwardingRulesRequestForwardingRulesRuleActions {
	s.RuleActionValue = &v
	return s
}

func (s *CreateForwardingRulesRequestForwardingRulesRuleActions) Validate() error {
	if s.ForwardGroupConfig != nil {
		if err := s.ForwardGroupConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateForwardingRulesRequestForwardingRulesRuleActionsForwardGroupConfig struct {
	// The configurations of an endpoint group.
	//
	// >  We recommend that you use **RuleActionType*	- and **RuleActionValue*	- rather than this parameter to configure forwarding conditions.
	//
	// This parameter is required.
	ServerGroupTuples []*CreateForwardingRulesRequestForwardingRulesRuleActionsForwardGroupConfigServerGroupTuples `json:"ServerGroupTuples,omitempty" xml:"ServerGroupTuples,omitempty" type:"Repeated"`
}

func (s CreateForwardingRulesRequestForwardingRulesRuleActionsForwardGroupConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateForwardingRulesRequestForwardingRulesRuleActionsForwardGroupConfig) GoString() string {
	return s.String()
}

func (s *CreateForwardingRulesRequestForwardingRulesRuleActionsForwardGroupConfig) GetServerGroupTuples() []*CreateForwardingRulesRequestForwardingRulesRuleActionsForwardGroupConfigServerGroupTuples {
	return s.ServerGroupTuples
}

func (s *CreateForwardingRulesRequestForwardingRulesRuleActionsForwardGroupConfig) SetServerGroupTuples(v []*CreateForwardingRulesRequestForwardingRulesRuleActionsForwardGroupConfigServerGroupTuples) *CreateForwardingRulesRequestForwardingRulesRuleActionsForwardGroupConfig {
	s.ServerGroupTuples = v
	return s
}

func (s *CreateForwardingRulesRequestForwardingRulesRuleActionsForwardGroupConfig) Validate() error {
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

type CreateForwardingRulesRequestForwardingRulesRuleActionsForwardGroupConfigServerGroupTuples struct {
	// The ID of the endpoint group.
	//
	// >  We recommend that you use **RuleActionType*	- and **RuleActionValue*	- rather than this parameter to configure forwarding conditions.
	//
	// This parameter is required.
	//
	// example:
	//
	// epg-bp1ieei9664r5nv****
	EndpointGroupId *string `json:"EndpointGroupId,omitempty" xml:"EndpointGroupId,omitempty"`
}

func (s CreateForwardingRulesRequestForwardingRulesRuleActionsForwardGroupConfigServerGroupTuples) String() string {
	return dara.Prettify(s)
}

func (s CreateForwardingRulesRequestForwardingRulesRuleActionsForwardGroupConfigServerGroupTuples) GoString() string {
	return s.String()
}

func (s *CreateForwardingRulesRequestForwardingRulesRuleActionsForwardGroupConfigServerGroupTuples) GetEndpointGroupId() *string {
	return s.EndpointGroupId
}

func (s *CreateForwardingRulesRequestForwardingRulesRuleActionsForwardGroupConfigServerGroupTuples) SetEndpointGroupId(v string) *CreateForwardingRulesRequestForwardingRulesRuleActionsForwardGroupConfigServerGroupTuples {
	s.EndpointGroupId = &v
	return s
}

func (s *CreateForwardingRulesRequestForwardingRulesRuleActionsForwardGroupConfigServerGroupTuples) Validate() error {
	return dara.Validate(s)
}

type CreateForwardingRulesRequestForwardingRulesRuleConditions struct {
	// The configuration of the domain name.
	//
	// >  We recommend that you use **RuleConditionType*	- and **RuleConditionValue*	- rather than this parameter to configure forwarding conditions.
	HostConfig *CreateForwardingRulesRequestForwardingRulesRuleConditionsHostConfig `json:"HostConfig,omitempty" xml:"HostConfig,omitempty" type:"Struct"`
	// The configuration of the path.
	//
	// >  We recommend that you use **RuleConditionType*	- and **RuleConditionValue*	- rather than this parameter to configure forwarding conditions.
	PathConfig *CreateForwardingRulesRequestForwardingRulesRuleConditionsPathConfig `json:"PathConfig,omitempty" xml:"PathConfig,omitempty" type:"Struct"`
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
	// 	- **SourceIP**: Requests are forwarded based on source IP addresses.
	//
	// example:
	//
	// Host
	RuleConditionType *string `json:"RuleConditionType,omitempty" xml:"RuleConditionType,omitempty"`
	// The value of the forwarding condition. You must specify different JSON strings based on **RuleConditionType**.
	//
	// 	- If **RuleConditionType*	- is set to **Host**, RuleConditionValue specifies a domain name condition. A forwarding rule can contain only one forwarding condition of the Host type. You can specify multiple domain names in a forwarding condition. The relationship between multiple domain names is OR. The domain name must be 3 to 128 characters in length, and can contain letters, digits, hyphens (-), and periods (.). You can use asterisks (\\*) and question marks (?) as wildcard characters. Example: `["www.example.com", "www.aliyun.com"]`.
	//
	// 	- If **RuleConditionType*	- is set to **Path**, RuleConditionValue specifies a path condition. A forwarding rule can contain multiple forwarding conditions of the Path type. The relationship between multiple path conditions is OR. You can specify multiple paths in a forwarding condition. The relationship between multiple paths is OR. The path must be 1 to 128 characters in length, and must start with a forward slash (/). The path can contain letters, digits, and the following special characters: $ - _ . + / & ~ @ : \\". Supported wildcard characters are asterisks (\\*) and question marks (?). Example: `["/a", "/b/"]`.
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

func (s CreateForwardingRulesRequestForwardingRulesRuleConditions) String() string {
	return dara.Prettify(s)
}

func (s CreateForwardingRulesRequestForwardingRulesRuleConditions) GoString() string {
	return s.String()
}

func (s *CreateForwardingRulesRequestForwardingRulesRuleConditions) GetHostConfig() *CreateForwardingRulesRequestForwardingRulesRuleConditionsHostConfig {
	return s.HostConfig
}

func (s *CreateForwardingRulesRequestForwardingRulesRuleConditions) GetPathConfig() *CreateForwardingRulesRequestForwardingRulesRuleConditionsPathConfig {
	return s.PathConfig
}

func (s *CreateForwardingRulesRequestForwardingRulesRuleConditions) GetRuleConditionType() *string {
	return s.RuleConditionType
}

func (s *CreateForwardingRulesRequestForwardingRulesRuleConditions) GetRuleConditionValue() *string {
	return s.RuleConditionValue
}

func (s *CreateForwardingRulesRequestForwardingRulesRuleConditions) SetHostConfig(v *CreateForwardingRulesRequestForwardingRulesRuleConditionsHostConfig) *CreateForwardingRulesRequestForwardingRulesRuleConditions {
	s.HostConfig = v
	return s
}

func (s *CreateForwardingRulesRequestForwardingRulesRuleConditions) SetPathConfig(v *CreateForwardingRulesRequestForwardingRulesRuleConditionsPathConfig) *CreateForwardingRulesRequestForwardingRulesRuleConditions {
	s.PathConfig = v
	return s
}

func (s *CreateForwardingRulesRequestForwardingRulesRuleConditions) SetRuleConditionType(v string) *CreateForwardingRulesRequestForwardingRulesRuleConditions {
	s.RuleConditionType = &v
	return s
}

func (s *CreateForwardingRulesRequestForwardingRulesRuleConditions) SetRuleConditionValue(v string) *CreateForwardingRulesRequestForwardingRulesRuleConditions {
	s.RuleConditionValue = &v
	return s
}

func (s *CreateForwardingRulesRequestForwardingRulesRuleConditions) Validate() error {
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

type CreateForwardingRulesRequestForwardingRulesRuleConditionsHostConfig struct {
	// The domain name.
	//
	// >  We recommend that you use **RuleConditionType*	- and **RuleConditionValue*	- rather than this parameter to configure forwarding conditions.
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s CreateForwardingRulesRequestForwardingRulesRuleConditionsHostConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateForwardingRulesRequestForwardingRulesRuleConditionsHostConfig) GoString() string {
	return s.String()
}

func (s *CreateForwardingRulesRequestForwardingRulesRuleConditionsHostConfig) GetValues() []*string {
	return s.Values
}

func (s *CreateForwardingRulesRequestForwardingRulesRuleConditionsHostConfig) SetValues(v []*string) *CreateForwardingRulesRequestForwardingRulesRuleConditionsHostConfig {
	s.Values = v
	return s
}

func (s *CreateForwardingRulesRequestForwardingRulesRuleConditionsHostConfig) Validate() error {
	return dara.Validate(s)
}

type CreateForwardingRulesRequestForwardingRulesRuleConditionsPathConfig struct {
	// The path.
	//
	// The path must be 1 to 128 characters in length, and must start with a forward slash (/). The path can contain letters, digits, and the following special characters: $ - _ . + / & ~ @ : \\". Supported wildcard characters are asterisks (\\*) and question marks (?).
	//
	// >  We recommend that you use **RuleConditionType*	- and **RuleConditionValue*	- rather than this parameter to configure forwarding conditions.
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s CreateForwardingRulesRequestForwardingRulesRuleConditionsPathConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateForwardingRulesRequestForwardingRulesRuleConditionsPathConfig) GoString() string {
	return s.String()
}

func (s *CreateForwardingRulesRequestForwardingRulesRuleConditionsPathConfig) GetValues() []*string {
	return s.Values
}

func (s *CreateForwardingRulesRequestForwardingRulesRuleConditionsPathConfig) SetValues(v []*string) *CreateForwardingRulesRequestForwardingRulesRuleConditionsPathConfig {
	s.Values = v
	return s
}

func (s *CreateForwardingRulesRequestForwardingRulesRuleConditionsPathConfig) Validate() error {
	return dara.Validate(s)
}
