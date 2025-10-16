// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDesktopSpecRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoPay(v bool) *ModifyDesktopSpecRequest
	GetAutoPay() *bool
	SetDesktopId(v string) *ModifyDesktopSpecRequest
	GetDesktopId() *string
	SetDesktopType(v string) *ModifyDesktopSpecRequest
	GetDesktopType() *string
	SetPromotionId(v string) *ModifyDesktopSpecRequest
	GetPromotionId() *string
	SetRegionId(v string) *ModifyDesktopSpecRequest
	GetRegionId() *string
	SetResellerOwnerUid(v int64) *ModifyDesktopSpecRequest
	GetResellerOwnerUid() *int64
	SetResourceSpecs(v []*ModifyDesktopSpecRequestResourceSpecs) *ModifyDesktopSpecRequest
	GetResourceSpecs() []*ModifyDesktopSpecRequestResourceSpecs
	SetResourceType(v string) *ModifyDesktopSpecRequest
	GetResourceType() *string
	SetRootDiskSizeGib(v int32) *ModifyDesktopSpecRequest
	GetRootDiskSizeGib() *int32
	SetUserDiskPerformanceLevel(v string) *ModifyDesktopSpecRequest
	GetUserDiskPerformanceLevel() *string
	SetUserDiskSizeGib(v int32) *ModifyDesktopSpecRequest
	GetUserDiskSizeGib() *int32
}

