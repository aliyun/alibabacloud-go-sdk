// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type AddonMeta struct {
	Alias                   *string                  `json:"alias,omitempty" xml:"alias,omitempty"`
	Categories              []*string                `json:"categories,omitempty" xml:"categories,omitempty" type:"Repeated"`
	Dashboards              []*AddonMetaDashboards   `json:"dashboards,omitempty" xml:"dashboards,omitempty" type:"Repeated"`
	Description             *string                  `json:"description,omitempty" xml:"description,omitempty"`
	Environments            []*AddonMetaEnvironments `json:"environments,omitempty" xml:"environments,omitempty" type:"Repeated"`
	Icon                    *string                  `json:"icon,omitempty" xml:"icon,omitempty"`
	Keywords                []*string                `json:"keywords,omitempty" xml:"keywords,omitempty" type:"Repeated"`
	Language                *string                  `json:"language,omitempty" xml:"language,omitempty"`
	LatestReleaseCreateTime *string                  `json:"latestReleaseCreateTime,omitempty" xml:"latestReleaseCreateTime,omitempty"`
	Name                    *string                  `json:"name,omitempty" xml:"name,omitempty"`
	Once                    *bool                    `json:"once,omitempty" xml:"once,omitempty"`
	Scene                   *string                  `json:"scene,omitempty" xml:"scene,omitempty"`
	Version                 *string                  `json:"version,omitempty" xml:"version,omitempty"`
	Weight                  *int32                   `json:"weight,omitempty" xml:"weight,omitempty"`
}

func (s AddonMeta) String() string {
	return tea.Prettify(s)
}

func (s AddonMeta) GoString() string {
	return s.String()
}

func (s *AddonMeta) SetAlias(v string) *AddonMeta {
	s.Alias = &v
	return s
}

func (s *AddonMeta) SetCategories(v []*string) *AddonMeta {
	s.Categories = v
	return s
}

func (s *AddonMeta) SetDashboards(v []*AddonMetaDashboards) *AddonMeta {
	s.Dashboards = v
	return s
}

func (s *AddonMeta) SetDescription(v string) *AddonMeta {
	s.Description = &v
	return s
}

func (s *AddonMeta) SetEnvironments(v []*AddonMetaEnvironments) *AddonMeta {
	s.Environments = v
	return s
}

func (s *AddonMeta) SetIcon(v string) *AddonMeta {
	s.Icon = &v
	return s
}

func (s *AddonMeta) SetKeywords(v []*string) *AddonMeta {
	s.Keywords = v
	return s
}

func (s *AddonMeta) SetLanguage(v string) *AddonMeta {
	s.Language = &v
	return s
}

func (s *AddonMeta) SetLatestReleaseCreateTime(v string) *AddonMeta {
	s.LatestReleaseCreateTime = &v
	return s
}

func (s *AddonMeta) SetName(v string) *AddonMeta {
	s.Name = &v
	return s
}

func (s *AddonMeta) SetOnce(v bool) *AddonMeta {
	s.Once = &v
	return s
}

func (s *AddonMeta) SetScene(v string) *AddonMeta {
	s.Scene = &v
	return s
}

func (s *AddonMeta) SetVersion(v string) *AddonMeta {
	s.Version = &v
	return s
}

func (s *AddonMeta) SetWeight(v int32) *AddonMeta {
	s.Weight = &v
	return s
}

type AddonMetaDashboards struct {
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	Name        *string `json:"name,omitempty" xml:"name,omitempty"`
	Url         *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s AddonMetaDashboards) String() string {
	return tea.Prettify(s)
}

func (s AddonMetaDashboards) GoString() string {
	return s.String()
}

func (s *AddonMetaDashboards) SetDescription(v string) *AddonMetaDashboards {
	s.Description = &v
	return s
}

func (s *AddonMetaDashboards) SetName(v string) *AddonMetaDashboards {
	s.Name = &v
	return s
}

func (s *AddonMetaDashboards) SetUrl(v string) *AddonMetaDashboards {
	s.Url = &v
	return s
}

type AddonMetaEnvironments struct {
	Dependencies *AddonMetaEnvironmentsDependencies `json:"dependencies,omitempty" xml:"dependencies,omitempty" type:"Struct"`
	Description  *string                            `json:"description,omitempty" xml:"description,omitempty"`
	Enable       *bool                              `json:"enable,omitempty" xml:"enable,omitempty"`
	Label        *string                            `json:"label,omitempty" xml:"label,omitempty"`
	Name         *string                            `json:"name,omitempty" xml:"name,omitempty"`
	Policies     *AddonMetaEnvironmentsPolicies     `json:"policies,omitempty" xml:"policies,omitempty" type:"Struct"`
	PolicyType   *string                            `json:"policyType,omitempty" xml:"policyType,omitempty"`
}

func (s AddonMetaEnvironments) String() string {
	return tea.Prettify(s)
}

func (s AddonMetaEnvironments) GoString() string {
	return s.String()
}

func (s *AddonMetaEnvironments) SetDependencies(v *AddonMetaEnvironmentsDependencies) *AddonMetaEnvironments {
	s.Dependencies = v
	return s
}

func (s *AddonMetaEnvironments) SetDescription(v string) *AddonMetaEnvironments {
	s.Description = &v
	return s
}

func (s *AddonMetaEnvironments) SetEnable(v bool) *AddonMetaEnvironments {
	s.Enable = &v
	return s
}

func (s *AddonMetaEnvironments) SetLabel(v string) *AddonMetaEnvironments {
	s.Label = &v
	return s
}

func (s *AddonMetaEnvironments) SetName(v string) *AddonMetaEnvironments {
	s.Name = &v
	return s
}

func (s *AddonMetaEnvironments) SetPolicies(v *AddonMetaEnvironmentsPolicies) *AddonMetaEnvironments {
	s.Policies = v
	return s
}

func (s *AddonMetaEnvironments) SetPolicyType(v string) *AddonMetaEnvironments {
	s.PolicyType = &v
	return s
}

type AddonMetaEnvironmentsDependencies struct {
	ClusterTypes []*string        `json:"clusterTypes,omitempty" xml:"clusterTypes,omitempty" type:"Repeated"`
	Features     map[string]*bool `json:"features,omitempty" xml:"features,omitempty"`
	Services     []*string        `json:"services,omitempty" xml:"services,omitempty" type:"Repeated"`
}

func (s AddonMetaEnvironmentsDependencies) String() string {
	return tea.Prettify(s)
}

func (s AddonMetaEnvironmentsDependencies) GoString() string {
	return s.String()
}

func (s *AddonMetaEnvironmentsDependencies) SetClusterTypes(v []*string) *AddonMetaEnvironmentsDependencies {
	s.ClusterTypes = v
	return s
}

func (s *AddonMetaEnvironmentsDependencies) SetFeatures(v map[string]*bool) *AddonMetaEnvironmentsDependencies {
	s.Features = v
	return s
}

func (s *AddonMetaEnvironmentsDependencies) SetServices(v []*string) *AddonMetaEnvironmentsDependencies {
	s.Services = v
	return s
}

type AddonMetaEnvironmentsPolicies struct {
	AlertDefaultStatus          *string                                       `json:"alertDefaultStatus,omitempty" xml:"alertDefaultStatus,omitempty"`
	BindDefaultPolicy           *bool                                         `json:"bindDefaultPolicy,omitempty" xml:"bindDefaultPolicy,omitempty"`
	BindEntity                  *AddonMetaEnvironmentsPoliciesBindEntity      `json:"bindEntity,omitempty" xml:"bindEntity,omitempty" type:"Struct"`
	DefaultInstall              *bool                                         `json:"defaultInstall,omitempty" xml:"defaultInstall,omitempty"`
	EnableServiceAccount        *bool                                         `json:"enableServiceAccount,omitempty" xml:"enableServiceAccount,omitempty"`
	MetricCheckRule             *AddonMetaEnvironmentsPoliciesMetricCheckRule `json:"metricCheckRule,omitempty" xml:"metricCheckRule,omitempty" type:"Struct"`
	NeedRestartAfterIntegration *bool                                         `json:"needRestartAfterIntegration,omitempty" xml:"needRestartAfterIntegration,omitempty"`
	Protocols                   []*AddonMetaEnvironmentsPoliciesProtocols     `json:"protocols,omitempty" xml:"protocols,omitempty" type:"Repeated"`
	TargetAddonName             *string                                       `json:"targetAddonName,omitempty" xml:"targetAddonName,omitempty"`
}

func (s AddonMetaEnvironmentsPolicies) String() string {
	return tea.Prettify(s)
}

func (s AddonMetaEnvironmentsPolicies) GoString() string {
	return s.String()
}

func (s *AddonMetaEnvironmentsPolicies) SetAlertDefaultStatus(v string) *AddonMetaEnvironmentsPolicies {
	s.AlertDefaultStatus = &v
	return s
}

func (s *AddonMetaEnvironmentsPolicies) SetBindDefaultPolicy(v bool) *AddonMetaEnvironmentsPolicies {
	s.BindDefaultPolicy = &v
	return s
}

func (s *AddonMetaEnvironmentsPolicies) SetBindEntity(v *AddonMetaEnvironmentsPoliciesBindEntity) *AddonMetaEnvironmentsPolicies {
	s.BindEntity = v
	return s
}

func (s *AddonMetaEnvironmentsPolicies) SetDefaultInstall(v bool) *AddonMetaEnvironmentsPolicies {
	s.DefaultInstall = &v
	return s
}

func (s *AddonMetaEnvironmentsPolicies) SetEnableServiceAccount(v bool) *AddonMetaEnvironmentsPolicies {
	s.EnableServiceAccount = &v
	return s
}

func (s *AddonMetaEnvironmentsPolicies) SetMetricCheckRule(v *AddonMetaEnvironmentsPoliciesMetricCheckRule) *AddonMetaEnvironmentsPolicies {
	s.MetricCheckRule = v
	return s
}

func (s *AddonMetaEnvironmentsPolicies) SetNeedRestartAfterIntegration(v bool) *AddonMetaEnvironmentsPolicies {
	s.NeedRestartAfterIntegration = &v
	return s
}

func (s *AddonMetaEnvironmentsPolicies) SetProtocols(v []*AddonMetaEnvironmentsPoliciesProtocols) *AddonMetaEnvironmentsPolicies {
	s.Protocols = v
	return s
}

func (s *AddonMetaEnvironmentsPolicies) SetTargetAddonName(v string) *AddonMetaEnvironmentsPolicies {
	s.TargetAddonName = &v
	return s
}

type AddonMetaEnvironmentsPoliciesBindEntity struct {
	EntityGroupMode  *bool   `json:"entityGroupMode,omitempty" xml:"entityGroupMode,omitempty"`
	EntityType       *string `json:"entityType,omitempty" xml:"entityType,omitempty"`
	SingleEntityMode *bool   `json:"singleEntityMode,omitempty" xml:"singleEntityMode,omitempty"`
	VpcIdFieldKey    *string `json:"vpcIdFieldKey,omitempty" xml:"vpcIdFieldKey,omitempty"`
}

func (s AddonMetaEnvironmentsPoliciesBindEntity) String() string {
	return tea.Prettify(s)
}

func (s AddonMetaEnvironmentsPoliciesBindEntity) GoString() string {
	return s.String()
}

func (s *AddonMetaEnvironmentsPoliciesBindEntity) SetEntityGroupMode(v bool) *AddonMetaEnvironmentsPoliciesBindEntity {
	s.EntityGroupMode = &v
	return s
}

func (s *AddonMetaEnvironmentsPoliciesBindEntity) SetEntityType(v string) *AddonMetaEnvironmentsPoliciesBindEntity {
	s.EntityType = &v
	return s
}

func (s *AddonMetaEnvironmentsPoliciesBindEntity) SetSingleEntityMode(v bool) *AddonMetaEnvironmentsPoliciesBindEntity {
	s.SingleEntityMode = &v
	return s
}

func (s *AddonMetaEnvironmentsPoliciesBindEntity) SetVpcIdFieldKey(v string) *AddonMetaEnvironmentsPoliciesBindEntity {
	s.VpcIdFieldKey = &v
	return s
}

type AddonMetaEnvironmentsPoliciesMetricCheckRule struct {
	PromQL []*string `json:"promQL,omitempty" xml:"promQL,omitempty" type:"Repeated"`
}

func (s AddonMetaEnvironmentsPoliciesMetricCheckRule) String() string {
	return tea.Prettify(s)
}

func (s AddonMetaEnvironmentsPoliciesMetricCheckRule) GoString() string {
	return s.String()
}

func (s *AddonMetaEnvironmentsPoliciesMetricCheckRule) SetPromQL(v []*string) *AddonMetaEnvironmentsPoliciesMetricCheckRule {
	s.PromQL = v
	return s
}

type AddonMetaEnvironmentsPoliciesProtocols struct {
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	Icon        *string `json:"icon,omitempty" xml:"icon,omitempty"`
	Label       *string `json:"label,omitempty" xml:"label,omitempty"`
	Name        *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s AddonMetaEnvironmentsPoliciesProtocols) String() string {
	return tea.Prettify(s)
}

func (s AddonMetaEnvironmentsPoliciesProtocols) GoString() string {
	return s.String()
}

func (s *AddonMetaEnvironmentsPoliciesProtocols) SetDescription(v string) *AddonMetaEnvironmentsPoliciesProtocols {
	s.Description = &v
	return s
}

func (s *AddonMetaEnvironmentsPoliciesProtocols) SetIcon(v string) *AddonMetaEnvironmentsPoliciesProtocols {
	s.Icon = &v
	return s
}

func (s *AddonMetaEnvironmentsPoliciesProtocols) SetLabel(v string) *AddonMetaEnvironmentsPoliciesProtocols {
	s.Label = &v
	return s
}

func (s *AddonMetaEnvironmentsPoliciesProtocols) SetName(v string) *AddonMetaEnvironmentsPoliciesProtocols {
	s.Name = &v
	return s
}

type AlertEventIntegrationPolicyForModify struct {
	// This parameter is required.
	AlertEventIntegrationPolicyName *string            `json:"alertEventIntegrationPolicyName,omitempty" xml:"alertEventIntegrationPolicyName,omitempty"`
	Description                     *string            `json:"description,omitempty" xml:"description,omitempty"`
	FilterSetting                   *FilterSetting     `json:"filterSetting,omitempty" xml:"filterSetting,omitempty"`
	IntegrationSetting              *string            `json:"integrationSetting,omitempty" xml:"integrationSetting,omitempty"`
	TransformerSetting              []*TransformAction `json:"transformerSetting,omitempty" xml:"transformerSetting,omitempty" type:"Repeated"`
	Type                            *string            `json:"type,omitempty" xml:"type,omitempty"`
}

func (s AlertEventIntegrationPolicyForModify) String() string {
	return tea.Prettify(s)
}

func (s AlertEventIntegrationPolicyForModify) GoString() string {
	return s.String()
}

func (s *AlertEventIntegrationPolicyForModify) SetAlertEventIntegrationPolicyName(v string) *AlertEventIntegrationPolicyForModify {
	s.AlertEventIntegrationPolicyName = &v
	return s
}

func (s *AlertEventIntegrationPolicyForModify) SetDescription(v string) *AlertEventIntegrationPolicyForModify {
	s.Description = &v
	return s
}

func (s *AlertEventIntegrationPolicyForModify) SetFilterSetting(v *FilterSetting) *AlertEventIntegrationPolicyForModify {
	s.FilterSetting = v
	return s
}

func (s *AlertEventIntegrationPolicyForModify) SetIntegrationSetting(v string) *AlertEventIntegrationPolicyForModify {
	s.IntegrationSetting = &v
	return s
}

func (s *AlertEventIntegrationPolicyForModify) SetTransformerSetting(v []*TransformAction) *AlertEventIntegrationPolicyForModify {
	s.TransformerSetting = v
	return s
}

func (s *AlertEventIntegrationPolicyForModify) SetType(v string) *AlertEventIntegrationPolicyForModify {
	s.Type = &v
	return s
}

type AlertEventIntegrationPolicyForView struct {
	AlertEventIntegrationPolicyId *string `json:"alertEventIntegrationPolicyId,omitempty" xml:"alertEventIntegrationPolicyId,omitempty"`
	// This parameter is required.
	AlertEventIntegrationPolicyName *string            `json:"alertEventIntegrationPolicyName,omitempty" xml:"alertEventIntegrationPolicyName,omitempty"`
	CreateTime                      *string            `json:"createTime,omitempty" xml:"createTime,omitempty"`
	Description                     *string            `json:"description,omitempty" xml:"description,omitempty"`
	Enable                          *bool              `json:"enable,omitempty" xml:"enable,omitempty"`
	FilterSetting                   *FilterSetting     `json:"filterSetting,omitempty" xml:"filterSetting,omitempty"`
	IntegrationSetting              *string            `json:"integrationSetting,omitempty" xml:"integrationSetting,omitempty"`
	Token                           *string            `json:"token,omitempty" xml:"token,omitempty"`
	TransformerSetting              []*TransformAction `json:"transformerSetting,omitempty" xml:"transformerSetting,omitempty" type:"Repeated"`
	Type                            *string            `json:"type,omitempty" xml:"type,omitempty"`
	UpdateTime                      *string            `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
	UserId                          *string            `json:"userId,omitempty" xml:"userId,omitempty"`
	Workspace                       *string            `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s AlertEventIntegrationPolicyForView) String() string {
	return tea.Prettify(s)
}

func (s AlertEventIntegrationPolicyForView) GoString() string {
	return s.String()
}

func (s *AlertEventIntegrationPolicyForView) SetAlertEventIntegrationPolicyId(v string) *AlertEventIntegrationPolicyForView {
	s.AlertEventIntegrationPolicyId = &v
	return s
}

func (s *AlertEventIntegrationPolicyForView) SetAlertEventIntegrationPolicyName(v string) *AlertEventIntegrationPolicyForView {
	s.AlertEventIntegrationPolicyName = &v
	return s
}

func (s *AlertEventIntegrationPolicyForView) SetCreateTime(v string) *AlertEventIntegrationPolicyForView {
	s.CreateTime = &v
	return s
}

func (s *AlertEventIntegrationPolicyForView) SetDescription(v string) *AlertEventIntegrationPolicyForView {
	s.Description = &v
	return s
}

func (s *AlertEventIntegrationPolicyForView) SetEnable(v bool) *AlertEventIntegrationPolicyForView {
	s.Enable = &v
	return s
}

func (s *AlertEventIntegrationPolicyForView) SetFilterSetting(v *FilterSetting) *AlertEventIntegrationPolicyForView {
	s.FilterSetting = v
	return s
}

func (s *AlertEventIntegrationPolicyForView) SetIntegrationSetting(v string) *AlertEventIntegrationPolicyForView {
	s.IntegrationSetting = &v
	return s
}

func (s *AlertEventIntegrationPolicyForView) SetToken(v string) *AlertEventIntegrationPolicyForView {
	s.Token = &v
	return s
}

func (s *AlertEventIntegrationPolicyForView) SetTransformerSetting(v []*TransformAction) *AlertEventIntegrationPolicyForView {
	s.TransformerSetting = v
	return s
}

func (s *AlertEventIntegrationPolicyForView) SetType(v string) *AlertEventIntegrationPolicyForView {
	s.Type = &v
	return s
}

func (s *AlertEventIntegrationPolicyForView) SetUpdateTime(v string) *AlertEventIntegrationPolicyForView {
	s.UpdateTime = &v
	return s
}

func (s *AlertEventIntegrationPolicyForView) SetUserId(v string) *AlertEventIntegrationPolicyForView {
	s.UserId = &v
	return s
}

func (s *AlertEventIntegrationPolicyForView) SetWorkspace(v string) *AlertEventIntegrationPolicyForView {
	s.Workspace = &v
	return s
}

type AlertRuleAction struct {
	Actions []*string `json:"actions,omitempty" xml:"actions,omitempty" type:"Repeated"`
}

func (s AlertRuleAction) String() string {
	return tea.Prettify(s)
}

func (s AlertRuleAction) GoString() string {
	return s.String()
}

func (s *AlertRuleAction) SetActions(v []*string) *AlertRuleAction {
	s.Actions = v
	return s
}

type AlertRuleAlertMetricFilterDef struct {
	Dim           *string                                       `json:"dim,omitempty" xml:"dim,omitempty"`
	DimDisabled   *bool                                         `json:"dimDisabled,omitempty" xml:"dimDisabled,omitempty"`
	DisplayNameCn *string                                       `json:"displayNameCn,omitempty" xml:"displayNameCn,omitempty"`
	DisplayNameEn *string                                       `json:"displayNameEn,omitempty" xml:"displayNameEn,omitempty"`
	Hidden        *bool                                         `json:"hidden,omitempty" xml:"hidden,omitempty"`
	LabelDisabled *bool                                         `json:"labelDisabled,omitempty" xml:"labelDisabled,omitempty"`
	Opt           *string                                       `json:"opt,omitempty" xml:"opt,omitempty"`
	SupportedOpts []*AlertRuleAlertMetricFilterDefSupportedOpts `json:"supportedOpts,omitempty" xml:"supportedOpts,omitempty" type:"Repeated"`
}

func (s AlertRuleAlertMetricFilterDef) String() string {
	return tea.Prettify(s)
}

func (s AlertRuleAlertMetricFilterDef) GoString() string {
	return s.String()
}

func (s *AlertRuleAlertMetricFilterDef) SetDim(v string) *AlertRuleAlertMetricFilterDef {
	s.Dim = &v
	return s
}

func (s *AlertRuleAlertMetricFilterDef) SetDimDisabled(v bool) *AlertRuleAlertMetricFilterDef {
	s.DimDisabled = &v
	return s
}

func (s *AlertRuleAlertMetricFilterDef) SetDisplayNameCn(v string) *AlertRuleAlertMetricFilterDef {
	s.DisplayNameCn = &v
	return s
}

func (s *AlertRuleAlertMetricFilterDef) SetDisplayNameEn(v string) *AlertRuleAlertMetricFilterDef {
	s.DisplayNameEn = &v
	return s
}

func (s *AlertRuleAlertMetricFilterDef) SetHidden(v bool) *AlertRuleAlertMetricFilterDef {
	s.Hidden = &v
	return s
}

func (s *AlertRuleAlertMetricFilterDef) SetLabelDisabled(v bool) *AlertRuleAlertMetricFilterDef {
	s.LabelDisabled = &v
	return s
}

func (s *AlertRuleAlertMetricFilterDef) SetOpt(v string) *AlertRuleAlertMetricFilterDef {
	s.Opt = &v
	return s
}

func (s *AlertRuleAlertMetricFilterDef) SetSupportedOpts(v []*AlertRuleAlertMetricFilterDefSupportedOpts) *AlertRuleAlertMetricFilterDef {
	s.SupportedOpts = v
	return s
}

type AlertRuleAlertMetricFilterDefSupportedOpts struct {
	DisplayNameCn *string `json:"displayNameCn,omitempty" xml:"displayNameCn,omitempty"`
	DisplayNameEn *string `json:"displayNameEn,omitempty" xml:"displayNameEn,omitempty"`
	Value         *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s AlertRuleAlertMetricFilterDefSupportedOpts) String() string {
	return tea.Prettify(s)
}

func (s AlertRuleAlertMetricFilterDefSupportedOpts) GoString() string {
	return s.String()
}

func (s *AlertRuleAlertMetricFilterDefSupportedOpts) SetDisplayNameCn(v string) *AlertRuleAlertMetricFilterDefSupportedOpts {
	s.DisplayNameCn = &v
	return s
}

func (s *AlertRuleAlertMetricFilterDefSupportedOpts) SetDisplayNameEn(v string) *AlertRuleAlertMetricFilterDefSupportedOpts {
	s.DisplayNameEn = &v
	return s
}

func (s *AlertRuleAlertMetricFilterDefSupportedOpts) SetValue(v string) *AlertRuleAlertMetricFilterDefSupportedOpts {
	s.Value = &v
	return s
}

type AlertRuleAlertMetricInput struct {
	FilterValues []*AlertRuleAlertMetricInputFilterValue `json:"filterValues,omitempty" xml:"filterValues,omitempty" type:"Repeated"`
	GroupId      *string                                 `json:"groupId,omitempty" xml:"groupId,omitempty"`
	MetricId     *string                                 `json:"metricId,omitempty" xml:"metricId,omitempty"`
	ParamValues  []*AlertRuleAlertMetricInputParamValue  `json:"paramValues,omitempty" xml:"paramValues,omitempty" type:"Repeated"`
}

func (s AlertRuleAlertMetricInput) String() string {
	return tea.Prettify(s)
}

func (s AlertRuleAlertMetricInput) GoString() string {
	return s.String()
}

func (s *AlertRuleAlertMetricInput) SetFilterValues(v []*AlertRuleAlertMetricInputFilterValue) *AlertRuleAlertMetricInput {
	s.FilterValues = v
	return s
}

func (s *AlertRuleAlertMetricInput) SetGroupId(v string) *AlertRuleAlertMetricInput {
	s.GroupId = &v
	return s
}

func (s *AlertRuleAlertMetricInput) SetMetricId(v string) *AlertRuleAlertMetricInput {
	s.MetricId = &v
	return s
}

func (s *AlertRuleAlertMetricInput) SetParamValues(v []*AlertRuleAlertMetricInputParamValue) *AlertRuleAlertMetricInput {
	s.ParamValues = v
	return s
}

