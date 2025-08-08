// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMcdpZoneRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *CreateMcdpZoneRequest
	GetAppId() *string
	SetMpaasMappcenterMcdpZoneCreateJsonStr(v string) *CreateMcdpZoneRequest
	GetMpaasMappcenterMcdpZoneCreateJsonStr() *string
	SetTenantId(v string) *CreateMcdpZoneRequest
	GetTenantId() *string
	SetWorkspaceId(v string) *CreateMcdpZoneRequest
	GetWorkspaceId() *string
}

type CreateMcdpZoneRequest struct {
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	MpaasMappcenterMcdpZoneCreateJsonStr *string `json:"MpaasMappcenterMcdpZoneCreateJsonStr,omitempty" xml:"MpaasMappcenterMcdpZoneCreateJsonStr,omitempty"`
	TenantId                             *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	WorkspaceId                          *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s CreateMcdpZoneRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateMcdpZoneRequest) GoString() string {
	return s.String()
}

func (s *CreateMcdpZoneRequest) GetAppId() *string {
	return s.AppId
}

func (s *CreateMcdpZoneRequest) GetMpaasMappcenterMcdpZoneCreateJsonStr() *string {
	return s.MpaasMappcenterMcdpZoneCreateJsonStr
}

func (s *CreateMcdpZoneRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *CreateMcdpZoneRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *CreateMcdpZoneRequest) SetAppId(v string) *CreateMcdpZoneRequest {
	s.AppId = &v
	return s
}

func (s *CreateMcdpZoneRequest) SetMpaasMappcenterMcdpZoneCreateJsonStr(v string) *CreateMcdpZoneRequest {
	s.MpaasMappcenterMcdpZoneCreateJsonStr = &v
	return s
}

func (s *CreateMcdpZoneRequest) SetTenantId(v string) *CreateMcdpZoneRequest {
	s.TenantId = &v
	return s
}

func (s *CreateMcdpZoneRequest) SetWorkspaceId(v string) *CreateMcdpZoneRequest {
	s.WorkspaceId = &v
	return s
}

func (s *CreateMcdpZoneRequest) Validate() error {
	return dara.Validate(s)
}
