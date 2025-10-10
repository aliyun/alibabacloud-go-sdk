// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRuleAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *UpdateRuleAttributeRequest
	GetClientToken() *string
	SetDryRun(v bool) *UpdateRuleAttributeRequest
	GetDryRun() *bool
	SetPriority(v int32) *UpdateRuleAttributeRequest
	GetPriority() *int32
	SetRuleActions(v []*UpdateRuleAttributeRequestRuleActions) *UpdateRuleAttributeRequest
	GetRuleActions() []*UpdateRuleAttributeRequestRuleActions
	SetRuleConditions(v []*UpdateRuleAttributeRequestRuleConditions) *UpdateRuleAttributeRequest
	GetRuleConditions() []*UpdateRuleAttributeRequestRuleConditions
	SetRuleId(v string) *UpdateRuleAttributeRequest
	GetRuleId() *string
	SetRuleName(v string) *UpdateRuleAttributeRequest
	GetRuleName() *string
}

type UpdateRuleAttributeRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 5A2CFF0E-5718-45B5-9D4D-70B3FF3898
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// 	- **true**: performs a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error code is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// 	- **false*	- (default): performs a dry run and performs the actual request. If the request passes the dry run, a `2xx HTTP` status code is returned and the operation is performed.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The priority of the forwarding rule. Valid values: **1 to 10000**. A lower value specifies a higher priority.
	//
	// > The priorities of the forwarding rules created for the same listener must be unique.
	//
	// example:
	//
	// 10
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The actions of the forwarding rule.
	RuleActions []*UpdateRuleAttributeRequestRuleActions `json:"RuleActions,omitempty" xml:"RuleActions,omitempty" type:"Repeated"`
	// The match conditions of the forwarding rule.
	RuleConditions []*UpdateRuleAttributeRequestRuleConditions `json:"RuleConditions,omitempty" xml:"RuleConditions,omitempty" type:"Repeated"`
	// The ID of the forwarding rule.
	//
	// This parameter is required.
	//
	// example:
	//
	// rule-4dp5i6ea****
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The name of the forwarding rule. The name must be 2 to 128 characters in length, and can contain letters, digits, periods (.), underscores (_), and hyphens (-). The name must start with a letter.
	//
	// example:
	//
	// rule-instance-test
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
}

func (s UpdateRuleAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateRuleAttributeRequest) GoString() string {
	return s.String()
}

func (s *UpdateRuleAttributeRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateRuleAttributeRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *UpdateRuleAttributeRequest) GetPriority() *int32 {
	return s.Priority
}

func (s *UpdateRuleAttributeRequest) GetRuleActions() []*UpdateRuleAttributeRequestRuleActions {
	return s.RuleActions
}

func (s *UpdateRuleAttributeRequest) GetRuleConditions() []*UpdateRuleAttributeRequestRuleConditions {
	return s.RuleConditions
}

func (s *UpdateRuleAttributeRequest) GetRuleId() *string {
	return s.RuleId
}

func (s *UpdateRuleAttributeRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *UpdateRuleAttributeRequest) SetClientToken(v string) *UpdateRuleAttributeRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateRuleAttributeRequest) SetDryRun(v bool) *UpdateRuleAttributeRequest {
	s.DryRun = &v
	return s
}

func (s *UpdateRuleAttributeRequest) SetPriority(v int32) *UpdateRuleAttributeRequest {
	s.Priority = &v
	return s
}

func (s *UpdateRuleAttributeRequest) SetRuleActions(v []*UpdateRuleAttributeRequestRuleActions) *UpdateRuleAttributeRequest {
	s.RuleActions = v
	return s
}

func (s *UpdateRuleAttributeRequest) SetRuleConditions(v []*UpdateRuleAttributeRequestRuleConditions) *UpdateRuleAttributeRequest {
	s.RuleConditions = v
	return s
}

func (s *UpdateRuleAttributeRequest) SetRuleId(v string) *UpdateRuleAttributeRequest {
	s.RuleId = &v
	return s
}

func (s *UpdateRuleAttributeRequest) SetRuleName(v string) *UpdateRuleAttributeRequest {
	s.RuleName = &v
	return s
}

func (s *UpdateRuleAttributeRequest) Validate() error {
	return dara.Validate(s)
}

type UpdateRuleAttributeRequestRuleActions struct {
	// The CORS configuration.
	CorsConfig *UpdateRuleAttributeRequestRuleActionsCorsConfig `json:"CorsConfig,omitempty" xml:"CorsConfig,omitempty" type:"Struct"`
	// The configuration of the custom response.
	FixedResponseConfig *UpdateRuleAttributeRequestRuleActionsFixedResponseConfig `json:"FixedResponseConfig,omitempty" xml:"FixedResponseConfig,omitempty" type:"Struct"`
	// The configuration of the server groups.
	ForwardGroupConfig *UpdateRuleAttributeRequestRuleActionsForwardGroupConfig `json:"ForwardGroupConfig,omitempty" xml:"ForwardGroupConfig,omitempty" type:"Struct"`
	// The configuration of the header to be inserted.
	InsertHeaderConfig *UpdateRuleAttributeRequestRuleActionsInsertHeaderConfig `json:"InsertHeaderConfig,omitempty" xml:"InsertHeaderConfig,omitempty" type:"Struct"`
	// The priority of the action. Valid values: **1 to 50000**. A smaller value specifies a higher priority. The actions of a forwarding rule are applied in descending order of priority. This parameter cannot be left empty. The priority of each action within a forwarding rule must be unique. You can specify up to 20 forwarding rule priorities.
	//
	// example:
	//
	// 1
	Order *int32 `json:"Order,omitempty" xml:"Order,omitempty"`
	// The configuration of the redirect action. You can specify up to 20 redirect actions.
	RedirectConfig *UpdateRuleAttributeRequestRuleActionsRedirectConfig `json:"RedirectConfig,omitempty" xml:"RedirectConfig,omitempty" type:"Struct"`
	// The HTTP header to be removed.
	RemoveHeaderConfig *UpdateRuleAttributeRequestRuleActionsRemoveHeaderConfig `json:"RemoveHeaderConfig,omitempty" xml:"RemoveHeaderConfig,omitempty" type:"Struct"`
	// The configuration of the rewrite action.
	RewriteConfig *UpdateRuleAttributeRequestRuleActionsRewriteConfig `json:"RewriteConfig,omitempty" xml:"RewriteConfig,omitempty" type:"Struct"`
	// The configuration of the action to throttle traffic.
	TrafficLimitConfig *UpdateRuleAttributeRequestRuleActionsTrafficLimitConfig `json:"TrafficLimitConfig,omitempty" xml:"TrafficLimitConfig,omitempty" type:"Struct"`
	// The configuration of the traffic mirroring action.
	TrafficMirrorConfig *UpdateRuleAttributeRequestRuleActionsTrafficMirrorConfig `json:"TrafficMirrorConfig,omitempty" xml:"TrafficMirrorConfig,omitempty" type:"Struct"`
	// The type of the task. You can specify up to 11 types of action. Valid values:
	//
	// 	- **ForwardGroup**: forwards a request to multiple vServer groups.
	//
	// 	- **Redirect**: redirects requests.
	//
	// 	- **FixedResponse**: returns a fixed response.
	//
	// 	- **Rewrite**: rewrites requests.
	//
	// 	- **InsertHeader**: inserts a header.
	//
	// 	- **RemoveHeader**: deletes the header of a request.
	//
	// 	- **TrafficLimit**: throttles traffic.
	//
	// 	- **trafficMirror**: mirrors network traffic.
	//
	// 	- **Cors**: forwards requests based on CORS.
	//
	// The preceding actions can be classified into two types:
	//
	// 	- **FinalType**: Each forwarding rule can contain only one FinalType action, which is performed at the end. You can specify only one of **ForwardGroup**, **Redirect**, and **FixedResponse**.
	//
	// 	- **ExtType**: Each forwarding rule can contain one or more **ExtType*	- actions, which are performed before the **FinalType*	- action. If you want to specify an ExtType action, you must also specify a **FinalType*	- action. You can specify multiple **InsertHeader*	- actions or one **Rewrite*	- action.
	//
	// example:
	//
	// ForwardGroup
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s UpdateRuleAttributeRequestRuleActions) String() string {
	return dara.Prettify(s)
}

func (s UpdateRuleAttributeRequestRuleActions) GoString() string {
	return s.String()
}

