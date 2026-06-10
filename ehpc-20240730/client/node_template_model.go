// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iNodeTemplate interface {
	dara.Model
	String() string
	GoString() string
	SetAutoRenew(v bool) *NodeTemplate
	GetAutoRenew() *bool
	SetAutoRenewPeriod(v int32) *NodeTemplate
	GetAutoRenewPeriod() *int32
	SetDataDisks(v []*NodeTemplateDataDisks) *NodeTemplate
	GetDataDisks() []*NodeTemplateDataDisks
	SetDuration(v int32) *NodeTemplate
	GetDuration() *int32
	SetEnableHT(v bool) *NodeTemplate
	GetEnableHT() *bool
	SetImageId(v string) *NodeTemplate
	GetImageId() *string
	SetInstanceChargeType(v string) *NodeTemplate
	GetInstanceChargeType() *string
	SetInstanceType(v string) *NodeTemplate
	GetInstanceType() *string
	SetPeriod(v int32) *NodeTemplate
	GetPeriod() *int32
	SetPeriodUnit(v string) *NodeTemplate
	GetPeriodUnit() *string
	SetSpotPriceLimit(v float32) *NodeTemplate
	GetSpotPriceLimit() *float32
	SetSpotStrategy(v string) *NodeTemplate
	GetSpotStrategy() *string
	SetSystemDisk(v *NodeTemplateSystemDisk) *NodeTemplate
	GetSystemDisk() *NodeTemplateSystemDisk
}

type NodeTemplate struct {
	// Specifies whether to enable auto-renewal. This parameter takes effect only when InstanceChargeType is set to PrePaid. Valid values:
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
	DataDisks []*NodeTemplateDataDisks `json:"DataDisks,omitempty" xml:"DataDisks,omitempty" type:"Repeated"`
	// The protection period for the preemptible instance. Unit: hours. Default value: 1. Valid values:
	//
	// 	- 1: After a preemptible instance is created, Alibaba Cloud ensures that the instance is not automatically released within one hour. After the one-hour protection period ends, the system compares the bid price with the market price and checks the resource inventory to determine whether to retain or release the instance.
	//
	// 	- 0: After a preemptible instance is created, Alibaba Cloud does not ensure that the instance runs for one hour. The system compares the bid price with the market price and checks the resource inventory to determine whether to retain or release the instance.
	//
	// Alibaba Cloud sends an ECS system event to notify you five minutes before the instance is released. The preemptible instance is billed by second. We recommend that you specify an appropriate protection period based on your business requirements.
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
	// example:
	//
	// aliyun_3_x64_20G_alibase_20221102.vhd
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The billing method. Valid values:
	//
	// 	- PrePaid: subscription
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
	// The ECS instance type.
	//
	// 	- To select an appropriate instance type, you can see [Instance families](https://help.aliyun.com/zh/ecs/user-guide/overview-of-instance-families?spm=api-workbench.API%20Document.0.0.7e5caef0GBcMYX) or call the [DescribeInstanceTypes](https://help.aliyun.com/zh/ecs/developer-reference/api-ecs-2014-05-26-describeinstancetypes?spm=api-workbench.API%20Document.0.0.7e5caef0GBcMYX) operation to learn the performance data about instance types.
	//
	// 	- To query the inventory of instance types in specified region or zone, you can call the [DescribeAvailableResource](https://help.aliyun.com/zh/ecs/developer-reference/api-ecs-2014-05-26-describeavailableresource?spm=api-workbench.API%20Document.0.0.7e5caef0GBcMYX) operation.
	//
	// example:
	//
	// ecs.c7.4xlarge
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
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
	// 	- SpotAsPriceGo: The instance is created as a preemptible instance for which the market price at the time of purchase is automatically used as the bidding price.
	//
	// Default value: NoSpot.
	//
	// example:
	//
	// NoSpot
	SpotStrategy *string `json:"SpotStrategy,omitempty" xml:"SpotStrategy,omitempty"`
	// The system disk configurations of the node.
	SystemDisk *NodeTemplateSystemDisk `json:"SystemDisk,omitempty" xml:"SystemDisk,omitempty" type:"Struct"`
}

func (s NodeTemplate) String() string {
	return dara.Prettify(s)
}

