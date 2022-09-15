// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	gatewayclient "github.com/alibabacloud-go/alibabacloud-gateway-sls/client"
	
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
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
	CreateTime     *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty"`
	LastModifyTime *string `json:"lastModifyTime,omitempty" xml:"lastModifyTime,omitempty"`
	Owner          *string `json:"owner,omitempty" xml:"owner,omitempty"`
	ProjectName    *string `json:"projectName,omitempty" xml:"projectName,omitempty"`
	Region         *string `json:"region,omitempty" xml:"region,omitempty"`
	Status         *string `json:"status,omitempty" xml:"status,omitempty"`
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

func (s *Project) SetStatus(v string) *Project {
	s.Status = &v
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

type BatchCreateEtlMetaRequest struct {
	EtlMetaList []*BatchCreateEtlMetaRequestEtlMetaList `json:"etlMetaList,omitempty" xml:"etlMetaList,omitempty" type:"Repeated"`
}

func (s BatchCreateEtlMetaRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchCreateEtlMetaRequest) GoString() string {
	return s.String()
}

func (s *BatchCreateEtlMetaRequest) SetEtlMetaList(v []*BatchCreateEtlMetaRequestEtlMetaList) *BatchCreateEtlMetaRequest {
	s.EtlMetaList = v
	return s
}

type BatchCreateEtlMetaRequestEtlMetaList struct {
	Enable       *bool                  `json:"enable,omitempty" xml:"enable,omitempty"`
	EtlMetaKey   *string                `json:"etlMetaKey,omitempty" xml:"etlMetaKey,omitempty"`
	EtlMetaName  *string                `json:"etlMetaName,omitempty" xml:"etlMetaName,omitempty"`
	EtlMetaTag   *string                `json:"etlMetaTag,omitempty" xml:"etlMetaTag,omitempty"`
	EtlMetaValue map[string]interface{} `json:"etlMetaValue,omitempty" xml:"etlMetaValue,omitempty"`
}

func (s BatchCreateEtlMetaRequestEtlMetaList) String() string {
	return tea.Prettify(s)
}

func (s BatchCreateEtlMetaRequestEtlMetaList) GoString() string {
	return s.String()
}

func (s *BatchCreateEtlMetaRequestEtlMetaList) SetEnable(v bool) *BatchCreateEtlMetaRequestEtlMetaList {
	s.Enable = &v
	return s
}

func (s *BatchCreateEtlMetaRequestEtlMetaList) SetEtlMetaKey(v string) *BatchCreateEtlMetaRequestEtlMetaList {
	s.EtlMetaKey = &v
	return s
}

func (s *BatchCreateEtlMetaRequestEtlMetaList) SetEtlMetaName(v string) *BatchCreateEtlMetaRequestEtlMetaList {
	s.EtlMetaName = &v
	return s
}

func (s *BatchCreateEtlMetaRequestEtlMetaList) SetEtlMetaTag(v string) *BatchCreateEtlMetaRequestEtlMetaList {
	s.EtlMetaTag = &v
	return s
}

func (s *BatchCreateEtlMetaRequestEtlMetaList) SetEtlMetaValue(v map[string]interface{}) *BatchCreateEtlMetaRequestEtlMetaList {
	s.EtlMetaValue = v
	return s
}

type BatchCreateEtlMetaResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
}

func (s BatchCreateEtlMetaResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchCreateEtlMetaResponse) GoString() string {
	return s.String()
}

func (s *BatchCreateEtlMetaResponse) SetHeaders(v map[string]*string) *BatchCreateEtlMetaResponse {
	s.Headers = v
	return s
}

func (s *BatchCreateEtlMetaResponse) SetStatusCode(v int32) *BatchCreateEtlMetaResponse {
	s.StatusCode = &v
	return s
}

type BatchModifyEtlMetaStatusRequest struct {
	EtlMetaKeyList []*string `json:"etlMetaKeyList,omitempty" xml:"etlMetaKeyList,omitempty" type:"Repeated"`
	EtlMetaName    *string   `json:"etlMetaName,omitempty" xml:"etlMetaName,omitempty"`
	EtlMetaTag     *string   `json:"etlMetaTag,omitempty" xml:"etlMetaTag,omitempty"`
	Range          *string   `json:"range,omitempty" xml:"range,omitempty"`
	Type           *string   `json:"type,omitempty" xml:"type,omitempty"`
}

func (s BatchModifyEtlMetaStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchModifyEtlMetaStatusRequest) GoString() string {
	return s.String()
}

func (s *BatchModifyEtlMetaStatusRequest) SetEtlMetaKeyList(v []*string) *BatchModifyEtlMetaStatusRequest {
	s.EtlMetaKeyList = v
	return s
}

func (s *BatchModifyEtlMetaStatusRequest) SetEtlMetaName(v string) *BatchModifyEtlMetaStatusRequest {
	s.EtlMetaName = &v
	return s
}

func (s *BatchModifyEtlMetaStatusRequest) SetEtlMetaTag(v string) *BatchModifyEtlMetaStatusRequest {
	s.EtlMetaTag = &v
	return s
}

func (s *BatchModifyEtlMetaStatusRequest) SetRange(v string) *BatchModifyEtlMetaStatusRequest {
	s.Range = &v
	return s
}

func (s *BatchModifyEtlMetaStatusRequest) SetType(v string) *BatchModifyEtlMetaStatusRequest {
	s.Type = &v
	return s
}

type BatchModifyEtlMetaStatusResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
}

func (s BatchModifyEtlMetaStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchModifyEtlMetaStatusResponse) GoString() string {
	return s.String()
}

func (s *BatchModifyEtlMetaStatusResponse) SetHeaders(v map[string]*string) *BatchModifyEtlMetaStatusResponse {
	s.Headers = v
	return s
}

func (s *BatchModifyEtlMetaStatusResponse) SetStatusCode(v int32) *BatchModifyEtlMetaStatusResponse {
	s.StatusCode = &v
	return s
}

type BatchUpdateEtlMetaRequest struct {
	EtlMetaList *BatchUpdateEtlMetaRequestEtlMetaList `json:"etlMetaList,omitempty" xml:"etlMetaList,omitempty" type:"Struct"`
}

func (s BatchUpdateEtlMetaRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchUpdateEtlMetaRequest) GoString() string {
	return s.String()
}

func (s *BatchUpdateEtlMetaRequest) SetEtlMetaList(v *BatchUpdateEtlMetaRequestEtlMetaList) *BatchUpdateEtlMetaRequest {
	s.EtlMetaList = v
	return s
}

type BatchUpdateEtlMetaRequestEtlMetaList struct {
	Enable       *bool                  `json:"enable,omitempty" xml:"enable,omitempty"`
	EtlMetaKey   *string                `json:"etlMetaKey,omitempty" xml:"etlMetaKey,omitempty"`
	EtlMetaName  *string                `json:"etlMetaName,omitempty" xml:"etlMetaName,omitempty"`
	EtlMetaTag   *string                `json:"etlMetaTag,omitempty" xml:"etlMetaTag,omitempty"`
	EtlMetaValue map[string]interface{} `json:"etlMetaValue,omitempty" xml:"etlMetaValue,omitempty"`
}

func (s BatchUpdateEtlMetaRequestEtlMetaList) String() string {
	return tea.Prettify(s)
}

func (s BatchUpdateEtlMetaRequestEtlMetaList) GoString() string {
	return s.String()
}

func (s *BatchUpdateEtlMetaRequestEtlMetaList) SetEnable(v bool) *BatchUpdateEtlMetaRequestEtlMetaList {
	s.Enable = &v
	return s
}

