// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDIJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPagingInfo(v *GetDIJobResponseBodyPagingInfo) *GetDIJobResponseBody
	GetPagingInfo() *GetDIJobResponseBodyPagingInfo
	SetRequestId(v string) *GetDIJobResponseBody
	GetRequestId() *string
}

type GetDIJobResponseBody struct {
	// The pagination information.
	PagingInfo *GetDIJobResponseBodyPagingInfo `json:"PagingInfo,omitempty" xml:"PagingInfo,omitempty" type:"Struct"`
	// The request ID. You can use the ID to query logs and troubleshoot issues.
	//
	// example:
	//
	// C99E2BE6-9DEA-5C2E-8F51-1DDCFEADE490
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetDIJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDIJobResponseBody) GoString() string {
	return s.String()
}

func (s *GetDIJobResponseBody) GetPagingInfo() *GetDIJobResponseBodyPagingInfo {
	return s.PagingInfo
}

func (s *GetDIJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDIJobResponseBody) SetPagingInfo(v *GetDIJobResponseBodyPagingInfo) *GetDIJobResponseBody {
	s.PagingInfo = v
	return s
}

func (s *GetDIJobResponseBody) SetRequestId(v string) *GetDIJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDIJobResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetDIJobResponseBodyPagingInfo struct {
	// Deprecated
	//
	// This parameter is deprecated. Use the Id parameter instead.
	//
	// example:
	//
	// 32601
	DIJobId *string `json:"DIJobId,omitempty" xml:"DIJobId,omitempty"`
	// The description of the synchronization task.
	//
	// example:
	//
	// description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The properties of the destination.
	DestinationDataSourceSettings []*GetDIJobResponseBodyPagingInfoDestinationDataSourceSettings `json:"DestinationDataSourceSettings,omitempty" xml:"DestinationDataSourceSettings,omitempty" type:"Repeated"`
	// The destination type. Valid values: Hologres, OSS-HDFS, OSS, MaxCompute, LogHub, StarRocks, DataHub, AnalyticDB_For_MySQL, Kafka, Hive.
	//
	// example:
	//
	// Hologres
	DestinationDataSourceType *string `json:"DestinationDataSourceType,omitempty" xml:"DestinationDataSourceType,omitempty"`
	// The ID of the synchronization task.
	//
	// example:
	//
	// 32601
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the synchronization task.
	//
	// example:
	//
	// imp_ods_dms_det_dealer_info_df
	JobName *string `json:"JobName,omitempty" xml:"JobName,omitempty"`
	// The runtime settings.
	JobSettings *GetDIJobResponseBodyPagingInfoJobSettings `json:"JobSettings,omitempty" xml:"JobSettings,omitempty" type:"Struct"`
	// The status of the job.
	//
	// example:
	//
	// Running
	JobStatus *string `json:"JobStatus,omitempty" xml:"JobStatus,omitempty"`
	// 任务类型
	//
	// - DatabaseRealtimeMigration(整库实时):将源端多个库的多个表进行流同步，支持仅全量，仅增量，或全量+增量。
	//
	// - DatabaseOfflineMigration(整库离线):将源端多个库的多个表进行批同步，支持仅全量，仅增量，或全量+增量。
	//
	// - SingleTableRealtimeMigration(单表实时):将源端单个表进行流同步。
	//
	// example:
	//
	// DatabaseRealtimeMigration
	JobType *string `json:"JobType,omitempty" xml:"JobType,omitempty"`
	// The synchronization type. Valid values:
	//
	// 	- FullAndRealtimeIncremental: full synchronization and real-time incremental synchronization of data in an entire database
	//
	// 	- RealtimeIncremental: real-time incremental synchronization of data in a single table
	//
	// 	- Full: full batch synchronization of data in an entire database
	//
	// 	- OfflineIncremental: batch incremental synchronization of data in an entire database
	//
	// 	- FullAndOfflineIncremental: full synchronization and batch incremental synchronization of data in an entire database
	//
	// example:
	//
	// FullAndRealtimeIncremental
	MigrationType *string `json:"MigrationType,omitempty" xml:"MigrationType,omitempty"`
	// The DataWorks workspace ID. You can log on to the [DataWorks console](https://workbench.data.aliyun.com/console) and go to the Workspace page to query the ID.
	//
	// This parameter indicates the DataWorks workspace to which the API operation is applied.
	//
	// example:
	//
	// 98330
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The resource settings.
	ResourceSettings *GetDIJobResponseBodyPagingInfoResourceSettings `json:"ResourceSettings,omitempty" xml:"ResourceSettings,omitempty" type:"Struct"`
	// The settings of the source. Only a single source is supported.
	SourceDataSourceSettings []*GetDIJobResponseBodyPagingInfoSourceDataSourceSettings `json:"SourceDataSourceSettings,omitempty" xml:"SourceDataSourceSettings,omitempty" type:"Repeated"`
	// The source type. Valid values: PolarDB, MySQL, Kafka, LogHub, Hologres, Oracle, OceanBase, MongoDB, RedShift, Hive, SQLServer, Doris, ClickHouse.
	//
	// example:
	//
	// Mysql
	SourceDataSourceType *string `json:"SourceDataSourceType,omitempty" xml:"SourceDataSourceType,omitempty"`
	// The list of mappings between rules used to select synchronization objects in the source and transformation rules applied to the selected synchronization objects. Each entry in the list displays a mapping between a rule used to select synchronization objects and a transformation rule applied to the selected synchronization objects.
	//
	// >  [ { "SourceObjectSelectionRules":[ { "ObjectType":"Database", "Action":"Include", "ExpressionType":"Exact", "Expression":"biz_db" }, { "ObjectType":"Schema", "Action":"Include", "ExpressionType":"Exact", "Expression":"s1" }, { "ObjectType":"Table", "Action":"Include", "ExpressionType":"Exact", "Expression":"table1" } ], "TransformationRuleNames":[ { "RuleName":"my_database_rename_rule", "RuleActionType":"Rename", "RuleTargetType":"Schema" } ] } ]
	TableMappings []*GetDIJobResponseBodyPagingInfoTableMappings `json:"TableMappings,omitempty" xml:"TableMappings,omitempty" type:"Repeated"`
	// The list of transformation rules that are applied to the synchronization objects selected from the source.
	//
	// >  [ { "RuleName":"my_database_rename_rule", "RuleActionType":"Rename", "RuleTargetType":"Schema", "RuleExpression":"{"expression":"${srcDatasoureName}_${srcDatabaseName}"}" } ]
	TransformationRules []*GetDIJobResponseBodyPagingInfoTransformationRules `json:"TransformationRules,omitempty" xml:"TransformationRules,omitempty" type:"Repeated"`
}

func (s GetDIJobResponseBodyPagingInfo) String() string {
	return dara.Prettify(s)
}

func (s GetDIJobResponseBodyPagingInfo) GoString() string {
	return s.String()
}

