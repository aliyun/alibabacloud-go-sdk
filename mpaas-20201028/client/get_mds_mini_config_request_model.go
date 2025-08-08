// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMdsMiniConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *GetMdsMiniConfigRequest
	GetAppId() *string
	SetH5Id(v string) *GetMdsMiniConfigRequest
	GetH5Id() *string
	SetTenantId(v string) *GetMdsMiniConfigRequest
	GetTenantId() *string
	SetWorkspaceId(v string) *GetMdsMiniConfigRequest
	GetWorkspaceId() *string
}

type GetMdsMiniConfigRequest struct {
	// This parameter is required.
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	H5Id *string `json:"H5Id,omitempty" xml:"H5Id,omitempty"`
	// This parameter is required.
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// This parameter is required.
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s GetMdsMiniConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s GetMdsMiniConfigRequest) GoString() string {
	return s.String()
}

func (s *GetMdsMiniConfigRequest) GetAppId() *string {
	return s.AppId
}

func (s *GetMdsMiniConfigRequest) GetH5Id() *string {
	return s.H5Id
}

func (s *GetMdsMiniConfigRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *GetMdsMiniConfigRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetMdsMiniConfigRequest) SetAppId(v string) *GetMdsMiniConfigRequest {
	s.AppId = &v
	return s
}

func (s *GetMdsMiniConfigRequest) SetH5Id(v string) *GetMdsMiniConfigRequest {
	s.H5Id = &v
	return s
}

func (s *GetMdsMiniConfigRequest) SetTenantId(v string) *GetMdsMiniConfigRequest {
	s.TenantId = &v
	return s
}

func (s *GetMdsMiniConfigRequest) SetWorkspaceId(v string) *GetMdsMiniConfigRequest {
	s.WorkspaceId = &v
	return s
}

func (s *GetMdsMiniConfigRequest) Validate() error {
	return dara.Validate(s)
}