func (s NodeTemplate) GoString() string {
	return s.String()
}

func (s *NodeTemplate) GetAutoRenew() *bool {
	return s.AutoRenew
}

func (s *NodeTemplate) GetAutoRenewPeriod() *int32 {
	return s.AutoRenewPeriod
}

func (s *NodeTemplate) GetDataDisks() []*NodeTemplateDataDisks {
	return s.DataDisks
}

func (s *NodeTemplate) GetDuration() *int32 {
	return s.Duration
}

func (s *NodeTemplate) GetEnableHT() *bool {
	return s.EnableHT
}

func (s *NodeTemplate) GetImageId() *string {
	return s.ImageId
}

func (s *NodeTemplate) GetInstanceChargeType() *string {
	return s.InstanceChargeType
}

func (s *NodeTemplate) GetInstanceType() *string {
	return s.InstanceType
}

func (s *NodeTemplate) GetPeriod() *int32 {
	return s.Period
}

func (s *NodeTemplate) GetPeriodUnit() *string {
	return s.PeriodUnit
}

func (s *NodeTemplate) GetSpotPriceLimit() *float32 {
	return s.SpotPriceLimit
}

func (s *NodeTemplate) GetSpotStrategy() *string {
	return s.SpotStrategy
}

func (s *NodeTemplate) GetSystemDisk() *NodeTemplateSystemDisk {
	return s.SystemDisk
}

func (s *NodeTemplate) SetAutoRenew(v bool) *NodeTemplate {
	s.AutoRenew = &v
	return s
}

func (s *NodeTemplate) SetAutoRenewPeriod(v int32) *NodeTemplate {
	s.AutoRenewPeriod = &v
	return s
}

func (s *NodeTemplate) SetDataDisks(v []*NodeTemplateDataDisks) *NodeTemplate {
	s.DataDisks = v
	return s
}

func (s *NodeTemplate) SetDuration(v int32) *NodeTemplate {
	s.Duration = &v
	return s
}

func (s *NodeTemplate) SetEnableHT(v bool) *NodeTemplate {
	s.EnableHT = &v
	return s
}

func (s *NodeTemplate) SetImageId(v string) *NodeTemplate {
	s.ImageId = &v
	return s
}

func (s *NodeTemplate) SetInstanceChargeType(v string) *NodeTemplate {
	s.InstanceChargeType = &v
	return s
}

func (s *NodeTemplate) SetInstanceType(v string) *NodeTemplate {
	s.InstanceType = &v
	return s
}

func (s *NodeTemplate) SetPeriod(v int32) *NodeTemplate {
	s.Period = &v
	return s
}

func (s *NodeTemplate) SetPeriodUnit(v string) *NodeTemplate {
	s.PeriodUnit = &v
	return s
}

func (s *NodeTemplate) SetSpotPriceLimit(v float32) *NodeTemplate {
	s.SpotPriceLimit = &v
	return s
}

func (s *NodeTemplate) SetSpotStrategy(v string) *NodeTemplate {
	s.SpotStrategy = &v
	return s
}

func (s *NodeTemplate) SetSystemDisk(v *NodeTemplateSystemDisk) *NodeTemplate {
	s.SystemDisk = v
	return s
}

