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
	// This parameter is deprecated. Use the Id parameter instead.
	//
	// example:
	//
	// 11588
	DIJobId *int64 `json:"DIJobId,omitempty" xml:"DIJobId,omitempty"`
	// The task description.
	//
	// example:
	//
	// The description of the synchronization task.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	FileSpec    *string `json:"FileSpec,omitempty" xml:"FileSpec,omitempty"`
	// The ID of the synchronization task.
	//
	// example:
	//
	// 11588
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The task-level settings, including DDL handling policies, column data type mapping between source and destination, and runtime parameters.
	JobSettings *UpdateDIJobRequestJobSettings `json:"JobSettings,omitempty" xml:"JobSettings,omitempty" type:"Struct"`
	// The task owner.
	//
	// example:
	//
	// 95279527
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The DataWorks workspace ID. You can call the [ListProjects](https://help.aliyun.com/document_detail/178393.html) operation to obtain the ID.
	//
	// example:
	//
	// 10000
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The resource settings.
	ResourceSettings *UpdateDIJobRequestResourceSettings `json:"ResourceSettings,omitempty" xml:"ResourceSettings,omitempty" type:"Struct"`
	// The list of synchronization object transformation mappings. Each element describes a set of source object selection rules and the transformation rules applied to those objects.
	//
	// >  [ { "SourceObjectSelectionRules":[ { "ObjectType":"Database", "Action":"Include", "ExpressionType":"Exact", "Expression":"biz_db" }, { "ObjectType":"Schema", "Action":"Include", "ExpressionType":"Exact", "Expression":"s1" }, { "ObjectType":"Table", "Action":"Include", "ExpressionType":"Exact", "Expression":"table1" } ], "TransformationRuleNames":[ { "RuleName":"my_database_rename_rule", "RuleActionType":"Rename", "RuleTargetType":"Schema" } ] } ]
	TableMappings []*UpdateDIJobRequestTableMappings `json:"TableMappings,omitempty" xml:"TableMappings,omitempty" type:"Repeated"`
	// The list of synchronization object transformation rule definitions.
	//
	// >  [ { "RuleName":"my_database_rename_rule", "RuleActionType":"Rename", "RuleTargetType":"Schema", "RuleExpression":"{"expression":"${srcDatasoureName}_${srcDatabaseName}"}" } ]
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
	ColumnDataTypeSettings []*UpdateDIJobRequestJobSettingsColumnDataTypeSettings `json:"ColumnDataTypeSettings,omitempty" xml:"ColumnDataTypeSettings,omitempty" type:"Repeated"`
	// The scheduled task settings.
	CycleScheduleSettings *UpdateDIJobRequestJobSettingsCycleScheduleSettings `json:"CycleScheduleSettings,omitempty" xml:"CycleScheduleSettings,omitempty" type:"Struct"`
	// The array of DDL handling settings.
	//
	// >  ["DDLHandlingSettings":[ { "Type":"Insert", "Action":"Normal" } ]
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
	// The batch synchronization resources.
	OfflineResourceSettings *UpdateDIJobRequestResourceSettingsOfflineResourceSettings `json:"OfflineResourceSettings,omitempty" xml:"OfflineResourceSettings,omitempty" type:"Struct"`
	// The real-time synchronization resources.
	RealtimeResourceSettings *UpdateDIJobRequestResourceSettingsRealtimeResourceSettings `json:"RealtimeResourceSettings,omitempty" xml:"RealtimeResourceSettings,omitempty" type:"Struct"`
	// The resource used for scheduling.
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
	// The CUs of the resource group for Data Integration used for batch synchronization.
	//
	// example:
	//
	// 2.0
	RequestedCu *float64 `json:"RequestedCu,omitempty" xml:"RequestedCu,omitempty"`
	// The name of the resource group for Data Integration that are used for batch synchronization.
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
	// The CUs of the resource group for Data Integration that are used for real-time synchronization.
	//
	// example:
	//
	// 2.0
	RequestedCu *float64 `json:"RequestedCu,omitempty" xml:"RequestedCu,omitempty"`
	// The name of the resource group for Data Integration that are used for real-time synchronization.
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
	// The CUs of the scheduling resource group for batch synchronization tasks.
	//
	// example:
	//
	// 2.0
	RequestedCu *float64 `json:"RequestedCu,omitempty" xml:"RequestedCu,omitempty"`
	// The name of the scheduling resource group used for batch synchronization tasks.
	//
	// example:
	//
	// S_res_group_235454102432001_1721021993437
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
	// Each rule can select different object types from the source, such as source databases and source tables.
	SourceObjectSelectionRules []*UpdateDIJobRequestTableMappingsSourceObjectSelectionRules `json:"SourceObjectSelectionRules,omitempty" xml:"SourceObjectSelectionRules,omitempty" type:"Repeated"`
	// The transformation rules applied to source objects.
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
	// Valid values: Include and Exclude.
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
	// Valid values:
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
	// Rename
	RuleActionType *string `json:"RuleActionType,omitempty" xml:"RuleActionType,omitempty"`
	// The rule name. The rule name must be unique for a given combination of action type and target type. The name cannot exceed 50 characters.
	//
	// example:
	//
	// rename_rule_1
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// Valid values:
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
	// Valid values:
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
	// 	- dmlType: The DML operation type. Valid values: Insert, Update, Delete.
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
