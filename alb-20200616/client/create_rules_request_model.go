// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateRulesRequest
	GetClientToken() *string
	SetDryRun(v bool) *CreateRulesRequest
	GetDryRun() *bool
	SetListenerId(v string) *CreateRulesRequest
	GetListenerId() *string
	SetRules(v []*CreateRulesRequestRules) *CreateRulesRequest
	GetRules() []*CreateRulesRequestRules
}

type CreateRulesRequest struct {
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
	// 	- **false*	- (default): performs a dry run and sends the request. If the request passes the dry run, a `2xx HTTP` status code is returned and the operation is performed.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ID of the Application Load Balancer (ALB) instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// lsr-bp1bpn0kn908w4nbw****
	ListenerId *string `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	// This parameter is required.
	Rules []*CreateRulesRequestRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
}

func (s CreateRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateRulesRequest) GoString() string {
	return s.String()
}

func (s *CreateRulesRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateRulesRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *CreateRulesRequest) GetListenerId() *string {
	return s.ListenerId
}

func (s *CreateRulesRequest) GetRules() []*CreateRulesRequestRules {
	return s.Rules
}

func (s *CreateRulesRequest) SetClientToken(v string) *CreateRulesRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateRulesRequest) SetDryRun(v bool) *CreateRulesRequest {
	s.DryRun = &v
	return s
}

func (s *CreateRulesRequest) SetListenerId(v string) *CreateRulesRequest {
	s.ListenerId = &v
	return s
}

func (s *CreateRulesRequest) SetRules(v []*CreateRulesRequestRules) *CreateRulesRequest {
	s.Rules = v
	return s
}

func (s *CreateRulesRequest) Validate() error {
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

type CreateRulesRequestRules struct {
	Direction *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	// This parameter is required.
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// This parameter is required.
	RuleActions []*CreateRulesRequestRulesRuleActions `json:"RuleActions,omitempty" xml:"RuleActions,omitempty" type:"Repeated"`
	// This parameter is required.
	RuleConditions []*CreateRulesRequestRulesRuleConditions `json:"RuleConditions,omitempty" xml:"RuleConditions,omitempty" type:"Repeated"`
	// This parameter is required.
	RuleName *string                       `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	Tag      []*CreateRulesRequestRulesTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s CreateRulesRequestRules) String() string {
	return dara.Prettify(s)
}

func (s CreateRulesRequestRules) GoString() string {
	return s.String()
}

func (s *CreateRulesRequestRules) GetDirection() *string {
	return s.Direction
}

func (s *CreateRulesRequestRules) GetPriority() *int32 {
	return s.Priority
}

func (s *CreateRulesRequestRules) GetRuleActions() []*CreateRulesRequestRulesRuleActions {
	return s.RuleActions
}

func (s *CreateRulesRequestRules) GetRuleConditions() []*CreateRulesRequestRulesRuleConditions {
	return s.RuleConditions
}

func (s *CreateRulesRequestRules) GetRuleName() *string {
	return s.RuleName
}

func (s *CreateRulesRequestRules) GetTag() []*CreateRulesRequestRulesTag {
	return s.Tag
}

func (s *CreateRulesRequestRules) SetDirection(v string) *CreateRulesRequestRules {
	s.Direction = &v
	return s
}

func (s *CreateRulesRequestRules) SetPriority(v int32) *CreateRulesRequestRules {
	s.Priority = &v
	return s
}

func (s *CreateRulesRequestRules) SetRuleActions(v []*CreateRulesRequestRulesRuleActions) *CreateRulesRequestRules {
	s.RuleActions = v
	return s
}

func (s *CreateRulesRequestRules) SetRuleConditions(v []*CreateRulesRequestRulesRuleConditions) *CreateRulesRequestRules {
	s.RuleConditions = v
	return s
}

func (s *CreateRulesRequestRules) SetRuleName(v string) *CreateRulesRequestRules {
	s.RuleName = &v
	return s
}

func (s *CreateRulesRequestRules) SetTag(v []*CreateRulesRequestRulesTag) *CreateRulesRequestRules {
	s.Tag = v
	return s
}