func (s *BatchUpdateEtlMetaRequestEtlMetaList) SetEtlMetaKey(v string) *BatchUpdateEtlMetaRequestEtlMetaList {
	s.EtlMetaKey = &v
	return s
}

func (s *BatchUpdateEtlMetaRequestEtlMetaList) SetEtlMetaName(v string) *BatchUpdateEtlMetaRequestEtlMetaList {
	s.EtlMetaName = &v
	return s
}

func (s *BatchUpdateEtlMetaRequestEtlMetaList) SetEtlMetaTag(v string) *BatchUpdateEtlMetaRequestEtlMetaList {
	s.EtlMetaTag = &v
	return s
}

func (s *BatchUpdateEtlMetaRequestEtlMetaList) SetEtlMetaValue(v map[string]interface{}) *BatchUpdateEtlMetaRequestEtlMetaList {
	s.EtlMetaValue = v
	return s
}

type BatchUpdateEtlMetaResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
}

func (s BatchUpdateEtlMetaResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchUpdateEtlMetaResponse) GoString() string {
	return s.String()
}

func (s *BatchUpdateEtlMetaResponse) SetHeaders(v map[string]*string) *BatchUpdateEtlMetaResponse {
	s.Headers = v
	return s
}

func (s *BatchUpdateEtlMetaResponse) SetStatusCode(v int32) *BatchUpdateEtlMetaResponse {
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

type CreateEtlJobRequest struct {
	Enable            *bool                              `json:"enable,omitempty" xml:"enable,omitempty"`
	EtlJobName        *string                            `json:"etlJobName,omitempty" xml:"etlJobName,omitempty"`
	FunctionConfig    *CreateEtlJobRequestFunctionConfig `json:"functionConfig,omitempty" xml:"functionConfig,omitempty" type:"Struct"`
	FunctionParameter map[string]interface{}             `json:"functionParameter,omitempty" xml:"functionParameter,omitempty"`
	LogConfig         *CreateEtlJobRequestLogConfig      `json:"logConfig,omitempty" xml:"logConfig,omitempty" type:"Struct"`
	SourceConfig      *CreateEtlJobRequestSourceConfig   `json:"sourceConfig,omitempty" xml:"sourceConfig,omitempty" type:"Struct"`
	TriggerConfig     *CreateEtlJobRequestTriggerConfig  `json:"triggerConfig,omitempty" xml:"triggerConfig,omitempty" type:"Struct"`
}

func (s CreateEtlJobRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateEtlJobRequest) GoString() string {
	return s.String()
}

func (s *CreateEtlJobRequest) SetEnable(v bool) *CreateEtlJobRequest {
	s.Enable = &v
	return s
}

func (s *CreateEtlJobRequest) SetEtlJobName(v string) *CreateEtlJobRequest {
	s.EtlJobName = &v
	return s
}

func (s *CreateEtlJobRequest) SetFunctionConfig(v *CreateEtlJobRequestFunctionConfig) *CreateEtlJobRequest {
	s.FunctionConfig = v
	return s
}

func (s *CreateEtlJobRequest) SetFunctionParameter(v map[string]interface{}) *CreateEtlJobRequest {
	s.FunctionParameter = v
	return s
}

func (s *CreateEtlJobRequest) SetLogConfig(v *CreateEtlJobRequestLogConfig) *CreateEtlJobRequest {
	s.LogConfig = v
	return s
}

func (s *CreateEtlJobRequest) SetSourceConfig(v *CreateEtlJobRequestSourceConfig) *CreateEtlJobRequest {
	s.SourceConfig = v
	return s
}

func (s *CreateEtlJobRequest) SetTriggerConfig(v *CreateEtlJobRequestTriggerConfig) *CreateEtlJobRequest {
	s.TriggerConfig = v
	return s
}

type CreateEtlJobRequestFunctionConfig struct {
	AccountId        *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
	Endpoint         *string `json:"endpoint,omitempty" xml:"endpoint,omitempty"`
	FunctionName     *string `json:"functionName,omitempty" xml:"functionName,omitempty"`
	FunctionProvider *string `json:"functionProvider,omitempty" xml:"functionProvider,omitempty"`
	RegionName       *string `json:"regionName,omitempty" xml:"regionName,omitempty"`
	RoleArn          *string `json:"roleArn,omitempty" xml:"roleArn,omitempty"`
	ServiceName      *string `json:"serviceName,omitempty" xml:"serviceName,omitempty"`
}

func (s CreateEtlJobRequestFunctionConfig) String() string {
	return tea.Prettify(s)
}

func (s CreateEtlJobRequestFunctionConfig) GoString() string {
	return s.String()
}

func (s *CreateEtlJobRequestFunctionConfig) SetAccountId(v string) *CreateEtlJobRequestFunctionConfig {
	s.AccountId = &v
	return s
}

func (s *CreateEtlJobRequestFunctionConfig) SetEndpoint(v string) *CreateEtlJobRequestFunctionConfig {
	s.Endpoint = &v
	return s
}

func (s *CreateEtlJobRequestFunctionConfig) SetFunctionName(v string) *CreateEtlJobRequestFunctionConfig {
	s.FunctionName = &v
	return s
}

func (s *CreateEtlJobRequestFunctionConfig) SetFunctionProvider(v string) *CreateEtlJobRequestFunctionConfig {
	s.FunctionProvider = &v
	return s
}

func (s *CreateEtlJobRequestFunctionConfig) SetRegionName(v string) *CreateEtlJobRequestFunctionConfig {
	s.RegionName = &v
	return s
}

func (s *CreateEtlJobRequestFunctionConfig) SetRoleArn(v string) *CreateEtlJobRequestFunctionConfig {
	s.RoleArn = &v
	return s
}

func (s *CreateEtlJobRequestFunctionConfig) SetServiceName(v string) *CreateEtlJobRequestFunctionConfig {
	s.ServiceName = &v
	return s
}

type CreateEtlJobRequestLogConfig struct {
	Endpoint     *string `json:"endpoint,omitempty" xml:"endpoint,omitempty"`
	LogstoreName *string `json:"logstoreName,omitempty" xml:"logstoreName,omitempty"`
	ProjectName  *string `json:"projectName,omitempty" xml:"projectName,omitempty"`
}

func (s CreateEtlJobRequestLogConfig) String() string {
	return tea.Prettify(s)
}

func (s CreateEtlJobRequestLogConfig) GoString() string {
	return s.String()
}

func (s *CreateEtlJobRequestLogConfig) SetEndpoint(v string) *CreateEtlJobRequestLogConfig {
	s.Endpoint = &v
	return s
}

func (s *CreateEtlJobRequestLogConfig) SetLogstoreName(v string) *CreateEtlJobRequestLogConfig {
	s.LogstoreName = &v
	return s
}

func (s *CreateEtlJobRequestLogConfig) SetProjectName(v string) *CreateEtlJobRequestLogConfig {
	s.ProjectName = &v
	return s
}

type CreateEtlJobRequestSourceConfig struct {
	LogstoreName *string `json:"logstoreName,omitempty" xml:"logstoreName,omitempty"`
}

func (s CreateEtlJobRequestSourceConfig) String() string {
	return tea.Prettify(s)
}

func (s CreateEtlJobRequestSourceConfig) GoString() string {
	return s.String()
}

func (s *CreateEtlJobRequestSourceConfig) SetLogstoreName(v string) *CreateEtlJobRequestSourceConfig {
	s.LogstoreName = &v
	return s
}

type CreateEtlJobRequestTriggerConfig struct {
	MaxRetryTime     *int32  `json:"maxRetryTime,omitempty" xml:"maxRetryTime,omitempty"`
	RoleArn          *string `json:"roleArn,omitempty" xml:"roleArn,omitempty"`
	StartingPosition *string `json:"startingPosition,omitempty" xml:"startingPosition,omitempty"`
	StartingUnixtime *int64  `json:"startingUnixtime,omitempty" xml:"startingUnixtime,omitempty"`
	TriggerInterval  *int32  `json:"triggerInterval,omitempty" xml:"triggerInterval,omitempty"`
}

func (s CreateEtlJobRequestTriggerConfig) String() string {
	return tea.Prettify(s)
}

func (s CreateEtlJobRequestTriggerConfig) GoString() string {
	return s.String()
}

func (s *CreateEtlJobRequestTriggerConfig) SetMaxRetryTime(v int32) *CreateEtlJobRequestTriggerConfig {
	s.MaxRetryTime = &v
	return s
}

func (s *CreateEtlJobRequestTriggerConfig) SetRoleArn(v string) *CreateEtlJobRequestTriggerConfig {
	s.RoleArn = &v
	return s
}

func (s *CreateEtlJobRequestTriggerConfig) SetStartingPosition(v string) *CreateEtlJobRequestTriggerConfig {
	s.StartingPosition = &v
	return s
}

func (s *CreateEtlJobRequestTriggerConfig) SetStartingUnixtime(v int64) *CreateEtlJobRequestTriggerConfig {
	s.StartingUnixtime = &v
	return s
}

func (s *CreateEtlJobRequestTriggerConfig) SetTriggerInterval(v int32) *CreateEtlJobRequestTriggerConfig {
	s.TriggerInterval = &v
	return s
}

type CreateEtlJobResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
}

