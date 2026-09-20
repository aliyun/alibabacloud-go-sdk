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
  // The vSwitch ID of the arbitration zone. The vSwitch must be in the zone specified by **ArbiterZoneId**.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // vsw-hangxzhouxb****
  ArbiterVSwitchId *string `json:"ArbiterVSwitchId,omitempty" xml:"ArbiterVSwitchId,omitempty"`
  // The zone ID of the arbitration zone.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // cn-hangzhou-b
  ArbiterZoneId *string `json:"ArbiterZoneId,omitempty" xml:"ArbiterZoneId,omitempty"`
  // The version of the deployment architecture. Currently, only the hbaseue engine type is supported. Set the value to **2.0**.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // 2.0
  ArchVersion *string `json:"ArchVersion,omitempty" xml:"ArchVersion,omitempty"`
  // The auto-renewal period of the instance. Unit: months.
  // 
  // > <ul><li>The default value is 0, which indicates that the instance is not automatically renewed after the instance expires.</li>
  // 
  // <li>For example, if the auto-renewal period is set to 2, the instance is automatically renewed for two months after the instance expires.</li></ul>
  // 
  // example:
  // 
  // 0
  AutoRenewPeriod *int32 `json:"AutoRenewPeriod,omitempty" xml:"AutoRenewPeriod,omitempty"`
  // The client token that is used to ensure the idempotence of the request. You can use the client to generate the value. Make sure that the value is unique among different requests. The value cannot exceed 64 ASCII characters in length and cannot contain non-ASCII characters.
  // 
  // example:
  // 
  // f4g8t5rd2gr94****
  ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
  // The cluster name. The following rules apply:
  // 
  // - The name must be 2 to 128 characters in length.
  // 
  // - The name must start with an uppercase letter, a lowercase letter, or a Chinese character.
  // 
  // - The name can contain digits or special characters, including periods (.), hyphens (-), and underscores (_).
  // 
  // example:
  // 
  // hbaseue_test
  ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
  // The disk size of the node. Valid values: 400 to 64000. Unit: GB. The value must be a multiple of 40.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // 400
  CoreDiskSize *int32 `json:"CoreDiskSize,omitempty" xml:"CoreDiskSize,omitempty"`
  // The disk type of the core node. Valid values:
  // 
  // - **cloud_efficiency**: ultra cloud disk.
  // 
  // - **cloud_ssd**: standard SSD.
  // 
  // - **local_hdd_pro**: throughput-intensive local disk.
  // 
  // - **local_ssd_pro**: I/O-intensive local disk.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // cloud_ssd
  CoreDiskType *string `json:"CoreDiskType,omitempty" xml:"CoreDiskType,omitempty"`
  // The node specifications of the core node. You can invoke the [DescribeInstanceType](https://help.aliyun.com/document_detail/145796.html) operation to query the node specifications.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // hbase.sn1.medium
  CoreInstanceType *string `json:"CoreInstanceType,omitempty" xml:"CoreInstanceType,omitempty"`
  // The number of core nodes. Valid values: 2 to 20. The value must be an even number.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // 4
  CoreNodeCount *int32 `json:"CoreNodeCount,omitempty" xml:"CoreNodeCount,omitempty"`
  // The service type. Currently, only ApsaraDB for HBase Performance-enhanced Edition is supported. Set the value to **hbaseue**.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // hbaseue
  Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
  // The version of the engine type. Set the value to **2.0**.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // 2.0
  EngineVersion *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
  // The disk size of the log node. Valid values: 400 to 64000. Unit: GB. The value must be a multiple of 40.
  // 
  // example:
  // 
  // 400
  LogDiskSize *int32 `json:"LogDiskSize,omitempty" xml:"LogDiskSize,omitempty"`
  // The disk type of the log node. Valid values:
  // 
  // - **cloud_efficiency**: ultra cloud disk.
  // 
  // - **cloud_ssd**: standard SSD.
  // 
  // - **local_hdd_pro**: throughput-intensive local disk.
  // 
  // - **local_ssd_pro**: I/O-intensive local disk.
  // 
  // example:
  // 
  // cloud_ssd
  LogDiskType *string `json:"LogDiskType,omitempty" xml:"LogDiskType,omitempty"`
  // The node specifications of the log node. You can invoke the [DescribeInstanceType](https://help.aliyun.com/document_detail/145796.html) operation to query the node specifications.
  // 
  // example:
  // 
  // hbase.sn1.medium
  LogInstanceType *string `json:"LogInstanceType,omitempty" xml:"LogInstanceType,omitempty"`
  // The number of log nodes. Valid values: 4 to 400. The value must be a multiple of 4.
  // 
  // example:
  // 
  // 4
  LogNodeCount *int32 `json:"LogNodeCount,omitempty" xml:"LogNodeCount,omitempty"`
  // The node specifications of the master node. You can invoke the [DescribeInstanceType](https://help.aliyun.com/document_detail/145796.html) operation to query the node specifications.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // hbase.sn1.medium
  MasterInstanceType *string `json:"MasterInstanceType,omitempty" xml:"MasterInstanceType,omitempty"`
  // <props="china">The zone combination. The following combinations are supported. You can go to the buy page or call the [DescribeMultiZoneAvailableRegions](https://help.aliyun.com/document_detail/203039.html) operation to view the supported zone combinations.
  // 
  // <props="intl">The zone combination. The following combinations are supported. You can go to the buy page to view the supported zone combinations..
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // cn-hangzhou-bef-aliyun-com
  MultiZoneCombination *string `json:"MultiZoneCombination,omitempty" xml:"MultiZoneCombination,omitempty"`
  // The billing method of the instance. Valid values:
  // 
  // - **Prepaid**: subscription.
  // 
  // - **Postpaid**: pay-as-you-go.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // Postpaid
  PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
  // The subscription duration of the subscription instance. Valid values:
  // 
  // - If PeriodUnit is set to year, valid values are 1 to 3.
  // 
  // - If PeriodUnit is set to month, valid values are 1 to 9.
  // 
  // > This parameter is required only when PayType is set to Prepaid.
  // 
  // example:
  // 
  // 1
  Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
  // The unit of the subscription duration for the subscription instance. Valid values:
  // 
  // - **year**
  // 
  // - **month**
  // 
  // > This parameter is required only when PayType is set to Prepaid.
  // 
  // example:
  // 
  // month
  PeriodUnit *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
  // The vSwitch ID of the primary zone instance. The vSwitch must be in the zone specified by **PrimaryZoneId**.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // vsw-hangxzhouxe*****
  PrimaryVSwitchId *string `json:"PrimaryVSwitchId,omitempty" xml:"PrimaryVSwitchId,omitempty"`
  // The zone ID of the primary zone instance.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // cn-hangzhou-e
  PrimaryZoneId *string `json:"PrimaryZoneId,omitempty" xml:"PrimaryZoneId,omitempty"`
  // The ID of the region in which the instance resides. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/144489.html) operation to query the region ID.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // cn-hangzhou
  RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
  // The IP addresses in the whitelist of the instance. Separate multiple IP addresses with commas (,).
  // 
  // > If the IP address is set to 127.0.0.1, all addresses are denied access to the instance. For example, 192.168.0.0/24 indicates that all IP addresses in the 192.168.0.XX range are allowed to access the instance.
  // 
  // example:
  // 
  // 127.0.0.1
  SecurityIPList *string `json:"SecurityIPList,omitempty" xml:"SecurityIPList,omitempty"`
  // The vSwitch ID of the secondary zone instance. The vSwitch must be in the zone specified by **StandbyZoneId**.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // vsw-hangxzhouxf****
  StandbyVSwitchId *string `json:"StandbyVSwitchId,omitempty" xml:"StandbyVSwitchId,omitempty"`
  // The zone ID of the secondary zone instance.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // cn-hangzhou-f
  StandbyZoneId *string `json:"StandbyZoneId,omitempty" xml:"StandbyZoneId,omitempty"`
  // The ID of the virtual private cloud (VPC). The VPC must be in the region specified by **RegionId**.
  // 
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

