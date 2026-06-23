// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDIJobShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateDIJobShrinkRequest
	GetDescription() *string
	SetDestinationDataSourceSettingsShrink(v string) *CreateDIJobShrinkRequest
	GetDestinationDataSourceSettingsShrink() *string
	SetDestinationDataSourceType(v string) *CreateDIJobShrinkRequest
	GetDestinationDataSourceType() *string
	SetFileSpec(v string) *CreateDIJobShrinkRequest
	GetFileSpec() *string
	SetJobName(v string) *CreateDIJobShrinkRequest
	GetJobName() *string
	SetJobSettingsShrink(v string) *CreateDIJobShrinkRequest
	GetJobSettingsShrink() *string
	SetJobType(v string) *CreateDIJobShrinkRequest
	GetJobType() *string
	SetMigrationType(v string) *CreateDIJobShrinkRequest
	GetMigrationType() *string
	SetName(v string) *CreateDIJobShrinkRequest
	GetName() *string
	SetOwner(v string) *CreateDIJobShrinkRequest
	GetOwner() *string
	SetProjectId(v int64) *CreateDIJobShrinkRequest
	GetProjectId() *int64
	SetResourceSettingsShrink(v string) *CreateDIJobShrinkRequest
	GetResourceSettingsShrink() *string
	SetSourceDataSourceSettingsShrink(v string) *CreateDIJobShrinkRequest
	GetSourceDataSourceSettingsShrink() *string
	SetSourceDataSourceType(v string) *CreateDIJobShrinkRequest
	GetSourceDataSourceType() *string
	SetTableMappingsShrink(v string) *CreateDIJobShrinkRequest
	GetTableMappingsShrink() *string
	SetTransformationRulesShrink(v string) *CreateDIJobShrinkRequest
	GetTransformationRulesShrink() *string
}

type CreateDIJobShrinkRequest struct {
	// The description of the job.
	//
	// example:
	//
	// DI Job Demo
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Settings for the destination data sources.
	DestinationDataSourceSettingsShrink *string `json:"DestinationDataSourceSettings,omitempty" xml:"DestinationDataSourceSettings,omitempty"`
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
	JobSettingsShrink *string `json:"JobSettings,omitempty" xml:"JobSettings,omitempty"`
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
	ResourceSettingsShrink *string `json:"ResourceSettings,omitempty" xml:"ResourceSettings,omitempty"`
	// Settings for the source data sources.
	SourceDataSourceSettingsShrink *string `json:"SourceDataSourceSettings,omitempty" xml:"SourceDataSourceSettings,omitempty"`
	// The type of the source data source. Valid values: `PolarDB`, `MySQL`, `Kafka`, `LogHub`, `Hologres`, `Oracle`, `OceanBase`, `MongoDB`, `Redshift`, `Hive`, `SQL Server`, `Doris`, and `ClickHouse`.
	//
	// example:
	//
	// MySQL
	SourceDataSourceType *string `json:"SourceDataSourceType,omitempty" xml:"SourceDataSourceType,omitempty"`
	// Transformation mappings for the objects to be synchronized. Each mapping defines selection rules for a group of source objects and the transformation rules to apply to them.
	//
	// > [ { "SourceObjectSelectionRules":[ { "ObjectType":"Database", "Action":"Include", "ExpressionType":"Exact", "Expression":"biz_db" }, { "ObjectType":"Schema", "Action":"Include", "ExpressionType":"Exact", "Expression":"s1" }, { "ObjectType":"Table", "Action":"Include", "ExpressionType":"Exact", "Expression":"table1" } ], "TransformationRuleNames":[ { "RuleName":"my_database_rename_rule", "RuleActionType":"Rename", "RuleTargetType":"Schema" } ] } ]
	TableMappingsShrink *string `json:"TableMappings,omitempty" xml:"TableMappings,omitempty"`
	// A list of transformation rules for the objects to be synchronized.
	//
	// > [ { "RuleName":"my_database_rename_rule", "RuleActionType":"Rename", "RuleTargetType":"Schema", "RuleExpression":"{\\\\"expression\\\\":\\\\"${srcDatasoureName}_${srcDatabaseName}\\\\"}" } ]
	TransformationRulesShrink *string `json:"TransformationRules,omitempty" xml:"TransformationRules,omitempty"`
}

func (s CreateDIJobShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDIJobShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateDIJobShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateDIJobShrinkRequest) GetDestinationDataSourceSettingsShrink() *string {
	return s.DestinationDataSourceSettingsShrink
}

