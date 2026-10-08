// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRCDiskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoPay(v bool) *CreateRCDiskRequest
	GetAutoPay() *bool
	SetAutoRenew(v bool) *CreateRCDiskRequest
	GetAutoRenew() *bool
	SetDescription(v string) *CreateRCDiskRequest
	GetDescription() *string
	SetDiskCategory(v string) *CreateRCDiskRequest
	GetDiskCategory() *string
	SetDiskName(v string) *CreateRCDiskRequest
	GetDiskName() *string
	SetInstanceChargeType(v string) *CreateRCDiskRequest
	GetInstanceChargeType() *string
	SetInstanceId(v string) *CreateRCDiskRequest
	GetInstanceId() *string
	SetPerformanceLevel(v string) *CreateRCDiskRequest
	GetPerformanceLevel() *string
	SetPeriod(v int32) *CreateRCDiskRequest
	GetPeriod() *int32
	SetPeriodUnit(v string) *CreateRCDiskRequest
	GetPeriodUnit() *string
	SetRegionId(v string) *CreateRCDiskRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreateRCDiskRequest
	GetResourceGroupId() *string
	SetSize(v int32) *CreateRCDiskRequest
	GetSize() *int32
	SetSnapshotId(v string) *CreateRCDiskRequest
	GetSnapshotId() *string
	SetTag(v []*CreateRCDiskRequestTag) *CreateRCDiskRequest
	GetTag() []*CreateRCDiskRequestTag
	SetZoneId(v string) *CreateRCDiskRequest
	GetZoneId() *string
}

