// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDIJobShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDIJobId(v int64) *UpdateDIJobShrinkRequest
	GetDIJobId() *int64
	SetDescription(v string) *UpdateDIJobShrinkRequest
	GetDescription() *string
	SetFileSpec(v string) *UpdateDIJobShrinkRequest
	GetFileSpec() *string
	SetId(v int64) *UpdateDIJobShrinkRequest
	GetId() *int64
	SetJobSettingsShrink(v string) *UpdateDIJobShrinkRequest
	GetJobSettingsShrink() *string
	SetOwner(v string) *UpdateDIJobShrinkRequest
	GetOwner() *string
	SetProjectId(v int64) *UpdateDIJobShrinkRequest
	GetProjectId() *int64
	SetResourceSettingsShrink(v string) *UpdateDIJobShrinkRequest
	GetResourceSettingsShrink() *string
	SetTableMappingsShrink(v string) *UpdateDIJobShrinkRequest
	GetTableMappingsShrink() *string
	SetTransformationRulesShrink(v string) *UpdateDIJobShrinkRequest
	GetTransformationRulesShrink() *string
}

type UpdateDIJobShrinkRequest struct {
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
	JobSettingsShrink *string `json:"JobSettings,omitempty" xml:"JobSettings,omitempty"`
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
	ResourceSettingsShrink *string `json:"ResourceSettings,omitempty" xml:"ResourceSettings,omitempty"`
	// A list of object transformation mappings. Each mapping specifies a set of selection rules for source objects and a list of transformation rules that apply to the selected objects.
	//
	// > [ { "SourceObjectSelectionRules":[ { "ObjectType":"Database", "Action":"Include", "ExpressionType":"Exact", "Expression":"biz_db" }, { "ObjectType":"Schema", "Action":"Include", "ExpressionType":"Exact", "Expression":"s1" }, { "ObjectType":"Table", "Action":"Include", "ExpressionType":"Exact", "Expression":"table1" } ], "TransformationRuleNames":[ { "RuleName":"my_database_rename_rule", "RuleActionType":"Rename", "RuleTargetType":"Schema" } ] } ]
	TableMappingsShrink *string `json:"TableMappings,omitempty" xml:"TableMappings,omitempty"`
	// A list of transformation rule definitions.
	//
	// > [ { "RuleName":"my_database_rename_rule", "RuleActionType":"Rename", "RuleTargetType":"Schema", "RuleExpression":"{"expression":"${srcDatasoureName}_${srcDatabaseName}"}" } ]
	TransformationRulesShrink *string `json:"TransformationRules,omitempty" xml:"TransformationRules,omitempty"`
}

func (s UpdateDIJobShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateDIJobShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateDIJobShrinkRequest) GetDIJobId() *int64 {
	return s.DIJobId
}

func (s *UpdateDIJobShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateDIJobShrinkRequest) GetFileSpec() *string {
	return s.FileSpec
}

func (s *UpdateDIJobShrinkRequest) GetId() *int64 {
	return s.Id
}

func (s *UpdateDIJobShrinkRequest) GetJobSettingsShrink() *string {
	return s.JobSettingsShrink
}

func (s *UpdateDIJobShrinkRequest) GetOwner() *string {
	return s.Owner
}

func (s *UpdateDIJobShrinkRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *UpdateDIJobShrinkRequest) GetResourceSettingsShrink() *string {
	return s.ResourceSettingsShrink
}

func (s *UpdateDIJobShrinkRequest) GetTableMappingsShrink() *string {
	return s.TableMappingsShrink
}

func (s *UpdateDIJobShrinkRequest) GetTransformationRulesShrink() *string {
	return s.TransformationRulesShrink
}

func (s *UpdateDIJobShrinkRequest) SetDIJobId(v int64) *UpdateDIJobShrinkRequest {
	s.DIJobId = &v
	return s
}

func (s *UpdateDIJobShrinkRequest) SetDescription(v string) *UpdateDIJobShrinkRequest {
	s.Description = &v
	return s
}

func (s *UpdateDIJobShrinkRequest) SetFileSpec(v string) *UpdateDIJobShrinkRequest {
	s.FileSpec = &v
	return s
}

func (s *UpdateDIJobShrinkRequest) SetId(v int64) *UpdateDIJobShrinkRequest {
	s.Id = &v
	return s
}

func (s *UpdateDIJobShrinkRequest) SetJobSettingsShrink(v string) *UpdateDIJobShrinkRequest {
	s.JobSettingsShrink = &v
	return s
}

func (s *UpdateDIJobShrinkRequest) SetOwner(v string) *UpdateDIJobShrinkRequest {
	s.Owner = &v
	return s
}

func (s *UpdateDIJobShrinkRequest) SetProjectId(v int64) *UpdateDIJobShrinkRequest {
	s.ProjectId = &v
	return s
}

func (s *UpdateDIJobShrinkRequest) SetResourceSettingsShrink(v string) *UpdateDIJobShrinkRequest {
	s.ResourceSettingsShrink = &v
	return s
}

func (s *UpdateDIJobShrinkRequest) SetTableMappingsShrink(v string) *UpdateDIJobShrinkRequest {
	s.TableMappingsShrink = &v
	return s
}

func (s *UpdateDIJobShrinkRequest) SetTransformationRulesShrink(v string) *UpdateDIJobShrinkRequest {
	s.TransformationRulesShrink = &v
	return s
}

func (s *UpdateDIJobShrinkRequest) Validate() error {
	return dara.Validate(s)
}
