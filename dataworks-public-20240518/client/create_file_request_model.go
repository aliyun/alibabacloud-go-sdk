// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAdvancedSettings(v string) *CreateFileRequest
	GetAdvancedSettings() *string
	SetApplyScheduleImmediately(v bool) *CreateFileRequest
	GetApplyScheduleImmediately() *bool
	SetAutoParsing(v bool) *CreateFileRequest
	GetAutoParsing() *bool
	SetAutoRerunIntervalMillis(v int32) *CreateFileRequest
	GetAutoRerunIntervalMillis() *int32
	SetAutoRerunTimes(v int32) *CreateFileRequest
	GetAutoRerunTimes() *int32
	SetConnectionName(v string) *CreateFileRequest
	GetConnectionName() *string
	SetContent(v string) *CreateFileRequest
	GetContent() *string
	SetCreateFolderIfNotExists(v bool) *CreateFileRequest
	GetCreateFolderIfNotExists() *bool
	SetCronExpress(v string) *CreateFileRequest
	GetCronExpress() *string
	SetCycleType(v string) *CreateFileRequest
	GetCycleType() *string
	SetDependentNodeIdList(v string) *CreateFileRequest
	GetDependentNodeIdList() *string
	SetDependentType(v string) *CreateFileRequest
	GetDependentType() *string
	SetEndEffectDate(v int64) *CreateFileRequest
	GetEndEffectDate() *int64
	SetFileDescription(v string) *CreateFileRequest
	GetFileDescription() *string
	SetFileFolderPath(v string) *CreateFileRequest
	GetFileFolderPath() *string
	SetFileName(v string) *CreateFileRequest
	GetFileName() *string
	SetFileType(v int32) *CreateFileRequest
	GetFileType() *int32
	SetIgnoreParentSkipRunningProperty(v bool) *CreateFileRequest
	GetIgnoreParentSkipRunningProperty() *bool
	SetImageId(v string) *CreateFileRequest
	GetImageId() *string
	SetInputList(v string) *CreateFileRequest
	GetInputList() *string
	SetInputParameters(v string) *CreateFileRequest
	GetInputParameters() *string
	SetOutputParameters(v string) *CreateFileRequest
	GetOutputParameters() *string
	SetOwner(v string) *CreateFileRequest
	GetOwner() *string
	SetParaValue(v string) *CreateFileRequest
	GetParaValue() *string
	SetProjectId(v int64) *CreateFileRequest
	GetProjectId() *int64
	SetProjectIdentifier(v string) *CreateFileRequest
	GetProjectIdentifier() *string
	SetRerunMode(v string) *CreateFileRequest
	GetRerunMode() *string
	SetResourceGroupId(v int64) *CreateFileRequest
	GetResourceGroupId() *int64
	SetResourceGroupIdentifier(v string) *CreateFileRequest
	GetResourceGroupIdentifier() *string
	SetSchedulerType(v string) *CreateFileRequest
	GetSchedulerType() *string
	SetStartEffectDate(v int64) *CreateFileRequest
	GetStartEffectDate() *int64
	SetStartImmediately(v bool) *CreateFileRequest
	GetStartImmediately() *bool
	SetStop(v bool) *CreateFileRequest
	GetStop() *bool
	SetTimeout(v int32) *CreateFileRequest
	GetTimeout() *int32
}