func (s *UpdateRuleAttributeRequestRuleActions) GetCorsConfig() *UpdateRuleAttributeRequestRuleActionsCorsConfig {
	return s.CorsConfig
}

func (s *UpdateRuleAttributeRequestRuleActions) GetFixedResponseConfig() *UpdateRuleAttributeRequestRuleActionsFixedResponseConfig {
	return s.FixedResponseConfig
}

func (s *UpdateRuleAttributeRequestRuleActions) GetForwardGroupConfig() *UpdateRuleAttributeRequestRuleActionsForwardGroupConfig {
	return s.ForwardGroupConfig
}

func (s *UpdateRuleAttributeRequestRuleActions) GetInsertHeaderConfig() *UpdateRuleAttributeRequestRuleActionsInsertHeaderConfig {
	return s.InsertHeaderConfig
}

func (s *UpdateRuleAttributeRequestRuleActions) GetOrder() *int32 {
	return s.Order
}

func (s *UpdateRuleAttributeRequestRuleActions) GetRedirectConfig() *UpdateRuleAttributeRequestRuleActionsRedirectConfig {
	return s.RedirectConfig
}

func (s *UpdateRuleAttributeRequestRuleActions) GetRemoveHeaderConfig() *UpdateRuleAttributeRequestRuleActionsRemoveHeaderConfig {
	return s.RemoveHeaderConfig
}

func (s *UpdateRuleAttributeRequestRuleActions) GetRewriteConfig() *UpdateRuleAttributeRequestRuleActionsRewriteConfig {
	return s.RewriteConfig
}

func (s *UpdateRuleAttributeRequestRuleActions) GetTrafficLimitConfig() *UpdateRuleAttributeRequestRuleActionsTrafficLimitConfig {
	return s.TrafficLimitConfig
}

func (s *UpdateRuleAttributeRequestRuleActions) GetTrafficMirrorConfig() *UpdateRuleAttributeRequestRuleActionsTrafficMirrorConfig {
	return s.TrafficMirrorConfig
}

func (s *UpdateRuleAttributeRequestRuleActions) GetType() *string {
	return s.Type
}

func (s *UpdateRuleAttributeRequestRuleActions) SetCorsConfig(v *UpdateRuleAttributeRequestRuleActionsCorsConfig) *UpdateRuleAttributeRequestRuleActions {
	s.CorsConfig = v
	return s
}

func (s *UpdateRuleAttributeRequestRuleActions) SetFixedResponseConfig(v *UpdateRuleAttributeRequestRuleActionsFixedResponseConfig) *UpdateRuleAttributeRequestRuleActions {
	s.FixedResponseConfig = v
	return s
}

func (s *UpdateRuleAttributeRequestRuleActions) SetForwardGroupConfig(v *UpdateRuleAttributeRequestRuleActionsForwardGroupConfig) *UpdateRuleAttributeRequestRuleActions {
	s.ForwardGroupConfig = v
	return s
}

func (s *UpdateRuleAttributeRequestRuleActions) SetInsertHeaderConfig(v *UpdateRuleAttributeRequestRuleActionsInsertHeaderConfig) *UpdateRuleAttributeRequestRuleActions {
	s.InsertHeaderConfig = v
	return s
}

func (s *UpdateRuleAttributeRequestRuleActions) SetOrder(v int32) *UpdateRuleAttributeRequestRuleActions {
	s.Order = &v
	return s
}

func (s *UpdateRuleAttributeRequestRuleActions) SetRedirectConfig(v *UpdateRuleAttributeRequestRuleActionsRedirectConfig) *UpdateRuleAttributeRequestRuleActions {
	s.RedirectConfig = v
	return s
}

func (s *UpdateRuleAttributeRequestRuleActions) SetRemoveHeaderConfig(v *UpdateRuleAttributeRequestRuleActionsRemoveHeaderConfig) *UpdateRuleAttributeRequestRuleActions {
	s.RemoveHeaderConfig = v
	return s
}

func (s *UpdateRuleAttributeRequestRuleActions) SetRewriteConfig(v *UpdateRuleAttributeRequestRuleActionsRewriteConfig) *UpdateRuleAttributeRequestRuleActions {
	s.RewriteConfig = v
	return s
}

func (s *UpdateRuleAttributeRequestRuleActions) SetTrafficLimitConfig(v *UpdateRuleAttributeRequestRuleActionsTrafficLimitConfig) *UpdateRuleAttributeRequestRuleActions {
	s.TrafficLimitConfig = v
	return s
}

func (s *UpdateRuleAttributeRequestRuleActions) SetTrafficMirrorConfig(v *UpdateRuleAttributeRequestRuleActionsTrafficMirrorConfig) *UpdateRuleAttributeRequestRuleActions {
	s.TrafficMirrorConfig = v
	return s
}

func (s *UpdateRuleAttributeRequestRuleActions) SetType(v string) *UpdateRuleAttributeRequestRuleActions {
	s.Type = &v
	return s
}

func (s *UpdateRuleAttributeRequestRuleActions) Validate() error {
	return dara.Validate(s)
}

type UpdateRuleAttributeRequestRuleActionsCorsConfig struct {
	// Specifies whether credentials can be carried in CORS requests. Valid values:
	//
	// 	- **on**
	//
	// 	- **off**
	//
	// example:
	//
	// on
	AllowCredentials *string `json:"AllowCredentials,omitempty" xml:"AllowCredentials,omitempty"`
	// The trusted headers of CORS requests.
	AllowHeaders []*string `json:"AllowHeaders,omitempty" xml:"AllowHeaders,omitempty" type:"Repeated"`
	// The trusted HTTP methods of CORS requests.
	AllowMethods []*string `json:"AllowMethods,omitempty" xml:"AllowMethods,omitempty" type:"Repeated"`
	// The trusted origins. You can specify one or more values, or only an asterisk (`*`).
	//
	// 	- The value must start with `http://` or `https://`, and be followed by a valid domain name, including top-level wildcard domain names. Example: `http://*.test.abc.example.com`.
	//
	// 	- You can specify ports for a single value. Valid values: **1*	- to **65535**.
	AllowOrigin []*string `json:"AllowOrigin,omitempty" xml:"AllowOrigin,omitempty" type:"Repeated"`
	// The headers that can be exposed.
	ExposeHeaders []*string `json:"ExposeHeaders,omitempty" xml:"ExposeHeaders,omitempty" type:"Repeated"`
	// The maximum cache time of dry runs in the browser. Unit: seconds.
	//
	// Valid values: **-1*	- to **172800**.
	//
	// example:
	//
	// 1000
	MaxAge *int64 `json:"MaxAge,omitempty" xml:"MaxAge,omitempty"`
}

func (s UpdateRuleAttributeRequestRuleActionsCorsConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateRuleAttributeRequestRuleActionsCorsConfig) GoString() string {
	return s.String()
}

func (s *UpdateRuleAttributeRequestRuleActionsCorsConfig) GetAllowCredentials() *string {
	return s.AllowCredentials
}

func (s *UpdateRuleAttributeRequestRuleActionsCorsConfig) GetAllowHeaders() []*string {
	return s.AllowHeaders
}

func (s *UpdateRuleAttributeRequestRuleActionsCorsConfig) GetAllowMethods() []*string {
	return s.AllowMethods
}

func (s *UpdateRuleAttributeRequestRuleActionsCorsConfig) GetAllowOrigin() []*string {
	return s.AllowOrigin
}

func (s *UpdateRuleAttributeRequestRuleActionsCorsConfig) GetExposeHeaders() []*string {
	return s.ExposeHeaders
}

func (s *UpdateRuleAttributeRequestRuleActionsCorsConfig) GetMaxAge() *int64 {
	return s.MaxAge
}

func (s *UpdateRuleAttributeRequestRuleActionsCorsConfig) SetAllowCredentials(v string) *UpdateRuleAttributeRequestRuleActionsCorsConfig {
	s.AllowCredentials = &v
	return s
}

func (s *UpdateRuleAttributeRequestRuleActionsCorsConfig) SetAllowHeaders(v []*string) *UpdateRuleAttributeRequestRuleActionsCorsConfig {
	s.AllowHeaders = v
	return s
}

func (s *UpdateRuleAttributeRequestRuleActionsCorsConfig) SetAllowMethods(v []*string) *UpdateRuleAttributeRequestRuleActionsCorsConfig {
	s.AllowMethods = v
	return s
}

func (s *UpdateRuleAttributeRequestRuleActionsCorsConfig) SetAllowOrigin(v []*string) *UpdateRuleAttributeRequestRuleActionsCorsConfig {
	s.AllowOrigin = v
	return s
}

