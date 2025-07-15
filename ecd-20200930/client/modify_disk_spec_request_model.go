// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDiskSpecRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoPay(v bool) *ModifyDiskSpecRequest
	GetAutoPay() *bool
	SetDesktopId(v string) *ModifyDiskSpecRequest
	GetDesktopId() *string
	SetPromotionId(v string) *ModifyDiskSpecRequest
	GetPromotionId() *string
	SetRegionId(v string) *ModifyDiskSpecRequest
	GetRegionId() *string
	SetResellerOwnerUid(v int64) *ModifyDiskSpecRequest
	GetResellerOwnerUid() *int64
	SetRootDiskPerformanceLevel(v string) *ModifyDiskSpecRequest
	GetRootDiskPerformanceLevel() *string
	SetUserDiskPerformanceLevel(v string) *ModifyDiskSpecRequest
	GetUserDiskPerformanceLevel() *string
}

type ModifyDiskSpecRequest struct {
	// Specifies whether to enable the automatic payment feature.
	//
	// 	- If you set the value to `true`, ensure your account has sufficient balance to avoid generating abnormal orders.
	//
	// 	- If you set the value to `false`, go to the **Expenses and Costs*	- page to complete the payment based on the order number.
	//
	// Valid values:
	//
	// 	- true (default): enables the automatic payment feature.
	//
	// 	- false: generates the order and manually complete the payment.
	//
	// example:
	//
	// false
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// The ID of the cloud computer.
	//
	// This parameter is required.
	//
	// example:
	//
	// ecd-2yjhqxo1monxxxxxx
	DesktopId *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
	// The ID of the sales promotion activity. You can call the DescribePrice operation to obtain the IDs of matching sales promotion activities.
	//
	// example:
	//
	// youhuiquan_promotion_option_id_for_blank
	PromotionId *string `json:"PromotionId,omitempty" xml:"PromotionId,omitempty"`
	// The ID of the region. You can call the [DescribeRegions](~~DescribeRegions~~) operation to query the list of regions where Elastic Desktop Service (EDS) Enterprise is available.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId         *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResellerOwnerUid *int64  `json:"ResellerOwnerUid,omitempty" xml:"ResellerOwnerUid,omitempty"`
	// The PL of the system disk. Only Enterprise Graphics or High Frequency cloud computers support disk PL adjustments.
	//
	// Valid values:
	//
	// 	- PL1
	//
	// 	- PL0
	//
	// 	- PL3
	//
	// 	- PL2
	//
	// example:
	//
	// PL1
	RootDiskPerformanceLevel *string `json:"RootDiskPerformanceLevel,omitempty" xml:"RootDiskPerformanceLevel,omitempty"`
	// The PL of the data disk. Only Enterprise Graphics or High Frequency cloud computers support disk PL adjustments.
	//
	// Valid values:
	//
	// 	- PL1
	//
	// 	- PL0
	//
	// 	- PL3
	//
	// 	- PL2
	//
	// example:
	//
	// PL1
	UserDiskPerformanceLevel *string `json:"UserDiskPerformanceLevel,omitempty" xml:"UserDiskPerformanceLevel,omitempty"`
}

func (s ModifyDiskSpecRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDiskSpecRequest) GoString() string {
	return s.String()
}

func (s *ModifyDiskSpecRequest) GetAutoPay() *bool {
	return s.AutoPay
}

func (s *ModifyDiskSpecRequest) GetDesktopId() *string {
	return s.DesktopId
}

func (s *ModifyDiskSpecRequest) GetPromotionId() *string {
	return s.PromotionId
}

func (s *ModifyDiskSpecRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyDiskSpecRequest) GetResellerOwnerUid() *int64 {
	return s.ResellerOwnerUid
}

func (s *ModifyDiskSpecRequest) GetRootDiskPerformanceLevel() *string {
	return s.RootDiskPerformanceLevel
}

func (s *ModifyDiskSpecRequest) GetUserDiskPerformanceLevel() *string {
	return s.UserDiskPerformanceLevel
}

func (s *ModifyDiskSpecRequest) SetAutoPay(v bool) *ModifyDiskSpecRequest {
	s.AutoPay = &v
	return s
}

func (s *ModifyDiskSpecRequest) SetDesktopId(v string) *ModifyDiskSpecRequest {
	s.DesktopId = &v
	return s
}

func (s *ModifyDiskSpecRequest) SetPromotionId(v string) *ModifyDiskSpecRequest {
	s.PromotionId = &v
	return s
}

func (s *ModifyDiskSpecRequest) SetRegionId(v string) *ModifyDiskSpecRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyDiskSpecRequest) SetResellerOwnerUid(v int64) *ModifyDiskSpecRequest {
	s.ResellerOwnerUid = &v
	return s
}

func (s *ModifyDiskSpecRequest) SetRootDiskPerformanceLevel(v string) *ModifyDiskSpecRequest {
	s.RootDiskPerformanceLevel = &v
	return s
}

func (s *ModifyDiskSpecRequest) SetUserDiskPerformanceLevel(v string) *ModifyDiskSpecRequest {
	s.UserDiskPerformanceLevel = &v
	return s
}

func (s *ModifyDiskSpecRequest) Validate() error {
	return dara.Validate(s)
}