func (s *GetDIJobResponseBodyPagingInfo) GetDIJobId() *string {
	return s.DIJobId
}

func (s *GetDIJobResponseBodyPagingInfo) GetDescription() *string {
	return s.Description
}

func (s *GetDIJobResponseBodyPagingInfo) GetDestinationDataSourceSettings() []*GetDIJobResponseBodyPagingInfoDestinationDataSourceSettings {
	return s.DestinationDataSourceSettings
}

func (s *GetDIJobResponseBodyPagingInfo) GetDestinationDataSourceType() *string {
	return s.DestinationDataSourceType
}

func (s *GetDIJobResponseBodyPagingInfo) GetId() *int64 {
	return s.Id
}

func (s *GetDIJobResponseBodyPagingInfo) GetJobName() *string {
	return s.JobName
}

func (s *GetDIJobResponseBodyPagingInfo) GetJobSettings() *GetDIJobResponseBodyPagingInfoJobSettings {
	return s.JobSettings
}

func (s *GetDIJobResponseBodyPagingInfo) GetJobStatus() *string {
	return s.JobStatus
}

func (s *GetDIJobResponseBodyPagingInfo) GetJobType() *string {
	return s.JobType
}

func (s *GetDIJobResponseBodyPagingInfo) GetMigrationType() *string {
	return s.MigrationType
}

func (s *GetDIJobResponseBodyPagingInfo) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetDIJobResponseBodyPagingInfo) GetResourceSettings() *GetDIJobResponseBodyPagingInfoResourceSettings {
	return s.ResourceSettings
}

func (s *GetDIJobResponseBodyPagingInfo) GetSourceDataSourceSettings() []*GetDIJobResponseBodyPagingInfoSourceDataSourceSettings {
	return s.SourceDataSourceSettings
}

func (s *GetDIJobResponseBodyPagingInfo) GetSourceDataSourceType() *string {
	return s.SourceDataSourceType
}

func (s *GetDIJobResponseBodyPagingInfo) GetTableMappings() []*GetDIJobResponseBodyPagingInfoTableMappings {
	return s.TableMappings
}

func (s *GetDIJobResponseBodyPagingInfo) GetTransformationRules() []*GetDIJobResponseBodyPagingInfoTransformationRules {
	return s.TransformationRules
}

func (s *GetDIJobResponseBodyPagingInfo) SetDIJobId(v string) *GetDIJobResponseBodyPagingInfo {
	s.DIJobId = &v
	return s
}

func (s *GetDIJobResponseBodyPagingInfo) SetDescription(v string) *GetDIJobResponseBodyPagingInfo {
	s.Description = &v
	return s
}

func (s *GetDIJobResponseBodyPagingInfo) SetDestinationDataSourceSettings(v []*GetDIJobResponseBodyPagingInfoDestinationDataSourceSettings) *GetDIJobResponseBodyPagingInfo {
	s.DestinationDataSourceSettings = v
	return s
}

func (s *GetDIJobResponseBodyPagingInfo) SetDestinationDataSourceType(v string) *GetDIJobResponseBodyPagingInfo {
	s.DestinationDataSourceType = &v
	return s
}

func (s *GetDIJobResponseBodyPagingInfo) SetId(v int64) *GetDIJobResponseBodyPagingInfo {
	s.Id = &v
	return s
}

func (s *GetDIJobResponseBodyPagingInfo) SetJobName(v string) *GetDIJobResponseBodyPagingInfo {
	s.JobName = &v
	return s
}

func (s *GetDIJobResponseBodyPagingInfo) SetJobSettings(v *GetDIJobResponseBodyPagingInfoJobSettings) *GetDIJobResponseBodyPagingInfo {
	s.JobSettings = v
	return s
}

func (s *GetDIJobResponseBodyPagingInfo) SetJobStatus(v string) *GetDIJobResponseBodyPagingInfo {
	s.JobStatus = &v
	return s
}

func (s *GetDIJobResponseBodyPagingInfo) SetJobType(v string) *GetDIJobResponseBodyPagingInfo {
	s.JobType = &v
	return s
}

func (s *GetDIJobResponseBodyPagingInfo) SetMigrationType(v string) *GetDIJobResponseBodyPagingInfo {
	s.MigrationType = &v
	return s
}

func (s *GetDIJobResponseBodyPagingInfo) SetProjectId(v int64) *GetDIJobResponseBodyPagingInfo {
	s.ProjectId = &v
	return s
}

func (s *GetDIJobResponseBodyPagingInfo) SetResourceSettings(v *GetDIJobResponseBodyPagingInfoResourceSettings) *GetDIJobResponseBodyPagingInfo {
	s.ResourceSettings = v
	return s
}

func (s *GetDIJobResponseBodyPagingInfo) SetSourceDataSourceSettings(v []*GetDIJobResponseBodyPagingInfoSourceDataSourceSettings) *GetDIJobResponseBodyPagingInfo {
	s.SourceDataSourceSettings = v
	return s
}

func (s *GetDIJobResponseBodyPagingInfo) SetSourceDataSourceType(v string) *GetDIJobResponseBodyPagingInfo {
	s.SourceDataSourceType = &v
	return s
}

func (s *GetDIJobResponseBodyPagingInfo) SetTableMappings(v []*GetDIJobResponseBodyPagingInfoTableMappings) *GetDIJobResponseBodyPagingInfo {
	s.TableMappings = v
	return s
}

func (s *GetDIJobResponseBodyPagingInfo) SetTransformationRules(v []*GetDIJobResponseBodyPagingInfoTransformationRules) *GetDIJobResponseBodyPagingInfo {
	s.TransformationRules = v
	return s
}

func (s *GetDIJobResponseBodyPagingInfo) Validate() error {
	return dara.Validate(s)
}

type GetDIJobResponseBodyPagingInfoDestinationDataSourceSettings struct {
	// The name of the data source.
	//
	// example:
	//
	// dw_mysql
	DataSourceName *string `json:"DataSourceName,omitempty" xml:"DataSourceName,omitempty"`
}

func (s GetDIJobResponseBodyPagingInfoDestinationDataSourceSettings) String() string {
	return dara.Prettify(s)
}

func (s GetDIJobResponseBodyPagingInfoDestinationDataSourceSettings) GoString() string {
	return s.String()
}

func (s *GetDIJobResponseBodyPagingInfoDestinationDataSourceSettings) GetDataSourceName() *string {
	return s.DataSourceName
}

func (s *GetDIJobResponseBodyPagingInfoDestinationDataSourceSettings) SetDataSourceName(v string) *GetDIJobResponseBodyPagingInfoDestinationDataSourceSettings {
	s.DataSourceName = &v
	return s
}

func (s *GetDIJobResponseBodyPagingInfoDestinationDataSourceSettings) Validate() error {
	return dara.Validate(s)
}

