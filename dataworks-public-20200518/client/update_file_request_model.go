// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateFileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAdvancedSettings(v string) *UpdateFileRequest
	GetAdvancedSettings() *string
	SetApplyScheduleImmediately(v bool) *UpdateFileRequest
	GetApplyScheduleImmediately() *bool
	SetAutoParsing(v bool) *UpdateFileRequest
	GetAutoParsing() *bool
	SetAutoRerunIntervalMillis(v int32) *UpdateFileRequest
	GetAutoRerunIntervalMillis() *int32
	SetAutoRerunTimes(v int32) *UpdateFileRequest
	GetAutoRerunTimes() *int32
	SetConnectionName(v string) *UpdateFileRequest
	GetConnectionName() *string
	SetContent(v string) *UpdateFileRequest
	GetContent() *string
	SetCronExpress(v string) *UpdateFileRequest
	GetCronExpress() *string
	SetCycleType(v string) *UpdateFileRequest
	GetCycleType() *string
	SetDependentNodeIdList(v string) *UpdateFileRequest
	GetDependentNodeIdList() *string
	SetDependentType(v string) *UpdateFileRequest
	GetDependentType() *string
	SetEndEffectDate(v int64) *UpdateFileRequest
	GetEndEffectDate() *int64
	SetFileDescription(v string) *UpdateFileRequest
	GetFileDescription() *string
	SetFileFolderPath(v string) *UpdateFileRequest
	GetFileFolderPath() *string
	SetFileId(v int64) *UpdateFileRequest
	GetFileId() *int64
	SetFileName(v string) *UpdateFileRequest
	GetFileName() *string
	SetIgnoreParentSkipRunningProperty(v bool) *UpdateFileRequest
	GetIgnoreParentSkipRunningProperty() *bool
	SetImageId(v string) *UpdateFileRequest
	GetImageId() *string
	SetInputList(v string) *UpdateFileRequest
	GetInputList() *string
	SetInputParameters(v string) *UpdateFileRequest
	GetInputParameters() *string
	SetOutputList(v string) *UpdateFileRequest
	GetOutputList() *string
	SetOutputParameters(v string) *UpdateFileRequest
	GetOutputParameters() *string
	SetOwner(v string) *UpdateFileRequest
	GetOwner() *string
	SetParaValue(v string) *UpdateFileRequest
	GetParaValue() *string
	SetProjectId(v int64) *UpdateFileRequest
	GetProjectId() *int64
	SetProjectIdentifier(v string) *UpdateFileRequest
	GetProjectIdentifier() *string
	SetRerunMode(v string) *UpdateFileRequest
	GetRerunMode() *string
	SetResourceGroupIdentifier(v string) *UpdateFileRequest
	GetResourceGroupIdentifier() *string
	SetSchedulerType(v string) *UpdateFileRequest
	GetSchedulerType() *string
	SetStartEffectDate(v int64) *UpdateFileRequest
	GetStartEffectDate() *int64
	SetStartImmediately(v bool) *UpdateFileRequest
	GetStartImmediately() *bool
	SetStop(v bool) *UpdateFileRequest
	GetStop() *bool
	SetTimeout(v int32) *UpdateFileRequest
	GetTimeout() *int32
}

