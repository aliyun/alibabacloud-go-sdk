// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	gatewayclient "github.com/alibabacloud-go/alibabacloud-gateway-sls/client"
	
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type ConsumerGroup struct {
	Name    *string `json:"name,omitempty" xml:"name,omitempty"`
	Order   *bool   `json:"order,omitempty" xml:"order,omitempty"`
	Timeout *int32  `json:"timeout,omitempty" xml:"timeout,omitempty"`
}

func (s ConsumerGroup) String() string {
	return tea.Prettify(s)
}

func (s ConsumerGroup) GoString() string {
	return s.String()
}

func (s *ConsumerGroup) SetName(v string) *ConsumerGroup {
	s.Name = &v
	return s
}

func (s *ConsumerGroup) SetOrder(v bool) *ConsumerGroup {
	s.Order = &v
	return s
}

func (s *ConsumerGroup) SetTimeout(v int32) *ConsumerGroup {
	s.Timeout = &v
	return s
}

type EncryptConf struct {
	Enable      *bool               `json:"enable,omitempty" xml:"enable,omitempty"`
	EncryptType *string             `json:"encrypt_type,omitempty" xml:"encrypt_type,omitempty"`
	UserCmkInfo *EncryptUserCmkConf `json:"user_cmk_info,omitempty" xml:"user_cmk_info,omitempty"`
}

func (s EncryptConf) String() string {
	return tea.Prettify(s)
}

func (s EncryptConf) GoString() string {
	return s.String()
}

func (s *EncryptConf) SetEnable(v bool) *EncryptConf {
	s.Enable = &v
	return s
}

func (s *EncryptConf) SetEncryptType(v string) *EncryptConf {
	s.EncryptType = &v
	return s
}

func (s *EncryptConf) SetUserCmkInfo(v *EncryptUserCmkConf) *EncryptConf {
	s.UserCmkInfo = v
	return s
}

type EncryptUserCmkConf struct {
	Arn      *string `json:"arn,omitempty" xml:"arn,omitempty"`
	CmkKeyId *string `json:"cmk_key_id,omitempty" xml:"cmk_key_id,omitempty"`
	RegionId *string `json:"region_id,omitempty" xml:"region_id,omitempty"`
}

func (s EncryptUserCmkConf) String() string {
	return tea.Prettify(s)
}

func (s EncryptUserCmkConf) GoString() string {
	return s.String()
}

func (s *EncryptUserCmkConf) SetArn(v string) *EncryptUserCmkConf {
	s.Arn = &v
	return s
}

func (s *EncryptUserCmkConf) SetCmkKeyId(v string) *EncryptUserCmkConf {
	s.CmkKeyId = &v
	return s
}

func (s *EncryptUserCmkConf) SetRegionId(v string) *EncryptUserCmkConf {
	s.RegionId = &v
	return s
}

type Histogram struct {
	Count    *int64  `json:"count,omitempty" xml:"count,omitempty"`
	From     *int32  `json:"from,omitempty" xml:"from,omitempty"`
	Progress *string `json:"progress,omitempty" xml:"progress,omitempty"`
	To       *int32  `json:"to,omitempty" xml:"to,omitempty"`
}

func (s Histogram) String() string {
	return tea.Prettify(s)
}

func (s Histogram) GoString() string {
	return s.String()
}

func (s *Histogram) SetCount(v int64) *Histogram {
	s.Count = &v
	return s
}

func (s *Histogram) SetFrom(v int32) *Histogram {
	s.From = &v
	return s
}

func (s *Histogram) SetProgress(v string) *Histogram {
	s.Progress = &v
	return s
}

func (s *Histogram) SetTo(v int32) *Histogram {
	s.To = &v
	return s
}

type LogtailConfig struct {
	ConfigName     *string                    `json:"configName,omitempty" xml:"configName,omitempty"`
	CreateTime     *int64                     `json:"createTime,omitempty" xml:"createTime,omitempty"`
	InputDetail    map[string]interface{}     `json:"inputDetail,omitempty" xml:"inputDetail,omitempty"`
	InputType      *string                    `json:"inputType,omitempty" xml:"inputType,omitempty"`
	LastModifyTime *int64                     `json:"lastModifyTime,omitempty" xml:"lastModifyTime,omitempty"`
	LogSample      *string                    `json:"logSample,omitempty" xml:"logSample,omitempty"`
	OutputDetail   *LogtailConfigOutputDetail `json:"outputDetail,omitempty" xml:"outputDetail,omitempty" type:"Struct"`
	OutputType     *string                    `json:"outputType,omitempty" xml:"outputType,omitempty"`
}

func (s LogtailConfig) String() string {
	return tea.Prettify(s)
}

func (s LogtailConfig) GoString() string {
	return s.String()
}

func (s *LogtailConfig) SetConfigName(v string) *LogtailConfig {
	s.ConfigName = &v
	return s
}

func (s *LogtailConfig) SetCreateTime(v int64) *LogtailConfig {
	s.CreateTime = &v
	return s
}

func (s *LogtailConfig) SetInputDetail(v map[string]interface{}) *LogtailConfig {
	s.InputDetail = v
	return s
}

func (s *LogtailConfig) SetInputType(v string) *LogtailConfig {
	s.InputType = &v
	return s
}

func (s *LogtailConfig) SetLastModifyTime(v int64) *LogtailConfig {
	s.LastModifyTime = &v
	return s
}

func (s *LogtailConfig) SetLogSample(v string) *LogtailConfig {
	s.LogSample = &v
	return s
}

func (s *LogtailConfig) SetOutputDetail(v *LogtailConfigOutputDetail) *LogtailConfig {
	s.OutputDetail = v
	return s
}

func (s *LogtailConfig) SetOutputType(v string) *LogtailConfig {
	s.OutputType = &v
	return s
}

type LogtailConfigOutputDetail struct {
	Endpoint     *string `json:"endpoint,omitempty" xml:"endpoint,omitempty"`
	LogstoreName *string `json:"logstoreName,omitempty" xml:"logstoreName,omitempty"`
	Region       *string `json:"region,omitempty" xml:"region,omitempty"`
}

func (s LogtailConfigOutputDetail) String() string {
	return tea.Prettify(s)
}

func (s LogtailConfigOutputDetail) GoString() string {
	return s.String()
}

func (s *LogtailConfigOutputDetail) SetEndpoint(v string) *LogtailConfigOutputDetail {
	s.Endpoint = &v
	return s
}

func (s *LogtailConfigOutputDetail) SetLogstoreName(v string) *LogtailConfigOutputDetail {
	s.LogstoreName = &v
	return s
}

func (s *LogtailConfigOutputDetail) SetRegion(v string) *LogtailConfigOutputDetail {
	s.Region = &v
	return s
}

type SavedSearch struct {
	DisplayName     *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	Logstore        *string `json:"logstore,omitempty" xml:"logstore,omitempty"`
	SavedsearchName *string `json:"savedsearchName,omitempty" xml:"savedsearchName,omitempty"`
	SearchQuery     *string `json:"searchQuery,omitempty" xml:"searchQuery,omitempty"`
	Topic           *string `json:"topic,omitempty" xml:"topic,omitempty"`
}

func (s SavedSearch) String() string {
	return tea.Prettify(s)
}

func (s SavedSearch) GoString() string {
	return s.String()
}

func (s *SavedSearch) SetDisplayName(v string) *SavedSearch {
	s.DisplayName = &v
	return s
}

func (s *SavedSearch) SetLogstore(v string) *SavedSearch {
	s.Logstore = &v
	return s
}

func (s *SavedSearch) SetSavedsearchName(v string) *SavedSearch {
	s.SavedsearchName = &v
	return s
}

func (s *SavedSearch) SetSearchQuery(v string) *SavedSearch {
	s.SearchQuery = &v
	return s
}

func (s *SavedSearch) SetTopic(v string) *SavedSearch {
	s.Topic = &v
	return s
}

type Alert struct {
	Configuration *Configuration `json:"configuration,omitempty" xml:"configuration,omitempty"`
	Description   *string        `json:"description,omitempty" xml:"description,omitempty"`
	DisplayName   *string        `json:"displayName,omitempty" xml:"displayName,omitempty"`
	Name          *string        `json:"name,omitempty" xml:"name,omitempty"`
	Schedule      *Schedule      `json:"schedule,omitempty" xml:"schedule,omitempty"`
	State         *string        `json:"state,omitempty" xml:"state,omitempty"`
	Type          *string        `json:"type,omitempty" xml:"type,omitempty"`
}

func (s Alert) String() string {
	return tea.Prettify(s)
}

func (s Alert) GoString() string {
	return s.String()
}

func (s *Alert) SetConfiguration(v *Configuration) *Alert {
	s.Configuration = v
	return s
}

func (s *Alert) SetDescription(v string) *Alert {
	s.Description = &v
	return s
}

func (s *Alert) SetDisplayName(v string) *Alert {
	s.DisplayName = &v
	return s
}

func (s *Alert) SetName(v string) *Alert {
	s.Name = &v
	return s
}

func (s *Alert) SetSchedule(v *Schedule) *Alert {
	s.Schedule = v
	return s
}

func (s *Alert) SetState(v string) *Alert {
	s.State = &v
	return s
}

func (s *Alert) SetType(v string) *Alert {
	s.Type = &v
	return s
}

type Chart struct {
	Action  map[string]interface{} `json:"action,omitempty" xml:"action,omitempty"`
	Display map[string]interface{} `json:"display,omitempty" xml:"display,omitempty"`
	Search  map[string]interface{} `json:"search,omitempty" xml:"search,omitempty"`
	Title   *string                `json:"title,omitempty" xml:"title,omitempty"`
	Type    *string                `json:"type,omitempty" xml:"type,omitempty"`
}

func (s Chart) String() string {
	return tea.Prettify(s)
}

func (s Chart) GoString() string {
	return s.String()
}

func (s *Chart) SetAction(v map[string]interface{}) *Chart {
	s.Action = v
	return s
}

func (s *Chart) SetDisplay(v map[string]interface{}) *Chart {
	s.Display = v
	return s
}

func (s *Chart) SetSearch(v map[string]interface{}) *Chart {
	s.Search = v
	return s
}

func (s *Chart) SetTitle(v string) *Chart {
	s.Title = &v
	return s
}

func (s *Chart) SetType(v string) *Chart {
	s.Type = &v
	return s
}

type Configuration struct {
	Annotations            []*ConfigurationAnnotations            `json:"annotations,omitempty" xml:"annotations,omitempty" type:"Repeated"`
	AutoAnnotation         *bool                                  `json:"autoAnnotation,omitempty" xml:"autoAnnotation,omitempty"`
	Dashboard              *string                                `json:"dashboard,omitempty" xml:"dashboard,omitempty"`
	GroupConfiguration     *ConfigurationGroupConfiguration       `json:"groupConfiguration,omitempty" xml:"groupConfiguration,omitempty" type:"Struct"`
	JoinConfigurations     []*ConfigurationJoinConfigurations     `json:"joinConfigurations,omitempty" xml:"joinConfigurations,omitempty" type:"Repeated"`
	Labels                 []*ConfigurationLabels                 `json:"labels,omitempty" xml:"labels,omitempty" type:"Repeated"`
	NoDataFire             *bool                                  `json:"noDataFire,omitempty" xml:"noDataFire,omitempty"`
	NoDataSeverity         *int64                                 `json:"noDataSeverity,omitempty" xml:"noDataSeverity,omitempty"`
	PolicyConfiguration    *ConfigurationPolicyConfiguration      `json:"policyConfiguration,omitempty" xml:"policyConfiguration,omitempty" type:"Struct"`
	QueryList              *ConfigurationQueryList                `json:"queryList,omitempty" xml:"queryList,omitempty" type:"Struct"`
	SendResolved           *bool                                  `json:"sendResolved,omitempty" xml:"sendResolved,omitempty"`
	SeverityConfigurations []*ConfigurationSeverityConfigurations `json:"severityConfigurations,omitempty" xml:"severityConfigurations,omitempty" type:"Repeated"`
	Tags                   []*string                              `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
	Threshold              *int64                                 `json:"threshold,omitempty" xml:"threshold,omitempty"`
	Type                   *string                                `json:"type,omitempty" xml:"type,omitempty"`
	Version                *string                                `json:"version,omitempty" xml:"version,omitempty"`
}

func (s Configuration) String() string {
	return tea.Prettify(s)
}

func (s Configuration) GoString() string {
	return s.String()
}

func (s *Configuration) SetAnnotations(v []*ConfigurationAnnotations) *Configuration {
	s.Annotations = v
	return s
}

func (s *Configuration) SetAutoAnnotation(v bool) *Configuration {
	s.AutoAnnotation = &v
	return s
}

func (s *Configuration) SetDashboard(v string) *Configuration {
	s.Dashboard = &v
	return s
}

func (s *Configuration) SetGroupConfiguration(v *ConfigurationGroupConfiguration) *Configuration {
	s.GroupConfiguration = v
	return s
}

func (s *Configuration) SetJoinConfigurations(v []*ConfigurationJoinConfigurations) *Configuration {
	s.JoinConfigurations = v
	return s
}

func (s *Configuration) SetLabels(v []*ConfigurationLabels) *Configuration {
	s.Labels = v
	return s
}

func (s *Configuration) SetNoDataFire(v bool) *Configuration {
	s.NoDataFire = &v
	return s
}

func (s *Configuration) SetNoDataSeverity(v int64) *Configuration {
	s.NoDataSeverity = &v
	return s
}

func (s *Configuration) SetPolicyConfiguration(v *ConfigurationPolicyConfiguration) *Configuration {
	s.PolicyConfiguration = v
	return s
}

func (s *Configuration) SetQueryList(v *ConfigurationQueryList) *Configuration {
	s.QueryList = v
	return s
}

func (s *Configuration) SetSendResolved(v bool) *Configuration {
	s.SendResolved = &v
	return s
}

func (s *Configuration) SetSeverityConfigurations(v []*ConfigurationSeverityConfigurations) *Configuration {
	s.SeverityConfigurations = v
	return s
}

func (s *Configuration) SetTags(v []*string) *Configuration {
	s.Tags = v
	return s
}

func (s *Configuration) SetThreshold(v int64) *Configuration {
	s.Threshold = &v
	return s
}

func (s *Configuration) SetType(v string) *Configuration {
	s.Type = &v
	return s
}

func (s *Configuration) SetVersion(v string) *Configuration {
	s.Version = &v
	return s
}

type ConfigurationAnnotations struct {
	Key   *string `json:"key,omitempty" xml:"key,omitempty"`
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s ConfigurationAnnotations) String() string {
	return tea.Prettify(s)
}

func (s ConfigurationAnnotations) GoString() string {
	return s.String()
}

func (s *ConfigurationAnnotations) SetKey(v string) *ConfigurationAnnotations {
	s.Key = &v
	return s
}

func (s *ConfigurationAnnotations) SetValue(v string) *ConfigurationAnnotations {
	s.Value = &v
	return s
}

type ConfigurationGroupConfiguration struct {
	Fields []*string `json:"fields,omitempty" xml:"fields,omitempty" type:"Repeated"`
	Type   *string   `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ConfigurationGroupConfiguration) String() string {
	return tea.Prettify(s)
}

func (s ConfigurationGroupConfiguration) GoString() string {
	return s.String()
}

func (s *ConfigurationGroupConfiguration) SetFields(v []*string) *ConfigurationGroupConfiguration {
	s.Fields = v
	return s
}

func (s *ConfigurationGroupConfiguration) SetType(v string) *ConfigurationGroupConfiguration {
	s.Type = &v
	return s
}

type ConfigurationJoinConfigurations struct {
	Condition *string `json:"condition,omitempty" xml:"condition,omitempty"`
	Type      *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ConfigurationJoinConfigurations) String() string {
	return tea.Prettify(s)
}

func (s ConfigurationJoinConfigurations) GoString() string {
	return s.String()
}

func (s *ConfigurationJoinConfigurations) SetCondition(v string) *ConfigurationJoinConfigurations {
	s.Condition = &v
	return s
}

func (s *ConfigurationJoinConfigurations) SetType(v string) *ConfigurationJoinConfigurations {
	s.Type = &v
	return s
}

type ConfigurationLabels struct {
	Key   *string `json:"key,omitempty" xml:"key,omitempty"`
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s ConfigurationLabels) String() string {
	return tea.Prettify(s)
}

func (s ConfigurationLabels) GoString() string {
	return s.String()
}

func (s *ConfigurationLabels) SetKey(v string) *ConfigurationLabels {
	s.Key = &v
	return s
}

func (s *ConfigurationLabels) SetValue(v string) *ConfigurationLabels {
	s.Value = &v
	return s
}

type ConfigurationPolicyConfiguration struct {
	ActionPolicyId *string `json:"actionPolicyId,omitempty" xml:"actionPolicyId,omitempty"`
	AlertPolicyId  *string `json:"alertPolicyId,omitempty" xml:"alertPolicyId,omitempty"`
	RepeatInterval *string `json:"repeatInterval,omitempty" xml:"repeatInterval,omitempty"`
	UseDefault     *bool   `json:"useDefault,omitempty" xml:"useDefault,omitempty"`
}

func (s ConfigurationPolicyConfiguration) String() string {
	return tea.Prettify(s)
}

func (s ConfigurationPolicyConfiguration) GoString() string {
	return s.String()
}

func (s *ConfigurationPolicyConfiguration) SetActionPolicyId(v string) *ConfigurationPolicyConfiguration {
	s.ActionPolicyId = &v
	return s
}

func (s *ConfigurationPolicyConfiguration) SetAlertPolicyId(v string) *ConfigurationPolicyConfiguration {
	s.AlertPolicyId = &v
	return s
}

func (s *ConfigurationPolicyConfiguration) SetRepeatInterval(v string) *ConfigurationPolicyConfiguration {
	s.RepeatInterval = &v
	return s
}

func (s *ConfigurationPolicyConfiguration) SetUseDefault(v bool) *ConfigurationPolicyConfiguration {
	s.UseDefault = &v
	return s
}

type ConfigurationQueryList struct {
	PowerSqlMode *string `json:"powerSqlMode,omitempty" xml:"powerSqlMode,omitempty"`
	Project      *string `json:"project,omitempty" xml:"project,omitempty"`
	Query        *string `json:"query,omitempty" xml:"query,omitempty"`
	Region       *string `json:"region,omitempty" xml:"region,omitempty"`
	RoleArn      *string `json:"roleArn,omitempty" xml:"roleArn,omitempty"`
	Start        *string `json:"start,omitempty" xml:"start,omitempty"`
	Store        *string `json:"store,omitempty" xml:"store,omitempty"`
	StoreType    *string `json:"storeType,omitempty" xml:"storeType,omitempty"`
	TimeSpanType *string `json:"timeSpanType,omitempty" xml:"timeSpanType,omitempty"`
}

func (s ConfigurationQueryList) String() string {
	return tea.Prettify(s)
}

func (s ConfigurationQueryList) GoString() string {
	return s.String()
}

func (s *ConfigurationQueryList) SetPowerSqlMode(v string) *ConfigurationQueryList {
	s.PowerSqlMode = &v
	return s
}

func (s *ConfigurationQueryList) SetProject(v string) *ConfigurationQueryList {
	s.Project = &v
	return s
}

func (s *ConfigurationQueryList) SetQuery(v string) *ConfigurationQueryList {
	s.Query = &v
	return s
}

func (s *ConfigurationQueryList) SetRegion(v string) *ConfigurationQueryList {
	s.Region = &v
	return s
}

func (s *ConfigurationQueryList) SetRoleArn(v string) *ConfigurationQueryList {
	s.RoleArn = &v
	return s
}

func (s *ConfigurationQueryList) SetStart(v string) *ConfigurationQueryList {
	s.Start = &v
	return s
}

func (s *ConfigurationQueryList) SetStore(v string) *ConfigurationQueryList {
	s.Store = &v
	return s
}

func (s *ConfigurationQueryList) SetStoreType(v string) *ConfigurationQueryList {
	s.StoreType = &v
	return s
}

func (s *ConfigurationQueryList) SetTimeSpanType(v string) *ConfigurationQueryList {
	s.TimeSpanType = &v
	return s
}

type ConfigurationSeverityConfigurations struct {
	EvalCondition *ConfigurationSeverityConfigurationsEvalCondition `json:"evalCondition,omitempty" xml:"evalCondition,omitempty" type:"Struct"`
	Severity      *int64                                            `json:"severity,omitempty" xml:"severity,omitempty"`
}

func (s ConfigurationSeverityConfigurations) String() string {
	return tea.Prettify(s)
}

func (s ConfigurationSeverityConfigurations) GoString() string {
	return s.String()
}

func (s *ConfigurationSeverityConfigurations) SetEvalCondition(v *ConfigurationSeverityConfigurationsEvalCondition) *ConfigurationSeverityConfigurations {
	s.EvalCondition = v
	return s
}

func (s *ConfigurationSeverityConfigurations) SetSeverity(v int64) *ConfigurationSeverityConfigurations {
	s.Severity = &v
	return s
}

type ConfigurationSeverityConfigurationsEvalCondition struct {
	Condition      *string `json:"condition,omitempty" xml:"condition,omitempty"`
	CountCondition *string `json:"countCondition,omitempty" xml:"countCondition,omitempty"`
}

func (s ConfigurationSeverityConfigurationsEvalCondition) String() string {
	return tea.Prettify(s)
}

func (s ConfigurationSeverityConfigurationsEvalCondition) GoString() string {
	return s.String()
}

func (s *ConfigurationSeverityConfigurationsEvalCondition) SetCondition(v string) *ConfigurationSeverityConfigurationsEvalCondition {
	s.Condition = &v
	return s
}

func (s *ConfigurationSeverityConfigurationsEvalCondition) SetCountCondition(v string) *ConfigurationSeverityConfigurationsEvalCondition {
	s.CountCondition = &v
	return s
}

type Dashboard struct {
	Attribute     map[string]*string `json:"attribute,omitempty" xml:"attribute,omitempty"`
	Charts        []*Chart           `json:"charts,omitempty" xml:"charts,omitempty" type:"Repeated"`
	DashboardName *string            `json:"dashboardName,omitempty" xml:"dashboardName,omitempty"`
	Description   *string            `json:"description,omitempty" xml:"description,omitempty"`
	DisplayName   *string            `json:"displayName,omitempty" xml:"displayName,omitempty"`
}

func (s Dashboard) String() string {
	return tea.Prettify(s)
}

func (s Dashboard) GoString() string {
	return s.String()
}

func (s *Dashboard) SetAttribute(v map[string]*string) *Dashboard {
	s.Attribute = v
	return s
}

func (s *Dashboard) SetCharts(v []*Chart) *Dashboard {
	s.Charts = v
	return s
}

func (s *Dashboard) SetDashboardName(v string) *Dashboard {
	s.DashboardName = &v
	return s
}

func (s *Dashboard) SetDescription(v string) *Dashboard {
	s.Description = &v
	return s
}

func (s *Dashboard) SetDisplayName(v string) *Dashboard {
	s.DisplayName = &v
	return s
}

type EtlFunctionConfig struct {
	FunctionProvider *string `json:"functionProvider,omitempty" xml:"functionProvider,omitempty"`
}

func (s EtlFunctionConfig) String() string {
	return tea.Prettify(s)
}

func (s EtlFunctionConfig) GoString() string {
	return s.String()
}

func (s *EtlFunctionConfig) SetFunctionProvider(v string) *EtlFunctionConfig {
	s.FunctionProvider = &v
	return s
}

type EtlJob struct {
	Enable            *bool                  `json:"enable,omitempty" xml:"enable,omitempty"`
	EtlJobName        *string                `json:"etlJobName,omitempty" xml:"etlJobName,omitempty"`
	FunctionConfig    *EtlJobFunctionConfig  `json:"functionConfig,omitempty" xml:"functionConfig,omitempty" type:"Struct"`
	FunctionParameter map[string]interface{} `json:"functionParameter,omitempty" xml:"functionParameter,omitempty"`
	LogConfig         *EtlJobLogConfig       `json:"logConfig,omitempty" xml:"logConfig,omitempty" type:"Struct"`
	SourceConfig      *EtlJobSourceConfig    `json:"sourceConfig,omitempty" xml:"sourceConfig,omitempty" type:"Struct"`
	TriggerConfig     *EtlJobTriggerConfig   `json:"triggerConfig,omitempty" xml:"triggerConfig,omitempty" type:"Struct"`
}

func (s EtlJob) String() string {
	return tea.Prettify(s)
}

func (s EtlJob) GoString() string {
	return s.String()
}

func (s *EtlJob) SetEnable(v bool) *EtlJob {
	s.Enable = &v
	return s
}

func (s *EtlJob) SetEtlJobName(v string) *EtlJob {
	s.EtlJobName = &v
	return s
}

func (s *EtlJob) SetFunctionConfig(v *EtlJobFunctionConfig) *EtlJob {
	s.FunctionConfig = v
	return s
}

func (s *EtlJob) SetFunctionParameter(v map[string]interface{}) *EtlJob {
	s.FunctionParameter = v
	return s
}

func (s *EtlJob) SetLogConfig(v *EtlJobLogConfig) *EtlJob {
	s.LogConfig = v
	return s
}

func (s *EtlJob) SetSourceConfig(v *EtlJobSourceConfig) *EtlJob {
	s.SourceConfig = v
	return s
}

func (s *EtlJob) SetTriggerConfig(v *EtlJobTriggerConfig) *EtlJob {
	s.TriggerConfig = v
	return s
}

type EtlJobFunctionConfig struct {
	AccountId        *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
	Endpoint         *string `json:"endpoint,omitempty" xml:"endpoint,omitempty"`
	FunctionName     *string `json:"functionName,omitempty" xml:"functionName,omitempty"`
	FunctionProvider *string `json:"functionProvider,omitempty" xml:"functionProvider,omitempty"`
	RegionName       *string `json:"regionName,omitempty" xml:"regionName,omitempty"`
	RoleArn          *string `json:"roleArn,omitempty" xml:"roleArn,omitempty"`
	ServiceName      *string `json:"serviceName,omitempty" xml:"serviceName,omitempty"`
}

func (s EtlJobFunctionConfig) String() string {
	return tea.Prettify(s)
}

func (s EtlJobFunctionConfig) GoString() string {
	return s.String()
}

func (s *EtlJobFunctionConfig) SetAccountId(v string) *EtlJobFunctionConfig {
	s.AccountId = &v
	return s
}

func (s *EtlJobFunctionConfig) SetEndpoint(v string) *EtlJobFunctionConfig {
	s.Endpoint = &v
	return s
}

func (s *EtlJobFunctionConfig) SetFunctionName(v string) *EtlJobFunctionConfig {
	s.FunctionName = &v
	return s
}

func (s *EtlJobFunctionConfig) SetFunctionProvider(v string) *EtlJobFunctionConfig {
	s.FunctionProvider = &v
	return s
}

func (s *EtlJobFunctionConfig) SetRegionName(v string) *EtlJobFunctionConfig {
	s.RegionName = &v
	return s
}

func (s *EtlJobFunctionConfig) SetRoleArn(v string) *EtlJobFunctionConfig {
	s.RoleArn = &v
	return s
}

func (s *EtlJobFunctionConfig) SetServiceName(v string) *EtlJobFunctionConfig {
	s.ServiceName = &v
	return s
}

type EtlJobLogConfig struct {
	Endpoint     *string `json:"endpoint,omitempty" xml:"endpoint,omitempty"`
	LogstoreName *string `json:"logstoreName,omitempty" xml:"logstoreName,omitempty"`
	ProjectName  *string `json:"projectName,omitempty" xml:"projectName,omitempty"`
}

func (s EtlJobLogConfig) String() string {
	return tea.Prettify(s)
}

func (s EtlJobLogConfig) GoString() string {
	return s.String()
}

func (s *EtlJobLogConfig) SetEndpoint(v string) *EtlJobLogConfig {
	s.Endpoint = &v
	return s
}

func (s *EtlJobLogConfig) SetLogstoreName(v string) *EtlJobLogConfig {
	s.LogstoreName = &v
	return s
}

func (s *EtlJobLogConfig) SetProjectName(v string) *EtlJobLogConfig {
	s.ProjectName = &v
	return s
}

type EtlJobSourceConfig struct {
	LogstoreName *string `json:"logstoreName,omitempty" xml:"logstoreName,omitempty"`
}

func (s EtlJobSourceConfig) String() string {
	return tea.Prettify(s)
}

func (s EtlJobSourceConfig) GoString() string {
	return s.String()
}

func (s *EtlJobSourceConfig) SetLogstoreName(v string) *EtlJobSourceConfig {
	s.LogstoreName = &v
	return s
}

type EtlJobTriggerConfig struct {
	MaxRetryTime     *int32  `json:"maxRetryTime,omitempty" xml:"maxRetryTime,omitempty"`
	RoleArn          *string `json:"roleArn,omitempty" xml:"roleArn,omitempty"`
	StartingPosition *string `json:"startingPosition,omitempty" xml:"startingPosition,omitempty"`
	StartingUnixtime *int64  `json:"startingUnixtime,omitempty" xml:"startingUnixtime,omitempty"`
	TriggerInterval  *int32  `json:"triggerInterval,omitempty" xml:"triggerInterval,omitempty"`
}

func (s EtlJobTriggerConfig) String() string {
	return tea.Prettify(s)
}

func (s EtlJobTriggerConfig) GoString() string {
	return s.String()
}

func (s *EtlJobTriggerConfig) SetMaxRetryTime(v int32) *EtlJobTriggerConfig {
	s.MaxRetryTime = &v
	return s
}

func (s *EtlJobTriggerConfig) SetRoleArn(v string) *EtlJobTriggerConfig {
	s.RoleArn = &v
	return s
}

func (s *EtlJobTriggerConfig) SetStartingPosition(v string) *EtlJobTriggerConfig {
	s.StartingPosition = &v
	return s
}

func (s *EtlJobTriggerConfig) SetStartingUnixtime(v int64) *EtlJobTriggerConfig {
	s.StartingUnixtime = &v
	return s
}

func (s *EtlJobTriggerConfig) SetTriggerInterval(v int32) *EtlJobTriggerConfig {
	s.TriggerInterval = &v
	return s
}

type EtlLogConfig struct {
	Endpoint     *string `json:"endpoint,omitempty" xml:"endpoint,omitempty"`
	LogstoreName *string `json:"logstoreName,omitempty" xml:"logstoreName,omitempty"`
	ProjectName  *string `json:"projectName,omitempty" xml:"projectName,omitempty"`
}

func (s EtlLogConfig) String() string {
	return tea.Prettify(s)
}

func (s EtlLogConfig) GoString() string {
	return s.String()
}

func (s *EtlLogConfig) SetEndpoint(v string) *EtlLogConfig {
	s.Endpoint = &v
	return s
}

func (s *EtlLogConfig) SetLogstoreName(v string) *EtlLogConfig {
	s.LogstoreName = &v
	return s
}

func (s *EtlLogConfig) SetProjectName(v string) *EtlLogConfig {
	s.ProjectName = &v
	return s
}

type EtlMeta struct {
	Enable       *bool   `json:"enable,omitempty" xml:"enable,omitempty"`
	EtlMetaKey   *string `json:"etlMetaKey,omitempty" xml:"etlMetaKey,omitempty"`
	EtlMetaName  *string `json:"etlMetaName,omitempty" xml:"etlMetaName,omitempty"`
	EtlMetaTag   *string `json:"etlMetaTag,omitempty" xml:"etlMetaTag,omitempty"`
	EtlMetaValue *string `json:"etlMetaValue,omitempty" xml:"etlMetaValue,omitempty"`
}

func (s EtlMeta) String() string {
	return tea.Prettify(s)
}

func (s EtlMeta) GoString() string {
	return s.String()
}

func (s *EtlMeta) SetEnable(v bool) *EtlMeta {
	s.Enable = &v
	return s
}

func (s *EtlMeta) SetEtlMetaKey(v string) *EtlMeta {
	s.EtlMetaKey = &v
	return s
}

func (s *EtlMeta) SetEtlMetaName(v string) *EtlMeta {
	s.EtlMetaName = &v
	return s
}

func (s *EtlMeta) SetEtlMetaTag(v string) *EtlMeta {
	s.EtlMetaTag = &v
	return s
}

func (s *EtlMeta) SetEtlMetaValue(v string) *EtlMeta {
	s.EtlMetaValue = &v
	return s
}

type EtlSourceConfig struct {
	LogstoreName *string `json:"logstoreName,omitempty" xml:"logstoreName,omitempty"`
}

func (s EtlSourceConfig) String() string {
	return tea.Prettify(s)
}

func (s EtlSourceConfig) GoString() string {
	return s.String()
}

func (s *EtlSourceConfig) SetLogstoreName(v string) *EtlSourceConfig {
	s.LogstoreName = &v
	return s
}

type EtlTriggerConfig struct {
	MaxRetryTime     *int32  `json:"maxRetryTime,omitempty" xml:"maxRetryTime,omitempty"`
	RoleArn          *string `json:"roleArn,omitempty" xml:"roleArn,omitempty"`
	StartingPosition *string `json:"startingPosition,omitempty" xml:"startingPosition,omitempty"`
	StartingUnixtime *int64  `json:"startingUnixtime,omitempty" xml:"startingUnixtime,omitempty"`
	TriggerInterval  *int32  `json:"triggerInterval,omitempty" xml:"triggerInterval,omitempty"`
}

