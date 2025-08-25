// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMcubeVhostRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *CreateMcubeVhostRequest
	GetAppId() *string
	SetTenantId(v string) *CreateMcubeVhostRequest
	GetTenantId() *string
	SetVhost(v string) *CreateMcubeVhostRequest
	GetVhost() *string
	SetWorkspaceId(v string) *CreateMcubeVhostRequest
	GetWorkspaceId() *string
}

type CreateMcubeVhostRequest struct {
	// This parameter is required.
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// This parameter is required.
	Vhost *string `json:"Vhost,omitempty" xml:"Vhost,omitempty"`
	// This parameter is required.
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s CreateMcubeVhostRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateMcubeVhostRequest) GoString() string {
	return s.String()
}

func (s *CreateMcubeVhostRequest) GetAppId() *string {
	return s.AppId
}

func (s *CreateMcubeVhostRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *CreateMcubeVhostRequest) GetVhost() *string {
	return s.Vhost
}

func (s *CreateMcubeVhostRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *CreateMcubeVhostRequest) SetAppId(v string) *CreateMcubeVhostRequest {
	s.AppId = &v
	return s
}

func (s *CreateMcubeVhostRequest) SetTenantId(v string) *CreateMcubeVhostRequest {
	s.TenantId = &v
	return s
}

func (s *CreateMcubeVhostRequest) SetVhost(v string) *CreateMcubeVhostRequest {
	s.Vhost = &v
	return s
}

func (s *CreateMcubeVhostRequest) SetWorkspaceId(v string) *CreateMcubeVhostRequest {
	s.WorkspaceId = &v
	return s
}

func (s *CreateMcubeVhostRequest) Validate() error {
	return dara.Validate(s)
}
