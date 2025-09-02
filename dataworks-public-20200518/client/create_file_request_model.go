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
	// The advanced configurations of the node.
	//
	// This parameter is valid only for an EMR Spark Streaming node or an EMR Streaming SQL node. This parameter corresponds to the Advanced Settings tab of the node in the [DataWorks console](https://workbench.data.aliyun.com/console).
	//
	// The value of this parameter must be in the JSON format.
	//
	// example:
	//
	// {"queue":"default","SPARK_CONF":"--conf spark.driver.memory=2g"}
	AdvancedSettings *string `json:"AdvancedSettings,omitempty" xml:"AdvancedSettings,omitempty"`
	// Specifies whether scheduling configurations immediately take effect after the node is deployed.
	//
	// example:
	//
	// true
	ApplyScheduleImmediately *bool `json:"ApplyScheduleImmediately,omitempty" xml:"ApplyScheduleImmediately,omitempty"`
	// Specifies whether to enable the automatic parsing feature for the file. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// This parameter corresponds to the Analyze Code parameter that is displayed after Same Cycle is selected in the Dependencies section of the Properties tab on the DataStudio page in the [DataWorks console](https://workbench.data.aliyun.com/console).
	//
	// example:
	//
	// true
	AutoParsing *bool `json:"AutoParsing,omitempty" xml:"AutoParsing,omitempty"`
	// The interval between automatic reruns after an error occurs. Unit: milliseconds. Maximum value: 1800000 (30 minutes).
	//
	// This parameter corresponds to the Rerun Interval parameter that is displayed after the Auto Rerun upon Error check box is selected in the Schedule section of the Properties tab in the [DataWorks console](https://workbench.data.aliyun.com/console).
	//
	// The interval that you specify in the DataWorks console is measured in minutes. Pay attention to the conversion between the units of time when you call the operation.
	//
	// example:
	//
	// 120000
	AutoRerunIntervalMillis *int32 `json:"AutoRerunIntervalMillis,omitempty" xml:"AutoRerunIntervalMillis,omitempty"`
	// The number of automatic reruns that are allowed after an error occurs. Maximum value: 10.
	//
	// example:
	//
	// 3
	AutoRerunTimes *int32 `json:"AutoRerunTimes,omitempty" xml:"AutoRerunTimes,omitempty"`
	// The name of the data source for which the node is run.
	//
	// You can call the [UpdateDataSource](https://help.aliyun.com/document_detail/211432.html) operation to query the available data sources in the workspace.
	//
	// example:
	//
	// odps_first
	ConnectionName *string `json:"ConnectionName,omitempty" xml:"ConnectionName,omitempty"`
	// The code for the file. The code format varies based on the file type. To view the code format for a specific file type, go to Operation Center, right-click a node of the file type, and then select View Code.
	//
	// example:
	//
	// SHOW TABLES;
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// Specifies whether to automatically create the directory that is specified by the FileFolderPath parameter if the directory does not exist. Valid values:
	//
	// 	- true: The system automatically creates the directory if the directory does not exist.
	//
	// 	- false: The system does not automatically create the directory if the directory does not exist. In this case, the call fails.
	//
	// example:
	//
	// false
	CreateFolderIfNotExists *bool `json:"CreateFolderIfNotExists,omitempty" xml:"CreateFolderIfNotExists,omitempty"`
	// The CRON expression that represents the periodic scheduling policy of the node. This parameter corresponds to the Cron Expression parameter in the Schedule section of the Properties tab on the DataStudio page in the [DataWorks console](https://workbench.data.aliyun.com/console). After you configure the Scheduling Cycle and Scheduled time parameters in the DataWorks console, DataWorks generates the value of the Cron Expression parameter.
	//
	// Examples:
	//
	// 	- CRON expression for a node that is scheduled to run at 05:30 every day: `00 30 05 	- 	- ?`
	//
	// 	- CRON expression for a node that is scheduled to run at the fifteenth minute of each hour: `00 15 00-23/1 	- 	- ?`
	//
	// 	- CRON expression for a node that is scheduled to run every 10 minutes: `00 00/10 	- 	- 	- ?`
	//
	// 	- CRON expression for a node that is scheduled to run every 10 minutes from 08:00 to 17:00 every day: `00 00-59/10 8-17 	- 	- 	- ?`
	//
	// 	- CRON expression for a node that is scheduled to run at 00:20 on the first day of each month: `00 20 00 1 	- ?`
	//
	// 	- CRON expression for a node that is scheduled to run every three months from 00:10 on January 1: `00 10 00 1 1-12/3 ?`
	//
	// 	- CRON expression for a node that is scheduled to run at 00:05 every Tuesday and Friday: `00 05 00 	- 	- 2,5`
	//
	// The scheduling system of DataWorks imposes the following limits on CRON expressions:
	//
	// 	- The minimum interval specified in a CRON expression to schedule a node is 5 minutes.
	//
	// 	- The earliest time specified in a CRON expression to schedule a node every day is 00:05.
	//
	// example:
	//
	// 00 05 00 	- 	- ?
	CronExpress *string `json:"CronExpress,omitempty" xml:"CronExpress,omitempty"`
	// The type of the scheduling cycle of the node that corresponds to the file. Valid values: NOT_DAY and DAY. The value NOT_DAY indicates that the node is scheduled to run by minute or hour. The value DAY indicates that the node is scheduled to run by day, week, or month.
	//
	// This parameter corresponds to the Scheduling Cycle parameter in the Schedule section of the Properties tab on the DataStudio page in the [DataWorks console](https://workbench.data.aliyun.com/console).
	//
	// example:
	//
	// DAY
	CycleType *string `json:"CycleType,omitempty" xml:"CycleType,omitempty"`
	// The IDs of the nodes that generate instances in the previous cycle on which the current node depends.
	//
	// example:
	//
	// abc
	DependentNodeIdList *string `json:"DependentNodeIdList,omitempty" xml:"DependentNodeIdList,omitempty"`
	// The type of the cross-cycle scheduling dependency of the node. Valid values:
	//
	// 	- SELF: The instance generated for the node in the current cycle depends on the instance generated for the node in the previous cycle.
	//
	// 	- CHILD: The instance generated for the node in the current cycle depends on the instances generated for the descendant nodes at the nearest level of the node in the previous cycle.
	//
	// 	- USER_DEFINE: The instance generated for the node in the current cycle depends on the instances generated for one or more specified nodes in the previous cycle.
	//
	// 	- NONE: No cross-cycle scheduling dependency type is selected for the node.
	//
	// 	- USER_DEFINE_AND_SELF: The instance generated for the node in the current cycle depends on the instance generated for the node in the previous cycle and the instances generated for one or more specified nodes in the previous cycle.
	//
	// 	- CHILD_AND_SELF: The instance generated for the node in the current cycle depends on the instances generated for the descendant nodes at the nearest level of the node in the previous cycle and the instance generated for the node in the previous cycle.
	//
	// example:
	//
	// NONE
	DependentType *string `json:"DependentType,omitempty" xml:"DependentType,omitempty"`
	// The end time of automatic scheduling. Set the value to a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// This parameter corresponds to the Validity Period parameter in the Schedule section of the Properties tab in the [DataWorks console](https://workbench.data.aliyun.com/console).
	//
	// example:
	//
	// 1671694850000
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
	// The type of the code for the file. The code for files varies based on the file type. For more information, see [DataWorks nodes](https://help.aliyun.com/document_detail/600169.html).
	//
	// You can call the [ListFileType](https://help.aliyun.com/document_detail/212428.html) operation to query the type of the code for the file.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	FileType *int32 `json:"FileType,omitempty" xml:"FileType,omitempty"`
	// Whether to use the last cycle empty run attribute. The values are as follows:
	//
	// - true: The empty run attribute of the previous cycle is used.
	//
	// - false: The empty run attribute of the previous cycle is not used.
	//
	// example:
	//
	// false
	IgnoreParentSkipRunningProperty *bool `json:"IgnoreParentSkipRunningProperty,omitempty" xml:"IgnoreParentSkipRunningProperty,omitempty"`
	// Custom image ID
	//
	// example:
	//
	// m-bp1h4b5a8ogkbll2f3tr
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The output name of the parent file on which the current file depends. If you specify multiple output names, separate them with commas (,).
	//
	// This parameter corresponds to the Output Name parameter under Parent Nodes in the Dependencies section of the Properties tab in the [DataWorks console](https://workbench.data.aliyun.com/console).
	//
	// example:
	//
	// project_root,project.file1,project.001_out
	InputList *string `json:"InputList,omitempty" xml:"InputList,omitempty"`
	// The input parameters of the node. The value of this parameter must be in the JSON format. For more information about the input parameters, see the InputContextParameterList parameter in the Response parameters section of the [GetFile](https://help.aliyun.com/document_detail/173954.html) operation.
	//
	// This parameter corresponds to the Input Parameters table in the Input and Output Parameters section of the Properties tab in the [DataWorks console](https://workbench.data.aliyun.com/console).
	//
	// example:
	//
	// [{"ValueSource": "project_001.first_node:bizdate_param","ParameterName": "bizdate_input"}]
	InputParameters *string `json:"InputParameters,omitempty" xml:"InputParameters,omitempty"`
	// The output parameters of the node. The value of this parameter must be in the JSON format. For more information about the output parameters, see the OutputContextParameterList parameter in the Response parameters section of the [GetFile](https://help.aliyun.com/document_detail/173954.html) operation.
	//
	// This parameter corresponds to the Output Parameters table in the Input and Output Parameters section of the Properties tab in the [DataWorks console](https://workbench.data.aliyun.com/console).
	//
	// example:
	//
	// [{"Type": 1,"Value": "${bizdate}","ParameterName": "bizdate_param"}]
	OutputParameters *string `json:"OutputParameters,omitempty" xml:"OutputParameters,omitempty"`
	// The ID of the Alibaba Cloud account used by the file owner. If this parameter is not configured, the ID of the Alibaba Cloud account of the user who calls the operation is used.
	//
	// example:
	//
	// 1000000000001
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The scheduling parameters of the node. Separate multiple parameters with spaces.
	//
	// This parameter corresponds to the Parameters section of the Properties tab in the [DataWorks console](https://workbench.data.aliyun.com/console). For more information about the configurations of the scheduling parameters, see [Configure scheduling parameters](https://help.aliyun.com/document_detail/137548.html).
	//
	// example:
	//
	// a=x b=y
	ParaValue *string `json:"ParaValue,omitempty" xml:"ParaValue,omitempty"`
	// The ID of the DataWorks workspace. You can log on to the [DataWorks console](https://workbench.data.aliyun.com/console) and go to the Workspace Management page to obtain the workspace ID.
	//
	// You must configure this parameter or the ProjectIdentifier parameter to determine the DataWorks workspace to which the operation is applied.
	//
	// example:
	//
	// 10000
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The name of the DataWorks workspace. You can log on to the [DataWorks console](https://workbench.data.aliyun.com/console) and go to the Workspace Management page to obtain the workspace name.
	//
	// You must configure this parameter or the ProjectId parameter to determine the DataWorks workspace to which the operation is applied.
	//
	// example:
	//
	// dw_project
	ProjectIdentifier *string `json:"ProjectIdentifier,omitempty" xml:"ProjectIdentifier,omitempty"`
	// Specifies whether the node that corresponds to the file can be rerun. Valid values:
	//
	// 	- ALL_ALLOWED: The node can be rerun regardless of whether it is successfully run or fails to run.
	//
	// 	- FAILURE_ALLOWED: The node can be rerun only after it fails to run.
	//
	// 	- ALL_DENIED: The node cannot be rerun regardless of whether it is successfully run or fails to run.
	//
	// This parameter corresponds to the Rerun parameter in the Schedule section of the Properties tab on the DataStudio page in the [DataWorks console](https://workbench.data.aliyun.com/console).
	//
	// example:
	//
	// ALL_ALLOWED
	RerunMode *string `json:"RerunMode,omitempty" xml:"RerunMode,omitempty"`
	// This parameter is deprecated. Do not use this parameter.
	//
	// The identifier of the resource group that is used to run the node. This parameter corresponds to the Resource Group parameter in the Resource Group section of the Properties tab in the DataWorks console. You must configure one of the ResourceGroupId and ResourceGroupIdentifier parameters to determine the resource group that is used to run the node.
	//
	// You can call the [ListResourceGroups](https://help.aliyun.com/document_detail/173913.html) operation to query the available resource groups in the workspace. When you call the operation, set the ResourceGroupType parameter to 1. The response parameter Id indicates the ID of an available resource group.
	//
	// example:
	//
	// 375827434852437
	ResourceGroupId *int64 `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The identifier of the resource group that is used to run the node. You can call the [ListResourceGroups](https://help.aliyun.com/document_detail/173913.html) operation to query the available resource groups in the workspace. The **Identifier*	- parameter in the response of the operation indicates the identifier of an available resource group.
	//
	// >  You must make sure that the available resource groups in the response of the ListResourceGroups operation are associated with the workspace for which you want to create a file by calling the CreateFile operation.
	//
	// example:
	//
	// group_375827434852437
	ResourceGroupIdentifier *string `json:"ResourceGroupIdentifier,omitempty" xml:"ResourceGroupIdentifier,omitempty"`
	// The scheduling type of the node. Valid values:
	//
	// 	- NORMAL: The node is an auto triggered node.
	//
	// 	- MANUAL: The node is a manually triggered node. Manually triggered nodes cannot be automatically triggered. They correspond to the nodes in the Manually Triggered Workflows pane.
	//
	// 	- PAUSE: The node is a paused node.
	//
	// 	- SKIP: The node is a dry-run node. Dry-run nodes are started as scheduled, but the system sets the status of the nodes to successful when it starts to run them
	//
	// example:
	//
	// NORMAL
	SchedulerType *string `json:"SchedulerType,omitempty" xml:"SchedulerType,omitempty"`
	// The start time of automatic scheduling. Set the value to a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// Configuring this parameter is equivalent to specifying a start time for the Validity Period parameter in the Schedule section of the Properties tab on the DataStudio page in the [DataWorks console](https://workbench.data.aliyun.com/console).
	//
	// example:
	//
	// 1671608450000
	StartEffectDate *int64 `json:"StartEffectDate,omitempty" xml:"StartEffectDate,omitempty"`
	// Specifies whether to immediately run a node after the node is deployed.
	//
	// This parameter is valid only for an EMR Spark Streaming node or an EMR Streaming SQL node. This parameter corresponds to the Start Method parameter in the Schedule section of the Configure tab in the [DataWorks console](https://workbench.data.aliyun.com/console).
	//
	// example:
	//
	// true
	StartImmediately *bool `json:"StartImmediately,omitempty" xml:"StartImmediately,omitempty"`
	// Specifies whether to suspend the scheduling of the node. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// This parameter corresponds to the Recurrence parameter in the Schedule section of the Properties tab on the DataStudio page in the [DataWorks console](https://workbench.data.aliyun.com/console).
	//
	// example:
	//
	// false
	Stop *bool `json:"Stop,omitempty" xml:"Stop,omitempty"`
	// The timeout period.
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
