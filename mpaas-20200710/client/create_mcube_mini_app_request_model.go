// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMcubeMiniAppRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *CreateMcubeMiniAppRequest
	GetAppId() *string
	SetH5Id(v string) *CreateMcubeMiniAppRequest
	GetH5Id() *string
	SetH5Name(v string) *CreateMcubeMiniAppRequest
	GetH5Name() *string
	SetTenantId(v string) *CreateMcubeMiniAppRequest
	GetTenantId() *string
	SetWorkspaceId(v string) *CreateMcubeMiniAppRequest
	GetWorkspaceId() *string
}

type CreateMcubeMiniAppRequest struct {
	// This parameter is required.
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	H5Id *string `json:"H5Id,omitempty" xml:"H5Id,omitempty"`
	// This parameter is required.
	H5Name *string `json:"H5Name,omitempty" xml:"H5Name,omitempty"`
	// This parameter is required.
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// This parameter is required.
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s CreateMcubeMiniAppRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateMcubeMiniAppRequest) GoString() string {
	return s.String()
}

func (s *CreateMcubeMiniAppRequest) GetAppId() *string {
	return s.AppId
}

func (s *CreateMcubeMiniAppRequest) GetH5Id() *string {
	return s.H5Id
}

func (s *CreateMcubeMiniAppRequest) GetH5Name() *string {
	return s.H5Name
}

func (s *CreateMcubeMiniAppRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *CreateMcubeMiniAppRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *CreateMcubeMiniAppRequest) SetAppId(v string) *CreateMcubeMiniAppRequest {
	s.AppId = &v
	return s
}

func (s *CreateMcubeMiniAppRequest) SetH5Id(v string) *CreateMcubeMiniAppRequest {
	s.H5Id = &v
	return s
}

func (s *CreateMcubeMiniAppRequest) SetH5Name(v string) *CreateMcubeMiniAppRequest {
	s.H5Name = &v
	return s
}

func (s *CreateMcubeMiniAppRequest) SetTenantId(v string) *CreateMcubeMiniAppRequest {
	s.TenantId = &v
	return s
}

func (s *CreateMcubeMiniAppRequest) SetWorkspaceId(v string) *CreateMcubeMiniAppRequest {
	s.WorkspaceId = &v
	return s
}

func (s *CreateMcubeMiniAppRequest) Validate() error {
	return dara.Validate(s)
}
