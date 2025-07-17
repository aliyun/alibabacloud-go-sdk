// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGrafanaWorkspaceDataBackup interface {
	dara.Model
	String() string
	GoString() string
	SetGmtCreate(v int64) *GrafanaWorkspaceDataBackup
	GetGmtCreate() *int64
	SetGmtModified(v int64) *GrafanaWorkspaceDataBackup
	GetGmtModified() *int64
	SetGrafanaWorkspaceId(v string) *GrafanaWorkspaceDataBackup
	GetGrafanaWorkspaceId() *string
	SetId(v int64) *GrafanaWorkspaceDataBackup
	GetId() *int64
	SetMsg(v string) *GrafanaWorkspaceDataBackup
	GetMsg() *string
	SetProcessName(v string) *GrafanaWorkspaceDataBackup
	GetProcessName() *string
	SetProcessStatus(v string) *GrafanaWorkspaceDataBackup
	GetProcessStatus() *string
	SetSubType(v string) *GrafanaWorkspaceDataBackup
	GetSubType() *string
	SetUserId(v string) *GrafanaWorkspaceDataBackup
	GetUserId() *string
}

type GrafanaWorkspaceDataBackup struct {
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

func (s GrafanaWorkspaceDataBackup) String() string {
	return dara.Prettify(s)
}

func (s GrafanaWorkspaceDataBackup) GoString() string {
	return s.String()
}

func (s *GrafanaWorkspaceDataBackup) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *GrafanaWorkspaceDataBackup) GetGmtModified() *int64 {
	return s.GmtModified
}

func (s *GrafanaWorkspaceDataBackup) GetGrafanaWorkspaceId() *string {
	return s.GrafanaWorkspaceId
}

func (s *GrafanaWorkspaceDataBackup) GetId() *int64 {
	return s.Id
}

func (s *GrafanaWorkspaceDataBackup) GetMsg() *string {
	return s.Msg
}

func (s *GrafanaWorkspaceDataBackup) GetProcessName() *string {
	return s.ProcessName
}

func (s *GrafanaWorkspaceDataBackup) GetProcessStatus() *string {
	return s.ProcessStatus
}

func (s *GrafanaWorkspaceDataBackup) GetSubType() *string {
	return s.SubType
}

func (s *GrafanaWorkspaceDataBackup) GetUserId() *string {
	return s.UserId
}

func (s *GrafanaWorkspaceDataBackup) SetGmtCreate(v int64) *GrafanaWorkspaceDataBackup {
	s.GmtCreate = &v
	return s
}

func (s *GrafanaWorkspaceDataBackup) SetGmtModified(v int64) *GrafanaWorkspaceDataBackup {
	s.GmtModified = &v
	return s
}

func (s *GrafanaWorkspaceDataBackup) SetGrafanaWorkspaceId(v string) *GrafanaWorkspaceDataBackup {
	s.GrafanaWorkspaceId = &v
	return s
}

func (s *GrafanaWorkspaceDataBackup) SetId(v int64) *GrafanaWorkspaceDataBackup {
	s.Id = &v
	return s
}

func (s *GrafanaWorkspaceDataBackup) SetMsg(v string) *GrafanaWorkspaceDataBackup {
	s.Msg = &v
	return s
}

func (s *GrafanaWorkspaceDataBackup) SetProcessName(v string) *GrafanaWorkspaceDataBackup {
	s.ProcessName = &v
	return s
}

func (s *GrafanaWorkspaceDataBackup) SetProcessStatus(v string) *GrafanaWorkspaceDataBackup {
	s.ProcessStatus = &v
	return s
}

func (s *GrafanaWorkspaceDataBackup) SetSubType(v string) *GrafanaWorkspaceDataBackup {
	s.SubType = &v
	return s
}

func (s *GrafanaWorkspaceDataBackup) SetUserId(v string) *GrafanaWorkspaceDataBackup {
	s.UserId = &v
	return s
}

func (s *GrafanaWorkspaceDataBackup) Validate() error {
	return dara.Validate(s)
}
