// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecWorkflowConnectivityRequest interface {
  dara.Model
  String() string
  GoString() string
  SetDsConfig(v string) *ExecWorkflowConnectivityRequest
  GetDsConfig() *string 
  SetDsName(v string) *ExecWorkflowConnectivityRequest
  GetDsName() *string 
  SetDsType(v string) *ExecWorkflowConnectivityRequest
  GetDsType() *string 
  SetDsVersion(v string) *ExecWorkflowConnectivityRequest
  GetDsVersion() *string 
  SetId(v int64) *ExecWorkflowConnectivityRequest
  GetId() *int64 
  SetIsModified(v bool) *ExecWorkflowConnectivityRequest
  GetIsModified() *bool 
}

type ExecWorkflowConnectivityRequest struct {
  // The datasource config. The value is a JSON character string whose structure is defined by each dsType. Parse the JSON string before use. Sensitive fields such as tokens are masked in the response.
  // 
  // example:
  // 
  // {"endpoint":"...","token":"******"}
  DsConfig *string `json:"dsConfig,omitempty" xml:"dsConfig,omitempty"`
  // The data source name. Exact match and fuzzy match are supported.
  // 
  // example:
  // 
  // test_ds318_hangzhou_0428
  DsName *string `json:"dsName,omitempty" xml:"dsName,omitempty"`
  // The data source type, such as Hive or MaxCompute.
  // 
  // example:
  // 
  // Hive
  DsType *string `json:"dsType,omitempty" xml:"dsType,omitempty"`
  // The data source version number.
  // 
  // example:
  // 
  // 3.2.0
  DsVersion *string `json:"dsVersion,omitempty" xml:"dsVersion,omitempty"`
  // The primary key ID that uniquely identifies a record.
  // 
  // example:
  // 
  // 10001
  Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
  // Specifies whether the configuration has been modified.
  IsModified *bool `json:"isModified,omitempty" xml:"isModified,omitempty"`
}

func (s ExecWorkflowConnectivityRequest) String() string {
  return dara.Prettify(s)
}

func (s ExecWorkflowConnectivityRequest) GoString() string {
  return s.String()
}

func (s *ExecWorkflowConnectivityRequest) GetDsConfig() *string  {
  return s.DsConfig
}

func (s *ExecWorkflowConnectivityRequest) GetDsName() *string  {
  return s.DsName
}

func (s *ExecWorkflowConnectivityRequest) GetDsType() *string  {
  return s.DsType
}

func (s *ExecWorkflowConnectivityRequest) GetDsVersion() *string  {
  return s.DsVersion
}

func (s *ExecWorkflowConnectivityRequest) GetId() *int64  {
  return s.Id
}

func (s *ExecWorkflowConnectivityRequest) GetIsModified() *bool  {
  return s.IsModified
}

func (s *ExecWorkflowConnectivityRequest) SetDsConfig(v string) *ExecWorkflowConnectivityRequest {
  s.DsConfig = &v
  return s
}

func (s *ExecWorkflowConnectivityRequest) SetDsName(v string) *ExecWorkflowConnectivityRequest {
  s.DsName = &v
  return s
}

func (s *ExecWorkflowConnectivityRequest) SetDsType(v string) *ExecWorkflowConnectivityRequest {
  s.DsType = &v
  return s
}

func (s *ExecWorkflowConnectivityRequest) SetDsVersion(v string) *ExecWorkflowConnectivityRequest {
  s.DsVersion = &v
  return s
}

func (s *ExecWorkflowConnectivityRequest) SetId(v int64) *ExecWorkflowConnectivityRequest {
  s.Id = &v
  return s
}

func (s *ExecWorkflowConnectivityRequest) SetIsModified(v bool) *ExecWorkflowConnectivityRequest {
  s.IsModified = &v
  return s
}

func (s *ExecWorkflowConnectivityRequest) Validate() error {
  return dara.Validate(s)
}

