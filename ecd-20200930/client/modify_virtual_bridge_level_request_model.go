// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyVirtualBridgeLevelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoPay(v bool) *ModifyVirtualBridgeLevelRequest
	GetAutoPay() *bool
	SetAutoRenew(v bool) *ModifyVirtualBridgeLevelRequest
	GetAutoRenew() *bool
	SetBridgeId(v string) *ModifyVirtualBridgeLevelRequest
	GetBridgeId() *string
	SetBridgeLevel(v string) *ModifyVirtualBridgeLevelRequest
	GetBridgeLevel() *string
	SetPaidCallBackUrl(v string) *ModifyVirtualBridgeLevelRequest
	GetPaidCallBackUrl() *string
	SetPeriod(v int32) *ModifyVirtualBridgeLevelRequest
	GetPeriod() *int32
	SetPeriodUnit(v string) *ModifyVirtualBridgeLevelRequest
	GetPeriodUnit() *string
	SetPromotionId(v string) *ModifyVirtualBridgeLevelRequest
	GetPromotionId() *string
	SetRegionId(v string) *ModifyVirtualBridgeLevelRequest
	GetRegionId() *string
}

type ModifyVirtualBridgeLevelRequest struct {
	// Specifies whether to enable automatic payment.
	//
	// example:
	//
	// true
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// Specifies whether to enable auto-renewal. This parameter takes effect and is optional only when the billing method is `PrePaid`.
	//
	// example:
	//
	// false
	AutoRenew *bool `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// The virtual bridge ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// vb-sfjoasjfosdfj**
	BridgeId *string `json:"BridgeId,omitempty" xml:"BridgeId,omitempty"`
	// The virtual bridge specifications.
	//
	// This parameter is required.
	//
	// example:
	//
	// vb.pro
	BridgeLevel *string `json:"BridgeLevel,omitempty" xml:"BridgeLevel,omitempty"`
	// The payment callback URL.
	//
	// example:
	//
	// https://wya.wuying.aliyun.com/mobileClaw
	PaidCallBackUrl *string `json:"PaidCallBackUrl,omitempty" xml:"PaidCallBackUrl,omitempty"`
	// The renewal duration. Valid values of this parameter are determined by the value of the `PeriodUnit` parameter.
	//
	// - If `PeriodUnit` is set to `Month`, valid values are 1, 2, 3, and 6.
	//
	// - If `PeriodUnit` is set to `Year`, valid values are 1, 2, and 3.
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The unit of the subscription duration for the prepaid cloud disk. This parameter takes effect and is required only when the `CdsChargeType` parameter is set to `PrePaid`.
	//
	// example:
	//
	// Month
	PeriodUnit *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	// The promotion ID. You can call the pricing query operation to obtain the list of matched promotion IDs.
	//
	// example:
	//
	// 23141
	PromotionId *string `json:"PromotionId,omitempty" xml:"PromotionId,omitempty"`
	// The region ID. You can call [DescribeRegions](~~DescribeRegions~~) to query the list of regions supported by WUYING Workspace.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifyVirtualBridgeLevelRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyVirtualBridgeLevelRequest) GoString() string {
	return s.String()
}

func (s *ModifyVirtualBridgeLevelRequest) GetAutoPay() *bool {
	return s.AutoPay
}

func (s *ModifyVirtualBridgeLevelRequest) GetAutoRenew() *bool {
	return s.AutoRenew
}

func (s *ModifyVirtualBridgeLevelRequest) GetBridgeId() *string {
	return s.BridgeId
}

func (s *ModifyVirtualBridgeLevelRequest) GetBridgeLevel() *string {
	return s.BridgeLevel
}

func (s *ModifyVirtualBridgeLevelRequest) GetPaidCallBackUrl() *string {
	return s.PaidCallBackUrl
}

func (s *ModifyVirtualBridgeLevelRequest) GetPeriod() *int32 {
	return s.Period
}

func (s *ModifyVirtualBridgeLevelRequest) GetPeriodUnit() *string {
	return s.PeriodUnit
}

func (s *ModifyVirtualBridgeLevelRequest) GetPromotionId() *string {
	return s.PromotionId
}

func (s *ModifyVirtualBridgeLevelRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyVirtualBridgeLevelRequest) SetAutoPay(v bool) *ModifyVirtualBridgeLevelRequest {
	s.AutoPay = &v
	return s
}

func (s *ModifyVirtualBridgeLevelRequest) SetAutoRenew(v bool) *ModifyVirtualBridgeLevelRequest {
	s.AutoRenew = &v
	return s
}

func (s *ModifyVirtualBridgeLevelRequest) SetBridgeId(v string) *ModifyVirtualBridgeLevelRequest {
	s.BridgeId = &v
	return s
}

func (s *ModifyVirtualBridgeLevelRequest) SetBridgeLevel(v string) *ModifyVirtualBridgeLevelRequest {
	s.BridgeLevel = &v
	return s
}

func (s *ModifyVirtualBridgeLevelRequest) SetPaidCallBackUrl(v string) *ModifyVirtualBridgeLevelRequest {
	s.PaidCallBackUrl = &v
	return s
}

func (s *ModifyVirtualBridgeLevelRequest) SetPeriod(v int32) *ModifyVirtualBridgeLevelRequest {
	s.Period = &v
	return s
}

func (s *ModifyVirtualBridgeLevelRequest) SetPeriodUnit(v string) *ModifyVirtualBridgeLevelRequest {
	s.PeriodUnit = &v
	return s
}

func (s *ModifyVirtualBridgeLevelRequest) SetPromotionId(v string) *ModifyVirtualBridgeLevelRequest {
	s.PromotionId = &v
	return s
}

func (s *ModifyVirtualBridgeLevelRequest) SetRegionId(v string) *ModifyVirtualBridgeLevelRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyVirtualBridgeLevelRequest) Validate() error {
	return dara.Validate(s)
}