type GetDIJobResponseBodyPagingInfoJobSettings struct {
	// The channel control settings for the synchronization task. You can configure special channel control settings for the following synchronization links: data synchronization between Hologres data sources and data synchronization from Hologres to Kafka.
	//
	// 1.  Holo2Kafka
	//
	// 	- Example: {"destinationChannelSettings":{"kafkaClientProperties":[{"key":"linger.ms","value":"100"}],"keyColumns":["col3"],"writeMode":"canal"}}
	//
	// 	- kafkaClientProperties: the parameters related to a Kafka producer, which are used when you write data to a Kafka data source.
	//
	// 	- keyColumns: the names of Kafka columns to which data is written.
	//
	// 	- writeMode: the writing format. Valid values: json and canal.
	//
	// 2.  Holo2Holo
	//
	// 	- Example: {"destinationChannelSettings":{"conflictMode":"replace","dynamicColumnAction":"replay","writeMode":"replay"}}
	//
	// 	- conflictMode: the policy used to handle a conflict that occurs during data writing to Hologres. Valid values: replace and ignore.
	//
	// 	- writeMode: the mode in which data is written to Hologres. Valid values: replay and insert.
	//
	// 	- dynamicColumnAction: the mode in which data is written to dynamic columns in a Hologres table. Valid values: replay, insert, and ignore.
	//
	// example:
	//
	// {"structInfo":"MANAGED","storageType":"TEXTFILE","writeMode":"APPEND","partitionColumns":[{"columnName":"pt","columnType":"STRING","comment":""}],"fieldDelimiter":""}
	ChannelSettings *string `json:"ChannelSettings,omitempty" xml:"ChannelSettings,omitempty"`
	// The data type mappings between source fields and destination fields.
	ColumnDataTypeSettings []*GetDIJobResponseBodyPagingInfoJobSettingsColumnDataTypeSettings `json:"ColumnDataTypeSettings,omitempty" xml:"ColumnDataTypeSettings,omitempty" type:"Repeated"`
	// The settings for periodic scheduling.
	CycleScheduleSettings *GetDIJobResponseBodyPagingInfoJobSettingsCycleScheduleSettings `json:"CycleScheduleSettings,omitempty" xml:"CycleScheduleSettings,omitempty" type:"Struct"`
	// The DDL operation types. Valid values:
	//
	// 	- RenameColumn
	//
	// 	- ModifyColumn
	//
	// 	- CreateTable
	//
	// 	- TruncateTable
	//
	// 	- DropTable
	//
	// 	- DropColumn
	//
	// 	- AddColumn
	DdlHandlingSettings []*GetDIJobResponseBodyPagingInfoJobSettingsDdlHandlingSettings `json:"DdlHandlingSettings,omitempty" xml:"DdlHandlingSettings,omitempty" type:"Repeated"`
	// The runtime settings.
	RuntimeSettings []*GetDIJobResponseBodyPagingInfoJobSettingsRuntimeSettings `json:"RuntimeSettings,omitempty" xml:"RuntimeSettings,omitempty" type:"Repeated"`
}

func (s GetDIJobResponseBodyPagingInfoJobSettings) String() string {
	return dara.Prettify(s)
}

func (s GetDIJobResponseBodyPagingInfoJobSettings) GoString() string {
	return s.String()
}

func (s *GetDIJobResponseBodyPagingInfoJobSettings) GetChannelSettings() *string {
	return s.ChannelSettings
}

func (s *GetDIJobResponseBodyPagingInfoJobSettings) GetColumnDataTypeSettings() []*GetDIJobResponseBodyPagingInfoJobSettingsColumnDataTypeSettings {
	return s.ColumnDataTypeSettings
}

func (s *GetDIJobResponseBodyPagingInfoJobSettings) GetCycleScheduleSettings() *GetDIJobResponseBodyPagingInfoJobSettingsCycleScheduleSettings {
	return s.CycleScheduleSettings
}

func (s *GetDIJobResponseBodyPagingInfoJobSettings) GetDdlHandlingSettings() []*GetDIJobResponseBodyPagingInfoJobSettingsDdlHandlingSettings {
	return s.DdlHandlingSettings
}

func (s *GetDIJobResponseBodyPagingInfoJobSettings) GetRuntimeSettings() []*GetDIJobResponseBodyPagingInfoJobSettingsRuntimeSettings {
	return s.RuntimeSettings
}

func (s *GetDIJobResponseBodyPagingInfoJobSettings) SetChannelSettings(v string) *GetDIJobResponseBodyPagingInfoJobSettings {
	s.ChannelSettings = &v
	return s
}

func (s *GetDIJobResponseBodyPagingInfoJobSettings) SetColumnDataTypeSettings(v []*GetDIJobResponseBodyPagingInfoJobSettingsColumnDataTypeSettings) *GetDIJobResponseBodyPagingInfoJobSettings {
	s.ColumnDataTypeSettings = v
	return s
}

func (s *GetDIJobResponseBodyPagingInfoJobSettings) SetCycleScheduleSettings(v *GetDIJobResponseBodyPagingInfoJobSettingsCycleScheduleSettings) *GetDIJobResponseBodyPagingInfoJobSettings {
	s.CycleScheduleSettings = v
	return s
}

func (s *GetDIJobResponseBodyPagingInfoJobSettings) SetDdlHandlingSettings(v []*GetDIJobResponseBodyPagingInfoJobSettingsDdlHandlingSettings) *GetDIJobResponseBodyPagingInfoJobSettings {
	s.DdlHandlingSettings = v
	return s
}

func (s *GetDIJobResponseBodyPagingInfoJobSettings) SetRuntimeSettings(v []*GetDIJobResponseBodyPagingInfoJobSettingsRuntimeSettings) *GetDIJobResponseBodyPagingInfoJobSettings {
	s.RuntimeSettings = v
	return s
}

func (s *GetDIJobResponseBodyPagingInfoJobSettings) Validate() error {
	return dara.Validate(s)
}

type GetDIJobResponseBodyPagingInfoJobSettingsColumnDataTypeSettings struct {
	// The data type of the destination field. Valid values: bigint, boolean, string, text, datetime, timestamp, decimal, and binary. Different types of data sources support different data types.
	//
	// example:
	//
	// text
	DestinationDataType *string `json:"DestinationDataType,omitempty" xml:"DestinationDataType,omitempty"`
	// The data type of the source field. Valid values: bigint, boolean, string, text, datetime, timestamp, decimal, and binary. Different types of data sources support different data types.
	//
	// example:
	//
	// bigint
	SourceDataType *string `json:"SourceDataType,omitempty" xml:"SourceDataType,omitempty"`
}

func (s GetDIJobResponseBodyPagingInfoJobSettingsColumnDataTypeSettings) String() string {
	return dara.Prettify(s)
}

func (s GetDIJobResponseBodyPagingInfoJobSettingsColumnDataTypeSettings) GoString() string {
	return s.String()
}

func (s *GetDIJobResponseBodyPagingInfoJobSettingsColumnDataTypeSettings) GetDestinationDataType() *string {
	return s.DestinationDataType
}

func (s *GetDIJobResponseBodyPagingInfoJobSettingsColumnDataTypeSettings) GetSourceDataType() *string {
	return s.SourceDataType
}