func (s EtlTriggerConfig) String() string {
	return tea.Prettify(s)
}

func (s EtlTriggerConfig) GoString() string {
	return s.String()
}

func (s *EtlTriggerConfig) SetMaxRetryTime(v int32) *EtlTriggerConfig {
	s.MaxRetryTime = &v
	return s
}

func (s *EtlTriggerConfig) SetRoleArn(v string) *EtlTriggerConfig {
	s.RoleArn = &v
	return s
}

func (s *EtlTriggerConfig) SetStartingPosition(v string) *EtlTriggerConfig {
	s.StartingPosition = &v
	return s
}

func (s *EtlTriggerConfig) SetStartingUnixtime(v int64) *EtlTriggerConfig {
	s.StartingUnixtime = &v
	return s
}

func (s *EtlTriggerConfig) SetTriggerInterval(v int32) *EtlTriggerConfig {
	s.TriggerInterval = &v
	return s
}

type ExternalStore struct {
	ExternalStoreName *string                `json:"externalStoreName,omitempty" xml:"externalStoreName,omitempty"`
	Parameter         map[string]interface{} `json:"parameter,omitempty" xml:"parameter,omitempty"`
	StoreType         *string                `json:"storeType,omitempty" xml:"storeType,omitempty"`
}

func (s ExternalStore) String() string {
	return tea.Prettify(s)
}

func (s ExternalStore) GoString() string {
	return s.String()
}

func (s *ExternalStore) SetExternalStoreName(v string) *ExternalStore {
	s.ExternalStoreName = &v
	return s
}

func (s *ExternalStore) SetParameter(v map[string]interface{}) *ExternalStore {
	s.Parameter = v
	return s
}

func (s *ExternalStore) SetStoreType(v string) *ExternalStore {
	s.StoreType = &v
	return s
}

type Index struct {
	Keys               map[string]*IndexKeysValue `json:"keys,omitempty" xml:"keys,omitempty"`
	LastModifyTime     *int64                     `json:"lastModifyTime,omitempty" xml:"lastModifyTime,omitempty"`
	Line               *IndexLine                 `json:"line,omitempty" xml:"line,omitempty" type:"Struct"`
	LogReduce          *bool                      `json:"log_reduce,omitempty" xml:"log_reduce,omitempty"`
	LogReduceBlackList []*string                  `json:"log_reduce_black_list,omitempty" xml:"log_reduce_black_list,omitempty" type:"Repeated"`
	LogReduceWhiteList []*string                  `json:"log_reduce_white_list,omitempty" xml:"log_reduce_white_list,omitempty" type:"Repeated"`
	MaxTextLen         *int32                     `json:"max_text_len,omitempty" xml:"max_text_len,omitempty"`
	Ttl                *int32                     `json:"ttl,omitempty" xml:"ttl,omitempty"`
}

func (s Index) String() string {
	return tea.Prettify(s)
}

func (s Index) GoString() string {
	return s.String()
}

func (s *Index) SetKeys(v map[string]*IndexKeysValue) *Index {
	s.Keys = v
	return s
}

func (s *Index) SetLastModifyTime(v int64) *Index {
	s.LastModifyTime = &v
	return s
}

func (s *Index) SetLine(v *IndexLine) *Index {
	s.Line = v
	return s
}

func (s *Index) SetLogReduce(v bool) *Index {
	s.LogReduce = &v
	return s
}

func (s *Index) SetLogReduceBlackList(v []*string) *Index {
	s.LogReduceBlackList = v
	return s
}

func (s *Index) SetLogReduceWhiteList(v []*string) *Index {
	s.LogReduceWhiteList = v
	return s
}

func (s *Index) SetMaxTextLen(v int32) *Index {
	s.MaxTextLen = &v
	return s
}

func (s *Index) SetTtl(v int32) *Index {
	s.Ttl = &v
	return s
}

type IndexLine struct {
	CaseSensitive *bool     `json:"caseSensitive,omitempty" xml:"caseSensitive,omitempty"`
	Chn           *bool     `json:"chn,omitempty" xml:"chn,omitempty"`
	ExcludeKeys   []*string `json:"exclude_keys,omitempty" xml:"exclude_keys,omitempty" type:"Repeated"`
	IncludeKeys   []*string `json:"include_keys,omitempty" xml:"include_keys,omitempty" type:"Repeated"`
	Token         []*string `json:"token,omitempty" xml:"token,omitempty" type:"Repeated"`
}

func (s IndexLine) String() string {
	return tea.Prettify(s)
}

func (s IndexLine) GoString() string {
	return s.String()
}

func (s *IndexLine) SetCaseSensitive(v bool) *IndexLine {
	s.CaseSensitive = &v
	return s
}

func (s *IndexLine) SetChn(v bool) *IndexLine {
	s.Chn = &v
	return s
}

func (s *IndexLine) SetExcludeKeys(v []*string) *IndexLine {
	s.ExcludeKeys = v
	return s
}

func (s *IndexLine) SetIncludeKeys(v []*string) *IndexLine {
	s.IncludeKeys = v
	return s
}

func (s *IndexLine) SetToken(v []*string) *IndexLine {
	s.Token = v
	return s
}

type Logging struct {
	LoggingDetails []*LoggingLoggingDetails `json:"loggingDetails,omitempty" xml:"loggingDetails,omitempty" type:"Repeated"`
	LoggingProject *string                  `json:"loggingProject,omitempty" xml:"loggingProject,omitempty"`
}

func (s Logging) String() string {
	return tea.Prettify(s)
}

func (s Logging) GoString() string {
	return s.String()
}

func (s *Logging) SetLoggingDetails(v []*LoggingLoggingDetails) *Logging {
	s.LoggingDetails = v
	return s
}

func (s *Logging) SetLoggingProject(v string) *Logging {
	s.LoggingProject = &v
	return s
}

type LoggingLoggingDetails struct {
	Logstore *string `json:"logstore,omitempty" xml:"logstore,omitempty"`
	Type     *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s LoggingLoggingDetails) String() string {
	return tea.Prettify(s)
}

func (s LoggingLoggingDetails) GoString() string {
	return s.String()
}

func (s *LoggingLoggingDetails) SetLogstore(v string) *LoggingLoggingDetails {
	s.Logstore = &v
	return s
}

func (s *LoggingLoggingDetails) SetType(v string) *LoggingLoggingDetails {
	s.Type = &v
	return s
}

type Logstore struct {
	AppendMeta     *bool        `json:"appendMeta,omitempty" xml:"appendMeta,omitempty"`
	AutoSplit      *bool        `json:"autoSplit,omitempty" xml:"autoSplit,omitempty"`
	CreateTime     *int32       `json:"createTime,omitempty" xml:"createTime,omitempty"`
	EnableTracking *bool        `json:"enable_tracking,omitempty" xml:"enable_tracking,omitempty"`
	EncryptConf    *EncryptConf `json:"encrypt_conf,omitempty" xml:"encrypt_conf,omitempty"`
	HotTtl         *int32       `json:"hot_ttl,omitempty" xml:"hot_ttl,omitempty"`
	LastModifyTime *int32       `json:"lastModifyTime,omitempty" xml:"lastModifyTime,omitempty"`
	LogstoreName   *string      `json:"logstoreName,omitempty" xml:"logstoreName,omitempty"`
	MaxSplitShard  *int32       `json:"maxSplitShard,omitempty" xml:"maxSplitShard,omitempty"`
	Mode           *string      `json:"mode,omitempty" xml:"mode,omitempty"`
	ProductType    *string      `json:"productType,omitempty" xml:"productType,omitempty"`
	ShardCount     *int32       `json:"shardCount,omitempty" xml:"shardCount,omitempty"`
	TelemetryType  *string      `json:"telemetryType,omitempty" xml:"telemetryType,omitempty"`
	Ttl            *int32       `json:"ttl,omitempty" xml:"ttl,omitempty"`
}

func (s Logstore) String() string {
	return tea.Prettify(s)
}

func (s Logstore) GoString() string {
	return s.String()
}

func (s *Logstore) SetAppendMeta(v bool) *Logstore {
	s.AppendMeta = &v
	return s
}

func (s *Logstore) SetAutoSplit(v bool) *Logstore {
	s.AutoSplit = &v
	return s
}

func (s *Logstore) SetCreateTime(v int32) *Logstore {
	s.CreateTime = &v
	return s
}

func (s *Logstore) SetEnableTracking(v bool) *Logstore {
	s.EnableTracking = &v
	return s
}

func (s *Logstore) SetEncryptConf(v *EncryptConf) *Logstore {
	s.EncryptConf = v
	return s
}

func (s *Logstore) SetHotTtl(v int32) *Logstore {
	s.HotTtl = &v
	return s
}

func (s *Logstore) SetLastModifyTime(v int32) *Logstore {
	s.LastModifyTime = &v
	return s
}

func (s *Logstore) SetLogstoreName(v string) *Logstore {
	s.LogstoreName = &v
	return s
}

func (s *Logstore) SetMaxSplitShard(v int32) *Logstore {
	s.MaxSplitShard = &v
	return s
}

func (s *Logstore) SetMode(v string) *Logstore {
	s.Mode = &v
	return s
}

func (s *Logstore) SetProductType(v string) *Logstore {
	s.ProductType = &v
	return s
}

func (s *Logstore) SetShardCount(v int32) *Logstore {
	s.ShardCount = &v
	return s
}

func (s *Logstore) SetTelemetryType(v string) *Logstore {
	s.TelemetryType = &v
	return s
}

func (s *Logstore) SetTtl(v int32) *Logstore {
	s.Ttl = &v
	return s
}

type Machine struct {
	Ip                *string `json:"ip,omitempty" xml:"ip,omitempty"`
	LastHeartbeatTime *int64  `json:"lastHeartbeatTime,omitempty" xml:"lastHeartbeatTime,omitempty"`
	MachineUniqueid   *string `json:"machine-uniqueid,omitempty" xml:"machine-uniqueid,omitempty"`
	UserdefinedId     *string `json:"userdefined-id,omitempty" xml:"userdefined-id,omitempty"`
}

func (s Machine) String() string {
	return tea.Prettify(s)
}

func (s Machine) GoString() string {
	return s.String()
}

func (s *Machine) SetIp(v string) *Machine {
	s.Ip = &v
	return s
}

func (s *Machine) SetLastHeartbeatTime(v int64) *Machine {
	s.LastHeartbeatTime = &v
	return s
}

func (s *Machine) SetMachineUniqueid(v string) *Machine {
	s.MachineUniqueid = &v
	return s
}

func (s *Machine) SetUserdefinedId(v string) *Machine {
	s.UserdefinedId = &v
	return s
}

type MachineGroup struct {
	GroupAttribute      *MachineGroupGroupAttribute `json:"groupAttribute,omitempty" xml:"groupAttribute,omitempty" type:"Struct"`
	GroupName           *string                     `json:"groupName,omitempty" xml:"groupName,omitempty"`
	GroupType           *string                     `json:"groupType,omitempty" xml:"groupType,omitempty"`
	MachineIdentifyType *string                     `json:"machineIdentifyType,omitempty" xml:"machineIdentifyType,omitempty"`
	MachineList         []*string                   `json:"machineList,omitempty" xml:"machineList,omitempty" type:"Repeated"`
}

func (s MachineGroup) String() string {
	return tea.Prettify(s)
}

func (s MachineGroup) GoString() string {
	return s.String()
}

func (s *MachineGroup) SetGroupAttribute(v *MachineGroupGroupAttribute) *MachineGroup {
	s.GroupAttribute = v
	return s
}

func (s *MachineGroup) SetGroupName(v string) *MachineGroup {
	s.GroupName = &v
	return s
}

func (s *MachineGroup) SetGroupType(v string) *MachineGroup {
	s.GroupType = &v
	return s
}

func (s *MachineGroup) SetMachineIdentifyType(v string) *MachineGroup {
	s.MachineIdentifyType = &v
	return s
}

func (s *MachineGroup) SetMachineList(v []*string) *MachineGroup {
	s.MachineList = v
	return s
}

type MachineGroupGroupAttribute struct {
	ExternalName *string `json:"externalName,omitempty" xml:"externalName,omitempty"`
	GroupTopic   *string `json:"groupTopic,omitempty" xml:"groupTopic,omitempty"`
}

func (s MachineGroupGroupAttribute) String() string {
	return tea.Prettify(s)
}

func (s MachineGroupGroupAttribute) GoString() string {
	return s.String()
}

func (s *MachineGroupGroupAttribute) SetExternalName(v string) *MachineGroupGroupAttribute {
	s.ExternalName = &v
	return s
}

func (s *MachineGroupGroupAttribute) SetGroupTopic(v string) *MachineGroupGroupAttribute {
	s.GroupTopic = &v
	return s
}

type Project struct {
	CreateTime      *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	Description     *string `json:"description,omitempty" xml:"description,omitempty"`
	LastModifyTime  *string `json:"lastModifyTime,omitempty" xml:"lastModifyTime,omitempty"`
	Owner           *string `json:"owner,omitempty" xml:"owner,omitempty"`
	ProjectName     *string `json:"projectName,omitempty" xml:"projectName,omitempty"`
	Region          *string `json:"region,omitempty" xml:"region,omitempty"`
	ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	Status          *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s Project) String() string {
	return tea.Prettify(s)
}

func (s Project) GoString() string {
	return s.String()
}

func (s *Project) SetCreateTime(v string) *Project {
	s.CreateTime = &v
	return s
}

func (s *Project) SetDescription(v string) *Project {
	s.Description = &v
	return s
}

func (s *Project) SetLastModifyTime(v string) *Project {
	s.LastModifyTime = &v
	return s
}

func (s *Project) SetOwner(v string) *Project {
	s.Owner = &v
	return s
}

func (s *Project) SetProjectName(v string) *Project {
	s.ProjectName = &v
	return s
}

func (s *Project) SetRegion(v string) *Project {
	s.Region = &v
	return s
}

func (s *Project) SetResourceGroupId(v string) *Project {
	s.ResourceGroupId = &v
	return s
}

func (s *Project) SetStatus(v string) *Project {
	s.Status = &v
	return s
}

type Schedule struct {
	CronExpression *string `json:"cronExpression,omitempty" xml:"cronExpression,omitempty"`
	DayOfWeek      *int64  `json:"dayOfWeek,omitempty" xml:"dayOfWeek,omitempty"`
	Hour           *int64  `json:"hour,omitempty" xml:"hour,omitempty"`
	Interval       *string `json:"interval,omitempty" xml:"interval,omitempty"`
	RunImmediately *bool   `json:"runImmediately,omitempty" xml:"runImmediately,omitempty"`
	Type           *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s Schedule) String() string {
	return tea.Prettify(s)
}

func (s Schedule) GoString() string {
	return s.String()
}

func (s *Schedule) SetCronExpression(v string) *Schedule {
	s.CronExpression = &v
	return s
}

func (s *Schedule) SetDayOfWeek(v int64) *Schedule {
	s.DayOfWeek = &v
	return s
}

func (s *Schedule) SetHour(v int64) *Schedule {
	s.Hour = &v
	return s
}

func (s *Schedule) SetInterval(v string) *Schedule {
	s.Interval = &v
	return s
}

func (s *Schedule) SetRunImmediately(v bool) *Schedule {
	s.RunImmediately = &v
	return s
}

func (s *Schedule) SetType(v string) *Schedule {
	s.Type = &v
	return s
}

type Shard struct {
	CreateTime        *int32  `json:"createTime,omitempty" xml:"createTime,omitempty"`
	ExclusiveEndKey   *string `json:"exclusiveEndKey,omitempty" xml:"exclusiveEndKey,omitempty"`
	InclusiveBeginKey *string `json:"inclusiveBeginKey,omitempty" xml:"inclusiveBeginKey,omitempty"`
	ShardID           *int32  `json:"shardID,omitempty" xml:"shardID,omitempty"`
	Status            *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s Shard) String() string {
	return tea.Prettify(s)
}

func (s Shard) GoString() string {
	return s.String()
}

func (s *Shard) SetCreateTime(v int32) *Shard {
	s.CreateTime = &v
	return s
}

func (s *Shard) SetExclusiveEndKey(v string) *Shard {
	s.ExclusiveEndKey = &v
	return s
}

func (s *Shard) SetInclusiveBeginKey(v string) *Shard {
	s.InclusiveBeginKey = &v
	return s
}

func (s *Shard) SetShardID(v int32) *Shard {
	s.ShardID = &v
	return s
}

func (s *Shard) SetStatus(v string) *Shard {
	s.Status = &v
	return s
}

type IndexKeysValue struct {
	Chn           *bool     `json:"chn,omitempty" xml:"chn,omitempty"`
	CaseSensitive *bool     `json:"caseSensitive,omitempty" xml:"caseSensitive,omitempty"`
	Token         []*string `json:"token,omitempty" xml:"token,omitempty" type:"Repeated"`
	Alias         *string   `json:"alias,omitempty" xml:"alias,omitempty"`
	Type          *string   `json:"type,omitempty" xml:"type,omitempty"`
	DocValue      *bool     `json:"doc_value,omitempty" xml:"doc_value,omitempty"`
}

func (s IndexKeysValue) String() string {
	return tea.Prettify(s)
}

func (s IndexKeysValue) GoString() string {
	return s.String()
}

func (s *IndexKeysValue) SetChn(v bool) *IndexKeysValue {
	s.Chn = &v
	return s
}

func (s *IndexKeysValue) SetCaseSensitive(v bool) *IndexKeysValue {
	s.CaseSensitive = &v
	return s
}

func (s *IndexKeysValue) SetToken(v []*string) *IndexKeysValue {
	s.Token = v
	return s
}

func (s *IndexKeysValue) SetAlias(v string) *IndexKeysValue {
	s.Alias = &v
	return s
}

func (s *IndexKeysValue) SetType(v string) *IndexKeysValue {
	s.Type = &v
	return s
}

func (s *IndexKeysValue) SetDocValue(v bool) *IndexKeysValue {
	s.DocValue = &v
	return s
}

type KeysValue struct {
	CaseSensitive *bool     `json:"caseSensitive,omitempty" xml:"caseSensitive,omitempty"`
	Chn           *bool     `json:"chn,omitempty" xml:"chn,omitempty"`
	Type          *string   `json:"type,omitempty" xml:"type,omitempty"`
	Alias         *string   `json:"alias,omitempty" xml:"alias,omitempty"`
	Token         []*string `json:"token,omitempty" xml:"token,omitempty" type:"Repeated"`
	DocValue      *bool     `json:"doc_value,omitempty" xml:"doc_value,omitempty"`
}

func (s KeysValue) String() string {
	return tea.Prettify(s)
}

func (s KeysValue) GoString() string {
	return s.String()
}

func (s *KeysValue) SetCaseSensitive(v bool) *KeysValue {
	s.CaseSensitive = &v
	return s
}

func (s *KeysValue) SetChn(v bool) *KeysValue {
	s.Chn = &v
	return s
}

func (s *KeysValue) SetType(v string) *KeysValue {
	s.Type = &v
	return s
}

func (s *KeysValue) SetAlias(v string) *KeysValue {
	s.Alias = &v
	return s
}

func (s *KeysValue) SetToken(v []*string) *KeysValue {
	s.Token = v
	return s
}

func (s *KeysValue) SetDocValue(v bool) *KeysValue {
	s.DocValue = &v
	return s
}

type ApplyConfigToMachineGroupResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
}

func (s ApplyConfigToMachineGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s ApplyConfigToMachineGroupResponse) GoString() string {
	return s.String()
}

func (s *ApplyConfigToMachineGroupResponse) SetHeaders(v map[string]*string) *ApplyConfigToMachineGroupResponse {
	s.Headers = v
	return s
}

func (s *ApplyConfigToMachineGroupResponse) SetStatusCode(v int32) *ApplyConfigToMachineGroupResponse {
	s.StatusCode = &v
	return s
}

type ChangeResourceGroupRequest struct {
	ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	ResourceId      *string `json:"resourceId,omitempty" xml:"resourceId,omitempty"`
	ResourceType    *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
}

func (s ChangeResourceGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s ChangeResourceGroupRequest) GoString() string {
	return s.String()
}

func (s *ChangeResourceGroupRequest) SetResourceGroupId(v string) *ChangeResourceGroupRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ChangeResourceGroupRequest) SetResourceId(v string) *ChangeResourceGroupRequest {
	s.ResourceId = &v
	return s
}

func (s *ChangeResourceGroupRequest) SetResourceType(v string) *ChangeResourceGroupRequest {
	s.ResourceType = &v
	return s
}

type ChangeResourceGroupResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
}

func (s ChangeResourceGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s ChangeResourceGroupResponse) GoString() string {
	return s.String()
}

func (s *ChangeResourceGroupResponse) SetHeaders(v map[string]*string) *ChangeResourceGroupResponse {
	s.Headers = v
	return s
}

func (s *ChangeResourceGroupResponse) SetStatusCode(v int32) *ChangeResourceGroupResponse {
	s.StatusCode = &v
	return s
}

type CreateAlertRequest struct {
	Body *Alert `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAlertRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAlertRequest) GoString() string {
	return s.String()
}

func (s *CreateAlertRequest) SetBody(v *Alert) *CreateAlertRequest {
	s.Body = v
	return s
}

type CreateAlertResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
}

func (s CreateAlertResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateAlertResponse) GoString() string {
	return s.String()
}

func (s *CreateAlertResponse) SetHeaders(v map[string]*string) *CreateAlertResponse {
	s.Headers = v
	return s
}

func (s *CreateAlertResponse) SetStatusCode(v int32) *CreateAlertResponse {
	s.StatusCode = &v
	return s
}

type CreateConsumerGroupRequest struct {
	ConsumerGroup *string `json:"consumerGroup,omitempty" xml:"consumerGroup,omitempty"`
	Order         *bool   `json:"order,omitempty" xml:"order,omitempty"`
	Timeout       *int32  `json:"timeout,omitempty" xml:"timeout,omitempty"`
}

func (s CreateConsumerGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateConsumerGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateConsumerGroupRequest) SetConsumerGroup(v string) *CreateConsumerGroupRequest {
	s.ConsumerGroup = &v
	return s
}

func (s *CreateConsumerGroupRequest) SetOrder(v bool) *CreateConsumerGroupRequest {
	s.Order = &v
	return s
}

func (s *CreateConsumerGroupRequest) SetTimeout(v int32) *CreateConsumerGroupRequest {
	s.Timeout = &v
	return s
}

type CreateConsumerGroupResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
}

func (s CreateConsumerGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateConsumerGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateConsumerGroupResponse) SetHeaders(v map[string]*string) *CreateConsumerGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateConsumerGroupResponse) SetStatusCode(v int32) *CreateConsumerGroupResponse {
	s.StatusCode = &v
	return s
}

type CreateDomainRequest struct {
	DomainName *string `json:"domainName,omitempty" xml:"domainName,omitempty"`
}

func (s CreateDomainRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDomainRequest) GoString() string {
	return s.String()
}

func (s *CreateDomainRequest) SetDomainName(v string) *CreateDomainRequest {
	s.DomainName = &v
	return s
}

type CreateDomainResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
}

func (s CreateDomainResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDomainResponse) GoString() string {
	return s.String()
}

func (s *CreateDomainResponse) SetHeaders(v map[string]*string) *CreateDomainResponse {
	s.Headers = v
	return s
}

func (s *CreateDomainResponse) SetStatusCode(v int32) *CreateDomainResponse {
	s.StatusCode = &v
	return s
}

type CreateETLJobRequest struct {
	Enable            *bool              `json:"enable,omitempty" xml:"enable,omitempty"`
	FunctionConfig    *EtlFunctionConfig `json:"functionConfig,omitempty" xml:"functionConfig,omitempty"`
	FunctionParameter *string            `json:"functionParameter,omitempty" xml:"functionParameter,omitempty"`
	JobName           *string            `json:"jobName,omitempty" xml:"jobName,omitempty"`
	LogConfig         *EtlLogConfig      `json:"logConfig,omitempty" xml:"logConfig,omitempty"`
	SourceConfig      *EtlSourceConfig   `json:"sourceConfig,omitempty" xml:"sourceConfig,omitempty"`
	TriggerConfig     *EtlTriggerConfig  `json:"triggerConfig,omitempty" xml:"triggerConfig,omitempty"`
}

func (s CreateETLJobRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateETLJobRequest) GoString() string {
	return s.String()
}

func (s *CreateETLJobRequest) SetEnable(v bool) *CreateETLJobRequest {
	s.Enable = &v
	return s
}

func (s *CreateETLJobRequest) SetFunctionConfig(v *EtlFunctionConfig) *CreateETLJobRequest {
	s.FunctionConfig = v
	return s
}

func (s *CreateETLJobRequest) SetFunctionParameter(v string) *CreateETLJobRequest {
	s.FunctionParameter = &v
	return s
}

func (s *CreateETLJobRequest) SetJobName(v string) *CreateETLJobRequest {
	s.JobName = &v
	return s
}

func (s *CreateETLJobRequest) SetLogConfig(v *EtlLogConfig) *CreateETLJobRequest {
	s.LogConfig = v
	return s
}

func (s *CreateETLJobRequest) SetSourceConfig(v *EtlSourceConfig) *CreateETLJobRequest {
	s.SourceConfig = v
	return s
}

func (s *CreateETLJobRequest) SetTriggerConfig(v *EtlTriggerConfig) *CreateETLJobRequest {
	s.TriggerConfig = v
	return s
}

type CreateETLJobResponseBody struct {
	Code      *int32                   `json:"code,omitempty" xml:"code,omitempty"`
	Data      []map[string]interface{} `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	Message   *string                  `json:"message,omitempty" xml:"message,omitempty"`
	RequestId *string                  `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *string                  `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CreateETLJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateETLJobResponseBody) GoString() string {
	return s.String()
}

func (s *CreateETLJobResponseBody) SetCode(v int32) *CreateETLJobResponseBody {
	s.Code = &v
	return s
}

func (s *CreateETLJobResponseBody) SetData(v []map[string]interface{}) *CreateETLJobResponseBody {
	s.Data = v
	return s
}

func (s *CreateETLJobResponseBody) SetMessage(v string) *CreateETLJobResponseBody {
	s.Message = &v
	return s
}

func (s *CreateETLJobResponseBody) SetRequestId(v string) *CreateETLJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateETLJobResponseBody) SetSuccess(v string) *CreateETLJobResponseBody {
	s.Success = &v
	return s
}

type CreateETLJobResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateETLJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateETLJobResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateETLJobResponse) GoString() string {
	return s.String()
}

func (s *CreateETLJobResponse) SetHeaders(v map[string]*string) *CreateETLJobResponse {
	s.Headers = v
	return s
}

func (s *CreateETLJobResponse) SetStatusCode(v int32) *CreateETLJobResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateETLJobResponse) SetBody(v *CreateETLJobResponseBody) *CreateETLJobResponse {
	s.Body = v
	return s
}

type CreateIndexRequest struct {
	Keys               map[string]*KeysValue   `json:"keys,omitempty" xml:"keys,omitempty"`
	Line               *CreateIndexRequestLine `json:"line,omitempty" xml:"line,omitempty" type:"Struct"`
	LogReduce          *bool                   `json:"log_reduce,omitempty" xml:"log_reduce,omitempty"`
	LogReduceBlackList []*string               `json:"log_reduce_black_list,omitempty" xml:"log_reduce_black_list,omitempty" type:"Repeated"`
	LogReduceWhiteList []*string               `json:"log_reduce_white_list,omitempty" xml:"log_reduce_white_list,omitempty" type:"Repeated"`
	MaxTextLen         *int32                  `json:"max_text_len,omitempty" xml:"max_text_len,omitempty"`
	Ttl                *int32                  `json:"ttl,omitempty" xml:"ttl,omitempty"`
}

func (s CreateIndexRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateIndexRequest) GoString() string {
	return s.String()
}

func (s *CreateIndexRequest) SetKeys(v map[string]*KeysValue) *CreateIndexRequest {
	s.Keys = v
	return s
}

func (s *CreateIndexRequest) SetLine(v *CreateIndexRequestLine) *CreateIndexRequest {
	s.Line = v
	return s
}

func (s *CreateIndexRequest) SetLogReduce(v bool) *CreateIndexRequest {
	s.LogReduce = &v
	return s
}

func (s *CreateIndexRequest) SetLogReduceBlackList(v []*string) *CreateIndexRequest {
	s.LogReduceBlackList = v
	return s
}

func (s *CreateIndexRequest) SetLogReduceWhiteList(v []*string) *CreateIndexRequest {
	s.LogReduceWhiteList = v
	return s
}

func (s *CreateIndexRequest) SetMaxTextLen(v int32) *CreateIndexRequest {
	s.MaxTextLen = &v
	return s
}

func (s *CreateIndexRequest) SetTtl(v int32) *CreateIndexRequest {
	s.Ttl = &v
	return s
}

type CreateIndexRequestLine struct {
	CaseSensitive *bool     `json:"caseSensitive,omitempty" xml:"caseSensitive,omitempty"`
	Chn           *bool     `json:"chn,omitempty" xml:"chn,omitempty"`
	ExcludeKeys   []*string `json:"exclude_keys,omitempty" xml:"exclude_keys,omitempty" type:"Repeated"`
	IncludeKeys   []*string `json:"include_keys,omitempty" xml:"include_keys,omitempty" type:"Repeated"`
	Token         []*string `json:"token,omitempty" xml:"token,omitempty" type:"Repeated"`
}

func (s CreateIndexRequestLine) String() string {
	return tea.Prettify(s)
}

func (s CreateIndexRequestLine) GoString() string {
	return s.String()
}

func (s *CreateIndexRequestLine) SetCaseSensitive(v bool) *CreateIndexRequestLine {
	s.CaseSensitive = &v
	return s
}

func (s *CreateIndexRequestLine) SetChn(v bool) *CreateIndexRequestLine {
	s.Chn = &v
	return s
}

func (s *CreateIndexRequestLine) SetExcludeKeys(v []*string) *CreateIndexRequestLine {
	s.ExcludeKeys = v
	return s
}

func (s *CreateIndexRequestLine) SetIncludeKeys(v []*string) *CreateIndexRequestLine {
	s.IncludeKeys = v
	return s
}

func (s *CreateIndexRequestLine) SetToken(v []*string) *CreateIndexRequestLine {
	s.Token = v
	return s
}

type CreateIndexResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
}

func (s CreateIndexResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateIndexResponse) GoString() string {
	return s.String()
}

func (s *CreateIndexResponse) SetHeaders(v map[string]*string) *CreateIndexResponse {
	s.Headers = v
	return s
}

func (s *CreateIndexResponse) SetStatusCode(v int32) *CreateIndexResponse {
	s.StatusCode = &v
	return s
}

type CreateJobRequest struct {
	Enable            *bool              `json:"enable,omitempty" xml:"enable,omitempty"`
	FunctionConfig    *EtlFunctionConfig `json:"functionConfig,omitempty" xml:"functionConfig,omitempty"`
	FunctionParameter *string            `json:"functionParameter,omitempty" xml:"functionParameter,omitempty"`
	JobName           *string            `json:"jobName,omitempty" xml:"jobName,omitempty"`
	LogConfig         *EtlLogConfig      `json:"logConfig,omitempty" xml:"logConfig,omitempty"`
	SourceConfig      *EtlSourceConfig   `json:"sourceConfig,omitempty" xml:"sourceConfig,omitempty"`
	TriggerConfig     *EtlTriggerConfig  `json:"triggerConfig,omitempty" xml:"triggerConfig,omitempty"`
}

func (s CreateJobRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateJobRequest) GoString() string {
	return s.String()
}

func (s *CreateJobRequest) SetEnable(v bool) *CreateJobRequest {
	s.Enable = &v
	return s
}

func (s *CreateJobRequest) SetFunctionConfig(v *EtlFunctionConfig) *CreateJobRequest {
	s.FunctionConfig = v
	return s
}

func (s *CreateJobRequest) SetFunctionParameter(v string) *CreateJobRequest {
	s.FunctionParameter = &v
	return s
}

func (s *CreateJobRequest) SetJobName(v string) *CreateJobRequest {
	s.JobName = &v
	return s
}

func (s *CreateJobRequest) SetLogConfig(v *EtlLogConfig) *CreateJobRequest {
	s.LogConfig = v
	return s
}

func (s *CreateJobRequest) SetSourceConfig(v *EtlSourceConfig) *CreateJobRequest {
	s.SourceConfig = v
	return s
}

func (s *CreateJobRequest) SetTriggerConfig(v *EtlTriggerConfig) *CreateJobRequest {
	s.TriggerConfig = v
	return s
}

