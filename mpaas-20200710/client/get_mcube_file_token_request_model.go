// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMcubeFileTokenRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *GetMcubeFileTokenRequest
	GetAppId() *string
	SetOnexFlag(v bool) *GetMcubeFileTokenRequest
	GetOnexFlag() *bool
	SetTenantId(v string) *GetMcubeFileTokenRequest
	GetTenantId() *string
	SetWorkspaceId(v string) *GetMcubeFileTokenRequest
	GetWorkspaceId() *string
}

type GetMcubeFileTokenRequest struct {
	// This parameter is required.
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	OnexFlag *bool `json:"OnexFlag,omitempty" xml:"OnexFlag,omitempty"`
	// This parameter is required.
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// This parameter is required.
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s GetMcubeFileTokenRequest) String() string {
	return dara.Prettify(s)
}

func (s GetMcubeFileTokenRequest) GoString() string {
	return s.String()
}

func (s *GetMcubeFileTokenRequest) GetAppId() *string {
	return s.AppId
}

func (s *GetMcubeFileTokenRequest) GetOnexFlag() *bool {
	return s.OnexFlag
}

func (s *GetMcubeFileTokenRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *GetMcubeFileTokenRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetMcubeFileTokenRequest) SetAppId(v string) *GetMcubeFileTokenRequest {
	s.AppId = &v
	return s
}

func (s *GetMcubeFileTokenRequest) SetOnexFlag(v bool) *GetMcubeFileTokenRequest {
	s.OnexFlag = &v
	return s
}

func (s *GetMcubeFileTokenRequest) SetTenantId(v string) *GetMcubeFileTokenRequest {
	s.TenantId = &v
	return s
}

func (s *GetMcubeFileTokenRequest) SetWorkspaceId(v string) *GetMcubeFileTokenRequest {
	s.WorkspaceId = &v
	return s
}

func (s *GetMcubeFileTokenRequest) Validate() error {
	return dara.Validate(s)
}
