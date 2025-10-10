// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRulesAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *UpdateRulesAttributeRequest
	GetClientToken() *string
	SetDryRun(v bool) *UpdateRulesAttributeRequest
	GetDryRun() *bool
	SetRules(v []*UpdateRulesAttributeRequestRules) *UpdateRulesAttributeRequest
	GetRules() []*UpdateRulesAttributeRequestRules
}

type UpdateRulesAttributeRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The client token can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// 	- **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// 	- **false*	- (default): performs a dry run and performs the actual request. If the request passes the dry run, a `2xx HTTP` status code is returned and the operation is performed.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// This parameter is required.
	Rules []*UpdateRulesAttributeRequestRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
}

func (s UpdateRulesAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateRulesAttributeRequest) GoString() string {
	return s.String()
}

func (s *UpdateRulesAttributeRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateRulesAttributeRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *UpdateRulesAttributeRequest) GetRules() []*UpdateRulesAttributeRequestRules {
	return s.Rules
}

func (s *UpdateRulesAttributeRequest) SetClientToken(v string) *UpdateRulesAttributeRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateRulesAttributeRequest) SetDryRun(v bool) *UpdateRulesAttributeRequest {
	s.DryRun = &v
	return s
}

func (s *UpdateRulesAttributeRequest) SetRules(v []*UpdateRulesAttributeRequestRules) *UpdateRulesAttributeRequest {
	s.Rules = v
	return s
}

func (s *UpdateRulesAttributeRequest) Validate() error {
	return dara.Validate(s)
}

type UpdateRulesAttributeRequestRules struct {
	Priority       *int32                                            `json:"Priority,omitempty" xml:"Priority,omitempty"`
	RuleActions    []*UpdateRulesAttributeRequestRulesRuleActions    `json:"RuleActions,omitempty" xml:"RuleActions,omitempty" type:"Repeated"`
	RuleConditions []*UpdateRulesAttributeRequestRulesRuleConditions `json:"RuleConditions,omitempty" xml:"RuleConditions,omitempty" type:"Repeated"`
	// This parameter is required.
	RuleId   *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
}

func (s UpdateRulesAttributeRequestRules) String() string {
	return dara.Prettify(s)
}

func (s UpdateRulesAttributeRequestRules) GoString() string {
	return s.String()
}

func (s *UpdateRulesAttributeRequestRules) GetPriority() *int32 {
	return s.Priority
}

func (s *UpdateRulesAttributeRequestRules) GetRuleActions() []*UpdateRulesAttributeRequestRulesRuleActions {
	return s.RuleActions
}

func (s *UpdateRulesAttributeRequestRules) GetRuleConditions() []*UpdateRulesAttributeRequestRulesRuleConditions {
	return s.RuleConditions
}

func (s *UpdateRulesAttributeRequestRules) GetRuleId() *string {
	return s.RuleId
}

func (s *UpdateRulesAttributeRequestRules) GetRuleName() *string {
	return s.RuleName
}

func (s *UpdateRulesAttributeRequestRules) SetPriority(v int32) *UpdateRulesAttributeRequestRules {
	s.Priority = &v
	return s
}

func (s *UpdateRulesAttributeRequestRules) SetRuleActions(v []*UpdateRulesAttributeRequestRulesRuleActions) *UpdateRulesAttributeRequestRules {
	s.RuleActions = v
	return s
}

func (s *UpdateRulesAttributeRequestRules) SetRuleConditions(v []*UpdateRulesAttributeRequestRulesRuleConditions) *UpdateRulesAttributeRequestRules {
	s.RuleConditions = v
	return s
}

func (s *UpdateRulesAttributeRequestRules) SetRuleId(v string) *UpdateRulesAttributeRequestRules {
	s.RuleId = &v
	return s
}

func (s *UpdateRulesAttributeRequestRules) SetRuleName(v string) *UpdateRulesAttributeRequestRules {
	s.RuleName = &v
	return s
}

func (s *UpdateRulesAttributeRequestRules) Validate() error {
	return dara.Validate(s)
}

type UpdateRulesAttributeRequestRulesRuleActions struct {
	CorsConfig          *UpdateRulesAttributeRequestRulesRuleActionsCorsConfig          `json:"CorsConfig,omitempty" xml:"CorsConfig,omitempty" type:"Struct"`
	FixedResponseConfig *UpdateRulesAttributeRequestRulesRuleActionsFixedResponseConfig `json:"FixedResponseConfig,omitempty" xml:"FixedResponseConfig,omitempty" type:"Struct"`
	ForwardGroupConfig  *UpdateRulesAttributeRequestRulesRuleActionsForwardGroupConfig  `json:"ForwardGroupConfig,omitempty" xml:"ForwardGroupConfig,omitempty" type:"Struct"`
	InsertHeaderConfig  *UpdateRulesAttributeRequestRulesRuleActionsInsertHeaderConfig  `json:"InsertHeaderConfig,omitempty" xml:"InsertHeaderConfig,omitempty" type:"Struct"`
	// This parameter is required.
	Order               *int32                                                          `json:"Order,omitempty" xml:"Order,omitempty"`
	RedirectConfig      *UpdateRulesAttributeRequestRulesRuleActionsRedirectConfig      `json:"RedirectConfig,omitempty" xml:"RedirectConfig,omitempty" type:"Struct"`
	RemoveHeaderConfig  *UpdateRulesAttributeRequestRulesRuleActionsRemoveHeaderConfig  `json:"RemoveHeaderConfig,omitempty" xml:"RemoveHeaderConfig,omitempty" type:"Struct"`
	RewriteConfig       *UpdateRulesAttributeRequestRulesRuleActionsRewriteConfig       `json:"RewriteConfig,omitempty" xml:"RewriteConfig,omitempty" type:"Struct"`
	TrafficLimitConfig  *UpdateRulesAttributeRequestRulesRuleActionsTrafficLimitConfig  `json:"TrafficLimitConfig,omitempty" xml:"TrafficLimitConfig,omitempty" type:"Struct"`
	TrafficMirrorConfig *UpdateRulesAttributeRequestRulesRuleActionsTrafficMirrorConfig `json:"TrafficMirrorConfig,omitempty" xml:"TrafficMirrorConfig,omitempty" type:"Struct"`
	// This parameter is required.
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s UpdateRulesAttributeRequestRulesRuleActions) String() string {
	return dara.Prettify(s)
}

func (s UpdateRulesAttributeRequestRulesRuleActions) GoString() string {
	return s.String()
}

func (s *UpdateRulesAttributeRequestRulesRuleActions) GetCorsConfig() *UpdateRulesAttributeRequestRulesRuleActionsCorsConfig {
	return s.CorsConfig
}

func (s *UpdateRulesAttributeRequestRulesRuleActions) GetFixedResponseConfig() *UpdateRulesAttributeRequestRulesRuleActionsFixedResponseConfig {
	return s.FixedResponseConfig
}

