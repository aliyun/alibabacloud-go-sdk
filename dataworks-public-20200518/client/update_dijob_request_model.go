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
	SetJobSettings(v *UpdateDIJobRequestJobSettings) *UpdateDIJobRequest
	GetJobSettings() *UpdateDIJobRequestJobSettings
	SetResourceSettings(v *UpdateDIJobRequestResourceSettings) *UpdateDIJobRequest
	GetResourceSettings() *UpdateDIJobRequestResourceSettings
	SetTableMappings(v []*UpdateDIJobRequestTableMappings) *UpdateDIJobRequest
	GetTableMappings() []*UpdateDIJobRequestTableMappings
	SetTransformationRules(v []*UpdateDIJobRequestTransformationRules) *UpdateDIJobRequest
	GetTransformationRules() []*UpdateDIJobRequestTransformationRules
}

type UpdateDIJobRequest struct {
	// The ID of the synchronization task.
	//
	// example:
	//
	// 11588
	DIJobId *int64 `json:"DIJobId,omitempty" xml:"DIJobId,omitempty"`
	// The description of the synchronization task.
	//
	// example:
	//
	// Synchronize mysql to hologres
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The settings for the dimension of the synchronization task. The settings include processing policies for DDL messages, policies for data type mappings between source fields and destination fields, and runtime parameters of the synchronization task.
	JobSettings *UpdateDIJobRequestJobSettings `json:"JobSettings,omitempty" xml:"JobSettings,omitempty" type:"Struct"`
	// The resource settings.
	ResourceSettings *UpdateDIJobRequestResourceSettings `json:"ResourceSettings,omitempty" xml:"ResourceSettings,omitempty" type:"Struct"`
	// The list of mappings between rules used to select synchronization objects in the source and transformation rules applied to the selected synchronization objects. Each entry in the list displays a mapping between a rule used to select synchronization objects and a transformation rule applied to the selected synchronization objects.
	TableMappings []*UpdateDIJobRequestTableMappings `json:"TableMappings,omitempty" xml:"TableMappings,omitempty" type:"Repeated"`
	// The list of transformation rules that you want to apply to the synchronization objects selected from the source. Each entry in the list defines a transformation rule.
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

func (s *UpdateDIJobRequest) GetJobSettings() *UpdateDIJobRequestJobSettings {
	return s.JobSettings
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

func (s *UpdateDIJobRequest) SetJobSettings(v *UpdateDIJobRequestJobSettings) *UpdateDIJobRequest {
	s.JobSettings = v
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
	return dara.Validate(s)
}

type UpdateDIJobRequestJobSettings struct {
	// The channel control settings for the synchronization task. The value of this parameter must be a JSON string.
	//
	// example:
	//
	// {"structInfo":"MANAGED","storageType":"TEXTFILE","writeMode":"APPEND","partitionColumns":[{"columnName":"pt","columnType":"STRING","comment":""}],"fieldDelimiter":""}
	ChannelSettings *string `json:"ChannelSettings,omitempty" xml:"ChannelSettings,omitempty"`
	// The settings for data type mappings between source fields and destination fields. The value of this parameter must be an array.
	ColumnDataTypeSettings []*UpdateDIJobRequestJobSettingsColumnDataTypeSettings `json:"ColumnDataTypeSettings,omitempty" xml:"ColumnDataTypeSettings,omitempty" type:"Repeated"`
	// The settings for periodic scheduling.
	CycleScheduleSettings *UpdateDIJobRequestJobSettingsCycleScheduleSettings `json:"CycleScheduleSettings,omitempty" xml:"CycleScheduleSettings,omitempty" type:"Struct"`
	// The settings for processing DDL messages. The value of this parameter must be an array.
	DdlHandlingSettings []*UpdateDIJobRequestJobSettingsDdlHandlingSettings `json:"DdlHandlingSettings,omitempty" xml:"DdlHandlingSettings,omitempty" type:"Repeated"`
	// The runtime settings. The value of this parameter must be an array.
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
	return dara.Validate(s)
}

type UpdateDIJobRequestJobSettingsColumnDataTypeSettings struct {
	// The data type of a destination field.
	//
	// example:
	//
	// text
	DestinationDataType *string `json:"DestinationDataType,omitempty" xml:"DestinationDataType,omitempty"`
	// The data type of a source field.
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
	// The processing policy for DDL messages. Valid values:
	//
	// 	- Ignore: ignores a DDL message.
	//
	// 	- Critical: reports an error for a DDL message.
	//
	// 	- Normal: normally processes a DDL message.
	//
	// example:
	//
	// Critical
	Action *string `json:"Action,omitempty" xml:"Action,omitempty"`
	// The type of the DDL operation. Valid values:
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
	// The name of the configuration item. Valid values:
	//
	// 	- runtime.offline.speed.limit.mb: specifies the maximum transmission rate that is allowed for a batch synchronization task. This configuration item takes effect only when runtime.offline.speed.limit.enable is set to true.
	//
	// 	- runtime.offline.speed.limit.enable: specifies whether throttling is enabled for a batch synchronization task.
	//
	// 	- dst.offline.connection.max: specifies the maximum number of connections that are allowed for writing data to the destination of a batch synchronization task.
	//
	// 	- runtime.offline.concurrent: specifies the maximum number of parallel threads that are allowed for a batch synchronization task.
	//
	// 	- dst.realtime.connection.max: specifies the maximum number of connections that are allowed for writing data to the destination of a real-time synchronization task.
	//
	// 	- runtime.enable.auto.create.schema: specifies whether schemas are automatically created in the destination of a synchronization task.
	//
	// 	- src.offline.datasource.max.connection: specifies the maximum number of connections that are allowed for reading data from the source of a batch synchronization task.
	//
	// 	- runtime.realtime.concurrent: specifies the maximum number of parallel threads that are allowed for a real-time synchronization task.
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
	// The resource used for batch synchronization.
	OfflineResourceSettings *UpdateDIJobRequestResourceSettingsOfflineResourceSettings `json:"OfflineResourceSettings,omitempty" xml:"OfflineResourceSettings,omitempty" type:"Struct"`
	// The resource used for real-time synchronization.
	RealtimeResourceSettings *UpdateDIJobRequestResourceSettingsRealtimeResourceSettings `json:"RealtimeResourceSettings,omitempty" xml:"RealtimeResourceSettings,omitempty" type:"Struct"`
	// The number of compute units (CUs) in the resource group that are used for full and incremental synchronization.
	//
	// example:
	//
	// 2.0
	RequestedCu *float32 `json:"RequestedCu,omitempty" xml:"RequestedCu,omitempty"`
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

func (s *UpdateDIJobRequestResourceSettings) GetRequestedCu() *float32 {
	return s.RequestedCu
}

func (s *UpdateDIJobRequestResourceSettings) SetOfflineResourceSettings(v *UpdateDIJobRequestResourceSettingsOfflineResourceSettings) *UpdateDIJobRequestResourceSettings {
	s.OfflineResourceSettings = v
	return s
}

func (s *UpdateDIJobRequestResourceSettings) SetRealtimeResourceSettings(v *UpdateDIJobRequestResourceSettingsRealtimeResourceSettings) *UpdateDIJobRequestResourceSettings {
	s.RealtimeResourceSettings = v
	return s
}

func (s *UpdateDIJobRequestResourceSettings) SetRequestedCu(v float32) *UpdateDIJobRequestResourceSettings {
	s.RequestedCu = &v
	return s
}

func (s *UpdateDIJobRequestResourceSettings) Validate() error {
	return dara.Validate(s)
}

type UpdateDIJobRequestResourceSettingsOfflineResourceSettings struct {
	// The identifier of the resource group for Data Integration used for batch synchronization.
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

func (s *UpdateDIJobRequestResourceSettingsOfflineResourceSettings) GetResourceGroupIdentifier() *string {
	return s.ResourceGroupIdentifier
}

func (s *UpdateDIJobRequestResourceSettingsOfflineResourceSettings) SetResourceGroupIdentifier(v string) *UpdateDIJobRequestResourceSettingsOfflineResourceSettings {
	s.ResourceGroupIdentifier = &v
	return s
}

func (s *UpdateDIJobRequestResourceSettingsOfflineResourceSettings) Validate() error {
	return dara.Validate(s)
}

type UpdateDIJobRequestResourceSettingsRealtimeResourceSettings struct {
	// The identifier of the resource group for Data Integration used for real-time synchronization.
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

func (s *UpdateDIJobRequestResourceSettingsRealtimeResourceSettings) GetResourceGroupIdentifier() *string {
	return s.ResourceGroupIdentifier
}

func (s *UpdateDIJobRequestResourceSettingsRealtimeResourceSettings) SetResourceGroupIdentifier(v string) *UpdateDIJobRequestResourceSettingsRealtimeResourceSettings {
	s.ResourceGroupIdentifier = &v
	return s
}

func (s *UpdateDIJobRequestResourceSettingsRealtimeResourceSettings) Validate() error {
	return dara.Validate(s)
}

type UpdateDIJobRequestTableMappings struct {
	// The list of rules that you want to use to select synchronization objects in the source.
	SourceObjectSelectionRules []*UpdateDIJobRequestTableMappingsSourceObjectSelectionRules `json:"SourceObjectSelectionRules,omitempty" xml:"SourceObjectSelectionRules,omitempty" type:"Repeated"`
	// The list of transformation rules that you want to apply to the synchronization objects selected from the source.
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
	return dara.Validate(s)
}

type UpdateDIJobRequestTableMappingsSourceObjectSelectionRules struct {
	// The expression.
	//
	// example:
	//
	// mysql_table_1
	Expression *string `json:"Expression,omitempty" xml:"Expression,omitempty"`
	// The type of the object. Valid values:
	//
	// 	- Table
	//
	// 	- Database
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

func (s *UpdateDIJobRequestTableMappingsSourceObjectSelectionRules) GetExpression() *string {
	return s.Expression
}

func (s *UpdateDIJobRequestTableMappingsSourceObjectSelectionRules) GetObjectType() *string {
	return s.ObjectType
}

func (s *UpdateDIJobRequestTableMappingsSourceObjectSelectionRules) SetExpression(v string) *UpdateDIJobRequestTableMappingsSourceObjectSelectionRules {
	s.Expression = &v
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
	// 	- DefineRuntimeSettings
	//
	// 	- DefinePartitionKey
	//
	// example:
	//
	// Rename
	RuleActionType *string `json:"RuleActionType,omitempty" xml:"RuleActionType,omitempty"`
	// The name of the rule. If the values of the RuleActionType parameter and the RuleTargetType parameter are the same for multiple transformation rules, you must make sure that the transformation rule names are unique.
	//
	// example:
	//
	// rename_rule_1
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The type of the object on which you want to perform the action. Valid values:
	//
	// 	- Table
	//
	// 	- Schema
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
	// 	- DefineRuntimeSettings
	//
	// 	- DefinePartitionKey
	//
	// example:
	//
	// Rename
	RuleActionType *string `json:"RuleActionType,omitempty" xml:"RuleActionType,omitempty"`
	// The expression of the rule. The expression must be a JSON string.
	//
	// Example of a renaming rule: {"expression":"${srcDatasourceName}_${srcDatabaseName}_0922","variables":[{"variableName":"srcDatabaseName","variableRules":[{"from":"fromdb","to":"todb"}]}]}.
	//
	// expression: the expression of the renaming rule. The expression may contain the following variables: ${srcDatasourceName}, ${srcDatabaseName}, and ${srcTableName}. ${srcDatasourceName} specifies the name of the source. ${srcDatabaseName} specifies the name of a source database. ${srcTableName} specifies the name of a source table. variables: the generation rule for a variable used in the expression of the renaming rule. The default value of the specified variable is the original value of the object indicated by the variable. You can define a group of string replacement rules to change the original values based on your business requirements. variableName: the name of the variable. Do not enclose the variable name in ${}. variableRules: the string replacement rules for variables. The system runs the string replacement rules in sequence. from specifies the original string. to specifies the new string. Example of a rule used to add a specific field to the destination and assign a value to the field: {"columns":[{"columnName":"my_add_column","columnValueType":"Constant","columnValue":"123"}]}.
	//
	// If you do not configure such a rule, no fields are added to the destination and no values are assigned by default. columnName: the name of the field that you want to add. columnValueType: the value type of the field. Valid values: Constant and Variable. columnValue: the value of the field that you want to add. If you set the valueType parameter to Constant, set the columnValue parameter to a custom constant of the STRING type. If you set the valueType parameter to Variable, set the columnValue to a built-in variable. The following built-in variables are supported: EXECUTE_TIME (LONG data type), DB_NAME_SRC (STRING data type), DATASOURCE_NAME_SRC (STRING data type), TABLE_NAME_SRC (STRING data type), DB_NAME_DEST (STRING data type), DATASOURCE_NAME_DEST (STRING data type), TABLE_NAME_DEST (STRING data type), and DB_NAME_SRC_TRANSED (STRING data type). EXECUTE_TIME specifies the execution time. DB_NAME_SRC indicates the name of a source database. DATASOURCE_NAME_SRC specifies the name of the source. TABLE_NAME_SRC specifies the name of a source table. DB_NAME_DEST specifies the name of a destination database. DATASOURCE_NAME_DEST specifies the name of the destination. TABLE_NAME_DEST specifies the name of a destination table. DB_NAME_SRC_TRANSED specifies the database name obtained after a transformation. Example of a rule used to specify primary key fields for a destination table: {"columns":["ukcolumn1","ukcolumn2"]}.
	//
	// If you do not configure such a rule, the primary key fields in the mapped source table are used for the destination table by default. If the destination table is an existing table, Data Integration does not modify the schema of the destination table. If the specified primary key fields do not exist in the destination table, an error is reported when the synchronization task starts to run. If the destination table is automatically created by the system, Data Integration automatically creates the schema of the destination table. The schema contains the primary key fields that you specify. If the specified primary key fields do not exist in the destination table, an error is reported when the synchronization task starts to run. Example of a rule used to process DML messages: {"dmlPolicies":[{"dmlType":"Delete","dmlAction":"Filter","filterCondition":"id > 1"}]}.
	//
	// If you do not configure such a rule, the default processing policy for messages generated for insert, update, and delete operations is Normal. dmlType: the DML operation. Valid values: Insert, Update, and Delete. dmlAction: the processing policy for DML messages. Valid values: Normal, Ignore, Filter, and LogicalDelete. Filter indicates conditional processing. You can set the dmlAction parameter to Filter only when the dmlType parameter is set to Update or Delete. filterCondition: the condition used to filter DML messages. This parameter is required only when the dmlAction parameter is set to Filter.
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
	// The type of the object on which you want to perform the action. Valid values:
	//
	// 	- Table
	//
	// 	- Schema
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
