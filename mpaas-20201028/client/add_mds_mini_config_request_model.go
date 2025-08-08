// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddMdsMiniConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *AddMdsMiniConfigRequest
	GetAppId() *string
	SetMpaasMappcenterMiniConfigAddJsonStr(v string) *AddMdsMiniConfigRequest
	GetMpaasMappcenterMiniConfigAddJsonStr() *string
	SetTenantId(v string) *AddMdsMiniConfigRequest
	GetTenantId() *string
	SetWorkspaceId(v string) *AddMdsMiniConfigRequest
	GetWorkspaceId() *string
}

type AddMdsMiniConfigRequest struct {
	// This parameter is required.
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	MpaasMappcenterMiniConfigAddJsonStr *string `json:"MpaasMappcenterMiniConfigAddJsonStr,omitempty" xml:"MpaasMappcenterMiniConfigAddJsonStr,omitempty"`
	// This parameter is required.
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// This parameter is required.
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s AddMdsMiniConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s AddMdsMiniConfigRequest) GoString() string {
	return s.String()
}

func (s *AddMdsMiniConfigRequest) GetAppId() *string {
	return s.AppId
}

func (s *AddMdsMiniConfigRequest) GetMpaasMappcenterMiniConfigAddJsonStr() *string {
	return s.MpaasMappcenterMiniConfigAddJsonStr
}

func (s *AddMdsMiniConfigRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *AddMdsMiniConfigRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *AddMdsMiniConfigRequest) SetAppId(v string) *AddMdsMiniConfigRequest {
	s.AppId = &v
	return s
}

func (s *AddMdsMiniConfigRequest) SetMpaasMappcenterMiniConfigAddJsonStr(v string) *AddMdsMiniConfigRequest {
	s.MpaasMappcenterMiniConfigAddJsonStr = &v
	return s
}

func (s *AddMdsMiniConfigRequest) SetTenantId(v string) *AddMdsMiniConfigRequest {
	s.TenantId = &v
	return s
}

func (s *AddMdsMiniConfigRequest) SetWorkspaceId(v string) *AddMdsMiniConfigRequest {
	s.WorkspaceId = &v
	return s
}

func (s *AddMdsMiniConfigRequest) Validate() error {
	return dara.Validate(s)
}