func (s *UpdateRuleAttributeRequestRuleActionsCorsConfig) SetExposeHeaders(v []*string) *UpdateRuleAttributeRequestRuleActionsCorsConfig {
	s.ExposeHeaders = v
	return s
}

func (s *UpdateRuleAttributeRequestRuleActionsCorsConfig) SetMaxAge(v int64) *UpdateRuleAttributeRequestRuleActionsCorsConfig {
	s.MaxAge = &v
	return s
}

func (s *UpdateRuleAttributeRequestRuleActionsCorsConfig) Validate() error {
	return dara.Validate(s)
}

type UpdateRuleAttributeRequestRuleActionsFixedResponseConfig struct {
	// The content of the response. The content can be up to 1 KB in size, and can contain only ASCII characters.
	//
	// example:
	//
	// dssacav
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The content type.
	//
	// Valid values: **text/plain**, **text/css**, **text/html**, **application/javascript**, and **application/json**.
	//
	// example:
	//
	// text/plain
	ContentType *string `json:"ContentType,omitempty" xml:"ContentType,omitempty"`
	// The HTTP status code in responses. Valid values: **2xx**, **4xx**, **5xx**. The value must be a numeric string. **x*	- must be a digit.
	//
	// example:
	//
	// HTTP_200
	HttpCode *string `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
}

func (s UpdateRuleAttributeRequestRuleActionsFixedResponseConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateRuleAttributeRequestRuleActionsFixedResponseConfig) GoString() string {
	return s.String()
}

func (s *UpdateRuleAttributeRequestRuleActionsFixedResponseConfig) GetContent() *string {
	return s.Content
}

func (s *UpdateRuleAttributeRequestRuleActionsFixedResponseConfig) GetContentType() *string {
	return s.ContentType
}

func (s *UpdateRuleAttributeRequestRuleActionsFixedResponseConfig) GetHttpCode() *string {
	return s.HttpCode
}

func (s *UpdateRuleAttributeRequestRuleActionsFixedResponseConfig) SetContent(v string) *UpdateRuleAttributeRequestRuleActionsFixedResponseConfig {
	s.Content = &v
	return s
}

func (s *UpdateRuleAttributeRequestRuleActionsFixedResponseConfig) SetContentType(v string) *UpdateRuleAttributeRequestRuleActionsFixedResponseConfig {
	s.ContentType = &v
	return s
}

func (s *UpdateRuleAttributeRequestRuleActionsFixedResponseConfig) SetHttpCode(v string) *UpdateRuleAttributeRequestRuleActionsFixedResponseConfig {
	s.HttpCode = &v
	return s
}

func (s *UpdateRuleAttributeRequestRuleActionsFixedResponseConfig) Validate() error {
	return dara.Validate(s)
}

type UpdateRuleAttributeRequestRuleActionsForwardGroupConfig struct {
	// The configuration of session persistence for server groups.
	ServerGroupStickySession *UpdateRuleAttributeRequestRuleActionsForwardGroupConfigServerGroupStickySession `json:"ServerGroupStickySession,omitempty" xml:"ServerGroupStickySession,omitempty" type:"Struct"`
	// The server groups to which requests are forwarded.
	ServerGroupTuples []*UpdateRuleAttributeRequestRuleActionsForwardGroupConfigServerGroupTuples `json:"ServerGroupTuples,omitempty" xml:"ServerGroupTuples,omitempty" type:"Repeated"`
}

func (s UpdateRuleAttributeRequestRuleActionsForwardGroupConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateRuleAttributeRequestRuleActionsForwardGroupConfig) GoString() string {
	return s.String()
}

func (s *UpdateRuleAttributeRequestRuleActionsForwardGroupConfig) GetServerGroupStickySession() *UpdateRuleAttributeRequestRuleActionsForwardGroupConfigServerGroupStickySession {
	return s.ServerGroupStickySession
}

func (s *UpdateRuleAttributeRequestRuleActionsForwardGroupConfig) GetServerGroupTuples() []*UpdateRuleAttributeRequestRuleActionsForwardGroupConfigServerGroupTuples {
	return s.ServerGroupTuples
}

func (s *UpdateRuleAttributeRequestRuleActionsForwardGroupConfig) SetServerGroupStickySession(v *UpdateRuleAttributeRequestRuleActionsForwardGroupConfigServerGroupStickySession) *UpdateRuleAttributeRequestRuleActionsForwardGroupConfig {
	s.ServerGroupStickySession = v
	return s
}

func (s *UpdateRuleAttributeRequestRuleActionsForwardGroupConfig) SetServerGroupTuples(v []*UpdateRuleAttributeRequestRuleActionsForwardGroupConfigServerGroupTuples) *UpdateRuleAttributeRequestRuleActionsForwardGroupConfig {
	s.ServerGroupTuples = v
	return s
}

func (s *UpdateRuleAttributeRequestRuleActionsForwardGroupConfig) Validate() error {
	return dara.Validate(s)
}

type UpdateRuleAttributeRequestRuleActionsForwardGroupConfigServerGroupStickySession struct {
	// Specifies whether to enable session persistence. Valid values:
	//
	// 	- **true**
	//
	// 	- **false*	- (default)
	//
	// example:
	//
	// false
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The timeout period for sessions. Unit: seconds. Valid values: 1 to 86400.
	//
	// example:
	//
	// 2
	Timeout *int32 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
}

func (s UpdateRuleAttributeRequestRuleActionsForwardGroupConfigServerGroupStickySession) String() string {
	return dara.Prettify(s)
}

func (s UpdateRuleAttributeRequestRuleActionsForwardGroupConfigServerGroupStickySession) GoString() string {
	return s.String()
}

func (s *UpdateRuleAttributeRequestRuleActionsForwardGroupConfigServerGroupStickySession) GetEnabled() *bool {
	return s.Enabled
}

func (s *UpdateRuleAttributeRequestRuleActionsForwardGroupConfigServerGroupStickySession) GetTimeout() *int32 {
	return s.Timeout
}

func (s *UpdateRuleAttributeRequestRuleActionsForwardGroupConfigServerGroupStickySession) SetEnabled(v bool) *UpdateRuleAttributeRequestRuleActionsForwardGroupConfigServerGroupStickySession {
	s.Enabled = &v
	return s
}

func (s *UpdateRuleAttributeRequestRuleActionsForwardGroupConfigServerGroupStickySession) SetTimeout(v int32) *UpdateRuleAttributeRequestRuleActionsForwardGroupConfigServerGroupStickySession {
	s.Timeout = &v
	return s
}

func (s *UpdateRuleAttributeRequestRuleActionsForwardGroupConfigServerGroupStickySession) Validate() error {
	return dara.Validate(s)
}

type UpdateRuleAttributeRequestRuleActionsForwardGroupConfigServerGroupTuples struct {
	// The ID of the server group to which requests are forwarded.
	//
	// example:
	//
	// sg--atstuj3rtoptyui****
	ServerGroupId *string `json:"ServerGroupId,omitempty" xml:"ServerGroupId,omitempty"`
	// The weight of the server group. A larger value specifies a higher weight. A server group with a higher weight receives more requests. Valid values: **0*	- to **100**.
	//
	// 	- If the number of destination server groups is 1, the default weight of the server group is **100**, unless you specify a weight.
	//
	// 	- If the number of destination server groups is larger than 1, you must specify a weight.
	//
	// example:
	//
	// 30
	Weight *int32 `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s UpdateRuleAttributeRequestRuleActionsForwardGroupConfigServerGroupTuples) String() string {
	return dara.Prettify(s)
}

func (s UpdateRuleAttributeRequestRuleActionsForwardGroupConfigServerGroupTuples) GoString() string {
	return s.String()
}

func (s *UpdateRuleAttributeRequestRuleActionsForwardGroupConfigServerGroupTuples) GetServerGroupId() *string {
	return s.ServerGroupId
}

func (s *UpdateRuleAttributeRequestRuleActionsForwardGroupConfigServerGroupTuples) GetWeight() *int32 {
	return s.Weight
}

func (s *UpdateRuleAttributeRequestRuleActionsForwardGroupConfigServerGroupTuples) SetServerGroupId(v string) *UpdateRuleAttributeRequestRuleActionsForwardGroupConfigServerGroupTuples {
	s.ServerGroupId = &v
	return s
}

