// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLindormInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetArbiterVSwitchId(v string) *CreateLindormInstanceRequest
	GetArbiterVSwitchId() *string
	SetArbiterZoneId(v string) *CreateLindormInstanceRequest
	GetArbiterZoneId() *string
	SetArchVersion(v string) *CreateLindormInstanceRequest
	GetArchVersion() *string
	SetAutoRenewDuration(v string) *CreateLindormInstanceRequest
	GetAutoRenewDuration() *string
	SetAutoRenewal(v bool) *CreateLindormInstanceRequest
	GetAutoRenewal() *bool
	SetColdStorage(v int32) *CreateLindormInstanceRequest
	GetColdStorage() *int32
	SetCoreSingleStorage(v int32) *CreateLindormInstanceRequest
	GetCoreSingleStorage() *int32
	SetCoreSpec(v string) *CreateLindormInstanceRequest
	GetCoreSpec() *string
	SetDiskCategory(v string) *CreateLindormInstanceRequest
	GetDiskCategory() *string
	SetDuration(v string) *CreateLindormInstanceRequest
	GetDuration() *string
	SetFilestoreNum(v int32) *CreateLindormInstanceRequest
	GetFilestoreNum() *int32
	SetFilestoreSpec(v string) *CreateLindormInstanceRequest
	GetFilestoreSpec() *string
	SetInstanceAlias(v string) *CreateLindormInstanceRequest
	GetInstanceAlias() *string
	SetInstanceStorage(v string) *CreateLindormInstanceRequest
	GetInstanceStorage() *string
	SetLindormNum(v int32) *CreateLindormInstanceRequest
	GetLindormNum() *int32
	SetLindormSpec(v string) *CreateLindormInstanceRequest
	GetLindormSpec() *string
	SetLogDiskCategory(v string) *CreateLindormInstanceRequest
	GetLogDiskCategory() *string
	SetLogNum(v int32) *CreateLindormInstanceRequest
	GetLogNum() *int32
	SetLogSingleStorage(v int32) *CreateLindormInstanceRequest
	GetLogSingleStorage() *int32
	SetLogSpec(v string) *CreateLindormInstanceRequest
	GetLogSpec() *string
	SetLtsNum(v string) *CreateLindormInstanceRequest
	GetLtsNum() *string
	SetLtsSpec(v string) *CreateLindormInstanceRequest
	GetLtsSpec() *string
	SetMultiZoneCombination(v string) *CreateLindormInstanceRequest
	GetMultiZoneCombination() *string
	SetOwnerAccount(v string) *CreateLindormInstanceRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateLindormInstanceRequest
	GetOwnerId() *int64
	SetPayType(v string) *CreateLindormInstanceRequest
	GetPayType() *string
	SetPricingCycle(v string) *CreateLindormInstanceRequest
	GetPricingCycle() *string
	SetPrimaryVSwitchId(v string) *CreateLindormInstanceRequest
	GetPrimaryVSwitchId() *string
	SetPrimaryZoneId(v string) *CreateLindormInstanceRequest
	GetPrimaryZoneId() *string
	SetRegionId(v string) *CreateLindormInstanceRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreateLindormInstanceRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *CreateLindormInstanceRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateLindormInstanceRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *CreateLindormInstanceRequest
	GetSecurityToken() *string
	SetSolrNum(v int32) *CreateLindormInstanceRequest
	GetSolrNum() *int32
	SetSolrSpec(v string) *CreateLindormInstanceRequest
	GetSolrSpec() *string
	SetStandbyVSwitchId(v string) *CreateLindormInstanceRequest
	GetStandbyVSwitchId() *string
	SetStandbyZoneId(v string) *CreateLindormInstanceRequest
	GetStandbyZoneId() *string
	SetStreamNum(v int32) *CreateLindormInstanceRequest
	GetStreamNum() *int32
	SetStreamSpec(v string) *CreateLindormInstanceRequest
	GetStreamSpec() *string
	SetTag(v []*CreateLindormInstanceRequestTag) *CreateLindormInstanceRequest
	GetTag() []*CreateLindormInstanceRequestTag
	SetTsdbNum(v int32) *CreateLindormInstanceRequest
	GetTsdbNum() *int32
	SetTsdbSpec(v string) *CreateLindormInstanceRequest
	GetTsdbSpec() *string
	SetVPCId(v string) *CreateLindormInstanceRequest
	GetVPCId() *string
	SetVSwitchId(v string) *CreateLindormInstanceRequest
	GetVSwitchId() *string
	SetZoneId(v string) *CreateLindormInstanceRequest
	GetZoneId() *string
}

