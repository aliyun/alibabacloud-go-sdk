// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnvironmentInfo interface {
  dara.Model
  String() string
  GoString() string
  SetAlias(v string) *EnvironmentInfo
  GetAlias() *string 
  SetCreateTimestamp(v int64) *EnvironmentInfo
  GetCreateTimestamp() *int64 
  SetDefault(v bool) *EnvironmentInfo
  GetDefault() *bool 
  SetDescription(v string) *EnvironmentInfo
  GetDescription() *string 
  SetEnvironmentId(v string) *EnvironmentInfo
  GetEnvironmentId() *string 
  SetGatewayInfo(v *GatewayInfo) *EnvironmentInfo
  GetGatewayInfo() *GatewayInfo 
  SetName(v string) *EnvironmentInfo
  GetName() *string 
  SetResourceGroupId(v string) *EnvironmentInfo
  GetResourceGroupId() *string 
  SetSubDomainInfos(v []*SubDomainInfo) *EnvironmentInfo
  GetSubDomainInfos() []*SubDomainInfo 
  SetUpdateTimestamp(v int64) *EnvironmentInfo
  GetUpdateTimestamp() *int64 
}

type EnvironmentInfo struct {
  Alias *string `json:"alias,omitempty" xml:"alias,omitempty"`
  CreateTimestamp *int64 `json:"createTimestamp,omitempty" xml:"createTimestamp,omitempty"`
  Default *bool `json:"default,omitempty" xml:"default,omitempty"`
  Description *string `json:"description,omitempty" xml:"description,omitempty"`
  EnvironmentId *string `json:"environmentId,omitempty" xml:"environmentId,omitempty"`
  GatewayInfo *GatewayInfo `json:"gatewayInfo,omitempty" xml:"gatewayInfo,omitempty"`
  Name *string `json:"name,omitempty" xml:"name,omitempty"`
  // example:
  // 
  // rg-xxxx
  ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
  SubDomainInfos []*SubDomainInfo `json:"subDomainInfos,omitempty" xml:"subDomainInfos,omitempty" type:"Repeated"`
  UpdateTimestamp *int64 `json:"updateTimestamp,omitempty" xml:"updateTimestamp,omitempty"`
}

func (s EnvironmentInfo) String() string {
  return dara.Prettify(s)
}

func (s EnvironmentInfo) GoString() string {
  return s.String()
}

func (s *EnvironmentInfo) GetAlias() *string  {
  return s.Alias
}

func (s *EnvironmentInfo) GetCreateTimestamp() *int64  {
  return s.CreateTimestamp
}

func (s *EnvironmentInfo) GetDefault() *bool  {
  return s.Default
}

func (s *EnvironmentInfo) GetDescription() *string  {
  return s.Description
}

func (s *EnvironmentInfo) GetEnvironmentId() *string  {
  return s.EnvironmentId
}

func (s *EnvironmentInfo) GetGatewayInfo() *GatewayInfo  {
  return s.GatewayInfo
}

func (s *EnvironmentInfo) GetName() *string  {
  return s.Name
}

func (s *EnvironmentInfo) GetResourceGroupId() *string  {
  return s.ResourceGroupId
}

func (s *EnvironmentInfo) GetSubDomainInfos() []*SubDomainInfo  {
  return s.SubDomainInfos
}

func (s *EnvironmentInfo) GetUpdateTimestamp() *int64  {
  return s.UpdateTimestamp
}

func (s *EnvironmentInfo) SetAlias(v string) *EnvironmentInfo {
  s.Alias = &v
  return s
}

func (s *EnvironmentInfo) SetCreateTimestamp(v int64) *EnvironmentInfo {
  s.CreateTimestamp = &v
  return s
}

func (s *EnvironmentInfo) SetDefault(v bool) *EnvironmentInfo {
  s.Default = &v
  return s
}

func (s *EnvironmentInfo) SetDescription(v string) *EnvironmentInfo {
  s.Description = &v
  return s
}

func (s *EnvironmentInfo) SetEnvironmentId(v string) *EnvironmentInfo {
  s.EnvironmentId = &v
  return s
}

func (s *EnvironmentInfo) SetGatewayInfo(v *GatewayInfo) *EnvironmentInfo {
  s.GatewayInfo = v
  return s
}

func (s *EnvironmentInfo) SetName(v string) *EnvironmentInfo {
  s.Name = &v
  return s
}

func (s *EnvironmentInfo) SetResourceGroupId(v string) *EnvironmentInfo {
  s.ResourceGroupId = &v
  return s
}

func (s *EnvironmentInfo) SetSubDomainInfos(v []*SubDomainInfo) *EnvironmentInfo {
  s.SubDomainInfos = v
  return s
}

func (s *EnvironmentInfo) SetUpdateTimestamp(v int64) *EnvironmentInfo {
  s.UpdateTimestamp = &v
  return s
}

func (s *EnvironmentInfo) Validate() error {
  if s.GatewayInfo != nil {
    if err := s.GatewayInfo.Validate(); err != nil {
      return err
    }
  }
  if s.SubDomainInfos != nil {
    for _, item := range s.SubDomainInfos {
      if item != nil {
        if err := item.Validate(); err != nil {
          return err
        }
      }
    }
  }
  return nil
}