type AlertRuleAlertMetricInputFilterValue struct {
	// This parameter is required.
	Dim *string `json:"dim,omitempty" xml:"dim,omitempty"`
	// This parameter is required.
	Opt   *string `json:"opt,omitempty" xml:"opt,omitempty"`
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s AlertRuleAlertMetricInputFilterValue) String() string {
	return tea.Prettify(s)
}

func (s AlertRuleAlertMetricInputFilterValue) GoString() string {
	return s.String()
}

func (s *AlertRuleAlertMetricInputFilterValue) SetDim(v string) *AlertRuleAlertMetricInputFilterValue {
	s.Dim = &v
	return s
}

func (s *AlertRuleAlertMetricInputFilterValue) SetOpt(v string) *AlertRuleAlertMetricInputFilterValue {
	s.Opt = &v
	return s
}

func (s *AlertRuleAlertMetricInputFilterValue) SetValue(v string) *AlertRuleAlertMetricInputFilterValue {
	s.Value = &v
	return s
}

type AlertRuleAlertMetricInputParamValue struct {
	// This parameter is required.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// This parameter is required.
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s AlertRuleAlertMetricInputParamValue) String() string {
	return tea.Prettify(s)
}

func (s AlertRuleAlertMetricInputParamValue) GoString() string {
	return s.String()
}

func (s *AlertRuleAlertMetricInputParamValue) SetName(v string) *AlertRuleAlertMetricInputParamValue {
	s.Name = &v
	return s
}

func (s *AlertRuleAlertMetricInputParamValue) SetValue(v string) *AlertRuleAlertMetricInputParamValue {
	s.Value = &v
	return s
}

type AlertRuleAlertMetricParamDef struct {
	MaxWidth      *int32                                `json:"maxWidth,omitempty" xml:"maxWidth,omitempty"`
	MinWidth      *int32                                `json:"minWidth,omitempty" xml:"minWidth,omitempty"`
	Name          *string                               `json:"name,omitempty" xml:"name,omitempty"`
	PlaceholderCn *string                               `json:"placeholderCn,omitempty" xml:"placeholderCn,omitempty"`
	PlaceholderEn *string                               `json:"placeholderEn,omitempty" xml:"placeholderEn,omitempty"`
	Type          *string                               `json:"type,omitempty" xml:"type,omitempty"`
	Value         *string                               `json:"value,omitempty" xml:"value,omitempty"`
	Values        []*AlertRuleAlertMetricParamDefValues `json:"values,omitempty" xml:"values,omitempty" type:"Repeated"`
}

func (s AlertRuleAlertMetricParamDef) String() string {
	return tea.Prettify(s)
}

func (s AlertRuleAlertMetricParamDef) GoString() string {
	return s.String()
}

func (s *AlertRuleAlertMetricParamDef) SetMaxWidth(v int32) *AlertRuleAlertMetricParamDef {
	s.MaxWidth = &v
	return s
}

func (s *AlertRuleAlertMetricParamDef) SetMinWidth(v int32) *AlertRuleAlertMetricParamDef {
	s.MinWidth = &v
	return s
}

func (s *AlertRuleAlertMetricParamDef) SetName(v string) *AlertRuleAlertMetricParamDef {
	s.Name = &v
	return s
}

func (s *AlertRuleAlertMetricParamDef) SetPlaceholderCn(v string) *AlertRuleAlertMetricParamDef {
	s.PlaceholderCn = &v
	return s
}

func (s *AlertRuleAlertMetricParamDef) SetPlaceholderEn(v string) *AlertRuleAlertMetricParamDef {
	s.PlaceholderEn = &v
	return s
}

func (s *AlertRuleAlertMetricParamDef) SetType(v string) *AlertRuleAlertMetricParamDef {
	s.Type = &v
	return s
}

func (s *AlertRuleAlertMetricParamDef) SetValue(v string) *AlertRuleAlertMetricParamDef {
	s.Value = &v
	return s
}

func (s *AlertRuleAlertMetricParamDef) SetValues(v []*AlertRuleAlertMetricParamDefValues) *AlertRuleAlertMetricParamDef {
	s.Values = v
	return s
}

type AlertRuleAlertMetricParamDefValues struct {
	LabelCn *string `json:"labelCn,omitempty" xml:"labelCn,omitempty"`
	LabelEn *string `json:"labelEn,omitempty" xml:"labelEn,omitempty"`
	Value   *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s AlertRuleAlertMetricParamDefValues) String() string {
	return tea.Prettify(s)
}

func (s AlertRuleAlertMetricParamDefValues) GoString() string {
	return s.String()
}

func (s *AlertRuleAlertMetricParamDefValues) SetLabelCn(v string) *AlertRuleAlertMetricParamDefValues {
	s.LabelCn = &v
	return s
}

func (s *AlertRuleAlertMetricParamDefValues) SetLabelEn(v string) *AlertRuleAlertMetricParamDefValues {
	s.LabelEn = &v
	return s
}

func (s *AlertRuleAlertMetricParamDefValues) SetValue(v string) *AlertRuleAlertMetricParamDefValues {
	s.Value = &v
	return s
}

type AlertRuleCondition struct {
	// type=SLS_CONDITION时指定，满足条件几次后告警，默认为1
	AlertCount *int32 `json:"alertCount,omitempty" xml:"alertCount,omitempty"`
	// type=SLS_CONDITION时指定
	CaseList          []*AlertRuleConditionCaseList    `json:"caseList,omitempty" xml:"caseList,omitempty" type:"Repeated"`
	CompareList       []*AlertRuleConditionCompareList `json:"compareList,omitempty" xml:"compareList,omitempty" type:"Repeated"`
	NoDataAppendValue *string                          `json:"noDataAppendValue,omitempty" xml:"noDataAppendValue,omitempty"`
	// 无数据时按什么级别告警，不指定则不对无数据报警
	NodataAlertLevel *string `json:"nodataAlertLevel,omitempty" xml:"nodataAlertLevel,omitempty"`
	Relation         *string `json:"relation,omitempty" xml:"relation,omitempty"`
	// 规则条件类型，可选值：SLS_CONDITION
	//
	// This parameter is required.
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s AlertRuleCondition) String() string {
	return tea.Prettify(s)
}

func (s AlertRuleCondition) GoString() string {
	return s.String()
}

func (s *AlertRuleCondition) SetAlertCount(v int32) *AlertRuleCondition {
	s.AlertCount = &v
	return s
}

func (s *AlertRuleCondition) SetCaseList(v []*AlertRuleConditionCaseList) *AlertRuleCondition {
	s.CaseList = v
	return s
}

func (s *AlertRuleCondition) SetCompareList(v []*AlertRuleConditionCompareList) *AlertRuleCondition {
	s.CompareList = v
	return s
}

func (s *AlertRuleCondition) SetNoDataAppendValue(v string) *AlertRuleCondition {
	s.NoDataAppendValue = &v
	return s
}

func (s *AlertRuleCondition) SetNodataAlertLevel(v string) *AlertRuleCondition {
	s.NodataAlertLevel = &v
	return s
}

func (s *AlertRuleCondition) SetRelation(v string) *AlertRuleCondition {
	s.Relation = &v
	return s
}

func (s *AlertRuleCondition) SetType(v string) *AlertRuleCondition {
	s.Type = &v
	return s
}

type AlertRuleConditionCaseList struct {
	Condition      *string `json:"condition,omitempty" xml:"condition,omitempty"`
	CountCondition *string `json:"countCondition,omitempty" xml:"countCondition,omitempty"`
	Level          *string `json:"level,omitempty" xml:"level,omitempty"`
	Type           *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s AlertRuleConditionCaseList) String() string {
	return tea.Prettify(s)
}

func (s AlertRuleConditionCaseList) GoString() string {
	return s.String()
}

func (s *AlertRuleConditionCaseList) SetCondition(v string) *AlertRuleConditionCaseList {
	s.Condition = &v
	return s
}

func (s *AlertRuleConditionCaseList) SetCountCondition(v string) *AlertRuleConditionCaseList {
	s.CountCondition = &v
	return s
}

func (s *AlertRuleConditionCaseList) SetLevel(v string) *AlertRuleConditionCaseList {
	s.Level = &v
	return s
}

func (s *AlertRuleConditionCaseList) SetType(v string) *AlertRuleConditionCaseList {
	s.Type = &v
	return s
}

type AlertRuleConditionCompareList struct {
	Aggregate      *string                                        `json:"aggregate,omitempty" xml:"aggregate,omitempty"`
	Oper           *string                                        `json:"oper,omitempty" xml:"oper,omitempty"`
	Value          *float64                                       `json:"value,omitempty" xml:"value,omitempty"`
	ValueLevelList []*AlertRuleConditionCompareListValueLevelList `json:"valueLevelList,omitempty" xml:"valueLevelList,omitempty" type:"Repeated"`
	YoyTimeUnit    *string                                        `json:"yoyTimeUnit,omitempty" xml:"yoyTimeUnit,omitempty"`
	YoyTimeValue   *int32                                         `json:"yoyTimeValue,omitempty" xml:"yoyTimeValue,omitempty"`
}

func (s AlertRuleConditionCompareList) String() string {
	return tea.Prettify(s)
}

func (s AlertRuleConditionCompareList) GoString() string {
	return s.String()
}

func (s *AlertRuleConditionCompareList) SetAggregate(v string) *AlertRuleConditionCompareList {
	s.Aggregate = &v
	return s
}

func (s *AlertRuleConditionCompareList) SetOper(v string) *AlertRuleConditionCompareList {
	s.Oper = &v
	return s
}

func (s *AlertRuleConditionCompareList) SetValue(v float64) *AlertRuleConditionCompareList {
	s.Value = &v
	return s
}

func (s *AlertRuleConditionCompareList) SetValueLevelList(v []*AlertRuleConditionCompareListValueLevelList) *AlertRuleConditionCompareList {
	s.ValueLevelList = v
	return s
}

func (s *AlertRuleConditionCompareList) SetYoyTimeUnit(v string) *AlertRuleConditionCompareList {
	s.YoyTimeUnit = &v
	return s
}

func (s *AlertRuleConditionCompareList) SetYoyTimeValue(v int32) *AlertRuleConditionCompareList {
	s.YoyTimeValue = &v
	return s
}

type AlertRuleConditionCompareListValueLevelList struct {
	Level *string  `json:"level,omitempty" xml:"level,omitempty"`
	Value *float64 `json:"value,omitempty" xml:"value,omitempty"`
}

func (s AlertRuleConditionCompareListValueLevelList) String() string {
	return tea.Prettify(s)
}

func (s AlertRuleConditionCompareListValueLevelList) GoString() string {
	return s.String()
}

func (s *AlertRuleConditionCompareListValueLevelList) SetLevel(v string) *AlertRuleConditionCompareListValueLevelList {
	s.Level = &v
	return s
}

func (s *AlertRuleConditionCompareListValueLevelList) SetValue(v float64) *AlertRuleConditionCompareListValueLevelList {
	s.Value = &v
	return s
}

type AlertRuleDataSource struct {
	AppType *string                      `json:"appType,omitempty" xml:"appType,omitempty"`
	DsList  []*AlertRuleDataSourceDsList `json:"dsList,omitempty" xml:"dsList,omitempty" type:"Repeated"`
	// 实例id，当type=PROMETHEUS_DS/ENTERPRISE_DS时必填，为prometheus实例的clusterId或指标仓库名称
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	Namespace  *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
	RegionId   *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// 数据源类型
	//
	// This parameter is required.
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s AlertRuleDataSource) String() string {
	return tea.Prettify(s)
}

func (s AlertRuleDataSource) GoString() string {
	return s.String()
}

func (s *AlertRuleDataSource) SetAppType(v string) *AlertRuleDataSource {
	s.AppType = &v
	return s
}

func (s *AlertRuleDataSource) SetDsList(v []*AlertRuleDataSourceDsList) *AlertRuleDataSource {
	s.DsList = v
	return s
}

func (s *AlertRuleDataSource) SetInstanceId(v string) *AlertRuleDataSource {
	s.InstanceId = &v
	return s
}

func (s *AlertRuleDataSource) SetNamespace(v string) *AlertRuleDataSource {
	s.Namespace = &v
	return s
}

func (s *AlertRuleDataSource) SetRegionId(v string) *AlertRuleDataSource {
	s.RegionId = &v
	return s
}

func (s *AlertRuleDataSource) SetType(v string) *AlertRuleDataSource {
	s.Type = &v
	return s
}

type AlertRuleDataSourceDsList struct {
	Project  *string `json:"project,omitempty" xml:"project,omitempty"`
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	Store    *string `json:"store,omitempty" xml:"store,omitempty"`
	Type     *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s AlertRuleDataSourceDsList) String() string {
	return tea.Prettify(s)
}

func (s AlertRuleDataSourceDsList) GoString() string {
	return s.String()
}

func (s *AlertRuleDataSourceDsList) SetProject(v string) *AlertRuleDataSourceDsList {
	s.Project = &v
	return s
}

func (s *AlertRuleDataSourceDsList) SetRegionId(v string) *AlertRuleDataSourceDsList {
	s.RegionId = &v
	return s
}

func (s *AlertRuleDataSourceDsList) SetStore(v string) *AlertRuleDataSourceDsList {
	s.Store = &v
	return s
}

func (s *AlertRuleDataSourceDsList) SetType(v string) *AlertRuleDataSourceDsList {
	s.Type = &v
	return s
}

type AlertRuleLabelFilter struct {
	Labels map[string]*string `json:"labels,omitempty" xml:"labels,omitempty"`
	Opt    *string            `json:"opt,omitempty" xml:"opt,omitempty"`
}

func (s AlertRuleLabelFilter) String() string {
	return tea.Prettify(s)
}

func (s AlertRuleLabelFilter) GoString() string {
	return s.String()
}

func (s *AlertRuleLabelFilter) SetLabels(v map[string]*string) *AlertRuleLabelFilter {
	s.Labels = v
	return s
}

func (s *AlertRuleLabelFilter) SetOpt(v string) *AlertRuleLabelFilter {
	s.Opt = &v
	return s
}

type AlertRuleNotification struct {
	Contacts       []*string          `json:"contacts,omitempty" xml:"contacts,omitempty" type:"Repeated"`
	CustomWebhooks []*string          `json:"customWebhooks,omitempty" xml:"customWebhooks,omitempty" type:"Repeated"`
	DingWebhooks   []*string          `json:"dingWebhooks,omitempty" xml:"dingWebhooks,omitempty" type:"Repeated"`
	FsWebhooks     []*string          `json:"fsWebhooks,omitempty" xml:"fsWebhooks,omitempty" type:"Repeated"`
	Groups         []*string          `json:"groups,omitempty" xml:"groups,omitempty" type:"Repeated"`
	NotifyTime     *AlertRuleTimeSpan `json:"notifyTime,omitempty" xml:"notifyTime,omitempty"`
	SilenceTime    *int64             `json:"silenceTime,omitempty" xml:"silenceTime,omitempty"`
	SlackWebhooks  []*string          `json:"slackWebhooks,omitempty" xml:"slackWebhooks,omitempty" type:"Repeated"`
	WxWebhooks     []*string          `json:"wxWebhooks,omitempty" xml:"wxWebhooks,omitempty" type:"Repeated"`
}

func (s AlertRuleNotification) String() string {
	return tea.Prettify(s)
}

func (s AlertRuleNotification) GoString() string {
	return s.String()
}

func (s *AlertRuleNotification) SetContacts(v []*string) *AlertRuleNotification {
	s.Contacts = v
	return s
}

func (s *AlertRuleNotification) SetCustomWebhooks(v []*string) *AlertRuleNotification {
	s.CustomWebhooks = v
	return s
}

func (s *AlertRuleNotification) SetDingWebhooks(v []*string) *AlertRuleNotification {
	s.DingWebhooks = v
	return s
}

func (s *AlertRuleNotification) SetFsWebhooks(v []*string) *AlertRuleNotification {
	s.FsWebhooks = v
	return s
}

func (s *AlertRuleNotification) SetGroups(v []*string) *AlertRuleNotification {
	s.Groups = v
	return s
}

func (s *AlertRuleNotification) SetNotifyTime(v *AlertRuleTimeSpan) *AlertRuleNotification {
	s.NotifyTime = v
	return s
}

func (s *AlertRuleNotification) SetSilenceTime(v int64) *AlertRuleNotification {
	s.SilenceTime = &v
	return s
}

func (s *AlertRuleNotification) SetSlackWebhooks(v []*string) *AlertRuleNotification {
	s.SlackWebhooks = v
	return s
}

func (s *AlertRuleNotification) SetWxWebhooks(v []*string) *AlertRuleNotification {
	s.WxWebhooks = v
	return s
}

type AlertRuleNotificationFilter struct {
	Contacts       []*string `json:"contacts,omitempty" xml:"contacts,omitempty" type:"Repeated"`
	CustomWebhooks []*string `json:"customWebhooks,omitempty" xml:"customWebhooks,omitempty" type:"Repeated"`
	DingWebhooks   []*string `json:"dingWebhooks,omitempty" xml:"dingWebhooks,omitempty" type:"Repeated"`
	FsWebhooks     []*string `json:"fsWebhooks,omitempty" xml:"fsWebhooks,omitempty" type:"Repeated"`
	Groups         []*string `json:"groups,omitempty" xml:"groups,omitempty" type:"Repeated"`
	SlackWebhooks  []*string `json:"slackWebhooks,omitempty" xml:"slackWebhooks,omitempty" type:"Repeated"`
	WxWebhooks     []*string `json:"wxWebhooks,omitempty" xml:"wxWebhooks,omitempty" type:"Repeated"`
}

func (s AlertRuleNotificationFilter) String() string {
	return tea.Prettify(s)
}

func (s AlertRuleNotificationFilter) GoString() string {
	return s.String()
}

func (s *AlertRuleNotificationFilter) SetContacts(v []*string) *AlertRuleNotificationFilter {
	s.Contacts = v
	return s
}

func (s *AlertRuleNotificationFilter) SetCustomWebhooks(v []*string) *AlertRuleNotificationFilter {
	s.CustomWebhooks = v
	return s
}

func (s *AlertRuleNotificationFilter) SetDingWebhooks(v []*string) *AlertRuleNotificationFilter {
	s.DingWebhooks = v
	return s
}

func (s *AlertRuleNotificationFilter) SetFsWebhooks(v []*string) *AlertRuleNotificationFilter {
	s.FsWebhooks = v
	return s
}

func (s *AlertRuleNotificationFilter) SetGroups(v []*string) *AlertRuleNotificationFilter {
	s.Groups = v
	return s
}

func (s *AlertRuleNotificationFilter) SetSlackWebhooks(v []*string) *AlertRuleNotificationFilter {
	s.SlackWebhooks = v
	return s
}

func (s *AlertRuleNotificationFilter) SetWxWebhooks(v []*string) *AlertRuleNotificationFilter {
	s.WxWebhooks = v
	return s
}

type AlertRuleQuery struct {
	CheckAfterDataComplete *bool                    `json:"checkAfterDataComplete,omitempty" xml:"checkAfterDataComplete,omitempty"`
	Duration               *int64                   `json:"duration,omitempty" xml:"duration,omitempty"`
	Expr                   *string                  `json:"expr,omitempty" xml:"expr,omitempty"`
	FirstJoin              *AlertRuleSlsQueryJoin   `json:"firstJoin,omitempty" xml:"firstJoin,omitempty"`
	GroupFieldList         []*string                `json:"groupFieldList,omitempty" xml:"groupFieldList,omitempty" type:"Repeated"`
	GroupType              *string                  `json:"groupType,omitempty" xml:"groupType,omitempty"`
	Queries                []*AlertRuleQueryQueries `json:"queries,omitempty" xml:"queries,omitempty" type:"Repeated"`
	SecondJoin             *AlertRuleSlsQueryJoin   `json:"secondJoin,omitempty" xml:"secondJoin,omitempty"`
	// 查询类型
	//
	// This parameter is required.
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s AlertRuleQuery) String() string {
	return tea.Prettify(s)
}

func (s AlertRuleQuery) GoString() string {
	return s.String()
}

func (s *AlertRuleQuery) SetCheckAfterDataComplete(v bool) *AlertRuleQuery {
	s.CheckAfterDataComplete = &v
	return s
}

func (s *AlertRuleQuery) SetDuration(v int64) *AlertRuleQuery {
	s.Duration = &v
	return s
}

func (s *AlertRuleQuery) SetExpr(v string) *AlertRuleQuery {
	s.Expr = &v
	return s
}

func (s *AlertRuleQuery) SetFirstJoin(v *AlertRuleSlsQueryJoin) *AlertRuleQuery {
	s.FirstJoin = v
	return s
}

func (s *AlertRuleQuery) SetGroupFieldList(v []*string) *AlertRuleQuery {
	s.GroupFieldList = v
	return s
}

func (s *AlertRuleQuery) SetGroupType(v string) *AlertRuleQuery {
	s.GroupType = &v
	return s
}

func (s *AlertRuleQuery) SetQueries(v []*AlertRuleQueryQueries) *AlertRuleQuery {
	s.Queries = v
	return s
}

func (s *AlertRuleQuery) SetSecondJoin(v *AlertRuleSlsQueryJoin) *AlertRuleQuery {
	s.SecondJoin = v
	return s
}

func (s *AlertRuleQuery) SetType(v string) *AlertRuleQuery {
	s.Type = &v
	return s
}

type AlertRuleQueryQueries struct {
	ApmAlertMetricId *string                            `json:"apmAlertMetricId,omitempty" xml:"apmAlertMetricId,omitempty"`
	ApmFilters       []*AlertRuleQueryQueriesApmFilters `json:"apmFilters,omitempty" xml:"apmFilters,omitempty" type:"Repeated"`
	ApmGroupBy       []*string                          `json:"apmGroupBy,omitempty" xml:"apmGroupBy,omitempty" type:"Repeated"`
	Duration         *int64                             `json:"duration,omitempty" xml:"duration,omitempty"`
	// 时间偏移结束时间(相对)，如果指定了start、end，则不指定window。
	End *int64 `json:"end,omitempty" xml:"end,omitempty"`
	// 查询表达式
	Expr *string `json:"expr,omitempty" xml:"expr,omitempty"`
	// sls查询的时间偏移开始时间(相对)，如果指定了start、end，则不指定window。  例如：start=15， timeUnit=minute，表示15分钟前
	Start *int64 `json:"start,omitempty" xml:"start,omitempty"`
	// start和end、window的时间单位： day/hour/minute/second
	TimeUnit *string `json:"timeUnit,omitempty" xml:"timeUnit,omitempty"`
	// 整点时间查询区间。  如果指定了window则不指定start、end
	Window *int64 `json:"window,omitempty" xml:"window,omitempty"`
}

func (s AlertRuleQueryQueries) String() string {
	return tea.Prettify(s)
}

func (s AlertRuleQueryQueries) GoString() string {
	return s.String()
}

func (s *AlertRuleQueryQueries) SetApmAlertMetricId(v string) *AlertRuleQueryQueries {
	s.ApmAlertMetricId = &v
	return s
}

func (s *AlertRuleQueryQueries) SetApmFilters(v []*AlertRuleQueryQueriesApmFilters) *AlertRuleQueryQueries {
	s.ApmFilters = v
	return s
}

func (s *AlertRuleQueryQueries) SetApmGroupBy(v []*string) *AlertRuleQueryQueries {
	s.ApmGroupBy = v
	return s
}

func (s *AlertRuleQueryQueries) SetDuration(v int64) *AlertRuleQueryQueries {
	s.Duration = &v
	return s
}

func (s *AlertRuleQueryQueries) SetEnd(v int64) *AlertRuleQueryQueries {
	s.End = &v
	return s
}

func (s *AlertRuleQueryQueries) SetExpr(v string) *AlertRuleQueryQueries {
	s.Expr = &v
	return s
}

func (s *AlertRuleQueryQueries) SetStart(v int64) *AlertRuleQueryQueries {
	s.Start = &v
	return s
}

func (s *AlertRuleQueryQueries) SetTimeUnit(v string) *AlertRuleQueryQueries {
	s.TimeUnit = &v
	return s
}

func (s *AlertRuleQueryQueries) SetWindow(v int64) *AlertRuleQueryQueries {
	s.Window = &v
	return s
}

type AlertRuleQueryQueriesApmFilters struct {
	Dim   *string `json:"dim,omitempty" xml:"dim,omitempty"`
	Type  *string `json:"type,omitempty" xml:"type,omitempty"`
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s AlertRuleQueryQueriesApmFilters) String() string {
	return tea.Prettify(s)
}

func (s AlertRuleQueryQueriesApmFilters) GoString() string {
	return s.String()
}

func (s *AlertRuleQueryQueriesApmFilters) SetDim(v string) *AlertRuleQueryQueriesApmFilters {
	s.Dim = &v
	return s
}

func (s *AlertRuleQueryQueriesApmFilters) SetType(v string) *AlertRuleQueryQueriesApmFilters {
	s.Type = &v
	return s
}

func (s *AlertRuleQueryQueriesApmFilters) SetValue(v string) *AlertRuleQueryQueriesApmFilters {
	s.Value = &v
	return s
}

type AlertRuleSend struct {
	Action       *AlertRuleAction       `json:"action,omitempty" xml:"action,omitempty"`
	Notification *AlertRuleNotification `json:"notification,omitempty" xml:"notification,omitempty"`
}