func (s *GetDIJobResponseBodyPagingInfoJobSettingsColumnDataTypeSettings) SetDestinationDataType(v string) *GetDIJobResponseBodyPagingInfoJobSettingsColumnDataTypeSettings {
	s.DestinationDataType = &v
	return s
}

func (s *GetDIJobResponseBodyPagingInfoJobSettingsColumnDataTypeSettings) SetSourceDataType(v string) *GetDIJobResponseBodyPagingInfoJobSettingsColumnDataTypeSettings {
	s.SourceDataType = &v
	return s
}

func (s *GetDIJobResponseBodyPagingInfoJobSettingsColumnDataTypeSettings) Validate() error {
	return dara.Validate(s)
}

type GetDIJobResponseBodyPagingInfoJobSettingsCycleScheduleSettings struct {
	// The synchronization type that requires periodic scheduling. Valid values:
	//
	// 	- Full: full synchronization
	//
	// 	- OfflineIncremental: batch incremental synchronization
	//
	// example:
	//
	// Full
	CycleMigrationType *string `json:"CycleMigrationType,omitempty" xml:"CycleMigrationType,omitempty"`
	// The scheduling parameters.
	//
	// example:
	//
	// bizdate=$bizdate
	ScheduleParameters *string `json:"ScheduleParameters,omitempty" xml:"ScheduleParameters,omitempty"`
}

func (s GetDIJobResponseBodyPagingInfoJobSettingsCycleScheduleSettings) String() string {
	return dara.Prettify(s)
}

func (s GetDIJobResponseBodyPagingInfoJobSettingsCycleScheduleSettings) GoString() string {
	return s.String()
}

func (s *GetDIJobResponseBodyPagingInfoJobSettingsCycleScheduleSettings) GetCycleMigrationType() *string {
	return s.CycleMigrationType
}

func (s *GetDIJobResponseBodyPagingInfoJobSettingsCycleScheduleSettings) GetScheduleParameters() *string {
	return s.ScheduleParameters
}

func (s *GetDIJobResponseBodyPagingInfoJobSettingsCycleScheduleSettings) SetCycleMigrationType(v string) *GetDIJobResponseBodyPagingInfoJobSettingsCycleScheduleSettings {
	s.CycleMigrationType = &v
	return s
}

func (s *GetDIJobResponseBodyPagingInfoJobSettingsCycleScheduleSettings) SetScheduleParameters(v string) *GetDIJobResponseBodyPagingInfoJobSettingsCycleScheduleSettings {
	s.ScheduleParameters = &v
	return s
}

func (s *GetDIJobResponseBodyPagingInfoJobSettingsCycleScheduleSettings) Validate() error {
	return dara.Validate(s)
}

type GetDIJobResponseBodyPagingInfoJobSettingsDdlHandlingSettings struct {
	// The processing policy for a specific type of DDL message. Valid values:
	//
	// 	- Ignore: ignores a DDL message.
	//
	// 	- Critical: reports an error for a DDL message.
	//
	// 	- Normal: normally processes a DDL message.
	//
	// example:
	//
	// Ignore
	Action *string `json:"Action,omitempty" xml:"Action,omitempty"`
	// The DDL operation type. Valid values:
	//
	// 	- RenameColumn
	//
	// 	- ModifyColumn
	//
	// 	- CreateTable
	//
	// 	- TruncateTable
	//
	// 	- DropTable
	//
	// example:
	//
	// CreateTable
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetDIJobResponseBodyPagingInfoJobSettingsDdlHandlingSettings) String() string {
	return dara.Prettify(s)
}

func (s GetDIJobResponseBodyPagingInfoJobSettingsDdlHandlingSettings) GoString() string {
	return s.String()
}

func (s *GetDIJobResponseBodyPagingInfoJobSettingsDdlHandlingSettings) GetAction() *string {
	return s.Action
}

func (s *GetDIJobResponseBodyPagingInfoJobSettingsDdlHandlingSettings) GetType() *string {
	return s.Type
}

func (s *GetDIJobResponseBodyPagingInfoJobSettingsDdlHandlingSettings) SetAction(v string) *GetDIJobResponseBodyPagingInfoJobSettingsDdlHandlingSettings {
	s.Action = &v
	return s
}

func (s *GetDIJobResponseBodyPagingInfoJobSettingsDdlHandlingSettings) SetType(v string) *GetDIJobResponseBodyPagingInfoJobSettingsDdlHandlingSettings {
	s.Type = &v
	return s
}

func (s *GetDIJobResponseBodyPagingInfoJobSettingsDdlHandlingSettings) Validate() error {
	return dara.Validate(s)
}

type GetDIJobResponseBodyPagingInfoJobSettingsRuntimeSettings struct {
	// The name of the configuration item. Valid values:
	//
	// 	- src.offline.datasource.max.connection: indicates the maximum number of connections that are allowed for reading data from the source of a batch synchronization task.
	//
	// 	- dst.offline.truncate: indicates whether to clear the destination table before data writing.
	//
	// 	- runtime.offline.speed.limit.enable: indicates whether throttling is enabled for a batch synchronization task.
	//
	// 	- runtime.offline.concurrent: indicates the maximum number of parallel threads that are allowed for a batch synchronization task.
	//
	// 	- runtime.enable.auto.create.schema: indicates whether schemas are automatically created in the destination of a synchronization task.
	//
	// 	- runtime.realtime.concurrent: indicates the maximum number of parallel threads that are allowed for a real-time synchronization task.
	//
	// 	- runtime.realtime.failover.minute.dataxcdc: indicates the maximum waiting duration before a synchronization task retries the next restart if the previous restart fails after failover occurs. Unit: minutes.
	//
	// 	- runtime.realtime.failover.times.dataxcdc: indicates the maximum number of failures that are allowed for restarting a synchronization task after failovers occur.
	//
	// example:
	//
	// runtime.offline.concurrent
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The value of the configuration item.
	//
	// example:
	//
	// 1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetDIJobResponseBodyPagingInfoJobSettingsRuntimeSettings) String() string {
	return dara.Prettify(s)
}

func (s GetDIJobResponseBodyPagingInfoJobSettingsRuntimeSettings) GoString() string {
	return s.String()
}

func (s *GetDIJobResponseBodyPagingInfoJobSettingsRuntimeSettings) GetName() *string {
	return s.Name
}

func (s *GetDIJobResponseBodyPagingInfoJobSettingsRuntimeSettings) GetValue() *string {
	return s.Value
}

func (s *GetDIJobResponseBodyPagingInfoJobSettingsRuntimeSettings) SetName(v string) *GetDIJobResponseBodyPagingInfoJobSettingsRuntimeSettings {
	s.Name = &v
	return s
}

func (s *GetDIJobResponseBodyPagingInfoJobSettingsRuntimeSettings) SetValue(v string) *GetDIJobResponseBodyPagingInfoJobSettingsRuntimeSettings {
	s.Value = &v
	return s
}

