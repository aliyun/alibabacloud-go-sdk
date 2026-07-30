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
	// Specifies whether to enable automatic payment.
	//
	// - If you set this parameter to `true`, ensure that your account balance is sufficient. Otherwise, abnormal orders are generated.
	//
	// - If you set this parameter to `false`, log on to the console and go to the **Expenses and Costs*	- page to complete the payment based on the returned order ID.
	//
	// example:
	//
	// false
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// The cloud computer ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ecd-2yjhqxo1monbf****
	DesktopId *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
	// The promotion ID. You can call the pricing API to obtain the list of matched promotion IDs.
	//
	// example:
	//
	// youhuiquan_promotion_option_id_for_blank
	PromotionId *string `json:"PromotionId,omitempty" xml:"PromotionId,omitempty"`
	// The region ID. Call [DescribeRegions](~~DescribeRegions~~) to query the regions supported by Elastic Desktop Service.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId         *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResellerOwnerUid *int64  `json:"ResellerOwnerUid,omitempty" xml:"ResellerOwnerUid,omitempty"`
	// The performance level of the system cloud disk. You can set the disk performance level when the cloud computer specification is Enterprise Graphics or High Frequency.
	//
	// example:
	//
	// PL1
	RootDiskPerformanceLevel *string `json:"RootDiskPerformanceLevel,omitempty" xml:"RootDiskPerformanceLevel,omitempty"`
	// The performance level of the data cloud disk. You can set the disk performance level when the cloud computer specification is Enterprise Graphics or High Frequency.
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
