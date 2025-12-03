// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateClusterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoRenewPeriod(v int32) *CreateClusterRequest
	GetAutoRenewPeriod() *int32
	SetClientToken(v string) *CreateClusterRequest
	GetClientToken() *string
	SetClusterName(v string) *CreateClusterRequest
	GetClusterName() *string
	SetColdStorageSize(v int32) *CreateClusterRequest
	GetColdStorageSize() *int32
	SetCoreInstanceType(v string) *CreateClusterRequest
	GetCoreInstanceType() *string
	SetDiskSize(v int32) *CreateClusterRequest
	GetDiskSize() *int32
	SetDiskType(v string) *CreateClusterRequest
	GetDiskType() *string
	SetEncryptionKey(v string) *CreateClusterRequest
	GetEncryptionKey() *string
	SetEngine(v string) *CreateClusterRequest
	GetEngine() *string
	SetEngineVersion(v string) *CreateClusterRequest
	GetEngineVersion() *string
	SetMasterInstanceType(v string) *CreateClusterRequest
	GetMasterInstanceType() *string
	SetNodeCount(v int32) *CreateClusterRequest
	GetNodeCount() *int32
	SetPayType(v string) *CreateClusterRequest
	GetPayType() *string
	SetPeriod(v int32) *CreateClusterRequest
	GetPeriod() *int32
	SetPeriodUnit(v string) *CreateClusterRequest
	GetPeriodUnit() *string
	SetRegionId(v string) *CreateClusterRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreateClusterRequest
	GetResourceGroupId() *string
	SetSecurityIPList(v string) *CreateClusterRequest
	GetSecurityIPList() *string
	SetVSwitchId(v string) *CreateClusterRequest
	GetVSwitchId() *string
	SetVpcId(v string) *CreateClusterRequest
	GetVpcId() *string
	SetZoneId(v string) *CreateClusterRequest
	GetZoneId() *string
}

type CreateClusterRequest struct {
	// example:
	//
	// 2
	AutoRenewPeriod *int32 `json:"AutoRenewPeriod,omitempty" xml:"AutoRenewPeriod,omitempty"`
	// example:
	//
	// ETnLKlblzczshOTUbOCz****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// example:
	//
	// hbase_test
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// example:
	//
	// 1024
	ColdStorageSize *int32 `json:"ColdStorageSize,omitempty" xml:"ColdStorageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// hbase.sn1.medium
	CoreInstanceType *string `json:"CoreInstanceType,omitempty" xml:"CoreInstanceType,omitempty"`
	// example:
	//
	// 400
	DiskSize *int32 `json:"DiskSize,omitempty" xml:"DiskSize,omitempty"`
	// example:
	//
	// cloud_ssd
	DiskType *string `json:"DiskType,omitempty" xml:"DiskType,omitempty"`
	// example:
	//
	// 0d2470df-da7b-4786-b981-9a164dae****
	EncryptionKey *string `json:"EncryptionKey,omitempty" xml:"EncryptionKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// hbase
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2.0
	EngineVersion *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	// example:
	//
	// hbase.sn1.medium
	MasterInstanceType *string `json:"MasterInstanceType,omitempty" xml:"MasterInstanceType,omitempty"`
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
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// rg-j4d53glb3****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// example:
	//
	// 116.62.XX.XX/24
	SecurityIPList *string `json:"SecurityIPList,omitempty" xml:"SecurityIPList,omitempty"`
	// example:
	//
	// vsw-bp191otqj1ssyl****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// example:
	//
	// vpc-bp120k6ixs4eog****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou-f
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s CreateClusterRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateClusterRequest) GoString() string {
	return s.String()
}

func (s *CreateClusterRequest) GetAutoRenewPeriod() *int32 {
	return s.AutoRenewPeriod
}

func (s *CreateClusterRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateClusterRequest) GetClusterName() *string {
	return s.ClusterName
}

func (s *CreateClusterRequest) GetColdStorageSize() *int32 {
	return s.ColdStorageSize
}

func (s *CreateClusterRequest) GetCoreInstanceType() *string {
	return s.CoreInstanceType
}

