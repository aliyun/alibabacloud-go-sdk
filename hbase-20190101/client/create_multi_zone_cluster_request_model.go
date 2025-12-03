// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMultiZoneClusterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetArbiterVSwitchId(v string) *CreateMultiZoneClusterRequest
	GetArbiterVSwitchId() *string
	SetArbiterZoneId(v string) *CreateMultiZoneClusterRequest
	GetArbiterZoneId() *string
	SetArchVersion(v string) *CreateMultiZoneClusterRequest
	GetArchVersion() *string
	SetAutoRenewPeriod(v int32) *CreateMultiZoneClusterRequest
	GetAutoRenewPeriod() *int32
	SetClientToken(v string) *CreateMultiZoneClusterRequest
	GetClientToken() *string
	SetClusterName(v string) *CreateMultiZoneClusterRequest
	GetClusterName() *string
	SetCoreDiskSize(v int32) *CreateMultiZoneClusterRequest
	GetCoreDiskSize() *int32
	SetCoreDiskType(v string) *CreateMultiZoneClusterRequest
	GetCoreDiskType() *string
	SetCoreInstanceType(v string) *CreateMultiZoneClusterRequest
	GetCoreInstanceType() *string
	SetCoreNodeCount(v int32) *CreateMultiZoneClusterRequest
	GetCoreNodeCount() *int32
	SetEngine(v string) *CreateMultiZoneClusterRequest
	GetEngine() *string
	SetEngineVersion(v string) *CreateMultiZoneClusterRequest
	GetEngineVersion() *string
	SetLogDiskSize(v int32) *CreateMultiZoneClusterRequest
	GetLogDiskSize() *int32
	SetLogDiskType(v string) *CreateMultiZoneClusterRequest
	GetLogDiskType() *string
	SetLogInstanceType(v string) *CreateMultiZoneClusterRequest
	GetLogInstanceType() *string
	SetLogNodeCount(v int32) *CreateMultiZoneClusterRequest
	GetLogNodeCount() *int32
	SetMasterInstanceType(v string) *CreateMultiZoneClusterRequest
	GetMasterInstanceType() *string
	SetMultiZoneCombination(v string) *CreateMultiZoneClusterRequest
	GetMultiZoneCombination() *string
	SetPayType(v string) *CreateMultiZoneClusterRequest
	GetPayType() *string
	SetPeriod(v int32) *CreateMultiZoneClusterRequest
	GetPeriod() *int32
	SetPeriodUnit(v string) *CreateMultiZoneClusterRequest
	GetPeriodUnit() *string
	SetPrimaryVSwitchId(v string) *CreateMultiZoneClusterRequest
	GetPrimaryVSwitchId() *string
	SetPrimaryZoneId(v string) *CreateMultiZoneClusterRequest
	GetPrimaryZoneId() *string
	SetRegionId(v string) *CreateMultiZoneClusterRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreateMultiZoneClusterRequest
	GetResourceGroupId() *string
	SetSecurityIPList(v string) *CreateMultiZoneClusterRequest
	GetSecurityIPList() *string
	SetStandbyVSwitchId(v string) *CreateMultiZoneClusterRequest
	GetStandbyVSwitchId() *string
	SetStandbyZoneId(v string) *CreateMultiZoneClusterRequest
	GetStandbyZoneId() *string
	SetVpcId(v string) *CreateMultiZoneClusterRequest
	GetVpcId() *string
}

type CreateMultiZoneClusterRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// vsw-hangxzhouxb*****
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
	// dfh3sf5gslfksfk****
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
	// This parameter is required.
	//
	// example:
	//
	// 400
	LogDiskSize *int32 `json:"LogDiskSize,omitempty" xml:"LogDiskSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cloud_ssd
	LogDiskType *string `json:"LogDiskType,omitempty" xml:"LogDiskType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// hbase.sn1.medium
	LogInstanceType *string `json:"LogInstanceType,omitempty" xml:"LogInstanceType,omitempty"`
	// This parameter is required.
	//
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
	// vsw-hangxzhouxe****
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
	// rg-gg3f4f5d5g5w****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
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
	// vpc-bp120k6ixs4eog****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s CreateMultiZoneClusterRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateMultiZoneClusterRequest) GoString() string {
	return s.String()
}