func (s *UpdateRuleAttributeRequestRuleActionsForwardGroupConfigServerGroupTuples) SetWeight(v int32) *UpdateRuleAttributeRequestRuleActionsForwardGroupConfigServerGroupTuples {
	s.Weight = &v
	return s
}

func (s *UpdateRuleAttributeRequestRuleActionsForwardGroupConfigServerGroupTuples) Validate() error {
	return dara.Validate(s)
}

type UpdateRuleAttributeRequestRuleActionsInsertHeaderConfig struct {
	// Specifies whether to overwrite the request header values. Valid values:
	//
	// 	- **true**
	//
	// 	- **false*	- (default)
	//
	// example:
	//
	// false
	CoverEnabled *bool `json:"CoverEnabled,omitempty" xml:"CoverEnabled,omitempty"`
	// The key of the header. The key must be 1 to 40 characters in length, and can contain letters, digits, underscores (_), and hyphens (-). The header keys specified by **InsertHeaderConfig*	- must be unique.
	//
	// >  You cannot specify the following header keys: `slb-id`, `slb-ip`, `x-forwarded-for`, `x-forwarded-proto`, `x-forwarded-eip`, `x-forwarded-port`, `x-forwarded-client-srcport`, `x-forwarded-host`, `connection`, `upgrade`, `content-length`, `transfer-encoding`, `keep-alive`, `te`, `host`, `cookie`, `remoteip`, and `authority`. The header keys are case-insensitive.
	//
	// example:
	//
	// key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the header.
	//
	// 	- If **ValueType*	- is set to **SystemDefined**, you can set the Value parameter to one of the following values:
	//
	//     	- **ClientSrcPort**: the client port.
	//
	//     	- **ClientSrcIp**: the IP address of the client.
	//
	//     	- **Protocol**: the request protocol (HTTP or HTTPS).
	//
	//     	- **SLBId**: the ID of the ALB instance.
	//
	//     	- **SLBPort**: the listener port of the ALB instance.
	//
	// 	- If **ValueType*	- is set to **UserDefined**, a custom header value is supported. The header value must be 1 to 128 characters in length, and can contain printable characters whose ASCII values are `greater than or equal to 32 and lower than 127`. You can use asterisks (\\*) and question marks (?) as wildcard characters. Quotation marks (`"`) are not supported. The header value cannot start or end with a space character, or end with a backslash (`\\`).
	//
	// 	- If **ValueType*	- is set to **ReferenceHeader**, you can reference a value from request headers. The value must be 1 to 128 characters in length, and can contain lowercase letters, digits, hyphens (-), and underscores (_).
	//
	// example:
	//
	// UserDefined
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
	// The type of the header. Valid values:
	//
	// 	- **UserDefined**: a custom header.
	//
	// 	- **ReferenceHeader**: a header that references one of the request headers.
	//
	// 	- **SystemDefined**: a system-defined header value.
	//
	// example:
	//
	// UserDefined
	ValueType *string `json:"ValueType,omitempty" xml:"ValueType,omitempty"`
}

func (s UpdateRuleAttributeRequestRuleActionsInsertHeaderConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateRuleAttributeRequestRuleActionsInsertHeaderConfig) GoString() string {
	return s.String()
}

func (s *UpdateRuleAttributeRequestRuleActionsInsertHeaderConfig) GetCoverEnabled() *bool {
	return s.CoverEnabled
}

func (s *UpdateRuleAttributeRequestRuleActionsInsertHeaderConfig) GetKey() *string {
	return s.Key
}

func (s *UpdateRuleAttributeRequestRuleActionsInsertHeaderConfig) GetValue() *string {
	return s.Value
}

func (s *UpdateRuleAttributeRequestRuleActionsInsertHeaderConfig) GetValueType() *string {
	return s.ValueType
}

func (s *UpdateRuleAttributeRequestRuleActionsInsertHeaderConfig) SetCoverEnabled(v bool) *UpdateRuleAttributeRequestRuleActionsInsertHeaderConfig {
	s.CoverEnabled = &v
	return s
}

func (s *UpdateRuleAttributeRequestRuleActionsInsertHeaderConfig) SetKey(v string) *UpdateRuleAttributeRequestRuleActionsInsertHeaderConfig {
	s.Key = &v
	return s
}

func (s *UpdateRuleAttributeRequestRuleActionsInsertHeaderConfig) SetValue(v string) *UpdateRuleAttributeRequestRuleActionsInsertHeaderConfig {
	s.Value = &v
	return s
}

func (s *UpdateRuleAttributeRequestRuleActionsInsertHeaderConfig) SetValueType(v string) *UpdateRuleAttributeRequestRuleActionsInsertHeaderConfig {
	s.ValueType = &v
	return s
}

func (s *UpdateRuleAttributeRequestRuleActionsInsertHeaderConfig) Validate() error {
	return dara.Validate(s)
}

type UpdateRuleAttributeRequestRuleActionsRedirectConfig struct {
	// The hostname to which requests are redirected. Valid values:
	//
	// 	- **${host}*	- (default): If ${host} is returned, no other character is appended.
	//
	// 	- If you want to specify a custom value, make sure that the following requirements are met:
	//
	//     	- The hostname must be 3 to 128 characters in length, and can contain lowercase letters, digits, hyphens (-), periods (.), asterisks (\\*), and question marks (?).
	//
	//     	- The hostname contains at least one period (.) but does not start or end with a period (.).
	//
	//     	- The rightmost domain label can contain only letters and wildcard characters. It cannot contain digits or hyphens (-).
	//
	//     	- The domain labels do not start or end with a hyphen (-).
	//
	//     	- You can use asterisks (\\*) and question marks (?) anywhere in a domain label as wildcard characters.
	//
	// example:
	//
	// www.example.com
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// The forwarding method. Valid values: **301**, **302**, **303**, **307**, and **308**.
	//
	// example:
	//
	// 301
	HttpCode *string `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	// The URL to which requests are redirected. Valid values:
	//
	// 	- Default value: **${path}**. \\*\\*${host}**, **${protocol}**, and **${port}\\*\\	- are also supported. Each variable can be specified only once. The preceding variables can be used at the same time or combined with a custom value.
	//
	// 	- If you want to specify a custom value, make sure that the following requirements are met:
	//
	//     	- The header value must be 1 to 128 characters in length.
	//
	//     	- It must start with a forward slash (/) and can contain letters, digits, and the following special characters: `$ - _ . + / & ~ @ :`. It does not contain the following special characters: `% # ; ! ( ) [ ] ^ , \\ "`. You can use asterisks (\\*) and question marks (?) as wildcard characters.
	//
	// example:
	//
	// /test
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// The port to which requests are redirected. Valid values:
	//
	// 	- **${port}*	- (default): If you set the value to ${port}, you cannot append other characters.
	//
	// 	- Other valid values: **1 to 63335**.
	//
	// example:
	//
	// 10
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// The redirect protocol. Valid values:
	//
	// 	- **${protocol}*	- (default): If you set the value to ${protocol}, you cannot append other characters.
	//
	// 	- **HTTP*	- or **HTTPS**.
	//
	// >  HTTPS listeners support only HTTPS redirects.
	//
	// example:
	//
	// HTTP
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// The query string of the URL to which requests are forwarded. Valid values:
	//
	// 	- Default value: **${query}**. \\*\\*${host}**, **${protocol}**, and **${port}\\*\\	- are also supported. Each variable can be specified only once. The preceding variables can be used at the same time or combined with a custom value.
	//
	// 	- If you want to specify a custom value, make sure that the following requirements are met:
	//
	//     	- The header value must be 1 to 128 characters in length.
	//
	//     	- It can contain printable characters, excluding space characters, the special characters `# [ ] { } \\ | < > "`, and uppercase letters.
	//
	// example:
	//
	// quert
	Query *string `json:"Query,omitempty" xml:"Query,omitempty"`
}

func (s UpdateRuleAttributeRequestRuleActionsRedirectConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateRuleAttributeRequestRuleActionsRedirectConfig) GoString() string {
	return s.String()
}

func (s *UpdateRuleAttributeRequestRuleActionsRedirectConfig) GetHost() *string {
	return s.Host
}

func (s *UpdateRuleAttributeRequestRuleActionsRedirectConfig) GetHttpCode() *string {
	return s.HttpCode
}

func (s *UpdateRuleAttributeRequestRuleActionsRedirectConfig) GetPath() *string {
	return s.Path
}

func (s *UpdateRuleAttributeRequestRuleActionsRedirectConfig) GetPort() *string {
	return s.Port
}

