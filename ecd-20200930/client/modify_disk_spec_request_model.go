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
	// - If set to `true`, ensure your account has a sufficient balance. Otherwise, an abnormal order is generated.
	//
	// - If set to `false`, log on to the console. Then, go to the **Expenses and Costs*	- page to pay for the order using the returned order ID.
	//
	// example:
	//
	// false
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// The ID of the cloud desktop.
	//
	// This parameter is required.
	//
	// example:
	//
	// ecd-2yjhqxo1monbf****
	DesktopId *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
	// The promotion ID. Call a pricing inquiry API to get a list of applicable promotion IDs.
	//
	// example:
	//
	// youhuiquan_promotion_option_id_for_blank
	PromotionId *string `json:"PromotionId,omitempty" xml:"PromotionId,omitempty"`
	// The region ID. Call [DescribeRegions](~~DescribeRegions~~) to get a list of supported regions for WUYING Workspace.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId         *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResellerOwnerUid *int64  `json:"ResellerOwnerUid,omitempty" xml:"ResellerOwnerUid,omitempty"`
	// The performance level of the system disk. You can set the performance level for Graphics or High-frequency workspaces.
	//
	// example:
	//
	// PL1
	RootDiskPerformanceLevel *string `json:"RootDiskPerformanceLevel,omitempty" xml:"RootDiskPerformanceLevel,omitempty"`
	// The performance level of the data disk. You can set the performance level for Graphics or High-frequency workspaces.
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
