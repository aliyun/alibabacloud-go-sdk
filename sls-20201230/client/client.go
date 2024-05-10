// This file is auto-generated, don't edit it. Thanks.
package client

import (
	gatewayclient "github.com/alibabacloud-go/alibabacloud-gateway-sls/client"
	spi "github.com/alibabacloud-go/alibabacloud-gateway-spi/client"
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type Alert struct {
	// This parameter is required.
	Configuration *AlertConfiguration `json:"configuration,omitempty" xml:"configuration,omitempty"`
	CreateTime    *int64              `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// Alert Desc
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// alertNameExample
	DisplayName      *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	LastModifiedTime *int64  `json:"lastModifiedTime,omitempty" xml:"lastModifiedTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// alert-123456
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// This parameter is required.
	Schedule *Schedule `json:"schedule,omitempty" xml:"schedule,omitempty"`
	// example:
	//
	// ENABLED/DISABLED
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s Alert) String() string {
	return tea.Prettify(s)
}

func (s Alert) GoString() string {
	return s.String()
}

func (s *Alert) SetConfiguration(v *AlertConfiguration) *Alert {
	s.Configuration = v
	return s
}

func (s *Alert) SetCreateTime(v int64) *Alert {
	s.CreateTime = &v
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

func (s *Alert) SetLastModifiedTime(v int64) *Alert {
	s.LastModifiedTime = &v
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

func (s *Alert) SetStatus(v string) *Alert {
	s.Status = &v
	return s
}

type AlertConfiguration struct {
	Annotations []*AlertTag `json:"annotations,omitempty" xml:"annotations,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// false
	AutoAnnotation         *bool                   `json:"autoAnnotation,omitempty" xml:"autoAnnotation,omitempty"`
	ConditionConfiguration *ConditionConfiguration `json:"conditionConfiguration,omitempty" xml:"conditionConfiguration,omitempty"`
	// example:
	//
	// dasnboardExample
	Dashboard *string `json:"dashboard,omitempty" xml:"dashboard,omitempty"`
	// This parameter is required.
	GroupConfiguration *GroupConfiguration  `json:"groupConfiguration,omitempty" xml:"groupConfiguration,omitempty"`
	JoinConfigurations []*JoinConfiguration `json:"joinConfigurations,omitempty" xml:"joinConfigurations,omitempty" type:"Repeated"`
	Labels             []*AlertTag          `json:"labels,omitempty" xml:"labels,omitempty" type:"Repeated"`
	// example:
	//
	// 1698907508
	MuteUntil *int64 `json:"muteUntil,omitempty" xml:"muteUntil,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// false
	NoDataFire *bool `json:"noDataFire,omitempty" xml:"noDataFire,omitempty"`
	// example:
	//
	// 6
	NoDataSeverity      *int32               `json:"noDataSeverity,omitempty" xml:"noDataSeverity,omitempty"`
	PolicyConfiguration *PolicyConfiguration `json:"policyConfiguration,omitempty" xml:"policyConfiguration,omitempty"`
	// This parameter is required.
	QueryList []*AlertQuery `json:"queryList,omitempty" xml:"queryList,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// false
	SendResolved *bool `json:"sendResolved,omitempty" xml:"sendResolved,omitempty"`
	// This parameter is required.
	SeverityConfigurations []*SeverityConfiguration     `json:"severityConfigurations,omitempty" xml:"severityConfigurations,omitempty" type:"Repeated"`
	SinkAlerthub           *SinkAlerthubConfiguration   `json:"sinkAlerthub,omitempty" xml:"sinkAlerthub,omitempty"`
	SinkCms                *SinkCmsConfiguration        `json:"sinkCms,omitempty" xml:"sinkCms,omitempty"`
	SinkEventStore         *SinkEventStoreConfiguration `json:"sinkEventStore,omitempty" xml:"sinkEventStore,omitempty"`
	Tags                   []*string                    `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
	TemplateConfiguration  *TemplateConfiguration       `json:"templateConfiguration,omitempty" xml:"templateConfiguration,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	Threshold *int32 `json:"threshold,omitempty" xml:"threshold,omitempty"`
	// example:
	//
	// default
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2.0
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s AlertConfiguration) String() string {
	return tea.Prettify(s)
}

func (s AlertConfiguration) GoString() string {
	return s.String()
}

func (s *AlertConfiguration) SetAnnotations(v []*AlertTag) *AlertConfiguration {
	s.Annotations = v
	return s
}

func (s *AlertConfiguration) SetAutoAnnotation(v bool) *AlertConfiguration {
	s.AutoAnnotation = &v
	return s
}

func (s *AlertConfiguration) SetConditionConfiguration(v *ConditionConfiguration) *AlertConfiguration {
	s.ConditionConfiguration = v
	return s
}

func (s *AlertConfiguration) SetDashboard(v string) *AlertConfiguration {
	s.Dashboard = &v
	return s
}

func (s *AlertConfiguration) SetGroupConfiguration(v *GroupConfiguration) *AlertConfiguration {
	s.GroupConfiguration = v
	return s
}

func (s *AlertConfiguration) SetJoinConfigurations(v []*JoinConfiguration) *AlertConfiguration {
	s.JoinConfigurations = v
	return s
}

func (s *AlertConfiguration) SetLabels(v []*AlertTag) *AlertConfiguration {
	s.Labels = v
	return s
}

func (s *AlertConfiguration) SetMuteUntil(v int64) *AlertConfiguration {
	s.MuteUntil = &v
	return s
}

func (s *AlertConfiguration) SetNoDataFire(v bool) *AlertConfiguration {
	s.NoDataFire = &v
	return s
}

func (s *AlertConfiguration) SetNoDataSeverity(v int32) *AlertConfiguration {
	s.NoDataSeverity = &v
	return s
}

func (s *AlertConfiguration) SetPolicyConfiguration(v *PolicyConfiguration) *AlertConfiguration {
	s.PolicyConfiguration = v
	return s
}

func (s *AlertConfiguration) SetQueryList(v []*AlertQuery) *AlertConfiguration {
	s.QueryList = v
	return s
}

func (s *AlertConfiguration) SetSendResolved(v bool) *AlertConfiguration {
	s.SendResolved = &v
	return s
}

func (s *AlertConfiguration) SetSeverityConfigurations(v []*SeverityConfiguration) *AlertConfiguration {
	s.SeverityConfigurations = v
	return s
}

func (s *AlertConfiguration) SetSinkAlerthub(v *SinkAlerthubConfiguration) *AlertConfiguration {
	s.SinkAlerthub = v
	return s
}

func (s *AlertConfiguration) SetSinkCms(v *SinkCmsConfiguration) *AlertConfiguration {
	s.SinkCms = v
	return s
}

func (s *AlertConfiguration) SetSinkEventStore(v *SinkEventStoreConfiguration) *AlertConfiguration {
	s.SinkEventStore = v
	return s
}

func (s *AlertConfiguration) SetTags(v []*string) *AlertConfiguration {
	s.Tags = v
	return s
}

func (s *AlertConfiguration) SetTemplateConfiguration(v *TemplateConfiguration) *AlertConfiguration {
	s.TemplateConfiguration = v
	return s
}

func (s *AlertConfiguration) SetThreshold(v int32) *AlertConfiguration {
	s.Threshold = &v
	return s
}

func (s *AlertConfiguration) SetType(v string) *AlertConfiguration {
	s.Type = &v
	return s
}

func (s *AlertConfiguration) SetVersion(v string) *AlertConfiguration {
	s.Version = &v
	return s
}

type AlertQuery struct {
	// example:
	//
	// chartExmaple
	ChartTitle *string `json:"chartTitle,omitempty" xml:"chartTitle,omitempty"`
	// example:
	//
	// dashboardExample
	DashboardId *string `json:"dashboardId,omitempty" xml:"dashboardId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// now
	End *string `json:"end,omitempty" xml:"end,omitempty"`
	// example:
	//
	// auto
	PowerSqlMode *string `json:"powerSqlMode,omitempty" xml:"powerSqlMode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// projectExample
	Project *string `json:"project,omitempty" xml:"project,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 	- | select *
	Query *string `json:"query,omitempty" xml:"query,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	Region  *string `json:"region,omitempty" xml:"region,omitempty"`
	RoleArn *string `json:"roleArn,omitempty" xml:"roleArn,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// -5m
	Start *string `json:"start,omitempty" xml:"start,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// store Example
	Store *string `json:"store,omitempty" xml:"store,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// log
	StoreType *string `json:"storeType,omitempty" xml:"storeType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Relative
	TimeSpanType *string `json:"timeSpanType,omitempty" xml:"timeSpanType,omitempty"`
	Ui           *string `json:"ui,omitempty" xml:"ui,omitempty"`
}

func (s AlertQuery) String() string {
	return tea.Prettify(s)
}

func (s AlertQuery) GoString() string {
	return s.String()
}

func (s *AlertQuery) SetChartTitle(v string) *AlertQuery {
	s.ChartTitle = &v
	return s
}

func (s *AlertQuery) SetDashboardId(v string) *AlertQuery {
	s.DashboardId = &v
	return s
}

func (s *AlertQuery) SetEnd(v string) *AlertQuery {
	s.End = &v
	return s
}

func (s *AlertQuery) SetPowerSqlMode(v string) *AlertQuery {
	s.PowerSqlMode = &v
	return s
}

func (s *AlertQuery) SetProject(v string) *AlertQuery {
	s.Project = &v
	return s
}

func (s *AlertQuery) SetQuery(v string) *AlertQuery {
	s.Query = &v
	return s
}

func (s *AlertQuery) SetRegion(v string) *AlertQuery {
	s.Region = &v
	return s
}

func (s *AlertQuery) SetRoleArn(v string) *AlertQuery {
	s.RoleArn = &v
	return s
}

func (s *AlertQuery) SetStart(v string) *AlertQuery {
	s.Start = &v
	return s
}

func (s *AlertQuery) SetStore(v string) *AlertQuery {
	s.Store = &v
	return s
}

func (s *AlertQuery) SetStoreType(v string) *AlertQuery {
	s.StoreType = &v
	return s
}

func (s *AlertQuery) SetTimeSpanType(v string) *AlertQuery {
	s.TimeSpanType = &v
	return s
}

func (s *AlertQuery) SetUi(v string) *AlertQuery {
	s.Ui = &v
	return s
}

type AlertTag struct {
	// example:
	//
	// title
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// example:
	//
	// example value
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s AlertTag) String() string {
	return tea.Prettify(s)
}

func (s AlertTag) GoString() string {
	return s.String()
}

func (s *AlertTag) SetKey(v string) *AlertTag {
	s.Key = &v
	return s
}

func (s *AlertTag) SetValue(v string) *AlertTag {
	s.Value = &v
	return s
}

type ConditionConfiguration struct {
	// example:
	//
	// cnt > 100
	Condition *string `json:"condition,omitempty" xml:"condition,omitempty"`
	// example:
	//
	// __count__ > 5
	CountCondition *string `json:"countCondition,omitempty" xml:"countCondition,omitempty"`
}

func (s ConditionConfiguration) String() string {
	return tea.Prettify(s)
}

func (s ConditionConfiguration) GoString() string {
	return s.String()
}

func (s *ConditionConfiguration) SetCondition(v string) *ConditionConfiguration {
	s.Condition = &v
	return s
}

func (s *ConditionConfiguration) SetCountCondition(v string) *ConditionConfiguration {
	s.CountCondition = &v
	return s
}

type ConsumerGroup struct {
	// example:
	//
	// test-group
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// false
	Order *bool `json:"order,omitempty" xml:"order,omitempty"`
	// example:
	//
	// 300
	Timeout *int32 `json:"timeout,omitempty" xml:"timeout,omitempty"`
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

type ETL struct {
	// This parameter is required.
	Configuration *ETLConfiguration `json:"configuration,omitempty" xml:"configuration,omitempty"`
	// example:
	//
	// 1714274900
	CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// 加工作业
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// etljob
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	// example:
	//
	// 1714274900
	LastModifiedTime *int64 `json:"lastModifiedTime,omitempty" xml:"lastModifiedTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// etl-20240426
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// f0eb655e501a8780808d1970ef6d04c4
	ScheduleId *string `json:"scheduleId,omitempty" xml:"scheduleId,omitempty"`
	// example:
	//
	// RUNNING
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s ETL) String() string {
	return tea.Prettify(s)
}

func (s ETL) GoString() string {
	return s.String()
}

func (s *ETL) SetConfiguration(v *ETLConfiguration) *ETL {
	s.Configuration = v
	return s
}

func (s *ETL) SetCreateTime(v int64) *ETL {
	s.CreateTime = &v
	return s
}

func (s *ETL) SetDescription(v string) *ETL {
	s.Description = &v
	return s
}

func (s *ETL) SetDisplayName(v string) *ETL {
	s.DisplayName = &v
	return s
}

func (s *ETL) SetLastModifiedTime(v int64) *ETL {
	s.LastModifiedTime = &v
	return s
}

func (s *ETL) SetName(v string) *ETL {
	s.Name = &v
	return s
}

func (s *ETL) SetScheduleId(v string) *ETL {
	s.ScheduleId = &v
	return s
}

func (s *ETL) SetStatus(v string) *ETL {
	s.Status = &v
	return s
}

type ETLConfiguration struct {
	// Deprecated
	AccessKeyId *string `json:"accessKeyId,omitempty" xml:"accessKeyId,omitempty"`
	// Deprecated
	AccessKeySecret *string `json:"accessKeySecret,omitempty" xml:"accessKeySecret,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1714274900
	FromTime *int64 `json:"fromTime,omitempty" xml:"fromTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test-logstore
	Logstore *string `json:"logstore,omitempty" xml:"logstore,omitempty"`
	// example:
	//
	// config.vpc.vpc_id.test1：vpc-uf6mskb0b****n9yj
	Parameters map[string]interface{} `json:"parameters,omitempty" xml:"parameters,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// acs:ram::13234:role/logtarget
	RoleArn *string `json:"roleArn,omitempty" xml:"roleArn,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// e_set("key","value")
	Script *string `json:"script,omitempty" xml:"script,omitempty"`
	// This parameter is required.
	Sinks []*ETLConfigurationSink `json:"sinks,omitempty" xml:"sinks,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 1714274970
	ToTime *int64 `json:"toTime,omitempty" xml:"toTime,omitempty"`
}

func (s ETLConfiguration) String() string {
	return tea.Prettify(s)
}

func (s ETLConfiguration) GoString() string {
	return s.String()
}

func (s *ETLConfiguration) SetAccessKeyId(v string) *ETLConfiguration {
	s.AccessKeyId = &v
	return s
}

func (s *ETLConfiguration) SetAccessKeySecret(v string) *ETLConfiguration {
	s.AccessKeySecret = &v
	return s
}

func (s *ETLConfiguration) SetFromTime(v int64) *ETLConfiguration {
	s.FromTime = &v
	return s
}

func (s *ETLConfiguration) SetLogstore(v string) *ETLConfiguration {
	s.Logstore = &v
	return s
}

func (s *ETLConfiguration) SetParameters(v map[string]interface{}) *ETLConfiguration {
	s.Parameters = v
	return s
}

func (s *ETLConfiguration) SetRoleArn(v string) *ETLConfiguration {
	s.RoleArn = &v
	return s
}

func (s *ETLConfiguration) SetScript(v string) *ETLConfiguration {
	s.Script = &v
	return s
}

func (s *ETLConfiguration) SetSinks(v []*ETLConfigurationSink) *ETLConfiguration {
	s.Sinks = v
	return s
}

func (s *ETLConfiguration) SetToTime(v int64) *ETLConfiguration {
	s.ToTime = &v
	return s
}

type ETLConfigurationSink struct {
	// Deprecated
	AccessKeyId *string `json:"accessKeyId,omitempty" xml:"accessKeyId,omitempty"`
	// Deprecated
	AccessKeySecret *string `json:"accessKeySecret,omitempty" xml:"accessKeySecret,omitempty"`
	Endpoint        *string `json:"endpoint,omitempty" xml:"endpoint,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test-logstore
	Logstore *string `json:"logstore,omitempty" xml:"logstore,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test-etljob
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test-project
	Project *string `json:"project,omitempty" xml:"project,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// acs:ram::13234:role/logtarget
	RoleArn *string `json:"roleArn,omitempty" xml:"roleArn,omitempty"`
}

func (s ETLConfigurationSink) String() string {
	return tea.Prettify(s)
}

func (s ETLConfigurationSink) GoString() string {
	return s.String()
}

func (s *ETLConfigurationSink) SetAccessKeyId(v string) *ETLConfigurationSink {
	s.AccessKeyId = &v
	return s
}

func (s *ETLConfigurationSink) SetAccessKeySecret(v string) *ETLConfigurationSink {
	s.AccessKeySecret = &v
	return s
}

func (s *ETLConfigurationSink) SetEndpoint(v string) *ETLConfigurationSink {
	s.Endpoint = &v
	return s
}

func (s *ETLConfigurationSink) SetLogstore(v string) *ETLConfigurationSink {
	s.Logstore = &v
	return s
}

func (s *ETLConfigurationSink) SetName(v string) *ETLConfigurationSink {
	s.Name = &v
	return s
}

func (s *ETLConfigurationSink) SetProject(v string) *ETLConfigurationSink {
	s.Project = &v
	return s
}

func (s *ETLConfigurationSink) SetRoleArn(v string) *ETLConfigurationSink {
	s.RoleArn = &v
	return s
}

type EncryptConf struct {
	// This parameter is required.
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
	// example:
	//
	// default
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
	// This parameter is required.
	Arn *string `json:"arn,omitempty" xml:"arn,omitempty"`
	// This parameter is required.
	CmkKeyId *string `json:"cmk_key_id,omitempty" xml:"cmk_key_id,omitempty"`
	// This parameter is required.
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

type GroupConfiguration struct {
	Fields []*string `json:"fields,omitempty" xml:"fields,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// custom
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GroupConfiguration) String() string {
	return tea.Prettify(s)
}

func (s GroupConfiguration) GoString() string {
	return s.String()
}

func (s *GroupConfiguration) SetFields(v []*string) *GroupConfiguration {
	s.Fields = v
	return s
}

func (s *GroupConfiguration) SetType(v string) *GroupConfiguration {
	s.Type = &v
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

type JoinConfiguration struct {
	// example:
	//
	// $0.id == $1.id
	Condition *string `json:"condition,omitempty" xml:"condition,omitempty"`
	// example:
	//
	// left_join
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s JoinConfiguration) String() string {
	return tea.Prettify(s)
}

func (s JoinConfiguration) GoString() string {
	return s.String()
}

func (s *JoinConfiguration) SetCondition(v string) *JoinConfiguration {
	s.Condition = &v
	return s
}

func (s *JoinConfiguration) SetType(v string) *JoinConfiguration {
	s.Type = &v
	return s
}

type LogContent struct {
	// This parameter is required.
	//
	// example:
	//
	// key-test
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// value-test
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s LogContent) String() string {
	return tea.Prettify(s)
}

func (s LogContent) GoString() string {
	return s.String()
}

func (s *LogContent) SetKey(v string) *LogContent {
	s.Key = &v
	return s
}

func (s *LogContent) SetValue(v string) *LogContent {
	s.Value = &v
	return s
}

type LogGroup struct {
	// This parameter is required.
	LogTags []*LogTag `json:"LogTags,omitempty" xml:"LogTags,omitempty" type:"Repeated"`
	// This parameter is required.
	Logs []*LogItem `json:"Logs,omitempty" xml:"Logs,omitempty" type:"Repeated"`
	// example:
	//
	// 192.1.1.1
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// topic-test
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
}

func (s LogGroup) String() string {
	return tea.Prettify(s)
}

func (s LogGroup) GoString() string {
	return s.String()
}

func (s *LogGroup) SetLogTags(v []*LogTag) *LogGroup {
	s.LogTags = v
	return s
}

func (s *LogGroup) SetLogs(v []*LogItem) *LogGroup {
	s.Logs = v
	return s
}

func (s *LogGroup) SetSource(v string) *LogGroup {
	s.Source = &v
	return s
}

func (s *LogGroup) SetTopic(v string) *LogGroup {
	s.Topic = &v
	return s
}

type LogItem struct {
	// This parameter is required.
	Contents []*LogContent `json:"Contents,omitempty" xml:"Contents,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 1690254376
	Time *int32 `json:"Time,omitempty" xml:"Time,omitempty"`
}

func (s LogItem) String() string {
	return tea.Prettify(s)
}

func (s LogItem) GoString() string {
	return s.String()
}

func (s *LogItem) SetContents(v []*LogContent) *LogItem {
	s.Contents = v
	return s
}

func (s *LogItem) SetTime(v int32) *LogItem {
	s.Time = &v
	return s
}

type LogTag struct {
	// This parameter is required.
	//
	// example:
	//
	// key-test
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// value-test
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s LogTag) String() string {
	return tea.Prettify(s)
}

func (s LogTag) GoString() string {
	return s.String()
}

func (s *LogTag) SetKey(v string) *LogTag {
	s.Key = &v
	return s
}

func (s *LogTag) SetValue(v string) *LogTag {
	s.Value = &v
	return s
}

type LogtailConfig struct {
	// This parameter is required.
	//
	// example:
	//
	// test-config
	ConfigName *string `json:"configName,omitempty" xml:"configName,omitempty"`
	// example:
	//
	// 1655176807
	CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// This parameter is required.
	InputDetail map[string]interface{} `json:"inputDetail,omitempty" xml:"inputDetail,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// file
	InputType *string `json:"inputType,omitempty" xml:"inputType,omitempty"`
	// example:
	//
	// 1655176807
	LastModifyTime *int64 `json:"lastModifyTime,omitempty" xml:"lastModifyTime,omitempty"`
	// example:
	//
	// 2022-06-14 11:13:29.796 | DEBUG    | __main__:<module>:1 - hello world
	LogSample *string `json:"logSample,omitempty" xml:"logSample,omitempty"`
	// This parameter is required.
	OutputDetail *LogtailConfigOutputDetail `json:"outputDetail,omitempty" xml:"outputDetail,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// LogService
	OutputType *string `json:"outputType,omitempty" xml:"outputType,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou-intranet.log.aliyuncs.com
	Endpoint *string `json:"endpoint,omitempty" xml:"endpoint,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test-logstore
	LogstoreName *string `json:"logstoreName,omitempty" xml:"logstoreName,omitempty"`
	// example:
	//
	// cn-hangzhou
	Region *string `json:"region,omitempty" xml:"region,omitempty"`
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

type LogtailPipelineConfig struct {
	Aggregators []map[string]interface{} `json:"aggregators,omitempty" xml:"aggregators,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// test-config
	ConfigName *string `json:"configName,omitempty" xml:"configName,omitempty"`
	// example:
	//
	// 1655176807
	CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// This parameter is required.
	Flushers []map[string]interface{} `json:"flushers,omitempty" xml:"flushers,omitempty" type:"Repeated"`
	Global   map[string]interface{}   `json:"global,omitempty" xml:"global,omitempty"`
	// This parameter is required.
	Inputs []map[string]interface{} `json:"inputs,omitempty" xml:"inputs,omitempty" type:"Repeated"`
	// example:
	//
	// 1655176807
	LastModifyTime *int64 `json:"lastModifyTime,omitempty" xml:"lastModifyTime,omitempty"`
	// example:
	//
	// 127.0.0.1 - - [10/Jun/2022:12:36:49 +0800] "GET /index.html HTTP/1.1" 200
	LogSample  *string                  `json:"logSample,omitempty" xml:"logSample,omitempty"`
	Processors []map[string]interface{} `json:"processors,omitempty" xml:"processors,omitempty" type:"Repeated"`
}

func (s LogtailPipelineConfig) String() string {
	return tea.Prettify(s)
}

func (s LogtailPipelineConfig) GoString() string {
	return s.String()
}

func (s *LogtailPipelineConfig) SetAggregators(v []map[string]interface{}) *LogtailPipelineConfig {
	s.Aggregators = v
	return s
}

func (s *LogtailPipelineConfig) SetConfigName(v string) *LogtailPipelineConfig {
	s.ConfigName = &v
	return s
}

func (s *LogtailPipelineConfig) SetCreateTime(v int64) *LogtailPipelineConfig {
	s.CreateTime = &v
	return s
}

func (s *LogtailPipelineConfig) SetFlushers(v []map[string]interface{}) *LogtailPipelineConfig {
	s.Flushers = v
	return s
}

func (s *LogtailPipelineConfig) SetGlobal(v map[string]interface{}) *LogtailPipelineConfig {
	s.Global = v
	return s
}

func (s *LogtailPipelineConfig) SetInputs(v []map[string]interface{}) *LogtailPipelineConfig {
	s.Inputs = v
	return s
}

func (s *LogtailPipelineConfig) SetLastModifyTime(v int64) *LogtailPipelineConfig {
	s.LastModifyTime = &v
	return s
}

func (s *LogtailPipelineConfig) SetLogSample(v string) *LogtailPipelineConfig {
	s.LogSample = &v
	return s
}

func (s *LogtailPipelineConfig) SetProcessors(v []map[string]interface{}) *LogtailPipelineConfig {
	s.Processors = v
	return s
}

type MLDataParam struct {
	// example:
	//
	// dc74b0f569126bb310e1ba6454c351ac
	AnnotationdataId *string                                 `json:"annotationdataId,omitempty" xml:"annotationdataId,omitempty"`
	Annotations      map[string]*MLDataParamAnnotationsValue `json:"annotations,omitempty" xml:"annotations,omitempty"`
	Config           map[string]*string                      `json:"config,omitempty" xml:"config,omitempty"`
	// example:
	//
	// 1695094335
	CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// 59db060bd89468245d76416a68a510ac
	DataHash *string `json:"dataHash,omitempty" xml:"dataHash,omitempty"`
	// example:
	//
	// a9bd488f6dd42d294495fb780858e83d
	DatasetId *string `json:"datasetId,omitempty" xml:"datasetId,omitempty"`
	// example:
	//
	// 1695094335
	LastModifyTime *int64                                  `json:"lastModifyTime,omitempty" xml:"lastModifyTime,omitempty"`
	Predictions    map[string]*MLDataParamPredictionsValue `json:"predictions,omitempty" xml:"predictions,omitempty"`
	// example:
	//
	// xxx/xxx/xxx/
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
	// example:
	//
	// oss
	ValueType *string `json:"valueType,omitempty" xml:"valueType,omitempty"`
}

func (s MLDataParam) String() string {
	return tea.Prettify(s)
}

func (s MLDataParam) GoString() string {
	return s.String()
}

func (s *MLDataParam) SetAnnotationdataId(v string) *MLDataParam {
	s.AnnotationdataId = &v
	return s
}

func (s *MLDataParam) SetAnnotations(v map[string]*MLDataParamAnnotationsValue) *MLDataParam {
	s.Annotations = v
	return s
}

func (s *MLDataParam) SetConfig(v map[string]*string) *MLDataParam {
	s.Config = v
	return s
}

func (s *MLDataParam) SetCreateTime(v int64) *MLDataParam {
	s.CreateTime = &v
	return s
}

func (s *MLDataParam) SetDataHash(v string) *MLDataParam {
	s.DataHash = &v
	return s
}

func (s *MLDataParam) SetDatasetId(v string) *MLDataParam {
	s.DatasetId = &v
	return s
}

func (s *MLDataParam) SetLastModifyTime(v int64) *MLDataParam {
	s.LastModifyTime = &v
	return s
}

func (s *MLDataParam) SetPredictions(v map[string]*MLDataParamPredictionsValue) *MLDataParam {
	s.Predictions = v
	return s
}

func (s *MLDataParam) SetValue(v string) *MLDataParam {
	s.Value = &v
	return s
}

func (s *MLDataParam) SetValueType(v string) *MLDataParam {
	s.ValueType = &v
	return s
}

type MLDataSetParam struct {
	// example:
	//
	// sls-console
	CreateBy *string `json:"createBy,omitempty" xml:"createBy,omitempty"`
	// example:
	//
	// 1695090077
	CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// Metric
	DataType *string `json:"dataType,omitempty" xml:"dataType,omitempty"`
	// example:
	//
	// d9bd488f6dd42d294495fb780858e83d
	DatasetId *string `json:"datasetId,omitempty" xml:"datasetId,omitempty"`
	// example:
	//
	// 数据集A
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// a191ae4ca615b0ccb93c211fc8a998af
	LabelId *string `json:"labelId,omitempty" xml:"labelId,omitempty"`
	// example:
	//
	// 1695090077
	LastModifyTime *int64 `json:"lastModifyTime,omitempty" xml:"lastModifyTime,omitempty"`
	// example:
	//
	// sls_builtin_dataset_metric.shapeclassification.anomalydetection
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// Metric.ShapeClassification.AnomalyDetection
	SettingType *string `json:"settingType,omitempty" xml:"settingType,omitempty"`
}

func (s MLDataSetParam) String() string {
	return tea.Prettify(s)
}

func (s MLDataSetParam) GoString() string {
	return s.String()
}

func (s *MLDataSetParam) SetCreateBy(v string) *MLDataSetParam {
	s.CreateBy = &v
	return s
}

func (s *MLDataSetParam) SetCreateTime(v int64) *MLDataSetParam {
	s.CreateTime = &v
	return s
}

func (s *MLDataSetParam) SetDataType(v string) *MLDataSetParam {
	s.DataType = &v
	return s
}

func (s *MLDataSetParam) SetDatasetId(v string) *MLDataSetParam {
	s.DatasetId = &v
	return s
}

func (s *MLDataSetParam) SetDescription(v string) *MLDataSetParam {
	s.Description = &v
	return s
}

func (s *MLDataSetParam) SetLabelId(v string) *MLDataSetParam {
	s.LabelId = &v
	return s
}

func (s *MLDataSetParam) SetLastModifyTime(v int64) *MLDataSetParam {
	s.LastModifyTime = &v
	return s
}

func (s *MLDataSetParam) SetName(v string) *MLDataSetParam {
	s.Name = &v
	return s
}

func (s *MLDataSetParam) SetSettingType(v string) *MLDataSetParam {
	s.SettingType = &v
	return s
}

type MLLabelParam struct {
	// example:
	//
	// 1695090077
	CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// 默认表
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// abbd488f6dd42d294495fb780858e83d
	LabelId *string `json:"labelId,omitempty" xml:"labelId,omitempty"`
	// example:
	//
	// 1695090077
	LastModifyTime *int64 `json:"lastModifyTime,omitempty" xml:"lastModifyTime,omitempty"`
	// example:
	//
	// 标签表
	Name     *string                 `json:"name,omitempty" xml:"name,omitempty"`
	Settings []*MLLabelParamSettings `json:"settings,omitempty" xml:"settings,omitempty" type:"Repeated"`
	// example:
	//
	// xxx
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s MLLabelParam) String() string {
	return tea.Prettify(s)
}

func (s MLLabelParam) GoString() string {
	return s.String()
}

func (s *MLLabelParam) SetCreateTime(v int64) *MLLabelParam {
	s.CreateTime = &v
	return s
}

func (s *MLLabelParam) SetDescription(v string) *MLLabelParam {
	s.Description = &v
	return s
}

func (s *MLLabelParam) SetLabelId(v string) *MLLabelParam {
	s.LabelId = &v
	return s
}

func (s *MLLabelParam) SetLastModifyTime(v int64) *MLLabelParam {
	s.LastModifyTime = &v
	return s
}

func (s *MLLabelParam) SetName(v string) *MLLabelParam {
	s.Name = &v
	return s
}

func (s *MLLabelParam) SetSettings(v []*MLLabelParamSettings) *MLLabelParam {
	s.Settings = v
	return s
}

func (s *MLLabelParam) SetType(v string) *MLLabelParam {
	s.Type = &v
	return s
}

type MLLabelParamSettings struct {
	// example:
	//
	// ""
	Config *string `json:"config,omitempty" xml:"config,omitempty"`
	// example:
	//
	// builtin
	Mode *string `json:"mode,omitempty" xml:"mode,omitempty"`
	// example:
	//
	// Trace.RCA
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// example:
	//
	// 0.01
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s MLLabelParamSettings) String() string {
	return tea.Prettify(s)
}

func (s MLLabelParamSettings) GoString() string {
	return s.String()
}

func (s *MLLabelParamSettings) SetConfig(v string) *MLLabelParamSettings {
	s.Config = &v
	return s
}

func (s *MLLabelParamSettings) SetMode(v string) *MLLabelParamSettings {
	s.Mode = &v
	return s
}

func (s *MLLabelParamSettings) SetType(v string) *MLLabelParamSettings {
	s.Type = &v
	return s
}

func (s *MLLabelParamSettings) SetVersion(v string) *MLLabelParamSettings {
	s.Version = &v
	return s
}

type MLServiceAnalysisParam struct {
	Input     []map[string]*string `json:"input,omitempty" xml:"input,omitempty" type:"Repeated"`
	Parameter map[string]*string   `json:"parameter,omitempty" xml:"parameter,omitempty"`
}

func (s MLServiceAnalysisParam) String() string {
	return tea.Prettify(s)
}

func (s MLServiceAnalysisParam) GoString() string {
	return s.String()
}

func (s *MLServiceAnalysisParam) SetInput(v []map[string]*string) *MLServiceAnalysisParam {
	s.Input = v
	return s
}

func (s *MLServiceAnalysisParam) SetParameter(v map[string]*string) *MLServiceAnalysisParam {
	s.Parameter = v
	return s
}

type MLServiceParam struct {
	// example:
	//
	// 某某服务
	Description *string              `json:"description,omitempty" xml:"description,omitempty"`
	Model       *MLServiceParamModel `json:"model,omitempty" xml:"model,omitempty" type:"Struct"`
	// example:
	//
	// service_name
	Name     *string                 `json:"name,omitempty" xml:"name,omitempty"`
	Resource *MLServiceParamResource `json:"resource,omitempty" xml:"resource,omitempty" type:"Struct"`
	// example:
	//
	// sls_builtin
	ServiceType *string `json:"serviceType,omitempty" xml:"serviceType,omitempty"`
	// example:
	//
	// running
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// 1695090077
	UpdateTimestamp *int64 `json:"updateTimestamp,omitempty" xml:"updateTimestamp,omitempty"`
}

func (s MLServiceParam) String() string {
	return tea.Prettify(s)
}

func (s MLServiceParam) GoString() string {
	return s.String()
}

func (s *MLServiceParam) SetDescription(v string) *MLServiceParam {
	s.Description = &v
	return s
}

func (s *MLServiceParam) SetModel(v *MLServiceParamModel) *MLServiceParam {
	s.Model = v
	return s
}

func (s *MLServiceParam) SetName(v string) *MLServiceParam {
	s.Name = &v
	return s
}

func (s *MLServiceParam) SetResource(v *MLServiceParamResource) *MLServiceParam {
	s.Resource = v
	return s
}

func (s *MLServiceParam) SetServiceType(v string) *MLServiceParam {
	s.ServiceType = &v
	return s
}

func (s *MLServiceParam) SetStatus(v string) *MLServiceParam {
	s.Status = &v
	return s
}

func (s *MLServiceParam) SetUpdateTimestamp(v int64) *MLServiceParam {
	s.UpdateTimestamp = &v
	return s
}

type MLServiceParamModel struct {
	// example:
	//
	// xxxx
	ModelResourceId *string `json:"modelResourceId,omitempty" xml:"modelResourceId,omitempty"`
	// example:
	//
	// xxx_type
	ModelResourceType *string `json:"modelResourceType,omitempty" xml:"modelResourceType,omitempty"`
}

func (s MLServiceParamModel) String() string {
	return tea.Prettify(s)
}

func (s MLServiceParamModel) GoString() string {
	return s.String()
}

func (s *MLServiceParamModel) SetModelResourceId(v string) *MLServiceParamModel {
	s.ModelResourceId = &v
	return s
}

func (s *MLServiceParamModel) SetModelResourceType(v string) *MLServiceParamModel {
	s.ModelResourceType = &v
	return s
}

type MLServiceParamResource struct {
	// example:
	//
	// 2
	CpuLimit *int32 `json:"cpuLimit,omitempty" xml:"cpuLimit,omitempty"`
	// example:
	//
	// 20
	Gpu *int32 `json:"gpu,omitempty" xml:"gpu,omitempty"`
	// example:
	//
	// 64
	MemoryLimit *int32 `json:"memoryLimit,omitempty" xml:"memoryLimit,omitempty"`
	// example:
	//
	// 2
	Replica *int32 `json:"replica,omitempty" xml:"replica,omitempty"`
}

func (s MLServiceParamResource) String() string {
	return tea.Prettify(s)
}

func (s MLServiceParamResource) GoString() string {
	return s.String()
}

func (s *MLServiceParamResource) SetCpuLimit(v int32) *MLServiceParamResource {
	s.CpuLimit = &v
	return s
}

func (s *MLServiceParamResource) SetGpu(v int32) *MLServiceParamResource {
	s.Gpu = &v
	return s
}

func (s *MLServiceParamResource) SetMemoryLimit(v int32) *MLServiceParamResource {
	s.MemoryLimit = &v
	return s
}

func (s *MLServiceParamResource) SetReplica(v int32) *MLServiceParamResource {
	s.Replica = &v
	return s
}

type MaxComputeExport struct {
	// This parameter is required.
	Configuration *MaxComputeExportConfiguration `json:"configuration,omitempty" xml:"configuration,omitempty"`
	CreateTime    *int64                         `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// MaxComputeExport
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// MaxComputeExport
	DisplayName      *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	LastModifiedTime *int64  `json:"lastModifiedTime,omitempty" xml:"lastModifiedTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// MaxComputeExport
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// RUNNING
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s MaxComputeExport) String() string {
	return tea.Prettify(s)
}

func (s MaxComputeExport) GoString() string {
	return s.String()
}

func (s *MaxComputeExport) SetConfiguration(v *MaxComputeExportConfiguration) *MaxComputeExport {
	s.Configuration = v
	return s
}

func (s *MaxComputeExport) SetCreateTime(v int64) *MaxComputeExport {
	s.CreateTime = &v
	return s
}

func (s *MaxComputeExport) SetDescription(v string) *MaxComputeExport {
	s.Description = &v
	return s
}

func (s *MaxComputeExport) SetDisplayName(v string) *MaxComputeExport {
	s.DisplayName = &v
	return s
}

func (s *MaxComputeExport) SetLastModifiedTime(v int64) *MaxComputeExport {
	s.LastModifiedTime = &v
	return s
}

func (s *MaxComputeExport) SetName(v string) *MaxComputeExport {
	s.Name = &v
	return s
}

func (s *MaxComputeExport) SetStatus(v string) *MaxComputeExport {
	s.Status = &v
	return s
}

type MaxComputeExportConfiguration struct {
	// This parameter is required.
	FromTime *int64 `json:"fromTime,omitempty" xml:"fromTime,omitempty"`
	// This parameter is required.
	Logstore *string `json:"logstore,omitempty" xml:"logstore,omitempty"`
	// This parameter is required.
	RoleArn *string `json:"roleArn,omitempty" xml:"roleArn,omitempty"`
	// This parameter is required.
	Sink *MaxComputeExportConfigurationSink `json:"sink,omitempty" xml:"sink,omitempty"`
	// This parameter is required.
	ToTime *int64 `json:"toTime,omitempty" xml:"toTime,omitempty"`
}

func (s MaxComputeExportConfiguration) String() string {
	return tea.Prettify(s)
}

func (s MaxComputeExportConfiguration) GoString() string {
	return s.String()
}

func (s *MaxComputeExportConfiguration) SetFromTime(v int64) *MaxComputeExportConfiguration {
	s.FromTime = &v
	return s
}

func (s *MaxComputeExportConfiguration) SetLogstore(v string) *MaxComputeExportConfiguration {
	s.Logstore = &v
	return s
}

func (s *MaxComputeExportConfiguration) SetRoleArn(v string) *MaxComputeExportConfiguration {
	s.RoleArn = &v
	return s
}

func (s *MaxComputeExportConfiguration) SetSink(v *MaxComputeExportConfigurationSink) *MaxComputeExportConfiguration {
	s.Sink = v
	return s
}

func (s *MaxComputeExportConfiguration) SetToTime(v int64) *MaxComputeExportConfiguration {
	s.ToTime = &v
	return s
}

type MaxComputeExportConfigurationSink struct {
	// This parameter is required.
	Fields []*string `json:"fields,omitempty" xml:"fields,omitempty" type:"Repeated"`
	// Deprecated
	OdpsAccessKeyId *string `json:"odpsAccessKeyId,omitempty" xml:"odpsAccessKeyId,omitempty"`
	// Deprecated
	OdpsAccessSecret *string `json:"odpsAccessSecret,omitempty" xml:"odpsAccessSecret,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// http://xxxxxxxx
	OdpsEndpoint *string `json:"odpsEndpoint,omitempty" xml:"odpsEndpoint,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// demo_project
	OdpsProject *string `json:"odpsProject,omitempty" xml:"odpsProject,omitempty"`
	// example:
	//
	// acs:ram::xxxxxxx
	OdpsRolearn *string `json:"odpsRolearn,omitempty" xml:"odpsRolearn,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// demo_table
	OdpsTable *string `json:"odpsTable,omitempty" xml:"odpsTable,omitempty"`
	// This parameter is required.
	OdpsTunnelEndpoint *string `json:"odpsTunnelEndpoint,omitempty" xml:"odpsTunnelEndpoint,omitempty"`
	// This parameter is required.
	PartitionColumn []*string `json:"partitionColumn,omitempty" xml:"partitionColumn,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// %Y_%m_%d
	PartitionTimeFormat *string `json:"partitionTimeFormat,omitempty" xml:"partitionTimeFormat,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// +0800
	TimeZone *string `json:"timeZone,omitempty" xml:"timeZone,omitempty"`
}

func (s MaxComputeExportConfigurationSink) String() string {
	return tea.Prettify(s)
}

func (s MaxComputeExportConfigurationSink) GoString() string {
	return s.String()
}

func (s *MaxComputeExportConfigurationSink) SetFields(v []*string) *MaxComputeExportConfigurationSink {
	s.Fields = v
	return s
}

func (s *MaxComputeExportConfigurationSink) SetOdpsAccessKeyId(v string) *MaxComputeExportConfigurationSink {
	s.OdpsAccessKeyId = &v
	return s
}

func (s *MaxComputeExportConfigurationSink) SetOdpsAccessSecret(v string) *MaxComputeExportConfigurationSink {
	s.OdpsAccessSecret = &v
	return s
}

func (s *MaxComputeExportConfigurationSink) SetOdpsEndpoint(v string) *MaxComputeExportConfigurationSink {
	s.OdpsEndpoint = &v
	return s
}

func (s *MaxComputeExportConfigurationSink) SetOdpsProject(v string) *MaxComputeExportConfigurationSink {
	s.OdpsProject = &v
	return s
}

func (s *MaxComputeExportConfigurationSink) SetOdpsRolearn(v string) *MaxComputeExportConfigurationSink {
	s.OdpsRolearn = &v
	return s
}

func (s *MaxComputeExportConfigurationSink) SetOdpsTable(v string) *MaxComputeExportConfigurationSink {
	s.OdpsTable = &v
	return s
}

func (s *MaxComputeExportConfigurationSink) SetOdpsTunnelEndpoint(v string) *MaxComputeExportConfigurationSink {
	s.OdpsTunnelEndpoint = &v
	return s
}

func (s *MaxComputeExportConfigurationSink) SetPartitionColumn(v []*string) *MaxComputeExportConfigurationSink {
	s.PartitionColumn = v
	return s
}

func (s *MaxComputeExportConfigurationSink) SetPartitionTimeFormat(v string) *MaxComputeExportConfigurationSink {
	s.PartitionTimeFormat = &v
	return s
}

func (s *MaxComputeExportConfigurationSink) SetTimeZone(v string) *MaxComputeExportConfigurationSink {
	s.TimeZone = &v
	return s
}

type OSSExport struct {
	Configuration *OSSExportConfiguration `json:"configuration,omitempty" xml:"configuration,omitempty"`
	// example:
	//
	// 1714284025
	CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// job-test
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// job-demo
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	// example:
	//
	// 1714284115
	LastModifiedTime *int64 `json:"lastModifiedTime,omitempty" xml:"lastModifiedTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// job-1714109458-123456
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// RUNNING
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s OSSExport) String() string {
	return tea.Prettify(s)
}

func (s OSSExport) GoString() string {
	return s.String()
}

func (s *OSSExport) SetConfiguration(v *OSSExportConfiguration) *OSSExport {
	s.Configuration = v
	return s
}

func (s *OSSExport) SetCreateTime(v int64) *OSSExport {
	s.CreateTime = &v
	return s
}

func (s *OSSExport) SetDescription(v string) *OSSExport {
	s.Description = &v
	return s
}

func (s *OSSExport) SetDisplayName(v string) *OSSExport {
	s.DisplayName = &v
	return s
}

func (s *OSSExport) SetLastModifiedTime(v int64) *OSSExport {
	s.LastModifiedTime = &v
	return s
}

func (s *OSSExport) SetName(v string) *OSSExport {
	s.Name = &v
	return s
}

func (s *OSSExport) SetStatus(v string) *OSSExport {
	s.Status = &v
	return s
}

type OSSExportConfiguration struct {
	// example:
	//
	// 1714123644
	FromTime *int64 `json:"fromTime,omitempty" xml:"fromTime,omitempty"`
	// example:
	//
	// logstore-demo
	Logstore *string `json:"logstore,omitempty" xml:"logstore,omitempty"`
	// example:
	//
	// acs:ram::123456789:role/aliyunlogdefaultrole
	RoleArn *string                     `json:"roleArn,omitempty" xml:"roleArn,omitempty"`
	Sink    *OSSExportConfigurationSink `json:"sink,omitempty" xml:"sink,omitempty" type:"Struct"`
	// example:
	//
	// 1714357112
	ToTime *int64 `json:"toTime,omitempty" xml:"toTime,omitempty"`
}

func (s OSSExportConfiguration) String() string {
	return tea.Prettify(s)
}

func (s OSSExportConfiguration) GoString() string {
	return s.String()
}

func (s *OSSExportConfiguration) SetFromTime(v int64) *OSSExportConfiguration {
	s.FromTime = &v
	return s
}

func (s *OSSExportConfiguration) SetLogstore(v string) *OSSExportConfiguration {
	s.Logstore = &v
	return s
}

func (s *OSSExportConfiguration) SetRoleArn(v string) *OSSExportConfiguration {
	s.RoleArn = &v
	return s
}

func (s *OSSExportConfiguration) SetSink(v *OSSExportConfigurationSink) *OSSExportConfiguration {
	s.Sink = v
	return s
}

func (s *OSSExportConfiguration) SetToTime(v int64) *OSSExportConfiguration {
	s.ToTime = &v
	return s
}

type OSSExportConfigurationSink struct {
	// This parameter is required.
	//
	// example:
	//
	// test-bucket
	Bucket *string `json:"bucket,omitempty" xml:"bucket,omitempty"`
	// example:
	//
	// 300
	BufferInterval *int64 `json:"bufferInterval,omitempty" xml:"bufferInterval,omitempty"`
	// example:
	//
	// 256
	BufferSize *int64 `json:"bufferSize,omitempty" xml:"bufferSize,omitempty"`
	// example:
	//
	// snappy
	CompressionType *string                `json:"compressionType,omitempty" xml:"compressionType,omitempty"`
	ContentDetail   map[string]interface{} `json:"contentDetail,omitempty" xml:"contentDetail,omitempty"`
	// example:
	//
	// json
	ContentType *string `json:"contentType,omitempty" xml:"contentType,omitempty"`
	// Deprecated
	//
	// example:
	//
	// 900
	DelaySec *int64 `json:"delaySec,omitempty" xml:"delaySec,omitempty"`
	// example:
	//
	// 900
	DelaySeconds *int64 `json:"delaySeconds,omitempty" xml:"delaySeconds,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// http://oss-cn-hangzhou-internal.aliyuncs.com
	Endpoint *string `json:"endpoint,omitempty" xml:"endpoint,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// %Y/%m/%d/%H/%M
	PathFormat *string `json:"pathFormat,omitempty" xml:"pathFormat,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// time
	PathFormatType *string `json:"pathFormatType,omitempty" xml:"pathFormatType,omitempty"`
	// example:
	//
	// demo/
	Prefix *string `json:"prefix,omitempty" xml:"prefix,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// acs:ram::123456789:role/aliyunlogdefaultrole
	RoleArn *string `json:"roleArn,omitempty" xml:"roleArn,omitempty"`
	// example:
	//
	// .json
	Suffix *string `json:"suffix,omitempty" xml:"suffix,omitempty"`
	// example:
	//
	// +0800
	TimeZone *string `json:"timeZone,omitempty" xml:"timeZone,omitempty"`
}

func (s OSSExportConfigurationSink) String() string {
	return tea.Prettify(s)
}

func (s OSSExportConfigurationSink) GoString() string {
	return s.String()
}

func (s *OSSExportConfigurationSink) SetBucket(v string) *OSSExportConfigurationSink {
	s.Bucket = &v
	return s
}

func (s *OSSExportConfigurationSink) SetBufferInterval(v int64) *OSSExportConfigurationSink {
	s.BufferInterval = &v
	return s
}

func (s *OSSExportConfigurationSink) SetBufferSize(v int64) *OSSExportConfigurationSink {
	s.BufferSize = &v
	return s
}

func (s *OSSExportConfigurationSink) SetCompressionType(v string) *OSSExportConfigurationSink {
	s.CompressionType = &v
	return s
}

func (s *OSSExportConfigurationSink) SetContentDetail(v map[string]interface{}) *OSSExportConfigurationSink {
	s.ContentDetail = v
	return s
}

func (s *OSSExportConfigurationSink) SetContentType(v string) *OSSExportConfigurationSink {
	s.ContentType = &v
	return s
}

func (s *OSSExportConfigurationSink) SetDelaySec(v int64) *OSSExportConfigurationSink {
	s.DelaySec = &v
	return s
}

func (s *OSSExportConfigurationSink) SetDelaySeconds(v int64) *OSSExportConfigurationSink {
	s.DelaySeconds = &v
	return s
}

func (s *OSSExportConfigurationSink) SetEndpoint(v string) *OSSExportConfigurationSink {
	s.Endpoint = &v
	return s
}

func (s *OSSExportConfigurationSink) SetPathFormat(v string) *OSSExportConfigurationSink {
	s.PathFormat = &v
	return s
}

func (s *OSSExportConfigurationSink) SetPathFormatType(v string) *OSSExportConfigurationSink {
	s.PathFormatType = &v
	return s
}

func (s *OSSExportConfigurationSink) SetPrefix(v string) *OSSExportConfigurationSink {
	s.Prefix = &v
	return s
}

func (s *OSSExportConfigurationSink) SetRoleArn(v string) *OSSExportConfigurationSink {
	s.RoleArn = &v
	return s
}

func (s *OSSExportConfigurationSink) SetSuffix(v string) *OSSExportConfigurationSink {
	s.Suffix = &v
	return s
}

func (s *OSSExportConfigurationSink) SetTimeZone(v string) *OSSExportConfigurationSink {
	s.TimeZone = &v
	return s
}

type OSSIngestion struct {
	// This parameter is required.
	Configuration *OSSIngestionConfiguration `json:"configuration,omitempty" xml:"configuration,omitempty"`
	// example:
	//
	// 1714360481
	CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// oss ingestion
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// oss ingestion
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	// example:
	//
	// 1714360481
	LastModifiedTime *int64 `json:"lastModifiedTime,omitempty" xml:"lastModifiedTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ingest-oss-123456
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// This parameter is required.
	Schedule *Schedule `json:"schedule,omitempty" xml:"schedule,omitempty"`
	// example:
	//
	// RUNNING
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s OSSIngestion) String() string {
	return tea.Prettify(s)
}

func (s OSSIngestion) GoString() string {
	return s.String()
}

func (s *OSSIngestion) SetConfiguration(v *OSSIngestionConfiguration) *OSSIngestion {
	s.Configuration = v
	return s
}

func (s *OSSIngestion) SetCreateTime(v int64) *OSSIngestion {
	s.CreateTime = &v
	return s
}

func (s *OSSIngestion) SetDescription(v string) *OSSIngestion {
	s.Description = &v
	return s
}

func (s *OSSIngestion) SetDisplayName(v string) *OSSIngestion {
	s.DisplayName = &v
	return s
}

func (s *OSSIngestion) SetLastModifiedTime(v int64) *OSSIngestion {
	s.LastModifiedTime = &v
	return s
}

func (s *OSSIngestion) SetName(v string) *OSSIngestion {
	s.Name = &v
	return s
}

func (s *OSSIngestion) SetSchedule(v *Schedule) *OSSIngestion {
	s.Schedule = v
	return s
}

func (s *OSSIngestion) SetStatus(v string) *OSSIngestion {
	s.Status = &v
	return s
}

type OSSIngestionConfiguration struct {
	// This parameter is required.
	//
	// example:
	//
	// myLogstore
	Logstore *string `json:"logstore,omitempty" xml:"logstore,omitempty"`
	// This parameter is required.
	Source *OSSIngestionConfigurationSource `json:"source,omitempty" xml:"source,omitempty"`
}

func (s OSSIngestionConfiguration) String() string {
	return tea.Prettify(s)
}

func (s OSSIngestionConfiguration) GoString() string {
	return s.String()
}

func (s *OSSIngestionConfiguration) SetLogstore(v string) *OSSIngestionConfiguration {
	s.Logstore = &v
	return s
}

func (s *OSSIngestionConfiguration) SetSource(v *OSSIngestionConfigurationSource) *OSSIngestionConfiguration {
	s.Source = v
	return s
}

type OSSIngestionConfigurationSource struct {
	// This parameter is required.
	//
	// example:
	//
	// ossbucket
	Bucket *string `json:"bucket,omitempty" xml:"bucket,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// none
	CompressionCodec *string `json:"compressionCodec,omitempty" xml:"compressionCodec,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// UTF-8
	Encoding *string `json:"encoding,omitempty" xml:"encoding,omitempty"`
	// example:
	//
	// 1714360481
	EndTime *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// oss-cn-hangzhou.aliyuncs.com
	Endpoint *string `json:"endpoint,omitempty" xml:"endpoint,omitempty"`
	// This parameter is required.
	Format map[string]interface{} `json:"format,omitempty" xml:"format,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// never
	Interval *string `json:"interval,omitempty" xml:"interval,omitempty"`
	// example:
	//
	// .*
	Pattern *string `json:"pattern,omitempty" xml:"pattern,omitempty"`
	// example:
	//
	// prefix
	Prefix *string `json:"prefix,omitempty" xml:"prefix,omitempty"`
	// example:
	//
	// true
	RestoreObjectEnabled *bool `json:"restoreObjectEnabled,omitempty" xml:"restoreObjectEnabled,omitempty"`
	// example:
	//
	// acs:ram::12345:role/aliyunlogdefaultrole
	RoleARN *string `json:"roleARN,omitempty" xml:"roleARN,omitempty"`
	// example:
	//
	// 1714274081
	StartTime *int64 `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// example:
	//
	// __time__
	TimeField *string `json:"timeField,omitempty" xml:"timeField,omitempty"`
	// example:
	//
	// yyyy-MM-dd HH:mm:ss
	TimeFormat *string `json:"timeFormat,omitempty" xml:"timeFormat,omitempty"`
	// example:
	//
	// [0-9]{0,2}\/[0-9a-zA-Z]+\/[0-9:,]+
	TimePattern *string `json:"timePattern,omitempty" xml:"timePattern,omitempty"`
	// example:
	//
	// GMT+08:00
	TimeZone *string `json:"timeZone,omitempty" xml:"timeZone,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// false
	UseMetaIndex *bool `json:"useMetaIndex,omitempty" xml:"useMetaIndex,omitempty"`
}

func (s OSSIngestionConfigurationSource) String() string {
	return tea.Prettify(s)
}

func (s OSSIngestionConfigurationSource) GoString() string {
	return s.String()
}

func (s *OSSIngestionConfigurationSource) SetBucket(v string) *OSSIngestionConfigurationSource {
	s.Bucket = &v
	return s
}

func (s *OSSIngestionConfigurationSource) SetCompressionCodec(v string) *OSSIngestionConfigurationSource {
	s.CompressionCodec = &v
	return s
}

func (s *OSSIngestionConfigurationSource) SetEncoding(v string) *OSSIngestionConfigurationSource {
	s.Encoding = &v
	return s
}

func (s *OSSIngestionConfigurationSource) SetEndTime(v int64) *OSSIngestionConfigurationSource {
	s.EndTime = &v
	return s
}

func (s *OSSIngestionConfigurationSource) SetEndpoint(v string) *OSSIngestionConfigurationSource {
	s.Endpoint = &v
	return s
}

func (s *OSSIngestionConfigurationSource) SetFormat(v map[string]interface{}) *OSSIngestionConfigurationSource {
	s.Format = v
	return s
}

func (s *OSSIngestionConfigurationSource) SetInterval(v string) *OSSIngestionConfigurationSource {
	s.Interval = &v
	return s
}

func (s *OSSIngestionConfigurationSource) SetPattern(v string) *OSSIngestionConfigurationSource {
	s.Pattern = &v
	return s
}

func (s *OSSIngestionConfigurationSource) SetPrefix(v string) *OSSIngestionConfigurationSource {
	s.Prefix = &v
	return s
}

func (s *OSSIngestionConfigurationSource) SetRestoreObjectEnabled(v bool) *OSSIngestionConfigurationSource {
	s.RestoreObjectEnabled = &v
	return s
}

func (s *OSSIngestionConfigurationSource) SetRoleARN(v string) *OSSIngestionConfigurationSource {
	s.RoleARN = &v
	return s
}

func (s *OSSIngestionConfigurationSource) SetStartTime(v int64) *OSSIngestionConfigurationSource {
	s.StartTime = &v
	return s
}

func (s *OSSIngestionConfigurationSource) SetTimeField(v string) *OSSIngestionConfigurationSource {
	s.TimeField = &v
	return s
}

func (s *OSSIngestionConfigurationSource) SetTimeFormat(v string) *OSSIngestionConfigurationSource {
	s.TimeFormat = &v
	return s
}

func (s *OSSIngestionConfigurationSource) SetTimePattern(v string) *OSSIngestionConfigurationSource {
	s.TimePattern = &v
	return s
}

func (s *OSSIngestionConfigurationSource) SetTimeZone(v string) *OSSIngestionConfigurationSource {
	s.TimeZone = &v
	return s
}

func (s *OSSIngestionConfigurationSource) SetUseMetaIndex(v bool) *OSSIngestionConfigurationSource {
	s.UseMetaIndex = &v
	return s
}

type PolicyConfiguration struct {
	// example:
	//
	// example_action_policy
	ActionPolicyId *string `json:"actionPolicyId,omitempty" xml:"actionPolicyId,omitempty"`
	// example:
	//
	// sls.builtin.dynamic
	AlertPolicyId *string `json:"alertPolicyId,omitempty" xml:"alertPolicyId,omitempty"`
	// example:
	//
	// 10m
	RepeatInterval *string `json:"repeatInterval,omitempty" xml:"repeatInterval,omitempty"`
}

func (s PolicyConfiguration) String() string {
	return tea.Prettify(s)
}

func (s PolicyConfiguration) GoString() string {
	return s.String()
}

func (s *PolicyConfiguration) SetActionPolicyId(v string) *PolicyConfiguration {
	s.ActionPolicyId = &v
	return s
}

func (s *PolicyConfiguration) SetAlertPolicyId(v string) *PolicyConfiguration {
	s.AlertPolicyId = &v
	return s
}

func (s *PolicyConfiguration) SetRepeatInterval(v string) *PolicyConfiguration {
	s.RepeatInterval = &v
	return s
}

type SavedSearch struct {
	// This parameter is required.
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	// This parameter is required.
	Logstore *string `json:"logstore,omitempty" xml:"logstore,omitempty"`
	// This parameter is required.
	SavedsearchName *string `json:"savedsearchName,omitempty" xml:"savedsearchName,omitempty"`
	// This parameter is required.
	SearchQuery *string `json:"searchQuery,omitempty" xml:"searchQuery,omitempty"`
	Topic       *string `json:"topic,omitempty" xml:"topic,omitempty"`
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

type Schedule struct {
	// example:
	//
	// 0/5 	- 	- 	- *
	CronExpression *string `json:"cronExpression,omitempty" xml:"cronExpression,omitempty"`
	// example:
	//
	// 4
	Delay *int32 `json:"delay,omitempty" xml:"delay,omitempty"`
	// example:
	//
	// 60s
	Interval *string `json:"interval,omitempty" xml:"interval,omitempty"`
	// example:
	//
	// false
	RunImmediately *bool `json:"runImmediately,omitempty" xml:"runImmediately,omitempty"`
	// example:
	//
	// +0800
	TimeZone *string `json:"timeZone,omitempty" xml:"timeZone,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// FixedRate
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
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

func (s *Schedule) SetDelay(v int32) *Schedule {
	s.Delay = &v
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

func (s *Schedule) SetTimeZone(v string) *Schedule {
	s.TimeZone = &v
	return s
}

func (s *Schedule) SetType(v string) *Schedule {
	s.Type = &v
	return s
}

type ScheduledSQL struct {
	// This parameter is required.
	Configuration *ScheduledSQLConfiguration `json:"configuration,omitempty" xml:"configuration,omitempty"`
	// example:
	//
	// 1714123644
	CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// schedule-sql-test
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// scheduleSqlTest
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	// example:
	//
	// 1714123644
	LastModifiedTime *int64 `json:"lastModifiedTime,omitempty" xml:"lastModifiedTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// sql-1714123463-225223
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// This parameter is required.
	Schedule *Schedule `json:"schedule,omitempty" xml:"schedule,omitempty"`
	// example:
	//
	// e73f43732852064ad5d091914e39342f
	ScheduleId *string `json:"scheduleId,omitempty" xml:"scheduleId,omitempty"`
	// example:
	//
	// ENABLED
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s ScheduledSQL) String() string {
	return tea.Prettify(s)
}

func (s ScheduledSQL) GoString() string {
	return s.String()
}

func (s *ScheduledSQL) SetConfiguration(v *ScheduledSQLConfiguration) *ScheduledSQL {
	s.Configuration = v
	return s
}

func (s *ScheduledSQL) SetCreateTime(v int64) *ScheduledSQL {
	s.CreateTime = &v
	return s
}

func (s *ScheduledSQL) SetDescription(v string) *ScheduledSQL {
	s.Description = &v
	return s
}

func (s *ScheduledSQL) SetDisplayName(v string) *ScheduledSQL {
	s.DisplayName = &v
	return s
}

func (s *ScheduledSQL) SetLastModifiedTime(v int64) *ScheduledSQL {
	s.LastModifiedTime = &v
	return s
}

func (s *ScheduledSQL) SetName(v string) *ScheduledSQL {
	s.Name = &v
	return s
}

func (s *ScheduledSQL) SetSchedule(v *Schedule) *ScheduledSQL {
	s.Schedule = v
	return s
}

func (s *ScheduledSQL) SetScheduleId(v string) *ScheduledSQL {
	s.ScheduleId = &v
	return s
}

func (s *ScheduledSQL) SetStatus(v string) *ScheduledSQL {
	s.Status = &v
	return s
}

type ScheduledSQLConfiguration struct {
	// This parameter is required.
	//
	// example:
	//
	// log2log
	DataFormat *string `json:"dataFormat,omitempty" xml:"dataFormat,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou-intranet.log.aliyuncs.com
	DestEndpoint *string `json:"destEndpoint,omitempty" xml:"destEndpoint,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// dest-logstore-demo
	DestLogstore *string `json:"destLogstore,omitempty" xml:"destLogstore,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// project-demo
	DestProject *string `json:"destProject,omitempty" xml:"destProject,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// acs:ram::123456789:role/aliyunlogetlrole
	DestRoleArn *string `json:"destRoleArn,omitempty" xml:"destRoleArn,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1712592000
	FromTime *int64 `json:"fromTime,omitempty" xml:"fromTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// @m-1m
	FromTimeExpr *string `json:"fromTimeExpr,omitempty" xml:"fromTimeExpr,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 20
	MaxRetries *int64 `json:"maxRetries,omitempty" xml:"maxRetries,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 600
	MaxRunTimeInSeconds *int64 `json:"maxRunTimeInSeconds,omitempty" xml:"maxRunTimeInSeconds,omitempty"`
	// This parameter is required.
	Parameters map[string]interface{} `json:"parameters,omitempty" xml:"parameters,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// enhanced
	ResourcePool *string `json:"resourcePool,omitempty" xml:"resourcePool,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// acs:ram::123456789:role/aliyunlogetlrole
	RoleArn *string `json:"roleArn,omitempty" xml:"roleArn,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 	- | select *
	Script *string `json:"script,omitempty" xml:"script,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// source-logstore-demo
	SourceLogstore *string `json:"sourceLogstore,omitempty" xml:"sourceLogstore,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// searchQuery
	SqlType *string `json:"sqlType,omitempty" xml:"sqlType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0
	ToTime *int64 `json:"toTime,omitempty" xml:"toTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// @m
	ToTimeExpr *string `json:"toTimeExpr,omitempty" xml:"toTimeExpr,omitempty"`
}

func (s ScheduledSQLConfiguration) String() string {
	return tea.Prettify(s)
}

func (s ScheduledSQLConfiguration) GoString() string {
	return s.String()
}

func (s *ScheduledSQLConfiguration) SetDataFormat(v string) *ScheduledSQLConfiguration {
	s.DataFormat = &v
	return s
}

func (s *ScheduledSQLConfiguration) SetDestEndpoint(v string) *ScheduledSQLConfiguration {
	s.DestEndpoint = &v
	return s
}

func (s *ScheduledSQLConfiguration) SetDestLogstore(v string) *ScheduledSQLConfiguration {
	s.DestLogstore = &v
	return s
}

func (s *ScheduledSQLConfiguration) SetDestProject(v string) *ScheduledSQLConfiguration {
	s.DestProject = &v
	return s
}

func (s *ScheduledSQLConfiguration) SetDestRoleArn(v string) *ScheduledSQLConfiguration {
	s.DestRoleArn = &v
	return s
}

func (s *ScheduledSQLConfiguration) SetFromTime(v int64) *ScheduledSQLConfiguration {
	s.FromTime = &v
	return s
}

func (s *ScheduledSQLConfiguration) SetFromTimeExpr(v string) *ScheduledSQLConfiguration {
	s.FromTimeExpr = &v
	return s
}

func (s *ScheduledSQLConfiguration) SetMaxRetries(v int64) *ScheduledSQLConfiguration {
	s.MaxRetries = &v
	return s
}

func (s *ScheduledSQLConfiguration) SetMaxRunTimeInSeconds(v int64) *ScheduledSQLConfiguration {
	s.MaxRunTimeInSeconds = &v
	return s
}

func (s *ScheduledSQLConfiguration) SetParameters(v map[string]interface{}) *ScheduledSQLConfiguration {
	s.Parameters = v
	return s
}

func (s *ScheduledSQLConfiguration) SetResourcePool(v string) *ScheduledSQLConfiguration {
	s.ResourcePool = &v
	return s
}

func (s *ScheduledSQLConfiguration) SetRoleArn(v string) *ScheduledSQLConfiguration {
	s.RoleArn = &v
	return s
}

func (s *ScheduledSQLConfiguration) SetScript(v string) *ScheduledSQLConfiguration {
	s.Script = &v
	return s
}

func (s *ScheduledSQLConfiguration) SetSourceLogstore(v string) *ScheduledSQLConfiguration {
	s.SourceLogstore = &v
	return s
}

func (s *ScheduledSQLConfiguration) SetSqlType(v string) *ScheduledSQLConfiguration {
	s.SqlType = &v
	return s
}

func (s *ScheduledSQLConfiguration) SetToTime(v int64) *ScheduledSQLConfiguration {
	s.ToTime = &v
	return s
}

func (s *ScheduledSQLConfiguration) SetToTimeExpr(v string) *ScheduledSQLConfiguration {
	s.ToTimeExpr = &v
	return s
}

type SeverityConfiguration struct {
	EvalCondition *ConditionConfiguration `json:"evalCondition,omitempty" xml:"evalCondition,omitempty"`
	// example:
	//
	// 8
	Severity *int32 `json:"severity,omitempty" xml:"severity,omitempty"`
}

func (s SeverityConfiguration) String() string {
	return tea.Prettify(s)
}

func (s SeverityConfiguration) GoString() string {
	return s.String()
}

func (s *SeverityConfiguration) SetEvalCondition(v *ConditionConfiguration) *SeverityConfiguration {
	s.EvalCondition = v
	return s
}

func (s *SeverityConfiguration) SetSeverity(v int32) *SeverityConfiguration {
	s.Severity = &v
	return s
}

type SinkAlerthubConfiguration struct {
	// example:
	//
	// true
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
}

func (s SinkAlerthubConfiguration) String() string {
	return tea.Prettify(s)
}

func (s SinkAlerthubConfiguration) GoString() string {
	return s.String()
}

func (s *SinkAlerthubConfiguration) SetEnabled(v bool) *SinkAlerthubConfiguration {
	s.Enabled = &v
	return s
}

type SinkCmsConfiguration struct {
	// example:
	//
	// true
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
}

func (s SinkCmsConfiguration) String() string {
	return tea.Prettify(s)
}

func (s SinkCmsConfiguration) GoString() string {
	return s.String()
}

func (s *SinkCmsConfiguration) SetEnabled(v bool) *SinkCmsConfiguration {
	s.Enabled = &v
	return s
}

type SinkEventStoreConfiguration struct {
	// example:
	//
	// true
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// example:
	//
	// cn-shanghai-intranet.log.aliyuncs.com
	Endpoint *string `json:"endpoint,omitempty" xml:"endpoint,omitempty"`
	// example:
	//
	// exampleStore
	EventStore *string `json:"eventStore,omitempty" xml:"eventStore,omitempty"`
	// example:
	//
	// exampleProject
	Project *string `json:"project,omitempty" xml:"project,omitempty"`
	// example:
	//
	// acs:ram::123456789:role/aliyunlogetlrole
	RoleArn *string `json:"roleArn,omitempty" xml:"roleArn,omitempty"`
}

func (s SinkEventStoreConfiguration) String() string {
	return tea.Prettify(s)
}

func (s SinkEventStoreConfiguration) GoString() string {
	return s.String()
}

func (s *SinkEventStoreConfiguration) SetEnabled(v bool) *SinkEventStoreConfiguration {
	s.Enabled = &v
	return s
}

func (s *SinkEventStoreConfiguration) SetEndpoint(v string) *SinkEventStoreConfiguration {
	s.Endpoint = &v
	return s
}

func (s *SinkEventStoreConfiguration) SetEventStore(v string) *SinkEventStoreConfiguration {
	s.EventStore = &v
	return s
}

func (s *SinkEventStoreConfiguration) SetProject(v string) *SinkEventStoreConfiguration {
	s.Project = &v
	return s
}

func (s *SinkEventStoreConfiguration) SetRoleArn(v string) *SinkEventStoreConfiguration {
	s.RoleArn = &v
	return s
}

type StoreViewStore struct {
	// This parameter is required.
	Project *string `json:"project,omitempty" xml:"project,omitempty"`
	Query   *string `json:"query,omitempty" xml:"query,omitempty"`
	// This parameter is required.
	StoreName *string `json:"storeName,omitempty" xml:"storeName,omitempty"`
}

func (s StoreViewStore) String() string {
	return tea.Prettify(s)
}

func (s StoreViewStore) GoString() string {
	return s.String()
}

func (s *StoreViewStore) SetProject(v string) *StoreViewStore {
	s.Project = &v
	return s
}

func (s *StoreViewStore) SetQuery(v string) *StoreViewStore {
	s.Query = &v
	return s
}

func (s *StoreViewStore) SetStoreName(v string) *StoreViewStore {
	s.StoreName = &v
	return s
}

type TemplateConfiguration struct {
	Aonotations map[string]interface{} `json:"aonotations,omitempty" xml:"aonotations,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// sls.app.ack.ip.not_enough
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// cn
	Lang   *string                `json:"lang,omitempty" xml:"lang,omitempty"`
	Tokens map[string]interface{} `json:"tokens,omitempty" xml:"tokens,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// sys
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// example:
	//
	// 1
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s TemplateConfiguration) String() string {
	return tea.Prettify(s)
}

func (s TemplateConfiguration) GoString() string {
	return s.String()
}

func (s *TemplateConfiguration) SetAonotations(v map[string]interface{}) *TemplateConfiguration {
	s.Aonotations = v
	return s
}

func (s *TemplateConfiguration) SetId(v string) *TemplateConfiguration {
	s.Id = &v
	return s
}

func (s *TemplateConfiguration) SetLang(v string) *TemplateConfiguration {
	s.Lang = &v
	return s
}

func (s *TemplateConfiguration) SetTokens(v map[string]interface{}) *TemplateConfiguration {
	s.Tokens = v
	return s
}

func (s *TemplateConfiguration) SetType(v string) *TemplateConfiguration {
	s.Type = &v
	return s
}

func (s *TemplateConfiguration) SetVersion(v string) *TemplateConfiguration {
	s.Version = &v
	return s
}

type Ticket struct {
	// example:
	//
	// 1000000000
	CallerUid *int64 `json:"callerUid,omitempty" xml:"callerUid,omitempty"`
	// example:
	//
	// 2023-09-06 14:57:07
	CreateDate *string `json:"createDate,omitempty" xml:"createDate,omitempty"`
	// example:
	//
	// 100
	ExpirationTime *int64 `json:"expirationTime,omitempty" xml:"expirationTime,omitempty"`
	// example:
	//
	// 2023-09-06 14:58:07
	ExpireDate *string `json:"expireDate,omitempty" xml:"expireDate,omitempty"`
	// example:
	//
	// {"xx":"yy"}
	Extra *string `json:"extra,omitempty" xml:"extra,omitempty"`
	// example:
	//
	// 测试
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 1
	Number *int32 `json:"number,omitempty" xml:"number,omitempty"`
	// example:
	//
	// {"type":"aliyun","ids":[1,2]}
	SharingTo *string `json:"sharingTo,omitempty" xml:"sharingTo,omitempty"`
	// example:
	//
	// xxxxx
	Ticket *string `json:"ticket,omitempty" xml:"ticket,omitempty"`
	// example:
	//
	// xxxxx
	TicketId *string `json:"ticketId,omitempty" xml:"ticketId,omitempty"`
	// example:
	//
	// 1
	UsedNumber *int32 `json:"usedNumber,omitempty" xml:"usedNumber,omitempty"`
	// example:
	//
	// false
	Valid *bool `json:"valid,omitempty" xml:"valid,omitempty"`
}

func (s Ticket) String() string {
	return tea.Prettify(s)
}

func (s Ticket) GoString() string {
	return s.String()
}

func (s *Ticket) SetCallerUid(v int64) *Ticket {
	s.CallerUid = &v
	return s
}

func (s *Ticket) SetCreateDate(v string) *Ticket {
	s.CreateDate = &v
	return s
}

func (s *Ticket) SetExpirationTime(v int64) *Ticket {
	s.ExpirationTime = &v
	return s
}

func (s *Ticket) SetExpireDate(v string) *Ticket {
	s.ExpireDate = &v
	return s
}

func (s *Ticket) SetExtra(v string) *Ticket {
	s.Extra = &v
	return s
}

func (s *Ticket) SetName(v string) *Ticket {
	s.Name = &v
	return s
}

func (s *Ticket) SetNumber(v int32) *Ticket {
	s.Number = &v
	return s
}

func (s *Ticket) SetSharingTo(v string) *Ticket {
	s.SharingTo = &v
	return s
}

func (s *Ticket) SetTicket(v string) *Ticket {
	s.Ticket = &v
	return s
}

func (s *Ticket) SetTicketId(v string) *Ticket {
	s.TicketId = &v
	return s
}

func (s *Ticket) SetUsedNumber(v int32) *Ticket {
	s.UsedNumber = &v
	return s
}

func (s *Ticket) SetValid(v bool) *Ticket {
	s.Valid = &v
	return s
}

type Chart struct {
	// This parameter is required.
	Action map[string]interface{} `json:"action,omitempty" xml:"action,omitempty"`
	// This parameter is required.
	Display map[string]interface{} `json:"display,omitempty" xml:"display,omitempty"`
	// This parameter is required.
	Search map[string]interface{} `json:"search,omitempty" xml:"search,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// my-chart
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// linepro
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
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

type Dashboard struct {
	Attribute map[string]*string `json:"attribute,omitempty" xml:"attribute,omitempty"`
	// This parameter is required.
	Charts []*Chart `json:"charts,omitempty" xml:"charts,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// dashboard-1609294922657-434834
	DashboardName *string `json:"dashboardName,omitempty" xml:"dashboardName,omitempty"`
	// example:
	//
	// 这是一个仪表盘。
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test-alert
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
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

type EtlJob struct {
	// This parameter is required.
	//
	// example:
	//
	// true
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
	// This parameter is required.
	EtlJobName *string `json:"etlJobName,omitempty" xml:"etlJobName,omitempty"`
	// This parameter is required.
	FunctionConfig *EtlJobFunctionConfig `json:"functionConfig,omitempty" xml:"functionConfig,omitempty" type:"Struct"`
	// This parameter is required.
	FunctionParameter map[string]interface{} `json:"functionParameter,omitempty" xml:"functionParameter,omitempty"`
	// This parameter is required.
	LogConfig *EtlJobLogConfig `json:"logConfig,omitempty" xml:"logConfig,omitempty" type:"Struct"`
	// This parameter is required.
	SourceConfig *EtlJobSourceConfig `json:"sourceConfig,omitempty" xml:"sourceConfig,omitempty" type:"Struct"`
	// This parameter is required.
	TriggerConfig *EtlJobTriggerConfig `json:"triggerConfig,omitempty" xml:"triggerConfig,omitempty" type:"Struct"`
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
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
	Endpoint  *string `json:"endpoint,omitempty" xml:"endpoint,omitempty"`
	// example:
	//
	// hello-wrold
	FunctionName *string `json:"functionName,omitempty" xml:"functionName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// FunctionCompute
	FunctionProvider *string `json:"functionProvider,omitempty" xml:"functionProvider,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionName *string `json:"regionName,omitempty" xml:"regionName,omitempty"`
	RoleArn    *string `json:"roleArn,omitempty" xml:"roleArn,omitempty"`
	// example:
	//
	// my-service
	ServiceName *string `json:"serviceName,omitempty" xml:"serviceName,omitempty"`
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
	// This parameter is required.
	Endpoint *string `json:"endpoint,omitempty" xml:"endpoint,omitempty"`
	// This parameter is required.
	LogstoreName *string `json:"logstoreName,omitempty" xml:"logstoreName,omitempty"`
	// This parameter is required.
	ProjectName *string `json:"projectName,omitempty" xml:"projectName,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// my-logstore
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
	// This parameter is required.
	MaxRetryTime *int32 `json:"maxRetryTime,omitempty" xml:"maxRetryTime,omitempty"`
	// This parameter is required.
	RoleArn *string `json:"roleArn,omitempty" xml:"roleArn,omitempty"`
	// example:
	//
	// at-unixtime
	StartingPosition *string `json:"startingPosition,omitempty" xml:"startingPosition,omitempty"`
	// example:
	//
	// 当 strtingPosition 为 at-unixtime 时生效
	StartingUnixtime *int64 `json:"startingUnixtime,omitempty" xml:"startingUnixtime,omitempty"`
	// This parameter is required.
	TriggerInterval *int32 `json:"triggerInterval,omitempty" xml:"triggerInterval,omitempty"`
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

type EtlMeta struct {
	// This parameter is required.
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
	// This parameter is required.
	EtlMetaKey *string `json:"etlMetaKey,omitempty" xml:"etlMetaKey,omitempty"`
	// This parameter is required.
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

type ExternalStore struct {
	// This parameter is required.
	//
	// example:
	//
	// rds_store
	ExternalStoreName *string `json:"externalStoreName,omitempty" xml:"externalStoreName,omitempty"`
	// This parameter is required.
	Parameter map[string]interface{} `json:"parameter,omitempty" xml:"parameter,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// rds-vpc
	StoreType *string `json:"storeType,omitempty" xml:"storeType,omitempty"`
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
	Keys map[string]*IndexKeysValue `json:"keys,omitempty" xml:"keys,omitempty"`
	// example:
	//
	// 1622186280
	LastModifyTime *int64     `json:"lastModifyTime,omitempty" xml:"lastModifyTime,omitempty"`
	Line           *IndexLine `json:"line,omitempty" xml:"line,omitempty" type:"Struct"`
	// example:
	//
	// true
	LogReduce          *bool     `json:"log_reduce,omitempty" xml:"log_reduce,omitempty"`
	LogReduceBlackList []*string `json:"log_reduce_black_list,omitempty" xml:"log_reduce_black_list,omitempty" type:"Repeated"`
	LogReduceWhiteList []*string `json:"log_reduce_white_list,omitempty" xml:"log_reduce_white_list,omitempty" type:"Repeated"`
	// example:
	//
	// 2048
	MaxTextLen *int32 `json:"max_text_len,omitempty" xml:"max_text_len,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30
	Ttl *int32 `json:"ttl,omitempty" xml:"ttl,omitempty"`
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
	// example:
	//
	// true
	CaseSensitive *bool `json:"caseSensitive,omitempty" xml:"caseSensitive,omitempty"`
	// example:
	//
	// true
	Chn         *bool     `json:"chn,omitempty" xml:"chn,omitempty"`
	ExcludeKeys []*string `json:"exclude_keys,omitempty" xml:"exclude_keys,omitempty" type:"Repeated"`
	IncludeKeys []*string `json:"include_keys,omitempty" xml:"include_keys,omitempty" type:"Repeated"`
	// This parameter is required.
	Token []*string `json:"token,omitempty" xml:"token,omitempty" type:"Repeated"`
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
	// This parameter is required.
	LoggingDetails []*LoggingLoggingDetails `json:"loggingDetails,omitempty" xml:"loggingDetails,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// my-project
	LoggingProject *string `json:"loggingProject,omitempty" xml:"loggingProject,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// my-logstore
	Logstore *string `json:"logstore,omitempty" xml:"logstore,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// consumergroup_log
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
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
	// example:
	//
	// true
	AppendMeta *bool `json:"appendMeta,omitempty" xml:"appendMeta,omitempty"`
	// example:
	//
	// true
	AutoSplit  *bool  `json:"autoSplit,omitempty" xml:"autoSplit,omitempty"`
	CreateTime *int32 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// false
	EnableTracking *bool        `json:"enable_tracking,omitempty" xml:"enable_tracking,omitempty"`
	EncryptConf    *EncryptConf `json:"encrypt_conf,omitempty" xml:"encrypt_conf,omitempty"`
	HotTtl         *int32       `json:"hot_ttl,omitempty" xml:"hot_ttl,omitempty"`
	// example:
	//
	// 30
	InfrequentAccessTTL *int32 `json:"infrequentAccessTTL,omitempty" xml:"infrequentAccessTTL,omitempty"`
	LastModifyTime      *int32 `json:"lastModifyTime,omitempty" xml:"lastModifyTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// my-logstore
	LogstoreName *string `json:"logstoreName,omitempty" xml:"logstoreName,omitempty"`
	// example:
	//
	// 2
	MaxSplitShard *int32 `json:"maxSplitShard,omitempty" xml:"maxSplitShard,omitempty"`
	// example:
	//
	// standard
	Mode        *string `json:"mode,omitempty" xml:"mode,omitempty"`
	ProductType *string `json:"productType,omitempty" xml:"productType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2
	ShardCount    *int32  `json:"shardCount,omitempty" xml:"shardCount,omitempty"`
	TelemetryType *string `json:"telemetryType,omitempty" xml:"telemetryType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30
	Ttl *int32 `json:"ttl,omitempty" xml:"ttl,omitempty"`
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

func (s *Logstore) SetInfrequentAccessTTL(v int32) *Logstore {
	s.InfrequentAccessTTL = &v
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
	// example:
	//
	// 192.168.x.x
	Ip *string `json:"ip,omitempty" xml:"ip,omitempty"`
	// example:
	//
	// 1657509674
	LastHeartbeatTime *int64 `json:"lastHeartbeatTime,omitempty" xml:"lastHeartbeatTime,omitempty"`
	// example:
	//
	// 3B70F4F1-80F7-46C4-A6C1-100D66C***47
	MachineUniqueid *string `json:"machine-uniqueid,omitempty" xml:"machine-uniqueid,omitempty"`
	// example:
	//
	// test
	UserdefinedId *string `json:"userdefined-id,omitempty" xml:"userdefined-id,omitempty"`
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
	GroupAttribute *MachineGroupGroupAttribute `json:"groupAttribute,omitempty" xml:"groupAttribute,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// test-group
	GroupName *string `json:"groupName,omitempty" xml:"groupName,omitempty"`
	GroupType *string `json:"groupType,omitempty" xml:"groupType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ip
	MachineIdentifyType *string `json:"machineIdentifyType,omitempty" xml:"machineIdentifyType,omitempty"`
	// This parameter is required.
	MachineList []*string `json:"machineList,omitempty" xml:"machineList,omitempty" type:"Repeated"`
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
	// example:
	//
	// test-group
	ExternalName *string `json:"externalName,omitempty" xml:"externalName,omitempty"`
	// example:
	//
	// test-topic
	GroupTopic *string `json:"groupTopic,omitempty" xml:"groupTopic,omitempty"`
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
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// This parameter is required.
	Description    *string `json:"description,omitempty" xml:"description,omitempty"`
	LastModifyTime *string `json:"lastModifyTime,omitempty" xml:"lastModifyTime,omitempty"`
	Owner          *string `json:"owner,omitempty" xml:"owner,omitempty"`
	// This parameter is required.
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

type ServiceStatus struct {
	// example:
	//
	// false
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// example:
	//
	// NotExist
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s ServiceStatus) String() string {
	return tea.Prettify(s)
}

func (s ServiceStatus) GoString() string {
	return s.String()
}

func (s *ServiceStatus) SetEnabled(v bool) *ServiceStatus {
	s.Enabled = &v
	return s
}

func (s *ServiceStatus) SetStatus(v string) *ServiceStatus {
	s.Status = &v
	return s
}

type Shard struct {
	// example:
	//
	// 1453949705
	CreateTime *int32 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// 8000000000000000000000000000000
	ExclusiveEndKey *string `json:"exclusiveEndKey,omitempty" xml:"exclusiveEndKey,omitempty"`
	// example:
	//
	// 00000000000000000000000000000000
	InclusiveBeginKey *string `json:"inclusiveBeginKey,omitempty" xml:"inclusiveBeginKey,omitempty"`
	// example:
	//
	// 0
	ShardID *int32 `json:"shardID,omitempty" xml:"shardID,omitempty"`
	// example:
	//
	// readwrite
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
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

type MLDataParamAnnotationsValue struct {
	// example:
	//
	// xxxx
	AnnotatedBy *string `json:"annotatedBy,omitempty" xml:"annotatedBy,omitempty"`
	// example:
	//
	// 1694761550
	UpdateTime *int64               `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
	Results    []map[string]*string `json:"results,omitempty" xml:"results,omitempty" type:"Repeated"`
}

func (s MLDataParamAnnotationsValue) String() string {
	return tea.Prettify(s)
}

func (s MLDataParamAnnotationsValue) GoString() string {
	return s.String()
}

func (s *MLDataParamAnnotationsValue) SetAnnotatedBy(v string) *MLDataParamAnnotationsValue {
	s.AnnotatedBy = &v
	return s
}

func (s *MLDataParamAnnotationsValue) SetUpdateTime(v int64) *MLDataParamAnnotationsValue {
	s.UpdateTime = &v
	return s
}

func (s *MLDataParamAnnotationsValue) SetResults(v []map[string]*string) *MLDataParamAnnotationsValue {
	s.Results = v
	return s
}

type MLDataParamPredictionsValue struct {
	// example:
	//
	// xxx
	AnnotatedBy *string `json:"annotatedBy,omitempty" xml:"annotatedBy,omitempty"`
	// example:
	//
	// 1694761550
	UpdateTime *int64               `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
	Results    []map[string]*string `json:"results,omitempty" xml:"results,omitempty" type:"Repeated"`
}

func (s MLDataParamPredictionsValue) String() string {
	return tea.Prettify(s)
}

func (s MLDataParamPredictionsValue) GoString() string {
	return s.String()
}

func (s *MLDataParamPredictionsValue) SetAnnotatedBy(v string) *MLDataParamPredictionsValue {
	s.AnnotatedBy = &v
	return s
}

func (s *MLDataParamPredictionsValue) SetUpdateTime(v int64) *MLDataParamPredictionsValue {
	s.UpdateTime = &v
	return s
}

func (s *MLDataParamPredictionsValue) SetResults(v []map[string]*string) *MLDataParamPredictionsValue {
	s.Results = v
	return s
}

type IndexKeysValue struct {
	// example:
	//
	// true
	Chn *bool `json:"chn,omitempty" xml:"chn,omitempty"`
	// example:
	//
	// true
	CaseSensitive *bool     `json:"caseSensitive,omitempty" xml:"caseSensitive,omitempty"`
	Token         []*string `json:"token,omitempty" xml:"token,omitempty" type:"Repeated"`
	// example:
	//
	// myAlias
	Alias *string `json:"alias,omitempty" xml:"alias,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// text
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// example:
	//
	// true
	DocValue *bool `json:"doc_value,omitempty" xml:"doc_value,omitempty"`
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
	// Specifies whether to enable case sensitivity. This parameter is required only when **type*	- is set to **text**. Valid values:
	//
	// 	- true
	//
	// 	- false (default)
	//
	// example:
	//
	// false
	CaseSensitive *bool `json:"caseSensitive,omitempty" xml:"caseSensitive,omitempty"`
	// Specifies whether to include Chinese characters. This parameter is required only when **type*	- is set to **text**. Valid values:
	//
	// 	- true
	//
	// 	- false (default)
	//
	// example:
	//
	// false
	Chn *bool `json:"chn,omitempty" xml:"chn,omitempty"`
	// The data type of the field value. Valid values: text, json, double, and long.
	//
	// This parameter is required.
	//
	// example:
	//
	// text
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The alias of the field.
	//
	// example:
	//
	// myAlias
	Alias *string `json:"alias,omitempty" xml:"alias,omitempty"`
	// The delimiters that are used to split text.
	Token []*string `json:"token,omitempty" xml:"token,omitempty" type:"Repeated"`
	// Specifies whether to turn on Enable Analytics for the field.
	//
	// example:
	//
	// false
	DocValue *bool `json:"doc_value,omitempty" xml:"doc_value,omitempty"`
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
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
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
	// The ID of the resource group.
	//
	// This parameter is required.
	//
	// example:
	//
	// rg-aek2i7nhaxifxey
	ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	// The ID of the resource.
	//
	// This parameter is required.
	//
	// example:
	//
	// ali-test-project
	ResourceId *string `json:"resourceId,omitempty" xml:"resourceId,omitempty"`
	// The type of the resource. Only PROJECT is supported. Set the value to PROJECT.
	//
	// example:
	//
	// PROJECT
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
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
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
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

type ConsumerGroupHeartBeatRequest struct {
	// The IDs of shards whose data is being consumed.
	//
	// This parameter is required.
	Body []*int32 `json:"body,omitempty" xml:"body,omitempty" type:"Repeated"`
	// The consumer.
	//
	// This parameter is required.
	//
	// example:
	//
	// consumer_1
	Consumer *string `json:"consumer,omitempty" xml:"consumer,omitempty"`
}

func (s ConsumerGroupHeartBeatRequest) String() string {
	return tea.Prettify(s)
}

func (s ConsumerGroupHeartBeatRequest) GoString() string {
	return s.String()
}

func (s *ConsumerGroupHeartBeatRequest) SetBody(v []*int32) *ConsumerGroupHeartBeatRequest {
	s.Body = v
	return s
}

func (s *ConsumerGroupHeartBeatRequest) SetConsumer(v string) *ConsumerGroupHeartBeatRequest {
	s.Consumer = &v
	return s
}

type ConsumerGroupHeartBeatResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       []*int32           `json:"body,omitempty" xml:"body,omitempty" type:"Repeated"`
}

func (s ConsumerGroupHeartBeatResponse) String() string {
	return tea.Prettify(s)
}

func (s ConsumerGroupHeartBeatResponse) GoString() string {
	return s.String()
}

func (s *ConsumerGroupHeartBeatResponse) SetHeaders(v map[string]*string) *ConsumerGroupHeartBeatResponse {
	s.Headers = v
	return s
}

func (s *ConsumerGroupHeartBeatResponse) SetStatusCode(v int32) *ConsumerGroupHeartBeatResponse {
	s.StatusCode = &v
	return s
}

func (s *ConsumerGroupHeartBeatResponse) SetBody(v []*int32) *ConsumerGroupHeartBeatResponse {
	s.Body = v
	return s
}

type ConsumerGroupUpdateCheckPointRequest struct {
	// Shard ID。
	Body []*ConsumerGroupUpdateCheckPointRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// consumer_1
	Consumer *string `json:"consumer,omitempty" xml:"consumer,omitempty"`
	// example:
	//
	// False
	ForceSuccess *bool `json:"forceSuccess,omitempty" xml:"forceSuccess,omitempty"`
}

func (s ConsumerGroupUpdateCheckPointRequest) String() string {
	return tea.Prettify(s)
}

func (s ConsumerGroupUpdateCheckPointRequest) GoString() string {
	return s.String()
}

func (s *ConsumerGroupUpdateCheckPointRequest) SetBody(v []*ConsumerGroupUpdateCheckPointRequestBody) *ConsumerGroupUpdateCheckPointRequest {
	s.Body = v
	return s
}

func (s *ConsumerGroupUpdateCheckPointRequest) SetConsumer(v string) *ConsumerGroupUpdateCheckPointRequest {
	s.Consumer = &v
	return s
}

func (s *ConsumerGroupUpdateCheckPointRequest) SetForceSuccess(v bool) *ConsumerGroupUpdateCheckPointRequest {
	s.ForceSuccess = &v
	return s
}

type ConsumerGroupUpdateCheckPointRequestBody struct {
	// This parameter is required.
	//
	// example:
	//
	// MTUyNDE1NTM3OTM3MzkwODQ5Ng==
	Checkpoint *string `json:"checkpoint,omitempty" xml:"checkpoint,omitempty"`
	// Shard ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	Shard *int32 `json:"shard,omitempty" xml:"shard,omitempty"`
}

func (s ConsumerGroupUpdateCheckPointRequestBody) String() string {
	return tea.Prettify(s)
}

func (s ConsumerGroupUpdateCheckPointRequestBody) GoString() string {
	return s.String()
}

func (s *ConsumerGroupUpdateCheckPointRequestBody) SetCheckpoint(v string) *ConsumerGroupUpdateCheckPointRequestBody {
	s.Checkpoint = &v
	return s
}

func (s *ConsumerGroupUpdateCheckPointRequestBody) SetShard(v int32) *ConsumerGroupUpdateCheckPointRequestBody {
	s.Shard = &v
	return s
}

type ConsumerGroupUpdateCheckPointResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s ConsumerGroupUpdateCheckPointResponse) String() string {
	return tea.Prettify(s)
}

func (s ConsumerGroupUpdateCheckPointResponse) GoString() string {
	return s.String()
}

func (s *ConsumerGroupUpdateCheckPointResponse) SetHeaders(v map[string]*string) *ConsumerGroupUpdateCheckPointResponse {
	s.Headers = v
	return s
}

func (s *ConsumerGroupUpdateCheckPointResponse) SetStatusCode(v int32) *ConsumerGroupUpdateCheckPointResponse {
	s.StatusCode = &v
	return s
}

type CreateAlertRequest struct {
	// This parameter is required.
	Configuration *AlertConfiguration `json:"configuration,omitempty" xml:"configuration,omitempty"`
	Description   *string             `json:"description,omitempty" xml:"description,omitempty"`
	// This parameter is required.
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	// This parameter is required.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// This parameter is required.
	Schedule *Schedule `json:"schedule,omitempty" xml:"schedule,omitempty"`
}

func (s CreateAlertRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAlertRequest) GoString() string {
	return s.String()
}

func (s *CreateAlertRequest) SetConfiguration(v *AlertConfiguration) *CreateAlertRequest {
	s.Configuration = v
	return s
}

func (s *CreateAlertRequest) SetDescription(v string) *CreateAlertRequest {
	s.Description = &v
	return s
}

func (s *CreateAlertRequest) SetDisplayName(v string) *CreateAlertRequest {
	s.DisplayName = &v
	return s
}

func (s *CreateAlertRequest) SetName(v string) *CreateAlertRequest {
	s.Name = &v
	return s
}

func (s *CreateAlertRequest) SetSchedule(v *Schedule) *CreateAlertRequest {
	s.Schedule = v
	return s
}

type CreateAlertResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
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

type CreateAnnotationDataSetRequest struct {
	// The data structure of the request.
	Body *MLDataSetParam `json:"body,omitempty" xml:"body,omitempty"`
	// The unique identifier of the dataset.
	//
	// example:
	//
	// cb8cc4eb51a85e823471cdb368fae9be
	DatasetId *string `json:"datasetId,omitempty" xml:"datasetId,omitempty"`
}

func (s CreateAnnotationDataSetRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAnnotationDataSetRequest) GoString() string {
	return s.String()
}

func (s *CreateAnnotationDataSetRequest) SetBody(v *MLDataSetParam) *CreateAnnotationDataSetRequest {
	s.Body = v
	return s
}

func (s *CreateAnnotationDataSetRequest) SetDatasetId(v string) *CreateAnnotationDataSetRequest {
	s.DatasetId = &v
	return s
}

type CreateAnnotationDataSetResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s CreateAnnotationDataSetResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateAnnotationDataSetResponse) GoString() string {
	return s.String()
}

func (s *CreateAnnotationDataSetResponse) SetHeaders(v map[string]*string) *CreateAnnotationDataSetResponse {
	s.Headers = v
	return s
}

func (s *CreateAnnotationDataSetResponse) SetStatusCode(v int32) *CreateAnnotationDataSetResponse {
	s.StatusCode = &v
	return s
}

type CreateAnnotationLabelRequest struct {
	// The data structure of the request.
	Body *MLLabelParam `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAnnotationLabelRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAnnotationLabelRequest) GoString() string {
	return s.String()
}

func (s *CreateAnnotationLabelRequest) SetBody(v *MLLabelParam) *CreateAnnotationLabelRequest {
	s.Body = v
	return s
}

type CreateAnnotationLabelResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s CreateAnnotationLabelResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateAnnotationLabelResponse) GoString() string {
	return s.String()
}

func (s *CreateAnnotationLabelResponse) SetHeaders(v map[string]*string) *CreateAnnotationLabelResponse {
	s.Headers = v
	return s
}

func (s *CreateAnnotationLabelResponse) SetStatusCode(v int32) *CreateAnnotationLabelResponse {
	s.StatusCode = &v
	return s
}

type CreateConfigRequest struct {
	// The body of the request.
	Body *LogtailConfig `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateConfigRequest) GoString() string {
	return s.String()
}

func (s *CreateConfigRequest) SetBody(v *LogtailConfig) *CreateConfigRequest {
	s.Body = v
	return s
}

type CreateConfigResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s CreateConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateConfigResponse) GoString() string {
	return s.String()
}

func (s *CreateConfigResponse) SetHeaders(v map[string]*string) *CreateConfigResponse {
	s.Headers = v
	return s
}

func (s *CreateConfigResponse) SetStatusCode(v int32) *CreateConfigResponse {
	s.StatusCode = &v
	return s
}

type CreateConsumerGroupRequest struct {
	// The name of the consumer group. The name must be unique in a project.
	//
	// This parameter is required.
	//
	// example:
	//
	// consumerGroupX
	ConsumerGroup *string `json:"consumerGroup,omitempty" xml:"consumerGroup,omitempty"`
	// Specifies whether to consume data in sequence. Valid values:
	//
	// 	- true
	//
	//     	- In a shard, data is consumed in ascending order based on the value of the \\*\\*__tag__:__receive_time__\\*\\	- field.
	//
	//     	- If a shard is split, data in the original shard is consumed first. Then, data in the new shards is consumed at the same time.
	//
	//     	- If shards are merged, data in the original shards is consumed first. Then, data in the new shard is consumed.
	//
	// 	- false Data in all shards is consumed at the same time. If a new shard is generated after a shard is split or after shards are merged, data in the new shard is immediately consumed.
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	Order *bool `json:"order,omitempty" xml:"order,omitempty"`
	// The timeout period. If the server does not receive heartbeats from a consumer within the timeout period, the server deletes the consumer. Unit: seconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 300
	Timeout *int32 `json:"timeout,omitempty" xml:"timeout,omitempty"`
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
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
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

type CreateDashboardRequest struct {
	// The data structure of the dashboard.
	//
	// This parameter is required.
	Body *Dashboard `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDashboardRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDashboardRequest) GoString() string {
	return s.String()
}

func (s *CreateDashboardRequest) SetBody(v *Dashboard) *CreateDashboardRequest {
	s.Body = v
	return s
}

type CreateDashboardResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s CreateDashboardResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDashboardResponse) GoString() string {
	return s.String()
}

func (s *CreateDashboardResponse) SetHeaders(v map[string]*string) *CreateDashboardResponse {
	s.Headers = v
	return s
}

func (s *CreateDashboardResponse) SetStatusCode(v int32) *CreateDashboardResponse {
	s.StatusCode = &v
	return s
}

type CreateDomainRequest struct {
	// The domain name.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
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
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
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

type CreateETLRequest struct {
	// This parameter is required.
	Configuration *ETLConfiguration `json:"configuration,omitempty" xml:"configuration,omitempty"`
	// example:
	//
	// this is ETL
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// sls-test-etl
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// etl-123456
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s CreateETLRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateETLRequest) GoString() string {
	return s.String()
}

func (s *CreateETLRequest) SetConfiguration(v *ETLConfiguration) *CreateETLRequest {
	s.Configuration = v
	return s
}

func (s *CreateETLRequest) SetDescription(v string) *CreateETLRequest {
	s.Description = &v
	return s
}

func (s *CreateETLRequest) SetDisplayName(v string) *CreateETLRequest {
	s.DisplayName = &v
	return s
}

func (s *CreateETLRequest) SetName(v string) *CreateETLRequest {
	s.Name = &v
	return s
}

type CreateETLResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s CreateETLResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateETLResponse) GoString() string {
	return s.String()
}

func (s *CreateETLResponse) SetHeaders(v map[string]*string) *CreateETLResponse {
	s.Headers = v
	return s
}

func (s *CreateETLResponse) SetStatusCode(v int32) *CreateETLResponse {
	s.StatusCode = &v
	return s
}

type CreateIndexRequest struct {
	// The configuration of field indexes. A field index is a key-value pair in which the key specifies the name of the field and the value specifies the index configuration of the field. You must specify this parameter, the line parameter, or both parameters. For more information, see Example.
	Keys map[string]*KeysValue `json:"keys,omitempty" xml:"keys,omitempty"`
	// The configuration of full-text indexes. You must specify this parameter, the keys parameter, or both parameters. For more information, see Example.
	Line *CreateIndexRequestLine `json:"line,omitempty" xml:"line,omitempty" type:"Struct"`
	// Specifies whether to turn on LogReduce. After you turn on LogReduce, either the whitelist or blacklist takes effect.
	//
	// example:
	//
	// false
	LogReduce *bool `json:"log_reduce,omitempty" xml:"log_reduce,omitempty"`
	// The fields in the blacklist that you want to use to cluster logs.
	LogReduceBlackList []*string `json:"log_reduce_black_list,omitempty" xml:"log_reduce_black_list,omitempty" type:"Repeated"`
	// The fields in the whitelist that you want to use to cluster logs.
	LogReduceWhiteList []*string `json:"log_reduce_white_list,omitempty" xml:"log_reduce_white_list,omitempty" type:"Repeated"`
	// The maximum length of a field value that can be retained. Default value: 2048. Unit: bytes. The default value is equal to 2 KB. You can change the value of max_text_len. Valid values: 64 to 16384.
	//
	// example:
	//
	// 2048
	MaxTextLen *int32 `json:"max_text_len,omitempty" xml:"max_text_len,omitempty"`
	// The retention period of logs. Unit: days. Valid values: 7, 30, and 90.
	//
	// example:
	//
	// 30
	Ttl *int32 `json:"ttl,omitempty" xml:"ttl,omitempty"`
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
	// Specifies whether to enable case sensitivity. Valid values:
	//
	// 	- true
	//
	// 	- false (default)
	//
	// example:
	//
	// false
	CaseSensitive *bool `json:"caseSensitive,omitempty" xml:"caseSensitive,omitempty"`
	// Specifies whether to include Chinese characters. Valid values:
	//
	// 	- true
	//
	// 	- false (default)
	//
	// example:
	//
	// false
	Chn *bool `json:"chn,omitempty" xml:"chn,omitempty"`
	// The excluded fields. You cannot specify both include_keys and exclude_keys.
	ExcludeKeys []*string `json:"exclude_keys,omitempty" xml:"exclude_keys,omitempty" type:"Repeated"`
	// The included fields. You cannot specify both include_keys and exclude_keys.
	IncludeKeys []*string `json:"include_keys,omitempty" xml:"include_keys,omitempty" type:"Repeated"`
	// The delimiters. You can specify a delimiter to delimit the content of a field value. For more information about delimiters, see Example.
	//
	// This parameter is required.
	Token []*string `json:"token,omitempty" xml:"token,omitempty" type:"Repeated"`
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
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
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

type CreateLogStoreRequest struct {
	// Specifies whether to record public IP addresses. Default value: false. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// false
	AppendMeta *bool `json:"appendMeta,omitempty" xml:"appendMeta,omitempty"`
	// Specifies whether to enable automatic sharding. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	AutoSplit *bool `json:"autoSplit,omitempty" xml:"autoSplit,omitempty"`
	// Specifies whether to enable the web tracking feature. Default value: false. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// false
	EnableTracking *bool `json:"enable_tracking,omitempty" xml:"enable_tracking,omitempty"`
	// The data structure of the encryption configuration.
	EncryptConf *EncryptConf `json:"encrypt_conf,omitempty" xml:"encrypt_conf,omitempty"`
	// The retention period of data in the hot storage tier of the Logstore. Unit: days. You can specify a value that ranges from 30 to the value of ttl.
	//
	// Hot data that is stored for longer than the period specified by hot_ttl is converted to cold data. For more information, see [Enable hot and cold-tiered storage for a Logstore](https://help.aliyun.com/document_detail/308645.html).
	//
	// example:
	//
	// 60
	HotTtl              *int32 `json:"hot_ttl,omitempty" xml:"hot_ttl,omitempty"`
	InfrequentAccessTTL *int32 `json:"infrequentAccessTTL,omitempty" xml:"infrequentAccessTTL,omitempty"`
	// The name of the Logstore. The name must meet the following requirements:
	//
	// 	- The name must be unique in a project.
	//
	// 	- The name can contain only lowercase letters, digits, hyphens (-), and underscores (_).
	//
	// 	- The name must start and end with a lowercase letter or a digit.
	//
	// 	- The name must be 3 to 63 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// my-logstore
	LogstoreName *string `json:"logstoreName,omitempty" xml:"logstoreName,omitempty"`
	// The maximum number of shards into which existing shards can be automatically split. Valid values: 1 to 64.
	//
	// > If you set autoSplit to true, you must configure this parameter.
	//
	// example:
	//
	// 64
	MaxSplitShard *int32 `json:"maxSplitShard,omitempty" xml:"maxSplitShard,omitempty"`
	// The type of the Logstore. Log Service provides the following types of Logstores: Standard Logstores and Query Logstores. Valid values:
	//
	// 	- **standard**: Standard Logstore. This type of Logstore supports the log analysis feature and is suitable for scenarios such as real-time monitoring and interactive analysis. You can also use this type of Logstore to build a comprehensive observability system.
	//
	// 	- **query**: Query Logstore. This type of Logstore supports high-performance queries. The index traffic fee of a Query Logstore is approximately half that of a Standard Logstore. Query Logstores do not support SQL analysis. Query Logstores are suitable for scenarios in which the amount of data is large, the log retention period is long, or log analysis is not required. Log retention periods of weeks or months are considered long.
	//
	// example:
	//
	// standard
	Mode *string `json:"mode,omitempty" xml:"mode,omitempty"`
	// The number of shards.
	//
	// > You cannot call the CreateLogStore operation to change the number of shards. You can call the SplitShard or MergeShards operation to change the number of shards.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2
	ShardCount *int32 `json:"shardCount,omitempty" xml:"shardCount,omitempty"`
	// The type of the observable data. Valid values:
	//
	// 	- None: logs
	//
	// 	- Metrics: metrics
	//
	// example:
	//
	// None
	TelemetryType *string `json:"telemetryType,omitempty" xml:"telemetryType,omitempty"`
	// The retention period of data. Unit: days. Valid values: 1 to 3000. If you set this parameter to 3650, data is permanently stored.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	Ttl *int32 `json:"ttl,omitempty" xml:"ttl,omitempty"`
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

func (s *CreateLogStoreRequest) SetInfrequentAccessTTL(v int32) *CreateLogStoreRequest {
	s.InfrequentAccessTTL = &v
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
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
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
	// The configurations of service logs.
	//
	// This parameter is required.
	LoggingDetails []*CreateLoggingRequestLoggingDetails `json:"loggingDetails,omitempty" xml:"loggingDetails,omitempty" type:"Repeated"`
	// The name of the project to which service logs are stored.
	//
	// This parameter is required.
	//
	// example:
	//
	// my-project
	LoggingProject *string `json:"loggingProject,omitempty" xml:"loggingProject,omitempty"`
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
	// The name of the Logstore to which service logs of the type are stored.
	//
	// This parameter is required.
	//
	// example:
	//
	// my-logstore
	Logstore *string `json:"logstore,omitempty" xml:"logstore,omitempty"`
	// The type of service logs. Valid values:
	//
	// 	- consumergroup_log: the consumption delay logs of consumer groups.
	//
	// 	- logtail_alarm: the alert logs of Logtail.
	//
	// 	- operation_log: the operation logs.
	//
	// 	- logtail_profile: the collection logs of Logtail.
	//
	// 	- metering: the metering logs.
	//
	// 	- logtail_status: the status logs of Logtail.
	//
	// 	- scheduledsqlalert: the run logs of Scheduled SQL jobs.
	//
	// 	- etl_alert: the run logs of data transformation jobs.
	//
	// This parameter is required.
	//
	// example:
	//
	// consumergroup_log
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
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
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
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

type CreateLogtailPipelineConfigRequest struct {
	// The aggregation plug-ins.
	Aggregators []map[string]interface{} `json:"aggregators,omitempty" xml:"aggregators,omitempty" type:"Repeated"`
	// The name of the configuration.
	//
	// This parameter is required.
	//
	// example:
	//
	// test-config
	ConfigName *string `json:"configName,omitempty" xml:"configName,omitempty"`
	// The data output plug-ins.
	//
	// This parameter is required.
	Flushers []map[string]interface{} `json:"flushers,omitempty" xml:"flushers,omitempty" type:"Repeated"`
	// The global configuration.
	Global map[string]interface{} `json:"global,omitempty" xml:"global,omitempty"`
	// The data source plug-ins.
	//
	// This parameter is required.
	Inputs []map[string]interface{} `json:"inputs,omitempty" xml:"inputs,omitempty" type:"Repeated"`
	// The sample log.
	//
	// example:
	//
	// 2022-06-14 11:13:29.796 | DEBUG    | __main__:<module>:1 - hello world
	LogSample *string `json:"logSample,omitempty" xml:"logSample,omitempty"`
	// The processing plug-ins.
	Processors []map[string]interface{} `json:"processors,omitempty" xml:"processors,omitempty" type:"Repeated"`
}

func (s CreateLogtailPipelineConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateLogtailPipelineConfigRequest) GoString() string {
	return s.String()
}

func (s *CreateLogtailPipelineConfigRequest) SetAggregators(v []map[string]interface{}) *CreateLogtailPipelineConfigRequest {
	s.Aggregators = v
	return s
}

func (s *CreateLogtailPipelineConfigRequest) SetConfigName(v string) *CreateLogtailPipelineConfigRequest {
	s.ConfigName = &v
	return s
}

func (s *CreateLogtailPipelineConfigRequest) SetFlushers(v []map[string]interface{}) *CreateLogtailPipelineConfigRequest {
	s.Flushers = v
	return s
}

func (s *CreateLogtailPipelineConfigRequest) SetGlobal(v map[string]interface{}) *CreateLogtailPipelineConfigRequest {
	s.Global = v
	return s
}

func (s *CreateLogtailPipelineConfigRequest) SetInputs(v []map[string]interface{}) *CreateLogtailPipelineConfigRequest {
	s.Inputs = v
	return s
}

func (s *CreateLogtailPipelineConfigRequest) SetLogSample(v string) *CreateLogtailPipelineConfigRequest {
	s.LogSample = &v
	return s
}

func (s *CreateLogtailPipelineConfigRequest) SetProcessors(v []map[string]interface{}) *CreateLogtailPipelineConfigRequest {
	s.Processors = v
	return s
}

type CreateLogtailPipelineConfigResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s CreateLogtailPipelineConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateLogtailPipelineConfigResponse) GoString() string {
	return s.String()
}

func (s *CreateLogtailPipelineConfigResponse) SetHeaders(v map[string]*string) *CreateLogtailPipelineConfigResponse {
	s.Headers = v
	return s
}

func (s *CreateLogtailPipelineConfigResponse) SetStatusCode(v int32) *CreateLogtailPipelineConfigResponse {
	s.StatusCode = &v
	return s
}

type CreateMachineGroupRequest struct {
	// The attributes of the machine group.
	GroupAttribute *CreateMachineGroupRequestGroupAttribute `json:"groupAttribute,omitempty" xml:"groupAttribute,omitempty" type:"Struct"`
	// The name of the machine group. The name must meet the following requirements:
	//
	// 	- The name of each machine group in a project must be unique.
	//
	// 	- It can contain only lowercase letters, digits, hyphens (-), and underscores (_).
	//
	// 	- It must start and end with a lowercase letter or a digit.
	//
	// 	- It must be 3 to 128 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// test-machine-group
	GroupName *string `json:"groupName,omitempty" xml:"groupName,omitempty"`
	// The type of the machine group. The parameter can be left empty.
	//
	// example:
	//
	// ""
	GroupType *string `json:"groupType,omitempty" xml:"groupType,omitempty"`
	// The type of the machine group identifier. Valid values:
	//
	// 	- ip: The machine group uses IP addresses as identifiers.
	//
	// 	- userdefined: The machine group uses custom identifiers.
	//
	// This parameter is required.
	//
	// example:
	//
	// ip
	MachineIdentifyType *string `json:"machineIdentifyType,omitempty" xml:"machineIdentifyType,omitempty"`
	// The identifiers of machine group.
	//
	// 	- If you set machineIdentifyType to ip, enter the IP address of the machine.
	//
	// 	- If you set machineIdentifyType to userdefined, enter a custom identifier.
	//
	// This parameter is required.
	MachineList []*string `json:"machineList,omitempty" xml:"machineList,omitempty" type:"Repeated"`
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
	// The identifier of the external management system on which the machine group depends.
	//
	// example:
	//
	// testgroup
	ExternalName *string `json:"externalName,omitempty" xml:"externalName,omitempty"`
	// The log topic of the machine group.
	//
	// example:
	//
	// testtopic
	GroupTopic *string `json:"groupTopic,omitempty" xml:"groupTopic,omitempty"`
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
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
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

type CreateOSSExportRequest struct {
	// This parameter is required.
	Configuration *OSSExportConfiguration `json:"configuration,omitempty" xml:"configuration,omitempty"`
	Description   *string                 `json:"description,omitempty" xml:"description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ali-test-oss-job
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// job-123456789-123456
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s CreateOSSExportRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateOSSExportRequest) GoString() string {
	return s.String()
}

func (s *CreateOSSExportRequest) SetConfiguration(v *OSSExportConfiguration) *CreateOSSExportRequest {
	s.Configuration = v
	return s
}

func (s *CreateOSSExportRequest) SetDescription(v string) *CreateOSSExportRequest {
	s.Description = &v
	return s
}

func (s *CreateOSSExportRequest) SetDisplayName(v string) *CreateOSSExportRequest {
	s.DisplayName = &v
	return s
}

func (s *CreateOSSExportRequest) SetName(v string) *CreateOSSExportRequest {
	s.Name = &v
	return s
}

type CreateOSSExportResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s CreateOSSExportResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateOSSExportResponse) GoString() string {
	return s.String()
}

func (s *CreateOSSExportResponse) SetHeaders(v map[string]*string) *CreateOSSExportResponse {
	s.Headers = v
	return s
}

func (s *CreateOSSExportResponse) SetStatusCode(v int32) *CreateOSSExportResponse {
	s.StatusCode = &v
	return s
}

type CreateOSSHDFSExportRequest struct {
	// This parameter is required.
	Configuration *OSSExportConfiguration `json:"configuration,omitempty" xml:"configuration,omitempty"`
	Description   *string                 `json:"description,omitempty" xml:"description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ali-test-oss-hdfs-job
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// job-123456789-123456
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s CreateOSSHDFSExportRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateOSSHDFSExportRequest) GoString() string {
	return s.String()
}

func (s *CreateOSSHDFSExportRequest) SetConfiguration(v *OSSExportConfiguration) *CreateOSSHDFSExportRequest {
	s.Configuration = v
	return s
}

func (s *CreateOSSHDFSExportRequest) SetDescription(v string) *CreateOSSHDFSExportRequest {
	s.Description = &v
	return s
}

func (s *CreateOSSHDFSExportRequest) SetDisplayName(v string) *CreateOSSHDFSExportRequest {
	s.DisplayName = &v
	return s
}

func (s *CreateOSSHDFSExportRequest) SetName(v string) *CreateOSSHDFSExportRequest {
	s.Name = &v
	return s
}

type CreateOSSHDFSExportResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s CreateOSSHDFSExportResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateOSSHDFSExportResponse) GoString() string {
	return s.String()
}

func (s *CreateOSSHDFSExportResponse) SetHeaders(v map[string]*string) *CreateOSSHDFSExportResponse {
	s.Headers = v
	return s
}

func (s *CreateOSSHDFSExportResponse) SetStatusCode(v int32) *CreateOSSHDFSExportResponse {
	s.StatusCode = &v
	return s
}

type CreateOSSIngestionRequest struct {
	// This parameter is required.
	Configuration *OSSIngestionConfiguration `json:"configuration,omitempty" xml:"configuration,omitempty"`
	Description   *string                    `json:"description,omitempty" xml:"description,omitempty"`
	// This parameter is required.
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ingest-oss-123456
	Name     *string   `json:"name,omitempty" xml:"name,omitempty"`
	Schedule *Schedule `json:"schedule,omitempty" xml:"schedule,omitempty"`
}

func (s CreateOSSIngestionRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateOSSIngestionRequest) GoString() string {
	return s.String()
}

func (s *CreateOSSIngestionRequest) SetConfiguration(v *OSSIngestionConfiguration) *CreateOSSIngestionRequest {
	s.Configuration = v
	return s
}

func (s *CreateOSSIngestionRequest) SetDescription(v string) *CreateOSSIngestionRequest {
	s.Description = &v
	return s
}

func (s *CreateOSSIngestionRequest) SetDisplayName(v string) *CreateOSSIngestionRequest {
	s.DisplayName = &v
	return s
}

func (s *CreateOSSIngestionRequest) SetName(v string) *CreateOSSIngestionRequest {
	s.Name = &v
	return s
}

func (s *CreateOSSIngestionRequest) SetSchedule(v *Schedule) *CreateOSSIngestionRequest {
	s.Schedule = v
	return s
}

type CreateOSSIngestionResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s CreateOSSIngestionResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateOSSIngestionResponse) GoString() string {
	return s.String()
}

func (s *CreateOSSIngestionResponse) SetHeaders(v map[string]*string) *CreateOSSIngestionResponse {
	s.Headers = v
	return s
}

func (s *CreateOSSIngestionResponse) SetStatusCode(v int32) *CreateOSSIngestionResponse {
	s.StatusCode = &v
	return s
}

type CreateOssExternalStoreRequest struct {
	// The name of the external store.
	//
	// This parameter is required.
	//
	// example:
	//
	// test_oss_store
	ExternalStoreName *string `json:"externalStoreName,omitempty" xml:"externalStoreName,omitempty"`
	// The parameters that are configured for the external store.
	//
	// This parameter is required.
	Parameter *CreateOssExternalStoreRequestParameter `json:"parameter,omitempty" xml:"parameter,omitempty" type:"Struct"`
	// The type of the external store. Set the value to oss.
	//
	// This parameter is required.
	//
	// example:
	//
	// oss
	StoreType *string `json:"storeType,omitempty" xml:"storeType,omitempty"`
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
	// The AccessKey ID of your account.
	//
	// This parameter is required.
	//
	// example:
	//
	// LTAI5tFsHGGeYry*****1Sz
	Accessid *string `json:"accessid,omitempty" xml:"accessid,omitempty"`
	// The AccessKey secret of your account.
	//
	// This parameter is required.
	//
	// example:
	//
	// GyviCLDVHkHrOztdkxuE6******Rp6
	Accesskey *string `json:"accesskey,omitempty" xml:"accesskey,omitempty"`
	// The name of the OSS bucket.
	//
	// This parameter is required.
	//
	// example:
	//
	// my-bucket
	Bucket *string `json:"bucket,omitempty" xml:"bucket,omitempty"`
	// The fields that are associated to the external store.
	//
	// This parameter is required.
	Columns []*CreateOssExternalStoreRequestParameterColumns `json:"columns,omitempty" xml:"columns,omitempty" type:"Repeated"`
	// The Object Storage Service (OSS) endpoint.
	//
	// This parameter is required.
	//
	// example:
	//
	// oss-cn-hangzhou.aliyuncs.com
	Endpoint *string `json:"endpoint,omitempty" xml:"endpoint,omitempty"`
	// The names of the OSS objects that are associated to the external store.
	//
	// This parameter is required.
	Objects []*string `json:"objects,omitempty" xml:"objects,omitempty" type:"Repeated"`
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
	// The name of the field.
	//
	// This parameter is required.
	//
	// example:
	//
	// auto-test
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The type of the field.
	//
	// This parameter is required.
	//
	// example:
	//
	// sls
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
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
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

type CreateProjectRequest struct {
	// Data redundancy type
	//
	// example:
	//
	// LRS
	DataRedundancyType *string `json:"dataRedundancyType,omitempty" xml:"dataRedundancyType,omitempty"`
	// The description of the project.
	//
	// This parameter is required.
	//
	// example:
	//
	// this is test
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The name of the project. The name must be unique in a region. You cannot change the name after you create the project. The name must meet the following requirements:
	//
	// 	- The name must be unique.
	//
	// 	- It can contain only lowercase letters, digits, and hyphens (-).
	//
	// 	- It must start and end with a lowercase letter or a digit.
	//
	// 	- It must be 3 to 63 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// test-project
	ProjectName *string `json:"projectName,omitempty" xml:"projectName,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-aekzf******sxby
	ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
}

func (s CreateProjectRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateProjectRequest) GoString() string {
	return s.String()
}

func (s *CreateProjectRequest) SetDataRedundancyType(v string) *CreateProjectRequest {
	s.DataRedundancyType = &v
	return s
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
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
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
	// The name of the external store. The name must be unique in a project and must be different from Logstore names.
	//
	// This parameter is required.
	//
	// example:
	//
	// rds_store
	ExternalStoreName *string `json:"externalStoreName,omitempty" xml:"externalStoreName,omitempty"`
	// The parameter struct.
	//
	// This parameter is required.
	Parameter *CreateRdsExternalStoreRequestParameter `json:"parameter,omitempty" xml:"parameter,omitempty" type:"Struct"`
	// The storage type. Set the value to rds-vpc, which indicates an ApsaraDB RDS for MySQL database in a virtual private cloud (VPC).
	//
	// This parameter is required.
	//
	// example:
	//
	// rds-vpc
	StoreType *string `json:"storeType,omitempty" xml:"storeType,omitempty"`
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
	// The name of the database in the ApsaraDB RDS for MySQL instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// meta
	Db *string `json:"db,omitempty" xml:"db,omitempty"`
	// The internal or public endpoint of the ApsaraDB RDS for MySQL instance.
	//
	// example:
	//
	// 192.168.XX.XX
	Host *string `json:"host,omitempty" xml:"host,omitempty"`
	// The ID of the ApsaraDB RDS for MySQL instance.
	//
	// example:
	//
	// i-bp1b6c719dfa08exf****
	InstanceId *string `json:"instance-id,omitempty" xml:"instance-id,omitempty"`
	// The password that is used to log on to the ApsaraDB RDS for MySQL instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// sfdsfldsfksfls****
	Password *string `json:"password,omitempty" xml:"password,omitempty"`
	// The internal or public port of the ApsaraDB RDS for MySQL instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3306
	Port *string `json:"port,omitempty" xml:"port,omitempty"`
	// The region where the ApsaraDB RDS for MySQL instance resides. Valid values: cn-qingdao, cn-beijing, and cn-hangzhou.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-qingdao
	Region *string `json:"region,omitempty" xml:"region,omitempty"`
	// The name of the database table in the ApsaraDB RDS for MySQL instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// join_meta
	Table *string `json:"table,omitempty" xml:"table,omitempty"`
	// The username that is used to log on to the ApsaraDB RDS for MySQL instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// root
	Username *string `json:"username,omitempty" xml:"username,omitempty"`
	// The ID of the VPC to which the ApsaraDB RDS for MySQL instance belongs.
	//
	// example:
	//
	// vpc-bp1aevy8sofi8mh1q****
	VpcId *string `json:"vpc-id,omitempty" xml:"vpc-id,omitempty"`
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
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
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
	// The display name.
	//
	// This parameter is required.
	//
	// example:
	//
	// displayname
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	// The name of the Logstore to which the saved search belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// aliyun-test-logstore
	Logstore *string `json:"logstore,omitempty" xml:"logstore,omitempty"`
	// The name of the saved search. The name must be 3 to 63 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// pv in minutes
	SavedsearchName *string `json:"savedsearchName,omitempty" xml:"savedsearchName,omitempty"`
	// The query statement of the saved search. A query statement consists of a search statement and an analytic statement in the `Search statement|Analytic statement` format. For more information about search statements and analytic statements, see [Log search overview](https://help.aliyun.com/document_detail/43772.html) and [Log analysis overview](https://help.aliyun.com/document_detail/53608.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// *|select date_format(__time__-__time__%60, \\"%H:%i:%s\\") as time, COUNT(*) as pv group by time
	SearchQuery *string `json:"searchQuery,omitempty" xml:"searchQuery,omitempty"`
	// The topic of the log.
	//
	// example:
	//
	// theme
	Topic *string `json:"topic,omitempty" xml:"topic,omitempty"`
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
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
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

type CreateScheduledSQLRequest struct {
	// This parameter is required.
	Configuration *ScheduledSQLConfiguration `json:"configuration,omitempty" xml:"configuration,omitempty"`
	Description   *string                    `json:"description,omitempty" xml:"description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ali-test-scheduled-sql
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// sql-123456789-123456
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// This parameter is required.
	Schedule *Schedule `json:"schedule,omitempty" xml:"schedule,omitempty"`
}

func (s CreateScheduledSQLRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateScheduledSQLRequest) GoString() string {
	return s.String()
}

func (s *CreateScheduledSQLRequest) SetConfiguration(v *ScheduledSQLConfiguration) *CreateScheduledSQLRequest {
	s.Configuration = v
	return s
}

func (s *CreateScheduledSQLRequest) SetDescription(v string) *CreateScheduledSQLRequest {
	s.Description = &v
	return s
}

func (s *CreateScheduledSQLRequest) SetDisplayName(v string) *CreateScheduledSQLRequest {
	s.DisplayName = &v
	return s
}

func (s *CreateScheduledSQLRequest) SetName(v string) *CreateScheduledSQLRequest {
	s.Name = &v
	return s
}

func (s *CreateScheduledSQLRequest) SetSchedule(v *Schedule) *CreateScheduledSQLRequest {
	s.Schedule = v
	return s
}

type CreateScheduledSQLResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s CreateScheduledSQLResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateScheduledSQLResponse) GoString() string {
	return s.String()
}

func (s *CreateScheduledSQLResponse) SetHeaders(v map[string]*string) *CreateScheduledSQLResponse {
	s.Headers = v
	return s
}

func (s *CreateScheduledSQLResponse) SetStatusCode(v int32) *CreateScheduledSQLResponse {
	s.StatusCode = &v
	return s
}

type CreateSqlInstanceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 2
	Cu *int32 `json:"cu,omitempty" xml:"cu,omitempty"`
	// This parameter is required.
	UseAsDefault *bool `json:"useAsDefault,omitempty" xml:"useAsDefault,omitempty"`
}

func (s CreateSqlInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateSqlInstanceRequest) GoString() string {
	return s.String()
}

func (s *CreateSqlInstanceRequest) SetCu(v int32) *CreateSqlInstanceRequest {
	s.Cu = &v
	return s
}

func (s *CreateSqlInstanceRequest) SetUseAsDefault(v bool) *CreateSqlInstanceRequest {
	s.UseAsDefault = &v
	return s
}

type CreateSqlInstanceResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s CreateSqlInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateSqlInstanceResponse) GoString() string {
	return s.String()
}

func (s *CreateSqlInstanceResponse) SetHeaders(v map[string]*string) *CreateSqlInstanceResponse {
	s.Headers = v
	return s
}

func (s *CreateSqlInstanceResponse) SetStatusCode(v int32) *CreateSqlInstanceResponse {
	s.StatusCode = &v
	return s
}

type CreateStoreViewRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// my_storeview
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// logstore
	StoreType *string `json:"storeType,omitempty" xml:"storeType,omitempty"`
	// This parameter is required.
	Stores []*StoreViewStore `json:"stores,omitempty" xml:"stores,omitempty" type:"Repeated"`
}

func (s CreateStoreViewRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateStoreViewRequest) GoString() string {
	return s.String()
}

func (s *CreateStoreViewRequest) SetName(v string) *CreateStoreViewRequest {
	s.Name = &v
	return s
}

func (s *CreateStoreViewRequest) SetStoreType(v string) *CreateStoreViewRequest {
	s.StoreType = &v
	return s
}

func (s *CreateStoreViewRequest) SetStores(v []*StoreViewStore) *CreateStoreViewRequest {
	s.Stores = v
	return s
}

type CreateStoreViewResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s CreateStoreViewResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateStoreViewResponse) GoString() string {
	return s.String()
}

func (s *CreateStoreViewResponse) SetHeaders(v map[string]*string) *CreateStoreViewResponse {
	s.Headers = v
	return s
}

func (s *CreateStoreViewResponse) SetStatusCode(v int32) *CreateStoreViewResponse {
	s.StatusCode = &v
	return s
}

type CreateTicketRequest struct {
	AccessTokenExpirationTime *int64 `json:"accessTokenExpirationTime,omitempty" xml:"accessTokenExpirationTime,omitempty"`
	ExpirationTime            *int64 `json:"expirationTime,omitempty" xml:"expirationTime,omitempty"`
}

func (s CreateTicketRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateTicketRequest) GoString() string {
	return s.String()
}

func (s *CreateTicketRequest) SetAccessTokenExpirationTime(v int64) *CreateTicketRequest {
	s.AccessTokenExpirationTime = &v
	return s
}

func (s *CreateTicketRequest) SetExpirationTime(v int64) *CreateTicketRequest {
	s.ExpirationTime = &v
	return s
}

type CreateTicketResponseBody struct {
	// example:
	//
	// eyJ***************.eyJ******************.KUT****************
	Ticket *string `json:"ticket,omitempty" xml:"ticket,omitempty"`
}

func (s CreateTicketResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateTicketResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTicketResponseBody) SetTicket(v string) *CreateTicketResponseBody {
	s.Ticket = &v
	return s
}

type CreateTicketResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateTicketResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateTicketResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateTicketResponse) GoString() string {
	return s.String()
}

func (s *CreateTicketResponse) SetHeaders(v map[string]*string) *CreateTicketResponse {
	s.Headers = v
	return s
}

func (s *CreateTicketResponse) SetStatusCode(v int32) *CreateTicketResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateTicketResponse) SetBody(v *CreateTicketResponseBody) *CreateTicketResponse {
	s.Body = v
	return s
}

type DeleteAlertResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
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

type DeleteAnnotationDataResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s DeleteAnnotationDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteAnnotationDataResponse) GoString() string {
	return s.String()
}

func (s *DeleteAnnotationDataResponse) SetHeaders(v map[string]*string) *DeleteAnnotationDataResponse {
	s.Headers = v
	return s
}

func (s *DeleteAnnotationDataResponse) SetStatusCode(v int32) *DeleteAnnotationDataResponse {
	s.StatusCode = &v
	return s
}

type DeleteAnnotationDataSetResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s DeleteAnnotationDataSetResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteAnnotationDataSetResponse) GoString() string {
	return s.String()
}

func (s *DeleteAnnotationDataSetResponse) SetHeaders(v map[string]*string) *DeleteAnnotationDataSetResponse {
	s.Headers = v
	return s
}

func (s *DeleteAnnotationDataSetResponse) SetStatusCode(v int32) *DeleteAnnotationDataSetResponse {
	s.StatusCode = &v
	return s
}

type DeleteAnnotationLabelResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s DeleteAnnotationLabelResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteAnnotationLabelResponse) GoString() string {
	return s.String()
}

func (s *DeleteAnnotationLabelResponse) SetHeaders(v map[string]*string) *DeleteAnnotationLabelResponse {
	s.Headers = v
	return s
}

func (s *DeleteAnnotationLabelResponse) SetStatusCode(v int32) *DeleteAnnotationLabelResponse {
	s.StatusCode = &v
	return s
}

type DeleteCollectionPolicyRequest struct {
	// example:
	//
	// access_log
	DataCode *string `json:"dataCode,omitempty" xml:"dataCode,omitempty"`
	// example:
	//
	// oss
	ProductCode *string `json:"productCode,omitempty" xml:"productCode,omitempty"`
}

func (s DeleteCollectionPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteCollectionPolicyRequest) GoString() string {
	return s.String()
}

func (s *DeleteCollectionPolicyRequest) SetDataCode(v string) *DeleteCollectionPolicyRequest {
	s.DataCode = &v
	return s
}

func (s *DeleteCollectionPolicyRequest) SetProductCode(v string) *DeleteCollectionPolicyRequest {
	s.ProductCode = &v
	return s
}

type DeleteCollectionPolicyResponseBody struct {
	// example:
	//
	// successful
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
}

func (s DeleteCollectionPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteCollectionPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCollectionPolicyResponseBody) SetMessage(v string) *DeleteCollectionPolicyResponseBody {
	s.Message = &v
	return s
}

type DeleteCollectionPolicyResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteCollectionPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteCollectionPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteCollectionPolicyResponse) GoString() string {
	return s.String()
}

func (s *DeleteCollectionPolicyResponse) SetHeaders(v map[string]*string) *DeleteCollectionPolicyResponse {
	s.Headers = v
	return s
}

func (s *DeleteCollectionPolicyResponse) SetStatusCode(v int32) *DeleteCollectionPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCollectionPolicyResponse) SetBody(v *DeleteCollectionPolicyResponseBody) *DeleteCollectionPolicyResponse {
	s.Body = v
	return s
}

type DeleteConfigResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s DeleteConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteConfigResponse) GoString() string {
	return s.String()
}

func (s *DeleteConfigResponse) SetHeaders(v map[string]*string) *DeleteConfigResponse {
	s.Headers = v
	return s
}

func (s *DeleteConfigResponse) SetStatusCode(v int32) *DeleteConfigResponse {
	s.StatusCode = &v
	return s
}

type DeleteConsumerGroupResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
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

type DeleteDashboardResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s DeleteDashboardResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDashboardResponse) GoString() string {
	return s.String()
}

func (s *DeleteDashboardResponse) SetHeaders(v map[string]*string) *DeleteDashboardResponse {
	s.Headers = v
	return s
}

func (s *DeleteDashboardResponse) SetStatusCode(v int32) *DeleteDashboardResponse {
	s.StatusCode = &v
	return s
}

type DeleteDomainResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
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

type DeleteETLResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s DeleteETLResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteETLResponse) GoString() string {
	return s.String()
}

func (s *DeleteETLResponse) SetHeaders(v map[string]*string) *DeleteETLResponse {
	s.Headers = v
	return s
}

func (s *DeleteETLResponse) SetStatusCode(v int32) *DeleteETLResponse {
	s.StatusCode = &v
	return s
}

type DeleteExternalStoreResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
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
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
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

type DeleteLogStoreResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
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
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
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

type DeleteLogtailPipelineConfigResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s DeleteLogtailPipelineConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteLogtailPipelineConfigResponse) GoString() string {
	return s.String()
}

func (s *DeleteLogtailPipelineConfigResponse) SetHeaders(v map[string]*string) *DeleteLogtailPipelineConfigResponse {
	s.Headers = v
	return s
}

func (s *DeleteLogtailPipelineConfigResponse) SetStatusCode(v int32) *DeleteLogtailPipelineConfigResponse {
	s.StatusCode = &v
	return s
}

type DeleteMachineGroupResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
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

type DeleteOSSExportResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s DeleteOSSExportResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteOSSExportResponse) GoString() string {
	return s.String()
}

func (s *DeleteOSSExportResponse) SetHeaders(v map[string]*string) *DeleteOSSExportResponse {
	s.Headers = v
	return s
}

func (s *DeleteOSSExportResponse) SetStatusCode(v int32) *DeleteOSSExportResponse {
	s.StatusCode = &v
	return s
}

type DeleteOSSHDFSExportResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s DeleteOSSHDFSExportResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteOSSHDFSExportResponse) GoString() string {
	return s.String()
}

func (s *DeleteOSSHDFSExportResponse) SetHeaders(v map[string]*string) *DeleteOSSHDFSExportResponse {
	s.Headers = v
	return s
}

func (s *DeleteOSSHDFSExportResponse) SetStatusCode(v int32) *DeleteOSSHDFSExportResponse {
	s.StatusCode = &v
	return s
}

type DeleteOSSIngestionResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s DeleteOSSIngestionResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteOSSIngestionResponse) GoString() string {
	return s.String()
}

func (s *DeleteOSSIngestionResponse) SetHeaders(v map[string]*string) *DeleteOSSIngestionResponse {
	s.Headers = v
	return s
}

func (s *DeleteOSSIngestionResponse) SetStatusCode(v int32) *DeleteOSSIngestionResponse {
	s.StatusCode = &v
	return s
}

type DeleteProjectResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
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
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
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
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
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

type DeleteScheduledSQLResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s DeleteScheduledSQLResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteScheduledSQLResponse) GoString() string {
	return s.String()
}

func (s *DeleteScheduledSQLResponse) SetHeaders(v map[string]*string) *DeleteScheduledSQLResponse {
	s.Headers = v
	return s
}

func (s *DeleteScheduledSQLResponse) SetStatusCode(v int32) *DeleteScheduledSQLResponse {
	s.StatusCode = &v
	return s
}

type DeleteShipperResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
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

type DeleteStoreViewResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s DeleteStoreViewResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteStoreViewResponse) GoString() string {
	return s.String()
}

func (s *DeleteStoreViewResponse) SetHeaders(v map[string]*string) *DeleteStoreViewResponse {
	s.Headers = v
	return s
}

func (s *DeleteStoreViewResponse) SetStatusCode(v int32) *DeleteStoreViewResponse {
	s.StatusCode = &v
	return s
}

type DisableAlertResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s DisableAlertResponse) String() string {
	return tea.Prettify(s)
}

func (s DisableAlertResponse) GoString() string {
	return s.String()
}

func (s *DisableAlertResponse) SetHeaders(v map[string]*string) *DisableAlertResponse {
	s.Headers = v
	return s
}

func (s *DisableAlertResponse) SetStatusCode(v int32) *DisableAlertResponse {
	s.StatusCode = &v
	return s
}

type DisableScheduledSQLResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s DisableScheduledSQLResponse) String() string {
	return tea.Prettify(s)
}

func (s DisableScheduledSQLResponse) GoString() string {
	return s.String()
}

func (s *DisableScheduledSQLResponse) SetHeaders(v map[string]*string) *DisableScheduledSQLResponse {
	s.Headers = v
	return s
}

func (s *DisableScheduledSQLResponse) SetStatusCode(v int32) *DisableScheduledSQLResponse {
	s.StatusCode = &v
	return s
}

type EnableAlertResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s EnableAlertResponse) String() string {
	return tea.Prettify(s)
}

func (s EnableAlertResponse) GoString() string {
	return s.String()
}

func (s *EnableAlertResponse) SetHeaders(v map[string]*string) *EnableAlertResponse {
	s.Headers = v
	return s
}

func (s *EnableAlertResponse) SetStatusCode(v int32) *EnableAlertResponse {
	s.StatusCode = &v
	return s
}

type EnableScheduledSQLResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s EnableScheduledSQLResponse) String() string {
	return tea.Prettify(s)
}

func (s EnableScheduledSQLResponse) GoString() string {
	return s.String()
}

func (s *EnableScheduledSQLResponse) SetHeaders(v map[string]*string) *EnableScheduledSQLResponse {
	s.Headers = v
	return s
}

func (s *EnableScheduledSQLResponse) SetStatusCode(v int32) *EnableScheduledSQLResponse {
	s.StatusCode = &v
	return s
}

type GetAlertResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *Alert             `json:"body,omitempty" xml:"body,omitempty"`
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

type GetAnnotationDataResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MLDataParam       `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAnnotationDataResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAnnotationDataResponse) GoString() string {
	return s.String()
}

func (s *GetAnnotationDataResponse) SetHeaders(v map[string]*string) *GetAnnotationDataResponse {
	s.Headers = v
	return s
}

func (s *GetAnnotationDataResponse) SetStatusCode(v int32) *GetAnnotationDataResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAnnotationDataResponse) SetBody(v *MLDataParam) *GetAnnotationDataResponse {
	s.Body = v
	return s
}

type GetAnnotationDataSetResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MLDataSetParam    `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAnnotationDataSetResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAnnotationDataSetResponse) GoString() string {
	return s.String()
}

func (s *GetAnnotationDataSetResponse) SetHeaders(v map[string]*string) *GetAnnotationDataSetResponse {
	s.Headers = v
	return s
}

func (s *GetAnnotationDataSetResponse) SetStatusCode(v int32) *GetAnnotationDataSetResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAnnotationDataSetResponse) SetBody(v *MLDataSetParam) *GetAnnotationDataSetResponse {
	s.Body = v
	return s
}

type GetAnnotationLabelResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MLLabelParam      `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAnnotationLabelResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAnnotationLabelResponse) GoString() string {
	return s.String()
}

func (s *GetAnnotationLabelResponse) SetHeaders(v map[string]*string) *GetAnnotationLabelResponse {
	s.Headers = v
	return s
}

func (s *GetAnnotationLabelResponse) SetStatusCode(v int32) *GetAnnotationLabelResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAnnotationLabelResponse) SetBody(v *MLLabelParam) *GetAnnotationLabelResponse {
	s.Body = v
	return s
}

type GetAppliedConfigsResponseBody struct {
	// The names of the Logtail configurations.
	Configs []*string `json:"configs,omitempty" xml:"configs,omitempty" type:"Repeated"`
	// The number of Logtail configurations.
	//
	// example:
	//
	// 2
	Count *int32 `json:"count,omitempty" xml:"count,omitempty"`
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
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAppliedConfigsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The number of returned machine groups.
	//
	// example:
	//
	// 2
	Count *int32 `json:"count,omitempty" xml:"count,omitempty"`
	// The names of the returned machine groups.
	//
	// example:
	//
	// [ "sample-group1","sample-group2" ]
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
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAppliedMachineGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The shard ID.
	//
	// 	- If the specified shard does not exist, an empty list is returned.
	//
	// 	- If no shard ID is specified, the checkpoints of all shards are returned.
	//
	// example:
	//
	// 1
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
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       []*GetCheckPointResponseBody `json:"body,omitempty" xml:"body,omitempty" type:"Repeated"`
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
	// The shard ID.
	//
	// example:
	//
	// 0
	Shard *int32 `json:"shard,omitempty" xml:"shard,omitempty"`
	// The value of the checkpoint.
	//
	// example:
	//
	// MTUyNDE1NTM3OTM3MzkwODQ5Ng==
	Checkpoint *string `json:"checkpoint,omitempty" xml:"checkpoint,omitempty"`
	// The time when the checkpoint was last updated. The value is a UNIX timestamp representing the number of seconds that have elapsed since the epoch time January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1524224984800922
	UpdateTime *int64 `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
	// The consumer at the checkpoint.
	//
	// example:
	//
	// consumer_1
	Consumer *string `json:"consumer,omitempty" xml:"consumer,omitempty"`
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

type GetCollectionPolicyRequest struct {
	// example:
	//
	// access_log
	DataCode *string `json:"dataCode,omitempty" xml:"dataCode,omitempty"`
	// example:
	//
	// oss
	ProductCode *string `json:"productCode,omitempty" xml:"productCode,omitempty"`
}

func (s GetCollectionPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s GetCollectionPolicyRequest) GoString() string {
	return s.String()
}

func (s *GetCollectionPolicyRequest) SetDataCode(v string) *GetCollectionPolicyRequest {
	s.DataCode = &v
	return s
}

func (s *GetCollectionPolicyRequest) SetProductCode(v string) *GetCollectionPolicyRequest {
	s.ProductCode = &v
	return s
}

type GetCollectionPolicyResponseBody struct {
	CollectionPolicy *GetCollectionPolicyResponseBodyCollectionPolicy `json:"collectionPolicy,omitempty" xml:"collectionPolicy,omitempty" type:"Struct"`
}

func (s GetCollectionPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetCollectionPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *GetCollectionPolicyResponseBody) SetCollectionPolicy(v *GetCollectionPolicyResponseBodyCollectionPolicy) *GetCollectionPolicyResponseBody {
	s.CollectionPolicy = v
	return s
}

type GetCollectionPolicyResponseBodyCollectionPolicy struct {
	Attribute        *GetCollectionPolicyResponseBodyCollectionPolicyAttribute        `json:"attribute,omitempty" xml:"attribute,omitempty" type:"Struct"`
	CentralizeConfig *GetCollectionPolicyResponseBodyCollectionPolicyCentralizeConfig `json:"centralizeConfig,omitempty" xml:"centralizeConfig,omitempty" type:"Struct"`
	// example:
	//
	// false
	CentralizeEnabled *bool `json:"centralizeEnabled,omitempty" xml:"centralizeEnabled,omitempty"`
	// example:
	//
	// access_log
	DataCode *string `json:"dataCode,omitempty" xml:"dataCode,omitempty"`
	// example:
	//
	// true
	Enabled      *string                                                      `json:"enabled,omitempty" xml:"enabled,omitempty"`
	PolicyConfig *GetCollectionPolicyResponseBodyCollectionPolicyPolicyConfig `json:"policyConfig,omitempty" xml:"policyConfig,omitempty" type:"Struct"`
	// example:
	//
	// your_log_policy
	PolicyName *string `json:"policyName,omitempty" xml:"policyName,omitempty"`
	// example:
	//
	// oss
	ProductCode *string `json:"productCode,omitempty" xml:"productCode,omitempty"`
}

func (s GetCollectionPolicyResponseBodyCollectionPolicy) String() string {
	return tea.Prettify(s)
}

func (s GetCollectionPolicyResponseBodyCollectionPolicy) GoString() string {
	return s.String()
}

func (s *GetCollectionPolicyResponseBodyCollectionPolicy) SetAttribute(v *GetCollectionPolicyResponseBodyCollectionPolicyAttribute) *GetCollectionPolicyResponseBodyCollectionPolicy {
	s.Attribute = v
	return s
}

func (s *GetCollectionPolicyResponseBodyCollectionPolicy) SetCentralizeConfig(v *GetCollectionPolicyResponseBodyCollectionPolicyCentralizeConfig) *GetCollectionPolicyResponseBodyCollectionPolicy {
	s.CentralizeConfig = v
	return s
}

func (s *GetCollectionPolicyResponseBodyCollectionPolicy) SetCentralizeEnabled(v bool) *GetCollectionPolicyResponseBodyCollectionPolicy {
	s.CentralizeEnabled = &v
	return s
}

func (s *GetCollectionPolicyResponseBodyCollectionPolicy) SetDataCode(v string) *GetCollectionPolicyResponseBodyCollectionPolicy {
	s.DataCode = &v
	return s
}

func (s *GetCollectionPolicyResponseBodyCollectionPolicy) SetEnabled(v string) *GetCollectionPolicyResponseBodyCollectionPolicy {
	s.Enabled = &v
	return s
}

func (s *GetCollectionPolicyResponseBodyCollectionPolicy) SetPolicyConfig(v *GetCollectionPolicyResponseBodyCollectionPolicyPolicyConfig) *GetCollectionPolicyResponseBodyCollectionPolicy {
	s.PolicyConfig = v
	return s
}

func (s *GetCollectionPolicyResponseBodyCollectionPolicy) SetPolicyName(v string) *GetCollectionPolicyResponseBodyCollectionPolicy {
	s.PolicyName = &v
	return s
}

func (s *GetCollectionPolicyResponseBodyCollectionPolicy) SetProductCode(v string) *GetCollectionPolicyResponseBodyCollectionPolicy {
	s.ProductCode = &v
	return s
}

type GetCollectionPolicyResponseBodyCollectionPolicyAttribute struct {
	// example:
	//
	// your-app-name
	App *string `json:"app,omitempty" xml:"app,omitempty"`
	// example:
	//
	// your-policy-group
	PolicyGroup *string `json:"policyGroup,omitempty" xml:"policyGroup,omitempty"`
}

func (s GetCollectionPolicyResponseBodyCollectionPolicyAttribute) String() string {
	return tea.Prettify(s)
}

func (s GetCollectionPolicyResponseBodyCollectionPolicyAttribute) GoString() string {
	return s.String()
}

func (s *GetCollectionPolicyResponseBodyCollectionPolicyAttribute) SetApp(v string) *GetCollectionPolicyResponseBodyCollectionPolicyAttribute {
	s.App = &v
	return s
}

func (s *GetCollectionPolicyResponseBodyCollectionPolicyAttribute) SetPolicyGroup(v string) *GetCollectionPolicyResponseBodyCollectionPolicyAttribute {
	s.PolicyGroup = &v
	return s
}

type GetCollectionPolicyResponseBodyCollectionPolicyCentralizeConfig struct {
	// example:
	//
	// your-sls-logstore-in-beijing
	DestLogstore *string `json:"destLogstore,omitempty" xml:"destLogstore,omitempty"`
	// example:
	//
	// your-sls-project-in-beijing
	DestProject *string `json:"destProject,omitempty" xml:"destProject,omitempty"`
	// example:
	//
	// cn-beijing
	DestRegion *string `json:"destRegion,omitempty" xml:"destRegion,omitempty"`
	// example:
	//
	// your-sls-logstore-ttl
	DestTTL *int32 `json:"destTTL,omitempty" xml:"destTTL,omitempty"`
}

func (s GetCollectionPolicyResponseBodyCollectionPolicyCentralizeConfig) String() string {
	return tea.Prettify(s)
}

func (s GetCollectionPolicyResponseBodyCollectionPolicyCentralizeConfig) GoString() string {
	return s.String()
}

func (s *GetCollectionPolicyResponseBodyCollectionPolicyCentralizeConfig) SetDestLogstore(v string) *GetCollectionPolicyResponseBodyCollectionPolicyCentralizeConfig {
	s.DestLogstore = &v
	return s
}

func (s *GetCollectionPolicyResponseBodyCollectionPolicyCentralizeConfig) SetDestProject(v string) *GetCollectionPolicyResponseBodyCollectionPolicyCentralizeConfig {
	s.DestProject = &v
	return s
}

func (s *GetCollectionPolicyResponseBodyCollectionPolicyCentralizeConfig) SetDestRegion(v string) *GetCollectionPolicyResponseBodyCollectionPolicyCentralizeConfig {
	s.DestRegion = &v
	return s
}

func (s *GetCollectionPolicyResponseBodyCollectionPolicyCentralizeConfig) SetDestTTL(v int32) *GetCollectionPolicyResponseBodyCollectionPolicyCentralizeConfig {
	s.DestTTL = &v
	return s
}

type GetCollectionPolicyResponseBodyCollectionPolicyPolicyConfig struct {
	InstanceIds []*string `json:"instanceIds,omitempty" xml:"instanceIds,omitempty" type:"Repeated"`
	Regions     []*string `json:"regions,omitempty" xml:"regions,omitempty" type:"Repeated"`
	// example:
	//
	// all
	ResourceMode *string                `json:"resourceMode,omitempty" xml:"resourceMode,omitempty"`
	ResourceTags map[string]interface{} `json:"resourceTags,omitempty" xml:"resourceTags,omitempty"`
}

func (s GetCollectionPolicyResponseBodyCollectionPolicyPolicyConfig) String() string {
	return tea.Prettify(s)
}

func (s GetCollectionPolicyResponseBodyCollectionPolicyPolicyConfig) GoString() string {
	return s.String()
}

func (s *GetCollectionPolicyResponseBodyCollectionPolicyPolicyConfig) SetInstanceIds(v []*string) *GetCollectionPolicyResponseBodyCollectionPolicyPolicyConfig {
	s.InstanceIds = v
	return s
}

func (s *GetCollectionPolicyResponseBodyCollectionPolicyPolicyConfig) SetRegions(v []*string) *GetCollectionPolicyResponseBodyCollectionPolicyPolicyConfig {
	s.Regions = v
	return s
}

func (s *GetCollectionPolicyResponseBodyCollectionPolicyPolicyConfig) SetResourceMode(v string) *GetCollectionPolicyResponseBodyCollectionPolicyPolicyConfig {
	s.ResourceMode = &v
	return s
}

func (s *GetCollectionPolicyResponseBodyCollectionPolicyPolicyConfig) SetResourceTags(v map[string]interface{}) *GetCollectionPolicyResponseBodyCollectionPolicyPolicyConfig {
	s.ResourceTags = v
	return s
}

type GetCollectionPolicyResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCollectionPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCollectionPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s GetCollectionPolicyResponse) GoString() string {
	return s.String()
}

func (s *GetCollectionPolicyResponse) SetHeaders(v map[string]*string) *GetCollectionPolicyResponse {
	s.Headers = v
	return s
}

func (s *GetCollectionPolicyResponse) SetStatusCode(v int32) *GetCollectionPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCollectionPolicyResponse) SetBody(v *GetCollectionPolicyResponseBody) *GetCollectionPolicyResponse {
	s.Body = v
	return s
}

type GetConfigResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *LogtailConfig     `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s GetConfigResponse) GoString() string {
	return s.String()
}

func (s *GetConfigResponse) SetHeaders(v map[string]*string) *GetConfigResponse {
	s.Headers = v
	return s
}

func (s *GetConfigResponse) SetStatusCode(v int32) *GetConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *GetConfigResponse) SetBody(v *LogtailConfig) *GetConfigResponse {
	s.Body = v
	return s
}

type GetContextLogsRequest struct {
	// The number of logs that you want to obtain and are generated before the generation time of the start log. Valid values: (0,100].
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	BackLines *int64 `json:"back_lines,omitempty" xml:"back_lines,omitempty"`
	// The number of logs that you want to obtain and are generated after the generation time of the start log. Valid values: (0,100].
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	ForwardLines *int64 `json:"forward_lines,omitempty" xml:"forward_lines,omitempty"`
	// The unique identifier of the log group to which the start log belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// 85C897C740352DC6-808
	PackId *string `json:"pack_id,omitempty" xml:"pack_id,omitempty"`
	// The unique context identifier of the start log in the log group.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2|MTY1NTcwNTUzODY5MTY0MDk1Mg==|3|0
	PackMeta *string `json:"pack_meta,omitempty" xml:"pack_meta,omitempty"`
	// The type of the data in the Logstore. Set the value to context_log.
	//
	// This parameter is required.
	//
	// example:
	//
	// context_log
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
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
	// The number of logs that are generated before the generation time of the start log.
	//
	// example:
	//
	// 100
	BackLines *int64 `json:"back_lines,omitempty" xml:"back_lines,omitempty"`
	// The number of logs that are generated after the generation time of the start log.
	//
	// example:
	//
	// 100
	ForwardLines *int64 `json:"forward_lines,omitempty" xml:"forward_lines,omitempty"`
	// The logs that are returned.
	Logs []map[string]interface{} `json:"logs,omitempty" xml:"logs,omitempty" type:"Repeated"`
	// Indicates whether the query and analysis results are complete. Valid values:
	//
	// 	- Complete: The query is successful, and the complete query and analysis results are returned.
	//
	// 	- Incomplete: The query is successful, but the query and analysis results are incomplete. To obtain the complete results, you must repeat the request.
	//
	// example:
	//
	// Complete
	Progress *string `json:"progress,omitempty" xml:"progress,omitempty"`
	// The total number of logs that are returned. The logs include the start log that is specified in the request.
	//
	// example:
	//
	// 201
	TotalLines *int64 `json:"total_lines,omitempty" xml:"total_lines,omitempty"`
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
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetContextLogsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The point in time that you want to use to query a cursor. Set the value to a UNIX timestamp or a string such as `begin` and `end`.
	//
	// This parameter is required.
	//
	// example:
	//
	// begin
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
	// The value of the cursor.
	//
	// example:
	//
	// MTQ0NzI5OTYwNjg5NjYzMjM1Ng==
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
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCursorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The cursor.
	//
	// This parameter is required.
	//
	// example:
	//
	// MTU0NzQ3MDY4MjM3NjUxMzQ0Ng==
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
	// The server time that is queried based on the cursor. The value is a UNIX timestamp representing the number of seconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1554260243
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
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCursorTimeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type GetDashboardResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *Dashboard         `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDashboardResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDashboardResponse) GoString() string {
	return s.String()
}

func (s *GetDashboardResponse) SetHeaders(v map[string]*string) *GetDashboardResponse {
	s.Headers = v
	return s
}

func (s *GetDashboardResponse) SetStatusCode(v int32) *GetDashboardResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDashboardResponse) SetBody(v *Dashboard) *GetDashboardResponse {
	s.Body = v
	return s
}

type GetETLResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ETL               `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetETLResponse) String() string {
	return tea.Prettify(s)
}

func (s GetETLResponse) GoString() string {
	return s.String()
}

func (s *GetETLResponse) SetHeaders(v map[string]*string) *GetETLResponse {
	s.Headers = v
	return s
}

func (s *GetETLResponse) SetStatusCode(v int32) *GetETLResponse {
	s.StatusCode = &v
	return s
}

func (s *GetETLResponse) SetBody(v *ETL) *GetETLResponse {
	s.Body = v
	return s
}

type GetExternalStoreResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ExternalStore     `json:"body,omitempty" xml:"body,omitempty"`
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
	// The start time of the subinterval. The value is a UNIX timestamp representing the number of seconds that have elapsed since the epoch time January 1, 1970, 00:00:00 UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1409529600
	From *int64 `json:"from,omitempty" xml:"from,omitempty"`
	// The search statement. Only search statements are supported. Analytic statements are not supported. For more information about the syntax of search statements, see [Log search overview](https://help.aliyun.com/document_detail/43772.html).
	//
	// example:
	//
	// with_pack_meta
	Query *string `json:"query,omitempty" xml:"query,omitempty"`
	// The end time of the subinterval. The value is a UNIX timestamp representing the number of seconds that have elapsed since the epoch time January 1, 1970, 00:00:00 UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1409569200
	To *int64 `json:"to,omitempty" xml:"to,omitempty"`
	// The topic of the logs.
	//
	// example:
	//
	// topic
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
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       []*GetHistogramsResponseBody `json:"body,omitempty" xml:"body,omitempty" type:"Repeated"`
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
	// The start time of the subinterval. The value is a UNIX timestamp representing the number of seconds that have elapsed since the epoch time January 1, 1970, 00:00:00 UTC.
	//
	// The time range that is specified in this operation is a left-closed, right-open interval. The interval includes the start time specified by the from parameter, but does not include the end time specified by the to parameter. If you specify the same value for the from and to parameters, the interval is invalid, and an error message is returned.
	//
	// example:
	//
	// 1409529600
	From *int64 `json:"from,omitempty" xml:"from,omitempty"`
	// The end time of the subinterval. The value is a UNIX timestamp representing the number of seconds that have elapsed since the epoch time January 1, 1970, 00:00:00 UTC.
	//
	// The time range that is specified in this operation is a left-closed, right-open interval. The interval includes the start time specified by the from parameter, but does not include the end time specified by the to parameter. If you specify the same value for the from and to parameters, the interval is invalid, and an error message is returned.
	//
	// example:
	//
	// 1409569200
	To *int64 `json:"to,omitempty" xml:"to,omitempty"`
	// The number of logs that are generated within the subinterval.
	//
	// example:
	//
	// 2
	Count *int64 `json:"count,omitempty" xml:"count,omitempty"`
	// Indicates whether the query and analysis results in the subinterval is complete. Valid values:
	//
	// Complete: The query is successful, and the complete query and analysis results are returned.
	//
	// Incomplete: The query is successful, but the query and analysis results are incomplete. To obtain the complete results, you must repeat the request.
	//
	// example:
	//
	// Complete
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
	// The type of the index.
	//
	// example:
	//
	// v2
	IndexMode *string `json:"index_mode,omitempty" xml:"index_mode,omitempty"`
	// The configurations of field indexes. A field index is in the key-value format in which the key specifies the name of the field and the value specifies the index configuration of the field.
	Keys map[string]*KeysValue `json:"keys,omitempty" xml:"keys,omitempty"`
	// The time when the index configurations were last updated. The value is a UNIX timestamp representing the number of seconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1524155379
	LastModifyTime *int64 `json:"lastModifyTime,omitempty" xml:"lastModifyTime,omitempty"`
	// The configurations of full-text indexes.
	Line *GetIndexResponseBodyLine `json:"line,omitempty" xml:"line,omitempty" type:"Struct"`
	// Indicates whether the log clustering feature is enabled.
	//
	// example:
	//
	// false
	LogReduce *bool `json:"log_reduce,omitempty" xml:"log_reduce,omitempty"`
	// The fields in the blacklist that are used to cluster logs. This parameter is valid only if the log clustering feature is enabled.
	LogReduceBlackList []*string `json:"log_reduce_black_list,omitempty" xml:"log_reduce_black_list,omitempty" type:"Repeated"`
	// The fields in the whitelist that are used to cluster logs. This parameter is valid only if the log clustering feature is enabled.
	LogReduceWhiteList []*string `json:"log_reduce_white_list,omitempty" xml:"log_reduce_white_list,omitempty" type:"Repeated"`
	// The maximum length of a field value that can be retained. Default value: 2048. Unit: bytes. The default value is equal to 2 KB. You can change the value of the max_text_len parameter. Valid values: 64 to 16384. Unit: bytes.
	//
	// example:
	//
	// 2048
	MaxTextLen *int32 `json:"max_text_len,omitempty" xml:"max_text_len,omitempty"`
	// The storage type. The value is fixed as pg.
	//
	// example:
	//
	// pg
	Storage *string `json:"storage,omitempty" xml:"storage,omitempty"`
	// The lifecycle of the index file. Valid values: 7, 30, and 90. Unit: day.
	//
	// example:
	//
	// 30
	Ttl *int32 `json:"ttl,omitempty" xml:"ttl,omitempty"`
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
	// Indicates whether case sensitivity is enabled. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// false
	CaseSensitive *bool `json:"caseSensitive,omitempty" xml:"caseSensitive,omitempty"`
	// Indicates whether Chinese characters are included. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// false
	Chn *bool `json:"chn,omitempty" xml:"chn,omitempty"`
	// The excluded fields.
	ExcludeKeys []*string `json:"exclude_keys,omitempty" xml:"exclude_keys,omitempty" type:"Repeated"`
	// The included fields.
	IncludeKeys []*string `json:"include_keys,omitempty" xml:"include_keys,omitempty" type:"Repeated"`
	// The delimiters.
	Token []*string `json:"token,omitempty" xml:"token,omitempty" type:"Repeated"`
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
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetIndexResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type GetLogStoreResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *Logstore          `json:"body,omitempty" xml:"body,omitempty"`
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

type GetLogStoreMeteringModeResponseBody struct {
	// example:
	//
	// ChargeByFunction
	MeteringMode *string `json:"meteringMode,omitempty" xml:"meteringMode,omitempty"`
}

func (s GetLogStoreMeteringModeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetLogStoreMeteringModeResponseBody) GoString() string {
	return s.String()
}

func (s *GetLogStoreMeteringModeResponseBody) SetMeteringMode(v string) *GetLogStoreMeteringModeResponseBody {
	s.MeteringMode = &v
	return s
}

type GetLogStoreMeteringModeResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetLogStoreMeteringModeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetLogStoreMeteringModeResponse) String() string {
	return tea.Prettify(s)
}

func (s GetLogStoreMeteringModeResponse) GoString() string {
	return s.String()
}

func (s *GetLogStoreMeteringModeResponse) SetHeaders(v map[string]*string) *GetLogStoreMeteringModeResponse {
	s.Headers = v
	return s
}

func (s *GetLogStoreMeteringModeResponse) SetStatusCode(v int32) *GetLogStoreMeteringModeResponse {
	s.StatusCode = &v
	return s
}

func (s *GetLogStoreMeteringModeResponse) SetBody(v *GetLogStoreMeteringModeResponseBody) *GetLogStoreMeteringModeResponse {
	s.Body = v
	return s
}

type GetLoggingResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *Logging           `json:"body,omitempty" xml:"body,omitempty"`
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
	// The beginning of the time range to query. The value is the log time that is specified when log data is written.
	//
	// 	- The time range that is specified in this operation is a left-closed, right-open interval. The interval includes the start time specified by the **from*	- parameter, but does not include the end time specified by the **to*	- parameter. If you specify the same value for the **from*	- and **to*	- parameters, the interval is invalid, and an error message is returned.
	//
	// 	- The value is a UNIX timestamp representing the number of seconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// > To ensure that full data can be queried, specify a query time range that is accurate to the minute. If you also specify a time range in an analytic statement, Simple Log Service uses the time range specified in the analytic statement for query and analysis.
	//
	// If you want to specify a time range that is accurate to the second in your analytic statement, you must use the from_unixtime or to_unixtime function to convert the time format. For more information about the functions, see [from_unixtime function](https://help.aliyun.com/document_detail/63451.html) and [to_unixtime function](https://help.aliyun.com/document_detail/63451.html). Examples:
	//
	// 	- `	- | SELECT 	- FROM log WHERE from_unixtime(__time__) > from_unixtime(1664186624) AND from_unixtime(__time__) < now()`
	//
	// 	- `	- | SELECT 	- FROM log WHERE __time__ > to_unixtime(date_parse(\\"2022-10-19 15:46:05\\", \\"%Y-%m-%d %H:%i:%s\\")) AND __time__ < to_unixtime(now())`
	//
	// This parameter is required.
	//
	// example:
	//
	// 1627268185
	From *int32 `json:"from,omitempty" xml:"from,omitempty"`
	// The maximum number of logs to return for the request. This parameter takes effect only when the query parameter is set to a search statement. Minimum value: 0. Maximum value: 100. Default value: 100.
	//
	// example:
	//
	// 100
	Line *int64 `json:"line,omitempty" xml:"line,omitempty"`
	// The line from which the query starts. This parameter takes effect only when the query parameter is set to a search statement. Default value: 0.
	//
	// example:
	//
	// 0
	Offset *int64 `json:"offset,omitempty" xml:"offset,omitempty"`
	// Specifies whether to enable the Dedicated SQL feature. For more information, see [Enable Dedicated SQL](https://help.aliyun.com/document_detail/223777.html). Valid values:
	//
	// 	- true: enables the Dedicated SQL feature.
	//
	// 	- false (default): enables the Standard SQL feature.
	//
	// You can use the powerSql or **query*	- parameter to configure Dedicated SQL.
	//
	// example:
	//
	// false
	PowerSql *bool `json:"powerSql,omitempty" xml:"powerSql,omitempty"`
	// The search statement or the query statement. For more information, see [Log search overview](https://help.aliyun.com/document_detail/43772.html) and [Log analysis overview](https://help.aliyun.com/document_detail/53608.html). If you add `set session parallel_sql=true;` to the analytic statement in the query parameter, Dedicated SQL is used. For example, you can set the query parameter to `	- | set session parallel_sql=true; select count(*) as pv`. For more information about common errors that may occur during log query and analysis, see [How do I resolve common errors that occur when I query and analyze logs?](https://help.aliyun.com/document_detail/61628.html)
	//
	// > If you specify an analytic statement in the value of the query parameter, the line and offset parameters do not take effect. In this case, we recommend that you set the line and offset parameters to 0 and use the LIMIT clause to limit the number of logs to return on each page. For more information, see [Paged query](https://help.aliyun.com/document_detail/89994.html).
	//
	// example:
	//
	// status: 401 | SELECT remote_addr,COUNT(*) as pv GROUP by remote_addr ORDER by pv desc limit 5
	Query *string `json:"query,omitempty" xml:"query,omitempty"`
	// Specifies whether to return logs in reverse chronological order of log timestamps. The log timestamps are accurate to the minute. Valid values:
	//
	// 	- true: returns logs in reverse chronological order of log timestamps.
	//
	// 	- false (default): returns logs in chronological order of log timestamps.
	//
	// >
	//
	// 	- The reverse parameter takes effect only when the query parameter is set to a search statement. The reverse parameter specifies the method used to sort returned logs.
	//
	// 	- If the query parameter is set to a query statement, the reverse parameter does not take effect. The method used to sort returned logs is specified by the ORDER BY clause in the analytic statement. If you use the keyword asc in the ORDER BY clause, the logs are sorted in chronological order. If you use the keyword desc in the ORDER BY clause, the logs are sorted in reverse chronological order. By default, asc is used in the ORDER BY clause.
	//
	// example:
	//
	// false
	Reverse *bool `json:"reverse,omitempty" xml:"reverse,omitempty"`
	// The end of the time range to query. The value is the log time that is specified when log data is written.
	//
	// 	- The time range that is specified in this operation is a left-closed, right-open interval. The interval includes the start time specified by the **from*	- parameter, but does not include the end time specified by the **to*	- parameter. If you specify the same value for the **from*	- and **to*	- parameters, the interval is invalid, and an error message is returned.
	//
	// 	- The value is a UNIX timestamp representing the number of seconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// > To ensure that full data can be queried, specify a query time range that is accurate to the minute. If you also specify a time range in an analytic statement, Simple Log Service uses the time range specified in the analytic statement for query and analysis.
	//
	// If you want to specify a time range that is accurate to the second in your analytic statement, you must use the from_unixtime or to_unixtime function to convert the time format. For more information about the functions, see [from_unixtime function](https://help.aliyun.com/document_detail/63451.html) and [to_unixtime function](https://help.aliyun.com/document_detail/63451.html). Examples:
	//
	// 	- `	- | SELECT 	- FROM log WHERE from_unixtime(__time__) > from_unixtime(1664186624) AND from_unixtime(__time__) < now()`
	//
	// 	- `	- | SELECT 	- FROM log WHERE __time__ > to_unixtime(date_parse(\\"2022-10-19 15:46:05\\", \\"%Y-%m-%d %H:%i:%s\\")) AND __time__ < to_unixtime(now())`
	//
	// This parameter is required.
	//
	// example:
	//
	// 1627269085
	To *int32 `json:"to,omitempty" xml:"to,omitempty"`
	// The topic of the logs. The default value is double quotation marks (""). For more information, see [Topic](https://help.aliyun.com/document_detail/48881.html).
	//
	// example:
	//
	// topic
	Topic *string `json:"topic,omitempty" xml:"topic,omitempty"`
}

func (s GetLogsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetLogsRequest) GoString() string {
	return s.String()
}

func (s *GetLogsRequest) SetFrom(v int32) *GetLogsRequest {
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

func (s *GetLogsRequest) SetTo(v int32) *GetLogsRequest {
	s.To = &v
	return s
}

func (s *GetLogsRequest) SetTopic(v string) *GetLogsRequest {
	s.Topic = &v
	return s
}

type GetLogsResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       []map[string]interface{} `json:"body,omitempty" xml:"body,omitempty" type:"Repeated"`
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

type GetLogsV2Headers struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// The compression method.
	//
	// This parameter is required.
	//
	// example:
	//
	// lz4
	AcceptEncoding *string `json:"Accept-Encoding,omitempty" xml:"Accept-Encoding,omitempty"`
}

func (s GetLogsV2Headers) String() string {
	return tea.Prettify(s)
}

func (s GetLogsV2Headers) GoString() string {
	return s.String()
}

func (s *GetLogsV2Headers) SetCommonHeaders(v map[string]*string) *GetLogsV2Headers {
	s.CommonHeaders = v
	return s
}

func (s *GetLogsV2Headers) SetAcceptEncoding(v string) *GetLogsV2Headers {
	s.AcceptEncoding = &v
	return s
}

type GetLogsV2Request struct {
	// Specifies whether to page forward or backward for the scan-based query or the phrase search.
	//
	// example:
	//
	// false
	Forward *bool `json:"forward,omitempty" xml:"forward,omitempty"`
	// The beginning of the time range to query. The value is the log time that is specified when log data is written.
	//
	// The time range that is specified in this operation is a left-closed, right-open interval. The interval includes the start time specified by the from parameter, but does not include the end time specified by the to parameter. If you specify the same value for the from and to parameters, the interval is invalid, and an error message is returned. The value is a UNIX timestamp representing the number of seconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1627268185
	From      *int32 `json:"from,omitempty" xml:"from,omitempty"`
	Highlight *bool  `json:"highlight,omitempty" xml:"highlight,omitempty"`
	// The maximum number of logs to return for the request. This parameter takes effect only when the query parameter is set to a search statement. Minimum value: 0. Maximum value: 100. Default value: 100.
	//
	// example:
	//
	// 100
	Line *int64 `json:"line,omitempty" xml:"line,omitempty"`
	// The line from which the query starts. This parameter takes effect only when the query parameter is set to a search statement. Default value: 0.
	//
	// example:
	//
	// 0
	Offset *int64 `json:"offset,omitempty" xml:"offset,omitempty"`
	// Specifies whether to enable the SQL enhancement feature. By default, the feature is disabled.
	//
	// example:
	//
	// false
	PowerSql *bool `json:"powerSql,omitempty" xml:"powerSql,omitempty"`
	// The search statement or the query statement. For more information, see the "Log search overview" and "Log analysis overview" topics.
	//
	// If you add set session parallel_sql=true; to the analytic statement in the query parameter, Dedicated SQL is used. For example, you can set the query parameter to \\	- | set session parallel_sql=true; select count(\\*) as pv.
	//
	// Note: If you specify an analytic statement in the query parameter, the line and offset parameters do not take effect in this operation. In this case, we recommend that you set the line and offset parameters to 0 and use the LIMIT clause to limit the number of logs to return on each page. For more information, see the "Perform paged queries" topic.
	//
	// example:
	//
	// status: 401 | SELECT remote_addr,COUNT(*) as pv GROUP by remote_addr ORDER by pv desc limit 5
	Query *string `json:"query,omitempty" xml:"query,omitempty"`
	// Specifies whether to return logs in reverse chronological order of log timestamps. The log timestamps are accurate to the minute. Valid values:
	//
	// true: Logs are returned in reverse chronological order of log timestamps. false (default): Logs are returned in chronological order of log timestamps. Note: The reverse parameter takes effect only when the query parameter is set to a search statement. The reverse parameter specifies the method used to sort returned logs. If the query parameter is set to a query statement, the reverse parameter does not take effect. The method used to sort returned logs is specified by the ORDER BY clause in the analytic statement. If you use the keyword asc in the ORDER BY clause, the logs are sorted in chronological order. If you use the keyword desc in the ORDER BY clause, the logs are sorted in reverse chronological order. By default, asc is used in the ORDER BY clause.
	//
	// example:
	//
	// false
	Reverse *bool `json:"reverse,omitempty" xml:"reverse,omitempty"`
	// The parameter that is used to query data.
	//
	// example:
	//
	// mode=scan
	Session *string `json:"session,omitempty" xml:"session,omitempty"`
	// The end of the time range to query. The value is the log time that is specified when log data is written.
	//
	// The time range that is specified in this operation is a left-closed, right-open interval. The interval includes the start time specified by the from parameter, but does not include the end time specified by the to parameter. If you specify the same value for the from and to parameters, the interval is invalid, and an error message is returned. The value is a UNIX timestamp representing the number of seconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1627268185
	To *int32 `json:"to,omitempty" xml:"to,omitempty"`
	// The topic of the logs. Default value: double quotation marks ("").
	//
	// example:
	//
	// ""
	Topic *string `json:"topic,omitempty" xml:"topic,omitempty"`
}

func (s GetLogsV2Request) String() string {
	return tea.Prettify(s)
}

func (s GetLogsV2Request) GoString() string {
	return s.String()
}

func (s *GetLogsV2Request) SetForward(v bool) *GetLogsV2Request {
	s.Forward = &v
	return s
}

func (s *GetLogsV2Request) SetFrom(v int32) *GetLogsV2Request {
	s.From = &v
	return s
}

func (s *GetLogsV2Request) SetHighlight(v bool) *GetLogsV2Request {
	s.Highlight = &v
	return s
}

func (s *GetLogsV2Request) SetLine(v int64) *GetLogsV2Request {
	s.Line = &v
	return s
}

func (s *GetLogsV2Request) SetOffset(v int64) *GetLogsV2Request {
	s.Offset = &v
	return s
}

func (s *GetLogsV2Request) SetPowerSql(v bool) *GetLogsV2Request {
	s.PowerSql = &v
	return s
}

func (s *GetLogsV2Request) SetQuery(v string) *GetLogsV2Request {
	s.Query = &v
	return s
}

func (s *GetLogsV2Request) SetReverse(v bool) *GetLogsV2Request {
	s.Reverse = &v
	return s
}

func (s *GetLogsV2Request) SetSession(v string) *GetLogsV2Request {
	s.Session = &v
	return s
}

func (s *GetLogsV2Request) SetTo(v int32) *GetLogsV2Request {
	s.To = &v
	return s
}

func (s *GetLogsV2Request) SetTopic(v string) *GetLogsV2Request {
	s.Topic = &v
	return s
}

type GetLogsV2ResponseBody struct {
	// The returned result.
	Data []map[string]*string `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// The metadata that is returned.
	Meta *GetLogsV2ResponseBodyMeta `json:"meta,omitempty" xml:"meta,omitempty" type:"Struct"`
}

func (s GetLogsV2ResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetLogsV2ResponseBody) GoString() string {
	return s.String()
}

func (s *GetLogsV2ResponseBody) SetData(v []map[string]*string) *GetLogsV2ResponseBody {
	s.Data = v
	return s
}

func (s *GetLogsV2ResponseBody) SetMeta(v *GetLogsV2ResponseBodyMeta) *GetLogsV2ResponseBody {
	s.Meta = v
	return s
}

type GetLogsV2ResponseBodyMeta struct {
	// The SQL statement after | in the query statement.
	//
	// example:
	//
	// select *
	AggQuery    *string   `json:"aggQuery,omitempty" xml:"aggQuery,omitempty"`
	ColumnTypes []*string `json:"columnTypes,omitempty" xml:"columnTypes,omitempty" type:"Repeated"`
	// The number of rows that are returned.
	//
	// example:
	//
	// 1
	Count    *int32   `json:"count,omitempty" xml:"count,omitempty"`
	CpuCores *int32   `json:"cpuCores,omitempty" xml:"cpuCores,omitempty"`
	CpuSec   *float64 `json:"cpuSec,omitempty" xml:"cpuSec,omitempty"`
	// The amount of time that is consumed by the request. Unit: milliseconds.
	//
	// example:
	//
	// 5
	ElapsedMillisecond *int64 `json:"elapsedMillisecond,omitempty" xml:"elapsedMillisecond,omitempty"`
	// Indicates whether the query is an SQL query.
	//
	// example:
	//
	// false
	HasSQL     *bool           `json:"hasSQL,omitempty" xml:"hasSQL,omitempty"`
	Highlights [][]*LogContent `json:"highlights,omitempty" xml:"highlights,omitempty" type:"Repeated"`
	// Indicates whether the returned result is accurate.
	//
	// example:
	//
	// true
	IsAccurate *bool `json:"isAccurate,omitempty" xml:"isAccurate,omitempty"`
	// All keys in the query result.
	Keys            []*string                                 `json:"keys,omitempty" xml:"keys,omitempty" type:"Repeated"`
	Limited         *int32                                    `json:"limited,omitempty" xml:"limited,omitempty"`
	Mode            *int32                                    `json:"mode,omitempty" xml:"mode,omitempty"`
	PhraseQueryInfo *GetLogsV2ResponseBodyMetaPhraseQueryInfo `json:"phraseQueryInfo,omitempty" xml:"phraseQueryInfo,omitempty" type:"Struct"`
	// The number of logs that are processed in the request.
	//
	// example:
	//
	// 10000
	ProcessedBytes *int64 `json:"processedBytes,omitempty" xml:"processedBytes,omitempty"`
	// The number of rows that are processed in the request.
	//
	// example:
	//
	// 10000
	ProcessedRows *int32 `json:"processedRows,omitempty" xml:"processedRows,omitempty"`
	// Indicates whether the query result is complete. Valid values:
	//
	// 	- Complete: The query was successful, and the complete result is returned.
	//
	// 	- Incomplete: The query was successful, but the query result is incomplete. To obtain the complete result, you must call the operation again.
	//
	// example:
	//
	// Complete
	Progress  *string `json:"progress,omitempty" xml:"progress,omitempty"`
	ScanBytes *int64  `json:"scanBytes,omitempty" xml:"scanBytes,omitempty"`
	// The type of observable data.
	//
	// example:
	//
	// None
	TelementryType *string `json:"telementryType,omitempty" xml:"telementryType,omitempty"`
	// All terms in the query statement.
	Terms []map[string]interface{} `json:"terms,omitempty" xml:"terms,omitempty" type:"Repeated"`
	// The part before | in the query statement.
	//
	// example:
	//
	// *
	WhereQuery *string `json:"whereQuery,omitempty" xml:"whereQuery,omitempty"`
}

func (s GetLogsV2ResponseBodyMeta) String() string {
	return tea.Prettify(s)
}

func (s GetLogsV2ResponseBodyMeta) GoString() string {
	return s.String()
}

func (s *GetLogsV2ResponseBodyMeta) SetAggQuery(v string) *GetLogsV2ResponseBodyMeta {
	s.AggQuery = &v
	return s
}

func (s *GetLogsV2ResponseBodyMeta) SetColumnTypes(v []*string) *GetLogsV2ResponseBodyMeta {
	s.ColumnTypes = v
	return s
}

func (s *GetLogsV2ResponseBodyMeta) SetCount(v int32) *GetLogsV2ResponseBodyMeta {
	s.Count = &v
	return s
}

func (s *GetLogsV2ResponseBodyMeta) SetCpuCores(v int32) *GetLogsV2ResponseBodyMeta {
	s.CpuCores = &v
	return s
}

func (s *GetLogsV2ResponseBodyMeta) SetCpuSec(v float64) *GetLogsV2ResponseBodyMeta {
	s.CpuSec = &v
	return s
}

func (s *GetLogsV2ResponseBodyMeta) SetElapsedMillisecond(v int64) *GetLogsV2ResponseBodyMeta {
	s.ElapsedMillisecond = &v
	return s
}

func (s *GetLogsV2ResponseBodyMeta) SetHasSQL(v bool) *GetLogsV2ResponseBodyMeta {
	s.HasSQL = &v
	return s
}

func (s *GetLogsV2ResponseBodyMeta) SetHighlights(v [][]*LogContent) *GetLogsV2ResponseBodyMeta {
	s.Highlights = v
	return s
}

func (s *GetLogsV2ResponseBodyMeta) SetIsAccurate(v bool) *GetLogsV2ResponseBodyMeta {
	s.IsAccurate = &v
	return s
}

func (s *GetLogsV2ResponseBodyMeta) SetKeys(v []*string) *GetLogsV2ResponseBodyMeta {
	s.Keys = v
	return s
}

func (s *GetLogsV2ResponseBodyMeta) SetLimited(v int32) *GetLogsV2ResponseBodyMeta {
	s.Limited = &v
	return s
}

func (s *GetLogsV2ResponseBodyMeta) SetMode(v int32) *GetLogsV2ResponseBodyMeta {
	s.Mode = &v
	return s
}

func (s *GetLogsV2ResponseBodyMeta) SetPhraseQueryInfo(v *GetLogsV2ResponseBodyMetaPhraseQueryInfo) *GetLogsV2ResponseBodyMeta {
	s.PhraseQueryInfo = v
	return s
}

func (s *GetLogsV2ResponseBodyMeta) SetProcessedBytes(v int64) *GetLogsV2ResponseBodyMeta {
	s.ProcessedBytes = &v
	return s
}

func (s *GetLogsV2ResponseBodyMeta) SetProcessedRows(v int32) *GetLogsV2ResponseBodyMeta {
	s.ProcessedRows = &v
	return s
}

func (s *GetLogsV2ResponseBodyMeta) SetProgress(v string) *GetLogsV2ResponseBodyMeta {
	s.Progress = &v
	return s
}

func (s *GetLogsV2ResponseBodyMeta) SetScanBytes(v int64) *GetLogsV2ResponseBodyMeta {
	s.ScanBytes = &v
	return s
}

func (s *GetLogsV2ResponseBodyMeta) SetTelementryType(v string) *GetLogsV2ResponseBodyMeta {
	s.TelementryType = &v
	return s
}

func (s *GetLogsV2ResponseBodyMeta) SetTerms(v []map[string]interface{}) *GetLogsV2ResponseBodyMeta {
	s.Terms = v
	return s
}

func (s *GetLogsV2ResponseBodyMeta) SetWhereQuery(v string) *GetLogsV2ResponseBodyMeta {
	s.WhereQuery = &v
	return s
}

type GetLogsV2ResponseBodyMetaPhraseQueryInfo struct {
	BeginOffset *int64 `json:"beginOffset,omitempty" xml:"beginOffset,omitempty"`
	EndOffset   *int64 `json:"endOffset,omitempty" xml:"endOffset,omitempty"`
	EndTime     *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	ScanAll     *bool  `json:"scanAll,omitempty" xml:"scanAll,omitempty"`
}

func (s GetLogsV2ResponseBodyMetaPhraseQueryInfo) String() string {
	return tea.Prettify(s)
}

func (s GetLogsV2ResponseBodyMetaPhraseQueryInfo) GoString() string {
	return s.String()
}

func (s *GetLogsV2ResponseBodyMetaPhraseQueryInfo) SetBeginOffset(v int64) *GetLogsV2ResponseBodyMetaPhraseQueryInfo {
	s.BeginOffset = &v
	return s
}

func (s *GetLogsV2ResponseBodyMetaPhraseQueryInfo) SetEndOffset(v int64) *GetLogsV2ResponseBodyMetaPhraseQueryInfo {
	s.EndOffset = &v
	return s
}

func (s *GetLogsV2ResponseBodyMetaPhraseQueryInfo) SetEndTime(v int64) *GetLogsV2ResponseBodyMetaPhraseQueryInfo {
	s.EndTime = &v
	return s
}

func (s *GetLogsV2ResponseBodyMetaPhraseQueryInfo) SetScanAll(v bool) *GetLogsV2ResponseBodyMetaPhraseQueryInfo {
	s.ScanAll = &v
	return s
}

type GetLogsV2Response struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetLogsV2ResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetLogsV2Response) String() string {
	return tea.Prettify(s)
}

func (s GetLogsV2Response) GoString() string {
	return s.String()
}

func (s *GetLogsV2Response) SetHeaders(v map[string]*string) *GetLogsV2Response {
	s.Headers = v
	return s
}

func (s *GetLogsV2Response) SetStatusCode(v int32) *GetLogsV2Response {
	s.StatusCode = &v
	return s
}

func (s *GetLogsV2Response) SetBody(v *GetLogsV2ResponseBody) *GetLogsV2Response {
	s.Body = v
	return s
}

type GetLogtailPipelineConfigResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *LogtailPipelineConfig `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetLogtailPipelineConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s GetLogtailPipelineConfigResponse) GoString() string {
	return s.String()
}

func (s *GetLogtailPipelineConfigResponse) SetHeaders(v map[string]*string) *GetLogtailPipelineConfigResponse {
	s.Headers = v
	return s
}

func (s *GetLogtailPipelineConfigResponse) SetStatusCode(v int32) *GetLogtailPipelineConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *GetLogtailPipelineConfigResponse) SetBody(v *LogtailPipelineConfig) *GetLogtailPipelineConfigResponse {
	s.Body = v
	return s
}

type GetMLServiceResultsRequest struct {
	// example:
	//
	// true
	AllowBuiltin *bool                   `json:"allowBuiltin,omitempty" xml:"allowBuiltin,omitempty"`
	Body         *MLServiceAnalysisParam `json:"body,omitempty" xml:"body,omitempty"`
	Version      *string                 `json:"version,omitempty" xml:"version,omitempty"`
}

func (s GetMLServiceResultsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetMLServiceResultsRequest) GoString() string {
	return s.String()
}

func (s *GetMLServiceResultsRequest) SetAllowBuiltin(v bool) *GetMLServiceResultsRequest {
	s.AllowBuiltin = &v
	return s
}

func (s *GetMLServiceResultsRequest) SetBody(v *MLServiceAnalysisParam) *GetMLServiceResultsRequest {
	s.Body = v
	return s
}

func (s *GetMLServiceResultsRequest) SetVersion(v string) *GetMLServiceResultsRequest {
	s.Version = &v
	return s
}

type GetMLServiceResultsResponseBody struct {
	Data   []map[string]*string `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	Status map[string]*string   `json:"status,omitempty" xml:"status,omitempty"`
}

func (s GetMLServiceResultsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetMLServiceResultsResponseBody) GoString() string {
	return s.String()
}

func (s *GetMLServiceResultsResponseBody) SetData(v []map[string]*string) *GetMLServiceResultsResponseBody {
	s.Data = v
	return s
}

func (s *GetMLServiceResultsResponseBody) SetStatus(v map[string]*string) *GetMLServiceResultsResponseBody {
	s.Status = v
	return s
}

type GetMLServiceResultsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMLServiceResultsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMLServiceResultsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetMLServiceResultsResponse) GoString() string {
	return s.String()
}

func (s *GetMLServiceResultsResponse) SetHeaders(v map[string]*string) *GetMLServiceResultsResponse {
	s.Headers = v
	return s
}

func (s *GetMLServiceResultsResponse) SetStatusCode(v int32) *GetMLServiceResultsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMLServiceResultsResponse) SetBody(v *GetMLServiceResultsResponseBody) *GetMLServiceResultsResponse {
	s.Body = v
	return s
}

type GetMachineGroupResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MachineGroup      `json:"body,omitempty" xml:"body,omitempty"`
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

type GetMetricStoreMeteringModeResponseBody struct {
	// example:
	//
	// ChargeByFunction
	MeteringMode *string `json:"meteringMode,omitempty" xml:"meteringMode,omitempty"`
}

func (s GetMetricStoreMeteringModeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetMetricStoreMeteringModeResponseBody) GoString() string {
	return s.String()
}

func (s *GetMetricStoreMeteringModeResponseBody) SetMeteringMode(v string) *GetMetricStoreMeteringModeResponseBody {
	s.MeteringMode = &v
	return s
}

type GetMetricStoreMeteringModeResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMetricStoreMeteringModeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMetricStoreMeteringModeResponse) String() string {
	return tea.Prettify(s)
}

func (s GetMetricStoreMeteringModeResponse) GoString() string {
	return s.String()
}

func (s *GetMetricStoreMeteringModeResponse) SetHeaders(v map[string]*string) *GetMetricStoreMeteringModeResponse {
	s.Headers = v
	return s
}

func (s *GetMetricStoreMeteringModeResponse) SetStatusCode(v int32) *GetMetricStoreMeteringModeResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMetricStoreMeteringModeResponse) SetBody(v *GetMetricStoreMeteringModeResponseBody) *GetMetricStoreMeteringModeResponse {
	s.Body = v
	return s
}

type GetOSSExportResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OSSExport         `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetOSSExportResponse) String() string {
	return tea.Prettify(s)
}

func (s GetOSSExportResponse) GoString() string {
	return s.String()
}

func (s *GetOSSExportResponse) SetHeaders(v map[string]*string) *GetOSSExportResponse {
	s.Headers = v
	return s
}

func (s *GetOSSExportResponse) SetStatusCode(v int32) *GetOSSExportResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOSSExportResponse) SetBody(v *OSSExport) *GetOSSExportResponse {
	s.Body = v
	return s
}

type GetOSSHDFSExportResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OSSExport         `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetOSSHDFSExportResponse) String() string {
	return tea.Prettify(s)
}

func (s GetOSSHDFSExportResponse) GoString() string {
	return s.String()
}

func (s *GetOSSHDFSExportResponse) SetHeaders(v map[string]*string) *GetOSSHDFSExportResponse {
	s.Headers = v
	return s
}

func (s *GetOSSHDFSExportResponse) SetStatusCode(v int32) *GetOSSHDFSExportResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOSSHDFSExportResponse) SetBody(v *OSSExport) *GetOSSHDFSExportResponse {
	s.Body = v
	return s
}

type GetOSSIngestionResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OSSIngestion      `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetOSSIngestionResponse) String() string {
	return tea.Prettify(s)
}

func (s GetOSSIngestionResponse) GoString() string {
	return s.String()
}

func (s *GetOSSIngestionResponse) SetHeaders(v map[string]*string) *GetOSSIngestionResponse {
	s.Headers = v
	return s
}

func (s *GetOSSIngestionResponse) SetStatusCode(v int32) *GetOSSIngestionResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOSSIngestionResponse) SetBody(v *OSSIngestion) *GetOSSIngestionResponse {
	s.Body = v
	return s
}

type GetProjectResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *Project           `json:"body,omitempty" xml:"body,omitempty"`
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
	// Specifies whether to enable the Dedicated SQL feature. For more information, see [Enable Dedicated SQL](https://help.aliyun.com/document_detail/223777.html). Valid values:
	//
	// 	- true
	//
	// 	- false (default): enables the Standard SQL feature.
	//
	// You can use the powerSql or **query*	- parameter to configure Dedicated SQL.
	//
	// example:
	//
	// false
	PowerSql *bool `json:"powerSql,omitempty" xml:"powerSql,omitempty"`
	// The standard SQL statement. In this example, the SQL statement queries the number of page views (PVs) from 2022-03-01 10:41:40 to 2022-03-01 10:56:40 in a Logstore whose name is nginx-moni.
	//
	// This parameter is required.
	//
	// example:
	//
	// SELECT COUNT(*) as pv FROM nginx-moni where __time__ &gt; 1646102500 and __time__ &lt; 1646103400
	Query *string `json:"query,omitempty" xml:"query,omitempty"`
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
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       []map[string]*string `json:"body,omitempty" xml:"body,omitempty" type:"Repeated"`
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
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *string            `json:"body,omitempty" xml:"body,omitempty"`
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
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SavedSearch       `json:"body,omitempty" xml:"body,omitempty"`
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

type GetScheduledSQLResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ScheduledSQL      `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetScheduledSQLResponse) String() string {
	return tea.Prettify(s)
}

func (s GetScheduledSQLResponse) GoString() string {
	return s.String()
}

func (s *GetScheduledSQLResponse) SetHeaders(v map[string]*string) *GetScheduledSQLResponse {
	s.Headers = v
	return s
}

func (s *GetScheduledSQLResponse) SetStatusCode(v int32) *GetScheduledSQLResponse {
	s.StatusCode = &v
	return s
}

func (s *GetScheduledSQLResponse) SetBody(v *ScheduledSQL) *GetScheduledSQLResponse {
	s.Body = v
	return s
}

type GetShipperStatusRequest struct {
	// The start time of the log shipping job. The value is a UNIX timestamp representing the number of seconds that have elapsed since the epoch time January 1, 1970, 00:00:00 UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1409529600
	From *int64 `json:"from,omitempty" xml:"from,omitempty"`
	// The line from which the query starts. Default value: 0.
	//
	// example:
	//
	// 0
	Offset *int32 `json:"offset,omitempty" xml:"offset,omitempty"`
	// The number of entries per page. Default value: 100. Maximum value: 500.
	//
	// example:
	//
	// 100
	Size *int32 `json:"size,omitempty" xml:"size,omitempty"`
	// The status of the log shipping job. This parameter is empty by default, which indicates that log shipping jobs in all states are returned. Valid values: success, fail, and running.
	//
	// example:
	//
	// success
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The end time of the log shipping job. The value is a UNIX timestamp representing the number of seconds that have elapsed since the epoch time January 1, 1970, 00:00:00 UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1627269085
	To *int64 `json:"to,omitempty" xml:"to,omitempty"`
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
	// The number of log shipping jobs returned on the current page.
	//
	// example:
	//
	// 10
	Count *int64 `json:"count,omitempty" xml:"count,omitempty"`
	// The statistics about log shipping jobs.
	Statistics *GetShipperStatusResponseBodyStatistics `json:"statistics,omitempty" xml:"statistics,omitempty" type:"Struct"`
	// The details of log shipping jobs.
	Tasks *GetShipperStatusResponseBodyTasks `json:"tasks,omitempty" xml:"tasks,omitempty" type:"Struct"`
	// The total number of log shipping jobs.
	//
	// example:
	//
	// 20
	Total *int64 `json:"total,omitempty" xml:"total,omitempty"`
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
	// The number of log shipping jobs that are in the fail state.
	//
	// example:
	//
	// 0
	Fail *int64 `json:"fail,omitempty" xml:"fail,omitempty"`
	// The number of log shipping jobs that are in the running state.
	//
	// example:
	//
	// 0
	Running *int64 `json:"running,omitempty" xml:"running,omitempty"`
	// The number of log shipping jobs that are in the success state.
	//
	// example:
	//
	// 20
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
	// The ID of the log shipping job.
	//
	// example:
	//
	// abcdefghijk
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// The error code of the log shipping job.
	//
	// example:
	//
	// UnAuthorized
	TaskCode *string `json:"taskCode,omitempty" xml:"taskCode,omitempty"`
	// The start time of the log shipping job. The value is a UNIX timestamp representing the number of seconds that have elapsed since the epoch time January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1448925013
	TaskCreateTime *int64 `json:"taskCreateTime,omitempty" xml:"taskCreateTime,omitempty"`
	// The number of logs that are shipped in the log shipping job.
	//
	// example:
	//
	// 0
	TaskDataLines *int32 `json:"taskDataLines,omitempty" xml:"taskDataLines,omitempty"`
	// The end time of the log shipping job. The value is a UNIX timestamp representing the number of seconds that have elapsed since the epoch time January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1448926013
	TaskFinishTime *int64 `json:"taskFinishTime,omitempty" xml:"taskFinishTime,omitempty"`
	// The time when Simple Log Service receives the most recent log of the log shipping job. The value is a UNIX timestamp representing the number of seconds that have elapsed since the epoch time January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1448915013
	TaskLastDataReceiveTime *int64 `json:"taskLastDataReceiveTime,omitempty" xml:"taskLastDataReceiveTime,omitempty"`
	// The error message of the log shipping job.
	//
	// example:
	//
	// Internal server error
	TaskMessage *string `json:"taskMessage,omitempty" xml:"taskMessage,omitempty"`
	// The status of the log shipping job. Valid values: running, success, and fail.
	//
	// example:
	//
	// success
	TaskStatus *string `json:"taskStatus,omitempty" xml:"taskStatus,omitempty"`
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
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetShipperStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type GetSlsServiceResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ServiceStatus     `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSlsServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSlsServiceResponse) GoString() string {
	return s.String()
}

func (s *GetSlsServiceResponse) SetHeaders(v map[string]*string) *GetSlsServiceResponse {
	s.Headers = v
	return s
}

func (s *GetSlsServiceResponse) SetStatusCode(v int32) *GetSlsServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSlsServiceResponse) SetBody(v *ServiceStatus) *GetSlsServiceResponse {
	s.Body = v
	return s
}

type GetSqlInstanceResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       []*GetSqlInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" type:"Repeated"`
}

func (s GetSqlInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSqlInstanceResponse) GoString() string {
	return s.String()
}

func (s *GetSqlInstanceResponse) SetHeaders(v map[string]*string) *GetSqlInstanceResponse {
	s.Headers = v
	return s
}

func (s *GetSqlInstanceResponse) SetStatusCode(v int32) *GetSqlInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSqlInstanceResponse) SetBody(v []*GetSqlInstanceResponseBody) *GetSqlInstanceResponse {
	s.Body = v
	return s
}

type GetSqlInstanceResponseBody struct {
	// example:
	//
	// project_name
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 2
	Cu *int32 `json:"cu,omitempty" xml:"cu,omitempty"`
	// example:
	//
	// 1710230272
	CreateTime *int32 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// 1710230272
	UpdateTime *int32 `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
	// example:
	//
	// true
	UseAsDefault *bool `json:"useAsDefault,omitempty" xml:"useAsDefault,omitempty"`
}

func (s GetSqlInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSqlInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *GetSqlInstanceResponseBody) SetName(v string) *GetSqlInstanceResponseBody {
	s.Name = &v
	return s
}

func (s *GetSqlInstanceResponseBody) SetCu(v int32) *GetSqlInstanceResponseBody {
	s.Cu = &v
	return s
}

func (s *GetSqlInstanceResponseBody) SetCreateTime(v int32) *GetSqlInstanceResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetSqlInstanceResponseBody) SetUpdateTime(v int32) *GetSqlInstanceResponseBody {
	s.UpdateTime = &v
	return s
}

func (s *GetSqlInstanceResponseBody) SetUseAsDefault(v bool) *GetSqlInstanceResponseBody {
	s.UseAsDefault = &v
	return s
}

type GetStoreViewResponseBody struct {
	// example:
	//
	// logstore
	StoreType *string           `json:"storeType,omitempty" xml:"storeType,omitempty"`
	Stores    []*StoreViewStore `json:"stores,omitempty" xml:"stores,omitempty" type:"Repeated"`
}

func (s GetStoreViewResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetStoreViewResponseBody) GoString() string {
	return s.String()
}

func (s *GetStoreViewResponseBody) SetStoreType(v string) *GetStoreViewResponseBody {
	s.StoreType = &v
	return s
}

func (s *GetStoreViewResponseBody) SetStores(v []*StoreViewStore) *GetStoreViewResponseBody {
	s.Stores = v
	return s
}

type GetStoreViewResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetStoreViewResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetStoreViewResponse) String() string {
	return tea.Prettify(s)
}

func (s GetStoreViewResponse) GoString() string {
	return s.String()
}

func (s *GetStoreViewResponse) SetHeaders(v map[string]*string) *GetStoreViewResponse {
	s.Headers = v
	return s
}

func (s *GetStoreViewResponse) SetStatusCode(v int32) *GetStoreViewResponse {
	s.StatusCode = &v
	return s
}

func (s *GetStoreViewResponse) SetBody(v *GetStoreViewResponseBody) *GetStoreViewResponse {
	s.Body = v
	return s
}

type GetStoreViewIndexResponseBody struct {
	Indexes []*GetStoreViewIndexResponseBodyIndexes `json:"indexes,omitempty" xml:"indexes,omitempty" type:"Repeated"`
}

func (s GetStoreViewIndexResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetStoreViewIndexResponseBody) GoString() string {
	return s.String()
}

func (s *GetStoreViewIndexResponseBody) SetIndexes(v []*GetStoreViewIndexResponseBodyIndexes) *GetStoreViewIndexResponseBody {
	s.Indexes = v
	return s
}

type GetStoreViewIndexResponseBodyIndexes struct {
	Index *Index `json:"index,omitempty" xml:"index,omitempty"`
	// example:
	//
	// my-logstore
	Logstore *string `json:"logstore,omitempty" xml:"logstore,omitempty"`
	// example:
	//
	// example-project
	Project *string `json:"project,omitempty" xml:"project,omitempty"`
}

func (s GetStoreViewIndexResponseBodyIndexes) String() string {
	return tea.Prettify(s)
}

func (s GetStoreViewIndexResponseBodyIndexes) GoString() string {
	return s.String()
}

func (s *GetStoreViewIndexResponseBodyIndexes) SetIndex(v *Index) *GetStoreViewIndexResponseBodyIndexes {
	s.Index = v
	return s
}

func (s *GetStoreViewIndexResponseBodyIndexes) SetLogstore(v string) *GetStoreViewIndexResponseBodyIndexes {
	s.Logstore = &v
	return s
}

func (s *GetStoreViewIndexResponseBodyIndexes) SetProject(v string) *GetStoreViewIndexResponseBodyIndexes {
	s.Project = &v
	return s
}

type GetStoreViewIndexResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetStoreViewIndexResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetStoreViewIndexResponse) String() string {
	return tea.Prettify(s)
}

func (s GetStoreViewIndexResponse) GoString() string {
	return s.String()
}

func (s *GetStoreViewIndexResponse) SetHeaders(v map[string]*string) *GetStoreViewIndexResponse {
	s.Headers = v
	return s
}

func (s *GetStoreViewIndexResponse) SetStatusCode(v int32) *GetStoreViewIndexResponse {
	s.StatusCode = &v
	return s
}

func (s *GetStoreViewIndexResponse) SetBody(v *GetStoreViewIndexResponseBody) *GetStoreViewIndexResponse {
	s.Body = v
	return s
}

type ListAlertsRequest struct {
	// example:
	//
	// ali-test-logstore
	Logstore *string `json:"logstore,omitempty" xml:"logstore,omitempty"`
	// example:
	//
	// 0
	Offset *int32 `json:"offset,omitempty" xml:"offset,omitempty"`
	// example:
	//
	// 10
	Size *int32 `json:"size,omitempty" xml:"size,omitempty"`
}

func (s ListAlertsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAlertsRequest) GoString() string {
	return s.String()
}

func (s *ListAlertsRequest) SetLogstore(v string) *ListAlertsRequest {
	s.Logstore = &v
	return s
}

func (s *ListAlertsRequest) SetOffset(v int32) *ListAlertsRequest {
	s.Offset = &v
	return s
}

func (s *ListAlertsRequest) SetSize(v int32) *ListAlertsRequest {
	s.Size = &v
	return s
}

type ListAlertsResponseBody struct {
	Count   *int32   `json:"count,omitempty" xml:"count,omitempty"`
	Results []*Alert `json:"results,omitempty" xml:"results,omitempty" type:"Repeated"`
	Total   *int32   `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListAlertsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAlertsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAlertsResponseBody) SetCount(v int32) *ListAlertsResponseBody {
	s.Count = &v
	return s
}

func (s *ListAlertsResponseBody) SetResults(v []*Alert) *ListAlertsResponseBody {
	s.Results = v
	return s
}

func (s *ListAlertsResponseBody) SetTotal(v int32) *ListAlertsResponseBody {
	s.Total = &v
	return s
}

type ListAlertsResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAlertsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAlertsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAlertsResponse) GoString() string {
	return s.String()
}

func (s *ListAlertsResponse) SetHeaders(v map[string]*string) *ListAlertsResponse {
	s.Headers = v
	return s
}

func (s *ListAlertsResponse) SetStatusCode(v int32) *ListAlertsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAlertsResponse) SetBody(v *ListAlertsResponseBody) *ListAlertsResponse {
	s.Body = v
	return s
}

type ListAnnotationDataRequest struct {
	// The line from which the query starts.
	//
	// example:
	//
	// 0
	Offset *int32 `json:"offset,omitempty" xml:"offset,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 100
	Size *int32 `json:"size,omitempty" xml:"size,omitempty"`
}

func (s ListAnnotationDataRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAnnotationDataRequest) GoString() string {
	return s.String()
}

func (s *ListAnnotationDataRequest) SetOffset(v int32) *ListAnnotationDataRequest {
	s.Offset = &v
	return s
}

func (s *ListAnnotationDataRequest) SetSize(v int32) *ListAnnotationDataRequest {
	s.Size = &v
	return s
}

type ListAnnotationDataResponseBody struct {
	// The data returned.
	Data []*MLDataParam `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 20
	Total *int32 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListAnnotationDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAnnotationDataResponseBody) GoString() string {
	return s.String()
}

func (s *ListAnnotationDataResponseBody) SetData(v []*MLDataParam) *ListAnnotationDataResponseBody {
	s.Data = v
	return s
}

func (s *ListAnnotationDataResponseBody) SetTotal(v int32) *ListAnnotationDataResponseBody {
	s.Total = &v
	return s
}

type ListAnnotationDataResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAnnotationDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAnnotationDataResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAnnotationDataResponse) GoString() string {
	return s.String()
}

func (s *ListAnnotationDataResponse) SetHeaders(v map[string]*string) *ListAnnotationDataResponse {
	s.Headers = v
	return s
}

func (s *ListAnnotationDataResponse) SetStatusCode(v int32) *ListAnnotationDataResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAnnotationDataResponse) SetBody(v *ListAnnotationDataResponseBody) *ListAnnotationDataResponse {
	s.Body = v
	return s
}

type ListAnnotationDataSetsRequest struct {
	// The line from which the query starts.
	//
	// example:
	//
	// 0
	Offset *int32 `json:"offset,omitempty" xml:"offset,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 100
	Size *int32 `json:"size,omitempty" xml:"size,omitempty"`
}

func (s ListAnnotationDataSetsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAnnotationDataSetsRequest) GoString() string {
	return s.String()
}

func (s *ListAnnotationDataSetsRequest) SetOffset(v int32) *ListAnnotationDataSetsRequest {
	s.Offset = &v
	return s
}

func (s *ListAnnotationDataSetsRequest) SetSize(v int32) *ListAnnotationDataSetsRequest {
	s.Size = &v
	return s
}

type ListAnnotationDataSetsResponseBody struct {
	// The data returned.
	Data []*MLDataSetParam `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 20
	Total *int32 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListAnnotationDataSetsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAnnotationDataSetsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAnnotationDataSetsResponseBody) SetData(v []*MLDataSetParam) *ListAnnotationDataSetsResponseBody {
	s.Data = v
	return s
}

func (s *ListAnnotationDataSetsResponseBody) SetTotal(v int32) *ListAnnotationDataSetsResponseBody {
	s.Total = &v
	return s
}

type ListAnnotationDataSetsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAnnotationDataSetsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAnnotationDataSetsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAnnotationDataSetsResponse) GoString() string {
	return s.String()
}

func (s *ListAnnotationDataSetsResponse) SetHeaders(v map[string]*string) *ListAnnotationDataSetsResponse {
	s.Headers = v
	return s
}

func (s *ListAnnotationDataSetsResponse) SetStatusCode(v int32) *ListAnnotationDataSetsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAnnotationDataSetsResponse) SetBody(v *ListAnnotationDataSetsResponseBody) *ListAnnotationDataSetsResponse {
	s.Body = v
	return s
}

type ListAnnotationLabelsRequest struct {
	// The line from which the query starts.
	//
	// example:
	//
	// 0
	Offset *int32 `json:"offset,omitempty" xml:"offset,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 100
	Size *int32 `json:"size,omitempty" xml:"size,omitempty"`
}

func (s ListAnnotationLabelsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAnnotationLabelsRequest) GoString() string {
	return s.String()
}

func (s *ListAnnotationLabelsRequest) SetOffset(v int32) *ListAnnotationLabelsRequest {
	s.Offset = &v
	return s
}

func (s *ListAnnotationLabelsRequest) SetSize(v int32) *ListAnnotationLabelsRequest {
	s.Size = &v
	return s
}

type ListAnnotationLabelsResponseBody struct {
	// The data returned.
	Data []*MLLabelParam `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// The total number of tags that meet the query conditions.
	//
	// example:
	//
	// 20
	Total *int32 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListAnnotationLabelsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAnnotationLabelsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAnnotationLabelsResponseBody) SetData(v []*MLLabelParam) *ListAnnotationLabelsResponseBody {
	s.Data = v
	return s
}

func (s *ListAnnotationLabelsResponseBody) SetTotal(v int32) *ListAnnotationLabelsResponseBody {
	s.Total = &v
	return s
}

type ListAnnotationLabelsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAnnotationLabelsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAnnotationLabelsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAnnotationLabelsResponse) GoString() string {
	return s.String()
}

func (s *ListAnnotationLabelsResponse) SetHeaders(v map[string]*string) *ListAnnotationLabelsResponse {
	s.Headers = v
	return s
}

func (s *ListAnnotationLabelsResponse) SetStatusCode(v int32) *ListAnnotationLabelsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAnnotationLabelsResponse) SetBody(v *ListAnnotationLabelsResponseBody) *ListAnnotationLabelsResponse {
	s.Body = v
	return s
}

type ListCollectionPoliciesRequest struct {
	Attribute *ListCollectionPoliciesRequestAttribute `json:"attribute,omitempty" xml:"attribute,omitempty" type:"Struct"`
	// example:
	//
	// access_log
	DataCode *string `json:"dataCode,omitempty" xml:"dataCode,omitempty"`
	// example:
	//
	// your-test-bucket1
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// example:
	//
	// 1
	PageNum *int32 `json:"pageNum,omitempty" xml:"pageNum,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// your_log_policy
	PolicyName *string `json:"policyName,omitempty" xml:"policyName,omitempty"`
	// example:
	//
	// oss
	ProductCode *string `json:"productCode,omitempty" xml:"productCode,omitempty"`
}

func (s ListCollectionPoliciesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListCollectionPoliciesRequest) GoString() string {
	return s.String()
}

func (s *ListCollectionPoliciesRequest) SetAttribute(v *ListCollectionPoliciesRequestAttribute) *ListCollectionPoliciesRequest {
	s.Attribute = v
	return s
}

func (s *ListCollectionPoliciesRequest) SetDataCode(v string) *ListCollectionPoliciesRequest {
	s.DataCode = &v
	return s
}

func (s *ListCollectionPoliciesRequest) SetInstanceId(v string) *ListCollectionPoliciesRequest {
	s.InstanceId = &v
	return s
}

func (s *ListCollectionPoliciesRequest) SetPageNum(v int32) *ListCollectionPoliciesRequest {
	s.PageNum = &v
	return s
}

func (s *ListCollectionPoliciesRequest) SetPageSize(v int32) *ListCollectionPoliciesRequest {
	s.PageSize = &v
	return s
}

func (s *ListCollectionPoliciesRequest) SetPolicyName(v string) *ListCollectionPoliciesRequest {
	s.PolicyName = &v
	return s
}

func (s *ListCollectionPoliciesRequest) SetProductCode(v string) *ListCollectionPoliciesRequest {
	s.ProductCode = &v
	return s
}

type ListCollectionPoliciesRequestAttribute struct {
	// example:
	//
	// your-app-name
	App *string `json:"app,omitempty" xml:"app,omitempty"`
	// example:
	//
	// your-policy-group
	PolicyGroup *string `json:"policyGroup,omitempty" xml:"policyGroup,omitempty"`
}

func (s ListCollectionPoliciesRequestAttribute) String() string {
	return tea.Prettify(s)
}

func (s ListCollectionPoliciesRequestAttribute) GoString() string {
	return s.String()
}

func (s *ListCollectionPoliciesRequestAttribute) SetApp(v string) *ListCollectionPoliciesRequestAttribute {
	s.App = &v
	return s
}

func (s *ListCollectionPoliciesRequestAttribute) SetPolicyGroup(v string) *ListCollectionPoliciesRequestAttribute {
	s.PolicyGroup = &v
	return s
}

type ListCollectionPoliciesShrinkRequest struct {
	AttributeShrink *string `json:"attribute,omitempty" xml:"attribute,omitempty"`
	// example:
	//
	// access_log
	DataCode *string `json:"dataCode,omitempty" xml:"dataCode,omitempty"`
	// example:
	//
	// your-test-bucket1
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// example:
	//
	// 1
	PageNum *int32 `json:"pageNum,omitempty" xml:"pageNum,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// your_log_policy
	PolicyName *string `json:"policyName,omitempty" xml:"policyName,omitempty"`
	// example:
	//
	// oss
	ProductCode *string `json:"productCode,omitempty" xml:"productCode,omitempty"`
}

func (s ListCollectionPoliciesShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListCollectionPoliciesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListCollectionPoliciesShrinkRequest) SetAttributeShrink(v string) *ListCollectionPoliciesShrinkRequest {
	s.AttributeShrink = &v
	return s
}

func (s *ListCollectionPoliciesShrinkRequest) SetDataCode(v string) *ListCollectionPoliciesShrinkRequest {
	s.DataCode = &v
	return s
}

func (s *ListCollectionPoliciesShrinkRequest) SetInstanceId(v string) *ListCollectionPoliciesShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *ListCollectionPoliciesShrinkRequest) SetPageNum(v int32) *ListCollectionPoliciesShrinkRequest {
	s.PageNum = &v
	return s
}

func (s *ListCollectionPoliciesShrinkRequest) SetPageSize(v int32) *ListCollectionPoliciesShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListCollectionPoliciesShrinkRequest) SetPolicyName(v string) *ListCollectionPoliciesShrinkRequest {
	s.PolicyName = &v
	return s
}

func (s *ListCollectionPoliciesShrinkRequest) SetProductCode(v string) *ListCollectionPoliciesShrinkRequest {
	s.ProductCode = &v
	return s
}

type ListCollectionPoliciesResponseBody struct {
	// example:
	//
	// 1
	CurrentCount *int32                                    `json:"currentCount,omitempty" xml:"currentCount,omitempty"`
	Data         []*ListCollectionPoliciesResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	TotalCount *int32 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListCollectionPoliciesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListCollectionPoliciesResponseBody) GoString() string {
	return s.String()
}

func (s *ListCollectionPoliciesResponseBody) SetCurrentCount(v int32) *ListCollectionPoliciesResponseBody {
	s.CurrentCount = &v
	return s
}

func (s *ListCollectionPoliciesResponseBody) SetData(v []*ListCollectionPoliciesResponseBodyData) *ListCollectionPoliciesResponseBody {
	s.Data = v
	return s
}

func (s *ListCollectionPoliciesResponseBody) SetTotalCount(v int32) *ListCollectionPoliciesResponseBody {
	s.TotalCount = &v
	return s
}

type ListCollectionPoliciesResponseBodyData struct {
	Attribute        *ListCollectionPoliciesResponseBodyDataAttribute        `json:"attribute,omitempty" xml:"attribute,omitempty" type:"Struct"`
	CentralizeConfig *ListCollectionPoliciesResponseBodyDataCentralizeConfig `json:"centralizeConfig,omitempty" xml:"centralizeConfig,omitempty" type:"Struct"`
	// example:
	//
	// false
	CentralizeEnabled *bool `json:"centralizeEnabled,omitempty" xml:"centralizeEnabled,omitempty"`
	// example:
	//
	// access_log
	DataCode *string `json:"dataCode,omitempty" xml:"dataCode,omitempty"`
	// example:
	//
	// true
	Enabled      *bool                                               `json:"enabled,omitempty" xml:"enabled,omitempty"`
	PolicyConfig *ListCollectionPoliciesResponseBodyDataPolicyConfig `json:"policyConfig,omitempty" xml:"policyConfig,omitempty" type:"Struct"`
	// example:
	//
	// your_log_policy
	PolicyName *string `json:"policyName,omitempty" xml:"policyName,omitempty"`
	// example:
	//
	// oss
	ProductCode *string `json:"productCode,omitempty" xml:"productCode,omitempty"`
}

func (s ListCollectionPoliciesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListCollectionPoliciesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListCollectionPoliciesResponseBodyData) SetAttribute(v *ListCollectionPoliciesResponseBodyDataAttribute) *ListCollectionPoliciesResponseBodyData {
	s.Attribute = v
	return s
}

func (s *ListCollectionPoliciesResponseBodyData) SetCentralizeConfig(v *ListCollectionPoliciesResponseBodyDataCentralizeConfig) *ListCollectionPoliciesResponseBodyData {
	s.CentralizeConfig = v
	return s
}

func (s *ListCollectionPoliciesResponseBodyData) SetCentralizeEnabled(v bool) *ListCollectionPoliciesResponseBodyData {
	s.CentralizeEnabled = &v
	return s
}

func (s *ListCollectionPoliciesResponseBodyData) SetDataCode(v string) *ListCollectionPoliciesResponseBodyData {
	s.DataCode = &v
	return s
}

func (s *ListCollectionPoliciesResponseBodyData) SetEnabled(v bool) *ListCollectionPoliciesResponseBodyData {
	s.Enabled = &v
	return s
}

func (s *ListCollectionPoliciesResponseBodyData) SetPolicyConfig(v *ListCollectionPoliciesResponseBodyDataPolicyConfig) *ListCollectionPoliciesResponseBodyData {
	s.PolicyConfig = v
	return s
}

func (s *ListCollectionPoliciesResponseBodyData) SetPolicyName(v string) *ListCollectionPoliciesResponseBodyData {
	s.PolicyName = &v
	return s
}

func (s *ListCollectionPoliciesResponseBodyData) SetProductCode(v string) *ListCollectionPoliciesResponseBodyData {
	s.ProductCode = &v
	return s
}

type ListCollectionPoliciesResponseBodyDataAttribute struct {
	// example:
	//
	// your-app-name
	App *string `json:"app,omitempty" xml:"app,omitempty"`
	// example:
	//
	// your-policy-group
	PolicyGroup *string `json:"policyGroup,omitempty" xml:"policyGroup,omitempty"`
}

func (s ListCollectionPoliciesResponseBodyDataAttribute) String() string {
	return tea.Prettify(s)
}

func (s ListCollectionPoliciesResponseBodyDataAttribute) GoString() string {
	return s.String()
}

func (s *ListCollectionPoliciesResponseBodyDataAttribute) SetApp(v string) *ListCollectionPoliciesResponseBodyDataAttribute {
	s.App = &v
	return s
}

func (s *ListCollectionPoliciesResponseBodyDataAttribute) SetPolicyGroup(v string) *ListCollectionPoliciesResponseBodyDataAttribute {
	s.PolicyGroup = &v
	return s
}

type ListCollectionPoliciesResponseBodyDataCentralizeConfig struct {
	// example:
	//
	// your-sls-logstore-in-beijing
	DestLogstore *string `json:"destLogstore,omitempty" xml:"destLogstore,omitempty"`
	// example:
	//
	// your-sls-project-in-beijing
	DestProject *string `json:"destProject,omitempty" xml:"destProject,omitempty"`
	// example:
	//
	// cn-beijing
	DestRegion *string `json:"destRegion,omitempty" xml:"destRegion,omitempty"`
	// example:
	//
	// your-sls-logstore-ttl
	DestTTL *int32 `json:"destTTL,omitempty" xml:"destTTL,omitempty"`
}

func (s ListCollectionPoliciesResponseBodyDataCentralizeConfig) String() string {
	return tea.Prettify(s)
}

func (s ListCollectionPoliciesResponseBodyDataCentralizeConfig) GoString() string {
	return s.String()
}

func (s *ListCollectionPoliciesResponseBodyDataCentralizeConfig) SetDestLogstore(v string) *ListCollectionPoliciesResponseBodyDataCentralizeConfig {
	s.DestLogstore = &v
	return s
}

func (s *ListCollectionPoliciesResponseBodyDataCentralizeConfig) SetDestProject(v string) *ListCollectionPoliciesResponseBodyDataCentralizeConfig {
	s.DestProject = &v
	return s
}

func (s *ListCollectionPoliciesResponseBodyDataCentralizeConfig) SetDestRegion(v string) *ListCollectionPoliciesResponseBodyDataCentralizeConfig {
	s.DestRegion = &v
	return s
}

func (s *ListCollectionPoliciesResponseBodyDataCentralizeConfig) SetDestTTL(v int32) *ListCollectionPoliciesResponseBodyDataCentralizeConfig {
	s.DestTTL = &v
	return s
}

type ListCollectionPoliciesResponseBodyDataPolicyConfig struct {
	InstanceIds []*string `json:"instanceIds,omitempty" xml:"instanceIds,omitempty" type:"Repeated"`
	Regions     []*string `json:"regions,omitempty" xml:"regions,omitempty" type:"Repeated"`
	// example:
	//
	// all
	ResourceMode *string                `json:"resourceMode,omitempty" xml:"resourceMode,omitempty"`
	ResourceTags map[string]interface{} `json:"resourceTags,omitempty" xml:"resourceTags,omitempty"`
}

func (s ListCollectionPoliciesResponseBodyDataPolicyConfig) String() string {
	return tea.Prettify(s)
}

func (s ListCollectionPoliciesResponseBodyDataPolicyConfig) GoString() string {
	return s.String()
}

func (s *ListCollectionPoliciesResponseBodyDataPolicyConfig) SetInstanceIds(v []*string) *ListCollectionPoliciesResponseBodyDataPolicyConfig {
	s.InstanceIds = v
	return s
}

func (s *ListCollectionPoliciesResponseBodyDataPolicyConfig) SetRegions(v []*string) *ListCollectionPoliciesResponseBodyDataPolicyConfig {
	s.Regions = v
	return s
}

func (s *ListCollectionPoliciesResponseBodyDataPolicyConfig) SetResourceMode(v string) *ListCollectionPoliciesResponseBodyDataPolicyConfig {
	s.ResourceMode = &v
	return s
}

func (s *ListCollectionPoliciesResponseBodyDataPolicyConfig) SetResourceTags(v map[string]interface{}) *ListCollectionPoliciesResponseBodyDataPolicyConfig {
	s.ResourceTags = v
	return s
}

type ListCollectionPoliciesResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCollectionPoliciesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCollectionPoliciesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListCollectionPoliciesResponse) GoString() string {
	return s.String()
}

func (s *ListCollectionPoliciesResponse) SetHeaders(v map[string]*string) *ListCollectionPoliciesResponse {
	s.Headers = v
	return s
}

func (s *ListCollectionPoliciesResponse) SetStatusCode(v int32) *ListCollectionPoliciesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCollectionPoliciesResponse) SetBody(v *ListCollectionPoliciesResponseBody) *ListCollectionPoliciesResponse {
	s.Body = v
	return s
}

type ListConfigRequest struct {
	// The name of the Logtail configuration.
	//
	// example:
	//
	// logtail-config-sample
	ConfigName *string `json:"configName,omitempty" xml:"configName,omitempty"`
	// The name of the Logstore.
	//
	// This parameter is required.
	//
	// example:
	//
	// ali-test-logstore
	LogstoreName *string `json:"logstoreName,omitempty" xml:"logstoreName,omitempty"`
	// The line from which the query starts. Default value: 0.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	Offset *int64 `json:"offset,omitempty" xml:"offset,omitempty"`
	// The number of entries per page. Maximum value: 500.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	Size *int64 `json:"size,omitempty" xml:"size,omitempty"`
}

func (s ListConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s ListConfigRequest) GoString() string {
	return s.String()
}

func (s *ListConfigRequest) SetConfigName(v string) *ListConfigRequest {
	s.ConfigName = &v
	return s
}

func (s *ListConfigRequest) SetLogstoreName(v string) *ListConfigRequest {
	s.LogstoreName = &v
	return s
}

func (s *ListConfigRequest) SetOffset(v int64) *ListConfigRequest {
	s.Offset = &v
	return s
}

func (s *ListConfigRequest) SetSize(v int64) *ListConfigRequest {
	s.Size = &v
	return s
}

type ListConfigResponseBody struct {
	// The Logtail configurations that are returned on the current page.
	Configs []*string `json:"configs,omitempty" xml:"configs,omitempty" type:"Repeated"`
	// The number of Logtail configurations that are returned on the current page.
	//
	// example:
	//
	// 3
	Count *int32 `json:"count,omitempty" xml:"count,omitempty"`
	// The total number of Logtail configurations that meet the query conditions.
	//
	// example:
	//
	// 2
	Total *int32 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ListConfigResponseBody) SetConfigs(v []*string) *ListConfigResponseBody {
	s.Configs = v
	return s
}

func (s *ListConfigResponseBody) SetCount(v int32) *ListConfigResponseBody {
	s.Count = &v
	return s
}

func (s *ListConfigResponseBody) SetTotal(v int32) *ListConfigResponseBody {
	s.Total = &v
	return s
}

type ListConfigResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s ListConfigResponse) GoString() string {
	return s.String()
}

func (s *ListConfigResponse) SetHeaders(v map[string]*string) *ListConfigResponse {
	s.Headers = v
	return s
}

func (s *ListConfigResponse) SetStatusCode(v int32) *ListConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *ListConfigResponse) SetBody(v *ListConfigResponseBody) *ListConfigResponse {
	s.Body = v
	return s
}

type ListConsumerGroupResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       []*ConsumerGroup   `json:"body,omitempty" xml:"body,omitempty" type:"Repeated"`
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

type ListDashboardRequest struct {
	// The line from which the query starts. Default value: 0.
	//
	// example:
	//
	// 0
	Offset *int32 `json:"offset,omitempty" xml:"offset,omitempty"`
	// The number of entries per page. Maximum value: 500. Default value: 500.
	//
	// example:
	//
	// 10
	Size *int32 `json:"size,omitempty" xml:"size,omitempty"`
}

func (s ListDashboardRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDashboardRequest) GoString() string {
	return s.String()
}

func (s *ListDashboardRequest) SetOffset(v int32) *ListDashboardRequest {
	s.Offset = &v
	return s
}

func (s *ListDashboardRequest) SetSize(v int32) *ListDashboardRequest {
	s.Size = &v
	return s
}

type ListDashboardResponseBody struct {
	// The details of the dashboard.
	DashboardItems []*ListDashboardResponseBodyDashboardItems `json:"dashboardItems,omitempty" xml:"dashboardItems,omitempty" type:"Repeated"`
	// The queried dashboards. Each dashboard in the array is specified by dashboardName.
	Dashboards []*string `json:"dashboards,omitempty" xml:"dashboards,omitempty" type:"Repeated"`
}

func (s ListDashboardResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDashboardResponseBody) GoString() string {
	return s.String()
}

func (s *ListDashboardResponseBody) SetDashboardItems(v []*ListDashboardResponseBodyDashboardItems) *ListDashboardResponseBody {
	s.DashboardItems = v
	return s
}

func (s *ListDashboardResponseBody) SetDashboards(v []*string) *ListDashboardResponseBody {
	s.Dashboards = v
	return s
}

type ListDashboardResponseBodyDashboardItems struct {
	// The dashboard ID. The ID must be unique in a project. Fuzzy search is supported. For example, if you enter da, all dashboards whose IDs start with da are queried.
	//
	// example:
	//
	// dashboard-1609294922657-434834
	DashboardName *string `json:"dashboardName,omitempty" xml:"dashboardName,omitempty"`
	// The display name of the dashboard.
	//
	// example:
	//
	// data-ingest
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
}

func (s ListDashboardResponseBodyDashboardItems) String() string {
	return tea.Prettify(s)
}

func (s ListDashboardResponseBodyDashboardItems) GoString() string {
	return s.String()
}

func (s *ListDashboardResponseBodyDashboardItems) SetDashboardName(v string) *ListDashboardResponseBodyDashboardItems {
	s.DashboardName = &v
	return s
}

func (s *ListDashboardResponseBodyDashboardItems) SetDisplayName(v string) *ListDashboardResponseBodyDashboardItems {
	s.DisplayName = &v
	return s
}

type ListDashboardResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDashboardResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDashboardResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDashboardResponse) GoString() string {
	return s.String()
}

func (s *ListDashboardResponse) SetHeaders(v map[string]*string) *ListDashboardResponse {
	s.Headers = v
	return s
}

func (s *ListDashboardResponse) SetStatusCode(v int32) *ListDashboardResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDashboardResponse) SetBody(v *ListDashboardResponseBody) *ListDashboardResponse {
	s.Body = v
	return s
}

type ListDomainsRequest struct {
	// The domain name that is used to match custom domain names. For example, if you set domainName to `example.com`, the matched domain names are `a.example.com` and `b.example.com`.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"domainName,omitempty" xml:"domainName,omitempty"`
	// The line from which the query starts. Default value: 0.
	//
	// example:
	//
	// 0
	Offset *int32 `json:"offset,omitempty" xml:"offset,omitempty"`
	// The number of entries per page. Default value: 500. Maximum value: 500.
	//
	// example:
	//
	// 10
	Size *int32 `json:"size,omitempty" xml:"size,omitempty"`
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
	// The number of domain names that are returned on the current page.
	//
	// example:
	//
	// 1
	Count *int64 `json:"count,omitempty" xml:"count,omitempty"`
	// The domain names.
	Domains []*string `json:"domains,omitempty" xml:"domains,omitempty" type:"Repeated"`
	// The total number of domain names that are returned.
	//
	// example:
	//
	// 1
	Total *int64 `json:"total,omitempty" xml:"total,omitempty"`
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
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDomainsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type ListETLsRequest struct {
	Logstore *string `json:"logstore,omitempty" xml:"logstore,omitempty"`
	// example:
	//
	// 0
	Offset *int32 `json:"offset,omitempty" xml:"offset,omitempty"`
	// example:
	//
	// 10
	Size *int32 `json:"size,omitempty" xml:"size,omitempty"`
}

func (s ListETLsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListETLsRequest) GoString() string {
	return s.String()
}

func (s *ListETLsRequest) SetLogstore(v string) *ListETLsRequest {
	s.Logstore = &v
	return s
}

func (s *ListETLsRequest) SetOffset(v int32) *ListETLsRequest {
	s.Offset = &v
	return s
}

func (s *ListETLsRequest) SetSize(v int32) *ListETLsRequest {
	s.Size = &v
	return s
}

type ListETLsResponseBody struct {
	Count   *int32 `json:"count,omitempty" xml:"count,omitempty"`
	Results []*ETL `json:"results,omitempty" xml:"results,omitempty" type:"Repeated"`
	Total   *int32 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListETLsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListETLsResponseBody) GoString() string {
	return s.String()
}

func (s *ListETLsResponseBody) SetCount(v int32) *ListETLsResponseBody {
	s.Count = &v
	return s
}

func (s *ListETLsResponseBody) SetResults(v []*ETL) *ListETLsResponseBody {
	s.Results = v
	return s
}

func (s *ListETLsResponseBody) SetTotal(v int32) *ListETLsResponseBody {
	s.Total = &v
	return s
}

type ListETLsResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListETLsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListETLsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListETLsResponse) GoString() string {
	return s.String()
}

func (s *ListETLsResponse) SetHeaders(v map[string]*string) *ListETLsResponse {
	s.Headers = v
	return s
}

func (s *ListETLsResponse) SetStatusCode(v int32) *ListETLsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListETLsResponse) SetBody(v *ListETLsResponseBody) *ListETLsResponse {
	s.Body = v
	return s
}

type ListExternalStoreRequest struct {
	// The name of the external store. You can query external stores that contain a specified string.
	//
	// example:
	//
	// store
	ExternalStoreName *string `json:"externalStoreName,omitempty" xml:"externalStoreName,omitempty"`
	// The line from which the query starts. Default value: 0.
	//
	// example:
	//
	// 0
	Offset *int32 `json:"offset,omitempty" xml:"offset,omitempty"`
	// The number of entries per page. Maximum value: 500.
	//
	// example:
	//
	// 10
	Sizs *int32 `json:"sizs,omitempty" xml:"sizs,omitempty"`
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
	// The number of external stores returned on the current page.
	//
	// example:
	//
	// 3
	Count *int32 `json:"count,omitempty" xml:"count,omitempty"`
	// The names of the external stores.
	Externalstores []*string `json:"externalstores,omitempty" xml:"externalstores,omitempty" type:"Repeated"`
	// The number of external stores that meet the query conditions.
	//
	// example:
	//
	// 3
	Total *int32 `json:"total,omitempty" xml:"total,omitempty"`
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

func (s *ListExternalStoreResponseBody) SetExternalstores(v []*string) *ListExternalStoreResponseBody {
	s.Externalstores = v
	return s
}

func (s *ListExternalStoreResponseBody) SetTotal(v int32) *ListExternalStoreResponseBody {
	s.Total = &v
	return s
}

type ListExternalStoreResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListExternalStoreResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type ListLogStoresRequest struct {
	// The name of the Logstore. Fuzzy match is supported. For example, if you enter test, Logstores whose name contains test are returned.
	//
	// example:
	//
	// my-logstore
	LogstoreName *string `json:"logstoreName,omitempty" xml:"logstoreName,omitempty"`
	// The type of the Logstore. Valid values: standard and query.
	//
	// 	- **standard**: Standard Logstore. This type of Logstore supports the log analysis feature and is suitable for scenarios such as real-time monitoring and interactive analysis. You can also use this type of Logstore to build a comprehensive observability system.
	//
	// 	- **query**: Query Logstore. This type of Logstore supports high-performance queries. The index traffic fee of a Query Logstore is approximately half that of a Standard Logstore. Query Logstores do not support SQL analysis. Query Logstores are suitable for scenarios in which the volume of data is large, the log retention period is long, or log analysis is not required. Log retention periods of weeks or months are considered long.
	//
	// example:
	//
	// standard
	Mode *string `json:"mode,omitempty" xml:"mode,omitempty"`
	// The line from which the query starts. Default value: 0.
	//
	// example:
	//
	// 0
	Offset *int32 `json:"offset,omitempty" xml:"offset,omitempty"`
	// The number of entries per page. Maximum value: 500. Default value: 500.
	//
	// example:
	//
	// 10
	Size *int32 `json:"size,omitempty" xml:"size,omitempty"`
	// The type of the data that you want to query. Valid values:
	//
	// 	- None: logs
	//
	// 	- Metrics: metrics
	//
	// example:
	//
	// None
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
	// The number of entries returned on the current page.
	//
	// example:
	//
	// 2
	Count *int32 `json:"count,omitempty" xml:"count,omitempty"`
	// The Logstores that meet the query conditions.
	//
	// example:
	//
	// ["test-1","test-2"]
	Logstores []*string `json:"logstores,omitempty" xml:"logstores,omitempty" type:"Repeated"`
	// The number of the Logstores that meet the query conditions.
	//
	// example:
	//
	// 2
	Total *int32 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListLogStoresResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListLogStoresResponseBody) GoString() string {
	return s.String()
}

func (s *ListLogStoresResponseBody) SetCount(v int32) *ListLogStoresResponseBody {
	s.Count = &v
	return s
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
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListLogStoresResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type ListLogtailPipelineConfigRequest struct {
	// The name of the Logtail pipeline configuration.
	//
	// example:
	//
	// logtail-config-sample
	ConfigName *string `json:"configName,omitempty" xml:"configName,omitempty"`
	// The name of the Logstore.
	//
	// example:
	//
	// test-logstore
	LogstoreName *string `json:"logstoreName,omitempty" xml:"logstoreName,omitempty"`
	// The line from which the query starts.
	//
	// example:
	//
	// 0
	Offset *int64 `json:"offset,omitempty" xml:"offset,omitempty"`
	// The number of Logtail pipeline configurations per page.
	//
	// example:
	//
	// 20
	Size *int64 `json:"size,omitempty" xml:"size,omitempty"`
}

func (s ListLogtailPipelineConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s ListLogtailPipelineConfigRequest) GoString() string {
	return s.String()
}

func (s *ListLogtailPipelineConfigRequest) SetConfigName(v string) *ListLogtailPipelineConfigRequest {
	s.ConfigName = &v
	return s
}

func (s *ListLogtailPipelineConfigRequest) SetLogstoreName(v string) *ListLogtailPipelineConfigRequest {
	s.LogstoreName = &v
	return s
}

func (s *ListLogtailPipelineConfigRequest) SetOffset(v int64) *ListLogtailPipelineConfigRequest {
	s.Offset = &v
	return s
}

func (s *ListLogtailPipelineConfigRequest) SetSize(v int64) *ListLogtailPipelineConfigRequest {
	s.Size = &v
	return s
}

type ListLogtailPipelineConfigResponseBody struct {
	// The Logtail pipeline configurations that are returned on the current page.
	Configs []*string `json:"configs,omitempty" xml:"configs,omitempty" type:"Repeated"`
	// The number of Logtail pipeline configurations that are returned on the current page.
	//
	// example:
	//
	// 10
	Count *int32 `json:"count,omitempty" xml:"count,omitempty"`
	// The total number of Logtail pipeline configurations in the current project.
	//
	// example:
	//
	// 20
	Total *int32 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListLogtailPipelineConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListLogtailPipelineConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ListLogtailPipelineConfigResponseBody) SetConfigs(v []*string) *ListLogtailPipelineConfigResponseBody {
	s.Configs = v
	return s
}

func (s *ListLogtailPipelineConfigResponseBody) SetCount(v int32) *ListLogtailPipelineConfigResponseBody {
	s.Count = &v
	return s
}

func (s *ListLogtailPipelineConfigResponseBody) SetTotal(v int32) *ListLogtailPipelineConfigResponseBody {
	s.Total = &v
	return s
}

type ListLogtailPipelineConfigResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListLogtailPipelineConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListLogtailPipelineConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s ListLogtailPipelineConfigResponse) GoString() string {
	return s.String()
}

func (s *ListLogtailPipelineConfigResponse) SetHeaders(v map[string]*string) *ListLogtailPipelineConfigResponse {
	s.Headers = v
	return s
}

func (s *ListLogtailPipelineConfigResponse) SetStatusCode(v int32) *ListLogtailPipelineConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *ListLogtailPipelineConfigResponse) SetBody(v *ListLogtailPipelineConfigResponseBody) *ListLogtailPipelineConfigResponse {
	s.Body = v
	return s
}

type ListMachineGroupRequest struct {
	// The name of the machine group. This parameter is used to filter machine groups. Partial match is supported.
	//
	// example:
	//
	// test-machine-group
	GroupName *string `json:"groupName,omitempty" xml:"groupName,omitempty"`
	// The line from which the query starts. Default value: 0.
	//
	// example:
	//
	// 1
	Offset *int32 `json:"offset,omitempty" xml:"offset,omitempty"`
	// The number of entries per page. Maximum value: 500.
	//
	// example:
	//
	// 10
	Size *int32 `json:"size,omitempty" xml:"size,omitempty"`
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
	// The number of machine groups that are returned on the current page.
	//
	// example:
	//
	// 2
	Count *int32 `json:"count,omitempty" xml:"count,omitempty"`
	// The machine groups that meet the query conditions.
	//
	// example:
	//
	// [ "test-machine-group-1", "test-machine-group-2" ]
	Machinegroups []*string `json:"machinegroups,omitempty" xml:"machinegroups,omitempty" type:"Repeated"`
	// The total number of machine groups that meet the query conditions.
	//
	// example:
	//
	// 2
	Total *int32 `json:"total,omitempty" xml:"total,omitempty"`
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
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListMachineGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The line from which the query starts. Default value: 0.
	//
	// example:
	//
	// 0
	Offset *int32 `json:"offset,omitempty" xml:"offset,omitempty"`
	// The number of entries per page. Default value: 100. Maximum value: 500.
	//
	// example:
	//
	// 10
	Size *int32 `json:"size,omitempty" xml:"size,omitempty"`
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
	// The number of machines that are returned on the current page.
	//
	// example:
	//
	// 3
	Count *int32 `json:"count,omitempty" xml:"count,omitempty"`
	// The machines that are returned.
	Machines []*Machine `json:"machines,omitempty" xml:"machines,omitempty" type:"Repeated"`
	// The total number of machines.
	//
	// example:
	//
	// 8
	Total *int32 `json:"total,omitempty" xml:"total,omitempty"`
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
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListMachinesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type ListOSSExportsRequest struct {
	Logstore *string `json:"logstore,omitempty" xml:"logstore,omitempty"`
	// example:
	//
	// 0
	Offset *int32 `json:"offset,omitempty" xml:"offset,omitempty"`
	// example:
	//
	// 100
	Size *int32 `json:"size,omitempty" xml:"size,omitempty"`
}

func (s ListOSSExportsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListOSSExportsRequest) GoString() string {
	return s.String()
}

func (s *ListOSSExportsRequest) SetLogstore(v string) *ListOSSExportsRequest {
	s.Logstore = &v
	return s
}

func (s *ListOSSExportsRequest) SetOffset(v int32) *ListOSSExportsRequest {
	s.Offset = &v
	return s
}

func (s *ListOSSExportsRequest) SetSize(v int32) *ListOSSExportsRequest {
	s.Size = &v
	return s
}

type ListOSSExportsResponseBody struct {
	// example:
	//
	// 2
	Count   *int32       `json:"count,omitempty" xml:"count,omitempty"`
	Results []*OSSExport `json:"results,omitempty" xml:"results,omitempty" type:"Repeated"`
	// example:
	//
	// 10
	Total *int32 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListOSSExportsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListOSSExportsResponseBody) GoString() string {
	return s.String()
}

func (s *ListOSSExportsResponseBody) SetCount(v int32) *ListOSSExportsResponseBody {
	s.Count = &v
	return s
}

func (s *ListOSSExportsResponseBody) SetResults(v []*OSSExport) *ListOSSExportsResponseBody {
	s.Results = v
	return s
}

func (s *ListOSSExportsResponseBody) SetTotal(v int32) *ListOSSExportsResponseBody {
	s.Total = &v
	return s
}

type ListOSSExportsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListOSSExportsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListOSSExportsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListOSSExportsResponse) GoString() string {
	return s.String()
}

func (s *ListOSSExportsResponse) SetHeaders(v map[string]*string) *ListOSSExportsResponse {
	s.Headers = v
	return s
}

func (s *ListOSSExportsResponse) SetStatusCode(v int32) *ListOSSExportsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListOSSExportsResponse) SetBody(v *ListOSSExportsResponseBody) *ListOSSExportsResponse {
	s.Body = v
	return s
}

type ListOSSHDFSExportsRequest struct {
	Logstore *string `json:"logstore,omitempty" xml:"logstore,omitempty"`
	// example:
	//
	// 0
	Offset *int32 `json:"offset,omitempty" xml:"offset,omitempty"`
	// example:
	//
	// 100
	Size *int32 `json:"size,omitempty" xml:"size,omitempty"`
}

func (s ListOSSHDFSExportsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListOSSHDFSExportsRequest) GoString() string {
	return s.String()
}

func (s *ListOSSHDFSExportsRequest) SetLogstore(v string) *ListOSSHDFSExportsRequest {
	s.Logstore = &v
	return s
}

func (s *ListOSSHDFSExportsRequest) SetOffset(v int32) *ListOSSHDFSExportsRequest {
	s.Offset = &v
	return s
}

func (s *ListOSSHDFSExportsRequest) SetSize(v int32) *ListOSSHDFSExportsRequest {
	s.Size = &v
	return s
}

type ListOSSHDFSExportsResponseBody struct {
	// example:
	//
	// 2
	Count   *int32       `json:"count,omitempty" xml:"count,omitempty"`
	Results []*OSSExport `json:"results,omitempty" xml:"results,omitempty" type:"Repeated"`
	// example:
	//
	// 10
	Total *int32 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListOSSHDFSExportsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListOSSHDFSExportsResponseBody) GoString() string {
	return s.String()
}

func (s *ListOSSHDFSExportsResponseBody) SetCount(v int32) *ListOSSHDFSExportsResponseBody {
	s.Count = &v
	return s
}

func (s *ListOSSHDFSExportsResponseBody) SetResults(v []*OSSExport) *ListOSSHDFSExportsResponseBody {
	s.Results = v
	return s
}

func (s *ListOSSHDFSExportsResponseBody) SetTotal(v int32) *ListOSSHDFSExportsResponseBody {
	s.Total = &v
	return s
}

type ListOSSHDFSExportsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListOSSHDFSExportsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListOSSHDFSExportsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListOSSHDFSExportsResponse) GoString() string {
	return s.String()
}

func (s *ListOSSHDFSExportsResponse) SetHeaders(v map[string]*string) *ListOSSHDFSExportsResponse {
	s.Headers = v
	return s
}

func (s *ListOSSHDFSExportsResponse) SetStatusCode(v int32) *ListOSSHDFSExportsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListOSSHDFSExportsResponse) SetBody(v *ListOSSHDFSExportsResponseBody) *ListOSSHDFSExportsResponse {
	s.Body = v
	return s
}

type ListOSSIngestionsRequest struct {
	Logstore *string `json:"logstore,omitempty" xml:"logstore,omitempty"`
	// example:
	//
	// 0
	Offset *int32 `json:"offset,omitempty" xml:"offset,omitempty"`
	// example:
	//
	// 10
	Size *int32 `json:"size,omitempty" xml:"size,omitempty"`
}

func (s ListOSSIngestionsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListOSSIngestionsRequest) GoString() string {
	return s.String()
}

func (s *ListOSSIngestionsRequest) SetLogstore(v string) *ListOSSIngestionsRequest {
	s.Logstore = &v
	return s
}

func (s *ListOSSIngestionsRequest) SetOffset(v int32) *ListOSSIngestionsRequest {
	s.Offset = &v
	return s
}

func (s *ListOSSIngestionsRequest) SetSize(v int32) *ListOSSIngestionsRequest {
	s.Size = &v
	return s
}

type ListOSSIngestionsResponseBody struct {
	Count   *int32          `json:"count,omitempty" xml:"count,omitempty"`
	Results []*OSSIngestion `json:"results,omitempty" xml:"results,omitempty" type:"Repeated"`
	Total   *int32          `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListOSSIngestionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListOSSIngestionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListOSSIngestionsResponseBody) SetCount(v int32) *ListOSSIngestionsResponseBody {
	s.Count = &v
	return s
}

func (s *ListOSSIngestionsResponseBody) SetResults(v []*OSSIngestion) *ListOSSIngestionsResponseBody {
	s.Results = v
	return s
}

func (s *ListOSSIngestionsResponseBody) SetTotal(v int32) *ListOSSIngestionsResponseBody {
	s.Total = &v
	return s
}

type ListOSSIngestionsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListOSSIngestionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListOSSIngestionsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListOSSIngestionsResponse) GoString() string {
	return s.String()
}

func (s *ListOSSIngestionsResponse) SetHeaders(v map[string]*string) *ListOSSIngestionsResponse {
	s.Headers = v
	return s
}

func (s *ListOSSIngestionsResponse) SetStatusCode(v int32) *ListOSSIngestionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListOSSIngestionsResponse) SetBody(v *ListOSSIngestionsResponseBody) *ListOSSIngestionsResponse {
	s.Body = v
	return s
}

type ListProjectRequest struct {
	// The line from which the query starts. Default value: 0.
	//
	// example:
	//
	// 0
	Offset *int32 `json:"offset,omitempty" xml:"offset,omitempty"`
	// The name of the project.
	//
	// example:
	//
	// ali-test-project
	ProjectName     *string `json:"projectName,omitempty" xml:"projectName,omitempty"`
	ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	// The number of entries per page. Default value: 100. This operation can return up to 500 projects.
	//
	// example:
	//
	// 10
	Size *int32 `json:"size,omitempty" xml:"size,omitempty"`
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

func (s *ListProjectRequest) SetResourceGroupId(v string) *ListProjectRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListProjectRequest) SetSize(v int32) *ListProjectRequest {
	s.Size = &v
	return s
}

type ListProjectResponseBody struct {
	// The number of returned projects on the current page.
	//
	// example:
	//
	// 2
	Count *int64 `json:"count,omitempty" xml:"count,omitempty"`
	// The projects that meet the query conditions.
	Projects []*Project `json:"projects,omitempty" xml:"projects,omitempty" type:"Repeated"`
	// The total number of projects that meet the query conditions.
	//
	// example:
	//
	// 11
	Total *int64 `json:"total,omitempty" xml:"total,omitempty"`
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
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListProjectResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The line from which the query starts. Default value: 0.
	//
	// example:
	//
	// 0
	Offset *int32 `json:"offset,omitempty" xml:"offset,omitempty"`
	// The number of entries per page. Maximum value: 500.
	//
	// example:
	//
	// 10
	Size *int32 `json:"size,omitempty" xml:"size,omitempty"`
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
	// The number of saved searches returned on the current page.
	//
	// example:
	//
	// 4
	Count *int32 `json:"count,omitempty" xml:"count,omitempty"`
	// The saved searches.
	//
	// example:
	//
	// [ "test-1", "test-2" ]
	SavedsearchItems []*SavedSearch `json:"savedsearchItems,omitempty" xml:"savedsearchItems,omitempty" type:"Repeated"`
	// The total number of saved searches that meet the query conditions.
	//
	// example:
	//
	// 4
	Total *int32 `json:"total,omitempty" xml:"total,omitempty"`
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
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSavedSearchResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type ListScheduledSQLsRequest struct {
	Logstore *string `json:"logstore,omitempty" xml:"logstore,omitempty"`
	// example:
	//
	// 0
	Offset *int64 `json:"offset,omitempty" xml:"offset,omitempty"`
	// example:
	//
	// 100
	Size *int64 `json:"size,omitempty" xml:"size,omitempty"`
}

func (s ListScheduledSQLsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListScheduledSQLsRequest) GoString() string {
	return s.String()
}

func (s *ListScheduledSQLsRequest) SetLogstore(v string) *ListScheduledSQLsRequest {
	s.Logstore = &v
	return s
}

func (s *ListScheduledSQLsRequest) SetOffset(v int64) *ListScheduledSQLsRequest {
	s.Offset = &v
	return s
}

func (s *ListScheduledSQLsRequest) SetSize(v int64) *ListScheduledSQLsRequest {
	s.Size = &v
	return s
}

type ListScheduledSQLsResponseBody struct {
	// example:
	//
	// 10
	Count   *int32          `json:"count,omitempty" xml:"count,omitempty"`
	Results []*ScheduledSQL `json:"results,omitempty" xml:"results,omitempty" type:"Repeated"`
	// example:
	//
	// 80
	Total *int32 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListScheduledSQLsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListScheduledSQLsResponseBody) GoString() string {
	return s.String()
}

func (s *ListScheduledSQLsResponseBody) SetCount(v int32) *ListScheduledSQLsResponseBody {
	s.Count = &v
	return s
}

func (s *ListScheduledSQLsResponseBody) SetResults(v []*ScheduledSQL) *ListScheduledSQLsResponseBody {
	s.Results = v
	return s
}

func (s *ListScheduledSQLsResponseBody) SetTotal(v int32) *ListScheduledSQLsResponseBody {
	s.Total = &v
	return s
}

type ListScheduledSQLsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListScheduledSQLsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListScheduledSQLsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListScheduledSQLsResponse) GoString() string {
	return s.String()
}

func (s *ListScheduledSQLsResponse) SetHeaders(v map[string]*string) *ListScheduledSQLsResponse {
	s.Headers = v
	return s
}

func (s *ListScheduledSQLsResponse) SetStatusCode(v int32) *ListScheduledSQLsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListScheduledSQLsResponse) SetBody(v *ListScheduledSQLsResponseBody) *ListScheduledSQLsResponse {
	s.Body = v
	return s
}

type ListShardsResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       []*Shard           `json:"body,omitempty" xml:"body,omitempty" type:"Repeated"`
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
	// The number of log shipping jobs returned.
	//
	// example:
	//
	// 3
	Count *int64 `json:"count,omitempty" xml:"count,omitempty"`
	// The names of the log shipping jobs.
	Shipper []*string `json:"shipper,omitempty" xml:"shipper,omitempty" type:"Repeated"`
	// The total number of log shipping jobs.
	//
	// example:
	//
	// 5
	Total *int64 `json:"total,omitempty" xml:"total,omitempty"`
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
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListShipperResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type ListStoreViewsRequest struct {
	// example:
	//
	// my_storeview
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 0
	Offset *int32 `json:"offset,omitempty" xml:"offset,omitempty"`
	// example:
	//
	// 100
	Size *int32 `json:"size,omitempty" xml:"size,omitempty"`
	// example:
	//
	// logstore
	StoreType *string `json:"storeType,omitempty" xml:"storeType,omitempty"`
}

func (s ListStoreViewsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListStoreViewsRequest) GoString() string {
	return s.String()
}

func (s *ListStoreViewsRequest) SetName(v string) *ListStoreViewsRequest {
	s.Name = &v
	return s
}

func (s *ListStoreViewsRequest) SetOffset(v int32) *ListStoreViewsRequest {
	s.Offset = &v
	return s
}

func (s *ListStoreViewsRequest) SetSize(v int32) *ListStoreViewsRequest {
	s.Size = &v
	return s
}

func (s *ListStoreViewsRequest) SetStoreType(v string) *ListStoreViewsRequest {
	s.StoreType = &v
	return s
}

type ListStoreViewsResponseBody struct {
	// example:
	//
	// 100
	Count      *int32    `json:"count,omitempty" xml:"count,omitempty"`
	Storeviews []*string `json:"storeviews,omitempty" xml:"storeviews,omitempty" type:"Repeated"`
	// example:
	//
	// 100
	Total *int32 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListStoreViewsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListStoreViewsResponseBody) GoString() string {
	return s.String()
}

func (s *ListStoreViewsResponseBody) SetCount(v int32) *ListStoreViewsResponseBody {
	s.Count = &v
	return s
}

func (s *ListStoreViewsResponseBody) SetStoreviews(v []*string) *ListStoreViewsResponseBody {
	s.Storeviews = v
	return s
}

func (s *ListStoreViewsResponseBody) SetTotal(v int32) *ListStoreViewsResponseBody {
	s.Total = &v
	return s
}

type ListStoreViewsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListStoreViewsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListStoreViewsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListStoreViewsResponse) GoString() string {
	return s.String()
}

func (s *ListStoreViewsResponse) SetHeaders(v map[string]*string) *ListStoreViewsResponse {
	s.Headers = v
	return s
}

func (s *ListStoreViewsResponse) SetStatusCode(v int32) *ListStoreViewsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListStoreViewsResponse) SetBody(v *ListStoreViewsResponseBody) *ListStoreViewsResponse {
	s.Body = v
	return s
}

type ListTagResourcesRequest struct {
	// The IDs of the resources for which you want to query tags. You must specify at least one of resourceId and tags.
	ResourceId []*string `json:"resourceId,omitempty" xml:"resourceId,omitempty" type:"Repeated"`
	// The type of the resource. Set the value to project.
	//
	// This parameter is required.
	//
	// example:
	//
	// project
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
	// The tags that you want to use to filter resources based on exact match. Each tag is a key-value pair. You must specify at least one of resourceId and tags.
	//
	// You can enter up to 20 tags.
	Tags []*ListTagResourcesRequestTags `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
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
	// The key of the tag that you want to use to filter resources. For example, if you set the key to `"test-key"`, only resources to which the key is added are returned.``
	//
	// This parameter is required.
	//
	// example:
	//
	// key1
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// The value of the tag that you want to use to filter resources. If you set the value to null, resources are filtered based only on the key of the tag.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// value1
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
	// The IDs of the resources for which you want to query tags. You must specify at least one of resourceId and tags.
	ResourceIdShrink *string `json:"resourceId,omitempty" xml:"resourceId,omitempty"`
	// The type of the resource. Set the value to project.
	//
	// This parameter is required.
	//
	// example:
	//
	// project
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
	// The tags that you want to use to filter resources based on exact match. Each tag is a key-value pair. You must specify at least one of resourceId and tags.
	//
	// You can enter up to 20 tags.
	TagsShrink *string `json:"tags,omitempty" xml:"tags,omitempty"`
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
	// The pagination token that is used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a4883
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The returned tags.
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
	// The ID of the resource.
	//
	// example:
	//
	// ali-test-project
	ResourceId *string `json:"resourceId,omitempty" xml:"resourceId,omitempty"`
	// The type of the resource.
	//
	// example:
	//
	// project
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
	// The key of the tag.
	//
	// example:
	//
	// key1
	TagKey *string `json:"tagKey,omitempty" xml:"tagKey,omitempty"`
	// The value of the tag.
	//
	// example:
	//
	// value1
	TagValue *string `json:"tagValue,omitempty" xml:"tagValue,omitempty"`
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
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type MergeShardResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       []*Shard           `json:"body,omitempty" xml:"body,omitempty" type:"Repeated"`
}

func (s MergeShardResponse) String() string {
	return tea.Prettify(s)
}

func (s MergeShardResponse) GoString() string {
	return s.String()
}

func (s *MergeShardResponse) SetHeaders(v map[string]*string) *MergeShardResponse {
	s.Headers = v
	return s
}

func (s *MergeShardResponse) SetStatusCode(v int32) *MergeShardResponse {
	s.StatusCode = &v
	return s
}

func (s *MergeShardResponse) SetBody(v []*Shard) *MergeShardResponse {
	s.Body = v
	return s
}

type OpenSlsServiceResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s OpenSlsServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s OpenSlsServiceResponse) GoString() string {
	return s.String()
}

func (s *OpenSlsServiceResponse) SetHeaders(v map[string]*string) *OpenSlsServiceResponse {
	s.Headers = v
	return s
}

func (s *OpenSlsServiceResponse) SetStatusCode(v int32) *OpenSlsServiceResponse {
	s.StatusCode = &v
	return s
}

type PutAnnotationDataRequest struct {
	// The unique identifier of the data.
	//
	// example:
	//
	// 2156d560fc7c01420542df92cd6365ds
	AnnotationdataId *string `json:"annotationdataId,omitempty" xml:"annotationdataId,omitempty"`
	// The data structure of the request.
	MlDataParam *MLDataParam `json:"mlDataParam,omitempty" xml:"mlDataParam,omitempty"`
	// The raw log data.
	RawLog []map[string]*string `json:"rawLog,omitempty" xml:"rawLog,omitempty" type:"Repeated"`
}

func (s PutAnnotationDataRequest) String() string {
	return tea.Prettify(s)
}

func (s PutAnnotationDataRequest) GoString() string {
	return s.String()
}

func (s *PutAnnotationDataRequest) SetAnnotationdataId(v string) *PutAnnotationDataRequest {
	s.AnnotationdataId = &v
	return s
}

func (s *PutAnnotationDataRequest) SetMlDataParam(v *MLDataParam) *PutAnnotationDataRequest {
	s.MlDataParam = v
	return s
}

func (s *PutAnnotationDataRequest) SetRawLog(v []map[string]*string) *PutAnnotationDataRequest {
	s.RawLog = v
	return s
}

type PutAnnotationDataResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s PutAnnotationDataResponse) String() string {
	return tea.Prettify(s)
}

func (s PutAnnotationDataResponse) GoString() string {
	return s.String()
}

func (s *PutAnnotationDataResponse) SetHeaders(v map[string]*string) *PutAnnotationDataResponse {
	s.Headers = v
	return s
}

func (s *PutAnnotationDataResponse) SetStatusCode(v int32) *PutAnnotationDataResponse {
	s.StatusCode = &v
	return s
}

type PutProjectPolicyRequest struct {
	// The project policy.
	//
	// example:
	//
	// { 	"Version": "1", 	"Statement": [{ 		"Action": ["log:PostLogStoreLogs"], 		"Resource": "acs:log:*:*:project/exampleproject/*", 		"Effect": "Deny", 		"Condition": { 			"StringNotLike": { 				"acs:SourceVpc": ["vpc-*"] 			} 		} 	}] }
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
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
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

type PutProjectTransferAccelerationRequest struct {
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
}

func (s PutProjectTransferAccelerationRequest) String() string {
	return tea.Prettify(s)
}

func (s PutProjectTransferAccelerationRequest) GoString() string {
	return s.String()
}

func (s *PutProjectTransferAccelerationRequest) SetEnabled(v bool) *PutProjectTransferAccelerationRequest {
	s.Enabled = &v
	return s
}

type PutProjectTransferAccelerationResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s PutProjectTransferAccelerationResponse) String() string {
	return tea.Prettify(s)
}

func (s PutProjectTransferAccelerationResponse) GoString() string {
	return s.String()
}

func (s *PutProjectTransferAccelerationResponse) SetHeaders(v map[string]*string) *PutProjectTransferAccelerationResponse {
	s.Headers = v
	return s
}

func (s *PutProjectTransferAccelerationResponse) SetStatusCode(v int32) *PutProjectTransferAccelerationResponse {
	s.StatusCode = &v
	return s
}

type PutWebtrackingRequest struct {
	// The logs. Each element is a JSON object that indicates a log.
	//
	// >  **Note**: The time in a log that is collected by using the web tracking feature is the time at which Simple Log Service receives the log. You do not need to configure the __time__ field for each log. If this field exists, it is overwritten by the time at which Simple Log Service receives the log.
	//
	// This parameter is required.
	Logs []map[string]*string `json:"__logs__,omitempty" xml:"__logs__,omitempty" type:"Repeated"`
	// The source of the logs.
	//
	// This parameter is required.
	//
	// example:
	//
	// source
	Source *string `json:"__source__,omitempty" xml:"__source__,omitempty"`
	// The tags of the logs.
	Tags map[string]*string `json:"__tags__,omitempty" xml:"__tags__,omitempty"`
	// The topic of the logs.
	//
	// example:
	//
	// topic
	Topic *string `json:"__topic__,omitempty" xml:"__topic__,omitempty"`
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
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
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

type QueryMLServiceResultsRequest struct {
	AllowBuiltin *bool                   `json:"allowBuiltin,omitempty" xml:"allowBuiltin,omitempty"`
	Body         *MLServiceAnalysisParam `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryMLServiceResultsRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryMLServiceResultsRequest) GoString() string {
	return s.String()
}

func (s *QueryMLServiceResultsRequest) SetAllowBuiltin(v bool) *QueryMLServiceResultsRequest {
	s.AllowBuiltin = &v
	return s
}

func (s *QueryMLServiceResultsRequest) SetBody(v *MLServiceAnalysisParam) *QueryMLServiceResultsRequest {
	s.Body = v
	return s
}

type QueryMLServiceResultsResponseBody struct {
	Data   []map[string]*string `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	Status map[string]*string   `json:"status,omitempty" xml:"status,omitempty"`
}

func (s QueryMLServiceResultsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryMLServiceResultsResponseBody) GoString() string {
	return s.String()
}

func (s *QueryMLServiceResultsResponseBody) SetData(v []map[string]*string) *QueryMLServiceResultsResponseBody {
	s.Data = v
	return s
}

func (s *QueryMLServiceResultsResponseBody) SetStatus(v map[string]*string) *QueryMLServiceResultsResponseBody {
	s.Status = v
	return s
}

type QueryMLServiceResultsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryMLServiceResultsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryMLServiceResultsResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryMLServiceResultsResponse) GoString() string {
	return s.String()
}

func (s *QueryMLServiceResultsResponse) SetHeaders(v map[string]*string) *QueryMLServiceResultsResponse {
	s.Headers = v
	return s
}

func (s *QueryMLServiceResultsResponse) SetStatusCode(v int32) *QueryMLServiceResultsResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryMLServiceResultsResponse) SetBody(v *QueryMLServiceResultsResponseBody) *QueryMLServiceResultsResponse {
	s.Body = v
	return s
}

type RefreshTokenRequest struct {
	// example:
	//
	// 600
	AccessTokenExpirationTime *int64 `json:"accessTokenExpirationTime,omitempty" xml:"accessTokenExpirationTime,omitempty"`
	// example:
	//
	// eyJ***************.eyJ******************.KUT****************
	Ticket *string `json:"ticket,omitempty" xml:"ticket,omitempty"`
}

func (s RefreshTokenRequest) String() string {
	return tea.Prettify(s)
}

func (s RefreshTokenRequest) GoString() string {
	return s.String()
}

func (s *RefreshTokenRequest) SetAccessTokenExpirationTime(v int64) *RefreshTokenRequest {
	s.AccessTokenExpirationTime = &v
	return s
}

func (s *RefreshTokenRequest) SetTicket(v string) *RefreshTokenRequest {
	s.Ticket = &v
	return s
}

type RefreshTokenResponseBody struct {
	// example:
	//
	// eyJ***************.eyJ******************.KUT****************
	AccessToken *string `json:"accessToken,omitempty" xml:"accessToken,omitempty"`
}

func (s RefreshTokenResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RefreshTokenResponseBody) GoString() string {
	return s.String()
}

func (s *RefreshTokenResponseBody) SetAccessToken(v string) *RefreshTokenResponseBody {
	s.AccessToken = &v
	return s
}

type RefreshTokenResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RefreshTokenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RefreshTokenResponse) String() string {
	return tea.Prettify(s)
}

func (s RefreshTokenResponse) GoString() string {
	return s.String()
}

func (s *RefreshTokenResponse) SetHeaders(v map[string]*string) *RefreshTokenResponse {
	s.Headers = v
	return s
}

func (s *RefreshTokenResponse) SetStatusCode(v int32) *RefreshTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *RefreshTokenResponse) SetBody(v *RefreshTokenResponseBody) *RefreshTokenResponse {
	s.Body = v
	return s
}

type RemoveConfigFromMachineGroupResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
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
	// The position where the shard is split.
	//
	// example:
	//
	// ef000000000000000000000000000000
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// The number of new shards that are generated after splitting.
	//
	// example:
	//
	// 2
	ShardCount *int32 `json:"shardCount,omitempty" xml:"shardCount,omitempty"`
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
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       []*Shard           `json:"body,omitempty" xml:"body,omitempty" type:"Repeated"`
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

type StartETLResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s StartETLResponse) String() string {
	return tea.Prettify(s)
}

func (s StartETLResponse) GoString() string {
	return s.String()
}

func (s *StartETLResponse) SetHeaders(v map[string]*string) *StartETLResponse {
	s.Headers = v
	return s
}

func (s *StartETLResponse) SetStatusCode(v int32) *StartETLResponse {
	s.StatusCode = &v
	return s
}

type StartOSSExportResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s StartOSSExportResponse) String() string {
	return tea.Prettify(s)
}

func (s StartOSSExportResponse) GoString() string {
	return s.String()
}

func (s *StartOSSExportResponse) SetHeaders(v map[string]*string) *StartOSSExportResponse {
	s.Headers = v
	return s
}

func (s *StartOSSExportResponse) SetStatusCode(v int32) *StartOSSExportResponse {
	s.StatusCode = &v
	return s
}

type StartOSSHDFSExportResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s StartOSSHDFSExportResponse) String() string {
	return tea.Prettify(s)
}

func (s StartOSSHDFSExportResponse) GoString() string {
	return s.String()
}

func (s *StartOSSHDFSExportResponse) SetHeaders(v map[string]*string) *StartOSSHDFSExportResponse {
	s.Headers = v
	return s
}

func (s *StartOSSHDFSExportResponse) SetStatusCode(v int32) *StartOSSHDFSExportResponse {
	s.StatusCode = &v
	return s
}

type StartOSSIngestionResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s StartOSSIngestionResponse) String() string {
	return tea.Prettify(s)
}

func (s StartOSSIngestionResponse) GoString() string {
	return s.String()
}

func (s *StartOSSIngestionResponse) SetHeaders(v map[string]*string) *StartOSSIngestionResponse {
	s.Headers = v
	return s
}

func (s *StartOSSIngestionResponse) SetStatusCode(v int32) *StartOSSIngestionResponse {
	s.StatusCode = &v
	return s
}

type StopETLResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s StopETLResponse) String() string {
	return tea.Prettify(s)
}

func (s StopETLResponse) GoString() string {
	return s.String()
}

func (s *StopETLResponse) SetHeaders(v map[string]*string) *StopETLResponse {
	s.Headers = v
	return s
}

func (s *StopETLResponse) SetStatusCode(v int32) *StopETLResponse {
	s.StatusCode = &v
	return s
}

type StopOSSExportResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s StopOSSExportResponse) String() string {
	return tea.Prettify(s)
}

func (s StopOSSExportResponse) GoString() string {
	return s.String()
}

func (s *StopOSSExportResponse) SetHeaders(v map[string]*string) *StopOSSExportResponse {
	s.Headers = v
	return s
}

func (s *StopOSSExportResponse) SetStatusCode(v int32) *StopOSSExportResponse {
	s.StatusCode = &v
	return s
}

type StopOSSHDFSExportResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s StopOSSHDFSExportResponse) String() string {
	return tea.Prettify(s)
}

func (s StopOSSHDFSExportResponse) GoString() string {
	return s.String()
}

func (s *StopOSSHDFSExportResponse) SetHeaders(v map[string]*string) *StopOSSHDFSExportResponse {
	s.Headers = v
	return s
}

func (s *StopOSSHDFSExportResponse) SetStatusCode(v int32) *StopOSSHDFSExportResponse {
	s.StatusCode = &v
	return s
}

type StopOSSIngestionResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s StopOSSIngestionResponse) String() string {
	return tea.Prettify(s)
}

func (s StopOSSIngestionResponse) GoString() string {
	return s.String()
}

func (s *StopOSSIngestionResponse) SetHeaders(v map[string]*string) *StopOSSIngestionResponse {
	s.Headers = v
	return s
}

func (s *StopOSSIngestionResponse) SetStatusCode(v int32) *StopOSSIngestionResponse {
	s.StatusCode = &v
	return s
}

type TagResourcesRequest struct {
	// The resource IDs. You can specify only one resource and add tags to the resource.
	//
	// This parameter is required.
	ResourceId []*string `json:"resourceId,omitempty" xml:"resourceId,omitempty" type:"Repeated"`
	// The type of the resource. Set the value to project.
	//
	// This parameter is required.
	//
	// example:
	//
	// project
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
	// The tags that you want to add to the resource. Up to 20 tags are supported at a time. Each tag is a key-value pair.
	//
	// This parameter is required.
	Tags []*TagResourcesRequestTags `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
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
	// The key of the tag. The key must meet the following requirements:
	//
	// 	- The key must be `1 to 128` characters in length.
	//
	// 	- The key cannot contain `"http://"` or `"https://"`.
	//
	// 	- The key cannot start with `"acs:"` or `"aliyun"`.
	//
	// This parameter is required.
	//
	// example:
	//
	// key1
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// The value of the tag. The value must meet the following requirements:
	//
	// 	- The value must be `1 to 128` characters in length.
	//
	// 	- The value cannot contain `"http://"` or `"https://"`.
	//
	// This parameter is required.
	//
	// example:
	//
	// value1
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
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
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
	// example:
	//
	// false
	All *bool `json:"all,omitempty" xml:"all,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ali-test-project
	ResourceId []*string `json:"resourceId,omitempty" xml:"resourceId,omitempty" type:"Repeated"`
	// example:
	//
	// project
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

func (s *UntagResourcesRequest) SetResourceId(v []*string) *UntagResourcesRequest {
	s.ResourceId = v
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
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
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
	// This parameter is required.
	Configuration *AlertConfiguration `json:"configuration,omitempty" xml:"configuration,omitempty"`
	Description   *string             `json:"description,omitempty" xml:"description,omitempty"`
	// This parameter is required.
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	// This parameter is required.
	Schedule *Schedule `json:"schedule,omitempty" xml:"schedule,omitempty"`
}

func (s UpdateAlertRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateAlertRequest) GoString() string {
	return s.String()
}

func (s *UpdateAlertRequest) SetConfiguration(v *AlertConfiguration) *UpdateAlertRequest {
	s.Configuration = v
	return s
}

func (s *UpdateAlertRequest) SetDescription(v string) *UpdateAlertRequest {
	s.Description = &v
	return s
}

func (s *UpdateAlertRequest) SetDisplayName(v string) *UpdateAlertRequest {
	s.DisplayName = &v
	return s
}

func (s *UpdateAlertRequest) SetSchedule(v *Schedule) *UpdateAlertRequest {
	s.Schedule = v
	return s
}

type UpdateAlertResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
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

type UpdateAnnotationDataSetRequest struct {
	// The data structure of the request.
	Body *MLDataSetParam `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAnnotationDataSetRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateAnnotationDataSetRequest) GoString() string {
	return s.String()
}

func (s *UpdateAnnotationDataSetRequest) SetBody(v *MLDataSetParam) *UpdateAnnotationDataSetRequest {
	s.Body = v
	return s
}

type UpdateAnnotationDataSetResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s UpdateAnnotationDataSetResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateAnnotationDataSetResponse) GoString() string {
	return s.String()
}

func (s *UpdateAnnotationDataSetResponse) SetHeaders(v map[string]*string) *UpdateAnnotationDataSetResponse {
	s.Headers = v
	return s
}

func (s *UpdateAnnotationDataSetResponse) SetStatusCode(v int32) *UpdateAnnotationDataSetResponse {
	s.StatusCode = &v
	return s
}

type UpdateAnnotationLabelRequest struct {
	// The data structure of the request.
	Body *MLLabelParam `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAnnotationLabelRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateAnnotationLabelRequest) GoString() string {
	return s.String()
}

func (s *UpdateAnnotationLabelRequest) SetBody(v *MLLabelParam) *UpdateAnnotationLabelRequest {
	s.Body = v
	return s
}

type UpdateAnnotationLabelResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s UpdateAnnotationLabelResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateAnnotationLabelResponse) GoString() string {
	return s.String()
}

func (s *UpdateAnnotationLabelResponse) SetHeaders(v map[string]*string) *UpdateAnnotationLabelResponse {
	s.Headers = v
	return s
}

func (s *UpdateAnnotationLabelResponse) SetStatusCode(v int32) *UpdateAnnotationLabelResponse {
	s.StatusCode = &v
	return s
}

type UpdateConfigRequest struct {
	// The body of the request.
	Body *LogtailConfig `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateConfigRequest) GoString() string {
	return s.String()
}

func (s *UpdateConfigRequest) SetBody(v *LogtailConfig) *UpdateConfigRequest {
	s.Body = v
	return s
}

type UpdateConfigResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s UpdateConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateConfigResponse) GoString() string {
	return s.String()
}

func (s *UpdateConfigResponse) SetHeaders(v map[string]*string) *UpdateConfigResponse {
	s.Headers = v
	return s
}

func (s *UpdateConfigResponse) SetStatusCode(v int32) *UpdateConfigResponse {
	s.StatusCode = &v
	return s
}

type UpdateConsumerGroupRequest struct {
	// Specifies whether to consume data in sequence. Valid values:
	//
	// 	- true: If a shard is split, the data in the original shard is consumed first. Then, the data in the new shards is consumed at the same time. If shards are merged, the data in the original shards is consumed first. Then, the data in the new shard is consumed.
	//
	// 	- false: The data in all shards is consumed at the same time. If a new shard is generated after a shard is split or shards are merged, the data in the new shard is immediately consumed.
	//
	// example:
	//
	// true
	Order *bool `json:"order,omitempty" xml:"order,omitempty"`
	// The timeout period. If Simple Log Service does not receive heartbeats from a consumer within the timeout period, Simple Log Service deletes the consumer. Unit: seconds.
	//
	// example:
	//
	// 300
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
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
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

type UpdateDashboardRequest struct {
	// The attributes of the dashboard.
	Attribute map[string]*string `json:"attribute,omitempty" xml:"attribute,omitempty"`
	// The charts on the dashboard.
	//
	// This parameter is required.
	Charts []*Chart `json:"charts,omitempty" xml:"charts,omitempty" type:"Repeated"`
	// The name of the dashboard.
	//
	// This parameter is required.
	//
	// example:
	//
	// dashboard-1609294922657-434834
	DashboardName *string `json:"dashboardName,omitempty" xml:"dashboardName,omitempty"`
	// The description of the dashboard.
	//
	// example:
	//
	// test dashboard.
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The display name of the dashboard.
	//
	// This parameter is required.
	//
	// example:
	//
	// Method pv
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
}

func (s UpdateDashboardRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateDashboardRequest) GoString() string {
	return s.String()
}

func (s *UpdateDashboardRequest) SetAttribute(v map[string]*string) *UpdateDashboardRequest {
	s.Attribute = v
	return s
}

func (s *UpdateDashboardRequest) SetCharts(v []*Chart) *UpdateDashboardRequest {
	s.Charts = v
	return s
}

func (s *UpdateDashboardRequest) SetDashboardName(v string) *UpdateDashboardRequest {
	s.DashboardName = &v
	return s
}

func (s *UpdateDashboardRequest) SetDescription(v string) *UpdateDashboardRequest {
	s.Description = &v
	return s
}

func (s *UpdateDashboardRequest) SetDisplayName(v string) *UpdateDashboardRequest {
	s.DisplayName = &v
	return s
}

type UpdateDashboardResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s UpdateDashboardResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateDashboardResponse) GoString() string {
	return s.String()
}

func (s *UpdateDashboardResponse) SetHeaders(v map[string]*string) *UpdateDashboardResponse {
	s.Headers = v
	return s
}

func (s *UpdateDashboardResponse) SetStatusCode(v int32) *UpdateDashboardResponse {
	s.StatusCode = &v
	return s
}

type UpdateETLRequest struct {
	// This parameter is required.
	Configuration *ETLConfiguration `json:"configuration,omitempty" xml:"configuration,omitempty"`
	// example:
	//
	// this is description
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// this is update
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
}

func (s UpdateETLRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateETLRequest) GoString() string {
	return s.String()
}

func (s *UpdateETLRequest) SetConfiguration(v *ETLConfiguration) *UpdateETLRequest {
	s.Configuration = v
	return s
}

func (s *UpdateETLRequest) SetDescription(v string) *UpdateETLRequest {
	s.Description = &v
	return s
}

func (s *UpdateETLRequest) SetDisplayName(v string) *UpdateETLRequest {
	s.DisplayName = &v
	return s
}

type UpdateETLResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s UpdateETLResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateETLResponse) GoString() string {
	return s.String()
}

func (s *UpdateETLResponse) SetHeaders(v map[string]*string) *UpdateETLResponse {
	s.Headers = v
	return s
}

func (s *UpdateETLResponse) SetStatusCode(v int32) *UpdateETLResponse {
	s.StatusCode = &v
	return s
}

type UpdateIndexRequest struct {
	// The configuration of field indexes. A field index is a key-value pair in which the key specifies the name of the field and the value specifies the index configuration of the field.
	Keys map[string]*KeysValue `json:"keys,omitempty" xml:"keys,omitempty"`
	// The configuration of full-text indexes.
	Line *UpdateIndexRequestLine `json:"line,omitempty" xml:"line,omitempty" type:"Struct"`
	// Specifies whether to turn on LogReduce. If you turn on LogReduce, only one of `log_reduce_white_list` and `log_reduce_black_list` takes effect.
	//
	// example:
	//
	// false
	LogReduce *bool `json:"log_reduce,omitempty" xml:"log_reduce,omitempty"`
	// The fields in the blacklist that you want to use to cluster logs.
	LogReduceBlackList []*string `json:"log_reduce_black_list,omitempty" xml:"log_reduce_black_list,omitempty" type:"Repeated"`
	// The fields in the whitelist that you want to use to cluster logs.
	LogReduceWhiteList []*string `json:"log_reduce_white_list,omitempty" xml:"log_reduce_white_list,omitempty" type:"Repeated"`
	// The maximum length of a field value that can be retained.
	//
	// example:
	//
	// 2048
	MaxTextLen *int32 `json:"max_text_len,omitempty" xml:"max_text_len,omitempty"`
	// The retention period of data. Unit: days. Valid values: 7, 30, and 90.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30
	Ttl *int32 `json:"ttl,omitempty" xml:"ttl,omitempty"`
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
	// Specifies whether to enable case sensitivity. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	CaseSensitive *bool `json:"caseSensitive,omitempty" xml:"caseSensitive,omitempty"`
	// Specifies whether to include Chinese characters. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// This parameter is required.
	//
	// example:
	//
	// false
	Chn *bool `json:"chn,omitempty" xml:"chn,omitempty"`
	// The excluded fields. You cannot specify both include_keys and exclude_keys.
	ExcludeKeys []*string `json:"exclude_keys,omitempty" xml:"exclude_keys,omitempty" type:"Repeated"`
	// The included fields. You cannot specify both include_keys and exclude_keys.
	IncludeKeys []*string `json:"include_keys,omitempty" xml:"include_keys,omitempty" type:"Repeated"`
	// The delimiters that are used to split text.
	//
	// This parameter is required.
	Token []*string `json:"token,omitempty" xml:"token,omitempty" type:"Repeated"`
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
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
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

type UpdateLogStoreRequest struct {
	// Specifies whether to record public IP addresses. Default value: false. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// false
	AppendMeta *bool `json:"appendMeta,omitempty" xml:"appendMeta,omitempty"`
	// Specifies whether to enable automatic sharding. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	AutoSplit *bool `json:"autoSplit,omitempty" xml:"autoSplit,omitempty"`
	// Specifies whether to enable the web tracking feature. Default value: false. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// false
	EnableTracking *bool `json:"enable_tracking,omitempty" xml:"enable_tracking,omitempty"`
	// The data structure of the encryption configuration.
	EncryptConf *EncryptConf `json:"encrypt_conf,omitempty" xml:"encrypt_conf,omitempty"`
	// The retention period of data in the hot storage tier of the Logstore. Minimum value: 30. Unit: day. You can specify a value that ranges from 30 to the value of ttl. Hot data that is stored for longer than the period specified by hot_ttl is converted to cold data. For more information, see [Enable hot and cold-tiered storage for a Logstore](https://help.aliyun.com/document_detail/308645.html).
	//
	// example:
	//
	// 60
	HotTtl              *int32 `json:"hot_ttl,omitempty" xml:"hot_ttl,omitempty"`
	InfrequentAccessTTL *int32 `json:"infrequentAccessTTL,omitempty" xml:"infrequentAccessTTL,omitempty"`
	// The name of the Logstore.
	//
	// This parameter is required.
	//
	// example:
	//
	// test-logstore
	LogstoreName *string `json:"logstoreName,omitempty" xml:"logstoreName,omitempty"`
	// The maximum number of shards into which existing shards can be automatically split. Valid values: 1 to 64.
	//
	// > If you set autoSplit to true, you must specify maxSplitShard.
	//
	// example:
	//
	// 64
	MaxSplitShard *int32 `json:"maxSplitShard,omitempty" xml:"maxSplitShard,omitempty"`
	// The type of the Logstore. Simple Log Service provides two types of Logstores: Standard Logstores and Query Logstores.
	//
	// 	- **standard**: Standard Logstore. This type of Logstore supports the log analysis feature and is suitable for scenarios such as real-time monitoring and interactive analysis. You can also use this type of Logstore to build a comprehensive observability system.
	//
	// 	- **query**: Query Logstore. This type of Logstore supports high-performance queries. The index traffic fee of a Query Logstore is approximately half that of a Standard Logstore. Query Logstores do not support SQL analysis. Query Logstores are suitable for scenarios in which the volume of data is large, the log retention period is long, or log analysis is not required. Log retention periods of weeks or months are considered long.
	//
	// example:
	//
	// standard
	Mode *string `json:"mode,omitempty" xml:"mode,omitempty"`
	// Deprecated
	//
	// The number of shards.
	//
	// > You cannot call the UpdateLogstore operation to change the number of shards. You can call the SplitShard or MergeShards operation to change the number of shards.
	//
	// example:
	//
	// 2
	ShardCount *int32 `json:"shardCount,omitempty" xml:"shardCount,omitempty"`
	// Deprecated
	//
	// The type of the log that you want to query. Valid values:
	//
	// 	- None: all types of logs.
	//
	// 	- Metrics: metrics.
	//
	// example:
	//
	// None
	TelemetryType *string `json:"telemetryType,omitempty" xml:"telemetryType,omitempty"`
	// The retention period of data. Unit: day. Valid values: 1 to 3650. If you set ttl to 3650, data is permanently stored.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30
	Ttl *int32 `json:"ttl,omitempty" xml:"ttl,omitempty"`
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

func (s *UpdateLogStoreRequest) SetInfrequentAccessTTL(v int32) *UpdateLogStoreRequest {
	s.InfrequentAccessTTL = &v
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
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
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

type UpdateLogStoreMeteringModeRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ChargeByFunction
	MeteringMode *string `json:"meteringMode,omitempty" xml:"meteringMode,omitempty"`
}

func (s UpdateLogStoreMeteringModeRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateLogStoreMeteringModeRequest) GoString() string {
	return s.String()
}

func (s *UpdateLogStoreMeteringModeRequest) SetMeteringMode(v string) *UpdateLogStoreMeteringModeRequest {
	s.MeteringMode = &v
	return s
}

type UpdateLogStoreMeteringModeResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s UpdateLogStoreMeteringModeResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateLogStoreMeteringModeResponse) GoString() string {
	return s.String()
}

func (s *UpdateLogStoreMeteringModeResponse) SetHeaders(v map[string]*string) *UpdateLogStoreMeteringModeResponse {
	s.Headers = v
	return s
}

func (s *UpdateLogStoreMeteringModeResponse) SetStatusCode(v int32) *UpdateLogStoreMeteringModeResponse {
	s.StatusCode = &v
	return s
}

type UpdateLoggingRequest struct {
	// The configurations of service logs.
	//
	// This parameter is required.
	LoggingDetails []*UpdateLoggingRequestLoggingDetails `json:"loggingDetails,omitempty" xml:"loggingDetails,omitempty" type:"Repeated"`
	// The name of the project to which you want to save service logs.
	//
	// This parameter is required.
	//
	// example:
	//
	// my-project
	LoggingProject *string `json:"loggingProject,omitempty" xml:"loggingProject,omitempty"`
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
	// The name of the Logstore to which you want to save service logs.
	//
	// This parameter is required.
	//
	// example:
	//
	// my-logstore
	Logstore *string `json:"logstore,omitempty" xml:"logstore,omitempty"`
	// The type of service logs. Valid values:
	//
	// 	- consumergroup_log: the consumption delay logs of consumer groups.
	//
	// 	- logtail_alarm: the alert logs of Logtail.
	//
	// 	- operation_log: the operation logs.
	//
	// 	- logtail_profile: the collection logs of Logtail.
	//
	// 	- metering: the metering logs.
	//
	// 	- logtail_status: the status logs of Logtail.
	//
	// 	- scheduledsqlalert: the operational logs of Scheduled SQL jobs.
	//
	// 	- etl_alert: the operational logs of data transformation jobs.
	//
	// This parameter is required.
	//
	// example:
	//
	// consumergroup_log
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
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
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
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

type UpdateLogtailPipelineConfigRequest struct {
	// The aggregation plug-ins.
	Aggregators []map[string]interface{} `json:"aggregators,omitempty" xml:"aggregators,omitempty" type:"Repeated"`
	// The name of the configuration.
	//
	// This parameter is required.
	//
	// example:
	//
	// test-config
	ConfigName *string `json:"configName,omitempty" xml:"configName,omitempty"`
	// The data output plug-ins.
	//
	// This parameter is required.
	Flushers []map[string]interface{} `json:"flushers,omitempty" xml:"flushers,omitempty" type:"Repeated"`
	// The global configuration.
	Global map[string]interface{} `json:"global,omitempty" xml:"global,omitempty"`
	// The data source plug-ins.
	//
	// This parameter is required.
	Inputs []map[string]interface{} `json:"inputs,omitempty" xml:"inputs,omitempty" type:"Repeated"`
	// The sample log.
	//
	// example:
	//
	// 2022-06-14 11:13:29.796 | DEBUG    | __main__:<module>:1 - hello world
	LogSample *string `json:"logSample,omitempty" xml:"logSample,omitempty"`
	// The processing plug-ins.
	Processors []map[string]interface{} `json:"processors,omitempty" xml:"processors,omitempty" type:"Repeated"`
}

func (s UpdateLogtailPipelineConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateLogtailPipelineConfigRequest) GoString() string {
	return s.String()
}

func (s *UpdateLogtailPipelineConfigRequest) SetAggregators(v []map[string]interface{}) *UpdateLogtailPipelineConfigRequest {
	s.Aggregators = v
	return s
}

func (s *UpdateLogtailPipelineConfigRequest) SetConfigName(v string) *UpdateLogtailPipelineConfigRequest {
	s.ConfigName = &v
	return s
}

func (s *UpdateLogtailPipelineConfigRequest) SetFlushers(v []map[string]interface{}) *UpdateLogtailPipelineConfigRequest {
	s.Flushers = v
	return s
}

func (s *UpdateLogtailPipelineConfigRequest) SetGlobal(v map[string]interface{}) *UpdateLogtailPipelineConfigRequest {
	s.Global = v
	return s
}

func (s *UpdateLogtailPipelineConfigRequest) SetInputs(v []map[string]interface{}) *UpdateLogtailPipelineConfigRequest {
	s.Inputs = v
	return s
}

func (s *UpdateLogtailPipelineConfigRequest) SetLogSample(v string) *UpdateLogtailPipelineConfigRequest {
	s.LogSample = &v
	return s
}

func (s *UpdateLogtailPipelineConfigRequest) SetProcessors(v []map[string]interface{}) *UpdateLogtailPipelineConfigRequest {
	s.Processors = v
	return s
}

type UpdateLogtailPipelineConfigResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s UpdateLogtailPipelineConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateLogtailPipelineConfigResponse) GoString() string {
	return s.String()
}

func (s *UpdateLogtailPipelineConfigResponse) SetHeaders(v map[string]*string) *UpdateLogtailPipelineConfigResponse {
	s.Headers = v
	return s
}

func (s *UpdateLogtailPipelineConfigResponse) SetStatusCode(v int32) *UpdateLogtailPipelineConfigResponse {
	s.StatusCode = &v
	return s
}

type UpdateMachineGroupRequest struct {
	// The attribute of the machine group. This parameter is empty by default.
	GroupAttribute *UpdateMachineGroupRequestGroupAttribute `json:"groupAttribute,omitempty" xml:"groupAttribute,omitempty" type:"Struct"`
	// The name of the machine group.
	//
	// This parameter is required.
	//
	// example:
	//
	// test-machine-group
	GroupName *string `json:"groupName,omitempty" xml:"groupName,omitempty"`
	// The type of the machine group. Set the value to an empty string.
	//
	// example:
	//
	// ""
	GroupType *string `json:"groupType,omitempty" xml:"groupType,omitempty"`
	// The identifier type of the machine group. Valid values:
	//
	// 	- ip: The machine group uses IP addresses as identifiers.
	//
	// 	- userdefined: The machine group uses custom identifiers.
	//
	// This parameter is required.
	//
	// example:
	//
	// userdefined
	MachineIdentifyType *string `json:"machineIdentifyType,omitempty" xml:"machineIdentifyType,omitempty"`
	// The identifiers of the machines in the machine group.
	//
	// 	- If you set machineIdentifyType to ip, enter the IP addresses of the machines.
	//
	// 	- If you set machineIdentifyType to userdefined, enter a custom identifier.
	//
	// This parameter is required.
	//
	// example:
	//
	// [uu_id_1，uu_id_2]
	MachineList []*string `json:"machineList,omitempty" xml:"machineList,omitempty" type:"Repeated"`
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
	// The identifier of the external management system on which the machine group depends. This parameter is empty by default.
	//
	// example:
	//
	// testgroup2
	ExternalName *string `json:"externalName,omitempty" xml:"externalName,omitempty"`
	// The topic of the machine group. This parameter is empty by default.
	//
	// example:
	//
	// testtopic2
	GroupTopic *string `json:"groupTopic,omitempty" xml:"groupTopic,omitempty"`
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
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
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
	// The operation on the machine. Valid values: add and delete. A value of add specifies to add the machine to the machine group. A value of delete specifies to remove the machine from the machine group.
	//
	// example:
	//
	// add
	Action *string `json:"action,omitempty" xml:"action,omitempty"`
	// The machines to be added or removed.
	//
	// This parameter is required.
	Body []*string `json:"body,omitempty" xml:"body,omitempty" type:"Repeated"`
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
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
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

type UpdateMetricStoreMeteringModeRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ChargeByFunction
	MeteringMode *string `json:"meteringMode,omitempty" xml:"meteringMode,omitempty"`
}

func (s UpdateMetricStoreMeteringModeRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateMetricStoreMeteringModeRequest) GoString() string {
	return s.String()
}

func (s *UpdateMetricStoreMeteringModeRequest) SetMeteringMode(v string) *UpdateMetricStoreMeteringModeRequest {
	s.MeteringMode = &v
	return s
}

type UpdateMetricStoreMeteringModeResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s UpdateMetricStoreMeteringModeResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateMetricStoreMeteringModeResponse) GoString() string {
	return s.String()
}

func (s *UpdateMetricStoreMeteringModeResponse) SetHeaders(v map[string]*string) *UpdateMetricStoreMeteringModeResponse {
	s.Headers = v
	return s
}

func (s *UpdateMetricStoreMeteringModeResponse) SetStatusCode(v int32) *UpdateMetricStoreMeteringModeResponse {
	s.StatusCode = &v
	return s
}

type UpdateOSSExportRequest struct {
	Configuration *OSSExportConfiguration `json:"configuration,omitempty" xml:"configuration,omitempty"`
	Description   *string                 `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// ali-test-oss-job
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
}

func (s UpdateOSSExportRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateOSSExportRequest) GoString() string {
	return s.String()
}

func (s *UpdateOSSExportRequest) SetConfiguration(v *OSSExportConfiguration) *UpdateOSSExportRequest {
	s.Configuration = v
	return s
}

func (s *UpdateOSSExportRequest) SetDescription(v string) *UpdateOSSExportRequest {
	s.Description = &v
	return s
}

func (s *UpdateOSSExportRequest) SetDisplayName(v string) *UpdateOSSExportRequest {
	s.DisplayName = &v
	return s
}

type UpdateOSSExportResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s UpdateOSSExportResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateOSSExportResponse) GoString() string {
	return s.String()
}

func (s *UpdateOSSExportResponse) SetHeaders(v map[string]*string) *UpdateOSSExportResponse {
	s.Headers = v
	return s
}

func (s *UpdateOSSExportResponse) SetStatusCode(v int32) *UpdateOSSExportResponse {
	s.StatusCode = &v
	return s
}

type UpdateOSSHDFSExportRequest struct {
	Configuration *OSSExportConfiguration `json:"configuration,omitempty" xml:"configuration,omitempty"`
	Description   *string                 `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// ali-test-oss-hdfs-job
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
}

func (s UpdateOSSHDFSExportRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateOSSHDFSExportRequest) GoString() string {
	return s.String()
}

func (s *UpdateOSSHDFSExportRequest) SetConfiguration(v *OSSExportConfiguration) *UpdateOSSHDFSExportRequest {
	s.Configuration = v
	return s
}

func (s *UpdateOSSHDFSExportRequest) SetDescription(v string) *UpdateOSSHDFSExportRequest {
	s.Description = &v
	return s
}

func (s *UpdateOSSHDFSExportRequest) SetDisplayName(v string) *UpdateOSSHDFSExportRequest {
	s.DisplayName = &v
	return s
}

type UpdateOSSHDFSExportResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s UpdateOSSHDFSExportResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateOSSHDFSExportResponse) GoString() string {
	return s.String()
}

func (s *UpdateOSSHDFSExportResponse) SetHeaders(v map[string]*string) *UpdateOSSHDFSExportResponse {
	s.Headers = v
	return s
}

func (s *UpdateOSSHDFSExportResponse) SetStatusCode(v int32) *UpdateOSSHDFSExportResponse {
	s.StatusCode = &v
	return s
}

type UpdateOSSIngestionRequest struct {
	// This parameter is required.
	Configuration *OSSIngestionConfiguration `json:"configuration,omitempty" xml:"configuration,omitempty"`
	Description   *string                    `json:"description,omitempty" xml:"description,omitempty"`
	// This parameter is required.
	DisplayName *string   `json:"displayName,omitempty" xml:"displayName,omitempty"`
	Schedule    *Schedule `json:"schedule,omitempty" xml:"schedule,omitempty"`
}

func (s UpdateOSSIngestionRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateOSSIngestionRequest) GoString() string {
	return s.String()
}

func (s *UpdateOSSIngestionRequest) SetConfiguration(v *OSSIngestionConfiguration) *UpdateOSSIngestionRequest {
	s.Configuration = v
	return s
}

func (s *UpdateOSSIngestionRequest) SetDescription(v string) *UpdateOSSIngestionRequest {
	s.Description = &v
	return s
}

func (s *UpdateOSSIngestionRequest) SetDisplayName(v string) *UpdateOSSIngestionRequest {
	s.DisplayName = &v
	return s
}

func (s *UpdateOSSIngestionRequest) SetSchedule(v *Schedule) *UpdateOSSIngestionRequest {
	s.Schedule = v
	return s
}

type UpdateOSSIngestionResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s UpdateOSSIngestionResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateOSSIngestionResponse) GoString() string {
	return s.String()
}

func (s *UpdateOSSIngestionResponse) SetHeaders(v map[string]*string) *UpdateOSSIngestionResponse {
	s.Headers = v
	return s
}

func (s *UpdateOSSIngestionResponse) SetStatusCode(v int32) *UpdateOSSIngestionResponse {
	s.StatusCode = &v
	return s
}

type UpdateOssExternalStoreRequest struct {
	// The name of the external store.
	//
	// This parameter is required.
	//
	// example:
	//
	// test-oss-store
	ExternalStoreName *string `json:"externalStoreName,omitempty" xml:"externalStoreName,omitempty"`
	// The parameters that are configured for the external store.
	//
	// This parameter is required.
	Parameter *UpdateOssExternalStoreRequestParameter `json:"parameter,omitempty" xml:"parameter,omitempty" type:"Struct"`
	// The type of the external store. Set the value to oss.
	//
	// This parameter is required.
	//
	// example:
	//
	// oss
	StoreType *string `json:"storeType,omitempty" xml:"storeType,omitempty"`
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
	// The AccessKey ID of your account.
	//
	// This parameter is required.
	//
	// example:
	//
	// LTAI5tFsHGGeYry*****1Sz
	Accessid *string `json:"accessid,omitempty" xml:"accessid,omitempty"`
	// The AccessKey secret of your account.
	//
	// This parameter is required.
	//
	// example:
	//
	// GyviCLDVHkHrOztdkxuE6******Rp6
	Accesskey *string `json:"accesskey,omitempty" xml:"accesskey,omitempty"`
	// The name of the OSS bucket.
	//
	// This parameter is required.
	//
	// example:
	//
	// my-bucket
	Bucket *string `json:"bucket,omitempty" xml:"bucket,omitempty"`
	// The fields that are associated to the external store.
	//
	// This parameter is required.
	Columns []*UpdateOssExternalStoreRequestParameterColumns `json:"columns,omitempty" xml:"columns,omitempty" type:"Repeated"`
	// The Object Storage Service (OSS) endpoint.
	//
	// This parameter is required.
	//
	// example:
	//
	// oss-cn-hangzhou.aliyuncs.com
	Endpoint *string `json:"endpoint,omitempty" xml:"endpoint,omitempty"`
	// The names of the OSS objects that are associated to the external store.
	//
	// This parameter is required.
	Objects []*string `json:"objects,omitempty" xml:"objects,omitempty" type:"Repeated"`
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
	// The name of the field.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The type of the field.
	//
	// This parameter is required.
	//
	// example:
	//
	// varchar
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
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
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

type UpdateProjectRequest struct {
	// The description of the project. The default value is an empty string.
	//
	// This parameter is required.
	//
	// example:
	//
	// Description of my-project-test
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
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
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
	// The name of the external store.
	//
	// This parameter is required.
	//
	// example:
	//
	// rds_store
	ExternalStoreName *string `json:"externalStoreName,omitempty" xml:"externalStoreName,omitempty"`
	// The parameter struct.
	//
	// This parameter is required.
	Parameter *UpdateRdsExternalStoreRequestParameter `json:"parameter,omitempty" xml:"parameter,omitempty" type:"Struct"`
	// The storage type. Set the value to rds-vpc, which indicates an ApsaraDB RDS for MySQL database in a virtual private cloud (VPC).
	//
	// This parameter is required.
	//
	// example:
	//
	// rds-vpc
	StoreType *string `json:"storeType,omitempty" xml:"storeType,omitempty"`
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
	// The name of the database in the ApsaraDB RDS for MySQL instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// meta
	Db *string `json:"db,omitempty" xml:"db,omitempty"`
	// The internal or public endpoint of the ApsaraDB RDS for MySQL instance.
	//
	// example:
	//
	// 192.168.XX.XX
	Host *string `json:"host,omitempty" xml:"host,omitempty"`
	// The ID of the ApsaraDB RDS for MySQL instance.
	//
	// example:
	//
	// i-bp1b6c719dfa08exf****
	InstanceId *string `json:"instance-id,omitempty" xml:"instance-id,omitempty"`
	// The password that is used to log on to the ApsaraDB RDS for MySQL instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// sfdsfldsfksfls****
	Password *string `json:"password,omitempty" xml:"password,omitempty"`
	// The internal or public port of the ApsaraDB RDS for MySQL instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3306
	Port *string `json:"port,omitempty" xml:"port,omitempty"`
	// The region where the ApsaraDB RDS for MySQL instance resides. Valid values: cn-qingdao, cn-beijing, and cn-hangzhou.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-qingdao
	Region *string `json:"region,omitempty" xml:"region,omitempty"`
	// The name of the database table in the ApsaraDB RDS for MySQL instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// join_meta
	Table *string `json:"table,omitempty" xml:"table,omitempty"`
	// The username that is used to log on to the ApsaraDB RDS for MySQL instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// root
	Username *string `json:"username,omitempty" xml:"username,omitempty"`
	// The ID of the VPC to which the ApsaraDB RDS for MySQL instance belongs.
	//
	// example:
	//
	// vpc-bp1aevy8sofi8mh1q****
	VpcId *string `json:"vpc-id,omitempty" xml:"vpc-id,omitempty"`
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
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
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
	// The display name.
	//
	// This parameter is required.
	//
	// example:
	//
	// displayname
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	// The name of the Logstore to which the saved search belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// aliyun-test-logstore
	Logstore *string `json:"logstore,omitempty" xml:"logstore,omitempty"`
	// The name of the saved search. The name must be 3 to 63 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// savedsearch-name
	SavedsearchName *string `json:"savedsearchName,omitempty" xml:"savedsearchName,omitempty"`
	// The search statement or the query statement of the saved search. A query statement consists of a search statement and an analytic statement in the Search statement|Analytic statement format.
	//
	// For more information, see Log search overview and Log analysis overview.
	//
	// This parameter is required.
	//
	// example:
	//
	// *|select date_format(__time__-__time__%60, \\"%H:%i:%s\\") as time, COUNT(*) as pv group by time
	SearchQuery *string `json:"searchQuery,omitempty" xml:"searchQuery,omitempty"`
	// The topic of the logs.
	//
	// example:
	//
	// theme
	Topic *string `json:"topic,omitempty" xml:"topic,omitempty"`
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
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
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

type UpdateScheduledSQLRequest struct {
	// This parameter is required.
	Configuration *ScheduledSQLConfiguration `json:"configuration,omitempty" xml:"configuration,omitempty"`
	Description   *string                    `json:"description,omitempty" xml:"description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ali-test-scheduled-sql
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	// This parameter is required.
	Schedule *Schedule `json:"schedule,omitempty" xml:"schedule,omitempty"`
}

func (s UpdateScheduledSQLRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateScheduledSQLRequest) GoString() string {
	return s.String()
}

func (s *UpdateScheduledSQLRequest) SetConfiguration(v *ScheduledSQLConfiguration) *UpdateScheduledSQLRequest {
	s.Configuration = v
	return s
}

func (s *UpdateScheduledSQLRequest) SetDescription(v string) *UpdateScheduledSQLRequest {
	s.Description = &v
	return s
}

func (s *UpdateScheduledSQLRequest) SetDisplayName(v string) *UpdateScheduledSQLRequest {
	s.DisplayName = &v
	return s
}

func (s *UpdateScheduledSQLRequest) SetSchedule(v *Schedule) *UpdateScheduledSQLRequest {
	s.Schedule = v
	return s
}

type UpdateScheduledSQLResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s UpdateScheduledSQLResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateScheduledSQLResponse) GoString() string {
	return s.String()
}

func (s *UpdateScheduledSQLResponse) SetHeaders(v map[string]*string) *UpdateScheduledSQLResponse {
	s.Headers = v
	return s
}

func (s *UpdateScheduledSQLResponse) SetStatusCode(v int32) *UpdateScheduledSQLResponse {
	s.StatusCode = &v
	return s
}

type UpdateSqlInstanceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 2
	Cu *int32 `json:"cu,omitempty" xml:"cu,omitempty"`
	// This parameter is required.
	UseAsDefault *bool `json:"useAsDefault,omitempty" xml:"useAsDefault,omitempty"`
}

func (s UpdateSqlInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateSqlInstanceRequest) GoString() string {
	return s.String()
}

func (s *UpdateSqlInstanceRequest) SetCu(v int32) *UpdateSqlInstanceRequest {
	s.Cu = &v
	return s
}

func (s *UpdateSqlInstanceRequest) SetUseAsDefault(v bool) *UpdateSqlInstanceRequest {
	s.UseAsDefault = &v
	return s
}

type UpdateSqlInstanceResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s UpdateSqlInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateSqlInstanceResponse) GoString() string {
	return s.String()
}

func (s *UpdateSqlInstanceResponse) SetHeaders(v map[string]*string) *UpdateSqlInstanceResponse {
	s.Headers = v
	return s
}

func (s *UpdateSqlInstanceResponse) SetStatusCode(v int32) *UpdateSqlInstanceResponse {
	s.StatusCode = &v
	return s
}

type UpdateStoreViewRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// logstore
	StoreType *string `json:"storeType,omitempty" xml:"storeType,omitempty"`
	// This parameter is required.
	Stores []*StoreViewStore `json:"stores,omitempty" xml:"stores,omitempty" type:"Repeated"`
}

func (s UpdateStoreViewRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateStoreViewRequest) GoString() string {
	return s.String()
}

func (s *UpdateStoreViewRequest) SetStoreType(v string) *UpdateStoreViewRequest {
	s.StoreType = &v
	return s
}

func (s *UpdateStoreViewRequest) SetStores(v []*StoreViewStore) *UpdateStoreViewRequest {
	s.Stores = v
	return s
}

type UpdateStoreViewResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s UpdateStoreViewResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateStoreViewResponse) GoString() string {
	return s.String()
}

func (s *UpdateStoreViewResponse) SetHeaders(v map[string]*string) *UpdateStoreViewResponse {
	s.Headers = v
	return s
}

func (s *UpdateStoreViewResponse) SetStatusCode(v int32) *UpdateStoreViewResponse {
	s.StatusCode = &v
	return s
}

type UpsertCollectionPolicyRequest struct {
	Attribute        *UpsertCollectionPolicyRequestAttribute        `json:"attribute,omitempty" xml:"attribute,omitempty" type:"Struct"`
	CentralizeConfig *UpsertCollectionPolicyRequestCentralizeConfig `json:"centralizeConfig,omitempty" xml:"centralizeConfig,omitempty" type:"Struct"`
	// example:
	//
	// false
	CentralizeEnabled *bool `json:"centralizeEnabled,omitempty" xml:"centralizeEnabled,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// access_log
	DataCode *string `json:"dataCode,omitempty" xml:"dataCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// true
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// This parameter is required.
	PolicyConfig *UpsertCollectionPolicyRequestPolicyConfig `json:"policyConfig,omitempty" xml:"policyConfig,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// your_log_policy
	PolicyName *string `json:"policyName,omitempty" xml:"policyName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// oss
	ProductCode *string `json:"productCode,omitempty" xml:"productCode,omitempty"`
}

func (s UpsertCollectionPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s UpsertCollectionPolicyRequest) GoString() string {
	return s.String()
}

func (s *UpsertCollectionPolicyRequest) SetAttribute(v *UpsertCollectionPolicyRequestAttribute) *UpsertCollectionPolicyRequest {
	s.Attribute = v
	return s
}

func (s *UpsertCollectionPolicyRequest) SetCentralizeConfig(v *UpsertCollectionPolicyRequestCentralizeConfig) *UpsertCollectionPolicyRequest {
	s.CentralizeConfig = v
	return s
}

func (s *UpsertCollectionPolicyRequest) SetCentralizeEnabled(v bool) *UpsertCollectionPolicyRequest {
	s.CentralizeEnabled = &v
	return s
}

func (s *UpsertCollectionPolicyRequest) SetDataCode(v string) *UpsertCollectionPolicyRequest {
	s.DataCode = &v
	return s
}

func (s *UpsertCollectionPolicyRequest) SetEnabled(v bool) *UpsertCollectionPolicyRequest {
	s.Enabled = &v
	return s
}

func (s *UpsertCollectionPolicyRequest) SetPolicyConfig(v *UpsertCollectionPolicyRequestPolicyConfig) *UpsertCollectionPolicyRequest {
	s.PolicyConfig = v
	return s
}

func (s *UpsertCollectionPolicyRequest) SetPolicyName(v string) *UpsertCollectionPolicyRequest {
	s.PolicyName = &v
	return s
}

func (s *UpsertCollectionPolicyRequest) SetProductCode(v string) *UpsertCollectionPolicyRequest {
	s.ProductCode = &v
	return s
}

type UpsertCollectionPolicyRequestAttribute struct {
	// example:
	//
	// your-app-name
	App *string `json:"app,omitempty" xml:"app,omitempty"`
	// example:
	//
	// your-policy-group
	PolicyGroup *string `json:"policyGroup,omitempty" xml:"policyGroup,omitempty"`
}

func (s UpsertCollectionPolicyRequestAttribute) String() string {
	return tea.Prettify(s)
}

func (s UpsertCollectionPolicyRequestAttribute) GoString() string {
	return s.String()
}

func (s *UpsertCollectionPolicyRequestAttribute) SetApp(v string) *UpsertCollectionPolicyRequestAttribute {
	s.App = &v
	return s
}

func (s *UpsertCollectionPolicyRequestAttribute) SetPolicyGroup(v string) *UpsertCollectionPolicyRequestAttribute {
	s.PolicyGroup = &v
	return s
}

type UpsertCollectionPolicyRequestCentralizeConfig struct {
	// example:
	//
	// your-sls-logstore-in-beijing
	DestLogstore *string `json:"destLogstore,omitempty" xml:"destLogstore,omitempty"`
	// example:
	//
	// your-sls-project-in-beijing
	DestProject *string `json:"destProject,omitempty" xml:"destProject,omitempty"`
	// example:
	//
	// cn-beijing
	DestRegion *string `json:"destRegion,omitempty" xml:"destRegion,omitempty"`
	// example:
	//
	// your-sls-logstore-ttl
	DestTTL *int32 `json:"destTTL,omitempty" xml:"destTTL,omitempty"`
}

func (s UpsertCollectionPolicyRequestCentralizeConfig) String() string {
	return tea.Prettify(s)
}

func (s UpsertCollectionPolicyRequestCentralizeConfig) GoString() string {
	return s.String()
}

func (s *UpsertCollectionPolicyRequestCentralizeConfig) SetDestLogstore(v string) *UpsertCollectionPolicyRequestCentralizeConfig {
	s.DestLogstore = &v
	return s
}

func (s *UpsertCollectionPolicyRequestCentralizeConfig) SetDestProject(v string) *UpsertCollectionPolicyRequestCentralizeConfig {
	s.DestProject = &v
	return s
}

func (s *UpsertCollectionPolicyRequestCentralizeConfig) SetDestRegion(v string) *UpsertCollectionPolicyRequestCentralizeConfig {
	s.DestRegion = &v
	return s
}

func (s *UpsertCollectionPolicyRequestCentralizeConfig) SetDestTTL(v int32) *UpsertCollectionPolicyRequestCentralizeConfig {
	s.DestTTL = &v
	return s
}

type UpsertCollectionPolicyRequestPolicyConfig struct {
	InstanceIds []*string `json:"instanceIds,omitempty" xml:"instanceIds,omitempty" type:"Repeated"`
	Regions     []*string `json:"regions,omitempty" xml:"regions,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// all
	ResourceMode *string                `json:"resourceMode,omitempty" xml:"resourceMode,omitempty"`
	ResourceTags map[string]interface{} `json:"resourceTags,omitempty" xml:"resourceTags,omitempty"`
}

func (s UpsertCollectionPolicyRequestPolicyConfig) String() string {
	return tea.Prettify(s)
}

func (s UpsertCollectionPolicyRequestPolicyConfig) GoString() string {
	return s.String()
}

func (s *UpsertCollectionPolicyRequestPolicyConfig) SetInstanceIds(v []*string) *UpsertCollectionPolicyRequestPolicyConfig {
	s.InstanceIds = v
	return s
}

func (s *UpsertCollectionPolicyRequestPolicyConfig) SetRegions(v []*string) *UpsertCollectionPolicyRequestPolicyConfig {
	s.Regions = v
	return s
}

func (s *UpsertCollectionPolicyRequestPolicyConfig) SetResourceMode(v string) *UpsertCollectionPolicyRequestPolicyConfig {
	s.ResourceMode = &v
	return s
}

func (s *UpsertCollectionPolicyRequestPolicyConfig) SetResourceTags(v map[string]interface{}) *UpsertCollectionPolicyRequestPolicyConfig {
	s.ResourceTags = v
	return s
}

type UpsertCollectionPolicyResponseBody struct {
	// example:
	//
	// successful
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
}

func (s UpsertCollectionPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpsertCollectionPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *UpsertCollectionPolicyResponseBody) SetMessage(v string) *UpsertCollectionPolicyResponseBody {
	s.Message = &v
	return s
}

type UpsertCollectionPolicyResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpsertCollectionPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpsertCollectionPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s UpsertCollectionPolicyResponse) GoString() string {
	return s.String()
}

func (s *UpsertCollectionPolicyResponse) SetHeaders(v map[string]*string) *UpsertCollectionPolicyResponse {
	s.Headers = v
	return s
}

func (s *UpsertCollectionPolicyResponse) SetStatusCode(v int32) *UpsertCollectionPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *UpsertCollectionPolicyResponse) SetBody(v *UpsertCollectionPolicyResponseBody) *UpsertCollectionPolicyResponse {
	s.Body = v
	return s
}

type Client struct {
	openapi.Client
	Client_ spi.Client
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
	client.SignatureAlgorithm = tea.String("v2")
	client.EndpointRule = tea.String("central")
	return nil
}

// Summary:
//
// Applies a Logtail configuration to a machine group.
//
// Description:
//
// Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ApplyConfigToMachineGroupResponse
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

// Summary:
//
// Applies a Logtail configuration to a machine group.
//
// Description:
//
// Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
//
// @return ApplyConfigToMachineGroupResponse
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

// Summary:
//
// Modifies the resource group of a resource.
//
// @param request - ChangeResourceGroupRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChangeResourceGroupResponse
func (client *Client) ChangeResourceGroupWithOptions(project *string, request *ChangeResourceGroupRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ChangeResourceGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["project"] = project
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
		HostMap: hostMap,
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

// Summary:
//
// Modifies the resource group of a resource.
//
// @param request - ChangeResourceGroupRequest
//
// @return ChangeResourceGroupResponse
func (client *Client) ChangeResourceGroup(project *string, request *ChangeResourceGroupRequest) (_result *ChangeResourceGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ChangeResourceGroupResponse{}
	_body, _err := client.ChangeResourceGroupWithOptions(project, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Sends heartbeats to a server from a consumer.
//
// Description:
//
// ### Usage notes
//
// 	- Connections between consumers and servers are established by sending heartbeats at regular intervals. If a server does not receive heartbeats from a consumer on schedule, the server deletes the consumer.
//
// 	- Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
//
// @param request - ConsumerGroupHeartBeatRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ConsumerGroupHeartBeatResponse
func (client *Client) ConsumerGroupHeartBeatWithOptions(project *string, logstore *string, consumerGroup *string, request *ConsumerGroupHeartBeatRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ConsumerGroupHeartBeatResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Consumer)) {
		query["consumer"] = request.Consumer
	}

	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    request.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("ConsumerGroupHeartBeat"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/logstores/" + tea.StringValue(logstore) + "/consumergroups/" + tea.StringValue(consumerGroup) + "?type=heartbeat"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("array"),
	}
	_result = &ConsumerGroupHeartBeatResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Sends heartbeats to a server from a consumer.
//
// Description:
//
// ### Usage notes
//
// 	- Connections between consumers and servers are established by sending heartbeats at regular intervals. If a server does not receive heartbeats from a consumer on schedule, the server deletes the consumer.
//
// 	- Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
//
// @param request - ConsumerGroupHeartBeatRequest
//
// @return ConsumerGroupHeartBeatResponse
func (client *Client) ConsumerGroupHeartBeat(project *string, logstore *string, consumerGroup *string, request *ConsumerGroupHeartBeatRequest) (_result *ConsumerGroupHeartBeatResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ConsumerGroupHeartBeatResponse{}
	_body, _err := client.ConsumerGroupHeartBeatWithOptions(project, logstore, consumerGroup, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新指定消费组消费数据时Shard的checkpoint。
//
// @param request - ConsumerGroupUpdateCheckPointRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ConsumerGroupUpdateCheckPointResponse
func (client *Client) ConsumerGroupUpdateCheckPointWithOptions(project *string, logstore *string, consumerGroup *string, request *ConsumerGroupUpdateCheckPointRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ConsumerGroupUpdateCheckPointResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Consumer)) {
		query["consumer"] = request.Consumer
	}

	if !tea.BoolValue(util.IsUnset(request.ForceSuccess)) {
		query["forceSuccess"] = request.ForceSuccess
	}

	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    util.ToArray(request.Body),
	}
	params := &openapi.Params{
		Action:      tea.String("ConsumerGroupUpdateCheckPoint"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/logstores/" + tea.StringValue(logstore) + "/consumergroups/" + tea.StringValue(consumerGroup) + "?type=checkpoint"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &ConsumerGroupUpdateCheckPointResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新指定消费组消费数据时Shard的checkpoint。
//
// @param request - ConsumerGroupUpdateCheckPointRequest
//
// @return ConsumerGroupUpdateCheckPointResponse
func (client *Client) ConsumerGroupUpdateCheckPoint(project *string, logstore *string, consumerGroup *string, request *ConsumerGroupUpdateCheckPointRequest) (_result *ConsumerGroupUpdateCheckPointResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ConsumerGroupUpdateCheckPointResponse{}
	_body, _err := client.ConsumerGroupUpdateCheckPointWithOptions(project, logstore, consumerGroup, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// CreateAlert
//
// @param request - CreateAlertRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAlertResponse
func (client *Client) CreateAlertWithOptions(project *string, request *CreateAlertRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateAlertResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Configuration)) {
		body["configuration"] = request.Configuration
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.DisplayName)) {
		body["displayName"] = request.DisplayName
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Schedule)) {
		body["schedule"] = request.Schedule
	}

	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateAlert"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/alerts"),
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

// Summary:
//
// CreateAlert
//
// @param request - CreateAlertRequest
//
// @return CreateAlertResponse
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

// Summary:
//
// Creates a dataset.
//
// @param request - CreateAnnotationDataSetRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAnnotationDataSetResponse
func (client *Client) CreateAnnotationDataSetWithOptions(request *CreateAnnotationDataSetRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateAnnotationDataSetResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DatasetId)) {
		query["datasetId"] = request.DatasetId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateAnnotationDataSet"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/ml/annotationdataset"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &CreateAnnotationDataSetResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a dataset.
//
// @param request - CreateAnnotationDataSetRequest
//
// @return CreateAnnotationDataSetResponse
func (client *Client) CreateAnnotationDataSet(request *CreateAnnotationDataSetRequest) (_result *CreateAnnotationDataSetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateAnnotationDataSetResponse{}
	_body, _err := client.CreateAnnotationDataSetWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a tag table.
//
// @param request - CreateAnnotationLabelRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAnnotationLabelResponse
func (client *Client) CreateAnnotationLabelWithOptions(request *CreateAnnotationLabelRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateAnnotationLabelResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateAnnotationLabel"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/ml/annotationlabel"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &CreateAnnotationLabelResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a tag table.
//
// @param request - CreateAnnotationLabelRequest
//
// @return CreateAnnotationLabelResponse
func (client *Client) CreateAnnotationLabel(request *CreateAnnotationLabelRequest) (_result *CreateAnnotationLabelResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateAnnotationLabelResponse{}
	_body, _err := client.CreateAnnotationLabelWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a Logtail configuration.
//
// Description:
//
// ### [](#)Usage notes
//
// 	- Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
//
// 	- An AccessKey pair is created and obtained. For more information, see [AccessKey pair](https://help.aliyun.com/document_detail/29009.html).
//
// The AccessKey pair of an Alibaba Cloud account has permissions on all API operations. Using these credentials to perform operations in Simple Log Service is a high-risk operation. We recommend that you use a RAM user to call API operations or perform routine O\\&M. To create a RAM user, log on to the RAM console. Make sure that the RAM user has the management permissions on Simple Log Service resources. For more information, see [Create a RAM user and authorize the RAM user to access Log Service](https://help.aliyun.com/document_detail/47664.html).
//
// 	- The information that is required to query logs is obtained. The information includes the name of the project to which the logs belong, the region of the project, and the name of the Logstore to which the logs belong. For more information, see [Manage a project](https://help.aliyun.com/document_detail/48984.html).
//
// 	- You can create up to 100 Logtail configurations in a project.
//
// 	- The Logtail configuration is planned out. For more information, see [Logtail configurations](https://help.aliyun.com/document_detail/29058.html).
//
// @param request - CreateConfigRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateConfigResponse
func (client *Client) CreateConfigWithOptions(project *string, request *CreateConfigRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateConfigResponse, _err error) {
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
		Action:      tea.String("CreateConfig"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/configs"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &CreateConfigResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a Logtail configuration.
//
// Description:
//
// ### [](#)Usage notes
//
// 	- Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
//
// 	- An AccessKey pair is created and obtained. For more information, see [AccessKey pair](https://help.aliyun.com/document_detail/29009.html).
//
// The AccessKey pair of an Alibaba Cloud account has permissions on all API operations. Using these credentials to perform operations in Simple Log Service is a high-risk operation. We recommend that you use a RAM user to call API operations or perform routine O\\&M. To create a RAM user, log on to the RAM console. Make sure that the RAM user has the management permissions on Simple Log Service resources. For more information, see [Create a RAM user and authorize the RAM user to access Log Service](https://help.aliyun.com/document_detail/47664.html).
//
// 	- The information that is required to query logs is obtained. The information includes the name of the project to which the logs belong, the region of the project, and the name of the Logstore to which the logs belong. For more information, see [Manage a project](https://help.aliyun.com/document_detail/48984.html).
//
// 	- You can create up to 100 Logtail configurations in a project.
//
// 	- The Logtail configuration is planned out. For more information, see [Logtail configurations](https://help.aliyun.com/document_detail/29058.html).
//
// @param request - CreateConfigRequest
//
// @return CreateConfigResponse
func (client *Client) CreateConfig(project *string, request *CreateConfigRequest) (_result *CreateConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateConfigResponse{}
	_body, _err := client.CreateConfigWithOptions(project, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a consumer group for a Logstore.
//
// Description:
//
// ### Usage notes
//
// 	- Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
//
// 	- You can create up to 30 consumer groups for a Logstore.
//
// 	- Simple Log Service provides examples of both regular log consumption and consumer group-based log consumption by using Simple Log Service SDKs for Java. For more information, see [Consume log data](https://help.aliyun.com/document_detail/120035.html) and [Use consumer groups to consume data](https://help.aliyun.com/document_detail/28998.html).
//
// @param request - CreateConsumerGroupRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateConsumerGroupResponse
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

// Summary:
//
// Creates a consumer group for a Logstore.
//
// Description:
//
// ### Usage notes
//
// 	- Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
//
// 	- You can create up to 30 consumer groups for a Logstore.
//
// 	- Simple Log Service provides examples of both regular log consumption and consumer group-based log consumption by using Simple Log Service SDKs for Java. For more information, see [Consume log data](https://help.aliyun.com/document_detail/120035.html) and [Use consumer groups to consume data](https://help.aliyun.com/document_detail/28998.html).
//
// @param request - CreateConsumerGroupRequest
//
// @return CreateConsumerGroupResponse
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

// Summary:
//
// Creates a dashboard.
//
// Description:
//
// ### [](#)Usage notes
//
// 	- Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
//
// 	- An AccessKey pair is created and obtained. For more information, see [AccessKey pair](https://help.aliyun.com/document_detail/29009.html).
//
// The AccessKey pair of an Alibaba Cloud account has permissions on all API operations. Using these credentials to perform operations in Simple Log Service is a high-risk operation. We recommend that you use a RAM user to call API operations or perform routine O\\&M. To create a RAM user, log on to the RAM console. Make sure that the RAM user has the management permissions on Simple Log Service resources. For more information, see [Create a RAM user and authorize the RAM user to access Simple Log Service](https://help.aliyun.com/document_detail/47664.html).
//
// 	- The information that is required to query logs is obtained. The information includes the name of the project to which the logs belong and the region of the project. For more information, see [Manage a project](https://help.aliyun.com/document_detail/48984.html).
//
// @param request - CreateDashboardRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDashboardResponse
func (client *Client) CreateDashboardWithOptions(project *string, request *CreateDashboardRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateDashboardResponse, _err error) {
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
		Action:      tea.String("CreateDashboard"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/dashboards"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &CreateDashboardResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a dashboard.
//
// Description:
//
// ### [](#)Usage notes
//
// 	- Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
//
// 	- An AccessKey pair is created and obtained. For more information, see [AccessKey pair](https://help.aliyun.com/document_detail/29009.html).
//
// The AccessKey pair of an Alibaba Cloud account has permissions on all API operations. Using these credentials to perform operations in Simple Log Service is a high-risk operation. We recommend that you use a RAM user to call API operations or perform routine O\\&M. To create a RAM user, log on to the RAM console. Make sure that the RAM user has the management permissions on Simple Log Service resources. For more information, see [Create a RAM user and authorize the RAM user to access Simple Log Service](https://help.aliyun.com/document_detail/47664.html).
//
// 	- The information that is required to query logs is obtained. The information includes the name of the project to which the logs belong and the region of the project. For more information, see [Manage a project](https://help.aliyun.com/document_detail/48984.html).
//
// @param request - CreateDashboardRequest
//
// @return CreateDashboardResponse
func (client *Client) CreateDashboard(project *string, request *CreateDashboardRequest) (_result *CreateDashboardResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateDashboardResponse{}
	_body, _err := client.CreateDashboardWithOptions(project, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Binds a new custom domain name to a project.
//
// Description:
//
// Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
//
// @param request - CreateDomainRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDomainResponse
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

// Summary:
//
// Binds a new custom domain name to a project.
//
// Description:
//
// Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
//
// @param request - CreateDomainRequest
//
// @return CreateDomainResponse
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

// Summary:
//
// 创建数据加工任务
//
// @param request - CreateETLRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateETLResponse
func (client *Client) CreateETLWithOptions(project *string, request *CreateETLRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateETLResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Configuration)) {
		body["configuration"] = request.Configuration
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.DisplayName)) {
		body["displayName"] = request.DisplayName
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateETL"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/etls"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &CreateETLResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建数据加工任务
//
// @param request - CreateETLRequest
//
// @return CreateETLResponse
func (client *Client) CreateETL(project *string, request *CreateETLRequest) (_result *CreateETLResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateETLResponse{}
	_body, _err := client.CreateETLWithOptions(project, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates indexes for a Logstore.
//
// Description:
//
// Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
//
// @param request - CreateIndexRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateIndexResponse
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

// Summary:
//
// Creates indexes for a Logstore.
//
// Description:
//
// Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
//
// @param request - CreateIndexRequest
//
// @return CreateIndexResponse
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

// Summary:
//
// Create a Logstore
//
// Description:
//
// ### Usage notes
//
// Host consists of a project name and a Log Service endpoint. You must specify a project in Host.
//
// @param request - CreateLogStoreRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateLogStoreResponse
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

	if !tea.BoolValue(util.IsUnset(request.InfrequentAccessTTL)) {
		body["infrequentAccessTTL"] = request.InfrequentAccessTTL
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

// Summary:
//
// Create a Logstore
//
// Description:
//
// ### Usage notes
//
// Host consists of a project name and a Log Service endpoint. You must specify a project in Host.
//
// @param request - CreateLogStoreRequest
//
// @return CreateLogStoreResponse
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

// Summary:
//
// Enables the service log feature for a project.
//
// Description:
//
// Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
//
// @param request - CreateLoggingRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateLoggingResponse
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

// Summary:
//
// Enables the service log feature for a project.
//
// Description:
//
// Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
//
// @param request - CreateLoggingRequest
//
// @return CreateLoggingResponse
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

// Summary:
//
// Creates a Logtail pipeline configuration.
//
// Description:
//
// The UK (London) region is supported. Supported regions are constantly updated.
//
// @param request - CreateLogtailPipelineConfigRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateLogtailPipelineConfigResponse
func (client *Client) CreateLogtailPipelineConfigWithOptions(project *string, request *CreateLogtailPipelineConfigRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateLogtailPipelineConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Aggregators)) {
		body["aggregators"] = request.Aggregators
	}

	if !tea.BoolValue(util.IsUnset(request.ConfigName)) {
		body["configName"] = request.ConfigName
	}

	if !tea.BoolValue(util.IsUnset(request.Flushers)) {
		body["flushers"] = request.Flushers
	}

	if !tea.BoolValue(util.IsUnset(request.Global)) {
		body["global"] = request.Global
	}

	if !tea.BoolValue(util.IsUnset(request.Inputs)) {
		body["inputs"] = request.Inputs
	}

	if !tea.BoolValue(util.IsUnset(request.LogSample)) {
		body["logSample"] = request.LogSample
	}

	if !tea.BoolValue(util.IsUnset(request.Processors)) {
		body["processors"] = request.Processors
	}

	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateLogtailPipelineConfig"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/pipelineconfigs"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &CreateLogtailPipelineConfigResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a Logtail pipeline configuration.
//
// Description:
//
// The UK (London) region is supported. Supported regions are constantly updated.
//
// @param request - CreateLogtailPipelineConfigRequest
//
// @return CreateLogtailPipelineConfigResponse
func (client *Client) CreateLogtailPipelineConfig(project *string, request *CreateLogtailPipelineConfigRequest) (_result *CreateLogtailPipelineConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateLogtailPipelineConfigResponse{}
	_body, _err := client.CreateLogtailPipelineConfigWithOptions(project, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a machine group.
//
// Description:
//
// ### Usage notes
//
// Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
//
// @param request - CreateMachineGroupRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateMachineGroupResponse
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

// Summary:
//
// Creates a machine group.
//
// Description:
//
// ### Usage notes
//
// Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
//
// @param request - CreateMachineGroupRequest
//
// @return CreateMachineGroupResponse
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

// Summary:
//
// 创建OSS投递任务
//
// @param request - CreateOSSExportRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateOSSExportResponse
func (client *Client) CreateOSSExportWithOptions(project *string, request *CreateOSSExportRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateOSSExportResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Configuration)) {
		body["configuration"] = request.Configuration
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.DisplayName)) {
		body["displayName"] = request.DisplayName
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateOSSExport"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/ossexports"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &CreateOSSExportResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建OSS投递任务
//
// @param request - CreateOSSExportRequest
//
// @return CreateOSSExportResponse
func (client *Client) CreateOSSExport(project *string, request *CreateOSSExportRequest) (_result *CreateOSSExportResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateOSSExportResponse{}
	_body, _err := client.CreateOSSExportWithOptions(project, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建OSSHDFS投递任务
//
// @param request - CreateOSSHDFSExportRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateOSSHDFSExportResponse
func (client *Client) CreateOSSHDFSExportWithOptions(project *string, request *CreateOSSHDFSExportRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateOSSHDFSExportResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Configuration)) {
		body["configuration"] = request.Configuration
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.DisplayName)) {
		body["displayName"] = request.DisplayName
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateOSSHDFSExport"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/osshdfsexports"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &CreateOSSHDFSExportResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建OSSHDFS投递任务
//
// @param request - CreateOSSHDFSExportRequest
//
// @return CreateOSSHDFSExportResponse
func (client *Client) CreateOSSHDFSExport(project *string, request *CreateOSSHDFSExportRequest) (_result *CreateOSSHDFSExportResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateOSSHDFSExportResponse{}
	_body, _err := client.CreateOSSHDFSExportWithOptions(project, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建OSS导入任务
//
// @param request - CreateOSSIngestionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateOSSIngestionResponse
func (client *Client) CreateOSSIngestionWithOptions(project *string, request *CreateOSSIngestionRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateOSSIngestionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Configuration)) {
		body["configuration"] = request.Configuration
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.DisplayName)) {
		body["displayName"] = request.DisplayName
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Schedule)) {
		body["schedule"] = request.Schedule
	}

	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateOSSIngestion"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/ossingestions"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &CreateOSSIngestionResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建OSS导入任务
//
// @param request - CreateOSSIngestionRequest
//
// @return CreateOSSIngestionResponse
func (client *Client) CreateOSSIngestion(project *string, request *CreateOSSIngestionRequest) (_result *CreateOSSIngestionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateOSSIngestionResponse{}
	_body, _err := client.CreateOSSIngestionWithOptions(project, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an Object Storage Service (OSS) external store.
//
// Description:
//
// ### [](#)Usage notes
//
// Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
//
// @param request - CreateOssExternalStoreRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateOssExternalStoreResponse
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

// Summary:
//
// Creates an Object Storage Service (OSS) external store.
//
// Description:
//
// ### [](#)Usage notes
//
// Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
//
// @param request - CreateOssExternalStoreRequest
//
// @return CreateOssExternalStoreResponse
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

// Summary:
//
// Creates a project.
//
// @param request - CreateProjectRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateProjectResponse
func (client *Client) CreateProjectWithOptions(request *CreateProjectRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateProjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DataRedundancyType)) {
		body["dataRedundancyType"] = request.DataRedundancyType
	}

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

// Summary:
//
// Creates a project.
//
// @param request - CreateProjectRequest
//
// @return CreateProjectResponse
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

// Summary:
//
// Creates an ApsaraDB RDS external store.
//
// Description:
//
// Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
//
// @param request - CreateRdsExternalStoreRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateRdsExternalStoreResponse
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

// Summary:
//
// Creates an ApsaraDB RDS external store.
//
// Description:
//
// Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
//
// @param request - CreateRdsExternalStoreRequest
//
// @return CreateRdsExternalStoreResponse
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

// Summary:
//
// Creates a saved search.
//
// Description:
//
// ### Usage notes
//
// Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
//
// @param request - CreateSavedSearchRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSavedSearchResponse
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

// Summary:
//
// Creates a saved search.
//
// Description:
//
// ### Usage notes
//
// Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
//
// @param request - CreateSavedSearchRequest
//
// @return CreateSavedSearchResponse
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

// Summary:
//
// 创建定时SQL任务
//
// @param request - CreateScheduledSQLRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateScheduledSQLResponse
func (client *Client) CreateScheduledSQLWithOptions(project *string, request *CreateScheduledSQLRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateScheduledSQLResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Configuration)) {
		body["configuration"] = request.Configuration
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.DisplayName)) {
		body["displayName"] = request.DisplayName
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Schedule)) {
		body["schedule"] = request.Schedule
	}

	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateScheduledSQL"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/scheduledsqls"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &CreateScheduledSQLResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建定时SQL任务
//
// @param request - CreateScheduledSQLRequest
//
// @return CreateScheduledSQLResponse
func (client *Client) CreateScheduledSQL(project *string, request *CreateScheduledSQLRequest) (_result *CreateScheduledSQLResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateScheduledSQLResponse{}
	_body, _err := client.CreateScheduledSQLWithOptions(project, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建独享sql实例
//
// @param request - CreateSqlInstanceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSqlInstanceResponse
func (client *Client) CreateSqlInstanceWithOptions(project *string, request *CreateSqlInstanceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateSqlInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Cu)) {
		body["cu"] = request.Cu
	}

	if !tea.BoolValue(util.IsUnset(request.UseAsDefault)) {
		body["useAsDefault"] = request.UseAsDefault
	}

	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateSqlInstance"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/sqlinstance"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &CreateSqlInstanceResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建独享sql实例
//
// @param request - CreateSqlInstanceRequest
//
// @return CreateSqlInstanceResponse
func (client *Client) CreateSqlInstance(project *string, request *CreateSqlInstanceRequest) (_result *CreateSqlInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateSqlInstanceResponse{}
	_body, _err := client.CreateSqlInstanceWithOptions(project, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建StoreView
//
// @param request - CreateStoreViewRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateStoreViewResponse
func (client *Client) CreateStoreViewWithOptions(project *string, request *CreateStoreViewRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateStoreViewResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.StoreType)) {
		body["storeType"] = request.StoreType
	}

	if !tea.BoolValue(util.IsUnset(request.Stores)) {
		body["stores"] = request.Stores
	}

	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateStoreView"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/storeviews"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &CreateStoreViewResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建StoreView
//
// @param request - CreateStoreViewRequest
//
// @return CreateStoreViewResponse
func (client *Client) CreateStoreView(project *string, request *CreateStoreViewRequest) (_result *CreateStoreViewResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateStoreViewResponse{}
	_body, _err := client.CreateStoreViewWithOptions(project, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建Ticket
//
// @param request - CreateTicketRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTicketResponse
func (client *Client) CreateTicketWithOptions(request *CreateTicketRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateTicketResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessTokenExpirationTime)) {
		query["accessTokenExpirationTime"] = request.AccessTokenExpirationTime
	}

	if !tea.BoolValue(util.IsUnset(request.ExpirationTime)) {
		query["expirationTime"] = request.ExpirationTime
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateTicket"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/tickets"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateTicketResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建Ticket
//
// @param request - CreateTicketRequest
//
// @return CreateTicketResponse
func (client *Client) CreateTicket(request *CreateTicketRequest) (_result *CreateTicketResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateTicketResponse{}
	_body, _err := client.CreateTicketWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除告警
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAlertResponse
func (client *Client) DeleteAlertWithOptions(project *string, alertName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteAlertResponse, _err error) {
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
		Pathname:    tea.String("/alerts/" + tea.StringValue(alertName)),
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

// Summary:
//
// 删除告警
//
// @return DeleteAlertResponse
func (client *Client) DeleteAlert(project *string, alertName *string) (_result *DeleteAlertResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteAlertResponse{}
	_body, _err := client.DeleteAlertWithOptions(project, alertName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Removes data from a dataset.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAnnotationDataResponse
func (client *Client) DeleteAnnotationDataWithOptions(datasetId *string, annotationdataId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteAnnotationDataResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteAnnotationData"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/ml/annotationdataset/" + tea.StringValue(datasetId) + "/annotationdata/" + tea.StringValue(annotationdataId)),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &DeleteAnnotationDataResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes data from a dataset.
//
// @return DeleteAnnotationDataResponse
func (client *Client) DeleteAnnotationData(datasetId *string, annotationdataId *string) (_result *DeleteAnnotationDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteAnnotationDataResponse{}
	_body, _err := client.DeleteAnnotationDataWithOptions(datasetId, annotationdataId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a dataset.
//
// Description:
//
// You can delete a dataset only if no data exists in the dataset.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAnnotationDataSetResponse
func (client *Client) DeleteAnnotationDataSetWithOptions(datasetId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteAnnotationDataSetResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteAnnotationDataSet"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/ml/annotationdataset/" + tea.StringValue(datasetId)),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &DeleteAnnotationDataSetResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a dataset.
//
// Description:
//
// You can delete a dataset only if no data exists in the dataset.
//
// @return DeleteAnnotationDataSetResponse
func (client *Client) DeleteAnnotationDataSet(datasetId *string) (_result *DeleteAnnotationDataSetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteAnnotationDataSetResponse{}
	_body, _err := client.DeleteAnnotationDataSetWithOptions(datasetId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a tag table.
//
// Description:
//
// Only non-built-in tags can be deleted.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAnnotationLabelResponse
func (client *Client) DeleteAnnotationLabelWithOptions(labelId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteAnnotationLabelResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteAnnotationLabel"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/ml/annotationlabel/" + tea.StringValue(labelId)),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &DeleteAnnotationLabelResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a tag table.
//
// Description:
//
// Only non-built-in tags can be deleted.
//
// @return DeleteAnnotationLabelResponse
func (client *Client) DeleteAnnotationLabel(labelId *string) (_result *DeleteAnnotationLabelResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteAnnotationLabelResponse{}
	_body, _err := client.DeleteAnnotationLabelWithOptions(labelId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 通过调用DeleteCollectionPolicy删除配置的日志采集规则
//
// @param request - DeleteCollectionPolicyRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCollectionPolicyResponse
func (client *Client) DeleteCollectionPolicyWithOptions(policyName *string, request *DeleteCollectionPolicyRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteCollectionPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DataCode)) {
		query["dataCode"] = request.DataCode
	}

	if !tea.BoolValue(util.IsUnset(request.ProductCode)) {
		query["productCode"] = request.ProductCode
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteCollectionPolicy"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/collectionpolicy/" + tea.StringValue(policyName)),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteCollectionPolicyResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 通过调用DeleteCollectionPolicy删除配置的日志采集规则
//
// @param request - DeleteCollectionPolicyRequest
//
// @return DeleteCollectionPolicyResponse
func (client *Client) DeleteCollectionPolicy(policyName *string, request *DeleteCollectionPolicyRequest) (_result *DeleteCollectionPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteCollectionPolicyResponse{}
	_body, _err := client.DeleteCollectionPolicyWithOptions(policyName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a Logtail configuration.
//
// Description:
//
// ### [](#)Usage notes
//
// 	- Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
//
// 	- If a Logtail configuration is applied to a machine group, you cannot collect data from the machine group after you delete the Logtail configuration.
//
// 	- An AccessKey pair is created and obtained. For more information, see [AccessKey pair](https://help.aliyun.com/document_detail/29009.html).
//
// The AccessKey pair of an Alibaba Cloud account has permissions on all API operations. Using these credentials to perform operations in Simple Log Service is a high-risk operation. We recommend that you use a RAM user to call API operations or perform routine O\\&M. To create a RAM user, log on to the RAM console. Make sure that the RAM user has the management permissions on Simple Log Service resources. For more information, see [Create a RAM user and authorize the RAM user to access Simple Log Service](https://help.aliyun.com/document_detail/47664.html).
//
// 	- The information that is required to query logs is obtained. The information includes the name of the project to which the logs belong and the region of the project. For more information, see [Manage a project](https://help.aliyun.com/document_detail/48984.html).
//
// 	- The name of the required Logtail configuration is obtained. For more information, see [ListConfig](https://help.aliyun.com/document_detail/29043.html).
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteConfigResponse
func (client *Client) DeleteConfigWithOptions(project *string, configName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteConfigResponse, _err error) {
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteConfig"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/configs/" + tea.StringValue(configName)),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &DeleteConfigResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a Logtail configuration.
//
// Description:
//
// ### [](#)Usage notes
//
// 	- Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
//
// 	- If a Logtail configuration is applied to a machine group, you cannot collect data from the machine group after you delete the Logtail configuration.
//
// 	- An AccessKey pair is created and obtained. For more information, see [AccessKey pair](https://help.aliyun.com/document_detail/29009.html).
//
// The AccessKey pair of an Alibaba Cloud account has permissions on all API operations. Using these credentials to perform operations in Simple Log Service is a high-risk operation. We recommend that you use a RAM user to call API operations or perform routine O\\&M. To create a RAM user, log on to the RAM console. Make sure that the RAM user has the management permissions on Simple Log Service resources. For more information, see [Create a RAM user and authorize the RAM user to access Simple Log Service](https://help.aliyun.com/document_detail/47664.html).
//
// 	- The information that is required to query logs is obtained. The information includes the name of the project to which the logs belong and the region of the project. For more information, see [Manage a project](https://help.aliyun.com/document_detail/48984.html).
//
// 	- The name of the required Logtail configuration is obtained. For more information, see [ListConfig](https://help.aliyun.com/document_detail/29043.html).
//
// @return DeleteConfigResponse
func (client *Client) DeleteConfig(project *string, configName *string) (_result *DeleteConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteConfigResponse{}
	_body, _err := client.DeleteConfigWithOptions(project, configName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a consumer group.
//
// Description:
//
// Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteConsumerGroupResponse
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

// Summary:
//
// Deletes a consumer group.
//
// Description:
//
// Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
//
// @return DeleteConsumerGroupResponse
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

// Summary:
//
// Deletes a dashboard.
//
// Description:
//
// ### [](#)Usage notes
//
// 	- Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
//
// 	- An AccessKey pair is created and obtained. For more information, see [AccessKey pair](https://help.aliyun.com/document_detail/29009.html).
//
// The AccessKey pair of an Alibaba Cloud account has permissions on all API operations. Using these credentials to perform operations in Simple Log Service is a high-risk operation. We recommend that you use a RAM user to call API operations or perform routine O\\&M. To create a RAM user, log on to the RAM console. Make sure that the RAM user has the management permissions on Simple Log Service resources. For more information, see [Create a RAM user and authorize the RAM user to access Simple Log Service](https://help.aliyun.com/document_detail/47664.html).
//
// 	- The information that is required to query logs is obtained. The information includes the name of the project to which the logs belong and the region of the project. For more information, see [Manage a project](https://help.aliyun.com/document_detail/48984.html).
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDashboardResponse
func (client *Client) DeleteDashboardWithOptions(project *string, dashboardName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteDashboardResponse, _err error) {
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteDashboard"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/dashboards/" + tea.StringValue(dashboardName)),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &DeleteDashboardResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a dashboard.
//
// Description:
//
// ### [](#)Usage notes
//
// 	- Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
//
// 	- An AccessKey pair is created and obtained. For more information, see [AccessKey pair](https://help.aliyun.com/document_detail/29009.html).
//
// The AccessKey pair of an Alibaba Cloud account has permissions on all API operations. Using these credentials to perform operations in Simple Log Service is a high-risk operation. We recommend that you use a RAM user to call API operations or perform routine O\\&M. To create a RAM user, log on to the RAM console. Make sure that the RAM user has the management permissions on Simple Log Service resources. For more information, see [Create a RAM user and authorize the RAM user to access Simple Log Service](https://help.aliyun.com/document_detail/47664.html).
//
// 	- The information that is required to query logs is obtained. The information includes the name of the project to which the logs belong and the region of the project. For more information, see [Manage a project](https://help.aliyun.com/document_detail/48984.html).
//
// @return DeleteDashboardResponse
func (client *Client) DeleteDashboard(project *string, dashboardName *string) (_result *DeleteDashboardResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteDashboardResponse{}
	_body, _err := client.DeleteDashboardWithOptions(project, dashboardName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a custom domain name that is bound to a project.
//
// Description:
//
// Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDomainResponse
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

// Summary:
//
// Deletes a custom domain name that is bound to a project.
//
// Description:
//
// Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
//
// @return DeleteDomainResponse
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

// Summary:
//
// 删除数据加工任务
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteETLResponse
func (client *Client) DeleteETLWithOptions(project *string, etlName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteETLResponse, _err error) {
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteETL"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/etls/" + tea.StringValue(etlName)),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &DeleteETLResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除数据加工任务
//
// @return DeleteETLResponse
func (client *Client) DeleteETL(project *string, etlName *string) (_result *DeleteETLResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteETLResponse{}
	_body, _err := client.DeleteETLWithOptions(project, etlName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes an external store.
//
// Description:
//
// Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteExternalStoreResponse
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

// Summary:
//
// Deletes an external store.
//
// Description:
//
// Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
//
// @return DeleteExternalStoreResponse
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

// Summary:
//
// Deletes an index of a Logstore.
//
// Description:
//
// ### Usage notes
//
// Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteIndexResponse
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

// Summary:
//
// Deletes an index of a Logstore.
//
// Description:
//
// ### Usage notes
//
// Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
//
// @return DeleteIndexResponse
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

// Summary:
//
// Deletes a Logstore, including all shards and indexes in the Logstore.
//
// Description:
//
// ### Usage notes
//
// Host consists of a project name and a Log Service endpoint. You must specify a project in Host.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteLogStoreResponse
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

// Summary:
//
// Deletes a Logstore, including all shards and indexes in the Logstore.
//
// Description:
//
// ### Usage notes
//
// Host consists of a project name and a Log Service endpoint. You must specify a project in Host.
//
// @return DeleteLogStoreResponse
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

// Summary:
//
// 关闭项目的服务日志记录。
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteLoggingResponse
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

// Summary:
//
// 关闭项目的服务日志记录。
//
// @return DeleteLoggingResponse
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

// Summary:
//
// Deletes a Logtail pipeline configuration.
//
// Description:
//
// The UK (London) region is supported. Supported regions are constantly updated.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteLogtailPipelineConfigResponse
func (client *Client) DeleteLogtailPipelineConfigWithOptions(project *string, configName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteLogtailPipelineConfigResponse, _err error) {
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteLogtailPipelineConfig"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/pipelineconfigs/" + tea.StringValue(configName)),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &DeleteLogtailPipelineConfigResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a Logtail pipeline configuration.
//
// Description:
//
// The UK (London) region is supported. Supported regions are constantly updated.
//
// @return DeleteLogtailPipelineConfigResponse
func (client *Client) DeleteLogtailPipelineConfig(project *string, configName *string) (_result *DeleteLogtailPipelineConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteLogtailPipelineConfigResponse{}
	_body, _err := client.DeleteLogtailPipelineConfigWithOptions(project, configName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a machine group. If the Logtail configurations for log collection are applied to a machine group, the configurations are disassociated from the machine group after the machine group is deleted.
//
// Description:
//
// ### Usage notes
//
// Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteMachineGroupResponse
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

// Summary:
//
// Deletes a machine group. If the Logtail configurations for log collection are applied to a machine group, the configurations are disassociated from the machine group after the machine group is deleted.
//
// Description:
//
// ### Usage notes
//
// Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
//
// @return DeleteMachineGroupResponse
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

// Summary:
//
// 删除OSS投递任务
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteOSSExportResponse
func (client *Client) DeleteOSSExportWithOptions(project *string, ossExportName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteOSSExportResponse, _err error) {
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteOSSExport"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/ossexports/" + tea.StringValue(ossExportName)),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &DeleteOSSExportResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除OSS投递任务
//
// @return DeleteOSSExportResponse
func (client *Client) DeleteOSSExport(project *string, ossExportName *string) (_result *DeleteOSSExportResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteOSSExportResponse{}
	_body, _err := client.DeleteOSSExportWithOptions(project, ossExportName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除OSSHDFS投递任务
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteOSSHDFSExportResponse
func (client *Client) DeleteOSSHDFSExportWithOptions(project *string, ossExportName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteOSSHDFSExportResponse, _err error) {
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteOSSHDFSExport"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/osshdfsexports/" + tea.StringValue(ossExportName)),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &DeleteOSSHDFSExportResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除OSSHDFS投递任务
//
// @return DeleteOSSHDFSExportResponse
func (client *Client) DeleteOSSHDFSExport(project *string, ossExportName *string) (_result *DeleteOSSHDFSExportResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteOSSHDFSExportResponse{}
	_body, _err := client.DeleteOSSHDFSExportWithOptions(project, ossExportName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除OSS导入任务
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteOSSIngestionResponse
func (client *Client) DeleteOSSIngestionWithOptions(project *string, ossIngestionName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteOSSIngestionResponse, _err error) {
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteOSSIngestion"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/ossingestions/" + tea.StringValue(ossIngestionName)),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &DeleteOSSIngestionResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除OSS导入任务
//
// @return DeleteOSSIngestionResponse
func (client *Client) DeleteOSSIngestion(project *string, ossIngestionName *string) (_result *DeleteOSSIngestionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteOSSIngestionResponse{}
	_body, _err := client.DeleteOSSIngestionWithOptions(project, ossIngestionName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除project
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteProjectResponse
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

// Summary:
//
// 删除project
//
// @return DeleteProjectResponse
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

// Summary:
//
// Deletes a project policy.
//
// Description:
//
// ### Usage notes
//
// Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteProjectPolicyResponse
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
		BodyType:    tea.String("none"),
	}
	_result = &DeleteProjectPolicyResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a project policy.
//
// Description:
//
// ### Usage notes
//
// Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
//
// @return DeleteProjectPolicyResponse
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

// Summary:
//
// Deletes a saved search.
//
// Description:
//
// ### Usage notes
//
// Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSavedSearchResponse
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

// Summary:
//
// Deletes a saved search.
//
// Description:
//
// ### Usage notes
//
// Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
//
// @return DeleteSavedSearchResponse
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

// Summary:
//
// 删除定时SQL任务
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteScheduledSQLResponse
func (client *Client) DeleteScheduledSQLWithOptions(project *string, scheduledSQLName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteScheduledSQLResponse, _err error) {
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteScheduledSQL"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/scheduledsqls/" + tea.StringValue(scheduledSQLName)),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &DeleteScheduledSQLResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除定时SQL任务
//
// @return DeleteScheduledSQLResponse
func (client *Client) DeleteScheduledSQL(project *string, scheduledSQLName *string) (_result *DeleteScheduledSQLResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteScheduledSQLResponse{}
	_body, _err := client.DeleteScheduledSQLWithOptions(project, scheduledSQLName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI DeleteShipper is deprecated
//
// Summary:
//
// Deletes the log shipping job of a Logstore.
//
// Description:
//
// Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteShipperResponse
// Deprecated
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

// Deprecated: OpenAPI DeleteShipper is deprecated
//
// Summary:
//
// Deletes the log shipping job of a Logstore.
//
// Description:
//
// Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
//
// @return DeleteShipperResponse
// Deprecated
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

// Summary:
//
// 删除StoreView
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteStoreViewResponse
func (client *Client) DeleteStoreViewWithOptions(project *string, name *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteStoreViewResponse, _err error) {
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteStoreView"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/storeviews/" + tea.StringValue(name)),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &DeleteStoreViewResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除StoreView
//
// @return DeleteStoreViewResponse
func (client *Client) DeleteStoreView(project *string, name *string) (_result *DeleteStoreViewResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteStoreViewResponse{}
	_body, _err := client.DeleteStoreViewWithOptions(project, name, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 禁用告警
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableAlertResponse
func (client *Client) DisableAlertWithOptions(project *string, alertName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DisableAlertResponse, _err error) {
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DisableAlert"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/alerts/" + tea.StringValue(alertName) + "?action=disable"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &DisableAlertResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 禁用告警
//
// @return DisableAlertResponse
func (client *Client) DisableAlert(project *string, alertName *string) (_result *DisableAlertResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DisableAlertResponse{}
	_body, _err := client.DisableAlertWithOptions(project, alertName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 禁用定时SQL任务
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableScheduledSQLResponse
func (client *Client) DisableScheduledSQLWithOptions(project *string, scheduledSQLName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DisableScheduledSQLResponse, _err error) {
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DisableScheduledSQL"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/scheduledsqls/" + tea.StringValue(scheduledSQLName) + "?action=disable"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &DisableScheduledSQLResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 禁用定时SQL任务
//
// @return DisableScheduledSQLResponse
func (client *Client) DisableScheduledSQL(project *string, scheduledSQLName *string) (_result *DisableScheduledSQLResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DisableScheduledSQLResponse{}
	_body, _err := client.DisableScheduledSQLWithOptions(project, scheduledSQLName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 启用告警
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableAlertResponse
func (client *Client) EnableAlertWithOptions(project *string, alertName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *EnableAlertResponse, _err error) {
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("EnableAlert"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/alerts/" + tea.StringValue(alertName) + "?action=enable"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &EnableAlertResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 启用告警
//
// @return EnableAlertResponse
func (client *Client) EnableAlert(project *string, alertName *string) (_result *EnableAlertResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &EnableAlertResponse{}
	_body, _err := client.EnableAlertWithOptions(project, alertName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 启用定时SQL任务
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableScheduledSQLResponse
func (client *Client) EnableScheduledSQLWithOptions(project *string, scheduledSQLName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *EnableScheduledSQLResponse, _err error) {
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("EnableScheduledSQL"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/scheduledsqls/" + tea.StringValue(scheduledSQLName) + "?action=enable"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &EnableScheduledSQLResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 启用定时SQL任务
//
// @return EnableScheduledSQLResponse
func (client *Client) EnableScheduledSQL(project *string, scheduledSQLName *string) (_result *EnableScheduledSQLResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &EnableScheduledSQLResponse{}
	_body, _err := client.EnableScheduledSQLWithOptions(project, scheduledSQLName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// GetAlert
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAlertResponse
func (client *Client) GetAlertWithOptions(project *string, alertName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetAlertResponse, _err error) {
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
		Pathname:    tea.String("/alerts/" + tea.StringValue(alertName)),
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

// Summary:
//
// GetAlert
//
// @return GetAlertResponse
func (client *Client) GetAlert(project *string, alertName *string) (_result *GetAlertResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetAlertResponse{}
	_body, _err := client.GetAlertWithOptions(project, alertName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries data in a dataset based on the unique identifier of the data.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAnnotationDataResponse
func (client *Client) GetAnnotationDataWithOptions(datasetId *string, annotationdataId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetAnnotationDataResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetAnnotationData"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/ml/annotationdataset/" + tea.StringValue(datasetId) + "/annotationdata/" + tea.StringValue(annotationdataId)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAnnotationDataResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries data in a dataset based on the unique identifier of the data.
//
// @return GetAnnotationDataResponse
func (client *Client) GetAnnotationData(datasetId *string, annotationdataId *string) (_result *GetAnnotationDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetAnnotationDataResponse{}
	_body, _err := client.GetAnnotationDataWithOptions(datasetId, annotationdataId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a dataset.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAnnotationDataSetResponse
func (client *Client) GetAnnotationDataSetWithOptions(datasetId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetAnnotationDataSetResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetAnnotationDataSet"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/ml/annotationdataset/" + tea.StringValue(datasetId)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAnnotationDataSetResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a dataset.
//
// @return GetAnnotationDataSetResponse
func (client *Client) GetAnnotationDataSet(datasetId *string) (_result *GetAnnotationDataSetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetAnnotationDataSetResponse{}
	_body, _err := client.GetAnnotationDataSetWithOptions(datasetId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a tag table by using a label id.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAnnotationLabelResponse
func (client *Client) GetAnnotationLabelWithOptions(labelId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetAnnotationLabelResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetAnnotationLabel"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/ml/annotationlabel/" + tea.StringValue(labelId)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAnnotationLabelResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a tag table by using a label id.
//
// @return GetAnnotationLabelResponse
func (client *Client) GetAnnotationLabel(labelId *string) (_result *GetAnnotationLabelResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetAnnotationLabelResponse{}
	_body, _err := client.GetAnnotationLabelWithOptions(labelId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the Logtail configurations that are applied to a machine group.
//
// Description:
//
// Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAppliedConfigsResponse
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

// Summary:
//
// Queries the Logtail configurations that are applied to a machine group.
//
// Description:
//
// Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
//
// @return GetAppliedConfigsResponse
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

// Summary:
//
// Queries the machine groups to which a Logtail configuration is bound.
//
// Description:
//
// Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAppliedMachineGroupsResponse
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

// Summary:
//
// Queries the machine groups to which a Logtail configuration is bound.
//
// Description:
//
// Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
//
// @return GetAppliedMachineGroupsResponse
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

// Summary:
//
// Queries the checkpoints of a shard from which data is consumed by a consumer group.
//
// Description:
//
// Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
//
// @param request - GetCheckPointRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCheckPointResponse
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

// Summary:
//
// Queries the checkpoints of a shard from which data is consumed by a consumer group.
//
// Description:
//
// Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
//
// @param request - GetCheckPointRequest
//
// @return GetCheckPointResponse
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

// Summary:
//
// 调用GetCollectionPolicy获取对应的规则
//
// @param request - GetCollectionPolicyRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCollectionPolicyResponse
func (client *Client) GetCollectionPolicyWithOptions(policyName *string, request *GetCollectionPolicyRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetCollectionPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DataCode)) {
		query["dataCode"] = request.DataCode
	}

	if !tea.BoolValue(util.IsUnset(request.ProductCode)) {
		query["productCode"] = request.ProductCode
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetCollectionPolicy"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/collectionpolicy/" + tea.StringValue(policyName)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetCollectionPolicyResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 调用GetCollectionPolicy获取对应的规则
//
// @param request - GetCollectionPolicyRequest
//
// @return GetCollectionPolicyResponse
func (client *Client) GetCollectionPolicy(policyName *string, request *GetCollectionPolicyRequest) (_result *GetCollectionPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetCollectionPolicyResponse{}
	_body, _err := client.GetCollectionPolicyWithOptions(policyName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of a Logtail configuration.
//
// Description:
//
// ### [](#)Usage notes
//
// 	- Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
//
// 	- An AccessKey pair is created and obtained. For more information, see [AccessKey pair](https://help.aliyun.com/document_detail/29009.html).
//
// The AccessKey pair of an Alibaba Cloud account has permissions on all API operations. Using these credentials to perform operations in Simple Log Service is a high-risk operation. We recommend that you use a RAM user to call API operations or perform routine O\\&M. To create a RAM user, log on to the RAM console. Make sure that the RAM user has the management permissions on Simple Log Service resources. For more information, see [Create a RAM user and authorize the RAM user to access Simple Log Service](https://help.aliyun.com/document_detail/47664.html).
//
// 	- The information that is required to query logs is obtained. The information includes the name of the project to which the logs belong and the region of the project. For more information, see [Manage a project](https://help.aliyun.com/document_detail/48984.html).
//
// 	- The name of the required Logtail configuration is obtained. For more information, see [ListConfig](https://help.aliyun.com/document_detail/29043.html).
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetConfigResponse
func (client *Client) GetConfigWithOptions(project *string, configName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetConfigResponse, _err error) {
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetConfig"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/configs/" + tea.StringValue(configName)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetConfigResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a Logtail configuration.
//
// Description:
//
// ### [](#)Usage notes
//
// 	- Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
//
// 	- An AccessKey pair is created and obtained. For more information, see [AccessKey pair](https://help.aliyun.com/document_detail/29009.html).
//
// The AccessKey pair of an Alibaba Cloud account has permissions on all API operations. Using these credentials to perform operations in Simple Log Service is a high-risk operation. We recommend that you use a RAM user to call API operations or perform routine O\\&M. To create a RAM user, log on to the RAM console. Make sure that the RAM user has the management permissions on Simple Log Service resources. For more information, see [Create a RAM user and authorize the RAM user to access Simple Log Service](https://help.aliyun.com/document_detail/47664.html).
//
// 	- The information that is required to query logs is obtained. The information includes the name of the project to which the logs belong and the region of the project. For more information, see [Manage a project](https://help.aliyun.com/document_detail/48984.html).
//
// 	- The name of the required Logtail configuration is obtained. For more information, see [ListConfig](https://help.aliyun.com/document_detail/29043.html).
//
// @return GetConfigResponse
func (client *Client) GetConfig(project *string, configName *string) (_result *GetConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetConfigResponse{}
	_body, _err := client.GetConfigWithOptions(project, configName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the contextual logs of a specified log.
//
// Description:
//
//   You can specify a log as the start log. The time range of a contextual query is one day before and one day after the generation time of the start log.
//
// 	- Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
//
// @param request - GetContextLogsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetContextLogsResponse
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

// Summary:
//
// Queries the contextual logs of a specified log.
//
// Description:
//
//   You can specify a log as the start log. The time range of a contextual query is one day before and one day after the generation time of the start log.
//
// 	- Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
//
// @param request - GetContextLogsRequest
//
// @return GetContextLogsResponse
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

// Summary:
//
// Queries a cursor based on a point in time.
//
// Description:
//
// ### Usage notes
//
// 	- Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
//
// 	- The following content describes the relationships among a cursor, project, Logstore, and shard:
//
//     	- A project can have multiple Logstores.
//
//     	- A Logstore can have multiple shards.
//
//     	- You can use a cursor to obtain a log in a shard.
//
// @param request - GetCursorRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCursorResponse
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

// Summary:
//
// Queries a cursor based on a point in time.
//
// Description:
//
// ### Usage notes
//
// 	- Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
//
// 	- The following content describes the relationships among a cursor, project, Logstore, and shard:
//
//     	- A project can have multiple Logstores.
//
//     	- A Logstore can have multiple shards.
//
//     	- You can use a cursor to obtain a log in a shard.
//
// @param request - GetCursorRequest
//
// @return GetCursorResponse
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

// Summary:
//
// Queries the server time of a cursor.
//
// @param request - GetCursorTimeRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCursorTimeResponse
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

// Summary:
//
// Queries the server time of a cursor.
//
// @param request - GetCursorTimeRequest
//
// @return GetCursorTimeResponse
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

// Summary:
//
// Queries a dashboard.
//
// Description:
//
// ### [](#)Usage notes
//
// 	- Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
//
// 	- An AccessKey pair is created and obtained. For more information, see [AccessKey pair](https://help.aliyun.com/document_detail/29009.html).
//
// The AccessKey pair of an Alibaba Cloud account has permissions on all API operations. Using these credentials to perform operations in Simple Log Service is a high-risk operation. We recommend that you use a RAM user to call API operations or perform routine O\\&M. To create a RAM user, log on to the RAM console. Make sure that the RAM user has the management permissions on Simple Log Service resources. For more information, see [Create a RAM user and authorize the RAM user to access Simple Log Service](https://help.aliyun.com/document_detail/47664.html).
//
// 	- The information that is required to query logs is obtained. The information includes the name of the project to which the logs belong and the region of the project. For more information, see [Manage a project](https://help.aliyun.com/document_detail/48984.html).
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDashboardResponse
func (client *Client) GetDashboardWithOptions(project *string, dashboardName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetDashboardResponse, _err error) {
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetDashboard"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/dashboards/" + tea.StringValue(dashboardName)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetDashboardResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a dashboard.
//
// Description:
//
// ### [](#)Usage notes
//
// 	- Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
//
// 	- An AccessKey pair is created and obtained. For more information, see [AccessKey pair](https://help.aliyun.com/document_detail/29009.html).
//
// The AccessKey pair of an Alibaba Cloud account has permissions on all API operations. Using these credentials to perform operations in Simple Log Service is a high-risk operation. We recommend that you use a RAM user to call API operations or perform routine O\\&M. To create a RAM user, log on to the RAM console. Make sure that the RAM user has the management permissions on Simple Log Service resources. For more information, see [Create a RAM user and authorize the RAM user to access Simple Log Service](https://help.aliyun.com/document_detail/47664.html).
//
// 	- The information that is required to query logs is obtained. The information includes the name of the project to which the logs belong and the region of the project. For more information, see [Manage a project](https://help.aliyun.com/document_detail/48984.html).
//
// @return GetDashboardResponse
func (client *Client) GetDashboard(project *string, dashboardName *string) (_result *GetDashboardResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetDashboardResponse{}
	_body, _err := client.GetDashboardWithOptions(project, dashboardName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取数据加工任务信息
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetETLResponse
func (client *Client) GetETLWithOptions(project *string, etlName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetETLResponse, _err error) {
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetETL"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/etls/" + tea.StringValue(etlName)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetETLResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取数据加工任务信息
//
// @return GetETLResponse
func (client *Client) GetETL(project *string, etlName *string) (_result *GetETLResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetETLResponse{}
	_body, _err := client.GetETLWithOptions(project, etlName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of an external store.
//
// Description:
//
//   The supported data sources of external stores include Object Storage Service (OSS) buckets and ApsaraDB RDS for MySQL databases in a virtual private cloud (VPC).
//
// 	- Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetExternalStoreResponse
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

// Summary:
//
// Queries the details of an external store.
//
// Description:
//
//   The supported data sources of external stores include Object Storage Service (OSS) buckets and ApsaraDB RDS for MySQL databases in a virtual private cloud (VPC).
//
// 	- Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
//
// @return GetExternalStoreResponse
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

// Summary:
//
// Queries the distribution of logs that meet the specified search conditions in a Logstore.
//
// Description:
//
// ### Usage notes
//
// 	- Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
//
// 	- The time range is evenly divided into subintervals in the responses. If the time range that is specified in the request remains unchanged, the subintervals in the responses also remain unchanged.
//
// 	- If the number of logs in a Logstore significantly changes, Simple Log Service cannot predict the number of times that you must call this operation to obtain the complete result. In this case, you must check the value of the progress parameter in the response of each request and determine whether to call this operation one more time to obtain the complete result. Each time you call this operation, the same number of charge units (CUs) are consumed.
//
// 	- After a log is written to a Logstore, you can call the GetHistograms or GetLogs operation to query the log after a short latency. The latency of a query varies based on the type of the log. Simple Log Service classifies logs into the following types based on the log time:
//
//     	- Real-time data: The difference between the time record in a log and the current time on Simple Log Service is within the interval (-180 seconds,900 seconds]. For example, if a log was generated at 12:03:00, September 25, 2014 (UTC) and Simple Log Service received the log at 12:05:00, September 25, 2014 (UTC), Simple Log Service processes the log as real-time data. After real-time data is written to a Logstore, the data can be queried with a maximum latency of 3 seconds.
//
//     	- Historical data: The difference between the time record in a log and the current time on Simple Log Service is within the interval [-604,800 seconds,-180 seconds). For example, if a log was generated at 12:00:00, September 25, 2014 (UTC) and Simple Log Service received the log at 12:05:00, September 25, 2014 (UTC), Simple Log Service processes the log as historical data. This type of log is usually generated in data backfill scenarios.
//
//     After real-time data is written to a Logstore, the data can be queried with a maximum latency of 3 seconds. For 99.9% of queries, the latency is no more than 1 second.
//
// > Simple Log Service calculates the difference between the log time that is specified by the __time__ field and the receiving time that is specified by the __tag__:__receive_time__ field for each log. The receiving time indicates the time at which Simple Log Service receives the log. If the difference is within the interval (-180 seconds,900 seconds], Simple Log Service processes the log as real-time data. If the difference is within the interval [-604,800 seconds,-180 seconds), Simple Log Service processes the log as historical data.
//
// 	- Simple Log Service provides examples on how to call the GetHistograms operation by using Simple Log Service SDK for Java. For more information, see [Use GetHistograms to query the distribution of logs](https://help.aliyun.com/document_detail/462234.html).
//
// @param request - GetHistogramsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetHistogramsResponse
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

// Summary:
//
// Queries the distribution of logs that meet the specified search conditions in a Logstore.
//
// Description:
//
// ### Usage notes
//
// 	- Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
//
// 	- The time range is evenly divided into subintervals in the responses. If the time range that is specified in the request remains unchanged, the subintervals in the responses also remain unchanged.
//
// 	- If the number of logs in a Logstore significantly changes, Simple Log Service cannot predict the number of times that you must call this operation to obtain the complete result. In this case, you must check the value of the progress parameter in the response of each request and determine whether to call this operation one more time to obtain the complete result. Each time you call this operation, the same number of charge units (CUs) are consumed.
//
// 	- After a log is written to a Logstore, you can call the GetHistograms or GetLogs operation to query the log after a short latency. The latency of a query varies based on the type of the log. Simple Log Service classifies logs into the following types based on the log time:
//
//     	- Real-time data: The difference between the time record in a log and the current time on Simple Log Service is within the interval (-180 seconds,900 seconds]. For example, if a log was generated at 12:03:00, September 25, 2014 (UTC) and Simple Log Service received the log at 12:05:00, September 25, 2014 (UTC), Simple Log Service processes the log as real-time data. After real-time data is written to a Logstore, the data can be queried with a maximum latency of 3 seconds.
//
//     	- Historical data: The difference between the time record in a log and the current time on Simple Log Service is within the interval [-604,800 seconds,-180 seconds). For example, if a log was generated at 12:00:00, September 25, 2014 (UTC) and Simple Log Service received the log at 12:05:00, September 25, 2014 (UTC), Simple Log Service processes the log as historical data. This type of log is usually generated in data backfill scenarios.
//
//     After real-time data is written to a Logstore, the data can be queried with a maximum latency of 3 seconds. For 99.9% of queries, the latency is no more than 1 second.
//
// > Simple Log Service calculates the difference between the log time that is specified by the __time__ field and the receiving time that is specified by the __tag__:__receive_time__ field for each log. The receiving time indicates the time at which Simple Log Service receives the log. If the difference is within the interval (-180 seconds,900 seconds], Simple Log Service processes the log as real-time data. If the difference is within the interval [-604,800 seconds,-180 seconds), Simple Log Service processes the log as historical data.
//
// 	- Simple Log Service provides examples on how to call the GetHistograms operation by using Simple Log Service SDK for Java. For more information, see [Use GetHistograms to query the distribution of logs](https://help.aliyun.com/document_detail/462234.html).
//
// @param request - GetHistogramsRequest
//
// @return GetHistogramsResponse
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

// Summary:
//
// Queries the index of a Logstore.
//
// Description:
//
// ### Usage notes
//
// Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetIndexResponse
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

// Summary:
//
// Queries the index of a Logstore.
//
// Description:
//
// ### Usage notes
//
// Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
//
// @return GetIndexResponse
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

// Summary:
//
// Queries the details of a Logstore.
//
// Description:
//
// ### Usage notes
//
// Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetLogStoreResponse
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

// Summary:
//
// Queries the details of a Logstore.
//
// Description:
//
// ### Usage notes
//
// Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
//
// @return GetLogStoreResponse
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

// Summary:
//
// 获取LogStore计量模式
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetLogStoreMeteringModeResponse
func (client *Client) GetLogStoreMeteringModeWithOptions(project *string, logstore *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetLogStoreMeteringModeResponse, _err error) {
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetLogStoreMeteringMode"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/logstores/" + tea.StringValue(logstore) + "/meteringmode"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetLogStoreMeteringModeResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取LogStore计量模式
//
// @return GetLogStoreMeteringModeResponse
func (client *Client) GetLogStoreMeteringMode(project *string, logstore *string) (_result *GetLogStoreMeteringModeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetLogStoreMeteringModeResponse{}
	_body, _err := client.GetLogStoreMeteringModeWithOptions(project, logstore, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the service log configuration of a project.
//
// Description:
//
// Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetLoggingResponse
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

// Summary:
//
// Queries the service log configuration of a project.
//
// Description:
//
// Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
//
// @return GetLoggingResponse
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

// Summary:
//
// Queries the logs of a Logstore in a project.
//
// Description:
//
// ### Usage notes
//
// > Simple Log Service allows you to create a Scheduled SQL job. For more information, see [Create a Scheduled SQL job](https://help.aliyun.com/document_detail/286457.html).
//
// 	- Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
//
// 	- If the number of logs in a Logstore significantly changes, Simple Log Service cannot forecast the number of times that you must call this operation to obtain the complete result. In this case, you must check the value of the x-log-progress parameter in the response of each request and determine whether to call this operation one more time to obtain the complete result. Each time you call this operation, the same number of charge units (CUs) are consumed.
//
// 	- After a log is written to a Logstore, you can call the GetHistograms or GetLogs operation to query the log after a short latency. The latency of a query varies based on the type of the log. Simple Log Service classifies logs into the following types based on the log time:
//
//         Real-time data: The difference between the time record in a log and the current time on Simple Log Service is within the interval (-180 seconds,900 seconds]. For example, if a log was generated at 12:03:00, September 25, 2014 (UTC) and Simple Log Service received the log at 12:05:00, September 25, 2014 (UTC), Simple Log Service processes the log as real-time data. This type of log is usually generated in common scenarios.
//
//     	- Historical data: The difference between the time record in a log and the current time on Simple Log Service is within the interval [-604,800 seconds,-180 seconds). For example, if a log was generated at 12:00:00, September 25, 2014 (UTC) and Simple Log Service received the log at 12:05:00, September 25, 2014 (UTC), Simple Log Service processes the log as historical data. This type of log is usually generated in data backfill scenarios.
//
//     After real-time data is written to a Logstore, the data can be queried with a maximum latency of 3 seconds. For 99.9% of queries, the latency is no more than 1 second.
//
// > Simple Log Service calculates the difference between the log time that is specified by the __time__ field and the receiving time that is specified by the __tag__:**receive_time*	- field for each log. The receiving time indicates the time at which Simple Log Service receives the log. If the difference is within the interval (-180 seconds,900 seconds], Simple Log Service processes the log as real-time data. If the difference is within the interval [-604,800 seconds,-180 seconds), Simple Log Service processes the log as historical data.
//
// 	- Simple Log Service provides examples on how to call the GetLogs operation by using Simple Log Service SDK for Java and Simple Log Service SDK for Python. For more information, see [Examples of calling the GetLogs operation by using Simple Log Service SDK for Java](https://help.aliyun.com/document_detail/407683.html) and [Examples of calling the GetLogs operation by using Simple Log Service SDK for Python](https://help.aliyun.com/document_detail/407684.html).
//
// @param request - GetLogsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetLogsResponse
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

// Summary:
//
// Queries the logs of a Logstore in a project.
//
// Description:
//
// ### Usage notes
//
// > Simple Log Service allows you to create a Scheduled SQL job. For more information, see [Create a Scheduled SQL job](https://help.aliyun.com/document_detail/286457.html).
//
// 	- Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
//
// 	- If the number of logs in a Logstore significantly changes, Simple Log Service cannot forecast the number of times that you must call this operation to obtain the complete result. In this case, you must check the value of the x-log-progress parameter in the response of each request and determine whether to call this operation one more time to obtain the complete result. Each time you call this operation, the same number of charge units (CUs) are consumed.
//
// 	- After a log is written to a Logstore, you can call the GetHistograms or GetLogs operation to query the log after a short latency. The latency of a query varies based on the type of the log. Simple Log Service classifies logs into the following types based on the log time:
//
//         Real-time data: The difference between the time record in a log and the current time on Simple Log Service is within the interval (-180 seconds,900 seconds]. For example, if a log was generated at 12:03:00, September 25, 2014 (UTC) and Simple Log Service received the log at 12:05:00, September 25, 2014 (UTC), Simple Log Service processes the log as real-time data. This type of log is usually generated in common scenarios.
//
//     	- Historical data: The difference between the time record in a log and the current time on Simple Log Service is within the interval [-604,800 seconds,-180 seconds). For example, if a log was generated at 12:00:00, September 25, 2014 (UTC) and Simple Log Service received the log at 12:05:00, September 25, 2014 (UTC), Simple Log Service processes the log as historical data. This type of log is usually generated in data backfill scenarios.
//
//     After real-time data is written to a Logstore, the data can be queried with a maximum latency of 3 seconds. For 99.9% of queries, the latency is no more than 1 second.
//
// > Simple Log Service calculates the difference between the log time that is specified by the __time__ field and the receiving time that is specified by the __tag__:**receive_time*	- field for each log. The receiving time indicates the time at which Simple Log Service receives the log. If the difference is within the interval (-180 seconds,900 seconds], Simple Log Service processes the log as real-time data. If the difference is within the interval [-604,800 seconds,-180 seconds), Simple Log Service processes the log as historical data.
//
// 	- Simple Log Service provides examples on how to call the GetLogs operation by using Simple Log Service SDK for Java and Simple Log Service SDK for Python. For more information, see [Examples of calling the GetLogs operation by using Simple Log Service SDK for Java](https://help.aliyun.com/document_detail/407683.html) and [Examples of calling the GetLogs operation by using Simple Log Service SDK for Python](https://help.aliyun.com/document_detail/407684.html).
//
// @param request - GetLogsRequest
//
// @return GetLogsResponse
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

// Summary:
//
// Queries the raw log data in a Logstore of a project. The returned result shows the raw log data in a specific time range. The returned results are compressed and transmitted.
//
// Description:
//
//   Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
//
// 	- If the number of logs in a Logstore significantly changes, Simple Log Service cannot predict the number of times you must call this API operation to obtain a complete result. In this case, you must check the value of the progress parameter in the response of each request and determine whether to call this operation again to obtain the complete result. Each time you call this operation, the same number of charge units (CUs) are consumed.
//
// 	- After a log is written to a Logstore, you can call the GetHistograms or GetLogs operation to query the log. The latency of the query varies based on the type of the log. Simple Log Service classifies logs into the following types based on log timestamps:
//
// 1.  1.  Real-time data: The difference between the time record in the log and the current server time is within the interval (-180 seconds,900 seconds]. For example, if a log was generated at 12:03:00, September 25, 2014 (UTC) and the server received the log at 12:05:00, September 25, 2014 (UTC), the server processes the log as real-time data. This type of log is usually generated in common scenarios.
//
// 2.  2.  Historical data: The difference between the time record in the log and the current server time is within the interval [-604,800 seconds,-180 seconds). For example, if a log was generated at 12:00:00, September 25, 2014 (UTC) and the server received the log at 12:05:00, September 25, 2014 (UTC), the server processes the log as historical data. This type of log is usually generated in data backfill scenarios. After real-time data is written to a Logstore, the data can be queried with a maximum latency of 3 seconds. For 99.9% of queries, the latency is no more than 1 second.
//
// @param request - GetLogsV2Request
//
// @param headers - GetLogsV2Headers
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetLogsV2Response
func (client *Client) GetLogsV2WithOptions(project *string, logstore *string, request *GetLogsV2Request, headers *GetLogsV2Headers, runtime *util.RuntimeOptions) (_result *GetLogsV2Response, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Forward)) {
		body["forward"] = request.Forward
	}

	if !tea.BoolValue(util.IsUnset(request.From)) {
		body["from"] = request.From
	}

	if !tea.BoolValue(util.IsUnset(request.Highlight)) {
		body["highlight"] = request.Highlight
	}

	if !tea.BoolValue(util.IsUnset(request.Line)) {
		body["line"] = request.Line
	}

	if !tea.BoolValue(util.IsUnset(request.Offset)) {
		body["offset"] = request.Offset
	}

	if !tea.BoolValue(util.IsUnset(request.PowerSql)) {
		body["powerSql"] = request.PowerSql
	}

	if !tea.BoolValue(util.IsUnset(request.Query)) {
		body["query"] = request.Query
	}

	if !tea.BoolValue(util.IsUnset(request.Reverse)) {
		body["reverse"] = request.Reverse
	}

	if !tea.BoolValue(util.IsUnset(request.Session)) {
		body["session"] = request.Session
	}

	if !tea.BoolValue(util.IsUnset(request.To)) {
		body["to"] = request.To
	}

	if !tea.BoolValue(util.IsUnset(request.Topic)) {
		body["topic"] = request.Topic
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.AcceptEncoding)) {
		realHeaders["Accept-Encoding"] = util.ToJSONString(headers.AcceptEncoding)
	}

	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetLogsV2"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/logstores/" + tea.StringValue(logstore) + "/logs"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetLogsV2Response{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the raw log data in a Logstore of a project. The returned result shows the raw log data in a specific time range. The returned results are compressed and transmitted.
//
// Description:
//
//   Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
//
// 	- If the number of logs in a Logstore significantly changes, Simple Log Service cannot predict the number of times you must call this API operation to obtain a complete result. In this case, you must check the value of the progress parameter in the response of each request and determine whether to call this operation again to obtain the complete result. Each time you call this operation, the same number of charge units (CUs) are consumed.
//
// 	- After a log is written to a Logstore, you can call the GetHistograms or GetLogs operation to query the log. The latency of the query varies based on the type of the log. Simple Log Service classifies logs into the following types based on log timestamps:
//
// 1.  1.  Real-time data: The difference between the time record in the log and the current server time is within the interval (-180 seconds,900 seconds]. For example, if a log was generated at 12:03:00, September 25, 2014 (UTC) and the server received the log at 12:05:00, September 25, 2014 (UTC), the server processes the log as real-time data. This type of log is usually generated in common scenarios.
//
// 2.  2.  Historical data: The difference between the time record in the log and the current server time is within the interval [-604,800 seconds,-180 seconds). For example, if a log was generated at 12:00:00, September 25, 2014 (UTC) and the server received the log at 12:05:00, September 25, 2014 (UTC), the server processes the log as historical data. This type of log is usually generated in data backfill scenarios. After real-time data is written to a Logstore, the data can be queried with a maximum latency of 3 seconds. For 99.9% of queries, the latency is no more than 1 second.
//
// @param request - GetLogsV2Request
//
// @return GetLogsV2Response
func (client *Client) GetLogsV2(project *string, logstore *string, request *GetLogsV2Request) (_result *GetLogsV2Response, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetLogsV2Headers{}
	_result = &GetLogsV2Response{}
	_body, _err := client.GetLogsV2WithOptions(project, logstore, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of a Logtail pipeline configuration.
//
// Description:
//
// The UK (London) region is supported. Supported regions are constantly updated.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetLogtailPipelineConfigResponse
func (client *Client) GetLogtailPipelineConfigWithOptions(project *string, configName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetLogtailPipelineConfigResponse, _err error) {
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetLogtailPipelineConfig"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/pipelineconfigs/" + tea.StringValue(configName)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetLogtailPipelineConfigResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a Logtail pipeline configuration.
//
// Description:
//
// The UK (London) region is supported. Supported regions are constantly updated.
//
// @return GetLogtailPipelineConfigResponse
func (client *Client) GetLogtailPipelineConfig(project *string, configName *string) (_result *GetLogtailPipelineConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetLogtailPipelineConfigResponse{}
	_body, _err := client.GetLogtailPipelineConfigWithOptions(project, configName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// GetMLServiceResults
//
// @param request - GetMLServiceResultsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMLServiceResultsResponse
func (client *Client) GetMLServiceResultsWithOptions(serviceName *string, request *GetMLServiceResultsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetMLServiceResultsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AllowBuiltin)) {
		query["allowBuiltin"] = request.AllowBuiltin
	}

	if !tea.BoolValue(util.IsUnset(request.Version)) {
		query["version"] = request.Version
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetMLServiceResults"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/ml/service/" + tea.StringValue(serviceName) + "/analysis"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetMLServiceResultsResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// GetMLServiceResults
//
// @param request - GetMLServiceResultsRequest
//
// @return GetMLServiceResultsResponse
func (client *Client) GetMLServiceResults(serviceName *string, request *GetMLServiceResultsRequest) (_result *GetMLServiceResultsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetMLServiceResultsResponse{}
	_body, _err := client.GetMLServiceResultsWithOptions(serviceName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of a machine group.
//
// Description:
//
// Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMachineGroupResponse
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

// Summary:
//
// Queries the details of a machine group.
//
// Description:
//
// Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
//
// @return GetMachineGroupResponse
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

// Summary:
//
// 获取 MetricStore 计量模式
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMetricStoreMeteringModeResponse
func (client *Client) GetMetricStoreMeteringModeWithOptions(project *string, metricStore *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetMetricStoreMeteringModeResponse, _err error) {
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetMetricStoreMeteringMode"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/metricstores/" + tea.StringValue(metricStore) + "/meteringmode"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetMetricStoreMeteringModeResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取 MetricStore 计量模式
//
// @return GetMetricStoreMeteringModeResponse
func (client *Client) GetMetricStoreMeteringMode(project *string, metricStore *string) (_result *GetMetricStoreMeteringModeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetMetricStoreMeteringModeResponse{}
	_body, _err := client.GetMetricStoreMeteringModeWithOptions(project, metricStore, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取OSS投递任务
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetOSSExportResponse
func (client *Client) GetOSSExportWithOptions(project *string, ossExportName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetOSSExportResponse, _err error) {
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetOSSExport"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/ossexports/" + tea.StringValue(ossExportName)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetOSSExportResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取OSS投递任务
//
// @return GetOSSExportResponse
func (client *Client) GetOSSExport(project *string, ossExportName *string) (_result *GetOSSExportResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetOSSExportResponse{}
	_body, _err := client.GetOSSExportWithOptions(project, ossExportName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Get OSSHDFS Exports
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetOSSHDFSExportResponse
func (client *Client) GetOSSHDFSExportWithOptions(project *string, ossExportName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetOSSHDFSExportResponse, _err error) {
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetOSSHDFSExport"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/osshdfsexports/" + tea.StringValue(ossExportName)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetOSSHDFSExportResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Get OSSHDFS Exports
//
// @return GetOSSHDFSExportResponse
func (client *Client) GetOSSHDFSExport(project *string, ossExportName *string) (_result *GetOSSHDFSExportResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetOSSHDFSExportResponse{}
	_body, _err := client.GetOSSHDFSExportWithOptions(project, ossExportName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取oss导入任务信息
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetOSSIngestionResponse
func (client *Client) GetOSSIngestionWithOptions(project *string, ossIngestionName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetOSSIngestionResponse, _err error) {
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetOSSIngestion"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/ossingestions/" + tea.StringValue(ossIngestionName)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetOSSIngestionResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取oss导入任务信息
//
// @return GetOSSIngestionResponse
func (client *Client) GetOSSIngestion(project *string, ossIngestionName *string) (_result *GetOSSIngestionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetOSSIngestionResponse{}
	_body, _err := client.GetOSSIngestionWithOptions(project, ossIngestionName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of a project.
//
// Description:
//
// ### Usage notes
//
// Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetProjectResponse
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

// Summary:
//
// Queries the details of a project.
//
// Description:
//
// ### Usage notes
//
// Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
//
// @return GetProjectResponse
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

// Summary:
//
// Queries logs in a project. You can use this operation to query logs at the project level.
//
// Description:
//
// ### Usage notes
//
// 	- You can use the query parameter to specify a standard SQL statement.
//
// 	- You must specify a project in the domain name of the request.
//
// 	- You must specify a Logstore in the FROM clause of the SQL statement. A Logstore can be used as an SQL table.
//
// 	- You must specify a time range in the SQL statement by using the __date__ parameter or __time__ parameter. The value of the __date__ parameter is a timestamp, and the value of the __time__ parameter is an integer. The unit of the __time__ parameter is seconds.
//
// 	- Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
//
// @param request - GetProjectLogsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetProjectLogsResponse
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

// Summary:
//
// Queries logs in a project. You can use this operation to query logs at the project level.
//
// Description:
//
// ### Usage notes
//
// 	- You can use the query parameter to specify a standard SQL statement.
//
// 	- You must specify a project in the domain name of the request.
//
// 	- You must specify a Logstore in the FROM clause of the SQL statement. A Logstore can be used as an SQL table.
//
// 	- You must specify a time range in the SQL statement by using the __date__ parameter or __time__ parameter. The value of the __date__ parameter is a timestamp, and the value of the __time__ parameter is an integer. The unit of the __time__ parameter is seconds.
//
// 	- Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
//
// @param request - GetProjectLogsRequest
//
// @return GetProjectLogsResponse
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

// Summary:
//
// Queries a project policy.
//
// Description:
//
// ### Usage notes
//
// Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetProjectPolicyResponse
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

// Summary:
//
// Queries a project policy.
//
// Description:
//
// ### Usage notes
//
// Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
//
// @return GetProjectPolicyResponse
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

// Summary:
//
// Queries a saved search.
//
// Description:
//
// ### Usage notes
//
// Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSavedSearchResponse
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

// Summary:
//
// Queries a saved search.
//
// Description:
//
// ### Usage notes
//
// Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
//
// @return GetSavedSearchResponse
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

// Summary:
//
// 查看定时SQL任务
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetScheduledSQLResponse
func (client *Client) GetScheduledSQLWithOptions(project *string, scheduledSQLName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetScheduledSQLResponse, _err error) {
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetScheduledSQL"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/scheduledsqls/" + tea.StringValue(scheduledSQLName)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetScheduledSQLResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查看定时SQL任务
//
// @return GetScheduledSQLResponse
func (client *Client) GetScheduledSQL(project *string, scheduledSQLName *string) (_result *GetScheduledSQLResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetScheduledSQLResponse{}
	_body, _err := client.GetScheduledSQLWithOptions(project, scheduledSQLName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI GetShipperStatus is deprecated
//
// Summary:
//
// Queries the status of a log shipping job.
//
// Description:
//
// Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
//
// @param request - GetShipperStatusRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetShipperStatusResponse
// Deprecated
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

// Deprecated: OpenAPI GetShipperStatus is deprecated
//
// Summary:
//
// Queries the status of a log shipping job.
//
// Description:
//
// Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
//
// @param request - GetShipperStatusRequest
//
// @return GetShipperStatusResponse
// Deprecated
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

// Summary:
//
// getSlsService
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSlsServiceResponse
func (client *Client) GetSlsServiceWithOptions(headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetSlsServiceResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetSlsService"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/slsservice"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetSlsServiceResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// getSlsService
//
// @return GetSlsServiceResponse
func (client *Client) GetSlsService() (_result *GetSlsServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetSlsServiceResponse{}
	_body, _err := client.GetSlsServiceWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询独享sql实例
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSqlInstanceResponse
func (client *Client) GetSqlInstanceWithOptions(project *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetSqlInstanceResponse, _err error) {
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetSqlInstance"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/sqlinstance"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("array"),
	}
	_result = &GetSqlInstanceResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询独享sql实例
//
// @return GetSqlInstanceResponse
func (client *Client) GetSqlInstance(project *string) (_result *GetSqlInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetSqlInstanceResponse{}
	_body, _err := client.GetSqlInstanceWithOptions(project, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询StoreView
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetStoreViewResponse
func (client *Client) GetStoreViewWithOptions(project *string, name *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetStoreViewResponse, _err error) {
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetStoreView"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/storeviews/" + tea.StringValue(name)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetStoreViewResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询StoreView
//
// @return GetStoreViewResponse
func (client *Client) GetStoreView(project *string, name *string) (_result *GetStoreViewResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetStoreViewResponse{}
	_body, _err := client.GetStoreViewWithOptions(project, name, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询StoreView索引
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetStoreViewIndexResponse
func (client *Client) GetStoreViewIndexWithOptions(project *string, name *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetStoreViewIndexResponse, _err error) {
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetStoreViewIndex"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/storeviews/" + tea.StringValue(name) + "/index"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetStoreViewIndexResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询StoreView索引
//
// @return GetStoreViewIndexResponse
func (client *Client) GetStoreViewIndex(project *string, name *string) (_result *GetStoreViewIndexResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetStoreViewIndexResponse{}
	_body, _err := client.GetStoreViewIndexWithOptions(project, name, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询告警列表
//
// @param request - ListAlertsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAlertsResponse
func (client *Client) ListAlertsWithOptions(project *string, request *ListAlertsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListAlertsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Logstore)) {
		query["logstore"] = request.Logstore
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
		Action:      tea.String("ListAlerts"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/alerts"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAlertsResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询告警列表
//
// @param request - ListAlertsRequest
//
// @return ListAlertsResponse
func (client *Client) ListAlerts(project *string, request *ListAlertsRequest) (_result *ListAlertsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListAlertsResponse{}
	_body, _err := client.ListAlertsWithOptions(project, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries data in a dataset.
//
// @param request - ListAnnotationDataRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAnnotationDataResponse
func (client *Client) ListAnnotationDataWithOptions(datasetId *string, request *ListAnnotationDataRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListAnnotationDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Offset)) {
		query["offset"] = request.Offset
	}

	if !tea.BoolValue(util.IsUnset(request.Size)) {
		query["size"] = request.Size
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListAnnotationData"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/ml/annotationdataset/" + tea.StringValue(datasetId) + "/annotationdata"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAnnotationDataResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries data in a dataset.
//
// @param request - ListAnnotationDataRequest
//
// @return ListAnnotationDataResponse
func (client *Client) ListAnnotationData(datasetId *string, request *ListAnnotationDataRequest) (_result *ListAnnotationDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListAnnotationDataResponse{}
	_body, _err := client.ListAnnotationDataWithOptions(datasetId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of datasets.
//
// @param request - ListAnnotationDataSetsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAnnotationDataSetsResponse
func (client *Client) ListAnnotationDataSetsWithOptions(request *ListAnnotationDataSetsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListAnnotationDataSetsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Offset)) {
		query["offset"] = request.Offset
	}

	if !tea.BoolValue(util.IsUnset(request.Size)) {
		query["size"] = request.Size
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListAnnotationDataSets"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/ml/annotationdataset"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAnnotationDataSetsResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of datasets.
//
// @param request - ListAnnotationDataSetsRequest
//
// @return ListAnnotationDataSetsResponse
func (client *Client) ListAnnotationDataSets(request *ListAnnotationDataSetsRequest) (_result *ListAnnotationDataSetsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListAnnotationDataSetsResponse{}
	_body, _err := client.ListAnnotationDataSetsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of tag tables.
//
// @param request - ListAnnotationLabelsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAnnotationLabelsResponse
func (client *Client) ListAnnotationLabelsWithOptions(request *ListAnnotationLabelsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListAnnotationLabelsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Offset)) {
		query["offset"] = request.Offset
	}

	if !tea.BoolValue(util.IsUnset(request.Size)) {
		query["size"] = request.Size
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListAnnotationLabels"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/ml/annotationlabel"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAnnotationLabelsResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of tag tables.
//
// @param request - ListAnnotationLabelsRequest
//
// @return ListAnnotationLabelsResponse
func (client *Client) ListAnnotationLabels(request *ListAnnotationLabelsRequest) (_result *ListAnnotationLabelsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListAnnotationLabelsResponse{}
	_body, _err := client.ListAnnotationLabelsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 通过调用ListCollectionPolicies接口查看配置的日志采集规则
//
// @param tmpReq - ListCollectionPoliciesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCollectionPoliciesResponse
func (client *Client) ListCollectionPoliciesWithOptions(tmpReq *ListCollectionPoliciesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListCollectionPoliciesResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListCollectionPoliciesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Attribute)) {
		request.AttributeShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Attribute, tea.String("attribute"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AttributeShrink)) {
		query["attribute"] = request.AttributeShrink
	}

	if !tea.BoolValue(util.IsUnset(request.DataCode)) {
		query["dataCode"] = request.DataCode
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["instanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNum)) {
		query["pageNum"] = request.PageNum
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.PolicyName)) {
		query["policyName"] = request.PolicyName
	}

	if !tea.BoolValue(util.IsUnset(request.ProductCode)) {
		query["productCode"] = request.ProductCode
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListCollectionPolicies"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/collectionpolicy"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListCollectionPoliciesResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 通过调用ListCollectionPolicies接口查看配置的日志采集规则
//
// @param request - ListCollectionPoliciesRequest
//
// @return ListCollectionPoliciesResponse
func (client *Client) ListCollectionPolicies(request *ListCollectionPoliciesRequest) (_result *ListCollectionPoliciesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListCollectionPoliciesResponse{}
	_body, _err := client.ListCollectionPoliciesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries all Logtail configurations in a project.
//
// Description:
//
// ### [](#)Usage notes
//
// 	- Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
//
// 	- An AccessKey pair is created and obtained. For more information, see [AccessKey pair](https://help.aliyun.com/document_detail/29009.html).
//
// The AccessKey pair of an Alibaba Cloud account has permissions on all API operations. Using these credentials to perform operations in Simple Log Service is a high-risk operation. We recommend that you use a RAM user to call API operations or perform routine O\\&M. To create a RAM user, log on to the RAM console. Make sure that the RAM user has the management permissions on Simple Log Service resources. For more information, see [Create a RAM user and authorize the RAM user to access Simple Log Service](https://help.aliyun.com/document_detail/47664.html).
//
// 	- The information that is required to query logs is obtained. The information includes the name of the project to which the logs belong and the region of the project. For more information, see [Manage a project](https://help.aliyun.com/document_detail/48984.html).
//
// @param request - ListConfigRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListConfigResponse
func (client *Client) ListConfigWithOptions(project *string, request *ListConfigRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConfigName)) {
		query["configName"] = request.ConfigName
	}

	if !tea.BoolValue(util.IsUnset(request.LogstoreName)) {
		query["logstoreName"] = request.LogstoreName
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
		Action:      tea.String("ListConfig"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/configs"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListConfigResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries all Logtail configurations in a project.
//
// Description:
//
// ### [](#)Usage notes
//
// 	- Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
//
// 	- An AccessKey pair is created and obtained. For more information, see [AccessKey pair](https://help.aliyun.com/document_detail/29009.html).
//
// The AccessKey pair of an Alibaba Cloud account has permissions on all API operations. Using these credentials to perform operations in Simple Log Service is a high-risk operation. We recommend that you use a RAM user to call API operations or perform routine O\\&M. To create a RAM user, log on to the RAM console. Make sure that the RAM user has the management permissions on Simple Log Service resources. For more information, see [Create a RAM user and authorize the RAM user to access Simple Log Service](https://help.aliyun.com/document_detail/47664.html).
//
// 	- The information that is required to query logs is obtained. The information includes the name of the project to which the logs belong and the region of the project. For more information, see [Manage a project](https://help.aliyun.com/document_detail/48984.html).
//
// @param request - ListConfigRequest
//
// @return ListConfigResponse
func (client *Client) ListConfig(project *string, request *ListConfigRequest) (_result *ListConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListConfigResponse{}
	_body, _err := client.ListConfigWithOptions(project, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries all consumer groups of a Logstore.
//
// Description:
//
// ### Usage notes
//
// Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListConsumerGroupResponse
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

// Summary:
//
// Queries all consumer groups of a Logstore.
//
// Description:
//
// ### Usage notes
//
// Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
//
// @return ListConsumerGroupResponse
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

// Summary:
//
// Queries a list of dashboards.
//
// Description:
//
// ### [](#)Usage notes
//
// 	- Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
//
// 	- An AccessKey pair is created and obtained. For more information, see [AccessKey pair](https://help.aliyun.com/document_detail/29009.html).
//
// The AccessKey pair of an Alibaba Cloud account has permissions on all API operations. Using these credentials to perform operations in Simple Log Service is a high-risk operation. We recommend that you use a RAM user to call API operations or perform routine O\\&M. To create a RAM user, log on to the RAM console. Make sure that the RAM user has the management permissions on Simple Log Service resources. For more information, see [Create a RAM user and authorize the RAM user to access Simple Log Service](https://help.aliyun.com/document_detail/47664.html).
//
// 	- The information that is required to query logs is obtained. The information includes the name of the project to which the logs belong and the region of the project. For more information, see [Manage a project](https://help.aliyun.com/document_detail/48984.html).
//
// @param request - ListDashboardRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDashboardResponse
func (client *Client) ListDashboardWithOptions(project *string, request *ListDashboardRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListDashboardResponse, _err error) {
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
		Action:      tea.String("ListDashboard"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/dashboards"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListDashboardResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of dashboards.
//
// Description:
//
// ### [](#)Usage notes
//
// 	- Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
//
// 	- An AccessKey pair is created and obtained. For more information, see [AccessKey pair](https://help.aliyun.com/document_detail/29009.html).
//
// The AccessKey pair of an Alibaba Cloud account has permissions on all API operations. Using these credentials to perform operations in Simple Log Service is a high-risk operation. We recommend that you use a RAM user to call API operations or perform routine O\\&M. To create a RAM user, log on to the RAM console. Make sure that the RAM user has the management permissions on Simple Log Service resources. For more information, see [Create a RAM user and authorize the RAM user to access Simple Log Service](https://help.aliyun.com/document_detail/47664.html).
//
// 	- The information that is required to query logs is obtained. The information includes the name of the project to which the logs belong and the region of the project. For more information, see [Manage a project](https://help.aliyun.com/document_detail/48984.html).
//
// @param request - ListDashboardRequest
//
// @return ListDashboardResponse
func (client *Client) ListDashboard(project *string, request *ListDashboardRequest) (_result *ListDashboardResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListDashboardResponse{}
	_body, _err := client.ListDashboardWithOptions(project, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the custom domain names that are bound to projects.
//
// Description:
//
//   Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
//
// 	- Only one custom domain name can be bound to each project.
//
// @param request - ListDomainsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDomainsResponse
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

// Summary:
//
// Queries the custom domain names that are bound to projects.
//
// Description:
//
//   Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
//
// 	- Only one custom domain name can be bound to each project.
//
// @param request - ListDomainsRequest
//
// @return ListDomainsResponse
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

// Summary:
//
// 列出数据加工任务
//
// @param request - ListETLsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListETLsResponse
func (client *Client) ListETLsWithOptions(project *string, request *ListETLsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListETLsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Logstore)) {
		query["logstore"] = request.Logstore
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
		Action:      tea.String("ListETLs"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/etls"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListETLsResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 列出数据加工任务
//
// @param request - ListETLsRequest
//
// @return ListETLsResponse
func (client *Client) ListETLs(project *string, request *ListETLsRequest) (_result *ListETLsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListETLsResponse{}
	_body, _err := client.ListETLsWithOptions(project, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of external stores.
//
// Description:
//
// Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
//
// @param request - ListExternalStoreRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListExternalStoreResponse
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

// Summary:
//
// Queries a list of external stores.
//
// Description:
//
// Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
//
// @param request - ListExternalStoreRequest
//
// @return ListExternalStoreResponse
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

// Summary:
//
// Queries all Logstores or Logstores that match specific conditions in a project.
//
// Description:
//
// ### Usage notes
//
// 	- Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
//
// 	- An AccessKey pair is created and obtained. For more information, see [AccessKey pair](https://help.aliyun.com/document_detail/29009.html).
//
// The AccessKey pair of an Alibaba Cloud account has permissions on all API operations. Using these credentials to perform operations in Simple Log Service is a high-risk operation. We recommend that you use a RAM user to call API operations or perform routine O&#x26;M. To create a RAM user, log on to the RAM console. Make sure that the RAM user has the management permissions on Simple Log Service resources. For more information, see [Create a RAM user and authorize the RAM user to access Simple Log Service](https://help.aliyun.com/document_detail/47664.html).
//
// 	- The information that is required to query logs is obtained. The information includes the name of the project to which the logs belong and the region of the project. For more information, see [Manage a project](https://help.aliyun.com/document_detail/48984.html).
//
// ### Authentication resources
//
// The following table describes the authorization information that is required for this operation. You can add the information to the Action element of a RAM policy statement to grant a RAM user or a RAM role the permissions to call this operation.
//
// |Action|Resource|
//
// |:---|:---|
//
// |`log:ListLogStores`|`acs:log:{#regionId}:{#accountId}:project/{#ProjectName}/logstore/*`|
//
// @param request - ListLogStoresRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListLogStoresResponse
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

// Summary:
//
// Queries all Logstores or Logstores that match specific conditions in a project.
//
// Description:
//
// ### Usage notes
//
// 	- Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
//
// 	- An AccessKey pair is created and obtained. For more information, see [AccessKey pair](https://help.aliyun.com/document_detail/29009.html).
//
// The AccessKey pair of an Alibaba Cloud account has permissions on all API operations. Using these credentials to perform operations in Simple Log Service is a high-risk operation. We recommend that you use a RAM user to call API operations or perform routine O&#x26;M. To create a RAM user, log on to the RAM console. Make sure that the RAM user has the management permissions on Simple Log Service resources. For more information, see [Create a RAM user and authorize the RAM user to access Simple Log Service](https://help.aliyun.com/document_detail/47664.html).
//
// 	- The information that is required to query logs is obtained. The information includes the name of the project to which the logs belong and the region of the project. For more information, see [Manage a project](https://help.aliyun.com/document_detail/48984.html).
//
// ### Authentication resources
//
// The following table describes the authorization information that is required for this operation. You can add the information to the Action element of a RAM policy statement to grant a RAM user or a RAM role the permissions to call this operation.
//
// |Action|Resource|
//
// |:---|:---|
//
// |`log:ListLogStores`|`acs:log:{#regionId}:{#accountId}:project/{#ProjectName}/logstore/*`|
//
// @param request - ListLogStoresRequest
//
// @return ListLogStoresResponse
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

// Summary:
//
// Queries a list of Logtail pipeline configurations that meet the specified conditions.
//
// Description:
//
// The UK (London) region is supported. Supported regions are constantly updated.
//
// @param request - ListLogtailPipelineConfigRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListLogtailPipelineConfigResponse
func (client *Client) ListLogtailPipelineConfigWithOptions(project *string, request *ListLogtailPipelineConfigRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListLogtailPipelineConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConfigName)) {
		query["configName"] = request.ConfigName
	}

	if !tea.BoolValue(util.IsUnset(request.LogstoreName)) {
		query["logstoreName"] = request.LogstoreName
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
		Action:      tea.String("ListLogtailPipelineConfig"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/pipelineconfigs"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListLogtailPipelineConfigResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of Logtail pipeline configurations that meet the specified conditions.
//
// Description:
//
// The UK (London) region is supported. Supported regions are constantly updated.
//
// @param request - ListLogtailPipelineConfigRequest
//
// @return ListLogtailPipelineConfigResponse
func (client *Client) ListLogtailPipelineConfig(project *string, request *ListLogtailPipelineConfigRequest) (_result *ListLogtailPipelineConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListLogtailPipelineConfigResponse{}
	_body, _err := client.ListLogtailPipelineConfigWithOptions(project, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the machine groups of a project.
//
// Description:
//
// Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
//
// @param request - ListMachineGroupRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMachineGroupResponse
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

// Summary:
//
// Queries the machine groups of a project.
//
// Description:
//
// Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
//
// @param request - ListMachineGroupRequest
//
// @return ListMachineGroupResponse
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

// Summary:
//
// Queries a list of machines that are connected to Simple Log Service in a specified machine group.
//
// Description:
//
// Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
//
// @param request - ListMachinesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMachinesResponse
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

// Summary:
//
// Queries a list of machines that are connected to Simple Log Service in a specified machine group.
//
// Description:
//
// Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
//
// @param request - ListMachinesRequest
//
// @return ListMachinesResponse
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

// Summary:
//
// 列出OSS投递任务
//
// @param request - ListOSSExportsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListOSSExportsResponse
func (client *Client) ListOSSExportsWithOptions(project *string, request *ListOSSExportsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListOSSExportsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Logstore)) {
		query["logstore"] = request.Logstore
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
		Action:      tea.String("ListOSSExports"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/ossexports"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListOSSExportsResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 列出OSS投递任务
//
// @param request - ListOSSExportsRequest
//
// @return ListOSSExportsResponse
func (client *Client) ListOSSExports(project *string, request *ListOSSExportsRequest) (_result *ListOSSExportsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListOSSExportsResponse{}
	_body, _err := client.ListOSSExportsWithOptions(project, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 列举OSSHDFS投递任务
//
// @param request - ListOSSHDFSExportsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListOSSHDFSExportsResponse
func (client *Client) ListOSSHDFSExportsWithOptions(project *string, request *ListOSSHDFSExportsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListOSSHDFSExportsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Logstore)) {
		query["logstore"] = request.Logstore
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
		Action:      tea.String("ListOSSHDFSExports"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/osshdfsexports"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListOSSHDFSExportsResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 列举OSSHDFS投递任务
//
// @param request - ListOSSHDFSExportsRequest
//
// @return ListOSSHDFSExportsResponse
func (client *Client) ListOSSHDFSExports(project *string, request *ListOSSHDFSExportsRequest) (_result *ListOSSHDFSExportsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListOSSHDFSExportsResponse{}
	_body, _err := client.ListOSSHDFSExportsWithOptions(project, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 列出oss导入任务
//
// @param request - ListOSSIngestionsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListOSSIngestionsResponse
func (client *Client) ListOSSIngestionsWithOptions(project *string, request *ListOSSIngestionsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListOSSIngestionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Logstore)) {
		query["logstore"] = request.Logstore
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
		Action:      tea.String("ListOSSIngestions"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/ossingestions"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListOSSIngestionsResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 列出oss导入任务
//
// @param request - ListOSSIngestionsRequest
//
// @return ListOSSIngestionsResponse
func (client *Client) ListOSSIngestions(project *string, request *ListOSSIngestionsRequest) (_result *ListOSSIngestionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListOSSIngestionsResponse{}
	_body, _err := client.ListOSSIngestionsWithOptions(project, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the projects that meet specified conditions.
//
// Description:
//
// ### [](#)Usage notes
//
// Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
//
// @param request - ListProjectRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListProjectResponse
func (client *Client) ListProjectWithOptions(request *ListProjectRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListProjectResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["resourceGroupId"] = request.ResourceGroupId
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

// Summary:
//
// Queries the projects that meet specified conditions.
//
// Description:
//
// ### [](#)Usage notes
//
// Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
//
// @param request - ListProjectRequest
//
// @return ListProjectResponse
func (client *Client) ListProject(request *ListProjectRequest) (_result *ListProjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListProjectResponse{}
	_body, _err := client.ListProjectWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of saved searches.
//
// Description:
//
// ### Usage notes
//
// Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
//
// @param request - ListSavedSearchRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSavedSearchResponse
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

// Summary:
//
// Queries a list of saved searches.
//
// Description:
//
// ### Usage notes
//
// Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
//
// @param request - ListSavedSearchRequest
//
// @return ListSavedSearchResponse
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

// Summary:
//
// 列举定时SQL任务
//
// @param request - ListScheduledSQLsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListScheduledSQLsResponse
func (client *Client) ListScheduledSQLsWithOptions(project *string, request *ListScheduledSQLsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListScheduledSQLsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Logstore)) {
		query["logstore"] = request.Logstore
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
		Action:      tea.String("ListScheduledSQLs"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/scheduledsqls"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListScheduledSQLsResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 列举定时SQL任务
//
// @param request - ListScheduledSQLsRequest
//
// @return ListScheduledSQLsResponse
func (client *Client) ListScheduledSQLs(project *string, request *ListScheduledSQLsRequest) (_result *ListScheduledSQLsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListScheduledSQLsResponse{}
	_body, _err := client.ListScheduledSQLsWithOptions(project, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of shards in a Logstore.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListShardsResponse
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

// Summary:
//
// Queries a list of shards in a Logstore.
//
// @return ListShardsResponse
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

// Deprecated: OpenAPI ListShipper is deprecated
//
// Summary:
//
// Queries a list of log shipping jobs in a Logstore.
//
// Description:
//
// Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListShipperResponse
// Deprecated
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

// Deprecated: OpenAPI ListShipper is deprecated
//
// Summary:
//
// Queries a list of log shipping jobs in a Logstore.
//
// Description:
//
// Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
//
// @return ListShipperResponse
// Deprecated
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

// Summary:
//
// 查询StoreView列表
//
// @param request - ListStoreViewsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListStoreViewsResponse
func (client *Client) ListStoreViewsWithOptions(project *string, request *ListStoreViewsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListStoreViewsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Offset)) {
		query["offset"] = request.Offset
	}

	if !tea.BoolValue(util.IsUnset(request.Size)) {
		query["size"] = request.Size
	}

	if !tea.BoolValue(util.IsUnset(request.StoreType)) {
		query["storeType"] = request.StoreType
	}

	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListStoreViews"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/storeviews"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListStoreViewsResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询StoreView列表
//
// @param request - ListStoreViewsRequest
//
// @return ListStoreViewsResponse
func (client *Client) ListStoreViews(project *string, request *ListStoreViewsRequest) (_result *ListStoreViewsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListStoreViewsResponse{}
	_body, _err := client.ListStoreViewsWithOptions(project, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of tags for one or more resources. You can query tags for resources by resource type or filter resources by tag. Each tag is a key-value pair.
//
// Description:
//
// ### Usage notes
//
// Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
//
// @param tmpReq - ListTagResourcesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTagResourcesResponse
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

// Summary:
//
// Queries a list of tags for one or more resources. You can query tags for resources by resource type or filter resources by tag. Each tag is a key-value pair.
//
// Description:
//
// ### Usage notes
//
// Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
//
// @param request - ListTagResourcesRequest
//
// @return ListTagResourcesResponse
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

// Summary:
//
// 合并两个相邻的readwrite状态的Shards。在参数中指定一个shardID，服务端自动找相邻的下一个Shard进行合并。
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return MergeShardResponse
func (client *Client) MergeShardWithOptions(project *string, logstore *string, shard *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *MergeShardResponse, _err error) {
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("MergeShard"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/logstores/" + tea.StringValue(logstore) + "/shards/" + tea.StringValue(shard) + "?action=merge"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("array"),
	}
	_result = &MergeShardResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 合并两个相邻的readwrite状态的Shards。在参数中指定一个shardID，服务端自动找相邻的下一个Shard进行合并。
//
// @return MergeShardResponse
func (client *Client) MergeShard(project *string, logstore *string, shard *string) (_result *MergeShardResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &MergeShardResponse{}
	_body, _err := client.MergeShardWithOptions(project, logstore, shard, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// openSlsService
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OpenSlsServiceResponse
func (client *Client) OpenSlsServiceWithOptions(headers map[string]*string, runtime *util.RuntimeOptions) (_result *OpenSlsServiceResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("OpenSlsService"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/slsservice"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &OpenSlsServiceResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// openSlsService
//
// @return OpenSlsServiceResponse
func (client *Client) OpenSlsService() (_result *OpenSlsServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &OpenSlsServiceResponse{}
	_body, _err := client.OpenSlsServiceWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds data to a dataset for storage.
//
// @param request - PutAnnotationDataRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PutAnnotationDataResponse
func (client *Client) PutAnnotationDataWithOptions(datasetId *string, request *PutAnnotationDataRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *PutAnnotationDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AnnotationdataId)) {
		query["annotationdataId"] = request.AnnotationdataId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MlDataParam)) {
		body["mlDataParam"] = request.MlDataParam
	}

	if !tea.BoolValue(util.IsUnset(request.RawLog)) {
		body["rawLog"] = request.RawLog
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("PutAnnotationData"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/ml/annotationdataset/" + tea.StringValue(datasetId) + "/annotationdata"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &PutAnnotationDataResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds data to a dataset for storage.
//
// @param request - PutAnnotationDataRequest
//
// @return PutAnnotationDataResponse
func (client *Client) PutAnnotationData(datasetId *string, request *PutAnnotationDataRequest) (_result *PutAnnotationDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &PutAnnotationDataResponse{}
	_body, _err := client.PutAnnotationDataWithOptions(datasetId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a project policy.
//
// Description:
//
// ### Usage notes
//
// 	- Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
//
// 	- Alibaba Cloud Simple Log Service allows you to configure a project policy to authorize other users to access the specified Log Service resources.
//
//     	- You must configure a project policy based on policy syntax. Before you configure a project policy, you must be familiar with the Action, Resource, and Condition parameters. For more information, see [RAM](https://help.aliyun.com/document_detail/128139.html).
//
//     	- If you set the Principal element to an asterisk (\\*) and do not configure the Condition element when you configure a project policy, the policy applies to all users except for the project owner. If you set the Principal element to an asterisk (\\*) and configure the Condition element when you configure a project policy, the policy applies to all users including the project owner.
//
//     	- You can configure multiple project policies for a project. The total size of the policies cannot exceed 16 KB.
//
// @param request - PutProjectPolicyRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PutProjectPolicyResponse
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
		BodyType:    tea.String("none"),
	}
	_result = &PutProjectPolicyResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a project policy.
//
// Description:
//
// ### Usage notes
//
// 	- Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
//
// 	- Alibaba Cloud Simple Log Service allows you to configure a project policy to authorize other users to access the specified Log Service resources.
//
//     	- You must configure a project policy based on policy syntax. Before you configure a project policy, you must be familiar with the Action, Resource, and Condition parameters. For more information, see [RAM](https://help.aliyun.com/document_detail/128139.html).
//
//     	- If you set the Principal element to an asterisk (\\*) and do not configure the Condition element when you configure a project policy, the policy applies to all users except for the project owner. If you set the Principal element to an asterisk (\\*) and configure the Condition element when you configure a project policy, the policy applies to all users including the project owner.
//
//     	- You can configure multiple project policies for a project. The total size of the policies cannot exceed 16 KB.
//
// @param request - PutProjectPolicyRequest
//
// @return PutProjectPolicyResponse
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

// Summary:
//
// 设置project传输加速状态
//
// @param request - PutProjectTransferAccelerationRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PutProjectTransferAccelerationResponse
func (client *Client) PutProjectTransferAccelerationWithOptions(project *string, request *PutProjectTransferAccelerationRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *PutProjectTransferAccelerationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Enabled)) {
		body["enabled"] = request.Enabled
	}

	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("PutProjectTransferAcceleration"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/transferacceleration"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &PutProjectTransferAccelerationResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 设置project传输加速状态
//
// @param request - PutProjectTransferAccelerationRequest
//
// @return PutProjectTransferAccelerationResponse
func (client *Client) PutProjectTransferAcceleration(project *string, request *PutProjectTransferAccelerationRequest) (_result *PutProjectTransferAccelerationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &PutProjectTransferAccelerationResponse{}
	_body, _err := client.PutProjectTransferAccelerationWithOptions(project, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Sends multiple logs to Simple Log Service in one request.
//
// Description:
//
// ### [](#)Usage notes
//
// 	- The information that is required to query logs is obtained. The information includes the name of the project to which the logs belong, the region of the project, and the name of the Logstore to which the logs belong. For more information, see [Manage a project](https://help.aliyun.com/document_detail/48984.html) and [Manage a Logstore](https://help.aliyun.com/document_detail/48990.html).
//
// 	- You can call this operation to collect logs from web pages or clients.
//
// 	- If you use web tracking to collect logs and you do not call this operation, you can send only one log to Simple Log Service in a request. For more information, see [Use web tracking to collect logs](https://help.aliyun.com/document_detail/31752.html).
//
// 	- If you want to collect a large amount of log data, you can call this operation to send multiple logs to Simple Log Service in one request.
//
// 	- Before you can call this operation to send logs to a Logstore, you must enable web tracking for the Logstore. For more information, see [Use web tracking to collect logs](https://help.aliyun.com/document_detail/31752.html).
//
// 	- You cannot call this operation to send the logs of multiple topics to Simple Log Service at a time.
//
// 	- If you call this operation, anonymous users from the Internet are granted the write permissions on the Logstore. This may generate dirty data because AccessKey pair-based authentication is not performed.
//
// @param request - PutWebtrackingRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PutWebtrackingResponse
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
		AuthType:    tea.String("Anonymous"),
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

// Summary:
//
// Sends multiple logs to Simple Log Service in one request.
//
// Description:
//
// ### [](#)Usage notes
//
// 	- The information that is required to query logs is obtained. The information includes the name of the project to which the logs belong, the region of the project, and the name of the Logstore to which the logs belong. For more information, see [Manage a project](https://help.aliyun.com/document_detail/48984.html) and [Manage a Logstore](https://help.aliyun.com/document_detail/48990.html).
//
// 	- You can call this operation to collect logs from web pages or clients.
//
// 	- If you use web tracking to collect logs and you do not call this operation, you can send only one log to Simple Log Service in a request. For more information, see [Use web tracking to collect logs](https://help.aliyun.com/document_detail/31752.html).
//
// 	- If you want to collect a large amount of log data, you can call this operation to send multiple logs to Simple Log Service in one request.
//
// 	- Before you can call this operation to send logs to a Logstore, you must enable web tracking for the Logstore. For more information, see [Use web tracking to collect logs](https://help.aliyun.com/document_detail/31752.html).
//
// 	- You cannot call this operation to send the logs of multiple topics to Simple Log Service at a time.
//
// 	- If you call this operation, anonymous users from the Internet are granted the write permissions on the Logstore. This may generate dirty data because AccessKey pair-based authentication is not performed.
//
// @param request - PutWebtrackingRequest
//
// @return PutWebtrackingResponse
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

// Deprecated: OpenAPI QueryMLServiceResults is deprecated
//
// Summary:
//
// queryMLServiceResults
//
// @param request - QueryMLServiceResultsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryMLServiceResultsResponse
// Deprecated
func (client *Client) QueryMLServiceResultsWithOptions(serviceName *string, request *QueryMLServiceResultsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *QueryMLServiceResultsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AllowBuiltin)) {
		query["allowBuiltin"] = request.AllowBuiltin
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryMLServiceResults"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/ml/service/" + tea.StringValue(serviceName) + "/analysis"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryMLServiceResultsResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI QueryMLServiceResults is deprecated
//
// Summary:
//
// queryMLServiceResults
//
// @param request - QueryMLServiceResultsRequest
//
// @return QueryMLServiceResultsResponse
// Deprecated
func (client *Client) QueryMLServiceResults(serviceName *string, request *QueryMLServiceResultsRequest) (_result *QueryMLServiceResultsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &QueryMLServiceResultsResponse{}
	_body, _err := client.QueryMLServiceResultsWithOptions(serviceName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 刷新token
//
// @param request - RefreshTokenRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RefreshTokenResponse
func (client *Client) RefreshTokenWithOptions(request *RefreshTokenRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *RefreshTokenResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessTokenExpirationTime)) {
		query["accessTokenExpirationTime"] = request.AccessTokenExpirationTime
	}

	if !tea.BoolValue(util.IsUnset(request.Ticket)) {
		query["ticket"] = request.Ticket
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RefreshToken"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/token/refresh"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &RefreshTokenResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 刷新token
//
// @param request - RefreshTokenRequest
//
// @return RefreshTokenResponse
func (client *Client) RefreshToken(request *RefreshTokenRequest) (_result *RefreshTokenResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RefreshTokenResponse{}
	_body, _err := client.RefreshTokenWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Removes a Logtail configuration from a machine group.
//
// Description:
//
// Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveConfigFromMachineGroupResponse
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

// Summary:
//
// Removes a Logtail configuration from a machine group.
//
// Description:
//
// Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
//
// @return RemoveConfigFromMachineGroupResponse
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

// Summary:
//
// Splits a shard in the readwrite state.
//
// Description:
//
// ### Usage notes
//
// 	- Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
//
// 	- Each shard has an MD5 hash range, and each range is a left-closed, right-open interval. The interval is in the `[BeginKey,EndKey)` format. A shard can be in the readwrite or readonly state. You can split a shard and merge shards. For more information, see [Shard](https://help.aliyun.com/document_detail/28976.html).
//
// @param request - SplitShardRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SplitShardResponse
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

// Summary:
//
// Splits a shard in the readwrite state.
//
// Description:
//
// ### Usage notes
//
// 	- Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
//
// 	- Each shard has an MD5 hash range, and each range is a left-closed, right-open interval. The interval is in the `[BeginKey,EndKey)` format. A shard can be in the readwrite or readonly state. You can split a shard and merge shards. For more information, see [Shard](https://help.aliyun.com/document_detail/28976.html).
//
// @param request - SplitShardRequest
//
// @return SplitShardResponse
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

// Summary:
//
// 启动数据加工任务
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartETLResponse
func (client *Client) StartETLWithOptions(project *string, etlName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *StartETLResponse, _err error) {
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("StartETL"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/etls/" + tea.StringValue(etlName) + "?action=START"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &StartETLResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 启动数据加工任务
//
// @return StartETLResponse
func (client *Client) StartETL(project *string, etlName *string) (_result *StartETLResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &StartETLResponse{}
	_body, _err := client.StartETLWithOptions(project, etlName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 启动OSS投递任务
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartOSSExportResponse
func (client *Client) StartOSSExportWithOptions(project *string, ossExportName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *StartOSSExportResponse, _err error) {
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("StartOSSExport"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/ossexports/" + tea.StringValue(ossExportName) + "?action=START"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &StartOSSExportResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 启动OSS投递任务
//
// @return StartOSSExportResponse
func (client *Client) StartOSSExport(project *string, ossExportName *string) (_result *StartOSSExportResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &StartOSSExportResponse{}
	_body, _err := client.StartOSSExportWithOptions(project, ossExportName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 启动OSSHDFS投递任务
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartOSSHDFSExportResponse
func (client *Client) StartOSSHDFSExportWithOptions(project *string, ossExportName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *StartOSSHDFSExportResponse, _err error) {
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("StartOSSHDFSExport"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/osshdfsexports/" + tea.StringValue(ossExportName) + "?action=START"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &StartOSSHDFSExportResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 启动OSSHDFS投递任务
//
// @return StartOSSHDFSExportResponse
func (client *Client) StartOSSHDFSExport(project *string, ossExportName *string) (_result *StartOSSHDFSExportResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &StartOSSHDFSExportResponse{}
	_body, _err := client.StartOSSHDFSExportWithOptions(project, ossExportName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 启动OSS导入任务
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartOSSIngestionResponse
func (client *Client) StartOSSIngestionWithOptions(project *string, ossIngestionName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *StartOSSIngestionResponse, _err error) {
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("StartOSSIngestion"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/ossingestions/" + tea.StringValue(ossIngestionName) + "?action=START"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &StartOSSIngestionResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 启动OSS导入任务
//
// @return StartOSSIngestionResponse
func (client *Client) StartOSSIngestion(project *string, ossIngestionName *string) (_result *StartOSSIngestionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &StartOSSIngestionResponse{}
	_body, _err := client.StartOSSIngestionWithOptions(project, ossIngestionName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 停止数据加工任务
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopETLResponse
func (client *Client) StopETLWithOptions(project *string, etlName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *StopETLResponse, _err error) {
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("StopETL"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/etls/" + tea.StringValue(etlName) + "?action=STOP"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &StopETLResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 停止数据加工任务
//
// @return StopETLResponse
func (client *Client) StopETL(project *string, etlName *string) (_result *StopETLResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &StopETLResponse{}
	_body, _err := client.StopETLWithOptions(project, etlName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 停止OSS投递任务
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopOSSExportResponse
func (client *Client) StopOSSExportWithOptions(project *string, ossExportName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *StopOSSExportResponse, _err error) {
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("StopOSSExport"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/ossexports/" + tea.StringValue(ossExportName) + "?action=STOP"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &StopOSSExportResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 停止OSS投递任务
//
// @return StopOSSExportResponse
func (client *Client) StopOSSExport(project *string, ossExportName *string) (_result *StopOSSExportResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &StopOSSExportResponse{}
	_body, _err := client.StopOSSExportWithOptions(project, ossExportName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 停止OSSHDFS投递任务
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopOSSHDFSExportResponse
func (client *Client) StopOSSHDFSExportWithOptions(project *string, ossExportName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *StopOSSHDFSExportResponse, _err error) {
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("StopOSSHDFSExport"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/osshdfsexports/" + tea.StringValue(ossExportName) + "?action=STOP"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &StopOSSHDFSExportResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 停止OSSHDFS投递任务
//
// @return StopOSSHDFSExportResponse
func (client *Client) StopOSSHDFSExport(project *string, ossExportName *string) (_result *StopOSSHDFSExportResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &StopOSSHDFSExportResponse{}
	_body, _err := client.StopOSSHDFSExportWithOptions(project, ossExportName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 停止OSS导入任务
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopOSSIngestionResponse
func (client *Client) StopOSSIngestionWithOptions(project *string, ossIngestionName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *StopOSSIngestionResponse, _err error) {
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("StopOSSIngestion"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/ossingestions/" + tea.StringValue(ossIngestionName) + "?action=STOP"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &StopOSSIngestionResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 停止OSS导入任务
//
// @return StopOSSIngestionResponse
func (client *Client) StopOSSIngestion(project *string, ossIngestionName *string) (_result *StopOSSIngestionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &StopOSSIngestionResponse{}
	_body, _err := client.StopOSSIngestionWithOptions(project, ossIngestionName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates and adds one or more tags to a specified resource. You can add tags only to projects.
//
// Description:
//
// ### Usage notes
//
// Host consists of a project name and a Log Service endpoint. You must specify a project in Host.
//
// @param request - TagResourcesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TagResourcesResponse
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

// Summary:
//
// Creates and adds one or more tags to a specified resource. You can add tags only to projects.
//
// Description:
//
// ### Usage notes
//
// Host consists of a project name and a Log Service endpoint. You must specify a project in Host.
//
// @param request - TagResourcesRequest
//
// @return TagResourcesResponse
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

// Summary:
//
// Detaches one or more tags from a resource. You can detach tags only from Simple Log Service projects. You can detach multiple or all tags from a Simple Log Service project at a time.
//
// Description:
//
// ### Usage notes
//
// Host consists of a project name and a Log Service endpoint. You must specify a project in Host.
//
// @param request - UntagResourcesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UntagResourcesResponse
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
		BodyType:    tea.String("none"),
	}
	_result = &UntagResourcesResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Detaches one or more tags from a resource. You can detach tags only from Simple Log Service projects. You can detach multiple or all tags from a Simple Log Service project at a time.
//
// Description:
//
// ### Usage notes
//
// Host consists of a project name and a Log Service endpoint. You must specify a project in Host.
//
// @param request - UntagResourcesRequest
//
// @return UntagResourcesResponse
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

// Summary:
//
// 更新告警
//
// @param request - UpdateAlertRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAlertResponse
func (client *Client) UpdateAlertWithOptions(project *string, alertName *string, request *UpdateAlertRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateAlertResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Configuration)) {
		body["configuration"] = request.Configuration
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.DisplayName)) {
		body["displayName"] = request.DisplayName
	}

	if !tea.BoolValue(util.IsUnset(request.Schedule)) {
		body["schedule"] = request.Schedule
	}

	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateAlert"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/alerts/" + tea.StringValue(alertName)),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &UpdateAlertResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新告警
//
// @param request - UpdateAlertRequest
//
// @return UpdateAlertResponse
func (client *Client) UpdateAlert(project *string, alertName *string, request *UpdateAlertRequest) (_result *UpdateAlertResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateAlertResponse{}
	_body, _err := client.UpdateAlertWithOptions(project, alertName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates a dataset.
//
// @param request - UpdateAnnotationDataSetRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAnnotationDataSetResponse
func (client *Client) UpdateAnnotationDataSetWithOptions(datasetId *string, request *UpdateAnnotationDataSetRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateAnnotationDataSetResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateAnnotationDataSet"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/ml/annotationdataset/" + tea.StringValue(datasetId)),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &UpdateAnnotationDataSetResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a dataset.
//
// @param request - UpdateAnnotationDataSetRequest
//
// @return UpdateAnnotationDataSetResponse
func (client *Client) UpdateAnnotationDataSet(datasetId *string, request *UpdateAnnotationDataSetRequest) (_result *UpdateAnnotationDataSetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateAnnotationDataSetResponse{}
	_body, _err := client.UpdateAnnotationDataSetWithOptions(datasetId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates a tag table.
//
// Description:
//
// You can update only the names of the tags in a tag set.
//
// @param request - UpdateAnnotationLabelRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAnnotationLabelResponse
func (client *Client) UpdateAnnotationLabelWithOptions(request *UpdateAnnotationLabelRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateAnnotationLabelResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateAnnotationLabel"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/ml/annotationlabel"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &UpdateAnnotationLabelResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a tag table.
//
// Description:
//
// You can update only the names of the tags in a tag set.
//
// @param request - UpdateAnnotationLabelRequest
//
// @return UpdateAnnotationLabelResponse
func (client *Client) UpdateAnnotationLabel(request *UpdateAnnotationLabelRequest) (_result *UpdateAnnotationLabelResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateAnnotationLabelResponse{}
	_body, _err := client.UpdateAnnotationLabelWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates a Logtail configuration.
//
// Description:
//
// ### [](#)Usage notes
//
// 	- Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
//
// 	- After you update a Logtail configuration that is applied to a machine group, the new configuration immediately takes effect.
//
// 	- An AccessKey pair is created and obtained. For more information, see [AccessKey pair](https://help.aliyun.com/document_detail/29009.html).
//
// The AccessKey pair of an Alibaba Cloud account has permissions on all API operations. Using these credentials to perform operations in Simple Log Service is a high-risk operation. We recommend that you use a RAM user to call API operations or perform routine O\\&M. To create a RAM user, log on to the RAM console. Make sure that the RAM user has the management permissions on Simple Log Service resources. For more information, see [Create a RAM user and authorize the RAM user to access Simple Log Service](https://help.aliyun.com/document_detail/47664.html).
//
// 	- The information that is required to query logs is obtained. The information includes the name of the project to which the logs belong and the region of the project. For more information, see [Manage a project](https://help.aliyun.com/document_detail/48984.html).
//
// 	- The Logtail configuration is planned out. For more information, see [Logtail configurations](https://help.aliyun.com/document_detail/29058.html).
//
// @param request - UpdateConfigRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateConfigResponse
func (client *Client) UpdateConfigWithOptions(project *string, configName *string, request *UpdateConfigRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateConfigResponse, _err error) {
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
		Action:      tea.String("UpdateConfig"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/configs/" + tea.StringValue(configName)),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &UpdateConfigResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a Logtail configuration.
//
// Description:
//
// ### [](#)Usage notes
//
// 	- Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
//
// 	- After you update a Logtail configuration that is applied to a machine group, the new configuration immediately takes effect.
//
// 	- An AccessKey pair is created and obtained. For more information, see [AccessKey pair](https://help.aliyun.com/document_detail/29009.html).
//
// The AccessKey pair of an Alibaba Cloud account has permissions on all API operations. Using these credentials to perform operations in Simple Log Service is a high-risk operation. We recommend that you use a RAM user to call API operations or perform routine O\\&M. To create a RAM user, log on to the RAM console. Make sure that the RAM user has the management permissions on Simple Log Service resources. For more information, see [Create a RAM user and authorize the RAM user to access Simple Log Service](https://help.aliyun.com/document_detail/47664.html).
//
// 	- The information that is required to query logs is obtained. The information includes the name of the project to which the logs belong and the region of the project. For more information, see [Manage a project](https://help.aliyun.com/document_detail/48984.html).
//
// 	- The Logtail configuration is planned out. For more information, see [Logtail configurations](https://help.aliyun.com/document_detail/29058.html).
//
// @param request - UpdateConfigRequest
//
// @return UpdateConfigResponse
func (client *Client) UpdateConfig(project *string, configName *string, request *UpdateConfigRequest) (_result *UpdateConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateConfigResponse{}
	_body, _err := client.UpdateConfigWithOptions(project, configName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the attributes of a consumer group.
//
// Description:
//
// ### Usage notes
//
// Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
//
// @param request - UpdateConsumerGroupRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateConsumerGroupResponse
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

// Summary:
//
// Updates the attributes of a consumer group.
//
// Description:
//
// ### Usage notes
//
// Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
//
// @param request - UpdateConsumerGroupRequest
//
// @return UpdateConsumerGroupResponse
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

// Summary:
//
// Updates a dashboard.
//
// Description:
//
// ### [](#)Usage notes
//
// Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
//
// @param request - UpdateDashboardRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDashboardResponse
func (client *Client) UpdateDashboardWithOptions(project *string, dashboardName *string, request *UpdateDashboardRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateDashboardResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Attribute)) {
		body["attribute"] = request.Attribute
	}

	if !tea.BoolValue(util.IsUnset(request.Charts)) {
		body["charts"] = request.Charts
	}

	if !tea.BoolValue(util.IsUnset(request.DashboardName)) {
		body["dashboardName"] = request.DashboardName
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.DisplayName)) {
		body["displayName"] = request.DisplayName
	}

	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateDashboard"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/dashboards/" + tea.StringValue(dashboardName)),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &UpdateDashboardResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a dashboard.
//
// Description:
//
// ### [](#)Usage notes
//
// Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
//
// @param request - UpdateDashboardRequest
//
// @return UpdateDashboardResponse
func (client *Client) UpdateDashboard(project *string, dashboardName *string, request *UpdateDashboardRequest) (_result *UpdateDashboardResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateDashboardResponse{}
	_body, _err := client.UpdateDashboardWithOptions(project, dashboardName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新数据加工任务
//
// @param request - UpdateETLRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateETLResponse
func (client *Client) UpdateETLWithOptions(project *string, etlName *string, request *UpdateETLRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateETLResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Configuration)) {
		body["configuration"] = request.Configuration
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.DisplayName)) {
		body["displayName"] = request.DisplayName
	}

	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateETL"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/etls/" + tea.StringValue(etlName)),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &UpdateETLResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新数据加工任务
//
// @param request - UpdateETLRequest
//
// @return UpdateETLResponse
func (client *Client) UpdateETL(project *string, etlName *string, request *UpdateETLRequest) (_result *UpdateETLResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateETLResponse{}
	_body, _err := client.UpdateETLWithOptions(project, etlName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the indexes of a Logstore.
//
// Description:
//
// ### Usage notes
//
// Host consists of a project name and a Log Service endpoint. You must specify a project in Host.
//
// @param request - UpdateIndexRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateIndexResponse
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

// Summary:
//
// Updates the indexes of a Logstore.
//
// Description:
//
// ### Usage notes
//
// Host consists of a project name and a Log Service endpoint. You must specify a project in Host.
//
// @param request - UpdateIndexRequest
//
// @return UpdateIndexResponse
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

// Summary:
//
// Updates the attributes of a Logstore.
//
// Description:
//
// ### Usage notes
//
// 	- Host consists of a project name and a Log Service endpoint. You must specify a project in Host.
//
// 	- You can call the UpdateLogStore operation to change only the time-to-live (TTL) attribute.
//
// @param request - UpdateLogStoreRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateLogStoreResponse
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

	if !tea.BoolValue(util.IsUnset(request.InfrequentAccessTTL)) {
		body["infrequentAccessTTL"] = request.InfrequentAccessTTL
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

// Summary:
//
// Updates the attributes of a Logstore.
//
// Description:
//
// ### Usage notes
//
// 	- Host consists of a project name and a Log Service endpoint. You must specify a project in Host.
//
// 	- You can call the UpdateLogStore operation to change only the time-to-live (TTL) attribute.
//
// @param request - UpdateLogStoreRequest
//
// @return UpdateLogStoreResponse
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

// Summary:
//
// 更新LogStore计量模式
//
// @param request - UpdateLogStoreMeteringModeRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateLogStoreMeteringModeResponse
func (client *Client) UpdateLogStoreMeteringModeWithOptions(project *string, logstore *string, request *UpdateLogStoreMeteringModeRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateLogStoreMeteringModeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MeteringMode)) {
		body["meteringMode"] = request.MeteringMode
	}

	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateLogStoreMeteringMode"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/logstores/" + tea.StringValue(logstore) + "/meteringmode"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &UpdateLogStoreMeteringModeResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新LogStore计量模式
//
// @param request - UpdateLogStoreMeteringModeRequest
//
// @return UpdateLogStoreMeteringModeResponse
func (client *Client) UpdateLogStoreMeteringMode(project *string, logstore *string, request *UpdateLogStoreMeteringModeRequest) (_result *UpdateLogStoreMeteringModeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateLogStoreMeteringModeResponse{}
	_body, _err := client.UpdateLogStoreMeteringModeWithOptions(project, logstore, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the service log configurations of a project.
//
// Description:
//
// Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
//
// @param request - UpdateLoggingRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateLoggingResponse
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

// Summary:
//
// Updates the service log configurations of a project.
//
// Description:
//
// Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
//
// @param request - UpdateLoggingRequest
//
// @return UpdateLoggingResponse
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

// Summary:
//
// Updates a Logtail pipeline configuration.
//
// Description:
//
// The UK (London) region is supported. Supported regions are constantly updated.
//
// @param request - UpdateLogtailPipelineConfigRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateLogtailPipelineConfigResponse
func (client *Client) UpdateLogtailPipelineConfigWithOptions(project *string, configName *string, request *UpdateLogtailPipelineConfigRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateLogtailPipelineConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Aggregators)) {
		body["aggregators"] = request.Aggregators
	}

	if !tea.BoolValue(util.IsUnset(request.ConfigName)) {
		body["configName"] = request.ConfigName
	}

	if !tea.BoolValue(util.IsUnset(request.Flushers)) {
		body["flushers"] = request.Flushers
	}

	if !tea.BoolValue(util.IsUnset(request.Global)) {
		body["global"] = request.Global
	}

	if !tea.BoolValue(util.IsUnset(request.Inputs)) {
		body["inputs"] = request.Inputs
	}

	if !tea.BoolValue(util.IsUnset(request.LogSample)) {
		body["logSample"] = request.LogSample
	}

	if !tea.BoolValue(util.IsUnset(request.Processors)) {
		body["processors"] = request.Processors
	}

	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateLogtailPipelineConfig"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/pipelineconfigs/" + tea.StringValue(configName)),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &UpdateLogtailPipelineConfigResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a Logtail pipeline configuration.
//
// Description:
//
// The UK (London) region is supported. Supported regions are constantly updated.
//
// @param request - UpdateLogtailPipelineConfigRequest
//
// @return UpdateLogtailPipelineConfigResponse
func (client *Client) UpdateLogtailPipelineConfig(project *string, configName *string, request *UpdateLogtailPipelineConfigRequest) (_result *UpdateLogtailPipelineConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateLogtailPipelineConfigResponse{}
	_body, _err := client.UpdateLogtailPipelineConfigWithOptions(project, configName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the configuration of a machine group.
//
// Description:
//
// Host consists of a project name and a Log Service endpoint. You must specify a project in Host.
//
// @param request - UpdateMachineGroupRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateMachineGroupResponse
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

// Summary:
//
// Modifies the configuration of a machine group.
//
// Description:
//
// Host consists of a project name and a Log Service endpoint. You must specify a project in Host.
//
// @param request - UpdateMachineGroupRequest
//
// @return UpdateMachineGroupResponse
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

// Summary:
//
// Modifies the machines in a machine group. You can add machine to or remove machines from the machine group.
//
// Description:
//
// Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
//
// @param request - UpdateMachineGroupMachineRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateMachineGroupMachineResponse
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

// Summary:
//
// Modifies the machines in a machine group. You can add machine to or remove machines from the machine group.
//
// Description:
//
// Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
//
// @param request - UpdateMachineGroupMachineRequest
//
// @return UpdateMachineGroupMachineResponse
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

// Summary:
//
// 更新 MetricStore 计量模式
//
// @param request - UpdateMetricStoreMeteringModeRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateMetricStoreMeteringModeResponse
func (client *Client) UpdateMetricStoreMeteringModeWithOptions(project *string, metricStore *string, request *UpdateMetricStoreMeteringModeRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateMetricStoreMeteringModeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MeteringMode)) {
		body["meteringMode"] = request.MeteringMode
	}

	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateMetricStoreMeteringMode"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/metricstores/" + tea.StringValue(metricStore) + "/meteringmode"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &UpdateMetricStoreMeteringModeResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新 MetricStore 计量模式
//
// @param request - UpdateMetricStoreMeteringModeRequest
//
// @return UpdateMetricStoreMeteringModeResponse
func (client *Client) UpdateMetricStoreMeteringMode(project *string, metricStore *string, request *UpdateMetricStoreMeteringModeRequest) (_result *UpdateMetricStoreMeteringModeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateMetricStoreMeteringModeResponse{}
	_body, _err := client.UpdateMetricStoreMeteringModeWithOptions(project, metricStore, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新OSS投递任务
//
// @param request - UpdateOSSExportRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateOSSExportResponse
func (client *Client) UpdateOSSExportWithOptions(project *string, ossExportName *string, request *UpdateOSSExportRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateOSSExportResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Configuration)) {
		body["configuration"] = request.Configuration
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.DisplayName)) {
		body["displayName"] = request.DisplayName
	}

	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateOSSExport"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/ossexports/" + tea.StringValue(ossExportName)),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &UpdateOSSExportResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新OSS投递任务
//
// @param request - UpdateOSSExportRequest
//
// @return UpdateOSSExportResponse
func (client *Client) UpdateOSSExport(project *string, ossExportName *string, request *UpdateOSSExportRequest) (_result *UpdateOSSExportResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateOSSExportResponse{}
	_body, _err := client.UpdateOSSExportWithOptions(project, ossExportName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新OSSHDFS投递任务
//
// @param request - UpdateOSSHDFSExportRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateOSSHDFSExportResponse
func (client *Client) UpdateOSSHDFSExportWithOptions(project *string, ossExportName *string, request *UpdateOSSHDFSExportRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateOSSHDFSExportResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Configuration)) {
		body["configuration"] = request.Configuration
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.DisplayName)) {
		body["displayName"] = request.DisplayName
	}

	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateOSSHDFSExport"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/osshdfsexports/" + tea.StringValue(ossExportName)),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &UpdateOSSHDFSExportResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新OSSHDFS投递任务
//
// @param request - UpdateOSSHDFSExportRequest
//
// @return UpdateOSSHDFSExportResponse
func (client *Client) UpdateOSSHDFSExport(project *string, ossExportName *string, request *UpdateOSSHDFSExportRequest) (_result *UpdateOSSHDFSExportResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateOSSHDFSExportResponse{}
	_body, _err := client.UpdateOSSHDFSExportWithOptions(project, ossExportName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新oss导入任务
//
// @param request - UpdateOSSIngestionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateOSSIngestionResponse
func (client *Client) UpdateOSSIngestionWithOptions(project *string, ossIngestionName *string, request *UpdateOSSIngestionRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateOSSIngestionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Configuration)) {
		body["configuration"] = request.Configuration
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.DisplayName)) {
		body["displayName"] = request.DisplayName
	}

	if !tea.BoolValue(util.IsUnset(request.Schedule)) {
		body["schedule"] = request.Schedule
	}

	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateOSSIngestion"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/ossingestions/" + tea.StringValue(ossIngestionName)),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &UpdateOSSIngestionResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新oss导入任务
//
// @param request - UpdateOSSIngestionRequest
//
// @return UpdateOSSIngestionResponse
func (client *Client) UpdateOSSIngestion(project *string, ossIngestionName *string, request *UpdateOSSIngestionRequest) (_result *UpdateOSSIngestionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateOSSIngestionResponse{}
	_body, _err := client.UpdateOSSIngestionWithOptions(project, ossIngestionName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates an Object Storage Service (OSS) external store.
//
// Description:
//
// ### [](#)Usage notes
//
// Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
//
// @param request - UpdateOssExternalStoreRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateOssExternalStoreResponse
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

// Summary:
//
// Updates an Object Storage Service (OSS) external store.
//
// Description:
//
// ### [](#)Usage notes
//
// Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
//
// @param request - UpdateOssExternalStoreRequest
//
// @return UpdateOssExternalStoreResponse
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

// Summary:
//
// Updates a project.
//
// Description:
//
// ### Usage notes
//
// Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
//
// @param request - UpdateProjectRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateProjectResponse
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

// Summary:
//
// Updates a project.
//
// Description:
//
// ### Usage notes
//
// Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
//
// @param request - UpdateProjectRequest
//
// @return UpdateProjectResponse
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

// Summary:
//
// Updates an ApsaraDB RDS external store.
//
// Description:
//
// Host consists of a project name and a Log Service endpoint. You must specify a project in Host.
//
// @param request - UpdateRdsExternalStoreRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateRdsExternalStoreResponse
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

// Summary:
//
// Updates an ApsaraDB RDS external store.
//
// Description:
//
// Host consists of a project name and a Log Service endpoint. You must specify a project in Host.
//
// @param request - UpdateRdsExternalStoreRequest
//
// @return UpdateRdsExternalStoreResponse
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

// Summary:
//
// Updates a saved search.
//
// @param request - UpdateSavedSearchRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateSavedSearchResponse
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

// Summary:
//
// Updates a saved search.
//
// @param request - UpdateSavedSearchRequest
//
// @return UpdateSavedSearchResponse
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

// Summary:
//
// 更新定时SQL任务
//
// @param request - UpdateScheduledSQLRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateScheduledSQLResponse
func (client *Client) UpdateScheduledSQLWithOptions(project *string, scheduledSQLName *string, request *UpdateScheduledSQLRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateScheduledSQLResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Configuration)) {
		body["configuration"] = request.Configuration
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.DisplayName)) {
		body["displayName"] = request.DisplayName
	}

	if !tea.BoolValue(util.IsUnset(request.Schedule)) {
		body["schedule"] = request.Schedule
	}

	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateScheduledSQL"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/scheduledsqls/" + tea.StringValue(scheduledSQLName)),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &UpdateScheduledSQLResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新定时SQL任务
//
// @param request - UpdateScheduledSQLRequest
//
// @return UpdateScheduledSQLResponse
func (client *Client) UpdateScheduledSQL(project *string, scheduledSQLName *string, request *UpdateScheduledSQLRequest) (_result *UpdateScheduledSQLResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateScheduledSQLResponse{}
	_body, _err := client.UpdateScheduledSQLWithOptions(project, scheduledSQLName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新独享sql实例
//
// @param request - UpdateSqlInstanceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateSqlInstanceResponse
func (client *Client) UpdateSqlInstanceWithOptions(project *string, request *UpdateSqlInstanceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateSqlInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Cu)) {
		body["cu"] = request.Cu
	}

	if !tea.BoolValue(util.IsUnset(request.UseAsDefault)) {
		body["useAsDefault"] = request.UseAsDefault
	}

	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateSqlInstance"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/sqlinstance"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &UpdateSqlInstanceResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新独享sql实例
//
// @param request - UpdateSqlInstanceRequest
//
// @return UpdateSqlInstanceResponse
func (client *Client) UpdateSqlInstance(project *string, request *UpdateSqlInstanceRequest) (_result *UpdateSqlInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateSqlInstanceResponse{}
	_body, _err := client.UpdateSqlInstanceWithOptions(project, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新StoreView
//
// @param request - UpdateStoreViewRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateStoreViewResponse
func (client *Client) UpdateStoreViewWithOptions(project *string, name *string, request *UpdateStoreViewRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateStoreViewResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.StoreType)) {
		body["storeType"] = request.StoreType
	}

	if !tea.BoolValue(util.IsUnset(request.Stores)) {
		body["stores"] = request.Stores
	}

	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateStoreView"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/storeviews/" + tea.StringValue(name)),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &UpdateStoreViewResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新StoreView
//
// @param request - UpdateStoreViewRequest
//
// @return UpdateStoreViewResponse
func (client *Client) UpdateStoreView(project *string, name *string, request *UpdateStoreViewRequest) (_result *UpdateStoreViewResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateStoreViewResponse{}
	_body, _err := client.UpdateStoreViewWithOptions(project, name, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 调用UpsertCollectionPolicy接口更新采集策略的属性信息
//
// @param request - UpsertCollectionPolicyRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpsertCollectionPolicyResponse
func (client *Client) UpsertCollectionPolicyWithOptions(request *UpsertCollectionPolicyRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpsertCollectionPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Attribute)) {
		body["attribute"] = request.Attribute
	}

	if !tea.BoolValue(util.IsUnset(request.CentralizeConfig)) {
		body["centralizeConfig"] = request.CentralizeConfig
	}

	if !tea.BoolValue(util.IsUnset(request.CentralizeEnabled)) {
		body["centralizeEnabled"] = request.CentralizeEnabled
	}

	if !tea.BoolValue(util.IsUnset(request.DataCode)) {
		body["dataCode"] = request.DataCode
	}

	if !tea.BoolValue(util.IsUnset(request.Enabled)) {
		body["enabled"] = request.Enabled
	}

	if !tea.BoolValue(util.IsUnset(request.PolicyConfig)) {
		body["policyConfig"] = request.PolicyConfig
	}

	if !tea.BoolValue(util.IsUnset(request.PolicyName)) {
		body["policyName"] = request.PolicyName
	}

	if !tea.BoolValue(util.IsUnset(request.ProductCode)) {
		body["productCode"] = request.ProductCode
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpsertCollectionPolicy"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/collectionpolicy"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpsertCollectionPolicyResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 调用UpsertCollectionPolicy接口更新采集策略的属性信息
//
// @param request - UpsertCollectionPolicyRequest
//
// @return UpsertCollectionPolicyResponse
func (client *Client) UpsertCollectionPolicy(request *UpsertCollectionPolicyRequest) (_result *UpsertCollectionPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpsertCollectionPolicyResponse{}
	_body, _err := client.UpsertCollectionPolicyWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