func (s *UpdateRuleAttributeRequestRuleActionsRedirectConfig) GetProtocol() *string {
	return s.Protocol
}

func (s *UpdateRuleAttributeRequestRuleActionsRedirectConfig) GetQuery() *string {
	return s.Query
}

func (s *UpdateRuleAttributeRequestRuleActionsRedirectConfig) SetHost(v string) *UpdateRuleAttributeRequestRuleActionsRedirectConfig {
	s.Host = &v
	return s
}

func (s *UpdateRuleAttributeRequestRuleActionsRedirectConfig) SetHttpCode(v string) *UpdateRuleAttributeRequestRuleActionsRedirectConfig {
	s.HttpCode = &v
	return s
}

func (s *UpdateRuleAttributeRequestRuleActionsRedirectConfig) SetPath(v string) *UpdateRuleAttributeRequestRuleActionsRedirectConfig {
	s.Path = &v
	return s
}

func (s *UpdateRuleAttributeRequestRuleActionsRedirectConfig) SetPort(v string) *UpdateRuleAttributeRequestRuleActionsRedirectConfig {
	s.Port = &v
	return s
}

func (s *UpdateRuleAttributeRequestRuleActionsRedirectConfig) SetProtocol(v string) *UpdateRuleAttributeRequestRuleActionsRedirectConfig {
	s.Protocol = &v
	return s
}

func (s *UpdateRuleAttributeRequestRuleActionsRedirectConfig) SetQuery(v string) *UpdateRuleAttributeRequestRuleActionsRedirectConfig {
	s.Query = &v
	return s
}

func (s *UpdateRuleAttributeRequestRuleActionsRedirectConfig) Validate() error {
	return dara.Validate(s)
}

type UpdateRuleAttributeRequestRuleActionsRemoveHeaderConfig struct {
	// The key of the header to be removed. The header key must be 1 to 40 characters in length, and can contain letters, digits, underscores (_), and hyphens (-). The header keys specified in RemoveHeader must be unique.
	//
	// 	- If Direction is set to Request, the following request headers cannot be removed: `slb-id`, `slb-ip`, `x-forwarded-for`, `x-forwarded-proto`, `x-forwarded-eip`, `x-forwarded-port`, `x-forwarded-client-srcport`, `x-forwarded-host`, `connection`, `upgrade`, `content-length`, `transfer-encoding`, `keep-alive`, `te`, `host`, `cookie`, `remoteip`, and `authority`. Request headers are not case-sensitive.
	//
	// 	- If Direction is set to Response, the following header keys are not supported: `connection`, `upgrade`, `content-length`, and `transfer-encoding`. The header keys are not case-sensitive.
	//
	// example:
	//
	// test
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
}

func (s UpdateRuleAttributeRequestRuleActionsRemoveHeaderConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateRuleAttributeRequestRuleActionsRemoveHeaderConfig) GoString() string {
	return s.String()
}

func (s *UpdateRuleAttributeRequestRuleActionsRemoveHeaderConfig) GetKey() *string {
	return s.Key
}

func (s *UpdateRuleAttributeRequestRuleActionsRemoveHeaderConfig) SetKey(v string) *UpdateRuleAttributeRequestRuleActionsRemoveHeaderConfig {
	s.Key = &v
	return s
}

func (s *UpdateRuleAttributeRequestRuleActionsRemoveHeaderConfig) Validate() error {
	return dara.Validate(s)
}

type UpdateRuleAttributeRequestRuleActionsRewriteConfig struct {
	// The hostname to which requests are rewritten. Valid values:
	//
	// 	- **${host}*	- (default): If you set the value to ${host}, you cannot append other characters.
	//
	// 	- If you want to specify a custom value, make sure that the following requirements are met:
	//
	//     	- The hostname must be 3 to 128 characters in length, and can contain lowercase letters, digits, hyphens (-), periods (.), asterisks (\\*), and question marks (?).
	//
	//     	- The hostname contains at least one period (.) but does not start or end with a period (.).
	//
	//     	- The rightmost domain label can contain only letters and wildcard characters. It cannot contain digits or hyphens (-).
	//
	//     	- The domain labels do not start or end with a hyphen (-). You can use asterisks (\\*) and question marks (?) anywhere in a domain label as wildcard characters.
	//
	// example:
	//
	// www.example.com
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// The URL to which requests are redirected. Valid values:
	//
	// 	- Default value: **${path}**. \\*\\*${host}**, **${protocol}**, and **${port}\\*\\	- are also supported. Each variable can be specified only once. The preceding variables can be used at the same time or combined with a custom value.
	//
	// 	- If you want to specify a custom value, make sure that the following requirements are met:
	//
	//     	- The header value must be 1 to 128 characters in length.
	//
	//     	- It must start with a forward slash (/) and can contain letters, digits, and the following special characters: `$ - _ . + / & ~ @ :`. It does not contain the following special characters: `% # ; ! ( ) [ ] ^ , \\ "`. You can use asterisks (\\*) and question marks (?) as wildcard characters.
	//
	// example:
	//
	// /tsdf
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// The query string to which requests are redirected. Valid values:
	//
	// 	- Default value: **${query}**. \\*\\*${host}**, **${protocol}**, and **${port}\\*\\	- are also supported. Each variable can be specified only once. The preceding variables can be used at the same time or combined with a custom value.
	//
	// 	- If you want to specify a custom value, make sure that the following requirements are met:
	//
	//     	- The header value must be 1 to 128 characters in length.
	//
	//     	- It can contain printable characters, excluding space characters, the special characters `# [ ] { } \\ | < > "`, and uppercase letters.
	//
	// example:
	//
	// quedsa
	Query *string `json:"Query,omitempty" xml:"Query,omitempty"`
}

func (s UpdateRuleAttributeRequestRuleActionsRewriteConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateRuleAttributeRequestRuleActionsRewriteConfig) GoString() string {
	return s.String()
}

func (s *UpdateRuleAttributeRequestRuleActionsRewriteConfig) GetHost() *string {
	return s.Host
}

func (s *UpdateRuleAttributeRequestRuleActionsRewriteConfig) GetPath() *string {
	return s.Path
}

func (s *UpdateRuleAttributeRequestRuleActionsRewriteConfig) GetQuery() *string {
	return s.Query
}

func (s *UpdateRuleAttributeRequestRuleActionsRewriteConfig) SetHost(v string) *UpdateRuleAttributeRequestRuleActionsRewriteConfig {
	s.Host = &v
	return s
}

func (s *UpdateRuleAttributeRequestRuleActionsRewriteConfig) SetPath(v string) *UpdateRuleAttributeRequestRuleActionsRewriteConfig {
	s.Path = &v
	return s
}

func (s *UpdateRuleAttributeRequestRuleActionsRewriteConfig) SetQuery(v string) *UpdateRuleAttributeRequestRuleActionsRewriteConfig {
	s.Query = &v
	return s
}

func (s *UpdateRuleAttributeRequestRuleActionsRewriteConfig) Validate() error {
	return dara.Validate(s)
}

type UpdateRuleAttributeRequestRuleActionsTrafficLimitConfig struct {
	// The number of requests per IP address. Value range: **1 to 1,000,000**.
	//
	// >  If both the **QPS*	- and **PerIpQps*	- parameters are specified, make sure that the value of the **QPS*	- parameter is smaller than the value of the PerIpQps parameter.
	//
	// example:
	//
	// 80
	PerIpQps *int32 `json:"PerIpQps,omitempty" xml:"PerIpQps,omitempty"`
	// The number of queries per second (QPS). Value range: **1 to 1,000,000**.
	//
	// example:
	//
	// 100
	QPS *int32 `json:"QPS,omitempty" xml:"QPS,omitempty"`
}

func (s UpdateRuleAttributeRequestRuleActionsTrafficLimitConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateRuleAttributeRequestRuleActionsTrafficLimitConfig) GoString() string {
	return s.String()
}

func (s *UpdateRuleAttributeRequestRuleActionsTrafficLimitConfig) GetPerIpQps() *int32 {
	return s.PerIpQps
}

func (s *UpdateRuleAttributeRequestRuleActionsTrafficLimitConfig) GetQPS() *int32 {
	return s.QPS
}

func (s *UpdateRuleAttributeRequestRuleActionsTrafficLimitConfig) SetPerIpQps(v int32) *UpdateRuleAttributeRequestRuleActionsTrafficLimitConfig {
	s.PerIpQps = &v
	return s
}