type CreateRCDiskRequest struct {
	// Specifies whether to enable automatic payment. Valid values:
	//
	// - **true*	- (default): enables automatic payment. Make sure that your account balance is sufficient.
	//
	// - **false**: generates an order without charging.
	//
	//
	//
	//
	// > If your payment method has insufficient balance, set this parameter to false. An unpaid order is generated, and you can log on to the ApsaraDB RDS console to complete the payment.
	//
	// >
	//
	// example:
	//
	// true
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// Specifies whether to enable auto-renewal. This parameter is valid only when you create a subscription data cloud disk. Valid values:
	//
	// - **true**: enables auto-renewal.
	//
	// - **false**: disables auto-renewal.
	//
	//  > If you purchase the cloud disk on a monthly basis, the auto-renewal epoch is one month.
	//
	//  If you purchase the cloud disk on a yearly basis, the auto-renewal epoch is one year.
	//
	// example:
	//
	// false
	AutoRenew *bool `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// The description of the cloud disk. The description must be 2 to 256 characters in length and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The category of the data cloud disk. Valid values:
	//
	// - **cloud_efficiency**: ultra cloud disk.
	//
	// - **cloud_ssd**: standard SSD.
	//
	// - **cloud_essd**: ESSD.
	//
	// - **cloud_auto*	- (default): premium performance disk.
	//
	// example:
	//
	// cloud_auto
	DiskCategory *string `json:"DiskCategory,omitempty" xml:"DiskCategory,omitempty"`
	// The name of the cloud disk. The name must be 2 to 128 characters in length and can contain characters that are categorized as letter in Unicode, including Chinese characters, English letters, and digits. The name can also contain colons (:), underscores (_), periods (.), and hyphens (-).
	//
	// example:
	//
	// testDisk
	DiskName *string `json:"DiskName,omitempty" xml:"DiskName,omitempty"`
	// The billing method. Valid values:
	//
	// - **Postpaid**: pay-as-you-go. Cloud disks with this billing method do not need to be mounted to an instance. You can also mount them to an instance of any billing method during creation as needed.
	//
	// - **Prepaid**: subscription. Cloud disks with this billing method must be mounted to a subscription instance. You must specify the **InstanceId*	- (instance ID) of a subscription instance.
	//
	// example:
	//
	// Postpaid
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" xml:"InstanceChargeType,omitempty"`
	// Instance ID of the instance to which the cloud disk is attached. If **InstanceChargeType*	- is set to **Prepaid*	- (subscription), you must specify instance ID of a subscription instance.
	//
	// example:
	//
	// rc-v28c6k3jupp61m2t****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The performance level (PL) of the ESSD cloud disk. Valid values:
	//
	// - **PL0**: A single cloud disk can deliver up to 10,000 random read/write IOPS.
	//
	// - **PL1*	- (default): A single cloud disk can deliver up to 50,000 random read/write IOPS.
	//
	// - **PL2**: A single cloud disk can deliver up to 100,000 random read/write IOPS.
	//
	// - **PL3**: A single cloud disk can deliver up to 1,000,000 random read/write IOPS.
	//
	// For more information about how to select an ESSD performance level, see [ESSD cloud disk](https://help.aliyun.com/document_detail/2859916.html).
	//
	// example:
	//
	// PL1
	PerformanceLevel *string `json:"PerformanceLevel,omitempty" xml:"PerformanceLevel,omitempty"`
	// A reserved parameter. You do not need to specify this parameter.
	//
	// example:
	//
	// none
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// A reserved parameter. You do not need to specify this parameter.
	//
	// example:
	//
	// none
	PeriodUnit *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	// The region ID. You can call the DescribeRegions operation to query region IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-ac****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The capacity size. Unit: GiB. You must specify a value for this parameter. Valid values:
	//
	// - **cloud_efficiency**: 20 to 32,768.
	//
	// - **cloud_ssd**: 20 to 32,768.
	//
	// - **cloud_auto**: 1 to 65,536.
	//
	// - **cloud_essd**: The valid value range depends on the value of **PerformanceLevel**.
	//
	//   - PL0: 1 to 65,536.
	//
	//   - PL1: 20 to 65,536.
	//
	//   - PL2: 461 to 65,536.
	//
	//   - PL3: 1,261 to 65,536.
	//
	// If **SnapshotId*	- is specified and the capacity of the corresponding snapshot is greater than the value of **Size**, snapshot size of the created cloud disk is the same as the snapshot capacity. If the snapshot capacity is less than the value of **Size**, snapshot size of the created cloud disk is the value of **Size**.
	//
	// example:
	//
	// 2000
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
	// The snapshot that is used to create the cloud disk.
	//
	// - RDS Custom snapshots and ECS snapshots (non-shared type) are supported.
	//
	// - If the capacity of the snapshot specified by **SnapshotId*	- is greater than the value of **Size**, snapshot size of the created cloud disk is the same as the snapshot capacity. If the snapshot capacity is less than the value of **Size**, snapshot size of the created cloud disk is the value of **Size**.
	//
	// - Creating elastic ephemeral disks from snapshots is not supported.
	//
	// - Snapshots created on or before July 15, 2013 cannot be used to create cloud disks.
	//
	// example:
	//
	// rcds-umtnkvevqbu****
	SnapshotId *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
	// The tags.
	Tag []*CreateRCDiskRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The zone ID.
	//
	// This parameter is required if the **InstanceId*	- parameter (the instance ID of the instance to which the cloud disk is mounted) is not specified.
	//
	// example:
	//
	// cn-hangzhou-h
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s CreateRCDiskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateRCDiskRequest) GoString() string {
	return s.String()
}

func (s *CreateRCDiskRequest) GetAutoPay() *bool {
	return s.AutoPay
}

func (s *CreateRCDiskRequest) GetAutoRenew() *bool {
	return s.AutoRenew
}

func (s *CreateRCDiskRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateRCDiskRequest) GetDiskCategory() *string {
	return s.DiskCategory
}

func (s *CreateRCDiskRequest) GetDiskName() *string {
	return s.DiskName
}

func (s *CreateRCDiskRequest) GetInstanceChargeType() *string {
	return s.InstanceChargeType
}

func (s *CreateRCDiskRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateRCDiskRequest) GetPerformanceLevel() *string {
	return s.PerformanceLevel
}

