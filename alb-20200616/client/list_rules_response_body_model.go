// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListRulesResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListRulesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListRulesResponseBody
	GetRequestId() *string
	SetRules(v []*ListRulesResponseBodyRules) *ListRulesResponseBody
	GetRules() []*ListRulesResponseBodyRules
	SetTotalCount(v int32) *ListRulesResponseBody
	GetTotalCount() *int32
}

type ListRulesResponseBody struct {
	// The maximum number of entries returned.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// A pagination token. It can be used in the next request to retrieve a new page of results. Valid values:
	//
	// 	- If **NextToken*	- is empty, no next page exists.
	//
	// 	- If a value is returned for **NextToken**, the value is the token that determines the start point of the next query.
	//
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// CEF72CEB-54B6-4AE8-B225-F876FF7BA984
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The details about the forwarding rule.
	Rules []*ListRulesResponseBodyRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1000
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListRulesResponseBody) GoString() string {
	return s.String()
}

func (s *ListRulesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListRulesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListRulesResponseBody) GetRules() []*ListRulesResponseBodyRules {
	return s.Rules
}

func (s *ListRulesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListRulesResponseBody) SetMaxResults(v int32) *ListRulesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListRulesResponseBody) SetNextToken(v string) *ListRulesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListRulesResponseBody) SetRequestId(v string) *ListRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRulesResponseBody) SetRules(v []*ListRulesResponseBodyRules) *ListRulesResponseBody {
	s.Rules = v
	return s
}

