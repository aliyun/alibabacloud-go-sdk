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
	// The advanced configurations of the node.
	//
	// This parameter is valid only for an EMR Spark Streaming node or an EMR Streaming SQL node. This parameter corresponds to the Advanced Settings tab of the node in the [DataWorks console](https://workbench.data.aliyun.com/console).
	//
	// This parameter is configured in the JSON format.
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
	// Specifies whether the automatic parsing feature is enabled for the file. Valid values:
	//
	// 	- true: The automatic parsing feature is enabled for the file.
	//
	// 	- false: The automatic parsing feature is not enabled for the file.
	//
	// This parameter corresponds to the Analyze Code parameter that is displayed after Same Cycle is selected in the Dependencies section of the Properties tab in the [DataWorks console](https://workbench.data.aliyun.com/console).
	//
	// example:
	//
	// true
	AutoParsing *bool `json:"AutoParsing,omitempty" xml:"AutoParsing,omitempty"`
	// The interval between two consecutive automatic reruns after an error occurs. Unit: milliseconds. Maximum value: 1800000 (30 minutes).
	//
	// This parameter corresponds to the Rerun Interval parameter that is displayed after the Auto Rerun upon Error check box is selected in the Schedule section of the Properties tab in the [DataWorks console](https://workbench.data.aliyun.com/console).
	//
	// The interval that you specify in the DataWorks console is measured in minutes. Pay attention to the conversion between the units of time when you call the operation.
	//
	// example:
	//
	// 120000
	AutoRerunIntervalMillis *int32 `json:"AutoRerunIntervalMillis,omitempty" xml:"AutoRerunIntervalMillis,omitempty"`
	// The number of automatic reruns that are allowed after an error occurs.
	//
	// example:
	//
	// 3
	AutoRerunTimes *int32 `json:"AutoRerunTimes,omitempty" xml:"AutoRerunTimes,omitempty"`
	// The name of the connected data source that is used to run the node. You can call the [ListDataSources](https://help.aliyun.com/document_detail/211431.html) operation to query the available data sources of the workspace.
	//
	// example:
	//
	// odps_first
	ConnectionName *string `json:"ConnectionName,omitempty" xml:"ConnectionName,omitempty"`
	// The code of the file. The code format varies based on the file type. To view the code format for a specific file type, go to Operation Center, right-click a node of the file type, and then select View Code.
	//
	// example:
	//
	// SELECT "1";
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The CRON expression that represents the periodic scheduling policy of the node. This parameter corresponds to the Cron Expression parameter in the Schedule section of the Properties tab in the [DataWorks console](https://workbench.data.aliyun.com/console). After you configure the Scheduling Cycle and Run At parameters in the DataWorks console, DataWorks automatically generates a value for the Cron Expression parameter.
	//
	// Examples:
	//
	// 	- CRON expression for a node that is scheduled to run at 05:30 every day: `00 30 05 	- 	- ?`
	//
	// 	- CRON expression for a node that is scheduled to run at the fifteenth minute of each hour: `00 15 	- 	- 	- ?`
	//
	// 	- CRON expression for a node that is scheduled to run every 10 minutes: `00 00/10 	- 	- 	- ?`
	//
	// 	- CRON expression for a node that is scheduled to run every 10 minutes from 08:00 to 17:00 every day: `00 00-59/10 8-23 	- 	- 	- ?`
	//
	// 	- CRON expression for a node that is scheduled to run at 00:20 on the first day of each month: `00 20 00 1 	- ?`
	//
	// 	- CRON expression for a node that is scheduled to run every three months starting from 00:10 on January 1: `00 10 00 1 1-12/3 ?`
	//
	// 	- CRON expression for a node that is scheduled to run at 00:05 every Tuesday and Friday: `00 05 00 	- 	- 2,5`
	//
	// The scheduling system of DataWorks imposes the following limits on CRON expressions:
	//
	// 	- A node can be scheduled to run at a minimum interval of 5 minutes.
	//
	// 	- A node can be scheduled to run at 00:05 every day at the earliest.
	//
	// example:
	//
	// 00 00-59/5 1-23 	- 	- ?
	CronExpress *string `json:"CronExpress,omitempty" xml:"CronExpress,omitempty"`
	// The type of the scheduling cycle of the node that corresponds to the file. Valid values: NOT_DAY and DAY. The value NOT_DAY indicates that the node is scheduled to run by minute or hour. The value DAY indicates that the node is scheduled to run by day, week, or month.
	//
	// This parameter corresponds to the Scheduling Cycle parameter in the Schedule section of the Properties tab in the [DataWorks console](https://workbench.data.aliyun.com/console).
	//
	// example:
	//
	// NOT_DAY
	CycleType *string `json:"CycleType,omitempty" xml:"CycleType,omitempty"`
	// The ID of the node on which the node corresponding to the file depends when the DependentType parameter is set to USER_DEFINE. Multiple IDs are separated by commas (,).
	//
	// The value of this parameter corresponds to the ID of the node that you specified after you select Previous Cycle and set Depend On to Other Nodes in the Dependencies section of the Properties tab in the [DataWorks console](https://workbench.data.aliyun.com/console).
	//
	// example:
	//
	// 5,10,15,20
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
	// example:
	//
	// USER_DEFINE
	DependentType *string `json:"DependentType,omitempty" xml:"DependentType,omitempty"`
	// The end time of automatic scheduling. Set this parameter to a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// This parameter corresponds to the end time specified for the Validity Period parameter in the Schedule section of the Properties tab in the [DataWorks console](https://workbench.data.aliyun.com/console).
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
	// The ID of the file. You can call the [ListFiles](https://help.aliyun.com/document_detail/173942.html) operation to obtain the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 100000001
	FileId *int64 `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// The name of the file. You can set the FileName parameter to a new value to change the file name.
	//
	// You can call the [ListFiles](https://help.aliyun.com/document_detail/173942.html) operation to query the ID of the file whose name you want to change. Then, you can set the FileId parameter to the ID and set the FileName parameter to a new value when you call the [UpdateFile](https://help.aliyun.com/document_detail/173951.html) operation.
	//
	// example:
	//
	// ods_user_info_d
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// Specifies whether to skip the dry-run property of the ancestor nodes of the node. This parameter corresponds to the Skip the dry-run property of the ancestor node parameter that is displayed after you configure the Depend On parameter in the Dependencies section of the Properties tab in the DataWorks console.
	//
	// example:
	//
	// true
	IgnoreParentSkipRunningProperty *bool `json:"IgnoreParentSkipRunningProperty,omitempty" xml:"IgnoreParentSkipRunningProperty,omitempty"`
	// The ID of the custom image.
	//
	// example:
	//
	// m-uf6d7npxk1hhek8ng0cb
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The output name of the parent file on which the current file depends. If you specify multiple output names, separate them with commas (,).
	//
	// This parameter corresponds to the Parent Nodes parameter that is displayed after you select Same Cycle in the Dependencies section of the Properties tab in the [DataWorks console](https://workbench.data.aliyun.com/console).
	//
	// >  You must configure this parameter when you call the CreateDISyncTask or UpdateFile operation to create a batch synchronization task.
	//
	// example:
	//
	// project_root,project.file1,project.001_out
	InputList *string `json:"InputList,omitempty" xml:"InputList,omitempty"`
	// The input parameters of the node. This parameter is configured in the JSON format. For more information about the input parameters, refer to the InputContextParameterList parameter in the Response parameters section of the [GetFile](https://help.aliyun.com/document_detail/173954.html) operation.
	//
	// This parameter corresponds to the Input Parameters table in the Input and Output Parameters section of the Properties tab in the [DataWorks console](https://workbench.data.aliyun.com/console).
	//
	// example:
	//
	// [{"ValueSource": "project_001.first_node:bizdate_param","ParameterName": "bizdate_input"}]
	InputParameters *string `json:"InputParameters,omitempty" xml:"InputParameters,omitempty"`
	// The output name of the current file.
	//
	// This parameter corresponds to the Output Name parameter in the Dependencies section of the Properties tab in the [DataWorks console](https://workbench.data.aliyun.com/console).
	//
	// example:
	//
	// dw_project.ods_user_info_d
	OutputList *string `json:"OutputList,omitempty" xml:"OutputList,omitempty"`
	// The output parameters of the node. This parameter is configured in the JSON format. For more information about the output parameters, refer to the OutputContextParameterList parameter in the Response parameters section of the [GetFile](https://help.aliyun.com/document_detail/173954.html) operation.
	//
	// This parameter corresponds to the Output Parameters table in the Input and Output Parameters section of the Properties tab in the [DataWorks console](https://workbench.data.aliyun.com/console).
	//
	// example:
	//
	// [{"Type": 1,"Value": "${bizdate}","ParameterName": "bizdate_param"}]
	OutputParameters *string `json:"OutputParameters,omitempty" xml:"OutputParameters,omitempty"`
	// The ID of the file owner.
	//
	// example:
	//
	// 18023848927592
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The scheduling parameters of the node.
	//
	// This parameter corresponds to the Parameters section of the Properties tab in the [DataWorks console](https://workbench.data.aliyun.com/console). For more information, see [Configure scheduling parameters](https://help.aliyun.com/document_detail/137548.html).
	//
	// example:
	//
	// x=a y=b z=c
	ParaValue *string `json:"ParaValue,omitempty" xml:"ParaValue,omitempty"`
	// The DataWorks workspace ID. You can log on to the [DataWorks console](https://workbench.data.aliyun.com/console) and go to the Workspace page to query the ID.
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
	// This parameter corresponds to the Rerun parameter in the Schedule section of the Properties tab in the [DataWorks console](https://workbench.data.aliyun.com/console).
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
	// The identifier of the resource group that is used to run the node. You can call the [ListResourceGroups](https://help.aliyun.com/document_detail/173913.html) operation to query the available resource groups in the workspace.
	//
	// example:
	//
	// default_group
	ResourceGroupIdentifier *string `json:"ResourceGroupIdentifier,omitempty" xml:"ResourceGroupIdentifier,omitempty"`
	// The scheduling type of the node. Valid values:
	//
	// 	- NORMAL: The node is an auto triggered node.
	//
	// 	- MANUAL: The node is a manually triggered node. Manually triggered nodes cannot be automatically triggered. They correspond to the nodes in the Manually Triggered Workflows pane.
	//
	// 	- PAUSE: The node is a paused node.
	//
	// 	- SKIP: The node is a dry-run node. Dry-run nodes are started as scheduled, but the system sets the status of the nodes to successful when it starts to run them.
	//
	// example:
	//
	// NORMAL
	SchedulerType *string `json:"SchedulerType,omitempty" xml:"SchedulerType,omitempty"`
	// The start time of automatic scheduling. Set this parameter to a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// This parameter corresponds to the Validity Period parameter in the Schedule section of the Properties tab in the [DataWorks console](https://workbench.data.aliyun.com/console).
	//
	// example:
	//
	// 936923400000
	StartEffectDate *int64 `json:"StartEffectDate,omitempty" xml:"StartEffectDate,omitempty"`
	// Specifies whether to immediately run a node after the node is deployed to the production environment. Valid values:
	//
	// 	- true: A node is immediately run after the node is deployed to the production environment.
	//
	// 	- false: A node is not immediately run after the node is deployed to the production environment.
	//
	// This parameter is valid only for an EMR Spark Streaming node or an EMR Streaming SQL node. This parameter corresponds to the Start Method parameter in the Schedule section of the Configure tab in the [DataWorks console](https://workbench.data.aliyun.com/console).
	//
	// example:
	//
	// true
	StartImmediately *bool `json:"StartImmediately,omitempty" xml:"StartImmediately,omitempty"`
	// Specifies whether to suspend the scheduling of the node. Valid values:
	//
	// 	- true: suspends the scheduling of the node.
	//
	// 	- false: does not suspend the scheduling of the node.
	//
	// This parameter corresponds to the Recurrence parameter in the Schedule section of the Properties tab in the [DataWorks console](https://workbench.data.aliyun.com/console).
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
