// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetTaskResponseBody
	GetRequestId() *string
	SetTask(v *GetTaskResponseBodyTask) *GetTaskResponseBody
	GetTask() *GetTaskResponseBodyTask
}

type GetTaskResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// C24C498A-09CF-54D3-8972-8DC074CF8614
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The task information.
	Task *GetTaskResponseBodyTask `json:"task,omitempty" xml:"task,omitempty" type:"Struct"`
}

func (s GetTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTaskResponseBody) GoString() string {
	return s.String()
}

func (s *GetTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTaskResponseBody) GetTask() *GetTaskResponseBodyTask {
	return s.Task
}

func (s *GetTaskResponseBody) SetRequestId(v string) *GetTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTaskResponseBody) SetTask(v *GetTaskResponseBodyTask) *GetTaskResponseBody {
	s.Task = v
	return s
}

func (s *GetTaskResponseBody) Validate() error {
	if s.Task != nil {
		if err := s.Task.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetTaskResponseBodyTask struct {
	// Indicates whether the task is automatically executed.
	//
	// example:
	//
	// true
	AutoApply *bool `json:"autoApply,omitempty" xml:"autoApply,omitempty"`
	// Indicates whether automatic deletion is enabled. When enabled, resources are automatically destroyed after the task is completed.
	//
	// example:
	//
	// false
	AutoDestroy *bool `json:"autoDestroy,omitempty" xml:"autoDestroy,omitempty"`
	// The time when the task was created.
	//
	// example:
	//
	// 2022-06-15T02:44:37Z
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// The job ID of the current task.
	//
	// example:
	//
	// job-absdf
	CurrentJobId *string `json:"currentJobId,omitempty" xml:"currentJobId,omitempty"`
	// The current job status.
	//
	// example:
	//
	// Planned
	CurrentJobStatus *string `json:"currentJobStatus,omitempty" xml:"currentJobStatus,omitempty"`
	// Indicates whether deletion protection is enabled.
	//
	// example:
	//
	// true
	DeletionProtection *bool `json:"deletionProtection,omitempty" xml:"deletionProtection,omitempty"`
	// The description of the task.
	//
	// example:
	//
	// this is description
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The group information.
	GroupInfo *GetTaskResponseBodyTaskGroupInfo `json:"groupInfo,omitempty" xml:"groupInfo,omitempty" type:"Struct"`
	// Specifies whether to use a state file. Default value: false. This parameter is applicable to templates that originate from resource export. Only one task can use this parameter at a time.
	//
	// example:
	//
	// false
	InitModuleState *bool `json:"initModuleState,omitempty" xml:"initModuleState,omitempty"`
	// The latest version number of the template.
	//
	// example:
	//
	// v3
	LatestModuleVersion *string `json:"latestModuleVersion,omitempty" xml:"latestModuleVersion,omitempty"`
	// The template ID.
	//
	// example:
	//
	// mod-4267dcfbf1b6d14625614ddbe15
	ModuleId *string `json:"moduleId,omitempty" xml:"moduleId,omitempty"`
	// The template name.
	//
	// example:
	//
	// moduleName
	ModuleName *string `json:"moduleName,omitempty" xml:"moduleName,omitempty"`
	// The template version.
	//
	// example:
	//
	// v2
	ModuleVersion *string `json:"moduleVersion,omitempty" xml:"moduleVersion,omitempty"`
	// The task name.
	//
	// example:
	//
	// TaskName
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The list of resource protection strategies.
	ProtectionStrategy []*string `json:"protectionStrategy,omitempty" xml:"protectionStrategy,omitempty" type:"Repeated"`
	// The RAM role.
	//
	// example:
	//
	// role
	RamRole *string `json:"ramRole,omitempty" xml:"ramRole,omitempty"`
	// Specifies whether to skip enumeration value validation. Default value: false.
	//
	// example:
	//
	// false
	SkipPropertyValidation *bool `json:"skipPropertyValidation,omitempty" xml:"skipPropertyValidation,omitempty"`
	SkipRegionValidation   *bool `json:"skipRegionValidation,omitempty" xml:"skipRegionValidation,omitempty"`
	// The task status. Valid values:
	//
	// - Available: the task is available and no job is running.
	//
	// - Running: a job is currently running.
	//
	// example:
	//
	// Running
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The list of task tags.
	Tags []*GetTaskResponseBodyTaskTags `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
	// The task backend configuration. After this parameter is configured, runtime log information is saved to the specified OSS bucket.
	TaskBackend *GetTaskResponseBodyTaskTaskBackend `json:"taskBackend,omitempty" xml:"taskBackend,omitempty" type:"Struct"`
	// The task ID.
	//
	// example:
	//
	// task-433aead756057154bda7f1c2e98
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
	// The task output path.
	//
	// example:
	//
	// /
	TaskOutputPath *string `json:"taskOutputPath,omitempty" xml:"taskOutputPath,omitempty"`
	// The Terraform version.
	//
	// example:
	//
	// 1.5.7
	TerraformVersion *string `json:"terraformVersion,omitempty" xml:"terraformVersion,omitempty"`
	// The job trigger method. Valid values:
	//
	// - Manual: manually triggered (default).
	//
	// - NewVersion: triggered when a new template version is published.
	//
	// - ParameterSetUpdated: triggered when the parameter set content changes or the parameter set attach relationship changes.
	//
	// - Auto: automatically triggered when the task properties change, such as task creation, execution version change, or job trigger policy change (when changed from another value to Auto).
	//
	// example:
	//
	// Manual
	TriggerStrategy *string `json:"triggerStrategy,omitempty" xml:"triggerStrategy,omitempty"`
}

func (s GetTaskResponseBodyTask) String() string {
	return dara.Prettify(s)
}

func (s GetTaskResponseBodyTask) GoString() string {
	return s.String()
}

func (s *GetTaskResponseBodyTask) GetAutoApply() *bool {
	return s.AutoApply
}

func (s *GetTaskResponseBodyTask) GetAutoDestroy() *bool {
	return s.AutoDestroy
}

func (s *GetTaskResponseBodyTask) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetTaskResponseBodyTask) GetCurrentJobId() *string {
	return s.CurrentJobId
}

func (s *GetTaskResponseBodyTask) GetCurrentJobStatus() *string {
	return s.CurrentJobStatus
}

func (s *GetTaskResponseBodyTask) GetDeletionProtection() *bool {
	return s.DeletionProtection
}

func (s *GetTaskResponseBodyTask) GetDescription() *string {
	return s.Description
}

func (s *GetTaskResponseBodyTask) GetGroupInfo() *GetTaskResponseBodyTaskGroupInfo {
	return s.GroupInfo
}

func (s *GetTaskResponseBodyTask) GetInitModuleState() *bool {
	return s.InitModuleState
}

func (s *GetTaskResponseBodyTask) GetLatestModuleVersion() *string {
	return s.LatestModuleVersion
}

func (s *GetTaskResponseBodyTask) GetModuleId() *string {
	return s.ModuleId
}

func (s *GetTaskResponseBodyTask) GetModuleName() *string {
	return s.ModuleName
}

func (s *GetTaskResponseBodyTask) GetModuleVersion() *string {
	return s.ModuleVersion
}

func (s *GetTaskResponseBodyTask) GetName() *string {
	return s.Name
}

func (s *GetTaskResponseBodyTask) GetProtectionStrategy() []*string {
	return s.ProtectionStrategy
}

func (s *GetTaskResponseBodyTask) GetRamRole() *string {
	return s.RamRole
}

func (s *GetTaskResponseBodyTask) GetSkipPropertyValidation() *bool {
	return s.SkipPropertyValidation
}

func (s *GetTaskResponseBodyTask) GetSkipRegionValidation() *bool {
	return s.SkipRegionValidation
}

func (s *GetTaskResponseBodyTask) GetStatus() *string {
	return s.Status
}

func (s *GetTaskResponseBodyTask) GetTags() []*GetTaskResponseBodyTaskTags {
	return s.Tags
}

func (s *GetTaskResponseBodyTask) GetTaskBackend() *GetTaskResponseBodyTaskTaskBackend {
	return s.TaskBackend
}

func (s *GetTaskResponseBodyTask) GetTaskId() *string {
	return s.TaskId
}

func (s *GetTaskResponseBodyTask) GetTaskOutputPath() *string {
	return s.TaskOutputPath
}

func (s *GetTaskResponseBodyTask) GetTerraformVersion() *string {
	return s.TerraformVersion
}

func (s *GetTaskResponseBodyTask) GetTriggerStrategy() *string {
	return s.TriggerStrategy
}

func (s *GetTaskResponseBodyTask) SetAutoApply(v bool) *GetTaskResponseBodyTask {
	s.AutoApply = &v
	return s
}

func (s *GetTaskResponseBodyTask) SetAutoDestroy(v bool) *GetTaskResponseBodyTask {
	s.AutoDestroy = &v
	return s
}

func (s *GetTaskResponseBodyTask) SetCreateTime(v string) *GetTaskResponseBodyTask {
	s.CreateTime = &v
	return s
}

func (s *GetTaskResponseBodyTask) SetCurrentJobId(v string) *GetTaskResponseBodyTask {
	s.CurrentJobId = &v
	return s
}

func (s *GetTaskResponseBodyTask) SetCurrentJobStatus(v string) *GetTaskResponseBodyTask {
	s.CurrentJobStatus = &v
	return s
}

func (s *GetTaskResponseBodyTask) SetDeletionProtection(v bool) *GetTaskResponseBodyTask {
	s.DeletionProtection = &v
	return s
}

func (s *GetTaskResponseBodyTask) SetDescription(v string) *GetTaskResponseBodyTask {
	s.Description = &v
	return s
}

func (s *GetTaskResponseBodyTask) SetGroupInfo(v *GetTaskResponseBodyTaskGroupInfo) *GetTaskResponseBodyTask {
	s.GroupInfo = v
	return s
}

func (s *GetTaskResponseBodyTask) SetInitModuleState(v bool) *GetTaskResponseBodyTask {
	s.InitModuleState = &v
	return s
}

func (s *GetTaskResponseBodyTask) SetLatestModuleVersion(v string) *GetTaskResponseBodyTask {
	s.LatestModuleVersion = &v
	return s
}

func (s *GetTaskResponseBodyTask) SetModuleId(v string) *GetTaskResponseBodyTask {
	s.ModuleId = &v
	return s
}

func (s *GetTaskResponseBodyTask) SetModuleName(v string) *GetTaskResponseBodyTask {
	s.ModuleName = &v
	return s
}

func (s *GetTaskResponseBodyTask) SetModuleVersion(v string) *GetTaskResponseBodyTask {
	s.ModuleVersion = &v
	return s
}

func (s *GetTaskResponseBodyTask) SetName(v string) *GetTaskResponseBodyTask {
	s.Name = &v
	return s
}

func (s *GetTaskResponseBodyTask) SetProtectionStrategy(v []*string) *GetTaskResponseBodyTask {
	s.ProtectionStrategy = v
	return s
}

func (s *GetTaskResponseBodyTask) SetRamRole(v string) *GetTaskResponseBodyTask {
	s.RamRole = &v
	return s
}

func (s *GetTaskResponseBodyTask) SetSkipPropertyValidation(v bool) *GetTaskResponseBodyTask {
	s.SkipPropertyValidation = &v
	return s
}

func (s *GetTaskResponseBodyTask) SetSkipRegionValidation(v bool) *GetTaskResponseBodyTask {
	s.SkipRegionValidation = &v
	return s
}

func (s *GetTaskResponseBodyTask) SetStatus(v string) *GetTaskResponseBodyTask {
	s.Status = &v
	return s
}

func (s *GetTaskResponseBodyTask) SetTags(v []*GetTaskResponseBodyTaskTags) *GetTaskResponseBodyTask {
	s.Tags = v
	return s
}

func (s *GetTaskResponseBodyTask) SetTaskBackend(v *GetTaskResponseBodyTaskTaskBackend) *GetTaskResponseBodyTask {
	s.TaskBackend = v
	return s
}

func (s *GetTaskResponseBodyTask) SetTaskId(v string) *GetTaskResponseBodyTask {
	s.TaskId = &v
	return s
}

func (s *GetTaskResponseBodyTask) SetTaskOutputPath(v string) *GetTaskResponseBodyTask {
	s.TaskOutputPath = &v
	return s
}

func (s *GetTaskResponseBodyTask) SetTerraformVersion(v string) *GetTaskResponseBodyTask {
	s.TerraformVersion = &v
	return s
}

func (s *GetTaskResponseBodyTask) SetTriggerStrategy(v string) *GetTaskResponseBodyTask {
	s.TriggerStrategy = &v
	return s
}

func (s *GetTaskResponseBodyTask) Validate() error {
	if s.GroupInfo != nil {
		if err := s.GroupInfo.Validate(); err != nil {
			return err
		}
	}
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.TaskBackend != nil {
		if err := s.TaskBackend.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetTaskResponseBodyTaskGroupInfo struct {
	// The group ID.
	//
	// example:
	//
	// g-59d8d22e78792ffe3d3eb6154d727
	GroupId *string `json:"groupId,omitempty" xml:"groupId,omitempty"`
	// The group name.
	//
	// example:
	//
	// abc
	GroupName *string `json:"groupName,omitempty" xml:"groupName,omitempty"`
	// The project ID.
	//
	// example:
	//
	// p-433aead756057fff47ecbfd94d76
	ProjectId *string `json:"projectId,omitempty" xml:"projectId,omitempty"`
	// The project name.
	//
	// example:
	//
	// abc
	ProjectName *string `json:"projectName,omitempty" xml:"projectName,omitempty"`
}

func (s GetTaskResponseBodyTaskGroupInfo) String() string {
	return dara.Prettify(s)
}

func (s GetTaskResponseBodyTaskGroupInfo) GoString() string {
	return s.String()
}

func (s *GetTaskResponseBodyTaskGroupInfo) GetGroupId() *string {
	return s.GroupId
}

func (s *GetTaskResponseBodyTaskGroupInfo) GetGroupName() *string {
	return s.GroupName
}

func (s *GetTaskResponseBodyTaskGroupInfo) GetProjectId() *string {
	return s.ProjectId
}

func (s *GetTaskResponseBodyTaskGroupInfo) GetProjectName() *string {
	return s.ProjectName
}

func (s *GetTaskResponseBodyTaskGroupInfo) SetGroupId(v string) *GetTaskResponseBodyTaskGroupInfo {
	s.GroupId = &v
	return s
}

func (s *GetTaskResponseBodyTaskGroupInfo) SetGroupName(v string) *GetTaskResponseBodyTaskGroupInfo {
	s.GroupName = &v
	return s
}

func (s *GetTaskResponseBodyTaskGroupInfo) SetProjectId(v string) *GetTaskResponseBodyTaskGroupInfo {
	s.ProjectId = &v
	return s
}

func (s *GetTaskResponseBodyTaskGroupInfo) SetProjectName(v string) *GetTaskResponseBodyTaskGroupInfo {
	s.ProjectName = &v
	return s
}

func (s *GetTaskResponseBodyTaskGroupInfo) Validate() error {
	return dara.Validate(s)
}

type GetTaskResponseBodyTaskTags struct {
	// The tag key of the task.
	//
	// example:
	//
	// TestKey
	TagKey *string `json:"tagKey,omitempty" xml:"tagKey,omitempty"`
	// The tag value of the task.
	//
	// example:
	//
	// TestValue
	TagValue *string `json:"tagValue,omitempty" xml:"tagValue,omitempty"`
}

func (s GetTaskResponseBodyTaskTags) String() string {
	return dara.Prettify(s)
}

func (s GetTaskResponseBodyTaskTags) GoString() string {
	return s.String()
}

func (s *GetTaskResponseBodyTaskTags) GetTagKey() *string {
	return s.TagKey
}

func (s *GetTaskResponseBodyTaskTags) GetTagValue() *string {
	return s.TagValue
}

func (s *GetTaskResponseBodyTaskTags) SetTagKey(v string) *GetTaskResponseBodyTaskTags {
	s.TagKey = &v
	return s
}

func (s *GetTaskResponseBodyTaskTags) SetTagValue(v string) *GetTaskResponseBodyTaskTags {
	s.TagValue = &v
	return s
}

func (s *GetTaskResponseBodyTaskTags) Validate() error {
	return dara.Validate(s)
}

type GetTaskResponseBodyTaskTaskBackend struct {
	// The endpoint information.
	//
	// example:
	//
	// ss-cn-beijing.aliyuncs.com
	BucketEndpoint *string `json:"bucketEndpoint,omitempty" xml:"bucketEndpoint,omitempty"`
	// The bucket name.
	//
	// example:
	//
	// iac-runtime-test
	BucketName *string `json:"bucketName,omitempty" xml:"bucketName,omitempty"`
	// The object path.
	//
	// example:
	//
	// /log
	ObjectPath *string `json:"objectPath,omitempty" xml:"objectPath,omitempty"`
}

func (s GetTaskResponseBodyTaskTaskBackend) String() string {
	return dara.Prettify(s)
}

func (s GetTaskResponseBodyTaskTaskBackend) GoString() string {
	return s.String()
}

func (s *GetTaskResponseBodyTaskTaskBackend) GetBucketEndpoint() *string {
	return s.BucketEndpoint
}

func (s *GetTaskResponseBodyTaskTaskBackend) GetBucketName() *string {
	return s.BucketName
}

func (s *GetTaskResponseBodyTaskTaskBackend) GetObjectPath() *string {
	return s.ObjectPath
}

func (s *GetTaskResponseBodyTaskTaskBackend) SetBucketEndpoint(v string) *GetTaskResponseBodyTaskTaskBackend {
	s.BucketEndpoint = &v
	return s
}

func (s *GetTaskResponseBodyTaskTaskBackend) SetBucketName(v string) *GetTaskResponseBodyTaskTaskBackend {
	s.BucketName = &v
	return s
}

func (s *GetTaskResponseBodyTaskTaskBackend) SetObjectPath(v string) *GetTaskResponseBodyTaskTaskBackend {
	s.ObjectPath = &v
	return s
}

func (s *GetTaskResponseBodyTaskTaskBackend) Validate() error {
	return dara.Validate(s)
}