func (s *UpdateRulesAttributeRequestRulesRuleActions) GetForwardGroupConfig() *UpdateRulesAttributeRequestRulesRuleActionsForwardGroupConfig {
	return s.ForwardGroupConfig
}

func (s *UpdateRulesAttributeRequestRulesRuleActions) GetInsertHeaderConfig() *UpdateRulesAttributeRequestRulesRuleActionsInsertHeaderConfig {
	return s.InsertHeaderConfig
}

func (s *UpdateRulesAttributeRequestRulesRuleActions) GetOrder() *int32 {
	return s.Order
}

func (s *UpdateRulesAttributeRequestRulesRuleActions) GetRedirectConfig() *UpdateRulesAttributeRequestRulesRuleActionsRedirectConfig {
	return s.RedirectConfig
}

func (s *UpdateRulesAttributeRequestRulesRuleActions) GetRemoveHeaderConfig() *UpdateRulesAttributeRequestRulesRuleActionsRemoveHeaderConfig {
	return s.RemoveHeaderConfig
}

func (s *UpdateRulesAttributeRequestRulesRuleActions) GetRewriteConfig() *UpdateRulesAttributeRequestRulesRuleActionsRewriteConfig {
	return s.RewriteConfig
}

func (s *UpdateRulesAttributeRequestRulesRuleActions) GetTrafficLimitConfig() *UpdateRulesAttributeRequestRulesRuleActionsTrafficLimitConfig {
	return s.TrafficLimitConfig
}

func (s *UpdateRulesAttributeRequestRulesRuleActions) GetTrafficMirrorConfig() *UpdateRulesAttributeRequestRulesRuleActionsTrafficMirrorConfig {
	return s.TrafficMirrorConfig
}

func (s *UpdateRulesAttributeRequestRulesRuleActions) GetType() *string {
	return s.Type
}

func (s *UpdateRulesAttributeRequestRulesRuleActions) SetCorsConfig(v *UpdateRulesAttributeRequestRulesRuleActionsCorsConfig) *UpdateRulesAttributeRequestRulesRuleActions {
	s.CorsConfig = v
	return s
}

func (s *UpdateRulesAttributeRequestRulesRuleActions) SetFixedResponseConfig(v *UpdateRulesAttributeRequestRulesRuleActionsFixedResponseConfig) *UpdateRulesAttributeRequestRulesRuleActions {
	s.FixedResponseConfig = v
	return s
}

func (s *UpdateRulesAttributeRequestRulesRuleActions) SetForwardGroupConfig(v *UpdateRulesAttributeRequestRulesRuleActionsForwardGroupConfig) *UpdateRulesAttributeRequestRulesRuleActions {
	s.ForwardGroupConfig = v
	return s
}

func (s *UpdateRulesAttributeRequestRulesRuleActions) SetInsertHeaderConfig(v *UpdateRulesAttributeRequestRulesRuleActionsInsertHeaderConfig) *UpdateRulesAttributeRequestRulesRuleActions {
	s.InsertHeaderConfig = v
	return s
}

func (s *UpdateRulesAttributeRequestRulesRuleActions) SetOrder(v int32) *UpdateRulesAttributeRequestRulesRuleActions {
	s.Order = &v
	return s
}

func (s *UpdateRulesAttributeRequestRulesRuleActions) SetRedirectConfig(v *UpdateRulesAttributeRequestRulesRuleActionsRedirectConfig) *UpdateRulesAttributeRequestRulesRuleActions {
	s.RedirectConfig = v
	return s
}

func (s *UpdateRulesAttributeRequestRulesRuleActions) SetRemoveHeaderConfig(v *UpdateRulesAttributeRequestRulesRuleActionsRemoveHeaderConfig) *UpdateRulesAttributeRequestRulesRuleActions {
	s.RemoveHeaderConfig = v
	return s
}

func (s *UpdateRulesAttributeRequestRulesRuleActions) SetRewriteConfig(v *UpdateRulesAttributeRequestRulesRuleActionsRewriteConfig) *UpdateRulesAttributeRequestRulesRuleActions {
	s.RewriteConfig = v
	return s
}

func (s *UpdateRulesAttributeRequestRulesRuleActions) SetTrafficLimitConfig(v *UpdateRulesAttributeRequestRulesRuleActionsTrafficLimitConfig) *UpdateRulesAttributeRequestRulesRuleActions {
	s.TrafficLimitConfig = v
	return s
}

func (s *UpdateRulesAttributeRequestRulesRuleActions) SetTrafficMirrorConfig(v *UpdateRulesAttributeRequestRulesRuleActionsTrafficMirrorConfig) *UpdateRulesAttributeRequestRulesRuleActions {
	s.TrafficMirrorConfig = v
	return s
}

func (s *UpdateRulesAttributeRequestRulesRuleActions) SetType(v string) *UpdateRulesAttributeRequestRulesRuleActions {
	s.Type = &v
	return s
}

func (s *UpdateRulesAttributeRequestRulesRuleActions) Validate() error {
	return dara.Validate(s)
}

type UpdateRulesAttributeRequestRulesRuleActionsCorsConfig struct {
	AllowCredentials *string   `json:"AllowCredentials,omitempty" xml:"AllowCredentials,omitempty"`
	AllowHeaders     []*string `json:"AllowHeaders,omitempty" xml:"AllowHeaders,omitempty" type:"Repeated"`
	AllowMethods     []*string `json:"AllowMethods,omitempty" xml:"AllowMethods,omitempty" type:"Repeated"`
	AllowOrigin      []*string `json:"AllowOrigin,omitempty" xml:"AllowOrigin,omitempty" type:"Repeated"`
	ExposeHeaders    []*string `json:"ExposeHeaders,omitempty" xml:"ExposeHeaders,omitempty" type:"Repeated"`
	MaxAge           *int64    `json:"MaxAge,omitempty" xml:"MaxAge,omitempty"`
}

func (s UpdateRulesAttributeRequestRulesRuleActionsCorsConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateRulesAttributeRequestRulesRuleActionsCorsConfig) GoString() string {
	return s.String()
}

func (s *UpdateRulesAttributeRequestRulesRuleActionsCorsConfig) GetAllowCredentials() *string {
	return s.AllowCredentials
}

func (s *UpdateRulesAttributeRequestRulesRuleActionsCorsConfig) GetAllowHeaders() []*string {
	return s.AllowHeaders
}

func (s *UpdateRulesAttributeRequestRulesRuleActionsCorsConfig) GetAllowMethods() []*string {
	return s.AllowMethods
}

func (s *UpdateRulesAttributeRequestRulesRuleActionsCorsConfig) GetAllowOrigin() []*string {
	return s.AllowOrigin
}

func (s *UpdateRulesAttributeRequestRulesRuleActionsCorsConfig) GetExposeHeaders() []*string {
	return s.ExposeHeaders
}

func (s *UpdateRulesAttributeRequestRulesRuleActionsCorsConfig) GetMaxAge() *int64 {
	return s.MaxAge
}

