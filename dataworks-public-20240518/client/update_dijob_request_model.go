// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDIJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDIJobId(v int64) *UpdateDIJobRequest
	GetDIJobId() *int64
	SetDescription(v string) *UpdateDIJobRequest
	GetDescription() *string
	SetFileSpec(v string) *UpdateDIJobRequest
	GetFileSpec() *string
	SetId(v int64) *UpdateDIJobRequest
	GetId() *int64
	SetJobSettings(v *UpdateDIJobRequestJobSettings) *UpdateDIJobRequest
	GetJobSettings() *UpdateDIJobRequestJobSettings
	SetOwner(v string) *UpdateDIJobRequest
	GetOwner() *string
	SetProjectId(v int64) *UpdateDIJobRequest
	GetProjectId() *int64
	SetResourceSettings(v *UpdateDIJobRequestResourceSettings) *UpdateDIJobRequest
	GetResourceSettings() *UpdateDIJobRequestResourceSettings
	SetTableMappings(v []*UpdateDIJobRequestTableMappings) *UpdateDIJobRequest
	GetTableMappings() []*UpdateDIJobRequestTableMappings
	SetTransformationRules(v []*UpdateDIJobRequestTransformationRules) *UpdateDIJobRequest
	GetTransformationRules() []*UpdateDIJobRequestTransformationRules
}

