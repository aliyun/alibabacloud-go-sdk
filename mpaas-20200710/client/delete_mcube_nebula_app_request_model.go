// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMcubeNebulaAppRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DeleteMcubeNebulaAppRequest
	GetAppId() *string
	SetH5Id(v string) *DeleteMcubeNebulaAppRequest
	GetH5Id() *string
	SetTenantId(v string) *DeleteMcubeNebulaAppRequest
	GetTenantId() *string
	SetWorkspaceId(v string) *DeleteMcubeNebulaAppRequest
	GetWorkspaceId() *string
}

type DeleteMcubeNebulaAppRequest struct {
	AppId       *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	H5Id        *string `json:"H5Id,omitempty" xml:"H5Id,omitempty"`
	TenantId    *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s DeleteMcubeNebulaAppRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteMcubeNebulaAppRequest) GoString() string {
	return s.String()
}

func (s *DeleteMcubeNebulaAppRequest) GetAppId() *string {
	return s.AppId
}

func (s *DeleteMcubeNebulaAppRequest) GetH5Id() *string {
	return s.H5Id
}

func (s *DeleteMcubeNebulaAppRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *DeleteMcubeNebulaAppRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *DeleteMcubeNebulaAppRequest) SetAppId(v string) *DeleteMcubeNebulaAppRequest {
	s.AppId = &v
	return s
}

func (s *DeleteMcubeNebulaAppRequest) SetH5Id(v string) *DeleteMcubeNebulaAppRequest {
	s.H5Id = &v
	return s
}

func (s *DeleteMcubeNebulaAppRequest) SetTenantId(v string) *DeleteMcubeNebulaAppRequest {
	s.TenantId = &v
	return s
}

func (s *DeleteMcubeNebulaAppRequest) SetWorkspaceId(v string) *DeleteMcubeNebulaAppRequest {
	s.WorkspaceId = &v
	return s
}

func (s *DeleteMcubeNebulaAppRequest) Validate() error {
	return dara.Validate(s)
}