func (s *CreateMultiZoneClusterRequest) GetArbiterVSwitchId() *string {
	return s.ArbiterVSwitchId
}

func (s *CreateMultiZoneClusterRequest) GetArbiterZoneId() *string {
	return s.ArbiterZoneId
}

func (s *CreateMultiZoneClusterRequest) GetArchVersion() *string {
	return s.ArchVersion
}

func (s *CreateMultiZoneClusterRequest) GetAutoRenewPeriod() *int32 {
	return s.AutoRenewPeriod
}

func (s *CreateMultiZoneClusterRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateMultiZoneClusterRequest) GetClusterName() *string {
	return s.ClusterName
}

func (s *CreateMultiZoneClusterRequest) GetCoreDiskSize() *int32 {
	return s.CoreDiskSize
}

func (s *CreateMultiZoneClusterRequest) GetCoreDiskType() *string {
	return s.CoreDiskType
}

func (s *CreateMultiZoneClusterRequest) GetCoreInstanceType() *string {
	return s.CoreInstanceType
}

func (s *CreateMultiZoneClusterRequest) GetCoreNodeCount() *int32 {
	return s.CoreNodeCount
}

func (s *CreateMultiZoneClusterRequest) GetEngine() *string {
	return s.Engine
}

func (s *CreateMultiZoneClusterRequest) GetEngineVersion() *string {
	return s.EngineVersion
}

func (s *CreateMultiZoneClusterRequest) GetLogDiskSize() *int32 {
	return s.LogDiskSize
}

func (s *CreateMultiZoneClusterRequest) GetLogDiskType() *string {
	return s.LogDiskType
}

func (s *CreateMultiZoneClusterRequest) GetLogInstanceType() *string {
	return s.LogInstanceType
}

func (s *CreateMultiZoneClusterRequest) GetLogNodeCount() *int32 {
	return s.LogNodeCount
}

func (s *CreateMultiZoneClusterRequest) GetMasterInstanceType() *string {
	return s.MasterInstanceType
}

func (s *CreateMultiZoneClusterRequest) GetMultiZoneCombination() *string {
	return s.MultiZoneCombination
}

func (s *CreateMultiZoneClusterRequest) GetPayType() *string {
	return s.PayType
}

func (s *CreateMultiZoneClusterRequest) GetPeriod() *int32 {
	return s.Period
}

func (s *CreateMultiZoneClusterRequest) GetPeriodUnit() *string {
	return s.PeriodUnit
}

func (s *CreateMultiZoneClusterRequest) GetPrimaryVSwitchId() *string {
	return s.PrimaryVSwitchId
}

func (s *CreateMultiZoneClusterRequest) GetPrimaryZoneId() *string {
	return s.PrimaryZoneId
}

func (s *CreateMultiZoneClusterRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateMultiZoneClusterRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateMultiZoneClusterRequest) GetSecurityIPList() *string {
	return s.SecurityIPList
}

func (s *CreateMultiZoneClusterRequest) GetStandbyVSwitchId() *string {
	return s.StandbyVSwitchId
}

func (s *CreateMultiZoneClusterRequest) GetStandbyZoneId() *string {
	return s.StandbyZoneId
}

func (s *CreateMultiZoneClusterRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *CreateMultiZoneClusterRequest) SetArbiterVSwitchId(v string) *CreateMultiZoneClusterRequest {
	s.ArbiterVSwitchId = &v
	return s
}

func (s *CreateMultiZoneClusterRequest) SetArbiterZoneId(v string) *CreateMultiZoneClusterRequest {
	s.ArbiterZoneId = &v
	return s
}

func (s *CreateMultiZoneClusterRequest) SetArchVersion(v string) *CreateMultiZoneClusterRequest {
	s.ArchVersion = &v
	return s
}

func (s *CreateMultiZoneClusterRequest) SetAutoRenewPeriod(v int32) *CreateMultiZoneClusterRequest {
	s.AutoRenewPeriod = &v
	return s
}

func (s *CreateMultiZoneClusterRequest) SetClientToken(v string) *CreateMultiZoneClusterRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateMultiZoneClusterRequest) SetClusterName(v string) *CreateMultiZoneClusterRequest {
	s.ClusterName = &v
	return s
}