func (s *GetDIJobResponseBodyPagingInfoJobSettingsRuntimeSettings) Validate() error {
	return dara.Validate(s)
}

type GetDIJobResponseBodyPagingInfoResourceSettings struct {
	// The resource used for batch synchronization.
	OfflineResourceSettings *GetDIJobResponseBodyPagingInfoResourceSettingsOfflineResourceSettings `json:"OfflineResourceSettings,omitempty" xml:"OfflineResourceSettings,omitempty" type:"Struct"`
	// The resource used for real-time synchronization.
	RealtimeResourceSettings *GetDIJobResponseBodyPagingInfoResourceSettingsRealtimeResourceSettings `json:"RealtimeResourceSettings,omitempty" xml:"RealtimeResourceSettings,omitempty" type:"Struct"`
	// The resource used for scheduling.
	ScheduleResourceSettings *GetDIJobResponseBodyPagingInfoResourceSettingsScheduleResourceSettings `json:"ScheduleResourceSettings,omitempty" xml:"ScheduleResourceSettings,omitempty" type:"Struct"`
}

func (s GetDIJobResponseBodyPagingInfoResourceSettings) String() string {
	return dara.Prettify(s)
}

func (s GetDIJobResponseBodyPagingInfoResourceSettings) GoString() string {
	return s.String()
}

func (s *GetDIJobResponseBodyPagingInfoResourceSettings) GetOfflineResourceSettings() *GetDIJobResponseBodyPagingInfoResourceSettingsOfflineResourceSettings {
	return s.OfflineResourceSettings
}

func (s *GetDIJobResponseBodyPagingInfoResourceSettings) GetRealtimeResourceSettings() *GetDIJobResponseBodyPagingInfoResourceSettingsRealtimeResourceSettings {
	return s.RealtimeResourceSettings
}

func (s *GetDIJobResponseBodyPagingInfoResourceSettings) GetScheduleResourceSettings() *GetDIJobResponseBodyPagingInfoResourceSettingsScheduleResourceSettings {
	return s.ScheduleResourceSettings
}

func (s *GetDIJobResponseBodyPagingInfoResourceSettings) SetOfflineResourceSettings(v *GetDIJobResponseBodyPagingInfoResourceSettingsOfflineResourceSettings) *GetDIJobResponseBodyPagingInfoResourceSettings {
	s.OfflineResourceSettings = v
	return s
}

func (s *GetDIJobResponseBodyPagingInfoResourceSettings) SetRealtimeResourceSettings(v *GetDIJobResponseBodyPagingInfoResourceSettingsRealtimeResourceSettings) *GetDIJobResponseBodyPagingInfoResourceSettings {
	s.RealtimeResourceSettings = v
	return s
}

func (s *GetDIJobResponseBodyPagingInfoResourceSettings) SetScheduleResourceSettings(v *GetDIJobResponseBodyPagingInfoResourceSettingsScheduleResourceSettings) *GetDIJobResponseBodyPagingInfoResourceSettings {
	s.ScheduleResourceSettings = v
	return s
}

func (s *GetDIJobResponseBodyPagingInfoResourceSettings) Validate() error {
	return dara.Validate(s)
}

type GetDIJobResponseBodyPagingInfoResourceSettingsOfflineResourceSettings struct {
	// The number of compute units (CUs) in the resource group for scheduling that are used for batch synchronization.
	//
	// example:
	//
	// 2.0
	RequestedCu *float64 `json:"RequestedCu,omitempty" xml:"RequestedCu,omitempty"`
	// The identifier of the resource group for Data Integration used for batch synchronization.
	//
	// example:
	//
	// S_res_group_7708_1667792816832
	ResourceGroupIdentifier *string `json:"ResourceGroupIdentifier,omitempty" xml:"ResourceGroupIdentifier,omitempty"`
}

func (s GetDIJobResponseBodyPagingInfoResourceSettingsOfflineResourceSettings) String() string {
	return dara.Prettify(s)
}

func (s GetDIJobResponseBodyPagingInfoResourceSettingsOfflineResourceSettings) GoString() string {
	return s.String()
}

func (s *GetDIJobResponseBodyPagingInfoResourceSettingsOfflineResourceSettings) GetRequestedCu() *float64 {
	return s.RequestedCu
}

func (s *GetDIJobResponseBodyPagingInfoResourceSettingsOfflineResourceSettings) GetResourceGroupIdentifier() *string {
	return s.ResourceGroupIdentifier
}

func (s *GetDIJobResponseBodyPagingInfoResourceSettingsOfflineResourceSettings) SetRequestedCu(v float64) *GetDIJobResponseBodyPagingInfoResourceSettingsOfflineResourceSettings {
	s.RequestedCu = &v
	return s
}

func (s *GetDIJobResponseBodyPagingInfoResourceSettingsOfflineResourceSettings) SetResourceGroupIdentifier(v string) *GetDIJobResponseBodyPagingInfoResourceSettingsOfflineResourceSettings {
	s.ResourceGroupIdentifier = &v
	return s
}

func (s *GetDIJobResponseBodyPagingInfoResourceSettingsOfflineResourceSettings) Validate() error {
	return dara.Validate(s)
}

type GetDIJobResponseBodyPagingInfoResourceSettingsRealtimeResourceSettings struct {
	// The number of CUs in the resource group for Data Integration that are used for real-time synchronization.
	//
	// example:
	//
	// 2.0
	RequestedCu *float64 `json:"RequestedCu,omitempty" xml:"RequestedCu,omitempty"`
	// The identifier of the resource group for Data Integration used for real-time synchronization.
	//
	// example:
	//
	// S_res_group_235454102432001_1579085295030
	ResourceGroupIdentifier *string `json:"ResourceGroupIdentifier,omitempty" xml:"ResourceGroupIdentifier,omitempty"`
}

func (s GetDIJobResponseBodyPagingInfoResourceSettingsRealtimeResourceSettings) String() string {
	return dara.Prettify(s)
}

func (s GetDIJobResponseBodyPagingInfoResourceSettingsRealtimeResourceSettings) GoString() string {
	return s.String()
}

func (s *GetDIJobResponseBodyPagingInfoResourceSettingsRealtimeResourceSettings) GetRequestedCu() *float64 {
	return s.RequestedCu
}

func (s *GetDIJobResponseBodyPagingInfoResourceSettingsRealtimeResourceSettings) GetResourceGroupIdentifier() *string {
	return s.ResourceGroupIdentifier
}

func (s *GetDIJobResponseBodyPagingInfoResourceSettingsRealtimeResourceSettings) SetRequestedCu(v float64) *GetDIJobResponseBodyPagingInfoResourceSettingsRealtimeResourceSettings {
	s.RequestedCu = &v
	return s
}

func (s *GetDIJobResponseBodyPagingInfoResourceSettingsRealtimeResourceSettings) SetResourceGroupIdentifier(v string) *GetDIJobResponseBodyPagingInfoResourceSettingsRealtimeResourceSettings {
	s.ResourceGroupIdentifier = &v
	return s
}

