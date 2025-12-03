// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEvaluateMultiZoneResourceRequest interface {
  dara.Model
  String() string
  GoString() string
  SetArbiterVSwitchId(v string) *EvaluateMultiZoneResourceRequest
  GetArbiterVSwitchId() *string 
  SetArbiterZoneId(v string) *EvaluateMultiZoneResourceRequest
  GetArbiterZoneId() *string 
  SetArchVersion(v string) *EvaluateMultiZoneResourceRequest
  GetArchVersion() *string 
  SetAutoRenewPeriod(v int32) *EvaluateMultiZoneResourceRequest
  GetAutoRenewPeriod() *int32 
  SetClientToken(v string) *EvaluateMultiZoneResourceRequest
  GetClientToken() *string 
  SetClusterName(v string) *EvaluateMultiZoneResourceRequest
  GetClusterName() *string 
  SetCoreDiskSize(v int32) *EvaluateMultiZoneResourceRequest
  GetCoreDiskSize() *int32 
  SetCoreDiskType(v string) *EvaluateMultiZoneResourceRequest
  GetCoreDiskType() *string 
  SetCoreInstanceType(v string) *EvaluateMultiZoneResourceRequest
  GetCoreInstanceType() *string 
  SetCoreNodeCount(v int32) *EvaluateMultiZoneResourceRequest
  GetCoreNodeCount() *int32 
  SetEngine(v string) *EvaluateMultiZoneResourceRequest
  GetEngine() *string 
  SetEngineVersion(v string) *EvaluateMultiZoneResourceRequest
  GetEngineVersion() *string 
  SetLogDiskSize(v int32) *EvaluateMultiZoneResourceRequest
  GetLogDiskSize() *int32 
  SetLogDiskType(v string) *EvaluateMultiZoneResourceRequest
  GetLogDiskType() *string 
  SetLogInstanceType(v string) *EvaluateMultiZoneResourceRequest
  GetLogInstanceType() *string 
  SetLogNodeCount(v int32) *EvaluateMultiZoneResourceRequest
  GetLogNodeCount() *int32 
  SetMasterInstanceType(v string) *EvaluateMultiZoneResourceRequest
  GetMasterInstanceType() *string 
  SetMultiZoneCombination(v string) *EvaluateMultiZoneResourceRequest
  GetMultiZoneCombination() *string 
  SetPayType(v string) *EvaluateMultiZoneResourceRequest
  GetPayType() *string 
  SetPeriod(v int32) *EvaluateMultiZoneResourceRequest
  GetPeriod() *int32 
  SetPeriodUnit(v string) *EvaluateMultiZoneResourceRequest
  GetPeriodUnit() *string 
  SetPrimaryVSwitchId(v string) *EvaluateMultiZoneResourceRequest
  GetPrimaryVSwitchId() *string 
  SetPrimaryZoneId(v string) *EvaluateMultiZoneResourceRequest
  GetPrimaryZoneId() *string 
  SetRegionId(v string) *EvaluateMultiZoneResourceRequest
  GetRegionId() *string 
  SetSecurityIPList(v string) *EvaluateMultiZoneResourceRequest
  GetSecurityIPList() *string 
  SetStandbyVSwitchId(v string) *EvaluateMultiZoneResourceRequest
  GetStandbyVSwitchId() *string 
  SetStandbyZoneId(v string) *EvaluateMultiZoneResourceRequest
  GetStandbyZoneId() *string 
  SetVpcId(v string) *EvaluateMultiZoneResourceRequest
  GetVpcId() *string 
}

