// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVulScanTasksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int64) *ListVulScanTasksRequest
	GetCurrentPage() *int64
	SetMatchMode(v string) *ListVulScanTasksRequest
	GetMatchMode() *string
	SetPageSize(v int64) *ListVulScanTasksRequest
	GetPageSize() *int64
	SetScheduledStrategyId(v string) *ListVulScanTasksRequest
	GetScheduledStrategyId() *string
	SetStatus(v string) *ListVulScanTasksRequest
	GetStatus() *string
	SetTaskIds(v []*string) *ListVulScanTasksRequest
	GetTaskIds() []*string
	SetTaskName(v string) *ListVulScanTasksRequest
	GetTaskName() *string
	SetTaskType(v string) *ListVulScanTasksRequest
	GetTaskType() *string
	SetUserGroupId(v string) *ListVulScanTasksRequest
	GetUserGroupId() *string
}

type ListVulScanTasksRequest struct {
	// The page number of the current page in a paged query. Valid values: 1 to 10000.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurrentPage *int64 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// Filters by the matching mode of the effective scope. Valid values:
	//
	// - **UserGroupAll**: applies to all users under the current Alibaba Cloud account.
	//
	// - **UserGroupNormal**: applies only to users within specified user groups.
	//
	// example:
	//
	// UserGroupNormal
	MatchMode *string `json:"MatchMode,omitempty" xml:"MatchMode,omitempty"`
	// The number of entries per page in a paged query. Valid values: 1 to 1000.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the vulnerability scheduled scan policy. This parameter is used to filter tasks triggered by the specified policy. Valid values are obtained from:
	//
	// - [ListVulScanScheduledStrategies](~~ListVulScanScheduledStrategies~~): lists vulnerability scheduled scan policies.
	//
	// - [CreateVulScanScheduledStrategy](~~CreateVulScanScheduledStrategy~~): creates a vulnerability scheduled scan policy.
	//
	// example:
	//
	// vul-scan-scheduled-strategy-8a3f6c2e91b7****
	ScheduledStrategyId *string `json:"ScheduledStrategyId,omitempty" xml:"ScheduledStrategyId,omitempty"`
	// Filters by task status. Valid values:
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
	// The vulnerability scanning task IDs used for filtering. A maximum of 100 IDs can be specified. Duplicate IDs are not allowed.
	TaskIds []*string `json:"TaskIds,omitempty" xml:"TaskIds,omitempty" type:"Repeated"`
	// The task name. Fuzzy match is supported. The name can be up to 128 characters in length.
	//
	// example:
	//
	// R&D department vulnerability scanning
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	// Filters by task type. Valid values:
	//
	// - **Instant**: an instant task created by CreateVulScanTask.
	//
	// - **Scheduled**: a scheduled task automatically created by a vulnerability scheduled scan policy on a periodic basis.
	//
	// example:
	//
	// Instant
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	// The user group ID. This parameter is used to filter records whose effective scope includes the specified user group. Valid values are obtained from:
	//
	// - [ListUserGroups](~~ListUserGroups~~): lists user groups.
	//
	// example:
	//
	// usergroup-9d4f2a7b3c1e****
	UserGroupId *string `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
}

func (s ListVulScanTasksRequest) String() string {
	return dara.Prettify(s)
}

func (s ListVulScanTasksRequest) GoString() string {
	return s.String()
}

func (s *ListVulScanTasksRequest) GetCurrentPage() *int64 {
	return s.CurrentPage
}

func (s *ListVulScanTasksRequest) GetMatchMode() *string {
	return s.MatchMode
}

func (s *ListVulScanTasksRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListVulScanTasksRequest) GetScheduledStrategyId() *string {
	return s.ScheduledStrategyId
}

func (s *ListVulScanTasksRequest) GetStatus() *string {
	return s.Status
}

func (s *ListVulScanTasksRequest) GetTaskIds() []*string {
	return s.TaskIds
}

func (s *ListVulScanTasksRequest) GetTaskName() *string {
	return s.TaskName
}

func (s *ListVulScanTasksRequest) GetTaskType() *string {
	return s.TaskType
}

func (s *ListVulScanTasksRequest) GetUserGroupId() *string {
	return s.UserGroupId
}

func (s *ListVulScanTasksRequest) SetCurrentPage(v int64) *ListVulScanTasksRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListVulScanTasksRequest) SetMatchMode(v string) *ListVulScanTasksRequest {
	s.MatchMode = &v
	return s
}

func (s *ListVulScanTasksRequest) SetPageSize(v int64) *ListVulScanTasksRequest {
	s.PageSize = &v
	return s
}

func (s *ListVulScanTasksRequest) SetScheduledStrategyId(v string) *ListVulScanTasksRequest {
	s.ScheduledStrategyId = &v
	return s
}

func (s *ListVulScanTasksRequest) SetStatus(v string) *ListVulScanTasksRequest {
	s.Status = &v
	return s
}

func (s *ListVulScanTasksRequest) SetTaskIds(v []*string) *ListVulScanTasksRequest {
	s.TaskIds = v
	return s
}

func (s *ListVulScanTasksRequest) SetTaskName(v string) *ListVulScanTasksRequest {
	s.TaskName = &v
	return s
}

func (s *ListVulScanTasksRequest) SetTaskType(v string) *ListVulScanTasksRequest {
	s.TaskType = &v
	return s
}

func (s *ListVulScanTasksRequest) SetUserGroupId(v string) *ListVulScanTasksRequest {
	s.UserGroupId = &v
	return s
}

func (s *ListVulScanTasksRequest) Validate() error {
	return dara.Validate(s)
}