type ModifyDesktopSpecRequest struct {
	// Specifies whether to enable the auto-payment feature.
	//
	// Default value: true. Valid values:
	//
	// 	- true: enables the auto-payment feature.
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     Make sure that you have sufficient balance in your Alibaba Cloud account. Otherwise, an exception occurs on your order.
	//
	//     <!-- -->
	//
	// 	- false: disables the auto-payment feature. In this case, an order is generated, and no payment is automatically made.
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     You can log on to the Elastic Desktop Service console and complete the payment based on the order ID on the Orders page.
	//
	//     <!-- -->
	//
	// example:
	//
	// false
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// The ID of a cloud computer.
	//
	// example:
	//
	// ecd-4543qyik164a4****
	DesktopId *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
	// The destination instance type. You can call the [DescribeDesktopTypes](https://help.aliyun.com/document_detail/188882.html) operation to query the instance types supported by cloud computers.
	//
	// This parameter is required.
	//
	// example:
	//
	// eds.general.2c8g
	DesktopType *string `json:"DesktopType,omitempty" xml:"DesktopType,omitempty"`
	// The ID of the promotional activity.
	//
	// example:
	//
	// 500033080110596
	PromotionId *string `json:"PromotionId,omitempty" xml:"PromotionId,omitempty"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/196646.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId         *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResellerOwnerUid *int64  `json:"ResellerOwnerUid,omitempty" xml:"ResellerOwnerUid,omitempty"`
	// The array of resource specification templates.
	ResourceSpecs []*ModifyDesktopSpecRequestResourceSpecs `json:"ResourceSpecs,omitempty" xml:"ResourceSpecs,omitempty" type:"Repeated"`
	// The resource type.
	//
	// > This parameter is optional for non-subscribed cloud computers.
	//
	// example:
	//
	// DesktopMonthPackage
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The size of the new system disk. Unit: GiB. Valid values: 80 to 500 GiB. The value must be a multiple of 10.
	//
	// example:
	//
	// 80
	RootDiskSizeGib *int32 `json:"RootDiskSizeGib,omitempty" xml:"RootDiskSizeGib,omitempty"`
	// The performance level (PL) of the data disk. Default value: PL0.
	//
	// Valid values:
	//
	// 	- PL1
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- PL0
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- PL3
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- PL2
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// PL0
	UserDiskPerformanceLevel *string `json:"UserDiskPerformanceLevel,omitempty" xml:"UserDiskPerformanceLevel,omitempty"`
	// The destination data disk size. Unit: GiB.
	//
	// 	- The data disk size of a non-graphical cloud computer ranges from 20 to 1020 GiB and must be a multiple of 10.
	//
	// 	- The data disk size of a graphical cloud computer ranges from 40 to 1020 GiB and must be a multiple of 10.
	//
	// example:
	//
	// 100
	UserDiskSizeGib *int32 `json:"UserDiskSizeGib,omitempty" xml:"UserDiskSizeGib,omitempty"`
}

func (s ModifyDesktopSpecRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDesktopSpecRequest) GoString() string {
	return s.String()
}

func (s *ModifyDesktopSpecRequest) GetAutoPay() *bool {
	return s.AutoPay
}

func (s *ModifyDesktopSpecRequest) GetDesktopId() *string {
	return s.DesktopId
}

func (s *ModifyDesktopSpecRequest) GetDesktopType() *string {
	return s.DesktopType
}

func (s *ModifyDesktopSpecRequest) GetPromotionId() *string {
	return s.PromotionId
}

func (s *ModifyDesktopSpecRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyDesktopSpecRequest) GetResellerOwnerUid() *int64 {
	return s.ResellerOwnerUid
}

func (s *ModifyDesktopSpecRequest) GetResourceSpecs() []*ModifyDesktopSpecRequestResourceSpecs {
	return s.ResourceSpecs
}

func (s *ModifyDesktopSpecRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *ModifyDesktopSpecRequest) GetRootDiskSizeGib() *int32 {
	return s.RootDiskSizeGib
}

func (s *ModifyDesktopSpecRequest) GetUserDiskPerformanceLevel() *string {
	return s.UserDiskPerformanceLevel
}

func (s *ModifyDesktopSpecRequest) GetUserDiskSizeGib() *int32 {
	return s.UserDiskSizeGib
}

func (s *ModifyDesktopSpecRequest) SetAutoPay(v bool) *ModifyDesktopSpecRequest {
	s.AutoPay = &v
	return s
}

func (s *ModifyDesktopSpecRequest) SetDesktopId(v string) *ModifyDesktopSpecRequest {
	s.DesktopId = &v
	return s
}

func (s *ModifyDesktopSpecRequest) SetDesktopType(v string) *ModifyDesktopSpecRequest {
	s.DesktopType = &v
	return s
}

func (s *ModifyDesktopSpecRequest) SetPromotionId(v string) *ModifyDesktopSpecRequest {
	s.PromotionId = &v
	return s
}

func (s *ModifyDesktopSpecRequest) SetRegionId(v string) *ModifyDesktopSpecRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyDesktopSpecRequest) SetResellerOwnerUid(v int64) *ModifyDesktopSpecRequest {
	s.ResellerOwnerUid = &v
	return s
}

func (s *ModifyDesktopSpecRequest) SetResourceSpecs(v []*ModifyDesktopSpecRequestResourceSpecs) *ModifyDesktopSpecRequest {
	s.ResourceSpecs = v
	return s
}

func (s *ModifyDesktopSpecRequest) SetResourceType(v string) *ModifyDesktopSpecRequest {
	s.ResourceType = &v
	return s
}

func (s *ModifyDesktopSpecRequest) SetRootDiskSizeGib(v int32) *ModifyDesktopSpecRequest {
	s.RootDiskSizeGib = &v
	return s
}

func (s *ModifyDesktopSpecRequest) SetUserDiskPerformanceLevel(v string) *ModifyDesktopSpecRequest {
	s.UserDiskPerformanceLevel = &v
	return s
}

func (s *ModifyDesktopSpecRequest) SetUserDiskSizeGib(v int32) *ModifyDesktopSpecRequest {
	s.UserDiskSizeGib = &v
	return s
}

func (s *ModifyDesktopSpecRequest) Validate() error {
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

type ModifyDesktopSpecRequestResourceSpecs struct {
	// The ID of the cloud computer.
	//
	// example:
	//
	// ecd-4543qyik164a4****
	DesktopId *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
	// The target size of the system disk. Valid values: 80-500 GiB. The value must be a multiple of 10.
	//
	// example:
	//
	// 80
	RootDiskSizeGib *int32 `json:"RootDiskSizeGib,omitempty" xml:"RootDiskSizeGib,omitempty"`
	// The target size of the data disk. Valid values: 80-500 GiB. The value must be a multiple of 10.
	//
	// example:
	//
	// 20
	UserDiskSizeGib *int32 `json:"UserDiskSizeGib,omitempty" xml:"UserDiskSizeGib,omitempty"`
}

func (s ModifyDesktopSpecRequestResourceSpecs) String() string {
	return dara.Prettify(s)
}

func (s ModifyDesktopSpecRequestResourceSpecs) GoString() string {
	return s.String()
}

func (s *ModifyDesktopSpecRequestResourceSpecs) GetDesktopId() *string {
	return s.DesktopId
}

func (s *ModifyDesktopSpecRequestResourceSpecs) GetRootDiskSizeGib() *int32 {
	return s.RootDiskSizeGib
}

func (s *ModifyDesktopSpecRequestResourceSpecs) GetUserDiskSizeGib() *int32 {
	return s.UserDiskSizeGib
}

func (s *ModifyDesktopSpecRequestResourceSpecs) SetDesktopId(v string) *ModifyDesktopSpecRequestResourceSpecs {
	s.DesktopId = &v
	return s
}

func (s *ModifyDesktopSpecRequestResourceSpecs) SetRootDiskSizeGib(v int32) *ModifyDesktopSpecRequestResourceSpecs {
	s.RootDiskSizeGib = &v
	return s
}

func (s *ModifyDesktopSpecRequestResourceSpecs) SetUserDiskSizeGib(v int32) *ModifyDesktopSpecRequestResourceSpecs {
	s.UserDiskSizeGib = &v
	return s
}

func (s *ModifyDesktopSpecRequestResourceSpecs) Validate() error {
	return dara.Validate(s)
}