type EvaluateMultiZoneResourceRequest struct {
  // This parameter is required.
  // 
  // example:
  // 
  // vsw-hangxzhouxb****
  ArbiterVSwitchId *string `json:"ArbiterVSwitchId,omitempty" xml:"ArbiterVSwitchId,omitempty"`
  // This parameter is required.
  // 
  // example:
  // 
  // cn-hangzhou-b
  ArbiterZoneId *string `json:"ArbiterZoneId,omitempty" xml:"ArbiterZoneId,omitempty"`
  // This parameter is required.
  // 
  // example:
  // 
  // 2.0
  ArchVersion *string `json:"ArchVersion,omitempty" xml:"ArchVersion,omitempty"`
  // example:
  // 
  // 0
  AutoRenewPeriod *int32 `json:"AutoRenewPeriod,omitempty" xml:"AutoRenewPeriod,omitempty"`
  // example:
  // 
  // f4g8t5rd2gr94****
  ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
  // example:
  // 
  // hbaseue_test
  ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
  // This parameter is required.
  // 
  // example:
  // 
  // 400
  CoreDiskSize *int32 `json:"CoreDiskSize,omitempty" xml:"CoreDiskSize,omitempty"`
  // This parameter is required.
  // 
  // example:
  // 
  // cloud_ssd
  CoreDiskType *string `json:"CoreDiskType,omitempty" xml:"CoreDiskType,omitempty"`
  // This parameter is required.
  // 
  // example:
  // 
  // hbase.sn1.medium
  CoreInstanceType *string `json:"CoreInstanceType,omitempty" xml:"CoreInstanceType,omitempty"`
  // This parameter is required.
  // 
  // example:
  // 
  // 4
  CoreNodeCount *int32 `json:"CoreNodeCount,omitempty" xml:"CoreNodeCount,omitempty"`
  // This parameter is required.
  // 
  // example:
  // 
  // hbaseue
  Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
  // This parameter is required.
  // 
  // example:
  // 
  // 2.0
  EngineVersion *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
  // example:
  // 
  // 400
  LogDiskSize *int32 `json:"LogDiskSize,omitempty" xml:"LogDiskSize,omitempty"`
  // example:
  // 
  // cloud_ssd
  LogDiskType *string `json:"LogDiskType,omitempty" xml:"LogDiskType,omitempty"`
  // example:
  // 
  // hbase.sn1.medium
  LogInstanceType *string `json:"LogInstanceType,omitempty" xml:"LogInstanceType,omitempty"`
  // example:
  // 
  // 4
  LogNodeCount *int32 `json:"LogNodeCount,omitempty" xml:"LogNodeCount,omitempty"`
  // This parameter is required.
  // 
  // example:
  // 
  // hbase.sn1.medium
  MasterInstanceType *string `json:"MasterInstanceType,omitempty" xml:"MasterInstanceType,omitempty"`
  // This parameter is required.
  // 
  // example:
  // 
  // cn-hangzhou-bef-aliyun-com
  MultiZoneCombination *string `json:"MultiZoneCombination,omitempty" xml:"MultiZoneCombination,omitempty"`
  // This parameter is required.
  // 
  // example:
  // 
  // Postpaid
  PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
  // example:
  // 
  // 1
  Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
  // example:
  // 
  // month
  PeriodUnit *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
  // This parameter is required.
  // 
  // example:
  // 
  // vsw-hangxzhouxe*****
  PrimaryVSwitchId *string `json:"PrimaryVSwitchId,omitempty" xml:"PrimaryVSwitchId,omitempty"`
  // This parameter is required.
  // 
  // example:
  // 
  // cn-hangzhou-e
  PrimaryZoneId *string `json:"PrimaryZoneId,omitempty" xml:"PrimaryZoneId,omitempty"`
  // This parameter is required.
  // 
  // example:
  // 
  // cn-hangzhou
  RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
  // example:
  // 
  // 127.0.0.1
  SecurityIPList *string `json:"SecurityIPList,omitempty" xml:"SecurityIPList,omitempty"`
  // This parameter is required.
  // 
  // example:
  // 
  // vsw-hangxzhouxf****
  StandbyVSwitchId *string `json:"StandbyVSwitchId,omitempty" xml:"StandbyVSwitchId,omitempty"`
  // This parameter is required.
  // 
  // example:
  // 
  // cn-hangzhou-f
  StandbyZoneId *string `json:"StandbyZoneId,omitempty" xml:"StandbyZoneId,omitempty"`
  // This parameter is required.
  // 
  // example:
  // 
  // vpc-bp120k6ixs4eog*****
  VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s EvaluateMultiZoneResourceRequest) String() string {
  return dara.Prettify(s)
}

func (s EvaluateMultiZoneResourceRequest) GoString() string {
  return s.String()
}

func (s *EvaluateMultiZoneResourceRequest) GetArbiterVSwitchId() *string  {
  return s.ArbiterVSwitchId
}

func (s *EvaluateMultiZoneResourceRequest) GetArbiterZoneId() *string  {
  return s.ArbiterZoneId
}

