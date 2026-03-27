// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDIJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateDIJobRequest
	GetDescription() *string
	SetDestinationDataSourceSettings(v []*CreateDIJobRequestDestinationDataSourceSettings) *CreateDIJobRequest
	GetDestinationDataSourceSettings() []*CreateDIJobRequestDestinationDataSourceSettings
	SetDestinationDataSourceType(v string) *CreateDIJobRequest
	GetDestinationDataSourceType() *string
	SetFileSpec(v string) *CreateDIJobRequest
	GetFileSpec() *string
	SetJobName(v string) *CreateDIJobRequest
	GetJobName() *string
	SetJobSettings(v *CreateDIJobRequestJobSettings) *CreateDIJobRequest
	GetJobSettings() *CreateDIJobRequestJobSettings
	SetJobType(v string) *CreateDIJobRequest
	GetJobType() *string
	SetMigrationType(v string) *CreateDIJobRequest
	GetMigrationType() *string
	SetName(v string) *CreateDIJobRequest
	GetName() *string
	SetOwner(v string) *CreateDIJobRequest
	GetOwner() *string
	SetProjectId(v int64) *CreateDIJobRequest
	GetProjectId() *int64
	SetResourceSettings(v *CreateDIJobRequestResourceSettings) *CreateDIJobRequest
	GetResourceSettings() *CreateDIJobRequestResourceSettings
	SetSourceDataSourceSettings(v []*CreateDIJobRequestSourceDataSourceSettings) *CreateDIJobRequest
	GetSourceDataSourceSettings() []*CreateDIJobRequestSourceDataSourceSettings
	SetSourceDataSourceType(v string) *CreateDIJobRequest
	GetSourceDataSourceType() *string
	SetTableMappings(v []*CreateDIJobRequestTableMappings) *CreateDIJobRequest
	GetTableMappings() []*CreateDIJobRequestTableMappings
	SetTransformationRules(v []*CreateDIJobRequestTransformationRules) *CreateDIJobRequest
	GetTransformationRules() []*CreateDIJobRequestTransformationRules
}