type CreateFileRequest struct {
	// The advanced settings of the node.
	//
	// This parameter corresponds to the Advanced Settings section in the right-side navigation pane on the configuration tab of EMR Spark Streaming and EMR Streaming SQL nodes in the [DataWorks console](https://workbench.data.aliyun.com/console).
	//
	// Only EMR Spark Streaming and EMR Streaming SQL nodes support this parameter. The value must be in the JSON format.
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
	// The number of automatic reruns after an error occurs. Maximum value: 10.
	//
	// example:
	//
	// 3
	AutoRerunTimes *int32 `json:"AutoRerunTimes,omitempty" xml:"AutoRerunTimes,omitempty"`
	// The data source used when the task published from the file is run.
	//
	// You can call the [UpdateDataSource](https://help.aliyun.com/document_detail/211432.html) operation to query the available data sources in the workspace.
	//
	// example:
	//
	// odps_source
	ConnectionName *string `json:"ConnectionName,omitempty" xml:"ConnectionName,omitempty"`
	// The file code content. Different code types (fileType) have different code formats. In Operation Center, you can find a task of the corresponding type, right-click it, and select View Code to view the specific code format.
	//
	// example:
	//
	// SHOW TABLES;
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// Specifies whether to automatically create the directory specified by FileFolderPath if the directory does not exist. Valid values:
	//
	// 	- true: If the directory does not exist, automatically create it.
	//
	// 	- false: If the directory does not exist, the call fails.
	//
	// example:
	//
	// false
	CreateFolderIfNotExists *bool `json:"CreateFolderIfNotExists,omitempty" xml:"CreateFolderIfNotExists,omitempty"`
	// The cron expression for scheduled execution. This parameter corresponds to the Cron Expression setting in Scheduling > Scheduling Time for Data Studio tasks in the [DataWorks console](https://workbench.data.aliyun.com/console). After you configure Scheduling Cycle and Scheduled Time, DataWorks automatically generates a cron expression.
	//
	// Examples:
	//
	// 	- Scheduled at 05:30 every day: `00 30 05 	- 	- ?`
	//
	// 	- Scheduled at the 15th minute of every hour: `00 15 00-23/1 	- 	- ?`
	//
	// 	- Scheduled every 10 minutes: `00 00/10 	- 	- 	- ?`
	//
	// 	- Scheduled every 10 minutes between 08:00 and 17:00 every day: `00 00-59/10 8-17 	- 	- 	- ?`
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
	// 00 05 00 	- 	- ?
	CronExpress *string `json:"CronExpress,omitempty" xml:"CronExpress,omitempty"`
	// The type of scheduling cycle. Valid values: NOT_DAY (minute, hour) and DAY (day, week, month).
	//
	// This parameter corresponds to the Scheduling Cycle setting in Scheduling > Scheduling Time for Data Studio tasks in the [DataWorks console](https://workbench.data.aliyun.com/console).
	//
	// example:
	//
	// DAY
	CycleType *string `json:"CycleType,omitempty" xml:"CycleType,omitempty"`
	// The IDs of the nodes on which the current node depends. This parameter takes effect only when the DependentType parameter is set to USER_DEFINE. Separate multiple node IDs with commas (,).
	//
	// This parameter corresponds to the Other Nodes option in Properties > Dependencies > Cross-cycle Dependency (Original Previous-cycle Dependency) for data development nodes in the [DataWorks console](https://workbench.data.aliyun.com/console).
	//
	// example:
	//
	// abc
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
	// 	- USER_DEFINE_AND_SELF: Depends on both the current node and other nodes in the previous cycle.
	//
	// 	- CHILD_AND_SELF: Depends on both the current node and its child nodes in the previous cycle.
	//
	// example:
	//
	// NONE
	DependentType *string `json:"DependentType,omitempty" xml:"DependentType,omitempty"`
	// The timestamp (in milliseconds) when automatic scheduling stops.
	//
	// This parameter corresponds to the end time of Effective Period in Scheduling > Scheduling Time for Data Studio tasks in the [DataWorks console](https://workbench.data.aliyun.com/console).
	//
	// example:
	//
	// 1671694850000
	EndEffectDate *int64 `json:"EndEffectDate,omitempty" xml:"EndEffectDate,omitempty"`
	// The description of the file.
	//
	// example:
	//
	// test
	FileDescription *string `json:"FileDescription,omitempty" xml:"FileDescription,omitempty"`
	// The file path.
	//
	// example:
	//
	// Business_process/First_Business_Process/MaxCompute/Folder_1/Folder_2
	FileFolderPath *string `json:"FileFolderPath,omitempty" xml:"FileFolderPath,omitempty"`
	// The file name.
	//
	// This parameter is required.
	//
	// example:
	//
	// File name
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// The code type of the file. Different file types have different code. For more information, see [DataWorks node types](https://help.aliyun.com/document_detail/600169.html). You can call the [ListFileType](https://help.aliyun.com/document_detail/212428.html) operation to query the code types of files.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	FileType *int32 `json:"FileType,omitempty" xml:"FileType,omitempty"`
	// Specifies whether to inherit the dry-run status from the previous cycle. Valid values:
	//
	// 	- true: Inherit the dry-run status from the previous cycle.
	//
	// 	- false: Do not inherit the dry-run status from the previous cycle.
	//
	// example:
	//
	// false
	IgnoreParentSkipRunningProperty *bool `json:"IgnoreParentSkipRunningProperty,omitempty" xml:"IgnoreParentSkipRunningProperty,omitempty"`
	// The custom image ID.
	//
	// example:
	//
	// m-bp1h4b5a8ogkbll2f3tr
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The output names of the ancestor nodes on which the current node depends. Separate multiple output names with commas (,).
	//
	// This parameter corresponds to the Output Name of Ancestor Node setting in Properties > Dependencies for data development nodes in the [DataWorks console](https://workbench.data.aliyun.com/console).
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
	// The output context parameters of the node. The value must be in the JSON format. For more information about the parameter structure, see the OutputContextParameterList parameter in the response parameters of the [GetFile](https://help.aliyun.com/document_detail/173954.html) operation.
	//
	// This parameter corresponds to the Output Parameters setting in Properties > Input and Output Parameters for data development nodes in the [DataWorks console](https://workbench.data.aliyun.com/console).
	//
	// example:
	//
	// [{"Type": 1,"Value": "${bizdate}","ParameterName": "bizdate_param"}]
	OutputParameters *string `json:"OutputParameters,omitempty" xml:"OutputParameters,omitempty"`
	// The Alibaba Cloud account ID of the file owner. If this parameter is not specified, the Alibaba Cloud account ID of the caller is used by default.
	//
	// example:
	//
	// 1000000000001
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The scheduling parameters of the node. Separate multiple parameters with spaces.
	//
	// This parameter corresponds to the Scheduling Parameter setting in Properties for data development nodes in the [DataWorks console](https://workbench.data.aliyun.com/console). For more information, see [Scheduling parameters](https://help.aliyun.com/document_detail/137548.html).
	//
	// example:
	//
	// a=x b=y
	ParaValue *string `json:"ParaValue,omitempty" xml:"ParaValue,omitempty"`
	// The DataWorks workspace ID. To obtain the workspace ID, log on to the [DataWorks console](https://workbench.data.aliyun.com/console) and navigate to the workspace configuration page. You must configure either this parameter or the ProjectIdentifier parameter to determine the DataWorks workspace to which the operation is applied.
	//
	// example:
	//
	// 10000
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
	// example:
	//
	// ALL_ALLOWED
	RerunMode *string `json:"RerunMode,omitempty" xml:"RerunMode,omitempty"`
	// This parameter is deprecated.
	//
	// example:
	//
	// 375827434852437
	ResourceGroupId *int64 `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The resource group for the task published from the file. To obtain the ID, log on to the [DataWorks console](https://workbench.data.aliyun.com/console), navigate to the workspace configuration page, and click Resource Groups in the left-side navigation pane to view the IDs of resource groups bound to the current workspace.
	//
	// example:
	//
	// S_res_group_559_1613715566828
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
	// 1671608450000
	StartEffectDate *int64 `json:"StartEffectDate,omitempty" xml:"StartEffectDate,omitempty"`
	// Specifies whether to immediately run the node after the node is deployed.
	//
	// This parameter corresponds to the Start Method setting in Settings > Schedule in the right-side navigation pane on the configuration tab of EMR Spark Streaming and EMR Streaming SQL nodes in the [DataWorks console](https://workbench.data.aliyun.com/console).
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

func (s CreateFileRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateFileRequest) GoString() string {
	return s.String()
}

func (s *CreateFileRequest) GetAdvancedSettings() *string {
	return s.AdvancedSettings
}

func (s *CreateFileRequest) GetApplyScheduleImmediately() *bool {
	return s.ApplyScheduleImmediately
}

func (s *CreateFileRequest) GetAutoParsing() *bool {
	return s.AutoParsing
}

func (s *CreateFileRequest) GetAutoRerunIntervalMillis() *int32 {
	return s.AutoRerunIntervalMillis
}

func (s *CreateFileRequest) GetAutoRerunTimes() *int32 {
	return s.AutoRerunTimes
}

func (s *CreateFileRequest) GetConnectionName() *string {
	return s.ConnectionName
}

func (s *CreateFileRequest) GetContent() *string {
	return s.Content
}

func (s *CreateFileRequest) GetCreateFolderIfNotExists() *bool {
	return s.CreateFolderIfNotExists
}

func (s *CreateFileRequest) GetCronExpress() *string {
	return s.CronExpress
}

func (s *CreateFileRequest) GetCycleType() *string {
	return s.CycleType
}

func (s *CreateFileRequest) GetDependentNodeIdList() *string {
	return s.DependentNodeIdList
}

func (s *CreateFileRequest) GetDependentType() *string {
	return s.DependentType
}

func (s *CreateFileRequest) GetEndEffectDate() *int64 {
	return s.EndEffectDate
}

func (s *CreateFileRequest) GetFileDescription() *string {
	return s.FileDescription
}

func (s *CreateFileRequest) GetFileFolderPath() *string {
	return s.FileFolderPath
}

func (s *CreateFileRequest) GetFileName() *string {
	return s.FileName
}

func (s *CreateFileRequest) GetFileType() *int32 {
	return s.FileType
}

func (s *CreateFileRequest) GetIgnoreParentSkipRunningProperty() *bool {
	return s.IgnoreParentSkipRunningProperty
}

func (s *CreateFileRequest) GetImageId() *string {
	return s.ImageId
}

func (s *CreateFileRequest) GetInputList() *string {
	return s.InputList
}

func (s *CreateFileRequest) GetInputParameters() *string {
	return s.InputParameters
}

func (s *CreateFileRequest) GetOutputParameters() *string {
	return s.OutputParameters
}

func (s *CreateFileRequest) GetOwner() *string {
	return s.Owner
}

func (s *CreateFileRequest) GetParaValue() *string {
	return s.ParaValue
}

func (s *CreateFileRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *CreateFileRequest) GetProjectIdentifier() *string {
	return s.ProjectIdentifier
}

func (s *CreateFileRequest) GetRerunMode() *string {
	return s.RerunMode
}

func (s *CreateFileRequest) GetResourceGroupId() *int64 {
	return s.ResourceGroupId
}

func (s *CreateFileRequest) GetResourceGroupIdentifier() *string {
	return s.ResourceGroupIdentifier
}

func (s *CreateFileRequest) GetSchedulerType() *string {
	return s.SchedulerType
}

func (s *CreateFileRequest) GetStartEffectDate() *int64 {
	return s.StartEffectDate
}

func (s *CreateFileRequest) GetStartImmediately() *bool {
	return s.StartImmediately
}

func (s *CreateFileRequest) GetStop() *bool {
	return s.Stop
}

func (s *CreateFileRequest) GetTimeout() *int32 {
	return s.Timeout
}

func (s *CreateFileRequest) SetAdvancedSettings(v string) *CreateFileRequest {
	s.AdvancedSettings = &v
	return s
}

func (s *CreateFileRequest) SetApplyScheduleImmediately(v bool) *CreateFileRequest {
	s.ApplyScheduleImmediately = &v
	return s
}

func (s *CreateFileRequest) SetAutoParsing(v bool) *CreateFileRequest {
	s.AutoParsing = &v
	return s
}

func (s *CreateFileRequest) SetAutoRerunIntervalMillis(v int32) *CreateFileRequest {
	s.AutoRerunIntervalMillis = &v
	return s
}

func (s *CreateFileRequest) SetAutoRerunTimes(v int32) *CreateFileRequest {
	s.AutoRerunTimes = &v
	return s
}

func (s *CreateFileRequest) SetConnectionName(v string) *CreateFileRequest {
	s.ConnectionName = &v
	return s
}

func (s *CreateFileRequest) SetContent(v string) *CreateFileRequest {
	s.Content = &v
	return s
}

func (s *CreateFileRequest) SetCreateFolderIfNotExists(v bool) *CreateFileRequest {
	s.CreateFolderIfNotExists = &v
	return s
}

func (s *CreateFileRequest) SetCronExpress(v string) *CreateFileRequest {
	s.CronExpress = &v
	return s
}

func (s *CreateFileRequest) SetCycleType(v string) *CreateFileRequest {
	s.CycleType = &v
	return s
}

func (s *CreateFileRequest) SetDependentNodeIdList(v string) *CreateFileRequest {
	s.DependentNodeIdList = &v
	return s
}

func (s *CreateFileRequest) SetDependentType(v string) *CreateFileRequest {
	s.DependentType = &v
	return s
}

func (s *CreateFileRequest) SetEndEffectDate(v int64) *CreateFileRequest {
	s.EndEffectDate = &v
	return s
}

func (s *CreateFileRequest) SetFileDescription(v string) *CreateFileRequest {
	s.FileDescription = &v
	return s
}

func (s *CreateFileRequest) SetFileFolderPath(v string) *CreateFileRequest {
	s.FileFolderPath = &v
	return s
}

func (s *CreateFileRequest) SetFileName(v string) *CreateFileRequest {
	s.FileName = &v
	return s
}

func (s *CreateFileRequest) SetFileType(v int32) *CreateFileRequest {
	s.FileType = &v
	return s
}

func (s *CreateFileRequest) SetIgnoreParentSkipRunningProperty(v bool) *CreateFileRequest {
	s.IgnoreParentSkipRunningProperty = &v
	return s
}

func (s *CreateFileRequest) SetImageId(v string) *CreateFileRequest {
	s.ImageId = &v
	return s
}

func (s *CreateFileRequest) SetInputList(v string) *CreateFileRequest {
	s.InputList = &v
	return s
}

func (s *CreateFileRequest) SetInputParameters(v string) *CreateFileRequest {
	s.InputParameters = &v
	return s
}

func (s *CreateFileRequest) SetOutputParameters(v string) *CreateFileRequest {
	s.OutputParameters = &v
	return s
}

func (s *CreateFileRequest) SetOwner(v string) *CreateFileRequest {
	s.Owner = &v
	return s
}

func (s *CreateFileRequest) SetParaValue(v string) *CreateFileRequest {
	s.ParaValue = &v
	return s
}

func (s *CreateFileRequest) SetProjectId(v int64) *CreateFileRequest {
	s.ProjectId = &v
	return s
}

func (s *CreateFileRequest) SetProjectIdentifier(v string) *CreateFileRequest {
	s.ProjectIdentifier = &v
	return s
}

func (s *CreateFileRequest) SetRerunMode(v string) *CreateFileRequest {
	s.RerunMode = &v
	return s
}

func (s *CreateFileRequest) SetResourceGroupId(v int64) *CreateFileRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateFileRequest) SetResourceGroupIdentifier(v string) *CreateFileRequest {
	s.ResourceGroupIdentifier = &v
	return s
}

func (s *CreateFileRequest) SetSchedulerType(v string) *CreateFileRequest {
	s.SchedulerType = &v
	return s
}

func (s *CreateFileRequest) SetStartEffectDate(v int64) *CreateFileRequest {
	s.StartEffectDate = &v
	return s
}

func (s *CreateFileRequest) SetStartImmediately(v bool) *CreateFileRequest {
	s.StartImmediately = &v
	return s
}

func (s *CreateFileRequest) SetStop(v bool) *CreateFileRequest {
	s.Stop = &v
	return s
}

func (s *CreateFileRequest) SetTimeout(v int32) *CreateFileRequest {
	s.Timeout = &v
	return s
}

func (s *CreateFileRequest) Validate() error {
	return dara.Validate(s)
}
