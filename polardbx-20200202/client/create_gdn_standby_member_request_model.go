// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGdnStandbyMemberRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoRenew(v bool) *CreateGdnStandbyMemberRequest
	GetAutoRenew() *bool
	SetCNNodeCount(v string) *CreateGdnStandbyMemberRequest
	GetCNNodeCount() *string
	SetClientToken(v string) *CreateGdnStandbyMemberRequest
	GetClientToken() *string
	SetCloneInstanceName(v string) *CreateGdnStandbyMemberRequest
	GetCloneInstanceName() *string
	SetCnClass(v string) *CreateGdnStandbyMemberRequest
	GetCnClass() *string
	SetDNNodeCount(v string) *CreateGdnStandbyMemberRequest
	GetDNNodeCount() *string
	SetDescription(v string) *CreateGdnStandbyMemberRequest
	GetDescription() *string
	SetDnClass(v string) *CreateGdnStandbyMemberRequest
	GetDnClass() *string
	SetEngineVersion(v string) *CreateGdnStandbyMemberRequest
	GetEngineVersion() *string
	SetNetworkType(v string) *CreateGdnStandbyMemberRequest
	GetNetworkType() *string
	SetPayType(v string) *CreateGdnStandbyMemberRequest
	GetPayType() *string
	SetPeriod(v string) *CreateGdnStandbyMemberRequest
	GetPeriod() *string
	SetPrimaryZone(v string) *CreateGdnStandbyMemberRequest
	GetPrimaryZone() *string
	SetRegionId(v string) *CreateGdnStandbyMemberRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreateGdnStandbyMemberRequest
	GetResourceGroupId() *string
	SetSecondaryZone(v string) *CreateGdnStandbyMemberRequest
	GetSecondaryZone() *string
	SetSeries(v string) *CreateGdnStandbyMemberRequest
	GetSeries() *string
	SetSourceInstanceRegion(v string) *CreateGdnStandbyMemberRequest
	GetSourceInstanceRegion() *string
	SetStorageType(v string) *CreateGdnStandbyMemberRequest
	GetStorageType() *string
	SetTertiaryZone(v string) *CreateGdnStandbyMemberRequest
	GetTertiaryZone() *string
	SetTopologyType(v string) *CreateGdnStandbyMemberRequest
	GetTopologyType() *string
	SetUsedTime(v int32) *CreateGdnStandbyMemberRequest
	GetUsedTime() *int32
	SetVPCId(v string) *CreateGdnStandbyMemberRequest
	GetVPCId() *string
	SetVSwitchId(v string) *CreateGdnStandbyMemberRequest
	GetVSwitchId() *string
}

type CreateGdnStandbyMemberRequest struct {
	// example:
	//
	// true
	AutoRenew *bool `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// example:
	//
	// 2
	CNNodeCount *string `json:"CNNodeCount,omitempty" xml:"CNNodeCount,omitempty"`
	// example:
	//
	// FEA5DC20-6D8A-5979-97AA-FC57546ADC20
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pxc-***
	CloneInstanceName *string `json:"CloneInstanceName,omitempty" xml:"CloneInstanceName,omitempty"`
	CnClass           *string `json:"CnClass,omitempty" xml:"CnClass,omitempty"`
	// example:
	//
	// 2
	DNNodeCount *string `json:"DNNodeCount,omitempty" xml:"DNNodeCount,omitempty"`
	// example:
	//
	// go-to-the-docks-for-french-fries
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DnClass     *string `json:"DnClass,omitempty" xml:"DnClass,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 5.7
	EngineVersion *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	// example:
	//
	// vpc
	NetworkType *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// PREPAY
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// example:
	//
	// Month
	Period *string `json:"Period,omitempty" xml:"Period,omitempty"`
	// example:
	//
	// cn-shenzhen-e
	PrimaryZone *string `json:"PrimaryZone,omitempty" xml:"PrimaryZone,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// null
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// example:
	//
	// cn-shenzhen-a
	SecondaryZone *string `json:"SecondaryZone,omitempty" xml:"SecondaryZone,omitempty"`
	// example:
	//
	// enterprise
	Series *string `json:"Series,omitempty" xml:"Series,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-chengdu
	SourceInstanceRegion *string `json:"SourceInstanceRegion,omitempty" xml:"SourceInstanceRegion,omitempty"`
	// example:
	//
	// cloud_auto
	StorageType *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	// example:
	//
	// cn-shenzhen-e
	TertiaryZone *string `json:"TertiaryZone,omitempty" xml:"TertiaryZone,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 3azones
	TopologyType *string `json:"TopologyType,omitempty" xml:"TopologyType,omitempty"`
	// example:
	//
	// 1
	UsedTime *int32 `json:"UsedTime,omitempty" xml:"UsedTime,omitempty"`
	// VPC ID。
	//
	// example:
	//
	// vpc-*****
	VPCId *string `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	// example:
	//
	// vsw-*********
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
}

func (s CreateGdnStandbyMemberRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateGdnStandbyMemberRequest) GoString() string {
	return s.String()
}

func (s *CreateGdnStandbyMemberRequest) GetAutoRenew() *bool {
	return s.AutoRenew
}

func (s *CreateGdnStandbyMemberRequest) GetCNNodeCount() *string {
	return s.CNNodeCount
}

func (s *CreateGdnStandbyMemberRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateGdnStandbyMemberRequest) GetCloneInstanceName() *string {
	return s.CloneInstanceName
}

func (s *CreateGdnStandbyMemberRequest) GetCnClass() *string {
	return s.CnClass
}

func (s *CreateGdnStandbyMemberRequest) GetDNNodeCount() *string {
	return s.DNNodeCount
}

func (s *CreateGdnStandbyMemberRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateGdnStandbyMemberRequest) GetDnClass() *string {
	return s.DnClass
}