type CreateLindormInstanceRequest struct {
	// The ID of the VSwitch for the arbiter zone of the multi-zone instance. The VSwitch must be in the zone specified by `ArbiterZoneId`. **This parameter is required for multi-zone instances.**
	//
	// example:
	//
	// vsw-uf6664pqjawb87k36****
	ArbiterVSwitchId *string `json:"ArbiterVSwitchId,omitempty" xml:"ArbiterVSwitchId,omitempty"`
	// The ID of the arbiter zone for the multi-zone instance. **This parameter is required for multi-zone instances.**
	//
	// example:
	//
	// cn-shanghai-g
	ArbiterZoneId *string `json:"ArbiterZoneId,omitempty" xml:"ArbiterZoneId,omitempty"`
	// The deployment architecture of the instance. Valid values:
	//
	// - **1.0**: Single-zone deployment.
	//
	// - **2.0**: Multi-zone deployment.
	//
	// The default value is 1.0. To create a multi-zone instance, set this parameter to 2.0. **This parameter is required for multi-zone instances.**
	//
	// example:
	//
	// 2.0
	ArchVersion *string `json:"ArchVersion,omitempty" xml:"ArchVersion,omitempty"`
	// The auto-renewal duration, in months.
	//
	// The value of this parameter ranges from **1*	- to **12**.
	//
	// > This parameter takes effect only when **AutoRenewal*	- is set to **true**.
	//
	// example:
	//
	// 1
	AutoRenewDuration *string `json:"AutoRenewDuration,omitempty" xml:"AutoRenewDuration,omitempty"`
	// Specifies whether to enable auto-renewal for the Subscription instance. Valid values:
	//
	// - **true**: Auto-renewal is enabled.
	//
	// - **false**: Auto-renewal is disabled.
	//
	// Default value: false.
	//
	// > This parameter takes effect only when the **PayType*	- parameter is set to **PREPAY**.
	//
	// example:
	//
	// false
	AutoRenewal *bool `json:"AutoRenewal,omitempty" xml:"AutoRenewal,omitempty"`
	// The cold storage capacity of the instance, in GB. The value of this parameter ranges from **800*	- to **1,000,000**. If you do not specify this parameter, cold storage is not enabled.
	//
	// example:
	//
	// 800
	ColdStorage *int32 `json:"ColdStorage,omitempty" xml:"ColdStorage,omitempty"`
	// The storage capacity of a single core node in the multi-zone instance. Unit: GB. The value of this parameter ranges from 400 to 64,000. **This parameter is required for multi-zone instances.**
	//
	// example:
	//
	// 400
	CoreSingleStorage *int32 `json:"CoreSingleStorage,omitempty" xml:"CoreSingleStorage,omitempty"`
	// The node specification for an instance that uses local disks.
	//
	// If the storage type is **local_ssd_pro**, valid values include the following: Note that I3-family specifications are available only for Subscription instances.
	//
	// - **lindorm.i4.xlarge**: 4 cores, 32 GB memory (I4).
	//
	// - **lindorm.i4.2xlarge**: 8 cores, 64 GB memory (I4).
	//
	// - **lindorm.i4.4xlarge**: 16 cores, 128 GB memory (I4).
	//
	// - **lindorm.i4.8xlarge**: 32 cores, 256 GB memory (I4).
	//
	// - **lindorm.i3.xlarge**: 4 cores, 32 GB memory (I3).
	//
	// - **lindorm.i3.2xlarge**: 8 cores, 64 GB memory (I3).
	//
	// - **lindorm.i3.4xlarge**: 16 cores, 128 GB memory (I3).
	//
	// - **lindorm.i3.8xlarge**: 32 cores, 256 GB memory (I3).
	//
	// - **lindorm.i2.xlarge**: 4 cores, 32 GB memory (I2).
	//
	// - **lindorm.i2.2xlarge**: 8 cores, 64 GB memory (I2).
	//
	// - **lindorm.i2.4xlarge**: 16 cores, 128 GB memory (I2).
	//
	// - **lindorm.i2.8xlarge**: 32 cores, 256 GB memory (I2).
	//
	// If the storage type is **local_hdd_pro**, valid values include:
	//
	// - **lindorm.sd3c.3xlarge**: 14 cores, 56 GB memory (D3C PRO).
	//
	// - **lindorm.sd3c.7xlarge**: 28 cores, 112 GB memory (D3C PRO).
	//
	// - **lindorm.sd3c.14xlarge**: 56 cores, 224 GB memory (D3C PRO).
	//
	// - **lindorm.d2c.6xlarge**: 24 cores, 88 GB memory (D2C).
	//
	// - **lindorm.d2c.12xlarge**: 48 cores, 176 GB memory (D2C).
	//
	// - **lindorm.d2c.24xlarge**: 96 cores, 352 GB memory (D2C).
	//
	// - **lindorm.d2s.5xlarge**: 20 cores, 88 GB memory (D2S).
	//
	// - **lindorm.d2s.10xlarge**: 40 cores, 176 GB memory (D2S).
	//
	// - **lindorm.d1.2xlarge**: 8 cores, 32 GB memory (D1NE).
	//
	// - **lindorm.d1.4xlarge**: 16 cores, 64 GB memory (D1NE).
	//
	// - **lindorm.d1.6xlarge**: 24 cores, 96 GB memory (D1NE).
	//
	// example:
	//
	// lindorm.i2.xlarge
	CoreSpec *string `json:"CoreSpec,omitempty" xml:"CoreSpec,omitempty"`
	// The storage type of the instance. Valid values:
	//
	// - **cloud_efficiency**: Efficiency cloud disk.
	//
	// - **cloud_ssd**: Performance cloud disk.
	//
	// - **cloud_essd**: Enhanced SSD (ESSD).
	//
	// - **cloud_essd_pl0**: ESSD PL0.
	//
	// - **capacity_cloud_storage**: Capacity-optimized cloud storage. (Not available for multi-zone instances.)
	//
	// - **local_ssd_pro**: Local SSD. (Not available for multi-zone instances.)
	//
	// - **local_hdd_pro**: Local HDD. (Not available for multi-zone instances.)
	//
	// This parameter is required.
	//
	// example:
	//
	// cloud_efficiency
	DiskCategory *string `json:"DiskCategory,omitempty" xml:"DiskCategory,omitempty"`
	// The subscription duration for the instance. Valid values:
	//
	// - If **PricingCycle*	- is set to **Month**, the value can range from **1*	- to **9**.
	//
	// - If **PricingCycle*	- is set to **Year**, the value can range from **1*	- to **3**.
	//
	// > This parameter is required if you set **PayType*	- to **PREPAY**.
	//
	// example:
	//
	// 1
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The number of nodes in the file engine. Valid values:
	//
	// - For a Subscription instance, the value of this parameter ranges from **0*	- to **60**.
	//
	// - For a Pay-As-You-Go instance, the value of this parameter ranges from **0*	- to **8**.
	//
	// example:
	//
	// 2
	FilestoreNum *int32 `json:"FilestoreNum,omitempty" xml:"FilestoreNum,omitempty"`
	// The specification of the file engine nodes. Valid values:
	//
	// - **lindorm.c.xlarge**: 4 cores, 8 GB memory (standard).
	//
	// example:
	//
	// lindorm.c.xlarge
	FilestoreSpec *string `json:"FilestoreSpec,omitempty" xml:"FilestoreSpec,omitempty"`
	// The name of the instance.
	//
	// example:
	//
	// lindorm_test
	InstanceAlias *string `json:"InstanceAlias,omitempty" xml:"InstanceAlias,omitempty"`
	// The storage capacity of the instance, in GB.
	//
	// example:
	//
	// 480
	InstanceStorage *string `json:"InstanceStorage,omitempty" xml:"InstanceStorage,omitempty"`
	// The number of nodes in the wide table engine.
	//
	// For a single-zone instance, the value of this parameter ranges from **0*	- to **90**.
	//
	// **This parameter is required for multi-zone instances.*	- For an instance that uses cloud disks, the value ranges from **4*	- to **400**. For an instance that uses local disks, the value ranges from **6*	- to **400**.
	//
	// example:
	//
	// 2
	LindormNum *int32 `json:"LindormNum,omitempty" xml:"LindormNum,omitempty"`
	// The specification of the wide table engine nodes. Valid values:
	//
	// - **lindorm.g.xlarge**: 4 cores, 16 GB memory (dedicated).
	//
	// - **lindorm.c.2xlarge**: 8 cores, 16 GB memory (dedicated).
	//
	// - **lindorm.g.2xlarge**: 8 cores, 32 GB memory (dedicated).
	//
	// - **lindorm.c.4xlarge**: 16 cores, 32 GB memory (dedicated).
	//
	// - **lindorm.g.4xlarge**: 16 cores, 64 GB memory (dedicated).
	//
	// - **lindorm.c.8xlarge**: 32 cores, 64 GB memory (dedicated).
	//
	// - **lindorm.g.8xlarge**: 32 cores, 128 GB memory (dedicated).
	//
	// example:
	//
	// lindorm.c.xlarge
	LindormSpec *string `json:"LindormSpec,omitempty" xml:"LindormSpec,omitempty"`
	// The storage type of the log nodes for the multi-zone instance. Valid values:
	//
	// - **cloud_efficiency**: Efficiency cloud disk.
	//
	// - **cloud_ssd**: Performance cloud disk.
	//
	// **This parameter is required for multi-zone instances.**
	//
	// example:
	//
	// cloud_ssd
	LogDiskCategory *string `json:"LogDiskCategory,omitempty" xml:"LogDiskCategory,omitempty"`
	// The number of log nodes for the multi-zone instance. The value of this parameter ranges from 4 to 400. **This parameter is required for multi-zone instances.**
	//
	// example:
	//
	// 4
	LogNum *int32 `json:"LogNum,omitempty" xml:"LogNum,omitempty"`
	// The storage capacity of a single log node in the multi-zone instance. Unit: GB. The value of this parameter ranges from 400 to 64,000. **This parameter is required for multi-zone instances.**
	//
	// example:
	//
	// 400
	LogSingleStorage *int32 `json:"LogSingleStorage,omitempty" xml:"LogSingleStorage,omitempty"`
	// The specification of the log nodes for the multi-zone instance. Valid values:
	//
	// - **lindorm.sn1.large**: 4 cores, 8 GB memory (dedicated).
	//
	// - **lindorm.sn1.2xlarge**: 8 cores, 16 GB memory (dedicated).
	//
	// **This parameter is required for multi-zone instances.**
	//
	// example:
	//
	// lindorm.sn1.large
	LogSpec *string `json:"LogSpec,omitempty" xml:"LogSpec,omitempty"`
	// The number of nodes in the LTS engine. The value of this parameter ranges from **0*	- to **60**.
	//
	// example:
	//
	// 2
	LtsNum *string `json:"LtsNum,omitempty" xml:"LtsNum,omitempty"`
	// The specification of the LTS engine nodes. Valid values:
	//
	// - **lindorm.c.xlarge**: 4 cores, 8 GB memory (dedicated).
	//
	// - **lindorm.g.xlarge**: 4 cores, 16 GB memory (dedicated).
	//
	// - **lindorm.c.2xlarge**: 8 cores, 16 GB memory (dedicated).
	//
	// - **lindorm.g.2xlarge**: 8 cores, 32 GB memory (dedicated).
	//
	// - **lindorm.c.4xlarge**: 16 cores, 32 GB memory (dedicated).
	//
	// - **lindorm.g.4xlarge**: 16 cores, 64 GB memory (dedicated).
	//
	// - **lindorm.c.8xlarge**: 32 cores, 64 GB memory (dedicated).
	//
	// - **lindorm.g.8xlarge**: 32 cores, 128 GB memory (dedicated).
	//
	// example:
	//
	// lindorm.g.xlarge
	LtsSpec *string `json:"LtsSpec,omitempty" xml:"LtsSpec,omitempty"`
	// The combination of zones for the multi-zone instance. For a list of supported combinations, refer to the instance purchase page.
	//
	// - **ap-southeast-5abc-aliyun**: Indonesia (Jakarta) A+B+C.
	//
	// - **cn-hangzhou-ehi-aliyun**: China (Hangzhou) E+H+I.
	//
	// - **cn-beijing-acd-aliyun**: China (Beijing) A+C+D.
	//
	// - **ap-southeast-1-abc-aliyun**: Singapore A+B+C.
	//
	// - **cn-zhangjiakou-abc-aliyun**: China (Zhangjiakou) A+B+C.
	//
	// - **cn-shanghai-efg-aliyun**: China (Shanghai) E+F+G.
	//
	// - **cn-shanghai-abd-aliyun**: China (Shanghai) A+B+D.
	//
	// - **cn-hangzhou-bef-aliyun**: China (Hangzhou) B+E+F.
	//
	// - **cn-hangzhou-bce-aliyun**: China (Hangzhou) B+C+E.
	//
	// - **cn-beijing-fgh-aliyun**: China (Beijing) F+G+H.
	//
	// - **cn-shenzhen-abc-aliyun**: China (Shenzhen) A+B+C.
	//
	// **This parameter is required for multi-zone instances.**
	//
	// example:
	//
	// cn-shanghai-efg-aliyun
	MultiZoneCombination *string `json:"MultiZoneCombination,omitempty" xml:"MultiZoneCombination,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The billing method of the instance. Valid values:
	//
	// - **PREPAY**: Subscription.
	//
	// - **POSTPAY**: Pay-As-You-Go.
	//
	// This parameter is required.
	//
	// example:
	//
	// POSTPAY
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// The billing cycle for the Subscription instance. Valid values:
	//
	// - **Month**
	//
	// - **Year**
	//
	// > This parameter is required if you set **PayType*	- to **PREPAY**.
	//
	// example:
	//
	// Month
	PricingCycle *string `json:"PricingCycle,omitempty" xml:"PricingCycle,omitempty"`
	// The ID of the VSwitch for the primary zone of the multi-zone instance. The VSwitch must be in the zone specified by `PrimaryZoneId`. **This parameter is required for multi-zone instances.**
	//
	// example:
	//
	// vsw-uf6fdqa7c0pipnqzq****
	PrimaryVSwitchId *string `json:"PrimaryVSwitchId,omitempty" xml:"PrimaryVSwitchId,omitempty"`
	// The ID of the primary zone for the multi-zone instance. **This parameter is required for multi-zone instances.**
	//
	// example:
	//
	// cn-shanghai-e
	PrimaryZoneId *string `json:"PrimaryZoneId,omitempty" xml:"PrimaryZoneId,omitempty"`
	// The ID of the region in which to create the instance. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/426062.html) operation to query the latest region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-aek2i6weeb4nfii
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The number of search engine nodes. The value of this parameter ranges from **0*	- to **60**.
	//
	// example:
	//
	// 2
	SolrNum *int32 `json:"SolrNum,omitempty" xml:"SolrNum,omitempty"`
	// The specification of the search engine nodes. Valid values:
	//
	// - **lindorm.g.xlarge**: 4 cores, 16 GB memory (dedicated).
	//
	// - **lindorm.c.2xlarge**: 8 cores, 16 GB memory (dedicated).
	//
	// - **lindorm.g.2xlarge**: 8 cores, 32 GB memory (dedicated).
	//
	// - **lindorm.c.4xlarge**: 16 cores, 32 GB memory (dedicated).
	//
	// - **lindorm.g.4xlarge**: 16 cores, 64 GB memory (dedicated).
	//
	// - **lindorm.c.8xlarge**: 32 cores, 64 GB memory (dedicated).
	//
	// - **lindorm.g.8xlarge**: 32 cores, 128 GB memory (dedicated).
	//
	// example:
	//
	// lindorm.g.xlarge
	SolrSpec *string `json:"SolrSpec,omitempty" xml:"SolrSpec,omitempty"`
	// The ID of the VSwitch for the standby zone of the multi-zone instance. The VSwitch must be in the zone specified by `StandbyZoneId`. **This parameter is required for multi-zone instances.**
	//
	// example:
	//
	// vsw-2zec0kcn08cgdtr6****
	StandbyVSwitchId *string `json:"StandbyVSwitchId,omitempty" xml:"StandbyVSwitchId,omitempty"`
	// The ID of the standby zone for the multi-zone instance. **This parameter is required for multi-zone instances.**
	//
	// example:
	//
	// cn-shanghai-f
	StandbyZoneId *string `json:"StandbyZoneId,omitempty" xml:"StandbyZoneId,omitempty"`
	// The number of nodes in the stream engine. The value of this parameter ranges from **0*	- to **60**.
	//
	// example:
	//
	// 2
	StreamNum *int32 `json:"StreamNum,omitempty" xml:"StreamNum,omitempty"`
	// The specification of the stream engine nodes. Valid values:
	//
	// - **lindorm.g.xlarge**: 4 cores, 16 GB memory (dedicated).
	//
	// - **lindorm.c.2xlarge**: 8 cores, 16 GB memory (dedicated).
	//
	// - **lindorm.g.2xlarge**: 8 cores, 32 GB memory (dedicated).
	//
	// - **lindorm.c.4xlarge**: 16 cores, 32 GB memory (dedicated).
	//
	// - **lindorm.g.4xlarge**: 16 cores, 64 GB memory (dedicated).
	//
	// - **lindorm.c.8xlarge**: 32 cores, 64 GB memory (dedicated).
	//
	// - **lindorm.g.8xlarge**: 32 cores, 128 GB memory (dedicated).
	//
	// example:
	//
	// lindorm.g.xlarge
	StreamSpec *string `json:"StreamSpec,omitempty" xml:"StreamSpec,omitempty"`
	// The tags to add to the instance. You can add up to 20 tags.
	Tag []*CreateLindormInstanceRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The number of nodes in the time series engine. Valid values:
	//
	// - For a Subscription instance, the value of this parameter ranges from **0*	- to **24**.
	//
	// - For a Pay-As-You-Go instance, the value of this parameter ranges from **0*	- to **32**.
	//
	// example:
	//
	// 2
	TsdbNum *int32 `json:"TsdbNum,omitempty" xml:"TsdbNum,omitempty"`
	// The specification of the time series engine nodes. Valid values:
	//
	// - **lindorm.g.xlarge**: 4 cores, 16 GB memory (dedicated).
	//
	// - **lindorm.g.2xlarge**: 8 cores, 32 GB memory (dedicated).
	//
	// - **lindorm.g.4xlarge**: 16 cores, 64 GB memory (dedicated).
	//
	// - **lindorm.g.8xlarge**: 32 cores, 128 GB memory (dedicated).
	//
	// example:
	//
	// lindorm.g.xlarge
	TsdbSpec *string `json:"TsdbSpec,omitempty" xml:"TsdbSpec,omitempty"`
	// The ID of the VPC where you want to create the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpc-bp1nme44gek34slfc****
	VPCId *string `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	// The ID of the VSwitch.
	//
	// This parameter is required.
	//
	// example:
	//
	// vsw-bp1e7clcw529l773d****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The ID of the zone where you want to create the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai-f
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s CreateLindormInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateLindormInstanceRequest) GoString() string {
	return s.String()
}

