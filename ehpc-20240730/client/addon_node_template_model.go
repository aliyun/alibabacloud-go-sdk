// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddonNodeTemplate interface {
	dara.Model
	String() string
	GoString() string
	SetAutoRenew(v bool) *AddonNodeTemplate
	GetAutoRenew() *bool
	SetAutoRenewPeriod(v int32) *AddonNodeTemplate
	GetAutoRenewPeriod() *int32
	SetDataDisks(v []*AddonNodeTemplateDataDisks) *AddonNodeTemplate
	GetDataDisks() []*AddonNodeTemplateDataDisks
	SetDuration(v int32) *AddonNodeTemplate
	GetDuration() *int32
	SetEnableHT(v bool) *AddonNodeTemplate
	GetEnableHT() *bool
	SetImageId(v string) *AddonNodeTemplate
	GetImageId() *string
	SetInstanceChargeType(v string) *AddonNodeTemplate
	GetInstanceChargeType() *string
	SetInstanceId(v string) *AddonNodeTemplate
	GetInstanceId() *string
	SetInstanceType(v string) *AddonNodeTemplate
	GetInstanceType() *string
	SetOsName(v string) *AddonNodeTemplate
	GetOsName() *string
	SetOsNameEN(v string) *AddonNodeTemplate
	GetOsNameEN() *string
	SetPeriod(v int32) *AddonNodeTemplate
	GetPeriod() *int32
	SetPeriodUnit(v string) *AddonNodeTemplate
	GetPeriodUnit() *string
	SetPrivateIpAddress(v string) *AddonNodeTemplate
	GetPrivateIpAddress() *string
	SetSpotPriceLimit(v float32) *AddonNodeTemplate
	GetSpotPriceLimit() *float32
	SetSpotStrategy(v string) *AddonNodeTemplate
	GetSpotStrategy() *string
	SetSystemDisk(v *AddonNodeTemplateSystemDisk) *AddonNodeTemplate
	GetSystemDisk() *AddonNodeTemplateSystemDisk
}

