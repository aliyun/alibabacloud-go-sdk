// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGrafanaWorkspaceMoveProcess interface {
	dara.Model
	String() string
	GoString() string
	SetGmtCreate(v int64) *GrafanaWorkspaceMoveProcess
	GetGmtCreate() *int64
	SetGmtModified(v int64) *GrafanaWorkspaceMoveProcess
	GetGmtModified() *int64
	SetGrafanaWorkspaceId(v string) *GrafanaWorkspaceMoveProcess
	GetGrafanaWorkspaceId() *string
	SetId(v int64) *GrafanaWorkspaceMoveProcess
	GetId() *int64
	SetMsg(v string) *GrafanaWorkspaceMoveProcess
	GetMsg() *string
	SetProcessName(v string) *GrafanaWorkspaceMoveProcess
	GetProcessName() *string
	SetProcessStatus(v string) *GrafanaWorkspaceMoveProcess
	GetProcessStatus() *string
	SetSubType(v string) *GrafanaWorkspaceMoveProcess
	GetSubType() *string
	SetUserId(v string) *GrafanaWorkspaceMoveProcess
	GetUserId() *string
}

type GrafanaWorkspaceMoveProcess struct {
	// example:
	//
	// 1680861352600
	GmtCreate *int64 `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// example:
	//
	// 1680861352600
	GmtModified *int64 `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// example:
	//
	// grafana-cn-**********
	GrafanaWorkspaceId *string `json:"grafanaWorkspaceId,omitempty" xml:"grafanaWorkspaceId,omitempty"`
	// example:
	//
	// 1
	Id            *int64  `json:"id,omitempty" xml:"id,omitempty"`
	Msg           *string `json:"msg,omitempty" xml:"msg,omitempty"`
	ProcessName   *string `json:"processName,omitempty" xml:"processName,omitempty"`
	ProcessStatus *string `json:"processStatus,omitempty" xml:"processStatus,omitempty"`
	// example:
	//
	// User
	SubType *string `json:"subType,omitempty" xml:"subType,omitempty"`
	UserId  *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s GrafanaWorkspaceMoveProcess) String() string {
	return dara.Prettify(s)
}

func (s GrafanaWorkspaceMoveProcess) GoString() string {
	return s.String()
}

func (s *GrafanaWorkspaceMoveProcess) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *GrafanaWorkspaceMoveProcess) GetGmtModified() *int64 {
	return s.GmtModified
}

func (s *GrafanaWorkspaceMoveProcess) GetGrafanaWorkspaceId() *string {
	return s.GrafanaWorkspaceId
}

func (s *GrafanaWorkspaceMoveProcess) GetId() *int64 {
	return s.Id
}

func (s *GrafanaWorkspaceMoveProcess) GetMsg() *string {
	return s.Msg
}

func (s *GrafanaWorkspaceMoveProcess) GetProcessName() *string {
	return s.ProcessName
}

func (s *GrafanaWorkspaceMoveProcess) GetProcessStatus() *string {
	return s.ProcessStatus
}

func (s *GrafanaWorkspaceMoveProcess) GetSubType() *string {
	return s.SubType
}

func (s *GrafanaWorkspaceMoveProcess) GetUserId() *string {
	return s.UserId
}

func (s *GrafanaWorkspaceMoveProcess) SetGmtCreate(v int64) *GrafanaWorkspaceMoveProcess {
	s.GmtCreate = &v
	return s
}

func (s *GrafanaWorkspaceMoveProcess) SetGmtModified(v int64) *GrafanaWorkspaceMoveProcess {
	s.GmtModified = &v
	return s
}

func (s *GrafanaWorkspaceMoveProcess) SetGrafanaWorkspaceId(v string) *GrafanaWorkspaceMoveProcess {
	s.GrafanaWorkspaceId = &v
	return s
}

func (s *GrafanaWorkspaceMoveProcess) SetId(v int64) *GrafanaWorkspaceMoveProcess {
	s.Id = &v
	return s
}

func (s *GrafanaWorkspaceMoveProcess) SetMsg(v string) *GrafanaWorkspaceMoveProcess {
	s.Msg = &v
	return s
}

func (s *GrafanaWorkspaceMoveProcess) SetProcessName(v string) *GrafanaWorkspaceMoveProcess {
	s.ProcessName = &v
	return s
}

func (s *GrafanaWorkspaceMoveProcess) SetProcessStatus(v string) *GrafanaWorkspaceMoveProcess {
	s.ProcessStatus = &v
	return s
}

func (s *GrafanaWorkspaceMoveProcess) SetSubType(v string) *GrafanaWorkspaceMoveProcess {
	s.SubType = &v
	return s
}

func (s *GrafanaWorkspaceMoveProcess) SetUserId(v string) *GrafanaWorkspaceMoveProcess {
	s.UserId = &v
	return s
}

func (s *GrafanaWorkspaceMoveProcess) Validate() error {
	return dara.Validate(s)
}
