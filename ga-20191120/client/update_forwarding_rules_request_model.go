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
	// You can generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **RequestId*	- of the request as the **client token**. The **RequestId*	- of each request is different.
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
	// The ID of the listener.
	//
	// This parameter is required.
	//
	// example:
	//
	// lsr-bp1s0vzbi5bxlx5****
	ListenerId *string `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	// The region ID of the Global Accelerator instance. Set the value to **cn-hangzhou**.
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
	// The ID of the forwarding rule.
	//
	// This parameter is required.
	//
	// example:
	//
	// frule-bp1dii16gu9qdvb34****
	ForwardingRuleId *string `json:"ForwardingRuleId,omitempty" xml:"ForwardingRuleId,omitempty"`
	// The name of the forwarding rule.
	//
	// The name must be 2 to 128 characters in length, start with a letter or a Chinese character, and can contain digits, periods (.), underscores (_), and hyphens (-).
	//
	// example:
	//
	// test
	ForwardingRuleName *string `json:"ForwardingRuleName,omitempty" xml:"ForwardingRuleName,omitempty"`
	// The priority of the forwarding rule. A smaller value indicates a higher priority. Valid values: **1*	- to **10000**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1000
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The actions that are performed when the forwarding conditions are met.
	//
	// This parameter is required.
	RuleActions []*UpdateForwardingRulesRequestForwardingRulesRuleActions `json:"RuleActions,omitempty" xml:"RuleActions,omitempty" type:"Repeated"`
	// The forwarding conditions.
	//
	// This parameter is required.
	RuleConditions []*UpdateForwardingRulesRequestForwardingRulesRuleConditions `json:"RuleConditions,omitempty" xml:"RuleConditions,omitempty" type:"Repeated"`
	// The direction of the rule. You do not need to specify this parameter.
	//
	// This parameter is set to **request*	- by default, which indicates that the rule applies to inbound requests.
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
	// > This parameter is deprecated. We recommend that you use the **RuleActionType*	- and **RuleActionValue*	- parameters.
	ForwardGroupConfig *UpdateForwardingRulesRequestForwardingRulesRuleActionsForwardGroupConfig `json:"ForwardGroupConfig,omitempty" xml:"ForwardGroupConfig,omitempty" type:"Struct"`
	// The priority of the action.
	//
	// > This parameter is not in use. You do not need to specify this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20
	Order *int32 `json:"Order,omitempty" xml:"Order,omitempty"`
	// The type of the action. Valid values:
	//
	// - **ForwardGroup**: forwards a request to an endpoint group.
	//
	// - **Redirect**: redirects a request.
	//
	// - **FixResponse**: returns a fixed response.
	//
	// - **Rewrite**: rewrites a request.
	//
	// - **AddHeader**: adds a header to a request.
	//
	// - **RemoveHeader**: removes a header from a request.
	//
	// - **Drop**: drops a request.
	//
	// This parameter is required.
	//
	// example:
	//
	// ForwardGroup
	RuleActionType *string `json:"RuleActionType,omitempty" xml:"RuleActionType,omitempty"`
	// The value of the action.
	//
	// The value is a JSON string that varies based on the value of **RuleActionType**.
	//
	// A forwarding rule can have at most one action of the **ForwardGroup**, **Redirect**, or **FixResponse*	- type. The `Rewrite`, `AddHeader`, and `RemoveHeader` actions must precede a `ForwardGroup` action.
	//
	// - If **RuleActionType*	- is set to **ForwardGroup**, this parameter specifies the endpoint group configuration. You can forward requests to only one endpoint group. Example: `{"type":"endpointgroup", "value":"epg-bp1enpdcrqhl78g6r****"}`, where:
	//
	//   - `type`: Set the value to `endpointgroup`.
	//
	//   - `value`: The ID of the destination endpoint group.
	//
	// - If **RuleActionType*	- is set to **Redirect**, this parameter specifies the redirect configuration. In a **Redirect*	- action, at least one of the following fields must be specified with a non-default value: `protocol`, `domain`, `port`, `path`, or `query`. Example: `{"protocol":"HTTP", "domain":"www.example.com", "port":"80", "path":"/a","query":"value1", "code":"301" }`, where:
	//
	//   - `protocol`: The protocol to which requests are redirected. Valid values: `${protocol}` (default), `HTTP`, and `HTTPS`.
	//
	//   - `domain`: The domain to which requests are redirected. The default value is `${host}`. You can also specify another domain. The domain must be 3 to 128 characters in length and can contain only lowercase letters, digits, and the following special characters: `.-?=~_-+/^*!$&|()[]`.
	//
	//   - `port`: The port to which requests are redirected. The default value is `${port}`. You can also specify another port. Valid values: 1 to 63335.
	//
	//   - `path`: The path to which requests are redirected. The default value is `${path}`. The path must be 1 to 128 characters in length. A regular expression-based path must start with a tilde (\\~) and can contain letters, digits, and the following special characters: `.-_/=?~^*$:()[]+|`. A path that is not a regular expression must start with a forward slash (/) and can contain letters, digits, and the following special characters: `.-_/=?:`.
	//
	//   - `query`: The query string to which requests are redirected. The default value is `${query}`. You can also specify another query string. The query string must be 1 to 128 characters in length and can contain printable ASCII characters whose character codes are in the range of `ch >= 32 && ch < 127`. The letters must be in lowercase. Spaces and the following characters are not allowed: `[]{}<>\\#|&`.
	//
	//   - `code`: The redirect type. Valid values: `301`, `302`, `303`, `307`, and `308`.
	//
	// - If **RuleActionType*	- is set to **FixResponse**, this parameter specifies the fixed response configuration. Example: `{"code":"200", "type":"text/plain", "content":"dssacav" }`, where:
	//
	//   - `code`: The HTTP status code. You can specify a numeric string that represents a `2xx`, `4xx`, or `5xx` status code, where `x` indicates a digit.
	//
	//   - `type`: The content type of the response body. Valid values: **text/plain**, **text/css**, **text/html**, **application/javascript**, and **application/json**.
	//
	//   - `content`: The content of the response body. The content can be up to 1,024 characters in length and cannot contain Chinese characters.
	//
	// - If **RuleActionType*	- is set to **AddHeader**, this parameter specifies the configuration for adding an HTTP header. An **AddHeader*	- action must be used together with a **ForwardGroup*	- action. Example: `[{"name":"header1","type":"user-defined", "value":"value"}]`, where:
	//
	//   - `name`: The name of the HTTP header. The name must be 1 to 40 characters in length and can contain letters, digits, hyphens (-), and underscores (_). The header names specified for **AddHeader*	- must be unique and cannot be the same as those specified for **RemoveHeader**.
	//
	//   - `type`: The content type of the HTTP header. Valid values: `user-defined`, `ref` (reference), and `system-defined`.
	//
	//   - `value`: The content of the HTTP header. This parameter cannot be left empty. If `type` is set to `user-defined`, the content must be 1 to 128 characters in length and can contain printable ASCII characters whose character codes are in the range of `ch >= 32 && ch < 127`. The content can include letters, digits, hyphens (-), and underscores (_). The content cannot start or end with a space. If `type` is set to `ref` (reference), the content must be 1 to 128 characters in length and can contain letters, digits, hyphens (-), and underscores (_). The content cannot start or end with a space. If `type` is set to `system-defined`, the only valid value is `ClientSrcIp`.
	//
	// - If **RuleActionType*	- is set to **RemoveHeader**, this parameter specifies the configuration for removing an HTTP header. A **RemoveHeader*	- action must be used together with a **ForwardGroup*	- action. The header name must be 1 to 40 characters in length and can contain letters, digits, hyphens (-), and underscores (_). Example: `["header1"]`.
	//
	// - If **RuleActionType*	- is set to **Rewrite**, this parameter specifies the rewrite configuration. A **Rewrite*	- action must be used together with a **ForwardGroup*	- action. Example: `{"domain":"value1", "path":"value2", "query":"value3"}`, where:
	//
	//   - `domain`: The domain to which requests are rewritten. The default value is `${host}`. You can also specify another domain. The domain must be 3 to 128 characters in length and can contain only lowercase letters, digits, and the following special characters: `.-?=~_-+/^*!$&|()[]`.
	//
	//   - `path`: The path to which requests are rewritten. The default value is `${path}`. The path must be 1 to 128 characters in length. A regular expression-based path must start with a tilde (\\~) and can contain letters, digits, and the following special characters: `.-_/=?~^*$:()[]+|`. A path that is not a regular expression must start with a forward slash (/) and can contain letters, digits, and the following special characters: `.-_/=?:`.
	//
	//   - `query`: The query string to which requests are rewritten. The default value is `${query}`. You can also specify another query string. The query string must be 1 to 128 characters in length and can contain printable ASCII characters whose character codes are in the range of `ch >= 32 && ch < 127`. The letters must be in lowercase. Spaces and the following characters are not allowed: `[]{}<>\\#|&`.
	//
	// - If **RuleActionType*	- is set to **Drop**, you do not need to specify this parameter.
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
	// The endpoint group configuration.
	//
	// > This parameter is deprecated. We recommend that you use the **RuleActionType*	- and **RuleActionValue*	- parameters.
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
	// > This parameter is deprecated. We recommend that you use the **RuleActionType*	- and **RuleActionValue*	- parameters.
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
	// The domain configuration.
	//
	// > This parameter is deprecated. We recommend that you use the **RuleConditionType*	- and **RuleConditionValue*	- parameters.
	HostConfig *UpdateForwardingRulesRequestForwardingRulesRuleConditionsHostConfig `json:"HostConfig,omitempty" xml:"HostConfig,omitempty" type:"Struct"`
	// The path configuration.
	//
	// > This parameter is deprecated. We recommend that you use the **RuleConditionType*	- and **RuleConditionValue*	- parameters.
	PathConfig *UpdateForwardingRulesRequestForwardingRulesRuleConditionsPathConfig `json:"PathConfig,omitempty" xml:"PathConfig,omitempty" type:"Struct"`
	// The type of the forwarding condition. Valid values:
	//
	// - **Host**: matches a request based on its domain name.
	//
	// - **Path**: matches a request based on its path.
	//
	// - **RequestHeader**: matches a request based on its HTTP header.
	//
	// - **Query**: matches a request based on its query string.
	//
	// - **Method**: matches a request based on its HTTP request method.
	//
	// - **Cookie**: matches a request based on its cookie.
	//
	// - **SourceIP**: matches a request based on its source IP address.
	//
	// This parameter is required.
	//
	// example:
	//
	// Host
	RuleConditionType *string `json:"RuleConditionType,omitempty" xml:"RuleConditionType,omitempty"`
	// The value of the forwarding condition.
	//
	// The value is a JSON string that varies based on the value of **RuleConditionType**.
	//
	// - If **RuleConditionType*	- is set to **Host**, this parameter specifies the domain configuration. A forwarding rule can contain only one Host-based rule condition. The condition can contain multiple domains that are evaluated with a logical OR. A domain must be 3 to 128 characters in length and can contain letters, digits, hyphens (-), and periods (.). You can use asterisks (\\*) and question marks (?) as wildcards. Example: `["www.example.com", "www.aliyun.com"]`.
	//
	// - If **RuleConditionType*	- is set to **Path**, this parameter specifies the path configuration. A forwarding rule can contain multiple path-based rule conditions, which are evaluated with a logical OR. Each condition can contain multiple paths, which are also evaluated with a logical OR. A path must be 1 to 128 characters in length, start with a forward slash (/), and contain only letters, digits, and the following special characters: `$`, `-`, `_`, `.`, `+`, `/`, `&`, `~`, `@`, `:`, and `\\"`. You can use asterisks (\\*) and question marks (?) as wildcards. Example: `["/a", "/b/"]`.
	//
	// - If **RuleConditionType*	- is set to **RequestHeader**, this parameter specifies the HTTP header configuration, which consists of key-value pairs. The values for a specific header must be unique. Example: `[{"header1":["value1","value2"]}]`.
	//
	//   - Key: The key of the HTTP header. The key must be 1 to 40 characters in length and can contain letters, digits, hyphens (-), and underscores (_).
	//
	//   - Value: The value of the HTTP header. The value must be 1 to 128 characters in length and can contain printable ASCII characters whose character codes are in the range of `ch >= 32 && ch < 127`. The value cannot start or end with a space.
	//
	// - If **RuleConditionType*	- is set to **Query**, this parameter specifies the query string configuration, which consists of key-value pairs. Example: `[{"query1":["value1"]}, {"query2":["value2"]}]`.
	//
	//   - Key: The key of the query string. The key must be 1 to 100 characters in length and can contain printable ASCII characters whose character codes are in the range of `ch >= 32 && ch < 127`. The letters must be in lowercase. Spaces and the following characters are not allowed: `[]{}<>\\;/?:@&=+,$%|"^~`.
	//
	//   - Value: The value of the query string. The value must be 1 to 128 characters in length and can contain printable ASCII characters whose character codes are in the range of `ch >= 32 && ch < 127`. The letters must be in lowercase. Spaces and the following characters are not allowed: `[]{}<>\\;/?:@&=+,$%|"^~`.
	//
	// - If **RuleConditionType*	- is set to **Method**, this parameter specifies the HTTP request method configuration. Valid values: **HEAD**, **GET**, **POST**, **OPTIONS**, **PUT**, **PATCH**, and **DELETE**. Example: `["GET", "OPTIONS", "POST"]`.
	//
	// - If **RuleConditionType*	- is set to **Cookie**, this parameter specifies the cookie configuration, which consists of key-value pairs. Example: `[{"cookie1":["value1"]}, {"cookie2":["value2"]}]`
	//
	//   - Key: The key of the cookie. The key must be 1 to 100 characters in length and can contain printable ASCII characters whose character codes are in the range of `ch >= 32 && ch < 127`. The letters must be in lowercase. Spaces and the following characters are not allowed: `#[]{}\\|<>&`.
	//
	//   - Value: The value of the cookie. The value must be 1 to 128 characters in length and can contain printable ASCII characters whose character codes are in the range of `ch >= 32 && ch < 127`. The letters must be in lowercase. Spaces and the following characters are not allowed: `#[]{}\\|<>&`.
	//
	// - If **RuleConditionType*	- is set to **SourceIP**, this parameter specifies the source IP configuration. You can specify IP addresses such as 1.1.XX.XX/32 or CIDR blocks such as 2.2.XX.XX/24. A forwarding rule can contain only one source IP-based rule condition, which can contain multiple source IP addresses or CIDR blocks that are evaluated with a logical OR. Example: `["1.1.XX.XX/32", "2.2.XX.XX/24"]`.
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
	// The domain configuration.
	//
	// > This parameter is deprecated. We recommend that you use the **RuleConditionType*	- and **RuleConditionValue*	- parameters.
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
	// > This parameter is deprecated. We recommend that you use the **RuleConditionType*	- and **RuleConditionValue*	- parameters.
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