func (s *CreateLindormInstanceRequest) GetArbiterVSwitchId() *string {
	return s.ArbiterVSwitchId
}

func (s *CreateLindormInstanceRequest) GetArbiterZoneId() *string {
	return s.ArbiterZoneId
}

func (s *CreateLindormInstanceRequest) GetArchVersion() *string {
	return s.ArchVersion
}

func (s *CreateLindormInstanceRequest) GetAutoRenewDuration() *string {
	return s.AutoRenewDuration
}

func (s *CreateLindormInstanceRequest) GetAutoRenewal() *bool {
	return s.AutoRenewal
}

func (s *CreateLindormInstanceRequest) GetColdStorage() *int32 {
	return s.ColdStorage
}

func (s *CreateLindormInstanceRequest) GetCoreSingleStorage() *int32 {
	return s.CoreSingleStorage
}

func (s *CreateLindormInstanceRequest) GetCoreSpec() *string {
	return s.CoreSpec
}

func (s *CreateLindormInstanceRequest) GetDiskCategory() *string {
	return s.DiskCategory
}

func (s *CreateLindormInstanceRequest) GetDuration() *string {
	return s.Duration
}

func (s *CreateLindormInstanceRequest) GetFilestoreNum() *int32 {
	return s.FilestoreNum
}

