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
	// The details of the data integration job.
	PagingInfo *GetDIJobResponseBodyPagingInfo `json:"PagingInfo,omitempty" xml:"PagingInfo,omitempty" type:"Struct"`
	// The request ID. You can use this ID to locate logs and troubleshoot issues.
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
	// This field is deprecated. Use the `Id` field instead.
	//
	// example:
	//
	// 32601
	DIJobId *string `json:"DIJobId,omitempty" xml:"DIJobId,omitempty"`
	// The description of the job.
	//
	// example:
	//
	// description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The settings for the destination data source.
	DestinationDataSourceSettings []*GetDIJobResponseBodyPagingInfoDestinationDataSourceSettings `json:"DestinationDataSourceSettings,omitempty" xml:"DestinationDataSourceSettings,omitempty" type:"Repeated"`
	// The type of the destination data source. Valid values: `Hologres`, `OSS-HDFS`, `OSS`, `MaxCompute`, `LogHub`, `StarRocks`, `DataHub`, `AnalyticDB for MySQL`, `Kafka`, and `Hive`.
	//
	// example:
	//
	// Hologres
	DestinationDataSourceType *string `json:"DestinationDataSourceType,omitempty" xml:"DestinationDataSourceType,omitempty"`
	// The job ID.
	//
	// example:
	//
	// 32601
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the job.
	//
	// example:
	//
	// imp_ods_dms_det_dealer_info_df
	JobName *string `json:"JobName,omitempty" xml:"JobName,omitempty"`
	// The job settings.
	JobSettings *GetDIJobResponseBodyPagingInfoJobSettings `json:"JobSettings,omitempty" xml:"JobSettings,omitempty" type:"Struct"`
	// The status of the job. Valid values:
	//
	// - `Finished`: The job is complete.
	//
	// - `Failed`: The job failed.
	//
	// - `Running`: The job is running.
	//
	// - `Initialized`: The job is initialized but has not started.
	//
	// - `Stopping`: The job is being stopped.
	//
	// - `Stop`: The job is stopped.
	//
	// example:
	//
	// Running
	JobStatus *string `json:"JobStatus,omitempty" xml:"JobStatus,omitempty"`
	// The job type.
	//
	// - `DatabaseRealtimeMigration`: real-time synchronization of multiple tables from multiple source databases. This type supports full, incremental, or both full and incremental synchronization.
	//
	// - `DatabaseOfflineMigration`: batch synchronization of multiple tables from multiple source databases. This type supports full, incremental, or both full and incremental synchronization.
	//
	// - `SingleTableRealtimeMigration`: real-time synchronization of a single source table.
	//
	// example:
	//
	// DatabaseRealtimeMigration
	JobType *string `json:"JobType,omitempty" xml:"JobType,omitempty"`
	// The synchronization type. Valid values:
	//
	// - `FullAndRealtimeIncremental`: one-time full synchronization and real-time incremental synchronization (for an entire database).
	//
	// - `RealtimeIncremental`: real-time incremental synchronization (for a single table).
	//
	// - `Full`: one-time full synchronization (for an entire database).
	//
	// - `OfflineIncremental`: offline incremental synchronization (for an entire database).
	//
	// - `FullAndOfflineIncremental`: one-time full synchronization and offline incremental synchronization (for an entire database).
	//
	// example:
	//
	// FullAndRealtimeIncremental
	MigrationType *string `json:"MigrationType,omitempty" xml:"MigrationType,omitempty"`
	Owner         *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The ID of the DataWorks workspace for the API call. You can obtain the workspace ID from the Workspace Configuration page in the [DataWorks console](https://workbench.data.aliyun.com/console).
	//
	// example:
	//
	// 98330
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The resource settings.
	ResourceSettings *GetDIJobResponseBodyPagingInfoResourceSettings `json:"ResourceSettings,omitempty" xml:"ResourceSettings,omitempty" type:"Struct"`
	// The settings for the source data source.
	SourceDataSourceSettings []*GetDIJobResponseBodyPagingInfoSourceDataSourceSettings `json:"SourceDataSourceSettings,omitempty" xml:"SourceDataSourceSettings,omitempty" type:"Repeated"`
	// The type of the source data source. Valid values: `PolarDB`, `MySQL`, `Kafka`, `LogHub`, `Hologres`, `Oracle`, `OceanBase`, `MongoDB`, `RedShift`, `Hive`, `SQLServer`, `Doris`, and `ClickHouse`.
	//
	// example:
	//
	// Mysql
	SourceDataSourceType *string `json:"SourceDataSourceType,omitempty" xml:"SourceDataSourceType,omitempty"`
	// A list of mappings for object transformation. Each element in the list describes a set of selection rules for source objects and a set of transformation rules that apply to the selected objects.
	//
	// > [
	//
	// > {
	//
	// > "SourceObjectSelectionRules":[
	//
	// > {
	//
	// > "ObjectType":"Database",
	//
	// > "Action":"Include",
	//
	// > "ExpressionType":"Exact",
	//
	// > "Expression":"biz_db"
	//
	// > },
	//
	// > {
	//
	// > "ObjectType":"Schema",
	//
	// > "Action":"Include",
	//
	// > "ExpressionType":"Exact",
	//
	// > "Expression":"s1"
	//
	// > },
	//
	// > {
	//
	// > "ObjectType":"Table",
	//
	// > "Action":"Include",
	//
	// > "ExpressionType":"Exact",
	//
	// > "Expression":"table1"
	//
	// > }
	//
	// > ],
	//
	// > "TransformationRuleNames":[
	//
	// > {
	//
	// > "RuleName":"my_database_rename_rule",
	//
	// > "RuleActionType":"Rename",
	//
	// > "RuleTargetType":"Schema"
	//
	// > }
	//
	// > ]
	//
	// > }
	//
	// > ]
	TableMappings []*GetDIJobResponseBodyPagingInfoTableMappings `json:"TableMappings,omitempty" xml:"TableMappings,omitempty" type:"Repeated"`
	// A list of definitions for object transformation rules.
	//
	// > [
	//
	// > {
	//
	// > "RuleName":"my_database_rename_rule",
	//
	// > "RuleActionType":"Rename",
	//
	// > "RuleTargetType":"Schema",
	//
	// > "RuleExpression":"{\\\\"expression\\\\":\\\\"${srcDatasoureName}_${srcDatabaseName}\\\\"}"
	//
	// > }
	//
	// > ]
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
	// The settings for channel-related jobs. You can configure special settings for specific channels. The following channels are supported: Holo2Holo (data synchronization from Hologres to Hologres) and Holo2Kafka (data synchronization from Hologres to Kafka).
	//
	// 1. Holo2Kafka
	//
	// - Example: `{"destinationChannelSettings":{"kafkaClientProperties":[{"key":"linger.ms","value":"100"}],"keyColumns":["col3"],"writeMode":"canal"}}`
	//
	// - `kafkaClientProperties`: The parameters for the Kafka producer, used when writing data to Kafka.
	//
	// - `keyColumns`: The columns whose values are used as the key for Kafka records.
	//
	// - `writeMode`: The format for writing data to Kafka. Valid values: `json` and `canal`.
	//
	// 2. Holo2Holo
	//
	// - Example: `{"destinationChannelSettings":{"conflictMode":"replace","dynamicColumnAction":"replay","writeMode":"replay"}}`
	//
	// - `conflictMode`: The conflict handling policy for writing data to Hologres. Valid values: `replace` (overwrite) and `ignore` (ignore).
	//
	// - `writeMode`: The method for writing data to Hologres. Valid values: `replay` and `insert`.
	//
	// - `dynamicColumnAction`: The method for handling dynamic columns when writing data to Hologres. Valid values: `replay`, `insert`, and `ignore`.
	//
	// example:
	//
	// {"structInfo":"MANAGED","storageType":"TEXTFILE","writeMode":"APPEND","partitionColumns":[{"columnName":"pt","columnType":"STRING","comment":""}],"fieldDelimiter":""}
	ChannelSettings *string `json:"ChannelSettings,omitempty" xml:"ChannelSettings,omitempty"`
	// The column data type mappings.
	ColumnDataTypeSettings []*GetDIJobResponseBodyPagingInfoJobSettingsColumnDataTypeSettings `json:"ColumnDataTypeSettings,omitempty" xml:"ColumnDataTypeSettings,omitempty" type:"Repeated"`
	// The settings for periodic scheduling.
	CycleScheduleSettings *GetDIJobResponseBodyPagingInfoJobSettingsCycleScheduleSettings `json:"CycleScheduleSettings,omitempty" xml:"CycleScheduleSettings,omitempty" type:"Struct"`
	// An array of settings for handling DDL messages. Each element specifies a DDL message type and the corresponding handling rule.
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
	// The data type in the destination, such as `bigint`, `boolean`, `string`, `text`, `datetime`, `timestamp`, `decimal`, and `binary`. Data types vary depending on the data source.
	//
	// example:
	//
	// text
	DestinationDataType *string `json:"DestinationDataType,omitempty" xml:"DestinationDataType,omitempty"`
	// The data type in the source, such as `bigint`, `boolean`, `string`, `text`, `datetime`, `timestamp`, `decimal`, and `binary`. Data types vary depending on the data source.
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
	// The synchronization type for periodic scheduling. Valid values:
	//
	// - `Full`: full
	//
	// - `OfflineIncremental`: offline incremental
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
	// - `Ignore`: Ignores the DDL message.
	//
	// - `Critical`: Reports an error.
	//
	// - `Normal`: Processes the DDL message.
	//
	// example:
	//
	// Ignore
	Action *string `json:"Action,omitempty" xml:"Action,omitempty"`
	// The DDL message type. Valid values:
	//
	// - `RenameColumn`
	//
	// - `ModifyColumn`
	//
	// - `CreateTable`
	//
	// - `TruncateTable`
	//
	// - `DropTable`
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
	// The name of the setting. Valid values:
	//
	// - `src.offline.datasource.max.connection`: the maximum number of connections to the source for an offline batch job.
	//
	// - `dst.offline.truncate`: Whether to truncate the destination table before the offline batch job starts.
	//
	// - `runtime.offline.speed.limit.enable`: Whether to enable throttling for an offline batch job.
	//
	// - `runtime.offline.concurrent`: the concurrency level for an offline batch synchronization job.
	//
	// - `runtime.enable.auto.create.schema`: Whether to automatically create a schema at the destination.
	//
	// - `runtime.realtime.concurrent`: the concurrency level for a real-time job.
	//
	// - `runtime.realtime.failover.minute.dataxcdc`: The wait duration (in minutes) before restarting a failed instance.
	//
	// - `runtime.realtime.failover.times.dataxcdc`: The maximum number of retries for a failed instance.
	//
	// example:
	//
	// runtime.offline.concurrent
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The value of the setting.
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
	// The resource settings for the offline synchronization job.
	OfflineResourceSettings *GetDIJobResponseBodyPagingInfoResourceSettingsOfflineResourceSettings `json:"OfflineResourceSettings,omitempty" xml:"OfflineResourceSettings,omitempty" type:"Struct"`
	// The resource settings for the real-time synchronization job.
	RealtimeResourceSettings *GetDIJobResponseBodyPagingInfoResourceSettingsRealtimeResourceSettings `json:"RealtimeResourceSettings,omitempty" xml:"RealtimeResourceSettings,omitempty" type:"Struct"`
	// The scheduling resource settings.
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
	// The number of CUs from the data integration resource group for the offline synchronization job.
	//
	// example:
	//
	// 2.0
	RequestedCu *float64 `json:"RequestedCu,omitempty" xml:"RequestedCu,omitempty"`
	// The name of the data integration resource group used by the offline synchronization job.
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
	// The number of CUs from the data integration resource group for the real-time synchronization job.
	//
	// example:
	//
	// 2.0
	RequestedCu *float64 `json:"RequestedCu,omitempty" xml:"RequestedCu,omitempty"`
	// The name of the data integration resource group used by the real-time job.
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
	// The number of CUs from the scheduling resource group for the offline scheduling job.
	//
	// example:
	//
	// 2.0
	RequestedCu *float64 `json:"RequestedCu,omitempty" xml:"RequestedCu,omitempty"`
	// The name of the scheduling resource group used by the offline scheduling job.
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
	// The encoding of the database.
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
	// Each rule selects a set of source objects to be synchronized. A combination of multiple rules selects one table.
	SourceObjectSelectionRules []*GetDIJobResponseBodyPagingInfoTableMappingsSourceObjectSelectionRules `json:"SourceObjectSelectionRules,omitempty" xml:"SourceObjectSelectionRules,omitempty" type:"Repeated"`
	// An array of object transformation rule definitions.
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
	// The selection action. Valid values: `Include` and `Exclude`.
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
	// The expression type. Valid values: `Exact` and `Regex`.
	//
	// example:
	//
	// Exact
	ExpressionType *string `json:"ExpressionType,omitempty" xml:"ExpressionType,omitempty"`
	// The object type. Valid values:
	//
	// - `Table`
	//
	// - `Schema`
	//
	// - `Database`
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
	// - `DefinePrimaryKey`
	//
	// - `Rename`
	//
	// - `AddColumn`
	//
	// - `HandleDml`
	//
	// example:
	//
	// AddColumn
	RuleActionType *string `json:"RuleActionType,omitempty" xml:"RuleActionType,omitempty"`
	// The name of the rule. The rule name must be unique for a specific action type (`RuleActionType`) and target type (`RuleTargetType`).
	//
	// example:
	//
	// rename_rule_1
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The target object type of the action. Valid values:
	//
	// - `Table`
	//
	// - `Schema`
	//
	// - `Database`
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
	// - `DefinePrimaryKey`
	//
	// - `Rename`
	//
	// - `AddColumn`
	//
	// - `HandleDml`
	//
	// - `DefineIncrementalCondition`
	//
	// - `DefineCycleScheduleSettings`
	//
	// - `DefinePartitionKey`
	//
	// example:
	//
	// Rename
	RuleActionType *string `json:"RuleActionType,omitempty" xml:"RuleActionType,omitempty"`
	// The rule expression, in JSON string format.
	//
	// 1. Rename rule (`Rename`)
	//
	// - Example: `{"expression":"${srcDatasourceName}_${srcDatabaseName}_0922"}`
	//
	// - `expression`: The expression for the rename transformation rule. The expression supports the following variables: `${srcDatasourceName}` (source data source name), `${srcDatabaseName}` (source database name), and `${srcTableName}` (source table name).
	//
	// 2. Add column rule (`AddColumn`)
	//
	// - Example: `{"columns":[{"columnName":"my_add_column","columnValueType":"Constant","columnValue":"123"}]}`
	//
	// - If you do not specify this parameter, no columns are added or copied by default.
	//
	// - `columnName`: The name of the column to add.
	//
	// - `columnValueType`: The value type of the added column. Valid values: `Constant` and `Variable`.
	//
	// - `columnValue`: The value of the added column. If `columnValueType` is `Constant`, the value is a custom string constant. If `columnValueType` is `Variable`, the value is a built-in variable. Valid built-in variables: `EXECUTE_TIME` (execution time, Long), `DB_NAME_SRC` (source database name, String), `DATASOURCE_NAME_SRC` (source data source name, String), `TABLE_NAME_SRC` (source table name, String), `DB_NAME_DEST` (destination database name, String), `DATASOURCE_NAME_DEST` (destination data source name, String), `TABLE_NAME_DEST` (destination table name, String), and `DB_NAME_SRC_TRANSED` (converted database name, String).
	//
	// 3. Define primary key rule (`DefinePrimaryKey`)
	//
	// - Example: `{"columns":["ukcolumn1","ukcolumn2"]}`
	//
	// - By default, the primary key columns from the source table are used.
	//
	// - If the destination table already exists, the data integration system does not modify the table schema. If the specified primary key columns are not in the destination table, the job fails to start.
	//
	// - If the destination table is automatically created, the data integration system automatically creates the table schema that includes the defined primary key columns. If the specified primary key columns are not in the destination table, the job fails to start.
	//
	// 4. DML handling rule (`HandleDml`)
	//
	// - Example: `{"dmlPolicies":[{"dmlType":"Delete","dmlAction":"Filter","filterCondition":"id > 1"}]}`
	//
	// - If you do not specify this parameter, the default value `Normal` is used for Insert, Update, and Delete operations.
	//
	// - `dmlType`: The DML operation type. Valid values: `Insert`, `Update`, and `Delete`.
	//
	// - `dmlAction`: The DML handling policy. Valid values: `Normal` (process normally), `Ignore` (ignore), `Filter` (process conditionally, used when `dmlType` is `Update` or `Delete`), and `LogicalDelete` (logically delete).
	//
	// - `filterCondition`: The DML filter condition. This parameter is used when `dmlAction` is `Filter`.
	//
	// 5. Define incremental condition rule (`DefineIncrementalCondition`)
	//
	// - Example: `{"where":"id > 0"}`
	//
	// - Specifies the filter condition for incremental synchronization.
	//
	// 6. Define cycle schedule settings rule (`DefineCycleScheduleSettings`)
	//
	// - Example: `{"cronExpress":" 	- 	- 	- 	- 	- *", "cycleType":"1"}`
	//
	// - Specifies the scheduling parameters for a periodic job.
	//
	// 7. Define partition key rule (`DefinePartitionKey`)
	//
	// - Example: `{"columns":["id"]}`
	//
	// - Specifies the partition key.
	//
	// example:
	//
	// {"expression":"${srcDatasoureName}_${srcDatabaseName}"}
	RuleExpression *string `json:"RuleExpression,omitempty" xml:"RuleExpression,omitempty"`
	// The name of the rule. The rule name must be unique for a specific action type (`RuleActionType`) and target type (`RuleTargetType`).
	//
	// example:
	//
	// rename_rule_1
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The target object type of the action. Valid values:
	//
	// - `Table`
	//
	// - `Schema`
	//
	// - `Database`
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
