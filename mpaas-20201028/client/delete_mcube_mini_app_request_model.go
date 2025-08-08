// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMcubeMiniAppRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DeleteMcubeMiniAppRequest
	GetAppId() *string
	SetH5Id(v string) *DeleteMcubeMiniAppRequest
	GetH5Id() *string
	SetTenantId(v string) *DeleteMcubeMiniAppRequest
	GetTenantId() *string
	SetWorkspaceId(v string) *DeleteMcubeMiniAppRequest
	GetWorkspaceId() *string
}

type DeleteMcubeMiniAppRequest struct {
	// This parameter is required.
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	H5Id *string `json:"H5Id,omitempty" xml:"H5Id,omitempty"`
	// This parameter is required.
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// This parameter is required.
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s DeleteMcubeMiniAppRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteMcubeMiniAppRequest) GoString() string {
	return s.String()
}

func (s *DeleteMcubeMiniAppRequest) GetAppId() *string {
	return s.AppId
}

func (s *DeleteMcubeMiniAppRequest) GetH5Id() *string {
	return s.H5Id
}

func (s *DeleteMcubeMiniAppRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *DeleteMcubeMiniAppRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *DeleteMcubeMiniAppRequest) SetAppId(v string) *DeleteMcubeMiniAppRequest {
	s.AppId = &v
	return s
}

func (s *DeleteMcubeMiniAppRequest) SetH5Id(v string) *DeleteMcubeMiniAppRequest {
	s.H5Id = &v
	return s
}

func (s *DeleteMcubeMiniAppRequest) SetTenantId(v string) *DeleteMcubeMiniAppRequest {
	s.TenantId = &v
	return s
}

func (s *DeleteMcubeMiniAppRequest) SetWorkspaceId(v string) *DeleteMcubeMiniAppRequest {
	s.WorkspaceId = &v
	return s
}

func (s *DeleteMcubeMiniAppRequest) Validate() error {
	return dara.Validate(s)
}
