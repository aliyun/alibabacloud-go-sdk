// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableHBaseueModuleRequest interface {
  dara.Model
  String() string
  GoString() string
  SetAutoRenewPeriod(v int32) *EnableHBaseueModuleRequest
  GetAutoRenewPeriod() *int32 
  SetBdsId(v string) *EnableHBaseueModuleRequest
  GetBdsId() *string 
  SetClientToken(v string) *EnableHBaseueModuleRequest
  GetClientToken() *string 
  SetCoreInstanceType(v string) *EnableHBaseueModuleRequest
  GetCoreInstanceType() *string 
  SetDiskSize(v int32) *EnableHBaseueModuleRequest
  GetDiskSize() *int32 
  SetDiskType(v string) *EnableHBaseueModuleRequest
  GetDiskType() *string 
  SetHbaseueClusterId(v string) *EnableHBaseueModuleRequest
  GetHbaseueClusterId() *string 
  SetMasterInstanceType(v string) *EnableHBaseueModuleRequest
  GetMasterInstanceType() *string 
  SetModuleClusterName(v string) *EnableHBaseueModuleRequest
  GetModuleClusterName() *string 
  SetModuleTypeName(v string) *EnableHBaseueModuleRequest
  GetModuleTypeName() *string 
  SetNodeCount(v int32) *EnableHBaseueModuleRequest
  GetNodeCount() *int32 
  SetPayType(v string) *EnableHBaseueModuleRequest
  GetPayType() *string 
  SetPeriod(v int32) *EnableHBaseueModuleRequest
  GetPeriod() *int32 
  SetPeriodUnit(v string) *EnableHBaseueModuleRequest
  GetPeriodUnit() *string 
  SetRegionId(v string) *EnableHBaseueModuleRequest
  GetRegionId() *string 
  SetVpcId(v string) *EnableHBaseueModuleRequest
  GetVpcId() *string 
  SetVswitchId(v string) *EnableHBaseueModuleRequest
  GetVswitchId() *string 
  SetZoneId(v string) *EnableHBaseueModuleRequest
  GetZoneId() *string 
}

type EnableHBaseueModuleRequest struct {
  // example:
  // 
  // 2
  AutoRenewPeriod *int32 `json:"AutoRenewPeriod,omitempty" xml:"AutoRenewPeriod,omitempty"`
  // example:
  // 
  // bds-bp174pm3tsk3****
  BdsId *string `json:"BdsId,omitempty" xml:"BdsId,omitempty"`
  // example:
  // 
  // ETnLKlblzczshOTUbOCz****
  ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
  // This parameter is required.
  // 
  // example:
  // 
  // hbase.sn1.large
  CoreInstanceType *string `json:"CoreInstanceType,omitempty" xml:"CoreInstanceType,omitempty"`
  // example:
  // 
  // 400
  DiskSize *int32 `json:"DiskSize,omitempty" xml:"DiskSize,omitempty"`
  // example:
  // 
  // cloud_ssd
  DiskType *string `json:"DiskType,omitempty" xml:"DiskType,omitempty"`
  // This parameter is required.
  // 
  // example:
  // 
  // ld-bp150tns0sjxs****
  HbaseueClusterId *string `json:"HbaseueClusterId,omitempty" xml:"HbaseueClusterId,omitempty"`
  // example:
  // 
  // hbase.sn1.large
  MasterInstanceType *string `json:"MasterInstanceType,omitempty" xml:"MasterInstanceType,omitempty"`
  // example:
  // 
  // cluster-name
  ModuleClusterName *string `json:"ModuleClusterName,omitempty" xml:"ModuleClusterName,omitempty"`
  // This parameter is required.
  // 
  // example:
  // 
  // solr
  ModuleTypeName *string `json:"ModuleTypeName,omitempty" xml:"ModuleTypeName,omitempty"`
  // This parameter is required.
  // 
  // example:
  // 
  // 2
  NodeCount *int32 `json:"NodeCount,omitempty" xml:"NodeCount,omitempty"`
  // This parameter is required.
  // 
  // example:
  // 
  // Prepaid
  PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
  // example:
  // 
  // 6
  Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
  // example:
  // 
  // month
  PeriodUnit *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
  // This parameter is required.
  // 
  // example:
  // 
  // cn-shenzhen
  RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
  // This parameter is required.
  // 
  // example:
  // 
  // vpc-bp120k6ixs4eog*****
  VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
  // This parameter is required.
  // 
  // example:
  // 
  // vsw-bp191ipotqj1ssyl*****
  VswitchId *string `json:"VswitchId,omitempty" xml:"VswitchId,omitempty"`
  // This parameter is required.
  // 
  // example:
  // 
  // cn-shenzhen-e
  ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s EnableHBaseueModuleRequest) String() string {
  return dara.Prettify(s)
}

func (s EnableHBaseueModuleRequest) GoString() string {
  return s.String()
}

func (s *EnableHBaseueModuleRequest) GetAutoRenewPeriod() *int32  {
  return s.AutoRenewPeriod
}