func (s *CreateMultiZoneClusterRequest) SetCoreDiskSize(v int32) *CreateMultiZoneClusterRequest {
	s.CoreDiskSize = &v
	return s
}

func (s *CreateMultiZoneClusterRequest) SetCoreDiskType(v string) *CreateMultiZoneClusterRequest {
	s.CoreDiskType = &v
	return s
}

func (s *CreateMultiZoneClusterRequest) SetCoreInstanceType(v string) *CreateMultiZoneClusterRequest {
	s.CoreInstanceType = &v
	return s
}

func (s *CreateMultiZoneClusterRequest) SetCoreNodeCount(v int32) *CreateMultiZoneClusterRequest {
	s.CoreNodeCount = &v
	return s
}

func (s *CreateMultiZoneClusterRequest) SetEngine(v string) *CreateMultiZoneClusterRequest {
	s.Engine = &v
	return s
}

func (s *CreateMultiZoneClusterRequest) SetEngineVersion(v string) *CreateMultiZoneClusterRequest {
	s.EngineVersion = &v
	return s
}

func (s *CreateMultiZoneClusterRequest) SetLogDiskSize(v int32) *CreateMultiZoneClusterRequest {
	s.LogDiskSize = &v
	return s
}

func (s *CreateMultiZoneClusterRequest) SetLogDiskType(v string) *CreateMultiZoneClusterRequest {
	s.LogDiskType = &v
	return s
}

func (s *CreateMultiZoneClusterRequest) SetLogInstanceType(v string) *CreateMultiZoneClusterRequest {
	s.LogInstanceType = &v
	return s
}

func (s *CreateMultiZoneClusterRequest) SetLogNodeCount(v int32) *CreateMultiZoneClusterRequest {
	s.LogNodeCount = &v
	return s
}

func (s *CreateMultiZoneClusterRequest) SetMasterInstanceType(v string) *CreateMultiZoneClusterRequest {
	s.MasterInstanceType = &v
	return s
}

func (s *CreateMultiZoneClusterRequest) SetMultiZoneCombination(v string) *CreateMultiZoneClusterRequest {
	s.MultiZoneCombination = &v
	return s
}

func (s *CreateMultiZoneClusterRequest) SetPayType(v string) *CreateMultiZoneClusterRequest {
	s.PayType = &v
	return s
}

func (s *CreateMultiZoneClusterRequest) SetPeriod(v int32) *CreateMultiZoneClusterRequest {
	s.Period = &v
	return s
}

func (s *CreateMultiZoneClusterRequest) SetPeriodUnit(v string) *CreateMultiZoneClusterRequest {
	s.PeriodUnit = &v
	return s
}

func (s *CreateMultiZoneClusterRequest) SetPrimaryVSwitchId(v string) *CreateMultiZoneClusterRequest {
	s.PrimaryVSwitchId = &v
	return s
}

func (s *CreateMultiZoneClusterRequest) SetPrimaryZoneId(v string) *CreateMultiZoneClusterRequest {
	s.PrimaryZoneId = &v
	return s
}

func (s *CreateMultiZoneClusterRequest) SetRegionId(v string) *CreateMultiZoneClusterRequest {
	s.RegionId = &v
	return s
}

func (s *CreateMultiZoneClusterRequest) SetResourceGroupId(v string) *CreateMultiZoneClusterRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateMultiZoneClusterRequest) SetSecurityIPList(v string) *CreateMultiZoneClusterRequest {
	s.SecurityIPList = &v
	return s
}

func (s *CreateMultiZoneClusterRequest) SetStandbyVSwitchId(v string) *CreateMultiZoneClusterRequest {
	s.StandbyVSwitchId = &v
	return s
}

func (s *CreateMultiZoneClusterRequest) SetStandbyZoneId(v string) *CreateMultiZoneClusterRequest {
	s.StandbyZoneId = &v
	return s
}

func (s *CreateMultiZoneClusterRequest) SetVpcId(v string) *CreateMultiZoneClusterRequest {
	s.VpcId = &v
	return s
}

func (s *CreateMultiZoneClusterRequest) Validate() error {
	return dara.Validate(s)
}
