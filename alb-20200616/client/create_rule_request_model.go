// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateRuleRequest
	GetClientToken() *string
	SetDirection(v string) *CreateRuleRequest
	GetDirection() *string
	SetDryRun(v bool) *CreateRuleRequest
	GetDryRun() *bool
	SetListenerId(v string) *CreateRuleRequest
	GetListenerId() *string
	SetPriority(v int32) *CreateRuleRequest
	GetPriority() *int32
	SetRuleActions(v []*CreateRuleRequestRuleActions) *CreateRuleRequest
	GetRuleActions() []*CreateRuleRequestRuleActions
	SetRuleConditions(v []*CreateRuleRequestRuleConditions) *CreateRuleRequest
	GetRuleConditions() []*CreateRuleRequestRuleConditions
	SetRuleName(v string) *CreateRuleRequest
	GetRuleName() *string
	SetTag(v []*CreateRuleRequestTag) *CreateRuleRequest
	GetTag() []*CreateRuleRequestTag
}

type CreateRuleRequest struct {
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
	// The direction to which the forwarding rule is applied. Valid values:
	//
	// 	- **Request*	- (default): The forwarding rule is applied to the requests received by ALB.
	//
	// 	- **Response**: The forwarding rule is applied to the responses returned by backend servers.
	//
	// > Basic ALB instances do not support the **Response*	- value.
	//
	// example:
	//
	// Request
	Direction *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
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
	// The listener ID of the ALB instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// lsr-bp1bpn0kn908w4nbw****
	ListenerId *string `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	// The priority of the forwarding rule. Valid values: **1 to 10000**. A smaller value indicates a higher priority.
	//
	// > The priorities of the forwarding rules created for the same listener must be unique.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The actions of the forwarding rule.
	//
	// This parameter is required.
	RuleActions []*CreateRuleRequestRuleActions `json:"RuleActions,omitempty" xml:"RuleActions,omitempty" type:"Repeated"`
	// The match conditions of the forwarding rule.
	//
	// This parameter is required.
	RuleConditions []*CreateRuleRequestRuleConditions `json:"RuleConditions,omitempty" xml:"RuleConditions,omitempty" type:"Repeated"`
	// The name of the forwarding rule.
	//
	// 	- The name must be 2 to 128 characters in length.
	//
	// 	- It can contain letters, digits, periods (.), underscores (_), and hyphens (-). It must start with a letter.
	//
	// This parameter is required.
	//
	// example:
	//
	// rule-doc
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The tags.
	Tag []*CreateRuleRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s CreateRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateRuleRequest) GoString() string {
	return s.String()
}

func (s *CreateRuleRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateRuleRequest) GetDirection() *string {
	return s.Direction
}

func (s *CreateRuleRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *CreateRuleRequest) GetListenerId() *string {
	return s.ListenerId
}

func (s *CreateRuleRequest) GetPriority() *int32 {
	return s.Priority
}

func (s *CreateRuleRequest) GetRuleActions() []*CreateRuleRequestRuleActions {
	return s.RuleActions
}

func (s *CreateRuleRequest) GetRuleConditions() []*CreateRuleRequestRuleConditions {
	return s.RuleConditions
}

func (s *CreateRuleRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *CreateRuleRequest) GetTag() []*CreateRuleRequestTag {
	return s.Tag
}

func (s *CreateRuleRequest) SetClientToken(v string) *CreateRuleRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateRuleRequest) SetDirection(v string) *CreateRuleRequest {
	s.Direction = &v
	return s
}

func (s *CreateRuleRequest) SetDryRun(v bool) *CreateRuleRequest {
	s.DryRun = &v
	return s
}

func (s *CreateRuleRequest) SetListenerId(v string) *CreateRuleRequest {
	s.ListenerId = &v
	return s
}

func (s *CreateRuleRequest) SetPriority(v int32) *CreateRuleRequest {
	s.Priority = &v
	return s
}

func (s *CreateRuleRequest) SetRuleActions(v []*CreateRuleRequestRuleActions) *CreateRuleRequest {
	s.RuleActions = v
	return s
}

func (s *CreateRuleRequest) SetRuleConditions(v []*CreateRuleRequestRuleConditions) *CreateRuleRequest {
	s.RuleConditions = v
	return s
}

func (s *CreateRuleRequest) SetRuleName(v string) *CreateRuleRequest {
	s.RuleName = &v
	return s
}

func (s *CreateRuleRequest) SetTag(v []*CreateRuleRequestTag) *CreateRuleRequest {
	s.Tag = v
	return s
}

func (s *CreateRuleRequest) Validate() error {
	return dara.Validate(s)
}

type CreateRuleRequestRuleActions struct {
	// Request forwarding based on CORS.
	CorsConfig *CreateRuleRequestRuleActionsCorsConfig `json:"CorsConfig,omitempty" xml:"CorsConfig,omitempty" type:"Struct"`
	// The configuration of the custom response.
	FixedResponseConfig *CreateRuleRequestRuleActionsFixedResponseConfig `json:"FixedResponseConfig,omitempty" xml:"FixedResponseConfig,omitempty" type:"Struct"`
	// The server groups to which requests are forwarded. You can specify at most five server groups for each forwarding rule.
	ForwardGroupConfig *CreateRuleRequestRuleActionsForwardGroupConfig `json:"ForwardGroupConfig,omitempty" xml:"ForwardGroupConfig,omitempty" type:"Struct"`
	// The key of the header to be inserted.
	InsertHeaderConfig *CreateRuleRequestRuleActionsInsertHeaderConfig `json:"InsertHeaderConfig,omitempty" xml:"InsertHeaderConfig,omitempty" type:"Struct"`
	// The priority of the action. Valid values: **1 to 50000**. A smaller value indicates a higher priority. The actions of a forwarding rule are applied in descending order of priority. This parameter cannot empty. The priority of each action within a forwarding rule must be unique.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	Order *int32 `json:"Order,omitempty" xml:"Order,omitempty"`
	// The configuration of the redirect action.
	//
	// >  Do not set all fields in **RedirectConfig*	- to default values, except for **httpCode**.
	RedirectConfig *CreateRuleRequestRuleActionsRedirectConfig `json:"RedirectConfig,omitempty" xml:"RedirectConfig,omitempty" type:"Struct"`
	// The HTTP headers to be removed.
	RemoveHeaderConfig *CreateRuleRequestRuleActionsRemoveHeaderConfig `json:"RemoveHeaderConfig,omitempty" xml:"RemoveHeaderConfig,omitempty" type:"Struct"`
	// The configuration of the rewrite action.
	//
	// >  If you specify multiple actions for a forwarding rule, you must configure the **ForwardGroup*	- parameter for the **RewriteConfig*	- action.
	RewriteConfig *CreateRuleRequestRuleActionsRewriteConfig `json:"RewriteConfig,omitempty" xml:"RewriteConfig,omitempty" type:"Struct"`
	// The configuration of the traffic throttling action.
	TrafficLimitConfig *CreateRuleRequestRuleActionsTrafficLimitConfig `json:"TrafficLimitConfig,omitempty" xml:"TrafficLimitConfig,omitempty" type:"Struct"`
	// The configuration of the traffic mirroring action.
	TrafficMirrorConfig *CreateRuleRequestRuleActionsTrafficMirrorConfig `json:"TrafficMirrorConfig,omitempty" xml:"TrafficMirrorConfig,omitempty" type:"Struct"`
	// The action. Valid values:
	//
	// 	- **ForwardGroup**: distributes requests to multiple vServer groups.
	//
	// 	- **Redirect**: redirects requests.
	//
	// 	- **FixedResponse**: returns a custom response.
	//
	// 	- **Rewrite**: rewrites requests.
	//
	// 	- **InsertHeader**: inserts headers.
	//
	// 	- **RemoveHeaderConfig:*	- deletes the header of a request.
	//
	// 	- **TrafficLimit**: throttles traffic.
	//
	// 	- **TrafficMirror**: mirrors network traffic.
	//
	// 	- **Cors**: enables cross-origin resource sharing (CORS).
	//
	// The following action types are supported:
	//
	// 	- **FinalType**: Each forwarding rule can contain only one FinalType action, which is performed at the end. You can specify only one of **ForwardGroup**, **Redirect**, and **FixedResponse**.
	//
	// 	- **ExtType**: Each forwarding rule can contain one or more **ExtType*	- actions, which are performed before the **FinalType*	- action. If you want to specify an ExtType action, you must also specify a **FinalType*	- action. You can specify multiple **InsertHeader*	- actions or one **Rewrite*	- action.
	//
	// This parameter is required.
	//
	// example:
	//
	// ForwardGroup
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateRuleRequestRuleActions) String() string {
	return dara.Prettify(s)
}

func (s CreateRuleRequestRuleActions) GoString() string {
	return s.String()
}

func (s *CreateRuleRequestRuleActions) GetCorsConfig() *CreateRuleRequestRuleActionsCorsConfig {
	return s.CorsConfig
}

func (s *CreateRuleRequestRuleActions) GetFixedResponseConfig() *CreateRuleRequestRuleActionsFixedResponseConfig {
	return s.FixedResponseConfig
}

func (s *CreateRuleRequestRuleActions) GetForwardGroupConfig() *CreateRuleRequestRuleActionsForwardGroupConfig {
	return s.ForwardGroupConfig
}

func (s *CreateRuleRequestRuleActions) GetInsertHeaderConfig() *CreateRuleRequestRuleActionsInsertHeaderConfig {
	return s.InsertHeaderConfig
}

func (s *CreateRuleRequestRuleActions) GetOrder() *int32 {
	return s.Order
}

func (s *CreateRuleRequestRuleActions) GetRedirectConfig() *CreateRuleRequestRuleActionsRedirectConfig {
	return s.RedirectConfig
}

func (s *CreateRuleRequestRuleActions) GetRemoveHeaderConfig() *CreateRuleRequestRuleActionsRemoveHeaderConfig {
	return s.RemoveHeaderConfig
}

func (s *CreateRuleRequestRuleActions) GetRewriteConfig() *CreateRuleRequestRuleActionsRewriteConfig {
	return s.RewriteConfig
}

func (s *CreateRuleRequestRuleActions) GetTrafficLimitConfig() *CreateRuleRequestRuleActionsTrafficLimitConfig {
	return s.TrafficLimitConfig
}

func (s *CreateRuleRequestRuleActions) GetTrafficMirrorConfig() *CreateRuleRequestRuleActionsTrafficMirrorConfig {
	return s.TrafficMirrorConfig
}

func (s *CreateRuleRequestRuleActions) GetType() *string {
	return s.Type
}

func (s *CreateRuleRequestRuleActions) SetCorsConfig(v *CreateRuleRequestRuleActionsCorsConfig) *CreateRuleRequestRuleActions {
	s.CorsConfig = v
	return s
}

func (s *CreateRuleRequestRuleActions) SetFixedResponseConfig(v *CreateRuleRequestRuleActionsFixedResponseConfig) *CreateRuleRequestRuleActions {
	s.FixedResponseConfig = v
	return s
}

func (s *CreateRuleRequestRuleActions) SetForwardGroupConfig(v *CreateRuleRequestRuleActionsForwardGroupConfig) *CreateRuleRequestRuleActions {
	s.ForwardGroupConfig = v
	return s
}

func (s *CreateRuleRequestRuleActions) SetInsertHeaderConfig(v *CreateRuleRequestRuleActionsInsertHeaderConfig) *CreateRuleRequestRuleActions {
	s.InsertHeaderConfig = v
	return s
}

func (s *CreateRuleRequestRuleActions) SetOrder(v int32) *CreateRuleRequestRuleActions {
	s.Order = &v
	return s
}

func (s *CreateRuleRequestRuleActions) SetRedirectConfig(v *CreateRuleRequestRuleActionsRedirectConfig) *CreateRuleRequestRuleActions {
	s.RedirectConfig = v
	return s
}

func (s *CreateRuleRequestRuleActions) SetRemoveHeaderConfig(v *CreateRuleRequestRuleActionsRemoveHeaderConfig) *CreateRuleRequestRuleActions {
	s.RemoveHeaderConfig = v
	return s
}

func (s *CreateRuleRequestRuleActions) SetRewriteConfig(v *CreateRuleRequestRuleActionsRewriteConfig) *CreateRuleRequestRuleActions {
	s.RewriteConfig = v
	return s
}

func (s *CreateRuleRequestRuleActions) SetTrafficLimitConfig(v *CreateRuleRequestRuleActionsTrafficLimitConfig) *CreateRuleRequestRuleActions {
	s.TrafficLimitConfig = v
	return s
}

func (s *CreateRuleRequestRuleActions) SetTrafficMirrorConfig(v *CreateRuleRequestRuleActionsTrafficMirrorConfig) *CreateRuleRequestRuleActions {
	s.TrafficMirrorConfig = v
	return s
}

func (s *CreateRuleRequestRuleActions) SetType(v string) *CreateRuleRequestRuleActions {
	s.Type = &v
	return s
}

func (s *CreateRuleRequestRuleActions) Validate() error {
	return dara.Validate(s)
}

type CreateRuleRequestRuleActionsCorsConfig struct {
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
	// The allowed HTTP methods of CORS requests.
	AllowMethods []*string `json:"AllowMethods,omitempty" xml:"AllowMethods,omitempty" type:"Repeated"`
	// The trusted origins of CORS requests. You can specify one or more values, or only the wildcard character (`*`).
	//
	// 	- Each value must start with `http://` or `https://`, which must be followed by a valid domain name, including top-level domain names. Example: `http://*.test.abc.example.com`.
	//
	// 	- You can specify a port in each value or leave the port empty. Valid values: **1*	- to **65535**.
	AllowOrigin []*string `json:"AllowOrigin,omitempty" xml:"AllowOrigin,omitempty" type:"Repeated"`
	// The headers that can be exposed.
	ExposeHeaders []*string `json:"ExposeHeaders,omitempty" xml:"ExposeHeaders,omitempty" type:"Repeated"`
	// The maximum cache time of dry run requests in the browser. Unit: seconds.
	//
	// Valid values: **-1*	- to **172800**.
	//
	// example:
	//
	// 1000
	MaxAge *int64 `json:"MaxAge,omitempty" xml:"MaxAge,omitempty"`
}

