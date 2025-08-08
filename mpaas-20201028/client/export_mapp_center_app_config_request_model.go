// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportMappCenterAppConfigRequest interface {
  dara.Model
  String() string
  GoString() string
  SetApkFileUrl(v string) *ExportMappCenterAppConfigRequest
  GetApkFileUrl() *string 
  SetAppId(v string) *ExportMappCenterAppConfigRequest
  GetAppId() *string 
  SetCertRsaBase64(v string) *ExportMappCenterAppConfigRequest
  GetCertRsaBase64() *string 
  SetIdentifier(v string) *ExportMappCenterAppConfigRequest
  GetIdentifier() *string 
  SetOnexFlag(v bool) *ExportMappCenterAppConfigRequest
  GetOnexFlag() *bool 
  SetSystemType(v string) *ExportMappCenterAppConfigRequest
  GetSystemType() *string 
  SetWorkspaceId(v string) *ExportMappCenterAppConfigRequest
  GetWorkspaceId() *string 
}

type ExportMappCenterAppConfigRequest struct {
  ApkFileUrl *string `json:"ApkFileUrl,omitempty" xml:"ApkFileUrl,omitempty"`
  AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
  CertRsaBase64 *string `json:"CertRsaBase64,omitempty" xml:"CertRsaBase64,omitempty"`
  // This parameter is required.
  Identifier *string `json:"Identifier,omitempty" xml:"Identifier,omitempty"`
  // This parameter is required.
  OnexFlag *bool `json:"OnexFlag,omitempty" xml:"OnexFlag,omitempty"`
  // This parameter is required.
  SystemType *string `json:"SystemType,omitempty" xml:"SystemType,omitempty"`
  WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ExportMappCenterAppConfigRequest) String() string {
  return dara.Prettify(s)
}

func (s ExportMappCenterAppConfigRequest) GoString() string {
  return s.String()
}

func (s *ExportMappCenterAppConfigRequest) GetApkFileUrl() *string  {
  return s.ApkFileUrl
}

func (s *ExportMappCenterAppConfigRequest) GetAppId() *string  {
  return s.AppId
}

func (s *ExportMappCenterAppConfigRequest) GetCertRsaBase64() *string  {
  return s.CertRsaBase64
}

func (s *ExportMappCenterAppConfigRequest) GetIdentifier() *string  {
  return s.Identifier
}

func (s *ExportMappCenterAppConfigRequest) GetOnexFlag() *bool  {
  return s.OnexFlag
}

func (s *ExportMappCenterAppConfigRequest) GetSystemType() *string  {
  return s.SystemType
}

func (s *ExportMappCenterAppConfigRequest) GetWorkspaceId() *string  {
  return s.WorkspaceId
}

func (s *ExportMappCenterAppConfigRequest) SetApkFileUrl(v string) *ExportMappCenterAppConfigRequest {
  s.ApkFileUrl = &v
  return s
}

func (s *ExportMappCenterAppConfigRequest) SetAppId(v string) *ExportMappCenterAppConfigRequest {
  s.AppId = &v
  return s
}

func (s *ExportMappCenterAppConfigRequest) SetCertRsaBase64(v string) *ExportMappCenterAppConfigRequest {
  s.CertRsaBase64 = &v
  return s
}

func (s *ExportMappCenterAppConfigRequest) SetIdentifier(v string) *ExportMappCenterAppConfigRequest {
  s.Identifier = &v
  return s
}

func (s *ExportMappCenterAppConfigRequest) SetOnexFlag(v bool) *ExportMappCenterAppConfigRequest {
  s.OnexFlag = &v
  return s
}

func (s *ExportMappCenterAppConfigRequest) SetSystemType(v string) *ExportMappCenterAppConfigRequest {
  s.SystemType = &v
  return s
}

func (s *ExportMappCenterAppConfigRequest) SetWorkspaceId(v string) *ExportMappCenterAppConfigRequest {
  s.WorkspaceId = &v
  return s
}

func (s *ExportMappCenterAppConfigRequest) Validate() error {
  return dara.Validate(s)
}

