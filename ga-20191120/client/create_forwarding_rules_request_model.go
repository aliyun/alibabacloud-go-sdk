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
	// The ID of the Global Accelerator instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// ga-bp17frjjh0udz4q****
	AcceleratorId *string `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can generate a client token from your client and make sure that the client token is unique among different requests. The client token can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **RequestId*	- of the request as the **ClientToken**. The **RequestId*	- of each request is different.
	//
	// example:
	//
	// 02fb3da4****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The forwarding rule configurations.
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
	// The ID of the region where the Global Accelerator instance is deployed. The only valid value is **cn-hangzhou**.
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
	// The name of the forwarding rule. The name must be 2 to 128 characters long. It must start with a letter or a Chinese character, and can contain letters, Chinese characters, digits, periods (.), underscores (_), and hyphens (-).
	//
	// example:
	//
	// test
	ForwardingRuleName *string `json:"ForwardingRuleName,omitempty" xml:"ForwardingRuleName,omitempty"`
	// The priority of the forwarding rule.
	//
	// Valid values: **1*	- to **10000**. A smaller value indicates a higher priority.
	//
	// example:
	//
	// 1
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The rule actions.
	//
	// This parameter is required.
	RuleActions []*CreateForwardingRulesRequestForwardingRulesRuleActions `json:"RuleActions,omitempty" xml:"RuleActions,omitempty" type:"Repeated"`
	// The rule conditions.
	//
	// This parameter is required.
	RuleConditions []*CreateForwardingRulesRequestForwardingRulesRuleConditions `json:"RuleConditions,omitempty" xml:"RuleConditions,omitempty" type:"Repeated"`
	// The direction in which the rule takes effect. This parameter does not need to be configured.
	//
	// By default, this parameter is set to **request**, which indicates that the rule applies to requests.
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
	// The forwarding configuration.
	//
	// > This parameter is deprecated. We recommend that you use **RuleActionType*	- and **RuleActionValue*	- to configure rule actions.
	ForwardGroupConfig *CreateForwardingRulesRequestForwardingRulesRuleActionsForwardGroupConfig `json:"ForwardGroupConfig,omitempty" xml:"ForwardGroupConfig,omitempty" type:"Struct"`
	// The priority of the action.
	//
	// > This parameter is not in use and can be ignored.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20
	Order *int32 `json:"Order,omitempty" xml:"Order,omitempty"`
	// The type of the rule action. Valid values:
	//
	// - **ForwardGroup**: Forwards requests.
	//
	// - **Redirect**: Redirects requests.
	//
	// - **FixResponse**: Returns a fixed response.
	//
	// - **Rewrite**: Rewrites requests.
	//
	// - **AddHeader**: Adds a header.
	//
	// - **RemoveHeader**: Removes a header.
	//
	// - **Drop**: Drops requests.
	//
	// This parameter is required.
	//
	// example:
	//
	// ForwardGroup
	RuleActionType *string `json:"RuleActionType,omitempty" xml:"RuleActionType,omitempty"`
	// The value for the rule action.
	//
	// This is a JSON-formatted string whose structure depends on the specified **RuleActionType**.
	//
	// A forwarding rule can have at most one action of type **ForwardGroup**, **Redirect**, or **FixResponse**. Actions of type **Rewrite**, **AddHeader**, and **RemoveHeader*	- must be specified before a **ForwardGroup*	- action.
	//
	// - If **RuleActionType*	- is set to **ForwardGroup**, this parameter specifies the endpoint group. You can forward requests to only one endpoint group. Example: `{"type":"endpointgroup", "value":"epg-bp1enpdcrqhl78g6r****"}`, where:
	//
	//   - `type`: Set the value to` endpointgroup`.
	//
	//   - `value`: The ID of the target endpoint group.
	//
	// - If **RuleActionType*	- is set to **Redirect**, this parameter specifies the redirect configuration. At least one of the `protocol`, `domain`, `port`, `path`, or `query` fields must be set to a value other than its default. Example: `{"protocol":"HTTP", "domain":"www.example.com", "port":"80", "path":"/a","query":"value1", "code":"301" }`, where:
	//
	//   - `protocol`: The protocol for the redirect. Valid values: `${protocol}` (default), `HTTP`, and `HTTPS`.
	//
	//   - `domain`: The domain name for the redirect. The default value is `${host}`. You can also specify another domain name. A domain name must be 3 to 128 characters long and can contain only lowercase letters, digits, and the following special characters:` .-=~_-+/^*!$&()[]?`.
	//
	//   - `port`: The port for the redirect. The default value is `${port}`. You can also specify a port number. Valid values: 1 to 63335.
	//
	//   - `path`: The path for the redirect. The default value is `${path}`. The path must be 1 to 128 characters long. For a regular expression path, it must start with a tilde (\\~) and can contain uppercase and lowercase letters, digits, and the following special characters:` .-_/=?~^*$:()[]+`. For a non-regular expression path, it must start with a forward slash (/) and can contain uppercase and lowercase letters, digits, and the following special characters:` .-_/=:?`.
	//
	//   - `query`: The query string for the redirect. The default value is `${query}`. You can also specify another query string. The query string must be 1 to 128 characters long and contain only printable characters within the ASCII range of` ch >= 32 && ch < 127`. Letters must be lowercase. Spaces and the following special characters are not supported:` []{}<>\\#&`.
	//
	//   - `code`: The redirect code. Valid values: `301`, `302`, `303`, `307`, and `308`.
	//
	// - If **RuleActionType*	- is set to **FixResponse**, this parameter specifies the fixed response configuration. Example: `{"code":"200", "type":"text/plain", "content":"dssacav" }`, where:
	//
	//   - `code`: The response status code. The value must be a numeric string in the `2xx`, `4xx`, or `5xx` format, where `x` is any digit.
	//
	//   - `type`: The content type of the response body. Valid values: **text/plain**, **text/css**, **text/html**, **application/javascript**, and **application/json**.
	//
	//   - `content`: The content of the response body. The content cannot exceed 1,000 characters and does not support Chinese characters.
	//
	// - If **RuleActionType*	- is set to **AddHeader**, this parameter specifies the configuration for adding an HTTP header. If a forwarding rule contains an **AddHeader*	- action, it must also contain a **ForwardGroup*	- action. Example: `[{"name":"header1","type":"user-defined", "value":"value"}]`, where:
	//
	//   - `name`: The name of the HTTP header. The name must be 1 to 40 characters long and can contain uppercase and lowercase letters, digits, hyphens (-), and underscores (_). The header names in **AddHeader*	- actions must be unique and cannot be the same as any header name in a **RemoveHeader*	- action.
	//
	//   - `type`: The type of the header value. Valid values: `user-defined`, `ref` (reference), and `system-defined`.
	//
	//   - `value`: The content of the HTTP header. This field cannot be empty. If `type` is `user-defined`, the value must be 1 to 128 characters long and contain only printable characters within the ASCII range of `ch >= 32 && ch < 127`. The value can contain uppercase and lowercase letters, digits, hyphens (-), and underscores (_), and cannot start or end with a space. If `type` is `ref`, the value must be 1 to 128 characters long and can contain uppercase and lowercase letters, digits, hyphens (-), and underscores (_). The value cannot start or end with a space. If `type` is `system-defined`, the only valid value is `ClientSrcIp`.
	//
	// - If **RuleActionType*	- is set to **RemoveHeader**, this parameter specifies the HTTP headers to remove. If a forwarding rule contains a **RemoveHeader*	- action, it must also contain a **ForwardGroup*	- action. The value must be 1 to 40 characters long and can contain uppercase and lowercase letters, digits, hyphens (-), and underscores (_). Example: `["header1"]`.
	//
	// - If **RuleActionType*	- is set to **Rewrite**, this parameter specifies the rewrite configuration. If a forwarding rule contains a **Rewrite*	- action, it must also contain a **ForwardGroup*	- action. Example: `{"domain":"value1", "path":"value2", "query":"value3"}`, where:
	//
	//   - `domain`: The domain name to rewrite. The default value is `${host}`. You can also specify another domain name. A domain name must be 3 to 128 characters long and can contain only lowercase letters, digits, and the following special characters:` .-=~_-+/^*!$&()[]?`.
	//
	//   - `path`: The path to rewrite. The default value is `${path}`. The path must be 1 to 128 characters long. For a regular expression path, it must start with a tilde (\\~) and can contain uppercase and lowercase letters, digits, and the following special characters:` .-_/=?~^*$:()[]+`. For a non-regular expression path, it must start with a forward slash (/) and can contain uppercase and lowercase letters, digits, and the following special characters:` .-_/=:?`.
	//
	//   - `query`: The query string to rewrite. The default value is `${query}`. You can also specify another query string. The query string must be 1 to 128 characters long and contain only printable characters within the ASCII range of` ch >= 32 && ch < 127`. Letters must be lowercase. Spaces and the following special characters are not supported:` []{}<>\\#&`.
	//
	// - If **RuleActionType*	- is set to **Drop**, you do not need to specify a value for this parameter.
	//
	// example:
	//
	// [{"type":"endpointgroup","value":"epg-bp1l49ltx6iengvf2ks5z****"}]
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
	// The endpoint group configuration.
	//
	// > This parameter is deprecated. We recommend that you use **RuleActionType*	- and **RuleActionValue*	- to configure rule actions.
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
	// > This parameter is deprecated. We recommend that you use **RuleActionType*	- and **RuleActionValue*	- to configure rule actions.
	//
	// This parameter is required.
	//
	// example:
	//
	// epg-bp1nktp3qgbcq9ih6****
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
	// The domain name configuration.
	//
	// > This parameter is deprecated. We recommend that you use **RuleConditionType*	- and **RuleConditionValue*	- to configure rule conditions.
	HostConfig *CreateForwardingRulesRequestForwardingRulesRuleConditionsHostConfig `json:"HostConfig,omitempty" xml:"HostConfig,omitempty" type:"Struct"`
	// The path configuration.
	//
	// > This parameter is deprecated. We recommend that you use **RuleConditionType*	- and **RuleConditionValue*	- to configure rule conditions.
	PathConfig *CreateForwardingRulesRequestForwardingRulesRuleConditionsPathConfig `json:"PathConfig,omitempty" xml:"PathConfig,omitempty" type:"Struct"`
	// The type of the rule condition. Valid values:
	//
	// - **Host**: Matches requests by domain name.
	//
	// - **Path**: Matches requests by path.
	//
	// - **RequestHeader**: Matches requests by HTTP header.
	//
	// - **Query**: Matches requests by query string.
	//
	// - **Method**: Matches requests by HTTP method.
	//
	// - **Cookie**: Matches requests by cookie.
	//
	// - **SourceIP**: Matches requests by source IP address.
	//
	// example:
	//
	// Host
	RuleConditionType *string `json:"RuleConditionType,omitempty" xml:"RuleConditionType,omitempty"`
	// The value of the rule condition.
	//
	// This is a JSON-formatted string whose structure depends on the specified **RuleConditionType**.
	//
	// - If **RuleConditionType*	- is set to **Host**, this parameter specifies the domain name conditions. A forwarding rule can have only one **Host*	- rule condition. This rule condition can contain multiple domain names, which are evaluated with a logical OR. A domain name must be 3 to 128 characters long and can contain letters, digits, hyphens (-), and periods (.). You can use asterisks (\\*) and question marks (?) as wildcards. Example: `["www.example.com", "www.aliyun.com"]`.
	//
	// - If **RuleConditionType*	- is set to **Path**, this parameter specifies the path conditions. A forwarding rule can have multiple **Path*	- rule conditions, which are evaluated with a logical OR. Each path rule condition can contain multiple paths, which are also evaluated with a logical OR. A path must be 1 to 128 characters long and must start with a forward slash (/). It can contain letters, digits, dollar signs ($), hyphens (-), underscores (_), periods (.), plus signs (+), forward slashes (/), ampersands (&), tildes (\\~), at signs (@), colons (:), and apostrophes (\\"). You can use asterisks (\\*) and question marks (?) as wildcards. Example: `["/a", "/b/"]`.
	//
	// - If **RuleConditionType*	- is set to **RequestHeader**, this parameter specifies the HTTP header conditions. The value is a key-value pair. The header values within the same rule condition must be unique. Example: `[{"header1":["value1","value2"]}]`.
	//
	//   - Key: The HTTP header key must be 1 to 40 characters long and can contain letters, digits, hyphens (-), and underscores (_).
	//
	//   - Value: The HTTP header value must be 1 to 128 characters long and contain only printable characters within the ASCII range of` ch >= 32 && ch < 127`. The value cannot start or end with a space.
	//
	// - If **RuleConditionType*	- is set to **Query**, this parameter specifies the query string conditions. The value is a key-value pair. Example: `[{"query1":["value1"]}, {"query2":["value2"]}]`.
	//
	//   - Key: The key must be 1 to 100 characters long and contain only printable characters within the ASCII range of` ch >= 32 && ch < 127`. Letters must be lowercase. Spaces and the following characters are not supported:` []{}<>\\;/?:@&=+,$%"^~`.
	//
	//   - Value: The value must be 1 to 128 characters long and contain only printable characters within the ASCII range of` ch >= 32 && ch < 127`. Letters must be lowercase. Spaces and the following characters are not supported:` []{}<>\\;/?:@&=+,$%"^~`.
	//
	// - If **RuleConditionType*	- is set to **Method**, this parameter specifies the HTTP method conditions. Valid values: **HEAD**, **GET**, **POST**, **OPTIONS**, **PUT**, **PATCH**, and **DELETE**. Example: `["GET", "OPTIONS", "POST"]`.
	//
	// - If **RuleConditionType*	- is set to **Cookie**, this parameter specifies the cookie conditions. The value is a key-value pair. Example: `[{"cookie1":["value1"]}, {"cookie2":["value2"]}]`
	//
	//   - Key: The cookie key must be 1 to 100 characters long and contain only printable characters within the ASCII range of` ch >= 32 && ch < 127`. Letters must be lowercase. Spaces and the following characters are not supported:` #[]{}\\<>&`.
	//
	//   - Value: The cookie value must be 1 to 128 characters long and contain only printable characters within the ASCII range of` ch >= 32 && ch < 127`. Letters must be lowercase. Spaces and the following characters are not supported:` #[]{}\\<>&`.
	//
	// - If **RuleConditionType*	- is set to **SourceIP**, this parameter specifies the source IP conditions. You can specify IP addresses, for example, `1.1.XX.XX/32`, or CIDR blocks, for example, `2.2.XX.XX/24`. A forwarding rule can have only one **SourceIP*	- rule condition. This rule condition can contain multiple source IP addresses, which are evaluated with a logical OR. Example: `["1.1.XX.XX/32", "2.2.XX.XX/24"]`.
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
	// The domain name configuration.
	//
	// > This parameter is deprecated. We recommend that you use **RuleConditionType*	- and **RuleConditionValue*	- to configure rule conditions.
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
	// The path configuration.
	//
	// A path must be 1 to 128 characters long and must start with a forward slash (/). It can contain letters, digits, dollar signs ($), hyphens (-), underscores (_), periods (.), plus signs (+), forward slashes (/), ampersands (&), tildes (\\~), at signs (@), colons (:), and apostrophes (\\"). You can use asterisks (\\*) and question marks (?) as wildcards.
	//
	// > This parameter is deprecated. We recommend that you use **RuleConditionType*	- and **RuleConditionValue*	- to configure rule conditions.
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