func (s *UpdateRulesAttributeRequestRulesRuleActionsCorsConfig) SetAllowCredentials(v string) *UpdateRulesAttributeRequestRulesRuleActionsCorsConfig {
	s.AllowCredentials = &v
	return s
}

func (s *UpdateRulesAttributeRequestRulesRuleActionsCorsConfig) SetAllowHeaders(v []*string) *UpdateRulesAttributeRequestRulesRuleActionsCorsConfig {
	s.AllowHeaders = v
	return s
}

func (s *UpdateRulesAttributeRequestRulesRuleActionsCorsConfig) SetAllowMethods(v []*string) *UpdateRulesAttributeRequestRulesRuleActionsCorsConfig {
	s.AllowMethods = v
	return s
}

func (s *UpdateRulesAttributeRequestRulesRuleActionsCorsConfig) SetAllowOrigin(v []*string) *UpdateRulesAttributeRequestRulesRuleActionsCorsConfig {
	s.AllowOrigin = v
	return s
}

func (s *UpdateRulesAttributeRequestRulesRuleActionsCorsConfig) SetExposeHeaders(v []*string) *UpdateRulesAttributeRequestRulesRuleActionsCorsConfig {
	s.ExposeHeaders = v
	return s
}

func (s *UpdateRulesAttributeRequestRulesRuleActionsCorsConfig) SetMaxAge(v int64) *UpdateRulesAttributeRequestRulesRuleActionsCorsConfig {
	s.MaxAge = &v
	return s
}

func (s *UpdateRulesAttributeRequestRulesRuleActionsCorsConfig) Validate() error {
	return dara.Validate(s)
}

type UpdateRulesAttributeRequestRulesRuleActionsFixedResponseConfig struct {
	Content     *string `json:"Content,omitempty" xml:"Content,omitempty"`
	ContentType *string `json:"ContentType,omitempty" xml:"ContentType,omitempty"`
	HttpCode    *string `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
}

func (s UpdateRulesAttributeRequestRulesRuleActionsFixedResponseConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateRulesAttributeRequestRulesRuleActionsFixedResponseConfig) GoString() string {
	return s.String()
}

func (s *UpdateRulesAttributeRequestRulesRuleActionsFixedResponseConfig) GetContent() *string {
	return s.Content
}

func (s *UpdateRulesAttributeRequestRulesRuleActionsFixedResponseConfig) GetContentType() *string {
	return s.ContentType
}

func (s *UpdateRulesAttributeRequestRulesRuleActionsFixedResponseConfig) GetHttpCode() *string {
	return s.HttpCode
}

func (s *UpdateRulesAttributeRequestRulesRuleActionsFixedResponseConfig) SetContent(v string) *UpdateRulesAttributeRequestRulesRuleActionsFixedResponseConfig {
	s.Content = &v
	return s
}

func (s *UpdateRulesAttributeRequestRulesRuleActionsFixedResponseConfig) SetContentType(v string) *UpdateRulesAttributeRequestRulesRuleActionsFixedResponseConfig {
	s.ContentType = &v
	return s
}

func (s *UpdateRulesAttributeRequestRulesRuleActionsFixedResponseConfig) SetHttpCode(v string) *UpdateRulesAttributeRequestRulesRuleActionsFixedResponseConfig {
	s.HttpCode = &v
	return s
}

func (s *UpdateRulesAttributeRequestRulesRuleActionsFixedResponseConfig) Validate() error {
	return dara.Validate(s)
}

type UpdateRulesAttributeRequestRulesRuleActionsForwardGroupConfig struct {
	ServerGroupStickySession *UpdateRulesAttributeRequestRulesRuleActionsForwardGroupConfigServerGroupStickySession `json:"ServerGroupStickySession,omitempty" xml:"ServerGroupStickySession,omitempty" type:"Struct"`
	ServerGroupTuples        []*UpdateRulesAttributeRequestRulesRuleActionsForwardGroupConfigServerGroupTuples      `json:"ServerGroupTuples,omitempty" xml:"ServerGroupTuples,omitempty" type:"Repeated"`
}

func (s UpdateRulesAttributeRequestRulesRuleActionsForwardGroupConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateRulesAttributeRequestRulesRuleActionsForwardGroupConfig) GoString() string {
	return s.String()
}

func (s *UpdateRulesAttributeRequestRulesRuleActionsForwardGroupConfig) GetServerGroupStickySession() *UpdateRulesAttributeRequestRulesRuleActionsForwardGroupConfigServerGroupStickySession {
	return s.ServerGroupStickySession
}

func (s *UpdateRulesAttributeRequestRulesRuleActionsForwardGroupConfig) GetServerGroupTuples() []*UpdateRulesAttributeRequestRulesRuleActionsForwardGroupConfigServerGroupTuples {
	return s.ServerGroupTuples
}

func (s *UpdateRulesAttributeRequestRulesRuleActionsForwardGroupConfig) SetServerGroupStickySession(v *UpdateRulesAttributeRequestRulesRuleActionsForwardGroupConfigServerGroupStickySession) *UpdateRulesAttributeRequestRulesRuleActionsForwardGroupConfig {
	s.ServerGroupStickySession = v
	return s
}

func (s *UpdateRulesAttributeRequestRulesRuleActionsForwardGroupConfig) SetServerGroupTuples(v []*UpdateRulesAttributeRequestRulesRuleActionsForwardGroupConfigServerGroupTuples) *UpdateRulesAttributeRequestRulesRuleActionsForwardGroupConfig {
	s.ServerGroupTuples = v
	return s
}

func (s *UpdateRulesAttributeRequestRulesRuleActionsForwardGroupConfig) Validate() error {
	return dara.Validate(s)
}

type UpdateRulesAttributeRequestRulesRuleActionsForwardGroupConfigServerGroupStickySession struct {
	Enabled *bool  `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	Timeout *int32 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
}

func (s UpdateRulesAttributeRequestRulesRuleActionsForwardGroupConfigServerGroupStickySession) String() string {
	return dara.Prettify(s)
}

func (s UpdateRulesAttributeRequestRulesRuleActionsForwardGroupConfigServerGroupStickySession) GoString() string {
	return s.String()
}

func (s *UpdateRulesAttributeRequestRulesRuleActionsForwardGroupConfigServerGroupStickySession) GetEnabled() *bool {
	return s.Enabled
}

func (s *UpdateRulesAttributeRequestRulesRuleActionsForwardGroupConfigServerGroupStickySession) GetTimeout() *int32 {
	return s.Timeout
}

func (s *UpdateRulesAttributeRequestRulesRuleActionsForwardGroupConfigServerGroupStickySession) SetEnabled(v bool) *UpdateRulesAttributeRequestRulesRuleActionsForwardGroupConfigServerGroupStickySession {
	s.Enabled = &v
	return s
}

func (s *UpdateRulesAttributeRequestRulesRuleActionsForwardGroupConfigServerGroupStickySession) SetTimeout(v int32) *UpdateRulesAttributeRequestRulesRuleActionsForwardGroupConfigServerGroupStickySession {
	s.Timeout = &v
	return s
}

