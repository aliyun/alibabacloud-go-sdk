// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateServerlessClusterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoRenewPeriod(v int32) *CreateServerlessClusterRequest
	GetAutoRenewPeriod() *int32
	SetClientToken(v string) *CreateServerlessClusterRequest
	GetClientToken() *string
	SetClientType(v string) *CreateServerlessClusterRequest
	GetClientType() *string
	SetClusterName(v string) *CreateServerlessClusterRequest
	GetClusterName() *string
	SetDiskType(v string) *CreateServerlessClusterRequest
	GetDiskType() *string
	SetEngine(v string) *CreateServerlessClusterRequest
	GetEngine() *string
	SetEngineVersion(v string) *CreateServerlessClusterRequest
	GetEngineVersion() *string
	SetPayType(v string) *CreateServerlessClusterRequest
	GetPayType() *string
	SetPeriod(v int32) *CreateServerlessClusterRequest
	GetPeriod() *int32
	SetPeriodUnit(v string) *CreateServerlessClusterRequest
	GetPeriodUnit() *string
	SetRegionId(v string) *CreateServerlessClusterRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreateServerlessClusterRequest
	GetResourceGroupId() *string
	SetServerlessCapability(v int32) *CreateServerlessClusterRequest
	GetServerlessCapability() *int32
	SetServerlessSpec(v string) *CreateServerlessClusterRequest
	GetServerlessSpec() *string
	SetServerlessStorage(v int32) *CreateServerlessClusterRequest
	GetServerlessStorage() *int32
	SetVSwitchId(v string) *CreateServerlessClusterRequest
	GetVSwitchId() *string
	SetVpcId(v string) *CreateServerlessClusterRequest
	GetVpcId() *string
	SetZoneId(v string) *CreateServerlessClusterRequest
	GetZoneId() *string
}

type CreateServerlessClusterRequest struct {
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
	// xx
	ClientType *string `json:"ClientType,omitempty" xml:"ClientType,omitempty"`
	// example:
	//
	// serverless-name
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// example:
	//
	// cloud_ssd
	DiskType *string `json:"DiskType,omitempty" xml:"DiskType,omitempty"`
	// example:
	//
	// serverlesshbase
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// example:
	//
	// 2.0
	EngineVersion *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
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
	// 1000
	ServerlessCapability *int32 `json:"ServerlessCapability,omitempty" xml:"ServerlessCapability,omitempty"`
	// example:
	//
	// serverless.small
	ServerlessSpec *string `json:"ServerlessSpec,omitempty" xml:"ServerlessSpec,omitempty"`
	// example:
	//
	// 100
	ServerlessStorage *int32 `json:"ServerlessStorage,omitempty" xml:"ServerlessStorage,omitempty"`
	// example:
	//
	// vsw-bp191ipotqj1ssyl*****
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

func (s CreateServerlessClusterRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateServerlessClusterRequest) GoString() string {
	return s.String()
}

func (s *CreateServerlessClusterRequest) GetAutoRenewPeriod() *int32 {
	return s.AutoRenewPeriod
}

func (s *CreateServerlessClusterRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateServerlessClusterRequest) GetClientType() *string {
	return s.ClientType
}

func (s *CreateServerlessClusterRequest) GetClusterName() *string {
	return s.ClusterName
}

func (s *CreateServerlessClusterRequest) GetDiskType() *string {
	return s.DiskType
}

func (s *CreateServerlessClusterRequest) GetEngine() *string {
	return s.Engine
}

func (s *CreateServerlessClusterRequest) GetEngineVersion() *string {
	return s.EngineVersion
}

func (s *CreateServerlessClusterRequest) GetPayType() *string {
	return s.PayType
}

func (s *CreateServerlessClusterRequest) GetPeriod() *int32 {
	return s.Period
}

func (s *CreateServerlessClusterRequest) GetPeriodUnit() *string {
	return s.PeriodUnit
}

func (s *CreateServerlessClusterRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateServerlessClusterRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateServerlessClusterRequest) GetServerlessCapability() *int32 {
	return s.ServerlessCapability
}

func (s *CreateServerlessClusterRequest) GetServerlessSpec() *string {
	return s.ServerlessSpec
}

func (s *CreateServerlessClusterRequest) GetServerlessStorage() *int32 {
	return s.ServerlessStorage
}

func (s *CreateServerlessClusterRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *CreateServerlessClusterRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *CreateServerlessClusterRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *CreateServerlessClusterRequest) SetAutoRenewPeriod(v int32) *CreateServerlessClusterRequest {
	s.AutoRenewPeriod = &v
	return s
}

func (s *CreateServerlessClusterRequest) SetClientToken(v string) *CreateServerlessClusterRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateServerlessClusterRequest) SetClientType(v string) *CreateServerlessClusterRequest {
	s.ClientType = &v
	return s
}

func (s *CreateServerlessClusterRequest) SetClusterName(v string) *CreateServerlessClusterRequest {
	s.ClusterName = &v
	return s
}

func (s *CreateServerlessClusterRequest) SetDiskType(v string) *CreateServerlessClusterRequest {
	s.DiskType = &v
	return s
}

func (s *CreateServerlessClusterRequest) SetEngine(v string) *CreateServerlessClusterRequest {
	s.Engine = &v
	return s
}

func (s *CreateServerlessClusterRequest) SetEngineVersion(v string) *CreateServerlessClusterRequest {
	s.EngineVersion = &v
	return s
}

func (s *CreateServerlessClusterRequest) SetPayType(v string) *CreateServerlessClusterRequest {
	s.PayType = &v
	return s
}

func (s *CreateServerlessClusterRequest) SetPeriod(v int32) *CreateServerlessClusterRequest {
	s.Period = &v
	return s
}

func (s *CreateServerlessClusterRequest) SetPeriodUnit(v string) *CreateServerlessClusterRequest {
	s.PeriodUnit = &v
	return s
}

func (s *CreateServerlessClusterRequest) SetRegionId(v string) *CreateServerlessClusterRequest {
	s.RegionId = &v
	return s
}

func (s *CreateServerlessClusterRequest) SetResourceGroupId(v string) *CreateServerlessClusterRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateServerlessClusterRequest) SetServerlessCapability(v int32) *CreateServerlessClusterRequest {
	s.ServerlessCapability = &v
	return s
}

func (s *CreateServerlessClusterRequest) SetServerlessSpec(v string) *CreateServerlessClusterRequest {
	s.ServerlessSpec = &v
	return s
}

func (s *CreateServerlessClusterRequest) SetServerlessStorage(v int32) *CreateServerlessClusterRequest {
	s.ServerlessStorage = &v
	return s
}

func (s *CreateServerlessClusterRequest) SetVSwitchId(v string) *CreateServerlessClusterRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateServerlessClusterRequest) SetVpcId(v string) *CreateServerlessClusterRequest {
	s.VpcId = &v
	return s
}

func (s *CreateServerlessClusterRequest) SetZoneId(v string) *CreateServerlessClusterRequest {
	s.ZoneId = &v
	return s
}

func (s *CreateServerlessClusterRequest) Validate() error {
	return dara.Validate(s)
}