type CreateJobResponseBody struct {
	Code      *int32  `json:"code,omitempty" xml:"code,omitempty"`
	Data      *string `json:"data,omitempty" xml:"data,omitempty"`
	Message   *string `json:"message,omitempty" xml:"message,omitempty"`
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *string `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CreateJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateJobResponseBody) GoString() string {
	return s.String()
}

func (s *CreateJobResponseBody) SetCode(v int32) *CreateJobResponseBody {
	s.Code = &v
	return s
}

func (s *CreateJobResponseBody) SetData(v string) *CreateJobResponseBody {
	s.Data = &v
	return s
}

func (s *CreateJobResponseBody) SetMessage(v string) *CreateJobResponseBody {
	s.Message = &v
	return s
}

func (s *CreateJobResponseBody) SetRequestId(v string) *CreateJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateJobResponseBody) SetSuccess(v string) *CreateJobResponseBody {
	s.Success = &v
	return s
}

type CreateJobResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateJobResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateJobResponse) GoString() string {
	return s.String()
}

func (s *CreateJobResponse) SetHeaders(v map[string]*string) *CreateJobResponse {
	s.Headers = v
	return s
}

func (s *CreateJobResponse) SetStatusCode(v int32) *CreateJobResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateJobResponse) SetBody(v *CreateJobResponseBody) *CreateJobResponse {
	s.Body = v
	return s
}

type CreateLogStoreRequest struct {
	AppendMeta     *bool        `json:"appendMeta,omitempty" xml:"appendMeta,omitempty"`
	AutoSplit      *bool        `json:"autoSplit,omitempty" xml:"autoSplit,omitempty"`
	EnableTracking *bool        `json:"enable_tracking,omitempty" xml:"enable_tracking,omitempty"`
	EncryptConf    *EncryptConf `json:"encrypt_conf,omitempty" xml:"encrypt_conf,omitempty"`
	HotTtl         *int32       `json:"hot_ttl,omitempty" xml:"hot_ttl,omitempty"`
	LogstoreName   *string      `json:"logstoreName,omitempty" xml:"logstoreName,omitempty"`
	MaxSplitShard  *int32       `json:"maxSplitShard,omitempty" xml:"maxSplitShard,omitempty"`
	Mode           *string      `json:"mode,omitempty" xml:"mode,omitempty"`
	ShardCount     *int32       `json:"shardCount,omitempty" xml:"shardCount,omitempty"`
	TelemetryType  *string      `json:"telemetryType,omitempty" xml:"telemetryType,omitempty"`
	Ttl            *int32       `json:"ttl,omitempty" xml:"ttl,omitempty"`
}

func (s CreateLogStoreRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateLogStoreRequest) GoString() string {
	return s.String()
}

func (s *CreateLogStoreRequest) SetAppendMeta(v bool) *CreateLogStoreRequest {
	s.AppendMeta = &v
	return s
}

func (s *CreateLogStoreRequest) SetAutoSplit(v bool) *CreateLogStoreRequest {
	s.AutoSplit = &v
	return s
}

func (s *CreateLogStoreRequest) SetEnableTracking(v bool) *CreateLogStoreRequest {
	s.EnableTracking = &v
	return s
}

func (s *CreateLogStoreRequest) SetEncryptConf(v *EncryptConf) *CreateLogStoreRequest {
	s.EncryptConf = v
	return s
}

func (s *CreateLogStoreRequest) SetHotTtl(v int32) *CreateLogStoreRequest {
	s.HotTtl = &v
	return s
}

func (s *CreateLogStoreRequest) SetLogstoreName(v string) *CreateLogStoreRequest {
	s.LogstoreName = &v
	return s
}

func (s *CreateLogStoreRequest) SetMaxSplitShard(v int32) *CreateLogStoreRequest {
	s.MaxSplitShard = &v
	return s
}

func (s *CreateLogStoreRequest) SetMode(v string) *CreateLogStoreRequest {
	s.Mode = &v
	return s
}

func (s *CreateLogStoreRequest) SetShardCount(v int32) *CreateLogStoreRequest {
	s.ShardCount = &v
	return s
}

func (s *CreateLogStoreRequest) SetTelemetryType(v string) *CreateLogStoreRequest {
	s.TelemetryType = &v
	return s
}

func (s *CreateLogStoreRequest) SetTtl(v int32) *CreateLogStoreRequest {
	s.Ttl = &v
	return s
}

type CreateLogStoreResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
}

func (s CreateLogStoreResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateLogStoreResponse) GoString() string {
	return s.String()
}

func (s *CreateLogStoreResponse) SetHeaders(v map[string]*string) *CreateLogStoreResponse {
	s.Headers = v
	return s
}

func (s *CreateLogStoreResponse) SetStatusCode(v int32) *CreateLogStoreResponse {
	s.StatusCode = &v
	return s
}

type CreateLoggingRequest struct {
	LoggingDetails []*CreateLoggingRequestLoggingDetails `json:"loggingDetails,omitempty" xml:"loggingDetails,omitempty" type:"Repeated"`
	LoggingProject *string                               `json:"loggingProject,omitempty" xml:"loggingProject,omitempty"`
}

func (s CreateLoggingRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateLoggingRequest) GoString() string {
	return s.String()
}

func (s *CreateLoggingRequest) SetLoggingDetails(v []*CreateLoggingRequestLoggingDetails) *CreateLoggingRequest {
	s.LoggingDetails = v
	return s
}

func (s *CreateLoggingRequest) SetLoggingProject(v string) *CreateLoggingRequest {
	s.LoggingProject = &v
	return s
}

type CreateLoggingRequestLoggingDetails struct {
	Logstore *string `json:"logstore,omitempty" xml:"logstore,omitempty"`
	Type     *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s CreateLoggingRequestLoggingDetails) String() string {
	return tea.Prettify(s)
}

func (s CreateLoggingRequestLoggingDetails) GoString() string {
	return s.String()
}

func (s *CreateLoggingRequestLoggingDetails) SetLogstore(v string) *CreateLoggingRequestLoggingDetails {
	s.Logstore = &v
	return s
}

func (s *CreateLoggingRequestLoggingDetails) SetType(v string) *CreateLoggingRequestLoggingDetails {
	s.Type = &v
	return s
}

type CreateLoggingResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
}

func (s CreateLoggingResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateLoggingResponse) GoString() string {
	return s.String()
}

func (s *CreateLoggingResponse) SetHeaders(v map[string]*string) *CreateLoggingResponse {
	s.Headers = v
	return s
}

func (s *CreateLoggingResponse) SetStatusCode(v int32) *CreateLoggingResponse {
	s.StatusCode = &v
	return s
}

type CreateMachineGroupRequest struct {
	GroupAttribute      *CreateMachineGroupRequestGroupAttribute `json:"groupAttribute,omitempty" xml:"groupAttribute,omitempty" type:"Struct"`
	GroupName           *string                                  `json:"groupName,omitempty" xml:"groupName,omitempty"`
	GroupType           *string                                  `json:"groupType,omitempty" xml:"groupType,omitempty"`
	MachineIdentifyType *string                                  `json:"machineIdentifyType,omitempty" xml:"machineIdentifyType,omitempty"`
	MachineList         []*string                                `json:"machineList,omitempty" xml:"machineList,omitempty" type:"Repeated"`
}

func (s CreateMachineGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateMachineGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateMachineGroupRequest) SetGroupAttribute(v *CreateMachineGroupRequestGroupAttribute) *CreateMachineGroupRequest {
	s.GroupAttribute = v
	return s
}

func (s *CreateMachineGroupRequest) SetGroupName(v string) *CreateMachineGroupRequest {
	s.GroupName = &v
	return s
}

func (s *CreateMachineGroupRequest) SetGroupType(v string) *CreateMachineGroupRequest {
	s.GroupType = &v
	return s
}

func (s *CreateMachineGroupRequest) SetMachineIdentifyType(v string) *CreateMachineGroupRequest {
	s.MachineIdentifyType = &v
	return s
}

func (s *CreateMachineGroupRequest) SetMachineList(v []*string) *CreateMachineGroupRequest {
	s.MachineList = v
	return s
}

type CreateMachineGroupRequestGroupAttribute struct {
	ExternalName *string `json:"externalName,omitempty" xml:"externalName,omitempty"`
	GroupTopic   *string `json:"groupTopic,omitempty" xml:"groupTopic,omitempty"`
}

func (s CreateMachineGroupRequestGroupAttribute) String() string {
	return tea.Prettify(s)
}

func (s CreateMachineGroupRequestGroupAttribute) GoString() string {
	return s.String()
}

func (s *CreateMachineGroupRequestGroupAttribute) SetExternalName(v string) *CreateMachineGroupRequestGroupAttribute {
	s.ExternalName = &v
	return s
}

func (s *CreateMachineGroupRequestGroupAttribute) SetGroupTopic(v string) *CreateMachineGroupRequestGroupAttribute {
	s.GroupTopic = &v
	return s
}

type CreateMachineGroupResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
}

func (s CreateMachineGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateMachineGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateMachineGroupResponse) SetHeaders(v map[string]*string) *CreateMachineGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateMachineGroupResponse) SetStatusCode(v int32) *CreateMachineGroupResponse {
	s.StatusCode = &v
	return s
}

type CreateOdpsShipperRequest struct {
	ShipperName         *string                                      `json:"shipperName,omitempty" xml:"shipperName,omitempty"`
	TargetConfiguration *CreateOdpsShipperRequestTargetConfiguration `json:"targetConfiguration,omitempty" xml:"targetConfiguration,omitempty" type:"Struct"`
	TargetType          *string                                      `json:"targetType,omitempty" xml:"targetType,omitempty"`
}

func (s CreateOdpsShipperRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateOdpsShipperRequest) GoString() string {
	return s.String()
}

func (s *CreateOdpsShipperRequest) SetShipperName(v string) *CreateOdpsShipperRequest {
	s.ShipperName = &v
	return s
}

func (s *CreateOdpsShipperRequest) SetTargetConfiguration(v *CreateOdpsShipperRequestTargetConfiguration) *CreateOdpsShipperRequest {
	s.TargetConfiguration = v
	return s
}

func (s *CreateOdpsShipperRequest) SetTargetType(v string) *CreateOdpsShipperRequest {
	s.TargetType = &v
	return s
}

type CreateOdpsShipperRequestTargetConfiguration struct {
	BufferInterval      *int32    `json:"bufferInterval,omitempty" xml:"bufferInterval,omitempty"`
	Enable              *bool     `json:"enable,omitempty" xml:"enable,omitempty"`
	Fields              []*string `json:"fields,omitempty" xml:"fields,omitempty" type:"Repeated"`
	OdpsEndpoint        *string   `json:"odpsEndpoint,omitempty" xml:"odpsEndpoint,omitempty"`
	OdpsProject         *string   `json:"odpsProject,omitempty" xml:"odpsProject,omitempty"`
	OdpsTable           *string   `json:"odpsTable,omitempty" xml:"odpsTable,omitempty"`
	PartitionColumn     []*string `json:"partitionColumn,omitempty" xml:"partitionColumn,omitempty" type:"Repeated"`
	PartitionTimeFormat *string   `json:"partitionTimeFormat,omitempty" xml:"partitionTimeFormat,omitempty"`
}

func (s CreateOdpsShipperRequestTargetConfiguration) String() string {
	return tea.Prettify(s)
}

func (s CreateOdpsShipperRequestTargetConfiguration) GoString() string {
	return s.String()
}

func (s *CreateOdpsShipperRequestTargetConfiguration) SetBufferInterval(v int32) *CreateOdpsShipperRequestTargetConfiguration {
	s.BufferInterval = &v
	return s
}

func (s *CreateOdpsShipperRequestTargetConfiguration) SetEnable(v bool) *CreateOdpsShipperRequestTargetConfiguration {
	s.Enable = &v
	return s
}

func (s *CreateOdpsShipperRequestTargetConfiguration) SetFields(v []*string) *CreateOdpsShipperRequestTargetConfiguration {
	s.Fields = v
	return s
}

func (s *CreateOdpsShipperRequestTargetConfiguration) SetOdpsEndpoint(v string) *CreateOdpsShipperRequestTargetConfiguration {
	s.OdpsEndpoint = &v
	return s
}

func (s *CreateOdpsShipperRequestTargetConfiguration) SetOdpsProject(v string) *CreateOdpsShipperRequestTargetConfiguration {
	s.OdpsProject = &v
	return s
}

func (s *CreateOdpsShipperRequestTargetConfiguration) SetOdpsTable(v string) *CreateOdpsShipperRequestTargetConfiguration {
	s.OdpsTable = &v
	return s
}

func (s *CreateOdpsShipperRequestTargetConfiguration) SetPartitionColumn(v []*string) *CreateOdpsShipperRequestTargetConfiguration {
	s.PartitionColumn = v
	return s
}

func (s *CreateOdpsShipperRequestTargetConfiguration) SetPartitionTimeFormat(v string) *CreateOdpsShipperRequestTargetConfiguration {
	s.PartitionTimeFormat = &v
	return s
}

type CreateOdpsShipperResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
}

func (s CreateOdpsShipperResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateOdpsShipperResponse) GoString() string {
	return s.String()
}

func (s *CreateOdpsShipperResponse) SetHeaders(v map[string]*string) *CreateOdpsShipperResponse {
	s.Headers = v
	return s
}

func (s *CreateOdpsShipperResponse) SetStatusCode(v int32) *CreateOdpsShipperResponse {
	s.StatusCode = &v
	return s
}

type CreateOssExternalStoreRequest struct {
	ExternalStoreName *string                                 `json:"externalStoreName,omitempty" xml:"externalStoreName,omitempty"`
	Parameter         *CreateOssExternalStoreRequestParameter `json:"parameter,omitempty" xml:"parameter,omitempty" type:"Struct"`
	StoreType         *string                                 `json:"storeType,omitempty" xml:"storeType,omitempty"`
}

func (s CreateOssExternalStoreRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateOssExternalStoreRequest) GoString() string {
	return s.String()
}

func (s *CreateOssExternalStoreRequest) SetExternalStoreName(v string) *CreateOssExternalStoreRequest {
	s.ExternalStoreName = &v
	return s
}

func (s *CreateOssExternalStoreRequest) SetParameter(v *CreateOssExternalStoreRequestParameter) *CreateOssExternalStoreRequest {
	s.Parameter = v
	return s
}

func (s *CreateOssExternalStoreRequest) SetStoreType(v string) *CreateOssExternalStoreRequest {
	s.StoreType = &v
	return s
}

type CreateOssExternalStoreRequestParameter struct {
	Accessid  *string                                          `json:"accessid,omitempty" xml:"accessid,omitempty"`
	Accesskey *string                                          `json:"accesskey,omitempty" xml:"accesskey,omitempty"`
	Bucket    *string                                          `json:"bucket,omitempty" xml:"bucket,omitempty"`
	Columns   []*CreateOssExternalStoreRequestParameterColumns `json:"columns,omitempty" xml:"columns,omitempty" type:"Repeated"`
	Endpoint  *string                                          `json:"endpoint,omitempty" xml:"endpoint,omitempty"`
	Objects   []*string                                        `json:"objects,omitempty" xml:"objects,omitempty" type:"Repeated"`
}

func (s CreateOssExternalStoreRequestParameter) String() string {
	return tea.Prettify(s)
}

func (s CreateOssExternalStoreRequestParameter) GoString() string {
	return s.String()
}

func (s *CreateOssExternalStoreRequestParameter) SetAccessid(v string) *CreateOssExternalStoreRequestParameter {
	s.Accessid = &v
	return s
}

func (s *CreateOssExternalStoreRequestParameter) SetAccesskey(v string) *CreateOssExternalStoreRequestParameter {
	s.Accesskey = &v
	return s
}

func (s *CreateOssExternalStoreRequestParameter) SetBucket(v string) *CreateOssExternalStoreRequestParameter {
	s.Bucket = &v
	return s
}

func (s *CreateOssExternalStoreRequestParameter) SetColumns(v []*CreateOssExternalStoreRequestParameterColumns) *CreateOssExternalStoreRequestParameter {
	s.Columns = v
	return s
}

func (s *CreateOssExternalStoreRequestParameter) SetEndpoint(v string) *CreateOssExternalStoreRequestParameter {
	s.Endpoint = &v
	return s
}

func (s *CreateOssExternalStoreRequestParameter) SetObjects(v []*string) *CreateOssExternalStoreRequestParameter {
	s.Objects = v
	return s
}

type CreateOssExternalStoreRequestParameterColumns struct {
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s CreateOssExternalStoreRequestParameterColumns) String() string {
	return tea.Prettify(s)
}

func (s CreateOssExternalStoreRequestParameterColumns) GoString() string {
	return s.String()
}

func (s *CreateOssExternalStoreRequestParameterColumns) SetName(v string) *CreateOssExternalStoreRequestParameterColumns {
	s.Name = &v
	return s
}

func (s *CreateOssExternalStoreRequestParameterColumns) SetType(v string) *CreateOssExternalStoreRequestParameterColumns {
	s.Type = &v
	return s
}

type CreateOssExternalStoreResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
}

func (s CreateOssExternalStoreResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateOssExternalStoreResponse) GoString() string {
	return s.String()
}

func (s *CreateOssExternalStoreResponse) SetHeaders(v map[string]*string) *CreateOssExternalStoreResponse {
	s.Headers = v
	return s
}

func (s *CreateOssExternalStoreResponse) SetStatusCode(v int32) *CreateOssExternalStoreResponse {
	s.StatusCode = &v
	return s
}

type CreateOssShipperRequest struct {
	ShipperName         *string                                     `json:"shipperName,omitempty" xml:"shipperName,omitempty"`
	TargetConfiguration *CreateOssShipperRequestTargetConfiguration `json:"targetConfiguration,omitempty" xml:"targetConfiguration,omitempty" type:"Struct"`
	TargetType          *string                                     `json:"targetType,omitempty" xml:"targetType,omitempty"`
}

func (s CreateOssShipperRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateOssShipperRequest) GoString() string {
	return s.String()
}

func (s *CreateOssShipperRequest) SetShipperName(v string) *CreateOssShipperRequest {
	s.ShipperName = &v
	return s
}

func (s *CreateOssShipperRequest) SetTargetConfiguration(v *CreateOssShipperRequestTargetConfiguration) *CreateOssShipperRequest {
	s.TargetConfiguration = v
	return s
}

func (s *CreateOssShipperRequest) SetTargetType(v string) *CreateOssShipperRequest {
	s.TargetType = &v
	return s
}

type CreateOssShipperRequestTargetConfiguration struct {
	BufferInterval *int32                                             `json:"bufferInterval,omitempty" xml:"bufferInterval,omitempty"`
	BufferSize     *int32                                             `json:"bufferSize,omitempty" xml:"bufferSize,omitempty"`
	CompressType   *string                                            `json:"compressType,omitempty" xml:"compressType,omitempty"`
	Enable         *bool                                              `json:"enable,omitempty" xml:"enable,omitempty"`
	OssBucket      *string                                            `json:"ossBucket,omitempty" xml:"ossBucket,omitempty"`
	OssPrefix      *string                                            `json:"ossPrefix,omitempty" xml:"ossPrefix,omitempty"`
	PathFormat     *string                                            `json:"pathFormat,omitempty" xml:"pathFormat,omitempty"`
	RoleArn        *string                                            `json:"roleArn,omitempty" xml:"roleArn,omitempty"`
	Storage        *CreateOssShipperRequestTargetConfigurationStorage `json:"storage,omitempty" xml:"storage,omitempty" type:"Struct"`
	TimeZone       *string                                            `json:"timeZone,omitempty" xml:"timeZone,omitempty"`
}

func (s CreateOssShipperRequestTargetConfiguration) String() string {
	return tea.Prettify(s)
}

func (s CreateOssShipperRequestTargetConfiguration) GoString() string {
	return s.String()
}

func (s *CreateOssShipperRequestTargetConfiguration) SetBufferInterval(v int32) *CreateOssShipperRequestTargetConfiguration {
	s.BufferInterval = &v
	return s
}

func (s *CreateOssShipperRequestTargetConfiguration) SetBufferSize(v int32) *CreateOssShipperRequestTargetConfiguration {
	s.BufferSize = &v
	return s
}

func (s *CreateOssShipperRequestTargetConfiguration) SetCompressType(v string) *CreateOssShipperRequestTargetConfiguration {
	s.CompressType = &v
	return s
}

func (s *CreateOssShipperRequestTargetConfiguration) SetEnable(v bool) *CreateOssShipperRequestTargetConfiguration {
	s.Enable = &v
	return s
}

func (s *CreateOssShipperRequestTargetConfiguration) SetOssBucket(v string) *CreateOssShipperRequestTargetConfiguration {
	s.OssBucket = &v
	return s
}

func (s *CreateOssShipperRequestTargetConfiguration) SetOssPrefix(v string) *CreateOssShipperRequestTargetConfiguration {
	s.OssPrefix = &v
	return s
}

func (s *CreateOssShipperRequestTargetConfiguration) SetPathFormat(v string) *CreateOssShipperRequestTargetConfiguration {
	s.PathFormat = &v
	return s
}

func (s *CreateOssShipperRequestTargetConfiguration) SetRoleArn(v string) *CreateOssShipperRequestTargetConfiguration {
	s.RoleArn = &v
	return s
}

func (s *CreateOssShipperRequestTargetConfiguration) SetStorage(v *CreateOssShipperRequestTargetConfigurationStorage) *CreateOssShipperRequestTargetConfiguration {
	s.Storage = v
	return s
}

func (s *CreateOssShipperRequestTargetConfiguration) SetTimeZone(v string) *CreateOssShipperRequestTargetConfiguration {
	s.TimeZone = &v
	return s
}

type CreateOssShipperRequestTargetConfigurationStorage struct {
	Detail map[string]interface{} `json:"detail,omitempty" xml:"detail,omitempty"`
	Format *string                `json:"format,omitempty" xml:"format,omitempty"`
}

func (s CreateOssShipperRequestTargetConfigurationStorage) String() string {
	return tea.Prettify(s)
}

func (s CreateOssShipperRequestTargetConfigurationStorage) GoString() string {
	return s.String()
}

func (s *CreateOssShipperRequestTargetConfigurationStorage) SetDetail(v map[string]interface{}) *CreateOssShipperRequestTargetConfigurationStorage {
	s.Detail = v
	return s
}

func (s *CreateOssShipperRequestTargetConfigurationStorage) SetFormat(v string) *CreateOssShipperRequestTargetConfigurationStorage {
	s.Format = &v
	return s
}

type CreateOssShipperResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
}

func (s CreateOssShipperResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateOssShipperResponse) GoString() string {
	return s.String()
}

func (s *CreateOssShipperResponse) SetHeaders(v map[string]*string) *CreateOssShipperResponse {
	s.Headers = v
	return s
}

func (s *CreateOssShipperResponse) SetStatusCode(v int32) *CreateOssShipperResponse {
	s.StatusCode = &v
	return s
}

type CreateProjectRequest struct {
	Description     *string `json:"description,omitempty" xml:"description,omitempty"`
	ProjectName     *string `json:"projectName,omitempty" xml:"projectName,omitempty"`
	ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
}

func (s CreateProjectRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateProjectRequest) GoString() string {
	return s.String()
}

func (s *CreateProjectRequest) SetDescription(v string) *CreateProjectRequest {
	s.Description = &v
	return s
}

func (s *CreateProjectRequest) SetProjectName(v string) *CreateProjectRequest {
	s.ProjectName = &v
	return s
}

func (s *CreateProjectRequest) SetResourceGroupId(v string) *CreateProjectRequest {
	s.ResourceGroupId = &v
	return s
}

type CreateProjectResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
}

func (s CreateProjectResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateProjectResponse) GoString() string {
	return s.String()
}

func (s *CreateProjectResponse) SetHeaders(v map[string]*string) *CreateProjectResponse {
	s.Headers = v
	return s
}

func (s *CreateProjectResponse) SetStatusCode(v int32) *CreateProjectResponse {
	s.StatusCode = &v
	return s
}

type CreateRdsExternalStoreRequest struct {
	ExternalStoreName *string                                 `json:"externalStoreName,omitempty" xml:"externalStoreName,omitempty"`
	Parameter         *CreateRdsExternalStoreRequestParameter `json:"parameter,omitempty" xml:"parameter,omitempty" type:"Struct"`
	StoreType         *string                                 `json:"storeType,omitempty" xml:"storeType,omitempty"`
}

func (s CreateRdsExternalStoreRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateRdsExternalStoreRequest) GoString() string {
	return s.String()
}

func (s *CreateRdsExternalStoreRequest) SetExternalStoreName(v string) *CreateRdsExternalStoreRequest {
	s.ExternalStoreName = &v
	return s
}

func (s *CreateRdsExternalStoreRequest) SetParameter(v *CreateRdsExternalStoreRequestParameter) *CreateRdsExternalStoreRequest {
	s.Parameter = v
	return s
}

func (s *CreateRdsExternalStoreRequest) SetStoreType(v string) *CreateRdsExternalStoreRequest {
	s.StoreType = &v
	return s
}

type CreateRdsExternalStoreRequestParameter struct {
	Db         *string `json:"db,omitempty" xml:"db,omitempty"`
	Host       *string `json:"host,omitempty" xml:"host,omitempty"`
	InstanceId *string `json:"instance-id,omitempty" xml:"instance-id,omitempty"`
	Password   *string `json:"password,omitempty" xml:"password,omitempty"`
	Port       *string `json:"port,omitempty" xml:"port,omitempty"`
	Region     *string `json:"region,omitempty" xml:"region,omitempty"`
	Table      *string `json:"table,omitempty" xml:"table,omitempty"`
	Username   *string `json:"username,omitempty" xml:"username,omitempty"`
	VpcId      *string `json:"vpc-id,omitempty" xml:"vpc-id,omitempty"`
}

func (s CreateRdsExternalStoreRequestParameter) String() string {
	return tea.Prettify(s)
}

func (s CreateRdsExternalStoreRequestParameter) GoString() string {
	return s.String()
}

func (s *CreateRdsExternalStoreRequestParameter) SetDb(v string) *CreateRdsExternalStoreRequestParameter {
	s.Db = &v
	return s
}

func (s *CreateRdsExternalStoreRequestParameter) SetHost(v string) *CreateRdsExternalStoreRequestParameter {
	s.Host = &v
	return s
}

func (s *CreateRdsExternalStoreRequestParameter) SetInstanceId(v string) *CreateRdsExternalStoreRequestParameter {
	s.InstanceId = &v
	return s
}

func (s *CreateRdsExternalStoreRequestParameter) SetPassword(v string) *CreateRdsExternalStoreRequestParameter {
	s.Password = &v
	return s
}

func (s *CreateRdsExternalStoreRequestParameter) SetPort(v string) *CreateRdsExternalStoreRequestParameter {
	s.Port = &v
	return s
}

func (s *CreateRdsExternalStoreRequestParameter) SetRegion(v string) *CreateRdsExternalStoreRequestParameter {
	s.Region = &v
	return s
}

func (s *CreateRdsExternalStoreRequestParameter) SetTable(v string) *CreateRdsExternalStoreRequestParameter {
	s.Table = &v
	return s
}

func (s *CreateRdsExternalStoreRequestParameter) SetUsername(v string) *CreateRdsExternalStoreRequestParameter {
	s.Username = &v
	return s
}

func (s *CreateRdsExternalStoreRequestParameter) SetVpcId(v string) *CreateRdsExternalStoreRequestParameter {
	s.VpcId = &v
	return s
}

type CreateRdsExternalStoreResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
}

func (s CreateRdsExternalStoreResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateRdsExternalStoreResponse) GoString() string {
	return s.String()
}

func (s *CreateRdsExternalStoreResponse) SetHeaders(v map[string]*string) *CreateRdsExternalStoreResponse {
	s.Headers = v
	return s
}

func (s *CreateRdsExternalStoreResponse) SetStatusCode(v int32) *CreateRdsExternalStoreResponse {
	s.StatusCode = &v
	return s
}

type CreateSavedSearchRequest struct {
	DisplayName     *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	Logstore        *string `json:"logstore,omitempty" xml:"logstore,omitempty"`
	SavedsearchName *string `json:"savedsearchName,omitempty" xml:"savedsearchName,omitempty"`
	SearchQuery     *string `json:"searchQuery,omitempty" xml:"searchQuery,omitempty"`
	Topic           *string `json:"topic,omitempty" xml:"topic,omitempty"`
}

func (s CreateSavedSearchRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateSavedSearchRequest) GoString() string {
	return s.String()
}

func (s *CreateSavedSearchRequest) SetDisplayName(v string) *CreateSavedSearchRequest {
	s.DisplayName = &v
	return s
}

func (s *CreateSavedSearchRequest) SetLogstore(v string) *CreateSavedSearchRequest {
	s.Logstore = &v
	return s
}

func (s *CreateSavedSearchRequest) SetSavedsearchName(v string) *CreateSavedSearchRequest {
	s.SavedsearchName = &v
	return s
}

func (s *CreateSavedSearchRequest) SetSearchQuery(v string) *CreateSavedSearchRequest {
	s.SearchQuery = &v
	return s
}

func (s *CreateSavedSearchRequest) SetTopic(v string) *CreateSavedSearchRequest {
	s.Topic = &v
	return s
}

type CreateSavedSearchResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
}

func (s CreateSavedSearchResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateSavedSearchResponse) GoString() string {
	return s.String()
}

func (s *CreateSavedSearchResponse) SetHeaders(v map[string]*string) *CreateSavedSearchResponse {
	s.Headers = v
	return s
}

func (s *CreateSavedSearchResponse) SetStatusCode(v int32) *CreateSavedSearchResponse {
	s.StatusCode = &v
	return s
}

type DeleteAlertResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
}

func (s DeleteAlertResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteAlertResponse) GoString() string {
	return s.String()
}

func (s *DeleteAlertResponse) SetHeaders(v map[string]*string) *DeleteAlertResponse {
	s.Headers = v
	return s
}

func (s *DeleteAlertResponse) SetStatusCode(v int32) *DeleteAlertResponse {
	s.StatusCode = &v
	return s
}

type DeleteConsumerGroupResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
}

func (s DeleteConsumerGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteConsumerGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteConsumerGroupResponse) SetHeaders(v map[string]*string) *DeleteConsumerGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteConsumerGroupResponse) SetStatusCode(v int32) *DeleteConsumerGroupResponse {
	s.StatusCode = &v
	return s
}

type DeleteDomainResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
}

func (s DeleteDomainResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDomainResponse) GoString() string {
	return s.String()
}

func (s *DeleteDomainResponse) SetHeaders(v map[string]*string) *DeleteDomainResponse {
	s.Headers = v
	return s
}

func (s *DeleteDomainResponse) SetStatusCode(v int32) *DeleteDomainResponse {
	s.StatusCode = &v
	return s
}

type DeleteETLJobResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *string            `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteETLJobResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteETLJobResponse) GoString() string {
	return s.String()
}

func (s *DeleteETLJobResponse) SetHeaders(v map[string]*string) *DeleteETLJobResponse {
	s.Headers = v
	return s
}

func (s *DeleteETLJobResponse) SetStatusCode(v int32) *DeleteETLJobResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteETLJobResponse) SetBody(v string) *DeleteETLJobResponse {
	s.Body = &v
	return s
}

type DeleteExternalStoreResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
}

func (s DeleteExternalStoreResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteExternalStoreResponse) GoString() string {
	return s.String()
}

func (s *DeleteExternalStoreResponse) SetHeaders(v map[string]*string) *DeleteExternalStoreResponse {
	s.Headers = v
	return s
}

func (s *DeleteExternalStoreResponse) SetStatusCode(v int32) *DeleteExternalStoreResponse {
	s.StatusCode = &v
	return s
}

type DeleteIndexResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
}

func (s DeleteIndexResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteIndexResponse) GoString() string {
	return s.String()
}

func (s *DeleteIndexResponse) SetHeaders(v map[string]*string) *DeleteIndexResponse {
	s.Headers = v
	return s
}

func (s *DeleteIndexResponse) SetStatusCode(v int32) *DeleteIndexResponse {
	s.StatusCode = &v
	return s
}

type DeleteJobResponseBody struct {
	Data interface{} `json:"data,omitempty" xml:"data,omitempty"`
}

func (s DeleteJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteJobResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteJobResponseBody) SetData(v interface{}) *DeleteJobResponseBody {
	s.Data = v
	return s
}

type DeleteJobResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteJobResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteJobResponse) GoString() string {
	return s.String()
}

func (s *DeleteJobResponse) SetHeaders(v map[string]*string) *DeleteJobResponse {
	s.Headers = v
	return s
}

func (s *DeleteJobResponse) SetStatusCode(v int32) *DeleteJobResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteJobResponse) SetBody(v *DeleteJobResponseBody) *DeleteJobResponse {
	s.Body = v
	return s
}

type DeleteLogStoreResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
}

func (s DeleteLogStoreResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteLogStoreResponse) GoString() string {
	return s.String()
}

func (s *DeleteLogStoreResponse) SetHeaders(v map[string]*string) *DeleteLogStoreResponse {
	s.Headers = v
	return s
}

func (s *DeleteLogStoreResponse) SetStatusCode(v int32) *DeleteLogStoreResponse {
	s.StatusCode = &v
	return s
}

type DeleteLoggingResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
}