func (s *NodeTemplate) Validate() error {
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

type NodeTemplateDataDisks struct {
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
	// Device
	//
	// example:
	//
	// /dev/xvdb
	Device *string `json:"Device,omitempty" xml:"Device,omitempty"`
	// The performance level of the ESSD that is used as a data disk. Valid values:
	//
	// 	- PL0: A single ESSD can deliver up to 10,000 random read/write IOPS.
	//
	// 	- PL1 (default): A single ESSD can deliver up to 50,000 random read/write IOPS.
	//
	// 	- PL2: A single ESSD can deliver up to 100,000 random read/write IOPS.
	//
	// 	- PL3: A single ESSD can deliver up to 1,000,000 random read/write IOPS. For more information about ESSD performance levels, see [ESSDs](https://help.aliyun.com/zh/ecs/user-guide/essds?spm=api-workbench.API%20Document.0.0.7e5caef0GBcMYX).
	//
	// example:
	//
	// PL0
	Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
	// MountDir
	//
	// example:
	//
	// /data1
	MountDir *string `json:"MountDir,omitempty" xml:"MountDir,omitempty"`
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
	// SnapshotId
	//
	// example:
	//
	// s-bp1ei2b44ripxuo46hym
	SnapshotId *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
}

func (s NodeTemplateDataDisks) String() string {
	return dara.Prettify(s)
}

func (s NodeTemplateDataDisks) GoString() string {
	return s.String()
}

func (s *NodeTemplateDataDisks) GetCategory() *string {
	return s.Category
}

func (s *NodeTemplateDataDisks) GetDeleteWithInstance() *bool {
	return s.DeleteWithInstance
}

func (s *NodeTemplateDataDisks) GetDevice() *string {
	return s.Device
}

func (s *NodeTemplateDataDisks) GetLevel() *string {
	return s.Level
}

func (s *NodeTemplateDataDisks) GetMountDir() *string {
	return s.MountDir
}

func (s *NodeTemplateDataDisks) GetSize() *int32 {
	return s.Size
}

func (s *NodeTemplateDataDisks) GetSnapshotId() *string {
	return s.SnapshotId
}

func (s *NodeTemplateDataDisks) SetCategory(v string) *NodeTemplateDataDisks {
	s.Category = &v
	return s
}

func (s *NodeTemplateDataDisks) SetDeleteWithInstance(v bool) *NodeTemplateDataDisks {
	s.DeleteWithInstance = &v
	return s
}

func (s *NodeTemplateDataDisks) SetDevice(v string) *NodeTemplateDataDisks {
	s.Device = &v
	return s
}

func (s *NodeTemplateDataDisks) SetLevel(v string) *NodeTemplateDataDisks {
	s.Level = &v
	return s
}

func (s *NodeTemplateDataDisks) SetMountDir(v string) *NodeTemplateDataDisks {
	s.MountDir = &v
	return s
}

func (s *NodeTemplateDataDisks) SetSize(v int32) *NodeTemplateDataDisks {
	s.Size = &v
	return s
}

func (s *NodeTemplateDataDisks) SetSnapshotId(v string) *NodeTemplateDataDisks {
	s.SnapshotId = &v
	return s
}

func (s *NodeTemplateDataDisks) Validate() error {
	return dara.Validate(s)
}

type NodeTemplateSystemDisk struct {
	// The category of the system disk. Valid values:
	//
	// 	- cloud_efficiency: utra disk
	//
	// 	- cloud_ssd: standard SSD
	//
	// 	- cloud_essd: Enterprise SSD (ESSD)
	//
	// 	- cloud: basic disk
	//
	// example:
	//
	// cloud_essd
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The performance level of the ESSD that is used as the system disk. Valid values:
	//
	// 	- PL0: A single ESSD can deliver up to 10,000 random read/write IOPS.
	//
	// 	- PL1 (default): A single ESSD can deliver up to 50,000 random read/write IOPS.
	//
	// 	- PL2: A single ESSD can deliver up to 100,000 random read/write IOPS.
	//
	// 	- PL3: A single ESSD can deliver up to 1,000,000 random read/write IOPS. For more information about ESSD performance levels, see [ESSDs](https://help.aliyun.com/zh/ecs/user-guide/essds?spm=api-workbench.API%20Document.0.0.7e5caef0GBcMYX).
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

func (s NodeTemplateSystemDisk) String() string {
	return dara.Prettify(s)
}

func (s NodeTemplateSystemDisk) GoString() string {
	return s.String()
}

func (s *NodeTemplateSystemDisk) GetCategory() *string {
	return s.Category
}

func (s *NodeTemplateSystemDisk) GetLevel() *string {
	return s.Level
}

func (s *NodeTemplateSystemDisk) GetSize() *int32 {
	return s.Size
}

func (s *NodeTemplateSystemDisk) SetCategory(v string) *NodeTemplateSystemDisk {
	s.Category = &v
	return s
}

func (s *NodeTemplateSystemDisk) SetLevel(v string) *NodeTemplateSystemDisk {
	s.Level = &v
	return s
}

func (s *NodeTemplateSystemDisk) SetSize(v int32) *NodeTemplateSystemDisk {
	s.Size = &v
	return s
}

func (s *NodeTemplateSystemDisk) Validate() error {
	return dara.Validate(s)
}
