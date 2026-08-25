// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTasksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListTasksResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListTasksResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListTasksResponseBody
	GetRequestId() *string
	SetTasks(v []*ListTasksResponseBodyTasks) *ListTasksResponseBody
	GetTasks() []*ListTasksResponseBodyTasks
	SetTotalCount(v int32) *ListTasksResponseBody
	GetTotalCount() *int32
}

type ListTasksResponseBody struct {
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The number of results returned per page. Default value: 20. Minimum value: 1. Maximum value: 100.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 98610149-488B-5E48-B981-8D4CE1AF77CD
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The list of tasks.
	Tasks []*ListTasksResponseBodyTasks `json:"tasks,omitempty" xml:"tasks,omitempty" type:"Repeated"`
	// The total number of records.
	//
	// example:
	//
	// 55
	TotalCount *int32 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListTasksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTasksResponseBody) GoString() string {
	return s.String()
}

func (s *ListTasksResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListTasksResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListTasksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTasksResponseBody) GetTasks() []*ListTasksResponseBodyTasks {
	return s.Tasks
}

func (s *ListTasksResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListTasksResponseBody) SetPageNumber(v int32) *ListTasksResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListTasksResponseBody) SetPageSize(v int32) *ListTasksResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListTasksResponseBody) SetRequestId(v string) *ListTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTasksResponseBody) SetTasks(v []*ListTasksResponseBodyTasks) *ListTasksResponseBody {
	s.Tasks = v
	return s
}

