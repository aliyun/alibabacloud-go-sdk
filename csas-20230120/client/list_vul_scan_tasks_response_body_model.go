// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVulScanTasksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListVulScanTasksResponseBody
	GetRequestId() *string
	SetTasks(v []*ListVulScanTasksResponseBodyTasks) *ListVulScanTasksResponseBody
	GetTasks() []*ListVulScanTasksResponseBodyTasks
	SetTotalNum(v int64) *ListVulScanTasksResponseBody
	GetTotalNum() *int64
}

type ListVulScanTasksResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 3D7EC0AF-DB2A-5D9C-90EC-F090A6BAAEA7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of vulnerability scanning tasks.
	Tasks []*ListVulScanTasksResponseBodyTasks `json:"Tasks,omitempty" xml:"Tasks,omitempty" type:"Repeated"`
	// The total number of vulnerability scanning tasks that match the query conditions.
	//
	// example:
	//
	// 37
	TotalNum *int64 `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
}

func (s ListVulScanTasksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListVulScanTasksResponseBody) GoString() string {
	return s.String()
}

func (s *ListVulScanTasksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListVulScanTasksResponseBody) GetTasks() []*ListVulScanTasksResponseBodyTasks {
	return s.Tasks
}

func (s *ListVulScanTasksResponseBody) GetTotalNum() *int64 {
	return s.TotalNum
}

func (s *ListVulScanTasksResponseBody) SetRequestId(v string) *ListVulScanTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListVulScanTasksResponseBody) SetTasks(v []*ListVulScanTasksResponseBodyTasks) *ListVulScanTasksResponseBody {
	s.Tasks = v
	return s
}

func (s *ListVulScanTasksResponseBody) SetTotalNum(v int64) *ListVulScanTasksResponseBody {
	s.TotalNum = &v
	return s
}

func (s *ListVulScanTasksResponseBody) Validate() error {
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

type ListVulScanTasksResponseBodyTasks struct {
	// The task creation time, in seconds-level UNIX timestamp.
	//
	// example:
	//
	// 1786291200
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The effective scope specified by organizational structure. An empty list is returned if no organizational structure is configured.
	CustomMatchGroup []*ListVulScanTasksResponseBodyTasksCustomMatchGroup `json:"CustomMatchGroup,omitempty" xml:"CustomMatchGroup,omitempty" type:"Repeated"`
	// The task expiration time, in seconds-level UNIX timestamp. After this time, endpoints no longer pull and execute this task.
	//
	// example:
	//
	// 1786291200
	EndTimestamp *int64 `json:"EndTimestamp,omitempty" xml:"EndTimestamp,omitempty"`
	// The matching mode of the effective scope. Valid values:
	//
	// - **UserGroupAll**: applies to all users under the current Alibaba Cloud account.
	//
	// - **UserGroupNormal**: applies only to users within specified user groups.
	//
	// example:
	//
	// UserGroupNormal
	MatchMode *string `json:"MatchMode,omitempty" xml:"MatchMode,omitempty"`
	// The collection of effective user group IDs. An empty list is returned when MatchMode is UserGroupAll.
	MatchTargetIds []*string `json:"MatchTargetIds,omitempty" xml:"MatchTargetIds,omitempty" type:"Repeated"`
	// The ID of the vulnerability scheduled scan policy that triggered this task. An empty string is returned when TaskType is Instant.
	//
	// example:
	//
	// vul-scan-scheduled-strategy-8a3f6c2e91b7****
	ScheduledStrategyId *string `json:"ScheduledStrategyId,omitempty" xml:"ScheduledStrategyId,omitempty"`
	// The task status. Valid values:
	//
	// - **Running**: the task is in progress and still within the validity period.
	//
	// - **Expired**: the task has expired and exceeded the validity period.
	//
	// - **Canceled**: the task has been canceled.
	//
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The execution statistics of this task on user endpoint devices within the effective scope.
	TargetDeviceCount *ListVulScanTasksResponseBodyTasksTargetDeviceCount `json:"TargetDeviceCount,omitempty" xml:"TargetDeviceCount,omitempty" type:"Struct"`
	// The task description. An empty string is returned if no description is specified.
	//
	// example:
	//
	// Execute a vulnerability scanning on R&D department endpoints
	TaskDescription *string `json:"TaskDescription,omitempty" xml:"TaskDescription,omitempty"`
	// The vulnerability scanning task ID.
	//
	// example:
	//
	// vul-scan-task-4d7b1e9a6c38****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The task name.
	//
	// example:
	//
	// R&D department vulnerability scanning
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	// The task type. Valid values:
	//
	// - **Instant**: an instant task created by CreateVulScanTask.
	//
	// - **Scheduled**: a scheduled task automatically created by a vulnerability scheduled scan policy on a periodic basis.
	//
	// example:
	//
	// Instant
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	// The total number of vulnerabilities detected by this task.
	//
	// example:
	//
	// 27
	VulCount *int64 `json:"VulCount,omitempty" xml:"VulCount,omitempty"`
	// The list of exempted users. Users in this list are excluded from the scan. An empty list is returned if no exemption is configured.
	Whitelist []*string `json:"Whitelist,omitempty" xml:"Whitelist,omitempty" type:"Repeated"`
}

func (s ListVulScanTasksResponseBodyTasks) String() string {
	return dara.Prettify(s)
}

func (s ListVulScanTasksResponseBodyTasks) GoString() string {
	return s.String()
}

func (s *ListVulScanTasksResponseBodyTasks) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListVulScanTasksResponseBodyTasks) GetCustomMatchGroup() []*ListVulScanTasksResponseBodyTasksCustomMatchGroup {
	return s.CustomMatchGroup
}

func (s *ListVulScanTasksResponseBodyTasks) GetEndTimestamp() *int64 {
	return s.EndTimestamp
}

func (s *ListVulScanTasksResponseBodyTasks) GetMatchMode() *string {
	return s.MatchMode
}

func (s *ListVulScanTasksResponseBodyTasks) GetMatchTargetIds() []*string {
	return s.MatchTargetIds
}

func (s *ListVulScanTasksResponseBodyTasks) GetScheduledStrategyId() *string {
	return s.ScheduledStrategyId
}

func (s *ListVulScanTasksResponseBodyTasks) GetStatus() *string {
	return s.Status
}

func (s *ListVulScanTasksResponseBodyTasks) GetTargetDeviceCount() *ListVulScanTasksResponseBodyTasksTargetDeviceCount {
	return s.TargetDeviceCount
}

func (s *ListVulScanTasksResponseBodyTasks) GetTaskDescription() *string {
	return s.TaskDescription
}

func (s *ListVulScanTasksResponseBodyTasks) GetTaskId() *string {
	return s.TaskId
}

func (s *ListVulScanTasksResponseBodyTasks) GetTaskName() *string {
	return s.TaskName
}

func (s *ListVulScanTasksResponseBodyTasks) GetTaskType() *string {
	return s.TaskType
}

func (s *ListVulScanTasksResponseBodyTasks) GetVulCount() *int64 {
	return s.VulCount
}

func (s *ListVulScanTasksResponseBodyTasks) GetWhitelist() []*string {
	return s.Whitelist
}

func (s *ListVulScanTasksResponseBodyTasks) SetCreateTime(v int64) *ListVulScanTasksResponseBodyTasks {
	s.CreateTime = &v
	return s
}

func (s *ListVulScanTasksResponseBodyTasks) SetCustomMatchGroup(v []*ListVulScanTasksResponseBodyTasksCustomMatchGroup) *ListVulScanTasksResponseBodyTasks {
	s.CustomMatchGroup = v
	return s
}

func (s *ListVulScanTasksResponseBodyTasks) SetEndTimestamp(v int64) *ListVulScanTasksResponseBodyTasks {
	s.EndTimestamp = &v
	return s
}

func (s *ListVulScanTasksResponseBodyTasks) SetMatchMode(v string) *ListVulScanTasksResponseBodyTasks {
	s.MatchMode = &v
	return s
}

func (s *ListVulScanTasksResponseBodyTasks) SetMatchTargetIds(v []*string) *ListVulScanTasksResponseBodyTasks {
	s.MatchTargetIds = v
	return s
}

func (s *ListVulScanTasksResponseBodyTasks) SetScheduledStrategyId(v string) *ListVulScanTasksResponseBodyTasks {
	s.ScheduledStrategyId = &v
	return s
}

func (s *ListVulScanTasksResponseBodyTasks) SetStatus(v string) *ListVulScanTasksResponseBodyTasks {
	s.Status = &v
	return s
}

func (s *ListVulScanTasksResponseBodyTasks) SetTargetDeviceCount(v *ListVulScanTasksResponseBodyTasksTargetDeviceCount) *ListVulScanTasksResponseBodyTasks {
	s.TargetDeviceCount = v
	return s
}

func (s *ListVulScanTasksResponseBodyTasks) SetTaskDescription(v string) *ListVulScanTasksResponseBodyTasks {
	s.TaskDescription = &v
	return s
}

func (s *ListVulScanTasksResponseBodyTasks) SetTaskId(v string) *ListVulScanTasksResponseBodyTasks {
	s.TaskId = &v
	return s
}

func (s *ListVulScanTasksResponseBodyTasks) SetTaskName(v string) *ListVulScanTasksResponseBodyTasks {
	s.TaskName = &v
	return s
}

func (s *ListVulScanTasksResponseBodyTasks) SetTaskType(v string) *ListVulScanTasksResponseBodyTasks {
	s.TaskType = &v
	return s
}

func (s *ListVulScanTasksResponseBodyTasks) SetVulCount(v int64) *ListVulScanTasksResponseBodyTasks {
	s.VulCount = &v
	return s
}

func (s *ListVulScanTasksResponseBodyTasks) SetWhitelist(v []*string) *ListVulScanTasksResponseBodyTasks {
	s.Whitelist = v
	return s
}

func (s *ListVulScanTasksResponseBodyTasks) Validate() error {
	if s.CustomMatchGroup != nil {
		for _, item := range s.CustomMatchGroup {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.TargetDeviceCount != nil {
		if err := s.TargetDeviceCount.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListVulScanTasksResponseBodyTasksCustomMatchGroup struct {
	// The collection of organizational structure nodes.
	Group []*string `json:"Group,omitempty" xml:"Group,omitempty" type:"Repeated"`
	// The identity provider ID.
	//
	// example:
	//
	// idp-7c3f9a2e5b18****
	IdpId *string `json:"IdpId,omitempty" xml:"IdpId,omitempty"`
}

func (s ListVulScanTasksResponseBodyTasksCustomMatchGroup) String() string {
	return dara.Prettify(s)
}

func (s ListVulScanTasksResponseBodyTasksCustomMatchGroup) GoString() string {
	return s.String()
}

func (s *ListVulScanTasksResponseBodyTasksCustomMatchGroup) GetGroup() []*string {
	return s.Group
}

func (s *ListVulScanTasksResponseBodyTasksCustomMatchGroup) GetIdpId() *string {
	return s.IdpId
}

func (s *ListVulScanTasksResponseBodyTasksCustomMatchGroup) SetGroup(v []*string) *ListVulScanTasksResponseBodyTasksCustomMatchGroup {
	s.Group = v
	return s
}

func (s *ListVulScanTasksResponseBodyTasksCustomMatchGroup) SetIdpId(v string) *ListVulScanTasksResponseBodyTasksCustomMatchGroup {
	s.IdpId = &v
	return s
}

func (s *ListVulScanTasksResponseBodyTasksCustomMatchGroup) Validate() error {
	return dara.Validate(s)
}

type ListVulScanTasksResponseBodyTasksTargetDeviceCount struct {
	// The number of user endpoint devices that have acknowledged receipt of this task.
	//
	// example:
	//
	// 12
	AckCount *int64 `json:"AckCount,omitempty" xml:"AckCount,omitempty"`
	// The number of user endpoint devices on which the scan failed.
	//
	// example:
	//
	// 1
	FailCount *int64 `json:"FailCount,omitempty" xml:"FailCount,omitempty"`
	// The number of user endpoint devices currently executing the scan. This value is calculated by subtracting SuccessCount and FailCount from AckCount.
	//
	// example:
	//
	// 3
	StartCount *int64 `json:"StartCount,omitempty" xml:"StartCount,omitempty"`
	// The number of user endpoint devices on which the scan succeeded.
	//
	// example:
	//
	// 8
	SuccessCount *int64 `json:"SuccessCount,omitempty" xml:"SuccessCount,omitempty"`
}

func (s ListVulScanTasksResponseBodyTasksTargetDeviceCount) String() string {
	return dara.Prettify(s)
}

func (s ListVulScanTasksResponseBodyTasksTargetDeviceCount) GoString() string {
	return s.String()
}

func (s *ListVulScanTasksResponseBodyTasksTargetDeviceCount) GetAckCount() *int64 {
	return s.AckCount
}

func (s *ListVulScanTasksResponseBodyTasksTargetDeviceCount) GetFailCount() *int64 {
	return s.FailCount
}

func (s *ListVulScanTasksResponseBodyTasksTargetDeviceCount) GetStartCount() *int64 {
	return s.StartCount
}

func (s *ListVulScanTasksResponseBodyTasksTargetDeviceCount) GetSuccessCount() *int64 {
	return s.SuccessCount
}

func (s *ListVulScanTasksResponseBodyTasksTargetDeviceCount) SetAckCount(v int64) *ListVulScanTasksResponseBodyTasksTargetDeviceCount {
	s.AckCount = &v
	return s
}

func (s *ListVulScanTasksResponseBodyTasksTargetDeviceCount) SetFailCount(v int64) *ListVulScanTasksResponseBodyTasksTargetDeviceCount {
	s.FailCount = &v
	return s
}

func (s *ListVulScanTasksResponseBodyTasksTargetDeviceCount) SetStartCount(v int64) *ListVulScanTasksResponseBodyTasksTargetDeviceCount {
	s.StartCount = &v
	return s
}

func (s *ListVulScanTasksResponseBodyTasksTargetDeviceCount) SetSuccessCount(v int64) *ListVulScanTasksResponseBodyTasksTargetDeviceCount {
	s.SuccessCount = &v
	return s
}

func (s *ListVulScanTasksResponseBodyTasksTargetDeviceCount) Validate() error {
	return dara.Validate(s)
}