type CreateDIJobRequest struct {
	// The task description.
	//
	// example:
	//
	// The description of the synchronization task.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The list of destination data source settings.
	DestinationDataSourceSettings []*CreateDIJobRequestDestinationDataSourceSettings `json:"DestinationDataSourceSettings,omitempty" xml:"DestinationDataSourceSettings,omitempty" type:"Repeated"`
	// The destination type. Valid values: Hologres, OSS-HDFS, OSS, MaxCompute, LogHub, StarRocks, DataHub, AnalyticDB for MySQL, Kafka, and Hive.
	//
	// example:
	//
	// Hologres
	DestinationDataSourceType *string `json:"DestinationDataSourceType,omitempty" xml:"DestinationDataSourceType,omitempty"`
	FileSpec                  *string `json:"FileSpec,omitempty" xml:"FileSpec,omitempty"`
	// Deprecated
	//
	// This parameter is deprecated and is replaced by the Name parameter.
	//
	// example:
	//
	// mysql_to_holo_sync_8772
	JobName *string `json:"JobName,omitempty" xml:"JobName,omitempty"`
	// The task-level settings, including DDL handling policies, column data type mapping between source and destination, and runtime parameters.
	JobSettings *CreateDIJobRequestJobSettings `json:"JobSettings,omitempty" xml:"JobSettings,omitempty" type:"Struct"`
	// The type of the synchronization task. Valid values:
	//
	// 	- DatabaseRealtimeMigration: A real-time synchronization task used to synchronize only full data, only incremental data, or full and incremental data in multiple tables of multiple databases in the source.
	//
	// 	- DatabaseOfflineMigration: A batch synchronization task used to synchronize only full data, only incremental data, or full and incremental data in multiple tables of multiple databases in the source.
	//
	// 	- SingleTableRealtimeMigration: A real-time synchronization task used to synchronize data only in a single table in the source.
	//
	// example:
	//
	// DatabaseRealtimeMigration
	JobType *string `json:"JobType,omitempty" xml:"JobType,omitempty"`
	// The synchronization type. Valid values:
	//
	// 	- FullAndRealtimeIncremental
	//
	// 	- RealtimeIncremental
	//
	// 	- Full
	//
	// 	- OfflineIncremental
	//
	// 	- FullAndOfflineIncremental
	//
	// example:
	//
	// FullAndRealtimeIncremental
	MigrationType *string `json:"MigrationType,omitempty" xml:"MigrationType,omitempty"`
	// The name of the synchronization task.
	//
	// example:
	//
	// mysql_to_holo_sync_8772
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The task owner.
	//
	// example:
	//
	// 3726346
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The DataWorks workspace ID. You can log on to the [DataWorks console](https://workbench.data.aliyun.com/console) and go to the Workspace page to obtain the ID.
	//
	// You must configure this parameter to specify the DataWorks workspace to which the API operation is applied.
	//
	// example:
	//
	// 10000
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The resource settings.
	ResourceSettings *CreateDIJobRequestResourceSettings `json:"ResourceSettings,omitempty" xml:"ResourceSettings,omitempty" type:"Struct"`
	// The list of source data source settings.
	SourceDataSourceSettings []*CreateDIJobRequestSourceDataSourceSettings `json:"SourceDataSourceSettings,omitempty" xml:"SourceDataSourceSettings,omitempty" type:"Repeated"`
	// The source type. Valid values: PolarDB, MySQL, Kafka, LogHub, Hologres, Oracle, OceanBase, MongoDB, Redshift, Hive, SQL Server, Doris, and ClickHouse.
	//
	// example:
	//
	// MySQL
	SourceDataSourceType *string `json:"SourceDataSourceType,omitempty" xml:"SourceDataSourceType,omitempty"`
	// The list of synchronization object transformation mappings. Each element describes a set of source object selection rules and the transformation rules applied to those objects.
	//
	// >  [ { "SourceObjectSelectionRules":[ { "ObjectType":"Database", "Action":"Include", "ExpressionType":"Exact", "Expression":"biz_db" }, { "ObjectType":"Schema", "Action":"Include", "ExpressionType":"Exact", "Expression":"s1" }, { "ObjectType":"Table", "Action":"Include", "ExpressionType":"Exact", "Expression":"table1" } ], "TransformationRuleNames":[ { "RuleName":"my_database_rename_rule", "RuleActionType":"Rename", "RuleTargetType":"Schema" } ] } ]
	TableMappings []*CreateDIJobRequestTableMappings `json:"TableMappings,omitempty" xml:"TableMappings,omitempty" type:"Repeated"`
	// The list of synchronization object transformation rule definitions.
	//
	// >  [ { "RuleName":"my_database_rename_rule", "RuleActionType":"Rename", "RuleTargetType":"Schema", "RuleExpression":"{"expression":"${srcDatasoureName}_${srcDatabaseName}"}" } ]
	TransformationRules []*CreateDIJobRequestTransformationRules `json:"TransformationRules,omitempty" xml:"TransformationRules,omitempty" type:"Repeated"`
}

func (s CreateDIJobRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDIJobRequest) GoString() string {
	return s.String()
}

func (s *CreateDIJobRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateDIJobRequest) GetDestinationDataSourceSettings() []*CreateDIJobRequestDestinationDataSourceSettings {
	return s.DestinationDataSourceSettings
}

func (s *CreateDIJobRequest) GetDestinationDataSourceType() *string {
	return s.DestinationDataSourceType
}

func (s *CreateDIJobRequest) GetFileSpec() *string {
	return s.FileSpec
}

func (s *CreateDIJobRequest) GetJobName() *string {
	return s.JobName
}

func (s *CreateDIJobRequest) GetJobSettings() *CreateDIJobRequestJobSettings {
	return s.JobSettings
}

func (s *CreateDIJobRequest) GetJobType() *string {
	return s.JobType
}

func (s *CreateDIJobRequest) GetMigrationType() *string {
	return s.MigrationType
}

func (s *CreateDIJobRequest) GetName() *string {
	return s.Name
}

func (s *CreateDIJobRequest) GetOwner() *string {
	return s.Owner
}

func (s *CreateDIJobRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *CreateDIJobRequest) GetResourceSettings() *CreateDIJobRequestResourceSettings {
	return s.ResourceSettings
}

func (s *CreateDIJobRequest) GetSourceDataSourceSettings() []*CreateDIJobRequestSourceDataSourceSettings {
	return s.SourceDataSourceSettings
}

func (s *CreateDIJobRequest) GetSourceDataSourceType() *string {
	return s.SourceDataSourceType
}

func (s *CreateDIJobRequest) GetTableMappings() []*CreateDIJobRequestTableMappings {
	return s.TableMappings
}

func (s *CreateDIJobRequest) GetTransformationRules() []*CreateDIJobRequestTransformationRules {
	return s.TransformationRules
}

func (s *CreateDIJobRequest) SetDescription(v string) *CreateDIJobRequest {
	s.Description = &v
	return s
}

func (s *CreateDIJobRequest) SetDestinationDataSourceSettings(v []*CreateDIJobRequestDestinationDataSourceSettings) *CreateDIJobRequest {
	s.DestinationDataSourceSettings = v
	return s
}

func (s *CreateDIJobRequest) SetDestinationDataSourceType(v string) *CreateDIJobRequest {
	s.DestinationDataSourceType = &v
	return s
}

func (s *CreateDIJobRequest) SetFileSpec(v string) *CreateDIJobRequest {
	s.FileSpec = &v
	return s
}

func (s *CreateDIJobRequest) SetJobName(v string) *CreateDIJobRequest {
	s.JobName = &v
	return s
}

func (s *CreateDIJobRequest) SetJobSettings(v *CreateDIJobRequestJobSettings) *CreateDIJobRequest {
	s.JobSettings = v
	return s
}

func (s *CreateDIJobRequest) SetJobType(v string) *CreateDIJobRequest {
	s.JobType = &v
	return s
}

func (s *CreateDIJobRequest) SetMigrationType(v string) *CreateDIJobRequest {
	s.MigrationType = &v
	return s
}

func (s *CreateDIJobRequest) SetName(v string) *CreateDIJobRequest {
	s.Name = &v
	return s
}

func (s *CreateDIJobRequest) SetOwner(v string) *CreateDIJobRequest {
	s.Owner = &v
	return s
}

func (s *CreateDIJobRequest) SetProjectId(v int64) *CreateDIJobRequest {
	s.ProjectId = &v
	return s
}

func (s *CreateDIJobRequest) SetResourceSettings(v *CreateDIJobRequestResourceSettings) *CreateDIJobRequest {
	s.ResourceSettings = v
	return s
}

func (s *CreateDIJobRequest) SetSourceDataSourceSettings(v []*CreateDIJobRequestSourceDataSourceSettings) *CreateDIJobRequest {
	s.SourceDataSourceSettings = v
	return s
}

func (s *CreateDIJobRequest) SetSourceDataSourceType(v string) *CreateDIJobRequest {
	s.SourceDataSourceType = &v
	return s
}

func (s *CreateDIJobRequest) SetTableMappings(v []*CreateDIJobRequestTableMappings) *CreateDIJobRequest {
	s.TableMappings = v
	return s
}

func (s *CreateDIJobRequest) SetTransformationRules(v []*CreateDIJobRequestTransformationRules) *CreateDIJobRequest {
	s.TransformationRules = v
	return s
}

func (s *CreateDIJobRequest) Validate() error {
	if s.DestinationDataSourceSettings != nil {
		for _, item := range s.DestinationDataSourceSettings {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.JobSettings != nil {
		if err := s.JobSettings.Validate(); err != nil {
			return err
		}
	}
	if s.ResourceSettings != nil {
		if err := s.ResourceSettings.Validate(); err != nil {
			return err
		}
	}
	if s.SourceDataSourceSettings != nil {
		for _, item := range s.SourceDataSourceSettings {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.TableMappings != nil {
		for _, item := range s.TableMappings {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.TransformationRules != nil {
		for _, item := range s.TransformationRules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateDIJobRequestDestinationDataSourceSettings struct {
	// The data source name.
	//
	// example:
	//
	// holo_datasource_1
	DataSourceName *string `json:"DataSourceName,omitempty" xml:"DataSourceName,omitempty"`
}

func (s CreateDIJobRequestDestinationDataSourceSettings) String() string {
	return dara.Prettify(s)
}

func (s CreateDIJobRequestDestinationDataSourceSettings) GoString() string {
	return s.String()
}

func (s *CreateDIJobRequestDestinationDataSourceSettings) GetDataSourceName() *string {
	return s.DataSourceName
}

func (s *CreateDIJobRequestDestinationDataSourceSettings) SetDataSourceName(v string) *CreateDIJobRequestDestinationDataSourceSettings {
	s.DataSourceName = &v
	return s
}

func (s *CreateDIJobRequestDestinationDataSourceSettings) Validate() error {
	return dara.Validate(s)
}

type CreateDIJobRequestJobSettings struct {
	// The channel-specific settings. You can configure special settings for specific channels. Currently supported: Holo2Holo (Hologres to Hologres) and Holo2Kafka (Hologres to Kafka).
	//
	// 1.  Holo2Kafka
	//
	// 	- Example: {"destinationChannelSettings":{"kafkaClientProperties":[{"key":"linger.ms","value":"100"}],"keyColumns":["col3"],"writeMode":"canal"}} kafkaClientProperties: Kafka producer parameters used when writing to Kafka.
	//
	// 	- keyColumns: The columns to write to Kafka.
	//
	// 	- writeMode: The Kafka write format. Valid values: json and canal.
	//
	// 2.  Holo2Holo
	//
	// 	- Example: {"destinationChannelSettings":{"conflictMode":"replace","dynamicColumnAction":"replay","writeMode":"replay"}}
	//
	// 	- conflictMode: The conflict handling policy when writing to Hologres. Valid values: replace (overwrite) and ignore.
	//
	// 	- writeMode: The write mode for Hologres. Valid values: replay and insert.
	//
	// 	- dynamicColumnAction: The dynamic column handling mode when writing to Hologres. Valid values: replay, insert, and ignore.
	//
	// example:
	//
	// {"structInfo":"MANAGED","storageType":"TEXTFILE","writeMode":"APPEND","partitionColumns":[{"columnName":"pt","columnType":"STRING","comment":""}],"fieldDelimiter":""}
	ChannelSettings *string `json:"ChannelSettings,omitempty" xml:"ChannelSettings,omitempty"`
	// The array of column type mappings.
	//
	// >  ["ColumnDataTypeSettings":[ { "SourceDataType":"Bigint", "DestinationDataType":"Text" } ]
	ColumnDataTypeSettings []*CreateDIJobRequestJobSettingsColumnDataTypeSettings `json:"ColumnDataTypeSettings,omitempty" xml:"ColumnDataTypeSettings,omitempty" type:"Repeated"`
	// The scheduled task settings.
	CycleScheduleSettings *CreateDIJobRequestJobSettingsCycleScheduleSettings `json:"CycleScheduleSettings,omitempty" xml:"CycleScheduleSettings,omitempty" type:"Struct"`
	// The array of DDL handling settings.
	//
	// >  ["DDLHandlingSettings":[ { "Type":"Insert", "Action":"Normal" } ]
	DdlHandlingSettings []*CreateDIJobRequestJobSettingsDdlHandlingSettings `json:"DdlHandlingSettings,omitempty" xml:"DdlHandlingSettings,omitempty" type:"Repeated"`
	// The runtime settings.
	RuntimeSettings []*CreateDIJobRequestJobSettingsRuntimeSettings `json:"RuntimeSettings,omitempty" xml:"RuntimeSettings,omitempty" type:"Repeated"`
}

func (s CreateDIJobRequestJobSettings) String() string {
	return dara.Prettify(s)
}

func (s CreateDIJobRequestJobSettings) GoString() string {
	return s.String()
}

func (s *CreateDIJobRequestJobSettings) GetChannelSettings() *string {
	return s.ChannelSettings
}

func (s *CreateDIJobRequestJobSettings) GetColumnDataTypeSettings() []*CreateDIJobRequestJobSettingsColumnDataTypeSettings {
	return s.ColumnDataTypeSettings
}

func (s *CreateDIJobRequestJobSettings) GetCycleScheduleSettings() *CreateDIJobRequestJobSettingsCycleScheduleSettings {
	return s.CycleScheduleSettings
}

func (s *CreateDIJobRequestJobSettings) GetDdlHandlingSettings() []*CreateDIJobRequestJobSettingsDdlHandlingSettings {
	return s.DdlHandlingSettings
}

func (s *CreateDIJobRequestJobSettings) GetRuntimeSettings() []*CreateDIJobRequestJobSettingsRuntimeSettings {
	return s.RuntimeSettings
}

func (s *CreateDIJobRequestJobSettings) SetChannelSettings(v string) *CreateDIJobRequestJobSettings {
	s.ChannelSettings = &v
	return s
}

func (s *CreateDIJobRequestJobSettings) SetColumnDataTypeSettings(v []*CreateDIJobRequestJobSettingsColumnDataTypeSettings) *CreateDIJobRequestJobSettings {
	s.ColumnDataTypeSettings = v
	return s
}

func (s *CreateDIJobRequestJobSettings) SetCycleScheduleSettings(v *CreateDIJobRequestJobSettingsCycleScheduleSettings) *CreateDIJobRequestJobSettings {
	s.CycleScheduleSettings = v
	return s
}

func (s *CreateDIJobRequestJobSettings) SetDdlHandlingSettings(v []*CreateDIJobRequestJobSettingsDdlHandlingSettings) *CreateDIJobRequestJobSettings {
	s.DdlHandlingSettings = v
	return s
}

func (s *CreateDIJobRequestJobSettings) SetRuntimeSettings(v []*CreateDIJobRequestJobSettingsRuntimeSettings) *CreateDIJobRequestJobSettings {
	s.RuntimeSettings = v
	return s
}

func (s *CreateDIJobRequestJobSettings) Validate() error {
	if s.ColumnDataTypeSettings != nil {
		for _, item := range s.ColumnDataTypeSettings {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.CycleScheduleSettings != nil {
		if err := s.CycleScheduleSettings.Validate(); err != nil {
			return err
		}
	}
	if s.DdlHandlingSettings != nil {
		for _, item := range s.DdlHandlingSettings {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.RuntimeSettings != nil {
		for _, item := range s.RuntimeSettings {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateDIJobRequestJobSettingsColumnDataTypeSettings struct {
	// The destination type, such as bigint, boolean, string, text, datetime, timestamp, decimal, or binary. Different data sources may have different types.
	//
	// example:
	//
	// text
	DestinationDataType *string `json:"DestinationDataType,omitempty" xml:"DestinationDataType,omitempty"`
	// The source type, such as bigint, boolean, string, text, datetime, timestamp, decimal, or binary. Different data sources may have different types.
	//
	// example:
	//
	// bigint
	SourceDataType *string `json:"SourceDataType,omitempty" xml:"SourceDataType,omitempty"`
}

func (s CreateDIJobRequestJobSettingsColumnDataTypeSettings) String() string {
	return dara.Prettify(s)
}

func (s CreateDIJobRequestJobSettingsColumnDataTypeSettings) GoString() string {
	return s.String()
}

func (s *CreateDIJobRequestJobSettingsColumnDataTypeSettings) GetDestinationDataType() *string {
	return s.DestinationDataType
}

func (s *CreateDIJobRequestJobSettingsColumnDataTypeSettings) GetSourceDataType() *string {
	return s.SourceDataType
}

func (s *CreateDIJobRequestJobSettingsColumnDataTypeSettings) SetDestinationDataType(v string) *CreateDIJobRequestJobSettingsColumnDataTypeSettings {
	s.DestinationDataType = &v
	return s
}

func (s *CreateDIJobRequestJobSettingsColumnDataTypeSettings) SetSourceDataType(v string) *CreateDIJobRequestJobSettingsColumnDataTypeSettings {
	s.SourceDataType = &v
	return s
}

func (s *CreateDIJobRequestJobSettingsColumnDataTypeSettings) Validate() error {
	return dara.Validate(s)
}

type CreateDIJobRequestJobSettingsCycleScheduleSettings struct {
	// The synchronization type that requires scheduling. Valid values:
	//
	// 	- Full: Full synchronization
	//
	// 	- OfflineIncremental: Batch incremental synchronization
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

func (s CreateDIJobRequestJobSettingsCycleScheduleSettings) String() string {
	return dara.Prettify(s)
}

func (s CreateDIJobRequestJobSettingsCycleScheduleSettings) GoString() string {
	return s.String()
}

func (s *CreateDIJobRequestJobSettingsCycleScheduleSettings) GetCycleMigrationType() *string {
	return s.CycleMigrationType
}

func (s *CreateDIJobRequestJobSettingsCycleScheduleSettings) GetScheduleParameters() *string {
	return s.ScheduleParameters
}

func (s *CreateDIJobRequestJobSettingsCycleScheduleSettings) SetCycleMigrationType(v string) *CreateDIJobRequestJobSettingsCycleScheduleSettings {
	s.CycleMigrationType = &v
	return s
}

func (s *CreateDIJobRequestJobSettingsCycleScheduleSettings) SetScheduleParameters(v string) *CreateDIJobRequestJobSettingsCycleScheduleSettings {
	s.ScheduleParameters = &v
	return s
}

func (s *CreateDIJobRequestJobSettingsCycleScheduleSettings) Validate() error {
	return dara.Validate(s)
}

type CreateDIJobRequestJobSettingsDdlHandlingSettings struct {
	// Valid values:
	//
	// 	- Ignore
	//
	// 	- Critical: Fail the task
	//
	// 	- Normal
	//
	// example:
	//
	// Critical
	Action *string `json:"Action,omitempty" xml:"Action,omitempty"`
	// The DDL type. Valid values:
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
	//
	// example:
	//
	// AddColumn
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateDIJobRequestJobSettingsDdlHandlingSettings) String() string {
	return dara.Prettify(s)
}

func (s CreateDIJobRequestJobSettingsDdlHandlingSettings) GoString() string {
	return s.String()
}

func (s *CreateDIJobRequestJobSettingsDdlHandlingSettings) GetAction() *string {
	return s.Action
}

func (s *CreateDIJobRequestJobSettingsDdlHandlingSettings) GetType() *string {
	return s.Type
}

func (s *CreateDIJobRequestJobSettingsDdlHandlingSettings) SetAction(v string) *CreateDIJobRequestJobSettingsDdlHandlingSettings {
	s.Action = &v
	return s
}

func (s *CreateDIJobRequestJobSettingsDdlHandlingSettings) SetType(v string) *CreateDIJobRequestJobSettingsDdlHandlingSettings {
	s.Type = &v
	return s
}

func (s *CreateDIJobRequestJobSettingsDdlHandlingSettings) Validate() error {
	return dara.Validate(s)
}

type CreateDIJobRequestJobSettingsRuntimeSettings struct {
	// The setting name. Valid values:
	//
	// 	- src.offline.datasource.max.connection: Specifies the maximum number of connections that are allowed for reading data from the source of a batch synchronization task.
	//
	// 	- dst.offline.truncate: Specifies whether to clear the destination table before data writing.
	//
	// 	- runtime.offline.speed.limit.enable: Specifies whether throttling is enabled for a batch synchronization task.
	//
	// 	- runtime.offline.concurrent: Specifies the maximum number of parallel threads that are allowed for a batch synchronization task.
	//
	// 	- runtime.enable.auto.create.schema: Specifies whether schemas are automatically created in the destination of a synchronization task.
	//
	// 	- runtime.realtime.concurrent: Specifies the maximum number of parallel threads that are allowed for a real-time synchronization task.
	//
	// 	- runtime.realtime.failover.minute.dataxcdc: Specifies the maximum waiting duration before a synchronization task retries the next restart if the previous restart fails after failover occurs. Unit: minutes.
	//
	// 	- runtime.realtime.failover.times.dataxcdc: Specifies the maximum number of failures that are allowed for restarting a synchronization task after failovers occur.
	//
	// example:
	//
	// runtime.offline.concurrent
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The setting value.
	//
	// example:
	//
	// 1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateDIJobRequestJobSettingsRuntimeSettings) String() string {
	return dara.Prettify(s)
}

func (s CreateDIJobRequestJobSettingsRuntimeSettings) GoString() string {
	return s.String()
}

func (s *CreateDIJobRequestJobSettingsRuntimeSettings) GetName() *string {
	return s.Name
}

func (s *CreateDIJobRequestJobSettingsRuntimeSettings) GetValue() *string {
	return s.Value
}

func (s *CreateDIJobRequestJobSettingsRuntimeSettings) SetName(v string) *CreateDIJobRequestJobSettingsRuntimeSettings {
	s.Name = &v
	return s
}

func (s *CreateDIJobRequestJobSettingsRuntimeSettings) SetValue(v string) *CreateDIJobRequestJobSettingsRuntimeSettings {
	s.Value = &v
	return s
}

func (s *CreateDIJobRequestJobSettingsRuntimeSettings) Validate() error {
	return dara.Validate(s)
}

type CreateDIJobRequestResourceSettings struct {
	// The batch synchronization resources.
	OfflineResourceSettings *CreateDIJobRequestResourceSettingsOfflineResourceSettings `json:"OfflineResourceSettings,omitempty" xml:"OfflineResourceSettings,omitempty" type:"Struct"`
	// The real-time synchronization resources.
	RealtimeResourceSettings *CreateDIJobRequestResourceSettingsRealtimeResourceSettings `json:"RealtimeResourceSettings,omitempty" xml:"RealtimeResourceSettings,omitempty" type:"Struct"`
	// The scheduling resources.
	ScheduleResourceSettings *CreateDIJobRequestResourceSettingsScheduleResourceSettings `json:"ScheduleResourceSettings,omitempty" xml:"ScheduleResourceSettings,omitempty" type:"Struct"`
}

func (s CreateDIJobRequestResourceSettings) String() string {
	return dara.Prettify(s)
}

func (s CreateDIJobRequestResourceSettings) GoString() string {
	return s.String()
}

func (s *CreateDIJobRequestResourceSettings) GetOfflineResourceSettings() *CreateDIJobRequestResourceSettingsOfflineResourceSettings {
	return s.OfflineResourceSettings
}

func (s *CreateDIJobRequestResourceSettings) GetRealtimeResourceSettings() *CreateDIJobRequestResourceSettingsRealtimeResourceSettings {
	return s.RealtimeResourceSettings
}

func (s *CreateDIJobRequestResourceSettings) GetScheduleResourceSettings() *CreateDIJobRequestResourceSettingsScheduleResourceSettings {
	return s.ScheduleResourceSettings
}

func (s *CreateDIJobRequestResourceSettings) SetOfflineResourceSettings(v *CreateDIJobRequestResourceSettingsOfflineResourceSettings) *CreateDIJobRequestResourceSettings {
	s.OfflineResourceSettings = v
	return s
}

func (s *CreateDIJobRequestResourceSettings) SetRealtimeResourceSettings(v *CreateDIJobRequestResourceSettingsRealtimeResourceSettings) *CreateDIJobRequestResourceSettings {
	s.RealtimeResourceSettings = v
	return s
}

func (s *CreateDIJobRequestResourceSettings) SetScheduleResourceSettings(v *CreateDIJobRequestResourceSettingsScheduleResourceSettings) *CreateDIJobRequestResourceSettings {
	s.ScheduleResourceSettings = v
	return s
}

func (s *CreateDIJobRequestResourceSettings) Validate() error {
	if s.OfflineResourceSettings != nil {
		if err := s.OfflineResourceSettings.Validate(); err != nil {
			return err
		}
	}
	if s.RealtimeResourceSettings != nil {
		if err := s.RealtimeResourceSettings.Validate(); err != nil {
			return err
		}
	}
	if s.ScheduleResourceSettings != nil {
		if err := s.ScheduleResourceSettings.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateDIJobRequestResourceSettingsOfflineResourceSettings struct {
	// The CU of the Data Integration resource group used for batch synchronization.
	//
	// example:
	//
	// 2.0
	RequestedCu *float64 `json:"RequestedCu,omitempty" xml:"RequestedCu,omitempty"`
	// The name of the Data Integration resource group used for batch synchronization.
	//
	// example:
	//
	// S_res_group_111_222
	ResourceGroupIdentifier *string `json:"ResourceGroupIdentifier,omitempty" xml:"ResourceGroupIdentifier,omitempty"`
}

func (s CreateDIJobRequestResourceSettingsOfflineResourceSettings) String() string {
	return dara.Prettify(s)
}

func (s CreateDIJobRequestResourceSettingsOfflineResourceSettings) GoString() string {
	return s.String()
}

func (s *CreateDIJobRequestResourceSettingsOfflineResourceSettings) GetRequestedCu() *float64 {
	return s.RequestedCu
}

func (s *CreateDIJobRequestResourceSettingsOfflineResourceSettings) GetResourceGroupIdentifier() *string {
	return s.ResourceGroupIdentifier
}

func (s *CreateDIJobRequestResourceSettingsOfflineResourceSettings) SetRequestedCu(v float64) *CreateDIJobRequestResourceSettingsOfflineResourceSettings {
	s.RequestedCu = &v
	return s
}

func (s *CreateDIJobRequestResourceSettingsOfflineResourceSettings) SetResourceGroupIdentifier(v string) *CreateDIJobRequestResourceSettingsOfflineResourceSettings {
	s.ResourceGroupIdentifier = &v
	return s
}

func (s *CreateDIJobRequestResourceSettingsOfflineResourceSettings) Validate() error {
	return dara.Validate(s)
}

type CreateDIJobRequestResourceSettingsRealtimeResourceSettings struct {
	// The CU of the Data Integration resource group used for real-time synchronization.
	//
	// example:
	//
	// 2.0
	RequestedCu *float64 `json:"RequestedCu,omitempty" xml:"RequestedCu,omitempty"`
	// The name of the Data Integration resource group used for real-time synchronization.
	//
	// example:
	//
	// S_res_group_111_222
	ResourceGroupIdentifier *string `json:"ResourceGroupIdentifier,omitempty" xml:"ResourceGroupIdentifier,omitempty"`
}

func (s CreateDIJobRequestResourceSettingsRealtimeResourceSettings) String() string {
	return dara.Prettify(s)
}

func (s CreateDIJobRequestResourceSettingsRealtimeResourceSettings) GoString() string {
	return s.String()
}

func (s *CreateDIJobRequestResourceSettingsRealtimeResourceSettings) GetRequestedCu() *float64 {
	return s.RequestedCu
}

func (s *CreateDIJobRequestResourceSettingsRealtimeResourceSettings) GetResourceGroupIdentifier() *string {
	return s.ResourceGroupIdentifier
}

func (s *CreateDIJobRequestResourceSettingsRealtimeResourceSettings) SetRequestedCu(v float64) *CreateDIJobRequestResourceSettingsRealtimeResourceSettings {
	s.RequestedCu = &v
	return s
}

func (s *CreateDIJobRequestResourceSettingsRealtimeResourceSettings) SetResourceGroupIdentifier(v string) *CreateDIJobRequestResourceSettingsRealtimeResourceSettings {
	s.ResourceGroupIdentifier = &v
	return s
}

func (s *CreateDIJobRequestResourceSettingsRealtimeResourceSettings) Validate() error {
	return dara.Validate(s)
}

type CreateDIJobRequestResourceSettingsScheduleResourceSettings struct {
	// The CU of the scheduling resource group for batch synchronization tasks.
	//
	// example:
	//
	// 2.0
	RequestedCu *float64 `json:"RequestedCu,omitempty" xml:"RequestedCu,omitempty"`
	// The name of the scheduling resource group for batch synchronization tasks.
	//
	// example:
	//
	// S_res_group_235454102432001_1579085295030
	ResourceGroupIdentifier *string `json:"ResourceGroupIdentifier,omitempty" xml:"ResourceGroupIdentifier,omitempty"`
}

func (s CreateDIJobRequestResourceSettingsScheduleResourceSettings) String() string {
	return dara.Prettify(s)
}

func (s CreateDIJobRequestResourceSettingsScheduleResourceSettings) GoString() string {
	return s.String()
}

func (s *CreateDIJobRequestResourceSettingsScheduleResourceSettings) GetRequestedCu() *float64 {
	return s.RequestedCu
}

func (s *CreateDIJobRequestResourceSettingsScheduleResourceSettings) GetResourceGroupIdentifier() *string {
	return s.ResourceGroupIdentifier
}

func (s *CreateDIJobRequestResourceSettingsScheduleResourceSettings) SetRequestedCu(v float64) *CreateDIJobRequestResourceSettingsScheduleResourceSettings {
	s.RequestedCu = &v
	return s
}

func (s *CreateDIJobRequestResourceSettingsScheduleResourceSettings) SetResourceGroupIdentifier(v string) *CreateDIJobRequestResourceSettingsScheduleResourceSettings {
	s.ResourceGroupIdentifier = &v
	return s
}

func (s *CreateDIJobRequestResourceSettingsScheduleResourceSettings) Validate() error {
	return dara.Validate(s)
}

type CreateDIJobRequestSourceDataSourceSettings struct {
	// The data source name.
	//
	// example:
	//
	// mysql_datasource_1
	DataSourceName *string `json:"DataSourceName,omitempty" xml:"DataSourceName,omitempty"`
	// The data source properties.
	DataSourceProperties *CreateDIJobRequestSourceDataSourceSettingsDataSourceProperties `json:"DataSourceProperties,omitempty" xml:"DataSourceProperties,omitempty" type:"Struct"`
}

func (s CreateDIJobRequestSourceDataSourceSettings) String() string {
	return dara.Prettify(s)
}

func (s CreateDIJobRequestSourceDataSourceSettings) GoString() string {
	return s.String()
}

func (s *CreateDIJobRequestSourceDataSourceSettings) GetDataSourceName() *string {
	return s.DataSourceName
}

func (s *CreateDIJobRequestSourceDataSourceSettings) GetDataSourceProperties() *CreateDIJobRequestSourceDataSourceSettingsDataSourceProperties {
	return s.DataSourceProperties
}

func (s *CreateDIJobRequestSourceDataSourceSettings) SetDataSourceName(v string) *CreateDIJobRequestSourceDataSourceSettings {
	s.DataSourceName = &v
	return s
}

func (s *CreateDIJobRequestSourceDataSourceSettings) SetDataSourceProperties(v *CreateDIJobRequestSourceDataSourceSettingsDataSourceProperties) *CreateDIJobRequestSourceDataSourceSettings {
	s.DataSourceProperties = v
	return s
}

func (s *CreateDIJobRequestSourceDataSourceSettings) Validate() error {
	if s.DataSourceProperties != nil {
		if err := s.DataSourceProperties.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateDIJobRequestSourceDataSourceSettingsDataSourceProperties struct {
	// The database encoding.
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

func (s CreateDIJobRequestSourceDataSourceSettingsDataSourceProperties) String() string {
	return dara.Prettify(s)
}

func (s CreateDIJobRequestSourceDataSourceSettingsDataSourceProperties) GoString() string {
	return s.String()
}

func (s *CreateDIJobRequestSourceDataSourceSettingsDataSourceProperties) GetEncoding() *string {
	return s.Encoding
}

func (s *CreateDIJobRequestSourceDataSourceSettingsDataSourceProperties) GetTimezone() *string {
	return s.Timezone
}

func (s *CreateDIJobRequestSourceDataSourceSettingsDataSourceProperties) SetEncoding(v string) *CreateDIJobRequestSourceDataSourceSettingsDataSourceProperties {
	s.Encoding = &v
	return s
}

func (s *CreateDIJobRequestSourceDataSourceSettingsDataSourceProperties) SetTimezone(v string) *CreateDIJobRequestSourceDataSourceSettingsDataSourceProperties {
	s.Timezone = &v
	return s
}

func (s *CreateDIJobRequestSourceDataSourceSettingsDataSourceProperties) Validate() error {
	return dara.Validate(s)
}

type CreateDIJobRequestTableMappings struct {
	// Each rule can select a set of source objects to synchronize. Multiple rules together select a table.
	SourceObjectSelectionRules []*CreateDIJobRequestTableMappingsSourceObjectSelectionRules `json:"SourceObjectSelectionRules,omitempty" xml:"SourceObjectSelectionRules,omitempty" type:"Repeated"`
	// The list of synchronization object transformation rule definitions. Each element represents a single transformation rule definition.
	TransformationRules []*CreateDIJobRequestTableMappingsTransformationRules `json:"TransformationRules,omitempty" xml:"TransformationRules,omitempty" type:"Repeated"`
}

func (s CreateDIJobRequestTableMappings) String() string {
	return dara.Prettify(s)
}

func (s CreateDIJobRequestTableMappings) GoString() string {
	return s.String()
}

func (s *CreateDIJobRequestTableMappings) GetSourceObjectSelectionRules() []*CreateDIJobRequestTableMappingsSourceObjectSelectionRules {
	return s.SourceObjectSelectionRules
}

func (s *CreateDIJobRequestTableMappings) GetTransformationRules() []*CreateDIJobRequestTableMappingsTransformationRules {
	return s.TransformationRules
}

func (s *CreateDIJobRequestTableMappings) SetSourceObjectSelectionRules(v []*CreateDIJobRequestTableMappingsSourceObjectSelectionRules) *CreateDIJobRequestTableMappings {
	s.SourceObjectSelectionRules = v
	return s
}

func (s *CreateDIJobRequestTableMappings) SetTransformationRules(v []*CreateDIJobRequestTableMappingsTransformationRules) *CreateDIJobRequestTableMappings {
	s.TransformationRules = v
	return s
}

func (s *CreateDIJobRequestTableMappings) Validate() error {
	if s.SourceObjectSelectionRules != nil {
		for _, item := range s.SourceObjectSelectionRules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.TransformationRules != nil {
		for _, item := range s.TransformationRules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateDIJobRequestTableMappingsSourceObjectSelectionRules struct {
	// The selection action. Valid values: Include and Exclude.
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

func (s CreateDIJobRequestTableMappingsSourceObjectSelectionRules) String() string {
	return dara.Prettify(s)
}

func (s CreateDIJobRequestTableMappingsSourceObjectSelectionRules) GoString() string {
	return s.String()
}

func (s *CreateDIJobRequestTableMappingsSourceObjectSelectionRules) GetAction() *string {
	return s.Action
}

func (s *CreateDIJobRequestTableMappingsSourceObjectSelectionRules) GetExpression() *string {
	return s.Expression
}

func (s *CreateDIJobRequestTableMappingsSourceObjectSelectionRules) GetExpressionType() *string {
	return s.ExpressionType
}

func (s *CreateDIJobRequestTableMappingsSourceObjectSelectionRules) GetObjectType() *string {
	return s.ObjectType
}

func (s *CreateDIJobRequestTableMappingsSourceObjectSelectionRules) SetAction(v string) *CreateDIJobRequestTableMappingsSourceObjectSelectionRules {
	s.Action = &v
	return s
}

func (s *CreateDIJobRequestTableMappingsSourceObjectSelectionRules) SetExpression(v string) *CreateDIJobRequestTableMappingsSourceObjectSelectionRules {
	s.Expression = &v
	return s
}

func (s *CreateDIJobRequestTableMappingsSourceObjectSelectionRules) SetExpressionType(v string) *CreateDIJobRequestTableMappingsSourceObjectSelectionRules {
	s.ExpressionType = &v
	return s
}

func (s *CreateDIJobRequestTableMappingsSourceObjectSelectionRules) SetObjectType(v string) *CreateDIJobRequestTableMappingsSourceObjectSelectionRules {
	s.ObjectType = &v
	return s
}

func (s *CreateDIJobRequestTableMappingsSourceObjectSelectionRules) Validate() error {
	return dara.Validate(s)
}

type CreateDIJobRequestTableMappingsTransformationRules struct {
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
	// The rule name. The rule name must be unique for a given combination of action type and target type. The name cannot exceed 50 characters.
	//
	// example:
	//
	// rename_rule_1
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The target type for the action. Valid values:
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

func (s CreateDIJobRequestTableMappingsTransformationRules) String() string {
	return dara.Prettify(s)
}

func (s CreateDIJobRequestTableMappingsTransformationRules) GoString() string {
	return s.String()
}

func (s *CreateDIJobRequestTableMappingsTransformationRules) GetRuleActionType() *string {
	return s.RuleActionType
}

func (s *CreateDIJobRequestTableMappingsTransformationRules) GetRuleName() *string {
	return s.RuleName
}

func (s *CreateDIJobRequestTableMappingsTransformationRules) GetRuleTargetType() *string {
	return s.RuleTargetType
}

func (s *CreateDIJobRequestTableMappingsTransformationRules) SetRuleActionType(v string) *CreateDIJobRequestTableMappingsTransformationRules {
	s.RuleActionType = &v
	return s
}

func (s *CreateDIJobRequestTableMappingsTransformationRules) SetRuleName(v string) *CreateDIJobRequestTableMappingsTransformationRules {
	s.RuleName = &v
	return s
}

func (s *CreateDIJobRequestTableMappingsTransformationRules) SetRuleTargetType(v string) *CreateDIJobRequestTableMappingsTransformationRules {
	s.RuleTargetType = &v
	return s
}

func (s *CreateDIJobRequestTableMappingsTransformationRules) Validate() error {
	return dara.Validate(s)
}

type CreateDIJobRequestTransformationRules struct {
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
	// The rule expression in JSON string format.
	//
	// 1.  Rename rule
	//
	// 	- Example: {"expression":"${srcDatasourceName}_${srcDatabaseName}_0922" }
	//
	// 	- expression: The rename transformation expression. Supported variables include: ${srcDatasourceName} (source data source name), ${srcDatabaseName} (source database name), and ${srcTableName} (source table name).
	//
	// 2.  AddColumn rule
	//
	// 	- Example: {"columns":[{"columnName":"my_add_column","columnValueType":"Constant","columnValue":"123"}]}
	//
	// 	- If not specified, the default behavior is to not add columns.
	//
	// 	- columnName: The name of the column to add.
	//
	// 	- columnValueType: The value type of the column to add. Valid values: Constant and Variable.
	//
	// 	- columnValue: The value of the column to add. When columnValueType is set to Constant, the value is a custom constant of the string type. When columnValueType is set to Variable, the value is a built-in variable. Built-in variables include: EXECUTE_TIME (execution time, long type), DB_NAME_SRC (source database name, string type), DATASOURCE_NAME_SRC (source data source name, string type), TABLE_NAME_SRC (source table name, string type), DB_NAME_DEST (destination database name, string type), DATASOURCE_NAME_DEST (destination data source name, string type), TABLE_NAME_DEST (destination table name, string type), and DB_NAME_SRC_TRANSED (transformed source database name, string type).
	//
	// 3.  DefinePrimaryKey
	//
	// 	- Example: {"columns":["ukcolumn1","ukcolumn2"]}
	//
	// 	- If not specified, the source primary key columns are used by default.
	//
	// 	- When the destination table already exists: Data Integration does not modify the destination table structure. If the specified primary key columns are not in the destination table, the task fails to start.
	//
	// 	- When the destination table is auto-created: Data Integration automatically creates the destination table structure with the defined primary key columns. If the specified primary key columns are not in the destination table, the task fails to start.
	//
	// 4.  HandleDml rule
	//
	// 	- Example of a rule used to process DML messages: {"dmlPolicies":[{"dmlType":"Delete","dmlAction":"Filter","filterCondition":"id > 1"}]}.
	//
	// 	- If not specified, the default rule is Normal for Insert, Update, and Delete.
	//
	// 	- dmlType: The DML operation type. Valid values: Insert, Update, and Delete.
	//
	// 	- dmlAction: The DML handling policy. Valid values: Normal, Ignore, Filter (conditional processing, used when dmlType is Update or Delete), and LogicalDelete.
	//
	// 	- filterCondition: The DML filter condition. This parameter is used when dmlAction is set to Filter.
	//
	// 5.  DefineIncrementalCondition
	//
	// 	- Example: {"where":"id > 0"}
	//
	// 	- Specifies the incremental filter condition.
	//
	// 6.  DefineCycleScheduleSettings
	//
	// 	- Example: {"cronExpress":" \\	- \\	- \\	- \\	- \\	- \\*", "cycleType":"1"}
	//
	// 	- Specifies the scheduled task parameters.
	//
	// 7.  DefinePartitionKey
	//
	// 	- Example: {"columns":["id"]}
	//
	// 	- Specifies the partition key.
	//
	// example:
	//
	// {"expression":"${srcDatasoureName}_${srcDatabaseName}"}
	RuleExpression *string `json:"RuleExpression,omitempty" xml:"RuleExpression,omitempty"`
	// The rule name. When the action type and target type are the same, the rule name must be unique. The name cannot exceed 50 characters.
	//
	// example:
	//
	// rename_rule_1
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The target type for the action. Valid values:
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

func (s CreateDIJobRequestTransformationRules) String() string {
	return dara.Prettify(s)
}

func (s CreateDIJobRequestTransformationRules) GoString() string {
	return s.String()
}

func (s *CreateDIJobRequestTransformationRules) GetRuleActionType() *string {
	return s.RuleActionType
}

func (s *CreateDIJobRequestTransformationRules) GetRuleExpression() *string {
	return s.RuleExpression
}

func (s *CreateDIJobRequestTransformationRules) GetRuleName() *string {
	return s.RuleName
}

func (s *CreateDIJobRequestTransformationRules) GetRuleTargetType() *string {
	return s.RuleTargetType
}

func (s *CreateDIJobRequestTransformationRules) SetRuleActionType(v string) *CreateDIJobRequestTransformationRules {
	s.RuleActionType = &v
	return s
}

func (s *CreateDIJobRequestTransformationRules) SetRuleExpression(v string) *CreateDIJobRequestTransformationRules {
	s.RuleExpression = &v
	return s
}

func (s *CreateDIJobRequestTransformationRules) SetRuleName(v string) *CreateDIJobRequestTransformationRules {
	s.RuleName = &v
	return s
}

func (s *CreateDIJobRequestTransformationRules) SetRuleTargetType(v string) *CreateDIJobRequestTransformationRules {
	s.RuleTargetType = &v
	return s
}

func (s *CreateDIJobRequestTransformationRules) Validate() error {
	return dara.Validate(s)
}
