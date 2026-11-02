// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGatewayApiGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *GetGatewayApiGroupRequest
	GetAppId() *string
	SetGroupId(v int64) *GetGatewayApiGroupRequest
	GetGroupId() *int64
	SetTenantId(v string) *GetGatewayApiGroupRequest
	GetTenantId() *string
	SetWorkspaceId(v string) *GetGatewayApiGroupRequest
	GetWorkspaceId() *string
}

type GetGatewayApiGroupRequest struct {
	// This parameter is required.
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	GroupId *int64 `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// This parameter is required.
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// This parameter is required.
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s GetGatewayApiGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s GetGatewayApiGroupRequest) GoString() string {
	return s.String()
}

func (s *GetGatewayApiGroupRequest) GetAppId() *string {
	return s.AppId
}

func (s *GetGatewayApiGroupRequest) GetGroupId() *int64 {
	return s.GroupId
}

func (s *GetGatewayApiGroupRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *GetGatewayApiGroupRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetGatewayApiGroupRequest) SetAppId(v string) *GetGatewayApiGroupRequest {
	s.AppId = &v
	return s
}

func (s *GetGatewayApiGroupRequest) SetGroupId(v int64) *GetGatewayApiGroupRequest {
	s.GroupId = &v
	return s
}

func (s *GetGatewayApiGroupRequest) SetTenantId(v string) *GetGatewayApiGroupRequest {
	s.TenantId = &v
	return s
}

func (s *GetGatewayApiGroupRequest) SetWorkspaceId(v string) *GetGatewayApiGroupRequest {
	s.WorkspaceId = &v
	return s
}

func (s *GetGatewayApiGroupRequest) Validate() error {
	return dara.Validate(s)
}