func (s *UpdateRulesAttributeRequestRulesRuleActionsForwardGroupConfigServerGroupStickySession) Validate() error {
	return dara.Validate(s)
}

type UpdateRulesAttributeRequestRulesRuleActionsForwardGroupConfigServerGroupTuples struct {
	ServerGroupId *string `json:"ServerGroupId,omitempty" xml:"ServerGroupId,omitempty"`
	Weight        *int32  `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s UpdateRulesAttributeRequestRulesRuleActionsForwardGroupConfigServerGroupTuples) String() string {
	return dara.Prettify(s)
}

func (s UpdateRulesAttributeRequestRulesRuleActionsForwardGroupConfigServerGroupTuples) GoString() string {
	return s.String()
}

func (s *UpdateRulesAttributeRequestRulesRuleActionsForwardGroupConfigServerGroupTuples) GetServerGroupId() *string {
	return s.ServerGroupId
}

func (s *UpdateRulesAttributeRequestRulesRuleActionsForwardGroupConfigServerGroupTuples) GetWeight() *int32 {
	return s.Weight
}

func (s *UpdateRulesAttributeRequestRulesRuleActionsForwardGroupConfigServerGroupTuples) SetServerGroupId(v string) *UpdateRulesAttributeRequestRulesRuleActionsForwardGroupConfigServerGroupTuples {
	s.ServerGroupId = &v
	return s
}

func (s *UpdateRulesAttributeRequestRulesRuleActionsForwardGroupConfigServerGroupTuples) SetWeight(v int32) *UpdateRulesAttributeRequestRulesRuleActionsForwardGroupConfigServerGroupTuples {
	s.Weight = &v
	return s
}

func (s *UpdateRulesAttributeRequestRulesRuleActionsForwardGroupConfigServerGroupTuples) Validate() error {
	return dara.Validate(s)
}

type UpdateRulesAttributeRequestRulesRuleActionsInsertHeaderConfig struct {
	CoverEnabled *bool   `json:"CoverEnabled,omitempty" xml:"CoverEnabled,omitempty"`
	Key          *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value        *string `json:"Value,omitempty" xml:"Value,omitempty"`
	ValueType    *string `json:"ValueType,omitempty" xml:"ValueType,omitempty"`
}

func (s UpdateRulesAttributeRequestRulesRuleActionsInsertHeaderConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateRulesAttributeRequestRulesRuleActionsInsertHeaderConfig) GoString() string {
	return s.String()
}

func (s *UpdateRulesAttributeRequestRulesRuleActionsInsertHeaderConfig) GetCoverEnabled() *bool {
	return s.CoverEnabled
}

func (s *UpdateRulesAttributeRequestRulesRuleActionsInsertHeaderConfig) GetKey() *string {
	return s.Key
}

func (s *UpdateRulesAttributeRequestRulesRuleActionsInsertHeaderConfig) GetValue() *string {
	return s.Value
}

func (s *UpdateRulesAttributeRequestRulesRuleActionsInsertHeaderConfig) GetValueType() *string {
	return s.ValueType
}

func (s *UpdateRulesAttributeRequestRulesRuleActionsInsertHeaderConfig) SetCoverEnabled(v bool) *UpdateRulesAttributeRequestRulesRuleActionsInsertHeaderConfig {
	s.CoverEnabled = &v
	return s
}

func (s *UpdateRulesAttributeRequestRulesRuleActionsInsertHeaderConfig) SetKey(v string) *UpdateRulesAttributeRequestRulesRuleActionsInsertHeaderConfig {
	s.Key = &v
	return s
}

func (s *UpdateRulesAttributeRequestRulesRuleActionsInsertHeaderConfig) SetValue(v string) *UpdateRulesAttributeRequestRulesRuleActionsInsertHeaderConfig {
	s.Value = &v
	return s
}

func (s *UpdateRulesAttributeRequestRulesRuleActionsInsertHeaderConfig) SetValueType(v string) *UpdateRulesAttributeRequestRulesRuleActionsInsertHeaderConfig {
	s.ValueType = &v
	return s
}

func (s *UpdateRulesAttributeRequestRulesRuleActionsInsertHeaderConfig) Validate() error {
	return dara.Validate(s)
}

type UpdateRulesAttributeRequestRulesRuleActionsRedirectConfig struct {
	Host     *string `json:"Host,omitempty" xml:"Host,omitempty"`
	HttpCode *string `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	Path     *string `json:"Path,omitempty" xml:"Path,omitempty"`
	Port     *string `json:"Port,omitempty" xml:"Port,omitempty"`
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	Query    *string `json:"Query,omitempty" xml:"Query,omitempty"`
}

func (s UpdateRulesAttributeRequestRulesRuleActionsRedirectConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateRulesAttributeRequestRulesRuleActionsRedirectConfig) GoString() string {
	return s.String()
}

func (s *UpdateRulesAttributeRequestRulesRuleActionsRedirectConfig) GetHost() *string {
	return s.Host
}

func (s *UpdateRulesAttributeRequestRulesRuleActionsRedirectConfig) GetHttpCode() *string {
	return s.HttpCode
}

func (s *UpdateRulesAttributeRequestRulesRuleActionsRedirectConfig) GetPath() *string {
	return s.Path
}

func (s *UpdateRulesAttributeRequestRulesRuleActionsRedirectConfig) GetPort() *string {
	return s.Port
}

func (s *UpdateRulesAttributeRequestRulesRuleActionsRedirectConfig) GetProtocol() *string {
	return s.Protocol
}

func (s *UpdateRulesAttributeRequestRulesRuleActionsRedirectConfig) GetQuery() *string {
	return s.Query
}

func (s *UpdateRulesAttributeRequestRulesRuleActionsRedirectConfig) SetHost(v string) *UpdateRulesAttributeRequestRulesRuleActionsRedirectConfig {
	s.Host = &v
	return s
}

func (s *UpdateRulesAttributeRequestRulesRuleActionsRedirectConfig) SetHttpCode(v string) *UpdateRulesAttributeRequestRulesRuleActionsRedirectConfig {
	s.HttpCode = &v
	return s
}

func (s *UpdateRulesAttributeRequestRulesRuleActionsRedirectConfig) SetPath(v string) *UpdateRulesAttributeRequestRulesRuleActionsRedirectConfig {
	s.Path = &v
	return s
}

func (s *UpdateRulesAttributeRequestRulesRuleActionsRedirectConfig) SetPort(v string) *UpdateRulesAttributeRequestRulesRuleActionsRedirectConfig {
	s.Port = &v
	return s
}

func (s *UpdateRulesAttributeRequestRulesRuleActionsRedirectConfig) SetProtocol(v string) *UpdateRulesAttributeRequestRulesRuleActionsRedirectConfig {
	s.Protocol = &v
	return s
}

