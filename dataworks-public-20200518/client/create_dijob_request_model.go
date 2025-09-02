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
	SetJobName(v string) *CreateDIJobRequest
	GetJobName() *string
	SetJobSettings(v *CreateDIJobRequestJobSettings) *CreateDIJobRequest
	GetJobSettings() *CreateDIJobRequestJobSettings
	SetMigrationType(v string) *CreateDIJobRequest
	GetMigrationType() *string
	SetProjectId(v int64) *CreateDIJobRequest
	GetProjectId() *int64
	SetResourceSettings(v *CreateDIJobRequestResourceSettings) *CreateDIJobRequest
	GetResourceSettings() *CreateDIJobRequestResourceSettings
	SetSourceDataSourceSettings(v []*CreateDIJobRequestSourceDataSourceSettings) *CreateDIJobRequest
	GetSourceDataSourceSettings() []*CreateDIJobRequestSourceDataSourceSettings
	SetSourceDataSourceType(v string) *CreateDIJobRequest
	GetSourceDataSourceType() *string
	SetSystemDebug(v string) *CreateDIJobRequest
	GetSystemDebug() *string
	SetTableMappings(v []*CreateDIJobRequestTableMappings) *CreateDIJobRequest
	GetTableMappings() []*CreateDIJobRequestTableMappings
	SetTransformationRules(v []*CreateDIJobRequestTransformationRules) *CreateDIJobRequest
	GetTransformationRules() []*CreateDIJobRequestTransformationRules
}

