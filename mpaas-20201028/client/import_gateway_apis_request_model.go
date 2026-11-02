// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportGatewayApisRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *ImportGatewayApisRequest
	GetAppId() *string
	SetFileUrl(v string) *ImportGatewayApisRequest
	GetFileUrl() *string
	SetImportMode(v string) *ImportGatewayApisRequest
	GetImportMode() *string
	SetTenantId(v string) *ImportGatewayApisRequest
	GetTenantId() *string
	SetWorkspaceId(v string) *ImportGatewayApisRequest
	GetWorkspaceId() *string
}

type ImportGatewayApisRequest struct {
	// This parameter is required.
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	FileUrl    *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	ImportMode *string `json:"ImportMode,omitempty" xml:"ImportMode,omitempty"`
	// This parameter is required.
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// This parameter is required.
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ImportGatewayApisRequest) String() string {
	return dara.Prettify(s)
}

func (s ImportGatewayApisRequest) GoString() string {
	return s.String()
}

func (s *ImportGatewayApisRequest) GetAppId() *string {
	return s.AppId
}

func (s *ImportGatewayApisRequest) GetFileUrl() *string {
	return s.FileUrl
}

func (s *ImportGatewayApisRequest) GetImportMode() *string {
	return s.ImportMode
}

func (s *ImportGatewayApisRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ImportGatewayApisRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ImportGatewayApisRequest) SetAppId(v string) *ImportGatewayApisRequest {
	s.AppId = &v
	return s
}

func (s *ImportGatewayApisRequest) SetFileUrl(v string) *ImportGatewayApisRequest {
	s.FileUrl = &v
	return s
}

func (s *ImportGatewayApisRequest) SetImportMode(v string) *ImportGatewayApisRequest {
	s.ImportMode = &v
	return s
}

func (s *ImportGatewayApisRequest) SetTenantId(v string) *ImportGatewayApisRequest {
	s.TenantId = &v
	return s
}

func (s *ImportGatewayApisRequest) SetWorkspaceId(v string) *ImportGatewayApisRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ImportGatewayApisRequest) Validate() error {
	return dara.Validate(s)
}