func (s *EnableHBaseueModuleRequest) GetBdsId() *string  {
  return s.BdsId
}

func (s *EnableHBaseueModuleRequest) GetClientToken() *string  {
  return s.ClientToken
}

func (s *EnableHBaseueModuleRequest) GetCoreInstanceType() *string  {
  return s.CoreInstanceType
}

func (s *EnableHBaseueModuleRequest) GetDiskSize() *int32  {
  return s.DiskSize
}

func (s *EnableHBaseueModuleRequest) GetDiskType() *string  {
  return s.DiskType
}

func (s *EnableHBaseueModuleRequest) GetHbaseueClusterId() *string  {
  return s.HbaseueClusterId
}

func (s *EnableHBaseueModuleRequest) GetMasterInstanceType() *string  {
  return s.MasterInstanceType
}

func (s *EnableHBaseueModuleRequest) GetModuleClusterName() *string  {
  return s.ModuleClusterName
}

func (s *EnableHBaseueModuleRequest) GetModuleTypeName() *string  {
  return s.ModuleTypeName
}

func (s *EnableHBaseueModuleRequest) GetNodeCount() *int32  {
  return s.NodeCount
}

func (s *EnableHBaseueModuleRequest) GetPayType() *string  {
  return s.PayType
}

func (s *EnableHBaseueModuleRequest) GetPeriod() *int32  {
  return s.Period
}

func (s *EnableHBaseueModuleRequest) GetPeriodUnit() *string  {
  return s.PeriodUnit
}

func (s *EnableHBaseueModuleRequest) GetRegionId() *string  {
  return s.RegionId
}

func (s *EnableHBaseueModuleRequest) GetVpcId() *string  {
  return s.VpcId
}

func (s *EnableHBaseueModuleRequest) GetVswitchId() *string  {
  return s.VswitchId
}

func (s *EnableHBaseueModuleRequest) GetZoneId() *string  {
  return s.ZoneId
}

func (s *EnableHBaseueModuleRequest) SetAutoRenewPeriod(v int32) *EnableHBaseueModuleRequest {
  s.AutoRenewPeriod = &v
  return s
}

func (s *EnableHBaseueModuleRequest) SetBdsId(v string) *EnableHBaseueModuleRequest {
  s.BdsId = &v
  return s
}

func (s *EnableHBaseueModuleRequest) SetClientToken(v string) *EnableHBaseueModuleRequest {
  s.ClientToken = &v
  return s
}

func (s *EnableHBaseueModuleRequest) SetCoreInstanceType(v string) *EnableHBaseueModuleRequest {
  s.CoreInstanceType = &v
  return s
}

func (s *EnableHBaseueModuleRequest) SetDiskSize(v int32) *EnableHBaseueModuleRequest {
  s.DiskSize = &v
  return s
}

func (s *EnableHBaseueModuleRequest) SetDiskType(v string) *EnableHBaseueModuleRequest {
  s.DiskType = &v
  return s
}

func (s *EnableHBaseueModuleRequest) SetHbaseueClusterId(v string) *EnableHBaseueModuleRequest {
  s.HbaseueClusterId = &v
  return s
}

func (s *EnableHBaseueModuleRequest) SetMasterInstanceType(v string) *EnableHBaseueModuleRequest {
  s.MasterInstanceType = &v
  return s
}

func (s *EnableHBaseueModuleRequest) SetModuleClusterName(v string) *EnableHBaseueModuleRequest {
  s.ModuleClusterName = &v
  return s
}

func (s *EnableHBaseueModuleRequest) SetModuleTypeName(v string) *EnableHBaseueModuleRequest {
  s.ModuleTypeName = &v
  return s
}

func (s *EnableHBaseueModuleRequest) SetNodeCount(v int32) *EnableHBaseueModuleRequest {
  s.NodeCount = &v
  return s
}

func (s *EnableHBaseueModuleRequest) SetPayType(v string) *EnableHBaseueModuleRequest {
  s.PayType = &v
  return s
}

func (s *EnableHBaseueModuleRequest) SetPeriod(v int32) *EnableHBaseueModuleRequest {
  s.Period = &v
  return s
}

func (s *EnableHBaseueModuleRequest) SetPeriodUnit(v string) *EnableHBaseueModuleRequest {
  s.PeriodUnit = &v
  return s
}

func (s *EnableHBaseueModuleRequest) SetRegionId(v string) *EnableHBaseueModuleRequest {
  s.RegionId = &v
  return s
}

func (s *EnableHBaseueModuleRequest) SetVpcId(v string) *EnableHBaseueModuleRequest {
  s.VpcId = &v
  return s
}

func (s *EnableHBaseueModuleRequest) SetVswitchId(v string) *EnableHBaseueModuleRequest {
  s.VswitchId = &v
  return s
}

func (s *EnableHBaseueModuleRequest) SetZoneId(v string) *EnableHBaseueModuleRequest {
  s.ZoneId = &v
  return s
}

func (s *EnableHBaseueModuleRequest) Validate() error {
  return dara.Validate(s)
}

