// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

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
	DisplayNameCn *string                                       `json:"displayNameCn,omitempty" xml:"displayNameCn,omitempty"`
	DisplayNameEn *string                                       `json:"displayNameEn,omitempty" xml:"displayNameEn,omitempty"`
	Hidden        *bool                                         `json:"hidden,omitempty" xml:"hidden,omitempty"`
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
	DsList []*AlertRuleDataSourceDsList `json:"dsList,omitempty" xml:"dsList,omitempty" type:"Repeated"`
	// 实例id，当type=PROMETHEUS_DS/ENTERPRISE_DS时必填，为prometheus实例的clusterId或指标仓库名称
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	Namespace  *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
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
	Duration       *int64                   `json:"duration,omitempty" xml:"duration,omitempty"`
	Expr           *string                  `json:"expr,omitempty" xml:"expr,omitempty"`
	FirstJoin      *AlertRuleSlsQueryJoin   `json:"firstJoin,omitempty" xml:"firstJoin,omitempty"`
	GroupFieldList []*string                `json:"groupFieldList,omitempty" xml:"groupFieldList,omitempty" type:"Repeated"`
	GroupType      *string                  `json:"groupType,omitempty" xml:"groupType,omitempty"`
	Queries        []*AlertRuleQueryQueries `json:"queries,omitempty" xml:"queries,omitempty" type:"Repeated"`
	SecondJoin     *AlertRuleSlsQueryJoin   `json:"secondJoin,omitempty" xml:"secondJoin,omitempty"`
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
	Duration *int64 `json:"duration,omitempty" xml:"duration,omitempty"`
	// 时间偏移结束时间(相对)，如果指定了start、end，则不指定window。
	End *int64 `json:"end,omitempty" xml:"end,omitempty"`
	// 查询表达式
	//
	// This parameter is required.
	Expr *string `json:"expr,omitempty" xml:"expr,omitempty"`
	// sls查询的时间偏移开始时间(相对)，如果指定了start、end，则不指定window。  例如：start=15， timeUnit=minute，表示15分钟前
	Start *int64 `json:"start,omitempty" xml:"start,omitempty"`
	// start和end、window的时间单位： day/hour/minute/second
	TimeUnit *string `json:"timeUnit,omitempty" xml:"timeUnit,omitempty"`
	// 整点时间查询区间。  如果指定了window则不指定start、end
	Window *string `json:"window,omitempty" xml:"window,omitempty"`
}

func (s AlertRuleQueryQueries) String() string {
	return tea.Prettify(s)
}

func (s AlertRuleQueryQueries) GoString() string {
	return s.String()
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

func (s *AlertRuleQueryQueries) SetWindow(v string) *AlertRuleQueryQueries {
	s.Window = &v
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

type MaintainWindowForModify struct {
	Description   *string        `json:"description,omitempty" xml:"description,omitempty"`
	Effective     *string        `json:"effective,omitempty" xml:"effective,omitempty"`
	EndTime       *string        `json:"endTime,omitempty" xml:"endTime,omitempty"`
	FilterSetting *FilterSetting `json:"filterSetting,omitempty" xml:"filterSetting,omitempty"`
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

type MaintainWindowForView struct {
	CreateTime       *string        `json:"createTime,omitempty" xml:"createTime,omitempty"`
	Description      *string        `json:"description,omitempty" xml:"description,omitempty"`
	Effective        *string        `json:"effective,omitempty" xml:"effective,omitempty"`
	Enable           *bool          `json:"enable,omitempty" xml:"enable,omitempty"`
	EndTime          *string        `json:"endTime,omitempty" xml:"endTime,omitempty"`
	FilterSetting    *FilterSetting `json:"filterSetting,omitempty" xml:"filterSetting,omitempty"`
	MaintainWindowId *string        `json:"maintainWindowId,omitempty" xml:"maintainWindowId,omitempty"`
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