type CreateDIJobRequest struct {
	// The description of the synchronization task.
	//
	// example:
	//
	// Synchronize mysql to hologres
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The settings of the destination. Only a single destination is supported.
	DestinationDataSourceSettings []*CreateDIJobRequestDestinationDataSourceSettings `json:"DestinationDataSourceSettings,omitempty" xml:"DestinationDataSourceSettings,omitempty" type:"Repeated"`
	// The destination type. Valid values: Hologres and Hive.
	//
	// example:
	//
	// Hologres
	DestinationDataSourceType *string `json:"DestinationDataSourceType,omitempty" xml:"DestinationDataSourceType,omitempty"`
	// The name of the synchronization task.
	//
	// example:
	//
	// mysql_to_holo_sync_8772
	JobName *string `json:"JobName,omitempty" xml:"JobName,omitempty"`
	// The settings for the dimension of the synchronization task. The settings include processing policies for DDL messages, policies for data type mappings between source fields and destination fields, and runtime parameters of the synchronization task.
	JobSettings *CreateDIJobRequestJobSettings `json:"JobSettings,omitempty" xml:"JobSettings,omitempty" type:"Struct"`
	// The synchronization type. Valid values:
	//
	// 	- FullAndRealtimeIncremental (one-time full synchronization and real-time incremental synchronization)
	//
	// 	- RealtimeIncremental (real-time incremental synchronization)
	//
	// 	- Full (full synchronization)
	//
	// 	- OfflineIncremental (batch incremental synchronization)
	//
	// 	- FullAndOfflineIncremental (one-time full synchronization and batch incremental synchronization)
	//
	// example:
	//
	// FullAndRealtimeIncremental
	MigrationType *string `json:"MigrationType,omitempty" xml:"MigrationType,omitempty"`
	// The DataWorks workspace ID. You can call the [ListProjects](https://help.aliyun.com/document_detail/178393.html) operation to obtain the ID.
	//
	// example:
	//
	// 10000
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The resource settings.
	ResourceSettings *CreateDIJobRequestResourceSettings `json:"ResourceSettings,omitempty" xml:"ResourceSettings,omitempty" type:"Struct"`
	// The settings of the source. Only a single source is supported.
	SourceDataSourceSettings []*CreateDIJobRequestSourceDataSourceSettings `json:"SourceDataSourceSettings,omitempty" xml:"SourceDataSourceSettings,omitempty" type:"Repeated"`
	// The source type. Set this parameter to MySQL.
	//
	// example:
	//
	// MySQL
	SourceDataSourceType *string `json:"SourceDataSourceType,omitempty" xml:"SourceDataSourceType,omitempty"`
	// Specifies whether to perform system debugging. Valid values: true and false. Default value: false.
	//
	// example:
	//
	// false
	SystemDebug *string `json:"SystemDebug,omitempty" xml:"SystemDebug,omitempty"`
	// The list of mappings between rules used to select synchronization objects in the source and transformation rules applied to the selected synchronization objects. Each entry in the list displays a mapping between a rule used to select synchronization objects and a transformation rule applied to the selected synchronization objects.
	TableMappings []*CreateDIJobRequestTableMappings `json:"TableMappings,omitempty" xml:"TableMappings,omitempty" type:"Repeated"`
	// The list of transformation rules that you want to apply to the synchronization objects selected from the source. Each entry in the list defines a transformation rule.
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

func (s *CreateDIJobRequest) GetJobName() *string {
	return s.JobName
}

func (s *CreateDIJobRequest) GetJobSettings() *CreateDIJobRequestJobSettings {
	return s.JobSettings
}

func (s *CreateDIJobRequest) GetMigrationType() *string {
	return s.MigrationType
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

func (s *CreateDIJobRequest) GetSystemDebug() *string {
	return s.SystemDebug
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

func (s *CreateDIJobRequest) SetJobName(v string) *CreateDIJobRequest {
	s.JobName = &v
	return s
}

func (s *CreateDIJobRequest) SetJobSettings(v *CreateDIJobRequestJobSettings) *CreateDIJobRequest {
	s.JobSettings = v
	return s
}

func (s *CreateDIJobRequest) SetMigrationType(v string) *CreateDIJobRequest {
	s.MigrationType = &v
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

func (s *CreateDIJobRequest) SetSystemDebug(v string) *CreateDIJobRequest {
	s.SystemDebug = &v
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
	return dara.Validate(s)
}

type CreateDIJobRequestDestinationDataSourceSettings struct {
	// The name of the data source.
	//
	// example:
	//
	// holo_datasource_1
	DataSourceName *string `json:"DataSourceName,omitempty" xml:"DataSourceName,omitempty"`
	// The properties of the data source.
	DataSourceProperties map[string]*string `json:"DataSourceProperties,omitempty" xml:"DataSourceProperties,omitempty"`
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

func (s *CreateDIJobRequestDestinationDataSourceSettings) GetDataSourceProperties() map[string]*string {
	return s.DataSourceProperties
}

func (s *CreateDIJobRequestDestinationDataSourceSettings) SetDataSourceName(v string) *CreateDIJobRequestDestinationDataSourceSettings {
	s.DataSourceName = &v
	return s
}

func (s *CreateDIJobRequestDestinationDataSourceSettings) SetDataSourceProperties(v map[string]*string) *CreateDIJobRequestDestinationDataSourceSettings {
	s.DataSourceProperties = v
	return s
}

func (s *CreateDIJobRequestDestinationDataSourceSettings) Validate() error {
	return dara.Validate(s)
}

type CreateDIJobRequestJobSettings struct {
	// The channel control settings for the synchronization task. The value of this parameter must be a JSON string.
	//
	// example:
	//
	// {"structInfo":"MANAGED","storageType":"TEXTFILE","writeMode":"APPEND","partitionColumns":[{"columnName":"pt","columnType":"STRING","comment":""}],"fieldDelimiter":""}
	ChannelSettings *string `json:"ChannelSettings,omitempty" xml:"ChannelSettings,omitempty"`
	// The settings for data type mappings between source fields and destination fields. The value of this parameter must be an array.
	ColumnDataTypeSettings []*CreateDIJobRequestJobSettingsColumnDataTypeSettings `json:"ColumnDataTypeSettings,omitempty" xml:"ColumnDataTypeSettings,omitempty" type:"Repeated"`
	// The settings for periodic scheduling.
	CycleScheduleSettings *CreateDIJobRequestJobSettingsCycleScheduleSettings `json:"CycleScheduleSettings,omitempty" xml:"CycleScheduleSettings,omitempty" type:"Struct"`
	// The processing settings for DDL messages.
	DdlHandlingSettings []*CreateDIJobRequestJobSettingsDdlHandlingSettings `json:"DdlHandlingSettings,omitempty" xml:"DdlHandlingSettings,omitempty" type:"Repeated"`
	// The import settings for the synchronization task.
	ImportRuleSettings *CreateDIJobRequestJobSettingsImportRuleSettings `json:"ImportRuleSettings,omitempty" xml:"ImportRuleSettings,omitempty" type:"Struct"`
	// The runtime settings. The value of this parameter must be an array.
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

func (s *CreateDIJobRequestJobSettings) GetImportRuleSettings() *CreateDIJobRequestJobSettingsImportRuleSettings {
	return s.ImportRuleSettings
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

func (s *CreateDIJobRequestJobSettings) SetImportRuleSettings(v *CreateDIJobRequestJobSettingsImportRuleSettings) *CreateDIJobRequestJobSettings {
	s.ImportRuleSettings = v
	return s
}

func (s *CreateDIJobRequestJobSettings) SetRuntimeSettings(v []*CreateDIJobRequestJobSettingsRuntimeSettings) *CreateDIJobRequestJobSettings {
	s.RuntimeSettings = v
	return s
}

func (s *CreateDIJobRequestJobSettings) Validate() error {
	return dara.Validate(s)
}

type CreateDIJobRequestJobSettingsColumnDataTypeSettings struct {
	// The data type of the destination field.
	//
	// example:
	//
	// text
	DestinationDataType *string `json:"DestinationDataType,omitempty" xml:"DestinationDataType,omitempty"`
	// The data type of the source field.
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
	// The processing policy. Valid values:
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

type CreateDIJobRequestJobSettingsImportRuleSettings struct {
	// The ID of the task to be imported.
	//
	// example:
	//
	// 10000
	FileId *string `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// The import source of the task. Set the value to Datastudio, which indicates synchronization tasks created in DataStudio.
	//
	// example:
	//
	// Datastudio
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
}

func (s CreateDIJobRequestJobSettingsImportRuleSettings) String() string {
	return dara.Prettify(s)
}

func (s CreateDIJobRequestJobSettingsImportRuleSettings) GoString() string {
	return s.String()
}

func (s *CreateDIJobRequestJobSettingsImportRuleSettings) GetFileId() *string {
	return s.FileId
}

func (s *CreateDIJobRequestJobSettingsImportRuleSettings) GetSource() *string {
	return s.Source
}

func (s *CreateDIJobRequestJobSettingsImportRuleSettings) SetFileId(v string) *CreateDIJobRequestJobSettingsImportRuleSettings {
	s.FileId = &v
	return s
}

func (s *CreateDIJobRequestJobSettingsImportRuleSettings) SetSource(v string) *CreateDIJobRequestJobSettingsImportRuleSettings {
	s.Source = &v
	return s
}

func (s *CreateDIJobRequestJobSettingsImportRuleSettings) Validate() error {
	return dara.Validate(s)
}

type CreateDIJobRequestJobSettingsRuntimeSettings struct {
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
	// The resource used for batch synchronization.
	OfflineResourceSettings *CreateDIJobRequestResourceSettingsOfflineResourceSettings `json:"OfflineResourceSettings,omitempty" xml:"OfflineResourceSettings,omitempty" type:"Struct"`
	// The resource used for real-time synchronization.
	RealtimeResourceSettings *CreateDIJobRequestResourceSettingsRealtimeResourceSettings `json:"RealtimeResourceSettings,omitempty" xml:"RealtimeResourceSettings,omitempty" type:"Struct"`
	// The number of compute units (CUs) in the resource group that are used for incremental and full synchronization.
	//
	// example:
	//
	// 2.0
	RequestedCu *float32 `json:"RequestedCu,omitempty" xml:"RequestedCu,omitempty"`
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

func (s *CreateDIJobRequestResourceSettings) GetRequestedCu() *float32 {
	return s.RequestedCu
}

func (s *CreateDIJobRequestResourceSettings) SetOfflineResourceSettings(v *CreateDIJobRequestResourceSettingsOfflineResourceSettings) *CreateDIJobRequestResourceSettings {
	s.OfflineResourceSettings = v
	return s
}

func (s *CreateDIJobRequestResourceSettings) SetRealtimeResourceSettings(v *CreateDIJobRequestResourceSettingsRealtimeResourceSettings) *CreateDIJobRequestResourceSettings {
	s.RealtimeResourceSettings = v
	return s
}

func (s *CreateDIJobRequestResourceSettings) SetRequestedCu(v float32) *CreateDIJobRequestResourceSettings {
	s.RequestedCu = &v
	return s
}

func (s *CreateDIJobRequestResourceSettings) Validate() error {
	return dara.Validate(s)
}

type CreateDIJobRequestResourceSettingsOfflineResourceSettings struct {
	// The identifier of the resource group for Data Integration used for batch synchronization.
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

func (s *CreateDIJobRequestResourceSettingsOfflineResourceSettings) GetResourceGroupIdentifier() *string {
	return s.ResourceGroupIdentifier
}

func (s *CreateDIJobRequestResourceSettingsOfflineResourceSettings) SetResourceGroupIdentifier(v string) *CreateDIJobRequestResourceSettingsOfflineResourceSettings {
	s.ResourceGroupIdentifier = &v
	return s
}

func (s *CreateDIJobRequestResourceSettingsOfflineResourceSettings) Validate() error {
	return dara.Validate(s)
}

type CreateDIJobRequestResourceSettingsRealtimeResourceSettings struct {
	// The identifier of the resource group for Data Integration used for real-time synchronization.
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

func (s *CreateDIJobRequestResourceSettingsRealtimeResourceSettings) GetResourceGroupIdentifier() *string {
	return s.ResourceGroupIdentifier
}

func (s *CreateDIJobRequestResourceSettingsRealtimeResourceSettings) SetResourceGroupIdentifier(v string) *CreateDIJobRequestResourceSettingsRealtimeResourceSettings {
	s.ResourceGroupIdentifier = &v
	return s
}

func (s *CreateDIJobRequestResourceSettingsRealtimeResourceSettings) Validate() error {
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
	DataSourceProperties map[string]*string `json:"DataSourceProperties,omitempty" xml:"DataSourceProperties,omitempty"`
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

func (s *CreateDIJobRequestSourceDataSourceSettings) GetDataSourceProperties() map[string]*string {
	return s.DataSourceProperties
}

func (s *CreateDIJobRequestSourceDataSourceSettings) SetDataSourceName(v string) *CreateDIJobRequestSourceDataSourceSettings {
	s.DataSourceName = &v
	return s
}

func (s *CreateDIJobRequestSourceDataSourceSettings) SetDataSourceProperties(v map[string]*string) *CreateDIJobRequestSourceDataSourceSettings {
	s.DataSourceProperties = v
	return s
}

func (s *CreateDIJobRequestSourceDataSourceSettings) Validate() error {
	return dara.Validate(s)
}

type CreateDIJobRequestTableMappings struct {
	// The rules used to select synchronization objects in the source.
	SourceObjectSelectionRules []*CreateDIJobRequestTableMappingsSourceObjectSelectionRules `json:"SourceObjectSelectionRules,omitempty" xml:"SourceObjectSelectionRules,omitempty" type:"Repeated"`
	// The list of transformation rules that you want to apply to the synchronization objects selected from the source.
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
	return dara.Validate(s)
}

type CreateDIJobRequestTableMappingsSourceObjectSelectionRules struct {
	// The expression.
	//
	// example:
	//
	// mysql_table_1
	Expression *string `json:"Expression,omitempty" xml:"Expression,omitempty"`
	// The object type. Valid values:
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

func (s CreateDIJobRequestTableMappingsSourceObjectSelectionRules) String() string {
	return dara.Prettify(s)
}

func (s CreateDIJobRequestTableMappingsSourceObjectSelectionRules) GoString() string {
	return s.String()
}

func (s *CreateDIJobRequestTableMappingsSourceObjectSelectionRules) GetExpression() *string {
	return s.Expression
}

func (s *CreateDIJobRequestTableMappingsSourceObjectSelectionRules) GetObjectType() *string {
	return s.ObjectType
}

func (s *CreateDIJobRequestTableMappingsSourceObjectSelectionRules) SetExpression(v string) *CreateDIJobRequestTableMappingsSourceObjectSelectionRules {
	s.Expression = &v
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
	// 	- DefineRuntimeSettings
	//
	// 	- DefinePartitionKey
	//
	// example:
	//
	// Rename
	RuleActionType *string `json:"RuleActionType,omitempty" xml:"RuleActionType,omitempty"`
	// The expression of the rule. An expression must be a JSON string.
	//
	// Example of a renaming rule: {"expression":"${srcDatasourceName}_${srcDatabaseName}_0922","variables":[{"variableName":"srcDatabaseName","variableRules":[{"from":"fromdb","to":"todb"}]}]}
	//
	// 	- expression: the expression of the renaming rule. You can use the following variables in an expression: ${srcDatasourceName}, ${srcDatabaseName}, and ${srcTableName}. ${srcDatasourceName} indicates the name of the source. ${srcDatabaseName} indicates the name of a source database. ${srcTableName} indicates the name of a source table.
	//
	// 	- variables: the generation rule for a variable used in the expression of the renaming rule. The default value of the specified variable is the original value of the object indicated by the variable. You can define a group of string replacement rules to change the original values based on your business requirements. variableName: the name of the variable. Do not enclose the variable name in ${}. variableRules: the string replacement rules for variables. The system runs the string replacement rules in sequence for string replacement. from specifies the original string. to specifies the new string.
	//
	// Example of a rule used to add a specific field to the destination and assign a value to the field: {"columns":[{"columnName":"my_add_column","columnValueType":"Constant","columnValue":"123"}]}
	//
	// 	- If you do not configure such a rule, no fields are added to the destination and no values are assigned by default.
	//
	// 	- columnName: the name of the field that is added.
	//
	// 	- columnValueType: the type of the value of the field. Valid values: Constant and Variable.
	//
	// 	- columnValue: the value of the field that you want to add. If you set the valueType parameter to Constant, set the columnValue parameter to a custom constant of the STRING type. If you set the valueType parameter to Variable, set the columnValue to a built-in variable. The following built-in variables are supported: EXECUTE_TIME (LONG data type), DB_NAME_SRC (STRING data type), DATASOURCE_NAME_SRC (STRING data type), TABLE_NAME_SRC (STRING data type), DB_NAME_DEST (STRING data type), DATASOURCE_NAME_DEST (STRING data type), TABLE_NAME_DEST (STRING data type), and DB_NAME_SRC_TRANSED (STRING data type). EXECUTE_TIME specifies the execution time. DB_NAME_SRC specifies the name of a source database. DATASOURCE_NAME_SRC specifies the name of the source. TABLE_NAME_SRC specifies the name of a source table. DB_NAME_DEST specifies the name of a destination database. DATASOURCE_NAME_DEST specifies the name of the destination. TABLE_NAME_DEST specifies the name of a destination table. DB_NAME_SRC_TRANSED specifies the database name obtained after a transformation.
	//
	// Example of a rule used to specify primary key fields for a destination table: {"columns":["ukcolumn1","ukcolumn2"]}
	//
	// 	- If you do not configure such a rule, the primary key fields in the mapped source table are used for the destination table by default.
	//
	// 	- If the destination table is an existing table, Data Integration does not modify the schema of the destination table. If the specified primary key fields do not exist in the destination table, an error is reported when the synchronization task starts to run.
	//
	// 	- If the destination table is automatically created by the system, Data Integration automatically creates the schema of the destination table. The schema contains the primary key fields that you specify. If the specified primary key fields do not exist in the destination table, an error is reported when the synchronization task starts to run.
	//
	// Example of a rule used to process DML messages: {"dmlPolicies":[{"dmlType":"Delete","dmlAction":"Filter","filterCondition":"id > 1"}]}
	//
	// 	- If you do not configure such a rule, the default processing policy for messages generated for insert, update, and delete operations is Normal.
	//
	// 	- dmlType: the DML operation. Valid values: Insert, Update, and Delete.
	//
	// 	- dmlAction: the processing policy for DML messages. Valid values: Normal, Ignore, Filter, and LogicalDelete. Filter indicates conditional processing. You can set the dmlAction parameter to Filter only when the dmlType parameter is set to Update or Delete.
	//
	// 	- filterCondition: the condition used to filter DML messages. This parameter is required only when the dmlAction parameter is set to Filter.
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
