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

type LogContent struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
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
	LogTags []*LogTag  `json:"LogTags,omitempty" xml:"LogTags,omitempty" type:"Repeated"`
	Logs    []*LogItem `json:"Logs,omitempty" xml:"Logs,omitempty" type:"Repeated"`
	Source  *string    `json:"Source,omitempty" xml:"Source,omitempty"`
	Topic   *string    `json:"Topic,omitempty" xml:"Topic,omitempty"`
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
	Contents []*LogContent `json:"Contents,omitempty" xml:"Contents,omitempty" type:"Repeated"`
	Time     *int32        `json:"Time,omitempty" xml:"Time,omitempty"`
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
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
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

type LogtailPipelineConfig struct {
	Aggregators    []map[string]interface{} `json:"aggregators,omitempty" xml:"aggregators,omitempty" type:"Repeated"`
	ConfigName     *string                  `json:"configName,omitempty" xml:"configName,omitempty"`
	CreateTime     *int64                   `json:"createTime,omitempty" xml:"createTime,omitempty"`
	Flushers       []map[string]interface{} `json:"flushers,omitempty" xml:"flushers,omitempty" type:"Repeated"`
	Global         map[string]interface{}   `json:"global,omitempty" xml:"global,omitempty"`
	Inputs         []map[string]interface{} `json:"inputs,omitempty" xml:"inputs,omitempty" type:"Repeated"`
	LastModifyTime *int64                   `json:"lastModifyTime,omitempty" xml:"lastModifyTime,omitempty"`
	LogSample      *string                  `json:"logSample,omitempty" xml:"logSample,omitempty"`
	Processors     []map[string]interface{} `json:"processors,omitempty" xml:"processors,omitempty" type:"Repeated"`
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
	AnnotationdataId *string                                 `json:"annotationdataId,omitempty" xml:"annotationdataId,omitempty"`
	Annotations      map[string]*MLDataParamAnnotationsValue `json:"annotations,omitempty" xml:"annotations,omitempty"`
	Config           map[string]*string                      `json:"config,omitempty" xml:"config,omitempty"`
	CreateTime       *int64                                  `json:"createTime,omitempty" xml:"createTime,omitempty"`
	DataHash         *string                                 `json:"dataHash,omitempty" xml:"dataHash,omitempty"`
	DatasetId        *string                                 `json:"datasetId,omitempty" xml:"datasetId,omitempty"`
	LastModifyTime   *int64                                  `json:"lastModifyTime,omitempty" xml:"lastModifyTime,omitempty"`
	Predictions      map[string]*MLDataParamPredictionsValue `json:"predictions,omitempty" xml:"predictions,omitempty"`
	Value            *string                                 `json:"value,omitempty" xml:"value,omitempty"`
	ValueType        *string                                 `json:"valueType,omitempty" xml:"valueType,omitempty"`
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
	CreateBy       *string `json:"createBy,omitempty" xml:"createBy,omitempty"`
	CreateTime     *int64  `json:"createTime,omitempty" xml:"createTime,omitempty"`
	DataType       *string `json:"dataType,omitempty" xml:"dataType,omitempty"`
	DatasetId      *string `json:"datasetId,omitempty" xml:"datasetId,omitempty"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty"`
	LabelId        *string `json:"labelId,omitempty" xml:"labelId,omitempty"`
	LastModifyTime *int64  `json:"lastModifyTime,omitempty" xml:"lastModifyTime,omitempty"`
	Name           *string `json:"name,omitempty" xml:"name,omitempty"`
	SettingType    *string `json:"settingType,omitempty" xml:"settingType,omitempty"`
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
	CreateTime     *int64                  `json:"createTime,omitempty" xml:"createTime,omitempty"`
	Description    *string                 `json:"description,omitempty" xml:"description,omitempty"`
	LabelId        *string                 `json:"labelId,omitempty" xml:"labelId,omitempty"`
	LastModifyTime *int64                  `json:"lastModifyTime,omitempty" xml:"lastModifyTime,omitempty"`
	Name           *string                 `json:"name,omitempty" xml:"name,omitempty"`
	Settings       []*MLLabelParamSettings `json:"settings,omitempty" xml:"settings,omitempty" type:"Repeated"`
	Type           *string                 `json:"type,omitempty" xml:"type,omitempty"`
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
	Config  *string `json:"config,omitempty" xml:"config,omitempty"`
	Mode    *string `json:"mode,omitempty" xml:"mode,omitempty"`
	Type    *string `json:"type,omitempty" xml:"type,omitempty"`
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
	Description     *string                 `json:"description,omitempty" xml:"description,omitempty"`
	Model           *MLServiceParamModel    `json:"model,omitempty" xml:"model,omitempty" type:"Struct"`
	Name            *string                 `json:"name,omitempty" xml:"name,omitempty"`
	Resource        *MLServiceParamResource `json:"resource,omitempty" xml:"resource,omitempty" type:"Struct"`
	ServiceType     *string                 `json:"serviceType,omitempty" xml:"serviceType,omitempty"`
	Status          *string                 `json:"status,omitempty" xml:"status,omitempty"`
	UpdateTimestamp *int64                  `json:"updateTimestamp,omitempty" xml:"updateTimestamp,omitempty"`
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
	ModelResourceId   *string `json:"modelResourceId,omitempty" xml:"modelResourceId,omitempty"`
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
	CpuLimit    *int32 `json:"cpuLimit,omitempty" xml:"cpuLimit,omitempty"`
	Gpu         *int32 `json:"gpu,omitempty" xml:"gpu,omitempty"`
	MemoryLimit *int32 `json:"memoryLimit,omitempty" xml:"memoryLimit,omitempty"`
	Replica     *int32 `json:"replica,omitempty" xml:"replica,omitempty"`
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

type Ticket struct {
	CallerUid      *int64  `json:"callerUid,omitempty" xml:"callerUid,omitempty"`
	CreateDate     *string `json:"createDate,omitempty" xml:"createDate,omitempty"`
	ExpirationTime *int64  `json:"expirationTime,omitempty" xml:"expirationTime,omitempty"`
	ExpireDate     *string `json:"expireDate,omitempty" xml:"expireDate,omitempty"`
	Extra          *string `json:"extra,omitempty" xml:"extra,omitempty"`
	Name           *string `json:"name,omitempty" xml:"name,omitempty"`
	Number         *int32  `json:"number,omitempty" xml:"number,omitempty"`
	Ticket         *string `json:"ticket,omitempty" xml:"ticket,omitempty"`
	TicketId       *string `json:"ticketId,omitempty" xml:"ticketId,omitempty"`
	UsedNumber     *int32  `json:"usedNumber,omitempty" xml:"usedNumber,omitempty"`
	Valid          *bool   `json:"valid,omitempty" xml:"valid,omitempty"`
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

type ServiceStatus struct {
	Enabled *bool   `json:"enabled,omitempty" xml:"enabled,omitempty"`
	Status  *string `json:"status,omitempty" xml:"status,omitempty"`
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

type MLDataParamAnnotationsValue struct {
	AnnotatedBy *string              `json:"annotatedBy,omitempty" xml:"annotatedBy,omitempty"`
	UpdateTime  *int64               `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
	Results     []map[string]*string `json:"results,omitempty" xml:"results,omitempty" type:"Repeated"`
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
	AnnotatedBy *string              `json:"annotatedBy,omitempty" xml:"annotatedBy,omitempty"`
	UpdateTime  *int64               `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
	Results     []map[string]*string `json:"results,omitempty" xml:"results,omitempty" type:"Repeated"`
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
	// Specifies whether to enable case sensitivity. This parameter is required only when **type** is set to **text**. Valid values:
	//
	// *   true
	// *   false (default)
	CaseSensitive *bool `json:"caseSensitive,omitempty" xml:"caseSensitive,omitempty"`
	// Specifies whether to include Chinese characters. This parameter is required only when **type** is set to **text**. Valid values:
	//
	// *   true
	// *   false (default)
	Chn *bool `json:"chn,omitempty" xml:"chn,omitempty"`
	// The data type of the field value. Valid values: text, json, double, and long.
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The alias of the field.
	Alias *string `json:"alias,omitempty" xml:"alias,omitempty"`
	// The delimiters that are used to split text.
	Token []*string `json:"token,omitempty" xml:"token,omitempty" type:"Repeated"`
	// Specifies whether to turn on Enable Analytics for the field.
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

type ConsumerGroupHeartBeatRequest struct {
	Body     []*int32 `json:"body,omitempty" xml:"body,omitempty" type:"Repeated"`
	Consumer *string  `json:"consumer,omitempty" xml:"consumer,omitempty"`
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
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       []*int32           `json:"body,omitempty" xml:"body,omitempty" require:"true" type:"Repeated"`
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

type CreateAnnotationDataSetRequest struct {
	Body      *MLDataSetParam `json:"body,omitempty" xml:"body,omitempty"`
	DatasetId *string         `json:"datasetId,omitempty" xml:"datasetId,omitempty"`
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
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
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
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
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
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
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
	ConsumerGroup *string `json:"consumerGroup,omitempty" xml:"consumerGroup,omitempty"`
	// Specifies whether to consume data in sequence. Valid values:
	//
	// *   true
	//
	//     *   In a shard, data is consumed in ascending order based on the value of the \*\*\__tag\_\_:\__receive_time\_\_\*\* field.
	//     *   If a shard is split, data in the original shard is consumed first. Then, data in the new shards is consumed at the same time.
	//     *   If shards are merged, data in the original shards is consumed first. Then, data in the new shard is consumed.
	//
	// *   false Data in all shards is consumed at the same time. If a new shard is generated after a shard is split or after shards are merged, data in the new shard is immediately consumed.
	Order *bool `json:"order,omitempty" xml:"order,omitempty"`
	// The timeout period. If the server does not receive heartbeats from a consumer within the timeout period, the server deletes the consumer. Unit: seconds.
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

type CreateDashboardRequest struct {
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
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
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

type CreateIndexRequest struct {
	// The configuration of field indexes. A field index is a key-value pair in which the key specifies the name of the field and the value specifies the index configuration of the field. You must specify this parameter, the line parameter, or both parameters. For more information, see Example.
	Keys map[string]*KeysValue `json:"keys,omitempty" xml:"keys,omitempty"`
	// The configuration of full-text indexes. You must specify this parameter, the keys parameter, or both parameters. For more information, see Example.
	Line *CreateIndexRequestLine `json:"line,omitempty" xml:"line,omitempty" type:"Struct"`
	// Specifies whether to turn on LogReduce. After you turn on LogReduce, either the whitelist or blacklist takes effect.
	LogReduce *bool `json:"log_reduce,omitempty" xml:"log_reduce,omitempty"`
	// The fields in the blacklist that you want to use to cluster logs.
	LogReduceBlackList []*string `json:"log_reduce_black_list,omitempty" xml:"log_reduce_black_list,omitempty" type:"Repeated"`
	// The fields in the whitelist that you want to use to cluster logs.
	LogReduceWhiteList []*string `json:"log_reduce_white_list,omitempty" xml:"log_reduce_white_list,omitempty" type:"Repeated"`
	// The maximum length of a field value that can be retained. Default value: 2048. Unit: bytes. The default value is equal to 2 KB. You can change the value of max_text_len. Valid values: 64 to 16384.
	MaxTextLen *int32 `json:"max_text_len,omitempty" xml:"max_text_len,omitempty"`
	// The retention period of logs. Unit: days. Valid values: 7, 30, and 90.
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
	// *   true
	// *   false (default)
	CaseSensitive *bool `json:"caseSensitive,omitempty" xml:"caseSensitive,omitempty"`
	// Specifies whether to include Chinese characters. Valid values:
	//
	// *   true
	// *   false (default)
	Chn *bool `json:"chn,omitempty" xml:"chn,omitempty"`
	// The excluded fields. You cannot specify both include_keys and exclude_keys.
	ExcludeKeys []*string `json:"exclude_keys,omitempty" xml:"exclude_keys,omitempty" type:"Repeated"`
	// The included fields. You cannot specify both include_keys and exclude_keys.
	IncludeKeys []*string `json:"include_keys,omitempty" xml:"include_keys,omitempty" type:"Repeated"`
	// The delimiters. You can specify a delimiter to delimit the content of a field value. For more information about delimiters, see Example.
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
	// Specifies whether to record public IP addresses. Default value: false. Valid values:
	//
	// *   true
	// *   false
	AppendMeta *bool `json:"appendMeta,omitempty" xml:"appendMeta,omitempty"`
	// Specifies whether to enable automatic sharding. Valid values:
	//
	// *   true
	// *   false
	AutoSplit *bool `json:"autoSplit,omitempty" xml:"autoSplit,omitempty"`
	// Specifies whether to enable the web tracking feature. Default value: false. Valid values:
	//
	// *   true
	// *   false
	EnableTracking *bool `json:"enable_tracking,omitempty" xml:"enable_tracking,omitempty"`
	// The data structure of the encryption configuration.
	EncryptConf *EncryptConf `json:"encrypt_conf,omitempty" xml:"encrypt_conf,omitempty"`
	// The retention period of data in the hot storage tier of the Logstore. Unit: days. You can specify a value that ranges from 30 to the value of ttl.
	//
	// Hot data that is stored for longer than the period specified by hot_ttl is converted to cold data. For more information, see [Enable hot and cold-tiered storage for a Logstore](~~308645~~).
	HotTtl *int32 `json:"hot_ttl,omitempty" xml:"hot_ttl,omitempty"`
	// The name of the Logstore. The name must meet the following requirements:
	//
	// *   The name must be unique in a project.
	// *   The name can contain only lowercase letters, digits, hyphens (-), and underscores (\_).
	// *   The name must start and end with a lowercase letter or a digit.
	// *   The name must be 3 to 63 characters in length.
	LogstoreName *string `json:"logstoreName,omitempty" xml:"logstoreName,omitempty"`
	// The maximum number of shards into which existing shards can be automatically split. Valid values: 1 to 64.
	//
	// > If you set autoSplit to true, you must configure this parameter.
	MaxSplitShard *int32 `json:"maxSplitShard,omitempty" xml:"maxSplitShard,omitempty"`
	// The type of the Logstore. Log Service provides the following types of Logstores: Standard Logstores and Query Logstores. Valid values:
	//
	// *   **standard**: Standard Logstore. This type of Logstore supports the log analysis feature and is suitable for scenarios such as real-time monitoring and interactive analysis. You can also use this type of Logstore to build a comprehensive observability system.
	// *   **query**: Query Logstore. This type of Logstore supports high-performance queries. The index traffic fee of a Query Logstore is approximately half that of a Standard Logstore. Query Logstores do not support SQL analysis. Query Logstores are suitable for scenarios in which the amount of data is large, the log retention period is long, or log analysis is not required. Log retention periods of weeks or months are considered long.
	Mode *string `json:"mode,omitempty" xml:"mode,omitempty"`
	// The number of shards.
	//
	// > You cannot call the CreateLogStore operation to change the number of shards. You can call the SplitShard or MergeShards operation to change the number of shards.
	ShardCount *int32 `json:"shardCount,omitempty" xml:"shardCount,omitempty"`
	// The type of the observable data. Valid values:
	//
	// *   None: logs
	// *   Metrics: metrics
	TelemetryType *string `json:"telemetryType,omitempty" xml:"telemetryType,omitempty"`
	// The retention period of data. Unit: days. Valid values: 1 to 3000. If you set this parameter to 3650, data is permanently stored.
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
	// The configurations of service logs.
	LoggingDetails []*CreateLoggingRequestLoggingDetails `json:"loggingDetails,omitempty" xml:"loggingDetails,omitempty" type:"Repeated"`
	// The name of the project to which service logs are stored.
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
	Logstore *string `json:"logstore,omitempty" xml:"logstore,omitempty"`
	// The type of service logs. Valid values:
	//
	// *   consumergroup_log: the consumption delay logs of consumer groups.
	// *   logtail_alarm: the alert logs of Logtail.
	// *   operation_log: the operation logs.
	// *   logtail_profile: the collection logs of Logtail.
	// *   metering: the metering logs.
	// *   logtail_status: the status logs of Logtail.
	// *   scheduledsqlalert: the run logs of Scheduled SQL jobs.
	// *   etl_alert: the run logs of data transformation jobs.
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

type CreateLogtailPipelineConfigRequest struct {
	Aggregators []map[string]interface{} `json:"aggregators,omitempty" xml:"aggregators,omitempty" type:"Repeated"`
	ConfigName  *string                  `json:"configName,omitempty" xml:"configName,omitempty"`
	Flushers    []map[string]interface{} `json:"flushers,omitempty" xml:"flushers,omitempty" type:"Repeated"`
	Global      map[string]interface{}   `json:"global,omitempty" xml:"global,omitempty"`
	Inputs      []map[string]interface{} `json:"inputs,omitempty" xml:"inputs,omitempty" type:"Repeated"`
	LogSample   *string                  `json:"logSample,omitempty" xml:"logSample,omitempty"`
	Processors  []map[string]interface{} `json:"processors,omitempty" xml:"processors,omitempty" type:"Repeated"`
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
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
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
	// *   The name of each machine group in a project must be unique.
	// *   It can contain only lowercase letters, digits, hyphens (-), and underscores (\_).
	// *   It must start and end with a lowercase letter or a digit.
	// *   It must be 3 to 128 characters in length.
	GroupName *string `json:"groupName,omitempty" xml:"groupName,omitempty"`
	// The type of the machine group. The parameter can be left empty.
	GroupType *string `json:"groupType,omitempty" xml:"groupType,omitempty"`
	// The type of the machine group identifier. Valid values:
	//
	// *   ip: The machine group uses IP addresses as identifiers.
	// *   userdefined: The machine group uses custom identifiers.
	MachineIdentifyType *string `json:"machineIdentifyType,omitempty" xml:"machineIdentifyType,omitempty"`
	// The identifiers of machine group.
	//
	// *   If you set machineIdentifyType to ip, enter the IP address of the machine.
	// *   If you set machineIdentifyType to userdefined, enter a custom identifier.
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
	ExternalName *string `json:"externalName,omitempty" xml:"externalName,omitempty"`
	// The log topic of the machine group.
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

type CreateOssExternalStoreRequest struct {
	// The name of the external store.
	ExternalStoreName *string `json:"externalStoreName,omitempty" xml:"externalStoreName,omitempty"`
	// The parameters that are configured for the external store.
	Parameter *CreateOssExternalStoreRequestParameter `json:"parameter,omitempty" xml:"parameter,omitempty" type:"Struct"`
	// The type of the external store. Set the value to oss.
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
	Accessid *string `json:"accessid,omitempty" xml:"accessid,omitempty"`
	// The AccessKey secret of your account.
	Accesskey *string `json:"accesskey,omitempty" xml:"accesskey,omitempty"`
	// The name of the OSS bucket.
	Bucket *string `json:"bucket,omitempty" xml:"bucket,omitempty"`
	// The associated fields.
	Columns []*CreateOssExternalStoreRequestParameterColumns `json:"columns,omitempty" xml:"columns,omitempty" type:"Repeated"`
	// The OSS endpoint.
	Endpoint *string `json:"endpoint,omitempty" xml:"endpoint,omitempty"`
	// The associated objects.
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
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The type of the field.
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

type CreateProjectRequest struct {
	// Data redundancy type
	DataRedundancyType *string `json:"dataRedundancyType,omitempty" xml:"dataRedundancyType,omitempty"`
	// The description of the project.
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The name of the project. The name must be unique in a region. You cannot change the name after you create the project. The name must meet the following requirements:
	//
	// *   The name must be unique.
	// *   It can contain only lowercase letters, digits, and hyphens (-).
	// *   It must start and end with a lowercase letter or a digit.
	// *   It must be 3 to 63 characters in length.
	ProjectName *string `json:"projectName,omitempty" xml:"projectName,omitempty"`
	// The ID of the resource group.
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
	// The name of the external store. The name must be unique in a project and must be different from Logstore names.
	ExternalStoreName *string `json:"externalStoreName,omitempty" xml:"externalStoreName,omitempty"`
	// The parameter struct.
	Parameter *CreateRdsExternalStoreRequestParameter `json:"parameter,omitempty" xml:"parameter,omitempty" type:"Struct"`
	// The storage type. Set the value to rds-vpc, which indicates an ApsaraDB RDS for MySQL database in a virtual private cloud (VPC).
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
	Db *string `json:"db,omitempty" xml:"db,omitempty"`
	// The internal or public endpoint of the ApsaraDB RDS for MySQL instance.
	Host *string `json:"host,omitempty" xml:"host,omitempty"`
	// The ID of the ApsaraDB RDS for MySQL instance.
	InstanceId *string `json:"instance-id,omitempty" xml:"instance-id,omitempty"`
	// The password that is used to log on to the ApsaraDB RDS for MySQL instance.
	Password *string `json:"password,omitempty" xml:"password,omitempty"`
	// The internal or public port of the ApsaraDB RDS for MySQL instance.
	Port *string `json:"port,omitempty" xml:"port,omitempty"`
	// The region where the ApsaraDB RDS for MySQL instance resides. Valid values: cn-qingdao, cn-beijing, and cn-hangzhou.
	Region *string `json:"region,omitempty" xml:"region,omitempty"`
	// The name of the database table in the ApsaraDB RDS for MySQL instance.
	Table *string `json:"table,omitempty" xml:"table,omitempty"`
	// The username that is used to log on to the ApsaraDB RDS for MySQL instance.
	Username *string `json:"username,omitempty" xml:"username,omitempty"`
	// The ID of the VPC to which the ApsaraDB RDS for MySQL instance belongs.
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
	// The display name.
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	// The name of the Logstore to which the saved search belongs.
	Logstore *string `json:"logstore,omitempty" xml:"logstore,omitempty"`
	// The name of the saved search. The name must be 3 to 63 characters in length.
	SavedsearchName *string `json:"savedsearchName,omitempty" xml:"savedsearchName,omitempty"`
	// The query statement of the saved search. A query statement consists of a search statement and an analytic statement in the `Search statement|Analytic statement` format. For more information about search statements and analytic statements, see [Log search overview](~~43772~~) and [Log analysis overview](~~53608~~).
	SearchQuery *string `json:"searchQuery,omitempty" xml:"searchQuery,omitempty"`
	// The topic of the log.
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

type DeleteAnnotationDataResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
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
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
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
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
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
	DataCode    *string `json:"dataCode,omitempty" xml:"dataCode,omitempty"`
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
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteCollectionPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
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

type DeleteDashboardResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
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

type DeleteLogtailPipelineConfigResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
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

type GetAnnotationDataResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *MLDataParam       `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *MLDataSetParam    `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *MLLabelParam      `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// The number of returned machine groups.
	Count *int32 `json:"count,omitempty" xml:"count,omitempty"`
	// The names of the returned machine groups.
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
	// The shard ID.
	//
	// *   If the specified shard does not exist, an empty list is returned.
	// *   If no shard ID is specified, the checkpoints of all shards are returned.
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
	// The shard ID.
	Shard *int32 `json:"shard,omitempty" xml:"shard,omitempty"`
	// The value of the checkpoint.
	Checkpoint *string `json:"checkpoint,omitempty" xml:"checkpoint,omitempty"`
	// The time when the checkpoint was last updated. The value is a UNIX timestamp representing the number of seconds that have elapsed since the epoch time January 1, 1970, 00:00:00 UTC.
	UpdateTime *int64 `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
	// The consumer at the checkpoint.
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
	DataCode    *string `json:"dataCode,omitempty" xml:"dataCode,omitempty"`
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
	Attribute         *GetCollectionPolicyResponseBodyCollectionPolicyAttribute        `json:"attribute,omitempty" xml:"attribute,omitempty" type:"Struct"`
	CentralizeConfig  *GetCollectionPolicyResponseBodyCollectionPolicyCentralizeConfig `json:"centralizeConfig,omitempty" xml:"centralizeConfig,omitempty" type:"Struct"`
	CentralizeEnabled *bool                                                            `json:"centralizeEnabled,omitempty" xml:"centralizeEnabled,omitempty"`
	DataCode          *string                                                          `json:"dataCode,omitempty" xml:"dataCode,omitempty"`
	Enabled           *string                                                          `json:"enabled,omitempty" xml:"enabled,omitempty"`
	PolicyConfig      *GetCollectionPolicyResponseBodyCollectionPolicyPolicyConfig     `json:"policyConfig,omitempty" xml:"policyConfig,omitempty" type:"Struct"`
	PolicyName        *string                                                          `json:"policyName,omitempty" xml:"policyName,omitempty"`
	ProductCode       *string                                                          `json:"productCode,omitempty" xml:"productCode,omitempty"`
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
	App         *string `json:"app,omitempty" xml:"app,omitempty"`
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
	DestLogstore *string `json:"destLogstore,omitempty" xml:"destLogstore,omitempty"`
	DestProject  *string `json:"destProject,omitempty" xml:"destProject,omitempty"`
	DestRegion   *string `json:"destRegion,omitempty" xml:"destRegion,omitempty"`
	DestTTL      *int32  `json:"destTTL,omitempty" xml:"destTTL,omitempty"`
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
	InstanceIds  []*string              `json:"instanceIds,omitempty" xml:"instanceIds,omitempty" type:"Repeated"`
	Regions      []*string              `json:"regions,omitempty" xml:"regions,omitempty" type:"Repeated"`
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
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetCollectionPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *LogtailConfig     `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	BackLines *int64 `json:"back_lines,omitempty" xml:"back_lines,omitempty"`
	// The number of logs that you want to obtain and are generated after the generation time of the start log. Valid values: (0,100].
	ForwardLines *int64 `json:"forward_lines,omitempty" xml:"forward_lines,omitempty"`
	// The unique identifier of the log group to which the start log belongs.
	PackId *string `json:"pack_id,omitempty" xml:"pack_id,omitempty"`
	// The unique context identifier of the start log in the log group.
	PackMeta *string `json:"pack_meta,omitempty" xml:"pack_meta,omitempty"`
	// The type of the data in the Logstore. Set the value to context_log.
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
	BackLines *int64 `json:"back_lines,omitempty" xml:"back_lines,omitempty"`
	// The number of logs that are generated after the generation time of the start log.
	ForwardLines *int64 `json:"forward_lines,omitempty" xml:"forward_lines,omitempty"`
	// The logs that are returned.
	Logs []map[string]interface{} `json:"logs,omitempty" xml:"logs,omitempty" type:"Repeated"`
	// Indicates whether the query and analysis results are complete. Valid values:
	//
	// *   Complete: The query is successful, and the complete query and analysis results are returned.
	// *   Incomplete: The query is successful, but the query and analysis results are incomplete. To obtain the complete results, you must repeat the request.
	Progress *string `json:"progress,omitempty" xml:"progress,omitempty"`
	// The total number of logs that are returned. The logs include the start log that is specified in the request.
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
	// The point in time that you want to use to query a cursor. Set the value to a UNIX timestamp or a string such as `begin` and `end`.
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
	// The cursor.
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

type GetDashboardResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *Dashboard         `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// The start time of the subinterval. The value is a UNIX timestamp representing the number of seconds that have elapsed since the epoch time January 1, 1970, 00:00:00 UTC.
	From *int64 `json:"from,omitempty" xml:"from,omitempty"`
	// The search statement. Only search statements are supported. Analytic statements are not supported. For more information about the syntax of search statements, see [Log search overview](~~43772~~).
	Query *string `json:"query,omitempty" xml:"query,omitempty"`
	// The end time of the subinterval. The value is a UNIX timestamp representing the number of seconds that have elapsed since the epoch time January 1, 1970, 00:00:00 UTC.
	To *int64 `json:"to,omitempty" xml:"to,omitempty"`
	// The topic of the logs.
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
	// The start time of the subinterval. The value is a UNIX timestamp representing the number of seconds that have elapsed since the epoch time January 1, 1970, 00:00:00 UTC.
	//
	// The time range that is specified in this operation is a left-closed, right-open interval. The interval includes the start time specified by the from parameter, but does not include the end time specified by the to parameter. If you specify the same value for the from and to parameters, the interval is invalid, and an error message is returned.
	From *int64 `json:"from,omitempty" xml:"from,omitempty"`
	// The end time of the subinterval. The value is a UNIX timestamp representing the number of seconds that have elapsed since the epoch time January 1, 1970, 00:00:00 UTC.
	//
	// The time range that is specified in this operation is a left-closed, right-open interval. The interval includes the start time specified by the from parameter, but does not include the end time specified by the to parameter. If you specify the same value for the from and to parameters, the interval is invalid, and an error message is returned.
	To *int64 `json:"to,omitempty" xml:"to,omitempty"`
	// The number of logs that are generated within the subinterval.
	Count *int64 `json:"count,omitempty" xml:"count,omitempty"`
	// Indicates whether the query and analysis results in the subinterval is complete. Valid values:
	//
	// Complete: The query is successful, and the complete query and analysis results are returned.
	//
	// Incomplete: The query is successful, but the query and analysis results are incomplete. To obtain the complete results, you must repeat the request.
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
	IndexMode *string `json:"index_mode,omitempty" xml:"index_mode,omitempty"`
	// The configurations of field indexes. A field index is in the key-value format in which the key specifies the name of the field and the value specifies the index configuration of the field.
	Keys map[string]*KeysValue `json:"keys,omitempty" xml:"keys,omitempty"`
	// The time when the index configurations were last updated. The value is a UNIX timestamp representing the number of seconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	LastModifyTime *int64 `json:"lastModifyTime,omitempty" xml:"lastModifyTime,omitempty"`
	// The configurations of full-text indexes.
	Line *GetIndexResponseBodyLine `json:"line,omitempty" xml:"line,omitempty" type:"Struct"`
	// Indicates whether the log clustering feature is enabled.
	LogReduce *bool `json:"log_reduce,omitempty" xml:"log_reduce,omitempty"`
	// The fields in the blacklist that are used to cluster logs. This parameter is valid only if the log clustering feature is enabled.
	LogReduceBlackList []*string `json:"log_reduce_black_list,omitempty" xml:"log_reduce_black_list,omitempty" type:"Repeated"`
	// The fields in the whitelist that are used to cluster logs. This parameter is valid only if the log clustering feature is enabled.
	LogReduceWhiteList []*string `json:"log_reduce_white_list,omitempty" xml:"log_reduce_white_list,omitempty" type:"Repeated"`
	// The maximum length of a field value that can be retained. Default value: 2048. Unit: bytes. The default value is equal to 2 KB. You can change the value of the max_text_len parameter. Valid values: 64 to 16384. Unit: bytes.
	MaxTextLen *int32 `json:"max_text_len,omitempty" xml:"max_text_len,omitempty"`
	// The storage type. The value is fixed as pg.
	Storage *string `json:"storage,omitempty" xml:"storage,omitempty"`
	// The lifecycle of the index file. Valid values: 7, 30, and 90. Unit: day.
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
	// *   true
	// *   false
	CaseSensitive *bool `json:"caseSensitive,omitempty" xml:"caseSensitive,omitempty"`
	// Indicates whether Chinese characters are included. Valid values:
	//
	// *   true
	// *   false
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

type GetLogStoreMeteringModeResponseBody struct {
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
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetLogStoreMeteringModeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// The beginning of the time range to query. The value is the log time that is specified when log data is written.
	//
	// *   The time range that is specified in this operation is a left-closed, right-open interval. The interval includes the start time specified by the **from** parameter, but does not include the end time specified by the **to** parameter. If you specify the same value for the **from** and **to** parameters, the interval is invalid, and an error message is returned.
	// *   The value is a UNIX timestamp representing the number of seconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// > To ensure that full data can be queried, specify a query time range that is accurate to the minute. If you also specify a time range in an analytic statement, Simple Log Service uses the time range specified in the analytic statement for query and analysis.
	//
	// If you want to specify a time range that is accurate to the second in your analytic statement, you must use the from_unixtime or to_unixtime function to convert the time format. For more information about the functions, see [from_unixtime function](~~63451~~) and [to_unixtime function](~~63451~~). Examples:
	//
	// *   `* | SELECT * FROM log WHERE from_unixtime(__time__) > from_unixtime(1664186624) AND from_unixtime(__time__) < now()`
	// *   `* | SELECT * FROM log WHERE __time__ > to_unixtime(date_parse(\"2022-10-19 15:46:05\", \"%Y-%m-%d %H:%i:%s\")) AND __time__ < to_unixtime(now())`
	From *int32 `json:"from,omitempty" xml:"from,omitempty"`
	// The maximum number of logs to return for the request. This parameter takes effect only when the query parameter is set to a search statement. Minimum value: 0. Maximum value: 100. Default value: 100.
	Line *int64 `json:"line,omitempty" xml:"line,omitempty"`
	// The line from which the query starts. This parameter takes effect only when the query parameter is set to a search statement. Default value: 0.
	Offset *int64 `json:"offset,omitempty" xml:"offset,omitempty"`
	// Specifies whether to enable the Dedicated SQL feature. For more information, see [Enable Dedicated SQL](~~223777~~). Valid values:
	//
	// *   true: enables the Dedicated SQL feature.
	// *   false (default): enables the Standard SQL feature.
	//
	// You can use the powerSql or **query** parameter to configure Dedicated SQL.
	PowerSql *bool `json:"powerSql,omitempty" xml:"powerSql,omitempty"`
	// The search statement or the query statement. For more information, see [Log search overview](~~43772~~) and [Log analysis overview](~~53608~~). If you add `set session parallel_sql=true;` to the analytic statement in the query parameter, Dedicated SQL is used. For example, you can set the query parameter to `* | set session parallel_sql=true; select count(*) as pv`. For more information about common errors that may occur during log query and analysis, see [How do I resolve common errors that occur when I query and analyze logs?](~~61628~~)
	//
	// > If you specify an analytic statement in the value of the query parameter, the line and offset parameters do not take effect. In this case, we recommend that you set the line and offset parameters to 0 and use the LIMIT clause to limit the number of logs to return on each page. For more information, see [Paged query](~~89994~~).
	Query *string `json:"query,omitempty" xml:"query,omitempty"`
	// Specifies whether to return logs in reverse chronological order of log timestamps. The log timestamps are accurate to the minute. Valid values:
	//
	// *   true: returns logs in reverse chronological order of log timestamps.
	// *   false (default): returns logs in chronological order of log timestamps.
	//
	// >
	//
	// *   The reverse parameter takes effect only when the query parameter is set to a search statement. The reverse parameter specifies the method used to sort returned logs.
	// *   If the query parameter is set to a query statement, the reverse parameter does not take effect. The method used to sort returned logs is specified by the ORDER BY clause in the analytic statement. If you use the keyword asc in the ORDER BY clause, the logs are sorted in chronological order. If you use the keyword desc in the ORDER BY clause, the logs are sorted in reverse chronological order. By default, asc is used in the ORDER BY clause.
	Reverse *bool `json:"reverse,omitempty" xml:"reverse,omitempty"`
	// The end of the time range to query. The value is the log time that is specified when log data is written.
	//
	// *   The time range that is specified in this operation is a left-closed, right-open interval. The interval includes the start time specified by the **from** parameter, but does not include the end time specified by the **to** parameter. If you specify the same value for the **from** and **to** parameters, the interval is invalid, and an error message is returned.
	// *   The value is a UNIX timestamp representing the number of seconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// > To ensure that full data can be queried, specify a query time range that is accurate to the minute. If you also specify a time range in an analytic statement, Simple Log Service uses the time range specified in the analytic statement for query and analysis.
	//
	// If you want to specify a time range that is accurate to the second in your analytic statement, you must use the from_unixtime or to_unixtime function to convert the time format. For more information about the functions, see [from_unixtime function](~~63451~~) and [to_unixtime function](~~63451~~). Examples:
	//
	// *   `* | SELECT * FROM log WHERE from_unixtime(__time__) > from_unixtime(1664186624) AND from_unixtime(__time__) < now()`
	// *   `* | SELECT * FROM log WHERE __time__ > to_unixtime(date_parse(\"2022-10-19 15:46:05\", \"%Y-%m-%d %H:%i:%s\")) AND __time__ < to_unixtime(now())`
	To *int32 `json:"to,omitempty" xml:"to,omitempty"`
	// The topic of the logs. The default value is double quotation marks (""). For more information, see [Topic](~~48881~~).
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

type GetLogsV2Headers struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// The compression method.
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
	// Specifies whether to page forward or backward for the scan-based query or the phrase query.
	Forward *bool `json:"forward,omitempty" xml:"forward,omitempty"`
	// The beginning of the time range to query. The value is the log time that is specified when log data is written.
	//
	// The time range specified by the from and to parameters is a left-closed and right-open interval. Each interval includes the specified start time but does not include the specified end time. If you specify the same value for the from and to parameters, the interval is invalid, and an error message is returned. The value is a UNIX timestamp representing the number of seconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	From *int32 `json:"from,omitempty" xml:"from,omitempty"`
	// The maximum number of logs to return for the request. This parameter takes effect only when the query parameter is set to a search statement. Valid values: 0 to 100. Default value: 100.
	Line *int64 `json:"line,omitempty" xml:"line,omitempty"`
	// The row from which the query starts. This parameter takes effect only when the query parameter is set to a search statement. Default value: 0.
	Offset *int64 `json:"offset,omitempty" xml:"offset,omitempty"`
	// Specifies whether to enable the SQL enhancement feature. By default, the feature is disabled.
	PowerSql *bool `json:"powerSql,omitempty" xml:"powerSql,omitempty"`
	// The search statement or the query statement. For more information, see the "Log search overview" and "Log analysis overview" topics.
	//
	// If you add set session parallel_sql=true; to the analytic statement in the query parameter, the dedicated SQL feature is enabled. Example: \* | set session parallel_sql=true; select count(\*) as pv.
	//
	// Note: If you specify an analytic statement in the query parameter, the line and offset parameters are invalid for this operation. In this case, we recommend that you set the line and offset parameters to 0 and use a LIMIT clause to limit the number of entries to return on each page. For more information, see the "Perform paged queries" topic.
	Query *string `json:"query,omitempty" xml:"query,omitempty"`
	// Specifies whether to return logs in reverse chronological order of log timestamps. The log timestamps are accurate to the minute. Valid values:
	//
	// true: returns logs in reverse chronological order of log timestamps. false (default): returns logs in chronological order of log timestamps. Note The reverse parameter takes effect only when the query parameter is set to a search statement. The reverse parameter specifies the method used to sort the returned logs. If the query parameter is set to a query statement, which consists of a search statement and an analytic statement, the reverse parameter does not take effect. The method used to sort the returned logs is specified by the ORDER BY clause in the analytic statement. If you use the keyword asc in the ORDER BY clause, the logs are sorted in chronological order. If you use the keyword desc in the ORDER BY clause, the logs are sorted in reverse chronological order. By default, asc is used in the ORDER BY clause.
	Reverse *bool   `json:"reverse,omitempty" xml:"reverse,omitempty"`
	Session *string `json:"session,omitempty" xml:"session,omitempty"`
	// The ID of the shard.
	Shard *int32 `json:"shard,omitempty" xml:"shard,omitempty"`
	// The end of the time range to query. The value is the log time that is specified when log data is written.
	//
	// The time range specified by the from and to parameters is a left-closed and right-open interval. Each interval includes the specified start time but does not include the specified end time. If you specify the same value for the from and to parameters, the interval is invalid, and an error message is returned. The value is a UNIX timestamp representing the number of seconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	To *int32 `json:"to,omitempty" xml:"to,omitempty"`
	// The topic of the logs. Default value: double quotation marks ("").
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

func (s *GetLogsV2Request) SetShard(v int32) *GetLogsV2Request {
	s.Shard = &v
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
	AggQuery *string `json:"aggQuery,omitempty" xml:"aggQuery,omitempty"`
	// The number of rows that are returned.
	Count *int32 `json:"count,omitempty" xml:"count,omitempty"`
	// The amount of time that is consumed by the request. Unit: milliseconds.
	ElapsedMillisecond *int64 `json:"elapsedMillisecond,omitempty" xml:"elapsedMillisecond,omitempty"`
	// Indicates whether the query is an SQL query.
	HasSQL *bool `json:"hasSQL,omitempty" xml:"hasSQL,omitempty"`
	// Indicates whether the returned result is accurate.
	IsAccurate *bool `json:"isAccurate,omitempty" xml:"isAccurate,omitempty"`
	// All keys in the query result.
	Keys []*string `json:"keys,omitempty" xml:"keys,omitempty" type:"Repeated"`
	// The number of logs that are processed in the request.
	ProcessedBytes *int64 `json:"processedBytes,omitempty" xml:"processedBytes,omitempty"`
	// The number of rows that are processed in the request.
	ProcessedRows *int32 `json:"processedRows,omitempty" xml:"processedRows,omitempty"`
	// Indicates whether the query result is complete. Valid values:
	//
	// *   Complete: The query was successful, and the complete result is returned.
	// *   Incomplete: The query was successful, but the query result is incomplete. To obtain the complete result, you must call the operation again.
	Progress *string `json:"progress,omitempty" xml:"progress,omitempty"`
	// The type of observable data.
	TelementryType *string `json:"telementryType,omitempty" xml:"telementryType,omitempty"`
	// All terms in the query statement.
	Terms []map[string]interface{} `json:"terms,omitempty" xml:"terms,omitempty" type:"Repeated"`
	// The part before | in the query statement.
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

func (s *GetLogsV2ResponseBodyMeta) SetCount(v int32) *GetLogsV2ResponseBodyMeta {
	s.Count = &v
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

func (s *GetLogsV2ResponseBodyMeta) SetIsAccurate(v bool) *GetLogsV2ResponseBodyMeta {
	s.IsAccurate = &v
	return s
}

func (s *GetLogsV2ResponseBodyMeta) SetKeys(v []*string) *GetLogsV2ResponseBodyMeta {
	s.Keys = v
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

type GetLogsV2Response struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetLogsV2ResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *LogtailPipelineConfig `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	AllowBuiltin *bool                   `json:"allowBuiltin,omitempty" xml:"allowBuiltin,omitempty"`
	Body         *MLServiceAnalysisParam `json:"body,omitempty" xml:"body,omitempty"`
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
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetMLServiceResultsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// Specifies whether to enable the Dedicated SQL feature. For more information, see [Enable Dedicated SQL](~~223777~~). Valid values:
	//
	// *   true
	// *   false (default): enables the Standard SQL feature.
	//
	// You can use the powerSql or **query** parameter to configure Dedicated SQL.
	PowerSql *bool `json:"powerSql,omitempty" xml:"powerSql,omitempty"`
	// The standard SQL statement. In this example, the SQL statement queries the number of page views (PVs) from 2022-03-01 10:41:40 to 2022-03-01 10:56:40 in a Logstore whose name is nginx-moni.
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
	// The start time of the log shipping job. The value is a UNIX timestamp representing the number of seconds that have elapsed since the epoch time January 1, 1970, 00:00:00 UTC.
	From *int64 `json:"from,omitempty" xml:"from,omitempty"`
	// The line from which the query starts. Default value: 0.
	Offset *int32 `json:"offset,omitempty" xml:"offset,omitempty"`
	// The number of entries per page. Default value: 100. Maximum value: 500.
	Size *int32 `json:"size,omitempty" xml:"size,omitempty"`
	// The status of the log shipping job. This parameter is empty by default, which indicates that log shipping jobs in all states are returned. Valid values: success, fail, and running.
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The end time of the log shipping job. The value is a UNIX timestamp representing the number of seconds that have elapsed since the epoch time January 1, 1970, 00:00:00 UTC.
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
	Count *int64 `json:"count,omitempty" xml:"count,omitempty"`
	// The statistics about log shipping jobs.
	Statistics *GetShipperStatusResponseBodyStatistics `json:"statistics,omitempty" xml:"statistics,omitempty" type:"Struct"`
	// The details of log shipping jobs.
	Tasks *GetShipperStatusResponseBodyTasks `json:"tasks,omitempty" xml:"tasks,omitempty" type:"Struct"`
	// The total number of log shipping jobs.
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
	Fail *int64 `json:"fail,omitempty" xml:"fail,omitempty"`
	// The number of log shipping jobs that are in the running state.
	Running *int64 `json:"running,omitempty" xml:"running,omitempty"`
	// The number of log shipping jobs that are in the success state.
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
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// The error code of the log shipping job.
	TaskCode *string `json:"taskCode,omitempty" xml:"taskCode,omitempty"`
	// The start time of the log shipping job. The value is a UNIX timestamp representing the number of seconds that have elapsed since the epoch time January 1, 1970, 00:00:00 UTC.
	TaskCreateTime *int64 `json:"taskCreateTime,omitempty" xml:"taskCreateTime,omitempty"`
	// The number of logs that are shipped in the log shipping job.
	TaskDataLines *int32 `json:"taskDataLines,omitempty" xml:"taskDataLines,omitempty"`
	// The end time of the log shipping job. The value is a UNIX timestamp representing the number of seconds that have elapsed since the epoch time January 1, 1970, 00:00:00 UTC.
	TaskFinishTime *int64 `json:"taskFinishTime,omitempty" xml:"taskFinishTime,omitempty"`
	// The time when Simple Log Service receives the most recent log of the log shipping job. The value is a UNIX timestamp representing the number of seconds that have elapsed since the epoch time January 1, 1970, 00:00:00 UTC.
	TaskLastDataReceiveTime *int64 `json:"taskLastDataReceiveTime,omitempty" xml:"taskLastDataReceiveTime,omitempty"`
	// The error message of the log shipping job.
	TaskMessage *string `json:"taskMessage,omitempty" xml:"taskMessage,omitempty"`
	// The status of the log shipping job. Valid values: running, success, and fail.
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

type ListAnnotationDataRequest struct {
	Offset *int32 `json:"offset,omitempty" xml:"offset,omitempty"`
	Size   *int32 `json:"size,omitempty" xml:"size,omitempty"`
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
	Data  []*MLDataParam `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	Total *int32         `json:"total,omitempty" xml:"total,omitempty"`
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
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListAnnotationDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Offset *int32 `json:"offset,omitempty" xml:"offset,omitempty"`
	Size   *int32 `json:"size,omitempty" xml:"size,omitempty"`
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
	Data  []*MLDataSetParam `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	Total *int32            `json:"total,omitempty" xml:"total,omitempty"`
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
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListAnnotationDataSetsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Offset *int32 `json:"offset,omitempty" xml:"offset,omitempty"`
	Size   *int32 `json:"size,omitempty" xml:"size,omitempty"`
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
	Data  []*MLLabelParam `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	Total *int32          `json:"total,omitempty" xml:"total,omitempty"`
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
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListAnnotationLabelsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Attribute   *ListCollectionPoliciesRequestAttribute `json:"attribute,omitempty" xml:"attribute,omitempty" type:"Struct"`
	DataCode    *string                                 `json:"dataCode,omitempty" xml:"dataCode,omitempty"`
	InstanceId  *string                                 `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	PageNum     *int32                                  `json:"pageNum,omitempty" xml:"pageNum,omitempty"`
	PageSize    *int32                                  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	PolicyName  *string                                 `json:"policyName,omitempty" xml:"policyName,omitempty"`
	ProductCode *string                                 `json:"productCode,omitempty" xml:"productCode,omitempty"`
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
	App         *string `json:"app,omitempty" xml:"app,omitempty"`
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
	DataCode        *string `json:"dataCode,omitempty" xml:"dataCode,omitempty"`
	InstanceId      *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	PageNum         *int32  `json:"pageNum,omitempty" xml:"pageNum,omitempty"`
	PageSize        *int32  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	PolicyName      *string `json:"policyName,omitempty" xml:"policyName,omitempty"`
	ProductCode     *string `json:"productCode,omitempty" xml:"productCode,omitempty"`
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
	CurrentCount *int32                                    `json:"currentCount,omitempty" xml:"currentCount,omitempty"`
	Data         []*ListCollectionPoliciesResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	TotalCount   *int32                                    `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
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
	Attribute         *ListCollectionPoliciesResponseBodyDataAttribute        `json:"attribute,omitempty" xml:"attribute,omitempty" type:"Struct"`
	CentralizeConfig  *ListCollectionPoliciesResponseBodyDataCentralizeConfig `json:"centralizeConfig,omitempty" xml:"centralizeConfig,omitempty" type:"Struct"`
	CentralizeEnabled *bool                                                   `json:"centralizeEnabled,omitempty" xml:"centralizeEnabled,omitempty"`
	DataCode          *string                                                 `json:"dataCode,omitempty" xml:"dataCode,omitempty"`
	Enabled           *bool                                                   `json:"enabled,omitempty" xml:"enabled,omitempty"`
	PolicyConfig      *ListCollectionPoliciesResponseBodyDataPolicyConfig     `json:"policyConfig,omitempty" xml:"policyConfig,omitempty" type:"Struct"`
	PolicyName        *string                                                 `json:"policyName,omitempty" xml:"policyName,omitempty"`
	ProductCode       *string                                                 `json:"productCode,omitempty" xml:"productCode,omitempty"`
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
	App         *string `json:"app,omitempty" xml:"app,omitempty"`
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
	DestLogstore *string `json:"destLogstore,omitempty" xml:"destLogstore,omitempty"`
	DestProject  *string `json:"destProject,omitempty" xml:"destProject,omitempty"`
	DestRegion   *string `json:"destRegion,omitempty" xml:"destRegion,omitempty"`
	DestTTL      *int32  `json:"destTTL,omitempty" xml:"destTTL,omitempty"`
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
	InstanceIds  []*string              `json:"instanceIds,omitempty" xml:"instanceIds,omitempty" type:"Repeated"`
	Regions      []*string              `json:"regions,omitempty" xml:"regions,omitempty" type:"Repeated"`
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
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListCollectionPoliciesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	ConfigName   *string `json:"configName,omitempty" xml:"configName,omitempty"`
	LogstoreName *string `json:"logstoreName,omitempty" xml:"logstoreName,omitempty"`
	Offset       *int64  `json:"offset,omitempty" xml:"offset,omitempty"`
	Size         *int64  `json:"size,omitempty" xml:"size,omitempty"`
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
	Configs []*string `json:"configs,omitempty" xml:"configs,omitempty" type:"Repeated"`
	Count   *int32    `json:"count,omitempty" xml:"count,omitempty"`
	Total   *int32    `json:"total,omitempty" xml:"total,omitempty"`
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
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type ListDashboardRequest struct {
	Offset *int32 `json:"offset,omitempty" xml:"offset,omitempty"`
	Size   *int32 `json:"size,omitempty" xml:"size,omitempty"`
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
	DashboardItems []*ListDashboardResponseBodyDashboardItems `json:"dashboardItems,omitempty" xml:"dashboardItems,omitempty" type:"Repeated"`
	Dashboards     []*string                                  `json:"dashboards,omitempty" xml:"dashboards,omitempty" type:"Repeated"`
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
	DashboardName *string `json:"dashboardName,omitempty" xml:"dashboardName,omitempty"`
	DisplayName   *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
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
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListDashboardResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	DomainName *string `json:"domainName,omitempty" xml:"domainName,omitempty"`
	// The line from which the query starts. Default value: 0.
	Offset *int32 `json:"offset,omitempty" xml:"offset,omitempty"`
	// The number of entries per page. Default value: 500. Maximum value: 500.
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
	Count *int64 `json:"count,omitempty" xml:"count,omitempty"`
	// The domain names.
	Domains []*string `json:"domains,omitempty" xml:"domains,omitempty" type:"Repeated"`
	// The total number of domain names that are returned.
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

type ListExternalStoreRequest struct {
	// The name of the external store. You can query external stores that contain a specified string.
	ExternalStoreName *string `json:"externalStoreName,omitempty" xml:"externalStoreName,omitempty"`
	// The line from which the query starts. Default value: 0.
	Offset *int32 `json:"offset,omitempty" xml:"offset,omitempty"`
	// The number of entries per page. Maximum value: 500.
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
	Count *int32 `json:"count,omitempty" xml:"count,omitempty"`
	// The names of the external stores.
	Externalstores []*ExternalStore `json:"externalstores,omitempty" xml:"externalstores,omitempty" type:"Repeated"`
	// The number of external stores that meet the query conditions.
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
	// The name of the Logstore. Fuzzy match is supported. For example, if you enter test, Logstores whose name contains test are returned.
	LogstoreName *string `json:"logstoreName,omitempty" xml:"logstoreName,omitempty"`
	// The type of the Logstore. Valid values: standard and query.
	//
	// *   **standard**: Standard Logstore. This type of Logstore supports the log analysis feature and is suitable for scenarios such as real-time monitoring and interactive analysis. You can also use this type of Logstore to build a comprehensive observability system.
	// *   **query**: Query Logstore. This type of Logstore supports high-performance queries. The index traffic fee of a Query Logstore is approximately half that of a Standard Logstore. Query Logstores do not support SQL analysis. Query Logstores are suitable for scenarios in which the volume of data is large, the log retention period is long, or log analysis is not required. Log retention periods of weeks or months are considered long.
	Mode *string `json:"mode,omitempty" xml:"mode,omitempty"`
	// The line from which the query starts. Default value: 0.
	Offset *int32 `json:"offset,omitempty" xml:"offset,omitempty"`
	// The number of entries per page. Maximum value: 500. Default value: 500.
	Size *int32 `json:"size,omitempty" xml:"size,omitempty"`
	// The type of the data that you want to query. Valid values:
	//
	// *   None: logs
	// *   Metrics: metrics
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
	// The number of entries that are returned on the current page.
	Count *int32 `json:"count,omitempty" xml:"count,omitempty"`
	// The Logstores that meet the query conditions.
	Logstores []*string `json:"logstores,omitempty" xml:"logstores,omitempty" type:"Repeated"`
	// The number of the Logstores that meet the query conditions.
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

type ListLogtailPipelineConfigRequest struct {
	ConfigName   *string `json:"configName,omitempty" xml:"configName,omitempty"`
	LogstoreName *string `json:"logstoreName,omitempty" xml:"logstoreName,omitempty"`
	Offset       *int64  `json:"offset,omitempty" xml:"offset,omitempty"`
	Size         *int64  `json:"size,omitempty" xml:"size,omitempty"`
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
	Configs []*string `json:"configs,omitempty" xml:"configs,omitempty" type:"Repeated"`
	Count   *int32    `json:"count,omitempty" xml:"count,omitempty"`
	Total   *int32    `json:"total,omitempty" xml:"total,omitempty"`
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
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListLogtailPipelineConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	GroupName *string `json:"groupName,omitempty" xml:"groupName,omitempty"`
	// The line from which the query starts. Default value: 0.
	Offset *int32 `json:"offset,omitempty" xml:"offset,omitempty"`
	// The number of entries per page. Maximum value: 500.
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
	Count *int32 `json:"count,omitempty" xml:"count,omitempty"`
	// The machine groups that meet the query conditions.
	Machinegroups []*string `json:"machinegroups,omitempty" xml:"machinegroups,omitempty" type:"Repeated"`
	// The total number of machine groups that meet the query conditions.
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
	// The line from which the query starts. Default value: 0.
	Offset *int32 `json:"offset,omitempty" xml:"offset,omitempty"`
	// The number of entries per page. Default value: 100. Maximum value: 500.
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
	Count *int32 `json:"count,omitempty" xml:"count,omitempty"`
	// The machines that are returned.
	Machines []*Machine `json:"machines,omitempty" xml:"machines,omitempty" type:"Repeated"`
	// The total number of machines.
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
	// The line from which the query starts. Default value: 0.
	Offset *int32 `json:"offset,omitempty" xml:"offset,omitempty"`
	// The name of the project.
	ProjectName     *string `json:"projectName,omitempty" xml:"projectName,omitempty"`
	ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	// The number of entries per page. Default value: 100. This operation can return up to 500 projects.
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
	Count *int64 `json:"count,omitempty" xml:"count,omitempty"`
	// The projects that meet the query conditions.
	Projects []*Project `json:"projects,omitempty" xml:"projects,omitempty" type:"Repeated"`
	// The total number of projects that meet the query conditions.
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
	// The line from which the query starts. Default value: 0.
	Offset *int32 `json:"offset,omitempty" xml:"offset,omitempty"`
	// The number of entries per page. Maximum value: 500.
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
	Count *int32 `json:"count,omitempty" xml:"count,omitempty"`
	// The saved searches.
	SavedsearchItems []*SavedSearch `json:"savedsearchItems,omitempty" xml:"savedsearchItems,omitempty" type:"Repeated"`
	// The total number of saved searches that meet the query conditions.
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
	// The number of log shipping jobs returned.
	Count *int64 `json:"count,omitempty" xml:"count,omitempty"`
	// The names of the log shipping jobs.
	Shipper []*string `json:"shipper,omitempty" xml:"shipper,omitempty" type:"Repeated"`
	// The total number of log shipping jobs.
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
	// The IDs of the resources for which you want to query tags. You must specify at least one of resourceId and tags.
	ResourceId []*string `json:"resourceId,omitempty" xml:"resourceId,omitempty" type:"Repeated"`
	// The type of the resource. Set the value to project.
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
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// The value of the tag that you want to use to filter resources. If you set the value to null, resources are filtered based only on the key of the tag.
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
	ResourceId *string `json:"resourceId,omitempty" xml:"resourceId,omitempty"`
	// The type of the resource.
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
	// The key of the tag.
	TagKey *string `json:"tagKey,omitempty" xml:"tagKey,omitempty"`
	// The value of the tag.
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

type PutAnnotationDataRequest struct {
	AnnotationdataId *string              `json:"annotationdataId,omitempty" xml:"annotationdataId,omitempty"`
	MlDataParam      *MLDataParam         `json:"mlDataParam,omitempty" xml:"mlDataParam,omitempty"`
	RawLog           []map[string]*string `json:"rawLog,omitempty" xml:"rawLog,omitempty" type:"Repeated"`
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
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
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
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryMLServiceResultsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// The position where the shard is split.
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// The number of new shards that are generated after splitting.
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
	// The resource IDs. You can specify only one resource and add tags to the resource.
	ResourceId []*string `json:"resourceId,omitempty" xml:"resourceId,omitempty" type:"Repeated"`
	// The type of the resource. Set the value to project.
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
	// The tags that you want to add to the resource. Up to 20 tags are supported at a time. Each tag is a key-value pair.
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
	// *   The key must be `1 to 128` characters in length.
	// *   The key cannot contain `"http://"` or `"https://"`.
	// *   The key cannot start with `"acs:"` or `"aliyun"`.
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// The value of the tag. The value must meet the following requirements:
	//
	// *   The value must be `1 to 128` characters in length.
	// *   The value cannot contain `"http://"` or `"https://"`.
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
	ResourceId   []*string `json:"resourceId,omitempty" xml:"resourceId,omitempty" type:"Repeated"`
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

type UpdateAnnotationDataSetRequest struct {
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
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
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
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
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
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
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
	// *   true: If a shard is split, the data in the original shard is consumed first. Then, the data in the new shards is consumed at the same time. If shards are merged, the data in the original shards is consumed first. Then, the data in the new shard is consumed.
	// *   false: The data in all shards is consumed at the same time. If a new shard is generated after a shard is split or shards are merged, the data in the new shard is immediately consumed.
	Order *bool `json:"order,omitempty" xml:"order,omitempty"`
	// The timeout period. If Simple Log Service does not receive heartbeats from a consumer within the timeout period, Simple Log Service deletes the consumer. Unit: seconds.
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

type UpdateDashboardRequest struct {
	Attribute     map[string]*string `json:"attribute,omitempty" xml:"attribute,omitempty"`
	Charts        []*Chart           `json:"charts,omitempty" xml:"charts,omitempty" type:"Repeated"`
	DashboardName *string            `json:"dashboardName,omitempty" xml:"dashboardName,omitempty"`
	Description   *string            `json:"description,omitempty" xml:"description,omitempty"`
	DisplayName   *string            `json:"displayName,omitempty" xml:"displayName,omitempty"`
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
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
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

type UpdateIndexRequest struct {
	// The configuration of field indexes. A field index is a key-value pair in which the key specifies the name of the field and the value specifies the index configuration of the field.
	Keys map[string]*KeysValue `json:"keys,omitempty" xml:"keys,omitempty"`
	// The configuration of full-text indexes.
	Line *UpdateIndexRequestLine `json:"line,omitempty" xml:"line,omitempty" type:"Struct"`
	// Specifies whether to turn on LogReduce. If you turn on LogReduce, only one of `log_reduce_white_list` and `log_reduce_black_list` takes effect.
	LogReduce *bool `json:"log_reduce,omitempty" xml:"log_reduce,omitempty"`
	// The fields in the blacklist that you want to use to cluster logs.
	LogReduceBlackList []*string `json:"log_reduce_black_list,omitempty" xml:"log_reduce_black_list,omitempty" type:"Repeated"`
	// The fields in the whitelist that you want to use to cluster logs.
	LogReduceWhiteList []*string `json:"log_reduce_white_list,omitempty" xml:"log_reduce_white_list,omitempty" type:"Repeated"`
	// The maximum length of a field value that can be retained.
	MaxTextLen *int32 `json:"max_text_len,omitempty" xml:"max_text_len,omitempty"`
	// The retention period of data. Unit: days. Valid values: 7, 30, and 90.
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
	// *   true
	// *   false
	CaseSensitive *bool `json:"caseSensitive,omitempty" xml:"caseSensitive,omitempty"`
	// Specifies whether to include Chinese characters. Valid values:
	//
	// *   true
	// *   false
	Chn *bool `json:"chn,omitempty" xml:"chn,omitempty"`
	// The excluded fields. You cannot specify both include_keys and exclude_keys.
	ExcludeKeys []*string `json:"exclude_keys,omitempty" xml:"exclude_keys,omitempty" type:"Repeated"`
	// The included fields. You cannot specify both include_keys and exclude_keys.
	IncludeKeys []*string `json:"include_keys,omitempty" xml:"include_keys,omitempty" type:"Repeated"`
	// The delimiters that are used to split text.
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
	// Specifies whether to record public IP addresses. Default value: false. Valid values:
	//
	// *   true
	// *   false
	AppendMeta *bool `json:"appendMeta,omitempty" xml:"appendMeta,omitempty"`
	// Specifies whether to enable automatic sharding. Valid values:
	//
	// *   true
	// *   false
	AutoSplit *bool `json:"autoSplit,omitempty" xml:"autoSplit,omitempty"`
	// Specifies whether to enable the web tracking feature. Default value: false. Valid values:
	//
	// *   true
	// *   false
	EnableTracking *bool `json:"enable_tracking,omitempty" xml:"enable_tracking,omitempty"`
	// The data structure of the encryption configuration.
	EncryptConf *EncryptConf `json:"encrypt_conf,omitempty" xml:"encrypt_conf,omitempty"`
	// The retention period of data in the hot storage tier of the Logstore. Minimum value: 30. Unit: day. You can specify a value that ranges from 30 to the value of ttl. Hot data that is stored for longer than the period specified by hot_ttl is converted to cold data. For more information, see [Enable hot and cold-tiered storage for a Logstore](~~308645~~).
	HotTtl *int32 `json:"hot_ttl,omitempty" xml:"hot_ttl,omitempty"`
	// The name of the Logstore.
	LogstoreName *string `json:"logstoreName,omitempty" xml:"logstoreName,omitempty"`
	// The maximum number of shards into which existing shards can be automatically split. Valid values: 1 to 64.
	//
	// > If you set autoSplit to true, you must specify maxSplitShard.
	MaxSplitShard *int32 `json:"maxSplitShard,omitempty" xml:"maxSplitShard,omitempty"`
	// The type of the Logstore. Simple Log Service provides two types of Logstores: Standard Logstores and Query Logstores.
	//
	// *   **standard**: Standard Logstore. This type of Logstore supports the log analysis feature and is suitable for scenarios such as real-time monitoring and interactive analysis. You can also use this type of Logstore to build a comprehensive observability system.
	// *   **query**: Query Logstore. This type of Logstore supports high-performance queries. The index traffic fee of a Query Logstore is approximately half that of a Standard Logstore. Query Logstores do not support SQL analysis. Query Logstores are suitable for scenarios in which the volume of data is large, the log retention period is long, or log analysis is not required. Log retention periods of weeks or months are considered long.
	Mode *string `json:"mode,omitempty" xml:"mode,omitempty"`
	// The number of shards.
	//
	// > You cannot call the UpdateLogstore operation to change the number of shards. You can call the SplitShard or MergeShards operation to change the number of shards.
	ShardCount *int32 `json:"shardCount,omitempty" xml:"shardCount,omitempty"`
	// The type of the log that you want to query. Valid values:
	//
	// *   None: all types of logs.
	// *   Metrics: metrics.
	TelemetryType *string `json:"telemetryType,omitempty" xml:"telemetryType,omitempty"`
	// The retention period of data. Unit: day. Valid values: 1 to 3650. If you set ttl to 3650, data is permanently stored.
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

type UpdateLogStoreMeteringModeRequest struct {
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
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
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
	LoggingDetails []*UpdateLoggingRequestLoggingDetails `json:"loggingDetails,omitempty" xml:"loggingDetails,omitempty" type:"Repeated"`
	// The name of the project to which you want to save service logs.
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
	Logstore *string `json:"logstore,omitempty" xml:"logstore,omitempty"`
	// The type of service logs. Valid values:
	//
	// *   consumergroup_log: the consumption delay logs of consumer groups.
	// *   logtail_alarm: the alert logs of Logtail.
	// *   operation_log: the operation logs.
	// *   logtail_profile: the collection logs of Logtail.
	// *   metering: the metering logs.
	// *   logtail_status: the status logs of Logtail.
	// *   scheduledsqlalert: the operational logs of Scheduled SQL jobs.
	// *   etl_alert: the operational logs of data transformation jobs.
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

type UpdateLogtailPipelineConfigRequest struct {
	Aggregators []map[string]interface{} `json:"aggregators,omitempty" xml:"aggregators,omitempty" type:"Repeated"`
	ConfigName  *string                  `json:"configName,omitempty" xml:"configName,omitempty"`
	Flushers    []map[string]interface{} `json:"flushers,omitempty" xml:"flushers,omitempty" type:"Repeated"`
	Global      map[string]interface{}   `json:"global,omitempty" xml:"global,omitempty"`
	Inputs      []map[string]interface{} `json:"inputs,omitempty" xml:"inputs,omitempty" type:"Repeated"`
	LogSample   *string                  `json:"logSample,omitempty" xml:"logSample,omitempty"`
	Processors  []map[string]interface{} `json:"processors,omitempty" xml:"processors,omitempty" type:"Repeated"`
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
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
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
	GroupName *string `json:"groupName,omitempty" xml:"groupName,omitempty"`
	// The type of the machine group. Set the value to an empty string.
	GroupType *string `json:"groupType,omitempty" xml:"groupType,omitempty"`
	// The identifier type of the machine group. Valid values:
	//
	// *   ip: The machine group uses IP addresses as identifiers.
	// *   userdefined: The machine group uses custom identifiers.
	MachineIdentifyType *string `json:"machineIdentifyType,omitempty" xml:"machineIdentifyType,omitempty"`
	// The identifiers of the machines in the machine group.
	//
	// *   If you set machineIdentifyType to ip, enter the IP addresses of the machines.
	// *   If you set machineIdentifyType to userdefined, enter a custom identifier.
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
	ExternalName *string `json:"externalName,omitempty" xml:"externalName,omitempty"`
	// The topic of the machine group. This parameter is empty by default.
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
	// The operation on the machine. Valid values: add and delete. A value of add specifies to add the machine to the machine group. A value of delete specifies to remove the machine from the machine group.
	Action *string `json:"action,omitempty" xml:"action,omitempty"`
	// The machines to be added or removed.
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

type UpdateOssExternalStoreRequest struct {
	// The name of the external store.
	ExternalStoreName *string `json:"externalStoreName,omitempty" xml:"externalStoreName,omitempty"`
	// The parameters that are configured for the external store.
	Parameter *UpdateOssExternalStoreRequestParameter `json:"parameter,omitempty" xml:"parameter,omitempty" type:"Struct"`
	// The type of the external store. Set the value to oss.
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
	Accessid *string `json:"accessid,omitempty" xml:"accessid,omitempty"`
	// The AccessKey secret of your account.
	Accesskey *string `json:"accesskey,omitempty" xml:"accesskey,omitempty"`
	// The name of the OSS bucket.
	Bucket *string `json:"bucket,omitempty" xml:"bucket,omitempty"`
	// The associated fields.
	Columns []*UpdateOssExternalStoreRequestParameterColumns `json:"columns,omitempty" xml:"columns,omitempty" type:"Repeated"`
	// The OSS endpoint.
	Endpoint *string `json:"endpoint,omitempty" xml:"endpoint,omitempty"`
	// The associated objects.
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
	// The key of the field.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The type of the field.
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

type UpdateProjectRequest struct {
	// The description of the project. The default value is an empty string.
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
	// The name of the external store.
	ExternalStoreName *string `json:"externalStoreName,omitempty" xml:"externalStoreName,omitempty"`
	// The parameter struct.
	Parameter *UpdateRdsExternalStoreRequestParameter `json:"parameter,omitempty" xml:"parameter,omitempty" type:"Struct"`
	// The storage type. Set the value to rds-vpc, which indicates an ApsaraDB RDS for MySQL database in a virtual private cloud (VPC).
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
	Db *string `json:"db,omitempty" xml:"db,omitempty"`
	// The internal or public endpoint of the ApsaraDB RDS for MySQL instance.
	Host *string `json:"host,omitempty" xml:"host,omitempty"`
	// The ID of the ApsaraDB RDS for MySQL instance.
	InstanceId *string `json:"instance-id,omitempty" xml:"instance-id,omitempty"`
	// The password that is used to log on to the ApsaraDB RDS for MySQL instance.
	Password *string `json:"password,omitempty" xml:"password,omitempty"`
	// The internal or public port of the ApsaraDB RDS for MySQL instance.
	Port *string `json:"port,omitempty" xml:"port,omitempty"`
	// The region where the ApsaraDB RDS for MySQL instance resides. Valid values: cn-qingdao, cn-beijing, and cn-hangzhou.
	Region *string `json:"region,omitempty" xml:"region,omitempty"`
	// The name of the database table in the ApsaraDB RDS for MySQL instance.
	Table *string `json:"table,omitempty" xml:"table,omitempty"`
	// The username that is used to log on to the ApsaraDB RDS for MySQL instance.
	Username *string `json:"username,omitempty" xml:"username,omitempty"`
	// The ID of the VPC to which the ApsaraDB RDS for MySQL instance belongs.
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
	// The display name.
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	// The name of the Logstore to which the saved search belongs.
	Logstore *string `json:"logstore,omitempty" xml:"logstore,omitempty"`
	// The name of the saved search. The name must be 3 to 63 characters in length.
	SavedsearchName *string `json:"savedsearchName,omitempty" xml:"savedsearchName,omitempty"`
	// The search statement or the query statement of the saved search. A query statement consists of a search statement and an analytic statement in the Search statement|Analytic statement format.
	//
	// For more information, see Log search overview and Log analysis overview.
	SearchQuery *string `json:"searchQuery,omitempty" xml:"searchQuery,omitempty"`
	// The topic of the logs.
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

type UpsertCollectionPolicyRequest struct {
	Attribute         *UpsertCollectionPolicyRequestAttribute        `json:"attribute,omitempty" xml:"attribute,omitempty" type:"Struct"`
	CentralizeConfig  *UpsertCollectionPolicyRequestCentralizeConfig `json:"centralizeConfig,omitempty" xml:"centralizeConfig,omitempty" type:"Struct"`
	CentralizeEnabled *bool                                          `json:"centralizeEnabled,omitempty" xml:"centralizeEnabled,omitempty"`
	DataCode          *string                                        `json:"dataCode,omitempty" xml:"dataCode,omitempty"`
	Enabled           *bool                                          `json:"enabled,omitempty" xml:"enabled,omitempty"`
	PolicyConfig      *UpsertCollectionPolicyRequestPolicyConfig     `json:"policyConfig,omitempty" xml:"policyConfig,omitempty" type:"Struct"`
	PolicyName        *string                                        `json:"policyName,omitempty" xml:"policyName,omitempty"`
	ProductCode       *string                                        `json:"productCode,omitempty" xml:"productCode,omitempty"`
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
	App         *string `json:"app,omitempty" xml:"app,omitempty"`
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
	DestLogstore *string `json:"destLogstore,omitempty" xml:"destLogstore,omitempty"`
	DestProject  *string `json:"destProject,omitempty" xml:"destProject,omitempty"`
	DestRegion   *string `json:"destRegion,omitempty" xml:"destRegion,omitempty"`
	DestTTL      *int32  `json:"destTTL,omitempty" xml:"destTTL,omitempty"`
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
	InstanceIds  []*string              `json:"instanceIds,omitempty" xml:"instanceIds,omitempty" type:"Repeated"`
	Regions      []*string              `json:"regions,omitempty" xml:"regions,omitempty" type:"Repeated"`
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
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpsertCollectionPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

/**
 * Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
 *
 * @param headers map
 * @param runtime runtime options for this request RuntimeOptions
 * @return ApplyConfigToMachineGroupResponse
 */
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

/**
 * Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
 *
 * @return ApplyConfigToMachineGroupResponse
 */
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

/**
 * ### Usage notes
 * *   Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
 * *   You can create up to 30 consumer groups for a Logstore.
 * *   Simple Log Service provides examples of both regular log consumption and consumer group-based log consumption by using Simple Log Service SDKs for Java. For more information, see [Consume log data](~~120035~~) and [Use consumer groups to consume data](~~28998~~).
 *
 * @param request CreateConsumerGroupRequest
 * @param headers map
 * @param runtime runtime options for this request RuntimeOptions
 * @return CreateConsumerGroupResponse
 */
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

/**
 * ### Usage notes
 * *   Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
 * *   You can create up to 30 consumer groups for a Logstore.
 * *   Simple Log Service provides examples of both regular log consumption and consumer group-based log consumption by using Simple Log Service SDKs for Java. For more information, see [Consume log data](~~120035~~) and [Use consumer groups to consume data](~~28998~~).
 *
 * @param request CreateConsumerGroupRequest
 * @return CreateConsumerGroupResponse
 */
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

/**
 * Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
 *
 * @param request CreateDomainRequest
 * @param headers map
 * @param runtime runtime options for this request RuntimeOptions
 * @return CreateDomainResponse
 */
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

/**
 * Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
 *
 * @param request CreateDomainRequest
 * @return CreateDomainResponse
 */
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

/**
 * Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
 *
 * @param request CreateIndexRequest
 * @param headers map
 * @param runtime runtime options for this request RuntimeOptions
 * @return CreateIndexResponse
 */
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

/**
 * Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
 *
 * @param request CreateIndexRequest
 * @return CreateIndexResponse
 */
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

/**
 * ### Usage notes
 * Host consists of a project name and a Log Service endpoint. You must specify a project in Host.
 *
 * @param request CreateLogStoreRequest
 * @param headers map
 * @param runtime runtime options for this request RuntimeOptions
 * @return CreateLogStoreResponse
 */
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

/**
 * ### Usage notes
 * Host consists of a project name and a Log Service endpoint. You must specify a project in Host.
 *
 * @param request CreateLogStoreRequest
 * @return CreateLogStoreResponse
 */
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

/**
 * Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
 *
 * @param request CreateLoggingRequest
 * @param headers map
 * @param runtime runtime options for this request RuntimeOptions
 * @return CreateLoggingResponse
 */
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

/**
 * Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
 *
 * @param request CreateLoggingRequest
 * @return CreateLoggingResponse
 */
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

/**
 * ### Usage notes
 * Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
 *
 * @param request CreateMachineGroupRequest
 * @param headers map
 * @param runtime runtime options for this request RuntimeOptions
 * @return CreateMachineGroupResponse
 */
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

/**
 * ### Usage notes
 * Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
 *
 * @param request CreateMachineGroupRequest
 * @return CreateMachineGroupResponse
 */
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

/**
 * Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
 *
 * @param request CreateOssExternalStoreRequest
 * @param headers map
 * @param runtime runtime options for this request RuntimeOptions
 * @return CreateOssExternalStoreResponse
 */
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

/**
 * Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
 *
 * @param request CreateOssExternalStoreRequest
 * @return CreateOssExternalStoreResponse
 */
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

/**
 * Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
 *
 * @param request CreateRdsExternalStoreRequest
 * @param headers map
 * @param runtime runtime options for this request RuntimeOptions
 * @return CreateRdsExternalStoreResponse
 */
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

/**
 * Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
 *
 * @param request CreateRdsExternalStoreRequest
 * @return CreateRdsExternalStoreResponse
 */
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

/**
 * ### Usage notes
 * Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
 *
 * @param request CreateSavedSearchRequest
 * @param headers map
 * @param runtime runtime options for this request RuntimeOptions
 * @return CreateSavedSearchResponse
 */
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

/**
 * ### Usage notes
 * Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
 *
 * @param request CreateSavedSearchRequest
 * @return CreateSavedSearchResponse
 */
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

/**
 * Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
 *
 * @param headers map
 * @param runtime runtime options for this request RuntimeOptions
 * @return DeleteConsumerGroupResponse
 */
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

/**
 * Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
 *
 * @return DeleteConsumerGroupResponse
 */
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

/**
 * Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
 *
 * @param headers map
 * @param runtime runtime options for this request RuntimeOptions
 * @return DeleteDomainResponse
 */
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

/**
 * Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
 *
 * @return DeleteDomainResponse
 */
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

/**
 * Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
 *
 * @param headers map
 * @param runtime runtime options for this request RuntimeOptions
 * @return DeleteExternalStoreResponse
 */
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

/**
 * Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
 *
 * @return DeleteExternalStoreResponse
 */
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

/**
 * ### Usage notes
 * Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
 *
 * @param headers map
 * @param runtime runtime options for this request RuntimeOptions
 * @return DeleteIndexResponse
 */
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

/**
 * ### Usage notes
 * Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
 *
 * @return DeleteIndexResponse
 */
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

/**
 * ### Usage notes
 * Host consists of a project name and a Log Service endpoint. You must specify a project in Host.
 *
 * @param headers map
 * @param runtime runtime options for this request RuntimeOptions
 * @return DeleteLogStoreResponse
 */
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

/**
 * ### Usage notes
 * Host consists of a project name and a Log Service endpoint. You must specify a project in Host.
 *
 * @return DeleteLogStoreResponse
 */
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

/**
 * ### Usage notes
 * Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
 *
 * @param headers map
 * @param runtime runtime options for this request RuntimeOptions
 * @return DeleteMachineGroupResponse
 */
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

/**
 * ### Usage notes
 * Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
 *
 * @return DeleteMachineGroupResponse
 */
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

/**
 * ### Usage notes
 * Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
 *
 * @param headers map
 * @param runtime runtime options for this request RuntimeOptions
 * @return DeleteProjectPolicyResponse
 */
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

/**
 * ### Usage notes
 * Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
 *
 * @return DeleteProjectPolicyResponse
 */
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

/**
 * ### Usage notes
 * Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
 *
 * @param headers map
 * @param runtime runtime options for this request RuntimeOptions
 * @return DeleteSavedSearchResponse
 */
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

/**
 * ### Usage notes
 * Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
 *
 * @return DeleteSavedSearchResponse
 */
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

/**
 * Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
 *
 * @param headers map
 * @param runtime runtime options for this request RuntimeOptions
 * @return DeleteShipperResponse
 */
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

/**
 * Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
 *
 * @return DeleteShipperResponse
 */
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

/**
 * Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
 *
 * @param headers map
 * @param runtime runtime options for this request RuntimeOptions
 * @return GetAppliedConfigsResponse
 */
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

/**
 * Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
 *
 * @return GetAppliedConfigsResponse
 */
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

/**
 * Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
 *
 * @param headers map
 * @param runtime runtime options for this request RuntimeOptions
 * @return GetAppliedMachineGroupsResponse
 */
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

/**
 * Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
 *
 * @return GetAppliedMachineGroupsResponse
 */
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

/**
 * Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
 *
 * @param request GetCheckPointRequest
 * @param headers map
 * @param runtime runtime options for this request RuntimeOptions
 * @return GetCheckPointResponse
 */
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

/**
 * Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
 *
 * @param request GetCheckPointRequest
 * @return GetCheckPointResponse
 */
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

/**
 * *   You can specify a log as the start log. The time range of a contextual query is one day before and one day after the generation time of the start log.
 * *   Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
 *
 * @param request GetContextLogsRequest
 * @param headers map
 * @param runtime runtime options for this request RuntimeOptions
 * @return GetContextLogsResponse
 */
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

/**
 * *   You can specify a log as the start log. The time range of a contextual query is one day before and one day after the generation time of the start log.
 * *   Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
 *
 * @param request GetContextLogsRequest
 * @return GetContextLogsResponse
 */
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

/**
 * ### Usage notes
 * *   Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
 * *   The following content describes the relationships among a cursor, project, Logstore, and shard:
 *     *   A project can have multiple Logstores.
 *     *   A Logstore can have multiple shards.
 *     *   You can use a cursor to obtain a log in a shard.
 *
 * @param request GetCursorRequest
 * @param headers map
 * @param runtime runtime options for this request RuntimeOptions
 * @return GetCursorResponse
 */
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

/**
 * ### Usage notes
 * *   Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
 * *   The following content describes the relationships among a cursor, project, Logstore, and shard:
 *     *   A project can have multiple Logstores.
 *     *   A Logstore can have multiple shards.
 *     *   You can use a cursor to obtain a log in a shard.
 *
 * @param request GetCursorRequest
 * @return GetCursorResponse
 */
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

/**
 * *   The supported data sources of external stores include Object Storage Service (OSS) buckets and ApsaraDB RDS for MySQL databases in a virtual private cloud (VPC).
 * *   Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
 *
 * @param headers map
 * @param runtime runtime options for this request RuntimeOptions
 * @return GetExternalStoreResponse
 */
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

/**
 * *   The supported data sources of external stores include Object Storage Service (OSS) buckets and ApsaraDB RDS for MySQL databases in a virtual private cloud (VPC).
 * *   Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
 *
 * @return GetExternalStoreResponse
 */
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

/**
 * ### Usage notes
 * *   Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
 * *   The time range is evenly divided into subintervals in the responses. If the time range that is specified in the request remains unchanged, the subintervals in the responses also remain unchanged.
 * *   If the number of logs in a Logstore significantly changes, Simple Log Service cannot predict the number of times that you must call this operation to obtain the complete result. In this case, you must check the value of the progress parameter in the response of each request and determine whether to call this operation one more time to obtain the complete result. Each time you call this operation, the same number of charge units (CUs) are consumed.
 * *   After a log is written to a Logstore, you can call the GetHistograms or GetLogs operation to query the log after a short latency. The latency of a query varies based on the type of the log. Simple Log Service classifies logs into the following types based on the log time:
 *     *   Real-time data: The difference between the time record in a log and the current time on Simple Log Service is within the interval (-180 seconds,900 seconds]. For example, if a log was generated at 12:03:00, September 25, 2014 (UTC) and Simple Log Service received the log at 12:05:00, September 25, 2014 (UTC), Simple Log Service processes the log as real-time data. After real-time data is written to a Logstore, the data can be queried with a maximum latency of 3 seconds.
 *     *   Historical data: The difference between the time record in a log and the current time on Simple Log Service is within the interval \\[-604,800 seconds,-180 seconds). For example, if a log was generated at 12:00:00, September 25, 2014 (UTC) and Simple Log Service received the log at 12:05:00, September 25, 2014 (UTC), Simple Log Service processes the log as historical data. This type of log is usually generated in data backfill scenarios.
 *     After real-time data is written to a Logstore, the data can be queried with a maximum latency of 3 seconds. For 99.9% of queries, the latency is no more than 1 second.
 * > Simple Log Service calculates the difference between the log time that is specified by the \\__time\\_\\_ field and the receiving time that is specified by the \\__tag\\_\\_:\\__receive_time\\_\\_ field for each log. The receiving time indicates the time at which Simple Log Service receives the log. If the difference is within the interval (-180 seconds,900 seconds], Simple Log Service processes the log as real-time data. If the difference is within the interval \\[-604,800 seconds,-180 seconds), Simple Log Service processes the log as historical data.
 * *   Simple Log Service provides examples on how to call the GetHistograms operation by using Simple Log Service SDK for Java. For more information, see [Use GetHistograms to query the distribution of logs](~~462234~~).
 *
 * @param request GetHistogramsRequest
 * @param headers map
 * @param runtime runtime options for this request RuntimeOptions
 * @return GetHistogramsResponse
 */
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

/**
 * ### Usage notes
 * *   Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
 * *   The time range is evenly divided into subintervals in the responses. If the time range that is specified in the request remains unchanged, the subintervals in the responses also remain unchanged.
 * *   If the number of logs in a Logstore significantly changes, Simple Log Service cannot predict the number of times that you must call this operation to obtain the complete result. In this case, you must check the value of the progress parameter in the response of each request and determine whether to call this operation one more time to obtain the complete result. Each time you call this operation, the same number of charge units (CUs) are consumed.
 * *   After a log is written to a Logstore, you can call the GetHistograms or GetLogs operation to query the log after a short latency. The latency of a query varies based on the type of the log. Simple Log Service classifies logs into the following types based on the log time:
 *     *   Real-time data: The difference between the time record in a log and the current time on Simple Log Service is within the interval (-180 seconds,900 seconds]. For example, if a log was generated at 12:03:00, September 25, 2014 (UTC) and Simple Log Service received the log at 12:05:00, September 25, 2014 (UTC), Simple Log Service processes the log as real-time data. After real-time data is written to a Logstore, the data can be queried with a maximum latency of 3 seconds.
 *     *   Historical data: The difference between the time record in a log and the current time on Simple Log Service is within the interval \\[-604,800 seconds,-180 seconds). For example, if a log was generated at 12:00:00, September 25, 2014 (UTC) and Simple Log Service received the log at 12:05:00, September 25, 2014 (UTC), Simple Log Service processes the log as historical data. This type of log is usually generated in data backfill scenarios.
 *     After real-time data is written to a Logstore, the data can be queried with a maximum latency of 3 seconds. For 99.9% of queries, the latency is no more than 1 second.
 * > Simple Log Service calculates the difference between the log time that is specified by the \\__time\\_\\_ field and the receiving time that is specified by the \\__tag\\_\\_:\\__receive_time\\_\\_ field for each log. The receiving time indicates the time at which Simple Log Service receives the log. If the difference is within the interval (-180 seconds,900 seconds], Simple Log Service processes the log as real-time data. If the difference is within the interval \\[-604,800 seconds,-180 seconds), Simple Log Service processes the log as historical data.
 * *   Simple Log Service provides examples on how to call the GetHistograms operation by using Simple Log Service SDK for Java. For more information, see [Use GetHistograms to query the distribution of logs](~~462234~~).
 *
 * @param request GetHistogramsRequest
 * @return GetHistogramsResponse
 */
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

/**
 * ### Usage notes
 * Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
 *
 * @param headers map
 * @param runtime runtime options for this request RuntimeOptions
 * @return GetIndexResponse
 */
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

/**
 * ### Usage notes
 * Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
 *
 * @return GetIndexResponse
 */
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

/**
 * ### Usage notes
 * Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
 *
 * @param headers map
 * @param runtime runtime options for this request RuntimeOptions
 * @return GetLogStoreResponse
 */
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

/**
 * ### Usage notes
 * Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
 *
 * @return GetLogStoreResponse
 */
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

/**
 * Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
 *
 * @param headers map
 * @param runtime runtime options for this request RuntimeOptions
 * @return GetLoggingResponse
 */
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

/**
 * Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
 *
 * @return GetLoggingResponse
 */
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

/**
 * ### Usage notes
 * > Simple Log Service allows you to create a Scheduled SQL job. For more information, see [Create a Scheduled SQL job](~~286457~~).
 * *   Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
 * *   If the number of logs in a Logstore significantly changes, Simple Log Service cannot forecast the number of times that you must call this operation to obtain the complete result. In this case, you must check the value of the x-log-progress parameter in the response of each request and determine whether to call this operation one more time to obtain the complete result. Each time you call this operation, the same number of charge units (CUs) are consumed.
 * *   After a log is written to a Logstore, you can call the GetHistograms or GetLogs operation to query the log after a short latency. The latency of a query varies based on the type of the log. Simple Log Service classifies logs into the following types based on the log time:
 *         Real-time data: The difference between the time record in a log and the current time on Simple Log Service is within the interval (-180 seconds,900 seconds]. For example, if a log was generated at 12:03:00, September 25, 2014 (UTC) and Simple Log Service received the log at 12:05:00, September 25, 2014 (UTC), Simple Log Service processes the log as real-time data. This type of log is usually generated in common scenarios.
 *     *   Historical data: The difference between the time record in a log and the current time on Simple Log Service is within the interval \\[-604,800 seconds,-180 seconds). For example, if a log was generated at 12:00:00, September 25, 2014 (UTC) and Simple Log Service received the log at 12:05:00, September 25, 2014 (UTC), Simple Log Service processes the log as historical data. This type of log is usually generated in data backfill scenarios.
 *     After real-time data is written to a Logstore, the data can be queried with a maximum latency of 3 seconds. For 99.9% of queries, the latency is no more than 1 second.
 * > Simple Log Service calculates the difference between the log time that is specified by the \\__time\\_\\_ field and the receiving time that is specified by the \\__tag\\_\\_:**receive_time** field for each log. The receiving time indicates the time at which Simple Log Service receives the log. If the difference is within the interval (-180 seconds,900 seconds], Simple Log Service processes the log as real-time data. If the difference is within the interval \\[-604,800 seconds,-180 seconds), Simple Log Service processes the log as historical data.
 * *   Simple Log Service provides examples on how to call the GetLogs operation by using Simple Log Service SDK for Java and Simple Log Service SDK for Python. For more information, see [Examples of calling the GetLogs operation by using Simple Log Service SDK for Java](~~407683~~) and [Examples of calling the GetLogs operation by using Simple Log Service SDK for Python](~~407684~~).
 *
 * @param request GetLogsRequest
 * @param headers map
 * @param runtime runtime options for this request RuntimeOptions
 * @return GetLogsResponse
 */
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

/**
 * ### Usage notes
 * > Simple Log Service allows you to create a Scheduled SQL job. For more information, see [Create a Scheduled SQL job](~~286457~~).
 * *   Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
 * *   If the number of logs in a Logstore significantly changes, Simple Log Service cannot forecast the number of times that you must call this operation to obtain the complete result. In this case, you must check the value of the x-log-progress parameter in the response of each request and determine whether to call this operation one more time to obtain the complete result. Each time you call this operation, the same number of charge units (CUs) are consumed.
 * *   After a log is written to a Logstore, you can call the GetHistograms or GetLogs operation to query the log after a short latency. The latency of a query varies based on the type of the log. Simple Log Service classifies logs into the following types based on the log time:
 *         Real-time data: The difference between the time record in a log and the current time on Simple Log Service is within the interval (-180 seconds,900 seconds]. For example, if a log was generated at 12:03:00, September 25, 2014 (UTC) and Simple Log Service received the log at 12:05:00, September 25, 2014 (UTC), Simple Log Service processes the log as real-time data. This type of log is usually generated in common scenarios.
 *     *   Historical data: The difference between the time record in a log and the current time on Simple Log Service is within the interval \\[-604,800 seconds,-180 seconds). For example, if a log was generated at 12:00:00, September 25, 2014 (UTC) and Simple Log Service received the log at 12:05:00, September 25, 2014 (UTC), Simple Log Service processes the log as historical data. This type of log is usually generated in data backfill scenarios.
 *     After real-time data is written to a Logstore, the data can be queried with a maximum latency of 3 seconds. For 99.9% of queries, the latency is no more than 1 second.
 * > Simple Log Service calculates the difference between the log time that is specified by the \\__time\\_\\_ field and the receiving time that is specified by the \\__tag\\_\\_:**receive_time** field for each log. The receiving time indicates the time at which Simple Log Service receives the log. If the difference is within the interval (-180 seconds,900 seconds], Simple Log Service processes the log as real-time data. If the difference is within the interval \\[-604,800 seconds,-180 seconds), Simple Log Service processes the log as historical data.
 * *   Simple Log Service provides examples on how to call the GetLogs operation by using Simple Log Service SDK for Java and Simple Log Service SDK for Python. For more information, see [Examples of calling the GetLogs operation by using Simple Log Service SDK for Java](~~407683~~) and [Examples of calling the GetLogs operation by using Simple Log Service SDK for Python](~~407684~~).
 *
 * @param request GetLogsRequest
 * @return GetLogsResponse
 */
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

/**
 * *   Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
 * *   If the number of logs in a Logstore significantly changes, Simple Log Service cannot predict the number of times you must call this API operation to obtain a complete result. In this case, you must check the value of the progress parameter in the response of each request and determine whether to call this operation again to obtain the complete result. Each time you call this operation, the same number of charge units (CUs) are consumed.
 * *   After a log is written to a Logstore, you can call the GetHistograms or GetLogs operation to query the log. The latency of the query varies based on the type of the log. Simple Log Service classifies logs into the following types based on log timestamps:
 * 1.  1.  Real-time data: The difference between the time record in the log and the current server time is within the interval (-180 seconds,900 seconds]. For example, if a log was generated at 12:03:00, September 25, 2014 (UTC) and the server received the log at 12:05:00, September 25, 2014 (UTC), the server processes the log as real-time data. This type of log is usually generated in common scenarios.
 * 2.  2.  Historical data: The difference between the time record in the log and the current server time is within the interval \\[-604,800 seconds,-180 seconds). For example, if a log was generated at 12:00:00, September 25, 2014 (UTC) and the server received the log at 12:05:00, September 25, 2014 (UTC), the server processes the log as historical data. This type of log is usually generated in data backfill scenarios. After real-time data is written to a Logstore, the data can be queried with a maximum latency of 3 seconds. For 99.9% of queries, the latency is no more than 1 second.
 *
 * @param request GetLogsV2Request
 * @param headers GetLogsV2Headers
 * @param runtime runtime options for this request RuntimeOptions
 * @return GetLogsV2Response
 */
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

	if !tea.BoolValue(util.IsUnset(request.Shard)) {
		body["shard"] = request.Shard
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

/**
 * *   Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
 * *   If the number of logs in a Logstore significantly changes, Simple Log Service cannot predict the number of times you must call this API operation to obtain a complete result. In this case, you must check the value of the progress parameter in the response of each request and determine whether to call this operation again to obtain the complete result. Each time you call this operation, the same number of charge units (CUs) are consumed.
 * *   After a log is written to a Logstore, you can call the GetHistograms or GetLogs operation to query the log. The latency of the query varies based on the type of the log. Simple Log Service classifies logs into the following types based on log timestamps:
 * 1.  1.  Real-time data: The difference between the time record in the log and the current server time is within the interval (-180 seconds,900 seconds]. For example, if a log was generated at 12:03:00, September 25, 2014 (UTC) and the server received the log at 12:05:00, September 25, 2014 (UTC), the server processes the log as real-time data. This type of log is usually generated in common scenarios.
 * 2.  2.  Historical data: The difference between the time record in the log and the current server time is within the interval \\[-604,800 seconds,-180 seconds). For example, if a log was generated at 12:00:00, September 25, 2014 (UTC) and the server received the log at 12:05:00, September 25, 2014 (UTC), the server processes the log as historical data. This type of log is usually generated in data backfill scenarios. After real-time data is written to a Logstore, the data can be queried with a maximum latency of 3 seconds. For 99.9% of queries, the latency is no more than 1 second.
 *
 * @param request GetLogsV2Request
 * @return GetLogsV2Response
 */
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

func (client *Client) GetMLServiceResultsWithOptions(serviceName *string, request *GetMLServiceResultsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetMLServiceResultsResponse, _err error) {
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

/**
 * Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
 *
 * @param headers map
 * @param runtime runtime options for this request RuntimeOptions
 * @return GetMachineGroupResponse
 */
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

/**
 * Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
 *
 * @return GetMachineGroupResponse
 */
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

/**
 * ### Usage notes
 * Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
 *
 * @param headers map
 * @param runtime runtime options for this request RuntimeOptions
 * @return GetProjectResponse
 */
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

/**
 * ### Usage notes
 * Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
 *
 * @return GetProjectResponse
 */
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

/**
 * ### Usage notes
 * *   You can use the query parameter to specify a standard SQL statement.
 * *   You must specify a project in the domain name of the request.
 * *   You must specify a Logstore in the FROM clause of the SQL statement. A Logstore can be used as an SQL table.
 * *   You must specify a time range in the SQL statement by using the \\__date\\_\\_ parameter or \\__time\\_\\_ parameter. The value of the \\__date\\_\\_ parameter is a timestamp, and the value of the \\__time\\_\\_ parameter is an integer. The unit of the \\__time\\_\\_ parameter is seconds.
 * *   Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
 *
 * @param request GetProjectLogsRequest
 * @param headers map
 * @param runtime runtime options for this request RuntimeOptions
 * @return GetProjectLogsResponse
 */
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

/**
 * ### Usage notes
 * *   You can use the query parameter to specify a standard SQL statement.
 * *   You must specify a project in the domain name of the request.
 * *   You must specify a Logstore in the FROM clause of the SQL statement. A Logstore can be used as an SQL table.
 * *   You must specify a time range in the SQL statement by using the \\__date\\_\\_ parameter or \\__time\\_\\_ parameter. The value of the \\__date\\_\\_ parameter is a timestamp, and the value of the \\__time\\_\\_ parameter is an integer. The unit of the \\__time\\_\\_ parameter is seconds.
 * *   Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
 *
 * @param request GetProjectLogsRequest
 * @return GetProjectLogsResponse
 */
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

/**
 * ### Usage notes
 * Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
 *
 * @param headers map
 * @param runtime runtime options for this request RuntimeOptions
 * @return GetProjectPolicyResponse
 */
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

/**
 * ### Usage notes
 * Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
 *
 * @return GetProjectPolicyResponse
 */
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

/**
 * ### Usage notes
 * Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
 *
 * @param headers map
 * @param runtime runtime options for this request RuntimeOptions
 * @return GetSavedSearchResponse
 */
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

/**
 * ### Usage notes
 * Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
 *
 * @return GetSavedSearchResponse
 */
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

/**
 * Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
 *
 * @param request GetShipperStatusRequest
 * @param headers map
 * @param runtime runtime options for this request RuntimeOptions
 * @return GetShipperStatusResponse
 */
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

/**
 * Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
 *
 * @param request GetShipperStatusRequest
 * @return GetShipperStatusResponse
 */
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

/**
 * ### Usage notes
 * Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
 *
 * @param headers map
 * @param runtime runtime options for this request RuntimeOptions
 * @return ListConsumerGroupResponse
 */
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

/**
 * ### Usage notes
 * Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
 *
 * @return ListConsumerGroupResponse
 */
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

/**
 * *   Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
 * *   Only one custom domain name can be bound to each project.
 *
 * @param request ListDomainsRequest
 * @param headers map
 * @param runtime runtime options for this request RuntimeOptions
 * @return ListDomainsResponse
 */
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

/**
 * *   Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
 * *   Only one custom domain name can be bound to each project.
 *
 * @param request ListDomainsRequest
 * @return ListDomainsResponse
 */
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

/**
 * Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
 *
 * @param request ListExternalStoreRequest
 * @param headers map
 * @param runtime runtime options for this request RuntimeOptions
 * @return ListExternalStoreResponse
 */
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

/**
 * Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
 *
 * @param request ListExternalStoreRequest
 * @return ListExternalStoreResponse
 */
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

/**
 * ### Usage notes
 * Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
 *
 * @param request ListLogStoresRequest
 * @param headers map
 * @param runtime runtime options for this request RuntimeOptions
 * @return ListLogStoresResponse
 */
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

/**
 * ### Usage notes
 * Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
 *
 * @param request ListLogStoresRequest
 * @return ListLogStoresResponse
 */
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

/**
 * Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
 *
 * @param request ListMachineGroupRequest
 * @param headers map
 * @param runtime runtime options for this request RuntimeOptions
 * @return ListMachineGroupResponse
 */
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

/**
 * Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
 *
 * @param request ListMachineGroupRequest
 * @return ListMachineGroupResponse
 */
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

/**
 * Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
 *
 * @param request ListMachinesRequest
 * @param headers map
 * @param runtime runtime options for this request RuntimeOptions
 * @return ListMachinesResponse
 */
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

/**
 * Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
 *
 * @param request ListMachinesRequest
 * @return ListMachinesResponse
 */
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

/**
 * ### Usage notes
 * Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
 *
 * @param request ListProjectRequest
 * @param headers map
 * @param runtime runtime options for this request RuntimeOptions
 * @return ListProjectResponse
 */
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

/**
 * ### Usage notes
 * Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
 *
 * @param request ListProjectRequest
 * @return ListProjectResponse
 */
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

/**
 * ### Usage notes
 * Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
 *
 * @param request ListSavedSearchRequest
 * @param headers map
 * @param runtime runtime options for this request RuntimeOptions
 * @return ListSavedSearchResponse
 */
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

/**
 * ### Usage notes
 * Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
 *
 * @param request ListSavedSearchRequest
 * @return ListSavedSearchResponse
 */
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

/**
 * Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
 *
 * @param headers map
 * @param runtime runtime options for this request RuntimeOptions
 * @return ListShipperResponse
 */
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

/**
 * Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
 *
 * @return ListShipperResponse
 */
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

/**
 * ### Usage notes
 * Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
 *
 * @param tmpReq ListTagResourcesRequest
 * @param headers map
 * @param runtime runtime options for this request RuntimeOptions
 * @return ListTagResourcesResponse
 */
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

/**
 * ### Usage notes
 * Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
 *
 * @param request ListTagResourcesRequest
 * @return ListTagResourcesResponse
 */
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

/**
 * ### Usage notes
 * *   Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
 * *   Alibaba Cloud Simple Log Service allows you to configure a project policy to authorize other users to access the specified Log Service resources.
 *     *   You must configure a project policy based on policy syntax. Before you configure a project policy, you must be familiar with the Action, Resource, and Condition parameters. For more information, see [RAM](~~128139~~).
 *     *   If you set the Principal element to an asterisk (\\*) and do not configure the Condition element when you configure a project policy, the policy applies to all users except for the project owner. If you set the Principal element to an asterisk (\\*) and configure the Condition element when you configure a project policy, the policy applies to all users including the project owner.
 *     *   You can configure multiple project policies for a project. The total size of the policies cannot exceed 16 KB.
 *
 * @param request PutProjectPolicyRequest
 * @param headers map
 * @param runtime runtime options for this request RuntimeOptions
 * @return PutProjectPolicyResponse
 */
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

/**
 * ### Usage notes
 * *   Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
 * *   Alibaba Cloud Simple Log Service allows you to configure a project policy to authorize other users to access the specified Log Service resources.
 *     *   You must configure a project policy based on policy syntax. Before you configure a project policy, you must be familiar with the Action, Resource, and Condition parameters. For more information, see [RAM](~~128139~~).
 *     *   If you set the Principal element to an asterisk (\\*) and do not configure the Condition element when you configure a project policy, the policy applies to all users except for the project owner. If you set the Principal element to an asterisk (\\*) and configure the Condition element when you configure a project policy, the policy applies to all users including the project owner.
 *     *   You can configure multiple project policies for a project. The total size of the policies cannot exceed 16 KB.
 *
 * @param request PutProjectPolicyRequest
 * @return PutProjectPolicyResponse
 */
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

/**
 * @deprecated
 *
 * @param request QueryMLServiceResultsRequest
 * @param headers map
 * @param runtime runtime options for this request RuntimeOptions
 * @return QueryMLServiceResultsResponse
 */
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

/**
 * @deprecated
 *
 * @param request QueryMLServiceResultsRequest
 * @return QueryMLServiceResultsResponse
 */
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

/**
 * Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
 *
 * @param headers map
 * @param runtime runtime options for this request RuntimeOptions
 * @return RemoveConfigFromMachineGroupResponse
 */
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

/**
 * Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
 *
 * @return RemoveConfigFromMachineGroupResponse
 */
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

/**
 * ### Usage notes
 * *   Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
 * *   Each shard has an MD5 hash range, and each range is a left-closed, right-open interval. The interval is in the `[BeginKey,EndKey)` format. A shard can be in the readwrite or readonly state. You can split a shard and merge shards. For more information, see [Shard](~~28976~~).
 *
 * @param request SplitShardRequest
 * @param headers map
 * @param runtime runtime options for this request RuntimeOptions
 * @return SplitShardResponse
 */
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

/**
 * ### Usage notes
 * *   Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
 * *   Each shard has an MD5 hash range, and each range is a left-closed, right-open interval. The interval is in the `[BeginKey,EndKey)` format. A shard can be in the readwrite or readonly state. You can split a shard and merge shards. For more information, see [Shard](~~28976~~).
 *
 * @param request SplitShardRequest
 * @return SplitShardResponse
 */
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

/**
 * ### Usage notes
 * Host consists of a project name and a Log Service endpoint. You must specify a project in Host.
 *
 * @param request TagResourcesRequest
 * @param headers map
 * @param runtime runtime options for this request RuntimeOptions
 * @return TagResourcesResponse
 */
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

/**
 * ### Usage notes
 * Host consists of a project name and a Log Service endpoint. You must specify a project in Host.
 *
 * @param request TagResourcesRequest
 * @return TagResourcesResponse
 */
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

/**
 * ### Usage notes
 * Host consists of a project name and a Log Service endpoint. You must specify a project in Host.
 *
 * @param request UntagResourcesRequest
 * @param headers map
 * @param runtime runtime options for this request RuntimeOptions
 * @return UntagResourcesResponse
 */
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

/**
 * ### Usage notes
 * Host consists of a project name and a Log Service endpoint. You must specify a project in Host.
 *
 * @param request UntagResourcesRequest
 * @return UntagResourcesResponse
 */
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

/**
 * ### Usage notes
 * Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
 *
 * @param request UpdateConsumerGroupRequest
 * @param headers map
 * @param runtime runtime options for this request RuntimeOptions
 * @return UpdateConsumerGroupResponse
 */
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

/**
 * ### Usage notes
 * Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
 *
 * @param request UpdateConsumerGroupRequest
 * @return UpdateConsumerGroupResponse
 */
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

/**
 * ### Usage notes
 * Host consists of a project name and a Log Service endpoint. You must specify a project in Host.
 *
 * @param request UpdateIndexRequest
 * @param headers map
 * @param runtime runtime options for this request RuntimeOptions
 * @return UpdateIndexResponse
 */
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

/**
 * ### Usage notes
 * Host consists of a project name and a Log Service endpoint. You must specify a project in Host.
 *
 * @param request UpdateIndexRequest
 * @return UpdateIndexResponse
 */
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

/**
 * ### Usage notes
 * *   Host consists of a project name and a Log Service endpoint. You must specify a project in Host.
 * *   You can call the UpdateLogStore operation to change only the time-to-live (TTL) attribute.
 *
 * @param request UpdateLogStoreRequest
 * @param headers map
 * @param runtime runtime options for this request RuntimeOptions
 * @return UpdateLogStoreResponse
 */
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

/**
 * ### Usage notes
 * *   Host consists of a project name and a Log Service endpoint. You must specify a project in Host.
 * *   You can call the UpdateLogStore operation to change only the time-to-live (TTL) attribute.
 *
 * @param request UpdateLogStoreRequest
 * @return UpdateLogStoreResponse
 */
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

/**
 * Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
 *
 * @param request UpdateLoggingRequest
 * @param headers map
 * @param runtime runtime options for this request RuntimeOptions
 * @return UpdateLoggingResponse
 */
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

/**
 * Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
 *
 * @param request UpdateLoggingRequest
 * @return UpdateLoggingResponse
 */
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

/**
 * Host consists of a project name and a Log Service endpoint. You must specify a project in Host.
 *
 * @param request UpdateMachineGroupRequest
 * @param headers map
 * @param runtime runtime options for this request RuntimeOptions
 * @return UpdateMachineGroupResponse
 */
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

/**
 * Host consists of a project name and a Log Service endpoint. You must specify a project in Host.
 *
 * @param request UpdateMachineGroupRequest
 * @return UpdateMachineGroupResponse
 */
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

/**
 * Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
 *
 * @param request UpdateMachineGroupMachineRequest
 * @param headers map
 * @param runtime runtime options for this request RuntimeOptions
 * @return UpdateMachineGroupMachineResponse
 */
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

/**
 * Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
 *
 * @param request UpdateMachineGroupMachineRequest
 * @return UpdateMachineGroupMachineResponse
 */
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

/**
 * Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
 *
 * @param request UpdateOssExternalStoreRequest
 * @param headers map
 * @param runtime runtime options for this request RuntimeOptions
 * @return UpdateOssExternalStoreResponse
 */
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

/**
 * Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
 *
 * @param request UpdateOssExternalStoreRequest
 * @return UpdateOssExternalStoreResponse
 */
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

/**
 * ### Usage notes
 * Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
 *
 * @param request UpdateProjectRequest
 * @param headers map
 * @param runtime runtime options for this request RuntimeOptions
 * @return UpdateProjectResponse
 */
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

/**
 * ### Usage notes
 * Host consists of a project name and a Simple Log Service endpoint. You must specify a project in Host.
 *
 * @param request UpdateProjectRequest
 * @return UpdateProjectResponse
 */
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

/**
 * Host consists of a project name and a Log Service endpoint. You must specify a project in Host.
 *
 * @param request UpdateRdsExternalStoreRequest
 * @param headers map
 * @param runtime runtime options for this request RuntimeOptions
 * @return UpdateRdsExternalStoreResponse
 */
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

/**
 * Host consists of a project name and a Log Service endpoint. You must specify a project in Host.
 *
 * @param request UpdateRdsExternalStoreRequest
 * @return UpdateRdsExternalStoreResponse
 */
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