func (s *GetDIJobResponseBodyPagingInfoResourceSettingsRealtimeResourceSettings) Validate() error {
	return dara.Validate(s)
}

type GetDIJobResponseBodyPagingInfoResourceSettingsScheduleResourceSettings struct {
	// The number of CUs in the resource group for Data Integration that are used for scheduling.
	//
	// example:
	//
	// 2.0
	RequestedCu *float64 `json:"RequestedCu,omitempty" xml:"RequestedCu,omitempty"`
	// The identifier of the resource group for scheduling used by the synchronization task.
	//
	// example:
	//
	// S_res_group_235454102432001_1718359176885
	ResourceGroupIdentifier *string `json:"ResourceGroupIdentifier,omitempty" xml:"ResourceGroupIdentifier,omitempty"`
}

func (s GetDIJobResponseBodyPagingInfoResourceSettingsScheduleResourceSettings) String() string {
	return dara.Prettify(s)
}

func (s GetDIJobResponseBodyPagingInfoResourceSettingsScheduleResourceSettings) GoString() string {
	return s.String()
}

func (s *GetDIJobResponseBodyPagingInfoResourceSettingsScheduleResourceSettings) GetRequestedCu() *float64 {
	return s.RequestedCu
}

func (s *GetDIJobResponseBodyPagingInfoResourceSettingsScheduleResourceSettings) GetResourceGroupIdentifier() *string {
	return s.ResourceGroupIdentifier
}

func (s *GetDIJobResponseBodyPagingInfoResourceSettingsScheduleResourceSettings) SetRequestedCu(v float64) *GetDIJobResponseBodyPagingInfoResourceSettingsScheduleResourceSettings {
	s.RequestedCu = &v
	return s
}

func (s *GetDIJobResponseBodyPagingInfoResourceSettingsScheduleResourceSettings) SetResourceGroupIdentifier(v string) *GetDIJobResponseBodyPagingInfoResourceSettingsScheduleResourceSettings {
	s.ResourceGroupIdentifier = &v
	return s
}

func (s *GetDIJobResponseBodyPagingInfoResourceSettingsScheduleResourceSettings) Validate() error {
	return dara.Validate(s)
}

type GetDIJobResponseBodyPagingInfoSourceDataSourceSettings struct {
	// The name of the data source.
	//
	// example:
	//
	// dw_mysql
	DataSourceName *string `json:"DataSourceName,omitempty" xml:"DataSourceName,omitempty"`
	// The properties of the data source.
	DataSourceProperties *GetDIJobResponseBodyPagingInfoSourceDataSourceSettingsDataSourceProperties `json:"DataSourceProperties,omitempty" xml:"DataSourceProperties,omitempty" type:"Struct"`
}

func (s GetDIJobResponseBodyPagingInfoSourceDataSourceSettings) String() string {
	return dara.Prettify(s)
}

func (s GetDIJobResponseBodyPagingInfoSourceDataSourceSettings) GoString() string {
	return s.String()
}

func (s *GetDIJobResponseBodyPagingInfoSourceDataSourceSettings) GetDataSourceName() *string {
	return s.DataSourceName
}

func (s *GetDIJobResponseBodyPagingInfoSourceDataSourceSettings) GetDataSourceProperties() *GetDIJobResponseBodyPagingInfoSourceDataSourceSettingsDataSourceProperties {
	return s.DataSourceProperties
}

func (s *GetDIJobResponseBodyPagingInfoSourceDataSourceSettings) SetDataSourceName(v string) *GetDIJobResponseBodyPagingInfoSourceDataSourceSettings {
	s.DataSourceName = &v
	return s
}

func (s *GetDIJobResponseBodyPagingInfoSourceDataSourceSettings) SetDataSourceProperties(v *GetDIJobResponseBodyPagingInfoSourceDataSourceSettingsDataSourceProperties) *GetDIJobResponseBodyPagingInfoSourceDataSourceSettings {
	s.DataSourceProperties = v
	return s
}

func (s *GetDIJobResponseBodyPagingInfoSourceDataSourceSettings) Validate() error {
	return dara.Validate(s)
}

type GetDIJobResponseBodyPagingInfoSourceDataSourceSettingsDataSourceProperties struct {
	// The encoding format of the database.
	//
	// example:
	//
	// UTF-8
	Encoding *string `json:"Encoding,omitempty" xml:"Encoding,omitempty"`
	// The time zone.
	//
	// example:
	//
	// GMT+8
	Timezone *string `json:"Timezone,omitempty" xml:"Timezone,omitempty"`
}

func (s GetDIJobResponseBodyPagingInfoSourceDataSourceSettingsDataSourceProperties) String() string {
	return dara.Prettify(s)
}

func (s GetDIJobResponseBodyPagingInfoSourceDataSourceSettingsDataSourceProperties) GoString() string {
	return s.String()
}

func (s *GetDIJobResponseBodyPagingInfoSourceDataSourceSettingsDataSourceProperties) GetEncoding() *string {
	return s.Encoding
}

func (s *GetDIJobResponseBodyPagingInfoSourceDataSourceSettingsDataSourceProperties) GetTimezone() *string {
	return s.Timezone
}

func (s *GetDIJobResponseBodyPagingInfoSourceDataSourceSettingsDataSourceProperties) SetEncoding(v string) *GetDIJobResponseBodyPagingInfoSourceDataSourceSettingsDataSourceProperties {
	s.Encoding = &v
	return s
}

func (s *GetDIJobResponseBodyPagingInfoSourceDataSourceSettingsDataSourceProperties) SetTimezone(v string) *GetDIJobResponseBodyPagingInfoSourceDataSourceSettingsDataSourceProperties {
	s.Timezone = &v
	return s
}

func (s *GetDIJobResponseBodyPagingInfoSourceDataSourceSettingsDataSourceProperties) Validate() error {
	return dara.Validate(s)
}

type GetDIJobResponseBodyPagingInfoTableMappings struct {
	// The list of rules used to select synchronization objects in the source.
	SourceObjectSelectionRules []*GetDIJobResponseBodyPagingInfoTableMappingsSourceObjectSelectionRules `json:"SourceObjectSelectionRules,omitempty" xml:"SourceObjectSelectionRules,omitempty" type:"Repeated"`
	// The list of transformation rules that are applied to the synchronization objects selected from the source. Each entry in the list defines a transformation rule.
	TransformationRules []*GetDIJobResponseBodyPagingInfoTableMappingsTransformationRules `json:"TransformationRules,omitempty" xml:"TransformationRules,omitempty" type:"Repeated"`
}

func (s GetDIJobResponseBodyPagingInfoTableMappings) String() string {
	return dara.Prettify(s)
}

func (s GetDIJobResponseBodyPagingInfoTableMappings) GoString() string {
	return s.String()
}

