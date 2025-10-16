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
	// The maximum public bandwidth. Unit: Mbit/s.
	//
	// >  Valid values when PayByTraffic is set to PayByBandwidth: 10 to 1000.
	//
	// example:
	//
	// 20
	Bandwidth *int32 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The ID of either the monthly subscription cloud computer with unlimited hours or the premium bandwidth plan.
	//
	// example:
	//
	// ecd-0gfv2z3sf95zvt****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The specifications.
	//
	// 	- Valid values when you set `ResourceType` to `Desktop`:
	//
	//     	- ecd.basic.small
	//
	//     	- ecd.basic.large
	//
	//     	- ecd.advanced.large
	//
	//     	- ecd.advanced.xlarge
	//
	//     	- ecd.performance.2xlarge
	//
	//     	- ecd.graphics.xlarge
	//
	//     	- ecd.graphics.2xlarge
	//
	//     	- ecd.advanced.xlarge_s8d2
	//
	//     	- ecd.advanced.xlarge_s8d7
	//
	//     	- ecd.graphics.1g72c
	//
	//     	- eds.general.2c2g
	//
	//     	- eds.general.2c4g
	//
	//     	- eds.general.2c8g
	//
	//     	- eds.general.4c8g
	//
	//     	- eds.general.4c16g
	//
	//     	- eds.general.8c16g
	//
	//     	- eds.general.8c32g
	//
	//     	- eds.general.16c32g
	//
	// 	- You can skip this parameter if `ResourceType` is set to `NetworkPackage`.
	//
	// example:
	//
	// eds.enterprise_office.8c16g
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	PromotionId  *string `json:"PromotionId,omitempty" xml:"PromotionId,omitempty"`
	// The region ID. You can call the [DescribeRegions](~~DescribeRegions~~) operation to query the list of regions where Elastic Desktop Service (EDS) Enterprise is available.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId         *string                                          `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResellerOwnerUid *int64                                           `json:"ResellerOwnerUid,omitempty" xml:"ResellerOwnerUid,omitempty"`
	ResourceSpecs    []*DescribeModificationPriceRequestResourceSpecs `json:"ResourceSpecs,omitempty" xml:"ResourceSpecs,omitempty" type:"Repeated"`
	// The resource type. The required parameters depend on the resource type.
	//
	// 	- When `ResourceType` is set to `Desktop`, the required parameters are `InstanceType`, `RootDiskSizeGib`, and `UserDiskSizeGib`.
	//
	// 	- When `ResourceType` is set to `NetworkPackage`, the required parameter is `Bandwidth`.
	//
	// Valid values:
	//
	// 	- Desktop (default): cloud computers.
	//
	// 	- NetworkPackage: premium bandwidth plans.
	//
	// example:
	//
	// Desktop
	ResourceType             *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	RootDiskPerformanceLevel *string `json:"RootDiskPerformanceLevel,omitempty" xml:"RootDiskPerformanceLevel,omitempty"`
	// The size of the system disk. Unit: GiB.
	//
	// example:
	//
	// 80
	RootDiskSizeGib          *int32  `json:"RootDiskSizeGib,omitempty" xml:"RootDiskSizeGib,omitempty"`
	UserDiskPerformanceLevel *string `json:"UserDiskPerformanceLevel,omitempty" xml:"UserDiskPerformanceLevel,omitempty"`
	// The size of the data disk. Unit: GiB.
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
	DesktopId       *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
	RootDiskSizeGib *int32  `json:"RootDiskSizeGib,omitempty" xml:"RootDiskSizeGib,omitempty"`
	UserDiskSizeGib *int32  `json:"UserDiskSizeGib,omitempty" xml:"UserDiskSizeGib,omitempty"`
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