func (s *CreateRCDiskRequest) GetPeriod() *int32 {
	return s.Period
}

func (s *CreateRCDiskRequest) GetPeriodUnit() *string {
	return s.PeriodUnit
}

func (s *CreateRCDiskRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateRCDiskRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateRCDiskRequest) GetSize() *int32 {
	return s.Size
}

func (s *CreateRCDiskRequest) GetSnapshotId() *string {
	return s.SnapshotId
}

func (s *CreateRCDiskRequest) GetTag() []*CreateRCDiskRequestTag {
	return s.Tag
}

func (s *CreateRCDiskRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *CreateRCDiskRequest) SetAutoPay(v bool) *CreateRCDiskRequest {
	s.AutoPay = &v
	return s
}

func (s *CreateRCDiskRequest) SetAutoRenew(v bool) *CreateRCDiskRequest {
	s.AutoRenew = &v
	return s
}

func (s *CreateRCDiskRequest) SetDescription(v string) *CreateRCDiskRequest {
	s.Description = &v
	return s
}

func (s *CreateRCDiskRequest) SetDiskCategory(v string) *CreateRCDiskRequest {
	s.DiskCategory = &v
	return s
}

func (s *CreateRCDiskRequest) SetDiskName(v string) *CreateRCDiskRequest {
	s.DiskName = &v
	return s
}

func (s *CreateRCDiskRequest) SetInstanceChargeType(v string) *CreateRCDiskRequest {
	s.InstanceChargeType = &v
	return s
}

func (s *CreateRCDiskRequest) SetInstanceId(v string) *CreateRCDiskRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateRCDiskRequest) SetPerformanceLevel(v string) *CreateRCDiskRequest {
	s.PerformanceLevel = &v
	return s
}

func (s *CreateRCDiskRequest) SetPeriod(v int32) *CreateRCDiskRequest {
	s.Period = &v
	return s
}

func (s *CreateRCDiskRequest) SetPeriodUnit(v string) *CreateRCDiskRequest {
	s.PeriodUnit = &v
	return s
}

func (s *CreateRCDiskRequest) SetRegionId(v string) *CreateRCDiskRequest {
	s.RegionId = &v
	return s
}

func (s *CreateRCDiskRequest) SetResourceGroupId(v string) *CreateRCDiskRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateRCDiskRequest) SetSize(v int32) *CreateRCDiskRequest {
	s.Size = &v
	return s
}

func (s *CreateRCDiskRequest) SetSnapshotId(v string) *CreateRCDiskRequest {
	s.SnapshotId = &v
	return s
}

func (s *CreateRCDiskRequest) SetTag(v []*CreateRCDiskRequestTag) *CreateRCDiskRequest {
	s.Tag = v
	return s
}

func (s *CreateRCDiskRequest) SetZoneId(v string) *CreateRCDiskRequest {
	s.ZoneId = &v
	return s
}

func (s *CreateRCDiskRequest) Validate() error {
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

type CreateRCDiskRequestTag struct {
	// The tag key. You can specify up to N tag keys at a time. Valid values of N: **1 to 20**. The tag key cannot be an empty string.
	//
	// example:
	//
	// testkey1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value that corresponds to the tag key. You can specify up to N tag values at a time. Valid values of N: **1*	- to **20**. The tag value can be an empty string.
	//
	// example:
	//
	// testvalue1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateRCDiskRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateRCDiskRequestTag) GoString() string {
	return s.String()
}

func (s *CreateRCDiskRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateRCDiskRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateRCDiskRequestTag) SetKey(v string) *CreateRCDiskRequestTag {
	s.Key = &v
	return s
}

func (s *CreateRCDiskRequestTag) SetValue(v string) *CreateRCDiskRequestTag {
	s.Value = &v
	return s
}

func (s *CreateRCDiskRequestTag) Validate() error {
	return dara.Validate(s)
}
