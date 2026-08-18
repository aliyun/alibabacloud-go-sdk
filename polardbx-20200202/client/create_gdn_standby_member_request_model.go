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
	// Specifies whether to enable auto-renewal. Default value: true.
	//
	// - **true**: Enabled.
	//
	// - **false**: Disabled.
	//
	// example:
	//
	// true
	AutoRenew *bool `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// The number of compute nodes.
	//
	// example:
	//
	// 2
	CNNodeCount *string `json:"CNNodeCount,omitempty" xml:"CNNodeCount,omitempty"`
	// The client token that is used to ensure the idempotence of the request. Make sure that the value is different for each request.
	//
	// example:
	//
	// FEA5DC20-6D8A-5979-97AA-FC57546ADC20
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The name of the source instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// pxc-***
	CloneInstanceName *string `json:"CloneInstanceName,omitempty" xml:"CloneInstanceName,omitempty"`
	// The compute node specifications. This parameter is required for Enterprise Edition instances and is not required for Standard Edition instances.
	//
	// Enterprise Edition with local disks:
	//
	// - **polarx.x4.medium.2e**: 2 cores, 8 GB (general-purpose)
	//
	// - **polarx.x4.large.2e**: 4 cores, 16 GB (general-purpose)
	//
	// - **polarx.x4.xlarge.2e**: 8 cores, 32 GB (general-purpose)
	//
	// - **polarx.x4.2xlarge.2e**: 16 cores, 64 GB (general-purpose)
	//
	// - **polarx.x8.large.2e**: 4 cores, 32 GB (dedicated)
	//
	// - **polarx.x2.large.2x**: 8 cores, 16 GB (dedicated)
	//
	// - **polarx.x4.xlarge.2x**: 8 cores, 32 GB (dedicated)
	//
	// - **polarx.x8.xlarge.2e**: 8 cores, 64 GB (dedicated)
	//
	// - **polarx.x8.2xlarge.2e**: 16 cores, 128 GB (dedicated)
	//
	// - **polarx.x4.4xlarge.2e**: 32 cores, 128 GB (dedicated)
	//
	// - **polarx.x8.4xlarge.2e**: 32 cores, 256 GB (dedicated)
	//
	// - **polarx.st.8xlarge.2e**: 60 cores, 470 GB (dedicated)
	//
	// - **polarx.st.12xlarge.2e**: 90 cores, 720 GB (dedicated)
	//
	//
	// Enterprise Edition with cloud disks:
	//
	// - **polarx.x4.medium.c2e**: 2 cores, 8 GB (general-purpose)
	//
	// - **polarx.x4.large.c2e**: 4 cores, 16 GB (general-purpose)
	//
	// - **polarx.x4.xlarge.c2e**: 8 cores, 32 GB (general-purpose)
	//
	// - **polarx.x4.2xlarge.c2e**: 16 cores, 64 GB (general-purpose)
	//
	// - **polarx.x8.large.c2e**: 4 cores, 32 GB (dedicated)
	//
	// - **polarx.x2.large.c2x**: 8 cores, 16 GB (dedicated)
	//
	// - **polarx.x4.xlarge.c2x**: 8 cores, 32 GB (dedicated)
	//
	// - **polarx.x8.xlarge.c2e**: 8 cores, 64 GB (dedicated)
	//
	// - **polarx.x8.2xlarge.c2e**: 16 cores, 128 GB (dedicated)
	//
	// - **polarx.x4.4xlarge.c2e**: 32 cores, 128 GB (dedicated)
	//
	// - **polarx.x8.4xlarge.c2e**: 32 cores, 256 GB (dedicated)
	//
	// - **polarx.st.8xlarge.c2e**: 60 cores, 470 GB (dedicated)
	//
	// - **polarx.st.12xlarge.c2e**: 90 cores, 720 GB (dedicated)
	//
	// example:
	//
	// 4 cores, 32 GB
	CnClass *string `json:"CnClass,omitempty" xml:"CnClass,omitempty"`
	// The number of storage nodes.
	//
	// example:
	//
	// 2
	DNNodeCount *string `json:"DNNodeCount,omitempty" xml:"DNNodeCount,omitempty"`
	// The description of the instance.
	//
	// example:
	//
	// go-to-the-docks-for-french-fries
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The storage node specifications. This parameter is required for Enterprise Edition instances and is not required for Standard Edition instances.
	//
	// Enterprise Edition with local disks:
	//
	// - **mysql.n2.medium.25**: 2 cores, 4 GB (general-purpose)
	//
	// - **mysql.n4.medium.25**: 2 cores, 8 GB (general-purpose)
	//
	// - **mysql.n2.large.25**: 4 cores, 8 GB (general-purpose)
	//
	// - **mysql.n4.large.25**: 4 cores, 16 GB (general-purpose)
	//
	// - **mysql.n4.xlarge.25**: 8 cores, 32 GB (general-purpose)
	//
	// - **mysql.n4.2xlarge.25**: 16 cores, 64 GB (general-purpose)
	//
	// - **mysql.x4.large.25**: 4 cores, 16 GB (dedicated)
	//
	// - **mysql.x8.large.25**: 4 cores, 32 GB (dedicated)
	//
	// - **mysql.x2.xlarge.25**: 8 cores, 16 GB (dedicated)
	//
	// - **mysql.x8.xlarge.25**: 8 cores, 64 GB (dedicated)
	//
	// - **mysql.x8.2xlarge.25**: 16 cores, 128 GB (dedicated)
	//
	// - **mysql.x4.4xlarge.25**: 32 cores, 128 GB (dedicated)
	//
	// - **mysql.x8.4xlarge.25**: 32 cores, 256 GB (dedicated)
	//
	// - **mysql.st.8xlarge.25**: 60 cores, 470 GB (dedicated)
	//
	// - **mysql.st.12xlarge.25**: 90 cores, 720 GB (dedicated)
	//
	// - **mysql.x8.45xlarge.25**: 180 cores, 1440 GB (dedicated)
	//
	// - **mysql.x8.60xlarge.25**: 240 cores, 1920 GB (dedicated)
	//
	//
	// Enterprise Edition with cloud disks:
	//
	// - **polarx.mysql.n2.medium.c25**: 2 cores, 4 GB (general-purpose)
	//
	// - **polarx.mysql.n4.medium.c25**: 2 cores, 8 GB (general-purpose)
	//
	// - **polarx.mysql.n2.large.c25**: 4 cores, 8 GB (general-purpose)
	//
	// - **polarx.mysql.n4.large.c25**: 4 cores, 16 GB (general-purpose)
	//
	// - **polarx.mysql.n4.xlarge.c25**: 8 cores, 32 GB (general-purpose)
	//
	// - **polarx.mysql.n4.2xlarge.c25**: 16 cores, 64 GB (general-purpose)
	//
	// - **polarx.mysql.x4.large.c25**: 4 cores, 16 GB (dedicated)
	//
	// - **polarx.mysql.x8.large.c25**: 4 cores, 32 GB (dedicated)
	//
	// - **polarx.mysql.x2.xlarge.c25**: 8 cores, 16 GB (dedicated)
	//
	// - **polarx.mysql.x8.xlarge.c25**: 8 cores, 64 GB (dedicated)
	//
	// - **polarx.mysql.x8.2xlarge.c25**: 16 cores, 128 GB (dedicated)
	//
	// - **polarx.mysql.x4.4xlarge.c25**: 32 cores, 128 GB (dedicated)
	//
	// - **polarx.mysql.x8.4xlarge.c25**: 32 cores, 256 GB (dedicated)
	//
	// - **polarx.mysql.st.8xlarge.c25**: 60 cores, 470 GB (dedicated)
	//
	// - **polarx.mysql.st.12xlarge.c25**: 90 cores, 720 GB (dedicated)
	//
	// - **polarx.mysql.x8.45xlarge.c25**: 180 cores, 1440 GB (dedicated)
	//
	// - **polarx.mysql.x8.60xlarge.c25**: 240 cores, 1920 GB (dedicated)
	//
	// example:
	//
	// 4 cores, 32 GB
	DnClass *string `json:"DnClass,omitempty" xml:"DnClass,omitempty"`
	// The MySQL DPI engine version. Valid values: 5.7 and 8.0.
	//
	// This parameter is required.
	//
	// example:
	//
	// 5.7
	EngineVersion *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	// The network type. Only VPC is supported.
	//
	// example:
	//
	// vpc
	NetworkType *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	// The billing method of the instance.
	//
	// - **PREPAY**: subscription
	//
	// - **POSTPAY**: pay-as-you-go
	//
	// This parameter is required.
	//
	// example:
	//
	// PREPAY
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// The unit of the subscription duration. Valid values:
	//
	// - Year
	//
	// - Month
	//
	// For pay-as-you-go instances, the default value is Hour.
	//
	// example:
	//
	// Month
	Period *string `json:"Period,omitempty" xml:"Period,omitempty"`
	// The primary zone.
	//
	// example:
	//
	// cn-shenzhen-e
	PrimaryZone *string `json:"PrimaryZone,omitempty" xml:"PrimaryZone,omitempty"`
	// The region in which the instance resides.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID. This parameter can be left empty. This parameter is not supported.
	//
	// example:
	//
	// null
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The secondary zone.
	//
	// example:
	//
	// cn-shenzhen-a
	SecondaryZone *string `json:"SecondaryZone,omitempty" xml:"SecondaryZone,omitempty"`
	// The edition of the instance. Valid values:
	//
	// - enterprise: Enterprise Edition.
	//
	// - standard: Standard Edition.
	//
	// example:
	//
	// enterprise
	Series *string `json:"Series,omitempty" xml:"Series,omitempty"`
	// The region in which the source instance resides.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-chengdu
	SourceInstanceRegion *string `json:"SourceInstanceRegion,omitempty" xml:"SourceInstanceRegion,omitempty"`
	// The storage type.
	//
	// example:
	//
	// cloud_auto
	StorageType *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	// The zone for Three-zone deployment.
	//
	// example:
	//
	// cn-shenzhen-e
	TertiaryZone *string `json:"TertiaryZone,omitempty" xml:"TertiaryZone,omitempty"`
	// The topology type. Valid values:
	//
	// - **3azones**: three-zone deployment.
	//
	// - **1azone**: single-zone deployment.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3azones
	TopologyType *string `json:"TopologyType,omitempty" xml:"TopologyType,omitempty"`
	// The subscription duration. Unit: months or years.
	//
	// > If Period is set to Year, valid values of this parameter are 1, 2, and 3.
	//
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
	// The vSwitch ID.
	//
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