func (s *CreateLindormInstanceRequest) GetFilestoreSpec() *string {
	return s.FilestoreSpec
}

func (s *CreateLindormInstanceRequest) GetInstanceAlias() *string {
	return s.InstanceAlias
}

func (s *CreateLindormInstanceRequest) GetInstanceStorage() *string {
	return s.InstanceStorage
}

func (s *CreateLindormInstanceRequest) GetLindormNum() *int32 {
	return s.LindormNum
}

func (s *CreateLindormInstanceRequest) GetLindormSpec() *string {
	return s.LindormSpec
}

func (s *CreateLindormInstanceRequest) GetLogDiskCategory() *string {
	return s.LogDiskCategory
}

func (s *CreateLindormInstanceRequest) GetLogNum() *int32 {
	return s.LogNum
}

func (s *CreateLindormInstanceRequest) GetLogSingleStorage() *int32 {
	return s.LogSingleStorage
}

func (s *CreateLindormInstanceRequest) GetLogSpec() *string {
	return s.LogSpec
}

func (s *CreateLindormInstanceRequest) GetLtsNum() *string {
	return s.LtsNum
}

func (s *CreateLindormInstanceRequest) GetLtsSpec() *string {
	return s.LtsSpec
}

func (s *CreateLindormInstanceRequest) GetMultiZoneCombination() *string {
	return s.MultiZoneCombination
}