type UpdateDIJobRequest struct {
	// Deprecated
	//
	// This parameter is deprecated. Use the `Id` parameter instead.
	//
	// example:
	//
	// 11588
	DIJobId *int64 `json:"DIJobId,omitempty" xml:"DIJobId,omitempty"`
	// The description of the synchronization job.
	//
	// example:
	//
	// DI Job Demo
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The job configuration in script mode.
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
	// The ID of the synchronization job.
	//
	// example:
	//
	// 11588
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The settings for the synchronization job. This includes DDL handling settings, data type mappings for columns between the source and destination, and runtime parameters.
	JobSettings *UpdateDIJobRequestJobSettings `json:"JobSettings,omitempty" xml:"JobSettings,omitempty" type:"Struct"`
	// The owner of the synchronization job.
	//
	// example:
	//
	// 95279527
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The ID of the DataWorks workspace. You can call the [ListProjects](https://help.aliyun.com/document_detail/178393.html) operation to get the workspace ID.
	//
	// example:
	//
	// 10000
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The resource settings.
	ResourceSettings *UpdateDIJobRequestResourceSettings `json:"ResourceSettings,omitempty" xml:"ResourceSettings,omitempty" type:"Struct"`
	// A list of object transformation mappings. Each mapping specifies a set of selection rules for source objects and a list of transformation rules that apply to the selected objects.
	//
	// > [ { "SourceObjectSelectionRules":[ { "ObjectType":"Database", "Action":"Include", "ExpressionType":"Exact", "Expression":"biz_db" }, { "ObjectType":"Schema", "Action":"Include", "ExpressionType":"Exact", "Expression":"s1" }, { "ObjectType":"Table", "Action":"Include", "ExpressionType":"Exact", "Expression":"table1" } ], "TransformationRuleNames":[ { "RuleName":"my_database_rename_rule", "RuleActionType":"Rename", "RuleTargetType":"Schema" } ] } ]
	TableMappings []*UpdateDIJobRequestTableMappings `json:"TableMappings,omitempty" xml:"TableMappings,omitempty" type:"Repeated"`
	// A list of transformation rule definitions.
	//
	// > [ { "RuleName":"my_database_rename_rule", "RuleActionType":"Rename", "RuleTargetType":"Schema", "RuleExpression":"{"expression":"${srcDatasoureName}_${srcDatabaseName}"}" } ]
	TransformationRules []*UpdateDIJobRequestTransformationRules `json:"TransformationRules,omitempty" xml:"TransformationRules,omitempty" type:"Repeated"`
}

func (s UpdateDIJobRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateDIJobRequest) GoString() string {
	return s.String()
}

func (s *UpdateDIJobRequest) GetDIJobId() *int64 {
	return s.DIJobId
}

func (s *UpdateDIJobRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateDIJobRequest) GetFileSpec() *string {
	return s.FileSpec
}

func (s *UpdateDIJobRequest) GetId() *int64 {
	return s.Id
}

func (s *UpdateDIJobRequest) GetJobSettings() *UpdateDIJobRequestJobSettings {
	return s.JobSettings
}

func (s *UpdateDIJobRequest) GetOwner() *string {
	return s.Owner
}

func (s *UpdateDIJobRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *UpdateDIJobRequest) GetResourceSettings() *UpdateDIJobRequestResourceSettings {
	return s.ResourceSettings
}

func (s *UpdateDIJobRequest) GetTableMappings() []*UpdateDIJobRequestTableMappings {
	return s.TableMappings
}

func (s *UpdateDIJobRequest) GetTransformationRules() []*UpdateDIJobRequestTransformationRules {
	return s.TransformationRules
}

func (s *UpdateDIJobRequest) SetDIJobId(v int64) *UpdateDIJobRequest {
	s.DIJobId = &v
	return s
}

func (s *UpdateDIJobRequest) SetDescription(v string) *UpdateDIJobRequest {
	s.Description = &v
	return s
}

func (s *UpdateDIJobRequest) SetFileSpec(v string) *UpdateDIJobRequest {
	s.FileSpec = &v
	return s
}

func (s *UpdateDIJobRequest) SetId(v int64) *UpdateDIJobRequest {
	s.Id = &v
	return s
}

func (s *UpdateDIJobRequest) SetJobSettings(v *UpdateDIJobRequestJobSettings) *UpdateDIJobRequest {
	s.JobSettings = v
	return s
}

func (s *UpdateDIJobRequest) SetOwner(v string) *UpdateDIJobRequest {
	s.Owner = &v
	return s
}

func (s *UpdateDIJobRequest) SetProjectId(v int64) *UpdateDIJobRequest {
	s.ProjectId = &v
	return s
}

func (s *UpdateDIJobRequest) SetResourceSettings(v *UpdateDIJobRequestResourceSettings) *UpdateDIJobRequest {
	s.ResourceSettings = v
	return s
}

func (s *UpdateDIJobRequest) SetTableMappings(v []*UpdateDIJobRequestTableMappings) *UpdateDIJobRequest {
	s.TableMappings = v
	return s
}

func (s *UpdateDIJobRequest) SetTransformationRules(v []*UpdateDIJobRequestTransformationRules) *UpdateDIJobRequest {
	s.TransformationRules = v
	return s
}

func (s *UpdateDIJobRequest) Validate() error {
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

type UpdateDIJobRequestJobSettings struct {
	// The job settings for specific data synchronization channels. You can apply special configurations to certain channels. Currently, `Holo2Holo` (synchronization from Hologres to Hologres) and `Holo2Kafka` (synchronization from Hologres to Kafka) are supported.
	//
	// 1. `Holo2Kafka`
	//
	// - Example: `{"destinationChannelSettings":{"kafkaClientProperties":[{"key":"linger.ms","value":"100"}],"keyColumns":["col3"],"writeMode":"canal"}}`
	//
	//   `kafkaClientProperties`: The Kafka producer parameters used when writing to Kafka.
	//
	// - `keyColumns`: The columns whose values are written to the key of a Kafka message.
	//
	// - `writeMode`: The format for writing data to Kafka. Valid values: `json` and `canal`.
	//
	// 2. `Holo2Holo`
	//
	// - Example: `{"destinationChannelSettings":{"conflictMode":"replace","dynamicColumnAction":"replay","writeMode":"replay"}}`
	//
	// - `conflictMode`: The conflict handling policy for writing data to Hologres. Valid values: `replace` (overwrite) and `ignore` (ignore).
	//
	// - `writeMode`: The method for writing data to Hologres. Valid values: `replay` and `insert`.
	//
	// - `dynamicColumnAction`: The action for handling dynamic columns when writing data to Hologres. Valid values: `replay`, `insert`, and `ignore`.
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
	// An array of column data type mappings.
	//
	// > ["ColumnDataTypeSettings":[ { "SourceDataType":"Bigint", "DestinationDataType":"Text" } ]
	ColumnDataTypeSettings []*UpdateDIJobRequestJobSettingsColumnDataTypeSettings `json:"ColumnDataTypeSettings,omitempty" xml:"ColumnDataTypeSettings,omitempty" type:"Repeated"`
	// The settings for periodic scheduling.
	CycleScheduleSettings *UpdateDIJobRequestJobSettingsCycleScheduleSettings `json:"CycleScheduleSettings,omitempty" xml:"CycleScheduleSettings,omitempty" type:"Struct"`
	// An array of DDL handling settings.
	//
	// > ["DDLHandlingSettings":[ { "Type":"Insert", "Action":"Normal" } ]
	DdlHandlingSettings []*UpdateDIJobRequestJobSettingsDdlHandlingSettings `json:"DdlHandlingSettings,omitempty" xml:"DdlHandlingSettings,omitempty" type:"Repeated"`
	// The runtime settings.
	RuntimeSettings []*UpdateDIJobRequestJobSettingsRuntimeSettings `json:"RuntimeSettings,omitempty" xml:"RuntimeSettings,omitempty" type:"Repeated"`
}

func (s UpdateDIJobRequestJobSettings) String() string {
	return dara.Prettify(s)
}

func (s UpdateDIJobRequestJobSettings) GoString() string {
	return s.String()
}

func (s *UpdateDIJobRequestJobSettings) GetChannelSettings() *string {
	return s.ChannelSettings
}

func (s *UpdateDIJobRequestJobSettings) GetColumnDataTypeSettings() []*UpdateDIJobRequestJobSettingsColumnDataTypeSettings {
	return s.ColumnDataTypeSettings
}

func (s *UpdateDIJobRequestJobSettings) GetCycleScheduleSettings() *UpdateDIJobRequestJobSettingsCycleScheduleSettings {
	return s.CycleScheduleSettings
}

func (s *UpdateDIJobRequestJobSettings) GetDdlHandlingSettings() []*UpdateDIJobRequestJobSettingsDdlHandlingSettings {
	return s.DdlHandlingSettings
}

func (s *UpdateDIJobRequestJobSettings) GetRuntimeSettings() []*UpdateDIJobRequestJobSettingsRuntimeSettings {
	return s.RuntimeSettings
}

func (s *UpdateDIJobRequestJobSettings) SetChannelSettings(v string) *UpdateDIJobRequestJobSettings {
	s.ChannelSettings = &v
	return s
}

func (s *UpdateDIJobRequestJobSettings) SetColumnDataTypeSettings(v []*UpdateDIJobRequestJobSettingsColumnDataTypeSettings) *UpdateDIJobRequestJobSettings {
	s.ColumnDataTypeSettings = v
	return s
}

func (s *UpdateDIJobRequestJobSettings) SetCycleScheduleSettings(v *UpdateDIJobRequestJobSettingsCycleScheduleSettings) *UpdateDIJobRequestJobSettings {
	s.CycleScheduleSettings = v
	return s
}

func (s *UpdateDIJobRequestJobSettings) SetDdlHandlingSettings(v []*UpdateDIJobRequestJobSettingsDdlHandlingSettings) *UpdateDIJobRequestJobSettings {
	s.DdlHandlingSettings = v
	return s
}

func (s *UpdateDIJobRequestJobSettings) SetRuntimeSettings(v []*UpdateDIJobRequestJobSettingsRuntimeSettings) *UpdateDIJobRequestJobSettings {
	s.RuntimeSettings = v
	return s
}

func (s *UpdateDIJobRequestJobSettings) Validate() error {
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

type UpdateDIJobRequestJobSettingsColumnDataTypeSettings struct {
	// The destination data type. Examples: `bigint`, `boolean`, `string`, `text`, `datetime`, `timestamp`, `decimal`, and `binary`. The supported data types depend on the destination data source.
	//
	// example:
	//
	// text
	DestinationDataType *string `json:"DestinationDataType,omitempty" xml:"DestinationDataType,omitempty"`
	// The source data type. Examples: `bigint`, `boolean`, `string`, `text`, `datetime`, `timestamp`, `decimal`, and `binary`. The supported data types depend on the source data source.
	//
	// example:
	//
	// bigint
	SourceDataType *string `json:"SourceDataType,omitempty" xml:"SourceDataType,omitempty"`
}

func (s UpdateDIJobRequestJobSettingsColumnDataTypeSettings) String() string {
	return dara.Prettify(s)
}

func (s UpdateDIJobRequestJobSettingsColumnDataTypeSettings) GoString() string {
	return s.String()
}

func (s *UpdateDIJobRequestJobSettingsColumnDataTypeSettings) GetDestinationDataType() *string {
	return s.DestinationDataType
}

func (s *UpdateDIJobRequestJobSettingsColumnDataTypeSettings) GetSourceDataType() *string {
	return s.SourceDataType
}

func (s *UpdateDIJobRequestJobSettingsColumnDataTypeSettings) SetDestinationDataType(v string) *UpdateDIJobRequestJobSettingsColumnDataTypeSettings {
	s.DestinationDataType = &v
	return s
}

func (s *UpdateDIJobRequestJobSettingsColumnDataTypeSettings) SetSourceDataType(v string) *UpdateDIJobRequestJobSettingsColumnDataTypeSettings {
	s.SourceDataType = &v
	return s
}

func (s *UpdateDIJobRequestJobSettingsColumnDataTypeSettings) Validate() error {
	return dara.Validate(s)
}

type UpdateDIJobRequestJobSettingsCycleScheduleSettings struct {
	// The scheduling parameters.
	//
	// example:
	//
	// bizdate=$bizdate
	ScheduleParameters *string `json:"ScheduleParameters,omitempty" xml:"ScheduleParameters,omitempty"`
}

func (s UpdateDIJobRequestJobSettingsCycleScheduleSettings) String() string {
	return dara.Prettify(s)
}

func (s UpdateDIJobRequestJobSettingsCycleScheduleSettings) GoString() string {
	return s.String()
}

func (s *UpdateDIJobRequestJobSettingsCycleScheduleSettings) GetScheduleParameters() *string {
	return s.ScheduleParameters
}

func (s *UpdateDIJobRequestJobSettingsCycleScheduleSettings) SetScheduleParameters(v string) *UpdateDIJobRequestJobSettingsCycleScheduleSettings {
	s.ScheduleParameters = &v
	return s
}

func (s *UpdateDIJobRequestJobSettingsCycleScheduleSettings) Validate() error {
	return dara.Validate(s)
}

type UpdateDIJobRequestJobSettingsDdlHandlingSettings struct {
	// The handling action. Valid values:
	//
	// - `Ignore`: Ignores the DDL message.
	//
	// - `Critical`: Reports an error and terminates the synchronization job.
	//
	// - `Normal`: Processes the DDL message normally.
	//
	// example:
	//
	// Critical
	Action *string `json:"Action,omitempty" xml:"Action,omitempty"`
	// The DDL type. Valid values:
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
	// - `DropColumn`
	//
	// - `AddColumn`
	//
	// example:
	//
	// AddColumn
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s UpdateDIJobRequestJobSettingsDdlHandlingSettings) String() string {
	return dara.Prettify(s)
}

func (s UpdateDIJobRequestJobSettingsDdlHandlingSettings) GoString() string {
	return s.String()
}

func (s *UpdateDIJobRequestJobSettingsDdlHandlingSettings) GetAction() *string {
	return s.Action
}

func (s *UpdateDIJobRequestJobSettingsDdlHandlingSettings) GetType() *string {
	return s.Type
}

func (s *UpdateDIJobRequestJobSettingsDdlHandlingSettings) SetAction(v string) *UpdateDIJobRequestJobSettingsDdlHandlingSettings {
	s.Action = &v
	return s
}

func (s *UpdateDIJobRequestJobSettingsDdlHandlingSettings) SetType(v string) *UpdateDIJobRequestJobSettingsDdlHandlingSettings {
	s.Type = &v
	return s
}

func (s *UpdateDIJobRequestJobSettingsDdlHandlingSettings) Validate() error {
	return dara.Validate(s)
}

type UpdateDIJobRequestJobSettingsRuntimeSettings struct {
	// The name of the setting. Valid values:
	//
	// - `src.offline.datasource.max.connection`: The maximum number of concurrent connections to the source for an offline synchronization job.
	//
	// - `dst.offline.truncate`: Specifies whether to truncate the destination table before an offline synchronization job.
	//
	// - `runtime.offline.speed.limit.enable`: Specifies whether to enable speed limiting for an offline synchronization job.
	//
	// - `runtime.offline.concurrent`: The concurrency level for an offline synchronization job.
	//
	// - `runtime.enable.auto.create.schema`: Specifies whether to automatically create a schema at the destination.
	//
	// - `runtime.realtime.concurrent`: The concurrency level for a real-time synchronization job.
	//
	// - `runtime.realtime.failover.minute.dataxcdc`: The number of minutes to wait before a failover retry.
	//
	// - `runtime.realtime.failover.times.dataxcdc`: The number of failover retries.
	//
	// example:
	//
	// src.offline.datasource.max.connection
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The value of the setting.
	//
	// example:
	//
	// 1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateDIJobRequestJobSettingsRuntimeSettings) String() string {
	return dara.Prettify(s)
}

func (s UpdateDIJobRequestJobSettingsRuntimeSettings) GoString() string {
	return s.String()
}

func (s *UpdateDIJobRequestJobSettingsRuntimeSettings) GetName() *string {
	return s.Name
}

func (s *UpdateDIJobRequestJobSettingsRuntimeSettings) GetValue() *string {
	return s.Value
}

func (s *UpdateDIJobRequestJobSettingsRuntimeSettings) SetName(v string) *UpdateDIJobRequestJobSettingsRuntimeSettings {
	s.Name = &v
	return s
}

func (s *UpdateDIJobRequestJobSettingsRuntimeSettings) SetValue(v string) *UpdateDIJobRequestJobSettingsRuntimeSettings {
	s.Value = &v
	return s
}

func (s *UpdateDIJobRequestJobSettingsRuntimeSettings) Validate() error {
	return dara.Validate(s)
}

type UpdateDIJobRequestResourceSettings struct {
	// The resource settings for the offline synchronization job.
	OfflineResourceSettings *UpdateDIJobRequestResourceSettingsOfflineResourceSettings `json:"OfflineResourceSettings,omitempty" xml:"OfflineResourceSettings,omitempty" type:"Struct"`
	// The resource settings for the real-time synchronization job.
	RealtimeResourceSettings *UpdateDIJobRequestResourceSettingsRealtimeResourceSettings `json:"RealtimeResourceSettings,omitempty" xml:"RealtimeResourceSettings,omitempty" type:"Struct"`
	// The scheduling resource settings.
	ScheduleResourceSettings *UpdateDIJobRequestResourceSettingsScheduleResourceSettings `json:"ScheduleResourceSettings,omitempty" xml:"ScheduleResourceSettings,omitempty" type:"Struct"`
}

func (s UpdateDIJobRequestResourceSettings) String() string {
	return dara.Prettify(s)
}

func (s UpdateDIJobRequestResourceSettings) GoString() string {
	return s.String()
}

func (s *UpdateDIJobRequestResourceSettings) GetOfflineResourceSettings() *UpdateDIJobRequestResourceSettingsOfflineResourceSettings {
	return s.OfflineResourceSettings
}

func (s *UpdateDIJobRequestResourceSettings) GetRealtimeResourceSettings() *UpdateDIJobRequestResourceSettingsRealtimeResourceSettings {
	return s.RealtimeResourceSettings
}

func (s *UpdateDIJobRequestResourceSettings) GetScheduleResourceSettings() *UpdateDIJobRequestResourceSettingsScheduleResourceSettings {
	return s.ScheduleResourceSettings
}

func (s *UpdateDIJobRequestResourceSettings) SetOfflineResourceSettings(v *UpdateDIJobRequestResourceSettingsOfflineResourceSettings) *UpdateDIJobRequestResourceSettings {
	s.OfflineResourceSettings = v
	return s
}

func (s *UpdateDIJobRequestResourceSettings) SetRealtimeResourceSettings(v *UpdateDIJobRequestResourceSettingsRealtimeResourceSettings) *UpdateDIJobRequestResourceSettings {
	s.RealtimeResourceSettings = v
	return s
}

func (s *UpdateDIJobRequestResourceSettings) SetScheduleResourceSettings(v *UpdateDIJobRequestResourceSettingsScheduleResourceSettings) *UpdateDIJobRequestResourceSettings {
	s.ScheduleResourceSettings = v
	return s
}

func (s *UpdateDIJobRequestResourceSettings) Validate() error {
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

type UpdateDIJobRequestResourceSettingsOfflineResourceSettings struct {
	// The number of CUs for the Data Integration resource group used by the offline synchronization job.
	//
	// example:
	//
	// 2
	RequestedCu *float64 `json:"RequestedCu,omitempty" xml:"RequestedCu,omitempty"`
	// The name of the Data Integration resource group used by the offline synchronization job.
	//
	// example:
	//
	// S_res_group_111_222
	ResourceGroupIdentifier *string `json:"ResourceGroupIdentifier,omitempty" xml:"ResourceGroupIdentifier,omitempty"`
}

func (s UpdateDIJobRequestResourceSettingsOfflineResourceSettings) String() string {
	return dara.Prettify(s)
}

func (s UpdateDIJobRequestResourceSettingsOfflineResourceSettings) GoString() string {
	return s.String()
}

func (s *UpdateDIJobRequestResourceSettingsOfflineResourceSettings) GetRequestedCu() *float64 {
	return s.RequestedCu
}

func (s *UpdateDIJobRequestResourceSettingsOfflineResourceSettings) GetResourceGroupIdentifier() *string {
	return s.ResourceGroupIdentifier
}

func (s *UpdateDIJobRequestResourceSettingsOfflineResourceSettings) SetRequestedCu(v float64) *UpdateDIJobRequestResourceSettingsOfflineResourceSettings {
	s.RequestedCu = &v
	return s
}

func (s *UpdateDIJobRequestResourceSettingsOfflineResourceSettings) SetResourceGroupIdentifier(v string) *UpdateDIJobRequestResourceSettingsOfflineResourceSettings {
	s.ResourceGroupIdentifier = &v
	return s
}

func (s *UpdateDIJobRequestResourceSettingsOfflineResourceSettings) Validate() error {
	return dara.Validate(s)
}

type UpdateDIJobRequestResourceSettingsRealtimeResourceSettings struct {
	// The number of CUs for the Data Integration resource group used by the real-time synchronization job.
	//
	// example:
	//
	// 2
	RequestedCu *float64 `json:"RequestedCu,omitempty" xml:"RequestedCu,omitempty"`
	// The name of the Data Integration resource group used by the real-time synchronization job.
	//
	// example:
	//
	// S_res_group_111_222
	ResourceGroupIdentifier *string `json:"ResourceGroupIdentifier,omitempty" xml:"ResourceGroupIdentifier,omitempty"`
}

func (s UpdateDIJobRequestResourceSettingsRealtimeResourceSettings) String() string {
	return dara.Prettify(s)
}

func (s UpdateDIJobRequestResourceSettingsRealtimeResourceSettings) GoString() string {
	return s.String()
}

func (s *UpdateDIJobRequestResourceSettingsRealtimeResourceSettings) GetRequestedCu() *float64 {
	return s.RequestedCu
}

func (s *UpdateDIJobRequestResourceSettingsRealtimeResourceSettings) GetResourceGroupIdentifier() *string {
	return s.ResourceGroupIdentifier
}

func (s *UpdateDIJobRequestResourceSettingsRealtimeResourceSettings) SetRequestedCu(v float64) *UpdateDIJobRequestResourceSettingsRealtimeResourceSettings {
	s.RequestedCu = &v
	return s
}

func (s *UpdateDIJobRequestResourceSettingsRealtimeResourceSettings) SetResourceGroupIdentifier(v string) *UpdateDIJobRequestResourceSettingsRealtimeResourceSettings {
	s.ResourceGroupIdentifier = &v
	return s
}

func (s *UpdateDIJobRequestResourceSettingsRealtimeResourceSettings) Validate() error {
	return dara.Validate(s)
}

type UpdateDIJobRequestResourceSettingsScheduleResourceSettings struct {
	// The number of CUs for the scheduling resource group used by the offline synchronization job.
	//
	// example:
	//
	// 2
	RequestedCu *float64 `json:"RequestedCu,omitempty" xml:"RequestedCu,omitempty"`
	// The name of the scheduling resource group used by the offline synchronization job.
	//
	// example:
	//
	// S_res_group_222_333
	ResourceGroupIdentifier *string `json:"ResourceGroupIdentifier,omitempty" xml:"ResourceGroupIdentifier,omitempty"`
}

func (s UpdateDIJobRequestResourceSettingsScheduleResourceSettings) String() string {
	return dara.Prettify(s)
}

func (s UpdateDIJobRequestResourceSettingsScheduleResourceSettings) GoString() string {
	return s.String()
}

func (s *UpdateDIJobRequestResourceSettingsScheduleResourceSettings) GetRequestedCu() *float64 {
	return s.RequestedCu
}

func (s *UpdateDIJobRequestResourceSettingsScheduleResourceSettings) GetResourceGroupIdentifier() *string {
	return s.ResourceGroupIdentifier
}

func (s *UpdateDIJobRequestResourceSettingsScheduleResourceSettings) SetRequestedCu(v float64) *UpdateDIJobRequestResourceSettingsScheduleResourceSettings {
	s.RequestedCu = &v
	return s
}

func (s *UpdateDIJobRequestResourceSettingsScheduleResourceSettings) SetResourceGroupIdentifier(v string) *UpdateDIJobRequestResourceSettingsScheduleResourceSettings {
	s.ResourceGroupIdentifier = &v
	return s
}

func (s *UpdateDIJobRequestResourceSettingsScheduleResourceSettings) Validate() error {
	return dara.Validate(s)
}

type UpdateDIJobRequestTableMappings struct {
	// The rules for selecting source objects. Each rule can select a different type of source object to synchronize, such as a source database or table.
	SourceObjectSelectionRules []*UpdateDIJobRequestTableMappingsSourceObjectSelectionRules `json:"SourceObjectSelectionRules,omitempty" xml:"SourceObjectSelectionRules,omitempty" type:"Repeated"`
	// The transformation rules for the source objects.
	TransformationRules []*UpdateDIJobRequestTableMappingsTransformationRules `json:"TransformationRules,omitempty" xml:"TransformationRules,omitempty" type:"Repeated"`
}

func (s UpdateDIJobRequestTableMappings) String() string {
	return dara.Prettify(s)
}

func (s UpdateDIJobRequestTableMappings) GoString() string {
	return s.String()
}

func (s *UpdateDIJobRequestTableMappings) GetSourceObjectSelectionRules() []*UpdateDIJobRequestTableMappingsSourceObjectSelectionRules {
	return s.SourceObjectSelectionRules
}

func (s *UpdateDIJobRequestTableMappings) GetTransformationRules() []*UpdateDIJobRequestTableMappingsTransformationRules {
	return s.TransformationRules
}

func (s *UpdateDIJobRequestTableMappings) SetSourceObjectSelectionRules(v []*UpdateDIJobRequestTableMappingsSourceObjectSelectionRules) *UpdateDIJobRequestTableMappings {
	s.SourceObjectSelectionRules = v
	return s
}

func (s *UpdateDIJobRequestTableMappings) SetTransformationRules(v []*UpdateDIJobRequestTableMappingsTransformationRules) *UpdateDIJobRequestTableMappings {
	s.TransformationRules = v
	return s
}

func (s *UpdateDIJobRequestTableMappings) Validate() error {
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

type UpdateDIJobRequestTableMappingsSourceObjectSelectionRules struct {
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
	// The type of the expression. Valid values: `Exact` and `Regex`.
	//
	// example:
	//
	// Exact
	ExpressionType *string `json:"ExpressionType,omitempty" xml:"ExpressionType,omitempty"`
	// The object type. Valid values:
	//
	// - `Table` (table)
	//
	// - `Schema` (schema)
	//
	// - `Database` (database)
	//
	// example:
	//
	// Table
	ObjectType *string `json:"ObjectType,omitempty" xml:"ObjectType,omitempty"`
}

func (s UpdateDIJobRequestTableMappingsSourceObjectSelectionRules) String() string {
	return dara.Prettify(s)
}

func (s UpdateDIJobRequestTableMappingsSourceObjectSelectionRules) GoString() string {
	return s.String()
}

func (s *UpdateDIJobRequestTableMappingsSourceObjectSelectionRules) GetAction() *string {
	return s.Action
}

func (s *UpdateDIJobRequestTableMappingsSourceObjectSelectionRules) GetExpression() *string {
	return s.Expression
}

func (s *UpdateDIJobRequestTableMappingsSourceObjectSelectionRules) GetExpressionType() *string {
	return s.ExpressionType
}

func (s *UpdateDIJobRequestTableMappingsSourceObjectSelectionRules) GetObjectType() *string {
	return s.ObjectType
}

func (s *UpdateDIJobRequestTableMappingsSourceObjectSelectionRules) SetAction(v string) *UpdateDIJobRequestTableMappingsSourceObjectSelectionRules {
	s.Action = &v
	return s
}

func (s *UpdateDIJobRequestTableMappingsSourceObjectSelectionRules) SetExpression(v string) *UpdateDIJobRequestTableMappingsSourceObjectSelectionRules {
	s.Expression = &v
	return s
}

func (s *UpdateDIJobRequestTableMappingsSourceObjectSelectionRules) SetExpressionType(v string) *UpdateDIJobRequestTableMappingsSourceObjectSelectionRules {
	s.ExpressionType = &v
	return s
}

func (s *UpdateDIJobRequestTableMappingsSourceObjectSelectionRules) SetObjectType(v string) *UpdateDIJobRequestTableMappingsSourceObjectSelectionRules {
	s.ObjectType = &v
	return s
}

func (s *UpdateDIJobRequestTableMappingsSourceObjectSelectionRules) Validate() error {
	return dara.Validate(s)
}

type UpdateDIJobRequestTableMappingsTransformationRules struct {
	// The type of the action. Valid values:
	//
	// - `DefinePrimaryKey`: Defines a primary key.
	//
	// - `Rename`: Renames an object.
	//
	// - `AddColumn`: Adds a column.
	//
	// - `HandleDml`: Handles DML operations.
	//
	// example:
	//
	// Rename
	RuleActionType *string `json:"RuleActionType,omitempty" xml:"RuleActionType,omitempty"`
	// The name of the transformation rule. The name must be unique for a specific combination of `RuleActionType` and `RuleTargetType` and can be up to 50 characters long.
	//
	// example:
	//
	// rename_rule_1
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The type of the target object. Valid values:
	//
	// - `Table` (table)
	//
	// - `Schema` (schema)
	//
	// - `Database` (database)
	//
	// example:
	//
	// Table
	RuleTargetType *string `json:"RuleTargetType,omitempty" xml:"RuleTargetType,omitempty"`
}

func (s UpdateDIJobRequestTableMappingsTransformationRules) String() string {
	return dara.Prettify(s)
}

func (s UpdateDIJobRequestTableMappingsTransformationRules) GoString() string {
	return s.String()
}

func (s *UpdateDIJobRequestTableMappingsTransformationRules) GetRuleActionType() *string {
	return s.RuleActionType
}

func (s *UpdateDIJobRequestTableMappingsTransformationRules) GetRuleName() *string {
	return s.RuleName
}

func (s *UpdateDIJobRequestTableMappingsTransformationRules) GetRuleTargetType() *string {
	return s.RuleTargetType
}

func (s *UpdateDIJobRequestTableMappingsTransformationRules) SetRuleActionType(v string) *UpdateDIJobRequestTableMappingsTransformationRules {
	s.RuleActionType = &v
	return s
}

func (s *UpdateDIJobRequestTableMappingsTransformationRules) SetRuleName(v string) *UpdateDIJobRequestTableMappingsTransformationRules {
	s.RuleName = &v
	return s
}

func (s *UpdateDIJobRequestTableMappingsTransformationRules) SetRuleTargetType(v string) *UpdateDIJobRequestTableMappingsTransformationRules {
	s.RuleTargetType = &v
	return s
}

func (s *UpdateDIJobRequestTableMappingsTransformationRules) Validate() error {
	return dara.Validate(s)
}

type UpdateDIJobRequestTransformationRules struct {
	// The type of the action. Valid values:
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
	// 1. Rename rule (`Rename`)
	//
	// - Example: `{"expression":"${srcDatasourceName}_${srcDatabaseName}_0922"}`
	//
	// - `expression`: The expression for the rename transformation rule. The expression supports variables, including `${srcDatasourceName}` (source data source name), `${srcDatabaseName}` (source database name), and `${srcTableName}` (source table name).
	//
	// 2. Add column rule (`AddColumn`)
	//
	// - Example: `{"columns":[{"columnName":"my_add_column","columnValueType":"Constant","columnValue":"123"}]}`
	//
	// - If this rule is not specified, no columns are added.
	//
	// - `columnName`: The name of the column to add.
	//
	// - `columnValueType`: The value type of the added column. Valid values: `Constant` and `Variable`.
	//
	// - `columnValue`: The value of the added column. If `columnValueType` is `Constant`, the value is a custom constant of the string type. If `columnValueType` is `Variable`, the value is a built-in variable. Valid built-in variables: `EXECUTE_TIME` (execution time, Long type), `DB_NAME_SRC` (source database name, String type), `DATASOURCE_NAME_SRC` (source data source name, String type), `TABLE_NAME_SRC` (source table name, String type), `DB_NAME_DEST` (destination database name, String type), `DATASOURCE_NAME_DEST` (destination data source name, String type), `TABLE_NAME_DEST` (destination table name, String type), and `DB_NAME_SRC_TRANSED` (transformed database name, String type).
	//
	// 3. Define primary key rule (`DefinePrimaryKey`)
	//
	// - Example: `{"columns":["ukcolumn1","ukcolumn2"]}`
	//
	// - If this rule is not specified, the primary key of the source is used by default.
	//
	// - Data Integration does not modify the structure of an existing destination table. If a specified primary key column does not exist in the table, the synchronization job fails.
	//
	// - When a destination table is automatically created, Data Integration includes the defined primary key columns in the structure. If a specified primary key column is not in the destination column set, the synchronization job fails.
	//
	// 4. DML handling rule (`HandleDml`)
	//
	// - Example: `{"dmlPolicies":[{"dmlType":"Delete","dmlAction":"Filter","filterCondition":"id > 1"}]}`
	//
	// - If this rule is not specified, the default action for `Insert`, `Update`, and `Delete` operations is `Normal`.
	//
	// - `dmlType`: The DML operation type. Valid values: `Insert`, `Update`, and `Delete`.
	//
	// - `dmlAction`: The DML handling policy. Valid values: `Normal` (process the operation), `Ignore` (ignore the operation), `Filter` (conditionally process the operation, used when `dmlType` is `Update` or `Delete`), and `LogicalDelete` (perform a logical delete).
	//
	// - `filterCondition`: The DML filter condition, used when `dmlAction` is `Filter`.
	//
	// 5. Incremental condition rule (`DefineIncrementalCondition`)
	//
	// - Example: `{"where":"id > 0"}`
	//
	// - The `WHERE` clause for the incremental condition.
	//
	// 6. Periodic scheduling rule (`DefineCycleScheduleSettings`)
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
	// The name of the transformation rule. The name must be unique for a specific combination of `RuleActionType` and `RuleTargetType` and can be up to 50 characters long.
	//
	// example:
	//
	// rename_rule_1
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The type of the target object. Valid values:
	//
	// - `Table` (table)
	//
	// - `Schema` (schema)
	//
	// - `Database` (database)
	//
	// example:
	//
	// Table
	RuleTargetType *string `json:"RuleTargetType,omitempty" xml:"RuleTargetType,omitempty"`
}

func (s UpdateDIJobRequestTransformationRules) String() string {
	return dara.Prettify(s)
}

func (s UpdateDIJobRequestTransformationRules) GoString() string {
	return s.String()
}

func (s *UpdateDIJobRequestTransformationRules) GetRuleActionType() *string {
	return s.RuleActionType
}

func (s *UpdateDIJobRequestTransformationRules) GetRuleExpression() *string {
	return s.RuleExpression
}

func (s *UpdateDIJobRequestTransformationRules) GetRuleName() *string {
	return s.RuleName
}

func (s *UpdateDIJobRequestTransformationRules) GetRuleTargetType() *string {
	return s.RuleTargetType
}

func (s *UpdateDIJobRequestTransformationRules) SetRuleActionType(v string) *UpdateDIJobRequestTransformationRules {
	s.RuleActionType = &v
	return s
}

func (s *UpdateDIJobRequestTransformationRules) SetRuleExpression(v string) *UpdateDIJobRequestTransformationRules {
	s.RuleExpression = &v
	return s
}

func (s *UpdateDIJobRequestTransformationRules) SetRuleName(v string) *UpdateDIJobRequestTransformationRules {
	s.RuleName = &v
	return s
}

func (s *UpdateDIJobRequestTransformationRules) SetRuleTargetType(v string) *UpdateDIJobRequestTransformationRules {
	s.RuleTargetType = &v
	return s
}

func (s *UpdateDIJobRequestTransformationRules) Validate() error {
	return dara.Validate(s)
}
