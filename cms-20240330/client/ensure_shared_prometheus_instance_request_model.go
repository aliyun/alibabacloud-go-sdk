// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnsureSharedPrometheusInstanceRequest interface {
  dara.Model
  String() string
  GoString() string
  SetAliyunLang(v string) *EnsureSharedPrometheusInstanceRequest
  GetAliyunLang() *string 
  SetShareScope(v string) *EnsureSharedPrometheusInstanceRequest
  GetShareScope() *string 
  SetWorkspace(v string) *EnsureSharedPrometheusInstanceRequest
  GetWorkspace() *string 
}

type EnsureSharedPrometheusInstanceRequest struct {
  // The language of the returned information. zh indicates Chinese, and en indicates English. Default value: zh.
  // 
  // example:
  // 
  // zh
  AliyunLang *string `json:"aliyunLang,omitempty" xml:"aliyunLang,omitempty"`
  // The share scope. Currently, only Workspace is supported, which indicates the shared Prometheus instance for the workspace in the requested region. Workspaces in the same region share the same instance, and no additional instances are created. Policy and Environment are not currently supported.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // Workspace
  ShareScope *string `json:"shareScope,omitempty" xml:"shareScope,omitempty"`
  // The name of the workspace. If the workspace does not exist, a 404 error is returned.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // team-workspace
  Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s EnsureSharedPrometheusInstanceRequest) String() string {
  return dara.Prettify(s)
}

func (s EnsureSharedPrometheusInstanceRequest) GoString() string {
  return s.String()
}

func (s *EnsureSharedPrometheusInstanceRequest) GetAliyunLang() *string  {
  return s.AliyunLang
}

func (s *EnsureSharedPrometheusInstanceRequest) GetShareScope() *string  {
  return s.ShareScope
}

func (s *EnsureSharedPrometheusInstanceRequest) GetWorkspace() *string  {
  return s.Workspace
}

func (s *EnsureSharedPrometheusInstanceRequest) SetAliyunLang(v string) *EnsureSharedPrometheusInstanceRequest {
  s.AliyunLang = &v
  return s
}

func (s *EnsureSharedPrometheusInstanceRequest) SetShareScope(v string) *EnsureSharedPrometheusInstanceRequest {
  s.ShareScope = &v
  return s
}

func (s *EnsureSharedPrometheusInstanceRequest) SetWorkspace(v string) *EnsureSharedPrometheusInstanceRequest {
  s.Workspace = &v
  return s
}

func (s *EnsureSharedPrometheusInstanceRequest) Validate() error {
  return dara.Validate(s)
}

