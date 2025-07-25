// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueueInfo interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueueInfo
	GetCode() *string
	SetCodeType(v string) *QueueInfo
	GetCodeType() *string
	SetGmtCreatedTime(v string) *QueueInfo
	GetGmtCreatedTime() *string
	SetGmtDequeuedTime(v string) *QueueInfo
	GetGmtDequeuedTime() *string
	SetGmtEnqueuedTime(v string) *QueueInfo
	GetGmtEnqueuedTime() *string
	SetGmtPositionModifiedTime(v string) *QueueInfo
	GetGmtPositionModifiedTime() *string
	SetName(v string) *QueueInfo
	GetName() *string
	SetPosition(v int64) *QueueInfo
	GetPosition() *int64
	SetPriority(v int64) *QueueInfo
	GetPriority() *int64
	SetQueueStrategy(v string) *QueueInfo
	GetQueueStrategy() *string
	SetQuotaId(v string) *QueueInfo
	GetQuotaId() *string
	SetReason(v string) *QueueInfo
	GetReason() *string
	SetResource(v *ResourceAmount) *QueueInfo
	GetResource() *ResourceAmount
	SetStatus(v string) *QueueInfo
	GetStatus() *string
	SetUseOversoldResource(v bool) *QueueInfo
	GetUseOversoldResource() *bool
	SetUserId(v string) *QueueInfo
	GetUserId() *string
	SetUserName(v string) *QueueInfo
	GetUserName() *string
	SetWorkloadId(v string) *QueueInfo
	GetWorkloadId() *string
	SetWorkloadName(v string) *QueueInfo
	GetWorkloadName() *string
	SetWorkloadStatus(v string) *QueueInfo
	GetWorkloadStatus() *string
	SetWorkloadType(v string) *QueueInfo
	GetWorkloadType() *string
	SetWorkspaceId(v string) *QueueInfo
	GetWorkspaceId() *string
	SetWorkspaceName(v string) *QueueInfo
	GetWorkspaceName() *string
}

type QueueInfo struct {
	// example:
	//
	// roleMaximumResource
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// ConfigRule
	CodeType       *string `json:"CodeType,omitempty" xml:"CodeType,omitempty"`
	GmtCreatedTime *string `json:"GmtCreatedTime,omitempty" xml:"GmtCreatedTime,omitempty"`
	// example:
	//
	// "2023-06-22T00:00:00Z"
	GmtDequeuedTime *string `json:"GmtDequeuedTime,omitempty" xml:"GmtDequeuedTime,omitempty"`
	// example:
	//
	// “2023-06-22T00:00:00Z”
	GmtEnqueuedTime *string `json:"GmtEnqueuedTime,omitempty" xml:"GmtEnqueuedTime,omitempty"`
	// example:
	//
	// "2023-06-22T00:00:00Z"
	GmtPositionModifiedTime *string `json:"GmtPositionModifiedTime,omitempty" xml:"GmtPositionModifiedTime,omitempty"`
	// example:
	//
	// test-label-79f5498dd-9qrzs
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 10
	Position *int64 `json:"Position,omitempty" xml:"Position,omitempty"`
	// example:
	//
	// 2
	Priority *int64 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// example:
	//
	// PaiStrategyIntelligent
	QueueStrategy *string `json:"QueueStrategy,omitempty" xml:"QueueStrategy,omitempty"`
	// example:
	//
	// “quotamtl37ge7gkvdz”
	QuotaId *string `json:"QuotaId,omitempty" xml:"QuotaId,omitempty"`
	// example:
	//
	// Current GPU Limit is 5, limited by Role PAI.AlgoDeveloper
	Reason   *string         `json:"Reason,omitempty" xml:"Reason,omitempty"`
	Resource *ResourceAmount `json:"Resource,omitempty" xml:"Resource,omitempty"`
	// example:
	//
	// Enqueued
	Status              *string `json:"Status,omitempty" xml:"Status,omitempty"`
	UseOversoldResource *bool   `json:"UseOversoldResource,omitempty" xml:"UseOversoldResource,omitempty"`
	UserId              *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	UserName            *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	// example:
	//
	// dlcxxxx
	WorkloadId   *string `json:"WorkloadId,omitempty" xml:"WorkloadId,omitempty"`
	WorkloadName *string `json:"WorkloadName,omitempty" xml:"WorkloadName,omitempty"`
	// example:
	//
	// PreAllocation
	WorkloadStatus *string `json:"WorkloadStatus,omitempty" xml:"WorkloadStatus,omitempty"`
	// example:
	//
	// dlc
	WorkloadType *string `json:"WorkloadType,omitempty" xml:"WorkloadType,omitempty"`
	// example:
	//
	// “432524”
	WorkspaceId   *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
	WorkspaceName *string `json:"WorkspaceName,omitempty" xml:"WorkspaceName,omitempty"`
}

func (s QueueInfo) String() string {
	return dara.Prettify(s)
}

func (s QueueInfo) GoString() string {
	return s.String()
}

