// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMcdpGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *CreateMcdpGroupRequest
	GetAppId() *string
	SetMpaasMappcenterMcdpGroupCreateJsonStr(v string) *CreateMcdpGroupRequest
	GetMpaasMappcenterMcdpGroupCreateJsonStr() *string
	SetTenantId(v string) *CreateMcdpGroupRequest
	GetTenantId() *string
	SetWorkspaceId(v string) *CreateMcdpGroupRequest
	GetWorkspaceId() *string
}

type CreateMcdpGroupRequest struct {
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	MpaasMappcenterMcdpGroupCreateJsonStr *string `json:"MpaasMappcenterMcdpGroupCreateJsonStr,omitempty" xml:"MpaasMappcenterMcdpGroupCreateJsonStr,omitempty"`
	TenantId                              *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	WorkspaceId                           *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s CreateMcdpGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateMcdpGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateMcdpGroupRequest) GetAppId() *string {
	return s.AppId
}

func (s *CreateMcdpGroupRequest) GetMpaasMappcenterMcdpGroupCreateJsonStr() *string {
	return s.MpaasMappcenterMcdpGroupCreateJsonStr
}

func (s *CreateMcdpGroupRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *CreateMcdpGroupRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *CreateMcdpGroupRequest) SetAppId(v string) *CreateMcdpGroupRequest {
	s.AppId = &v
	return s
}

func (s *CreateMcdpGroupRequest) SetMpaasMappcenterMcdpGroupCreateJsonStr(v string) *CreateMcdpGroupRequest {
	s.MpaasMappcenterMcdpGroupCreateJsonStr = &v
	return s
}

func (s *CreateMcdpGroupRequest) SetTenantId(v string) *CreateMcdpGroupRequest {
	s.TenantId = &v
	return s
}

func (s *CreateMcdpGroupRequest) SetWorkspaceId(v string) *CreateMcdpGroupRequest {
	s.WorkspaceId = &v
	return s
}

func (s *CreateMcdpGroupRequest) Validate() error {
	return dara.Validate(s)
}