func (s AlertRuleSend) String() string {
	return tea.Prettify(s)
}

func (s AlertRuleSend) GoString() string {
	return s.String()
}

func (s *AlertRuleSend) SetAction(v *AlertRuleAction) *AlertRuleSend {
	s.Action = v
	return s
}

func (s *AlertRuleSend) SetNotification(v *AlertRuleNotification) *AlertRuleSend {
	s.Notification = v
	return s
}

type AlertRuleSlsQueryJoin struct {
	Conditions []*AlertRuleSlsQueryJoinConditions `json:"conditions,omitempty" xml:"conditions,omitempty" type:"Repeated"`
	// 集合操作类型。
	//
	//   ● CrossJoin： 笛卡尔积
	//
	//   ● FullJoin：全联
	//
	//   ● InnerJoin：内联
	//
	//   ● LeftExclude： 左斥
	//
	//   ● RightExclude：右斥
	//
	//   ● LeftJoin：左联
	//
	//   ● RightJoin：右联
	//
	//   ● NoJoin：不合并
	//
	//   ● Concat： 拼接
	//
	//   https://help.aliyun.com/zh/sls/user-guide/set-query-statistics-statement
	//
	// This parameter is required.
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s AlertRuleSlsQueryJoin) String() string {
	return tea.Prettify(s)
}

func (s AlertRuleSlsQueryJoin) GoString() string {
	return s.String()
}

func (s *AlertRuleSlsQueryJoin) SetConditions(v []*AlertRuleSlsQueryJoinConditions) *AlertRuleSlsQueryJoin {
	s.Conditions = v
	return s
}

func (s *AlertRuleSlsQueryJoin) SetType(v string) *AlertRuleSlsQueryJoin {
	s.Type = &v
	return s
}

type AlertRuleSlsQueryJoinConditions struct {
	// 条件的左操作参数，格式为$<query_idx>.<结果集字段名>
	FirstField *string `json:"firstField,omitempty" xml:"firstField,omitempty"`
	// <, >, ==, !=, <=, >=
	Oper *string `json:"oper,omitempty" xml:"oper,omitempty"`
	// 条件的右操作参数，格式为$<query_idx>.<结果集字段名>
	SecondField *string `json:"secondField,omitempty" xml:"secondField,omitempty"`
}

func (s AlertRuleSlsQueryJoinConditions) String() string {
	return tea.Prettify(s)
}

func (s AlertRuleSlsQueryJoinConditions) GoString() string {
	return s.String()
}

func (s *AlertRuleSlsQueryJoinConditions) SetFirstField(v string) *AlertRuleSlsQueryJoinConditions {
	s.FirstField = &v
	return s
}

func (s *AlertRuleSlsQueryJoinConditions) SetOper(v string) *AlertRuleSlsQueryJoinConditions {
	s.Oper = &v
	return s
}

func (s *AlertRuleSlsQueryJoinConditions) SetSecondField(v string) *AlertRuleSlsQueryJoinConditions {
	s.SecondField = &v
	return s
}

type AlertRuleTimeSpan struct {
	DayOfWeek []*int32 `json:"dayOfWeek,omitempty" xml:"dayOfWeek,omitempty" type:"Repeated"`
	EndTime   *string  `json:"endTime,omitempty" xml:"endTime,omitempty"`
	GmtOffset *string  `json:"gmtOffset,omitempty" xml:"gmtOffset,omitempty"`
	StartTime *string  `json:"startTime,omitempty" xml:"startTime,omitempty"`
}

func (s AlertRuleTimeSpan) String() string {
	return tea.Prettify(s)
}

func (s AlertRuleTimeSpan) GoString() string {
	return s.String()
}

func (s *AlertRuleTimeSpan) SetDayOfWeek(v []*int32) *AlertRuleTimeSpan {
	s.DayOfWeek = v
	return s
}

func (s *AlertRuleTimeSpan) SetEndTime(v string) *AlertRuleTimeSpan {
	s.EndTime = &v
	return s
}

func (s *AlertRuleTimeSpan) SetGmtOffset(v string) *AlertRuleTimeSpan {
	s.GmtOffset = &v
	return s
}

func (s *AlertRuleTimeSpan) SetStartTime(v string) *AlertRuleTimeSpan {
	s.StartTime = &v
	return s
}

type BizTraceConfig struct {
	AdvancedConfig *string `json:"advancedConfig,omitempty" xml:"advancedConfig,omitempty"`
	BizTraceCode   *string `json:"bizTraceCode,omitempty" xml:"bizTraceCode,omitempty"`
	BizTraceId     *string `json:"bizTraceId,omitempty" xml:"bizTraceId,omitempty"`
	BizTraceName   *string `json:"bizTraceName,omitempty" xml:"bizTraceName,omitempty"`
	CreateTime     *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	RegionId       *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	RuleConfig     *string `json:"ruleConfig,omitempty" xml:"ruleConfig,omitempty"`
	Workspace      *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s BizTraceConfig) String() string {
	return tea.Prettify(s)
}

func (s BizTraceConfig) GoString() string {
	return s.String()
}

func (s *BizTraceConfig) SetAdvancedConfig(v string) *BizTraceConfig {
	s.AdvancedConfig = &v
	return s
}

func (s *BizTraceConfig) SetBizTraceCode(v string) *BizTraceConfig {
	s.BizTraceCode = &v
	return s
}

func (s *BizTraceConfig) SetBizTraceId(v string) *BizTraceConfig {
	s.BizTraceId = &v
	return s
}

func (s *BizTraceConfig) SetBizTraceName(v string) *BizTraceConfig {
	s.BizTraceName = &v
	return s
}

func (s *BizTraceConfig) SetCreateTime(v string) *BizTraceConfig {
	s.CreateTime = &v
	return s
}

func (s *BizTraceConfig) SetRegionId(v string) *BizTraceConfig {
	s.RegionId = &v
	return s
}

func (s *BizTraceConfig) SetRuleConfig(v string) *BizTraceConfig {
	s.RuleConfig = &v
	return s
}

func (s *BizTraceConfig) SetWorkspace(v string) *BizTraceConfig {
	s.Workspace = &v
	return s
}

type DataStorageItem struct {
	DataType  *string `json:"dataType,omitempty" xml:"dataType,omitempty"`
	Project   *string `json:"project,omitempty" xml:"project,omitempty"`
	RegionId  *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	StoreName *string `json:"storeName,omitempty" xml:"storeName,omitempty"`
	StoreType *string `json:"storeType,omitempty" xml:"storeType,omitempty"`
}

func (s DataStorageItem) String() string {
	return tea.Prettify(s)
}

func (s DataStorageItem) GoString() string {
	return s.String()
}

func (s *DataStorageItem) SetDataType(v string) *DataStorageItem {
	s.DataType = &v
	return s
}

func (s *DataStorageItem) SetProject(v string) *DataStorageItem {
	s.Project = &v
	return s
}

func (s *DataStorageItem) SetRegionId(v string) *DataStorageItem {
	s.RegionId = &v
	return s
}

func (s *DataStorageItem) SetStoreName(v string) *DataStorageItem {
	s.StoreName = &v
	return s
}

func (s *DataStorageItem) SetStoreType(v string) *DataStorageItem {
	s.StoreType = &v
	return s
}

type FilterSetting struct {
	Conditions []*FilterSettingConditions `json:"conditions,omitempty" xml:"conditions,omitempty" type:"Repeated"`
	Expression *string                    `json:"expression,omitempty" xml:"expression,omitempty"`
	Relation   *string                    `json:"relation,omitempty" xml:"relation,omitempty"`
}

func (s FilterSetting) String() string {
	return tea.Prettify(s)
}

func (s FilterSetting) GoString() string {
	return s.String()
}

func (s *FilterSetting) SetConditions(v []*FilterSettingConditions) *FilterSetting {
	s.Conditions = v
	return s
}

func (s *FilterSetting) SetExpression(v string) *FilterSetting {
	s.Expression = &v
	return s
}

func (s *FilterSetting) SetRelation(v string) *FilterSetting {
	s.Relation = &v
	return s
}

type FilterSettingConditions struct {
	Field *string `json:"field,omitempty" xml:"field,omitempty"`
	Op    *string `json:"op,omitempty" xml:"op,omitempty"`
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s FilterSettingConditions) String() string {
	return tea.Prettify(s)
}

func (s FilterSettingConditions) GoString() string {
	return s.String()
}

func (s *FilterSettingConditions) SetField(v string) *FilterSettingConditions {
	s.Field = &v
	return s
}

func (s *FilterSettingConditions) SetOp(v string) *FilterSettingConditions {
	s.Op = &v
	return s
}

func (s *FilterSettingConditions) SetValue(v string) *FilterSettingConditions {
	s.Value = &v
	return s
}

type IncidentContactStruct struct {
	Channel     []*string `json:"channel,omitempty" xml:"channel,omitempty" type:"Repeated"`
	ContactId   *string   `json:"contactId,omitempty" xml:"contactId,omitempty"`
	ContactType *string   `json:"contactType,omitempty" xml:"contactType,omitempty"`
}

func (s IncidentContactStruct) String() string {
	return tea.Prettify(s)
}

func (s IncidentContactStruct) GoString() string {
	return s.String()
}

func (s *IncidentContactStruct) SetChannel(v []*string) *IncidentContactStruct {
	s.Channel = v
	return s
}

func (s *IncidentContactStruct) SetContactId(v string) *IncidentContactStruct {
	s.ContactId = &v
	return s
}

func (s *IncidentContactStruct) SetContactType(v string) *IncidentContactStruct {
	s.ContactType = &v
	return s
}

type IncidentEscalationStageStruct struct {
	Contact             []*IncidentContactStruct `json:"contact,omitempty" xml:"contact,omitempty" type:"Repeated"`
	CycleNotifyCount    *int32                   `json:"cycleNotifyCount,omitempty" xml:"cycleNotifyCount,omitempty"`
	CycleNotifyTime     *int32                   `json:"cycleNotifyTime,omitempty" xml:"cycleNotifyTime,omitempty"`
	Description         *string                  `json:"description,omitempty" xml:"description,omitempty"`
	EffectTime          *string                  `json:"effectTime,omitempty" xml:"effectTime,omitempty"`
	Name                *string                  `json:"name,omitempty" xml:"name,omitempty"`
	StageIndex          *int32                   `json:"stageIndex,omitempty" xml:"stageIndex,omitempty"`
	TimeZone            *string                  `json:"timeZone,omitempty" xml:"timeZone,omitempty"`
	WaitToNextStageTime *int32                   `json:"waitToNextStageTime,omitempty" xml:"waitToNextStageTime,omitempty"`
}

func (s IncidentEscalationStageStruct) String() string {
	return tea.Prettify(s)
}

func (s IncidentEscalationStageStruct) GoString() string {
	return s.String()
}

func (s *IncidentEscalationStageStruct) SetContact(v []*IncidentContactStruct) *IncidentEscalationStageStruct {
	s.Contact = v
	return s
}

func (s *IncidentEscalationStageStruct) SetCycleNotifyCount(v int32) *IncidentEscalationStageStruct {
	s.CycleNotifyCount = &v
	return s
}

func (s *IncidentEscalationStageStruct) SetCycleNotifyTime(v int32) *IncidentEscalationStageStruct {
	s.CycleNotifyTime = &v
	return s
}

func (s *IncidentEscalationStageStruct) SetDescription(v string) *IncidentEscalationStageStruct {
	s.Description = &v
	return s
}

func (s *IncidentEscalationStageStruct) SetEffectTime(v string) *IncidentEscalationStageStruct {
	s.EffectTime = &v
	return s
}

func (s *IncidentEscalationStageStruct) SetName(v string) *IncidentEscalationStageStruct {
	s.Name = &v
	return s
}

func (s *IncidentEscalationStageStruct) SetStageIndex(v int32) *IncidentEscalationStageStruct {
	s.StageIndex = &v
	return s
}

func (s *IncidentEscalationStageStruct) SetTimeZone(v string) *IncidentEscalationStageStruct {
	s.TimeZone = &v
	return s
}

func (s *IncidentEscalationStageStruct) SetWaitToNextStageTime(v int32) *IncidentEscalationStageStruct {
	s.WaitToNextStageTime = &v
	return s
}

type IncidentEscalationStruct struct {
	CreateTime           *int64                           `json:"createTime,omitempty" xml:"createTime,omitempty"`
	Description          *string                          `json:"description,omitempty" xml:"description,omitempty"`
	IncidentEscalationId *string                          `json:"incidentEscalationId,omitempty" xml:"incidentEscalationId,omitempty"`
	ModifyTime           *int64                           `json:"modifyTime,omitempty" xml:"modifyTime,omitempty"`
	Name                 *string                          `json:"name,omitempty" xml:"name,omitempty"`
	RegionId             *string                          `json:"regionId,omitempty" xml:"regionId,omitempty"`
	Stage                []*IncidentEscalationStageStruct `json:"stage,omitempty" xml:"stage,omitempty" type:"Repeated"`
	Workspace            *string                          `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s IncidentEscalationStruct) String() string {
	return tea.Prettify(s)
}

func (s IncidentEscalationStruct) GoString() string {
	return s.String()
}

func (s *IncidentEscalationStruct) SetCreateTime(v int64) *IncidentEscalationStruct {
	s.CreateTime = &v
	return s
}

func (s *IncidentEscalationStruct) SetDescription(v string) *IncidentEscalationStruct {
	s.Description = &v
	return s
}

func (s *IncidentEscalationStruct) SetIncidentEscalationId(v string) *IncidentEscalationStruct {
	s.IncidentEscalationId = &v
	return s
}

func (s *IncidentEscalationStruct) SetModifyTime(v int64) *IncidentEscalationStruct {
	s.ModifyTime = &v
	return s
}

func (s *IncidentEscalationStruct) SetName(v string) *IncidentEscalationStruct {
	s.Name = &v
	return s
}

func (s *IncidentEscalationStruct) SetRegionId(v string) *IncidentEscalationStruct {
	s.RegionId = &v
	return s
}

func (s *IncidentEscalationStruct) SetStage(v []*IncidentEscalationStageStruct) *IncidentEscalationStruct {
	s.Stage = v
	return s
}

func (s *IncidentEscalationStruct) SetWorkspace(v string) *IncidentEscalationStruct {
	s.Workspace = &v
	return s
}

type IncidentEventStruct struct {
	AutoRecoverTime *int64             `json:"autoRecoverTime,omitempty" xml:"autoRecoverTime,omitempty"`
	Content         *string            `json:"content,omitempty" xml:"content,omitempty"`
	Count           *int32             `json:"count,omitempty" xml:"count,omitempty"`
	Dimension       map[string]*string `json:"dimension,omitempty" xml:"dimension,omitempty"`
	GroupBy         map[string]*string `json:"groupBy,omitempty" xml:"groupBy,omitempty"`
	IncidentEventId *string            `json:"incidentEventId,omitempty" xml:"incidentEventId,omitempty"`
	IncidentId      *string            `json:"incidentId,omitempty" xml:"incidentId,omitempty"`
	LastTime        *int64             `json:"lastTime,omitempty" xml:"lastTime,omitempty"`
	RecoverTime     *int64             `json:"recoverTime,omitempty" xml:"recoverTime,omitempty"`
	Resource        map[string]*string `json:"resource,omitempty" xml:"resource,omitempty"`
	Status          *int64             `json:"status,omitempty" xml:"status,omitempty"`
	Time            *string            `json:"time,omitempty" xml:"time,omitempty"`
	Title           *string            `json:"title,omitempty" xml:"title,omitempty"`
	UserId          *string            `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s IncidentEventStruct) String() string {
	return tea.Prettify(s)
}

func (s IncidentEventStruct) GoString() string {
	return s.String()
}

func (s *IncidentEventStruct) SetAutoRecoverTime(v int64) *IncidentEventStruct {
	s.AutoRecoverTime = &v
	return s
}

func (s *IncidentEventStruct) SetContent(v string) *IncidentEventStruct {
	s.Content = &v
	return s
}

func (s *IncidentEventStruct) SetCount(v int32) *IncidentEventStruct {
	s.Count = &v
	return s
}

func (s *IncidentEventStruct) SetDimension(v map[string]*string) *IncidentEventStruct {
	s.Dimension = v
	return s
}

func (s *IncidentEventStruct) SetGroupBy(v map[string]*string) *IncidentEventStruct {
	s.GroupBy = v
	return s
}

func (s *IncidentEventStruct) SetIncidentEventId(v string) *IncidentEventStruct {
	s.IncidentEventId = &v
	return s
}

func (s *IncidentEventStruct) SetIncidentId(v string) *IncidentEventStruct {
	s.IncidentId = &v
	return s
}

func (s *IncidentEventStruct) SetLastTime(v int64) *IncidentEventStruct {
	s.LastTime = &v
	return s
}

func (s *IncidentEventStruct) SetRecoverTime(v int64) *IncidentEventStruct {
	s.RecoverTime = &v
	return s
}

func (s *IncidentEventStruct) SetResource(v map[string]*string) *IncidentEventStruct {
	s.Resource = v
	return s
}

func (s *IncidentEventStruct) SetStatus(v int64) *IncidentEventStruct {
	s.Status = &v
	return s
}

func (s *IncidentEventStruct) SetTime(v string) *IncidentEventStruct {
	s.Time = &v
	return s
}

func (s *IncidentEventStruct) SetTitle(v string) *IncidentEventStruct {
	s.Title = &v
	return s
}

func (s *IncidentEventStruct) SetUserId(v string) *IncidentEventStruct {
	s.UserId = &v
	return s
}

type IncidentMemberStruct struct {
	Acknowledge      *IncidentMemberStructAcknowledge   `json:"acknowledge,omitempty" xml:"acknowledge,omitempty" type:"Struct"`
	ContactId        *string                            `json:"contactId,omitempty" xml:"contactId,omitempty"`
	Contacts         []*IncidentMemberStructContacts    `json:"contacts,omitempty" xml:"contacts,omitempty" type:"Repeated"`
	Escalation       *IncidentMemberStructEscalation    `json:"escalation,omitempty" xml:"escalation,omitempty" type:"Struct"`
	IncidentId       *string                            `json:"incidentId,omitempty" xml:"incidentId,omitempty"`
	IncidentMemberId *string                            `json:"incidentMemberId,omitempty" xml:"incidentMemberId,omitempty"`
	ScheduleGroup    *IncidentMemberStructScheduleGroup `json:"scheduleGroup,omitempty" xml:"scheduleGroup,omitempty" type:"Struct"`
	Time             *int64                             `json:"time,omitempty" xml:"time,omitempty"`
	UserId           *int64                             `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s IncidentMemberStruct) String() string {
	return tea.Prettify(s)
}

func (s IncidentMemberStruct) GoString() string {
	return s.String()
}

func (s *IncidentMemberStruct) SetAcknowledge(v *IncidentMemberStructAcknowledge) *IncidentMemberStruct {
	s.Acknowledge = v
	return s
}

func (s *IncidentMemberStruct) SetContactId(v string) *IncidentMemberStruct {
	s.ContactId = &v
	return s
}

func (s *IncidentMemberStruct) SetContacts(v []*IncidentMemberStructContacts) *IncidentMemberStruct {
	s.Contacts = v
	return s
}

func (s *IncidentMemberStruct) SetEscalation(v *IncidentMemberStructEscalation) *IncidentMemberStruct {
	s.Escalation = v
	return s
}

func (s *IncidentMemberStruct) SetIncidentId(v string) *IncidentMemberStruct {
	s.IncidentId = &v
	return s
}

func (s *IncidentMemberStruct) SetIncidentMemberId(v string) *IncidentMemberStruct {
	s.IncidentMemberId = &v
	return s
}

func (s *IncidentMemberStruct) SetScheduleGroup(v *IncidentMemberStructScheduleGroup) *IncidentMemberStruct {
	s.ScheduleGroup = v
	return s
}

func (s *IncidentMemberStruct) SetTime(v int64) *IncidentMemberStruct {
	s.Time = &v
	return s
}

func (s *IncidentMemberStruct) SetUserId(v int64) *IncidentMemberStruct {
	s.UserId = &v
	return s
}

type IncidentMemberStructAcknowledge struct {
	BreakLevel *string `json:"breakLevel,omitempty" xml:"breakLevel,omitempty"`
	VerifyTime *int64  `json:"verifyTime,omitempty" xml:"verifyTime,omitempty"`
}

func (s IncidentMemberStructAcknowledge) String() string {
	return tea.Prettify(s)
}

func (s IncidentMemberStructAcknowledge) GoString() string {
	return s.String()
}

func (s *IncidentMemberStructAcknowledge) SetBreakLevel(v string) *IncidentMemberStructAcknowledge {
	s.BreakLevel = &v
	return s
}

func (s *IncidentMemberStructAcknowledge) SetVerifyTime(v int64) *IncidentMemberStructAcknowledge {
	s.VerifyTime = &v
	return s
}

type IncidentMemberStructContacts struct {
	Channel     *string `json:"channel,omitempty" xml:"channel,omitempty"`
	ContactMask *string `json:"contactMask,omitempty" xml:"contactMask,omitempty"`
}

func (s IncidentMemberStructContacts) String() string {
	return tea.Prettify(s)
}

func (s IncidentMemberStructContacts) GoString() string {
	return s.String()
}

func (s *IncidentMemberStructContacts) SetChannel(v string) *IncidentMemberStructContacts {
	s.Channel = &v
	return s
}

func (s *IncidentMemberStructContacts) SetContactMask(v string) *IncidentMemberStructContacts {
	s.ContactMask = &v
	return s
}

type IncidentMemberStructEscalation struct {
	Description          *string `json:"description,omitempty" xml:"description,omitempty"`
	IncidentEscalationId *string `json:"incidentEscalationId,omitempty" xml:"incidentEscalationId,omitempty"`
	Name                 *string `json:"name,omitempty" xml:"name,omitempty"`
	StageIndex           *string `json:"stageIndex,omitempty" xml:"stageIndex,omitempty"`
	Title                *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s IncidentMemberStructEscalation) String() string {
	return tea.Prettify(s)
}

func (s IncidentMemberStructEscalation) GoString() string {
	return s.String()
}

func (s *IncidentMemberStructEscalation) SetDescription(v string) *IncidentMemberStructEscalation {
	s.Description = &v
	return s
}

func (s *IncidentMemberStructEscalation) SetIncidentEscalationId(v string) *IncidentMemberStructEscalation {
	s.IncidentEscalationId = &v
	return s
}

func (s *IncidentMemberStructEscalation) SetName(v string) *IncidentMemberStructEscalation {
	s.Name = &v
	return s
}

func (s *IncidentMemberStructEscalation) SetStageIndex(v string) *IncidentMemberStructEscalation {
	s.StageIndex = &v
	return s
}

func (s *IncidentMemberStructEscalation) SetTitle(v string) *IncidentMemberStructEscalation {
	s.Title = &v
	return s
}

type IncidentMemberStructScheduleGroup struct {
	ContactId *string `json:"contactId,omitempty" xml:"contactId,omitempty"`
	Name      *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s IncidentMemberStructScheduleGroup) String() string {
	return tea.Prettify(s)
}

func (s IncidentMemberStructScheduleGroup) GoString() string {
	return s.String()
}

func (s *IncidentMemberStructScheduleGroup) SetContactId(v string) *IncidentMemberStructScheduleGroup {
	s.ContactId = &v
	return s
}

func (s *IncidentMemberStructScheduleGroup) SetName(v string) *IncidentMemberStructScheduleGroup {
	s.Name = &v
	return s
}

type IncidentNoteStruct struct {
	Content    *string                     `json:"content,omitempty" xml:"content,omitempty"`
	Format     *string                     `json:"format,omitempty" xml:"format,omitempty"`
	IncidentId *string                     `json:"incidentId,omitempty" xml:"incidentId,omitempty"`
	NoteId     *string                     `json:"noteId,omitempty" xml:"noteId,omitempty"`
	Operator   *IncidentNoteStructOperator `json:"operator,omitempty" xml:"operator,omitempty" type:"Struct"`
	Time       *int64                      `json:"time,omitempty" xml:"time,omitempty"`
	Type       *string                     `json:"type,omitempty" xml:"type,omitempty"`
}

func (s IncidentNoteStruct) String() string {
	return tea.Prettify(s)
}

func (s IncidentNoteStruct) GoString() string {
	return s.String()
}

func (s *IncidentNoteStruct) SetContent(v string) *IncidentNoteStruct {
	s.Content = &v
	return s
}

func (s *IncidentNoteStruct) SetFormat(v string) *IncidentNoteStruct {
	s.Format = &v
	return s
}

func (s *IncidentNoteStruct) SetIncidentId(v string) *IncidentNoteStruct {
	s.IncidentId = &v
	return s
}

func (s *IncidentNoteStruct) SetNoteId(v string) *IncidentNoteStruct {
	s.NoteId = &v
	return s
}

func (s *IncidentNoteStruct) SetOperator(v *IncidentNoteStructOperator) *IncidentNoteStruct {
	s.Operator = v
	return s
}

func (s *IncidentNoteStruct) SetTime(v int64) *IncidentNoteStruct {
	s.Time = &v
	return s
}

func (s *IncidentNoteStruct) SetType(v string) *IncidentNoteStruct {
	s.Type = &v
	return s
}

type IncidentNoteStructOperator struct {
	Contact   *string `json:"contact,omitempty" xml:"contact,omitempty"`
	ContactId *string `json:"contactId,omitempty" xml:"contactId,omitempty"`
	Name      *string `json:"name,omitempty" xml:"name,omitempty"`
	UserId    *int64  `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s IncidentNoteStructOperator) String() string {
	return tea.Prettify(s)
}