func (s *UpdateRulesAttributeRequestRulesRuleActionsRedirectConfig) SetQuery(v string) *UpdateRulesAttributeRequestRulesRuleActionsRedirectConfig {
	s.Query = &v
	return s
}

func (s *UpdateRulesAttributeRequestRulesRuleActionsRedirectConfig) Validate() error {
	return dara.Validate(s)
}

type UpdateRulesAttributeRequestRulesRuleActionsRemoveHeaderConfig struct {
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
}

func (s UpdateRulesAttributeRequestRulesRuleActionsRemoveHeaderConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateRulesAttributeRequestRulesRuleActionsRemoveHeaderConfig) GoString() string {
	return s.String()
}

func (s *UpdateRulesAttributeRequestRulesRuleActionsRemoveHeaderConfig) GetKey() *string {
	return s.Key
}

func (s *UpdateRulesAttributeRequestRulesRuleActionsRemoveHeaderConfig) SetKey(v string) *UpdateRulesAttributeRequestRulesRuleActionsRemoveHeaderConfig {
	s.Key = &v
	return s
}

func (s *UpdateRulesAttributeRequestRulesRuleActionsRemoveHeaderConfig) Validate() error {
	return dara.Validate(s)
}

type UpdateRulesAttributeRequestRulesRuleActionsRewriteConfig struct {
	Host  *string `json:"Host,omitempty" xml:"Host,omitempty"`
	Path  *string `json:"Path,omitempty" xml:"Path,omitempty"`
	Query *string `json:"Query,omitempty" xml:"Query,omitempty"`
}

func (s UpdateRulesAttributeRequestRulesRuleActionsRewriteConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateRulesAttributeRequestRulesRuleActionsRewriteConfig) GoString() string {
	return s.String()
}

func (s *UpdateRulesAttributeRequestRulesRuleActionsRewriteConfig) GetHost() *string {
	return s.Host
}

func (s *UpdateRulesAttributeRequestRulesRuleActionsRewriteConfig) GetPath() *string {
	return s.Path
}

func (s *UpdateRulesAttributeRequestRulesRuleActionsRewriteConfig) GetQuery() *string {
	return s.Query
}

func (s *UpdateRulesAttributeRequestRulesRuleActionsRewriteConfig) SetHost(v string) *UpdateRulesAttributeRequestRulesRuleActionsRewriteConfig {
	s.Host = &v
	return s
}

func (s *UpdateRulesAttributeRequestRulesRuleActionsRewriteConfig) SetPath(v string) *UpdateRulesAttributeRequestRulesRuleActionsRewriteConfig {
	s.Path = &v
	return s
}

func (s *UpdateRulesAttributeRequestRulesRuleActionsRewriteConfig) SetQuery(v string) *UpdateRulesAttributeRequestRulesRuleActionsRewriteConfig {
	s.Query = &v
	return s
}

func (s *UpdateRulesAttributeRequestRulesRuleActionsRewriteConfig) Validate() error {
	return dara.Validate(s)
}

type UpdateRulesAttributeRequestRulesRuleActionsTrafficLimitConfig struct {
	PerIpQps *int32 `json:"PerIpQps,omitempty" xml:"PerIpQps,omitempty"`
	QPS      *int32 `json:"QPS,omitempty" xml:"QPS,omitempty"`
}

func (s UpdateRulesAttributeRequestRulesRuleActionsTrafficLimitConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateRulesAttributeRequestRulesRuleActionsTrafficLimitConfig) GoString() string {
	return s.String()
}

func (s *UpdateRulesAttributeRequestRulesRuleActionsTrafficLimitConfig) GetPerIpQps() *int32 {
	return s.PerIpQps
}

func (s *UpdateRulesAttributeRequestRulesRuleActionsTrafficLimitConfig) GetQPS() *int32 {
	return s.QPS
}

func (s *UpdateRulesAttributeRequestRulesRuleActionsTrafficLimitConfig) SetPerIpQps(v int32) *UpdateRulesAttributeRequestRulesRuleActionsTrafficLimitConfig {
	s.PerIpQps = &v
	return s
}

func (s *UpdateRulesAttributeRequestRulesRuleActionsTrafficLimitConfig) SetQPS(v int32) *UpdateRulesAttributeRequestRulesRuleActionsTrafficLimitConfig {
	s.QPS = &v
	return s
}

func (s *UpdateRulesAttributeRequestRulesRuleActionsTrafficLimitConfig) Validate() error {
	return dara.Validate(s)
}

