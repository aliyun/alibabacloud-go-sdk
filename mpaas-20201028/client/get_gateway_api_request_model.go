// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGatewayApiRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiId(v int64) *GetGatewayApiRequest
	GetApiId() *int64
	SetAppId(v string) *GetGatewayApiRequest
	GetAppId() *string
	SetTenantId(v string) *GetGatewayApiRequest
	GetTenantId() *string
	SetWorkspaceId(v string) *GetGatewayApiRequest
	GetWorkspaceId() *string
}

type GetGatewayApiRequest struct {
	// This parameter is required.
	ApiId *int64 `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	// This parameter is required.
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// This parameter is required.
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s GetGatewayApiRequest) String() string {
	return dara.Prettify(s)
}

func (s GetGatewayApiRequest) GoString() string {
	return s.String()
}

func (s *GetGatewayApiRequest) GetApiId() *int64 {
	return s.ApiId
}

func (s *GetGatewayApiRequest) GetAppId() *string {
	return s.AppId
}

func (s *GetGatewayApiRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *GetGatewayApiRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetGatewayApiRequest) SetApiId(v int64) *GetGatewayApiRequest {
	s.ApiId = &v
	return s
}

func (s *GetGatewayApiRequest) SetAppId(v string) *GetGatewayApiRequest {
	s.AppId = &v
	return s
}

func (s *GetGatewayApiRequest) SetTenantId(v string) *GetGatewayApiRequest {
	s.TenantId = &v
	return s
}

func (s *GetGatewayApiRequest) SetWorkspaceId(v string) *GetGatewayApiRequest {
	s.WorkspaceId = &v
	return s
}

func (s *GetGatewayApiRequest) Validate() error {
	return dara.Validate(s)
}