func (s IncidentNoteStructOperator) GoString() string {
	return s.String()
}

func (s *IncidentNoteStructOperator) SetContact(v string) *IncidentNoteStructOperator {
	s.Contact = &v
	return s
}

func (s *IncidentNoteStructOperator) SetContactId(v string) *IncidentNoteStructOperator {
	s.ContactId = &v
	return s
}

func (s *IncidentNoteStructOperator) SetName(v string) *IncidentNoteStructOperator {
	s.Name = &v
	return s
}

func (s *IncidentNoteStructOperator) SetUserId(v int64) *IncidentNoteStructOperator {
	s.UserId = &v
	return s
}

type IncidentPlanCorporationStruct struct {
	Channel *string `json:"channel,omitempty" xml:"channel,omitempty"`
	RobotId *string `json:"robotId,omitempty" xml:"robotId,omitempty"`
}

func (s IncidentPlanCorporationStruct) String() string {
	return tea.Prettify(s)
}

func (s IncidentPlanCorporationStruct) GoString() string {
	return s.String()
}

func (s *IncidentPlanCorporationStruct) SetChannel(v string) *IncidentPlanCorporationStruct {
	s.Channel = &v
	return s
}

func (s *IncidentPlanCorporationStruct) SetRobotId(v string) *IncidentPlanCorporationStruct {
	s.RobotId = &v
	return s
}

type IncidentPlanFieldPath struct {
	FieldAlias *string   `json:"fieldAlias,omitempty" xml:"fieldAlias,omitempty"`
	FieldPath  []*string `json:"fieldPath,omitempty" xml:"fieldPath,omitempty" type:"Repeated"`
}

func (s IncidentPlanFieldPath) String() string {
	return tea.Prettify(s)
}

func (s IncidentPlanFieldPath) GoString() string {
	return s.String()
}

func (s *IncidentPlanFieldPath) SetFieldAlias(v string) *IncidentPlanFieldPath {
	s.FieldAlias = &v
	return s
}

func (s *IncidentPlanFieldPath) SetFieldPath(v []*string) *IncidentPlanFieldPath {
	s.FieldPath = v
	return s
}

type IncidentPlanStruct struct {
	AutoRecoverSeconds *int32                           `json:"autoRecoverSeconds,omitempty" xml:"autoRecoverSeconds,omitempty"`
	CloseExpire        *int64                           `json:"closeExpire,omitempty" xml:"closeExpire,omitempty"`
	Corporation        []*IncidentPlanCorporationStruct `json:"corporation,omitempty" xml:"corporation,omitempty" type:"Repeated"`
	Description        *string                          `json:"description,omitempty" xml:"description,omitempty"`
	EscalationId       []*string                        `json:"escalationId,omitempty" xml:"escalationId,omitempty" type:"Repeated"`
	GmtCreate          *int64                           `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	GmtModified        *int64                           `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	GroupBy            []*IncidentPlanFieldPath         `json:"groupBy,omitempty" xml:"groupBy,omitempty" type:"Repeated"`
	IncidentPlanId     *string                          `json:"incidentPlanId,omitempty" xml:"incidentPlanId,omitempty"`
	Name               *string                          `json:"name,omitempty" xml:"name,omitempty"`
	ResourceFiled      []*IncidentPlanFieldPath         `json:"resourceFiled,omitempty" xml:"resourceFiled,omitempty" type:"Repeated"`
	Status             *string                          `json:"status,omitempty" xml:"status,omitempty"`
	UserId             *int64                           `json:"userId,omitempty" xml:"userId,omitempty"`
	Workspace          *string                          `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s IncidentPlanStruct) String() string {
	return tea.Prettify(s)
}

func (s IncidentPlanStruct) GoString() string {
	return s.String()
}

func (s *IncidentPlanStruct) SetAutoRecoverSeconds(v int32) *IncidentPlanStruct {
	s.AutoRecoverSeconds = &v
	return s
}

func (s *IncidentPlanStruct) SetCloseExpire(v int64) *IncidentPlanStruct {
	s.CloseExpire = &v
	return s
}

func (s *IncidentPlanStruct) SetCorporation(v []*IncidentPlanCorporationStruct) *IncidentPlanStruct {
	s.Corporation = v
	return s
}

func (s *IncidentPlanStruct) SetDescription(v string) *IncidentPlanStruct {
	s.Description = &v
	return s
}

func (s *IncidentPlanStruct) SetEscalationId(v []*string) *IncidentPlanStruct {
	s.EscalationId = v
	return s
}

func (s *IncidentPlanStruct) SetGmtCreate(v int64) *IncidentPlanStruct {
	s.GmtCreate = &v
	return s
}

func (s *IncidentPlanStruct) SetGmtModified(v int64) *IncidentPlanStruct {
	s.GmtModified = &v
	return s
}

func (s *IncidentPlanStruct) SetGroupBy(v []*IncidentPlanFieldPath) *IncidentPlanStruct {
	s.GroupBy = v
	return s
}

func (s *IncidentPlanStruct) SetIncidentPlanId(v string) *IncidentPlanStruct {
	s.IncidentPlanId = &v
	return s
}

func (s *IncidentPlanStruct) SetName(v string) *IncidentPlanStruct {
	s.Name = &v
	return s
}

func (s *IncidentPlanStruct) SetResourceFiled(v []*IncidentPlanFieldPath) *IncidentPlanStruct {
	s.ResourceFiled = v
	return s
}

func (s *IncidentPlanStruct) SetStatus(v string) *IncidentPlanStruct {
	s.Status = &v
	return s
}

func (s *IncidentPlanStruct) SetUserId(v int64) *IncidentPlanStruct {
	s.UserId = &v
	return s
}

func (s *IncidentPlanStruct) SetWorkspace(v string) *IncidentPlanStruct {
	s.Workspace = &v
	return s
}

type IncidentResourceDetail struct {
	ExtraId    *string                `json:"extraId,omitempty" xml:"extraId,omitempty"`
	ResourceId map[string]interface{} `json:"resourceId,omitempty" xml:"resourceId,omitempty"`
	Type       *string                `json:"type,omitempty" xml:"type,omitempty"`
}

func (s IncidentResourceDetail) String() string {
	return tea.Prettify(s)
}

func (s IncidentResourceDetail) GoString() string {
	return s.String()
}

func (s *IncidentResourceDetail) SetExtraId(v string) *IncidentResourceDetail {
	s.ExtraId = &v
	return s
}

func (s *IncidentResourceDetail) SetResourceId(v map[string]interface{}) *IncidentResourceDetail {
	s.ResourceId = v
	return s
}

func (s *IncidentResourceDetail) SetType(v string) *IncidentResourceDetail {
	s.Type = &v
	return s
}

type IncidentResourceStruct struct {
	Description        *string                 `json:"description,omitempty" xml:"description,omitempty"`
	IncidentId         *string                 `json:"incidentId,omitempty" xml:"incidentId,omitempty"`
	IncidentResourceId *string                 `json:"incidentResourceId,omitempty" xml:"incidentResourceId,omitempty"`
	Resource           *IncidentResourceDetail `json:"resource,omitempty" xml:"resource,omitempty"`
	Source             *string                 `json:"source,omitempty" xml:"source,omitempty"`
	Time               *int64                  `json:"time,omitempty" xml:"time,omitempty"`
	UserId             *int64                  `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s IncidentResourceStruct) String() string {
	return tea.Prettify(s)
}

func (s IncidentResourceStruct) GoString() string {
	return s.String()
}

func (s *IncidentResourceStruct) SetDescription(v string) *IncidentResourceStruct {
	s.Description = &v
	return s
}

func (s *IncidentResourceStruct) SetIncidentId(v string) *IncidentResourceStruct {
	s.IncidentId = &v
	return s
}

func (s *IncidentResourceStruct) SetIncidentResourceId(v string) *IncidentResourceStruct {
	s.IncidentResourceId = &v
	return s
}

func (s *IncidentResourceStruct) SetResource(v *IncidentResourceDetail) *IncidentResourceStruct {
	s.Resource = v
	return s
}

func (s *IncidentResourceStruct) SetSource(v string) *IncidentResourceStruct {
	s.Source = &v
	return s
}

func (s *IncidentResourceStruct) SetTime(v int64) *IncidentResourceStruct {
	s.Time = &v
	return s
}

func (s *IncidentResourceStruct) SetUserId(v int64) *IncidentResourceStruct {
	s.UserId = &v
	return s
}

type IncidentStruct struct {
	Content      *string                     `json:"content,omitempty" xml:"content,omitempty"`
	Escalations  []*IncidentEscalationStruct `json:"escalations,omitempty" xml:"escalations,omitempty" type:"Repeated"`
	IncidentId   *string                     `json:"incidentId,omitempty" xml:"incidentId,omitempty"`
	IncidentPlan *IncidentPlanStruct         `json:"incidentPlan,omitempty" xml:"incidentPlan,omitempty"`
	Resource     *IncidentResourceDetail     `json:"resource,omitempty" xml:"resource,omitempty"`
	Severity     *string                     `json:"severity,omitempty" xml:"severity,omitempty"`
	Status       *string                     `json:"status,omitempty" xml:"status,omitempty"`
	Time         *int64                      `json:"time,omitempty" xml:"time,omitempty"`
	Title        *string                     `json:"title,omitempty" xml:"title,omitempty"`
	UserId       *string                     `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s IncidentStruct) String() string {
	return tea.Prettify(s)
}

func (s IncidentStruct) GoString() string {
	return s.String()
}

func (s *IncidentStruct) SetContent(v string) *IncidentStruct {
	s.Content = &v
	return s
}

func (s *IncidentStruct) SetEscalations(v []*IncidentEscalationStruct) *IncidentStruct {
	s.Escalations = v
	return s
}

func (s *IncidentStruct) SetIncidentId(v string) *IncidentStruct {
	s.IncidentId = &v
	return s
}

func (s *IncidentStruct) SetIncidentPlan(v *IncidentPlanStruct) *IncidentStruct {
	s.IncidentPlan = v
	return s
}

func (s *IncidentStruct) SetResource(v *IncidentResourceDetail) *IncidentStruct {
	s.Resource = v
	return s
}

func (s *IncidentStruct) SetSeverity(v string) *IncidentStruct {
	s.Severity = &v
	return s
}

func (s *IncidentStruct) SetStatus(v string) *IncidentStruct {
	s.Status = &v
	return s
}

func (s *IncidentStruct) SetTime(v int64) *IncidentStruct {
	s.Time = &v
	return s
}

func (s *IncidentStruct) SetTitle(v string) *IncidentStruct {
	s.Title = &v
	return s
}

func (s *IncidentStruct) SetUserId(v string) *IncidentStruct {
	s.UserId = &v
	return s
}

type IncidentTimeline struct {
	ChildType          *string `json:"childType,omitempty" xml:"childType,omitempty"`
	Content            *string `json:"content,omitempty" xml:"content,omitempty"`
	IncidentId         *string `json:"incidentId,omitempty" xml:"incidentId,omitempty"`
	IncidentTimelineId *string `json:"incidentTimelineId,omitempty" xml:"incidentTimelineId,omitempty"`
	Time               *int64  `json:"time,omitempty" xml:"time,omitempty"`
	TimelineId         *string `json:"timelineId,omitempty" xml:"timelineId,omitempty"`
	Title              *string `json:"title,omitempty" xml:"title,omitempty"`
	Type               *string `json:"type,omitempty" xml:"type,omitempty"`
	UserId             *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s IncidentTimeline) String() string {
	return tea.Prettify(s)
}

func (s IncidentTimeline) GoString() string {
	return s.String()
}

func (s *IncidentTimeline) SetChildType(v string) *IncidentTimeline {
	s.ChildType = &v
	return s
}

func (s *IncidentTimeline) SetContent(v string) *IncidentTimeline {
	s.Content = &v
	return s
}

func (s *IncidentTimeline) SetIncidentId(v string) *IncidentTimeline {
	s.IncidentId = &v
	return s
}

func (s *IncidentTimeline) SetIncidentTimelineId(v string) *IncidentTimeline {
	s.IncidentTimelineId = &v
	return s
}

func (s *IncidentTimeline) SetTime(v int64) *IncidentTimeline {
	s.Time = &v
	return s
}

func (s *IncidentTimeline) SetTimelineId(v string) *IncidentTimeline {
	s.TimelineId = &v
	return s
}

func (s *IncidentTimeline) SetTitle(v string) *IncidentTimeline {
	s.Title = &v
	return s
}

func (s *IncidentTimeline) SetType(v string) *IncidentTimeline {
	s.Type = &v
	return s
}

func (s *IncidentTimeline) SetUserId(v string) *IncidentTimeline {
	s.UserId = &v
	return s
}

type MaintainWindowForModify struct {
	Description     *string                                 `json:"description,omitempty" xml:"description,omitempty"`
	EffectTimeRange *MaintainWindowForModifyEffectTimeRange `json:"effectTimeRange,omitempty" xml:"effectTimeRange,omitempty" type:"Struct"`
	Effective       *string                                 `json:"effective,omitempty" xml:"effective,omitempty"`
	EndTime         *string                                 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	FilterSetting   *FilterSetting                          `json:"filterSetting,omitempty" xml:"filterSetting,omitempty"`
	// This parameter is required.
	MaintainWindowName *string `json:"maintainWindowName,omitempty" xml:"maintainWindowName,omitempty"`
	StartTime          *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
}

func (s MaintainWindowForModify) String() string {
	return tea.Prettify(s)
}

func (s MaintainWindowForModify) GoString() string {
	return s.String()
}

func (s *MaintainWindowForModify) SetDescription(v string) *MaintainWindowForModify {
	s.Description = &v
	return s
}

func (s *MaintainWindowForModify) SetEffectTimeRange(v *MaintainWindowForModifyEffectTimeRange) *MaintainWindowForModify {
	s.EffectTimeRange = v
	return s
}

func (s *MaintainWindowForModify) SetEffective(v string) *MaintainWindowForModify {
	s.Effective = &v
	return s
}

func (s *MaintainWindowForModify) SetEndTime(v string) *MaintainWindowForModify {
	s.EndTime = &v
	return s
}

func (s *MaintainWindowForModify) SetFilterSetting(v *FilterSetting) *MaintainWindowForModify {
	s.FilterSetting = v
	return s
}

func (s *MaintainWindowForModify) SetMaintainWindowName(v string) *MaintainWindowForModify {
	s.MaintainWindowName = &v
	return s
}

func (s *MaintainWindowForModify) SetStartTime(v string) *MaintainWindowForModify {
	s.StartTime = &v
	return s
}

type MaintainWindowForModifyEffectTimeRange struct {
	DayInWeek         []*int32 `json:"dayInWeek,omitempty" xml:"dayInWeek,omitempty" type:"Repeated"`
	EndTimeInMinute   *int32   `json:"endTimeInMinute,omitempty" xml:"endTimeInMinute,omitempty"`
	StartTimeInMinute *int32   `json:"startTimeInMinute,omitempty" xml:"startTimeInMinute,omitempty"`
	TimeZone          *string  `json:"timeZone,omitempty" xml:"timeZone,omitempty"`
}

func (s MaintainWindowForModifyEffectTimeRange) String() string {
	return tea.Prettify(s)
}

func (s MaintainWindowForModifyEffectTimeRange) GoString() string {
	return s.String()
}

func (s *MaintainWindowForModifyEffectTimeRange) SetDayInWeek(v []*int32) *MaintainWindowForModifyEffectTimeRange {
	s.DayInWeek = v
	return s
}

func (s *MaintainWindowForModifyEffectTimeRange) SetEndTimeInMinute(v int32) *MaintainWindowForModifyEffectTimeRange {
	s.EndTimeInMinute = &v
	return s
}

func (s *MaintainWindowForModifyEffectTimeRange) SetStartTimeInMinute(v int32) *MaintainWindowForModifyEffectTimeRange {
	s.StartTimeInMinute = &v
	return s
}

func (s *MaintainWindowForModifyEffectTimeRange) SetTimeZone(v string) *MaintainWindowForModifyEffectTimeRange {
	s.TimeZone = &v
	return s
}

type MaintainWindowForView struct {
	CreateTime       *string                               `json:"createTime,omitempty" xml:"createTime,omitempty"`
	Description      *string                               `json:"description,omitempty" xml:"description,omitempty"`
	EffectTimeRange  *MaintainWindowForViewEffectTimeRange `json:"effectTimeRange,omitempty" xml:"effectTimeRange,omitempty" type:"Struct"`
	Effective        *string                               `json:"effective,omitempty" xml:"effective,omitempty"`
	Enable           *bool                                 `json:"enable,omitempty" xml:"enable,omitempty"`
	EndTime          *string                               `json:"endTime,omitempty" xml:"endTime,omitempty"`
	FilterSetting    *FilterSetting                        `json:"filterSetting,omitempty" xml:"filterSetting,omitempty"`
	MaintainWindowId *string                               `json:"maintainWindowId,omitempty" xml:"maintainWindowId,omitempty"`
	// This parameter is required.
	MaintainWindowName *string `json:"maintainWindowName,omitempty" xml:"maintainWindowName,omitempty"`
	StartTime          *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
	UpdateTime         *string `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
	UserId             *string `json:"userId,omitempty" xml:"userId,omitempty"`
	Workspace          *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s MaintainWindowForView) String() string {
	return tea.Prettify(s)
}

func (s MaintainWindowForView) GoString() string {
	return s.String()
}

func (s *MaintainWindowForView) SetCreateTime(v string) *MaintainWindowForView {
	s.CreateTime = &v
	return s
}

func (s *MaintainWindowForView) SetDescription(v string) *MaintainWindowForView {
	s.Description = &v
	return s
}

func (s *MaintainWindowForView) SetEffectTimeRange(v *MaintainWindowForViewEffectTimeRange) *MaintainWindowForView {
	s.EffectTimeRange = v
	return s
}

func (s *MaintainWindowForView) SetEffective(v string) *MaintainWindowForView {
	s.Effective = &v
	return s
}

func (s *MaintainWindowForView) SetEnable(v bool) *MaintainWindowForView {
	s.Enable = &v
	return s
}

func (s *MaintainWindowForView) SetEndTime(v string) *MaintainWindowForView {
	s.EndTime = &v
	return s
}

func (s *MaintainWindowForView) SetFilterSetting(v *FilterSetting) *MaintainWindowForView {
	s.FilterSetting = v
	return s
}

func (s *MaintainWindowForView) SetMaintainWindowId(v string) *MaintainWindowForView {
	s.MaintainWindowId = &v
	return s
}

func (s *MaintainWindowForView) SetMaintainWindowName(v string) *MaintainWindowForView {
	s.MaintainWindowName = &v
	return s
}

func (s *MaintainWindowForView) SetStartTime(v string) *MaintainWindowForView {
	s.StartTime = &v
	return s
}

func (s *MaintainWindowForView) SetUpdateTime(v string) *MaintainWindowForView {
	s.UpdateTime = &v
	return s
}

func (s *MaintainWindowForView) SetUserId(v string) *MaintainWindowForView {
	s.UserId = &v
	return s
}

func (s *MaintainWindowForView) SetWorkspace(v string) *MaintainWindowForView {
	s.Workspace = &v
	return s
}

type MaintainWindowForViewEffectTimeRange struct {
	DayInWeek         []*int32 `json:"dayInWeek,omitempty" xml:"dayInWeek,omitempty" type:"Repeated"`
	EndTimeInMinute   *int32   `json:"endTimeInMinute,omitempty" xml:"endTimeInMinute,omitempty"`
	StartTimeInMinute *int32   `json:"startTimeInMinute,omitempty" xml:"startTimeInMinute,omitempty"`
	TimeZone          *string  `json:"timeZone,omitempty" xml:"timeZone,omitempty"`
}

func (s MaintainWindowForViewEffectTimeRange) String() string {
	return tea.Prettify(s)
}

func (s MaintainWindowForViewEffectTimeRange) GoString() string {
	return s.String()
}

func (s *MaintainWindowForViewEffectTimeRange) SetDayInWeek(v []*int32) *MaintainWindowForViewEffectTimeRange {
	s.DayInWeek = v
	return s
}

func (s *MaintainWindowForViewEffectTimeRange) SetEndTimeInMinute(v int32) *MaintainWindowForViewEffectTimeRange {
	s.EndTimeInMinute = &v
	return s
}

func (s *MaintainWindowForViewEffectTimeRange) SetStartTimeInMinute(v int32) *MaintainWindowForViewEffectTimeRange {
	s.StartTimeInMinute = &v
	return s
}

func (s *MaintainWindowForViewEffectTimeRange) SetTimeZone(v string) *MaintainWindowForViewEffectTimeRange {
	s.TimeZone = &v
	return s
}

type NotifyStrategyForModify struct {
	CustomTemplateEntries []*NotifyStrategyForModifyCustomTemplateEntries `json:"customTemplateEntries,omitempty" xml:"customTemplateEntries,omitempty" type:"Repeated"`
	Description           *string                                         `json:"description,omitempty" xml:"description,omitempty"`
	// This parameter is required.
	GroupingSetting            *NotifyStrategyForModifyGroupingSetting `json:"groupingSetting,omitempty" xml:"groupingSetting,omitempty" type:"Struct"`
	IgnoreRestoredNotification *bool                                   `json:"ignoreRestoredNotification,omitempty" xml:"ignoreRestoredNotification,omitempty"`
	// This parameter is required.
	NotifyStrategyName *string `json:"notifyStrategyName,omitempty" xml:"notifyStrategyName,omitempty"`
	// This parameter is required.
	Routes []*NotifyStrategyForModifyRoutes `json:"routes,omitempty" xml:"routes,omitempty" type:"Repeated"`
}

func (s NotifyStrategyForModify) String() string {
	return tea.Prettify(s)
}

func (s NotifyStrategyForModify) GoString() string {
	return s.String()
}

func (s *NotifyStrategyForModify) SetCustomTemplateEntries(v []*NotifyStrategyForModifyCustomTemplateEntries) *NotifyStrategyForModify {
	s.CustomTemplateEntries = v
	return s
}

func (s *NotifyStrategyForModify) SetDescription(v string) *NotifyStrategyForModify {
	s.Description = &v
	return s
}

func (s *NotifyStrategyForModify) SetGroupingSetting(v *NotifyStrategyForModifyGroupingSetting) *NotifyStrategyForModify {
	s.GroupingSetting = v
	return s
}

func (s *NotifyStrategyForModify) SetIgnoreRestoredNotification(v bool) *NotifyStrategyForModify {
	s.IgnoreRestoredNotification = &v
	return s
}

func (s *NotifyStrategyForModify) SetNotifyStrategyName(v string) *NotifyStrategyForModify {
	s.NotifyStrategyName = &v
	return s
}

func (s *NotifyStrategyForModify) SetRoutes(v []*NotifyStrategyForModifyRoutes) *NotifyStrategyForModify {
	s.Routes = v
	return s
}

type NotifyStrategyForModifyCustomTemplateEntries struct {
	// This parameter is required.
	TargetType *string `json:"targetType,omitempty" xml:"targetType,omitempty"`
	// This parameter is required.
	TemplateUuid *string `json:"templateUuid,omitempty" xml:"templateUuid,omitempty"`
}

func (s NotifyStrategyForModifyCustomTemplateEntries) String() string {
	return tea.Prettify(s)
}

func (s NotifyStrategyForModifyCustomTemplateEntries) GoString() string {
	return s.String()
}

func (s *NotifyStrategyForModifyCustomTemplateEntries) SetTargetType(v string) *NotifyStrategyForModifyCustomTemplateEntries {
	s.TargetType = &v
	return s
}

func (s *NotifyStrategyForModifyCustomTemplateEntries) SetTemplateUuid(v string) *NotifyStrategyForModifyCustomTemplateEntries {
	s.TemplateUuid = &v
	return s
}

type NotifyStrategyForModifyGroupingSetting struct {
	GroupingKeys []*string `json:"groupingKeys,omitempty" xml:"groupingKeys,omitempty" type:"Repeated"`
	PeriodMin    *int32    `json:"periodMin,omitempty" xml:"periodMin,omitempty"`
	SilenceSec   *int32    `json:"silenceSec,omitempty" xml:"silenceSec,omitempty"`
	Times        *int32    `json:"times,omitempty" xml:"times,omitempty"`
}

func (s NotifyStrategyForModifyGroupingSetting) String() string {
	return tea.Prettify(s)
}

func (s NotifyStrategyForModifyGroupingSetting) GoString() string {
	return s.String()
}

func (s *NotifyStrategyForModifyGroupingSetting) SetGroupingKeys(v []*string) *NotifyStrategyForModifyGroupingSetting {
	s.GroupingKeys = v
	return s
}

func (s *NotifyStrategyForModifyGroupingSetting) SetPeriodMin(v int32) *NotifyStrategyForModifyGroupingSetting {
	s.PeriodMin = &v
	return s
}

func (s *NotifyStrategyForModifyGroupingSetting) SetSilenceSec(v int32) *NotifyStrategyForModifyGroupingSetting {
	s.SilenceSec = &v
	return s
}

