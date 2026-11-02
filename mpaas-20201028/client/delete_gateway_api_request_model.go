// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteGatewayApiRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiId(v int64) *DeleteGatewayApiRequest
	GetApiId() *int64
	SetAppId(v string) *DeleteGatewayApiRequest
	GetAppId() *string
	SetTenantId(v string) *DeleteGatewayApiRequest
	GetTenantId() *string
	SetWorkspaceId(v string) *DeleteGatewayApiRequest
	GetWorkspaceId() *string
}

type DeleteGatewayApiRequest struct {
	// This parameter is required.
	ApiId *int64 `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	// This parameter is required.
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// This parameter is required.
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s DeleteGatewayApiRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteGatewayApiRequest) GoString() string {
	return s.String()
}

func (s *DeleteGatewayApiRequest) GetApiId() *int64 {
	return s.ApiId
}

func (s *DeleteGatewayApiRequest) GetAppId() *string {
	return s.AppId
}

func (s *DeleteGatewayApiRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *DeleteGatewayApiRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *DeleteGatewayApiRequest) SetApiId(v int64) *DeleteGatewayApiRequest {
	s.ApiId = &v
	return s
}

func (s *DeleteGatewayApiRequest) SetAppId(v string) *DeleteGatewayApiRequest {
	s.AppId = &v
	return s
}

func (s *DeleteGatewayApiRequest) SetTenantId(v string) *DeleteGatewayApiRequest {
	s.TenantId = &v
	return s
}

func (s *DeleteGatewayApiRequest) SetWorkspaceId(v string) *DeleteGatewayApiRequest {
	s.WorkspaceId = &v
	return s
}

func (s *DeleteGatewayApiRequest) Validate() error {
	return dara.Validate(s)
}
