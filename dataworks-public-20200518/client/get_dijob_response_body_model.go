// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDIJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetDIJobResponseBodyData) *GetDIJobResponseBody
	GetData() *GetDIJobResponseBodyData
	SetRequestId(v string) *GetDIJobResponseBody
	GetRequestId() *string
}

type GetDIJobResponseBody struct {
	// The information about the synchronization task.
	Data *GetDIJobResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 0000-ABCD-EFG****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetDIJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDIJobResponseBody) GoString() string {
	return s.String()
}

func (s *GetDIJobResponseBody) GetData() *GetDIJobResponseBodyData {
	return s.Data
}

func (s *GetDIJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDIJobResponseBody) SetData(v *GetDIJobResponseBodyData) *GetDIJobResponseBody {
	s.Data = v
	return s
}

func (s *GetDIJobResponseBody) SetRequestId(v string) *GetDIJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDIJobResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetDIJobResponseBodyData struct {
	// The timestamp when the synchronization task was created. The timestamp is accurate to the second.
	//
	// example:
	//
	// 1671516776
	CreatedTime *int64 `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// The ID of the user who creates the synchronization task.
	//
	// example:
	//
	// 100000001
	CreatedUid *string `json:"CreatedUid,omitempty" xml:"CreatedUid,omitempty"`
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
	// The settings of the destination. Only a single destination is supported.
	DestinationDataSourceSettings []*GetDIJobResponseBodyDataDestinationDataSourceSettings `json:"DestinationDataSourceSettings,omitempty" xml:"DestinationDataSourceSettings,omitempty" type:"Repeated"`
	// The destination type. Valid values: Hologres and Hive.
	//
	// example:
	//
	// Hologres
	DestinationDataSourceType *string `json:"DestinationDataSourceType,omitempty" xml:"DestinationDataSourceType,omitempty"`
	// The error message returned if the value of the JobStatus parameter is Failed.
	//
	// example:
	//
	// error details xxx
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The name of the synchronization task.
	//
	// example:
	//
	// mysql_to_holo_sync_445
	JobName *string `json:"JobName,omitempty" xml:"JobName,omitempty"`
	// The settings for the dimension of the synchronization task. The settings include processing policies for DDL messages, policies for data type mappings between source fields and destination fields, and runtime parameters of the synchronization task.
	JobSettings *GetDIJobResponseBodyDataJobSettings `json:"JobSettings,omitempty" xml:"JobSettings,omitempty" type:"Struct"`
	// The status of the synchronization task. Valid values:
	//
	// 	- Finished
	//
	// 	- Initialized
	//
	// 	- Stopped
	//
	// 	- Failed
	//
	// 	- Running
	//
	// 	- Stopping
	//
	// example:
	//
	// Finished
	JobStatus *string `json:"JobStatus,omitempty" xml:"JobStatus,omitempty"`
	// The synchronization type. Valid values:
	//
	// 	- FullAndRealtimeIncremental: one-time full synchronization and real-time incremental synchronization
	//
	// 	- RealtimeIncremental: real-time incremental synchronization
	//
	// 	- Full: full synchronization
	//
	// 	- OfflineIncremental: batch incremental synchronization
	//
	// 	- FullAndOfflineIncremental: one-time full synchronization and batch incremental synchronization
	//
	// example:
	//
	// FullAndRealtimeIncremental
	MigrationType *string `json:"MigrationType,omitempty" xml:"MigrationType,omitempty"`
	// The workspace ID.
	//
	// example:
	//
	// 22
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The resource settings.
	ResourceSettings *GetDIJobResponseBodyDataResourceSettings `json:"ResourceSettings,omitempty" xml:"ResourceSettings,omitempty" type:"Struct"`
	// The information about the running of the synchronization task.
	RunStats map[string]*string `json:"RunStats,omitempty" xml:"RunStats,omitempty"`
	// The settings of the source. Only a single source is supported.
	SourceDataSourceSettings []*GetDIJobResponseBodyDataSourceDataSourceSettings `json:"SourceDataSourceSettings,omitempty" xml:"SourceDataSourceSettings,omitempty" type:"Repeated"`
	// The source type. The value MySQL is returned.
	//
	// example:
	//
	// MySQL
	SourceDataSourceType *string `json:"SourceDataSourceType,omitempty" xml:"SourceDataSourceType,omitempty"`
	// The timestamp when the synchronization task was last started. The timestamp is accurate to the second.
	//
	// example:
	//
	// 1673859999
	StartedTime *int64 `json:"StartedTime,omitempty" xml:"StartedTime,omitempty"`
	// The ID of the user who last starts the synchronization task.
	//
	// example:
	//
	// 100000001
	StartedUid *string `json:"StartedUid,omitempty" xml:"StartedUid,omitempty"`
	// The list of mappings between rules used to select synchronization objects in the source and transformation rules applied to the selected synchronization objects. Each entry in the list displays a mapping between a rule used to select synchronization objects and a transformation rule applied to the selected synchronization objects.
	TableMappings []*GetDIJobResponseBodyDataTableMappings `json:"TableMappings,omitempty" xml:"TableMappings,omitempty" type:"Repeated"`
	// The list of transformation rules that are applied to the synchronization objects selected from the source. Each entry in the list defines a transformation rule.
	TransformationRules []*GetDIJobResponseBodyDataTransformationRules `json:"TransformationRules,omitempty" xml:"TransformationRules,omitempty" type:"Repeated"`
	// The timestamp when the synchronization task was last modified. The timestamp is accurate to the second.
	//
	// example:
	//
	// 1673859985
	UpdatedTime *int64 `json:"UpdatedTime,omitempty" xml:"UpdatedTime,omitempty"`
	// The ID of the user who last modifies the synchronization task.
	//
	// example:
	//
	// 100000001
	UpdatedUid *string `json:"UpdatedUid,omitempty" xml:"UpdatedUid,omitempty"`
}

func (s GetDIJobResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetDIJobResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetDIJobResponseBodyData) GetCreatedTime() *int64 {
	return s.CreatedTime
}

func (s *GetDIJobResponseBodyData) GetCreatedUid() *string {
	return s.CreatedUid
}

func (s *GetDIJobResponseBodyData) GetDIJobId() *int64 {
	return s.DIJobId
}

func (s *GetDIJobResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *GetDIJobResponseBodyData) GetDestinationDataSourceSettings() []*GetDIJobResponseBodyDataDestinationDataSourceSettings {
	return s.DestinationDataSourceSettings
}

func (s *GetDIJobResponseBodyData) GetDestinationDataSourceType() *string {
	return s.DestinationDataSourceType
}

func (s *GetDIJobResponseBodyData) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetDIJobResponseBodyData) GetJobName() *string {
	return s.JobName
}

func (s *GetDIJobResponseBodyData) GetJobSettings() *GetDIJobResponseBodyDataJobSettings {
	return s.JobSettings
}

func (s *GetDIJobResponseBodyData) GetJobStatus() *string {
	return s.JobStatus
}

func (s *GetDIJobResponseBodyData) GetMigrationType() *string {
	return s.MigrationType
}

func (s *GetDIJobResponseBodyData) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetDIJobResponseBodyData) GetResourceSettings() *GetDIJobResponseBodyDataResourceSettings {
	return s.ResourceSettings
}

func (s *GetDIJobResponseBodyData) GetRunStats() map[string]*string {
	return s.RunStats
}

func (s *GetDIJobResponseBodyData) GetSourceDataSourceSettings() []*GetDIJobResponseBodyDataSourceDataSourceSettings {
	return s.SourceDataSourceSettings
}

func (s *GetDIJobResponseBodyData) GetSourceDataSourceType() *string {
	return s.SourceDataSourceType
}

func (s *GetDIJobResponseBodyData) GetStartedTime() *int64 {
	return s.StartedTime
}

func (s *GetDIJobResponseBodyData) GetStartedUid() *string {
	return s.StartedUid
}

func (s *GetDIJobResponseBodyData) GetTableMappings() []*GetDIJobResponseBodyDataTableMappings {
	return s.TableMappings
}

func (s *GetDIJobResponseBodyData) GetTransformationRules() []*GetDIJobResponseBodyDataTransformationRules {
	return s.TransformationRules
}

func (s *GetDIJobResponseBodyData) GetUpdatedTime() *int64 {
	return s.UpdatedTime
}

func (s *GetDIJobResponseBodyData) GetUpdatedUid() *string {
	return s.UpdatedUid
}

func (s *GetDIJobResponseBodyData) SetCreatedTime(v int64) *GetDIJobResponseBodyData {
	s.CreatedTime = &v
	return s
}

func (s *GetDIJobResponseBodyData) SetCreatedUid(v string) *GetDIJobResponseBodyData {
	s.CreatedUid = &v
	return s
}

func (s *GetDIJobResponseBodyData) SetDIJobId(v int64) *GetDIJobResponseBodyData {
	s.DIJobId = &v
	return s
}

func (s *GetDIJobResponseBodyData) SetDescription(v string) *GetDIJobResponseBodyData {
	s.Description = &v
	return s
}

func (s *GetDIJobResponseBodyData) SetDestinationDataSourceSettings(v []*GetDIJobResponseBodyDataDestinationDataSourceSettings) *GetDIJobResponseBodyData {
	s.DestinationDataSourceSettings = v
	return s
}

func (s *GetDIJobResponseBodyData) SetDestinationDataSourceType(v string) *GetDIJobResponseBodyData {
	s.DestinationDataSourceType = &v
	return s
}

func (s *GetDIJobResponseBodyData) SetErrorMessage(v string) *GetDIJobResponseBodyData {
	s.ErrorMessage = &v
	return s
}

func (s *GetDIJobResponseBodyData) SetJobName(v string) *GetDIJobResponseBodyData {
	s.JobName = &v
	return s
}

func (s *GetDIJobResponseBodyData) SetJobSettings(v *GetDIJobResponseBodyDataJobSettings) *GetDIJobResponseBodyData {
	s.JobSettings = v
	return s
}

func (s *GetDIJobResponseBodyData) SetJobStatus(v string) *GetDIJobResponseBodyData {
	s.JobStatus = &v
	return s
}

func (s *GetDIJobResponseBodyData) SetMigrationType(v string) *GetDIJobResponseBodyData {
	s.MigrationType = &v
	return s
}

func (s *GetDIJobResponseBodyData) SetProjectId(v int64) *GetDIJobResponseBodyData {
	s.ProjectId = &v
	return s
}

func (s *GetDIJobResponseBodyData) SetResourceSettings(v *GetDIJobResponseBodyDataResourceSettings) *GetDIJobResponseBodyData {
	s.ResourceSettings = v
	return s
}

func (s *GetDIJobResponseBodyData) SetRunStats(v map[string]*string) *GetDIJobResponseBodyData {
	s.RunStats = v
	return s
}

func (s *GetDIJobResponseBodyData) SetSourceDataSourceSettings(v []*GetDIJobResponseBodyDataSourceDataSourceSettings) *GetDIJobResponseBodyData {
	s.SourceDataSourceSettings = v
	return s
}

func (s *GetDIJobResponseBodyData) SetSourceDataSourceType(v string) *GetDIJobResponseBodyData {
	s.SourceDataSourceType = &v
	return s
}

func (s *GetDIJobResponseBodyData) SetStartedTime(v int64) *GetDIJobResponseBodyData {
	s.StartedTime = &v
	return s
}

func (s *GetDIJobResponseBodyData) SetStartedUid(v string) *GetDIJobResponseBodyData {
	s.StartedUid = &v
	return s
}

func (s *GetDIJobResponseBodyData) SetTableMappings(v []*GetDIJobResponseBodyDataTableMappings) *GetDIJobResponseBodyData {
	s.TableMappings = v
	return s
}

func (s *GetDIJobResponseBodyData) SetTransformationRules(v []*GetDIJobResponseBodyDataTransformationRules) *GetDIJobResponseBodyData {
	s.TransformationRules = v
	return s
}

func (s *GetDIJobResponseBodyData) SetUpdatedTime(v int64) *GetDIJobResponseBodyData {
	s.UpdatedTime = &v
	return s
}

func (s *GetDIJobResponseBodyData) SetUpdatedUid(v string) *GetDIJobResponseBodyData {
	s.UpdatedUid = &v
	return s
}

func (s *GetDIJobResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type GetDIJobResponseBodyDataDestinationDataSourceSettings struct {
	// The name of the data source.
	//
	// example:
	//
	// holo_datasource_1
	DataSourceName *string `json:"DataSourceName,omitempty" xml:"DataSourceName,omitempty"`
	// The properties of the data source.
	DataSourceProperties map[string]*string `json:"DataSourceProperties,omitempty" xml:"DataSourceProperties,omitempty"`
}

func (s GetDIJobResponseBodyDataDestinationDataSourceSettings) String() string {
	return dara.Prettify(s)
}

func (s GetDIJobResponseBodyDataDestinationDataSourceSettings) GoString() string {
	return s.String()
}

func (s *GetDIJobResponseBodyDataDestinationDataSourceSettings) GetDataSourceName() *string {
	return s.DataSourceName
}

func (s *GetDIJobResponseBodyDataDestinationDataSourceSettings) GetDataSourceProperties() map[string]*string {
	return s.DataSourceProperties
}

func (s *GetDIJobResponseBodyDataDestinationDataSourceSettings) SetDataSourceName(v string) *GetDIJobResponseBodyDataDestinationDataSourceSettings {
	s.DataSourceName = &v
	return s
}

func (s *GetDIJobResponseBodyDataDestinationDataSourceSettings) SetDataSourceProperties(v map[string]*string) *GetDIJobResponseBodyDataDestinationDataSourceSettings {
	s.DataSourceProperties = v
	return s
}

func (s *GetDIJobResponseBodyDataDestinationDataSourceSettings) Validate() error {
	return dara.Validate(s)
}

type GetDIJobResponseBodyDataJobSettings struct {
	// The channel control settings for the synchronization task. The value of this parameter is a JSON string.
	//
	// example:
	//
	// {"structInfo":"MANAGED","storageType":"TEXTFILE","writeMode":"APPEND","partitionColumns":[{"columnName":"pt","columnType":"STRING","comment":""}],"fieldDelimiter":""}
	ChannelSettings *string `json:"ChannelSettings,omitempty" xml:"ChannelSettings,omitempty"`
	// The settings for data type mappings between source fields and destination fields. The value of this parameter is an array.
	ColumnDataTypeSettings []*GetDIJobResponseBodyDataJobSettingsColumnDataTypeSettings `json:"ColumnDataTypeSettings,omitempty" xml:"ColumnDataTypeSettings,omitempty" type:"Repeated"`
	// The settings for periodic scheduling.
	CycleScheduleSettings *GetDIJobResponseBodyDataJobSettingsCycleScheduleSettings `json:"CycleScheduleSettings,omitempty" xml:"CycleScheduleSettings,omitempty" type:"Struct"`
	// The settings for processing DDL messages. The value of this parameter is an array.
	DdlHandlingSettings []*GetDIJobResponseBodyDataJobSettingsDdlHandlingSettings `json:"DdlHandlingSettings,omitempty" xml:"DdlHandlingSettings,omitempty" type:"Repeated"`
	// The runtime settings. The value of this parameter is an array.
	RuntimeSettings []*GetDIJobResponseBodyDataJobSettingsRuntimeSettings `json:"RuntimeSettings,omitempty" xml:"RuntimeSettings,omitempty" type:"Repeated"`
}

func (s GetDIJobResponseBodyDataJobSettings) String() string {
	return dara.Prettify(s)
}

func (s GetDIJobResponseBodyDataJobSettings) GoString() string {
	return s.String()
}

func (s *GetDIJobResponseBodyDataJobSettings) GetChannelSettings() *string {
	return s.ChannelSettings
}

func (s *GetDIJobResponseBodyDataJobSettings) GetColumnDataTypeSettings() []*GetDIJobResponseBodyDataJobSettingsColumnDataTypeSettings {
	return s.ColumnDataTypeSettings
}

func (s *GetDIJobResponseBodyDataJobSettings) GetCycleScheduleSettings() *GetDIJobResponseBodyDataJobSettingsCycleScheduleSettings {
	return s.CycleScheduleSettings
}

func (s *GetDIJobResponseBodyDataJobSettings) GetDdlHandlingSettings() []*GetDIJobResponseBodyDataJobSettingsDdlHandlingSettings {
	return s.DdlHandlingSettings
}

func (s *GetDIJobResponseBodyDataJobSettings) GetRuntimeSettings() []*GetDIJobResponseBodyDataJobSettingsRuntimeSettings {
	return s.RuntimeSettings
}

func (s *GetDIJobResponseBodyDataJobSettings) SetChannelSettings(v string) *GetDIJobResponseBodyDataJobSettings {
	s.ChannelSettings = &v
	return s
}

func (s *GetDIJobResponseBodyDataJobSettings) SetColumnDataTypeSettings(v []*GetDIJobResponseBodyDataJobSettingsColumnDataTypeSettings) *GetDIJobResponseBodyDataJobSettings {
	s.ColumnDataTypeSettings = v
	return s
}

func (s *GetDIJobResponseBodyDataJobSettings) SetCycleScheduleSettings(v *GetDIJobResponseBodyDataJobSettingsCycleScheduleSettings) *GetDIJobResponseBodyDataJobSettings {
	s.CycleScheduleSettings = v
	return s
}

func (s *GetDIJobResponseBodyDataJobSettings) SetDdlHandlingSettings(v []*GetDIJobResponseBodyDataJobSettingsDdlHandlingSettings) *GetDIJobResponseBodyDataJobSettings {
	s.DdlHandlingSettings = v
	return s
}

func (s *GetDIJobResponseBodyDataJobSettings) SetRuntimeSettings(v []*GetDIJobResponseBodyDataJobSettingsRuntimeSettings) *GetDIJobResponseBodyDataJobSettings {
	s.RuntimeSettings = v
	return s
}

func (s *GetDIJobResponseBodyDataJobSettings) Validate() error {
	return dara.Validate(s)
}

type GetDIJobResponseBodyDataJobSettingsColumnDataTypeSettings struct {
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

func (s GetDIJobResponseBodyDataJobSettingsColumnDataTypeSettings) String() string {
	return dara.Prettify(s)
}

func (s GetDIJobResponseBodyDataJobSettingsColumnDataTypeSettings) GoString() string {
	return s.String()
}

func (s *GetDIJobResponseBodyDataJobSettingsColumnDataTypeSettings) GetDestinationDataType() *string {
	return s.DestinationDataType
}

func (s *GetDIJobResponseBodyDataJobSettingsColumnDataTypeSettings) GetSourceDataType() *string {
	return s.SourceDataType
}

func (s *GetDIJobResponseBodyDataJobSettingsColumnDataTypeSettings) SetDestinationDataType(v string) *GetDIJobResponseBodyDataJobSettingsColumnDataTypeSettings {
	s.DestinationDataType = &v
	return s
}

func (s *GetDIJobResponseBodyDataJobSettingsColumnDataTypeSettings) SetSourceDataType(v string) *GetDIJobResponseBodyDataJobSettingsColumnDataTypeSettings {
	s.SourceDataType = &v
	return s
}

func (s *GetDIJobResponseBodyDataJobSettingsColumnDataTypeSettings) Validate() error {
	return dara.Validate(s)
}

type GetDIJobResponseBodyDataJobSettingsCycleScheduleSettings struct {
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

func (s GetDIJobResponseBodyDataJobSettingsCycleScheduleSettings) String() string {
	return dara.Prettify(s)
}

func (s GetDIJobResponseBodyDataJobSettingsCycleScheduleSettings) GoString() string {
	return s.String()
}

func (s *GetDIJobResponseBodyDataJobSettingsCycleScheduleSettings) GetCycleMigrationType() *string {
	return s.CycleMigrationType
}

func (s *GetDIJobResponseBodyDataJobSettingsCycleScheduleSettings) GetScheduleParameters() *string {
	return s.ScheduleParameters
}

func (s *GetDIJobResponseBodyDataJobSettingsCycleScheduleSettings) SetCycleMigrationType(v string) *GetDIJobResponseBodyDataJobSettingsCycleScheduleSettings {
	s.CycleMigrationType = &v
	return s
}

func (s *GetDIJobResponseBodyDataJobSettingsCycleScheduleSettings) SetScheduleParameters(v string) *GetDIJobResponseBodyDataJobSettingsCycleScheduleSettings {
	s.ScheduleParameters = &v
	return s
}

func (s *GetDIJobResponseBodyDataJobSettingsCycleScheduleSettings) Validate() error {
	return dara.Validate(s)
}

type GetDIJobResponseBodyDataJobSettingsDdlHandlingSettings struct {
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

func (s GetDIJobResponseBodyDataJobSettingsDdlHandlingSettings) String() string {
	return dara.Prettify(s)
}

func (s GetDIJobResponseBodyDataJobSettingsDdlHandlingSettings) GoString() string {
	return s.String()
}

func (s *GetDIJobResponseBodyDataJobSettingsDdlHandlingSettings) GetAction() *string {
	return s.Action
}

func (s *GetDIJobResponseBodyDataJobSettingsDdlHandlingSettings) GetType() *string {
	return s.Type
}

func (s *GetDIJobResponseBodyDataJobSettingsDdlHandlingSettings) SetAction(v string) *GetDIJobResponseBodyDataJobSettingsDdlHandlingSettings {
	s.Action = &v
	return s
}

func (s *GetDIJobResponseBodyDataJobSettingsDdlHandlingSettings) SetType(v string) *GetDIJobResponseBodyDataJobSettingsDdlHandlingSettings {
	s.Type = &v
	return s
}

func (s *GetDIJobResponseBodyDataJobSettingsDdlHandlingSettings) Validate() error {
	return dara.Validate(s)
}

type GetDIJobResponseBodyDataJobSettingsRuntimeSettings struct {
	// The name of the configuration item. Valid values:
	//
	// 	- runtime.offline.speed.limit.mb: indicates the maximum transmission rate that is allowed for a batch synchronization task. This configuration item takes effect only when runtime.offline.speed.limit.enable is set to true.
	//
	// 	- runtime.offline.speed.limit.enable: indicates whether throttling is enabled for a batch synchronization task.
	//
	// 	- dst.offline.connection.max: indicates the maximum number of connections that are allowed for writing data to the destination of a batch synchronization task.
	//
	// 	- runtime.offline.concurrent: indicates the maximum number of parallel threads that are allowed for a batch synchronization task.
	//
	// 	- dst.realtime.connection.max: indicates the maximum number of connections that are allowed for writing data to the destination of a real-time synchronization task.
	//
	// 	- runtime.enable.auto.create.schema: indicates whether schemas are automatically created in the destination of a synchronization task.
	//
	// 	- src.offline.datasource.max.connection: indicates the maximum number of connections that are allowed for reading data from the source of a batch synchronization task.
	//
	// 	- runtime.realtime.concurrent: indicates the maximum number of parallel threads that are allowed for a real-time synchronization task.
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

func (s GetDIJobResponseBodyDataJobSettingsRuntimeSettings) String() string {
	return dara.Prettify(s)
}

func (s GetDIJobResponseBodyDataJobSettingsRuntimeSettings) GoString() string {
	return s.String()
}

func (s *GetDIJobResponseBodyDataJobSettingsRuntimeSettings) GetName() *string {
	return s.Name
}

func (s *GetDIJobResponseBodyDataJobSettingsRuntimeSettings) GetValue() *string {
	return s.Value
}

func (s *GetDIJobResponseBodyDataJobSettingsRuntimeSettings) SetName(v string) *GetDIJobResponseBodyDataJobSettingsRuntimeSettings {
	s.Name = &v
	return s
}

func (s *GetDIJobResponseBodyDataJobSettingsRuntimeSettings) SetValue(v string) *GetDIJobResponseBodyDataJobSettingsRuntimeSettings {
	s.Value = &v
	return s
}

func (s *GetDIJobResponseBodyDataJobSettingsRuntimeSettings) Validate() error {
	return dara.Validate(s)
}

type GetDIJobResponseBodyDataResourceSettings struct {
	// The resource used for batch synchronization.
	OfflineResourceSettings *GetDIJobResponseBodyDataResourceSettingsOfflineResourceSettings `json:"OfflineResourceSettings,omitempty" xml:"OfflineResourceSettings,omitempty" type:"Struct"`
	// The resource used for real-time synchronization.
	RealtimeResourceSettings *GetDIJobResponseBodyDataResourceSettingsRealtimeResourceSettings `json:"RealtimeResourceSettings,omitempty" xml:"RealtimeResourceSettings,omitempty" type:"Struct"`
	// The number of compute units (CUs) in the resource group that are used for incremental and full synchronization.
	//
	// example:
	//
	// 2.0
	RequestedCu *float32 `json:"RequestedCu,omitempty" xml:"RequestedCu,omitempty"`
}

func (s GetDIJobResponseBodyDataResourceSettings) String() string {
	return dara.Prettify(s)
}

func (s GetDIJobResponseBodyDataResourceSettings) GoString() string {
	return s.String()
}

func (s *GetDIJobResponseBodyDataResourceSettings) GetOfflineResourceSettings() *GetDIJobResponseBodyDataResourceSettingsOfflineResourceSettings {
	return s.OfflineResourceSettings
}

func (s *GetDIJobResponseBodyDataResourceSettings) GetRealtimeResourceSettings() *GetDIJobResponseBodyDataResourceSettingsRealtimeResourceSettings {
	return s.RealtimeResourceSettings
}

func (s *GetDIJobResponseBodyDataResourceSettings) GetRequestedCu() *float32 {
	return s.RequestedCu
}

func (s *GetDIJobResponseBodyDataResourceSettings) SetOfflineResourceSettings(v *GetDIJobResponseBodyDataResourceSettingsOfflineResourceSettings) *GetDIJobResponseBodyDataResourceSettings {
	s.OfflineResourceSettings = v
	return s
}

func (s *GetDIJobResponseBodyDataResourceSettings) SetRealtimeResourceSettings(v *GetDIJobResponseBodyDataResourceSettingsRealtimeResourceSettings) *GetDIJobResponseBodyDataResourceSettings {
	s.RealtimeResourceSettings = v
	return s
}

func (s *GetDIJobResponseBodyDataResourceSettings) SetRequestedCu(v float32) *GetDIJobResponseBodyDataResourceSettings {
	s.RequestedCu = &v
	return s
}

func (s *GetDIJobResponseBodyDataResourceSettings) Validate() error {
	return dara.Validate(s)
}

type GetDIJobResponseBodyDataResourceSettingsOfflineResourceSettings struct {
	// The identifier of the resource group for Data Integration used for batch synchronization.
	//
	// example:
	//
	// S_res_group_111_222
	ResourceGroupIdentifier *string `json:"ResourceGroupIdentifier,omitempty" xml:"ResourceGroupIdentifier,omitempty"`
}

func (s GetDIJobResponseBodyDataResourceSettingsOfflineResourceSettings) String() string {
	return dara.Prettify(s)
}

func (s GetDIJobResponseBodyDataResourceSettingsOfflineResourceSettings) GoString() string {
	return s.String()
}

func (s *GetDIJobResponseBodyDataResourceSettingsOfflineResourceSettings) GetResourceGroupIdentifier() *string {
	return s.ResourceGroupIdentifier
}

func (s *GetDIJobResponseBodyDataResourceSettingsOfflineResourceSettings) SetResourceGroupIdentifier(v string) *GetDIJobResponseBodyDataResourceSettingsOfflineResourceSettings {
	s.ResourceGroupIdentifier = &v
	return s
}

func (s *GetDIJobResponseBodyDataResourceSettingsOfflineResourceSettings) Validate() error {
	return dara.Validate(s)
}

type GetDIJobResponseBodyDataResourceSettingsRealtimeResourceSettings struct {
	// The identifier of the resource group for Data Integration used for real-time synchronization.
	//
	// example:
	//
	// S_res_group_111_222
	ResourceGroupIdentifier *string `json:"ResourceGroupIdentifier,omitempty" xml:"ResourceGroupIdentifier,omitempty"`
}

func (s GetDIJobResponseBodyDataResourceSettingsRealtimeResourceSettings) String() string {
	return dara.Prettify(s)
}

func (s GetDIJobResponseBodyDataResourceSettingsRealtimeResourceSettings) GoString() string {
	return s.String()
}

func (s *GetDIJobResponseBodyDataResourceSettingsRealtimeResourceSettings) GetResourceGroupIdentifier() *string {
	return s.ResourceGroupIdentifier
}

func (s *GetDIJobResponseBodyDataResourceSettingsRealtimeResourceSettings) SetResourceGroupIdentifier(v string) *GetDIJobResponseBodyDataResourceSettingsRealtimeResourceSettings {
	s.ResourceGroupIdentifier = &v
	return s
}

func (s *GetDIJobResponseBodyDataResourceSettingsRealtimeResourceSettings) Validate() error {
	return dara.Validate(s)
}

type GetDIJobResponseBodyDataSourceDataSourceSettings struct {
	// The name of the data source.
	//
	// example:
	//
	// mysql_datasource_1
	DataSourceName *string `json:"DataSourceName,omitempty" xml:"DataSourceName,omitempty"`
	// The properties of the data source.
	DataSourceProperties map[string]*string `json:"DataSourceProperties,omitempty" xml:"DataSourceProperties,omitempty"`
}

func (s GetDIJobResponseBodyDataSourceDataSourceSettings) String() string {
	return dara.Prettify(s)
}

func (s GetDIJobResponseBodyDataSourceDataSourceSettings) GoString() string {
	return s.String()
}

func (s *GetDIJobResponseBodyDataSourceDataSourceSettings) GetDataSourceName() *string {
	return s.DataSourceName
}

func (s *GetDIJobResponseBodyDataSourceDataSourceSettings) GetDataSourceProperties() map[string]*string {
	return s.DataSourceProperties
}

func (s *GetDIJobResponseBodyDataSourceDataSourceSettings) SetDataSourceName(v string) *GetDIJobResponseBodyDataSourceDataSourceSettings {
	s.DataSourceName = &v
	return s
}

func (s *GetDIJobResponseBodyDataSourceDataSourceSettings) SetDataSourceProperties(v map[string]*string) *GetDIJobResponseBodyDataSourceDataSourceSettings {
	s.DataSourceProperties = v
	return s
}

func (s *GetDIJobResponseBodyDataSourceDataSourceSettings) Validate() error {
	return dara.Validate(s)
}

type GetDIJobResponseBodyDataTableMappings struct {
	// The list of rules used to select synchronization objects in the source.
	SourceObjectSelectionRules []*GetDIJobResponseBodyDataTableMappingsSourceObjectSelectionRules `json:"SourceObjectSelectionRules,omitempty" xml:"SourceObjectSelectionRules,omitempty" type:"Repeated"`
	// The list of transformation rules that are applied to the synchronization objects selected from the source.
	TransformationRules []*GetDIJobResponseBodyDataTableMappingsTransformationRules `json:"TransformationRules,omitempty" xml:"TransformationRules,omitempty" type:"Repeated"`
}

func (s GetDIJobResponseBodyDataTableMappings) String() string {
	return dara.Prettify(s)
}

func (s GetDIJobResponseBodyDataTableMappings) GoString() string {
	return s.String()
}

func (s *GetDIJobResponseBodyDataTableMappings) GetSourceObjectSelectionRules() []*GetDIJobResponseBodyDataTableMappingsSourceObjectSelectionRules {
	return s.SourceObjectSelectionRules
}

func (s *GetDIJobResponseBodyDataTableMappings) GetTransformationRules() []*GetDIJobResponseBodyDataTableMappingsTransformationRules {
	return s.TransformationRules
}

func (s *GetDIJobResponseBodyDataTableMappings) SetSourceObjectSelectionRules(v []*GetDIJobResponseBodyDataTableMappingsSourceObjectSelectionRules) *GetDIJobResponseBodyDataTableMappings {
	s.SourceObjectSelectionRules = v
	return s
}

func (s *GetDIJobResponseBodyDataTableMappings) SetTransformationRules(v []*GetDIJobResponseBodyDataTableMappingsTransformationRules) *GetDIJobResponseBodyDataTableMappings {
	s.TransformationRules = v
	return s
}

func (s *GetDIJobResponseBodyDataTableMappings) Validate() error {
	return dara.Validate(s)
}

type GetDIJobResponseBodyDataTableMappingsSourceObjectSelectionRules struct {
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

func (s GetDIJobResponseBodyDataTableMappingsSourceObjectSelectionRules) String() string {
	return dara.Prettify(s)
}

func (s GetDIJobResponseBodyDataTableMappingsSourceObjectSelectionRules) GoString() string {
	return s.String()
}

func (s *GetDIJobResponseBodyDataTableMappingsSourceObjectSelectionRules) GetExpression() *string {
	return s.Expression
}

func (s *GetDIJobResponseBodyDataTableMappingsSourceObjectSelectionRules) GetObjectType() *string {
	return s.ObjectType
}

func (s *GetDIJobResponseBodyDataTableMappingsSourceObjectSelectionRules) SetExpression(v string) *GetDIJobResponseBodyDataTableMappingsSourceObjectSelectionRules {
	s.Expression = &v
	return s
}

func (s *GetDIJobResponseBodyDataTableMappingsSourceObjectSelectionRules) SetObjectType(v string) *GetDIJobResponseBodyDataTableMappingsSourceObjectSelectionRules {
	s.ObjectType = &v
	return s
}

func (s *GetDIJobResponseBodyDataTableMappingsSourceObjectSelectionRules) Validate() error {
	return dara.Validate(s)
}

type GetDIJobResponseBodyDataTableMappingsTransformationRules struct {
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
	// The type of the object on which the action is performed. Valid values:
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

func (s GetDIJobResponseBodyDataTableMappingsTransformationRules) String() string {
	return dara.Prettify(s)
}

func (s GetDIJobResponseBodyDataTableMappingsTransformationRules) GoString() string {
	return s.String()
}

func (s *GetDIJobResponseBodyDataTableMappingsTransformationRules) GetRuleActionType() *string {
	return s.RuleActionType
}

func (s *GetDIJobResponseBodyDataTableMappingsTransformationRules) GetRuleName() *string {
	return s.RuleName
}

func (s *GetDIJobResponseBodyDataTableMappingsTransformationRules) GetRuleTargetType() *string {
	return s.RuleTargetType
}

func (s *GetDIJobResponseBodyDataTableMappingsTransformationRules) SetRuleActionType(v string) *GetDIJobResponseBodyDataTableMappingsTransformationRules {
	s.RuleActionType = &v
	return s
}

func (s *GetDIJobResponseBodyDataTableMappingsTransformationRules) SetRuleName(v string) *GetDIJobResponseBodyDataTableMappingsTransformationRules {
	s.RuleName = &v
	return s
}

func (s *GetDIJobResponseBodyDataTableMappingsTransformationRules) SetRuleTargetType(v string) *GetDIJobResponseBodyDataTableMappingsTransformationRules {
	s.RuleTargetType = &v
	return s
}

func (s *GetDIJobResponseBodyDataTableMappingsTransformationRules) Validate() error {
	return dara.Validate(s)
}

type GetDIJobResponseBodyDataTransformationRules struct {
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
	// The expression of the rule. The expression is a JSON string.
	//
	// 	- Example of a renaming rule: `{"expression":"${srcDatasourceName}_${srcDatabaseName}_0922","variables":[{"variableName":"srcDatabaseName","variableRules":[{"from":"fromdb","to":"todb"}\\]}\\]}`.
	//
	//     	- expression: the expression of the renaming rule. The expression may contain the following variables:
	//
	//         	- ${srcDatasourceName}
	//
	//         	- ${srcDatabaseName}
	//
	//         	- ${srcTableName}
	//
	//     	- variables: the generation rule for a variable used in the expression of the renaming rule. The default value of the specified variable is the original value of the object indicated by the variable. A group of string replacement rules used to change the original values may be returned.
	//
	//         	- variableName: the name of the variable. The variable name is not enclosed in ${}.
	//
	//         	- variableRules: the string replacement rules for variables. The system runs the string replacement rules in sequence. from indicates the original string. to indicates the new string.
	//
	// 	- Example of a rule used to add a specific field to the destination and assign a value to the field: `{"columns":[{"columnName":"my_add_column","columnValueType":"Constant","columnValue":"123"}\\]}`.
	//
	//     If no rule of this type is configured, no fields are added to the destination and no values are assigned by default.
	//
	//     	- columnName: the name of the field that is added.
	//
	//     	- columnValueType: the value type of the field. Valid values: Constant and Variable.
	//
	//     	- columnValue: the value of the field.
	//
	//         	- If the value of the columnValueType parameter is Constant, the value of the columnValue parameter is a constant of the STRING data type.
	//
	//         	- If the value of the columnValueType parameter is Variable, the value of the columnValue parameter is a built-in variable. The following built-in variables are supported: EXECUTE_TIME (LONG data type), DB_NAME_SRC (STRING data type), DATASOURCE_NAME_SRC (STRING data type), TABLE_NAME_SRC (STRING data type), DB_NAME_DEST (STRING data type), DATASOURCE_NAME_DEST (STRING data type), TABLE_NAME_DEST (STRING data type), and DB_NAME_SRC_TRANSED (STRING data type). EXECUTE_TIME indicates the execution time. DB_NAME_SRC indicates the name of a source database. DATASOURCE_NAME_SRC indicates the name of the source. TABLE_NAME_SRC indicates the name of a source table. DB_NAME_DEST indicates the name of a destination database. DATASOURCE_NAME_DEST indicates the name of the destination. TABLE_NAME_DEST indicates the name of a destination table. DB_NAME_SRC_TRANSED indicates the database name obtained after a transformation.
	//
	// 	- Example of a rule used to specify primary key fields for a destination table: `{"columns":["ukcolumn1","ukcolumn2"\\]}`.
	//
	//     If no rule of this type is configured, the primary key fields in the mapped source table are used for the destination table by default.
	//
	//     	- If the destination table is an existing table, Data Integration does not modify the schema of the destination table. If the specified primary key fields do not exist in the destination table, an error is reported when the synchronization task starts to run.
	//
	//     	- If the destination table is automatically created by the system, Data Integration automatically creates the schema of the destination table. The schema contains the primary key fields that you specify. If the specified primary key fields do not exist in the destination table, an error is reported when the synchronization task starts to run.
	//
	// 	- Example of a rule used to process DML messages: `{"dmlPolicies":[{"dmlType":"Delete","dmlAction":"Filter","filterCondition":"id > 1"}\\]}`.
	//
	//     If no rule of this type is configured, the default processing policy for messages generated for insert, update, and delete operations is Normal.
	//
	//     	- dmlType: the DML operation. Valid values: Insert, Update, and Delete.
	//
	//     	- dmlAction: the processing policy for DML messages. Valid values: Normal, Ignore, Filter, and LogicalDelete. Filter indicates conditional processing. The value Filter is returned for the dmlAction parameter only when the value of the dmlType parameter is Update or Delete.
	//
	//     	- filterCondition: the condition used to filter DML messages. This parameter is returned only when the value of the dmlAction parameter is Filter.
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
	// The type of the object on which the action is performed. Valid values:
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

func (s GetDIJobResponseBodyDataTransformationRules) String() string {
	return dara.Prettify(s)
}

func (s GetDIJobResponseBodyDataTransformationRules) GoString() string {
	return s.String()
}

func (s *GetDIJobResponseBodyDataTransformationRules) GetRuleActionType() *string {
	return s.RuleActionType
}

func (s *GetDIJobResponseBodyDataTransformationRules) GetRuleExpression() *string {
	return s.RuleExpression
}

func (s *GetDIJobResponseBodyDataTransformationRules) GetRuleName() *string {
	return s.RuleName
}

func (s *GetDIJobResponseBodyDataTransformationRules) GetRuleTargetType() *string {
	return s.RuleTargetType
}

func (s *GetDIJobResponseBodyDataTransformationRules) SetRuleActionType(v string) *GetDIJobResponseBodyDataTransformationRules {
	s.RuleActionType = &v
	return s
}

func (s *GetDIJobResponseBodyDataTransformationRules) SetRuleExpression(v string) *GetDIJobResponseBodyDataTransformationRules {
	s.RuleExpression = &v
	return s
}

func (s *GetDIJobResponseBodyDataTransformationRules) SetRuleName(v string) *GetDIJobResponseBodyDataTransformationRules {
	s.RuleName = &v
	return s
}

func (s *GetDIJobResponseBodyDataTransformationRules) SetRuleTargetType(v string) *GetDIJobResponseBodyDataTransformationRules {
	s.RuleTargetType = &v
	return s
}

func (s *GetDIJobResponseBodyDataTransformationRules) Validate() error {
	return dara.Validate(s)
}