func (s *NotifyStrategyForModifyGroupingSetting) SetTimes(v int32) *NotifyStrategyForModifyGroupingSetting {
	s.Times = &v
	return s
}

type NotifyStrategyForModifyRoutes struct {
	Channels        []*NotifyStrategyForModifyRoutesChannels      `json:"channels,omitempty" xml:"channels,omitempty" type:"Repeated"`
	EffectTimeRange *NotifyStrategyForModifyRoutesEffectTimeRange `json:"effectTimeRange,omitempty" xml:"effectTimeRange,omitempty" type:"Struct"`
	FilterSetting   *FilterSetting                                `json:"filterSetting,omitempty" xml:"filterSetting,omitempty"`
	Severities      []*string                                     `json:"severities,omitempty" xml:"severities,omitempty" type:"Repeated"`
}

func (s NotifyStrategyForModifyRoutes) String() string {
	return tea.Prettify(s)
}

func (s NotifyStrategyForModifyRoutes) GoString() string {
	return s.String()
}

func (s *NotifyStrategyForModifyRoutes) SetChannels(v []*NotifyStrategyForModifyRoutesChannels) *NotifyStrategyForModifyRoutes {
	s.Channels = v
	return s
}

func (s *NotifyStrategyForModifyRoutes) SetEffectTimeRange(v *NotifyStrategyForModifyRoutesEffectTimeRange) *NotifyStrategyForModifyRoutes {
	s.EffectTimeRange = v
	return s
}

func (s *NotifyStrategyForModifyRoutes) SetFilterSetting(v *FilterSetting) *NotifyStrategyForModifyRoutes {
	s.FilterSetting = v
	return s
}

func (s *NotifyStrategyForModifyRoutes) SetSeverities(v []*string) *NotifyStrategyForModifyRoutes {
	s.Severities = v
	return s
}

type NotifyStrategyForModifyRoutesChannels struct {
	// This parameter is required.
	ChannelType        *string   `json:"channelType,omitempty" xml:"channelType,omitempty"`
	EnabledSubChannels []*string `json:"enabledSubChannels,omitempty" xml:"enabledSubChannels,omitempty" type:"Repeated"`
	// This parameter is required.
	Receivers []*string `json:"receivers,omitempty" xml:"receivers,omitempty" type:"Repeated"`
}

func (s NotifyStrategyForModifyRoutesChannels) String() string {
	return tea.Prettify(s)
}

func (s NotifyStrategyForModifyRoutesChannels) GoString() string {
	return s.String()
}

func (s *NotifyStrategyForModifyRoutesChannels) SetChannelType(v string) *NotifyStrategyForModifyRoutesChannels {
	s.ChannelType = &v
	return s
}

func (s *NotifyStrategyForModifyRoutesChannels) SetEnabledSubChannels(v []*string) *NotifyStrategyForModifyRoutesChannels {
	s.EnabledSubChannels = v
	return s
}

func (s *NotifyStrategyForModifyRoutesChannels) SetReceivers(v []*string) *NotifyStrategyForModifyRoutesChannels {
	s.Receivers = v
	return s
}

type NotifyStrategyForModifyRoutesEffectTimeRange struct {
	DayInWeek         []*int32 `json:"dayInWeek,omitempty" xml:"dayInWeek,omitempty" type:"Repeated"`
	EndTimeInMinute   *int32   `json:"endTimeInMinute,omitempty" xml:"endTimeInMinute,omitempty"`
	StartTimeInMinute *int32   `json:"startTimeInMinute,omitempty" xml:"startTimeInMinute,omitempty"`
	TimeZone          *string  `json:"timeZone,omitempty" xml:"timeZone,omitempty"`
}

func (s NotifyStrategyForModifyRoutesEffectTimeRange) String() string {
	return tea.Prettify(s)
}

func (s NotifyStrategyForModifyRoutesEffectTimeRange) GoString() string {
	return s.String()
}

func (s *NotifyStrategyForModifyRoutesEffectTimeRange) SetDayInWeek(v []*int32) *NotifyStrategyForModifyRoutesEffectTimeRange {
	s.DayInWeek = v
	return s
}

func (s *NotifyStrategyForModifyRoutesEffectTimeRange) SetEndTimeInMinute(v int32) *NotifyStrategyForModifyRoutesEffectTimeRange {
	s.EndTimeInMinute = &v
	return s
}

func (s *NotifyStrategyForModifyRoutesEffectTimeRange) SetStartTimeInMinute(v int32) *NotifyStrategyForModifyRoutesEffectTimeRange {
	s.StartTimeInMinute = &v
	return s
}

func (s *NotifyStrategyForModifyRoutesEffectTimeRange) SetTimeZone(v string) *NotifyStrategyForModifyRoutesEffectTimeRange {
	s.TimeZone = &v
	return s
}

type NotifyStrategyForView struct {
	CreateTime            *string                                       `json:"createTime,omitempty" xml:"createTime,omitempty"`
	CustomTemplateEntries []*NotifyStrategyForViewCustomTemplateEntries `json:"customTemplateEntries,omitempty" xml:"customTemplateEntries,omitempty" type:"Repeated"`
	Description           *string                                       `json:"description,omitempty" xml:"description,omitempty"`
	Enable                *bool                                         `json:"enable,omitempty" xml:"enable,omitempty"`
	// This parameter is required.
	GroupingSetting            *NotifyStrategyForViewGroupingSetting `json:"groupingSetting,omitempty" xml:"groupingSetting,omitempty" type:"Struct"`
	IgnoreRestoredNotification *bool                                 `json:"ignoreRestoredNotification,omitempty" xml:"ignoreRestoredNotification,omitempty"`
	NotifyStrategyId           *string                               `json:"notifyStrategyId,omitempty" xml:"notifyStrategyId,omitempty"`
	// This parameter is required.
	NotifyStrategyName *string `json:"notifyStrategyName,omitempty" xml:"notifyStrategyName,omitempty"`
	// This parameter is required.
	Routes     []*NotifyStrategyForViewRoutes `json:"routes,omitempty" xml:"routes,omitempty" type:"Repeated"`
	UpdateTime *string                        `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
	UserId     *string                        `json:"userId,omitempty" xml:"userId,omitempty"`
	Workspace  *string                        `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s NotifyStrategyForView) String() string {
	return tea.Prettify(s)
}

func (s NotifyStrategyForView) GoString() string {
	return s.String()
}

func (s *NotifyStrategyForView) SetCreateTime(v string) *NotifyStrategyForView {
	s.CreateTime = &v
	return s
}

func (s *NotifyStrategyForView) SetCustomTemplateEntries(v []*NotifyStrategyForViewCustomTemplateEntries) *NotifyStrategyForView {
	s.CustomTemplateEntries = v
	return s
}

func (s *NotifyStrategyForView) SetDescription(v string) *NotifyStrategyForView {
	s.Description = &v
	return s
}

func (s *NotifyStrategyForView) SetEnable(v bool) *NotifyStrategyForView {
	s.Enable = &v
	return s
}

func (s *NotifyStrategyForView) SetGroupingSetting(v *NotifyStrategyForViewGroupingSetting) *NotifyStrategyForView {
	s.GroupingSetting = v
	return s
}

func (s *NotifyStrategyForView) SetIgnoreRestoredNotification(v bool) *NotifyStrategyForView {
	s.IgnoreRestoredNotification = &v
	return s
}

func (s *NotifyStrategyForView) SetNotifyStrategyId(v string) *NotifyStrategyForView {
	s.NotifyStrategyId = &v
	return s
}

func (s *NotifyStrategyForView) SetNotifyStrategyName(v string) *NotifyStrategyForView {
	s.NotifyStrategyName = &v
	return s
}

func (s *NotifyStrategyForView) SetRoutes(v []*NotifyStrategyForViewRoutes) *NotifyStrategyForView {
	s.Routes = v
	return s
}

func (s *NotifyStrategyForView) SetUpdateTime(v string) *NotifyStrategyForView {
	s.UpdateTime = &v
	return s
}

func (s *NotifyStrategyForView) SetUserId(v string) *NotifyStrategyForView {
	s.UserId = &v
	return s
}

func (s *NotifyStrategyForView) SetWorkspace(v string) *NotifyStrategyForView {
	s.Workspace = &v
	return s
}

type NotifyStrategyForViewCustomTemplateEntries struct {
	// This parameter is required.
	TargetType *string `json:"targetType,omitempty" xml:"targetType,omitempty"`
	// This parameter is required.
	TemplateUuid *string `json:"templateUuid,omitempty" xml:"templateUuid,omitempty"`
}

func (s NotifyStrategyForViewCustomTemplateEntries) String() string {
	return tea.Prettify(s)
}

func (s NotifyStrategyForViewCustomTemplateEntries) GoString() string {
	return s.String()
}

func (s *NotifyStrategyForViewCustomTemplateEntries) SetTargetType(v string) *NotifyStrategyForViewCustomTemplateEntries {
	s.TargetType = &v
	return s
}

func (s *NotifyStrategyForViewCustomTemplateEntries) SetTemplateUuid(v string) *NotifyStrategyForViewCustomTemplateEntries {
	s.TemplateUuid = &v
	return s
}

type NotifyStrategyForViewGroupingSetting struct {
	GroupingKeys []*string `json:"groupingKeys,omitempty" xml:"groupingKeys,omitempty" type:"Repeated"`
	PeriodMin    *int32    `json:"periodMin,omitempty" xml:"periodMin,omitempty"`
	SilenceSec   *int32    `json:"silenceSec,omitempty" xml:"silenceSec,omitempty"`
	Times        *int32    `json:"times,omitempty" xml:"times,omitempty"`
}

func (s NotifyStrategyForViewGroupingSetting) String() string {
	return tea.Prettify(s)
}

func (s NotifyStrategyForViewGroupingSetting) GoString() string {
	return s.String()
}

func (s *NotifyStrategyForViewGroupingSetting) SetGroupingKeys(v []*string) *NotifyStrategyForViewGroupingSetting {
	s.GroupingKeys = v
	return s
}

func (s *NotifyStrategyForViewGroupingSetting) SetPeriodMin(v int32) *NotifyStrategyForViewGroupingSetting {
	s.PeriodMin = &v
	return s
}

func (s *NotifyStrategyForViewGroupingSetting) SetSilenceSec(v int32) *NotifyStrategyForViewGroupingSetting {
	s.SilenceSec = &v
	return s
}

func (s *NotifyStrategyForViewGroupingSetting) SetTimes(v int32) *NotifyStrategyForViewGroupingSetting {
	s.Times = &v
	return s
}

type NotifyStrategyForViewRoutes struct {
	Channels        []*NotifyStrategyForViewRoutesChannels      `json:"channels,omitempty" xml:"channels,omitempty" type:"Repeated"`
	EffectTimeRange *NotifyStrategyForViewRoutesEffectTimeRange `json:"effectTimeRange,omitempty" xml:"effectTimeRange,omitempty" type:"Struct"`
	FilterSetting   *FilterSetting                              `json:"filterSetting,omitempty" xml:"filterSetting,omitempty"`
	Severities      []*string                                   `json:"severities,omitempty" xml:"severities,omitempty" type:"Repeated"`
}

func (s NotifyStrategyForViewRoutes) String() string {
	return tea.Prettify(s)
}

func (s NotifyStrategyForViewRoutes) GoString() string {
	return s.String()
}

func (s *NotifyStrategyForViewRoutes) SetChannels(v []*NotifyStrategyForViewRoutesChannels) *NotifyStrategyForViewRoutes {
	s.Channels = v
	return s
}

func (s *NotifyStrategyForViewRoutes) SetEffectTimeRange(v *NotifyStrategyForViewRoutesEffectTimeRange) *NotifyStrategyForViewRoutes {
	s.EffectTimeRange = v
	return s
}

func (s *NotifyStrategyForViewRoutes) SetFilterSetting(v *FilterSetting) *NotifyStrategyForViewRoutes {
	s.FilterSetting = v
	return s
}

func (s *NotifyStrategyForViewRoutes) SetSeverities(v []*string) *NotifyStrategyForViewRoutes {
	s.Severities = v
	return s
}

type NotifyStrategyForViewRoutesChannels struct {
	// This parameter is required.
	ChannelType        *string   `json:"channelType,omitempty" xml:"channelType,omitempty"`
	EnabledSubChannels []*string `json:"enabledSubChannels,omitempty" xml:"enabledSubChannels,omitempty" type:"Repeated"`
	// This parameter is required.
	Receivers []*string `json:"receivers,omitempty" xml:"receivers,omitempty" type:"Repeated"`
}

func (s NotifyStrategyForViewRoutesChannels) String() string {
	return tea.Prettify(s)
}

func (s NotifyStrategyForViewRoutesChannels) GoString() string {
	return s.String()
}

func (s *NotifyStrategyForViewRoutesChannels) SetChannelType(v string) *NotifyStrategyForViewRoutesChannels {
	s.ChannelType = &v
	return s
}

func (s *NotifyStrategyForViewRoutesChannels) SetEnabledSubChannels(v []*string) *NotifyStrategyForViewRoutesChannels {
	s.EnabledSubChannels = v
	return s
}

func (s *NotifyStrategyForViewRoutesChannels) SetReceivers(v []*string) *NotifyStrategyForViewRoutesChannels {
	s.Receivers = v
	return s
}

type NotifyStrategyForViewRoutesEffectTimeRange struct {
	DayInWeek         []*int32 `json:"dayInWeek,omitempty" xml:"dayInWeek,omitempty" type:"Repeated"`
	EndTimeInMinute   *int32   `json:"endTimeInMinute,omitempty" xml:"endTimeInMinute,omitempty"`
	StartTimeInMinute *int32   `json:"startTimeInMinute,omitempty" xml:"startTimeInMinute,omitempty"`
	TimeZone          *string  `json:"timeZone,omitempty" xml:"timeZone,omitempty"`
}

func (s NotifyStrategyForViewRoutesEffectTimeRange) String() string {
	return tea.Prettify(s)
}

func (s NotifyStrategyForViewRoutesEffectTimeRange) GoString() string {
	return s.String()
}

func (s *NotifyStrategyForViewRoutesEffectTimeRange) SetDayInWeek(v []*int32) *NotifyStrategyForViewRoutesEffectTimeRange {
	s.DayInWeek = v
	return s
}

func (s *NotifyStrategyForViewRoutesEffectTimeRange) SetEndTimeInMinute(v int32) *NotifyStrategyForViewRoutesEffectTimeRange {
	s.EndTimeInMinute = &v
	return s
}

func (s *NotifyStrategyForViewRoutesEffectTimeRange) SetStartTimeInMinute(v int32) *NotifyStrategyForViewRoutesEffectTimeRange {
	s.StartTimeInMinute = &v
	return s
}

func (s *NotifyStrategyForViewRoutesEffectTimeRange) SetTimeZone(v string) *NotifyStrategyForViewRoutesEffectTimeRange {
	s.TimeZone = &v
	return s
}

type PrometheusManagedInstance struct {
	CreateTime             *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	InstanceType           *string `json:"instanceType,omitempty" xml:"instanceType,omitempty"`
	PrometheusInstanceId   *string `json:"prometheusInstanceId,omitempty" xml:"prometheusInstanceId,omitempty"`
	PrometheusInstanceName *string `json:"prometheusInstanceName,omitempty" xml:"prometheusInstanceName,omitempty"`
	RegionId               *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	Status                 *string `json:"status,omitempty" xml:"status,omitempty"`
	Workspace              *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s PrometheusManagedInstance) String() string {
	return tea.Prettify(s)
}

func (s PrometheusManagedInstance) GoString() string {
	return s.String()
}

func (s *PrometheusManagedInstance) SetCreateTime(v string) *PrometheusManagedInstance {
	s.CreateTime = &v
	return s
}

func (s *PrometheusManagedInstance) SetInstanceType(v string) *PrometheusManagedInstance {
	s.InstanceType = &v
	return s
}

func (s *PrometheusManagedInstance) SetPrometheusInstanceId(v string) *PrometheusManagedInstance {
	s.PrometheusInstanceId = &v
	return s
}

func (s *PrometheusManagedInstance) SetPrometheusInstanceName(v string) *PrometheusManagedInstance {
	s.PrometheusInstanceName = &v
	return s
}

func (s *PrometheusManagedInstance) SetRegionId(v string) *PrometheusManagedInstance {
	s.RegionId = &v
	return s
}

func (s *PrometheusManagedInstance) SetStatus(v string) *PrometheusManagedInstance {
	s.Status = &v
	return s
}

func (s *PrometheusManagedInstance) SetWorkspace(v string) *PrometheusManagedInstance {
	s.Workspace = &v
	return s
}

type RumDnsResponse struct {
	Domain  *string `json:"domain,omitempty" xml:"domain,omitempty"`
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	Result  *bool   `json:"result,omitempty" xml:"result,omitempty"`
}

func (s RumDnsResponse) String() string {
	return tea.Prettify(s)
}

func (s RumDnsResponse) GoString() string {
	return s.String()
}

func (s *RumDnsResponse) SetDomain(v string) *RumDnsResponse {
	s.Domain = &v
	return s
}

func (s *RumDnsResponse) SetMessage(v string) *RumDnsResponse {
	s.Message = &v
	return s
}

func (s *RumDnsResponse) SetResult(v bool) *RumDnsResponse {
	s.Result = &v
	return s
}

type SubscriptionForModify struct {
	Description      *string                              `json:"description,omitempty" xml:"description,omitempty"`
	FilterSetting    *FilterSetting                       `json:"filterSetting,omitempty" xml:"filterSetting,omitempty"`
	NotifyStrategyId *string                              `json:"notifyStrategyId,omitempty" xml:"notifyStrategyId,omitempty"`
	PushingSetting   *SubscriptionForModifyPushingSetting `json:"pushingSetting,omitempty" xml:"pushingSetting,omitempty" type:"Struct"`
	// This parameter is required.
	SubscriptionName *string `json:"subscriptionName,omitempty" xml:"subscriptionName,omitempty"`
}

func (s SubscriptionForModify) String() string {
	return tea.Prettify(s)
}

func (s SubscriptionForModify) GoString() string {
	return s.String()
}

func (s *SubscriptionForModify) SetDescription(v string) *SubscriptionForModify {
	s.Description = &v
	return s
}

func (s *SubscriptionForModify) SetFilterSetting(v *FilterSetting) *SubscriptionForModify {
	s.FilterSetting = v
	return s
}

func (s *SubscriptionForModify) SetNotifyStrategyId(v string) *SubscriptionForModify {
	s.NotifyStrategyId = &v
	return s
}

func (s *SubscriptionForModify) SetPushingSetting(v *SubscriptionForModifyPushingSetting) *SubscriptionForModify {
	s.PushingSetting = v
	return s
}

func (s *SubscriptionForModify) SetSubscriptionName(v string) *SubscriptionForModify {
	s.SubscriptionName = &v
	return s
}

type SubscriptionForModifyPushingSetting struct {
	AlertActionIds   []*string `json:"alertActionIds,omitempty" xml:"alertActionIds,omitempty" type:"Repeated"`
	ResponsePlanId   *string   `json:"responsePlanId,omitempty" xml:"responsePlanId,omitempty"`
	RestoreActionIds []*string `json:"restoreActionIds,omitempty" xml:"restoreActionIds,omitempty" type:"Repeated"`
	TemplateUuid     *string   `json:"templateUuid,omitempty" xml:"templateUuid,omitempty"`
}

func (s SubscriptionForModifyPushingSetting) String() string {
	return tea.Prettify(s)
}

func (s SubscriptionForModifyPushingSetting) GoString() string {
	return s.String()
}

func (s *SubscriptionForModifyPushingSetting) SetAlertActionIds(v []*string) *SubscriptionForModifyPushingSetting {
	s.AlertActionIds = v
	return s
}

func (s *SubscriptionForModifyPushingSetting) SetResponsePlanId(v string) *SubscriptionForModifyPushingSetting {
	s.ResponsePlanId = &v
	return s
}

func (s *SubscriptionForModifyPushingSetting) SetRestoreActionIds(v []*string) *SubscriptionForModifyPushingSetting {
	s.RestoreActionIds = v
	return s
}

func (s *SubscriptionForModifyPushingSetting) SetTemplateUuid(v string) *SubscriptionForModifyPushingSetting {
	s.TemplateUuid = &v
	return s
}

type SubscriptionForView struct {
	CreateTime       *string                            `json:"createTime,omitempty" xml:"createTime,omitempty"`
	Description      *string                            `json:"description,omitempty" xml:"description,omitempty"`
	Enable           *bool                              `json:"enable,omitempty" xml:"enable,omitempty"`
	FilterSetting    *FilterSetting                     `json:"filterSetting,omitempty" xml:"filterSetting,omitempty"`
	NotifyStrategyId *string                            `json:"notifyStrategyId,omitempty" xml:"notifyStrategyId,omitempty"`
	PushingSetting   *SubscriptionForViewPushingSetting `json:"pushingSetting,omitempty" xml:"pushingSetting,omitempty" type:"Struct"`
	SubscriptionId   *string                            `json:"subscriptionId,omitempty" xml:"subscriptionId,omitempty"`
	// This parameter is required.
	SubscriptionName *string `json:"subscriptionName,omitempty" xml:"subscriptionName,omitempty"`
	UpdateTime       *string `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
	UserId           *string `json:"userId,omitempty" xml:"userId,omitempty"`
	Workspace        *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s SubscriptionForView) String() string {
	return tea.Prettify(s)
}

func (s SubscriptionForView) GoString() string {
	return s.String()
}

func (s *SubscriptionForView) SetCreateTime(v string) *SubscriptionForView {
	s.CreateTime = &v
	return s
}

func (s *SubscriptionForView) SetDescription(v string) *SubscriptionForView {
	s.Description = &v
	return s
}

func (s *SubscriptionForView) SetEnable(v bool) *SubscriptionForView {
	s.Enable = &v
	return s
}

func (s *SubscriptionForView) SetFilterSetting(v *FilterSetting) *SubscriptionForView {
	s.FilterSetting = v
	return s
}

func (s *SubscriptionForView) SetNotifyStrategyId(v string) *SubscriptionForView {
	s.NotifyStrategyId = &v
	return s
}

func (s *SubscriptionForView) SetPushingSetting(v *SubscriptionForViewPushingSetting) *SubscriptionForView {
	s.PushingSetting = v
	return s
}

func (s *SubscriptionForView) SetSubscriptionId(v string) *SubscriptionForView {
	s.SubscriptionId = &v
	return s
}

func (s *SubscriptionForView) SetSubscriptionName(v string) *SubscriptionForView {
	s.SubscriptionName = &v
	return s
}

func (s *SubscriptionForView) SetUpdateTime(v string) *SubscriptionForView {
	s.UpdateTime = &v
	return s
}

func (s *SubscriptionForView) SetUserId(v string) *SubscriptionForView {
	s.UserId = &v
	return s
}

func (s *SubscriptionForView) SetWorkspace(v string) *SubscriptionForView {
	s.Workspace = &v
	return s
}

type SubscriptionForViewPushingSetting struct {
	AlertActionIds   []*string `json:"alertActionIds,omitempty" xml:"alertActionIds,omitempty" type:"Repeated"`
	ResponsePlanId   *string   `json:"responsePlanId,omitempty" xml:"responsePlanId,omitempty"`
	RestoreActionIds []*string `json:"restoreActionIds,omitempty" xml:"restoreActionIds,omitempty" type:"Repeated"`
	TemplateUuid     *string   `json:"templateUuid,omitempty" xml:"templateUuid,omitempty"`
}

func (s SubscriptionForViewPushingSetting) String() string {
	return tea.Prettify(s)
}

func (s SubscriptionForViewPushingSetting) GoString() string {
	return s.String()
}

func (s *SubscriptionForViewPushingSetting) SetAlertActionIds(v []*string) *SubscriptionForViewPushingSetting {
	s.AlertActionIds = v
	return s
}

func (s *SubscriptionForViewPushingSetting) SetResponsePlanId(v string) *SubscriptionForViewPushingSetting {
	s.ResponsePlanId = &v
	return s
}

func (s *SubscriptionForViewPushingSetting) SetRestoreActionIds(v []*string) *SubscriptionForViewPushingSetting {
	s.RestoreActionIds = v
	return s
}

func (s *SubscriptionForViewPushingSetting) SetTemplateUuid(v string) *SubscriptionForViewPushingSetting {
	s.TemplateUuid = &v
	return s
}

type TransformAction struct {
	FilterSetting *FilterSetting     `json:"filterSetting,omitempty" xml:"filterSetting,omitempty"`
	LabelKey      *string            `json:"labelKey,omitempty" xml:"labelKey,omitempty"`
	Mapping       map[string]*string `json:"mapping,omitempty" xml:"mapping,omitempty"`
	RegExp        *string            `json:"regExp,omitempty" xml:"regExp,omitempty"`
	Source        *string            `json:"source,omitempty" xml:"source,omitempty"`
	Target        *string            `json:"target,omitempty" xml:"target,omitempty"`
	Type          *string            `json:"type,omitempty" xml:"type,omitempty"`
	Value         *string            `json:"value,omitempty" xml:"value,omitempty"`
	Variable      *string            `json:"variable,omitempty" xml:"variable,omitempty"`
}

func (s TransformAction) String() string {
	return tea.Prettify(s)
}

func (s TransformAction) GoString() string {
	return s.String()
}

func (s *TransformAction) SetFilterSetting(v *FilterSetting) *TransformAction {
	s.FilterSetting = v
	return s
}

func (s *TransformAction) SetLabelKey(v string) *TransformAction {
	s.LabelKey = &v
	return s
}

func (s *TransformAction) SetMapping(v map[string]*string) *TransformAction {
	s.Mapping = v
	return s
}

func (s *TransformAction) SetRegExp(v string) *TransformAction {
	s.RegExp = &v
	return s
}

func (s *TransformAction) SetSource(v string) *TransformAction {
	s.Source = &v
	return s
}

func (s *TransformAction) SetTarget(v string) *TransformAction {
	s.Target = &v
	return s
}

func (s *TransformAction) SetType(v string) *TransformAction {
	s.Type = &v
	return s
}

func (s *TransformAction) SetValue(v string) *TransformAction {
	s.Value = &v
	return s
}

func (s *TransformAction) SetVariable(v string) *TransformAction {
	s.Variable = &v
	return s
}

type TransformerForModify struct {
	Actions        []*TransformAction `json:"actions,omitempty" xml:"actions,omitempty" type:"Repeated"`
	Description    *string            `json:"description,omitempty" xml:"description,omitempty"`
	FilterSetting  *FilterSetting     `json:"filterSetting,omitempty" xml:"filterSetting,omitempty"`
	QuitAfterMatch *bool              `json:"quitAfterMatch,omitempty" xml:"quitAfterMatch,omitempty"`
	SortId         *int32             `json:"sortId,omitempty" xml:"sortId,omitempty"`
	// This parameter is required.
	TransformerName *string `json:"transformerName,omitempty" xml:"transformerName,omitempty"`
}

func (s TransformerForModify) String() string {
	return tea.Prettify(s)
}

func (s TransformerForModify) GoString() string {
	return s.String()
}

func (s *TransformerForModify) SetActions(v []*TransformAction) *TransformerForModify {
	s.Actions = v
	return s
}

func (s *TransformerForModify) SetDescription(v string) *TransformerForModify {
	s.Description = &v
	return s
}

func (s *TransformerForModify) SetFilterSetting(v *FilterSetting) *TransformerForModify {
	s.FilterSetting = v
	return s
}

func (s *TransformerForModify) SetQuitAfterMatch(v bool) *TransformerForModify {
	s.QuitAfterMatch = &v
	return s
}

func (s *TransformerForModify) SetSortId(v int32) *TransformerForModify {
	s.SortId = &v
	return s
}

func (s *TransformerForModify) SetTransformerName(v string) *TransformerForModify {
	s.TransformerName = &v
	return s
}

type TransformerForView struct {
	Actions        []*TransformAction `json:"actions,omitempty" xml:"actions,omitempty" type:"Repeated"`
	CreateTime     *string            `json:"createTime,omitempty" xml:"createTime,omitempty"`
	Description    *string            `json:"description,omitempty" xml:"description,omitempty"`
	Enable         *bool              `json:"enable,omitempty" xml:"enable,omitempty"`
	FilterSetting  *FilterSetting     `json:"filterSetting,omitempty" xml:"filterSetting,omitempty"`
	QuitAfterMatch *bool              `json:"quitAfterMatch,omitempty" xml:"quitAfterMatch,omitempty"`
	SortId         *int32             `json:"sortId,omitempty" xml:"sortId,omitempty"`
	TransformerId  *string            `json:"transformerId,omitempty" xml:"transformerId,omitempty"`
	// This parameter is required.
	TransformerName *string `json:"transformerName,omitempty" xml:"transformerName,omitempty"`
	UpdateTime      *string `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
	UserId          *string `json:"userId,omitempty" xml:"userId,omitempty"`
	Workspace       *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s TransformerForView) String() string {
	return tea.Prettify(s)
}