func (s *CreateGdnStandbyMemberRequest) GetEngineVersion() *string {
	return s.EngineVersion
}

func (s *CreateGdnStandbyMemberRequest) GetNetworkType() *string {
	return s.NetworkType
}

func (s *CreateGdnStandbyMemberRequest) GetPayType() *string {
	return s.PayType
}

func (s *CreateGdnStandbyMemberRequest) GetPeriod() *string {
	return s.Period
}

func (s *CreateGdnStandbyMemberRequest) GetPrimaryZone() *string {
	return s.PrimaryZone
}

func (s *CreateGdnStandbyMemberRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateGdnStandbyMemberRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateGdnStandbyMemberRequest) GetSecondaryZone() *string {
	return s.SecondaryZone
}

func (s *CreateGdnStandbyMemberRequest) GetSeries() *string {
	return s.Series
}

func (s *CreateGdnStandbyMemberRequest) GetSourceInstanceRegion() *string {
	return s.SourceInstanceRegion
}

func (s *CreateGdnStandbyMemberRequest) GetStorageType() *string {
	return s.StorageType
}

func (s *CreateGdnStandbyMemberRequest) GetTertiaryZone() *string {
	return s.TertiaryZone
}

func (s *CreateGdnStandbyMemberRequest) GetTopologyType() *string {
	return s.TopologyType
}

func (s *CreateGdnStandbyMemberRequest) GetUsedTime() *int32 {
	return s.UsedTime
}

func (s *CreateGdnStandbyMemberRequest) GetVPCId() *string {
	return s.VPCId
}

func (s *CreateGdnStandbyMemberRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *CreateGdnStandbyMemberRequest) SetAutoRenew(v bool) *CreateGdnStandbyMemberRequest {
	s.AutoRenew = &v
	return s
}

func (s *CreateGdnStandbyMemberRequest) SetCNNodeCount(v string) *CreateGdnStandbyMemberRequest {
	s.CNNodeCount = &v
	return s
}

func (s *CreateGdnStandbyMemberRequest) SetClientToken(v string) *CreateGdnStandbyMemberRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateGdnStandbyMemberRequest) SetCloneInstanceName(v string) *CreateGdnStandbyMemberRequest {
	s.CloneInstanceName = &v
	return s
}

func (s *CreateGdnStandbyMemberRequest) SetCnClass(v string) *CreateGdnStandbyMemberRequest {
	s.CnClass = &v
	return s
}

func (s *CreateGdnStandbyMemberRequest) SetDNNodeCount(v string) *CreateGdnStandbyMemberRequest {
	s.DNNodeCount = &v
	return s
}

func (s *CreateGdnStandbyMemberRequest) SetDescription(v string) *CreateGdnStandbyMemberRequest {
	s.Description = &v
	return s
}

func (s *CreateGdnStandbyMemberRequest) SetDnClass(v string) *CreateGdnStandbyMemberRequest {
	s.DnClass = &v
	return s
}

func (s *CreateGdnStandbyMemberRequest) SetEngineVersion(v string) *CreateGdnStandbyMemberRequest {
	s.EngineVersion = &v
	return s
}

func (s *CreateGdnStandbyMemberRequest) SetNetworkType(v string) *CreateGdnStandbyMemberRequest {
	s.NetworkType = &v
	return s
}

func (s *CreateGdnStandbyMemberRequest) SetPayType(v string) *CreateGdnStandbyMemberRequest {
	s.PayType = &v
	return s
}

func (s *CreateGdnStandbyMemberRequest) SetPeriod(v string) *CreateGdnStandbyMemberRequest {
	s.Period = &v
	return s
}

func (s *CreateGdnStandbyMemberRequest) SetPrimaryZone(v string) *CreateGdnStandbyMemberRequest {
	s.PrimaryZone = &v
	return s
}

func (s *CreateGdnStandbyMemberRequest) SetRegionId(v string) *CreateGdnStandbyMemberRequest {
	s.RegionId = &v
	return s
}

func (s *CreateGdnStandbyMemberRequest) SetResourceGroupId(v string) *CreateGdnStandbyMemberRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateGdnStandbyMemberRequest) SetSecondaryZone(v string) *CreateGdnStandbyMemberRequest {
	s.SecondaryZone = &v
	return s
}

func (s *CreateGdnStandbyMemberRequest) SetSeries(v string) *CreateGdnStandbyMemberRequest {
	s.Series = &v
	return s
}

func (s *CreateGdnStandbyMemberRequest) SetSourceInstanceRegion(v string) *CreateGdnStandbyMemberRequest {
	s.SourceInstanceRegion = &v
	return s
}

func (s *CreateGdnStandbyMemberRequest) SetStorageType(v string) *CreateGdnStandbyMemberRequest {
	s.StorageType = &v
	return s
}

func (s *CreateGdnStandbyMemberRequest) SetTertiaryZone(v string) *CreateGdnStandbyMemberRequest {
	s.TertiaryZone = &v
	return s
}

func (s *CreateGdnStandbyMemberRequest) SetTopologyType(v string) *CreateGdnStandbyMemberRequest {
	s.TopologyType = &v
	return s
}

func (s *CreateGdnStandbyMemberRequest) SetUsedTime(v int32) *CreateGdnStandbyMemberRequest {
	s.UsedTime = &v
	return s
}

func (s *CreateGdnStandbyMemberRequest) SetVPCId(v string) *CreateGdnStandbyMemberRequest {
	s.VPCId = &v
	return s
}

func (s *CreateGdnStandbyMemberRequest) SetVSwitchId(v string) *CreateGdnStandbyMemberRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateGdnStandbyMemberRequest) Validate() error {
	return dara.Validate(s)
}