func (s CreateEtlJobResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateEtlJobResponse) GoString() string {
	return s.String()
}

func (s *CreateEtlJobResponse) SetHeaders(v map[string]*string) *CreateEtlJobResponse {
	s.Headers = v
	return s
}

func (s *CreateEtlJobResponse) SetStatusCode(v int32) *CreateEtlJobResponse {
	s.StatusCode = &v
	return s
}

type CreateEtlMetaRequest struct {
	Enable       *bool                  `json:"enable,omitempty" xml:"enable,omitempty"`
	EtlMetaKey   *string                `json:"etlMetaKey,omitempty" xml:"etlMetaKey,omitempty"`
	EtlMetaName  *string                `json:"etlMetaName,omitempty" xml:"etlMetaName,omitempty"`
	EtlMetaTag   *string                `json:"etlMetaTag,omitempty" xml:"etlMetaTag,omitempty"`
	EtlMetaValue map[string]interface{} `json:"etlMetaValue,omitempty" xml:"etlMetaValue,omitempty"`
}

func (s CreateEtlMetaRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateEtlMetaRequest) GoString() string {
	return s.String()
}

func (s *CreateEtlMetaRequest) SetEnable(v bool) *CreateEtlMetaRequest {
	s.Enable = &v
	return s
}

func (s *CreateEtlMetaRequest) SetEtlMetaKey(v string) *CreateEtlMetaRequest {
	s.EtlMetaKey = &v
	return s
}

func (s *CreateEtlMetaRequest) SetEtlMetaName(v string) *CreateEtlMetaRequest {
	s.EtlMetaName = &v
	return s
}

func (s *CreateEtlMetaRequest) SetEtlMetaTag(v string) *CreateEtlMetaRequest {
	s.EtlMetaTag = &v
	return s
}

func (s *CreateEtlMetaRequest) SetEtlMetaValue(v map[string]interface{}) *CreateEtlMetaRequest {
	s.EtlMetaValue = v
	return s
}

type CreateEtlMetaResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
}

func (s CreateEtlMetaResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateEtlMetaResponse) GoString() string {
	return s.String()
}

func (s *CreateEtlMetaResponse) SetHeaders(v map[string]*string) *CreateEtlMetaResponse {
	s.Headers = v
	return s
}

func (s *CreateEtlMetaResponse) SetStatusCode(v int32) *CreateEtlMetaResponse {
	s.StatusCode = &v
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
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	ProjectName *string `json:"projectName,omitempty" xml:"projectName,omitempty"`
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

type DeleteEtlJobResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
}

func (s DeleteEtlJobResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteEtlJobResponse) GoString() string {
	return s.String()
}

func (s *DeleteEtlJobResponse) SetHeaders(v map[string]*string) *DeleteEtlJobResponse {
	s.Headers = v
	return s
}

func (s *DeleteEtlJobResponse) SetStatusCode(v int32) *DeleteEtlJobResponse {
	s.StatusCode = &v
	return s
}

type DeleteEtlMetaRequest struct {
	EtlMetaKey  *string `json:"etlMetaKey,omitempty" xml:"etlMetaKey,omitempty"`
	EtlMetaName *string `json:"etlMetaName,omitempty" xml:"etlMetaName,omitempty"`
	EtlMetaTag  *string `json:"etlMetaTag,omitempty" xml:"etlMetaTag,omitempty"`
}

func (s DeleteEtlMetaRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteEtlMetaRequest) GoString() string {
	return s.String()
}

func (s *DeleteEtlMetaRequest) SetEtlMetaKey(v string) *DeleteEtlMetaRequest {
	s.EtlMetaKey = &v
	return s
}

func (s *DeleteEtlMetaRequest) SetEtlMetaName(v string) *DeleteEtlMetaRequest {
	s.EtlMetaName = &v
	return s
}

func (s *DeleteEtlMetaRequest) SetEtlMetaTag(v string) *DeleteEtlMetaRequest {
	s.EtlMetaTag = &v
	return s
}

type DeleteEtlMetaResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
}

func (s DeleteEtlMetaResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteEtlMetaResponse) GoString() string {
	return s.String()
}

func (s *DeleteEtlMetaResponse) SetHeaders(v map[string]*string) *DeleteEtlMetaResponse {
	s.Headers = v
	return s
}

