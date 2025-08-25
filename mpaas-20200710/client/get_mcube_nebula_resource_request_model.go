// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMcubeNebulaResourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *GetMcubeNebulaResourceRequest
	GetAppId() *string
	SetId(v string) *GetMcubeNebulaResourceRequest
	GetId() *string
	SetTenantId(v string) *GetMcubeNebulaResourceRequest
	GetTenantId() *string
	SetWorkspaceId(v string) *GetMcubeNebulaResourceRequest
	GetWorkspaceId() *string
}

type GetMcubeNebulaResourceRequest struct {
	AppId       *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	Id          *string `json:"Id,omitempty" xml:"Id,omitempty"`
	TenantId    *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s GetMcubeNebulaResourceRequest) String() string {
	return dara.Prettify(s)
}

func (s GetMcubeNebulaResourceRequest) GoString() string {
	return s.String()
}

func (s *GetMcubeNebulaResourceRequest) GetAppId() *string {
	return s.AppId
}

func (s *GetMcubeNebulaResourceRequest) GetId() *string {
	return s.Id
}

func (s *GetMcubeNebulaResourceRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *GetMcubeNebulaResourceRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetMcubeNebulaResourceRequest) SetAppId(v string) *GetMcubeNebulaResourceRequest {
	s.AppId = &v
	return s
}

func (s *GetMcubeNebulaResourceRequest) SetId(v string) *GetMcubeNebulaResourceRequest {
	s.Id = &v
	return s
}

func (s *GetMcubeNebulaResourceRequest) SetTenantId(v string) *GetMcubeNebulaResourceRequest {
	s.TenantId = &v
	return s
}

func (s *GetMcubeNebulaResourceRequest) SetWorkspaceId(v string) *GetMcubeNebulaResourceRequest {
	s.WorkspaceId = &v
	return s
}

func (s *GetMcubeNebulaResourceRequest) Validate() error {
	return dara.Validate(s)
}