func (s *CreateLindormInstanceRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateLindormInstanceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateLindormInstanceRequest) GetPayType() *string {
	return s.PayType
}

func (s *CreateLindormInstanceRequest) GetPricingCycle() *string {
	return s.PricingCycle
}

func (s *CreateLindormInstanceRequest) GetPrimaryVSwitchId() *string {
	return s.PrimaryVSwitchId
}

func (s *CreateLindormInstanceRequest) GetPrimaryZoneId() *string {
	return s.PrimaryZoneId
}

func (s *CreateLindormInstanceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateLindormInstanceRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateLindormInstanceRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateLindormInstanceRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateLindormInstanceRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *CreateLindormInstanceRequest) GetSolrNum() *int32 {
	return s.SolrNum
}

func (s *CreateLindormInstanceRequest) GetSolrSpec() *string {
	return s.SolrSpec
}

func (s *CreateLindormInstanceRequest) GetStandbyVSwitchId() *string {
	return s.StandbyVSwitchId
}

func (s *CreateLindormInstanceRequest) GetStandbyZoneId() *string {
	return s.StandbyZoneId
}

func (s *CreateLindormInstanceRequest) GetStreamNum() *int32 {
	return s.StreamNum
}

func (s *CreateLindormInstanceRequest) GetStreamSpec() *string {
	return s.StreamSpec
}

