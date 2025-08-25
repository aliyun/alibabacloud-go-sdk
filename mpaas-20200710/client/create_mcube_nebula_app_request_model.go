// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMcubeNebulaAppRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *CreateMcubeNebulaAppRequest
	GetAppId() *string
	SetH5Id(v string) *CreateMcubeNebulaAppRequest
	GetH5Id() *string
	SetH5Name(v string) *CreateMcubeNebulaAppRequest
	GetH5Name() *string
	SetTenantId(v string) *CreateMcubeNebulaAppRequest
	GetTenantId() *string
	SetWorkspaceId(v string) *CreateMcubeNebulaAppRequest
	GetWorkspaceId() *string
}

type CreateMcubeNebulaAppRequest struct {
	AppId       *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	H5Id        *string `json:"H5Id,omitempty" xml:"H5Id,omitempty"`
	H5Name      *string `json:"H5Name,omitempty" xml:"H5Name,omitempty"`
	TenantId    *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s CreateMcubeNebulaAppRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateMcubeNebulaAppRequest) GoString() string {
	return s.String()
}

func (s *CreateMcubeNebulaAppRequest) GetAppId() *string {
	return s.AppId
}

func (s *CreateMcubeNebulaAppRequest) GetH5Id() *string {
	return s.H5Id
}

func (s *CreateMcubeNebulaAppRequest) GetH5Name() *string {
	return s.H5Name
}

func (s *CreateMcubeNebulaAppRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *CreateMcubeNebulaAppRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *CreateMcubeNebulaAppRequest) SetAppId(v string) *CreateMcubeNebulaAppRequest {
	s.AppId = &v
	return s
}

func (s *CreateMcubeNebulaAppRequest) SetH5Id(v string) *CreateMcubeNebulaAppRequest {
	s.H5Id = &v
	return s
}

func (s *CreateMcubeNebulaAppRequest) SetH5Name(v string) *CreateMcubeNebulaAppRequest {
	s.H5Name = &v
	return s
}

func (s *CreateMcubeNebulaAppRequest) SetTenantId(v string) *CreateMcubeNebulaAppRequest {
	s.TenantId = &v
	return s
}

func (s *CreateMcubeNebulaAppRequest) SetWorkspaceId(v string) *CreateMcubeNebulaAppRequest {
	s.WorkspaceId = &v
	return s
}

func (s *CreateMcubeNebulaAppRequest) Validate() error {
	return dara.Validate(s)
}