func (s *ListTasksResponseBody) SetTotalCount(v int32) *ListTasksResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListTasksResponseBody) Validate() error {
	if s.Tasks != nil {
		for _, item := range s.Tasks {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListTasksResponseBodyTasks struct {
	// Indicates whether the task is automatically executed.
	//
	// example:
	//
	// false
	AutoApply *bool `json:"autoApply,omitempty" xml:"autoApply,omitempty"`
	// The time when the task was created, in UTC in the ISO 8601 format of YYYY-MM-DDTHH:mm:ssZ.
	//
	// example:
	//
	// 2022-07-11T15:09:53Z
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// The job ID of the current task.
	//
	// example:
	//
	// job-123asd
	CurrentJobId *string `json:"currentJobId,omitempty" xml:"currentJobId,omitempty"`
	// The status of the current job.
	//
	// example:
	//
	// Pending
	CurrentJobStatus *string `json:"currentJobStatus,omitempty" xml:"currentJobStatus,omitempty"`
	// Indicates whether deletion protection is enabled. Deletion protection is automatically enabled when managed resources exist.
	//
	// example:
	//
	// false
	DeletionProtection *bool `json:"deletionProtection,omitempty" xml:"deletionProtection,omitempty"`
	// The task group information.
	GroupInfo *ListTasksResponseBodyTasksGroupInfo `json:"groupInfo,omitempty" xml:"groupInfo,omitempty" type:"Struct"`
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
	// mod-518855d9a058c331e9c60bc0ce
	ModuleId *string `json:"moduleId,omitempty" xml:"moduleId,omitempty"`
	// The template name.
	//
	// example:
	//
	// mod-name
	ModuleName *string `json:"moduleName,omitempty" xml:"moduleName,omitempty"`
	// The template version.
	//
	// example:
	//
	// v1
	ModuleVersion *string `json:"moduleVersion,omitempty" xml:"moduleVersion,omitempty"`
	// The task name.
	//
	// example:
	//
	// TaskName
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The task status. Valid values:
	//
	// - Available: The task is in an available state with no job running.
	//
	// - Running: The task is in a running state with a current job in progress.
	//
	// example:
	//
	// Running
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The list of task tags.
	Tags []*ListTasksResponseBodyTasksTags `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
	// The task ID.
	//
	// example:
	//
	// task-1525e992f1b621b0ca51647876e
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s ListTasksResponseBodyTasks) String() string {
	return dara.Prettify(s)
}

func (s ListTasksResponseBodyTasks) GoString() string {
	return s.String()
}

func (s *ListTasksResponseBodyTasks) GetAutoApply() *bool {
	return s.AutoApply
}

func (s *ListTasksResponseBodyTasks) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListTasksResponseBodyTasks) GetCurrentJobId() *string {
	return s.CurrentJobId
}

func (s *ListTasksResponseBodyTasks) GetCurrentJobStatus() *string {
	return s.CurrentJobStatus
}

func (s *ListTasksResponseBodyTasks) GetDeletionProtection() *bool {
	return s.DeletionProtection
}

func (s *ListTasksResponseBodyTasks) GetGroupInfo() *ListTasksResponseBodyTasksGroupInfo {
	return s.GroupInfo
}

func (s *ListTasksResponseBodyTasks) GetLatestModuleVersion() *string {
	return s.LatestModuleVersion
}

func (s *ListTasksResponseBodyTasks) GetModuleId() *string {
	return s.ModuleId
}

func (s *ListTasksResponseBodyTasks) GetModuleName() *string {
	return s.ModuleName
}

func (s *ListTasksResponseBodyTasks) GetModuleVersion() *string {
	return s.ModuleVersion
}

func (s *ListTasksResponseBodyTasks) GetName() *string {
	return s.Name
}

func (s *ListTasksResponseBodyTasks) GetStatus() *string {
	return s.Status
}

func (s *ListTasksResponseBodyTasks) GetTags() []*ListTasksResponseBodyTasksTags {
	return s.Tags
}

func (s *ListTasksResponseBodyTasks) GetTaskId() *string {
	return s.TaskId
}

func (s *ListTasksResponseBodyTasks) SetAutoApply(v bool) *ListTasksResponseBodyTasks {
	s.AutoApply = &v
	return s
}

func (s *ListTasksResponseBodyTasks) SetCreateTime(v string) *ListTasksResponseBodyTasks {
	s.CreateTime = &v
	return s
}

func (s *ListTasksResponseBodyTasks) SetCurrentJobId(v string) *ListTasksResponseBodyTasks {
	s.CurrentJobId = &v
	return s
}

func (s *ListTasksResponseBodyTasks) SetCurrentJobStatus(v string) *ListTasksResponseBodyTasks {
	s.CurrentJobStatus = &v
	return s
}

func (s *ListTasksResponseBodyTasks) SetDeletionProtection(v bool) *ListTasksResponseBodyTasks {
	s.DeletionProtection = &v
	return s
}

func (s *ListTasksResponseBodyTasks) SetGroupInfo(v *ListTasksResponseBodyTasksGroupInfo) *ListTasksResponseBodyTasks {
	s.GroupInfo = v
	return s
}

func (s *ListTasksResponseBodyTasks) SetLatestModuleVersion(v string) *ListTasksResponseBodyTasks {
	s.LatestModuleVersion = &v
	return s
}

func (s *ListTasksResponseBodyTasks) SetModuleId(v string) *ListTasksResponseBodyTasks {
	s.ModuleId = &v
	return s
}

func (s *ListTasksResponseBodyTasks) SetModuleName(v string) *ListTasksResponseBodyTasks {
	s.ModuleName = &v
	return s
}

func (s *ListTasksResponseBodyTasks) SetModuleVersion(v string) *ListTasksResponseBodyTasks {
	s.ModuleVersion = &v
	return s
}

func (s *ListTasksResponseBodyTasks) SetName(v string) *ListTasksResponseBodyTasks {
	s.Name = &v
	return s
}

func (s *ListTasksResponseBodyTasks) SetStatus(v string) *ListTasksResponseBodyTasks {
	s.Status = &v
	return s
}

func (s *ListTasksResponseBodyTasks) SetTags(v []*ListTasksResponseBodyTasksTags) *ListTasksResponseBodyTasks {
	s.Tags = v
	return s
}

func (s *ListTasksResponseBodyTasks) SetTaskId(v string) *ListTasksResponseBodyTasks {
	s.TaskId = &v
	return s
}

func (s *ListTasksResponseBodyTasks) Validate() error {
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
	return nil
}

type ListTasksResponseBodyTasksGroupInfo struct {
	// The group ID.
	//
	// example:
	//
	// g-4267dcfbf1b6d1e0652bfbbe995
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
	// p-433aead7560571cf1b2bfbbe92b
	ProjectId *string `json:"projectId,omitempty" xml:"projectId,omitempty"`
	// The project name.
	//
	// example:
	//
	// abc
	ProjectName *string `json:"projectName,omitempty" xml:"projectName,omitempty"`
}

func (s ListTasksResponseBodyTasksGroupInfo) String() string {
	return dara.Prettify(s)
}

func (s ListTasksResponseBodyTasksGroupInfo) GoString() string {
	return s.String()
}

func (s *ListTasksResponseBodyTasksGroupInfo) GetGroupId() *string {
	return s.GroupId
}

func (s *ListTasksResponseBodyTasksGroupInfo) GetGroupName() *string {
	return s.GroupName
}

func (s *ListTasksResponseBodyTasksGroupInfo) GetProjectId() *string {
	return s.ProjectId
}

func (s *ListTasksResponseBodyTasksGroupInfo) GetProjectName() *string {
	return s.ProjectName
}

func (s *ListTasksResponseBodyTasksGroupInfo) SetGroupId(v string) *ListTasksResponseBodyTasksGroupInfo {
	s.GroupId = &v
	return s
}

func (s *ListTasksResponseBodyTasksGroupInfo) SetGroupName(v string) *ListTasksResponseBodyTasksGroupInfo {
	s.GroupName = &v
	return s
}

func (s *ListTasksResponseBodyTasksGroupInfo) SetProjectId(v string) *ListTasksResponseBodyTasksGroupInfo {
	s.ProjectId = &v
	return s
}

func (s *ListTasksResponseBodyTasksGroupInfo) SetProjectName(v string) *ListTasksResponseBodyTasksGroupInfo {
	s.ProjectName = &v
	return s
}

func (s *ListTasksResponseBodyTasksGroupInfo) Validate() error {
	return dara.Validate(s)
}

type ListTasksResponseBodyTasksTags struct {
	// The task tag key.
	//
	// example:
	//
	// name
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// The tag key of the template.
	//
	// example:
	//
	// name
	TagKey *string `json:"tagKey,omitempty" xml:"tagKey,omitempty"`
	// The tag value of the task.
	//
	// example:
	//
	// iac-demo
	TagValue *string `json:"tagValue,omitempty" xml:"tagValue,omitempty"`
	// The task tag value.
	//
	// example:
	//
	// iac-demo
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s ListTasksResponseBodyTasksTags) String() string {
	return dara.Prettify(s)
}

func (s ListTasksResponseBodyTasksTags) GoString() string {
	return s.String()
}

func (s *ListTasksResponseBodyTasksTags) GetKey() *string {
	return s.Key
}

func (s *ListTasksResponseBodyTasksTags) GetTagKey() *string {
	return s.TagKey
}

func (s *ListTasksResponseBodyTasksTags) GetTagValue() *string {
	return s.TagValue
}

func (s *ListTasksResponseBodyTasksTags) GetValue() *string {
	return s.Value
}

func (s *ListTasksResponseBodyTasksTags) SetKey(v string) *ListTasksResponseBodyTasksTags {
	s.Key = &v
	return s
}

func (s *ListTasksResponseBodyTasksTags) SetTagKey(v string) *ListTasksResponseBodyTasksTags {
	s.TagKey = &v
	return s
}

func (s *ListTasksResponseBodyTasksTags) SetTagValue(v string) *ListTasksResponseBodyTasksTags {
	s.TagValue = &v
	return s
}

func (s *ListTasksResponseBodyTasksTags) SetValue(v string) *ListTasksResponseBodyTasksTags {
	s.Value = &v
	return s
}

func (s *ListTasksResponseBodyTasksTags) Validate() error {
	return dara.Validate(s)
}
