// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteGatewayApiGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DeleteGatewayApiGroupRequest
	GetAppId() *string
	SetGroupId(v int64) *DeleteGatewayApiGroupRequest
	GetGroupId() *int64
	SetTenantId(v string) *DeleteGatewayApiGroupRequest
	GetTenantId() *string
	SetWorkspaceId(v string) *DeleteGatewayApiGroupRequest
	GetWorkspaceId() *string
}

type DeleteGatewayApiGroupRequest struct {
	// This parameter is required.
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	GroupId *int64 `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// This parameter is required.
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// This parameter is required.
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s DeleteGatewayApiGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteGatewayApiGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteGatewayApiGroupRequest) GetAppId() *string {
	return s.AppId
}

func (s *DeleteGatewayApiGroupRequest) GetGroupId() *int64 {
	return s.GroupId
}

func (s *DeleteGatewayApiGroupRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *DeleteGatewayApiGroupRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *DeleteGatewayApiGroupRequest) SetAppId(v string) *DeleteGatewayApiGroupRequest {
	s.AppId = &v
	return s
}

func (s *DeleteGatewayApiGroupRequest) SetGroupId(v int64) *DeleteGatewayApiGroupRequest {
	s.GroupId = &v
	return s
}

func (s *DeleteGatewayApiGroupRequest) SetTenantId(v string) *DeleteGatewayApiGroupRequest {
	s.TenantId = &v
	return s
}

func (s *DeleteGatewayApiGroupRequest) SetWorkspaceId(v string) *DeleteGatewayApiGroupRequest {
	s.WorkspaceId = &v
	return s
}

func (s *DeleteGatewayApiGroupRequest) Validate() error {
	return dara.Validate(s)
}
