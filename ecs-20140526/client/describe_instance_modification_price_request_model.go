// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceModificationPriceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSystemDisk(v *DescribeInstanceModificationPriceRequestSystemDisk) *DescribeInstanceModificationPriceRequest
	GetSystemDisk() *DescribeInstanceModificationPriceRequestSystemDisk
	SetDataDisk(v []*DescribeInstanceModificationPriceRequestDataDisk) *DescribeInstanceModificationPriceRequest
	GetDataDisk() []*DescribeInstanceModificationPriceRequestDataDisk
	SetEndTime(v string) *DescribeInstanceModificationPriceRequest
	GetEndTime() *string
	SetISP(v string) *DescribeInstanceModificationPriceRequest
	GetISP() *string
	SetImageId(v string) *DescribeInstanceModificationPriceRequest
	GetImageId() *string
	SetInstanceId(v string) *DescribeInstanceModificationPriceRequest
	GetInstanceId() *string
	SetInstanceType(v string) *DescribeInstanceModificationPriceRequest
	GetInstanceType() *string
	SetInternetChargeType(v string) *DescribeInstanceModificationPriceRequest
	GetInternetChargeType() *string
	SetInternetMaxBandwidthOut(v int32) *DescribeInstanceModificationPriceRequest
	GetInternetMaxBandwidthOut() *int32
	SetOwnerAccount(v string) *DescribeInstanceModificationPriceRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeInstanceModificationPriceRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeInstanceModificationPriceRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeInstanceModificationPriceRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeInstanceModificationPriceRequest
	GetResourceOwnerId() *int64
	SetStartTime(v string) *DescribeInstanceModificationPriceRequest
	GetStartTime() *string
}

type DescribeInstanceModificationPriceRequest struct {
	SystemDisk *DescribeInstanceModificationPriceRequestSystemDisk `json:"SystemDisk,omitempty" xml:"SystemDisk,omitempty" type:"Struct"`
	// The information about data disk types.
	DataDisk []*DescribeInstanceModificationPriceRequestDataDisk `json:"DataDisk,omitempty" xml:"DataDisk,omitempty" type:"Repeated"`
	// example:
	//
	// 2025-12-06T22Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// BGP
	ISP *string `json:"ISP,omitempty" xml:"ISP,omitempty"`
	// example:
	//
	// aliyun_2_1903_x64_20G_alibase_20200324.vhd
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The instance ID of the instance for which you want to query the upgrade price.
	//
	// This parameter is required.
	//
	// example:
	//
	// i-bp1f2o4ldh8l****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The target instance type for the upgrade. Call [DescribeResourcesModification](https://help.aliyun.com/document_detail/66187.html) to query the instance types available for upgrade in a specified zone.
	//
	// > The instance type parameter (`InstanceType`) and data disk parameters (`DataDisk.N.*`) cannot both be empty. You must specify at least one.
	//
	// example:
	//
	// ecs.g6e.large
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// example:
	//
	// PayByTraffic
	InternetChargeType *string `json:"InternetChargeType,omitempty" xml:"InternetChargeType,omitempty"`
	// example:
	//
	// 10
	InternetMaxBandwidthOut *int32  `json:"InternetMaxBandwidthOut,omitempty" xml:"InternetMaxBandwidthOut,omitempty"`
	OwnerAccount            *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId                 *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID. You can call [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// example:
	//
	// 2025-12-05T22:40Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeInstanceModificationPriceRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceModificationPriceRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceModificationPriceRequest) GetSystemDisk() *DescribeInstanceModificationPriceRequestSystemDisk {
	return s.SystemDisk
}

func (s *DescribeInstanceModificationPriceRequest) GetDataDisk() []*DescribeInstanceModificationPriceRequestDataDisk {
	return s.DataDisk
}

func (s *DescribeInstanceModificationPriceRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeInstanceModificationPriceRequest) GetISP() *string {
	return s.ISP
}

func (s *DescribeInstanceModificationPriceRequest) GetImageId() *string {
	return s.ImageId
}

func (s *DescribeInstanceModificationPriceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeInstanceModificationPriceRequest) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DescribeInstanceModificationPriceRequest) GetInternetChargeType() *string {
	return s.InternetChargeType
}