func (s *CreateLindormInstanceRequest) GetTag() []*CreateLindormInstanceRequestTag {
	return s.Tag
}

func (s *CreateLindormInstanceRequest) GetTsdbNum() *int32 {
	return s.TsdbNum
}

func (s *CreateLindormInstanceRequest) GetTsdbSpec() *string {
	return s.TsdbSpec
}

func (s *CreateLindormInstanceRequest) GetVPCId() *string {
	return s.VPCId
}

func (s *CreateLindormInstanceRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *CreateLindormInstanceRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *CreateLindormInstanceRequest) SetArbiterVSwitchId(v string) *CreateLindormInstanceRequest {
	s.ArbiterVSwitchId = &v
	return s
}

func (s *CreateLindormInstanceRequest) SetArbiterZoneId(v string) *CreateLindormInstanceRequest {
	s.ArbiterZoneId = &v
	return s
}

func (s *CreateLindormInstanceRequest) SetArchVersion(v string) *CreateLindormInstanceRequest {
	s.ArchVersion = &v
	return s
}

func (s *CreateLindormInstanceRequest) SetAutoRenewDuration(v string) *CreateLindormInstanceRequest {
	s.AutoRenewDuration = &v
	return s
}

func (s *CreateLindormInstanceRequest) SetAutoRenewal(v bool) *CreateLindormInstanceRequest {
	s.AutoRenewal = &v
	return s
}