func (s TransformerForView) GoString() string {
	return s.String()
}

func (s *TransformerForView) SetActions(v []*TransformAction) *TransformerForView {
	s.Actions = v
	return s
}

func (s *TransformerForView) SetCreateTime(v string) *TransformerForView {
	s.CreateTime = &v
	return s
}

func (s *TransformerForView) SetDescription(v string) *TransformerForView {
	s.Description = &v
	return s
}

func (s *TransformerForView) SetEnable(v bool) *TransformerForView {
	s.Enable = &v
	return s
}

func (s *TransformerForView) SetFilterSetting(v *FilterSetting) *TransformerForView {
	s.FilterSetting = v
	return s
}

func (s *TransformerForView) SetQuitAfterMatch(v bool) *TransformerForView {
	s.QuitAfterMatch = &v
	return s
}

func (s *TransformerForView) SetSortId(v int32) *TransformerForView {
	s.SortId = &v
	return s
}

func (s *TransformerForView) SetTransformerId(v string) *TransformerForView {
	s.TransformerId = &v
	return s
}

func (s *TransformerForView) SetTransformerName(v string) *TransformerForView {
	s.TransformerName = &v
	return s
}

func (s *TransformerForView) SetUpdateTime(v string) *TransformerForView {
	s.UpdateTime = &v
	return s
}

func (s *TransformerForView) SetUserId(v string) *TransformerForView {
	s.UserId = &v
	return s
}

func (s *TransformerForView) SetWorkspace(v string) *TransformerForView {
	s.Workspace = &v
	return s
}

type CreateEntityStoreResponseBody struct {
	// example:
	//
	// 264C3E89-XXXX-XXXX-XXXX-CE9C2196C7DC
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// workspace-test-001
	WorkspaceName *string `json:"workspaceName,omitempty" xml:"workspaceName,omitempty"`
}

func (s CreateEntityStoreResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateEntityStoreResponseBody) GoString() string {
	return s.String()
}

func (s *CreateEntityStoreResponseBody) SetRequestId(v string) *CreateEntityStoreResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateEntityStoreResponseBody) SetWorkspaceName(v string) *CreateEntityStoreResponseBody {
	s.WorkspaceName = &v
	return s
}

type CreateEntityStoreResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateEntityStoreResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateEntityStoreResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateEntityStoreResponse) GoString() string {
	return s.String()
}

func (s *CreateEntityStoreResponse) SetHeaders(v map[string]*string) *CreateEntityStoreResponse {
	s.Headers = v
	return s
}

func (s *CreateEntityStoreResponse) SetStatusCode(v int32) *CreateEntityStoreResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateEntityStoreResponse) SetBody(v *CreateEntityStoreResponseBody) *CreateEntityStoreResponse {
	s.Body = v
	return s
}

type CreatePrometheusInstanceRequest struct {
	// if can be null:
	// true
	//
	// example:
	//
	// 60
	ArchiveDuration *int32 `json:"archiveDuration,omitempty" xml:"archiveDuration,omitempty"`
	// example:
	//
	// 0.0.0.0/0
	AuthFreeReadPolicy *string `json:"authFreeReadPolicy,omitempty" xml:"authFreeReadPolicy,omitempty"`
	// example:
	//
	// 0.0.0.0/0
	AuthFreeWritePolicy *string `json:"authFreeWritePolicy,omitempty" xml:"authFreeWritePolicy,omitempty"`
	// example:
	//
	// true
	EnableAuthFreeRead *bool `json:"enableAuthFreeRead,omitempty" xml:"enableAuthFreeRead,omitempty"`
	// example:
	//
	// true
	EnableAuthFreeWrite *bool `json:"enableAuthFreeWrite,omitempty" xml:"enableAuthFreeWrite,omitempty"`
	// example:
	//
	// true
	EnableAuthToken *bool `json:"enableAuthToken,omitempty" xml:"enableAuthToken,omitempty"`
	// example:
	//
	// POSTPAY
	PaymentType *string `json:"paymentType,omitempty" xml:"paymentType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// name1
	PrometheusInstanceName *string `json:"prometheusInstanceName,omitempty" xml:"prometheusInstanceName,omitempty"`
	// example:
	//
	// Running
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// 90
	StorageDuration *int32                                 `json:"storageDuration,omitempty" xml:"storageDuration,omitempty"`
	Tags            []*CreatePrometheusInstanceRequestTags `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
	// example:
	//
	// wokspace1
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s CreatePrometheusInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreatePrometheusInstanceRequest) GoString() string {
	return s.String()
}

func (s *CreatePrometheusInstanceRequest) SetArchiveDuration(v int32) *CreatePrometheusInstanceRequest {
	s.ArchiveDuration = &v
	return s
}

func (s *CreatePrometheusInstanceRequest) SetAuthFreeReadPolicy(v string) *CreatePrometheusInstanceRequest {
	s.AuthFreeReadPolicy = &v
	return s
}

func (s *CreatePrometheusInstanceRequest) SetAuthFreeWritePolicy(v string) *CreatePrometheusInstanceRequest {
	s.AuthFreeWritePolicy = &v
	return s
}

func (s *CreatePrometheusInstanceRequest) SetEnableAuthFreeRead(v bool) *CreatePrometheusInstanceRequest {
	s.EnableAuthFreeRead = &v
	return s
}

func (s *CreatePrometheusInstanceRequest) SetEnableAuthFreeWrite(v bool) *CreatePrometheusInstanceRequest {
	s.EnableAuthFreeWrite = &v
	return s
}

func (s *CreatePrometheusInstanceRequest) SetEnableAuthToken(v bool) *CreatePrometheusInstanceRequest {
	s.EnableAuthToken = &v
	return s
}

func (s *CreatePrometheusInstanceRequest) SetPaymentType(v string) *CreatePrometheusInstanceRequest {
	s.PaymentType = &v
	return s
}

func (s *CreatePrometheusInstanceRequest) SetPrometheusInstanceName(v string) *CreatePrometheusInstanceRequest {
	s.PrometheusInstanceName = &v
	return s
}

func (s *CreatePrometheusInstanceRequest) SetStatus(v string) *CreatePrometheusInstanceRequest {
	s.Status = &v
	return s
}

func (s *CreatePrometheusInstanceRequest) SetStorageDuration(v int32) *CreatePrometheusInstanceRequest {
	s.StorageDuration = &v
	return s
}

func (s *CreatePrometheusInstanceRequest) SetTags(v []*CreatePrometheusInstanceRequestTags) *CreatePrometheusInstanceRequest {
	s.Tags = v
	return s
}

func (s *CreatePrometheusInstanceRequest) SetWorkspace(v string) *CreatePrometheusInstanceRequest {
	s.Workspace = &v
	return s
}

type CreatePrometheusInstanceRequestTags struct {
	// example:
	//
	// key1
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// example:
	//
	// 110109200001214284
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s CreatePrometheusInstanceRequestTags) String() string {
	return tea.Prettify(s)
}

func (s CreatePrometheusInstanceRequestTags) GoString() string {
	return s.String()
}

func (s *CreatePrometheusInstanceRequestTags) SetKey(v string) *CreatePrometheusInstanceRequestTags {
	s.Key = &v
	return s
}

func (s *CreatePrometheusInstanceRequestTags) SetValue(v string) *CreatePrometheusInstanceRequestTags {
	s.Value = &v
	return s
}

type CreatePrometheusInstanceResponseBody struct {
	// example:
	//
	// rw-abc123
	PrometheusInstanceId *string `json:"prometheusInstanceId,omitempty" xml:"prometheusInstanceId,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 264C3E89-BE6E-5F82-A484-CE9C2196C7DC
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreatePrometheusInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreatePrometheusInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePrometheusInstanceResponseBody) SetPrometheusInstanceId(v string) *CreatePrometheusInstanceResponseBody {
	s.PrometheusInstanceId = &v
	return s
}

func (s *CreatePrometheusInstanceResponseBody) SetRequestId(v string) *CreatePrometheusInstanceResponseBody {
	s.RequestId = &v
	return s
}

type CreatePrometheusInstanceResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreatePrometheusInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePrometheusInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s CreatePrometheusInstanceResponse) GoString() string {
	return s.String()
}

func (s *CreatePrometheusInstanceResponse) SetHeaders(v map[string]*string) *CreatePrometheusInstanceResponse {
	s.Headers = v
	return s
}

func (s *CreatePrometheusInstanceResponse) SetStatusCode(v int32) *CreatePrometheusInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePrometheusInstanceResponse) SetBody(v *CreatePrometheusInstanceResponseBody) *CreatePrometheusInstanceResponse {
	s.Body = v
	return s
}

type CreateUmodelRequest struct {
	CommonSchemaRef []*CreateUmodelRequestCommonSchemaRef `json:"commonSchemaRef,omitempty" xml:"commonSchemaRef,omitempty" type:"Repeated"`
	// example:
	//
	// workspace test
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
}

func (s CreateUmodelRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateUmodelRequest) GoString() string {
	return s.String()
}

func (s *CreateUmodelRequest) SetCommonSchemaRef(v []*CreateUmodelRequestCommonSchemaRef) *CreateUmodelRequest {
	s.CommonSchemaRef = v
	return s
}

func (s *CreateUmodelRequest) SetDescription(v string) *CreateUmodelRequest {
	s.Description = &v
	return s
}

type CreateUmodelRequestCommonSchemaRef struct {
	// example:
	//
	// test-job-123
	Group *string   `json:"group,omitempty" xml:"group,omitempty"`
	Items []*string `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s CreateUmodelRequestCommonSchemaRef) String() string {
	return tea.Prettify(s)
}

func (s CreateUmodelRequestCommonSchemaRef) GoString() string {
	return s.String()
}

func (s *CreateUmodelRequestCommonSchemaRef) SetGroup(v string) *CreateUmodelRequestCommonSchemaRef {
	s.Group = &v
	return s
}

func (s *CreateUmodelRequestCommonSchemaRef) SetItems(v []*string) *CreateUmodelRequestCommonSchemaRef {
	s.Items = v
	return s
}

func (s *CreateUmodelRequestCommonSchemaRef) SetVersion(v string) *CreateUmodelRequestCommonSchemaRef {
	s.Version = &v
	return s
}

type CreateUmodelResponseBody struct {
	// example:
	//
	// 123-0F43-23423-AC43-34234
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// workspace-test
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s CreateUmodelResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateUmodelResponseBody) GoString() string {
	return s.String()
}

func (s *CreateUmodelResponseBody) SetRequestId(v string) *CreateUmodelResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateUmodelResponseBody) SetWorkspace(v string) *CreateUmodelResponseBody {
	s.Workspace = &v
	return s
}

type CreateUmodelResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateUmodelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateUmodelResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateUmodelResponse) GoString() string {
	return s.String()
}

func (s *CreateUmodelResponse) SetHeaders(v map[string]*string) *CreateUmodelResponse {
	s.Headers = v
	return s
}

func (s *CreateUmodelResponse) SetStatusCode(v int32) *CreateUmodelResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateUmodelResponse) SetBody(v *CreateUmodelResponseBody) *CreateUmodelResponse {
	s.Body = v
	return s
}

type DeleteEntityStoreResponseBody struct {
	// example:
	//
	// 264C3E89-XXXX-XXXX-XXXX-CE9C2196C7DC
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeleteEntityStoreResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteEntityStoreResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteEntityStoreResponseBody) SetRequestId(v string) *DeleteEntityStoreResponseBody {
	s.RequestId = &v
	return s
}

type DeleteEntityStoreResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteEntityStoreResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteEntityStoreResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteEntityStoreResponse) GoString() string {
	return s.String()
}

func (s *DeleteEntityStoreResponse) SetHeaders(v map[string]*string) *DeleteEntityStoreResponse {
	s.Headers = v
	return s
}

func (s *DeleteEntityStoreResponse) SetStatusCode(v int32) *DeleteEntityStoreResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteEntityStoreResponse) SetBody(v *DeleteEntityStoreResponseBody) *DeleteEntityStoreResponse {
	s.Body = v
	return s
}

type DeleteUmodelResponseBody struct {
	// example:
	//
	// 123123-3213-345-9941-345345345
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeleteUmodelResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteUmodelResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteUmodelResponseBody) SetRequestId(v string) *DeleteUmodelResponseBody {
	s.RequestId = &v
	return s
}

type DeleteUmodelResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteUmodelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteUmodelResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteUmodelResponse) GoString() string {
	return s.String()
}

func (s *DeleteUmodelResponse) SetHeaders(v map[string]*string) *DeleteUmodelResponse {
	s.Headers = v
	return s
}

func (s *DeleteUmodelResponse) SetStatusCode(v int32) *DeleteUmodelResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteUmodelResponse) SetBody(v *DeleteUmodelResponseBody) *DeleteUmodelResponse {
	s.Body = v
	return s
}

type DeleteUmodelDataRequest struct {
	// example:
	//
	// apm
	Domain *string `json:"domain,omitempty" xml:"domain,omitempty"`
	// example:
	//
	// metric_set
	Kind *string `json:"kind,omitempty" xml:"kind,omitempty"`
	// example:
	//
	// test
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s DeleteUmodelDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteUmodelDataRequest) GoString() string {
	return s.String()
}

func (s *DeleteUmodelDataRequest) SetDomain(v string) *DeleteUmodelDataRequest {
	s.Domain = &v
	return s
}

func (s *DeleteUmodelDataRequest) SetKind(v string) *DeleteUmodelDataRequest {
	s.Kind = &v
	return s
}

func (s *DeleteUmodelDataRequest) SetName(v string) *DeleteUmodelDataRequest {
	s.Name = &v
	return s
}

type DeleteUmodelDataResponseBody struct {
	// example:
	//
	// 111111-222-333-1111-33333
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeleteUmodelDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteUmodelDataResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteUmodelDataResponseBody) SetRequestId(v string) *DeleteUmodelDataResponseBody {
	s.RequestId = &v
	return s
}

type DeleteUmodelDataResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteUmodelDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteUmodelDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteUmodelDataResponse) GoString() string {
	return s.String()
}

func (s *DeleteUmodelDataResponse) SetHeaders(v map[string]*string) *DeleteUmodelDataResponse {
	s.Headers = v
	return s
}

func (s *DeleteUmodelDataResponse) SetStatusCode(v int32) *DeleteUmodelDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteUmodelDataResponse) SetBody(v *DeleteUmodelDataResponseBody) *DeleteUmodelDataResponse {
	s.Body = v
	return s
}

type DeleteWorkspaceResponseBody struct {
	// example:
	//
	// 264C3E89-XXXX-XXXX-XXXX-CE9C2196C7DC
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeleteWorkspaceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteWorkspaceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteWorkspaceResponseBody) SetRequestId(v string) *DeleteWorkspaceResponseBody {
	s.RequestId = &v
	return s
}

type DeleteWorkspaceResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteWorkspaceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteWorkspaceResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteWorkspaceResponse) GoString() string {
	return s.String()
}

func (s *DeleteWorkspaceResponse) SetHeaders(v map[string]*string) *DeleteWorkspaceResponse {
	s.Headers = v
	return s
}

func (s *DeleteWorkspaceResponse) SetStatusCode(v int32) *DeleteWorkspaceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteWorkspaceResponse) SetBody(v *DeleteWorkspaceResponseBody) *DeleteWorkspaceResponse {
	s.Body = v
	return s
}

type GetEntityStoreResponseBody struct {
	// example:
	//
	// cn-heyuan
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// example:
	//
	// 264C3E89-XXXX-XXXX-XXXX-CE9C2196C7DC
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// workspace-test-001
	WorkspaceName *string `json:"workspaceName,omitempty" xml:"workspaceName,omitempty"`
}

func (s GetEntityStoreResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetEntityStoreResponseBody) GoString() string {
	return s.String()
}

func (s *GetEntityStoreResponseBody) SetRegionId(v string) *GetEntityStoreResponseBody {
	s.RegionId = &v
	return s
}

func (s *GetEntityStoreResponseBody) SetRequestId(v string) *GetEntityStoreResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetEntityStoreResponseBody) SetWorkspaceName(v string) *GetEntityStoreResponseBody {
	s.WorkspaceName = &v
	return s
}

type GetEntityStoreResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetEntityStoreResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetEntityStoreResponse) String() string {
	return tea.Prettify(s)
}

func (s GetEntityStoreResponse) GoString() string {
	return s.String()
}

func (s *GetEntityStoreResponse) SetHeaders(v map[string]*string) *GetEntityStoreResponse {
	s.Headers = v
	return s
}

func (s *GetEntityStoreResponse) SetStatusCode(v int32) *GetEntityStoreResponse {
	s.StatusCode = &v
	return s
}

func (s *GetEntityStoreResponse) SetBody(v *GetEntityStoreResponseBody) *GetEntityStoreResponse {
	s.Body = v
	return s
}

type GetEntityStoreDataHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// gzip
	AcceptEncoding *string `json:"acceptEncoding,omitempty" xml:"acceptEncoding,omitempty"`
}

func (s GetEntityStoreDataHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetEntityStoreDataHeaders) GoString() string {
	return s.String()
}

func (s *GetEntityStoreDataHeaders) SetCommonHeaders(v map[string]*string) *GetEntityStoreDataHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetEntityStoreDataHeaders) SetAcceptEncoding(v string) *GetEntityStoreDataHeaders {
	s.AcceptEncoding = &v
	return s
}

type GetEntityStoreDataRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1721767203
	From *int32 `json:"from,omitempty" xml:"from,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// .entity with(domain=\\"acs\\", type=\\"acs.k8s.node\\") | limit 0, 10
	Query *string `json:"query,omitempty" xml:"query,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1721767283
	To *int32 `json:"to,omitempty" xml:"to,omitempty"`
}

func (s GetEntityStoreDataRequest) String() string {
	return tea.Prettify(s)
}

func (s GetEntityStoreDataRequest) GoString() string {
	return s.String()
}

func (s *GetEntityStoreDataRequest) SetFrom(v int32) *GetEntityStoreDataRequest {
	s.From = &v
	return s
}

func (s *GetEntityStoreDataRequest) SetQuery(v string) *GetEntityStoreDataRequest {
	s.Query = &v
	return s
}

func (s *GetEntityStoreDataRequest) SetTo(v int32) *GetEntityStoreDataRequest {
	s.To = &v
	return s
}

type GetEntityStoreDataResponseBody struct {
	Data   [][]*string `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	Header []*string   `json:"header,omitempty" xml:"header,omitempty" type:"Repeated"`
	// example:
	//
	// 264C3E89-XXXX-XXXX-XXXX-CE9C2196C7DC
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetEntityStoreDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetEntityStoreDataResponseBody) GoString() string {
	return s.String()
}

func (s *GetEntityStoreDataResponseBody) SetData(v [][]*string) *GetEntityStoreDataResponseBody {
	s.Data = v
	return s
}

func (s *GetEntityStoreDataResponseBody) SetHeader(v []*string) *GetEntityStoreDataResponseBody {
	s.Header = v
	return s
}

func (s *GetEntityStoreDataResponseBody) SetRequestId(v string) *GetEntityStoreDataResponseBody {
	s.RequestId = &v
	return s
}

type GetEntityStoreDataResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetEntityStoreDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetEntityStoreDataResponse) String() string {
	return tea.Prettify(s)
}

func (s GetEntityStoreDataResponse) GoString() string {
	return s.String()
}

func (s *GetEntityStoreDataResponse) SetHeaders(v map[string]*string) *GetEntityStoreDataResponse {
	s.Headers = v
	return s
}

func (s *GetEntityStoreDataResponse) SetStatusCode(v int32) *GetEntityStoreDataResponse {
	s.StatusCode = &v
	return s
}

func (s *GetEntityStoreDataResponse) SetBody(v *GetEntityStoreDataResponseBody) *GetEntityStoreDataResponse {
	s.Body = v
	return s
}

type GetUmodelResponseBody struct {
	CommonSchemaRef []*GetUmodelResponseBodyCommonSchemaRef `json:"commonSchemaRef,omitempty" xml:"commonSchemaRef,omitempty" type:"Repeated"`
	// example:
	//
	// workspace test
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// cn-heyuan
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// example:
	//
	// 123-123123-sdf-435-3123
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// workspace-test
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s GetUmodelResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetUmodelResponseBody) GoString() string {
	return s.String()
}

func (s *GetUmodelResponseBody) SetCommonSchemaRef(v []*GetUmodelResponseBodyCommonSchemaRef) *GetUmodelResponseBody {
	s.CommonSchemaRef = v
	return s
}

func (s *GetUmodelResponseBody) SetDescription(v string) *GetUmodelResponseBody {
	s.Description = &v
	return s
}

func (s *GetUmodelResponseBody) SetRegionId(v string) *GetUmodelResponseBody {
	s.RegionId = &v
	return s
}

func (s *GetUmodelResponseBody) SetRequestId(v string) *GetUmodelResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetUmodelResponseBody) SetWorkspace(v string) *GetUmodelResponseBody {
	s.Workspace = &v
	return s
}

type GetUmodelResponseBodyCommonSchemaRef struct {
	// example:
	//
	// test-job-123123
	Group *string   `json:"group,omitempty" xml:"group,omitempty"`
	Items []*string `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// example:
	//
	// 5
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s GetUmodelResponseBodyCommonSchemaRef) String() string {
	return tea.Prettify(s)
}

func (s GetUmodelResponseBodyCommonSchemaRef) GoString() string {
	return s.String()
}

func (s *GetUmodelResponseBodyCommonSchemaRef) SetGroup(v string) *GetUmodelResponseBodyCommonSchemaRef {
	s.Group = &v
	return s
}

