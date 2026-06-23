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
	// The description of the job.
	//
	// example:
	//
	// DI Job Demo
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Settings for the destination data sources.
	DestinationDataSourceSettings []*CreateDIJobRequestDestinationDataSourceSettings `json:"DestinationDataSourceSettings,omitempty" xml:"DestinationDataSourceSettings,omitempty" type:"Repeated"`
	// The type of the destination data source. Valid values: `Hologres`, `OSS-HDFS`, `OSS`, `MaxCompute`, `LogHub`, `StarRocks`, `DataHub`, `AnalyticDB for MySQL`, `Kafka`, and `Hive`.
	//
	// example:
	//
	// Hologres
	DestinationDataSourceType *string `json:"DestinationDataSourceType,omitempty" xml:"DestinationDataSourceType,omitempty"`
	// The code for a job created in script mode.
	//
	// example:
	//
	// {
	//
	//     "resourceSettings": {
	//
	//         "realtimeResourceSettings": {
	//
	//             "requestedCu": 2,
	//
	//             "resourceGroupIdentifier": "Serverless_res_group_123_456"
	//
	//         },
	//
	//         "offlineResourceSettings": {
	//
	//             "requestedCu": 2,
	//
	//             "resourceGroupIdentifier": "Serverless_res_group_123_456"
	//
	//         }
	//
	//     },
	//
	//     "tableMappings": [
	//
	//         {
	//
	//             "sourceObjectSelectionRules": [
	//
	//                 {
	//
	//                     "expression": "autotest_hologres",
	//
	//                     "action": "Include",
	//
	//                     "expressionType": "Exact",
	//
	//                     "objectType": "Datasource"
	//
	//                 },
	//
	//                 {
	//
	//                     "expression": "auto_holo_2661647",
	//
	//                     "action": "Include",
	//
	//                     "expressionType": "Exact",
	//
	//                     "objectType": "Table"
	//
	//                 },
	//
	//                 {
	//
	//                     "expression": "public",
	//
	//                     "action": "Include",
	//
	//                     "expressionType": "Exact",
	//
	//                     "objectType": "Schema"
	//
	//                 }
	//
	//             ],
	//
	//             "transformationRules": [
	//
	//                 {
	//
	//                     "ruleTargetType": "Table",
	//
	//                     "ruleActionType": "SourceSchema",
	//
	//                     "ruleName": "SourceSchema_Table_BStf8aXPSCJjOWGe"
	//
	//                 },
	//
	//                 {
	//
	//                     "ruleTargetType": "Schema",
	//
	//                     "ruleActionType": "Rename",
	//
	//                     "ruleName": "Rename_Schema_3qWNOIsljtInvKJy"
	//
	//                 },
	//
	//                 {
	//
	//                     "ruleTargetType": "Table",
	//
	//                     "ruleActionType": "Rename",
	//
	//                     "ruleName": "Rename_Table_o3PVQq1aIKDGoVVW"
	//
	//                 },
	//
	//                 {
	//
	//                     "ruleTargetType": "Table",
	//
	//                     "ruleActionType": "DefineDstTableSettings",
	//
	//                     "ruleName": "DefineDstTableSettings_Table_BhJltOmOCIc81fzi"
	//
	//                 },
	//
	//                 {
	//
	//                     "ruleTargetType": "Table",
	//
	//                     "ruleActionType": "ColumnMapping",
	//
	//                     "ruleName": "ColumnMapping_Table_nP4hJPX1wh2W3fpo"
	//
	//                 }
	//
	//             ]
	//
	//         }
	//
	//     ],
	//
	//     "sourceDataSourceSettings": [
	//
	//         {
	//
	//             "dataSourceProperties": {
	//
	//                 "timeZone": "Asia/Shanghai"
	//
	//             },
	//
	//             "dataSourceName": "autotest_hologres"
	//
	//         }
	//
	//     ],
	//
	//     "jobSettings": {
	//
	//         "runtimeSettings": [
	//
	//         ],
	//
	//         "ddlHandlingSettings": [
	//
	//         ],
	//
	//         "columnDataTypeSettings": [
	//
	//         ],
	//
	//         "cycleScheduleSettings": {
	//
	//         },
	//
	//         "channelSettings": {
	//
	//             "destinationChannelSettings": {
	//
	//                 "conflictMode": "replace",
	//
	//                 "dynamicColumnAction": "replay",
	//
	//                 "writeMode": "replay"
	//
	//             },
	//
	//             "sourceChannelSettings": {
	//
	//             }
	//
	//         }
	//
	//     },
	//
	//     "destinationDataSourceType": "Hologres",
	//
	//     "transformationRules": [
	//
	//         {
	//
	//             "ruleTargetType": "Table",
	//
	//             "ruleName": "SourceSchema_Table_BStf8aXPSCJjOWGe",
	//
	//             "ruleActionType": "SourceSchema",
	//
	//             "ruleExpression": {
	//
	//                 "columns": [
	//
	//                     {
	//
	//                         "name": "id",
	//
	//                         "category": "normal",
	//
	//                         "type": "BIGINT"
	//
	//                     },
	//
	//                     {
	//
	//                         "name": "decimal",
	//
	//                         "category": "normal",
	//
	//                         "type": "DECIMAL"
	//
	//                     }
	//
	//                 ]
	//
	//             }
	//
	//         },
	//
	//         {
	//
	//             "ruleTargetType": "Schema",
	//
	//             "ruleName": "Rename_Schema_3qWNOIsljtInvKJy",
	//
	//             "ruleActionType": "Rename",
	//
	//             "ruleExpression": {
	//
	//                 "expression": "public"
	//
	//             }
	//
	//         },
	//
	//         {
	//
	//             "ruleTargetType": "Table",
	//
	//             "ruleName": "Rename_Table_o3PVQq1aIKDGoVVW",
	//
	//             "ruleActionType": "Rename",
	//
	//             "ruleExpression": {
	//
	//                 "expression": "auto_holo_2661647_dst"
	//
	//             }
	//
	//         },
	//
	//         {
	//
	//             "ruleTargetType": "Table",
	//
	//             "ruleName": "DefineDstTableSettings_Table_BhJltOmOCIc81fzi",
	//
	//             "ruleActionType": "DefineDstTableSettings",
	//
	//             "ruleExpression": {
	//
	//                 "ddlString": "BEGIN;
	//
	// CREATE TABLE IF NOT EXISTS public.auto_holo_2661647_dst (
	//
	//    id          BIGINT PRIMARY KEY,
	//
	//    "decimal"   DECIMAL(38,18)
	//
	// );
	//
	// CALL SET_TABLE_PROPERTY(\\"public.auto_holo_2661647_dst\\", \\"time_to_live_in_seconds\\", \\"3153600000\\");
	//
	// CALL SET_TABLE_PROPERTY(\\"public.auto_holo_2661647_dst\\", \\"orientation\\", \\"column\\");
	//
	// CALL SET_TABLE_PROPERTY(\\"public.auto_holo_2661647_dst\\", \\"binlog.level\\", \\"replica\\");
	//
	// CALL SET_TABLE_PROPERTY(\\"public.auto_holo_2661647_dst\\", \\"binlog.ttl\\", \\"2592000\\");
	//
	// CALL SET_TABLE_PROPERTY(\\"public.auto_holo_2661647_dst\\", \\"bitmap_columns\\", \\""text","char","varchar"\\");
	//
	// CALL SET_TABLE_PROPERTY(\\"public.auto_holo_2661647_dst\\", \\"dictionary_encoding_columns\\", \\""text":auto,"bytea":auto,"char":auto,"varchar":auto\\");
	//
	// CALL SET_TABLE_PROPERTY(\\"public.auto_holo_2661647_dst\\", \\"distribution_key\\", \\""id"\\");
	//
	// COMMIT;
	//
	// ",
	//
	//                 "ddlType": "STRUCT"
	//
	//             }
	//
	//         },
	//
	//         {
	//
	//             "ruleTargetType": "Table",
	//
	//             "ruleName": "ColumnMapping_Table_nP4hJPX1wh2W3fpo",
	//
	//             "ruleActionType": "ColumnMapping",
	//
	//             "ruleExpression": {
	//
	//                 "columnMapping": [
	//
	//                     {
	//
	//                         "sourceColName": "id",
	//
	//                         "dstColName": "id"
	//
	//                     },
	//
	//                     {
	//
	//                         "sourceColName": "decimal",
	//
	//                         "dstColName": "decimal"
	//
	//                     }
	//
	//                 ]
	//
	//             }
	//
	//         }
	//
	//     ],
	//
	//     "migrationType": "FullAndRealtimeIncremental",
	//
	//     "destinationDataSourceSettings": [
	//
	//         {
	//
	//             "dataSourceProperties": {
	//
	//             },
	//
	//             "dataSourceName": "autotest_hologres"
	//
	//         }
	//
	//     ],
	//
	//     "sourceDataSourceType": "Hologres"
	//
	// }
	FileSpec *string `json:"FileSpec,omitempty" xml:"FileSpec,omitempty"`
	// Deprecated
	//
	// This parameter is deprecated. Use the `Name` parameter instead.
	//
	// example:
	//
	// mysql_to_holo_sync_8772
	JobName *string `json:"JobName,omitempty" xml:"JobName,omitempty"`
	// The settings for the synchronization job, including DDL processing policies, data type mappings between source and destination columns, and runtime parameters.
	JobSettings *CreateDIJobRequestJobSettings `json:"JobSettings,omitempty" xml:"JobSettings,omitempty" type:"Struct"`
	// The job type. Valid values:
	//
	// - `DatabaseRealtimeMigration`: Synchronizes multiple tables from multiple source databases in real time (stream synchronization). This type supports full, incremental, or both full and incremental synchronization.
	//
	// - `DatabaseOfflineMigration`: Synchronizes multiple tables from multiple source databases in batches. This type supports full, incremental, or both full and incremental synchronization.
	//
	// - `SingleTableRealtimeMigration`: Synchronizes a single source table in real time (stream synchronization).
	//
	// example:
	//
	// DatabaseRealtimeMigration
	JobType *string `json:"JobType,omitempty" xml:"JobType,omitempty"`
	// The synchronization type. Valid values:
	//
	// - `FullAndRealtimeIncremental`: Full and real-time incremental synchronization for an entire database.
	//
	// - `RealtimeIncremental`: Real-time incremental synchronization for a single table.
	//
	// - `Full`: Full batch synchronization for an entire database.
	//
	// - `OfflineIncremental`: Incremental synchronization in batch mode.
	//
	// - `FullAndOfflineIncremental`: Full and incremental batch synchronization for an entire database.
	//
	// example:
	//
	// FullAndRealtimeIncremental
	MigrationType *string `json:"MigrationType,omitempty" xml:"MigrationType,omitempty"`
	// The name of the job.
	//
	// example:
	//
	// mysql_to_holo_sync_8772
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The job owner.
	//
	// example:
	//
	// 3726346
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The ID of the DataWorks workspace for this API call. To obtain the workspace ID, log on to the [DataWorks console](https://workbench.data.aliyun.com/console) and go to the Workspace Management page.
	//
	// example:
	//
	// 10000
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The resource settings.
	ResourceSettings *CreateDIJobRequestResourceSettings `json:"ResourceSettings,omitempty" xml:"ResourceSettings,omitempty" type:"Struct"`
	// Settings for the source data sources.
	SourceDataSourceSettings []*CreateDIJobRequestSourceDataSourceSettings `json:"SourceDataSourceSettings,omitempty" xml:"SourceDataSourceSettings,omitempty" type:"Repeated"`
	// The type of the source data source. Valid values: `PolarDB`, `MySQL`, `Kafka`, `LogHub`, `Hologres`, `Oracle`, `OceanBase`, `MongoDB`, `Redshift`, `Hive`, `SQL Server`, `Doris`, and `ClickHouse`.
	//
	// example:
	//
	// MySQL
	SourceDataSourceType *string `json:"SourceDataSourceType,omitempty" xml:"SourceDataSourceType,omitempty"`
	// Transformation mappings for the objects to be synchronized. Each mapping defines selection rules for a group of source objects and the transformation rules to apply to them.
	//
	// > [ { "SourceObjectSelectionRules":[ { "ObjectType":"Database", "Action":"Include", "ExpressionType":"Exact", "Expression":"biz_db" }, { "ObjectType":"Schema", "Action":"Include", "ExpressionType":"Exact", "Expression":"s1" }, { "ObjectType":"Table", "Action":"Include", "ExpressionType":"Exact", "Expression":"table1" } ], "TransformationRuleNames":[ { "RuleName":"my_database_rename_rule", "RuleActionType":"Rename", "RuleTargetType":"Schema" } ] } ]
	TableMappings []*CreateDIJobRequestTableMappings `json:"TableMappings,omitempty" xml:"TableMappings,omitempty" type:"Repeated"`
	// A list of transformation rules for the objects to be synchronized.
	//
	// > [ { "RuleName":"my_database_rename_rule", "RuleActionType":"Rename", "RuleTargetType":"Schema", "RuleExpression":"{\\\\"expression\\\\":\\\\"${srcDatasoureName}_${srcDatabaseName}\\\\"}" } ]
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
	// The name of the data source.
	//
	// example:
	//
	// holo_datasource_1
	DataSourceName *string `json:"DataSourceName,omitempty" xml:"DataSourceName,omitempty"`
	// The properties of the data source.
	DataSourceProperties *CreateDIJobRequestDestinationDataSourceSettingsDataSourceProperties `json:"DataSourceProperties,omitempty" xml:"DataSourceProperties,omitempty" type:"Struct"`
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

func (s *CreateDIJobRequestDestinationDataSourceSettings) GetDataSourceProperties() *CreateDIJobRequestDestinationDataSourceSettingsDataSourceProperties {
	return s.DataSourceProperties
}

func (s *CreateDIJobRequestDestinationDataSourceSettings) SetDataSourceName(v string) *CreateDIJobRequestDestinationDataSourceSettings {
	s.DataSourceName = &v
	return s
}

func (s *CreateDIJobRequestDestinationDataSourceSettings) SetDataSourceProperties(v *CreateDIJobRequestDestinationDataSourceSettingsDataSourceProperties) *CreateDIJobRequestDestinationDataSourceSettings {
	s.DataSourceProperties = v
	return s
}

func (s *CreateDIJobRequestDestinationDataSourceSettings) Validate() error {
	if s.DataSourceProperties != nil {
		if err := s.DataSourceProperties.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateDIJobRequestDestinationDataSourceSettingsDataSourceProperties struct {
	// Custom connection settings for the data source, such as instance ID, access credentials, and instance region. You must specify this parameter or `DataSourceName`.
	//
	// This parameter applies only when the data source is configured in instance mode (`ConnectionPropertiesMode`). The property format varies by data source. For more information, see [ConnectionProperties for data sources](https://help.aliyun.com/document_detail/2852465.html).
	//
	// example:
	//
	// { "instanceId": "rm-2ze09gn3x6xxx", "password": "xxxx", "database": "agent", "username": "zmtest" "regionId": "cn-beijing" }
	ConnectionProperties *string `json:"ConnectionProperties,omitempty" xml:"ConnectionProperties,omitempty"`
}

func (s CreateDIJobRequestDestinationDataSourceSettingsDataSourceProperties) String() string {
	return dara.Prettify(s)
}

func (s CreateDIJobRequestDestinationDataSourceSettingsDataSourceProperties) GoString() string {
	return s.String()
}

func (s *CreateDIJobRequestDestinationDataSourceSettingsDataSourceProperties) GetConnectionProperties() *string {
	return s.ConnectionProperties
}

func (s *CreateDIJobRequestDestinationDataSourceSettingsDataSourceProperties) SetConnectionProperties(v string) *CreateDIJobRequestDestinationDataSourceSettingsDataSourceProperties {
	s.ConnectionProperties = &v
	return s
}

func (s *CreateDIJobRequestDestinationDataSourceSettingsDataSourceProperties) Validate() error {
	return dara.Validate(s)
}

type CreateDIJobRequestJobSettings struct {
	// Settings for data synchronization channels. You can configure special settings for specific channels. The following channels are supported: synchronization from Hologres to Hologres (Holo2Holo) and from Hologres to Kafka (Holo2Kafka).
	//
	// 1. Holo2Kafka
	//
	// - Example: `{"destinationChannelSettings":{"kafkaClientProperties":[{"key":"linger.ms","value":"100"}],"keyColumns":["col3"],"writeMode":"canal"}}`
	//
	//   `kafkaClientProperties`: Parameters for the Kafka producer.
	//
	// - `keyColumns`: The columns whose values are used as keys for data written to Kafka.
	//
	// - `writeMode`: The data format for writing to Kafka. Valid values: `json` and `canal`.
	//
	// 2. Holo2Holo
	//
	// - Example: `{"destinationChannelSettings":{"conflictMode":"replace","dynamicColumnAction":"replay","writeMode":"replay"}}`
	//
	// - `conflictMode`: The conflict handling policy for writing data to Hologres. Valid values: `replace` (overwrite) and `ignore`.
	//
	// - `writeMode`: The method for writing data to Hologres. Valid values: `replay` and `insert`.
	//
	// - `dynamicColumnAction`: The method for handling dynamic columns when writing data to Hologres. Valid values: `replay`, `insert`, and `ignore`.
	//
	// example:
	//
	// {
	//
	//       "structInfo": "MANAGED",
	//
	//       "storageType": "TEXTFILE",
	//
	//       "writeMode": "APPEND",
	//
	//       "partitionColumns": [
	//
	//             {
	//
	//                   "columnName": "pt",
	//
	//                   "columnType": "STRING",
	//
	//                   "comment": ""
	//
	//             }
	//
	//       ],
	//
	//       "fieldDelimiter": ""
	//
	// }
	ChannelSettings *string `json:"ChannelSettings,omitempty" xml:"ChannelSettings,omitempty"`
	// Column data type mappings.
	//
	// > "ColumnDataTypeSettings":[ { "SourceDataType":"Bigint", "DestinationDataType":"Text" } ]
	ColumnDataTypeSettings []*CreateDIJobRequestJobSettingsColumnDataTypeSettings `json:"ColumnDataTypeSettings,omitempty" xml:"ColumnDataTypeSettings,omitempty" type:"Repeated"`
	// The periodic scheduling settings.
	CycleScheduleSettings *CreateDIJobRequestJobSettingsCycleScheduleSettings `json:"CycleScheduleSettings,omitempty" xml:"CycleScheduleSettings,omitempty" type:"Struct"`
	// DDL handling settings.
	//
	// > "DDLHandlingSettings":[ { "Type":"Insert", "Action":"Normal" } ]
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
	// The destination data type. For example: `bigint`, `boolean`, `string`, `text`, `datetime`, `timestamp`, `decimal`, or `binary`. Available data types vary by data source.
	//
	// example:
	//
	// text
	DestinationDataType *string `json:"DestinationDataType,omitempty" xml:"DestinationDataType,omitempty"`
	// The source data type. For example: `bigint`, `boolean`, `string`, `text`, `datetime`, `timestamp`, `decimal`, or `binary`. Available data types vary by data source.
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
	// The synchronization type for periodic scheduling. Valid values:
	//
	// - `Full`: Full synchronization.
	//
	// - `OfflineIncremental`: Incremental synchronization in batch mode.
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
	// The handling action. Valid values:
	//
	// - `Ignore`: Ignores the DDL message.
	//
	// - `Critical`: Reports an error.
	//
	// - `Normal`: Processes the DDL message normally.
	//
	// example:
	//
	// Critical
	Action *string `json:"Action,omitempty" xml:"Action,omitempty"`
	// The DDL type. Valid values:
	//
	// - `RenameColumn`: Renames a column.
	//
	// - `ModifyColumn`: Modifies a column.
	//
	// - `CreateTable`: Creates a table.
	//
	// - `TruncateTable`: Truncates a table.
	//
	// - `DropTable`: Drops a table.
	//
	// - `DropColumn`: Drops a column.
	//
	// - `AddColumn`: Adds a column.
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
	// The name of the setting. Valid values:
	//
	// - `src.offline.datasource.max.connection`: The maximum number of connections to the source of a batch synchronization job.
	//
	// - `dst.offline.truncate`: Specifies whether to truncate the destination table before a batch job starts.
	//
	// - `runtime.offline.speed.limit.enable`: Specifies whether to enable throttling for a batch synchronization job.
	//
	// - `runtime.offline.concurrent`: The concurrency level of a batch synchronization job.
	//
	// - `runtime.enable.auto.create.schema`: Specifies whether to automatically create a destination schema.
	//
	// - `runtime.realtime.concurrent`: The concurrency level of a real-time synchronization job.
	//
	// - `runtime.realtime.failover.minute.dataxcdc`: The wait time in minutes for a failover restart.
	//
	// - `runtime.realtime.failover.times.dataxcdc`: The number of failover restart attempts.
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
	// Resource settings for batch synchronization.
	OfflineResourceSettings *CreateDIJobRequestResourceSettingsOfflineResourceSettings `json:"OfflineResourceSettings,omitempty" xml:"OfflineResourceSettings,omitempty" type:"Struct"`
	// The resources for real-time synchronization.
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
	// The number of CUs for the resource group for data integration that is used for batch synchronization.
	//
	// example:
	//
	// 2
	RequestedCu *float64 `json:"RequestedCu,omitempty" xml:"RequestedCu,omitempty"`
	// The identifier of the resource group for data integration used for batch synchronization.
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
	// The number of CUs for the resource group for data integration that is used for real-time synchronization.
	//
	// example:
	//
	// 2
	RequestedCu *float64 `json:"RequestedCu,omitempty" xml:"RequestedCu,omitempty"`
	// The identifier of the resource group for data integration used for real-time synchronization.
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
	// The number of CUs for the scheduling resource group that is used for batch synchronization jobs.
	//
	// example:
	//
	// 2
	RequestedCu *float64 `json:"RequestedCu,omitempty" xml:"RequestedCu,omitempty"`
	// The identifier of the scheduling resource group used for batch synchronization jobs.
	//
	// example:
	//
	// S_res_group_222_333
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
	// The name of the data source.
	//
	// example:
	//
	// mysql_datasource_1
	DataSourceName *string `json:"DataSourceName,omitempty" xml:"DataSourceName,omitempty"`
	// The properties of the data source.
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
	// Custom connection settings for the data source, such as instance ID, access credentials, and instance region. You must specify this parameter or `DataSourceName`.
	//
	// This parameter applies only when the data source is configured in instance mode (`ConnectionPropertiesMode`). The property format varies by data source. For more information, see [ConnectionProperties for data sources](https://help.aliyun.com/document_detail/2852465.html).
	//
	// example:
	//
	// { "instanceId": "rm-2ze09gn3x6xxx", "password": "xxxx", "database": "agent", "username": "zmtest" "regionId": "cn-beijing" }
	ConnectionProperties *string `json:"ConnectionProperties,omitempty" xml:"ConnectionProperties,omitempty"`
	// The database encoding format.
	//
	// example:
	//
	// UTF-8
	Encoding *string `json:"Encoding,omitempty" xml:"Encoding,omitempty"`
	// The time zone.
	//
	// example:
	//
	// Asia/Shanghai
	Timezone *string `json:"Timezone,omitempty" xml:"Timezone,omitempty"`
}

func (s CreateDIJobRequestSourceDataSourceSettingsDataSourceProperties) String() string {
	return dara.Prettify(s)
}

func (s CreateDIJobRequestSourceDataSourceSettingsDataSourceProperties) GoString() string {
	return s.String()
}

func (s *CreateDIJobRequestSourceDataSourceSettingsDataSourceProperties) GetConnectionProperties() *string {
	return s.ConnectionProperties
}

func (s *CreateDIJobRequestSourceDataSourceSettingsDataSourceProperties) GetEncoding() *string {
	return s.Encoding
}

func (s *CreateDIJobRequestSourceDataSourceSettingsDataSourceProperties) GetTimezone() *string {
	return s.Timezone
}

func (s *CreateDIJobRequestSourceDataSourceSettingsDataSourceProperties) SetConnectionProperties(v string) *CreateDIJobRequestSourceDataSourceSettingsDataSourceProperties {
	s.ConnectionProperties = &v
	return s
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
	// Each rule can select a set of source objects to synchronize. Multiple rules combine to select one table.
	SourceObjectSelectionRules []*CreateDIJobRequestTableMappingsSourceObjectSelectionRules `json:"SourceObjectSelectionRules,omitempty" xml:"SourceObjectSelectionRules,omitempty" type:"Repeated"`
	// The names of the transformation rules to apply to the selected objects.
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
	// - `DefinePrimaryKey`: Defines a primary key.
	//
	// - `Rename`: Renames an object.
	//
	// - `AddColumn`: Adds a column.
	//
	// - `HandleDml`: Handles DML operations.
	//
	// - `DefineIncrementalCondition`: Defines an incremental condition.
	//
	// - `DefineCycleScheduleSettings`: Defines periodic scheduling settings.
	//
	// - `DefinePartitionKey`: Defines a partition key.
	//
	// example:
	//
	// Rename
	RuleActionType *string `json:"RuleActionType,omitempty" xml:"RuleActionType,omitempty"`
	// The name of the transformation rule. The rule name must be unique for a specific action type and target object type. Maximum length: 50 characters.
	//
	// example:
	//
	// rename_rule_1
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The type of the object to which the action applies. Valid values:
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
	// - `DefinePrimaryKey`: Defines a primary key.
	//
	// - `Rename`: Renames an object.
	//
	// - `AddColumn`: Adds a column.
	//
	// - `HandleDml`: Handles DML operations.
	//
	// - `DefineIncrementalCondition`: Defines an incremental condition.
	//
	// - `DefineCycleScheduleSettings`: Defines periodic scheduling settings.
	//
	// - `DefinePartitionKey`: Defines a partition key.
	//
	// example:
	//
	// Rename
	RuleActionType *string `json:"RuleActionType,omitempty" xml:"RuleActionType,omitempty"`
	// The rule expression, specified as a JSON string.
	//
	// 1. Renaming rule (`Rename`)
	//
	// - Example: `{"expression":"${srcDatasourceName}_${srcDatabaseName}_0922" }`
	//
	// - `expression`: The renaming expression. You can use the following variables: `${srcDatasourceName}` (name of the source data source), `${srcDatabaseName}` (name of the source database), and `${srcTableName}` (name of the source table).
	//
	// 2. Rule for adding a column (`AddColumn`)
	//
	// - Example: `{"columns":[{"columnName":"my_add_column","columnValueType":"Constant","columnValue":"123"}]}`
	//
	// - If you do not specify this rule, no columns are added.
	//
	// - `columnName`: The name of the column to add.
	//
	// - `columnValueType`: The value type of the added column. Valid values: `Constant` and `Variable`.
	//
	// - `columnValue`: The value of the added column. If `columnValueType` is `Constant`, the value is a custom string constant. If `columnValueType` is `Variable`, the value is a built-in variable. Valid built-in variables include: `EXECUTE_TIME` (execution time, Long), `DB_NAME_SRC` (source database name, String), `DATASOURCE_NAME_SRC` (source data source name, String), `TABLE_NAME_SRC` (source table name, String), `DB_NAME_DEST` (destination database name, String), `DATASOURCE_NAME_DEST` (destination data source name, String), `TABLE_NAME_DEST` (destination table name, String), and `DB_NAME_SRC_TRANSED` (transformed database name, String).
	//
	// 3. Rule for defining the primary key columns of a destination table (`DefinePrimaryKey`)
	//
	// - Example: `{"columns":["ukcolumn1","ukcolumn2"]}`
	//
	// - If you do not specify this rule, the primary key columns of the source table are used by default.
	//
	// - If the destination is an existing table, Data Integration does not modify its schema. If a specified primary key column does not exist in the destination table, the job fails to start and an error is reported.
	//
	// - If the destination table is automatically created, Data Integration automatically creates its schema with the defined primary key columns. If a specified primary key column does not exist in the source table, the job fails to start and an error is reported.
	//
	// 4. DML handling rule (`HandleDml`)
	//
	// - Example: `{"dmlPolicies":[{"dmlType":"Delete","dmlAction":"Filter","filterCondition":"id > 1"}]}`
	//
	// - If you do not specify this rule, the default `dmlAction` is `Normal` for `Insert`, `Update`, and `Delete` operations.
	//
	// - `dmlType`: The DML operation type. Valid values: `Insert`, `Update`, and `Delete`.
	//
	// - `dmlAction`: The DML handling policy. Valid values: `Normal` (normal processing), `Ignore`, `Filter` (conditional processing, used when `dmlType` is `Update` or `Delete`), and `LogicalDelete` (logical deletion).
	//
	// - `filterCondition`: The DML filter condition, used when `dmlAction` is `Filter`.
	//
	// 5. Incremental condition (`DefineIncrementalCondition`)
	//
	// - Example: `{"where":"id > 0"}`
	//
	// - Specifies the filter condition for incremental synchronization.
	//
	// 6. Parameters for periodic scheduling (`DefineCycleScheduleSettings`)
	//
	// - Example: `{"cronExpress":" 	- 	- 	- 	- 	- *", "cycleType":"1"}`
	//
	// - Specifies the parameters for periodically scheduling a job.
	//
	// 7. Rule to define a partition key (`DefinePartitionKey`)
	//
	// - Example: `{"columns":["id"]}`
	//
	// - Specifies a partition key.
	//
	// example:
	//
	// {
	//
	//       "expression": "${srcDatasoureName}_${srcDatabaseName}"
	//
	// }
	RuleExpression *string `json:"RuleExpression,omitempty" xml:"RuleExpression,omitempty"`
	// The name of the rule. The rule name must be unique for a specific action type and target object type. Maximum length: 50 characters.
	//
	// example:
	//
	// rename_rule_1
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The type of the object to which the action applies. Valid values:
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