func (s *CreateClusterRequest) GetDiskSize() *int32 {
	return s.DiskSize
}

func (s *CreateClusterRequest) GetDiskType() *string {
	return s.DiskType
}

func (s *CreateClusterRequest) GetEncryptionKey() *string {
	return s.EncryptionKey
}

func (s *CreateClusterRequest) GetEngine() *string {
	return s.Engine
}

func (s *CreateClusterRequest) GetEngineVersion() *string {
	return s.EngineVersion
}

func (s *CreateClusterRequest) GetMasterInstanceType() *string {
	return s.MasterInstanceType
}

func (s *CreateClusterRequest) GetNodeCount() *int32 {
	return s.NodeCount
}

func (s *CreateClusterRequest) GetPayType() *string {
	return s.PayType
}

func (s *CreateClusterRequest) GetPeriod() *int32 {
	return s.Period
}

func (s *CreateClusterRequest) GetPeriodUnit() *string {
	return s.PeriodUnit
}

func (s *CreateClusterRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateClusterRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateClusterRequest) GetSecurityIPList() *string {
	return s.SecurityIPList
}

func (s *CreateClusterRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *CreateClusterRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *CreateClusterRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *CreateClusterRequest) SetAutoRenewPeriod(v int32) *CreateClusterRequest {
	s.AutoRenewPeriod = &v
	return s
}

func (s *CreateClusterRequest) SetClientToken(v string) *CreateClusterRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateClusterRequest) SetClusterName(v string) *CreateClusterRequest {
	s.ClusterName = &v
	return s
}

func (s *CreateClusterRequest) SetColdStorageSize(v int32) *CreateClusterRequest {
	s.ColdStorageSize = &v
	return s
}

func (s *CreateClusterRequest) SetCoreInstanceType(v string) *CreateClusterRequest {
	s.CoreInstanceType = &v
	return s
}

func (s *CreateClusterRequest) SetDiskSize(v int32) *CreateClusterRequest {
	s.DiskSize = &v
	return s
}

func (s *CreateClusterRequest) SetDiskType(v string) *CreateClusterRequest {
	s.DiskType = &v
	return s
}

func (s *CreateClusterRequest) SetEncryptionKey(v string) *CreateClusterRequest {
	s.EncryptionKey = &v
	return s
}

func (s *CreateClusterRequest) SetEngine(v string) *CreateClusterRequest {
	s.Engine = &v
	return s
}

func (s *CreateClusterRequest) SetEngineVersion(v string) *CreateClusterRequest {
	s.EngineVersion = &v
	return s
}

func (s *CreateClusterRequest) SetMasterInstanceType(v string) *CreateClusterRequest {
	s.MasterInstanceType = &v
	return s
}

func (s *CreateClusterRequest) SetNodeCount(v int32) *CreateClusterRequest {
	s.NodeCount = &v
	return s
}

func (s *CreateClusterRequest) SetPayType(v string) *CreateClusterRequest {
	s.PayType = &v
	return s
}

func (s *CreateClusterRequest) SetPeriod(v int32) *CreateClusterRequest {
	s.Period = &v
	return s
}

func (s *CreateClusterRequest) SetPeriodUnit(v string) *CreateClusterRequest {
	s.PeriodUnit = &v
	return s
}

func (s *CreateClusterRequest) SetRegionId(v string) *CreateClusterRequest {
	s.RegionId = &v
	return s
}

func (s *CreateClusterRequest) SetResourceGroupId(v string) *CreateClusterRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateClusterRequest) SetSecurityIPList(v string) *CreateClusterRequest {
	s.SecurityIPList = &v
	return s
}

func (s *CreateClusterRequest) SetVSwitchId(v string) *CreateClusterRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateClusterRequest) SetVpcId(v string) *CreateClusterRequest {
	s.VpcId = &v
	return s
}

func (s *CreateClusterRequest) SetZoneId(v string) *CreateClusterRequest {
	s.ZoneId = &v
	return s
}

func (s *CreateClusterRequest) Validate() error {
	return dara.Validate(s)
}
