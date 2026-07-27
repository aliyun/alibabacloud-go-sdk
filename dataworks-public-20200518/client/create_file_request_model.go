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
	SetOutputList(v string) *CreateFileRequest
	GetOutputList() *string
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
	// This parameter corresponds to the **Advanced Settings*	- in the right-side navigation pane of the editing page for EMR Spark Streaming and EMR Streaming SQL DataStudio tasks in the [DataWorks console](https://workbench.data.aliyun.com/console).
	//
	// Currently, only EMR Spark Streaming and EMR Streaming SQL tasks support this parameter, and the parameter value must be in JSON format.
	//
	// example:
	//
	// {"queue":"default","SPARK_CONF":"--conf spark.driver.memory=2g"}
	AdvancedSettings *string `json:"AdvancedSettings,omitempty" xml:"AdvancedSettings,omitempty"`
	// Specifies whether the scheduling configuration takes effect immediately after the file is published.
	//
	// example:
	//
	// true
	ApplyScheduleImmediately *bool `json:"ApplyScheduleImmediately,omitempty" xml:"ApplyScheduleImmediately,omitempty"`
	// Specifies whether to enable automatic parsing for the file. Valid values:
	//
	// - true: The file automatically parses code.
	//
	// - false: The file does not automatically parse code.
	//
	// This parameter corresponds to the **Code Parsing*	- setting when **Same Cycle*	- is selected under **Scheduling Configuration > Scheduling Dependency*	- of a DataStudio task in the [DataWorks console](https://workbench.data.aliyun.com/console).
	//
	// example:
	//
	// true
	AutoParsing *bool `json:"AutoParsing,omitempty" xml:"AutoParsing,omitempty"`
	// The interval between automatic reruns upon an error, in milliseconds. The maximum value is 1800000 milliseconds (30 minutes).
	//
	// This parameter corresponds to the **Rerun Interval*	- setting under **Scheduling Configuration > Time Properties > Auto Rerun upon Error*	- of a DataStudio task in the [DataWorks console](https://workbench.data.aliyun.com/console).
	//
	// The **Rerun Interval*	- in the console is in minutes. Make sure to convert the time unit when calling this operation.
	//
	// example:
	//
	// 120000
	AutoRerunIntervalMillis *int32 `json:"AutoRerunIntervalMillis,omitempty" xml:"AutoRerunIntervalMillis,omitempty"`
	// The number of automatic reruns allowed upon an error. The maximum value is 10.
	//
	// example:
	//
	// 3
	AutoRerunTimes *int32 `json:"AutoRerunTimes,omitempty" xml:"AutoRerunTimes,omitempty"`
	// The data source that the node connects to when the file is published as a node and executed.
	//
	// You can call the [UpdateDataSource](https://help.aliyun.com/document_detail/211432.html) operation to obtain the list of available data sources for the workspace.
	//
	// example:
	//
	// odps_source
	ConnectionName *string `json:"ConnectionName,omitempty" xml:"ConnectionName,omitempty"`
	// The code content of the file. Different code types (fileType) have different code formats. You can find the corresponding type of node in Operation Center, right-click the node, and then click View Code to view the specific code format.
	//
	// example:
	//
	// SHOW TABLES;
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// Specifies whether to automatically create the folder if the specified folder path (FileFolderPath) does not exist in the system. Valid values:
	//
	// - true: The folder is automatically created if it does not exist.
	//
	// - false: The invocation fails if the folder does not exist.
	//
	// example:
	//
	// false
	CreateFolderIfNotExists *bool `json:"CreateFolderIfNotExists,omitempty" xml:"CreateFolderIfNotExists,omitempty"`
	// The cron expression for periodic scheduling. This parameter corresponds to the **cron Expression*	- setting under **Scheduling Configuration > Time Property > cron Expression*	- of a DataStudio node in the [DataWorks console](https://workbench.data.aliyun.com/console). After you configure the **Scheduling Cycle*	- and **Timed Scheduling Time**, DataWorks automatically generates the corresponding cron expression.
	//
	// Examples:
	//
	// - Timed scheduling at 05:30 every day: `00 30 05 	- 	- ?`
	//
	// - Timed scheduling at the 15th minute of every hour: `00 15 00-23/1 	- 	- ?`
	//
	// - Schedule every 10 minutes: `00 00/10 	- 	- 	- ?`
	//
	// - Schedule every 10 minutes from 08:00 to 17:00 every day: `00 00-59/10 8-17 	- 	- 	- ?`
	//
	// - Timed scheduling at 00:20 on the 1st of every month: `00 20 00 1 	- ?`
	//
	// - Schedule every 3 months starting from 00:10 on January 1: `00 10 00 1 1-12/3 ?`
	//
	// - Timed scheduling at 00:05 every Tuesday and Friday: `00 05 00 	- 	- 2,5`
	//
	// Due to the rules of the DataWorks scheduling system, the cron expression has the following limits:
	//
	// - The minimum scheduling interval is 5 minutes.
	//
	// - The earliest scheduling time each day is 00:05.
	//
	// example:
	//
	// 00 05 00 	- 	- ?
	CronExpress *string `json:"CronExpress,omitempty" xml:"CronExpress,omitempty"`
	// The type of the scheduling cycle. Valid values: NOT_DAY (minute or hour) and DAY (day, week, or month).
	//
	// This parameter corresponds to the **Scheduling Cycle*	- setting under **Scheduling Configuration > Time Properties*	- of a DataStudio task in the [DataWorks console](https://workbench.data.aliyun.com/console).
	//
	// example:
	//
	// DAY
	CycleType *string `json:"CycleType,omitempty" xml:"CycleType,omitempty"`
	// The list of nodes that the current node depends on from the previous cycle.
	//
	// example:
	//
	// abc
	DependentNodeIdList *string `json:"DependentNodeIdList,omitempty" xml:"DependentNodeIdList,omitempty"`
	// The mode of cross-cycle dependency. Valid values:
	//
	// - SELF: The dependency is set to the current node.
	//
	// - CHILD: The dependency is set to first-level child nodes.
	//
	// - USER_DEFINE: The dependency is set to other nodes.
	//
	// - NONE: No dependency is selected, which means the node does not depend on the previous cycle.
	//
	// - USER_DEFINE_AND_SELF: The dependency is set to a combination of the current node and other nodes across cycles.
	//
	// - CHILD_AND_SELF: The dependency is set to a combination of first-level child nodes and the current node across cycles.
	//
	// example:
	//
	// NONE
	DependentType *string `json:"DependentType,omitempty" xml:"DependentType,omitempty"`
	// The timestamp in milliseconds when automatic scheduling stops.
	//
	// This parameter corresponds to the end time (in milliseconds) of the **Effective Date*	- setting under **Scheduling Configuration > Time Properties*	- of a DataStudio task in the [DataWorks console](https://workbench.data.aliyun.com/console).
	//
	// example:
	//
	// 1671694850000
	EndEffectDate *int64 `json:"EndEffectDate,omitempty" xml:"EndEffectDate,omitempty"`
	// The description of the file.
	//
	// example:
	//
	// This is a file description.
	FileDescription *string `json:"FileDescription,omitempty" xml:"FileDescription,omitempty"`
	// The path of the file.
	//
	// example:
	//
	// Business_process/First_Business_Process/MaxCompute/Folder_1/Folder_2
	FileFolderPath *string `json:"FileFolderPath,omitempty" xml:"FileFolderPath,omitempty"`
	// The name of the file.
	//
	// This parameter is required.
	//
	// example:
	//
	// File name
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// The code type of the file.
	//
	// Different file types have different codes. For more information, see [DataWorks nodes](https://help.aliyun.com/document_detail/600169.html).
	//
	// You can call the [ListFileType](https://help.aliyun.com/document_detail/212428.html) operation to query the code types of files.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	FileType *int32 `json:"FileType,omitempty" xml:"FileType,omitempty"`
	// Specifies whether to inherit the dry-run property from the previous cycle. Valid values:
	//
	// - true: Inherit the dry-run property from the previous cycle.
	//
	// - false: Do not inherit the dry-run property from the previous cycle.
	//
	// example:
	//
	// false
	IgnoreParentSkipRunningProperty *bool `json:"IgnoreParentSkipRunningProperty,omitempty" xml:"IgnoreParentSkipRunningProperty,omitempty"`
	// The ID of the custom image.
	//
	// example:
	//
	// m-bp1h4b5a8ogkbll2f3tr
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The output names of the upstream files that the file depends on. Separate multiple output names with commas (,).
	//
	// This parameter corresponds to the **Parent Node Output Name*	- setting when **Same Cycle*	- is selected under **Scheduling Configuration > Scheduling Dependency*	- of a DataStudio task in the [DataWorks console](https://workbench.data.aliyun.com/console).
	//
	// example:
	//
	// project_root,project.file1,project.001_out
	InputList *string `json:"InputList,omitempty" xml:"InputList,omitempty"`
	// The context input parameters of the node. The parameter value is in JSON format. For the fields included, see the InputContextParameterList parameter structure in the response of the [GetFile](https://help.aliyun.com/document_detail/173954.html) operation.
	//
	// This parameter corresponds to the **Input Parameters of This Node*	- setting under **Scheduling Configuration > Node Context*	- of a DataStudio task in the [DataWorks console](https://workbench.data.aliyun.com/console).
	//
	// example:
	//
	// [{"ValueSource": "project_001.first_node:bizdate_param","ParameterName": "bizdate_input"}]
	InputParameters *string `json:"InputParameters,omitempty" xml:"InputParameters,omitempty"`
	OutputList      *string `json:"OutputList,omitempty" xml:"OutputList,omitempty"`
	// The context output parameters of the node. The parameter value is in JSON format. For the fields included, see the OutputContextParameterList parameter structure in the response of the [GetFile](https://help.aliyun.com/document_detail/173954.html) operation.
	//
	// This parameter corresponds to the **Output Parameters of This Node*	- setting under **Scheduling Configuration > Node Context*	- of a DataStudio task in the [DataWorks console](https://workbench.data.aliyun.com/console).
	//
	// example:
	//
	// [{"Type": 1,"Value": "${bizdate}","ParameterName": "bizdate_param"}]
	OutputParameters *string `json:"OutputParameters,omitempty" xml:"OutputParameters,omitempty"`
	// The Alibaba Cloud user ID of the file owner. If this parameter is left empty, the Alibaba Cloud user ID of the caller is used by default.
	//
	// example:
	//
	// 1000000000001
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The scheduling parameters. Separate multiple parameters with spaces.
	//
	// This parameter corresponds to the **Parameters*	- setting under **Scheduling Configuration*	- of a DataStudio task in the [DataWorks console](https://workbench.data.aliyun.com/console). For more information, see [Scheduling parameters](https://help.aliyun.com/document_detail/137548.html).
	//
	// example:
	//
	// a=x b=y
	ParaValue *string `json:"ParaValue,omitempty" xml:"ParaValue,omitempty"`
	// The ID of the DataWorks workspace. You can log on to the [DataWorks console](https://workbench.data.aliyun.com/console) and go to the Workspace Settings page to obtain the workspace ID.
	//
	// You must specify either this parameter or ProjectIdentifier to determine the DataWorks workspace for this API call.
	//
	// example:
	//
	// 10000
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
	// This parameter corresponds to the **Rerun Property*	- setting under **Scheduling Configuration > Time Properties > Rerun Property*	- of a DataStudio task in the [DataWorks console](https://workbench.data.aliyun.com/console).
	//
	// example:
	//
	// ALL_ALLOWED
	RerunMode *string `json:"RerunMode,omitempty" xml:"RerunMode,omitempty"`
	// This parameter is deprecated. Do not use it.
	//
	// The schedule resource used when the file is published as a node and executed. This parameter corresponds to the **Scheduling Configuration > Resource Properties > Scheduling Resource Group*	- setting on the page. You can specify either this parameter or ResourceGroupIdentifier.
	//
	// You can call the [ListResourceGroups](https://help.aliyun.com/document_detail/173913.html) operation to obtain the list of available resource groups for the workspace. Set ResourceGroupType to 1 and use the ID field from the response.
	//
	// example:
	//
	// 375827434852437
	ResourceGroupId *int64 `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The schedule resource used when the file is published as a node and executed. You can call the [ListResourceGroups](https://help.aliyun.com/document_detail/173913.html) operation and use the **identifier*	- field to obtain the list of available resource groups for the workspace.
	//
	//
	// > Make sure that the resource group returned by the ListResourceGroups operation is bound to the workspace used to create the file. The resource group can be used in CreateFile only after it is bound.
	//
	// example:
	//
	// group_375827434852437
	ResourceGroupIdentifier *string `json:"ResourceGroupIdentifier,omitempty" xml:"ResourceGroupIdentifier,omitempty"`
	// The type of scheduling. Valid values:
	//
	// - NORMAL: The node is a normal scheduled node.
	//
	// - MANUAL: The node is a manual node that is not included in daily scheduling. This corresponds to nodes under manual workflows.
	//
	// - PAUSE: The node is a paused node.
	//
	// - SKIP: The node is a dry-run node that is included in daily scheduling but is immediately set to successful when triggered.
	//
	// example:
	//
	// NORMAL
	SchedulerType *string `json:"SchedulerType,omitempty" xml:"SchedulerType,omitempty"`
	// The timestamp in milliseconds when automatic scheduling starts.
	//
	// This parameter corresponds to the start time (in milliseconds) of the **Effective Date*	- setting under **Scheduling Configuration > Time Properties*	- of a DataStudio task in the [DataWorks console](https://workbench.data.aliyun.com/console).
	//
	// example:
	//
	// 1671608450000
	StartEffectDate *int64 `json:"StartEffectDate,omitempty" xml:"StartEffectDate,omitempty"`
	// Specifies whether to start the node immediately after it is published.
	//
	// This parameter corresponds to the **Start Mode*	- setting under **Configuration > Time Properties*	- in the right-side navigation pane of the editing page for EMR Spark Streaming and EMR Streaming SQL DataStudio tasks in the [DataWorks console](https://workbench.data.aliyun.com/console).
	//
	// example:
	//
	// true
	StartImmediately *bool `json:"StartImmediately,omitempty" xml:"StartImmediately,omitempty"`
	// Specifies whether to suspend scheduling. Valid values:
	//
	// - true: Suspend scheduling.
	//
	// - false: Do not suspend scheduling.
	//
	// This parameter corresponds to setting the **Scheduling Type*	- to **Suspend Scheduling*	- under **Scheduling Configuration > Time Properties*	- of a DataStudio task in the [DataWorks console](https://workbench.data.aliyun.com/console).
	//
	// example:
	//
	// false
	Stop *bool `json:"Stop,omitempty" xml:"Stop,omitempty"`
	// The timeout period defined in the scheduling configuration.
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

func (s *CreateFileRequest) GetOutputList() *string {
	return s.OutputList
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

func (s *CreateFileRequest) SetOutputList(v string) *CreateFileRequest {
	s.OutputList = &v
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
