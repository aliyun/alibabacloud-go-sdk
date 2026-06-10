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
	// Specifies whether to enable automatic payment.
	//
	// example:
	//
	// false
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// The ID of the cloud desktop.
	//
	// example:
	//
	// ecd-4543qyik164a4****
	DesktopId *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
	// The new desktop type. You can call the [DescribeDesktopTypes](~~DescribeDesktopTypes~~) operation to query the supported desktop types.
	//
	// This parameter is required.
	//
	// example:
	//
	// eds.general.2c4g
	DesktopType *string `json:"DesktopType,omitempty" xml:"DesktopType,omitempty"`
	// The promotion ID.
	//
	// example:
	//
	// 50003308011****
	PromotionId *string `json:"PromotionId,omitempty" xml:"PromotionId,omitempty"`
	// The ID of the region. You can call the [DescribeRegions](~~DescribeRegions~~) operation to obtain a list of regions that Elastic Desktop Service supports.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId         *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResellerOwnerUid *int64  `json:"ResellerOwnerUid,omitempty" xml:"ResellerOwnerUid,omitempty"`
	// A list of resource specification templates.
	ResourceSpecs []*ModifyDesktopSpecRequestResourceSpecs `json:"ResourceSpecs,omitempty" xml:"ResourceSpecs,omitempty" type:"Repeated"`
	// The resource type.
	//
	// > This parameter is required only for cloud desktops that use the subscription billing method.
	//
	// example:
	//
	// DesktopMonthPackage
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The new size of the system disk, in GiB. The value must be a multiple of 10 in the range of 80 to 500.
	//
	// example:
	//
	// 80
	RootDiskSizeGib *int32 `json:"RootDiskSizeGib,omitempty" xml:"RootDiskSizeGib,omitempty"`
	// The performance level of the data disk.
	//
	// example:
	//
	// PL0
	UserDiskPerformanceLevel *string `json:"UserDiskPerformanceLevel,omitempty" xml:"UserDiskPerformanceLevel,omitempty"`
	// The new size of the data disk, in GiB.
	//
	// - For non-graphics-accelerated desktop types, the value must be a multiple of 10 in the range of 20 to 1,020.
	//
	// - For graphics-accelerated desktop types, the value must be a multiple of 10 in the range of 40 to 1,020.
	//
	// example:
	//
	// 40
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
	// The ID of the cloud desktop.
	//
	// example:
	//
	// ecd-4543qyik164a4****
	DesktopId *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
	// The new size of the system disk, in GiB. The value must be a multiple of 10 in the range of 80 to 500.
	//
	// example:
	//
	// 80
	RootDiskSizeGib *int32 `json:"RootDiskSizeGib,omitempty" xml:"RootDiskSizeGib,omitempty"`
	// The new size of the data disk, in GiB. The value must be a multiple of 10 in the range of 20 to 2,040.
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
