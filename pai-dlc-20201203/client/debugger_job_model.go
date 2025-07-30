// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDebuggerJob interface {
	dara.Model
	String() string
	GoString() string
	SetDebuggerJobId(v string) *DebuggerJob
	GetDebuggerJobId() *string
	SetDisplayName(v string) *DebuggerJob
	GetDisplayName() *string
	SetDuration(v string) *DebuggerJob
	GetDuration() *string
	SetGmtCreateTime(v string) *DebuggerJob
	GetGmtCreateTime() *string
	SetGmtFailedTime(v string) *DebuggerJob
	GetGmtFailedTime() *string
	SetGmtFinishTime(v string) *DebuggerJob
	GetGmtFinishTime() *string
	SetGmtRunningTime(v string) *DebuggerJob
	GetGmtRunningTime() *string
	SetGmtStoppedTime(v string) *DebuggerJob
	GetGmtStoppedTime() *string
	SetGmtSubmittedTime(v string) *DebuggerJob
	GetGmtSubmittedTime() *string
	SetGmtSucceedTime(v string) *DebuggerJob
	GetGmtSucceedTime() *string
	SetStatus(v string) *DebuggerJob
	GetStatus() *string
	SetUserId(v string) *DebuggerJob
	GetUserId() *string
	SetWorkspaceId(v string) *DebuggerJob
	GetWorkspaceId() *string
	SetWorkspaceName(v string) *DebuggerJob
	GetWorkspaceName() *string
}

type DebuggerJob struct {
	// example:
	//
	// dlc20210126170216-mtl37ge7gkvdz
	DebuggerJobId *string `json:"DebuggerJobId,omitempty" xml:"DebuggerJobId,omitempty"`
	// example:
	//
	// dlc debugger test
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// example:
	//
	// 2932
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// example:
	//
	// 2021-01-12T14:35:00Z
	GmtCreateTime    *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	GmtFailedTime    *string `json:"GmtFailedTime,omitempty" xml:"GmtFailedTime,omitempty"`
	GmtFinishTime    *string `json:"GmtFinishTime,omitempty" xml:"GmtFinishTime,omitempty"`
	GmtRunningTime   *string `json:"GmtRunningTime,omitempty" xml:"GmtRunningTime,omitempty"`
	GmtStoppedTime   *string `json:"GmtStoppedTime,omitempty" xml:"GmtStoppedTime,omitempty"`
	GmtSubmittedTime *string `json:"GmtSubmittedTime,omitempty" xml:"GmtSubmittedTime,omitempty"`
	GmtSucceedTime   *string `json:"GmtSucceedTime,omitempty" xml:"GmtSucceedTime,omitempty"`
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 12344556
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// example:
	//
	// workspace01
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
	// example:
	//
	// public
	WorkspaceName *string `json:"WorkspaceName,omitempty" xml:"WorkspaceName,omitempty"`
}

func (s DebuggerJob) String() string {
	return dara.Prettify(s)
}

func (s DebuggerJob) GoString() string {
	return s.String()
}

func (s *DebuggerJob) GetDebuggerJobId() *string {
	return s.DebuggerJobId
}

func (s *DebuggerJob) GetDisplayName() *string {
	return s.DisplayName
}

func (s *DebuggerJob) GetDuration() *string {
	return s.Duration
}

func (s *DebuggerJob) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *DebuggerJob) GetGmtFailedTime() *string {
	return s.GmtFailedTime
}

func (s *DebuggerJob) GetGmtFinishTime() *string {
	return s.GmtFinishTime
}

func (s *DebuggerJob) GetGmtRunningTime() *string {
	return s.GmtRunningTime
}

func (s *DebuggerJob) GetGmtStoppedTime() *string {
	return s.GmtStoppedTime
}

func (s *DebuggerJob) GetGmtSubmittedTime() *string {
	return s.GmtSubmittedTime
}

func (s *DebuggerJob) GetGmtSucceedTime() *string {
	return s.GmtSucceedTime
}

func (s *DebuggerJob) GetStatus() *string {
	return s.Status
}

func (s *DebuggerJob) GetUserId() *string {
	return s.UserId
}

func (s *DebuggerJob) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *DebuggerJob) GetWorkspaceName() *string {
	return s.WorkspaceName
}

func (s *DebuggerJob) SetDebuggerJobId(v string) *DebuggerJob {
	s.DebuggerJobId = &v
	return s
}

func (s *DebuggerJob) SetDisplayName(v string) *DebuggerJob {
	s.DisplayName = &v
	return s
}

func (s *DebuggerJob) SetDuration(v string) *DebuggerJob {
	s.Duration = &v
	return s
}

func (s *DebuggerJob) SetGmtCreateTime(v string) *DebuggerJob {
	s.GmtCreateTime = &v
	return s
}

func (s *DebuggerJob) SetGmtFailedTime(v string) *DebuggerJob {
	s.GmtFailedTime = &v
	return s
}

func (s *DebuggerJob) SetGmtFinishTime(v string) *DebuggerJob {
	s.GmtFinishTime = &v
	return s
}

func (s *DebuggerJob) SetGmtRunningTime(v string) *DebuggerJob {
	s.GmtRunningTime = &v
	return s
}

func (s *DebuggerJob) SetGmtStoppedTime(v string) *DebuggerJob {
	s.GmtStoppedTime = &v
	return s
}

func (s *DebuggerJob) SetGmtSubmittedTime(v string) *DebuggerJob {
	s.GmtSubmittedTime = &v
	return s
}

func (s *DebuggerJob) SetGmtSucceedTime(v string) *DebuggerJob {
	s.GmtSucceedTime = &v
	return s
}

func (s *DebuggerJob) SetStatus(v string) *DebuggerJob {
	s.Status = &v
	return s
}

func (s *DebuggerJob) SetUserId(v string) *DebuggerJob {
	s.UserId = &v
	return s
}

func (s *DebuggerJob) SetWorkspaceId(v string) *DebuggerJob {
	s.WorkspaceId = &v
	return s
}

func (s *DebuggerJob) SetWorkspaceName(v string) *DebuggerJob {
	s.WorkspaceName = &v
	return s
}

func (s *DebuggerJob) Validate() error {
	return dara.Validate(s)
}