func (s *DeleteEtlMetaResponse) SetStatusCode(v int32) *DeleteEtlMetaResponse {
	s.StatusCode = &v
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
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
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

func (s *GetCursorRequest) SetType(v string) *GetCursorRequest {
	s.Type = &v
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
	Type   *string `json:"type,omitempty" xml:"type,omitempty"`
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

func (s *GetCursorTimeRequest) SetType(v string) *GetCursorTimeRequest {
	s.Type = &v
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

type GetEtlJobResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *EtlJob            `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetEtlJobResponse) SetBody(v *EtlJob) *GetEtlJobResponse {
	s.Body = v
	return s
}

type GetEtlMetaRequest struct {
	EtlMetaKey  *string `json:"etlMetaKey,omitempty" xml:"etlMetaKey,omitempty"`
	EtlMetaName *string `json:"etlMetaName,omitempty" xml:"etlMetaName,omitempty"`
	EtlMetaTag  *string `json:"etlMetaTag,omitempty" xml:"etlMetaTag,omitempty"`
}

func (s GetEtlMetaRequest) String() string {
	return tea.Prettify(s)
}

func (s GetEtlMetaRequest) GoString() string {
	return s.String()
}

func (s *GetEtlMetaRequest) SetEtlMetaKey(v string) *GetEtlMetaRequest {
	s.EtlMetaKey = &v
	return s
}

func (s *GetEtlMetaRequest) SetEtlMetaName(v string) *GetEtlMetaRequest {
	s.EtlMetaName = &v
	return s
}

func (s *GetEtlMetaRequest) SetEtlMetaTag(v string) *GetEtlMetaRequest {
	s.EtlMetaTag = &v
	return s
}

type GetEtlMetaResponseBody struct {
	EtlMetaList []*EtlMeta `json:"etlMetaList,omitempty" xml:"etlMetaList,omitempty" type:"Repeated"`
	Total       *int32     `json:"total,omitempty" xml:"total,omitempty"`
}

func (s GetEtlMetaResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetEtlMetaResponseBody) GoString() string {
	return s.String()
}

func (s *GetEtlMetaResponseBody) SetEtlMetaList(v []*EtlMeta) *GetEtlMetaResponseBody {
	s.EtlMetaList = v
	return s
}

func (s *GetEtlMetaResponseBody) SetTotal(v int32) *GetEtlMetaResponseBody {
	s.Total = &v
	return s
}

type GetEtlMetaResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetEtlMetaResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetEtlMetaResponse) String() string {
	return tea.Prettify(s)
}

func (s GetEtlMetaResponse) GoString() string {
	return s.String()
}

func (s *GetEtlMetaResponse) SetHeaders(v map[string]*string) *GetEtlMetaResponse {
	s.Headers = v
	return s
}

func (s *GetEtlMetaResponse) SetStatusCode(v int32) *GetEtlMetaResponse {
	s.StatusCode = &v
	return s
}

func (s *GetEtlMetaResponse) SetBody(v *GetEtlMetaResponseBody) *GetEtlMetaResponse {
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

type GetShipperConfigResponseBody struct {
	CreateTime          *int64                 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	LastModifyTime      *int64                 `json:"lastModifyTime,omitempty" xml:"lastModifyTime,omitempty"`
	ShipperName         *string                `json:"shipperName,omitempty" xml:"shipperName,omitempty"`
	TargetConfiguration map[string]interface{} `json:"targetConfiguration,omitempty" xml:"targetConfiguration,omitempty"`
	TargetType          *string                `json:"targetType,omitempty" xml:"targetType,omitempty"`
}

func (s GetShipperConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetShipperConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GetShipperConfigResponseBody) SetCreateTime(v int64) *GetShipperConfigResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetShipperConfigResponseBody) SetLastModifyTime(v int64) *GetShipperConfigResponseBody {
	s.LastModifyTime = &v
	return s
}

func (s *GetShipperConfigResponseBody) SetShipperName(v string) *GetShipperConfigResponseBody {
	s.ShipperName = &v
	return s
}

func (s *GetShipperConfigResponseBody) SetTargetConfiguration(v map[string]interface{}) *GetShipperConfigResponseBody {
	s.TargetConfiguration = v
	return s
}

func (s *GetShipperConfigResponseBody) SetTargetType(v string) *GetShipperConfigResponseBody {
	s.TargetType = &v
	return s
}

type GetShipperConfigResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetShipperConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetShipperConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s GetShipperConfigResponse) GoString() string {
	return s.String()
}

func (s *GetShipperConfigResponse) SetHeaders(v map[string]*string) *GetShipperConfigResponse {
	s.Headers = v
	return s
}

func (s *GetShipperConfigResponse) SetStatusCode(v int32) *GetShipperConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *GetShipperConfigResponse) SetBody(v *GetShipperConfigResponseBody) *GetShipperConfigResponse {
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

type ListEtlJobRequest struct {
	Offset *int32 `json:"offset,omitempty" xml:"offset,omitempty"`
	Size   *int32 `json:"size,omitempty" xml:"size,omitempty"`
}

func (s ListEtlJobRequest) String() string {
	return tea.Prettify(s)
}

func (s ListEtlJobRequest) GoString() string {
	return s.String()
}

func (s *ListEtlJobRequest) SetOffset(v int32) *ListEtlJobRequest {
	s.Offset = &v
	return s
}

func (s *ListEtlJobRequest) SetSize(v int32) *ListEtlJobRequest {
	s.Size = &v
	return s
}

type ListEtlJobResponseBody struct {
	Count          *int32    `json:"count,omitempty" xml:"count,omitempty"`
	EtlJobNameList []*string `json:"etlJobNameList,omitempty" xml:"etlJobNameList,omitempty" type:"Repeated"`
	Total          *int32    `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListEtlJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListEtlJobResponseBody) GoString() string {
	return s.String()
}

func (s *ListEtlJobResponseBody) SetCount(v int32) *ListEtlJobResponseBody {
	s.Count = &v
	return s
}

func (s *ListEtlJobResponseBody) SetEtlJobNameList(v []*string) *ListEtlJobResponseBody {
	s.EtlJobNameList = v
	return s
}

func (s *ListEtlJobResponseBody) SetTotal(v int32) *ListEtlJobResponseBody {
	s.Total = &v
	return s
}

type ListEtlJobResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListEtlJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListEtlJobResponse) String() string {
	return tea.Prettify(s)
}

func (s ListEtlJobResponse) GoString() string {
	return s.String()
}

func (s *ListEtlJobResponse) SetHeaders(v map[string]*string) *ListEtlJobResponse {
	s.Headers = v
	return s
}

func (s *ListEtlJobResponse) SetStatusCode(v int32) *ListEtlJobResponse {
	s.StatusCode = &v
	return s
}

func (s *ListEtlJobResponse) SetBody(v *ListEtlJobResponseBody) *ListEtlJobResponse {
	s.Body = v
	return s
}

type ListEtlMetaRequest struct {
	EtlMetaKey  *string `json:"etlMetaKey,omitempty" xml:"etlMetaKey,omitempty"`
	EtlMetaName *string `json:"etlMetaName,omitempty" xml:"etlMetaName,omitempty"`
	EtlMetaTag  *string `json:"etlMetaTag,omitempty" xml:"etlMetaTag,omitempty"`
	Offset      *int32  `json:"offset,omitempty" xml:"offset,omitempty"`
	Size        *int32  `json:"size,omitempty" xml:"size,omitempty"`
}

func (s ListEtlMetaRequest) String() string {
	return tea.Prettify(s)
}

func (s ListEtlMetaRequest) GoString() string {
	return s.String()
}

func (s *ListEtlMetaRequest) SetEtlMetaKey(v string) *ListEtlMetaRequest {
	s.EtlMetaKey = &v
	return s
}

func (s *ListEtlMetaRequest) SetEtlMetaName(v string) *ListEtlMetaRequest {
	s.EtlMetaName = &v
	return s
}

func (s *ListEtlMetaRequest) SetEtlMetaTag(v string) *ListEtlMetaRequest {
	s.EtlMetaTag = &v
	return s
}

func (s *ListEtlMetaRequest) SetOffset(v int32) *ListEtlMetaRequest {
	s.Offset = &v
	return s
}

func (s *ListEtlMetaRequest) SetSize(v int32) *ListEtlMetaRequest {
	s.Size = &v
	return s
}

type ListEtlMetaResponseBody struct {
	EtlMetaList []*EtlMeta `json:"etlMetaList,omitempty" xml:"etlMetaList,omitempty" type:"Repeated"`
	Total       *int32     `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListEtlMetaResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListEtlMetaResponseBody) GoString() string {
	return s.String()
}

func (s *ListEtlMetaResponseBody) SetEtlMetaList(v []*EtlMeta) *ListEtlMetaResponseBody {
	s.EtlMetaList = v
	return s
}

func (s *ListEtlMetaResponseBody) SetTotal(v int32) *ListEtlMetaResponseBody {
	s.Total = &v
	return s
}

type ListEtlMetaResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListEtlMetaResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListEtlMetaResponse) String() string {
	return tea.Prettify(s)
}

func (s ListEtlMetaResponse) GoString() string {
	return s.String()
}

func (s *ListEtlMetaResponse) SetHeaders(v map[string]*string) *ListEtlMetaResponse {
	s.Headers = v
	return s
}

func (s *ListEtlMetaResponse) SetStatusCode(v int32) *ListEtlMetaResponse {
	s.StatusCode = &v
	return s
}

func (s *ListEtlMetaResponse) SetBody(v *ListEtlMetaResponseBody) *ListEtlMetaResponse {
	s.Body = v
	return s
}

type ListEtlMetaNameRequest struct {
	Offset *int32 `json:"offset,omitempty" xml:"offset,omitempty"`
	Size   *int32 `json:"size,omitempty" xml:"size,omitempty"`
}

func (s ListEtlMetaNameRequest) String() string {
	return tea.Prettify(s)
}

func (s ListEtlMetaNameRequest) GoString() string {
	return s.String()
}

func (s *ListEtlMetaNameRequest) SetOffset(v int32) *ListEtlMetaNameRequest {
	s.Offset = &v
	return s
}

func (s *ListEtlMetaNameRequest) SetSize(v int32) *ListEtlMetaNameRequest {
	s.Size = &v
	return s
}

type ListEtlMetaNameResponseBody struct {
	Count           *int32    `json:"count,omitempty" xml:"count,omitempty"`
	EtlMetaNameList []*string `json:"etlMetaNameList,omitempty" xml:"etlMetaNameList,omitempty" type:"Repeated"`
	Total           *int32    `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListEtlMetaNameResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListEtlMetaNameResponseBody) GoString() string {
	return s.String()
}

func (s *ListEtlMetaNameResponseBody) SetCount(v int32) *ListEtlMetaNameResponseBody {
	s.Count = &v
	return s
}

func (s *ListEtlMetaNameResponseBody) SetEtlMetaNameList(v []*string) *ListEtlMetaNameResponseBody {
	s.EtlMetaNameList = v
	return s
}

func (s *ListEtlMetaNameResponseBody) SetTotal(v int32) *ListEtlMetaNameResponseBody {
	s.Total = &v
	return s
}

type ListEtlMetaNameResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListEtlMetaNameResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListEtlMetaNameResponse) String() string {
	return tea.Prettify(s)
}