func (s *CreateLindormInstanceRequest) SetColdStorage(v int32) *CreateLindormInstanceRequest {
	s.ColdStorage = &v
	return s
}

func (s *CreateLindormInstanceRequest) SetCoreSingleStorage(v int32) *CreateLindormInstanceRequest {
	s.CoreSingleStorage = &v
	return s
}

func (s *CreateLindormInstanceRequest) SetCoreSpec(v string) *CreateLindormInstanceRequest {
	s.CoreSpec = &v
	return s
}

func (s *CreateLindormInstanceRequest) SetDiskCategory(v string) *CreateLindormInstanceRequest {
	s.DiskCategory = &v
	return s
}

func (s *CreateLindormInstanceRequest) SetDuration(v string) *CreateLindormInstanceRequest {
	s.Duration = &v
	return s
}

func (s *CreateLindormInstanceRequest) SetFilestoreNum(v int32) *CreateLindormInstanceRequest {
	s.FilestoreNum = &v
	return s
}

func (s *CreateLindormInstanceRequest) SetFilestoreSpec(v string) *CreateLindormInstanceRequest {
	s.FilestoreSpec = &v
	return s
}

func (s *CreateLindormInstanceRequest) SetInstanceAlias(v string) *CreateLindormInstanceRequest {
	s.InstanceAlias = &v
	return s
}

func (s *CreateLindormInstanceRequest) SetInstanceStorage(v string) *CreateLindormInstanceRequest {
	s.InstanceStorage = &v
	return s
}

func (s *CreateLindormInstanceRequest) SetLindormNum(v int32) *CreateLindormInstanceRequest {
	s.LindormNum = &v
	return s
}

func (s *CreateLindormInstanceRequest) SetLindormSpec(v string) *CreateLindormInstanceRequest {
	s.LindormSpec = &v
	return s
}

func (s *CreateLindormInstanceRequest) SetLogDiskCategory(v string) *CreateLindormInstanceRequest {
	s.LogDiskCategory = &v
	return s
}

func (s *CreateLindormInstanceRequest) SetLogNum(v int32) *CreateLindormInstanceRequest {
	s.LogNum = &v
	return s
}

func (s *CreateLindormInstanceRequest) SetLogSingleStorage(v int32) *CreateLindormInstanceRequest {
	s.LogSingleStorage = &v
	return s
}

func (s *CreateLindormInstanceRequest) SetLogSpec(v string) *CreateLindormInstanceRequest {
	s.LogSpec = &v
	return s
}

func (s *CreateLindormInstanceRequest) SetLtsNum(v string) *CreateLindormInstanceRequest {
	s.LtsNum = &v
	return s
}

func (s *CreateLindormInstanceRequest) SetLtsSpec(v string) *CreateLindormInstanceRequest {
	s.LtsSpec = &v
	return s
}

