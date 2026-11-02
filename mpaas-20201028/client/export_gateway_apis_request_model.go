// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportGatewayApisRequest interface {
  dara.Model
  String() string
  GoString() string
  SetApiIds(v []*int64) *ExportGatewayApisRequest
  GetApiIds() []*int64 
  SetAppId(v string) *ExportGatewayApisRequest
  GetAppId() *string 
  SetTenantId(v string) *ExportGatewayApisRequest
  GetTenantId() *string 
  SetWorkspaceId(v string) *ExportGatewayApisRequest
  GetWorkspaceId() *string 
}

type ExportGatewayApisRequest struct {
  // This parameter is required.
  ApiIds []*int64 `json:"ApiIds,omitempty" xml:"ApiIds,omitempty" type:"Repeated"`
  // This parameter is required.
  AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
  // This parameter is required.
  TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
  // This parameter is required.
  WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ExportGatewayApisRequest) String() string {
  return dara.Prettify(s)
}

func (s ExportGatewayApisRequest) GoString() string {
  return s.String()
}

func (s *ExportGatewayApisRequest) GetApiIds() []*int64  {
  return s.ApiIds
}

func (s *ExportGatewayApisRequest) GetAppId() *string  {
  return s.AppId
}

func (s *ExportGatewayApisRequest) GetTenantId() *string  {
  return s.TenantId
}

func (s *ExportGatewayApisRequest) GetWorkspaceId() *string  {
  return s.WorkspaceId
}

func (s *ExportGatewayApisRequest) SetApiIds(v []*int64) *ExportGatewayApisRequest {
  s.ApiIds = v
  return s
}

func (s *ExportGatewayApisRequest) SetAppId(v string) *ExportGatewayApisRequest {
  s.AppId = &v
  return s
}

func (s *ExportGatewayApisRequest) SetTenantId(v string) *ExportGatewayApisRequest {
  s.TenantId = &v
  return s
}

func (s *ExportGatewayApisRequest) SetWorkspaceId(v string) *ExportGatewayApisRequest {
  s.WorkspaceId = &v
  return s
}

func (s *ExportGatewayApisRequest) Validate() error {
  return dara.Validate(s)
}