func (s CreateRuleRequestRuleActionsCorsConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateRuleRequestRuleActionsCorsConfig) GoString() string {
	return s.String()
}

func (s *CreateRuleRequestRuleActionsCorsConfig) GetAllowCredentials() *string {
	return s.AllowCredentials
}

func (s *CreateRuleRequestRuleActionsCorsConfig) GetAllowHeaders() []*string {
	return s.AllowHeaders
}

func (s *CreateRuleRequestRuleActionsCorsConfig) GetAllowMethods() []*string {
	return s.AllowMethods
}

func (s *CreateRuleRequestRuleActionsCorsConfig) GetAllowOrigin() []*string {
	return s.AllowOrigin
}

func (s *CreateRuleRequestRuleActionsCorsConfig) GetExposeHeaders() []*string {
	return s.ExposeHeaders
}

func (s *CreateRuleRequestRuleActionsCorsConfig) GetMaxAge() *int64 {
	return s.MaxAge
}

func (s *CreateRuleRequestRuleActionsCorsConfig) SetAllowCredentials(v string) *CreateRuleRequestRuleActionsCorsConfig {
	s.AllowCredentials = &v
	return s
}

func (s *CreateRuleRequestRuleActionsCorsConfig) SetAllowHeaders(v []*string) *CreateRuleRequestRuleActionsCorsConfig {
	s.AllowHeaders = v
	return s
}