func (s DeleteLoggingResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteLoggingResponse) GoString() string {
	return s.String()
}

func (s *DeleteLoggingResponse) SetHeaders(v map[string]*string) *DeleteLoggingResponse {
	s.Headers = v
	return s
}

func (s *DeleteLoggingResponse) SetStatusCode(v int32) *DeleteLoggingResponse {
	s.StatusCode = &v
	return s
}

type DeleteMachineGroupResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
}

func (s DeleteMachineGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteMachineGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteMachineGroupResponse) SetHeaders(v map[string]*string) *DeleteMachineGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteMachineGroupResponse) SetStatusCode(v int32) *DeleteMachineGroupResponse {
	s.StatusCode = &v
	return s
}

type DeleteProjectResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
}

func (s DeleteProjectResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteProjectResponse) GoString() string {
	return s.String()
}

func (s *DeleteProjectResponse) SetHeaders(v map[string]*string) *DeleteProjectResponse {
	s.Headers = v
	return s
}

func (s *DeleteProjectResponse) SetStatusCode(v int32) *DeleteProjectResponse {
	s.StatusCode = &v
	return s
}

type DeleteProjectPolicyResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
}

func (s DeleteProjectPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteProjectPolicyResponse) GoString() string {
	return s.String()
}

func (s *DeleteProjectPolicyResponse) SetHeaders(v map[string]*string) *DeleteProjectPolicyResponse {
	s.Headers = v
	return s
}

func (s *DeleteProjectPolicyResponse) SetStatusCode(v int32) *DeleteProjectPolicyResponse {
	s.StatusCode = &v
	return s
}

type DeleteSavedSearchResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
}

func (s DeleteSavedSearchResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteSavedSearchResponse) GoString() string {
	return s.String()
}

func (s *DeleteSavedSearchResponse) SetHeaders(v map[string]*string) *DeleteSavedSearchResponse {
	s.Headers = v
	return s
}

func (s *DeleteSavedSearchResponse) SetStatusCode(v int32) *DeleteSavedSearchResponse {
	s.StatusCode = &v
	return s
}

type DeleteShipperResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
}

func (s DeleteShipperResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteShipperResponse) GoString() string {
	return s.String()
}

func (s *DeleteShipperResponse) SetHeaders(v map[string]*string) *DeleteShipperResponse {
	s.Headers = v
	return s
}

func (s *DeleteShipperResponse) SetStatusCode(v int32) *DeleteShipperResponse {
	s.StatusCode = &v
	return s
}

type GetAlertResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *Alert             `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetAlertResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAlertResponse) GoString() string {
	return s.String()
}

func (s *GetAlertResponse) SetHeaders(v map[string]*string) *GetAlertResponse {
	s.Headers = v
	return s
}

func (s *GetAlertResponse) SetStatusCode(v int32) *GetAlertResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAlertResponse) SetBody(v *Alert) *GetAlertResponse {
	s.Body = v
	return s
}

type GetAppliedConfigsResponseBody struct {
	Configs []*string `json:"configs,omitempty" xml:"configs,omitempty" type:"Repeated"`
	Count   *int32    `json:"count,omitempty" xml:"count,omitempty"`
}

func (s GetAppliedConfigsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAppliedConfigsResponseBody) GoString() string {
	return s.String()
}

func (s *GetAppliedConfigsResponseBody) SetConfigs(v []*string) *GetAppliedConfigsResponseBody {
	s.Configs = v
	return s
}

func (s *GetAppliedConfigsResponseBody) SetCount(v int32) *GetAppliedConfigsResponseBody {
	s.Count = &v
	return s
}

type GetAppliedConfigsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetAppliedConfigsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetAppliedConfigsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAppliedConfigsResponse) GoString() string {
	return s.String()
}

func (s *GetAppliedConfigsResponse) SetHeaders(v map[string]*string) *GetAppliedConfigsResponse {
	s.Headers = v
	return s
}

func (s *GetAppliedConfigsResponse) SetStatusCode(v int32) *GetAppliedConfigsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAppliedConfigsResponse) SetBody(v *GetAppliedConfigsResponseBody) *GetAppliedConfigsResponse {
	s.Body = v
	return s
}

type GetAppliedMachineGroupsResponseBody struct {
	Count         *int32    `json:"count,omitempty" xml:"count,omitempty"`
	Machinegroups []*string `json:"machinegroups,omitempty" xml:"machinegroups,omitempty" type:"Repeated"`
}

func (s GetAppliedMachineGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAppliedMachineGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *GetAppliedMachineGroupsResponseBody) SetCount(v int32) *GetAppliedMachineGroupsResponseBody {
	s.Count = &v
	return s
}

func (s *GetAppliedMachineGroupsResponseBody) SetMachinegroups(v []*string) *GetAppliedMachineGroupsResponseBody {
	s.Machinegroups = v
	return s
}

type GetAppliedMachineGroupsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetAppliedMachineGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetAppliedMachineGroupsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAppliedMachineGroupsResponse) GoString() string {
	return s.String()
}

func (s *GetAppliedMachineGroupsResponse) SetHeaders(v map[string]*string) *GetAppliedMachineGroupsResponse {
	s.Headers = v
	return s
}

func (s *GetAppliedMachineGroupsResponse) SetStatusCode(v int32) *GetAppliedMachineGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAppliedMachineGroupsResponse) SetBody(v *GetAppliedMachineGroupsResponseBody) *GetAppliedMachineGroupsResponse {
	s.Body = v
	return s
}

type GetCheckPointRequest struct {
	Shard *int32 `json:"shard,omitempty" xml:"shard,omitempty"`
}

func (s GetCheckPointRequest) String() string {
	return tea.Prettify(s)
}

func (s GetCheckPointRequest) GoString() string {
	return s.String()
}

func (s *GetCheckPointRequest) SetShard(v int32) *GetCheckPointRequest {
	s.Shard = &v
	return s
}

type GetCheckPointResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       []*GetCheckPointResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true" type:"Repeated"`
}

func (s GetCheckPointResponse) String() string {
	return tea.Prettify(s)
}

func (s GetCheckPointResponse) GoString() string {
	return s.String()
}

func (s *GetCheckPointResponse) SetHeaders(v map[string]*string) *GetCheckPointResponse {
	s.Headers = v
	return s
}

func (s *GetCheckPointResponse) SetStatusCode(v int32) *GetCheckPointResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCheckPointResponse) SetBody(v []*GetCheckPointResponseBody) *GetCheckPointResponse {
	s.Body = v
	return s
}

type GetCheckPointResponseBody struct {
	// Shard ID。
	Shard      *int32  `json:"shard,omitempty" xml:"shard,omitempty"`
	Checkpoint *string `json:"checkpoint,omitempty" xml:"checkpoint,omitempty"`
	UpdateTime *int64  `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
	Consumer   *string `json:"consumer,omitempty" xml:"consumer,omitempty"`
}

func (s GetCheckPointResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetCheckPointResponseBody) GoString() string {
	return s.String()
}

func (s *GetCheckPointResponseBody) SetShard(v int32) *GetCheckPointResponseBody {
	s.Shard = &v
	return s
}

func (s *GetCheckPointResponseBody) SetCheckpoint(v string) *GetCheckPointResponseBody {
	s.Checkpoint = &v
	return s
}

func (s *GetCheckPointResponseBody) SetUpdateTime(v int64) *GetCheckPointResponseBody {
	s.UpdateTime = &v
	return s
}

func (s *GetCheckPointResponseBody) SetConsumer(v string) *GetCheckPointResponseBody {
	s.Consumer = &v
	return s
}

type GetContextLogsRequest struct {
	BackLines    *int64  `json:"back_lines,omitempty" xml:"back_lines,omitempty"`
	ForwardLines *int64  `json:"forward_lines,omitempty" xml:"forward_lines,omitempty"`
	PackId       *string `json:"pack_id,omitempty" xml:"pack_id,omitempty"`
	PackMeta     *string `json:"pack_meta,omitempty" xml:"pack_meta,omitempty"`
	Type         *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetContextLogsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetContextLogsRequest) GoString() string {
	return s.String()
}

func (s *GetContextLogsRequest) SetBackLines(v int64) *GetContextLogsRequest {
	s.BackLines = &v
	return s
}

func (s *GetContextLogsRequest) SetForwardLines(v int64) *GetContextLogsRequest {
	s.ForwardLines = &v
	return s
}

func (s *GetContextLogsRequest) SetPackId(v string) *GetContextLogsRequest {
	s.PackId = &v
	return s
}

func (s *GetContextLogsRequest) SetPackMeta(v string) *GetContextLogsRequest {
	s.PackMeta = &v
	return s
}

func (s *GetContextLogsRequest) SetType(v string) *GetContextLogsRequest {
	s.Type = &v
	return s
}

type GetContextLogsResponseBody struct {
	BackLines    *int64                   `json:"back_lines,omitempty" xml:"back_lines,omitempty"`
	ForwardLines *int64                   `json:"forward_lines,omitempty" xml:"forward_lines,omitempty"`
	Logs         []map[string]interface{} `json:"logs,omitempty" xml:"logs,omitempty" type:"Repeated"`
	Progress     *string                  `json:"progress,omitempty" xml:"progress,omitempty"`
	TotalLines   *int64                   `json:"total_lines,omitempty" xml:"total_lines,omitempty"`
}

func (s GetContextLogsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetContextLogsResponseBody) GoString() string {
	return s.String()
}

func (s *GetContextLogsResponseBody) SetBackLines(v int64) *GetContextLogsResponseBody {
	s.BackLines = &v
	return s
}

func (s *GetContextLogsResponseBody) SetForwardLines(v int64) *GetContextLogsResponseBody {
	s.ForwardLines = &v
	return s
}

func (s *GetContextLogsResponseBody) SetLogs(v []map[string]interface{}) *GetContextLogsResponseBody {
	s.Logs = v
	return s
}

func (s *GetContextLogsResponseBody) SetProgress(v string) *GetContextLogsResponseBody {
	s.Progress = &v
	return s
}

func (s *GetContextLogsResponseBody) SetTotalLines(v int64) *GetContextLogsResponseBody {
	s.TotalLines = &v
	return s
}

type GetContextLogsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetContextLogsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetContextLogsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetContextLogsResponse) GoString() string {
	return s.String()
}

func (s *GetContextLogsResponse) SetHeaders(v map[string]*string) *GetContextLogsResponse {
	s.Headers = v
	return s
}

func (s *GetContextLogsResponse) SetStatusCode(v int32) *GetContextLogsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetContextLogsResponse) SetBody(v *GetContextLogsResponseBody) *GetContextLogsResponse {
	s.Body = v
	return s
}

type GetCursorRequest struct {
	From *string `json:"from,omitempty" xml:"from,omitempty"`
}

func (s GetCursorRequest) String() string {
	return tea.Prettify(s)
}

func (s GetCursorRequest) GoString() string {
	return s.String()
}

func (s *GetCursorRequest) SetFrom(v string) *GetCursorRequest {
	s.From = &v
	return s
}

type GetCursorResponseBody struct {
	Cursor *string `json:"cursor,omitempty" xml:"cursor,omitempty"`
}

func (s GetCursorResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetCursorResponseBody) GoString() string {
	return s.String()
}

func (s *GetCursorResponseBody) SetCursor(v string) *GetCursorResponseBody {
	s.Cursor = &v
	return s
}

type GetCursorResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetCursorResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetCursorResponse) String() string {
	return tea.Prettify(s)
}

func (s GetCursorResponse) GoString() string {
	return s.String()
}

func (s *GetCursorResponse) SetHeaders(v map[string]*string) *GetCursorResponse {
	s.Headers = v
	return s
}

func (s *GetCursorResponse) SetStatusCode(v int32) *GetCursorResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCursorResponse) SetBody(v *GetCursorResponseBody) *GetCursorResponse {
	s.Body = v
	return s
}

type GetCursorTimeRequest struct {
	Cursor *string `json:"cursor,omitempty" xml:"cursor,omitempty"`
}

func (s GetCursorTimeRequest) String() string {
	return tea.Prettify(s)
}

func (s GetCursorTimeRequest) GoString() string {
	return s.String()
}

func (s *GetCursorTimeRequest) SetCursor(v string) *GetCursorTimeRequest {
	s.Cursor = &v
	return s
}

type GetCursorTimeResponseBody struct {
	CursorTime *string `json:"cursor_time,omitempty" xml:"cursor_time,omitempty"`
}

func (s GetCursorTimeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetCursorTimeResponseBody) GoString() string {
	return s.String()
}

func (s *GetCursorTimeResponseBody) SetCursorTime(v string) *GetCursorTimeResponseBody {
	s.CursorTime = &v
	return s
}

type GetCursorTimeResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetCursorTimeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetCursorTimeResponse) String() string {
	return tea.Prettify(s)
}

func (s GetCursorTimeResponse) GoString() string {
	return s.String()
}

func (s *GetCursorTimeResponse) SetHeaders(v map[string]*string) *GetCursorTimeResponse {
	s.Headers = v
	return s
}

func (s *GetCursorTimeResponse) SetStatusCode(v int32) *GetCursorTimeResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCursorTimeResponse) SetBody(v *GetCursorTimeResponseBody) *GetCursorTimeResponse {
	s.Body = v
	return s
}

type GetEtlJobResponseBody struct {
	Code      *int32                   `json:"code,omitempty" xml:"code,omitempty"`
	Data      []map[string]interface{} `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	Message   *string                  `json:"message,omitempty" xml:"message,omitempty"`
	RequestId *string                  `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *string                  `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetEtlJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetEtlJobResponseBody) GoString() string {
	return s.String()
}

func (s *GetEtlJobResponseBody) SetCode(v int32) *GetEtlJobResponseBody {
	s.Code = &v
	return s
}

func (s *GetEtlJobResponseBody) SetData(v []map[string]interface{}) *GetEtlJobResponseBody {
	s.Data = v
	return s
}

func (s *GetEtlJobResponseBody) SetMessage(v string) *GetEtlJobResponseBody {
	s.Message = &v
	return s
}

func (s *GetEtlJobResponseBody) SetRequestId(v string) *GetEtlJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetEtlJobResponseBody) SetSuccess(v string) *GetEtlJobResponseBody {
	s.Success = &v
	return s
}

type GetEtlJobResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetEtlJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetEtlJobResponse) String() string {
	return tea.Prettify(s)
}

func (s GetEtlJobResponse) GoString() string {
	return s.String()
}

func (s *GetEtlJobResponse) SetHeaders(v map[string]*string) *GetEtlJobResponse {
	s.Headers = v
	return s
}

func (s *GetEtlJobResponse) SetStatusCode(v int32) *GetEtlJobResponse {
	s.StatusCode = &v
	return s
}

func (s *GetEtlJobResponse) SetBody(v *GetEtlJobResponseBody) *GetEtlJobResponse {
	s.Body = v
	return s
}

type GetExternalStoreResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ExternalStore     `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetExternalStoreResponse) String() string {
	return tea.Prettify(s)
}

func (s GetExternalStoreResponse) GoString() string {
	return s.String()
}

func (s *GetExternalStoreResponse) SetHeaders(v map[string]*string) *GetExternalStoreResponse {
	s.Headers = v
	return s
}

func (s *GetExternalStoreResponse) SetStatusCode(v int32) *GetExternalStoreResponse {
	s.StatusCode = &v
	return s
}

func (s *GetExternalStoreResponse) SetBody(v *ExternalStore) *GetExternalStoreResponse {
	s.Body = v
	return s
}

type GetHistogramsRequest struct {
	From  *int64  `json:"from,omitempty" xml:"from,omitempty"`
	Query *string `json:"query,omitempty" xml:"query,omitempty"`
	To    *int64  `json:"to,omitempty" xml:"to,omitempty"`
	Topic *string `json:"topic,omitempty" xml:"topic,omitempty"`
}

func (s GetHistogramsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetHistogramsRequest) GoString() string {
	return s.String()
}

func (s *GetHistogramsRequest) SetFrom(v int64) *GetHistogramsRequest {
	s.From = &v
	return s
}

func (s *GetHistogramsRequest) SetQuery(v string) *GetHistogramsRequest {
	s.Query = &v
	return s
}

func (s *GetHistogramsRequest) SetTo(v int64) *GetHistogramsRequest {
	s.To = &v
	return s
}

func (s *GetHistogramsRequest) SetTopic(v string) *GetHistogramsRequest {
	s.Topic = &v
	return s
}

type GetHistogramsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       []*GetHistogramsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true" type:"Repeated"`
}

func (s GetHistogramsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetHistogramsResponse) GoString() string {
	return s.String()
}

func (s *GetHistogramsResponse) SetHeaders(v map[string]*string) *GetHistogramsResponse {
	s.Headers = v
	return s
}

func (s *GetHistogramsResponse) SetStatusCode(v int32) *GetHistogramsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetHistogramsResponse) SetBody(v []*GetHistogramsResponseBody) *GetHistogramsResponse {
	s.Body = v
	return s
}

type GetHistogramsResponseBody struct {
	From     *int64  `json:"from,omitempty" xml:"from,omitempty"`
	To       *int64  `json:"to,omitempty" xml:"to,omitempty"`
	Count    *int64  `json:"count,omitempty" xml:"count,omitempty"`
	Progress *string `json:"progress,omitempty" xml:"progress,omitempty"`
}

func (s GetHistogramsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetHistogramsResponseBody) GoString() string {
	return s.String()
}

func (s *GetHistogramsResponseBody) SetFrom(v int64) *GetHistogramsResponseBody {
	s.From = &v
	return s
}

func (s *GetHistogramsResponseBody) SetTo(v int64) *GetHistogramsResponseBody {
	s.To = &v
	return s
}

func (s *GetHistogramsResponseBody) SetCount(v int64) *GetHistogramsResponseBody {
	s.Count = &v
	return s
}

func (s *GetHistogramsResponseBody) SetProgress(v string) *GetHistogramsResponseBody {
	s.Progress = &v
	return s
}

type GetIndexResponseBody struct {
	IndexMode          *string                   `json:"index_mode,omitempty" xml:"index_mode,omitempty"`
	Keys               map[string]*KeysValue     `json:"keys,omitempty" xml:"keys,omitempty"`
	LastModifyTime     *int64                    `json:"lastModifyTime,omitempty" xml:"lastModifyTime,omitempty"`
	Line               *GetIndexResponseBodyLine `json:"line,omitempty" xml:"line,omitempty" type:"Struct"`
	LogReduce          *bool                     `json:"log_reduce,omitempty" xml:"log_reduce,omitempty"`
	LogReduceBlackList []*string                 `json:"log_reduce_black_list,omitempty" xml:"log_reduce_black_list,omitempty" type:"Repeated"`
	LogReduceWhiteList []*string                 `json:"log_reduce_white_list,omitempty" xml:"log_reduce_white_list,omitempty" type:"Repeated"`
	MaxTextLen         *int32                    `json:"max_text_len,omitempty" xml:"max_text_len,omitempty"`
	Storage            *string                   `json:"storage,omitempty" xml:"storage,omitempty"`
	Ttl                *int32                    `json:"ttl,omitempty" xml:"ttl,omitempty"`
}

func (s GetIndexResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetIndexResponseBody) GoString() string {
	return s.String()
}

func (s *GetIndexResponseBody) SetIndexMode(v string) *GetIndexResponseBody {
	s.IndexMode = &v
	return s
}

func (s *GetIndexResponseBody) SetKeys(v map[string]*KeysValue) *GetIndexResponseBody {
	s.Keys = v
	return s
}

func (s *GetIndexResponseBody) SetLastModifyTime(v int64) *GetIndexResponseBody {
	s.LastModifyTime = &v
	return s
}

func (s *GetIndexResponseBody) SetLine(v *GetIndexResponseBodyLine) *GetIndexResponseBody {
	s.Line = v
	return s
}

func (s *GetIndexResponseBody) SetLogReduce(v bool) *GetIndexResponseBody {
	s.LogReduce = &v
	return s
}

func (s *GetIndexResponseBody) SetLogReduceBlackList(v []*string) *GetIndexResponseBody {
	s.LogReduceBlackList = v
	return s
}

func (s *GetIndexResponseBody) SetLogReduceWhiteList(v []*string) *GetIndexResponseBody {
	s.LogReduceWhiteList = v
	return s
}

func (s *GetIndexResponseBody) SetMaxTextLen(v int32) *GetIndexResponseBody {
	s.MaxTextLen = &v
	return s
}

func (s *GetIndexResponseBody) SetStorage(v string) *GetIndexResponseBody {
	s.Storage = &v
	return s
}

func (s *GetIndexResponseBody) SetTtl(v int32) *GetIndexResponseBody {
	s.Ttl = &v
	return s
}

type GetIndexResponseBodyLine struct {
	CaseSensitive *bool     `json:"caseSensitive,omitempty" xml:"caseSensitive,omitempty"`
	Chn           *bool     `json:"chn,omitempty" xml:"chn,omitempty"`
	ExcludeKeys   []*string `json:"exclude_keys,omitempty" xml:"exclude_keys,omitempty" type:"Repeated"`
	IncludeKeys   []*string `json:"include_keys,omitempty" xml:"include_keys,omitempty" type:"Repeated"`
	Token         []*string `json:"token,omitempty" xml:"token,omitempty" type:"Repeated"`
}

func (s GetIndexResponseBodyLine) String() string {
	return tea.Prettify(s)
}

func (s GetIndexResponseBodyLine) GoString() string {
	return s.String()
}

func (s *GetIndexResponseBodyLine) SetCaseSensitive(v bool) *GetIndexResponseBodyLine {
	s.CaseSensitive = &v
	return s
}

func (s *GetIndexResponseBodyLine) SetChn(v bool) *GetIndexResponseBodyLine {
	s.Chn = &v
	return s
}

func (s *GetIndexResponseBodyLine) SetExcludeKeys(v []*string) *GetIndexResponseBodyLine {
	s.ExcludeKeys = v
	return s
}

func (s *GetIndexResponseBodyLine) SetIncludeKeys(v []*string) *GetIndexResponseBodyLine {
	s.IncludeKeys = v
	return s
}

func (s *GetIndexResponseBodyLine) SetToken(v []*string) *GetIndexResponseBodyLine {
	s.Token = v
	return s
}

type GetIndexResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetIndexResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetIndexResponse) String() string {
	return tea.Prettify(s)
}

func (s GetIndexResponse) GoString() string {
	return s.String()
}

func (s *GetIndexResponse) SetHeaders(v map[string]*string) *GetIndexResponse {
	s.Headers = v
	return s
}

func (s *GetIndexResponse) SetStatusCode(v int32) *GetIndexResponse {
	s.StatusCode = &v
	return s
}

func (s *GetIndexResponse) SetBody(v *GetIndexResponseBody) *GetIndexResponse {
	s.Body = v
	return s
}

type GetJobResponseBody struct {
	Code      *int32               `json:"code,omitempty" xml:"code,omitempty"`
	Data      []map[string]*string `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	Message   *string              `json:"message,omitempty" xml:"message,omitempty"`
	RequestId *string              `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *string              `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetJobResponseBody) GoString() string {
	return s.String()
}

func (s *GetJobResponseBody) SetCode(v int32) *GetJobResponseBody {
	s.Code = &v
	return s
}

func (s *GetJobResponseBody) SetData(v []map[string]*string) *GetJobResponseBody {
	s.Data = v
	return s
}

func (s *GetJobResponseBody) SetMessage(v string) *GetJobResponseBody {
	s.Message = &v
	return s
}

func (s *GetJobResponseBody) SetRequestId(v string) *GetJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetJobResponseBody) SetSuccess(v string) *GetJobResponseBody {
	s.Success = &v
	return s
}

type GetJobResponse struct {
	Headers    map[string]*string  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetJobResponse) String() string {
	return tea.Prettify(s)
}

func (s GetJobResponse) GoString() string {
	return s.String()
}

func (s *GetJobResponse) SetHeaders(v map[string]*string) *GetJobResponse {
	s.Headers = v
	return s
}

func (s *GetJobResponse) SetStatusCode(v int32) *GetJobResponse {
	s.StatusCode = &v
	return s
}

func (s *GetJobResponse) SetBody(v *GetJobResponseBody) *GetJobResponse {
	s.Body = v
	return s
}

type GetLogStoreResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *Logstore          `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetLogStoreResponse) String() string {
	return tea.Prettify(s)
}

func (s GetLogStoreResponse) GoString() string {
	return s.String()
}

func (s *GetLogStoreResponse) SetHeaders(v map[string]*string) *GetLogStoreResponse {
	s.Headers = v
	return s
}

func (s *GetLogStoreResponse) SetStatusCode(v int32) *GetLogStoreResponse {
	s.StatusCode = &v
	return s
}

func (s *GetLogStoreResponse) SetBody(v *Logstore) *GetLogStoreResponse {
	s.Body = v
	return s
}

type GetLoggingResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *Logging           `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetLoggingResponse) String() string {
	return tea.Prettify(s)
}

func (s GetLoggingResponse) GoString() string {
	return s.String()
}

func (s *GetLoggingResponse) SetHeaders(v map[string]*string) *GetLoggingResponse {
	s.Headers = v
	return s
}

func (s *GetLoggingResponse) SetStatusCode(v int32) *GetLoggingResponse {
	s.StatusCode = &v
	return s
}

func (s *GetLoggingResponse) SetBody(v *Logging) *GetLoggingResponse {
	s.Body = v
	return s
}

type GetLogsRequest struct {
	From     *int64  `json:"from,omitempty" xml:"from,omitempty"`
	Line     *int64  `json:"line,omitempty" xml:"line,omitempty"`
	Offset   *int64  `json:"offset,omitempty" xml:"offset,omitempty"`
	PowerSql *bool   `json:"powerSql,omitempty" xml:"powerSql,omitempty"`
	Query    *string `json:"query,omitempty" xml:"query,omitempty"`
	Reverse  *bool   `json:"reverse,omitempty" xml:"reverse,omitempty"`
	To       *int64  `json:"to,omitempty" xml:"to,omitempty"`
	Topic    *string `json:"topic,omitempty" xml:"topic,omitempty"`
}

func (s GetLogsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetLogsRequest) GoString() string {
	return s.String()
}

func (s *GetLogsRequest) SetFrom(v int64) *GetLogsRequest {
	s.From = &v
	return s
}

func (s *GetLogsRequest) SetLine(v int64) *GetLogsRequest {
	s.Line = &v
	return s
}

func (s *GetLogsRequest) SetOffset(v int64) *GetLogsRequest {
	s.Offset = &v
	return s
}

func (s *GetLogsRequest) SetPowerSql(v bool) *GetLogsRequest {
	s.PowerSql = &v
	return s
}

func (s *GetLogsRequest) SetQuery(v string) *GetLogsRequest {
	s.Query = &v
	return s
}

func (s *GetLogsRequest) SetReverse(v bool) *GetLogsRequest {
	s.Reverse = &v
	return s
}

func (s *GetLogsRequest) SetTo(v int64) *GetLogsRequest {
	s.To = &v
	return s
}

func (s *GetLogsRequest) SetTopic(v string) *GetLogsRequest {
	s.Topic = &v
	return s
}

type GetLogsResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       []map[string]interface{} `json:"body,omitempty" xml:"body,omitempty" require:"true" type:"Repeated"`
}

func (s GetLogsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetLogsResponse) GoString() string {
	return s.String()
}

func (s *GetLogsResponse) SetHeaders(v map[string]*string) *GetLogsResponse {
	s.Headers = v
	return s
}

func (s *GetLogsResponse) SetStatusCode(v int32) *GetLogsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetLogsResponse) SetBody(v []map[string]interface{}) *GetLogsResponse {
	s.Body = v
	return s
}

type GetMachineGroupResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *MachineGroup      `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetMachineGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s GetMachineGroupResponse) GoString() string {
	return s.String()
}

func (s *GetMachineGroupResponse) SetHeaders(v map[string]*string) *GetMachineGroupResponse {
	s.Headers = v
	return s
}

func (s *GetMachineGroupResponse) SetStatusCode(v int32) *GetMachineGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMachineGroupResponse) SetBody(v *MachineGroup) *GetMachineGroupResponse {
	s.Body = v
	return s
}

type GetProjectResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *Project           `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetProjectResponse) String() string {
	return tea.Prettify(s)
}

func (s GetProjectResponse) GoString() string {
	return s.String()
}

func (s *GetProjectResponse) SetHeaders(v map[string]*string) *GetProjectResponse {
	s.Headers = v
	return s
}

func (s *GetProjectResponse) SetStatusCode(v int32) *GetProjectResponse {
	s.StatusCode = &v
	return s
}

func (s *GetProjectResponse) SetBody(v *Project) *GetProjectResponse {
	s.Body = v
	return s
}

type GetProjectLogsRequest struct {
	PowerSql *bool   `json:"powerSql,omitempty" xml:"powerSql,omitempty"`
	Query    *string `json:"query,omitempty" xml:"query,omitempty"`
}

func (s GetProjectLogsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetProjectLogsRequest) GoString() string {
	return s.String()
}

func (s *GetProjectLogsRequest) SetPowerSql(v bool) *GetProjectLogsRequest {
	s.PowerSql = &v
	return s
}

func (s *GetProjectLogsRequest) SetQuery(v string) *GetProjectLogsRequest {
	s.Query = &v
	return s
}

type GetProjectLogsResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       []map[string]*string `json:"body,omitempty" xml:"body,omitempty" require:"true" type:"Repeated"`
}

func (s GetProjectLogsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetProjectLogsResponse) GoString() string {
	return s.String()
}

func (s *GetProjectLogsResponse) SetHeaders(v map[string]*string) *GetProjectLogsResponse {
	s.Headers = v
	return s
}

func (s *GetProjectLogsResponse) SetStatusCode(v int32) *GetProjectLogsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetProjectLogsResponse) SetBody(v []map[string]*string) *GetProjectLogsResponse {
	s.Body = v
	return s
}

type GetProjectPolicyResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *string            `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetProjectPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s GetProjectPolicyResponse) GoString() string {
	return s.String()
}

func (s *GetProjectPolicyResponse) SetHeaders(v map[string]*string) *GetProjectPolicyResponse {
	s.Headers = v
	return s
}

func (s *GetProjectPolicyResponse) SetStatusCode(v int32) *GetProjectPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *GetProjectPolicyResponse) SetBody(v string) *GetProjectPolicyResponse {
	s.Body = &v
	return s
}

type GetSavedSearchResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SavedSearch       `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetSavedSearchResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSavedSearchResponse) GoString() string {
	return s.String()
}

func (s *GetSavedSearchResponse) SetHeaders(v map[string]*string) *GetSavedSearchResponse {
	s.Headers = v
	return s
}

func (s *GetSavedSearchResponse) SetStatusCode(v int32) *GetSavedSearchResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSavedSearchResponse) SetBody(v *SavedSearch) *GetSavedSearchResponse {
	s.Body = v
	return s
}

type GetShipperStatusRequest struct {
	From   *int64  `json:"from,omitempty" xml:"from,omitempty"`
	Offset *int32  `json:"offset,omitempty" xml:"offset,omitempty"`
	Size   *int32  `json:"size,omitempty" xml:"size,omitempty"`
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	To     *int64  `json:"to,omitempty" xml:"to,omitempty"`
}

func (s GetShipperStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s GetShipperStatusRequest) GoString() string {
	return s.String()
}

func (s *GetShipperStatusRequest) SetFrom(v int64) *GetShipperStatusRequest {
	s.From = &v
	return s
}

func (s *GetShipperStatusRequest) SetOffset(v int32) *GetShipperStatusRequest {
	s.Offset = &v
	return s
}

func (s *GetShipperStatusRequest) SetSize(v int32) *GetShipperStatusRequest {
	s.Size = &v
	return s
}

func (s *GetShipperStatusRequest) SetStatus(v string) *GetShipperStatusRequest {
	s.Status = &v
	return s
}

func (s *GetShipperStatusRequest) SetTo(v int64) *GetShipperStatusRequest {
	s.To = &v
	return s
}

type GetShipperStatusResponseBody struct {
	Count      *int64                                  `json:"count,omitempty" xml:"count,omitempty"`
	Statistics *GetShipperStatusResponseBodyStatistics `json:"statistics,omitempty" xml:"statistics,omitempty" type:"Struct"`
	Tasks      *GetShipperStatusResponseBodyTasks      `json:"tasks,omitempty" xml:"tasks,omitempty" type:"Struct"`
	Total      *int64                                  `json:"total,omitempty" xml:"total,omitempty"`
}

func (s GetShipperStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetShipperStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetShipperStatusResponseBody) SetCount(v int64) *GetShipperStatusResponseBody {
	s.Count = &v
	return s
}

func (s *GetShipperStatusResponseBody) SetStatistics(v *GetShipperStatusResponseBodyStatistics) *GetShipperStatusResponseBody {
	s.Statistics = v
	return s
}

func (s *GetShipperStatusResponseBody) SetTasks(v *GetShipperStatusResponseBodyTasks) *GetShipperStatusResponseBody {
	s.Tasks = v
	return s
}

func (s *GetShipperStatusResponseBody) SetTotal(v int64) *GetShipperStatusResponseBody {
	s.Total = &v
	return s
}

type GetShipperStatusResponseBodyStatistics struct {
	Fail    *int64 `json:"fail,omitempty" xml:"fail,omitempty"`
	Running *int64 `json:"running,omitempty" xml:"running,omitempty"`
	Success *int64 `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetShipperStatusResponseBodyStatistics) String() string {
	return tea.Prettify(s)
}

