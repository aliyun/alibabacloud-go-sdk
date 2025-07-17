// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGrafanaWorkspaceTrans interface {
	dara.Model
	String() string
	GoString() string
	SetApiUrl(v string) *GrafanaWorkspaceTrans
	GetApiUrl() *string
	SetAuthType(v string) *GrafanaWorkspaceTrans
	GetAuthType() *string
	SetGmtCreate(v float32) *GrafanaWorkspaceTrans
	GetGmtCreate() *float32
	SetGmtModified(v float32) *GrafanaWorkspaceTrans
	GetGmtModified() *float32
	SetGrafanaWorkspaceId(v string) *GrafanaWorkspaceTrans
	GetGrafanaWorkspaceId() *string
	SetId(v int64) *GrafanaWorkspaceTrans
	GetId() *int64
	SetMsg(v string) *GrafanaWorkspaceTrans
	GetMsg() *string
	SetProcessStatus(v string) *GrafanaWorkspaceTrans
	GetProcessStatus() *string
	SetTransDetails(v []*GrafanaWorkspaceTransDetail) *GrafanaWorkspaceTrans
	GetTransDetails() []*GrafanaWorkspaceTransDetail
	SetUserId(v string) *GrafanaWorkspaceTrans
	GetUserId() *string
}

type GrafanaWorkspaceTrans struct {
	ApiUrl             *string                        `json:"apiUrl,omitempty" xml:"apiUrl,omitempty"`
	AuthType           *string                        `json:"authType,omitempty" xml:"authType,omitempty"`
	GmtCreate          *float32                       `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	GmtModified        *float32                       `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	GrafanaWorkspaceId *string                        `json:"grafanaWorkspaceId,omitempty" xml:"grafanaWorkspaceId,omitempty"`
	Id                 *int64                         `json:"id,omitempty" xml:"id,omitempty"`
	Msg                *string                        `json:"msg,omitempty" xml:"msg,omitempty"`
	ProcessStatus      *string                        `json:"processStatus,omitempty" xml:"processStatus,omitempty"`
	TransDetails       []*GrafanaWorkspaceTransDetail `json:"transDetails,omitempty" xml:"transDetails,omitempty" type:"Repeated"`
	UserId             *string                        `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s GrafanaWorkspaceTrans) String() string {
	return dara.Prettify(s)
}

func (s GrafanaWorkspaceTrans) GoString() string {
	return s.String()
}

func (s *GrafanaWorkspaceTrans) GetApiUrl() *string {
	return s.ApiUrl
}

func (s *GrafanaWorkspaceTrans) GetAuthType() *string {
	return s.AuthType
}

func (s *GrafanaWorkspaceTrans) GetGmtCreate() *float32 {
	return s.GmtCreate
}

func (s *GrafanaWorkspaceTrans) GetGmtModified() *float32 {
	return s.GmtModified
}

func (s *GrafanaWorkspaceTrans) GetGrafanaWorkspaceId() *string {
	return s.GrafanaWorkspaceId
}

func (s *GrafanaWorkspaceTrans) GetId() *int64 {
	return s.Id
}

func (s *GrafanaWorkspaceTrans) GetMsg() *string {
	return s.Msg
}

func (s *GrafanaWorkspaceTrans) GetProcessStatus() *string {
	return s.ProcessStatus
}

func (s *GrafanaWorkspaceTrans) GetTransDetails() []*GrafanaWorkspaceTransDetail {
	return s.TransDetails
}

func (s *GrafanaWorkspaceTrans) GetUserId() *string {
	return s.UserId
}

func (s *GrafanaWorkspaceTrans) SetApiUrl(v string) *GrafanaWorkspaceTrans {
	s.ApiUrl = &v
	return s
}

func (s *GrafanaWorkspaceTrans) SetAuthType(v string) *GrafanaWorkspaceTrans {
	s.AuthType = &v
	return s
}

func (s *GrafanaWorkspaceTrans) SetGmtCreate(v float32) *GrafanaWorkspaceTrans {
	s.GmtCreate = &v
	return s
}

func (s *GrafanaWorkspaceTrans) SetGmtModified(v float32) *GrafanaWorkspaceTrans {
	s.GmtModified = &v
	return s
}

func (s *GrafanaWorkspaceTrans) SetGrafanaWorkspaceId(v string) *GrafanaWorkspaceTrans {
	s.GrafanaWorkspaceId = &v
	return s
}

func (s *GrafanaWorkspaceTrans) SetId(v int64) *GrafanaWorkspaceTrans {
	s.Id = &v
	return s
}

func (s *GrafanaWorkspaceTrans) SetMsg(v string) *GrafanaWorkspaceTrans {
	s.Msg = &v
	return s
}

func (s *GrafanaWorkspaceTrans) SetProcessStatus(v string) *GrafanaWorkspaceTrans {
	s.ProcessStatus = &v
	return s
}

func (s *GrafanaWorkspaceTrans) SetTransDetails(v []*GrafanaWorkspaceTransDetail) *GrafanaWorkspaceTrans {
	s.TransDetails = v
	return s
}

func (s *GrafanaWorkspaceTrans) SetUserId(v string) *GrafanaWorkspaceTrans {
	s.UserId = &v
	return s
}

func (s *GrafanaWorkspaceTrans) Validate() error {
	return dara.Validate(s)
}
