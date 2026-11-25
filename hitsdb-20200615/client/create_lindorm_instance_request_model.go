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
	// The vSwitch ID of the arbiter zone for a multi-zone instance. The vSwitch must be in the zone specified by ArbiterZoneId. **To create a multi-zone instance, this parameter is required.**
	//
	// example:
	//
	// vsw-uf6664pqjawb87k36****
	ArbiterVSwitchId *string `json:"ArbiterVSwitchId,omitempty" xml:"ArbiterVSwitchId,omitempty"`
	// The zone ID of the arbiter zone for a multi-zone instance. **To create a multi-zone instance, this parameter is required.**
	//
	// example:
	//
	// cn-shanghai-g
	ArbiterZoneId *string `json:"ArbiterZoneId,omitempty" xml:"ArbiterZoneId,omitempty"`
	// The deployment architecture. Valid values:
	//
	// - **1.0**: single-zone.
	//
	// - **2.0**: multi-zone.
	//
	// If this parameter is not specified, the default value is 1.0. To create a multi-zone instance, set this parameter to 2.0. **To create a multi-zone instance, this parameter is required.**
	//
	// example:
	//
	// 2.0
	ArchVersion *string `json:"ArchVersion,omitempty" xml:"ArchVersion,omitempty"`
	// The auto-renewal duration. Unit: months.
	//
	// Valid values: **1*	- to **12**.
	//
	// > This parameter takes effect only when **AutoRenewal*	- is set to **true**.
	//
	// example:
	//
	// 1
	AutoRenewDuration *string `json:"AutoRenewDuration,omitempty" xml:"AutoRenewDuration,omitempty"`
	// Specifies whether to enable auto-renewal for the instance. Valid values:
	//
	// - **true**: Auto-renewal is enabled.
	//
	// - **false**: Auto-renewal is disabled.
	//
	// Default value: false.
	//
	// > This parameter takes effect only when **PayType*	- is set to **PREPAY*	- (subscription).
	//
	// example:
	//
	// false
	AutoRenewal *bool `json:"AutoRenewal,omitempty" xml:"AutoRenewal,omitempty"`
	// The capacity cloud storage capacity of the instance. Unit: GB. If this parameter is not specified, capacity cloud storage is not enabled. Valid values: **800*	- to **1000000**.
	//
	// example:
	//
	// 800
	ColdStorage *int32 `json:"ColdStorage,omitempty" xml:"ColdStorage,omitempty"`
	// The storage capacity per core node for a multi-zone instance. Valid values: 400 to 64000. Unit: GB. **To create a multi-zone instance, this parameter is required.**
	//
	// example:
	//
	// 400
	CoreSingleStorage *int32 `json:"CoreSingleStorage,omitempty" xml:"CoreSingleStorage,omitempty"`
	// The node specifications for local disk type of the instance.
	//
	// When the local storage type is local_ssd_pro, valid values are as follows. I3 instance types currently support only subscription:
	//
	// - **lindorm.i4.xlarge**: 4 cores, 32 GB (I4).
	//
	// - **lindorm.i4.2xlarge**: 8 cores, 64 GB (I4).
	//
	// - **lindorm.i4.4xlarge**: 16 cores, 128 GB (I4).
	//
	// - **lindorm.i4.8xlarge**: 32 cores, 256 GB (I4).
	//
	// - **lindorm.i3.xlarge**: 4 cores, 32 GB (I3).
	//
	// - **lindorm.i3.2xlarge**: 8 cores, 64 GB (I3).
	//
	// - **lindorm.i3.4xlarge**: 16 cores, 128 GB (I3).
	//
	// - **lindorm.i3.8xlarge**: 32 cores, 256 GB (I3).
	//
	// - **lindorm.i2.xlarge**: 4 cores, 32 GB (I2).
	//
	// - **lindorm.i2.2xlarge**: 8 cores, 64 GB (I2).
	//
	// - **lindorm.i2.4xlarge**: 16 cores, 128 GB (I2).
	//
	// - **lindorm.i2.8xlarge**: 32 cores, 256 GB (I2).
	//
	// When the local storage type is local_hdd_pro, valid values are as follows:
	//
	// - **lindorm.sd3c.3xlarge**: 14 cores, 56 GB (D3C PRO).
	//
	// - **lindorm.sd3c.7xlarge**: 28 cores, 112 GB (D3C PRO).
	//
	// - **lindorm.sd3c.14xlarge**: 56 cores, 224 GB (D3C PRO).
	//
	// - **lindorm.d2c.6xlarge**: 24 cores, 88 GB (D2C).
	//
	// - **lindorm.d2c.12xlarge**: 48 cores, 176 GB (D2C).
	//
	// - **lindorm.d2c.24xlarge**: 96 cores, 352 GB (D2C).
	//
	// - **lindorm.d2s.5xlarge**: 20 cores, 88 GB (D2S).
	//
	// - **lindorm.d2s.10xlarge**: 40 cores, 176 GB (D2S).
	//
	// - **lindorm.d1.2xlarge**: 8 cores, 32 GB (D1NE).
	//
	// - **lindorm.d1.4xlarge**: 16 cores, 64 GB (D1NE).
	//
	// - **lindorm.d1.6xlarge**: 24 cores, 96 GB (D1NE).
	//
	// example:
	//
	// lindorm.i2.xlarge
	CoreSpec *string `json:"CoreSpec,omitempty" xml:"CoreSpec,omitempty"`
	// The storage type of the instance. Valid values:
	//
	// - **cloud_efficiency**: standard cloud storage.
	//
	// - **cloud_ssd**: performance cloud storage.
	//
	// - **cloud_essd**: enhanced performance cloud storage.
	//
	// - **cloud_essd_pl0**: performance cloud storage PL0.
	//
	// - **capacity_cloud_storage**: capacity cloud storage (not supported for multi-zone instances).
	//
	// - **local_ssd_pro**: local SSD (not supported for multi-zone instances).
	//
	// - **local_hdd_pro**: local HDD (not supported for multi-zone instances).
	//
	// This parameter is required.
	//
	// example:
	//
	// cloud_efficiency
	DiskCategory *string `json:"DiskCategory,omitempty" xml:"DiskCategory,omitempty"`
	// The subscription duration of the instance. Valid values:
	//
	// - If PricingCycle is set to **Month**, valid values are **1*	- to **9**.
	//
	// - If PricingCycle is set to **Year**, valid values are **1*	- to **3**.
	//
	// > This parameter is available and required when PayType is set to **PREPAY**.
	//
	// example:
	//
	// 1
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The number of file engine nodes of the instance. Valid values:
	//
	// - If the billing method of the instance is **PREPAY**, valid values are **0*	- to **60**.
	//
	// - If the billing method of the instance is **POSTPAY**, valid values are **0*	- to **8**.
	//
	// example:
	//
	// 2
	FilestoreNum *int32 `json:"FilestoreNum,omitempty" xml:"FilestoreNum,omitempty"`
	// The file engine node specifications of the instance. Valid values:
	//
	// - **lindorm.c.xlarge**: 4 cores, 8 GB (standard).
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
	// The storage capacity of the instance. Unit: GB.
	//
	// example:
	//
	// 480
	InstanceStorage *string `json:"InstanceStorage,omitempty" xml:"InstanceStorage,omitempty"`
	// The number of LindormTable engine nodes of the instance.
	//
	// To create a single-zone instance, valid values are **0*	- to **90**.
	//
	// **To create a multi-zone instance, this parameter is required.*	- For instances with non-local disk storage types, valid values are **4*	- to **400**. For instances with local disk storage types, valid values are **6*	- to **400**.
	//
	// example:
	//
	// 2
	LindormNum *int32 `json:"LindormNum,omitempty" xml:"LindormNum,omitempty"`
	// The LindormTable engine node specifications of the instance. Valid values:
	//
	// - **lindorm.g.xlarge**: 4 cores, 16 GB (dedicated).
	//
	// - **lindorm.c.2xlarge**: 8 cores, 16 GB (dedicated).
	//
	// - **lindorm.g.2xlarge**: 8 cores, 32 GB (dedicated).
	//
	// - **lindorm.c.4xlarge**: 16 cores, 32 GB (dedicated).
	//
	// - **lindorm.g.4xlarge**: 16 cores, 64 GB (dedicated).
	//
	// - **lindorm.c.8xlarge**: 32 cores, 64 GB (dedicated).
	//
	// - **lindorm.g.8xlarge**: 32 cores, 128 GB (dedicated).
	//
	// example:
	//
	// lindorm.g.4xlarge
	LindormSpec *string `json:"LindormSpec,omitempty" xml:"LindormSpec,omitempty"`
	// The disk type of log nodes for a multi-zone instance. Valid values:
	//
	// - **cloud_efficiency**: standard cloud storage.
	//
	// - **cloud_ssd**: performance cloud storage.
	//
	// **To create a multi-zone instance, this parameter is required.**
	//
	// example:
	//
	// cloud_ssd
	LogDiskCategory *string `json:"LogDiskCategory,omitempty" xml:"LogDiskCategory,omitempty"`
	// The number of log nodes for a multi-zone instance. Valid values: 4 to 400. **To create a multi-zone instance, this parameter is required.**
	//
	// example:
	//
	// 4
	LogNum *int32 `json:"LogNum,omitempty" xml:"LogNum,omitempty"`
	// The storage capacity per log node for a multi-zone instance. Valid values: 400 to 64000. Unit: GB. **To create a multi-zone instance, this parameter is required.**
	//
	// example:
	//
	// 400
	LogSingleStorage *int32 `json:"LogSingleStorage,omitempty" xml:"LogSingleStorage,omitempty"`
	// The log node specifications for a multi-zone instance. Valid values:
	//
	// - **lindorm.sn1.large**: 4 cores, 8 GB (dedicated).
	//
	// - **lindorm.sn1.2xlarge**: 8 cores, 16 GB (dedicated).
	//
	// **To create a multi-zone instance, this parameter is required and active.**
	//
	// example:
	//
	// lindorm.sn1.large
	LogSpec *string `json:"LogSpec,omitempty" xml:"LogSpec,omitempty"`
	// The number of LTS engine nodes of the instance. Valid values: **0*	- to **60**.
	//
	// example:
	//
	// 2
	LtsNum *string `json:"LtsNum,omitempty" xml:"LtsNum,omitempty"`
	// The LTS engine node specifications of the instance. Valid values:
	//
	// - **lindorm.c.xlarge**: 4 cores, 8 GB (dedicated).
	//
	// - **lindorm.g.xlarge**: 4 cores, 16 GB (dedicated).
	//
	// - **lindorm.c.2xlarge**: 8 cores, 16 GB (dedicated).
	//
	// - **lindorm.g.2xlarge**: 8 cores, 32 GB (dedicated).
	//
	// - **lindorm.c.4xlarge**: 16 cores, 32 GB (dedicated).
	//
	// - **lindorm.g.4xlarge**: 16 cores, 64 GB (dedicated).
	//
	// - **lindorm.c.8xlarge**: 32 cores, 64 GB (dedicated).
	//
	// - **lindorm.g.8xlarge**: 32 cores, 128 GB (dedicated).
	//
	// example:
	//
	// lindorm.g.xlarge
	LtsSpec *string `json:"LtsSpec,omitempty" xml:"LtsSpec,omitempty"`
	// The multi-zone combination. You can check the supported zone combinations on the buy page.
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
	// **To create a multi-zone instance, this parameter is required.**
	//
	// example:
	//
	// cn-shanghai-efg-aliyun
	MultiZoneCombination *string `json:"MultiZoneCombination,omitempty" xml:"MultiZoneCombination,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The billing method of the instance. Valid values:
	//
	// - **PREPAY**: subscription.
	//
	// - **POSTPAY**: pay-as-you-go.
	//
	// This parameter is required.
	//
	// example:
	//
	// POSTPAY
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// The billing cycle of the instance. Valid values:
	//
	// - **Month**: monthly billing.
	//
	// - **Year**: yearly billing.
	//
	// > This parameter is available and required when PayType is set to **PREPAY**.
	//
	// example:
	//
	// Month
	PricingCycle *string `json:"PricingCycle,omitempty" xml:"PricingCycle,omitempty"`
	// The vSwitch ID of the primary zone for a multi-zone instance. The vSwitch must be in the zone specified by PrimaryZoneId. **To create a multi-zone instance, this parameter is required.**
	//
	// example:
	//
	// vsw-uf6fdqa7c0pipnqzq****
	PrimaryVSwitchId *string `json:"PrimaryVSwitchId,omitempty" xml:"PrimaryVSwitchId,omitempty"`
	// The zone ID of the primary zone for a multi-zone instance. **To create a multi-zone instance, this parameter is required.**
	//
	// example:
	//
	// cn-shanghai-e
	PrimaryZoneId *string `json:"PrimaryZoneId,omitempty" xml:"PrimaryZoneId,omitempty"`
	// The region ID of the instance. You can call [DescribeRegions](https://help.aliyun.com/document_detail/426062.html) to query available regions. Use this parameter to specify the region in which to create the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-aek2i6weeb4nfii
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The number of search index engine nodes of the instance. Valid values: **0*	- to **60**.
	//
	// example:
	//
	// 2
	SolrNum *int32 `json:"SolrNum,omitempty" xml:"SolrNum,omitempty"`
	// The search index engine node specifications of the instance. Valid values:
	//
	// - **lindorm.g.xlarge**: 4 cores, 16 GB (dedicated).
	//
	// - **lindorm.c.2xlarge**: 8 cores, 16 GB (dedicated).
	//
	// - **lindorm.g.2xlarge**: 8 cores, 32 GB (dedicated).
	//
	// - **lindorm.c.4xlarge**: 16 cores, 32 GB (dedicated).
	//
	// - **lindorm.g.4xlarge**: 16 cores, 64 GB (dedicated).
	//
	// - **lindorm.c.8xlarge**: 32 cores, 64 GB (dedicated).
	//
	// - **lindorm.g.8xlarge**: 32 cores, 128 GB (dedicated).
	//
	// example:
	//
	// lindorm.g.xlarge
	SolrSpec *string `json:"SolrSpec,omitempty" xml:"SolrSpec,omitempty"`
	// The vSwitch ID of the secondary zone for a multi-zone instance. The vSwitch must be in the zone specified by StandbyZoneId. **To create a multi-zone instance, this parameter is required.**
	//
	// example:
	//
	// vsw-2zec0kcn08cgdtr6****
	StandbyVSwitchId *string `json:"StandbyVSwitchId,omitempty" xml:"StandbyVSwitchId,omitempty"`
	// The zone ID of the secondary zone for a multi-zone instance. **To create a multi-zone instance, this parameter is required.**
	//
	// example:
	//
	// cn-shanghai-f
	StandbyZoneId *string `json:"StandbyZoneId,omitempty" xml:"StandbyZoneId,omitempty"`
	// The number of stream engine nodes of the instance. Valid values: **0*	- to **60**.
	//
	// example:
	//
	// 2
	StreamNum *int32 `json:"StreamNum,omitempty" xml:"StreamNum,omitempty"`
	// The stream engine node specifications of the instance. Valid values:
	//
	// - **lindorm.g.xlarge**: 4 cores, 16 GB (dedicated).
	//
	// - **lindorm.c.2xlarge**: 8 cores, 16 GB (dedicated).
	//
	// - **lindorm.g.2xlarge**: 8 cores, 32 GB (dedicated).
	//
	// - **lindorm.c.4xlarge**: 16 cores, 32 GB (dedicated).
	//
	// - **lindorm.g.4xlarge**: 16 cores, 64 GB (dedicated).
	//
	// - **lindorm.c.8xlarge**: 32 cores, 64 GB (dedicated).
	//
	// - **lindorm.g.8xlarge**: 32 cores, 128 GB (dedicated).
	//
	// example:
	//
	// lindorm.g.xlarge
	StreamSpec *string `json:"StreamSpec,omitempty" xml:"StreamSpec,omitempty"`
	// The list of tags.
	Tag []*CreateLindormInstanceRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The number of LindormTSDB engine nodes of the instance. Valid values:
	//
	// - If the billing method of the instance is **PREPAY**, valid values are **0*	- to **24**.
	//
	// - If the billing method of the instance is **POSTPAY**, valid values are **0*	- to **32**.
	//
	// example:
	//
	// 2
	TsdbNum *int32 `json:"TsdbNum,omitempty" xml:"TsdbNum,omitempty"`
	// The LindormTSDB engine node specifications of the instance. Valid values:
	//
	// - **lindorm.g.xlarge**: 4 cores, 16 GB (dedicated).
	//
	// - **lindorm.g.2xlarge**: 8 cores, 32 GB (dedicated).
	//
	// - **lindorm.g.4xlarge**: 16 cores, 64 GB (dedicated).
	//
	// - **lindorm.g.8xlarge**: 32 cores, 128 GB (dedicated).
	//
	// example:
	//
	// lindorm.g.xlarge
	TsdbSpec *string `json:"TsdbSpec,omitempty" xml:"TsdbSpec,omitempty"`
	// The VPC ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpc-bp1nme44gek34slfc****
	VPCId *string `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	// The vSwitch ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// vsw-bp1e7clcw529l773d****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The zone ID of the instance.
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
	// The tag key. Valid values of N: 1 to 20.
	//
	// > You can specify multiple tag keys. For example, the Key in the first pair specifies the first tag key, and the Key in the second pair specifies the second tag key.
	//
	// example:
	//
	// test
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value. Valid values of N: 1 to 20.
	//
	// > You can specify multiple tag values. For example, the Value in the first pair specifies the first tag value, and the Value in the second pair specifies the second tag value.
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