func (s *CreateRuleRequestRuleActionsCorsConfig) SetAllowMethods(v []*string) *CreateRuleRequestRuleActionsCorsConfig {
	s.AllowMethods = v
	return s
}

func (s *CreateRuleRequestRuleActionsCorsConfig) SetAllowOrigin(v []*string) *CreateRuleRequestRuleActionsCorsConfig {
	s.AllowOrigin = v
	return s
}

func (s *CreateRuleRequestRuleActionsCorsConfig) SetExposeHeaders(v []*string) *CreateRuleRequestRuleActionsCorsConfig {
	s.ExposeHeaders = v
	return s
}

func (s *CreateRuleRequestRuleActionsCorsConfig) SetMaxAge(v int64) *CreateRuleRequestRuleActionsCorsConfig {
	s.MaxAge = &v
	return s
}

func (s *CreateRuleRequestRuleActionsCorsConfig) Validate() error {
	return dara.Validate(s)
}

type CreateRuleRequestRuleActionsFixedResponseConfig struct {
	// The content of the custom response. The content cannot exceed 1 KB in size, and can contain only ASCII characters.
	//
	// example:
	//
	// dssacav
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The format of the response.
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

func (s CreateRuleRequestRuleActionsFixedResponseConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateRuleRequestRuleActionsFixedResponseConfig) GoString() string {
	return s.String()
}

func (s *CreateRuleRequestRuleActionsFixedResponseConfig) GetContent() *string {
	return s.Content
}

func (s *CreateRuleRequestRuleActionsFixedResponseConfig) GetContentType() *string {
	return s.ContentType
}

func (s *CreateRuleRequestRuleActionsFixedResponseConfig) GetHttpCode() *string {
	return s.HttpCode
}

func (s *CreateRuleRequestRuleActionsFixedResponseConfig) SetContent(v string) *CreateRuleRequestRuleActionsFixedResponseConfig {
	s.Content = &v
	return s
}

func (s *CreateRuleRequestRuleActionsFixedResponseConfig) SetContentType(v string) *CreateRuleRequestRuleActionsFixedResponseConfig {
	s.ContentType = &v
	return s
}

func (s *CreateRuleRequestRuleActionsFixedResponseConfig) SetHttpCode(v string) *CreateRuleRequestRuleActionsFixedResponseConfig {
	s.HttpCode = &v
	return s
}

func (s *CreateRuleRequestRuleActionsFixedResponseConfig) Validate() error {
	return dara.Validate(s)
}

type CreateRuleRequestRuleActionsForwardGroupConfig struct {
	// The configuration of session persistence for the server groups.
	ServerGroupStickySession *CreateRuleRequestRuleActionsForwardGroupConfigServerGroupStickySession `json:"ServerGroupStickySession,omitempty" xml:"ServerGroupStickySession,omitempty" type:"Struct"`
	// The server groups to which requests are forwarded. You can specify at most five server groups for each forwarding rule.
	ServerGroupTuples []*CreateRuleRequestRuleActionsForwardGroupConfigServerGroupTuples `json:"ServerGroupTuples,omitempty" xml:"ServerGroupTuples,omitempty" type:"Repeated"`
}

func (s CreateRuleRequestRuleActionsForwardGroupConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateRuleRequestRuleActionsForwardGroupConfig) GoString() string {
	return s.String()
}

func (s *CreateRuleRequestRuleActionsForwardGroupConfig) GetServerGroupStickySession() *CreateRuleRequestRuleActionsForwardGroupConfigServerGroupStickySession {
	return s.ServerGroupStickySession
}

func (s *CreateRuleRequestRuleActionsForwardGroupConfig) GetServerGroupTuples() []*CreateRuleRequestRuleActionsForwardGroupConfigServerGroupTuples {
	return s.ServerGroupTuples
}

func (s *CreateRuleRequestRuleActionsForwardGroupConfig) SetServerGroupStickySession(v *CreateRuleRequestRuleActionsForwardGroupConfigServerGroupStickySession) *CreateRuleRequestRuleActionsForwardGroupConfig {
	s.ServerGroupStickySession = v
	return s
}

func (s *CreateRuleRequestRuleActionsForwardGroupConfig) SetServerGroupTuples(v []*CreateRuleRequestRuleActionsForwardGroupConfigServerGroupTuples) *CreateRuleRequestRuleActionsForwardGroupConfig {
	s.ServerGroupTuples = v
	return s
}

func (s *CreateRuleRequestRuleActionsForwardGroupConfig) Validate() error {
	return dara.Validate(s)
}

type CreateRuleRequestRuleActionsForwardGroupConfigServerGroupStickySession struct {
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
	// The timeout period of sessions. Unit: seconds. Valid values: **1*	- to **86400**. Default value: **1000**.
	//
	// example:
	//
	// 100
	Timeout *int32 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
}

func (s CreateRuleRequestRuleActionsForwardGroupConfigServerGroupStickySession) String() string {
	return dara.Prettify(s)
}

func (s CreateRuleRequestRuleActionsForwardGroupConfigServerGroupStickySession) GoString() string {
	return s.String()
}

func (s *CreateRuleRequestRuleActionsForwardGroupConfigServerGroupStickySession) GetEnabled() *bool {
	return s.Enabled
}

func (s *CreateRuleRequestRuleActionsForwardGroupConfigServerGroupStickySession) GetTimeout() *int32 {
	return s.Timeout
}

func (s *CreateRuleRequestRuleActionsForwardGroupConfigServerGroupStickySession) SetEnabled(v bool) *CreateRuleRequestRuleActionsForwardGroupConfigServerGroupStickySession {
	s.Enabled = &v
	return s
}

func (s *CreateRuleRequestRuleActionsForwardGroupConfigServerGroupStickySession) SetTimeout(v int32) *CreateRuleRequestRuleActionsForwardGroupConfigServerGroupStickySession {
	s.Timeout = &v
	return s
}

func (s *CreateRuleRequestRuleActionsForwardGroupConfigServerGroupStickySession) Validate() error {
	return dara.Validate(s)
}

type CreateRuleRequestRuleActionsForwardGroupConfigServerGroupTuples struct {
	// The server group to which requests are forwarded.
	//
	// example:
	//
	// sgp-k86c1ov501id6p****
	ServerGroupId *string `json:"ServerGroupId,omitempty" xml:"ServerGroupId,omitempty"`
	// The weight of the server group. A larger value specifies a higher weight. A server group with a higher weight receives more requests. Valid values: **0*	- to **100**.
	//
	// 	- If the number of destination server groups is 1, the default weight of the server group is **100**, unless you specify a weight.
	//
	// 	- If the number of destination server groups is larger than 1, you must specify a weight.
	//
	// example:
	//
	// 100
	Weight *int32 `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s CreateRuleRequestRuleActionsForwardGroupConfigServerGroupTuples) String() string {
	return dara.Prettify(s)
}

func (s CreateRuleRequestRuleActionsForwardGroupConfigServerGroupTuples) GoString() string {
	return s.String()
}

func (s *CreateRuleRequestRuleActionsForwardGroupConfigServerGroupTuples) GetServerGroupId() *string {
	return s.ServerGroupId
}

func (s *CreateRuleRequestRuleActionsForwardGroupConfigServerGroupTuples) GetWeight() *int32 {
	return s.Weight
}

func (s *CreateRuleRequestRuleActionsForwardGroupConfigServerGroupTuples) SetServerGroupId(v string) *CreateRuleRequestRuleActionsForwardGroupConfigServerGroupTuples {
	s.ServerGroupId = &v
	return s
}

func (s *CreateRuleRequestRuleActionsForwardGroupConfigServerGroupTuples) SetWeight(v int32) *CreateRuleRequestRuleActionsForwardGroupConfigServerGroupTuples {
	s.Weight = &v
	return s
}

func (s *CreateRuleRequestRuleActionsForwardGroupConfigServerGroupTuples) Validate() error {
	return dara.Validate(s)
}

type CreateRuleRequestRuleActionsInsertHeaderConfig struct {
	// The key of the header. The header key must be 1 to 40 characters in length, and can contain letters, digits, underscores (_), and hyphens (-). The header keys specified by **InsertHeaderConfig*	- must be unique.
	//
	// >  You cannot specify the following header keys: `slb-id`, `slb-ip`, `x-forwarded-for`, `x-forwarded-proto`, `x-forwarded-eip`, `x-forwarded-port`, `x-forwarded-client-srcport`, `connection`, `upgrade`, `content-length`, `transfer-encoding`, `keep-alive`, `te`, `host`, `cookie`, `remoteip`, `authority`, and `x-forwarded-host`. The header keys are case-insensitive.
	//
	// example:
	//
	// key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the header to be inserted.
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
	//     	- **SLBPort**: the listener port.
	//
	// 	- If **ValueType*	- is set to **UserDefined**, a custom header value is supported. The header value must be 1 to 128 characters in length, and can contain printable characters whose ASCII values are `greater than or equal to 32 and lower than 127`. You can use asterisks (\\*) and question marks (?) as wildcard characters. `Quotation marks (")` are not supported. The header value cannot start or end with a space character, or end with a backslash (`\\`).
	//
	// 	- If **ValueType*	- is set to **ReferenceHeader**, you can reference a value from request headers. The value must be 1 to 128 characters in length, and can contain lowercase letters, digits, hyphens (-), and underscores (_).
	//
	// example:
	//
	// UserDefined
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
	// The type of the header. Valid values:
	//
	// 	- **UserDefined**: a custom header value.
	//
	// 	- **ReferenceHeader**: a header value that is referenced from one of the request headers.
	//
	// 	- **SystemDefined**: a header predefined by the system.
	//
	// example:
	//
	// UserDefined
	ValueType *string `json:"ValueType,omitempty" xml:"ValueType,omitempty"`
}

func (s CreateRuleRequestRuleActionsInsertHeaderConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateRuleRequestRuleActionsInsertHeaderConfig) GoString() string {
	return s.String()
}

func (s *CreateRuleRequestRuleActionsInsertHeaderConfig) GetKey() *string {
	return s.Key
}

func (s *CreateRuleRequestRuleActionsInsertHeaderConfig) GetValue() *string {
	return s.Value
}

func (s *CreateRuleRequestRuleActionsInsertHeaderConfig) GetValueType() *string {
	return s.ValueType
}

func (s *CreateRuleRequestRuleActionsInsertHeaderConfig) SetKey(v string) *CreateRuleRequestRuleActionsInsertHeaderConfig {
	s.Key = &v
	return s
}

func (s *CreateRuleRequestRuleActionsInsertHeaderConfig) SetValue(v string) *CreateRuleRequestRuleActionsInsertHeaderConfig {
	s.Value = &v
	return s
}

func (s *CreateRuleRequestRuleActionsInsertHeaderConfig) SetValueType(v string) *CreateRuleRequestRuleActionsInsertHeaderConfig {
	s.ValueType = &v
	return s
}

func (s *CreateRuleRequestRuleActionsInsertHeaderConfig) Validate() error {
	return dara.Validate(s)
}

type CreateRuleRequestRuleActionsRedirectConfig struct {
	// The hostname to which requests are forwarded. Valid values:
	//
	// 	- **${host}*	- (default): If you set the value to ${host}, you cannot append other characters.
	//
	// 	- If you want to specify a custom value, make sure that the following requirements are met:
	//
	//     	- The hostname must be 3 to 128 characters in length, and can contain lowercase letters, digits, and the following characters: - . \\	- = ~ _ + \\ ^ ! $ & | ( ) [ ] ?.
	//
	//     	- The hostname must contain at least one period (.) but cannot start or end with a period (.).
	//
	//     	- The rightmost domain label can contain only letters and wildcard characters, and cannot contain digits or hyphens (-). The leftmost `domain label` can be an asterisk (\\*).
	//
	//     	- The domain labels cannot start or end with hyphens (-).
	//
	//     	- You can use asterisks (\\*) and question marks (?) anywhere in a domain label as wildcard characters.
	//
	// example:
	//
	// ${host}
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// The HTTP status code that indicates the redirect type. Valid values: **301**, **302**, **303**, **307**, and **308**.
	//
	// example:
	//
	// 301
	HttpCode *string `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	// The URL to which requests are redirected. Valid values:
	//
	// 	- Default value: **${path}**. **${host}**, **${protocol}**, and **${port}*	- are also supported. Each variable can be specified only once. You can specify one or more of the preceding variables in each request. You can also combine them with a custom value.
	//
	// 	- If you want to specify a custom value, make sure that the following requirements are met:
	//
	//     	- The URL must be 1 to 128 characters in length, and is case-sensitive. You can use asterisks (\\*) and question marks (?) as wildcard characters.
	//
	//     	- The URL must start with a forward slash (/) and can contain letters, digits, and the following special characters: `$ - _ . + / & ~ @ : \\" 	- ?`. It cannot contain the following special characters: `% # ; ! ( ) [ ] ^ , \\ "`. You can use asterisks (\\*) and question marks (?) as wildcard characters.
	//
	// example:
	//
	// /test
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// The port to which requests are distributed.
	//
	// 	- **${port}*	- (default): If you set the value to ${port}, you cannot append other characters to the value.
	//
	// 	- You can also enter a port number. Valid values: **1 to 63335**.
	//
	// example:
	//
	// 10
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// The redirect protocol. Valid values:
	//
	// 	- **${protocol}*	- (default): If you set the value to ${protocol}, you cannot modify the value or append other characters.
	//
	// 	- **HTTP**
	//
	// 	- **HTTPS**
	//
	// > 	- HTTPS listeners support only HTTPS redirection.
	//
	// >	- HTTP listeners support HTTP and HTTPS redirection.
	//
	// example:
	//
	// HTTP
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// The query string to which requests are redirected.
	//
	// 	- Default value: **${query}**. **${host}**, **${protocol}**, and **${port}*	- are also supported. Each variable can be specified only once. You can specify one or more of the preceding variables in each request. You can also combine them with a custom value.
	//
	// 	- If you want to specify a custom value, make sure that the following requirements are met:
	//
	//     	- The query string must be 1 to 128 characters in length.
	//
	//     	- It can contain printable characters, excluding space characters, the special characters `# [ ] { } \\ | < > "`, and lowercase letters.
	//
	// example:
	//
	// ${query}
	Query *string `json:"Query,omitempty" xml:"Query,omitempty"`
}

func (s CreateRuleRequestRuleActionsRedirectConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateRuleRequestRuleActionsRedirectConfig) GoString() string {
	return s.String()
}

func (s *CreateRuleRequestRuleActionsRedirectConfig) GetHost() *string {
	return s.Host
}

func (s *CreateRuleRequestRuleActionsRedirectConfig) GetHttpCode() *string {
	return s.HttpCode
}

func (s *CreateRuleRequestRuleActionsRedirectConfig) GetPath() *string {
	return s.Path
}

func (s *CreateRuleRequestRuleActionsRedirectConfig) GetPort() *string {
	return s.Port
}

func (s *CreateRuleRequestRuleActionsRedirectConfig) GetProtocol() *string {
	return s.Protocol
}

func (s *CreateRuleRequestRuleActionsRedirectConfig) GetQuery() *string {
	return s.Query
}

func (s *CreateRuleRequestRuleActionsRedirectConfig) SetHost(v string) *CreateRuleRequestRuleActionsRedirectConfig {
	s.Host = &v
	return s
}

func (s *CreateRuleRequestRuleActionsRedirectConfig) SetHttpCode(v string) *CreateRuleRequestRuleActionsRedirectConfig {
	s.HttpCode = &v
	return s
}

func (s *CreateRuleRequestRuleActionsRedirectConfig) SetPath(v string) *CreateRuleRequestRuleActionsRedirectConfig {
	s.Path = &v
	return s
}

func (s *CreateRuleRequestRuleActionsRedirectConfig) SetPort(v string) *CreateRuleRequestRuleActionsRedirectConfig {
	s.Port = &v
	return s
}

func (s *CreateRuleRequestRuleActionsRedirectConfig) SetProtocol(v string) *CreateRuleRequestRuleActionsRedirectConfig {
	s.Protocol = &v
	return s
}

func (s *CreateRuleRequestRuleActionsRedirectConfig) SetQuery(v string) *CreateRuleRequestRuleActionsRedirectConfig {
	s.Query = &v
	return s
}

func (s *CreateRuleRequestRuleActionsRedirectConfig) Validate() error {
	return dara.Validate(s)
}

type CreateRuleRequestRuleActionsRemoveHeaderConfig struct {
	// The keys of the headers to be removed. The header keys must be 1 to 40 characters in length, and can contain letters, digits, underscores (_), and hyphens (-). The header keys specified in RemoveHeader must be unique.
	//
	// 	- If Direction is set to Request, the following request headers cannot be removed: `slb-id`, `slb-ip`, `x-forwarded-for`, `x-forwarded-proto`, `x-forwarded-eip`, `x-forwarded-port`, `x-forwarded-client-srcport`, `connection`, `upgrade`, `content-length`, `transfer-encoding`, `keep-alive`, `te`, `host`, `cookie`, `remoteip`, `authority`, and `x-forwarded-host`. Request headers are not case-sensitive.
	//
	// 	- If Direction is set to Response, the following response header keys are not supported: `connection`, `upgrade`, `content-length`, and `transfer-encoding`. The header keys are not case-sensitive.
	//
	// example:
	//
	// test
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
}

func (s CreateRuleRequestRuleActionsRemoveHeaderConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateRuleRequestRuleActionsRemoveHeaderConfig) GoString() string {
	return s.String()
}

func (s *CreateRuleRequestRuleActionsRemoveHeaderConfig) GetKey() *string {
	return s.Key
}

func (s *CreateRuleRequestRuleActionsRemoveHeaderConfig) SetKey(v string) *CreateRuleRequestRuleActionsRemoveHeaderConfig {
	s.Key = &v
	return s
}

func (s *CreateRuleRequestRuleActionsRemoveHeaderConfig) Validate() error {
	return dara.Validate(s)
}

type CreateRuleRequestRuleActionsRewriteConfig struct {
	// The hostname to which requests are redirected. Valid values:
	//
	// 	- **${host}*	- (default): If you set the value to ${host}, you cannot append other characters.
	//
	// 	- If you want to specify a custom value, make sure that the following requirements are met:
	//
	//     	- The hostname must be 3 to 128 characters in length, and can contain lowercase letters, digits, and the following characters: - . \\	- = ~ _ + \\ ^ ! $ & | ( ) [ ] ?.
	//
	//     	- The hostname must contain at least one period (.) but cannot start or end with a period (.).
	//
	//     	- The rightmost domain label can contain only letters and wildcard characters, and cannot contain digits or hyphens (-). The leftmost `domain label` can be an asterisk (\\*).
	//
	//     	- The domain labels cannot start or end with hyphens (-). You can use asterisks (\\*) and question marks (?) anywhere in a domain label as wildcard characters.
	//
	// example:
	//
	// www.example.com
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// The URL to which requests are redirected. Valid values:
	//
	// 	- Default value: **${path}**. **${host}**, **${protocol}**, and **${port}*	- are also supported. Each variable can be specified only once. You can specify one or more of the preceding variables in each request. You can also combine them with a custom value.
	//
	// 	- If you want to specify a custom value, make sure that the following requirements are met:
	//
	//     	- The URL must be 1 to 128 characters in length, and is case-sensitive. You can use asterisks (\\*) and question marks (?) as wildcard characters.
	//
	//     	- The URL must start with a forward slash (/) and can contain letters, digits, and the following special characters: `$ - _ . + / & ~ @ : \\" 	- ?`. It cannot contain the following special characters: `% # ; ! ( ) [ ] ^ , \\ "`. You can use asterisks (\\*) and question marks (?) as wildcard characters.
	//
	// example:
	//
	// /tsdf
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// The query string of the URL to which requests are distributed.
	//
	// 	- Default value: **${query}**. **${host}**, **${protocol}**, and **${port}*	- are also supported. Each variable can be specified only once. The preceding variables can be used at the same time or combined with a custom value.
	//
	// 	- If you want to specify a custom value, make sure that the following requirements are met:
	//
	//     	- The path must be 1 to 128 characters in length.
	//
	//     	- It can contain printable characters, excluding space characters, the special characters `# [ ] { } \\ | < > "` and lowercase letters.
	//
	// example:
	//
	// ${query}
	Query *string `json:"Query,omitempty" xml:"Query,omitempty"`
}

func (s CreateRuleRequestRuleActionsRewriteConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateRuleRequestRuleActionsRewriteConfig) GoString() string {
	return s.String()
}

func (s *CreateRuleRequestRuleActionsRewriteConfig) GetHost() *string {
	return s.Host
}

func (s *CreateRuleRequestRuleActionsRewriteConfig) GetPath() *string {
	return s.Path
}

func (s *CreateRuleRequestRuleActionsRewriteConfig) GetQuery() *string {
	return s.Query
}

func (s *CreateRuleRequestRuleActionsRewriteConfig) SetHost(v string) *CreateRuleRequestRuleActionsRewriteConfig {
	s.Host = &v
	return s
}

func (s *CreateRuleRequestRuleActionsRewriteConfig) SetPath(v string) *CreateRuleRequestRuleActionsRewriteConfig {
	s.Path = &v
	return s
}

func (s *CreateRuleRequestRuleActionsRewriteConfig) SetQuery(v string) *CreateRuleRequestRuleActionsRewriteConfig {
	s.Query = &v
	return s
}

func (s *CreateRuleRequestRuleActionsRewriteConfig) Validate() error {
	return dara.Validate(s)
}

type CreateRuleRequestRuleActionsTrafficLimitConfig struct {
	// The number of requests per IP address. Value values: **1 to 1000000**.
	//
	// >  If both the **QPS*	- and **PerIpQps*	- parameters are specified, the value of the **QPS*	- parameter is smaller than the value of the PerIpQps parameter.
	//
	// example:
	//
	// 80
	PerIpQps *int32 `json:"PerIpQps,omitempty" xml:"PerIpQps,omitempty"`
	// The queries per second (QPS). Valid values: **1 to 1000000**.
	//
	// example:
	//
	// 100
	QPS *int32 `json:"QPS,omitempty" xml:"QPS,omitempty"`
}

func (s CreateRuleRequestRuleActionsTrafficLimitConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateRuleRequestRuleActionsTrafficLimitConfig) GoString() string {
	return s.String()
}

func (s *CreateRuleRequestRuleActionsTrafficLimitConfig) GetPerIpQps() *int32 {
	return s.PerIpQps
}

func (s *CreateRuleRequestRuleActionsTrafficLimitConfig) GetQPS() *int32 {
	return s.QPS
}

func (s *CreateRuleRequestRuleActionsTrafficLimitConfig) SetPerIpQps(v int32) *CreateRuleRequestRuleActionsTrafficLimitConfig {
	s.PerIpQps = &v
	return s
}

func (s *CreateRuleRequestRuleActionsTrafficLimitConfig) SetQPS(v int32) *CreateRuleRequestRuleActionsTrafficLimitConfig {
	s.QPS = &v
	return s
}

func (s *CreateRuleRequestRuleActionsTrafficLimitConfig) Validate() error {
	return dara.Validate(s)
}

type CreateRuleRequestRuleActionsTrafficMirrorConfig struct {
	// The configuration of the server group to which traffic is mirrored.
	MirrorGroupConfig *CreateRuleRequestRuleActionsTrafficMirrorConfigMirrorGroupConfig `json:"MirrorGroupConfig,omitempty" xml:"MirrorGroupConfig,omitempty" type:"Struct"`
	// The type of destination to which network traffic is mirrored. Valid values:
	//
	// 	- **ForwardGroupMirror**: a server group.
	//
	// example:
	//
	// ForwardGroupMirror
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
}

func (s CreateRuleRequestRuleActionsTrafficMirrorConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateRuleRequestRuleActionsTrafficMirrorConfig) GoString() string {
	return s.String()
}

func (s *CreateRuleRequestRuleActionsTrafficMirrorConfig) GetMirrorGroupConfig() *CreateRuleRequestRuleActionsTrafficMirrorConfigMirrorGroupConfig {
	return s.MirrorGroupConfig
}

func (s *CreateRuleRequestRuleActionsTrafficMirrorConfig) GetTargetType() *string {
	return s.TargetType
}

func (s *CreateRuleRequestRuleActionsTrafficMirrorConfig) SetMirrorGroupConfig(v *CreateRuleRequestRuleActionsTrafficMirrorConfigMirrorGroupConfig) *CreateRuleRequestRuleActionsTrafficMirrorConfig {
	s.MirrorGroupConfig = v
	return s
}

func (s *CreateRuleRequestRuleActionsTrafficMirrorConfig) SetTargetType(v string) *CreateRuleRequestRuleActionsTrafficMirrorConfig {
	s.TargetType = &v
	return s
}

func (s *CreateRuleRequestRuleActionsTrafficMirrorConfig) Validate() error {
	return dara.Validate(s)
}

type CreateRuleRequestRuleActionsTrafficMirrorConfigMirrorGroupConfig struct {
	// The configuration of the server group to which traffic is mirrored.
	ServerGroupTuples []*CreateRuleRequestRuleActionsTrafficMirrorConfigMirrorGroupConfigServerGroupTuples `json:"ServerGroupTuples,omitempty" xml:"ServerGroupTuples,omitempty" type:"Repeated"`
}

func (s CreateRuleRequestRuleActionsTrafficMirrorConfigMirrorGroupConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateRuleRequestRuleActionsTrafficMirrorConfigMirrorGroupConfig) GoString() string {
	return s.String()
}

func (s *CreateRuleRequestRuleActionsTrafficMirrorConfigMirrorGroupConfig) GetServerGroupTuples() []*CreateRuleRequestRuleActionsTrafficMirrorConfigMirrorGroupConfigServerGroupTuples {
	return s.ServerGroupTuples
}

func (s *CreateRuleRequestRuleActionsTrafficMirrorConfigMirrorGroupConfig) SetServerGroupTuples(v []*CreateRuleRequestRuleActionsTrafficMirrorConfigMirrorGroupConfigServerGroupTuples) *CreateRuleRequestRuleActionsTrafficMirrorConfigMirrorGroupConfig {
	s.ServerGroupTuples = v
	return s
}

func (s *CreateRuleRequestRuleActionsTrafficMirrorConfigMirrorGroupConfig) Validate() error {
	return dara.Validate(s)
}

type CreateRuleRequestRuleActionsTrafficMirrorConfigMirrorGroupConfigServerGroupTuples struct {
	// The server group ID.
	//
	// example:
	//
	// sgp-00mkgijak0w4qgz9****
	ServerGroupId *string `json:"ServerGroupId,omitempty" xml:"ServerGroupId,omitempty"`
}

func (s CreateRuleRequestRuleActionsTrafficMirrorConfigMirrorGroupConfigServerGroupTuples) String() string {
	return dara.Prettify(s)
}

func (s CreateRuleRequestRuleActionsTrafficMirrorConfigMirrorGroupConfigServerGroupTuples) GoString() string {
	return s.String()
}

func (s *CreateRuleRequestRuleActionsTrafficMirrorConfigMirrorGroupConfigServerGroupTuples) GetServerGroupId() *string {
	return s.ServerGroupId
}

func (s *CreateRuleRequestRuleActionsTrafficMirrorConfigMirrorGroupConfigServerGroupTuples) SetServerGroupId(v string) *CreateRuleRequestRuleActionsTrafficMirrorConfigMirrorGroupConfigServerGroupTuples {
	s.ServerGroupId = &v
	return s
}

func (s *CreateRuleRequestRuleActionsTrafficMirrorConfigMirrorGroupConfigServerGroupTuples) Validate() error {
	return dara.Validate(s)
}

type CreateRuleRequestRuleConditions struct {
	// The key-value pairs of the cookie.
	CookieConfig *CreateRuleRequestRuleConditionsCookieConfig `json:"CookieConfig,omitempty" xml:"CookieConfig,omitempty" type:"Struct"`
	// The configuration of the header.
	HeaderConfig *CreateRuleRequestRuleConditionsHeaderConfig `json:"HeaderConfig,omitempty" xml:"HeaderConfig,omitempty" type:"Struct"`
	// The configuration of the host.
	HostConfig *CreateRuleRequestRuleConditionsHostConfig `json:"HostConfig,omitempty" xml:"HostConfig,omitempty" type:"Struct"`
	// The configurations of the request method.
	MethodConfig *CreateRuleRequestRuleConditionsMethodConfig `json:"MethodConfig,omitempty" xml:"MethodConfig,omitempty" type:"Struct"`
	// The configurations of the URL to which requests are forwarded.
	PathConfig *CreateRuleRequestRuleConditionsPathConfig `json:"PathConfig,omitempty" xml:"PathConfig,omitempty" type:"Struct"`
	// The configuration of the query strings.
	QueryStringConfig *CreateRuleRequestRuleConditionsQueryStringConfig `json:"QueryStringConfig,omitempty" xml:"QueryStringConfig,omitempty" type:"Struct"`
	// The configuration of headers.
	ResponseHeaderConfig *CreateRuleRequestRuleConditionsResponseHeaderConfig `json:"ResponseHeaderConfig,omitempty" xml:"ResponseHeaderConfig,omitempty" type:"Struct"`
	// The configuration of the response status codes.
	ResponseStatusCodeConfig *CreateRuleRequestRuleConditionsResponseStatusCodeConfig `json:"ResponseStatusCodeConfig,omitempty" xml:"ResponseStatusCodeConfig,omitempty" type:"Struct"`
	// Configuration of traffic matching based on source IP addresses. This parameter is required and valid when **Type*	- is set to **SourceIP**.
	SourceIpConfig *CreateRuleRequestRuleConditionsSourceIpConfig `json:"SourceIpConfig,omitempty" xml:"SourceIpConfig,omitempty" type:"Struct"`
	// The type of forwarding rule. Valid values:
	//
	// 	- **Host**: Requests are distributed based on hosts.
	//
	// 	- **Path**: Requests are distributed based on paths.
	//
	// 	- **Header**: Requests are distributed based on HTTP headers.
	//
	// 	- **QueryString**: Requests are distributed based on query strings.
	//
	// 	- **Method**: Requests are distributed based on request methods.
	//
	// 	- **Cookie**: Requests are distributed based on cookies.
	//
	// 	- **SourceIp**: Requests are forwarded based on source IP addresses.
	//
	// 	- **ResponseHeader**: Requests are forwarded based on HTTP response headers.
	//
	// 	- **ResponseStatusCode**: Requests are forwarded based on response status codes.
	//
	// This parameter is required.
	//
	// example:
	//
	// Host
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateRuleRequestRuleConditions) String() string {
	return dara.Prettify(s)
}

func (s CreateRuleRequestRuleConditions) GoString() string {
	return s.String()
}

func (s *CreateRuleRequestRuleConditions) GetCookieConfig() *CreateRuleRequestRuleConditionsCookieConfig {
	return s.CookieConfig
}

func (s *CreateRuleRequestRuleConditions) GetHeaderConfig() *CreateRuleRequestRuleConditionsHeaderConfig {
	return s.HeaderConfig
}

func (s *CreateRuleRequestRuleConditions) GetHostConfig() *CreateRuleRequestRuleConditionsHostConfig {
	return s.HostConfig
}

func (s *CreateRuleRequestRuleConditions) GetMethodConfig() *CreateRuleRequestRuleConditionsMethodConfig {
	return s.MethodConfig
}

func (s *CreateRuleRequestRuleConditions) GetPathConfig() *CreateRuleRequestRuleConditionsPathConfig {
	return s.PathConfig
}

func (s *CreateRuleRequestRuleConditions) GetQueryStringConfig() *CreateRuleRequestRuleConditionsQueryStringConfig {
	return s.QueryStringConfig
}

func (s *CreateRuleRequestRuleConditions) GetResponseHeaderConfig() *CreateRuleRequestRuleConditionsResponseHeaderConfig {
	return s.ResponseHeaderConfig
}

func (s *CreateRuleRequestRuleConditions) GetResponseStatusCodeConfig() *CreateRuleRequestRuleConditionsResponseStatusCodeConfig {
	return s.ResponseStatusCodeConfig
}

func (s *CreateRuleRequestRuleConditions) GetSourceIpConfig() *CreateRuleRequestRuleConditionsSourceIpConfig {
	return s.SourceIpConfig
}

func (s *CreateRuleRequestRuleConditions) GetType() *string {
	return s.Type
}

func (s *CreateRuleRequestRuleConditions) SetCookieConfig(v *CreateRuleRequestRuleConditionsCookieConfig) *CreateRuleRequestRuleConditions {
	s.CookieConfig = v
	return s
}

func (s *CreateRuleRequestRuleConditions) SetHeaderConfig(v *CreateRuleRequestRuleConditionsHeaderConfig) *CreateRuleRequestRuleConditions {
	s.HeaderConfig = v
	return s
}

func (s *CreateRuleRequestRuleConditions) SetHostConfig(v *CreateRuleRequestRuleConditionsHostConfig) *CreateRuleRequestRuleConditions {
	s.HostConfig = v
	return s
}

func (s *CreateRuleRequestRuleConditions) SetMethodConfig(v *CreateRuleRequestRuleConditionsMethodConfig) *CreateRuleRequestRuleConditions {
	s.MethodConfig = v
	return s
}

func (s *CreateRuleRequestRuleConditions) SetPathConfig(v *CreateRuleRequestRuleConditionsPathConfig) *CreateRuleRequestRuleConditions {
	s.PathConfig = v
	return s
}

func (s *CreateRuleRequestRuleConditions) SetQueryStringConfig(v *CreateRuleRequestRuleConditionsQueryStringConfig) *CreateRuleRequestRuleConditions {
	s.QueryStringConfig = v
	return s
}

func (s *CreateRuleRequestRuleConditions) SetResponseHeaderConfig(v *CreateRuleRequestRuleConditionsResponseHeaderConfig) *CreateRuleRequestRuleConditions {
	s.ResponseHeaderConfig = v
	return s
}

func (s *CreateRuleRequestRuleConditions) SetResponseStatusCodeConfig(v *CreateRuleRequestRuleConditionsResponseStatusCodeConfig) *CreateRuleRequestRuleConditions {
	s.ResponseStatusCodeConfig = v
	return s
}

func (s *CreateRuleRequestRuleConditions) SetSourceIpConfig(v *CreateRuleRequestRuleConditionsSourceIpConfig) *CreateRuleRequestRuleConditions {
	s.SourceIpConfig = v
	return s
}

func (s *CreateRuleRequestRuleConditions) SetType(v string) *CreateRuleRequestRuleConditions {
	s.Type = &v
	return s
}

func (s *CreateRuleRequestRuleConditions) Validate() error {
	return dara.Validate(s)
}

type CreateRuleRequestRuleConditionsCookieConfig struct {
	// The cookie values.
	Values []*CreateRuleRequestRuleConditionsCookieConfigValues `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s CreateRuleRequestRuleConditionsCookieConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateRuleRequestRuleConditionsCookieConfig) GoString() string {
	return s.String()
}

func (s *CreateRuleRequestRuleConditionsCookieConfig) GetValues() []*CreateRuleRequestRuleConditionsCookieConfigValues {
	return s.Values
}

func (s *CreateRuleRequestRuleConditionsCookieConfig) SetValues(v []*CreateRuleRequestRuleConditionsCookieConfigValues) *CreateRuleRequestRuleConditionsCookieConfig {
	s.Values = v
	return s
}

func (s *CreateRuleRequestRuleConditionsCookieConfig) Validate() error {
	return dara.Validate(s)
}

type CreateRuleRequestRuleConditionsCookieConfigValues struct {
	// The cookie key.
	//
	// 	- The cookie key must be 1 to 100 characters in length.
	//
	// 	- You can use asterisks (\\*) and question marks (?) as wildcard characters.
	//
	// 	- The value can contain printable characters, excluding uppercase letters, space characters, and the following special characters: `; # [ ] { } \\ | < > & "`.
	//
	// example:
	//
	// test
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The cookie value.
	//
	// 	- The cookie value must be 1 to 100 characters in length.
	//
	// 	- You can use asterisks (\\*) and question marks (?) as wildcard characters.
	//
	// 	- The value can contain printable characters, excluding uppercase letters, space characters, and the following special characters: `; # [ ] { } \\ | < > & "`.
	//
	// example:
	//
	// test
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateRuleRequestRuleConditionsCookieConfigValues) String() string {
	return dara.Prettify(s)
}

func (s CreateRuleRequestRuleConditionsCookieConfigValues) GoString() string {
	return s.String()
}

func (s *CreateRuleRequestRuleConditionsCookieConfigValues) GetKey() *string {
	return s.Key
}

func (s *CreateRuleRequestRuleConditionsCookieConfigValues) GetValue() *string {
	return s.Value
}

func (s *CreateRuleRequestRuleConditionsCookieConfigValues) SetKey(v string) *CreateRuleRequestRuleConditionsCookieConfigValues {
	s.Key = &v
	return s
}

func (s *CreateRuleRequestRuleConditionsCookieConfigValues) SetValue(v string) *CreateRuleRequestRuleConditionsCookieConfigValues {
	s.Value = &v
	return s
}

func (s *CreateRuleRequestRuleConditionsCookieConfigValues) Validate() error {
	return dara.Validate(s)
}

type CreateRuleRequestRuleConditionsHeaderConfig struct {
	// The header key.
	//
	// 	- The header key must be 1 to 40 characters in length.
	//
	// 	- The key can contain letters, digits, hyphens (-), and underscores (_).
	//
	// 	- Cookie and Host are not supported.
	//
	// example:
	//
	// Port
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The header values.
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s CreateRuleRequestRuleConditionsHeaderConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateRuleRequestRuleConditionsHeaderConfig) GoString() string {
	return s.String()
}

func (s *CreateRuleRequestRuleConditionsHeaderConfig) GetKey() *string {
	return s.Key
}

func (s *CreateRuleRequestRuleConditionsHeaderConfig) GetValues() []*string {
	return s.Values
}

func (s *CreateRuleRequestRuleConditionsHeaderConfig) SetKey(v string) *CreateRuleRequestRuleConditionsHeaderConfig {
	s.Key = &v
	return s
}

func (s *CreateRuleRequestRuleConditionsHeaderConfig) SetValues(v []*string) *CreateRuleRequestRuleConditionsHeaderConfig {
	s.Values = v
	return s
}

func (s *CreateRuleRequestRuleConditionsHeaderConfig) Validate() error {
	return dara.Validate(s)
}

type CreateRuleRequestRuleConditionsHostConfig struct {
	// The hostnames.
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s CreateRuleRequestRuleConditionsHostConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateRuleRequestRuleConditionsHostConfig) GoString() string {
	return s.String()
}

func (s *CreateRuleRequestRuleConditionsHostConfig) GetValues() []*string {
	return s.Values
}

func (s *CreateRuleRequestRuleConditionsHostConfig) SetValues(v []*string) *CreateRuleRequestRuleConditionsHostConfig {
	s.Values = v
	return s
}

func (s *CreateRuleRequestRuleConditionsHostConfig) Validate() error {
	return dara.Validate(s)
}

type CreateRuleRequestRuleConditionsMethodConfig struct {
	// The request methods.
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s CreateRuleRequestRuleConditionsMethodConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateRuleRequestRuleConditionsMethodConfig) GoString() string {
	return s.String()
}

func (s *CreateRuleRequestRuleConditionsMethodConfig) GetValues() []*string {
	return s.Values
}

func (s *CreateRuleRequestRuleConditionsMethodConfig) SetValues(v []*string) *CreateRuleRequestRuleConditionsMethodConfig {
	s.Values = v
	return s
}

func (s *CreateRuleRequestRuleConditionsMethodConfig) Validate() error {
	return dara.Validate(s)
}

type CreateRuleRequestRuleConditionsPathConfig struct {
	// The forwarding URLs.
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s CreateRuleRequestRuleConditionsPathConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateRuleRequestRuleConditionsPathConfig) GoString() string {
	return s.String()
}

func (s *CreateRuleRequestRuleConditionsPathConfig) GetValues() []*string {
	return s.Values
}

func (s *CreateRuleRequestRuleConditionsPathConfig) SetValues(v []*string) *CreateRuleRequestRuleConditionsPathConfig {
	s.Values = v
	return s
}

func (s *CreateRuleRequestRuleConditionsPathConfig) Validate() error {
	return dara.Validate(s)
}

type CreateRuleRequestRuleConditionsQueryStringConfig struct {
	// The query strings.
	Values []*CreateRuleRequestRuleConditionsQueryStringConfigValues `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s CreateRuleRequestRuleConditionsQueryStringConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateRuleRequestRuleConditionsQueryStringConfig) GoString() string {
	return s.String()
}

func (s *CreateRuleRequestRuleConditionsQueryStringConfig) GetValues() []*CreateRuleRequestRuleConditionsQueryStringConfigValues {
	return s.Values
}

func (s *CreateRuleRequestRuleConditionsQueryStringConfig) SetValues(v []*CreateRuleRequestRuleConditionsQueryStringConfigValues) *CreateRuleRequestRuleConditionsQueryStringConfig {
	s.Values = v
	return s
}

func (s *CreateRuleRequestRuleConditionsQueryStringConfig) Validate() error {
	return dara.Validate(s)
}

type CreateRuleRequestRuleConditionsQueryStringConfigValues struct {
	// They key of the query string.
	//
	// 	- The key must be 1 to 100 characters in length.
	//
	// 	- You can use asterisks (\\*) and question marks (?) as wildcard characters. It can contain printable characters, excluding uppercase letters, space characters, and the following special characters: `# [ ] { } \\ | < > & "`.
	//
	// example:
	//
	// test
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the query string.
	//
	// 	- The value must be 1 to 128 characters in length.
	//
	// 	- It can contain printable characters, but cannot contain uppercase letters, space characters, or the following special characters: `# [ ] { } \\ | < > &`. You can use asterisks (\\*) and question marks (?) as wildcard characters.
	//
	// example:
	//
	// test
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateRuleRequestRuleConditionsQueryStringConfigValues) String() string {
	return dara.Prettify(s)
}

func (s CreateRuleRequestRuleConditionsQueryStringConfigValues) GoString() string {
	return s.String()
}

func (s *CreateRuleRequestRuleConditionsQueryStringConfigValues) GetKey() *string {
	return s.Key
}

func (s *CreateRuleRequestRuleConditionsQueryStringConfigValues) GetValue() *string {
	return s.Value
}

func (s *CreateRuleRequestRuleConditionsQueryStringConfigValues) SetKey(v string) *CreateRuleRequestRuleConditionsQueryStringConfigValues {
	s.Key = &v
	return s
}

func (s *CreateRuleRequestRuleConditionsQueryStringConfigValues) SetValue(v string) *CreateRuleRequestRuleConditionsQueryStringConfigValues {
	s.Value = &v
	return s
}

func (s *CreateRuleRequestRuleConditionsQueryStringConfigValues) Validate() error {
	return dara.Validate(s)
}

type CreateRuleRequestRuleConditionsResponseHeaderConfig struct {
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

func (s CreateRuleRequestRuleConditionsResponseHeaderConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateRuleRequestRuleConditionsResponseHeaderConfig) GoString() string {
	return s.String()
}

func (s *CreateRuleRequestRuleConditionsResponseHeaderConfig) GetKey() *string {
	return s.Key
}

func (s *CreateRuleRequestRuleConditionsResponseHeaderConfig) GetValues() []*string {
	return s.Values
}

func (s *CreateRuleRequestRuleConditionsResponseHeaderConfig) SetKey(v string) *CreateRuleRequestRuleConditionsResponseHeaderConfig {
	s.Key = &v
	return s
}

func (s *CreateRuleRequestRuleConditionsResponseHeaderConfig) SetValues(v []*string) *CreateRuleRequestRuleConditionsResponseHeaderConfig {
	s.Values = v
	return s
}

func (s *CreateRuleRequestRuleConditionsResponseHeaderConfig) Validate() error {
	return dara.Validate(s)
}

type CreateRuleRequestRuleConditionsResponseStatusCodeConfig struct {
	// The response status codes.
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s CreateRuleRequestRuleConditionsResponseStatusCodeConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateRuleRequestRuleConditionsResponseStatusCodeConfig) GoString() string {
	return s.String()
}

func (s *CreateRuleRequestRuleConditionsResponseStatusCodeConfig) GetValues() []*string {
	return s.Values
}

func (s *CreateRuleRequestRuleConditionsResponseStatusCodeConfig) SetValues(v []*string) *CreateRuleRequestRuleConditionsResponseStatusCodeConfig {
	s.Values = v
	return s
}

func (s *CreateRuleRequestRuleConditionsResponseStatusCodeConfig) Validate() error {
	return dara.Validate(s)
}

type CreateRuleRequestRuleConditionsSourceIpConfig struct {
	// The source IP addresses used for traffic matching.
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s CreateRuleRequestRuleConditionsSourceIpConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateRuleRequestRuleConditionsSourceIpConfig) GoString() string {
	return s.String()
}

func (s *CreateRuleRequestRuleConditionsSourceIpConfig) GetValues() []*string {
	return s.Values
}

func (s *CreateRuleRequestRuleConditionsSourceIpConfig) SetValues(v []*string) *CreateRuleRequestRuleConditionsSourceIpConfig {
	s.Values = v
	return s
}

func (s *CreateRuleRequestRuleConditionsSourceIpConfig) Validate() error {
	return dara.Validate(s)
}

type CreateRuleRequestTag struct {
	// The tag key. The tag key can be up to 128 characters in length. It cannot start with aliyun or acs: and cannot contain http:// or https://.
	//
	// example:
	//
	// env
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value can be up to 128 characters in length. It cannot start with aliyun or acs: and cannot contain http:// or https://.
	//
	// example:
	//
	// product
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateRuleRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateRuleRequestTag) GoString() string {
	return s.String()
}

func (s *CreateRuleRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateRuleRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateRuleRequestTag) SetKey(v string) *CreateRuleRequestTag {
	s.Key = &v
	return s
}

func (s *CreateRuleRequestTag) SetValue(v string) *CreateRuleRequestTag {
	s.Value = &v
	return s
}

func (s *CreateRuleRequestTag) Validate() error {
	return dara.Validate(s)
}
