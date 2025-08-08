// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMcubeWhitelistRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DeleteMcubeWhitelistRequest
	GetAppId() *string
	SetId(v int64) *DeleteMcubeWhitelistRequest
	GetId() *int64
	SetTenantId(v string) *DeleteMcubeWhitelistRequest
	GetTenantId() *string
	SetWorkspaceId(v string) *DeleteMcubeWhitelistRequest
	GetWorkspaceId() *string
}

type DeleteMcubeWhitelistRequest struct {
	// This parameter is required.
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// This parameter is required.
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s DeleteMcubeWhitelistRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteMcubeWhitelistRequest) GoString() string {
	return s.String()
}

func (s *DeleteMcubeWhitelistRequest) GetAppId() *string {
	return s.AppId
}

func (s *DeleteMcubeWhitelistRequest) GetId() *int64 {
	return s.Id
}

func (s *DeleteMcubeWhitelistRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *DeleteMcubeWhitelistRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *DeleteMcubeWhitelistRequest) SetAppId(v string) *DeleteMcubeWhitelistRequest {
	s.AppId = &v
	return s
}

func (s *DeleteMcubeWhitelistRequest) SetId(v int64) *DeleteMcubeWhitelistRequest {
	s.Id = &v
	return s
}

func (s *DeleteMcubeWhitelistRequest) SetTenantId(v string) *DeleteMcubeWhitelistRequest {
	s.TenantId = &v
	return s
}

func (s *DeleteMcubeWhitelistRequest) SetWorkspaceId(v string) *DeleteMcubeWhitelistRequest {
	s.WorkspaceId = &v
	return s
}

func (s *DeleteMcubeWhitelistRequest) Validate() error {
	return dara.Validate(s)
}