type UpdateFileRequest struct {
	// The advanced settings of the node.
	//
	// This parameter corresponds to the Advanced Settings in the right-side navigation pane on the editing page for EMR Spark Streaming and EMR Streaming SQL data development nodes in the [DataWorks console](https://workbench.data.aliyun.com/console).
	//
	// Currently, only EMR Spark Streaming and EMR Streaming SQL nodes support this parameter. The parameter value is in JSON format.
	//
	// example:
	//
	// {"queue":"default","SPARK_CONF":"--conf spark.driver.memory=2g"}
	AdvancedSettings *string `json:"AdvancedSettings,omitempty" xml:"AdvancedSettings,omitempty"`
	// Specifies whether the scheduling configuration takes effect immediately after publishing.
	//
	// example:
	//
	// true
	ApplyScheduleImmediately *bool `json:"ApplyScheduleImmediately,omitempty" xml:"ApplyScheduleImmediately,omitempty"`
	// Specifies whether to enable the automatic parsing feature for the file. Valid values:
	//
	// - true: The file automatically parses code.
	//
	// - false: The file does not automatically parse code.
	//
	// This parameter corresponds to the Code Parsing setting when you select Same Cycle under Scheduling Configuration > Scheduling Dependency for a data development node in the [DataWorks console](https://workbench.data.aliyun.com/console).
	//
	// example:
	//
	// true
	AutoParsing *bool `json:"AutoParsing,omitempty" xml:"AutoParsing,omitempty"`
	// The interval between automatic reruns upon an error, in milliseconds. The maximum value is 1800000 milliseconds (30 minutes).
	//
	// This parameter corresponds to the Rerun Interval setting under Scheduling Configuration > Time Properties > Auto Rerun upon Error for a data development node in the [DataWorks console](https://workbench.data.aliyun.com/console).
	//
	// The time unit for Rerun Interval in the console is minutes. Convert the time accordingly when calling this operation.
	//
	// example:
	//
	// 120000
	AutoRerunIntervalMillis *int32 `json:"AutoRerunIntervalMillis,omitempty" xml:"AutoRerunIntervalMillis,omitempty"`
	// The number of automatic reruns after an error occurs.
	//
	// example:
	//
	// 3
	AutoRerunTimes *int32 `json:"AutoRerunTimes,omitempty" xml:"AutoRerunTimes,omitempty"`
	// The identifier of the data source used when the node corresponding to the file runs. You can call the [ListDataSources](https://help.aliyun.com/document_detail/211431.html) operation to obtain the list of available data sources.
	//
	// example:
	//
	// odps_source
	ConnectionName *string `json:"ConnectionName,omitempty" xml:"ConnectionName,omitempty"`
	// The code content of the file. Files of different code types (fileType) have different code formats. In Operation Center, right-click a node of the corresponding type and select View Code to view the specific code format.
	//
	// example:
	//
	// SELECT "1";
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The cron expression for timed scheduling. This parameter corresponds to the cron Expression setting under Scheduling Configuration > Time Property for a data development node in the [DataWorks console](https://workbench.data.aliyun.com/console). After you configure the Scheduling Epoch and Timed Scheduling time, DataWorks automatically generates the corresponding cron expression.
	//
	// Examples:
	//
	// - Timed scheduling at 05:30 every day: `00 30 05 	- 	- ?`.
	//
	// - Timed scheduling at the 15th minute of every hour: `00 15 	- 	- 	- ?`.
	//
	// - Schedule every 10 minutes: `00 00/10 	- 	- 	- ?`.
	//
	// - Schedule every 10 minutes from 08:00 to 17:00 every day: `00 00-59/10 8-23 	- 	- 	- ?`.
	//
	// - Timed scheduling at 00:20 on the 1st of every month: `00 20 00 1 	- ?`.
	//
	// - Schedule every 3 months starting from 00:10 on January 1: `00 10 00 1 1-12/3 ?`.
	//
	// - Timed scheduling at 00:05 every Tuesday and Friday: `00 05 00 	- 	- 2,5`.
	//
	//
	// The cron expression has the following limits due to the DataWorks scheduling system rules:
	//
	// - The minimum scheduling interval is 5 minutes.
	//
	// - The earliest scheduling time each day is 00:05.
	//
	// example:
	//
	// 00 00-59/5 1-23 	- 	- ?
	CronExpress *string `json:"CronExpress,omitempty" xml:"CronExpress,omitempty"`
	// The type of the scheduling cycle. Valid values: NOT_DAY (minute or hour) and DAY (day, week, or month).
	//
	// This parameter corresponds to the Scheduling Cycle setting under Scheduling Configuration > Time Properties for a data development node in the [DataWorks console](https://workbench.data.aliyun.com/console).
	//
	// example:
	//
	// NOT_DAY
	CycleType *string `json:"CycleType,omitempty" xml:"CycleType,omitempty"`
	// The IDs of the nodes on which the current file depends when DependentType is set to USER_DEFINE. Separate multiple node IDs with commas (,).
	//
	// This parameter corresponds to the Settings when you select Other Nodes as the dependency after configuring Scheduling Configuration > Scheduling Dependency to Previous Epoch for a data development node in the [DataWorks console](https://workbench.data.aliyun.com/console). This is part of the parameter settings for scheduling dependencies.
	//
	// example:
	//
	// 5,10,15,20
	DependentNodeIdList *string `json:"DependentNodeIdList,omitempty" xml:"DependentNodeIdList,omitempty"`
	// The mode in which the node depends on the previous cycle. Valid values:
	//
	// - SELF: The dependency is set to the current node.
	//
	// - CHILD: The dependency is set to first-level child nodes.
	//
	// - USER_DEFINE: The dependency is set to other nodes.
	//
	// - NONE: No dependency is selected. The node does not depend on the previous cycle.
	//
	// example:
	//
	// USER_DEFINE
	DependentType *string `json:"DependentType,omitempty" xml:"DependentType,omitempty"`
	// The timestamp in milliseconds when automatic scheduling stops.
	//
	// This parameter corresponds to the end time in milliseconds under Scheduling Configuration > Time Properties > Effective Date for a data development node in the [DataWorks console](https://workbench.data.aliyun.com/console).
	//
	// example:
	//
	// 4155787800000
	EndEffectDate *int64 `json:"EndEffectDate,omitempty" xml:"EndEffectDate,omitempty"`
	// The description of the file.
	//
	// example:
	//
	// Here is the file description
	FileDescription *string `json:"FileDescription,omitempty" xml:"FileDescription,omitempty"`
	// The path of the file.
	//
	// example:
	//
	// Business_process/First_Business_Process/data_integration/Folder_1/Folder_2
	FileFolderPath *string `json:"FileFolderPath,omitempty" xml:"FileFolderPath,omitempty"`
	// The ID of the file. You can call the [ListFiles](https://help.aliyun.com/document_detail/173942.html) operation to obtain the file ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 100000001
	FileId *int64 `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// The name of the file. You can modify the file name by setting FileName to a new value.
	//
	// For example, call the [ListFiles](https://help.aliyun.com/document_detail/173942.html) operation to query the file ID in the target folder, and then call the [UpdateFile](https://help.aliyun.com/document_detail/173951.html) operation to specify the file ID for the FileId parameter and configure the FileName parameter to rename the file.
	//
	// example:
	//
	// ods_user_info_d
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// Specifies whether to skip the dry-run property of the upstream node under Scheduling Configuration > Previous Cycle.
	//
	// example:
	//
	// true
	IgnoreParentSkipRunningProperty *bool `json:"IgnoreParentSkipRunningProperty,omitempty" xml:"IgnoreParentSkipRunningProperty,omitempty"`
	// The custom image ID.
	//
	// example:
	//
	// m-uf6d7npxk1hhek8ng0cb
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The output names of the upstream files on which the current file depends. Separate multiple output names with commas (,).
	//
	// This parameter corresponds to the Parent Node Output Name setting when you select Same Cycle under Scheduling Configuration > Scheduling Dependency for a data development node in the [DataWorks console](https://workbench.data.aliyun.com/console).
	//
	// > This parameter is required when you create a batch synchronization node by calling CreateDISyncTask and UpdateFile.
	//
	// example:
	//
	// project_root,project.file1,project.001_out
	InputList *string `json:"InputList,omitempty" xml:"InputList,omitempty"`
	// The input context parameters of the node. The parameter value is in JSON format. For the fields included, refer to the InputContextParameterList parameter structure in the response of the [GetFile](https://help.aliyun.com/document_detail/173954.html) operation.
	//
	// This parameter corresponds to the Input Parameters of Current Node setting under Scheduling Configuration > Node Context for a data development node in the [DataWorks console](https://workbench.data.aliyun.com/console).
	//
	// example:
	//
	// [{"ValueSource": "project_001.first_node:bizdate_param","ParameterName": "bizdate_input"}]
	InputParameters *string `json:"InputParameters,omitempty" xml:"InputParameters,omitempty"`
	// The output of the file.
	//
	// This parameter corresponds to the Output Name of Current Node setting when you select Same Cycle under Scheduling Configuration > Scheduling Dependency for a data development node in the [DataWorks console](https://workbench.data.aliyun.com/console).
	//
	// example:
	//
	// dw_project.ods_user_info_d
	OutputList *string `json:"OutputList,omitempty" xml:"OutputList,omitempty"`
	// The output context parameters of the node. The parameter value is in JSON format. For the fields included, refer to the OutputContextParameterList parameter structure in the response of the [GetFile](https://help.aliyun.com/document_detail/173954.html) operation.
	//
	// This parameter corresponds to the Output Parameters of Current Node setting under Scheduling Configuration > Node Context for a data development node in the [DataWorks console](https://workbench.data.aliyun.com/console).
	//
	// example:
	//
	// [{"Type": 1,"Value": "${bizdate}","ParameterName": "bizdate_param"}]
	OutputParameters *string `json:"OutputParameters,omitempty" xml:"OutputParameters,omitempty"`
	// The user ID of the file owner.
	//
	// example:
	//
	// 18023848927592
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The scheduling parameters.
	//
	// This parameter corresponds to the Parameters setting under Scheduling Configuration for a data development node in the [DataWorks console](https://workbench.data.aliyun.com/console). For more information, see [Scheduling parameters](https://help.aliyun.com/document_detail/137548.html).
	//
	// example:
	//
	// x=a y=b z=c
	ParaValue *string `json:"ParaValue,omitempty" xml:"ParaValue,omitempty"`
	// The ID of the DataWorks workspace. You can logon to the [DataWorks console](https://workbench.data.aliyun.com/console) and go to the Storage Management page to obtain the ID.
	//
	// example:
	//
	// 100001
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The name of the DataWorks workspace. You can log on to the [DataWorks console](https://workbench.data.aliyun.com/console) and go to the Workspace Settings page to obtain the workspace name.
	//
	// You must specify either this parameter or ProjectId to determine the DataWorks workspace for this API call.
	//
	// example:
	//
	// dw_project
	ProjectIdentifier *string `json:"ProjectIdentifier,omitempty" xml:"ProjectIdentifier,omitempty"`
	// The rerun property. Valid values:
	//
	// - ALL_ALLOWED: The node can be rerun regardless of whether it runs successfully or fails.
	//
	// - FAILURE_ALLOWED: The node can be rerun only after it fails.
	//
	// - ALL_DENIED: The node cannot be rerun regardless of whether it runs successfully or fails.
	//
	// This parameter corresponds to the Rerun Property setting under Scheduling Configuration > Time Properties > Rerun Property for a data development node in the [DataWorks console](https://workbench.data.aliyun.com/console).
	//
	// example:
	//
	// ALL_ALLOWED
	RerunMode *string `json:"RerunMode,omitempty" xml:"RerunMode,omitempty"`
	// The schedule resource used when the file is published as a node and the node runs. You can call the [ListResourceGroups](https://help.aliyun.com/document_detail/173913.html) operation to obtain the list of available resource groups for the workspace.
	//
	// example:
	//
	// default_group
	ResourceGroupIdentifier *string `json:"ResourceGroupIdentifier,omitempty" xml:"ResourceGroupIdentifier,omitempty"`
	// The scheduling type. Valid values:
	//
	// - NORMAL: A normal scheduling node.
	//
	// - MANUAL: A manual node that is not scheduled on a daily basis. This corresponds to nodes in a manual workflow.
	//
	// - PAUSE: A paused node.
	//
	// - SKIP: A dry-run node that is scheduled on a daily basis but is directly set to successful when scheduling starts.
	//
	// example:
	//
	// NORMAL
	SchedulerType *string `json:"SchedulerType,omitempty" xml:"SchedulerType,omitempty"`
	// The timestamp in milliseconds when automatic scheduling starts.
	//
	// This parameter corresponds to the start time in milliseconds under Scheduling Configuration > Time Properties > Effective Date for a data development node in the [DataWorks console](https://workbench.data.aliyun.com/console).
	//
	// example:
	//
	// 936923400000
	StartEffectDate *int64 `json:"StartEffectDate,omitempty" xml:"StartEffectDate,omitempty"`
	// Specifies whether to start the node immediately after publishing. Valid values:
	//
	// - true: Starts immediately after publishing.
	//
	// - false: Does not start after publishing.
	//
	// This parameter corresponds to the Start Mode setting under Configuration > Time Properties in the right-side navigation pane on the editing page for EMR Spark Streaming and EMR Streaming SQL data development nodes in the [DataWorks console](https://workbench.data.aliyun.com/console).
	//
	// example:
	//
	// true
	StartImmediately *bool `json:"StartImmediately,omitempty" xml:"StartImmediately,omitempty"`
	// Specifies whether to pause scheduling. Valid values:
	//
	// - true: Pauses scheduling.
	//
	// - false: Does not pause scheduling.
	//
	// This parameter corresponds to the setting when Scheduling Type is set to Pause Scheduling under Scheduling Configuration > Time Properties > Scheduling Type for a data development node in the [DataWorks console](https://workbench.data.aliyun.com/console).
	//
	// example:
	//
	// false
	Stop *bool `json:"Stop,omitempty" xml:"Stop,omitempty"`
	// The timeout setting for the scheduling configuration.
	//
	// example:
	//
	// 1
	Timeout *int32 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
}

func (s UpdateFileRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateFileRequest) GoString() string {
	return s.String()
}

func (s *UpdateFileRequest) GetAdvancedSettings() *string {
	return s.AdvancedSettings
}

func (s *UpdateFileRequest) GetApplyScheduleImmediately() *bool {
	return s.ApplyScheduleImmediately
}

func (s *UpdateFileRequest) GetAutoParsing() *bool {
	return s.AutoParsing
}

func (s *UpdateFileRequest) GetAutoRerunIntervalMillis() *int32 {
	return s.AutoRerunIntervalMillis
}

func (s *UpdateFileRequest) GetAutoRerunTimes() *int32 {
	return s.AutoRerunTimes
}

func (s *UpdateFileRequest) GetConnectionName() *string {
	return s.ConnectionName
}

func (s *UpdateFileRequest) GetContent() *string {
	return s.Content
}

func (s *UpdateFileRequest) GetCronExpress() *string {
	return s.CronExpress
}

func (s *UpdateFileRequest) GetCycleType() *string {
	return s.CycleType
}

func (s *UpdateFileRequest) GetDependentNodeIdList() *string {
	return s.DependentNodeIdList
}

func (s *UpdateFileRequest) GetDependentType() *string {
	return s.DependentType
}

func (s *UpdateFileRequest) GetEndEffectDate() *int64 {
	return s.EndEffectDate
}

func (s *UpdateFileRequest) GetFileDescription() *string {
	return s.FileDescription
}

func (s *UpdateFileRequest) GetFileFolderPath() *string {
	return s.FileFolderPath
}

func (s *UpdateFileRequest) GetFileId() *int64 {
	return s.FileId
}

func (s *UpdateFileRequest) GetFileName() *string {
	return s.FileName
}

func (s *UpdateFileRequest) GetIgnoreParentSkipRunningProperty() *bool {
	return s.IgnoreParentSkipRunningProperty
}

func (s *UpdateFileRequest) GetImageId() *string {
	return s.ImageId
}

func (s *UpdateFileRequest) GetInputList() *string {
	return s.InputList
}

func (s *UpdateFileRequest) GetInputParameters() *string {
	return s.InputParameters
}

func (s *UpdateFileRequest) GetOutputList() *string {
	return s.OutputList
}

func (s *UpdateFileRequest) GetOutputParameters() *string {
	return s.OutputParameters
}

func (s *UpdateFileRequest) GetOwner() *string {
	return s.Owner
}

func (s *UpdateFileRequest) GetParaValue() *string {
	return s.ParaValue
}

func (s *UpdateFileRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *UpdateFileRequest) GetProjectIdentifier() *string {
	return s.ProjectIdentifier
}

func (s *UpdateFileRequest) GetRerunMode() *string {
	return s.RerunMode
}

func (s *UpdateFileRequest) GetResourceGroupIdentifier() *string {
	return s.ResourceGroupIdentifier
}

func (s *UpdateFileRequest) GetSchedulerType() *string {
	return s.SchedulerType
}

func (s *UpdateFileRequest) GetStartEffectDate() *int64 {
	return s.StartEffectDate
}

func (s *UpdateFileRequest) GetStartImmediately() *bool {
	return s.StartImmediately
}

func (s *UpdateFileRequest) GetStop() *bool {
	return s.Stop
}

func (s *UpdateFileRequest) GetTimeout() *int32 {
	return s.Timeout
}

func (s *UpdateFileRequest) SetAdvancedSettings(v string) *UpdateFileRequest {
	s.AdvancedSettings = &v
	return s
}

func (s *UpdateFileRequest) SetApplyScheduleImmediately(v bool) *UpdateFileRequest {
	s.ApplyScheduleImmediately = &v
	return s
}

func (s *UpdateFileRequest) SetAutoParsing(v bool) *UpdateFileRequest {
	s.AutoParsing = &v
	return s
}

func (s *UpdateFileRequest) SetAutoRerunIntervalMillis(v int32) *UpdateFileRequest {
	s.AutoRerunIntervalMillis = &v
	return s
}

func (s *UpdateFileRequest) SetAutoRerunTimes(v int32) *UpdateFileRequest {
	s.AutoRerunTimes = &v
	return s
}

func (s *UpdateFileRequest) SetConnectionName(v string) *UpdateFileRequest {
	s.ConnectionName = &v
	return s
}

func (s *UpdateFileRequest) SetContent(v string) *UpdateFileRequest {
	s.Content = &v
	return s
}

func (s *UpdateFileRequest) SetCronExpress(v string) *UpdateFileRequest {
	s.CronExpress = &v
	return s
}

func (s *UpdateFileRequest) SetCycleType(v string) *UpdateFileRequest {
	s.CycleType = &v
	return s
}

func (s *UpdateFileRequest) SetDependentNodeIdList(v string) *UpdateFileRequest {
	s.DependentNodeIdList = &v
	return s
}

func (s *UpdateFileRequest) SetDependentType(v string) *UpdateFileRequest {
	s.DependentType = &v
	return s
}

func (s *UpdateFileRequest) SetEndEffectDate(v int64) *UpdateFileRequest {
	s.EndEffectDate = &v
	return s
}

func (s *UpdateFileRequest) SetFileDescription(v string) *UpdateFileRequest {
	s.FileDescription = &v
	return s
}

func (s *UpdateFileRequest) SetFileFolderPath(v string) *UpdateFileRequest {
	s.FileFolderPath = &v
	return s
}

func (s *UpdateFileRequest) SetFileId(v int64) *UpdateFileRequest {
	s.FileId = &v
	return s
}

func (s *UpdateFileRequest) SetFileName(v string) *UpdateFileRequest {
	s.FileName = &v
	return s
}

func (s *UpdateFileRequest) SetIgnoreParentSkipRunningProperty(v bool) *UpdateFileRequest {
	s.IgnoreParentSkipRunningProperty = &v
	return s
}

func (s *UpdateFileRequest) SetImageId(v string) *UpdateFileRequest {
	s.ImageId = &v
	return s
}

func (s *UpdateFileRequest) SetInputList(v string) *UpdateFileRequest {
	s.InputList = &v
	return s
}

func (s *UpdateFileRequest) SetInputParameters(v string) *UpdateFileRequest {
	s.InputParameters = &v
	return s
}

func (s *UpdateFileRequest) SetOutputList(v string) *UpdateFileRequest {
	s.OutputList = &v
	return s
}

func (s *UpdateFileRequest) SetOutputParameters(v string) *UpdateFileRequest {
	s.OutputParameters = &v
	return s
}

func (s *UpdateFileRequest) SetOwner(v string) *UpdateFileRequest {
	s.Owner = &v
	return s
}

func (s *UpdateFileRequest) SetParaValue(v string) *UpdateFileRequest {
	s.ParaValue = &v
	return s
}

func (s *UpdateFileRequest) SetProjectId(v int64) *UpdateFileRequest {
	s.ProjectId = &v
	return s
}

func (s *UpdateFileRequest) SetProjectIdentifier(v string) *UpdateFileRequest {
	s.ProjectIdentifier = &v
	return s
}

func (s *UpdateFileRequest) SetRerunMode(v string) *UpdateFileRequest {
	s.RerunMode = &v
	return s
}

func (s *UpdateFileRequest) SetResourceGroupIdentifier(v string) *UpdateFileRequest {
	s.ResourceGroupIdentifier = &v
	return s
}

func (s *UpdateFileRequest) SetSchedulerType(v string) *UpdateFileRequest {
	s.SchedulerType = &v
	return s
}

func (s *UpdateFileRequest) SetStartEffectDate(v int64) *UpdateFileRequest {
	s.StartEffectDate = &v
	return s
}

func (s *UpdateFileRequest) SetStartImmediately(v bool) *UpdateFileRequest {
	s.StartImmediately = &v
	return s
}

func (s *UpdateFileRequest) SetStop(v bool) *UpdateFileRequest {
	s.Stop = &v
	return s
}

func (s *UpdateFileRequest) SetTimeout(v int32) *UpdateFileRequest {
	s.Timeout = &v
	return s
}

func (s *UpdateFileRequest) Validate() error {
	return dara.Validate(s)
}
