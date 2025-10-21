// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEngineVersionSupportedFeatures interface {
  dara.Model
  String() string
  GoString() string
  SetSupportNativeSavepoint(v bool) *EngineVersionSupportedFeatures
  GetSupportNativeSavepoint() *bool 
  SetUseForSqlDeployments(v bool) *EngineVersionSupportedFeatures
  GetUseForSqlDeployments() *bool 
}

type EngineVersionSupportedFeatures struct {
  // example:
  // 
  // true
  SupportNativeSavepoint *bool `json:"supportNativeSavepoint,omitempty" xml:"supportNativeSavepoint,omitempty"`
  // example:
  // 
  // true
  UseForSqlDeployments *bool `json:"useForSqlDeployments,omitempty" xml:"useForSqlDeployments,omitempty"`
}

func (s EngineVersionSupportedFeatures) String() string {
  return dara.Prettify(s)
}

func (s EngineVersionSupportedFeatures) GoString() string {
  return s.String()
}

func (s *EngineVersionSupportedFeatures) GetSupportNativeSavepoint() *bool  {
  return s.SupportNativeSavepoint
}

func (s *EngineVersionSupportedFeatures) GetUseForSqlDeployments() *bool  {
  return s.UseForSqlDeployments
}

func (s *EngineVersionSupportedFeatures) SetSupportNativeSavepoint(v bool) *EngineVersionSupportedFeatures {
  s.SupportNativeSavepoint = &v
  return s
}

func (s *EngineVersionSupportedFeatures) SetUseForSqlDeployments(v bool) *EngineVersionSupportedFeatures {
  s.UseForSqlDeployments = &v
  return s
}

func (s *EngineVersionSupportedFeatures) Validate() error {
  return dara.Validate(s)
}

