// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAgenticDmsInstanceSyncTask interface {
	dara.Model
	String() string
	GoString() string
	SetActorId(v string) *AgenticDmsInstanceSyncTask
	GetActorId() *string
	SetActorName(v string) *AgenticDmsInstanceSyncTask
	GetActorName() *string
	SetActorType(v string) *AgenticDmsInstanceSyncTask
	GetActorType() *string
	SetErrorCode(v string) *AgenticDmsInstanceSyncTask
	GetErrorCode() *string
	SetErrorSummary(v string) *AgenticDmsInstanceSyncTask
	GetErrorSummary() *string
	SetFailedCount(v int32) *AgenticDmsInstanceSyncTask
	GetFailedCount() *int32
	SetGmtCreate(v string) *AgenticDmsInstanceSyncTask
	GetGmtCreate() *string
	SetGmtModified(v string) *AgenticDmsInstanceSyncTask
	GetGmtModified() *string
	SetSkippedCount(v int32) *AgenticDmsInstanceSyncTask
	GetSkippedCount() *int32
	SetStatus(v string) *AgenticDmsInstanceSyncTask
	GetStatus() *string
	SetSuccessCount(v int32) *AgenticDmsInstanceSyncTask
	GetSuccessCount() *int32
	SetSyncUserDataPermission(v bool) *AgenticDmsInstanceSyncTask
	GetSyncUserDataPermission() *bool
	SetTaskId(v string) *AgenticDmsInstanceSyncTask
	GetTaskId() *string
	SetTotalCount(v int32) *AgenticDmsInstanceSyncTask
	GetTotalCount() *int32
}

type AgenticDmsInstanceSyncTask struct {
	ActorId                *string `json:"ActorId,omitempty" xml:"ActorId,omitempty"`
	ActorName              *string `json:"ActorName,omitempty" xml:"ActorName,omitempty"`
	ActorType              *string `json:"ActorType,omitempty" xml:"ActorType,omitempty"`
	ErrorCode              *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorSummary           *string `json:"ErrorSummary,omitempty" xml:"ErrorSummary,omitempty"`
	FailedCount            *int32  `json:"FailedCount,omitempty" xml:"FailedCount,omitempty"`
	GmtCreate              *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified            *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	SkippedCount           *int32  `json:"SkippedCount,omitempty" xml:"SkippedCount,omitempty"`
	Status                 *string `json:"Status,omitempty" xml:"Status,omitempty"`
	SuccessCount           *int32  `json:"SuccessCount,omitempty" xml:"SuccessCount,omitempty"`
	SyncUserDataPermission *bool   `json:"SyncUserDataPermission,omitempty" xml:"SyncUserDataPermission,omitempty"`
	TaskId                 *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	TotalCount             *int32  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s AgenticDmsInstanceSyncTask) String() string {
	return dara.Prettify(s)
}

func (s AgenticDmsInstanceSyncTask) GoString() string {
	return s.String()
}

func (s *AgenticDmsInstanceSyncTask) GetActorId() *string {
	return s.ActorId
}

func (s *AgenticDmsInstanceSyncTask) GetActorName() *string {
	return s.ActorName
}

func (s *AgenticDmsInstanceSyncTask) GetActorType() *string {
	return s.ActorType
}

func (s *AgenticDmsInstanceSyncTask) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *AgenticDmsInstanceSyncTask) GetErrorSummary() *string {
	return s.ErrorSummary
}

func (s *AgenticDmsInstanceSyncTask) GetFailedCount() *int32 {
	return s.FailedCount
}

func (s *AgenticDmsInstanceSyncTask) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *AgenticDmsInstanceSyncTask) GetGmtModified() *string {
	return s.GmtModified
}

func (s *AgenticDmsInstanceSyncTask) GetSkippedCount() *int32 {
	return s.SkippedCount
}

func (s *AgenticDmsInstanceSyncTask) GetStatus() *string {
	return s.Status
}

func (s *AgenticDmsInstanceSyncTask) GetSuccessCount() *int32 {
	return s.SuccessCount
}

func (s *AgenticDmsInstanceSyncTask) GetSyncUserDataPermission() *bool {
	return s.SyncUserDataPermission
}

func (s *AgenticDmsInstanceSyncTask) GetTaskId() *string {
	return s.TaskId
}

func (s *AgenticDmsInstanceSyncTask) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *AgenticDmsInstanceSyncTask) SetActorId(v string) *AgenticDmsInstanceSyncTask {
	s.ActorId = &v
	return s
}

func (s *AgenticDmsInstanceSyncTask) SetActorName(v string) *AgenticDmsInstanceSyncTask {
	s.ActorName = &v
	return s
}

func (s *AgenticDmsInstanceSyncTask) SetActorType(v string) *AgenticDmsInstanceSyncTask {
	s.ActorType = &v
	return s
}

func (s *AgenticDmsInstanceSyncTask) SetErrorCode(v string) *AgenticDmsInstanceSyncTask {
	s.ErrorCode = &v
	return s
}

func (s *AgenticDmsInstanceSyncTask) SetErrorSummary(v string) *AgenticDmsInstanceSyncTask {
	s.ErrorSummary = &v
	return s
}

func (s *AgenticDmsInstanceSyncTask) SetFailedCount(v int32) *AgenticDmsInstanceSyncTask {
	s.FailedCount = &v
	return s
}

func (s *AgenticDmsInstanceSyncTask) SetGmtCreate(v string) *AgenticDmsInstanceSyncTask {
	s.GmtCreate = &v
	return s
}

func (s *AgenticDmsInstanceSyncTask) SetGmtModified(v string) *AgenticDmsInstanceSyncTask {
	s.GmtModified = &v
	return s
}

func (s *AgenticDmsInstanceSyncTask) SetSkippedCount(v int32) *AgenticDmsInstanceSyncTask {
	s.SkippedCount = &v
	return s
}

func (s *AgenticDmsInstanceSyncTask) SetStatus(v string) *AgenticDmsInstanceSyncTask {
	s.Status = &v
	return s
}

func (s *AgenticDmsInstanceSyncTask) SetSuccessCount(v int32) *AgenticDmsInstanceSyncTask {
	s.SuccessCount = &v
	return s
}

func (s *AgenticDmsInstanceSyncTask) SetSyncUserDataPermission(v bool) *AgenticDmsInstanceSyncTask {
	s.SyncUserDataPermission = &v
	return s
}

func (s *AgenticDmsInstanceSyncTask) SetTaskId(v string) *AgenticDmsInstanceSyncTask {
	s.TaskId = &v
	return s
}

func (s *AgenticDmsInstanceSyncTask) SetTotalCount(v int32) *AgenticDmsInstanceSyncTask {
	s.TotalCount = &v
	return s
}

func (s *AgenticDmsInstanceSyncTask) Validate() error {
	return dara.Validate(s)
}