type UpdateRulesAttributeRequestRulesRuleActionsTrafficMirrorConfig struct {
	MirrorGroupConfig *UpdateRulesAttributeRequestRulesRuleActionsTrafficMirrorConfigMirrorGroupConfig `json:"MirrorGroupConfig,omitempty" xml:"MirrorGroupConfig,omitempty" type:"Struct"`
	TargetType        *string                                                                          `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
}

func (s UpdateRulesAttributeRequestRulesRuleActionsTrafficMirrorConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateRulesAttributeRequestRulesRuleActionsTrafficMirrorConfig) GoString() string {
	return s.String()
}

func (s *UpdateRulesAttributeRequestRulesRuleActionsTrafficMirrorConfig) GetMirrorGroupConfig() *UpdateRulesAttributeRequestRulesRuleActionsTrafficMirrorConfigMirrorGroupConfig {
	return s.MirrorGroupConfig
}

func (s *UpdateRulesAttributeRequestRulesRuleActionsTrafficMirrorConfig) GetTargetType() *string {
	return s.TargetType
}

func (s *UpdateRulesAttributeRequestRulesRuleActionsTrafficMirrorConfig) SetMirrorGroupConfig(v *UpdateRulesAttributeRequestRulesRuleActionsTrafficMirrorConfigMirrorGroupConfig) *UpdateRulesAttributeRequestRulesRuleActionsTrafficMirrorConfig {
	s.MirrorGroupConfig = v
	return s
}

func (s *UpdateRulesAttributeRequestRulesRuleActionsTrafficMirrorConfig) SetTargetType(v string) *UpdateRulesAttributeRequestRulesRuleActionsTrafficMirrorConfig {
	s.TargetType = &v
	return s
}

func (s *UpdateRulesAttributeRequestRulesRuleActionsTrafficMirrorConfig) Validate() error {
	return dara.Validate(s)
}

type UpdateRulesAttributeRequestRulesRuleActionsTrafficMirrorConfigMirrorGroupConfig struct {
	ServerGroupTuples []*UpdateRulesAttributeRequestRulesRuleActionsTrafficMirrorConfigMirrorGroupConfigServerGroupTuples `json:"ServerGroupTuples,omitempty" xml:"ServerGroupTuples,omitempty" type:"Repeated"`
}

func (s UpdateRulesAttributeRequestRulesRuleActionsTrafficMirrorConfigMirrorGroupConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateRulesAttributeRequestRulesRuleActionsTrafficMirrorConfigMirrorGroupConfig) GoString() string {
	return s.String()
}

func (s *UpdateRulesAttributeRequestRulesRuleActionsTrafficMirrorConfigMirrorGroupConfig) GetServerGroupTuples() []*UpdateRulesAttributeRequestRulesRuleActionsTrafficMirrorConfigMirrorGroupConfigServerGroupTuples {
	return s.ServerGroupTuples
}

func (s *UpdateRulesAttributeRequestRulesRuleActionsTrafficMirrorConfigMirrorGroupConfig) SetServerGroupTuples(v []*UpdateRulesAttributeRequestRulesRuleActionsTrafficMirrorConfigMirrorGroupConfigServerGroupTuples) *UpdateRulesAttributeRequestRulesRuleActionsTrafficMirrorConfigMirrorGroupConfig {
	s.ServerGroupTuples = v
	return s
}

func (s *UpdateRulesAttributeRequestRulesRuleActionsTrafficMirrorConfigMirrorGroupConfig) Validate() error {
	return dara.Validate(s)
}

type UpdateRulesAttributeRequestRulesRuleActionsTrafficMirrorConfigMirrorGroupConfigServerGroupTuples struct {
	ServerGroupId *string `json:"ServerGroupId,omitempty" xml:"ServerGroupId,omitempty"`
}

func (s UpdateRulesAttributeRequestRulesRuleActionsTrafficMirrorConfigMirrorGroupConfigServerGroupTuples) String() string {
	return dara.Prettify(s)
}

func (s UpdateRulesAttributeRequestRulesRuleActionsTrafficMirrorConfigMirrorGroupConfigServerGroupTuples) GoString() string {
	return s.String()
}

func (s *UpdateRulesAttributeRequestRulesRuleActionsTrafficMirrorConfigMirrorGroupConfigServerGroupTuples) GetServerGroupId() *string {
	return s.ServerGroupId
}

func (s *UpdateRulesAttributeRequestRulesRuleActionsTrafficMirrorConfigMirrorGroupConfigServerGroupTuples) SetServerGroupId(v string) *UpdateRulesAttributeRequestRulesRuleActionsTrafficMirrorConfigMirrorGroupConfigServerGroupTuples {
	s.ServerGroupId = &v
	return s
}

func (s *UpdateRulesAttributeRequestRulesRuleActionsTrafficMirrorConfigMirrorGroupConfigServerGroupTuples) Validate() error {
	return dara.Validate(s)
}

type UpdateRulesAttributeRequestRulesRuleConditions struct {
	CookieConfig             *UpdateRulesAttributeRequestRulesRuleConditionsCookieConfig             `json:"CookieConfig,omitempty" xml:"CookieConfig,omitempty" type:"Struct"`
	HeaderConfig             *UpdateRulesAttributeRequestRulesRuleConditionsHeaderConfig             `json:"HeaderConfig,omitempty" xml:"HeaderConfig,omitempty" type:"Struct"`
	HostConfig               *UpdateRulesAttributeRequestRulesRuleConditionsHostConfig               `json:"HostConfig,omitempty" xml:"HostConfig,omitempty" type:"Struct"`
	MethodConfig             *UpdateRulesAttributeRequestRulesRuleConditionsMethodConfig             `json:"MethodConfig,omitempty" xml:"MethodConfig,omitempty" type:"Struct"`
	PathConfig               *UpdateRulesAttributeRequestRulesRuleConditionsPathConfig               `json:"PathConfig,omitempty" xml:"PathConfig,omitempty" type:"Struct"`
	QueryStringConfig        *UpdateRulesAttributeRequestRulesRuleConditionsQueryStringConfig        `json:"QueryStringConfig,omitempty" xml:"QueryStringConfig,omitempty" type:"Struct"`
	ResponseHeaderConfig     *UpdateRulesAttributeRequestRulesRuleConditionsResponseHeaderConfig     `json:"ResponseHeaderConfig,omitempty" xml:"ResponseHeaderConfig,omitempty" type:"Struct"`
	ResponseStatusCodeConfig *UpdateRulesAttributeRequestRulesRuleConditionsResponseStatusCodeConfig `json:"ResponseStatusCodeConfig,omitempty" xml:"ResponseStatusCodeConfig,omitempty" type:"Struct"`
	SourceIpConfig           *UpdateRulesAttributeRequestRulesRuleConditionsSourceIpConfig           `json:"SourceIpConfig,omitempty" xml:"SourceIpConfig,omitempty" type:"Struct"`
	Type                     *string                                                                 `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s UpdateRulesAttributeRequestRulesRuleConditions) String() string {
	return dara.Prettify(s)
}

func (s UpdateRulesAttributeRequestRulesRuleConditions) GoString() string {
	return s.String()
}

func (s *UpdateRulesAttributeRequestRulesRuleConditions) GetCookieConfig() *UpdateRulesAttributeRequestRulesRuleConditionsCookieConfig {
	return s.CookieConfig
}

func (s *UpdateRulesAttributeRequestRulesRuleConditions) GetHeaderConfig() *UpdateRulesAttributeRequestRulesRuleConditionsHeaderConfig {
	return s.HeaderConfig
}

func (s *UpdateRulesAttributeRequestRulesRuleConditions) GetHostConfig() *UpdateRulesAttributeRequestRulesRuleConditionsHostConfig {
	return s.HostConfig
}

func (s *UpdateRulesAttributeRequestRulesRuleConditions) GetMethodConfig() *UpdateRulesAttributeRequestRulesRuleConditionsMethodConfig {
	return s.MethodConfig
}

func (s *UpdateRulesAttributeRequestRulesRuleConditions) GetPathConfig() *UpdateRulesAttributeRequestRulesRuleConditionsPathConfig {
	return s.PathConfig
}

func (s *UpdateRulesAttributeRequestRulesRuleConditions) GetQueryStringConfig() *UpdateRulesAttributeRequestRulesRuleConditionsQueryStringConfig {
	return s.QueryStringConfig
}

func (s *UpdateRulesAttributeRequestRulesRuleConditions) GetResponseHeaderConfig() *UpdateRulesAttributeRequestRulesRuleConditionsResponseHeaderConfig {
	return s.ResponseHeaderConfig
}

func (s *UpdateRulesAttributeRequestRulesRuleConditions) GetResponseStatusCodeConfig() *UpdateRulesAttributeRequestRulesRuleConditionsResponseStatusCodeConfig {
	return s.ResponseStatusCodeConfig
}

func (s *UpdateRulesAttributeRequestRulesRuleConditions) GetSourceIpConfig() *UpdateRulesAttributeRequestRulesRuleConditionsSourceIpConfig {
	return s.SourceIpConfig
}

func (s *UpdateRulesAttributeRequestRulesRuleConditions) GetType() *string {
	return s.Type
}

func (s *UpdateRulesAttributeRequestRulesRuleConditions) SetCookieConfig(v *UpdateRulesAttributeRequestRulesRuleConditionsCookieConfig) *UpdateRulesAttributeRequestRulesRuleConditions {
	s.CookieConfig = v
	return s
}