func (s GetShipperStatusResponseBodyStatistics) GoString() string {
	return s.String()
}

func (s *GetShipperStatusResponseBodyStatistics) SetFail(v int64) *GetShipperStatusResponseBodyStatistics {
	s.Fail = &v
	return s
}

func (s *GetShipperStatusResponseBodyStatistics) SetRunning(v int64) *GetShipperStatusResponseBodyStatistics {
	s.Running = &v
	return s
}

func (s *GetShipperStatusResponseBodyStatistics) SetSuccess(v int64) *GetShipperStatusResponseBodyStatistics {
	s.Success = &v
	return s
}

type GetShipperStatusResponseBodyTasks struct {
	Id                      *string `json:"id,omitempty" xml:"id,omitempty"`
	TaskCode                *string `json:"taskCode,omitempty" xml:"taskCode,omitempty"`
	TaskCreateTime          *int64  `json:"taskCreateTime,omitempty" xml:"taskCreateTime,omitempty"`
	TaskDataLines           *int32  `json:"taskDataLines,omitempty" xml:"taskDataLines,omitempty"`
	TaskFinishTime          *int64  `json:"taskFinishTime,omitempty" xml:"taskFinishTime,omitempty"`
	TaskLastDataReceiveTime *int64  `json:"taskLastDataReceiveTime,omitempty" xml:"taskLastDataReceiveTime,omitempty"`
	TaskMessage             *string `json:"taskMessage,omitempty" xml:"taskMessage,omitempty"`
	TaskStatus              *string `json:"taskStatus,omitempty" xml:"taskStatus,omitempty"`
}

func (s GetShipperStatusResponseBodyTasks) String() string {
	return tea.Prettify(s)
}

func (s GetShipperStatusResponseBodyTasks) GoString() string {
	return s.String()
}

func (s *GetShipperStatusResponseBodyTasks) SetId(v string) *GetShipperStatusResponseBodyTasks {
	s.Id = &v
	return s
}

func (s *GetShipperStatusResponseBodyTasks) SetTaskCode(v string) *GetShipperStatusResponseBodyTasks {
	s.TaskCode = &v
	return s
}

func (s *GetShipperStatusResponseBodyTasks) SetTaskCreateTime(v int64) *GetShipperStatusResponseBodyTasks {
	s.TaskCreateTime = &v
	return s
}

func (s *GetShipperStatusResponseBodyTasks) SetTaskDataLines(v int32) *GetShipperStatusResponseBodyTasks {
	s.TaskDataLines = &v
	return s
}

func (s *GetShipperStatusResponseBodyTasks) SetTaskFinishTime(v int64) *GetShipperStatusResponseBodyTasks {
	s.TaskFinishTime = &v
	return s
}

func (s *GetShipperStatusResponseBodyTasks) SetTaskLastDataReceiveTime(v int64) *GetShipperStatusResponseBodyTasks {
	s.TaskLastDataReceiveTime = &v
	return s
}

func (s *GetShipperStatusResponseBodyTasks) SetTaskMessage(v string) *GetShipperStatusResponseBodyTasks {
	s.TaskMessage = &v
	return s
}

func (s *GetShipperStatusResponseBodyTasks) SetTaskStatus(v string) *GetShipperStatusResponseBodyTasks {
	s.TaskStatus = &v
	return s
}

type GetShipperStatusResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetShipperStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetShipperStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s GetShipperStatusResponse) GoString() string {
	return s.String()
}

func (s *GetShipperStatusResponse) SetHeaders(v map[string]*string) *GetShipperStatusResponse {
	s.Headers = v
	return s
}

func (s *GetShipperStatusResponse) SetStatusCode(v int32) *GetShipperStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetShipperStatusResponse) SetBody(v *GetShipperStatusResponseBody) *GetShipperStatusResponse {
	s.Body = v
	return s
}

type ListAlertResponseBody struct {
	Count   *int64   `json:"count,omitempty" xml:"count,omitempty"`
	Results []*Alert `json:"results,omitempty" xml:"results,omitempty" type:"Repeated"`
	Total   *int64   `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListAlertResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAlertResponseBody) GoString() string {
	return s.String()
}

func (s *ListAlertResponseBody) SetCount(v int64) *ListAlertResponseBody {
	s.Count = &v
	return s
}

func (s *ListAlertResponseBody) SetResults(v []*Alert) *ListAlertResponseBody {
	s.Results = v
	return s
}

func (s *ListAlertResponseBody) SetTotal(v int64) *ListAlertResponseBody {
	s.Total = &v
	return s
}

type ListAlertResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListAlertResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListAlertResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAlertResponse) GoString() string {
	return s.String()
}

func (s *ListAlertResponse) SetHeaders(v map[string]*string) *ListAlertResponse {
	s.Headers = v
	return s
}

func (s *ListAlertResponse) SetStatusCode(v int32) *ListAlertResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAlertResponse) SetBody(v *ListAlertResponseBody) *ListAlertResponse {
	s.Body = v
	return s
}

type ListConsumerGroupResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       []*ConsumerGroup   `json:"body,omitempty" xml:"body,omitempty" require:"true" type:"Repeated"`
}

func (s ListConsumerGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s ListConsumerGroupResponse) GoString() string {
	return s.String()
}

func (s *ListConsumerGroupResponse) SetHeaders(v map[string]*string) *ListConsumerGroupResponse {
	s.Headers = v
	return s
}

func (s *ListConsumerGroupResponse) SetStatusCode(v int32) *ListConsumerGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *ListConsumerGroupResponse) SetBody(v []*ConsumerGroup) *ListConsumerGroupResponse {
	s.Body = v
	return s
}

type ListDomainsRequest struct {
	DomainName *string `json:"domainName,omitempty" xml:"domainName,omitempty"`
	Offset     *int32  `json:"offset,omitempty" xml:"offset,omitempty"`
	Size       *int32  `json:"size,omitempty" xml:"size,omitempty"`
}

func (s ListDomainsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDomainsRequest) GoString() string {
	return s.String()
}

func (s *ListDomainsRequest) SetDomainName(v string) *ListDomainsRequest {
	s.DomainName = &v
	return s
}

func (s *ListDomainsRequest) SetOffset(v int32) *ListDomainsRequest {
	s.Offset = &v
	return s
}

func (s *ListDomainsRequest) SetSize(v int32) *ListDomainsRequest {
	s.Size = &v
	return s
}

type ListDomainsResponseBody struct {
	Count   *int64    `json:"count,omitempty" xml:"count,omitempty"`
	Domains []*string `json:"domains,omitempty" xml:"domains,omitempty" type:"Repeated"`
	Total   *int64    `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListDomainsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDomainsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDomainsResponseBody) SetCount(v int64) *ListDomainsResponseBody {
	s.Count = &v
	return s
}

func (s *ListDomainsResponseBody) SetDomains(v []*string) *ListDomainsResponseBody {
	s.Domains = v
	return s
}

func (s *ListDomainsResponseBody) SetTotal(v int64) *ListDomainsResponseBody {
	s.Total = &v
	return s
}

type ListDomainsResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListDomainsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListDomainsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDomainsResponse) GoString() string {
	return s.String()
}

func (s *ListDomainsResponse) SetHeaders(v map[string]*string) *ListDomainsResponse {
	s.Headers = v
	return s
}

func (s *ListDomainsResponse) SetStatusCode(v int32) *ListDomainsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDomainsResponse) SetBody(v *ListDomainsResponseBody) *ListDomainsResponse {
	s.Body = v
	return s
}

type ListETLJobsResponseBody struct {
	Results []*EtlJob `json:"results,omitempty" xml:"results,omitempty" type:"Repeated"`
	Total   *int32    `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListETLJobsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListETLJobsResponseBody) GoString() string {
	return s.String()
}

func (s *ListETLJobsResponseBody) SetResults(v []*EtlJob) *ListETLJobsResponseBody {
	s.Results = v
	return s
}

func (s *ListETLJobsResponseBody) SetTotal(v int32) *ListETLJobsResponseBody {
	s.Total = &v
	return s
}

type ListETLJobsResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListETLJobsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListETLJobsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListETLJobsResponse) GoString() string {
	return s.String()
}

func (s *ListETLJobsResponse) SetHeaders(v map[string]*string) *ListETLJobsResponse {
	s.Headers = v
	return s
}

func (s *ListETLJobsResponse) SetStatusCode(v int32) *ListETLJobsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListETLJobsResponse) SetBody(v *ListETLJobsResponseBody) *ListETLJobsResponse {
	s.Body = v
	return s
}

type ListExternalStoreRequest struct {
	ExternalStoreName *string `json:"externalStoreName,omitempty" xml:"externalStoreName,omitempty"`
	Offset            *int32  `json:"offset,omitempty" xml:"offset,omitempty"`
	Sizs              *int32  `json:"sizs,omitempty" xml:"sizs,omitempty"`
}

func (s ListExternalStoreRequest) String() string {
	return tea.Prettify(s)
}

func (s ListExternalStoreRequest) GoString() string {
	return s.String()
}

func (s *ListExternalStoreRequest) SetExternalStoreName(v string) *ListExternalStoreRequest {
	s.ExternalStoreName = &v
	return s
}

func (s *ListExternalStoreRequest) SetOffset(v int32) *ListExternalStoreRequest {
	s.Offset = &v
	return s
}

func (s *ListExternalStoreRequest) SetSizs(v int32) *ListExternalStoreRequest {
	s.Sizs = &v
	return s
}

type ListExternalStoreResponseBody struct {
	Count          *int32           `json:"count,omitempty" xml:"count,omitempty"`
	Externalstores []*ExternalStore `json:"externalstores,omitempty" xml:"externalstores,omitempty" type:"Repeated"`
	Total          *int32           `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListExternalStoreResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListExternalStoreResponseBody) GoString() string {
	return s.String()
}

func (s *ListExternalStoreResponseBody) SetCount(v int32) *ListExternalStoreResponseBody {
	s.Count = &v
	return s
}

func (s *ListExternalStoreResponseBody) SetExternalstores(v []*ExternalStore) *ListExternalStoreResponseBody {
	s.Externalstores = v
	return s
}

func (s *ListExternalStoreResponseBody) SetTotal(v int32) *ListExternalStoreResponseBody {
	s.Total = &v
	return s
}

type ListExternalStoreResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListExternalStoreResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListExternalStoreResponse) String() string {
	return tea.Prettify(s)
}

func (s ListExternalStoreResponse) GoString() string {
	return s.String()
}

func (s *ListExternalStoreResponse) SetHeaders(v map[string]*string) *ListExternalStoreResponse {
	s.Headers = v
	return s
}

func (s *ListExternalStoreResponse) SetStatusCode(v int32) *ListExternalStoreResponse {
	s.StatusCode = &v
	return s
}

func (s *ListExternalStoreResponse) SetBody(v *ListExternalStoreResponseBody) *ListExternalStoreResponse {
	s.Body = v
	return s
}

type ListJobsResponseBody struct {
	Results interface{} `json:"results,omitempty" xml:"results,omitempty"`
	Total   *int32      `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListJobsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListJobsResponseBody) GoString() string {
	return s.String()
}

func (s *ListJobsResponseBody) SetResults(v interface{}) *ListJobsResponseBody {
	s.Results = v
	return s
}

func (s *ListJobsResponseBody) SetTotal(v int32) *ListJobsResponseBody {
	s.Total = &v
	return s
}

type ListJobsResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListJobsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListJobsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListJobsResponse) GoString() string {
	return s.String()
}

func (s *ListJobsResponse) SetHeaders(v map[string]*string) *ListJobsResponse {
	s.Headers = v
	return s
}

func (s *ListJobsResponse) SetStatusCode(v int32) *ListJobsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListJobsResponse) SetBody(v *ListJobsResponseBody) *ListJobsResponse {
	s.Body = v
	return s
}

type ListLogStoresRequest struct {
	LogstoreName  *string `json:"logstoreName,omitempty" xml:"logstoreName,omitempty"`
	Mode          *string `json:"mode,omitempty" xml:"mode,omitempty"`
	Offset        *int32  `json:"offset,omitempty" xml:"offset,omitempty"`
	Size          *int32  `json:"size,omitempty" xml:"size,omitempty"`
	TelemetryType *string `json:"telemetryType,omitempty" xml:"telemetryType,omitempty"`
}

func (s ListLogStoresRequest) String() string {
	return tea.Prettify(s)
}

func (s ListLogStoresRequest) GoString() string {
	return s.String()
}

func (s *ListLogStoresRequest) SetLogstoreName(v string) *ListLogStoresRequest {
	s.LogstoreName = &v
	return s
}

func (s *ListLogStoresRequest) SetMode(v string) *ListLogStoresRequest {
	s.Mode = &v
	return s
}

func (s *ListLogStoresRequest) SetOffset(v int32) *ListLogStoresRequest {
	s.Offset = &v
	return s
}

func (s *ListLogStoresRequest) SetSize(v int32) *ListLogStoresRequest {
	s.Size = &v
	return s
}

func (s *ListLogStoresRequest) SetTelemetryType(v string) *ListLogStoresRequest {
	s.TelemetryType = &v
	return s
}

type ListLogStoresResponseBody struct {
	Logstores []*string `json:"logstores,omitempty" xml:"logstores,omitempty" type:"Repeated"`
	Total     *int32    `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListLogStoresResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListLogStoresResponseBody) GoString() string {
	return s.String()
}

func (s *ListLogStoresResponseBody) SetLogstores(v []*string) *ListLogStoresResponseBody {
	s.Logstores = v
	return s
}

func (s *ListLogStoresResponseBody) SetTotal(v int32) *ListLogStoresResponseBody {
	s.Total = &v
	return s
}

type ListLogStoresResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListLogStoresResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListLogStoresResponse) String() string {
	return tea.Prettify(s)
}

func (s ListLogStoresResponse) GoString() string {
	return s.String()
}

func (s *ListLogStoresResponse) SetHeaders(v map[string]*string) *ListLogStoresResponse {
	s.Headers = v
	return s
}

func (s *ListLogStoresResponse) SetStatusCode(v int32) *ListLogStoresResponse {
	s.StatusCode = &v
	return s
}

func (s *ListLogStoresResponse) SetBody(v *ListLogStoresResponseBody) *ListLogStoresResponse {
	s.Body = v
	return s
}

type ListMachineGroupRequest struct {
	GroupName *string `json:"groupName,omitempty" xml:"groupName,omitempty"`
	Offset    *int32  `json:"offset,omitempty" xml:"offset,omitempty"`
	Size      *int32  `json:"size,omitempty" xml:"size,omitempty"`
}

func (s ListMachineGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s ListMachineGroupRequest) GoString() string {
	return s.String()
}

func (s *ListMachineGroupRequest) SetGroupName(v string) *ListMachineGroupRequest {
	s.GroupName = &v
	return s
}

func (s *ListMachineGroupRequest) SetOffset(v int32) *ListMachineGroupRequest {
	s.Offset = &v
	return s
}

func (s *ListMachineGroupRequest) SetSize(v int32) *ListMachineGroupRequest {
	s.Size = &v
	return s
}

type ListMachineGroupResponseBody struct {
	Count         *int32    `json:"count,omitempty" xml:"count,omitempty"`
	Machinegroups []*string `json:"machinegroups,omitempty" xml:"machinegroups,omitempty" type:"Repeated"`
	Total         *int32    `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListMachineGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListMachineGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ListMachineGroupResponseBody) SetCount(v int32) *ListMachineGroupResponseBody {
	s.Count = &v
	return s
}

func (s *ListMachineGroupResponseBody) SetMachinegroups(v []*string) *ListMachineGroupResponseBody {
	s.Machinegroups = v
	return s
}

func (s *ListMachineGroupResponseBody) SetTotal(v int32) *ListMachineGroupResponseBody {
	s.Total = &v
	return s
}

type ListMachineGroupResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListMachineGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListMachineGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s ListMachineGroupResponse) GoString() string {
	return s.String()
}

func (s *ListMachineGroupResponse) SetHeaders(v map[string]*string) *ListMachineGroupResponse {
	s.Headers = v
	return s
}

func (s *ListMachineGroupResponse) SetStatusCode(v int32) *ListMachineGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMachineGroupResponse) SetBody(v *ListMachineGroupResponseBody) *ListMachineGroupResponse {
	s.Body = v
	return s
}

type ListMachinesRequest struct {
	Offset *int32 `json:"offset,omitempty" xml:"offset,omitempty"`
	Size   *int32 `json:"size,omitempty" xml:"size,omitempty"`
}

func (s ListMachinesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListMachinesRequest) GoString() string {
	return s.String()
}

func (s *ListMachinesRequest) SetOffset(v int32) *ListMachinesRequest {
	s.Offset = &v
	return s
}

func (s *ListMachinesRequest) SetSize(v int32) *ListMachinesRequest {
	s.Size = &v
	return s
}

type ListMachinesResponseBody struct {
	Count    *int32     `json:"count,omitempty" xml:"count,omitempty"`
	Machines []*Machine `json:"machines,omitempty" xml:"machines,omitempty" type:"Repeated"`
	Total    *int32     `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListMachinesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListMachinesResponseBody) GoString() string {
	return s.String()
}

func (s *ListMachinesResponseBody) SetCount(v int32) *ListMachinesResponseBody {
	s.Count = &v
	return s
}

func (s *ListMachinesResponseBody) SetMachines(v []*Machine) *ListMachinesResponseBody {
	s.Machines = v
	return s
}

func (s *ListMachinesResponseBody) SetTotal(v int32) *ListMachinesResponseBody {
	s.Total = &v
	return s
}

type ListMachinesResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListMachinesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListMachinesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListMachinesResponse) GoString() string {
	return s.String()
}

func (s *ListMachinesResponse) SetHeaders(v map[string]*string) *ListMachinesResponse {
	s.Headers = v
	return s
}

func (s *ListMachinesResponse) SetStatusCode(v int32) *ListMachinesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMachinesResponse) SetBody(v *ListMachinesResponseBody) *ListMachinesResponse {
	s.Body = v
	return s
}

type ListProjectRequest struct {
	Offset      *int32  `json:"offset,omitempty" xml:"offset,omitempty"`
	ProjectName *string `json:"projectName,omitempty" xml:"projectName,omitempty"`
	Size        *int32  `json:"size,omitempty" xml:"size,omitempty"`
}

func (s ListProjectRequest) String() string {
	return tea.Prettify(s)
}

func (s ListProjectRequest) GoString() string {
	return s.String()
}

func (s *ListProjectRequest) SetOffset(v int32) *ListProjectRequest {
	s.Offset = &v
	return s
}

func (s *ListProjectRequest) SetProjectName(v string) *ListProjectRequest {
	s.ProjectName = &v
	return s
}

func (s *ListProjectRequest) SetSize(v int32) *ListProjectRequest {
	s.Size = &v
	return s
}

type ListProjectResponseBody struct {
	Count    *int64     `json:"count,omitempty" xml:"count,omitempty"`
	Projects []*Project `json:"projects,omitempty" xml:"projects,omitempty" type:"Repeated"`
	Total    *int64     `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListProjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListProjectResponseBody) GoString() string {
	return s.String()
}

func (s *ListProjectResponseBody) SetCount(v int64) *ListProjectResponseBody {
	s.Count = &v
	return s
}

func (s *ListProjectResponseBody) SetProjects(v []*Project) *ListProjectResponseBody {
	s.Projects = v
	return s
}

func (s *ListProjectResponseBody) SetTotal(v int64) *ListProjectResponseBody {
	s.Total = &v
	return s
}

type ListProjectResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListProjectResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListProjectResponse) String() string {
	return tea.Prettify(s)
}

func (s ListProjectResponse) GoString() string {
	return s.String()
}

func (s *ListProjectResponse) SetHeaders(v map[string]*string) *ListProjectResponse {
	s.Headers = v
	return s
}

func (s *ListProjectResponse) SetStatusCode(v int32) *ListProjectResponse {
	s.StatusCode = &v
	return s
}

func (s *ListProjectResponse) SetBody(v *ListProjectResponseBody) *ListProjectResponse {
	s.Body = v
	return s
}

type ListSavedSearchRequest struct {
	Offset *int32 `json:"offset,omitempty" xml:"offset,omitempty"`
	Size   *int32 `json:"size,omitempty" xml:"size,omitempty"`
}

func (s ListSavedSearchRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSavedSearchRequest) GoString() string {
	return s.String()
}

func (s *ListSavedSearchRequest) SetOffset(v int32) *ListSavedSearchRequest {
	s.Offset = &v
	return s
}

func (s *ListSavedSearchRequest) SetSize(v int32) *ListSavedSearchRequest {
	s.Size = &v
	return s
}

type ListSavedSearchResponseBody struct {
	Count            *int32         `json:"count,omitempty" xml:"count,omitempty"`
	SavedsearchItems []*SavedSearch `json:"savedsearchItems,omitempty" xml:"savedsearchItems,omitempty" type:"Repeated"`
	Total            *int32         `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListSavedSearchResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSavedSearchResponseBody) GoString() string {
	return s.String()
}

func (s *ListSavedSearchResponseBody) SetCount(v int32) *ListSavedSearchResponseBody {
	s.Count = &v
	return s
}

func (s *ListSavedSearchResponseBody) SetSavedsearchItems(v []*SavedSearch) *ListSavedSearchResponseBody {
	s.SavedsearchItems = v
	return s
}

func (s *ListSavedSearchResponseBody) SetTotal(v int32) *ListSavedSearchResponseBody {
	s.Total = &v
	return s
}

type ListSavedSearchResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListSavedSearchResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListSavedSearchResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSavedSearchResponse) GoString() string {
	return s.String()
}

func (s *ListSavedSearchResponse) SetHeaders(v map[string]*string) *ListSavedSearchResponse {
	s.Headers = v
	return s
}

func (s *ListSavedSearchResponse) SetStatusCode(v int32) *ListSavedSearchResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSavedSearchResponse) SetBody(v *ListSavedSearchResponseBody) *ListSavedSearchResponse {
	s.Body = v
	return s
}

type ListShardsResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       []*Shard           `json:"body,omitempty" xml:"body,omitempty" require:"true" type:"Repeated"`
}

func (s ListShardsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListShardsResponse) GoString() string {
	return s.String()
}

func (s *ListShardsResponse) SetHeaders(v map[string]*string) *ListShardsResponse {
	s.Headers = v
	return s
}

func (s *ListShardsResponse) SetStatusCode(v int32) *ListShardsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListShardsResponse) SetBody(v []*Shard) *ListShardsResponse {
	s.Body = v
	return s
}

type ListShipperResponseBody struct {
	Count   *int64    `json:"count,omitempty" xml:"count,omitempty"`
	Shipper []*string `json:"shipper,omitempty" xml:"shipper,omitempty" type:"Repeated"`
	Total   *int64    `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListShipperResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListShipperResponseBody) GoString() string {
	return s.String()
}

func (s *ListShipperResponseBody) SetCount(v int64) *ListShipperResponseBody {
	s.Count = &v
	return s
}

func (s *ListShipperResponseBody) SetShipper(v []*string) *ListShipperResponseBody {
	s.Shipper = v
	return s
}

func (s *ListShipperResponseBody) SetTotal(v int64) *ListShipperResponseBody {
	s.Total = &v
	return s
}

type ListShipperResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListShipperResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListShipperResponse) String() string {
	return tea.Prettify(s)
}

func (s ListShipperResponse) GoString() string {
	return s.String()
}

func (s *ListShipperResponse) SetHeaders(v map[string]*string) *ListShipperResponse {
	s.Headers = v
	return s
}

func (s *ListShipperResponse) SetStatusCode(v int32) *ListShipperResponse {
	s.StatusCode = &v
	return s
}

func (s *ListShipperResponse) SetBody(v *ListShipperResponseBody) *ListShipperResponse {
	s.Body = v
	return s
}

type ListTagResourcesRequest struct {
	ResourceId   []*string                      `json:"resourceId,omitempty" xml:"resourceId,omitempty" type:"Repeated"`
	ResourceType *string                        `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
	Tags         []*ListTagResourcesRequestTags `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
}

func (s ListTagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesRequest) GoString() string {
	return s.String()
}

func (s *ListTagResourcesRequest) SetResourceId(v []*string) *ListTagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *ListTagResourcesRequest) SetResourceType(v string) *ListTagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesRequest) SetTags(v []*ListTagResourcesRequestTags) *ListTagResourcesRequest {
	s.Tags = v
	return s
}

type ListTagResourcesRequestTags struct {
	Key   *string `json:"key,omitempty" xml:"key,omitempty"`
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s ListTagResourcesRequestTags) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesRequestTags) GoString() string {
	return s.String()
}

func (s *ListTagResourcesRequestTags) SetKey(v string) *ListTagResourcesRequestTags {
	s.Key = &v
	return s
}

func (s *ListTagResourcesRequestTags) SetValue(v string) *ListTagResourcesRequestTags {
	s.Value = &v
	return s
}

type ListTagResourcesShrinkRequest struct {
	ResourceIdShrink *string `json:"resourceId,omitempty" xml:"resourceId,omitempty"`
	ResourceType     *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
	TagsShrink       *string `json:"tags,omitempty" xml:"tags,omitempty"`
}

func (s ListTagResourcesShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListTagResourcesShrinkRequest) SetResourceIdShrink(v string) *ListTagResourcesShrinkRequest {
	s.ResourceIdShrink = &v
	return s
}

func (s *ListTagResourcesShrinkRequest) SetResourceType(v string) *ListTagResourcesShrinkRequest {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesShrinkRequest) SetTagsShrink(v string) *ListTagResourcesShrinkRequest {
	s.TagsShrink = &v
	return s
}

type ListTagResourcesResponseBody struct {
	NextToken    *string                                     `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	TagResources []*ListTagResourcesResponseBodyTagResources `json:"tagResources,omitempty" xml:"tagResources,omitempty" type:"Repeated"`
}

func (s ListTagResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBody) SetNextToken(v string) *ListTagResourcesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTagResourcesResponseBody) SetTagResources(v []*ListTagResourcesResponseBodyTagResources) *ListTagResourcesResponseBody {
	s.TagResources = v
	return s
}

type ListTagResourcesResponseBodyTagResources struct {
	ResourceId   *string `json:"resourceId,omitempty" xml:"resourceId,omitempty"`
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
	TagKey       *string `json:"tagKey,omitempty" xml:"tagKey,omitempty"`
	TagValue     *string `json:"tagValue,omitempty" xml:"tagValue,omitempty"`
}

func (s ListTagResourcesResponseBodyTagResources) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseBodyTagResources) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBodyTagResources) SetResourceId(v string) *ListTagResourcesResponseBodyTagResources {
	s.ResourceId = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResources) SetResourceType(v string) *ListTagResourcesResponseBodyTagResources {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResources) SetTagKey(v string) *ListTagResourcesResponseBodyTagResources {
	s.TagKey = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResources) SetTagValue(v string) *ListTagResourcesResponseBodyTagResources {
	s.TagValue = &v
	return s
}

type ListTagResourcesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListTagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListTagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponse) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponse) SetHeaders(v map[string]*string) *ListTagResourcesResponse {
	s.Headers = v
	return s
}

func (s *ListTagResourcesResponse) SetStatusCode(v int32) *ListTagResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTagResourcesResponse) SetBody(v *ListTagResourcesResponseBody) *ListTagResourcesResponse {
	s.Body = v
	return s
}

type PullDataRequest struct {
	Count     *string `json:"count,omitempty" xml:"count,omitempty"`
	Cursor    *string `json:"cursor,omitempty" xml:"cursor,omitempty"`
	EndCursor *string `json:"endCursor,omitempty" xml:"endCursor,omitempty"`
}

func (s PullDataRequest) String() string {
	return tea.Prettify(s)
}

func (s PullDataRequest) GoString() string {
	return s.String()
}

func (s *PullDataRequest) SetCount(v string) *PullDataRequest {
	s.Count = &v
	return s
}

func (s *PullDataRequest) SetCursor(v string) *PullDataRequest {
	s.Cursor = &v
	return s
}

func (s *PullDataRequest) SetEndCursor(v string) *PullDataRequest {
	s.EndCursor = &v
	return s
}

type PullDataResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
}

func (s PullDataResponse) String() string {
	return tea.Prettify(s)
}

func (s PullDataResponse) GoString() string {
	return s.String()
}

func (s *PullDataResponse) SetHeaders(v map[string]*string) *PullDataResponse {
	s.Headers = v
	return s
}

func (s *PullDataResponse) SetStatusCode(v int32) *PullDataResponse {
	s.StatusCode = &v
	return s
}

