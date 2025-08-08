// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMcdpMaterialRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *CreateMcdpMaterialRequest
	GetAppId() *string
	SetMpaasMappcenterMcdpMaterialCreateJsonStr(v string) *CreateMcdpMaterialRequest
	GetMpaasMappcenterMcdpMaterialCreateJsonStr() *string
	SetTenantId(v string) *CreateMcdpMaterialRequest
	GetTenantId() *string
	SetWorkspaceId(v string) *CreateMcdpMaterialRequest
	GetWorkspaceId() *string
}

type CreateMcdpMaterialRequest struct {
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	MpaasMappcenterMcdpMaterialCreateJsonStr *string `json:"MpaasMappcenterMcdpMaterialCreateJsonStr,omitempty" xml:"MpaasMappcenterMcdpMaterialCreateJsonStr,omitempty"`
	TenantId                                 *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	WorkspaceId                              *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s CreateMcdpMaterialRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateMcdpMaterialRequest) GoString() string {
	return s.String()
}

func (s *CreateMcdpMaterialRequest) GetAppId() *string {
	return s.AppId
}

func (s *CreateMcdpMaterialRequest) GetMpaasMappcenterMcdpMaterialCreateJsonStr() *string {
	return s.MpaasMappcenterMcdpMaterialCreateJsonStr
}

func (s *CreateMcdpMaterialRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *CreateMcdpMaterialRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *CreateMcdpMaterialRequest) SetAppId(v string) *CreateMcdpMaterialRequest {
	s.AppId = &v
	return s
}

func (s *CreateMcdpMaterialRequest) SetMpaasMappcenterMcdpMaterialCreateJsonStr(v string) *CreateMcdpMaterialRequest {
	s.MpaasMappcenterMcdpMaterialCreateJsonStr = &v
	return s
}

func (s *CreateMcdpMaterialRequest) SetTenantId(v string) *CreateMcdpMaterialRequest {
	s.TenantId = &v
	return s
}

func (s *CreateMcdpMaterialRequest) SetWorkspaceId(v string) *CreateMcdpMaterialRequest {
	s.WorkspaceId = &v
	return s
}

func (s *CreateMcdpMaterialRequest) Validate() error {
	return dara.Validate(s)
}