func (s *UpdateRulesAttributeRequestRulesRuleConditions) SetHeaderConfig(v *UpdateRulesAttributeRequestRulesRuleConditionsHeaderConfig) *UpdateRulesAttributeRequestRulesRuleConditions {
	s.HeaderConfig = v
	return s
}

func (s *UpdateRulesAttributeRequestRulesRuleConditions) SetHostConfig(v *UpdateRulesAttributeRequestRulesRuleConditionsHostConfig) *UpdateRulesAttributeRequestRulesRuleConditions {
	s.HostConfig = v
	return s
}

func (s *UpdateRulesAttributeRequestRulesRuleConditions) SetMethodConfig(v *UpdateRulesAttributeRequestRulesRuleConditionsMethodConfig) *UpdateRulesAttributeRequestRulesRuleConditions {
	s.MethodConfig = v
	return s
}

func (s *UpdateRulesAttributeRequestRulesRuleConditions) SetPathConfig(v *UpdateRulesAttributeRequestRulesRuleConditionsPathConfig) *UpdateRulesAttributeRequestRulesRuleConditions {
	s.PathConfig = v
	return s
}

func (s *UpdateRulesAttributeRequestRulesRuleConditions) SetQueryStringConfig(v *UpdateRulesAttributeRequestRulesRuleConditionsQueryStringConfig) *UpdateRulesAttributeRequestRulesRuleConditions {
	s.QueryStringConfig = v
	return s
}

func (s *UpdateRulesAttributeRequestRulesRuleConditions) SetResponseHeaderConfig(v *UpdateRulesAttributeRequestRulesRuleConditionsResponseHeaderConfig) *UpdateRulesAttributeRequestRulesRuleConditions {
	s.ResponseHeaderConfig = v
	return s
}

func (s *UpdateRulesAttributeRequestRulesRuleConditions) SetResponseStatusCodeConfig(v *UpdateRulesAttributeRequestRulesRuleConditionsResponseStatusCodeConfig) *UpdateRulesAttributeRequestRulesRuleConditions {
	s.ResponseStatusCodeConfig = v
	return s
}

func (s *UpdateRulesAttributeRequestRulesRuleConditions) SetSourceIpConfig(v *UpdateRulesAttributeRequestRulesRuleConditionsSourceIpConfig) *UpdateRulesAttributeRequestRulesRuleConditions {
	s.SourceIpConfig = v
	return s
}

func (s *UpdateRulesAttributeRequestRulesRuleConditions) SetType(v string) *UpdateRulesAttributeRequestRulesRuleConditions {
	s.Type = &v
	return s
}

func (s *UpdateRulesAttributeRequestRulesRuleConditions) Validate() error {
	return dara.Validate(s)
}

type UpdateRulesAttributeRequestRulesRuleConditionsCookieConfig struct {
	Values []*UpdateRulesAttributeRequestRulesRuleConditionsCookieConfigValues `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s UpdateRulesAttributeRequestRulesRuleConditionsCookieConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateRulesAttributeRequestRulesRuleConditionsCookieConfig) GoString() string {
	return s.String()
}

func (s *UpdateRulesAttributeRequestRulesRuleConditionsCookieConfig) GetValues() []*UpdateRulesAttributeRequestRulesRuleConditionsCookieConfigValues {
	return s.Values
}

func (s *UpdateRulesAttributeRequestRulesRuleConditionsCookieConfig) SetValues(v []*UpdateRulesAttributeRequestRulesRuleConditionsCookieConfigValues) *UpdateRulesAttributeRequestRulesRuleConditionsCookieConfig {
	s.Values = v
	return s
}

func (s *UpdateRulesAttributeRequestRulesRuleConditionsCookieConfig) Validate() error {
	return dara.Validate(s)
}

type UpdateRulesAttributeRequestRulesRuleConditionsCookieConfigValues struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateRulesAttributeRequestRulesRuleConditionsCookieConfigValues) String() string {
	return dara.Prettify(s)
}

func (s UpdateRulesAttributeRequestRulesRuleConditionsCookieConfigValues) GoString() string {
	return s.String()
}

func (s *UpdateRulesAttributeRequestRulesRuleConditionsCookieConfigValues) GetKey() *string {
	return s.Key
}

func (s *UpdateRulesAttributeRequestRulesRuleConditionsCookieConfigValues) GetValue() *string {
	return s.Value
}

func (s *UpdateRulesAttributeRequestRulesRuleConditionsCookieConfigValues) SetKey(v string) *UpdateRulesAttributeRequestRulesRuleConditionsCookieConfigValues {
	s.Key = &v
	return s
}

func (s *UpdateRulesAttributeRequestRulesRuleConditionsCookieConfigValues) SetValue(v string) *UpdateRulesAttributeRequestRulesRuleConditionsCookieConfigValues {
	s.Value = &v
	return s
}

func (s *UpdateRulesAttributeRequestRulesRuleConditionsCookieConfigValues) Validate() error {
	return dara.Validate(s)
}

type UpdateRulesAttributeRequestRulesRuleConditionsHeaderConfig struct {
	Key    *string   `json:"Key,omitempty" xml:"Key,omitempty"`
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s UpdateRulesAttributeRequestRulesRuleConditionsHeaderConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateRulesAttributeRequestRulesRuleConditionsHeaderConfig) GoString() string {
	return s.String()
}

func (s *UpdateRulesAttributeRequestRulesRuleConditionsHeaderConfig) GetKey() *string {
	return s.Key
}

func (s *UpdateRulesAttributeRequestRulesRuleConditionsHeaderConfig) GetValues() []*string {
	return s.Values
}

func (s *UpdateRulesAttributeRequestRulesRuleConditionsHeaderConfig) SetKey(v string) *UpdateRulesAttributeRequestRulesRuleConditionsHeaderConfig {
	s.Key = &v
	return s
}

func (s *UpdateRulesAttributeRequestRulesRuleConditionsHeaderConfig) SetValues(v []*string) *UpdateRulesAttributeRequestRulesRuleConditionsHeaderConfig {
	s.Values = v
	return s
}

func (s *UpdateRulesAttributeRequestRulesRuleConditionsHeaderConfig) Validate() error {
	return dara.Validate(s)
}

type UpdateRulesAttributeRequestRulesRuleConditionsHostConfig struct {
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s UpdateRulesAttributeRequestRulesRuleConditionsHostConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateRulesAttributeRequestRulesRuleConditionsHostConfig) GoString() string {
	return s.String()
}

func (s *UpdateRulesAttributeRequestRulesRuleConditionsHostConfig) GetValues() []*string {
	return s.Values
}

func (s *UpdateRulesAttributeRequestRulesRuleConditionsHostConfig) SetValues(v []*string) *UpdateRulesAttributeRequestRulesRuleConditionsHostConfig {
	s.Values = v
	return s
}

func (s *UpdateRulesAttributeRequestRulesRuleConditionsHostConfig) Validate() error {
	return dara.Validate(s)
}

type UpdateRulesAttributeRequestRulesRuleConditionsMethodConfig struct {
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s UpdateRulesAttributeRequestRulesRuleConditionsMethodConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateRulesAttributeRequestRulesRuleConditionsMethodConfig) GoString() string {
	return s.String()
}

func (s *UpdateRulesAttributeRequestRulesRuleConditionsMethodConfig) GetValues() []*string {
	return s.Values
}

func (s *UpdateRulesAttributeRequestRulesRuleConditionsMethodConfig) SetValues(v []*string) *UpdateRulesAttributeRequestRulesRuleConditionsMethodConfig {
	s.Values = v
	return s
}

func (s *UpdateRulesAttributeRequestRulesRuleConditionsMethodConfig) Validate() error {
	return dara.Validate(s)
}

type UpdateRulesAttributeRequestRulesRuleConditionsPathConfig struct {
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s UpdateRulesAttributeRequestRulesRuleConditionsPathConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateRulesAttributeRequestRulesRuleConditionsPathConfig) GoString() string {
	return s.String()
}

func (s *UpdateRulesAttributeRequestRulesRuleConditionsPathConfig) GetValues() []*string {
	return s.Values
}

func (s *UpdateRulesAttributeRequestRulesRuleConditionsPathConfig) SetValues(v []*string) *UpdateRulesAttributeRequestRulesRuleConditionsPathConfig {
	s.Values = v
	return s
}

func (s *UpdateRulesAttributeRequestRulesRuleConditionsPathConfig) Validate() error {
	return dara.Validate(s)
}

type UpdateRulesAttributeRequestRulesRuleConditionsQueryStringConfig struct {
	Values []*UpdateRulesAttributeRequestRulesRuleConditionsQueryStringConfigValues `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s UpdateRulesAttributeRequestRulesRuleConditionsQueryStringConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateRulesAttributeRequestRulesRuleConditionsQueryStringConfig) GoString() string {
	return s.String()
}