func (s *CreateDIJobShrinkRequest) GetDestinationDataSourceType() *string {
	return s.DestinationDataSourceType
}

func (s *CreateDIJobShrinkRequest) GetFileSpec() *string {
	return s.FileSpec
}

func (s *CreateDIJobShrinkRequest) GetJobName() *string {
	return s.JobName
}

func (s *CreateDIJobShrinkRequest) GetJobSettingsShrink() *string {
	return s.JobSettingsShrink
}

func (s *CreateDIJobShrinkRequest) GetJobType() *string {
	return s.JobType
}

func (s *CreateDIJobShrinkRequest) GetMigrationType() *string {
	return s.MigrationType
}

func (s *CreateDIJobShrinkRequest) GetName() *string {
	return s.Name
}

func (s *CreateDIJobShrinkRequest) GetOwner() *string {
	return s.Owner
}

func (s *CreateDIJobShrinkRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *CreateDIJobShrinkRequest) GetResourceSettingsShrink() *string {
	return s.ResourceSettingsShrink
}

func (s *CreateDIJobShrinkRequest) GetSourceDataSourceSettingsShrink() *string {
	return s.SourceDataSourceSettingsShrink
}

func (s *CreateDIJobShrinkRequest) GetSourceDataSourceType() *string {
	return s.SourceDataSourceType
}

func (s *CreateDIJobShrinkRequest) GetTableMappingsShrink() *string {
	return s.TableMappingsShrink
}

func (s *CreateDIJobShrinkRequest) GetTransformationRulesShrink() *string {
	return s.TransformationRulesShrink
}

func (s *CreateDIJobShrinkRequest) SetDescription(v string) *CreateDIJobShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreateDIJobShrinkRequest) SetDestinationDataSourceSettingsShrink(v string) *CreateDIJobShrinkRequest {
	s.DestinationDataSourceSettingsShrink = &v
	return s
}

func (s *CreateDIJobShrinkRequest) SetDestinationDataSourceType(v string) *CreateDIJobShrinkRequest {
	s.DestinationDataSourceType = &v
	return s
}

func (s *CreateDIJobShrinkRequest) SetFileSpec(v string) *CreateDIJobShrinkRequest {
	s.FileSpec = &v
	return s
}

func (s *CreateDIJobShrinkRequest) SetJobName(v string) *CreateDIJobShrinkRequest {
	s.JobName = &v
	return s
}

func (s *CreateDIJobShrinkRequest) SetJobSettingsShrink(v string) *CreateDIJobShrinkRequest {
	s.JobSettingsShrink = &v
	return s
}

func (s *CreateDIJobShrinkRequest) SetJobType(v string) *CreateDIJobShrinkRequest {
	s.JobType = &v
	return s
}

func (s *CreateDIJobShrinkRequest) SetMigrationType(v string) *CreateDIJobShrinkRequest {
	s.MigrationType = &v
	return s
}

func (s *CreateDIJobShrinkRequest) SetName(v string) *CreateDIJobShrinkRequest {
	s.Name = &v
	return s
}

func (s *CreateDIJobShrinkRequest) SetOwner(v string) *CreateDIJobShrinkRequest {
	s.Owner = &v
	return s
}

func (s *CreateDIJobShrinkRequest) SetProjectId(v int64) *CreateDIJobShrinkRequest {
	s.ProjectId = &v
	return s
}

func (s *CreateDIJobShrinkRequest) SetResourceSettingsShrink(v string) *CreateDIJobShrinkRequest {
	s.ResourceSettingsShrink = &v
	return s
}

func (s *CreateDIJobShrinkRequest) SetSourceDataSourceSettingsShrink(v string) *CreateDIJobShrinkRequest {
	s.SourceDataSourceSettingsShrink = &v
	return s
}

func (s *CreateDIJobShrinkRequest) SetSourceDataSourceType(v string) *CreateDIJobShrinkRequest {
	s.SourceDataSourceType = &v
	return s
}

func (s *CreateDIJobShrinkRequest) SetTableMappingsShrink(v string) *CreateDIJobShrinkRequest {
	s.TableMappingsShrink = &v
	return s
}

func (s *CreateDIJobShrinkRequest) SetTransformationRulesShrink(v string) *CreateDIJobShrinkRequest {
	s.TransformationRulesShrink = &v
	return s
}

func (s *CreateDIJobShrinkRequest) Validate() error {
	return dara.Validate(s)
}
