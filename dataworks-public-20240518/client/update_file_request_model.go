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
	// The advanced settings for the task.
	//
	// This parameter corresponds to the Advanced Settings in the right-side navigation pane on the editing page for EMR Spark Streaming and EMR Streaming SQL tasks in Data Studio in the [DataWorks console](https://workbench.data.aliyun.com/console).
	//
	// Currently, only EMR Spark Streaming and EMR Streaming SQL tasks support this parameter, and the parameter must be in JSON format.
	//
	// example:
	//
	// {"queue":"default","SPARK_CONF":"--conf spark.driver.memory=2g"}
	AdvancedSettings *string `json:"AdvancedSettings,omitempty" xml:"AdvancedSettings,omitempty"`
	// Specifies whether to apply the scheduling configuration immediately after the file is published.
	//
	// example:
	//
	// true
	ApplyScheduleImmediately *bool `json:"ApplyScheduleImmediately,omitempty" xml:"ApplyScheduleImmediately,omitempty"`
	// Specifies whether to enable automatic parsing for the file. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// This parameter corresponds to the Analyze Code setting in Properties > Dependencies for data development nodes in the [DataWorks console](https://workbench.data.aliyun.com/console).
	//
	// example:
	//
	// true
	AutoParsing *bool `json:"AutoParsing,omitempty" xml:"AutoParsing,omitempty"`
	// The interval at which the node is automatically rerun after a failure. Unit: milliseconds. Maximum value: 1800000 milliseconds (30 minutes).
	//
	// This parameter corresponds to the Rerun interval parameter in Properties > Schedule > Auto Rerun upon Failure for data development nodes in the [DataWorks console](https://workbench.data.aliyun.com/console). In the console, the unit of the rerun interval is minutes. Convert the time unit when you call this operation.
	//
	// example:
	//
	// 120000
	AutoRerunIntervalMillis *int32 `json:"AutoRerunIntervalMillis,omitempty" xml:"AutoRerunIntervalMillis,omitempty"`
	// The number of automatic reruns after the file execution fails.
	//
	// example:
	//
	// 3
	AutoRerunTimes *int32 `json:"AutoRerunTimes,omitempty" xml:"AutoRerunTimes,omitempty"`
	// The name of the data source that is used to run the node. You can call the [ListDataSources](https://help.aliyun.com/document_detail/211431.html) operation to query the available data sources.
	//
	// example:
	//
	// odps_source
	ConnectionName *string `json:"ConnectionName,omitempty" xml:"ConnectionName,omitempty"`
	// The file code content. Different code types (fileType) have different code formats. In Operation Center, you can right-click a task of the corresponding type and select View Code to view the specific code format.
	//
	// example:
	//
	// SELECT "1";
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The cron expression for scheduled execution. This parameter corresponds to the Cron Expression setting in Scheduling > Scheduling Time for Data Studio tasks in the [DataWorks console](https://workbench.data.aliyun.com/console). After you configure Scheduling Cycle and Scheduled Time, DataWorks automatically generates a cron expression.
	//
	// Examples:
	//
	// 	- Scheduled at 05:30 every day: `00 30 05 	- 	- ?`
	//
	// 	- Scheduled at the 15th minute of every hour: `00 15 	- 	- 	- ?`
	//
	// 	- Scheduled every 10 minutes: `00 00/10 	- 	- 	- ?`
	//
	// 	- Scheduled every 10 minutes between 08:00 and 23:00 every day: `00 00-59/10 8-23 	- 	- 	- ?`
	//
	// 	- Scheduled at 00:20 on the 1st day of every month: `00 20 00 1 	- ?`
	//
	// 	- Scheduled every 3 months starting from 00:10 on January 1: `00 10 00 1 1-12/3 ?`
	//
	// 	- Scheduled at 00:05 on every Tuesday and Friday: `00 05 00 	- 	- 2,5`
	//
	// Due to the rules of the DataWorks scheduling system, cron expressions have the following restrictions:
	//
	// 	- The minimum scheduling interval is 5 minutes.
	//
	// 	- The earliest scheduling time each day is 00:05.
	//
	// example:
	//
	// 00 00-59/5 1-23 	- 	- ?
	CronExpress *string `json:"CronExpress,omitempty" xml:"CronExpress,omitempty"`
	// The type of scheduling cycle. Valid values: NOT_DAY (minute, hour) and DAY (day, week, month).
	//
	// This parameter corresponds to the Scheduling Cycle setting in Scheduling > Scheduling Time for Data Studio tasks in the [DataWorks console](https://workbench.data.aliyun.com/console).
	//
	// example:
	//
	// NOT_DAY
	CycleType *string `json:"CycleType,omitempty" xml:"CycleType,omitempty"`
	// The IDs of the nodes on which the current node depends. This parameter takes effect only when the DependentType parameter is set to USER_DEFINE. Separate multiple node IDs with commas (,).
	//
	// This parameter corresponds to the Other Nodes option in Properties > Dependencies > Cross-cycle Dependency (Original Previous-cycle Dependency) for data development nodes in the [DataWorks console](https://workbench.data.aliyun.com/console).
	//
	// example:
	//
	// 5,10,15,20
	DependentNodeIdList *string `json:"DependentNodeIdList,omitempty" xml:"DependentNodeIdList,omitempty"`
	// The dependency mode on the previous cycle. Valid values:
	//
	// 	- SELF: Depends on the current node.
	//
	// 	- CHILD: Depends on the child nodes.
	//
	// 	- USER_DEFINE: Depends on other nodes.
	//
	// 	- NONE: No dependencies. Does not depend on the previous cycle.
	//
	// example:
	//
	// USER_DEFINE
	DependentType *string `json:"DependentType,omitempty" xml:"DependentType,omitempty"`
	// The timestamp (in milliseconds) when automatic scheduling stops.
	//
	// This parameter corresponds to the end time of Effective Period in Scheduling > Scheduling Time for Data Studio tasks in the [DataWorks console](https://workbench.data.aliyun.com/console).
	//
	// example:
	//
	// 4155787800000
	EndEffectDate *int64 `json:"EndEffectDate,omitempty" xml:"EndEffectDate,omitempty"`
	// The file description.
	//
	// example:
	//
	// Here is the file description
	FileDescription *string `json:"FileDescription,omitempty" xml:"FileDescription,omitempty"`
	// The path to the folder where the file is located.
	//
	// example:
	//
	// Business_process/First_Business_Process/data_integration/Folder_1/Folder_2
	FileFolderPath *string `json:"FileFolderPath,omitempty" xml:"FileFolderPath,omitempty"`
	// The file ID. You can call the [ListFiles](https://help.aliyun.com/document_detail/173942.html) operation to obtain the file ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 100000001
	FileId *int64 `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// The file name. You can modify the file name by setting a new value for FileName. For example, you can call the [ListFiles](https://help.aliyun.com/document_detail/173942.html) operation to query the file ID in the target directory, and then call the [UpdateFile](https://help.aliyun.com/document_detail/173951.html) operation with the file ID specified in the FileId parameter and a new value specified in the FileName parameter to modify the file name.
	//
	// example:
	//
	// ods_user_info_d
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// This parameter corresponds to the Skip The Dry-Run Property Of The Ancestor Node option in Properties > Dependencies > Cross-cycle Dependency (Original Previous-cycle Dependency) when Instances of Current Node or Level-1 Child Node is selected for data development nodes in the [DataWorks console](https://workbench.data.aliyun.com/console).
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
	// The output names of the ancestor nodes on which the current node depends. Separate multiple output names with commas (,).
	//
	// This parameter corresponds to the Output Name of Ancestor Node setting in Properties > Dependencies for data development nodes in the [DataWorks console](https://workbench.data.aliyun.com/console).
	//
	// > This parameter is required when you call the CreateDISyncTask or UpdateFile operation to create a batch synchronization node.
	//
	// example:
	//
	// project_root,project.file1,project.001_out
	InputList *string `json:"InputList,omitempty" xml:"InputList,omitempty"`
	// The input context parameters of the node. The value must be in the JSON format. For more information about the parameter structure, see the InputContextParameterList parameter in the response parameters of the [GetFile](https://help.aliyun.com/document_detail/173954.html) operation.
	//
	// This parameter corresponds to the Input Parameters setting in Properties > Input and Output Parameters for data development nodes in the [DataWorks console](https://workbench.data.aliyun.com/console).
	//
	// example:
	//
	// [{"ValueSource": "project_001.first_node:bizdate_param","ParameterName": "bizdate_input"}]
	InputParameters *string `json:"InputParameters,omitempty" xml:"InputParameters,omitempty"`
	// The outputs of the node.
	//
	// This parameter corresponds to the Output Name setting in Properties > Dependencies for data development nodes in the [DataWorks console](https://workbench.data.aliyun.com/console).
	//
	// example:
	//
	// dw_project.ods_user_info_d
	OutputList *string `json:"OutputList,omitempty" xml:"OutputList,omitempty"`
	// The output context parameters of the node. The value must be in the JSON format. For more information about the parameter structure, see the OutputContextParameterList parameter in the response parameters of the [GetFile](https://help.aliyun.com/document_detail/173954.html) operation.
	//
	// This parameter corresponds to the Output Parameters setting in Properties > Input and Output Parameters for data development nodes in the [DataWorks console](https://workbench.data.aliyun.com/console).
	//
	// example:
	//
	// [{"Type": 1,"Value": "${bizdate}","ParameterName": "bizdate_param"}]
	OutputParameters *string `json:"OutputParameters,omitempty" xml:"OutputParameters,omitempty"`
	// The file owner ID.
	//
	// example:
	//
	// 18023848927592
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The scheduling parameters of the node.
	//
	// This parameter corresponds to the Scheduling Parameter setting in Properties for data development nodes in the [DataWorks console](https://workbench.data.aliyun.com/console). For more information, see [Scheduling parameters](https://help.aliyun.com/document_detail/137548.html).
	//
	// example:
	//
	// x=a y=b z=c
	ParaValue *string `json:"ParaValue,omitempty" xml:"ParaValue,omitempty"`
	// The DataWorks workspace ID. To obtain the ID, log on to the [DataWorks console](https://workbench.data.aliyun.com/console) and navigate to the workspace management page.
	//
	// example:
	//
	// 100001
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The DataWorks workspace name. To obtain the workspace name, log on to the [DataWorks console](https://workbench.data.aliyun.com/console) and navigate to the workspace configuration page.
	//
	// You must specify either this parameter or ProjectId to identify the target DataWorks workspace for this API call.
	//
	// example:
	//
	// dw_project
	ProjectIdentifier *string `json:"ProjectIdentifier,omitempty" xml:"ProjectIdentifier,omitempty"`
	// The rerun policy. Valid values:
	//
	// 	- ALL_ALLOWED: Reruns are allowed regardless of whether the task succeeds or fails.
	//
	// 	- FAILURE_ALLOWED: Reruns are allowed only when the task fails.
	//
	// 	- ALL_DENIED: Reruns are not allowed regardless of whether the task succeeds or fails.
	//
	// This parameter corresponds to the Support for Rerun setting in Scheduling > Scheduling Policies for Data Studio tasks in the [DataWorks console](https://workbench.data.aliyun.com/console).
	//
	// Valid values:
	//
	// 	- ALL_ALLOWD
	//
	// 	- FAILURE_ALLOWED
	//
	// 	- ALL_DENIED
	//
	// 	- ALL_ALLOWED
	//
	// example:
	//
	// ALL_ALLOWED
	RerunMode *string `json:"RerunMode,omitempty" xml:"RerunMode,omitempty"`
	// The resource group for the task published from the file. You can call the [ListResourceGroups](https://help.aliyun.com/document_detail/173913.html) operation to query the available resource groups in the workspace.
	//
	// example:
	//
	// default_group
	ResourceGroupIdentifier *string `json:"ResourceGroupIdentifier,omitempty" xml:"ResourceGroupIdentifier,omitempty"`
	// The scheduling type. Valid values:
	//
	// 	- NORMAL: Normal scheduled task.
	//
	// 	- MANUAL: Manually triggered node. Not scheduled for daily execution. Corresponds to nodes in manually triggered workflows.
	//
	// 	- PAUSE: Paused task.
	//
	// 	- SKIP: Dry-run task. Scheduled for daily execution but is directly marked as successful when scheduling starts.
	//
	// example:
	//
	// NORMAL
	SchedulerType *string `json:"SchedulerType,omitempty" xml:"SchedulerType,omitempty"`
	// The timestamp (in milliseconds) when automatic scheduling starts.
	//
	// This parameter corresponds to the start time of Effective Period in Scheduling > Scheduling Time for Data Studio tasks in the [DataWorks console](https://workbench.data.aliyun.com/console).
	//
	// example:
	//
	// 936923400000
	StartEffectDate *int64 `json:"StartEffectDate,omitempty" xml:"StartEffectDate,omitempty"`
	// Specifies whether to start the task immediately after it is published. Valid values:
	//
	// 	- true: Start the task immediately after it is published.
	//
	// 	- false: Do not start the task immediately after it is published.
	//
	// This parameter corresponds to the Start Method setting in Configuration > Scheduling Policies in the right-side navigation pane on the editing page for EMR Spark Streaming and EMR Streaming SQL tasks in Data Studio in the [DataWorks console](https://workbench.data.aliyun.com/console).
	//
	// example:
	//
	// true
	StartImmediately *bool `json:"StartImmediately,omitempty" xml:"StartImmediately,omitempty"`
	// Specifies whether to skip execution. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// This parameter corresponds to the Skip Execution option in Properties > Schedule > Recurrence for data development nodes in the [DataWorks console](https://workbench.data.aliyun.com/console).
	//
	// example:
	//
	// false
	Stop *bool `json:"Stop,omitempty" xml:"Stop,omitempty"`
	// The timeout settings for scheduling configuration.
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
