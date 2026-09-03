// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeModificationPriceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBandwidth(v int32) *DescribeModificationPriceRequest
	GetBandwidth() *int32
	SetInstanceId(v string) *DescribeModificationPriceRequest
	GetInstanceId() *string
	SetInstanceType(v string) *DescribeModificationPriceRequest
	GetInstanceType() *string
	SetPromotionId(v string) *DescribeModificationPriceRequest
	GetPromotionId() *string
	SetRegionId(v string) *DescribeModificationPriceRequest
	GetRegionId() *string
	SetResellerOwnerUid(v int64) *DescribeModificationPriceRequest
	GetResellerOwnerUid() *int64
	SetResourceSpecs(v []*DescribeModificationPriceRequestResourceSpecs) *DescribeModificationPriceRequest
	GetResourceSpecs() []*DescribeModificationPriceRequestResourceSpecs
	SetResourceType(v string) *DescribeModificationPriceRequest
	GetResourceType() *string
	SetRootDiskPerformanceLevel(v string) *DescribeModificationPriceRequest
	GetRootDiskPerformanceLevel() *string
	SetRootDiskSizeGib(v int32) *DescribeModificationPriceRequest
	GetRootDiskSizeGib() *int32
	SetUserDiskPerformanceLevel(v string) *DescribeModificationPriceRequest
	GetUserDiskPerformanceLevel() *string
	SetUserDiskSizeGib(v int32) *DescribeModificationPriceRequest
	GetUserDiskSizeGib() *int32
}

type DescribeModificationPriceRequest struct {
	// The peak Internet bandwidth. Unit: Mbit/s.
	//
	// > If you use the pay-by-fixed-bandwidth billing method, valid values are 10 to 1000.
	//
	// example:
	//
	// 20
	Bandwidth *int32 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The instance ID. The value can be the ID of a monthly subscription (unlimited duration) cloud computer or the ID of a premium Internet bandwidth instance.
	//
	// example:
	//
	// ecd-0gfv2z3sf95zvt****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The resource specification.
	//
	// - If ResourceType is set to Desktop, valid values include:
	//
	//     - ecd.basic.small
	//
	//     - ecd.basic.large
	//
	//     - ecd.advanced.large
	//
	//     - ecd.advanced.xlarge
	//
	//     - ecd.performance.2xlarge
	//
	//     - ecd.graphics.xlarge
	//
	//     - ecd.graphics.2xlarge
	//
	//     - ecd.advanced.xlarge_s8d2
	//
	//     - ecd.advanced.xlarge_s8d7
	//
	//     - ecd.graphics.1g72c
	//
	//     - eds.general.2c2g
	//
	//     - eds.general.2c4g
	//
	//     - eds.general.2c8g
	//
	//     - eds.general.4c8g
	//
	//     - eds.general.4c16g
	//
	//     - eds.general.8c16g
	//
	//     - eds.general.8c32g
	//
	//     - eds.general.16c32g
	//
	// - If ResourceType is set to NetworkPackage, you do not need to specify this parameter.
	//
	// example:
	//
	// eds.enterprise_office.8c16g
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The promotion ID.
	//
	// example:
	//
	// youhuiquan_promotion_option_id_for_blank
	PromotionId *string `json:"PromotionId,omitempty" xml:"PromotionId,omitempty"`
	// The region ID. You can call [DescribeRegions](~~DescribeRegions~~) to query the list of regions supported by Elastic Desktop Service.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource ownership user ID in the reseller pattern. You do not need to specify this parameter in non-reseller pattern.
	//
	// example:
	//
	// 1422724566551XXX
	ResellerOwnerUid *int64 `json:"ResellerOwnerUid,omitempty" xml:"ResellerOwnerUid,omitempty"`
	// The list of resource specification templates.
	ResourceSpecs []*DescribeModificationPriceRequestResourceSpecs `json:"ResourceSpecs,omitempty" xml:"ResourceSpecs,omitempty" type:"Repeated"`
	// The resource type. The required parameters vary based on the resource type for which you want to query the upgrade/downgrade price:
	//
	// - If ResourceType is set to Desktop, you must specify the InstanceType, RootDiskSizeGib, and UserDiskSizeGib parameters.
	//
	// - If ResourceType is set to NetworkPackage, you must specify the Bandwidth parameter.
	//
	// example:
	//
	// Desktop
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The performance level (PL) of the system cloud disk. You can set the disk performance level when the cloud computer specification in Settings is Graphics or High Frequency. For more information about the differences between disk performance levels, see [ESSDs](https://help.aliyun.com/document_detail/122389.html). standard SSD and ESSD have different performance levels.
	//
	// example:
	//
	// PL0
	RootDiskPerformanceLevel *string `json:"RootDiskPerformanceLevel,omitempty" xml:"RootDiskPerformanceLevel,omitempty"`
	// The size of the system cloud disk. Unit: GiB.
	//
	// example:
	//
	// 80
	RootDiskSizeGib *int32 `json:"RootDiskSizeGib,omitempty" xml:"RootDiskSizeGib,omitempty"`
	// The performance level (PL) of the data cloud disk. You can set the disk performance level when the cloud computer specification in Settings is Graphics or High Frequency. For more information about the differences between disk performance levels, see [ESSDs](https://help.aliyun.com/document_detail/122389.html). standard SSD and ESSD have different performance levels.
	//
	// example:
	//
	// PL0
	UserDiskPerformanceLevel *string `json:"UserDiskPerformanceLevel,omitempty" xml:"UserDiskPerformanceLevel,omitempty"`
	// The size of the data cloud disk. Unit: GiB.
	//
	// example:
	//
	// 50
	UserDiskSizeGib *int32 `json:"UserDiskSizeGib,omitempty" xml:"UserDiskSizeGib,omitempty"`
}

func (s DescribeModificationPriceRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeModificationPriceRequest) GoString() string {
	return s.String()
}

func (s *DescribeModificationPriceRequest) GetBandwidth() *int32 {
	return s.Bandwidth
}

func (s *DescribeModificationPriceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeModificationPriceRequest) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DescribeModificationPriceRequest) GetPromotionId() *string {
	return s.PromotionId
}