func (s *DescribeInstanceModificationPriceRequest) GetInternetMaxBandwidthOut() *int32 {
	return s.InternetMaxBandwidthOut
}

func (s *DescribeInstanceModificationPriceRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeInstanceModificationPriceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeInstanceModificationPriceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeInstanceModificationPriceRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeInstanceModificationPriceRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeInstanceModificationPriceRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeInstanceModificationPriceRequest) SetSystemDisk(v *DescribeInstanceModificationPriceRequestSystemDisk) *DescribeInstanceModificationPriceRequest {
	s.SystemDisk = v
	return s
}

func (s *DescribeInstanceModificationPriceRequest) SetDataDisk(v []*DescribeInstanceModificationPriceRequestDataDisk) *DescribeInstanceModificationPriceRequest {
	s.DataDisk = v
	return s
}

func (s *DescribeInstanceModificationPriceRequest) SetEndTime(v string) *DescribeInstanceModificationPriceRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeInstanceModificationPriceRequest) SetISP(v string) *DescribeInstanceModificationPriceRequest {
	s.ISP = &v
	return s
}

func (s *DescribeInstanceModificationPriceRequest) SetImageId(v string) *DescribeInstanceModificationPriceRequest {
	s.ImageId = &v
	return s
}

func (s *DescribeInstanceModificationPriceRequest) SetInstanceId(v string) *DescribeInstanceModificationPriceRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstanceModificationPriceRequest) SetInstanceType(v string) *DescribeInstanceModificationPriceRequest {
	s.InstanceType = &v
	return s
}

func (s *DescribeInstanceModificationPriceRequest) SetInternetChargeType(v string) *DescribeInstanceModificationPriceRequest {
	s.InternetChargeType = &v
	return s
}

func (s *DescribeInstanceModificationPriceRequest) SetInternetMaxBandwidthOut(v int32) *DescribeInstanceModificationPriceRequest {
	s.InternetMaxBandwidthOut = &v
	return s
}

