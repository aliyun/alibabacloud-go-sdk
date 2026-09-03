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
	// Default value: true. Valid values:
	//
	// - true: Automatic payment is enabled. Make sure that your Alibaba Cloud account balance is sufficient. Otherwise, abnormal orders may be generated.
	//
	// - false: Only an order is generated. Automatic payment is not enabled.
	//
	// example:
	//
	// false
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// The cloud computer ID.
	//
	// example:
	//
	// ecd-4543qyik164a4****
	DesktopId *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
	// The target instance type. You can call [DescribeDesktopTypes](https://help.aliyun.com/document_detail/188882.html) to query the instance types supported by cloud computers.
	//
	// This parameter is required.
	//
	// example:
	//
	// eds.general.2c8g
	DesktopType *string `json:"DesktopType,omitempty" xml:"DesktopType,omitempty"`
	// The promotion ID.
	//
	// example:
	//
	// 500033080110596
	PromotionId *string `json:"PromotionId,omitempty" xml:"PromotionId,omitempty"`
	// The region ID. You can call [DescribeRegions](https://help.aliyun.com/document_detail/196646.html) to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The user ID of the resource ownership in the reseller pattern. This parameter is not required in the non-reseller pattern.
	//
	// example:
	//
	// 1422724566551XXX
	ResellerOwnerUid *int64 `json:"ResellerOwnerUid,omitempty" xml:"ResellerOwnerUid,omitempty"`
	// The resource specification templates.
	ResourceSpecs []*ModifyDesktopSpecRequestResourceSpecs `json:"ResourceSpecs,omitempty" xml:"ResourceSpecs,omitempty" type:"Repeated"`
	// The resource type.
	//
	// > This parameter is not required for non-subscription cloud computers.
	//
	// example:
	//
	// DesktopMonthPackage
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The system cloud disk size after the change. Unit: GiB. Valid values: 80 to 500. The value must be a multiple of 10.
	//
	// example:
	//
	// 80
	RootDiskSizeGib *int32 `json:"RootDiskSizeGib,omitempty" xml:"RootDiskSizeGib,omitempty"`
	// The performance level (PL) of the data cloud disk. Default value: PL0.
	//
	// Valid values:
	//
	// - PL0
	//
	// - PL1
	//
	// - PL2
	//
	// - PL3
	//
	// example:
	//
	// PL0
	UserDiskPerformanceLevel *string `json:"UserDiskPerformanceLevel,omitempty" xml:"UserDiskPerformanceLevel,omitempty"`
	// The data cloud disk size after the change. Unit: GiB.
	//
	// - For non-graphics cloud computers, valid values: 20 to 1020. The value must be a multiple of 10.
	//
	// - For graphics cloud computers, valid values: 40 to 1020. The value must be a multiple of 10.
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
	// The cloud computer ID.
	//
	// example:
	//
	// ecd-4543qyik164a4****
	DesktopId *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
	// The target system cloud disk size. Valid values: 80 to 500 GiB. The value must be a multiple of 10.
	//
	// example:
	//
	// 80
	RootDiskSizeGib *int32 `json:"RootDiskSizeGib,omitempty" xml:"RootDiskSizeGib,omitempty"`
	// The target data cloud disk size. Valid values: 80 to 500 GiB. The value must be a multiple of 10.
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
