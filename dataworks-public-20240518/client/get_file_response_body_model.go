// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFileResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetFileResponseBodyData) *GetFileResponseBody
	GetData() *GetFileResponseBodyData
	SetErrorCode(v string) *GetFileResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetFileResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v int32) *GetFileResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *GetFileResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetFileResponseBody
	GetSuccess() *bool
}

type GetFileResponseBody struct {
	// The details of the file.
	Data *GetFileResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code.
	//
	// example:
	//
	// Invalid.Tenant.ConnectionNotExists
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// The connection does not exist.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0000-ABCD-EFG****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetFileResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetFileResponseBody) GoString() string {
	return s.String()
}

func (s *GetFileResponseBody) GetData() *GetFileResponseBodyData {
	return s.Data
}

func (s *GetFileResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetFileResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetFileResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetFileResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetFileResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetFileResponseBody) SetData(v *GetFileResponseBodyData) *GetFileResponseBody {
	s.Data = v
	return s
}

func (s *GetFileResponseBody) SetErrorCode(v string) *GetFileResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetFileResponseBody) SetErrorMessage(v string) *GetFileResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetFileResponseBody) SetHttpStatusCode(v int32) *GetFileResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetFileResponseBody) SetRequestId(v string) *GetFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetFileResponseBody) SetSuccess(v bool) *GetFileResponseBody {
	s.Success = &v
	return s
}