func (s *DescribeInstanceModificationPriceRequest) SetOwnerAccount(v string) *DescribeInstanceModificationPriceRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeInstanceModificationPriceRequest) SetOwnerId(v int64) *DescribeInstanceModificationPriceRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeInstanceModificationPriceRequest) SetRegionId(v string) *DescribeInstanceModificationPriceRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeInstanceModificationPriceRequest) SetResourceOwnerAccount(v string) *DescribeInstanceModificationPriceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeInstanceModificationPriceRequest) SetResourceOwnerId(v int64) *DescribeInstanceModificationPriceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeInstanceModificationPriceRequest) SetStartTime(v string) *DescribeInstanceModificationPriceRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeInstanceModificationPriceRequest) Validate() error {
	if s.SystemDisk != nil {
		if err := s.SystemDisk.Validate(); err != nil {
			return err
		}
	}
	if s.DataDisk != nil {
		for _, item := range s.DataDisk {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeInstanceModificationPriceRequestSystemDisk struct {
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// example:
	//
	// PL0
	PerformanceLevel *string `json:"PerformanceLevel,omitempty" xml:"PerformanceLevel,omitempty"`
	// example:
	//
	// 40
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s DescribeInstanceModificationPriceRequestSystemDisk) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceModificationPriceRequestSystemDisk) GoString() string {
	return s.String()
}

func (s *DescribeInstanceModificationPriceRequestSystemDisk) GetCategory() *string {
	return s.Category
}

func (s *DescribeInstanceModificationPriceRequestSystemDisk) GetPerformanceLevel() *string {
	return s.PerformanceLevel
}

func (s *DescribeInstanceModificationPriceRequestSystemDisk) GetSize() *int32 {
	return s.Size
}

func (s *DescribeInstanceModificationPriceRequestSystemDisk) SetCategory(v string) *DescribeInstanceModificationPriceRequestSystemDisk {
	s.Category = &v
	return s
}

func (s *DescribeInstanceModificationPriceRequestSystemDisk) SetPerformanceLevel(v string) *DescribeInstanceModificationPriceRequestSystemDisk {
	s.PerformanceLevel = &v
	return s
}

func (s *DescribeInstanceModificationPriceRequestSystemDisk) SetSize(v int32) *DescribeInstanceModificationPriceRequestSystemDisk {
	s.Size = &v
	return s
}

func (s *DescribeInstanceModificationPriceRequestSystemDisk) Validate() error {
	return dara.Validate(s)
}

type DescribeInstanceModificationPriceRequestDataDisk struct {
	// The category of the data disk. Specify this parameter when you want to query the price of new subscription data disks to be attached to the ECS instance. Valid values of N: 1 to 16. Valid values:
	//
	// - cloud_efficiency: ultra disk.
	//
	// - cloud_ssd: standard SSD.
	//
	// - cloud_essd: enterprise SSD.
	//
	// - cloud: basic disk.
	//
	// Default value: null.
	//
	// > The instance type parameter (`InstanceType`) and data disk parameters (`DataDisk.N.*`) cannot both be empty. You must specify at least one.
	//
	// example:
	//
	// cloud_essd
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// example:
	//
	// d-bf4rupt9****
	DiskId *string `json:"DiskId,omitempty" xml:"DiskId,omitempty"`
	// The performance level of the data disk when the data disk is an enterprise SSD. The value of N must be the same as that in `DataDisk.N.Category=cloud_essd`. Valid values:
	//
	// - PL0: a single disk can deliver up to 10,000 random read/write IOPS.
	//
	// - PL1: a single disk can deliver up to 50,000 random read/write IOPS.
	//
	// - PL2: a single disk can deliver up to 100,000 random read/write IOPS.
	//
	// - PL3: a single disk can deliver up to 1,000,000 random read/write IOPS.
	//
	// Default value: PL1.
	//
	// For information about how to select an ESSD performance level, see [ESSD](https://help.aliyun.com/document_detail/122389.html).
	//
	// example:
	//
	// PL1
	PerformanceLevel *string `json:"PerformanceLevel,omitempty" xml:"PerformanceLevel,omitempty"`
	// The capacity of the data disk. Valid values of N: 1 to 16. Unit: GiB. Valid values:
	//
	// - cloud_efficiency: 20 to 32768.
	//
	// - cloud_ssd: 20 to 32768.
	//
	// - cloud_essd: The valid values depend on the value of `DataDisk.N.PerformanceLevel`.
	//
	//     - PL0: 1 to 32768.
	//
	//     - PL1: 20 to 32768.
	//
	//     - PL2: 461 to 32768.
	//
	//     - PL3: 1261 to 32768.
	//
	// - cloud: 5 to 2000.
	//
	// Default value: the minimum capacity for the specified data disk category.
	//
	// example:
	//
	// 100
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s DescribeInstanceModificationPriceRequestDataDisk) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceModificationPriceRequestDataDisk) GoString() string {
	return s.String()
}

func (s *DescribeInstanceModificationPriceRequestDataDisk) GetCategory() *string {
	return s.Category
}

func (s *DescribeInstanceModificationPriceRequestDataDisk) GetDiskId() *string {
	return s.DiskId
}

func (s *DescribeInstanceModificationPriceRequestDataDisk) GetPerformanceLevel() *string {
	return s.PerformanceLevel
}

func (s *DescribeInstanceModificationPriceRequestDataDisk) GetSize() *int32 {
	return s.Size
}

func (s *DescribeInstanceModificationPriceRequestDataDisk) SetCategory(v string) *DescribeInstanceModificationPriceRequestDataDisk {
	s.Category = &v
	return s
}

func (s *DescribeInstanceModificationPriceRequestDataDisk) SetDiskId(v string) *DescribeInstanceModificationPriceRequestDataDisk {
	s.DiskId = &v
	return s
}

func (s *DescribeInstanceModificationPriceRequestDataDisk) SetPerformanceLevel(v string) *DescribeInstanceModificationPriceRequestDataDisk {
	s.PerformanceLevel = &v
	return s
}

func (s *DescribeInstanceModificationPriceRequestDataDisk) SetSize(v int32) *DescribeInstanceModificationPriceRequestDataDisk {
	s.Size = &v
	return s
}

func (s *DescribeInstanceModificationPriceRequestDataDisk) Validate() error {
	return dara.Validate(s)
}
