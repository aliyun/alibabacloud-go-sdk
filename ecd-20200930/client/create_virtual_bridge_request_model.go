// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVirtualBridgeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoPay(v bool) *CreateVirtualBridgeRequest
	GetAutoPay() *bool
	SetAutoRenew(v bool) *CreateVirtualBridgeRequest
	GetAutoRenew() *bool
	SetBridgeLevel(v string) *CreateVirtualBridgeRequest
	GetBridgeLevel() *string
	SetOfficeSiteId(v string) *CreateVirtualBridgeRequest
	GetOfficeSiteId() *string
	SetPaidCallBackUrl(v string) *CreateVirtualBridgeRequest
	GetPaidCallBackUrl() *string
	SetPeriod(v int32) *CreateVirtualBridgeRequest
	GetPeriod() *int32
	SetPeriodUnit(v string) *CreateVirtualBridgeRequest
	GetPeriodUnit() *string
	SetPromotionId(v string) *CreateVirtualBridgeRequest
	GetPromotionId() *string
	SetRegionId(v string) *CreateVirtualBridgeRequest
	GetRegionId() *string
}

type CreateVirtualBridgeRequest struct {
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
	// The bridge specifications.
	//
	// example:
	//
	// vb.ultra
	BridgeLevel *string `json:"BridgeLevel,omitempty" xml:"BridgeLevel,omitempty"`
	// The office network ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou+dir-467671****
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	// The payment callback URL.
	//
	// example:
	//
	// https://wya.wuying.aliyun.com/mobileClaw
	PaidCallBackUrl *string `json:"PaidCallBackUrl,omitempty" xml:"PaidCallBackUrl,omitempty"`
	// The duration for which you want to purchase the resource. The unit is specified by `PeriodUnit`. This parameter takes effect and is required only when `ChargeType` is set to `PrePaid`.
	//
	// - If `PeriodUnit` is set to `Month`, valid values:
	//
	//      - 1
	//
	//     -  2
	//
	//     - 3
	//
	//     - 6
	//
	// - If `PeriodUnit` is set to `Year`, valid values:
	//
	//     - 1
	//
	//     - 2
	//
	//     - 3
	//
	//     - 4
	//
	//     - 5
	//
	// example:
	//
	// 1
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The unit of duration for the subscription billable methods.
	//
	// example:
	//
	// Week
	PeriodUnit *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	// The promotion ID.
	//
	// example:
	//
	// youhuiquan_promotion_option_id_for_blank
	PromotionId *string `json:"PromotionId,omitempty" xml:"PromotionId,omitempty"`
	// The region ID. This feature is not region-specific. Set this parameter to cn-shanghai.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateVirtualBridgeRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateVirtualBridgeRequest) GoString() string {
	return s.String()
}

func (s *CreateVirtualBridgeRequest) GetAutoPay() *bool {
	return s.AutoPay
}

func (s *CreateVirtualBridgeRequest) GetAutoRenew() *bool {
	return s.AutoRenew
}

func (s *CreateVirtualBridgeRequest) GetBridgeLevel() *string {
	return s.BridgeLevel
}

func (s *CreateVirtualBridgeRequest) GetOfficeSiteId() *string {
	return s.OfficeSiteId
}

func (s *CreateVirtualBridgeRequest) GetPaidCallBackUrl() *string {
	return s.PaidCallBackUrl
}

func (s *CreateVirtualBridgeRequest) GetPeriod() *int32 {
	return s.Period
}

func (s *CreateVirtualBridgeRequest) GetPeriodUnit() *string {
	return s.PeriodUnit
}

func (s *CreateVirtualBridgeRequest) GetPromotionId() *string {
	return s.PromotionId
}

func (s *CreateVirtualBridgeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateVirtualBridgeRequest) SetAutoPay(v bool) *CreateVirtualBridgeRequest {
	s.AutoPay = &v
	return s
}

func (s *CreateVirtualBridgeRequest) SetAutoRenew(v bool) *CreateVirtualBridgeRequest {
	s.AutoRenew = &v
	return s
}

func (s *CreateVirtualBridgeRequest) SetBridgeLevel(v string) *CreateVirtualBridgeRequest {
	s.BridgeLevel = &v
	return s
}

func (s *CreateVirtualBridgeRequest) SetOfficeSiteId(v string) *CreateVirtualBridgeRequest {
	s.OfficeSiteId = &v
	return s
}

func (s *CreateVirtualBridgeRequest) SetPaidCallBackUrl(v string) *CreateVirtualBridgeRequest {
	s.PaidCallBackUrl = &v
	return s
}

func (s *CreateVirtualBridgeRequest) SetPeriod(v int32) *CreateVirtualBridgeRequest {
	s.Period = &v
	return s
}

func (s *CreateVirtualBridgeRequest) SetPeriodUnit(v string) *CreateVirtualBridgeRequest {
	s.PeriodUnit = &v
	return s
}

func (s *CreateVirtualBridgeRequest) SetPromotionId(v string) *CreateVirtualBridgeRequest {
	s.PromotionId = &v
	return s
}

func (s *CreateVirtualBridgeRequest) SetRegionId(v string) *CreateVirtualBridgeRequest {
	s.RegionId = &v
	return s
}

func (s *CreateVirtualBridgeRequest) Validate() error {
	return dara.Validate(s)
}
