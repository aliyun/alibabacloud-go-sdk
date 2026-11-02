// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportGatewayApisShrinkRequest interface {
  dara.Model
  String() string
  GoString() string
  SetApiIdsShrink(v string) *ExportGatewayApisShrinkRequest
  GetApiIdsShrink() *string 
  SetAppId(v string) *ExportGatewayApisShrinkRequest
  GetAppId() *string 
  SetTenantId(v string) *ExportGatewayApisShrinkRequest
  GetTenantId() *string 
  SetWorkspaceId(v string) *ExportGatewayApisShrinkRequest
  GetWorkspaceId() *string 
}

type ExportGatewayApisShrinkRequest struct {
  // This parameter is required.
  ApiIdsShrink *string `json:"ApiIds,omitempty" xml:"ApiIds,omitempty"`
  // This parameter is required.
  AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
  // This parameter is required.
  TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
  // This parameter is required.
  WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ExportGatewayApisShrinkRequest) String() string {
  return dara.Prettify(s)
}

func (s ExportGatewayApisShrinkRequest) GoString() string {
  return s.String()
}

func (s *ExportGatewayApisShrinkRequest) GetApiIdsShrink() *string  {
  return s.ApiIdsShrink
}

func (s *ExportGatewayApisShrinkRequest) GetAppId() *string  {
  return s.AppId
}

func (s *ExportGatewayApisShrinkRequest) GetTenantId() *string  {
  return s.TenantId
}

func (s *ExportGatewayApisShrinkRequest) GetWorkspaceId() *string  {
  return s.WorkspaceId
}

func (s *ExportGatewayApisShrinkRequest) SetApiIdsShrink(v string) *ExportGatewayApisShrinkRequest {
  s.ApiIdsShrink = &v
  return s
}

func (s *ExportGatewayApisShrinkRequest) SetAppId(v string) *ExportGatewayApisShrinkRequest {
  s.AppId = &v
  return s
}

func (s *ExportGatewayApisShrinkRequest) SetTenantId(v string) *ExportGatewayApisShrinkRequest {
  s.TenantId = &v
  return s
}

func (s *ExportGatewayApisShrinkRequest) SetWorkspaceId(v string) *ExportGatewayApisShrinkRequest {
  s.WorkspaceId = &v
  return s
}

func (s *ExportGatewayApisShrinkRequest) Validate() error {
  return dara.Validate(s)
}