func (s *GetFileResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetFileResponseBodyData struct {
	// The basic information about the file.
	File *GetFileResponseBodyDataFile `json:"File,omitempty" xml:"File,omitempty" type:"Struct"`
	// The scheduling configurations of the file.
	NodeConfiguration *GetFileResponseBodyDataNodeConfiguration `json:"NodeConfiguration,omitempty" xml:"NodeConfiguration,omitempty" type:"Struct"`
	// The download URL of the resource.
	ResourceDownloadLink *GetFileResponseBodyDataResourceDownloadLink `json:"ResourceDownloadLink,omitempty" xml:"ResourceDownloadLink,omitempty" type:"Struct"`
}

func (s GetFileResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetFileResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetFileResponseBodyData) GetFile() *GetFileResponseBodyDataFile {
	return s.File
}

func (s *GetFileResponseBodyData) GetNodeConfiguration() *GetFileResponseBodyDataNodeConfiguration {
	return s.NodeConfiguration
}

func (s *GetFileResponseBodyData) GetResourceDownloadLink() *GetFileResponseBodyDataResourceDownloadLink {
	return s.ResourceDownloadLink
}

func (s *GetFileResponseBodyData) SetFile(v *GetFileResponseBodyDataFile) *GetFileResponseBodyData {
	s.File = v
	return s
}

func (s *GetFileResponseBodyData) SetNodeConfiguration(v *GetFileResponseBodyDataNodeConfiguration) *GetFileResponseBodyData {
	s.NodeConfiguration = v
	return s
}

func (s *GetFileResponseBodyData) SetResourceDownloadLink(v *GetFileResponseBodyDataResourceDownloadLink) *GetFileResponseBodyData {
	s.ResourceDownloadLink = v
	return s
}

func (s *GetFileResponseBodyData) Validate() error {
	if s.File != nil {
		if err := s.File.Validate(); err != nil {
			return err
		}
	}
	if s.NodeConfiguration != nil {
		if err := s.NodeConfiguration.Validate(); err != nil {
			return err
		}
	}
	if s.ResourceDownloadLink != nil {
		if err := s.ResourceDownloadLink.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetFileResponseBodyDataFile struct {
	// The advanced configurations of the node.
	//
	// This parameter is valid for an EMR node. This parameter corresponds to the Advanced Settings tab in the right-side navigation pane on the configuration tab of the node in the [DataWorks console](https://workbench.data.aliyun.com/console).
	//
	// >  You cannot configure advanced parameters for EMR Shell nodes.
	//
	// For information about the advanced parameters of each type of EMR node, see [Develop EMR tasks](https://help.aliyun.com/document_detail/473077.html).
	//
	// example:
	//
	// {\\"priority\\":\\"1\\",\\"ENABLE_SPARKSQL_JDBC\\":false,\\"FLOW_SKIP_SQL_ANALYZE\\":false,\\"queue\\":\\"default\\"}
	AdvancedSettings *string `json:"AdvancedSettings,omitempty" xml:"AdvancedSettings,omitempty"`
	// Indicates whether the automatic parsing feature is enabled for the file. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// This parameter corresponds to the Automatic Parsing From Code Before Node Committing parameter that is displayed after you select Same Cycle in the Dependencies section of the Properties tab on the DataStudio page in the [DataWorks console](https://workbench.data.aliyun.com/console).
	//
	// example:
	//
	// true
	AutoParsing *bool `json:"AutoParsing,omitempty" xml:"AutoParsing,omitempty"`
	// The ID of the workflow to which the file belongs. This parameter is deprecated and replaced by the BusinessId parameter.
	//
	// example:
	//
	// 1000001
	BizId *int64 `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// The ID of the workflow to which the file belongs.
	//
	// example:
	//
	// 1000001
	BusinessId *int64 `json:"BusinessId,omitempty" xml:"BusinessId,omitempty"`
	// Indicates whether the latest code in the file is committed. Valid values: 0 and 1. The value 0 indicates that the latest code in the file is not committed. The value 1 indicates that the latest code in the file is committed.
	//
	// example:
	//
	// 0
	CommitStatus *int32 `json:"CommitStatus,omitempty" xml:"CommitStatus,omitempty"`
	// The name of the data source that is used to run the node that corresponds to the file.
	//
	// example:
	//
	// odps_source
	ConnectionName *string `json:"ConnectionName,omitempty" xml:"ConnectionName,omitempty"`
	// The code in the file.
	//
	// example:
	//
	// SHOW TABLES;
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The time when the file was created. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1593879116000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the Alibaba Cloud account used to create the file.
	//
	// example:
	//
	// 424732****
	CreateUser *string `json:"CreateUser,omitempty" xml:"CreateUser,omitempty"`
	// The latest version number of the file.
	//
	// example:
	//
	// 3
	CurrentVersion *int32 `json:"CurrentVersion,omitempty" xml:"CurrentVersion,omitempty"`
	// The status of the file. Valid values:
	//
	// 	- NORMAL: The file is not deleted.
	//
	// 	- RECYCLE_BIN: The file is stored in the recycle bin.
	//
	// 	- DELETED: The file is deleted.
	//
	// example:
	//
	// RECYCLE
	DeletedStatus *string `json:"DeletedStatus,omitempty" xml:"DeletedStatus,omitempty"`
	// The description of the file.
	//
	// example:
	//
	// My first DataWorks file
	FileDescription *string `json:"FileDescription,omitempty" xml:"FileDescription,omitempty"`
	// The ID of the folder to which the file belongs.
	//
	// example:
	//
	// 2735c2****
	FileFolderId *string `json:"FileFolderId,omitempty" xml:"FileFolderId,omitempty"`
	// The file ID.
	//
	// example:
	//
	// 100000001
	FileId *int64 `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// The name of the file.
	//
	// example:
	//
	// ods_user_info_d
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// The type of the code for the file. The code for files varies based on the file type. For more information, see [DataWorks nodes](https://help.aliyun.com/document_detail/600169.html).
	//
	// example:
	//
	// 10
	FileType *int32 `json:"FileType,omitempty" xml:"FileType,omitempty"`
	// Indicates whether the resource file needs to be uploaded to MaxCompute. This parameter is returned only if the file is a MaxCompute resource file.
	//
	// example:
	//
	// true
	IsMaxCompute *bool `json:"IsMaxCompute,omitempty" xml:"IsMaxCompute,omitempty"`
	// The time when the file was last modified. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1593879116000
	LastEditTime *int64 `json:"LastEditTime,omitempty" xml:"LastEditTime,omitempty"`
	// The ID of the Alibaba Cloud account used to last modify the file.
	//
	// example:
	//
	// 424732****
	LastEditUser *string `json:"LastEditUser,omitempty" xml:"LastEditUser,omitempty"`
	// The ID of the auto triggered node that is generated in the scheduling system after the file is committed.
	//
	// example:
	//
	// 300001
	NodeId *int64 `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The ID of the Alibaba Cloud account used by the file owner.
	//
	// example:
	//
	// 7775674356****
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The ID of the node group file to which the current file belongs. This parameter is returned only if the current file is an inner file of the node group file.
	//
	// example:
	//
	// -1
	ParentId *int64 `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
	// The module to which the file belongs. Valid values:
	//
	// 	- NORMAL: The file is used for DataStudio.
	//
	// 	- MANUAL: The file is used for a manually triggered node.
	//
	// 	- MANUAL_BIZ: The file is used for a manually triggered workflow.
	//
	// 	- SKIP: The file is used for a dry-run node in DataStudio.
	//
	// 	- ADHOCQUERY: The file is used for an ad hoc query.
	//
	// 	- COMPONENT: The file is used for a script template.
	//
	// example:
	//
	// NORMAL
	UseType *string `json:"UseType,omitempty" xml:"UseType,omitempty"`
}

func (s GetFileResponseBodyDataFile) String() string {
	return dara.Prettify(s)
}

func (s GetFileResponseBodyDataFile) GoString() string {
	return s.String()
}

func (s *GetFileResponseBodyDataFile) GetAdvancedSettings() *string {
	return s.AdvancedSettings
}

func (s *GetFileResponseBodyDataFile) GetAutoParsing() *bool {
	return s.AutoParsing
}

func (s *GetFileResponseBodyDataFile) GetBizId() *int64 {
	return s.BizId
}

func (s *GetFileResponseBodyDataFile) GetBusinessId() *int64 {
	return s.BusinessId
}

func (s *GetFileResponseBodyDataFile) GetCommitStatus() *int32 {
	return s.CommitStatus
}

func (s *GetFileResponseBodyDataFile) GetConnectionName() *string {
	return s.ConnectionName
}

func (s *GetFileResponseBodyDataFile) GetContent() *string {
	return s.Content
}

func (s *GetFileResponseBodyDataFile) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *GetFileResponseBodyDataFile) GetCreateUser() *string {
	return s.CreateUser
}

func (s *GetFileResponseBodyDataFile) GetCurrentVersion() *int32 {
	return s.CurrentVersion
}

func (s *GetFileResponseBodyDataFile) GetDeletedStatus() *string {
	return s.DeletedStatus
}

func (s *GetFileResponseBodyDataFile) GetFileDescription() *string {
	return s.FileDescription
}

func (s *GetFileResponseBodyDataFile) GetFileFolderId() *string {
	return s.FileFolderId
}

func (s *GetFileResponseBodyDataFile) GetFileId() *int64 {
	return s.FileId
}

func (s *GetFileResponseBodyDataFile) GetFileName() *string {
	return s.FileName
}

func (s *GetFileResponseBodyDataFile) GetFileType() *int32 {
	return s.FileType
}

func (s *GetFileResponseBodyDataFile) GetIsMaxCompute() *bool {
	return s.IsMaxCompute
}

func (s *GetFileResponseBodyDataFile) GetLastEditTime() *int64 {
	return s.LastEditTime
}

func (s *GetFileResponseBodyDataFile) GetLastEditUser() *string {
	return s.LastEditUser
}

func (s *GetFileResponseBodyDataFile) GetNodeId() *int64 {
	return s.NodeId
}

func (s *GetFileResponseBodyDataFile) GetOwner() *string {
	return s.Owner
}

func (s *GetFileResponseBodyDataFile) GetParentId() *int64 {
	return s.ParentId
}

func (s *GetFileResponseBodyDataFile) GetUseType() *string {
	return s.UseType
}

func (s *GetFileResponseBodyDataFile) SetAdvancedSettings(v string) *GetFileResponseBodyDataFile {
	s.AdvancedSettings = &v
	return s
}

func (s *GetFileResponseBodyDataFile) SetAutoParsing(v bool) *GetFileResponseBodyDataFile {
	s.AutoParsing = &v
	return s
}

func (s *GetFileResponseBodyDataFile) SetBizId(v int64) *GetFileResponseBodyDataFile {
	s.BizId = &v
	return s
}

func (s *GetFileResponseBodyDataFile) SetBusinessId(v int64) *GetFileResponseBodyDataFile {
	s.BusinessId = &v
	return s
}

func (s *GetFileResponseBodyDataFile) SetCommitStatus(v int32) *GetFileResponseBodyDataFile {
	s.CommitStatus = &v
	return s
}

func (s *GetFileResponseBodyDataFile) SetConnectionName(v string) *GetFileResponseBodyDataFile {
	s.ConnectionName = &v
	return s
}

func (s *GetFileResponseBodyDataFile) SetContent(v string) *GetFileResponseBodyDataFile {
	s.Content = &v
	return s
}

func (s *GetFileResponseBodyDataFile) SetCreateTime(v int64) *GetFileResponseBodyDataFile {
	s.CreateTime = &v
	return s
}

func (s *GetFileResponseBodyDataFile) SetCreateUser(v string) *GetFileResponseBodyDataFile {
	s.CreateUser = &v
	return s
}

func (s *GetFileResponseBodyDataFile) SetCurrentVersion(v int32) *GetFileResponseBodyDataFile {
	s.CurrentVersion = &v
	return s
}

func (s *GetFileResponseBodyDataFile) SetDeletedStatus(v string) *GetFileResponseBodyDataFile {
	s.DeletedStatus = &v
	return s
}

func (s *GetFileResponseBodyDataFile) SetFileDescription(v string) *GetFileResponseBodyDataFile {
	s.FileDescription = &v
	return s
}

func (s *GetFileResponseBodyDataFile) SetFileFolderId(v string) *GetFileResponseBodyDataFile {
	s.FileFolderId = &v
	return s
}

func (s *GetFileResponseBodyDataFile) SetFileId(v int64) *GetFileResponseBodyDataFile {
	s.FileId = &v
	return s
}

func (s *GetFileResponseBodyDataFile) SetFileName(v string) *GetFileResponseBodyDataFile {
	s.FileName = &v
	return s
}

func (s *GetFileResponseBodyDataFile) SetFileType(v int32) *GetFileResponseBodyDataFile {
	s.FileType = &v
	return s
}

func (s *GetFileResponseBodyDataFile) SetIsMaxCompute(v bool) *GetFileResponseBodyDataFile {
	s.IsMaxCompute = &v
	return s
}

func (s *GetFileResponseBodyDataFile) SetLastEditTime(v int64) *GetFileResponseBodyDataFile {
	s.LastEditTime = &v
	return s
}

func (s *GetFileResponseBodyDataFile) SetLastEditUser(v string) *GetFileResponseBodyDataFile {
	s.LastEditUser = &v
	return s
}

func (s *GetFileResponseBodyDataFile) SetNodeId(v int64) *GetFileResponseBodyDataFile {
	s.NodeId = &v
	return s
}

func (s *GetFileResponseBodyDataFile) SetOwner(v string) *GetFileResponseBodyDataFile {
	s.Owner = &v
	return s
}

func (s *GetFileResponseBodyDataFile) SetParentId(v int64) *GetFileResponseBodyDataFile {
	s.ParentId = &v
	return s
}

func (s *GetFileResponseBodyDataFile) SetUseType(v string) *GetFileResponseBodyDataFile {
	s.UseType = &v
	return s
}

func (s *GetFileResponseBodyDataFile) Validate() error {
	return dara.Validate(s)
}

type GetFileResponseBodyDataNodeConfiguration struct {
	// Indicates whether scheduling configurations immediately take effect after the deployment.
	//
	// example:
	//
	// true
	ApplyScheduleImmediately *string `json:"ApplyScheduleImmediately,omitempty" xml:"ApplyScheduleImmediately,omitempty"`
	// The interval between automatic reruns after an error occurs. Unit: milliseconds.
	//
	// This parameter corresponds to the Rerun interval parameter that is displayed after the Auto Rerun upon Failure check box is selected in the Schedule section of the Properties tab on the DataStudio page in the [DataWorks console](https://workbench.data.aliyun.com/console). The interval that you specify in the DataWorks console is measured in minutes. Pay attention to the conversion between the units of time when you call the operation.
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
	// The cron expression that represents the periodic scheduling policy of the node.
	//
	// example:
	//
	// 00 05 00 	- 	- ?
	CronExpress *string `json:"CronExpress,omitempty" xml:"CronExpress,omitempty"`
	// The type of the scheduling cycle. Valid values: NOT_DAY and DAY. The value NOT_DAY indicates that the node is scheduled to run by minute or hour. The value DAY indicates that the node is scheduled to run by day, week, or month.
	//
	// This parameter corresponds to the Scheduling Cycle parameter in the Schedule section of the Properties tab on the DataStudio page in the [DataWorks console](https://workbench.data.aliyun.com/console).
	//
	// example:
	//
	// DAY
	CycleType *string `json:"CycleType,omitempty" xml:"CycleType,omitempty"`
	// The ID of the node on which the node that corresponds to the file depends when the DependentType parameter is set to USER_DEFINE. Multiple IDs are separated by commas (,).
	//
	// The value of this parameter is equivalent to the ID of the node that you specified after you select Previous Cycle and set Depend On to Other Nodes in the Dependencies section of the Properties tab on the DataStudio page in the [DataWorks console](https://workbench.data.aliyun.com/console).
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
	// The end of the time range for automatic scheduling. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// Configuring this parameter is equivalent to specifying an end time for the Validity Period parameter in the Schedule section of the Properties tab on the DataStudio page in the [DataWorks console](https://workbench.data.aliyun.com/console).
	//
	// example:
	//
	// 4155787800000
	EndEffectDate *int64 `json:"EndEffectDate,omitempty" xml:"EndEffectDate,omitempty"`
	// Indicates whether the dry-run property of the ancestor nodes of the node is skipped. This parameter corresponds to the Skip the dry-run property of the ancestor node parameter that is displayed after you configure the Depend On parameter in the Dependencies section of the Properties tab on the DataStudio page in the DataWorks console.
	//
	// example:
	//
	// true
	IgnoreParentSkipRunningProperty *string `json:"IgnoreParentSkipRunningProperty,omitempty" xml:"IgnoreParentSkipRunningProperty,omitempty"`
	// The custom image ID.
	//
	// example:
	//
	// m-bp1h4b5a8ogkbll2f3tr
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The output information about the parent files on which the current file depends.
	InputList []*GetFileResponseBodyDataNodeConfigurationInputList `json:"InputList,omitempty" xml:"InputList,omitempty" type:"Repeated"`
	// The input parameters of the node.
	InputParameters []*GetFileResponseBodyDataNodeConfigurationInputParameters `json:"InputParameters,omitempty" xml:"InputParameters,omitempty" type:"Repeated"`
	// The output information about the current file.
	OutputList []*GetFileResponseBodyDataNodeConfigurationOutputList `json:"OutputList,omitempty" xml:"OutputList,omitempty" type:"Repeated"`
	// The output parameters of the node.
	OutputParameters []*GetFileResponseBodyDataNodeConfigurationOutputParameters `json:"OutputParameters,omitempty" xml:"OutputParameters,omitempty" type:"Repeated"`
	// The scheduling parameters of the node.
	//
	// This parameter corresponds to the Scheduling Parameter section of the Properties tab on the DataStudio page in the [DataWorks console](https://workbench.data.aliyun.com/console). For more information about the configurations of scheduling parameters, see [Configure scheduling parameters](https://help.aliyun.com/document_detail/137548.html).
	//
	// example:
	//
	// a=x b=y
	ParaValue *string `json:"ParaValue,omitempty" xml:"ParaValue,omitempty"`
	// Indicates whether the node that corresponds to the file can be rerun. Valid values:
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
	// The ID of the resource group that is used to run the node that corresponds to the file. You can call the [ListResourceGroups](https://help.aliyun.com/document_detail/173913.html) operation to query the available resource groups in the workspace.
	//
	// example:
	//
	// 375827434852437
	ResourceGroupId *int64 `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
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
	// The beginning of the time range for automatic scheduling. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// Configuring this parameter is equivalent to specifying a start time for the Validity Period parameter in the Schedule section of the Properties tab on the DataStudio page in the [DataWorks console](https://workbench.data.aliyun.com/console).
	//
	// example:
	//
	// 936923400000
	StartEffectDate *int64 `json:"StartEffectDate,omitempty" xml:"StartEffectDate,omitempty"`
	// Indicates whether a node is immediately run after the node is deployed to the production environment.
	//
	// This parameter is valid only for an EMR Spark Streaming node or an EMR Streaming SQL node. This parameter corresponds to the Start Method parameter in the Schedule section of the Configure tab on the DataStudio page in the [DataWorks console](https://workbench.data.aliyun.com/console).
	//
	// example:
	//
	// true
	StartImmediately *bool `json:"StartImmediately,omitempty" xml:"StartImmediately,omitempty"`
	// Indicates whether the scheduling for the node is suspended Valid values:
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

func (s GetFileResponseBodyDataNodeConfiguration) String() string {
	return dara.Prettify(s)
}

func (s GetFileResponseBodyDataNodeConfiguration) GoString() string {
	return s.String()
}

func (s *GetFileResponseBodyDataNodeConfiguration) GetApplyScheduleImmediately() *string {
	return s.ApplyScheduleImmediately
}

func (s *GetFileResponseBodyDataNodeConfiguration) GetAutoRerunIntervalMillis() *int32 {
	return s.AutoRerunIntervalMillis
}

func (s *GetFileResponseBodyDataNodeConfiguration) GetAutoRerunTimes() *int32 {
	return s.AutoRerunTimes
}

func (s *GetFileResponseBodyDataNodeConfiguration) GetCronExpress() *string {
	return s.CronExpress
}

func (s *GetFileResponseBodyDataNodeConfiguration) GetCycleType() *string {
	return s.CycleType
}

func (s *GetFileResponseBodyDataNodeConfiguration) GetDependentNodeIdList() *string {
	return s.DependentNodeIdList
}

func (s *GetFileResponseBodyDataNodeConfiguration) GetDependentType() *string {
	return s.DependentType
}

func (s *GetFileResponseBodyDataNodeConfiguration) GetEndEffectDate() *int64 {
	return s.EndEffectDate
}

func (s *GetFileResponseBodyDataNodeConfiguration) GetIgnoreParentSkipRunningProperty() *string {
	return s.IgnoreParentSkipRunningProperty
}

func (s *GetFileResponseBodyDataNodeConfiguration) GetImageId() *string {
	return s.ImageId
}

func (s *GetFileResponseBodyDataNodeConfiguration) GetInputList() []*GetFileResponseBodyDataNodeConfigurationInputList {
	return s.InputList
}

func (s *GetFileResponseBodyDataNodeConfiguration) GetInputParameters() []*GetFileResponseBodyDataNodeConfigurationInputParameters {
	return s.InputParameters
}

func (s *GetFileResponseBodyDataNodeConfiguration) GetOutputList() []*GetFileResponseBodyDataNodeConfigurationOutputList {
	return s.OutputList
}

func (s *GetFileResponseBodyDataNodeConfiguration) GetOutputParameters() []*GetFileResponseBodyDataNodeConfigurationOutputParameters {
	return s.OutputParameters
}

func (s *GetFileResponseBodyDataNodeConfiguration) GetParaValue() *string {
	return s.ParaValue
}

func (s *GetFileResponseBodyDataNodeConfiguration) GetRerunMode() *string {
	return s.RerunMode
}

func (s *GetFileResponseBodyDataNodeConfiguration) GetResourceGroupId() *int64 {
	return s.ResourceGroupId
}

func (s *GetFileResponseBodyDataNodeConfiguration) GetSchedulerType() *string {
	return s.SchedulerType
}

func (s *GetFileResponseBodyDataNodeConfiguration) GetStartEffectDate() *int64 {
	return s.StartEffectDate
}

func (s *GetFileResponseBodyDataNodeConfiguration) GetStartImmediately() *bool {
	return s.StartImmediately
}

func (s *GetFileResponseBodyDataNodeConfiguration) GetStop() *bool {
	return s.Stop
}

func (s *GetFileResponseBodyDataNodeConfiguration) GetTimeout() *int32 {
	return s.Timeout
}

func (s *GetFileResponseBodyDataNodeConfiguration) SetApplyScheduleImmediately(v string) *GetFileResponseBodyDataNodeConfiguration {
	s.ApplyScheduleImmediately = &v
	return s
}

func (s *GetFileResponseBodyDataNodeConfiguration) SetAutoRerunIntervalMillis(v int32) *GetFileResponseBodyDataNodeConfiguration {
	s.AutoRerunIntervalMillis = &v
	return s
}

func (s *GetFileResponseBodyDataNodeConfiguration) SetAutoRerunTimes(v int32) *GetFileResponseBodyDataNodeConfiguration {
	s.AutoRerunTimes = &v
	return s
}

func (s *GetFileResponseBodyDataNodeConfiguration) SetCronExpress(v string) *GetFileResponseBodyDataNodeConfiguration {
	s.CronExpress = &v
	return s
}

func (s *GetFileResponseBodyDataNodeConfiguration) SetCycleType(v string) *GetFileResponseBodyDataNodeConfiguration {
	s.CycleType = &v
	return s
}

func (s *GetFileResponseBodyDataNodeConfiguration) SetDependentNodeIdList(v string) *GetFileResponseBodyDataNodeConfiguration {
	s.DependentNodeIdList = &v
	return s
}

func (s *GetFileResponseBodyDataNodeConfiguration) SetDependentType(v string) *GetFileResponseBodyDataNodeConfiguration {
	s.DependentType = &v
	return s
}

func (s *GetFileResponseBodyDataNodeConfiguration) SetEndEffectDate(v int64) *GetFileResponseBodyDataNodeConfiguration {
	s.EndEffectDate = &v
	return s
}

func (s *GetFileResponseBodyDataNodeConfiguration) SetIgnoreParentSkipRunningProperty(v string) *GetFileResponseBodyDataNodeConfiguration {
	s.IgnoreParentSkipRunningProperty = &v
	return s
}

func (s *GetFileResponseBodyDataNodeConfiguration) SetImageId(v string) *GetFileResponseBodyDataNodeConfiguration {
	s.ImageId = &v
	return s
}

func (s *GetFileResponseBodyDataNodeConfiguration) SetInputList(v []*GetFileResponseBodyDataNodeConfigurationInputList) *GetFileResponseBodyDataNodeConfiguration {
	s.InputList = v
	return s
}

func (s *GetFileResponseBodyDataNodeConfiguration) SetInputParameters(v []*GetFileResponseBodyDataNodeConfigurationInputParameters) *GetFileResponseBodyDataNodeConfiguration {
	s.InputParameters = v
	return s
}

func (s *GetFileResponseBodyDataNodeConfiguration) SetOutputList(v []*GetFileResponseBodyDataNodeConfigurationOutputList) *GetFileResponseBodyDataNodeConfiguration {
	s.OutputList = v
	return s
}

func (s *GetFileResponseBodyDataNodeConfiguration) SetOutputParameters(v []*GetFileResponseBodyDataNodeConfigurationOutputParameters) *GetFileResponseBodyDataNodeConfiguration {
	s.OutputParameters = v
	return s
}

func (s *GetFileResponseBodyDataNodeConfiguration) SetParaValue(v string) *GetFileResponseBodyDataNodeConfiguration {
	s.ParaValue = &v
	return s
}

func (s *GetFileResponseBodyDataNodeConfiguration) SetRerunMode(v string) *GetFileResponseBodyDataNodeConfiguration {
	s.RerunMode = &v
	return s
}

func (s *GetFileResponseBodyDataNodeConfiguration) SetResourceGroupId(v int64) *GetFileResponseBodyDataNodeConfiguration {
	s.ResourceGroupId = &v
	return s
}

func (s *GetFileResponseBodyDataNodeConfiguration) SetSchedulerType(v string) *GetFileResponseBodyDataNodeConfiguration {
	s.SchedulerType = &v
	return s
}

func (s *GetFileResponseBodyDataNodeConfiguration) SetStartEffectDate(v int64) *GetFileResponseBodyDataNodeConfiguration {
	s.StartEffectDate = &v
	return s
}

func (s *GetFileResponseBodyDataNodeConfiguration) SetStartImmediately(v bool) *GetFileResponseBodyDataNodeConfiguration {
	s.StartImmediately = &v
	return s
}

func (s *GetFileResponseBodyDataNodeConfiguration) SetStop(v bool) *GetFileResponseBodyDataNodeConfiguration {
	s.Stop = &v
	return s
}

func (s *GetFileResponseBodyDataNodeConfiguration) SetTimeout(v int32) *GetFileResponseBodyDataNodeConfiguration {
	s.Timeout = &v
	return s
}

func (s *GetFileResponseBodyDataNodeConfiguration) Validate() error {
	if s.InputList != nil {
		for _, item := range s.InputList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.InputParameters != nil {
		for _, item := range s.InputParameters {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.OutputList != nil {
		for _, item := range s.OutputList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.OutputParameters != nil {
		for _, item := range s.OutputParameters {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetFileResponseBodyDataNodeConfigurationInputList struct {
	// The output name of the parent file on which the current file depends.
	//
	// This parameter corresponds to the Output Name of Ancestor Node parameter under Parent Nodes after Same Cycle is selected in the Dependencies section of the Properties tab on the DataStudio page in the [DataWorks console](https://workbench.data.aliyun.com/console).
	//
	// example:
	//
	// project.001_out
	Input *string `json:"Input,omitempty" xml:"Input,omitempty"`
	// The mode of the configuration file dependency. Valid values:
	//
	// 	- MANUAL: Scheduling dependencies are manually configured.
	//
	// 	- AUTO: Scheduling dependencies are automatically parsed.
	//
	// example:
	//
	// MANUAL
	ParseType *string `json:"ParseType,omitempty" xml:"ParseType,omitempty"`
}

func (s GetFileResponseBodyDataNodeConfigurationInputList) String() string {
	return dara.Prettify(s)
}

func (s GetFileResponseBodyDataNodeConfigurationInputList) GoString() string {
	return s.String()
}

func (s *GetFileResponseBodyDataNodeConfigurationInputList) GetInput() *string {
	return s.Input
}

func (s *GetFileResponseBodyDataNodeConfigurationInputList) GetParseType() *string {
	return s.ParseType
}

func (s *GetFileResponseBodyDataNodeConfigurationInputList) SetInput(v string) *GetFileResponseBodyDataNodeConfigurationInputList {
	s.Input = &v
	return s
}

func (s *GetFileResponseBodyDataNodeConfigurationInputList) SetParseType(v string) *GetFileResponseBodyDataNodeConfigurationInputList {
	s.ParseType = &v
	return s
}

func (s *GetFileResponseBodyDataNodeConfigurationInputList) Validate() error {
	return dara.Validate(s)
}

type GetFileResponseBodyDataNodeConfigurationInputParameters struct {
	// The name of the input parameter of the node. In the code, you can use the ${...} method to reference the input parameter of the node.
	//
	// This parameter corresponds to the Parameter Name parameter in the Input Parameters table in the Input and Output Parameters section of the Properties tab on the DataStudio page in the [DataWorks console](https://workbench.data.aliyun.com/console).
	//
	// example:
	//
	// input
	ParameterName *string `json:"ParameterName,omitempty" xml:"ParameterName,omitempty"`
	// The value source of the input parameter of the node.
	//
	// This parameter corresponds to the Value Source parameter in the Input Parameters table in the Input and Output Parameters section of the Properties tab on the DataStudio page in the [DataWorks console](https://workbench.data.aliyun.com/console).
	//
	// example:
	//
	// project_001.parent_node:outputs
	ValueSource *string `json:"ValueSource,omitempty" xml:"ValueSource,omitempty"`
}

func (s GetFileResponseBodyDataNodeConfigurationInputParameters) String() string {
	return dara.Prettify(s)
}

func (s GetFileResponseBodyDataNodeConfigurationInputParameters) GoString() string {
	return s.String()
}

func (s *GetFileResponseBodyDataNodeConfigurationInputParameters) GetParameterName() *string {
	return s.ParameterName
}

func (s *GetFileResponseBodyDataNodeConfigurationInputParameters) GetValueSource() *string {
	return s.ValueSource
}

func (s *GetFileResponseBodyDataNodeConfigurationInputParameters) SetParameterName(v string) *GetFileResponseBodyDataNodeConfigurationInputParameters {
	s.ParameterName = &v
	return s
}

func (s *GetFileResponseBodyDataNodeConfigurationInputParameters) SetValueSource(v string) *GetFileResponseBodyDataNodeConfigurationInputParameters {
	s.ValueSource = &v
	return s
}

func (s *GetFileResponseBodyDataNodeConfigurationInputParameters) Validate() error {
	return dara.Validate(s)
}

type GetFileResponseBodyDataNodeConfigurationOutputList struct {
	// The output name of the current file.
	//
	// This parameter corresponds to the Output Name parameter under Output after Same Cycle is selected in the Dependencies section of the Properties tab on the DataStudio page in the [DataWorks console](https://workbench.data.aliyun.com/console).
	//
	// example:
	//
	// dw_project.002_out
	Output *string `json:"Output,omitempty" xml:"Output,omitempty"`
	// The output table name of the current file.
	//
	// This parameter corresponds to the Output Table Name parameter under Output after Same Cycle is selected in the Dependencies section of the Properties tab on the DataStudio page in the [DataWorks console](https://workbench.data.aliyun.com/console).
	//
	// example:
	//
	// ods_user_info_d
	RefTableName *string `json:"RefTableName,omitempty" xml:"RefTableName,omitempty"`
}

func (s GetFileResponseBodyDataNodeConfigurationOutputList) String() string {
	return dara.Prettify(s)
}

func (s GetFileResponseBodyDataNodeConfigurationOutputList) GoString() string {
	return s.String()
}

func (s *GetFileResponseBodyDataNodeConfigurationOutputList) GetOutput() *string {
	return s.Output
}

func (s *GetFileResponseBodyDataNodeConfigurationOutputList) GetRefTableName() *string {
	return s.RefTableName
}

func (s *GetFileResponseBodyDataNodeConfigurationOutputList) SetOutput(v string) *GetFileResponseBodyDataNodeConfigurationOutputList {
	s.Output = &v
	return s
}

func (s *GetFileResponseBodyDataNodeConfigurationOutputList) SetRefTableName(v string) *GetFileResponseBodyDataNodeConfigurationOutputList {
	s.RefTableName = &v
	return s
}

func (s *GetFileResponseBodyDataNodeConfigurationOutputList) Validate() error {
	return dara.Validate(s)
}

type GetFileResponseBodyDataNodeConfigurationOutputParameters struct {
	// The description of the output parameter of the node.
	//
	// example:
	//
	// It\\"s a context output parameter.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the output parameter of the node.
	//
	// This parameter corresponds to the Parameter Name parameter in the Output Parameters table in the Input and Output Parameters section of the Properties tab on the DataStudio page in the [DataWorks console](https://workbench.data.aliyun.com/console).
	//
	// example:
	//
	// output
	ParameterName *string `json:"ParameterName,omitempty" xml:"ParameterName,omitempty"`
	// The type of the output parameter of the node. Valid values:
	//
	// 	- 1: indicates a constant.
	//
	// 	- 2: indicates a variable.
	//
	// 	- 3: indicates a pass-through variable.
	//
	// This parameter corresponds to the Type parameter in the Output Parameters table in the Input and Output Parameters section of the Properties tab on the DataStudio page in the [DataWorks console](https://workbench.data.aliyun.com/console).
	//
	// example:
	//
	// 1
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The value of the output parameter of the node.
	//
	// This parameter corresponds to the Value parameter in the Output Parameters table in the Input and Output Parameters section of the Properties tab on the DataStudio page in the [DataWorks console](https://workbench.data.aliyun.com/console).
	//
	// example:
	//
	// ${bizdate}
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetFileResponseBodyDataNodeConfigurationOutputParameters) String() string {
	return dara.Prettify(s)
}

func (s GetFileResponseBodyDataNodeConfigurationOutputParameters) GoString() string {
	return s.String()
}

func (s *GetFileResponseBodyDataNodeConfigurationOutputParameters) GetDescription() *string {
	return s.Description
}

func (s *GetFileResponseBodyDataNodeConfigurationOutputParameters) GetParameterName() *string {
	return s.ParameterName
}

func (s *GetFileResponseBodyDataNodeConfigurationOutputParameters) GetType() *string {
	return s.Type
}

func (s *GetFileResponseBodyDataNodeConfigurationOutputParameters) GetValue() *string {
	return s.Value
}

func (s *GetFileResponseBodyDataNodeConfigurationOutputParameters) SetDescription(v string) *GetFileResponseBodyDataNodeConfigurationOutputParameters {
	s.Description = &v
	return s
}

func (s *GetFileResponseBodyDataNodeConfigurationOutputParameters) SetParameterName(v string) *GetFileResponseBodyDataNodeConfigurationOutputParameters {
	s.ParameterName = &v
	return s
}

func (s *GetFileResponseBodyDataNodeConfigurationOutputParameters) SetType(v string) *GetFileResponseBodyDataNodeConfigurationOutputParameters {
	s.Type = &v
	return s
}

func (s *GetFileResponseBodyDataNodeConfigurationOutputParameters) SetValue(v string) *GetFileResponseBodyDataNodeConfigurationOutputParameters {
	s.Value = &v
	return s
}

func (s *GetFileResponseBodyDataNodeConfigurationOutputParameters) Validate() error {
	return dara.Validate(s)
}

type GetFileResponseBodyDataResourceDownloadLink struct {
	// The download URL of the resource.
	//
	// example:
	//
	// http://xx
	DownloadLink *string `json:"downloadLink,omitempty" xml:"downloadLink,omitempty"`
}

func (s GetFileResponseBodyDataResourceDownloadLink) String() string {
	return dara.Prettify(s)
}

func (s GetFileResponseBodyDataResourceDownloadLink) GoString() string {
	return s.String()
}

func (s *GetFileResponseBodyDataResourceDownloadLink) GetDownloadLink() *string {
	return s.DownloadLink
}

func (s *GetFileResponseBodyDataResourceDownloadLink) SetDownloadLink(v string) *GetFileResponseBodyDataResourceDownloadLink {
	s.DownloadLink = &v
	return s
}

func (s *GetFileResponseBodyDataResourceDownloadLink) Validate() error {
	return dara.Validate(s)
}