type PutProjectPolicyRequest struct {
	Body *string `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PutProjectPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s PutProjectPolicyRequest) GoString() string {
	return s.String()
}

func (s *PutProjectPolicyRequest) SetBody(v string) *PutProjectPolicyRequest {
	s.Body = &v
	return s
}

type PutProjectPolicyResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
}

func (s PutProjectPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s PutProjectPolicyResponse) GoString() string {
	return s.String()
}

func (s *PutProjectPolicyResponse) SetHeaders(v map[string]*string) *PutProjectPolicyResponse {
	s.Headers = v
	return s
}

func (s *PutProjectPolicyResponse) SetStatusCode(v int32) *PutProjectPolicyResponse {
	s.StatusCode = &v
	return s
}

type PutWebtrackingRequest struct {
	Logs   []map[string]*string `json:"__logs__,omitempty" xml:"__logs__,omitempty" type:"Repeated"`
	Source *string              `json:"__source__,omitempty" xml:"__source__,omitempty"`
	Tags   map[string]*string   `json:"__tags__,omitempty" xml:"__tags__,omitempty"`
	Topic  *string              `json:"__topic__,omitempty" xml:"__topic__,omitempty"`
}

func (s PutWebtrackingRequest) String() string {
	return tea.Prettify(s)
}

func (s PutWebtrackingRequest) GoString() string {
	return s.String()
}

func (s *PutWebtrackingRequest) SetLogs(v []map[string]*string) *PutWebtrackingRequest {
	s.Logs = v
	return s
}

func (s *PutWebtrackingRequest) SetSource(v string) *PutWebtrackingRequest {
	s.Source = &v
	return s
}

func (s *PutWebtrackingRequest) SetTags(v map[string]*string) *PutWebtrackingRequest {
	s.Tags = v
	return s
}

func (s *PutWebtrackingRequest) SetTopic(v string) *PutWebtrackingRequest {
	s.Topic = &v
	return s
}

type PutWebtrackingResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
}

func (s PutWebtrackingResponse) String() string {
	return tea.Prettify(s)
}

func (s PutWebtrackingResponse) GoString() string {
	return s.String()
}

func (s *PutWebtrackingResponse) SetHeaders(v map[string]*string) *PutWebtrackingResponse {
	s.Headers = v
	return s
}

func (s *PutWebtrackingResponse) SetStatusCode(v int32) *PutWebtrackingResponse {
	s.StatusCode = &v
	return s
}

type RemoveConfigFromMachineGroupResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
}

func (s RemoveConfigFromMachineGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveConfigFromMachineGroupResponse) GoString() string {
	return s.String()
}

func (s *RemoveConfigFromMachineGroupResponse) SetHeaders(v map[string]*string) *RemoveConfigFromMachineGroupResponse {
	s.Headers = v
	return s
}

func (s *RemoveConfigFromMachineGroupResponse) SetStatusCode(v int32) *RemoveConfigFromMachineGroupResponse {
	s.StatusCode = &v
	return s
}

type SplitShardRequest struct {
	Key        *string `json:"key,omitempty" xml:"key,omitempty"`
	ShardCount *int32  `json:"shardCount,omitempty" xml:"shardCount,omitempty"`
}

func (s SplitShardRequest) String() string {
	return tea.Prettify(s)
}

func (s SplitShardRequest) GoString() string {
	return s.String()
}

func (s *SplitShardRequest) SetKey(v string) *SplitShardRequest {
	s.Key = &v
	return s
}

func (s *SplitShardRequest) SetShardCount(v int32) *SplitShardRequest {
	s.ShardCount = &v
	return s
}

type SplitShardResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       []*Shard           `json:"body,omitempty" xml:"body,omitempty" require:"true" type:"Repeated"`
}

func (s SplitShardResponse) String() string {
	return tea.Prettify(s)
}

func (s SplitShardResponse) GoString() string {
	return s.String()
}

func (s *SplitShardResponse) SetHeaders(v map[string]*string) *SplitShardResponse {
	s.Headers = v
	return s
}

func (s *SplitShardResponse) SetStatusCode(v int32) *SplitShardResponse {
	s.StatusCode = &v
	return s
}

func (s *SplitShardResponse) SetBody(v []*Shard) *SplitShardResponse {
	s.Body = v
	return s
}

type TagResourcesRequest struct {
	ResourceId   []*string                  `json:"resourceId,omitempty" xml:"resourceId,omitempty" type:"Repeated"`
	ResourceType *string                    `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
	Tags         []*TagResourcesRequestTags `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
}

func (s TagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesRequest) GoString() string {
	return s.String()
}

func (s *TagResourcesRequest) SetResourceId(v []*string) *TagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *TagResourcesRequest) SetResourceType(v string) *TagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *TagResourcesRequest) SetTags(v []*TagResourcesRequestTags) *TagResourcesRequest {
	s.Tags = v
	return s
}

type TagResourcesRequestTags struct {
	Key   *string `json:"key,omitempty" xml:"key,omitempty"`
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s TagResourcesRequestTags) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesRequestTags) GoString() string {
	return s.String()
}

func (s *TagResourcesRequestTags) SetKey(v string) *TagResourcesRequestTags {
	s.Key = &v
	return s
}

func (s *TagResourcesRequestTags) SetValue(v string) *TagResourcesRequestTags {
	s.Value = &v
	return s
}

type TagResourcesResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
}

func (s TagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesResponse) GoString() string {
	return s.String()
}

func (s *TagResourcesResponse) SetHeaders(v map[string]*string) *TagResourcesResponse {
	s.Headers = v
	return s
}

func (s *TagResourcesResponse) SetStatusCode(v int32) *TagResourcesResponse {
	s.StatusCode = &v
	return s
}

type UntagResourcesRequest struct {
	All          *bool     `json:"all,omitempty" xml:"all,omitempty"`
	ResourceId   *string   `json:"resourceId,omitempty" xml:"resourceId,omitempty"`
	ResourceType *string   `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
	Tags         []*string `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
}

func (s UntagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesRequest) GoString() string {
	return s.String()
}

func (s *UntagResourcesRequest) SetAll(v bool) *UntagResourcesRequest {
	s.All = &v
	return s
}

func (s *UntagResourcesRequest) SetResourceId(v string) *UntagResourcesRequest {
	s.ResourceId = &v
	return s
}

func (s *UntagResourcesRequest) SetResourceType(v string) *UntagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *UntagResourcesRequest) SetTags(v []*string) *UntagResourcesRequest {
	s.Tags = v
	return s
}

type UntagResourcesResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
}

func (s UntagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesResponse) GoString() string {
	return s.String()
}

func (s *UntagResourcesResponse) SetHeaders(v map[string]*string) *UntagResourcesResponse {
	s.Headers = v
	return s
}

func (s *UntagResourcesResponse) SetStatusCode(v int32) *UntagResourcesResponse {
	s.StatusCode = &v
	return s
}

type UpdateAlertRequest struct {
	Alert *Alert `json:"alert,omitempty" xml:"alert,omitempty"`
}

func (s UpdateAlertRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateAlertRequest) GoString() string {
	return s.String()
}

func (s *UpdateAlertRequest) SetAlert(v *Alert) *UpdateAlertRequest {
	s.Alert = v
	return s
}

type UpdateAlertResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
}

func (s UpdateAlertResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateAlertResponse) GoString() string {
	return s.String()
}

func (s *UpdateAlertResponse) SetHeaders(v map[string]*string) *UpdateAlertResponse {
	s.Headers = v
	return s
}

func (s *UpdateAlertResponse) SetStatusCode(v int32) *UpdateAlertResponse {
	s.StatusCode = &v
	return s
}

type UpdateConsumerGroupRequest struct {
	Order   *bool  `json:"order,omitempty" xml:"order,omitempty"`
	Timeout *int32 `json:"timeout,omitempty" xml:"timeout,omitempty"`
}

func (s UpdateConsumerGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateConsumerGroupRequest) GoString() string {
	return s.String()
}

func (s *UpdateConsumerGroupRequest) SetOrder(v bool) *UpdateConsumerGroupRequest {
	s.Order = &v
	return s
}

func (s *UpdateConsumerGroupRequest) SetTimeout(v int32) *UpdateConsumerGroupRequest {
	s.Timeout = &v
	return s
}

type UpdateConsumerGroupResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
}

func (s UpdateConsumerGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateConsumerGroupResponse) GoString() string {
	return s.String()
}

func (s *UpdateConsumerGroupResponse) SetHeaders(v map[string]*string) *UpdateConsumerGroupResponse {
	s.Headers = v
	return s
}

func (s *UpdateConsumerGroupResponse) SetStatusCode(v int32) *UpdateConsumerGroupResponse {
	s.StatusCode = &v
	return s
}

type UpdateEtlJobRequest struct {
	Enable            *bool              `json:"enable,omitempty" xml:"enable,omitempty"`
	FunctionConfig    *EtlFunctionConfig `json:"functionConfig,omitempty" xml:"functionConfig,omitempty"`
	FunctionParameter *string            `json:"functionParameter,omitempty" xml:"functionParameter,omitempty"`
	JobName           *string            `json:"jobName,omitempty" xml:"jobName,omitempty"`
	LogConfig         *EtlLogConfig      `json:"logConfig,omitempty" xml:"logConfig,omitempty"`
	SourceConfig      *EtlSourceConfig   `json:"sourceConfig,omitempty" xml:"sourceConfig,omitempty"`
	TriggerConfig     *EtlTriggerConfig  `json:"triggerConfig,omitempty" xml:"triggerConfig,omitempty"`
}

func (s UpdateEtlJobRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateEtlJobRequest) GoString() string {
	return s.String()
}

func (s *UpdateEtlJobRequest) SetEnable(v bool) *UpdateEtlJobRequest {
	s.Enable = &v
	return s
}

func (s *UpdateEtlJobRequest) SetFunctionConfig(v *EtlFunctionConfig) *UpdateEtlJobRequest {
	s.FunctionConfig = v
	return s
}

func (s *UpdateEtlJobRequest) SetFunctionParameter(v string) *UpdateEtlJobRequest {
	s.FunctionParameter = &v
	return s
}

func (s *UpdateEtlJobRequest) SetJobName(v string) *UpdateEtlJobRequest {
	s.JobName = &v
	return s
}

func (s *UpdateEtlJobRequest) SetLogConfig(v *EtlLogConfig) *UpdateEtlJobRequest {
	s.LogConfig = v
	return s
}

func (s *UpdateEtlJobRequest) SetSourceConfig(v *EtlSourceConfig) *UpdateEtlJobRequest {
	s.SourceConfig = v
	return s
}

func (s *UpdateEtlJobRequest) SetTriggerConfig(v *EtlTriggerConfig) *UpdateEtlJobRequest {
	s.TriggerConfig = v
	return s
}

type UpdateEtlJobResponseBody struct {
	Code      *int32                   `json:"code,omitempty" xml:"code,omitempty"`
	Data      []map[string]interface{} `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	Message   *string                  `json:"message,omitempty" xml:"message,omitempty"`
	RequestId *string                  `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *string                  `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateEtlJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateEtlJobResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateEtlJobResponseBody) SetCode(v int32) *UpdateEtlJobResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateEtlJobResponseBody) SetData(v []map[string]interface{}) *UpdateEtlJobResponseBody {
	s.Data = v
	return s
}

func (s *UpdateEtlJobResponseBody) SetMessage(v string) *UpdateEtlJobResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateEtlJobResponseBody) SetRequestId(v string) *UpdateEtlJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateEtlJobResponseBody) SetSuccess(v string) *UpdateEtlJobResponseBody {
	s.Success = &v
	return s
}

type UpdateEtlJobResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateEtlJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateEtlJobResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateEtlJobResponse) GoString() string {
	return s.String()
}

func (s *UpdateEtlJobResponse) SetHeaders(v map[string]*string) *UpdateEtlJobResponse {
	s.Headers = v
	return s
}

func (s *UpdateEtlJobResponse) SetStatusCode(v int32) *UpdateEtlJobResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateEtlJobResponse) SetBody(v *UpdateEtlJobResponseBody) *UpdateEtlJobResponse {
	s.Body = v
	return s
}

type UpdateIndexRequest struct {
	Keys               map[string]*KeysValue   `json:"keys,omitempty" xml:"keys,omitempty"`
	Line               *UpdateIndexRequestLine `json:"line,omitempty" xml:"line,omitempty" type:"Struct"`
	LogReduce          *bool                   `json:"log_reduce,omitempty" xml:"log_reduce,omitempty"`
	LogReduceBlackList []*string               `json:"log_reduce_black_list,omitempty" xml:"log_reduce_black_list,omitempty" type:"Repeated"`
	LogReduceWhiteList []*string               `json:"log_reduce_white_list,omitempty" xml:"log_reduce_white_list,omitempty" type:"Repeated"`
	MaxTextLen         *int32                  `json:"max_text_len,omitempty" xml:"max_text_len,omitempty"`
	Ttl                *int32                  `json:"ttl,omitempty" xml:"ttl,omitempty"`
}

func (s UpdateIndexRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateIndexRequest) GoString() string {
	return s.String()
}

func (s *UpdateIndexRequest) SetKeys(v map[string]*KeysValue) *UpdateIndexRequest {
	s.Keys = v
	return s
}

func (s *UpdateIndexRequest) SetLine(v *UpdateIndexRequestLine) *UpdateIndexRequest {
	s.Line = v
	return s
}

func (s *UpdateIndexRequest) SetLogReduce(v bool) *UpdateIndexRequest {
	s.LogReduce = &v
	return s
}

func (s *UpdateIndexRequest) SetLogReduceBlackList(v []*string) *UpdateIndexRequest {
	s.LogReduceBlackList = v
	return s
}

func (s *UpdateIndexRequest) SetLogReduceWhiteList(v []*string) *UpdateIndexRequest {
	s.LogReduceWhiteList = v
	return s
}

func (s *UpdateIndexRequest) SetMaxTextLen(v int32) *UpdateIndexRequest {
	s.MaxTextLen = &v
	return s
}

func (s *UpdateIndexRequest) SetTtl(v int32) *UpdateIndexRequest {
	s.Ttl = &v
	return s
}

type UpdateIndexRequestLine struct {
	CaseSensitive *bool     `json:"caseSensitive,omitempty" xml:"caseSensitive,omitempty"`
	Chn           *bool     `json:"chn,omitempty" xml:"chn,omitempty"`
	ExcludeKeys   []*string `json:"exclude_keys,omitempty" xml:"exclude_keys,omitempty" type:"Repeated"`
	IncludeKeys   []*string `json:"include_keys,omitempty" xml:"include_keys,omitempty" type:"Repeated"`
	Token         []*string `json:"token,omitempty" xml:"token,omitempty" type:"Repeated"`
}

func (s UpdateIndexRequestLine) String() string {
	return tea.Prettify(s)
}

func (s UpdateIndexRequestLine) GoString() string {
	return s.String()
}

func (s *UpdateIndexRequestLine) SetCaseSensitive(v bool) *UpdateIndexRequestLine {
	s.CaseSensitive = &v
	return s
}

func (s *UpdateIndexRequestLine) SetChn(v bool) *UpdateIndexRequestLine {
	s.Chn = &v
	return s
}

func (s *UpdateIndexRequestLine) SetExcludeKeys(v []*string) *UpdateIndexRequestLine {
	s.ExcludeKeys = v
	return s
}

func (s *UpdateIndexRequestLine) SetIncludeKeys(v []*string) *UpdateIndexRequestLine {
	s.IncludeKeys = v
	return s
}

func (s *UpdateIndexRequestLine) SetToken(v []*string) *UpdateIndexRequestLine {
	s.Token = v
	return s
}

type UpdateIndexResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
}

func (s UpdateIndexResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateIndexResponse) GoString() string {
	return s.String()
}

func (s *UpdateIndexResponse) SetHeaders(v map[string]*string) *UpdateIndexResponse {
	s.Headers = v
	return s
}

func (s *UpdateIndexResponse) SetStatusCode(v int32) *UpdateIndexResponse {
	s.StatusCode = &v
	return s
}

type UpdateJobRequest struct {
	Body *UpdateJobRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Struct"`
}

func (s UpdateJobRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateJobRequest) GoString() string {
	return s.String()
}

func (s *UpdateJobRequest) SetBody(v *UpdateJobRequestBody) *UpdateJobRequest {
	s.Body = v
	return s
}

type UpdateJobRequestBody struct {
	Enable            *string            `json:"enable,omitempty" xml:"enable,omitempty"`
	FunctionConfig    *EtlFunctionConfig `json:"functionConfig,omitempty" xml:"functionConfig,omitempty"`
	FunctionParameter *string            `json:"functionParameter,omitempty" xml:"functionParameter,omitempty"`
	JobName           *string            `json:"jobName,omitempty" xml:"jobName,omitempty"`
	LogConfig         *EtlLogConfig      `json:"logConfig,omitempty" xml:"logConfig,omitempty"`
	SourceConfig      *EtlSourceConfig   `json:"sourceConfig,omitempty" xml:"sourceConfig,omitempty"`
	TriggerConfig     *EtlTriggerConfig  `json:"triggerConfig,omitempty" xml:"triggerConfig,omitempty"`
}

func (s UpdateJobRequestBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateJobRequestBody) GoString() string {
	return s.String()
}

func (s *UpdateJobRequestBody) SetEnable(v string) *UpdateJobRequestBody {
	s.Enable = &v
	return s
}

func (s *UpdateJobRequestBody) SetFunctionConfig(v *EtlFunctionConfig) *UpdateJobRequestBody {
	s.FunctionConfig = v
	return s
}

func (s *UpdateJobRequestBody) SetFunctionParameter(v string) *UpdateJobRequestBody {
	s.FunctionParameter = &v
	return s
}

func (s *UpdateJobRequestBody) SetJobName(v string) *UpdateJobRequestBody {
	s.JobName = &v
	return s
}

func (s *UpdateJobRequestBody) SetLogConfig(v *EtlLogConfig) *UpdateJobRequestBody {
	s.LogConfig = v
	return s
}

func (s *UpdateJobRequestBody) SetSourceConfig(v *EtlSourceConfig) *UpdateJobRequestBody {
	s.SourceConfig = v
	return s
}

func (s *UpdateJobRequestBody) SetTriggerConfig(v *EtlTriggerConfig) *UpdateJobRequestBody {
	s.TriggerConfig = v
	return s
}

type UpdateJobShrinkRequest struct {
	BodyShrink *string `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateJobShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateJobShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateJobShrinkRequest) SetBodyShrink(v string) *UpdateJobShrinkRequest {
	s.BodyShrink = &v
	return s
}

type UpdateJobResponseBody struct {
	Code      *string `json:"code,omitempty" xml:"code,omitempty"`
	Data      *string `json:"data,omitempty" xml:"data,omitempty"`
	Message   *string `json:"message,omitempty" xml:"message,omitempty"`
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *string `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateJobResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateJobResponseBody) SetCode(v string) *UpdateJobResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateJobResponseBody) SetData(v string) *UpdateJobResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateJobResponseBody) SetMessage(v string) *UpdateJobResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateJobResponseBody) SetRequestId(v string) *UpdateJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateJobResponseBody) SetSuccess(v string) *UpdateJobResponseBody {
	s.Success = &v
	return s
}

type UpdateJobResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateJobResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateJobResponse) GoString() string {
	return s.String()
}

func (s *UpdateJobResponse) SetHeaders(v map[string]*string) *UpdateJobResponse {
	s.Headers = v
	return s
}

func (s *UpdateJobResponse) SetStatusCode(v int32) *UpdateJobResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateJobResponse) SetBody(v *UpdateJobResponseBody) *UpdateJobResponse {
	s.Body = v
	return s
}

type UpdateLogStoreRequest struct {
	AppendMeta     *bool        `json:"appendMeta,omitempty" xml:"appendMeta,omitempty"`
	AutoSplit      *bool        `json:"autoSplit,omitempty" xml:"autoSplit,omitempty"`
	EnableTracking *bool        `json:"enable_tracking,omitempty" xml:"enable_tracking,omitempty"`
	EncryptConf    *EncryptConf `json:"encrypt_conf,omitempty" xml:"encrypt_conf,omitempty"`
	HotTtl         *int32       `json:"hot_ttl,omitempty" xml:"hot_ttl,omitempty"`
	LogstoreName   *string      `json:"logstoreName,omitempty" xml:"logstoreName,omitempty"`
	MaxSplitShard  *int32       `json:"maxSplitShard,omitempty" xml:"maxSplitShard,omitempty"`
	Mode           *string      `json:"mode,omitempty" xml:"mode,omitempty"`
	ShardCount     *int32       `json:"shardCount,omitempty" xml:"shardCount,omitempty"`
	TelemetryType  *string      `json:"telemetryType,omitempty" xml:"telemetryType,omitempty"`
	Ttl            *int32       `json:"ttl,omitempty" xml:"ttl,omitempty"`
}

func (s UpdateLogStoreRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateLogStoreRequest) GoString() string {
	return s.String()
}

func (s *UpdateLogStoreRequest) SetAppendMeta(v bool) *UpdateLogStoreRequest {
	s.AppendMeta = &v
	return s
}

func (s *UpdateLogStoreRequest) SetAutoSplit(v bool) *UpdateLogStoreRequest {
	s.AutoSplit = &v
	return s
}

func (s *UpdateLogStoreRequest) SetEnableTracking(v bool) *UpdateLogStoreRequest {
	s.EnableTracking = &v
	return s
}

func (s *UpdateLogStoreRequest) SetEncryptConf(v *EncryptConf) *UpdateLogStoreRequest {
	s.EncryptConf = v
	return s
}

func (s *UpdateLogStoreRequest) SetHotTtl(v int32) *UpdateLogStoreRequest {
	s.HotTtl = &v
	return s
}

func (s *UpdateLogStoreRequest) SetLogstoreName(v string) *UpdateLogStoreRequest {
	s.LogstoreName = &v
	return s
}

func (s *UpdateLogStoreRequest) SetMaxSplitShard(v int32) *UpdateLogStoreRequest {
	s.MaxSplitShard = &v
	return s
}

func (s *UpdateLogStoreRequest) SetMode(v string) *UpdateLogStoreRequest {
	s.Mode = &v
	return s
}

func (s *UpdateLogStoreRequest) SetShardCount(v int32) *UpdateLogStoreRequest {
	s.ShardCount = &v
	return s
}

func (s *UpdateLogStoreRequest) SetTelemetryType(v string) *UpdateLogStoreRequest {
	s.TelemetryType = &v
	return s
}

func (s *UpdateLogStoreRequest) SetTtl(v int32) *UpdateLogStoreRequest {
	s.Ttl = &v
	return s
}

type UpdateLogStoreResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
}

func (s UpdateLogStoreResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateLogStoreResponse) GoString() string {
	return s.String()
}

func (s *UpdateLogStoreResponse) SetHeaders(v map[string]*string) *UpdateLogStoreResponse {
	s.Headers = v
	return s
}

func (s *UpdateLogStoreResponse) SetStatusCode(v int32) *UpdateLogStoreResponse {
	s.StatusCode = &v
	return s
}

type UpdateLoggingRequest struct {
	LoggingDetails []*UpdateLoggingRequestLoggingDetails `json:"loggingDetails,omitempty" xml:"loggingDetails,omitempty" type:"Repeated"`
	LoggingProject *string                               `json:"loggingProject,omitempty" xml:"loggingProject,omitempty"`
}

func (s UpdateLoggingRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateLoggingRequest) GoString() string {
	return s.String()
}

func (s *UpdateLoggingRequest) SetLoggingDetails(v []*UpdateLoggingRequestLoggingDetails) *UpdateLoggingRequest {
	s.LoggingDetails = v
	return s
}

func (s *UpdateLoggingRequest) SetLoggingProject(v string) *UpdateLoggingRequest {
	s.LoggingProject = &v
	return s
}

type UpdateLoggingRequestLoggingDetails struct {
	Logstore *string `json:"logstore,omitempty" xml:"logstore,omitempty"`
	Type     *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s UpdateLoggingRequestLoggingDetails) String() string {
	return tea.Prettify(s)
}

func (s UpdateLoggingRequestLoggingDetails) GoString() string {
	return s.String()
}

func (s *UpdateLoggingRequestLoggingDetails) SetLogstore(v string) *UpdateLoggingRequestLoggingDetails {
	s.Logstore = &v
	return s
}

func (s *UpdateLoggingRequestLoggingDetails) SetType(v string) *UpdateLoggingRequestLoggingDetails {
	s.Type = &v
	return s
}

type UpdateLoggingResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
}

func (s UpdateLoggingResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateLoggingResponse) GoString() string {
	return s.String()
}

func (s *UpdateLoggingResponse) SetHeaders(v map[string]*string) *UpdateLoggingResponse {
	s.Headers = v
	return s
}

func (s *UpdateLoggingResponse) SetStatusCode(v int32) *UpdateLoggingResponse {
	s.StatusCode = &v
	return s
}

type UpdateMachineGroupRequest struct {
	GroupAttribute      *UpdateMachineGroupRequestGroupAttribute `json:"groupAttribute,omitempty" xml:"groupAttribute,omitempty" type:"Struct"`
	GroupName           *string                                  `json:"groupName,omitempty" xml:"groupName,omitempty"`
	GroupType           *string                                  `json:"groupType,omitempty" xml:"groupType,omitempty"`
	MachineIdentifyType *string                                  `json:"machineIdentifyType,omitempty" xml:"machineIdentifyType,omitempty"`
	MachineList         []*string                                `json:"machineList,omitempty" xml:"machineList,omitempty" type:"Repeated"`
}

func (s UpdateMachineGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateMachineGroupRequest) GoString() string {
	return s.String()
}

func (s *UpdateMachineGroupRequest) SetGroupAttribute(v *UpdateMachineGroupRequestGroupAttribute) *UpdateMachineGroupRequest {
	s.GroupAttribute = v
	return s
}

func (s *UpdateMachineGroupRequest) SetGroupName(v string) *UpdateMachineGroupRequest {
	s.GroupName = &v
	return s
}

func (s *UpdateMachineGroupRequest) SetGroupType(v string) *UpdateMachineGroupRequest {
	s.GroupType = &v
	return s
}

func (s *UpdateMachineGroupRequest) SetMachineIdentifyType(v string) *UpdateMachineGroupRequest {
	s.MachineIdentifyType = &v
	return s
}

func (s *UpdateMachineGroupRequest) SetMachineList(v []*string) *UpdateMachineGroupRequest {
	s.MachineList = v
	return s
}

type UpdateMachineGroupRequestGroupAttribute struct {
	ExternalName *string `json:"externalName,omitempty" xml:"externalName,omitempty"`
	GroupTopic   *string `json:"groupTopic,omitempty" xml:"groupTopic,omitempty"`
}

func (s UpdateMachineGroupRequestGroupAttribute) String() string {
	return tea.Prettify(s)
}

func (s UpdateMachineGroupRequestGroupAttribute) GoString() string {
	return s.String()
}

func (s *UpdateMachineGroupRequestGroupAttribute) SetExternalName(v string) *UpdateMachineGroupRequestGroupAttribute {
	s.ExternalName = &v
	return s
}

func (s *UpdateMachineGroupRequestGroupAttribute) SetGroupTopic(v string) *UpdateMachineGroupRequestGroupAttribute {
	s.GroupTopic = &v
	return s
}

type UpdateMachineGroupResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
}

func (s UpdateMachineGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateMachineGroupResponse) GoString() string {
	return s.String()
}

func (s *UpdateMachineGroupResponse) SetHeaders(v map[string]*string) *UpdateMachineGroupResponse {
	s.Headers = v
	return s
}

func (s *UpdateMachineGroupResponse) SetStatusCode(v int32) *UpdateMachineGroupResponse {
	s.StatusCode = &v
	return s
}

type UpdateMachineGroupMachineRequest struct {
	Action *string   `json:"action,omitempty" xml:"action,omitempty"`
	Body   []*string `json:"body,omitempty" xml:"body,omitempty" type:"Repeated"`
}

func (s UpdateMachineGroupMachineRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateMachineGroupMachineRequest) GoString() string {
	return s.String()
}

func (s *UpdateMachineGroupMachineRequest) SetAction(v string) *UpdateMachineGroupMachineRequest {
	s.Action = &v
	return s
}

func (s *UpdateMachineGroupMachineRequest) SetBody(v []*string) *UpdateMachineGroupMachineRequest {
	s.Body = v
	return s
}

type UpdateMachineGroupMachineResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
}

func (s UpdateMachineGroupMachineResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateMachineGroupMachineResponse) GoString() string {
	return s.String()
}

func (s *UpdateMachineGroupMachineResponse) SetHeaders(v map[string]*string) *UpdateMachineGroupMachineResponse {
	s.Headers = v
	return s
}

func (s *UpdateMachineGroupMachineResponse) SetStatusCode(v int32) *UpdateMachineGroupMachineResponse {
	s.StatusCode = &v
	return s
}

type UpdateOdpsShipperRequest struct {
	ShipperName         *string                                      `json:"shipperName,omitempty" xml:"shipperName,omitempty"`
	TargetConfiguration *UpdateOdpsShipperRequestTargetConfiguration `json:"targetConfiguration,omitempty" xml:"targetConfiguration,omitempty" type:"Struct"`
	TargetType          *string                                      `json:"targetType,omitempty" xml:"targetType,omitempty"`
}

func (s UpdateOdpsShipperRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateOdpsShipperRequest) GoString() string {
	return s.String()
}

func (s *UpdateOdpsShipperRequest) SetShipperName(v string) *UpdateOdpsShipperRequest {
	s.ShipperName = &v
	return s
}

func (s *UpdateOdpsShipperRequest) SetTargetConfiguration(v *UpdateOdpsShipperRequestTargetConfiguration) *UpdateOdpsShipperRequest {
	s.TargetConfiguration = v
	return s
}

func (s *UpdateOdpsShipperRequest) SetTargetType(v string) *UpdateOdpsShipperRequest {
	s.TargetType = &v
	return s
}

type UpdateOdpsShipperRequestTargetConfiguration struct {
	BufferInterval      *int32    `json:"bufferInterval,omitempty" xml:"bufferInterval,omitempty"`
	Enable              *bool     `json:"enable,omitempty" xml:"enable,omitempty"`
	Fields              []*string `json:"fields,omitempty" xml:"fields,omitempty" type:"Repeated"`
	OdpsEndpoint        *string   `json:"odpsEndpoint,omitempty" xml:"odpsEndpoint,omitempty"`
	OdpsProject         *string   `json:"odpsProject,omitempty" xml:"odpsProject,omitempty"`
	OdpsTable           *string   `json:"odpsTable,omitempty" xml:"odpsTable,omitempty"`
	PartitionColumn     []*string `json:"partitionColumn,omitempty" xml:"partitionColumn,omitempty" type:"Repeated"`
	PartitionTimeFormat *string   `json:"partitionTimeFormat,omitempty" xml:"partitionTimeFormat,omitempty"`
}

func (s UpdateOdpsShipperRequestTargetConfiguration) String() string {
	return tea.Prettify(s)
}

func (s UpdateOdpsShipperRequestTargetConfiguration) GoString() string {
	return s.String()
}

func (s *UpdateOdpsShipperRequestTargetConfiguration) SetBufferInterval(v int32) *UpdateOdpsShipperRequestTargetConfiguration {
	s.BufferInterval = &v
	return s
}

func (s *UpdateOdpsShipperRequestTargetConfiguration) SetEnable(v bool) *UpdateOdpsShipperRequestTargetConfiguration {
	s.Enable = &v
	return s
}

func (s *UpdateOdpsShipperRequestTargetConfiguration) SetFields(v []*string) *UpdateOdpsShipperRequestTargetConfiguration {
	s.Fields = v
	return s
}

func (s *UpdateOdpsShipperRequestTargetConfiguration) SetOdpsEndpoint(v string) *UpdateOdpsShipperRequestTargetConfiguration {
	s.OdpsEndpoint = &v
	return s
}

func (s *UpdateOdpsShipperRequestTargetConfiguration) SetOdpsProject(v string) *UpdateOdpsShipperRequestTargetConfiguration {
	s.OdpsProject = &v
	return s
}

func (s *UpdateOdpsShipperRequestTargetConfiguration) SetOdpsTable(v string) *UpdateOdpsShipperRequestTargetConfiguration {
	s.OdpsTable = &v
	return s
}

func (s *UpdateOdpsShipperRequestTargetConfiguration) SetPartitionColumn(v []*string) *UpdateOdpsShipperRequestTargetConfiguration {
	s.PartitionColumn = v
	return s
}

func (s *UpdateOdpsShipperRequestTargetConfiguration) SetPartitionTimeFormat(v string) *UpdateOdpsShipperRequestTargetConfiguration {
	s.PartitionTimeFormat = &v
	return s
}

type UpdateOdpsShipperResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
}

func (s UpdateOdpsShipperResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateOdpsShipperResponse) GoString() string {
	return s.String()
}

func (s *UpdateOdpsShipperResponse) SetHeaders(v map[string]*string) *UpdateOdpsShipperResponse {
	s.Headers = v
	return s
}

func (s *UpdateOdpsShipperResponse) SetStatusCode(v int32) *UpdateOdpsShipperResponse {
	s.StatusCode = &v
	return s
}

type UpdateOssExternalStoreRequest struct {
	ExternalStoreName *string                                 `json:"externalStoreName,omitempty" xml:"externalStoreName,omitempty"`
	Parameter         *UpdateOssExternalStoreRequestParameter `json:"parameter,omitempty" xml:"parameter,omitempty" type:"Struct"`
	StoreType         *string                                 `json:"storeType,omitempty" xml:"storeType,omitempty"`
}

func (s UpdateOssExternalStoreRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateOssExternalStoreRequest) GoString() string {
	return s.String()
}

func (s *UpdateOssExternalStoreRequest) SetExternalStoreName(v string) *UpdateOssExternalStoreRequest {
	s.ExternalStoreName = &v
	return s
}

func (s *UpdateOssExternalStoreRequest) SetParameter(v *UpdateOssExternalStoreRequestParameter) *UpdateOssExternalStoreRequest {
	s.Parameter = v
	return s
}

func (s *UpdateOssExternalStoreRequest) SetStoreType(v string) *UpdateOssExternalStoreRequest {
	s.StoreType = &v
	return s
}

type UpdateOssExternalStoreRequestParameter struct {
	Accessid  *string                                          `json:"accessid,omitempty" xml:"accessid,omitempty"`
	Accesskey *string                                          `json:"accesskey,omitempty" xml:"accesskey,omitempty"`
	Bucket    *string                                          `json:"bucket,omitempty" xml:"bucket,omitempty"`
	Columns   []*UpdateOssExternalStoreRequestParameterColumns `json:"columns,omitempty" xml:"columns,omitempty" type:"Repeated"`
	Endpoint  *string                                          `json:"endpoint,omitempty" xml:"endpoint,omitempty"`
	Objects   []*string                                        `json:"objects,omitempty" xml:"objects,omitempty" type:"Repeated"`
}

func (s UpdateOssExternalStoreRequestParameter) String() string {
	return tea.Prettify(s)
}

func (s UpdateOssExternalStoreRequestParameter) GoString() string {
	return s.String()
}

func (s *UpdateOssExternalStoreRequestParameter) SetAccessid(v string) *UpdateOssExternalStoreRequestParameter {
	s.Accessid = &v
	return s
}

func (s *UpdateOssExternalStoreRequestParameter) SetAccesskey(v string) *UpdateOssExternalStoreRequestParameter {
	s.Accesskey = &v
	return s
}

func (s *UpdateOssExternalStoreRequestParameter) SetBucket(v string) *UpdateOssExternalStoreRequestParameter {
	s.Bucket = &v
	return s
}

func (s *UpdateOssExternalStoreRequestParameter) SetColumns(v []*UpdateOssExternalStoreRequestParameterColumns) *UpdateOssExternalStoreRequestParameter {
	s.Columns = v
	return s
}

func (s *UpdateOssExternalStoreRequestParameter) SetEndpoint(v string) *UpdateOssExternalStoreRequestParameter {
	s.Endpoint = &v
	return s
}

func (s *UpdateOssExternalStoreRequestParameter) SetObjects(v []*string) *UpdateOssExternalStoreRequestParameter {
	s.Objects = v
	return s
}

type UpdateOssExternalStoreRequestParameterColumns struct {
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s UpdateOssExternalStoreRequestParameterColumns) String() string {
	return tea.Prettify(s)
}

func (s UpdateOssExternalStoreRequestParameterColumns) GoString() string {
	return s.String()
}

func (s *UpdateOssExternalStoreRequestParameterColumns) SetName(v string) *UpdateOssExternalStoreRequestParameterColumns {
	s.Name = &v
	return s
}

func (s *UpdateOssExternalStoreRequestParameterColumns) SetType(v string) *UpdateOssExternalStoreRequestParameterColumns {
	s.Type = &v
	return s
}

type UpdateOssExternalStoreResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
}

func (s UpdateOssExternalStoreResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateOssExternalStoreResponse) GoString() string {
	return s.String()
}

func (s *UpdateOssExternalStoreResponse) SetHeaders(v map[string]*string) *UpdateOssExternalStoreResponse {
	s.Headers = v
	return s
}

