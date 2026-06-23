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
	// Details of the file.
	Data *GetFileResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Error code.
	//
	// example:
	//
	// Invalid.Tenant.ConnectionNotExists
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// Error message.
	//
	// example:
	//
	// The connection does not exist.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Request ID. Used for troubleshooting when a fault occurs.
	//
	// example:
	//
	// 0000-ABCD-EFG****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the invocation succeeded. Valid values:
	//
	// - true: The invocation succeeded.
	//
	// - false: Failed to invoke.
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
	// Basic information about the file.
	File *GetFileResponseBodyDataFile `json:"File,omitempty" xml:"File,omitempty" type:"Struct"`
	// The schedule configuration of the file.
	NodeConfiguration *GetFileResponseBodyDataNodeConfiguration `json:"NodeConfiguration,omitempty" xml:"NodeConfiguration,omitempty" type:"Struct"`
	// Resource download link.
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
	// Advanced configuration of the job.
	//
	// This parameter corresponds to "Advanced Settings" in the right-side navigation bar on the editing page of an EMR Data Development job in the [DataWorks console](https://workbench.data.aliyun.com/console).
	//
	// > Currently, EMR Shell jobs do not support advanced parameters.
	//
	// For details about advanced parameters for different EMR job types, see [EMR Job Development](https://help.aliyun.com/document_detail/473077.html).
	//
	// example:
	//
	// {\\"priority\\":\\"1\\",\\"ENABLE_SPARKSQL_JDBC\\":false,\\"FLOW_SKIP_SQL_ANALYZE\\":false,\\"queue\\":\\"default\\"}
	AdvancedSettings *string `json:"AdvancedSettings,omitempty" xml:"AdvancedSettings,omitempty"`
	// Indicates whether automatic parsing is enabled for the file. Valid values:
	//
	// - true: The code in the file is automatically parsed.
	//
	// - false: The code in the file is not automatically parsed.
	//
	// This parameter corresponds to the "Code Parsing" option in the DataWorks console (https\\://workbench.data.aliyun.com/console) when you select "Same Cycle" under Schedule Configuration > Schedule Dependency for a Data Development job.
	//
	// example:
	//
	// true
	AutoParsing *bool `json:"AutoParsing,omitempty" xml:"AutoParsing,omitempty"`
	// The ID of the Business Process to which the file belongs. This field is deprecated. Use the BusinessId field instead.
	//
	// example:
	//
	// 1000001
	BizId *int64 `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// The Business Process ID of the file.
	//
	// example:
	//
	// 1000001
	BusinessId *int64 `json:"BusinessId,omitempty" xml:"BusinessId,omitempty"`
	// The current commit status of the file. Valid values:
	//
	// - 0: The latest code has not been submitted.
	//
	// - 1: The latest code has been submitted.
	//
	// example:
	//
	// 0
	CommitStatus *int32 `json:"CommitStatus,omitempty" xml:"CommitStatus,omitempty"`
	// The name of the data source used when executing the job corresponding to the file.
	//
	// example:
	//
	// odps_source
	ConnectionName *string `json:"ConnectionName,omitempty" xml:"ConnectionName,omitempty"`
	// The code of the file.
	//
	// example:
	//
	// SHOW TABLES;
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// UNIX timestamp when the file was created, in milliseconds.
	//
	// example:
	//
	// 1593879116000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The Alibaba Cloud User ID of the file creator.
	//
	// example:
	//
	// 424732****
	CreateUser *string `json:"CreateUser,omitempty" xml:"CreateUser,omitempty"`
	// Version number of the latest submitted version of the file.
	//
	// example:
	//
	// 3
	CurrentVersion *int32 `json:"CurrentVersion,omitempty" xml:"CurrentVersion,omitempty"`
	// The deletion status of the file. Valid values:
	//
	// - NORMAL: Not deleted.
	//
	// - RECYCLE_BIN: In the recycle bin.
	//
	// - DELETED: Deleted.
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
	// The ID of the file.
	//
	// example:
	//
	// 100000001
	FileId *int64 `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// Name of the file.
	//
	// example:
	//
	// ods_user_info_d
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// The code type of the file. Different file types use different code. For more information, see [DataWorks Edge Zone Collection](https://help.aliyun.com/document_detail/600169.html).
	//
	// example:
	//
	// 10
	FileType *int32 `json:"FileType,omitempty" xml:"FileType,omitempty"`
	// Indicates whether the resource file needs to be uploaded to MaxCompute.
	//
	// Configure this parameter only when the file is a MaxCompute resource file.
	//
	// example:
	//
	// true
	IsMaxCompute *bool `json:"IsMaxCompute,omitempty" xml:"IsMaxCompute,omitempty"`
	// The UNIX timestamp of the most recent edit to the file, in milliseconds.
	//
	// example:
	//
	// 1593879116000
	LastEditTime *int64 `json:"LastEditTime,omitempty" xml:"LastEditTime,omitempty"`
	// The Alibaba Cloud User ID of the user who last edited the file.
	//
	// example:
	//
	// 424732****
	LastEditUser *string `json:"LastEditUser,omitempty" xml:"LastEditUser,omitempty"`
	// The ID of the scheduling task generated in the CDN mapping system after the file is submitted.
	//
	// example:
	//
	// 300001
	NodeId *int64 `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// Alibaba Cloud User ID of the file owner.
	//
	// example:
	//
	// 7775674356****
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// If the current file is an internal file of a composite edge zone file, this field identifies the ID of the corresponding composite edge zone file.
	//
	// example:
	//
	// -1
	ParentId *int64 `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
	// The function module to which the file belongs. Valid values:
	//
	// - NORMAL: Data Development.
	//
	// - MANUAL: One-time task.
	//
	// - MANUAL_BIZ: Manually triggered workflow.
	//
	// - SKIP: Dry-run scheduling in Data Development.
	//
	// - ADHOCQUERY: Ad-hoc query.
	//
	// - COMPONENT: Widget Management.
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
	// Whether to apply the schedule configuration immediately after publishing.
	//
	// example:
	//
	// true
	ApplyScheduleImmediately *string `json:"ApplyScheduleImmediately,omitempty" xml:"ApplyScheduleImmediately,omitempty"`
	// The time interval between automatic reruns after an error, in milliseconds.
	//
	// This parameter corresponds to the "Rerun Interval" setting under "Schedule Configuration > Time Properties > Auto Rerun on Error" for a Data Development job in the [DataWorks console](https://workbench.data.aliyun.com/console).<br>
	//
	// Note that the time unit for "Rerun Interval" in the console is minutes; convert the time accordingly when invoking the API.
	//
	// example:
	//
	// 120000
	AutoRerunIntervalMillis *int32 `json:"AutoRerunIntervalMillis,omitempty" xml:"AutoRerunIntervalMillis,omitempty"`
	// The number of automatic reruns after an error.
	//
	// example:
	//
	// 3
	AutoRerunTimes *int32 `json:"AutoRerunTimes,omitempty" xml:"AutoRerunTimes,omitempty"`
	// The Cron Expression for timed scheduling of the file.
	//
	// example:
	//
	// 00 05 00 	- 	- ?
	CronExpress *string `json:"CronExpress,omitempty" xml:"CronExpress,omitempty"`
	// The type of recurrence, including NOT_DAY (minute, hour) and DAY (day, week, month).
	//
	// This parameter corresponds to "Schedule Configuration > Time Properties > Recurrence" for a Data Development job in the [DataWorks console](https://workbench.data.aliyun.com/console).
	//
	// example:
	//
	// DAY
	CycleType *string `json:"CycleType,omitempty" xml:"CycleType,omitempty"`
	// When the DependentType parameter is set to USER_DEFINE, this parameter specifies the IDs of the nodes on which the current file depends. Separate multiple node IDs with commas (,).
	//
	// This parameter corresponds to the configuration when, in the [DataWorks console](https://workbench.data.aliyun.com/console), the "Schedule Configuration > Schedule Dependency" of a Data Development job is set to "Previous Cycle" and the dependency option is set to "Other Nodes".
	//
	// example:
	//
	// 5,10,15,20
	DependentNodeIdList *string `json:"DependentNodeIdList,omitempty" xml:"DependentNodeIdList,omitempty"`
	// The method of depending on the previous cycle. Valid values:
	//
	// - SELF: The dependency is the current node itself.
	//
	// - CHILD: The dependency is direct child nodes.
	//
	// - USER_DEFINE: The dependency is other specified nodes.
	//
	// - NONE: No dependency is selected, meaning the node does not depend on the previous cycle.
	//
	// example:
	//
	// USER_DEFINE
	DependentType *string `json:"DependentType,omitempty" xml:"DependentType,omitempty"`
	// The UNIX timestamp, in milliseconds, when automatic scheduling stops.
	//
	// This parameter corresponds to the millisecond UNIX timestamp of the end time configured in the "Scan Configuration > Time Properties > Effective Date" setting for a Data Development job in the [DataWorks console](https://workbench.data.aliyun.com/console).
	//
	// example:
	//
	// 4155787800000
	EndEffectDate *int64 `json:"EndEffectDate,omitempty" xml:"EndEffectDate,omitempty"`
	// Schedule Configuration > Previous Cycle > Whether to ignore the upstream dry-run property.
	//
	// example:
	//
	// true
	IgnoreParentSkipRunningProperty *string `json:"IgnoreParentSkipRunningProperty,omitempty" xml:"IgnoreParentSkipRunningProperty,omitempty"`
	// Custom image ID
	//
	// example:
	//
	// m-bp1h4b5a8ogkbll2f3tr
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// Information about outputs from upstream files on which this file depends.
	InputList []*GetFileResponseBodyDataNodeConfigurationInputList `json:"InputList,omitempty" xml:"InputList,omitempty" type:"Repeated"`
	// Return Result.
	InputParameters []*GetFileResponseBodyDataNodeConfigurationInputParameters `json:"InputParameters,omitempty" xml:"InputParameters,omitempty" type:"Repeated"`
	// Output information of the file.
	OutputList []*GetFileResponseBodyDataNodeConfigurationOutputList `json:"OutputList,omitempty" xml:"OutputList,omitempty" type:"Repeated"`
	// Return Result.
	OutputParameters []*GetFileResponseBodyDataNodeConfigurationOutputParameters `json:"OutputParameters,omitempty" xml:"OutputParameters,omitempty" type:"Repeated"`
	// Schedule parameter.
	//
	// This parameter corresponds to the "Scan Configuration > Parameters" setting for a Data Development job in the [DataWorks console](https://workbench.data.aliyun.com/console). You can refer to the [Schedule Parameters](https://help.aliyun.com/document_detail/137548.html) documentation for configuration details.
	//
	// example:
	//
	// a=x b=y
	ParaValue *string `json:"ParaValue,omitempty" xml:"ParaValue,omitempty"`
	// Rerun property. Valid values:
	//
	// - ALL_ALLOWED: The job can be rerun regardless of whether it previously Succeeded or failed.
	//
	// - FAILURE_ALLOWED: The job cannot be rerun if it previously Succeeded, but can be rerun if it previously failed.
	//
	// - ALL_DENIED: The job cannot be rerun regardless of whether it previously Succeeded or failed.
	//
	// This parameter corresponds to the "Scan Configuration > Time Properties > Rerun Property" setting for a Data Development job in the [DataWorks console](https://workbench.data.aliyun.com/console).
	//
	// example:
	//
	// ALL_ALLOWED
	RerunMode *string `json:"RerunMode,omitempty" xml:"RerunMode,omitempty"`
	// The resource group used when the file is published as a Job and executed. You can call [ListResourceGroups](https://help.aliyun.com/document_detail/173913.html) to obtain the list of available resource groups in the workspace.
	//
	// example:
	//
	// 375827434852437
	ResourceGroupId *int64 `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The schedule type. Valid values:
	//
	// - NORMAL: Normal scheduling task.
	//
	// - MANUAL: One-time task, which is not included in regular scheduling and corresponds to a node in a manually triggered workflow.
	//
	// - PAUSE: Paused task.
	//
	// - SKIP: Dry-run task, which is included in regular scheduling but is immediately marked as Succeeded when scheduled.
	//
	// example:
	//
	// NORMAL
	SchedulerType *string `json:"SchedulerType,omitempty" xml:"SchedulerType,omitempty"`
	// The UNIX timestamp (in milliseconds) indicating when automatic scheduling starts.
	//
	// This parameter corresponds to the start time (as a UNIX timestamp in milliseconds) configured under "Schedule Configuration > Time Properties > Effective Date" for a Data Development job in the [DataWorks console](https://workbench.data.aliyun.com/console).
	//
	// example:
	//
	// 936923400000
	StartEffectDate *int64 `json:"StartEffectDate,omitempty" xml:"StartEffectDate,omitempty"`
	// Indicates whether to start immediately after publishing.
	//
	// This parameter corresponds to the "Start Method" setting under "Configuration > Time Properties" in the right-side navigation bar on the editing page for EMR Spark Streaming and EMR Streaming SQL Data Development jobs in the [DataWorks console](https://workbench.data.aliyun.com/console).
	//
	// example:
	//
	// true
	StartImmediately *bool `json:"StartImmediately,omitempty" xml:"StartImmediately,omitempty"`
	// Indicates whether to skip execution. Valid values:
	//
	// - true: Skip execution.
	//
	// - false: Do not skip execution.
	//
	// This parameter corresponds to the setting "Schedule Type" under "Schedule Configuration > Time Properties" for a Data Development job in the [DataWorks console](https://workbench.data.aliyun.com/console), when it is set to "skip execution".
	//
	// example:
	//
	// false
	Stop *bool `json:"Stop,omitempty" xml:"Stop,omitempty"`
	// Timeout definition for scheduling configuration.
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
	// The output name of the upstream file on which this file depends.
	//
	// This parameter corresponds to "Parent Node Output Name" when "Same Cycle" is selected under "Schedule Configuration > Schedule Dependency" for a Data Development job in the [DataWorks console](https://workbench.data.aliyun.com/console).
	//
	// example:
	//
	// project.001_out
	Input *string `json:"Input,omitempty" xml:"Input,omitempty"`
	// The method for configuring file dependencies. Valid values:
	//
	// - MANUAL: Manually configured.
	//
	// - AUTO: Automatically parsed.
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
	// The parameter name of the input parameter in the node context. You can reference this parameter in code by using the ${...} syntax.
	//
	// This parameter corresponds to the "Parameter Name" field under "Schedule Configuration > Node Context > Input Parameters of This Node" in the [DataWorks console](https://workbench.data.aliyun.com/console).
	//
	// example:
	//
	// input
	ParameterName *string `json:"ParameterName,omitempty" xml:"ParameterName,omitempty"`
	// The value source of the input parameter in the node context.
	//
	// This parameter corresponds to the "Value Source" field under "Schedule Configuration > Node Context > Input Parameters of This Node" in the [DataWorks console](https://workbench.data.aliyun.com/console).
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
	// Output name of the file.
	//
	// This parameter corresponds to the value in the "Output Name" column when "Same Cycle" is selected under "Scan Configuration > Schedule Dependency" for a Data Development job in the [DataWorks console](https://workbench.data.aliyun.com/console).
	//
	// example:
	//
	// dw_project.002_out
	Output *string `json:"Output,omitempty" xml:"Output,omitempty"`
	// Output value of the file.
	//
	// This parameter corresponds to the value in the "Output Table" column when "Same Cycle" is selected under "Scan Configuration > Schedule Dependency" for a Data Development job in the [DataWorks console](https://workbench.data.aliyun.com/console).
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
	// The description of the output parameter in the edge zone context.
	//
	// example:
	//
	// It\\"s a context output parameter.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The parameter name of the output parameter in the node context.
	//
	// This parameter corresponds to the "Parameter Name" field under "Schedule Configuration > Node Context > Output Parameters of This Node" for a Data Development job in the [DataWorks console](https://workbench.data.aliyun.com/console).
	//
	// example:
	//
	// output
	ParameterName *string `json:"ParameterName,omitempty" xml:"ParameterName,omitempty"`
	// The type of the expression for the edge zone context output parameter. Valid values are as follows:
	//
	// - 1: constant
	//
	// - 2: variable
	//
	// - 3: pass-through variable from a parameter node
	//
	// This parameter corresponds to the "Type" field in the "Scan Configuration > Edge Zone Context > Output Parameters of This Node" section for a Data Development job in the [DataWorks console](https://workbench.data.aliyun.com/console).
	//
	// example:
	//
	// 1
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The expression of the output parameter in the edge zone context.
	//
	// This parameter corresponds to the "Value" field in the "Scan Configuration > Edge Zone Context > Output Parameters of This Node" section for a Data Development job in the [DataWorks console](https://workbench.data.aliyun.com/console).
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
	// Link for downloading the resource.
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
