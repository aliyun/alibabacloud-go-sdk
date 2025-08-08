// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMcdpZoneRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DeleteMcdpZoneRequest
	GetAppId() *string
	SetMpaasMappcenterMcdpZoneDeleteJsonStr(v string) *DeleteMcdpZoneRequest
	GetMpaasMappcenterMcdpZoneDeleteJsonStr() *string
	SetTenantId(v string) *DeleteMcdpZoneRequest
	GetTenantId() *string
	SetWorkspaceId(v string) *DeleteMcdpZoneRequest
	GetWorkspaceId() *string
}

type DeleteMcdpZoneRequest struct {
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	MpaasMappcenterMcdpZoneDeleteJsonStr *string `json:"MpaasMappcenterMcdpZoneDeleteJsonStr,omitempty" xml:"MpaasMappcenterMcdpZoneDeleteJsonStr,omitempty"`
	TenantId                             *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	WorkspaceId                          *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s DeleteMcdpZoneRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteMcdpZoneRequest) GoString() string {
	return s.String()
}

func (s *DeleteMcdpZoneRequest) GetAppId() *string {
	return s.AppId
}

func (s *DeleteMcdpZoneRequest) GetMpaasMappcenterMcdpZoneDeleteJsonStr() *string {
	return s.MpaasMappcenterMcdpZoneDeleteJsonStr
}

func (s *DeleteMcdpZoneRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *DeleteMcdpZoneRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *DeleteMcdpZoneRequest) SetAppId(v string) *DeleteMcdpZoneRequest {
	s.AppId = &v
	return s
}

func (s *DeleteMcdpZoneRequest) SetMpaasMappcenterMcdpZoneDeleteJsonStr(v string) *DeleteMcdpZoneRequest {
	s.MpaasMappcenterMcdpZoneDeleteJsonStr = &v
	return s
}

func (s *DeleteMcdpZoneRequest) SetTenantId(v string) *DeleteMcdpZoneRequest {
	s.TenantId = &v
	return s
}

func (s *DeleteMcdpZoneRequest) SetWorkspaceId(v string) *DeleteMcdpZoneRequest {
	s.WorkspaceId = &v
	return s
}

func (s *DeleteMcdpZoneRequest) Validate() error {
	return dara.Validate(s)
}
