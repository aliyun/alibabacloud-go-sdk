// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMcdpAimRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DeleteMcdpAimRequest
	GetAppId() *string
	SetMpaasMappcenterMcdpAimDeleteJsonStr(v string) *DeleteMcdpAimRequest
	GetMpaasMappcenterMcdpAimDeleteJsonStr() *string
	SetTenantId(v string) *DeleteMcdpAimRequest
	GetTenantId() *string
	SetWorkspaceId(v string) *DeleteMcdpAimRequest
	GetWorkspaceId() *string
}

type DeleteMcdpAimRequest struct {
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	MpaasMappcenterMcdpAimDeleteJsonStr *string `json:"MpaasMappcenterMcdpAimDeleteJsonStr,omitempty" xml:"MpaasMappcenterMcdpAimDeleteJsonStr,omitempty"`
	TenantId                            *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	WorkspaceId                         *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s DeleteMcdpAimRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteMcdpAimRequest) GoString() string {
	return s.String()
}

func (s *DeleteMcdpAimRequest) GetAppId() *string {
	return s.AppId
}

func (s *DeleteMcdpAimRequest) GetMpaasMappcenterMcdpAimDeleteJsonStr() *string {
	return s.MpaasMappcenterMcdpAimDeleteJsonStr
}

func (s *DeleteMcdpAimRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *DeleteMcdpAimRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *DeleteMcdpAimRequest) SetAppId(v string) *DeleteMcdpAimRequest {
	s.AppId = &v
	return s
}

func (s *DeleteMcdpAimRequest) SetMpaasMappcenterMcdpAimDeleteJsonStr(v string) *DeleteMcdpAimRequest {
	s.MpaasMappcenterMcdpAimDeleteJsonStr = &v
	return s
}

func (s *DeleteMcdpAimRequest) SetTenantId(v string) *DeleteMcdpAimRequest {
	s.TenantId = &v
	return s
}

func (s *DeleteMcdpAimRequest) SetWorkspaceId(v string) *DeleteMcdpAimRequest {
	s.WorkspaceId = &v
	return s
}

func (s *DeleteMcdpAimRequest) Validate() error {
	return dara.Validate(s)
}