func (s *GetUmodelResponseBodyCommonSchemaRef) SetItems(v []*string) *GetUmodelResponseBodyCommonSchemaRef {
	s.Items = v
	return s
}

func (s *GetUmodelResponseBodyCommonSchemaRef) SetVersion(v string) *GetUmodelResponseBodyCommonSchemaRef {
	s.Version = &v
	return s
}

type GetUmodelResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetUmodelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetUmodelResponse) String() string {
	return tea.Prettify(s)
}

func (s GetUmodelResponse) GoString() string {
	return s.String()
}

func (s *GetUmodelResponse) SetHeaders(v map[string]*string) *GetUmodelResponse {
	s.Headers = v
	return s
}

func (s *GetUmodelResponse) SetStatusCode(v int32) *GetUmodelResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUmodelResponse) SetBody(v *GetUmodelResponseBody) *GetUmodelResponse {
	s.Body = v
	return s
}

type GetUmodelDataRequest struct {
	// example:
	//
	// {
	//
	// 	"filter": {
	//
	// 		"domains": []
	//
	// 	},
	//
	// 	"offset": 0,
	//
	// 	"size": 100000
	//
	// }
	Content interface{} `json:"content,omitempty" xml:"content,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ListData
	Method *string `json:"method,omitempty" xml:"method,omitempty"`
}

func (s GetUmodelDataRequest) String() string {
	return tea.Prettify(s)
}

func (s GetUmodelDataRequest) GoString() string {
	return s.String()
}

func (s *GetUmodelDataRequest) SetContent(v interface{}) *GetUmodelDataRequest {
	s.Content = v
	return s
}

func (s *GetUmodelDataRequest) SetMethod(v string) *GetUmodelDataRequest {
	s.Method = &v
	return s
}

type GetUmodelDataResponseBody struct {
	Errors []*GetUmodelDataResponseBodyErrors `json:"errors,omitempty" xml:"errors,omitempty" type:"Repeated"`
	Links  []interface{}                      `json:"links,omitempty" xml:"links,omitempty" type:"Repeated"`
	Nodes  []interface{}                      `json:"nodes,omitempty" xml:"nodes,omitempty" type:"Repeated"`
	// example:
	//
	// 123-123-234-345-123
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 0
	TotalLinksCount *int32 `json:"totalLinksCount,omitempty" xml:"totalLinksCount,omitempty"`
	// example:
	//
	// 0
	TotalNodesCount *int32 `json:"totalNodesCount,omitempty" xml:"totalNodesCount,omitempty"`
}

func (s GetUmodelDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetUmodelDataResponseBody) GoString() string {
	return s.String()
}

func (s *GetUmodelDataResponseBody) SetErrors(v []*GetUmodelDataResponseBodyErrors) *GetUmodelDataResponseBody {
	s.Errors = v
	return s
}

func (s *GetUmodelDataResponseBody) SetLinks(v []interface{}) *GetUmodelDataResponseBody {
	s.Links = v
	return s
}

func (s *GetUmodelDataResponseBody) SetNodes(v []interface{}) *GetUmodelDataResponseBody {
	s.Nodes = v
	return s
}

func (s *GetUmodelDataResponseBody) SetRequestId(v string) *GetUmodelDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetUmodelDataResponseBody) SetTotalLinksCount(v int32) *GetUmodelDataResponseBody {
	s.TotalLinksCount = &v
	return s
}

func (s *GetUmodelDataResponseBody) SetTotalNodesCount(v int32) *GetUmodelDataResponseBody {
	s.TotalNodesCount = &v
	return s
}

type GetUmodelDataResponseBodyErrors struct {
	// example:
	//
	// ok
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// external
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetUmodelDataResponseBodyErrors) String() string {
	return tea.Prettify(s)
}

func (s GetUmodelDataResponseBodyErrors) GoString() string {
	return s.String()
}

func (s *GetUmodelDataResponseBodyErrors) SetMessage(v string) *GetUmodelDataResponseBodyErrors {
	s.Message = &v
	return s
}

func (s *GetUmodelDataResponseBodyErrors) SetType(v string) *GetUmodelDataResponseBodyErrors {
	s.Type = &v
	return s
}

type GetUmodelDataResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetUmodelDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetUmodelDataResponse) String() string {
	return tea.Prettify(s)
}

func (s GetUmodelDataResponse) GoString() string {
	return s.String()
}

func (s *GetUmodelDataResponse) SetHeaders(v map[string]*string) *GetUmodelDataResponse {
	s.Headers = v
	return s
}

func (s *GetUmodelDataResponse) SetStatusCode(v int32) *GetUmodelDataResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUmodelDataResponse) SetBody(v *GetUmodelDataResponseBody) *GetUmodelDataResponse {
	s.Body = v
	return s
}

type GetWorkspaceResponseBody struct {
	// Use the UTC time format: yyyy-MM-ddTHH:mm:ss.SSSZ
	//
	// example:
	//
	// 2025-03-11T08:21:58Z
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// 工作空间描述
	//
	// example:
	//
	// workspace test
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// workspace-test
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	// Use the UTC time format: yyyy-MM-ddTHH:mm:ssZ
	//
	// example:
	//
	// 2025-03-11T08:21:58Z
	LastModifyTime *string `json:"lastModifyTime,omitempty" xml:"lastModifyTime,omitempty"`
	// 地域ID
	//
	// example:
	//
	// cn-heyuan
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// example:
	//
	// 264C3E89-XXXX-XXXX-XXXX-CE9C2196C7DC
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// 工作空间绑定的日志服务项目名称
	//
	// example:
	//
	// sls-project-test-001
	SlsProject *string `json:"slsProject,omitempty" xml:"slsProject,omitempty"`
	// 工作空间名称
	//
	// This parameter is required.
	//
	// example:
	//
	// workspace-test-001
	WorkspaceName *string `json:"workspaceName,omitempty" xml:"workspaceName,omitempty"`
}

func (s GetWorkspaceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetWorkspaceResponseBody) GoString() string {
	return s.String()
}

func (s *GetWorkspaceResponseBody) SetCreateTime(v string) *GetWorkspaceResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetWorkspaceResponseBody) SetDescription(v string) *GetWorkspaceResponseBody {
	s.Description = &v
	return s
}

func (s *GetWorkspaceResponseBody) SetDisplayName(v string) *GetWorkspaceResponseBody {
	s.DisplayName = &v
	return s
}

func (s *GetWorkspaceResponseBody) SetLastModifyTime(v string) *GetWorkspaceResponseBody {
	s.LastModifyTime = &v
	return s
}

func (s *GetWorkspaceResponseBody) SetRegionId(v string) *GetWorkspaceResponseBody {
	s.RegionId = &v
	return s
}

func (s *GetWorkspaceResponseBody) SetRequestId(v string) *GetWorkspaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetWorkspaceResponseBody) SetSlsProject(v string) *GetWorkspaceResponseBody {
	s.SlsProject = &v
	return s
}

func (s *GetWorkspaceResponseBody) SetWorkspaceName(v string) *GetWorkspaceResponseBody {
	s.WorkspaceName = &v
	return s
}

type GetWorkspaceResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetWorkspaceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetWorkspaceResponse) String() string {
	return tea.Prettify(s)
}

func (s GetWorkspaceResponse) GoString() string {
	return s.String()
}

func (s *GetWorkspaceResponse) SetHeaders(v map[string]*string) *GetWorkspaceResponse {
	s.Headers = v
	return s
}

func (s *GetWorkspaceResponse) SetStatusCode(v int32) *GetWorkspaceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetWorkspaceResponse) SetBody(v *GetWorkspaceResponseBody) *GetWorkspaceResponse {
	s.Body = v
	return s
}

type ListAlertActionsRequest struct {
	AlertActionIds  []*string `json:"alertActionIds,omitempty" xml:"alertActionIds,omitempty" type:"Repeated"`
	AlertActionName *string   `json:"alertActionName,omitempty" xml:"alertActionName,omitempty"`
	PageNumber      *int32    `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize        *int32    `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	Type            *string   `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListAlertActionsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAlertActionsRequest) GoString() string {
	return s.String()
}

func (s *ListAlertActionsRequest) SetAlertActionIds(v []*string) *ListAlertActionsRequest {
	s.AlertActionIds = v
	return s
}

func (s *ListAlertActionsRequest) SetAlertActionName(v string) *ListAlertActionsRequest {
	s.AlertActionName = &v
	return s
}

func (s *ListAlertActionsRequest) SetPageNumber(v int32) *ListAlertActionsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListAlertActionsRequest) SetPageSize(v int32) *ListAlertActionsRequest {
	s.PageSize = &v
	return s
}

func (s *ListAlertActionsRequest) SetType(v string) *ListAlertActionsRequest {
	s.Type = &v
	return s
}

type ListAlertActionsShrinkRequest struct {
	AlertActionIdsShrink *string `json:"alertActionIds,omitempty" xml:"alertActionIds,omitempty"`
	AlertActionName      *string `json:"alertActionName,omitempty" xml:"alertActionName,omitempty"`
	PageNumber           *int32  `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize             *int32  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	Type                 *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListAlertActionsShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAlertActionsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListAlertActionsShrinkRequest) SetAlertActionIdsShrink(v string) *ListAlertActionsShrinkRequest {
	s.AlertActionIdsShrink = &v
	return s
}

func (s *ListAlertActionsShrinkRequest) SetAlertActionName(v string) *ListAlertActionsShrinkRequest {
	s.AlertActionName = &v
	return s
}

func (s *ListAlertActionsShrinkRequest) SetPageNumber(v int32) *ListAlertActionsShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListAlertActionsShrinkRequest) SetPageSize(v int32) *ListAlertActionsShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListAlertActionsShrinkRequest) SetType(v string) *ListAlertActionsShrinkRequest {
	s.Type = &v
	return s
}

type ListAlertActionsResponseBody struct {
	AlertActions []*ListAlertActionsResponseBodyAlertActions `json:"alertActions,omitempty" xml:"alertActions,omitempty" type:"Repeated"`
	PageNumber   *int64                                      `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize     *int64                                      `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	RequestId    *string                                     `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Total        *int32                                      `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListAlertActionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAlertActionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAlertActionsResponseBody) SetAlertActions(v []*ListAlertActionsResponseBodyAlertActions) *ListAlertActionsResponseBody {
	s.AlertActions = v
	return s
}

func (s *ListAlertActionsResponseBody) SetPageNumber(v int64) *ListAlertActionsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListAlertActionsResponseBody) SetPageSize(v int64) *ListAlertActionsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListAlertActionsResponseBody) SetRequestId(v string) *ListAlertActionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAlertActionsResponseBody) SetTotal(v int32) *ListAlertActionsResponseBody {
	s.Total = &v
	return s
}

type ListAlertActionsResponseBodyAlertActions struct {
	AlertActionId   *string                                                 `json:"alertActionId,omitempty" xml:"alertActionId,omitempty"`
	AlertActionName *string                                                 `json:"alertActionName,omitempty" xml:"alertActionName,omitempty"`
	EssParam        *ListAlertActionsResponseBodyAlertActionsEssParam       `json:"essParam,omitempty" xml:"essParam,omitempty" type:"Struct"`
	FcParam         *ListAlertActionsResponseBodyAlertActionsFcParam        `json:"fcParam,omitempty" xml:"fcParam,omitempty" type:"Struct"`
	MnsParam        *ListAlertActionsResponseBodyAlertActionsMnsParam       `json:"mnsParam,omitempty" xml:"mnsParam,omitempty" type:"Struct"`
	PagerDutyParam  *ListAlertActionsResponseBodyAlertActionsPagerDutyParam `json:"pagerDutyParam,omitempty" xml:"pagerDutyParam,omitempty" type:"Struct"`
	SlsParam        *ListAlertActionsResponseBodyAlertActionsSlsParam       `json:"slsParam,omitempty" xml:"slsParam,omitempty" type:"Struct"`
	Type            *string                                                 `json:"type,omitempty" xml:"type,omitempty"`
	WebhookParam    *ListAlertActionsResponseBodyAlertActionsWebhookParam   `json:"webhookParam,omitempty" xml:"webhookParam,omitempty" type:"Struct"`
}

func (s ListAlertActionsResponseBodyAlertActions) String() string {
	return tea.Prettify(s)
}

func (s ListAlertActionsResponseBodyAlertActions) GoString() string {
	return s.String()
}

func (s *ListAlertActionsResponseBodyAlertActions) SetAlertActionId(v string) *ListAlertActionsResponseBodyAlertActions {
	s.AlertActionId = &v
	return s
}

func (s *ListAlertActionsResponseBodyAlertActions) SetAlertActionName(v string) *ListAlertActionsResponseBodyAlertActions {
	s.AlertActionName = &v
	return s
}

func (s *ListAlertActionsResponseBodyAlertActions) SetEssParam(v *ListAlertActionsResponseBodyAlertActionsEssParam) *ListAlertActionsResponseBodyAlertActions {
	s.EssParam = v
	return s
}

func (s *ListAlertActionsResponseBodyAlertActions) SetFcParam(v *ListAlertActionsResponseBodyAlertActionsFcParam) *ListAlertActionsResponseBodyAlertActions {
	s.FcParam = v
	return s
}

func (s *ListAlertActionsResponseBodyAlertActions) SetMnsParam(v *ListAlertActionsResponseBodyAlertActionsMnsParam) *ListAlertActionsResponseBodyAlertActions {
	s.MnsParam = v
	return s
}

func (s *ListAlertActionsResponseBodyAlertActions) SetPagerDutyParam(v *ListAlertActionsResponseBodyAlertActionsPagerDutyParam) *ListAlertActionsResponseBodyAlertActions {
	s.PagerDutyParam = v
	return s
}

func (s *ListAlertActionsResponseBodyAlertActions) SetSlsParam(v *ListAlertActionsResponseBodyAlertActionsSlsParam) *ListAlertActionsResponseBodyAlertActions {
	s.SlsParam = v
	return s
}

func (s *ListAlertActionsResponseBodyAlertActions) SetType(v string) *ListAlertActionsResponseBodyAlertActions {
	s.Type = &v
	return s
}

func (s *ListAlertActionsResponseBodyAlertActions) SetWebhookParam(v *ListAlertActionsResponseBodyAlertActionsWebhookParam) *ListAlertActionsResponseBodyAlertActions {
	s.WebhookParam = v
	return s
}

type ListAlertActionsResponseBodyAlertActionsEssParam struct {
	EssGroupId *string `json:"essGroupId,omitempty" xml:"essGroupId,omitempty"`
	EssRuleId  *string `json:"essRuleId,omitempty" xml:"essRuleId,omitempty"`
	RegionId   *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
}

func (s ListAlertActionsResponseBodyAlertActionsEssParam) String() string {
	return tea.Prettify(s)
}

func (s ListAlertActionsResponseBodyAlertActionsEssParam) GoString() string {
	return s.String()
}

func (s *ListAlertActionsResponseBodyAlertActionsEssParam) SetEssGroupId(v string) *ListAlertActionsResponseBodyAlertActionsEssParam {
	s.EssGroupId = &v
	return s
}

func (s *ListAlertActionsResponseBodyAlertActionsEssParam) SetEssRuleId(v string) *ListAlertActionsResponseBodyAlertActionsEssParam {
	s.EssRuleId = &v
	return s
}

func (s *ListAlertActionsResponseBodyAlertActionsEssParam) SetRegionId(v string) *ListAlertActionsResponseBodyAlertActionsEssParam {
	s.RegionId = &v
	return s
}

type ListAlertActionsResponseBodyAlertActionsFcParam struct {
	Function *string `json:"function,omitempty" xml:"function,omitempty"`
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	Service  *string `json:"service,omitempty" xml:"service,omitempty"`
}

func (s ListAlertActionsResponseBodyAlertActionsFcParam) String() string {
	return tea.Prettify(s)
}

func (s ListAlertActionsResponseBodyAlertActionsFcParam) GoString() string {
	return s.String()
}

func (s *ListAlertActionsResponseBodyAlertActionsFcParam) SetFunction(v string) *ListAlertActionsResponseBodyAlertActionsFcParam {
	s.Function = &v
	return s
}

func (s *ListAlertActionsResponseBodyAlertActionsFcParam) SetRegionId(v string) *ListAlertActionsResponseBodyAlertActionsFcParam {
	s.RegionId = &v
	return s
}

func (s *ListAlertActionsResponseBodyAlertActionsFcParam) SetService(v string) *ListAlertActionsResponseBodyAlertActionsFcParam {
	s.Service = &v
	return s
}

type ListAlertActionsResponseBodyAlertActionsMnsParam struct {
	MnsType  *string `json:"mnsType,omitempty" xml:"mnsType,omitempty"`
	Name     *string `json:"name,omitempty" xml:"name,omitempty"`
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
}

func (s ListAlertActionsResponseBodyAlertActionsMnsParam) String() string {
	return tea.Prettify(s)
}

func (s ListAlertActionsResponseBodyAlertActionsMnsParam) GoString() string {
	return s.String()
}

func (s *ListAlertActionsResponseBodyAlertActionsMnsParam) SetMnsType(v string) *ListAlertActionsResponseBodyAlertActionsMnsParam {
	s.MnsType = &v
	return s
}

func (s *ListAlertActionsResponseBodyAlertActionsMnsParam) SetName(v string) *ListAlertActionsResponseBodyAlertActionsMnsParam {
	s.Name = &v
	return s
}

func (s *ListAlertActionsResponseBodyAlertActionsMnsParam) SetRegionId(v string) *ListAlertActionsResponseBodyAlertActionsMnsParam {
	s.RegionId = &v
	return s
}

type ListAlertActionsResponseBodyAlertActionsPagerDutyParam struct {
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s ListAlertActionsResponseBodyAlertActionsPagerDutyParam) String() string {
	return tea.Prettify(s)
}

func (s ListAlertActionsResponseBodyAlertActionsPagerDutyParam) GoString() string {
	return s.String()
}

func (s *ListAlertActionsResponseBodyAlertActionsPagerDutyParam) SetKey(v string) *ListAlertActionsResponseBodyAlertActionsPagerDutyParam {
	s.Key = &v
	return s
}

func (s *ListAlertActionsResponseBodyAlertActionsPagerDutyParam) SetUrl(v string) *ListAlertActionsResponseBodyAlertActionsPagerDutyParam {
	s.Url = &v
	return s
}

type ListAlertActionsResponseBodyAlertActionsSlsParam struct {
	Logstore *string `json:"logstore,omitempty" xml:"logstore,omitempty"`
	Project  *string `json:"project,omitempty" xml:"project,omitempty"`
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
}

func (s ListAlertActionsResponseBodyAlertActionsSlsParam) String() string {
	return tea.Prettify(s)
}

func (s ListAlertActionsResponseBodyAlertActionsSlsParam) GoString() string {
	return s.String()
}

func (s *ListAlertActionsResponseBodyAlertActionsSlsParam) SetLogstore(v string) *ListAlertActionsResponseBodyAlertActionsSlsParam {
	s.Logstore = &v
	return s
}

func (s *ListAlertActionsResponseBodyAlertActionsSlsParam) SetProject(v string) *ListAlertActionsResponseBodyAlertActionsSlsParam {
	s.Project = &v
	return s
}

func (s *ListAlertActionsResponseBodyAlertActionsSlsParam) SetRegionId(v string) *ListAlertActionsResponseBodyAlertActionsSlsParam {
	s.RegionId = &v
	return s
}

type ListAlertActionsResponseBodyAlertActionsWebhookParam struct {
	ContentType *string            `json:"contentType,omitempty" xml:"contentType,omitempty"`
	Headers     map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	Method      *string            `json:"method,omitempty" xml:"method,omitempty"`
	Url         *string            `json:"url,omitempty" xml:"url,omitempty"`
}

func (s ListAlertActionsResponseBodyAlertActionsWebhookParam) String() string {
	return tea.Prettify(s)
}

func (s ListAlertActionsResponseBodyAlertActionsWebhookParam) GoString() string {
	return s.String()
}

func (s *ListAlertActionsResponseBodyAlertActionsWebhookParam) SetContentType(v string) *ListAlertActionsResponseBodyAlertActionsWebhookParam {
	s.ContentType = &v
	return s
}

func (s *ListAlertActionsResponseBodyAlertActionsWebhookParam) SetHeaders(v map[string]*string) *ListAlertActionsResponseBodyAlertActionsWebhookParam {
	s.Headers = v
	return s
}

func (s *ListAlertActionsResponseBodyAlertActionsWebhookParam) SetMethod(v string) *ListAlertActionsResponseBodyAlertActionsWebhookParam {
	s.Method = &v
	return s
}

func (s *ListAlertActionsResponseBodyAlertActionsWebhookParam) SetUrl(v string) *ListAlertActionsResponseBodyAlertActionsWebhookParam {
	s.Url = &v
	return s
}

type ListAlertActionsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAlertActionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAlertActionsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAlertActionsResponse) GoString() string {
	return s.String()
}

func (s *ListAlertActionsResponse) SetHeaders(v map[string]*string) *ListAlertActionsResponse {
	s.Headers = v
	return s
}

func (s *ListAlertActionsResponse) SetStatusCode(v int32) *ListAlertActionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAlertActionsResponse) SetBody(v *ListAlertActionsResponseBody) *ListAlertActionsResponse {
	s.Body = v
	return s
}

type ListWorkspacesRequest struct {
	// example:
	//
	// 10
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// xxxxxxxxxx
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// if can be null:
	// true
	//
	// example:
	//
	// cn-heyuan
	Region *string `json:"region,omitempty" xml:"region,omitempty"`
	// example:
	//
	// workspace-test-001
	WorkspaceName *string `json:"workspaceName,omitempty" xml:"workspaceName,omitempty"`
	// if can be null:
	// true
	//
	// example:
	//
	// workspace-test-001
	WorkspaceNameList []*string `json:"workspaceNameList,omitempty" xml:"workspaceNameList,omitempty" type:"Repeated"`
}

func (s ListWorkspacesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListWorkspacesRequest) GoString() string {
	return s.String()
}

func (s *ListWorkspacesRequest) SetMaxResults(v int32) *ListWorkspacesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListWorkspacesRequest) SetNextToken(v string) *ListWorkspacesRequest {
	s.NextToken = &v
	return s
}

func (s *ListWorkspacesRequest) SetRegion(v string) *ListWorkspacesRequest {
	s.Region = &v
	return s
}

func (s *ListWorkspacesRequest) SetWorkspaceName(v string) *ListWorkspacesRequest {
	s.WorkspaceName = &v
	return s
}

func (s *ListWorkspacesRequest) SetWorkspaceNameList(v []*string) *ListWorkspacesRequest {
	s.WorkspaceNameList = v
	return s
}

type ListWorkspacesShrinkRequest struct {
	// example:
	//
	// 10
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// xxxxxxxxxx
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// if can be null:
	// true
	//
	// example:
	//
	// cn-heyuan
	Region *string `json:"region,omitempty" xml:"region,omitempty"`
	// example:
	//
	// workspace-test-001
	WorkspaceName *string `json:"workspaceName,omitempty" xml:"workspaceName,omitempty"`
	// if can be null:
	// true
	//
	// example:
	//
	// workspace-test-001
	WorkspaceNameListShrink *string `json:"workspaceNameList,omitempty" xml:"workspaceNameList,omitempty"`
}

func (s ListWorkspacesShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListWorkspacesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListWorkspacesShrinkRequest) SetMaxResults(v int32) *ListWorkspacesShrinkRequest {
	s.MaxResults = &v
	return s
}

func (s *ListWorkspacesShrinkRequest) SetNextToken(v string) *ListWorkspacesShrinkRequest {
	s.NextToken = &v
	return s
}

func (s *ListWorkspacesShrinkRequest) SetRegion(v string) *ListWorkspacesShrinkRequest {
	s.Region = &v
	return s
}

func (s *ListWorkspacesShrinkRequest) SetWorkspaceName(v string) *ListWorkspacesShrinkRequest {
	s.WorkspaceName = &v
	return s
}

func (s *ListWorkspacesShrinkRequest) SetWorkspaceNameListShrink(v string) *ListWorkspacesShrinkRequest {
	s.WorkspaceNameListShrink = &v
	return s
}

type ListWorkspacesResponseBody struct {
	// example:
	//
	// 10
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// xxxxxxxxxx
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// example:
	//
	// 264C3E89-XXXX-XXXX-XXXX-CE9C2196C7DC
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 1
	Total      *int32                                  `json:"total,omitempty" xml:"total,omitempty"`
	Workspaces []*ListWorkspacesResponseBodyWorkspaces `json:"workspaces,omitempty" xml:"workspaces,omitempty" type:"Repeated"`
}

func (s ListWorkspacesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListWorkspacesResponseBody) GoString() string {
	return s.String()
}

func (s *ListWorkspacesResponseBody) SetMaxResults(v int32) *ListWorkspacesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListWorkspacesResponseBody) SetNextToken(v string) *ListWorkspacesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListWorkspacesResponseBody) SetRequestId(v string) *ListWorkspacesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListWorkspacesResponseBody) SetTotal(v int32) *ListWorkspacesResponseBody {
	s.Total = &v
	return s
}

func (s *ListWorkspacesResponseBody) SetWorkspaces(v []*ListWorkspacesResponseBodyWorkspaces) *ListWorkspacesResponseBody {
	s.Workspaces = v
	return s
}

type ListWorkspacesResponseBodyWorkspaces struct {
	// Use the UTC time format: yyyy-MM-ddTHH:mm:ssZ
	//
	// example:
	//
	// 2025-03-11T08:21:58Z
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// 工作空间描述
	//
	// example:
	//
	// workspace test
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// workspace-test
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	// Use the UTC time format: yyyy-MM-ddTHH:mm:ssZ
	//
	// example:
	//
	// 2025-03-11T08:21:58Z
	LastModifyTime *string `json:"lastModifyTime,omitempty" xml:"lastModifyTime,omitempty"`
	// 地域ID
	//
	// example:
	//
	// cn-heyuan
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// 工作空间绑定的日志服务项目名称
	//
	// example:
	//
	// sls-project-test-001
	SlsProject *string `json:"slsProject,omitempty" xml:"slsProject,omitempty"`
	// 工作空间名称
	//
	// This parameter is required.
	//
	// example:
	//
	// workspace-test-001
	WorkspaceName *string `json:"workspaceName,omitempty" xml:"workspaceName,omitempty"`
}

func (s ListWorkspacesResponseBodyWorkspaces) String() string {
	return tea.Prettify(s)
}

func (s ListWorkspacesResponseBodyWorkspaces) GoString() string {
	return s.String()
}

func (s *ListWorkspacesResponseBodyWorkspaces) SetCreateTime(v string) *ListWorkspacesResponseBodyWorkspaces {
	s.CreateTime = &v
	return s
}

func (s *ListWorkspacesResponseBodyWorkspaces) SetDescription(v string) *ListWorkspacesResponseBodyWorkspaces {
	s.Description = &v
	return s
}

func (s *ListWorkspacesResponseBodyWorkspaces) SetDisplayName(v string) *ListWorkspacesResponseBodyWorkspaces {
	s.DisplayName = &v
	return s
}

func (s *ListWorkspacesResponseBodyWorkspaces) SetLastModifyTime(v string) *ListWorkspacesResponseBodyWorkspaces {
	s.LastModifyTime = &v
	return s
}

func (s *ListWorkspacesResponseBodyWorkspaces) SetRegionId(v string) *ListWorkspacesResponseBodyWorkspaces {
	s.RegionId = &v
	return s
}

func (s *ListWorkspacesResponseBodyWorkspaces) SetSlsProject(v string) *ListWorkspacesResponseBodyWorkspaces {
	s.SlsProject = &v
	return s
}

func (s *ListWorkspacesResponseBodyWorkspaces) SetWorkspaceName(v string) *ListWorkspacesResponseBodyWorkspaces {
	s.WorkspaceName = &v
	return s
}

type ListWorkspacesResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListWorkspacesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListWorkspacesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListWorkspacesResponse) GoString() string {
	return s.String()
}

func (s *ListWorkspacesResponse) SetHeaders(v map[string]*string) *ListWorkspacesResponse {
	s.Headers = v
	return s
}

func (s *ListWorkspacesResponse) SetStatusCode(v int32) *ListWorkspacesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListWorkspacesResponse) SetBody(v *ListWorkspacesResponseBody) *ListWorkspacesResponse {
	s.Body = v
	return s
}

type PutWorkspaceRequest struct {
	// 工作空间描述
	//
	// example:
	//
	// workspace test
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// workspace-test
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	// 工作空间绑定的日志服务项目名称
	//
	// This parameter is required.
	//
	// example:
	//
	// sls-project-test-001
	SlsProject *string `json:"slsProject,omitempty" xml:"slsProject,omitempty"`
}

func (s PutWorkspaceRequest) String() string {
	return tea.Prettify(s)
}

func (s PutWorkspaceRequest) GoString() string {
	return s.String()
}

func (s *PutWorkspaceRequest) SetDescription(v string) *PutWorkspaceRequest {
	s.Description = &v
	return s
}

func (s *PutWorkspaceRequest) SetDisplayName(v string) *PutWorkspaceRequest {
	s.DisplayName = &v
	return s
}

func (s *PutWorkspaceRequest) SetSlsProject(v string) *PutWorkspaceRequest {
	s.SlsProject = &v
	return s
}

type PutWorkspaceResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 264C3E89-XXXX-XXXX-XXXX-CE9C2196C7DC
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// workspace-test-001
	WorkspaceName *string `json:"workspaceName,omitempty" xml:"workspaceName,omitempty"`
}

func (s PutWorkspaceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PutWorkspaceResponseBody) GoString() string {
	return s.String()
}

func (s *PutWorkspaceResponseBody) SetRequestId(v string) *PutWorkspaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *PutWorkspaceResponseBody) SetWorkspaceName(v string) *PutWorkspaceResponseBody {
	s.WorkspaceName = &v
	return s
}

type PutWorkspaceResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PutWorkspaceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PutWorkspaceResponse) String() string {
	return tea.Prettify(s)
}

func (s PutWorkspaceResponse) GoString() string {
	return s.String()
}

func (s *PutWorkspaceResponse) SetHeaders(v map[string]*string) *PutWorkspaceResponse {
	s.Headers = v
	return s
}

func (s *PutWorkspaceResponse) SetStatusCode(v int32) *PutWorkspaceResponse {
	s.StatusCode = &v
	return s
}

func (s *PutWorkspaceResponse) SetBody(v *PutWorkspaceResponseBody) *PutWorkspaceResponse {
	s.Body = v
	return s
}

type UpdateUmodelRequest struct {
	CommonSchemaRef []*UpdateUmodelRequestCommonSchemaRef `json:"commonSchemaRef,omitempty" xml:"commonSchemaRef,omitempty" type:"Repeated"`
	// example:
	//
	// workspace test
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
}

func (s UpdateUmodelRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateUmodelRequest) GoString() string {
	return s.String()
}

func (s *UpdateUmodelRequest) SetCommonSchemaRef(v []*UpdateUmodelRequestCommonSchemaRef) *UpdateUmodelRequest {
	s.CommonSchemaRef = v
	return s
}

func (s *UpdateUmodelRequest) SetDescription(v string) *UpdateUmodelRequest {
	s.Description = &v
	return s
}

type UpdateUmodelRequestCommonSchemaRef struct {
	// example:
	//
	// test-bmp-123123
	Group *string   `json:"group,omitempty" xml:"group,omitempty"`
	Items []*string `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// example:
	//
	// 2.5.
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s UpdateUmodelRequestCommonSchemaRef) String() string {
	return tea.Prettify(s)
}