func (s *UpdateOssExternalStoreResponse) SetStatusCode(v int32) *UpdateOssExternalStoreResponse {
	s.StatusCode = &v
	return s
}

type UpdateOssShipperRequest struct {
	ShipperName         *string                                     `json:"shipperName,omitempty" xml:"shipperName,omitempty"`
	TargetConfiguration *UpdateOssShipperRequestTargetConfiguration `json:"targetConfiguration,omitempty" xml:"targetConfiguration,omitempty" type:"Struct"`
	TargetType          *string                                     `json:"targetType,omitempty" xml:"targetType,omitempty"`
}

func (s UpdateOssShipperRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateOssShipperRequest) GoString() string {
	return s.String()
}

func (s *UpdateOssShipperRequest) SetShipperName(v string) *UpdateOssShipperRequest {
	s.ShipperName = &v
	return s
}

func (s *UpdateOssShipperRequest) SetTargetConfiguration(v *UpdateOssShipperRequestTargetConfiguration) *UpdateOssShipperRequest {
	s.TargetConfiguration = v
	return s
}

func (s *UpdateOssShipperRequest) SetTargetType(v string) *UpdateOssShipperRequest {
	s.TargetType = &v
	return s
}

type UpdateOssShipperRequestTargetConfiguration struct {
	BufferInterval *int32                                             `json:"bufferInterval,omitempty" xml:"bufferInterval,omitempty"`
	BufferSize     *int32                                             `json:"bufferSize,omitempty" xml:"bufferSize,omitempty"`
	CompressType   *string                                            `json:"compressType,omitempty" xml:"compressType,omitempty"`
	Enable         *bool                                              `json:"enable,omitempty" xml:"enable,omitempty"`
	OssBucket      *string                                            `json:"ossBucket,omitempty" xml:"ossBucket,omitempty"`
	OssPrefix      *string                                            `json:"ossPrefix,omitempty" xml:"ossPrefix,omitempty"`
	PathFormat     *string                                            `json:"pathFormat,omitempty" xml:"pathFormat,omitempty"`
	RoleArn        *string                                            `json:"roleArn,omitempty" xml:"roleArn,omitempty"`
	Storage        *UpdateOssShipperRequestTargetConfigurationStorage `json:"storage,omitempty" xml:"storage,omitempty" type:"Struct"`
	TimeZone       *string                                            `json:"timeZone,omitempty" xml:"timeZone,omitempty"`
}

func (s UpdateOssShipperRequestTargetConfiguration) String() string {
	return tea.Prettify(s)
}

func (s UpdateOssShipperRequestTargetConfiguration) GoString() string {
	return s.String()
}

func (s *UpdateOssShipperRequestTargetConfiguration) SetBufferInterval(v int32) *UpdateOssShipperRequestTargetConfiguration {
	s.BufferInterval = &v
	return s
}

func (s *UpdateOssShipperRequestTargetConfiguration) SetBufferSize(v int32) *UpdateOssShipperRequestTargetConfiguration {
	s.BufferSize = &v
	return s
}

func (s *UpdateOssShipperRequestTargetConfiguration) SetCompressType(v string) *UpdateOssShipperRequestTargetConfiguration {
	s.CompressType = &v
	return s
}

func (s *UpdateOssShipperRequestTargetConfiguration) SetEnable(v bool) *UpdateOssShipperRequestTargetConfiguration {
	s.Enable = &v
	return s
}

func (s *UpdateOssShipperRequestTargetConfiguration) SetOssBucket(v string) *UpdateOssShipperRequestTargetConfiguration {
	s.OssBucket = &v
	return s
}

func (s *UpdateOssShipperRequestTargetConfiguration) SetOssPrefix(v string) *UpdateOssShipperRequestTargetConfiguration {
	s.OssPrefix = &v
	return s
}

func (s *UpdateOssShipperRequestTargetConfiguration) SetPathFormat(v string) *UpdateOssShipperRequestTargetConfiguration {
	s.PathFormat = &v
	return s
}

func (s *UpdateOssShipperRequestTargetConfiguration) SetRoleArn(v string) *UpdateOssShipperRequestTargetConfiguration {
	s.RoleArn = &v
	return s
}

func (s *UpdateOssShipperRequestTargetConfiguration) SetStorage(v *UpdateOssShipperRequestTargetConfigurationStorage) *UpdateOssShipperRequestTargetConfiguration {
	s.Storage = v
	return s
}

func (s *UpdateOssShipperRequestTargetConfiguration) SetTimeZone(v string) *UpdateOssShipperRequestTargetConfiguration {
	s.TimeZone = &v
	return s
}

type UpdateOssShipperRequestTargetConfigurationStorage struct {
	Detail map[string]interface{} `json:"detail,omitempty" xml:"detail,omitempty"`
	Format *string                `json:"format,omitempty" xml:"format,omitempty"`
}

func (s UpdateOssShipperRequestTargetConfigurationStorage) String() string {
	return tea.Prettify(s)
}

func (s UpdateOssShipperRequestTargetConfigurationStorage) GoString() string {
	return s.String()
}

func (s *UpdateOssShipperRequestTargetConfigurationStorage) SetDetail(v map[string]interface{}) *UpdateOssShipperRequestTargetConfigurationStorage {
	s.Detail = v
	return s
}

func (s *UpdateOssShipperRequestTargetConfigurationStorage) SetFormat(v string) *UpdateOssShipperRequestTargetConfigurationStorage {
	s.Format = &v
	return s
}

type UpdateOssShipperResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
}

func (s UpdateOssShipperResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateOssShipperResponse) GoString() string {
	return s.String()
}

func (s *UpdateOssShipperResponse) SetHeaders(v map[string]*string) *UpdateOssShipperResponse {
	s.Headers = v
	return s
}

func (s *UpdateOssShipperResponse) SetStatusCode(v int32) *UpdateOssShipperResponse {
	s.StatusCode = &v
	return s
}

type UpdateProjectRequest struct {
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
}

func (s UpdateProjectRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateProjectRequest) GoString() string {
	return s.String()
}

func (s *UpdateProjectRequest) SetDescription(v string) *UpdateProjectRequest {
	s.Description = &v
	return s
}

type UpdateProjectResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
}

func (s UpdateProjectResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateProjectResponse) GoString() string {
	return s.String()
}

func (s *UpdateProjectResponse) SetHeaders(v map[string]*string) *UpdateProjectResponse {
	s.Headers = v
	return s
}

func (s *UpdateProjectResponse) SetStatusCode(v int32) *UpdateProjectResponse {
	s.StatusCode = &v
	return s
}

type UpdateRdsExternalStoreRequest struct {
	ExternalStoreName *string                                 `json:"externalStoreName,omitempty" xml:"externalStoreName,omitempty"`
	Parameter         *UpdateRdsExternalStoreRequestParameter `json:"parameter,omitempty" xml:"parameter,omitempty" type:"Struct"`
	StoreType         *string                                 `json:"storeType,omitempty" xml:"storeType,omitempty"`
}

func (s UpdateRdsExternalStoreRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateRdsExternalStoreRequest) GoString() string {
	return s.String()
}

func (s *UpdateRdsExternalStoreRequest) SetExternalStoreName(v string) *UpdateRdsExternalStoreRequest {
	s.ExternalStoreName = &v
	return s
}

func (s *UpdateRdsExternalStoreRequest) SetParameter(v *UpdateRdsExternalStoreRequestParameter) *UpdateRdsExternalStoreRequest {
	s.Parameter = v
	return s
}

func (s *UpdateRdsExternalStoreRequest) SetStoreType(v string) *UpdateRdsExternalStoreRequest {
	s.StoreType = &v
	return s
}

type UpdateRdsExternalStoreRequestParameter struct {
	Db         *string `json:"db,omitempty" xml:"db,omitempty"`
	Host       *string `json:"host,omitempty" xml:"host,omitempty"`
	InstanceId *string `json:"instance-id,omitempty" xml:"instance-id,omitempty"`
	Password   *string `json:"password,omitempty" xml:"password,omitempty"`
	Port       *string `json:"port,omitempty" xml:"port,omitempty"`
	Region     *string `json:"region,omitempty" xml:"region,omitempty"`
	Table      *string `json:"table,omitempty" xml:"table,omitempty"`
	Username   *string `json:"username,omitempty" xml:"username,omitempty"`
	VpcId      *string `json:"vpc-id,omitempty" xml:"vpc-id,omitempty"`
}

func (s UpdateRdsExternalStoreRequestParameter) String() string {
	return tea.Prettify(s)
}

func (s UpdateRdsExternalStoreRequestParameter) GoString() string {
	return s.String()
}

func (s *UpdateRdsExternalStoreRequestParameter) SetDb(v string) *UpdateRdsExternalStoreRequestParameter {
	s.Db = &v
	return s
}

func (s *UpdateRdsExternalStoreRequestParameter) SetHost(v string) *UpdateRdsExternalStoreRequestParameter {
	s.Host = &v
	return s
}

func (s *UpdateRdsExternalStoreRequestParameter) SetInstanceId(v string) *UpdateRdsExternalStoreRequestParameter {
	s.InstanceId = &v
	return s
}

func (s *UpdateRdsExternalStoreRequestParameter) SetPassword(v string) *UpdateRdsExternalStoreRequestParameter {
	s.Password = &v
	return s
}

func (s *UpdateRdsExternalStoreRequestParameter) SetPort(v string) *UpdateRdsExternalStoreRequestParameter {
	s.Port = &v
	return s
}

func (s *UpdateRdsExternalStoreRequestParameter) SetRegion(v string) *UpdateRdsExternalStoreRequestParameter {
	s.Region = &v
	return s
}

func (s *UpdateRdsExternalStoreRequestParameter) SetTable(v string) *UpdateRdsExternalStoreRequestParameter {
	s.Table = &v
	return s
}

func (s *UpdateRdsExternalStoreRequestParameter) SetUsername(v string) *UpdateRdsExternalStoreRequestParameter {
	s.Username = &v
	return s
}

func (s *UpdateRdsExternalStoreRequestParameter) SetVpcId(v string) *UpdateRdsExternalStoreRequestParameter {
	s.VpcId = &v
	return s
}

type UpdateRdsExternalStoreResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
}

func (s UpdateRdsExternalStoreResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateRdsExternalStoreResponse) GoString() string {
	return s.String()
}

func (s *UpdateRdsExternalStoreResponse) SetHeaders(v map[string]*string) *UpdateRdsExternalStoreResponse {
	s.Headers = v
	return s
}

func (s *UpdateRdsExternalStoreResponse) SetStatusCode(v int32) *UpdateRdsExternalStoreResponse {
	s.StatusCode = &v
	return s
}

type UpdateSavedSearchRequest struct {
	DisplayName     *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	Logstore        *string `json:"logstore,omitempty" xml:"logstore,omitempty"`
	SavedsearchName *string `json:"savedsearchName,omitempty" xml:"savedsearchName,omitempty"`
	SearchQuery     *string `json:"searchQuery,omitempty" xml:"searchQuery,omitempty"`
	Topic           *string `json:"topic,omitempty" xml:"topic,omitempty"`
}

func (s UpdateSavedSearchRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateSavedSearchRequest) GoString() string {
	return s.String()
}

func (s *UpdateSavedSearchRequest) SetDisplayName(v string) *UpdateSavedSearchRequest {
	s.DisplayName = &v
	return s
}

func (s *UpdateSavedSearchRequest) SetLogstore(v string) *UpdateSavedSearchRequest {
	s.Logstore = &v
	return s
}

func (s *UpdateSavedSearchRequest) SetSavedsearchName(v string) *UpdateSavedSearchRequest {
	s.SavedsearchName = &v
	return s
}

func (s *UpdateSavedSearchRequest) SetSearchQuery(v string) *UpdateSavedSearchRequest {
	s.SearchQuery = &v
	return s
}

func (s *UpdateSavedSearchRequest) SetTopic(v string) *UpdateSavedSearchRequest {
	s.Topic = &v
	return s
}

type UpdateSavedSearchResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
}

func (s UpdateSavedSearchResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateSavedSearchResponse) GoString() string {
	return s.String()
}

func (s *UpdateSavedSearchResponse) SetHeaders(v map[string]*string) *UpdateSavedSearchResponse {
	s.Headers = v
	return s
}

func (s *UpdateSavedSearchResponse) SetStatusCode(v int32) *UpdateSavedSearchResponse {
	s.StatusCode = &v
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
	interfaceSPI, _err := gatewayclient.NewClient()
	if _err != nil {
		return _err
	}

	client.Spi = interfaceSPI
	client.EndpointRule = tea.String("central")
	return nil
}