func (s *EvaluateMultiZoneResourceRequest) GetArchVersion() *string  {
  return s.ArchVersion
}

func (s *EvaluateMultiZoneResourceRequest) GetAutoRenewPeriod() *int32  {
  return s.AutoRenewPeriod
}

func (s *EvaluateMultiZoneResourceRequest) GetClientToken() *string  {
  return s.ClientToken
}

func (s *EvaluateMultiZoneResourceRequest) GetClusterName() *string  {
  return s.ClusterName
}

func (s *EvaluateMultiZoneResourceRequest) GetCoreDiskSize() *int32  {
  return s.CoreDiskSize
}

func (s *EvaluateMultiZoneResourceRequest) GetCoreDiskType() *string  {
  return s.CoreDiskType
}

func (s *EvaluateMultiZoneResourceRequest) GetCoreInstanceType() *string  {
  return s.CoreInstanceType
}

func (s *EvaluateMultiZoneResourceRequest) GetCoreNodeCount() *int32  {
  return s.CoreNodeCount
}

func (s *EvaluateMultiZoneResourceRequest) GetEngine() *string  {
  return s.Engine
}

func (s *EvaluateMultiZoneResourceRequest) GetEngineVersion() *string  {
  return s.EngineVersion
}

func (s *EvaluateMultiZoneResourceRequest) GetLogDiskSize() *int32  {
  return s.LogDiskSize
}

func (s *EvaluateMultiZoneResourceRequest) GetLogDiskType() *string  {
  return s.LogDiskType
}

func (s *EvaluateMultiZoneResourceRequest) GetLogInstanceType() *string  {
  return s.LogInstanceType
}

func (s *EvaluateMultiZoneResourceRequest) GetLogNodeCount() *int32  {
  return s.LogNodeCount
}

func (s *EvaluateMultiZoneResourceRequest) GetMasterInstanceType() *string  {
  return s.MasterInstanceType
}

func (s *EvaluateMultiZoneResourceRequest) GetMultiZoneCombination() *string  {
  return s.MultiZoneCombination
}

func (s *EvaluateMultiZoneResourceRequest) GetPayType() *string  {
  return s.PayType
}

func (s *EvaluateMultiZoneResourceRequest) GetPeriod() *int32  {
  return s.Period
}

func (s *EvaluateMultiZoneResourceRequest) GetPeriodUnit() *string  {
  return s.PeriodUnit
}

func (s *EvaluateMultiZoneResourceRequest) GetPrimaryVSwitchId() *string  {
  return s.PrimaryVSwitchId
}

func (s *EvaluateMultiZoneResourceRequest) GetPrimaryZoneId() *string  {
  return s.PrimaryZoneId
}

func (s *EvaluateMultiZoneResourceRequest) GetRegionId() *string  {
  return s.RegionId
}

func (s *EvaluateMultiZoneResourceRequest) GetSecurityIPList() *string  {
  return s.SecurityIPList
}

func (s *EvaluateMultiZoneResourceRequest) GetStandbyVSwitchId() *string  {
  return s.StandbyVSwitchId
}

func (s *EvaluateMultiZoneResourceRequest) GetStandbyZoneId() *string  {
  return s.StandbyZoneId
}

func (s *EvaluateMultiZoneResourceRequest) GetVpcId() *string  {
  return s.VpcId
}

func (s *EvaluateMultiZoneResourceRequest) SetArbiterVSwitchId(v string) *EvaluateMultiZoneResourceRequest {
  s.ArbiterVSwitchId = &v
  return s
}

func (s *EvaluateMultiZoneResourceRequest) SetArbiterZoneId(v string) *EvaluateMultiZoneResourceRequest {
  s.ArbiterZoneId = &v
  return s
}

func (s *EvaluateMultiZoneResourceRequest) SetArchVersion(v string) *EvaluateMultiZoneResourceRequest {
  s.ArchVersion = &v
  return s
}

func (s *EvaluateMultiZoneResourceRequest) SetAutoRenewPeriod(v int32) *EvaluateMultiZoneResourceRequest {
  s.AutoRenewPeriod = &v
  return s
}

func (s *EvaluateMultiZoneResourceRequest) SetClientToken(v string) *EvaluateMultiZoneResourceRequest {
  s.ClientToken = &v
  return s
}

func (s *EvaluateMultiZoneResourceRequest) SetClusterName(v string) *EvaluateMultiZoneResourceRequest {
  s.ClusterName = &v
  return s
}

