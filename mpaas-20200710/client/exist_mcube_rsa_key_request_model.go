// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExistMcubeRsaKeyRequest interface {
  dara.Model
  String() string
  GoString() string
  SetAppId(v string) *ExistMcubeRsaKeyRequest
  GetAppId() *string 
  SetTenantId(v string) *ExistMcubeRsaKeyRequest
  GetTenantId() *string 
  SetWorkspaceId(v string) *ExistMcubeRsaKeyRequest
  GetWorkspaceId() *string 
}

type ExistMcubeRsaKeyRequest struct {
  // This parameter is required.
  AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
  // This parameter is required.
  TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
  // This parameter is required.
  WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ExistMcubeRsaKeyRequest) String() string {
  return dara.Prettify(s)
}

func (s ExistMcubeRsaKeyRequest) GoString() string {
  return s.String()
}

func (s *ExistMcubeRsaKeyRequest) GetAppId() *string  {
  return s.AppId
}

func (s *ExistMcubeRsaKeyRequest) GetTenantId() *string  {
  return s.TenantId
}

func (s *ExistMcubeRsaKeyRequest) GetWorkspaceId() *string  {
  return s.WorkspaceId
}

func (s *ExistMcubeRsaKeyRequest) SetAppId(v string) *ExistMcubeRsaKeyRequest {
  s.AppId = &v
  return s
}

func (s *ExistMcubeRsaKeyRequest) SetTenantId(v string) *ExistMcubeRsaKeyRequest {
  s.TenantId = &v
  return s
}

func (s *ExistMcubeRsaKeyRequest) SetWorkspaceId(v string) *ExistMcubeRsaKeyRequest {
  s.WorkspaceId = &v
  return s
}

func (s *ExistMcubeRsaKeyRequest) Validate() error {
  return dara.Validate(s)
}

