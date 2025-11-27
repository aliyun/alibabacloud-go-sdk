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
	// 	- **true*	- (default): enables automatic payment. Make sure that your account balance is sufficient.
	//
	// 	- **false**: does not automatically complete the payment. An unpaid order is generated.
	//
	// >  If your account balance is insufficient, you can set the parameter to false. In this case, an unpaid order is generated. You can complete the payment in the Expenses and Costs console.
	//
	// example:
	//
	// true
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// Specifies whether to enable auto-renewal. You must specify this parameter only when the data disk uses the subscription billing method. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// >  The auto-renewal cycle is one month for a monthly subscription. The auto-renewal cycle is one year for a yearly subscription.
	//
	// example:
	//
	// false
	AutoRenew *bool `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// The disk description. The description must be 2 to 256 characters in length and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The data disk type. Valid values:
	//
	// 	- **cloud_efficiency**: ultra disk.
	//
	// 	- **cloud_ssd**: standard SSD
	//
	// 	- **cloud_essd**: ESSD
	//
	// 	- **cloud_auto*	- (default): Premium ESSD
	//
	// example:
	//
	// cloud_ssd
	DiskCategory *string `json:"DiskCategory,omitempty" xml:"DiskCategory,omitempty"`
	// The name of the data disk. The name must be 2 to 128 characters in length and can contain letters and digits. The name can contain colons (:), underscores (_), periods (.), and hyphens (-).
	//
	// example:
	//
	// ZStack-Hybrid-Test-ECS-Instance
	DiskName *string `json:"DiskName,omitempty" xml:"DiskName,omitempty"`
	// The billing method. Valid values:
	//
	// 	- **Postpaid**: pay-as-you-go Pay-as-you-go disks do not require to be attached. You can also attach the pay-as-you-go disk to an instance of any billing method based on your business requirements.
	//
	// 	- **Prepaid**: subscription Subscription disks must be attached to a subscription instance. Set **InstanceId*	- to the ID of a subscription instance.
	//
	// example:
	//
	// Postpaid
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" xml:"InstanceChargeType,omitempty"`
	// The ID of the instance to which you want to attach the disk. If you set **InstanceChargeType*	- to **Prepaid**, you must set InstanceId to the ID of a subscription instance.
	//
	// example:
	//
	// rc-v28c6k3jupp61m2t****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The performance level (PL) of ESSDs. Valid values:
	//
	// 	- **PL0**: A single ESSD delivers up to 10,000 random read/write IOPS.
	//
	// 	- **PL1: An ESSD delivers up to 50,000 random read/write IOPS.**
	//
	// 	- **PL2**: A single ESSD delivers up to 100,000 random read/write IOPS.
	//
	// 	- **PL3**: A single ESSD delivers up to 1,000,000 random read/write IOPS.
	//
	// For information about ESSD PLs, see [ESSDs](https://help.aliyun.com/document_detail/2859916.html).
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
	// The region ID. You can call the DescribeRegions operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-ac****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The disk size. Unit: GiB. This parameter is required. Valid values:
	//
	// 	- Valid values if you set DiskCategory to **cloud_efficiency**: 20 to 32768.
	//
	// 	- Valid values if you set DiskCategory to **cloud_ssd**: 20 to 32768.
	//
	// 	- Valid values if you set DiskCategory to **cloud_auto**: 1 to 65536.
	//
	// 	- Valid values when DiskCategory is set to cloud_essd: depending on the value of **PerformanceLevel**.****
	//
	//     	- Valid values if PerformanceLevel is set to PL0: 1 to 65536
	//
	//     	- Valid values if PerformanceLevel is set to PL1: 20 to 65536
	//
	//     	- Valid values if PerformanceLevel is set to PL2: 461 to 65536
	//
	//     	- Valid values if PerformanceLevel is set to PL3: 1261 to 65536
	//
	// If **SnapshotId*	- is specified and the size of the corresponding snapshot is greater than the **Size*	- value, the size of the created disk is the same as that of the snapshot. If the snapshot size is less than the **Size*	- value, the size of the created disk is equal to the **Size*	- value.
	//
	// example:
	//
	// 2000
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
	// The snapshot that you want to use to create the disk.
	//
	// 	- The snapshots of RDS Custom instances and the non-shared snapshots of ECS instances are supported.
	//
	// 	- If the size of the snapshot specified by **SnapshotId*	- is greater than the value of **Size**, the size of the created disk is equal to the specified snapshot size. If the snapshot size is less than the **Size*	- value, the size of the created disk is equal to the **Size*	- value.
	//
	// 	- You cannot create elastic ephemeral disks from snapshots.
	//
	// 	- Snapshots that were created on or before July 15, 2013 cannot be used to create disks.
	//
	// example:
	//
	// rcds-umtnkvevqbu****
	SnapshotId *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
	// The list of tags.
	Tag []*CreateRCDiskRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The zone ID.
	//
	// This parameter is required if you do not specify **InstanceId**.
	//
	// example:
	//
	// cn-hangzhou-a
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
	// The tag key. You can create N tag keys at a time. Valid values of N: **1 to 20**. The tag key cannot be an empty string.
	//
	// example:
	//
	// testkey1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value. You can query N values at a time. Valid values of N: **1*	- to **20**. The tag value can be an empty string.
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