func (s *CreateRulesRequestRules) Validate() error {
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

type CreateRulesRequestRulesRuleActions struct {
	CorsConfig          *CreateRulesRequestRulesRuleActionsCorsConfig          `json:"CorsConfig,omitempty" xml:"CorsConfig,omitempty" type:"Struct"`
	FixedResponseConfig *CreateRulesRequestRulesRuleActionsFixedResponseConfig `json:"FixedResponseConfig,omitempty" xml:"FixedResponseConfig,omitempty" type:"Struct"`
	ForwardGroupConfig  *CreateRulesRequestRulesRuleActionsForwardGroupConfig  `json:"ForwardGroupConfig,omitempty" xml:"ForwardGroupConfig,omitempty" type:"Struct"`
	InsertHeaderConfig  *CreateRulesRequestRulesRuleActionsInsertHeaderConfig  `json:"InsertHeaderConfig,omitempty" xml:"InsertHeaderConfig,omitempty" type:"Struct"`
	// This parameter is required.
	Order               *int32                                                 `json:"Order,omitempty" xml:"Order,omitempty"`
	RedirectConfig      *CreateRulesRequestRulesRuleActionsRedirectConfig      `json:"RedirectConfig,omitempty" xml:"RedirectConfig,omitempty" type:"Struct"`
	RemoveHeaderConfig  *CreateRulesRequestRulesRuleActionsRemoveHeaderConfig  `json:"RemoveHeaderConfig,omitempty" xml:"RemoveHeaderConfig,omitempty" type:"Struct"`
	RewriteConfig       *CreateRulesRequestRulesRuleActionsRewriteConfig       `json:"RewriteConfig,omitempty" xml:"RewriteConfig,omitempty" type:"Struct"`
	TrafficLimitConfig  *CreateRulesRequestRulesRuleActionsTrafficLimitConfig  `json:"TrafficLimitConfig,omitempty" xml:"TrafficLimitConfig,omitempty" type:"Struct"`
	TrafficMirrorConfig *CreateRulesRequestRulesRuleActionsTrafficMirrorConfig `json:"TrafficMirrorConfig,omitempty" xml:"TrafficMirrorConfig,omitempty" type:"Struct"`
	// This parameter is required.
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateRulesRequestRulesRuleActions) String() string {
	return dara.Prettify(s)
}

func (s CreateRulesRequestRulesRuleActions) GoString() string {
	return s.String()
}

func (s *CreateRulesRequestRulesRuleActions) GetCorsConfig() *CreateRulesRequestRulesRuleActionsCorsConfig {
	return s.CorsConfig
}

func (s *CreateRulesRequestRulesRuleActions) GetFixedResponseConfig() *CreateRulesRequestRulesRuleActionsFixedResponseConfig {
	return s.FixedResponseConfig
}

func (s *CreateRulesRequestRulesRuleActions) GetForwardGroupConfig() *CreateRulesRequestRulesRuleActionsForwardGroupConfig {
	return s.ForwardGroupConfig
}

func (s *CreateRulesRequestRulesRuleActions) GetInsertHeaderConfig() *CreateRulesRequestRulesRuleActionsInsertHeaderConfig {
	return s.InsertHeaderConfig
}

func (s *CreateRulesRequestRulesRuleActions) GetOrder() *int32 {
	return s.Order
}

func (s *CreateRulesRequestRulesRuleActions) GetRedirectConfig() *CreateRulesRequestRulesRuleActionsRedirectConfig {
	return s.RedirectConfig
}

func (s *CreateRulesRequestRulesRuleActions) GetRemoveHeaderConfig() *CreateRulesRequestRulesRuleActionsRemoveHeaderConfig {
	return s.RemoveHeaderConfig
}

func (s *CreateRulesRequestRulesRuleActions) GetRewriteConfig() *CreateRulesRequestRulesRuleActionsRewriteConfig {
	return s.RewriteConfig
}

func (s *CreateRulesRequestRulesRuleActions) GetTrafficLimitConfig() *CreateRulesRequestRulesRuleActionsTrafficLimitConfig {
	return s.TrafficLimitConfig
}

func (s *CreateRulesRequestRulesRuleActions) GetTrafficMirrorConfig() *CreateRulesRequestRulesRuleActionsTrafficMirrorConfig {
	return s.TrafficMirrorConfig
}

func (s *CreateRulesRequestRulesRuleActions) GetType() *string {
	return s.Type
}

func (s *CreateRulesRequestRulesRuleActions) SetCorsConfig(v *CreateRulesRequestRulesRuleActionsCorsConfig) *CreateRulesRequestRulesRuleActions {
	s.CorsConfig = v
	return s
}

func (s *CreateRulesRequestRulesRuleActions) SetFixedResponseConfig(v *CreateRulesRequestRulesRuleActionsFixedResponseConfig) *CreateRulesRequestRulesRuleActions {
	s.FixedResponseConfig = v
	return s
}

func (s *CreateRulesRequestRulesRuleActions) SetForwardGroupConfig(v *CreateRulesRequestRulesRuleActionsForwardGroupConfig) *CreateRulesRequestRulesRuleActions {
	s.ForwardGroupConfig = v
	return s
}

func (s *CreateRulesRequestRulesRuleActions) SetInsertHeaderConfig(v *CreateRulesRequestRulesRuleActionsInsertHeaderConfig) *CreateRulesRequestRulesRuleActions {
	s.InsertHeaderConfig = v
	return s
}

func (s *CreateRulesRequestRulesRuleActions) SetOrder(v int32) *CreateRulesRequestRulesRuleActions {
	s.Order = &v
	return s
}

func (s *CreateRulesRequestRulesRuleActions) SetRedirectConfig(v *CreateRulesRequestRulesRuleActionsRedirectConfig) *CreateRulesRequestRulesRuleActions {
	s.RedirectConfig = v
	return s
}

func (s *CreateRulesRequestRulesRuleActions) SetRemoveHeaderConfig(v *CreateRulesRequestRulesRuleActionsRemoveHeaderConfig) *CreateRulesRequestRulesRuleActions {
	s.RemoveHeaderConfig = v
	return s
}

func (s *CreateRulesRequestRulesRuleActions) SetRewriteConfig(v *CreateRulesRequestRulesRuleActionsRewriteConfig) *CreateRulesRequestRulesRuleActions {
	s.RewriteConfig = v
	return s
}

func (s *CreateRulesRequestRulesRuleActions) SetTrafficLimitConfig(v *CreateRulesRequestRulesRuleActionsTrafficLimitConfig) *CreateRulesRequestRulesRuleActions {
	s.TrafficLimitConfig = v
	return s
}

func (s *CreateRulesRequestRulesRuleActions) SetTrafficMirrorConfig(v *CreateRulesRequestRulesRuleActionsTrafficMirrorConfig) *CreateRulesRequestRulesRuleActions {
	s.TrafficMirrorConfig = v
	return s
}

func (s *CreateRulesRequestRulesRuleActions) SetType(v string) *CreateRulesRequestRulesRuleActions {
	s.Type = &v
	return s
}

func (s *CreateRulesRequestRulesRuleActions) Validate() error {
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

type CreateRulesRequestRulesRuleActionsCorsConfig struct {
	AllowCredentials *string   `json:"AllowCredentials,omitempty" xml:"AllowCredentials,omitempty"`
	AllowHeaders     []*string `json:"AllowHeaders,omitempty" xml:"AllowHeaders,omitempty" type:"Repeated"`
	AllowMethods     []*string `json:"AllowMethods,omitempty" xml:"AllowMethods,omitempty" type:"Repeated"`
	AllowOrigin      []*string `json:"AllowOrigin,omitempty" xml:"AllowOrigin,omitempty" type:"Repeated"`
	ExposeHeaders    []*string `json:"ExposeHeaders,omitempty" xml:"ExposeHeaders,omitempty" type:"Repeated"`
	MaxAge           *int64    `json:"MaxAge,omitempty" xml:"MaxAge,omitempty"`
}

func (s CreateRulesRequestRulesRuleActionsCorsConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateRulesRequestRulesRuleActionsCorsConfig) GoString() string {
	return s.String()
}

func (s *CreateRulesRequestRulesRuleActionsCorsConfig) GetAllowCredentials() *string {
	return s.AllowCredentials
}

func (s *CreateRulesRequestRulesRuleActionsCorsConfig) GetAllowHeaders() []*string {
	return s.AllowHeaders
}

func (s *CreateRulesRequestRulesRuleActionsCorsConfig) GetAllowMethods() []*string {
	return s.AllowMethods
}

func (s *CreateRulesRequestRulesRuleActionsCorsConfig) GetAllowOrigin() []*string {
	return s.AllowOrigin
}

func (s *CreateRulesRequestRulesRuleActionsCorsConfig) GetExposeHeaders() []*string {
	return s.ExposeHeaders
}

func (s *CreateRulesRequestRulesRuleActionsCorsConfig) GetMaxAge() *int64 {
	return s.MaxAge
}

func (s *CreateRulesRequestRulesRuleActionsCorsConfig) SetAllowCredentials(v string) *CreateRulesRequestRulesRuleActionsCorsConfig {
	s.AllowCredentials = &v
	return s
}

func (s *CreateRulesRequestRulesRuleActionsCorsConfig) SetAllowHeaders(v []*string) *CreateRulesRequestRulesRuleActionsCorsConfig {
	s.AllowHeaders = v
	return s
}

func (s *CreateRulesRequestRulesRuleActionsCorsConfig) SetAllowMethods(v []*string) *CreateRulesRequestRulesRuleActionsCorsConfig {
	s.AllowMethods = v
	return s
}

func (s *CreateRulesRequestRulesRuleActionsCorsConfig) SetAllowOrigin(v []*string) *CreateRulesRequestRulesRuleActionsCorsConfig {
	s.AllowOrigin = v
	return s
}

func (s *CreateRulesRequestRulesRuleActionsCorsConfig) SetExposeHeaders(v []*string) *CreateRulesRequestRulesRuleActionsCorsConfig {
	s.ExposeHeaders = v
	return s
}

func (s *CreateRulesRequestRulesRuleActionsCorsConfig) SetMaxAge(v int64) *CreateRulesRequestRulesRuleActionsCorsConfig {
	s.MaxAge = &v
	return s
}

func (s *CreateRulesRequestRulesRuleActionsCorsConfig) Validate() error {
	return dara.Validate(s)
}

type CreateRulesRequestRulesRuleActionsFixedResponseConfig struct {
	Content     *string `json:"Content,omitempty" xml:"Content,omitempty"`
	ContentType *string `json:"ContentType,omitempty" xml:"ContentType,omitempty"`
	HttpCode    *string `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
}

func (s CreateRulesRequestRulesRuleActionsFixedResponseConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateRulesRequestRulesRuleActionsFixedResponseConfig) GoString() string {
	return s.String()
}

func (s *CreateRulesRequestRulesRuleActionsFixedResponseConfig) GetContent() *string {
	return s.Content
}

func (s *CreateRulesRequestRulesRuleActionsFixedResponseConfig) GetContentType() *string {
	return s.ContentType
}

func (s *CreateRulesRequestRulesRuleActionsFixedResponseConfig) GetHttpCode() *string {
	return s.HttpCode
}

func (s *CreateRulesRequestRulesRuleActionsFixedResponseConfig) SetContent(v string) *CreateRulesRequestRulesRuleActionsFixedResponseConfig {
	s.Content = &v
	return s
}

func (s *CreateRulesRequestRulesRuleActionsFixedResponseConfig) SetContentType(v string) *CreateRulesRequestRulesRuleActionsFixedResponseConfig {
	s.ContentType = &v
	return s
}

func (s *CreateRulesRequestRulesRuleActionsFixedResponseConfig) SetHttpCode(v string) *CreateRulesRequestRulesRuleActionsFixedResponseConfig {
	s.HttpCode = &v
	return s
}

func (s *CreateRulesRequestRulesRuleActionsFixedResponseConfig) Validate() error {
	return dara.Validate(s)
}

type CreateRulesRequestRulesRuleActionsForwardGroupConfig struct {
	ServerGroupStickySession *CreateRulesRequestRulesRuleActionsForwardGroupConfigServerGroupStickySession `json:"ServerGroupStickySession,omitempty" xml:"ServerGroupStickySession,omitempty" type:"Struct"`
	ServerGroupTuples        []*CreateRulesRequestRulesRuleActionsForwardGroupConfigServerGroupTuples      `json:"ServerGroupTuples,omitempty" xml:"ServerGroupTuples,omitempty" type:"Repeated"`
}

func (s CreateRulesRequestRulesRuleActionsForwardGroupConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateRulesRequestRulesRuleActionsForwardGroupConfig) GoString() string {
	return s.String()
}

func (s *CreateRulesRequestRulesRuleActionsForwardGroupConfig) GetServerGroupStickySession() *CreateRulesRequestRulesRuleActionsForwardGroupConfigServerGroupStickySession {
	return s.ServerGroupStickySession
}

func (s *CreateRulesRequestRulesRuleActionsForwardGroupConfig) GetServerGroupTuples() []*CreateRulesRequestRulesRuleActionsForwardGroupConfigServerGroupTuples {
	return s.ServerGroupTuples
}

func (s *CreateRulesRequestRulesRuleActionsForwardGroupConfig) SetServerGroupStickySession(v *CreateRulesRequestRulesRuleActionsForwardGroupConfigServerGroupStickySession) *CreateRulesRequestRulesRuleActionsForwardGroupConfig {
	s.ServerGroupStickySession = v
	return s
}

func (s *CreateRulesRequestRulesRuleActionsForwardGroupConfig) SetServerGroupTuples(v []*CreateRulesRequestRulesRuleActionsForwardGroupConfigServerGroupTuples) *CreateRulesRequestRulesRuleActionsForwardGroupConfig {
	s.ServerGroupTuples = v
	return s
}

func (s *CreateRulesRequestRulesRuleActionsForwardGroupConfig) Validate() error {
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

type CreateRulesRequestRulesRuleActionsForwardGroupConfigServerGroupStickySession struct {
	Enabled *bool  `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	Timeout *int32 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
}

func (s CreateRulesRequestRulesRuleActionsForwardGroupConfigServerGroupStickySession) String() string {
	return dara.Prettify(s)
}

func (s CreateRulesRequestRulesRuleActionsForwardGroupConfigServerGroupStickySession) GoString() string {
	return s.String()
}

func (s *CreateRulesRequestRulesRuleActionsForwardGroupConfigServerGroupStickySession) GetEnabled() *bool {
	return s.Enabled
}

func (s *CreateRulesRequestRulesRuleActionsForwardGroupConfigServerGroupStickySession) GetTimeout() *int32 {
	return s.Timeout
}

func (s *CreateRulesRequestRulesRuleActionsForwardGroupConfigServerGroupStickySession) SetEnabled(v bool) *CreateRulesRequestRulesRuleActionsForwardGroupConfigServerGroupStickySession {
	s.Enabled = &v
	return s
}

func (s *CreateRulesRequestRulesRuleActionsForwardGroupConfigServerGroupStickySession) SetTimeout(v int32) *CreateRulesRequestRulesRuleActionsForwardGroupConfigServerGroupStickySession {
	s.Timeout = &v
	return s
}

func (s *CreateRulesRequestRulesRuleActionsForwardGroupConfigServerGroupStickySession) Validate() error {
	return dara.Validate(s)
}

type CreateRulesRequestRulesRuleActionsForwardGroupConfigServerGroupTuples struct {
	ServerGroupId *string `json:"ServerGroupId,omitempty" xml:"ServerGroupId,omitempty"`
	Weight        *int32  `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s CreateRulesRequestRulesRuleActionsForwardGroupConfigServerGroupTuples) String() string {
	return dara.Prettify(s)
}

func (s CreateRulesRequestRulesRuleActionsForwardGroupConfigServerGroupTuples) GoString() string {
	return s.String()
}

func (s *CreateRulesRequestRulesRuleActionsForwardGroupConfigServerGroupTuples) GetServerGroupId() *string {
	return s.ServerGroupId
}

func (s *CreateRulesRequestRulesRuleActionsForwardGroupConfigServerGroupTuples) GetWeight() *int32 {
	return s.Weight
}

func (s *CreateRulesRequestRulesRuleActionsForwardGroupConfigServerGroupTuples) SetServerGroupId(v string) *CreateRulesRequestRulesRuleActionsForwardGroupConfigServerGroupTuples {
	s.ServerGroupId = &v
	return s
}

func (s *CreateRulesRequestRulesRuleActionsForwardGroupConfigServerGroupTuples) SetWeight(v int32) *CreateRulesRequestRulesRuleActionsForwardGroupConfigServerGroupTuples {
	s.Weight = &v
	return s
}

func (s *CreateRulesRequestRulesRuleActionsForwardGroupConfigServerGroupTuples) Validate() error {
	return dara.Validate(s)
}

type CreateRulesRequestRulesRuleActionsInsertHeaderConfig struct {
	Key       *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value     *string `json:"Value,omitempty" xml:"Value,omitempty"`
	ValueType *string `json:"ValueType,omitempty" xml:"ValueType,omitempty"`
}

func (s CreateRulesRequestRulesRuleActionsInsertHeaderConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateRulesRequestRulesRuleActionsInsertHeaderConfig) GoString() string {
	return s.String()
}

func (s *CreateRulesRequestRulesRuleActionsInsertHeaderConfig) GetKey() *string {
	return s.Key
}

func (s *CreateRulesRequestRulesRuleActionsInsertHeaderConfig) GetValue() *string {
	return s.Value
}

func (s *CreateRulesRequestRulesRuleActionsInsertHeaderConfig) GetValueType() *string {
	return s.ValueType
}

func (s *CreateRulesRequestRulesRuleActionsInsertHeaderConfig) SetKey(v string) *CreateRulesRequestRulesRuleActionsInsertHeaderConfig {
	s.Key = &v
	return s
}

func (s *CreateRulesRequestRulesRuleActionsInsertHeaderConfig) SetValue(v string) *CreateRulesRequestRulesRuleActionsInsertHeaderConfig {
	s.Value = &v
	return s
}

func (s *CreateRulesRequestRulesRuleActionsInsertHeaderConfig) SetValueType(v string) *CreateRulesRequestRulesRuleActionsInsertHeaderConfig {
	s.ValueType = &v
	return s
}

func (s *CreateRulesRequestRulesRuleActionsInsertHeaderConfig) Validate() error {
	return dara.Validate(s)
}

type CreateRulesRequestRulesRuleActionsRedirectConfig struct {
	Host     *string `json:"Host,omitempty" xml:"Host,omitempty"`
	HttpCode *string `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	Path     *string `json:"Path,omitempty" xml:"Path,omitempty"`
	Port     *string `json:"Port,omitempty" xml:"Port,omitempty"`
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	Query    *string `json:"Query,omitempty" xml:"Query,omitempty"`
}

func (s CreateRulesRequestRulesRuleActionsRedirectConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateRulesRequestRulesRuleActionsRedirectConfig) GoString() string {
	return s.String()
}

func (s *CreateRulesRequestRulesRuleActionsRedirectConfig) GetHost() *string {
	return s.Host
}

func (s *CreateRulesRequestRulesRuleActionsRedirectConfig) GetHttpCode() *string {
	return s.HttpCode
}

func (s *CreateRulesRequestRulesRuleActionsRedirectConfig) GetPath() *string {
	return s.Path
}

func (s *CreateRulesRequestRulesRuleActionsRedirectConfig) GetPort() *string {
	return s.Port
}

func (s *CreateRulesRequestRulesRuleActionsRedirectConfig) GetProtocol() *string {
	return s.Protocol
}

func (s *CreateRulesRequestRulesRuleActionsRedirectConfig) GetQuery() *string {
	return s.Query
}

func (s *CreateRulesRequestRulesRuleActionsRedirectConfig) SetHost(v string) *CreateRulesRequestRulesRuleActionsRedirectConfig {
	s.Host = &v
	return s
}

func (s *CreateRulesRequestRulesRuleActionsRedirectConfig) SetHttpCode(v string) *CreateRulesRequestRulesRuleActionsRedirectConfig {
	s.HttpCode = &v
	return s
}

func (s *CreateRulesRequestRulesRuleActionsRedirectConfig) SetPath(v string) *CreateRulesRequestRulesRuleActionsRedirectConfig {
	s.Path = &v
	return s
}

func (s *CreateRulesRequestRulesRuleActionsRedirectConfig) SetPort(v string) *CreateRulesRequestRulesRuleActionsRedirectConfig {
	s.Port = &v
	return s
}

func (s *CreateRulesRequestRulesRuleActionsRedirectConfig) SetProtocol(v string) *CreateRulesRequestRulesRuleActionsRedirectConfig {
	s.Protocol = &v
	return s
}

func (s *CreateRulesRequestRulesRuleActionsRedirectConfig) SetQuery(v string) *CreateRulesRequestRulesRuleActionsRedirectConfig {
	s.Query = &v
	return s
}

func (s *CreateRulesRequestRulesRuleActionsRedirectConfig) Validate() error {
	return dara.Validate(s)
}

type CreateRulesRequestRulesRuleActionsRemoveHeaderConfig struct {
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
}

func (s CreateRulesRequestRulesRuleActionsRemoveHeaderConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateRulesRequestRulesRuleActionsRemoveHeaderConfig) GoString() string {
	return s.String()
}

func (s *CreateRulesRequestRulesRuleActionsRemoveHeaderConfig) GetKey() *string {
	return s.Key
}

func (s *CreateRulesRequestRulesRuleActionsRemoveHeaderConfig) SetKey(v string) *CreateRulesRequestRulesRuleActionsRemoveHeaderConfig {
	s.Key = &v
	return s
}

func (s *CreateRulesRequestRulesRuleActionsRemoveHeaderConfig) Validate() error {
	return dara.Validate(s)
}

type CreateRulesRequestRulesRuleActionsRewriteConfig struct {
	Host  *string `json:"Host,omitempty" xml:"Host,omitempty"`
	Path  *string `json:"Path,omitempty" xml:"Path,omitempty"`
	Query *string `json:"Query,omitempty" xml:"Query,omitempty"`
}

func (s CreateRulesRequestRulesRuleActionsRewriteConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateRulesRequestRulesRuleActionsRewriteConfig) GoString() string {
	return s.String()
}

func (s *CreateRulesRequestRulesRuleActionsRewriteConfig) GetHost() *string {
	return s.Host
}

func (s *CreateRulesRequestRulesRuleActionsRewriteConfig) GetPath() *string {
	return s.Path
}

func (s *CreateRulesRequestRulesRuleActionsRewriteConfig) GetQuery() *string {
	return s.Query
}

func (s *CreateRulesRequestRulesRuleActionsRewriteConfig) SetHost(v string) *CreateRulesRequestRulesRuleActionsRewriteConfig {
	s.Host = &v
	return s
}

func (s *CreateRulesRequestRulesRuleActionsRewriteConfig) SetPath(v string) *CreateRulesRequestRulesRuleActionsRewriteConfig {
	s.Path = &v
	return s
}

func (s *CreateRulesRequestRulesRuleActionsRewriteConfig) SetQuery(v string) *CreateRulesRequestRulesRuleActionsRewriteConfig {
	s.Query = &v
	return s
}

func (s *CreateRulesRequestRulesRuleActionsRewriteConfig) Validate() error {
	return dara.Validate(s)
}

type CreateRulesRequestRulesRuleActionsTrafficLimitConfig struct {
	PerIpQps *int32 `json:"PerIpQps,omitempty" xml:"PerIpQps,omitempty"`
	QPS      *int32 `json:"QPS,omitempty" xml:"QPS,omitempty"`
}

func (s CreateRulesRequestRulesRuleActionsTrafficLimitConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateRulesRequestRulesRuleActionsTrafficLimitConfig) GoString() string {
	return s.String()
}

func (s *CreateRulesRequestRulesRuleActionsTrafficLimitConfig) GetPerIpQps() *int32 {
	return s.PerIpQps
}

func (s *CreateRulesRequestRulesRuleActionsTrafficLimitConfig) GetQPS() *int32 {
	return s.QPS
}

func (s *CreateRulesRequestRulesRuleActionsTrafficLimitConfig) SetPerIpQps(v int32) *CreateRulesRequestRulesRuleActionsTrafficLimitConfig {
	s.PerIpQps = &v
	return s
}

func (s *CreateRulesRequestRulesRuleActionsTrafficLimitConfig) SetQPS(v int32) *CreateRulesRequestRulesRuleActionsTrafficLimitConfig {
	s.QPS = &v
	return s
}

func (s *CreateRulesRequestRulesRuleActionsTrafficLimitConfig) Validate() error {
	return dara.Validate(s)
}

type CreateRulesRequestRulesRuleActionsTrafficMirrorConfig struct {
	MirrorGroupConfig *CreateRulesRequestRulesRuleActionsTrafficMirrorConfigMirrorGroupConfig `json:"MirrorGroupConfig,omitempty" xml:"MirrorGroupConfig,omitempty" type:"Struct"`
	TargetType        *string                                                                 `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
}

func (s CreateRulesRequestRulesRuleActionsTrafficMirrorConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateRulesRequestRulesRuleActionsTrafficMirrorConfig) GoString() string {
	return s.String()
}

func (s *CreateRulesRequestRulesRuleActionsTrafficMirrorConfig) GetMirrorGroupConfig() *CreateRulesRequestRulesRuleActionsTrafficMirrorConfigMirrorGroupConfig {
	return s.MirrorGroupConfig
}

func (s *CreateRulesRequestRulesRuleActionsTrafficMirrorConfig) GetTargetType() *string {
	return s.TargetType
}

func (s *CreateRulesRequestRulesRuleActionsTrafficMirrorConfig) SetMirrorGroupConfig(v *CreateRulesRequestRulesRuleActionsTrafficMirrorConfigMirrorGroupConfig) *CreateRulesRequestRulesRuleActionsTrafficMirrorConfig {
	s.MirrorGroupConfig = v
	return s
}

func (s *CreateRulesRequestRulesRuleActionsTrafficMirrorConfig) SetTargetType(v string) *CreateRulesRequestRulesRuleActionsTrafficMirrorConfig {
	s.TargetType = &v
	return s
}

func (s *CreateRulesRequestRulesRuleActionsTrafficMirrorConfig) Validate() error {
	if s.MirrorGroupConfig != nil {
		if err := s.MirrorGroupConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateRulesRequestRulesRuleActionsTrafficMirrorConfigMirrorGroupConfig struct {
	ServerGroupTuples []*CreateRulesRequestRulesRuleActionsTrafficMirrorConfigMirrorGroupConfigServerGroupTuples `json:"ServerGroupTuples,omitempty" xml:"ServerGroupTuples,omitempty" type:"Repeated"`
}

func (s CreateRulesRequestRulesRuleActionsTrafficMirrorConfigMirrorGroupConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateRulesRequestRulesRuleActionsTrafficMirrorConfigMirrorGroupConfig) GoString() string {
	return s.String()
}

func (s *CreateRulesRequestRulesRuleActionsTrafficMirrorConfigMirrorGroupConfig) GetServerGroupTuples() []*CreateRulesRequestRulesRuleActionsTrafficMirrorConfigMirrorGroupConfigServerGroupTuples {
	return s.ServerGroupTuples
}

func (s *CreateRulesRequestRulesRuleActionsTrafficMirrorConfigMirrorGroupConfig) SetServerGroupTuples(v []*CreateRulesRequestRulesRuleActionsTrafficMirrorConfigMirrorGroupConfigServerGroupTuples) *CreateRulesRequestRulesRuleActionsTrafficMirrorConfigMirrorGroupConfig {
	s.ServerGroupTuples = v
	return s
}

func (s *CreateRulesRequestRulesRuleActionsTrafficMirrorConfigMirrorGroupConfig) Validate() error {
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

type CreateRulesRequestRulesRuleActionsTrafficMirrorConfigMirrorGroupConfigServerGroupTuples struct {
	ServerGroupId *string `json:"ServerGroupId,omitempty" xml:"ServerGroupId,omitempty"`
}

func (s CreateRulesRequestRulesRuleActionsTrafficMirrorConfigMirrorGroupConfigServerGroupTuples) String() string {
	return dara.Prettify(s)
}

func (s CreateRulesRequestRulesRuleActionsTrafficMirrorConfigMirrorGroupConfigServerGroupTuples) GoString() string {
	return s.String()
}

func (s *CreateRulesRequestRulesRuleActionsTrafficMirrorConfigMirrorGroupConfigServerGroupTuples) GetServerGroupId() *string {
	return s.ServerGroupId
}

func (s *CreateRulesRequestRulesRuleActionsTrafficMirrorConfigMirrorGroupConfigServerGroupTuples) SetServerGroupId(v string) *CreateRulesRequestRulesRuleActionsTrafficMirrorConfigMirrorGroupConfigServerGroupTuples {
	s.ServerGroupId = &v
	return s
}

func (s *CreateRulesRequestRulesRuleActionsTrafficMirrorConfigMirrorGroupConfigServerGroupTuples) Validate() error {
	return dara.Validate(s)
}

type CreateRulesRequestRulesRuleConditions struct {
	CookieConfig             *CreateRulesRequestRulesRuleConditionsCookieConfig             `json:"CookieConfig,omitempty" xml:"CookieConfig,omitempty" type:"Struct"`
	HeaderConfig             *CreateRulesRequestRulesRuleConditionsHeaderConfig             `json:"HeaderConfig,omitempty" xml:"HeaderConfig,omitempty" type:"Struct"`
	HostConfig               *CreateRulesRequestRulesRuleConditionsHostConfig               `json:"HostConfig,omitempty" xml:"HostConfig,omitempty" type:"Struct"`
	MethodConfig             *CreateRulesRequestRulesRuleConditionsMethodConfig             `json:"MethodConfig,omitempty" xml:"MethodConfig,omitempty" type:"Struct"`
	PathConfig               *CreateRulesRequestRulesRuleConditionsPathConfig               `json:"PathConfig,omitempty" xml:"PathConfig,omitempty" type:"Struct"`
	QueryStringConfig        *CreateRulesRequestRulesRuleConditionsQueryStringConfig        `json:"QueryStringConfig,omitempty" xml:"QueryStringConfig,omitempty" type:"Struct"`
	ResponseHeaderConfig     *CreateRulesRequestRulesRuleConditionsResponseHeaderConfig     `json:"ResponseHeaderConfig,omitempty" xml:"ResponseHeaderConfig,omitempty" type:"Struct"`
	ResponseStatusCodeConfig *CreateRulesRequestRulesRuleConditionsResponseStatusCodeConfig `json:"ResponseStatusCodeConfig,omitempty" xml:"ResponseStatusCodeConfig,omitempty" type:"Struct"`
	SourceIpConfig           *CreateRulesRequestRulesRuleConditionsSourceIpConfig           `json:"SourceIpConfig,omitempty" xml:"SourceIpConfig,omitempty" type:"Struct"`
	// This parameter is required.
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateRulesRequestRulesRuleConditions) String() string {
	return dara.Prettify(s)
}

func (s CreateRulesRequestRulesRuleConditions) GoString() string {
	return s.String()
}

func (s *CreateRulesRequestRulesRuleConditions) GetCookieConfig() *CreateRulesRequestRulesRuleConditionsCookieConfig {
	return s.CookieConfig
}

func (s *CreateRulesRequestRulesRuleConditions) GetHeaderConfig() *CreateRulesRequestRulesRuleConditionsHeaderConfig {
	return s.HeaderConfig
}

func (s *CreateRulesRequestRulesRuleConditions) GetHostConfig() *CreateRulesRequestRulesRuleConditionsHostConfig {
	return s.HostConfig
}

func (s *CreateRulesRequestRulesRuleConditions) GetMethodConfig() *CreateRulesRequestRulesRuleConditionsMethodConfig {
	return s.MethodConfig
}

func (s *CreateRulesRequestRulesRuleConditions) GetPathConfig() *CreateRulesRequestRulesRuleConditionsPathConfig {
	return s.PathConfig
}

func (s *CreateRulesRequestRulesRuleConditions) GetQueryStringConfig() *CreateRulesRequestRulesRuleConditionsQueryStringConfig {
	return s.QueryStringConfig
}

func (s *CreateRulesRequestRulesRuleConditions) GetResponseHeaderConfig() *CreateRulesRequestRulesRuleConditionsResponseHeaderConfig {
	return s.ResponseHeaderConfig
}

func (s *CreateRulesRequestRulesRuleConditions) GetResponseStatusCodeConfig() *CreateRulesRequestRulesRuleConditionsResponseStatusCodeConfig {
	return s.ResponseStatusCodeConfig
}

func (s *CreateRulesRequestRulesRuleConditions) GetSourceIpConfig() *CreateRulesRequestRulesRuleConditionsSourceIpConfig {
	return s.SourceIpConfig
}

func (s *CreateRulesRequestRulesRuleConditions) GetType() *string {
	return s.Type
}

func (s *CreateRulesRequestRulesRuleConditions) SetCookieConfig(v *CreateRulesRequestRulesRuleConditionsCookieConfig) *CreateRulesRequestRulesRuleConditions {
	s.CookieConfig = v
	return s
}

func (s *CreateRulesRequestRulesRuleConditions) SetHeaderConfig(v *CreateRulesRequestRulesRuleConditionsHeaderConfig) *CreateRulesRequestRulesRuleConditions {
	s.HeaderConfig = v
	return s
}

func (s *CreateRulesRequestRulesRuleConditions) SetHostConfig(v *CreateRulesRequestRulesRuleConditionsHostConfig) *CreateRulesRequestRulesRuleConditions {
	s.HostConfig = v
	return s
}

func (s *CreateRulesRequestRulesRuleConditions) SetMethodConfig(v *CreateRulesRequestRulesRuleConditionsMethodConfig) *CreateRulesRequestRulesRuleConditions {
	s.MethodConfig = v
	return s
}

func (s *CreateRulesRequestRulesRuleConditions) SetPathConfig(v *CreateRulesRequestRulesRuleConditionsPathConfig) *CreateRulesRequestRulesRuleConditions {
	s.PathConfig = v
	return s
}

func (s *CreateRulesRequestRulesRuleConditions) SetQueryStringConfig(v *CreateRulesRequestRulesRuleConditionsQueryStringConfig) *CreateRulesRequestRulesRuleConditions {
	s.QueryStringConfig = v
	return s
}

func (s *CreateRulesRequestRulesRuleConditions) SetResponseHeaderConfig(v *CreateRulesRequestRulesRuleConditionsResponseHeaderConfig) *CreateRulesRequestRulesRuleConditions {
	s.ResponseHeaderConfig = v
	return s
}

func (s *CreateRulesRequestRulesRuleConditions) SetResponseStatusCodeConfig(v *CreateRulesRequestRulesRuleConditionsResponseStatusCodeConfig) *CreateRulesRequestRulesRuleConditions {
	s.ResponseStatusCodeConfig = v
	return s
}

func (s *CreateRulesRequestRulesRuleConditions) SetSourceIpConfig(v *CreateRulesRequestRulesRuleConditionsSourceIpConfig) *CreateRulesRequestRulesRuleConditions {
	s.SourceIpConfig = v
	return s
}

func (s *CreateRulesRequestRulesRuleConditions) SetType(v string) *CreateRulesRequestRulesRuleConditions {
	s.Type = &v
	return s
}

func (s *CreateRulesRequestRulesRuleConditions) Validate() error {
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

type CreateRulesRequestRulesRuleConditionsCookieConfig struct {
	Values []*CreateRulesRequestRulesRuleConditionsCookieConfigValues `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s CreateRulesRequestRulesRuleConditionsCookieConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateRulesRequestRulesRuleConditionsCookieConfig) GoString() string {
	return s.String()
}

func (s *CreateRulesRequestRulesRuleConditionsCookieConfig) GetValues() []*CreateRulesRequestRulesRuleConditionsCookieConfigValues {
	return s.Values
}

func (s *CreateRulesRequestRulesRuleConditionsCookieConfig) SetValues(v []*CreateRulesRequestRulesRuleConditionsCookieConfigValues) *CreateRulesRequestRulesRuleConditionsCookieConfig {
	s.Values = v
	return s
}

func (s *CreateRulesRequestRulesRuleConditionsCookieConfig) Validate() error {
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

type CreateRulesRequestRulesRuleConditionsCookieConfigValues struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateRulesRequestRulesRuleConditionsCookieConfigValues) String() string {
	return dara.Prettify(s)
}

func (s CreateRulesRequestRulesRuleConditionsCookieConfigValues) GoString() string {
	return s.String()
}

func (s *CreateRulesRequestRulesRuleConditionsCookieConfigValues) GetKey() *string {
	return s.Key
}

func (s *CreateRulesRequestRulesRuleConditionsCookieConfigValues) GetValue() *string {
	return s.Value
}

func (s *CreateRulesRequestRulesRuleConditionsCookieConfigValues) SetKey(v string) *CreateRulesRequestRulesRuleConditionsCookieConfigValues {
	s.Key = &v
	return s
}

func (s *CreateRulesRequestRulesRuleConditionsCookieConfigValues) SetValue(v string) *CreateRulesRequestRulesRuleConditionsCookieConfigValues {
	s.Value = &v
	return s
}

func (s *CreateRulesRequestRulesRuleConditionsCookieConfigValues) Validate() error {
	return dara.Validate(s)
}

type CreateRulesRequestRulesRuleConditionsHeaderConfig struct {
	Key    *string   `json:"Key,omitempty" xml:"Key,omitempty"`
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s CreateRulesRequestRulesRuleConditionsHeaderConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateRulesRequestRulesRuleConditionsHeaderConfig) GoString() string {
	return s.String()
}

func (s *CreateRulesRequestRulesRuleConditionsHeaderConfig) GetKey() *string {
	return s.Key
}

func (s *CreateRulesRequestRulesRuleConditionsHeaderConfig) GetValues() []*string {
	return s.Values
}

func (s *CreateRulesRequestRulesRuleConditionsHeaderConfig) SetKey(v string) *CreateRulesRequestRulesRuleConditionsHeaderConfig {
	s.Key = &v
	return s
}

func (s *CreateRulesRequestRulesRuleConditionsHeaderConfig) SetValues(v []*string) *CreateRulesRequestRulesRuleConditionsHeaderConfig {
	s.Values = v
	return s
}

func (s *CreateRulesRequestRulesRuleConditionsHeaderConfig) Validate() error {
	return dara.Validate(s)
}

type CreateRulesRequestRulesRuleConditionsHostConfig struct {
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s CreateRulesRequestRulesRuleConditionsHostConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateRulesRequestRulesRuleConditionsHostConfig) GoString() string {
	return s.String()
}

func (s *CreateRulesRequestRulesRuleConditionsHostConfig) GetValues() []*string {
	return s.Values
}

func (s *CreateRulesRequestRulesRuleConditionsHostConfig) SetValues(v []*string) *CreateRulesRequestRulesRuleConditionsHostConfig {
	s.Values = v
	return s
}

func (s *CreateRulesRequestRulesRuleConditionsHostConfig) Validate() error {
	return dara.Validate(s)
}

type CreateRulesRequestRulesRuleConditionsMethodConfig struct {
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s CreateRulesRequestRulesRuleConditionsMethodConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateRulesRequestRulesRuleConditionsMethodConfig) GoString() string {
	return s.String()
}

func (s *CreateRulesRequestRulesRuleConditionsMethodConfig) GetValues() []*string {
	return s.Values
}

func (s *CreateRulesRequestRulesRuleConditionsMethodConfig) SetValues(v []*string) *CreateRulesRequestRulesRuleConditionsMethodConfig {
	s.Values = v
	return s
}

func (s *CreateRulesRequestRulesRuleConditionsMethodConfig) Validate() error {
	return dara.Validate(s)
}

type CreateRulesRequestRulesRuleConditionsPathConfig struct {
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s CreateRulesRequestRulesRuleConditionsPathConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateRulesRequestRulesRuleConditionsPathConfig) GoString() string {
	return s.String()
}

func (s *CreateRulesRequestRulesRuleConditionsPathConfig) GetValues() []*string {
	return s.Values
}

func (s *CreateRulesRequestRulesRuleConditionsPathConfig) SetValues(v []*string) *CreateRulesRequestRulesRuleConditionsPathConfig {
	s.Values = v
	return s
}

func (s *CreateRulesRequestRulesRuleConditionsPathConfig) Validate() error {
	return dara.Validate(s)
}

type CreateRulesRequestRulesRuleConditionsQueryStringConfig struct {
	Values []*CreateRulesRequestRulesRuleConditionsQueryStringConfigValues `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s CreateRulesRequestRulesRuleConditionsQueryStringConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateRulesRequestRulesRuleConditionsQueryStringConfig) GoString() string {
	return s.String()
}

func (s *CreateRulesRequestRulesRuleConditionsQueryStringConfig) GetValues() []*CreateRulesRequestRulesRuleConditionsQueryStringConfigValues {
	return s.Values
}

func (s *CreateRulesRequestRulesRuleConditionsQueryStringConfig) SetValues(v []*CreateRulesRequestRulesRuleConditionsQueryStringConfigValues) *CreateRulesRequestRulesRuleConditionsQueryStringConfig {
	s.Values = v
	return s
}

func (s *CreateRulesRequestRulesRuleConditionsQueryStringConfig) Validate() error {
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

type CreateRulesRequestRulesRuleConditionsQueryStringConfigValues struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateRulesRequestRulesRuleConditionsQueryStringConfigValues) String() string {
	return dara.Prettify(s)
}

func (s CreateRulesRequestRulesRuleConditionsQueryStringConfigValues) GoString() string {
	return s.String()
}

func (s *CreateRulesRequestRulesRuleConditionsQueryStringConfigValues) GetKey() *string {
	return s.Key
}

func (s *CreateRulesRequestRulesRuleConditionsQueryStringConfigValues) GetValue() *string {
	return s.Value
}

func (s *CreateRulesRequestRulesRuleConditionsQueryStringConfigValues) SetKey(v string) *CreateRulesRequestRulesRuleConditionsQueryStringConfigValues {
	s.Key = &v
	return s
}

func (s *CreateRulesRequestRulesRuleConditionsQueryStringConfigValues) SetValue(v string) *CreateRulesRequestRulesRuleConditionsQueryStringConfigValues {
	s.Value = &v
	return s
}

func (s *CreateRulesRequestRulesRuleConditionsQueryStringConfigValues) Validate() error {
	return dara.Validate(s)
}

type CreateRulesRequestRulesRuleConditionsResponseHeaderConfig struct {
	Key    *string   `json:"Key,omitempty" xml:"Key,omitempty"`
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s CreateRulesRequestRulesRuleConditionsResponseHeaderConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateRulesRequestRulesRuleConditionsResponseHeaderConfig) GoString() string {
	return s.String()
}

func (s *CreateRulesRequestRulesRuleConditionsResponseHeaderConfig) GetKey() *string {
	return s.Key
}

func (s *CreateRulesRequestRulesRuleConditionsResponseHeaderConfig) GetValues() []*string {
	return s.Values
}

func (s *CreateRulesRequestRulesRuleConditionsResponseHeaderConfig) SetKey(v string) *CreateRulesRequestRulesRuleConditionsResponseHeaderConfig {
	s.Key = &v
	return s
}

func (s *CreateRulesRequestRulesRuleConditionsResponseHeaderConfig) SetValues(v []*string) *CreateRulesRequestRulesRuleConditionsResponseHeaderConfig {
	s.Values = v
	return s
}

func (s *CreateRulesRequestRulesRuleConditionsResponseHeaderConfig) Validate() error {
	return dara.Validate(s)
}

type CreateRulesRequestRulesRuleConditionsResponseStatusCodeConfig struct {
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s CreateRulesRequestRulesRuleConditionsResponseStatusCodeConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateRulesRequestRulesRuleConditionsResponseStatusCodeConfig) GoString() string {
	return s.String()
}

func (s *CreateRulesRequestRulesRuleConditionsResponseStatusCodeConfig) GetValues() []*string {
	return s.Values
}

func (s *CreateRulesRequestRulesRuleConditionsResponseStatusCodeConfig) SetValues(v []*string) *CreateRulesRequestRulesRuleConditionsResponseStatusCodeConfig {
	s.Values = v
	return s
}

func (s *CreateRulesRequestRulesRuleConditionsResponseStatusCodeConfig) Validate() error {
	return dara.Validate(s)
}

type CreateRulesRequestRulesRuleConditionsSourceIpConfig struct {
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s CreateRulesRequestRulesRuleConditionsSourceIpConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateRulesRequestRulesRuleConditionsSourceIpConfig) GoString() string {
	return s.String()
}

func (s *CreateRulesRequestRulesRuleConditionsSourceIpConfig) GetValues() []*string {
	return s.Values
}

func (s *CreateRulesRequestRulesRuleConditionsSourceIpConfig) SetValues(v []*string) *CreateRulesRequestRulesRuleConditionsSourceIpConfig {
	s.Values = v
	return s
}

func (s *CreateRulesRequestRulesRuleConditionsSourceIpConfig) Validate() error {
	return dara.Validate(s)
}

type CreateRulesRequestRulesTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateRulesRequestRulesTag) String() string {
	return dara.Prettify(s)
}

func (s CreateRulesRequestRulesTag) GoString() string {
	return s.String()
}

func (s *CreateRulesRequestRulesTag) GetKey() *string {
	return s.Key
}

func (s *CreateRulesRequestRulesTag) GetValue() *string {
	return s.Value
}

func (s *CreateRulesRequestRulesTag) SetKey(v string) *CreateRulesRequestRulesTag {
	s.Key = &v
	return s
}

func (s *CreateRulesRequestRulesTag) SetValue(v string) *CreateRulesRequestRulesTag {
	s.Value = &v
	return s
}

func (s *CreateRulesRequestRulesTag) Validate() error {
	return dara.Validate(s)
}