func (client *Client) ApplyConfigToMachineGroupWithOptions(project *string, machineGroup *string, configName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ApplyConfigToMachineGroupResponse, _err error) {
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("ApplyConfigToMachineGroup"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/machinegroups/" + tea.StringValue(machineGroup) + "/configs/" + tea.StringValue(configName)),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &ApplyConfigToMachineGroupResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ApplyConfigToMachineGroup(project *string, machineGroup *string, configName *string) (_result *ApplyConfigToMachineGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ApplyConfigToMachineGroupResponse{}
	_body, _err := client.ApplyConfigToMachineGroupWithOptions(project, machineGroup, configName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ChangeResourceGroupWithOptions(request *ChangeResourceGroupRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ChangeResourceGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		body["resourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		body["resourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		body["resourceType"] = request.ResourceType
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ChangeResourceGroup"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/resourcegroup"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &ChangeResourceGroupResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ChangeResourceGroup(request *ChangeResourceGroupRequest) (_result *ChangeResourceGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ChangeResourceGroupResponse{}
	_body, _err := client.ChangeResourceGroupWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateAlertWithOptions(project *string, request *CreateAlertRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateAlertResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateAlert"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/jobs"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &CreateAlertResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateAlert(project *string, request *CreateAlertRequest) (_result *CreateAlertResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateAlertResponse{}
	_body, _err := client.CreateAlertWithOptions(project, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateConsumerGroupWithOptions(project *string, logstore *string, request *CreateConsumerGroupRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateConsumerGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConsumerGroup)) {
		body["consumerGroup"] = request.ConsumerGroup
	}

	if !tea.BoolValue(util.IsUnset(request.Order)) {
		body["order"] = request.Order
	}

	if !tea.BoolValue(util.IsUnset(request.Timeout)) {
		body["timeout"] = request.Timeout
	}

	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateConsumerGroup"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/logstores/" + tea.StringValue(logstore) + "/consumergroups"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &CreateConsumerGroupResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateConsumerGroup(project *string, logstore *string, request *CreateConsumerGroupRequest) (_result *CreateConsumerGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateConsumerGroupResponse{}
	_body, _err := client.CreateConsumerGroupWithOptions(project, logstore, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateDomainWithOptions(project *string, request *CreateDomainRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateDomainResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		body["domainName"] = request.DomainName
	}

	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateDomain"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/domains"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &CreateDomainResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateDomain(project *string, request *CreateDomainRequest) (_result *CreateDomainResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateDomainResponse{}
	_body, _err := client.CreateDomainWithOptions(project, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateETLJobWithOptions(project *string, request *CreateETLJobRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateETLJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Enable)) {
		body["enable"] = request.Enable
	}

	if !tea.BoolValue(util.IsUnset(request.FunctionConfig)) {
		body["functionConfig"] = request.FunctionConfig
	}

	if !tea.BoolValue(util.IsUnset(request.FunctionParameter)) {
		body["functionParameter"] = request.FunctionParameter
	}

	if !tea.BoolValue(util.IsUnset(request.JobName)) {
		body["jobName"] = request.JobName
	}

	if !tea.BoolValue(util.IsUnset(request.LogConfig)) {
		body["logConfig"] = request.LogConfig
	}

	if !tea.BoolValue(util.IsUnset(request.SourceConfig)) {
		body["sourceConfig"] = request.SourceConfig
	}

	if !tea.BoolValue(util.IsUnset(request.TriggerConfig)) {
		body["triggerConfig"] = request.TriggerConfig
	}

	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateETLJob"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/jobs"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateETLJobResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateETLJob(project *string, request *CreateETLJobRequest) (_result *CreateETLJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateETLJobResponse{}
	_body, _err := client.CreateETLJobWithOptions(project, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateIndexWithOptions(project *string, logstore *string, request *CreateIndexRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateIndexResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Keys)) {
		body["keys"] = request.Keys
	}

	if !tea.BoolValue(util.IsUnset(request.Line)) {
		body["line"] = request.Line
	}

	if !tea.BoolValue(util.IsUnset(request.LogReduce)) {
		body["log_reduce"] = request.LogReduce
	}

	if !tea.BoolValue(util.IsUnset(request.LogReduceBlackList)) {
		body["log_reduce_black_list"] = request.LogReduceBlackList
	}

	if !tea.BoolValue(util.IsUnset(request.LogReduceWhiteList)) {
		body["log_reduce_white_list"] = request.LogReduceWhiteList
	}

	if !tea.BoolValue(util.IsUnset(request.MaxTextLen)) {
		body["max_text_len"] = request.MaxTextLen
	}

	if !tea.BoolValue(util.IsUnset(request.Ttl)) {
		body["ttl"] = request.Ttl
	}

	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateIndex"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/logstores/" + tea.StringValue(logstore) + "/index"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &CreateIndexResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateIndex(project *string, logstore *string, request *CreateIndexRequest) (_result *CreateIndexResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateIndexResponse{}
	_body, _err := client.CreateIndexWithOptions(project, logstore, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateJobWithOptions(project *string, request *CreateJobRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Enable)) {
		body["enable"] = request.Enable
	}

	if !tea.BoolValue(util.IsUnset(request.FunctionConfig)) {
		body["functionConfig"] = request.FunctionConfig
	}

	if !tea.BoolValue(util.IsUnset(request.FunctionParameter)) {
		body["functionParameter"] = request.FunctionParameter
	}

	if !tea.BoolValue(util.IsUnset(request.JobName)) {
		body["jobName"] = request.JobName
	}

	if !tea.BoolValue(util.IsUnset(request.LogConfig)) {
		body["logConfig"] = request.LogConfig
	}

	if !tea.BoolValue(util.IsUnset(request.SourceConfig)) {
		body["sourceConfig"] = request.SourceConfig
	}

	if !tea.BoolValue(util.IsUnset(request.TriggerConfig)) {
		body["triggerConfig"] = request.TriggerConfig
	}

	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateJob"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/jobs"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateJobResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateJob(project *string, request *CreateJobRequest) (_result *CreateJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateJobResponse{}
	_body, _err := client.CreateJobWithOptions(project, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateLogStoreWithOptions(project *string, request *CreateLogStoreRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateLogStoreResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppendMeta)) {
		body["appendMeta"] = request.AppendMeta
	}

	if !tea.BoolValue(util.IsUnset(request.AutoSplit)) {
		body["autoSplit"] = request.AutoSplit
	}

	if !tea.BoolValue(util.IsUnset(request.EnableTracking)) {
		body["enable_tracking"] = request.EnableTracking
	}

	if !tea.BoolValue(util.IsUnset(request.EncryptConf)) {
		body["encrypt_conf"] = request.EncryptConf
	}

	if !tea.BoolValue(util.IsUnset(request.HotTtl)) {
		body["hot_ttl"] = request.HotTtl
	}

	if !tea.BoolValue(util.IsUnset(request.LogstoreName)) {
		body["logstoreName"] = request.LogstoreName
	}

	if !tea.BoolValue(util.IsUnset(request.MaxSplitShard)) {
		body["maxSplitShard"] = request.MaxSplitShard
	}

	if !tea.BoolValue(util.IsUnset(request.Mode)) {
		body["mode"] = request.Mode
	}

	if !tea.BoolValue(util.IsUnset(request.ShardCount)) {
		body["shardCount"] = request.ShardCount
	}

	if !tea.BoolValue(util.IsUnset(request.TelemetryType)) {
		body["telemetryType"] = request.TelemetryType
	}

	if !tea.BoolValue(util.IsUnset(request.Ttl)) {
		body["ttl"] = request.Ttl
	}

	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateLogStore"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/logstores"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &CreateLogStoreResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateLogStore(project *string, request *CreateLogStoreRequest) (_result *CreateLogStoreResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateLogStoreResponse{}
	_body, _err := client.CreateLogStoreWithOptions(project, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateLoggingWithOptions(project *string, request *CreateLoggingRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateLoggingResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.LoggingDetails)) {
		body["loggingDetails"] = request.LoggingDetails
	}

	if !tea.BoolValue(util.IsUnset(request.LoggingProject)) {
		body["loggingProject"] = request.LoggingProject
	}

	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateLogging"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/logging"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &CreateLoggingResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateLogging(project *string, request *CreateLoggingRequest) (_result *CreateLoggingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateLoggingResponse{}
	_body, _err := client.CreateLoggingWithOptions(project, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateMachineGroupWithOptions(project *string, request *CreateMachineGroupRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateMachineGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GroupAttribute)) {
		body["groupAttribute"] = request.GroupAttribute
	}

	if !tea.BoolValue(util.IsUnset(request.GroupName)) {
		body["groupName"] = request.GroupName
	}

	if !tea.BoolValue(util.IsUnset(request.GroupType)) {
		body["groupType"] = request.GroupType
	}

	if !tea.BoolValue(util.IsUnset(request.MachineIdentifyType)) {
		body["machineIdentifyType"] = request.MachineIdentifyType
	}

	if !tea.BoolValue(util.IsUnset(request.MachineList)) {
		body["machineList"] = request.MachineList
	}

	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateMachineGroup"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/machinegroups"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &CreateMachineGroupResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateMachineGroup(project *string, request *CreateMachineGroupRequest) (_result *CreateMachineGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateMachineGroupResponse{}
	_body, _err := client.CreateMachineGroupWithOptions(project, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateOdpsShipperWithOptions(project *string, logstore *string, request *CreateOdpsShipperRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateOdpsShipperResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ShipperName)) {
		body["shipperName"] = request.ShipperName
	}

	if !tea.BoolValue(util.IsUnset(request.TargetConfiguration)) {
		body["targetConfiguration"] = request.TargetConfiguration
	}

	if !tea.BoolValue(util.IsUnset(request.TargetType)) {
		body["targetType"] = request.TargetType
	}

	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateOdpsShipper"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/logstores/" + tea.StringValue(logstore) + "/shipper"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &CreateOdpsShipperResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateOdpsShipper(project *string, logstore *string, request *CreateOdpsShipperRequest) (_result *CreateOdpsShipperResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateOdpsShipperResponse{}
	_body, _err := client.CreateOdpsShipperWithOptions(project, logstore, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateOssExternalStoreWithOptions(project *string, request *CreateOssExternalStoreRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateOssExternalStoreResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ExternalStoreName)) {
		body["externalStoreName"] = request.ExternalStoreName
	}

	if !tea.BoolValue(util.IsUnset(request.Parameter)) {
		body["parameter"] = request.Parameter
	}

	if !tea.BoolValue(util.IsUnset(request.StoreType)) {
		body["storeType"] = request.StoreType
	}

	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateOssExternalStore"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/externalstores"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &CreateOssExternalStoreResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateOssExternalStore(project *string, request *CreateOssExternalStoreRequest) (_result *CreateOssExternalStoreResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateOssExternalStoreResponse{}
	_body, _err := client.CreateOssExternalStoreWithOptions(project, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateOssShipperWithOptions(project *string, logstore *string, request *CreateOssShipperRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateOssShipperResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ShipperName)) {
		body["shipperName"] = request.ShipperName
	}

	if !tea.BoolValue(util.IsUnset(request.TargetConfiguration)) {
		body["targetConfiguration"] = request.TargetConfiguration
	}

	if !tea.BoolValue(util.IsUnset(request.TargetType)) {
		body["targetType"] = request.TargetType
	}

	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateOssShipper"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/logstores/" + tea.StringValue(logstore) + "/shipper"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &CreateOssShipperResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateOssShipper(project *string, logstore *string, request *CreateOssShipperRequest) (_result *CreateOssShipperResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateOssShipperResponse{}
	_body, _err := client.CreateOssShipperWithOptions(project, logstore, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateProjectWithOptions(request *CreateProjectRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateProjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectName)) {
		body["projectName"] = request.ProjectName
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		body["resourceGroupId"] = request.ResourceGroupId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateProject"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &CreateProjectResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateProject(request *CreateProjectRequest) (_result *CreateProjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateProjectResponse{}
	_body, _err := client.CreateProjectWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateRdsExternalStoreWithOptions(project *string, request *CreateRdsExternalStoreRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateRdsExternalStoreResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ExternalStoreName)) {
		body["externalStoreName"] = request.ExternalStoreName
	}

	if !tea.BoolValue(util.IsUnset(request.Parameter)) {
		body["parameter"] = request.Parameter
	}

	if !tea.BoolValue(util.IsUnset(request.StoreType)) {
		body["storeType"] = request.StoreType
	}

	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateRdsExternalStore"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/externalstores"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &CreateRdsExternalStoreResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateRdsExternalStore(project *string, request *CreateRdsExternalStoreRequest) (_result *CreateRdsExternalStoreResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateRdsExternalStoreResponse{}
	_body, _err := client.CreateRdsExternalStoreWithOptions(project, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateSavedSearchWithOptions(project *string, request *CreateSavedSearchRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateSavedSearchResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DisplayName)) {
		body["displayName"] = request.DisplayName
	}

	if !tea.BoolValue(util.IsUnset(request.Logstore)) {
		body["logstore"] = request.Logstore
	}

	if !tea.BoolValue(util.IsUnset(request.SavedsearchName)) {
		body["savedsearchName"] = request.SavedsearchName
	}

	if !tea.BoolValue(util.IsUnset(request.SearchQuery)) {
		body["searchQuery"] = request.SearchQuery
	}

	if !tea.BoolValue(util.IsUnset(request.Topic)) {
		body["topic"] = request.Topic
	}

	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateSavedSearch"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/savedsearches"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &CreateSavedSearchResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateSavedSearch(project *string, request *CreateSavedSearchRequest) (_result *CreateSavedSearchResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateSavedSearchResponse{}
	_body, _err := client.CreateSavedSearchWithOptions(project, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteAlertWithOptions(project *string, alert *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteAlertResponse, _err error) {
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteAlert"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/jobs/" + tea.StringValue(alert)),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &DeleteAlertResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteAlert(project *string, alert *string) (_result *DeleteAlertResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteAlertResponse{}
	_body, _err := client.DeleteAlertWithOptions(project, alert, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteConsumerGroupWithOptions(project *string, logstore *string, consumerGroup *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteConsumerGroupResponse, _err error) {
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteConsumerGroup"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/logstores/" + tea.StringValue(logstore) + "/consumergroups/" + tea.StringValue(consumerGroup)),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &DeleteConsumerGroupResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteConsumerGroup(project *string, logstore *string, consumerGroup *string) (_result *DeleteConsumerGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteConsumerGroupResponse{}
	_body, _err := client.DeleteConsumerGroupWithOptions(project, logstore, consumerGroup, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteDomainWithOptions(project *string, domainName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteDomainResponse, _err error) {
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteDomain"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/domains/" + tea.StringValue(domainName)),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &DeleteDomainResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteDomain(project *string, domainName *string) (_result *DeleteDomainResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteDomainResponse{}
	_body, _err := client.DeleteDomainWithOptions(project, domainName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteETLJobWithOptions(project *string, etlJobName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteETLJobResponse, _err error) {
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteETLJob"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/jobs/" + tea.StringValue(etlJobName)),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("string"),
	}
	_result = &DeleteETLJobResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteETLJob(project *string, etlJobName *string) (_result *DeleteETLJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteETLJobResponse{}
	_body, _err := client.DeleteETLJobWithOptions(project, etlJobName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteExternalStoreWithOptions(project *string, externalStoreName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteExternalStoreResponse, _err error) {
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteExternalStore"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/externalstores/" + tea.StringValue(externalStoreName)),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &DeleteExternalStoreResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteExternalStore(project *string, externalStoreName *string) (_result *DeleteExternalStoreResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteExternalStoreResponse{}
	_body, _err := client.DeleteExternalStoreWithOptions(project, externalStoreName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteIndexWithOptions(project *string, logstore *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteIndexResponse, _err error) {
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteIndex"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/logstores/" + tea.StringValue(logstore) + "/index"),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &DeleteIndexResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteIndex(project *string, logstore *string) (_result *DeleteIndexResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteIndexResponse{}
	_body, _err := client.DeleteIndexWithOptions(project, logstore, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteJobWithOptions(project *string, etlJobName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteJobResponse, _err error) {
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteJob"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/jobs/" + tea.StringValue(etlJobName)),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteJobResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteJob(project *string, etlJobName *string) (_result *DeleteJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteJobResponse{}
	_body, _err := client.DeleteJobWithOptions(project, etlJobName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteLogStoreWithOptions(project *string, logstore *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteLogStoreResponse, _err error) {
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteLogStore"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/logstores/" + tea.StringValue(logstore)),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &DeleteLogStoreResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteLogStore(project *string, logstore *string) (_result *DeleteLogStoreResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteLogStoreResponse{}
	_body, _err := client.DeleteLogStoreWithOptions(project, logstore, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteLoggingWithOptions(project *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteLoggingResponse, _err error) {
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteLogging"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/logging"),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &DeleteLoggingResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteLogging(project *string) (_result *DeleteLoggingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteLoggingResponse{}
	_body, _err := client.DeleteLoggingWithOptions(project, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteMachineGroupWithOptions(project *string, machineGroup *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteMachineGroupResponse, _err error) {
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteMachineGroup"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/machinegroups/" + tea.StringValue(machineGroup)),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &DeleteMachineGroupResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteMachineGroup(project *string, machineGroup *string) (_result *DeleteMachineGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteMachineGroupResponse{}
	_body, _err := client.DeleteMachineGroupWithOptions(project, machineGroup, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteProjectWithOptions(project *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteProjectResponse, _err error) {
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteProject"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &DeleteProjectResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteProject(project *string) (_result *DeleteProjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteProjectResponse{}
	_body, _err := client.DeleteProjectWithOptions(project, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteProjectPolicyWithOptions(project *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteProjectPolicyResponse, _err error) {
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteProjectPolicy"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/policy"),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteProjectPolicyResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteProjectPolicy(project *string) (_result *DeleteProjectPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteProjectPolicyResponse{}
	_body, _err := client.DeleteProjectPolicyWithOptions(project, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteSavedSearchWithOptions(project *string, savedsearchName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteSavedSearchResponse, _err error) {
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteSavedSearch"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/savedsearches/" + tea.StringValue(savedsearchName)),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &DeleteSavedSearchResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteSavedSearch(project *string, savedsearchName *string) (_result *DeleteSavedSearchResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteSavedSearchResponse{}
	_body, _err := client.DeleteSavedSearchWithOptions(project, savedsearchName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteShipperWithOptions(project *string, logstore *string, shipperName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteShipperResponse, _err error) {
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteShipper"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/logstores/" + tea.StringValue(logstore) + "/shipper/" + tea.StringValue(shipperName)),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &DeleteShipperResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteShipper(project *string, logstore *string, shipperName *string) (_result *DeleteShipperResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteShipperResponse{}
	_body, _err := client.DeleteShipperWithOptions(project, logstore, shipperName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAlertWithOptions(project *string, alert *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetAlertResponse, _err error) {
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetAlert"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/jobs/" + tea.StringValue(alert)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAlertResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAlert(project *string, alert *string) (_result *GetAlertResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetAlertResponse{}
	_body, _err := client.GetAlertWithOptions(project, alert, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAppliedConfigsWithOptions(project *string, machineGroup *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetAppliedConfigsResponse, _err error) {
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetAppliedConfigs"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/machinegroups/" + tea.StringValue(machineGroup) + "/configs"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAppliedConfigsResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAppliedConfigs(project *string, machineGroup *string) (_result *GetAppliedConfigsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetAppliedConfigsResponse{}
	_body, _err := client.GetAppliedConfigsWithOptions(project, machineGroup, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAppliedMachineGroupsWithOptions(project *string, configName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetAppliedMachineGroupsResponse, _err error) {
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetAppliedMachineGroups"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/configs/" + tea.StringValue(configName) + "/machinegroups"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAppliedMachineGroupsResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAppliedMachineGroups(project *string, configName *string) (_result *GetAppliedMachineGroupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetAppliedMachineGroupsResponse{}
	_body, _err := client.GetAppliedMachineGroupsWithOptions(project, configName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetCheckPointWithOptions(project *string, logstore *string, consumerGroup *string, request *GetCheckPointRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetCheckPointResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Shard)) {
		query["shard"] = request.Shard
	}

	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetCheckPoint"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/logstores/" + tea.StringValue(logstore) + "/consumergroups/" + tea.StringValue(consumerGroup)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("array"),
	}
	_result = &GetCheckPointResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetCheckPoint(project *string, logstore *string, consumerGroup *string, request *GetCheckPointRequest) (_result *GetCheckPointResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetCheckPointResponse{}
	_body, _err := client.GetCheckPointWithOptions(project, logstore, consumerGroup, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetContextLogsWithOptions(project *string, logstore *string, request *GetContextLogsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetContextLogsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BackLines)) {
		query["back_lines"] = request.BackLines
	}

	if !tea.BoolValue(util.IsUnset(request.ForwardLines)) {
		query["forward_lines"] = request.ForwardLines
	}

	if !tea.BoolValue(util.IsUnset(request.PackId)) {
		query["pack_id"] = request.PackId
	}

	if !tea.BoolValue(util.IsUnset(request.PackMeta)) {
		query["pack_meta"] = request.PackMeta
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetContextLogs"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/logstores/" + tea.StringValue(logstore)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetContextLogsResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetContextLogs(project *string, logstore *string, request *GetContextLogsRequest) (_result *GetContextLogsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetContextLogsResponse{}
	_body, _err := client.GetContextLogsWithOptions(project, logstore, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetCursorWithOptions(project *string, logstore *string, shardId *string, request *GetCursorRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetCursorResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.From)) {
		query["from"] = request.From
	}

	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetCursor"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/logstores/" + tea.StringValue(logstore) + "/shards/" + tea.StringValue(shardId) + "?type=cursor"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetCursorResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetCursor(project *string, logstore *string, shardId *string, request *GetCursorRequest) (_result *GetCursorResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetCursorResponse{}
	_body, _err := client.GetCursorWithOptions(project, logstore, shardId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetCursorTimeWithOptions(project *string, logstore *string, shardId *string, request *GetCursorTimeRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetCursorTimeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Cursor)) {
		query["cursor"] = request.Cursor
	}

	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetCursorTime"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/logstores/" + tea.StringValue(logstore) + "/shards/" + tea.StringValue(shardId) + "?type=cursor_time"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetCursorTimeResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetCursorTime(project *string, logstore *string, shardId *string, request *GetCursorTimeRequest) (_result *GetCursorTimeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetCursorTimeResponse{}
	_body, _err := client.GetCursorTimeWithOptions(project, logstore, shardId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetEtlJobWithOptions(project *string, etlJobName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetEtlJobResponse, _err error) {
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetEtlJob"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/jobs/" + tea.StringValue(etlJobName)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetEtlJobResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetEtlJob(project *string, etlJobName *string) (_result *GetEtlJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetEtlJobResponse{}
	_body, _err := client.GetEtlJobWithOptions(project, etlJobName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetExternalStoreWithOptions(project *string, externalStoreName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetExternalStoreResponse, _err error) {
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetExternalStore"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/externalstores/" + tea.StringValue(externalStoreName)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetExternalStoreResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetExternalStore(project *string, externalStoreName *string) (_result *GetExternalStoreResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetExternalStoreResponse{}
	_body, _err := client.GetExternalStoreWithOptions(project, externalStoreName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetHistogramsWithOptions(project *string, logstore *string, request *GetHistogramsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetHistogramsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.From)) {
		query["from"] = request.From
	}

	if !tea.BoolValue(util.IsUnset(request.Query)) {
		query["query"] = request.Query
	}

	if !tea.BoolValue(util.IsUnset(request.To)) {
		query["to"] = request.To
	}

	if !tea.BoolValue(util.IsUnset(request.Topic)) {
		query["topic"] = request.Topic
	}

	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetHistograms"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/logstores/" + tea.StringValue(logstore) + "/index?type=histogram"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("array"),
	}
	_result = &GetHistogramsResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetHistograms(project *string, logstore *string, request *GetHistogramsRequest) (_result *GetHistogramsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetHistogramsResponse{}
	_body, _err := client.GetHistogramsWithOptions(project, logstore, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetIndexWithOptions(project *string, logstore *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetIndexResponse, _err error) {
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetIndex"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/logstores/" + tea.StringValue(logstore) + "/index"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetIndexResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetIndex(project *string, logstore *string) (_result *GetIndexResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetIndexResponse{}
	_body, _err := client.GetIndexWithOptions(project, logstore, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetJobWithOptions(project *string, etlJobName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetJobResponse, _err error) {
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetJob"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/jobs/" + tea.StringValue(etlJobName)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetJobResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetJob(project *string, etlJobName *string) (_result *GetJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetJobResponse{}
	_body, _err := client.GetJobWithOptions(project, etlJobName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetLogStoreWithOptions(project *string, logstore *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetLogStoreResponse, _err error) {
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetLogStore"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/logstores/" + tea.StringValue(logstore)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetLogStoreResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetLogStore(project *string, logstore *string) (_result *GetLogStoreResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetLogStoreResponse{}
	_body, _err := client.GetLogStoreWithOptions(project, logstore, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetLoggingWithOptions(project *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetLoggingResponse, _err error) {
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetLogging"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/logging"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetLoggingResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetLogging(project *string) (_result *GetLoggingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetLoggingResponse{}
	_body, _err := client.GetLoggingWithOptions(project, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetLogsWithOptions(project *string, logstore *string, request *GetLogsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetLogsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.From)) {
		query["from"] = request.From
	}

	if !tea.BoolValue(util.IsUnset(request.Line)) {
		query["line"] = request.Line
	}

	if !tea.BoolValue(util.IsUnset(request.Offset)) {
		query["offset"] = request.Offset
	}

	if !tea.BoolValue(util.IsUnset(request.PowerSql)) {
		query["powerSql"] = request.PowerSql
	}

	if !tea.BoolValue(util.IsUnset(request.Query)) {
		query["query"] = request.Query
	}

	if !tea.BoolValue(util.IsUnset(request.Reverse)) {
		query["reverse"] = request.Reverse
	}

	if !tea.BoolValue(util.IsUnset(request.To)) {
		query["to"] = request.To
	}

	if !tea.BoolValue(util.IsUnset(request.Topic)) {
		query["topic"] = request.Topic
	}

	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetLogs"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/logstores/" + tea.StringValue(logstore) + "?type=log"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("array"),
	}
	_result = &GetLogsResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetLogs(project *string, logstore *string, request *GetLogsRequest) (_result *GetLogsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetLogsResponse{}
	_body, _err := client.GetLogsWithOptions(project, logstore, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetMachineGroupWithOptions(project *string, machineGroup *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetMachineGroupResponse, _err error) {
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetMachineGroup"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/machinegroups/" + tea.StringValue(machineGroup)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetMachineGroupResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetMachineGroup(project *string, machineGroup *string) (_result *GetMachineGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetMachineGroupResponse{}
	_body, _err := client.GetMachineGroupWithOptions(project, machineGroup, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetProjectWithOptions(project *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetProjectResponse, _err error) {
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetProject"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetProjectResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetProject(project *string) (_result *GetProjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetProjectResponse{}
	_body, _err := client.GetProjectWithOptions(project, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetProjectLogsWithOptions(project *string, request *GetProjectLogsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetProjectLogsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PowerSql)) {
		query["powerSql"] = request.PowerSql
	}

	if !tea.BoolValue(util.IsUnset(request.Query)) {
		query["query"] = request.Query
	}

	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetProjectLogs"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/logs"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("array"),
	}
	_result = &GetProjectLogsResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetProjectLogs(project *string, request *GetProjectLogsRequest) (_result *GetProjectLogsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetProjectLogsResponse{}
	_body, _err := client.GetProjectLogsWithOptions(project, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetProjectPolicyWithOptions(project *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetProjectPolicyResponse, _err error) {
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetProjectPolicy"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/policy"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("string"),
	}
	_result = &GetProjectPolicyResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetProjectPolicy(project *string) (_result *GetProjectPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetProjectPolicyResponse{}
	_body, _err := client.GetProjectPolicyWithOptions(project, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetSavedSearchWithOptions(project *string, savedsearchName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetSavedSearchResponse, _err error) {
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetSavedSearch"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/savedsearches/" + tea.StringValue(savedsearchName)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetSavedSearchResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetSavedSearch(project *string, savedsearchName *string) (_result *GetSavedSearchResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetSavedSearchResponse{}
	_body, _err := client.GetSavedSearchWithOptions(project, savedsearchName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetShipperStatusWithOptions(project *string, logstore *string, shipperName *string, request *GetShipperStatusRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetShipperStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.From)) {
		query["from"] = request.From
	}

	if !tea.BoolValue(util.IsUnset(request.Offset)) {
		query["offset"] = request.Offset
	}

	if !tea.BoolValue(util.IsUnset(request.Size)) {
		query["size"] = request.Size
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.To)) {
		query["to"] = request.To
	}

	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetShipperStatus"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/logstores/" + tea.StringValue(logstore) + "/shipper/" + tea.StringValue(shipperName) + "/tasks"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetShipperStatusResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetShipperStatus(project *string, logstore *string, shipperName *string, request *GetShipperStatusRequest) (_result *GetShipperStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetShipperStatusResponse{}
	_body, _err := client.GetShipperStatusWithOptions(project, logstore, shipperName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListAlertWithOptions(project *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListAlertResponse, _err error) {
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("ListAlert"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/jobs"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAlertResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListAlert(project *string) (_result *ListAlertResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListAlertResponse{}
	_body, _err := client.ListAlertWithOptions(project, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListConsumerGroupWithOptions(project *string, logstore *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListConsumerGroupResponse, _err error) {
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("ListConsumerGroup"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/logstores/" + tea.StringValue(logstore) + "/consumergroups"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("array"),
	}
	_result = &ListConsumerGroupResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListConsumerGroup(project *string, logstore *string) (_result *ListConsumerGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListConsumerGroupResponse{}
	_body, _err := client.ListConsumerGroupWithOptions(project, logstore, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListDomainsWithOptions(project *string, request *ListDomainsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListDomainsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["domainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.Offset)) {
		query["offset"] = request.Offset
	}

	if !tea.BoolValue(util.IsUnset(request.Size)) {
		query["size"] = request.Size
	}

	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListDomains"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/domains"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListDomainsResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListDomains(project *string, request *ListDomainsRequest) (_result *ListDomainsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListDomainsResponse{}
	_body, _err := client.ListDomainsWithOptions(project, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListETLJobsWithOptions(project *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListETLJobsResponse, _err error) {
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("ListETLJobs"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/jobs"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListETLJobsResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListETLJobs(project *string) (_result *ListETLJobsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListETLJobsResponse{}
	_body, _err := client.ListETLJobsWithOptions(project, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListExternalStoreWithOptions(project *string, request *ListExternalStoreRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListExternalStoreResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ExternalStoreName)) {
		query["externalStoreName"] = request.ExternalStoreName
	}

	if !tea.BoolValue(util.IsUnset(request.Offset)) {
		query["offset"] = request.Offset
	}

	if !tea.BoolValue(util.IsUnset(request.Sizs)) {
		query["sizs"] = request.Sizs
	}

	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListExternalStore"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/externalstores"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListExternalStoreResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListExternalStore(project *string, request *ListExternalStoreRequest) (_result *ListExternalStoreResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListExternalStoreResponse{}
	_body, _err := client.ListExternalStoreWithOptions(project, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListJobsWithOptions(project *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListJobsResponse, _err error) {
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("ListJobs"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/jobs"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListJobsResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListJobs(project *string) (_result *ListJobsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListJobsResponse{}
	_body, _err := client.ListJobsWithOptions(project, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListLogStoresWithOptions(project *string, request *ListLogStoresRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListLogStoresResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.LogstoreName)) {
		query["logstoreName"] = request.LogstoreName
	}

	if !tea.BoolValue(util.IsUnset(request.Mode)) {
		query["mode"] = request.Mode
	}

	if !tea.BoolValue(util.IsUnset(request.Offset)) {
		query["offset"] = request.Offset
	}

	if !tea.BoolValue(util.IsUnset(request.Size)) {
		query["size"] = request.Size
	}

	if !tea.BoolValue(util.IsUnset(request.TelemetryType)) {
		query["telemetryType"] = request.TelemetryType
	}

	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListLogStores"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/logstores"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListLogStoresResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListLogStores(project *string, request *ListLogStoresRequest) (_result *ListLogStoresResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListLogStoresResponse{}
	_body, _err := client.ListLogStoresWithOptions(project, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListMachineGroupWithOptions(project *string, request *ListMachineGroupRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListMachineGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GroupName)) {
		query["groupName"] = request.GroupName
	}

	if !tea.BoolValue(util.IsUnset(request.Offset)) {
		query["offset"] = request.Offset
	}

	if !tea.BoolValue(util.IsUnset(request.Size)) {
		query["size"] = request.Size
	}

	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListMachineGroup"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/machinegroups"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListMachineGroupResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListMachineGroup(project *string, request *ListMachineGroupRequest) (_result *ListMachineGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListMachineGroupResponse{}
	_body, _err := client.ListMachineGroupWithOptions(project, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListMachinesWithOptions(project *string, machineGroup *string, request *ListMachinesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListMachinesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Offset)) {
		query["offset"] = request.Offset
	}

	if !tea.BoolValue(util.IsUnset(request.Size)) {
		query["size"] = request.Size
	}

	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListMachines"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/machinegroups/" + tea.StringValue(machineGroup) + "/machines"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListMachinesResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListMachines(project *string, machineGroup *string, request *ListMachinesRequest) (_result *ListMachinesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListMachinesResponse{}
	_body, _err := client.ListMachinesWithOptions(project, machineGroup, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListProjectWithOptions(resourceGroupId *string, request *ListProjectRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListProjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Offset)) {
		query["offset"] = request.Offset
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectName)) {
		query["projectName"] = request.ProjectName
	}

	if !tea.BoolValue(util.IsUnset(request.Size)) {
		query["size"] = request.Size
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListProject"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListProjectResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListProject(resourceGroupId *string, request *ListProjectRequest) (_result *ListProjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListProjectResponse{}
	_body, _err := client.ListProjectWithOptions(resourceGroupId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListSavedSearchWithOptions(project *string, request *ListSavedSearchRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListSavedSearchResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Offset)) {
		query["offset"] = request.Offset
	}

	if !tea.BoolValue(util.IsUnset(request.Size)) {
		query["size"] = request.Size
	}

	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListSavedSearch"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/savedsearches"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListSavedSearchResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListSavedSearch(project *string, request *ListSavedSearchRequest) (_result *ListSavedSearchResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListSavedSearchResponse{}
	_body, _err := client.ListSavedSearchWithOptions(project, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListShardsWithOptions(project *string, logstore *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListShardsResponse, _err error) {
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("ListShards"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/logstores/" + tea.StringValue(logstore) + "/shards"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("array"),
	}
	_result = &ListShardsResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListShards(project *string, logstore *string) (_result *ListShardsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListShardsResponse{}
	_body, _err := client.ListShardsWithOptions(project, logstore, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListShipperWithOptions(project *string, logstore *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListShipperResponse, _err error) {
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("ListShipper"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/logstores/" + tea.StringValue(logstore) + "/shipper"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListShipperResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListShipper(project *string, logstore *string) (_result *ListShipperResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListShipperResponse{}
	_body, _err := client.ListShipperWithOptions(project, logstore, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListTagResourcesWithOptions(tmpReq *ListTagResourcesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListTagResourcesResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListTagResourcesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.ResourceId)) {
		request.ResourceIdShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ResourceId, tea.String("resourceId"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Tags)) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, tea.String("tags"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ResourceIdShrink)) {
		query["resourceId"] = request.ResourceIdShrink
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["resourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.TagsShrink)) {
		query["tags"] = request.TagsShrink
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTagResources"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/tags"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTagResourcesResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListTagResources(request *ListTagResourcesRequest) (_result *ListTagResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListTagResourcesResponse{}
	_body, _err := client.ListTagResourcesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PullDataWithOptions(project *string, logstore *string, shard *string, request *PullDataRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *PullDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Count)) {
		query["count"] = request.Count
	}

	if !tea.BoolValue(util.IsUnset(request.Cursor)) {
		query["cursor"] = request.Cursor
	}

	if !tea.BoolValue(util.IsUnset(request.EndCursor)) {
		query["endCursor"] = request.EndCursor
	}

	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("PullData"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/logstores/" + tea.StringValue(logstore) + "/shards/" + tea.StringValue(shard) + "?type=log"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &PullDataResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PullData(project *string, logstore *string, shard *string, request *PullDataRequest) (_result *PullDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &PullDataResponse{}
	_body, _err := client.PullDataWithOptions(project, logstore, shard, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PutProjectPolicyWithOptions(project *string, request *PutProjectPolicyRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *PutProjectPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Body:    request.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("PutProjectPolicy"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/policy"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &PutProjectPolicyResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PutProjectPolicy(project *string, request *PutProjectPolicyRequest) (_result *PutProjectPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &PutProjectPolicyResponse{}
	_body, _err := client.PutProjectPolicyWithOptions(project, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PutWebtrackingWithOptions(project *string, logstoreName *string, request *PutWebtrackingRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *PutWebtrackingResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Logs)) {
		body["__logs__"] = request.Logs
	}

	if !tea.BoolValue(util.IsUnset(request.Source)) {
		body["__source__"] = request.Source
	}

	if !tea.BoolValue(util.IsUnset(request.Tags)) {
		body["__tags__"] = request.Tags
	}

	if !tea.BoolValue(util.IsUnset(request.Topic)) {
		body["__topic__"] = request.Topic
	}

	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("PutWebtracking"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/logstores/" + tea.StringValue(logstoreName) + "/track"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &PutWebtrackingResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PutWebtracking(project *string, logstoreName *string, request *PutWebtrackingRequest) (_result *PutWebtrackingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &PutWebtrackingResponse{}
	_body, _err := client.PutWebtrackingWithOptions(project, logstoreName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RemoveConfigFromMachineGroupWithOptions(project *string, machineGroup *string, configName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *RemoveConfigFromMachineGroupResponse, _err error) {
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("RemoveConfigFromMachineGroup"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/machinegroups/" + tea.StringValue(machineGroup) + "/configs/" + tea.StringValue(configName)),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &RemoveConfigFromMachineGroupResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RemoveConfigFromMachineGroup(project *string, machineGroup *string, configName *string) (_result *RemoveConfigFromMachineGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RemoveConfigFromMachineGroupResponse{}
	_body, _err := client.RemoveConfigFromMachineGroupWithOptions(project, machineGroup, configName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SplitShardWithOptions(project *string, logstore *string, shard *string, request *SplitShardRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *SplitShardResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Key)) {
		query["key"] = request.Key
	}

	if !tea.BoolValue(util.IsUnset(request.ShardCount)) {
		query["shardCount"] = request.ShardCount
	}

	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SplitShard"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/logstores/" + tea.StringValue(logstore) + "/shards/" + tea.StringValue(shard) + "?action=split"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("array"),
	}
	_result = &SplitShardResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SplitShard(project *string, logstore *string, shard *string, request *SplitShardRequest) (_result *SplitShardResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &SplitShardResponse{}
	_body, _err := client.SplitShardWithOptions(project, logstore, shard, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) TagResourcesWithOptions(request *TagResourcesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *TagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		body["resourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		body["resourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.Tags)) {
		body["tags"] = request.Tags
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("TagResources"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/tag"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &TagResourcesResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) TagResources(request *TagResourcesRequest) (_result *TagResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &TagResourcesResponse{}
	_body, _err := client.TagResourcesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UntagResourcesWithOptions(request *UntagResourcesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UntagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.All)) {
		body["all"] = request.All
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		body["resourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		body["resourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.Tags)) {
		body["tags"] = request.Tags
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UntagResources"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/untag"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UntagResourcesResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UntagResources(request *UntagResourcesRequest) (_result *UntagResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UntagResourcesResponse{}
	_body, _err := client.UntagResourcesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateAlertWithOptions(project *string, alert *string, request *UpdateAlertRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateAlertResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Alert),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateAlert"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/jobs/" + tea.StringValue(alert)),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateAlertResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateAlert(project *string, alert *string, request *UpdateAlertRequest) (_result *UpdateAlertResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateAlertResponse{}
	_body, _err := client.UpdateAlertWithOptions(project, alert, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateConsumerGroupWithOptions(project *string, logstore *string, consumerGroup *string, request *UpdateConsumerGroupRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateConsumerGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Order)) {
		body["order"] = request.Order
	}

	if !tea.BoolValue(util.IsUnset(request.Timeout)) {
		body["timeout"] = request.Timeout
	}

	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateConsumerGroup"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/logstores/" + tea.StringValue(logstore) + "/consumergroups/" + tea.StringValue(consumerGroup)),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &UpdateConsumerGroupResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateConsumerGroup(project *string, logstore *string, consumerGroup *string, request *UpdateConsumerGroupRequest) (_result *UpdateConsumerGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateConsumerGroupResponse{}
	_body, _err := client.UpdateConsumerGroupWithOptions(project, logstore, consumerGroup, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateEtlJobWithOptions(project *string, etlJob *string, request *UpdateEtlJobRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateEtlJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Enable)) {
		body["enable"] = request.Enable
	}

	if !tea.BoolValue(util.IsUnset(request.FunctionConfig)) {
		body["functionConfig"] = request.FunctionConfig
	}

	if !tea.BoolValue(util.IsUnset(request.FunctionParameter)) {
		body["functionParameter"] = request.FunctionParameter
	}

	if !tea.BoolValue(util.IsUnset(request.JobName)) {
		body["jobName"] = request.JobName
	}

	if !tea.BoolValue(util.IsUnset(request.LogConfig)) {
		body["logConfig"] = request.LogConfig
	}

	if !tea.BoolValue(util.IsUnset(request.SourceConfig)) {
		body["sourceConfig"] = request.SourceConfig
	}

	if !tea.BoolValue(util.IsUnset(request.TriggerConfig)) {
		body["triggerConfig"] = request.TriggerConfig
	}

	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateEtlJob"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/jobs/" + tea.StringValue(etlJob)),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateEtlJobResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateEtlJob(project *string, etlJob *string, request *UpdateEtlJobRequest) (_result *UpdateEtlJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateEtlJobResponse{}
	_body, _err := client.UpdateEtlJobWithOptions(project, etlJob, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateIndexWithOptions(project *string, logstore *string, request *UpdateIndexRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateIndexResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Keys)) {
		body["keys"] = request.Keys
	}

	if !tea.BoolValue(util.IsUnset(request.Line)) {
		body["line"] = request.Line
	}

	if !tea.BoolValue(util.IsUnset(request.LogReduce)) {
		body["log_reduce"] = request.LogReduce
	}

	if !tea.BoolValue(util.IsUnset(request.LogReduceBlackList)) {
		body["log_reduce_black_list"] = request.LogReduceBlackList
	}

	if !tea.BoolValue(util.IsUnset(request.LogReduceWhiteList)) {
		body["log_reduce_white_list"] = request.LogReduceWhiteList
	}

	if !tea.BoolValue(util.IsUnset(request.MaxTextLen)) {
		body["max_text_len"] = request.MaxTextLen
	}

	if !tea.BoolValue(util.IsUnset(request.Ttl)) {
		body["ttl"] = request.Ttl
	}

	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateIndex"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/logstores/" + tea.StringValue(logstore) + "/index"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &UpdateIndexResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateIndex(project *string, logstore *string, request *UpdateIndexRequest) (_result *UpdateIndexResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateIndexResponse{}
	_body, _err := client.UpdateIndexWithOptions(project, logstore, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateJobWithOptions(project *string, etlJob *string, tmpReq *UpdateJobRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateJobResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	request := &UpdateJobShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Body)) {
		request.BodyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Body, tea.String("body"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BodyShrink)) {
		query["body"] = request.BodyShrink
	}

	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateJob"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/jobs/" + tea.StringValue(etlJob)),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateJobResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateJob(project *string, etlJob *string, request *UpdateJobRequest) (_result *UpdateJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateJobResponse{}
	_body, _err := client.UpdateJobWithOptions(project, etlJob, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateLogStoreWithOptions(project *string, logstore *string, request *UpdateLogStoreRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateLogStoreResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppendMeta)) {
		body["appendMeta"] = request.AppendMeta
	}

	if !tea.BoolValue(util.IsUnset(request.AutoSplit)) {
		body["autoSplit"] = request.AutoSplit
	}

	if !tea.BoolValue(util.IsUnset(request.EnableTracking)) {
		body["enable_tracking"] = request.EnableTracking
	}

	if !tea.BoolValue(util.IsUnset(request.EncryptConf)) {
		body["encrypt_conf"] = request.EncryptConf
	}

	if !tea.BoolValue(util.IsUnset(request.HotTtl)) {
		body["hot_ttl"] = request.HotTtl
	}

	if !tea.BoolValue(util.IsUnset(request.LogstoreName)) {
		body["logstoreName"] = request.LogstoreName
	}

	if !tea.BoolValue(util.IsUnset(request.MaxSplitShard)) {
		body["maxSplitShard"] = request.MaxSplitShard
	}

	if !tea.BoolValue(util.IsUnset(request.Mode)) {
		body["mode"] = request.Mode
	}

	if !tea.BoolValue(util.IsUnset(request.ShardCount)) {
		body["shardCount"] = request.ShardCount
	}

	if !tea.BoolValue(util.IsUnset(request.TelemetryType)) {
		body["telemetryType"] = request.TelemetryType
	}

	if !tea.BoolValue(util.IsUnset(request.Ttl)) {
		body["ttl"] = request.Ttl
	}

	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateLogStore"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/logstores/" + tea.StringValue(logstore)),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &UpdateLogStoreResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateLogStore(project *string, logstore *string, request *UpdateLogStoreRequest) (_result *UpdateLogStoreResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateLogStoreResponse{}
	_body, _err := client.UpdateLogStoreWithOptions(project, logstore, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateLoggingWithOptions(project *string, request *UpdateLoggingRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateLoggingResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.LoggingDetails)) {
		body["loggingDetails"] = request.LoggingDetails
	}

	if !tea.BoolValue(util.IsUnset(request.LoggingProject)) {
		body["loggingProject"] = request.LoggingProject
	}

	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateLogging"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/logging"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &UpdateLoggingResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateLogging(project *string, request *UpdateLoggingRequest) (_result *UpdateLoggingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateLoggingResponse{}
	_body, _err := client.UpdateLoggingWithOptions(project, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateMachineGroupWithOptions(project *string, groupName *string, request *UpdateMachineGroupRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateMachineGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GroupAttribute)) {
		body["groupAttribute"] = request.GroupAttribute
	}

	if !tea.BoolValue(util.IsUnset(request.GroupName)) {
		body["groupName"] = request.GroupName
	}

	if !tea.BoolValue(util.IsUnset(request.GroupType)) {
		body["groupType"] = request.GroupType
	}

	if !tea.BoolValue(util.IsUnset(request.MachineIdentifyType)) {
		body["machineIdentifyType"] = request.MachineIdentifyType
	}

	if !tea.BoolValue(util.IsUnset(request.MachineList)) {
		body["machineList"] = request.MachineList
	}

	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateMachineGroup"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/machinegroups/" + tea.StringValue(groupName)),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &UpdateMachineGroupResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateMachineGroup(project *string, groupName *string, request *UpdateMachineGroupRequest) (_result *UpdateMachineGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateMachineGroupResponse{}
	_body, _err := client.UpdateMachineGroupWithOptions(project, groupName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateMachineGroupMachineWithOptions(project *string, machineGroup *string, request *UpdateMachineGroupMachineRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateMachineGroupMachineResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Action)) {
		query["action"] = request.Action
	}

	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    request.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateMachineGroupMachine"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/machinegroups/" + tea.StringValue(machineGroup) + "/machines"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &UpdateMachineGroupMachineResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateMachineGroupMachine(project *string, machineGroup *string, request *UpdateMachineGroupMachineRequest) (_result *UpdateMachineGroupMachineResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateMachineGroupMachineResponse{}
	_body, _err := client.UpdateMachineGroupMachineWithOptions(project, machineGroup, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateOdpsShipperWithOptions(project *string, logstore *string, shipperName *string, request *UpdateOdpsShipperRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateOdpsShipperResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ShipperName)) {
		body["shipperName"] = request.ShipperName
	}

	if !tea.BoolValue(util.IsUnset(request.TargetConfiguration)) {
		body["targetConfiguration"] = request.TargetConfiguration
	}

	if !tea.BoolValue(util.IsUnset(request.TargetType)) {
		body["targetType"] = request.TargetType
	}

	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateOdpsShipper"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/logstores/" + tea.StringValue(logstore) + "/shipper/" + tea.StringValue(shipperName)),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &UpdateOdpsShipperResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateOdpsShipper(project *string, logstore *string, shipperName *string, request *UpdateOdpsShipperRequest) (_result *UpdateOdpsShipperResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateOdpsShipperResponse{}
	_body, _err := client.UpdateOdpsShipperWithOptions(project, logstore, shipperName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateOssExternalStoreWithOptions(project *string, externalStoreName *string, request *UpdateOssExternalStoreRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateOssExternalStoreResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ExternalStoreName)) {
		body["externalStoreName"] = request.ExternalStoreName
	}

	if !tea.BoolValue(util.IsUnset(request.Parameter)) {
		body["parameter"] = request.Parameter
	}

	if !tea.BoolValue(util.IsUnset(request.StoreType)) {
		body["storeType"] = request.StoreType
	}

	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateOssExternalStore"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/externalstores/" + tea.StringValue(externalStoreName)),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &UpdateOssExternalStoreResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateOssExternalStore(project *string, externalStoreName *string, request *UpdateOssExternalStoreRequest) (_result *UpdateOssExternalStoreResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateOssExternalStoreResponse{}
	_body, _err := client.UpdateOssExternalStoreWithOptions(project, externalStoreName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateOssShipperWithOptions(project *string, logstore *string, shipperName *string, request *UpdateOssShipperRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateOssShipperResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ShipperName)) {
		body["shipperName"] = request.ShipperName
	}

	if !tea.BoolValue(util.IsUnset(request.TargetConfiguration)) {
		body["targetConfiguration"] = request.TargetConfiguration
	}

	if !tea.BoolValue(util.IsUnset(request.TargetType)) {
		body["targetType"] = request.TargetType
	}

	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateOssShipper"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/logstores/" + tea.StringValue(logstore) + "/shipper/" + tea.StringValue(shipperName)),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &UpdateOssShipperResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateOssShipper(project *string, logstore *string, shipperName *string, request *UpdateOssShipperRequest) (_result *UpdateOssShipperResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateOssShipperResponse{}
	_body, _err := client.UpdateOssShipperWithOptions(project, logstore, shipperName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateProjectWithOptions(project *string, request *UpdateProjectRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateProjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateProject"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &UpdateProjectResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateProject(project *string, request *UpdateProjectRequest) (_result *UpdateProjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateProjectResponse{}
	_body, _err := client.UpdateProjectWithOptions(project, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateRdsExternalStoreWithOptions(project *string, externalStoreName *string, request *UpdateRdsExternalStoreRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateRdsExternalStoreResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ExternalStoreName)) {
		body["externalStoreName"] = request.ExternalStoreName
	}

	if !tea.BoolValue(util.IsUnset(request.Parameter)) {
		body["parameter"] = request.Parameter
	}

	if !tea.BoolValue(util.IsUnset(request.StoreType)) {
		body["storeType"] = request.StoreType
	}

	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateRdsExternalStore"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/externalstores/" + tea.StringValue(externalStoreName)),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &UpdateRdsExternalStoreResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateRdsExternalStore(project *string, externalStoreName *string, request *UpdateRdsExternalStoreRequest) (_result *UpdateRdsExternalStoreResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateRdsExternalStoreResponse{}
	_body, _err := client.UpdateRdsExternalStoreWithOptions(project, externalStoreName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateSavedSearchWithOptions(project *string, savedsearchName *string, request *UpdateSavedSearchRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateSavedSearchResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DisplayName)) {
		body["displayName"] = request.DisplayName
	}

	if !tea.BoolValue(util.IsUnset(request.Logstore)) {
		body["logstore"] = request.Logstore
	}

	if !tea.BoolValue(util.IsUnset(request.SavedsearchName)) {
		body["savedsearchName"] = request.SavedsearchName
	}

	if !tea.BoolValue(util.IsUnset(request.SearchQuery)) {
		body["searchQuery"] = request.SearchQuery
	}

	if !tea.BoolValue(util.IsUnset(request.Topic)) {
		body["topic"] = request.Topic
	}

	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateSavedSearch"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/savedsearches/" + tea.StringValue(savedsearchName)),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &UpdateSavedSearchResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateSavedSearch(project *string, savedsearchName *string, request *UpdateSavedSearchRequest) (_result *UpdateSavedSearchResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateSavedSearchResponse{}
	_body, _err := client.UpdateSavedSearchWithOptions(project, savedsearchName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
