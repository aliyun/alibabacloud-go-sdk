// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iWorkloadInfo interface {
	dara.Model
	String() string
	GoString() string
	SetIsScheduled(v string) *WorkloadInfo
	GetIsScheduled() *string
	SetPriority(v int32) *WorkloadInfo
	GetPriority() *int32
	SetQueueMetas(v []*QueueMeta) *WorkloadInfo
	GetQueueMetas() []*QueueMeta
	SetTenantId(v string) *WorkloadInfo
	GetTenantId() *string
	SetUserId(v string) *WorkloadInfo
	GetUserId() *string
	SetUserName(v string) *WorkloadInfo
	GetUserName() *string
	SetWorkloadCreatedTime(v string) *WorkloadInfo
	GetWorkloadCreatedTime() *string
	SetWorkloadId(v string) *WorkloadInfo
	GetWorkloadId() *string
	SetWorkloadName(v string) *WorkloadInfo
	GetWorkloadName() *string
	SetWorkloadStatus(v string) *WorkloadInfo
	GetWorkloadStatus() *string
	SetWorkloadType(v string) *WorkloadInfo
	GetWorkloadType() *string
	SetWorkspaceId(v string) *WorkloadInfo
	GetWorkspaceId() *string
	SetWorkspaceName(v string) *WorkloadInfo
	GetWorkspaceName() *string
}

type WorkloadInfo struct {
	// Specifies whether the task is scheduled.
	IsScheduled *string `json:"IsScheduled,omitempty" xml:"IsScheduled,omitempty"`
	// The task priority.
	//
	// example:
	//
	// 0
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The information about the queue where the task resides.
	QueueMetas []*QueueMeta `json:"QueueMetas,omitempty" xml:"QueueMetas,omitempty" type:"Repeated"`
	// The ID of the Alibaba Cloud account that created the task.
	//
	// example:
	//
	// 123456789
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// The ID of the RAM user that created the task.
	//
	// example:
	//
	// 123456789
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The username of the task creator.
	//
	// example:
	//
	// test
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	// The time when the task was created.
	//
	// example:
	//
	// 2023-06-22T00:00:00Z
	WorkloadCreatedTime *string `json:"WorkloadCreatedTime,omitempty" xml:"WorkloadCreatedTime,omitempty"`
	// The task ID.
	//
	// example:
	//
	// dlcxxxx
	WorkloadId *string `json:"WorkloadId,omitempty" xml:"WorkloadId,omitempty"`
	// The task name.
	//
	// example:
	//
	// test-job
	WorkloadName *string `json:"WorkloadName,omitempty" xml:"WorkloadName,omitempty"`
	// The task status.
	//
	// example:
	//
	// Running
	WorkloadStatus *string `json:"WorkloadStatus,omitempty" xml:"WorkloadStatus,omitempty"`
	// The task type.
	//
	// example:
	//
	// dlc
	WorkloadType *string `json:"WorkloadType,omitempty" xml:"WorkloadType,omitempty"`
	// The ID of the workspace where the task resides.
	//
	// example:
	//
	// 12345
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
	// The name of the workspace where the task resides.
	//
	// example:
	//
	// myshellprod
	WorkspaceName *string `json:"WorkspaceName,omitempty" xml:"WorkspaceName,omitempty"`
}

func (s WorkloadInfo) String() string {
	return dara.Prettify(s)
}

func (s WorkloadInfo) GoString() string {
	return s.String()
}

func (s *WorkloadInfo) GetIsScheduled() *string {
	return s.IsScheduled
}

func (s *WorkloadInfo) GetPriority() *int32 {
	return s.Priority
}

func (s *WorkloadInfo) GetQueueMetas() []*QueueMeta {
	return s.QueueMetas
}

func (s *WorkloadInfo) GetTenantId() *string {
	return s.TenantId
}

func (s *WorkloadInfo) GetUserId() *string {
	return s.UserId
}

func (s *WorkloadInfo) GetUserName() *string {
	return s.UserName
}

func (s *WorkloadInfo) GetWorkloadCreatedTime() *string {
	return s.WorkloadCreatedTime
}

func (s *WorkloadInfo) GetWorkloadId() *string {
	return s.WorkloadId
}

func (s *WorkloadInfo) GetWorkloadName() *string {
	return s.WorkloadName
}

func (s *WorkloadInfo) GetWorkloadStatus() *string {
	return s.WorkloadStatus
}

func (s *WorkloadInfo) GetWorkloadType() *string {
	return s.WorkloadType
}

func (s *WorkloadInfo) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *WorkloadInfo) GetWorkspaceName() *string {
	return s.WorkspaceName
}

func (s *WorkloadInfo) SetIsScheduled(v string) *WorkloadInfo {
	s.IsScheduled = &v
	return s
}

func (s *WorkloadInfo) SetPriority(v int32) *WorkloadInfo {
	s.Priority = &v
	return s
}

func (s *WorkloadInfo) SetQueueMetas(v []*QueueMeta) *WorkloadInfo {
	s.QueueMetas = v
	return s
}

func (s *WorkloadInfo) SetTenantId(v string) *WorkloadInfo {
	s.TenantId = &v
	return s
}

func (s *WorkloadInfo) SetUserId(v string) *WorkloadInfo {
	s.UserId = &v
	return s
}

func (s *WorkloadInfo) SetUserName(v string) *WorkloadInfo {
	s.UserName = &v
	return s
}

func (s *WorkloadInfo) SetWorkloadCreatedTime(v string) *WorkloadInfo {
	s.WorkloadCreatedTime = &v
	return s
}

func (s *WorkloadInfo) SetWorkloadId(v string) *WorkloadInfo {
	s.WorkloadId = &v
	return s
}

func (s *WorkloadInfo) SetWorkloadName(v string) *WorkloadInfo {
	s.WorkloadName = &v
	return s
}

func (s *WorkloadInfo) SetWorkloadStatus(v string) *WorkloadInfo {
	s.WorkloadStatus = &v
	return s
}

func (s *WorkloadInfo) SetWorkloadType(v string) *WorkloadInfo {
	s.WorkloadType = &v
	return s
}

func (s *WorkloadInfo) SetWorkspaceId(v string) *WorkloadInfo {
	s.WorkspaceId = &v
	return s
}

func (s *WorkloadInfo) SetWorkspaceName(v string) *WorkloadInfo {
	s.WorkspaceName = &v
	return s
}

func (s *WorkloadInfo) Validate() error {
	if s.QueueMetas != nil {
		for _, item := range s.QueueMetas {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