func (s ListEtlMetaNameResponse) GoString() string {
	return s.String()
}

func (s *ListEtlMetaNameResponse) SetHeaders(v map[string]*string) *ListEtlMetaNameResponse {
	s.Headers = v
	return s
}

func (s *ListEtlMetaNameResponse) SetStatusCode(v int32) *ListEtlMetaNameResponse {
	s.StatusCode = &v
	return s
}

func (s *ListEtlMetaNameResponse) SetBody(v *ListEtlMetaNameResponseBody) *ListEtlMetaNameResponse {
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

type MergeShardsRequest struct {
	Action *string `json:"action,omitempty" xml:"action,omitempty"`
}

func (s MergeShardsRequest) String() string {
	return tea.Prettify(s)
}

func (s MergeShardsRequest) GoString() string {
	return s.String()
}

func (s *MergeShardsRequest) SetAction(v string) *MergeShardsRequest {
	s.Action = &v
	return s
}

type MergeShardsResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       []*Shard           `json:"body,omitempty" xml:"body,omitempty" require:"true" type:"Repeated"`
}

func (s MergeShardsResponse) String() string {
	return tea.Prettify(s)
}

func (s MergeShardsResponse) GoString() string {
	return s.String()
}

func (s *MergeShardsResponse) SetHeaders(v map[string]*string) *MergeShardsResponse {
	s.Headers = v
	return s
}

func (s *MergeShardsResponse) SetStatusCode(v int32) *MergeShardsResponse {
	s.StatusCode = &v
	return s
}

func (s *MergeShardsResponse) SetBody(v []*Shard) *MergeShardsResponse {
	s.Body = v
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
	Action     *string `json:"action,omitempty" xml:"action,omitempty"`
	Key        *string `json:"key,omitempty" xml:"key,omitempty"`
	ShardCount *int32  `json:"shardCount,omitempty" xml:"shardCount,omitempty"`
}

func (s SplitShardRequest) String() string {
	return tea.Prettify(s)
}

func (s SplitShardRequest) GoString() string {
	return s.String()
}

func (s *SplitShardRequest) SetAction(v string) *SplitShardRequest {
	s.Action = &v
	return s
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

type UnTagResourcesRequest struct {
	All          *bool     `json:"all,omitempty" xml:"all,omitempty"`
	ResourceId   []*string `json:"resourceId,omitempty" xml:"resourceId,omitempty" type:"Repeated"`
	ResourceType *string   `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
	Tags         []*string `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
}

func (s UnTagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s UnTagResourcesRequest) GoString() string {
	return s.String()
}

func (s *UnTagResourcesRequest) SetAll(v bool) *UnTagResourcesRequest {
	s.All = &v
	return s
}

func (s *UnTagResourcesRequest) SetResourceId(v []*string) *UnTagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *UnTagResourcesRequest) SetResourceType(v string) *UnTagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *UnTagResourcesRequest) SetTags(v []*string) *UnTagResourcesRequest {
	s.Tags = v
	return s
}

type UnTagResourcesResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
}

func (s UnTagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s UnTagResourcesResponse) GoString() string {
	return s.String()
}

func (s *UnTagResourcesResponse) SetHeaders(v map[string]*string) *UnTagResourcesResponse {
	s.Headers = v
	return s
}

func (s *UnTagResourcesResponse) SetStatusCode(v int32) *UnTagResourcesResponse {
	s.StatusCode = &v
	return s
}

type UpdateCheckPointRequest struct {
	Checkpoint   *string `json:"checkpoint,omitempty" xml:"checkpoint,omitempty"`
	Shard        *int32  `json:"shard,omitempty" xml:"shard,omitempty"`
	Consumer     *string `json:"consumer,omitempty" xml:"consumer,omitempty"`
	ForceSuccess *bool   `json:"forceSuccess,omitempty" xml:"forceSuccess,omitempty"`
	Type         *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s UpdateCheckPointRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateCheckPointRequest) GoString() string {
	return s.String()
}

func (s *UpdateCheckPointRequest) SetCheckpoint(v string) *UpdateCheckPointRequest {
	s.Checkpoint = &v
	return s
}

func (s *UpdateCheckPointRequest) SetShard(v int32) *UpdateCheckPointRequest {
	s.Shard = &v
	return s
}

func (s *UpdateCheckPointRequest) SetConsumer(v string) *UpdateCheckPointRequest {
	s.Consumer = &v
	return s
}

func (s *UpdateCheckPointRequest) SetForceSuccess(v bool) *UpdateCheckPointRequest {
	s.ForceSuccess = &v
	return s
}

func (s *UpdateCheckPointRequest) SetType(v string) *UpdateCheckPointRequest {
	s.Type = &v
	return s
}

type UpdateCheckPointResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
}

func (s UpdateCheckPointResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateCheckPointResponse) GoString() string {
	return s.String()
}

func (s *UpdateCheckPointResponse) SetHeaders(v map[string]*string) *UpdateCheckPointResponse {
	s.Headers = v
	return s
}