func (s *EvaluateMultiZoneResourceRequest) SetCoreDiskSize(v int32) *EvaluateMultiZoneResourceRequest {
  s.CoreDiskSize = &v
  return s
}

func (s *EvaluateMultiZoneResourceRequest) SetCoreDiskType(v string) *EvaluateMultiZoneResourceRequest {
  s.CoreDiskType = &v
  return s
}

func (s *EvaluateMultiZoneResourceRequest) SetCoreInstanceType(v string) *EvaluateMultiZoneResourceRequest {
  s.CoreInstanceType = &v
  return s
}

func (s *EvaluateMultiZoneResourceRequest) SetCoreNodeCount(v int32) *EvaluateMultiZoneResourceRequest {
  s.CoreNodeCount = &v
  return s
}

func (s *EvaluateMultiZoneResourceRequest) SetEngine(v string) *EvaluateMultiZoneResourceRequest {
  s.Engine = &v
  return s
}

func (s *EvaluateMultiZoneResourceRequest) SetEngineVersion(v string) *EvaluateMultiZoneResourceRequest {
  s.EngineVersion = &v
  return s
}

func (s *EvaluateMultiZoneResourceRequest) SetLogDiskSize(v int32) *EvaluateMultiZoneResourceRequest {
  s.LogDiskSize = &v
  return s
}

func (s *EvaluateMultiZoneResourceRequest) SetLogDiskType(v string) *EvaluateMultiZoneResourceRequest {
  s.LogDiskType = &v
  return s
}

func (s *EvaluateMultiZoneResourceRequest) SetLogInstanceType(v string) *EvaluateMultiZoneResourceRequest {
  s.LogInstanceType = &v
  return s
}

func (s *EvaluateMultiZoneResourceRequest) SetLogNodeCount(v int32) *EvaluateMultiZoneResourceRequest {
  s.LogNodeCount = &v
  return s
}

func (s *EvaluateMultiZoneResourceRequest) SetMasterInstanceType(v string) *EvaluateMultiZoneResourceRequest {
  s.MasterInstanceType = &v
  return s
}

func (s *EvaluateMultiZoneResourceRequest) SetMultiZoneCombination(v string) *EvaluateMultiZoneResourceRequest {
  s.MultiZoneCombination = &v
  return s
}

func (s *EvaluateMultiZoneResourceRequest) SetPayType(v string) *EvaluateMultiZoneResourceRequest {
  s.PayType = &v
  return s
}

func (s *EvaluateMultiZoneResourceRequest) SetPeriod(v int32) *EvaluateMultiZoneResourceRequest {
  s.Period = &v
  return s
}

func (s *EvaluateMultiZoneResourceRequest) SetPeriodUnit(v string) *EvaluateMultiZoneResourceRequest {
  s.PeriodUnit = &v
  return s
}

func (s *EvaluateMultiZoneResourceRequest) SetPrimaryVSwitchId(v string) *EvaluateMultiZoneResourceRequest {
  s.PrimaryVSwitchId = &v
  return s
}

func (s *EvaluateMultiZoneResourceRequest) SetPrimaryZoneId(v string) *EvaluateMultiZoneResourceRequest {
  s.PrimaryZoneId = &v
  return s
}

func (s *EvaluateMultiZoneResourceRequest) SetRegionId(v string) *EvaluateMultiZoneResourceRequest {
  s.RegionId = &v
  return s
}

func (s *EvaluateMultiZoneResourceRequest) SetSecurityIPList(v string) *EvaluateMultiZoneResourceRequest {
  s.SecurityIPList = &v
  return s
}

func (s *EvaluateMultiZoneResourceRequest) SetStandbyVSwitchId(v string) *EvaluateMultiZoneResourceRequest {
  s.StandbyVSwitchId = &v
  return s
}

func (s *EvaluateMultiZoneResourceRequest) SetStandbyZoneId(v string) *EvaluateMultiZoneResourceRequest {
  s.StandbyZoneId = &v
  return s
}

func (s *EvaluateMultiZoneResourceRequest) SetVpcId(v string) *EvaluateMultiZoneResourceRequest {
  s.VpcId = &v
  return s
}

func (s *EvaluateMultiZoneResourceRequest) Validate() error {
  return dara.Validate(s)
}