func (s UpdateUmodelRequestCommonSchemaRef) GoString() string {
	return s.String()
}

func (s *UpdateUmodelRequestCommonSchemaRef) SetGroup(v string) *UpdateUmodelRequestCommonSchemaRef {
	s.Group = &v
	return s
}

func (s *UpdateUmodelRequestCommonSchemaRef) SetItems(v []*string) *UpdateUmodelRequestCommonSchemaRef {
	s.Items = v
	return s
}

func (s *UpdateUmodelRequestCommonSchemaRef) SetVersion(v string) *UpdateUmodelRequestCommonSchemaRef {
	s.Version = &v
	return s
}

type UpdateUmodelResponseBody struct {
	// example:
	//
	// 234324-123-123-123-23423
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// workspace-test
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s UpdateUmodelResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateUmodelResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateUmodelResponseBody) SetRequestId(v string) *UpdateUmodelResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateUmodelResponseBody) SetWorkspace(v string) *UpdateUmodelResponseBody {
	s.Workspace = &v
	return s
}

type UpdateUmodelResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateUmodelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateUmodelResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateUmodelResponse) GoString() string {
	return s.String()
}

func (s *UpdateUmodelResponse) SetHeaders(v map[string]*string) *UpdateUmodelResponse {
	s.Headers = v
	return s
}

func (s *UpdateUmodelResponse) SetStatusCode(v int32) *UpdateUmodelResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateUmodelResponse) SetBody(v *UpdateUmodelResponseBody) *UpdateUmodelResponse {
	s.Body = v
	return s
}

type UpsertUmodelDataRequest struct {
	Elements []interface{} `json:"elements,omitempty" xml:"elements,omitempty" type:"Repeated"`
	// example:
	//
	// Upsert
	Method *string `json:"method,omitempty" xml:"method,omitempty"`
}

func (s UpsertUmodelDataRequest) String() string {
	return tea.Prettify(s)
}

func (s UpsertUmodelDataRequest) GoString() string {
	return s.String()
}

func (s *UpsertUmodelDataRequest) SetElements(v []interface{}) *UpsertUmodelDataRequest {
	s.Elements = v
	return s
}

func (s *UpsertUmodelDataRequest) SetMethod(v string) *UpsertUmodelDataRequest {
	s.Method = &v
	return s
}

type UpsertUmodelDataResponseBody struct {
	// example:
	//
	// 0CEC5375-C554-562B-A65F-9A629907C1F0
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpsertUmodelDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpsertUmodelDataResponseBody) GoString() string {
	return s.String()
}

func (s *UpsertUmodelDataResponseBody) SetRequestId(v string) *UpsertUmodelDataResponseBody {
	s.RequestId = &v
	return s
}

type UpsertUmodelDataResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpsertUmodelDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpsertUmodelDataResponse) String() string {
	return tea.Prettify(s)
}

func (s UpsertUmodelDataResponse) GoString() string {
	return s.String()
}

func (s *UpsertUmodelDataResponse) SetHeaders(v map[string]*string) *UpsertUmodelDataResponse {
	s.Headers = v
	return s
}

func (s *UpsertUmodelDataResponse) SetStatusCode(v int32) *UpsertUmodelDataResponse {
	s.StatusCode = &v
	return s
}

func (s *UpsertUmodelDataResponse) SetBody(v *UpsertUmodelDataResponseBody) *UpsertUmodelDataResponse {
	s.Body = v
	return s
}

type Client struct {
	openapi.Client
}

func NewClient(config *openapi.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapi.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("cms"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !tea.BoolValue(util.Empty(endpoint)) {
		_result = endpoint
		return _result, _err
	}

	if !tea.BoolValue(util.IsUnset(endpointMap)) && !tea.BoolValue(util.Empty(endpointMap[tea.StringValue(regionId)])) {
		_result = endpointMap[tea.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := endpointutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建EntityStore相关存储
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateEntityStoreResponse
func (client *Client) CreateEntityStoreWithOptions(workspaceName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateEntityStoreResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("CreateEntityStore"),
		Version:     tea.String("2024-03-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/workspace/" + tea.StringValue(openapiutil.GetEncodeParam(workspaceName)) + "/entitystore"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateEntityStoreResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建EntityStore相关存储
//
// @return CreateEntityStoreResponse
func (client *Client) CreateEntityStore(workspaceName *string) (_result *CreateEntityStoreResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateEntityStoreResponse{}
	_body, _err := client.CreateEntityStoreWithOptions(workspaceName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建Prometheus监控实例
//
// @param request - CreatePrometheusInstanceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePrometheusInstanceResponse
func (client *Client) CreatePrometheusInstanceWithOptions(request *CreatePrometheusInstanceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreatePrometheusInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ArchiveDuration)) {
		body["archiveDuration"] = request.ArchiveDuration
	}

	if !tea.BoolValue(util.IsUnset(request.AuthFreeReadPolicy)) {
		body["authFreeReadPolicy"] = request.AuthFreeReadPolicy
	}

	if !tea.BoolValue(util.IsUnset(request.AuthFreeWritePolicy)) {
		body["authFreeWritePolicy"] = request.AuthFreeWritePolicy
	}

	if !tea.BoolValue(util.IsUnset(request.EnableAuthFreeRead)) {
		body["enableAuthFreeRead"] = request.EnableAuthFreeRead
	}

	if !tea.BoolValue(util.IsUnset(request.EnableAuthFreeWrite)) {
		body["enableAuthFreeWrite"] = request.EnableAuthFreeWrite
	}

	if !tea.BoolValue(util.IsUnset(request.EnableAuthToken)) {
		body["enableAuthToken"] = request.EnableAuthToken
	}

	if !tea.BoolValue(util.IsUnset(request.PaymentType)) {
		body["paymentType"] = request.PaymentType
	}

	if !tea.BoolValue(util.IsUnset(request.PrometheusInstanceName)) {
		body["prometheusInstanceName"] = request.PrometheusInstanceName
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		body["status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.StorageDuration)) {
		body["storageDuration"] = request.StorageDuration
	}

	if !tea.BoolValue(util.IsUnset(request.Tags)) {
		body["tags"] = request.Tags
	}

	if !tea.BoolValue(util.IsUnset(request.Workspace)) {
		body["workspace"] = request.Workspace
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreatePrometheusInstance"),
		Version:     tea.String("2024-03-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/prometheus-instances"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreatePrometheusInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建Prometheus监控实例
//
// @param request - CreatePrometheusInstanceRequest
//
// @return CreatePrometheusInstanceResponse
func (client *Client) CreatePrometheusInstance(request *CreatePrometheusInstanceRequest) (_result *CreatePrometheusInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreatePrometheusInstanceResponse{}
	_body, _err := client.CreatePrometheusInstanceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建Umodel配置
//
// @param request - CreateUmodelRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateUmodelResponse
func (client *Client) CreateUmodelWithOptions(workspace *string, request *CreateUmodelRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateUmodelResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CommonSchemaRef)) {
		body["commonSchemaRef"] = request.CommonSchemaRef
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateUmodel"),
		Version:     tea.String("2024-03-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/workspace/" + tea.StringValue(openapiutil.GetEncodeParam(workspace)) + "/umodel"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateUmodelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建Umodel配置
//
// @param request - CreateUmodelRequest
//
// @return CreateUmodelResponse
func (client *Client) CreateUmodel(workspace *string, request *CreateUmodelRequest) (_result *CreateUmodelResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateUmodelResponse{}
	_body, _err := client.CreateUmodelWithOptions(workspace, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除EntityStore相关存储
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteEntityStoreResponse
func (client *Client) DeleteEntityStoreWithOptions(workspaceName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteEntityStoreResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteEntityStore"),
		Version:     tea.String("2024-03-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/workspace/" + tea.StringValue(openapiutil.GetEncodeParam(workspaceName)) + "/entitystore"),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteEntityStoreResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除EntityStore相关存储
//
// @return DeleteEntityStoreResponse
func (client *Client) DeleteEntityStore(workspaceName *string) (_result *DeleteEntityStoreResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteEntityStoreResponse{}
	_body, _err := client.DeleteEntityStoreWithOptions(workspaceName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除Umodel配置信息
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteUmodelResponse
func (client *Client) DeleteUmodelWithOptions(workspace *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteUmodelResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteUmodel"),
		Version:     tea.String("2024-03-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/workspace/" + tea.StringValue(openapiutil.GetEncodeParam(workspace)) + "/umodel"),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteUmodelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除Umodel配置信息
//
// @return DeleteUmodelResponse
func (client *Client) DeleteUmodel(workspace *string) (_result *DeleteUmodelResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteUmodelResponse{}
	_body, _err := client.DeleteUmodelWithOptions(workspace, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除 Umodel Elements
//
// @param request - DeleteUmodelDataRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteUmodelDataResponse
func (client *Client) DeleteUmodelDataWithOptions(workspace *string, request *DeleteUmodelDataRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteUmodelDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Domain)) {
		query["domain"] = request.Domain
	}

	if !tea.BoolValue(util.IsUnset(request.Kind)) {
		query["kind"] = request.Kind
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["name"] = request.Name
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteUmodelData"),
		Version:     tea.String("2024-03-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/workspace/" + tea.StringValue(openapiutil.GetEncodeParam(workspace)) + "/umodel/data"),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteUmodelDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除 Umodel Elements
//
// @param request - DeleteUmodelDataRequest
//
// @return DeleteUmodelDataResponse
func (client *Client) DeleteUmodelData(workspace *string, request *DeleteUmodelDataRequest) (_result *DeleteUmodelDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteUmodelDataResponse{}
	_body, _err := client.DeleteUmodelDataWithOptions(workspace, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除工作空间
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteWorkspaceResponse
func (client *Client) DeleteWorkspaceWithOptions(workspaceName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteWorkspaceResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteWorkspace"),
		Version:     tea.String("2024-03-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/workspace/" + tea.StringValue(openapiutil.GetEncodeParam(workspaceName))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteWorkspaceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除工作空间
//
// @return DeleteWorkspaceResponse
func (client *Client) DeleteWorkspace(workspaceName *string) (_result *DeleteWorkspaceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteWorkspaceResponse{}
	_body, _err := client.DeleteWorkspaceWithOptions(workspaceName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取EntityStore相关存储信息
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetEntityStoreResponse
func (client *Client) GetEntityStoreWithOptions(workspaceName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetEntityStoreResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetEntityStore"),
		Version:     tea.String("2024-03-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/workspace/" + tea.StringValue(openapiutil.GetEncodeParam(workspaceName)) + "/entitystore"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetEntityStoreResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取EntityStore相关存储信息
//
// @return GetEntityStoreResponse
func (client *Client) GetEntityStore(workspaceName *string) (_result *GetEntityStoreResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetEntityStoreResponse{}
	_body, _err := client.GetEntityStoreWithOptions(workspaceName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询指定Workspace下的实体和关系数据，返回结果显示某时间区间中的实体数据（返回结果压缩后传输）。
//
// @param request - GetEntityStoreDataRequest
//
// @param headers - GetEntityStoreDataHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetEntityStoreDataResponse
func (client *Client) GetEntityStoreDataWithOptions(workspace *string, request *GetEntityStoreDataRequest, headers *GetEntityStoreDataHeaders, runtime *util.RuntimeOptions) (_result *GetEntityStoreDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.From)) {
		body["from"] = request.From
	}

	if !tea.BoolValue(util.IsUnset(request.Query)) {
		body["query"] = request.Query
	}

	if !tea.BoolValue(util.IsUnset(request.To)) {
		body["to"] = request.To
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.AcceptEncoding)) {
		realHeaders["acceptEncoding"] = util.ToJSONString(headers.AcceptEncoding)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetEntityStoreData"),
		Version:     tea.String("2024-03-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/workspace/" + tea.StringValue(openapiutil.GetEncodeParam(workspace)) + "/entitiesAndRelations"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetEntityStoreDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询指定Workspace下的实体和关系数据，返回结果显示某时间区间中的实体数据（返回结果压缩后传输）。
//
// @param request - GetEntityStoreDataRequest
//
// @return GetEntityStoreDataResponse
func (client *Client) GetEntityStoreData(workspace *string, request *GetEntityStoreDataRequest) (_result *GetEntityStoreDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetEntityStoreDataHeaders{}
	_result = &GetEntityStoreDataResponse{}
	_body, _err := client.GetEntityStoreDataWithOptions(workspace, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取Umodel配置信息
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUmodelResponse
func (client *Client) GetUmodelWithOptions(workspace *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetUmodelResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetUmodel"),
		Version:     tea.String("2024-03-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/workspace/" + tea.StringValue(openapiutil.GetEncodeParam(workspace)) + "/umodel"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetUmodelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取Umodel配置信息
//
// @return GetUmodelResponse
func (client *Client) GetUmodel(workspace *string) (_result *GetUmodelResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetUmodelResponse{}
	_body, _err := client.GetUmodelWithOptions(workspace, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取相关联的 Umodel 图数据
//
// @param request - GetUmodelDataRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUmodelDataResponse
func (client *Client) GetUmodelDataWithOptions(workspace *string, request *GetUmodelDataRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetUmodelDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Method)) {
		query["method"] = request.Method
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Content)) {
		body["content"] = request.Content
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetUmodelData"),
		Version:     tea.String("2024-03-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/workspace/" + tea.StringValue(openapiutil.GetEncodeParam(workspace)) + "/umodel/graph"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetUmodelDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取相关联的 Umodel 图数据
//
// @param request - GetUmodelDataRequest
//
// @return GetUmodelDataResponse
func (client *Client) GetUmodelData(workspace *string, request *GetUmodelDataRequest) (_result *GetUmodelDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetUmodelDataResponse{}
	_body, _err := client.GetUmodelDataWithOptions(workspace, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取工作空间
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetWorkspaceResponse
func (client *Client) GetWorkspaceWithOptions(workspaceName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetWorkspaceResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetWorkspace"),
		Version:     tea.String("2024-03-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/workspace/" + tea.StringValue(openapiutil.GetEncodeParam(workspaceName))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetWorkspaceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取工作空间
//
// @return GetWorkspaceResponse
func (client *Client) GetWorkspace(workspaceName *string) (_result *GetWorkspaceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetWorkspaceResponse{}
	_body, _err := client.GetWorkspaceWithOptions(workspaceName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询告警动作
//
// @param tmpReq - ListAlertActionsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAlertActionsResponse
func (client *Client) ListAlertActionsWithOptions(tmpReq *ListAlertActionsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListAlertActionsResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListAlertActionsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.AlertActionIds)) {
		request.AlertActionIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AlertActionIds, tea.String("alertActionIds"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AlertActionIdsShrink)) {
		query["alertActionIds"] = request.AlertActionIdsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.AlertActionName)) {
		query["alertActionName"] = request.AlertActionName
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListAlertActions"),
		Version:     tea.String("2024-03-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/alertActions"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAlertActionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询告警动作
//
// @param request - ListAlertActionsRequest
//
// @return ListAlertActionsResponse
func (client *Client) ListAlertActions(request *ListAlertActionsRequest) (_result *ListAlertActionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListAlertActionsResponse{}
	_body, _err := client.ListAlertActionsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取工作空间列表
//
// @param tmpReq - ListWorkspacesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListWorkspacesResponse
func (client *Client) ListWorkspacesWithOptions(tmpReq *ListWorkspacesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListWorkspacesResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListWorkspacesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.WorkspaceNameList)) {
		request.WorkspaceNameListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.WorkspaceNameList, tea.String("workspaceNameList"), tea.String("simple"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["maxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["nextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.Region)) {
		query["region"] = request.Region
	}

	if !tea.BoolValue(util.IsUnset(request.WorkspaceName)) {
		query["workspaceName"] = request.WorkspaceName
	}

	if !tea.BoolValue(util.IsUnset(request.WorkspaceNameListShrink)) {
		query["workspaceNameList"] = request.WorkspaceNameListShrink
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListWorkspaces"),
		Version:     tea.String("2024-03-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/workspace"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListWorkspacesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取工作空间列表
//
// @param request - ListWorkspacesRequest
//
// @return ListWorkspacesResponse
func (client *Client) ListWorkspaces(request *ListWorkspacesRequest) (_result *ListWorkspacesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListWorkspacesResponse{}
	_body, _err := client.ListWorkspacesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建工作空间
//
// @param request - PutWorkspaceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PutWorkspaceResponse
func (client *Client) PutWorkspaceWithOptions(workspaceName *string, request *PutWorkspaceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *PutWorkspaceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.DisplayName)) {
		body["displayName"] = request.DisplayName
	}

	if !tea.BoolValue(util.IsUnset(request.SlsProject)) {
		body["slsProject"] = request.SlsProject
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("PutWorkspace"),
		Version:     tea.String("2024-03-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/workspace/" + tea.StringValue(openapiutil.GetEncodeParam(workspaceName))),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &PutWorkspaceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建工作空间
//
// @param request - PutWorkspaceRequest
//
// @return PutWorkspaceResponse
func (client *Client) PutWorkspace(workspaceName *string, request *PutWorkspaceRequest) (_result *PutWorkspaceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &PutWorkspaceResponse{}
	_body, _err := client.PutWorkspaceWithOptions(workspaceName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新Umodel配置信息
//
// @param request - UpdateUmodelRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateUmodelResponse
func (client *Client) UpdateUmodelWithOptions(workspace *string, request *UpdateUmodelRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateUmodelResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CommonSchemaRef)) {
		body["commonSchemaRef"] = request.CommonSchemaRef
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateUmodel"),
		Version:     tea.String("2024-03-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/workspace/" + tea.StringValue(openapiutil.GetEncodeParam(workspace)) + "/umodel"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateUmodelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新Umodel配置信息
//
// @param request - UpdateUmodelRequest
//
// @return UpdateUmodelResponse
func (client *Client) UpdateUmodel(workspace *string, request *UpdateUmodelRequest) (_result *UpdateUmodelResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateUmodelResponse{}
	_body, _err := client.UpdateUmodelWithOptions(workspace, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 写入 Umodel Elements
//
// @param request - UpsertUmodelDataRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpsertUmodelDataResponse
func (client *Client) UpsertUmodelDataWithOptions(workspace *string, request *UpsertUmodelDataRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpsertUmodelDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Method)) {
		query["method"] = request.Method
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Elements)) {
		body["elements"] = request.Elements
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpsertUmodelData"),
		Version:     tea.String("2024-03-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/workspace/" + tea.StringValue(openapiutil.GetEncodeParam(workspace)) + "/umodel/data"),
		Method:      tea.String("PATCH"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpsertUmodelDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 写入 Umodel Elements
//
// @param request - UpsertUmodelDataRequest
//
// @return UpsertUmodelDataResponse
func (client *Client) UpsertUmodelData(workspace *string, request *UpsertUmodelDataRequest) (_result *UpsertUmodelDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpsertUmodelDataResponse{}
	_body, _err := client.UpsertUmodelDataWithOptions(workspace, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