func (s *UpdateCheckPointResponse) SetStatusCode(v int32) *UpdateCheckPointResponse {
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
	Enable            *bool                              `json:"enable,omitempty" xml:"enable,omitempty"`
	EtlJobName        *string                            `json:"etlJobName,omitempty" xml:"etlJobName,omitempty"`
	FunctionConfig    *UpdateEtlJobRequestFunctionConfig `json:"functionConfig,omitempty" xml:"functionConfig,omitempty" type:"Struct"`
	FunctionParameter map[string]interface{}             `json:"functionParameter,omitempty" xml:"functionParameter,omitempty"`
	LogConfig         *UpdateEtlJobRequestLogConfig      `json:"logConfig,omitempty" xml:"logConfig,omitempty" type:"Struct"`
	SourceConfig      *UpdateEtlJobRequestSourceConfig   `json:"sourceConfig,omitempty" xml:"sourceConfig,omitempty" type:"Struct"`
	TriggerConfig     *UpdateEtlJobRequestTriggerConfig  `json:"triggerConfig,omitempty" xml:"triggerConfig,omitempty" type:"Struct"`
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

func (s *UpdateEtlJobRequest) SetEtlJobName(v string) *UpdateEtlJobRequest {
	s.EtlJobName = &v
	return s
}

func (s *UpdateEtlJobRequest) SetFunctionConfig(v *UpdateEtlJobRequestFunctionConfig) *UpdateEtlJobRequest {
	s.FunctionConfig = v
	return s
}

func (s *UpdateEtlJobRequest) SetFunctionParameter(v map[string]interface{}) *UpdateEtlJobRequest {
	s.FunctionParameter = v
	return s
}

func (s *UpdateEtlJobRequest) SetLogConfig(v *UpdateEtlJobRequestLogConfig) *UpdateEtlJobRequest {
	s.LogConfig = v
	return s
}

func (s *UpdateEtlJobRequest) SetSourceConfig(v *UpdateEtlJobRequestSourceConfig) *UpdateEtlJobRequest {
	s.SourceConfig = v
	return s
}

func (s *UpdateEtlJobRequest) SetTriggerConfig(v *UpdateEtlJobRequestTriggerConfig) *UpdateEtlJobRequest {
	s.TriggerConfig = v
	return s
}

type UpdateEtlJobRequestFunctionConfig struct {
	AccountId        *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
	Endpoint         *string `json:"endpoint,omitempty" xml:"endpoint,omitempty"`
	FunctionName     *string `json:"functionName,omitempty" xml:"functionName,omitempty"`
	FunctionProvider *string `json:"functionProvider,omitempty" xml:"functionProvider,omitempty"`
	RegionName       *string `json:"regionName,omitempty" xml:"regionName,omitempty"`
	RoleArn          *string `json:"roleArn,omitempty" xml:"roleArn,omitempty"`
	ServiceName      *string `json:"serviceName,omitempty" xml:"serviceName,omitempty"`
}

func (s UpdateEtlJobRequestFunctionConfig) String() string {
	return tea.Prettify(s)
}

func (s UpdateEtlJobRequestFunctionConfig) GoString() string {
	return s.String()
}

func (s *UpdateEtlJobRequestFunctionConfig) SetAccountId(v string) *UpdateEtlJobRequestFunctionConfig {
	s.AccountId = &v
	return s
}

func (s *UpdateEtlJobRequestFunctionConfig) SetEndpoint(v string) *UpdateEtlJobRequestFunctionConfig {
	s.Endpoint = &v
	return s
}

func (s *UpdateEtlJobRequestFunctionConfig) SetFunctionName(v string) *UpdateEtlJobRequestFunctionConfig {
	s.FunctionName = &v
	return s
}

func (s *UpdateEtlJobRequestFunctionConfig) SetFunctionProvider(v string) *UpdateEtlJobRequestFunctionConfig {
	s.FunctionProvider = &v
	return s
}

func (s *UpdateEtlJobRequestFunctionConfig) SetRegionName(v string) *UpdateEtlJobRequestFunctionConfig {
	s.RegionName = &v
	return s
}

func (s *UpdateEtlJobRequestFunctionConfig) SetRoleArn(v string) *UpdateEtlJobRequestFunctionConfig {
	s.RoleArn = &v
	return s
}

func (s *UpdateEtlJobRequestFunctionConfig) SetServiceName(v string) *UpdateEtlJobRequestFunctionConfig {
	s.ServiceName = &v
	return s
}

type UpdateEtlJobRequestLogConfig struct {
	Endpoint     *string `json:"endpoint,omitempty" xml:"endpoint,omitempty"`
	LogstoreName *string `json:"logstoreName,omitempty" xml:"logstoreName,omitempty"`
	ProjectName  *string `json:"projectName,omitempty" xml:"projectName,omitempty"`
}

func (s UpdateEtlJobRequestLogConfig) String() string {
	return tea.Prettify(s)
}

func (s UpdateEtlJobRequestLogConfig) GoString() string {
	return s.String()
}

func (s *UpdateEtlJobRequestLogConfig) SetEndpoint(v string) *UpdateEtlJobRequestLogConfig {
	s.Endpoint = &v
	return s
}

func (s *UpdateEtlJobRequestLogConfig) SetLogstoreName(v string) *UpdateEtlJobRequestLogConfig {
	s.LogstoreName = &v
	return s
}

func (s *UpdateEtlJobRequestLogConfig) SetProjectName(v string) *UpdateEtlJobRequestLogConfig {
	s.ProjectName = &v
	return s
}

type UpdateEtlJobRequestSourceConfig struct {
	LogstoreName *string `json:"logstoreName,omitempty" xml:"logstoreName,omitempty"`
}

func (s UpdateEtlJobRequestSourceConfig) String() string {
	return tea.Prettify(s)
}

func (s UpdateEtlJobRequestSourceConfig) GoString() string {
	return s.String()
}

func (s *UpdateEtlJobRequestSourceConfig) SetLogstoreName(v string) *UpdateEtlJobRequestSourceConfig {
	s.LogstoreName = &v
	return s
}

type UpdateEtlJobRequestTriggerConfig struct {
	MaxRetryTime     *int32  `json:"maxRetryTime,omitempty" xml:"maxRetryTime,omitempty"`
	RoleArn          *string `json:"roleArn,omitempty" xml:"roleArn,omitempty"`
	StartingPosition *string `json:"startingPosition,omitempty" xml:"startingPosition,omitempty"`
	StartingUnixtime *int64  `json:"startingUnixtime,omitempty" xml:"startingUnixtime,omitempty"`
	TriggerInterval  *int32  `json:"triggerInterval,omitempty" xml:"triggerInterval,omitempty"`
}

func (s UpdateEtlJobRequestTriggerConfig) String() string {
	return tea.Prettify(s)
}

func (s UpdateEtlJobRequestTriggerConfig) GoString() string {
	return s.String()
}

func (s *UpdateEtlJobRequestTriggerConfig) SetMaxRetryTime(v int32) *UpdateEtlJobRequestTriggerConfig {
	s.MaxRetryTime = &v
	return s
}

func (s *UpdateEtlJobRequestTriggerConfig) SetRoleArn(v string) *UpdateEtlJobRequestTriggerConfig {
	s.RoleArn = &v
	return s
}

func (s *UpdateEtlJobRequestTriggerConfig) SetStartingPosition(v string) *UpdateEtlJobRequestTriggerConfig {
	s.StartingPosition = &v
	return s
}

func (s *UpdateEtlJobRequestTriggerConfig) SetStartingUnixtime(v int64) *UpdateEtlJobRequestTriggerConfig {
	s.StartingUnixtime = &v
	return s
}

func (s *UpdateEtlJobRequestTriggerConfig) SetTriggerInterval(v int32) *UpdateEtlJobRequestTriggerConfig {
	s.TriggerInterval = &v
	return s
}

type UpdateEtlJobResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
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

type UpdateEtlMetaRequest struct {
	Enable       *bool                  `json:"enable,omitempty" xml:"enable,omitempty"`
	EtlMetaKey   *string                `json:"etlMetaKey,omitempty" xml:"etlMetaKey,omitempty"`
	EtlMetaName  *string                `json:"etlMetaName,omitempty" xml:"etlMetaName,omitempty"`
	EtlMetaTag   *string                `json:"etlMetaTag,omitempty" xml:"etlMetaTag,omitempty"`
	EtlMetaValue map[string]interface{} `json:"etlMetaValue,omitempty" xml:"etlMetaValue,omitempty"`
}

func (s UpdateEtlMetaRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateEtlMetaRequest) GoString() string {
	return s.String()
}

func (s *UpdateEtlMetaRequest) SetEnable(v bool) *UpdateEtlMetaRequest {
	s.Enable = &v
	return s
}

func (s *UpdateEtlMetaRequest) SetEtlMetaKey(v string) *UpdateEtlMetaRequest {
	s.EtlMetaKey = &v
	return s
}

func (s *UpdateEtlMetaRequest) SetEtlMetaName(v string) *UpdateEtlMetaRequest {
	s.EtlMetaName = &v
	return s
}

func (s *UpdateEtlMetaRequest) SetEtlMetaTag(v string) *UpdateEtlMetaRequest {
	s.EtlMetaTag = &v
	return s
}

func (s *UpdateEtlMetaRequest) SetEtlMetaValue(v map[string]interface{}) *UpdateEtlMetaRequest {
	s.EtlMetaValue = v
	return s
}

type UpdateEtlMetaResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
}

