// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRenewDesktopsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoPay(v bool) *RenewDesktopsRequest
	GetAutoPay() *bool
	SetAutoRenew(v bool) *RenewDesktopsRequest
	GetAutoRenew() *bool
	SetDesktopId(v []*string) *RenewDesktopsRequest
	GetDesktopId() []*string
	SetPeriod(v int32) *RenewDesktopsRequest
	GetPeriod() *int32
	SetPeriodUnit(v string) *RenewDesktopsRequest
	GetPeriodUnit() *string
	SetPromotionId(v string) *RenewDesktopsRequest
	GetPromotionId() *string
	SetRegionId(v string) *RenewDesktopsRequest
	GetRegionId() *string
	SetResellerOwnerUid(v int64) *RenewDesktopsRequest
	GetResellerOwnerUid() *int64
	SetResourceType(v string) *RenewDesktopsRequest
	GetResourceType() *string
}

type RenewDesktopsRequest struct {
	// Specifies whether to enable automatic payment.
	//
	// example:
	//
	// true
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// Specifies whether to enable auto-renewal.
	//
	// example:
	//
	// false
	AutoRenew *bool `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// A list of WUYING Workspace instance IDs. You can renew only instances purchased on a monthly basis.
	//
	// This parameter is required.
	//
	// example:
	//
	// ecd-ia2zw38bi6cm7****
	DesktopId []*string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty" type:"Repeated"`
	// The renewal duration. Valid values depend on the value of `PeriodUnit`.
	//
	// - If `PeriodUnit` is `Month`, valid values are 1, 2, 3, and 6.
	//
	// - If `PeriodUnit` is `Year`, valid values are 1 to 5.
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The unit for the renewal duration, which applies to the `Period` parameter.
	//
	// example:
	//
	// Month
	PeriodUnit *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	// The promotion ID.
	//
	// example:
	//
	// 50003098015****
	PromotionId *string `json:"PromotionId,omitempty" xml:"PromotionId,omitempty"`
	// The region ID. Call [DescribeRegions](~~DescribeRegions~~) to list the regions where WUYING Workspace is available.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId         *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResellerOwnerUid *int64  `json:"ResellerOwnerUid,omitempty" xml:"ResellerOwnerUid,omitempty"`
	// > This field is not available for public use.
	//
	// example:
	//
	// null
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s RenewDesktopsRequest) String() string {
	return dara.Prettify(s)
}

func (s RenewDesktopsRequest) GoString() string {
	return s.String()
}

func (s *RenewDesktopsRequest) GetAutoPay() *bool {
	return s.AutoPay
}

func (s *RenewDesktopsRequest) GetAutoRenew() *bool {
	return s.AutoRenew
}

func (s *RenewDesktopsRequest) GetDesktopId() []*string {
	return s.DesktopId
}

func (s *RenewDesktopsRequest) GetPeriod() *int32 {
	return s.Period
}

func (s *RenewDesktopsRequest) GetPeriodUnit() *string {
	return s.PeriodUnit
}

func (s *RenewDesktopsRequest) GetPromotionId() *string {
	return s.PromotionId
}

func (s *RenewDesktopsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *RenewDesktopsRequest) GetResellerOwnerUid() *int64 {
	return s.ResellerOwnerUid
}

func (s *RenewDesktopsRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *RenewDesktopsRequest) SetAutoPay(v bool) *RenewDesktopsRequest {
	s.AutoPay = &v
	return s
}

func (s *RenewDesktopsRequest) SetAutoRenew(v bool) *RenewDesktopsRequest {
	s.AutoRenew = &v
	return s
}

func (s *RenewDesktopsRequest) SetDesktopId(v []*string) *RenewDesktopsRequest {
	s.DesktopId = v
	return s
}

func (s *RenewDesktopsRequest) SetPeriod(v int32) *RenewDesktopsRequest {
	s.Period = &v
	return s
}

func (s *RenewDesktopsRequest) SetPeriodUnit(v string) *RenewDesktopsRequest {
	s.PeriodUnit = &v
	return s
}

func (s *RenewDesktopsRequest) SetPromotionId(v string) *RenewDesktopsRequest {
	s.PromotionId = &v
	return s
}

func (s *RenewDesktopsRequest) SetRegionId(v string) *RenewDesktopsRequest {
	s.RegionId = &v
	return s
}

func (s *RenewDesktopsRequest) SetResellerOwnerUid(v int64) *RenewDesktopsRequest {
	s.ResellerOwnerUid = &v
	return s
}

func (s *RenewDesktopsRequest) SetResourceType(v string) *RenewDesktopsRequest {
	s.ResourceType = &v
	return s
}

func (s *RenewDesktopsRequest) Validate() error {
	return dara.Validate(s)
}