func (s *ListRulesResponseBody) SetTotalCount(v int32) *ListRulesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListRulesResponseBody) Validate() error {
	if s.Rules != nil {
		for _, item := range s.Rules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListRulesResponseBodyRules struct {
	// The direction to which the forwarding rule is applied. Valid values:
	//
	// 	- Request (default): The forwarding rule is applied to requests. The forwarding action is performed on packets that are forwarded from clients to ALB.
	//
	// 	- Responses: The forwarding rule is applied to responses. The forwarding action is performed on packets that are returned from backend servers to ALB.
	//
	// >  Basic ALB instances support only the Response direction.
	//
	// example:
	//
	// Request
	Direction *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	// The ID of the listener that is associated with the forwarding rule.
	//
	// example:
	//
	// lsn-i35udpz3pxsmnf****
	ListenerId *string `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	// The ID of the Application Load Balancer (ALB) instance that is associated with the forwarding rule.
	//
	// example:
	//
	// alb-x30o38azsuj0sx****
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	// The priority of the forwarding rule. Valid values: **1 to 10000**. A smaller value indicates a higher priority.
	//
	// >  The priority of each forwarding rule added to a listener must be unique.
	//
	// example:
	//
	// 1
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The action of the forwarding rule.
	RuleActions []*ListRulesResponseBodyRulesRuleActions `json:"RuleActions,omitempty" xml:"RuleActions,omitempty" type:"Repeated"`
	// The conditions of the forwarding rule.
	RuleConditions []*ListRulesResponseBodyRulesRuleConditions `json:"RuleConditions,omitempty" xml:"RuleConditions,omitempty" type:"Repeated"`
	// The ID of the forwarding rule.
	//
	// example:
	//
	// rule-bpn0kn908w4nbw****
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The name of the forwarding rule. The name must be 2 to 128 letters in length, and can contain letters, digits, periods (.), underscores (_), and hyphens (-). The name must start with a letter.
	//
	// example:
	//
	// rule-instance-test
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The status of the forwarding rule. Valid values:
	//
	// 	- **Provisioning**: The forwarding rule is being created.
	//
	// 	- **Configuring**: The forwarding rule is being modified.
	//
	// 	- **Available**: The forwarding rule is available.
	//
	// example:
	//
	// Available
	RuleStatus *string `json:"RuleStatus,omitempty" xml:"RuleStatus,omitempty"`
	// The tags.
	Tags []*ListRulesResponseBodyRulesTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s ListRulesResponseBodyRules) String() string {
	return dara.Prettify(s)
}

func (s ListRulesResponseBodyRules) GoString() string {
	return s.String()
}

func (s *ListRulesResponseBodyRules) GetDirection() *string {
	return s.Direction
}

func (s *ListRulesResponseBodyRules) GetListenerId() *string {
	return s.ListenerId
}

func (s *ListRulesResponseBodyRules) GetLoadBalancerId() *string {
	return s.LoadBalancerId
}

func (s *ListRulesResponseBodyRules) GetPriority() *int32 {
	return s.Priority
}

func (s *ListRulesResponseBodyRules) GetRuleActions() []*ListRulesResponseBodyRulesRuleActions {
	return s.RuleActions
}

func (s *ListRulesResponseBodyRules) GetRuleConditions() []*ListRulesResponseBodyRulesRuleConditions {
	return s.RuleConditions
}

func (s *ListRulesResponseBodyRules) GetRuleId() *string {
	return s.RuleId
}

func (s *ListRulesResponseBodyRules) GetRuleName() *string {
	return s.RuleName
}

func (s *ListRulesResponseBodyRules) GetRuleStatus() *string {
	return s.RuleStatus
}

func (s *ListRulesResponseBodyRules) GetTags() []*ListRulesResponseBodyRulesTags {
	return s.Tags
}

func (s *ListRulesResponseBodyRules) SetDirection(v string) *ListRulesResponseBodyRules {
	s.Direction = &v
	return s
}

func (s *ListRulesResponseBodyRules) SetListenerId(v string) *ListRulesResponseBodyRules {
	s.ListenerId = &v
	return s
}

func (s *ListRulesResponseBodyRules) SetLoadBalancerId(v string) *ListRulesResponseBodyRules {
	s.LoadBalancerId = &v
	return s
}

func (s *ListRulesResponseBodyRules) SetPriority(v int32) *ListRulesResponseBodyRules {
	s.Priority = &v
	return s
}

func (s *ListRulesResponseBodyRules) SetRuleActions(v []*ListRulesResponseBodyRulesRuleActions) *ListRulesResponseBodyRules {
	s.RuleActions = v
	return s
}

func (s *ListRulesResponseBodyRules) SetRuleConditions(v []*ListRulesResponseBodyRulesRuleConditions) *ListRulesResponseBodyRules {
	s.RuleConditions = v
	return s
}

func (s *ListRulesResponseBodyRules) SetRuleId(v string) *ListRulesResponseBodyRules {
	s.RuleId = &v
	return s
}

func (s *ListRulesResponseBodyRules) SetRuleName(v string) *ListRulesResponseBodyRules {
	s.RuleName = &v
	return s
}

func (s *ListRulesResponseBodyRules) SetRuleStatus(v string) *ListRulesResponseBodyRules {
	s.RuleStatus = &v
	return s
}

func (s *ListRulesResponseBodyRules) SetTags(v []*ListRulesResponseBodyRulesTags) *ListRulesResponseBodyRules {
	s.Tags = v
	return s
}

func (s *ListRulesResponseBodyRules) Validate() error {
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
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListRulesResponseBodyRulesRuleActions struct {
	// The CORS configuration.
	CorsConfig *ListRulesResponseBodyRulesRuleActionsCorsConfig `json:"CorsConfig,omitempty" xml:"CorsConfig,omitempty" type:"Struct"`
	// The configuration of the custom response.
	FixedResponseConfig *ListRulesResponseBodyRulesRuleActionsFixedResponseConfig `json:"FixedResponseConfig,omitempty" xml:"FixedResponseConfig,omitempty" type:"Struct"`
	// The configurations of the server groups.
	ForwardGroupConfig *ListRulesResponseBodyRulesRuleActionsForwardGroupConfig `json:"ForwardGroupConfig,omitempty" xml:"ForwardGroupConfig,omitempty" type:"Struct"`
	// The key of the header to be inserted.
	InsertHeaderConfig *ListRulesResponseBodyRulesRuleActionsInsertHeaderConfig `json:"InsertHeaderConfig,omitempty" xml:"InsertHeaderConfig,omitempty" type:"Struct"`
	// The priority of the action. Valid values: **1 to 50000**. A smaller value indicates a higher priority. The actions of a forwarding rule are applied in descending order of priority. This parameter cannot empty. The priority of each action within a forwarding rule must be unique.
	//
	// example:
	//
	// 1
	Order *int32 `json:"Order,omitempty" xml:"Order,omitempty"`
	// The configuration of the redirect action.
	RedirectConfig *ListRulesResponseBodyRulesRuleActionsRedirectConfig `json:"RedirectConfig,omitempty" xml:"RedirectConfig,omitempty" type:"Struct"`
	// The HTTP header to be removed.
	RemoveHeaderConfig *ListRulesResponseBodyRulesRuleActionsRemoveHeaderConfig `json:"RemoveHeaderConfig,omitempty" xml:"RemoveHeaderConfig,omitempty" type:"Struct"`
	// The configuration of the rewrite action.
	RewriteConfig *ListRulesResponseBodyRulesRuleActionsRewriteConfig `json:"RewriteConfig,omitempty" xml:"RewriteConfig,omitempty" type:"Struct"`
	// The configuration of traffic throttling.
	TrafficLimitConfig *ListRulesResponseBodyRulesRuleActionsTrafficLimitConfig `json:"TrafficLimitConfig,omitempty" xml:"TrafficLimitConfig,omitempty" type:"Struct"`
	// The configuration of traffic mirroring.
	TrafficMirrorConfig *ListRulesResponseBodyRulesRuleActionsTrafficMirrorConfig `json:"TrafficMirrorConfig,omitempty" xml:"TrafficMirrorConfig,omitempty" type:"Struct"`
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
	// 	- **RemoveHeaderConfig**: removes headers.
	//
	// 	- **TrafficLimitConfig**: throttles network traffic.
	//
	// 	- **TrafficMirrorConfig**: mirrors network traffic.
	//
	// 	- **CorsConfig**: forwards requests based on CORS.
	//
	// The preceding actions can be classified into two broad types:
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

func (s ListRulesResponseBodyRulesRuleActions) String() string {
	return dara.Prettify(s)
}

func (s ListRulesResponseBodyRulesRuleActions) GoString() string {
	return s.String()
}

func (s *ListRulesResponseBodyRulesRuleActions) GetCorsConfig() *ListRulesResponseBodyRulesRuleActionsCorsConfig {
	return s.CorsConfig
}

func (s *ListRulesResponseBodyRulesRuleActions) GetFixedResponseConfig() *ListRulesResponseBodyRulesRuleActionsFixedResponseConfig {
	return s.FixedResponseConfig
}

func (s *ListRulesResponseBodyRulesRuleActions) GetForwardGroupConfig() *ListRulesResponseBodyRulesRuleActionsForwardGroupConfig {
	return s.ForwardGroupConfig
}

func (s *ListRulesResponseBodyRulesRuleActions) GetInsertHeaderConfig() *ListRulesResponseBodyRulesRuleActionsInsertHeaderConfig {
	return s.InsertHeaderConfig
}

func (s *ListRulesResponseBodyRulesRuleActions) GetOrder() *int32 {
	return s.Order
}

func (s *ListRulesResponseBodyRulesRuleActions) GetRedirectConfig() *ListRulesResponseBodyRulesRuleActionsRedirectConfig {
	return s.RedirectConfig
}

func (s *ListRulesResponseBodyRulesRuleActions) GetRemoveHeaderConfig() *ListRulesResponseBodyRulesRuleActionsRemoveHeaderConfig {
	return s.RemoveHeaderConfig
}

func (s *ListRulesResponseBodyRulesRuleActions) GetRewriteConfig() *ListRulesResponseBodyRulesRuleActionsRewriteConfig {
	return s.RewriteConfig
}

func (s *ListRulesResponseBodyRulesRuleActions) GetTrafficLimitConfig() *ListRulesResponseBodyRulesRuleActionsTrafficLimitConfig {
	return s.TrafficLimitConfig
}

func (s *ListRulesResponseBodyRulesRuleActions) GetTrafficMirrorConfig() *ListRulesResponseBodyRulesRuleActionsTrafficMirrorConfig {
	return s.TrafficMirrorConfig
}

func (s *ListRulesResponseBodyRulesRuleActions) GetType() *string {
	return s.Type
}

func (s *ListRulesResponseBodyRulesRuleActions) SetCorsConfig(v *ListRulesResponseBodyRulesRuleActionsCorsConfig) *ListRulesResponseBodyRulesRuleActions {
	s.CorsConfig = v
	return s
}

func (s *ListRulesResponseBodyRulesRuleActions) SetFixedResponseConfig(v *ListRulesResponseBodyRulesRuleActionsFixedResponseConfig) *ListRulesResponseBodyRulesRuleActions {
	s.FixedResponseConfig = v
	return s
}

func (s *ListRulesResponseBodyRulesRuleActions) SetForwardGroupConfig(v *ListRulesResponseBodyRulesRuleActionsForwardGroupConfig) *ListRulesResponseBodyRulesRuleActions {
	s.ForwardGroupConfig = v
	return s
}

func (s *ListRulesResponseBodyRulesRuleActions) SetInsertHeaderConfig(v *ListRulesResponseBodyRulesRuleActionsInsertHeaderConfig) *ListRulesResponseBodyRulesRuleActions {
	s.InsertHeaderConfig = v
	return s
}

func (s *ListRulesResponseBodyRulesRuleActions) SetOrder(v int32) *ListRulesResponseBodyRulesRuleActions {
	s.Order = &v
	return s
}

func (s *ListRulesResponseBodyRulesRuleActions) SetRedirectConfig(v *ListRulesResponseBodyRulesRuleActionsRedirectConfig) *ListRulesResponseBodyRulesRuleActions {
	s.RedirectConfig = v
	return s
}

func (s *ListRulesResponseBodyRulesRuleActions) SetRemoveHeaderConfig(v *ListRulesResponseBodyRulesRuleActionsRemoveHeaderConfig) *ListRulesResponseBodyRulesRuleActions {
	s.RemoveHeaderConfig = v
	return s
}

func (s *ListRulesResponseBodyRulesRuleActions) SetRewriteConfig(v *ListRulesResponseBodyRulesRuleActionsRewriteConfig) *ListRulesResponseBodyRulesRuleActions {
	s.RewriteConfig = v
	return s
}

func (s *ListRulesResponseBodyRulesRuleActions) SetTrafficLimitConfig(v *ListRulesResponseBodyRulesRuleActionsTrafficLimitConfig) *ListRulesResponseBodyRulesRuleActions {
	s.TrafficLimitConfig = v
	return s
}

func (s *ListRulesResponseBodyRulesRuleActions) SetTrafficMirrorConfig(v *ListRulesResponseBodyRulesRuleActionsTrafficMirrorConfig) *ListRulesResponseBodyRulesRuleActions {
	s.TrafficMirrorConfig = v
	return s
}

func (s *ListRulesResponseBodyRulesRuleActions) SetType(v string) *ListRulesResponseBodyRulesRuleActions {
	s.Type = &v
	return s
}

func (s *ListRulesResponseBodyRulesRuleActions) Validate() error {
	if s.CorsConfig != nil {
		if err := s.CorsConfig.Validate(); err != nil {
			return err
		}
	}
	if s.FixedResponseConfig != nil {
		if err := s.FixedResponseConfig.Validate(); err != nil {
			return err
		}
	}
	if s.ForwardGroupConfig != nil {
		if err := s.ForwardGroupConfig.Validate(); err != nil {
			return err
		}
	}
	if s.InsertHeaderConfig != nil {
		if err := s.InsertHeaderConfig.Validate(); err != nil {
			return err
		}
	}
	if s.RedirectConfig != nil {
		if err := s.RedirectConfig.Validate(); err != nil {
			return err
		}
	}
	if s.RemoveHeaderConfig != nil {
		if err := s.RemoveHeaderConfig.Validate(); err != nil {
			return err
		}
	}
	if s.RewriteConfig != nil {
		if err := s.RewriteConfig.Validate(); err != nil {
			return err
		}
	}
	if s.TrafficLimitConfig != nil {
		if err := s.TrafficLimitConfig.Validate(); err != nil {
			return err
		}
	}
	if s.TrafficMirrorConfig != nil {
		if err := s.TrafficMirrorConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListRulesResponseBodyRulesRuleActionsCorsConfig struct {
	// Indicates whether credentials can be carried in CORS requests. Valid values:
	//
	// 	- **on**
	//
	// 	- **off**
	//
	// example:
	//
	// on
	AllowCredentials *string `json:"AllowCredentials,omitempty" xml:"AllowCredentials,omitempty"`
	// The allowed headers of CORS requests.
	AllowHeaders []*string `json:"AllowHeaders,omitempty" xml:"AllowHeaders,omitempty" type:"Repeated"`
	// The allowed HTTP methods of CORS requests.
	AllowMethods []*string `json:"AllowMethods,omitempty" xml:"AllowMethods,omitempty" type:"Repeated"`
	// The allowed origins of CORS requests.
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

func (s ListRulesResponseBodyRulesRuleActionsCorsConfig) String() string {
	return dara.Prettify(s)
}

func (s ListRulesResponseBodyRulesRuleActionsCorsConfig) GoString() string {
	return s.String()
}

func (s *ListRulesResponseBodyRulesRuleActionsCorsConfig) GetAllowCredentials() *string {
	return s.AllowCredentials
}

func (s *ListRulesResponseBodyRulesRuleActionsCorsConfig) GetAllowHeaders() []*string {
	return s.AllowHeaders
}

func (s *ListRulesResponseBodyRulesRuleActionsCorsConfig) GetAllowMethods() []*string {
	return s.AllowMethods
}

func (s *ListRulesResponseBodyRulesRuleActionsCorsConfig) GetAllowOrigin() []*string {
	return s.AllowOrigin
}

func (s *ListRulesResponseBodyRulesRuleActionsCorsConfig) GetExposeHeaders() []*string {
	return s.ExposeHeaders
}

func (s *ListRulesResponseBodyRulesRuleActionsCorsConfig) GetMaxAge() *int64 {
	return s.MaxAge
}

func (s *ListRulesResponseBodyRulesRuleActionsCorsConfig) SetAllowCredentials(v string) *ListRulesResponseBodyRulesRuleActionsCorsConfig {
	s.AllowCredentials = &v
	return s
}

func (s *ListRulesResponseBodyRulesRuleActionsCorsConfig) SetAllowHeaders(v []*string) *ListRulesResponseBodyRulesRuleActionsCorsConfig {
	s.AllowHeaders = v
	return s
}

func (s *ListRulesResponseBodyRulesRuleActionsCorsConfig) SetAllowMethods(v []*string) *ListRulesResponseBodyRulesRuleActionsCorsConfig {
	s.AllowMethods = v
	return s
}

func (s *ListRulesResponseBodyRulesRuleActionsCorsConfig) SetAllowOrigin(v []*string) *ListRulesResponseBodyRulesRuleActionsCorsConfig {
	s.AllowOrigin = v
	return s
}

func (s *ListRulesResponseBodyRulesRuleActionsCorsConfig) SetExposeHeaders(v []*string) *ListRulesResponseBodyRulesRuleActionsCorsConfig {
	s.ExposeHeaders = v
	return s
}

func (s *ListRulesResponseBodyRulesRuleActionsCorsConfig) SetMaxAge(v int64) *ListRulesResponseBodyRulesRuleActionsCorsConfig {
	s.MaxAge = &v
	return s
}

func (s *ListRulesResponseBodyRulesRuleActionsCorsConfig) Validate() error {
	return dara.Validate(s)
}

type ListRulesResponseBodyRulesRuleActionsFixedResponseConfig struct {
	// The content of the custom response. The content can be up to 1 KB in size, and can contain only ASCII characters.
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
	// The HTTP status code in responses. Valid values: **HTTP_2xx**, **HTTP_4xx**, and **HTTP_5xx**. **x*	- is a digit.
	//
	// example:
	//
	// HTTP_2xx
	HttpCode *string `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
}

func (s ListRulesResponseBodyRulesRuleActionsFixedResponseConfig) String() string {
	return dara.Prettify(s)
}

func (s ListRulesResponseBodyRulesRuleActionsFixedResponseConfig) GoString() string {
	return s.String()
}

func (s *ListRulesResponseBodyRulesRuleActionsFixedResponseConfig) GetContent() *string {
	return s.Content
}

func (s *ListRulesResponseBodyRulesRuleActionsFixedResponseConfig) GetContentType() *string {
	return s.ContentType
}

func (s *ListRulesResponseBodyRulesRuleActionsFixedResponseConfig) GetHttpCode() *string {
	return s.HttpCode
}

func (s *ListRulesResponseBodyRulesRuleActionsFixedResponseConfig) SetContent(v string) *ListRulesResponseBodyRulesRuleActionsFixedResponseConfig {
	s.Content = &v
	return s
}

func (s *ListRulesResponseBodyRulesRuleActionsFixedResponseConfig) SetContentType(v string) *ListRulesResponseBodyRulesRuleActionsFixedResponseConfig {
	s.ContentType = &v
	return s
}

func (s *ListRulesResponseBodyRulesRuleActionsFixedResponseConfig) SetHttpCode(v string) *ListRulesResponseBodyRulesRuleActionsFixedResponseConfig {
	s.HttpCode = &v
	return s
}

func (s *ListRulesResponseBodyRulesRuleActionsFixedResponseConfig) Validate() error {
	return dara.Validate(s)
}

type ListRulesResponseBodyRulesRuleActionsForwardGroupConfig struct {
	// The session persistence configurations of the server group.
	ServerGroupStickySession *ListRulesResponseBodyRulesRuleActionsForwardGroupConfigServerGroupStickySession `json:"ServerGroupStickySession,omitempty" xml:"ServerGroupStickySession,omitempty" type:"Struct"`
	// The server groups to which requests are forwarded.
	ServerGroupTuples []*ListRulesResponseBodyRulesRuleActionsForwardGroupConfigServerGroupTuples `json:"ServerGroupTuples,omitempty" xml:"ServerGroupTuples,omitempty" type:"Repeated"`
}

func (s ListRulesResponseBodyRulesRuleActionsForwardGroupConfig) String() string {
	return dara.Prettify(s)
}

func (s ListRulesResponseBodyRulesRuleActionsForwardGroupConfig) GoString() string {
	return s.String()
}

func (s *ListRulesResponseBodyRulesRuleActionsForwardGroupConfig) GetServerGroupStickySession() *ListRulesResponseBodyRulesRuleActionsForwardGroupConfigServerGroupStickySession {
	return s.ServerGroupStickySession
}

func (s *ListRulesResponseBodyRulesRuleActionsForwardGroupConfig) GetServerGroupTuples() []*ListRulesResponseBodyRulesRuleActionsForwardGroupConfigServerGroupTuples {
	return s.ServerGroupTuples
}

func (s *ListRulesResponseBodyRulesRuleActionsForwardGroupConfig) SetServerGroupStickySession(v *ListRulesResponseBodyRulesRuleActionsForwardGroupConfigServerGroupStickySession) *ListRulesResponseBodyRulesRuleActionsForwardGroupConfig {
	s.ServerGroupStickySession = v
	return s
}

func (s *ListRulesResponseBodyRulesRuleActionsForwardGroupConfig) SetServerGroupTuples(v []*ListRulesResponseBodyRulesRuleActionsForwardGroupConfigServerGroupTuples) *ListRulesResponseBodyRulesRuleActionsForwardGroupConfig {
	s.ServerGroupTuples = v
	return s
}

func (s *ListRulesResponseBodyRulesRuleActionsForwardGroupConfig) Validate() error {
	if s.ServerGroupStickySession != nil {
		if err := s.ServerGroupStickySession.Validate(); err != nil {
			return err
		}
	}
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

type ListRulesResponseBodyRulesRuleActionsForwardGroupConfigServerGroupStickySession struct {
	// If the value of N in ServerGroupTuple.N is larger than 1, you can enable or disable session persistence for server groups.
	//
	// example:
	//
	// true
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// If Enabled is set to True, you can specify a session persistence timeout period.
	//
	// example:
	//
	// 100
	Timeout *int32 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
}

func (s ListRulesResponseBodyRulesRuleActionsForwardGroupConfigServerGroupStickySession) String() string {
	return dara.Prettify(s)
}

func (s ListRulesResponseBodyRulesRuleActionsForwardGroupConfigServerGroupStickySession) GoString() string {
	return s.String()
}

func (s *ListRulesResponseBodyRulesRuleActionsForwardGroupConfigServerGroupStickySession) GetEnabled() *bool {
	return s.Enabled
}

func (s *ListRulesResponseBodyRulesRuleActionsForwardGroupConfigServerGroupStickySession) GetTimeout() *int32 {
	return s.Timeout
}

func (s *ListRulesResponseBodyRulesRuleActionsForwardGroupConfigServerGroupStickySession) SetEnabled(v bool) *ListRulesResponseBodyRulesRuleActionsForwardGroupConfigServerGroupStickySession {
	s.Enabled = &v
	return s
}

func (s *ListRulesResponseBodyRulesRuleActionsForwardGroupConfigServerGroupStickySession) SetTimeout(v int32) *ListRulesResponseBodyRulesRuleActionsForwardGroupConfigServerGroupStickySession {
	s.Timeout = &v
	return s
}

func (s *ListRulesResponseBodyRulesRuleActionsForwardGroupConfigServerGroupStickySession) Validate() error {
	return dara.Validate(s)
}

type ListRulesResponseBodyRulesRuleActionsForwardGroupConfigServerGroupTuples struct {
	// The server group to which requests are forwarded.
	//
	// example:
	//
	// sg-atstuj3rtoptyui****
	ServerGroupId *string `json:"ServerGroupId,omitempty" xml:"ServerGroupId,omitempty"`
	// The weight of the server group. Valid values: **0*	- to **100**.
	//
	// example:
	//
	// 2
	Weight *int32 `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s ListRulesResponseBodyRulesRuleActionsForwardGroupConfigServerGroupTuples) String() string {
	return dara.Prettify(s)
}

func (s ListRulesResponseBodyRulesRuleActionsForwardGroupConfigServerGroupTuples) GoString() string {
	return s.String()
}

func (s *ListRulesResponseBodyRulesRuleActionsForwardGroupConfigServerGroupTuples) GetServerGroupId() *string {
	return s.ServerGroupId
}

func (s *ListRulesResponseBodyRulesRuleActionsForwardGroupConfigServerGroupTuples) GetWeight() *int32 {
	return s.Weight
}

func (s *ListRulesResponseBodyRulesRuleActionsForwardGroupConfigServerGroupTuples) SetServerGroupId(v string) *ListRulesResponseBodyRulesRuleActionsForwardGroupConfigServerGroupTuples {
	s.ServerGroupId = &v
	return s
}

func (s *ListRulesResponseBodyRulesRuleActionsForwardGroupConfigServerGroupTuples) SetWeight(v int32) *ListRulesResponseBodyRulesRuleActionsForwardGroupConfigServerGroupTuples {
	s.Weight = &v
	return s
}

func (s *ListRulesResponseBodyRulesRuleActionsForwardGroupConfigServerGroupTuples) Validate() error {
	return dara.Validate(s)
}

type ListRulesResponseBodyRulesRuleActionsInsertHeaderConfig struct {
	// The key of the header. The header key must be 1 to 40 characters in length, and can contain letters, digits, underscores (_), and hyphens (-). The header key specified in `InsertHeader` must be unique.
	//
	// >  **Cookie*	- and **Host*	- are not supported.
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
	// 	- If **ValueType*	- is set to **UserDefined**, you can specify a custom header value. The header value must be 1 to 128 characters in length, and can contain wildcard characters, such as asterisks (\\*) and question marks (?), and printable characters whose ASCII values are `larger than or equal to 32 and smaller than 127`. The header value cannot start or end with a space character.
	//
	// 	- If **ValueType*	- is set to **ReferenceHeader**, you can reference a value from a request header. The header value must be 1 to 128 characters in length, and can contain lowercase letters, digits, hyphens (-), and underscores (_).
	//
	// example:
	//
	// ClientSrcPort
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
	// The type of the header value. Valid values:
	//
	// 	- **UserDefined**: a user-defined header value.
	//
	// 	- **ReferenceHeader**: a header value that is referenced from a request header.
	//
	// 	- **SystemDefined:*	- a system-defined header value.
	//
	// example:
	//
	// SystemDefined
	ValueType *string `json:"ValueType,omitempty" xml:"ValueType,omitempty"`
}

func (s ListRulesResponseBodyRulesRuleActionsInsertHeaderConfig) String() string {
	return dara.Prettify(s)
}

func (s ListRulesResponseBodyRulesRuleActionsInsertHeaderConfig) GoString() string {
	return s.String()
}

func (s *ListRulesResponseBodyRulesRuleActionsInsertHeaderConfig) GetKey() *string {
	return s.Key
}

func (s *ListRulesResponseBodyRulesRuleActionsInsertHeaderConfig) GetValue() *string {
	return s.Value
}

func (s *ListRulesResponseBodyRulesRuleActionsInsertHeaderConfig) GetValueType() *string {
	return s.ValueType
}

func (s *ListRulesResponseBodyRulesRuleActionsInsertHeaderConfig) SetKey(v string) *ListRulesResponseBodyRulesRuleActionsInsertHeaderConfig {
	s.Key = &v
	return s
}

func (s *ListRulesResponseBodyRulesRuleActionsInsertHeaderConfig) SetValue(v string) *ListRulesResponseBodyRulesRuleActionsInsertHeaderConfig {
	s.Value = &v
	return s
}

func (s *ListRulesResponseBodyRulesRuleActionsInsertHeaderConfig) SetValueType(v string) *ListRulesResponseBodyRulesRuleActionsInsertHeaderConfig {
	s.ValueType = &v
	return s
}

func (s *ListRulesResponseBodyRulesRuleActionsInsertHeaderConfig) Validate() error {
	return dara.Validate(s)
}

type ListRulesResponseBodyRulesRuleActionsRedirectConfig struct {
	// The hostname to which requests are redirected. Valid values:
	//
	// 	- **${host}*	- (default): If ${host} is returned, no other characters are appended.
	//
	// 	- A custom value. Make sure that the custom value meets the following requirements:
	//
	//     	- The hostname must be 3 to 128 characters in length, and can contain lowercase letters, digits, hyphens (-), periods (.), asterisks (\\*), and question marks (?).
	//
	//     	- The hostname must contain at least one period (.) but cannot start or end with a period (.).
	//
	//     	- The rightmost domain label can contain only letters and wildcard characters. It cannot contain digits or hyphens (-).
	//
	//     	- The domain labels cannot start or end with a hyphen (-).
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
	// 	- **${path}*	- (default): You can reference \\*\\*${host}**, **${protocol}**, and**${port}**. The URL can consist of **${host}**,**${protocol}**, and **${port}\\*\\*. Each variable can be used only once. The preceding variables can be used at the same time or combined with a custom value.
	//
	// 	- A custom value. Make sure that the custom value meets the following requirements:
	//
	//     	- The URL must be 1 to 128 characters in length.
	//
	//     	- It must start with a forward slash (/) and can contain letters, digits, and the following special characters: `$ - _ .+ / & ~ @ :`. It cannot contain the following special characters: `" % # ; ! ( ) [ ] ^ , "`. You can use asterisks (\\*) and question marks (?) as wildcard characters.
	//
	// example:
	//
	// /test
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// The port to which requests are redirected. Valid values:
	//
	// 	- **${port}*	- (default): If ${port} is returned, no other characters are appended.
	//
	// 	- Other valid values: **1 to 63335**.
	//
	// example:
	//
	// 10
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// The redirect protocol. Valid values:
	//
	// 	- **${protocol}*	- (default): If ${protocol} is returned, no other characters are appended.
	//
	// 	- **HTTP*	- or **HTTPS**
	//
	// >  HTTPS listeners supports only HTTPS redirects.
	//
	// example:
	//
	// HTTP
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// The query string of the URL to which requests are redirected. The query string must be 1 to 128 characters in length, and can contain printable characters, excluding uppercase letters and the following special characters: `# [ ] { } \\ | < > &`.
	//
	// example:
	//
	// quert
	Query *string `json:"Query,omitempty" xml:"Query,omitempty"`
}

func (s ListRulesResponseBodyRulesRuleActionsRedirectConfig) String() string {
	return dara.Prettify(s)
}

func (s ListRulesResponseBodyRulesRuleActionsRedirectConfig) GoString() string {
	return s.String()
}

func (s *ListRulesResponseBodyRulesRuleActionsRedirectConfig) GetHost() *string {
	return s.Host
}

func (s *ListRulesResponseBodyRulesRuleActionsRedirectConfig) GetHttpCode() *string {
	return s.HttpCode
}

func (s *ListRulesResponseBodyRulesRuleActionsRedirectConfig) GetPath() *string {
	return s.Path
}

func (s *ListRulesResponseBodyRulesRuleActionsRedirectConfig) GetPort() *string {
	return s.Port
}

func (s *ListRulesResponseBodyRulesRuleActionsRedirectConfig) GetProtocol() *string {
	return s.Protocol
}

func (s *ListRulesResponseBodyRulesRuleActionsRedirectConfig) GetQuery() *string {
	return s.Query
}

func (s *ListRulesResponseBodyRulesRuleActionsRedirectConfig) SetHost(v string) *ListRulesResponseBodyRulesRuleActionsRedirectConfig {
	s.Host = &v
	return s
}

func (s *ListRulesResponseBodyRulesRuleActionsRedirectConfig) SetHttpCode(v string) *ListRulesResponseBodyRulesRuleActionsRedirectConfig {
	s.HttpCode = &v
	return s
}

func (s *ListRulesResponseBodyRulesRuleActionsRedirectConfig) SetPath(v string) *ListRulesResponseBodyRulesRuleActionsRedirectConfig {
	s.Path = &v
	return s
}

func (s *ListRulesResponseBodyRulesRuleActionsRedirectConfig) SetPort(v string) *ListRulesResponseBodyRulesRuleActionsRedirectConfig {
	s.Port = &v
	return s
}

func (s *ListRulesResponseBodyRulesRuleActionsRedirectConfig) SetProtocol(v string) *ListRulesResponseBodyRulesRuleActionsRedirectConfig {
	s.Protocol = &v
	return s
}

func (s *ListRulesResponseBodyRulesRuleActionsRedirectConfig) SetQuery(v string) *ListRulesResponseBodyRulesRuleActionsRedirectConfig {
	s.Query = &v
	return s
}

func (s *ListRulesResponseBodyRulesRuleActionsRedirectConfig) Validate() error {
	return dara.Validate(s)
}

type ListRulesResponseBodyRulesRuleActionsRemoveHeaderConfig struct {
	// The key of the header to be removed. The header key must be 1 to 40 characters in length, and can contain letters, digits, underscores (_), and hyphens (-). The header keys specified in RemoveHeader must be unique.
	//
	// 	- If Direction is set to Request, the specified headers are removed from requests. The following header keys are not supported (not case-sensitive): `slb-id`, `slb-ip`, `x-forwarded-for`, `x-forwarded-proto`, `x-forwarded-eip`, `x-forwarded-port`, `x-forwarded-client-srcport`, `connection`, `upgrade`, `content-length`, `transfer-encoding`, `keep-alive`, `te`, `host`, `cookie`, `remoteip`, and `authority`.
	//
	// 	- If Direction is set to Response, the specified headers are removed from responses. The following header keys are not supported (not case-sensitive): `connection`, `upgrade`, `content-length`, and `transfer-encoding`.
	//
	// example:
	//
	// key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
}

func (s ListRulesResponseBodyRulesRuleActionsRemoveHeaderConfig) String() string {
	return dara.Prettify(s)
}

func (s ListRulesResponseBodyRulesRuleActionsRemoveHeaderConfig) GoString() string {
	return s.String()
}

func (s *ListRulesResponseBodyRulesRuleActionsRemoveHeaderConfig) GetKey() *string {
	return s.Key
}

func (s *ListRulesResponseBodyRulesRuleActionsRemoveHeaderConfig) SetKey(v string) *ListRulesResponseBodyRulesRuleActionsRemoveHeaderConfig {
	s.Key = &v
	return s
}

func (s *ListRulesResponseBodyRulesRuleActionsRemoveHeaderConfig) Validate() error {
	return dara.Validate(s)
}

type ListRulesResponseBodyRulesRuleActionsRewriteConfig struct {
	// The hostname to which requests are redirected. Valid values:
	//
	// 	- **${host}*	- (default): If ${host} is returned, no other characters are appended.
	//
	// 	- A custom value. Make sure that the custom value meets the following requirements:
	//
	//     	- The hostname must be 3 to 128 characters in length, and can contain lowercase letters, digits, hyphens (-), periods (.), asterisks (\\*), and question marks (?).
	//
	//     	- The hostname must contain at least one period (.) but cannot start or end with a period (.).
	//
	//     	- The rightmost domain label can contain only letters and wildcard characters. It cannot contain digits or hyphens (-).
	//
	//     	- The domain labels cannot start or end with a hyphen (-).
	//
	//     	- You can use asterisks (\\*) and question marks (?) anywhere in a domain label as wildcard characters.
	//
	// example:
	//
	// www.example.com
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// The URL to which requests are redirected. The URL must be 1 to 128 characters in length, and can contain letters, digits, asterisks (\\*), question marks (?), and the following special characters: `$ - _ . + / & ~ @ :`. It must start with a forward slash (/) and does not contain the following special characters: `" % # ; ! ( ) [ ] ^ , "`.
	//
	// example:
	//
	// /tsdf
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// The query string of the URL to which requests are redirected. The query string must be 1 to 128 characters in length, and can contain printable characters, excluding uppercase letters and the following special characters: `# [ ] { } \\ | < > &`.
	//
	// example:
	//
	// quedsa
	Query *string `json:"Query,omitempty" xml:"Query,omitempty"`
}

func (s ListRulesResponseBodyRulesRuleActionsRewriteConfig) String() string {
	return dara.Prettify(s)
}

func (s ListRulesResponseBodyRulesRuleActionsRewriteConfig) GoString() string {
	return s.String()
}

func (s *ListRulesResponseBodyRulesRuleActionsRewriteConfig) GetHost() *string {
	return s.Host
}

func (s *ListRulesResponseBodyRulesRuleActionsRewriteConfig) GetPath() *string {
	return s.Path
}

func (s *ListRulesResponseBodyRulesRuleActionsRewriteConfig) GetQuery() *string {
	return s.Query
}

func (s *ListRulesResponseBodyRulesRuleActionsRewriteConfig) SetHost(v string) *ListRulesResponseBodyRulesRuleActionsRewriteConfig {
	s.Host = &v
	return s
}

func (s *ListRulesResponseBodyRulesRuleActionsRewriteConfig) SetPath(v string) *ListRulesResponseBodyRulesRuleActionsRewriteConfig {
	s.Path = &v
	return s
}

func (s *ListRulesResponseBodyRulesRuleActionsRewriteConfig) SetQuery(v string) *ListRulesResponseBodyRulesRuleActionsRewriteConfig {
	s.Query = &v
	return s
}

func (s *ListRulesResponseBodyRulesRuleActionsRewriteConfig) Validate() error {
	return dara.Validate(s)
}

type ListRulesResponseBodyRulesRuleActionsTrafficLimitConfig struct {
	// The number of requests per IP address. Valid values: **1 to 100000**.
	//
	// >  If both the **QPS*	- and **PerIpQps*	- parameters are specified, the value of the **QPS*	- parameter is smaller than the value of the PerIpQps parameter.
	//
	// example:
	//
	// 80
	PerIpQps *int32 `json:"PerIpQps,omitempty" xml:"PerIpQps,omitempty"`
	// The number of queries per second (QPS). Valid values: **1*	- to **100000**.
	//
	// example:
	//
	// 4
	QPS *int32 `json:"QPS,omitempty" xml:"QPS,omitempty"`
}

func (s ListRulesResponseBodyRulesRuleActionsTrafficLimitConfig) String() string {
	return dara.Prettify(s)
}

func (s ListRulesResponseBodyRulesRuleActionsTrafficLimitConfig) GoString() string {
	return s.String()
}

func (s *ListRulesResponseBodyRulesRuleActionsTrafficLimitConfig) GetPerIpQps() *int32 {
	return s.PerIpQps
}

func (s *ListRulesResponseBodyRulesRuleActionsTrafficLimitConfig) GetQPS() *int32 {
	return s.QPS
}

func (s *ListRulesResponseBodyRulesRuleActionsTrafficLimitConfig) SetPerIpQps(v int32) *ListRulesResponseBodyRulesRuleActionsTrafficLimitConfig {
	s.PerIpQps = &v
	return s
}

func (s *ListRulesResponseBodyRulesRuleActionsTrafficLimitConfig) SetQPS(v int32) *ListRulesResponseBodyRulesRuleActionsTrafficLimitConfig {
	s.QPS = &v
	return s
}

func (s *ListRulesResponseBodyRulesRuleActionsTrafficLimitConfig) Validate() error {
	return dara.Validate(s)
}

type ListRulesResponseBodyRulesRuleActionsTrafficMirrorConfig struct {
	// The configuration of the server group to which traffic is mirrored.
	MirrorGroupConfig *ListRulesResponseBodyRulesRuleActionsTrafficMirrorConfigMirrorGroupConfig `json:"MirrorGroupConfig,omitempty" xml:"MirrorGroupConfig,omitempty" type:"Struct"`
	// The destination to which traffic is mirrored. The destination can be a server group.
	//
	// example:
	//
	// ForwardGroupMirror
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
}

func (s ListRulesResponseBodyRulesRuleActionsTrafficMirrorConfig) String() string {
	return dara.Prettify(s)
}

func (s ListRulesResponseBodyRulesRuleActionsTrafficMirrorConfig) GoString() string {
	return s.String()
}

func (s *ListRulesResponseBodyRulesRuleActionsTrafficMirrorConfig) GetMirrorGroupConfig() *ListRulesResponseBodyRulesRuleActionsTrafficMirrorConfigMirrorGroupConfig {
	return s.MirrorGroupConfig
}

func (s *ListRulesResponseBodyRulesRuleActionsTrafficMirrorConfig) GetTargetType() *string {
	return s.TargetType
}

func (s *ListRulesResponseBodyRulesRuleActionsTrafficMirrorConfig) SetMirrorGroupConfig(v *ListRulesResponseBodyRulesRuleActionsTrafficMirrorConfigMirrorGroupConfig) *ListRulesResponseBodyRulesRuleActionsTrafficMirrorConfig {
	s.MirrorGroupConfig = v
	return s
}

func (s *ListRulesResponseBodyRulesRuleActionsTrafficMirrorConfig) SetTargetType(v string) *ListRulesResponseBodyRulesRuleActionsTrafficMirrorConfig {
	s.TargetType = &v
	return s
}

func (s *ListRulesResponseBodyRulesRuleActionsTrafficMirrorConfig) Validate() error {
	if s.MirrorGroupConfig != nil {
		if err := s.MirrorGroupConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListRulesResponseBodyRulesRuleActionsTrafficMirrorConfigMirrorGroupConfig struct {
	// The server group to which traffic is mirrored.
	ServerGroupTuples []*ListRulesResponseBodyRulesRuleActionsTrafficMirrorConfigMirrorGroupConfigServerGroupTuples `json:"ServerGroupTuples,omitempty" xml:"ServerGroupTuples,omitempty" type:"Repeated"`
}

func (s ListRulesResponseBodyRulesRuleActionsTrafficMirrorConfigMirrorGroupConfig) String() string {
	return dara.Prettify(s)
}

func (s ListRulesResponseBodyRulesRuleActionsTrafficMirrorConfigMirrorGroupConfig) GoString() string {
	return s.String()
}

func (s *ListRulesResponseBodyRulesRuleActionsTrafficMirrorConfigMirrorGroupConfig) GetServerGroupTuples() []*ListRulesResponseBodyRulesRuleActionsTrafficMirrorConfigMirrorGroupConfigServerGroupTuples {
	return s.ServerGroupTuples
}

func (s *ListRulesResponseBodyRulesRuleActionsTrafficMirrorConfigMirrorGroupConfig) SetServerGroupTuples(v []*ListRulesResponseBodyRulesRuleActionsTrafficMirrorConfigMirrorGroupConfigServerGroupTuples) *ListRulesResponseBodyRulesRuleActionsTrafficMirrorConfigMirrorGroupConfig {
	s.ServerGroupTuples = v
	return s
}

func (s *ListRulesResponseBodyRulesRuleActionsTrafficMirrorConfigMirrorGroupConfig) Validate() error {
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

type ListRulesResponseBodyRulesRuleActionsTrafficMirrorConfigMirrorGroupConfigServerGroupTuples struct {
	// The ID of the server group.
	//
	// example:
	//
	// srg-00mkgijak0w4qgz9****
	ServerGroupId *string `json:"ServerGroupId,omitempty" xml:"ServerGroupId,omitempty"`
	// The weight of the server group. Valid values: **0*	- to **100**.
	//
	// example:
	//
	// 2
	Weight *int32 `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s ListRulesResponseBodyRulesRuleActionsTrafficMirrorConfigMirrorGroupConfigServerGroupTuples) String() string {
	return dara.Prettify(s)
}

func (s ListRulesResponseBodyRulesRuleActionsTrafficMirrorConfigMirrorGroupConfigServerGroupTuples) GoString() string {
	return s.String()
}

func (s *ListRulesResponseBodyRulesRuleActionsTrafficMirrorConfigMirrorGroupConfigServerGroupTuples) GetServerGroupId() *string {
	return s.ServerGroupId
}

func (s *ListRulesResponseBodyRulesRuleActionsTrafficMirrorConfigMirrorGroupConfigServerGroupTuples) GetWeight() *int32 {
	return s.Weight
}

func (s *ListRulesResponseBodyRulesRuleActionsTrafficMirrorConfigMirrorGroupConfigServerGroupTuples) SetServerGroupId(v string) *ListRulesResponseBodyRulesRuleActionsTrafficMirrorConfigMirrorGroupConfigServerGroupTuples {
	s.ServerGroupId = &v
	return s
}

func (s *ListRulesResponseBodyRulesRuleActionsTrafficMirrorConfigMirrorGroupConfigServerGroupTuples) SetWeight(v int32) *ListRulesResponseBodyRulesRuleActionsTrafficMirrorConfigMirrorGroupConfigServerGroupTuples {
	s.Weight = &v
	return s
}

func (s *ListRulesResponseBodyRulesRuleActionsTrafficMirrorConfigMirrorGroupConfigServerGroupTuples) Validate() error {
	return dara.Validate(s)
}

type ListRulesResponseBodyRulesRuleConditions struct {
	// The key-value pairs of the cookie.
	CookieConfig *ListRulesResponseBodyRulesRuleConditionsCookieConfig `json:"CookieConfig,omitempty" xml:"CookieConfig,omitempty" type:"Struct"`
	// The configuration of the header.
	HeaderConfig *ListRulesResponseBodyRulesRuleConditionsHeaderConfig `json:"HeaderConfig,omitempty" xml:"HeaderConfig,omitempty" type:"Struct"`
	// The configuration of the hosts.
	HostConfig *ListRulesResponseBodyRulesRuleConditionsHostConfig `json:"HostConfig,omitempty" xml:"HostConfig,omitempty" type:"Struct"`
	// The configurations of the request methods.
	MethodConfig *ListRulesResponseBodyRulesRuleConditionsMethodConfig `json:"MethodConfig,omitempty" xml:"MethodConfig,omitempty" type:"Struct"`
	// The configurations of the forwarding URLs.
	PathConfig *ListRulesResponseBodyRulesRuleConditionsPathConfig `json:"PathConfig,omitempty" xml:"PathConfig,omitempty" type:"Struct"`
	// The configurations of the query strings.
	QueryStringConfig *ListRulesResponseBodyRulesRuleConditionsQueryStringConfig `json:"QueryStringConfig,omitempty" xml:"QueryStringConfig,omitempty" type:"Struct"`
	// The HTTP header in responses.
	ResponseHeaderConfig *ListRulesResponseBodyRulesRuleConditionsResponseHeaderConfig `json:"ResponseHeaderConfig,omitempty" xml:"ResponseHeaderConfig,omitempty" type:"Struct"`
	// The configurations of the response status codes.
	ResponseStatusCodeConfig *ListRulesResponseBodyRulesRuleConditionsResponseStatusCodeConfig `json:"ResponseStatusCodeConfig,omitempty" xml:"ResponseStatusCodeConfig,omitempty" type:"Struct"`
	// Traffic matching based on source IP addresses.
	SourceIpConfig *ListRulesResponseBodyRulesRuleConditionsSourceIpConfig `json:"SourceIpConfig,omitempty" xml:"SourceIpConfig,omitempty" type:"Struct"`
	// The type of forwarding rule. Valid values:
	//
	// 	- **Host**: Responses are forwarded based on hosts.
	//
	// 	- **Path**: Responses are forwarded based on URLs.
	//
	// 	- **Header**: Responses are forwarded based on HTTP headers.
	//
	// 	- **QueryString**: Responses are forwarded based on query strings.
	//
	// 	- **Method**: Responses are forwarded based on request methods.
	//
	// 	- **Cookie**: Responses are forwarded based on cookies.
	//
	// 	- **SourceIp**: Responses are forwarded based on source IP addresses.
	//
	// example:
	//
	// Host
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListRulesResponseBodyRulesRuleConditions) String() string {
	return dara.Prettify(s)
}

func (s ListRulesResponseBodyRulesRuleConditions) GoString() string {
	return s.String()
}

func (s *ListRulesResponseBodyRulesRuleConditions) GetCookieConfig() *ListRulesResponseBodyRulesRuleConditionsCookieConfig {
	return s.CookieConfig
}

func (s *ListRulesResponseBodyRulesRuleConditions) GetHeaderConfig() *ListRulesResponseBodyRulesRuleConditionsHeaderConfig {
	return s.HeaderConfig
}

func (s *ListRulesResponseBodyRulesRuleConditions) GetHostConfig() *ListRulesResponseBodyRulesRuleConditionsHostConfig {
	return s.HostConfig
}

func (s *ListRulesResponseBodyRulesRuleConditions) GetMethodConfig() *ListRulesResponseBodyRulesRuleConditionsMethodConfig {
	return s.MethodConfig
}

func (s *ListRulesResponseBodyRulesRuleConditions) GetPathConfig() *ListRulesResponseBodyRulesRuleConditionsPathConfig {
	return s.PathConfig
}

func (s *ListRulesResponseBodyRulesRuleConditions) GetQueryStringConfig() *ListRulesResponseBodyRulesRuleConditionsQueryStringConfig {
	return s.QueryStringConfig
}

func (s *ListRulesResponseBodyRulesRuleConditions) GetResponseHeaderConfig() *ListRulesResponseBodyRulesRuleConditionsResponseHeaderConfig {
	return s.ResponseHeaderConfig
}

func (s *ListRulesResponseBodyRulesRuleConditions) GetResponseStatusCodeConfig() *ListRulesResponseBodyRulesRuleConditionsResponseStatusCodeConfig {
	return s.ResponseStatusCodeConfig
}

func (s *ListRulesResponseBodyRulesRuleConditions) GetSourceIpConfig() *ListRulesResponseBodyRulesRuleConditionsSourceIpConfig {
	return s.SourceIpConfig
}

func (s *ListRulesResponseBodyRulesRuleConditions) GetType() *string {
	return s.Type
}

func (s *ListRulesResponseBodyRulesRuleConditions) SetCookieConfig(v *ListRulesResponseBodyRulesRuleConditionsCookieConfig) *ListRulesResponseBodyRulesRuleConditions {
	s.CookieConfig = v
	return s
}

func (s *ListRulesResponseBodyRulesRuleConditions) SetHeaderConfig(v *ListRulesResponseBodyRulesRuleConditionsHeaderConfig) *ListRulesResponseBodyRulesRuleConditions {
	s.HeaderConfig = v
	return s
}

func (s *ListRulesResponseBodyRulesRuleConditions) SetHostConfig(v *ListRulesResponseBodyRulesRuleConditionsHostConfig) *ListRulesResponseBodyRulesRuleConditions {
	s.HostConfig = v
	return s
}

func (s *ListRulesResponseBodyRulesRuleConditions) SetMethodConfig(v *ListRulesResponseBodyRulesRuleConditionsMethodConfig) *ListRulesResponseBodyRulesRuleConditions {
	s.MethodConfig = v
	return s
}

func (s *ListRulesResponseBodyRulesRuleConditions) SetPathConfig(v *ListRulesResponseBodyRulesRuleConditionsPathConfig) *ListRulesResponseBodyRulesRuleConditions {
	s.PathConfig = v
	return s
}

func (s *ListRulesResponseBodyRulesRuleConditions) SetQueryStringConfig(v *ListRulesResponseBodyRulesRuleConditionsQueryStringConfig) *ListRulesResponseBodyRulesRuleConditions {
	s.QueryStringConfig = v
	return s
}

func (s *ListRulesResponseBodyRulesRuleConditions) SetResponseHeaderConfig(v *ListRulesResponseBodyRulesRuleConditionsResponseHeaderConfig) *ListRulesResponseBodyRulesRuleConditions {
	s.ResponseHeaderConfig = v
	return s
}

func (s *ListRulesResponseBodyRulesRuleConditions) SetResponseStatusCodeConfig(v *ListRulesResponseBodyRulesRuleConditionsResponseStatusCodeConfig) *ListRulesResponseBodyRulesRuleConditions {
	s.ResponseStatusCodeConfig = v
	return s
}

func (s *ListRulesResponseBodyRulesRuleConditions) SetSourceIpConfig(v *ListRulesResponseBodyRulesRuleConditionsSourceIpConfig) *ListRulesResponseBodyRulesRuleConditions {
	s.SourceIpConfig = v
	return s
}

func (s *ListRulesResponseBodyRulesRuleConditions) SetType(v string) *ListRulesResponseBodyRulesRuleConditions {
	s.Type = &v
	return s
}

func (s *ListRulesResponseBodyRulesRuleConditions) Validate() error {
	if s.CookieConfig != nil {
		if err := s.CookieConfig.Validate(); err != nil {
			return err
		}
	}
	if s.HeaderConfig != nil {
		if err := s.HeaderConfig.Validate(); err != nil {
			return err
		}
	}
	if s.HostConfig != nil {
		if err := s.HostConfig.Validate(); err != nil {
			return err
		}
	}
	if s.MethodConfig != nil {
		if err := s.MethodConfig.Validate(); err != nil {
			return err
		}
	}
	if s.PathConfig != nil {
		if err := s.PathConfig.Validate(); err != nil {
			return err
		}
	}
	if s.QueryStringConfig != nil {
		if err := s.QueryStringConfig.Validate(); err != nil {
			return err
		}
	}
	if s.ResponseHeaderConfig != nil {
		if err := s.ResponseHeaderConfig.Validate(); err != nil {
			return err
		}
	}
	if s.ResponseStatusCodeConfig != nil {
		if err := s.ResponseStatusCodeConfig.Validate(); err != nil {
			return err
		}
	}
	if s.SourceIpConfig != nil {
		if err := s.SourceIpConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListRulesResponseBodyRulesRuleConditionsCookieConfig struct {
	// The cookie value.
	Values []*ListRulesResponseBodyRulesRuleConditionsCookieConfigValues `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s ListRulesResponseBodyRulesRuleConditionsCookieConfig) String() string {
	return dara.Prettify(s)
}

func (s ListRulesResponseBodyRulesRuleConditionsCookieConfig) GoString() string {
	return s.String()
}

func (s *ListRulesResponseBodyRulesRuleConditionsCookieConfig) GetValues() []*ListRulesResponseBodyRulesRuleConditionsCookieConfigValues {
	return s.Values
}

func (s *ListRulesResponseBodyRulesRuleConditionsCookieConfig) SetValues(v []*ListRulesResponseBodyRulesRuleConditionsCookieConfigValues) *ListRulesResponseBodyRulesRuleConditionsCookieConfig {
	s.Values = v
	return s
}

func (s *ListRulesResponseBodyRulesRuleConditionsCookieConfig) Validate() error {
	if s.Values != nil {
		for _, item := range s.Values {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListRulesResponseBodyRulesRuleConditionsCookieConfigValues struct {
	// The cookie key. The cookie key must be 1 to 100 characters in length, and can contain lowercase letters, printable ASCII characters, asterisks (\\*), and question marks (?). It cannot contain space characters or the following special characters: `# [ ] { } \\ | < > &`.
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

func (s ListRulesResponseBodyRulesRuleConditionsCookieConfigValues) String() string {
	return dara.Prettify(s)
}

func (s ListRulesResponseBodyRulesRuleConditionsCookieConfigValues) GoString() string {
	return s.String()
}

func (s *ListRulesResponseBodyRulesRuleConditionsCookieConfigValues) GetKey() *string {
	return s.Key
}

func (s *ListRulesResponseBodyRulesRuleConditionsCookieConfigValues) GetValue() *string {
	return s.Value
}

func (s *ListRulesResponseBodyRulesRuleConditionsCookieConfigValues) SetKey(v string) *ListRulesResponseBodyRulesRuleConditionsCookieConfigValues {
	s.Key = &v
	return s
}

func (s *ListRulesResponseBodyRulesRuleConditionsCookieConfigValues) SetValue(v string) *ListRulesResponseBodyRulesRuleConditionsCookieConfigValues {
	s.Value = &v
	return s
}

func (s *ListRulesResponseBodyRulesRuleConditionsCookieConfigValues) Validate() error {
	return dara.Validate(s)
}

type ListRulesResponseBodyRulesRuleConditionsHeaderConfig struct {
	// The key of the header. The header key must be 1 to 40 characters in length. It can contain letters, digits, hyphens (-), and underscores (_). Cookie and Host are not supported.
	//
	// example:
	//
	// Port
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the header.
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s ListRulesResponseBodyRulesRuleConditionsHeaderConfig) String() string {
	return dara.Prettify(s)
}

func (s ListRulesResponseBodyRulesRuleConditionsHeaderConfig) GoString() string {
	return s.String()
}

func (s *ListRulesResponseBodyRulesRuleConditionsHeaderConfig) GetKey() *string {
	return s.Key
}

func (s *ListRulesResponseBodyRulesRuleConditionsHeaderConfig) GetValues() []*string {
	return s.Values
}

func (s *ListRulesResponseBodyRulesRuleConditionsHeaderConfig) SetKey(v string) *ListRulesResponseBodyRulesRuleConditionsHeaderConfig {
	s.Key = &v
	return s
}

func (s *ListRulesResponseBodyRulesRuleConditionsHeaderConfig) SetValues(v []*string) *ListRulesResponseBodyRulesRuleConditionsHeaderConfig {
	s.Values = v
	return s
}

func (s *ListRulesResponseBodyRulesRuleConditionsHeaderConfig) Validate() error {
	return dara.Validate(s)
}

type ListRulesResponseBodyRulesRuleConditionsHostConfig struct {
	// The hostnames.
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s ListRulesResponseBodyRulesRuleConditionsHostConfig) String() string {
	return dara.Prettify(s)
}

func (s ListRulesResponseBodyRulesRuleConditionsHostConfig) GoString() string {
	return s.String()
}

func (s *ListRulesResponseBodyRulesRuleConditionsHostConfig) GetValues() []*string {
	return s.Values
}

func (s *ListRulesResponseBodyRulesRuleConditionsHostConfig) SetValues(v []*string) *ListRulesResponseBodyRulesRuleConditionsHostConfig {
	s.Values = v
	return s
}

func (s *ListRulesResponseBodyRulesRuleConditionsHostConfig) Validate() error {
	return dara.Validate(s)
}

type ListRulesResponseBodyRulesRuleConditionsMethodConfig struct {
	// The request methods.
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s ListRulesResponseBodyRulesRuleConditionsMethodConfig) String() string {
	return dara.Prettify(s)
}

func (s ListRulesResponseBodyRulesRuleConditionsMethodConfig) GoString() string {
	return s.String()
}

func (s *ListRulesResponseBodyRulesRuleConditionsMethodConfig) GetValues() []*string {
	return s.Values
}

func (s *ListRulesResponseBodyRulesRuleConditionsMethodConfig) SetValues(v []*string) *ListRulesResponseBodyRulesRuleConditionsMethodConfig {
	s.Values = v
	return s
}

func (s *ListRulesResponseBodyRulesRuleConditionsMethodConfig) Validate() error {
	return dara.Validate(s)
}

type ListRulesResponseBodyRulesRuleConditionsPathConfig struct {
	// The URLs to which requests are forwarded.
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s ListRulesResponseBodyRulesRuleConditionsPathConfig) String() string {
	return dara.Prettify(s)
}

func (s ListRulesResponseBodyRulesRuleConditionsPathConfig) GoString() string {
	return s.String()
}

func (s *ListRulesResponseBodyRulesRuleConditionsPathConfig) GetValues() []*string {
	return s.Values
}

func (s *ListRulesResponseBodyRulesRuleConditionsPathConfig) SetValues(v []*string) *ListRulesResponseBodyRulesRuleConditionsPathConfig {
	s.Values = v
	return s
}

func (s *ListRulesResponseBodyRulesRuleConditionsPathConfig) Validate() error {
	return dara.Validate(s)
}

type ListRulesResponseBodyRulesRuleConditionsQueryStringConfig struct {
	// The query string.
	Values []*ListRulesResponseBodyRulesRuleConditionsQueryStringConfigValues `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s ListRulesResponseBodyRulesRuleConditionsQueryStringConfig) String() string {
	return dara.Prettify(s)
}

func (s ListRulesResponseBodyRulesRuleConditionsQueryStringConfig) GoString() string {
	return s.String()
}

func (s *ListRulesResponseBodyRulesRuleConditionsQueryStringConfig) GetValues() []*ListRulesResponseBodyRulesRuleConditionsQueryStringConfigValues {
	return s.Values
}

func (s *ListRulesResponseBodyRulesRuleConditionsQueryStringConfig) SetValues(v []*ListRulesResponseBodyRulesRuleConditionsQueryStringConfigValues) *ListRulesResponseBodyRulesRuleConditionsQueryStringConfig {
	s.Values = v
	return s
}

func (s *ListRulesResponseBodyRulesRuleConditionsQueryStringConfig) Validate() error {
	if s.Values != nil {
		for _, item := range s.Values {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListRulesResponseBodyRulesRuleConditionsQueryStringConfigValues struct {
	// They key of the query string. The key must be 1 to 100 characters in length, and can contain lowercase letters, printable ASCII characters, asterisks (\\*), and question marks (?). It cannot contain space characters or the following special characters: `# [ ] { } \\ | < > &`.
	//
	// example:
	//
	// test
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the query string. The value must be 1 to 128 characters in length, and can contain lowercase letters, printable ASCII characters, asterisks (\\*), and question marks (?). It cannot contain space characters or the following special characters: `# [ ] { } \\ | < > &`.
	//
	// example:
	//
	// test
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListRulesResponseBodyRulesRuleConditionsQueryStringConfigValues) String() string {
	return dara.Prettify(s)
}

func (s ListRulesResponseBodyRulesRuleConditionsQueryStringConfigValues) GoString() string {
	return s.String()
}

func (s *ListRulesResponseBodyRulesRuleConditionsQueryStringConfigValues) GetKey() *string {
	return s.Key
}

func (s *ListRulesResponseBodyRulesRuleConditionsQueryStringConfigValues) GetValue() *string {
	return s.Value
}

func (s *ListRulesResponseBodyRulesRuleConditionsQueryStringConfigValues) SetKey(v string) *ListRulesResponseBodyRulesRuleConditionsQueryStringConfigValues {
	s.Key = &v
	return s
}

func (s *ListRulesResponseBodyRulesRuleConditionsQueryStringConfigValues) SetValue(v string) *ListRulesResponseBodyRulesRuleConditionsQueryStringConfigValues {
	s.Value = &v
	return s
}

func (s *ListRulesResponseBodyRulesRuleConditionsQueryStringConfigValues) Validate() error {
	return dara.Validate(s)
}

type ListRulesResponseBodyRulesRuleConditionsResponseHeaderConfig struct {
	// The key of the HTTP header. The header key must be 1 to 40 characters in length, It can contain letters, digits, hyphens (-), and underscores (_). Cookie and Host are not supported.
	//
	// example:
	//
	// key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The values of the HTTP header.
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s ListRulesResponseBodyRulesRuleConditionsResponseHeaderConfig) String() string {
	return dara.Prettify(s)
}

func (s ListRulesResponseBodyRulesRuleConditionsResponseHeaderConfig) GoString() string {
	return s.String()
}

func (s *ListRulesResponseBodyRulesRuleConditionsResponseHeaderConfig) GetKey() *string {
	return s.Key
}

func (s *ListRulesResponseBodyRulesRuleConditionsResponseHeaderConfig) GetValues() []*string {
	return s.Values
}

func (s *ListRulesResponseBodyRulesRuleConditionsResponseHeaderConfig) SetKey(v string) *ListRulesResponseBodyRulesRuleConditionsResponseHeaderConfig {
	s.Key = &v
	return s
}

func (s *ListRulesResponseBodyRulesRuleConditionsResponseHeaderConfig) SetValues(v []*string) *ListRulesResponseBodyRulesRuleConditionsResponseHeaderConfig {
	s.Values = v
	return s
}

func (s *ListRulesResponseBodyRulesRuleConditionsResponseHeaderConfig) Validate() error {
	return dara.Validate(s)
}

type ListRulesResponseBodyRulesRuleConditionsResponseStatusCodeConfig struct {
	// The response status codes.
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s ListRulesResponseBodyRulesRuleConditionsResponseStatusCodeConfig) String() string {
	return dara.Prettify(s)
}

func (s ListRulesResponseBodyRulesRuleConditionsResponseStatusCodeConfig) GoString() string {
	return s.String()
}

func (s *ListRulesResponseBodyRulesRuleConditionsResponseStatusCodeConfig) GetValues() []*string {
	return s.Values
}

func (s *ListRulesResponseBodyRulesRuleConditionsResponseStatusCodeConfig) SetValues(v []*string) *ListRulesResponseBodyRulesRuleConditionsResponseStatusCodeConfig {
	s.Values = v
	return s
}

func (s *ListRulesResponseBodyRulesRuleConditionsResponseStatusCodeConfig) Validate() error {
	return dara.Validate(s)
}

type ListRulesResponseBodyRulesRuleConditionsSourceIpConfig struct {
	// The source IP addresses.
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s ListRulesResponseBodyRulesRuleConditionsSourceIpConfig) String() string {
	return dara.Prettify(s)
}

func (s ListRulesResponseBodyRulesRuleConditionsSourceIpConfig) GoString() string {
	return s.String()
}

func (s *ListRulesResponseBodyRulesRuleConditionsSourceIpConfig) GetValues() []*string {
	return s.Values
}

func (s *ListRulesResponseBodyRulesRuleConditionsSourceIpConfig) SetValues(v []*string) *ListRulesResponseBodyRulesRuleConditionsSourceIpConfig {
	s.Values = v
	return s
}

func (s *ListRulesResponseBodyRulesRuleConditionsSourceIpConfig) Validate() error {
	return dara.Validate(s)
}

type ListRulesResponseBodyRulesTags struct {
	// The tag key. The tag key can be up to 128 characters in length. It cannot start with aliyun or acs: and cannot contain http:// or https://.
	//
	// example:
	//
	// env
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value. The tag value can be up to 128 characters in length. It cannot start with aliyun or acs: and cannot contain http:// or https://.
	//
	// example:
	//
	// product
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListRulesResponseBodyRulesTags) String() string {
	return dara.Prettify(s)
}

func (s ListRulesResponseBodyRulesTags) GoString() string {
	return s.String()
}

func (s *ListRulesResponseBodyRulesTags) GetKey() *string {
	return s.Key
}

func (s *ListRulesResponseBodyRulesTags) GetValue() *string {
	return s.Value
}

func (s *ListRulesResponseBodyRulesTags) SetKey(v string) *ListRulesResponseBodyRulesTags {
	s.Key = &v
	return s
}

func (s *ListRulesResponseBodyRulesTags) SetValue(v string) *ListRulesResponseBodyRulesTags {
	s.Value = &v
	return s
}

func (s *ListRulesResponseBodyRulesTags) Validate() error {
	return dara.Validate(s)
}
