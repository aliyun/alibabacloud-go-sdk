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
	// Specifies whether the workload is scheduled. Valid values are `true` and `false`.
	IsScheduled *string `json:"IsScheduled,omitempty" xml:"IsScheduled,omitempty"`
	// The scheduling priority of the workload. A higher value indicates a higher priority.
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// A list of metadata for the queues associated with the workload.
	QueueMetas []*QueueMeta `json:"QueueMetas,omitempty" xml:"QueueMetas,omitempty" type:"Repeated"`
	// The ID of the tenant that contains the workload.
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// The ID of the user who created the workload.
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The name of the user who created the workload.
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	// The creation time of the workload, in UTC format.
	WorkloadCreatedTime *string `json:"WorkloadCreatedTime,omitempty" xml:"WorkloadCreatedTime,omitempty"`
	// The unique, system-generated ID for the workload.
	WorkloadId *string `json:"WorkloadId,omitempty" xml:"WorkloadId,omitempty"`
	// The user-defined name of the workload.
	WorkloadName *string `json:"WorkloadName,omitempty" xml:"WorkloadName,omitempty"`
	// The current status of the workload. Possible values include `Pending`, `Running`, `Succeeded`, and `Failed`.
	WorkloadStatus *string `json:"WorkloadStatus,omitempty" xml:"WorkloadStatus,omitempty"`
	// The type of the workload.
	WorkloadType *string `json:"WorkloadType,omitempty" xml:"WorkloadType,omitempty"`
	// The ID of the workspace that contains the workload.
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
	// The name of the workspace that contains the workload.
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