func (s *GetDIJobResponseBodyPagingInfoTableMappings) GetSourceObjectSelectionRules() []*GetDIJobResponseBodyPagingInfoTableMappingsSourceObjectSelectionRules {
	return s.SourceObjectSelectionRules
}

func (s *GetDIJobResponseBodyPagingInfoTableMappings) GetTransformationRules() []*GetDIJobResponseBodyPagingInfoTableMappingsTransformationRules {
	return s.TransformationRules
}

func (s *GetDIJobResponseBodyPagingInfoTableMappings) SetSourceObjectSelectionRules(v []*GetDIJobResponseBodyPagingInfoTableMappingsSourceObjectSelectionRules) *GetDIJobResponseBodyPagingInfoTableMappings {
	s.SourceObjectSelectionRules = v
	return s
}

func (s *GetDIJobResponseBodyPagingInfoTableMappings) SetTransformationRules(v []*GetDIJobResponseBodyPagingInfoTableMappingsTransformationRules) *GetDIJobResponseBodyPagingInfoTableMappings {
	s.TransformationRules = v
	return s
}

func (s *GetDIJobResponseBodyPagingInfoTableMappings) Validate() error {
	return dara.Validate(s)
}

type GetDIJobResponseBodyPagingInfoTableMappingsSourceObjectSelectionRules struct {
	// The operation that is performed to select objects. Valid values: Include and Exclude.
	//
	// example:
	//
	// Include
	Action *string `json:"Action,omitempty" xml:"Action,omitempty"`
	// The expression.
	//
	// example:
	//
	// mysql_table_1
	Expression *string `json:"Expression,omitempty" xml:"Expression,omitempty"`
	// The expression type. Valid values: Exact and Regex.
	//
	// example:
	//
	// Exact
	ExpressionType *string `json:"ExpressionType,omitempty" xml:"ExpressionType,omitempty"`
	// The object type. Valid values:
	//
	// 	- Table
	//
	// 	- Schema
	//
	// 	- Database
	//
	// example:
	//
	// Table
	ObjectType *string `json:"ObjectType,omitempty" xml:"ObjectType,omitempty"`
}

func (s GetDIJobResponseBodyPagingInfoTableMappingsSourceObjectSelectionRules) String() string {
	return dara.Prettify(s)
}

func (s GetDIJobResponseBodyPagingInfoTableMappingsSourceObjectSelectionRules) GoString() string {
	return s.String()
}

func (s *GetDIJobResponseBodyPagingInfoTableMappingsSourceObjectSelectionRules) GetAction() *string {
	return s.Action
}

func (s *GetDIJobResponseBodyPagingInfoTableMappingsSourceObjectSelectionRules) GetExpression() *string {
	return s.Expression
}

func (s *GetDIJobResponseBodyPagingInfoTableMappingsSourceObjectSelectionRules) GetExpressionType() *string {
	return s.ExpressionType
}

func (s *GetDIJobResponseBodyPagingInfoTableMappingsSourceObjectSelectionRules) GetObjectType() *string {
	return s.ObjectType
}

func (s *GetDIJobResponseBodyPagingInfoTableMappingsSourceObjectSelectionRules) SetAction(v string) *GetDIJobResponseBodyPagingInfoTableMappingsSourceObjectSelectionRules {
	s.Action = &v
	return s
}

func (s *GetDIJobResponseBodyPagingInfoTableMappingsSourceObjectSelectionRules) SetExpression(v string) *GetDIJobResponseBodyPagingInfoTableMappingsSourceObjectSelectionRules {
	s.Expression = &v
	return s
}

func (s *GetDIJobResponseBodyPagingInfoTableMappingsSourceObjectSelectionRules) SetExpressionType(v string) *GetDIJobResponseBodyPagingInfoTableMappingsSourceObjectSelectionRules {
	s.ExpressionType = &v
	return s
}

func (s *GetDIJobResponseBodyPagingInfoTableMappingsSourceObjectSelectionRules) SetObjectType(v string) *GetDIJobResponseBodyPagingInfoTableMappingsSourceObjectSelectionRules {
	s.ObjectType = &v
	return s
}

func (s *GetDIJobResponseBodyPagingInfoTableMappingsSourceObjectSelectionRules) Validate() error {
	return dara.Validate(s)
}

type GetDIJobResponseBodyPagingInfoTableMappingsTransformationRules struct {
	// The action type. Valid values:
	//
	// 	- DefinePrimaryKey
	//
	// 	- Rename
	//
	// 	- AddColumn
	//
	// 	- HandleDml
	//
	// example:
	//
	// AddColumn
	RuleActionType *string `json:"RuleActionType,omitempty" xml:"RuleActionType,omitempty"`
	// The name of the rule. If the values of the RuleActionType parameter and the RuleTargetType parameter are the same for multiple transformation rules, you must make sure that the transformation rule names are unique.
	//
	// example:
	//
	// rename_rule_1
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The type of the object on which the action is performed. Valid values:
	//
	// 	- Table
	//
	// 	- Schema
	//
	// 	- Database
	//
	// example:
	//
	// Table
	RuleTargetType *string `json:"RuleTargetType,omitempty" xml:"RuleTargetType,omitempty"`
}

func (s GetDIJobResponseBodyPagingInfoTableMappingsTransformationRules) String() string {
	return dara.Prettify(s)
}

func (s GetDIJobResponseBodyPagingInfoTableMappingsTransformationRules) GoString() string {
	return s.String()
}

func (s *GetDIJobResponseBodyPagingInfoTableMappingsTransformationRules) GetRuleActionType() *string {
	return s.RuleActionType
}

func (s *GetDIJobResponseBodyPagingInfoTableMappingsTransformationRules) GetRuleName() *string {
	return s.RuleName
}

func (s *GetDIJobResponseBodyPagingInfoTableMappingsTransformationRules) GetRuleTargetType() *string {
	return s.RuleTargetType
}

func (s *GetDIJobResponseBodyPagingInfoTableMappingsTransformationRules) SetRuleActionType(v string) *GetDIJobResponseBodyPagingInfoTableMappingsTransformationRules {
	s.RuleActionType = &v
	return s
}

func (s *GetDIJobResponseBodyPagingInfoTableMappingsTransformationRules) SetRuleName(v string) *GetDIJobResponseBodyPagingInfoTableMappingsTransformationRules {
	s.RuleName = &v
	return s
}

func (s *GetDIJobResponseBodyPagingInfoTableMappingsTransformationRules) SetRuleTargetType(v string) *GetDIJobResponseBodyPagingInfoTableMappingsTransformationRules {
	s.RuleTargetType = &v
	return s
}

func (s *GetDIJobResponseBodyPagingInfoTableMappingsTransformationRules) Validate() error {
	return dara.Validate(s)
}