func (s *UpdateRuleAttributeRequestRuleActionsTrafficLimitConfig) SetQPS(v int32) *UpdateRuleAttributeRequestRuleActionsTrafficLimitConfig {
	s.QPS = &v
	return s
}

func (s *UpdateRuleAttributeRequestRuleActionsTrafficLimitConfig) Validate() error {
	return dara.Validate(s)
}

type UpdateRuleAttributeRequestRuleActionsTrafficMirrorConfig struct {
	// The server group to which network traffic is mirrored.
	MirrorGroupConfig *UpdateRuleAttributeRequestRuleActionsTrafficMirrorConfigMirrorGroupConfig `json:"MirrorGroupConfig,omitempty" xml:"MirrorGroupConfig,omitempty" type:"Struct"`
	// The type of destination to which network traffic is mirrored. Valid values:
	//
	// 	- **ForwardGroupMirror**: a server group
	//
	// example:
	//
	// ForwardGroupMirror
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
}

func (s UpdateRuleAttributeRequestRuleActionsTrafficMirrorConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateRuleAttributeRequestRuleActionsTrafficMirrorConfig) GoString() string {
	return s.String()
}

func (s *UpdateRuleAttributeRequestRuleActionsTrafficMirrorConfig) GetMirrorGroupConfig() *UpdateRuleAttributeRequestRuleActionsTrafficMirrorConfigMirrorGroupConfig {
	return s.MirrorGroupConfig
}

func (s *UpdateRuleAttributeRequestRuleActionsTrafficMirrorConfig) GetTargetType() *string {
	return s.TargetType
}

func (s *UpdateRuleAttributeRequestRuleActionsTrafficMirrorConfig) SetMirrorGroupConfig(v *UpdateRuleAttributeRequestRuleActionsTrafficMirrorConfigMirrorGroupConfig) *UpdateRuleAttributeRequestRuleActionsTrafficMirrorConfig {
	s.MirrorGroupConfig = v
	return s
}

func (s *UpdateRuleAttributeRequestRuleActionsTrafficMirrorConfig) SetTargetType(v string) *UpdateRuleAttributeRequestRuleActionsTrafficMirrorConfig {
	s.TargetType = &v
	return s
}

func (s *UpdateRuleAttributeRequestRuleActionsTrafficMirrorConfig) Validate() error {
	return dara.Validate(s)
}

type UpdateRuleAttributeRequestRuleActionsTrafficMirrorConfigMirrorGroupConfig struct {
	// The server group to which traffic is mirrored.
	ServerGroupTuples []*UpdateRuleAttributeRequestRuleActionsTrafficMirrorConfigMirrorGroupConfigServerGroupTuples `json:"ServerGroupTuples,omitempty" xml:"ServerGroupTuples,omitempty" type:"Repeated"`
}

func (s UpdateRuleAttributeRequestRuleActionsTrafficMirrorConfigMirrorGroupConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateRuleAttributeRequestRuleActionsTrafficMirrorConfigMirrorGroupConfig) GoString() string {
	return s.String()
}

func (s *UpdateRuleAttributeRequestRuleActionsTrafficMirrorConfigMirrorGroupConfig) GetServerGroupTuples() []*UpdateRuleAttributeRequestRuleActionsTrafficMirrorConfigMirrorGroupConfigServerGroupTuples {
	return s.ServerGroupTuples
}

func (s *UpdateRuleAttributeRequestRuleActionsTrafficMirrorConfigMirrorGroupConfig) SetServerGroupTuples(v []*UpdateRuleAttributeRequestRuleActionsTrafficMirrorConfigMirrorGroupConfigServerGroupTuples) *UpdateRuleAttributeRequestRuleActionsTrafficMirrorConfigMirrorGroupConfig {
	s.ServerGroupTuples = v
	return s
}

func (s *UpdateRuleAttributeRequestRuleActionsTrafficMirrorConfigMirrorGroupConfig) Validate() error {
	return dara.Validate(s)
}

type UpdateRuleAttributeRequestRuleActionsTrafficMirrorConfigMirrorGroupConfigServerGroupTuples struct {
	// The server group ID.
	//
	// example:
	//
	// srg-00mkgijak0w4qgz9****
	ServerGroupId *string `json:"ServerGroupId,omitempty" xml:"ServerGroupId,omitempty"`
}

func (s UpdateRuleAttributeRequestRuleActionsTrafficMirrorConfigMirrorGroupConfigServerGroupTuples) String() string {
	return dara.Prettify(s)
}

func (s UpdateRuleAttributeRequestRuleActionsTrafficMirrorConfigMirrorGroupConfigServerGroupTuples) GoString() string {
	return s.String()
}

func (s *UpdateRuleAttributeRequestRuleActionsTrafficMirrorConfigMirrorGroupConfigServerGroupTuples) GetServerGroupId() *string {
	return s.ServerGroupId
}

func (s *UpdateRuleAttributeRequestRuleActionsTrafficMirrorConfigMirrorGroupConfigServerGroupTuples) SetServerGroupId(v string) *UpdateRuleAttributeRequestRuleActionsTrafficMirrorConfigMirrorGroupConfigServerGroupTuples {
	s.ServerGroupId = &v
	return s
}

func (s *UpdateRuleAttributeRequestRuleActionsTrafficMirrorConfigMirrorGroupConfigServerGroupTuples) Validate() error {
	return dara.Validate(s)
}

type UpdateRuleAttributeRequestRuleConditions struct {
	// The key-value pairs of the cookie.
	CookieConfig *UpdateRuleAttributeRequestRuleConditionsCookieConfig `json:"CookieConfig,omitempty" xml:"CookieConfig,omitempty" type:"Struct"`
	// The configuration of the header.
	HeaderConfig *UpdateRuleAttributeRequestRuleConditionsHeaderConfig `json:"HeaderConfig,omitempty" xml:"HeaderConfig,omitempty" type:"Struct"`
	// The configuration of the hosts.
	HostConfig *UpdateRuleAttributeRequestRuleConditionsHostConfig `json:"HostConfig,omitempty" xml:"HostConfig,omitempty" type:"Struct"`
	// The configuration of the request method.
	MethodConfig *UpdateRuleAttributeRequestRuleConditionsMethodConfig `json:"MethodConfig,omitempty" xml:"MethodConfig,omitempty" type:"Struct"`
	// The configuration of the forwarding URL.
	PathConfig *UpdateRuleAttributeRequestRuleConditionsPathConfig `json:"PathConfig,omitempty" xml:"PathConfig,omitempty" type:"Struct"`
	// The configuration of the query strings.
	QueryStringConfig *UpdateRuleAttributeRequestRuleConditionsQueryStringConfig `json:"QueryStringConfig,omitempty" xml:"QueryStringConfig,omitempty" type:"Struct"`
	// The configuration of headers.
	ResponseHeaderConfig *UpdateRuleAttributeRequestRuleConditionsResponseHeaderConfig `json:"ResponseHeaderConfig,omitempty" xml:"ResponseHeaderConfig,omitempty" type:"Struct"`
	// The configuration of the response status codes.
	ResponseStatusCodeConfig *UpdateRuleAttributeRequestRuleConditionsResponseStatusCodeConfig `json:"ResponseStatusCodeConfig,omitempty" xml:"ResponseStatusCodeConfig,omitempty" type:"Struct"`
	// Traffic matching based on source IP addresses. You can specify up to five IP addresses, including CIDR blocks.
	SourceIpConfig *UpdateRuleAttributeRequestRuleConditionsSourceIpConfig `json:"SourceIpConfig,omitempty" xml:"SourceIpConfig,omitempty" type:"Struct"`
	// The type of forwarding rule. You can specify up to seven types of forwarding rule. Valid values:
	//
	// 	- **Host**: Requests are forwarded based on hosts.
	//
	// 	- **Path**: Requests are forwarded based on URLs.
	//
	// 	- **Header**: Requests are forwarded based on HTTP headers.
	//
	// 	- **QueryString**: Requests are forwarded based on query strings.
	//
	// 	- **Method**: Requests are forwarded based on request methods.
	//
	// 	- **Cookie**: Requests are forwarded based on cookies.
	//
	// 	- **SourceIp**: Requests are forwarded based on source IP addresses.
	//
	// 	- **ResponseHeader**: Requests are forwarded based on HTTP response headers.
	//
	// 	- **ResponseStatusCode**: Requests are forwarded based on response status codes.
	//
	// example:
	//
	// Host
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s UpdateRuleAttributeRequestRuleConditions) String() string {
	return dara.Prettify(s)
}