func (s *QueueInfo) GetCode() *string {
	return s.Code
}

func (s *QueueInfo) GetCodeType() *string {
	return s.CodeType
}

func (s *QueueInfo) GetGmtCreatedTime() *string {
	return s.GmtCreatedTime
}

func (s *QueueInfo) GetGmtDequeuedTime() *string {
	return s.GmtDequeuedTime
}

func (s *QueueInfo) GetGmtEnqueuedTime() *string {
	return s.GmtEnqueuedTime
}

func (s *QueueInfo) GetGmtPositionModifiedTime() *string {
	return s.GmtPositionModifiedTime
}

func (s *QueueInfo) GetName() *string {
	return s.Name
}

func (s *QueueInfo) GetPosition() *int64 {
	return s.Position
}

func (s *QueueInfo) GetPriority() *int64 {
	return s.Priority
}

func (s *QueueInfo) GetQueueStrategy() *string {
	return s.QueueStrategy
}

func (s *QueueInfo) GetQuotaId() *string {
	return s.QuotaId
}

func (s *QueueInfo) GetReason() *string {
	return s.Reason
}

func (s *QueueInfo) GetResource() *ResourceAmount {
	return s.Resource
}

func (s *QueueInfo) GetStatus() *string {
	return s.Status
}

func (s *QueueInfo) GetUseOversoldResource() *bool {
	return s.UseOversoldResource
}

func (s *QueueInfo) GetUserId() *string {
	return s.UserId
}

func (s *QueueInfo) GetUserName() *string {
	return s.UserName
}

func (s *QueueInfo) GetWorkloadId() *string {
	return s.WorkloadId
}

func (s *QueueInfo) GetWorkloadName() *string {
	return s.WorkloadName
}

func (s *QueueInfo) GetWorkloadStatus() *string {
	return s.WorkloadStatus
}

func (s *QueueInfo) GetWorkloadType() *string {
	return s.WorkloadType
}

func (s *QueueInfo) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *QueueInfo) GetWorkspaceName() *string {
	return s.WorkspaceName
}

func (s *QueueInfo) SetCode(v string) *QueueInfo {
	s.Code = &v
	return s
}

func (s *QueueInfo) SetCodeType(v string) *QueueInfo {
	s.CodeType = &v
	return s
}

func (s *QueueInfo) SetGmtCreatedTime(v string) *QueueInfo {
	s.GmtCreatedTime = &v
	return s
}

func (s *QueueInfo) SetGmtDequeuedTime(v string) *QueueInfo {
	s.GmtDequeuedTime = &v
	return s
}

func (s *QueueInfo) SetGmtEnqueuedTime(v string) *QueueInfo {
	s.GmtEnqueuedTime = &v
	return s
}

func (s *QueueInfo) SetGmtPositionModifiedTime(v string) *QueueInfo {
	s.GmtPositionModifiedTime = &v
	return s
}

func (s *QueueInfo) SetName(v string) *QueueInfo {
	s.Name = &v
	return s
}

func (s *QueueInfo) SetPosition(v int64) *QueueInfo {
	s.Position = &v
	return s
}

func (s *QueueInfo) SetPriority(v int64) *QueueInfo {
	s.Priority = &v
	return s
}

func (s *QueueInfo) SetQueueStrategy(v string) *QueueInfo {
	s.QueueStrategy = &v
	return s
}

func (s *QueueInfo) SetQuotaId(v string) *QueueInfo {
	s.QuotaId = &v
	return s
}

func (s *QueueInfo) SetReason(v string) *QueueInfo {
	s.Reason = &v
	return s
}

func (s *QueueInfo) SetResource(v *ResourceAmount) *QueueInfo {
	s.Resource = v
	return s
}

func (s *QueueInfo) SetStatus(v string) *QueueInfo {
	s.Status = &v
	return s
}

func (s *QueueInfo) SetUseOversoldResource(v bool) *QueueInfo {
	s.UseOversoldResource = &v
	return s
}

func (s *QueueInfo) SetUserId(v string) *QueueInfo {
	s.UserId = &v
	return s
}

func (s *QueueInfo) SetUserName(v string) *QueueInfo {
	s.UserName = &v
	return s
}

func (s *QueueInfo) SetWorkloadId(v string) *QueueInfo {
	s.WorkloadId = &v
	return s
}

func (s *QueueInfo) SetWorkloadName(v string) *QueueInfo {
	s.WorkloadName = &v
	return s
}

func (s *QueueInfo) SetWorkloadStatus(v string) *QueueInfo {
	s.WorkloadStatus = &v
	return s
}

func (s *QueueInfo) SetWorkloadType(v string) *QueueInfo {
	s.WorkloadType = &v
	return s
}

func (s *QueueInfo) SetWorkspaceId(v string) *QueueInfo {
	s.WorkspaceId = &v
	return s
}

func (s *QueueInfo) SetWorkspaceName(v string) *QueueInfo {
	s.WorkspaceName = &v
	return s
}

func (s *QueueInfo) Validate() error {
	return dara.Validate(s)
}