type AddonNodeTemplate struct {
	// Specifies whether to enable auto-renewal for the reserved instance. This parameter takes effect only when InstanceChargeType is set to PrePaid. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// Default value: false.
	//
	// example:
	//
	// false
	AutoRenew *bool `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// The auto-renewal period. Valid values:
	//
	// 	- Valid values when PeriodUnit is set to Week: 1, 2, and 3
	//
	// 	- Valid values when PeriodUnit is set to Month: 1, 2, 3, 6, 12, 24, 36, 48, and 60
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	AutoRenewPeriod *int32 `json:"AutoRenewPeriod,omitempty" xml:"AutoRenewPeriod,omitempty"`
	// The data disks.
	//
	// >  You can specify up to 16 data disks.
	DataDisks []*AddonNodeTemplateDataDisks `json:"DataDisks,omitempty" xml:"DataDisks,omitempty" type:"Repeated"`
	// The protection period for the preemptible instance. Unit: hours. Default value: 1. Valid values:
	//
	// 	- 1: After a preemptible instance is created, Alibaba Cloud ensures that the instance is not automatically released within one hour. After the one-hour protection period ends, the system compares the bid price with the market price and checks the resource inventory to determine whether to retain or release the instance.
	//
	// 	- 0: After a preemptible instance is created, Alibaba Cloud does not ensure that the instance runs for one hour. The system compares the bid price with the market price and checks the resource inventory to determine whether to retain or release the instance.
	//
	// Alibaba Cloud sends an ECS system event to notify you five minutes before the instance is released. Preemptible instances are billed by second. We recommend that you specify an appropriate protection period based on your business requirements.
	//
	// example:
	//
	// 1
	Duration *int32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// Specifies whether to enable hyper-threading for the node. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	EnableHT *bool `json:"EnableHT,omitempty" xml:"EnableHT,omitempty"`
	// The ID of the image to be used for instance booting. You can call the [DescribeImages](https://help.aliyun.com/zh/ecs/developer-reference/api-ecs-2014-05-26-describeimages?spm=api-workbench.API%20Document.0.0.7e5caef0GBcMYX) operation to query the available image resources.
	//
	// This parameter is required.
	//
	// example:
	//
	// aliyun_3_x64_20G_alibase_20221102.vhd
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The billing method. Valid values:
	//
	// 	- Prepaid: subscription
	//
	// 	- PostPaid: pay-as-you-go
	//
	// Default value: PostPaid.
	//
	// If you set this parameter to PrePaid, you must make sure that your account supports payment by balance or credit. Otherwise, the InvalidPayMethod error message will be returned.
	//
	// example:
	//
	// PostPaid
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" xml:"InstanceChargeType,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// i-bp1bzqq4rj1eemun****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ECS instance type.
	//
	// 	- To select an appropriate instance type, you can see [Instance families](https://help.aliyun.com/zh/ecs/user-guide/overview-of-instance-families?spm=api-workbench.API%20Document.0.0.7e5caef0GBcMYX) or call the [DescribeInstanceTypes](https://help.aliyun.com/zh/ecs/developer-reference/api-ecs-2014-05-26-describeinstancetypes?spm=api-workbench.API%20Document.0.0.7e5caef0GBcMYX) operation to learn the performance data about instance types.
	//
	// 	- To query the inventory of instance types in specified region or zone, you can call the [DescribeAvailableResource](https://help.aliyun.com/zh/ecs/developer-reference/api-ecs-2014-05-26-describeavailableresource?spm=api-workbench.API%20Document.0.0.7e5caef0GBcMYX) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// ecs.c7.4xlarge
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// OsName
	//
	// This parameter is required.
	//
	// example:
	//
	// CentOS  7.6 64 位
	OsName *string `json:"OsName,omitempty" xml:"OsName,omitempty"`
	// OsNameEN
	//
	// This parameter is required.
	//
	// example:
	//
	// CentOS  7.6 64 bit
	OsNameEN *string `json:"OsNameEN,omitempty" xml:"OsNameEN,omitempty"`
	// The subscription period of the instance. The unit is specified by the PeriodUnit parameter. This parameter takes effect and is required only when InstanceChargeType is set to PrePaid. If the DedicatedHostId parameter is specified, the subscription duration of the instance must be no longer than that of the dedicated host. Valid values:
	//
	// 	- Valid values when PeriodUnit is set to Week: 1, 2, 3, and 4
	//
	// 	- Valid values when PeriodUnit is set to Month: 1, 2, 3, 4, 5, 6, 7, 8, 9, 12, 24, 36, 48, and 60
	//
	// example:
	//
	// 1
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The unit of the subscription period. Valid values:
	//
	// 	- Week
	//
	// 	- Month (default)
	//
	// example:
	//
	// Month
	PeriodUnit *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	// The IP address of the virtual private cloud (VPC) in which the ECS instance is deployed.
	//
	// example:
	//
	// ``172.16.**.**``
	PrivateIpAddress *string `json:"PrivateIpAddress,omitempty" xml:"PrivateIpAddress,omitempty"`
	// The maximum hourly price of the preemptible instance. This parameter takes effect only when SpotStrategy is set to SpotWithPriceLimit. A maximum of three decimal places can be specified.
	//
	// example:
	//
	// 0.97
	SpotPriceLimit *float32 `json:"SpotPriceLimit,omitempty" xml:"SpotPriceLimit,omitempty"`
	// The bidding policy for the pay-as-you-go instance. This parameter is valid only when InstanceChargeType is set to PostPaid. Valid values:
	//
	// 	- NoSpot: The instance is created as a pay-as-you-go instance.
	//
	// 	- SpotWithPriceLimit: The instance is created as a preemptible instance with a user-defined maximum hourly price.
	//
	// 	- SpotAsPriceGo: The instances are created as preemptible instances for which the market price at the time of purchase is automatically used as the bidding price.
	//
	// Default value: NoSpot.
	//
	// example:
	//
	// NoSpot
	SpotStrategy *string `json:"SpotStrategy,omitempty" xml:"SpotStrategy,omitempty"`
	// The system disk configurations of the node.
	SystemDisk *AddonNodeTemplateSystemDisk `json:"SystemDisk,omitempty" xml:"SystemDisk,omitempty" type:"Struct"`
}

func (s AddonNodeTemplate) String() string {
	return dara.Prettify(s)
}

func (s AddonNodeTemplate) GoString() string {
	return s.String()
}

func (s *AddonNodeTemplate) GetAutoRenew() *bool {
	return s.AutoRenew
}

func (s *AddonNodeTemplate) GetAutoRenewPeriod() *int32 {
	return s.AutoRenewPeriod
}

func (s *AddonNodeTemplate) GetDataDisks() []*AddonNodeTemplateDataDisks {
	return s.DataDisks
}

func (s *AddonNodeTemplate) GetDuration() *int32 {
	return s.Duration
}

func (s *AddonNodeTemplate) GetEnableHT() *bool {
	return s.EnableHT
}

func (s *AddonNodeTemplate) GetImageId() *string {
	return s.ImageId
}

func (s *AddonNodeTemplate) GetInstanceChargeType() *string {
	return s.InstanceChargeType
}

func (s *AddonNodeTemplate) GetInstanceId() *string {
	return s.InstanceId
}

func (s *AddonNodeTemplate) GetInstanceType() *string {
	return s.InstanceType
}

func (s *AddonNodeTemplate) GetOsName() *string {
	return s.OsName
}

func (s *AddonNodeTemplate) GetOsNameEN() *string {
	return s.OsNameEN
}

func (s *AddonNodeTemplate) GetPeriod() *int32 {
	return s.Period
}

func (s *AddonNodeTemplate) GetPeriodUnit() *string {
	return s.PeriodUnit
}

func (s *AddonNodeTemplate) GetPrivateIpAddress() *string {
	return s.PrivateIpAddress
}

func (s *AddonNodeTemplate) GetSpotPriceLimit() *float32 {
	return s.SpotPriceLimit
}

func (s *AddonNodeTemplate) GetSpotStrategy() *string {
	return s.SpotStrategy
}

func (s *AddonNodeTemplate) GetSystemDisk() *AddonNodeTemplateSystemDisk {
	return s.SystemDisk
}

func (s *AddonNodeTemplate) SetAutoRenew(v bool) *AddonNodeTemplate {
	s.AutoRenew = &v
	return s
}

func (s *AddonNodeTemplate) SetAutoRenewPeriod(v int32) *AddonNodeTemplate {
	s.AutoRenewPeriod = &v
	return s
}

func (s *AddonNodeTemplate) SetDataDisks(v []*AddonNodeTemplateDataDisks) *AddonNodeTemplate {
	s.DataDisks = v
	return s
}

func (s *AddonNodeTemplate) SetDuration(v int32) *AddonNodeTemplate {
	s.Duration = &v
	return s
}

func (s *AddonNodeTemplate) SetEnableHT(v bool) *AddonNodeTemplate {
	s.EnableHT = &v
	return s
}

func (s *AddonNodeTemplate) SetImageId(v string) *AddonNodeTemplate {
	s.ImageId = &v
	return s
}

func (s *AddonNodeTemplate) SetInstanceChargeType(v string) *AddonNodeTemplate {
	s.InstanceChargeType = &v
	return s
}

func (s *AddonNodeTemplate) SetInstanceId(v string) *AddonNodeTemplate {
	s.InstanceId = &v
	return s
}

func (s *AddonNodeTemplate) SetInstanceType(v string) *AddonNodeTemplate {
	s.InstanceType = &v
	return s
}

func (s *AddonNodeTemplate) SetOsName(v string) *AddonNodeTemplate {
	s.OsName = &v
	return s
}

func (s *AddonNodeTemplate) SetOsNameEN(v string) *AddonNodeTemplate {
	s.OsNameEN = &v
	return s
}

func (s *AddonNodeTemplate) SetPeriod(v int32) *AddonNodeTemplate {
	s.Period = &v
	return s
}

func (s *AddonNodeTemplate) SetPeriodUnit(v string) *AddonNodeTemplate {
	s.PeriodUnit = &v
	return s
}

func (s *AddonNodeTemplate) SetPrivateIpAddress(v string) *AddonNodeTemplate {
	s.PrivateIpAddress = &v
	return s
}

func (s *AddonNodeTemplate) SetSpotPriceLimit(v float32) *AddonNodeTemplate {
	s.SpotPriceLimit = &v
	return s
}

func (s *AddonNodeTemplate) SetSpotStrategy(v string) *AddonNodeTemplate {
	s.SpotStrategy = &v
	return s
}

func (s *AddonNodeTemplate) SetSystemDisk(v *AddonNodeTemplateSystemDisk) *AddonNodeTemplate {
	s.SystemDisk = v
	return s
}

func (s *AddonNodeTemplate) Validate() error {
	if s.DataDisks != nil {
		for _, item := range s.DataDisks {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.SystemDisk != nil {
		if err := s.SystemDisk.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AddonNodeTemplateDataDisks struct {
	// The disk category. Valid values:
	//
	// 	- cloud_efficiency: utra disk
	//
	// 	- cloud_ssd: standard SSD
	//
	// 	- cloud_essd: ESSD
	//
	// example:
	//
	// cloud_essd
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// Specifies whether to release data disk N when the instance is released. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// Default value: true.
	//
	// example:
	//
	// false
	DeleteWithInstance *bool `json:"DeleteWithInstance,omitempty" xml:"DeleteWithInstance,omitempty"`
	// The performance level of the ESSD to be used as the data disk. Valid values:
	//
	// 	- PL0: A single ESSD can deliver up to 10000 random read/write IOPS.
	//
	// 	- PL1 (default): A single ESSD can deliver up to 50,000 random read/write IOPS.
	//
	// 	- PL2: A single ESSD can deliver up to 100000 random read/write IOPS.
	//
	// 	- PL3: A single ESSD can deliver up to 1000000 random read/write IOPS. For more information about ESSD performance levels, see [ESSDs](https://help.aliyun.com/zh/ecs/user-guide/essds?spm=api-workbench.API%20Document.0.0.7e5caef0GBcMYX).
	//
	// example:
	//
	// PL0
	Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
	// The disk size. Valid values of N: 1 to 16. Unit: GiB. Valid values:
	//
	// 	- cloud_efficiency: 40 to 32,768
	//
	// 	- cloud_ssd: 40 to 32,768
	//
	// 	- Valid values when Category is set to cloud_essd depends on the value of the DataDisk.N.PerformanceLevel parameter. Specifically:
	//
	//     	- PL0: 40 to 65,536
	//
	//     	- PL1: 40 to 65,536
	//
	//     	- PL2: 461 to 65,536
	//
	//     	- PL3: 1,261 to 65,536
	//
	// example:
	//
	// 40
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s AddonNodeTemplateDataDisks) String() string {
	return dara.Prettify(s)
}

func (s AddonNodeTemplateDataDisks) GoString() string {
	return s.String()
}

func (s *AddonNodeTemplateDataDisks) GetCategory() *string {
	return s.Category
}

func (s *AddonNodeTemplateDataDisks) GetDeleteWithInstance() *bool {
	return s.DeleteWithInstance
}

func (s *AddonNodeTemplateDataDisks) GetLevel() *string {
	return s.Level
}

func (s *AddonNodeTemplateDataDisks) GetSize() *int32 {
	return s.Size
}

func (s *AddonNodeTemplateDataDisks) SetCategory(v string) *AddonNodeTemplateDataDisks {
	s.Category = &v
	return s
}

func (s *AddonNodeTemplateDataDisks) SetDeleteWithInstance(v bool) *AddonNodeTemplateDataDisks {
	s.DeleteWithInstance = &v
	return s
}

func (s *AddonNodeTemplateDataDisks) SetLevel(v string) *AddonNodeTemplateDataDisks {
	s.Level = &v
	return s
}

func (s *AddonNodeTemplateDataDisks) SetSize(v int32) *AddonNodeTemplateDataDisks {
	s.Size = &v
	return s
}

func (s *AddonNodeTemplateDataDisks) Validate() error {
	return dara.Validate(s)
}

type AddonNodeTemplateSystemDisk struct {
	// The category of the system disk. Valid values:
	//
	// 	- cloud_efficiency: utra disk
	//
	// 	- cloud_ssd: standard SSD
	//
	// 	- cloud_essd: enterprise SSD (ESSD)
	//
	// 	- cloud: basic disk
	//
	// example:
	//
	// cloud_essd
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The performance level of the ESSD that is used as the system disk. Valid values:
	//
	// 	- PL0: A single ESSD can deliver up to 10000 random read/write IOPS.
	//
	// 	- PL1 (default): A single ESSD can deliver up to 50,000 random read/write IOPS.
	//
	// 	- PL2: A single ESSD can deliver up to 100000 random read/write IOPS.
	//
	// 	- PL3: A single ESSD can deliver up to 1000000 random read/write IOPS. For more information about ESSD performance levels, see [ESSDs](https://help.aliyun.com/zh/ecs/user-guide/essds?spm=api-workbench.API%20Document.0.0.7e5caef0GBcMYX).
	//
	// example:
	//
	// PL0
	Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
	// The disk size. Unit: GiB. Valid values:
	//
	// 	- cloud_efficiency: 40 to 32,768
	//
	// 	- cloud_ssd: 40 to 32,768
	//
	// 	- Valid values when Category is set to cloud_essd depends on the value of the DataDisk.N.PerformanceLevel parameter. Specifically:
	//
	//     	- PL0: 40 to 65,536
	//
	//     	- PL1: 40 to 65,536
	//
	//     	- PL2: 461 to 65,536
	//
	//     	- PL3: 1,261 to 65,536
	//
	// 	- cloud: 40 to 500
	//
	// example:
	//
	// 40
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s AddonNodeTemplateSystemDisk) String() string {
	return dara.Prettify(s)
}

func (s AddonNodeTemplateSystemDisk) GoString() string {
	return s.String()
}

func (s *AddonNodeTemplateSystemDisk) GetCategory() *string {
	return s.Category
}

func (s *AddonNodeTemplateSystemDisk) GetLevel() *string {
	return s.Level
}

func (s *AddonNodeTemplateSystemDisk) GetSize() *int32 {
	return s.Size
}

func (s *AddonNodeTemplateSystemDisk) SetCategory(v string) *AddonNodeTemplateSystemDisk {
	s.Category = &v
	return s
}

func (s *AddonNodeTemplateSystemDisk) SetLevel(v string) *AddonNodeTemplateSystemDisk {
	s.Level = &v
	return s
}

func (s *AddonNodeTemplateSystemDisk) SetSize(v int32) *AddonNodeTemplateSystemDisk {
	s.Size = &v
	return s
}

func (s *AddonNodeTemplateSystemDisk) Validate() error {
	return dara.Validate(s)
}