func (s UpdateRuleAttributeRequestRuleConditions) GoString() string {
	return s.String()
}

func (s *UpdateRuleAttributeRequestRuleConditions) GetCookieConfig() *UpdateRuleAttributeRequestRuleConditionsCookieConfig {
	return s.CookieConfig
}

func (s *UpdateRuleAttributeRequestRuleConditions) GetHeaderConfig() *UpdateRuleAttributeRequestRuleConditionsHeaderConfig {
	return s.HeaderConfig
}

func (s *UpdateRuleAttributeRequestRuleConditions) GetHostConfig() *UpdateRuleAttributeRequestRuleConditionsHostConfig {
	return s.HostConfig
}

func (s *UpdateRuleAttributeRequestRuleConditions) GetMethodConfig() *UpdateRuleAttributeRequestRuleConditionsMethodConfig {
	return s.MethodConfig
}

func (s *UpdateRuleAttributeRequestRuleConditions) GetPathConfig() *UpdateRuleAttributeRequestRuleConditionsPathConfig {
	return s.PathConfig
}

func (s *UpdateRuleAttributeRequestRuleConditions) GetQueryStringConfig() *UpdateRuleAttributeRequestRuleConditionsQueryStringConfig {
	return s.QueryStringConfig
}

func (s *UpdateRuleAttributeRequestRuleConditions) GetResponseHeaderConfig() *UpdateRuleAttributeRequestRuleConditionsResponseHeaderConfig {
	return s.ResponseHeaderConfig
}

func (s *UpdateRuleAttributeRequestRuleConditions) GetResponseStatusCodeConfig() *UpdateRuleAttributeRequestRuleConditionsResponseStatusCodeConfig {
	return s.ResponseStatusCodeConfig
}

func (s *UpdateRuleAttributeRequestRuleConditions) GetSourceIpConfig() *UpdateRuleAttributeRequestRuleConditionsSourceIpConfig {
	return s.SourceIpConfig
}

func (s *UpdateRuleAttributeRequestRuleConditions) GetType() *string {
	return s.Type
}

func (s *UpdateRuleAttributeRequestRuleConditions) SetCookieConfig(v *UpdateRuleAttributeRequestRuleConditionsCookieConfig) *UpdateRuleAttributeRequestRuleConditions {
	s.CookieConfig = v
	return s
}

func (s *UpdateRuleAttributeRequestRuleConditions) SetHeaderConfig(v *UpdateRuleAttributeRequestRuleConditionsHeaderConfig) *UpdateRuleAttributeRequestRuleConditions {
	s.HeaderConfig = v
	return s
}

func (s *UpdateRuleAttributeRequestRuleConditions) SetHostConfig(v *UpdateRuleAttributeRequestRuleConditionsHostConfig) *UpdateRuleAttributeRequestRuleConditions {
	s.HostConfig = v
	return s
}

func (s *UpdateRuleAttributeRequestRuleConditions) SetMethodConfig(v *UpdateRuleAttributeRequestRuleConditionsMethodConfig) *UpdateRuleAttributeRequestRuleConditions {
	s.MethodConfig = v
	return s
}

func (s *UpdateRuleAttributeRequestRuleConditions) SetPathConfig(v *UpdateRuleAttributeRequestRuleConditionsPathConfig) *UpdateRuleAttributeRequestRuleConditions {
	s.PathConfig = v
	return s
}

func (s *UpdateRuleAttributeRequestRuleConditions) SetQueryStringConfig(v *UpdateRuleAttributeRequestRuleConditionsQueryStringConfig) *UpdateRuleAttributeRequestRuleConditions {
	s.QueryStringConfig = v
	return s
}

func (s *UpdateRuleAttributeRequestRuleConditions) SetResponseHeaderConfig(v *UpdateRuleAttributeRequestRuleConditionsResponseHeaderConfig) *UpdateRuleAttributeRequestRuleConditions {
	s.ResponseHeaderConfig = v
	return s
}

func (s *UpdateRuleAttributeRequestRuleConditions) SetResponseStatusCodeConfig(v *UpdateRuleAttributeRequestRuleConditionsResponseStatusCodeConfig) *UpdateRuleAttributeRequestRuleConditions {
	s.ResponseStatusCodeConfig = v
	return s
}

func (s *UpdateRuleAttributeRequestRuleConditions) SetSourceIpConfig(v *UpdateRuleAttributeRequestRuleConditionsSourceIpConfig) *UpdateRuleAttributeRequestRuleConditions {
	s.SourceIpConfig = v
	return s
}

func (s *UpdateRuleAttributeRequestRuleConditions) SetType(v string) *UpdateRuleAttributeRequestRuleConditions {
	s.Type = &v
	return s
}

func (s *UpdateRuleAttributeRequestRuleConditions) Validate() error {
	return dara.Validate(s)
}

type UpdateRuleAttributeRequestRuleConditionsCookieConfig struct {
	// The key-value pairs of the cookie.
	Values []*UpdateRuleAttributeRequestRuleConditionsCookieConfigValues `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s UpdateRuleAttributeRequestRuleConditionsCookieConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateRuleAttributeRequestRuleConditionsCookieConfig) GoString() string {
	return s.String()
}

func (s *UpdateRuleAttributeRequestRuleConditionsCookieConfig) GetValues() []*UpdateRuleAttributeRequestRuleConditionsCookieConfigValues {
	return s.Values
}

func (s *UpdateRuleAttributeRequestRuleConditionsCookieConfig) SetValues(v []*UpdateRuleAttributeRequestRuleConditionsCookieConfigValues) *UpdateRuleAttributeRequestRuleConditionsCookieConfig {
	s.Values = v
	return s
}

func (s *UpdateRuleAttributeRequestRuleConditionsCookieConfig) Validate() error {
	return dara.Validate(s)
}

type UpdateRuleAttributeRequestRuleConditionsCookieConfigValues struct {
	// The cookie key. The key must be 1 to 100 characters in length, and can contain printable characters such as lowercase letters, asterisks (\\*), and question marks (?). The key cannot contain uppercase letters, space characters, or the following special characters: `# [ ] { } \\ | < > & " ;`.
	//
	// example:
	//
	// test
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The cookie value. The cookie value must be 1 to 128 characters in length, and can contain lowercase letters, printable ASCII characters, asterisks (\\*), and question marks (?). It cannot contain space characters or the following special characters: `# [ ] { } \\ | < > &`.
	//
	// example:
	//
	// test
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateRuleAttributeRequestRuleConditionsCookieConfigValues) String() string {
	return dara.Prettify(s)
}

func (s UpdateRuleAttributeRequestRuleConditionsCookieConfigValues) GoString() string {
	return s.String()
}

func (s *UpdateRuleAttributeRequestRuleConditionsCookieConfigValues) GetKey() *string {
	return s.Key
}

func (s *UpdateRuleAttributeRequestRuleConditionsCookieConfigValues) GetValue() *string {
	return s.Value
}

func (s *UpdateRuleAttributeRequestRuleConditionsCookieConfigValues) SetKey(v string) *UpdateRuleAttributeRequestRuleConditionsCookieConfigValues {
	s.Key = &v
	return s
}

func (s *UpdateRuleAttributeRequestRuleConditionsCookieConfigValues) SetValue(v string) *UpdateRuleAttributeRequestRuleConditionsCookieConfigValues {
	s.Value = &v
	return s
}

func (s *UpdateRuleAttributeRequestRuleConditionsCookieConfigValues) Validate() error {
	return dara.Validate(s)
}

type UpdateRuleAttributeRequestRuleConditionsHeaderConfig struct {
	// The header key. The header key must be 1 to 40 characters in length, and can contain letters, digits, hyphens (-), and underscores (_). Cookie and Host are not supported.
	//
	// example:
	//
	// Port
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The header values.
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s UpdateRuleAttributeRequestRuleConditionsHeaderConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateRuleAttributeRequestRuleConditionsHeaderConfig) GoString() string {
	return s.String()
}

func (s *UpdateRuleAttributeRequestRuleConditionsHeaderConfig) GetKey() *string {
	return s.Key
}

func (s *UpdateRuleAttributeRequestRuleConditionsHeaderConfig) GetValues() []*string {
	return s.Values
}

func (s *UpdateRuleAttributeRequestRuleConditionsHeaderConfig) SetKey(v string) *UpdateRuleAttributeRequestRuleConditionsHeaderConfig {
	s.Key = &v
	return s
}