func (s *DescribeModificationPriceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeModificationPriceRequest) GetResellerOwnerUid() *int64 {
	return s.ResellerOwnerUid
}

func (s *DescribeModificationPriceRequest) GetResourceSpecs() []*DescribeModificationPriceRequestResourceSpecs {
	return s.ResourceSpecs
}

func (s *DescribeModificationPriceRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *DescribeModificationPriceRequest) GetRootDiskPerformanceLevel() *string {
	return s.RootDiskPerformanceLevel
}

func (s *DescribeModificationPriceRequest) GetRootDiskSizeGib() *int32 {
	return s.RootDiskSizeGib
}

func (s *DescribeModificationPriceRequest) GetUserDiskPerformanceLevel() *string {
	return s.UserDiskPerformanceLevel
}

func (s *DescribeModificationPriceRequest) GetUserDiskSizeGib() *int32 {
	return s.UserDiskSizeGib
}

func (s *DescribeModificationPriceRequest) SetBandwidth(v int32) *DescribeModificationPriceRequest {
	s.Bandwidth = &v
	return s
}

func (s *DescribeModificationPriceRequest) SetInstanceId(v string) *DescribeModificationPriceRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeModificationPriceRequest) SetInstanceType(v string) *DescribeModificationPriceRequest {
	s.InstanceType = &v
	return s
}

func (s *DescribeModificationPriceRequest) SetPromotionId(v string) *DescribeModificationPriceRequest {
	s.PromotionId = &v
	return s
}

func (s *DescribeModificationPriceRequest) SetRegionId(v string) *DescribeModificationPriceRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeModificationPriceRequest) SetResellerOwnerUid(v int64) *DescribeModificationPriceRequest {
	s.ResellerOwnerUid = &v
	return s
}

func (s *DescribeModificationPriceRequest) SetResourceSpecs(v []*DescribeModificationPriceRequestResourceSpecs) *DescribeModificationPriceRequest {
	s.ResourceSpecs = v
	return s
}

func (s *DescribeModificationPriceRequest) SetResourceType(v string) *DescribeModificationPriceRequest {
	s.ResourceType = &v
	return s
}

func (s *DescribeModificationPriceRequest) SetRootDiskPerformanceLevel(v string) *DescribeModificationPriceRequest {
	s.RootDiskPerformanceLevel = &v
	return s
}

func (s *DescribeModificationPriceRequest) SetRootDiskSizeGib(v int32) *DescribeModificationPriceRequest {
	s.RootDiskSizeGib = &v
	return s
}

func (s *DescribeModificationPriceRequest) SetUserDiskPerformanceLevel(v string) *DescribeModificationPriceRequest {
	s.UserDiskPerformanceLevel = &v
	return s
}

func (s *DescribeModificationPriceRequest) SetUserDiskSizeGib(v int32) *DescribeModificationPriceRequest {
	s.UserDiskSizeGib = &v
	return s
}

func (s *DescribeModificationPriceRequest) Validate() error {
	if s.ResourceSpecs != nil {
		for _, item := range s.ResourceSpecs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeModificationPriceRequestResourceSpecs struct {
	// The cloud computer ID.
	//
	// example:
	//
	// ecd-6ghhzivgmnzgeyXXX
	DesktopId *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
	// The size of the system cloud disk. Unit: GiB.
	//
	// example:
	//
	// 80
	RootDiskSizeGib *int32 `json:"RootDiskSizeGib,omitempty" xml:"RootDiskSizeGib,omitempty"`
	// The size of the data cloud disk. Unit: GiB.
	//
	// example:
	//
	// 100
	UserDiskSizeGib *int32 `json:"UserDiskSizeGib,omitempty" xml:"UserDiskSizeGib,omitempty"`
}

func (s DescribeModificationPriceRequestResourceSpecs) String() string {
	return dara.Prettify(s)
}

func (s DescribeModificationPriceRequestResourceSpecs) GoString() string {
	return s.String()
}

func (s *DescribeModificationPriceRequestResourceSpecs) GetDesktopId() *string {
	return s.DesktopId
}

func (s *DescribeModificationPriceRequestResourceSpecs) GetRootDiskSizeGib() *int32 {
	return s.RootDiskSizeGib
}

func (s *DescribeModificationPriceRequestResourceSpecs) GetUserDiskSizeGib() *int32 {
	return s.UserDiskSizeGib
}

func (s *DescribeModificationPriceRequestResourceSpecs) SetDesktopId(v string) *DescribeModificationPriceRequestResourceSpecs {
	s.DesktopId = &v
	return s
}

func (s *DescribeModificationPriceRequestResourceSpecs) SetRootDiskSizeGib(v int32) *DescribeModificationPriceRequestResourceSpecs {
	s.RootDiskSizeGib = &v
	return s
}

func (s *DescribeModificationPriceRequestResourceSpecs) SetUserDiskSizeGib(v int32) *DescribeModificationPriceRequestResourceSpecs {
	s.UserDiskSizeGib = &v
	return s
}

func (s *DescribeModificationPriceRequestResourceSpecs) Validate() error {
	return dara.Validate(s)
}