type GetDIJobResponseBodyPagingInfoTransformationRules struct {
	// The action type. Valid values:
	//
	// 	- DefinePrimaryKey
	//
	// 	- Rename
	//
	// 	- AddColumn
	//
	// 	- HandleDml
	//
	// 	- DefineIncrementalCondition
	//
	// 	- DefineCycleScheduleSettings
	//
	// 	- DefinePartitionKey
	//
	// example:
	//
	// Rename
	RuleActionType *string `json:"RuleActionType,omitempty" xml:"RuleActionType,omitempty"`
	// The expression of the rule. The expression is a JSON string.
	//
	// 1.  Example of a renaming rule
	//
	// 	- Example: {"expression":"${srcDatasourceName}_${srcDatabaseName}_0922" }
	//
	// 	- expression: the expression of the renaming rule. You can use the following variables in an expression: ${srcDatasourceName}, ${srcDatabaseName}, and ${srcTableName}. ${srcDatasourceName} indicates the name of the source. ${srcDatabaseName} indicates the name of a source database. ${srcTableName} indicates the name of a source table.
	//
	// 2.  Example of a column addition rule
	//
	// 	- Example: {"columns":[{"columnName":"my_add_column","columnValueType":"Constant","columnValue":"123"}]}
	//
	// 	- If no rule of this type is configured, no fields are added to the destination and no values are assigned by default.
	//
	// 	- columnName: the name of the field that is added.
	//
	// 	- columnValueType: the value type of the field. Valid values: Constant and Variable.
	//
	// 	- columnValue: the value of the field. If the columnValueType parameter is set to Constant, the value of the columnValue parameter is a constant of the STRING data type. If the columnValueType parameter is set to Variable, the value of the columnValue parameter is a built-in variable. The following built-in variables are supported: EXECUTE_TIME (LONG data type), DB_NAME_SRC (STRING data type), DATASOURCE_NAME_SRC (STRING data type), TABLE_NAME_SRC (STRING data type), DB_NAME_DEST (STRING data type), DATASOURCE_NAME_DEST (STRING data type), TABLE_NAME_DEST (STRING data type), and DB_NAME_SRC_TRANSED (STRING data type). EXECUTE_TIME indicates the execution time. DB_NAME_SRC indicates the name of a source database. DATASOURCE_NAME_SRC indicates the name of the source. TABLE_NAME_SRC indicates the name of a source table. DB_NAME_DEST indicates the name of a destination database. DATASOURCE_NAME_DEST indicates the name of the destination. TABLE_NAME_DEST indicates the name of a destination table. DB_NAME_SRC_TRANSED indicates the database name obtained after a transformation.
	//
	// 3.  Example of a rule used to specify primary key fields for a destination table
	//
	// 	- Example: {"columns":["ukcolumn1","ukcolumn2"]}
	//
	// 	- If no rule of this type is configured, the primary key fields in the mapped source table are used for the destination table by default.
	//
	// 	- If the destination table is an existing table, Data Integration does not modify the schema of the destination table. If the specified primary key fields do not exist in the destination table, an error is reported when the synchronization task starts to run.
	//
	// 	- If the destination table is automatically created by the system, Data Integration automatically creates the schema of the destination table. The schema contains the primary key fields that you specify. If the specified primary key fields do not exist in the destination table, an error is reported when the synchronization task starts to run.
	//
	// 4.  Example of a rule used to process DML messages
	//
	// 	- Example: {"dmlPolicies":[{"dmlType":"Delete","dmlAction":"Filter","filterCondition":"id > 1"}]}
	//
	// 	- If no rule of this type is configured, the default processing policy for messages generated for insert, update, and delete operations is Normal.
	//
	// 	- dmlType: the DML operation. Valid values: Insert, Update, and Delete.
	//
	// 	- dmlAction: the processing policy for DML messages. Valid values: Normal, Ignore, Filter, and LogicalDelete. Filter indicates conditional processing. The value Filter is returned for the dmlAction parameter only when the value of the dmlType parameter is Update or Delete.
	//
	// 	- filterCondition: the condition used to filter DML messages. This parameter is returned only when the value of the dmlAction parameter is Filter.
	//
	// 5.  Example of a rule used to perform incremental synchronization
	//
	// 	- Example: {"where":"id > 0"}
	//
	// 	- The rule used to perform incremental synchronization is returned.
	//
	// 6.  Example of a rule used to configure scheduling parameters for an auto triggered task
	//
	// 	- Example: {"cronExpress":" \\	- \\	- \\	- \\	- \\	- \\*", "cycleType":"1"}
	//
	// 	- The rule used to configure scheduling parameters for an auto triggered task is returned.
	//
	// 7.  Example of a rule used to specify a partition key
	//
	// 	- Example: {"columns":["id"]}
	//
	// 	- The rule used to specify a partition key is returned.
	//
	// example:
	//
	// {"expression":"${srcDatasoureName}_${srcDatabaseName}"}
	RuleExpression *string `json:"RuleExpression,omitempty" xml:"RuleExpression,omitempty"`
	// The name of the rule. If the values of the RuleActionType parameter and the RuleTargetType parameter are the same for multiple transformation rules, you must make sure that the transformation rule names are unique.
	//
	// example:
	//
	// rename_rule_1
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The type of the object on which the action is performed. Valid values:
	//
	// 	- Table
	//
	// 	- Schema
	//
	// 	- Database
	//
	// example:
	//
	// Table
	RuleTargetType *string `json:"RuleTargetType,omitempty" xml:"RuleTargetType,omitempty"`
}

func (s GetDIJobResponseBodyPagingInfoTransformationRules) String() string {
	return dara.Prettify(s)
}

func (s GetDIJobResponseBodyPagingInfoTransformationRules) GoString() string {
	return s.String()
}

func (s *GetDIJobResponseBodyPagingInfoTransformationRules) GetRuleActionType() *string {
	return s.RuleActionType
}

func (s *GetDIJobResponseBodyPagingInfoTransformationRules) GetRuleExpression() *string {
	return s.RuleExpression
}

func (s *GetDIJobResponseBodyPagingInfoTransformationRules) GetRuleName() *string {
	return s.RuleName
}

func (s *GetDIJobResponseBodyPagingInfoTransformationRules) GetRuleTargetType() *string {
	return s.RuleTargetType
}

func (s *GetDIJobResponseBodyPagingInfoTransformationRules) SetRuleActionType(v string) *GetDIJobResponseBodyPagingInfoTransformationRules {
	s.RuleActionType = &v
	return s
}

func (s *GetDIJobResponseBodyPagingInfoTransformationRules) SetRuleExpression(v string) *GetDIJobResponseBodyPagingInfoTransformationRules {
	s.RuleExpression = &v
	return s
}

func (s *GetDIJobResponseBodyPagingInfoTransformationRules) SetRuleName(v string) *GetDIJobResponseBodyPagingInfoTransformationRules {
	s.RuleName = &v
	return s
}

func (s *GetDIJobResponseBodyPagingInfoTransformationRules) SetRuleTargetType(v string) *GetDIJobResponseBodyPagingInfoTransformationRules {
	s.RuleTargetType = &v
	return s
}

func (s *GetDIJobResponseBodyPagingInfoTransformationRules) Validate() error {
	return dara.Validate(s)
}
