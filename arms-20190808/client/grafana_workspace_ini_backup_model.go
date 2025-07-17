// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGrafanaWorkspaceIniBackup interface {
	dara.Model
	String() string
	GoString() string
	SetExt(v string) *GrafanaWorkspaceIniBackup
	GetExt() *string
	SetGmtCreate(v int64) *GrafanaWorkspaceIniBackup
	GetGmtCreate() *int64
	SetGmtModified(v int64) *GrafanaWorkspaceIniBackup
	GetGmtModified() *int64
	SetGrafanaWorkspaceId(v string) *GrafanaWorkspaceIniBackup
	GetGrafanaWorkspaceId() *string
	SetId(v int64) *GrafanaWorkspaceIniBackup
	GetId() *int64
	SetMsg(v string) *GrafanaWorkspaceIniBackup
	GetMsg() *string
	SetProcessName(v string) *GrafanaWorkspaceIniBackup
	GetProcessName() *string
	SetProcessStatus(v string) *GrafanaWorkspaceIniBackup
	GetProcessStatus() *string
	SetSubType(v string) *GrafanaWorkspaceIniBackup
	GetSubType() *string
	SetUserId(v string) *GrafanaWorkspaceIniBackup
	GetUserId() *string
}

type GrafanaWorkspaceIniBackup struct {
	Ext *string `json:"ext,omitempty" xml:"ext,omitempty"`
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

func (s GrafanaWorkspaceIniBackup) String() string {
	return dara.Prettify(s)
}

func (s GrafanaWorkspaceIniBackup) GoString() string {
	return s.String()
}

func (s *GrafanaWorkspaceIniBackup) GetExt() *string {
	return s.Ext
}

func (s *GrafanaWorkspaceIniBackup) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *GrafanaWorkspaceIniBackup) GetGmtModified() *int64 {
	return s.GmtModified
}

func (s *GrafanaWorkspaceIniBackup) GetGrafanaWorkspaceId() *string {
	return s.GrafanaWorkspaceId
}

func (s *GrafanaWorkspaceIniBackup) GetId() *int64 {
	return s.Id
}

func (s *GrafanaWorkspaceIniBackup) GetMsg() *string {
	return s.Msg
}

func (s *GrafanaWorkspaceIniBackup) GetProcessName() *string {
	return s.ProcessName
}

func (s *GrafanaWorkspaceIniBackup) GetProcessStatus() *string {
	return s.ProcessStatus
}

func (s *GrafanaWorkspaceIniBackup) GetSubType() *string {
	return s.SubType
}

func (s *GrafanaWorkspaceIniBackup) GetUserId() *string {
	return s.UserId
}

func (s *GrafanaWorkspaceIniBackup) SetExt(v string) *GrafanaWorkspaceIniBackup {
	s.Ext = &v
	return s
}

func (s *GrafanaWorkspaceIniBackup) SetGmtCreate(v int64) *GrafanaWorkspaceIniBackup {
	s.GmtCreate = &v
	return s
}

func (s *GrafanaWorkspaceIniBackup) SetGmtModified(v int64) *GrafanaWorkspaceIniBackup {
	s.GmtModified = &v
	return s
}

func (s *GrafanaWorkspaceIniBackup) SetGrafanaWorkspaceId(v string) *GrafanaWorkspaceIniBackup {
	s.GrafanaWorkspaceId = &v
	return s
}

func (s *GrafanaWorkspaceIniBackup) SetId(v int64) *GrafanaWorkspaceIniBackup {
	s.Id = &v
	return s
}

func (s *GrafanaWorkspaceIniBackup) SetMsg(v string) *GrafanaWorkspaceIniBackup {
	s.Msg = &v
	return s
}

func (s *GrafanaWorkspaceIniBackup) SetProcessName(v string) *GrafanaWorkspaceIniBackup {
	s.ProcessName = &v
	return s
}

func (s *GrafanaWorkspaceIniBackup) SetProcessStatus(v string) *GrafanaWorkspaceIniBackup {
	s.ProcessStatus = &v
	return s
}

func (s *GrafanaWorkspaceIniBackup) SetSubType(v string) *GrafanaWorkspaceIniBackup {
	s.SubType = &v
	return s
}

func (s *GrafanaWorkspaceIniBackup) SetUserId(v string) *GrafanaWorkspaceIniBackup {
	s.UserId = &v
	return s
}

func (s *GrafanaWorkspaceIniBackup) Validate() error {
	return dara.Validate(s)
}