func (s UpdateEtlMetaResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateEtlMetaResponse) GoString() string {
	return s.String()
}

func (s *UpdateEtlMetaResponse) SetHeaders(v map[string]*string) *UpdateEtlMetaResponse {
	s.Headers = v
	return s
}

func (s *UpdateEtlMetaResponse) SetStatusCode(v int32) *UpdateEtlMetaResponse {
	s.StatusCode = &v
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

func (client *Client) BatchCreateEtlMeta(project *string, request *BatchCreateEtlMetaRequest) (_result *BatchCreateEtlMetaResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &BatchCreateEtlMetaResponse{}
	_body, _err := client.BatchCreateEtlMetaWithOptions(project, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BatchCreateEtlMetaWithOptions(project *string, request *BatchCreateEtlMetaRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *BatchCreateEtlMetaResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EtlMetaList)) {
		body["etlMetaList"] = request.EtlMetaList
	}

	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("BatchCreateEtlMeta"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/etlmetas"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &BatchCreateEtlMetaResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BatchModifyEtlMetaStatus(project *string, request *BatchModifyEtlMetaStatusRequest) (_result *BatchModifyEtlMetaStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &BatchModifyEtlMetaStatusResponse{}
	_body, _err := client.BatchModifyEtlMetaStatusWithOptions(project, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BatchModifyEtlMetaStatusWithOptions(project *string, request *BatchModifyEtlMetaStatusRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *BatchModifyEtlMetaStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["type"] = request.Type
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EtlMetaKeyList)) {
		body["etlMetaKeyList"] = request.EtlMetaKeyList
	}

	if !tea.BoolValue(util.IsUnset(request.EtlMetaName)) {
		body["etlMetaName"] = request.EtlMetaName
	}

	if !tea.BoolValue(util.IsUnset(request.EtlMetaTag)) {
		body["etlMetaTag"] = request.EtlMetaTag
	}

	if !tea.BoolValue(util.IsUnset(request.Range)) {
		body["range"] = request.Range
	}

	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("BatchModifyEtlMetaStatus"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/etlmetas"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &BatchModifyEtlMetaStatusResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BatchUpdateEtlMeta(project *string, request *BatchUpdateEtlMetaRequest) (_result *BatchUpdateEtlMetaResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &BatchUpdateEtlMetaResponse{}
	_body, _err := client.BatchUpdateEtlMetaWithOptions(project, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BatchUpdateEtlMetaWithOptions(project *string, request *BatchUpdateEtlMetaRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *BatchUpdateEtlMetaResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.EtlMetaList))) {
		body["etlMetaList"] = request.EtlMetaList
	}

	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("BatchUpdateEtlMeta"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/etlmetas"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &BatchUpdateEtlMetaResponse{}
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

func (client *Client) CreateEtlJob(project *string, request *CreateEtlJobRequest) (_result *CreateEtlJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateEtlJobResponse{}
	_body, _err := client.CreateEtlJobWithOptions(project, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateEtlJobWithOptions(project *string, request *CreateEtlJobRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateEtlJobResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.EtlJobName)) {
		body["etlJobName"] = request.EtlJobName
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.FunctionConfig))) {
		body["functionConfig"] = request.FunctionConfig
	}

	if !tea.BoolValue(util.IsUnset(request.FunctionParameter)) {
		body["functionParameter"] = request.FunctionParameter
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.LogConfig))) {
		body["logConfig"] = request.LogConfig
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.SourceConfig))) {
		body["sourceConfig"] = request.SourceConfig
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.TriggerConfig))) {
		body["triggerConfig"] = request.TriggerConfig
	}

	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateEtlJob"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/etljobs"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &CreateEtlJobResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateEtlMeta(project *string, request *CreateEtlMetaRequest) (_result *CreateEtlMetaResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateEtlMetaResponse{}
	_body, _err := client.CreateEtlMetaWithOptions(project, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateEtlMetaWithOptions(project *string, request *CreateEtlMetaRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateEtlMetaResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.EtlMetaKey)) {
		body["etlMetaKey"] = request.EtlMetaKey
	}

	if !tea.BoolValue(util.IsUnset(request.EtlMetaName)) {
		body["etlMetaName"] = request.EtlMetaName
	}

	if !tea.BoolValue(util.IsUnset(request.EtlMetaTag)) {
		body["etlMetaTag"] = request.EtlMetaTag
	}

	if !tea.BoolValue(util.IsUnset(request.EtlMetaValue)) {
		body["etlMetaValue"] = request.EtlMetaValue
	}

	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateEtlMeta"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/etlmetas"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &CreateEtlMetaResponse{}
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

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.Line))) {
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

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.EncryptConf))) {
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

func (client *Client) CreateMachineGroupWithOptions(project *string, request *CreateMachineGroupRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateMachineGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.GroupAttribute))) {
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

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.TargetConfiguration))) {
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

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.Parameter))) {
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

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.TargetConfiguration))) {
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

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.Parameter))) {
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

func (client *Client) DeleteEtlJob(project *string, etlJobName *string) (_result *DeleteEtlJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteEtlJobResponse{}
	_body, _err := client.DeleteEtlJobWithOptions(project, etlJobName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteEtlJobWithOptions(project *string, etlJobName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteEtlJobResponse, _err error) {
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteEtlJob"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/etljobs/" + tea.StringValue(etlJobName)),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &DeleteEtlJobResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteEtlMeta(project *string, request *DeleteEtlMetaRequest) (_result *DeleteEtlMetaResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteEtlMetaResponse{}
	_body, _err := client.DeleteEtlMetaWithOptions(project, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteEtlMetaWithOptions(project *string, request *DeleteEtlMetaRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteEtlMetaResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EtlMetaKey)) {
		query["etlMetaKey"] = request.EtlMetaKey
	}

	if !tea.BoolValue(util.IsUnset(request.EtlMetaName)) {
		query["etlMetaName"] = request.EtlMetaName
	}

	if !tea.BoolValue(util.IsUnset(request.EtlMetaTag)) {
		query["etlMetaTag"] = request.EtlMetaTag
	}

	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteEtlMeta"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/etlmetas"),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &DeleteEtlMetaResponse{}
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

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["type"] = request.Type
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
		Pathname:    tea.String("/logstores/" + tea.StringValue(logstore) + "/shards/" + tea.StringValue(shardId)),
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

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["type"] = request.Type
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
		Pathname:    tea.String("/logstores/" + tea.StringValue(logstore) + "/shards/" + tea.StringValue(shardId)),
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
		Pathname:    tea.String("/etljobs/" + tea.StringValue(etlJobName)),
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

func (client *Client) GetEtlMeta(project *string, request *GetEtlMetaRequest) (_result *GetEtlMetaResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetEtlMetaResponse{}
	_body, _err := client.GetEtlMetaWithOptions(project, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetEtlMetaWithOptions(project *string, request *GetEtlMetaRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetEtlMetaResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EtlMetaKey)) {
		query["etlMetaKey"] = request.EtlMetaKey
	}

	if !tea.BoolValue(util.IsUnset(request.EtlMetaName)) {
		query["etlMetaName"] = request.EtlMetaName
	}

	if !tea.BoolValue(util.IsUnset(request.EtlMetaTag)) {
		query["etlMetaTag"] = request.EtlMetaTag
	}

	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetEtlMeta"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/etlmetas"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetEtlMetaResponse{}
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

func (client *Client) GetShipperConfig(project *string, logstore *string, shipperName *string) (_result *GetShipperConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetShipperConfigResponse{}
	_body, _err := client.GetShipperConfigWithOptions(project, logstore, shipperName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetShipperConfigWithOptions(project *string, logstore *string, shipperName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetShipperConfigResponse, _err error) {
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetShipperConfig"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/logstores/" + tea.StringValue(logstore) + "/shipper/" + tea.StringValue(shipperName)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetShipperConfigResponse{}
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

func (client *Client) ListEtlJob(project *string, request *ListEtlJobRequest) (_result *ListEtlJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListEtlJobResponse{}
	_body, _err := client.ListEtlJobWithOptions(project, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListEtlJobWithOptions(project *string, request *ListEtlJobRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListEtlJobResponse, _err error) {
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
		Action:      tea.String("ListEtlJob"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/etljobs"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListEtlJobResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListEtlMeta(project *string, request *ListEtlMetaRequest) (_result *ListEtlMetaResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListEtlMetaResponse{}
	_body, _err := client.ListEtlMetaWithOptions(project, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListEtlMetaWithOptions(project *string, request *ListEtlMetaRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListEtlMetaResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EtlMetaKey)) {
		query["etlMetaKey"] = request.EtlMetaKey
	}

	if !tea.BoolValue(util.IsUnset(request.EtlMetaName)) {
		query["etlMetaName"] = request.EtlMetaName
	}

	if !tea.BoolValue(util.IsUnset(request.EtlMetaTag)) {
		query["etlMetaTag"] = request.EtlMetaTag
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
		Action:      tea.String("ListEtlMeta"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/etlmetas"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListEtlMetaResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListEtlMetaName(project *string, request *ListEtlMetaNameRequest) (_result *ListEtlMetaNameResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListEtlMetaNameResponse{}
	_body, _err := client.ListEtlMetaNameWithOptions(project, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListEtlMetaNameWithOptions(project *string, request *ListEtlMetaNameRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListEtlMetaNameResponse, _err error) {
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
		Action:      tea.String("ListEtlMetaName"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/etlmetanames"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListEtlMetaNameResponse{}
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

func (client *Client) MergeShards(project *string, logstore *string, shardID *string, request *MergeShardsRequest) (_result *MergeShardsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &MergeShardsResponse{}
	_body, _err := client.MergeShardsWithOptions(project, logstore, shardID, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) MergeShardsWithOptions(project *string, logstore *string, shardID *string, request *MergeShardsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *MergeShardsResponse, _err error) {
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
	}
	params := &openapi.Params{
		Action:      tea.String("MergeShards"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/logstores/" + tea.StringValue(logstore) + "/shards/" + tea.StringValue(shardID)),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("array"),
	}
	_result = &MergeShardsResponse{}
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

func (client *Client) SplitShard(project *string, logstore *string, shardID *string, request *SplitShardRequest) (_result *SplitShardResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &SplitShardResponse{}
	_body, _err := client.SplitShardWithOptions(project, logstore, shardID, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SplitShardWithOptions(project *string, logstore *string, shardID *string, request *SplitShardRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *SplitShardResponse, _err error) {
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
		Pathname:    tea.String("/logstores/" + tea.StringValue(logstore) + "/shards/" + tea.StringValue(shardID)),
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

func (client *Client) UnTagResources(request *UnTagResourcesRequest) (_result *UnTagResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UnTagResourcesResponse{}
	_body, _err := client.UnTagResourcesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UnTagResourcesWithOptions(request *UnTagResourcesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UnTagResourcesResponse, _err error) {
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
		Action:      tea.String("UnTagResources"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/untag"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &UnTagResourcesResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateCheckPoint(project *string, logstore *string, consumerGroup *string, request *UpdateCheckPointRequest) (_result *UpdateCheckPointResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateCheckPointResponse{}
	_body, _err := client.UpdateCheckPointWithOptions(project, logstore, consumerGroup, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateCheckPointWithOptions(project *string, logstore *string, consumerGroup *string, request *UpdateCheckPointRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateCheckPointResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["type"] = request.Type
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Checkpoint)) {
		body["checkpoint"] = request.Checkpoint
	}

	if !tea.BoolValue(util.IsUnset(request.Shard)) {
		body["shard"] = request.Shard
	}

	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateCheckPoint"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/logstores/" + tea.StringValue(logstore) + "/consumergroups/" + tea.StringValue(consumerGroup)),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &UpdateCheckPointResponse{}
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

func (client *Client) UpdateEtlJob(project *string, etlJobName *string, request *UpdateEtlJobRequest) (_result *UpdateEtlJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateEtlJobResponse{}
	_body, _err := client.UpdateEtlJobWithOptions(project, etlJobName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateEtlJobWithOptions(project *string, etlJobName *string, request *UpdateEtlJobRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateEtlJobResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.EtlJobName)) {
		body["etlJobName"] = request.EtlJobName
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.FunctionConfig))) {
		body["functionConfig"] = request.FunctionConfig
	}

	if !tea.BoolValue(util.IsUnset(request.FunctionParameter)) {
		body["functionParameter"] = request.FunctionParameter
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.LogConfig))) {
		body["logConfig"] = request.LogConfig
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.SourceConfig))) {
		body["sourceConfig"] = request.SourceConfig
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.TriggerConfig))) {
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
		Pathname:    tea.String("/etljobs/" + tea.StringValue(etlJobName)),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &UpdateEtlJobResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateEtlMeta(project *string, request *UpdateEtlMetaRequest) (_result *UpdateEtlMetaResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateEtlMetaResponse{}
	_body, _err := client.UpdateEtlMetaWithOptions(project, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateEtlMetaWithOptions(project *string, request *UpdateEtlMetaRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateEtlMetaResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.EtlMetaKey)) {
		body["etlMetaKey"] = request.EtlMetaKey
	}

	if !tea.BoolValue(util.IsUnset(request.EtlMetaName)) {
		body["etlMetaName"] = request.EtlMetaName
	}

	if !tea.BoolValue(util.IsUnset(request.EtlMetaTag)) {
		body["etlMetaTag"] = request.EtlMetaTag
	}

	if !tea.BoolValue(util.IsUnset(request.EtlMetaValue)) {
		body["etlMetaValue"] = request.EtlMetaValue
	}

	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateEtlMeta"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/etlmetas"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &UpdateEtlMetaResponse{}
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

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.Line))) {
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

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.EncryptConf))) {
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

func (client *Client) UpdateMachineGroupWithOptions(project *string, groupName *string, request *UpdateMachineGroupRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateMachineGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.GroupAttribute))) {
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

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.TargetConfiguration))) {
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

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.Parameter))) {
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

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.TargetConfiguration))) {
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

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.Parameter))) {
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