func (s *CreateLindormInstanceRequest) SetMultiZoneCombination(v string) *CreateLindormInstanceRequest {
	s.MultiZoneCombination = &v
	return s
}

func (s *CreateLindormInstanceRequest) SetOwnerAccount(v string) *CreateLindormInstanceRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateLindormInstanceRequest) SetOwnerId(v int64) *CreateLindormInstanceRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateLindormInstanceRequest) SetPayType(v string) *CreateLindormInstanceRequest {
	s.PayType = &v
	return s
}

func (s *CreateLindormInstanceRequest) SetPricingCycle(v string) *CreateLindormInstanceRequest {
	s.PricingCycle = &v
	return s
}

func (s *CreateLindormInstanceRequest) SetPrimaryVSwitchId(v string) *CreateLindormInstanceRequest {
	s.PrimaryVSwitchId = &v
	return s
}

func (s *CreateLindormInstanceRequest) SetPrimaryZoneId(v string) *CreateLindormInstanceRequest {
	s.PrimaryZoneId = &v
	return s
}

func (s *CreateLindormInstanceRequest) SetRegionId(v string) *CreateLindormInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *CreateLindormInstanceRequest) SetResourceGroupId(v string) *CreateLindormInstanceRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateLindormInstanceRequest) SetResourceOwnerAccount(v string) *CreateLindormInstanceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateLindormInstanceRequest) SetResourceOwnerId(v int64) *CreateLindormInstanceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateLindormInstanceRequest) SetSecurityToken(v string) *CreateLindormInstanceRequest {
	s.SecurityToken = &v
	return s
}

func (s *CreateLindormInstanceRequest) SetSolrNum(v int32) *CreateLindormInstanceRequest {
	s.SolrNum = &v
	return s
}

func (s *CreateLindormInstanceRequest) SetSolrSpec(v string) *CreateLindormInstanceRequest {
	s.SolrSpec = &v
	return s
}

func (s *CreateLindormInstanceRequest) SetStandbyVSwitchId(v string) *CreateLindormInstanceRequest {
	s.StandbyVSwitchId = &v
	return s
}

func (s *CreateLindormInstanceRequest) SetStandbyZoneId(v string) *CreateLindormInstanceRequest {
	s.StandbyZoneId = &v
	return s
}

func (s *CreateLindormInstanceRequest) SetStreamNum(v int32) *CreateLindormInstanceRequest {
	s.StreamNum = &v
	return s
}

func (s *CreateLindormInstanceRequest) SetStreamSpec(v string) *CreateLindormInstanceRequest {
	s.StreamSpec = &v
	return s
}

func (s *CreateLindormInstanceRequest) SetTag(v []*CreateLindormInstanceRequestTag) *CreateLindormInstanceRequest {
	s.Tag = v
	return s
}

func (s *CreateLindormInstanceRequest) SetTsdbNum(v int32) *CreateLindormInstanceRequest {
	s.TsdbNum = &v
	return s
}

func (s *CreateLindormInstanceRequest) SetTsdbSpec(v string) *CreateLindormInstanceRequest {
	s.TsdbSpec = &v
	return s
}

func (s *CreateLindormInstanceRequest) SetVPCId(v string) *CreateLindormInstanceRequest {
	s.VPCId = &v
	return s
}

func (s *CreateLindormInstanceRequest) SetVSwitchId(v string) *CreateLindormInstanceRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateLindormInstanceRequest) SetZoneId(v string) *CreateLindormInstanceRequest {
	s.ZoneId = &v
	return s
}

func (s *CreateLindormInstanceRequest) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateLindormInstanceRequestTag struct {
	// The key of a tag.
	//
	// > You can specify the keys of multiple tags. For example, `Tag.1.Key` specifies the key of the first tag and `Tag.2.Key` specifies the key of the second tag.
	//
	// example:
	//
	// test
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of a tag.
	//
	// > You can specify the values of multiple tags. For example, `Tag.1.Value` specifies the value of the first tag and `Tag.2.Value` specifies the value of the second tag.
	//
	// example:
	//
	// value
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateLindormInstanceRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateLindormInstanceRequestTag) GoString() string {
	return s.String()
}

func (s *CreateLindormInstanceRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateLindormInstanceRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateLindormInstanceRequestTag) SetKey(v string) *CreateLindormInstanceRequestTag {
	s.Key = &v
	return s
}

func (s *CreateLindormInstanceRequestTag) SetValue(v string) *CreateLindormInstanceRequestTag {
	s.Value = &v
	return s
}

func (s *CreateLindormInstanceRequestTag) Validate() error {
	return dara.Validate(s)
}