func (s *UpdateRuleAttributeRequestRuleConditionsHeaderConfig) SetValues(v []*string) *UpdateRuleAttributeRequestRuleConditionsHeaderConfig {
	s.Values = v
	return s
}

func (s *UpdateRuleAttributeRequestRuleConditionsHeaderConfig) Validate() error {
	return dara.Validate(s)
}

type UpdateRuleAttributeRequestRuleConditionsHostConfig struct {
	// The hostnames.
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s UpdateRuleAttributeRequestRuleConditionsHostConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateRuleAttributeRequestRuleConditionsHostConfig) GoString() string {
	return s.String()
}

func (s *UpdateRuleAttributeRequestRuleConditionsHostConfig) GetValues() []*string {
	return s.Values
}

func (s *UpdateRuleAttributeRequestRuleConditionsHostConfig) SetValues(v []*string) *UpdateRuleAttributeRequestRuleConditionsHostConfig {
	s.Values = v
	return s
}

func (s *UpdateRuleAttributeRequestRuleConditionsHostConfig) Validate() error {
	return dara.Validate(s)
}

type UpdateRuleAttributeRequestRuleConditionsMethodConfig struct {
	// The request methods.
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s UpdateRuleAttributeRequestRuleConditionsMethodConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateRuleAttributeRequestRuleConditionsMethodConfig) GoString() string {
	return s.String()
}

func (s *UpdateRuleAttributeRequestRuleConditionsMethodConfig) GetValues() []*string {
	return s.Values
}

func (s *UpdateRuleAttributeRequestRuleConditionsMethodConfig) SetValues(v []*string) *UpdateRuleAttributeRequestRuleConditionsMethodConfig {
	s.Values = v
	return s
}

func (s *UpdateRuleAttributeRequestRuleConditionsMethodConfig) Validate() error {
	return dara.Validate(s)
}

type UpdateRuleAttributeRequestRuleConditionsPathConfig struct {
	// The forwarding URLs.
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s UpdateRuleAttributeRequestRuleConditionsPathConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateRuleAttributeRequestRuleConditionsPathConfig) GoString() string {
	return s.String()
}

func (s *UpdateRuleAttributeRequestRuleConditionsPathConfig) GetValues() []*string {
	return s.Values
}

func (s *UpdateRuleAttributeRequestRuleConditionsPathConfig) SetValues(v []*string) *UpdateRuleAttributeRequestRuleConditionsPathConfig {
	s.Values = v
	return s
}

func (s *UpdateRuleAttributeRequestRuleConditionsPathConfig) Validate() error {
	return dara.Validate(s)
}

type UpdateRuleAttributeRequestRuleConditionsQueryStringConfig struct {
	// The query strings. You can specify up to 20 query strings.
	Values []*UpdateRuleAttributeRequestRuleConditionsQueryStringConfigValues `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s UpdateRuleAttributeRequestRuleConditionsQueryStringConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateRuleAttributeRequestRuleConditionsQueryStringConfig) GoString() string {
	return s.String()
}

func (s *UpdateRuleAttributeRequestRuleConditionsQueryStringConfig) GetValues() []*UpdateRuleAttributeRequestRuleConditionsQueryStringConfigValues {
	return s.Values
}

func (s *UpdateRuleAttributeRequestRuleConditionsQueryStringConfig) SetValues(v []*UpdateRuleAttributeRequestRuleConditionsQueryStringConfigValues) *UpdateRuleAttributeRequestRuleConditionsQueryStringConfig {
	s.Values = v
	return s
}

func (s *UpdateRuleAttributeRequestRuleConditionsQueryStringConfig) Validate() error {
	return dara.Validate(s)
}

type UpdateRuleAttributeRequestRuleConditionsQueryStringConfigValues struct {
	// The key of the query string. The key must be 1 to 100 characters in length, and can contain printable characters such as lowercase letters, asterisks (\\*), and question marks (?). The key cannot contain uppercase letters, space characters, or the following special characters: `# [ ] { } \\ | < > & "`.
	//
	// example:
	//
	// test
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the query string. The value must be 1 to 128 characters in length, and can contain printable characters such as lowercase letters, asterisks (\\*), and question marks (?). The value cannot contain uppercase letters, space characters, or the following special characters: `# [ ] { } \\ | < > & "`.
	//
	// example:
	//
	// test
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateRuleAttributeRequestRuleConditionsQueryStringConfigValues) String() string {
	return dara.Prettify(s)
}

func (s UpdateRuleAttributeRequestRuleConditionsQueryStringConfigValues) GoString() string {
	return s.String()
}

func (s *UpdateRuleAttributeRequestRuleConditionsQueryStringConfigValues) GetKey() *string {
	return s.Key
}

func (s *UpdateRuleAttributeRequestRuleConditionsQueryStringConfigValues) GetValue() *string {
	return s.Value
}

func (s *UpdateRuleAttributeRequestRuleConditionsQueryStringConfigValues) SetKey(v string) *UpdateRuleAttributeRequestRuleConditionsQueryStringConfigValues {
	s.Key = &v
	return s
}

func (s *UpdateRuleAttributeRequestRuleConditionsQueryStringConfigValues) SetValue(v string) *UpdateRuleAttributeRequestRuleConditionsQueryStringConfigValues {
	s.Value = &v
	return s
}

func (s *UpdateRuleAttributeRequestRuleConditionsQueryStringConfigValues) Validate() error {
	return dara.Validate(s)
}

type UpdateRuleAttributeRequestRuleConditionsResponseHeaderConfig struct {
	// The header key.
	//
	// 	- The key must be 1 to 40 characters in length.
	//
	// 	- It can contain letters, digits, hyphens (-), and underscores (_).
	//
	// 	- Cookie and Host are not supported.
	//
	// example:
	//
	// test
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The header values.
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s UpdateRuleAttributeRequestRuleConditionsResponseHeaderConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateRuleAttributeRequestRuleConditionsResponseHeaderConfig) GoString() string {
	return s.String()
}

func (s *UpdateRuleAttributeRequestRuleConditionsResponseHeaderConfig) GetKey() *string {
	return s.Key
}

func (s *UpdateRuleAttributeRequestRuleConditionsResponseHeaderConfig) GetValues() []*string {
	return s.Values
}

func (s *UpdateRuleAttributeRequestRuleConditionsResponseHeaderConfig) SetKey(v string) *UpdateRuleAttributeRequestRuleConditionsResponseHeaderConfig {
	s.Key = &v
	return s
}

func (s *UpdateRuleAttributeRequestRuleConditionsResponseHeaderConfig) SetValues(v []*string) *UpdateRuleAttributeRequestRuleConditionsResponseHeaderConfig {
	s.Values = v
	return s
}

func (s *UpdateRuleAttributeRequestRuleConditionsResponseHeaderConfig) Validate() error {
	return dara.Validate(s)
}

type UpdateRuleAttributeRequestRuleConditionsResponseStatusCodeConfig struct {
	// The response status codes.
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s UpdateRuleAttributeRequestRuleConditionsResponseStatusCodeConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateRuleAttributeRequestRuleConditionsResponseStatusCodeConfig) GoString() string {
	return s.String()
}

func (s *UpdateRuleAttributeRequestRuleConditionsResponseStatusCodeConfig) GetValues() []*string {
	return s.Values
}

func (s *UpdateRuleAttributeRequestRuleConditionsResponseStatusCodeConfig) SetValues(v []*string) *UpdateRuleAttributeRequestRuleConditionsResponseStatusCodeConfig {
	s.Values = v
	return s
}

func (s *UpdateRuleAttributeRequestRuleConditionsResponseStatusCodeConfig) Validate() error {
	return dara.Validate(s)
}

type UpdateRuleAttributeRequestRuleConditionsSourceIpConfig struct {
	// You can add one or more IP addresses, including CIDR blocks.
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s UpdateRuleAttributeRequestRuleConditionsSourceIpConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateRuleAttributeRequestRuleConditionsSourceIpConfig) GoString() string {
	return s.String()
}

func (s *UpdateRuleAttributeRequestRuleConditionsSourceIpConfig) GetValues() []*string {
	return s.Values
}

func (s *UpdateRuleAttributeRequestRuleConditionsSourceIpConfig) SetValues(v []*string) *UpdateRuleAttributeRequestRuleConditionsSourceIpConfig {
	s.Values = v
	return s
}

func (s *UpdateRuleAttributeRequestRuleConditionsSourceIpConfig) Validate() error {
	return dara.Validate(s)
}
