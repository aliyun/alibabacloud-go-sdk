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
	// The ID of the vSwitch that is specified for the zone for the coordinate node of the instance. The vSwitch must be deployed in the zone specified by the ArbiterZoneId parameter. **This parameter is required if you want to create a multi-zone instance**.
	//
	// example:
	//
	// vsw-uf6664pqjawb87k36****
	ArbiterVSwitchId *string `json:"ArbiterVSwitchId,omitempty" xml:"ArbiterVSwitchId,omitempty"`
	// The ID of the zone for the coordinate node of the instance. **This parameter is required if you want to create a multi-zone instance**.
	//
	// example:
	//
	// cn-shanghai-g
	ArbiterZoneId *string `json:"ArbiterZoneId,omitempty" xml:"ArbiterZoneId,omitempty"`
	// The architecture of the instance. Valid values:
	//
	// 	- **1.0**: The instance that you want to create is a single-zone instance.
	//
	// 	- **2.0**: The instance that you want to create is a multi-zone instance.
	//
	// By default, the value of this parameter is 1.0. To create a multi-zone instance, set this parameter to 2.0. **This parameter is required if you want to create a multi-zone instance**.
	//
	// example:
	//
	// 2.0
	ArchVersion *string `json:"ArchVersion,omitempty" xml:"ArchVersion,omitempty"`
	// The auto-renewal duration. Unit: month.
	//
	// Valid values: **1*	- to **12**.
	//
	// >  This parameter is available only when the **AutoRenewal*	- parameter is set to **true**.
	//
	// example:
	//
	// 1
	AutoRenewDuration *string `json:"AutoRenewDuration,omitempty" xml:"AutoRenewDuration,omitempty"`
	// Specifies whether to enable auto-renewal for the instance. Valid values:
	//
	// 	- **true**: enables auto-renewal.
	//
	// 	- **false**: disables auto-renewal.
	//
	// Default value: false.
	//
	// >  This parameter is available only when the **PayType*	- parameter is set to **PREPAY**.
	//
	// example:
	//
	// false
	AutoRenewal *bool `json:"AutoRenewal,omitempty" xml:"AutoRenewal,omitempty"`
	// The cold storage capacity of the instance. By default, if you leave this parameter unspecified, cold storage is not enabled for the instance. Unit: GB. Valid values: **800*	- to **1000000**.
	//
	// example:
	//
	// 800
	ColdStorage *int32 `json:"ColdStorage,omitempty" xml:"ColdStorage,omitempty"`
	// The storage capacity of the disk of a single core node. Valid values: 400 to 64000. Unit: GB. **This parameter is required if you want to create a multi-zone instance**.
	//
	// example:
	//
	// 400
	CoreSingleStorage *int32 `json:"CoreSingleStorage,omitempty" xml:"CoreSingleStorage,omitempty"`
	// The specification of the nodes in the instance if you set DiskCategory to local_ssd_pro or local_hdd_pro.
	//
	// Valid values when DiskCategory is set to local_ssd_pro (i3 instance types support only subscription instances):
	//
	// 	- **lindorm.i4.xlarge**: Each node has 4 CPU cores and 32 GB of memory.
	//
	// 	- **lindorm.i4.2xlarge**: Each node has 8 CPU cores and 64 GB of memory.
	//
	// 	- **lindorm.i4.4xlarge**: Each node has 16 CPU cores and 128 GB of memory.
	//
	// 	- **lindorm.i4.8xlarge**: Each node has 32 CPU cores and 256 GB of memory.
	//
	// 	- **lindorm.i3.xlarge**: Each node has 4 CPU cores and 32 GB of memory.
	//
	// 	- **lindorm.i3.2xlarge**: Each node has 8 CPU cores and 64 GB of memory.
	//
	// 	- **lindorm.i3.4xlarge**: Each node has 16 CPU cores and 128 GB of memory.
	//
	// 	- **lindorm.i3.8xlarge**: Each node has 32 CPU cores and 256 GB of memory.
	//
	// 	- **lindorm.i2.xlarge**: Each node has 4 CPU cores and 32 GB of memory.
	//
	// 	- **lindorm.i2.2xlarge**: Each node has 8 CPU cores and 64 GB of memory.
	//
	// 	- **lindorm.i2.4xlarge**: Each node has 16 CPU cores and 128 GB of memory.
	//
	// 	- **lindorm.i2.8xlarge**: Each node has 32 CPU cores and 256 GB of memory.
	//
	// Valid values when DiskCategory is set to local_hhd_pro:
	//
	// 	- **lindorm.sd3c.3xlarge**: Each node has 14 CPU cores and 56 GB of memory.
	//
	// 	- **lindorm.sd3c.7xlarge**: Each node has 28 CPU cores and 112 GB of memory.
	//
	// 	- **lindorm.sd3c.14xlarge**: Each node has 56 CPU cores and 224 GB of memory.
	//
	// 	- **lindorm.d2c.6xlarge**: Each node has 24 CPU cores and 88 GB of memory.
	//
	// 	- **lindorm.d2c.12xlarge**: Each node has 48 CPU cores and 176 GB of memory.
	//
	// 	- **lindorm.d2c.24xlarge**: Each node has 96 CPU cores and 352 GB of memory.
	//
	// 	- **lindorm.d2s.5xlarge**: Each node has 20 CPU cores and 88 GB of memory.
	//
	// 	- **lindorm.d2s.10xlarge**: Each node has 40 CPU cores and 176 GB of memory.
	//
	// 	- **lindorm.d1.2xlarge**: Each node has 8 CPU cores and 32 GB of memory.
	//
	// 	- **lindorm.d1.4xlarge**: Each node has 16 CPU cores and 64 GB of memory.
	//
	// 	- **lindorm.d1.6xlarge**: Each node has 24 CPU cores and 96 GB of memory.
	//
	// example:
	//
	// lindorm.i2.xlarge
	CoreSpec *string `json:"CoreSpec,omitempty" xml:"CoreSpec,omitempty"`
	// The storage type of the instance. Valid values:
	//
	// 	- **cloud_efficiency**: This instance uses the Standard type of storage.
	//
	// 	- **cloud_ssd**: This instance uses the Performance type of storage.
	//
	// 	- **capacity_cloud_storage**: This instance uses the Capacity type of storage.
	//
	// 	- **local_ssd_pro**: This instance uses local SSDs.
	//
	// 	- **local_hdd_pro**: This instance uses local HDDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// cloud_efficiency
	DiskCategory *string `json:"DiskCategory,omitempty" xml:"DiskCategory,omitempty"`
	// The subscription period of the instance. The valid values of this parameter depend on the value of the PricingCycle parameter.
	//
	// 	- If PricingCycle is set to **Month**, set this parameter to an integer that ranges from **1*	- to **9**.
	//
	// 	- If PricingCycle is set to **Year**, set this parameter to an integer that ranges from **1*	- to **3**.
	//
	// > This parameter is available and required when the PayType parameter is set to **PREPAY**.
	//
	// example:
	//
	// 1
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The number of LindormDFS nodes in the instance. The valid values of this parameter depend on the value of the PayType parameter.
	//
	// 	- If the PayType parameter is set to **PREPAY**, set this parameter to an integer that ranges from **0*	- to **60**.
	//
	// 	- If the PayType parameter is set to **POSTPAY**, set this parameter to an integer that ranges from **0*	- to **8**.
	//
	// example:
	//
	// 2
	FilestoreNum *int32 `json:"FilestoreNum,omitempty" xml:"FilestoreNum,omitempty"`
	// The specification of LindormDFS nodes in the instance. Set the value of this parameter to **lindorm.c.xlarge**, which indicates that each node has 4 dedicated CPU cores and 8 GB of dedicated memory.
	//
	// example:
	//
	// lindorm.c.xlarge
	FilestoreSpec *string `json:"FilestoreSpec,omitempty" xml:"FilestoreSpec,omitempty"`
	// The name of the instance that you want to create.
	//
	// example:
	//
	// lindorm_test
	InstanceAlias *string `json:"InstanceAlias,omitempty" xml:"InstanceAlias,omitempty"`
	// The storage capacity of the instance you want to create. Unit: GB.
	//
	// example:
	//
	// 480
	InstanceStorage *string `json:"InstanceStorage,omitempty" xml:"InstanceStorage,omitempty"`
	// The number of LindormTable nodes in the instance. The valid values of this parameter depend on the value of the PayType parameter.
	//
	// 	- If the PayType parameter is set to **PREPAY**, set this parameter to an integer that ranges from **0*	- to **90**.
	//
	// 	- If the PayType parameter is set to **POSTPAY**, set this parameter to an integer that ranges from **0*	- to **400**.
	//
	// **This parameter is required if you want to create a multi-zone instance**.  The valid values of this parameter range from 4 to 400 if you want to create a multi-zone instance.
	//
	// example:
	//
	// 2
	LindormNum *int32 `json:"LindormNum,omitempty" xml:"LindormNum,omitempty"`
	// The specification of LindormTable nodes in the instance. Valid values:
	//
	// 	- **lindorm.c.xlarge**: Each node has 4 dedicated CPU cores and 8 GB of dedicated memory.
	//
	// 	- **lindorm.c.2xlarge**: Each node has 8 dedicated CPU cores and 16 GB of dedicated memory.
	//
	// 	- **lindorm.c.4xlarge**: Each node has 16 dedicated CPU cores and 32 GB of dedicated memory.
	//
	// 	- **lindorm.c.8xlarge**: Each node has 32 dedicated CPU cores and 64 GB of dedicated memory.
	//
	// example:
	//
	// lindorm.c.xlarge
	LindormSpec *string `json:"LindormSpec,omitempty" xml:"LindormSpec,omitempty"`
	// The disk type of the log nodes. Valid values:
	//
	// 	- **cloud_efficiency**: This instance uses the Standard type of storage.
	//
	// 	- **cloud_ssd**: This instance uses the Performance type of storage.
	//
	// **This parameter is required if you want to create a multi-zone instance**.
	//
	// example:
	//
	// cloud_ssd
	LogDiskCategory *string `json:"LogDiskCategory,omitempty" xml:"LogDiskCategory,omitempty"`
	// The number of the log nodes. Valid values: 4 to 400. **This parameter is required if you want to create a multi-zone instance**.
	//
	// example:
	//
	// 4
	LogNum *int32 `json:"LogNum,omitempty" xml:"LogNum,omitempty"`
	// The storage capacity of the disk of a single log node. Valid values: 400 to 64000. Unit: GB. **This parameter is required if you want to create a multi-zone instance**.
	//
	// example:
	//
	// 400
	LogSingleStorage *int32 `json:"LogSingleStorage,omitempty" xml:"LogSingleStorage,omitempty"`
	// The type of the log nodes. Valid values:
	//
	// 	- **lindorm.sn1.xlarge**: Each node has 4 dedicated CPU cores and 8 GB of dedicated memory.
	//
	// 	- **lindorm.sn1.2xlarge**: Each node has 8 dedicated CPU cores and 16 GB of dedicated memory.
	//
	// **This parameter is required if you want to create a multi-zone instance**.
	//
	// example:
	//
	// lindorm.sn1.large
	LogSpec *string `json:"LogSpec,omitempty" xml:"LogSpec,omitempty"`
	// The number of LTS nodes in the instance. Valid values: **0*	- to **60**.
	//
	// example:
	//
	// 2
	LtsNum *string `json:"LtsNum,omitempty" xml:"LtsNum,omitempty"`
	// The specification of LTS nodes in the instance. Valid values:
	//
	// 	- **lindorm.c.xlarge**: Each node has 4 dedicated CPU cores and 8 GB of dedicated memory.
	//
	// 	- **lindorm.g.xlarge**: Each node has 4 dedicated CPU cores and 16 GB of dedicated memory.
	//
	// 	- **lindorm.c.2xlarge**: Each node has 8 dedicated CPU cores and 16 GB of dedicated memory.
	//
	// 	- **lindorm.g.2xlarge**: Each node has 8 dedicated CPU cores and 32 GB of dedicated memory.
	//
	// 	- **lindorm.c.4xlarge**: Each node has 16 dedicated CPU cores and 32 GB of dedicated memory.
	//
	// 	- **lindorm.g.4xlarge**: Each node has 16 dedicated CPU cores and 64 GB of dedicated memory.
	//
	// 	- **lindorm.c.8xlarge**: Each node has 32 dedicated CPU cores and 64 GB of dedicated memory.
	//
	// 	- **lindorm.g.8xlarge**: Each node has 32 dedicated CPU cores and 128 GB of dedicated memory.
	//
	// example:
	//
	// lindorm.g.xlarge
	LtsSpec *string `json:"LtsSpec,omitempty" xml:"LtsSpec,omitempty"`
	// The combinations of zones that are available for the multi-zone instance. You can go to the purchase page of Lindorm to view the supported zone combinations.
	//
	// 	- **ap-southeast-5abc-aliyun**: Zone A+B+C in the Indonesia (Jakarta) region.
	//
	// 	- **cn-hangzhou-ehi-aliyun**: Zone E+H+I in the China (Hangzhou) region.
	//
	// 	- **cn-beijing-acd-aliyun**: Zone A+C+D in the China (Beijing) region.
	//
	// 	- **ap-southeast-1-abc-aliyun**: Zone A+B+C in the Singapore region.
	//
	// 	- **cn-zhangjiakou-abc-aliyun**: Zone A+B+C in the China (Zhangjiakou) region.
	//
	// 	- **cn-shanghai-efg-aliyun**: Zone E+F+G in the China (Shanghai) region.
	//
	// 	- **cn-shanghai-abd-aliyun**: Zone A+B+D in the China (Shanghai) region.
	//
	// 	- **cn-hangzhou-bef-aliyun**: Zone B+E+F in the China (Hangzhou) region.
	//
	// 	- **cn-hangzhou-bce-aliyun**: Zone B+C+E in the China (Hangzhou) region.
	//
	// 	- **cn-beijing-fgh-aliyun**: Zone F+G+H in the China (Beijing) region.
	//
	// 	- **cn-shenzhen-abc-aliyun**: Zone A+B+C in the China (Shenzhen) region.
	//
	// **This parameter is required if you want to create a multi-zone instance**.
	//
	// example:
	//
	// cn-shanghai-efg-aliyun
	MultiZoneCombination *string `json:"MultiZoneCombination,omitempty" xml:"MultiZoneCombination,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The billing method of the instance you want to create. Valid values:
	//
	// 	- **PREPAY**: subscription.
	//
	// 	- **POSTPAY**: pay-as-you-go.
	//
	// This parameter is required.
	//
	// example:
	//
	// POSTPAY
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// The period based on which you are charged for the instance. Valid values:
	//
	// 	- **Month**: You are charged for the instance on a monthly basis.
	//
	// 	- **Year**: You are charged for the instance on a yearly basis.
	//
	// > This parameter is available and required when the PayType parameter is set to **PREPAY**.
	//
	// example:
	//
	// Month
	PricingCycle *string `json:"PricingCycle,omitempty" xml:"PricingCycle,omitempty"`
	// The ID of the vSwitch that is specified for the secondary zone of the instance. The vSwitch must be deployed in the zone specified by the StandbyZoneId parameter. **This parameter is required if you want to create a multi-zone instance**.
	//
	// example:
	//
	// vsw-uf6fdqa7c0pipnqzq****
	PrimaryVSwitchId *string `json:"PrimaryVSwitchId,omitempty" xml:"PrimaryVSwitchId,omitempty"`
	// Multi-zone instance, availability zone ID of the primary zone. **This parameter is required if you need to create a multi-zone instance.**
	//
	// example:
	//
	// cn-shanghai-e
	PrimaryZoneId *string `json:"PrimaryZoneId,omitempty" xml:"PrimaryZoneId,omitempty"`
	// The ID of the region in which you want to create the instance. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/426062.html) operation to query the region in which you can create the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the Lindorm instance belongs.
	//
	// example:
	//
	// rg-aek2i6weeb4nfii
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The number of LindormSearch nodes in the instance. Valid values: integers from **0*	- to **60**.
	//
	// example:
	//
	// 2
	SolrNum *int32 `json:"SolrNum,omitempty" xml:"SolrNum,omitempty"`
	// The specification of the LindormSearch nodes in the instance. Valid values:
	//
	// 	- **lindorm.g.xlarge**: Each node has 4 dedicated CPU cores and 16 GB of dedicated memory.
	//
	// 	- **lindorm.c.2xlarge**: Each node has 8 dedicated CPU cores and 16 GB of dedicated memory.
	//
	// 	- **lindorm.g.2xlarge**: Each node has 8 dedicated CPU cores and 32 GB of dedicated memory.
	//
	// 	- **lindorm.c.4xlarge**: Each node has 16 dedicated CPU cores and 32 GB of dedicated memory.
	//
	// 	- **lindorm.g.4xlarge**: Each node has 16 dedicated CPU cores and 64 GB of dedicated memory.
	//
	// 	- **lindorm.c.8xlarge**: Each node has 32 dedicated CPU cores and 64 GB of dedicated memory.
	//
	// 	- **lindorm.g.8xlarge**: Each node has 32 dedicated CPU cores and 128 GB of dedicated memory.
	//
	// example:
	//
	// lindorm.g.xlarge
	SolrSpec *string `json:"SolrSpec,omitempty" xml:"SolrSpec,omitempty"`
	// The ID of the vSwitch that is specified for the secondary zone of the instance. The vSwitch must be deployed in the zone specified by the StandbyZoneId parameter. **This parameter is required if you want to create a multi-zone instance**.
	//
	// example:
	//
	// vsw-2zec0kcn08cgdtr6****
	StandbyVSwitchId *string `json:"StandbyVSwitchId,omitempty" xml:"StandbyVSwitchId,omitempty"`
	// The ID of the secondary zone of the instance. **This parameter is required if you want to create a multi-zone instance**.
	//
	// example:
	//
	// cn-shanghai-f
	StandbyZoneId *string `json:"StandbyZoneId,omitempty" xml:"StandbyZoneId,omitempty"`
	// The number of LindormStream nodes in the instance. Valid values: integers from **0*	- to **60**.
	//
	// example:
	//
	// 2
	StreamNum *int32 `json:"StreamNum,omitempty" xml:"StreamNum,omitempty"`
	// The specification of the LindormStream nodes in the instance. Valid values:
	//
	// 	- **lindorm.g.xlarge**: Each node has 4 dedicated CPU cores and 16 GB of dedicated memory.
	//
	// 	- **lindorm.g.2xlarge**: Each node has 8 dedicated CPU cores and 32 GB of dedicated memory.
	//
	// 	- **lindorm.g.4xlarge**: Each node has 16 dedicated CPU cores and 64 GB of dedicated memory.
	//
	// 	- **lindorm.g.8xlarge**: Each node has 32 dedicated CPU cores and 128 GB of dedicated memory.
	//
	// example:
	//
	// lindorm.g.xlarge
	StreamSpec *string `json:"StreamSpec,omitempty" xml:"StreamSpec,omitempty"`
	// The tags that are added to instances.
	Tag []*CreateLindormInstanceRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The number of the LindormTSDB nodes in the instance. The valid values of this parameter depend on the value of the PayType parameter.
	//
	// 	- If the PayType parameter is set to **PREPAY**, set this parameter to an integer that ranges from **0*	- to **24**.
	//
	// 	- If the PayType parameter is set to **POSTPAY**, set this parameter to an integer that ranges from **0*	- to **32**.
	//
	// example:
	//
	// 2
	TsdbNum *int32 `json:"TsdbNum,omitempty" xml:"TsdbNum,omitempty"`
	// The specification of the LindormTSDB nodes in the instance. Valid values:
	//
	// 	- **lindorm.g.xlarge**: Each node has 4 dedicated CPU cores and 16 GB of dedicated memory.
	//
	// 	- **lindorm.g.2xlarge**: Each node has 8 dedicated CPU cores and 32 GB of dedicated memory.
	//
	// 	- **lindorm.g.4xlarge**: Each node has 16 dedicated CPU cores and 64 GB of dedicated memory.
	//
	// 	- **lindorm.g.8xlarge**: Each node has 32 dedicated CPU cores and 128 GB of dedicated memory.
	//
	// example:
	//
	// lindorm.g.xlarge
	TsdbSpec *string `json:"TsdbSpec,omitempty" xml:"TsdbSpec,omitempty"`
	// The ID of the VPC in which you want to create the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpc-bp1nme44gek34slfc****
	VPCId *string `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	// The ID of the vSwitch to which you want the instance to connect.
	//
	// This parameter is required.
	//
	// example:
	//
	// vsw-bp1e7clcw529l773d****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The ID of the zone in which you want to create the instance.
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
	return dara.Validate(s)
}

type CreateLindormInstanceRequestTag struct {
	// The tag key. Valid values of N: 1 to 20.
	//
	// >  You can specify the keys of multiple tags. For example, you can specify the key of the first tag in the first key-value pair contained in the value of this parameter and specify the key of the second tag in the second key-value pair.
	//
	// example:
	//
	// test
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value. Valid values of N: 1 to 20.
	//
	// >  You can specify the values of multiple tags. For example, you can specify the value of the first tag in the first key-value pair contained in the value of this parameter and specify the value of the second tag in the second key-value pair.
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