func (s *UpdateRulesAttributeRequestRulesRuleConditionsQueryStringConfig) GetValues() []*UpdateRulesAttributeRequestRulesRuleConditionsQueryStringConfigValues {
	return s.Values
}

func (s *UpdateRulesAttributeRequestRulesRuleConditionsQueryStringConfig) SetValues(v []*UpdateRulesAttributeRequestRulesRuleConditionsQueryStringConfigValues) *UpdateRulesAttributeRequestRulesRuleConditionsQueryStringConfig {
	s.Values = v
	return s
}

func (s *UpdateRulesAttributeRequestRulesRuleConditionsQueryStringConfig) Validate() error {
	return dara.Validate(s)
}

type UpdateRulesAttributeRequestRulesRuleConditionsQueryStringConfigValues struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateRulesAttributeRequestRulesRuleConditionsQueryStringConfigValues) String() string {
	return dara.Prettify(s)
}

func (s UpdateRulesAttributeRequestRulesRuleConditionsQueryStringConfigValues) GoString() string {
	return s.String()
}

func (s *UpdateRulesAttributeRequestRulesRuleConditionsQueryStringConfigValues) GetKey() *string {
	return s.Key
}

func (s *UpdateRulesAttributeRequestRulesRuleConditionsQueryStringConfigValues) GetValue() *string {
	return s.Value
}

func (s *UpdateRulesAttributeRequestRulesRuleConditionsQueryStringConfigValues) SetKey(v string) *UpdateRulesAttributeRequestRulesRuleConditionsQueryStringConfigValues {
	s.Key = &v
	return s
}

func (s *UpdateRulesAttributeRequestRulesRuleConditionsQueryStringConfigValues) SetValue(v string) *UpdateRulesAttributeRequestRulesRuleConditionsQueryStringConfigValues {
	s.Value = &v
	return s
}

func (s *UpdateRulesAttributeRequestRulesRuleConditionsQueryStringConfigValues) Validate() error {
	return dara.Validate(s)
}

type UpdateRulesAttributeRequestRulesRuleConditionsResponseHeaderConfig struct {
	Key    *string   `json:"Key,omitempty" xml:"Key,omitempty"`
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s UpdateRulesAttributeRequestRulesRuleConditionsResponseHeaderConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateRulesAttributeRequestRulesRuleConditionsResponseHeaderConfig) GoString() string {
	return s.String()
}

func (s *UpdateRulesAttributeRequestRulesRuleConditionsResponseHeaderConfig) GetKey() *string {
	return s.Key
}

func (s *UpdateRulesAttributeRequestRulesRuleConditionsResponseHeaderConfig) GetValues() []*string {
	return s.Values
}

func (s *UpdateRulesAttributeRequestRulesRuleConditionsResponseHeaderConfig) SetKey(v string) *UpdateRulesAttributeRequestRulesRuleConditionsResponseHeaderConfig {
	s.Key = &v
	return s
}

func (s *UpdateRulesAttributeRequestRulesRuleConditionsResponseHeaderConfig) SetValues(v []*string) *UpdateRulesAttributeRequestRulesRuleConditionsResponseHeaderConfig {
	s.Values = v
	return s
}

func (s *UpdateRulesAttributeRequestRulesRuleConditionsResponseHeaderConfig) Validate() error {
	return dara.Validate(s)
}

type UpdateRulesAttributeRequestRulesRuleConditionsResponseStatusCodeConfig struct {
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s UpdateRulesAttributeRequestRulesRuleConditionsResponseStatusCodeConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateRulesAttributeRequestRulesRuleConditionsResponseStatusCodeConfig) GoString() string {
	return s.String()
}

func (s *UpdateRulesAttributeRequestRulesRuleConditionsResponseStatusCodeConfig) GetValues() []*string {
	return s.Values
}

func (s *UpdateRulesAttributeRequestRulesRuleConditionsResponseStatusCodeConfig) SetValues(v []*string) *UpdateRulesAttributeRequestRulesRuleConditionsResponseStatusCodeConfig {
	s.Values = v
	return s
}

func (s *UpdateRulesAttributeRequestRulesRuleConditionsResponseStatusCodeConfig) Validate() error {
	return dara.Validate(s)
}

type UpdateRulesAttributeRequestRulesRuleConditionsSourceIpConfig struct {
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s UpdateRulesAttributeRequestRulesRuleConditionsSourceIpConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateRulesAttributeRequestRulesRuleConditionsSourceIpConfig) GoString() string {
	return s.String()
}

func (s *UpdateRulesAttributeRequestRulesRuleConditionsSourceIpConfig) GetValues() []*string {
	return s.Values
}

func (s *UpdateRulesAttributeRequestRulesRuleConditionsSourceIpConfig) SetValues(v []*string) *UpdateRulesAttributeRequestRulesRuleConditionsSourceIpConfig {
	s.Values = v
	return s
}

func (s *UpdateRulesAttributeRequestRulesRuleConditionsSourceIpConfig) Validate() error {
	return dara.Validate(s)
}
