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
	// The request ID. You can use the ID to locate logs and troubleshoot issues.
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
	if s.PagingInfo != nil {
		if err := s.PagingInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetDIJobResponseBodyPagingInfo struct {
	// Deprecated
	//
	// **[Deprecated]*	- This field is deprecated. Use the Id field instead.
	//
	// example:
	//
	// 32601
	DIJobId *string `json:"DIJobId,omitempty" xml:"DIJobId,omitempty"`
	// The description.
	//
	// example:
	//
	// description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The list of destination data source properties.
	DestinationDataSourceSettings []*GetDIJobResponseBodyPagingInfoDestinationDataSourceSettings `json:"DestinationDataSourceSettings,omitempty" xml:"DestinationDataSourceSettings,omitempty" type:"Repeated"`
	// The type of the destination data source. Valid values: Hologres, OSS-HDFS, OSS, MaxCompute, LogHub, StarRocks, DataHub, AnalyticDB_For_MySQL, Kafka, and Hive.
	//
	// example:
	//
	// Hologres
	DestinationDataSourceType *string `json:"DestinationDataSourceType,omitempty" xml:"DestinationDataSourceType,omitempty"`
	// The task ID.
	//
	// example:
	//
	// 32601
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The task name.
	//
	// example:
	//
	// imp_ods_dms_det_dealer_info_df
	JobName *string `json:"JobName,omitempty" xml:"JobName,omitempty"`
	// The runtime settings.
	JobSettings *GetDIJobResponseBodyPagingInfoJobSettings `json:"JobSettings,omitempty" xml:"JobSettings,omitempty" type:"Struct"`
	// The task status. Valid values:
	//
	// - Finished: the task ran successfully and has finished.
	//
	// - Failed: the task failed to run.
	//
	// - Running: the task is running.
	//
	// - Initialized: the task is initialized but not started.
	//
	// - Stopping: the task is being stopped.
	//
	// - Stop: the task is stopped.
	//
	// example:
	//
	// Running
	JobStatus *string `json:"JobStatus,omitempty" xml:"JobStatus,omitempty"`
	// The task type. Valid values:
	//
	// - DatabaseRealtimeMigration: performs stream synchronization for multiple tables across multiple source databases. Supports full data only, incremental data only, or both full and incremental data.
	//
	// - DatabaseOfflineMigration: performs batch synchronization for multiple tables across multiple source databases. Supports full data only, incremental data only, or both full and incremental data.
	//
	// - SingleTableRealtimeMigration: performs stream synchronization for a single source table.
	//
	// example:
	//
	// DatabaseRealtimeMigration
	JobType *string `json:"JobType,omitempty" xml:"JobType,omitempty"`
	// The synchronization type. Valid values:
	//
	// - FullAndRealtimeIncremental: full and real-time incremental synchronization (real-time synchronization for the entire database).
	//
	// - RealtimeIncremental: real-time incremental synchronization (real-time synchronization for a single table).
	//
	// - Full: full data synchronization (batch synchronization for the entire database).
	//
	// - OfflineIncremental: batch incremental synchronization (batch synchronization for the entire database).
	//
	// - FullAndOfflineIncremental: full and batch incremental synchronization (batch synchronization for the entire database).
	//
	// example:
	//
	// FullAndRealtimeIncremental
	MigrationType *string `json:"MigrationType,omitempty" xml:"MigrationType,omitempty"`
	// The task owner.
	//
	// example:
	//
	// 3726346
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The ID of the DataWorks workspace. You can log on to the [DataWorks console](https://workbench.data.aliyun.com/console) and go to the workspace configuration page to obtain the workspace ID.
	//
	// This parameter specifies the DataWorks workspace used for this API call.
	//
	// example:
	//
	// 98330
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The resource settings.
	ResourceSettings *GetDIJobResponseBodyPagingInfoResourceSettings `json:"ResourceSettings,omitempty" xml:"ResourceSettings,omitempty" type:"Struct"`
	// The array of source data source settings.
	SourceDataSourceSettings []*GetDIJobResponseBodyPagingInfoSourceDataSourceSettings `json:"SourceDataSourceSettings,omitempty" xml:"SourceDataSourceSettings,omitempty" type:"Repeated"`
	// The type of the source data source. Valid values: PolarDB, MySQL, Kafka, LogHub, Hologres, Oracle, OceanBase, MongoDB, RedShift, Hive, SQLServer, Doris, and ClickHouse.
	//
	// example:
	//
	// Mysql
	SourceDataSourceType *string `json:"SourceDataSourceType,omitempty" xml:"SourceDataSourceType,omitempty"`
	// The list of synchronization object transformation mappings. Each element in the list describes a set of source object selection rules and a list of transformation rules applied to the selected synchronization objects.
	//
	// > [
	//
	//   {
	//
	//     "SourceObjectSelectionRules":[
	//
	//       {
	//
	//         "ObjectType":"Database",
	//
	//         "Action":"Include",
	//
	//         "ExpressionType":"Exact",
	//
	//         "Expression":"biz_db"
	//
	//       },
	//
	//     {
	//
	//         "ObjectType":"Schema",
	//
	//         "Action":"Include",
	//
	//         "ExpressionType":"Exact",
	//
	//         "Expression":"s1"
	//
	//       },
	//
	//     {
	//
	//         "ObjectType":"Table",
	//
	//         "Action":"Include",
	//
	//         "ExpressionType":"Exact",
	//
	//         "Expression":"table1"
	//
	//       }
	//
	//     ],
	//
	//     "TransformationRuleNames":[
	//
	//       {
	//
	//         "RuleName":"my_database_rename_rule",
	//
	//         "RuleActionType":"Rename",
	//
	//         "RuleTargetType":"Schema"
	//
	//   		}
	//
	//     ]
	//
	//   }
	//
	// ]
	TableMappings []*GetDIJobResponseBodyPagingInfoTableMappings `json:"TableMappings,omitempty" xml:"TableMappings,omitempty" type:"Repeated"`
	// The list of transformation rule definitions for synchronization objects.
	//
	// > [
	//
	//   {
	//
	//     "RuleName":"my_database_rename_rule",
	//
	//     "RuleActionType":"Rename",
	//
	//     "RuleTargetType":"Schema",
	//
	//     "RuleExpression":"{\\"expression\\":\\"${srcDatasoureName}_${srcDatabaseName}\\"}"
	//
	//   }
	//
	// ]
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

func (s *GetDIJobResponseBodyPagingInfo) GetOwner() *string {
	return s.Owner
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

func (s *GetDIJobResponseBodyPagingInfo) SetOwner(v string) *GetDIJobResponseBodyPagingInfo {
	s.Owner = &v
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

type GetDIJobResponseBodyPagingInfoDestinationDataSourceSettings struct {
	// The name of the destination data source.
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
	// The channel-related task settings. You can configure specific channels. Currently, Holo2Holo (synchronization from Hologres to Hologres) and Holo2Kafka (synchronization from Hologres to Kafka) are supported.
	//
	// 1. Holo2Kafka
	//
	// - Example: {"destinationChannelSettings":{"kafkaClientProperties":[{"key":"linger.ms","value":"100"}],"keyColumns":["col3"],"writeMode":"canal"}}
	//
	// - kafkaClientProperties: the Kafka producer parameters used when writing data to Kafka.
	//
	// - keyColumns: the Kafka columns to which data is written.
	//
	// - writeMode: the format in which data is written to Kafka. Valid values: json and canal.
	//
	// 2. Holo2Holo
	//
	// - Example: {"destinationChannelSettings":{"conflictMode":"replace","dynamicColumnAction":"replay","writeMode":"replay"}}
	//
	// - conflictMode: the conflict handling policy when writing data to Hologres. Valid values: replace (overwrite) and ignore (ignore).
	//
	// - writeMode: the method used to write data to Hologres. Valid values: replay and insert.
	//
	// - dynamicColumnAction: the method used to write dynamic columns to Hologres. Valid values: replay, insert, and ignore.
	//
	// example:
	//
	// {"structInfo":"MANAGED","storageType":"TEXTFILE","writeMode":"APPEND","partitionColumns":[{"columnName":"pt","columnType":"STRING","comment":""}],"fieldDelimiter":""}
	ChannelSettings *string `json:"ChannelSettings,omitempty" xml:"ChannelSettings,omitempty"`
	// The column type mappings.
	ColumnDataTypeSettings []*GetDIJobResponseBodyPagingInfoJobSettingsColumnDataTypeSettings `json:"ColumnDataTypeSettings,omitempty" xml:"ColumnDataTypeSettings,omitempty" type:"Repeated"`
	// The periodic scheduling settings.
	CycleScheduleSettings *GetDIJobResponseBodyPagingInfoJobSettingsCycleScheduleSettings `json:"CycleScheduleSettings,omitempty" xml:"CycleScheduleSettings,omitempty" type:"Struct"`
	// The array of DDL handling settings, which includes a list of DDL types to be handled and their corresponding rules.
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

type GetDIJobResponseBodyPagingInfoJobSettingsColumnDataTypeSettings struct {
	// The destination data type, such as bigint, boolean, string, text, datetime, timestamp, decimal, and binary. Data types vary based on the data source type.
	//
	// example:
	//
	// text
	DestinationDataType *string `json:"DestinationDataType,omitempty" xml:"DestinationDataType,omitempty"`
	// The source data type, such as bigint, boolean, string, text, datetime, timestamp, decimal, and binary. Data types vary based on the data source type.
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
	// - Full: full data synchronization.
	//
	// - OfflineIncremental: batch incremental synchronization.
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
	// The handling action. Valid values:
	//
	// - Ignore: ignores the DDL statement.
	//
	// - Critical: reports an error.
	//
	// - Normal: processes the DDL statement normally.
	//
	// example:
	//
	// Ignore
	Action *string `json:"Action,omitempty" xml:"Action,omitempty"`
	// The DDL type. Valid values:
	//
	// - RenameColumn: renames a column.
	//
	// - ModifyColumn: modifies a column.
	//
	// - CreateTable: creates a table.
	//
	// - TruncateTable: truncates a table.
	//
	// - DropTable: drops a table.
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
	// The setting name. Valid values:
	//
	// - src.offline.datasource.max.connection: the maximum number of connections to the source for batch tasks.
	//
	// - dst.offline.truncate: specifies whether to truncate the destination table.
	//
	// - runtime.offline.speed.limit.enable: specifies whether to enable throttling for batch tasks.
	//
	// - runtime.offline.concurrent: the concurrency of batch synchronization tasks.
	//
	// - runtime.enable.auto.create.schema: specifies whether to automatically create a schema at the destination.
	//
	// - runtime.realtime.concurrent: the concurrency of real-time tasks.
	//
	// - runtime.realtime.failover.minute.dataxcdc: the wait time in minutes before restarting after a failover failure.
	//
	// - runtime.realtime.failover.times.dataxcdc: the number of restart attempts after a failover failure.
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
	// The batch synchronization resources.
	OfflineResourceSettings *GetDIJobResponseBodyPagingInfoResourceSettingsOfflineResourceSettings `json:"OfflineResourceSettings,omitempty" xml:"OfflineResourceSettings,omitempty" type:"Struct"`
	// The real-time synchronization resources.
	RealtimeResourceSettings *GetDIJobResponseBodyPagingInfoResourceSettingsRealtimeResourceSettings `json:"RealtimeResourceSettings,omitempty" xml:"RealtimeResourceSettings,omitempty" type:"Struct"`
	// The scheduling resource.
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

type GetDIJobResponseBodyPagingInfoResourceSettingsOfflineResourceSettings struct {
	// The number of compute units (CUs) of the scheduling resource group used by the batch synchronization task.
	//
	// example:
	//
	// 2.0
	RequestedCu *float64 `json:"RequestedCu,omitempty" xml:"RequestedCu,omitempty"`
	// The name of the data integration resource group used for batch synchronization.
	//
	// example:
	//
	// di_resourcegroup_v1
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
	// The number of CUs of the data integration resource group used for real-time synchronization.
	//
	// example:
	//
	// 2.0
	RequestedCu *float64 `json:"RequestedCu,omitempty" xml:"RequestedCu,omitempty"`
	// The name of the data integration resource group used for real-time tasks.
	//
	// example:
	//
	// di_resourcegroup_v1
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
	// The number of CUs of the data integration resource group used by the batch scheduling task.
	//
	// example:
	//
	// 2.0
	RequestedCu *float64 `json:"RequestedCu,omitempty" xml:"RequestedCu,omitempty"`
	// The name of the scheduling resource group used by the batch scheduling task.
	//
	// example:
	//
	// schedual_resourcegroup_v1
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
	if s.DataSourceProperties != nil {
		if err := s.DataSourceProperties.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetDIJobResponseBodyPagingInfoSourceDataSourceSettingsDataSourceProperties struct {
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
	// Each rule selects a set of source objects to be synchronized. Multiple rules combine to select a single table.
	SourceObjectSelectionRules []*GetDIJobResponseBodyPagingInfoTableMappingsSourceObjectSelectionRules `json:"SourceObjectSelectionRules,omitempty" xml:"SourceObjectSelectionRules,omitempty" type:"Repeated"`
	// The list of transformation rule definitions for synchronization objects. Each element in the list is a transformation rule definition.
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

type GetDIJobResponseBodyPagingInfoTableMappingsSourceObjectSelectionRules struct {
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
	// - Table: table.
	//
	// - Schema: schema.
	//
	// - Database: database.
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
	// - DefinePrimaryKey: defines the primary key.
	//
	// - Rename: renames the object.
	//
	// - AddColumn: adds a column.
	//
	// - HandleDml: handles DML operations.
	//
	// example:
	//
	// AddColumn
	RuleActionType *string `json:"RuleActionType,omitempty" xml:"RuleActionType,omitempty"`
	// The rule name. The rule name must be unique when the action type and the target type of the action are the same.
	//
	// example:
	//
	// rename_rule_1
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The target type of the action. Valid values:
	//
	// - Table: table.
	//
	// - Schema: schema.
	//
	// - Database: database.
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
	// - DefinePrimaryKey: defines the primary key.
	//
	// - Rename: renames the object.
	//
	// - AddColumn: adds a column.
	//
	// - HandleDml: handles DML operations.
	//
	// - DefineIncrementalCondition: defines the incremental condition.
	//
	// - DefineCycleScheduleSettings: defines the periodic scheduling settings.
	//
	// - DefinePartitionKey: defines the partition key.
	//
	// example:
	//
	// Rename
	RuleActionType *string `json:"RuleActionType,omitempty" xml:"RuleActionType,omitempty"`
	// The rule expression in JSON string format.
	//
	// 1. Rename rule (Rename)
	//
	// - Example: {"expression":"${srcDatasourceName}_${srcDatabaseName}_0922"}
	//
	// - expression: the expression for the rename transformation rule. Supported variables in the expression include: ${srcDatasourceName} (source data source name), ${srcDatabaseName} (source database name), and ${srcTableName} (source table name).
	//
	// 2. Add column rule (AddColumn)
	//
	// - Example: {"columns":[{"columnName":"my_add_column","columnValueType":"Constant","columnValue":"123"}]}
	//
	// - If not specified, the default rule is to add no columns and perform no replication.
	//
	// - columnName: the name of the appended column.
	//
	// - columnValueType: the value type of the appended column. Valid values: Constant and Variable.
	//
	// - columnValue: the value of the appended column. When columnValueType is Constant, the value is a custom constant of the String type. When columnValueType is Variable, the value is a built-in variable. Valid built-in variables include: EXECUTE_TIME (execution time, Long type), DB_NAME_SRC (source database name, String type), DATASOURCE_NAME_SRC (source data source name, String type), TABLE_NAME_SRC (source table name, String type), DB_NAME_DEST (destination database name, String type), DATASOURCE_NAME_DEST (destination data source name, String type), TABLE_NAME_DEST (destination table name, String type), and DB_NAME_SRC_TRANSED (transformed database name, String type).
	//
	// 3. Specify the primary key column names for the destination table (DefinePrimaryKey)
	//
	// - Example: {"columns":["ukcolumn1","ukcolumn2"]}
	//
	// - If not specified, the source primary key columns are used by default.
	//
	// - When the destination is an existing table: the data integration system does not modify the destination table schema. If the specified primary key columns are not in the destination column set, an error is reported when the task starts.
	//
	// - When the destination is a table to be created automatically: the data integration system automatically creates the destination table schema, which includes the defined primary key columns. If the specified primary key columns are not in the destination column set, an error is reported when the task starts.
	//
	// 4. DML handling rule (HandleDml)
	//
	// - Example:
	//
	// {"dmlPolicies":[{"dmlType":"Delete","dmlAction":"Filter","filterCondition":"id > 1"}]}
	//
	// - If not specified, the default rule is Normal for Insert, Update, and Delete.
	//
	// - dmlType: the DML operations type. Valid values: Insert, Update, and Delete.
	//
	// - dmlAction: the DML handling policy. Valid values: Normal (normal processing), Ignore (ignore), Filter (conditional normal processing, used when dmlType is Update or Delete), and LogicalDelete (logical deletion).
	//
	// - filterCondition: the DML filter condition, used when dmlAction is Filter.
	//
	// 5. Incremental condition (DefineIncrementalCondition)
	//
	// - Example:
	//
	// {"where":"id > 0"}
	//
	// - Specifies the incremental filter condition.
	//
	// 6. Periodic scheduling parameters (DefineCycleScheduleSettings)
	//
	// - Example:
	//
	// {"cronExpress":" 	- 	- 	- 	- 	- *", "cycleType":"1"}
	//
	// - Specifies the scheduling parameters for the scheduled task.
	//
	// 7. Specify the partition key (DefinePartitionKey)
	//
	// - Example:
	//
	// {"columns":["id"]}
	//
	// - Specifies the partition key.
	//
	// example:
	//
	// {"expression":"${srcDatasoureName}_${srcDatabaseName}"}
	RuleExpression *string `json:"RuleExpression,omitempty" xml:"RuleExpression,omitempty"`
	// The rule name. The rule name must be unique under a specific combination of action type and target type.
	//
	// example:
	//
	// rename_rule_1
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The target type of the action. Valid values:
	//
	// - Table: table.
	//
	// - Schema: schema.
	//
	// - Database: database.
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
